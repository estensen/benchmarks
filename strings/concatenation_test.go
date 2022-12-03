package strings

import (
	"math/rand"
	"testing"
)

const wordlength = 10

var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func BenchmarkConcatString(b *testing.B) {
	benchmarks := []struct {
		name  string
		words []string
	}{
		{
			name:  "len 0",
			words: RandomStrings(0),
		},
		{
			name:  "len 1",
			words: RandomStrings(1),
		},
		{
			name:  "len 5",
			words: RandomStrings(5),
		},
		{
			name:  "len 10",
			words: RandomStrings(10),
		},
		{
			name:  "len 50",
			words: RandomStrings(50),
		},
		{
			name:  "len 1000",
			words: RandomStrings(1000),
		},
	}

	for _, bm := range benchmarks {
		b.Run("concat string for "+bm.name, func(b *testing.B) {
			b.ReportAllocs()
			for i := 0; i < b.N; i++ {
				ConcatString(bm.words)
			}
		})

		b.Run("concat buffer for "+bm.name, func(b *testing.B) {
			b.ReportAllocs()
			for i := 0; i < b.N; i++ {
				ConcatBuffer(bm.words)
			}
		})

		b.Run("concat builder for "+bm.name, func(b *testing.B) {
			b.ReportAllocs()
			for i := 0; i < b.N; i++ {
				ConcatBuilder(bm.words)
			}
		})

		b.Run("concat grown builder for "+bm.name, func(b *testing.B) {
			b.ReportAllocs()
			for i := 0; i < b.N; i++ {
				ConcatGrownBuilder(bm.words)
			}
		})
	}
}

func RandomStrings(n int) []string {
	strings := make([]string, n)
	for i := range strings {
		strings[i] = RandStringBytes(wordlength)
	}
	return strings
}

func RandStringBytes(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}
