# go3270 ![Build Status](https://travis-ci.org/msalcantara/go3270.svg?branch=master) [![Release](https://img.shields.io/github/release/msalcantara/go3270.svg)](https://github.com/msalcantara/go3270/releases) [![License: MIT](https://img.shields.io/badge/License-MIT-blue.svg)](https://opensource.org/licenses/MIT) [![GoDoc](https://godoc.org/github.com/msalcantara/go3270?status.png)](https://godoc.org/github.com/msalcantara/go3270)

go3270 is a Golang interface to 3270 terminal emulator.

## Requirements
go3270 require  [x3270](http://x3270.bgp.nu/index.html) executables in your PATH

## Install
```$ go get github.com/msalcantara/go3270```

## Example
```go
package main

import (
	"log"

	"github.com/msalcantara/go3270"
)

func main() {

	e := go3270.Emulator{
		Host: "0.0.0.0",
		Port: 992,
	}

	if err := e.Connect(); err != nil {
		log.Fatalf("error to create connection: %v\n", err)
	}

	if err := e.SetString("my_user"); err != nil {
		log.Fatalf("error to set user: %v\n", err)
	}

	if err := e.Press(go3270.Tab); err != nil {
		log.Fatalf("error to press tab: %v\n", err)
	}

	if err := e.SetString("my_password"); err != nil {
		log.Fatalf("error to set password: %v\n", err)
	}

	if err := e.Press(go3270.Enter); err != nil {
		log.Fatalf("error to press enter: %v\n", err)
	}

	v, err := e.GetValue(0, 1, 4)
	if err != nil {
		log.Fatalf("error to get value: %v", err)
	}
	log.Println(v)

	if err := e.Disconnect(); err != nil {
		log.Fatalf("error to disconnect: %v\n", err)
	}

}
```
## License

go3270 is licensed under the [MIT][mit] license.

[mit]: https://raw.githubusercontent.com/tebeka/selenium/master/LICENSE
