package events

import (
	"context"
	"math/big"

	"github.com/scroll-tech/go-ethereum/common"

	"github.com/scroll-tech/chain-monitor/internal/logic/contracts/abi/il1erc1155gateway"
	"github.com/scroll-tech/chain-monitor/internal/logic/contracts/abi/il2erc1155gateway"
	"github.com/scroll-tech/chain-monitor/internal/types"
)

// ERC1155GatewayEventUnmarshaler is a struct that helps unmarshal events from the ERC1155 Gateway.
type ERC1155GatewayEventUnmarshaler struct {
	Layer        types.LayerType
	Type         types.EventType
	Number       uint64
	TxHash       common.Hash
	TokenIds     []*big.Int
	Amounts      []*big.Int
	Index        uint
	MessageHash  common.Hash
	TokenAddress common.Address
}

// Unmarshal takes a context, layer type, and a list of iterators and unmarshals each iterator
// into an EventUnmarshaler, returning a list of these unmarshaled events.
func (e *ERC1155GatewayEventUnmarshaler) Unmarshal(context context.Context, layerType types.LayerType, iterators []types.WrapIterator) []EventUnmarshaler {
	var events []EventUnmarshaler
	for _, it := range iterators {
		for it.Iter.Next() {
			events = append(events, e.erc1155(layerType, it.Iter, it.EventType))
		}
	}
	return events
}

