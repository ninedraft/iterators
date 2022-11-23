package iterators

import "context"

func TransduceErr[F, T any](iter Iterator[F], fn func(F) (T, error)) IteratorErr[T] {
	var ierr = &iterErr[T]{}
	ierr.Iterator = Fn[T](func(ctx context.Context) (T, bool) {
		if ierr.err != nil {
			return none[T]()
		}
		var value, ok = iter.Next(ctx)
		if !ok {
			return none[T]()
		}
		var result, err = fn(value)
		if err != nil {
			ierr.err = err
			return none[T]()
		}
		return result, true
	})
	return ierr
}

type iterErr[E any] struct {
	err error
	Iterator[E]
}

func (iter *iterErr[E]) Err() error {
	return iter.err
}

func TransduceErrWrap[F, T any](iter Iterator[F], fn func(F) (T, error)) Pairs[T, error] {
	return Fn[Pair[T, error]](func(ctx context.Context) (Pair[T, error], bool) {
		var value, ok = iter.Next(ctx)
		if !ok {
			return none[Pair[T, error]]()
		}
		var result, err = fn(value)
		return Pair[T, error]{
			A: result,
			B: err,
		}, true
	})
}
