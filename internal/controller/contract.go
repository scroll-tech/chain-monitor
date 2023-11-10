package controller

import (
	"context"
	"fmt"

	"github.com/scroll-tech/go-ethereum/accounts/abi/bind"
	"github.com/scroll-tech/go-ethereum/ethclient"
	"github.com/scroll-tech/go-ethereum/log"
	"github.com/scroll-tech/go-ethereum/rpc"
	"gorm.io/gorm"

	"github.com/scroll-tech/chain-monitor/internal/config"
	"github.com/scroll-tech/chain-monitor/internal/logic/checker"
	"github.com/scroll-tech/chain-monitor/internal/logic/contracts"
	"github.com/scroll-tech/chain-monitor/internal/logic/events"
	"github.com/scroll-tech/chain-monitor/internal/logic/message_match"
	"github.com/scroll-tech/chain-monitor/internal/types"
	"github.com/scroll-tech/chain-monitor/internal/utils"
)

const maxBlockFetchSize uint64 = 200

type ContractController struct {
	l1ETHClient       *ethclient.Client
	l2ETHClient       *ethclient.Client
	conf              config.Config
	eventGatherLogic  *events.EventGather
	contractsLogic    *contracts.Contracts
	checker           *checker.Checker
	messageMatchLogic *message_match.MessageMatchLogic

	l1EventCategoryList []types.TxEventCategory
	l2EventCategoryList []types.TxEventCategory
}

func NewContractController(conf config.Config, db *gorm.DB, l1RPCClient, l2RPCClient *rpc.Client) *ContractController {
	l1ETHClient := ethclient.NewClient(l1RPCClient)
	l2ETHClient := ethclient.NewClient(l2RPCClient)
	c := &ContractController{
		l1ETHClient:       l1ETHClient,
		l2ETHClient:       l2ETHClient,
		conf:              conf,
		eventGatherLogic:  events.NewEventGather(),
		contractsLogic:    contracts.NewContracts(l1RPCClient, l2RPCClient),
		checker:           checker.NewChecker(db),
		messageMatchLogic: message_match.NewTransactionsMatchLogic(db),
	}

	if err := c.contractsLogic.Register(c.conf); err != nil {
		log.Crit("contract register failure", "error", err)
		return nil
	}

	c.l1EventCategoryList = append(c.l1EventCategoryList, types.ERC20EventCategory)
	c.l1EventCategoryList = append(c.l1EventCategoryList, types.ERC721EventCategory)
	c.l1EventCategoryList = append(c.l1EventCategoryList, types.ERC1155EventCategory)
	c.l1EventCategoryList = append(c.l1EventCategoryList, types.ETHEventCategory)
	c.l1EventCategoryList = append(c.l1EventCategoryList, types.MessengerEventCategory)

	c.l2EventCategoryList = append(c.l2EventCategoryList, types.ERC20EventCategory)
	c.l2EventCategoryList = append(c.l2EventCategoryList, types.ERC721EventCategory)
	c.l2EventCategoryList = append(c.l2EventCategoryList, types.ERC1155EventCategory)
	c.l2EventCategoryList = append(c.l2EventCategoryList, types.ETHEventCategory)
	c.l2EventCategoryList = append(c.l2EventCategoryList, types.MessengerEventCategory)

	return c
}

// Watch the l1/l2 events, contains gateways events, transfer events, messenger events
func (c *ContractController) Watch(ctx context.Context) error {
	go c.WatcherStart(ctx, c.l1ETHClient, types.Layer1, c.conf.L1Config.Confirm)
	go c.WatcherStart(ctx, c.l2ETHClient, types.Layer2, c.conf.L2Config.Confirm)
	return nil
}

