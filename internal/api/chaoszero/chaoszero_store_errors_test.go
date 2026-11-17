package chaoszero

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/jackc/pgx/v5/pgxpool"

	"github.com/PredictionExplorer/augur-explorer/internal/api/httpx"
	"github.com/PredictionExplorer/augur-explorer/internal/store"
)

// failingStoreAPI builds the module over a store whose pool dials a port
// nothing listens on: construction is lazy, so every query — and only the
// queries — fails. This is the unit-level stand-in for a database outage.
func failingStoreAPI(t *testing.T) *API {
	t.Helper()
	pool, err := pgxpool.New(context.Background(),
		"postgres://cz:cz@127.0.0.1:1/cz?sslmode=disable&connect_timeout=1")
	if err != nil {
		t.Fatalf("building lazy pool: %v", err)
	}
	t.Cleanup(pool.Close)
	return New(store.NewFromPool(pool), Options{})
}

func routedRequest(t *testing.T, a *API, path string) *httptest.ResponseRecorder {
	t.Helper()
	r := httpx.NewRouter()
	a.RegisterRoutes(r)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, httptest.NewRequest(http.MethodGet, path, nil))
	return w
}

// TestStoreFailuresAnswerInternalError drives every handler shape through a
// database outage: the response must be the opaque HTTP-500 envelope
// {"status":0,"error":"Internal server error"} — never a stack trace or a
// driver message.
func TestStoreFailuresAnswerInternalError(t *testing.T) {
	a := failingStoreAPI(t)

	paths := []string{
		"/api/chaoszero/statistics",
		"/api/chaoszero/contracts",
		"/api/chaoszero/rounds/list/0/10",
		"/api/chaoszero/rounds/info/0",
		"/api/chaoszero/trades/list/0/10",
		"/api/chaoszero/trades/by_round/0/0/10",
		"/api/chaoszero/traders/list/0/10",
		"/api/chaoszero/funding/list/0/10",
		"/api/chaoszero/splits/list/0/10",
		"/api/chaoszero/merges/list/0/10",
		"/api/chaoszero/settlements/list/0/10",
		"/api/chaoszero/claims/list/0/10",
		"/api/chaoszero/answers/list/0/10",
		"/api/chaoszero/escrows/list/0/10",
		"/api/chaoszero/payouts/list/0/10",
		"/api/chaoszero/transfers/list/0/10",
		"/api/chaoszero/approvals/list/0/10",
	}
	for _, path := range paths {
		t.Run(path, func(t *testing.T) {
			w := routedRequest(t, a, path)
			if w.Code != http.StatusInternalServerError {
				t.Fatalf("status = %d, want 500\n%s", w.Code, w.Body.String())
			}
			var envelope struct {
				Status int    `json:"status"`
				Error  string `json:"error"`
			}
			if err := json.Unmarshal(w.Body.Bytes(), &envelope); err != nil {
				t.Fatalf("body is not the envelope: %v\n%s", err, w.Body.String())
			}
			if envelope.Status != 0 || envelope.Error != "Internal server error" {
				t.Fatalf("envelope = %+v, want the opaque internal error", envelope)
			}
		})
	}
}

// TestByUserLookupFailureAnswersEmptyList pins the by-user convention: an
// address the store cannot resolve (here: outage; in production: an address
// never seen on chain) answers an empty 200 list with UserAid 0 instead of
// an error.
func TestByUserLookupFailureAnswersEmptyList(t *testing.T) {
	a := failingStoreAPI(t)

	for _, path := range []string{
		"/api/chaoszero/trades/by_user/0x2100000000000000000000000000000000000021/0/10",
		"/api/chaoszero/claims/by_user/0x2100000000000000000000000000000000000021/0/10",
		"/api/chaoszero/transfers/by_user/0x2100000000000000000000000000000000000021/0/10",
	} {
		t.Run(path, func(t *testing.T) {
			w := routedRequest(t, a, path)
			if w.Code != http.StatusOK {
				t.Fatalf("status = %d, want 200\n%s", w.Code, w.Body.String())
			}
			var body struct {
				Status  int   `json:"status"`
				UserAid int64 `json:"UserAid"`
			}
			if err := json.Unmarshal(w.Body.Bytes(), &body); err != nil {
				t.Fatal(err)
			}
			if body.Status != 1 || body.UserAid != 0 {
				t.Fatalf("body = %+v, want status 1 with UserAid 0", body)
			}
		})
	}
}

// TestEmptyPathParameterGuards pins the parameter guards that run before any
// store access; an unrouted request carries no path values, which is the
// only way the empty-parameter branches can trigger.
func TestEmptyPathParameterGuards(t *testing.T) {
	a := failingStoreAPI(t)

	cases := []struct {
		name    string
		handler func(*httpx.Context)
		wantErr string
	}{
		{"round_num", a.handleRoundInfo, "'round_num' parameter is not set"},
		{"user_addr", a.handleTradesByUser, "'user_addr' parameter is not set"},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			w := httptest.NewRecorder()
			c := httpx.NewContext(w, httptest.NewRequest(http.MethodGet, "/unrouted", nil))
			tc.handler(c)
			var envelope struct {
				Status int    `json:"status"`
				Error  string `json:"error"`
			}
			if err := json.Unmarshal(w.Body.Bytes(), &envelope); err != nil {
				t.Fatalf("body is not the envelope: %v\n%s", err, w.Body.String())
			}
			if envelope.Status != 0 || envelope.Error != tc.wantErr {
				t.Fatalf("envelope = %+v, want error %q", envelope, tc.wantErr)
			}
		})
	}
}
