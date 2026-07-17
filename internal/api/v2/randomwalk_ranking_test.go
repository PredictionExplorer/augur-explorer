package v2

import (
	"bytes"
	"context"
	"encoding/base64"
	"encoding/hex"
	"errors"
	"fmt"
	"io"
	"log/slog"
	"math"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"

	"github.com/ethereum/go-ethereum/accounts"
	"github.com/ethereum/go-ethereum/crypto"

	"github.com/PredictionExplorer/augur-explorer/internal/api/common"
	"github.com/PredictionExplorer/augur-explorer/internal/api/cosmicgame/contractstate"
	"github.com/PredictionExplorer/augur-explorer/internal/api/httpx"
	"github.com/PredictionExplorer/augur-explorer/internal/beautyrank"
	rwmodel "github.com/PredictionExplorer/augur-explorer/internal/model/randomwalk"
	"github.com/PredictionExplorer/augur-explorer/internal/store"
	rwstore "github.com/PredictionExplorer/augur-explorer/internal/store/randomwalk"
)

// rankingTestVoterKey is the deterministic secp256k1 key the signed-vote
// tests use; its address is 0x7E5F4552091A69125d5DfCb7b8C2659029395Bdf.
const rankingTestVoterKey = "0000000000000000000000000000000000000000000000000000000000000001"

type fakeRankingRepository struct {
	addrs        func(context.Context) (rwmodel.ContractAddresses, error)
	explore      func(context.Context, int64, int64, int) ([]int64, error)
	addressID    func(context.Context, string) (int64, error)
	hasVote      func(context.Context, int64, int64, int64) (bool, error)
	ratingsPage  func(context.Context, int64, *rwstore.RankingRatingPageCursor, int) ([]rwstore.RankingRatingRecord, bool, error)
	statistics   func(context.Context) (rwstore.RankingStatisticsRecord, error)
	countMatches func(context.Context) (int64, error)
	ratingPair   func(context.Context, int64, int64) (float64, float64, error)
	ensureVoter  func(context.Context, string) (int64, error)
	createNonce  func(context.Context, string, time.Duration) (time.Time, error)
	recordMatch  func(context.Context, int64, int64, bool, float64, float64) error
	recordVote   func(context.Context, string, int64, int64, bool, float64, float64, int64) error
}

// fakeNonceExpiry is what the fake store clock answers for challenge
// creation.
var fakeNonceExpiry = time.Unix(1767230900, 0).UTC()

func (f fakeRankingRepository) ContractAddrs(ctx context.Context) (rwmodel.ContractAddresses, error) {
	if f.addrs == nil {
		return rwmodel.ContractAddresses{
			MarketPlace:    "0x1200000000000000000000000000000000000012",
			RandomWalk:     "0x8000000000000000000000000000000000000008",
			MarketPlaceAid: 12,
			RandomWalkAid:  8,
		}, nil
	}
	return f.addrs(ctx)
}

func (f fakeRankingRepository) ExploreRandomTokenIDs(ctx context.Context, rwalkAid, maxID int64, limit int) ([]int64, error) {
	if f.explore == nil {
		out := make([]int64, 0, limit)
		for i := range limit {
			out = append(out, int64(10+i))
		}
		return out, nil
	}
	return f.explore(ctx, rwalkAid, maxID, limit)
}

func (f fakeRankingRepository) UserAddressID(ctx context.Context, address string) (int64, error) {
	if f.addressID == nil {
		return 0, store.ErrNotFound
	}
	return f.addressID(ctx, address)
}

func (f fakeRankingRepository) HasRankingVoteForVoterPair(ctx context.Context, voterAid, first, second int64) (bool, error) {
	if f.hasVote == nil {
		return false, nil
	}
	return f.hasVote(ctx, voterAid, first, second)
}

func (f fakeRankingRepository) RankingRatingsPage(
	ctx context.Context,
	rwalkAid int64,
	after *rwstore.RankingRatingPageCursor,
	limit int,
) ([]rwstore.RankingRatingRecord, bool, error) {
	if f.ratingsPage == nil {
		return []rwstore.RankingRatingRecord{}, false, nil
	}
	return f.ratingsPage(ctx, rwalkAid, after, limit)
}

func (f fakeRankingRepository) RankingStatistics(ctx context.Context) (rwstore.RankingStatisticsRecord, error) {
	if f.statistics == nil {
		return rwstore.RankingStatisticsRecord{}, nil
	}
	return f.statistics(ctx)
}

func (f fakeRankingRepository) CountRankingMatches(ctx context.Context) (int64, error) {
	if f.countMatches == nil {
		return 0, nil
	}
	return f.countMatches(ctx)
}

func (f fakeRankingRepository) RatingPair(ctx context.Context, first, second int64) (float64, float64, error) {
	if f.ratingPair == nil {
		return beautyrank.DefaultRating, beautyrank.DefaultRating, nil
	}
	return f.ratingPair(ctx, first, second)
}

func (f fakeRankingRepository) EnsureVoterAddress(ctx context.Context, addr string) (int64, error) {
	if f.ensureVoter == nil {
		return 77, nil
	}
	return f.ensureVoter(ctx, addr)
}

func (f fakeRankingRepository) CreateRankingVoteNonce(ctx context.Context, nonce string, ttl time.Duration) (time.Time, error) {
	if f.createNonce == nil {
		return fakeNonceExpiry, nil
	}
	return f.createNonce(ctx, nonce, ttl)
}

func (f fakeRankingRepository) RecordRankingMatch(ctx context.Context, first, second int64, firstWon bool, newFirst, newSecond float64) error {
	if f.recordMatch == nil {
		return nil
	}
	return f.recordMatch(ctx, first, second, firstWon, newFirst, newSecond)
}

func (f fakeRankingRepository) RecordSignedRankingVote(
	ctx context.Context,
	nonce string,
	first, second int64,
	firstWon bool,
	newFirst, newSecond float64,
	voterAid int64,
) error {
	if f.recordVote == nil {
		return nil
	}
	return f.recordVote(ctx, nonce, first, second, firstWon, newFirst, newSecond, voterAid)
}

// countingEntropy yields a fresh deterministic nonce per read so repeated
// challenge creations in one test do not collide.
type countingEntropy struct {
	calls int
}

func (c *countingEntropy) Read(p []byte) (int, error) {
	c.calls++
	for i := range p {
		p[i] = byte(c.calls & 0xff)
	}
	return len(p), nil
}

// generousRankingLimits keeps the write rate limiter out of the way for
// tests that exercise handler behavior rather than the limiter itself.
func generousRankingLimits() *RankingWriteLimits {
	return &RankingWriteLimits{
		Challenges: RateLimitSpec{PerSecond: 10_000, Burst: 10_000},
		Votes:      RateLimitSpec{PerSecond: 10_000, Burst: 10_000},
		Matches:    RateLimitSpec{PerSecond: 10_000, Burst: 10_000},
	}
}

func newRankingTestServer(t *testing.T, ranking rankingRepository, config RankingConfig) *Server {
	t.Helper()
	if config.WriteLimits == nil {
		config.WriteLimits = generousRankingLimits()
	}
	server, err := newServer(
		nil,
		fakeBidReader{},
		fakeRoundReader{},
		fakeCurrentRoundReader{},
		fakeRoundPrizeReader{},
		fakeRoundRaffleReader{},
		fakeRoundDonationReader{},
		fakeStatisticsReader{},
		fakeBiddingAnalyticsReader{},
		fakeContractAddressReader{},
		fakeParticipantReader{},
		fakeUserReader{},
		fakeUserHistoryReader{},
		fakeUserStakingReader{},
		fakeUserActivityReader{},
		fakeGlobalDirectoryReader{},
		fakeGlobalStakingReader{},
		fakeRandomWalkReader{},
		ranking,
		fakeContractState{},
		slog.New(slog.NewTextHandler(io.Discard, nil)),
	)
	if err != nil {
		t.Fatalf("newServer: %v", err)
	}
	server.rankingConfig = config
	server.now = func() time.Time { return time.Unix(1767230000, 0) }
	server.entropy = &countingEntropy{}
	return server
}

