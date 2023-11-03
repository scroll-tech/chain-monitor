package checker

import (
	"context"
	"github.com/scroll-tech/chain-monitor/internal/logic/events"
	"github.com/scroll-tech/chain-monitor/internal/orm"
	"github.com/scroll-tech/chain-monitor/internal/types"
	"github.com/scroll-tech/go-ethereum/log"
	"github.com/shopspring/decimal"
	"gorm.io/gorm"
)

type Checker struct {
	transactionMatchOrm *orm.TransactionMatch
	transferMatcher     *TransferEventMatcher
	crossChainMatcher   *CrossEventMatcher
}

func NewChecker(db *gorm.DB) *Checker {
	return &Checker{
		transactionMatchOrm: orm.NewTransactionMatch(db),
		transferMatcher:     NewTransferEventMatcher(),
		crossChainMatcher:   NewCrossEventMatcher(),
	}
}

func (c *Checker) CrossChainCheck(_ context.Context, layer types.LayerType, transactionMatch orm.TransactionMatch) types.MismatchType {
	if layer == types.Layer1 {
		if !c.crossChainMatcher.L1EventMatchL2(transactionMatch) {
			return types.MismatchTypeL2EventNotExist
		}
	}

	if layer == types.Layer2 {
		if !c.crossChainMatcher.L2EventMatchL1(transactionMatch) {
			return types.MismatchTypeL1EventNotExist
		}
	}

	if !c.crossChainMatcher.CrossChainAmountMatch(transactionMatch) {
		return types.MismatchTypeAmount
	}

	if !c.crossChainMatcher.EventTypeMatch(transactionMatch) {
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
	var transactionMatches []orm.TransactionMatch
	erc20TransferEventMap := make(map[string]*events.ERC20GatewayEventUnmarshaler)
	erc20GatewayEventMap := make(map[string]*events.ERC20GatewayEventUnmarshaler)
	for _, eventData := range eventDataList {
		erc20EventUnmarshaler := eventData.(*events.ERC20GatewayEventUnmarshaler)
		eventType := erc20EventUnmarshaler.Type
		if erc20EventUnmarshaler.Transfer {
			erc20TransferEventMap[erc20EventUnmarshaler.TxHash] = erc20EventUnmarshaler
		} else {
			erc20GatewayEventMap[erc20EventUnmarshaler.TxHash] = erc20EventUnmarshaler
		}

		tmpTransactionMatch := orm.TransactionMatch{
			TokenType:     int(types.TokenTypeERC20),
			L1EventType:   int(eventType),
			L1BlockNumber: erc20EventUnmarshaler.Number,
			L1TxHash:      erc20EventUnmarshaler.TxHash,
			L1Value:       decimal.NewFromBigInt(erc20EventUnmarshaler.Amount, 10),
		}

		transactionMatches = append(transactionMatches, tmpTransactionMatch)
	}

	effectRows, err := c.transactionMatchOrm.InsertOrUpdate(ctx, transactionMatches)
	if err != nil || effectRows != len(transactionMatches) {
		log.Error("erc20EventUnmarshaler orm insert failed, err:%w", err)
	}

	return c.transferMatcher.Erc20Matcher(erc20TransferEventMap, erc20GatewayEventMap)
}
