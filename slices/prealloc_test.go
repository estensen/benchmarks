package slices

import (
	"math"
	"math/rand"
	"testing"
)

func BenchmarkAppendSlice(b *testing.B) {
	numbers := RandIntSlices(10)

	b.Run("no pre-alloc", func(b *testing.B) {
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			AppendInts(numbers)
		}
	})

	b.Run("pre-alloc", func(b *testing.B) {
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			AppendIntsPreAlloc(numbers)
		}
	})
}

func RandIntSlices(n int) []int {
	result := make([]int, 0, n)
	for i := 0; i < n; i++ {
		result = append(result, rand.Intn(math.MaxInt))
	}
	return result
}
