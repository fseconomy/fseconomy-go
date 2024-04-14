package data

import (
	"encoding/xml"
	"github.com/fseconomy/fseconomy-go/internal/data"
)

// AircraftStatus represents the data provided by the Aircraft Status data feed
type AircraftStatus struct {
	SerialNumber int    `xml:"SerialNumber"`
	Status       string `xml:"Status"`
	Location     string `xml:"Location"`
}

type AircraftConfig struct {
	MakeModel   string  `xml:"MakeModel"`
	Crew        int64   `xml:"Crew"`
	Seats       int64   `xml:"Seats"`
	CruiseSpeed int64   `xml:"CruiseSpeed"`
	GPH         int64   `xml:"GPH"`
	FuelType    int64   `xml:"FuelType"`
	MTOW        int64   `xml:"MTOW"`
	EmptyWeight int64   `xml:"EmptyWeight"`
	Price       float64 `xml:"Price"`
	Ext1        int64   `xml:"Ext1"`
	LTip        int64   `xml:"LTip"`
	LAux        int64   `xml:"LAux"`
	LMain       int64   `xml:"LMain"`
	Center1     int64   `xml:"Center1"`
	Center2     int64   `xml:"Center2"`
	Center3     int64   `xml:"Center3"`
	RMain       int64   `xml:"RMain"`
	RAux        int64   `xml:"RAux"`
	RTip        int64   `xml:"RTip"`
	Ext2        int64   `xml:"Ext2"`
	Engines     int64   `xml:"Engines"`
	EnginePrice float64 `xml:"EnginePrice"`
	ModelId     int64   `xml:"ModelId"`
	MaxCargo    int64   `xml:"MaxCargo"`
}

// AircraftStatusByRegistration extracts data from the Aircraft Status By Registration data feed
func (d *Data) AircraftStatusByRegistration(registration string) (*AircraftStatus, error) {
	keys, err := d.Keys()
	if err != nil {
		return nil, err
	}
	feed, err := data.GetDataFeed("Aircraft Status By Registration")
	if err != nil {
		return nil, err
	}
	resp, err := feed.QueryFeed(map[string]string{"aircraftreg": registration}, keys)
	if err != nil {
		return nil, err
	}
	var status struct {
		Status *AircraftStatus `xml:"Aircraft"`
	}
	err = xml.Unmarshal(resp.ByteData, &status)
	if err != nil {
		return nil, err
	}
	return status.Status, nil
}

// AircraftConfigs extracts data from the Aircraft Configs data feed
func (d *Data) AircraftConfigs() ([]*AircraftConfig, error) {
	keys, err := d.Keys()
	if err != nil {
		return nil, err
	}
	feed, err := data.GetDataFeed("Aircraft Configs")
	if err != nil {
		return nil, err
	}
	resp, err := feed.QueryFeed(nil, keys)
	if err != nil {
		return nil, err
	}
	var items struct {
		Config []*AircraftConfig `xml:"AircraftConfig"`
	}
	err = xml.Unmarshal(resp.ByteData, &items)
	if err != nil {
		return nil, err
	}
	return items.Config, nil
}
