package security

import (
	"github.com/fseconomy/fseconomy-go/exceptions"
	"github.com/fseconomy/fseconomy-go/internal/api"
)

const AuthApi string = "https://server.fseconomy.net/rest/fse/api"

type AuthService struct {
	api.Service
}

var AuthServices = map[string]AuthService{
	"login": {
		Service: api.Service{
			Api:    AuthApi,
			Url:    "/login",
			Method: "POST",
			Params: []string{"username", "password"},
		},
	},
	"logout": {
		Service: api.Service{
			Api:    AuthApi,
			Url:    "/logout",
			Method: "POST",
			Params: []string{"authtoken", "username"},
		},
	},
}

func GetAuthService(name string) (*AuthService, error) {
	elem, ok := AuthServices[name]
	if ok {
		return &elem, nil
	}
	return nil, exceptions.InvalidServiceError
}
