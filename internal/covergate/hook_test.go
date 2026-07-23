package covergate

import (
	"bytes"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"strings"
	"testing"
)

func TestCursorCommitHook(t *testing.T) {
	t.Parallel()
	repoRoot := covergateRepoRoot(t)
	script := filepath.Join(repoRoot, ".cursor", "hooks", "coverage-commit-gate.sh")
	tracked := filepath.Join(t.TempDir(), "tracked-hook")
	installed := filepath.Join(t.TempDir(), "installed-hook")
	enabledPolicy := filepath.Join(t.TempDir(), "enabled-policy.json")
	if err := os.WriteFile(tracked, []byte("managed hook\n"), 0o600); err != nil {
		t.Fatal(err)
	}
	policy := strings.NewReplacer(
		`"commitGateEnabled": false`, `"commitGateEnabled": true`,
		`"internalFloor": 70`, `"internalFloor": 90`,
	).Replace(validPolicyJSON)
	if err := os.WriteFile(enabledPolicy, []byte(policy), 0o600); err != nil {
		t.Fatal(err)
	}

	runHook := func(input string) string {
		t.Helper()
		// #nosec G204 -- the script path is derived from this test's repository.
		command := exec.Command("bash", script)
		command.Dir = repoRoot
		command.Env = append(
			os.Environ(),
			"COVERAGE_TRACKED_HOOK="+tracked,
			"COVERAGE_INSTALLED_HOOK="+installed,
			"COVERAGE_POLICY="+enabledPolicy,
		)
		command.Stdin = strings.NewReader(input)
		output, err := command.CombinedOutput()
		if err != nil {
			t.Fatalf("hook failed: %v\n%s", err, output)
		}
		return string(output)
	}

	if output := runHook(`{"command":"git commit -m test"}`); !strings.Contains(output, `"permission": "deny"`) ||
		!strings.Contains(output, "missing or stale") {
		t.Fatalf("missing-hook output = %s", output)
	}

	// #nosec G304 -- tracked is inside this test's temporary directory.
	data, err := os.ReadFile(tracked)
	if err != nil {
		t.Fatal(err)
	}
	// #nosec G703 -- installed is inside this test's temporary directory.
	if err := os.WriteFile(installed, data, 0o600); err != nil {
		t.Fatal(err)
	}
	// #nosec G302 -- the fixture must be executable to model an installed hook.
	if err := os.Chmod(installed, 0o700); err != nil {
		t.Fatal(err)
	}
	if output := runHook(`{"command":"git commit -m test"}`); !strings.Contains(output, `"permission": "allow"`) {
		t.Fatalf("installed-hook output = %s", output)
	}
	if output := runHook(`{"command":"git commit --no-verify -m test"}`); !strings.Contains(output, `"permission": "deny"`) ||
		!strings.Contains(output, "forbids bypassing") {
		t.Fatalf("bypass output = %s", output)
	}
}

func TestCursorCommitHookDeniesBypassUnderRealPolicy(t *testing.T) {
	t.Parallel()
	repoRoot := covergateRepoRoot(t)
	script := filepath.Join(repoRoot, ".cursor", "hooks", "coverage-commit-gate.sh")
	// The tracked policy has the commit gate enabled (activated when the
	// race-enabled handwritten internal coverage first exceeded 90%), so
	// bypassing the pre-commit hook must be denied.
	// #nosec G204 -- the script path is derived from this test's repository.
	command := exec.Command("bash", script)
	command.Dir = repoRoot
	command.Stdin = strings.NewReader(`{"command":"git commit --no-verify -m test"}`)
	output, err := command.CombinedOutput()
	if err != nil {
		t.Fatalf("Cursor hook failed under the real policy: %v\n%s", err, output)
	}
	if !strings.Contains(string(output), `"permission": "deny"`) ||
		!strings.Contains(string(output), "forbids bypassing") {
		t.Fatalf("real-policy bypass output = %s", output)
	}
}

func TestCursorHookConfigurationReferencesFailClosedGate(t *testing.T) {
	t.Parallel()
	repoRoot := covergateRepoRoot(t)
	// #nosec G304 -- the path is derived from this test's repository.
	data, err := os.ReadFile(filepath.Join(repoRoot, ".cursor", "hooks.json"))
	if err != nil {
		t.Fatal(err)
	}
	for _, required := range [][]byte{
		[]byte(`"beforeShellExecution"`),
		[]byte(`"failClosed": true`),
		[]byte(`coverage-commit-gate.sh`),
		[]byte(`git\\s+commit`),
	} {
		if !bytes.Contains(data, required) {
			t.Fatalf("hooks.json missing %q:\n%s", required, data)
		}
	}
}

func TestNativeCommitGateIsEnabledAtNinetyPercent(t *testing.T) {
	t.Parallel()
	repoRoot := covergateRepoRoot(t)
	// The pre-commit hook itself would launch the full coverage run;
	// asserting the policy status pins the activation without that cost.
	command := exec.Command("go", "run", "./cmd/covergate",
		"-policy", "coverage/policy.json", "-commit-status")
	command.Dir = repoRoot
	output, err := command.CombinedOutput()
	if err != nil {
		t.Fatalf("commit-status query failed: %v\n%s", err, output)
	}
	status := strings.TrimSpace(string(output))
	if status != "enabled 90.00" {
		t.Fatalf("commit gate status = %q, want permanently enabled at the 90%% floor", status)
	}
}

func TestNativeCommitHookUsesRaceProfileAndDistinctCacheKey(t *testing.T) {
	t.Parallel()
	repoRoot := covergateRepoRoot(t)
	for path, required := range map[string][]byte{
		".githooks/pre-commit":     []byte("export COVERAGE_RACE=1"),
		"scripts/coverage-gate.sh": []byte(`coverage_race=${COVERAGE_RACE:-0}`),
	} {
		// #nosec G304 -- both paths are fixed repository files.
		data, err := os.ReadFile(filepath.Join(repoRoot, path))
		if err != nil {
			t.Fatalf("reading %s: %v", path, err)
		}
		if !bytes.Contains(data, required) {
			t.Fatalf("%s missing %q", path, required)
		}
	}
}

func covergateRepoRoot(t *testing.T) string {
	t.Helper()
	_, file, _, ok := runtime.Caller(0)
	if !ok {
		t.Fatal("runtime.Caller failed")
	}
	return filepath.Clean(filepath.Join(filepath.Dir(file), "..", ".."))
}
