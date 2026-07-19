package main

import (
	"context"
	"strings"
	"testing"

	"github.com/jackc/pgx/v5/pgxpool"
)

func TestOpenOpsDBBoundsPoolWithoutDialing(t *testing.T) {
	db, err := openOpsDB(4)(context.Background(),
		"postgres://user:pass@127.0.0.1:1/database?sslmode=disable")
	if err != nil {
		t.Fatalf("lazy pool construction: %v", err)
	}
	defer db.Close()
	pool, ok := db.(*pgxpool.Pool)
	if !ok {
		t.Fatalf("handle = %T, want *pgxpool.Pool", db)
	}
	if got := pool.Config().MaxConns; got != 4 {
		t.Fatalf("max conns = %d, want 4", got)
	}
}

func TestOpenOpsDBRejectsMalformedDSNWithoutEchoingSecrets(t *testing.T) {
	_, err := openOpsDB(2)(context.Background(), "postgres://user:supersecret@[::1")
	if err == nil || !strings.Contains(err.Error(), "db pool config") {
		t.Fatalf("error = %v, want config failure", err)
	}
	if strings.Contains(err.Error(), "supersecret") {
		t.Fatalf("error echoes the password: %v", err)
	}
}
