package controller

import (
	"context"

	"github.com/scroll-tech/go-ethereum/accounts/abi/bind"
	"github.com/scroll-tech/go-ethereum/ethclient"
	"github.com/scroll-tech/go-ethereum/log"
	"gorm.io/gorm"

	"github.com/scroll-tech/chain-monitor/internal/config"
	"github.com/scroll-tech/chain-monitor/internal/logic/checker"
	"github.com/scroll-tech/chain-monitor/internal/logic/contracts"
	"github.com/scroll-tech/chain-monitor/internal/logic/events"
	"github.com/scroll-tech/chain-monitor/internal/types"
)

type ContractController struct {
	conf             config.Config
	eventGatherLogic *events.EventGather
	contractsLogic   *contracts.Contracts
	checker          *checker.Checker

	l1EventCategoryList []types.TxEventCategory
	l2EventCategoryList []types.TxEventCategory
}

func NewContractController(conf config.Config, db *gorm.DB, client *ethclient.Client) *ContractController {
	c := &ContractController{
		conf:             conf,
		eventGatherLogic: events.NewEventGather(),
		contractsLogic:   contracts.NewContracts(client),
		checker:          checker.NewChecker(db),
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
func (c *ContractController) Watch() error {
	// todo 怎么处理 start end
	return nil
}

// @param Start of the queried range
// @param End of the range (nil = latest)
func (c *ContractController) l1Watch(ctx context.Context, start uint64, end *uint64) {
	for _, eventCategory := range c.l1EventCategoryList {
		opt := bind.FilterOpts{
			Start:   start,
			End:     end,
			Context: ctx,
		}

		wrapIterList, err := c.contractsLogic.Iterator(ctx, &opt, types.Layer1, eventCategory)
		if err != nil {
			log.Error("get contract iterator failed", "layer", types.Layer1, "eventCategory", eventCategory, "error", err)
			continue
		}

		transferEvents, err := c.contractsLogic.GetGatewayTransfer(ctx, &opt, types.Layer1, eventCategory)
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

func (c *ContractController) l2Watch(ctx context.Context, start uint64, end *uint64) {
	for _, eventCategory := range c.l2EventCategoryList {
		opt := bind.FilterOpts{
			Start:   start,
			End:     end,
			Context: ctx,
		}

		wrapIterList, err := c.contractsLogic.Iterator(ctx, &opt, types.Layer2, eventCategory)
		if err != nil {
			log.Error("get contract iterator failed", "layer", types.Layer2, "eventCategory", eventCategory, "error", err)
			continue
		}

		transferEvents, err := c.contractsLogic.GetGatewayTransfer(ctx, &opt, types.Layer2, eventCategory)
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
}
