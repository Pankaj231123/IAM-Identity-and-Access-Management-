package hash

import (
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"golang.org/x/crypto/argon2"
	"errors"
	"strings"
	"crypto/subtle"
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

func VerifyPassword(password, encodedHash string) (bool, error) {
	// Split the encoded string apart

	parts := strings.Split(encodedHash, "$")
	if len(parts) != 6 {
		return false, errors.New("invalid hash format")
	}

	// Parse out memory, iterations, parallelism from parts[3]
	var memory, iterations uint32
	var parallelism uint8
	_, err := fmt.Sscanf(parts[3], "m=%d,t=%d,p=%d", &memory, &iterations, &parallelism)
	if err != nil {
		return false, err
	}

	// Decode the salt and hash back to raw bytes
	salt, err := base64.RawStdEncoding.DecodeString(parts[4])
	if err != nil {
		return false, err
	}
	storedHash, err := base64.RawStdEncoding.DecodeString(parts[5])
	if err != nil {
		return false, err
	}

	// Re-derive the hash from the given password
	computedHash := argon2.IDKey([]byte(password), salt, iterations, memory, parallelism, uint32(len(storedHash)))

	// Constant-time comparison
	if subtle.ConstantTimeCompare(computedHash, storedHash) == 1 {
		return true, nil
	}
	return false, nil
}
