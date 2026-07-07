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
		// A panic line is always reported as such; anything else must have
		// come from a [GIN] access line.
		if strings.Contains(line, "panic:") {
			if !strings.HasPrefix(rec, "PANIC | ") {
				t.Fatalf("panic line %q produced non-PANIC record %q", line, rec)
			}
		} else if !strings.HasPrefix(line, "[GIN]") {
			t.Fatalf("non-GIN line %q was accepted: %q", line, rec)
		}
	})
}
