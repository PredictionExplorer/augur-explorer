// Package cosmicgame provides HTTP handlers for CosmicGame functionality
package cosmicgame

import (
	"log"

	"github.com/ethereum/go-ethereum/ethclient"
	ethrpc "github.com/ethereum/go-ethereum/rpc"
	"github.com/gin-gonic/gin"

	. "github.com/PredictionExplorer/augur-explorer/rwcg/dbs/cosmicgame"
	"github.com/PredictionExplorer/augur-explorer/rwcg/websrv/api/common"
)

var (
	// EthClient for blockchain calls
	EthClient *ethclient.Client

	// RpcClient for direct RPC calls
	rpcclient *ethrpc.Client

	// ArbStoragew is the CosmicGame database wrapper
	ArbStoragew SQLStorageWrapper

	// Loggers
	Info  *log.Logger
	Error *log.Logger

	// Enabled indicates whether the cosmicgame module is enabled
	Enabled bool
)

// Init initializes the cosmicgame package with required dependencies
// If enabled is false, the module will be disabled and routes will not be functional
func Init(ethClient *ethclient.Client, rpcClient *ethrpc.Client, info, errorLog *log.Logger, enabled bool) {
	Enabled = enabled
	if !enabled {
		info.Printf("CosmicGame module is disabled")
		return
	}

	EthClient = ethClient
	rpcclient = rpcClient
	Info = info
	Error = errorLog

	// Initialize the storage wrapper
	if common.Ctx != nil && common.Ctx.Db != nil {
		ArbStoragew.S = common.Ctx.Db
		ArbStoragew.S.Db_set_schema_name("public")
	}

	// Call the cosmic game initialization
	cosmic_game_init()
}

// Helper to check if database is initialized
func dbInitialized() bool {
	return common.Ctx != nil && common.Ctx.Db != nil
}

