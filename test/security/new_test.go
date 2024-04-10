package security_test

import (
	"github.com/fseconomy/fseconomy-go/security"
	"os"
	"testing"
)

const testKeyGood string = "2ad0394ad6c8e400"
const testKeyBad string = "2ad0394ad6"

func SetupGood() {
	_ = os.Setenv("FSE_SERVICE_KEY", testKeyGood)
	_ = os.Setenv("FSE_USER_KEY", testKeyGood)
	_ = os.Setenv("FSE_ACCESS_KEY", testKeyGood)
}

func SetupBad() {
	_ = os.Setenv("FSE_SERVICE_KEY", testKeyBad)
	_ = os.Setenv("FSE_USER_KEY", testKeyBad)
	_ = os.Setenv("FSE_ACCESS_KEY", testKeyBad)
}

func TearDown() {
	_ = os.Unsetenv("FSE_SERVICE_KEY")
	_ = os.Unsetenv("FSE_USER_KEY")
	_ = os.Unsetenv("FSE_ACCESS_KEY")
}

func TestNewBlank(t *testing.T) {
	_, err := security.New()
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
}

func TestNewEnvironGood(t *testing.T) {
	SetupGood()
	defer TearDown()
	_, err := security.New()
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
}

func TestNewEnvironBad(t *testing.T) {
	SetupBad()
	defer TearDown()
	_, err := security.New()
	if err == nil {
		t.Errorf("Expected invalid key error, got: %v", err)
	}
}

func TestNewInvalid(t *testing.T) {

}
