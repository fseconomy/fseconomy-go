package util

import (
	"reflect"
	"testing"
)

func TestMergeMaps(t *testing.T) {
	type args[K comparable, V any] struct {
		maps []map[K]V
	}
	type testCase[K comparable, V any] struct {
		name string
		args args[K, V]
		want map[K]V
	}
	tests := []testCase[string, string]{
		{
			name: "map with different keys",
			args: args[string, string]{
				maps: []map[string]string{
					{"foo": "bar"},
					{"boo": "far"},
				},
			},
			want: map[string]string{"foo": "bar", "boo": "far"},
		},
		{
			name: "map with same keys",
			args: args[string, string]{
				maps: []map[string]string{
					{"foo": "bar"},
					{"foo": "far"},
				},
			},
			want: map[string]string{"foo": "far"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := MergeMaps(tt.args.maps...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("mergeMaps() = %v, want %v", got, tt.want)
			}
		})
	}
}
