package main

import (
	"encoding/json"
	"io"
	"log"
	"net"
	"net/http"
	"os"
	"time"
    "partisci/version"
)

const partisci_version = "0.1"
const listenAddr = "localhost:7777"
const updateInterval = 10 * time.Second

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

type OpStats struct {
	updates int64
}

func handleUpdateUDP(conn net.PacketConn, updates chan<- version.Version) {
	for {
		b := make([]byte, 2048)
		n, addr, err := conn.ReadFrom(b)
		if err != nil {
			l.Print("Error reading UDP packet:\n  ", err)
			continue
		}
		ip := addr.(*net.UDPAddr).IP
		update, err := version.ParsePacket(ip.String(), b[:n])
		if err != nil {
			l.Print("parsePacket: ", err)
		}
		updates <- update
	}
}

func processUpdates(updates <-chan version.Version) {
	stats := OpStats{}
	ticker := time.NewTicker(updateInterval)
	go func() {
		for {
			select {
			case <-ticker.C:
				l.Printf("%v updates in last %v", stats.updates, updateInterval)
				stats.updates = 0
			case v := <-updates:
				stats.updates++
				l.Println(v)
			}
		}
	}()
}

// hello world, the web server
func HelloServer(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "hello, world!\n")
}

type InfoRes struct {
    Version string `json:"version"`
}

func ApiPartisci(w http.ResponseWriter, req *http.Request) {
    info := InfoRes{partisci_version}
    data, _ := json.Marshal(info)
    w.Write(data)
}

func main() {
	l.Print("Starting.")

	conn, err := net.ListenPacket("udp", listenAddr)
	if err != nil {
		l.Fatalf("Error opening listen socket: %v\n  %v", listenAddr, err)
	}
	l.Print("listening on: ", conn.LocalAddr())

	updates := make(chan version.Version)
	go processUpdates(updates)
	go handleUpdateUDP(conn, updates)

	http.HandleFunc("/api/v1/_partisci", ApiPartisci)
	http.HandleFunc("/hello", HelloServer)
	err = http.ListenAndServe(listenAddr, nil)
	if err != nil {
		l.Fatal("ListenAndServe: ", err)
	}
	l.Print("Exit.")
}
