package main

import (
	"encoding/json"
	"io"
	"log"
	"net"
	"net/http"
	"os"
	"strings"
	"time"
)

const listenAddr = "localhost:7777"

var l = log.New(os.Stderr, "", log.Ldate|log.Ltime)

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
		l.Print("parsePacket: ", err)
	}
	v.AppId = AppNameToID(v.Name)
	l.Print(v)
	return
}

func handleUpdateUDP(conn net.PacketConn) {
	for {
		b := make([]byte, 2048)
		n, addr, err := conn.ReadFrom(b)
		if err != nil {
			l.Print("Error reading UDP packet:\n  ", err)
			continue
		}
		ip := addr.(*net.UDPAddr).IP
		parsePacket(ip.String(), b[:n])
	}
}

// hello world, the web server
func HelloServer(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "hello, world!\n")
}

func main() {
	l.Print("Starting.")

	conn, err := net.ListenPacket("udp", listenAddr)
	if err != nil {
		l.Fatalf("Error opening listen socket: %v\n  %v", listenAddr, err)
	}
	l.Print("listening on: ", conn.LocalAddr())

	go handleUpdateUDP(conn)

	http.HandleFunc("/", HelloServer)
	err = http.ListenAndServe(listenAddr, nil)
	if err != nil {
		l.Fatal("ListenAndServe: ", err)
	}
	l.Print("Exit.")
}
