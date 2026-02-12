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
	r.GET("/api/rwalk/current_offers/:rwalk_addr/:market_addr/:order_by", apiRwalkCurrentOffers)
	r.GET("/api/rwalk/floor_price/:rwalk_addr/:market_addr", apiRwalkFloorPrice)
	r.GET("/api/rwalk/tokens/list/sequential/:rwalk_addr", apiRwalkTokenListSeq)
	r.GET("/api/rwalk/tokens/list/sequential/:rwalk_addr/:offset/:limit", apiRwalkTokenListSeq)
	r.GET("/api/rwalk/tokens/list/by_period/:rwalk_addr/:init_ts/:fin_ts", apiRwalkTokenListPeriod)
	r.GET("/api/rwalk/tokens/info/:rwalk_addr/:token_id", apiRwalkTokenInfo)
	r.GET("/api/rwalk/tokens/name_changes/:token_id", apiRwalkTokenNameHistory)
	r.GET("/api/rwalk/trading/history/:market_addr/:offset/:limit", apiRwalkTradingHistory)
	r.GET("/api/rwalk/trading/by_user/:user_aid/:offset/:limit", apiRwalkTradingHistoryByUser)
	r.GET("/api/rwalk/trading/sales/:market_addr/:offset/:limit", apiRwalkSaleHistory)
	r.GET("/api/rwalk/tokens/history/:token_id/:rwalk_addr/:offset/:limit", apiRwalkTokenHistory)
	r.GET("/api/rwalk/tokens/by_user/:user_aid", apiRwalkTokensByUser)
	r.GET("/api/rwalk/statistics/by_token/:rwalk_addr", apiRwalkTokenStats)
	r.GET("/api/rwalk/statistics/by_market/:market_addr", apiRwalkMarketStats)
	r.GET("/api/rwalk/statistics/trading_volume/:market_addr/:init_ts/:fin_ts/:interval_secs", apiRwalkTradingVolumeByPeriod)
	r.GET("/api/rwalk/statistics/mint_intervals/:rwalk_addr", apiRwalkMintIntervals)
	r.GET("/api/rwalk/statistics/floor_price/:market_addr/:rwalk_addr/:init_ts/:fin_ts/:interval_secs", apiRwalkFloorPriceOverTime)
	r.GET("/api/rwalk/statistics/withdrawal_chart/:rwalk_addr", apiRwalkWithdrawalChart)
	r.GET("/api/rwalk/user/info/:user_aid/:rwalk_addr", apiRwalkUserInfo)
	r.GET("/api/rwalk/top5tokens", apiRwalkTop5TradedTokens)
	r.GET("/api/rwalk/mint_report", apiRwalkMintReport)
}

// RegisterHTMLRoutes registers all RandomWalk HTML page routes
func RegisterHTMLRoutes(r *gin.Engine) {
	r.GET("/black/rwalk", rwalk_index_page)
	r.GET("/black/rwalk/", rwalk_index_page)
	r.GET("/black/rwalk/current_offers/:rwalk_addr/:market_addr/:order_by", rwalk_current_offers)
	r.GET("/black/rwalk/floor_price/:rwalk_addr/:market_addr", rwalk_floor_price)
	r.GET("/black/rwalk/tokens/list/sequential/:rwalk_addr", rwalk_token_list_seq)
	r.GET("/black/rwalk/tokens/list/sequential/:rwalk_addr/:offset/:limit", rwalk_token_list_seq)
	r.GET("/black/rwalk/tokens/list/by_period/:rwalk_addr/:init_ts/:fin_ts", rwalk_token_list_period)
	r.GET("/black/rwalk/tokens/history/:token_id/:rwalk_addr", rwalk_token_history)
	r.GET("/black/rwalk/tokens/history/:token_id/:rwalk_addr/:offest/:limit", rwalk_token_history)
	r.GET("/black/rwalk/tokens/name_changes/:token_id", rwalk_token_name_history)
	r.GET("/black/rwalk/tokens/info/:rwalk_addr/:token_id", rwalk_token_info)
	r.GET("/black/rwalk/tokens/by_user/:user_aid", rwalk_tokens_by_user)
	r.GET("/black/rwalk/trading/history/:market_addr/:offset/:limit", rwalk_trading_history)
	r.GET("/black/rwalk/trading/history/:market_addr", rwalk_trading_history)
	r.GET("/black/rwalk/trading/by_user/:user_aid", rwalk_trading_history_by_user)
	r.GET("/black/rwalk/trading/sales/:market_addr", rwalk_sale_history)
	r.GET("/black/rwalk/statistics/by_token/:rwalk_addr", rwalk_token_stats)
	r.GET("/black/rwalk/statistics/by_market/:market_addr", rwalk_market_stats)
	r.GET("/black/rwalk/statistics/trading_volume/:market_addr/:init_ts/:fin_ts/:interval_secs", rwalk_trading_volume_by_period)
	r.GET("/black/rwalk/statistics/top_users", rwalk_top_users)
	r.GET("/black/rwalk/statistics/mint_intervals/:rwalk_addr", rwalk_mint_intervals)
	r.GET("/black/rwalk/statistics/withdrawal_chart/:rwalk_addr", rwalk_withdrawal_chart)
	r.GET("/black/rwalk/statistics/floor_price/:market_addr/:rwalk_addr/:init_ts/:fin_ts/:interval_secs", rwalk_floor_price_over_time)
	r.GET("/black/rwalk/user/info/:user_aid/:rwalk_addr", rwalk_user_info)
	r.GET("/black/rwalk/download_mints/:rwalk_addr", rwalk_token_csv_export)
	r.GET("/black/rwalk/mint_report", rwalk_mint_report)
}

// Helper to check if database is initialized
func dbInitialized() bool {
	return common.Ctx != nil && common.Ctx.Db != nil
}

