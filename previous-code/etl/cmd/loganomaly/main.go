// Command loganomaly scans the websrv (Gin) access log and writes a compact
// list of the most recent anomalies to an output file.
//
// An "anomaly" is a genuine server-side problem, not routine bot noise:
//   - any Gin request logged with an HTTP status >= 500 (configurable), and
//   - any line containing a Go panic.
//
// Ordinary 4xx responses (404 probes for /.env, /wp-admin, etc.) are ignored.
//
// It is meant to run periodically on the production host (e.g. cosmic1) via
// cron or a run-loop, writing to $HOME/ae_logs/webserver_anomalies.log, which
// srvmonitor then fetches via scp and displays.
//
// Usage:
//
//	loganomaly [-in LOG] [-out FILE] [-min-status 500] [-keep 50]
package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

func main() {
	home, _ := os.UserHomeDir()
	defaultIn := filepath.Join(home, "ae_logs", "webserver_cosmic_nohup.log")
	defaultOut := filepath.Join(home, "ae_logs", "webserver_anomalies.log")

	inPath := flag.String("in", defaultIn, "path to the Gin access log to scan")
	outPath := flag.String("out", defaultOut, "path to write the anomalies file")
	minStatus := flag.Int("min-status", 500, "minimum HTTP status treated as an anomaly")
	keep := flag.Int("keep", 50, "max number of most-recent anomalies to keep")
	flag.Parse()

	anomalies, err := scan(*inPath, *minStatus, *keep)
	if err != nil {
		fmt.Fprintf(os.Stderr, "loganomaly: %v\n", err)
		os.Exit(1)
	}

	if err := writeOut(*outPath, anomalies); err != nil {
		fmt.Fprintf(os.Stderr, "loganomaly: %v\n", err)
		os.Exit(1)
	}
}

// scan reads the log once and returns up to keep most-recent anomalies in
// chronological order (oldest first).
func scan(path string, minStatus, keep int) ([]string, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	ring := make([]string, 0, keep)
	add := func(s string) {
		ring = append(ring, s)
		if keep > 0 && len(ring) > keep {
			ring = ring[len(ring)-keep:]
		}
	}

	sc := bufio.NewScanner(f)
	// Gin panic/recovery lines can be long; allow up to 1 MiB per line.
	sc.Buffer(make([]byte, 0, 64*1024), 1024*1024)
	for sc.Scan() {
		if rec, ok := parseAnomaly(sc.Text(), minStatus); ok {
			add(rec)
		}
	}
	if err := sc.Err(); err != nil {
		return nil, err
	}
	return ring, nil
}

// parseAnomaly returns a compact one-line record when the log line represents
// an anomaly, and false otherwise.
func parseAnomaly(line string, minStatus int) (string, bool) {
	// Panics are always anomalies, regardless of format.
	if strings.Contains(line, "panic:") {
		return "PANIC | " + strings.TrimSpace(line), true
	}

	// Gin access lines look like:
	//   [GIN] 2026/07/03 - 13:21:06 | 500 |   50.7ms | 69.10.55.2 | GET  "/api/..."
	if !strings.HasPrefix(line, "[GIN]") {
		return "", false
	}
	parts := strings.Split(line, "|")
	if len(parts) < 5 {
		return "", false
	}

	status, err := strconv.Atoi(strings.TrimSpace(parts[1]))
	if err != nil || status < minStatus {
		return "", false
	}

	// Timestamp: strip "[GIN] " and collapse " - " into a single space.
	ts := strings.TrimSpace(strings.TrimPrefix(parts[0], "[GIN]"))
	ts = strings.Replace(ts, " - ", " ", 1)

	latency := strings.TrimSpace(parts[2])

	// The request field is the remainder; rejoin in case a path contained '|'.
	method, path := splitRequest(strings.TrimSpace(strings.Join(parts[4:], "|")))

	return fmt.Sprintf("%s | %d | %s %s | %s", ts, status, method, path, latency), true
}

// splitRequest parses `GET      "/api/..."` into method and unquoted path.
func splitRequest(s string) (method, path string) {
	fields := strings.Fields(s)
	if len(fields) == 0 {
		return "", s
	}
	method = fields[0]
	rest := strings.TrimSpace(strings.TrimPrefix(s, method))
	return method, strings.Trim(rest, "\"")
}

// writeOut writes the anomalies atomically (temp file + rename) so a concurrent
// scp never observes a half-written file.
func writeOut(path string, lines []string) error {
	if dir := filepath.Dir(path); dir != "" {
		if err := os.MkdirAll(dir, 0o755); err != nil {
			return err
		}
	}

	tmp := path + ".tmp"
	f, err := os.Create(tmp)
	if err != nil {
		return err
	}

	w := bufio.NewWriter(f)
	for _, l := range lines {
		if _, err := fmt.Fprintln(w, l); err != nil {
			f.Close()
			return err
		}
	}
	if err := w.Flush(); err != nil {
		f.Close()
		return err
	}
	if err := f.Close(); err != nil {
		return err
	}
	return os.Rename(tmp, path)
}
