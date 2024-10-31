package util

import (
	"crypto/rand"
	"encoding/base64"
)

// Rand is generates a random string of the specified byte size, encoded in base64 URL format.
func Rand(byteSize int) (string, error) {
	b := make([]byte, byteSize)
	_, err := rand.Read(b)
	if err != nil {
		return "", err
	}
	r := base64.URLEncoding.EncodeToString(b)
	return r, nil
}
