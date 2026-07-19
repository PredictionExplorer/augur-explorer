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
	setFlags(t, map[string]string{"ancientDir": "", "receiptsCidx": ""})
	err := runMain()
	if err == nil || !strings.Contains(err.Error(), "--ancientDir or --receiptsCidx is required") {
		t.Fatalf("runMain() = %v, want missing-flag error", err)
	}
}

// TestMainExitsNonZeroOnError drives the real entrypoint through its failure
// arm (no freezer location flags) via the stubbed exit seam.
func TestMainExitsNonZeroOnError(t *testing.T) {
	setFlags(t, map[string]string{"ancientDir": "", "receiptsCidx": ""})
	oldArgs := os.Args
	t.Cleanup(func() { os.Args = oldArgs; osExit = os.Exit })
	os.Args = []string{"freezer-scan"}
	code := 0
	osExit = func(c int) { code = c }

	main()

	if code != 1 {
		t.Fatalf("main() exit code = %d, want 1", code)
	}
}
