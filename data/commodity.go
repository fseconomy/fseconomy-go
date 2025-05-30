package data

import (
	"encoding/xml"
	"github.com/fseconomy/fseconomy-go/internal/data"
	"github.com/fseconomy/fseconomy-go/types"
)

type Commodity struct {
	Location string                   `xml:"Location"`
	Type     string                   `xml:"Type"`
	Amount   types.FseCommodityAmount `xml:"Amount"`
}

// CommoditiesByKey extracts data from the Assignments By Key data feed
func (d *Data) CommoditiesByKey() ([]*Commodity, error) {
	keys, err := d.Keys()
	if err != nil {
		return nil, err
	}
	feed, err := data.GetDataFeed("Commodities By Key")
	if err != nil {
		return nil, err
	}
	resp, err := feed.QueryFeed(nil, keys)
	if err != nil {
		return nil, err
	}
	var items struct {
		Commodity []*Commodity `xml:"Commodity"`
	}
	err = xml.Unmarshal(resp.ByteData, &items)
	if err != nil {
		return nil, err
	}
	return items.Commodity, nil
}
