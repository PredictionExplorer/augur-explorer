package randomwalk

import (
	"context"
	"crypto/rand"
	"errors"
	"fmt"
	"net/http"
	"slices"
	"strconv"
	"strings"

	ethcommon "github.com/ethereum/go-ethereum/common"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"

	"github.com/PredictionExplorer/augur-explorer/internal/api/common"
	"github.com/PredictionExplorer/augur-explorer/internal/api/httpx"
	"github.com/PredictionExplorer/augur-explorer/internal/beautyrank"
	rwdb "github.com/PredictionExplorer/augur-explorer/internal/store/randomwalk"
)

var (
	errRankingBadPair                 = errors.New("randomwalk: nft1 and nft2 must be distinct non-negative token ids")
	errRankingDuplicateVoterPair      = errors.New("randomwalk: already voted on this pair")
	errRankingVoteCredentialsRequired = errors.New("randomwalk: sign_nonce and signature are required")
	errRankingVoteChainNotAllowed     = errors.New("randomwalk: chain_id not allowed for beauty votes")
	errRankingVoteInvalidSignature    = errors.New("randomwalk: invalid signature")
	errRankingVoteInvalidNonce        = errors.New("randomwalk: invalid or expired sign_nonce")
)

type rankingMatchDependencies struct {
	countRankingMatches func(context.Context) (int64, error)
	ratingPair          func(context.Context, int64, int64) (float64, float64, error)
	withTransaction     func(context.Context, func(pgx.Tx) error) error
	applyRankingMatch   func(context.Context, pgx.Tx, int64, int64, bool, float64, float64, *int64) error
}

type signedBeautyVoteDependencies struct {
	rankingMatchDependencies

	lookupOrCreateAddress   func(context.Context, string, int64, int64) (int64, error)
	consumeRankingVoteNonce func(context.Context, pgx.Tx, string) (bool, error)
	chainAllowed            func(int64) bool
}

func (a *API) productionRankingMatchDependencies() rankingMatchDependencies {
	return rankingMatchDependencies{
		countRankingMatches: func(ctx context.Context) (int64, error) {
			return a.repo.CountRankingMatches(ctx)
		},
		ratingPair: func(ctx context.Context, nft1, nft2 int64) (float64, float64, error) {
			return a.repo.RatingPair(ctx, nft1, nft2)
		},
		withTransaction:   a.withRankingTransaction,
		applyRankingMatch: rwdb.ApplyRankingMatch,
	}
}

func (a *API) productionSignedBeautyVoteDependencies() signedBeautyVoteDependencies {
	return signedBeautyVoteDependencies{
		rankingMatchDependencies: a.productionRankingMatchDependencies(),
		lookupOrCreateAddress: func(ctx context.Context, addr string, blockNum, txID int64) (int64, error) {
			return a.store.LookupOrCreateAddress(ctx, addr, blockNum, txID)
		},
		consumeRankingVoteNonce: rwdb.ConsumeRankingVoteNonce,
		chainAllowed:            a.rankingVoteChainAllowed,
	}
}

func (a *API) withRankingTransaction(ctx context.Context, fn func(pgx.Tx) error) error {
	tx, err := a.store.Pool().Begin(ctx)
	if err != nil {
		return err
	}
	defer func() { _ = tx.Rollback(ctx) }()

	if err := fn(tx); err != nil {
		return err
	}
	return tx.Commit(ctx)
}

func (a *API) performRankingMatch(ctx context.Context, nft1, nft2 int64, nft1Won bool) (raNew, rbNew float64, err error) {
	return performRankingMatchWithDependencies(ctx, nft1, nft2, nft1Won, a.productionRankingMatchDependencies())
}

func performRankingMatchWithDependencies(
	ctx context.Context,
	nft1, nft2 int64,
	nft1Won bool,
	deps rankingMatchDependencies,
) (raNew, rbNew float64, err error) {
	if nft1 < 0 || nft2 < 0 || nft1 == nft2 {
		return 0, 0, errRankingBadPair
	}
	n, err := deps.countRankingMatches(ctx)
	if err != nil {
		return 0, 0, err
	}
	ra, rb, err := deps.ratingPair(ctx, nft1, nft2)
	if err != nil {
		return 0, 0, err
	}
	score := 0.0
	if nft1Won {
		score = 1.0
	}
	raNew, rbNew = beautyrank.EloUpdate(ra, rb, score, n)

	err = deps.withTransaction(ctx, func(tx pgx.Tx) error {
		return deps.applyRankingMatch(ctx, tx, nft1, nft2, nft1Won, raNew, rbNew, nil)
	})
	if err != nil {
		return 0, 0, err
	}
	return raNew, rbNew, nil
}

