package api

import (
	"bytes"
	"github.com/fseconomy/fseconomy-go/exceptions"
	"io"
	"net/http"
	"net/url"
	"strings"
)

const maintenance = "Currently Closed for Maintenance"

type QueryResult struct {
	status      int
	contentType string
	byteData    []byte
	stringData  string
}

// RawQuery assembles and executes the http request to a service, returning a result object
func (service *Service) RawQuery(urlParams map[string]string, formParams map[string]string, headerParams map[string]string) (*QueryResult, error) {

	// build final url from service data and url parameters
	reqUrl, err := url.Parse(service.Api + service.Url)
	if err != nil {
		return nil, err
	}

	// prepare url parameters
	if urlParams != nil {
		up := url.Values{}
		for k, v := range urlParams {
			up.Set(k, v)
		}
		reqUrl.RawQuery = up.Encode()
	}

	// prepare form parameters
	var body io.Reader
	if formParams != nil {
		fp := url.Values{}
		for k, v := range formParams {
			fp.Add(k, v)
		}
		body = bytes.NewBufferString(fp.Encode())
	}

	// create request
	req, err := http.NewRequest(service.Method, reqUrl.String(), body)
	if err != nil {
		return nil, err
	}

	// set header params
	// Set the Content-Type header for form data
	if formParams != nil {
		req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	}

	if headerParams != nil {
		for k, v := range headerParams {
			req.Header.Set(k, v)
		}
	}

	// create a http client and execute the request
	client := http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	// read data and populate response object
	var result QueryResult
	var readErr error
	result.status = resp.StatusCode
	result.byteData, readErr = io.ReadAll(resp.Body)
	closeErr := resp.Body.Close()
	if readErr != nil {
		return nil, readErr
	}
	if closeErr != nil {
		return nil, closeErr
	}

	// convert string data and check for server maintenance downtime
	result.stringData = string(result.byteData)
	if strings.Contains(result.stringData, maintenance) {
		return nil, exceptions.ServerMaintenanceError
	}

	// make sure we know the content type downstream
	result.contentType = resp.Header.Get("Content-Type")

	return &result, nil
}
