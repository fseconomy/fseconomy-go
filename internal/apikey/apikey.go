package apikey

import "encoding/hex"

// ValidateKey checks if a apikey string represents a valid FSEconomy API apikey.
func ValidateKey(key string) bool {
	decoded, err := hex.DecodeString(key)
	if err != nil {
		return false
	}
	return len(decoded) == 8
}
