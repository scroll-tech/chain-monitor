package events

import (
	"context"
	"math/big"

	"github.com/scroll-tech/go-ethereum/common"

	"github.com/scroll-tech/chain-monitor/internal/logic/contracts"
	"github.com/scroll-tech/chain-monitor/internal/logic/contracts/abi/il1erc20gateway"
	"github.com/scroll-tech/chain-monitor/internal/logic/contracts/abi/il2erc20gateway"
	"github.com/scroll-tech/chain-monitor/internal/logic/contracts/abi/iscrollerc20"
	"github.com/scroll-tech/chain-monitor/internal/types"
)

type ERC20GatewayEventUnmarshaler struct {
	Transfer bool
	Layer    types.LayerType
	Type     types.EventType
	Number   uint64
	TxHash   common.Hash
	Amount   *big.Int

	MsgHash common.Hash
}

func NewERC20GatewayEventUnmarshaler() *ERC20GatewayEventUnmarshaler {
	return &ERC20GatewayEventUnmarshaler{}
}

func (e *ERC20GatewayEventUnmarshaler) Unmarshal(context context.Context, layerType types.LayerType, iterators []contracts.WrapIterator) []EventUnmarshaler {
	var events []EventUnmarshaler
	for _, it := range iterators {
		for it.Iter.Next() {
			if it.Transfer {
				events = append(events, e.transfer(layerType, it.Iter))
			} else {
				events = append(events, e.erc20(layerType, it.Iter, it.EventType))
			}
		}
	}
	return events
}

func (e *ERC20GatewayEventUnmarshaler) transfer(layerType types.LayerType, it contracts.Iterator) EventUnmarshaler {
	iter := it.(*iscrollerc20.Iscrollerc20TransferIterator)
	event := &ERC20GatewayEventUnmarshaler{
		Transfer: true,
		Layer:    layerType,
		Number:   iter.Event.Raw.BlockNumber,
		TxHash:   iter.Event.Raw.TxHash,
		Amount:   iter.Event.Value,
	}
	return event
}

func (e *ERC20GatewayEventUnmarshaler) erc20(layerType types.LayerType, it contracts.Iterator, eventType types.EventType) EventUnmarshaler {
	var event EventUnmarshaler
	switch eventType {
	case types.L1DepositERC20:
		iter := it.(*il1erc20gateway.Il1erc20gatewayDepositERC20Iterator)
		event = &ERC20GatewayEventUnmarshaler{
			Transfer: false,
			Layer:    layerType,
			Type:     eventType,
			Number:   iter.Event.Raw.BlockNumber,
			TxHash:   iter.Event.Raw.TxHash,
			Amount:   iter.Event.Amount,
		}
	case types.L1FinalizeWithdrawERC20:
		iter := it.(*il1erc20gateway.Il1erc20gatewayFinalizeWithdrawERC20Iterator)
		event = &ERC20GatewayEventUnmarshaler{
			Transfer: false,
			Layer:    layerType,
			Type:     eventType,
			Number:   iter.Event.Raw.BlockNumber,
			TxHash:   iter.Event.Raw.TxHash,
			Amount:   iter.Event.Amount,
		}
	case types.L1RefundERC20:
		iter := it.(*il1erc20gateway.Il1erc20gatewayRefundERC20Iterator)
		event = &ERC20GatewayEventUnmarshaler{
			Transfer: false,
			Layer:    layerType,
			Type:     eventType,
			Number:   iter.Event.Raw.BlockNumber,
			TxHash:   iter.Event.Raw.TxHash,
			Amount:   iter.Event.Amount,
		}
	case types.L2WithdrawERC20:
		iter := it.(*il2erc20gateway.Il2erc20gatewayWithdrawERC20Iterator)
		event = &ERC20GatewayEventUnmarshaler{
			Transfer: false,
			Layer:    layerType,
			Type:     eventType,
			Number:   iter.Event.Raw.BlockNumber,
			TxHash:   iter.Event.Raw.TxHash,
			Amount:   iter.Event.Amount,
		}
	case types.L2FinalizeDepositERC20:
		iter := it.(*il2erc20gateway.Il2erc20gatewayFinalizeDepositERC20Iterator)
		event = &ERC20GatewayEventUnmarshaler{
			Transfer: false,
			Layer:    layerType,
			Type:     eventType,
			Number:   iter.Event.Raw.BlockNumber,
			TxHash:   iter.Event.Raw.TxHash,
			Amount:   iter.Event.Amount,
		}
	}
	return event
}
