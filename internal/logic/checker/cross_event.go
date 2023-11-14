package checker

import (
	"github.com/scroll-tech/chain-monitor/internal/orm"
	"github.com/scroll-tech/chain-monitor/internal/types"
)

type CrossEventMatcher struct {
	eventMatchMap map[types.EventType]types.EventType
}

// NewCrossEventMatcher initializes a new instance of CrossEventMatcher.
func NewCrossEventMatcher() *CrossEventMatcher {
	c := &CrossEventMatcher{
		eventMatchMap: make(map[types.EventType]types.EventType),
	}

	c.eventMatchMap[types.L2FinalizeDepositERC20] = types.L1DepositERC20
	c.eventMatchMap[types.L1FinalizeWithdrawERC20] = types.L2WithdrawERC20

	c.eventMatchMap[types.L2FinalizeDepositERC721] = types.L1DepositERC721
	c.eventMatchMap[types.L1FinalizeWithdrawERC721] = types.L2WithdrawERC721

	c.eventMatchMap[types.L2FinalizeDepositERC1155] = types.L1DepositERC1155
	c.eventMatchMap[types.L1FinalizeWithdrawERC1155] = types.L2WithdrawERC1155

	c.eventMatchMap[types.L2FinalizeBatchDepositERC721] = types.L1BatchDepositERC721
	c.eventMatchMap[types.L1FinalizeBatchWithdrawERC721] = types.L2BatchWithdrawERC721

	c.eventMatchMap[types.L2FinalizeBatchDepositERC1155] = types.L1BatchDepositERC1155
	c.eventMatchMap[types.L1FinalizeBatchWithdrawERC1155] = types.L2BatchWithdrawERC1155

	c.eventMatchMap[types.L2RelayedMessage] = types.L1SentMessage
	c.eventMatchMap[types.L1RelayedMessage] = types.L2SentMessage

	return c
}

// check every L1FializedWithdraw/L1RelayedMessage has corresponding L2 event.
func (c *CrossEventMatcher) checkL1EventMatchL2(messageMatch orm.MessageMatch) bool {
	matchingEvent, isPresent := c.eventMatchMap[types.EventType(messageMatch.L1EventType)]
	if !isPresent {
		// If the L1 event type is not in the checklist, skip the check
		return true
	}

	if matchingEvent != types.EventType(messageMatch.L2EventType) {
		// If the matching event is not equal to the L2 event type, return false
		return false
	}

	if messageMatch.L2TxHash == "" {
		return false
	}

	if messageMatch.L2BlockNumber == 0 {
		return false
	}

	if messageMatch.L1Amount != messageMatch.L2Amount {
		return false
	}

	if messageMatch.L1TokenId != messageMatch.L2TokenId {
		return false
	}

	return true
}

// check every L2FializedDeposit/L2RelayedMessage has corresponding L1 event.
func (c *CrossEventMatcher) checkL2EventMatchL1(messageMatch orm.MessageMatch) bool {
	matchingEvent, isPresent := c.eventMatchMap[types.EventType(messageMatch.L2EventType)]
	if !isPresent {
		// If the L2 event type is not in the checklist, skip the check
		return true
	}

	if matchingEvent != types.EventType(messageMatch.L1EventType) {
		// If the matching event is not equal to the L1 event type, return false
		return false
	}

	if messageMatch.L1TxHash == "" {
		return false
	}

	if messageMatch.L1BlockNumber == 0 {
		return false
	}

	if messageMatch.L1Amount != messageMatch.L2Amount {
		return false
	}

	if messageMatch.L1TokenId != messageMatch.L2TokenId {
		return false
	}

	return true
}
