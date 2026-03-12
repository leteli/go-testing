package hashutil

import (
	"crypto/sha256"
	"encoding/hex"
)

func HashSHA256(s string) string {
	sum := sha256.Sum256([]byte(s))
	return hex.EncodeToString(sum[:])
}
