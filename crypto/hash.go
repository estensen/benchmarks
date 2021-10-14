package crypto

import (
	"crypto/md5"
	"crypto/sha256"

	"github.com/cespare/xxhash/v2"
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

func MD5(input string) [16]byte {
	return md5.Sum([]byte(input))
}

func xxHash(input string) uint64 {
	return xxhash.Sum64([]byte(input))
}
