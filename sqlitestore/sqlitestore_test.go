// +build ignore

package sqlitestore

import (
	"github.com/briandorsey/partisci/sharedtest"
	"log"
	"os"
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
		// log.Print("cleanup(): removed file")
	}
}

func prepStore(t *testing.T) (s *SQLiteStore) {
	path := filepath.Join(os.TempDir(), dbPath)
	cleanup(path)
	s, err := NewSQLiteStore(path)
	if err != nil {
		t.Fatal(err)
	}
	return s
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
	s := prepStore(t)
	sharedtest.USTestAppSummary(s, t)
}

func TestHostSummary(t *testing.T) {
	s := prepStore(t)
	sharedtest.USTestHostSummary(s, t)
}

func TestClearUpdate(t *testing.T) {
	s := prepStore(t)
	sharedtest.USTestClearUpdate(s, t)
}

func TestTrim(t *testing.T) {
	s := prepStore(t)
	sharedtest.USTestTrim(s, t)
}
