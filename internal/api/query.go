package api

import (
	"github.com/fseconomy/fseconomy-go/internal/core"
	"net/http"
	"net/url"
)

func Query(context *core.Fse, service *Service) {
	path, err := url.JoinPath(service.Api, service.Url)
	if err != nil {
		return
	}
	req, err := http.NewRequest(service.Method, path, nil)
	client := http.Client{}
	_, _ = client.Do(req)
}
