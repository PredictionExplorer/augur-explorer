package main

import (

	"github.com/gin-gonic/gin"
)

func set_api_routing_augur_v2(r *gin.Engine) {

	r.GET("/api/active_market_ids",a1_active_market_ids)
	r.GET("/api/active_markets/:start/:num_rows",a1_active_markets)
	r.GET("/api/mkt_card/:market_aid",a1_market_card)
	r.GET("/api/mkt_cards/:market_aid_list",a1_multiple_market_cards)
	r.GET("/api/market/:market",  a1_market_info)
	r.GET("/api/user/:user",  a1_user_info)
	r.GET("/api/funds/:user",  a1_user_funds)
	r.GET("/api/user_traded_markets/:user/:active",  a1_user_traded_markets)
	r.GET("/api/user_created_markets/:user",  a1_user_created_markets)
	r.GET("/api/utrades/:user/:market",  a1_user_trades_for_market)
	r.GET("/api/user_pl/:user",  a1_user_profit_loss)
	r.GET("/api/user_opos/:user",  a1_user_open_positions)
	r.GET("/api/user_reports/:user",  a1_user_reports)
	r.GET("/api/user_oorders/:user",  a1_user_open_orders)
	r.GET("/api/mkt_oo/:market/:outcome", a1_market_open_orders)
	r.GET("/api/top_users",  a1_top_users)
	r.GET("/api/stats_main",  a1_stats_main)
	r.GET("/api/stats_cashflow",  a1_stats_cashflow)
	r.GET("/api/stats_uniqaddr",  a1_stats_uniqaddr)
	r.GET("/api/stats_gasused",  a1_stats_gasused)
	r.GET("/api/stats_gasused_accum/:init_ts/:fin_ts/:interval_secs",a1_stats_gasused_accum)
	r.GET("/api/stats_txcost_accum/:init_ts/:fin_ts/:interval_secs",a1_stats_txcost_accum)
	r.GET("/api/stats_txcost/:init_ts/:fin_ts",  a1_stats_txcost)
	r.GET("/api/price_history/:market/:init_ts/:fin_ts/:interval_secs", a1_price_history_zoomed)
	r.GET("/api/stats_accum_trades/:init_ts/:fin_ts/:interval_secs",  a1_stats_accum_trades)
	r.GET("/api/stats_accum_oi/:init_ts/:fin_ts/:interval_secs",  a1_stats_accum_oi)
	r.GET("/api/trade_history/:market",a1_mkt_trade_history)
	r.GET("/api/mkt_stbc/:market/:offset/:limit",a1_market_share_token_balance_changes)
	r.GET("/api/search",a1_search)
	r.GET("/api/categories",a1_categories)
	r.GET("/api/whats_new_augur/:code",a1_whats_new_augur)
	r.GET("/api/tx/:hash",a1_transaction_info)
	r.GET("/api/block/:block_num",a1_block_info)
	r.GET("/api/reporting_table/:market",a1_reporting_table)
	r.GET("/api/user_rep_pl/:user",a1_user_rep_profit_loss)
	r.GET("/api/noshow_bond",a1_noshow_bond_prices)
	r.GET("/api/validity_bond",a1_validity_bond_prices)
}
func set_api_routing_uniswap(r *gin.Engine) {

	r.GET("/api/user_uniswaps/:user/:offset/:limit",a1_user_uniswap_swaps)
	r.GET("/api/mkt_uniswaps/:market",a1_market_uniswap_pairs)
	r.GET("/api/uniswap_swaps/:address/:offset/:limit",a1_uniswap_pair_swaps)
	r.GET("/api/mkt_uniswap_vol/:market/:outcome/:init_ts/:fin_ts/:interval_secs",a1_uniswap_volume)
	r.GET("/api/upair_price_hist/:pair/:inverse/:init_ts/:fin_ts/:interval_secs",a1_upair_price_history)
	r.GET("/api/uni_swap/:id",a1_single_uniswap_swap)
	r.GET("/api/uniswap_slippage/:pair",a1_uniswap_slippage)
}
func set_api_routing_balancer(r *gin.Engine) {

	r.GET("/api/user_balswaps/:user/:offset/:limit",a1_user_balancer_swaps)
	r.GET("/api/pool_swaps/:address/:offset/:limit",a1_pool_swaps)
	r.GET("/api/mkt_pool_vol/:market/:outcome/:init_ts/:fin_ts/:interval_secs",a1_balancer_volume)
	r.GET("/api/pool_price_hist/:pool/:token1/:token2/:init_ts/:fin_ts/:interval_secs",a1_pool_price_history)
	r.GET("/api/bal_swap/:id",a1_single_balancer_swap)
	r.GET("/api/bal_calc_slip/:pool/:tok_in/:tok_out/:amount",a1_balancer_calculate_slippage)
	r.GET("/api/pool_slippage/:pool",a1_pool_slippage)
	r.GET("/api/uni_calc_slip/:pair/:tok_in/:amount",a1_uniswap_calculate_slippage)
	r.GET("/api/aa/pools",a1_arbitrum_augur_pools)
}
func set_api_routing_ens(r *gin.Engine) {

	r.GET("/api/user_ens_names/:user/:offset/:limit",a1_user_ens_names)
	r.GET("/api/rlookup/:address",a1_ens_reverse_lookup)
	r.GET("/api/node_text_data/:node",a1_node_text_key_value_pairs)
	r.GET("/api/ens_name_info/:fqdn",a1_ens_name_info)
	r.GET("/api/ens_lookup/:user",a1_ens_name_lookup)
}
func set_api_routing_augur_foundry(r *gin.Engine) {

	r.GET("/api/wrapped_tokens/:market",a1_wrapped_tokens)
	r.GET("/api/wr_transfers/:address/:offset/:limit",a1_wrapped_token_transfers)
	r.GET("/api/user_wr_transfers/:user/:wrapper/:offset/:limit",a1_user_wrapped_token_transfers)
	r.GET("/api/wr_vol/:address/:init_ts/:fin_ts/:interval_secs",a1_wrapped_token_volume)
	r.GET("/api/wshtok_balances/:user",a1_wrapped_shtoken_balances)
	r.GET("/api/augur_foundry",a1_augur_foundry_contracts)
}
func set_api_routing_augur_amm(r *gin.Engine) {

	r.GET("/api/augur_amm/markets/sports/:status/:sort/:offset/:limit",a1_arbitrum_markets_sports)
	r.GET("/api/augur_amm/markets/info/sports/:factory_aid/:market_id",a1_arbitrum_market_info_sports)
	r.GET("/api/augur_amm/liquidity/:factory_aid/:market_id/:offset/:limit",a1_arbitrum_liquidity_changed)
	r.GET("/api/augur_amm/swaps/:contract_aid/:market_id/:offset/:limit",a1_arbitrum_shares_swapped)
	r.GET("/api/augur_amm/user/swaps/:user/:offset/:limit",a1_amm_user_swaps)
	r.GET("/api/augur_amm/user/liquidity/:user/:offset/:limit",a1_amm_user_liquidity)
	r.GET("/api/augur_amm/market/liquidity/providers/:factory_aid/:market_id",a1_arbitrum_market_liquidity_providers)
	r.GET("/api/augur_amm/market/outside/shares_burned/:factory_aid/:market_id/:offset/:limit",a1_arbitrum_market_outside_augur_shares_burned)
	r.GET("/api/augur_amm/market/outside/shares_minted/:factory_aid/:market_id/:offset/:limit",a1_arbitrum_market_outside_augur_shares_minted)
	r.GET("/api/augur_amm/market/outside/balancer_swaps/:factory_aid/:market_id/:offset/:limit",a1_arbitrum_market_outside_augur_balancer_swaps)
	r.GET("/api/augur_amm/market/outside/erc20_transfers/:factory_aid/:market_id/:offset/:limit",a1_arbitrum_market_outside_augur_erc20_transfers)
}
func set_api_routing_polymarket(r *gin.Engine) {

	// Polymarket API requests
	r.GET("/api/poly/markets/buysell/list/:market_id/:offset/:limit",a1_poly_buysell_operations)
	r.GET("/api/poly/markets/buysell/info/:id",a1_poly_market_buysell_info)
	r.GET("/api/poly/markets/liquidity/:market_id/:offset/:limit",a1_poly_liquidity_operations)
	r.GET("/api/poly/markets/info/:market_id",a1_poly_market_info)
	r.GET("/api/poly/markets/statistics/:market_id",a1_poly_market_stats)
	r.GET("/api/poly/markets/volume/liquidity/:market_id/:init_ts/:fin_ts/:interval_secs",a1_poly_market_liquidity_periods)
	r.GET("/api/poly/markets/volume/trading/:market_id/:init_ts/:fin_ts/:interval_secs",a1_poly_market_trading_periods)
	r.GET("/api/poly/markets/userlist/:market_id",a1_poly_user_list)
	r.GET("/api/poly/markets/traderops/:market_id/:user_aid/:offset/:limit",a1_poly_trader_operations)
	r.GET("/api/poly/markets/funderops/:market_id/:user_aid/:offset/:limit",a1_poly_funder_operations)
	r.GET("/api/poly/markets/redemptions/:market_id/:offset/:limit",a1_poly_market_payout_redemptions)
	r.GET("/api/poly/markets/list/:status/:sort",a1_poly_markets_listing)
	r.GET("/api/poly/markets/list/:status",a1_poly_markets_listing)
	r.GET("/api/poly/markets/list",a1_poly_markets_listing)
	r.GET("/api/poly/markets/open_positions/:market_id",a1_poly_market_open_positions)
	r.GET("/api/poly/market/erc1155/:market_id/:offset/:limit",a1_poly_market_erc1155_transfers)
	r.GET("/api/poly/markets/funder/share_ratio/:market_id",a1_poly_market_funder_share_ratio)
	r.GET("/api/poly/markets/price_history/:market_id/:outcome",a1_poly_market_price_history)
	r.GET("/api/poly/markets/open_interest_history/:market_id/:offset/:limit",a1_market_open_interest_history)
	r.GET("/api/poly/user/open_positions/:user_aid",a1_poly_market_user_open_positions)
	r.GET("/api/poly/user/info/:user",a1_poly_user_info)
	r.GET("/api/poly/user/traded_markets/:user",a1_poly_user_traded_markets)
	r.GET("/api/poly/top_users",  a1_poly_top_users)
	r.GET("/api/poly/categories",a1_poly_categories)
	r.GET("/api/poly/search",a1_poly_market_search)
	r.GET("/api/poly/datafeed",a1_poly_datafeed)
	r.GET("/api/poly/unique_users/:init_ts/:fin_ts",a1_poly_unique_users)
	r.GET("/api/poly/stats/global_liquidity/:init_ts/:fin_ts/:interval_secs",a1_poly_liq_hist_global)
	r.GET("/api/poly/stats/global_trading/:init_ts/:fin_ts/:interval_secs",a1_poly_trade_hist_global)
}
func set_api_routing_randomwalk(r *gin.Engine) {
	r.GET("/api/rwalk/current_offers/:rwalk_addr/:market_addr/:order_by",api_rwalk_current_offers)
	r.GET("/api/rwalk/floor_price/:rwalk_addr/:market_addr",api_rwalk_floor_price)
	r.GET("/api/rwalk/tokens/list/sequential/:rwalk_addr",api_rwalk_token_list_seq)
	r.GET("/api/rwalk/tokens/list/sequential/:rwalk_addr/:offset/:limit",api_rwalk_token_list_seq)
	r.GET("/api/rwalk/tokens/list/by_period/:rwalk_addr/:init_ts/:fin_ts",api_rwalk_token_list_period)
	r.GET("/api/rwalk/tokens/info/:rwalk_addr/:token_id",api_rwalk_token_info)
	r.GET("/api/rwalk/tokens/name_changes/:token_id",api_rwalk_token_name_history)
	r.GET("/api/rwalk/trading/history/:market_addr/:offset/:limit",api_rwalk_trading_history)
	r.GET("/api/rwalk/trading/by_user/:user_aid/:offset/:limit",api_rwalk_trading_history_by_user)
	r.GET("/api/rwalk/trading/sales/:market_addr/:offset/:limit",api_rwalk_sale_history)
	r.GET("/api/rwalk/tokens/history/:token_id/:rwalk_addr/:offset/:limit",api_rwalk_token_history)
	r.GET("/api/rwalk/tokens/by_user/:user_aid",api_rwalk_tokens_by_user)
	r.GET("/api/rwalk/statistics/by_token/:rwalk_addr",api_rwalk_token_stats)
	r.GET("/api/rwalk/statistics/by_market/:market_addr",api_rwalk_market_stats)
	r.GET("/api/rwalk/statistics/trading_volume/:market_addr/:init_ts/:fin_ts/:interval_secs",api_rwalk_trading_volume_by_period)
	r.GET("/api/rwalk/statistics/mint_intervals/:rwalk_addr",api_rwalk_mint_intervals)
	r.GET("/api/rwalk/statistics/floor_price/:market_addr/:rwalk_addr/:init_ts/:fin_ts/:interval_secs",api_rwalk_floor_price_over_time)
	r.GET("/api/rwalk/statistics/withdrawal_chart/:rwalk_addr",api_rwalk_withdrawal_chart)
	r.GET("/api/rwalk/user/info/:user_aid/:rwalk_addr",api_rwalk_user_info)
	r.GET("/api/rwalk/top5tokens",api_rwalk_top5_traded_tokens)
	r.GET("/api/rwalk/mint_report",api_rwalk_mint_report)
}
func set_api_routing_biddingwar(r *gin.Engine) {

	r.GET("/api/cosmicgame/statistics/dashboard",api_biddingwar_dashboard)
	r.GET("/api/cosmicgame/statistics/counters",api_biddingwar_record_counters)
	r.GET("/api/cosmicgame/prize/list/:offset/:limit",api_biddingwar_prize_list)
	r.GET("/api/cosmicgame/prize/info/:prize_num",api_biddingwar_prize_info)
	r.GET("/api/cosmicgame/bid/list/:offset/:limit",api_biddingwar_bid_list)
	r.GET("/api/cosmicgame/bid/info/:evtlog_id",api_biddingwar_bid_info)
	r.GET("/api/cosmicgame/bid/list_by_round/:round_num/:sort/:offset/:limit",api_biddingwar_bid_list_by_round)
	r.GET("/api/cosmicgame/cst/list/:offset/:limit",api_biddingwar_cosmic_signature_token_list)
	r.GET("/api/cosmicgame/cst/info/:token_id",api_biddingwar_cosmic_signature_token_info)
	r.GET("/api/cosmicgame/user/info/:user_addr",api_biddingwar_user_info)
	r.GET("/api/cosmicgame/user/raffle_deposits/:user_addr",api_biddingwar_user_raffle_deposits)
	r.GET("/api/cosmicgame/user/raffle_nft_claims/:user_addr",api_biddingwar_user_raffle_nft_claims)
	r.GET("/api/cosmicgame/user/raffle_nft_winnings/:user_addr",api_biddingwar_user_raffle_nft_winnings)
	r.GET("/api/cosmicgame/user/unique_bidders",api_biddingwar_user_unique_bidders)
	r.GET("/api/cosmicgame/user/unique_winners",api_biddingwar_user_unique_winners)
	r.GET("/api/cosmicgame/user/nft/claims/:user_addr",api_biddingwar_donated_nft_claims_by_user)
	r.GET("/api/cosmicgame/donations/eth",api_biddingwar_donations_eth)
	r.GET("/api/cosmicgame/donations/charity",api_biddingwar_charity_donations)
	r.GET("/api/cosmicgame/donations/nft/list/:offset/:limit",api_biddingwar_donations_nft_list)
	r.GET("/api/cosmicgame/donations/nft/info/:record_id",api_biddingwar_donated_nft_info)
	r.GET("/api/cosmicgame/donations/nft/claims/:offset/:limit",api_biddingwar_donated_nft_claims_all)
	r.GET("/api/cosmicgame/donations/nft/statistics",api_biddingwar_nft_donation_stats)
	r.GET("/api/cosmicgame/donations/nft/by_prize/:prize_num",api_biddingwar_nft_donations_by_prize)
	r.GET("/api/cosmicgame/raffle/deposits/list/:offset/:limit",api_biddingwar_raffle_deposits_list)
	r.GET("/api/cosmicgame/raffle/deposits/by_round/:round_num",api_biddingwar_raffle_deposits_by_round)
	r.GET("/api/cosmicgame/raffle/nft_winners/list/:offset/:limit",api_biddingwar_raffle_nft_winners_list)
	r.GET("/api/cosmicgame/raffle/nft_winners/by_round/:round_num",api_biddingwar_raffle_nft_winners_by_round)
	r.GET("/api/cosmicgame/raffle/nft_claims/:offset/:limit",api_biddingwar_raffle_nft_claims)
	r.GET("/api/cosmicgame/time/current",api_biddingwar_time_current)
	r.GET("/api/cosmicgame/time/until_prize",api_biddingwar_time_until_prize)
}
func set_api_routing_statistics(r *gin.Engine) {

	r.GET("/api/statistics/main_net/:init_ts/:fin_ts",api_stats_main_statistics_main_net)
	r.GET("/api/statistics/arbitrum/:init_ts/:fin_ts",api_stats_main_statistics_arbitrum)
	r.GET("/api/statistics/timeframe_ranges",api_stats_get_timeframe_ranges)
}
func set_routing_api(r *gin.Engine) {
	set_api_routing_augur_v2(r)
	set_api_routing_uniswap(r)
	set_api_routing_balancer(r)
	set_api_routing_ens(r)
	set_api_routing_augur_foundry(r)
	set_api_routing_augur_amm(r)
	set_api_routing_polymarket(r)
	set_api_routing_randomwalk(r)
	set_api_routing_biddingwar(r)
	set_api_routing_statistics(r)
}
