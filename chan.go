package iterators

import "context"

// Chan wraps a read-only channel.
type Chan[E any] <-chan E

// Len returns number of unread elements from the channel iterator.
func (ch Chan[E]) Len() int {
	return len(ch)
}

// Next implements iterator interface.
func (ch Chan[E]) Next(ctx context.Context) (_ E, ok bool) {
	select {
	case <-ctx.Done():
		return none[E]()
	case v, ok := <-ch:
		return v, ok
	}
}
