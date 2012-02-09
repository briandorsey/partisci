package version

import (
	"encoding/json"
	"strings"
	"time"
)

type Version struct {
	AppId       string    `json:"app_id,omitempty"`
	App         string    `json:"app,omitempty"`
	Ver         string    `json:"ver,omitempty"`
	Host        string    `json:"host,omitempty"`
	Instance    uint16    `json:"instance,omitempty"`
	HostIP      string    `json:"host_ip,omitempty"`
	LastUpdate  int64     `json:"last_update,omitempty"`
	ExactUpdate time.Time `json:"-"`
}

type AppSummary struct {
	AppId      string `json:"app_id"`
	App        string `json:"app"`
	LastUpdate int64  `json:"last_update"`
	HostCount  int32  `json:"host_count"`
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
	v.HostIP = host
	v.ExactUpdate = time.Now()
	v.LastUpdate = v.ExactUpdate.Unix()
	err = json.Unmarshal(b[:len(b)], &v)
	if err != nil {
		return v, err
	}
	v.AppId = AppIdToId(v.App)
	return
}
