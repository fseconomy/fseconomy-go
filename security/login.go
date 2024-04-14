package security

import (
	"github.com/fseconomy/fseconomy-go/exceptions"
	"github.com/fseconomy/fseconomy-go/internal/security"
)

// Login attempts to log in to the FSEconomy server with a given username and password.
// In case of success, the security context's UserName and AuthToken properties are populated,
// and the method returns nil. In case of an error, the method returns the corresponding error.
func (sec *Security) Login(username string, password string) error {
	loginService, err := security.GetAuthService("login")
	if err != nil {
		return err
	}
	resp, err := loginService.Query(map[string]string{"username": username, "password": password}, nil, nil)
	if err != nil {
		return err
	}
	if resp.Meta.Code == 200 {
		sec.SetUserName(username)
		sec.SetAuthToken(resp.Data)
		return nil
	}
	return exceptions.InvalidCredentialsError
}
