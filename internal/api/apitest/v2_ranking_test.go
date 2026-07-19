//go:build integration

package apitest

// V2 ranking slice suite: deterministic goldens for the four reads and the
// full write flows of the first v2 mutations (ADR-0008) — challenge →
// EIP-191 signed vote, admin match, replay/duplicate/tamper rejection,
// fail-closed admin auth and the spec-declared 429. Every mutation restores
// the shared ranking fixture so ordering (and -shuffle) cannot change any
// other golden.

import (
	"context"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io"
	"log/slog"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"

	"github.com/ethereum/go-ethereum/accounts"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/getkin/kin-openapi/openapi3"
	"github.com/getkin/kin-openapi/openapi3filter"
	"github.com/getkin/kin-openapi/routers"

	"github.com/PredictionExplorer/augur-explorer/internal/api/common"
	"github.com/PredictionExplorer/augur-explorer/internal/api/routes"
	apiv2 "github.com/PredictionExplorer/augur-explorer/internal/api/v2"
	"github.com/PredictionExplorer/augur-explorer/internal/beautyrank"
)

const (
	// #nosec G101 -- route templates, not credentials.
	v2RankingRandomTokens = "/api/v2/randomwalk/ranking/random-tokens"
	v2RankingPair         = "/api/v2/randomwalk/ranking/pair"
	v2RankingRatings      = "/api/v2/randomwalk/ranking/ratings"
	v2RankingStatistics   = "/api/v2/randomwalk/ranking/statistics"
	v2RankingChallenges   = "/api/v2/randomwalk/ranking/challenges"
	v2RankingVotes        = "/api/v2/randomwalk/ranking/votes"
	v2RankingMatches      = "/api/v2/randomwalk/ranking/matches"
)

// v2VoterPrivKeyHex is a dedicated wallet for the v2 vote flow, distinct
// from the v1 suite's voter so the two flows can never interfere through
// the (voter, pair) uniqueness constraint or the address cache.
// Address: 0x6f4c950442e1Af093BcfF730381E63Ae9171b87a.
const v2VoterPrivKeyHex = "0000000000000000000000000000000000000000000000000000000000000042"

// signV2RankingVote signs the canonical beauty-vote message for the v2
// vote request.
func signV2RankingVote(t *testing.T, chainID int64, nonce string, first, second, winner int64) (signature, signer string) {
	t.Helper()
	key, err := crypto.HexToECDSA(v2VoterPrivKeyHex)
	if err != nil {
		t.Fatalf("parsing v2 voter key: %v", err)
	}
	message := beautyrank.VoteMessage(chainID, nonce, first, second, winner)
	sig, err := crypto.Sign(accounts.TextHash([]byte(message)), key)
	if err != nil {
		t.Fatalf("signing vote: %v", err)
	}
	return hex.EncodeToString(sig), crypto.PubkeyToAddress(key.PublicKey).Hex()
}

// restoreV2RankingState reverts every ranking mutation this suite performs:
// fixture matches/ratings/addresses via the shared v1 helper (which also
// covers the voter rows signed votes create), the challenge nonces, and the
// store's address cache so deleted voter rows cannot leak stale ids into
// later tests.
func restoreV2RankingState(t *testing.T, h *harness) {
	t.Helper()
	restoreRankingFixture(t, h)
	if _, err := h.db.Exec(`DELETE FROM rw_ranking_vote_nonce`); err != nil {
		t.Fatalf("clearing ranking vote nonces: %v", err)
	}
	h.store.ResetAddressCache()
}

// v2RankingSpec loads and validates the embedded v2 spec once per test.
func v2RankingSpec(t *testing.T) *openapi3.T {
	t.Helper()
	spec, err := apiv2.GetSpec()
	if err != nil {
		t.Fatalf("loading embedded v2 spec: %v", err)
	}
	if err := spec.Validate(context.Background()); err != nil {
		t.Fatalf("validating embedded v2 spec: %v", err)
	}
	return spec
}

