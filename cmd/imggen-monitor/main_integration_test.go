//go:build integration

package main

import (
	"context"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"testing"

	"github.com/PredictionExplorer/augur-explorer/internal/testdb"
	"github.com/PredictionExplorer/augur-explorer/internal/testfixtures"
)

// TestRunScanModeIntegration drives the production wiring end to end: env
// config -> store -> repository token source -> artifact probes.
func TestRunScanModeIntegration(t *testing.T) {
	db := testdb.New(t)
	ctx := context.Background()
	if err := testfixtures.Apply(ctx, db.SQL); err != nil {
		t.Fatalf("applying fixtures: %v", err)
	}

	// Every artifact "exists": the scan reports presence for each token.
	artifacts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
		w.WriteHeader(http.StatusOK)
	}))
	t.Cleanup(artifacts.Close)

	u, err := url.Parse(db.ConnString)
	if err != nil {
		t.Fatal(err)
	}
	password, _ := u.User.Password()
	env := imEnv(artifacts.URL)
	env["PGSQL_USERNAME"] = u.User.Username()
	env["PGSQL_PASSWORD"] = password
	env["PGSQL_DATABASE"] = strings.TrimPrefix(u.Path, "/")
	env["PGSQL_HOST"] = u.Host
	for k, v := range env {
		t.Setenv(k, v)
	}

	var out, errOut strings.Builder
	if err := run(ctx, nil, envFunc(env), &out, &errOut); err != nil {
		t.Fatal(err)
	}

	report := out.String()
	if !strings.Contains(report, "Checking image/video presence") {
		t.Fatalf("report = %q", report)
	}
	if !strings.Contains(report, "image/video present") {
		t.Fatalf("report = %q, want at least one present token", report)
	}
	if strings.Contains(report, "doesn't exist") || strings.Contains(report, "error:") {
		t.Fatalf("report = %q, want clean scan", report)
	}

	// A broken token query surfaces as a scan failure rather than an empty
	// success (drop the mint table the token source reads).
	if _, err := db.Pool.Exec(ctx, "DROP TABLE cg_mint_event CASCADE"); err != nil {
		t.Fatal(err)
	}
	out.Reset()
	if err := run(ctx, nil, envFunc(env), &out, &errOut); err == nil ||
		!strings.Contains(err.Error(), "failed to list tokens") {
		t.Fatalf("err = %v", err)
	}
}

// TestRunScanModeIntegrationDBFailure proves a bad database configuration
// surfaces as a connection error rather than a hang or a partial scan.
func TestRunScanModeIntegrationDBFailure(t *testing.T) {
	env := imEnv("http://127.0.0.1:1")
	env["PGSQL_HOST"] = "127.0.0.1:1"
	env["PGSQL_USERNAME"] = "nobody"
	env["PGSQL_DATABASE"] = "nothing"
	env["PGSQL_PASSWORD"] = "wrong"
	for k, v := range env {
		t.Setenv(k, v)
	}

	var out, errOut strings.Builder
	err := run(context.Background(), nil, envFunc(env), &out, &errOut)
	if err == nil || !strings.Contains(err.Error(), "failed to connect to storage") {
		t.Fatalf("err = %v", err)
	}
}
