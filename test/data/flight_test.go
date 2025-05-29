package data

import (
	"errors"
	"github.com/fseconomy/fseconomy-go/data"
	"github.com/fseconomy/fseconomy-go/exceptions"
	"testing"
)

func TestFlightLogsByKeyMonthYear(t *testing.T) {
	d, err := data.New()
	if err != nil {
		t.Errorf("unexpected error creating data context: %v", err)
	}
	_, err = d.FlightLogsByKeyMonthYear(1, 2025)
	if err != nil && !errors.Is(err, exceptions.FseDataKeyError) && !errors.Is(err, exceptions.ServerMaintenanceError) {
		t.Errorf("unexpected error getting facilities by key: %v", err)
	}
}

func TestFlightLogsByRegMonthYear(t *testing.T) {
	d, err := data.New()
	if err != nil {
		t.Errorf("unexpected error creating data context: %v", err)
	}
	_, err = d.FlightLogsByRegMonthYear("D-ESTE", 1, 2025)
	if err != nil && !errors.Is(err, exceptions.FseDataKeyError) && !errors.Is(err, exceptions.ServerMaintenanceError) {
		t.Errorf("unexpected error getting facilities by key: %v", err)
	}
}

func TestFlightLogsBySerialnumberMonthYear(t *testing.T) {
	d, err := data.New()
	if err != nil {
		t.Errorf("unexpected error creating data context: %v", err)
	}
	_, err = d.FlightLogsBySerialnumberMonthYear(32019, 1, 2025)
	if err != nil && !errors.Is(err, exceptions.FseDataKeyError) && !errors.Is(err, exceptions.ServerMaintenanceError) {
		t.Errorf("unexpected error getting facilities by key: %v", err)
	}
}

func TestFlightLogsByKeyFromID(t *testing.T) {
	d, err := data.New()
	if err != nil {
		t.Errorf("unexpected error creating data context: %v", err)
	}
	_, err = d.FlightLogsByKeyFromId(0)
	if err != nil && !errors.Is(err, exceptions.FseDataKeyError) && !errors.Is(err, exceptions.ServerMaintenanceError) {
		t.Errorf("unexpected error getting facilities by key: %v", err)
	}
}

func TestFlightLogsByKeyFromIdForAllGroupAircraft(t *testing.T) {
	d, err := data.New()
	if err != nil {
		t.Errorf("unexpected error creating data context: %v", err)
	}
	_, err = d.FlightLogsByKeyFromIdForAllGroupAircraft(0)
	if err != nil && !errors.Is(err, exceptions.FseDataKeyError) && !errors.Is(err, exceptions.ServerMaintenanceError) {
		t.Errorf("unexpected error getting facilities by key: %v", err)
	}
}

func TestFlightLogsByRegFromId(t *testing.T) {
	d, err := data.New()
	if err != nil {
		t.Errorf("unexpected error creating data context: %v", err)
	}
	_, err = d.FlightLogsByRegFromId("D-ESTE", 0)
	if err != nil && !errors.Is(err, exceptions.FseDataKeyError) && !errors.Is(err, exceptions.ServerMaintenanceError) {
		t.Errorf("unexpected error getting facilities by key: %v", err)
	}
}

func TestFlightLogsBySerialnumberFromId(t *testing.T) {
	d, err := data.New()
	if err != nil {
		t.Errorf("unexpected error creating data context: %v", err)
	}
	_, err = d.FlightLogsBySerialnumberFromId(32019, 0)
	if err != nil && !errors.Is(err, exceptions.FseDataKeyError) && !errors.Is(err, exceptions.ServerMaintenanceError) {
		t.Errorf("unexpected error getting facilities by key: %v", err)
	}
}