// servePost drives one POST with an optional JSON body and headers
// through the full generated router, middleware chain included.
func servePost(
	t *testing.T,
	server *Server,
	target, body string,
	headers map[string]string,
) *httptest.ResponseRecorder {
	t.Helper()
	router := httpx.NewRouter()
	server.RegisterRoutes(router)
	var reader io.Reader
	if body != "" {
		reader = strings.NewReader(body)
	}
	request := httptest.NewRequest(http.MethodPost, target, reader)
	if body != "" {
		request.Header.Set("Content-Type", "application/json")
	}
	for name, value := range headers {
		request.Header.Set(name, value)
	}
	response := httptest.NewRecorder()
	router.ServeHTTP(response, request)
	return response
}

func signRankingVote(t *testing.T, chainID int64, nonce string, first, second, winner int64) (signature, signer string) {
	t.Helper()
	key, err := crypto.HexToECDSA(rankingTestVoterKey)
	if err != nil {
		t.Fatalf("parse deterministic key: %v", err)
	}
	message := beautyrank.VoteMessage(chainID, nonce, first, second, winner)
	sig, err := crypto.Sign(accounts.TextHash([]byte(message)), key)
	if err != nil {
		t.Fatalf("sign vote: %v", err)
	}
	return hex.EncodeToString(sig), crypto.PubkeyToAddress(key.PublicKey).Hex()
}

func rankingVoteBody(chainID int64, nonce, signature string, first, second, winner int64) string {
	return fmt.Sprintf(
		`{"firstTokenId":%d,"secondTokenId":%d,"winnerTokenId":%d,"chainId":%d,"nonce":%q,"signature":%q}`,
		first, second, winner, chainID, nonce, signature,
	)
}

func assertProblemKind(t *testing.T, response *httptest.ResponseRecorder, status int, kind string) Problem {
	t.Helper()
	assertProblem(t, response, status)
	var problem Problem
	decodeResponse(t, response, &problem)
	if problem.Type != problemTypeBase+kind {
		t.Fatalf("problem type = %q, want %q; body=%s", problem.Type, problemTypeBase+kind, response.Body.String())
	}
	return problem
}

func TestGetRandomWalkRankingRandomTokens(t *testing.T) {
	t.Parallel()

	t.Run("default sample size is two", func(t *testing.T) {
		t.Parallel()
		var gotLimit int
		var gotMaxID int64
		server := newRankingTestServer(t, fakeRankingRepository{
			explore: func(_ context.Context, rwalkAid, maxID int64, limit int) ([]int64, error) {
				if rwalkAid != 8 {
					t.Errorf("rwalkAid = %d, want 8", rwalkAid)
				}
				gotLimit, gotMaxID = limit, maxID
				return []int64{4, 9}, nil
			},
		}, RankingConfig{})
		response := serve(t, server, "/api/v2/randomwalk/ranking/random-tokens")
		if response.Code != http.StatusOK {
			t.Fatalf("status = %d; body=%s", response.Code, response.Body.String())
		}
		if gotLimit != 2 || gotMaxID != defaultExploreMaxTokenID {
			t.Fatalf("explore called with limit=%d maxID=%d, want 2/%d", gotLimit, gotMaxID, defaultExploreMaxTokenID)
		}
		var result RankingRandomTokens
		decodeResponse(t, response, &result)
		if len(result.TokenIds) != 2 || result.TokenIds[0] != 4 || result.TokenIds[1] != 9 {
			t.Fatalf("tokenIds = %v", result.TokenIds)
		}
	})

	t.Run("configured bound and explicit sample size", func(t *testing.T) {
		t.Parallel()
		server := newRankingTestServer(t, fakeRankingRepository{
			explore: func(_ context.Context, _ int64, maxID int64, limit int) ([]int64, error) {
				if maxID != 42 || limit != 5 {
					t.Errorf("explore(maxID=%d, limit=%d), want (42, 5)", maxID, limit)
				}
				return []int64{1, 2, 3}, nil
			},
		}, RankingConfig{ExploreMaxTokenID: 42})
		response := serve(t, server, "/api/v2/randomwalk/ranking/random-tokens?sampleSize=5")
		if response.Code != http.StatusOK {
			t.Fatalf("status = %d; body=%s", response.Code, response.Body.String())
		}
		var result RankingRandomTokens
		decodeResponse(t, response, &result)
		if len(result.TokenIds) != 3 {
			t.Fatalf("tokenIds = %v", result.TokenIds)
		}
	})

	t.Run("empty collection is an empty list", func(t *testing.T) {
		t.Parallel()
		server := newRankingTestServer(t, fakeRankingRepository{
			explore: func(context.Context, int64, int64, int) ([]int64, error) {
				return nil, nil
			},
		}, RankingConfig{})
		response := serve(t, server, "/api/v2/randomwalk/ranking/random-tokens")
		if response.Code != http.StatusOK {
			t.Fatalf("status = %d", response.Code)
		}
		if !strings.Contains(response.Body.String(), `"tokenIds":[]`) {
			t.Fatalf("body = %s", response.Body.String())
		}
	})

	t.Run("sample size bounds", func(t *testing.T) {
		t.Parallel()
		server := newRankingTestServer(t, fakeRankingRepository{}, RankingConfig{})
		for _, size := range []string{"0", "101", "-3"} {
			response := serve(t, server, "/api/v2/randomwalk/ranking/random-tokens?sampleSize="+size)
			assertProblemKind(t, response, http.StatusBadRequest, "invalid-parameter")
		}
	})

	t.Run("failure families", func(t *testing.T) {
		t.Parallel()
		cases := []struct {
			name    string
			ranking fakeRankingRepository
		}{
			{"registry failure", fakeRankingRepository{
				addrs: func(context.Context) (rwmodel.ContractAddresses, error) {
					return rwmodel.ContractAddresses{}, errors.New("registry unreadable")
				},
			}},
			{"explore failure", fakeRankingRepository{
				explore: func(context.Context, int64, int64, int) ([]int64, error) {
					return nil, errors.New("query failed")
				},
			}},
			{"over-cardinality", fakeRankingRepository{
				explore: func(context.Context, int64, int64, int) ([]int64, error) {
					return []int64{1, 2, 3}, nil
				},
			}},
			{"negative id", fakeRankingRepository{
				explore: func(context.Context, int64, int64, int) ([]int64, error) {
					return []int64{-1, 2}, nil
				},
			}},
			{"duplicate id", fakeRankingRepository{
				explore: func(context.Context, int64, int64, int) ([]int64, error) {
					return []int64{2, 2}, nil
				},
			}},
		}
		for _, tc := range cases {
			t.Run(tc.name, func(t *testing.T) {
				t.Parallel()
				server := newRankingTestServer(t, tc.ranking, RankingConfig{})
				response := serve(t, server, "/api/v2/randomwalk/ranking/random-tokens")
				assertProblemKind(t, response, http.StatusInternalServerError, "internal")
				if strings.Contains(response.Body.String(), "query failed") ||
					strings.Contains(response.Body.String(), "registry unreadable") {
					t.Fatalf("internal detail leaked: %s", response.Body.String())
				}
			})
		}
	})
}

