package controller

import (
	"context"
	"errors"
	"fmt"
	"sync/atomic"
	"time"

	"github.com/scroll-tech/go-ethereum/accounts/abi/bind"
	"github.com/scroll-tech/go-ethereum/ethclient"
	"github.com/scroll-tech/go-ethereum/log"
	"github.com/scroll-tech/go-ethereum/rpc"
	"golang.org/x/sync/errgroup"
	"gorm.io/gorm"

	"github.com/scroll-tech/chain-monitor/internal/config"
	"github.com/scroll-tech/chain-monitor/internal/logic/checker"
	"github.com/scroll-tech/chain-monitor/internal/logic/contracts"
	"github.com/scroll-tech/chain-monitor/internal/logic/events"
	messagematch "github.com/scroll-tech/chain-monitor/internal/logic/message_match"
	"github.com/scroll-tech/chain-monitor/internal/orm"
	"github.com/scroll-tech/chain-monitor/internal/types"
	"github.com/scroll-tech/chain-monitor/internal/utils"
)

const maxBlockFetchSize uint64 = 199

var errEmptyConfirmation = errors.New("block number > ConfirmationNumber")

// ContractController is a struct that manages the interaction with contracts on Layer 1 and Layer 2.
type ContractController struct {
	l1Client          *rpc.Client
	l2Client          *rpc.Client
	conf              *config.Config
	eventGatherLogic  *events.EventGather
	contractsLogic    *contracts.Contracts
	checker           *checker.Checker
	messageMatchLogic *messagematch.LogicMessageMatch

	stopTimeoutChan     chan struct{}
	l1EventCategoryList []types.EventCategory
	l2EventCategoryList []types.EventCategory
}

