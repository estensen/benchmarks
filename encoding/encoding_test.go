package encoding

import (
	"crypto/rand"
	"testing"
)

func BenchmarkByteToHex(b *testing.B) {
	benchmarks := []struct {
		name        string
		randomBytes []byte
	}{
		{
			name:        "1 byte",
			randomBytes: GenerateRandomBytes(1),
		},
		{
			name:        "5 bytes",
			randomBytes: GenerateRandomBytes(5),
		},
		{
			name:        "10 bytes",
			randomBytes: GenerateRandomBytes(10),
		},
		{
			name:        "50 bytes",
			randomBytes: GenerateRandomBytes(50),
		},
		{
			name:        "100 bytes",
			randomBytes: GenerateRandomBytes(100),
		},
	}

	for _, bm := range benchmarks {
		b.Run("encode bytes to hex string with fmt.Sprintf for "+bm.name, func(b *testing.B) {
			b.ReportAllocs()
			for i := 0; i < b.N; i++ {
				ByteToHexStringSprintf(bm.randomBytes)
			}
		})

		b.Run("encode bytes to hex string with hex.EncodeToString for "+bm.name, func(b *testing.B) {
			b.ReportAllocs()
			for i := 0; i < b.N; i++ {
				ByteToHexStringHex(bm.randomBytes)
			}
		})
	}

}

func GenerateRandomBytes(n int) []byte {
	b := make([]byte, n)
	rand.Read(b)
	return b
}
