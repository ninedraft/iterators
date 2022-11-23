package iterators_test

import (
	"context"
	"testing"

	"github.com/ninedraft/iterators"
	"golang.org/x/exp/maps"
)

func TestCollectMap(test *testing.T) {
	var mapped = iterators.TransduceArg(
		iterators.Range(1, 10),
		func(x int) int {
			return 2 * x
		})

	var got = map[int]int{}
	iterators.CollectMap(context.Background(), got, mapped)

	var expected = map[int]int{
		1: 2, 2: 4, 3: 6,
		4: 8, 5: 10, 6: 12,
		7: 14, 8: 16, 9: 18,
	}
	if !maps.Equal(got, expected) {
		test.Errorf("got:      %v", got)
		test.Errorf("expected: %v", expected)
	}
}
