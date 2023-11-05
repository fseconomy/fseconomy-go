package key

import "testing"

// TestValidateKeyGood calls key.ValidateKey with a valid key, checking for a valid return.
func TestValidateKeyGood(t *testing.T) {
	TestKey := "2ad0394ad6c8e400"
	if !ValidateKey(TestKey) {
		t.Fatalf(`ValidateKey(%q) expected to be true`, TestKey)
	}
}

// TestValidateKeyTooShort calls key.ValidateKey with a too short key, checking for an error.
func TestValidateKeyTooShort(t *testing.T) {
	TestKey := "2ad0394ad6c8e4"
	if ValidateKey(TestKey) {
		t.Fatalf(`ValidateKey(%q) expected to be false`, TestKey)
	}
}

// TestValidateKeyUneven calls key.ValidateKey with an uneven number of hex digits, checking for an error.
func TestValidateKeyUneven(t *testing.T) {
	TestKey := "2ad0394ad6c8e40"
	if ValidateKey(TestKey) {
		t.Fatalf(`ValidateKey(%q) expected to be false`, TestKey)
	}
}

// TestValidateKeyNoHex calls key.ValidateKey with an invalid hex string, checking for an error.
func TestValidateKeyNoHex(t *testing.T) {
	TestKey := "2ad0394ad6c8e40g"
	if ValidateKey(TestKey) {
		t.Fatalf(`ValidateKey(%q) expected to be false`, TestKey)
	}
}
