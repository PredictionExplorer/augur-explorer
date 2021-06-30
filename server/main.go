package main

import (
	"log"
	"os"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/autotls"
	"github.com/ethereum/go-ethereum/ethclient"

	"golang.org/x/crypto/acme/autocert"

	. "github.com/PredictionExplorer/augur-explorer/primitives"
	. "github.com/PredictionExplorer/augur-explorer/amm"
)
var (
	RPC_URL = os.Getenv("AUGUR_ETH_NODE_RPC_URL")
	rpcclient *ethclient.Client

	augur_srv *AugurServer

	// contracts
	all_contracts map[string]interface{}

	caddrs *ContractAddresses

	ctrct_dai_token *DAICash
	ctrct_rep_token *RepTok

	REP_ETH_UNISWAP_PAIR_ADDR string = "0x8979A3Ef9D540480342AC0F56e9D4c88807b1CBa"
	// thes variables should be removed on the next code reorg task
	market_order_id int64 = 0
	fill_order_id int64 = 0

	Error   *log.Logger
	Info    *log.Logger

	amm_constants		AMM_Constants
	amm_contracts		AA_ContractAddrs
)
func initialize() {

	caddrs_obj,err := augur_srv.db_augur.Get_contract_addresses()
	if err!=nil {
		Fatalf("Can't find contract addresses in 'contract_addresses' table")
	}
	caddrs=&caddrs_obj

	amm_constants = Load_amm_constants("./amm_constants")
	if augur_srv.db_matic != nil {
//		amm_contracts = augur_srv.db_matic.Get_arbitrum_augur_contract_addresses()
	}
}
func secure_https(r http.Handler) {
	autotls.Run(r, "localhost")
}
func main() {


	augur_srv = create_augur_server()

	initialize()


	if len(RPC_URL) == 0 {
		fmt.Printf("Configuration error: RPC URL of Ethereum node is not set."+
			"Calls to contracts are disabled. " +
			" Please set AUGUR_ETH_NODE_RPC environment variable")
	} else {
		var err error
		rpcclient, err = ethclient.Dial(RPC_URL)
		if err != nil {
			log.Fatal(err)
		}
		// init contracts
		fmt.Printf("init DAI contract with addr %v\n",caddrs.Dai.String())
		ctrct_dai_token,err = NewDAICash(caddrs.Dai,rpcclient)
		if err != nil {
			Fatalf("Couldn't initialize DAI Cash contract: %v\n",err)
		}

		fmt.Printf("init REP contract with addr %v\n",caddrs.Reputation.String())
		ctrct_rep_token,err = NewRepTok(caddrs.Reputation,rpcclient)
		if err != nil {
			Fatalf("Couldn't initialize Rep Token contract: %v\n",err)
		}
	}

	port_plain := os.Getenv("AUGUR_HTTP_PORT")

	r := gin.New()
	//r.RedirectTrailingSlash=false
	//r.RedirectFixedPath = false
	r.LoadHTMLGlob("html/templates/*html")

	r.Use(gin.Logger())

	r.Static("/ui","/home/frontend/pe-ui/build")
	r.StaticFile("/index.html", "/home/frontend/pe-ui/build/index.html")
	r.StaticFile("/index.htm", "/home/frontend/pe-ui/build/index.html")
	r.StaticFile("/INDEX.HTM", "/home/frontend/pe-ui/build/index.html")
	r.StaticFile("/favicon.ico", "/home/frontend/pe-ui/build/favicon.ico")
	r.StaticFile("/sw.js", "/home/frontend/pe-ui/build/sw.js")
	r.StaticFile("/sw.js.gz", "/home/frontend/pe-ui/build/sw.js.gz")
	r.StaticFile("/","/home/frontend/pe-ui/build/index.html")

	// Main HTML
	r.GET("/black/", main_page)
	r.GET("/black/index.html", main_page)
	r.GET("/black/index.htm", main_page)
	r.GET("/black/INDEX.HTM", main_page)
	// Old version of the site with black templates
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
	r.GET("/black/user_wr_transfers/:user/:wrapper",user_wrapped_token_transfers)
	r.GET("/black/user_uswaps/:user",user_uniswap_swaps)
	r.GET("/black/user_bswaps/:user",user_balancer_swaps)
	r.GET("/black/user_ens_names/:user",user_ens_names)
	r.GET("/black/user_rep_pl/:user",user_rep_profit_loss)
	r.GET("/black/statement/:addr",account_statement)
	r.GET("/black/oohist/:addr",open_order_history)
	r.GET("/black/pehist/:market/:outcome", price_estimate_history)
	r.GET("/black/wrapped/:market",wrapped_tokens)
	r.GET("/black/wr_transfers/:address",wrapped_token_transfers)
	r.GET("/black/pool_swaps/:address",pool_swaps)
	r.GET("/black/stbc/:market",sharetoken_balance_changes)
	r.GET("/black/mkt_uniswaps/:market",market_uniswap_pairs)
	r.GET("/black/uniswap_swaps/:address",uniswap_swaps)
	r.GET("/black/text_search",do_text_search)
	r.GET("/black/text_search_form",show_text_search_form)
	r.GET("/black/pool_swap_price/:pool_aid/:token1_aid/:token2_aid/:init_ts/:fin_ts/:interval_secs",show_pool_swap_prices)
	r.GET("/black/upair_swap_price/:pair_aid/:inverse/:init_ts/:fin_ts/:interval_secs",show_upair_swap_prices)
	r.GET("/black/uni_swap/:id",show_single_uniswap_swap)
	r.GET("/black/bal_swap/:id",show_single_balancer_swap)
	r.GET("/black/wrtok_info/:address",wrapped_token_info)
	r.GET("/black/pool_slippage/:pool",show_pool_slippage)
	r.GET("/black/uniswap_slippage/:pair",show_uniswap_slippage)
	r.GET("/black/rt_uniswap_slippage/:pair",rt_show_uniswap_slippage)
	r.GET("/black/ethusd/",show_ethusd_price)
	r.GET("/black/whats_new_augur/",whats_new_in_augur)
	r.GET("/black/node_text_data/:node",show_node_text_data)
	r.GET("/black/augur_foundry",show_augur_foundry_contracts)
	r.GET("/black/reports_table/:market",show_reporting_table)
	r.GET("/black/noshow_bond",augur_noshow_bond_prices)
	r.GET("/black/validity_bond",augur_validity_bond_prices)
	r.GET("/black/ens_name_info/:fqdn",ens_name_info)
	r.GET("/black/aa/pools",arbitrum_augur_pools)
	r.GET("/black/arbitrum/markets/sports/:status/:sort",arbitrum_markets_sports)
	r.GET("/black/arbitrum/liquidity/:factory_aid/:market_id/:offset/:limit",arbitrum_liquidity_changed)
	r.GET("/black/arbitrum/swaps/:contract_aid/:market_id/:offset/:limit",arbitrum_shares_swapped)
	r.GET("/black/arbitrum/user/swaps/:user/:offset/:limit",amm_user_swaps)
	r.GET("/black/arbitrum/user/liquidity/:user/:offset/:limit",amm_user_liquidity)
	r.GET("/black/arbitrum/markets/info/:contract_aid/:market_id",arbitrum_market_info)

	r.Static("/black/imgs", "./html/imgs")
	r.Static("/black/res", "./html/res")			// resources (static)
	// API calls for the new FrontEnd
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
	r.GET("/api/user_uniswaps/:user/:offset/:limit",a1_user_uniswap_swaps)
	r.GET("/api/user_balswaps/:user/:offset/:limit",a1_user_balancer_swaps)
	r.GET("/api/user_ens_names/:user/:offset/:limit",a1_user_ens_names)
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
	r.GET("/api/wrapped_tokens/:market",a1_wrapped_tokens)
	r.GET("/api/wr_transfers/:address/:offset/:limit",a1_wrapped_token_transfers)
	r.GET("/api/user_wr_transfers/:user/:wrapper/:offset/:limit",a1_user_wrapped_token_transfers)
	r.GET("/api/wr_vol/:address/:init_ts/:fin_ts/:interval_secs",a1_wrapped_token_volume)
	r.GET("/api/pool_swaps/:address/:offset/:limit",a1_pool_swaps)
	r.GET("/api/mkt_pool_vol/:market/:outcome/:init_ts/:fin_ts/:interval_secs",a1_balancer_volume)
	r.GET("/api/mkt_stbc/:market/:offset/:limit",a1_market_share_token_balance_changes)
	r.GET("/api/mkt_uniswaps/:market",a1_market_uniswap_pairs)
	r.GET("/api/uniswap_swaps/:address/:offset/:limit",a1_uniswap_pair_swaps)
	r.GET("/api/mkt_uniswap_vol/:market/:outcome/:init_ts/:fin_ts/:interval_secs",a1_uniswap_volume)
	r.GET("/api/search",a1_search)
	r.GET("/api/categories",a1_categories)
	r.GET("/api/pool_price_hist/:pool/:token1/:token2/:init_ts/:fin_ts/:interval_secs",a1_pool_price_history)
	r.GET("/api/upair_price_hist/:pair/:inverse/:init_ts/:fin_ts/:interval_secs",a1_upair_price_history)
	r.GET("/api/uni_swap/:id",a1_single_uniswap_swap)
	r.GET("/api/bal_swap/:id",a1_single_balancer_swap)
	r.GET("/api/bal_calc_slip/:pool/:tok_in/:tok_out/:amount",a1_balancer_calculate_slippage)
	r.GET("/api/pool_slippage/:pool",a1_pool_slippage)
	r.GET("/api/uni_calc_slip/:pair/:tok_in/:amount",a1_uniswap_calculate_slippage)
	r.GET("/api/uniswap_slippage/:pair",a1_uniswap_slippage)
	r.GET("/api/wshtok_balances/:user",a1_wrapped_shtoken_balances)
	r.GET("/api/rlookup/:address",a1_ens_reverse_lookup)
	r.GET("/api/whats_new_augur/:code",a1_whats_new_augur)
	r.GET("/api/node_text_data/:node",a1_node_text_key_value_pairs)
	r.GET("/api/augur_foundry",a1_augur_foundry_contracts)
	r.GET("/api/tx/:hash",a1_transaction_info)
	r.GET("/api/block/:block_num",a1_block_info)
	r.GET("/api/reporting_table/:market",a1_reporting_table)
	r.GET("/api/user_rep_pl/:user",a1_user_rep_profit_loss)
	r.GET("/api/noshow_bond",a1_noshow_bond_prices)
	r.GET("/api/validity_bond",a1_validity_bond_prices)
	r.GET("/api/ens_name_info/:fqdn",a1_ens_name_info)
	r.GET("/api/ens_lookup/:user",a1_ens_name_lookup)
	r.GET("/api/aa/pools",a1_arbitrum_augur_pools)
	r.GET("/api/arbitrum/markets/sports/:status/:sort/:offset/:limit",a1_arbitrum_markets_sports)
	r.GET("/api/arbitrum/markets/info/sports/:factory_aid/:market_id",a1_arbitrum_market_info_sports)
	r.GET("/api/arbitrum/liquidity/:factory_aid/:market_id/:offset/:limit",a1_arbitrum_liquidity_changed)
	r.GET("/api/arbitrum/swaps/:contract_aid/:market_id/:offset/:limit",a1_arbitrum_shares_swapped)
	r.GET("/api/arbitrum/user/swaps/:user/:offset/:limit",a1_amm_user_swaps)
	r.GET("/api/arbitrum/user/liquidity/:user/:offset/:limit",a1_amm_user_liquidity)

	m := autocert.Manager{
		Prompt:     autocert.AcceptTOS,
		HostPolicy: autocert.HostWhitelist("api.predictionexplorer.com"),
		Cache:      autocert.DirCache(os.Getenv("HOME")+".tls-autocert-cache"),
	}
	_ = m
	// Listen and serve on defined port
	log.Printf("Listening on port %s", port_plain)

	go func() {
		log.Printf("%v",autotls.Run(r, "api.predictionexplorer.com"))
	}()
	if len(port_plain) > 0 {
		go func() {
			r.Run(":" + port_plain)
		}()
	}
	select{} // infinite suspend for the main go-routine
}
