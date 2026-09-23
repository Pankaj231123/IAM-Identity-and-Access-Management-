package hash

import (
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"golang.org/x/crypto/argon2"
)

const (
	argonMemory      = 65536
	argonIterations  = 3
	argonParallelism = 2
	argonSaltLength  = 16
	argonKeyLength   = 32
)

func HashPassword(password string) (string, error) {
	// create and fill the salt
	salt := make([]byte, argonSaltLength)
	if _, err := rand.Read(salt); err != nil {
		return "", err
	}
	// derive the hash
	hash := argon2.IDKey([]byte(password), salt, argonIterations, argonMemory, argonParallelism, argonKeyLength)

	// encode the salt to base64
	encodedSalt := base64.RawStdEncoding.EncodeToString(salt)

	// encode the hash to base64
	encodedHash := base64.RawStdEncoding.EncodeToString(hash)

	// assemble and return the final string
	finalHash := fmt.Sprintf("$argon2id$v=%d$m=%d,t=%d,p=%d$%s$%s",
		argon2.Version, argonMemory, argonIterations, argonParallelism, encodedSalt, encodedHash)
	return finalHash, nil
}
