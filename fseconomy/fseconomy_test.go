package fseconomy

import (
	"os"
	"testing"
)

const testKey string = "2ad0394ad6c8e400"

func Setup() {
	_ = os.Setenv("FSE_SERVICE_KEY", testKey)
	_ = os.Setenv("FSE_USER_KEY", testKey)
	_ = os.Setenv("FSE_ACCESS_KEY", testKey)
}

func TearDown() {
	_ = os.Unsetenv("FSE_SERVICE_KEY")
	_ = os.Unsetenv("FSE_USER_KEY")
	_ = os.Unsetenv("FSE_ACCESS_KEY")
}

// TestNewNoEnviron calls New without providing environment variables
func TestNewNoEnviron(t *testing.T) {
	f, err := New()
	if err != nil {
		t.Fatalf(`Unexpected error: %v`, err)
	}
	if (f.AccessKey != "") || (f.ServiceKey != "") || (f.UserKey != "") {
		t.Fatalf(`All keys expected to be empty strings`)
	}
}

// TestNewWithEnviron calls New with providing environment variables
func TestNewWithEnviron(t *testing.T) {
	Setup()
	defer TearDown()
	f, err := New()
	if err != nil {
		t.Fatalf(`Unexpected error: %v`, err)
	}
	if (f.AccessKey != testKey) || (f.ServiceKey != testKey) || (f.UserKey != testKey) {
		t.Fatalf(`All keys expected to be %q`, testKey)
	}
}
