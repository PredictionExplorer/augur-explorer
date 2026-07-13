package store

import (
	"math"
	"strings"
	"testing"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
)

func TestConfigFromEnv(t *testing.T) {
	t.Setenv("DATABASE_URL", "")
	t.Setenv("PGSQL_USERNAME", "alice")
	t.Setenv("PGSQL_PASSWORD", "s3cret")
	t.Setenv("PGSQL_DATABASE", "gamedb")
	t.Setenv("PGSQL_HOST", "db.example.com:5433")

	cfg := ConfigFromEnv()
	want := Config{User: "alice", Password: "s3cret", Database: "gamedb", Host: "db.example.com:5433"}
	if cfg != want {
		t.Errorf("ConfigFromEnv() = %+v, want %+v", cfg, want)
	}
}

func TestConfigFromEnvDatabaseURL(t *testing.T) {
	t.Setenv("DATABASE_URL", " postgres://u:p@db.example.com:5433/gamedb ")
	t.Setenv("PGSQL_USERNAME", "alice")
	t.Setenv("PGSQL_PASSWORD", "s3cret")
	t.Setenv("PGSQL_DATABASE", "gamedb")
	t.Setenv("PGSQL_HOST", "db.example.com:5433")

	cfg := ConfigFromEnv()
	if cfg.URL != "postgres://u:p@db.example.com:5433/gamedb" { //nolint:gosec // G101: fake test credentials
		t.Errorf("URL = %q, want the trimmed DATABASE_URL", cfg.URL)
	}
	// The legacy fields are still read (diagnostics), but poolConfig uses
	// the URL when present.
	if cfg.User != "alice" {
		t.Errorf("User = %q, want alice", cfg.User)
	}
}

func TestPoolConfigPrefersURL(t *testing.T) {
	t.Parallel()
	cfg := Config{ //nolint:gosec // G101: fake test credentials
		URL:      "postgres://urluser:urlpass@url.example.com:6543/urldb",
		User:     "ignored",
		Password: "ignored",
		Database: "ignored",
		Host:     "ignored.example.com:1",
	}
	poolCfg, err := cfg.poolConfig()
	if err != nil {
		t.Fatalf("poolConfig: %v", err)
	}
	cc := poolCfg.ConnConfig
	if cc.Host != "url.example.com" || cc.Port != 6543 || cc.User != "urluser" || cc.Database != "urldb" {
		t.Errorf("URL was not preferred: host=%q port=%d user=%q db=%q", cc.Host, cc.Port, cc.User, cc.Database)
	}
	if cc.ConnectTimeout != dbConnectTimeout {
		t.Errorf("ConnectTimeout = %v, want the store default %v", cc.ConnectTimeout, dbConnectTimeout)
	}
	// The store invariants apply to URL configs too.
	if cc.RuntimeParams["timezone"] != "UTC" || cc.RuntimeParams["search_path"] != "public" {
		t.Errorf("runtime params not pinned: %v", cc.RuntimeParams)
	}
	if poolCfg.MaxConns != DefaultMaxConns {
		t.Errorf("MaxConns = %d, want %d", poolCfg.MaxConns, DefaultMaxConns)
	}
	if cc.DialFunc == nil {
		t.Error("keepalive dialer not applied to URL config")
	}
}

func TestPoolConfigURLKeepsExplicitConnectTimeout(t *testing.T) {
	t.Parallel()
	cfg := Config{URL: "postgres://u:p@h:5432/d?connect_timeout=3"} //nolint:gosec // G101: fake test credentials
	poolCfg, err := cfg.poolConfig()
	if err != nil {
		t.Fatalf("poolConfig: %v", err)
	}
	if got := poolCfg.ConnConfig.ConnectTimeout; got != 3*time.Second {
		t.Errorf("ConnectTimeout = %v, want the URL's 3s", got)
	}
}

func TestPoolConfigURLParseErrorDoesNotEchoSecret(t *testing.T) {
	t.Parallel()
	cfg := Config{URL: "postgres://user:hunter2@bad host/db"} //nolint:gosec // G101: fake credentials proving parse errors stay secret-free
	_, err := cfg.poolConfig()
	if err == nil {
		t.Fatal("malformed URL was accepted")
	}
	if strings.Contains(err.Error(), "hunter2") {
		t.Errorf("parse error leaked the password: %v", err)
	}
}

func TestNewRejectsMalformedURL(t *testing.T) {
	t.Parallel()
	// The parse failure happens before any pool is created, so no database
	// is needed; the error must not echo the URL (it embeds the password).
	_, err := New(t.Context(), Config{URL: "postgres://user:hunter2@bad host/db"}) //nolint:gosec // G101: fake test credentials
	if err == nil {
		t.Fatal("malformed DATABASE_URL was accepted")
	}
	if strings.Contains(err.Error(), "hunter2") {
		t.Errorf("New leaked the password: %v", err)
	}
}

func TestPoolConfigFieldsWithoutURL(t *testing.T) {
	t.Parallel()
	cfg := Config{User: "u", Password: "p", Database: "d", Host: "h:5433", MaxConns: 4}
	poolCfg, err := cfg.poolConfig()
	if err != nil {
		t.Fatalf("poolConfig: %v", err)
	}
	if poolCfg.ConnConfig.Host != "h" || poolCfg.ConnConfig.Port != 5433 {
		t.Errorf("host:port = %s:%d, want h:5433", poolCfg.ConnConfig.Host, poolCfg.ConnConfig.Port)
	}
	if poolCfg.MaxConns != 4 {
		t.Errorf("MaxConns = %d, want the configured 4", poolCfg.MaxConns)
	}
}

