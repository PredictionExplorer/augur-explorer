//go:build integration

package srvmonitor

import (
	"context"
	"net"
	"strconv"
	"strings"
	"testing"

	"github.com/jackc/pgx/v5"

	"github.com/PredictionExplorer/augur-explorer/internal/testdb"
)

// dbConfigFromConnString converts testdb's connection string (a URL) into
// the monitor's DatabaseConfig shape.
func dbConfigFromConnString(t *testing.T, connString string) DatabaseConfig {
	t.Helper()
	parsed, err := pgx.ParseConfig(connString)
	if err != nil {
		t.Fatalf("parsing conn string: %v", err)
	}
	return DatabaseConfig{
		Name:   "test",
		Host:   net.JoinHostPort(parsed.Host, strconv.Itoa(int(parsed.Port))),
		DBName: parsed.Database,
		User:   parsed.User,
		Pass:   parsed.Password,
	}
}

func TestConnectPostgresIntegration(t *testing.T) {
	db := testdb.New(t)
	cfg := dbConfigFromConnString(t, db.ConnString)
	ctx := context.Background()

	conn, err := ConnectPostgres(ctx, cfg)
	if err != nil {
		t.Fatalf("ConnectPostgres: %v", err)
	}
	defer func() { _ = conn.Close(ctx) }()

	// The session runs in UTC.
	var tz string
	if err := conn.QueryRow(ctx, "SHOW timezone").Scan(&tz); err != nil {
		t.Fatal(err)
	}
	if tz != "UTC" {
		t.Fatalf("timezone = %q, want UTC", tz)
	}

	// The migrated schema serves the monitors' queries.
	var n int64
	if err := conn.QueryRow(ctx, "SELECT COUNT(*) FROM block").Scan(&n); err != nil {
		t.Fatalf("querying block table: %v", err)
	}
}

func TestConnectPostgresIntegrationDatabaseMonitor(t *testing.T) {
	db := testdb.New(t)
	ctx := context.Background()

	// Seed one block, and insert a second one during the blockWait so the
	// monitor sees progress through the real connector.
	if _, err := db.Pool.Exec(ctx,
		`INSERT INTO block(block_hash, parent_hash, block_num, ts, num_tx) VALUES ('0xa','0x9',100,NOW(),0)`); err != nil {
		t.Fatal(err)
	}

	cfg := dbConfigFromConnString(t, db.ConnString)
	m := NewDatabaseMonitor([]DatabaseConfig{cfg}, testIntervals())
	m.blockWait = func(ctx context.Context) {
		if _, err := db.Pool.Exec(ctx,
			`INSERT INTO block(block_hash, parent_hash, block_num, ts, num_tx) VALUES ('0xb','0xa',101,NOW(),0)`); err != nil {
			t.Errorf("inserting second block: %v", err)
		}
	}

	disp := newFakeDisplay()
	errCh := make(chan string, 10)
	m.check(ctx, disp, errCh)

	if msgs := drain(errCh); len(msgs) != 0 {
		t.Fatalf("unexpected errors: %v", msgs)
	}
	st := m.statuses[0]
	if !st.Alive || st.LastBlockNum != 101 {
		t.Fatalf("status = %+v, want alive at block 101", st)
	}
}

func TestConnectPostgresIntegrationErrors(t *testing.T) {
	ctx := context.Background()

	t.Run("invalid port", func(t *testing.T) {
		_, err := ConnectPostgres(ctx, DatabaseConfig{Host: "db.example:notaport"})
		if err == nil || !strings.Contains(err.Error(), "invalid port") {
			t.Fatalf("err = %v", err)
		}
	})

	t.Run("unreachable", func(t *testing.T) {
		_, err := ConnectPostgres(ctx, DatabaseConfig{Host: "127.0.0.1:1", DBName: "x", User: "u"})
		if err == nil || !strings.Contains(err.Error(), "error connecting") {
			t.Fatalf("err = %v", err)
		}
	})

	t.Run("bare host defaults to 5432", func(t *testing.T) {
		// 127.0.0.1 without a port: connection is attempted on 5432. We
		// only assert the error is a connection error, not a port parse
		// error, proving the default applied.
		cancelled, cancel := context.WithCancel(ctx)
		cancel()
		_, err := ConnectPostgres(cancelled, DatabaseConfig{Host: "127.0.0.1", DBName: "x", User: "u"})
		if err == nil || strings.Contains(err.Error(), "invalid port") {
			t.Fatalf("err = %v", err)
		}
	})
}
