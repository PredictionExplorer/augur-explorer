package randomwalk

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/PredictionExplorer/augur-explorer/internal/api/common"
	"github.com/PredictionExplorer/augur-explorer/internal/api/httpx"
	"github.com/PredictionExplorer/augur-explorer/internal/store"
	rwdb "github.com/PredictionExplorer/augur-explorer/internal/store/randomwalk"
)

func preserveRandomwalkTestState(t *testing.T) {
	t.Helper()

	oldRepo := rwRepo
	oldStore := rwStore
	oldRoutesEnabled := routesEnabled
	oldCommonContext := common.Ctx
	t.Cleanup(func() {
		rwRepo = oldRepo
		rwStore = oldStore
		routesEnabled = oldRoutesEnabled
		common.Ctx = oldCommonContext
	})
}

func invokeRandomwalkHandler(method, target, body string, handler httpx.HandlerFunc) *httptest.ResponseRecorder {
	request := httptest.NewRequest(method, target, strings.NewReader(body))
	recorder := httptest.NewRecorder()
	handler(httpx.NewContext(recorder, request))
	return recorder
}

func decodeResponseObject(t *testing.T, recorder *httptest.ResponseRecorder) map[string]any {
	t.Helper()

	var body map[string]any
	if err := json.Unmarshal(recorder.Body.Bytes(), &body); err != nil {
		t.Fatalf("decode response %q: %v", recorder.Body.String(), err)
	}
	return body
}

func TestRandomwalkHandlersRejectUnconfiguredDatabase(t *testing.T) {
	preserveRandomwalkTestState(t)
	rwRepo = nil
	rwStore = nil

	handlers := []struct {
		name    string
		method  string
		handler httpx.HandlerFunc
	}{
		{name: "current offers", method: http.MethodGet, handler: apiRwalkCurrentOffers},
		{name: "floor price", method: http.MethodGet, handler: apiRwalkFloorPrice},
		{name: "token list sequential", method: http.MethodGet, handler: apiRwalkTokenListSeq},
		{name: "token list period", method: http.MethodGet, handler: apiRwalkTokenListPeriod},
		{name: "token info", method: http.MethodGet, handler: apiRwalkTokenInfo},
		{name: "token history", method: http.MethodGet, handler: apiRwalkTokenHistory},
		{name: "token name history", method: http.MethodGet, handler: apiRwalkTokenNameHistory},
		{name: "tokens by user", method: http.MethodGet, handler: apiRwalkTokensByUser},
		{name: "trading history", method: http.MethodGet, handler: apiRwalkTradingHistory},
		{name: "sale history", method: http.MethodGet, handler: apiRwalkSaleHistory},
		{name: "trading by user", method: http.MethodGet, handler: apiRwalkTradingHistoryByUser},
		{name: "token stats", method: http.MethodGet, handler: apiRwalkTokenStats},
		{name: "market stats", method: http.MethodGet, handler: apiRwalkMarketStats},
		{name: "trading volume", method: http.MethodGet, handler: apiRwalkTradingVolumeByPeriod},
		{name: "mint intervals", method: http.MethodGet, handler: apiRwalkMintIntervals},
		{name: "withdrawal chart", method: http.MethodGet, handler: apiRwalkWithdrawalChart},
		{name: "floor price over time", method: http.MethodGet, handler: apiRwalkFloorPriceOverTime},
		{name: "top tokens", method: http.MethodGet, handler: apiRwalkTop5TradedTokens},
		{name: "mint report", method: http.MethodGet, handler: apiRwalkMintReport},
		{name: "user info", method: http.MethodGet, handler: apiRwalkUserInfo},
		{name: "contracts", method: http.MethodGet, handler: apiRwalkContracts},
		{name: "explore random", method: http.MethodGet, handler: apiRandomwalkExploreRandom},
		{name: "beauty pair", method: http.MethodGet, handler: apiRandomwalkRankingBeautyPairIDs},
		{name: "vote count", method: http.MethodGet, handler: apiRandomwalkVoteCount},
		{name: "ranking order", method: http.MethodGet, handler: apiRandomwalkTokenRankingOrder},
		{name: "ranking match", method: http.MethodPost, handler: apiRandomwalkTokenRankingMatch},
		{name: "ranking sign challenge", method: http.MethodGet, handler: apiRandomwalkRankingSignChallenge},
		{name: "add game", method: http.MethodPost, handler: apiRandomwalkAddGameLegacy},
		{name: "metadata exported handler", method: http.MethodGet, handler: TokenMetadataHandler},
	}

	const wantBody = `{"error":"Database link wasn't configured","status":0}`
	for _, tt := range handlers {
		t.Run(tt.name, func(t *testing.T) {
			recorder := invokeRandomwalkHandler(tt.method, "/", `{}`, tt.handler)
			if recorder.Code != http.StatusBadRequest {
				t.Fatalf("status = %d, want %d", recorder.Code, http.StatusBadRequest)
			}
			if got := recorder.Body.String(); got != wantBody {
				t.Fatalf("body = %q, want %q", got, wantBody)
			}
			if got := recorder.Header().Get("Content-Type"); got != "application/json; charset=utf-8" {
				t.Fatalf("Content-Type = %q, want application/json; charset=utf-8", got)
			}
		})
	}
}

