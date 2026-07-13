package toolutil

import (
	"net/url"
	"strings"
	"testing"
)

func TestPostgresConnStringPrefersDatabaseURL(t *testing.T) {
	t.Setenv("DATABASE_URL", " postgres://u:p@url.example.com/urldb ")
	t.Setenv("PGSQL_HOST", "ignored.example.com")
	t.Setenv("PGSQL_USERNAME", "ignored")
	t.Setenv("PGSQL_DATABASE", "ignored")
	connection, err := PostgresConnStringFromEnv()
	if err != nil {
		t.Fatal(err)
	}
	if connection != "postgres://u:p@url.example.com/urldb" {
		t.Fatalf("connection = %q, want the trimmed DATABASE_URL", connection)
	}
}

func TestPostgresConnStringFromEnv(t *testing.T) {
	t.Setenv("DATABASE_URL", "")
	t.Setenv("PGSQL_HOST", "db.example.com:5433")
	t.Setenv("PGSQL_USERNAME", "alice@example.com")
	t.Setenv("PGSQL_DATABASE", "game db")
	t.Setenv("PGSQL_PASSWORD", "p@ss:/word")
	connection, err := PostgresConnStringFromEnv()
	if err != nil {
		t.Fatal(err)
	}
	parsed, err := url.Parse(connection)
	if err != nil {
		t.Fatal(err)
	}
	password, _ := parsed.User.Password()
	if parsed.Scheme != "postgres" ||
		parsed.Host != "db.example.com:5433" ||
		parsed.User.Username() != "alice@example.com" ||
		password != "p@ss:/word" ||
		parsed.Path != "/game db" ||
		parsed.Query().Get("sslmode") != "disable" {
		t.Fatalf("connection URL = %s", connection)
	}
	database, err := OpenPostgresFromEnv()
	if err != nil {
		t.Fatal(err)
	}
	if err := database.Close(); err != nil {
		t.Fatal(err)
	}
}

func TestPostgresConnStringRequiresCoreFields(t *testing.T) {
	for _, missing := range []string{"PGSQL_HOST", "PGSQL_USERNAME", "PGSQL_DATABASE"} {
		t.Run(missing, func(t *testing.T) {
			t.Setenv("DATABASE_URL", "")
			t.Setenv("PGSQL_HOST", "host")
			t.Setenv("PGSQL_USERNAME", "user")
			t.Setenv("PGSQL_DATABASE", "db")
			t.Setenv(missing, "")
			if _, err := PostgresConnStringFromEnv(); err == nil ||
				!strings.Contains(err.Error(), "must be set") {
				t.Fatalf("error = %v", err)
			}
		})
	}
}
