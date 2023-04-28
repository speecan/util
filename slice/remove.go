package slice

func Remove[T any](src []T, i int) []T {
	if isOutOfSlice(src, i) {
		return src
	}
	j := i + 1
	if j > len(src) {
		j = i
	}
	return append(src[:i], src[j:]...)
}

func RemoveCopy[T any](src []T, i int) []T {
	if isOutOfSlice(src, i) {
		return src
	}
	j := i + 1
	if j > len(src) {
		j = i
	}
	return src[:i+copy(src[i:], src[j:])]
}

func RemoveUnsorted[T any](src []T, i int) []T {
	if isOutOfSlice(src, i) {
		return src
	}
	src[i] = src[len(src)-1]
	return src[:len(src)-1]
}

func isOutOfSlice[T any](src []T, i int) bool {
	if i < 0 || i >= len(src) {
		return true
	}
	return false
}

func (me Slice[T]) Remove(i int) Slice[T] {
	return Slice[T](Remove(me, i))
}
func (me Slice[T]) RemoveCopy(i int) Slice[T] {
	return Slice[T](RemoveCopy(me, i))
}
func (me Slice[T]) RemoveUnsorted(i int) Slice[T] {
	return Slice[T](RemoveUnsorted(me, i))
}
