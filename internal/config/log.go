package config

import (
	"fmt"
	"io"
	"log/slog"
	"strings"
)

// Log selects the process logger's output format and minimum level, shared
// by every service configuration. Services log to stdout and journald owns
// persistence; production env files set LOG_FORMAT=json, developers keep
// the text default.
type Log struct {
	// Format is "text" (dev default) or "json" (production).
	Format string `env:"LOG_FORMAT" default:"text"`
	// Level is the minimum record level: debug, info, warn or error.
	Level string `env:"LOG_LEVEL" default:"info"`
}

// validate returns the Log-specific problems (enum checks) in Load's
// problem-line format.
func (l Log) validate() []string {
	var problems []string
	if _, err := l.slogLevel(); err != nil {
		problems = append(problems, fmt.Sprintf("LOG_LEVEL: %v", err))
	}
	switch strings.ToLower(l.Format) {
	case "text", "json":
	default:
		problems = append(problems, fmt.Sprintf("LOG_FORMAT: unknown format %q (use text or json)", l.Format))
	}
	return problems
}

// slogLevel maps the configured level name onto a slog.Level.
func (l Log) slogLevel() (slog.Level, error) {
	switch strings.ToLower(l.Level) {
	case "debug":
		return slog.LevelDebug, nil
	case "info", "":
		return slog.LevelInfo, nil
	case "warn", "warning":
		return slog.LevelWarn, nil
	case "error":
		return slog.LevelError, nil
	}
	return 0, fmt.Errorf("unknown level %q (use debug, info, warn or error)", l.Level)
}

// NewLogger builds the process logger writing to w: a JSON handler when
// Format is "json", a text handler otherwise, filtering below the
// configured level. Call it once per process and inject the logger (or a
// .With derivative) everywhere.
func (l Log) NewLogger(w io.Writer) *slog.Logger {
	level, err := l.slogLevel()
	if err != nil {
		level = slog.LevelInfo // Load validated already; be forgiving here
	}
	opts := &slog.HandlerOptions{Level: level}
	if strings.EqualFold(l.Format, "json") {
		return slog.New(slog.NewJSONHandler(w, opts))
	}
	return slog.New(slog.NewTextHandler(w, opts))
}
