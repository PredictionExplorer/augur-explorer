package randomwalk

import (
	"math"
	"net/http"
	"os"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"

	rwdb "github.com/PredictionExplorer/augur-explorer/rwcg/dbs/randomwalk"
	"github.com/PredictionExplorer/augur-explorer/rwcg/websrv/api/common"
)

// GET /api/randomwalk/explore/random — legacy parity with Python GET /random (up to 2 token IDs).
func apiRandomwalkExploreRandom(c *gin.Context) {
	if !dbInitialized() {
		common.RespondErrorJSON(c, "Database link wasn't configured")
		return
	}
	maxID := int64(3766)
	if s := strings.TrimSpace(os.Getenv("RANDOMWALK_EXPLORE_MAX_TOKEN_ID")); s != "" {
		if v, err := strconv.ParseInt(s, 10, 64); err == nil && v > 0 {
			maxID = v
		}
	}
	addrs := rw_storagew.Get_randomwalk_contract_addresses()
	ids, err := rw_storagew.Get_explore_random_token_ids(addrs.RandomWalkAid, maxID, 2)
	if err != nil {
		ids, err = rw_storagew.Get_fallback_random_token_ids(addrs.RandomWalkAid, maxID, 2)
	}
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, ids)
}

// GET /api/randomwalk/token-ranking/order — legacy parity with GET /rating_order.
func apiRandomwalkTokenRankingOrder(c *gin.Context) {
	if !dbInitialized() {
		common.RespondErrorJSON(c, "Database link wasn't configured")
		return
	}
	addrs := rw_storagew.Get_randomwalk_contract_addresses()
	ids, err := rw_storagew.Get_rating_order(addrs.RandomWalkAid)
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

// POST /api/randomwalk/token-ranking/match — records pairwise result and updates Elo (protected by RANKING_ADMIN_KEY).
func apiRandomwalkTokenRankingMatch(c *gin.Context) {
	if !dbInitialized() {
		common.RespondErrorJSON(c, "Database link wasn't configured")
		return
	}
	key := strings.TrimSpace(os.Getenv("RANKING_ADMIN_KEY"))
	if key == "" || c.GetHeader("X-Ranking-Admin-Key") != key {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid or missing X-Ranking-Admin-Key"})
		return
	}
	var body rankingMatchBody
	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if body.Nft1 <= 0 || body.Nft2 <= 0 || body.Nft1 == body.Nft2 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "nft1 and nft2 must be positive and distinct"})
		return
	}
	n, err := rw_storagew.Count_ranking_matches()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ra, rb, err := rw_storagew.Get_rating_pair(body.Nft1, body.Nft2)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	score := 0.0
	if body.Nft1Won {
		score = 1.0
	}
	raNew, rbNew := computeEloUpdate(ra, rb, score, n)

	tx, err := rw_storagew.S.Db().Begin()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer func() { _ = tx.Rollback() }()

	if err := rwdb.Apply_ranking_match_tx(tx, body.Nft1, body.Nft2, body.Nft1Won, raNew, rbNew); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	if err := tx.Commit(); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"nft1": body.Nft1, "nft2": body.Nft2, "rating_nft1": raNew, "rating_nft2": rbNew})
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
