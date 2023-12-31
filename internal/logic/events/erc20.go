package events

import (
	"context"
	"math/big"

	"github.com/scroll-tech/go-ethereum/common"

	"github.com/scroll-tech/chain-monitor/internal/logic/contracts/abi/il1erc20gateway"
	"github.com/scroll-tech/chain-monitor/internal/logic/contracts/abi/il2erc20gateway"
	"github.com/scroll-tech/chain-monitor/internal/types"
)

// ERC20GatewayEventUnmarshaler is a struct that helps unmarshal events from the ERC20 Gateway.
type ERC20GatewayEventUnmarshaler struct {
	Layer        types.LayerType
	Type         types.EventType
	Number       uint64
	TxHash       common.Hash
	Amount       *big.Int
	Index        uint
	MessageHash  common.Hash
	TokenAddress common.Address
}

// Unmarshal takes a context, layer type, and a list of iterators and unmarshals each iterator
// into an EventUnmarshaler, returning a list of these unmarshalled events.
func (e *ERC20GatewayEventUnmarshaler) Unmarshal(context context.Context, layerType types.LayerType, iterators []types.WrapIterator) []EventUnmarshaler {
	var events []EventUnmarshaler
	for _, it := range iterators {
		for it.Iter.Next() {
			events = append(events, e.erc20(layerType, it.Iter, it.EventType))
		}
	}
	return events
}

func (e *ERC20GatewayEventUnmarshaler) erc20(layerType types.LayerType, it types.Iterator, eventType types.EventType) EventUnmarshaler {
	var event EventUnmarshaler
	switch eventType {
	case types.L1DepositERC20:
		iter := it.(*il1erc20gateway.Il1erc20gatewayDepositERC20Iterator)
		event = &ERC20GatewayEventUnmarshaler{
			Layer:        layerType,
			Type:         eventType,
			Number:       iter.Event.Raw.BlockNumber,
			TxHash:       iter.Event.Raw.TxHash,
			Amount:       iter.Event.Amount,
			Index:        iter.Event.Raw.Index,
			TokenAddress: iter.Event.L1Token,
		}
	case types.L1FinalizeWithdrawERC20:
		iter := it.(*il1erc20gateway.Il1erc20gatewayFinalizeWithdrawERC20Iterator)
		event = &ERC20GatewayEventUnmarshaler{
			Layer:        layerType,
			Type:         eventType,
			Number:       iter.Event.Raw.BlockNumber,
			TxHash:       iter.Event.Raw.TxHash,
			Amount:       iter.Event.Amount,
			Index:        iter.Event.Raw.Index,
			TokenAddress: iter.Event.L1Token,
		}
	case types.L1RefundERC20:
		iter := it.(*il1erc20gateway.Il1erc20gatewayRefundERC20Iterator)
		event = &ERC20GatewayEventUnmarshaler{
			Layer:        layerType,
			Type:         eventType,
			Number:       iter.Event.Raw.BlockNumber,
			TxHash:       iter.Event.Raw.TxHash,
			Amount:       iter.Event.Amount,
			Index:        iter.Event.Raw.Index,
			TokenAddress: iter.Event.Token,
		}
	case types.L2WithdrawERC20:
		iter := it.(*il2erc20gateway.Il2erc20gatewayWithdrawERC20Iterator)
		event = &ERC20GatewayEventUnmarshaler{
			Layer:        layerType,
			Type:         eventType,
			Number:       iter.Event.Raw.BlockNumber,
			TxHash:       iter.Event.Raw.TxHash,
			Amount:       iter.Event.Amount,
			Index:        iter.Event.Raw.Index,
			TokenAddress: iter.Event.L2Token,
		}
	case types.L2FinalizeDepositERC20:
		iter := it.(*il2erc20gateway.Il2erc20gatewayFinalizeDepositERC20Iterator)
		event = &ERC20GatewayEventUnmarshaler{
			Layer:        layerType,
			Type:         eventType,
			Number:       iter.Event.Raw.BlockNumber,
			TxHash:       iter.Event.Raw.TxHash,
			Amount:       iter.Event.Amount,
			Index:        iter.Event.Raw.Index,
			TokenAddress: iter.Event.L2Token,
		}
	}
	return event
}
