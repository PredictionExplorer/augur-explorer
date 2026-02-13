// Package dbs provides database storage for rwcg: connection, SQLStorage,
// and schema helpers. Domain-specific access is in subpackages randomwalk and cosmicgame.
package dbs

import (
	"database/sql"
	"fmt"
	"log"
	"net"
	"os"
	"strings"

	_ "github.com/lib/pq"
)

// escapeConnParam escapes a value for use inside a single-quoted libpq connection parameter.
func escapeConnParam(s string) string {
	s = strings.ReplaceAll(s, "\\", "\\\\")
	s = strings.ReplaceAll(s, "'", "\\'")
	return s
}

// openDB opens a PostgreSQL connection using env vars (PGSQL_*).
// If schema != "", adds search_path to the connection string.
// Caller must not use returned db if err != nil.
func openDB(schema string) (*sql.DB, error) {
	host, port, err := net.SplitHostPort(os.Getenv("PGSQL_HOST"))
	if err != nil {
		host = os.Getenv("PGSQL_HOST")
		port = "5432"
	}
	connStr := "user='" + escapeConnParam(os.Getenv("PGSQL_USERNAME")) +
		"' dbname='" + escapeConnParam(os.Getenv("PGSQL_DATABASE")) +
		"' password='" + escapeConnParam(os.Getenv("PGSQL_PASSWORD")) +
		"' host='" + escapeConnParam(host) +
		"' port='" + escapeConnParam(port) + "'"
	if schema != "" {
		connStr += " search_path='" + escapeConnParam(schema) + "'"
	}
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, err
	}
	_, err = db.Exec("SET timezone TO 0")
	if err != nil {
		db.Close()
		return nil, err
	}
	return db, nil
}

func show_connect_error() {
	fmt.Printf(`Websrv: can't connect to PostgreSQL database.
				Check that you have set PGSQL_USERNAME,PGSQL_PASSWORD,PGSQL_DATABASE
				PGSQL_HOST environment variables`)
}

type SQLStorage struct {
	db         *sql.DB
	db_logger  *log.Logger
	Info       *log.Logger
	schema_name string
}

func (ss *SQLStorage) SchemaName() string { return ss.schema_name }
func (ss *SQLStorage) Db() *sql.DB        { return ss.db }

// Connect_to_storage connects using PGSQL_* env vars and returns SQLStorage.
// On connection failure prints an error and returns nil; callers should check for nil.
func Connect_to_storage(info_log *log.Logger) *SQLStorage {
	db, err := openDB("")
	if err != nil {
		show_connect_error()
		return nil
	}
	ss := new(SQLStorage)
	ss.db = db
	ss.Info = info_log
	return ss
}

func (ss *SQLStorage) Init_log(fname string) {
	f, err := os.OpenFile(fname, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Printf("Exiting extractor with error: %v", err)
		os.Exit(1)
	}
	ss.db_logger = log.New(f, "DB: ", log.LstdFlags)
}

func (ss *SQLStorage) Log_msg(msg string) {
	if ss.db_logger != nil {
		ss.db_logger.Printf(msg)
	} else {
		ss.Info.Printf(msg)
	}
}

func (ss *SQLStorage) Db_set_schema_name(name string) {
	ss.schema_name = name
}