// RegisterAPIRoutes registers all CosmicGame JSON API routes
// If the module is disabled, this function returns without registering any routes
func RegisterAPIRoutes(r *gin.Engine) {
	if !Enabled {
		return
	}
	// Statistics
	r.GET("/api/cosmicgame/statistics/dashboard", api_cosmic_game_dashboard)
	r.GET("/api/cosmicgame/statistics/counters", api_cosmic_game_record_counters)
	r.GET("/api/cosmicgame/statistics/unique/bidders", api_cosmic_game_user_unique_bidders)
	r.GET("/api/cosmicgame/statistics/unique/winners", api_cosmic_game_user_unique_winners)
	r.GET("/api/cosmicgame/statistics/unique/donors", api_cosmic_game_user_unique_donors)
	r.GET("/api/cosmicgame/statistics/unique/stakers/cst", api_cosmic_game_user_unique_stakers_cst)
	r.GET("/api/cosmicgame/statistics/unique/stakers/rwalk", api_cosmic_game_user_unique_stakers_rwalk)
	r.GET("/api/cosmicgame/statistics/unique/stakers/both", api_cosmic_game_user_unique_stakers_both)

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
	r.GET("/api/cosmicgame/prizes/deposits/raffle/by_user/:use_addr", api_cosmic_game_prize_deposits_raffle_eth_by_user)
	r.GET("/api/cosmicgame/prizes/deposits/chrono_warrior/by_user/:user_addr", api_cosmic_game_prize_deposits_chrono_warrior_by_user)
	r.GET("/api/cosmicgame/prizes/deposits/unclaimed/by_user/:user_addr/:offset/:limit", api_cosmic_game_unclaimed_prize_deposits_by_user)

	// Bids
	r.GET("/api/cosmicgame/bid/list/all/:offset/:limit", api_cosmic_game_bid_list)
	r.GET("/api/cosmicgame/bid/info/:evtlog_id", api_cosmic_game_bid_info)
	r.GET("/api/cosmicgame/bid/list/by_round/:round_num/:sort/:offset/:limit", api_cosmic_game_bid_list_by_round)
	r.GET("/api/cosmicgame/bid/used_rwalk_nfts", api_cosmic_game_used_rwalk_nfts)
	r.GET("/api/cosmicgame/bid/cst_price", api_cosmic_game_get_cst_price)
	r.GET("/api/cosmicgame/bid/eth_price", api_cosmic_game_get_eth_price)
	r.GET("/api/cosmicgame/bid/current_special_winners", api_cosmic_game_bid_special_winners)

	// CST Tokens
	r.GET("/api/cosmicgame/cst/list/all/:offset/:limit", api_cosmic_game_cosmic_signature_token_list)
	r.GET("/api/cosmicgame/cst/list/by_user/:user_addr/:offset/:limit", api_cosmic_game_cosmic_signature_token_list_by_user)
	r.GET("/api/cosmicgame/cst/info/:token_id", api_cosmic_game_cosmic_signature_token_info)
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
	r.GET("/api/cosmicgame/donations/nft/unclaimed/by_round/:prize_num", api_cosmic_game_unclaimed_donated_nfts_by_prize)
	r.GET("/api/cosmicgame/donations/nft/unclaimed/by_user/:user_addr", api_cosmic_game_unclaimed_donated_nfts_by_user)
	r.GET("/api/cosmicgame/donations/erc20/by_round/detailed/:round_num", api_cosmic_game_donations_erc20_by_round_detailed)
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

	// Staking RWalk
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

// RegisterHTMLRoutes registers all CosmicGame HTML page routes
// If the module is disabled, this function returns without registering any routes
func RegisterHTMLRoutes(r *gin.Engine) {
	if !Enabled {
		return
	}
	r.GET("/black/cosmicgame", cosmic_game_index_page)
	r.GET("/black/cosmicgame/index.html", cosmic_game_index_page)
	r.GET("/black/cosmicgame/api_docs", cosmic_game_api_docs)
	r.GET("/black/cosmicgame/statistics/unique/bidders", cosmic_game_unique_bidders)
	r.GET("/black/cosmicgame/statistics/unique/winners", cosmic_game_unique_winners)
	r.GET("/black/cosmicgame/statistics/unique/donors", cosmic_game_unique_donors)
	r.GET("/black/cosmicgame/statistics/unique/stakers/cst", cosmic_game_unique_stakers_cst)
	r.GET("/black/cosmicgame/statistics/unique/stakers/rwalk", cosmic_game_unique_stakers_rwalk)
	r.GET("/black/cosmicgame/statistics/unique/stakers/both", cosmic_game_unique_stakers_both)
	r.GET("/black/cosmicgame/rounds/list", cosmic_game_prize_claims)
	r.GET("/black/cosmicgame/rounds/info/:prize_num", cosmic_game_round_info)
	r.GET("/black/cosmicgame/prizes/history/global", cosmic_game_global_claim_history_detail)
	r.GET("/black/cosmicgame/prizes/history/by_user/:user_addr/:offset/:limit", cosmic_game_prize_history_detail_by_user)
	r.GET("/black/cosmicgame/prizes/eth/all", cosmic_game_prizes_eth_all)
	r.GET("/black/cosmicgame/prizes/eth/all/:offset/:limit", cosmic_game_prizes_eth_all)
	r.GET("/black/cosmicgame/prizes/eth/raffle", cosmic_game_prizes_eth_raffle)
	r.GET("/black/cosmicgame/prizes/eth/raffle/:offset/:limit", cosmic_game_prizes_eth_raffle)
	r.GET("/black/cosmicgame/prizes/eth/chronowarrior", cosmic_game_prizes_eth_chronowarrior)
	r.GET("/black/cosmicgame/prizes/eth/chronowarrior/:offset/:limit", cosmic_game_prizes_eth_chronowarrior)
	r.GET("/black/cosmicgame/prizes/eth/round/:round_num", cosmic_game_prizes_eth_round)
	r.GET("/black/cosmicgame/prizes/eth/user/:user_addr", cosmic_game_prizes_eth_user)
	r.GET("/black/cosmicgame/prizes/eth/user_raffle/:user_addr", cosmic_game_prizes_eth_user_raffle)
	r.GET("/black/cosmicgame/prizes/eth/user_chronowarrior/:user_addr", cosmic_game_prizes_eth_user_chronowarrior)
	r.GET("/black/cosmicgame/prizes/eth/user_unclaimed/:user_addr", cosmic_game_prizes_eth_user_unclaimed)
	r.GET("/black/cosmicgame/bid/list/all", cosmic_game_bids)
	r.GET("/black/cosmicgame/bid/info/:evtlog_id", cosmic_game_bid_info)
	r.GET("/black/cosmicgame/bid/list/by_round/:round_num/:sort/:offset/:limit", cosmic_game_bid_list_by_round)
	r.GET("/black/cosmicgame/bid/used_rwalk_nfts", cosmic_game_used_rwalk_nfts)
	r.GET("/black/cosmicgame/bid/cst_price", cosmic_game_get_cst_price)
	r.GET("/black/cosmicgame/bid/eth_price", cosmic_game_get_eth_price)
	r.GET("/black/cosmicgame/bid/current_special_winners", cosmic_game_bid_special_winners)
	r.GET("/black/cosmicgame/cst/list/all", cosmic_game_cosmic_signature_token_list)
	r.GET("/black/cosmicgame/cst/list/by_user/:user_addr", cosmic_game_cosmic_signature_token_list_by_user)
	r.GET("/black/cosmicgame/cst/info/:token_id", cosmic_game_cosmic_signature_token_info)
	r.GET("/black/cosmicgame/cst/names/history/:token_id", cosmic_game_token_name_history)
	r.GET("/black/cosmicgame/cst/names/search/:name", cosmic_game_token_name_search)
	r.GET("/black/cosmicgame/cst/names/named_only", cosmic_game_named_tokens_only)
	r.GET("/black/cosmicgame/cst/transfers/all/:token_id", cosmic_game_token_ownership_transfers)
	r.GET("/black/cosmicgame/cst/transfers/by_user/:user_addr", cosmic_game_cosmic_signature_transfers_by_user)
	r.GET("/black/cosmicgame/cst/distribution", cosmic_game_cs_token_distribution)
	r.GET("/black/cosmicgame/ct/balances", cosmic_game_cosmic_token_balances)
	r.GET("/black/cosmicgame/ct/statistics", cosmic_game_cosmic_token_statistics)
	r.GET("/black/cosmicgame/ct/summary/by_user/:user_addr", cosmic_game_cosmic_token_summary_by_user)
	r.GET("/black/cosmicgame/ct/transfers/by_user/:user_addr", cosmic_game_cosmic_token_transfers_by_user)
	r.GET("/black/cosmicgame/user/info/:user_addr", cosmic_game_user_info)
	r.GET("/black/cosmicgame/user/notif_red_box/:user_addr", cosmic_game_user_notif_red_box_rewards)
	r.GET("/black/cosmicgame/user/balances/:user_addr", cosmic_game_user_balances)
	r.GET("/black/cosmicgame/donations/eth/simple/list", cosmic_game_donations_cg_simple_list)
	r.GET("/black/cosmicgame/donations/eth/simple/by_round/:round_num", cosmic_game_donations_cg_simple_by_round)
	r.GET("/black/cosmicgame/donations/eth/with_info/list", cosmic_game_donations_cg_with_info_list)
	r.GET("/black/cosmicgame/donations/eth/with_info/by_round/:round_num", cosmic_game_donations_cg_with_info_by_round)
	r.GET("/black/cosmicgame/donations/eth/with_info/info/:record_id", cosmic_game_donations_cg_with_info_record_info)
	r.GET("/black/cosmicgame/donations/eth/by_user/:user_addr", cosmic_game_donations_by_user)
	r.GET("/black/cosmicgame/donations/eth/both/by_round/:round_num", cosmic_game_donations_cg_both_by_round)
	r.GET("/black/cosmicgame/donations/eth/both/all", cosmic_game_donations_cg_both_all)
	r.GET("/black/cosmicgame/donations/charity/deposits", cosmic_game_charity_donations_deposits)
	r.GET("/black/cosmicgame/donations/charity/cg_deposits", cosmic_game_charity_cosmicgame_deposits)
	r.GET("/black/cosmicgame/donations/charity/voluntary", cosmic_game_charity_voluntary_deposits)
	r.GET("/black/cosmicgame/donations/charity/withdrawals", cosmic_game_charity_donations_withdrawals)
	r.GET("/black/cosmicgame/donations/nft/list", cosmic_game_donations_nft)
	r.GET("/black/cosmicgame/donations/nft/info/:record_id", cosmic_game_donations_nft_info)
	r.GET("/black/cosmicgame/donations/nft/by_user/:user_addr", cosmic_game_nft_donations_by_user)
	r.GET("/black/cosmicgame/donations/nft/claims", cosmic_game_donated_nft_claims_all)
	r.GET("/black/cosmicgame/donations/nft/claims/:offset/:limit", cosmic_game_donated_nft_claims_all)
	r.GET("/black/cosmicgame/donations/nft/claims/by_user/:user_addr", cosmic_game_donated_nft_claims_by_user)
	r.GET("/black/cosmicgame/donations/nft/statistics", cosmic_game_nft_donation_stats)
	r.GET("/black/cosmicgame/donations/nft/by_round/:prize_num", cosmic_game_nft_donations_by_prize)
	r.GET("/black/cosmicgame/donations/nft/unclaimed/by_round/:prize_num", cosmic_game_unclaimed_donated_nfts_by_prize)
	r.GET("/black/cosmicgame/donations/nft/unclaimed/by_user/:user_addr", cosmic_game_unclaimed_donated_nfts_by_user)
	r.GET("/black/cosmicgame/donations/erc20/by_round/detailed/:round_num", cosmic_game_donations_erc20_by_round_detailed)
	r.GET("/black/cosmicgame/donations/erc20/by_round/summarized/:round_num", cosmic_game_donations_erc20_by_round_summarized)
	r.GET("/black/cosmicgame/donations/erc20/donated/by_user/:user_addr", cosmic_game_donations_erc20_donated_by_user)
	r.GET("/black/cosmicgame/donations/erc20/by_user/:user_addr", cosmic_game_donations_erc20_by_user)
	r.GET("/black/cosmicgame/donations/erc20/global", cosmic_game_donations_erc20_global)
	r.GET("/black/cosmicgame/donations/erc20/info/:record_id", cosmic_game_donations_erc20_info)
	r.GET("/black/cosmicgame/donations/erc20/claims", cosmic_game_erc20_claims_global)
	r.GET("/black/cosmicgame/donations/erc20/claims/:offset/:limit", cosmic_game_erc20_claims_global)
	r.GET("/black/cosmicgame/donations/erc20/claims/by_user/:user_addr", cosmic_game_erc20_claims_by_user)
	r.GET("/black/cosmicgame/raffle/nft/all/list", cosmic_game_raffle_nft_winners_list)
	r.GET("/black/cosmicgame/raffle/nft/all/list/:offset/:limit", cosmic_game_raffle_nft_winners_list)
	r.GET("/black/cosmicgame/raffle/nft/by_round/:round_num", cosmic_game_raffle_nft_winners_by_round)
	r.GET("/black/cosmicgame/raffle/nft/by_user/:user_addr", cosmic_game_user_raffle_nft_winnings)
	r.GET("/black/cosmicgame/staking/cst/staked_tokens/all", cosmic_game_staked_tokens_cst_global)
	r.GET("/black/cosmicgame/staking/cst/staked_tokens/by_user/:user_addr", cosmic_game_staked_tokens_cst_by_user)
	r.GET("/black/cosmicgame/staking/cst/actions/global", cosmic_game_staking_cst_actions_global)
	r.GET("/black/cosmicgame/staking/cst/actions/by_user/:user_addr", cosmic_game_staking_actions_cst_by_user)
	r.GET("/black/cosmicgame/staking/cst/actions/info/:action_id", cosmic_game_staking_cst_action_info)
	r.GET("/black/cosmicgame/staking/cst/rewards/global", cosmic_game_staking_cst_rewards_global)
	r.GET("/black/cosmicgame/staking/cst/rewards/to_claim/by_user/:user_addr", cosmic_game_staking_cst_rewards_to_claim_by_user)
	r.GET("/black/cosmicgame/staking/cst/rewards/collected/by_user/:user_addr", cosmic_game_staking_cst_rewards_collected_by_user)
	r.GET("/black/cosmicgame/staking/cst/rewards/action_ids_by_deposit/:user_addr/:deposit_id", cosmic_game_staking_cst_rewards_action_ids_by_deposit)
	r.GET("/black/cosmicgame/staking/cst/rewards/by_user/by_token/summary/:user_addr", cosmic_game_staking_cst_by_user_by_token_rewards)
	r.GET("/black/cosmicgame/staking/cst/rewards/by_user/by_token/details/:user_addr/:token_id", cosmic_game_staking_cst_by_user_by_token_rewards_details)
	r.GET("/black/cosmicgame/staking/cst/rewards/by_user/by_deposit/:user_addr", cosmic_game_staking_cst_by_user_by_deposit_rewards)
	r.GET("/black/cosmicgame/staking/cst/rewards/by_round/:round_num", cosmic_game_staking_cst_rewards_by_round)
	r.GET("/black/cosmicgame/staking/cst/mints/global", cosmic_game_staking_cst_mints_global)
	r.GET("/black/cosmicgame/staking/cst/mints/by_user/:user_addr", cosmic_game_staking_cst_mints_by_user)
	r.GET("/black/cosmicgame/staking/rwalk/actions/info/:action_id", cosmic_game_staking_action_rwalk_info)
	r.GET("/black/cosmicgame/staking/rwalk/actions/global", cosmic_game_staking_actions_rwalk_global)
	r.GET("/black/cosmicgame/staking/rwalk/actions/by_user/:user_addr", cosmic_game_staking_actions_rwalk_by_user)
	r.GET("/black/cosmicgame/staking/rwalk/mints/global", cosmic_game_staking_rwalk_mints_global)
	r.GET("/black/cosmicgame/staking/rwalk/mints/by_user/:user_addr", cosmic_game_staking_rwalk_mints_by_user)
	r.GET("/black/cosmicgame/staking/rwalk/staked_tokens/all", cosmic_game_staked_tokens_rwalk_global)
	r.GET("/black/cosmicgame/staking/rwalk/staked_tokens/by_user/:user_addr", cosmic_game_staked_tokens_rwalk_by_user)
	r.GET("/black/cosmicgame/marketing/rewards/global", cosmic_game_marketing_rewards_global)
	r.GET("/black/cosmicgame/marketing/rewards/by_user/:user_addr", cosmic_game_marketing_rewards_by_user)
	r.GET("/black/cosmicgame/marketing/config/current", cosmic_game_marketing_config_current)
	r.GET("/black/cosmicgame/time/current", cosmic_game_time_current)
	r.GET("/black/cosmicgame/time/until_prize", cosmic_game_time_until_prize)
	r.GET("/black/cosmicgame/dev_funcs", cosmic_game_dev_funcs)
	r.GET("/black/cosmicgame/system/modelist", cosmic_game_sysmode_changes)
	r.GET("/black/cosmicgame/system/modelist/:offset/:limit", cosmic_game_sysmode_changes)
	r.GET("/black/cosmicgame/system/admin_events/:evtlog_start/:evtlog_end", cosmic_game_admin_events_in_range)
}

