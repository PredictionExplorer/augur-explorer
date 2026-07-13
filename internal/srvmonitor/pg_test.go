package srvmonitor

import (
	"context"
	"strings"
	"testing"
)

func TestConnectPostgresInvalidPort(t *testing.T) {
	t.Parallel()
	_, err := ConnectPostgres(context.Background(), DatabaseConfig{Host: "db.example:notaport"})
	if err == nil || !strings.Contains(err.Error(), "invalid port") {
		t.Fatalf("err = %v", err)
	}
}

func TestConnectPostgresBadEnvironmentConfig(t *testing.T) {
	// pgx.ParseConfig("") honors libpq environment variables; a malformed
	// one must surface as a config error, not a connect attempt.
	t.Setenv("PGCONNECT_TIMEOUT", "notanumber")
	_, err := ConnectPostgres(context.Background(), DatabaseConfig{Host: "db.example:5432"})
	if err == nil || !strings.Contains(err.Error(), "building connection config") {
		t.Fatalf("err = %v", err)
	}
}
