package iterators

import "context"

// Chain pulls values from provided iterators.
// It will return false after last iterator is drained.
func Chain[E any](iters ...Iterator[E]) Iterator[E] {
	return Fn[E](func(ctx context.Context) (E, bool) {
		for len(iters) > 0 {
			var iter = iters[0]
			var value, ok = iter.Next(ctx)
			if !ok {
				iters[0] = nil
				iters = iters[1:]
				continue
			}
			return value, true
		}
		return none[E]()
	})
}

// Tap calls the provided fn on each result from iter.
func Tap[E any](iter Iterator[E], fn func(E, bool)) Iterator[E] {
	return Fn[E](func(ctx context.Context) (E, bool) {
		var value, ok = iter.Next(ctx)
		fn(value, ok)
		return value, ok
	})
}
