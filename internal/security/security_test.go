package security

import (
	"errors"
	"github.com/fseconomy/fseconomy-go/exceptions"
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

func TestSecurity_SetSecretIv(t *testing.T) {
	tests := []struct {
		name string
		iv   string
		want error
	}{
		{"good", testIvGood, nil},
		{"short", testIvBadShort, exceptions.BadSecretIvError},
		{"inval", testIvBadInval, exceptions.BadSecretIvError},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Security{}
			got := s.SetSecretIvFromString(tt.iv)
			if !errors.Is(got, tt.want) {
				t.Errorf("SetSecretIvFromString() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSecurity_SetSecretKey(t *testing.T) {
	tests := []struct {
		name string
		iv   string
		want error
	}{
		{"good", testKeyGood, nil},
		{"short", testKeyBadShort, exceptions.BadSecretKeyError},
		{"inval", testKeyBadInval, exceptions.BadSecretKeyError},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Security{}
			got := s.SetSecretKeyFromString(tt.iv)
			if !errors.Is(got, tt.want) {
				t.Errorf("SetSecretKeyFromString() got = %v, want %v", got, tt.want)
			}
		})
	}
}
