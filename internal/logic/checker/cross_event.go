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

func (c *CrossEventMatcher) L1EventMatchL2(messageMatch orm.MessageMatch) bool {
	if messageMatch.L2EventType == 0 {
		return false
	}

	if messageMatch.L2Amount == decimal.NewFromInt(0) {
		return false
	}

	if messageMatch.L2TxHash == "" {
		return false
	}

	if messageMatch.L2BlockNumber == 0 {
		return false
	}

	return true
}

func (c *CrossEventMatcher) L2EventMatchL1(messageMatch orm.MessageMatch) bool {
	if messageMatch.L1EventType == 0 {
		return false
	}

	if messageMatch.L1Amount == decimal.NewFromInt(0) {
		return false
	}

	if messageMatch.L1TxHash == "" {
		return false
	}

	if messageMatch.L1BlockNumber == 0 {
		return false
	}

	return true
}

func (c *CrossEventMatcher) CrossChainAmountMatch(messageMatch orm.MessageMatch) bool {
	// how to check erc1155
	// todo need calculate the refund value to eth
	return messageMatch.L2Amount == messageMatch.L1Amount
}

func (c *CrossEventMatcher) EventTypeMatch(messageMatch orm.MessageMatch) bool {
	checkType, ok := c.eventMatchMap[types.EventType(messageMatch.L1EventType)]
	if !ok {
		return false
	}

	return checkType == types.EventType(messageMatch.L2EventType)
}
