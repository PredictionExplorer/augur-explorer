package main

import (
	"os"
	"testing"
)

// TestRunMainPropagatesError pins the exit-path contract: runMain surfaces
// failures as errors (so main alone decides the exit code after every
// deferred cleanup has run) instead of exiting from library code.
func TestRunMainPropagatesError(t *testing.T) {
	// Blank out the required RPC endpoint so the typed config load fails
	// deterministically regardless of the developer environment.
	t.Setenv("RPC_URL", "")
	if err := runMain(true, false); err == nil {
		t.Fatal("runMain(true, false) = nil, want configuration error")
	}
}

// TestMainExitsNonZeroOnError drives the real entrypoint through both
// failure arms — missing mode flags, then a config failure with --twitter —
// via the stubbed exit seam.
func TestMainExitsNonZeroOnError(t *testing.T) {
	t.Setenv("RPC_URL", "")
	oldArgs := os.Args
	t.Cleanup(func() { os.Args = oldArgs; osExit = os.Exit })
	code := 0
	osExit = func(c int) { code = c }

	os.Args = []string{"notibot"}
	main()
	if code != 1 {
		t.Fatalf("main() without mode flags: exit code = %d, want 1", code)
	}

	code = 0
	os.Args = []string{"notibot", "--twitter"}
	main()
	if code != 1 {
		t.Fatalf("main() with config failure: exit code = %d, want 1", code)
	}
}
