package security

import (
	"github.com/fseconomy/fseconomy-go/exceptions"
	"github.com/fseconomy/fseconomy-go/internal/security"
)

// Check verifies the presented token belongs to a valid session. If the token is valid,
// the method returns nil. In case of an error, the method returns the corresponding error.
func (sec *Security) Check() error {
	if (sec.UserName() == "") || (sec.AuthToken() == "") {
		return exceptions.UserNotLoggedInError
	}
	checkService, err := security.GetAuthService("check")
	if err != nil {
		return err
	}
	resp, err := checkService.Query(nil, map[string]string{"authToken": sec.AuthToken()}, nil)
	if err != nil {
		return err
	}
	if resp.Meta.Code == 200 {
		return nil
	}
	return exceptions.InvalidCredentialsError
}
