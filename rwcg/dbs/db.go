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
// If PGSQL_HOST is empty or unset, connects via Unix socket (same as psql without -h); otherwise uses TCP.
// Caller must not use returned db if err != nil.
func openDB(schema string) (*sql.DB, error) {
	hostEnv := os.Getenv("PGSQL_HOST")
	useSocket := hostEnv == ""
	connStr := "user='" + escapeConnParam(os.Getenv("PGSQL_USERNAME")) +
		"' dbname='" + escapeConnParam(os.Getenv("PGSQL_DATABASE")) + "'"
	if !useSocket {
		connStr += " password='" + escapeConnParam(os.Getenv("PGSQL_PASSWORD")) + "'"
	}
	if hostEnv != "" {
		host, port, err := net.SplitHostPort(hostEnv)
		if err != nil {
			host = hostEnv
			port = "5432"
		}
		connStr += " host='" + escapeConnParam(host) + "' port='" + escapeConnParam(port) + "'"
	}
	// When useSocket: no host/port and no password → Unix socket + trust/peer auth (like psql -U user)
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

func show_connect_error(err error) {
	fmt.Println("Can't connect to PostgreSQL database.")
	if err != nil {
		fmt.Printf("Connection error: %v\n", err)
		if strings.Contains(err.Error(), "password authentication failed") {
			fmt.Println("Hint: PGSQL_PASSWORD does not match the password set in PostgreSQL for this user.")
			fmt.Println("  Option A – Make PostgreSQL use your env password:")
			fmt.Println("    psql -U postgres -c \"ALTER USER cosmicgame PASSWORD 'YOUR_ENV_PASSWORD';\"")
			fmt.Println("  Option B – Discover the correct password and set it in your env:")
			fmt.Println("    psql -h 127.0.0.1 -U cosmicgame -d cosmicgame -W   # type the working password, then set that in PGSQL_PASSWORD")
			fmt.Println("  Option C – Use Unix socket (no password) if pg_hba allows trust/peer for local:")
			fmt.Println("    unset PGSQL_HOST")
		}
	}
	fmt.Println("Environment variable status:")
	for _, name := range []string{"PGSQL_USERNAME", "PGSQL_PASSWORD", "PGSQL_DATABASE", "PGSQL_HOST"} {
		v := os.Getenv(name)
		if v == "" {
			fmt.Printf("  %s: not set (or empty)\n", name)
		} else {
			if name == "PGSQL_PASSWORD" {
				fmt.Printf("  %s: set (length %d)\n", name, len(v))
			} else {
				fmt.Printf("  %s: set = %q\n", name, v)
			}
		}
	}
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
		show_connect_error(err)
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
