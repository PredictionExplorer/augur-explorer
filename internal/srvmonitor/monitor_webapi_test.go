package srvmonitor

import (
	"context"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"testing"
)

// splitHostPort extracts host and port from an httptest server URL.
func splitHostPort(t *testing.T, serverURL string) (string, string) {
	t.Helper()
	u, err := url.Parse(serverURL)
	if err != nil {
		t.Fatalf("parsing %q: %v", serverURL, err)
	}
	return u.Hostname(), u.Port()
}

func TestWebAPIMonitorInternalAndPublicAlive(t *testing.T) {
	t.Parallel()
	internal := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
		w.WriteHeader(http.StatusOK)
	}))
	t.Cleanup(internal.Close)
	public := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
		_, _ = w.Write([]byte(`{"status":1}`))
	}))
	t.Cleanup(public.Close)

	host, port := splitHostPort(t, internal.URL)
	apis := []WebAPIConfig{{Title: "API 1", Host: host, Port: port, URI: "/healthz", PublicURL: public.URL}}
	m := NewWebAPIMonitor(apis, 30, testIntervals())

	disp := newFakeDisplay()
	errCh := make(chan string, 10)
	m.check(context.Background(), disp, errCh)

	st := m.statuses[0]
	if !st.Alive || !st.PublicAlive || st.ErrStr != "" || st.PublicErrStr != "" {
		t.Fatalf("status = %+v, want both probes alive", st)
	}
	if msgs := drain(errCh); len(msgs) != 0 {
		t.Fatalf("unexpected errors: %v", msgs)
	}
	row := disp.Row(31)
	if !strings.Contains(row, "Alive") || !strings.Contains(row, "API 1") || !strings.Contains(row, host+":"+port) {
		t.Fatalf("row = %q", row)
	}
	if header := disp.Row(30); !strings.Contains(header, "Web API ( Int / Pub )") {
		t.Fatalf("header = %q", header)
	}
}

func TestWebAPIMonitorInternalLenientPublicStrict(t *testing.T) {
	t.Parallel()
	// Internal answers 404: lenient rule counts 2xx-4xx as alive.
	internal := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
		http.Error(w, "not found", http.StatusNotFound)
	}))
	t.Cleanup(internal.Close)
	// Public answers 404: strict rule fails anything but 200.
	public := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
		http.Error(w, "not found", http.StatusNotFound)
	}))
	t.Cleanup(public.Close)

	host, port := splitHostPort(t, internal.URL)
	apis := []WebAPIConfig{{Title: "API", Host: host, Port: port, URI: "/", PublicURL: public.URL}}
	m := NewWebAPIMonitor(apis, 30, testIntervals())

	errCh := make(chan string, 10)
	m.check(context.Background(), newFakeDisplay(), errCh)

	st := m.statuses[0]
	if !st.Alive {
		t.Fatalf("internal 404 must count as alive (lenient rule), status = %+v", st)
	}
	if st.PublicAlive || st.PublicErrStr != "HTTP 404" {
		t.Fatalf("public status = %+v, want strict HTTP 404 failure", st)
	}
	msgs := drain(errCh)
	if len(msgs) != 1 || !strings.Contains(msgs[0], "(public") {
		t.Fatalf("errors = %v, want one public-probe error", msgs)
	}
}

func TestWebAPIMonitorInternal5xxNotAlive(t *testing.T) {
	t.Parallel()
	internal := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
		http.Error(w, "boom", http.StatusInternalServerError)
	}))
	t.Cleanup(internal.Close)

	host, port := splitHostPort(t, internal.URL)
	m := NewWebAPIMonitor([]WebAPIConfig{{Title: "API", Host: host, Port: port, URI: "/"}}, 30, testIntervals())

	disp := newFakeDisplay()
	errCh := make(chan string, 10)
	m.check(context.Background(), disp, errCh)

	st := m.statuses[0]
	if st.Alive {
		t.Fatal("500 must not count as alive")
	}
	msgs := drain(errCh)
	if len(msgs) != 1 || !strings.Contains(msgs[0], "unexpected HTTP status") {
		t.Fatalf("errors = %v", msgs)
	}
	// Without a public URL the second column renders a dash.
	if row := disp.Row(31); !strings.Contains(row, "-") {
		t.Fatalf("row = %q, want dash placeholder", row)
	}
}

func TestWebAPIMonitorPublicEmptyBodyFails(t *testing.T) {
	t.Parallel()
	internal := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
		w.WriteHeader(http.StatusOK)
	}))
	t.Cleanup(internal.Close)
	public := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
		w.WriteHeader(http.StatusOK) // 200 with empty body
	}))
	t.Cleanup(public.Close)

	host, port := splitHostPort(t, internal.URL)
	m := NewWebAPIMonitor([]WebAPIConfig{{Title: "API", Host: host, Port: port, URI: "/", PublicURL: public.URL}}, 30, testIntervals())

	errCh := make(chan string, 10)
	m.check(context.Background(), newFakeDisplay(), errCh)

	st := m.statuses[0]
	if st.PublicAlive || st.PublicErrStr != "HTTP 200 but empty body" {
		t.Fatalf("public status = %+v", st)
	}
}

func TestWebAPIMonitorUnreachable(t *testing.T) {
	t.Parallel()
	m := NewWebAPIMonitor([]WebAPIConfig{{Title: "API", Host: "127.0.0.1", Port: "1", URI: "/"}}, 30, testIntervals())

	errCh := make(chan string, 10)
	m.check(context.Background(), newFakeDisplay(), errCh)

	st := m.statuses[0]
	if st.Alive || st.ErrStr == "" {
		t.Fatalf("status = %+v, want network error", st)
	}
	msgs := drain(errCh)
	if len(msgs) != 1 || !strings.Contains(msgs[0], "(internal)") {
		t.Fatalf("errors = %v", msgs)
	}
}

func TestWebAPIMonitorMalformedHost(t *testing.T) {
	t.Parallel()
	// A control character makes the probe URL unparsable: the request
	// builder itself must fail cleanly.
	m := NewWebAPIMonitor([]WebAPIConfig{{Title: "API", Host: "bad\x7fhost", Port: "80", URI: "/"}}, 30, testIntervals())

	errCh := make(chan string, 10)
	m.check(context.Background(), newFakeDisplay(), errCh)

	st := m.statuses[0]
	if st.Alive || !strings.Contains(st.ErrStr, "invalid control character") {
		t.Fatalf("status = %+v, want URL parse error", st)
	}
}

func TestWebAPIMonitorStartLoopStopsOnCancel(t *testing.T) {
	t.Parallel()
	cycles := make(chan struct{}, 100)
	internal := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
		select {
		case cycles <- struct{}{}:
		default:
		}
		w.WriteHeader(http.StatusOK)
	}))
	t.Cleanup(internal.Close)

	host, port := splitHostPort(t, internal.URL)
	m := NewWebAPIMonitor([]WebAPIConfig{{Title: "API", Host: host, Port: port, URI: "/"}}, 30, testIntervals())

	ctx, cancel := context.WithCancel(context.Background())
	done := make(chan struct{})
	go func() {
		m.Start(ctx, newFakeDisplay(), make(chan string, 100))
		close(done)
	}()

	<-cycles
	<-cycles
	cancel()
	waitFor(t, "loop exit", func() bool {
		select {
		case <-done:
			return true
		default:
			return false
		}
	})
}
