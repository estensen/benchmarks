package strings

import "testing"

func BenchmarkMatchString(b *testing.B) {
	longString := RandStringBytes(20)
	pattern := "haavard@example.com"

	b.Run("match string", func(b *testing.B) {
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			MatchString(pattern, longString)
		}
	})
}

func BenchmarkMatchStringCompiled(b *testing.B) {
	longString := RandStringBytes(20)
	pattern := `^[A-Za-z0-9._%+-]+@[A-Za-z0-9.-]+\.[A-Za-z]+$`
	compiledPattern := CompilePattern(pattern)

	b.Run("match string", func(b *testing.B) {
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			MatchStringCompiled(compiledPattern, longString)
		}
	})
}
