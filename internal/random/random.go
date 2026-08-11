package random

import (
	"crypto/rand"
	"encoding/base64"
	"errors"
)

func GenerateBase64String() (string, error) {
	// Create random 8 bytes token
	token := make([]byte, 8)

	_, err := rand.Read(token)
	if err != nil {
		return "", errors.New("Failed to generate random bytes.")
	}

	// Encode token into base64 string
	return base64.RawURLEncoding.EncodeToString(token)[:7], nil
}
