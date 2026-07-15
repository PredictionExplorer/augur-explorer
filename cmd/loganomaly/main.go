// Command loganomaly scans the websrv access log and writes a compact list
// of the most recent anomalies to an output file. It understands the current
// slog formats — JSON records (production LOG_FORMAT=json under journald)
// and text records (msg=request access lines, msg="panic recovered") — plus
// the legacy [GIN] format still present in older log files.
//
// An "anomaly" is a genuine server-side problem, not routine bot noise:
//   - any request logged with an HTTP status >= 500 (configurable), and
//   - any recovered or fatal Go panic.
//
// Ordinary 4xx responses (404 probes for /.env, /wp-admin, etc.) are ignored.
//
// It is meant to run periodically on the production host (e.g. cosmic1) via
// cron, scanning a journald export (journalctl -u rwcg-apiserver@cosmic
// -o cat > file, see docs/operations.md) or a legacy capture file, writing
// to $HOME/ae_logs/webserver_anomalies.log, which srvmonitor then fetches
// via scp and displays.
//
// Usage:
//
//	loganomaly [-in LOG] [-out FILE] [-min-status 500] [-keep 50]
package main

import (
	"bufio"
	"encoding/json"
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/PredictionExplorer/augur-explorer/internal/version"
)

func main() {
	// Before flag.Parse: --version must win over flag validation.
	if version.HandleFlag(os.Args[1:], os.Stdout) {
		return
	}
	home, _ := os.UserHomeDir()
	defaultIn := filepath.Join(home, "ae_logs", "webserver_cosmic_nohup.log")
	defaultOut := filepath.Join(home, "ae_logs", "webserver_anomalies.log")

	inPath := flag.String("in", defaultIn, "path to the websrv access log to scan")
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
	f, err := os.Open(filepath.Clean(path))
	if err != nil {
		return nil, err
	}
	defer func() { _ = f.Close() }() // best-effort close on read path

	ring := make([]string, 0, keep)
	add := func(s string) {
		ring = append(ring, s)
		if keep > 0 && len(ring) > keep {
			ring = ring[len(ring)-keep:]
		}
	}

	sc := bufio.NewScanner(f)
	// Panic/recovery lines carry stack traces; allow up to 1 MiB per line.
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
	// Structured JSON records are inspected first: their stack fields may
	// contain "panic:" without the record being an anomaly by itself.
	if strings.HasPrefix(line, "{") {
		return parseJSONAnomaly(line, minStatus)
	}
	// Fatal panics are always anomalies, regardless of format.
	if strings.Contains(line, "panic:") {
		return "PANIC | " + strings.TrimSpace(line), true
	}
	if strings.HasPrefix(line, "[GIN]") {
		return parseGinAnomaly(line, minStatus)
	}
	if strings.HasPrefix(line, "time=") {
		return parseSlogAnomaly(line, minStatus)
	}
	return "", false
}

// parseJSONAnomaly handles the slog JSON format production services emit
// (LOG_FORMAT=json; one record per line):
//
//	{"time":"2026-07-13T02:42:09.121-05:00","level":"INFO","msg":"request","method":"GET","path":"/api/x?y=1","status":500,"duration_ms":1.2,...}
//	{"time":"...","level":"ERROR","msg":"panic recovered","method":"GET","path":"/api/x","panic":"...","stack":"..."}
func parseJSONAnomaly(line string, minStatus int) (string, bool) {
	var rec map[string]any
	if err := json.Unmarshal([]byte(line), &rec); err != nil {
		return "", false
	}
	msg, _ := rec["msg"].(string)
	switch msg {
	case "panic recovered":
		return "PANIC | " + strings.TrimSpace(line), true
	case "request":
		status, ok := rec["status"].(float64)
		if !ok || int(status) < minStatus {
			return "", false
		}
		var latency string
		if d, ok := rec["duration_ms"].(float64); ok {
			latency = strconv.FormatFloat(d, 'f', -1, 64) + "ms"
		}
		timeStr, _ := rec["time"].(string)
		method, _ := rec["method"].(string)
		path, _ := rec["path"].(string)
		return fmt.Sprintf("%s | %d | %s %s | %s", timeStr, int(status), method, path, latency), true
	default:
		return "", false
	}
}

// parseGinAnomaly handles the legacy framework access-log format:
//
//	[GIN] 2026/07/03 - 13:21:06 | 500 |   50.7ms | 69.10.55.2 | GET  "/api/..."
func parseGinAnomaly(line string, minStatus int) (string, bool) {
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

// parseSlogAnomaly handles the slog text format the stdlib-router websrv
// emits (internal/api/common AccessLog and Recovery middleware):
//
//	time=2026-07-09T02:42:09.121-05:00 level=INFO msg=request method=GET path="/api/x?y=1" route=... status=500 bytes=11 duration_ms=1.2 ip=1.2.3.4
//	time=2026-07-09T02:42:09.121-05:00 level=ERROR msg="panic recovered" method=GET path=/api/x panic=... stack="..."
func parseSlogAnomaly(line string, minStatus int) (string, bool) {
	kv := parseLogfmt(line)
	switch kv["msg"] {
	case "panic recovered":
		return "PANIC | " + strings.TrimSpace(line), true
	case "request":
		status, err := strconv.Atoi(kv["status"])
		if err != nil || status < minStatus {
			return "", false
		}
		latency := kv["duration_ms"]
		if latency != "" {
			latency += "ms"
		}
		return fmt.Sprintf("%s | %d | %s %s | %s",
			kv["time"], status, kv["method"], kv["path"], latency), true
	default:
		return "", false
	}
}

// parseLogfmt splits a slog text line into key/value pairs. Values may be
// bare tokens or Go-quoted strings (as slog emits them); malformed input
// never panics — unparsable segments are skipped.
func parseLogfmt(line string) map[string]string {
	kv := make(map[string]string)
	i := 0
	for i < len(line) {
		for i < len(line) && line[i] == ' ' {
			i++
		}
		start := i
		for i < len(line) && line[i] != '=' && line[i] != ' ' {
			i++
		}
		if i >= len(line) || line[i] != '=' {
			continue // bare token without '='; the space skip advances i
		}
		key := line[start:i]
		i++ // consume '='

		var val string
		if i < len(line) && line[i] == '"' {
			j := i + 1
			for j < len(line) && line[j] != '"' {
				if line[j] == '\\' {
					j++ // skip the escaped byte
				}
				j++
			}
			if j >= len(line) {
				val = line[i+1:] // unterminated quote: take the rest raw
				i = len(line)
			} else {
				if unquoted, err := strconv.Unquote(line[i : j+1]); err == nil {
					val = unquoted
				} else {
					val = line[i+1 : j]
				}
				i = j + 1
			}
		} else {
			j := i
			for j < len(line) && line[j] != ' ' {
				j++
			}
			val = line[i:j]
			i = j
		}
		if key != "" {
			kv[key] = val
		}
	}
	return kv
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
		if err := os.MkdirAll(dir, 0o750); err != nil {
			return err
		}
	}

	tmp := path + ".tmp"
	f, err := os.Create(filepath.Clean(tmp))
	if err != nil {
		return err
	}

	w := bufio.NewWriter(f)
	for _, l := range lines {
		if _, err := fmt.Fprintln(w, l); err != nil {
			_ = f.Close() // best-effort cleanup on error path
			return err
		}
	}
	if err := w.Flush(); err != nil {
		_ = f.Close() // best-effort cleanup on error path
		return err
	}
	if err := f.Close(); err != nil {
		return err
	}
	return os.Rename(tmp, path)
}