func TestGetRandomWalkRankingPair(t *testing.T) {
	t.Parallel()
	const voter = "0x7E5F4552091A69125d5DfCb7b8C2659029395Bdf"

	t.Run("no voter returns first pair", func(t *testing.T) {
		t.Parallel()
		calls := 0
		server := newRankingTestServer(t, fakeRankingRepository{
			explore: func(_ context.Context, _, _ int64, limit int) ([]int64, error) {
				calls++
				if limit != 2 {
					t.Errorf("limit = %d, want 2", limit)
				}
				return []int64{5, 9}, nil
			},
			hasVote: func(context.Context, int64, int64, int64) (bool, error) {
				t.Error("vote filter must not run without a voter")
				return false, nil
			},
		}, RankingConfig{})
		response := serve(t, server, "/api/v2/randomwalk/ranking/pair")
		if response.Code != http.StatusOK {
			t.Fatalf("status = %d; body=%s", response.Code, response.Body.String())
		}
		var pair RankingPair
		decodeResponse(t, response, &pair)
		if pair.FirstTokenId != 5 || pair.SecondTokenId != 9 || pair.PairExhausted {
			t.Fatalf("pair = %+v", pair)
		}
		if calls != 1 {
			t.Fatalf("explore calls = %d, want 1", calls)
		}
	})

	t.Run("invalid voter is a 400", func(t *testing.T) {
		t.Parallel()
		server := newRankingTestServer(t, fakeRankingRepository{}, RankingConfig{})
		for _, bad := range []string{"nonsense", "0x123", "7E5F4552091A69125d5DfCb7b8C2659029395Bdf"} {
			response := serve(t, server, "/api/v2/randomwalk/ranking/pair?voter="+bad)
			assertProblemKind(t, response, http.StatusBadRequest, "invalid-parameter")
		}
	})

	t.Run("unindexed voter skips the filter", func(t *testing.T) {
		t.Parallel()
		server := newRankingTestServer(t, fakeRankingRepository{
			hasVote: func(context.Context, int64, int64, int64) (bool, error) {
				t.Error("vote filter must not run for an unindexed voter")
				return false, nil
			},
		}, RankingConfig{})
		response := serve(t, server, "/api/v2/randomwalk/ranking/pair?voter="+voter)
		if response.Code != http.StatusOK {
			t.Fatalf("status = %d; body=%s", response.Code, response.Body.String())
		}
		var pair RankingPair
		decodeResponse(t, response, &pair)
		if pair.PairExhausted {
			t.Fatalf("pair = %+v", pair)
		}
	})

	t.Run("voter re-rolls to an unvoted pair", func(t *testing.T) {
		t.Parallel()
		pairs := [][]int64{{1, 2}, {3, 4}}
		fetch := 0
		server := newRankingTestServer(t, fakeRankingRepository{
			addressID: func(_ context.Context, address string) (int64, error) {
				if address != voter {
					t.Errorf("address = %q, want checksummed %q", address, voter)
				}
				return 44, nil
			},
			explore: func(context.Context, int64, int64, int) ([]int64, error) {
				ids := pairs[fetch%len(pairs)]
				fetch++
				return ids, nil
			},
			hasVote: func(_ context.Context, voterAid, first, second int64) (bool, error) {
				if voterAid != 44 {
					t.Errorf("voterAid = %d, want 44", voterAid)
				}
				return first == 1 && second == 2, nil
			},
		}, RankingConfig{})
		response := serve(t, server, "/api/v2/randomwalk/ranking/pair?voter="+voter)
		if response.Code != http.StatusOK {
			t.Fatalf("status = %d; body=%s", response.Code, response.Body.String())
		}
		var pair RankingPair
		decodeResponse(t, response, &pair)
		if pair.FirstTokenId != 3 || pair.SecondTokenId != 4 || pair.PairExhausted {
			t.Fatalf("pair = %+v", pair)
		}
	})

	t.Run("exhausted voter still gets the last pair", func(t *testing.T) {
		t.Parallel()
		fetches := 0
		server := newRankingTestServer(t, fakeRankingRepository{
			addressID: func(context.Context, string) (int64, error) { return 44, nil },
			explore: func(context.Context, int64, int64, int) ([]int64, error) {
				fetches++
				return []int64{6, 7}, nil
			},
			hasVote: func(context.Context, int64, int64, int64) (bool, error) {
				return true, nil
			},
		}, RankingConfig{})
		response := serve(t, server, "/api/v2/randomwalk/ranking/pair?voter="+voter)
		if response.Code != http.StatusOK {
			t.Fatalf("status = %d; body=%s", response.Code, response.Body.String())
		}
		var pair RankingPair
		decodeResponse(t, response, &pair)
		if pair.FirstTokenId != 6 || pair.SecondTokenId != 7 || !pair.PairExhausted {
			t.Fatalf("pair = %+v", pair)
		}
		if fetches != maxRankingPairAttempts {
			t.Fatalf("fetches = %d, want %d", fetches, maxRankingPairAttempts)
		}
	})

	t.Run("fewer than two tokens is a 404", func(t *testing.T) {
		t.Parallel()
		server := newRankingTestServer(t, fakeRankingRepository{
			explore: func(context.Context, int64, int64, int) ([]int64, error) {
				return []int64{3}, nil
			},
		}, RankingConfig{})
		response := serve(t, server, "/api/v2/randomwalk/ranking/pair")
		assertProblemKind(t, response, http.StatusNotFound, "ranking-pair-unavailable")
	})

	t.Run("failure families", func(t *testing.T) {
		t.Parallel()
		cases := []struct {
			name    string
			ranking fakeRankingRepository
			voter   string
		}{
			{"explore failure", fakeRankingRepository{
				explore: func(context.Context, int64, int64, int) ([]int64, error) {
					return nil, errors.New("query failed")
				},
			}, ""},
			{"voter lookup failure", fakeRankingRepository{
				addressID: func(context.Context, string) (int64, error) {
					return 0, errors.New("lookup failed")
				},
			}, voter},
			{"vote filter failure", fakeRankingRepository{
				addressID: func(context.Context, string) (int64, error) { return 44, nil },
				hasVote: func(context.Context, int64, int64, int64) (bool, error) {
					return false, errors.New("filter failed")
				},
			}, voter},
			{"identical pair from repository", fakeRankingRepository{
				explore: func(context.Context, int64, int64, int) ([]int64, error) {
					return []int64{7, 7}, nil
				},
			}, ""},
		}
		for _, tc := range cases {
			t.Run(tc.name, func(t *testing.T) {
				t.Parallel()
				server := newRankingTestServer(t, tc.ranking, RankingConfig{})
				target := "/api/v2/randomwalk/ranking/pair"
				if tc.voter != "" {
					target += "?voter=" + tc.voter
				}
				response := serve(t, server, target)
				assertProblemKind(t, response, http.StatusInternalServerError, "internal")
				if strings.Contains(response.Body.String(), "failed") {
					t.Fatalf("internal detail leaked: %s", response.Body.String())
				}
			})
		}
	})
}

func rankingRatingRecordsForTest() []rwstore.RankingRatingRecord {
	return []rwstore.RankingRatingRecord{
		{TokenID: 12, Rating: 1075, MatchCount: 1},
		{TokenID: 10, Rating: 1200, MatchCount: 0},
		{TokenID: 11, Rating: 1200, MatchCount: 2},
	}
}

