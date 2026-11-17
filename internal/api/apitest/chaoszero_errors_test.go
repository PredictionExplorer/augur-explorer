//go:build integration

package apitest

import (
	"net/http"
	"testing"
)

// TestChaosZeroRoundNotFound pins the missing-round envelope: a round the
// indexer never saw answers the legacy 400 envelope with a stable message,
// not an empty 200 or an internal error.
func TestChaosZeroRoundNotFound(t *testing.T) {
	h := server(t)

	w := h.get(t, "/api/chaoszero/rounds/info/999")
	if w.Code != http.StatusBadRequest {
		t.Fatalf("status = %d, want 400\n%s", w.Code, w.Body.String())
	}
	status, errMsg := decodeEnvelope(t, w.Body.Bytes())
	if status != 0 || errMsg != "Round not found" {
		t.Fatalf("envelope = (%d, %q), want (0, \"Round not found\")", status, errMsg)
	}
}
