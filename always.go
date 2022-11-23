package iterators

import "context"

// Empty is a always drained iterator.
func Empty[E any]() Iterator[E] {
	return Fn[E](func(context.Context) (E, bool) {
		return empty[E](), false
	})
}

// Always return provided value.
// Returns false if context is canceled.
func Always[E any](value E) Iterator[E] {
	return Fn[E](func(ctx context.Context) (E, bool) {
		return value, ctx.Err() == nil
	})
}
