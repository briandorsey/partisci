// +build ignore

package main

import (
	"log"
	"github.com/briandorsey/partisci/clients/go/client"
	"github.com/briandorsey/partisci/version"
)

func main() {
	v := version.Version{App: "app", Ver: "1.2.3",
		Host: "example.com", Instance: 0}
	err := client.SendUDP("localhost", 7777, v)
	if err != nil {
		log.Fatal(err)
	}
}
