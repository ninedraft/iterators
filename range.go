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

// RangeStep emits a bounded sequence of numbers [from, to).
// Returns non-empty iterator if from<to && step>0 or from>to && step<0.
// Returns an empty iterator otherwise.
// Returns an empty iterator if any of arguments is NaN.
func RangeStep[N Number](from, to, step N) Iterator[N] {
	var i = from
	var stop = func() bool { return i >= to }
	switch {
	case isNaN(from), isNaN(to), isNaN(step):
		return Empty[N]()
	case step == 0:
		return Empty[N]()
	case from > to && step > 0:
		return Empty[N]()
	case from > to:
		stop = func() bool { return i <= to }
	case from < to && step < 0:
		return Empty[N]()
	}

	return Fn[N](func(context.Context) (N, bool) {
		if stop() {
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
