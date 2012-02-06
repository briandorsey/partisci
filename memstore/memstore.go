package memstore

import (
	"partisci/version"
)

type MemoryStore struct {
	Version map[string]version.Version
	App     map[string]version.Version
	Host    map[string]version.Version
}

func initMemoryStore(m *MemoryStore) {
	m.Version = make(map[string]version.Version)
	m.App = make(map[string]version.Version)
	m.Host = make(map[string]version.Version)
}

func NewMemoryStore() (m *MemoryStore) {
	m = new(MemoryStore)
	initMemoryStore(m)
	return
}

func (s *MemoryStore) Apps() []version.Version {
	vs := make([]version.Version, 0)
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

func (s *MemoryStore) Versions() []version.Version {
	vs := make([]version.Version, 0)
	for _, v := range s.Version {
		vs = append(vs, v)
	}
	return vs
}

func (s *MemoryStore) Update(v version.Version) (err error) {
	key := versionToKey(v)
	s.Version[key] = v

	// app map
	appv := version.Version{
		App:        v.App,
		AppId:      v.AppId,
		LastUpdate: v.LastUpdate,
	}
	s.App[v.AppId] = appv

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
