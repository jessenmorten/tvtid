# 📺 tvtid

[![Build](https://github.com/jessenmorten/tvtid/actions/workflows/build.yml/badge.svg)](https://github.com/jessenmorten/tvtid/actions/workflows/build.yml)
![GitHub go.mod Go version (subdirectory of monorepo)](https://img.shields.io/github/go-mod/go-version/jessenmorten/tvtid)
[![GoDoc](https://godoc.org/github.com/jessenmorten/tvtid?status.svg)](https://godoc.org/github.com/jessenmorten/tvtid)
[![Go Report Card](https://goreportcard.com/badge/github.com/jessenmorten/tvtid)](https://goreportcard.com/report/github.com/jessenmorten/tvtid)
![License](https://img.shields.io/github/license/jessenmorten/tvtid)

A reverse-engineered implementation of the **[tvtid](https://tvtid.tv2.dk/)** API in Go.

## Getting started

### Install with Go Modules

With [Go Modules](https://github.com/golang/go/wiki/Modules) support, simply import the package to start using it:

```go
import "github.com/jessenmorten/tvtid"
```

### Install with go get

With [Go](https://go.dev) installed, run the following command to install the package:

```bash
go get -u github.com/jessenmorten/tvtid
```

## Usage example

Here's a simple example of how to use the package.
For more information, see the [documentation](https://pkg.go.dev/github.com/jessenmorten/tvtid).

```go
package main

import (
    "fmt"
    "time"

    "github.com/jessenmorten/tvtid"
)

func main() {
    // Create a new client, using the default HTTP client
    client := tvtid.NewDefaultClient()

    // Get all channels
    channels, err := client.GetChannels()
    panicIfErr(err)
    fmt.Println("Found", len(channels), "channels")

    // Get programs for a channel on a specific date
    channelId := channels[0].Id
    date := time.Now()
    programs, err := client.GetPrograms(channelId, date)
    panicIfErr(err)
    fmt.Println("Found", len(programs), "programs")

    // Get program details
    programId := programs[0].Id
    details, err := client.GetProgramDetails(channelId, programId)
    panicIfErr(err)
    fmt.Println("Program details for", programId)

    // Print program details
    fmt.Println("Title:", details.Title)
    fmt.Println("Description:", details.Description)
    fmt.Println("Url:", details.Url)
    fmt.Println("Start:", programs[0].StartTime)
    fmt.Println("Stop:", programs[0].StopTime)
    // ...
}

func panicIfErr(err error) {
    if err != nil {
        panic(err)
    }
}
```
