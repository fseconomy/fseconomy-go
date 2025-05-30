package data

import (
	"encoding/xml"
	"github.com/fseconomy/fseconomy-go/internal/data"
)

type GroupMember struct {
	Name   string `xml:"Name"`
	Status string `xml:"Status"`
}

// GroupMembers extracts data from the Group Members data feed
func (d *Data) GroupMembers() ([]*GroupMember, error) {
	keys, err := d.Keys()
	if err != nil {
		return nil, err
	}
	feed, err := data.GetDataFeed("Group Members")
	if err != nil {
		return nil, err
	}
	resp, err := feed.QueryFeed(nil, keys)
	if err != nil {
		return nil, err
	}
	var items struct {
		GroupMember []*GroupMember `xml:"FlightLog"`
	}
	err = xml.Unmarshal(resp.ByteData, &items)
	if err != nil {
		return nil, err
	}
	return items.GroupMember, nil
}
