package events

import (
	"context"
	"math/big"

	"github.com/scroll-tech/go-ethereum/common"

	"github.com/scroll-tech/chain-monitor/internal/logic/contracts/abi/il1erc721gateway"
	"github.com/scroll-tech/chain-monitor/internal/logic/contracts/abi/il2erc721gateway"
	"github.com/scroll-tech/chain-monitor/internal/types"
)

type ERC721GatewayEventUnmarshaler struct {
	Layer   types.LayerType
	Type    types.EventType
	Number  uint64
	TxHash  common.Hash
	TokenId *big.Int
	Index   uint

	MessageHash  common.Hash
	TokenAddress common.Address
}

func NewERC721GatewayEventUnmarshaler() *ERC721GatewayEventUnmarshaler {
	return &ERC721GatewayEventUnmarshaler{}
}

func (e *ERC721GatewayEventUnmarshaler) Unmarshal(context context.Context, layerType types.LayerType, iterators []types.WrapIterator) []EventUnmarshaler {
	var events []EventUnmarshaler
	for _, it := range iterators {
		for it.Iter.Next() {
			events = append(events, e.erc721(layerType, it.Iter, it.EventType))
		}
	}
	return events
}

func (e *ERC721GatewayEventUnmarshaler) erc721(layerType types.LayerType, it types.Iterator, eventType types.EventType) EventUnmarshaler {
	var event EventUnmarshaler
	switch eventType {
	case types.L1DepositERC721:
		iter := it.(*il1erc721gateway.Il1erc721gatewayDepositERC721Iterator)
		event = &ERC721GatewayEventUnmarshaler{
			Layer:        layerType,
			Type:         eventType,
			Number:       iter.Event.Raw.BlockNumber,
			TxHash:       iter.Event.Raw.TxHash,
			TokenId:      iter.Event.TokenId,
			Index:        iter.Event.Raw.Index,
			TokenAddress: iter.Event.Raw.Address,
		}
	case types.L1BatchDepositERC721:

	case types.L1FinalizeWithdrawERC721:
		iter := it.(*il1erc721gateway.Il1erc721gatewayFinalizeWithdrawERC721Iterator)
		event = &ERC721GatewayEventUnmarshaler{
			Layer:        layerType,
			Type:         eventType,
			Number:       iter.Event.Raw.BlockNumber,
			TxHash:       iter.Event.Raw.TxHash,
			TokenId:      iter.Event.TokenId,
			Index:        iter.Event.Raw.Index,
			TokenAddress: iter.Event.Raw.Address,
		}
	case types.L1FinalizeBatchWithdrawERC721:

	case types.L1RefundERC721:
		iter := it.(*il1erc721gateway.Il1erc721gatewayRefundERC721Iterator)
		event = &ERC721GatewayEventUnmarshaler{
			Layer:        layerType,
			Type:         eventType,
			Number:       iter.Event.Raw.BlockNumber,
			TxHash:       iter.Event.Raw.TxHash,
			TokenId:      iter.Event.TokenId,
			Index:        iter.Event.Raw.Index,
			TokenAddress: iter.Event.Raw.Address,
		}
	case types.L1BatchRefundERC721:

	case types.L2WithdrawERC721:
		iter := it.(*il2erc721gateway.Il2erc721gatewayWithdrawERC721Iterator)
		event = &ERC721GatewayEventUnmarshaler{
			Layer:        layerType,
			Type:         eventType,
			Number:       iter.Event.Raw.BlockNumber,
			TxHash:       iter.Event.Raw.TxHash,
			TokenId:      iter.Event.TokenId,
			Index:        iter.Event.Raw.Index,
			TokenAddress: iter.Event.Raw.Address,
		}
	case types.L2BatchWithdrawERC721:

	case types.L2FinalizeDepositERC721:
		iter := it.(*il2erc721gateway.Il2erc721gatewayFinalizeDepositERC721Iterator)
		event = &ERC721GatewayEventUnmarshaler{
			Layer:        layerType,
			Type:         eventType,
			Number:       iter.Event.Raw.BlockNumber,
			TxHash:       iter.Event.Raw.TxHash,
			TokenId:      iter.Event.TokenId,
			Index:        iter.Event.Raw.Index,
			TokenAddress: iter.Event.Raw.Address,
		}
	case types.L2FinalizeBatchDepositERC721:
	}
	return event
}
