package data

import (
	"encoding/xml"
	"github.com/fseconomy/fseconomy-go/internal/data"
	"github.com/fseconomy/fseconomy-go/types"
)

type Statistic struct {
	Account         string             `xml:"account,attr"`
	PersonalBalance float64            `xml:"Personal_balance"`
	BankBalance     float64            `xml:"Bank_balance"`
	Flights         int64              `xml:"flights"`
	TotalMiles      int64              `xml:"Total_Miles"`
	TimeFlown       types.FseHobbsTime `xml:"Time_Flown"`
}

// StatisticsByKey extracts data from the Statistics By Key data feed
func (d *Data) StatisticsByKey() ([]*Statistic, error) {
	keys, err := d.Keys()
	if err != nil {
		return nil, err
	}
	feed, err := data.GetDataFeed("Statistics By Key")
	if err != nil {
		return nil, err
	}
	resp, err := feed.QueryFeed(nil, keys)
	if err != nil {
		return nil, err
	}
	var items struct {
		Statistics []*Statistic `xml:"Statistic"`
	}
	err = xml.Unmarshal(resp.ByteData, &items)
	if err != nil {
		return nil, err
	}
	return items.Statistics, nil
}
