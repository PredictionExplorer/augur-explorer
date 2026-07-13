//go:build integration

package apitest

import (
	"io"
	"log"
	"log/slog"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"

	"github.com/PredictionExplorer/augur-explorer/internal/api/cosmicgame"
	"github.com/PredictionExplorer/augur-explorer/internal/api/faq"
	"github.com/PredictionExplorer/augur-explorer/internal/api/randomwalk"
	"github.com/PredictionExplorer/augur-explorer/internal/api/routes"
	v2 "github.com/PredictionExplorer/augur-explorer/internal/api/v2"
)

// TestTwoIndependentRoutersInOneProcess pins the Phase 4 DI guarantee: the
// v1 modules hold no package-level state, so a second fully wired router can
// be constructed next to the shared harness router and both answer requests
// independently. Before the idiomatic-sweep sprint this was impossible —
// module Init could run only once per process (the old §4.1 caveat).
func TestTwoIndependentRoutersInOneProcess(t *testing.T) {
	h := server(t)
	discard := log.New(io.Discard, "", 0)

	// Second, independently constructed module set over the same store and
	// fake chain. A separate FAQ upstream proves per-instance configuration.
	secondFAQ := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		_, _ = w.Write([]byte(`{"service":"second-faq"}`))
	}))
	t.Cleanup(secondFAQ.Close)

	cgAPI, err := cosmicgame.New(t.Context(), cosmicgame.Config{
		Store:     h.store,
		EthClient: h.ethClient,
		RPCClient: h.rpcClient,
		Info:      discard,
		Error:     discard,
	})
	if err != nil {
		t.Fatalf("second cosmicgame module: %v", err)
	}
	v2Server, err := v2.NewServer(
		h.store,
		cgAPI.ContractState(),
		slog.New(slog.NewTextHandler(io.Discard, nil)),
		v2.WithClock(func() time.Time { return time.Unix(1767230000, 0) }),
	)
	if err != nil {
		t.Fatalf("second v2 server: %v", err)
	}
	second := routes.New(h.store, routes.Options{
		CosmicGame: cgAPI,
		RandomWalk: randomwalk.New(h.store, discard, discard),
		FAQ:        faq.New(faq.Options{UpstreamURL: secondFAQ.URL}),
		V2:         v2Server,
	})

	get := func(t *testing.T, path string) *httptest.ResponseRecorder {
		t.Helper()
		req := httptest.NewRequest(http.MethodGet, path, nil)
		req.RemoteAddr = "10.77.77.77:4242"
		w := httptest.NewRecorder()
		second.ServeHTTP(w, req)
		return w
	}

	t.Run("both routers serve the same v1 surface", func(t *testing.T) {
		paths := []string{
			"/api/cosmicgame/statistics/counters",
			"/api/randomwalk/floor_price",
			"/api/v2/cosmicgame/rounds/1",
		}
		for _, path := range paths {
			first := h.get(t, path)
			w := get(t, path)
			if w.Code != first.Code {
				t.Fatalf("%s: second router status %d, harness router %d\n%s",
					path, w.Code, first.Code, w.Body.String())
			}
			if w.Body.String() != first.Body.String() {
				t.Fatalf("%s: second router body diverged\nsecond: %s\nfirst:  %s",
					path, w.Body.String(), first.Body.String())
			}
		}
	})

	t.Run("per-instance config stays isolated", func(t *testing.T) {
		w := get(t, "/api/cosmicgame/faq/health")
		if w.Code != http.StatusOK || !strings.Contains(w.Body.String(), "second-faq") {
			t.Fatalf("second router FAQ = %d\n%s", w.Code, w.Body.String())
		}
		first := h.get(t, "/api/cosmicgame/faq/health")
		if first.Code != http.StatusOK || !strings.Contains(first.Body.String(), "faq-bot-stub") {
			t.Fatalf("harness router FAQ changed = %d\n%s", first.Code, first.Body.String())
		}
	})

	t.Run("route tables are identical", func(t *testing.T) {
		firstRoutes := make(map[string]bool)
		for _, route := range h.router.Routes() {
			firstRoutes[route.Method+" "+route.Pattern] = true
		}
		for _, route := range second.Routes() {
			if !firstRoutes[route.Method+" "+route.Pattern] {
				t.Fatalf("second router registered %s %s, absent from the harness router",
					route.Method, route.Pattern)
			}
		}
		if got, want := len(second.Routes()), len(h.router.Routes()); got != want {
			t.Fatalf("second router has %d routes, harness router %d", got, want)
		}
	})
}
