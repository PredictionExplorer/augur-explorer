package main

import (
	"context"
	"net/http"
	"net/http/httptest"
	"os"
	"path/filepath"
	"strings"
	"testing"
	"time"
)

func writeURLList(t *testing.T, content string) string {
	t.Helper()
	path := filepath.Join(t.TempDir(), "urls.txt")
	if err := os.WriteFile(path, []byte(content), 0o600); err != nil {
		t.Fatal(err)
	}
	return path
}

func validEnv() map[string]string {
	return map[string]string{
		"PHONE_LIST":        "alice:+1555",
		"WHATSAPP_TOKEN":    "token",
		"WHATSAPP_PHONE_ID": "42",
	}
}

func envFunc(m map[string]string) func(string) string {
	return func(k string) string { return m[k] }
}

func TestRunConfigErrors(t *testing.T) {
	t.Parallel()
	urlFile := writeURLList(t, "http://a.example/x\tA down\n")

	cases := map[string]struct {
		args    []string
		env     map[string]string
		wantErr string
	}{
		"no args":  {args: nil, env: validEnv(), wantErr: "usage:"},
		"two args": {args: []string{"a", "b"}, env: validEnv(), wantErr: "usage:"},
		"missing file": {
			args: []string{filepath.Join(t.TempDir(), "missing.txt")},
			env:  validEnv(), wantErr: "reading url list",
		},
		"bad url list": {
			args: []string{writeURLList(t, "no tab here\n")},
			env:  validEnv(), wantErr: "parsing url list",
		},
		"bad phone list": {
			args: []string{urlFile},
			env: map[string]string{
				"PHONE_LIST": "justname", "WHATSAPP_TOKEN": "t", "WHATSAPP_PHONE_ID": "1",
			},
			wantErr: "parsing PHONE_LIST",
		},
		"empty phone list": {
			args: []string{urlFile},
			env: map[string]string{
				"WHATSAPP_TOKEN": "t", "WHATSAPP_PHONE_ID": "1",
			},
			wantErr: "PHONE_LIST: required but not set",
		},
		"missing token": {
			args: []string{urlFile},
			env: map[string]string{
				"PHONE_LIST": "alice:+1555", "WHATSAPP_PHONE_ID": "1",
			},
			wantErr: "WHATSAPP_TOKEN: required but not set",
		},
		"missing phone id": {
			args: []string{urlFile},
			env: map[string]string{
				"PHONE_LIST": "alice:+1555", "WHATSAPP_TOKEN": "t",
			},
			wantErr: "WHATSAPP_PHONE_ID: required but not set",
		},
	}

	for name, tc := range cases {
		t.Run(name, func(t *testing.T) {
			t.Parallel()
			var out strings.Builder
			err := run(context.Background(), tc.args, envFunc(tc.env), &out)
			if err == nil || !strings.Contains(err.Error(), tc.wantErr) {
				t.Fatalf("err = %v, want containing %q", err, tc.wantErr)
			}
		})
	}
}

func TestRunLoopsUntilCancelled(t *testing.T) {
	t.Parallel()
	checks := make(chan struct{}, 100)
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
		select {
		case checks <- struct{}{}:
		default:
		}
		w.WriteHeader(http.StatusOK)
	}))
	t.Cleanup(server.Close)

	urlFile := writeURLList(t, server.URL+"\tAPI down\n")

	ctx, cancel := context.WithCancel(context.Background())
	var out strings.Builder
	done := make(chan error, 1)
	go func() {
		done <- run(ctx, []string{urlFile}, envFunc(validEnv()), &out)
	}()

	// The URL is probed at least once before we stop.
	select {
	case <-checks:
	case <-time.After(5 * time.Second):
		t.Fatal("no URL check happened")
	}
	cancel()

	select {
	case err := <-done:
		if err != nil {
			t.Fatalf("run returned %v, want nil on graceful shutdown", err)
		}
	case <-time.After(5 * time.Second):
		t.Fatal("run did not stop on cancellation")
	}

	log := out.String()
	if !strings.Contains(log, "msg=\"loaded URLs\" count=1") ||
		!strings.Contains(log, "msg=\"loaded phones for notification\" count=1") ||
		!strings.Contains(log, "exiting upon user request") {
		t.Fatalf("log = %q", log)
	}
	// The startup record redacts the WhatsApp token.
	if !strings.Contains(log, "WHATSAPP_TOKEN=[set]") || strings.Contains(log, "token") {
		t.Fatalf("effective configuration leaked or missing: %q", log)
	}
}

func TestRunReturnsNonCancellationErrors(t *testing.T) {
	t.Parallel()
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
		w.WriteHeader(http.StatusOK)
	}))
	t.Cleanup(server.Close)
	urlFile := writeURLList(t, server.URL+"\tAPI down\n")

	// A deadline (not a cancellation) must surface as an error: only a
	// user-requested shutdown exits cleanly.
	ctx, cancel := context.WithTimeout(context.Background(), 50*time.Millisecond)
	defer cancel()

	var out strings.Builder
	err := run(ctx, []string{urlFile}, envFunc(validEnv()), &out)
	if err == nil || !strings.Contains(err.Error(), "deadline") {
		t.Fatalf("err = %v, want deadline error", err)
	}
}
