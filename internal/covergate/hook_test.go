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

func TestCursorCommitHookAllowsBypassWhileDeferred(t *testing.T) {
	t.Parallel()
	repoRoot := covergateRepoRoot(t)
	script := filepath.Join(repoRoot, ".cursor", "hooks", "coverage-commit-gate.sh")
	// #nosec G204 -- the script path is derived from this test's repository.
	command := exec.Command("bash", script)
	command.Dir = repoRoot
	command.Stdin = strings.NewReader(`{"command":"git commit --no-verify -m test"}`)
	output, err := command.CombinedOutput()
	if err != nil {
		t.Fatalf("deferred Cursor hook failed: %v\n%s", err, output)
	}
	if !strings.Contains(string(output), `"permission": "allow"`) {
		t.Fatalf("deferred Cursor hook output = %s", output)
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

func TestNativeCommitGateIsDeferredBeforeNinetyPercent(t *testing.T) {
	t.Parallel()
	repoRoot := covergateRepoRoot(t)
	script := filepath.Join(repoRoot, ".githooks", "pre-commit")
	// #nosec G204 -- the script path is derived from this test's repository.
	command := exec.Command("bash", script)
	command.Dir = repoRoot
	output, err := command.CombinedOutput()
	if err != nil {
		t.Fatalf("native hook failed before activation: %v\n%s", err, output)
	}
	if !strings.Contains(string(output), "deferred until handwritten internal coverage reaches 90%") {
		t.Fatalf("native hook output = %s", output)
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
