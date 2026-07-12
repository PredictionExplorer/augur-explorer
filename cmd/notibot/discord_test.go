package main

// The discordSink is exercised against a stub Discord REST API: a
// RoundTripper rewrites every request to an httptest server, so the real
// disgord client marshals real REST calls without touching discord.com.

import (
	"context"
	"encoding/json"
	"io"
	"log/slog"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"testing"
	"time"

	"github.com/andersfylling/disgord"

	"github.com/PredictionExplorer/augur-explorer/internal/notify/rwbot"
)

// rewriteTransport sends every request to the stub server regardless of the
// original host.
type rewriteTransport struct {
	target *url.URL
}

func (t rewriteTransport) RoundTrip(req *http.Request) (*http.Response, error) {
	req.URL.Scheme = t.target.Scheme
	req.URL.Host = t.target.Host
	return http.DefaultTransport.RoundTrip(req)
}

// stubDiscord starts a fake Discord REST endpoint and a sink wired to it.
type stubDiscordCall struct {
	method string
	path   string
	body   string
}

func stubDiscord(t *testing.T, status int, respBody string) (*discordSink, *[]stubDiscordCall) {
	t.Helper()
	var calls []stubDiscordCall
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// disgord.New probes the bot identity at construction and panics on
		// failure; always satisfy it.
		if strings.HasSuffix(r.URL.Path, "/users/@me") {
			_, _ = w.Write([]byte(`{"id":"1","username":"bot"}`))
			return
		}
		body, _ := io.ReadAll(r.Body)
		calls = append(calls, stubDiscordCall{method: r.Method, path: r.URL.Path, body: string(body)})
		w.WriteHeader(status)
		_, _ = w.Write([]byte(respBody))
	}))
	t.Cleanup(srv.Close)

	target, err := url.Parse(srv.URL)
	if err != nil {
		t.Fatalf("parsing stub URL: %v", err)
	}
	// The deprecated HTTPClient field is the one disgord.New's identity
	// probe actually uses; the replacement HttpClient only covers REST.
	client := disgord.New(disgord.Config{
		BotToken:   "test-token",
		HTTPClient: &http.Client{Transport: rewriteTransport{target: target}},
	})
	keys := discordKeys{
		TokenKey:          "test-token",
		ChannelId:         101,
		MintStatsChanId:   102,
		PriceStatsChanId:  103,
		DateStatsChanId:   104,
		RewardStatsChanId: 105,
	}
	sink := newDiscordSink(client, keys, slog.New(slog.DiscardHandler))
	sink.sleep = func(time.Duration) {}
	return sink, &calls
}

func TestDiscordSinkSendMessage(t *testing.T) {
	sink, calls := stubDiscord(t, http.StatusOK, `{"id":"900"}`)

	err := sink.SendMessage(context.Background(), "hello #42", []byte("png"), "https://detail/42")
	if err != nil {
		t.Fatalf("SendMessage: %v", err)
	}
	if len(*calls) != 1 {
		t.Fatalf("REST calls = %d, want 1", len(*calls))
	}
	call := (*calls)[0]
	if call.method != http.MethodPost || !strings.Contains(call.path, "/channels/101/messages") {
		t.Errorf("call = %s %s, want POST to channel 101 messages", call.method, call.path)
	}
	// Multipart upload: the text and the embed URL travel in the payload,
	// the image as a file part.
	for _, want := range []string{"hello #42", "https://detail/42", "token.png", "png"} {
		if !strings.Contains(call.body, want) {
			t.Errorf("message payload missing %q", want)
		}
	}
}

func TestDiscordSinkSendMessageError(t *testing.T) {
	sink, _ := stubDiscord(t, http.StatusForbidden, `{"message":"Missing Permissions","code":50013}`)
	if err := sink.SendMessage(context.Background(), "x", nil, "u"); err == nil {
		t.Error("SendMessage against a 403 API returned nil")
	}
}

func TestDiscordSinkSetChannelName(t *testing.T) {
	sink, calls := stubDiscord(t, http.StatusOK, `{"id":"103","name":"Cur. price"}`)

	err := sink.SetChannelName(context.Background(), rwbot.ChannelPrice, "Cur. price 💲 : 0.5000")
	if err != nil {
		t.Fatalf("SetChannelName: %v", err)
	}
	if len(*calls) != 1 {
		t.Fatalf("REST calls = %d, want 1", len(*calls))
	}
	call := (*calls)[0]
	if !strings.Contains(call.path, "/channels/103") {
		t.Errorf("call path = %s, want the price channel id", call.path)
	}
	var payload struct {
		Name string `json:"name"`
	}
	if err := json.Unmarshal([]byte(call.body), &payload); err != nil {
		t.Fatalf("parsing rename payload %q: %v", call.body, err)
	}
	if payload.Name != "Cur. price 💲 : 0.5000" {
		t.Errorf("rename payload name = %q", payload.Name)
	}
}

func TestDiscordSinkSetChannelNameSkipsUnconfigured(t *testing.T) {
	sink, calls := stubDiscord(t, http.StatusOK, `{}`)
	sink.channels[rwbot.ChannelLastDate] = 0

	if err := sink.SetChannelName(context.Background(), rwbot.ChannelLastDate, "x"); err != nil {
		t.Fatalf("SetChannelName: %v", err)
	}
	if len(*calls) != 0 {
		t.Errorf("REST calls = %d, want none for an unconfigured channel", len(*calls))
	}
}

func TestDiscordSinkRateLimitedRenameRetries(t *testing.T) {
	var calls int
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if strings.HasSuffix(r.URL.Path, "/users/@me") {
			_, _ = w.Write([]byte(`{"id":"1","username":"bot"}`))
			return
		}
		calls++
		if calls == 1 {
			w.WriteHeader(http.StatusTooManyRequests)
			_, _ = w.Write([]byte(`{"message": "You are being rate limited.", "retry_after": 1.2, "global": false}`))
			return
		}
		_, _ = w.Write([]byte(`{"id":"103"}`))
	}))
	defer srv.Close()
	target, err := url.Parse(srv.URL)
	if err != nil {
		t.Fatal(err)
	}
	client := disgord.New(disgord.Config{
		BotToken:   "t",
		HTTPClient: &http.Client{Transport: rewriteTransport{target: target}},
	})
	sink := newDiscordSink(client, discordKeys{PriceStatsChanId: 103}, slog.New(slog.DiscardHandler))
	var slept time.Duration
	sink.sleep = func(d time.Duration) { slept = d }

	if err := sink.SetChannelName(context.Background(), rwbot.ChannelPrice, "n"); err != nil {
		t.Fatalf("SetChannelName after rate limit: %v (calls %d)", err, calls)
	}
	if calls < 2 {
		t.Errorf("REST calls = %d, want a retry after the 429", calls)
	}
	if slept < time.Second {
		t.Errorf("slept %v, want at least the reported retry_after + safety", slept)
	}
}
