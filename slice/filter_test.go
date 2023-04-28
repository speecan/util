package slice

import (
	"testing"

	"golang.org/x/exp/slices"
)

func TestFilter(t *testing.T) {
	x := Slice[int]([]int{1, 2, 3, 4, 5, 6, 7})
	expected := []int{4, 5, 6, 7}
	res := x.SlowFilter(func(index, v int) bool {
		return v > 3
	})
	if !slices.Equal(expected, res) {
		t.Error("unexpected result", expected, res)
	}
}

func TestMap(t *testing.T) {
	x := Slice[int]([]int{1, 2, 3, 4, 5, 6, 7})
	expected := []int{3, 6, 9, 12, 15, 18, 21}
	res := x.SlowMap(func(index, v int) int {
		return v * 3
	})
	if !slices.Equal(expected, res) {
		t.Error("unexpected result", expected, res)
	}
}

func BenchmarkFilter(b *testing.B) {
	l := 10000
	value := make([]int, l)
	filterFn := func(index int, v int) bool {
		return index%2 == 0
	}
	mapFn := func(index int, v int) int {
		return v * 2
	}
	b.Run("filter", func(b *testing.B) {
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			for j := 0; j < l; j++ {
				res := SlowFilter(value, filterFn)
				if len(res) != l/2 {
					b.Error("filtered slice must have half length")
				}
			}
		}
	})
	b.Run("map", func(b *testing.B) {
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			for j := 0; j < l; j++ {
				res := SlowMap(value, mapFn)
				if len(res) != l {
					b.Error("maped slice must have same length")
				}
			}
		}
	})
}
