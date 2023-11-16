package types

// Iterator is an interface that provides a way to access elements in a sequence or collection.
// The Next() method is used to move to the next element in the sequence.
type Iterator interface {
	Next() bool
}

// WrapIterator is a struct that wraps an Iterator with an additional EventType field.
// EventType provides context or metadata about the type of events the iterator is dealing with.
// Iter is the instance of an Iterator that this WrapIterator is enhancing.
type WrapIterator struct {
	EventType EventType
	Iter      Iterator
}
