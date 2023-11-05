package fseconomy

import (
	"github.com/fseconomy/fseconomy-go/internal/key"
	"os"
)

// Fse is the top level FSEconomy service context
type Fse struct {
	ServiceKey string
	AccessKey  string
	UserKey    string
}

// SetServiceKey sets the service key for a Fse instance
func SetServiceKey(context *Fse, ServiceKey string) error {
	if key.ValidateKey(ServiceKey) {
		context.ServiceKey = ServiceKey
		return nil
	}
	return ErrBadKey
}

// SetUserKey sets the user key for a Fse instance
func SetUserKey(context *Fse, UserKey string) error {
	if key.ValidateKey(UserKey) {
		context.UserKey = UserKey
		return nil
	}
	return ErrBadKey
}

// SetAccessKey sets the read access key for a Fse instance
func SetAccessKey(context *Fse, AccessKey string) error {
	if key.ValidateKey(AccessKey) {
		context.AccessKey = AccessKey
		return nil
	}
	return ErrBadKey
}

// New returns a pointer to a new Fse instance
func New() (*Fse, error) {
	context := &Fse{
		ServiceKey: "",
		AccessKey:  "",
		UserKey:    "",
	}

	val, ok := os.LookupEnv("FSE_SERVICE_KEY")
	if ok {
		err := SetServiceKey(context, val)
		if err != nil {
			return nil, err
		}
	}

	val, ok = os.LookupEnv("FSE_USER_KEY")
	if ok {
		err := SetUserKey(context, val)
		if err != nil {
			return nil, err
		}
	}

	val, ok = os.LookupEnv("FSE_ACCESS_KEY")
	if ok {
		err := SetAccessKey(context, val)
		if err != nil {
			return nil, err
		}
	}

	return context, nil
}
