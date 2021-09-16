package crypto

import (
	"math/rand"
	"testing"
)

func BenchmarkHashString(b *testing.B) {
	word := RandStringBytes(10)

	b.Run("SHA256", func(b *testing.B) {
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			SHA256(word)
		}
	})

	b.Run("SHA256 directly", func(b *testing.B) {
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			SHA256Directly(word)
		}
	})

	b.Run("BLAKE2b", func(b *testing.B) {
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			BLAKE2b(word)
		}
	})
}

var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func RandStringBytes(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}
