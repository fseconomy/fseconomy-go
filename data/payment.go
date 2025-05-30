package data

import (
	"encoding/xml"
	"github.com/fseconomy/fseconomy-go/internal/data"
	"github.com/fseconomy/fseconomy-go/types"
	"strconv"
)

type Payment struct {
	Id         int64                  `xml:"Id"`
	Date       types.FseFlightLogTime `xml:"Date"`
	To         string                 `xml:"To"`
	From       string                 `xml:"From"`
	Amount     float64                `xml:"Amount"`
	Reason     string                 `xml:"Reason"`
	Fbo        types.FseString        `xml:"Fbo"`
	Location   types.FseString        `xml:"Location"`
	Aircraft   types.FseString        `xml:"Aircraft"`
	Comment    types.FseString        `xml:"Comment"`
	AircraftId int64                  `xml:"AircraftId"`
}

// PaymentsByMonthYear extracts data from the Payments By Month Year data feed
func (d *Data) PaymentsByMonthYear(month int, year int, archive bool) ([]*Payment, error) {
	keys, err := d.Keys()
	if err != nil {
		return nil, err
	}
	feed, err := data.GetDataFeed("Payments By Month Year")
	if err != nil {
		return nil, err
	}
	params := make(map[string]string)
	params["month"] = strconv.Itoa(month)
	params["year"] = strconv.Itoa(year)
	if archive {
		params["archive"] = "true"
	}
	resp, err := feed.QueryFeed(params, keys)
	if err != nil {
		return nil, err
	}
	var items struct {
		Payment []*Payment `xml:"Payment"`
	}
	err = xml.Unmarshal(resp.ByteData, &items)
	if err != nil {
		return nil, err
	}
	return items.Payment, nil
}

// PaymentsFromId extracts data from the Payments From Id data feed
func (d *Data) PaymentsFromId(id int64, archive bool) ([]*Payment, error) {
	keys, err := d.Keys()
	if err != nil {
		return nil, err
	}
	feed, err := data.GetDataFeed("Payments From Id")
	if err != nil {
		return nil, err
	}
	params := make(map[string]string)
	params["id"] = strconv.FormatInt(id, 10)
	if archive {
		params["archive"] = "true"
	}
	resp, err := feed.QueryFeed(params, keys)
	if err != nil {
		return nil, err
	}
	var items struct {
		Payment []*Payment `xml:"Payment"`
	}
	err = xml.Unmarshal(resp.ByteData, &items)
	if err != nil {
		return nil, err
	}
	return items.Payment, nil
}
