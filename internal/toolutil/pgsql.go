package toolutil

import (
	"database/sql"
	"fmt"
	"net/url"
	"os"
	"strings"

	_ "github.com/lib/pq"
)

// OpenPostgresFromEnv connects using DATABASE_URL when set, otherwise
// PGSQL_HOST, PGSQL_USERNAME, PGSQL_DATABASE, PGSQL_PASSWORD.
// Host may be host:port (e.g. 69.10.55.2:5432).
func OpenPostgresFromEnv() (*sql.DB, error) {
	conn, err := PostgresConnStringFromEnv()
	if err != nil {
		return nil, err
	}
	return sql.Open("postgres", conn)
}

// PostgresConnStringFromEnv returns DATABASE_URL when set, otherwise a
// postgres URL built from the standard PGSQL_* variables.
func PostgresConnStringFromEnv() (string, error) {
	if u := strings.TrimSpace(os.Getenv("DATABASE_URL")); u != "" {
		return u, nil
	}
	host := os.Getenv("PGSQL_HOST")
	user := os.Getenv("PGSQL_USERNAME")
	dbName := os.Getenv("PGSQL_DATABASE")
	pass := os.Getenv("PGSQL_PASSWORD")
	if host == "" || user == "" || dbName == "" {
		return "", fmt.Errorf("DATABASE_URL, or PGSQL_HOST, PGSQL_USERNAME, and PGSQL_DATABASE must be set (or pass -db)")
	}
	u := &url.URL{
		Scheme: "postgres",
		User:   url.UserPassword(user, pass),
		Host:   host,
		Path:   dbName,
	}
	q := u.Query()
	q.Set("sslmode", "disable")
	u.RawQuery = q.Encode()
	return u.String(), nil
}
