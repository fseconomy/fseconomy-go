package security

import (
	"encoding/json"
	"github.com/fseconomy/fseconomy-go/exceptions"
	"log"
	"strings"
)

type Response struct {
	Meta struct {
		Code  int    `json:"code"`
		Error string `json:"error"`
		Info  string `json:"info"`
	} `json:"meta"`
	Data string `json:"data"`
}

func (auth *AuthService) Query(fp map[string]string, hp map[string]string, up map[string]string) (*Response, error) {
	resp, err := auth.RawQuery(up, fp, hp)
	if err != nil {
		return nil, err
	}
	if !strings.Contains(resp.ContentType, "application/json") {
		log.Printf("Content Type: %s", resp.ContentType)
		return nil, exceptions.InvalidContentTypeError
	}
	var response Response
	err = json.Unmarshal(resp.ByteData, &response)
	if err != nil {
		return nil, err
	}
	return &response, nil
}
