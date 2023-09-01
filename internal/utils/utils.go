package utils

import (
	"context"
	"time"

	"github.com/modern-go/reflect2"
)

// LoopWithContext Run the f func with context periodically.
func LoopWithContext(ctx context.Context, period time.Duration, f func(ctx context.Context)) {
	tick := time.NewTicker(period)
	defer tick.Stop()
	for true {
		select {
		case <-ctx.Done():
			return
		case <-tick.C:
			f(ctx)
		}
	}
}

// IsNil Check if the interface is empty.
func IsNil(i interface{}) bool {
	return i == nil || reflect2.IsNil(i)
}
