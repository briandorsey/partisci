// Package store defines the UpdateStore interface for version persistence.
package store

import (
	"github.com/briandorsey/partisci/version"
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
	Apps() (as []version.AppSummary, err error)

	// Host returns a HostSummary for the given Host.
	// The value of ok follows map indexing conventions:
	//   true if Host is present, false otherwise.
	Host(Host string) (hs version.HostSummary, ok bool)

	// Hosts returns summary information about each host,
	// based on the known Versions.
	Hosts() (hs []version.HostSummary, err error)

	// Versions returns full Version structs where their values match app_id, host
	// and ver. Zero length strings are considered a match for all Versions.
	Versions(AppId string, Host string, Ver string) (
		vs []version.Version, err error)

	// Clear empties the MemoryStore.
	Clear() (err error)

	// Trim removes old versions.
	Trim(t time.Time) (c uint64, err error)
}
