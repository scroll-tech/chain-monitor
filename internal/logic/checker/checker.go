package checker

import (
	"context"
	"github.com/scroll-tech/go-ethereum/rpc"
	"math"

	"github.com/scroll-tech/go-ethereum/common"
	"gorm.io/gorm"

	"github.com/scroll-tech/chain-monitor/internal/logic/events"
	"github.com/scroll-tech/chain-monitor/internal/orm"
	"github.com/scroll-tech/chain-monitor/internal/types"
)

type messageEventKey struct {
	TxHash   common.Hash
	LogIndex uint
}

// Checker is a structure that helps in verifying the data integrity
// in the blockchain by checking the message matches and the events.
type Checker struct {
	messengerMessageMatchOrm *orm.MessengerMessageMatch

	transferMatcher   *TransferEventMatcher
	crossChainMatcher *CrossEventMatcher
}

// NewChecker returns a new Checker instance.
func NewChecker(db *gorm.DB) *Checker {
	return &Checker{
		messengerMessageMatchOrm: orm.NewMessengerMessageMatch(db),
		transferMatcher:          NewTransferEventMatcher(),
		crossChainMatcher:        NewCrossEventMatcher(),
	}
}

// CrossChainCheck checks the cross chain events.
func (c *Checker) CrossChainCheck(layer types.LayerType, messageMatch orm.MessageMatch) types.MismatchType {
	switch layer {
	case types.Layer1:
		return c.crossChainMatcher.checkL1EventAndAmountMatchL2(messageMatch)
	case types.Layer2:
		return c.crossChainMatcher.checkL2EventAndAmountMatchL1(messageMatch)
	}
	return types.MismatchTypeValid
}

// GatewayCheck checks the gateway events.
func (c *Checker) GatewayCheck(eventCategory types.EventCategory, gatewayEvents, messengerEvents, transferEvents []events.EventUnmarshaler) ([]orm.GatewayMessageMatch, error) {
	switch eventCategory {
	case types.ERC20EventCategory:
		return c.erc20EventUnmarshaler(gatewayEvents, messengerEvents, transferEvents)
	case types.ERC721EventCategory:
		return c.erc721EventUnmarshaler(gatewayEvents, messengerEvents, transferEvents)
	case types.ERC1155EventCategory:
		return c.erc1155EventUnmarshaler(gatewayEvents, messengerEvents, transferEvents)
	}
	return nil, nil
}

// CheckL2WithdrawRoots checks the L2 withdraw roots.
func (c *Checker) CheckL2WithdrawRoots(ctx context.Context, startBlockNumber, endBlockNumber uint64, client *rpc.Client, messageQueueAddr common.Address) (*orm.MessengerMessageMatch, error) {
	return c.checkL2WithdrawRoots(ctx, startBlockNumber, endBlockNumber, client, messageQueueAddr)
}

// MessengerCheck checks the messenger events.
func (c *Checker) MessengerCheck(messengerEvents []events.EventUnmarshaler) ([]orm.MessengerMessageMatch, error) {
	return c.messengerCheck(messengerEvents)
}

func (c *Checker) findNextMessageEvent(txHash common.Hash, logIndex uint, messageHashes map[messageEventKey]common.Hash) (common.Hash, bool) {
	var nextMessageHash common.Hash
	var found bool
	var smallestDiff uint = math.MaxUint
	for key, msgHash := range messageHashes {
		if key.TxHash == txHash && key.LogIndex > logIndex {
			if diff := key.LogIndex - logIndex; diff < smallestDiff {
				smallestDiff = diff
				nextMessageHash = msgHash
				found = true
			}
		}
	}
	return nextMessageHash, found
}

func (c *Checker) findPrevMessageEvent(txHash common.Hash, logIndex uint, messageHashes map[messageEventKey]common.Hash) (common.Hash, bool) {
	var prevMessageHash common.Hash
	var found bool
	var smallestDiff uint = math.MaxUint
	for key, msgHash := range messageHashes {
		if key.TxHash == txHash && key.LogIndex < logIndex {
			if diff := logIndex - key.LogIndex; diff < smallestDiff {
				smallestDiff = diff
				prevMessageHash = msgHash
				found = true
			}
		}
	}
	return prevMessageHash, found
}
