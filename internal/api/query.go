package api

import (
	"github.com/fseconomy/fseconomy-go/fseconomy"
	"net/http"
	"net/url"
)

func Query(context *fseconomy.Fse, service *Service) {
	path, err := url.JoinPath(service.api, service.url)
	if err != nil {
		return
	}
	req, err := http.NewRequest(service.method, path, nil)
	client := http.Client{}
	resp, err := client.Do(req)
}