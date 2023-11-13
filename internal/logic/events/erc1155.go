package events

import (
	"context"
	"math/big"

	"github.com/scroll-tech/go-ethereum/common"

	"github.com/scroll-tech/chain-monitor/internal/types"
)

type ERC1155GatewayEventUnmarshaler struct {
	Layer   types.LayerType
	Type    types.EventType
	Number  uint64
	TxHash  common.Hash
	TokenId *big.Int
	Index   uint

	MessageHash  common.Hash
	TokenAddress common.Address
}

func NewERC1155GatewayEventUnmarshaler() *ERC1155GatewayEventUnmarshaler {
	return &ERC1155GatewayEventUnmarshaler{}
}

func (e *ERC1155GatewayEventUnmarshaler) Unmarshal(context context.Context, layerType types.LayerType, iterators []types.WrapIterator) []EventUnmarshaler {
	var events []EventUnmarshaler
	for _, it := range iterators {
		for it.Iter.Next() {
			events = append(events, e.erc20(layerType, it.Iter, it.EventType))
		}
	}
	return events
}
