package iterators

import (
	"context"
	"reflect"
	"unsafe"

	"golang.org/x/exp/maps"
)

// Map returns a sequence of key-value pairs from provided map.
// Resulting iterator follows the same iteration semantics as a range statement
func Map[K comparable, V any, M ~map[K]V](m M) Sized[Pair[K, V]] {
	var keySizes = unsafe.Sizeof(empty[K]()) * uintptr(len(m))
	var result = &sizeProxy[Pair[K, V]]{
		len: func() int { return len(m) },
	}

	switch {
	case keySizes <= 2048:
		var keys = Slice(maps.Keys(m))
		result.Iterator = TransduceFilter[K](keys, func(key K) (Pair[K, V], bool) {
			var value, ok = m[key]
			if !ok {
				return none[Pair[K, V]]()
			}
			return Pair[K, V]{A: key, B: value}, true
		})
	default:
		var mapIter = reflect.ValueOf(m).MapRange()
		result.Iterator = Fn[Pair[K, V]](func(context.Context) (Pair[K, V], bool) {
			if !mapIter.Next() {
				return none[Pair[K, V]]()
			}
			var key = mapIter.Key().Interface().(K)
			var value = mapIter.Value().Interface().(V)
			return Pair[K, V]{A: key, B: value}, true
		})
	}
	return result
}

type sizeProxy[E any] struct {
	len func() int
	Iterator[E]
}

func (iter *sizeProxy[E]) Len() int {
	return iter.len()
}
