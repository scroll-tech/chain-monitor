package events

import (
	"context"

	"github.com/scroll-tech/go-ethereum/log"

	"github.com/scroll-tech/chain-monitor/internal/types"
)

// EventGather is a struct that holds a map of transaction event categories to event unmarshalers.
type EventGather struct {
	gathers map[types.EventCategory]EventUnmarshaler
}

// NewEventGather creates a new instance of EventGather and initializes it with different types of
// event unmarshalers for each transaction event category.
func NewEventGather() *EventGather {
	g := &EventGather{
		gathers: make(map[types.EventCategory]EventUnmarshaler),
	}

	g.gathers[types.ERC20EventCategory] = &ERC20GatewayEventUnmarshaler{}
	g.gathers[types.ERC721EventCategory] = &ERC721GatewayEventUnmarshaler{}
	g.gathers[types.ERC1155EventCategory] = &ERC1155GatewayEventUnmarshaler{}
	g.gathers[types.MessengerEventCategory] = &MessengerEventUnmarshaler{}

	return g
}

// Dispatch takes a context, layer type, event category, and a list of iterators, and dispatches
// the events to the appropriate unmarshaler based on the event category. It returns a list of the
// unmarshalled events.
func (e *EventGather) Dispatch(ctx context.Context, layer types.LayerType, eventCategory types.EventCategory, iterators []types.WrapIterator) []EventUnmarshaler {
	gather, exist := e.gathers[eventCategory]
	if !exist {
		log.Error("EventGather.Dispatch failed: there no event unmarshaler exist", "layer", layer.String(), "event category", eventCategory.String())
		return nil
	}

	return gather.Unmarshal(ctx, layer, iterators)
}