// validateV2MutationResponse validates one non-GET exchange against the
// embedded spec (request body included; the admin apiKey scheme passes
// through a noop authenticator since enforcement is middleware behavior
// under test elsewhere).
func validateV2MutationResponse(
	t *testing.T,
	spec *openapi3.T,
	method, template, target, requestBody string,
	headers map[string]string,
	response *httptest.ResponseRecorder,
) {
	t.Helper()
	pathItem := spec.Paths.Value(template)
	if pathItem == nil {
		t.Fatalf("spec has no path item for %s", template)
	}
	operation := pathItem.GetOperation(method)
	if operation == nil {
		t.Fatalf("spec has no %s operation for %s", method, template)
	}
	var body io.Reader
	if requestBody != "" {
		body = strings.NewReader(requestBody)
	}
	request := httptest.NewRequest(method, target, body)
	if requestBody != "" {
		request.Header.Set("Content-Type", "application/json")
	}
	for name, value := range headers {
		request.Header.Set(name, value)
	}
	route := &routers.Route{
		Spec:      spec,
		Path:      template,
		PathItem:  pathItem,
		Method:    method,
		Operation: operation,
	}
	requestInput := &openapi3filter.RequestValidationInput{
		Request:    request,
		PathParams: map[string]string{},
		Route:      route,
		Options: &openapi3filter.Options{
			AuthenticationFunc: openapi3filter.NoopAuthenticationFunc,
		},
	}
	if response.Code < http.StatusBadRequest && requestBody != "" {
		if err := openapi3filter.ValidateRequest(context.Background(), requestInput); err != nil {
			t.Fatalf("%s %s request violates OpenAPI v2: %v", method, target, err)
		}
	}
	responseInput := &openapi3filter.ResponseValidationInput{
		RequestValidationInput: requestInput,
		Status:                 response.Code,
		Header:                 response.Header(),
	}
	responseInput.SetBodyBytes(response.Body.Bytes())
	if err := openapi3filter.ValidateResponse(context.Background(), responseInput); err != nil {
		t.Fatalf("%s %s response violates OpenAPI v2 (status %d): %v\n%s",
			method, target, response.Code, err, response.Body.String())
	}
}

// pinV2MutationGolden validates a mutation exchange against the spec and
// pins it as a golden, applying redact to volatile fields first.
func pinV2MutationGolden(
	t *testing.T,
	spec *openapi3.T,
	name, template, target, requestBody string,
	headers map[string]string,
	recorder *httptest.ResponseRecorder,
	redact redactor,
) {
	t.Helper()
	validateV2MutationResponse(t, spec, http.MethodPost, template, target, requestBody, headers, recorder)
	body := canonicalJSON(t, recorder.Body.Bytes())
	if redact != nil {
		body = redact(t, body)
	}
	compareV2Golden(t, name, response{
		Status:      recorder.Code,
		ContentType: contentTypeOf(recorder),
		Body:        body,
	})
}

func postV2(t *testing.T, h *harness, path, body string, headers map[string]string) *httptest.ResponseRecorder {
	t.Helper()
	var reader io.Reader
	if body != "" {
		reader = strings.NewReader(body)
	}
	return h.do(t, request{method: http.MethodPost, path: path, body: reader, headers: headers})
}

