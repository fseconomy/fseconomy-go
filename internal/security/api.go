package security

import (
	"github.com/fseconomy/fseconomy-go/exceptions"
	"github.com/fseconomy/fseconomy-go/internal/api"
)

const AuthApi string = "https://server.fseconomy.net/rest/fse/api"

var AuthServices = map[string]api.Service{
	"login": {
		Api:    AuthApi,
		Url:    "/login",
		Method: "POST",
		Params: []string{"username", "password"},
	},
	"logout": {
		Api:    AuthApi,
		Url:    "/logout",
		Method: "POST",
		Params: []string{"authtoken", "username"},
	},
}

func GetAuthService(name string) (*api.Service, error) {
	elem, ok := AuthServices[name]
	if ok {
		return &elem, nil
	}
	return nil, exceptions.InvalidServiceError
}
