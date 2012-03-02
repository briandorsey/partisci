// Package sqlitestore is a SQLite-backed implementation of the UpdateStore
// interface.
package sqlitestore

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
	"partisci/version"
	"time"
)

// SQLiteStore is an a SQLite-backed implementation of the UpdateStore interface.
type SQLiteStore struct {
	db   *sql.DB
	path string
}

// NewSQLiteStore returns a new, initialized SQLiteStore.
func NewSQLiteStore(path string) (s *SQLiteStore, err error) {
	s = new(SQLiteStore)
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
        group by app_id`, AppId)
	err := row.Scan(&as.AppId, &as.App, &as.LastUpdate, &as.HostCount)
	if err != nil {
		// TODO: should App() & Host() return errors instead of ok?
		return as, false
	}
	return as, true
}

func (s *SQLiteStore) Apps() []version.AppSummary {
	as := make([]version.AppSummary, 0)
	return as
}

func (s *SQLiteStore) Host(Host string) (hs version.HostSummary, ok bool) {
	row := s.db.QueryRow(`
        select host, max(last_update), count(app)
        from version
        where host = ?
        group by host`, Host)
	err := row.Scan(&hs.Host, &hs.LastUpdate, &hs.AppCount)
	if err != nil {
		// TODO: should App() & Host() return errors instead of ok?
		return hs, false
	}
	return hs, true
}

func (s *SQLiteStore) Hosts() []version.HostSummary {
	vs := make([]version.HostSummary, 0)
	return vs
}

func (s *SQLiteStore) Update(v version.Version) (err error) {
	tx, err := s.db.Begin()
	if err != nil {
		return
	}
	stmt, err := tx.Prepare(
		`insert into version(key, app_id, app, ver, host, 
                    instance, host_ip, last_update, exact_update)
        values(?, ?, ?, ?, ?, ?, ?, ?, ?)`)
	if err != nil {
		return
	}
	defer stmt.Close()

	_, err = stmt.Exec(v.Key(), v.AppId, v.App, v.Ver, v.Host,
		v.Instance, v.HostIP, v.LastUpdate, v.ExactUpdate.Format(time.RFC3339))
	if err != nil {
		return
	}
	err = tx.Commit()
	if err != nil {
		return
	}
	return
}

func (s *SQLiteStore) Versions(app_id string,
	host string, ver string) []version.Version {
	v := make([]version.Version, 0)
	return v
}

func (s *SQLiteStore) Clear() {
}

func (s *SQLiteStore) Trim(t time.Time) (c uint64) {
	c = 0
	return
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
