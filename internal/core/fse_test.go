package core

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
	if (f.accessKey != "") || (f.serviceKey != "") || (f.userKey != "") {
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
	if (f.accessKey != testKey) || (f.serviceKey != testKey) || (f.userKey != testKey) {
		t.Fatalf(`All keys expected to be %q`, testKey)
	}
}

// TestSetAccessKey calls SetAccessKey for an empty context and verifies the key has been set correctly
func TestSetAccessKey(t *testing.T) {
	f, err := New()
	if err != nil {
		t.Fatalf(`Unexpected error: %v`, err)
	}
	err = f.SetAccessKey(testKey)
	if err != nil {
		t.Fatalf(`Unexpected error: %v`, err)
	}
	if f.accessKey != testKey {
		t.Fatalf(`Access key expected to be %q, got %q instead`, testKey, f.accessKey)
	}
}

// TestSetServiceKey calls SetServiceKey for an empty context and verifies the key has been set correctly
func TestSetServiceKey(t *testing.T) {
	f, err := New()
	if err != nil {
		t.Fatalf(`Unexpected error: %v`, err)
	}
	err = f.SetServiceKey(testKey)
	if err != nil {
		t.Fatalf(`Unexpected error: %v`, err)
	}
	if f.serviceKey != testKey {
		t.Fatalf(`Service key expected to be %q, got %q instead`, testKey, f.serviceKey)
	}
}

// TestSetUserKey calls SetUserKey for an empty context and verifies the key has been set correctly
func TestSetUserKey(t *testing.T) {
	f, err := New()
	if err != nil {
		t.Fatalf(`Unexpected error: %v`, err)
	}
	err = f.SetUserKey(testKey)
	if err != nil {
		t.Fatalf(`Unexpected error: %v`, err)
	}
	if f.userKey != testKey {
		t.Fatalf(`User key expected to be %q, got %q instead`, testKey, f.userKey)
	}
}