func TestListRandomWalkRankingRatings(t *testing.T) {
	t.Parallel()

	t.Run("page with continuation cursor", func(t *testing.T) {
		t.Parallel()
		var cursorSeen bool
		server := newRankingTestServer(t, fakeRankingRepository{
			ratingsPage: func(_ context.Context, rwalkAid int64, after *rwstore.RankingRatingPageCursor, limit int) ([]rwstore.RankingRatingRecord, bool, error) {
				if rwalkAid != 8 || limit != 2 {
					t.Errorf("page(rwalkAid=%d, limit=%d), want (8, 2)", rwalkAid, limit)
				}
				cursorSeen = after != nil
				return rankingRatingRecordsForTest()[:2], true, nil
			},
		}, RankingConfig{})
		response := serve(t, server, "/api/v2/randomwalk/ranking/ratings?limit=2")
		if response.Code != http.StatusOK {
			t.Fatalf("status = %d; body=%s", response.Code, response.Body.String())
		}
		if cursorSeen {
			t.Fatal("first page must not pass a cursor")
		}
		var page RankingRatingPage
		decodeResponse(t, response, &page)
		if len(page.Data) != 2 || page.Data[0].TokenId != 12 || page.Data[0].Rating != 1075 ||
			page.Data[0].MatchCount != 1 || page.Data[1].TokenId != 10 {
			t.Fatalf("page = %+v", page)
		}
		if page.Meta.Limit != 2 || page.Meta.NextCursor == nil {
			t.Fatalf("meta = %+v", page.Meta)
		}

		decoded, err := decodeRandomWalkRankingRatingCursor(*page.Meta.NextCursor)
		if err != nil {
			t.Fatalf("decode next cursor: %v", err)
		}
		if decoded.Rating != 1200 || decoded.TokenID != 10 {
			t.Fatalf("cursor = %+v", decoded)
		}
	})

	t.Run("continuation resumes at the cursor position", func(t *testing.T) {
		t.Parallel()
		encoded, err := encodeRandomWalkRankingRatingCursor(randomWalkRankingRatingCursor{
			Version:  randomWalkCursorVersion,
			Resource: randomWalkResourceRankingRatings,
			Rating:   1200,
			TokenID:  10,
		})
		if err != nil {
			t.Fatalf("encode cursor: %v", err)
		}
		server := newRankingTestServer(t, fakeRankingRepository{
			ratingsPage: func(_ context.Context, _ int64, after *rwstore.RankingRatingPageCursor, _ int) ([]rwstore.RankingRatingRecord, bool, error) {
				if after == nil || after.Rating != 1200 || after.TokenID != 10 {
					t.Errorf("cursor = %+v, want {1200 10}", after)
				}
				return rankingRatingRecordsForTest()[2:], false, nil
			},
		}, RankingConfig{})
		response := serve(t, server, "/api/v2/randomwalk/ranking/ratings?cursor="+encoded)
		if response.Code != http.StatusOK {
			t.Fatalf("status = %d; body=%s", response.Code, response.Body.String())
		}
		var page RankingRatingPage
		decodeResponse(t, response, &page)
		if len(page.Data) != 1 || page.Data[0].TokenId != 11 || page.Meta.NextCursor != nil {
			t.Fatalf("page = %+v", page)
		}
	})

	t.Run("empty page", func(t *testing.T) {
		t.Parallel()
		server := newRankingTestServer(t, fakeRankingRepository{}, RankingConfig{})
		response := serve(t, server, "/api/v2/randomwalk/ranking/ratings")
		if response.Code != http.StatusOK {
			t.Fatalf("status = %d", response.Code)
		}
		if !strings.Contains(response.Body.String(), `"data":[]`) {
			t.Fatalf("body = %s", response.Body.String())
		}
	})

	t.Run("invalid limits and cursors", func(t *testing.T) {
		t.Parallel()
		server := newRankingTestServer(t, fakeRankingRepository{}, RankingConfig{})
		for _, target := range []string{
			"/api/v2/randomwalk/ranking/ratings?limit=0",
			"/api/v2/randomwalk/ranking/ratings?limit=201",
		} {
			response := serve(t, server, target)
			assertProblemKind(t, response, http.StatusBadRequest, "invalid-parameter")
		}
		for _, cursor := range []string{"not-a-cursor", mustEncodeLedgerCursorForTest(t)} {
			response := serve(t, server, "/api/v2/randomwalk/ranking/ratings?cursor="+cursor)
			assertProblemKind(t, response, http.StatusBadRequest, "invalid-cursor")
		}
	})

	t.Run("failure families", func(t *testing.T) {
		t.Parallel()
		outOfOrder := rankingRatingRecordsForTest()
		outOfOrder[0], outOfOrder[2] = outOfOrder[2], outOfOrder[0]
		cases := []struct {
			name    string
			ranking fakeRankingRepository
		}{
			{"registry failure", fakeRankingRepository{
				addrs: func(context.Context) (rwmodel.ContractAddresses, error) {
					return rwmodel.ContractAddresses{}, errors.New("registry unreadable")
				},
			}},
			{"page failure", fakeRankingRepository{
				ratingsPage: func(context.Context, int64, *rwstore.RankingRatingPageCursor, int) ([]rwstore.RankingRatingRecord, bool, error) {
					return nil, false, errors.New("query failed")
				},
			}},
			{"over-cardinality", fakeRankingRepository{
				ratingsPage: func(_ context.Context, _ int64, _ *rwstore.RankingRatingPageCursor, limit int) ([]rwstore.RankingRatingRecord, bool, error) {
					records := make([]rwstore.RankingRatingRecord, limit+1)
					for i := range records {
						records[i] = rwstore.RankingRatingRecord{TokenID: int64(i), Rating: float64(1000 + i)}
					}
					return records, false, nil
				},
			}},
			{"out of order", fakeRankingRepository{
				ratingsPage: func(context.Context, int64, *rwstore.RankingRatingPageCursor, int) ([]rwstore.RankingRatingRecord, bool, error) {
					return outOfOrder, false, nil
				},
			}},
			{"duplicate keyset position", fakeRankingRepository{
				ratingsPage: func(context.Context, int64, *rwstore.RankingRatingPageCursor, int) ([]rwstore.RankingRatingRecord, bool, error) {
					return []rwstore.RankingRatingRecord{
						{TokenID: 10, Rating: 1200},
						{TokenID: 10, Rating: 1200},
					}, false, nil
				},
			}},
			{"negative match count", fakeRankingRepository{
				ratingsPage: func(context.Context, int64, *rwstore.RankingRatingPageCursor, int) ([]rwstore.RankingRatingRecord, bool, error) {
					return []rwstore.RankingRatingRecord{{TokenID: 10, Rating: 1200, MatchCount: -1}}, false, nil
				},
			}},
			{"negative token id", fakeRankingRepository{
				ratingsPage: func(context.Context, int64, *rwstore.RankingRatingPageCursor, int) ([]rwstore.RankingRatingRecord, bool, error) {
					return []rwstore.RankingRatingRecord{{TokenID: -1, Rating: 1200}}, false, nil
				},
			}},
			{"non-finite rating", fakeRankingRepository{
				ratingsPage: func(context.Context, int64, *rwstore.RankingRatingPageCursor, int) ([]rwstore.RankingRatingRecord, bool, error) {
					return []rwstore.RankingRatingRecord{{TokenID: 10, Rating: math.NaN()}}, false, nil
				},
			}},
		}
		for _, tc := range cases {
			t.Run(tc.name, func(t *testing.T) {
				t.Parallel()
				server := newRankingTestServer(t, tc.ranking, RankingConfig{})
				response := serve(t, server, "/api/v2/randomwalk/ranking/ratings")
				assertProblemKind(t, response, http.StatusInternalServerError, "internal")
			})
		}
	})

	t.Run("rows behind the cursor fail loudly", func(t *testing.T) {
		t.Parallel()
		encoded, err := encodeRandomWalkRankingRatingCursor(randomWalkRankingRatingCursor{
			Version:  randomWalkCursorVersion,
			Resource: randomWalkResourceRankingRatings,
			Rating:   1200,
			TokenID:  10,
		})
		if err != nil {
			t.Fatalf("encode cursor: %v", err)
		}
		server := newRankingTestServer(t, fakeRankingRepository{
			ratingsPage: func(context.Context, int64, *rwstore.RankingRatingPageCursor, int) ([]rwstore.RankingRatingRecord, bool, error) {
				return []rwstore.RankingRatingRecord{{TokenID: 10, Rating: 1200}}, false, nil
			},
		}, RankingConfig{})
		response := serve(t, server, "/api/v2/randomwalk/ranking/ratings?cursor="+encoded)
		assertProblemKind(t, response, http.StatusInternalServerError, "internal")
	})
}

// mustEncodeLedgerCursorForTest builds a valid cursor of a different
// RandomWalk resource, proving cross-resource rejection.
func mustEncodeLedgerCursorForTest(t *testing.T) string {
	t.Helper()
	encoded, err := encodeRandomWalkLedgerCursor(randomWalkLedgerCursor{
		Version:    randomWalkCursorVersion,
		Resource:   randomWalkResourceWithdrawals,
		EventLogID: 10,
	})
	if err != nil {
		t.Fatalf("encode ledger cursor: %v", err)
	}
	return encoded
}

