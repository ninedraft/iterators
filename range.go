package iterators

import (
	"context"
	"math"

	"golang.org/x/exp/constraints"
)

// Number describes a generic integer or float number.
type Number interface {
	constraints.Signed | constraints.Unsigned | constraints.Float
}

// Range emits a bounded sequence of numbers [from, to).
// If from<to, then it's equal to RangeStep(from, to, 1).
// If from>to, then returns an empty iterator.
func Range[N Number](from, to N) Iterator[N] {
	if from > to {
		return Empty[N]()
	}
	return RangeStep(from, to, 1)
}

func RangeStep[N Number](from, to, step N) Iterator[N] {
	switch {
	case isNaN(from), isNaN(to), isNaN(step):
		return Empty[N]()
	case step == 0:
		return Empty[N]()
	case from > to && step > 0:
		return Empty[N]()
	case from < to && step < 0:
		return Empty[N]()
	}

	var i = from
	return Fn[N](func(context.Context) (N, bool) {
		if i > to {
			return 0, false
		}
		var x = i
		i += step
		return x, true
	})
}

func isNaN[N Number](x N) bool {
	return math.IsNaN(float64(x))
}
