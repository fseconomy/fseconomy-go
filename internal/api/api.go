package api

const DataApi string = "https://server.fseconomy.net/data"
const RestApi string = "https://server.fseconomy.net/rest/api"
const RestApiV2 string = "https://server.fseconomy.net/rest/api/v2"

type Service struct {
	Api    string
	Url    string
	Method string
	Params []string
}

func (service *Service) ValidateParams(params map[string]string) bool {
	for _, param := range service.Params {
		if _, ok := params[param]; !ok {
			return false
		}
	}
	return true
}
