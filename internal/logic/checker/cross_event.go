package checker

import (
	"math/big"
	"strings"

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

	if messageMatch.L2Amounts == "" {
		return false
	}

	if messageMatch.L2TxHash == "" {
		return false
	}

	if messageMatch.L2BlockNumber == 0 {
		return false
	}

	return c.crossChainAmountMatch(messageMatch)
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

	if messageMatch.L1Amounts == "" {
		return false
	}

	if messageMatch.L1TxHash == "" {
		return false
	}

	if messageMatch.L1BlockNumber == 0 {
		return false
	}

	return c.crossChainAmountMatch(messageMatch)
}

func (c *CrossEventMatcher) crossChainAmountMatch(messageMatch orm.MessageMatch) bool {
	var l1Amounts, l2Amounts []*big.Int
	var l1TokenIds, l2TokenIds []*big.Int

	if messageMatch.L1Amounts != "" {
		l1AmountSplits := strings.Split(messageMatch.L1Amounts, ",")
		for _, l1AmountSplit := range l1AmountSplits {
			l1Amount, ok := new(big.Int).SetString(l1AmountSplit, 0)
			if !ok {
				return false
			}
			l1Amounts = append(l1Amounts, l1Amount)
		}
	}

	if messageMatch.L1TokenIds != "" {
		l1TokenIdSplits := strings.Split(messageMatch.L1TokenIds, ",")
		for _, l1TokenIdSplit := range l1TokenIdSplits {
			l1Token, ok := new(big.Int).SetString(l1TokenIdSplit, 0)
			if !ok {
				return false
			}
			l1TokenIds = append(l1TokenIds, l1Token)
		}
	}

	if messageMatch.L2Amounts != "" {
		l2AmountSplits := strings.Split(messageMatch.L2Amounts, ",")
		for _, l2AmountSplit := range l2AmountSplits {
			l2Amount, ok := new(big.Int).SetString(l2AmountSplit, 0)
			if !ok {
				return false
			}
			l2Amounts = append(l2Amounts, l2Amount)
		}
	}

	if messageMatch.L2TokenIds != "" {
		l2TokenIDSplits := strings.Split(messageMatch.L2TokenIds, ",")
		for _, l2TokenIDSplit := range l2TokenIDSplits {
			l2TokenID, ok := new(big.Int).SetString(l2TokenIDSplit, 0)
			if !ok {
				return false
			}
			l2TokenIds = append(l2TokenIds, l2TokenID)
		}
	}

	switch types.TokenType(messageMatch.TokenType) {
	case types.TokenTypeERC20, types.TokenTypeETH:
		return len(messageMatch.L1Amounts) == len(messageMatch.L2Amounts) &&
			len(messageMatch.L2Amounts) == 1 && messageMatch.L1Amounts[0] == messageMatch.L2Amounts[0]
	case types.TokenTypeERC721:
		if len(l1TokenIds) != len(l2TokenIds) {
			return false
		}
		for l1Idx, l1TokenId := range l1TokenIds {
			l2TokenID := l2TokenIds[l1Idx]
			if l1TokenId != l2TokenID {
				return false
			}
		}
	case types.TokenTypeERC1155:
		if len(l1TokenIds) != len(l2TokenIds) || len(l1Amounts) != len(l2Amounts) || len(l1TokenIds) != len(l1Amounts) {
			return false
		}
		for l1TokenIdx, l1TokenID := range l1TokenIds {
			l2TokenID := l2TokenIds[l1TokenIdx]
			if l1TokenID != l2TokenID {
				return false
			}
			l1Amount := l1Amounts[l1TokenIdx]
			l2Amount := l2Amounts[l1TokenIdx]
			if l1Amount != l2Amount {
				return false
			}
		}
	}
	return false
}
