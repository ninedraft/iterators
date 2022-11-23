package iterators

import "context"

// StopOnCancel returns false if context is canceled.
func StopOnCancel[E any](iter Iterator[E]) Iterator[E] {
	return Fn[E](func(ctx context.Context) (E, bool) {
		if ctx.Err() != nil {
			return none[E]()
		}
		return iter.Next(ctx)
	})
}
