// Package client is a client wrapper for the Partisci API.
package client

import (
	"encoding/json"
	"fmt"
	"net"
	"partisci/version"
)

func SendUDP(server string, port int, v *version.Version) (err error) {
	fmt.Printf("sending version update to %v:%v\n", server, port)
	b, _ := json.Marshal(v)
	fmt.Println(string(b))
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
