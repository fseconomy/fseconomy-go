package api

import (
	"testing"
)

func TestService_RawQuery(t *testing.T) {
	type fields struct {
		Api          string
		Url          string
		Method       string
		FormParams   []string
		UrlParams    []string
		HeaderParams []string
	}
	type args struct {
		urlParams    map[string]string
		formParams   map[string]string
		headerParams map[string]string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: "cat fact without length",
			fields: fields{
				Api:    "https://catfact.ninja",
				Url:    "/fact",
				Method: "GET",
			},
			args: args{
				urlParams:    nil,
				formParams:   nil,
				headerParams: nil,
			},
			wantErr: false,
		},
		{
			name: "cat fact with length",
			fields: fields{
				Api:       "https://catfact.ninja",
				Url:       "/fact",
				Method:    "GET",
				UrlParams: []string{"max_length"},
			},
			args: args{
				urlParams:    map[string]string{"max_length": "50"},
				formParams:   nil,
				headerParams: nil,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			service := &Service{
				Api:          tt.fields.Api,
				Url:          tt.fields.Url,
				Method:       tt.fields.Method,
				FormParams:   tt.fields.FormParams,
				HeaderParams: tt.fields.HeaderParams,
				UrlParams:    tt.fields.UrlParams,
			}
			got, err := service.RawQuery(tt.args.urlParams, tt.args.formParams, tt.args.headerParams)
			if (err != nil) != tt.wantErr {
				t.Errorf("RawQuery() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got.Status != 200 {
				t.Errorf("RawQuery() got status = %d, want %d", got.Status, 200)
			}
		})
	}
}
