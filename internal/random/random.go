package random

import (
	"crypto/rand"
	"encoding/hex"
	"errors"
)

func GenerateHexString() (string, error) {
	// Create random 8 bytes token
	token := make([]byte, 8)

	_, err := rand.Read(token)
	if err != nil {
		return "", errors.New("Failed to generate random bytes.")
	}

	// Encode token into 16 (8 * 2) character hex string
	return hex.EncodeToString(token), nil
}
