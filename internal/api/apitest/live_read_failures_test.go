//go:build integration

package apitest

import (
	"context"
	"encoding/json"
	"errors"
	"io"
	"log"
	"math/big"
	"net/http"
	"os"
	"strings"
	"testing"

	ethcommon "github.com/ethereum/go-ethereum/common"

	"github.com/PredictionExplorer/augur-explorer/internal/api/faq"
	"github.com/PredictionExplorer/augur-explorer/internal/api/httpx"
	"github.com/PredictionExplorer/augur-explorer/internal/testchain"
)

// addrHex parses a fixture address literal.
func addrHex(s string) ethcommon.Address { return ethcommon.HexToAddress(s) }

// rpcFail makes one stubbed contract method revert for the duration of a
// case; restore re-installs the fixture value afterwards.
func rpcFail([]any) ([]any, error) {
	return nil, errors.New("forced rpc failure")
}

// decodeEnvelope parses the legacy {"status":..,"error":..} body.
func decodeEnvelope(t *testing.T, body []byte) (int, string) {
	t.Helper()
	var envelope struct {
		Status int    `json:"status"`
		Error  string `json:"error"`
	}
	if err := json.Unmarshal(body, &envelope); err != nil {
		t.Fatalf("response is not the legacy envelope: %v\n%s", err, body)
	}
	return envelope.Status, envelope.Error
}

