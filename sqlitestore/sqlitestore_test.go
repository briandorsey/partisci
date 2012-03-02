package sqlitestore

import (
	"log"
	"os"
	"partisci/store"
	"path/filepath"
	"testing"
)

const dbPath = "partisci_testing_db.sqlite"

func cleanup(path string) {
	err := os.Remove(path)
	if err != nil {
		if os.IsNotExist(err) {
			log.Print("cleanup(): didn't exist")
		} else {
			log.Print("cleanup(): ", err)
		}
	} else {
		log.Print("cleanup(): removed file")
	}
}

func TestNewSQLiteStore(t *testing.T) {
	path := filepath.Join(os.TempDir(), dbPath)
	log.Print(path)
	cleanup(path)
	s, err := NewSQLiteStore(path)
	if err != nil {
		t.Fatal(err)
	} else {
		s.Close()
	}
	// the sqlite database should be created
	if _, err := os.Stat(path); err != nil {
		if os.IsNotExist(err) {
			t.Fatal("SQLite db file not found: ", err)
		} else {
			t.Fatal(err)
		}
	}

	// ensure we can open it again, implying db init idempotency
	s2, err := NewSQLiteStore(path)
	if err != nil {
		t.Fatal(err)
	} else {
		s2.Close()
	}
}

func TestAppSummary(t *testing.T) {
	path := filepath.Join(os.TempDir(), dbPath)
	s, err := NewSQLiteStore(path)
	if err != nil {
		t.Fatal(err)
    }
	store.USTestAppSummary(s, t)
}
