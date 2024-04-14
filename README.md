# FSEconomy Go Bindings

![Build and Test State](https://github.com/fseconomy/fseconomy-go/actions/workflows/go.yml/badge.svg?event=push)
[![Go Reference](https://pkg.go.dev/badge/github.com/fseconomy/fseconomy-go.svg)](https://pkg.go.dev/github.com/fseconomy/fseconomy-go)
![MIT License](https://img.shields.io/badge/license-MIT-green.svg)

This package provides [Go](https://go.dev/) bindings for various [FSEconomy](https://www.fseconomy.net) APIs.

**Important Links**

* Documentation: https://pkg.go.dev/github.com/fseconomy/fseconomy-go
* GitHub: https://github.com/fseconomy/fseconomy-go
* Bug Tracker: https://github.com/fseconomy/fseconomy-go/issues

## Installing

Use `go get` to retrieve the SDK and add it to your `GOPATH` workspace, or project's Go module dependencies:

```shell
go get github.com/fseconomy/fseconomy-go
```

To update the SDK, use `go get -U` to retrieve the latest version of the SDK:

```shell
go get -U github.com/fseconomy/fseconomy-go
```

## Usage

### Data Feed API

```go
package main

import (
	"github.com/fseconomy/fseconomy-go/data"
	"log"
)

func main() {
	d, err := data.New(data.WithUserKey("my user key here"))
	if err != nil {
		log.Fatalf(`failed to create FSEconomy data feed context: %v\n`, err)
	}
	
	// retrieve data feeds
	configs, err := d.AircraftConfigs()
	status, err := d.AircraftStatusByRegistration("D-ESTE")
	
	log.Printf("loaded number of configs: %d", len(configs))
	log.Printf("status for D-ESTE: %v", status)
}
```

### Security & Authentication API

```go
package main

import (
	"github.com/fseconomy/fseconomy-go/security"
	"log"
)

func main() {
	s, err := security.New()
	if err != nil {
		log.Fatalf(`failed to create FSEconomy Security context: %v\n`, err)
	}
	
	// login with user and password
	err = s.Login("user", "password")
	
	// check if currently logged in
	err = s.Check()
	
	// logout
	err = s.Logout()
}
```