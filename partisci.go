package main

import (
	"encoding/json"
	"io"
	"log"
	"net"
	"net/http"
	"os"
	"partisci/memstore"
	"partisci/version"
	"time"
)

const partisci_version = "0.1"
const listenAddr = ":7777"
const updateInterval = 10 * time.Second

var l = log.New(os.Stderr, "", log.Ldate|log.Ltime)

type OpStats struct {
	updates int64
}

type UpdateStore interface {
	Apps() (vs []version.Version)
	Update(v version.Version) (err error)
}

func handleUpdateUDP(conn net.PacketConn, updates chan<- version.Version) {
	for {
		b := make([]byte, 2048)
		n, addr, err := conn.ReadFrom(b)
		if err != nil {
			l.Print("ERROR: handleUpdateUDP: ReadFrom:\n  ", err)
			continue
		}
		ip := addr.(*net.UDPAddr).IP
		update, err := version.ParsePacket(ip.String(), b[:n])
		if err != nil {
			l.Print("ERROR: handleUpdateUDP: parsePacket:\n  ", err)
			continue
		}
		updates <- update
	}
}

func processUpdates(updates <-chan version.Version, store UpdateStore) {
	stats := OpStats{}
	ticker := time.NewTicker(updateInterval)
	go func() {
		for {
			select {
			case <-ticker.C:
				l.Printf("STAT: %v updates in last %v", stats.updates, updateInterval)
				stats.updates = 0
			case v := <-updates:
				stats.updates++
				store.Update(v)
			}
		}
	}()
}

type InfoRes struct {
	Version string `json:"version"`
}

type ErrorRes struct {
	Error string `json:"error"`
}

type DataRes struct {
	Data []version.Version `json:"data"`
}

func NewDataRes() (r *DataRes) {
	r = new(DataRes)
	r.Data = make([]version.Version, 0)
	return r
}

// HTTP handlers

// hello world, the web server
func HelloServer(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "hello, world!\n")
}

func ApiPartisci(w http.ResponseWriter, req *http.Request) {
	info := InfoRes{partisci_version}
	data, _ := json.Marshal(info)
	w.Write(data)
}

func ApiApp(w http.ResponseWriter, req *http.Request, s UpdateStore) {
	r := NewDataRes()
	r.Data = s.Apps()
	data, err := json.Marshal(r)
	if err != nil {
		m := "ERROR: ApiApp: " + err.Error()
		l.Print(m)
		errRes := ErrorRes{Error: m}
		data, _ := json.Marshal(errRes)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write(data)
		return
	}
	w.Write(data)
}

func makeStoreHandler(fn func(w http.ResponseWriter, req *http.Request, s UpdateStore), s UpdateStore) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		fn(w, req, s)
	}
}

func main() {
	l.Print("Starting.")

	conn, err := net.ListenPacket("udp", listenAddr)
	if err != nil {
		l.Fatalf("Error opening listen socket: %v\n  %v", listenAddr, err)
	}
	l.Print("listening on: ", conn.LocalAddr())

	updates := make(chan version.Version)
	store := memstore.NewMemoryStore()
	go processUpdates(updates, store)
	go handleUpdateUDP(conn, updates)

	http.HandleFunc("/api/v1/_partisci/", ApiPartisci)
	http.HandleFunc("/api/v1/app/", makeStoreHandler(ApiApp, store))
	http.HandleFunc("/hello", HelloServer)
	err = http.ListenAndServe(listenAddr, nil)
	if err != nil {
		l.Fatal("ListenAndServe: ", err)
	}
	l.Print("Exit.")
}
