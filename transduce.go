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

// TransduceArg maps values of type F to values of T using provided function fn.
// If input iter is drained, then output iterator will return false.
// Results of fn are attached to corresponding arguments of fn.
func TransduceArg[F, T any](iter Iterator[F], fn func(F) T) Pairs[F, T] {
	return Fn[Pair[F, T]](func(ctx context.Context) (Pair[F, T], bool) {
		var value, ok = iter.Next(ctx)
		if !ok {
			return none[Pair[F, T]]()
		}
		return Pair[F, T]{
			A: value,
			B: fn(value),
		}, true
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
			return value, true
		}
	})
}

func Compact[E comparable](iter Iterator[E]) Iterator[E] {
	var prev E
	var first = true
	return Fn[E](func(ctx context.Context) (E, bool) {
		for {
			var value, ok = iter.Next(ctx)
			if !ok {
				return none[E]()
			}
			if !first && prev == value {
				continue
			}
			prev = value
			return value, true
		}
	})
}

func CompactFn[E any](iter Iterator[E], fn func(a, B E) bool) Iterator[E] {
	var prev E
	var first = true
	return Fn[E](func(ctx context.Context) (E, bool) {
		for {
			var value, ok = iter.Next(ctx)
			if !ok {
				return none[E]()
			}
			if !first && fn(prev, value) {
				continue
			}
			prev = value
			return value, true
		}
	})
}
