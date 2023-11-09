package events

import (
	"context"
	"fmt"

	"github.com/scroll-tech/go-ethereum/log"

	"github.com/scroll-tech/chain-monitor/internal/types"
)

type EventGather struct {
	gathers map[types.TxEventCategory]EventUnmarshaler
}

func NewEventGather() *EventGather {
	g := &EventGather{
		gathers: make(map[types.TxEventCategory]EventUnmarshaler),
	}

	g.gathers[types.ERC20EventCategory] = NewERC20GatewayEventUnmarshaler()

	return g
}

func (e *EventGather) Dispatch(ctx context.Context, layer types.LayerType, eventCategory types.TxEventCategory, iterators []types.WrapIterator) []EventUnmarshaler {
	gather, exist := e.gathers[eventCategory]
	if !exist {
		err := fmt.Errorf("there no event unmarshaler exist, layer type:%d, event category:%d", layer, eventCategory)
		log.Error("EventGather.Dispatch failed", "error", err)
		return nil
	}

	return gather.Unmarshal(ctx, layer, iterators)
}
