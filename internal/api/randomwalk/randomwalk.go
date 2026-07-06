// Package randomwalk provides HTTP handlers for RandomWalk NFT functionality
package randomwalk

import (
	"github.com/gin-gonic/gin"

	rwdb "github.com/PredictionExplorer/augur-explorer/internal/store/randomwalk"
	"github.com/PredictionExplorer/augur-explorer/internal/api/common"
)

// Response format flags for handlers
const (
	JSON = true
	HTTP = false
)

// Package-level storage wrapper for randomwalk database operations
var rw_storagew rwdb.SQLStorageWrapper

// routesEnabled mirrors websrv env ENABLE_ROUTES_RANDOMWALK (default true). When false, no RandomWalk HTTP routes are registered.
var routesEnabled = true

// Init initializes the randomwalk package. enableRoutes controls whether RegisterAPIRoutes adds handlers (from ENABLE_ROUTES_RANDOMWALK).
func Init(enableRoutes bool) {
	routesEnabled = enableRoutes
	if common.Ctx != nil && common.Ctx.Db != nil {
		rw_storagew.S = common.Ctx.Db
		rw_storagew.S.Db_set_schema_name("public")
	}
}

// RegisterAPIRoutes registers all RandomWalk JSON API routes
func RegisterAPIRoutes(r *gin.Engine) {
	if !routesEnabled {
		return
	}
	r.GET("/api/randomwalk/current_offers/:order_by", apiRwalkCurrentOffers)
	r.GET("/api/randomwalk/floor_price", apiRwalkFloorPrice)
	r.GET("/api/randomwalk/tokens/list/sequential", apiRwalkTokenListSeq)
	r.GET("/api/randomwalk/tokens/list/by_period/:init_ts/:fin_ts", apiRwalkTokenListPeriod)
	r.GET("/api/randomwalk/tokens/info/:token_id", apiRwalkTokenInfo)
	// Same JSON as above; use this path when a reverse proxy only forwards /api/cosmicgame/* to the app.
	r.GET("/api/cosmicgame/randomwalk/tokens/info/:token_id", apiRwalkTokenInfo)
	r.GET("/api/randomwalk/tokens/name_changes/:token_id", apiRwalkTokenNameHistory)
	r.GET("/api/randomwalk/trading/history/:offset/:limit", apiRwalkTradingHistory)
	r.GET("/api/randomwalk/trading/by_user/:user_aid/:offset/:limit", apiRwalkTradingHistoryByUser)
	r.GET("/api/randomwalk/trading/sales/:offset/:limit", apiRwalkSaleHistory)
	r.GET("/api/randomwalk/tokens/history/:token_id/:offset/:limit", apiRwalkTokenHistory)
	r.GET("/api/randomwalk/tokens/by_user/:user_aid", apiRwalkTokensByUser)
	r.GET("/api/randomwalk/statistics/by_token", apiRwalkTokenStats)
	r.GET("/api/randomwalk/statistics/by_market", apiRwalkMarketStats)
	r.GET("/api/randomwalk/statistics/trading_volume/:init_ts/:fin_ts/:interval_secs", apiRwalkTradingVolumeByPeriod)
	r.GET("/api/randomwalk/statistics/mint_intervals", apiRwalkMintIntervals)
	r.GET("/api/randomwalk/statistics/floor_price/:init_ts/:fin_ts/:interval_secs", apiRwalkFloorPriceOverTime)
	r.GET("/api/randomwalk/statistics/withdrawal_chart", apiRwalkWithdrawalChart)
	r.GET("/api/randomwalk/user/info/:user_aid", apiRwalkUserInfo)
	r.GET("/api/randomwalk/top5tokens", apiRwalkTop5TradedTokens)
	r.GET("/api/randomwalk/mint_report", apiRwalkMintReport)
	r.GET("/api/randomwalk/contracts", apiRwalkContracts)

	// NFT metadata + explore (legacy Python backend parity)
	r.GET("/api/randomwalk/explore/random", apiRandomwalkExploreRandom)
	r.GET("/api/randomwalk/random", apiRandomwalkExploreRandom)
	r.GET("/api/randomwalk/token-ranking/order", apiRandomwalkTokenRankingOrder)
	r.GET("/api/randomwalk/rating_order", apiRandomwalkTokenRankingOrder)
	r.GET("/api/randomwalk/vote_count", apiRandomwalkVoteCount)
	r.GET("/api/randomwalk/ranking/sign-challenge", apiRandomwalkRankingSignChallenge)
	r.GET("/api/randomwalk/ranking/beauty-pair-ids", apiRandomwalkRankingBeautyPairIDs)
	// Direct match recording is admin-only (fails closed when no key is set).
	// Public voting goes through /add_game, which is wallet-signature-verified
	// and rate limited instead.
	r.POST("/api/randomwalk/token-ranking/match",
		common.RateLimit(2, 5),
		common.RequireAdminKey("X-Ranking-Admin-Key", "RANKING_ADMIN_KEY", "ADMIN_API_KEY"),
		apiRandomwalkTokenRankingMatch)
	r.POST("/api/randomwalk/add_game", common.RateLimit(1, 10), apiRandomwalkAddGameLegacy)
	r.GET("/api/randomwalk/metadata/:token_id", apiRandomwalkTokenMetadata)

	// NOTE: the bare ERC-721 tokenURI route GET /metadata/:token_id is registered
	// centrally in main.go with host-aware dispatch, because the same webserv
	// serves both RandomWalk and Cosmic Signature hosts (and the Cosmic Signature
	// contract's tokenURI base is https://nfts.cosmicsignature.com/metadata/).
}

// Helper to check if database is initialized
func dbInitialized() bool {
	return common.Ctx != nil && common.Ctx.Db != nil
}
