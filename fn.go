package iterators

import "context"

// Fn is a the most simple iterator - hook, which can be called multiple times until drained.
type Fn[E any] func(ctx context.Context) (_ E, ok bool)

// Next implements the Iterator interface.
// Always returns false if Fn is nil.
func (fn Fn[E]) Next(ctx context.Context) (_ E, ok bool) {
	if fn == nil {
		return empty[E](), false
	}
	return fn(ctx)
}
