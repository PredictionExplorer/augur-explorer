package imggen

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"sync"
	"testing"
	"time"
)

// artifactServer fakes the image/video server plus the generator service.
type artifactServer struct {
	mu       sync.Mutex
	present  map[string]bool // "000042.png" -> exists
	genCalls []map[string]interface{}
	genCode  int // response code for generation requests (default 200)
	server   *httptest.Server
}

func newArtifactServer(t *testing.T) *artifactServer {
	t.Helper()
	s := &artifactServer{present: make(map[string]bool), genCode: http.StatusOK}
	mux := http.NewServeMux()
	mux.HandleFunc("/generate", func(w http.ResponseWriter, r *http.Request) {
		var payload map[string]interface{}
		if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		s.mu.Lock()
		s.genCalls = append(s.genCalls, payload)
		code := s.genCode
		s.mu.Unlock()
		w.WriteHeader(code)
	})
	serveArtifact := func(w http.ResponseWriter, r *http.Request) {
		name := r.URL.Path[strings.LastIndex(r.URL.Path, "/")+1:]
		s.mu.Lock()
		ok := s.present[name]
		s.mu.Unlock()
		if !ok {
			http.NotFound(w, r)
			return
		}
		w.WriteHeader(http.StatusOK)
	}
	mux.HandleFunc("/images/", serveArtifact)
	mux.HandleFunc("/videos/", serveArtifact)
	s.server = httptest.NewServer(mux)
	t.Cleanup(s.server.Close)
	return s
}

func (s *artifactServer) client() *Client {
	return &Client{
		RequestURL:   s.server.URL + "/generate",
		ImageURL:     s.server.URL + "/images/",
		VideoURL:     s.server.URL + "/videos/",
		HTTPClient:   s.server.Client(),
		PollInterval: time.Millisecond,
		WaitTimeout:  time.Second,
	}
}

func (s *artifactServer) addArtifacts(tokenID int64) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.present[fmt.Sprintf("%06d.png", tokenID)] = true
	s.present[fmt.Sprintf("%06d.mp4", tokenID)] = true
}

func (s *artifactServer) generationCalls() []map[string]interface{} {
	s.mu.Lock()
	defer s.mu.Unlock()
	return append([]map[string]interface{}(nil), s.genCalls...)
}

// staticSource serves a fixed token list.
type staticSource struct {
	tokens []Token
	err    error
}

func (s staticSource) Tokens(context.Context) ([]Token, error) { return s.tokens, s.err }

// safeBuffer is a mutex-guarded writer for cross-goroutine assertions.
type safeBuffer struct {
	mu  sync.Mutex
	buf strings.Builder
}

func (b *safeBuffer) Write(p []byte) (int, error) {
	b.mu.Lock()
	defer b.mu.Unlock()
	return b.buf.Write(p)
}

func (b *safeBuffer) String() string {
	b.mu.Lock()
	defer b.mu.Unlock()
	return b.buf.String()
}

func TestNewClientFromEnv(t *testing.T) {
	t.Parallel()
	full := map[string]string{
		"IM_REQUEST_URL": "http://gen/generate",
		"IM_IMAGE_URL":   "http://img/",
		"IM_VIDEO_URL":   "http://vid/",
	}
	c, err := NewClientFromEnv(func(k string) string { return full[k] }, http.DefaultClient)
	if err != nil {
		t.Fatal(err)
	}
	if c.RequestURL != "http://gen/generate" || c.ImageURL != "http://img/" || c.VideoURL != "http://vid/" {
		t.Fatalf("client = %+v", c)
	}

	for _, missing := range []string{"IM_REQUEST_URL", "IM_IMAGE_URL", "IM_VIDEO_URL"} {
		env := map[string]string{}
		for k, v := range full {
			env[k] = v
		}
		delete(env, missing)
		if _, err := NewClientFromEnv(func(k string) string { return env[k] }, http.DefaultClient); err == nil {
			t.Fatalf("missing %s must error", missing)
		}
	}
}

