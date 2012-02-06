package main

import (
	"encoding/json"
	"flag"
	"fmt"
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
const updateInterval = 10 * time.Second

var l = log.New(os.Stderr, "", log.Ldate|log.Ltime)
var port *int = flag.Int("port", 7777, "listening port (both UDP and HTTP server)")
var danger *bool = flag.Bool("danger", false, "enable dangerous commands for testing")

type OpStats struct {
	updates int64
}

type UpdateStore interface {
	Update(v version.Version) (err error)
	Apps() (vs []version.Version)
	Hosts() (vs []version.Version)
	Versions() (vs []version.Version)
	Clear()
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
	data, err := json.Marshal(info)
	if handleError(err, "ApiPartisci", w, http.StatusInternalServerError) {
		return
	}
	w.Write(data)
}

func ApiApp(w http.ResponseWriter, req *http.Request, s UpdateStore) {
	r := NewDataRes()
	r.Data = s.Apps()
	data, err := json.Marshal(r)
	if handleError(err, "ApiApp", w, http.StatusInternalServerError) {
		return
	}
	w.Write(data)
}

func ApiHost(w http.ResponseWriter, req *http.Request, s UpdateStore) {
	r := NewDataRes()
	r.Data = s.Hosts()
	data, err := json.Marshal(r)
	if handleError(err, "ApiHost", w, http.StatusInternalServerError) {
		return
	}
	w.Write(data)
}

func ApiVersion(w http.ResponseWriter, req *http.Request, s UpdateStore) {
	r := NewDataRes()
	r.Data = s.Versions()
	data, err := json.Marshal(r)
	if handleError(err, "ApiVersion", w, http.StatusInternalServerError) {
		return
	}
	w.Write(data)
}

func ApiClear(w http.ResponseWriter, req *http.Request, s UpdateStore) {
	if req.Method == "POST" {
		l.Print("WARNING: Version database cleared via testing hook.")
		s.Clear()
	} else {
		m := "ERROR: ApiClear: only accepts POST requests"
		l.Print(m)
		errRes := ErrorRes{Error: m}
		data, _ := json.Marshal(errRes)
		w.WriteHeader(http.StatusMethodNotAllowed)
		w.Write(data)
		return
	}
}

func handleError(err error, source string, w http.ResponseWriter, code int) bool {
	if err != nil {
		m := fmt.Sprintf("ERROR: %s:\n  %s", source, err.Error())
		l.Print(m)
		errRes := ErrorRes{Error: m}
		data, _ := json.Marshal(errRes)
		w.WriteHeader(code)
		w.Write(data)
		return true
	}
	return false
}

func makeStoreHandler(fn func(w http.ResponseWriter, req *http.Request, s UpdateStore), s UpdateStore) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		fn(w, req, s)
	}
}

func main() {
	flag.Parse()
	listenAddr := fmt.Sprintf(":%d", *port)

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
	http.HandleFunc("/api/v1/summary/app/", makeStoreHandler(ApiApp, store))
	http.HandleFunc("/api/v1/summary/host/", makeStoreHandler(ApiHost, store))
	http.HandleFunc("/api/v1/version/", makeStoreHandler(ApiVersion, store))
	if *danger {
		http.HandleFunc("/api/v1/_danger/clear/", makeStoreHandler(ApiClear, store))
	}
	http.HandleFunc("/hello", HelloServer)
	err = http.ListenAndServe(listenAddr, nil)
	if err != nil {
		l.Fatal("ListenAndServe: ", err)
	}
	l.Print("Exit.")
}
