package security_test

import (
	"github.com/fseconomy/fseconomy-go/security"
	"os"
	"testing"
)

const (
	testKeyGood     string = "869c7ede03e6435e3d13a984b3dcdf91"
	testKeyBadShort string = "193f47bba353a619b5f15f1fea70"
	testKeyBadInval string = "i amm no hex string"
	testIvGood      string = "21688c2b3998ce52fd3b2179"
	testIvBadShort  string = "dff833d28cf07a956b58"
	testIvBadInval  string = "also no hex string"
)

func SetupGood() {
	_ = os.Setenv("FSE_SECRET_KEY", testKeyGood)
	_ = os.Setenv("FSE_SECRET_IV", testIvGood)
}

func SetupBadShort() {
	_ = os.Setenv("FSE_SECRET_KEY", testKeyBadShort)
	_ = os.Setenv("FSE_SECRET_IV", testIvBadShort)
}

func SetupBadInval() {
	_ = os.Setenv("FSE_SECRET_KEY", testKeyBadInval)
	_ = os.Setenv("FSE_SECRET_IV", testIvBadInval)
}

func TearDown() {
	_ = os.Unsetenv("FSE_SECRET_KEY")
	_ = os.Unsetenv("FSE_SECRET_IV")
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

func TestNewEnvironBadShort(t *testing.T) {
	SetupBadShort()
	defer TearDown()
	_, err := security.New()
	if err == nil {
		t.Errorf("Expected invalid secret key error, got: %v", err)
	}
}

func TestNewEnvironBadInval(t *testing.T) {
	SetupBadInval()
	defer TearDown()
	_, err := security.New()
	if err == nil {
		t.Errorf("Expected invalid secret key error, got: %v", err)
	}
}

func TestNewWithSecretKey(t *testing.T) {
	tests := []struct {
		name    string
		key     string
		wantErr bool
	}{
		{"Good key", testKeyGood, false},
		{"Short key", testKeyBadShort, true},
		{"Inval key", testKeyBadInval, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := security.New(security.WithSecretKey(tt.key))
			if (err != nil) != tt.wantErr {
				t.Errorf("New() with key error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}

func TestNewWithSecretIv(t *testing.T) {
	tests := []struct {
		name    string
		iv      string
		wantErr bool
	}{
		{"Good IV", testIvGood, false},
		{"Short IV", testIvBadShort, true},
		{"Inval IV", testIvBadInval, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := security.New(security.WithSecretIv(tt.iv))
			if (err != nil) != tt.wantErr {
				t.Errorf("New() with IV error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}

func TestNewWithUser(t *testing.T) {
	s, err := security.New(security.WithUser("foo"))
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	if s.UserName() != "foo" {
		t.Errorf("Unexpected user name: %s (want: 'foo')", s.UserName())
	}
}

func TestNewWithToken(t *testing.T) {
	s, err := security.New(security.WithToken("abc"))
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	if s.AuthToken() != "abc" {
		t.Errorf("Unexpected token: %s (want: 'abc')", s.AuthToken())
	}
}
