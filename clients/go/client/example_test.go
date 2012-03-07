package client_test

import "github.com/briandorsey/partisci/clients/go/client"
import "github.com/briandorsey/partisci/version"

func ExampleSendUDP() {
	v := version.Version{App: "app", Ver: "1.2.3", Host: "example.com"}
	err := client.SendUDP("localhost", 7777, v)
	if err != nil {
		// handle error
	}
}
