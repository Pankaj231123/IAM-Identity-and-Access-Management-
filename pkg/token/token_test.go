package token

import "testing"

func TestGenerateOpaqueToken(t *testing.T) {
	// Check 1: no error, non-empty result
	token1, err := GenerateOpaqueToken()
	if err != nil {
		t.Fatalf("GenerateOpaqueToken returned an error: %v", err)
	}
	if token1 == "" {
		t.Errorf("GenerateOpaqueToken returned an empty string")
	}

	// Check 2: two calls produce different tokens
	token2, err := GenerateOpaqueToken()
	if err != nil {
		t.Fatalf("GenerateOpaqueToken returned an error on second call: %v", err)
	}
	if token1 == token2 {
		t.Errorf("expected two different tokens, got the same value twice: %s", token1)
	}
}

func TestHashToken(t *testing.T) {
	// Check 3: same input produces same hash, every time
	input := "some-fixed-test-value"

	hash1 := HashToken(input)
	hash2 := HashToken(input)

	if hash1 != hash2 {
		t.Errorf("expected same hash for same input, got %s and %s", hash1, hash2)
	}
}