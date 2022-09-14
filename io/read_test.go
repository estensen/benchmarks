package io

import (
	"bytes"
	"testing"
)

func BenchmarkReaders(b *testing.B) {
	data := make([]byte, 100)
	for i := range data {
		data[i] = byte(i)
	}

	reader := bytes.NewReader(data)

	b.Run("ReadAll", func(b *testing.B) {
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			IOReadAll(reader)
		}
	})

	b.Run("Copy", func(b *testing.B) {
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			IOCopy(reader)
		}
	})

	b.Run("ReadFrom", func(b *testing.B) {
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			IOBufferReadFrom(reader)
		}
	})
}