func TestConnStringTCP(t *testing.T) {
	cfg := Config{User: "u", Password: "p", Database: "d", Host: "db.example.com:5433"}
	got := cfg.connString()

	poolCfg, err := pgxpool.ParseConfig(got)
	if err != nil {
		t.Fatalf("connString produced an unparsable string %q: %v", got, err)
	}
	if poolCfg.ConnConfig.Host != "db.example.com" {
		t.Errorf("host = %q, want db.example.com", poolCfg.ConnConfig.Host)
	}
	if poolCfg.ConnConfig.Port != 5433 {
		t.Errorf("port = %d, want 5433", poolCfg.ConnConfig.Port)
	}
	if poolCfg.ConnConfig.User != "u" || poolCfg.ConnConfig.Password != "p" || poolCfg.ConnConfig.Database != "d" {
		t.Errorf("user/password/database = %q/%q/%q, want u/p/d",
			poolCfg.ConnConfig.User, poolCfg.ConnConfig.Password, poolCfg.ConnConfig.Database)
	}
	if !strings.Contains(got, "connect_timeout=10") {
		t.Errorf("connString %q missing connect_timeout", got)
	}
}

func TestConnStringHostWithoutPort(t *testing.T) {
	cfg := Config{User: "u", Password: "p", Database: "d", Host: "db.example.com"}
	poolCfg, err := pgxpool.ParseConfig(cfg.connString())
	if err != nil {
		t.Fatalf("parse: %v", err)
	}
	if poolCfg.ConnConfig.Host != "db.example.com" || poolCfg.ConnConfig.Port != 5432 {
		t.Errorf("host:port = %s:%d, want db.example.com:5432", poolCfg.ConnConfig.Host, poolCfg.ConnConfig.Port)
	}
}

func TestConnStringUnixSocket(t *testing.T) {
	cfg := Config{User: "u", Password: "ignored", Database: "d"}
	got := cfg.connString()
	poolCfg, err := pgxpool.ParseConfig(got)
	if err != nil {
		t.Fatalf("parse: %v", err)
	}
	if !strings.HasPrefix(poolCfg.ConnConfig.Host, "/") {
		t.Errorf("host = %q, want a Unix socket directory", poolCfg.ConnConfig.Host)
	}
	// Socket connections use peer/trust auth; the password must not leak in.
	if strings.Contains(got, "ignored") {
		t.Errorf("connString %q contains the password for a socket connection", got)
	}
}

func TestConnStringEscapesCredentials(t *testing.T) {
	cfg := Config{User: "u", Password: `pa'ss\word`, Database: "d", Host: "h:5432"}
	poolCfg, err := pgxpool.ParseConfig(cfg.connString())
	if err != nil {
		t.Fatalf("parse: %v", err)
	}
	if poolCfg.ConnConfig.Password != `pa'ss\word` {
		t.Errorf("password round-trip = %q, want pa'ss\\word", poolCfg.ConnConfig.Password)
	}
}

func TestConnectHint(t *testing.T) {
	if got := ConnectHint(nil); got != "" {
		t.Errorf("ConnectHint(nil) = %q, want empty", got)
	}

	t.Setenv("DATABASE_URL", "postgres://alice:hunter2@h:1/db")
	t.Setenv("PGSQL_USERNAME", "alice")
	t.Setenv("PGSQL_PASSWORD", "hunter2")
	t.Setenv("PGSQL_DATABASE", "")
	t.Setenv("PGSQL_HOST", "h:1")

	err := textError("FATAL: password authentication failed for user \"alice\"")
	hint := ConnectHint(err)
	for _, want := range []string{
		"PGSQL_PASSWORD does not match",
		"DATABASE_URL: set (wins over PGSQL_*; never echoed — it embeds the password)",
		`PGSQL_USERNAME: set = "alice"`,
		"PGSQL_PASSWORD: set (length 7)", // length only: the secret must not be echoed
		"PGSQL_DATABASE: not set (or empty)",
	} {
		if !strings.Contains(hint, want) {
			t.Errorf("ConnectHint missing %q in:\n%s", want, hint)
		}
	}
	if strings.Contains(hint, "hunter2") {
		t.Errorf("ConnectHint leaked the password:\n%s", hint)
	}
	t.Setenv("DATABASE_URL", "")
	if !strings.Contains(ConnectHint(err), "DATABASE_URL: not set (or empty)") {
		t.Error("ConnectHint missing the unset DATABASE_URL status")
	}

	// A non-password failure still reports the environment status.
	other := ConnectHint(textError("connection refused"))
	if strings.Contains(other, "PGSQL_PASSWORD does not match") {
		t.Errorf("password hint shown for unrelated error:\n%s", other)
	}
	if !strings.Contains(other, "Environment variable status:") {
		t.Errorf("env status missing for unrelated error:\n%s", other)
	}
}

func TestNextEventLogIndexRejectsInvalidInput(t *testing.T) {
	t.Parallel()
	var store Store
	if _, err := store.NextEventLogIndex(t.Context(), -1, 0); err == nil {
		t.Fatal("negative block was accepted")
	}
	if uint64(^uint(0)) > math.MaxInt32 {
		if _, err := store.NextEventLogIndex(
			t.Context(), 0, uint(math.MaxInt32)+1,
		); err == nil {
			t.Fatal("out-of-range minimum was accepted")
		}
	}
}

type textError string

func (e textError) Error() string { return string(e) }
