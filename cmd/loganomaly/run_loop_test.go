package main

import (
	"bytes"
	"context"
	"errors"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"testing"
	"time"
)

func copyRunLoop(t *testing.T) string {
	t.Helper()
	data, err := os.ReadFile("run-loop.sh")
	if err != nil {
		t.Fatal(err)
	}
	path := filepath.Join(t.TempDir(), "run-loop.sh")
	if err := os.WriteFile(path, data, 0o600); err != nil { //nolint:gosec // fixed basename under t.TempDir
		t.Fatal(err)
	}
	if err := os.Chmod(path, 0o755); err != nil { //nolint:gosec // copied script must be executable for behavior tests
		t.Fatal(err)
	}
	return path
}

func runLoopEnv(t *testing.T, values ...string) []string {
	t.Helper()
	env := make([]string, 0, len(os.Environ())+len(values)+1)
	for _, entry := range os.Environ() {
		name, _, _ := strings.Cut(entry, "=")
		if strings.HasPrefix(name, "LOGANOMALY_") || name == "CAPTURE" || name == "HOME" {
			continue
		}
		env = append(env, entry)
	}
	env = append(env, "HOME="+t.TempDir())
	return append(env, values...)
}

func executeRunLoop(t *testing.T, script, workingDir string, env ...string) (string, error) {
	t.Helper()
	ctx, cancel := context.WithTimeout(t.Context(), 5*time.Second)
	defer cancel()
	var stderr bytes.Buffer
	cmd := exec.CommandContext(ctx, "/bin/bash", script) //nolint:gosec // fixed shell executes the copied repository script
	cmd.Dir = workingDir
	cmd.Env = runLoopEnv(t, env...)
	cmd.Stderr = &stderr
	err := cmd.Run()
	if errors.Is(ctx.Err(), context.DeadlineExceeded) {
		t.Fatalf("run-loop did not terminate; stderr: %s", stderr.String())
	}
	return stderr.String(), err
}

func requireExitCode(t *testing.T, err error, want int) {
	t.Helper()
	var exitErr *exec.ExitError
	if !errors.As(err, &exitErr) {
		t.Fatalf("error = %v, want exit status %d", err, want)
	}
	if got := exitErr.ExitCode(); got != want {
		t.Fatalf("exit code = %d, want %d", got, want)
	}
}

func TestRunLoopMissingBinaryExitsNonZero(t *testing.T) {
	t.Parallel()
	script := copyRunLoop(t)
	stderr, err := executeRunLoop(t, script, t.TempDir(), "LOGANOMALY_INTERVAL=0")
	requireExitCode(t, err, 1)
	if !strings.Contains(stderr, "missing or not executable") {
		t.Fatalf("stderr = %q", stderr)
	}
}

func TestRunLoopNonExecutableBinaryExitsNonZero(t *testing.T) {
	t.Parallel()
	script := copyRunLoop(t)
	bin := filepath.Join(t.TempDir(), "loganomaly")
	if err := os.WriteFile(bin, []byte("#!/bin/bash\nexit 0\n"), 0o600); err != nil {
		t.Fatal(err)
	}

	stderr, err := executeRunLoop(t, script, t.TempDir(),
		"LOGANOMALY_BIN="+bin,
		"LOGANOMALY_INTERVAL=0",
	)
	requireExitCode(t, err, 1)
	if !strings.Contains(stderr, "binary '"+bin+"' missing or not executable") {
		t.Fatalf("stderr = %q", stderr)
	}
}

func TestRunLoopResolvesDefaultBinaryRelativeToScript(t *testing.T) {
	t.Parallel()
	script := copyRunLoop(t)
	scriptDir := filepath.Dir(script)
	bin := filepath.Join(scriptDir, "loganomaly")
	capture := filepath.Join(t.TempDir(), "args")
	fake := "#!/bin/bash\nprintf '%s\\n' \"$0\" \"$@\" > \"$CAPTURE\"\nrm -f -- \"$0\"\n"
	if err := os.WriteFile(bin, []byte(fake), 0o600); err != nil {
		t.Fatal(err)
	}
	if err := os.Chmod(bin, 0o755); err != nil { //nolint:gosec // fake binary must be executable to exercise the loop
		t.Fatal(err)
	}

	otherDir := t.TempDir()
	input := filepath.Join(t.TempDir(), "input with spaces.log")
	output := filepath.Join(t.TempDir(), "output with spaces.log")
	stderr, err := executeRunLoop(t, script, otherDir,
		"CAPTURE="+capture,
		"LOGANOMALY_IN="+input,
		"LOGANOMALY_OUT="+output,
		"LOGANOMALY_MIN_STATUS=501",
		"LOGANOMALY_KEEP=7",
		"LOGANOMALY_INTERVAL=0",
	)
	requireExitCode(t, err, 1)
	if !strings.Contains(stderr, "missing or not executable") {
		t.Fatalf("stderr = %q", stderr)
	}

	data, err := os.ReadFile(capture) //nolint:gosec // test path under t.TempDir
	if err != nil {
		t.Fatal(err)
	}
	got := strings.Split(strings.TrimSpace(string(data)), "\n")
	resolvedScriptDir, err := filepath.EvalSymlinks(scriptDir)
	if err != nil {
		t.Fatal(err)
	}
	want := []string{filepath.Join(resolvedScriptDir, "loganomaly"), "-in", input, "-out", output, "-min-status", "501", "-keep", "7"}
	if len(got) != len(want) {
		t.Fatalf("arguments = %q, want %q", got, want)
	}
	for i := range want {
		if got[i] != want[i] {
			t.Fatalf("argument %d = %q, want %q (all: %q)", i, got[i], want[i], got)
		}
	}
}
