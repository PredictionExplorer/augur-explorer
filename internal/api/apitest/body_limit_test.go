//go:build integration

package apitest

// Request-body cap suite: every mutating route on the production router —
// frozen v1, FAQ proxy and generated v2 — rejects oversized bodies with 413
// in its family's error shape, whether the client declares the length or
// streams past the cap. The middleware is global, so this matrix discovers
// future routes automatically through the route table.

import (
	"encoding/json"
	"io"
	"net/http"
	"strings"
	"testing"

	"github.com/PredictionExplorer/augur-explorer/internal/api/common"
	apiv2 "github.com/PredictionExplorer/augur-explorer/internal/api/v2"
)

// undeclaredLengthReader hides its size from httptest.NewRequest so the
// request carries no Content-Length and only the read-time cap can fire.
type undeclaredLengthReader struct{ r io.Reader }

func (u undeclaredLengthReader) Read(p []byte) (int, error) { return u.r.Read(p) }

// oversizedJSON is one valid-JSON-shaped body one byte past the production
// cap, so only the size (never a parse error) can explain a rejection.
func oversizedJSON() string {
	return `{"padding":"` + strings.Repeat("x", int(common.MaxRequestBodyBytes)-12) + `"}`
}

// TestBodyCapOnEveryMutatingRoute walks the real route table: every non-GET
// route must answer a declared oversized body with 413 before any handler,
// auth guard or upstream call runs — RFC 9457 problems under /api/v2/, the
// legacy envelope everywhere else.
func TestBodyCapOnEveryMutatingRoute(t *testing.T) {
	h := server(t)
	body := oversizedJSON()
	if int64(len(body)) <= common.MaxRequestBodyBytes {
		t.Fatalf("test body must exceed the %d-byte cap, got %d", common.MaxRequestBodyBytes, len(body))
	}

	mutating := 0
	for _, route := range h.router.Routes() {
		if route.Method == http.MethodGet {
			continue
		}
		mutating++
		target := strings.NewReplacer("{bidId}", "2001").Replace(route.Pattern)
		t.Run(route.Method+" "+route.Pattern, func(t *testing.T) {
			w := h.do(t, request{method: route.Method, path: target, body: strings.NewReader(body)})
			if w.Code != http.StatusRequestEntityTooLarge {
				t.Fatalf("status = %d, want 413\n%.200s", w.Code, w.Body.String())
			}
			if strings.HasPrefix(route.Pattern, "/api/v2/") {
				if ct := w.Header().Get("Content-Type"); ct != "application/problem+json" {
					t.Fatalf("Content-Type = %q, want application/problem+json", ct)
				}
				var problem apiv2.Problem
				if err := json.Unmarshal(w.Body.Bytes(), &problem); err != nil {
					t.Fatalf("decode problem: %v\n%.200s", err, w.Body.String())
				}
				if problem.Status != http.StatusRequestEntityTooLarge ||
					!strings.HasSuffix(problem.Type, "request-too-large") {
					t.Fatalf("problem = %+v", problem)
				}
				return
			}
			var envelope struct {
				Status *int   `json:"status"`
				Error  string `json:"error"`
			}
			if err := json.Unmarshal(w.Body.Bytes(), &envelope); err != nil {
				t.Fatalf("decode legacy envelope: %v\n%.200s", err, w.Body.String())
			}
			if envelope.Status == nil || *envelope.Status != 0 ||
				!strings.Contains(envelope.Error, "request body exceeds") {
				t.Fatalf("envelope = %+v", envelope)
			}
		})
	}
	// The write surface is small and pinned (ADR-0008/0009 + v1 freeze):
	// 6 v1/FAQ POSTs and 5 generated v2 mutations.
	if mutating != 11 {
		t.Errorf("mutating routes = %d, want 11 — update this matrix alongside the route table", mutating)
	}
}

