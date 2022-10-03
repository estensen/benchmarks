package slices

import "testing"

func BenchmarkSlicePointers(b *testing.B) {
	b.ReportAllocs()

	for i := 0; i < b.N; i++ {
		AppendSlicePointers(i)
	}
}

func BenchmarkSliceNoPointers(b *testing.B) {
	b.ReportAllocs()

	for i := 0; i < b.N; i++ {
		AppendSliceNoPointers(i)
	}
}

func BenchmarkSliceHybrid(b *testing.B) {
	b.ReportAllocs()

	for i := 0; i < b.N; i++ {
		AppendSliceHybrid(i)
	}
}
