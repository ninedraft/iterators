package iterators

import (
	"context"
)

// Zip combines two sequences into sequence of pairs.
// If any of iterators is drained, then resulting iterator will return false.
func Zip[A, B any](aa Iterator[A], bb Iterator[B]) Pairs[A, B] {
	var stopped = false
	return Fn[Pair[A, B]](func(ctx context.Context) (Pair[A, B], bool) {
		if stopped {
			return none[Pair[A, B]]()
		}
		var a, aOk = aa.Next(ctx)
		if !aOk {
			stopped = true
			return none[Pair[A, B]]()
		}
		var b, bOk = bb.Next(ctx)
		if !bOk {
			stopped = true
			return none[Pair[A, B]]()
		}
		return Pair[A, B]{A: a, B: b}, true
	})
}
