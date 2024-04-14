package security

import (
	"github.com/fseconomy/fseconomy-go/exceptions"
	"github.com/fseconomy/fseconomy-go/internal/security"
)

// Logout terminates the current user session. If successful, the method sets UserName and AuthToken to
// empty strings and returns nil. Otherwise, the method returns the error describing the problem encountered.
func (sec *Security) Logout() error {
	if (sec.UserName() == "") || (sec.AuthToken() == "") {
		return exceptions.UserNotLoggedInError
	}
	logoutService, err := security.GetAuthService("logout")
	if err != nil {
		return err
	}
	resp, err := logoutService.Query(map[string]string{"username": sec.UserName()}, map[string]string{"authtoken": sec.AuthToken()}, nil)
	if err != nil {
		return err
	}
	if resp.Meta.Code == 200 {
		sec.SetAuthToken("")
		sec.SetUserName("")
		return nil
	}
	return exceptions.InvalidCredentialsError
}
