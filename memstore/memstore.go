// Package memstore is a memory-backed implementation of the UpdateStore interface.
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

// NewMemoryStore returns a new, initialized MemoryStore.
func NewMemoryStore() (m *MemoryStore) {
	m = new(MemoryStore)
	initMemoryStore(m)
	return
}

func (s *MemoryStore) App(AppId string) (as version.AppSummary, ok bool) {
	as, ok = s.app[AppId]
	return
}

func (s *MemoryStore) Apps() []version.AppSummary {
	vs := make([]version.AppSummary, 0)
	for _, v := range s.app {
		vs = append(vs, v)
	}
	return vs
}

func (s *MemoryStore) Host(Host string) (hs version.HostSummary, ok bool) {
	hs, ok = s.host[Host]
	return
}

func (s *MemoryStore) Hosts() []version.HostSummary {
	vs := make([]version.HostSummary, 0)
	for _, v := range s.host {
		vs = append(vs, v)
	}
	return vs
}

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

func (s *MemoryStore) Update(v version.Version) (err error) {
	key := v.Key()
	_, vpresent := s.version[key]
	if v.ExactUpdate.After(s.threshold) {
		s.version[key] = v
	}

	// app map
	if as, present := s.app[v.AppId]; present {
		as.LastUpdate = v.LastUpdate
		if !vpresent {
			as.HostCount++
			s.app[v.AppId] = as
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
	if hs, present := s.host[v.Host]; present {
		hs.LastUpdate = v.LastUpdate
		if !vpresent {
			hs.AppCount++
			s.host[v.Host] = hs
		}
	} else {
		hostv := version.HostSummary{
			Host:       v.Host,
			LastUpdate: v.LastUpdate,
			AppCount:   1,
		}
		s.host[v.Host] = hostv
	}
	return
}

func (s *MemoryStore) Clear() {
	initMemoryStore(s)
}

func (s *MemoryStore) Trim(t time.Time) (c uint64) {
	s.threshold = t
	for k, v := range s.version {
		if v.ExactUpdate.Before(t) {
			c++
			delete(s.version, k)
			if as, ok := s.app[v.AppId]; ok {
				as.HostCount -= 1
				s.app[v.AppId] = as
				if as.HostCount < 1 {
					delete(s.app, v.AppId)
				}
			}
			if hs, ok := s.host[v.Host]; ok {
				hs.AppCount -= 1
				s.host[v.Host] = hs
				if hs.AppCount < 1 {
					delete(s.host, v.Host)
				}
			}
		}
	}
	return
}
