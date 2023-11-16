package events

import (
	"context"

	"github.com/scroll-tech/chain-monitor/internal/types"
)

// EventUnmarshaler is an interface that defines the Unmarshal method.
type EventUnmarshaler interface {
	Unmarshal(context.Context, types.LayerType, []types.WrapIterator) []EventUnmarshaler
}
