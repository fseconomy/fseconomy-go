package data

import (
	"os"
	"reflect"
	"testing"
)

const testKey string = "2ad0394ad6c8e400"

func Setup() {
	_ = os.Setenv("FSE_SERVICE_KEY", testKey)
	_ = os.Setenv("FSE_USER_KEY", testKey)
	_ = os.Setenv("FSE_ACCESS_KEY", testKey)
}

func TearDown() {
	_ = os.Unsetenv("FSE_SERVICE_KEY")
	_ = os.Unsetenv("FSE_USER_KEY")
	_ = os.Unsetenv("FSE_ACCESS_KEY")
}

func TestData_AccessKey(t *testing.T) {
	type fields struct {
		serviceKey string
		userKey    string
		accessKey  string
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{"valid access key", fields{accessKey: "2ad0394ad6c8e400"}, "2ad0394ad6c8e400"},
		{"no access key", fields{}, ""},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := &Data{
				serviceKey: tt.fields.serviceKey,
				userKey:    tt.fields.userKey,
				accessKey:  tt.fields.accessKey,
			}
			if got := d.AccessKey(); got != tt.want {
				t.Errorf("AccessKey() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestData_Keys(t *testing.T) {
	type fields struct {
		serviceKey string
		userKey    string
		accessKey  string
	}
	tests := []struct {
		name    string
		fields  fields
		want    map[string]string
		wantErr bool
	}{
		{
			name:    "valid keys with service key and access key",
			fields:  fields{serviceKey: "2ad0394ad6c8e400", accessKey: "2bd0394ad6c8e400"},
			want:    map[string]string{"readaccesskey": "2bd0394ad6c8e400", "servicekey": "2ad0394ad6c8e400"},
			wantErr: false,
		},
		{
			name:    "valid keys with user key and access key",
			fields:  fields{userKey: "2ad0394ad6c8e400", accessKey: "2bd0394ad6c8e400"},
			want:    map[string]string{"readaccesskey": "2bd0394ad6c8e400", "userkey": "2ad0394ad6c8e400"},
			wantErr: false,
		},
		{
			name:    "valid keys with user key",
			fields:  fields{userKey: "2ad0394ad6c8e400"},
			want:    map[string]string{"readaccesskey": "2ad0394ad6c8e400", "userkey": "2ad0394ad6c8e400"},
			wantErr: false,
		},
		{
			name:    "valid keys with user key and service key",
			fields:  fields{userKey: "2ad0394ad6c8e400", serviceKey: "2bd0394ad6c8e400"},
			want:    map[string]string{"readaccesskey": "2ad0394ad6c8e400", "servicekey": "2bd0394ad6c8e400"},
			wantErr: false,
		},
		{
			name:    "invalid keys with service key only",
			fields:  fields{serviceKey: "2bd0394ad6c8e400"},
			want:    nil,
			wantErr: true,
		},
		{
			name:    "invalid keys with access key only",
			fields:  fields{accessKey: "2ad0394ad6c8e400"},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := &Data{
				serviceKey: tt.fields.serviceKey,
				userKey:    tt.fields.userKey,
				accessKey:  tt.fields.accessKey,
			}
			got, err := d.Keys()
			if (err != nil) != tt.wantErr {
				t.Errorf("Keys() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Keys() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestData_ServiceKey(t *testing.T) {
	type fields struct {
		serviceKey string
		userKey    string
		accessKey  string
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{"valid service key", fields{serviceKey: "2ad0394ad6c8e400"}, "2ad0394ad6c8e400"},
		{"no service key", fields{}, ""},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := &Data{
				serviceKey: tt.fields.serviceKey,
				userKey:    tt.fields.userKey,
				accessKey:  tt.fields.accessKey,
			}
			if got := d.ServiceKey(); got != tt.want {
				t.Errorf("ServiceKey() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestData_SetAccessKey(t *testing.T) {
	type fields struct {
		serviceKey string
		userKey    string
		accessKey  string
	}
	type args struct {
		accessKey string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{"valid access key", fields{}, args{accessKey: "2ad0394ad6c8e400"}, false},
		{"invalid access key", fields{}, args{accessKey: "2ad0394ad6c8e40"}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := &Data{
				serviceKey: tt.fields.serviceKey,
				userKey:    tt.fields.userKey,
				accessKey:  tt.fields.accessKey,
			}
			if err := d.SetAccessKey(tt.args.accessKey); (err != nil) != tt.wantErr {
				t.Errorf("SetAccessKey() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestData_SetKeysFromEnv(t *testing.T) {
	d := &Data{}
	Setup()
	defer TearDown()
	err := d.SetKeysFromEnv()
	if err != nil {
		t.Errorf("SetKeysFromEnv() error = %v", err)
	}
	if (d.AccessKey() != testKey) || (d.ServiceKey() != testKey) || (d.UserKey() != testKey) {
		t.Errorf("All keys expected to be %s, got: AccessKey %s, ServiceKey %s, UserKey %s", testKey, d.AccessKey(), d.ServiceKey(), d.UserKey())
	}
}

func TestData_SetServiceKey(t *testing.T) {
	type fields struct {
		serviceKey string
		userKey    string
		accessKey  string
	}
	type args struct {
		serviceKey string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{"valid service key", fields{}, args{serviceKey: "2ad0394ad6c8e400"}, false},
		{"invalid service key", fields{}, args{serviceKey: "2ad0394ad6c8e40"}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := &Data{
				serviceKey: tt.fields.serviceKey,
				userKey:    tt.fields.userKey,
				accessKey:  tt.fields.accessKey,
			}
			if err := d.SetServiceKey(tt.args.serviceKey); (err != nil) != tt.wantErr {
				t.Errorf("SetServiceKey() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestData_SetUserKey(t *testing.T) {
	type fields struct {
		serviceKey string
		userKey    string
		accessKey  string
	}
	type args struct {
		userKey string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{"valid user key", fields{}, args{userKey: "2ad0394ad6c8e400"}, false},
		{"invalid user key", fields{}, args{userKey: "2ad0394ad6c8e40"}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := &Data{
				serviceKey: tt.fields.serviceKey,
				userKey:    tt.fields.userKey,
				accessKey:  tt.fields.accessKey,
			}
			if err := d.SetUserKey(tt.args.userKey); (err != nil) != tt.wantErr {
				t.Errorf("SetUserKey() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestData_UserKey(t *testing.T) {
	type fields struct {
		serviceKey string
		userKey    string
		accessKey  string
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{"valid user key", fields{userKey: "2ad0394ad6c8e400"}, "2ad0394ad6c8e400"},
		{"no user key", fields{}, ""},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := &Data{
				serviceKey: tt.fields.serviceKey,
				userKey:    tt.fields.userKey,
				accessKey:  tt.fields.accessKey,
			}
			if got := d.UserKey(); got != tt.want {
				t.Errorf("UserKey() = %v, want %v", got, tt.want)
			}
		})
	}
}
