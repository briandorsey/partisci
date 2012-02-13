package version

import (
	"testing"
)

func TestVersionKey(t *testing.T) {
	v := NewVersion()
	v.App = "app name"
	v.AppId = AppIdToId(v.App)
	v.Host = "hostname"
	v.Instance = 0

	// same values
	vs := NewVersion()
	vs.App = "app name"
	vs.AppId = AppIdToId(v.App)
	vs.Host = "hostname"
	vs.Instance = 0
	if v.Key() != vs.Key() {
		t.Errorf("keys not equal: %v != %v", v.Key(), vs.Key())
	}

	// only instance varies
	vi := NewVersion()
	vi.App = "app name"
	vi.AppId = AppIdToId(v.App)
	vi.Host = "hostname"
	vi.Instance = 1
	if v.Key() == vi.Key() {
		t.Errorf("varied instance, but keys equal: %v == %v", v.Key(), vi.Key())
	}

}

func TestAppIdToId(t *testing.T) {
	if "lower" != AppIdToId("LoWeR") {
		t.Error("ids should be all lowercase")
	}
	if "___________" != AppIdToId("_!@#$%^&*( ") {
		t.Error("non alpha should be converted to underscores")
	}
	if "0123456789" != AppIdToId("0123456789") {
		t.Error("digits should be preserved")
	}
}

func BenchmarkVersionKey(b *testing.B) {
	v := NewVersion()
	v.App = "benchmark version key app name"
	v.AppId = AppIdToId(v.App)
	v.Host = "hostname"
	v.Instance = 0
	for i := 0; i < b.N; i++ {
		_ = v.Key()
	}
}

func BenchmarkAppToID(b *testing.B) {
	s := "Longish Application Name!"
	for i := 0; i < b.N; i++ {
		_ = AppIdToId(s)
	}
}

func BenchmarkParsePacket(b *testing.B) {
	s := `{"instance": 0, "host": "hostname", "ver": "0.1test", "app": "test"}`
	for i := 0; i < b.N; i++ {
		_, _ = ParsePacket("0.0.0.0", []byte(s))
	}
}
