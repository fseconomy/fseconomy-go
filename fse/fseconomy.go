package fse

import (
	"github.com/fseconomy/fseconomy-go/internal/key"
	"os"
)

// Fse is the top level FSEconomy service context
type Fse struct {
	ServiceKey string
	AccessKey  string
	UserKey    string
	UserName   string
	UserToken  string
}

// An Option is a functional option for configuring a Fse context
type Option func(*Fse)

// WithServiceKey sets the service key
func WithServiceKey(serviceKey string) Option {
	return func(fse *Fse) {
		fse.ServiceKey = serviceKey
	}
}

// WithAccessKey sets the access key
func WithAccessKey(accessKey string) Option {
	return func(fse *Fse) {
		fse.AccessKey = accessKey
	}
}

// WithUserKey sets the user key
func WithUserKey(userKey string) Option {
	return func(fse *Fse) {
		fse.UserKey = userKey
	}
}

// SetServiceKey sets the service key for a Fse instance
func (context *Fse) SetServiceKey(serviceKey string) error {
	if key.ValidateKey(serviceKey) {
		context.ServiceKey = serviceKey
		return nil
	}
	return ErrBadKey
}

// SetUserKey sets the user key for a Fse instance
func (context *Fse) SetUserKey(userKey string) error {
	if key.ValidateKey(userKey) {
		context.UserKey = userKey
		return nil
	}
	return ErrBadKey
}

// SetAccessKey sets the read access key for a Fse instance
func (context *Fse) SetAccessKey(accessKey string) error {
	if key.ValidateKey(accessKey) {
		context.AccessKey = accessKey
		return nil
	}
	return ErrBadKey
}

// New returns a pointer to a new Fse instance
func New(options ...Option) (*Fse, error) {
	context := &Fse{
		ServiceKey: "",
		AccessKey:  "",
		UserKey:    "",
		UserName:   "",
		UserToken:  "",
	}

	val, ok := os.LookupEnv("FSE_SERVICE_KEY")
	if ok {
		err := context.SetServiceKey(val)
		if err != nil {
			return nil, err
		}
	}

	val, ok = os.LookupEnv("FSE_USER_KEY")
	if ok {
		err := context.SetUserKey(val)
		if err != nil {
			return nil, err
		}
	}

	val, ok = os.LookupEnv("FSE_ACCESS_KEY")
	if ok {
		err := context.SetAccessKey(val)
		if err != nil {
			return nil, err
		}
	}

	for _, option := range options {
		option(context)
	}

	return context, nil
}
