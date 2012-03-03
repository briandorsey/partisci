package memstore

import (
	"partisci/sharedtest"
	"testing"
)

func TestAppSummary(t *testing.T) {
	s := NewMemoryStore()
	sharedtest.USTestAppSummary(s, t)
}

func TestHostSummary(t *testing.T) {
	s := NewMemoryStore()
	sharedtest.USTestHostSummary(s, t)
}

func TestClearUpdate(t *testing.T) {
	s := NewMemoryStore()
	sharedtest.USTestClearUpdate(s, t)
}

func TestTrim(t *testing.T) {
	s := NewMemoryStore()
	sharedtest.USTestTrim(s, t)
}
