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
