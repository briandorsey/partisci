package memstore

import (
	"log"
	"partisci/version"
	"testing"
)

// test Clear() & Update() interactions
func TestClearUpdate(t *testing.T) {
	s := NewMemoryStore()
	log.Print(s)
	if len(s.Versions("", "", "")) > 0 {
		t.Error("Versions should be empty")
	}
	v := *version.NewVersion()
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
