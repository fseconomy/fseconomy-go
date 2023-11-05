/*
Package fseconomy is the core FSEconomy SDK package.

# Configuring Credentials

You will require FSEconomy API keys when using the FSEconomy SDK:

* A Service Key is required to use the REST API
* A combination of User Key and Access Key is required to use the data feeds
* Alternatively, the data feeds can also be used with a Service Key instead of a User Key

FSEconomy API keys can be provided by setting these environment variables:

* FSE_SERVICE_KEY => Service Key
* FSE_ACCESS_KEY => Data Access Key
* FSE_USER_KEY => User Key

This is the default way of working:

	f, err := fseconomy.New()
	if err != nil {
		fmt.Fprintf(os.Stderr, "failed to create new fseconomy instance: %v\n", err)
	}
*/
package fseconomy
