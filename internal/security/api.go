package security

import (
	"github.com/fseconomy/fseconomy-go/exceptions"
	"github.com/fseconomy/fseconomy-go/internal/api"
)

const AuthApi string = "https://server.fseconomy.net/rest/fse/api"

// An AuthService is a service driver for one of FSEconomy's authentication services
type AuthService struct {
	api.Service
}

// AuthServices are service drivers for one of FSEconomy's authentication services
var AuthServices = map[string]AuthService{
	"login": {
		Service: api.Service{
			Api:        AuthApi,
			Url:        "/login",
			Method:     "POST",
			FormParams: []string{"username", "password"},
		},
	},
	"logout": {
		Service: api.Service{
			Api:          AuthApi,
			Url:          "/logout",
			Method:       "POST",
			FormParams:   []string{"username"},
			HeaderParams: []string{"authtoken"},
		},
	},
	"check": {
		Service: api.Service{
			Api:          AuthApi,
			Url:          "/accountcheck",
			Method:       "GET",
			HeaderParams: []string{"authToken"},
		},
	},
}

// GetAuthService returns the requested authentication service (or nil and an error if the
// requested service doesn't exist)
func GetAuthService(name string) (*AuthService, error) {
	elem, ok := AuthServices[name]
	if ok {
		return &elem, nil
	}
	return nil, exceptions.InvalidServiceError
}
