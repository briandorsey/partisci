package memstore

import (
	"log"
	"partisci/store"
	"partisci/version"
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

// test Clear() & Update() interactions
func TestClearUpdate(t *testing.T) {
	s := NewMemoryStore()
	log.Print(s)
	if len(s.Versions("", "", "")) > 0 {
		t.Error("Versions should be empty")
	}
	v := *new(version.Version)
	v.Prepare()
	s.Update(v)
	if len(s.Versions("", "", "")) != 1 {
		log.Print(s.threshold, s.threshold.Unix())
		log.Print(v.ExactUpdate, v.LastUpdate)
		t.Error("Versions should have one entry")
	}
	s.Clear()
	if len(s.Versions("", "", "")) > 0 {
		t.Error("Versions should be empty")
	}
	s.Update(v)
	if len(s.Versions("", "", "")) > 0 {
		t.Error("updates older than threshold should be discarded")
	}
}

func TestTrim(t *testing.T) {
	s := NewMemoryStore()
	store.USTestTrim(s, t)
}
