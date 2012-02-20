// Package memstore is an in-memory implementation of the UpdateStore interface.
package memstore

import (
	"partisci/version"
	"time"
)

// MemoryStore is an in-memory implementation of the UpdateStore interface.
type MemoryStore struct {
	version   map[string]version.Version
	app       map[string]version.AppSummary
	host      map[string]version.HostSummary
	threshold time.Time
}

func initMemoryStore(m *MemoryStore) {
	m.version = make(map[string]version.Version)
	m.app = make(map[string]version.AppSummary)
	m.host = make(map[string]version.HostSummary)
	m.threshold = time.Now()
}

// NewMemoryStore returns a empty MemoryStore.
func NewMemoryStore() (m *MemoryStore) {
	m = new(MemoryStore)
	initMemoryStore(m)
	return
}

// Apps returns summary information about each application, based on the known Versions.
func (s *MemoryStore) Apps() []version.AppSummary {
	vs := make([]version.AppSummary, 0)
	for _, v := range s.app {
		vs = append(vs, v)
	}
	return vs
}

// Hosts returns summary information about each host, based on the known Versions.
func (s *MemoryStore) Hosts() []version.HostSummary {
	vs := make([]version.HostSummary, 0)
	for _, v := range s.host {
		vs = append(vs, v)
	}
	return vs
}

// Versions returns full Version structs where their values match app_id, host
// and ver. Zero length strings are considered a match for all Versions.
func (s *MemoryStore) Versions(app_id string,
	host string, ver string) []version.Version {
	vs := make([]version.Version, 0)
	for _, v := range s.version {
		if (len(app_id) == 0 || app_id == v.AppId) &&
			(len(host) == 0 || host == v.Host) &&
			(len(ver) == 0 || ver == v.Ver) {
			vs = append(vs, v)
		}
	}
	return vs
}

// Update stores a Version and updates app and host summaries.
func (s *MemoryStore) Update(v version.Version) (err error) {
	key := v.Key()
	_, vpresent := s.version[key]
	if v.ExactUpdate.After(s.threshold) {
		s.version[key] = v
	}

	// app map
	as, present := s.app[v.AppId]
	if present {
		as.LastUpdate = v.LastUpdate
		if !vpresent {
			as.HostCount++
		}
	} else {
		appv := version.AppSummary{
			App:        v.App,
			AppId:      v.AppId,
			LastUpdate: v.LastUpdate,
			HostCount:  1,
		}
		s.app[v.AppId] = appv
	}

	// host map
	hostv := version.HostSummary{
		Host:       v.Host,
		LastUpdate: v.LastUpdate,
	}
	s.host[v.Host] = hostv
	return
}

// Clear empties the MemoryStore.
func (s *MemoryStore) Clear() {
	initMemoryStore(s)
}
