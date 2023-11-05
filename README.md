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

### Set Up the FSEconomy SDK

```go
package main

import (
	"fmt"
	"os"
	
	"github.com/fseconomy/fseconomy-go/fseconomy"
)

func main() {
	f, err := fseconomy.New()
	if err != nil {
		_, _ = fmt.Fprintf(os.Stderr, `failed to create FSEconomy SDK context: %v\n`, err)
		os.Exit(1)
	}
}
```