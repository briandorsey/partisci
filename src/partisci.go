package main

import (
	"encoding/json"
	logpkg "log"
	"net"
	"os"
	"time"
    "strings"
)

const listenAddr = "localhost:7777"

var log = logpkg.New(os.Stderr, "", logpkg.Ldate|logpkg.Ltime)

type Version struct {
	AppId      string `json:"app_id"`
	Name       string `json:"name"`
	Version    string `json:"version"`
	Host       string `json:"host"`
	Instance   uint16 `json:"instance"`
	HostIP     string `json:"host_ip"`
	LastUpdate int64  `json:"last_update"`
}

func safeRunes(r rune) rune {
    if 'a' <= r && r <= 'z' {
        return r
    }
    return '_'
}

func AppNameToID(app string) (id string) {
    id = strings.ToLower(app)
    id = strings.Map(safeRunes, id)
    return
}

func parsePacket(host string, b []byte) (v Version, err error) {
	v.HostIP = host
	v.LastUpdate = time.Now().Unix()
	err = json.Unmarshal(b[:len(b)], &v)
	if err != nil {
		log.Print("parsePacket: ", err)
	}
    v.AppId = AppNameToID(v.Name)
    log.Print(v)
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
		n, addr, err := conn.ReadFrom(b)
		if err != nil {
			log.Print("Error reading UDP packet:\n  ", err)
			continue
		}
		ip := addr.(*net.UDPAddr).IP
		parsePacket(ip.String(), b[:n])
	}

	log.Print("Exit.")
}
