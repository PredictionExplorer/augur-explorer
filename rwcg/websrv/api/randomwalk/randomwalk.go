// Package randomwalk provides HTTP handlers for RandomWalk NFT functionality
package randomwalk

import (
	"github.com/gin-gonic/gin"

	rwdb "github.com/PredictionExplorer/augur-explorer/rwcg/dbs/randomwalk"
	"github.com/PredictionExplorer/augur-explorer/rwcg/websrv/api/common"
)

// Response format flags for handlers
const (
	JSON = true
	HTTP = false
)

// Package-level storage wrapper for randomwalk database operations
var rw_storagew rwdb.SQLStorageWrapper

// Init initializes the randomwalk package with database connection
func Init() {
	if common.Ctx != nil && common.Ctx.Db != nil {
		rw_storagew.S = common.Ctx.Db
		rw_storagew.S.Db_set_schema_name("public")
	}
}

// RegisterAPIRoutes registers all RandomWalk JSON API routes
func RegisterAPIRoutes(r *gin.Engine) {
	r.GET("/api/randomwalk/current_offers/:order_by", apiRwalkCurrentOffers)
	r.GET("/api/randomwalk/floor_price", apiRwalkFloorPrice)
	r.GET("/api/randomwalk/tokens/list/sequential", apiRwalkTokenListSeq)
	r.GET("/api/randomwalk/tokens/list/by_period/:init_ts/:fin_ts", apiRwalkTokenListPeriod)
	r.GET("/api/randomwalk/tokens/info/:token_id", apiRwalkTokenInfo)
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
	r.GET("/api/randomwalk/token-ranking/order", apiRandomwalkTokenRankingOrder)
	r.GET("/api/randomwalk/vote_count", apiRandomwalkVoteCount)
	r.GET("/api/randomwalk/ranking/sign-challenge", apiRandomwalkRankingSignChallenge)
	r.GET("/api/randomwalk/ranking/beauty-pair-ids", apiRandomwalkRankingBeautyPairIDs)
	r.POST("/api/randomwalk/token-ranking/match", apiRandomwalkTokenRankingMatch)
	r.GET("/api/randomwalk/metadata/:token_id", apiRandomwalkTokenMetadata)

	// Legacy paths (same responses as old FastAPI app at host root)
	r.GET("/random", apiRandomwalkExploreRandom)
	r.GET("/rating_order", apiRandomwalkTokenRankingOrder)
	r.GET("/vote_count", apiRandomwalkVoteCount)
	r.POST("/add_game", apiRandomwalkAddGameLegacy)
	r.GET("/metadata/:token_id", apiRandomwalkTokenMetadata)
}

// RegisterHTMLRoutes registers all RandomWalk HTML page routes
func RegisterHTMLRoutes(r *gin.Engine) {
	r.GET("/black/randomwalk", rwalk_index_page)
	r.GET("/black/randomwalk/", rwalk_index_page)
	r.GET("/black/randomwalk/current_offers/:order_by", rwalk_current_offers)
	r.GET("/black/randomwalk/floor_price", rwalk_floor_price)
	r.GET("/black/randomwalk/tokens/list/sequential", rwalk_token_list_seq)
	r.GET("/black/randomwalk/tokens/list/by_period/:init_ts/:fin_ts", rwalk_token_list_period)
	r.GET("/black/randomwalk/tokens/history/:token_id/:offset/:limit", rwalk_token_history)
	r.GET("/black/randomwalk/tokens/history/:token_id", rwalk_token_history)
	r.GET("/black/randomwalk/tokens/name_changes/:token_id", rwalk_token_name_history)
	r.GET("/black/randomwalk/tokens/info/:token_id", rwalk_token_info)
	r.GET("/black/randomwalk/tokens/by_user/:user_aid", rwalk_tokens_by_user)
	r.GET("/black/randomwalk/trading/history/:offset/:limit", rwalk_trading_history)
	r.GET("/black/randomwalk/trading/history", rwalk_trading_history)
	r.GET("/black/randomwalk/trading/by_user/:user_aid", rwalk_trading_history_by_user)
	r.GET("/black/randomwalk/trading/sales", rwalk_sale_history)
	r.GET("/black/randomwalk/statistics/by_token", rwalk_token_stats)
	r.GET("/black/randomwalk/statistics/by_market", rwalk_market_stats)
	r.GET("/black/randomwalk/statistics/trading_volume/:init_ts/:fin_ts/:interval_secs", rwalk_trading_volume_by_period)
	r.GET("/black/randomwalk/statistics/top_users", rwalk_top_users)
	r.GET("/black/randomwalk/statistics/mint_intervals", rwalk_mint_intervals)
	r.GET("/black/randomwalk/statistics/withdrawal_chart", rwalk_withdrawal_chart)
	r.GET("/black/randomwalk/statistics/floor_price/:init_ts/:fin_ts/:interval_secs", rwalk_floor_price_over_time)
	r.GET("/black/randomwalk/user/info/:user_aid", rwalk_user_info)
	r.GET("/black/randomwalk/download_mints", rwalk_token_csv_export)
	r.GET("/black/randomwalk/mint_report", rwalk_mint_report)
}

// Helper to check if database is initialized
func dbInitialized() bool {
	return common.Ctx != nil && common.Ctx.Db != nil
}

