package api

import (
	"github.com/fseconomy/fseconomy-go/internal/core"
	"net/http"
	"net/url"
)

func Query(context *core.Fse, service *Service) {
	path, err := url.JoinPath(service.api, service.url)
	if err != nil {
		return
	}
	req, err := http.NewRequest(service.method, path, nil)
	client := http.Client{}
	_, _ = client.Do(req)
}