// TestAPIV2RankingReads pins the four bounded read resources. The fixture
// selection is deterministic: every token has exactly one recorded match,
// so the explorer prefers ascending rating (11, 12, 13, 10) and no RANDOM()
// fallback exists on the v2 path.
func TestAPIV2RankingReads(t *testing.T) {
	h := server(t)
	spec := v2RankingSpec(t)

	firstPage := h.get(t, v2RankingRatings+"?limit=2")
	var page apiv2.RankingRatingPage
	decodeV2JSON(t, firstPage, &page)
	if page.Meta.NextCursor == nil {
		t.Fatal("two-row rating page did not return a continuation cursor")
	}

	cancelledCtx, cancel := context.WithCancel(context.Background())
	cancel()
	cases := []v2GoldenCase{
		{
			name:       "ranking_random_tokens",
			target:     v2RankingRandomTokens,
			template:   v2RankingRandomTokens,
			pathParams: map[string]string{},
		},
		{
			name:       "ranking_random_tokens_sample_size",
			target:     v2RankingRandomTokens + "?sampleSize=4",
			template:   v2RankingRandomTokens,
			pathParams: map[string]string{},
		},
		{
			name:       "ranking_random_tokens_error_sample_size",
			target:     v2RankingRandomTokens + "?sampleSize=101",
			template:   v2RankingRandomTokens,
			pathParams: map[string]string{},
		},
		{
			name:       "ranking_random_tokens_error_internal",
			target:     v2RankingRandomTokens,
			template:   v2RankingRandomTokens,
			pathParams: map[string]string{},
			ctx:        cancelledCtx,
		},
		{
			name:       "ranking_pair",
			target:     v2RankingPair,
			template:   v2RankingPair,
			pathParams: map[string]string{},
		},
		{
			// Alice voted on (10, 11); the deterministic first candidate
			// (11, 12) is fresh for her, so no re-roll is consumed.
			name:       "ranking_pair_voter_alice",
			target:     v2RankingPair + "?voter=" + addrAlice,
			template:   v2RankingPair,
			pathParams: map[string]string{},
		},
		{
			name:       "ranking_pair_voter_unindexed",
			target:     v2RankingPair + "?voter=0x9900000000000000000000000000000000000099",
			template:   v2RankingPair,
			pathParams: map[string]string{},
		},
		{
			name:       "ranking_pair_error_voter",
			target:     v2RankingPair + "?voter=not-a-wallet",
			template:   v2RankingPair,
			pathParams: map[string]string{},
		},
		{
			name:       "ranking_ratings",
			target:     v2RankingRatings,
			template:   v2RankingRatings,
			pathParams: map[string]string{},
		},
		{
			name:       "ranking_ratings_first_page",
			target:     v2RankingRatings + "?limit=2",
			template:   v2RankingRatings,
			pathParams: map[string]string{},
		},
		{
			name:       "ranking_ratings_next_page",
			target:     v2RankingRatings + "?limit=2&cursor=" + *page.Meta.NextCursor,
			template:   v2RankingRatings,
			pathParams: map[string]string{},
		},
		{
			name:       "ranking_ratings_error_limit",
			target:     v2RankingRatings + "?limit=0",
			template:   v2RankingRatings,
			pathParams: map[string]string{},
		},
		{
			name:       "ranking_ratings_error_cursor",
			target:     v2RankingRatings + "?cursor=not-a-cursor",
			template:   v2RankingRatings,
			pathParams: map[string]string{},
		},
		{
			name:       "ranking_statistics",
			target:     v2RankingStatistics,
			template:   v2RankingStatistics,
			pathParams: map[string]string{},
		},
		{
			name:       "ranking_statistics_error_internal",
			target:     v2RankingStatistics,
			template:   v2RankingStatistics,
			pathParams: map[string]string{},
			ctx:        cancelledCtx,
		},
	}
	runV2GoldenCases(t, h, spec, cases)

	// The exhausted-voter re-roll: alice has voted on the only remaining
	// pair when the collection is narrowed... not reproducible against the
	// live fixture without mutations, so pairExhausted=true stays pinned by
	// the unit suite; here we prove the voter filter accepts a wallet that
	// exhausts nothing.
	fresh := h.get(t, v2RankingPair+"?voter="+addrBob)
	var pair apiv2.RankingPair
	decodeV2JSON(t, fresh, &pair)
	if pair.PairExhausted {
		t.Fatalf("bob exhausted the fixture pairs: %+v", pair)
	}
}

