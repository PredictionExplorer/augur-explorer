package config

import (
	"encoding/json"
	"strings"
	"testing"
)

func TestLogValidate(t *testing.T) {
	t.Parallel()
	if p := (Log{Format: "text", Level: "info"}).validate(); len(p) != 0 {
		t.Errorf("valid Log reported problems: %v", p)
	}
	if p := (Log{Format: "JSON", Level: "WARN"}).validate(); len(p) != 0 {
		t.Errorf("case-insensitive values rejected: %v", p)
	}
	p := (Log{Format: "xml", Level: "loud"}).validate()
	if len(p) != 2 {
		t.Fatalf("got %d problems, want 2: %v", len(p), p)
	}
	if !strings.Contains(p[0], "LOG_LEVEL") || !strings.Contains(p[1], "LOG_FORMAT") {
		t.Errorf("problems misattributed: %v", p)
	}
}

func TestNewLoggerTextFormat(t *testing.T) {
	t.Parallel()
	var sb strings.Builder
	logger := Log{Format: "text", Level: "info"}.NewLogger(&sb)
	logger.Debug("hidden")
	logger.Info("shown", "k", "v")
	out := sb.String()
	if strings.Contains(out, "hidden") {
		t.Errorf("debug record passed an info-level logger: %s", out)
	}
	if !strings.Contains(out, "msg=shown") || !strings.Contains(out, "k=v") {
		t.Errorf("text record malformed: %s", out)
	}
	if strings.HasPrefix(strings.TrimSpace(out), "{") {
		t.Errorf("text format produced JSON: %s", out)
	}
}

func TestNewLoggerJSONFormat(t *testing.T) {
	t.Parallel()
	var sb strings.Builder
	logger := Log{Format: "json", Level: "debug"}.NewLogger(&sb)
	logger.Debug("dbg", "n", 1)
	line := strings.TrimSpace(sb.String())
	var rec map[string]any
	if err := json.Unmarshal([]byte(line), &rec); err != nil {
		t.Fatalf("JSON record does not parse: %v (%s)", err, line)
	}
	if rec["msg"] != "dbg" || rec["level"] != "DEBUG" || rec["n"] != float64(1) {
		t.Errorf("JSON record fields = %v", rec)
	}
}

func TestNewLoggerLevels(t *testing.T) {
	t.Parallel()
	for level, wantWarn := range map[string]bool{"debug": true, "info": true, "warn": true, "error": false} {
		var sb strings.Builder
		logger := Log{Format: "text", Level: level}.NewLogger(&sb)
		logger.Warn("w")
		if got := strings.Contains(sb.String(), "msg=w"); got != wantWarn {
			t.Errorf("level %s: warn shown = %v, want %v", level, got, wantWarn)
		}
	}
}

func TestNewLoggerToleratesInvalidLevel(t *testing.T) {
	t.Parallel()
	var sb strings.Builder
	// Load validates levels; NewLogger falls back to info for robustness.
	logger := Log{Format: "text", Level: "bogus"}.NewLogger(&sb)
	logger.Info("still works")
	if !strings.Contains(sb.String(), "still works") {
		t.Errorf("fallback logger dropped the record: %s", sb.String())
	}
}
