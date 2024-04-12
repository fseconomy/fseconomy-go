package core

import (
	"github.com/fseconomy/fseconomy-go/exceptions"
	"github.com/fseconomy/fseconomy-go/internal/key"
	"os"
)

// Fse is the top level FSEconomy service context
type Fse struct {
	serviceKey string
	accessKey  string
	userKey    string
	userName   string
	userToken  string
}

// SetServiceKey sets the service key for a Fse instance
func (context *Fse) SetServiceKey(serviceKey string) error {
	if key.ValidateKey(serviceKey) {
		context.serviceKey = serviceKey
		return nil
	}
	return exceptions.BadFseKeyError
}

// SetUserKey sets the user key for a Fse instance
func (context *Fse) SetUserKey(userKey string) error {
	if key.ValidateKey(userKey) {
		context.userKey = userKey
		return nil
	}
	return exceptions.BadFseKeyError
}

// SetAccessKey sets the read access key for a Fse instance
func (context *Fse) SetAccessKey(accessKey string) error {
	if key.ValidateKey(accessKey) {
		context.accessKey = accessKey
		return nil
	}
	return exceptions.BadFseKeyError
}

func (context *Fse) SetKeysFromEnv() error {
	val, ok := os.LookupEnv("FSE_SERVICE_KEY")
	if ok {
		err := context.SetServiceKey(val)
		if err != nil {
			return err
		}
	}
	val, ok = os.LookupEnv("FSE_USER_KEY")
	if ok {
		err := context.SetUserKey(val)
		if err != nil {
			return err
		}
	}
	val, ok = os.LookupEnv("FSE_ACCESS_KEY")
	if ok {
		err := context.SetAccessKey(val)
		if err != nil {
			return err
		}
	}
	return nil
}

// New returns a pointer to a new Fse instance
func New() (*Fse, error) {
	context := &Fse{
		serviceKey: "",
		accessKey:  "",
		userKey:    "",
		userName:   "",
		userToken:  "",
	}

	err := context.SetKeysFromEnv()
	if err != nil {
		return nil, err
	}

	return context, nil
}
