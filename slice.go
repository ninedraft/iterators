package iterators

import "context"

// Slice creates an iterator from provided slice.
// Resulting iterator will emit all values in order index=0...len.
// Length of iterator reports the number of uncunsomed value.
func Slice[E any, S ~[]E](slice S) Sized[E] {
	return &iterSlice[E]{
		slice: slice,
	}
}

type iterSlice[E any] struct {
	slice []E
}

func (iter *iterSlice[E]) Len() int {
	return len(iter.slice)
}

func (iter *iterSlice[E]) Next(ctx context.Context) (E, bool) {
	if len(iter.slice) == 0 {
		return none[E]()
	}
	var value = iter.slice[0]
	iter.slice = iter.slice[1:]
	return value, true
}