func TestRankingHandlersValidateBeforeStoreAccess(t *testing.T) {
	preserveRandomwalkTestState(t)
	rwRepo = new(rwdb.Repo)
	rwStore = new(store.Store)
	t.Setenv("RANKING_VOTE_CHAIN_IDS", "1")

	tests := []struct {
		name          string
		handler       httpx.HandlerFunc
		body          string
		wantError     string
		errorContains string
	}{
		{
			name: "admin match malformed JSON", handler: apiRandomwalkTokenRankingMatch,
			body: `{`, errorContains: "unexpected EOF",
		},
		{
			name: "admin match wrong JSON type", handler: apiRandomwalkTokenRankingMatch,
			body: `{"nft1":"seven","nft2":9}`, errorContains: "cannot unmarshal string",
		},
		{
			name: "admin match negative token", handler: apiRandomwalkTokenRankingMatch,
			body: `{"nft1":-1,"nft2":9,"nft1_won":true}`, wantError: errRankingBadPair.Error(),
		},
		{
			name: "admin match repeated token", handler: apiRandomwalkTokenRankingMatch,
			body: `{"nft1":9,"nft2":9,"nft1_won":false}`, wantError: errRankingBadPair.Error(),
		},
		{
			name: "beauty vote malformed JSON", handler: apiRandomwalkAddGameLegacy,
			body: `not-json`, errorContains: "invalid character",
		},
		{
			name: "beauty vote winner below range", handler: apiRandomwalkAddGameLegacy,
			body: `{"nft1":1,"nft2":2,"nft1_win":-1}`, wantError: "nft1_win must be 0 or 1",
		},
		{
			name: "beauty vote winner above range", handler: apiRandomwalkAddGameLegacy,
			body: `{"nft1":1,"nft2":2,"nft1_win":2}`, wantError: "nft1_win must be 0 or 1",
		},
		{
			name: "beauty vote bad pair precedes credentials", handler: apiRandomwalkAddGameLegacy,
			body: `{"nft1":2,"nft2":2,"nft1_win":1}`, wantError: errRankingBadPair.Error(),
		},
		{
			name: "beauty vote credentials required", handler: apiRandomwalkAddGameLegacy,
			body: `{"nft1":1,"nft2":2,"nft1_win":1,"chain_id":1}`, wantError: errRankingVoteCredentialsRequired.Error(),
		},
		{
			name: "beauty vote chain rejected", handler: apiRandomwalkAddGameLegacy,
			body:      `{"nft1":1,"nft2":2,"nft1_win":1,"chain_id":42161,"sign_nonce":"n","signature":"00"}`,
			wantError: errRankingVoteChainNotAllowed.Error(),
		},
		{
			name: "beauty vote signature rejected", handler: apiRandomwalkAddGameLegacy,
			body:          `{"nft1":1,"nft2":2,"nft1_win":1,"chain_id":1,"sign_nonce":"n","signature":"zz"}`,
			errorContains: errRankingVoteInvalidSignature.Error(),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			recorder := invokeRandomwalkHandler(http.MethodPost, "/", tt.body, tt.handler)
			if recorder.Code != http.StatusBadRequest {
				t.Fatalf("status = %d, want %d; body = %q", recorder.Code, http.StatusBadRequest, recorder.Body.String())
			}
			body := decodeResponseObject(t, recorder)
			if len(body) != 1 {
				t.Fatalf("response keys = %v, want only error", body)
			}
			gotError, ok := body["error"].(string)
			if !ok {
				t.Fatalf("error field = %#v, want string", body["error"])
			}
			if tt.wantError != "" && gotError != tt.wantError {
				t.Fatalf("error = %q, want %q", gotError, tt.wantError)
			}
			if tt.errorContains != "" && !strings.Contains(gotError, tt.errorContains) {
				t.Fatalf("error = %q, want substring %q", gotError, tt.errorContains)
			}
		})
	}
}

