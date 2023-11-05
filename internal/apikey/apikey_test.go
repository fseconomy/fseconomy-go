package apikey

import "testing"

// TestValidateKeyGood calls apikey.ValidateKey with a valid apikey, checking for a valid return.
func TestValidateKeyGood(t *testing.T) {
	TestKey := "2ad0394ad6c8e400"
	if !ValidateKey(TestKey) {
		t.Fatalf(`ValidateKey(%q) expected to be true`, TestKey)
	}
}

// TestValidateKeyTooShort calls apikey.ValidateKey with a too short apikey, checking for an error.
func TestValidateKeyTooShort(t *testing.T) {
	TestKey := "2ad0394ad6c8e4"
	if ValidateKey(TestKey) {
		t.Fatalf(`ValidateKey(%q) expected to be false`, TestKey)
	}
}

// TestValidateKeyUneven calls apikey.ValidateKey with an uneven number of hex digits, checking for an error.
func TestValidateKeyUneven(t *testing.T) {
	TestKey := "2ad0394ad6c8e40"
	if ValidateKey(TestKey) {
		t.Fatalf(`ValidateKey(%q) expected to be false`, TestKey)
	}
}

// TestValidateKeyNoHex calls apikey.ValidateKey with an invalid hex string, checking for an error.
func TestValidateKeyNoHex(t *testing.T) {
	TestKey := "2ad0394ad6c8e40g"
	if ValidateKey(TestKey) {
		t.Fatalf(`ValidateKey(%q) expected to be false`, TestKey)
	}
}
