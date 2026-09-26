package token
import (
	"crypto/rand"
	"encoding/base64"
)

func GenerateOpaqueToken() (string, error) {
	// Generate a random 32-byte token
	tokenBytes := make([]byte, 32)
	_, err := rand.Read(tokenBytes)
	if err != nil {
		return "", err
	}

	// Encode the token to a base64 string
	token := base64.URLEncoding.EncodeToString(tokenBytes)

	return token, nil
}
