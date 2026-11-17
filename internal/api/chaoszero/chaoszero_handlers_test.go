package chaoszero

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/PredictionExplorer/augur-explorer/internal/api/httpx"
	"github.com/PredictionExplorer/augur-explorer/internal/store"
)

func invokeHandler(method, target string, handler httpx.HandlerFunc) *httptest.ResponseRecorder {
	request := httptest.NewRequest(method, target, strings.NewReader(""))
	recorder := httptest.NewRecorder()
	handler(httpx.NewContext(recorder, request))
	return recorder
}

// routedRecorder sends a request through a real router so path parameters
// resolve, using the module's own route registrations.
func routedRecorder(t *testing.T, a *API, target string) *httptest.ResponseRecorder {
	t.Helper()
	r := httpx.NewRouter()
	a.RegisterRoutes(r)
	request := httptest.NewRequest(http.MethodGet, target, nil)
	recorder := httptest.NewRecorder()
	r.ServeHTTP(recorder, request)
	return recorder
}

func decodeEnvelope(t *testing.T, recorder *httptest.ResponseRecorder) map[string]any {
	t.Helper()
	var body map[string]any
	if err := json.Unmarshal(recorder.Body.Bytes(), &body); err != nil {
		t.Fatalf("decode response %q: %v", recorder.Body.String(), err)
	}
	return body
}

// TestHandlersRejectUnconfiguredDatabase drives every registered handler of
// a bare module: each must answer the shared guard envelope without touching
// a database.
func TestHandlersRejectUnconfiguredDatabase(t *testing.T) {
	a := NewBare()

	handlers := []struct {
		name    string
		handler httpx.HandlerFunc
	}{
		{name: "statistics", handler: a.handleStatistics},
		{name: "contracts", handler: a.handleContracts},
		{name: "rounds list", handler: a.handleRoundsList},
		{name: "round info", handler: a.handleRoundInfo},
		{name: "trades list", handler: a.handleTradesList},
		{name: "trades by round", handler: a.handleTradesByRound},
		{name: "trades by user", handler: a.handleTradesByUser},
		{name: "traders list", handler: a.handleTradersList},
		{name: "funding list", handler: a.handleFundingList},
		{name: "splits list", handler: a.handleSplitsList},
		{name: "merges list", handler: a.handleMergesList},
		{name: "settlements list", handler: a.handleSettlementsList},
		{name: "claims list", handler: a.handleClaimsList},
		{name: "claims by user", handler: a.handleClaimsByUser},
		{name: "answers list", handler: a.handleAnswersList},
		{name: "escrows list", handler: a.handleEscrowsList},
		{name: "payouts list", handler: a.handlePayoutsList},
		{name: "transfers list", handler: a.handleTransfersList},
		{name: "transfers by user", handler: a.handleTransfersByUser},
		{name: "approvals list", handler: a.handleApprovalsList},
	}

	const wantBody = `{"error":"Database link wasn't configured","status":0}`
	for _, tt := range handlers {
		t.Run(tt.name, func(t *testing.T) {
			recorder := invokeHandler(http.MethodGet, "/", tt.handler)
			if recorder.Code != http.StatusBadRequest {
				t.Fatalf("status = %d, want %d", recorder.Code, http.StatusBadRequest)
			}
			if got := recorder.Body.String(); got != wantBody {
				t.Fatalf("body = %q, want %q", got, wantBody)
			}
			if got := recorder.Header().Get("Access-Control-Allow-Origin"); got != "*" {
				t.Fatalf("Access-Control-Allow-Origin = %q, want *", got)
			}
		})
	}
}

// TestRegisterRoutesServesEveryEndpoint pins the route table: all twenty
// ChaosZero paths register, and a bare module answers each through the
// router (guard envelope, never a 404).
func TestRegisterRoutesServesEveryEndpoint(t *testing.T) {
	a := NewBare()
	r := httpx.NewRouter()
	a.RegisterRoutes(r)

	if got := len(r.Routes()); got != 20 {
		t.Errorf("registered routes = %d, want 20", got)
	}

	paths := []string{
		"/api/chaoszero/statistics",
		"/api/chaoszero/contracts",
		"/api/chaoszero/rounds/list/0/10",
		"/api/chaoszero/rounds/info/0",
		"/api/chaoszero/trades/list/0/10",
		"/api/chaoszero/trades/by_round/0/0/10",
		"/api/chaoszero/trades/by_user/0x2100000000000000000000000000000000000021/0/10",
		"/api/chaoszero/traders/list/0/10",
		"/api/chaoszero/funding/list/0/10",
		"/api/chaoszero/splits/list/0/10",
		"/api/chaoszero/merges/list/0/10",
		"/api/chaoszero/settlements/list/0/10",
		"/api/chaoszero/claims/list/0/10",
		"/api/chaoszero/claims/by_user/0x2100000000000000000000000000000000000021/0/10",
		"/api/chaoszero/answers/list/0/10",
		"/api/chaoszero/escrows/list/0/10",
		"/api/chaoszero/payouts/list/0/10",
		"/api/chaoszero/transfers/list/0/10",
		"/api/chaoszero/transfers/by_user/0x2400000000000000000000000000000000000024/0/10",
		"/api/chaoszero/approvals/list/0/10",
	}
	for _, path := range paths {
		request := httptest.NewRequest(http.MethodGet, path, nil)
		recorder := httptest.NewRecorder()
		r.ServeHTTP(recorder, request)
		if recorder.Code == http.StatusNotFound {
			t.Errorf("%s: not routed (404)", path)
			continue
		}
		body := decodeEnvelope(t, recorder)
		if body["status"] != float64(0) || body["error"] != "Database link wasn't configured" {
			t.Errorf("%s: body = %v, want the bare-module guard envelope", path, body)
		}
	}
}

// TestParameterValidationRunsBeforeStoreAccess drives the parameter guards
// of a module whose store is a zero value: a malformed path parameter must
// be rejected before any database call is attempted.
func TestParameterValidationRunsBeforeStoreAccess(t *testing.T) {
	a := New(new(store.Store), Options{})

	cases := []struct {
		name      string
		target    string
		wantError string
	}{
		{
			name:      "round info rejects non-numeric round",
			target:    "/api/chaoszero/rounds/info/not-a-number",
			wantError: "parameter",
		},
		{
			name:      "trades by round rejects non-numeric round",
			target:    "/api/chaoszero/trades/by_round/not-a-number/0/10",
			wantError: "parameter",
		},
		{
			name:      "rounds list rejects non-numeric offset",
			target:    "/api/chaoszero/rounds/list/not-a-number/10",
			wantError: "parameter",
		},
		{
			name:      "trades by user rejects non-numeric offset",
			target:    "/api/chaoszero/trades/by_user/0x2100000000000000000000000000000000000021/not-a-number/10",
			wantError: "parameter",
		},
	}
	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			recorder := routedRecorder(t, a, tt.target)
			if recorder.Code != http.StatusBadRequest {
				t.Fatalf("status = %d, want %d\n%s", recorder.Code, http.StatusBadRequest, recorder.Body.String())
			}
			body := decodeEnvelope(t, recorder)
			errStr, _ := body["error"].(string)
			if body["status"] != float64(0) || !strings.Contains(strings.ToLower(errStr), tt.wantError) {
				t.Fatalf("body = %v, want status 0 mentioning %q", body, tt.wantError)
			}
		})
	}
}
