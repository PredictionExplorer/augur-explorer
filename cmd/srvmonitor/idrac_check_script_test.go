package main

import (
	"context"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"testing"
	"time"
)

// selFixture builds a Redfish SEL Entries payload mirroring a real iDRAC9
// response: a fatal bus error, an informational OEM event and a recovered
// temperature excursion.
func selFixture(fatalCreated string) string {
	return fmt.Sprintf(`{
  "Members": [
    {"Id": "1", "Created": %q, "Severity": "Critical",
     "Message": "A bus fatal error was detected on a component at slot 2."},
    {"Id": "2", "Created": %q, "Severity": "Critical",
     "Message": "A fatal error was detected on a component at bus 93 device 0 function 0."},
    {"Id": "3", "Created": %q, "Severity": "OK",
     "Message": "An OEM diagnostic event occurred."},
    {"Id": "4", "Created": "2026-06-18T01:27:24-05:00", "Severity": "Critical",
     "Message": "CPU 2 temperature is greater than the upper critical threshold."},
    {"Id": "5", "Created": "2026-06-18T01:27:39-05:00", "Severity": "OK",
     "Message": "CPU 2 temperature is within range."}
  ]
}`, fatalCreated, fatalCreated, fatalCreated)
}

// runIDRACScript executes the real idrac_check.sh with a stub curl serving
// the fixture, returning trimmed output and the exit code.
func runIDRACScript(t *testing.T, fixture string, curlExit int) (string, int) {
	t.Helper()
	for _, tool := range []string{"bash", "jq"} {
		if _, err := exec.LookPath(tool); err != nil {
			t.Skipf("%s not available: %v", tool, err)
		}
	}

	dir := t.TempDir()
	fixturePath := filepath.Join(dir, "sel.json")
	if err := os.WriteFile(fixturePath, []byte(fixture), 0o600); err != nil {
		t.Fatal(err)
	}
	stub := fmt.Sprintf("#!/usr/bin/env bash\nif [ %d -ne 0 ]; then echo 'curl: (7) Failed to connect' >&2; exit %d; fi\ncat %q\n",
		curlExit, curlExit, fixturePath)
	if err := os.WriteFile(filepath.Join(dir, "curl"), []byte(stub), 0o700); err != nil { // #nosec G306 -- test stub must be executable
		t.Fatal(err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	cmd := exec.CommandContext(ctx, "bash", "idrac_check.sh")
	cmd.Env = append(os.Environ(),
		"PATH="+dir+string(os.PathListSeparator)+os.Getenv("PATH"),
		"IDRAC=10.0.0.1",
		"IDRAC_USER=mon",
		"IDRAC_PASS=pw",
	)
	out, err := cmd.CombinedOutput()
	exitCode := 0
	if err != nil {
		exitErr, ok := err.(*exec.ExitError)
		if !ok {
			t.Fatalf("run script: %v (%s)", err, out)
		}
		exitCode = exitErr.ExitCode()
	}
	return strings.TrimSpace(string(out)), exitCode
}

func TestIDRACScriptDetectsRecentFatalError(t *testing.T) {
	t.Parallel()
	recent := time.Now().Add(-5 * time.Minute).Format(time.RFC3339)
	out, code := runIDRACScript(t, selFixture(recent), 0)
	if code != 2 {
		t.Fatalf("exit = %d (%s), want 2", code, out)
	}
	if !strings.HasPrefix(out, "CRASH ") || !strings.Contains(out, "fatal error was detected") {
		t.Fatalf("output = %q", out)
	}
}

func TestIDRACScriptIgnoresOldFatalError(t *testing.T) {
	t.Parallel()
	old := time.Now().Add(-48 * time.Hour).Format(time.RFC3339)
	out, code := runIDRACScript(t, selFixture(old), 0)
	if code != 0 || out != "OK" {
		t.Fatalf("exit = %d, output = %q, want 0/OK", code, out)
	}
}

func TestIDRACScriptIgnoresNonFatalCritical(t *testing.T) {
	t.Parallel()
	// Only the recovered temperature excursion is present: Critical severity
	// but not a fatal/halting error, so it must not raise a crash.
	fixture := `{
  "Members": [
    {"Id": "4", "Created": "2026-06-18T01:27:24-05:00", "Severity": "Critical",
     "Message": "CPU 2 temperature is greater than the upper critical threshold."}
  ]
}`
	out, code := runIDRACScript(t, fixture, 0)
	if code != 0 || out != "OK" {
		t.Fatalf("exit = %d, output = %q, want 0/OK", code, out)
	}
}

func TestIDRACScriptReportsCheckFailure(t *testing.T) {
	t.Parallel()
	out, code := runIDRACScript(t, "", 7)
	if code == 0 || code == 2 {
		t.Fatalf("exit = %d (%s), want a plain failure", code, out)
	}
	if !strings.Contains(out, "Failed to connect") {
		t.Fatalf("output = %q", out)
	}
}

// TestIDRACScriptIsReadOnly pins the script's core guarantee: it must never
// contain a Redfish action or any non-GET HTTP invocation.
func TestIDRACScriptIsReadOnly(t *testing.T) {
	t.Parallel()
	script, err := os.ReadFile("idrac_check.sh")
	if err != nil {
		t.Fatal(err)
	}
	text := string(script)
	for _, forbidden := range []string{
		"ComputerSystem.Reset", "ForceOff", "ForceRestart", "GracefulShutdown",
		"-X POST", "--request POST", "-X PATCH", "--request PATCH",
		"-X DELETE", "--request DELETE", "--data",
	} {
		if strings.Contains(text, forbidden) {
			t.Fatalf("script contains forbidden token %q", forbidden)
		}
	}
	if !strings.Contains(text, "--request GET") {
		t.Fatal("script must pin --request GET explicitly")
	}
}