// TestAPIV2RankingChallenge pins the challenge mutation: 201 with a fresh
// one-time nonce (redacted in the golden; the expiry is clock-pinned) that
// lands in the nonce table.
func TestAPIV2RankingChallenge(t *testing.T) {
	h := server(t)
	spec := v2RankingSpec(t)
	t.Cleanup(func() { restoreV2RankingState(t, h) })

	response := postV2(t, h, v2RankingChallenges, "", nil)
	if response.Code != http.StatusCreated {
		t.Fatalf("challenge status = %d; body=%s", response.Code, response.Body.String())
	}
	var challenge apiv2.RankingChallenge
	if err := json.Unmarshal(response.Body.Bytes(), &challenge); err != nil {
		t.Fatalf("decoding challenge: %v", err)
	}
	if len(challenge.Nonce) != 64 {
		t.Fatalf("nonce = %q, want 64 hex characters", challenge.Nonce)
	}
	// The expiry is issued by the database clock (NOW() + 15m), the same
	// clock the vote's consumption compares against; allow container skew.
	wantExpiry := time.Now().UTC().Add(beautyrank.ChallengeTTL)
	if drift := challenge.ExpiresAt.Sub(wantExpiry); drift < -time.Minute || drift > time.Minute {
		t.Fatalf("expiresAt = %v, want about %v", challenge.ExpiresAt, wantExpiry)
	}
	var stored int
	if err := h.db.QueryRow(
		`SELECT COUNT(*) FROM rw_ranking_vote_nonce WHERE nonce = $1`, challenge.Nonce,
	).Scan(&stored); err != nil {
		t.Fatalf("counting stored nonce: %v", err)
	}
	if stored != 1 {
		t.Fatalf("stored nonce rows = %d, want 1", stored)
	}

	second := postV2(t, h, v2RankingChallenges, "", nil)
	var next apiv2.RankingChallenge
	if err := json.Unmarshal(second.Body.Bytes(), &next); err != nil {
		t.Fatalf("decoding second challenge: %v", err)
	}
	if next.Nonce == challenge.Nonce {
		t.Fatal("consecutive challenges returned the same nonce")
	}

	// Both the nonce (random) and expiresAt (database NOW()) are
	// legitimately volatile; the golden pins the response shape.
	pinV2MutationGolden(t, spec, "ranking_challenge_created",
		v2RankingChallenges, v2RankingChallenges, "", nil,
		response, redactStringFields("nonce", "expiresAt"))
}

