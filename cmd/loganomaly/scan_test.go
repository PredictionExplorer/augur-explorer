package main

import (
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestScan(t *testing.T) {
	t.Parallel()

	writeLog := func(t *testing.T, content string) string {
		t.Helper()
		path := filepath.Join(t.TempDir(), "access.log")
		if err := os.WriteFile(path, []byte(content), 0o600); err != nil {
			t.Fatal(err)
		}
		return path
	}

	t.Run("mixed formats and noise", func(t *testing.T) {
		t.Parallel()
		log := strings.Join([]string{
			`[GIN] 2026/07/03 - 13:21:06 | 500 |   50.7ms | 69.10.55.2 | GET  "/api/legacy"`,
			`[GIN] 2026/07/03 - 13:21:07 | 404 |    1.2ms | 69.10.55.2 | GET  "/wp-admin"`,
			`time=2026-07-09T02:42:09.121-05:00 level=INFO msg=request method=GET path="/api/x" status=502 bytes=11 duration_ms=1.2 ip=1.2.3.4`,
			`time=2026-07-09T02:42:10.121-05:00 level=INFO msg=request method=GET path="/api/ok" status=200 bytes=11 duration_ms=0.4 ip=1.2.3.4`,
			`random noise line`,
		}, "\n") + "\n"

		anomalies, err := scan(writeLog(t, log), 500, 50)
		if err != nil {
			t.Fatal(err)
		}
		if len(anomalies) != 2 {
			t.Fatalf("anomalies = %v", anomalies)
		}
		if !strings.Contains(anomalies[0], "500") || !strings.Contains(anomalies[0], "/api/legacy") {
			t.Fatalf("anomalies[0] = %q", anomalies[0])
		}
		if !strings.Contains(anomalies[1], "502") || !strings.Contains(anomalies[1], "/api/x") {
			t.Fatalf("anomalies[1] = %q", anomalies[1])
		}
	})

	t.Run("keep bounds the ring", func(t *testing.T) {
		t.Parallel()
		var lines []string
		for i := 0; i < 10; i++ {
			lines = append(lines,
				`time=2026-07-09T02:42:09.121-05:00 level=INFO msg=request method=GET path="/api/`+strings.Repeat("x", i+1)+`" status=500 duration_ms=1 ip=1.1.1.1`)
		}
		anomalies, err := scan(writeLog(t, strings.Join(lines, "\n")), 500, 3)
		if err != nil {
			t.Fatal(err)
		}
		if len(anomalies) != 3 {
			t.Fatalf("anomalies = %v", anomalies)
		}
		// The most recent three survive, oldest first.
		if !strings.Contains(anomalies[0], "/api/xxxxxxxx") || !strings.Contains(anomalies[2], "/api/xxxxxxxxxx") {
			t.Fatalf("anomalies = %v", anomalies)
		}
	})

	t.Run("missing file", func(t *testing.T) {
		t.Parallel()
		if _, err := scan(filepath.Join(t.TempDir(), "missing.log"), 500, 10); err == nil {
			t.Fatal("missing file must error")
		}
	})

	t.Run("oversized line", func(t *testing.T) {
		t.Parallel()
		big := "panic: " + strings.Repeat("x", 2*1024*1024)
		if _, err := scan(writeLog(t, big), 500, 10); err == nil {
			t.Fatal("a >1MiB line must surface the scanner error")
		}
	})

	t.Run("empty log", func(t *testing.T) {
		t.Parallel()
		anomalies, err := scan(writeLog(t, ""), 500, 10)
		if err != nil || len(anomalies) != 0 {
			t.Fatalf("anomalies = %v, err = %v", anomalies, err)
		}
	})
}

func TestWriteOut(t *testing.T) {
	t.Parallel()

	t.Run("writes lines and replaces previous file", func(t *testing.T) {
		t.Parallel()
		path := filepath.Join(t.TempDir(), "sub", "anomalies.log")

		if err := writeOut(path, []string{"first", "second"}); err != nil {
			t.Fatal(err)
		}
		data, err := os.ReadFile(path) //nolint:gosec // test path under t.TempDir
		if err != nil {
			t.Fatal(err)
		}
		if string(data) != "first\nsecond\n" {
			t.Fatalf("content = %q", data)
		}

		// A rewrite fully replaces the file and leaves no temp file behind.
		if err := writeOut(path, []string{"only"}); err != nil {
			t.Fatal(err)
		}
		data, err = os.ReadFile(path) //nolint:gosec // test path under t.TempDir
		if err != nil {
			t.Fatal(err)
		}
		if string(data) != "only\n" {
			t.Fatalf("content = %q", data)
		}
		if _, err := os.Stat(path + ".tmp"); !os.IsNotExist(err) {
			t.Fatalf("temp file left behind: %v", err)
		}
	})

	t.Run("empty anomaly list writes empty file", func(t *testing.T) {
		t.Parallel()
		path := filepath.Join(t.TempDir(), "anomalies.log")
		if err := writeOut(path, nil); err != nil {
			t.Fatal(err)
		}
		data, err := os.ReadFile(path) //nolint:gosec // test path under t.TempDir
		if err != nil {
			t.Fatal(err)
		}
		if len(data) != 0 {
			t.Fatalf("content = %q", data)
		}
	})

	t.Run("unwritable directory", func(t *testing.T) {
		t.Parallel()
		if err := writeOut("/dev/null/nope/anomalies.log", []string{"x"}); err == nil {
			t.Fatal("unwritable dir must error")
		}
	})
}
