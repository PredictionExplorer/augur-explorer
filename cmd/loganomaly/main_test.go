package main

import (
	"strings"
	"testing"
)

func TestParseAnomaly(t *testing.T) {
	cases := []struct {
		name      string
		line      string
		minStatus int
		want      string
		wantOK    bool
	}{
		{
			name:      "gin 500",
			line:      `[GIN] 2026/07/03 - 13:21:06 | 500 |   50.7ms | 69.10.55.2 | GET  "/api/foo"`,
			minStatus: 500,
			want:      "2026/07/03 13:21:06 | 500 | GET /api/foo | 50.7ms",
			wantOK:    true,
		},
		{
			name:      "gin 404 below threshold",
			line:      `[GIN] 2026/07/03 - 13:21:06 | 404 |   1.2ms | 1.2.3.4 | GET  "/wp-admin"`,
			minStatus: 500,
			wantOK:    false,
		},
		{
			name:      "gin 404 with lowered threshold",
			line:      `[GIN] 2026/07/03 - 13:21:06 | 404 |   1.2ms | 1.2.3.4 | GET  "/wp-admin"`,
			minStatus: 400,
			want:      "2026/07/03 13:21:06 | 404 | GET /wp-admin | 1.2ms",
			wantOK:    true,
		},
		{
			name:      "panic line",
			line:      "  panic: runtime error: index out of range",
			minStatus: 500,
			want:      "PANIC | panic: runtime error: index out of range",
			wantOK:    true,
		},
		{
			name:      "path containing pipe is rejoined",
			line:      `[GIN] 2026/07/03 - 13:21:06 | 502 |   9ms | 1.2.3.4 | GET  "/api/a|b"`,
			minStatus: 500,
			want:      "2026/07/03 13:21:06 | 502 | GET /api/a|b | 9ms",
			wantOK:    true,
		},
		{name: "non-gin line", line: "hello world", minStatus: 500, wantOK: false},
		{name: "gin line too few fields", line: "[GIN] x | 500", minStatus: 500, wantOK: false},
		{name: "gin line bad status", line: "[GIN] ts | NaN | 1ms | ip | GET \"/\"", minStatus: 500, wantOK: false},
		{name: "empty", line: "", minStatus: 500, wantOK: false},
		{
			name:      "slog 500",
			line:      `time=2026-07-09T02:42:09.121-05:00 level=INFO msg=request method=GET path="/api/foo?limit=2" route=/api/foo status=500 bytes=11 duration_ms=1.25 ip=10.0.0.1`,
			minStatus: 500,
			want:      "2026-07-09T02:42:09.121-05:00 | 500 | GET /api/foo?limit=2 | 1.25ms",
			wantOK:    true,
		},
		{
			name:      "slog 200 below threshold",
			line:      `time=2026-07-09T02:42:09.121-05:00 level=INFO msg=request method=GET path=/ok status=200 duration_ms=0.5 ip=10.0.0.1`,
			minStatus: 500,
			wantOK:    false,
		},
		{
			name:      "slog 404 with lowered threshold",
			line:      `time=2026-07-09T02:42:09.121-05:00 level=INFO msg=request method=GET path=/wp-admin route="" status=404 duration_ms=0.1 ip=10.0.0.1`,
			minStatus: 400,
			want:      "2026-07-09T02:42:09.121-05:00 | 404 | GET /wp-admin | 0.1ms",
			wantOK:    true,
		},
		{
			name:      "slog recovered panic",
			line:      `time=2026-07-09T02:42:09.121-05:00 level=ERROR msg="panic recovered" method=GET path=/boom panic=kaboom`,
			minStatus: 500,
			want:      `PANIC | time=2026-07-09T02:42:09.121-05:00 level=ERROR msg="panic recovered" method=GET path=/boom panic=kaboom`,
			wantOK:    true,
		},
		{
			name:      "slog non-request message ignored",
			line:      `time=2026-07-09T02:42:09.121-05:00 level=INFO msg=startup listeners=2`,
			minStatus: 500,
			wantOK:    false,
		},
		{
			name:      "json 500",
			line:      `{"time":"2026-07-13T02:42:09.121-05:00","level":"INFO","msg":"request","method":"GET","path":"/api/foo?limit=2","route":"/api/foo","status":500,"bytes":11,"duration_ms":1.25,"ip":"10.0.0.1"}`,
			minStatus: 500,
			want:      "2026-07-13T02:42:09.121-05:00 | 500 | GET /api/foo?limit=2 | 1.25ms",
			wantOK:    true,
		},
		{
			name:      "json 200 below threshold",
			line:      `{"time":"2026-07-13T02:42:09.121-05:00","level":"INFO","msg":"request","method":"GET","path":"/ok","status":200,"duration_ms":0.5}`,
			minStatus: 500,
			wantOK:    false,
		},
		{
			name:      "json 404 with lowered threshold",
			line:      `{"time":"2026-07-13T02:42:09.121-05:00","level":"INFO","msg":"request","method":"GET","path":"/wp-admin","status":404,"duration_ms":0.1}`,
			minStatus: 400,
			want:      "2026-07-13T02:42:09.121-05:00 | 404 | GET /wp-admin | 0.1ms",
			wantOK:    true,
		},
		{
			name:      "json recovered panic",
			line:      `{"time":"2026-07-13T02:42:09.121-05:00","level":"ERROR","msg":"panic recovered","method":"GET","path":"/boom","panic":"kaboom","stack":"goroutine 1..."}`,
			minStatus: 500,
			want:      `PANIC | {"time":"2026-07-13T02:42:09.121-05:00","level":"ERROR","msg":"panic recovered","method":"GET","path":"/boom","panic":"kaboom","stack":"goroutine 1..."}`,
			wantOK:    true,
		},
		{
			name:      "json record with panic in a field is not a fatal panic",
			line:      `{"time":"t","level":"INFO","msg":"startup","note":"recovered from panic: earlier"}`,
			minStatus: 500,
			wantOK:    false,
		},
		{
			name:      "json non-request message ignored",
			line:      `{"time":"t","level":"INFO","msg":"effective configuration","HTTP_PORT":"8080"}`,
			minStatus: 500,
			wantOK:    false,
		},
		{
			name:      "json request without numeric status ignored",
			line:      `{"time":"t","msg":"request","status":"broken"}`,
			minStatus: 500,
			wantOK:    false,
		},
		{name: "malformed json ignored", line: `{"msg":"request",`, minStatus: 0, wantOK: false},
	}
	for _, tc := range cases {
		got, ok := parseAnomaly(tc.line, tc.minStatus)
		if ok != tc.wantOK {
			t.Errorf("%s: parseAnomaly ok = %v, want %v", tc.name, ok, tc.wantOK)
			continue
		}
		if ok && got != tc.want {
			t.Errorf("%s: parseAnomaly = %q, want %q", tc.name, got, tc.want)
		}
	}
}

