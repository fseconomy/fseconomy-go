package api

import "testing"

func TestService_ValidateParams(t *testing.T) {
	service := Service{
		Api:          "SomeApiConstant",
		Url:          "http://localhost",
		Method:       "POST",
		FormParams:   []string{"foo", "bar"},
		HeaderParams: []string{"baz", "quk"},
		UrlParams:    []string{"quz", "flop"},
	}
	tests := []struct {
		name         string
		formParams   map[string]string
		headerParams map[string]string
		urlParams    map[string]string
		want         bool
	}{
		{
			name:         "valid parameters",
			formParams:   map[string]string{"foo": "", "bar": ""},
			headerParams: map[string]string{"baz": "", "quk": ""},
			urlParams:    map[string]string{"quz": "", "flop": ""},
			want:         true,
		},
		{
			name:         "invalid form parameters",
			formParams:   map[string]string{"foo": "", "baz": ""},
			headerParams: map[string]string{"baz": "", "quk": ""},
			urlParams:    map[string]string{"quz": "", "flop": ""},
			want:         false,
		},
		{
			name:         "invalid header parameters",
			formParams:   map[string]string{"foo": "", "bar": ""},
			headerParams: map[string]string{"baz": "", "qek": ""},
			urlParams:    map[string]string{"quz": "", "flop": ""},
			want:         false,
		},
		{
			name:         "invalid url parameters",
			formParams:   map[string]string{"foo": "", "bar": ""},
			headerParams: map[string]string{"baz": "", "quk": ""},
			urlParams:    map[string]string{"quz": "", "flup": ""},
			want:         false,
		},
		{
			name:         "nil parameters",
			formParams:   nil,
			headerParams: nil,
			urlParams:    nil,
			want:         false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if (service.ValidateFormParams(tt.formParams) && service.ValidateHeaderParams(tt.headerParams) && service.ValidateUrlParams(tt.urlParams)) != tt.want {
				t.Error("Expected valid params")
			}
		})
	}
}
