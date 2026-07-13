package cosmicgame

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/PredictionExplorer/augur-explorer/internal/api/httpx"
	cgmodel "github.com/PredictionExplorer/augur-explorer/internal/model/cosmicgame"
)

// TestMutationGuardsBeforeInitialization pins the POST handlers' database
// guard, which the GET-only route sweep cannot reach (the admin middleware
// answers first through the router).
func TestMutationGuardsBeforeInitialization(t *testing.T) {
	a := NewBare()
	handlers := map[string]httpx.HandlerFunc{
		"ban_bid":   a.handleBanBid,
		"unban_bid": a.handleUnbanBid,
	}
	for name, handler := range handlers {
		t.Run(name, func(t *testing.T) {
			w := httptest.NewRecorder()
			req := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(`{}`))
			handler(httpx.NewContext(w, req))
			if w.Code != http.StatusBadRequest {
				t.Fatalf("status = %d, body=%s", w.Code, w.Body.String())
			}
			if !strings.Contains(w.Body.String(), "Database link wasn't configured") {
				t.Fatalf("body = %s", w.Body.String())
			}
		})
	}
}

// TestEnrichAdminEventsGuards pins the early returns of the contract-backed
// admin-event enrichment: no Ethereum client (bare module) and no events.
func TestEnrichAdminEventsGuards(t *testing.T) {
	bare := NewBare()
	events := []cgmodel.CGAdminEvent{{RecordType: 20, IntegerValue: 4}}
	bare.enrichAdminEventsResolvedValues(events) // nil client: no panic, no change
	if events[0].ResolvedValue != "" {
		t.Fatalf("resolved value changed without a client: %+v", events[0])
	}
	bare.enrichAdminEventsResolvedValues(nil) // empty slice: no-op
}