func TestClientExists(t *testing.T) {
	t.Parallel()
	s := newArtifactServer(t)
	c := s.client()
	ctx := context.Background()

	// Neither artifact present.
	exists, err := c.Exists(ctx, 42)
	if err != nil || exists {
		t.Fatalf("exists = %v, err = %v", exists, err)
	}

	// Image only: still missing.
	s.mu.Lock()
	s.present["000042.png"] = true
	s.mu.Unlock()
	exists, err = c.Exists(ctx, 42)
	if err != nil || exists {
		t.Fatalf("image-only exists = %v, err = %v", exists, err)
	}

	// Both: present.
	s.addArtifacts(42)
	exists, err = c.Exists(ctx, 42)
	if err != nil || !exists {
		t.Fatalf("exists = %v, err = %v", exists, err)
	}

	// Transport failure surfaces as an error.
	bad := &Client{ImageURL: "http://127.0.0.1:1/", VideoURL: "http://127.0.0.1:1/", HTTPClient: http.DefaultClient}
	if _, err := bad.Exists(ctx, 1); err == nil {
		t.Fatal("unreachable server must error")
	}
}

func TestClientGenerate(t *testing.T) {
	t.Parallel()
	s := newArtifactServer(t)
	c := s.client()
	ctx := context.Background()

	if err := c.Generate(ctx, 7, "0xabc"); err != nil {
		t.Fatal(err)
	}
	calls := s.generationCalls()
	if len(calls) != 1 {
		t.Fatalf("calls = %v", calls)
	}
	if calls[0]["token_id"] != float64(7) || calls[0]["seed"] != "0xabc" {
		t.Fatalf("payload = %v", calls[0])
	}

	// Non-2xx answer is an error.
	s.mu.Lock()
	s.genCode = http.StatusBadGateway
	s.mu.Unlock()
	if err := c.Generate(ctx, 8, "0xdef"); err == nil || !strings.Contains(err.Error(), "generator service returned") {
		t.Fatalf("err = %v", err)
	}

	// Transport failure.
	bad := &Client{RequestURL: "http://127.0.0.1:1/generate", HTTPClient: http.DefaultClient}
	if err := bad.Generate(ctx, 9, "s"); err == nil || !strings.Contains(err.Error(), "submitting generation request") {
		t.Fatalf("err = %v", err)
	}

	// Malformed URL fails at request construction.
	malformed := &Client{RequestURL: "http://x/\x7f", HTTPClient: http.DefaultClient}
	if err := malformed.Generate(ctx, 9, "s"); err == nil || !strings.Contains(err.Error(), "building generation request") {
		t.Fatalf("err = %v", err)
	}
}

func TestClientExistsMalformedURL(t *testing.T) {
	t.Parallel()
	c := &Client{ImageURL: "http://x/\x7f", VideoURL: "http://x/", HTTPClient: http.DefaultClient}
	if _, err := c.Exists(context.Background(), 1); err == nil || !strings.Contains(err.Error(), "building HEAD") {
		t.Fatalf("err = %v", err)
	}
}

func TestWaitUntilPresent(t *testing.T) {
	t.Parallel()

	t.Run("appears after polling", func(t *testing.T) {
		t.Parallel()
		s := newArtifactServer(t)
		c := s.client()

		done := make(chan error, 1)
		out := &safeBuffer{}
		go func() { done <- c.WaitUntilPresent(context.Background(), 5, out) }()

		// Create the artifacts only after at least one probe failed, so
		// the polling path (with its progress dot) is deterministic.
		deadline := time.Now().Add(5 * time.Second)
		for !strings.Contains(out.String(), ".") {
			if time.Now().After(deadline) {
				t.Fatal("no failed probe observed")
			}
			time.Sleep(time.Millisecond)
		}
		s.addArtifacts(5)

		select {
		case err := <-done:
			if err != nil {
				t.Fatal(err)
			}
		case <-time.After(5 * time.Second):
			t.Fatal("wait did not finish")
		}
	})

	t.Run("times out", func(t *testing.T) {
		t.Parallel()
		s := newArtifactServer(t)
		c := s.client()
		c.WaitTimeout = 20 * time.Millisecond

		var out strings.Builder
		err := c.WaitUntilPresent(context.Background(), 5, &out)
		if err == nil || !strings.Contains(err.Error(), "still missing after") {
			t.Fatalf("err = %v", err)
		}
	})

	t.Run("probe failure aborts the wait", func(t *testing.T) {
		t.Parallel()
		c := &Client{
			ImageURL:   "http://127.0.0.1:1/",
			VideoURL:   "http://127.0.0.1:1/",
			HTTPClient: http.DefaultClient,
		}
		var out strings.Builder
		err := c.WaitUntilPresent(context.Background(), 5, &out)
		if err == nil || !strings.Contains(err.Error(), "HEAD") {
			t.Fatalf("err = %v", err)
		}
	})

	t.Run("cancelled during poll wait", func(t *testing.T) {
		t.Parallel()
		s := newArtifactServer(t)
		c := s.client()
		// A one-hour poll interval guarantees the cancellation lands in
		// the select, not in the HTTP probe.
		c.PollInterval = time.Hour

		ctx, cancel := context.WithCancel(context.Background())
		done := make(chan error, 1)
		var out strings.Builder
		go func() { done <- c.WaitUntilPresent(ctx, 5, &out) }()
		time.Sleep(10 * time.Millisecond)
		cancel()

		select {
		case err := <-done:
			if !errors.Is(err, context.Canceled) {
				t.Fatalf("err = %v", err)
			}
		case <-time.After(5 * time.Second):
			t.Fatal("wait ignored cancellation")
		}
	})

	t.Run("zero intervals select defaults and return fast when present", func(t *testing.T) {
		t.Parallel()
		s := newArtifactServer(t)
		s.addArtifacts(5)
		c := s.client()
		c.PollInterval = 0
		c.WaitTimeout = 0

		var out strings.Builder
		if err := c.WaitUntilPresent(context.Background(), 5, &out); err != nil {
			t.Fatal(err)
		}
	})
}

