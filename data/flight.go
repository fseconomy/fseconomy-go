package data

import (
	"encoding/xml"
	"github.com/fseconomy/fseconomy-go/internal/data"
	"github.com/fseconomy/fseconomy-go/types"
	"strconv"
)

type FlightLog struct {
	Id              int64                  `xml:"Id"`
	Type            string                 `xml:"Type"`
	Time            types.FseFlightLogTime `xml:"Time"`
	Distance        int64                  `xml:"Distance"`
	Pilot           string                 `xml:"Pilot"`
	SerialNumber    int64                  `xml:"SerialNumber"`
	Aircraft        string                 `xml:"Aircraft"`
	MakeModel       string                 `xml:"MakeModel"`
	From            string                 `xml:"From"`
	To              string                 `xml:"To"`
	TotalEngineTime types.FseHobbsTime     `xml:"TotalEngineTime"`
	FlightTime      types.FseHobbsTime     `xml:"FlightTime"`
	GroupName       string                 `xml:"GroupName"`
	Income          float64                `xml:"Income"`
	PilotFee        float64                `xml:"PilotFee"`
	CrewCost        float64                `xml:"CrewCost"`
	BookingFee      float64                `xml:"BookingFee"`
	Bonus           float64                `xml:"Bonus"`
	FuelCost        float64                `xml:"FuelCost"`
	Gcf             float64                `xml:"GCF"`
	RentalPrice     float64                `xml:"RentalPrice"`
	RentalType      string                 `xml:"RentalType"`
	RentalUnits     types.FseHobbsTime     `xml:"RentalUnits"`
	RentalCost      float64                `xml:"RentalCost"`
}

// FlightLogsByKeyMonthYear extracts data from the Flight Logs By Key Month Year data feed
func (d *Data) FlightLogsByKeyMonthYear(month int, year int) ([]*FlightLog, error) {
	keys, err := d.Keys()
	if err != nil {
		return nil, err
	}
	feed, err := data.GetDataFeed("Flight Logs By Key Month Year")
	if err != nil {
		return nil, err
	}
	resp, err := feed.QueryFeed(map[string]string{"month": strconv.Itoa(month), "year": strconv.Itoa(year)}, keys)
	if err != nil {
		return nil, err
	}
	var items struct {
		FlightLog []*FlightLog `xml:"FlightLog"`
	}
	err = xml.Unmarshal(resp.ByteData, &items)
	if err != nil {
		return nil, err
	}
	return items.FlightLog, nil
}

// FlightLogsByRegMonthYear extracts data from the Flight Logs By Reg Month Year data feed
func (d *Data) FlightLogsByRegMonthYear(registration string, month int, year int) ([]*FlightLog, error) {
	keys, err := d.Keys()
	if err != nil {
		return nil, err
	}
	feed, err := data.GetDataFeed("Flight Logs By Reg Month Year")
	if err != nil {
		return nil, err
	}
	resp, err := feed.QueryFeed(map[string]string{"aircraftreg": registration, "month": strconv.Itoa(month), "year": strconv.Itoa(year)}, keys)
	if err != nil {
		return nil, err
	}
	var items struct {
		FlightLog []*FlightLog `xml:"FlightLog"`
	}
	err = xml.Unmarshal(resp.ByteData, &items)
	if err != nil {
		return nil, err
	}
	return items.FlightLog, nil
}

// FlightLogsBySerialnumberMonthYear extracts data from the Flight Logs By serialnumber Month Year data feed
func (d *Data) FlightLogsBySerialnumberMonthYear(serialNumber int64, month int, year int) ([]*FlightLog, error) {
	keys, err := d.Keys()
	if err != nil {
		return nil, err
	}
	feed, err := data.GetDataFeed("Flight Logs By serialnumber Month Year")
	if err != nil {
		return nil, err
	}
	resp, err := feed.QueryFeed(map[string]string{"serialnumber": strconv.FormatInt(serialNumber, 10), "month": strconv.Itoa(month), "year": strconv.Itoa(year)}, keys)
	if err != nil {
		return nil, err
	}
	var items struct {
		FlightLog []*FlightLog `xml:"FlightLog"`
	}
	err = xml.Unmarshal(resp.ByteData, &items)
	if err != nil {
		return nil, err
	}
	return items.FlightLog, nil
}

