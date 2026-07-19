package srvmonitor

import (
	"maps"
	"strings"
	"testing"
	"time"
)

// envMap adapts a map to the getenv seam.
func envMap(m map[string]string) func(string) string {
	return func(key string) string { return m[key] }
}

// minimalEnv satisfies Validate: one RPC node and one Layer1 database.
func minimalEnv() map[string]string {
	return map[string]string{
		"RPC0_NAME":       "node0",
		"RPC0_URL":        "http://rpc0.example",
		"RPC0_CHAINID":    "1",
		"DB_L1_NAME_SRV1": "l1",
		"DB_L1_HOST_SRV1": "db1.example:5432",
	}
}

func TestLoadFromEnvMinimal(t *testing.T) {
	t.Parallel()
	cfg, err := LoadFromEnv(envMap(minimalEnv()))
	if err != nil {
		t.Fatal(err)
	}
	if len(cfg.RPCNodes) != 1 || cfg.RPCNodes[0].Name != "node0" || cfg.RPCNodes[0].IsOfficial {
		t.Fatalf("RPCNodes = %+v", cfg.RPCNodes)
	}
	if len(cfg.Layer1DBs) != 1 || cfg.Layer1DBs[0].Name != "l1" {
		t.Fatalf("Layer1DBs = %+v", cfg.Layer1DBs)
	}
	if cfg.Intervals != DefaultIntervals() {
		t.Fatalf("Intervals = %+v, want defaults", cfg.Intervals)
	}
	if cfg.MobileNotif || cfg.Anomaly.Enabled() {
		t.Fatalf("cfg = %+v, want notifications and anomaly disabled", cfg)
	}
}

func TestLoadFromEnvFull(t *testing.T) {
	t.Parallel()
	env := minimalEnv()
	maps.Copy(env, map[string]string{
		"OFFICIAL_RPC_MAINNET": "node0",
		"RPC1_NAME":            "follower",
		"RPC1_URL":             "http://rpc1.example",
		"RPC1_CHAINID":         "1",

		"DB_L1EVT1_NAME":   "evt1",
		"DB_L1EVT1_HOST":   "evt1.example",
		"DB_L1EVT2_NAME":   "evt2",
		"DB_L1EVT2_HOST":   "evt2.example",
		"DB_L1EVT2_TABLE":  "rw_proc_status",
		"DB_L1EVT2_COLUMN": "last_block",

		"APP_STATUS_SRV1_TITLE": "cg app",
		"APP_STATUS_SRV1_HOST":  "app1.example",

		"SRV1_WEB_API_NAME":       "api1",
		"SRV1_WEB_API_HOST":       "10.0.0.5",
		"SRV1_WEB_API_PORT":       "8090",
		"SRV1_WEB_API_URI":        "/api/time/current",
		"SRV1_WEB_API_PUBLIC_URL": "https://a1.example/api",

		"SSH_CMD_DF_SRV1_NAME":    "disk1",
		"SSH_CMD_DF_SRV1_USER":    "ops",
		"SSH_CMD_DF_SRV1_IP":      "10.0.0.6",
		"SSH_CMD_DF_SRV1_DEVICES": "/dev/sda1",

		"SSL_CERT1_HOST":       "a1.example",
		"SSL_CERT2_HOST":       "b2.example",
		"SSL_CERT2_PORT":       "8443",
		"SSL_CERT2_NAME":       "b2",
		"SSL_CERT2_SERVERNAME": "b2.sni",

		"ANOMALY_SSH_USER":    "ops",
		"ANOMALY_SSH_HOST":    "websrv",
		"ANOMALY_REMOTE_FILE": "/var/log/anomalies.log",

		"DB_RWLK_NAME_SRV":    "rw",
		"DB_RWLK_HOST_SRV":    "rw.example",
		"RWALK_CONTRACT_ADDR": "0x1111111111111111111111111111111111111111",

		"MOBILE_NOTIF": "yes",
	})

	cfg, err := LoadFromEnv(envMap(env))
	if err != nil {
		t.Fatal(err)
	}

	if !cfg.RPCNodes[0].IsOfficial || cfg.RPCNodes[1].IsOfficial {
		t.Fatalf("official flags = %+v", cfg.RPCNodes)
	}

	if len(cfg.EventTableDBs) != 2 {
		t.Fatalf("EventTableDBs = %+v", cfg.EventTableDBs)
	}
	// Defaults apply to evt1; explicit values to evt2.
	if cfg.EventTableDBs[0].TableName != "cg_proc_status" || cfg.EventTableDBs[0].ColumnName != "last_evt_id" {
		t.Fatalf("evt1 = %+v", cfg.EventTableDBs[0])
	}
	if cfg.EventTableDBs[1].TableName != "rw_proc_status" || cfg.EventTableDBs[1].ColumnName != "last_block" {
		t.Fatalf("evt2 = %+v", cfg.EventTableDBs[1])
	}

	if len(cfg.ApplicationDBs) != 1 || cfg.ApplicationDBs[0].Name != "cg app" {
		t.Fatalf("ApplicationDBs = %+v", cfg.ApplicationDBs)
	}
	if len(cfg.WebAPIs) != 1 || cfg.WebAPIs[0].PublicURL != "https://a1.example/api" {
		t.Fatalf("WebAPIs = %+v", cfg.WebAPIs)
	}
	if len(cfg.DiskMonitors) != 1 || cfg.DiskMonitors[0].DeviceList != "/dev/sda1" {
		t.Fatalf("DiskMonitors = %+v", cfg.DiskMonitors)
	}

	if len(cfg.SSLCerts) != 2 {
		t.Fatalf("SSLCerts = %+v", cfg.SSLCerts)
	}
	if cfg.SSLCerts[1].Port != "8443" || cfg.SSLCerts[1].ServerName != "b2.sni" {
		t.Fatalf("SSLCerts[1] = %+v", cfg.SSLCerts[1])
	}

	if !cfg.Anomaly.Enabled() {
		t.Fatalf("Anomaly = %+v, want enabled", cfg.Anomaly)
	}
	if cfg.RWalkDB.Host != "rw.example" || cfg.RWalkImage.ContractAddr == "" {
		t.Fatalf("RWalk config = %+v / %+v", cfg.RWalkDB, cfg.RWalkImage)
	}
	if !cfg.MobileNotif {
		t.Fatal("MobileNotif = false, want true")
	}
}

