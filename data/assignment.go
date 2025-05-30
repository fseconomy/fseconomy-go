package data

import (
	"encoding/xml"
	"github.com/fseconomy/fseconomy-go/internal/data"
	"github.com/fseconomy/fseconomy-go/types"
)

type Assignment struct {
	XMLName        xml.Name                `xml:"Assignment"`
	Id             int64                   `xml:"Id"`
	Status         string                  `xml:"Status"`
	Location       string                  `xml:"Location"`
	From           string                  `xml:"From"`
	Destination    string                  `xml:"Destination"`
	AssignmentText string                  `xml:"Assignment"`
	Amount         int64                   `xml:"Amount"`
	Units          string                  `xml:"Units"`
	Pay            float64                 `xml:"Pay"`
	PilotFee       float64                 `xml:"PilotFee"`
	Expires        string                  `xml:"Expires"`
	ExpireDateTime types.FseAssignmentTime `xml:"ExpireDateTime"`
	Type           string                  `xml:"Type"`
	Express        bool                    `xml:"Express"`
	PtAssignment   bool                    `xml:"PtAssignment"`
	Locked         string                  `xml:"Locked"`
	Comment        string                  `xml:"Comment"`
}

// AssignmentsByKey extracts data from the Assignments By Key data feed
func (d *Data) AssignmentsByKey() ([]*Assignment, error) {
	keys, err := d.Keys()
	if err != nil {
		return nil, err
	}
	feed, err := data.GetDataFeed("Assignments By Key")
	if err != nil {
		return nil, err
	}
	resp, err := feed.QueryFeed(nil, keys)
	if err != nil {
		return nil, err
	}
	var items struct {
		Assignment []*Assignment `xml:"Assignment"`
	}
	err = xml.Unmarshal(resp.ByteData, &items)
	if err != nil {
		return nil, err
	}
	return items.Assignment, nil
}
