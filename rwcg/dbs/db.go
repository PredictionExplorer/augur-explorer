// Package dbs provides database storage for rwcg: connection, SQLStorage,
// and schema helpers. Domain-specific access is in subpackages randomwalk and cosmicgame.
package dbs

import (
	"context"
	"database/sql"
	"database/sql/driver"
	"fmt"
	"log"
	"net"
	"os"
	"strings"
	"time"

	"github.com/lib/pq"
)

// Connection resilience tuning. These mirror libpq's connect_timeout +
// keepalives_idle / keepalives_interval / keepalives_count / tcp_user_timeout,
// which the pure-Go lib/pq driver does NOT accept as connection-string keys, so
// we apply them ourselves via a custom dialer (see keepaliveDialer below).
const (
	dbConnectTimeout    = 10 * time.Second
	dbKeepaliveIdle     = 30 * time.Second
	dbKeepaliveInterval = 10 * time.Second
	dbKeepaliveCount    = 5
	dbTCPUserTimeout    = 15 * time.Second

	// Connect retry: a transient link failure (e.g. "no route to host" on a
	// flaky Wi-Fi link) should not propagate an error to callers that react by
	// calling os.Exit. We retry the dial several times with capped backoff so a
	// short blip is absorbed before the error can reach those call sites.
	dbConnectMaxAttempts = 10
	dbConnectRetryDelay  = 1 * time.Second
	dbConnectRetryMaxWait = 5 * time.Second
)

// keepaliveDialer dials with a bounded connect timeout and TCP keepalive probing
// (plus TCP_USER_TIMEOUT on Linux) so a dead or flaky link is detected in
// seconds instead of the kernel default of minutes/hours.
type keepaliveDialer struct{ d net.Dialer }

func newKeepaliveDialer() keepaliveDialer {
	return keepaliveDialer{d: net.Dialer{
		Timeout:   dbConnectTimeout,
		KeepAlive: dialerKeepAlive,    // platform-specific (see db_keepalive_*.go)
		Control:   tcpKeepaliveControl, // platform-specific socket-option tuning
	}}
}

func (k keepaliveDialer) Dial(network, address string) (net.Conn, error) {
	return k.d.Dial(network, address)
}

func (k keepaliveDialer) DialTimeout(network, address string, timeout time.Duration) (net.Conn, error) {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	return k.d.DialContext(ctx, network, address)
}

func (k keepaliveDialer) DialContext(ctx context.Context, network, address string) (net.Conn, error) {
	return k.d.DialContext(ctx, network, address)
}

// keepaliveConnector is a database/sql driver.Connector that opens lib/pq
// connections through keepaliveDialer. lib/pq v1.10.4's Connector keeps its
// dialer unexported with no setter, so we wrap pq.DialOpen directly.
type keepaliveConnector struct{ dsn string }

func (c keepaliveConnector) Connect(ctx context.Context) (driver.Conn, error) {
	delay := dbConnectRetryDelay
	var lastErr error
	for attempt := 1; attempt <= dbConnectMaxAttempts; attempt++ {
		conn, err := pq.DialOpen(newKeepaliveDialer(), c.dsn)
		if err == nil {
			if attempt > 1 {
				log.Printf("DB: reconnected on attempt %d/%d", attempt, dbConnectMaxAttempts)
			}
			return conn, nil
		}
		lastErr = err
		log.Printf("DB: connect attempt %d/%d failed: %v", attempt, dbConnectMaxAttempts, err)
		if attempt == dbConnectMaxAttempts {
			break
		}
		// Back off, but bail out early if the caller's context is cancelled.
		select {
		case <-ctx.Done():
			return nil, ctx.Err()
		case <-time.After(delay):
		}
		if delay *= 2; delay > dbConnectRetryMaxWait {
			delay = dbConnectRetryMaxWait
		}
	}
	return nil, fmt.Errorf("db connect failed after %d attempts: %w", dbConnectMaxAttempts, lastErr)
}

func (c keepaliveConnector) Driver() driver.Driver { return pq.Driver{} }

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
	// connect_timeout makes lib/pq call our dialer's DialTimeout, so a routing
	// blip fails fast instead of hanging. Keepalive + TCP_USER_TIMEOUT are applied
	// at the socket level by keepaliveDialer (lib/pq ignores those DSN keys).
	connStr += fmt.Sprintf(" connect_timeout=%d", int(dbConnectTimeout.Seconds()))

	db := sql.OpenDB(keepaliveConnector{dsn: connStr})
	// First statement forces a real connection, surfacing connect errors here
	// (same early-failure behavior as the previous sql.Open + Exec).
	if _, err := db.Exec("SET timezone TO 0"); err != nil {
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
	return NewSQLStorageFromDB(db, info_log)
}

// NewSQLStorageFromDB wraps an existing database handle (e.g. tools using -db DSN).
func NewSQLStorageFromDB(db *sql.DB, info_log *log.Logger) *SQLStorage {
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
