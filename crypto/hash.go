package crypto

import (
	"crypto/sha256"

	"golang.org/x/crypto/blake2b"
)

func SHA256(input string) []byte {
	h := sha256.New()
	h.Write([]byte(input))
	return h.Sum(nil)
}

func SHA256Directly(input string) [32]byte {
	return sha256.Sum256([]byte(input))
}

func BLAKE2b(input string) [32]byte {
	return blake2b.Sum256([]byte(input))
}