func exploreRandomLimit(raw string) int {
	limit := 2
	if s := strings.TrimSpace(raw); s != "" {
		if v, err := strconv.Atoi(s); err == nil && v > 0 && v <= 100 {
			limit = v
		}
	}
	return limit
}

func (a *API) fetchExploreRandomTokenIDs(ctx context.Context, limit int) ([]int64, error) {
	maxID := a.exploreMaxTokenID
	addrs, err := a.repo.ContractAddrs(ctx)
	if err != nil {
		return nil, err
	}
	ids, err := a.repo.ExploreRandomTokenIDs(ctx, addrs.RandomWalkAid, maxID, limit)
	if err != nil {
		ids, err = a.repo.FallbackRandomTokenIDs(ctx, addrs.RandomWalkAid, maxID, limit)
	}
	return ids, err
}

// GET /api/randomwalk/explore/random and GET /api/randomwalk/random — default 2 token IDs unless ?limit=.
// Optional query: limit=1..100 (homepage / legacy random_token uses a higher limit).
func (a *API) handleExploreRandom(c *httpx.Context) {
	if !a.dbInitialized() {
		common.RespondErrorJSON(c, "Database link wasn't configured")
		return
	}
	limit := exploreRandomLimit(c.Query("limit"))
	ids, err := a.fetchExploreRandomTokenIDs(c.Request.Context(), limit)
	if err != nil {
		a.respondStoreError(c, err)
		return
	}
	c.JSON(http.StatusOK, ids)
}

// GET /api/randomwalk/ranking/beauty-pair-ids — two token ids for the beauty contest.
// Optional query voter=0x… re-rolls up to 50 times to avoid pairs this wallet already voted on.
// skip_pair_filter=1 ignores that (one random pair; vote may still 409 if already voted).
func (a *API) handleRankingBeautyPairIDs(c *httpx.Context) {
	if !a.dbInitialized() {
		common.RespondErrorJSON(c, "Database link wasn't configured")
		return
	}
	skipPairFilter := strings.TrimSpace(c.Query("skip_pair_filter")) == "1"
	voterAid := int64(0)
	if !skipPairFilter {
		v := strings.TrimSpace(c.Query("voter"))
		if v != "" && ethcommon.IsHexAddress(v) {
			aid, err := a.store.LookupAddressID(c.Request.Context(), ethcommon.HexToAddress(v).Hex())
			if err == nil && aid > 0 {
				voterAid = aid
			}
		}
	}

	var ids []int64
	var err error
	pairExhausted := false
	const maxAttempts = 50

	if voterAid > 0 {
		found := false
		for range maxAttempts {
			ids, err = a.fetchExploreRandomTokenIDs(c.Request.Context(), 2)
			if err != nil {
				a.respondStoreError(c, err)
				return
			}
			if len(ids) < 2 {
				break
			}
			voted, err := a.repo.HasRankingVoteForVoterPair(c.Request.Context(), voterAid, ids[0], ids[1])
			if err != nil {
				a.respondStoreError(c, err)
				return
			}
			if !voted {
				found = true
				break
			}
		}
		if len(ids) >= 2 && !found {
			pairExhausted = true
		}
	} else {
		ids, err = a.fetchExploreRandomTokenIDs(c.Request.Context(), 2)
		if err != nil {
			a.respondStoreError(c, err)
			return
		}
	}

	c.JSON(http.StatusOK, httpx.H{
		"token_ids":      ids,
		"pair_exhausted": pairExhausted,
	})
}

