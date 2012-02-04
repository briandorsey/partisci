package memstore

import (
	"partisci/version"
)

type MemoryStore struct {
	Apps     map[string]version.Version
	Versions map[string]version.Version
}

func NewMemoryStore() (m *MemoryStore) {
	m = new(MemoryStore)
	m.Apps = make(map[string]version.Version)
	m.Versions = make(map[string]version.Version)
	return
}

func (s *MemoryStore) GetApps() (vs []version.Version) {
	for _, v := range s.Apps {
		vs = append(vs, v)
	}
	return
}

func (s *MemoryStore) Update(v version.Version) (err error) {
	_, ok := s.Apps[v.Id]
	if !ok {
		// store a simplified version in the app map
		appv := version.Version{Name: v.Name, Id: v.Id}
		s.Apps[v.Id] = appv
	}
	return
}
