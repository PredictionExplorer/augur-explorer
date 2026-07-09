// Package cosmicgame provides HTTP handlers for CosmicGame functionality
package cosmicgame

import (
	"log"

	"github.com/ethereum/go-ethereum/ethclient"
	ethrpc "github.com/ethereum/go-ethereum/rpc"
	"github.com/gin-gonic/gin"

	"github.com/PredictionExplorer/augur-explorer/internal/api/common"
)

var (
	// EthClient for blockchain calls
	EthClient *ethclient.Client

	// RpcClient for direct RPC calls
	rpcclient *ethrpc.Client

	// Loggers
	Info  *log.Logger
	Error *log.Logger

	// Enabled is set from websrv env ENABLE_ROUTES_COSMICGAME (default true). When false, API routes are not registered.
	Enabled bool
)

// Init initializes the cosmicgame package with required dependencies.
// If enabled is false, RegisterAPIRoutes is a no-op; Init returns without loading contract state.
func Init(ethClient *ethclient.Client, rpcClient *ethrpc.Client, info, errorLog *log.Logger, enabled bool) {
	Enabled = enabled
	if !enabled {
		info.Printf("CosmicGame module init skipped (ENABLE_ROUTES_COSMICGAME=false)")
		return
	}

	EthClient = ethClient
	rpcclient = rpcClient
	Info = info
	Error = errorLog

	// Call the cosmic game initialization
	cosmic_game_init()
}

// Helper to check if database is initialized
func dbInitialized() bool {
	return common.Ctx != nil && common.Ctx.Store != nil
}