func TestGetRandomWalkRankingStatistics(t *testing.T) {
	t.Parallel()

	t.Run("happy snapshot", func(t *testing.T) {
		t.Parallel()
		server := newRankingTestServer(t, fakeRankingRepository{
			statistics: func(context.Context) (rwstore.RankingStatisticsRecord, error) {
				return rwstore.RankingStatisticsRecord{
					TotalVotes:     10,
					WalletVotes:    7,
					DistinctVoters: 3,
					RatedTokens:    12,
				}, nil
			},
		}, RankingConfig{})
		response := serve(t, server, "/api/v2/randomwalk/ranking/statistics")
		if response.Code != http.StatusOK {
			t.Fatalf("status = %d; body=%s", response.Code, response.Body.String())
		}
		var statistics RankingStatistics
		decodeResponse(t, response, &statistics)
		if statistics.TotalVotes != 10 || statistics.WalletVotes != 7 ||
			statistics.DistinctVoters != 3 || statistics.RatedTokens != 12 {
			t.Fatalf("statistics = %+v", statistics)
		}
	})

	t.Run("invariant violations become opaque 500s", func(t *testing.T) {
		t.Parallel()
		cases := map[string]rwstore.RankingStatisticsRecord{
			"negative counter":            {TotalVotes: -1},
			"wallet exceeds total":        {TotalVotes: 1, WalletVotes: 2, DistinctVoters: 1},
			"voters exceed wallet votes":  {TotalVotes: 5, WalletVotes: 2, DistinctVoters: 3},
			"negative rated tokens":       {TotalVotes: 5, WalletVotes: 0, DistinctVoters: 0, RatedTokens: -1},
			"wallet votes without voters": {TotalVotes: 5, WalletVotes: 2, DistinctVoters: 0},
		}
		for name, record := range cases {
			t.Run(name, func(t *testing.T) {
				t.Parallel()
				server := newRankingTestServer(t, fakeRankingRepository{
					statistics: func(context.Context) (rwstore.RankingStatisticsRecord, error) {
						return record, nil
					},
				}, RankingConfig{})
				response := serve(t, server, "/api/v2/randomwalk/ranking/statistics")
				assertProblemKind(t, response, http.StatusInternalServerError, "internal")
			})
		}
	})

	t.Run("store failure", func(t *testing.T) {
		t.Parallel()
		server := newRankingTestServer(t, fakeRankingRepository{
			statistics: func(context.Context) (rwstore.RankingStatisticsRecord, error) {
				return rwstore.RankingStatisticsRecord{}, errors.New("query failed")
			},
		}, RankingConfig{})
		response := serve(t, server, "/api/v2/randomwalk/ranking/statistics")
		assertProblemKind(t, response, http.StatusInternalServerError, "internal")
	})
}

func TestCreateRandomWalkRankingChallenge(t *testing.T) {
	t.Parallel()

	t.Run("issues a nonce with the store's expiry", func(t *testing.T) {
		t.Parallel()
		var storedNonce string
		var requestedTTL time.Duration
		server := newRankingTestServer(t, fakeRankingRepository{
			createNonce: func(_ context.Context, nonce string, ttl time.Duration) (time.Time, error) {
				storedNonce, requestedTTL = nonce, ttl
				return fakeNonceExpiry, nil
			},
		}, RankingConfig{})
		response := servePost(t, server, "/api/v2/randomwalk/ranking/challenges", "", nil)
		if response.Code != http.StatusCreated {
			t.Fatalf("status = %d; body=%s", response.Code, response.Body.String())
		}
		var challenge RankingChallenge
		decodeResponse(t, response, &challenge)
		if len(challenge.Nonce) != 64 {
			t.Fatalf("nonce = %q, want 64 hex characters", challenge.Nonce)
		}
		if _, err := hex.DecodeString(challenge.Nonce); err != nil {
			t.Fatalf("nonce is not hex: %v", err)
		}
		if challenge.Nonce != storedNonce {
			t.Fatalf("returned nonce %q differs from stored %q", challenge.Nonce, storedNonce)
		}
		if requestedTTL != beautyrank.ChallengeTTL {
			t.Fatalf("ttl = %v, want %v", requestedTTL, beautyrank.ChallengeTTL)
		}
		// The response reports the storage-issued expiry (the database
		// clock), never a process-clock derivation that could drift from
		// the consumption check.
		if !challenge.ExpiresAt.Equal(fakeNonceExpiry) {
			t.Fatalf("expiresAt = %v, want the store expiry %v", challenge.ExpiresAt, fakeNonceExpiry)
		}
	})

	t.Run("consecutive challenges differ", func(t *testing.T) {
		t.Parallel()
		server := newRankingTestServer(t, fakeRankingRepository{}, RankingConfig{})
		first := servePost(t, server, "/api/v2/randomwalk/ranking/challenges", "", nil)
		second := servePost(t, server, "/api/v2/randomwalk/ranking/challenges", "", nil)
		var a, b RankingChallenge
		decodeResponse(t, first, &a)
		decodeResponse(t, second, &b)
		if a.Nonce == b.Nonce {
			t.Fatalf("consecutive nonces collide: %q", a.Nonce)
		}
	})

	t.Run("entropy failure is an opaque 500", func(t *testing.T) {
		t.Parallel()
		server := newRankingTestServer(t, fakeRankingRepository{}, RankingConfig{})
		server.entropy = errorReaderForTest{}
		response := servePost(t, server, "/api/v2/randomwalk/ranking/challenges", "", nil)
		assertProblemKind(t, response, http.StatusInternalServerError, "internal")
	})

	t.Run("store failure is an opaque 500", func(t *testing.T) {
		t.Parallel()
		server := newRankingTestServer(t, fakeRankingRepository{
			createNonce: func(context.Context, string, time.Duration) (time.Time, error) {
				return time.Time{}, errors.New("insert failed")
			},
		}, RankingConfig{})
		response := servePost(t, server, "/api/v2/randomwalk/ranking/challenges", "", nil)
		assertProblemKind(t, response, http.StatusInternalServerError, "internal")
	})
}

type errorReaderForTest struct{}

func (errorReaderForTest) Read([]byte) (int, error) {
	return 0, errors.New("entropy unavailable")
}

