package urlalarm

import (
	"context"
	"errors"
	"log"
	"net/http"
	"net/http/httptest"
	"reflect"
	"strings"
	"sync"
	"testing"
	"time"
)

// fakeNotifier records notifications and can fail per phone number.
type fakeNotifier struct {
	mu   sync.Mutex
	sent [][2]string // {phone, text}
	fail map[string]error
}

func (f *fakeNotifier) SendText(phone, text string) (map[string]interface{}, error) {
	f.mu.Lock()
	defer f.mu.Unlock()
	if err, ok := f.fail[phone]; ok {
		return map[string]interface{}{"error": "scripted"}, err
	}
	f.sent = append(f.sent, [2]string{phone, text})
	return map[string]interface{}{"messages": []any{}}, nil
}

func (f *fakeNotifier) sentCount() int {
	f.mu.Lock()
	defer f.mu.Unlock()
	return len(f.sent)
}

// flakyServer answers with the scripted status codes in sequence, repeating
// the last one.
type flakyServer struct {
	mu     sync.Mutex
	codes  []int
	served int
}

func (s *flakyServer) handler(w http.ResponseWriter, _ *http.Request) {
	s.mu.Lock()
	idx := s.served
	if idx >= len(s.codes) {
		idx = len(s.codes) - 1
	}
	code := s.codes[idx]
	s.served++
	s.mu.Unlock()
	w.WriteHeader(code)
	_, _ = w.Write([]byte("body"))
}

func newTestEngine(urls map[string]string, notifier Notifier) (*Engine, *strings.Builder) {
	var buf strings.Builder
	logger := log.New(&buf, "", 0)
	cfg := Config{
		URLs:             urls,
		People:           map[string]string{"alice": "+1555", "bob": "+1666"},
		FailureThreshold: 3,
		PollInterval:     time.Millisecond,
	}
	return New(cfg, notifier, nil, logger), &buf
}

func TestEngineHealthyURLNeverNotifies(t *testing.T) {
	t.Parallel()
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
		w.WriteHeader(http.StatusOK)
	}))
	t.Cleanup(server.Close)

	notifier := &fakeNotifier{}
	e, _ := newTestEngine(map[string]string{server.URL: "API down"}, notifier)

	for i := 0; i < 10; i++ {
		e.CheckAll(context.Background())
	}
	if notifier.sentCount() != 0 {
		t.Fatalf("sent = %d, want 0", notifier.sentCount())
	}
	if e.numFails[server.URL] != 0 {
		t.Fatalf("numFails = %d", e.numFails[server.URL])
	}
}

func TestEngineNotifiesAfterThresholdAndResets(t *testing.T) {
	t.Parallel()
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
		http.Error(w, "boom", http.StatusInternalServerError)
	}))
	t.Cleanup(server.Close)

	notifier := &fakeNotifier{}
	e, logBuf := newTestEngine(map[string]string{server.URL: "API down"}, notifier)

	// Two failures: below the threshold of 3, nothing sent.
	e.CheckAll(context.Background())
	e.CheckAll(context.Background())
	if notifier.sentCount() != 0 {
		t.Fatalf("sent early: %v", notifier.sent)
	}

	// Third failure: both people notified, counter reset.
	e.CheckAll(context.Background())
	if notifier.sentCount() != 2 {
		t.Fatalf("sent = %d, want 2", notifier.sentCount())
	}
	if e.numFails[server.URL] != 0 {
		t.Fatalf("numFails = %d, want reset", e.numFails[server.URL])
	}

	notifier.mu.Lock()
	phones := map[string]bool{}
	var text string
	for _, s := range notifier.sent {
		phones[s[0]] = true
		text = s[1]
	}
	notifier.mu.Unlock()
	if !phones["+1555"] || !phones["+1666"] {
		t.Fatalf("phones = %v", phones)
	}
	if want := "API down. HTTP status: 500"; text != want {
		t.Fatalf("text = %q, want %q", text, want)
	}
	if !strings.Contains(logBuf.String(), "Notifying failure of url") {
		t.Fatalf("log = %q", logBuf.String())
	}
}

func TestEngineRecoveryResetsCounter(t *testing.T) {
	t.Parallel()
	flaky := &flakyServer{codes: []int{500, 500, 200, 500, 500, 500}}
	server := httptest.NewServer(http.HandlerFunc(flaky.handler))
	t.Cleanup(server.Close)

	notifier := &fakeNotifier{}
	e, _ := newTestEngine(map[string]string{server.URL: "API"}, notifier)

	// 500, 500 (2 fails), 200 (reset), 500, 500 (2 fails): never reaches 3.
	for i := 0; i < 5; i++ {
		e.CheckAll(context.Background())
	}
	if notifier.sentCount() != 0 {
		t.Fatalf("sent = %d, want 0", notifier.sentCount())
	}
	// The sixth check is the third consecutive failure.
	e.CheckAll(context.Background())
	if notifier.sentCount() != 2 {
		t.Fatalf("sent = %d, want 2", notifier.sentCount())
	}
}

func TestEngineMalformedURL(t *testing.T) {
	t.Parallel()
	notifier := &fakeNotifier{}
	// A control character makes request construction itself fail; the
	// failure counts like any network error.
	e, _ := newTestEngine(map[string]string{"http://x/\x7f": "API"}, notifier)

	for i := 0; i < 3; i++ {
		e.CheckAll(context.Background())
	}
	if notifier.sentCount() != 2 {
		t.Fatalf("sent = %d, want 2", notifier.sentCount())
	}
	notifier.mu.Lock()
	text := notifier.sent[0][1]
	notifier.mu.Unlock()
	if !strings.HasPrefix(text, "Networking error: ") {
		t.Fatalf("text = %q", text)
	}
}

