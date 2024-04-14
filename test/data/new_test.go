package data_test

import (
	"github.com/fseconomy/fseconomy-go/data"
	"os"
	"testing"
)

const (
	keyGood     string = "2ad0394ad6c8e400"
	keyBadShort string = "2ad0394ad6c8e4"
	keyBadInval string = "im no hex string"
)

func setupGood() {
	_ = os.Setenv("FSE_SERVICE_KEY", keyGood)
	_ = os.Setenv("FSE_ACCESS_KEY", keyGood)
	_ = os.Setenv("FSE_USER_KEY", keyGood)
}

func setupBadShort() {
	_ = os.Setenv("FSE_SERVICE_KEY", keyBadShort)
	_ = os.Setenv("FSE_ACCESS_KEY", keyBadShort)
	_ = os.Setenv("FSE_USER_KEY", keyBadShort)
}

func setupBadInval() {
	_ = os.Setenv("FSE_SERVICE_KEY", keyBadInval)
	_ = os.Setenv("FSE_ACCESS_KEY", keyBadInval)
	_ = os.Setenv("FSE_USER_KEY", keyBadInval)
}

func tearDown() {
	_ = os.Unsetenv("FSE_SERVICE_KEY")
	_ = os.Unsetenv("FSE_ACCESS_KEY")
	_ = os.Unsetenv("FSE_USER_KEY")
}

func TestNewBlank(t *testing.T) {
	_, err := data.New()
	if err != nil {
		t.Errorf("Unexpected error when creating a blank data context: %v", err)
	}
}

func TestNewEnvironGood(t *testing.T) {
	setupGood()
	defer tearDown()
	_, err := data.New()
	if err != nil {
		t.Errorf("Unexpected error when creating a blank data context: %v", err)
	}
}

func TestNewEnvironBadShort(t *testing.T) {
	setupBadShort()
	defer tearDown()
	_, err := data.New()
	if err == nil {
		t.Errorf("Creating new data context should have failed: expected BadFseKeyError, got: %v", err)
	}
}

func TestNewEnvironBadInval(t *testing.T) {
	setupBadInval()
	defer tearDown()
	_, err := data.New()
	if err == nil {
		t.Errorf("Creating new data context should have failed: expected BadFseKeyError, got: %v", err)
	}
}

func TestWithServiceKey(t *testing.T) {
	d, err := data.New(data.WithServiceKey(keyGood))
	if err != nil {
		t.Errorf("Unexpected error when creating a data context with service key: %v", err)
	}
	if d.ServiceKey() != keyGood {
		t.Errorf("Service key does not match expected value, expected: %s, got: %s", keyGood, d.ServiceKey())
	}
}

func TestWithAccessKey(t *testing.T) {
	d, err := data.New(data.WithAccessKey(keyGood))
	if err != nil {
		t.Errorf("Unexpected error when creating a data context with access key: %v", err)
	}
	if d.AccessKey() != keyGood {
		t.Errorf("Access key does not match expected value, expected: %s, got: %s", keyGood, d.AccessKey())
	}
}

func TestWithUserKey(t *testing.T) {
	d, err := data.New(data.WithUserKey(keyGood))
	if err != nil {
		t.Errorf("Unexpected error when creating a data context with user key: %v", err)
	}
	if d.UserKey() != keyGood {
		t.Errorf("User key does not match expected value, expected: %s, got: %s", keyGood, d.UserKey())
	}
}
