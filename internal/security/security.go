package security

import (
	"encoding/hex"
	"github.com/fseconomy/fseconomy-go/exceptions"
	"os"
)

// Security is the internal security structure
type Security struct {
	secretKey []byte
	secretIv  []byte
	userName  string
	authToken string
}

// UserName returns the username
func (s *Security) UserName() string {
	return s.userName
}

// SetUserName sets the username
func (s *Security) SetUserName(userName string) {
	s.userName = userName
}

// AuthToken returns the FSEconomy authentication token
func (s *Security) AuthToken() string {
	return s.authToken
}

// SetAuthToken sets the FSEconomy authentication token
func (s *Security) SetAuthToken(token string) {
	s.authToken = token
}

// SetSecretKey sets the secret key (16 bytes hex string expected)
func (s *Security) SetSecretKey(key []byte) error {
	if len(key) != 16 {
		return exceptions.BadSecretKeyError
	}
	s.secretKey = key
	return nil
}

// SetSecretKeyFromString sets secret key if submitted value is a hex string representing 16 bytes of data
func (s *Security) SetSecretKeyFromString(key string) error {
	val, err := hex.DecodeString(key)
	if err != nil {
		return exceptions.BadSecretKeyError
	}
	err = s.SetSecretKey(val)
	if err != nil {
		return err
	}
	return nil
}

// SetSecretKeyFromEnv checks for the environment variable FSE_SECRET_KEY and
// sets its content as secret key (if it is a hex string representing 16 bytes of data)
func (s *Security) SetSecretKeyFromEnv() error {
	val, ok := os.LookupEnv("FSE_SECRET_KEY")
	if ok {
		err := s.SetSecretKeyFromString(val)
		if err != nil {
			return err
		}
	}
	return nil
}

// SetSecretIv sets the initialization vector (12 bytes hex string expected)
func (s *Security) SetSecretIv(iv []byte) error {
	if len(iv) != 12 {
		return exceptions.BadSecretIvError
	}
	s.secretIv = iv
	return nil
}

// SetSecretIvFromString sets secret key if submitted value is a hex string representing 12 bytes of data
func (s *Security) SetSecretIvFromString(key string) error {
	val, err := hex.DecodeString(key)
	if err != nil {
		return exceptions.BadSecretIvError
	}
	err = s.SetSecretIv(val)
	if err != nil {
		return err
	}
	return nil
}

// SetSecretIvFromEnv checks for the environment variable FSE_SECRET_IV and
// sets its content as initialization vector (if it is a hex string representing 12 bytes of data)
func (s *Security) SetSecretIvFromEnv() error {
	val, ok := os.LookupEnv("FSE_SECRET_IV")
	if ok {

		err := s.SetSecretIvFromString(val)
		if err != nil {
			return err
		}
	}
	return nil
}
