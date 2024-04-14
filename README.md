# FSEconomy Go Bindings

![Build and Test State](https://github.com/fseconomy/fseconomy-go/actions/workflows/go.yml/badge.svg?event=push)

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