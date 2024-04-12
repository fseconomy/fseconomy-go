package exceptions

import (
	"errors"
)

var (
	// BadFseKeyError malformed key
	BadFseKeyError = errors.New("malformed key")

	// BadSecretKeyError malformed secret key
	BadSecretKeyError = errors.New("malformed secret key (must be 16 bytes long)")

	// BadSecretIvError malformed secret IV
	BadSecretIvError = errors.New("malformed secret IV (must be 12 bytes long)")

	// InvalidServiceError service doesn't exist
	InvalidServiceError = errors.New("service does not exist")
)