// GET /api/randomwalk/vote_count — total pairwise beauty votes (rw_ranking_match count).
// total pairwise comparisons (Python GameModel count == rows in rw_ranking_match).
func (a *API) handleVoteCount(c *httpx.Context) {
	if !a.dbInitialized() {
		common.RespondErrorJSON(c, "Database link wasn't configured")
		return
	}
	n, err := a.repo.CountRankingMatches(c.Request.Context())
	if err != nil {
		a.respondStoreError(c, err)
		return
	}
	c.JSON(http.StatusOK, httpx.H{"total_count": n})
}

// GET /api/randomwalk/token-ranking/order and GET /api/randomwalk/rating_order — token ids by Elo ascending.
func (a *API) handleTokenRankingOrder(c *httpx.Context) {
	if !a.dbInitialized() {
		common.RespondErrorJSON(c, "Database link wasn't configured")
		return
	}
	addrs, err := a.repo.ContractAddrs(c.Request.Context())
	if err != nil {
		a.respondStoreError(c, err)
		return
	}
	ids, err := a.repo.RatingOrder(c.Request.Context(), addrs.RandomWalkAid)
	if err != nil {
		a.respondStoreError(c, err)
		return
	}
	c.JSON(http.StatusOK, ids)
}

type rankingMatchBody struct {
	Nft1    int64 `json:"nft1"`
	Nft2    int64 `json:"nft2"`
	Nft1Won bool  `json:"nft1_won"`
}

// POST /api/randomwalk/token-ranking/match — records pairwise result and updates Elo.
// Authentication is enforced by the RequireAdminKey middleware at registration
// (X-Ranking-Admin-Key, fails closed when no key is configured).
func (a *API) handleTokenRankingMatch(c *httpx.Context) {
	if !a.dbInitialized() {
		common.RespondErrorJSON(c, "Database link wasn't configured")
		return
	}
	var body rankingMatchBody
	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, httpx.H{"error": err.Error()})
		return
	}
	raNew, rbNew, err := a.performRankingMatch(c.Request.Context(), body.Nft1, body.Nft2, body.Nft1Won)
	if err != nil {
		if errors.Is(err, errRankingBadPair) {
			c.JSON(http.StatusBadRequest, httpx.H{"error": err.Error()})
			return
		}
		a.respondStoreError(c, err)
		return
	}
	c.JSON(http.StatusOK, httpx.H{"nft1": body.Nft1, "nft2": body.Nft2, "rating_nft1": raNew, "rating_nft2": rbNew})
}

// addGameLegacyBody matches POST /api/randomwalk/add_game (beauty contest frontend).
// Wallet path: sign_nonce + signature + chain_id (EIP-191 personal_sign over beautyVoteSignMessage).
type addGameLegacyBody struct {
	Nft1      int64  `json:"nft1"`
	Nft2      int64  `json:"nft2"`
	Nft1Win   int    `json:"nft1_win"`
	SignNonce string `json:"sign_nonce"`
	Signature string `json:"signature"`
	ChainID   int64  `json:"chain_id"`
}

// rankingVoteChainAllowed reports whether the wallet-signed vote's chain id
// passes the configured allowlist (empty allows any positive chain id).
func (a *API) rankingVoteChainAllowed(chainID int64) bool {
	if chainID <= 0 {
		return false
	}
	if len(a.voteChainIDs) == 0 {
		return true
	}
	return slices.Contains(a.voteChainIDs, chainID)
}

func (a *API) performSignedBeautyVote(ctx context.Context, nft1, nft2 int64, nft1Won bool, chainID int64, signNonce, signature string) error {
	return performSignedBeautyVoteWithDependencies(
		ctx,
		nft1,
		nft2,
		nft1Won,
		chainID,
		signNonce,
		signature,
		a.productionSignedBeautyVoteDependencies(),
	)
}

