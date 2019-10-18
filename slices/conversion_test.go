package slices

import "testing"

func BenchmarkSliceConversion(b *testing.B) {
	numbers := make([]int, 100)
	for i := range numbers {
		numbers[i] = i
	}

	b.Run("bad", func(b *testing.B) {
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			NumbersToStringsBad(numbers)
		}
	})

	b.Run("better", func(b *testing.B) {
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			NumbersToStringsBetter(numbers)
		}
	})

	b.Run("best", func(b *testing.B) {
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			NumbersToStringsBest(numbers)
		}
	})
}
