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

// TestSetAccessKey calls SetAccessKey for an empty context and verifies the key has been set correctly
func TestSetAccessKey(t *testing.T) {
	f, err := New()
	if err != nil {
		t.Fatalf(`Unexpected error: %v`, err)
	}
	err = SetAccessKey(f, testKey)
	if err != nil {
		t.Fatalf(`Unexpected error: %v`, err)
	}
	if f.AccessKey != testKey {
		t.Fatalf(`Access key expected to be %q, got %q instead`, testKey, f.AccessKey)
	}
}

// TestSetServiceKey calls SetServiceKey for an empty context and verifies the key has been set correctly
func TestSetServiceKey(t *testing.T) {
	f, err := New()
	if err != nil {
		t.Fatalf(`Unexpected error: %v`, err)
	}
	err = SetServiceKey(f, testKey)
	if err != nil {
		t.Fatalf(`Unexpected error: %v`, err)
	}
	if f.ServiceKey != testKey {
		t.Fatalf(`Service key expected to be %q, got %q instead`, testKey, f.ServiceKey)
	}
}

// TestSetUserKey calls SetUserKey for an empty context and verifies the key has been set correctly
func TestSetUserKey(t *testing.T) {
	f, err := New()
	if err != nil {
		t.Fatalf(`Unexpected error: %v`, err)
	}
	err = SetUserKey(f, testKey)
	if err != nil {
		t.Fatalf(`Unexpected error: %v`, err)
	}
	if f.UserKey != testKey {
		t.Fatalf(`User key expected to be %q, got %q instead`, testKey, f.UserKey)
	}
}