func TestCreateRandomWalkRankingVote(t *testing.T) {
	t.Parallel()
	const (
		first   = int64(10)
		second  = int64(13)
		chainID = int64(42161)
		nonce   = "a-nonce"
	)

	t.Run("records the vote and returns both ratings", func(t *testing.T) {
		t.Parallel()
		signature, signer := signRankingVote(t, chainID, nonce, first, second, second)
		wantFirst, wantSecond := beautyrank.EloUpdate(1100, 1250, 0, 17)
		var recorded bool
		server := newRankingTestServer(t, fakeRankingRepository{
			countMatches: func(context.Context) (int64, error) { return 17, nil },
			ratingPair: func(_ context.Context, gotFirst, gotSecond int64) (float64, float64, error) {
				if gotFirst != first || gotSecond != second {
					t.Errorf("ratingPair(%d, %d), want (%d, %d)", gotFirst, gotSecond, first, second)
				}
				return 1100, 1250, nil
			},
			ensureVoter: func(_ context.Context, addr string) (int64, error) {
				if addr != signer {
					t.Errorf("voter = %q, want %q", addr, signer)
				}
				return 91, nil
			},
			recordVote: func(_ context.Context, gotNonce string, gotFirst, gotSecond int64, firstWon bool, newFirst, newSecond float64, voterAid int64) error {
				recorded = true
				if gotNonce != nonce || gotFirst != first || gotSecond != second || firstWon {
					t.Errorf("recordVote(%q, %d, %d, %v)", gotNonce, gotFirst, gotSecond, firstWon)
				}
				if newFirst != wantFirst || newSecond != wantSecond {
					t.Errorf("ratings = (%v, %v), want (%v, %v)", newFirst, newSecond, wantFirst, wantSecond)
				}
				if voterAid != 91 {
					t.Errorf("voterAid = %d, want 91", voterAid)
				}
				return nil
			},
		}, RankingConfig{VoteChainIDs: []int64{chainID}})

		body := rankingVoteBody(chainID, nonce, signature, first, second, second)
		response := servePost(t, server, "/api/v2/randomwalk/ranking/votes", body, nil)
		if response.Code != http.StatusCreated {
			t.Fatalf("status = %d; body=%s", response.Code, response.Body.String())
		}
		if !recorded {
			t.Fatal("vote was not recorded")
		}
		var result RankingVoteResult
		decodeResponse(t, response, &result)
		if result.FirstToken.TokenId != first || result.FirstToken.Rating != wantFirst ||
			result.SecondToken.TokenId != second || result.SecondToken.Rating != wantSecond ||
			result.WinnerTokenId != second || result.VoterAddress != signer {
			t.Fatalf("result = %+v", result)
		}
	})

	t.Run("validation problems", func(t *testing.T) {
		t.Parallel()
		valid, _ := signRankingVote(t, chainID, nonce, first, second, first)
		cases := []struct {
			name   string
			body   string
			status int
			kind   string
		}{
			{"malformed json", "{", http.StatusBadRequest, "invalid-request"},
			{"negative token", rankingVoteBody(chainID, nonce, valid, -1, second, second), http.StatusBadRequest, "invalid-pair"},
			{"identical pair", rankingVoteBody(chainID, nonce, valid, first, first, first), http.StatusBadRequest, "invalid-pair"},
			{"winner outside pair", rankingVoteBody(chainID, nonce, valid, first, second, 999), http.StatusBadRequest, "invalid-pair"},
			{"missing nonce", rankingVoteBody(chainID, " ", valid, first, second, first), http.StatusBadRequest, "invalid-nonce"},
			{"missing signature", rankingVoteBody(chainID, nonce, " ", first, second, first), http.StatusBadRequest, "invalid-signature"},
			{"zero chain", rankingVoteBody(0, nonce, valid, first, second, first), http.StatusBadRequest, "chain-not-allowed"},
			{"chain outside allowlist", rankingVoteBody(1, nonce, valid, first, second, first), http.StatusBadRequest, "chain-not-allowed"},
			{"undecodable signature", rankingVoteBody(chainID, nonce, "zz", first, second, first), http.StatusBadRequest, "invalid-signature"},
		}
		for _, tc := range cases {
			t.Run(tc.name, func(t *testing.T) {
				t.Parallel()
				server := newRankingTestServer(t, fakeRankingRepository{
					recordVote: func(context.Context, string, int64, int64, bool, float64, float64, int64) error {
						t.Error("recordVote must not run for invalid requests")
						return nil
					},
				}, RankingConfig{VoteChainIDs: []int64{chainID}})
				response := servePost(t, server, "/api/v2/randomwalk/ranking/votes", tc.body, nil)
				assertProblemKind(t, response, tc.status, tc.kind)
			})
		}
	})

	t.Run("tampered winner does not attribute to the signer", func(t *testing.T) {
		t.Parallel()
		// Signed for token 7 winning against 9, submitted claiming 9 won:
		// the recovered address differs from the intended signer, so the
		// vote must never be attributed to the original wallet. The
		// distinct nonce and pair prove every field is bound into the
		// signed message.
		const tamperNonce = "tamper-nonce"
		signature, signer := signRankingVote(t, chainID, tamperNonce, 7, 9, 7)
		server := newRankingTestServer(t, fakeRankingRepository{
			ensureVoter: func(_ context.Context, addr string) (int64, error) {
				if addr == signer {
					t.Errorf("tampered message attributed to the original signer %q", addr)
				}
				return 55, nil
			},
		}, RankingConfig{VoteChainIDs: []int64{chainID}})
		body := rankingVoteBody(chainID, tamperNonce, signature, 7, 9, 9)
		response := servePost(t, server, "/api/v2/randomwalk/ranking/votes", body, nil)
		if response.Code != http.StatusCreated {
			t.Fatalf("status = %d; body=%s", response.Code, response.Body.String())
		}
		var result RankingVoteResult
		decodeResponse(t, response, &result)
		if result.VoterAddress == signer {
			t.Fatalf("tampered vote attributed to original signer: %+v", result)
		}
	})

	t.Run("store outcomes", func(t *testing.T) {
		t.Parallel()
		// A different allowlisted chain also proves the chain id is bound
		// into the signed message rather than fixed.
		const outcomeChain = int64(1)
		signature, _ := signRankingVote(t, outcomeChain, nonce, first, second, first)
		body := rankingVoteBody(outcomeChain, nonce, signature, first, second, first)
		cases := []struct {
			name    string
			ranking fakeRankingRepository
			status  int
			kind    string
		}{
			{"invalid nonce", fakeRankingRepository{
				recordVote: func(context.Context, string, int64, int64, bool, float64, float64, int64) error {
					return rwstore.ErrRankingNonceInvalid
				},
			}, http.StatusBadRequest, "invalid-nonce"},
			{"duplicate pair", fakeRankingRepository{
				recordVote: func(context.Context, string, int64, int64, bool, float64, float64, int64) error {
					return fmt.Errorf("record signed ranking vote: %w", store.ErrConflict)
				},
			}, http.StatusConflict, "already-voted"},
			{"count failure", fakeRankingRepository{
				countMatches: func(context.Context) (int64, error) { return 0, errors.New("count failed") },
			}, http.StatusInternalServerError, "internal"},
			{"rating failure", fakeRankingRepository{
				ratingPair: func(context.Context, int64, int64) (float64, float64, error) {
					return 0, 0, errors.New("rating failed")
				},
			}, http.StatusInternalServerError, "internal"},
			{"voter resolution failure", fakeRankingRepository{
				ensureVoter: func(context.Context, string) (int64, error) {
					return 0, errors.New("address failed")
				},
			}, http.StatusInternalServerError, "internal"},
			{"write failure", fakeRankingRepository{
				recordVote: func(context.Context, string, int64, int64, bool, float64, float64, int64) error {
					return errors.New("write failed")
				},
			}, http.StatusInternalServerError, "internal"},
		}
		for _, tc := range cases {
			t.Run(tc.name, func(t *testing.T) {
				t.Parallel()
				server := newRankingTestServer(t, tc.ranking, RankingConfig{VoteChainIDs: []int64{outcomeChain}})
				response := servePost(t, server, "/api/v2/randomwalk/ranking/votes", body, nil)
				assertProblemKind(t, response, tc.status, tc.kind)
				if strings.Contains(response.Body.String(), "failed") {
					t.Fatalf("internal detail leaked: %s", response.Body.String())
				}
			})
		}
	})
}

