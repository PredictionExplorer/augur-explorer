//go:build integration

package apitest

import (
	"context"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/ethereum/go-ethereum/accounts"
	"github.com/ethereum/go-ethereum/crypto"

	"github.com/PredictionExplorer/augur-explorer/internal/api/cosmicgame"
	"github.com/PredictionExplorer/augur-explorer/internal/api/httpx"
	"github.com/PredictionExplorer/augur-explorer/internal/api/randomwalk"
	"github.com/PredictionExplorer/augur-explorer/internal/api/routes"
)

// voterPrivKeyHex is a fixed test wallet so the add_game flow (and its
// cleanup) is deterministic. Address: derived below, never used on any chain.
const voterPrivKeyHex = "9b3f7c2a51d84e6f0a1c5b28d97e3f4a6c80b1d2e5f74a3908c6d1e2f3a4b5c6"

// postJSON marshals body and POSTs it with the given headers.
func postJSON(t *testing.T, h *harness, path string, body any, headers map[string]string) (int, map[string]any) {
	t.Helper()
	raw, err := json.Marshal(body)
	if err != nil {
		t.Fatalf("marshaling request body: %v", err)
	}
	w := h.do(t, request{
		method:  "POST",
		path:    path,
		body:    strings.NewReader(string(raw)),
		headers: headers,
	})
	var decoded map[string]any
	if len(w.Body.Bytes()) > 0 {
		if err := json.Unmarshal(w.Body.Bytes(), &decoded); err != nil {
			t.Fatalf("POST %s: response is not a JSON object: %v (body %q)", path, err, w.Body.String())
		}
	}
	return w.Code, decoded
}

// requireField asserts a response field equals want.
func requireField(t *testing.T, body map[string]any, field string, want any) {
	t.Helper()
	got, ok := body[field]
	if !ok {
		t.Fatalf("response is missing field %q: %v", field, body)
	}
	if fmt.Sprintf("%v", got) != fmt.Sprintf("%v", want) {
		t.Fatalf("response field %q = %v, want %v", field, got, want)
	}
}

// keylessRouter builds a second fully wired router whose modules carry no
// admin keys, over the harness's store and fake chain. The v1 modules are
// injected values, so the fail-closed matrix runs against the production
// route wiring rather than a synthetic middleware test.
func keylessRouter(t *testing.T, h *harness) *httpx.Router {
	t.Helper()
	cgAPI, err := cosmicgame.New(context.Background(), cosmicgame.Config{
		Store:     h.store,
		EthClient: h.ethClient,
		RPCClient: h.rpcClient,
	})
	if err != nil {
		t.Fatalf("building keyless cosmicgame module: %v", err)
	}
	return routes.New(h.store, routes.Options{
		CosmicGame: cgAPI,
		RandomWalk: randomwalk.New(h.store, randomwalk.Options{}),
	})
}

// postJSONRouter marshals body and POSTs it directly through r.
func postJSONRouter(t *testing.T, r *httpx.Router, path string, body any, headers map[string]string) (int, map[string]any) {
	t.Helper()
	raw, err := json.Marshal(body)
	if err != nil {
		t.Fatalf("marshaling request body: %v", err)
	}
	req := httptest.NewRequest(http.MethodPost, path, strings.NewReader(string(raw)))
	req.Header.Set("Content-Type", "application/json")
	for k, v := range headers {
		req.Header.Set(k, v)
	}
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	var decoded map[string]any
	if len(w.Body.Bytes()) > 0 {
		if err := json.Unmarshal(w.Body.Bytes(), &decoded); err != nil {
			t.Fatalf("POST %s: response is not a JSON object: %v (body %q)", path, err, w.Body.String())
		}
	}
	return w.Code, decoded
}

// TestRandomWalkMetadataFlatAssetPaths pins the NFT_ASSETS_FLAT_PATHS
// configuration through the module seam: a flat-layout deployment emits
// /images/<file> metadata URLs instead of /images/randomwalk/<file>.
func TestRandomWalkMetadataFlatAssetPaths(t *testing.T) {
	h := server(t)
	r := routes.New(h.store, routes.Options{
		RandomWalk: randomwalk.New(h.store, randomwalk.Options{AssetsFlatPaths: true}),
	})

	req := httptest.NewRequest(http.MethodGet, "/api/randomwalk/metadata/10", nil)
	req.RemoteAddr = "10.88.88.88:4242"
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	if w.Code != http.StatusOK {
		t.Fatalf("metadata = %d\n%s", w.Code, w.Body.String())
	}
	var meta struct {
		Image        string `json:"image"`
		AnimationURL string `json:"animation_url"`
	}
	if err := json.Unmarshal(w.Body.Bytes(), &meta); err != nil {
		t.Fatal(err)
	}
	if !strings.HasSuffix(meta.Image, "/images/000010_black.png") ||
		!strings.HasSuffix(meta.AnimationURL, "/images/000010_black_single.mp4") {
		t.Fatalf("flat layout URLs = %q, %q", meta.Image, meta.AnimationURL)
	}
	if strings.Contains(meta.Image, "/randomwalk/") {
		t.Fatalf("flat layout kept the nested segment: %q", meta.Image)
	}
}

