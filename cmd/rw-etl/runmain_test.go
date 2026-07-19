package main

import (
	"os"
	"testing"
)

// TestRunMainPropagatesError pins the exit-path contract: runMain surfaces
// failures as errors (so main alone decides the exit code after every
// deferred cleanup has run) instead of exiting from library code. Failing
// before wiring also proves the default Prometheus registry is untouched on
// the error path (a second registration would panic).
func TestRunMainPropagatesError(t *testing.T) {
	// Blank out the required RPC endpoint so the typed config load fails
	// deterministically regardless of the developer environment.
	t.Setenv("RPC_URL", "")
	if err := runMain(); err == nil {
		t.Fatal("runMain() = nil, want configuration error")
	}
}

// TestMainExitsNonZeroOnError drives the real entrypoint through its failure
// arm via the stubbed exit seam.
func TestMainExitsNonZeroOnError(t *testing.T) {
	t.Setenv("RPC_URL", "")
	oldArgs := os.Args
	t.Cleanup(func() { os.Args = oldArgs; osExit = os.Exit })
	os.Args = []string{"rw-etl"}
	code := 0
	osExit = func(c int) { code = c }

	main()

	if code != 1 {
		t.Fatalf("main() exit code = %d, want 1", code)
	}
}