func TestSplitRequest(t *testing.T) {
	cases := []struct {
		in, wantMethod, wantPath string
	}{
		{`GET      "/api/foo"`, "GET", "/api/foo"},
		{`POST "/x"`, "POST", "/x"},
		{``, "", ""},
		{`GET`, "GET", ""},
	}
	for _, tc := range cases {
		m, p := splitRequest(tc.in)
		if m != tc.wantMethod || p != tc.wantPath {
			t.Errorf("splitRequest(%q) = (%q,%q), want (%q,%q)", tc.in, m, p, tc.wantMethod, tc.wantPath)
		}
	}
}

func FuzzLogAnomalyScan(f *testing.F) {
	f.Add(`[GIN] 2026/07/03 - 13:21:06 | 500 |   50.7ms | 69.10.55.2 | GET  "/api/foo"`, 500)
	f.Add("panic: boom", 500)
	f.Add("[GIN] |||||", 500)
	f.Add("", 0)
	f.Add("[GIN] a | 99999999999999999999 | b | c | d", 500)
	f.Add(`time=x level=INFO msg=request method=GET path="/a b" status=502 duration_ms=9 ip=1.2.3.4`, 500)
	f.Add(`time=x msg="panic recovered" path="unterminated`, 500)
	f.Add(`time=x msg=request status=`, 500)
	f.Add("time== = =\"", 500)
	f.Add(`{"time":"t","msg":"request","status":503,"method":"GET","path":"/x","duration_ms":1.5}`, 500)
	f.Add(`{"time":"t","msg":"panic recovered","stack":"panic: deep"}`, 500)
	f.Add(`{"msg":"request","status":1e309}`, 500)
	f.Add(`{"msg":"request","status":{}}`, 500)
	f.Add(`{`, 500)
	f.Fuzz(func(t *testing.T, line string, minStatus int) {
		rec, ok := parseAnomaly(line, minStatus)
		if !ok {
			if rec != "" {
				t.Fatalf("parseAnomaly(%q) rejected but returned text %q", line, rec)
			}
			return
		}
		if rec == "" {
			t.Fatalf("parseAnomaly(%q) accepted but returned empty record", line)
		}
		switch {
		case strings.HasPrefix(line, "{"):
			// JSON records are accepted only through their msg dispatch; a
			// "panic:" substring inside a field must not force a PANIC record.
			if strings.HasPrefix(rec, "PANIC | ") && !strings.Contains(line, `"panic recovered"`) {
				t.Fatalf("JSON line %q produced an unexpected PANIC record %q", line, rec)
			}
		case strings.Contains(line, "panic:"):
			// A fatal panic line is always reported as such.
			if !strings.HasPrefix(rec, "PANIC | ") {
				t.Fatalf("panic line %q produced non-PANIC record %q", line, rec)
			}
		default:
			// Anything else must come from a legacy [GIN] line or a slog
			// text line.
			if !strings.HasPrefix(line, "[GIN]") && !strings.HasPrefix(line, "time=") {
				t.Fatalf("unrecognized line %q was accepted: %q", line, rec)
			}
		}
	})
}
