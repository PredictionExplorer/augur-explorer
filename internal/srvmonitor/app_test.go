package srvmonitor

import (
	"reflect"
	"strings"
	"testing"
)

// fullConfig returns a config enabling every monitor.
func fullConfig() *Config {
	return &Config{
		RPCNodes:       []RPCConfig{{Name: "n", URL: "http://rpc"}},
		Layer1DBs:      []DatabaseConfig{dbConfig("l1a"), dbConfig("l1b")},
		ApplicationDBs: []DatabaseConfig{dbConfig("app1"), dbConfig("app2"), dbConfig("app3")},
		EventTableDBs: []EventTableConfig{
			evtConfig("evt1"),
			evtConfig("evt2"),
			evtConfig("evt3"),
			evtConfig("evt4"),
			evtConfig("evt5"),
			evtConfig("evt6"),
		},
		WebAPIs:      []WebAPIConfig{{Title: "api1"}, {Title: "api2"}},
		DiskMonitors: []DiskConfig{{Title: "disk1"}},
		SSLCerts:     []SSLCertConfig{{Host: "a"}, {Host: "b"}, {Host: "c"}},
		Anomaly:      anomalyConfig(),
		RWalkDB:      dbConfig("rwalk"),
		RWalkImage:   ImageServerConfig{URL: "http://img"},
		MobileNotif:  true,
		Intervals:    testIntervals(),
	}
}

func TestComputeLayoutFullConfig(t *testing.T) {
	t.Parallel()
	cfg := fullConfig()
	l := ComputeLayout(cfg)

	// 3 application DBs: event table entries start at 25+3.
	if l.EventTableBaseY != 28 {
		t.Fatalf("EventTableBaseY = %d, want 28", l.EventTableBaseY)
	}
	// Web API header: 25 + 3 apps + 6 event DBs + 1 gap = 35.
	if l.WebAPIBaseY != 35 {
		t.Fatalf("WebAPIBaseY = %d, want 35", l.WebAPIBaseY)
	}
	// Image: 35 + 1 header + 2 APIs + 1 gap = 39.
	if l.ImageBaseY != 39 {
		t.Fatalf("ImageBaseY = %d, want 39", l.ImageBaseY)
	}
	// Anomaly: image header + 2 content lines below.
	if l.AnomalyBaseY != 42 {
		t.Fatalf("AnomalyBaseY = %d, want 42", l.AnomalyBaseY)
	}
	// SSL column is fixed.
	if l.SSLBaseX != 62 || l.SSLBaseY != 24 {
		t.Fatalf("SSL position = (%d,%d)", l.SSLBaseX, l.SSLBaseY)
	}
	// With SSL certs, errors go under the SSL column: 24 + 1 header + 3 certs + 1.
	if l.ErrorX != 62 || l.ErrorY != 29 {
		t.Fatalf("Error position = (%d,%d), want (62,29)", l.ErrorX, l.ErrorY)
	}
}

func TestComputeLayoutWithoutSSLAndAnomaly(t *testing.T) {
	t.Parallel()
	cfg := fullConfig()
	cfg.SSLCerts = nil
	cfg.Anomaly = AnomalyConfig{}
	l := ComputeLayout(cfg)

	// Errors fall to the bottom of the left column: ImageBaseY+3+1.
	if l.ErrorX != 1 || l.ErrorY != l.ImageBaseY+4 {
		t.Fatalf("Error position = (%d,%d), want (1,%d)", l.ErrorX, l.ErrorY, l.ImageBaseY+4)
	}
}

func TestComputeLayoutAnomalyExtendsLeftColumn(t *testing.T) {
	t.Parallel()
	cfg := fullConfig()
	cfg.SSLCerts = nil
	l := ComputeLayout(cfg)

	// With anomaly enabled, the error area sits below the anomaly section.
	if l.ErrorY != l.AnomalyBaseY+AnomalyDisplayCount+1 {
		t.Fatalf("ErrorY = %d, want %d", l.ErrorY, l.AnomalyBaseY+AnomalyDisplayCount+1)
	}
}

func TestBuildManagerRegistersEverything(t *testing.T) {
	t.Parallel()
	logger, sink := newTestLogger()
	mgr, layout := BuildManager(fullConfig(), newFakeDisplay(), logger, t.TempDir())

	want := []string{
		"RPC Monitor",
		"Database Monitor",
		"Application Layer Monitor",
		"Event Table Monitor",
		"Web API Monitor",
		"Disk Usage Monitor",
		"Image Monitor",
		"WebSrv Anomaly Monitor",
		"SSL Monitor",
	}
	if got := mgr.MonitorNames(); !reflect.DeepEqual(got, want) {
		t.Fatalf("monitors = %v, want %v", got, want)
	}
	if layout != ComputeLayout(fullConfig()) {
		t.Fatalf("layout = %+v", layout)
	}
	if !mgr.alarmTracker.enabled {
		t.Fatal("alarm tracker must be enabled with MobileNotif")
	}
	log := sink.String()
	for _, name := range want {
		if !strings.Contains(log, "Registered: "+name) {
			t.Fatalf("log missing registration of %q:\n%s", name, log)
		}
	}
	if !strings.Contains(log, "Mobile notifications enabled") {
		t.Fatalf("log = %q", log)
	}
}

func TestBuildManagerGatesOptionalMonitors(t *testing.T) {
	t.Parallel()
	logger, _ := newTestLogger()
	cfg := &Config{
		RPCNodes:  []RPCConfig{{Name: "n", URL: "u"}},
		Layer1DBs: []DatabaseConfig{dbConfig("l1")},
		Intervals: testIntervals(),
	}
	mgr, _ := BuildManager(cfg, newFakeDisplay(), logger, t.TempDir())

	want := []string{"RPC Monitor", "Database Monitor"}
	if got := mgr.MonitorNames(); !reflect.DeepEqual(got, want) {
		t.Fatalf("monitors = %v, want %v", got, want)
	}
	if mgr.alarmTracker.enabled {
		t.Fatal("alarm tracker must be disabled without MobileNotif")
	}
}

func TestBuildManagerImageMonitorNeedsBothSettings(t *testing.T) {
	t.Parallel()
	logger, _ := newTestLogger()

	// URL without a database host: no image monitor.
	cfg := &Config{
		RPCNodes:   []RPCConfig{{Name: "n", URL: "u"}},
		Layer1DBs:  []DatabaseConfig{dbConfig("l1")},
		RWalkImage: ImageServerConfig{URL: "http://img"},
		Intervals:  testIntervals(),
	}
	mgr, _ := BuildManager(cfg, newFakeDisplay(), logger, t.TempDir())
	for _, name := range mgr.MonitorNames() {
		if name == "Image Monitor" {
			t.Fatal("image monitor must not register without a database host")
		}
	}
}
