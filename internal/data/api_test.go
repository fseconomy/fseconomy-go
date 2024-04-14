package data

import (
	"reflect"
	"testing"
)

func TestGetDataFeed(t *testing.T) {
	acStatusByReg := Feeds["Aircraft Status By Registration"]

	type args struct {
		name string
	}
	tests := []struct {
		name    string
		args    args
		want    *Feed
		wantErr bool
	}{
		{"valid data feed", args{"Aircraft Status By Registration"}, &acStatusByReg, false},
		{"invalid data feed", args{"Those Magnificent Men in Their Flying Machines"}, nil, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GetDataFeed(tt.args.name)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetDataFeed() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetDataFeed() got = %v, want %v", got, tt.want)
			}
		})
	}
}
