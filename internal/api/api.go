package api

const DataApi string = "https://server.fseconomy.net/data"
const RestApi string = "https://server.fseconomy.net/rest/api"
const RestApiV2 string = "https://server.fseconomy.net/rest/api/v2"

type Service struct {
	api    string
	url    string
	method string
	params []string
}
