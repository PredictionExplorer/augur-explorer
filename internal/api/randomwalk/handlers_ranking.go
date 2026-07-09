package randomwalk

import (
	"context"
	"crypto/rand"
	"encoding/hex"
	"errors"
	"fmt"
	"math"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/ethereum/go-ethereum/accounts"
	ethcommon "github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgconn"

	"github.com/PredictionExplorer/augur-explorer/internal/api/common"
	rwdb "github.com/PredictionExplorer/augur-explorer/internal/store/randomwalk"
)

var errRankingBadPair = errors.New("randomwalk: nft1 and nft2 must be distinct non-negative token ids")

func performRankingMatch(ctx context.Context, nft1, nft2 int64, nft1Won bool) (raNew, rbNew float64, err error) {
	if nft1 < 0 || nft2 < 0 || nft1 == nft2 {
		return 0, 0, errRankingBadPair
	}
	n, err := rwRepo.CountRankingMatches(ctx)
	if err != nil {
		return 0, 0, err
	}
	ra, rb, err := rwRepo.RatingPair(ctx, nft1, nft2)
	if err != nil {
		return 0, 0, err
	}
	score := 0.0
	if nft1Won {
		score = 1.0
	}
	raNew, rbNew = computeEloUpdate(ra, rb, score, n)

	tx, err := rwStore.Pool().Begin(ctx)
	if err != nil {
		return 0, 0, err
	}
	defer func() { _ = tx.Rollback(ctx) }()

	if err := rwdb.ApplyRankingMatch(ctx, tx, nft1, nft2, nft1Won, raNew, rbNew, nil); err != nil {
		return 0, 0, err
	}
	if err := tx.Commit(ctx); err != nil {
		return 0, 0, err
	}
	return raNew, rbNew, nil
}

func exploreRandomMaxTokenID() int64 {
	maxID := int64(3766)
	if s := strings.TrimSpace(os.Getenv("RANDOMWALK_EXPLORE_MAX_TOKEN_ID")); s != "" {
		if v, err := strconv.ParseInt(s, 10, 64); err == nil && v > 0 {
			maxID = v
		}
	}
	return maxID
}

func fetchExploreRandomTokenIDs(ctx context.Context, limit int) ([]int64, error) {
	maxID := exploreRandomMaxTokenID()
	addrs, err := rwRepo.ContractAddrs(ctx)
	if err != nil {
		return nil, err
	}
	ids, err := rwRepo.ExploreRandomTokenIDs(ctx, addrs.RandomWalkAid, maxID, limit)
	if err != nil {
		ids, err = rwRepo.FallbackRandomTokenIDs(ctx, addrs.RandomWalkAid, maxID, limit)
	}
	return ids, err
}

// GET /api/randomwalk/explore/random and GET /api/randomwalk/random — default 2 token IDs unless ?limit=.
// Optional query: limit=1..100 (homepage / legacy random_token uses a higher limit).
func apiRandomwalkExploreRandom(c *gin.Context) {
	if !dbInitialized() {
		common.RespondErrorJSON(c, "Database link wasn't configured")
		return
	}
	limit := 2
	if s := strings.TrimSpace(c.Query("limit")); s != "" {
		if v, err := strconv.Atoi(s); err == nil && v > 0 && v <= 100 {
			limit = v
		}
	}
	ids, err := fetchExploreRandomTokenIDs(c.Request.Context(), limit)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, ids)
}

