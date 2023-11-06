package events

import (
	"context"

	"github.com/scroll-tech/chain-monitor/internal/types"
)

type EventUnmarshaler interface {
	Unmarshal(context.Context, types.LayerType, []types.WrapIterator) []EventUnmarshaler
}
