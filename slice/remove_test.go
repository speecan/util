package slice

import (
	"testing"

	"golang.org/x/exp/slices"
)

func TestRemove(t *testing.T) {
	values := []struct {
		src      []int
		index    int
		expected []int
	}{
		{[]int{1, 2, 3}, 0, []int{2, 3}},
		{[]int{1, 2, 3}, -1, []int{1, 2, 3}},
		{[]int{1, 2, 3}, 1, []int{1, 3}},
		{[]int{1, 2, 3}, 2, []int{1, 2}},
		{[]int{1, 2, 3}, 3, []int{1, 2, 3}},
		{[]int{1, 2, 3}, 5, []int{1, 2, 3}},
		{[]int{5}, 0, []int{}},
		{[]int{5}, 1, []int{5}},
	}
	for _, v := range values {
		res := Remove(v.src, v.index)
		if !slices.Equal(res, v.expected) {
			t.Error("unexpected result", v.index, res, v.expected)
		}
	}
}

func TestRemoveCopy(t *testing.T) {
	values := []struct {
		src      []int
		index    int
		expected []int
	}{
		{[]int{1, 2, 3}, 0, []int{2, 3}},
		{[]int{1, 2, 3}, -1, []int{1, 2, 3}},
		{[]int{1, 2, 3}, 1, []int{1, 3}},
		{[]int{1, 2, 3}, 2, []int{1, 2}},
		{[]int{1, 2, 3}, 3, []int{1, 2, 3}},
		{[]int{1, 2, 3}, 5, []int{1, 2, 3}},
		{[]int{5}, 0, []int{}},
		{[]int{5}, 1, []int{5}},
	}
	for _, v := range values {
		res := RemoveCopy(v.src, v.index)
		if !slices.Equal(res, v.expected) {
			t.Error("unexpected result", v.index, res, v.expected)
		}
	}
}

func TestRemoveUnsorted(t *testing.T) {
	values := []struct {
		src      []int
		index    int
		expected []int
	}{
		{[]int{1, 2, 3}, 0, []int{3, 2}},
		{[]int{1, 2, 3}, -1, []int{1, 2, 3}},
		{[]int{1, 2, 3}, 1, []int{1, 3}},
		{[]int{1, 2, 3}, 2, []int{1, 2}},
		{[]int{1, 2, 3}, 3, []int{1, 2, 3}},
		{[]int{1, 2, 3}, 5, []int{1, 2, 3}},
		{[]int{5}, 0, []int{}},
		{[]int{5}, 1, []int{5}},
	}
	for _, v := range values {
		res := RemoveUnsorted(v.src, v.index)
		if !slices.Equal(res, v.expected) {
			t.Error("unexpected result", v.index, res, v.expected)
		}
	}
}

func BenchmarkRemove(b *testing.B) {
	l := 10000
	value := make([]int, l)
	b.Run("append", func(b *testing.B) {
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			for j := 0; j < l; j++ {
				res := Remove(value, j)
				if len(res) != l-1 {
					b.Error("remove must not be change src slice")
				}
			}
		}
	})
	b.Run("copy", func(b *testing.B) {
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			for j := 0; j < l; j++ {
				res := RemoveCopy(value, j)
				if len(res) != l-1 {
					b.Error("remove must not be change src slice")
				}
			}
		}
	})
	b.Run("unsorted", func(b *testing.B) {
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			for j := 0; j < l; j++ {
				res := RemoveUnsorted(value, j)
				if len(res) != l-1 {
					b.Error("remove must not be change src slice")
				}
			}
		}
	})
}
