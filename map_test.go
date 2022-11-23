package iterators_test

import (
	"context"
	"testing"

	"github.com/ninedraft/iterators"
	"golang.org/x/exp/slices"
)

func TestMapSmall(test *testing.T) {
	test.Parallel()

	var ctx = context.Background()
	var iter = iterators.Map(map[string]int{
		"a": 1,
		"b": 2,
		"c": 3,
	})
	var got = iterators.Collect[iterators.Pair[string, int]](ctx, []iterators.Pair[string, int]{}, iter)
	slices.SortFunc(got, func(a, b iterators.Pair[string, int]) bool {
		return a.A < b.A
	})

	var expected = []iterators.Pair[string, int]{
		{A: "a", B: 1},
		{A: "b", B: 2},
		{A: "c", B: 3},
	}

	if !slices.Equal(got, expected) {
		test.Errorf("got:      %v", got)
		test.Errorf("expected: %v", expected)
	}
}

func TestMapSmall_Mutate(test *testing.T) {
	test.Parallel()

	var ctx = context.Background()
	var m = map[string]int{
		"a": 1,
		"b": 2,
		"c": 3,
	}
	var iter = iterators.Map(m)
	iterators.ForEach[iterators.Pair[string, int]](ctx, iter, func(iterators.Pair[string, int]) bool {
		delete(m, "c")
		return true
	})
}

func TestMapLen(test *testing.T) {
	test.Parallel()

	var ctx = context.Background()
	var m = map[string]int{
		"a": 1,
		"b": 2,
		"c": 3,
	}
	var iter = iterators.Map(m)
	var n = len(m)
	iterators.ForEach[iterators.Pair[string, int]](ctx, iter, func(iterators.Pair[string, int]) bool {
		n--
		if iter.Len() != n {
			test.Errorf("iterator expected to have len=%d, but have %d", n, iter.Len())
		}
		return true
	})
}

func TestMapBig(test *testing.T) {
	test.Parallel()

	const N = 10_000
	var ctx = context.Background()
	var m = map[int]int{}
	for i := 0; i < N; i++ {
		m[i] = i
	}
	var iter = iterators.Map(m)
	var got = iterators.Collect[iterators.Pair[int, int]](ctx, []iterators.Pair[int, int]{}, iter)
	slices.SortFunc(got, func(a, b iterators.Pair[int, int]) bool {
		return a.A < b.A
	})

	var expected = []iterators.Pair[int, int]{}
	for i := 0; i < N; i++ {
		expected = append(expected, iterators.Pair[int, int]{
			A: i, B: i,
		})
	}

	if !slices.Equal(got, expected) {
		test.Errorf("got:      %v", got)
		test.Errorf("expected: %v", expected)
	}
}

func TestMapBig_Mutate(test *testing.T) {
	test.Parallel()

	const N = 10_000
	var ctx = context.Background()
	var m = map[int]int{}
	for i := 0; i < N; i++ {
		m[i] = i
	}
	var iter = iterators.Map(m)
	iterators.ForEach[iterators.Pair[int, int]](ctx, iter, func(p iterators.Pair[int, int]) bool {
		delete(m, p.A+1)
		return true
	})
}
