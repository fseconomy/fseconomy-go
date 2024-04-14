package security

import (
	"reflect"
	"testing"
)

func TestGetAuthService(t *testing.T) {
	loginService := AuthServices["login"]
	logoutService := AuthServices["logout"]

	type args struct {
		name string
	}
	tests := []struct {
		name    string
		args    args
		want    *AuthService
		wantErr bool
	}{
		{"login service", args{"login"}, &loginService, false},
		{"logout service", args{"logout"}, &logoutService, false},
		{"invalid service", args{"invalid"}, nil, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GetAuthService(tt.args.name)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetAuthService() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetAuthService() got = %v, want %v", got, tt.want)
			}
		})
	}
}
