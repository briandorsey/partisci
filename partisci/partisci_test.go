package partisci

import (
	"testing"
)

func TestAppNameToID(t *testing.T) {
	if "lower" != AppNameToID("LoWeR") {
		t.Error("ids should be all lowercase")
	}
	if "___________" != AppNameToID("_!@#$%^&*( ") {
		t.Error("non alpha should be converted to underscores")
	}
}

func BenchmarkAppNameToID(b *testing.B) {
	s := "Longish Application Name!"
	for i := 0; i < b.N; i++ {
		_ = AppNameToID(s)
	}
}

func BenchmarkParsePacket(b *testing.B) {
	s := `{"instance": 0, "host": "hostname", "version": "0.1test", "name": "test"}`
	for i := 0; i < b.N; i++ {
		_, _ = parsePacket("0.0.0.0", []byte(s))
	}
}
