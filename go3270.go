/*
Package go3270 is a x3270 terminal emulator and screen-scraping library written entirely in Go.

Go3270 use x3270 scripts to emulators terminal, so, you need to set x3270 in your $PATH to go3270 works.

Example:
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
*/
package go3270

//version of package - based on Semantic Versioning 2.0.0 http://semver.org/
const version = "0.1.0"