func TestRespondRankingVoteErrorClassificationAndSecrecy(t *testing.T) {
	preserveRandomwalkTestState(t)
	common.Ctx = nil

	invalidSignature := fmt.Errorf("%w: bad recovery id", errRankingVoteInvalidSignature)
	const secret = "database signature column failed; sign_nonce=top-secret"
	tests := []struct {
		name      string
		err       error
		wantCode  int
		wantError string
	}{
		{
			name: "bad pair", err: errRankingBadPair,
			wantCode: http.StatusBadRequest, wantError: errRankingBadPair.Error(),
		},
		{
			name: "duplicate pair", err: errRankingDuplicateVoterPair,
			wantCode: http.StatusConflict, wantError: "already voted on this pair",
		},
		{
			name: "credentials required", err: errRankingVoteCredentialsRequired,
			wantCode: http.StatusBadRequest, wantError: errRankingVoteCredentialsRequired.Error(),
		},
		{
			name: "chain rejected", err: errRankingVoteChainNotAllowed,
			wantCode: http.StatusBadRequest, wantError: errRankingVoteChainNotAllowed.Error(),
		},
		{
			name: "invalid signature", err: invalidSignature,
			wantCode: http.StatusBadRequest, wantError: invalidSignature.Error(),
		},
		{
			name: "invalid or replayed nonce", err: errRankingVoteInvalidNonce,
			wantCode: http.StatusBadRequest, wantError: errRankingVoteInvalidNonce.Error(),
		},
		{
			name: "internal error with deceptive client words", err: errors.New(secret),
			wantCode: http.StatusInternalServerError, wantError: "Internal server error",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			recorder := invokeRandomwalkHandler(
				http.MethodPost,
				"/api/randomwalk/add_game",
				"",
				func(c *httpx.Context) { respondRankingVoteError(c, tt.err) },
			)
			if recorder.Code != tt.wantCode {
				t.Fatalf("status = %d, want %d", recorder.Code, tt.wantCode)
			}
			body := decodeResponseObject(t, recorder)
			if got := body["error"]; got != tt.wantError {
				t.Fatalf("error = %#v, want %q", got, tt.wantError)
			}
			if tt.wantCode == http.StatusInternalServerError {
				if got := body["status"]; got != float64(0) {
					t.Fatalf("status field = %#v, want 0", got)
				}
				if strings.Contains(recorder.Body.String(), "top-secret") {
					t.Fatalf("internal detail leaked in response: %q", recorder.Body.String())
				}
			} else if len(body) != 1 {
				t.Fatalf("response keys = %v, want only error", body)
			}
		})
	}
}

func TestRegisterAPIRoutesHonorsToggleAndKeepsRankingAliases(t *testing.T) {
	preserveRandomwalkTestState(t)

	routesEnabled = false
	disabledRouter := httpx.NewRouter()
	RegisterAPIRoutes(disabledRouter)
	if got := len(disabledRouter.Routes()); got != 0 {
		t.Fatalf("disabled route count = %d, want 0", got)
	}

	routesEnabled = true
	router := httpx.NewRouter()
	RegisterAPIRoutes(router)
	routes := router.Routes()
	if got, want := len(routes), 32; got != want {
		t.Fatalf("route count = %d, want %d", got, want)
	}

	registered := make(map[string]bool, len(routes))
	for _, route := range routes {
		registered[route.Method+" "+route.Pattern] = true
	}
	wantRankingRoutes := []string{
		"GET /api/randomwalk/explore/random",
		"GET /api/randomwalk/random",
		"GET /api/randomwalk/token-ranking/order",
		"GET /api/randomwalk/rating_order",
		"GET /api/randomwalk/vote_count",
		"GET /api/randomwalk/ranking/sign-challenge",
		"GET /api/randomwalk/ranking/beauty-pair-ids",
		"POST /api/randomwalk/token-ranking/match",
		"POST /api/randomwalk/add_game",
		"GET /api/randomwalk/metadata/{token_id}",
	}
	for _, route := range wantRankingRoutes {
		if !registered[route] {
			t.Errorf("ranking route %q is not registered", route)
		}
	}
}