// GET /api/randomwalk/ranking/beauty-pair-ids — two token ids for the beauty contest.
// Optional query voter=0x… re-rolls up to 50 times to avoid pairs this wallet already voted on.
// skip_pair_filter=1 ignores that (one random pair; vote may still 409 if already voted).
func apiRandomwalkRankingBeautyPairIDs(c *gin.Context) {
	if !dbInitialized() {
		common.RespondErrorJSON(c, "Database link wasn't configured")
		return
	}
	skipPairFilter := strings.TrimSpace(c.Query("skip_pair_filter")) == "1"
	voterAid := int64(0)
	if !skipPairFilter {
		v := strings.TrimSpace(c.Query("voter"))
		if v != "" && ethcommon.IsHexAddress(v) {
			aid, err := rwStore.LookupAddressID(c.Request.Context(), ethcommon.HexToAddress(v).Hex())
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
		for attempt := 0; attempt < maxAttempts; attempt++ {
			ids, err = fetchExploreRandomTokenIDs(c.Request.Context(), 2)
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
				return
			}
			if len(ids) < 2 {
				break
			}
			voted, err := rwRepo.HasRankingVoteForVoterPair(c.Request.Context(), voterAid, ids[0], ids[1])
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
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
		ids, err = fetchExploreRandomTokenIDs(c.Request.Context(), 2)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
	}

	c.JSON(http.StatusOK, gin.H{
		"token_ids":      ids,
		"pair_exhausted": pairExhausted,
	})
}

// GET /api/randomwalk/vote_count — total pairwise beauty votes (rw_ranking_match count).
// total pairwise comparisons (Python GameModel count == rows in rw_ranking_match).
func apiRandomwalkVoteCount(c *gin.Context) {
	if !dbInitialized() {
		common.RespondErrorJSON(c, "Database link wasn't configured")
		return
	}
	n, err := rwRepo.CountRankingMatches(c.Request.Context())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"total_count": n})
}

// GET /api/randomwalk/token-ranking/order and GET /api/randomwalk/rating_order — token ids by Elo ascending.
func apiRandomwalkTokenRankingOrder(c *gin.Context) {
	if !dbInitialized() {
		common.RespondErrorJSON(c, "Database link wasn't configured")
		return
	}
	addrs, err := rwRepo.ContractAddrs(c.Request.Context())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ids, err := rwRepo.RatingOrder(c.Request.Context(), addrs.RandomWalkAid)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
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
func apiRandomwalkTokenRankingMatch(c *gin.Context) {
	if !dbInitialized() {
		common.RespondErrorJSON(c, "Database link wasn't configured")
		return
	}
	var body rankingMatchBody
	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	raNew, rbNew, err := performRankingMatch(c.Request.Context(), body.Nft1, body.Nft2, body.Nft1Won)
	if err != nil {
		if errors.Is(err, errRankingBadPair) {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"nft1": body.Nft1, "nft2": body.Nft2, "rating_nft1": raNew, "rating_nft2": rbNew})
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

func beautyVoteSignMessage(chainID int64, nonce string, nft1, nft2, winner int64) string {
	return fmt.Sprintf(
		"RandomWalk beauty vote\nVersion: 1\nchainId: %d\nnonce: %s\nnft1: %d\nnft2: %d\nwinner: %d",
		chainID, nonce, nft1, nft2, winner,
	)
}

func recoverPersonalSignSigner(message string, sigHex string) (ethcommon.Address, error) {
	raw := strings.TrimSpace(sigHex)
	raw = strings.TrimPrefix(raw, "0x")
	sig, err := hex.DecodeString(raw)
	if err != nil {
		return ethcommon.Address{}, fmt.Errorf("decode signature: %w", err)
	}
	if len(sig) != 65 {
		return ethcommon.Address{}, errors.New("signature must be 65 bytes")
	}
	if sig[64] == 27 || sig[64] == 28 {
		sig[64] -= 27
	}
	h := accounts.TextHash([]byte(message))
	pub, err := crypto.SigToPub(h, sig)
	if err != nil {
		return ethcommon.Address{}, err
	}
	return crypto.PubkeyToAddress(*pub), nil
}

func rankingVoteChainAllowed(chainID int64) bool {
	if chainID <= 0 {
		return false
	}
	env := strings.TrimSpace(os.Getenv("RANKING_VOTE_CHAIN_IDS"))
	if env == "" {
		return true
	}
	for _, p := range strings.Split(env, ",") {
		p = strings.TrimSpace(p)
		if p == "" {
			continue
		}
		v, err := strconv.ParseInt(p, 10, 64)
		if err == nil && v == chainID {
			return true
		}
	}
	return false
}

func performSignedBeautyVote(ctx context.Context, nft1, nft2 int64, nft1Won bool, chainID int64, signNonce, signature string) error {
	if nft1 < 0 || nft2 < 0 || nft1 == nft2 {
		return errRankingBadPair
	}
	if strings.TrimSpace(signNonce) == "" || strings.TrimSpace(signature) == "" {
		return errors.New("randomwalk: sign_nonce and signature are required")
	}
	if !rankingVoteChainAllowed(chainID) {
		return errors.New("randomwalk: chain_id not allowed for beauty votes")
	}
	winner := nft2
	if nft1Won {
		winner = nft1
	}
	msg := beautyVoteSignMessage(chainID, signNonce, nft1, nft2, winner)
	signer, err := recoverPersonalSignSigner(msg, signature)
	if err != nil {
		return fmt.Errorf("randomwalk: invalid signature: %w", err)
	}

	n, err := rwRepo.CountRankingMatches(ctx)
	if err != nil {
		return err
	}
	ra, rb, err := rwRepo.RatingPair(ctx, nft1, nft2)
	if err != nil {
		return err
	}
	score := 0.0
	if nft1Won {
		score = 1.0
	}
	raNew, rbNew := computeEloUpdate(ra, rb, score, n)

	voterAid, err := rwStore.LookupOrCreateAddress(ctx, signer.Hex(), 0, 0)
	if err != nil {
		return fmt.Errorf("randomwalk: voter address: %w", err)
	}

	tx, err := rwStore.Pool().Begin(ctx)
	if err != nil {
		return err
	}
	defer func() { _ = tx.Rollback(ctx) }()

	ok, err := rwdb.ConsumeRankingVoteNonce(ctx, tx, signNonce)
	if err != nil {
		return err
	}
	if !ok {
		return errors.New("randomwalk: invalid or expired sign_nonce")
	}

	if err := rwdb.ApplyRankingMatch(ctx, tx, nft1, nft2, nft1Won, raNew, rbNew, &voterAid); err != nil {
		// unique_violation: this wallet already voted on the pair.
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) && pgErr.Code == "23505" {
			return errRankingDuplicateVoterPair
		}
		return err
	}
	return tx.Commit(ctx)
}

var errRankingDuplicateVoterPair = errors.New("randomwalk: already voted on this pair")

// GET /api/randomwalk/ranking/sign-challenge — issue one-time nonce for wallet-signed /api/randomwalk/add_game.
func apiRandomwalkRankingSignChallenge(c *gin.Context) {
	if !dbInitialized() {
		common.RespondErrorJSON(c, "Database link wasn't configured")
		return
	}
	var b [32]byte
	if _, err := rand.Read(b[:]); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	nonce := hex.EncodeToString(b[:])
	if err := rwRepo.InsertRankingVoteNonce(c.Request.Context(), nonce, 15*time.Minute); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"nonce": nonce})
}

