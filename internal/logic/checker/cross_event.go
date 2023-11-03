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

	c.eventMatchMap[types.L1DepositWETH] = types.L2FinalizeDepositWETH
	c.eventMatchMap[types.L1FinalizeWithdrawWETH] = types.L2WithdrawWETH

	c.eventMatchMap[types.L1DepositDAI] = types.L2FinalizeDepositDAI
	c.eventMatchMap[types.L1FinalizeWithdrawDAI] = types.L2WithdrawDAI

	c.eventMatchMap[types.L1DepositCustomERC20] = types.L2FinalizeDepositCustomERC20
	c.eventMatchMap[types.L1FinalizeWithdrawCustomERC20] = types.L2WithdrawCustomERC20

	c.eventMatchMap[types.L1DepositStandardERC20] = types.L2FinalizeDepositStandardERC20
	c.eventMatchMap[types.L1FinalizeWithdrawStandardERC20] = types.L2WithdrawStandardERC20

	c.eventMatchMap[types.L1LIDODepositERC20] = types.L2LIDOFinalizeDepositERC20
	c.eventMatchMap[types.L1LIDOFinalizeWithdrawERC20] = types.L2LIDOWithdrawERC20

	c.eventMatchMap[types.L1USDCDepositERC20] = types.L2USDCFinalizeDepositERC20
	c.eventMatchMap[types.L1USDCFinalizeWithdrawERC20] = types.L2USDCWithdrawERC20

	// add others

	return c
}

func (c *CrossEventMatcher) L1EventMatchL2(transactionMatch orm.TransactionMatch) bool {
	if transactionMatch.L2EventType == 0 {
		return false
	}

	if transactionMatch.L2Value == decimal.NewFromInt(0) {
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

	if transactionMatch.L1Value == decimal.NewFromInt(0) {
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
	return transactionMatch.L2Value == transactionMatch.L1Value
}

func (c *CrossEventMatcher) EventTypeMatch(transactionMatch orm.TransactionMatch) bool {
	checkType, ok := c.eventMatchMap[types.EventType(transactionMatch.L1EventType)]
	if !ok {
		return false
	}

	return checkType == types.EventType(transactionMatch.L2EventType)
}