func TestLoadFromEnvMobileNotifSpellings(t *testing.T) {
	t.Parallel()
	for value, want := range map[string]bool{
		"yes": true, "TRUE": true, "1": true,
		"no": false, "0": false, "": false, "on": false,
	} {
		env := minimalEnv()
		env["MOBILE_NOTIF"] = value
		cfg, err := LoadFromEnv(envMap(env))
		if err != nil {
			t.Fatal(err)
		}
		if cfg.MobileNotif != want {
			t.Fatalf("MOBILE_NOTIF=%q => %v, want %v", value, cfg.MobileNotif, want)
		}
	}
}

func TestLoadFromEnvValidation(t *testing.T) {
	t.Parallel()

	t.Run("no rpc", func(t *testing.T) {
		t.Parallel()
		env := minimalEnv()
		delete(env, "RPC0_NAME")
		_, err := LoadFromEnv(envMap(env))
		if err == nil || !strings.Contains(err.Error(), "at least one RPC node") {
			t.Fatalf("err = %v", err)
		}
	})

	t.Run("no layer1 db", func(t *testing.T) {
		t.Parallel()
		env := minimalEnv()
		delete(env, "DB_L1_NAME_SRV1")
		_, err := LoadFromEnv(envMap(env))
		if err == nil || !strings.Contains(err.Error(), "at least one Layer1 database") {
			t.Fatalf("err = %v", err)
		}
	})

	t.Run("both missing lists both", func(t *testing.T) {
		t.Parallel()
		_, err := LoadFromEnv(envMap(map[string]string{}))
		if err == nil || !strings.Contains(err.Error(), "RPC node") || !strings.Contains(err.Error(), "Layer1") {
			t.Fatalf("err = %v", err)
		}
	})
}

func TestDefaultIntervals(t *testing.T) {
	t.Parallel()
	iv := DefaultIntervals()
	// Pin the production periods: they mirror the legacy hardcoded
	// constants and operators depend on their cadence.
	want := Intervals{
		RPC:            60 * time.Second,
		RPCBlockWait:   60 * time.Second,
		DB:             60 * time.Second,
		DBBlockWait:    60 * time.Second,
		EventTable:     60 * time.Second,
		EventTableWait: 120 * time.Second,
		Disk:           600 * time.Second,
		Image:          900 * time.Second,
		SSL:            3600 * time.Second,
		Anomaly:        300 * time.Second,
	}
	if iv != want {
		t.Fatalf("DefaultIntervals() = %+v, want %+v", iv, want)
	}
}
