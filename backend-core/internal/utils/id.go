package utils

import (
	"crypto/rand"
	"encoding/hex"
)

func NewID() string {
	b := make([]byte, 16)
	_, err := rand.Read(b)
	if err != nil {
		return ""
	}
	return hex.EncodeToString(b)
}
