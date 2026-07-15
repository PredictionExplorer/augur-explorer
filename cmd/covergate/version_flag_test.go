package main

import (
	"io"
	"os"
	"strings"
	"testing"
)

// TestMainVersionFlag pins the real entrypoint contract: --version prints
// the build identity and returns before any wiring (no config, database or
// side effects — the test would hang or fail otherwise).
func TestMainVersionFlag(t *testing.T) {
	oldArgs, oldStdout := os.Args, os.Stdout
	t.Cleanup(func() { os.Args, os.Stdout = oldArgs, oldStdout })

	r, w, err := os.Pipe()
	if err != nil {
		t.Fatal(err)
	}
	os.Args = []string{"covergate", "--version"}
	os.Stdout = w

	main()

	_ = w.Close()
	os.Stdout = oldStdout
	out, err := io.ReadAll(r)
	if err != nil {
		t.Fatal(err)
	}
	if !strings.Contains(string(out), "commit") {
		t.Fatalf("--version output = %q, want the build identity line", out)
	}
}
