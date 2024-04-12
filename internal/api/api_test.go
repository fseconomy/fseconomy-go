package api

import "testing"

func TestService_ValidateParams(t *testing.T) {
	service := Service{
		Api:    "SomeApiConstant",
		Url:    "http://localhost",
		Method: "POST",
		Params: []string{"foo", "bar"},
	}
	tests := []struct {
		name   string
		params map[string]string
		want   bool
	}{
		{"valid parameters", map[string]string{"foo": "", "bar": ""}, true},
		{"invalid parameters", map[string]string{"foo": "", "baz": ""}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if service.ValidateParams(tt.params) != tt.want {
				t.Error("Expected valid params")
			}
		})
	}
}
