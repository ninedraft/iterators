package iterators

import "context"

// Zip combines two sequences into sequence of pairs.
// If any of iterators is drained, then resulting iterator will return false.
func Zip[A, B any](aa Iterator[A], bb Iterator[B]) Pairs[A, B] {
	return Fn[Pair[A, B]](func(ctx context.Context) (Pair[A, B], bool) {
		var a, aOk = aa.Next(ctx)
		if !aOk {
			return none[Pair[A, B]]()
		}
		var b, bOk = bb.Next(ctx)
		if !bOk {
			return none[Pair[A, B]]()
		}
		return Pair[A, B]{A: a, B: b}, true
	})
}
