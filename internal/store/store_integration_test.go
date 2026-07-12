//go:build integration

package store_test

import (
	"context"
	"net"
	"strconv"
	"testing"
	"time"

	"github.com/jackc/pgx/v5"

	"github.com/PredictionExplorer/augur-explorer/internal/store"
	"github.com/PredictionExplorer/augur-explorer/internal/testdb"
)

// configFromConnString converts the testcontainer's URL connection string
// into the PGSQL_*-shaped store.Config that production builds from the
// environment.
func configFromConnString(t *testing.T, connString string) store.Config {
	t.Helper()
	parsed, err := pgx.ParseConfig(connString)
	if err != nil {
		t.Fatalf("parsing container conn string: %v", err)
	}
	return store.Config{
		User:     parsed.User,
		Password: parsed.Password,
		Database: parsed.Database,
		Host:     net.JoinHostPort(parsed.Host, strconv.Itoa(int(parsed.Port))),
	}
}

// TestNewConnectsQueriesAndCloses covers the production constructor end to
// end against a real PostgreSQL: keepalive dialer, UTC/search_path runtime
// params, first-attempt ping and a query round trip through the pool.
func TestNewConnectsQueriesAndCloses(t *testing.T) {
	db := testdb.New(t)
	ctx := context.Background()

	st, err := store.New(ctx, configFromConnString(t, db.ConnString))
	if err != nil {
		t.Fatalf("store.New: %v", err)
	}
	defer st.Close()

	// The pool must be usable and pinned to UTC.
	var tz string
	if err := st.Pool().QueryRow(ctx, "SHOW timezone").Scan(&tz); err != nil {
		t.Fatalf("querying timezone: %v", err)
	}
	if tz != "UTC" {
		t.Errorf("timezone = %q, want UTC", tz)
	}

	// Address round trip proves the Store wiring (pool + LRU cache).
	const addr = "0x00000000000000000000000000000000000000AB"
	aid, err := st.LookupOrCreateAddress(ctx, addr, 1, 1)
	if err != nil {
		t.Fatalf("LookupOrCreateAddress: %v", err)
	}
	again, err := st.LookupAddressID(ctx, addr)
	if err != nil || again != aid {
		t.Errorf("LookupAddressID = (%d, %v), want (%d, nil)", again, err, aid)
	}
}

// TestNewUnreachableHostHonorsContext pins the startup retry loop's
// cancellation exit: a dead endpoint fails the first ping and the context
// deadline interrupts the backoff wait instead of burning all ten attempts.
func TestNewUnreachableHostHonorsContext(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), 1500*time.Millisecond)
	defer cancel()

	start := time.Now()
	_, err := store.New(ctx, store.Config{
		User:     "u",
		Password: "p",
		Database: "d",
		Host:     "127.0.0.1:1", // nothing listens on port 1
	})
	if err == nil {
		t.Fatal("store.New against a dead endpoint should fail")
	}
	if elapsed := time.Since(start); elapsed > 10*time.Second {
		t.Errorf("store.New took %v; the context deadline should bound the retry loop", elapsed)
	}
}

// TestLastBlockNumWatermarkStates covers the three shapes of the singleton
// watermark row: the migration-seeded NULL, an explicit value, and the
// defensive no-row fallback.
func TestLastBlockNumWatermarkStates(t *testing.T) {
	db := testdb.New(t)
	ctx := context.Background()
	st := store.NewFromPool(db.Pool)

	// Migration 00005 seeds the singleton row with a NULL block_num.
	got, err := st.LastBlockNum(ctx)
	if err != nil || got != 0 {
		t.Errorf("fresh database LastBlockNum = (%d, %v), want (0, nil)", got, err)
	}

	if err := st.SetLastBlockNum(ctx, 12345); err != nil {
		t.Fatalf("SetLastBlockNum: %v", err)
	}
	got, err = st.LastBlockNum(ctx)
	if err != nil || got != 12345 {
		t.Errorf("LastBlockNum after set = (%d, %v), want (12345, nil)", got, err)
	}

	// A rowless table (pre-00005 databases) still reads as zero.
	if _, err := db.Pool.Exec(ctx, "DELETE FROM last_block"); err != nil {
		t.Fatalf("clearing last_block: %v", err)
	}
	got, err = st.LastBlockNum(ctx)
	if err != nil || got != 0 {
		t.Errorf("rowless LastBlockNum = (%d, %v), want (0, nil)", got, err)
	}
}
