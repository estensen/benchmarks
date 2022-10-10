package strings

import "testing"

const hello = "hello.txt"
const dir = "tmp"

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

func BenchmarkConcatFilename(b *testing.B) {
	b.Run("concat with print", func(b *testing.B) {
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			FileConcatPrint(dir, hello)
		}
	})
	b.Run("concat", func(b *testing.B) {
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			FileConcat(dir, hello)
		}
	})
}
