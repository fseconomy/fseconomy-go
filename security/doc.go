/*
Package security is the FSEconomy authentication SDK package.

# Configuring Credentials

You will require FSEconomy API keys when using the FSEconomy SDK:

  - A Service Key is required to use the REST API
  - A combination of User Key and Access Key is required to use the data feeds
  - Alternatively, the data feeds can also be used with a Service Key instead of a User Key

FSEconomy API keys can be provided by setting these environment variables:

  - FSE_SERVICE_KEY => Service Key
  - FSE_ACCESS_KEY => Data Access Key
  - FSE_USER_KEY => User Key

This is the default way of working:

	s, err := security.New()
	if err != nil {
		fmt.Fprintf(os.Stderr, "failed to create new fse instance: %v\n", err)
	}

If you need to pass keys explicitly, you can do so when creating the new context:

	a, err := security.New(security.WithServiceKey("service key here"))

Alternatively, keys can be set explicitly using the corresponding setter function:

	s, err := security.New()
	err = s.SetServiceKey("service key string goes here")
	err = s.SetUserKey("user key string goes here")
	err = s.SetAccessKey("access key string goes here")
*/
package security
