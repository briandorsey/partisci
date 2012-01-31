package main

import (
	logpkg "log"
	"net"
	"os"
)

const listenAddr = "localhost:7777"

func main() {
	log := logpkg.New(os.Stderr, "", logpkg.Ldate|logpkg.Ltime)
	log.Print("Starting.")

	address, err := net.ResolveUDPAddr("udp", listenAddr)
	if err != nil {
		log.Fatalf("Unable to resolve address: %v\n  %v", listenAddr, err)
	}

	conn, err := net.ListenUDP("udp", address)
	if err != nil {
		log.Fatalf("Error opening listen socket: %v\n  %v", listenAddr, err)
	}
	log.Print("listening on:", conn.LocalAddr())


	log.Print("Exit.")
}