// RegisterAPIRoutes registers all CosmicGame JSON API routes.
// If ENABLE_ROUTES_COSMICGAME is false, this function returns without registering any routes.
func RegisterAPIRoutes(r *gin.Engine) {
	if !Enabled {
		return
	}
	// Statistics
	r.GET("/api/cosmicgame/statistics/dashboard", api_cosmic_game_dashboard)
	r.GET("/api/cosmicgame/statistics/counters", api_cosmic_game_record_counters)
	r.GET("/api/cosmicgame/statistics/unique/bidders", api_cosmic_game_user_unique_bidders)
	r.GET("/api/cosmicgame/statistics/unique/winners", api_cosmic_game_user_unique_winners)
	r.GET("/api/cosmicgame/statistics/leaderboard/roi", api_cosmic_game_roi_leaderboard)
	r.GET("/api/cosmicgame/statistics/claims/by_round", api_cosmic_game_claims_by_round)
	r.GET("/api/cosmicgame/statistics/claims/detail/:round_num", api_cosmic_game_claim_detail_by_round)
	r.GET("/api/cosmicgame/statistics/unique/donors", api_cosmic_game_user_unique_donors)
	r.GET("/api/cosmicgame/statistics/unique/stakers/cst", api_cosmic_game_user_unique_stakers_cst)
	r.GET("/api/cosmicgame/statistics/unique/stakers/randomwalk", api_cosmic_game_user_unique_stakers_rwalk)
	r.GET("/api/cosmicgame/statistics/unique/stakers/rwalk", api_cosmic_game_user_unique_stakers_rwalk) // legacy alias
	r.GET("/api/cosmicgame/statistics/unique/stakers/both", api_cosmic_game_user_unique_stakers_both)
	r.GET("/api/cosmicgame/statistics/bidding/activity/:init_ts/:fin_ts/:interval_secs", api_cosmic_game_bidding_activity)
	r.GET("/api/cosmicgame/statistics/bidding/frequency/:init_ts/:fin_ts/:interval_secs", api_cosmic_game_bidding_frequency)
	r.GET("/api/cosmicgame/statistics/bidding/top_active_periods/:n/:init_ts/:fin_ts", api_cosmic_game_bidding_top_active_periods)
	r.GET("/api/cosmicgame/statistics/bidding/time_bounds", api_cosmic_game_bidding_time_bounds)

	// Rounds
	r.GET("/api/cosmicgame/rounds/list/:offset/:limit", api_cosmic_game_prize_list)
	r.GET("/api/cosmicgame/rounds/info/:prize_num", api_cosmic_game_round_info)
	r.GET("/api/cosmicgame/rounds/current/time", api_cosmic_game_prize_cur_round_time)

	// Prizes
	r.GET("/api/cosmicgame/prizes/history/global/:offset/:limit", api_cosmic_game_global_claim_history_detail)
	r.GET("/api/cosmicgame/prizes/history/by_user/:user_addr/:offset/:limit", api_cosmic_game_prize_history_detail_by_user)
	r.GET("/api/cosmicgame/prizes/eth/all/global", api_cosmic_game_all_eth_deposits_list)
	r.GET("/api/cosmicgame/prizes/eth/all/global/:offset/:limit", api_cosmic_game_all_eth_deposits_list)
	r.GET("/api/cosmicgame/prizes/eth/raffle/global", api_cosmic_game_raffle_eth_deposits_list)
	r.GET("/api/cosmicgame/prizes/eth/raffle/global/:offset/:limit", api_cosmic_game_raffle_eth_deposits_list)
	r.GET("/api/cosmicgame/prizes/eth/chronowarrior/global", api_cosmic_game_chronowarrior_eth_deposits_list)
	r.GET("/api/cosmicgame/prizes/eth/chronowarrior/global/:offset/:limit", api_cosmic_game_chronowarrior_eth_deposits_list)
	r.GET("/api/cosmicgame/prizes/eth/all/by_user/:user_addr", api_cosmic_game_unified_eth_all_by_user)
	r.GET("/api/cosmicgame/prizes/eth/raffle/by_user/:user_addr", api_cosmic_game_unified_eth_raffle_by_user)
	r.GET("/api/cosmicgame/prizes/eth/chronowarrior/by_user/:user_addr", api_cosmic_game_unified_eth_chronowarrior_by_user)
	r.GET("/api/cosmicgame/prizes/eth/unclaimed/by_user/:user_addr/:offset/:limit", api_cosmic_game_unclaimed_prize_deposits_by_user)
	r.GET("/api/cosmicgame/prizes/deposits/raffle/by_user/:user_addr", api_cosmic_game_prize_deposits_raffle_eth_by_user)
	r.GET("/api/cosmicgame/prizes/deposits/chrono_warrior/by_user/:user_addr", api_cosmic_game_prize_deposits_chrono_warrior_by_user)
	r.GET("/api/cosmicgame/prizes/deposits/unclaimed/by_user/:user_addr/:offset/:limit", api_cosmic_game_unclaimed_prize_deposits_by_user)

	// Bids
	r.GET("/api/cosmicgame/bid/list/all/:offset/:limit", api_cosmic_game_bid_list)
	r.GET("/api/cosmicgame/bid/info/:evtlog_id", api_cosmic_game_bid_info)
	r.GET("/api/cosmicgame/bid/info_by_pos/:round_num/:bid_position", api_cosmic_game_bid_info_by_pos)
	r.GET("/api/cosmicgame/bid/with_message/by_round/:round", api_cosmic_game_bid_with_message_by_round)
	r.GET("/api/cosmicgame/bid/list/by_round/:round_num/:sort/:offset/:limit", api_cosmic_game_bid_list_by_round)
	r.GET("/api/cosmicgame/bid/bid_type_ratio", api_cosmic_game_bid_type_ratio)
	r.GET("/api/cosmicgame/bid/used_randomwalk_nfts", api_cosmic_game_used_rwalk_nfts)
	r.GET("/api/cosmicgame/bid/used_rwalk_nfts", api_cosmic_game_used_rwalk_nfts) // legacy path (same handler)
	r.GET("/api/cosmicgame/bid/cst_price", api_cosmic_game_get_cst_price)
	r.GET("/api/cosmicgame/bid/eth_price", api_cosmic_game_get_eth_price)
	r.GET("/api/cosmicgame/bid/current_special_winners", api_cosmic_game_bid_special_winners)
	r.GET("/api/cosmicgame/get_banned_bids", api_cosmic_game_get_banned_bids)
	// Bid moderation is admin-only: requires X-Admin-Key matching ADMIN_API_KEY
	// (503 when unset — fails closed) and is strictly rate limited.
	adminGuard := common.RequireAdminKey("X-Admin-Key", "ADMIN_API_KEY")
	adminRate := common.RateLimit(2, 5)
	r.POST("/api/cosmicgame/ban_bid", adminRate, adminGuard, api_cosmic_game_ban_bid)
	r.POST("/api/cosmicgame/unban_bid", adminRate, adminGuard, api_cosmic_game_unban_bid)

	// CST Tokens
	r.GET("/api/cosmicgame/cst/list/all/:offset/:limit", api_cosmic_game_cosmic_signature_token_list)
	r.GET("/api/cosmicgame/cst/list/by_user/:user_addr/:offset/:limit", api_cosmic_game_cosmic_signature_token_list_by_user)
	r.GET("/api/cosmicgame/cst/info/:token_id", api_cosmic_game_cosmic_signature_token_info)
	r.GET("/api/cosmicgame/cst/metadata/:token_id", api_cosmic_game_cst_metadata)
	r.GET("/cg/metadata/:token_id", api_cosmic_game_cst_metadata) // legacy Python path
	r.GET("/api/cosmicgame/cst/names/history/:token_id", api_cosmic_game_token_name_history)
	r.GET("/api/cosmicgame/cst/names/search/:name", api_cosmic_game_token_name_search)
	r.GET("/api/cosmicgame/cst/names/named_only", api_cosmic_game_named_tokens_only)
	r.GET("/api/cosmicgame/cst/transfers/all/:token_id/:offset/:limit", api_cosmic_game_token_ownership_transfers)
	r.GET("/api/cosmicgame/cst/transfers/by_user/:user_addr/:offset/:limit", api_cosmic_game_cosmic_signature_transfers_by_user)
	r.GET("/api/cosmicgame/cst/distribution", api_cosmic_game_cs_token_distribution)

	// Cosmic Token (CT)
	r.GET("/api/cosmicgame/ct/balances", api_cosmic_game_cosmic_token_balances)
	r.GET("/api/cosmicgame/ct/statistics", api_cosmic_game_cosmic_token_statistics)
	r.GET("/api/cosmicgame/ct/summary/by_user/:user_addr", api_cosmic_game_cosmic_token_summary_by_user)
	r.GET("/api/cosmicgame/ct/transfers/by_user/:user_addr/:offset/:limit", api_cosmic_game_cosmic_token_transfers_by_user)
	r.GET("/api/cosmicgame/ct/total_supply_history_by_bid", api_cosmic_game_cosmic_token_total_supply_history_by_bid)
	r.GET("/api/cosmicgame/ct/total_supply_history_by_date/:from_date/:to_date", api_cosmic_game_cosmic_token_total_supply_history_by_date)

	// User
	r.GET("/api/cosmicgame/user/info/:user_addr", api_cosmic_game_user_info)
	r.GET("/api/cosmicgame/user/notif_red_box/:user_addr", api_cosmic_game_user_global_winnings)
	r.GET("/api/cosmicgame/user/balances/:user_addr", api_cosmic_game_user_balances)

	// Donations
	r.GET("/api/cosmicgame/donations/eth/simple/list/:offset/:limit", api_cosmic_game_donations_cg_simple_list)
	r.GET("/api/cosmicgame/donations/eth/simple/by_round/:round_num", api_cosmic_game_donations_cg_simple_by_round)
	r.GET("/api/cosmicgame/donations/eth/with_info/list/:offset/:limit", api_cosmic_game_donations_cg_with_info_list)
	r.GET("/api/cosmicgame/donations/eth/with_info/by_round/:round_num", api_cosmic_game_donations_cg_with_info_by_round)
	r.GET("/api/cosmicgame/donations/eth/with_info/info/:record_id", api_cosmic_game_donations_cg_with_info_record_info)
	r.GET("/api/cosmicgame/donations/eth/by_user/:user_addr", api_cosmic_game_donations_by_user)
	r.GET("/api/cosmicgame/donations/eth/both/by_round/:round_num", api_cosmic_game_donations_cg_both_by_round)
	r.GET("/api/cosmicgame/donations/eth/both/all", api_cosmic_game_donations_cg_both_all)
	r.GET("/api/cosmicgame/donations/charity/deposits", api_cosmic_game_charity_donations_deposits)
	r.GET("/api/cosmicgame/donations/charity/cg_deposits", api_cosmic_game_charity_cosmicgame_deposits)
	r.GET("/api/cosmicgame/donations/charity/voluntary", api_cosmic_game_charity_voluntary_deposits)
	r.GET("/api/cosmicgame/donations/charity/withdrawals", api_cosmic_game_charity_donations_withdrawals)
	r.GET("/api/cosmicgame/donations/nft/list/:offset/:limit", api_cosmic_game_donations_nft_list)
	r.GET("/api/cosmicgame/donations/nft/info/:record_id", api_cosmic_game_donated_nft_info)
	r.GET("/api/cosmicgame/donations/nft/by_user/:user_addr", api_cosmic_game_nft_donations_by_user)
	r.GET("/api/cosmicgame/donations/nft/claims", api_cosmic_game_donated_nft_claims_all)
	r.GET("/api/cosmicgame/donations/nft/claims/:offset/:limit", api_cosmic_game_donated_nft_claims_all)
	r.GET("/api/cosmicgame/donations/nft/claims/by_user/:user_addr", api_cosmic_game_donated_nft_claims_by_user)
	r.GET("/api/cosmicgame/donations/nft/statistics", api_cosmic_game_nft_donation_stats)
	r.GET("/api/cosmicgame/donations/nft/by_round/:prize_num", api_cosmic_game_nft_donations_by_prize)
	r.GET("/api/cosmicgame/donations/nft/by_token/:token_addr", api_cosmic_game_nft_donations_by_token)
	r.GET("/api/cosmicgame/donations/nft/unclaimed/by_round/:prize_num", api_cosmic_game_unclaimed_donated_nfts_by_prize)
	r.GET("/api/cosmicgame/donations/nft/unclaimed/by_user/:user_addr", api_cosmic_game_unclaimed_donated_nfts_by_user)
	r.GET("/api/cosmicgame/donations/erc20/by_round/detailed/:round_num", api_cosmic_game_donations_erc20_by_round_detailed)
	r.GET("/api/cosmicgame/donations/erc20/by_round/all/:round_num", api_cosmic_game_donations_erc20_by_round_all)
	r.GET("/api/cosmicgame/donations/erc20/by_round/summarized/:round_num", api_cosmic_game_donations_erc20_by_round_summarized)
	r.GET("/api/cosmicgame/donations/erc20/donated/by_user/:user_addr", api_cosmic_game_donations_erc20_donated_by_user)
	r.GET("/api/cosmicgame/donations/erc20/by_user/:user_addr", api_cosmic_game_donations_erc20_by_user)
	r.GET("/api/cosmicgame/donations/erc20/global/:offset/:limit", api_cosmic_game_donations_erc20_global)
	r.GET("/api/cosmicgame/donations/erc20/info/:record_id", api_cosmic_game_donated_erc20_info)
	r.GET("/api/cosmicgame/donations/erc20/claims", api_cosmic_game_erc20_claims_global)
	r.GET("/api/cosmicgame/donations/erc20/claims/:offset/:limit", api_cosmic_game_erc20_claims_global)
	r.GET("/api/cosmicgame/donations/erc20/claims/by_user/:user_addr", api_cosmic_game_erc20_claims_by_user)
	r.GET("/api/cosmicgame/donations/erc20/claims/by_round/:round_num", api_cosmic_game_erc20_claims_by_round)

	// Raffle
	r.GET("/api/cosmicgame/raffle/deposits/list", api_cosmic_game_prize_deposits_list)
	r.GET("/api/cosmicgame/raffle/deposits/list/:offset/:limit", api_cosmic_game_prize_deposits_list)
	r.GET("/api/cosmicgame/raffle/deposits/by_round/:round_num", api_cosmic_game_prize_deposits_by_round)
	r.GET("/api/cosmicgame/eth_deposits/all/list/:offset/:limit", api_cosmic_game_all_eth_deposits_list)
	r.GET("/api/cosmicgame/eth_deposits/raffle_eth/list/:offset/:limit", api_cosmic_game_raffle_eth_deposits_list)
	r.GET("/api/cosmicgame/eth_deposits/chronowarrior_eth/list/:offset/:limit", api_cosmic_game_chronowarrior_eth_deposits_list)
	r.GET("/api/cosmicgame/raffle/nft/all/list", api_cosmic_game_raffle_nft_winners_list)
	r.GET("/api/cosmicgame/raffle/nft/all/list/:offset/:limit", api_cosmic_game_raffle_nft_winners_list)
	r.GET("/api/cosmicgame/raffle/nft/by_round/:round_num", api_cosmic_game_raffle_nft_winners_by_round)
	r.GET("/api/cosmicgame/raffle/nft/by_user/:user_addr", api_cosmic_game_user_raffle_nft_winnings)

	// Staking CST
	r.GET("/api/cosmicgame/staking/cst/staked_tokens/all", api_cosmic_game_staked_tokens_cst_global)
	r.GET("/api/cosmicgame/staking/cst/staked_tokens/by_user/:user_addr", api_cosmic_game_staked_tokens_cst_by_user)
	r.GET("/api/cosmicgame/staking/cst/actions/global/:offset/:limit", api_cosmic_game_staking_actions_cst_global)
	r.GET("/api/cosmicgame/staking/cst/actions/by_user/:user_addr/:offset/:limit", api_cosmic_game_staking_cst_actions_by_user)
	r.GET("/api/cosmicgame/staking/cst/actions/info/:action_id", api_cosmic_game_staking_action_cst_info)
	r.GET("/api/cosmicgame/staking/cst/rewards/global", api_cosmic_game_staking_cst_rewards_global)
	r.GET("/api/cosmicgame/staking/cst/rewards/to_claim/by_user/:user_addr", api_cosmic_game_staking_cst_rewards_to_claim_by_user)
	r.GET("/api/cosmicgame/staking/cst/rewards/collected/by_user/:user_addr/:offset/:limit", api_cosmic_game_staking_cst_rewards_collected_by_user)
	r.GET("/api/cosmicgame/staking/cst/rewards/action_ids_by_deposit/:user_addr/:deposit_id", api_cosmic_game_staking_cst_rewards_action_ids_by_deposit)
	r.GET("/api/cosmicgame/staking/cst/rewards/by_user/by_token/summary/:user_addr", api_cosmic_game_staking_cst_by_user_by_token_rewards)
	r.GET("/api/cosmicgame/staking/cst/rewards/by_user/by_token/details/:user_addr/:token_id", api_cosmic_game_staking_cst_by_user_by_token_rewards_details)
	r.GET("/api/cosmicgame/staking/cst/rewards/by_user/by_deposit/:user_addr", api_cosmic_game_staking_cst_by_user_by_deposit_rewards)
	r.GET("/api/cosmicgame/staking/cst/rewards/by_round/:round_num", api_cosmic_game_staking_cst_rewards_by_round)
	r.GET("/api/cosmicgame/staking/cst/mints/global/:offset/:limit", api_cosmic_game_staking_cst_mints_global)
	r.GET("/api/cosmicgame/staking/cst/mints/by_user/:user_addr", api_cosmic_game_staking_cst_mints_by_user)

	// Staking RWalk (canonical: .../randomwalk/...; legacy: .../rwalk/... — same handlers)
	r.GET("/api/cosmicgame/staking/randomwalk/actions/info/:action_id", api_cosmic_game_staking_action_rwalk_info)
	r.GET("/api/cosmicgame/staking/randomwalk/actions/global/:offset/:limit", api_cosmic_game_staking_actions_rwalk_global)
	r.GET("/api/cosmicgame/staking/randomwalk/actions/by_user/:user_addr/:offset/:limit", api_cosmic_game_staking_actions_rwalk_by_user)
	r.GET("/api/cosmicgame/staking/randomwalk/mints/global/:offset/:limit", api_cosmic_game_staking_rwalk_mints_global)
	r.GET("/api/cosmicgame/staking/randomwalk/mints/by_user/:user_addr", api_cosmic_game_staking_rwalk_mints_by_user)
	r.GET("/api/cosmicgame/staking/randomwalk/staked_tokens/all", api_cosmic_game_staked_tokens_rwalk_global)
	r.GET("/api/cosmicgame/staking/randomwalk/staked_tokens/by_user/:user_addr", api_cosmic_game_staked_tokens_rwalk_by_user)
	r.GET("/api/cosmicgame/staking/rwalk/actions/info/:action_id", api_cosmic_game_staking_action_rwalk_info)
	r.GET("/api/cosmicgame/staking/rwalk/actions/global/:offset/:limit", api_cosmic_game_staking_actions_rwalk_global)
	r.GET("/api/cosmicgame/staking/rwalk/actions/by_user/:user_addr/:offset/:limit", api_cosmic_game_staking_actions_rwalk_by_user)
	r.GET("/api/cosmicgame/staking/rwalk/mints/global/:offset/:limit", api_cosmic_game_staking_rwalk_mints_global)
	r.GET("/api/cosmicgame/staking/rwalk/mints/by_user/:user_addr", api_cosmic_game_staking_rwalk_mints_by_user)
	r.GET("/api/cosmicgame/staking/rwalk/staked_tokens/all", api_cosmic_game_staked_tokens_rwalk_global)
	r.GET("/api/cosmicgame/staking/rwalk/staked_tokens/by_user/:user_addr", api_cosmic_game_staked_tokens_rwalk_by_user)

	// Marketing
	r.GET("/api/cosmicgame/marketing/rewards/global/:offset/:limit", api_cosmic_game_marketing_rewards_global)
	r.GET("/api/cosmicgame/marketing/rewards/by_user/:user_addr/:offset/:limit", api_cosmic_game_marketing_rewards_by_user)
	r.GET("/api/cosmicgame/marketing/config/current", api_cosmic_game_marketing_config_current)

	// Time
	r.GET("/api/cosmicgame/time/current", api_cosmic_game_time_current)
	r.GET("/api/cosmicgame/time/until_prize", api_cosmic_game_time_until_prize)

	// System
	r.GET("/api/cosmicgame/system/modelist", api_cosmic_game_sysmode_changes)
	r.GET("/api/cosmicgame/system/modelist/:offset/:limit", api_cosmic_game_sysmode_changes)
	r.GET("/api/cosmicgame/system/admin_events/:evtlog_start/:evtlog_end", api_cosmic_game_admin_events_in_range)
}
