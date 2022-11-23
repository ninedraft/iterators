package iterators

import (
	"context"
	"io"
)

// Iterator is a generic stateful of values.
type Iterator[E any] interface {
	// Next returns false if there is no more values in the sequence.
	// Iterator can ignore provided context or return false if context is canceled.
	Next(ctx context.Context) (_ E, ok bool)
}

// IteratorErr describes an iterator, which can enter an erroneous state.
// If there is an error condition, then the Err method must return a non-nil error.
type IteratorErr[E any] interface {
	Iterator[E]
	Err() error
}

// IteratorCloser describes an iterator, which can be closed.
// Multiple close method calls are a valid usage and must not panic.
// After first method close call the instance of iterator must return false from Next method.
type IteratorCloser[E any] interface {
	Iterator[E]
	io.Closer
}

// Sized describe an iterator with known size.
// If Len()==0, then Next() must return false.
// If Next() returns false, then Len() must be 0.
type Sized[E any] interface {
	Iterator[E]
	Len() int
}
