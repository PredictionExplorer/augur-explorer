package main
import (
	"os"
	"fmt"
	"log"
	"time"
	"bytes"
	"strconv"
	"net/http"
	"encoding/hex"
	"github.com/gin-gonic/gin"
	"math/big"
	"context"
	"strings"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/wealdtech/go-ens/v3"

	. "github.com/PredictionExplorer/augur-explorer/primitives"
	. "github.com/PredictionExplorer/augur-explorer/dbs"
)
func categories(c *gin.Context) {
	blknum,_ := augur_srv.db_augur.Get_last_block_num()
	categories := augur_srv.db_augur.Get_categories()
	c.HTML(http.StatusOK, "categories.html", gin.H{
			"title": "Augur Market Categories",
			"block_num" : blknum,
			"Categories" : categories,
	})
}
func statistics(c *gin.Context) {

	stats := augur_srv.db_augur.Get_main_stats()
	cash_flow_entries := augur_srv.db_augur.Get_cash_flow()
	gas_usage := augur_srv.db_augur.Get_gas_usage_global()
	uniq_addr_entries := augur_srv.db_augur.Get_unique_addresses()
	cash_flow_data := build_js_cash_flow_data(&cash_flow_entries)
	uniq_addrs_data := build_js_uniq_addrs(&uniq_addr_entries)
	// Gas Used
	gas_usage_trading := build_js_global_gas_usage_data(&gas_usage,0)
	gas_usage_reporting := build_js_global_gas_usage_data(&gas_usage,1)
	gas_usage_markets := build_js_global_gas_usage_data(&gas_usage,2)
	gas_usage_total := build_js_global_gas_usage_data(&gas_usage,3)
	// Transaction Cost
	tx_fees_trading := build_js_global_gas_usage_data(&gas_usage,4)
	tx_fees_reporting := build_js_global_gas_usage_data(&gas_usage,5)
	tx_fees_markets := build_js_global_gas_usage_data(&gas_usage,6)
	tx_fees_total := build_js_global_gas_usage_data(&gas_usage,7)

	c.HTML(http.StatusOK, "augur_v2/statistics.html", gin.H{
			"title": "Augur Market Statistics",
			"MainStats" : stats,
			"CashFlowData" : cash_flow_data,
			"UniqAddrsData" : uniq_addrs_data,
			"GasUsageTrading" : gas_usage_trading,
			"GasUsageReporting" : gas_usage_reporting,
			"GasUsageMarkets" : gas_usage_markets,
			"GasUsageTotal" : gas_usage_total,
			"TxFeesTrading" : tx_fees_trading,
			"TxFeesReporting" : tx_fees_reporting,
			"TxFeesMarkets" : tx_fees_markets,
			"TxFeesTotal" : tx_fees_total,
	})
}
func complete_and_output_market_info(c *gin.Context,json_output bool,minfo InfoMarket) {
	trades := augur_srv.db_augur.Get_mkt_trades(minfo.MktAddr,10000000)
	outcome_vols,_ := augur_srv.db_augur.Get_outcome_volumes(minfo.MktAddr,minfo.MktAid,0,minfo.LowPriceLimit)
	price_estimates := augur_srv.db_augur.Get_price_estimates(minfo.MktAid,outcome_vols,minfo.LowPriceLimit)
	reports := augur_srv.db_augur.Get_market_reports(minfo.MktAid,0)
	price_history := augur_srv.db_augur.Get_full_price_history(minfo.MktAddr,minfo.MktAid,minfo.LowPriceLimit)
	balancer_pools := augur_srv.db_augur.Get_market_balancer_pools(minfo.MktAid)
	uniswap_pairs := augur_srv.db_augur.Get_market_uniswap_pairs(minfo.MktAid)
	wrappers := augur_srv.db_augur.Get_wrapped_tokens_for_market(minfo.MktAid)

	if json_output {
		c.JSON(http.StatusOK,gin.H{
			"Trades" : trades,
			"Reports" : reports,
			"Market": minfo ,
			"OutcomeVols" : outcome_vols,
			"PriceHistory" : price_history,
			"PriceEstimates" : price_estimates,
			"BalancerPools" : balancer_pools,
			"UniswapPairs":  uniswap_pairs,
			"WrappedContracts": wrappers,
		})
	} else {
		c.HTML(http.StatusOK, "augur_v2/market_info.html", gin.H{
			"title": "DISCONTINUED",
			"Trades" : trades,
			"Reports" : reports,
			"Market": minfo ,
			"OutcomeVols" : outcome_vols,
			"PriceHistory" : price_history,
			"PriceEstimates" : price_estimates,
			"BalancerPools" : balancer_pools,
			"UniswapPairs" : uniswap_pairs,
		})
	}
}
func explorer(c *gin.Context) {
	blknum,res := augur_srv.db_augur.Get_last_block_num()
	_ = res
	c.HTML(http.StatusOK, "augur_v2/explorer.html", gin.H{
			"title": "Augur Event Explorer",
			"block_num" : blknum,
	})
}
func market_info(c *gin.Context) {

	market := c.Param("market")
	market_addr,valid:=is_address_valid(c,false,market)
	if !valid {
		return
	}
	market_info,err := augur_srv.db_augur.Get_market_info(market_addr,0,false)
	if err != nil {
		show_market_not_found_error(c,false,&market_addr)
		return
	}
	complete_and_output_market_info(c,false,market_info)
}
func full_trade_list(c *gin.Context) {

	market := c.Param("market")
	market_addr,valid := is_address_valid(c,false,market)
	if !valid {
		return
	}
	market_info,err := augur_srv.db_augur.Get_market_info(market_addr,0,false)
	if err != nil {
		show_market_not_found_error(c,false,&market_addr)
		return
	}
	trades := augur_srv.db_augur.Get_mkt_trades(market_addr,0)
	c.HTML(http.StatusOK, "augur_v2/full_trade_list.html", gin.H{
			"title": "Trades for market",
			"Trades" : trades,
			"Market": market_info,
	})
}
func market_depth(c *gin.Context) {

	// Market Depth Info: https://en.wikipedia.org/wiki/Order_book_(trading)
	market := c.Param("market")
	market_addr,valid := is_address_valid(c,false,market)
	if !valid {
		return
	}
	p_outcome := c.Param("outcome")
	var outcome int
	if len(p_outcome) > 0 {
		p, err := strconv.Atoi(p_outcome)
		if err != nil {
			c.HTML(http.StatusBadRequest, "error.html", gin.H{
				"title": "Augur Markets: Error",
				"ErrDescr": "Can't parse outcome",
			})
			return
		}
		outcome = int(p)
	} else {
		c.HTML(http.StatusBadRequest, "error.html", gin.H{
			"title": "Augur Markets: Error",
			"ErrDescr": "Can't parse outcome",
		})
		return
	}

	market_info,err := augur_srv.db_augur.Get_market_info(market_addr,outcome,true)
	if err != nil {
		show_market_not_found_error(c,false,&market_addr)
		return
	}
	mdepth,last_oo_id := augur_srv.db_augur.Get_mkt_depth(market_info.MktAid,outcome)
	num_orders:=len(mdepth.Bids) + len(mdepth.Asks)
	js_bid_data,js_ask_data := build_js_data_obj(mdepth)
	c.HTML(http.StatusOK, "augur_v2/market_depth.html", gin.H{
			"title": "Market Depth",
			"Market": market_info,
			"LastOOID": last_oo_id,
			"NumOrders" : num_orders,
			"Bids": mdepth.Bids,
			"Asks": mdepth.Asks,
			"JSAskData": js_ask_data,
			"JSBidData": js_bid_data,
			"OutcomeIdx" : outcome,
	})
}
func market_depth_ajax(c *gin.Context) {

	p_outcome := c.Param("outcome")
	var outcome int64
	if len(p_outcome) > 0 {
		var success bool
		outcome,success = parse_int_from_remote_or_error(c,true,&p_outcome)
		if !success {
			return
		}
	} else {
		respond_error(c,"No outcome provided")
		return
	}

	p_market_aid := c.Param("market_aid")
	var market_aid int64
	if len(p_market_aid) > 0 {
		var success bool
		market_aid,success = parse_int_from_remote_or_error(c,true,&p_market_aid)
		if !success {
			return
		}
	} else {
		respond_error(c,"No outcome provided")
		return
	}
	mdepth,last_oo_id := augur_srv.db_augur.Get_mkt_depth(market_aid,int(outcome))
	js_bid_data,js_ask_data := build_js_data_obj(mdepth)
	c.JSON(http.StatusOK,gin.H{
		"bids":js_bid_data,
		"asks":js_ask_data,
		"LastOOID":last_oo_id,
	})
}
func market_price_history(c *gin.Context) {

	market := c.Param("market")
	market_addr,valid := is_address_valid(c,false,market)
	if !valid {
		return
	}
	p_outcome := c.Param("outcome")
	var outcome int
	if len(p_outcome) > 0 {
		p, err := strconv.Atoi(p_outcome)
		if err != nil {
			c.HTML(http.StatusBadRequest, "error.html", gin.H{
				"title": "Augur Markets: Error",
				"ErrDescr": "Can't parse outcome",
			})
			return
		}
		outcome = int(p)
	} else {
		c.HTML(http.StatusBadRequest, "error.html", gin.H{
			"title": "Augur Markets: Error",
			"ErrDescr": "Can't parse outcome",
		})
		return
	}

	market_info,err := augur_srv.db_augur.Get_market_info(market_addr,outcome,true)
	if err != nil {
		show_market_not_found_error(c,false,&market_addr)
		return
	}
	mkt_price_hist := augur_srv.db_augur.Get_price_history_for_outcome(market_info.MktAid,outcome,market_info.LowPriceLimit)
	js_price_history := build_js_price_history(&mkt_price_hist)
	fmt.Printf("js price history = %v\n",js_price_history)
	c.HTML(http.StatusOK, "augur_v2/price_history.html", gin.H{
			"title": "Market Price History",
			"Market": market_info,
			"Prices": mkt_price_hist,
			"JSPriceData": js_price_history,
	})
}
func serve_user_info_page(c *gin.Context,addr string) {

	eoa_aid,err := augur_srv.db_augur.Nonfatal_lookup_address_id(addr)
	if err == nil {
		user_info,err := augur_srv.db_augur.Get_user_info(eoa_aid)
		if err == nil {
			pl_entries := augur_srv.db_augur.Get_profit_loss(eoa_aid)
			open_pos_entries := augur_srv.db_augur.Get_open_positions(eoa_aid)
			js_pl_data := build_js_profit_loss_history(&pl_entries)
			js_open_pos_data := build_js_open_positions(&open_pos_entries)
			user_reports := augur_srv.db_augur.Get_user_reports(eoa_aid,DEFAULT_USER_REPORTS_LIMIT)
			user_active_markets := augur_srv.db_augur.Get_traded_markets_for_user(eoa_aid,1)
			var has_active_markets bool = false
			if len(user_active_markets) > 0 {
				has_active_markets = true
			}
			open_orders := augur_srv.db_augur.Get_user_open_orders(user_info.Aid)
			gas_spent,_ := augur_srv.db_augur.Get_gas_spent_for_user(eoa_aid)
			shtoken_balances := augur_srv.db_augur.Get_wrapped_shtoken_balances(eoa_aid)
			active_names,active_total_rows := augur_srv.db_augur.Get_user_ens_names_active(eoa_aid,0,1000000)
			inactive_names,inactive_total_rows := augur_srv.db_augur.Get_user_ens_names_inactive(eoa_aid,0,1000000)

			c.HTML(http.StatusOK, "augur_v2/user_info.html", gin.H{
				"title": "User "+addr,
				"user_addr": addr,
				"UserInfo" : user_info,
				"PLEntries" : pl_entries,
				"JSPLData" : js_pl_data,
				"JSOpenPosData" : js_open_pos_data,
				"OpenOrders": open_orders,
				"UserReports" : user_reports,
				"UserActiveMarkets" : user_active_markets,
				"HasActiveMarkets" : has_active_markets,
				"GasSpent" : gas_spent,
				"ShtokBalances" : shtoken_balances,
				"ENS_Names_Active" : active_names,
				"ENS_Names_Inactive" : inactive_names,
				"Total_ENS_Names_Active": active_total_rows,
				"Total_ENS_Names_History": inactive_total_rows,
			})
		} else {
			c.HTML(http.StatusOK, "user_not_found.html", gin.H{
				"title": "User "+addr,
				"user_addr": addr,
			})
		}
	} else {
		c.HTML(http.StatusOK, "user_not_found.html", gin.H{
			"title": "User "+addr,
			"user_addr": addr,
		})
	}
}
func serve_tx_info_page(c *gin.Context,tx_hash string) {

	tx_info,err := augur_srv.db_augur.Get_transaction(tx_hash)
	if err == nil {
		c.HTML(http.StatusOK, "augur_v2/tx_info.html", gin.H{
			"title": "Transaction " + tx_hash,
			"TxInfo" : tx_info,
		})
	} else {
		c.HTML(http.StatusOK, "augur_v2/tx_not_found.html", gin.H{
			"title": "Transaction "+tx_hash,
			"tx_hash": tx_hash,
		})
	}
}
func serve_user_funds_v2(c *gin.Context,addr *string) {
	// the input address must be EOA, from that we can get Wallet addr
	var (
		dai_balance float64 = 0.0
		rep_balance float64 = 0.0
		rep_balance_usd float64 = 0.0
		eth_balance float64 = 0.0
		eth_balance_usd float64 = 0.0

	)
	var status_code int = 0
	var error_text  string = ""
	var total_usd float64 = 0
	eoa_aid,err := augur_srv.db_augur.Nonfatal_lookup_address_id(*addr)
	if err != nil {
		error_text = "Address lookup failed"
	}
	if eoa_aid > 0 {
		addr := common.HexToAddress(*addr)
		dai_balance = get_token_balance(0,&addr)
		rep_balance = get_token_balance(1,&addr)
		eth_balance = get_eth_balance(&addr)
		status_code = 1
	}

	rep_in_eth,err :=get_REP_token_price_in_ETH()
	if err == nil {
		ethusd,err := augur_srv.db_augur.Get_last_ethusd_price()
		if err == nil {
			rep_balance_usd = rep_balance / rep_in_eth * ethusd
			total_usd = dai_balance + eth_balance * ethusd + rep_balance_usd
			eth_balance_usd = eth_balance * ethusd
		}
	} else {
		error_text = fmt.Sprintf("Error at checking REP price: %v\n",err.Error())
	}
	scode := http.StatusOK
	if status_code == 0 {
		status_code = http.StatusBadRequest
	}
	c.JSON(scode, gin.H{
		"status": status_code,
		"error": error_text,
		"Eth": fmt.Sprintf("%v",eth_balance),
		"EthInUsd": fmt.Sprintf("%v",eth_balance_usd),
		"DAI": fmt.Sprintf("%v",dai_balance),
		"REP": fmt.Sprintf("%v",rep_balance),
		"REPUSD" : fmt.Sprintf("%v",rep_balance_usd),
		"TotalUSDAcctValue" : total_usd,
	})
}
func serve_user_funds_v1(c *gin.Context,addr common.Address) {
	// the input address must be EOA, from that we can get Wallet addr
	// Note: this request is becoming obsolete, and will be removed later

	var wallet_aid int64 = 0
	eoa_aid,err := augur_srv.db_augur.Nonfatal_lookup_address_id(addr.String())
	if err == nil {
		wallet_aid,_ = augur_srv.db_augur.Lookup_wallet_aid(eoa_aid)
	} else {
		c.JSON(http.StatusOK,gin.H{
			"eoa_eth":0,"wallet_eth":0,"eoa_dai":0,"wallet_dai":0,"eoa_rep":0,"wallet_rep":0,
		})
		return
	}
	eoa_dai_balance := get_token_balance(0,&addr)
	eoa_rep_balance := get_token_balance(1,&addr)
	eoa_eth_balance := get_eth_balance(&addr)

	var wallet_dai_balance float64 = 0.0
	var wallet_rep_balance float64 = 0.0
	var wallet_eth_balance float64 = 0.0

	if wallet_aid != 0 {
		wallet_addr,err := augur_srv.db_augur.Lookup_address(wallet_aid)
		if err == nil {
			waddr := common.HexToAddress(wallet_addr)
			wallet_dai_balance = get_token_balance(0,&waddr)
			wallet_rep_balance = get_token_balance(1,&waddr)
			wallet_eth_balance = get_eth_balance(&waddr)
		} else {
			fmt.Printf("address lookup for wallet_aid = %v failed: %v",wallet_aid,err)
		}
	}

	c.JSON(http.StatusOK, gin.H{
			"eoa_eth": fmt.Sprintf("%v",eoa_eth_balance),
			"wallet_eth": fmt.Sprintf("%v",wallet_eth_balance),
			"eoa_dai": fmt.Sprintf("%v",eoa_dai_balance),
			"wallet_dai": fmt.Sprintf("%v",wallet_dai_balance),
			"eoa_rep": fmt.Sprintf("%v",eoa_rep_balance),
			"wallet_rep": fmt.Sprintf("%v",wallet_rep_balance),
	})
}
func search(c *gin.Context) {

	keyword := c.Query("keywords")
	if (len(keyword) == 40) || (len(keyword) == 42) { // address
		if len(keyword) == 42 {	// strip 0x prefix
			keyword = keyword[2:]
		}
		addr_bytes,err := hex.DecodeString(keyword)
		if err == nil {
			addr := common.BytesToAddress(addr_bytes)
			addr_str := addr.String()
			_,err:=augur_srv.db_augur.Nonfatal_lookup_address_id(addr_str)
			if err==nil {
				market_info,err := augur_srv.db_augur.Get_market_info(addr_str,0,false)
				if err == nil {
					complete_and_output_market_info(c,false,market_info)
					return
				}
				serve_user_info_page(c,addr_str)
				return
			} else {
				c.HTML(http.StatusBadRequest, "error.html", gin.H{
					"title": "Augur Markets: Error",
					"ErrDescr": fmt.Sprintf("Address %v not found",addr_str),
				})
				return
			}
		} else {
			c.HTML(http.StatusBadRequest, "error.html", gin.H{
				"title": "Augur Markets: Error",
				"ErrDescr": fmt.Sprintf("Invalid HEX string in address parameter: %v",keyword),
			})
			return
		}
	}
	if (len(keyword) == 64) || (len(keyword) == 66) { // Hash (Tx hash)
		if len(keyword) == 66 {	// strip 0x prefix
			keyword = keyword[2:]
		}
		var hash string
		hash_bytes,err := hex.DecodeString(keyword)
		if err == nil {
			tmp_hash := common.BytesToHash(hash_bytes)
			hash = tmp_hash.String()
		} else {
			c.HTML(http.StatusBadRequest, "error.html", gin.H{
				"title": "Augur Markets: Error",
				"ErrDescr": "Invalid HEX string in hash parameter",
			})
			return
		}
		if augur_srv.db_augur.Tx_exists(hash) {
			serve_tx_info_page(c,hash)
			return
		}
		orders := augur_srv.db_augur.Get_filling_orders_by_hash(hash)
		if len(orders) > 0 {
			output_filling_orders(c,hash,orders)
			return
		}
		c.HTML(http.StatusBadRequest, "error.html", gin.H{
			"title": "Augur Markets: Error",
			"ErrDescr": fmt.Sprintf("Obhect with hash %v wasn't found",hash),
		})
		return
	} else {
		_, err := strconv.Atoi(keyword)
		if err == nil {
			serve_block_info(keyword,c)
		} else {
			c.HTML(http.StatusBadRequest, "error.html", gin.H{
				"title": "Augur Markets: Error",
				"ErrDescr": "Search object not found",
			})
			return
		}
	}
}
func search_v2(c *gin.Context) {

	keyword := c.Query("q")
	sr := execute_search(keyword)
	if !sr.Found {
		c.HTML(http.StatusBadRequest, "error.html", gin.H{
			"title": "Augur Markets: Error",
			"ErrDescr": fmt.Sprintf("Search object %v not found: %v",keyword,sr.ErrStr),
		})
		return
	}
	switch (sr.SRType) {
	case SR_MarketOrders:
		orders:=sr.Object.(*[]OrderInfo)
		output_filling_orders(c,sr.Query,*orders)
	case SR_Address:
		addr := sr.Object.(*common.Address)
		c.HTML(http.StatusBadRequest, "address.html", gin.H{
			"Address" : addr.String(),
		})
	case SR_Hash:
		hash := sr.Object.(*common.Hash)
		c.HTML(http.StatusBadRequest, "hash.html", gin.H{
			"Hash" : hash.String(),
		})
	case SR_Transaction:
		tx_info := sr.Object.(*TxInfo)
		c.HTML(http.StatusOK, "tx_info.html", gin.H{
			"title": "Transaction " + keyword,
			"TxInfo" : tx_info,
		})
	case SR_Block:
		block_info := sr.Object.(*BlockInfo)
		c.HTML(http.StatusOK, "block/block_info.html", gin.H{
			"BlockInfo" : block_info,
		})
	case SR_UserInfo:
		user_info := sr.Object.(*UserInfo)
		serve_user_info_page(c,user_info.Addr)
	case SR_ShareTokenWrapper:
		winfo := sr.Object.(*ERC20ShTokContract)
		c.HTML(http.StatusOK, "wrapped_shtok_info.html", gin.H{
			"WrapperInfo" : winfo,
		})
	case SR_AugurMarketInfo:
		minfo:=sr.Object.(*InfoMarket)
		complete_and_output_market_info(c,false,*minfo)
		return
	case SR_BalancerPool:
		pool_obj := sr.Object.(*BalancerPoolInfo)
		pool_info,_ := augur_srv.db_augur.Get_pool_info(pool_obj.PoolAid)
		swaps := augur_srv.db_augur.Get_pool_swaps(pool_info.PoolAid,0,200)
		c.HTML(http.StatusOK, "balancer/pool_swaps.html", gin.H{
				"PoolInfo" : pool_info,
				"PoolSwaps" : swaps,
		})
	case SR_UniswapPair:
		now_ts := time.Now().Unix()
		past_ts := now_ts - 100 * 3600 * 24
		pair_info := sr.Object.(*MarketUPair)//.storage.Get_uniswap_pair_info(aid)
		swaps := augur_srv.db_augur.Get_uniswap_swaps(pair_info.PairAid,0,200)
		c.HTML(http.StatusOK, "uniswap/swaps.html", gin.H{
				"PairInfo" : pair_info,
				"PairSwaps" : swaps,
				"SampleFinTs" : now_ts,
				"SampleInitTs" : past_ts,
		})
	case SR_Unknown:
		fallthrough
	default:
		c.HTML(http.StatusBadRequest, "error.html", gin.H{
			"title": "Augur Markets: Error",
			"ErrDescr": fmt.Sprintf("Search object %v found but no HTML template for output found for object type=%v",keyword,sr.SRType),
		})
	}
}
func execute_search(keyword string) SearchResultObject {

	if len(keyword) == 0 {
		var iface interface{}
		return SearchResultObject {
			SRType:			SR_Block,
			Found:			false,
			ErrStr:			"keyword is empty",
			Object:			iface,
		}
	}
	idx := strings.Index(keyword, ":") // if there is a colon, then it is a range of block search
	if idx != -1 {
		block_range := strings.Split(keyword,":")
		if len(block_range) !=2  {
			var iface interface{}
			return SearchResultObject {
				SRType:			SR_Block,
				Found:			false,
				ErrStr:			"Invalid block range, two blocks are needed",
				Query:			keyword,
				Object:			iface,
			}
		}
		block_num_from,err := strconv.Atoi(block_range[0])
		if err != nil {
			var iface interface{}
			return SearchResultObject {
				SRType:		SR_Block,
				Found:		false,
				ErrStr:		fmt.Sprintf("Error in converting first block num: %v",err.Error()),
				Query:		keyword,
				Object:		iface,
			}
		}
		block_num_to,err := strconv.Atoi(block_range[1])
		if err != nil {
			var iface interface{}
			return SearchResultObject {
				SRType:		SR_Block,
				Found:		false,
				ErrStr:		fmt.Sprintf("Error in converting second block num: %v",err.Error()),
				Query:		keyword,
				Object:		iface,
			}
		}
		block_info,err := augur_srv.db_augur.Get_block_info(int64(block_num_from),int64(block_num_to))
		if err != nil {
			var iface interface{}
			return SearchResultObject {
				SRType:		SR_Block,
				Found:		false,
				ErrStr:		err.Error(),
				Query:		keyword,
				Object:		iface,
			}
		}
		var iface interface{}
		iface = &block_info
		return SearchResultObject {
			SRType:		SR_Block,
			Found:		true,
			ErrStr:		"",
			Query:		keyword,
			Object:		iface,
		}
	}
	idx = strings.Index(keyword, " ") // if there is a space, then it is some text to search
	if idx == -1 {
		var hex_data []byte
		var err error
		if has0xPrefix(keyword) {
			hex_data,err = hex.DecodeString(keyword[2:])
		} else {
			hex_data,err = hex.DecodeString(keyword)
		}
		if err == nil {
			// could be: Hash or Address
			if (len(keyword) == 40) || (len(keyword) == 42) { // Address
				addr := common.BytesToAddress(hex_data) // corrects any lower-case input
				addr_str := addr.String()
				aid,err:=augur_srv.db_augur.Nonfatal_lookup_address_id(addr_str)
				if err != nil {
					var iface interface{}
					iface = &addr
					return SearchResultObject {
						SRType:		SR_Address,
						Found:		true,
						ErrStr:		"",
						Query:		keyword,
						Object:		iface,
					}
				}
				market_info,err := augur_srv.db_augur.Get_market_info(addr_str,0,false)
				if err == nil {
					var iface interface{}
					iface = &market_info
					return SearchResultObject {
						SRType:		SR_AugurMarketInfo,
						Found:		true,
						ErrStr:		"",
						Query:		keyword,
						Object:		iface,
					}
				}
				pool_info,err := augur_srv.db_augur.Get_pool_info(aid)
				if err == nil {
					var iface interface{}
					iface = &pool_info
					return SearchResultObject {
						SRType:		SR_BalancerPool,
						Found:		true,
						ErrStr:		"",
						Query:		keyword,
						Object:		iface,
					}
				}
				uniswap_info,err := augur_srv.db_augur.Get_uniswap_pair_info(aid)
				if err == nil {
					var iface interface{}
					iface = &uniswap_info
					return SearchResultObject {
						SRType:		SR_UniswapPair,
						Found:		true,
						ErrStr:		"",
						Query:		keyword,
						Object:		iface,
					}
				}
				af_wrapper,err := augur_srv.db_augur.Get_wrapped_token_info(aid)
				if err == nil {
					var iface interface{}
					iface = &af_wrapper
					return SearchResultObject {
						SRType:		SR_ShareTokenWrapper,
						Found:		true,
						ErrStr:		"",
						Object:		iface,
					}
				}
				user_info,err := augur_srv.db_augur.Get_user_info(aid)
				if err == nil {
					var iface interface{}
					iface= &user_info
					return SearchResultObject {
						SRType:		SR_UserInfo,
						Found:		true,
						ErrStr:		"",
						Query:		keyword,
						Object:		iface,
					}
				}
				var iface interface{}
				iface = &addr
				return SearchResultObject {
					SRType:		SR_Address,
					Found:		true,
					Query:		keyword,
					Object:		iface,
				}
			}
			if (len(keyword) == 64) || (len(keyword) == 66) { // Hash (Tx hash)
				hash := common.BytesToHash(hex_data)	// corrects any lower-case input
				hash_str := hash.String()
				tx_info,err := augur_srv.db_augur.Get_transaction(hash_str)
				if err == nil {
					var iface interface{}
					iface = &tx_info
					return SearchResultObject {
						SRType:		SR_Transaction,
						Found:		true,
						ErrStr:		"",
						Query:		keyword,
						Object:		iface,
					}
				}
				orders := augur_srv.db_augur.Get_filling_orders_by_hash(hash_str)
				if len(orders) > 0 {
					var iface interface{}
					iface = &orders
					return SearchResultObject {
						SRType:		SR_MarketOrders,
						Found:		true,
						ErrStr:		"",
						Query:		keyword,
						Object:		iface,
					}
				}
				block_num,err := augur_srv.db_augur.Get_block_num_by_hash(hash_str)
				if err == nil {
					block_info,err := augur_srv.db_augur.Get_block_info(block_num,block_num)
					if err != nil {
						var iface interface{}
						return SearchResultObject {
							SRType:		SR_Block,
							Found:		false,
							ErrStr:		err.Error(),
							Query:		keyword,
							Object:		iface,
						}
					}
					var iface interface{}
					iface = &block_info
					return SearchResultObject {
						SRType:		SR_Block,
						Found:		true,
						ErrStr:		"",
						Object:		iface,
					}
				}
				var iface interface{}
				iface = &hash
				return SearchResultObject {
					SRType:			SR_Hash,
					Found:			true,
					ErrStr:			"",
					Query:			keyword,
					Object:			iface,
				}
			}
		}
		block_num,err := strconv.Atoi(keyword)
		if err == nil {
			if block_num <= 0 {
				var iface interface{}
				return SearchResultObject {
					SRType:		SR_Block,
					Found:		false,
					ErrStr:		"Given block number is not a positive number",
					Query:		keyword,
					Object:		iface,
				}
			}
			block_info,err := augur_srv.db_augur.Get_block_info(int64(block_num),int64(block_num))
			if err != nil {
				var iface interface{}
				return SearchResultObject {
					SRType:		SR_Block,
					Found:		false,
					ErrStr:		err.Error(),
					Query:		keyword,
					Object:		iface,
				}
			}
			var iface interface{}
			iface = &block_info
			return SearchResultObject {
				SRType:		SR_Block,
				Found:		true,
				ErrStr:		"",
				Query:		keyword,
				Object:		iface,
			}
		}
	}

	search_results := augur_srv.db_augur.Search_keywords_in_markets(keyword)
	var iface interface{}
	iface = &search_results
	return SearchResultObject {
		SRType:		SR_TextSearchResults,
		Found:		true,
		ErrStr:		"",
		Query:		keyword,
		Object:		iface,
	}
}
func output_filling_orders(c *gin.Context,order_hash string,orders []OrderInfo) {
	c.HTML(http.StatusOK, "augur_v2/filling_orders.html", gin.H{
		"title": "Order "+order_hash,
		"OrderHash" : order_hash,
		"FillingOrders" : orders,
	})
}
func order(c *gin.Context) {
	// Order can be queried by OrderId or OrderHash
	// In case of OrderID we return single record
	// In case of OrderHash we return many rows (all orders filling the order for this OrderHash)
	p_order := c.Param("order")

	// Case 1 : query by orderId
	order_id, err := strconv.ParseInt(p_order,10,64)
	if err == nil {
		order,err := augur_srv.db_augur.Get_order_info_by_id(order_id)
		if err != nil {
			c.HTML(http.StatusBadRequest, "error.html", gin.H{
				"title": "Augur Markets: Error",
				"ErrDescr": fmt.Sprintf("Can't find the order with id=%v",order_id),
			})
			return
		}
		c.HTML(http.StatusOK, "order_info.html", gin.H{
			"title": fmt.Sprintf("Order "+order.OrderHash),
			"OrderInfo" : order,
		})
		return
	}
	// Case 2 : query by OrderHash
	order_hash := p_order
	orders := augur_srv.db_augur.Get_filling_orders_by_hash(order_hash)
	output_filling_orders(c,order_hash,orders)
}
func category(c *gin.Context) {

	p_catid:= c.Param("catid")

	cat_id,success := parse_int_from_remote_or_error(c,false,&p_catid)
	if !success {
		return
	}
	cat_markets := augur_srv.db_augur.Get_category_markets(int64(cat_id))
	c.HTML(http.StatusOK, "augur_v2/category_markets.html", gin.H{
			"title": "Category Markets",
			"Markets": cat_markets,
	})
}
func user_info(c *gin.Context) {
	p_addr := c.Param("addr")
	if (len(p_addr) == 40) || (len(p_addr) == 42) { // address
		if len(p_addr) == 42 {	// strip 0x prefix
			p_addr = p_addr[2:]
		}
	} else {
		c.HTML(http.StatusBadRequest, "error.html", gin.H{
			"title": "Augur Markets: Error",
			"ErrDescr": "Invalid length of address parameter",
		})
	}
	addr_bytes,err := hex.DecodeString(p_addr)
	if err == nil {
		addr := common.BytesToAddress(addr_bytes)
		addr_str := addr.String()
		serve_user_info_page(c,addr_str)
	} else {
		c.HTML(http.StatusBadRequest, "error.html", gin.H{
			"title": "Augur Markets: Error",
			"ErrDescr": fmt.Sprintf("Invalid HEX string in address parameter: %v",err),
		})
	}
}
func full_reports(c *gin.Context) {

	p_addr := c.Param("addr")
	if (len(p_addr) == 40) || (len(p_addr) == 42) { // address
		if len(p_addr) == 42 {	// strip 0x prefix
			p_addr = p_addr[2:]
		}
	} else {
		c.HTML(http.StatusBadRequest, "error.html", gin.H{
			"title": "Augur Markets: Error",
			"ErrDescr": "Invalid length of address parameter",
		})
	}
	addr_bytes,err := hex.DecodeString(p_addr)
	if err == nil {
		addr := common.BytesToAddress(addr_bytes)
		addr_str := addr.String()
		aid,err:=augur_srv.db_augur.Nonfatal_lookup_address_id(addr_str)
		if err==nil {
			user_info,err := augur_srv.db_augur.Get_user_info(aid)
			if err!= nil {
				c.HTML(http.StatusBadRequest, "error.html", gin.H{
					"title": "Augur Markets: Error",
					"ErrDescr": fmt.Sprintf("No records found for address: %v",addr_str),
				})
			} else {
				user_reports := augur_srv.db_augur.Get_user_reports(aid,0)
				c.HTML(http.StatusOK, "augur_v2/full_user_reports.html", gin.H{
					"title": fmt.Sprintf("User Reports %v",addr),
					"UserInfo" : user_info,
					"UserReports" : user_reports,
				})
			}
		} else {
			c.HTML(http.StatusBadRequest, "error.html", gin.H{
				"title": "Augur Markets: Error",
				"ErrDescr": fmt.Sprintf("DB error: %v",err),
			})
		}
	} else {
		c.HTML(http.StatusBadRequest, "error.html", gin.H{
			"title": "Augur Markets: Error",
			"ErrDescr": fmt.Sprintf("Invalid HEX string in address parameter: %v",err),
		})
	}
}
func user_markets(c *gin.Context) {

	p_addr := c.Param("addr")
	if (len(p_addr) == 40) || (len(p_addr) == 42) { // address
		if len(p_addr) == 42 {	// strip 0x prefix
			p_addr = p_addr[2:]
		}
	} else {
		c.HTML(http.StatusBadRequest, "error.html", gin.H{
			"title": "Augur Markets: Error",
			"ErrDescr": "Invalid length of address parameter",
		})
	}
	addr_bytes,err := hex.DecodeString(p_addr)
	if err == nil {
		addr := common.BytesToAddress(addr_bytes)
		addr_str := addr.String()
		aid,err:=augur_srv.db_augur.Nonfatal_lookup_address_id(addr_str)
		if err==nil {
			user_info,err := augur_srv.db_augur.Get_user_info(aid)
			if err!= nil {
				c.HTML(http.StatusBadRequest, "error.html", gin.H{
					"title": "Augur Markets: Error",
					"ErrDescr": fmt.Sprintf("No records found for address: %v",addr_str),
				})
			} else {
				user_reports := augur_srv.db_augur.Get_user_markets(aid)
				c.HTML(http.StatusOK, "augur_v2/user_markets.html", gin.H{
					"title": fmt.Sprintf("User Markets %v",addr),
					"UserInfo" : user_info,
					"Markets" : user_reports,
				})
			}
		} else {
			c.HTML(http.StatusBadRequest, "error.html", gin.H{
				"title": "Augur Markets: Error",
				"ErrDescr": fmt.Sprintf("DB error: %v",err),
			})
		}
	} else {
		c.HTML(http.StatusBadRequest, "error.html", gin.H{
			"title": "Augur Markets: Error",
			"ErrDescr": fmt.Sprintf("Invalid HEX string in address parameter: %v",err),
		})
	}
}
func user_deposits_withdrawals(c *gin.Context) {

	p_addr := c.Param("addr")
	if (len(p_addr) == 40) || (len(p_addr) == 42) { // address
		if len(p_addr) == 42 {	// strip 0x prefix
			p_addr = p_addr[2:]
		}
	} else {
		c.HTML(http.StatusBadRequest, "error.html", gin.H{
			"title": "Augur Markets: Error",
			"ErrDescr": "Invalid length of address parameter",
		})
	}
	addr_bytes,err := hex.DecodeString(p_addr)
	if err == nil {
		addr := common.BytesToAddress(addr_bytes)
		addr_str := addr.String()
		aid,err:=augur_srv.db_augur.Nonfatal_lookup_address_id(addr_str)
		if err==nil {
			user_info,err := augur_srv.db_augur.Get_user_info(aid)
			if err!= nil {
				c.HTML(http.StatusBadRequest, "error.html", gin.H{
					"title": "Augur Markets: Error",
					"ErrDescr": fmt.Sprintf("No records found for address: %v",addr_str),
				})
			} else {
				wallet_aid,err := augur_srv.db_augur.Lookup_wallet_aid(aid)
				if err == nil {
					user_deposits_withdrawals := augur_srv.db_augur.Get_deposits_withdrawals(wallet_aid)
					c.HTML(http.StatusOK, "augur_v2/user_deposits_withdrawals.html", gin.H{
						"title": fmt.Sprintf("User %v Deposits/Withdrawals",addr),
						"UserInfo" : user_info,
						"Operations" : user_deposits_withdrawals,
					})
				} else {
					c.HTML(http.StatusBadRequest, "error.html", gin.H{
						"title": "Augur Markets: Error",
						"ErrDescr": fmt.Sprintf("User %v doesn't have a Wallet in Augur",addr.String()),
					})
				}
			}
		} else {
			c.HTML(http.StatusBadRequest, "error.html", gin.H{
				"title": "Augur Markets: Error",
				"ErrDescr": fmt.Sprintf("DB error: %v",err),
			})
		}
	} else {
		c.HTML(http.StatusBadRequest, "error.html", gin.H{
			"title": "Augur Markets: Error",
			"ErrDescr": fmt.Sprintf("Invalid HEX string in address parameter: %v",err),
		})
	}
}
func block_info(c *gin.Context) {

	p_block_num := c.Param("block_num")
	serve_block_info(p_block_num,c)
}
func serve_block_info(p_block_num string,c *gin.Context) {

	var block_num int64
	if len(p_block_num ) > 0 {
		var success bool
		block_num,success = parse_int_from_remote_or_error(c,false,&p_block_num)
		if !success {
			return
		}
	}
	block_info,err := augur_srv.db_augur.Get_block_info(block_num,block_num)
	if err == nil {
		c.HTML(http.StatusOK, "block_info.html", gin.H{
			"title": fmt.Sprintf("Block Number %v",block_num),
			"BlockInfo" : block_info,
		})
	} else {
		c.HTML(http.StatusBadRequest, "error.html", gin.H{
			"title": "Augur Markets: Error",
			"ErrDescr": fmt.Sprintf("DB error: %v",err),
		})
	}
}
func top_users(c *gin.Context) {

	top_profit_makers := augur_srv.db_augur.Get_top_profit_makers()
	top_trade_makers := augur_srv.db_augur.Get_top_trade_makers()
	top_volume_makers := augur_srv.db_augur.Get_top_volume_makers()
	c.HTML(http.StatusOK, "augur_v2/top_users.html", gin.H{
			"title": "Top 100 Users of Augur Markets",
			"ProfitMakers" : top_profit_makers,
			"TradeMakers" : top_trade_makers,
			"VolumeMakers" : top_volume_makers,
	})
}
func market_depth_status(c *gin.Context) {

	p_market_aid := c.Param("market_aid")
	var market_aid int64
	if len(p_market_aid) > 0 {
		var success bool
		market_aid,success = parse_int_from_remote_or_error(c,false,&p_market_aid)
		if !success {
			return
		}
	} else {
		respond_error(c,"Market ID is not set")
		return
	}

	p_outcome_idx := c.Param("outcome_idx")
	var outcome_idx int64
	if len(p_outcome_idx) > 0 {
		var success bool
		outcome_idx,success = parse_int_from_remote_or_error(c,false,&p_outcome_idx)
		if !success {
			return
		}
	} else {
		respond_error(c,"Outcome not set")
		return
	}

	p_last_oo_id := c.Param("last_oo_id")
	var last_oo_id int64
	if len(p_last_oo_id) > 0 {
		var success bool
		last_oo_id,success = parse_int_from_remote_or_error(c,false,&p_last_oo_id)
		if !success {
			return
		}
	} else {
		respond_error(c,"Last open order ID is not set")
		return
	}

	status,err := augur_srv.db_augur.Get_mdepth_status(market_aid,int(outcome_idx),last_oo_id)
	if err!=nil {
		respond_error(c,fmt.Sprintf("Error: %v",err))
		return
	}
	c.JSON(http.StatusOK,gin.H{
		"num_orders":status.NumOrders,
		"last_oo_id":status.LastOOID,
	})
}
func user_trades_for_market(c *gin.Context) {

	p_addr := c.Query("addr")
	user_addr_str,valid := is_address_valid(c,false,p_addr)
	if !valid {
		return
	}

	p_addr = c.Query("market")
	mkt_addr_str,valid := is_address_valid(c,false,p_addr)
	if !valid {
		return
	}

	aid,err:=augur_srv.db_augur.Nonfatal_lookup_address_id(user_addr_str)
	if err!=nil {
		c.HTML(http.StatusBadRequest, "error.html", gin.H{
			"title": "Augur Markets: Error",
			"ErrDescr": fmt.Sprintf("Such address wasn't found: %v",user_addr_str),
		})
		return
	}
	user_info,err := augur_srv.db_augur.Get_user_info(aid)
	if err!= nil {
		c.HTML(http.StatusBadRequest, "error.html", gin.H{
			"title": "Augur Markets: Error",
			"ErrDescr": fmt.Sprintf("No records found for address: %v",user_addr_str),
		})
		return
	}

	market_info,err := augur_srv.db_augur.Get_market_info(mkt_addr_str,0,false)
	if err!= nil {
		show_market_not_found_error(c,false,&mkt_addr_str)
		return
	}
	trades := augur_srv.db_augur.Get_user_trades_for_market(aid,market_info.MktAid)
	c.HTML(http.StatusOK, "augur_v2/user_trades.html", gin.H{
		"title": fmt.Sprintf("Trades for User %v",user_addr_str),
		"UTrades" : trades,
		"UserInfo" : user_info,
		"Market" : market_info,
	})
}
func account_statement(c *gin.Context) {

	addr := c.Param("addr")
	_,valid := is_address_valid(c,false,addr)
	if !valid {
		return
	}
	aid,err := augur_srv.db_augur.Nonfatal_lookup_address_id(addr)
	if err!= nil {
		c.HTML(http.StatusBadRequest, "error.html", gin.H{
			"title": "Augur Markets: Error",
			"ErrDescr": fmt.Sprintf("Address %v not found",addr),
		})
		return
	}
	transfers := augur_srv.db_augur.Get_account_statement(aid)
	c.HTML(http.StatusOK, "augur_v2/account_statement.html", gin.H{
			"Address" : addr,
			"Transfers": transfers,
	})
}
func open_order_history(c *gin.Context) {

	p_addr := c.Param("addr")
	user_addr_str,valid := is_address_valid(c,false,p_addr)
	if !valid {
		return
	}
	aid,err:=augur_srv.db_augur.Nonfatal_lookup_address_id(user_addr_str)
	if err!=nil {
		c.HTML(http.StatusBadRequest, "error.html", gin.H{
			"title": "Augur Markets: Error",
			"ErrDescr": fmt.Sprintf("Such address wasn't found: %v",user_addr_str),
		})
		return
	}
	user_info,err := augur_srv.db_augur.Get_user_info(aid)
	if err!= nil {
		c.HTML(http.StatusBadRequest, "error.html", gin.H{
			"title": "Augur Markets: Error",
			"ErrDescr": fmt.Sprintf("No records found for address: %v",user_addr_str),
		})
		return
	}
	oo_history := augur_srv.db_augur.Get_user_oo_history(aid)
	c.HTML(http.StatusOK, "augur_v2/user_oo_history.html", gin.H{
		"UserInfo" : user_info,
		"OOHistory" : oo_history,
	})
}
func price_estimate_history(c *gin.Context) {

	// Market Depth Info: https://en.wikipedia.org/wiki/Order_book_(trading)
	market := c.Param("market")
	market_addr,valid := is_address_valid(c,false,market)
	if !valid {
		return
	}
	p_outcome := c.Param("outcome")
	var outcome int
	if len(p_outcome) > 0 {
		p, err := strconv.Atoi(p_outcome)
		if err != nil {
			c.HTML(http.StatusBadRequest, "error.html", gin.H{
				"title": "Augur Markets: Error",
				"ErrDescr": "Can't parse outcome",
			})
			return
		}
		outcome = int(p)
	} else {
		c.HTML(http.StatusBadRequest, "error.html", gin.H{
			"title": "Augur Markets: Error",
			"ErrDescr": "Can't parse outcome",
		})
		return
	}
	market_info,err := augur_srv.db_augur.Get_market_info(market_addr,outcome,true)
	if err != nil {
		show_market_not_found_error(c,false,&market_addr)
		return
	}
	price_estimates := augur_srv.db_augur.Get_price_estimate_history(market_info.MktAid,outcome)
	js_price_estimate_data := build_js_price_estimate_history(&price_estimates)
	js_weighted_price_data := build_js_weighted_price_history(&price_estimates)
	c.HTML(http.StatusOK, "augur_v2/price_estimate.html", gin.H{
		"Market": market_info,
		"OutcomeIdx" : outcome,
		"PriceHistory" : price_estimates ,
		"JSPriceEst" : js_price_estimate_data,
		"JSWeightedPrice" : js_weighted_price_data,
	})
}
