package encoding

import (
	"encoding/hex"
	"fmt"
)

func ByteToHexStringSprintf(b []byte) string {
	return fmt.Sprintf("%x", b)
}

func ByteToHexStringHex(b []byte) string {
	return hex.EncodeToString(b)
}
