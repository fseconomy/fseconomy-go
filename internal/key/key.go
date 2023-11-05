package key

import "encoding/hex"

// validateKey checks if a key string represents a valid FSEconomy API key.
func validateKey(key string) bool {
	decoded, err := hex.DecodeString(key)
	if err != nil {
		return false
	}
	return len(decoded) == 8
}
