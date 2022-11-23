package iterators

import "context"

// Transduce maps values of type F to values of T using provided function fn.
// If input iter is drained, then output iterator will return false.
func Transduce[F, T any](iter Iterator[F], fn func(F) T) Iterator[T] {
	return Fn[T](func(ctx context.Context) (T, bool) {
		var value, ok = iter.Next(ctx)
		if !ok {
			return none[T]()
		}
		return fn(value), true
	})
}

// TransduceFilter maps values of type F to values of T using provided function fn.
// If fn returns false, then corresponding result is skipped.
// If input iter is drained, then output iterator will return false.
func TransduceFilter[F, T any](iter Iterator[F], fn func(F) (T, bool)) Iterator[T] {
	return Fn[T](func(ctx context.Context) (T, bool) {
		for {
			var value, ok = iter.Next(ctx)
			if !ok {
				return none[T]()
			}
			var result, resultOk = fn(value)
			if !resultOk {
				continue
			}
			return result, true
		}
	})
}

// Filter pulls values from iter and filters them using provided fn.
// Output iterator emits only `fn(value)==true` values.
func Filter[E any](iter Iterator[E], fn func(E) bool) Iterator[E] {
	return Fn[E](func(ctx context.Context) (E, bool) {
		for {
			var value, ok = iter.Next(ctx)
			if !ok {
				return none[E]()
			}
			if !fn(value) {
				continue
			}
			return value
		}
	})
}
