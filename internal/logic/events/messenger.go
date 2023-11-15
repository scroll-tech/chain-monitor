package events

import (
	"context"
	"math/big"

	"github.com/scroll-tech/chain-monitor/internal/logic/contracts/abi/il1scrollmessenger"
	"github.com/scroll-tech/chain-monitor/internal/logic/contracts/abi/il2scrollmessenger"
	"github.com/scroll-tech/chain-monitor/internal/types"
	"github.com/scroll-tech/chain-monitor/internal/utils"
	"github.com/scroll-tech/go-ethereum/common"
)

// MessengerEventUnmarshaler represents a SentMessage event unmarshaler raised by the L1/L2 scroll messenger contract.
type MessengerEventUnmarshaler struct {
	Layer        types.LayerType
	Type         types.EventType
	Number       uint64
	TxHash       common.Hash
	Index        uint
	MessageNonce *big.Int
	Message      []byte
	MessageHash  common.Hash
	Value        *big.Int
}

func NewMessengerEventUnmarshaler() *MessengerEventUnmarshaler {
	return &MessengerEventUnmarshaler{}
}

func (e *MessengerEventUnmarshaler) Unmarshal(context context.Context, layerType types.LayerType, iterators []types.WrapIterator) []EventUnmarshaler {
	var events []EventUnmarshaler
	for _, it := range iterators {
		for it.Iter.Next() {
			events = append(events, e.messenger(layerType, it.Iter, it.EventType))
		}
	}
	return events
}

func (e *MessengerEventUnmarshaler) messenger(layerType types.LayerType, it types.Iterator, eventType types.EventType) EventUnmarshaler {
	var event EventUnmarshaler
	switch eventType {
	case types.L1SentMessage:
		iter := it.(*il1scrollmessenger.Il1scrollmessengerSentMessageIterator)
		msgHash := utils.ComputeMessageHash(iter.Event.Sender, iter.Event.Target, iter.Event.Value, iter.Event.MessageNonce, iter.Event.Message)
		event = &MessengerEventUnmarshaler{
			Layer:        layerType,
			Type:         eventType,
			Number:       iter.Event.Raw.BlockNumber,
			TxHash:       iter.Event.Raw.TxHash,
			Index:        iter.Event.Raw.Index,
			MessageNonce: iter.Event.MessageNonce,
			Message:      iter.Event.Message,
			MessageHash:  msgHash,
			Value:        iter.Event.Value,
		}
	case types.L2SentMessage:
		iter := it.(*il2scrollmessenger.Il2scrollmessengerSentMessageIterator)
		msgHash := utils.ComputeMessageHash(iter.Event.Sender, iter.Event.Target, iter.Event.Value, iter.Event.MessageNonce, iter.Event.Message)
		event = &MessengerEventUnmarshaler{
			Layer:        layerType,
			Type:         eventType,
			Number:       iter.Event.Raw.BlockNumber,
			TxHash:       iter.Event.Raw.TxHash,
			Index:        iter.Event.Raw.Index,
			MessageNonce: iter.Event.MessageNonce,
			Message:      iter.Event.Message,
			MessageHash:  msgHash,
			Value:        iter.Event.Value,
		}
	case types.L1RelayedMessage:
		iter := it.(*il1scrollmessenger.Il1scrollmessengerRelayedMessageIterator)
		event = &MessengerEventUnmarshaler{
			Layer:       layerType,
			Type:        eventType,
			Number:      iter.Event.Raw.BlockNumber,
			TxHash:      iter.Event.Raw.TxHash,
			Index:       iter.Event.Raw.Index,
			MessageHash: iter.Event.MessageHash,
		}
	case types.L2RelayedMessage:
		iter := it.(*il2scrollmessenger.Il2scrollmessengerRelayedMessageIterator)
		event = &MessengerEventUnmarshaler{
			Layer:       layerType,
			Type:        eventType,
			Number:      iter.Event.Raw.BlockNumber,
			TxHash:      iter.Event.Raw.TxHash,
			Index:       iter.Event.Raw.Index,
			MessageHash: iter.Event.MessageHash,
		}
	}
	return event
}