// TestAPIV2RankingVoteFlow drives the complete wallet flow through the real
// router: challenge → sign → 201 with both new ratings, then replayed
// nonces, duplicate pairs, tampered messages and validation problems.
func TestAPIV2RankingVoteFlow(t *testing.T) {
	h := server(t)
	spec := v2RankingSpec(t)
	restoreV2RankingState(t, h) // deterministic Elo inputs under -shuffle
	t.Cleanup(func() { restoreV2RankingState(t, h) })

	const chainID = int64(42161)
	fetchNonce := func() string {
		t.Helper()
		response := postV2(t, h, v2RankingChallenges, "", nil)
		if response.Code != http.StatusCreated {
			t.Fatalf("challenge status = %d; body=%s", response.Code, response.Body.String())
		}
		var challenge apiv2.RankingChallenge
		if err := json.Unmarshal(response.Body.Bytes(), &challenge); err != nil {
			t.Fatalf("decoding challenge: %v", err)
		}
		return challenge.Nonce
	}
	voteBody := func(nonce, signature string, first, second, winner int64) string {
		return fmt.Sprintf(
			`{"firstTokenId":%d,"secondTokenId":%d,"winnerTokenId":%d,"chainId":%d,"nonce":%q,"signature":%q}`,
			first, second, winner, chainID, nonce, signature,
		)
	}

	// Happy path: bob's fixture wallet has not voted on (12, 13); the vote
	// applies the exact Elo update over the fixture ratings (1195/1205 at
	// two recorded matches).
	nonce := fetchNonce()
	signature, signer := signV2RankingVote(t, chainID, nonce, 12, 13, 12)
	body := voteBody(nonce, signature, 12, 13, 12)
	created := postV2(t, h, v2RankingVotes, body, nil)
	if created.Code != http.StatusCreated {
		t.Fatalf("vote status = %d; body=%s", created.Code, created.Body.String())
	}
	var result apiv2.RankingVoteResult
	if err := json.Unmarshal(created.Body.Bytes(), &result); err != nil {
		t.Fatalf("decoding vote result: %v", err)
	}
	wantFirst, wantSecond := beautyrank.EloUpdate(1195, 1205, 1, 2)
	if result.FirstToken.Rating != wantFirst || result.SecondToken.Rating != wantSecond ||
		result.VoterAddress != signer || result.WinnerTokenId != 12 {
		t.Fatalf("vote result = %+v, want ratings (%v, %v) for %s", result, wantFirst, wantSecond, signer)
	}
	pinV2MutationGolden(t, spec, "ranking_vote_created",
		v2RankingVotes, v2RankingVotes, body, nil,
		created, redactStringFields("nonce"))

	// The vote is visible through the read resources.
	var statistics apiv2.RankingStatistics
	decodeV2JSON(t, h.get(t, v2RankingStatistics), &statistics)
	if statistics.TotalVotes != 3 || statistics.WalletVotes != 2 || statistics.DistinctVoters != 2 {
		t.Fatalf("statistics after vote = %+v", statistics)
	}
	var ratings apiv2.RankingRatingPage
	decodeV2JSON(t, h.get(t, v2RankingRatings), &ratings)
	moved := map[int64]float64{}
	for _, rating := range ratings.Data {
		moved[rating.TokenId] = rating.Rating
	}
	if moved[12] != wantFirst || moved[13] != wantSecond {
		t.Fatalf("ratings after vote = %v, want 12=%v 13=%v", moved, wantFirst, wantSecond)
	}

	// Replaying the consumed nonce is rejected before any write.
	replay := postV2(t, h, v2RankingVotes, body, nil)
	pinV2MutationGolden(t, spec, "ranking_vote_error_replayed_nonce",
		v2RankingVotes, v2RankingVotes, body, nil,
		replay, nil)
	if replay.Code != http.StatusBadRequest {
		t.Fatalf("replay status = %d", replay.Code)
	}

	// A fresh nonce cannot bypass the one-vote-per-pair constraint; the
	// rejected nonce survives for a retry on a different pair.
	dupNonce := fetchNonce()
	dupSignature, _ := signV2RankingVote(t, chainID, dupNonce, 13, 12, 13)
	dupBody := voteBody(dupNonce, dupSignature, 13, 12, 13)
	conflict := postV2(t, h, v2RankingVotes, dupBody, nil)
	if conflict.Code != http.StatusConflict {
		t.Fatalf("duplicate status = %d; body=%s", conflict.Code, conflict.Body.String())
	}
	pinV2MutationGolden(t, spec, "ranking_vote_error_already_voted",
		v2RankingVotes, v2RankingVotes, dupBody, nil,
		conflict, nil)
	var nonceAlive int
	if err := h.db.QueryRow(
		`SELECT COUNT(*) FROM rw_ranking_vote_nonce WHERE nonce = $1`, dupNonce,
	).Scan(&nonceAlive); err != nil {
		t.Fatalf("counting surviving nonce: %v", err)
	}
	if nonceAlive != 1 {
		t.Fatal("conflicting vote consumed its nonce; the wallet cannot retry")
	}

	// Tampering with the winner after signing recovers a different wallet:
	// the vote must not be attributed to the original signer.
	tamperNonce := fetchNonce()
	tamperSignature, tamperSigner := signV2RankingVote(t, chainID, tamperNonce, 10, 13, 10)
	tampered := postV2(t, h, v2RankingVotes, voteBody(tamperNonce, tamperSignature, 10, 13, 13), nil)
	if tampered.Code != http.StatusCreated {
		t.Fatalf("tampered-winner status = %d; body=%s", tampered.Code, tampered.Body.String())
	}
	var tamperedResult apiv2.RankingVoteResult
	if err := json.Unmarshal(tampered.Body.Bytes(), &tamperedResult); err != nil {
		t.Fatalf("decoding tampered result: %v", err)
	}
	if tamperedResult.VoterAddress == tamperSigner {
		t.Fatalf("tampered vote attributed to the original signer %s", tamperSigner)
	}

	// Validation problems answer specific RFC 9457 kinds.
	problems := []struct {
		name string
		body string
		want int
	}{
		{"ranking_vote_error_invalid_pair", voteBody(fetchNonce(), signature, 12, 12, 12), http.StatusBadRequest},
		{"ranking_vote_error_chain", strings.Replace(voteBody(fetchNonce(), signature, 12, 13, 12), fmt.Sprintf(`"chainId":%d`, chainID), `"chainId":31337`, 1), http.StatusBadRequest},
		{"ranking_vote_error_signature", voteBody(fetchNonce(), "zz", 12, 13, 12), http.StatusBadRequest},
		{"ranking_vote_error_body", `{"firstTokenId": ["not-a-number"]}`, http.StatusBadRequest},
	}
	for _, problem := range problems {
		response := postV2(t, h, v2RankingVotes, problem.body, nil)
		if response.Code != problem.want {
			t.Fatalf("%s status = %d, want %d; body=%s", problem.name, response.Code, problem.want, response.Body.String())
		}
		pinV2MutationGolden(t, spec, problem.name,
			v2RankingVotes, v2RankingVotes, problem.body, nil,
			response, nil)
	}
}