// POST /api/randomwalk/add_game — beauty contest with EIP-191 wallet signature; stores voter_aid and enforces one vote per pair per wallet.
// Response shape matches actionResponseSchema: {"result":"success"}.
func apiRandomwalkAddGameLegacy(c *gin.Context) {
	if !dbInitialized() {
		common.RespondErrorJSON(c, "Database link wasn't configured")
		return
	}
	var body addGameLegacyBody
	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if body.Nft1Win != 0 && body.Nft1Win != 1 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "nft1_win must be 0 or 1"})
		return
	}
	nft1Won := body.Nft1Win != 0

	err := performSignedBeautyVote(c.Request.Context(), body.Nft1, body.Nft2, nft1Won, body.ChainID, body.SignNonce, body.Signature)
	if err != nil {
		if errors.Is(err, errRankingBadPair) {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		if errors.Is(err, errRankingDuplicateVoterPair) {
			c.JSON(http.StatusConflict, gin.H{"error": "already voted on this pair"})
			return
		}
		if strings.Contains(err.Error(), "sign_nonce") || strings.Contains(err.Error(), "signature") ||
			strings.Contains(err.Error(), "chain_id") {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"result": "success"})
}

func computeEloUpdate(ra, rb, score float64, totalMatches int64) (raNew, rbNew float64) {
	k := 250.0 - float64(totalMatches)*0.00525
	if k < 1 {
		k = 1
	}
	es := 1.0 / (1.0 + math.Pow(10, (rb-ra)/400))
	raNew = ra + k*(score-es)
	rbNew = rb - k*(score-es)
	return raNew, rbNew
}
