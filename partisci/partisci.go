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
		return v, err
	}
	v.AppId = AppNameToID(v.Name)
	return
}

func handleUpdateUDP(conn net.PacketConn, updates chan<- Version) {
	for {
		b := make([]byte, 2048)
		n, addr, err := conn.ReadFrom(b)
		if err != nil {
			l.Print("Error reading UDP packet:\n  ", err)
			continue
		}
		ip := addr.(*net.UDPAddr).IP
		update, err := parsePacket(ip.String(), b[:n])
		if err != nil {
			l.Print("parsePacket: ", err)
		}
		updates <- update
	}
}

func processUpdates(updates <-chan Version) {
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

func main() {
	l.Print("Starting.")

	conn, err := net.ListenPacket("udp", listenAddr)
	if err != nil {
		l.Fatalf("Error opening listen socket: %v\n  %v", listenAddr, err)
	}
	l.Print("listening on: ", conn.LocalAddr())

	updates := make(chan Version)
	go processUpdates(updates)
	go handleUpdateUDP(conn, updates)

	http.HandleFunc("/", HelloServer)
	err = http.ListenAndServe(listenAddr, nil)
	if err != nil {
		l.Fatal("ListenAndServe: ", err)
	}
	l.Print("Exit.")
}