func TestRecordRandomWalkRankingMatch(t *testing.T) {
	t.Parallel()
	adminConfig := RankingConfig{
		AdminKeys: []common.AdminKey{
			{Name: "RANKING_ADMIN_KEY", Value: "secret-key"},
			{Name: "ADMIN_API_KEY", Value: ""},
		},
	}
	adminHeaders := map[string]string{"X-Ranking-Admin-Key": "secret-key"}
	matchBody := `{"firstTokenId":10,"secondTokenId":13,"winnerTokenId":10}`

	t.Run("records the match with the admin key", func(t *testing.T) {
		t.Parallel()
		wantFirst, wantSecond := beautyrank.EloUpdate(beautyrank.DefaultRating, beautyrank.DefaultRating, 1, 0)
		var recorded bool
		server := newRankingTestServer(t, fakeRankingRepository{
			recordMatch: func(_ context.Context, first, second int64, firstWon bool, newFirst, newSecond float64) error {
				recorded = true
				if first != 10 || second != 13 || !firstWon || newFirst != wantFirst || newSecond != wantSecond {
					t.Errorf("recordMatch(%d, %d, %v, %v, %v)", first, second, firstWon, newFirst, newSecond)
				}
				return nil
			},
		}, adminConfig)
		response := servePost(t, server, "/api/v2/randomwalk/ranking/matches", matchBody, adminHeaders)
		if response.Code != http.StatusCreated {
			t.Fatalf("status = %d; body=%s", response.Code, response.Body.String())
		}
		if !recorded {
			t.Fatal("match was not recorded")
		}
		var result RankingMatchResult
		decodeResponse(t, response, &result)
		if result.FirstToken.Rating != wantFirst || result.SecondToken.Rating != wantSecond ||
			result.WinnerTokenId != 10 {
			t.Fatalf("result = %+v", result)
		}
	})

	t.Run("authentication fails closed", func(t *testing.T) {
		t.Parallel()
		refuse := fakeRankingRepository{
			recordMatch: func(context.Context, int64, int64, bool, float64, float64) error {
				t.Error("recordMatch must not run without authentication")
				return nil
			},
		}

		t.Run("no key configured answers 503", func(t *testing.T) {
			t.Parallel()
			server := newRankingTestServer(t, refuse, RankingConfig{
				AdminKeys: []common.AdminKey{
					{Name: "RANKING_ADMIN_KEY", Value: ""},
					{Name: "ADMIN_API_KEY", Value: "  "},
				},
			})
			response := servePost(t, server, "/api/v2/randomwalk/ranking/matches", matchBody, adminHeaders)
			problem := assertProblemKind(t, response, http.StatusServiceUnavailable, "admin-disabled")
			if problem.Detail == nil || !strings.Contains(*problem.Detail, "RANKING_ADMIN_KEY or ADMIN_API_KEY") {
				t.Fatalf("problem = %+v", problem)
			}
		})

		t.Run("zero config answers 503", func(t *testing.T) {
			t.Parallel()
			server := newRankingTestServer(t, refuse, RankingConfig{})
			response := servePost(t, server, "/api/v2/randomwalk/ranking/matches", matchBody, adminHeaders)
			assertProblemKind(t, response, http.StatusServiceUnavailable, "admin-disabled")
		})

		t.Run("missing header answers 401", func(t *testing.T) {
			t.Parallel()
			server := newRankingTestServer(t, refuse, adminConfig)
			response := servePost(t, server, "/api/v2/randomwalk/ranking/matches", matchBody, nil)
			assertProblemKind(t, response, http.StatusUnauthorized, "unauthorized")
		})

		t.Run("wrong header answers 401", func(t *testing.T) {
			t.Parallel()
			server := newRankingTestServer(t, refuse, adminConfig)
			response := servePost(t, server, "/api/v2/randomwalk/ranking/matches", matchBody,
				map[string]string{"X-Ranking-Admin-Key": "wrong"})
			assertProblemKind(t, response, http.StatusUnauthorized, "unauthorized")
		})

		t.Run("fallback admin key authenticates", func(t *testing.T) {
			t.Parallel()
			server := newRankingTestServer(t, fakeRankingRepository{}, RankingConfig{
				AdminKeys: []common.AdminKey{
					{Name: "RANKING_ADMIN_KEY", Value: ""},
					{Name: "ADMIN_API_KEY", Value: "fallback"},
				},
			})
			response := servePost(t, server, "/api/v2/randomwalk/ranking/matches", matchBody,
				map[string]string{"X-Ranking-Admin-Key": "fallback"})
			if response.Code != http.StatusCreated {
				t.Fatalf("status = %d; body=%s", response.Code, response.Body.String())
			}
		})
	})

	t.Run("read operations bypass admin auth", func(t *testing.T) {
		t.Parallel()
		server := newRankingTestServer(t, fakeRankingRepository{}, RankingConfig{})
		response := serve(t, server, "/api/v2/randomwalk/ranking/statistics")
		if response.Code != http.StatusOK {
			t.Fatalf("status = %d; body=%s", response.Code, response.Body.String())
		}
	})

	t.Run("validation and failure", func(t *testing.T) {
		t.Parallel()
		cases := []struct {
			name   string
			body   string
			status int
			kind   string
		}{
			{"malformed json", "not json", http.StatusBadRequest, "invalid-request"},
			{"identical pair", `{"firstTokenId":5,"secondTokenId":5,"winnerTokenId":5}`, http.StatusBadRequest, "invalid-pair"},
			{"negative token", `{"firstTokenId":-4,"secondTokenId":5,"winnerTokenId":5}`, http.StatusBadRequest, "invalid-pair"},
			{"winner outside pair", `{"firstTokenId":4,"secondTokenId":5,"winnerTokenId":6}`, http.StatusBadRequest, "invalid-pair"},
		}
		for _, tc := range cases {
			t.Run(tc.name, func(t *testing.T) {
				t.Parallel()
				server := newRankingTestServer(t, fakeRankingRepository{}, adminConfig)
				response := servePost(t, server, "/api/v2/randomwalk/ranking/matches", tc.body, adminHeaders)
				assertProblemKind(t, response, tc.status, tc.kind)
			})
		}

		t.Run("write failure is an opaque 500", func(t *testing.T) {
			t.Parallel()
			server := newRankingTestServer(t, fakeRankingRepository{
				recordMatch: func(context.Context, int64, int64, bool, float64, float64) error {
					return errors.New("write failed")
				},
			}, adminConfig)
			response := servePost(t, server, "/api/v2/randomwalk/ranking/matches", matchBody, adminHeaders)
			assertProblemKind(t, response, http.StatusInternalServerError, "internal")
			if strings.Contains(response.Body.String(), "write failed") {
				t.Fatalf("internal detail leaked: %s", response.Body.String())
			}
		})

		t.Run("outcome computation failure is an opaque 500", func(t *testing.T) {
			t.Parallel()
			server := newRankingTestServer(t, fakeRankingRepository{
				countMatches: func(context.Context) (int64, error) {
					return 0, errors.New("count failed")
				},
				recordMatch: func(context.Context, int64, int64, bool, float64, float64) error {
					t.Error("recordMatch must not run when the outcome computation fails")
					return nil
				},
			}, adminConfig)
			response := servePost(t, server, "/api/v2/randomwalk/ranking/matches", matchBody, adminHeaders)
			assertProblemKind(t, response, http.StatusInternalServerError, "internal")
			if strings.Contains(response.Body.String(), "count failed") {
				t.Fatalf("internal detail leaked: %s", response.Body.String())
			}
		})
	})
}

// TestRankingVoteChainAllowlist pins the allowlist semantics: an empty
// configuration allows any positive chain id end to end.
func TestRankingVoteChainAllowlist(t *testing.T) {
	t.Parallel()

	t.Run("matrix", func(t *testing.T) {
		t.Parallel()
		cases := []struct {
			name    string
			allowed []int64
			chainID int64
			want    bool
		}{
			{"positive allowed when unset", nil, 7, true},
			{"positive allowed when empty", []int64{}, 42161, true},
			{"zero always rejected", nil, 0, false},
			{"negative always rejected", nil, -1, false},
			{"allowlisted", []int64{1, 42161}, 42161, true},
			{"not allowlisted", []int64{1, 42161}, 10, false},
		}
		for _, tc := range cases {
			config := RankingConfig{VoteChainIDs: tc.allowed}
			if got := config.voteChainAllowed(tc.chainID); got != tc.want {
				t.Errorf("%s: voteChainAllowed(%d) = %v, want %v", tc.name, tc.chainID, got, tc.want)
			}
		}
	})

	t.Run("empty allowlist accepts a vote on any positive chain", func(t *testing.T) {
		t.Parallel()
		const anyChain = int64(31337)
		nonce := "open-chain-nonce"
		signature, signer := signRankingVote(t, anyChain, nonce, 3, 4, 4)
		server := newRankingTestServer(t, fakeRankingRepository{
			ensureVoter: func(_ context.Context, addr string) (int64, error) {
				if addr != signer {
					t.Errorf("voter = %q, want %q", addr, signer)
				}
				return 12, nil
			},
		}, RankingConfig{})
		body := rankingVoteBody(anyChain, nonce, signature, 3, 4, 4)
		response := servePost(t, server, "/api/v2/randomwalk/ranking/votes", body, nil)
		if response.Code != http.StatusCreated {
			t.Fatalf("status = %d; body=%s", response.Code, response.Body.String())
		}
	})
}

// TestWithEntropyOption pins the constructor seam the challenge nonce
// randomness travels through.
func TestWithEntropyOption(t *testing.T) {
	t.Parallel()

	if _, err := NewServer(
		&store.Store{},
		&contractstate.State{},
		nil,
		WithEntropy(nil),
	); err == nil {
		t.Fatal("NewServer accepted a nil entropy source")
	}

	configured, err := NewServer(
		&store.Store{},
		&contractstate.State{},
		nil,
		WithEntropy(errorReaderForTest{}),
	)
	if err != nil {
		t.Fatalf("NewServer with entropy option: %v", err)
	}
	if _, err := configured.entropy.Read(nil); err == nil {
		t.Fatal("entropy option was not installed")
	}
}