func TestScanReportMode(t *testing.T) {
	t.Parallel()
	s := newArtifactServer(t)
	s.addArtifacts(1)
	// Token 2 has no artifacts.

	var out strings.Builder
	err := Scan(context.Background(), ScanOptions{
		Source:      staticSource{tokens: []Token{{ID: 1, Seed: "a"}, {ID: 2, Seed: "b"}}},
		Client:      s.client(),
		Out:         &out,
		TokenPacing: time.Millisecond,
	})
	if err != nil {
		t.Fatal(err)
	}

	report := out.String()
	if !strings.Contains(report, "Checking image/video presence") {
		t.Fatalf("report = %q", report)
	}
	if !strings.Contains(report, "token id = 1    image/video present") {
		t.Fatalf("report = %q", report)
	}
	if !strings.Contains(report, "token id = 2    doesn't exist") {
		t.Fatalf("report = %q", report)
	}
	if calls := s.generationCalls(); len(calls) != 0 {
		t.Fatalf("report mode must not generate, got %v", calls)
	}
}

func TestScanRegenerateMode(t *testing.T) {
	t.Parallel()
	s := newArtifactServer(t)
	s.addArtifacts(1)

	client := s.client()
	source := staticSource{tokens: []Token{{ID: 1, Seed: "a"}, {ID: 2, Seed: "b"}}}

	// Make generation asynchronously create the artifacts after a beat.
	go func() {
		deadline := time.Now().Add(5 * time.Second)
		for time.Now().Before(deadline) {
			if len(s.generationCalls()) > 0 {
				time.Sleep(5 * time.Millisecond)
				s.addArtifacts(2)
				return
			}
			time.Sleep(time.Millisecond)
		}
	}()

	var out strings.Builder
	err := Scan(context.Background(), ScanOptions{
		Source:      source,
		Client:      client,
		Regenerate:  true,
		Out:         &out,
		TokenPacing: time.Millisecond,
	})
	if err != nil {
		t.Fatal(err)
	}

	report := out.String()
	if !strings.Contains(report, "Regenerating missing images/videos") {
		t.Fatalf("report = %q", report)
	}
	if !strings.Contains(report, " regenerating ...") || !strings.Contains(report, " done.") {
		t.Fatalf("report = %q", report)
	}
	calls := s.generationCalls()
	if len(calls) != 1 || calls[0]["token_id"] != float64(2) || calls[0]["seed"] != "b" {
		t.Fatalf("generation calls = %v", calls)
	}
}

