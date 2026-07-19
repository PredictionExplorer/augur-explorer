// Package urlalarm implements the cmd/rwalk-alarm engine: it polls a set of
// URLs and notifies a list of people over WhatsApp once a URL has failed
// several consecutive checks.
//
// The engine is injected with its HTTP client and notifier, so the whole
// check/notify pipeline is testable against httptest servers; the binary
// wires the production WhatsApp client (internal/notify/wanotif) and file/
// environment configuration.
package urlalarm

import (
	"context"
	"errors"
	"fmt"
	"log/slog"
	"net/http"
	"strings"
	"time"
)

const (
	// DefaultFailureThreshold is how many consecutive failed checks trigger
	// a notification.
	DefaultFailureThreshold = 5
	// DefaultPollInterval is the pause between check rounds.
	DefaultPollInterval = 2 * time.Second
	// defaultProbeTimeout bounds one URL check. The legacy http.Get had no
	// timeout, so one hung URL stalled the whole loop forever.
	defaultProbeTimeout = 30 * time.Second
)

// Notifier sends a text notification to a phone number. It is satisfied by
// *wanotif.Whatsapp.
type Notifier interface {
	SendText(toPhoneNumber string, text string) (map[string]any, error)
}

// Config holds the engine configuration.
type Config struct {
	// URLs maps each URL to check to its notification message header.
	URLs map[string]string
	// People maps person names to the phone numbers to notify.
	People map[string]string
	// FailureThreshold is how many consecutive failures trigger a
	// notification; zero selects DefaultFailureThreshold.
	FailureThreshold int
	// PollInterval is the pause between check rounds; zero selects
	// DefaultPollInterval.
	PollInterval time.Duration
}

// Engine polls the configured URLs and dispatches notifications.
type Engine struct {
	cfg      Config
	notifier Notifier
	client   *http.Client
	logger   *slog.Logger
	numFails map[string]int
}

// New creates an Engine. httpClient may be nil, selecting a client with the
// default probe timeout.
func New(cfg Config, notifier Notifier, httpClient *http.Client, logger *slog.Logger) *Engine {
	if cfg.FailureThreshold <= 0 {
		cfg.FailureThreshold = DefaultFailureThreshold
	}
	if cfg.PollInterval <= 0 {
		cfg.PollInterval = DefaultPollInterval
	}
	if httpClient == nil {
		httpClient = &http.Client{Timeout: defaultProbeTimeout}
	}
	return &Engine{
		cfg:      cfg,
		notifier: notifier,
		client:   httpClient,
		logger:   logger,
		numFails: make(map[string]int),
	}
}

// Run checks all URLs once per poll interval until ctx is cancelled.
func (e *Engine) Run(ctx context.Context) error {
	timer := time.NewTimer(0)
	defer timer.Stop()
	<-timer.C
	for {
		e.CheckAll(ctx)
		timer.Reset(e.cfg.PollInterval)
		select {
		case <-ctx.Done():
			return ctx.Err()
		case <-timer.C:
		}
	}
}

// CheckAll performs one round over every configured URL.
func (e *Engine) CheckAll(ctx context.Context) {
	for url, header := range e.cfg.URLs {
		e.checkURL(ctx, url, header)
		if ctx.Err() != nil {
			return
		}
	}
}

// checkURL probes one URL and notifies when the consecutive-failure
// threshold is reached.
func (e *Engine) checkURL(ctx context.Context, url, msgHeader string) {
	errStr, ok := e.probe(ctx, url, msgHeader)
	if ok {
		e.numFails[url] = 0
		return
	}
	if ctx.Err() != nil {
		// A cancellation-aborted probe is not a URL failure.
		return
	}

	numFails := e.numFails[url] + 1
	e.numFails[url] = numFails
	if numFails >= e.cfg.FailureThreshold {
		e.logger.Warn("notifying URL failure", "url", url)
		e.notifyFailure(errStr)
		e.numFails[url] = 0
		return
	}
	e.logger.Info("URL check failed",
		"url", url, "err", errStr, "num_fails", numFails, "threshold", e.cfg.FailureThreshold)
}

// probe issues one GET; ok reports whether the URL is healthy, otherwise
// errStr carries the notification text.
func (e *Engine) probe(ctx context.Context, url, msgHeader string) (errStr string, ok bool) {
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return fmt.Sprintf("Networking error: %v", err.Error()), false
	}
	resp, err := e.client.Do(req)
	if err != nil {
		return fmt.Sprintf("Networking error: %v", err.Error()), false
	}
	defer func() { _ = resp.Body.Close() }() // best-effort close on probe path
	if resp.StatusCode == http.StatusOK {
		return "", true
	}
	return fmt.Sprintf("%v. HTTP status: %v", msgHeader, resp.StatusCode), false
}

// notifyFailure sends the message to every configured person.
func (e *Engine) notifyFailure(notifMsg string) {
	for person, phone := range e.cfg.People {
		res, err := e.notifier.SendText(phone, notifMsg)
		if err != nil {
			e.logger.Error(fmt.Sprintf(
				"Error sending whatsapp request to %v: %v (res=%+v,  phone=%v, msg=%v)",
				person, err, res, phone, notifMsg,
			))
			continue
		}
		e.logger.Info("notified failure", "person", person, "phone", phone, "msg", notifMsg)
	}
}

// ParseURLList parses the URL list file format: one "URL<TAB>Message header"
// entry per line; empty lines are skipped.
func ParseURLList(data []byte) (map[string]string, error) {
	urls := make(map[string]string)
	for rowNum, entry := range strings.Split(string(data), "\n") {
		if entry == "" {
			continue
		}
		fields := strings.Split(entry, "\t")
		if len(fields) != 2 {
			return nil, fmt.Errorf("missing tab separator at line %v (%v)", rowNum, entry)
		}
		urls[fields[0]] = fields[1]
	}
	if len(urls) == 0 {
		return nil, errors.New("URL list is empty")
	}
	return urls, nil
}

// ParsePhoneList parses the PHONE_LIST format: comma-separated
// "person:phone" records.
func ParsePhoneList(s string) (map[string]string, error) {
	people := make(map[string]string)
	if s == "" {
		return nil, errors.New("phone list is empty")
	}
	for rowNum, entry := range strings.Split(s, ",") {
		personPhone := strings.Split(entry, ":")
		if len(personPhone) != 2 {
			return nil, fmt.Errorf("phone list entry %v has invalid format (read %v fields, want 2)", rowNum, len(personPhone))
		}
		people[personPhone[0]] = personPhone[1]
	}
	return people, nil
}