// FlightLogsByKeyFromId extracts data from the Flight Logs By Key From Id data feed
func (d *Data) FlightLogsByKeyFromId(id int64) ([]*FlightLog, error) {
	keys, err := d.Keys()
	if err != nil {
		return nil, err
	}
	feed, err := data.GetDataFeed("Flight Logs By Key From Id")
	if err != nil {
		return nil, err
	}
	resp, err := feed.QueryFeed(map[string]string{"fromid": strconv.FormatInt(id, 10)}, keys)
	if err != nil {
		return nil, err
	}
	var items struct {
		FlightLog []*FlightLog `xml:"FlightLog"`
	}
	err = xml.Unmarshal(resp.ByteData, &items)
	if err != nil {
		return nil, err
	}
	return items.FlightLog, nil
}

// FlightLogsByKeyFromIdForAllGroupAircraft extracts data from the Flight Logs By Key From Id for ALL group aircraft data feed
func (d *Data) FlightLogsByKeyFromIdForAllGroupAircraft(id int64) ([]*FlightLog, error) {
	keys, err := d.Keys()
	if err != nil {
		return nil, err
	}
	feed, err := data.GetDataFeed("Flight Logs By Key From Id for ALL group aircraft")
	if err != nil {
		return nil, err
	}
	resp, err := feed.QueryFeed(map[string]string{"fromid": strconv.FormatInt(id, 10), "type": "groupaircraft"}, keys)
	if err != nil {
		return nil, err
	}
	var items struct {
		FlightLog []*FlightLog `xml:"FlightLog"`
	}
	err = xml.Unmarshal(resp.ByteData, &items)
	if err != nil {
		return nil, err
	}
	return items.FlightLog, nil
}

// FlightLogsByRegFromId extracts data from the Flight Logs By Reg From Id data feed
func (d *Data) FlightLogsByRegFromId(registration string, id int64) ([]*FlightLog, error) {
	keys, err := d.Keys()
	if err != nil {
		return nil, err
	}
	feed, err := data.GetDataFeed("Flight Logs By Rey From Id")
	if err != nil {
		return nil, err
	}
	resp, err := feed.QueryFeed(map[string]string{"aircraftreg": registration, "fromid": strconv.FormatInt(id, 10)}, keys)
	if err != nil {
		return nil, err
	}
	var items struct {
		FlightLog []*FlightLog `xml:"FlightLog"`
	}
	err = xml.Unmarshal(resp.ByteData, &items)
	if err != nil {
		return nil, err
	}
	return items.FlightLog, nil
}

// FlightLogsBySerialnumberFromId extracts data from the Flight Logs By serialnumber From Id data feed
func (d *Data) FlightLogsBySerialnumberFromId(serialNumber int64, id int64) ([]*FlightLog, error) {
	keys, err := d.Keys()
	if err != nil {
		return nil, err
	}
	feed, err := data.GetDataFeed("Flight Logs By serialnumber From Id")
	if err != nil {
		return nil, err
	}
	resp, err := feed.QueryFeed(map[string]string{"serialnumber": strconv.FormatInt(serialNumber, 10), "fromid": strconv.FormatInt(id, 10)}, keys)
	if err != nil {
		return nil, err
	}
	var items struct {
		FlightLog []*FlightLog `xml:"FlightLog"`
	}
	err = xml.Unmarshal(resp.ByteData, &items)
	if err != nil {
		return nil, err
	}
	return items.FlightLog, nil
}
