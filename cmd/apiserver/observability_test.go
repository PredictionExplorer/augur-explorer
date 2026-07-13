package main

import (
	"log/slog"
	"net"
	"strings"
	"sync"
	"testing"
	"time"
)

// logSink is a mutex-guarded writer for records emitted by server goroutines.
type logSink struct {
	mu  sync.Mutex
	buf strings.Builder
}

func (s *logSink) Write(p []byte) (int, error) {
	s.mu.Lock()
	defer s.mu.Unlock()
	return s.buf.Write(p)
}

func (s *logSink) String() string {
	s.mu.Lock()
	defer s.mu.Unlock()
	return s.buf.String()
}

func TestStartInternalServerDisabledWhenUnset(t *testing.T) {
	t.Parallel()
	if srv := startInternalServer("  ", slog.New(slog.DiscardHandler)); srv != nil {
		t.Fatal("empty METRICS_ADDR must disable the internal server")
	}
}

func TestStartInternalServerServesMetrics(t *testing.T) {
	t.Parallel()
	sink := &logSink{}
	logger := slog.New(slog.NewTextHandler(sink, nil))
	srv := startInternalServer("127.0.0.1:0", logger)
	if srv == nil {
		t.Fatal("internal server not started")
	}
	t.Cleanup(func() { _ = srv.Close() })

	// ListenAndServe binds asynchronously on the configured address; with
	// port 0 the bound port is not observable, so this test only proves the
	// startup record appears and the server participates in Close.
	deadline := time.Now().Add(5 * time.Second)
	for !strings.Contains(sink.String(), "internal metrics/pprof server listening") {
		if time.Now().After(deadline) {
			t.Fatalf("startup record missing: %q", sink.String())
		}
		time.Sleep(2 * time.Millisecond)
	}
}

func TestStartInternalServerLogsListenFailure(t *testing.T) {
	t.Parallel()
	// Squat the port first so ListenAndServe fails.
	ln, err := net.Listen("tcp", "127.0.0.1:0")
	if err != nil {
		t.Fatal(err)
	}
	t.Cleanup(func() { _ = ln.Close() })

	sink := &logSink{}
	logger := slog.New(slog.NewTextHandler(sink, nil))
	srv := startInternalServer(ln.Addr().String(), logger)
	if srv == nil {
		t.Fatal("internal server not constructed")
	}
	t.Cleanup(func() { _ = srv.Close() })

	deadline := time.Now().Add(5 * time.Second)
	for !strings.Contains(sink.String(), "internal metrics server:") {
		if time.Now().After(deadline) {
			t.Fatalf("listen failure was not logged: %q", sink.String())
		}
		time.Sleep(2 * time.Millisecond)
	}
}

// TestStatusClass pins the metric label mapping.
func TestStatusClass(t *testing.T) {
	t.Parallel()
	for code, want := range map[int]string{
		200: "2xx", 301: "3xx", 404: "4xx", 500: "5xx", 503: "5xx", 101: "2xx",
	} {
		if got := statusClass(code); got != want {
			t.Errorf("statusClass(%d) = %q, want %q", code, got, want)
		}
	}
}
