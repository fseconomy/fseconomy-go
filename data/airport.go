package data

import (
	"bytes"
	"encoding/csv"
	"github.com/fseconomy/fseconomy-go/internal/data"
	"io"
	"strconv"
)

type Airport struct {
	Icao    string
	Lat     float64
	Lon     float64
	Type    string
	Size    int64
	Name    string
	City    string
	State   string
	Country string
}

// Airports extracts data from the Airports data feed
func (d *Data) Airports() ([]*Airport, error) {
	feed, err := data.GetDataFeed("Airports")
	if err != nil {
		return nil, err
	}
	resp, err := feed.QueryCsv()
	if err != nil {
		return nil, err
	}
	reader := csv.NewReader(bytes.NewReader(resp.ByteData))
	reader.TrimLeadingSpace = true
	reader.FieldsPerRecord = -1

	// read header
	if _, err = reader.Read(); err != nil {
		return nil, err
	}

	var items struct {
		Airports []*Airport
	}
	for {
		record, err := reader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			return nil, err
		}
		if len(record) < 9 {
			continue
		}
		lat, _ := strconv.ParseFloat(record[1], 64)
		lon, _ := strconv.ParseFloat(record[2], 64)
		size, _ := strconv.ParseInt(record[4], 10, 64)
		airport := &Airport{
			Icao:    record[0],
			Lat:     lat,
			Lon:     lon,
			Type:    record[3],
			Size:    size,
			Name:    record[5],
			City:    record[6],
			State:   record[7],
			Country: record[8],
		}
		items.Airports = append(items.Airports, airport)
	}
	return items.Airports, nil
}
