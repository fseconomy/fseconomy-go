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
	"Aircraft By Key": {
		Service: api.Service{Api: FeedApi, Method: "GET"},
		Query:   "aircraft",
		Search:  "key",
	},
	"Assignments By Key": {
		Service: api.Service{Api: FeedApi, Method: "GET"},
		Query:   "assignments",
		Search:  "key",
	},
	"Commodities By Key": {
		Service: api.Service{Api: FeedApi, Method: "GET"},
		Query:   "commodities",
		Search:  "key",
	},
	"Facilities By Key": {
		Service: api.Service{Api: FeedApi, Method: "GET"},
		Query:   "facilities",
		Search:  "key",
	},
	"FBOs By Key": {
		Service: api.Service{Api: FeedApi, Method: "GET"},
		Query:   "fbos",
		Search:  "key",
	},
	"FBOs For Sale": {
		Service: api.Service{Api: FeedApi, Method: "GET"},
		Query:   "fbos",
		Search:  "forsale",
	},
	"FBO Monthly Summary By ICAO": {
		Service: api.Service{Api: FeedApi, Method: "GET"},
		Query:   "fbos",
		Search:  "monthlysummary",
		Params:  []string{"month", "year", "icao"},
	},
	"Flight Logs By Key Month Year": {
		Service: api.Service{Api: FeedApi, Method: "GET"},
		Query:   "flightlogs",
		Search:  "monthyear",
		Params:  []string{"month", "year"},
	},
	"Flight Logs By Reg Month Year": {
		Service: api.Service{Api: FeedApi, Method: "GET"},
		Query:   "flightlogs",
		Search:  "monthyear",
		Params:  []string{"aircraftreg", "month", "year"},
	},
	"Flight Logs By serialnumber Month Year": {
		Service: api.Service{Api: FeedApi, Method: "GET"},
		Query:   "flightlogs",
		Search:  "monthyear",
		Params:  []string{"serialnumber", "month", "year"},
	},
	"Flight Logs By Key From Id": {
		Service: api.Service{Api: FeedApi, Method: "GET"},
		Query:   "flightlogs",
		Search:  "id",
		Params:  []string{"fromid"},
	},
	"Flight Logs By Key From Id for ALL group aircraft": {
		Service: api.Service{Api: FeedApi, Method: "GET"},
		Query:   "flightlogs",
		Search:  "id",
		Params:  []string{"fromid", "type"},
	},
	"Flight Logs By Reg From Id": {
		Service: api.Service{Api: FeedApi, Method: "GET"},
		Query:   "flightlogs",
		Search:  "id",
		Params:  []string{"aircraftreg", "fromid"},
	},
	"Flight Logs By serialnumber From Id": {
		Service: api.Service{Api: FeedApi, Method: "GET"},
		Query:   "flightlogs",
		Search:  "id",
		Params:  []string{"serialnumber", "fromid"},
	},
	"Group Members": {
		Service: api.Service{Api: FeedApi, Method: "GET"},
		Query:   "group",
		Search:  "members",
	},
	"ICAO Listing of Aircraft": {
		Service: api.Service{Api: FeedApi, Method: "GET"},
		Query:   "icao",
		Search:  "aircraft",
		Params:  []string{"icao"},
	},
	"ICAO Listing of FBOs": {
		Service: api.Service{Api: FeedApi, Method: "GET"},
		Query:   "icao",
		Search:  "fbo",
		Params:  []string{"icao"},
	},
	"ICAO Jobs To": {
		Service: api.Service{Api: FeedApi, Method: "GET"},
		Query:   "icao",
		Search:  "jobsto",
		Params:  []string{"icaos"},
	},
	"ICAO Jobs From": {
		Service: api.Service{Api: FeedApi, Method: "GET"},
		Query:   "icao",
		Search:  "jobsfrom",
		Params:  []string{"icaos"},
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
