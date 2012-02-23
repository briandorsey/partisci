package version

import (
	"testing"
        "time"
)

func TestVersionKey(t *testing.T) {
	v := new(Version)
	v.App = "app name"
	v.Host = "hostname"
	v.Instance = 0
	v.Prepare()

	// same values
	vs := new(Version)
	vs.App = "app name"
	vs.Host = "hostname"
	vs.Instance = 0
	vs.Prepare()
	if v.Key() != vs.Key() {
		t.Errorf("keys not equal: %v != %v", v.Key(), vs.Key())
	}

	// only instance varies
	vi := new(Version)
	vi.App = "app name"
	vi.Host = "hostname"
	vi.Instance = 1
	vi.Prepare()
	if v.Key() == vi.Key() {
		t.Errorf("varied instance, but keys equal: %v == %v", v.Key(), vi.Key())
	}
}

func TestVersionPrepare(t *testing.T) {
    v := Version{App: "app", Ver: "ver"}
    v.Prepare()
    zerot := new(time.Time)
    if v.ExactUpdate == *zerot {
        t.Errorf("v.Prepare should initialize ExactUpdate if needed:\n%v", v)
    }
    if v.LastUpdate == 0 {
        t.Errorf("v.Prepare should initialize LastUpdate if needed:\n%v", v)
    }
}

func TestAppIdToId(t *testing.T) {
	if "lower" != appIdToId("LoWeR") {
		t.Error("ids should be all lowercase")
	}
	if "___________" != appIdToId("_!@#$%^&*( ") {
		t.Error("non alpha should be converted to underscores")
	}
	if "0123456789" != appIdToId("0123456789") {
		t.Error("digits should be preserved")
	}
}

func BenchmarkVersionKey(b *testing.B) {
	v := new(Version)
	v.App = "benchmark version key app name"
	v.Host = "hostname"
	v.Instance = 0
	v.Prepare()
	for i := 0; i < b.N; i++ {
		_ = v.Key()
	}
}

func BenchmarkAppToID(b *testing.B) {
	s := "Longish Application Name!"
	for i := 0; i < b.N; i++ {
		_ = appIdToId(s)
	}
}

func BenchmarkParsePacket(b *testing.B) {
	s := `{"instance": 0, "host": "hostname", "ver": "0.1test", "app": "test"}`
	for i := 0; i < b.N; i++ {
		_, _ = ParsePacket("0.0.0.0", []byte(s))
	}
}