// TestAPIV2RankingMatchAdmin drives the admin mutation: the full auth
// matrix (fail-closed 503 on a keyless deployment, 401, 201 with the key)
// and the Elo outcome over fixture state.
func TestAPIV2RankingMatchAdmin(t *testing.T) {
	h := server(t)
	spec := v2RankingSpec(t)
	restoreV2RankingState(t, h)
	t.Cleanup(func() { restoreV2RankingState(t, h) })

	matchBody := `{"firstTokenId":10,"secondTokenId":11,"winnerTokenId":11}`
	adminHeaders := map[string]string{"X-Ranking-Admin-Key": adminKey}

	unauthorized := postV2(t, h, v2RankingMatches, matchBody, nil)
	pinV2MutationGolden(t, spec, "ranking_match_error_unauthorized",
		v2RankingMatches, v2RankingMatches, matchBody, nil,
		unauthorized, nil)
	if unauthorized.Code != http.StatusUnauthorized {
		t.Fatalf("missing key status = %d", unauthorized.Code)
	}
	wrongKey := postV2(t, h, v2RankingMatches, matchBody, map[string]string{"X-Ranking-Admin-Key": "wrong"})
	if wrongKey.Code != http.StatusUnauthorized {
		t.Fatalf("wrong key status = %d", wrongKey.Code)
	}

	invalid := postV2(t, h, v2RankingMatches,
		`{"firstTokenId":10,"secondTokenId":11,"winnerTokenId":12}`, adminHeaders)
	pinV2MutationGolden(t, spec, "ranking_match_error_winner_outside_pair",
		v2RankingMatches, v2RankingMatches,
		`{"firstTokenId":10,"secondTokenId":11,"winnerTokenId":12}`, adminHeaders,
		invalid, nil)
	if invalid.Code != http.StatusBadRequest {
		t.Fatalf("invalid pair status = %d", invalid.Code)
	}

	created := postV2(t, h, v2RankingMatches, matchBody, adminHeaders)
	if created.Code != http.StatusCreated {
		t.Fatalf("match status = %d; body=%s", created.Code, created.Body.String())
	}
	var result apiv2.RankingMatchResult
	if err := json.Unmarshal(created.Body.Bytes(), &result); err != nil {
		t.Fatalf("decoding match result: %v", err)
	}
	wantFirst, wantSecond := beautyrank.EloUpdate(1210.5, 1189.5, 0, 2)
	if result.FirstToken.Rating != wantFirst || result.SecondToken.Rating != wantSecond ||
		result.WinnerTokenId != 11 {
		t.Fatalf("match result = %+v, want ratings (%v, %v)", result, wantFirst, wantSecond)
	}
	pinV2MutationGolden(t, spec, "ranking_match_recorded",
		v2RankingMatches, v2RankingMatches, matchBody, adminHeaders,
		created, nil)

	// Admin matches carry no voter of record.
	var statistics apiv2.RankingStatistics
	decodeV2JSON(t, h.get(t, v2RankingStatistics), &statistics)
	if statistics.TotalVotes != 3 || statistics.WalletVotes != 1 {
		t.Fatalf("statistics after admin match = %+v", statistics)
	}
}

