package main

import (
	"os/user"
	"log"

	"gopkg.in/ini.v1"
)

const __CONF__ = ".anet/config"

var homeDir string

var conf *config

func setHomeDir() error {
	user, err := user.Current()
	if err != nil {
		return err
	}

	homeDir = user.HomeDir + "/"

	return nil
}

type config struct {
	release bool
	host string
}

func (c *config) read () error {
	cfg, err := ini.Load(homeDir + __CONF__)
	if err != nil {
		return err
	}

	c.release, err = cfg.Section("").Key("release").Bool()
	if err != nil {
		log.Fatal(err)
	}
	c.host = cfg.Section("").Key("host").String()

	return nil
}
