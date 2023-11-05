package auth

import (
	"github.com/fseconomy/fseconomy-go/fseconomy"
	"github.com/fseconomy/fseconomy-go/internal/api"
)

func Login(context *fseconomy.Fse, username string, password string) {
	_, err := api.GetAuthService("login")
	if err != nil {
		return
	}
}
