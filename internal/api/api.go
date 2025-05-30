package api

const RestApi string = "https://server.fseconomy.net/rest/api"
const RestApiV2 string = "https://server.fseconomy.net/rest/api/v2"

type Service struct {
	Api          string
	Url          string
	Method       string
	FormParams   []string
	HeaderParams []string
	UrlParams    []string
}

// ValidateFormParams validates if all required form parameters are provided
func (service *Service) ValidateFormParams(params map[string]string) bool {
	for _, param := range service.FormParams {
		if _, ok := params[param]; !ok {
			return false
		}
	}
	return true
}

// ValidateHeaderParams validates if all required header parameters are provided
func (service *Service) ValidateHeaderParams(params map[string]string) bool {
	for _, param := range service.HeaderParams {
		if _, ok := params[param]; !ok {
			return false
		}
	}
	return true
}

// ValidateUrlParams validates if all required url parameters are provided
func (service *Service) ValidateUrlParams(params map[string]string) bool {
	for _, param := range service.UrlParams {
		if _, ok := params[param]; !ok {
			return false
		}
	}
	return true
}