func performSignedBeautyVoteWithDependencies(
	ctx context.Context,
	nft1, nft2 int64,
	nft1Won bool,
	chainID int64,
	signNonce, signature string,
	deps signedBeautyVoteDependencies,
) error {
	if nft1 < 0 || nft2 < 0 || nft1 == nft2 {
		return errRankingBadPair
	}
	if strings.TrimSpace(signNonce) == "" || strings.TrimSpace(signature) == "" {
		return errRankingVoteCredentialsRequired
	}
	if !deps.chainAllowed(chainID) {
		return errRankingVoteChainNotAllowed
	}
	winner := nft2
	if nft1Won {
		winner = nft1
	}
	msg := beautyrank.VoteMessage(chainID, signNonce, nft1, nft2, winner)
	signer, err := beautyrank.RecoverSigner(msg, signature)
	if err != nil {
		return fmt.Errorf("%w: %w", errRankingVoteInvalidSignature, err)
	}

	n, err := deps.countRankingMatches(ctx)
	if err != nil {
		return err
	}
	ra, rb, err := deps.ratingPair(ctx, nft1, nft2)
	if err != nil {
		return err
	}
	score := 0.0
	if nft1Won {
		score = 1.0
	}
	raNew, rbNew := beautyrank.EloUpdate(ra, rb, score, n)

	voterAid, err := deps.lookupOrCreateAddress(ctx, signer.Hex(), 0, 0)
	if err != nil {
		return fmt.Errorf("randomwalk: voter address: %w", err)
	}

	return deps.withTransaction(ctx, func(tx pgx.Tx) error {
		ok, err := deps.consumeRankingVoteNonce(ctx, tx, signNonce)
		if err != nil {
			return err
		}
		if !ok {
			return errRankingVoteInvalidNonce
		}

		applyErr := deps.applyRankingMatch(ctx, tx, nft1, nft2, nft1Won, raNew, rbNew, &voterAid)
		return classifyRankingMatchApplyError(applyErr)
	})
}

func classifyRankingMatchApplyError(err error) error {
	if err == nil {
		return nil
	}
	var pgErr *pgconn.PgError
	if errors.As(err, &pgErr) && pgErr.Code == "23505" {
		return errRankingDuplicateVoterPair
	}
	return err
}

func (a *API) respondRankingVoteError(c *httpx.Context, err error) {
	switch {
	case errors.Is(err, errRankingBadPair):
		c.JSON(http.StatusBadRequest, httpx.H{"error": err.Error()})
	case errors.Is(err, errRankingDuplicateVoterPair):
		c.JSON(http.StatusConflict, httpx.H{"error": "already voted on this pair"})
	case errors.Is(err, errRankingVoteCredentialsRequired),
		errors.Is(err, errRankingVoteChainNotAllowed),
		errors.Is(err, errRankingVoteInvalidSignature),
		errors.Is(err, errRankingVoteInvalidNonce):
		c.JSON(http.StatusBadRequest, httpx.H{"error": err.Error()})
	default:
		a.respondStoreError(c, err)
	}
}

// GET /api/randomwalk/ranking/sign-challenge — issue one-time nonce for wallet-signed /api/randomwalk/add_game.
func (a *API) handleRankingSignChallenge(c *httpx.Context) {
	if !a.dbInitialized() {
		common.RespondErrorJSON(c, "Database link wasn't configured")
		return
	}
	nonce, err := beautyrank.NewNonce(rand.Reader)
	if err != nil {
		common.RespondInternalErrorJSON(c)
		return
	}
	if err := a.repo.InsertRankingVoteNonce(c.Request.Context(), nonce, beautyrank.ChallengeTTL); err != nil {
		a.respondStoreError(c, err)
		return
	}
	c.JSON(http.StatusOK, httpx.H{"nonce": nonce})
}

// POST /api/randomwalk/add_game — beauty contest with EIP-191 wallet signature; stores voter_aid and enforces one vote per pair per wallet.
// Response shape matches actionResponseSchema: {"result":"success"}.
func (a *API) handleAddGameLegacy(c *httpx.Context) {
	if !a.dbInitialized() {
		common.RespondErrorJSON(c, "Database link wasn't configured")
		return
	}
	var body addGameLegacyBody
	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, httpx.H{"error": err.Error()})
		return
	}
	if body.Nft1Win != 0 && body.Nft1Win != 1 {
		c.JSON(http.StatusBadRequest, httpx.H{"error": "nft1_win must be 0 or 1"})
		return
	}
	nft1Won := body.Nft1Win != 0

	err := a.performSignedBeautyVote(c.Request.Context(), body.Nft1, body.Nft2, nft1Won, body.ChainID, body.SignNonce, body.Signature)
	if err != nil {
		a.respondRankingVoteError(c, err)
		return
	}
	c.JSON(http.StatusOK, httpx.H{"result": "success"})
}
