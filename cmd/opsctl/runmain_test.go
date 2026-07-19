package main

import (
	"os"
	"strings"
	"testing"
)

// TestRunMainPropagatesError pins the exit-path contract: runMain surfaces
// failures as errors (so main alone decides the exit code after every
// deferred cleanup has run) instead of exiting from library code.
func TestRunMainPropagatesError(t *testing.T) {
	err := runMain([]string{"no-such-subcommand"})
	if err == nil || !strings.Contains(err.Error(), "unknown command") {
		t.Fatalf("runMain(no-such-subcommand) = %v, want unknown-command error", err)
	}
}

// TestMainExitsNonZeroOnError drives the real entrypoint through its failure
// arm (unknown subcommand) via the stubbed exit seam.
func TestMainExitsNonZeroOnError(t *testing.T) {
	oldArgs := os.Args
	t.Cleanup(func() { os.Args = oldArgs; osExit = os.Exit })
	os.Args = []string{"opsctl", "no-such-subcommand"}
	code := 0
	osExit = func(c int) { code = c }

	main()

	if code != 1 {
		t.Fatalf("main() exit code = %d, want 1", code)
	}
}
