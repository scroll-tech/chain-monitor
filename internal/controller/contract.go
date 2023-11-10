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
	l1Client          *rpc.Client
	l2Client          *rpc.Client
	conf              config.Config
	eventGatherLogic  *events.EventGather
	contractsLogic    *contracts.Contracts
	checker           *checker.Checker
	messageMatchLogic *message_match.MessageMatchLogic

	l1EventCategoryList []types.TxEventCategory
	l2EventCategoryList []types.TxEventCategory
}

func NewContractController(conf config.Config, db *gorm.DB, l1Client, l2Client *rpc.Client) *ContractController {
	c := &ContractController{
		l1Client:          l1Client,
		l2Client:          l2Client,
		conf:              conf,
		eventGatherLogic:  events.NewEventGather(),
		contractsLogic:    contracts.NewContracts(ethclient.NewClient(l1Client), ethclient.NewClient(l2Client)),
		checker:           checker.NewChecker(db),
		messageMatchLogic: message_match.NewTransactionsMatchLogic(db),
	}

	if err := c.contractsLogic.Register(c.conf); err != nil {
		log.Crit("contract register failure", "error", err)
		return nil
	}

	// eth balance is checked by other means.
	c.l1EventCategoryList = append(c.l1EventCategoryList, types.ERC20EventCategory)
	c.l1EventCategoryList = append(c.l1EventCategoryList, types.ERC721EventCategory)
	c.l1EventCategoryList = append(c.l1EventCategoryList, types.ERC1155EventCategory)

	c.l2EventCategoryList = append(c.l2EventCategoryList, types.ERC20EventCategory)
	c.l2EventCategoryList = append(c.l2EventCategoryList, types.ERC721EventCategory)
	c.l2EventCategoryList = append(c.l2EventCategoryList, types.ERC1155EventCategory)

	return c
}

// Watch the l1/l2 events, contains gateways events, transfer events, messenger events
func (c *ContractController) Watch(ctx context.Context) error {
	go c.WatcherStart(ctx, ethclient.NewClient(c.l1Client), types.Layer1, c.conf.L1Config.Confirm)
	go c.WatcherStart(ctx, ethclient.NewClient(c.l2Client), types.Layer2, c.conf.L2Config.Confirm)
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

func (c *ContractController) l1Watch(ctx context.Context, start uint64, end uint64) {
	opts := bind.FilterOpts{
		Start:   start,
		End:     &end,
		Context: ctx,
	}

	messengerIterList, err := c.contractsLogic.Iterator(ctx, &opts, types.Layer2, types.MessengerEventCategory)
	if err != nil {
		log.Error("get gateway related transfer events failed", "layer", types.Layer1, "eventCategory", types.MessengerEventCategory, "error", err)
		return
	}
	messengerEvents := c.eventGatherLogic.Dispatch(ctx, types.Layer2, types.MessengerEventCategory, messengerIterList)
	if messengerEvents == nil {
		log.Info("dispatch messenger events returns empty data", "layer", types.Layer2, "eventCategory", types.MessengerEventCategory)
	}

	for _, eventCategory := range c.l1EventCategoryList {
		wrapIterList, err := c.contractsLogic.Iterator(ctx, &opts, types.Layer1, eventCategory)
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
		gatewayEvents := c.eventGatherLogic.Dispatch(ctx, types.Layer1, eventCategory, wrapIterList)
		if gatewayEvents == nil {
			log.Error("event gather deal event return empty data", "layer", types.Layer1, "eventCategory", eventCategory)
			continue
		}

		// match transfer event
		if checkErr := c.checker.GatewayCheck(ctx, eventCategory, gatewayEvents, messengerEvents, transferEvents); checkErr != nil {
			log.Error("event matcher deal failed", "layer", types.Layer1, "eventCategory", eventCategory, "error", checkErr)
			continue
		}
	}
}

func (c *ContractController) l2Watch(ctx context.Context, start uint64, end uint64) {
	opts := bind.FilterOpts{
		Start:   start,
		End:     &end,
		Context: ctx,
	}

	messengerIterList, err := c.contractsLogic.Iterator(ctx, &opts, types.Layer2, types.MessengerEventCategory)
	if err != nil {
		log.Error("get gateway related transfer events failed", "layer", types.Layer1, "eventCategory", types.MessengerEventCategory, "error", err)
		return
	}
	messengerEvents := c.eventGatherLogic.Dispatch(ctx, types.Layer2, types.MessengerEventCategory, messengerIterList)
	if messengerEvents == nil {
		log.Info("dispatch messenger events returns empty data", "layer", types.Layer2, "eventCategory", types.MessengerEventCategory)
	}

	for _, eventCategory := range c.l2EventCategoryList {
		wrapIterList, err := c.contractsLogic.Iterator(ctx, &opts, types.Layer2, eventCategory)
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
		gatewayEvents := c.eventGatherLogic.Dispatch(ctx, types.Layer2, eventCategory, wrapIterList)
		if gatewayEvents == nil {
			log.Info("dispatch gateway events returns empty data", "layer", types.Layer2, "eventCategory", eventCategory)
			continue
		}

		// match transfer event
		if checkErr := c.checker.GatewayCheck(ctx, eventCategory, gatewayEvents, messengerEvents, transferEvents); checkErr != nil {
			log.Error("event matcher deal failed", "layer", types.Layer2, "eventCategory", eventCategory, "error", checkErr)
			continue
		}
	}

	withdrawRootsMap, err := utils.GetL2WithdrawRootsInRange(ctx, c.l2Client, c.conf.L2Config.L2Contracts.MessageQueue, start, end)
	if err != nil {
		log.Error("get l2 withdraw roots in range failed", "message queue addr", c.conf.L2Config.L2Contracts.MessageQueue, "start", start, "end", end, "error", err)
		return
	}

	if checkErr := c.checker.CheckL2WithdrawRoots(ctx, start, end, messengerEvents, withdrawRootsMap); checkErr != nil {
		log.Error("check withdraw roots failed", "layer", types.Layer2, "error", checkErr)
		return
	}
}
