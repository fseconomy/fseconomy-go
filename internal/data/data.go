package data

import (
	"github.com/fseconomy/fseconomy-go/exceptions"
	"github.com/fseconomy/fseconomy-go/internal/key"
	"os"
)

// Data is the internal data provider structure
type Data struct {
	serviceKey string
	userKey    string
	accessKey  string
}

// ServiceKey returns the service key
func (d *Data) ServiceKey() string {
	return d.serviceKey
}

// SetServiceKey sets the service key
func (d *Data) SetServiceKey(serviceKey string) error {
	if key.ValidateKey(serviceKey) {
		d.serviceKey = serviceKey
		return nil
	}
	return exceptions.BadFseKeyError
}

// UserKey returns the user key
func (d *Data) UserKey() string {
	return d.userKey
}

// SetUserKey sets the user key
func (d *Data) SetUserKey(userKey string) error {
	if key.ValidateKey(userKey) {
		d.userKey = userKey
		return nil
	}
	return exceptions.BadFseKeyError
}

// AccessKey returns the access key
func (d *Data) AccessKey() string {
	return d.accessKey
}

// SetAccessKey sets the access key
func (d *Data) SetAccessKey(accessKey string) error {
	if key.ValidateKey(accessKey) {
		d.accessKey = accessKey
		return nil
	}
	return exceptions.BadFseKeyError
}

// SetKeysFromEnv sets keys from environment variables
func (d *Data) SetKeysFromEnv() error {
	val, ok := os.LookupEnv("FSE_SERVICE_KEY")
	if ok {
		err := d.SetServiceKey(val)
		if err != nil {
			return err
		}
	}
	val, ok = os.LookupEnv("FSE_USER_KEY")
	if ok {
		err := d.SetUserKey(val)
		if err != nil {
			return err
		}
	}
	val, ok = os.LookupEnv("FSE_ACCESS_KEY")
	if ok {
		err := d.SetAccessKey(val)
		if err != nil {
			return err
		}
	}
	return nil
}

// Keys provides the keys required to query an FSEconomy data feed
//
// This method uses a set of rules to deliver a valid result even if not all
// keys are set:
//
//   - If available, the service key is preferred over the (personal) user key.
//   - If no access key has been defined, the user key is used as default.
//
// The function must be able to establish an authentication key (user or
// service key), and a data access key (read access or user key). In case of
// a failure, it raises either an FseDataKeyError or an FseAuthKeyError.
func (d *Data) Keys() (map[string]string, error) {
	result := make(map[string]string)
	if key.ValidateKey(d.AccessKey()) {
		result["readaccesskey"] = d.AccessKey()
	} else if key.ValidateKey(d.userKey) {
		result["readaccesskey"] = d.UserKey()
	} else {
		return nil, exceptions.FseDataKeyError
	}
	if key.ValidateKey(d.serviceKey) {
		result["servicekey"] = d.ServiceKey()
	} else if key.ValidateKey(d.userKey) {
		result["userkey"] = d.UserKey()
	} else {
		return nil, exceptions.FseAuthKeyError
	}
	return result, nil
}
