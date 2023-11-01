package contracts

import "github.com/scroll-tech/chain-monitor/internal/types"

type Iterator interface {
	Next() bool
}

type WrapIterator struct {
	Transfer  bool
	EventType types.EventType
	Iter      Iterator
}
