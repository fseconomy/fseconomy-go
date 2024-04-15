package data

import (
	"github.com/fseconomy/fseconomy-go/exceptions"
	"github.com/fseconomy/fseconomy-go/internal/api"
)

const FeedApi string = "https://server.fseconomy.net/data"

// A Feed is a service driver for one of FSEconomy's data feeds
type Feed struct {
	api.Service
	Query  string
	Search string
	Params []string
}

// Feeds are service drivers for one of FSEconomy's data feeds
var Feeds = map[string]Feed{
	"Aircraft Status By Registration": {
		Service: api.Service{Api: FeedApi, Method: "GET"},
		Query:   "aircraft",
		Search:  "status",
		Params:  []string{"aircraftreg"},
	},
	"Aircraft Configs": {
		Service: api.Service{Api: FeedApi, Method: "GET"},
		Query:   "aircraft",
		Search:  "configs",
	},
	"Aircraft Aliases": {
		Service: api.Service{Api: FeedApi, Method: "GET"},
		Query:   "aircraft",
		Search:  "aliases",
	},
	"Aircraft For Sale": {
		Service: api.Service{Api: FeedApi, Method: "GET"},
		Query:   "aircraft",
		Search:  "forsale",
	},
	"Aircraft By MakeModel": {
		Service: api.Service{Api: FeedApi, Method: "GET"},
		Query:   "aircraft",
		Search:  "makemodel",
		Params:  []string{"makemodel"},
	},
	"Aircraft By Owner Name": {
		Service: api.Service{Api: FeedApi, Method: "GET"},
		Query:   "aircraft",
		Search:  "ownername",
		Params:  []string{"ownername"},
	},
	"Aircraft By Registration": {
		Service: api.Service{Api: FeedApi, Method: "GET"},
		Query:   "aircraft",
		Search:  "registration",
		Params:  []string{"aircraftreg"},
	},
	"Aircraft By Id": {
		Service: api.Service{Api: FeedApi, Method: "GET"},
		Query:   "aircraft",
		Search:  "serialnumber",
		Params:  []string{"serialnumber"},
	},
}

// GetDataFeed returns the requested data feed (or nil and an error if the
// requested service doesn't exist)
func GetDataFeed(name string) (*Feed, error) {
	elem, ok := Feeds[name]
	if ok {
		return &elem, nil
	}
	return nil, exceptions.InvalidServiceError
}