func TestRankingWriteRateLimits(t *testing.T) {
	t.Parallel()

	t.Run("over-limit writes answer 429 problems", func(t *testing.T) {
		t.Parallel()
		server := newRankingTestServer(t, fakeRankingRepository{}, RankingConfig{
			WriteLimits: &RankingWriteLimits{
				Challenges: RateLimitSpec{PerSecond: 0.001, Burst: 1},
				Votes:      RateLimitSpec{PerSecond: 10_000, Burst: 10_000},
				Matches:    RateLimitSpec{PerSecond: 10_000, Burst: 10_000},
			},
		})
		router := httpx.NewRouter()
		server.RegisterRoutes(router)
		post := func() *httptest.ResponseRecorder {
			request := httptest.NewRequest(http.MethodPost, "/api/v2/randomwalk/ranking/challenges", nil)
			response := httptest.NewRecorder()
			router.ServeHTTP(response, request)
			return response
		}
		if first := post(); first.Code != http.StatusCreated {
			t.Fatalf("first status = %d; body=%s", first.Code, first.Body.String())
		}
		second := post()
		assertProblemKind(t, second, http.StatusTooManyRequests, "rate-limited")
		if second.Header().Get("Retry-After") != "1" {
			t.Fatalf("Retry-After = %q, want 1", second.Header().Get("Retry-After"))
		}
	})

	t.Run("limits are per operation", func(t *testing.T) {
		t.Parallel()
		server := newRankingTestServer(t, fakeRankingRepository{}, RankingConfig{
			AdminKeys: []common.AdminKey{{Name: "RANKING_ADMIN_KEY", Value: "k"}},
			WriteLimits: &RankingWriteLimits{
				Challenges: RateLimitSpec{PerSecond: 0.001, Burst: 1},
				Votes:      RateLimitSpec{PerSecond: 10_000, Burst: 10_000},
				Matches:    RateLimitSpec{PerSecond: 10_000, Burst: 10_000},
			},
		})
		router := httpx.NewRouter()
		server.RegisterRoutes(router)

		for range 2 {
			exhaust := httptest.NewRequest(http.MethodPost, "/api/v2/randomwalk/ranking/challenges", nil)
			router.ServeHTTP(httptest.NewRecorder(), exhaust)
		}
		// The match bucket is untouched by the exhausted challenge bucket.
		request := httptest.NewRequest(http.MethodPost, "/api/v2/randomwalk/ranking/matches",
			strings.NewReader(`{"firstTokenId":1,"secondTokenId":2,"winnerTokenId":1}`))
		request.Header.Set("Content-Type", "application/json")
		request.Header.Set("X-Ranking-Admin-Key", "k")
		response := httptest.NewRecorder()
		router.ServeHTTP(response, request)
		if response.Code != http.StatusCreated {
			t.Fatalf("status = %d; body=%s", response.Code, response.Body.String())
		}
	})

	t.Run("reads are not write-limited", func(t *testing.T) {
		t.Parallel()
		server := newRankingTestServer(t, fakeRankingRepository{}, RankingConfig{
			WriteLimits: &RankingWriteLimits{
				Challenges: RateLimitSpec{PerSecond: 0.001, Burst: 1},
				Votes:      RateLimitSpec{PerSecond: 0.001, Burst: 1},
				Matches:    RateLimitSpec{PerSecond: 0.001, Burst: 1},
			},
		})
		router := httpx.NewRouter()
		server.RegisterRoutes(router)
		for i := range 5 {
			request := httptest.NewRequest(http.MethodGet, "/api/v2/randomwalk/ranking/statistics", nil)
			response := httptest.NewRecorder()
			router.ServeHTTP(response, request)
			if response.Code != http.StatusOK {
				t.Fatalf("read %d status = %d", i, response.Code)
			}
		}
	})

	t.Run("default limits mirror v1", func(t *testing.T) {
		t.Parallel()
		limits := defaultRankingWriteLimits()
		if limits.Votes != (RateLimitSpec{PerSecond: 1, Burst: 10}) {
			t.Fatalf("votes = %+v", limits.Votes)
		}
		if limits.Matches != (RateLimitSpec{PerSecond: 2, Burst: 5}) {
			t.Fatalf("matches = %+v", limits.Matches)
		}
		if limits.Challenges.Burst < limits.Votes.Burst {
			t.Fatalf("challenge burst %d starves the vote burst %d", limits.Challenges.Burst, limits.Votes.Burst)
		}
		if (RankingConfig{}).writeLimits() != limits {
			t.Fatal("zero config does not apply the default limits")
		}
	})
}

func TestRankingRequestBodyDecodeProblem(t *testing.T) {
	t.Parallel()
	server := newRankingTestServer(t, fakeRankingRepository{}, RankingConfig{})
	response := servePost(t, server, "/api/v2/randomwalk/ranking/votes",
		`{"firstTokenId": "password-super-secret"}`, nil)
	problem := assertProblemKind(t, response, http.StatusBadRequest, "invalid-request")
	if problem.Detail == nil || !strings.Contains(*problem.Detail, "request body") {
		t.Fatalf("problem = %+v", problem)
	}
	if strings.Contains(response.Body.String(), "password-super-secret") {
		t.Fatalf("client payload leaked: %s", response.Body.String())
	}
}

func TestRankingRatingCursorCodec(t *testing.T) {
	t.Parallel()

	t.Run("round trip", func(t *testing.T) {
		t.Parallel()
		for _, rating := range []float64{-312.75, 0, 1075.0625, beautyrank.DefaultRating, 2412.3333333333335} {
			cursor := randomWalkRankingRatingCursor{
				Version:  randomWalkCursorVersion,
				Resource: randomWalkResourceRankingRatings,
				Rating:   rating,
				TokenID:  17,
			}
			encoded, err := encodeRandomWalkRankingRatingCursor(cursor)
			if err != nil {
				t.Fatalf("encode(%v): %v", rating, err)
			}
			decoded, err := decodeRandomWalkRankingRatingCursor(encoded)
			if err != nil {
				t.Fatalf("decode(%v): %v", rating, err)
			}
			if decoded != cursor {
				t.Fatalf("round trip = %+v, want %+v", decoded, cursor)
			}
		}
	})

	t.Run("rejections", func(t *testing.T) {
		t.Parallel()
		if _, err := encodeRandomWalkRankingRatingCursor(randomWalkRankingRatingCursor{
			Version:  randomWalkCursorVersion,
			Resource: randomWalkResourceRankingRatings,
			TokenID:  -1,
		}); err == nil {
			t.Fatal("encode accepted a negative token id")
		}
		if _, err := encodeRandomWalkRankingRatingCursor(randomWalkRankingRatingCursor{
			Version:  randomWalkCursorVersion,
			Resource: randomWalkResourceTokens,
			TokenID:  1,
		}); err == nil {
			t.Fatal("encode accepted a foreign resource")
		}
		// Syntactically well-formed payloads with hostile field values must
		// fail the semantic validation, not just the JSON decode.
		wellFormed := func(payload string) string {
			return base64.RawURLEncoding.EncodeToString([]byte(payload))
		}
		for name, encoded := range map[string]string{
			"garbage":           "###",
			"empty":             "",
			"wrong resource":    mustEncodeLedgerCursorForTest(t),
			"negative token id": wellFormed(`{"v":1,"k":"rwalkRankingRatings","r":1200,"t":-1}`),
			"wrong version":     wellFormed(`{"v":2,"k":"rwalkRankingRatings","r":1200,"t":1}`),
			"foreign resource":  wellFormed(`{"v":1,"k":"rwalkTokens","r":1200,"t":1}`),
			"missing resource":  wellFormed(`{"v":1,"r":1200,"t":1}`),
		} {
			if _, err := decodeRandomWalkRankingRatingCursor(encoded); err == nil {
				t.Fatalf("%s: decode accepted %q", name, encoded)
			}
		}
	})
}

// TestRankingVoteMessageCompatibility pins that the v2 vote flow signs the
// byte-identical v1 message, so a frontend can migrate without changing
// its wallet prompt.
func TestRankingVoteMessageCompatibility(t *testing.T) {
	t.Parallel()
	want := "RandomWalk beauty vote\nVersion: 1\nchainId: 42161\nnonce: n1\nnft1: 10\nnft2: 13\nwinner: 13"
	if got := beautyrank.VoteMessage(42161, "n1", 10, 13, 13); got != want {
		t.Fatalf("message = %q, want %q", got, want)
	}
}

// TestRankingChallengeNonceMatchesBeautyrank pins that the served nonce is
// exactly the beautyrank encoding of the entropy bytes.
func TestRankingChallengeNonceMatchesBeautyrank(t *testing.T) {
	t.Parallel()
	entropy := bytes.Repeat([]byte{0xAB}, 32)
	server := newRankingTestServer(t, fakeRankingRepository{}, RankingConfig{})
	server.entropy = bytes.NewReader(entropy)
	response := servePost(t, server, "/api/v2/randomwalk/ranking/challenges", "", nil)
	if response.Code != http.StatusCreated {
		t.Fatalf("status = %d; body=%s", response.Code, response.Body.String())
	}
	var challenge RankingChallenge
	decodeResponse(t, response, &challenge)
	if want := hex.EncodeToString(entropy); challenge.Nonce != want {
		t.Fatalf("nonce = %q, want %q", challenge.Nonce, want)
	}
}
