package data

import (
	"errors"
	"github.com/fseconomy/fseconomy-go/data"
	"github.com/fseconomy/fseconomy-go/exceptions"
	"testing"
)

func TestAirports(t *testing.T) {
	d, err := data.New()
	if err != nil {
		t.Errorf("unexpected error creating data context: %v", err)
	}
	airports, err := d.Airports()
	if err != nil && !errors.Is(err, exceptions.FseDataKeyError) && !errors.Is(err, exceptions.ServerMaintenanceError) {
		t.Errorf("unexpected error getting statistics by key: %v", err)
	}
	if len(airports) < 23780 {
		t.Errorf("not enough airports, expected at least 23780, got: %d", len(airports))
	}
}
