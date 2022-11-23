package iterators

// Pair describes a sequence of pairs.
type Pairs[A, B any] Iterator[Pair[A, B]]

// Pair describes a sequence of pairs with known size.
type PairsSized[A, B any] Sized[Pair[A, B]]

// Pair describes two connected values.
type Pair[A, B any] struct {
	A A
	B B
}

// Unpack returns both values from pair.
// It's useful then you need a (T, error) function.
func (pair *Pair[A, B]) Unpack() (A, B) {
	return pair.A, pair.B
}
