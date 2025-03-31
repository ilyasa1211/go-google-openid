package utils

import (
	"crypto/rand"
	"encoding/hex"
)

func GenerateRandomState() string {
	b := make([]byte, 30)
	_, _ = rand.Read(b)

	return hex.EncodeToString(b)
}