// NewContractController creates a new ContractController object.
func NewContractController(conf *config.Config, db *gorm.DB, l1Client, l2Client *rpc.Client) *ContractController {
	c := &ContractController{
		l1Client:          l1Client,
		l2Client:          l2Client,
		conf:              conf,
		eventGatherLogic:  events.NewEventGather(),
		contractsLogic:    contracts.NewContracts(ethclient.NewClient(l1Client), ethclient.NewClient(l2Client)),
		checker:           checker.NewChecker(db),
		messageMatchLogic: messagematch.NewMessageMatchLogic(conf, db),
		stopTimeoutChan:   make(chan struct{}),
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

// Watch is an exported function that starts watching the Layer 1 and Layer 2 events, which include gateways events, transfer events, and messenger events.
func (c *ContractController) Watch(ctx context.Context) {
	go c.watcherStart(ctx, ethclient.NewClient(c.l1Client), types.Layer1, c.conf.L1Config.Confirm, 5)
	// l2 watcher only supports concurrency 1 (no concurrency).
	go c.watcherStart(ctx, ethclient.NewClient(c.l2Client), types.Layer2, c.conf.L2Config.Confirm, 1)
}

// Stop the contract controller
func (c *ContractController) Stop() {
	c.stopTimeoutChan <- struct{}{}
}

func (c *ContractController) watcherStart(ctx context.Context, client *ethclient.Client, layer types.LayerType, confirmation rpc.BlockNumber, concurrency int) {
	log.Info("contract controller start successful", "layer", layer.String(), "confirmation", confirmation)
	//defer func() {
	//	if err := recover(); err != nil {
	//		log.Warn("watcherStart panic", "error", err)
	//	}
	//}()
	// 1. get the max l1_number and l2_number
	blockNumberInDB, err := c.messageMatchLogic.GetLatestBlockNumber(ctx, layer)
	if err != nil {
		log.Error("ContractController.Watch get latest block number failed", "layer", layer, "err", err)
		return
	}
	start := blockNumberInDB + 1

	for {
		select {
		case <-ctx.Done():
			if ctx.Err() != nil {
				log.Error("CrossChainController proposer canceled with error", "error", ctx.Err())
			}
			return
		case <-c.stopTimeoutChan:
			log.Info("CrossChainController proposer the run loop exit")
			return
		default:
		}

		var eg errgroup.Group
		for i := 0; i < concurrency; i++ {
			eg.Go(func() error {
				// 2. get latest chain confirmation number
				confirmationNumber, err := utils.GetLatestConfirmedBlockNumber(ctx, client, confirmation)
				if err != nil {
					log.Error("ContractController.Watch get latest confirmation block number failed", "layer", layer.String(), "err", err)
					return fmt.Errorf("ContractController.Watch, :%v", err)
				}

				if start > confirmationNumber {
					log.Info("Watcher start block number > ConfirmationNumber",
						"layer", layer.String(),
						"startBlockNumber", blockNumberInDB,
						"confirmationNumber", confirmationNumber,
						"err", err,
					)
					return errEmptyConfirmation
				}

				// 3. get the max fetch number
				end := start + maxBlockFetchSize
				if start+maxBlockFetchSize > confirmationNumber {
					end = confirmationNumber
				}

				currentStart := start
				nextStart := end + 1
				atomic.StoreUint64(&start, nextStart)

				switch layer {
				case types.Layer1:
					c.l1Watch(ctx, currentStart, end)
				case types.Layer2:
					c.l2Watch(ctx, currentStart, end)
				}
				return nil
			})
		}

		if egErr := eg.Wait(); egErr != nil {
			if errors.Is(egErr, errEmptyConfirmation) {
				time.Sleep(time.Millisecond * 500)
			}
		}
	}
}

// nolint
func (c *ContractController) l1Watch(ctx context.Context, start uint64, end uint64) {
	log.Info("watching block number", "layer", types.Layer1, "start", start, "end", end)
	opts := bind.FilterOpts{
		Start:   start,
		End:     &end,
		Context: ctx,
	}

	messengerIterList, err := c.contractsLogic.Iterator(ctx, &opts, types.Layer1, types.MessengerEventCategory)
	if err != nil {
		log.Error("get messenger iterator failed", "layer", types.Layer1, "eventCategory", types.MessengerEventCategory, "error", err)
		return
	}
	messengerEvents := c.eventGatherLogic.Dispatch(ctx, types.Layer1, types.MessengerEventCategory, messengerIterList)
	messengerMessageMatches, err := c.checker.MessengerCheck(ctx, types.Layer1, messengerEvents)
	if err != nil {
		log.Error("generate messenger message match failed", "layer", types.Layer2, "eventCategory", types.MessengerEventCategory, "error", err)
		return
	}

	if len(messengerMessageMatches) == 0 {
		return
	}

	var l1GatewayMessageMatches []orm.MessageMatch
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
			log.Debug("event gather deal event return empty data", "layer", types.Layer1, "eventCategory", eventCategory)
			continue
		}

		// match transfer event
		retL1MessageMatches, checkErr := c.checker.GatewayCheck(ctx, eventCategory, gatewayEvents, messengerEvents, transferEvents)
		l1GatewayMessageMatches = append(l1GatewayMessageMatches, retL1MessageMatches...)
		if checkErr != nil {
			log.Error("event matcher deal failed", "layer", types.Layer1, "eventCategory", eventCategory, "error", checkErr)
			continue
		}
	}

	c.replaceGatewayEventInfo(types.Layer1, l1GatewayMessageMatches, messengerMessageMatches)
	if err := c.messageMatchLogic.InsertOrUpdateMessageMatches(ctx, types.Layer1, messengerMessageMatches); err != nil {
		log.Error("insert message events failed", "layer", types.Layer1, "error", err)
		return
	}

	// update l1 block status
	if err := c.checker.UpdateBlockStatus(ctx, types.Layer1, start, end); err != nil {
		log.Error("update block status failed", "layer", types.Layer1, "start", start, "end", end, "error", err)
		return
	}
}

