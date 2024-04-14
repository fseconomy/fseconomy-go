package data

import (
	"github.com/fseconomy/fseconomy-go/exceptions"
	"github.com/fseconomy/fseconomy-go/internal/api"
	"github.com/fseconomy/fseconomy-go/internal/util"
	"strings"
)

func (d *Feed) QueryFeed(params map[string]string, keys map[string]string) (*api.QueryResult, error) {
	// Check parameter completeness
	for _, param := range d.Params {
		if _, ok := params[param]; !ok {
			return nil, exceptions.MissingParamError
		}
	}

	// Prepare URL parameter set
	urlParams := util.MergeMaps(params, keys)
	urlParams["format"] = "xml"
	urlParams["query"] = d.Query
	urlParams["search"] = d.Search

	// execute query and validate content type
	resp, err := d.RawQuery(urlParams, nil, nil)
	if err != nil {
		return nil, err
	}
	if !strings.Contains(resp.ContentType, "text/xml") {
		return nil, exceptions.InvalidContentTypeError
	}
	if resp.Status != 200 {
		return nil, exceptions.ServerStatusError
	}
	return resp, nil
}
