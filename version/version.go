// Package version provides types and functions for manipulating version data.
package version

import (
	"encoding/json"
	"errors"
	"strings"
	"time"
)

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

func NewVersion() (v *Version) {
	v = new(Version)
	v.ExactUpdate = time.Now()
	v.LastUpdate = v.ExactUpdate.Unix()
	return
}

func (v *Version) Key() string {
	return v.AppId + v.Host + string(v.Instance)
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

func AppIdToId(app string) (id string) {
	id = strings.ToLower(app)
	id = strings.Map(safeRunes, id)
	return
}

func ParsePacket(host string, b []byte) (v Version, err error) {
	v = *NewVersion()
	v.HostIP = host
	err = json.Unmarshal(b[:len(b)], &v)
	if err != nil {
		return
	}
	v.AppId = AppIdToId(v.App)

	// ensure minimal values were given
	if len(v.App) == 0 ||
		len(v.Ver) == 0 {
		err = errors.New("value for app & ver must be specified")
		return
	}
	return
}
