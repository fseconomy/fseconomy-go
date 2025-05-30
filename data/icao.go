package data

import (
	"encoding/xml"
	"fmt"
	"github.com/fseconomy/fseconomy-go/internal/data"
	"github.com/fseconomy/fseconomy-go/types"
	"strings"
)

type Job struct {
	Id             int                     `xml:"Id"`
	Location       string                  `xml:"Location"`
	ToIcao         string                  `xml:"ToIcao"`
	FromIcao       string                  `xml:"FromIcao"`
	Amount         int64                   `xml:"Amount"`
	UnitType       string                  `xml:"UnitType"`
	Commodity      string                  `xml:"Commodity"`
	Pay            float64                 `xml:"Pay"`
	Expires        string                  `xml:"Expires"`
	ExpireDateTime types.FseAssignmentTime `xml:"ExpireDateTime"`
	Express        bool                    `xml:"Express"`
	PtAssignment   bool                    `xml:"PtAssignment"`
	Type           string                  `xml:"Type"`
	AircraftId     int64                   `xml:"AircraftId"`
}

// IcaoListingOfAircraft extracts data from the ICAO Listing of Aircraft data feed
func (d *Data) IcaoListingOfAircraft(icao string) ([]*Aircraft, error) {
	keys, err := d.Keys()
	if err != nil {
		return nil, err
	}
	feed, err := data.GetDataFeed("ICAO Listing of Aircraft")
	if err != nil {
		return nil, err
	}
	resp, err := feed.QueryFeed(map[string]string{"icao": icao}, keys)
	if err != nil {
		return nil, err
	}
	var items struct {
		Aircraft []*Aircraft `xml:"Aircraft"`
	}
	err = xml.Unmarshal(resp.ByteData, &items)
	if err != nil {
		return nil, err
	}
	return items.Aircraft, nil
}

// IcaoListingOfFbos extracts data from the ICAO Listing of FBOs data feed
func (d *Data) IcaoListingOfFbos(icao string) ([]*Fbo, error) {
	keys, err := d.Keys()
	if err != nil {
		return nil, err
	}
	feed, err := data.GetDataFeed("ICAO Listing of FBOs")
	if err != nil {
		return nil, err
	}
	resp, err := feed.QueryFeed(map[string]string{"icao": icao}, keys)
	if err != nil {
		return nil, err
	}
	var items struct {
		Fbo []*Fbo `xml:"FBO"`
	}
	err = xml.Unmarshal(resp.ByteData, &items)
	if err != nil {
		return nil, err
	}
	return items.Fbo, nil
}

// IcaoJobsTo extracts data from the ICAO Jobs To data feed
func (d *Data) IcaoJobsTo(icaos []string) ([]*Job, error) {
	keys, err := d.Keys()
	if err != nil {
		return nil, err
	}
	icaoList := strings.Join(icaos, "-")
	if icaoList == "" {
		return nil, fmt.Errorf("requires at least one ICAO code")
	}
	feed, err := data.GetDataFeed("ICAO Jobs To")
	if err != nil {
		return nil, err
	}
	resp, err := feed.QueryFeed(map[string]string{"icaos": icaoList}, keys)
	if err != nil {
		return nil, err
	}
	var items struct {
		Job []*Job `xml:"Assignment"`
	}
	err = xml.Unmarshal(resp.ByteData, &items)
	if err != nil {
		return nil, err
	}
	return items.Job, nil
}

// IcaoJobsFrom extracts data from the ICAO Jobs From data feed
func (d *Data) IcaoJobsFrom(icaos []string) ([]*Job, error) {
	keys, err := d.Keys()
	if err != nil {
		return nil, err
	}
	icaoList := strings.Join(icaos, "-")
	if icaoList == "" {
		return nil, fmt.Errorf("requires at least one ICAO code")
	}
	feed, err := data.GetDataFeed("ICAO Jobs From")
	if err != nil {
		return nil, err
	}
	resp, err := feed.QueryFeed(map[string]string{"icaos": icaoList}, keys)
	if err != nil {
		return nil, err
	}
	var items struct {
		Job []*Job `xml:"Assignment"`
	}
	err = xml.Unmarshal(resp.ByteData, &items)
	if err != nil {
		return nil, err
	}
	return items.Job, nil
}
