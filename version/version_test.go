package version

import (
	"testing"
)

func TestAppToID(t *testing.T) {
	if "lower" != AppToID("LoWeR") {
		t.Error("ids should be all lowercase")
	}
	if "___________" != AppToID("_!@#$%^&*( ") {
		t.Error("non alpha should be converted to underscores")
	}
	if "0123456789" != AppToID("0123456789") {
		t.Error("digits should be preserved")
	}
}

func BenchmarkAppToID(b *testing.B) {
	s := "Longish Application Name!"
	for i := 0; i < b.N; i++ {
		_ = AppToID(s)
	}
}

func BenchmarkParsePacket(b *testing.B) {
	s := `{"instance": 0, "host": "hostname", "version": "0.1test", "app": "test"}`
	for i := 0; i < b.N; i++ {
		_, _ = ParsePacket("0.0.0.0", []byte(s))
	}
}
