package memstore_test

import "fmt"
import "partisci/memstore"
import "partisci/version"

var s *memstore.MemoryStore

func init() {
	s = memstore.NewMemoryStore()
}

func ExampleMemoryStore_Versions() {
	// returns all known Version structs
	vers := s.Versions("", "", "")
	for v := range vers {
		fmt.Println(v)
	}
}

func ExampleMemoryStore_Update() {
	v := version.Version{App: "app", Ver: "1.2.3", Host: "example.com"}
    err := s.Update(v)
	if err != nil {
		// handle error
	}
}