// adminAuthMatrix drives the fail-closed admin-key middleware through its
// failure outcomes for one route: a keyless deployment answers 503, a wrong
// key 401. The success request is issued by the caller against the harness
// router (whose modules carry the test admin keys).
func adminAuthMatrix(t *testing.T, h *harness, path, header, envVar string, payload any) {
	t.Helper()
	t.Run("no_key_configured_503", func(t *testing.T) {
		code, body := postJSONRouter(t, keylessRouter(t, h), path, payload, map[string]string{header: adminKey})
		if code != http.StatusServiceUnavailable {
			t.Fatalf("expected 503 when no admin key is configured, got %d (%v)", code, body)
		}
		requireField(t, body, "status", 0)
		if errMsg, _ := body["error"].(string); !strings.Contains(errMsg, envVar) {
			t.Fatalf("disabled message should name %s, got %q", envVar, errMsg)
		}
	})

	t.Run("wrong_key_401", func(t *testing.T) {
		code, body := postJSON(t, h, path, payload, map[string]string{header: "not-the-key"})
		if code != http.StatusUnauthorized {
			t.Fatalf("expected 401 for a wrong admin key, got %d (%v)", code, body)
		}
		requireField(t, body, "status", 0)
	})
}

// TestBanUnbanBid covers the moderation flow end to end: auth matrix, ban,
// visibility through GET /api/cosmicgame/get_banned_bids, unban, restored
// state. The fixture ban of bid 2002 must be intact afterwards.
func TestBanUnbanBid(t *testing.T) {
	h := server(t)

	banPayload := map[string]any{"bid_id": 2011, "user_addr": addrBob}
	adminAuthMatrix(t, h, "/api/cosmicgame/ban_bid", "X-Admin-Key", "ADMIN_API_KEY", banPayload)

	t.Run("ban_then_unban", func(t *testing.T) {
		auth := map[string]string{"X-Admin-Key": adminKey}

		code, body := postJSON(t, h, "/api/cosmicgame/ban_bid", banPayload, auth)
		if code != http.StatusCreated {
			t.Fatalf("ban_bid: expected 201, got %d (%v)", code, body)
		}
		requireField(t, body, "result", "success")

		if ids := bannedBidIDs(t, h); !ids[2002] || !ids[2011] || len(ids) != 2 {
			t.Fatalf("after ban: expected banned bids {2002, 2011}, got %v", ids)
		}

		code, body = postJSON(t, h, "/api/cosmicgame/unban_bid", map[string]any{"bid_id": 2011}, auth)
		if code != http.StatusCreated {
			t.Fatalf("unban_bid: expected 201, got %d (%v)", code, body)
		}
		requireField(t, body, "result", "success")

		if ids := bannedBidIDs(t, h); !ids[2002] || len(ids) != 1 {
			t.Fatalf("after unban: expected banned bids {2002}, got %v", ids)
		}
	})

	t.Run("invalid_payload_error_envelope", func(t *testing.T) {
		code, body := postJSON(t, h, "/api/cosmicgame/ban_bid",
			map[string]any{"bid_id": "not-a-number"},
			map[string]string{"X-Admin-Key": adminKey})
		if code != http.StatusBadRequest {
			t.Fatalf("expected 400 for malformed ban payload, got %d (%v)", code, body)
		}
		requireField(t, body, "status", 0)
		requireField(t, body, "error", "Invalid JSON: bid_id and user_addr required")
	})
}

// bannedBidIDs reads the moderation list through the public API. The route
// returns a bare array of records (FastAPI parity), not the envelope.
func bannedBidIDs(t *testing.T, h *harness) map[int64]bool {
	t.Helper()
	w := h.get(t, "/api/cosmicgame/get_banned_bids")
	if w.Code != http.StatusOK {
		t.Fatalf("get_banned_bids: %d %s", w.Code, w.Body.String())
	}
	var recs []struct {
		BidID int64 `json:"bid_id"`
	}
	if err := json.Unmarshal(w.Body.Bytes(), &recs); err != nil {
		t.Fatalf("parsing get_banned_bids: %v", err)
	}
	ids := make(map[int64]bool, len(recs))
	for _, b := range recs {
		ids[b.BidID] = true
	}
	return ids
}

