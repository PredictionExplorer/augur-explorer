package main

import (
	"os"
	"testing"
)

// blankImggenEnv blanks the required generator endpoints so the typed config
// load fails deterministically regardless of the developer environment.
func blankImggenEnv(t *testing.T) {
	t.Helper()
	t.Setenv("IM_REQUEST_URL", "")
	t.Setenv("IM_IMAGE_URL", "")
	t.Setenv("IM_VIDEO_URL", "")
}

// TestRunMainPropagatesError pins the exit-path contract: runMain surfaces
// failures as errors (so main alone decides the exit code after every
// deferred cleanup has run) instead of exiting from library code.
func TestRunMainPropagatesError(t *testing.T) {
	blankImggenEnv(t)
	if err := runMain(nil); err == nil {
		t.Fatal("runMain(nil) = nil, want configuration error")
	}
}

// TestMainExitsNonZeroOnError drives the real entrypoint through its failure
// arm via the stubbed exit seam.
func TestMainExitsNonZeroOnError(t *testing.T) {
	blankImggenEnv(t)
	oldArgs := os.Args
	t.Cleanup(func() { os.Args = oldArgs; osExit = os.Exit })
	os.Args = []string{"imggen-monitor"}
	code := 0
	osExit = func(c int) { code = c }

	main()

	if code != 1 {
		t.Fatalf("main() exit code = %d, want 1", code)
	}
}
