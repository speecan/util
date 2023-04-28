package slice

func SlowFilter[T any](x []T, fn func(index int, v T) bool) []T {
	res := make([]T, 0, len(x))
	for i, v := range x {
		if fn(i, v) {
			res = append(res, v)
		}
	}
	return res
}

func SlowMap[T any](x []T, fn func(index int, v T) T) []T {
	res := make([]T, len(x))
	for i, v := range x {
		res[i] = fn(i, v)
	}
	return res
}

func (me Slice[T]) SlowFilter(fn func(index int, v T) bool) Slice[T] {
	return Slice[T](SlowFilter(me, fn))
}
func (me Slice[T]) SlowMap(fn func(index int, v T) T) Slice[T] {
	return Slice[T](SlowMap(me, fn))
}
