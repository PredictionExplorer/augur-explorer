package toolutil

import (
	"database/sql"
	"fmt"
	"net/url"
	"os"

	_ "github.com/lib/pq"
)

// OpenPostgresFromEnv connects using PGSQL_HOST, PGSQL_USERNAME, PGSQL_DATABASE, PGSQL_PASSWORD.
// Host may be host:port (e.g. 69.10.55.2:5432).
func OpenPostgresFromEnv() (*sql.DB, error) {
	conn, err := PostgresConnStringFromEnv()
	if err != nil {
		return nil, err
	}
	return sql.Open("postgres", conn)
}

// PostgresConnStringFromEnv builds a postgres URL from standard cg-prod.env variables.
func PostgresConnStringFromEnv() (string, error) {
	host := os.Getenv("PGSQL_HOST")
	user := os.Getenv("PGSQL_USERNAME")
	dbName := os.Getenv("PGSQL_DATABASE")
	pass := os.Getenv("PGSQL_PASSWORD")
	if host == "" || user == "" || dbName == "" {
		return "", fmt.Errorf("PGSQL_HOST, PGSQL_USERNAME, and PGSQL_DATABASE must be set (or pass -db)")
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
