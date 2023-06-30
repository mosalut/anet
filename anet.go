package main

var anet struct {
	run func() error
	monitor func() error
}

func anetInit() {
	anet.run = func() error {
		err := listenUDP()
		if err != nil {
			return err
		}
	}

	anet.monitor = func(uint8) error {
	}
}
