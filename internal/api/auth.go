package api

import (
	"github.com/fseconomy/fseconomy-go/exceptions"
)

const AuthApi string = "https://server.fseconomy.net/rest/fse/api"

var AuthServices = map[string]Service{
	"login": {
		api:    AuthApi,
		url:    "/login",
		method: "POST",
		params: []string{"username", "password"},
	},
	"logout": {
		api:    AuthApi,
		url:    "/logout",
		method: "POST",
		params: []string{"authtoken", "username"},
	},
}

func GetAuthService(name string) (*Service, error) {
	elem, ok := AuthServices[name]
	if ok {
		return &elem, nil
	}
	return nil, exceptions.InvalidServiceError
}
