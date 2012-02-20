package client_test

import "partisci/client"
import "partisci/version"

func ExampleSendUDP() {
	v := version.Version{App: "app", Ver: "1.2.3", Host: "example.com"}
	err := client.SendUDP("localhost", 7777, v)
	if err != nil {
		// handle error
	}
}
