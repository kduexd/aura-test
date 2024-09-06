package utils

import (
	"crypto/rand"
	"encoding/hex"
)

func GenerateToken() string {
	b := make([]byte, 32)
	if _, err := rand.Read(b); err != nil {
		return ""
	}
	return hex.EncodeToString(b)
}
