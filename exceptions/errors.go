package exceptions

import (
	"errors"
)

var (
	// BadKeyError malformed key
	BadKeyError = errors.New("malformed key")

	// InvalidServiceError service doesn't exist
	InvalidServiceError = errors.New("service does not exist")
)
