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
	passed, err := runMain(nil)
	if passed || err == nil || !strings.Contains(err.Error(), "--input is required") {
		t.Fatalf("runMain(nil) = (%v, %v), want missing --input error", passed, err)
	}
}

// TestMainExitsNonZeroOnError drives the real entrypoint through its failure
// arm (missing --input) via the stubbed exit seam.
func TestMainExitsNonZeroOnError(t *testing.T) {
	oldArgs := os.Args
	t.Cleanup(func() { os.Args = oldArgs; osExit = os.Exit })
	os.Args = []string{"freezer-verify"}
	code := 0
	osExit = func(c int) { code = c }

	main()

	if code != 1 {
		t.Fatalf("main() exit code = %d, want 1", code)
	}
}
