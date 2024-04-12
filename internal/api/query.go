package api

import (
	"bytes"
	"io"
	"net/http"
	"net/url"
)

func (service *Service) RawQuery(urlParams map[string]string, formParams map[string]string, headerParams map[string]string) ([]byte, error) {
	// prepare url parameters
	up := url.Values{}
	for k, v := range urlParams {
		up.Set(k, v)
	}

	// prepare form parameters
	fp := url.Values{}
	for k, v := range formParams {
		fp.Add(k, v)
	}

	// build final url from service data and url parameters
	reqUrl, err := url.Parse(service.Api + service.Url)
	if err != nil {
		return nil, err
	}
	reqUrl.RawQuery = up.Encode()

	// create request
	req, err := http.NewRequest(service.Method, reqUrl.String(), bytes.NewBufferString(fp.Encode()))
	if err != nil {
		return nil, err
	}

	// set header params
	// Set the Content-Type header for form data
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	for k, v := range headerParams {
		req.Header.Set(k, v)
	}

	// create a http client and execute the request
	client := http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	rawData, readErr := io.ReadAll(resp.Body)
	closeErr := resp.Body.Close()
	if readErr != nil {
		return nil, readErr
	}
	if closeErr != nil {
		return nil, closeErr
	}
	return rawData, nil
}
