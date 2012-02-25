// Package version provides types and functions for manipulating version data.
package version

import (
	"encoding/json"
	"errors"
	"strings"
	"time"
)

// UpdateStore defines an interface for persisting application version information.
type UpdateStore interface {
	Update(v Version) (err error)
	Apps() (vs []AppSummary)
	Hosts() (vs []HostSummary)
	Versions(app_id string, host string, ver string) (vs []Version)
	Clear()
	Trim(t time.Time) (c uint64)
}

type Version struct {
	AppId       string    `json:"app_id,omitempty"`
	App         string    `json:"app"`
	Ver         string    `json:"ver"`
	Host        string    `json:"host"`
	Instance    uint16    `json:"instance"`
	HostIP      string    `json:"host_ip,omitempty"`
	LastUpdate  int64     `json:"last_update,omitempty"`
	ExactUpdate time.Time `json:"-"`
}

// Key returns a suitable unique id for storing in a database.
// It is calculated using AppId, Host & Instance so later version changes
// will result in updates.
func (v *Version) Key() string {
	return v.AppId + v.Host + string(v.Instance)
}

// Prepare readies a Version for use by calculating fields.
// Prepare *must* be called after populating fields and before passing to a store.
func (v *Version) Prepare() {
	v.AppId = appIdToId(v.App)
	if v.LastUpdate == 0 {
		v.ExactUpdate = time.Now()
		v.LastUpdate = v.ExactUpdate.Unix()
	}
}

type AppSummary struct {
	AppId      string `json:"app_id"`
	App        string `json:"app"`
	LastUpdate int64  `json:"last_update"`
	HostCount  int32  `json:"host_count"`
}

type HostSummary struct {
	Host       string `json:"host"`
	LastUpdate int64  `json:"last_update"`
	AppCount   int32  `json:"app_count"`
}

func safeRunes(r rune) rune {
	if '0' <= r && r <= '9' {
		return r
	}
	if 'a' <= r && r <= 'z' {
		return r
	}
	return '_'
}

func appIdToId(app string) (id string) {
	id = strings.ToLower(app)
	id = strings.Map(safeRunes, id)
	return
}

func ParsePacket(host string, b []byte) (v Version, err error) {
	v = *new(Version)
	v.HostIP = host
	err = json.Unmarshal(b[:len(b)], &v)
	if err != nil {
		return
	}
	v.Prepare()

	// ensure minimal values were given
	if len(v.App) == 0 ||
		len(v.Ver) == 0 {
		err = errors.New("value for app & ver must be specified")
		return
	}
	return
}