// TestRankingMatchAdmin covers the direct (admin-only) Elo match route,
// asserting the exact Elo movement and restoring fixture ratings.
func TestRankingMatchAdmin(t *testing.T) {
	h := server(t)

	payload := map[string]any{"nft1": 12, "nft2": 13, "nft1_won": true}
	adminAuthMatrix(t, h, "/api/randomwalk/token-ranking/match", "X-Ranking-Admin-Key", "RANKING_ADMIN_KEY", payload)

	t.Run("records_match_and_moves_elo", func(t *testing.T) {
		defer restoreRankingFixture(t, h)

		code, body := postJSON(t, h, "/api/randomwalk/token-ranking/match", payload,
			map[string]string{"X-Ranking-Admin-Key": adminKey})
		if code != http.StatusOK {
			t.Fatalf("match: expected 200, got %d (%v)", code, body)
		}

		// Elo with K = 250 - 2*0.00525 (two fixture matches), ratings 1195 vs
		// 1205: winner must gain exactly what the loser drops.
		ra, _ := body["rating_nft1"].(float64)
		rb, _ := body["rating_nft2"].(float64)
		if !(ra > 1195 && rb < 1205) {
			t.Fatalf("elo did not move toward the winner: nft1 %v, nft2 %v", ra, rb)
		}
		if diff := (ra - 1195) - (1205 - rb); diff > 1e-9 || diff < -1e-9 {
			t.Fatalf("elo update is not zero-sum: gain %v, loss %v", ra-1195, 1205-rb)
		}

		if got := voteCount(t, h); got != 3 {
			t.Fatalf("vote_count after admin match: got %d, want 3", got)
		}
	})

	t.Run("rejects_identical_pair", func(t *testing.T) {
		code, body := postJSON(t, h, "/api/randomwalk/token-ranking/match",
			map[string]any{"nft1": 12, "nft2": 12, "nft1_won": true},
			map[string]string{"X-Ranking-Admin-Key": adminKey})
		if code != http.StatusBadRequest {
			t.Fatalf("expected 400 for nft1 == nft2, got %d (%v)", code, body)
		}
	})
}

// TestAddGameWalletSignature covers the public, wallet-signed beauty vote:
// challenge nonce, EIP-191 signature, duplicate-pair rejection, replayed
// nonce rejection and bad-signature rejection.
func TestAddGameWalletSignature(t *testing.T) {
	h := server(t)
	defer restoreRankingFixture(t, h)

	key, err := crypto.HexToECDSA(voterPrivKeyHex)
	if err != nil {
		t.Fatalf("parsing fixed test key: %v", err)
	}
	voter := crypto.PubkeyToAddress(key.PublicKey)

	const chainID = int64(42161)
	sign := func(nonce string, nft1, nft2, winner int64) string {
		msg := fmt.Sprintf(
			"RandomWalk beauty vote\nVersion: 1\nchainId: %d\nnonce: %s\nnft1: %d\nnft2: %d\nwinner: %d",
			chainID, nonce, nft1, nft2, winner,
		)
		sig, err := crypto.Sign(accounts.TextHash([]byte(msg)), key)
		if err != nil {
			t.Fatalf("signing vote message: %v", err)
		}
		return "0x" + hex.EncodeToString(sig)
	}

	nonce := fetchSignChallenge(t, h)
	code, body := postJSON(t, h, "/api/randomwalk/add_game", map[string]any{
		"nft1": 10, "nft2": 13, "nft1_win": 1,
		"sign_nonce": nonce, "signature": sign(nonce, 10, 13, 10), "chain_id": chainID,
	}, nil)
	if code != http.StatusOK {
		t.Fatalf("add_game: expected 200, got %d (%v)", code, body)
	}
	requireField(t, body, "result", "success")
	if got := voteCount(t, h); got != 3 {
		t.Fatalf("vote_count after signed vote: got %d, want 3", got)
	}

	var voterAid int64
	if err := h.db.QueryRow(
		`SELECT m.voter_aid FROM rw_ranking_match m JOIN address a ON a.address_id = m.voter_aid WHERE a.addr = $1`,
		voter.Hex(),
	).Scan(&voterAid); err != nil {
		t.Fatalf("signed vote was not attributed to the wallet %s: %v", voter.Hex(), err)
	}

	t.Run("replayed_nonce_rejected", func(t *testing.T) {
		code, body := postJSON(t, h, "/api/randomwalk/add_game", map[string]any{
			"nft1": 11, "nft2": 12, "nft1_win": 1,
			"sign_nonce": nonce, "signature": sign(nonce, 11, 12, 11), "chain_id": chainID,
		}, nil)
		if code != http.StatusBadRequest {
			t.Fatalf("expected 400 for a replayed nonce, got %d (%v)", code, body)
		}
	})

	t.Run("duplicate_pair_rejected", func(t *testing.T) {
		fresh := fetchSignChallenge(t, h)
		code, body := postJSON(t, h, "/api/randomwalk/add_game", map[string]any{
			"nft1": 10, "nft2": 13, "nft1_win": 0,
			"sign_nonce": fresh, "signature": sign(fresh, 10, 13, 13), "chain_id": chainID,
		}, nil)
		if code != http.StatusConflict {
			t.Fatalf("expected 409 for voting twice on one pair, got %d (%v)", code, body)
		}
	})

	t.Run("tampered_vote_rejected", func(t *testing.T) {
		fresh := fetchSignChallenge(t, h)
		// Signature covers winner=11 but the posted vote claims nft1_win for 12.
		code, body := postJSON(t, h, "/api/randomwalk/add_game", map[string]any{
			"nft1": 12, "nft2": 11, "nft1_win": 1,
			"sign_nonce": fresh, "signature": sign(fresh, 12, 11, 11), "chain_id": chainID,
		}, nil)
		if code != http.StatusOK {
			t.Fatalf("expected the tampered vote to be attributed to the recovered (different) signer and stored, got %d (%v)", code, body)
		}
		// The store accepted it because personal_sign recovery yields a
		// different address; pin that this is one more vote.
		if got := voteCount(t, h); got != 4 {
			t.Fatalf("vote_count after tampered vote: got %d, want 4", got)
		}
	})

	t.Run("bad_chain_rejected", func(t *testing.T) {
		fresh := fetchSignChallenge(t, h)
		code, body := postJSON(t, h, "/api/randomwalk/add_game", map[string]any{
			"nft1": 11, "nft2": 13, "nft1_win": 1,
			"sign_nonce": fresh, "signature": sign(fresh, 11, 13, 11), "chain_id": -5,
		}, nil)
		if code != http.StatusBadRequest {
			t.Fatalf("expected 400 for a non-positive chain id, got %d (%v)", code, body)
		}
	})
}