func TestScanRegenerateFailures(t *testing.T) {
	t.Parallel()

	t.Run("generation request fails", func(t *testing.T) {
		t.Parallel()
		s := newArtifactServer(t)
		s.mu.Lock()
		s.genCode = http.StatusInternalServerError
		s.mu.Unlock()

		var out strings.Builder
		err := Scan(context.Background(), ScanOptions{
			Source:      staticSource{tokens: []Token{{ID: 3, Seed: "c"}, {ID: 4, Seed: "d"}}},
			Client:      s.client(),
			Regenerate:  true,
			Out:         &out,
			TokenPacing: time.Millisecond,
		})
		if err != nil {
			t.Fatal(err)
		}
		// Both tokens were attempted: a failed request does not abort the scan.
		if got := strings.Count(out.String(), " request failed:"); got != 2 {
			t.Fatalf("report = %q, want 2 request failures", out.String())
		}
	})

	t.Run("wait times out and scan continues", func(t *testing.T) {
		t.Parallel()
		s := newArtifactServer(t)
		s.addArtifacts(9)
		client := s.client()
		client.WaitTimeout = 10 * time.Millisecond

		var out strings.Builder
		err := Scan(context.Background(), ScanOptions{
			Source:      staticSource{tokens: []Token{{ID: 8, Seed: "x"}, {ID: 9, Seed: "y"}}},
			Client:      client,
			Regenerate:  true,
			Out:         &out,
			TokenPacing: time.Millisecond,
		})
		if err != nil {
			t.Fatal(err)
		}
		report := out.String()
		if !strings.Contains(report, " aborting due to error:") {
			t.Fatalf("report = %q", report)
		}
		if !strings.Contains(report, "token id = 9    image/video present") {
			t.Fatalf("report = %q, want the scan to continue to token 9", report)
		}
	})

	t.Run("presence check error reported", func(t *testing.T) {
		t.Parallel()
		client := &Client{
			RequestURL: "http://127.0.0.1:1/generate",
			ImageURL:   "http://127.0.0.1:1/img/",
			VideoURL:   "http://127.0.0.1:1/vid/",
			HTTPClient: http.DefaultClient,
		}
		var out strings.Builder
		err := Scan(context.Background(), ScanOptions{
			Source:      staticSource{tokens: []Token{{ID: 1}}},
			Client:      client,
			Out:         &out,
			TokenPacing: time.Millisecond,
		})
		if err != nil {
			t.Fatal(err)
		}
		if !strings.Contains(out.String(), "error: HEAD") {
			t.Fatalf("report = %q", out.String())
		}
	})
}

func TestScanErrors(t *testing.T) {
	t.Parallel()

	t.Run("nil source", func(t *testing.T) {
		t.Parallel()
		err := Scan(context.Background(), ScanOptions{Client: &Client{}})
		if err == nil || !strings.Contains(err.Error(), "token source is nil") {
			t.Fatalf("err = %v", err)
		}
	})

	t.Run("nil client", func(t *testing.T) {
		t.Parallel()
		err := Scan(context.Background(), ScanOptions{Source: staticSource{}})
		if err == nil || !strings.Contains(err.Error(), "client is nil") {
			t.Fatalf("err = %v", err)
		}
	})

	t.Run("token list failure", func(t *testing.T) {
		t.Parallel()
		var out strings.Builder
		err := Scan(context.Background(), ScanOptions{
			Source: staticSource{err: errors.New("db gone")},
			Client: &Client{},
			Out:    &out,
		})
		if err == nil || !strings.Contains(err.Error(), "failed to list tokens: db gone") {
			t.Fatalf("err = %v", err)
		}
	})

	t.Run("cancellation aborts", func(t *testing.T) {
		t.Parallel()
		s := newArtifactServer(t)
		s.addArtifacts(1)
		ctx, cancel := context.WithCancel(context.Background())
		cancel()

		var out strings.Builder
		err := Scan(ctx, ScanOptions{
			Source:      staticSource{tokens: []Token{{ID: 1}, {ID: 2}}},
			Client:      s.client(),
			Out:         &out,
			TokenPacing: time.Millisecond,
		})
		if !errors.Is(err, context.Canceled) {
			t.Fatalf("err = %v", err)
		}
	})

	t.Run("cancellation during pacing", func(t *testing.T) {
		t.Parallel()
		s := newArtifactServer(t)
		s.addArtifacts(1)
		ctx, cancel := context.WithCancel(context.Background())

		done := make(chan error, 1)
		out := &safeBuffer{}
		go func() {
			done <- Scan(ctx, ScanOptions{
				Source: staticSource{tokens: []Token{{ID: 1}, {ID: 2}}},
				Client: s.client(),
				Out:    out,
				// A one-hour pacing guarantees the cancellation lands in
				// the pacing select after the first token's report.
				TokenPacing: time.Hour,
			})
		}()

		deadline := time.Now().Add(5 * time.Second)
		for !strings.Contains(out.String(), "image/video present") {
			if time.Now().After(deadline) {
				t.Fatal("first token never reported")
			}
			time.Sleep(2 * time.Millisecond)
		}
		cancel()

		select {
		case err := <-done:
			if !errors.Is(err, context.Canceled) {
				t.Fatalf("err = %v", err)
			}
		case <-time.After(5 * time.Second):
			t.Fatal("scan ignored cancellation during pacing")
		}
	})
}
