package key

import "testing"

// TestValidateKeyGood calls key.validateKey with a valid key, checking for a valid return.
func TestValidateKeyGood(t *testing.T) {
	TestKey := "2ad0394ad6c8e400"
	if !validateKey(TestKey) {
		t.Fatalf(`validateKey(%q) expected to be true`, TestKey)
	}
}

// TestValidateKeyTooShort calls key.validateKey with a too short key, checking for an error.
func TestValidateKeyTooShort(t *testing.T) {
	TestKey := "2ad0394ad6c8e4"
	if validateKey(TestKey) {
		t.Fatalf(`validateKey(%q) expected to be false`, TestKey)
	}
}

// TestValidateKeyUneven calls key.validateKey with an uneven number of hex digits, checking for an error.
func TestValidateKeyUneven(t *testing.T) {
	TestKey := "2ad0394ad6c8e40"
	if validateKey(TestKey) {
		t.Fatalf(`validateKey(%q) expected to be false`, TestKey)
	}
}

// TestValidateKeyNoHex calls key.validateKey with an invalid hex string, checking for an error.
func TestValidateKeyNoHex(t *testing.T) {
	TestKey := "2ad0394ad6c8e40g"
	if validateKey(TestKey) {
		t.Fatalf(`validateKey(%q) expected to be false`, TestKey)
	}
}
