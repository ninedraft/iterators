package iterators

import (
	"context"
	"sync/atomic"
)

// Closer makes provided iter closable.
// If provided iter is a closer already, then it returned as is.
// If it's not, then returned iterator will always return false after
// method close is called.
func Closer[E any](iter Iterator[E]) IteratorCloser[E] {
	if it, ok := iter.(IteratorCloser[E]); ok {
		return it
	}
	return &iterCloser[E]{
		iter: iter,
	}
}

type iterCloser[E any] struct {
	isClosed atomic.Bool
	iter     Iterator[E]
}

func (iter *iterCloser[E]) Close() error {
	iter.isClosed.Store(true)
	return nil
}

func (iter *iterCloser[E]) Next(ctx context.Context) (E, bool) {
	if iter.isClosed.Load() {
		return none[E]()
	}
	return iter.iter.Next(ctx)
}
