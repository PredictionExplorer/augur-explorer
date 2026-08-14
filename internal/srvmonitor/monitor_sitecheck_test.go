package srvmonitor

import (
	"context"
	"errors"
	"log/slog"
	"strings"
	"sync"
	"testing"
)

// scriptedSiteCheckRunner serves one canned subprocess result and records
// the command it received.
type scriptedSiteCheckRunner struct {
	mu   sync.Mutex
	out  []byte
	err  error
	name string
	args []string
}

func (r *scriptedSiteCheckRunner) run(_ context.Context, name string, args ...string) ([]byte, error) {
	r.mu.Lock()
	defer r.mu.Unlock()
	r.name = name
	r.args = args
	return r.out, r.err
}

func newSiteCheckMonitorForTest(cfg SiteCheckConfig, runner *scriptedSiteCheckRunner) *SiteCheckMonitor {
	m := NewSiteCheckMonitor(cfg, 50, slog.New(slog.DiscardHandler), testIntervals())
	m.run = runner.run
	return m
}

func TestSiteCheckMonitorHealthy(t *testing.T) {
	t.Parallel()
	runner := &scriptedSiteCheckRunner{
		out: []byte(`{"sites":[{"name":"cosmicsignature","authentic":[]}],"endpoints":[],"ssl":[],"authenticCount":0}`),
	}
	m := newSiteCheckMonitorForTest(SiteCheckConfig{
		Script: "/opt/tools/site-checker/check-sites.js",
		Node:   "/usr/local/bin/node",
	}, runner)
	disp := newFakeDisplay()
	errCh := make(chan string, 10)

	m.check(context.Background(), disp, errCh)

	if got := disp.Row(50); !strings.Contains(got, "Site Checker") {
		t.Fatalf("header row = %q", got)
	}
	if got := disp.Row(51); !strings.Contains(got, "All sites OK") {
		t.Fatalf("row = %q", got)
	}
	if disp.FgAt(1, 51) != ColorGreen {
		t.Fatalf("fg = %v, want green", disp.FgAt(1, 51))
	}
	if msgs := drain(errCh); len(msgs) != 0 {
		t.Fatalf("unexpected errors: %v", msgs)
	}

	// The configured node binary runs the script with the JSON-to-stdout flag.
	runner.mu.Lock()
	defer runner.mu.Unlock()
	if runner.name != "/usr/local/bin/node" {
		t.Fatalf("binary = %q", runner.name)
	}
	joined := strings.Join(runner.args, " ")
	if !strings.Contains(joined, "/opt/tools/site-checker/check-sites.js --json -") {
		t.Fatalf("args = %q", joined)
	}
}

func TestSiteCheckMonitorShowsFirstTwoErrorsOnly(t *testing.T) {
	t.Parallel()
	runner := &scriptedSiteCheckRunner{
		out: []byte(`{
			"sites":[{"name":"axiomzero","authentic":[
				{"severity":"critical","kind":"http-status","site":"axiomzero","message":"Main document returned HTTP 500"},
				{"severity":"error","kind":"js-error","site":"axiomzero","message":"Uncaught JS exception: boom"}
			]}],
			"endpoints":[{"name":"cosmic API a1","authentic":[
				{"severity":"critical","kind":"endpoint","site":"cosmic API a1","message":"Endpoint returned HTTP 502"}
			]}],
			"ssl":[{"host":"a1.cosmicsignature.com:443","findings":[
				{"severity":"warning","kind":"ssl","site":"a1.cosmicsignature.com:443","message":"expires in 3 day(s)"}
			]}],
			"authenticCount":3}`),
		err: errors.New("exit status 1"), // alarm exit code, JSON still authoritative
	}
	m := newSiteCheckMonitorForTest(SiteCheckConfig{Script: "check-sites.js"}, runner)
	disp := newFakeDisplay()
	errCh := make(chan string, 10)

	m.check(context.Background(), disp, errCh)

	row1 := disp.Row(51)
	if !strings.Contains(row1, "axiomzero: Main document returned HTTP 500") {
		t.Fatalf("row1 = %q", row1)
	}
	if disp.FgAt(1, 51) != ColorRed {
		t.Fatalf("fg = %v, want red", disp.FgAt(1, 51))
	}
	// Only two lines are shown; the second carries the overflow count.
	// SSL warnings are excluded, so 3 problems total -> "+1 more".
	row2 := disp.Row(52)
	if !strings.Contains(row2, "Uncaught JS exception: boom") || !strings.Contains(row2, "(+1 more)") {
		t.Fatalf("row2 = %q", row2)
	}
	if got := disp.Row(53); got != "" {
		t.Fatalf("row3 should be empty, got %q", got)
	}

	// Only the first two problems reach the shared error area.
	msgs := drain(errCh)
	if len(msgs) != 2 {
		t.Fatalf("errors = %v, want 2", msgs)
	}
	if !strings.Contains(msgs[0], "Main document returned HTTP 500") {
		t.Fatalf("msgs[0] = %q", msgs[0])
	}
}

