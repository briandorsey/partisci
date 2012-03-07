package memstore_test

import "fmt"
import "github.com/briandorsey/partisci/memstore"
import "github.com/briandorsey/partisci/version"

var s *memstore.MemoryStore

func init() {
	s = memstore.NewMemoryStore()
}

func ExampleMemoryStore_Versions_all() {
	// returns all known Version structs
	vers, err := s.Versions("", "", "")
	if err != nil {
		// handle error
	}
	for v := range vers {
		fmt.Println(v)
	}
}

func ExampleMemoryStore_Versions_appId() {
	// returns all known Version structs for a specific AppId
	vers, err := s.Versions("app1", "", "")
	if err != nil {
		// handle error
	}
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