func (e *ERC1155GatewayEventUnmarshaler) erc1155(layerType types.LayerType, it types.Iterator, eventType types.EventType) EventUnmarshaler {
	var event EventUnmarshaler
	switch eventType {
	case types.L1DepositERC1155:
		iter := it.(*il1erc1155gateway.Il1erc1155gatewayDepositERC1155Iterator)
		event = &ERC1155GatewayEventUnmarshaler{
			Layer:        layerType,
			Type:         eventType,
			Number:       iter.Event.Raw.BlockNumber,
			TxHash:       iter.Event.Raw.TxHash,
			TokenIds:     []*big.Int{iter.Event.TokenId},
			Amounts:      []*big.Int{iter.Event.Amount},
			Index:        iter.Event.Raw.Index,
			TokenAddress: iter.Event.L1Token,
		}
	case types.L1BatchDepositERC1155:
		iter := it.(*il1erc1155gateway.Il1erc1155gatewayBatchDepositERC1155Iterator)
		event = &ERC1155GatewayEventUnmarshaler{
			Layer:        layerType,
			Type:         eventType,
			Number:       iter.Event.Raw.BlockNumber,
			TxHash:       iter.Event.Raw.TxHash,
			TokenIds:     iter.Event.TokenIds,
			Amounts:      iter.Event.Amounts,
			Index:        iter.Event.Raw.Index,
			TokenAddress: iter.Event.L1Token,
		}
	case types.L1FinalizeWithdrawERC1155:
		iter := it.(*il1erc1155gateway.Il1erc1155gatewayFinalizeWithdrawERC1155Iterator)
		event = &ERC1155GatewayEventUnmarshaler{
			Layer:        layerType,
			Type:         eventType,
			Number:       iter.Event.Raw.BlockNumber,
			TxHash:       iter.Event.Raw.TxHash,
			TokenIds:     []*big.Int{iter.Event.TokenId},
			Amounts:      []*big.Int{iter.Event.Amount},
			Index:        iter.Event.Raw.Index,
			TokenAddress: iter.Event.L1Token,
		}
	case types.L1FinalizeBatchWithdrawERC1155:
		iter := it.(*il1erc1155gateway.Il1erc1155gatewayFinalizeBatchWithdrawERC1155Iterator)
		event = &ERC1155GatewayEventUnmarshaler{
			Layer:        layerType,
			Type:         eventType,
			Number:       iter.Event.Raw.BlockNumber,
			TxHash:       iter.Event.Raw.TxHash,
			TokenIds:     iter.Event.TokenIds,
			Amounts:      iter.Event.Amounts,
			Index:        iter.Event.Raw.Index,
			TokenAddress: iter.Event.L1Token,
		}
	case types.L1RefundERC1155:
		iter := it.(*il1erc1155gateway.Il1erc1155gatewayRefundERC1155Iterator)
		event = &ERC1155GatewayEventUnmarshaler{
			Layer:        layerType,
			Type:         eventType,
			Number:       iter.Event.Raw.BlockNumber,
			TxHash:       iter.Event.Raw.TxHash,
			TokenIds:     []*big.Int{iter.Event.TokenId},
			Amounts:      []*big.Int{iter.Event.Amount},
			Index:        iter.Event.Raw.Index,
			TokenAddress: iter.Event.Token,
		}
	case types.L1BatchRefundERC1155:
		iter := it.(*il1erc1155gateway.Il1erc1155gatewayBatchRefundERC1155Iterator)
		event = &ERC1155GatewayEventUnmarshaler{
			Layer:        layerType,
			Type:         eventType,
			Number:       iter.Event.Raw.BlockNumber,
			TxHash:       iter.Event.Raw.TxHash,
			TokenIds:     iter.Event.TokenIds,
			Amounts:      iter.Event.Amounts,
			Index:        iter.Event.Raw.Index,
			TokenAddress: iter.Event.Token,
		}
	case types.L2WithdrawERC1155:
		iter := it.(*il2erc1155gateway.Il2erc1155gatewayWithdrawERC1155Iterator)
		event = &ERC1155GatewayEventUnmarshaler{
			Layer:        layerType,
			Type:         eventType,
			Number:       iter.Event.Raw.BlockNumber,
			TxHash:       iter.Event.Raw.TxHash,
			TokenIds:     []*big.Int{iter.Event.TokenId},
			Amounts:      []*big.Int{iter.Event.Amount},
			Index:        iter.Event.Raw.Index,
			TokenAddress: iter.Event.L2Token,
		}
	case types.L2BatchWithdrawERC1155:
		iter := it.(*il2erc1155gateway.Il2erc1155gatewayBatchWithdrawERC1155Iterator)
		event = &ERC1155GatewayEventUnmarshaler{
			Layer:        layerType,
			Type:         eventType,
			Number:       iter.Event.Raw.BlockNumber,
			TxHash:       iter.Event.Raw.TxHash,
			TokenIds:     iter.Event.TokenIds,
			Amounts:      iter.Event.Amounts,
			Index:        iter.Event.Raw.Index,
			TokenAddress: iter.Event.L2Token,
		}
	case types.L2FinalizeDepositERC1155:
		iter := it.(*il2erc1155gateway.Il2erc1155gatewayFinalizeDepositERC1155Iterator)
		event = &ERC1155GatewayEventUnmarshaler{
			Layer:        layerType,
			Type:         eventType,
			Number:       iter.Event.Raw.BlockNumber,
			TxHash:       iter.Event.Raw.TxHash,
			TokenIds:     []*big.Int{iter.Event.TokenId},
			Amounts:      []*big.Int{iter.Event.Amount},
			Index:        iter.Event.Raw.Index,
			TokenAddress: iter.Event.L2Token,
		}
	case types.L2FinalizeBatchDepositERC1155:
		iter := it.(*il2erc1155gateway.Il2erc1155gatewayFinalizeBatchDepositERC1155Iterator)
		event = &ERC1155GatewayEventUnmarshaler{
			Layer:        layerType,
			Type:         eventType,
			Number:       iter.Event.Raw.BlockNumber,
			TxHash:       iter.Event.Raw.TxHash,
			TokenIds:     iter.Event.TokenIds,
			Amounts:      iter.Event.Amounts,
			Index:        iter.Event.Raw.Index,
			TokenAddress: iter.Event.L2Token,
		}
	}
	return event
}
