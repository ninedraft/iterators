package iterators_test

import (
	"context"
	"math"
	"testing"

	"github.com/ninedraft/iterators"
	"golang.org/x/exp/slices"
)

func TestRange(test *testing.T) {
	test.Parallel()

	var ctx = context.Background()
	var iter = iterators.Range(0, 10)
	var got = iterators.Collect(ctx, []int{}, iter)

	var expected = []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	if !slices.Equal(got, expected) {
		test.Errorf("got:      %v", got)
		test.Errorf("expected: %v", expected)
	}
}

func TestRange_Equal(test *testing.T) {
	test.Parallel()

	var ctx = context.Background()
	var iter = iterators.Range(0, 0)
	var got = iterators.Collect(ctx, []int{}, iter)

	var expected = []int{}
	if !slices.Equal(got, expected) {
		test.Errorf("got:      %v", got)
		test.Errorf("expected: %v", expected)
	}
}

func TestRange_Inverse(test *testing.T) {
	test.Parallel()

	var ctx = context.Background()
	var iter = iterators.Range(10, 0)
	var got = iterators.Collect(ctx, []int{}, iter)

	var expected = []int{}
	if !slices.Equal(got, expected) {
		test.Errorf("got:      %v", got)
		test.Errorf("expected: %v", expected)
	}
}

func TestRangeStep(test *testing.T) {
	test.Parallel()

	var ctx = context.Background()
	var iter = iterators.RangeStep(0, 10, 2)
	var got = iterators.Collect(ctx, []int{}, iter)

	var expected = []int{0, 2, 4, 6, 8}
	if !slices.Equal(got, expected) {
		test.Errorf("got:      %v", got)
		test.Errorf("expected: %v", expected)
	}
}

func TestRangeStep_NegativeStep(test *testing.T) {
	test.Parallel()

	var ctx = context.Background()
	var iter = iterators.RangeStep(0, 10, -2)
	var got = iterators.Collect(ctx, []int{}, iter)

	var expected = []int{}
	if !slices.Equal(got, expected) {
		test.Errorf("got:      %v", got)
		test.Errorf("expected: %v", expected)
	}
}

func TestRangeStep_Inverse(test *testing.T) {
	test.Parallel()

	var ctx = context.Background()
	var iter = iterators.RangeStep(10, 2, -2)
	var got = iterators.Collect(ctx, []int{}, iter)

	var expected = []int{10, 8, 6, 4}
	if !slices.Equal(got, expected) {
		test.Errorf("got:      %v", got)
		test.Errorf("expected: %v", expected)
	}
}

func TestRangeStep_InversePositiveStep(test *testing.T) {
	test.Parallel()

	var ctx = context.Background()
	var iter = iterators.RangeStep(10, 2, 2)
	var got = iterators.Collect(ctx, []int{}, iter)

	var expected = []int{}
	if !slices.Equal(got, expected) {
		test.Errorf("got:      %v", got)
		test.Errorf("expected: %v", expected)
	}
}

func TestRangeStep_ZeroStep(test *testing.T) {
	test.Parallel()

	var ctx = context.Background()
	var iter = iterators.RangeStep(1, 10, 0)
	var got = iterators.Collect(ctx, []int{}, iter)

	var expected = []int{}
	if !slices.Equal(got, expected) {
		test.Errorf("got:      %v", got)
		test.Errorf("expected: %v", expected)
	}
}

func TestRangeStep_Equal(test *testing.T) {
	test.Parallel()

	var ctx = context.Background()
	var iter = iterators.RangeStep(5, 5, 1)
	var got = iterators.Collect(ctx, []int{}, iter)

	var expected = []int{}
	if !slices.Equal(got, expected) {
		test.Errorf("got:      %v", got)
		test.Errorf("expected: %v", expected)
	}
}

func TestRangeStep_StepNaN(test *testing.T) {
	test.Parallel()

	var ctx = context.Background()
	var iter = iterators.RangeStep(10, 2, math.NaN())
	var got = iterators.Collect(ctx, []float64{}, iter)

	var expected = []float64{}
	if !slices.Equal(got, nil) {
		test.Errorf("got:      %v", got)
		test.Errorf("expected: %v", expected)
	}
}

func TestRangeStep_FromNaN(test *testing.T) {
	test.Parallel()

	var ctx = context.Background()
	var iter = iterators.RangeStep(math.NaN(), 2, 1)
	var got = iterators.Collect(ctx, []float64{}, iter)

	var expected = []float64{}
	if !slices.Equal(got, nil) {
		test.Errorf("got:      %v", got)
		test.Errorf("expected: %v", expected)
	}
}

func TestRangeStep_ToNaN(test *testing.T) {
	test.Parallel()

	var ctx = context.Background()
	var iter = iterators.RangeStep(1, math.NaN(), 1)
	var got = iterators.Collect(ctx, []float64{}, iter)

	var expected = []float64{}
	if !slices.Equal(got, nil) {
		test.Errorf("got:      %v", got)
		test.Errorf("expected: %v", expected)
	}
}
