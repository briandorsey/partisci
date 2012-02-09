package memstore

import (
	"partisci/version"
	"time"
)

type MemoryStore struct {
	Version   map[string]version.Version
	App       map[string]version.AppSummary
	Host      map[string]version.Version
	threshold time.Time
}

func initMemoryStore(m *MemoryStore) {
	m.Version = make(map[string]version.Version)
	m.App = make(map[string]version.AppSummary)
	m.Host = make(map[string]version.Version)
	m.threshold = time.Now()
}

func NewMemoryStore() (m *MemoryStore) {
	m = new(MemoryStore)
	initMemoryStore(m)
	return
}

func (s *MemoryStore) Apps() []version.AppSummary {
	vs := make([]version.AppSummary, 0)
	for _, v := range s.App {
		vs = append(vs, v)
	}
	return vs
}

func (s *MemoryStore) Hosts() []version.Version {
	vs := make([]version.Version, 0)
	for _, v := range s.Host {
		vs = append(vs, v)
	}
	return vs
}

func (s *MemoryStore) Versions(app_id string,
	host string, ver string) []version.Version {
	vs := make([]version.Version, 0)
	for _, v := range s.Version {
		if (len(app_id) == 0 || app_id == v.AppId) &&
			(len(host) == 0 || host == v.Host) &&
			(len(ver) == 0 || ver == v.Ver) {
			vs = append(vs, v)
		}
	}
	return vs
}

func (s *MemoryStore) Update(v version.Version) (err error) {
	key := versionToKey(v)
	_, vpresent := s.Version[key]
	if v.ExactUpdate.After(s.threshold) {
		s.Version[key] = v
	}

	// app map
	as, present := s.App[v.AppId]
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
		s.App[v.AppId] = appv
	}

	// host map
	hostv := version.Version{
		Host:       v.Host,
		LastUpdate: v.LastUpdate,
	}
	s.Host[v.Host] = hostv
	return
}

func (s *MemoryStore) Clear() {
	initMemoryStore(s)
}

func versionToKey(v version.Version) string {
	return v.AppId + v.Host
}
