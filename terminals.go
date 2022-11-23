package iterators

import "context"

// Collect values into a dst slice.
// It will run until iter is drained.
func Collect[E any, S ~[]E](ctx context.Context, dst S, iter Iterator[E]) S {
	for {
		var value, ok = iter.Next(ctx)
		if !ok {
			return dst
		}
		dst = append(dst, value)
	}
}

// CollectN pulls a most n values from iter into dst slice.
// It will stop if iter is drained.
func CollectN[E any, S ~[]E](ctx context.Context, dst S, iter Iterator[E], n int) S {
	for i := 0; i < n; i++ {
		var value, ok = iter.Next(ctx)
		if !ok {
			return dst
		}
		dst = append(dst, value)
	}
	return dst
}

// ForEach calls fn for each value from provided iterator.
// It will run until iter is drained.
func ForEach[E any](ctx context.Context, iter Iterator[E], fn func(value E) bool) {
	for {
		var value, ok = iter.Next(ctx)
		if !ok {
			return
		}
		if !fn(value) {
			return
		}
	}
}

// Drain will consume values from iter until it is drained.
func Drain[E any](ctx context.Context, iter Iterator[E]) {
	ForEach(ctx, iter, func(E) bool {
		return true
	})
}
