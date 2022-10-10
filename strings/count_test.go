package strings

import "testing"

const rows = "a\nb\nc\nd\ne"

func BenchmarkCount(b *testing.B) {
	b.Run("split strings and count", func(b *testing.B) {
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			StringSplit(rows, "\n")
		}
	})

	b.Run("count newlines", func(b *testing.B) {
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			StringCount(rows, "\n")
		}
	})
}

func TestStringSplit(t *testing.T) {
	want := 5
	got := StringSplit(rows, "\n")
	if want != got {
		t.Errorf("got %d, want %d", got, want)
	}
}

func TestStringCount(t *testing.T) {
	want := 5
	got := StringCount(rows, "\n")
	if want != got {
		t.Errorf("got %d, want %d", got, want)
	}
}
