package data

import (
	"errors"
	"github.com/fseconomy/fseconomy-go/data"
	"github.com/fseconomy/fseconomy-go/exceptions"
	"testing"
)

func TestGroupMembers(t *testing.T) {
	d, err := data.New()
	if err != nil {
		t.Errorf("unexpected error creating data context: %v", err)
	}
	_, err = d.GroupMembers()
	if err != nil && !errors.Is(err, exceptions.FseDataKeyError) && !errors.Is(err, exceptions.ServerMaintenanceError) {
		t.Errorf("unexpected error getting group members: %v", err)
	}
}
