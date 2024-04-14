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

	// ServerMaintenanceError FSEconomy server is in maintenance
	ServerMaintenanceError = errors.New("server is in maintenance")

	// MissingParamError for query
	MissingParamError = errors.New("missing parameter for query")

	// InvalidContentTypeError FSEconomy server response has unexpected content type
	InvalidContentTypeError = errors.New("invalid content type")

	// InvalidCredentialsError FSEconomy server rejected the login attempt with the given credentials
	InvalidCredentialsError = errors.New("invalid user credentials (username, password, authtoken)")

	// UserNotLoggedInError the security context doesn't hold a logged-in user
	UserNotLoggedInError = errors.New("user not logged in")

	// FseDataKeyError could not find a valid data access key (user or access key)
	FseDataKeyError = errors.New("could not find a valid data access key (user or access key)")

	// FseAuthKeyError could not find a valid authentication key (user or service key)
	FseAuthKeyError = errors.New("could not find a valid authentication key (user or service key)")
)
