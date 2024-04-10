package security

import "github.com/fseconomy/fseconomy-go/internal/core"

type FseSec struct {
	core.Fse
}

// WithServiceKey sets the service key
func WithServiceKey(serviceKey string) func(*FseSec) error {
	return func(a *FseSec) error {
		return a.SetServiceKey(serviceKey)
	}
}

// WithAccessKey sets the access key
func WithAccessKey(accessKey string) func(*FseSec) error {
	return func(a *FseSec) error {
		return a.SetAccessKey(accessKey)
	}
}

// WithUserKey sets the user key
func WithUserKey(userKey string) func(*FseSec) error {
	return func(a *FseSec) error {
		return a.SetUserKey(userKey)
	}
}

// New returns a new Fseconomy security object
func New(options ...func(*FseSec) error) (*FseSec, error) {
	auth := &FseSec{}
	err := auth.SetKeysFromEnv()
	if err != nil {
		return nil, err
	}
	for _, option := range options {
		if err := option(auth); err != nil {
			return nil, err
		}
	}
	return auth, nil
}
