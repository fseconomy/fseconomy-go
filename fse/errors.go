package fse

import (
	"errors"
)

var (
	// ErrBadKey malformed key
	ErrBadKey = errors.New("malformed key")

	// ErrInvalidService service doesn't exist
	ErrInvalidService = errors.New("service does not exist")
)
