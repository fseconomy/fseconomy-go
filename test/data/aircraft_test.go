package data_test

import (
	"errors"
	"github.com/fseconomy/fseconomy-go/data"
	"github.com/fseconomy/fseconomy-go/exceptions"
	"testing"
)

func TestAircraftStatusByRegistration(t *testing.T) {
	d, err := data.New()
	if err != nil {
		t.Errorf("unexpected error creating data context: %v", err)
	}
	_, err = d.AircraftStatusByRegistration("D-ESTE")
	if err != nil && !errors.Is(err, exceptions.FseDataKeyError) {
		t.Errorf("unexpected error getting aircraft status: %v", err)
	}
}

func TestAircraftConfigs(t *testing.T) {
	d, err := data.New(data.WithUserKey("0015D19A457E2FCD"))
	if err != nil {
		t.Errorf("unexpected error creating data context: %v", err)
	}
	_, err = d.AircraftConfigs()
	if err != nil && !errors.Is(err, exceptions.FseDataKeyError) {
		t.Errorf("unexpected error getting aircraft configs: %v", err)
	}
}
