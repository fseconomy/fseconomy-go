package security_test

import (
	"errors"
	"github.com/fseconomy/fseconomy-go/exceptions"
	"github.com/fseconomy/fseconomy-go/security"
	"testing"
)

func TestCheck(t *testing.T) {
	sec, err := security.New()
	if err != nil {
		t.Errorf("Unexpected error while creating security context: %v", err)
	}
	sec.SetUserName("test")
	sec.SetAuthToken("abcdef1234567890")
	err = sec.Check()
	if (err != nil) && !errors.Is(err, exceptions.ServerMaintenanceError) && !errors.Is(err, exceptions.InvalidCredentialsError) {
		t.Errorf("Unexpected error while attempting to log out: %v", err)
	}
}
