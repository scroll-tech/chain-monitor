package types

type Iterator interface {
	Next() bool
}

type WrapIterator struct {
	EventType EventType
	Iter      Iterator
}
