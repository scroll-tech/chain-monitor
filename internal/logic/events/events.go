package events

import (
	"context"
	"fmt"

	"github.com/scroll-tech/go-ethereum/log"

	"github.com/scroll-tech/chain-monitor/internal/types"
)

// EventGather is a struct that holds a map of transaction event categories to event unmarshalers.
type EventGather struct {
	gathers map[types.TxEventCategory]EventUnmarshaler
}

// NewEventGather creates a new instance of EventGather and initializes it with different types of
// event unmarshalers for each transaction event category.
func NewEventGather() *EventGather {
	g := &EventGather{
		gathers: make(map[types.TxEventCategory]EventUnmarshaler),
	}

	g.gathers[types.ERC20EventCategory] = &ERC20GatewayEventUnmarshaler{}
	g.gathers[types.ERC721EventCategory] = &ERC721GatewayEventUnmarshaler{}
	g.gathers[types.ERC1155EventCategory] = &ERC1155GatewayEventUnmarshaler{}
	g.gathers[types.MessengerEventCategory] = &MessengerEventUnmarshaler{}

	return g
}

// Dispatch takes a context, layer type, event category, and a list of iterators, and dispatches
// the events to the appropriate unmarshaler based on the event category. It returns a list of the
// unmarshaled events.
func (e *EventGather) Dispatch(ctx context.Context, layer types.LayerType, eventCategory types.TxEventCategory, iterators []types.WrapIterator) []EventUnmarshaler {
	gather, exist := e.gathers[eventCategory]
	if !exist {
		err := fmt.Errorf("there no event unmarshaler exist, layer type:%d, event category:%d", layer, eventCategory)
		log.Error("EventGather.Dispatch failed", "error", err)
		return nil
	}

	return gather.Unmarshal(ctx, layer, iterators)
}
