package data

import "github.com/fseconomy/fseconomy-go/internal/data"

type Data struct {
	data.Data
}

// WithServiceKey sets the service key
func WithServiceKey(key string) func(*Data) error {
	return func(d *Data) error {
		return d.SetServiceKey(key)
	}
}

// WithAccessKey sets the access key
func WithAccessKey(key string) func(*Data) error {
	return func(d *Data) error {
		return d.SetAccessKey(key)
	}
}

// WithUserKey sets the user key
func WithUserKey(key string) func(*Data) error {
	return func(d *Data) error {
		return d.SetUserKey(key)
	}
}

// New returns a new FSEconomy data object
func New(options ...func(*Data) error) (*Data, error) {
	d := &Data{}
	err := d.SetKeysFromEnv()
	if err != nil {
		return nil, err
	}
	for _, option := range options {
		if err := option(d); err != nil {
			return nil, err
		}
	}
	return d, nil
}