// TestLiveContractReadFailures drives every request-time RPC read of the
// frozen v1 API through a per-method revert: each handler must answer with
// the legacy HTTP-400 envelope naming the failed stage (these are live
// contract reads, not cached state — the degraded-cache sentinels are pinned
// separately by the contractstate unit tests).
func TestLiveContractReadFailures(t *testing.T) {
	h := server(t)

	cases := []struct {
		name    string
		method  string
		restore func()
		path    string
		wantErr string
	}{
		{
			name: "cst_price next-price read fails", method: "getNextCstBidPrice",
			restore: func() { h.gameStub.Return("getNextCstBidPrice", wei("55000000000000000000")) },
			path:    "/api/cosmicgame/bid/cst_price", wantErr: "forced rpc failure",
		},
		{
			name: "cst_price auction-durations read fails", method: "getCstDutchAuctionDurations",
			restore: func() {
				h.gameStub.Return("getCstDutchAuctionDurations", big.NewInt(28800), big.NewInt(3600))
			},
			path: "/api/cosmicgame/bid/cst_price", wantErr: "forced rpc failure",
		},
		{
			name: "eth_price next-price read fails", method: "getNextEthBidPrice",
			restore: func() { h.gameStub.Return("getNextEthBidPrice", wei("1010000000000000")) },
			path:    "/api/cosmicgame/bid/eth_price", wantErr: "forced rpc failure",
		},
		{
			name: "eth_price auction-durations read fails", method: "getEthDutchAuctionDurations",
			restore: func() {
				h.gameStub.Return("getEthDutchAuctionDurations", big.NewInt(86400), big.NewInt(7200))
			},
			path: "/api/cosmicgame/bid/eth_price", wantErr: "forced rpc failure",
		},
		{
			name: "special winners champion read fails", method: "tryGetCurrentChampions",
			restore: func() {
				h.gameStub.Return("tryGetCurrentChampions",
					addrHex(addrBob), big.NewInt(600), addrHex(addrCarol), big.NewInt(800))
			},
			path: "/api/cosmicgame/bid/current_special_winners", wantErr: "TryGetCurrentChampions",
		},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			h.gameStub.Handle(tc.method, rpcFail)
			t.Cleanup(tc.restore)

			w := h.get(t, tc.path)
			if w.Code != http.StatusBadRequest {
				t.Fatalf("status = %d, want 400\n%s", w.Code, w.Body.String())
			}
			status, errStr := decodeEnvelope(t, w.Body.Bytes())
			if status != 0 || !strings.Contains(errStr, tc.wantErr) {
				t.Fatalf("envelope = (%d, %q), want status 0 with %q", status, errStr, tc.wantErr)
			}
		})
	}

	// The token metadata surface: each of the four reads has its own guard.
	tokenReads := []struct {
		method  string
		wantErr string
		restore func()
	}{
		{"name", "Error reading token name", func() { h.tokenStub.Return("name", "CosmicToken") }},
		{"symbol", "Error reading token symbol", func() { h.tokenStub.Return("symbol", "CST") }},
		{"decimals", "Error reading decimals", func() { h.tokenStub.Return("decimals", uint8(18)) }},
		{"game", "Error reading game address", func() {
			h.tokenStub.Return("game", addrHex("0x2000000000000000000000000000000000000002"))
		}},
	}
	for _, tc := range tokenReads {
		t.Run("ct_statistics "+tc.method+" fails", func(t *testing.T) {
			h.tokenStub.Handle(tc.method, rpcFail)
			t.Cleanup(tc.restore)

			w := h.get(t, "/api/cosmicgame/ct/statistics")
			if w.Code != http.StatusBadRequest {
				t.Fatalf("status = %d, want 400\n%s", w.Code, w.Body.String())
			}
			if status, errStr := decodeEnvelope(t, w.Body.Bytes()); status != 0 || !strings.Contains(errStr, tc.wantErr) {
				t.Fatalf("envelope = (%d, %q), want %q", status, errStr, tc.wantErr)
			}
		})
	}

	// The marketing-wallet configuration surface.
	marketingReads := []struct {
		method  string
		wantErr string
		restore func()
	}{
		{"treasurerAddress", "Error reading TreasurerAddress", func() { h.marketingStub.Return("treasurerAddress", addrHex(addrEmma)) }},
		{"token", "Error reading Token address", func() {
			h.marketingStub.Return("token", addrHex("0x4000000000000000000000000000000000000004"))
		}},
		{"owner", "Error reading Owner address", func() { h.marketingStub.Return("owner", addrHex(addrAlice)) }},
	}
	for _, tc := range marketingReads {
		t.Run("marketing_config "+tc.method+" fails", func(t *testing.T) {
			h.marketingStub.Handle(tc.method, rpcFail)
			t.Cleanup(tc.restore)

			w := h.get(t, "/api/cosmicgame/marketing/config/current")
			if w.Code != http.StatusBadRequest {
				t.Fatalf("status = %d, want 400\n%s", w.Code, w.Body.String())
			}
			if status, errStr := decodeEnvelope(t, w.Body.Bytes()); status != 0 || !strings.Contains(errStr, tc.wantErr) {
				t.Fatalf("envelope = (%d, %q), want %q", status, errStr, tc.wantErr)
			}
		})
	}

	t.Run("user balances token read fails", func(t *testing.T) {
		h.tokenStub.Handle("balanceOf", rpcFail)
		t.Cleanup(func() { h.tokenStub.Return("balanceOf", wei("150000000000000000000")) })

		w := h.get(t, "/api/cosmicgame/user/balances/"+addrAlice)
		if w.Code != http.StatusBadRequest {
			t.Fatalf("status = %d, want 400\n%s", w.Code, w.Body.String())
		}
		if status, errStr := decodeEnvelope(t, w.Body.Bytes()); status != 0 || !strings.Contains(errStr, "Error at BalanceOf() call") {
			t.Fatalf("envelope = (%d, %q)", status, errStr)
		}
	})

	t.Run("prize time keeps 200 shape when the contract is unreachable", func(t *testing.T) {
		h.gameStub.Handle("mainPrizeTime", rpcFail)
		t.Cleanup(func() {
			h.gameStub.Return("mainPrizeTime", big.NewInt(int64(testchain.BlockTime(chainTipBlock))+3600))
		})

		w := h.get(t, "/api/cosmicgame/rounds/current/time")
		if w.Code != http.StatusOK {
			t.Fatalf("status = %d, want the legacy 200-with-error shape\n%s", w.Code, w.Body.String())
		}
		var resp struct {
			Status            int             `json:"status"`
			Error             string          `json:"error"`
			CurRoundPrizeTime json.RawMessage `json:"CurRoundPrizeTime"`
		}
		if err := json.Unmarshal(w.Body.Bytes(), &resp); err != nil {
			t.Fatalf("decode: %v", err)
		}
		if resp.Status != 0 || !strings.Contains(resp.Error, "MainPrizeTime call failed") {
			t.Fatalf("response = %+v", resp)
		}
		if string(resp.CurRoundPrizeTime) != "0" {
			t.Fatalf("CurRoundPrizeTime = %s, want 0", resp.CurRoundPrizeTime)
		}
	})
}

