package jobqueue

import (
	"crypto/rand"
	"encoding/hex"
)

func newTestSuffix() string {
	b := make([]byte, 6)
	_, _ = rand.Read(b)
	return hex.EncodeToString(b)
}
