// Package store provides database access for rwcg: the pgxpool-backed Store
// (see store.go) plus the legacy SQLStorage handle that unconverted query
// methods still hang off. Domain-specific access is in subpackages randomwalk
// and cosmicgame.
package store

import (
	"database/sql"
	"fmt"
	"log"
	"os"
)

// SQLStorage is the legacy database handle: methods on it predate the Phase 1
// store rewrite (no context, errors logged or fatal). It shrinks as files are
// converted onto Store/pgx and is deleted when the conversion completes.
type SQLStorage struct {
	db          *sql.DB
	db_logger   *log.Logger
	Info        *log.Logger
	schema_name string
}

func (ss *SQLStorage) SchemaName() string { return ss.schema_name }
func (ss *SQLStorage) Db() *sql.DB        { return ss.db }

// NewSQLStorageFromDB wraps an existing database handle: the pool view from
// Store.DB(), or a separately opened pool for tools using -db DSN flags.
func NewSQLStorageFromDB(db *sql.DB, info_log *log.Logger) *SQLStorage {
	ss := new(SQLStorage)
	ss.db = db
	ss.Info = info_log
	return ss
}

// Init_log opens fname for appending and directs subsequent DB log output
// (Log_msg) there. It returns an error if the log file can't be opened.
func (ss *SQLStorage) Init_log(fname string) error {
	f, err := os.OpenFile(fname, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return fmt.Errorf("open db log file %v: %w", fname, err)
	}
	ss.db_logger = log.New(f, "DB: ", log.LstdFlags)
	return nil
}

func (ss *SQLStorage) Log_msg(msg string) {
	if ss.db_logger != nil {
		ss.db_logger.Print(msg)
	} else {
		ss.Info.Print(msg)
	}
}

func (ss *SQLStorage) Db_set_schema_name(name string) {
	ss.schema_name = name
}
