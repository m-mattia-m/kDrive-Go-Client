# kDrive-go-client

[![GitHub tag (latest SemVer)](https://img.shields.io/github/v/tag/m-mattia-m/kdrive-go-client?label=go%20module)](https://github.com/m-mattia-m/kdrive-go-client/tags)
[![Go Reference](https://pkg.go.dev/badge/github.com/m-mattia-m/kdrive-go-client.svg)](https://pkg.go.dev/github.com/m-mattia-m/kdrive-go-client)
[![Test](https://github.com/m-mattia-m/kdrive-go-client/actions/workflows/test.yml/badge.svg)](https://github.com/m-mattia-m/kdrive-go-client/actions/workflows/test.yml)

> Is currently under development and does not yet include all functions. Feel free to fork and make pull requests to further develop the client together.

This is a client for the integration of [Infomaniak's](https://www.infomaniak.com/de) [kDrive](https://www.infomaniak.com/de/kdrive). 

# Installation

```
$ go get github.com/m-mattia-m/kdrive-go-client
```

# Getting started
The official [development website of Infomaniak](https://developer.infomaniak.com/) is used as API-documentation.

## Example

Make a new `Client`

```go
import "github.com/m-mattia-m/kdrive-go-client"
client := kDrive.NewClient("your-token")
```