// TestAPIV2RankingAdminFailsClosed builds a second router whose v2 server
// has no admin key configured (a misdeployed instance) and proves the
// match operation answers the spec-declared 503 problem.
func TestAPIV2RankingAdminFailsClosed(t *testing.T) {
	h := server(t)
	spec := v2RankingSpec(t)

	keyless, err := apiv2.NewServer(
		h.store,
		h.state,
		slog.New(slog.DiscardHandler),
		apiv2.WithClock(func() time.Time { return time.Unix(1767230000, 0) }),
		apiv2.WithRanking(apiv2.RankingConfig{
			AdminKeys: []common.AdminKey{
				{Name: "RANKING_ADMIN_KEY", Value: ""},
				{Name: "ADMIN_API_KEY", Value: ""},
			},
		}),
	)
	if err != nil {
		t.Fatalf("building keyless v2 server: %v", err)
	}
	router := routes.New(h.store, routes.Options{V2: keyless})

	matchBody := `{"firstTokenId":10,"secondTokenId":11,"winnerTokenId":11}`
	request := httptest.NewRequest(http.MethodPost, v2RankingMatches, strings.NewReader(matchBody))
	request.Header.Set("Content-Type", "application/json")
	request.Header.Set("X-Ranking-Admin-Key", adminKey)
	request.RemoteAddr = "10.99.99.1:4242"
	recorder := httptest.NewRecorder()
	router.ServeHTTP(recorder, request)

	if recorder.Code != http.StatusServiceUnavailable {
		t.Fatalf("keyless status = %d; body=%s", recorder.Code, recorder.Body.String())
	}
	pinV2MutationGolden(t, spec, "ranking_match_error_admin_disabled",
		v2RankingMatches, v2RankingMatches, matchBody, nil,
		recorder, nil)
}

// TestAPIV2RankingWriteRateLimit exhausts the challenge bucket from one
// pinned client IP and pins the spec-declared 429 problem with its
// Retry-After header; a different client IP stays unaffected.
func TestAPIV2RankingWriteRateLimit(t *testing.T) {
	h := server(t)
	spec := v2RankingSpec(t)
	t.Cleanup(func() { restoreV2RankingState(t, h) })

	const limitedIP = "10.77.77.7:4242"
	limits := 0
	var limited *httptest.ResponseRecorder
	// The default challenge bucket is 2 rps with burst 20; drain it from
	// one address.
	for i := range 25 {
		response := h.do(t, request{
			method:     http.MethodPost,
			path:       v2RankingChallenges,
			remoteAddr: limitedIP,
		})
		switch response.Code {
		case http.StatusCreated:
			continue
		case http.StatusTooManyRequests:
			limits++
			limited = response
		default:
			t.Fatalf("challenge %d status = %d; body=%s", i, response.Code, response.Body.String())
		}
	}
	if limits == 0 {
		t.Fatal("25 challenges from one IP never hit the write rate limit")
	}
	if limited.Header().Get("Retry-After") != "1" {
		t.Fatalf("Retry-After = %q, want 1", limited.Header().Get("Retry-After"))
	}
	pinV2MutationGolden(t, spec, "ranking_challenge_error_rate_limited",
		v2RankingChallenges, v2RankingChallenges, "", nil,
		limited, nil)

	// A different client is not affected by the drained bucket.
	fresh := h.do(t, request{method: http.MethodPost, path: v2RankingChallenges})
	if fresh.Code != http.StatusCreated {
		t.Fatalf("fresh client status = %d; body=%s", fresh.Code, fresh.Body.String())
	}
}
