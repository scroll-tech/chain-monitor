package events

import (
	"context"
	"math/big"

	"github.com/scroll-tech/chain-monitor/internal/logic/contracts"
	"github.com/scroll-tech/chain-monitor/internal/logic/contracts/abi/il1erc20gateway"
	"github.com/scroll-tech/chain-monitor/internal/logic/contracts/abi/iscrollerc20"
	"github.com/scroll-tech/chain-monitor/internal/types"
)

type ERC20GatewayEventUnmarshaler struct {
	Transfer bool
	Layer    types.LayerType
	Type     types.EventType
	Number   uint64
	TxHash   string
	MsgHash  string
	Amount   *big.Int

	erc20mapping map[types.EventType]types.ERC20
}

func NewERC20GatewayEventUnmarshaler() *ERC20GatewayEventUnmarshaler {
	e := &ERC20GatewayEventUnmarshaler{
		erc20mapping: make(map[types.EventType]types.ERC20),
	}

	e.erc20mapping[types.L1DepositWETH] = types.WETH
	e.erc20mapping[types.L1FinalizeWithdrawWETH] = types.WETH
	e.erc20mapping[types.L1RefundWETH] = types.WETH

	return e
}

func (e *ERC20GatewayEventUnmarshaler) Unmarshal(context context.Context, layerType types.LayerType, iterators []contracts.WrapIterator) []EventUnmarshaler {
	var events []EventUnmarshaler
	for _, it := range iterators {
		for it.Iter.Next() {
			if it.Transfer {
				transferEvent := e.transfer(layerType, it.Iter)
				events = append(events, transferEvent)
				continue
			}

			erc20Type := e.erc20mapping[it.EventType]
			var erc20events []EventUnmarshaler
			switch erc20Type {
			case types.WETH:
				erc20events = e.weth(layerType, it.Iter, it.EventType)
			case types.DAI:
			case types.LIDO:
			case types.CustomERC20:
			case types.StandardERC20:
			case types.USDCERC20:
			}
			events = append(events, erc20events...)
		}

	}
	return events
}

func (e *ERC20GatewayEventUnmarshaler) transfer(layerType types.LayerType, it contracts.Iterator) EventUnmarshaler {
	transferIter := it.(*iscrollerc20.Iscrollerc20TransferIterator)
	event := &ERC20GatewayEventUnmarshaler{
		Transfer: true,
		Layer:    layerType,
		Number:   transferIter.Event.Raw.BlockNumber,
		TxHash:   transferIter.Event.Raw.TxHash.Hex(),
		MsgHash:  "",
		Amount:   transferIter.Event.Value,
	}
	return event
}

func (e *ERC20GatewayEventUnmarshaler) weth(layerType types.LayerType, it contracts.Iterator, eventType types.EventType) []EventUnmarshaler {
	var events []EventUnmarshaler
	switch eventType {
	case types.L1DepositWETH:
		erc20DepositIter := it.(*il1erc20gateway.Il1erc20gatewayDepositERC20Iterator)
		event := &ERC20GatewayEventUnmarshaler{
			Transfer: false,
			Layer:    layerType,
			Type:     types.L1DepositWETH,
			Number:   erc20DepositIter.Event.Raw.BlockNumber,
			TxHash:   erc20DepositIter.Event.Raw.TxHash.Hex(),
			MsgHash:  "",
			Amount:   erc20DepositIter.Event.Amount,
		}
		events = append(events, event)
	case types.L1FinalizeWithdrawWETH:
		erc20FinalizeWithdrawIter := it.(*il1erc20gateway.Il1erc20gatewayFinalizeWithdrawERC20Iterator)
		event := &ERC20GatewayEventUnmarshaler{
			Transfer: false,
			Layer:    layerType,
			Type:     types.L1FinalizeWithdrawWETH,
			Number:   erc20FinalizeWithdrawIter.Event.Raw.BlockNumber,
			TxHash:   erc20FinalizeWithdrawIter.Event.Raw.TxHash.Hex(),
			MsgHash:  "",
			Amount:   erc20FinalizeWithdrawIter.Event.Amount,
		}
		events = append(events, event)
	case types.L1RefundWETH:
		erc20RefundIter := it.(*il1erc20gateway.Il1erc20gatewayRefundERC20Iterator)
		event := &ERC20GatewayEventUnmarshaler{
			Transfer: false,
			Layer:    layerType,
			Type:     types.L1RefundWETH,
			Number:   erc20RefundIter.Event.Raw.BlockNumber,
			TxHash:   erc20RefundIter.Event.Raw.TxHash.Hex(),
			MsgHash:  "",
			Amount:   erc20RefundIter.Event.Amount,
		}
		events = append(events, event)
	}
	return events
}
