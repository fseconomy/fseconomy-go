package fseconomy

import (
	"github.com/fseconomy/fseconomy-go/internal/apikey"
)

// Fse is the top level FSEconomy service context
type Fse struct {
	ServiceKey string
	AccessKey  string
	UserKey    string
}

// New returns a pointer to a new Fse instance
func New() *Fse {
	context := &Fse{
		ServiceKey: "",
		AccessKey:  "",
		UserKey:    "",
	}
	return context
}

// setServiceKey sets the service key for a Fse instance
func setServiceKey(context *Fse, ServiceKey string) error {
	if apikey.ValidateKey(ServiceKey) {
		context.ServiceKey = ServiceKey
		return nil
	}
	return ErrBadKey
}

// setUserKey sets the user key for a Fse instance
func setUserKey(context *Fse, UserKey string) error {
	if apikey.ValidateKey(UserKey) {
		context.UserKey = UserKey
		return nil
	}
	return ErrBadKey
}

// setAccessKey sets the read access key for a Fse instance
func setAccessKey(context *Fse, AccessKey string) error {
	if apikey.ValidateKey(AccessKey) {
		context.AccessKey = AccessKey
		return nil
	}
	return ErrBadKey
}
