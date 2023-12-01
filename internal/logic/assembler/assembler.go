package assembler

import (
	"context"
	"math"

	"github.com/scroll-tech/go-ethereum/common"
	"github.com/scroll-tech/go-ethereum/rpc"
	"gorm.io/gorm"

	"github.com/scroll-tech/chain-monitor/internal/logic/events"
	"github.com/scroll-tech/chain-monitor/internal/orm"
	"github.com/scroll-tech/chain-monitor/internal/types"
)

type messageEventKey struct {
	TxHash   common.Hash
	LogIndex uint
}

// MessageMatchAssembler is a structure that helps in verifying the data integrity
// in the blockchain by checking the message matches and the events.
type MessageMatchAssembler struct {
	messengerMessageMatchOrm *orm.MessengerMessageMatch

	transferMatcher *TransferEventMatcher
}

// NewMessageMatchAssembler returns a new message match instance.
func NewMessageMatchAssembler(db *gorm.DB) *MessageMatchAssembler {
	return &MessageMatchAssembler{
		messengerMessageMatchOrm: orm.NewMessengerMessageMatch(db),
		transferMatcher:          NewTransferEventMatcher(),
	}
}

// GatewayCheck checks the gateway events.
func (c *MessageMatchAssembler) GatewayCheck(eventCategory types.EventCategory, gatewayEvents, messengerEvents, transferEvents []events.EventUnmarshaler) ([]orm.GatewayMessageMatch, error) {
	switch eventCategory {
	case types.ERC20EventCategory:
		return c.erc20EventMessageMatchAssembler(gatewayEvents, messengerEvents, transferEvents)
	case types.ERC721EventCategory:
		return c.erc721EventMessageMatchAssembler(gatewayEvents, messengerEvents, transferEvents)
	case types.ERC1155EventCategory:
		return c.erc1155EventMessageMatchAssembler(gatewayEvents, messengerEvents, transferEvents)
	}
	return nil, nil
}

// CheckL2WithdrawRoots checks the L2 withdraw roots.
func (c *MessageMatchAssembler) CheckL2WithdrawRoots(ctx context.Context, startBlockNumber, endBlockNumber uint64, client *rpc.Client, messageQueueAddr common.Address) (*orm.MessengerMessageMatch, error) {
	return c.checkL2WithdrawRoots(ctx, startBlockNumber, endBlockNumber, client, messageQueueAddr)
}

// MessengerCheck checks the messenger events.
func (c *MessageMatchAssembler) MessengerCheck(messengerEvents []events.EventUnmarshaler) ([]orm.MessengerMessageMatch, error) {
	return c.messengerMessageMatchAssembler(messengerEvents)
}

func (c *MessageMatchAssembler) findNextMessageEvent(txHash common.Hash, logIndex uint, messageHashes map[messageEventKey]common.Hash) (common.Hash, bool) {
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

func (c *MessageMatchAssembler) findPrevMessageEvent(txHash common.Hash, logIndex uint, messageHashes map[messageEventKey]common.Hash) (common.Hash, bool) {
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
