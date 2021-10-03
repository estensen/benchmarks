package strings

import (
	"math/rand"
	"testing"
)

func BenchmarkCovertIntToString(b *testing.B) {
	number := rand.Int()

	b.Run("convert fmt", func(b *testing.B) {
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			ConvertFmt(number)
		}
	})

	b.Run("convert strconv", func(b *testing.B) {
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			ConvertStrconv(number)
		}
	})
}

func RandInt() int {
	return rand.Int()
}