func (c *ContractController) WatcherStart(ctx context.Context, client *ethclient.Client, layer types.LayerType, confirmation rpc.BlockNumber) {
	// 1. get the max l1_number and l2_number
	blockNumberInDB, err := c.messageMatchLogic.GetLatestBlockNumber(ctx, layer)
	if err != nil {
		log.Error("ContractController.Watch get latest block number failed", "layer", layer, "err", err)
		return
	}
	start := blockNumberInDB

	// 2. get latest chain confirmation number
	confirmationNumber, err := utils.GetLatestConfirmedBlockNumber(ctx, client, confirmation)
	if err != nil {
		log.Error("ContractController.Watch get latest confirmation block number failed", "err", err)
		return
	}

	if start >= confirmationNumber {
		err := fmt.Errorf("ContractController.Watch startBlockNumber:%d >= l1ConfirmationNumber:%d", blockNumberInDB, confirmationNumber)
		log.Error("l1Watcher starting", "err", err)
		return
	}

	// 3. get the max fetch number
	end := start + maxBlockFetchSize
	if start+maxBlockFetchSize > confirmationNumber {
		end = confirmationNumber
	}

	switch layer {
	case types.Layer1:
		c.l1Watch(ctx, start, end)
	case types.Layer2:
		c.l2Watch(ctx, start, end)
	}
}

// @param Start of the queried range
// @param End of the range (nil = latest)
func (c *ContractController) l1Watch(ctx context.Context, start uint64, end uint64) {
	for _, eventCategory := range c.l1EventCategoryList {
		opt := bind.FilterOpts{
			Start:   start,
			End:     &end,
			Context: ctx,
		}

		wrapIterList, err := c.contractsLogic.Iterator(ctx, &opt, types.Layer1, eventCategory)
		if err != nil {
			log.Error("get contract iterator failed", "layer", types.Layer1, "eventCategory", eventCategory, "error", err)
			continue
		}

		transferEvents, err := c.contractsLogic.GetGatewayTransfer(ctx, start, end, types.Layer1, eventCategory)
		if err != nil {
			log.Error("get gateway related transfer events failed", "layer", types.Layer1, "eventCategory", eventCategory, "error", err)
			continue
		}

		// parse the gateway and messenger event data
		eventDataList := c.eventGatherLogic.Dispatch(ctx, types.Layer1, eventCategory, wrapIterList)
		if eventDataList == nil {
			log.Error("event gather deal event return empty data", "layer", types.Layer1, "eventCategory", eventCategory)
			continue
		}

		// match transfer event
		if checkErr := c.checker.GatewayCheck(ctx, eventCategory, eventDataList, transferEvents); checkErr != nil {
			log.Error("event matcher deal failed", "layer", types.Layer1, "eventCategory", eventCategory, "error", checkErr)
			continue
		}
	}
}

func (c *ContractController) l2Watch(ctx context.Context, start uint64, end uint64) {
	for _, eventCategory := range c.l2EventCategoryList {
		opt := bind.FilterOpts{
			Start:   start,
			End:     &end,
			Context: ctx,
		}

		wrapIterList, err := c.contractsLogic.Iterator(ctx, &opt, types.Layer2, eventCategory)
		if err != nil {
			log.Error("get contract iterator failed", "layer", types.Layer2, "eventCategory", eventCategory, "error", err)
			continue
		}

		transferEvents, err := c.contractsLogic.GetGatewayTransfer(ctx, start, end, types.Layer2, eventCategory)
		if err != nil {
			log.Error("get gateway related transfer events failed", "layer", types.Layer2, "eventCategory", eventCategory, "error", err)
			continue
		}

		// parse the event data
		eventDataList := c.eventGatherLogic.Dispatch(ctx, types.Layer2, eventCategory, wrapIterList)
		if eventDataList == nil {
			log.Error("event gather deal event return empty data", "layer", types.Layer2, "eventCategory", eventCategory)
			continue
		}

		// match transfer event
		if checkErr := c.checker.GatewayCheck(ctx, eventCategory, eventDataList, transferEvents); checkErr != nil {
			log.Error("event matcher deal failed", "layer", types.Layer2, "eventCategory", eventCategory, "error", checkErr)
			continue
		}
	}

	sentMessageEventsMap, withdrawRootsMap, err := c.contractsLogic.GetL2SentMessageEventsAndWithdrawRoots(ctx, start, end)
	if err != nil {
		log.Error("get messenger sent message events failed", "layer", types.Layer2, "error", err)
		return
	}

	if checkErr := c.checker.CheckL2WithdrawRoots(ctx, start, end, sentMessageEventsMap, withdrawRootsMap); checkErr != nil {
		log.Error("check withdraw roots failed", "layer", types.Layer2, "error", checkErr)
		return
	}
}
