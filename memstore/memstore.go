package memstore

import (
	"partisci/version"
)

type MemoryStore struct {
	App     map[string]version.Version
	Version map[string]version.Version
}

func NewMemoryStore() (m *MemoryStore) {
	m = new(MemoryStore)
	m.App = make(map[string]version.Version)
	m.Version = make(map[string]version.Version)
	return
}

func (s *MemoryStore) Apps() []version.Version {
	vs := make([]version.Version, 0)
	for _, v := range s.App {
		vs = append(vs, v)
	}
	return vs
}

func (s *MemoryStore) Update(v version.Version) (err error) {
	key := versionToKey(v)
	s.Version[key] = v

	// app map
	_, ok := s.App[v.AppId]
	if !ok {
		// store a simplified version in the app map
		appv := version.Version{
            App: v.App,
            AppId: v.AppId,
            LastUpdate: v.LastUpdate,
        }
		s.App[v.AppId] = appv
	}
	return
}

func versionToKey(v version.Version) string {
	return v.AppId + v.Host
}
