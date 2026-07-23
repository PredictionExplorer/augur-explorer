package store

import (
	"testing"
	"time"
)

// TestPoolConfigServerSideTimeouts pins the D22 defense-in-depth wiring:
// positive StatementTimeout/IdleInTxSessionTimeout become the corresponding
// PostgreSQL session GUCs in milliseconds, and zero values leave the server
// defaults untouched (operator CLIs with legitimately heavy statements).
func TestPoolConfigServerSideTimeouts(t *testing.T) {
	t.Parallel()
	cfg := Config{
		User: "u", Password: "p", Database: "d", Host: "h:5432",
		StatementTimeout:       30 * time.Second,
		IdleInTxSessionTimeout: 5 * time.Minute,
	}
	poolCfg, err := cfg.poolConfig()
	if err != nil {
		t.Fatalf("poolConfig: %v", err)
	}
	params := poolCfg.ConnConfig.RuntimeParams
	if got := params["statement_timeout"]; got != "30000" {
		t.Errorf("statement_timeout = %q, want 30000 (milliseconds)", got)
	}
	if got := params["idle_in_transaction_session_timeout"]; got != "300000" {
		t.Errorf("idle_in_transaction_session_timeout = %q, want 300000 (milliseconds)", got)
	}
}

func TestPoolConfigZeroTimeoutsLeaveServerDefaults(t *testing.T) {
	t.Parallel()
	cfg := Config{User: "u", Password: "p", Database: "d", Host: "h:5432"}
	poolCfg, err := cfg.poolConfig()
	if err != nil {
		t.Fatalf("poolConfig: %v", err)
	}
	params := poolCfg.ConnConfig.RuntimeParams
	if _, ok := params["statement_timeout"]; ok {
		t.Error("statement_timeout set for a zero config; operator CLIs must keep the server default")
	}
	if _, ok := params["idle_in_transaction_session_timeout"]; ok {
		t.Error("idle_in_transaction_session_timeout set for a zero config")
	}
}
