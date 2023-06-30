package main

import (
	"log"
	"net"
)

func listenUDP() error {
	laddr, err := net.ResolveUDPAddr("udp4", conf.host)
	if err != nil {
		return err
	}

	conn, err := net.ListenUDP("udp4", laddr)
	if err != nil {
		return err
	}
	defer conn.Close()

	log.Println("UDP Listening...")
	err = handleClient(conn)
	if err != nil {
		return err
	}

	return nil
}

func handleClient(conn *net.UDPConn) error {
	data := make([]byte, 65506, 65506)
	n, remoteAddr, err := conn.ReadFromUDP(data)
	if err != nil {
		return err
	}

	log.Println(remoteAddr)
	log.Println(string(data))
	log.Println(n)

	conn.Close()
//	sendUDP(remoteAddr)

	return nil
}
