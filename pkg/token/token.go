package token

import (
	"crypto/rand"
	"crypto/sha256"
	"encoding/base64"
	"encoding/hex"
)

func GenerateOpaqueToken() (string, error) {
	// Generate a random 32-byte token
	tokenBytes := make([]byte, 32)
	_, err := rand.Read(tokenBytes)
	if err != nil {
		return "", err
	}

	// Encode the token to a base64 string
	token := base64.RawURLEncoding.EncodeToString(tokenBytes)

	return token, nil
}

func HashToken(token string) string {
	// hash the token bytes
	hash := sha256.Sum256([]byte(token))

	// hex-encode it
	return hex.EncodeToString(hash[:])


}
