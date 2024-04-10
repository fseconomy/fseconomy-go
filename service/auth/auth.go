package auth

import (
	"github.com/fseconomy/fseconomy-go/fse"
	"github.com/fseconomy/fseconomy-go/internal/api"
)

func Login(context *fse.Fse, username string, password string) {
	_, err := api.GetAuthService("login")
	if err != nil {
		return
	}
}