// TestDashboardLiveBidPriceFallback pins the dashboard's cache-recovery
// path: a failed cache refresh leaves the "error" sentinel, and the next
// dashboard request re-reads the live price and writes it back to the cache.
func TestDashboardLiveBidPriceFallback(t *testing.T) {
	h := server(t)

	// Poison the cached bid price via a failing refresh.
	h.gameStub.Handle("getNextEthBidPrice", rpcFail)
	h.state.LoadInitial(context.Background())
	t.Cleanup(func() {
		h.gameStub.Return("getNextEthBidPrice", wei("1010000000000000"))
		h.state.LoadInitial(context.Background())
	})

	// While the RPC stays down, the dashboard reports the sentinel: the
	// live fallback cannot improve on the failed cache.
	w := h.get(t, "/api/cosmicgame/statistics/dashboard")
	if w.Code != http.StatusOK {
		t.Fatalf("dashboard while RPC down = %d\n%s", w.Code, w.Body.String())
	}
	var resp struct {
		BidPrice    string  `json:"BidPrice"`
		BidPriceEth float64 `json:"BidPriceEth"`
	}
	if err := json.Unmarshal(w.Body.Bytes(), &resp); err != nil {
		t.Fatalf("decode: %v", err)
	}
	if resp.BidPrice != "error" {
		t.Fatalf("BidPrice while RPC down = %q, want the error sentinel", resp.BidPrice)
	}

	// RPC recovers, but no refresh has run yet: the dashboard's live
	// fallback must fetch the price itself and write it back to the cache.
	h.gameStub.Return("getNextEthBidPrice", wei("1010000000000000"))
	w = h.get(t, "/api/cosmicgame/statistics/dashboard")
	if err := json.Unmarshal(w.Body.Bytes(), &resp); err != nil {
		t.Fatalf("decode recovered dashboard: %v", err)
	}
	if resp.BidPrice != "1010000000000000" {
		t.Fatalf("BidPrice after recovery = %q, want the live value", resp.BidPrice)
	}
	if snap := h.state.Snapshot(); snap.BidPrice != "1010000000000000" {
		t.Fatalf("cache write-back missing: snapshot BidPrice = %q", snap.BidPrice)
	}
}

// TestBanUnbanStoreFailures covers the moderation mutations' repository
// failure arms: a cancelled request context makes the insert/delete fail and
// the handler answers the legacy 400 envelope.
func TestBanUnbanStoreFailures(t *testing.T) {
	h := server(t)
	t.Setenv("ADMIN_API_KEY", adminKey)
	cancelled, cancel := context.WithCancel(context.Background())
	cancel()

	cases := []struct {
		path    string
		payload string
		wantErr string
	}{
		{"/api/cosmicgame/ban_bid", `{"bid_id":2011,"user_addr":"` + addrBob + `"}`, "Failed to insert banned bid"},
		{"/api/cosmicgame/unban_bid", `{"bid_id":2011}`, "Failed to unban bid"},
	}
	for _, tc := range cases {
		t.Run(tc.path, func(t *testing.T) {
			w := h.do(t, request{
				method:  http.MethodPost,
				path:    tc.path,
				body:    strings.NewReader(tc.payload),
				headers: map[string]string{"X-Admin-Key": adminKey},
				ctx:     cancelled,
			})
			if w.Code != http.StatusBadRequest {
				t.Fatalf("status = %d, want 400\n%s", w.Code, w.Body.String())
			}
			if status, errStr := decodeEnvelope(t, w.Body.Bytes()); status != 0 || !strings.Contains(errStr, tc.wantErr) {
				t.Fatalf("envelope = (%d, %q), want %q", status, errStr, tc.wantErr)
			}
		})
	}
}

