// Package store defines the UpdateStore interface for version persistence.
//
// The USTest* functions are tests which should be run by implementations
// of this interface to ensure compatability.
package store

import (
	"partisci/version"
	"testing"
	"time"
)

// UpdateStore defines an interface for persisting application version information.
type UpdateStore interface {
	// Update stores a Version and updates app and host summaries.
	Update(v version.Version) (err error)

	// App returns an AppSummary for the given AppId.
	// The value of ok follows map indexing conventions: 
	//     true if AppId is present, false otherwise.
	App(AppId string) (as version.AppSummary, ok bool)

	// Apps returns summary information about each application, 
	// based on the known Versions.
	Apps() (vs []version.AppSummary)

	// Host returns a HostSummary for the given Host.
	// The value of ok follows map indexing conventions: 
	//   true if Host is present, false otherwise.
	Host(Host string) (hs version.HostSummary, ok bool)

	// Hosts returns summary information about each host, 
	// based on the known Versions.
	Hosts() (vs []version.HostSummary)

	// Versions returns full Version structs where their values match app_id, host
	// and ver. Zero length strings are considered a match for all Versions.
	Versions(app_id string, host string, ver string) (vs []version.Version)

	// Clear empties the MemoryStore.
	Clear()

	// Trim removes old versions.
	Trim(t time.Time) (c uint64)
}

func USTestAppSummary(s UpdateStore, t *testing.T) {
	v := version.Version{App: "app1", Ver: "ver", Host: "a"}
	v.Prepare()
	s.Update(v)

	if _, ok := s.App("non-existant"); ok {
		t.Error("got ok for non-existant AppId")
	}
	if as, ok := s.App("app1"); ok {
		if as.HostCount != 1 {
			t.Error("expected HostCount: 1, actual: ", as.HostCount)
		}
	} else {
		t.Error("missing expected AppId")
	}

	v2 := version.Version{App: "app1", Ver: "ver", Host: "b"}
	v2.Prepare()
	s.Update(v2)
	if as, ok := s.App("app1"); ok {
		if as.HostCount != 2 {
			t.Error("expected HostCount: 2, actual: ", as.HostCount)
		}
	}
}

func USTestHostSummary(s UpdateStore, t *testing.T) {
	v := version.Version{App: "app1", Ver: "ver", Host: "a"}
	v.Prepare()
	s.Update(v)

	if _, ok := s.Host("non-existant"); ok {
		t.Error("got ok for non-existant Host")
	}
	if as, ok := s.Host("a"); ok {
		if as.AppCount != 1 {
			t.Error("expected AppCount: 1, actual: ", as.AppCount)
		}
	} else {
		t.Error("missing expected AppId")
	}

	v2 := version.Version{App: "app2", Ver: "ver", Host: "a"}
	v2.Prepare()
	s.Update(v2)
	if as, ok := s.Host("a"); ok {
		if as.AppCount != 2 {
			t.Error("expected AppCount: 2, actual: ", as.AppCount)
		}
	}
}

func USTestTrim(s UpdateStore, t *testing.T) {
	// setup one version in the future and a few more
	v1a := version.Version{App: "app1", Ver: "ver", Host: "a"}
	v1a.Prepare()
	v1a.ExactUpdate = v1a.ExactUpdate.Add(time.Duration(10 * time.Second))
	s.Update(v1a)

	v1b := version.Version{App: "app1", Ver: "ver", Host: "b"}
	v1b.Prepare()
	s.Update(v1b)

	v2 := version.Version{App: "app2", Ver: "ver", Host: "a"}
	v2.Prepare()
	s.Update(v2)

	// sanity check
	if l := len(s.Versions("", "", "")); l != 3 {
		t.Error("before: version count - expected: 2, actual: ", l)
	}
	if l := len(s.Hosts()); l != 2 {
		t.Error("before: host count - expected: 2, actual: ", l)
	}
	if l := len(s.Apps()); l != 2 {
		t.Error("before: app count - expected: 2, actual: ", l)
	}

	// trim every version before 1 second in the future of one version
	count := s.Trim(v2.ExactUpdate.Add(time.Duration(1 * time.Second)))
	if count != 2 {
		t.Error("after: trim should have removed 2 versions")
	}
	if l := len(s.Versions("", "", "")); l != 1 {
		t.Error("after: version count - expected: 1, actual: ", l)
	}
	if l := len(s.Hosts()); l != 1 {
		t.Error("after: host count - expected: 1, actual: ", l)
	}
	if l := len(s.Apps()); l != 1 {
		t.Error("after: app count - expected: 1, actual: ", l)
	}

	// trim every version
	count = s.Trim(v2.ExactUpdate.Add(time.Duration(20 * time.Second)))
	if count != 1 {
		t.Error("after all: trim should have removed the last one version")
	}
	if l := len(s.Versions("", "", "")); l != 0 {
		t.Error("after all: version count - expected: 0, actual: ", l)
	}
	if l := len(s.Hosts()); l != 0 {
		t.Error(s.Hosts())
		t.Error("after all: host count - expected: 0, actual: ", l)
	}
	if l := len(s.Apps()); l != 0 {
		t.Error(s.Apps())
		t.Error("after all: app count - expected: 0, actual: ", l)
	}
}
