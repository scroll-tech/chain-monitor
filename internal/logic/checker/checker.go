package checker

import (
	"context"

	"github.com/scroll-tech/go-ethereum/log"
	"github.com/shopspring/decimal"
	"gorm.io/gorm"

	"github.com/scroll-tech/chain-monitor/internal/logic/events"
	"github.com/scroll-tech/chain-monitor/internal/orm"
	"github.com/scroll-tech/chain-monitor/internal/types"
)

type Checker struct {
	messageMatchOrm   *orm.MessageMatch
	transferMatcher   *TransferEventMatcher
	crossChainMatcher *CrossEventMatcher
}

func NewChecker(db *gorm.DB) *Checker {
	return &Checker{
		messageMatchOrm:   orm.NewMessageMatch(db),
		transferMatcher:   NewTransferEventMatcher(),
		crossChainMatcher: NewCrossEventMatcher(),
	}
}

func (c *Checker) CrossChainCheck(_ context.Context, layer types.LayerType, messageMatch orm.MessageMatch) types.MismatchType {
	if layer == types.Layer1 {
		if !c.crossChainMatcher.L1EventMatchL2(messageMatch) {
			return types.MismatchTypeL2EventNotExist
		}
	}

	if layer == types.Layer2 {
		if !c.crossChainMatcher.L2EventMatchL1(messageMatch) {
			return types.MismatchTypeL1EventNotExist
		}
	}

	if !c.crossChainMatcher.CrossChainAmountMatch(messageMatch) {
		return types.MismatchTypeAmount
	}

	if !c.crossChainMatcher.EventTypeMatch(messageMatch) {
		return types.MismatchTypeCrossChainTypeNotMatch
	}

	return types.MismatchTypeOk
}

func (c *Checker) GatewayCheck(ctx context.Context, eventCategory types.TxEventCategory, eventDataList []events.EventUnmarshaler) error {
	switch eventCategory {
	case types.ERC20EventCategory:
		return c.erc20EventUnmarshaler(ctx, eventDataList)
	}
	return nil
}

func (c *Checker) erc20EventUnmarshaler(ctx context.Context, eventDataList []events.EventUnmarshaler) error {
	var messageMatches []orm.MessageMatch
	erc20TransferEventMap := make(map[string]*events.ERC20GatewayEventUnmarshaler)
	erc20GatewayEventMap := make(map[string]*events.ERC20GatewayEventUnmarshaler)
	for _, eventData := range eventDataList {
		erc20EventUnmarshaler := eventData.(*events.ERC20GatewayEventUnmarshaler)
		eventType := erc20EventUnmarshaler.Type
		if erc20EventUnmarshaler.Transfer {
			erc20TransferEventMap[erc20EventUnmarshaler.TxHash.Hex()] = erc20EventUnmarshaler
		} else {
			erc20GatewayEventMap[erc20EventUnmarshaler.TxHash.Hex()] = erc20EventUnmarshaler
		}

		var tmpMessageMatch orm.MessageMatch
		if eventType == types.L1DepositERC20 || eventType == types.L1FinalizeWithdrawERC20 || eventType == types.L1RefundERC20 {
			tmpMessageMatch = orm.MessageMatch{
				TokenType:     int(types.TokenTypeERC20),
				L1EventType:   int(eventType),
				L1BlockNumber: erc20EventUnmarshaler.Number,
				L1TxHash:      erc20EventUnmarshaler.TxHash.Hex(),
				L1Amount:      decimal.NewFromBigInt(erc20EventUnmarshaler.Amount, 10),
			}
		} else if eventType == types.L2WithdrawERC20 || eventType == types.L2FinalizeDepositERC20 {
			tmpMessageMatch = orm.MessageMatch{
				TokenType:     int(types.TokenTypeERC20),
				L2EventType:   int(eventType),
				L2BlockNumber: erc20EventUnmarshaler.Number,
				L2TxHash:      erc20EventUnmarshaler.TxHash.Hex(),
				L2Amount:      decimal.NewFromBigInt(erc20EventUnmarshaler.Amount, 10),
			}
		} else {
			log.Error("unknown erc20 event type")
			continue
		}

		messageMatches = append(messageMatches, tmpMessageMatch)
	}

	effectRows, err := c.messageMatchOrm.InsertOrUpdate(ctx, messageMatches)
	if err != nil || effectRows != len(messageMatches) {
		log.Error("erc20EventUnmarshaler orm insert failed, err:%w", err)
	}

	return c.transferMatcher.Erc20Matcher(erc20TransferEventMap, erc20GatewayEventMap)
}
