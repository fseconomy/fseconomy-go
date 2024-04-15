package data

import (
	"encoding/xml"
	"github.com/fseconomy/fseconomy-go/internal/data"
	"strconv"
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

type AircraftAlias struct {
	MakeModel string   `xml:"MakeModel"`
	Alias     []string `xml:"Alias"`
}

type Aircraft struct {
	SerialNumber  int64   `xml:"SerialNumber"`
	MakeModel     string  `xml:"MakeModel"`
	Registration  string  `xml:"Registration"`
	Owner         string  `xml:"Owner"`
	Location      string  `xml:"Location"`
	LocationName  string  `xml:"LocationName"`
	Home          string  `xml:"Home"`
	SalePrice     float64 `xml:"SalePrice"`
	SellbackPrice float64 `xml:"SellbackPrice"`
	Equipment     string  `xml:"Equipment"`
	RentalDry     float64 `xml:"RentalDry"`
	RentalWet     float64 `xml:"RentalWet"`
	RentalType    string  `xml:"RentalType"`
	Bonus         int64   `xml:"Bonus"`
	RentalTime    int64   `xml:"RentalTime"`
	RentedBy      string  `xml:"RentedBy"`
	FuelPct       float64 `xml:"FuelPct"`
	NeedsRepair   bool    `xml:"NeedsRepair"`
	AirframeTime  string  `xml:"AirframeTime"`
	EngineTime    string  `xml:"EngineTime"`
	TimeLast100hr string  `xml:"TimeLast100hr"`
	LeasedFrom    string  `xml:"LeasedFrom"`
	MonthlyFee    float64 `xml:"MonthlyFee"`
	FeeOwed       float64 `xml:"FeeOwed"`
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

// AircraftAliases extracts data from the Aircraft Aliases data feed
func (d *Data) AircraftAliases() ([]*AircraftAlias, error) {
	keys, err := d.Keys()
	if err != nil {
		return nil, err
	}
	feed, err := data.GetDataFeed("Aircraft Aliases")
	if err != nil {
		return nil, err
	}
	resp, err := feed.QueryFeed(nil, keys)
	if err != nil {
		return nil, err
	}
	var items struct {
		Alias []*AircraftAlias `xml:"AircraftAliases"`
	}
	err = xml.Unmarshal(resp.ByteData, &items)
	if err != nil {
		return nil, err
	}
	return items.Alias, nil
}

// AircraftForSale extracts data from the Aircraft For Sale data feed
func (d *Data) AircraftForSale() ([]*Aircraft, error) {
	keys, err := d.Keys()
	if err != nil {
		return nil, err
	}
	feed, err := data.GetDataFeed("Aircraft For Sale")
	if err != nil {
		return nil, err
	}
	resp, err := feed.QueryFeed(nil, keys)
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

// AircraftByMakeModel extracts data from the Aircraft By MakeModel data feed
func (d *Data) AircraftByMakeModel(makeModel string) ([]*Aircraft, error) {
	keys, err := d.Keys()
	if err != nil {
		return nil, err
	}
	feed, err := data.GetDataFeed("Aircraft By MakeModel")
	if err != nil {
		return nil, err
	}
	resp, err := feed.QueryFeed(map[string]string{"makemodel": makeModel}, keys)
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

// AircraftByOwnerName extracts data from the Aircraft By Owner Name data feed
func (d *Data) AircraftByOwnerName(ownerName string) ([]*Aircraft, error) {
	keys, err := d.Keys()
	if err != nil {
		return nil, err
	}
	feed, err := data.GetDataFeed("Aircraft By Owner Name")
	if err != nil {
		return nil, err
	}
	resp, err := feed.QueryFeed(map[string]string{"ownername": ownerName}, keys)
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

// AircraftByRegistration extracts data from the Aircraft By Registration data feed
func (d *Data) AircraftByRegistration(registration string) (*Aircraft, error) {
	keys, err := d.Keys()
	if err != nil {
		return nil, err
	}
	feed, err := data.GetDataFeed("Aircraft By Registration")
	if err != nil {
		return nil, err
	}
	resp, err := feed.QueryFeed(map[string]string{"aircraftreg": registration}, keys)
	if err != nil {
		return nil, err
	}
	var items struct {
		Aircraft *Aircraft `xml:"Aircraft"`
	}
	err = xml.Unmarshal(resp.ByteData, &items)
	if err != nil {
		return nil, err
	}
	return items.Aircraft, nil
}

// AircraftById extracts data from the Aircraft By Id data feed
func (d *Data) AircraftById(serialNumber int) (*Aircraft, error) {
	keys, err := d.Keys()
	if err != nil {
		return nil, err
	}
	feed, err := data.GetDataFeed("Aircraft By Id")
	if err != nil {
		return nil, err
	}
	resp, err := feed.QueryFeed(map[string]string{"serialnumber": strconv.Itoa(serialNumber)}, keys)
	if err != nil {
		return nil, err
	}
	var items struct {
		Aircraft *Aircraft `xml:"Aircraft"`
	}
	err = xml.Unmarshal(resp.ByteData, &items)
	if err != nil {
		return nil, err
	}
	return items.Aircraft, nil
}
