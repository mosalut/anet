package main

import (
	"log"
)

func run() {
	err := listenUDP()
	if err != nil {
		log.Fatal(err)
	}
}
