package iterators_test

import (
	"context"
	"testing"

	"github.com/ninedraft/iterators"
	"golang.org/x/exp/slices"
)

func TestZip(test *testing.T) {
	test.Parallel()
	test.Log(
		"Zip iterator must emit pairs from connected values of two sequences of different legths",
	)
	var aa = iterators.Range(1, 10)
	var bb = iterators.Range(100, 110)
	var pairs = iterators.Zip(aa, bb)
	var got = iterators.Collect[iterators.Pair[int, int]](context.Background(),
		[]iterators.Pair[int, int]{},
		pairs)

	var expected = []iterators.Pair[int, int]{
		{1, 100}, {2, 101}, {3, 102}, {4, 103},
		{5, 104}, {6, 105}, {7, 106}, {8, 107},
		{9, 108},
	}

	if !slices.Equal(got, expected) {
		test.Errorf("got:      %v", got)
		test.Errorf("expected: %v", expected)
	}
}

func TestZip_Empty(test *testing.T) {
	test.Parallel()
	test.Log(
		"Zip iterator must emit no pairs from connected values of two sequences of different legths",
	)
	var aa = iterators.Empty[int]()
	var bb = iterators.Empty[int]()
	var pairs = iterators.Zip(aa, bb)
	var got = iterators.Collect[iterators.Pair[int, int]](context.Background(),
		[]iterators.Pair[int, int]{},
		pairs)

	var expected = []iterators.Pair[int, int]{}

	if !slices.Equal(got, expected) {
		test.Errorf("got:      %v", got)
		test.Errorf("expected: %v", expected)
	}
}

func TestZip_CallsFirstDrained(test *testing.T) {
	test.Parallel()
	test.Log(
		"Zip iterator must emit pairs from connected values of two sequences of different legths",
	)
	var aa = iterators.Range(1, 2)
	var bb = iterators.Range(1, 100)
	var pairs = iterators.Zip(aa, bb)
	var ctx = context.Background()

	for i := 0; i < 10; i++ {
		_, _ = pairs.Next(ctx)
	}

	var b, ok = bb.Next(ctx)
	if !ok {
		test.Errorf("bb iterator must be not drained")
	}
	if b != 2 {
		test.Errorf("bb expected to emit 2, got %d ", b)
	}
}

func TestZip_CallsSecondDrained(test *testing.T) {
	test.Parallel()
	test.Log(
		"Zip iterator must emit pairs from connected values of two sequences of different legths",
	)
	var aa = iterators.Range(1, 100)
	var bb = iterators.Range(1, 2)
	var pairs = iterators.Zip(aa, bb)
	var ctx = context.Background()

	for i := 0; i < 10; i++ {
		_, _ = pairs.Next(ctx)
	}

	var a, ok = aa.Next(ctx)
	if !ok {
		test.Errorf("aa iterator must be not drained")
	}
	if a != 3 {
		test.Errorf("aa expected to emit 3, got %d ", a)
	}
}
