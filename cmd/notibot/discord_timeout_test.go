package main

import (
	"context"
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

// hungDiscordSink builds a sink whose REST endpoint answers the identity
// probe and then black-holes every call.
func hungDiscordSink(t *testing.T) *discordSink {
	t.Helper()
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if strings.HasSuffix(r.URL.Path, "/users/@me") {
			_, _ = w.Write([]byte(`{"id":"1","username":"bot"}`))
			return
		}
		// Drain the body so the server notices the aborted client and
		// cancels the request context (otherwise Close would wait forever).
		_, _ = io.Copy(io.Discard, r.Body)
		<-r.Context().Done()
	}))
	t.Cleanup(srv.Close)
	target, err := url.Parse(srv.URL)
	if err != nil {
		t.Fatal(err)
	}
	client := disgord.New(disgord.Config{
		BotToken:   "t",
		HTTPClient: &http.Client{Transport: rewriteTransport{target: target}},
	})
	sink := newDiscordSink(client, discordKeys{ChannelID: 101, PriceStatsChanID: 103}, slog.New(slog.DiscardHandler))
	sink.callTimeout = 100 * time.Millisecond
	return sink
}

// TestDiscordSinkBoundsHungAPI proves a black-holed Discord API fails one
// notification within the per-call bound instead of stalling the engine
// loop forever — with no external cancellation at all.
func TestDiscordSinkBoundsHungAPI(t *testing.T) {
	t.Parallel()
	sink := hungDiscordSink(t)

	start := time.Now()
	if err := sink.SendMessage(context.Background(), "text", []byte("png"), "url"); err == nil {
		t.Fatal("SendMessage against a hung API returned nil")
	}
	if elapsed := time.Since(start); elapsed > 5*time.Second {
		t.Fatalf("SendMessage took %v, want ~100ms (the per-call bound)", elapsed)
	}

	start = time.Now()
	if err := sink.SetChannelName(context.Background(), rwbot.ChannelPrice, "name"); err == nil {
		t.Fatal("SetChannelName against a hung API returned nil")
	}
	if elapsed := time.Since(start); elapsed > 5*time.Second {
		t.Fatalf("SetChannelName took %v, want ~100ms (the per-call bound)", elapsed)
	}
}

// TestDiscordSinkDefaultCallTimeout pins the production per-call bound.
func TestDiscordSinkDefaultCallTimeout(t *testing.T) {
	t.Parallel()
	sink := newDiscordSink(nil, discordKeys{}, slog.New(slog.DiscardHandler))
	if sink.callTimeout != discordCallTimeout || discordCallTimeout != 30*time.Second {
		t.Fatalf("callTimeout = %v (const %v), want the documented 30s", sink.callTimeout, discordCallTimeout)
	}
}