func TestSiteCheckMonitorCheckerFailure(t *testing.T) {
	t.Parallel()
	runner := &scriptedSiteCheckRunner{
		out: nil,
		err: errors.New("exec: \"node\": executable file not found in $PATH"),
	}
	m := newSiteCheckMonitorForTest(SiteCheckConfig{Script: "check-sites.js"}, runner)
	disp := newFakeDisplay()
	errCh := make(chan string, 10)

	m.check(context.Background(), disp, errCh)

	if got := disp.Row(51); !strings.Contains(got, "checker failed") {
		t.Fatalf("row = %q", got)
	}
	if disp.FgAt(1, 51) != ColorRed {
		t.Fatalf("fg = %v, want red", disp.FgAt(1, 51))
	}
	if msgs := drain(errCh); len(msgs) != 1 || !strings.Contains(msgs[0], "checker failed") {
		t.Fatalf("errors = %v", msgs)
	}
}

func TestSiteCheckMonitorBeforeFirstCheck(t *testing.T) {
	t.Parallel()
	m := newSiteCheckMonitorForTest(SiteCheckConfig{Script: "check-sites.js"}, &scriptedSiteCheckRunner{})
	disp := newFakeDisplay()

	m.display(disp)

	if got := disp.Row(51); !strings.Contains(got, "Checking...") {
		t.Fatalf("row = %q", got)
	}
	if disp.FgAt(1, 51) != ColorYellow {
		t.Fatalf("fg = %v, want yellow", disp.FgAt(1, 51))
	}
}

func TestLoadFromEnvSiteCheck(t *testing.T) {
	t.Parallel()
	env := map[string]string{
		// Minimum valid config
		"RPC0_NAME":       "mainnet",
		"RPC0_URL":        "http://localhost:8545",
		"DB_L1_NAME_SRV1": "db1",
		"DB_L1_HOST_SRV1": "localhost",
		// An iDRAC so the layout assertion below exercises the "site check
		// sits below the iDRAC block" placement.
		"IDRAC1_HOST": "10.0.0.1",
		"IDRAC1_USER": "mon",
		"IDRAC1_PASS": "pw",
		// Site checker
		"SITE_CHECK_SCRIPT": "/opt/tools/site-checker/check-sites.js",
		"SITE_CHECK_NODE":   "/usr/local/bin/node",
		"SITE_CHECK_CONFIG": "/opt/tools/site-checker/config.json",
		"SITE_CHECK_TITLE":  "Websites",
	}
	cfg, err := LoadFromEnv(func(key string) string { return env[key] })
	if err != nil {
		t.Fatalf("LoadFromEnv: %v", err)
	}
	if !cfg.SiteCheck.Enabled() {
		t.Fatal("site check should be enabled")
	}
	want := SiteCheckConfig{
		Title:  "Websites",
		Script: "/opt/tools/site-checker/check-sites.js",
		Node:   "/usr/local/bin/node",
		Config: "/opt/tools/site-checker/config.json",
	}
	if cfg.SiteCheck != want {
		t.Fatalf("SiteCheck = %+v, want %+v", cfg.SiteCheck, want)
	}
	if cfg.Intervals.SiteCheck.Seconds() != 300 {
		t.Fatalf("interval = %v, want 5m", cfg.Intervals.SiteCheck)
	}

	// Layout: the section lands below the iDRAC block and extends the left
	// column.
	withoutSiteCheck := *cfg
	withoutSiteCheck.SiteCheck = SiteCheckConfig{}
	layoutWith := ComputeLayout(cfg)
	layoutWithout := ComputeLayout(&withoutSiteCheck)
	if layoutWith.SiteCheckBaseY <= layoutWith.IDRACBaseY {
		t.Fatalf("SiteCheckBaseY = %d, IDRACBaseY = %d", layoutWith.SiteCheckBaseY, layoutWith.IDRACBaseY)
	}
	if layoutWith.ErrorY <= layoutWithout.ErrorY {
		t.Fatalf("error area should move down: with=%d without=%d", layoutWith.ErrorY, layoutWithout.ErrorY)
	}
}
