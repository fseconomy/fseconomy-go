package key

import "encoding/hex"

// ValidateKey checks if a key string represents a valid FSEconomy API key.
func ValidateKey(key string) bool {
	decoded, err := hex.DecodeString(key)
	if err != nil {
		return false
	}
	return len(decoded) == 8
}
