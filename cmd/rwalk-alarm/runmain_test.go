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
	err := runMain(nil)
	if err == nil || !strings.Contains(err.Error(), "usage: rwalk-alarm") {
		t.Fatalf("runMain(nil) = %v, want usage error", err)
	}
}

// TestMainExitsNonZeroOnError drives the real entrypoint through its failure
// arm (no url-list argument) via the stubbed exit seam.
func TestMainExitsNonZeroOnError(t *testing.T) {
	oldArgs := os.Args
	t.Cleanup(func() { os.Args = oldArgs; osExit = os.Exit })
	os.Args = []string{"rwalk-alarm"}
	code := 0
	osExit = func(c int) { code = c }

	main()

	if code != 1 {
		t.Fatalf("main() exit code = %d, want 1", code)
	}
}