func (c *ContractController) l2Watch(ctx context.Context, start uint64, end uint64) {
	log.Info("watching block number", "layer", types.Layer2, "start", start, "end", end)
	opts := bind.FilterOpts{
		Start:   start,
		End:     &end,
		Context: ctx,
	}

	messengerIterList, err := c.contractsLogic.Iterator(ctx, &opts, types.Layer2, types.MessengerEventCategory)
	if err != nil {
		log.Error("get messenger iterator failed", "layer", types.Layer2, "eventCategory", types.MessengerEventCategory, "error", err)
		return
	}
	messengerEvents := c.eventGatherLogic.Dispatch(ctx, types.Layer2, types.MessengerEventCategory, messengerIterList)
	messengerMessageMatches, err := c.checker.MessengerCheck(ctx, types.Layer2, messengerEvents)
	if err != nil {
		log.Error("generate messenger message match failed", "layer", types.Layer2, "eventCategory", types.MessengerEventCategory, "error", err)
		return
	}

	var l2GatewayMessageMatches []orm.MessageMatch
	for _, eventCategory := range c.l2EventCategoryList {
		var wrapIterList []types.WrapIterator
		wrapIterList, err = c.contractsLogic.Iterator(ctx, &opts, types.Layer2, eventCategory)
		if err != nil {
			log.Error("get contract iterator failed", "layer", types.Layer2, "eventCategory", eventCategory, "error", err)
			continue
		}

		var transferEvents []events.EventUnmarshaler
		transferEvents, err = c.contractsLogic.GetGatewayTransfer(ctx, start, end, types.Layer2, eventCategory)
		if err != nil {
			log.Error("get gateway related transfer events failed", "layer", types.Layer2, "eventCategory", eventCategory, "error", err)
			continue
		}

		// parse the event data
		gatewayEvents := c.eventGatherLogic.Dispatch(ctx, types.Layer2, eventCategory, wrapIterList)
		if gatewayEvents == nil {
			log.Debug("dispatch gateway events returns empty data", "layer", types.Layer2, "eventCategory", eventCategory)
			continue
		}

		// match transfer event
		retL2MessageMatches, checkErr := c.checker.GatewayCheck(ctx, eventCategory, gatewayEvents, messengerEvents, transferEvents)
		l2GatewayMessageMatches = append(l2GatewayMessageMatches, retL2MessageMatches...)
		if checkErr != nil {
			log.Error("event matcher deal failed", "layer", types.Layer2, "eventCategory", eventCategory, "error", checkErr)
			continue
		}
	}

	c.replaceGatewayEventInfo(types.Layer2, l2GatewayMessageMatches, messengerMessageMatches)
	if err = c.messageMatchLogic.InsertOrUpdateMessageMatches(ctx, types.Layer2, messengerMessageMatches); err != nil {
		log.Error("insert message events failed", "layer", types.Layer2, "error", err)
		return
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

	// update l2 block status
	if err := c.checker.UpdateBlockStatus(ctx, types.Layer2, start, end); err != nil {
		log.Error("update block status failed", "layer", types.Layer2, "start", start, "end", end, "error", err)
		return
	}
}

func (c *ContractController) replaceGatewayEventInfo(layer types.LayerType, gatewayMessages, messengerMessages []orm.MessageMatch) {
	messageHashGatewayMessageMatchMap := make(map[string]orm.MessageMatch)
	for _, gatewayMessage := range gatewayMessages {
		messageHashGatewayMessageMatchMap[gatewayMessage.MessageHash] = gatewayMessage
	}

	for i := 0; i < len(messengerMessages); i++ {
		m := messengerMessages[i]
		gatewayMessageMatch, ok := messageHashGatewayMessageMatchMap[m.MessageHash]
		if !ok {
			continue
		}

		messengerMessages[i].TokenType = gatewayMessageMatch.TokenType
		switch layer {
		case types.Layer1:
			messengerMessages[i].L1EventType = gatewayMessageMatch.L1EventType
			messengerMessages[i].L1TokenIds = gatewayMessageMatch.L1TokenIds
			messengerMessages[i].L1Amounts = gatewayMessageMatch.L1Amounts
		case types.Layer2:
			messengerMessages[i].L2EventType = gatewayMessageMatch.L2EventType
			messengerMessages[i].L2TokenIds = gatewayMessageMatch.L2TokenIds
			messengerMessages[i].L2Amounts = gatewayMessageMatch.L2Amounts
		}
	}
}
