package strings

import "testing"

const hello = "hello.txt"

func BenchmarkFileSplit(b *testing.B) {
	b.Run("split string to parse", func(b *testing.B) {
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			FileSplit(hello)
		}
	})

	b.Run("cut string to parse", func(b *testing.B) {
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			FileSplit(hello)
		}
	})
}
