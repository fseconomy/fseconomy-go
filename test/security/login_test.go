package security_test

import (
	"errors"
	"github.com/fseconomy/fseconomy-go/exceptions"
	"github.com/fseconomy/fseconomy-go/security"
	"testing"
)

func TestLogin(t *testing.T) {
	sec, err := security.New()
	if err != nil {
		t.Errorf("Unexpected error with New(): %v", err)
	}
	err = sec.Login("test", "test")
	if (err != nil) && !errors.Is(err, exceptions.ServerMaintenanceError) && !errors.Is(err, exceptions.InvalidCredentialsError) {
		t.Errorf("Unexpected error with Login(): %v", err)
	}
}