// TestBodyCapUndeclaredLengthPerFamily streams past the cap without a
// Content-Length so the read-time arm fires inside each consuming layer:
// the generated v2 decoder answers the spec-declared 413 problem, the FAQ
// proxy the legacy 413 envelope, and the frozen v1 handlers keep their 400
// bind-error shape with Go's stable message.
func TestBodyCapUndeclaredLengthPerFamily(t *testing.T) {
	h := server(t)
	spec := v2RankingSpec(t)

	t.Run("v2 decoder answers the declared 413 problem", func(t *testing.T) {
		w := h.do(t, request{
			method: http.MethodPost,
			path:   v2RankingVotes,
			body:   undeclaredLengthReader{strings.NewReader(oversizedJSON())},
		})
		if w.Code != http.StatusRequestEntityTooLarge {
			t.Fatalf("status = %d, want 413\n%.200s", w.Code, w.Body.String())
		}
		validateV2MutationResponse(t, spec, http.MethodPost, v2RankingVotes, v2RankingVotes, "", nil, w)
		var problem apiv2.Problem
		if err := json.Unmarshal(w.Body.Bytes(), &problem); err != nil {
			t.Fatalf("decode problem: %v\n%.300s", err, w.Body.String())
		}
		if problem.Detail == nil || *problem.Detail != common.RequestBodyTooLargeDetail(common.MaxRequestBodyBytes) {
			t.Fatalf("problem = %+v", problem)
		}
	})

	t.Run("faq proxy answers the legacy 413 envelope", func(t *testing.T) {
		w := h.do(t, request{
			method: http.MethodPost,
			path:   "/api/cosmicgame/faq/query",
			body:   undeclaredLengthReader{strings.NewReader(oversizedJSON())},
		})
		if w.Code != http.StatusRequestEntityTooLarge {
			t.Fatalf("status = %d, want 413\n%.200s", w.Code, w.Body.String())
		}
		if !strings.Contains(w.Body.String(), "request body exceeds") {
			t.Fatalf("body = %.200s", w.Body.String())
		}
	})

	t.Run("frozen v1 bind keeps its 400 shape", func(t *testing.T) {
		w := h.do(t, request{
			method:  http.MethodPost,
			path:    "/api/randomwalk/token-ranking/match",
			body:    undeclaredLengthReader{strings.NewReader(oversizedJSON())},
			headers: map[string]string{"X-Ranking-Admin-Key": adminKey},
		})
		if w.Code != http.StatusBadRequest {
			t.Fatalf("status = %d, want the frozen 400 bind-error shape\n%.200s", w.Code, w.Body.String())
		}
		if !strings.Contains(w.Body.String(), "http: request body too large") {
			t.Fatalf("body = %.200s", w.Body.String())
		}
	})

	t.Run("bodies just under the cap still bind", func(t *testing.T) {
		// A syntactically valid vote body padded to one byte under the cap
		// decodes and reaches handler semantics (the garbage signature is
		// rejected, not the size) — the cap admits every legitimate payload.
		payload := `{"nft1":11,"nft2":12,"winner":11,"chainId":1,"nonce":"` +
			strings.Repeat("a", 64) + `","signature":"0x` + strings.Repeat("b", 130) + `","padding":"`
		payload += strings.Repeat("p", int(common.MaxRequestBodyBytes)-len(payload)-3) + `"}`
		if int64(len(payload)) != common.MaxRequestBodyBytes-1 {
			t.Fatalf("padded body = %d bytes, want cap-1 = %d", len(payload), common.MaxRequestBodyBytes-1)
		}
		w := postV2(t, h, v2RankingVotes, payload, nil)
		if w.Code != http.StatusBadRequest {
			t.Fatalf("status = %d, want 400 (signature semantics), body:\n%.300s", w.Code, w.Body.String())
		}
		if !strings.Contains(w.Body.String(), "signature") {
			t.Fatalf("rejection must be handler semantics, not the size: %.300s", w.Body.String())
		}
	})
}
