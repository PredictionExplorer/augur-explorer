package config

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"runtime"
	"strings"
	"testing"
)

// cliOnlyVars are environment variables documented in .env.example that are
// read outside the tag-driven service structs: operator CLI wiring and
// srvmonitor's indexed-family loader. Each entry names its reader.
var cliOnlyVars = func() map[string]string {
	vars := map[string]string{
		"PKEY_HEX":                         "cmd/cgctl, cmd/rwctl (env.go)",
		"CGAME_ADDR":                       "cmd/cgctl (env.go)",
		"GAS_PRICE_MULTIPLIER":             "cmd/cgctl, cmd/rwctl (env.go)",
		"API_BASE":                         "cmd/opsctl (cmd_smoketest.go)",
		"MAX_ETH_BID":                      "cmd/cgctl (cmd_autobid.go)",
		"MAX_CST_BID":                      "cmd/cgctl (cmd_autobid.go)",
		"RWALK_MIN_PRICE":                  "cmd/cgctl (cmd_autobid.go)",
		"TIME_BEFORE_PRIZE":                "cmd/cgctl (cmd_autobid.go)",
		"CST_BID_ANYWAY":                   "cmd/cgctl (cmd_autobid.go)",
		"AT_STARTUP_BID_UP_TO_PRICE_LEVEL": "cmd/cgctl (cmd_autobid.go)",
		"SRV1_WEB_API_NAME":                "internal/srvmonitor (config.go)",
		"SRV1_WEB_API_HOST":                "internal/srvmonitor (config.go)",
		"SRV1_WEB_API_PORT":                "internal/srvmonitor (config.go)",
		"SRV1_WEB_API_URI":                 "internal/srvmonitor (config.go)",
		"SRV1_WEB_API_PUBLIC_URL":          "internal/srvmonitor (config.go)",
		"ANOMALY_TITLE":                    "internal/srvmonitor (config.go)",
		"ANOMALY_SSH_USER":                 "internal/srvmonitor (config.go)",
		"ANOMALY_SSH_HOST":                 "internal/srvmonitor (config.go)",
		"ANOMALY_REMOTE_FILE":              "internal/srvmonitor (config.go)",
		"ANOMALY_STALE_SECS":               "internal/srvmonitor (config.go)",
	}
	for i := 1; i <= 6; i++ {
		for _, suffix := range []string{"NAME", "HOST", "DBNAME", "USER", "PASS", "TABLE", "COLUMN"} {
			vars[fmt.Sprintf("DB_L1EVT%d_%s", i, suffix)] = "internal/srvmonitor (config.go)"
		}
	}
	return vars
}()

// envExampleEntry is one (possibly commented) VAR=value line.
type envExampleEntry struct {
	value     string
	commented bool
}

// readEnvExample parses the repository's .env.example.
func readEnvExample(t *testing.T) map[string]envExampleEntry {
	t.Helper()
	_, thisFile, _, ok := runtime.Caller(0)
	if !ok {
		t.Fatal("cannot locate this source file")
	}
	path := filepath.Join(filepath.Dir(thisFile), "..", "..", ".env.example")
	f, err := os.Open(path) //nolint:gosec // repo-relative documentation file
	if err != nil {
		t.Fatalf("open .env.example: %v", err)
	}
	defer f.Close() //nolint:errcheck // read-only handle

	lineRe := regexp.MustCompile(`^(#?)([A-Z][A-Z0-9_]*)=(.*)$`)
	entries := map[string]envExampleEntry{}
	sc := bufio.NewScanner(f)
	for sc.Scan() {
		m := lineRe.FindStringSubmatch(strings.TrimSpace(sc.Text()))
		if m == nil {
			continue
		}
		entries[m[2]] = envExampleEntry{value: m[3], commented: m[1] == "#"}
	}
	if err := sc.Err(); err != nil {
		t.Fatalf("read .env.example: %v", err)
	}
	return entries
}

// TestEnvExampleMatchesServiceConfigs is the documentation gate for §8.3:
// .env.example and the typed service configurations can never drift apart.
func TestEnvExampleMatchesServiceConfigs(t *testing.T) {
	t.Parallel()
	entries := readEnvExample(t)

	declared := map[string]Var{}
	apiserverRequired := map[string]bool{}
	for svc, cfg := range ServiceConfigs() {
		for _, v := range Vars(cfg) {
			declared[v.Name] = v
			if svc == "apiserver" && v.Required {
				apiserverRequired[v.Name] = true
			}
		}
	}

	// Every variable a service reads is documented.
	for name, v := range declared {
		entry, ok := entries[name]
		if !ok {
			t.Errorf(".env.example is missing %s (declared by a service config)", name)
			continue
		}
		// Documented example values for defaulted variables must state the
		// real default, or they would mislead operators.
		if v.Default != "" && entry.commented && entry.value != "" && entry.value != v.Default {
			t.Errorf(".env.example documents %s=%s but the code default is %q", name, entry.value, v.Default)
		}
		// Variables the API server requires should be presented uncommented
		// (ready to fill in). Optional-service requirements (imggen,
		// rwalk-alarm) stay commented: most deployments don't run them.
		if apiserverRequired[name] && entry.commented {
			t.Errorf(".env.example comments out %s, which the API server requires", name)
		}
	}

	// Every documented variable maps to a known reader.
	for name := range entries {
		if _, ok := declared[name]; ok {
			continue
		}
		if _, ok := cliOnlyVars[name]; ok {
			continue
		}
		t.Errorf(".env.example documents %s, which no service config declares (add it to a struct or the cliOnlyVars registry)", name)
	}
}

// TestCLIOnlyVarsAreReferencedInCode keeps the cliOnlyVars registry honest:
// a variable removed from the CLIs must leave the registry (and the example
// file) too. It scans the referenced cmd/ packages for the literal name.
func TestCLIOnlyVarsAreReferencedInCode(t *testing.T) {
	t.Parallel()
	_, thisFile, _, ok := runtime.Caller(0)
	if !ok {
		t.Fatal("cannot locate this source file")
	}
	root := filepath.Join(filepath.Dir(thisFile), "..", "..")

	var sources strings.Builder
	for _, dir := range []string{"cmd/cgctl", "cmd/rwctl", "cmd/opsctl", "internal/srvmonitor"} {
		files, err := filepath.Glob(filepath.Join(root, dir, "*.go"))
		if err != nil {
			t.Fatalf("glob %s: %v", dir, err)
		}
		for _, f := range files {
			b, err := os.ReadFile(f) //nolint:gosec // repo source files
			if err != nil {
				t.Fatalf("read %s: %v", f, err)
			}
			sources.Write(b)
		}
	}
	indexedDigits := regexp.MustCompile(`[0-9]+`)
	for name, reader := range cliOnlyVars {
		referenced := strings.Contains(sources.String(), `"`+name+`"`)
		for _, match := range indexedDigits.FindAllStringIndex(name, -1) {
			template := name[:match[0]] + "%d" + name[match[1]:]
			referenced = referenced || strings.Contains(sources.String(), `"`+template+`"`)
		}
		if !referenced {
			t.Errorf("cliOnlyVars entry %s claims reader %q but no configured source references it", name, reader)
		}
	}
}
