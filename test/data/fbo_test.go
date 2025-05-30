package data

import (
	"errors"
	"github.com/fseconomy/fseconomy-go/data"
	"github.com/fseconomy/fseconomy-go/exceptions"
	"testing"
)

func TestFacilitiesByKey(t *testing.T) {
	d, err := data.New()
	if err != nil {
		t.Errorf("unexpected error creating data context: %v", err)
	}
	_, err = d.FacilitiesByKey()
	if err != nil && !errors.Is(err, exceptions.FseDataKeyError) && !errors.Is(err, exceptions.ServerMaintenanceError) {
		t.Errorf("unexpected error getting facilities by key: %v", err)
	}
}

func TestFbosByKey(t *testing.T) {
	d, err := data.New()
	if err != nil {
		t.Errorf("unexpected error creating data context: %v", err)
	}
	_, err = d.FbosByKey()
	if err != nil && !errors.Is(err, exceptions.FseDataKeyError) && !errors.Is(err, exceptions.ServerMaintenanceError) {
		t.Errorf("unexpected error getting FBOs by key: %v", err)
	}
}

func TestFbosForSale(t *testing.T) {
	d, err := data.New()
	if err != nil {
		t.Errorf("unexpected error creating data context: %v", err)
	}
	_, err = d.FbosForSale()
	if err != nil && !errors.Is(err, exceptions.FseDataKeyError) && !errors.Is(err, exceptions.ServerMaintenanceError) {
		t.Errorf("unexpected error getting FBOs for sale: %v", err)
	}
}

func TestFboMonthlySummaryByIcao(t *testing.T) {
	d, err := data.New()
	if err != nil {
		t.Errorf("unexpected error creating data context: %v", err)
	}
	_, err = d.FboMonthlySummaryByIcao(1, 2025, "EDHS")
	if err != nil && !errors.Is(err, exceptions.FseDataKeyError) && !errors.Is(err, exceptions.ServerMaintenanceError) {
		t.Errorf("unexpected error getting FBO monthly summary by ICAO: %v", err)
	}
}
