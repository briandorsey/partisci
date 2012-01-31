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

    conn, err := net.ListenPacket("udp", listenAddr)
	if err != nil {
		log.Fatalf("Error opening listen socket: %v\n  %v", listenAddr, err)
	}
	log.Print("listening on: ", conn.LocalAddr())

    for {
        b := make([]byte, 2048)
        n, addr, _ := conn.ReadFrom(b)
        log.Print(n, addr, string(b))
    }

	log.Print("Exit.")
}