func TestEngineNetworkErrorText(t *testing.T) {
	t.Parallel()
	notifier := &fakeNotifier{}
	e, _ := newTestEngine(map[string]string{"http://127.0.0.1:1/x": "API"}, notifier)

	for i := 0; i < 3; i++ {
		e.CheckAll(context.Background())
	}
	if notifier.sentCount() != 2 {
		t.Fatalf("sent = %d, want 2", notifier.sentCount())
	}
	notifier.mu.Lock()
	text := notifier.sent[0][1]
	notifier.mu.Unlock()
	if !strings.HasPrefix(text, "Networking error: ") {
		t.Fatalf("text = %q", text)
	}
}

func TestEngineNotifierFailureLoggedAndOthersStillNotified(t *testing.T) {
	t.Parallel()
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
		http.Error(w, "down", http.StatusBadGateway)
	}))
	t.Cleanup(server.Close)

	notifier := &fakeNotifier{fail: map[string]error{"+1555": errors.New("graph api rejected")}}
	e, logBuf := newTestEngine(map[string]string{server.URL: "API"}, notifier)

	for i := 0; i < 3; i++ {
		e.CheckAll(context.Background())
	}
	// Only bob's message went through; alice's failure is logged.
	if notifier.sentCount() != 1 {
		t.Fatalf("sent = %v", notifier.sent)
	}
	if !strings.Contains(logBuf.String(), "Error sending whatsapp request to alice") {
		t.Fatalf("log = %q", logBuf.String())
	}
}

func TestEngineRunLoopStopsOnCancel(t *testing.T) {
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

	e, _ := newTestEngine(map[string]string{server.URL: "API"}, &fakeNotifier{})

	ctx, cancel := context.WithCancel(context.Background())
	done := make(chan error, 1)
	go func() { done <- e.Run(ctx) }()

	<-checks
	<-checks // at least two rounds
	cancel()
	select {
	case err := <-done:
		if !errors.Is(err, context.Canceled) {
			t.Fatalf("Run returned %v, want context.Canceled", err)
		}
	case <-time.After(5 * time.Second):
		t.Fatal("Run did not stop on cancellation")
	}
}

func TestEngineCancellationIsNotAFailure(t *testing.T) {
	t.Parallel()
	// A cancelled context aborts the probe; the URL's failure counter must
	// not advance (a shutdown is not an outage).
	e, _ := newTestEngine(map[string]string{"http://127.0.0.1:1/x": "API"}, &fakeNotifier{})
	ctx, cancel := context.WithCancel(context.Background())
	cancel()

	e.CheckAll(ctx)
	if got := e.numFails["http://127.0.0.1:1/x"]; got != 0 {
		t.Fatalf("numFails = %d, want 0 after cancelled probe", got)
	}
}

func TestEngineDefaults(t *testing.T) {
	t.Parallel()
	var buf strings.Builder
	e := New(Config{}, &fakeNotifier{}, nil, log.New(&buf, "", 0))
	if e.cfg.FailureThreshold != DefaultFailureThreshold {
		t.Fatalf("threshold = %d", e.cfg.FailureThreshold)
	}
	if e.cfg.PollInterval != DefaultPollInterval {
		t.Fatalf("interval = %v", e.cfg.PollInterval)
	}
	if e.client == nil || e.client.Timeout == 0 {
		t.Fatal("default client must carry a timeout")
	}
}

func TestParseURLList(t *testing.T) {
	t.Parallel()

	t.Run("valid", func(t *testing.T) {
		t.Parallel()
		data := "http://a.example/x\tA down\n\nhttp://b.example/y\tB down\n"
		urls, err := ParseURLList([]byte(data))
		if err != nil {
			t.Fatal(err)
		}
		want := map[string]string{
			"http://a.example/x": "A down",
			"http://b.example/y": "B down",
		}
		if !reflect.DeepEqual(urls, want) {
			t.Fatalf("urls = %v", urls)
		}
	})

	t.Run("missing tab", func(t *testing.T) {
		t.Parallel()
		_, err := ParseURLList([]byte("http://a.example/x A down\n"))
		if err == nil || !strings.Contains(err.Error(), "missing tab separator at line 0") {
			t.Fatalf("err = %v", err)
		}
	})

	t.Run("empty", func(t *testing.T) {
		t.Parallel()
		for _, data := range []string{"", "\n\n"} {
			if _, err := ParseURLList([]byte(data)); err == nil {
				t.Fatalf("data %q must error", data)
			}
		}
	})
}

func TestParsePhoneList(t *testing.T) {
	t.Parallel()

	t.Run("valid", func(t *testing.T) {
		t.Parallel()
		people, err := ParsePhoneList("alice:+1555,bob:+1666")
		if err != nil {
			t.Fatal(err)
		}
		want := map[string]string{"alice": "+1555", "bob": "+1666"}
		if !reflect.DeepEqual(people, want) {
			t.Fatalf("people = %v", people)
		}
	})

	t.Run("invalid entry", func(t *testing.T) {
		t.Parallel()
		_, err := ParsePhoneList("alice:+1555,bob")
		if err == nil || !strings.Contains(err.Error(), "entry 1 has invalid format") {
			t.Fatalf("err = %v", err)
		}
	})

	t.Run("empty", func(t *testing.T) {
		t.Parallel()
		if _, err := ParsePhoneList(""); err == nil {
			t.Fatal("empty list must error")
		}
	})
}
