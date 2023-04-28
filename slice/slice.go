package slice

type (
	Slice[T any] []T
)

func (me Slice[T]) Slice() []T {
	return me
}
