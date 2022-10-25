package main

import (

	"github.com/gin-gonic/gin"
)
func set_routing_augur_v2(r *gin.Engine) {

	r.GET("/black/markets.html",markets)
	r.GET("/black/statistics.html",statistics)
	r.GET("/black/categories.html",categories)
	r.GET("/black/explorer.html",explorer)
	r.GET("/black/market/:market",  market_info)
	r.GET("/black/fulltradelist/:market",  full_trade_list)
	r.GET("/black/mdepth/:market/:outcome", market_depth)
	r.GET("/black/deptha/:market_aid/:outcome", market_depth_ajax)
	r.GET("/black/mphist/:market/:outcome", market_price_history)
	r.GET("/black/search", search_v2)
	r.GET("/black/money/:addr",  read_money)
	r.GET("/black/order/:order",  order)
	r.GET("/black/category/:catid",  category)
	r.GET("/black/user/:addr",  user_info)
	r.GET("/black/fullreports/:addr",  full_reports)
	r.GET("/black/umarkets/:addr",  user_markets)
	r.GET("/black/udw/:addr",  user_deposits_withdrawals)
	r.GET("/black/block/:block_num",  block_info)
	r.GET("/black/topusers.html",top_users)
	r.GET("/black/mdstat/:market_aid/:outcome_idx/:last_oo_id",market_depth_status)
	r.GET("/black/umtrades.html",user_trades_for_market)
	r.GET("/black/user_rep_pl/:user",user_rep_profit_loss)
	r.GET("/black/statement/:addr",account_statement)
	r.GET("/black/oohist/:addr",open_order_history)
	r.GET("/black/pehist/:market/:outcome", price_estimate_history)
	r.GET("/black/stbc/:market",sharetoken_balance_changes)
	r.GET("/black/text_search",do_text_search)
	r.GET("/black/text_search_form",show_text_search_form)
	r.GET("/black/whats_new_augur/",whats_new_in_augur)
	r.GET("/black/reports_table/:market",show_reporting_table)
	r.GET("/black/noshow_bond",augur_noshow_bond_prices)
	r.GET("/black/validity_bond",augur_validity_bond_prices)
}
func set_routing_augur_foundry(r *gin.Engine) {

	r.GET("/black/user_wr_transfers/:user/:wrapper",user_wrapped_token_transfers)
	r.GET("/black/wrapped/:market",wrapped_tokens)
	r.GET("/black/wr_transfers/:address",wrapped_token_transfers)
	r.GET("/black/wrtok_info/:address",wrapped_token_info)
	r.GET("/black/augur_foundry",show_augur_foundry_contracts)
}
func set_routing_uniswap(r *gin.Engine) {

	r.GET("/black/user_uswaps/:user",user_uniswap_swaps)
	r.GET("/black/mkt_uniswaps/:market",market_uniswap_pairs)
	r.GET("/black/uniswap_swaps/:address",uniswap_swaps)
	r.GET("/black/upair_swap_price/:pair_aid/:inverse/:init_ts/:fin_ts/:interval_secs",show_upair_swap_prices)
	r.GET("/black/uni_swap/:id",show_single_uniswap_swap)
	r.GET("/black/uniswap_slippage/:pair",show_uniswap_slippage)
	r.GET("/black/rt_uniswap_slippage/:pair",rt_show_uniswap_slippage)
}
func set_routing_balancer(r *gin.Engine) {

	r.GET("/black/user_bswaps/:user",user_balancer_swaps)
	r.GET("/black/pool_swaps/:address",pool_swaps)
	r.GET("/black/pool_swap_price/:pool_aid/:token1_aid/:token2_aid/:init_ts/:fin_ts/:interval_secs",show_pool_swap_prices)
	r.GET("/black/bal_swap/:id",show_single_balancer_swap)
	r.GET("/black/pool_slippage/:pool",show_pool_slippage)
}
func set_routing_ens(r *gin.Engine) {

	r.GET("/black/user_ens_names/:user",user_ens_names)
	r.GET("/black/node_text_data/:node",show_node_text_data)
	r.GET("/black/ens_name_info/:fqdn",ens_name_info)
}
func set_routing_augur_amm(r *gin.Engine) {

	r.GET("/black/augur_amm/pools",arbitrum_augur_pools)
	r.GET("/black/augur_amm/markets/sports/:status/:sort",arbitrum_markets_sports)
	r.GET("/black/augur_amm/liquidity/:factory_aid/:market_id/:offset/:limit",arbitrum_liquidity_changed)
	r.GET("/black/augur_amm/swaps/:contract_aid/:market_id/:offset/:limit",arbitrum_shares_swapped)
	r.GET("/black/augur_amm/user/swaps/:user/:offset/:limit",amm_user_swaps)
	r.GET("/black/augur_amm/user/liquidity/:user/:offset/:limit",amm_user_liquidity)
	r.GET("/black/augur_amm/markets/info/:contract_aid/:market_id",arbitrum_market_info)
	r.GET("/black/augur_amm/market/liquidity/providers/:contract_aid/:market_id",arbitrum_market_liquidity_providers)
	r.GET("/black/augur_amm/market/outside/shares_burned/:contract_aid/:market_id/:offset/:limit",
															arbitrum_market_outside_augur_shares_burned)
	r.GET("/black/augur_amm/market/outside/shares_minted/:contract_aid/:market_id/:offset/:limit",
															arbitrum_market_outside_augur_shares_minted)
	r.GET("/black/augur_amm/market/outside/balancer_swaps/:contract_aid/:market_id/:offset/:limit",
															arbitrum_market_outside_augur_balancer_swaps)
	r.GET("/black/augur_amm/market/outside/erc20_transfers/:contract_aid/:market_id/:offset/:limit",
															arbitrum_market_outside_augur_erc20_transfers)
}
func set_routing_polymarket(r *gin.Engine) {

	r.GET("/black/poly/markets/buysell/:market_id/:offset/:limit",poly_buysell_operations)
	r.GET("/black/poly/markets/liquidity/:market_id/:offset/:limit",poly_liquidity_operations)
	r.GET("/black/poly/markets/info/:market_id",poly_market_info)
	r.GET("/black/poly/markets/volume/liquidity/:market_id/:init_ts/:fin_ts/:interval_secs",poly_market_liquidity_periods)
	r.GET("/black/poly/markets/redemptions/:market_id",poly_market_payout_redemptions)
	r.GET("/black/poly/markets/statistics/:market_id",poly_market_stats)
	r.GET("/black/poly/stats/global_liquidity/:init_ts/:fin_ts/:interval_secs",poly_liq_hist_global)
	r.GET("/black/poly/markets/erc1155/:market_id/:offset/:limit",poly_market_erc1155_transfers)
	r.GET("/black/poly/markets/userlist/:market_id",poly_user_list)
	r.GET("/black/poly/markets/traderops/:market_id/:user_aid/:offset/:limit",poly_market_trader_operations)
	r.GET("/black/poly/markets/funderops/:market_id/:user_aid/:offset/:limit",poly_market_funder_operations)
	r.GET("/black/poly/markets/open_positions/:market_id",poly_market_open_positions)
	r.GET("/black/poly/markets/open_interest_history/:market_id/:offset/:limit",poly_market_open_interest_history)
	r.GET("/black/poly/user/open_positions/:user_aid",poly_market_user_open_positions)
	r.GET("/black/poly/user/info/:user",poly_user_info)
	r.GET("/black/poly/user/traded_markets/:user",poly_user_traded_markets)
	r.GET("/black/poly/markets/funder/share_ratio/:market_id",poly_market_funder_share_ratio)
	r.GET("/black/poly/markets/list/:status",poly_markets_listing)
	r.GET("/black/poly/markets/list/:status/:sort",poly_markets_listing)
	r.GET("/black/poly/markets/list",poly_markets_listing)
	r.GET("/black/poly/topusers.html",poly_top_users)
	r.GET("/black/poly/categories/list",poly_market_categories)
	r.GET("/black/poly/search",poly_market_search)
}
func set_routing_randomwalk(r *gin.Engine) {
	r.GET("/black/rwalk",rwalk_index_page)
	r.GET("/black/rwalk/",rwalk_index_page)
	r.GET("/black/rwalk/current_offers/:rwalk_addr/:market_addr/:order_by",rwalk_current_offers)
	r.GET("/black/rwalk/floor_price/:rwalk_addr/:market_addr",rwalk_floor_price)
	r.GET("/black/rwalk/tokens/list/sequential/:rwalk_addr",rwalk_token_list_seq)
	r.GET("/black/rwalk/tokens/list/sequential/:rwalk_addr/:offset/:limit",rwalk_token_list_seq)
	r.GET("/black/rwalk/tokens/list/by_period/:rwalk_addr/:init_ts/:fin_ts",rwalk_token_list_period)
	r.GET("/black/rwalk/tokens/history/:token_id/:rwalk_addr",rwalk_token_history)
	r.GET("/black/rwalk/tokens/history/:token_id/:rwalk_addr/:offest/:limit",rwalk_token_history)
	r.GET("/black/rwalk/tokens/name_changes/:token_id",rwalk_token_name_history)
	r.GET("/black/rwalk/tokens/info/:rwalk_addr/:token_id",rwalk_token_info)
	r.GET("/black/rwalk/tokens/by_user/:user_aid",rwalk_tokens_by_user)
	r.GET("/black/rwalk/trading/history/:market_addr/:offset/:limit",rwalk_trading_history)
	r.GET("/black/rwalk/trading/history/:market_addr",rwalk_trading_history)
	r.GET("/black/rwalk/trading/by_user/:user_aid",rwalk_trading_history_by_user)
	r.GET("/black/rwalk/trading/sales/:market_addr",rwalk_sale_history)
	r.GET("/black/rwalk/statistics/by_token/:rwalk_addr",rwalk_token_stats)
	r.GET("/black/rwalk/statistics/by_market/:market_addr",rwalk_market_stats)
	r.GET("/black/rwalk/statistics/trading_volume/:market_addr/:init_ts/:fin_ts/:interval_secs",rwalk_trading_volume_by_period)
	r.GET("/black/rwalk/statistics/top_users",rwalk_top_users)
	r.GET("/black/rwalk/statistics/mint_intervals/:rwalk_addr",rwalk_mint_intervals)
	r.GET("/black/rwalk/statistics/withdrawal_chart/:rwalk_addr",rwalk_withdrawal_chart)
	r.GET("/black/rwalk/statistics/floor_price/:market_addr/:rwalk_addr/:init_ts/:fin_ts/:interval_secs",rwalk_floor_price_over_time)
	r.GET("/black/rwalk/user/info/:user_aid/:rwalk_addr",rwalk_user_info)
	r.GET("/black/rwalk/download_mints/:rwalk_addr",rwalk_token_csv_export)
	r.GET("/black/rwalk/mint_report",rwalk_mint_report)
}
func set_routing_statistics(r *gin.Engine) {

	r.GET("/black/statistics/main_net/:init_ts/:fin_ts",stats_main_statistics_main_net)
	r.GET("/black/statistics/arbitrum/:init_ts/:fin_ts",stats_main_statistics_arbitrum)
	r.GET("/black/statistics/",stats_index_page)
}
func set_routing_black_templates(r *gin.Engine) {

	set_routing_augur_v2(r)
	set_routing_augur_foundry(r)
	set_routing_uniswap(r)
	set_routing_balancer(r)
	set_routing_ens(r)
	set_routing_augur_amm(r)
	set_routing_polymarket(r)
	set_routing_randomwalk(r)
	set_routing_statistics(r)

	r.GET("/black/ethusd/",show_ethusd_price)

}
