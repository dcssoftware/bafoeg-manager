package random

import (
	"crypto/rand"
	"encoding/hex"
)

func generateRandomHash() (string, error) {
	randomBytesLength := 16
	bytes := make([]byte, randomBytesLength)
	if _, err := rand.Read(bytes); err != nil {
		return "", err
	}
	return hex.EncodeToString(bytes), nil
}
