package strings

import "testing"

func BenchmarkEqualStrings(b *testing.B) {
	first := "hello, World!"
	second := "HelLo, world!"

	b.Run("equal strings op", func(b *testing.B) {
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			EqualOp(first, second)
		}
	})

	b.Run("equal strings fold", func(b *testing.B) {
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			EqualFold(first, second)
		}
	})
}