// TestFAQProxy covers the POST proxy routes end to end (body and Accept
// forwarding, upstream content-type passthrough), the upstream-down 502 and
// the Init environment fallbacks that only fire outside the harness default.
func TestFAQProxy(t *testing.T) {
	h := server(t)
	discard := log.New(io.Discard, "", 0)
	stubURL := os.Getenv("AI_BOT_BACKEND_URL")

	t.Run("query forwards body and accept", func(t *testing.T) {
		w := h.do(t, request{
			method:  http.MethodPost,
			path:    "/api/cosmicgame/faq/query",
			body:    strings.NewReader(`{"question":"what is a round?"}`),
			headers: map[string]string{"Accept": "application/json"},
		})
		if w.Code != http.StatusOK {
			t.Fatalf("query = %d\n%s", w.Code, w.Body.String())
		}
		var resp struct {
			Answer   string `json:"answer"`
			Received string `json:"received"`
			Accept   string `json:"accept"`
		}
		if err := json.Unmarshal(w.Body.Bytes(), &resp); err != nil {
			t.Fatalf("decode: %v", err)
		}
		if resp.Answer != "stub answer" || resp.Received != `{"question":"what is a round?"}` {
			t.Fatalf("upstream did not receive the forwarded body: %+v", resp)
		}
		if resp.Accept != "application/json" {
			t.Fatalf("Accept header not forwarded: %+v", resp)
		}
		if ct := contentTypeOf(w); ct != "application/json" {
			t.Fatalf("Content-Type = %q", ct)
		}
	})

	t.Run("reindex proxies", func(t *testing.T) {
		w := h.do(t, request{method: http.MethodPost, path: "/api/cosmicgame/faq/reindex"})
		if w.Code != http.StatusOK || !strings.Contains(w.Body.String(), "reindexed") {
			t.Fatalf("reindex = %d\n%s", w.Code, w.Body.String())
		}
	})

	t.Run("upstream down answers 502", func(t *testing.T) {
		t.Setenv("AI_BOT_BACKEND_URL", "http://127.0.0.1:1")
		faq.Init(discard, discard, true)
		t.Cleanup(func() {
			_ = os.Setenv("AI_BOT_BACKEND_URL", stubURL)
			faq.Init(discard, discard, true)
		})

		w := h.do(t, request{method: http.MethodPost, path: "/api/cosmicgame/faq/query", body: strings.NewReader(`{}`)})
		if w.Code != http.StatusBadGateway {
			t.Fatalf("status = %d, want 502\n%s", w.Code, w.Body.String())
		}
		if !strings.Contains(w.Body.String(), "FAQ service unavailable") {
			t.Fatalf("body = %s", w.Body.String())
		}
	})

	t.Run("legacy env alias selects the upstream", func(t *testing.T) {
		t.Setenv("AI_BOT_BACKEND_URL", "")
		t.Setenv("FAQ_BOT_UPSTREAM_URL", stubURL)
		faq.Init(discard, discard, true)
		t.Cleanup(func() {
			_ = os.Setenv("AI_BOT_BACKEND_URL", stubURL)
			faq.Init(discard, discard, true)
		})

		w := h.get(t, "/api/cosmicgame/faq/health")
		if w.Code != http.StatusOK || !strings.Contains(w.Body.String(), "faq-bot-stub") {
			t.Fatalf("health through legacy alias = %d\n%s", w.Code, w.Body.String())
		}
	})

	t.Run("disabled module registers nothing", func(t *testing.T) {
		faq.Init(discard, discard, false)
		t.Cleanup(func() { faq.Init(discard, discard, true) })
		if faq.Enabled {
			t.Fatal("Init(false) left the module enabled")
		}
		r := httpx.NewRouter()
		faq.RegisterAPIRoutes(r)
		if n := len(r.Routes()); n != 0 {
			t.Fatalf("disabled module registered %d routes", n)
		}
	})
}
