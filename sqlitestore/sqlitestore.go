// Package sqlitestore is a SQLite-backed implementation of the UpdateStore
// interface.
package sqlitestore

import (
	"database/sql"
	"errors"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
	"github.com/briandorsey/partisci/version"
	"time"
)

// SQLiteStore is an a SQLite-backed implementation of the UpdateStore interface.
type SQLiteStore struct {
	db        *sql.DB
	path      string
	threshold time.Time
}

// NewSQLiteStore returns a new, initialized SQLiteStore.
func NewSQLiteStore(path string) (s *SQLiteStore, err error) {
	s = new(SQLiteStore)
	s.threshold = time.Now()

	db, err := sql.Open("sqlite3", path)
	if err != nil {
		return nil, err
	}
	s.db = db
	err = s.checkInit()
	if err != nil {
		return nil, err
	}
	return s, nil
}

func (s *SQLiteStore) checkInit() (err error) {
	for _, sql := range sqls {
		_, err = s.db.Exec(sql)
		if err != nil {
			return err
		}
	}
	return nil
}

func (s *SQLiteStore) Close() (err error) {
	err = s.db.Close()
	return
}

func (s *SQLiteStore) App(AppId string) (as version.AppSummary, ok bool) {
	row := s.db.QueryRow(`
        select app_id, max(app), max(last_update), count(host)
        from version
        where app_id = ?
        group by app_id;`, AppId)
	err := row.Scan(&as.AppId, &as.App, &as.LastUpdate, &as.HostCount)
	if err != nil {
		// TODO: should App() & Host() return errors instead of ok?
		return as, false
	}
	return as, true
}

func (s *SQLiteStore) Apps() (as []version.AppSummary, err error) {
	as = make([]version.AppSummary, 0)
	rows, err := s.db.Query(`
        select app_id, max(app), max(last_update), count(host)
        from version
        group by app_id;`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		a := version.AppSummary{}
		err = rows.Scan(&a.AppId, &a.App, &a.LastUpdate, &a.HostCount)
		if err != nil {
			return nil, err
		}
		as = append(as, a)
	}
	return as, nil
}

func (s *SQLiteStore) Host(Host string) (hs version.HostSummary, ok bool) {
	row := s.db.QueryRow(`
        select host, max(last_update), count(app_id)
        from version
        where host = ?
        group by host;`, Host)
	err := row.Scan(&hs.Host, &hs.LastUpdate, &hs.AppCount)
	if err != nil {
		// TODO: should App() & Host() return errors instead of ok?
		return hs, false
	}
	return hs, true
}

func (s *SQLiteStore) Hosts() (hs []version.HostSummary, err error) {
	hs = make([]version.HostSummary, 0)
	rows, err := s.db.Query(`
        select host, max(last_update), count(app_id)
        from version
        group by host;`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		h := version.HostSummary{}
		err = rows.Scan(&h.Host, &h.LastUpdate, &h.AppCount)
		if err != nil {
			return nil, err
		}
		hs = append(hs, h)
	}
	return hs, nil
}

func (s *SQLiteStore) Update(v version.Version) (err error) {
	if v.ExactUpdate.Before(s.threshold) {
		return
	}
	_, err = s.db.Exec(
		`insert into version(key, app_id, app, ver, host, 
                    instance, host_ip, last_update, exact_update)
        values(?, ?, ?, ?, ?, ?, ?, ?, ?);`,
		v.Key(), v.AppId, v.App, v.Ver, v.Host,
		v.Instance, v.HostIP, v.LastUpdate, v.ExactUpdate.UnixNano())
	return err
}

func (s *SQLiteStore) Versions(AppId string, Host string, Ver string) (
	vs []version.Version, err error) {
	vs = make([]version.Version, 0)
	if AppId == "" {
		AppId = "%"
	}
	if Host == "" {
		Host = "%"
	}
	if Ver == "" {
		Ver = "%"
	}
	rows, err := s.db.Query(`
        select app_id, app, ver, host, 
            instance, host_ip, last_update, exact_update
        from version
        where app_id like ?
            and host like ?
            and ver like ?;`,
		AppId, Host, Ver)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		v := version.Version{}
		var d int64
		err = rows.Scan(&v.AppId, &v.App, &v.Ver, &v.Host,
			&v.Instance, &v.HostIP, &v.LastUpdate, &d)
		if err != nil {
			return nil, err
		}
		v.ExactUpdate = time.Unix(0, d)
		vs = append(vs, v)
	}
	return vs, nil
}

func (s *SQLiteStore) Clear() (err error) {
	_, err = s.db.Exec(`delete from version;`)
	if err != nil {
		return err
	}
	s.threshold = time.Now()
	return nil
}

func (s *SQLiteStore) Trim(t time.Time) (c uint64, err error) {
	c = 0
	un := t.UnixNano()
	r, err := s.db.Exec(`delete from version where exact_update < ?;`, un)
	if err != nil {
		return 0, err
	}
	ra, err := r.RowsAffected()
	if err != nil {
		return 0, nil
	}
	if ra >= 0 {
		c = uint64(ra)
	} else {
		return 0, errors.New(fmt.Sprint(
			"negative value returned from RowsAffected(): ", ra))
	}
	return c, nil
}

var sqls = []string{
	`create table if not exists version 
        (key varchar not null on conflict replace primary key,
        app_id varchar not null,
        app varchar not null,
        ver varchar not null,
        host varchar not null,
        instance int not null,
        host_ip varchar not null,
        last_update int not null,
        exact_update datetime not null);`,
	`create index if not exists version_app_id on version (app_id);`,
	`create index if not exists version_ver on version (ver);`,
	`create index if not exists version_host on version (host);`,
}
