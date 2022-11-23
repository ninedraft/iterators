package iterators

import "context"

// Enumerate assignes a sequential increasing number starting from 0 to values from provided iterator.
func Enumerate[E any](iter Iterator[E]) Pairs[int, E] {
	var i = 0
	return Fn[Pair[int, E]](func(ctx context.Context) (Pair[int, E], bool) {
		var value, ok = iter.Next(ctx)
		if !ok {
			return none[Pair[int, E]]()
		}
		var result = Pair[int, E]{
			A: i,
			B: value,
		}
		i++
		return result, true
	})
}
