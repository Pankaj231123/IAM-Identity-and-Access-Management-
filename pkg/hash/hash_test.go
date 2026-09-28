package hash

import "testing"

func TestHashAndVerifyPassword(t *testing.T) {
	password := "Correct horse battery staple!42"

	encodedHash, err := HashPassword(password)
	if err != nil {
		t.Fatalf("HashPassword() error = %v", err)
	}
	if encodedHash == "" {
		t.Fatal("HashPassword() returned an empty hash")
	}
	if encodedHash == password {
		t.Fatal("HashPassword() returned the plaintext password")
	}

	matched, err := VerifyPassword(password, encodedHash)
	if err != nil {
		t.Fatalf("VerifyPassword() with correct password error = %v", err)
	}
	if !matched {
		t.Error("VerifyPassword() did not match the correct password")
	}

	matched, err = VerifyPassword("wrong password", encodedHash)
	if err != nil {
		t.Fatalf("VerifyPassword() with wrong password error = %v", err)
	}
	if matched {
		t.Error("VerifyPassword() matched an incorrect password")
	}
}
