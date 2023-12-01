package crosschain

import (
	"github.com/scroll-tech/chain-monitor/internal/orm"
	"github.com/scroll-tech/chain-monitor/internal/types"
)

// MessengerCrossEventMatcher is a utility struct used for verifying the consistency of messenger events across different blockchain layers (L1 and L2).
type MessengerCrossEventMatcher struct {
	eventMatchMap map[types.EventType]types.EventType
}

// NewMessengerCrossEventMatcher initializes a new instance of MessengerCrossEventMatcher.
func NewMessengerCrossEventMatcher() *MessengerCrossEventMatcher {
	c := &MessengerCrossEventMatcher{
		eventMatchMap: make(map[types.EventType]types.EventType),
	}

	c.eventMatchMap[types.L2RelayedMessage] = types.L1SentMessage
	c.eventMatchMap[types.L1RelayedMessage] = types.L2SentMessage

	return c
}

// MessengerCrossChainCheck checks the cross chain events.
func (c *MessengerCrossEventMatcher) MessengerCrossChainCheck(layer types.LayerType, messageMatch *orm.MessengerMessageMatch) types.MismatchType {
	switch layer {
	case types.Layer1:
		return c.checkL1EventAndAmountMatchL2(messageMatch)
	case types.Layer2:
		return c.checkL2EventAndAmountMatchL1(messageMatch)
	}
	return types.MismatchTypeValid
}

// checkL1EventAndAmountMatchL2 checks that every L1FinalizeWithdraw/L1RelayedMessage has a corresponding L2 event.
func (c *MessengerCrossEventMatcher) checkL1EventAndAmountMatchL2(messageMatch *orm.MessengerMessageMatch) types.MismatchType {
	if !c.checkL1EventMatchL2(messageMatch) {
		return types.MismatchTypeL1EventNotMatch
	}
	return types.MismatchTypeValid
}

func (c *MessengerCrossEventMatcher) checkL1EventMatchL2(messageMatch *orm.MessengerMessageMatch) bool {
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
	return true
}

// checkL2EventAndAmountMatchL1  checks that every L2FinalizeDeposit/L2RelayedMessage has a corresponding L1 event.
func (c *MessengerCrossEventMatcher) checkL2EventAndAmountMatchL1(messageMatch *orm.MessengerMessageMatch) types.MismatchType {
	if !c.checkL2EventMatchL1(messageMatch) {
		return types.MismatchTypeL2EventNotMatch
	}
	return types.MismatchTypeValid
}

// checkL2EventMatchL1 checks that every L2FinalizeDeposit/L2RelayedMessage has a corresponding L1 event.
func (c *MessengerCrossEventMatcher) checkL2EventMatchL1(messageMatch *orm.MessengerMessageMatch) bool {
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

	return true
}
