package controller

import (
	"context"
	"fmt"

	"github.com/scroll-tech/go-ethereum/accounts/abi/bind"
	"github.com/scroll-tech/go-ethereum/common"
	"github.com/scroll-tech/go-ethereum/ethclient"
	"github.com/scroll-tech/go-ethereum/log"
	"github.com/scroll-tech/go-ethereum/rpc"
	"gorm.io/gorm"

	"github.com/scroll-tech/chain-monitor/internal/config"
	"github.com/scroll-tech/chain-monitor/internal/logic/checker"
	"github.com/scroll-tech/chain-monitor/internal/logic/contracts"
	"github.com/scroll-tech/chain-monitor/internal/logic/events"
	messagematch "github.com/scroll-tech/chain-monitor/internal/logic/message_match"
	"github.com/scroll-tech/chain-monitor/internal/types"
	"github.com/scroll-tech/chain-monitor/internal/utils"
)

const maxBlockFetchSize uint64 = 200

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
	defer func() {
		if err := recover(); err != nil {
			nerr := fmt.Errorf("ContractController watch panic error: %v", err)
			log.Warn(nerr.Error())
		}
	}()

	log.Info("contract controller start successful")

	// go c.watcherStart(ctx, ethclient.NewClient(c.l1Client), types.Layer1, c.conf.L1Config.Confirm)
	go c.watcherStart(ctx, ethclient.NewClient(c.l2Client), types.Layer2, c.conf.L2Config.Confirm)
}

// Stop the contract controller
func (c *ContractController) Stop() {
	c.stopTimeoutChan <- struct{}{}
}

func (c *ContractController) watcherStart(ctx context.Context, client *ethclient.Client, layer types.LayerType, confirmation rpc.BlockNumber) {
	//defer func() {
	//	if err := recover(); err != nil {
	//		log.Warn("watcherStart panic", "error", err)
	//	}
	//}()
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
			log.Error("ContractController.Watch get latest confirmation block number failed", "layer", layer.String(), "err", err)
			return
		}

		if start >= confirmationNumber {
			log.Error("Watcher start block number >= l1ConfirmationNumber", "layer", layer.String(), "startBlockNumber", blockNumberInDB, "confirmationNumber", confirmationNumber, "err", err)
			return
		}

		// 3. get the max fetch number
		end := start + maxBlockFetchSize
		if start+maxBlockFetchSize > confirmationNumber {
			end = confirmationNumber
		}

		switch layer {
		case types.Layer1:
			// c.l1Watch(ctx, start, end)
		case types.Layer2:
			c.l2Watch(ctx, start, end)
		}
	}
}

// nolint
func (c *ContractController) l1Watch(ctx context.Context, start uint64, end uint64) {
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

	var ethMessengerEvent []events.EventUnmarshaler
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

		// filter the non-eth messenger messages
		retMessengerEvents := c.filterNonETHMessenger(eventCategory, gatewayEvents, messengerEvents)
		ethMessengerEvent = append(ethMessengerEvent, retMessengerEvents...)
	}

	// because the messengerEvents contains the non-eth messenger messages and eth messenger messages,
	// but only eth messenger messages need insert to db. so need filter the non-eth messenger messages.
	if err = c.checker.MessengerCheck(ctx, ethMessengerEvent); err != nil {
		log.Error("insert message events failed", "layer", types.Layer1, "eventCategory", types.MessengerEventCategory, "error", err)
		return
	}

	// update l1 block status
	if err := c.checker.UpdateBlockStatus(ctx, types.Layer1, start, end); err != nil {
		log.Error("update block status failed", "layer", types.Layer1, "start", start, "end", end, "error", err)
		return
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
		log.Error("get messenger iterator failed", "layer", types.Layer2, "eventCategory", types.MessengerEventCategory, "error", err)
		return
	}
	messengerEvents := c.eventGatherLogic.Dispatch(ctx, types.Layer2, types.MessengerEventCategory, messengerIterList)

	var ethMessengerEvent []events.EventUnmarshaler
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
		if checkErr := c.checker.GatewayCheck(ctx, eventCategory, gatewayEvents, messengerEvents, transferEvents); checkErr != nil {
			log.Error("event matcher deal failed", "layer", types.Layer2, "eventCategory", eventCategory, "error", checkErr)
			continue
		}

		// filter the non-eth messenger messages
		retMessengerEvents := c.filterNonETHMessenger(eventCategory, gatewayEvents, messengerEvents)
		ethMessengerEvent = append(ethMessengerEvent, retMessengerEvents...)
	}

	// because the messengerEvents contains the non-eth messenger messages and eth messenger messages,
	// but only eth messenger messages need insert to db. so need filter the non-eth messenger messages.
	if err = c.checker.MessengerCheck(ctx, ethMessengerEvent); err != nil {
		log.Error("insert message events failed", "layer", types.Layer2, "eventCategory", types.MessengerEventCategory, "error", err)
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

func (c *ContractController) filterNonETHMessenger(eventCategory types.EventCategory, gatewayEvents, messengerEvents []events.EventUnmarshaler) []events.EventUnmarshaler {
	gatewayEventMap := make(map[common.Hash]struct{})
	for _, gatewayEvent := range gatewayEvents {
		switch eventCategory {
		case types.ERC20EventCategory:
			event := gatewayEvent.(*events.ERC20GatewayEventUnmarshaler)
			log.Info("xxx", "xxx", event.MessageHash.Hex())
			gatewayEventMap[event.MessageHash] = struct{}{}
		case types.ERC721EventCategory:
			event := gatewayEvent.(*events.ERC721GatewayEventUnmarshaler)
			log.Info("xxx", "xxx", event.MessageHash.Hex())
			gatewayEventMap[event.MessageHash] = struct{}{}
		case types.ERC1155EventCategory:
			event := gatewayEvent.(*events.ERC1155GatewayEventUnmarshaler)
			log.Info("xxx", "xxx", event.MessageHash.Hex())
			gatewayEventMap[event.MessageHash] = struct{}{}
		}
	}

	var ret []events.EventUnmarshaler
	for _, messengerEvent := range messengerEvents {
		event := messengerEvent.(*events.MessengerEventUnmarshaler)
		if _, ok := gatewayEventMap[event.MessageHash]; ok {
			continue
		}
		ret = append(ret, messengerEvent)
	}

	return ret
}
