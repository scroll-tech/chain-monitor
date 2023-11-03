package checker

import (
	"github.com/scroll-tech/chain-monitor/internal/orm"
	"github.com/scroll-tech/chain-monitor/internal/types"
	"github.com/shopspring/decimal"
)

type CrossEventMatcher struct {
	eventMatchMap map[types.EventType]types.EventType
}

func NewCrossEventMatcher() *CrossEventMatcher {
	c := &CrossEventMatcher{
		eventMatchMap: make(map[types.EventType]types.EventType),
	}

	c.eventMatchMap[types.L1DepositETH] = types.L2FinalizeDepositETH
	c.eventMatchMap[types.L1FinalizeWithdrawETH] = types.L2WithdrawETH

	c.eventMatchMap[types.L1DepositERC20] = types.L2FinalizeDepositERC20
	c.eventMatchMap[types.L1FinalizeWithdrawERC20] = types.L2WithdrawERC20

	// add others

	return c
}

func (c *CrossEventMatcher) L1EventMatchL2(transactionMatch orm.TransactionMatch) bool {
	if transactionMatch.L2EventType == 0 {
		return false
	}

	if transactionMatch.L2Amount == decimal.NewFromInt(0) {
		return false
	}

	if transactionMatch.L2TxHash == "" {
		return false
	}

	if transactionMatch.L2BlockNumber == 0 {
		return false
	}

	return true
}

func (c *CrossEventMatcher) L2EventMatchL1(transactionMatch orm.TransactionMatch) bool {
	if transactionMatch.L1EventType == 0 {
		return false
	}

	if transactionMatch.L1Amount == decimal.NewFromInt(0) {
		return false
	}

	if transactionMatch.L1TxHash == "" {
		return false
	}

	if transactionMatch.L1BlockNumber == 0 {
		return false
	}

	return true
}

func (c *CrossEventMatcher) CrossChainAmountMatch(transactionMatch orm.TransactionMatch) bool {
	// how to check erc1155
	// todo need calculate the refund value to eth
	return transactionMatch.L2Amount == transactionMatch.L1Amount
}

func (c *CrossEventMatcher) EventTypeMatch(transactionMatch orm.TransactionMatch) bool {
	checkType, ok := c.eventMatchMap[types.EventType(transactionMatch.L1EventType)]
	if !ok {
		return false
	}

	return checkType == types.EventType(transactionMatch.L2EventType)
}
