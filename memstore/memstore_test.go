package memstore

import (
	"partisci/store"
	"testing"
)

func TestAppSummary(t *testing.T) {
	s := NewMemoryStore()
	store.USTestAppSummary(s, t)
}

func TestHostSummary(t *testing.T) {
	s := NewMemoryStore()
	store.USTestHostSummary(s, t)
}

func TestClearUpdate(t *testing.T) {
	s := NewMemoryStore()
	store.USTestClearUpdate(s, t)
}

func TestTrim(t *testing.T) {
	s := NewMemoryStore()
	store.USTestTrim(s, t)
}
