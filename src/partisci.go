package main

import (
	logpkg "log"
	"net"
	"os"
	"time"
)

const listenAddr = "localhost:7777"

var log = logpkg.New(os.Stderr, "", logpkg.Ldate|logpkg.Ltime)

type Version struct {
	Name       string
	Version    string
	Host       string
	Instance   uint16
	HostIP     string `json:"host_ip"`
	LastUpdate int64
}

func parsePacket(host string, b []byte) (v Version, err error) {
	v.HostIP = host
	v.LastUpdate = time.Now().Unix()
	log.Println(host, string(b))
	log.Println(v)
	return
}

func main() {
	log.Print("Starting.")

	conn, err := net.ListenPacket("udp", listenAddr)
	if err != nil {
		log.Fatalf("Error opening listen socket: %v\n  %v", listenAddr, err)
	}
	log.Print("listening on: ", conn.LocalAddr())

	for {
		b := make([]byte, 2048)
		_, addr, err := conn.ReadFrom(b)
		if err != nil {
			log.Print("Error reading UDP packet:\n  ", err)
			continue
		}
		ip := addr.(*net.UDPAddr).IP
		parsePacket(ip.String(), b)
	}

	log.Print("Exit.")
}
