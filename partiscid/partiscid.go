// Partiscid is the server executable, it listens for version updates.
package main

import (
	"encoding/json"
	"expvar"
	"flag"
	"fmt"
	"io/ioutil"
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
var listenip *string = flag.String("listenip", "", "listen only on this IP (defaults to all)")
var verbose *bool = flag.Bool("v", false, "log more details")
var danger *bool = flag.Bool("danger", false, "enable dangerous commands for testing")

func init() {
	ver := expvar.NewString("version")
	ver.Set(partisci_version)
}

// OpStats contains operational statisics about this server.
type OpStats struct {
	updates int64
}

type UpdateStore interface {
	Update(v version.Version) (err error)
	Apps() (vs []version.AppSummary)
	Hosts() (vs []version.HostSummary)
	Versions(app_id string, host string, ver string) (vs []version.Version)
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
		v, err := version.ParsePacket(ip.String(), b[:n])
		if err != nil {
			l.Print("ERROR: handleUpdateUDP: parsePacket:\n  ", err,
				"\n  packet:", string(b[:n]))
			continue
		}
		updates <- v
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
				if *verbose {
					log.Print("UPDATE: ", v)
				}
			}
		}
	}()
}

type ErrorRes struct {
	Error string `json:"error"`
}

type DataRes struct {
	Data []interface{} `json:"data"`
}

func NewDataRes() (r *DataRes) {
	r = new(DataRes)
	r.Data = make([]interface{}, 0)
	return r
}

type storeServer struct {
	store   UpdateStore
	updates chan<- version.Version
	danger  bool
}

func (ss storeServer) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	switch req.URL.Path {
	case "app/":
		ss.ApiApp(w, req)
	case "host/":
		ss.ApiHost(w, req)
	case "version/":
		ss.ApiVersion(w, req)
	case "update/":
		ss.ApiUpdate(w, req)
	case "_danger/clear/":
		if ss.danger {
			ss.ApiClear(w, req)
		}
	default:
		l.Print("INFO: 404: ", req.URL)
		http.Error(w, "404 page not found", http.StatusNotFound)
	}
	return
}

func (ss *storeServer) ApiVersion(w http.ResponseWriter, req *http.Request) {
	r := NewDataRes()
	app_id := req.FormValue("app_id")
	host := req.FormValue("host")
	ver := req.FormValue("ver")
	vers := ss.store.Versions(app_id, host, ver)
	for _, ver := range vers {
		r.Data = append(r.Data, ver)
	}
	data, err := json.Marshal(r)
	if handleError(err, "ApiVersion", w, http.StatusInternalServerError) {
		return
	}
	w.Write(data)
}

func (ss storeServer) ApiApp(w http.ResponseWriter, req *http.Request) {
	r := NewDataRes()
	apps := ss.store.Apps()
	for _, app := range apps {
		r.Data = append(r.Data, app)
	}
	data, err := json.Marshal(r)
	if handleError(err, "ApiApp", w, http.StatusInternalServerError) {
		return
	}
	w.Write(data)
}

func (ss storeServer) ApiHost(w http.ResponseWriter, req *http.Request) {
	r := NewDataRes()
	hosts := ss.store.Hosts()
	for _, host := range hosts {
		r.Data = append(r.Data, host)
	}
	data, err := json.Marshal(r)
	if handleError(err, "ApiHost", w, http.StatusInternalServerError) {
		return
	}
	w.Write(data)
}

func (ss *storeServer) ApiClear(w http.ResponseWriter, req *http.Request) {
	if req.Method == "POST" {
		l.Print("WARNING: Version database cleared via _danger/clear/ hook.")
		ss.store.Clear()
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

func (ss storeServer) ApiUpdate(w http.ResponseWriter, req *http.Request) {
	if req.Method != "POST" {
		m := "ERROR: ApiUpdate: only accepts POST requests"
		l.Print(m)
		errRes := ErrorRes{Error: m}
		data, _ := json.Marshal(errRes)
		w.WriteHeader(http.StatusMethodNotAllowed)
		w.Write(data)
		return
	}
	b, err := ioutil.ReadAll(req.Body)
	if handleError(err, "ApiUpdate: ReadAll", w,
		http.StatusInternalServerError) {
		return
	}
	host, _, err := net.SplitHostPort(req.RemoteAddr)
	if err != nil {
		host = ""
	}
	v, err := version.ParsePacket(host, b)
	if err != nil {
		handleError(err, "ApiUpdate: parsePacket", w,
			http.StatusUnsupportedMediaType)
		l.Print("ERROR: packet:", string(b))
		return
	}
	ss.updates <- v
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

func main() {
	flag.Parse()
	listenAddr := fmt.Sprintf("%s:%d", *listenip, *port)

	l.Print("Starting.")

	conn, err := net.ListenPacket("udp", listenAddr)
	if err != nil {
		l.Fatalf("Error opening listen socket: %v\n  %v", listenAddr, err)
	}
	l.Print("listening on: ", conn.LocalAddr())

	updates := make(chan version.Version)
	store := memstore.NewMemoryStore()
	ss := storeServer{store, updates, *danger}
	go processUpdates(updates, store)
	go handleUpdateUDP(conn, updates)

	apiRoot := http.StripPrefix("/api/v1/", ss)
	http.Handle("/api/v1/", apiRoot)
	err = http.ListenAndServe(listenAddr, nil)
	if err != nil {
		l.Fatal("ListenAndServe: ", err)
	}
	l.Print("Exit.")
}
