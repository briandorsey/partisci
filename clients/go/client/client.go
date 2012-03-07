// Package client implements a wrapper for the Partisci API.
package client

import (
	"encoding/json"
	"fmt"
	"github.com/briandorsey/partisci/version"
	"net"
)

// SendUDP serializes and sends a single Version to a partiscid server via UDP.
func SendUDP(server string, port int, v version.Version) (err error) {
	b, err := json.Marshal(v)
	if err != nil {
		return err
	}
	addr := fmt.Sprintf("%s:%d", server, port)

	conn, err := net.Dial("udp", addr)
	if err != nil {
		return err
	}
	_, err = conn.Write(b)
	if err != nil {
		return err
	}
	err = conn.Close()
	if err != nil {
		return err
	}
	return nil
}
