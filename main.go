package main

import (
	"log"
)

func init() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	err := setHomeDir()
	if err != nil {
		log.Fatal(err)
	}

	conf = &config{}
	err = conf.read()
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	run()
}
