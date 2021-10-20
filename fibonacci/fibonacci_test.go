package fibonacci

import "testing"

func TestRecursiveFibonacci(t *testing.T) {
	tests := []struct {
		n    int
		want int
	}{
		{0, 0},
		{1, 1},
		{2, 1},
		{3, 2},
		{4, 3},
		{5, 5},
		{6, 8},
		{10, 55},
		{42, 267914296},
	}

	for _, test := range tests {
		if got := RecursiveFibonacci(test.n); got != test.want {
			t.Errorf("got %d, want %d", got, test.want)
		}
	}
}

func TestSequentialFibonacci(t *testing.T) {
	tests := []struct {
		n    int
		want int
	}{
		{0, 0},
		{1, 1},
		{2, 1},
		{3, 2},
		{4, 3},
		{5, 5},
		{6, 8},
		{10, 55},
		{42, 267914296},
	}

	for _, test := range tests {
		if got := SequentialFibonacci(test.n); got != test.want {
			t.Errorf("got %d, want %d", got, test.want)
		}
	}
}

func BenchmarkRecursiveFibonacci_10(b *testing.B) {
	for i := 0; i < b.N; i++ {
		RecursiveFibonacci(10)
	}
}

func BenchmarkRecursiveFibonacci_20(b *testing.B) {
	for i := 0; i < b.N; i++ {
		RecursiveFibonacci(20)
	}
}

func BenchmarkSequentialFibonacci_10(b *testing.B) {
	for i := 0; i < b.N; i++ {
		SequentialFibonacci(10)
	}
}

func BenchmarkSequentialFibonacci_20(b *testing.B) {
	for i := 0; i < b.N; i++ {
		SequentialFibonacci(20)
	}
}