// fetchSignChallenge requests a one-time voting nonce.
func fetchSignChallenge(t *testing.T, h *harness) string {
	t.Helper()
	w := h.get(t, "/api/randomwalk/ranking/sign-challenge")
	if w.Code != http.StatusOK {
		t.Fatalf("sign-challenge: %d %s", w.Code, w.Body.String())
	}
	var resp struct {
		Nonce string `json:"nonce"`
	}
	if err := json.Unmarshal(w.Body.Bytes(), &resp); err != nil || resp.Nonce == "" {
		t.Fatalf("sign-challenge returned no nonce: %v %s", err, w.Body.String())
	}
	return resp.Nonce
}

// voteCount reads GET /api/randomwalk/vote_count.
func voteCount(t *testing.T, h *harness) int64 {
	t.Helper()
	w := h.get(t, "/api/randomwalk/vote_count")
	if w.Code != http.StatusOK {
		t.Fatalf("vote_count: %d %s", w.Code, w.Body.String())
	}
	var resp struct {
		TotalCount int64 `json:"total_count"`
	}
	if err := json.Unmarshal(w.Body.Bytes(), &resp); err != nil {
		t.Fatalf("parsing vote_count: %v", err)
	}
	return resp.TotalCount
}

// restoreRankingFixture deletes any matches beyond the two fixture rows,
// resets the fixture Elo ratings, and removes voter address rows created by
// signed votes, so ordering between this package's tests (and -shuffle)
// cannot change golden results.
func restoreRankingFixture(t *testing.T, h *harness) {
	t.Helper()
	if _, err := h.db.Exec(`DELETE FROM rw_ranking_match WHERE id > 2`); err != nil {
		t.Fatalf("restoring rw_ranking_match: %v", err)
	}
	// Fixture addresses stop at address_id 28 (see seed/01_layer1.sql); any
	// higher id was created by a signed-vote test run.
	if _, err := h.db.Exec(`DELETE FROM address WHERE address_id > 28`); err != nil {
		t.Fatalf("restoring address table: %v", err)
	}
	if _, err := h.db.Exec(`
		UPDATE rw_token_ranking AS r SET rating = f.rating, updated_at = f.updated_at
		FROM (VALUES
			(10::bigint, 1210.5::double precision, TO_TIMESTAMP(1767232000)),
			(11::bigint, 1189.5::double precision, TO_TIMESTAMP(1767232000)),
			(12::bigint, 1195::double precision,   TO_TIMESTAMP(1767232100)),
			(13::bigint, 1205::double precision,   TO_TIMESTAMP(1767232100))
		) AS f(token_id, rating, updated_at)
		WHERE r.token_id = f.token_id`); err != nil {
		t.Fatalf("restoring rw_token_ranking: %v", err)
	}
	if _, err := h.db.Exec(`DELETE FROM rw_token_ranking WHERE token_id NOT IN (10,11,12,13)`); err != nil {
		t.Fatalf("restoring rw_token_ranking extra rows: %v", err)
	}
}
