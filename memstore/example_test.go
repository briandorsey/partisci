package memstore_test

import "fmt"
import "partisci/memstore"
import "partisci/version"

var s *memstore.MemoryStore

func init() {
	s = memstore.NewMemoryStore()
}

func ExampleMemoryStore_Versions_all() {
	// returns all known Version structs
	vers := s.Versions("", "", "")
	for v := range vers {
		fmt.Println(v)
	}
}

func ExampleMemoryStore_Versions_appId() {
	// returns all known Version structs for a specific AppId
	vers := s.Versions("app1", "", "")
	for v := range vers {
		fmt.Println(v)
	}
}

func ExampleMemoryStore_Update() {
	v := version.Version{App: "app1", Ver: "1.2.3", Host: "host1.example.com"}
	err := s.Update(v)
	if err != nil {
		// handle error
	}
}
