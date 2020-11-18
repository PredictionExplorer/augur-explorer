package main
import (
	"fmt"
	"log"
	"strconv"
	"net/http"
	"encoding/hex"
	"github.com/gin-gonic/gin"
	//"html/template"
	"math/big"
	"context"
	//"strings"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"

	. "github.com/PredictionExplorer/augur-explorer/primitives"
	. "github.com/PredictionExplorer/augur-explorer/dbs"
)
const (
	DEFAULT_DB_LOG_FILE_NAME = "/var/tmp/backend-db.log"
	DEFAULT_MARKET_ROWS_LIMIT int	= 500
	DEFAILT_MARKET_TRADES_LIMIT int = 20
	DEFAULT_USER_REPORTS_LIMIT int = 10
	DEFAULT_MARKET_REPORTS_LIMIT int = 20
)
type AugurServer struct {
	storage		*SQLStorage
}
func create_augur_server(mkt_order_id_ptr *int64,dblog_fname string,info_log *log.Logger) *AugurServer {

	srv := new(AugurServer)
	srv.storage = Connect_to_storage(mkt_order_id_ptr,info_log)
	srv.storage.Init_log(dblog_fname)
	return srv
}
func respond_error(c *gin.Context,error_text string) {

	c.HTML(http.StatusBadRequest, "error.html", gin.H{
		"title": "Augur Markets: Error",
		"ErrDescr": error_text,
	})
}
func parse_int_from_remote_or_error(c *gin.Context,json_output bool,ascii_int *string) (int64,bool) {
	p, err := strconv.ParseInt(*ascii_int,10,64)
	if err != nil {
		if json_output {
			c.JSON(http.StatusBadRequest, gin.H{
				"status" : 0,
				"error": fmt.Sprintf("Can't parse integer parameter : ",err),
			})
		} else {
			c.HTML(http.StatusBadRequest, "error.html", gin.H{
				"title": "Augur Markets: Error",
				"ErrDescr": fmt.Sprintf("Can't parse integer parameter : ",err),
			})
		}
		return 0,false
	}
	return p,true
}
func mkt_depth_entry_to_js_obj(de *DepthEntry) string {

	var output string
	output = "{" +
				"x:" + fmt.Sprintf("%v",de.Price)  + "," +
				"y:"  + fmt.Sprintf("%v",de.AccumVol) + "," +
				"price: " + fmt.Sprintf("%v",de.Price) + "," +
				"addr: \"" + de.AddrSh + "\"," +
				"expires: \"" + de.Expires + "\"," +
				"volume: " + fmt.Sprintf("%v",de.Volume) + "," +
				"click: function() {load_order_data(\"" +
					de.AddrSh +"\"," +
					fmt.Sprintf("%v,%v,%v,%v,%v,%v,\"%v\",\"%v\"",
										de.MktAid,
										de.OutcomeIdx,
										de.Volume,
										de.TotalBids,
										de.TotalAsks,
										de.TotalCancel,
										de.DateCreated,
										de.Expires,
					) +
				")}" +
				"}"
	return output
}
func main_page(c *gin.Context) {
	blknum,_:= augur_srv.storage.Get_last_block_num()
	blknum_thousand_separated := ThousandsFormat(int64(blknum))
	stats := augur_srv.storage.Get_front_page_stats()
	c.HTML(http.StatusOK, "index.html", gin.H{
			"title": "Augur Prediction Markets",
			"block_num" : blknum_thousand_separated,
			"Stats" : stats,
	})
}
func markets(c *gin.Context) {
	off_str := c.Query("off")
	var off int = 0
	var err error
	if len(off_str) > 0 {
		off, err = strconv.Atoi(off_str)
		if err != nil {
			c.HTML(http.StatusBadRequest, "error.html", gin.H{
				"title": "Augur Markets: Error",
				"ErrDescr": "Can't parse offset",
			})
			return
		}
	}
	markets := augur_srv.storage.Get_active_market_list(off,DEFAULT_MARKET_ROWS_LIMIT)
	c.HTML(http.StatusOK, "markets.html", gin.H{
			"title": "Augur Markets",
			"Markets" : markets,
	})
}
func categories(c *gin.Context) {
	blknum,_ := augur_srv.storage.Get_last_block_num()
	categories := augur_srv.storage.Get_categories()
	c.HTML(http.StatusOK, "categories.html", gin.H{
			"title": "Augur Market Categories",
			"block_num" : blknum,
			"Categories" : categories,
	})
}
func statistics(c *gin.Context) {

	stats := augur_srv.storage.Get_main_stats()
	cash_flow_entries := augur_srv.storage.Get_cash_flow()
	gas_usage := augur_srv.storage.Get_gas_usage_global()
	uniq_addr_entries := augur_srv.storage.Get_unique_addresses()
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

	c.HTML(http.StatusOK, "statistics.html", gin.H{
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
func explorer(c *gin.Context) {
	blknum,res := augur_srv.storage.Get_last_block_num()
	_ = res
	c.HTML(http.StatusOK, "explorer.html", gin.H{
			"title": "Augur Event Explorer",
			"block_num" : blknum,
	})
}
func complete_and_output_market_info(c *gin.Context,json_output bool,minfo InfoMarket) {
/*
	var limit int64 = int64(DEFAILT_MARKET_TRADES_LIMIT);
	p_limit := c.Query("limit")
	if len(p_limit) > 0 {
		var success bool
		limit,success = parse_int_from_remote_or_error(c,json_output,&p_limit)
		if !success {
			return
		}
	}
*/
	trades := augur_srv.storage.Get_mkt_trades(minfo.MktAddr,10000000)
	outcome_vols,_ := augur_srv.storage.Get_outcome_volumes(minfo.MktAddr,minfo.MktAid,0,minfo.LowPriceLimit)
	price_estimates := augur_srv.storage.Get_price_estimates(minfo.MktAid,outcome_vols,minfo.LowPriceLimit)
	reports := augur_srv.storage.Get_market_reports(minfo.MktAid,DEFAULT_MARKET_REPORTS_LIMIT)
	price_history := augur_srv.storage.Get_full_price_history(minfo.MktAddr,minfo.MktAid,minfo.LowPriceLimit)
	balancer_pools := augur_srv.storage.Get_market_balancer_pools(minfo.MktAid)
	wrappers := augur_srv.storage.Get_wrapped_tokens_for_market(minfo.MktAid)

	if json_output {
		c.JSON(200,gin.H{
			"Trades" : trades,
			"Reports" : reports,
			"Market": minfo ,
			"OutcomeVols" : outcome_vols,
			"PriceHistory" : price_history,
			"PriceEstimates" : price_estimates,
			"BalancerPools" : balancer_pools,
			"WrappedContracts": wrappers,
		})
	} else {
		c.HTML(http.StatusOK, "market_info.html", gin.H{
			"title": "DISCONTINUED",
			"Trades" : trades,
			"Reports" : reports,
			"Market": minfo ,
			"OutcomeVols" : outcome_vols,
			"PriceHistory" : price_history,
			"PriceEstimates" : price_estimates,
			"BalancerPools" : balancer_pools,
		})
	}
}
func is_address_valid(c *gin.Context,json_output bool,addr string) (string,bool) {

	if (len(addr) != 40) && (len(addr)!=42) {
		var err_msg = fmt.Sprintf("Provided address has invalid length (len=%v)",len(addr))
		if json_output {
			c.JSON(200,gin.H{
				"status": 0,
				"error": err_msg,
			})
		} else {
			c.HTML(http.StatusBadRequest, "error.html", gin.H{
				"title": "Augur Markets: Error",
				"ErrDescr": err_msg,
			})
		}
		return "",false
	}
	if (addr[0]=='0') && (addr[1] == 'x') {
		addr = addr[2:]
	}
	if len(addr) != 40 {
		if json_output {
			c.JSON(200,gin.H{
				"status": 0,
				"error": fmt.Sprintf("Invalid address length"),
			})
		} else {
			c.HTML(http.StatusBadRequest, "error.html", gin.H{
				"title": "Augur Markets: Error",
				"ErrDescr": fmt.Sprintf("Invalid address length"),
			})
		}
		return "",false
	}
	var formatted_addr string
	addr_bytes,err := hex.DecodeString(addr)
	if err == nil {
		addr := common.BytesToAddress(addr_bytes)
		formatted_addr = addr.String()
	} else {
		if json_output {
			c.JSON(200,gin.H{
				"status": 0,
				"error": fmt.Sprintf("Provided address parameter is invalid : %v",err),
			})
		} else {
			c.HTML(http.StatusBadRequest, "error.html", gin.H{
				"title": "Augur Markets: Error",
				"ErrDescr": fmt.Sprintf("Provided address parameter is invalid : %v",err),
			})
		}
		return "",false
	}
	return formatted_addr,true
}
func json_validate_and_lookup_address_or_aid(c *gin.Context,p_addr *string) (string,int64,bool) {
	// Note: this function transforms address in checksumed format
	var aid int64 = 0
	if len(*p_addr) > 0 {
		aid, err := strconv.ParseInt(*p_addr,10,64)
		if err == nil {
			var addr string
			addr,err = augur_srv.storage.Lookup_address(aid)
			if err!=nil {
				c.JSON(http.StatusBadRequest,gin.H{
					"status":0,
					"error":fmt.Sprintf("Address with ID=%v not found",aid),
				})
				return "",aid,false
			}
			return addr,aid,true
		} else {
			aid = 0
		}
	} else {
		c.JSON(http.StatusBadRequest,gin.H{
			"status":0,
			"error":fmt.Sprintf("Empty 'address' parameter for lookup"),
		})
		return "",0,false
	}
	address,valid:=is_address_valid(c,true,*p_addr)
	if !valid {
		return "",0,false
	}
	aid,err := augur_srv.storage.Nonfatal_lookup_address_id(*p_addr)
	if err != nil {
		c.JSON(http.StatusBadRequest,gin.H{
			"status":0,
			"error":fmt.Sprintf("Address not found in the DB"),
		})
		return "",0,false
	}
	return address,aid,true
}
func show_market_not_found_error(c *gin.Context,json_output bool,addr *string) {

	if json_output {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": 0,
			"error": fmt.Sprintf("Market with address %v wasn't found",*addr),
		})
	} else {
		c.HTML(http.StatusBadRequest, "error.html", gin.H{
			"title": "Augur Markets: Error",
			"ErrDescr": fmt.Sprintf("Market with address %v wasn't found",*addr),
		})
	}
}
func market_info(c *gin.Context) {

	market := c.Param("market")
	market_addr,valid:=is_address_valid(c,false,market)
	if !valid {
		return
	}
	market_info,err := augur_srv.storage.Get_market_info(market_addr,0,false)
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
	market_info,err := augur_srv.storage.Get_market_info(market_addr,0,false)
	if err != nil {
		show_market_not_found_error(c,false,&market_addr)
		return
	}
	trades := augur_srv.storage.Get_mkt_trades(market_addr,0)
	c.HTML(http.StatusOK, "full_trade_list.html", gin.H{
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

	market_info,err := augur_srv.storage.Get_market_info(market_addr,outcome,true)
	if err != nil {
		show_market_not_found_error(c,false,&market_addr)
		return
	}
	mdepth,last_oo_id := augur_srv.storage.Get_mkt_depth(market_info.MktAid,outcome)
	num_orders:=len(mdepth.Bids) + len(mdepth.Asks)
	js_bid_data,js_ask_data := build_js_data_obj(mdepth)
	c.HTML(http.StatusOK, "market_depth.html", gin.H{
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
	mdepth,last_oo_id := augur_srv.storage.Get_mkt_depth(market_aid,int(outcome))
	js_bid_data,js_ask_data := build_js_data_obj(mdepth)
	c.JSON(200,gin.H{
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

	market_info,err := augur_srv.storage.Get_market_info(market_addr,outcome,true)
	if err != nil {
		show_market_not_found_error(c,false,&market_addr)
		return
	}
	mkt_price_hist := augur_srv.storage.Get_price_history_for_outcome(market_info.MktAid,outcome,market_info.LowPriceLimit)
	js_price_history := build_js_price_history(&mkt_price_hist)
	fmt.Printf("js price history = %v\n",js_price_history)
	c.HTML(http.StatusOK, "price_history.html", gin.H{
			"title": "Market Price History",
			"Market": market_info,
			"Prices": mkt_price_hist,
			"JSPriceData": js_price_history,
	})
}
func serve_user_info_page(c *gin.Context,addr string) {

	eoa_aid,err := augur_srv.storage.Nonfatal_lookup_address_id(addr)
	if err == nil {
		user_info,err := augur_srv.storage.Get_user_info(eoa_aid)
		if err == nil {
			pl_entries := augur_srv.storage.Get_profit_loss(eoa_aid)
			open_pos_entries := augur_srv.storage.Get_open_positions(eoa_aid)
			js_pl_data := build_js_profit_loss_history(&pl_entries)
			js_open_pos_data := build_js_open_positions(&open_pos_entries)
			user_reports := augur_srv.storage.Get_user_reports(eoa_aid,DEFAULT_USER_REPORTS_LIMIT)
			user_active_markets := augur_srv.storage.Get_traded_markets_for_user(eoa_aid,1)
			var has_active_markets bool = false
			if len(user_active_markets) > 0 {
				has_active_markets = true
			}
			open_orders := augur_srv.storage.Get_user_open_orders(user_info.Aid)
			gas_spent,_ := augur_srv.storage.Get_gas_spent_for_user(eoa_aid)

			c.HTML(http.StatusOK, "user_info.html", gin.H{
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

	tx_info,err := augur_srv.storage.Get_transaction(tx_hash)
	if err == nil {
		c.HTML(http.StatusOK, "tx_info.html", gin.H{
			"title": "Transaction " + tx_hash,
			"TxInfo" : tx_info,
		})
	} else {
		c.HTML(http.StatusOK, "tx_not_found.html", gin.H{
			"title": "Transaction "+tx_hash,
			"tx_hash": tx_hash,
		})
	}
}
func get_token_balance(token_type int,addr *common.Address) float64 {
	// input: token_type = 0 => DAI token; token_type = 1 => Rep token
	switch token_type {	//  null pointer error prevention
		case 0:
			if ctrct_dai_token == nil {
				return 0.0
			}
		case 1:
			if ctrct_rep_token == nil {
				return 0.0
			}
	}
	big_float_balance := big.NewFloat(0.0)
	var copts = new(bind.CallOpts)
	var err error
	var int_balance *big.Int
	switch token_type {
		case 0:
			int_balance,err = ctrct_dai_token.BalanceOf(copts,*addr)
			fmt.Printf("switch: DAI int_balance=%v\n",int_balance.String())
		case 1:
			int_balance,err = ctrct_rep_token.BalanceOf(copts,*addr)
			fmt.Printf("switch: REP int_balance=%v\n",int_balance.String())
		default:
			Fatalf("get_token_balance(): undefined behavior")
	}
	if err == nil {
		f_balance :=big.NewFloat(0.0)
		f_balance.SetInt(int_balance)
		divisor:=big.NewFloat(0.0)
		divisor.SetString("1000000000000000000.0")
		div_result:=new(big.Float).Quo(f_balance,divisor)
		big_float_balance.Set(div_result)
	} else {
		fmt.Printf("Error retrieving token (type=%v) balance for addr %v: %v\n",
							token_type,addr.String(),err)
	}
	balance,_:=big_float_balance.Float64()
	return balance
}
func get_eth_balance(addr *common.Address) float64 {
	ctx := context.Background()
	var float_eth_balance float64 = 0.0
	big_eth_balance,err := rpcclient.BalanceAt(ctx,*addr,nil)
	if err == nil {
		big_float_eth_balance := big.NewFloat(0.0)
		big_float_eth_balance.SetInt(big_eth_balance)
		divisor:=big.NewFloat(0.0)
		divisor.SetString("1000000000000000000.0")
		div_result:=new(big.Float).Quo(big_float_eth_balance,divisor)
		float_eth_balance,_ = div_result.Float64()
	}
	return float_eth_balance
}
func serve_user_funds_v2(c *gin.Context,addr *string) {
	// the input address must be EOA, from that we can get Wallet addr
	var (
		eoa_dai_balance float64 = 0.0
		eoa_rep_balance float64 = 0.0
		eoa_eth_balance float64 = 0.0
		wallet_dai_balance float64 = 0.0
		wallet_rep_balance float64 = 0.0
		wallet_eth_balance float64 = 0.0
	)

	var status_code int = 0
	var error_text  string = ""
	var wallet_aid int64 = 0
	eoa_aid,err := augur_srv.storage.Nonfatal_lookup_address_id(*addr)
	if err == nil {
		wallet_aid,_ = augur_srv.storage.Lookup_wallet_aid(eoa_aid)
	} else {
		error_text = "Address lookup failed"
	}
	if eoa_aid > 0 {
		addr := common.HexToAddress(*addr)
		eoa_dai_balance = get_token_balance(0,&addr)
		eoa_rep_balance = get_token_balance(1,&addr)
		eoa_eth_balance = get_eth_balance(&addr)
		status_code = 1
	}

	if wallet_aid != 0 {
		wallet_addr,err := augur_srv.storage.Lookup_address(wallet_aid)
		if err == nil {
			waddr := common.HexToAddress(wallet_addr)
			wallet_dai_balance = get_token_balance(0,&waddr)
			wallet_rep_balance = get_token_balance(1,&waddr)
			wallet_eth_balance = get_eth_balance(&waddr)
			status_code = 1
		}
	}

	c.JSON(200, gin.H{
		"status": status_code,
		"error": error_text,
		"eoa_eth": fmt.Sprintf("%v",eoa_eth_balance),
		"wallet_eth": fmt.Sprintf("%v",wallet_eth_balance),
		"eoa_dai": fmt.Sprintf("%v",eoa_dai_balance),
		"wallet_dai": fmt.Sprintf("%v",wallet_dai_balance),
		"eoa_rep": fmt.Sprintf("%v",eoa_rep_balance),
		"wallet_rep": fmt.Sprintf("%v",wallet_rep_balance),
		"eoa_usd" : 0,
		"wallet_usd" : 0,
	})
}
func serve_user_funds_v1(c *gin.Context,addr common.Address) {
	// the input address must be EOA, from that we can get Wallet addr
	// Note: this request is becoming obsolete, and will be removed later

	var wallet_aid int64 = 0
	eoa_aid,err := augur_srv.storage.Nonfatal_lookup_address_id(addr.String())
	if err == nil {
		wallet_aid,_ = augur_srv.storage.Lookup_wallet_aid(eoa_aid)
	} else {
		c.JSON(200,gin.H{
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
		wallet_addr,err := augur_srv.storage.Lookup_address(wallet_aid)
		if err == nil {
			waddr := common.HexToAddress(wallet_addr)
			wallet_dai_balance = get_token_balance(0,&waddr)
			wallet_rep_balance = get_token_balance(1,&waddr)
			wallet_eth_balance = get_eth_balance(&waddr)
		} else {
			fmt.Printf("address lookup for wallet_aid = %v failed: %v",wallet_aid,err)
		}
	}

	c.JSON(200, gin.H{
			"eoa_eth": fmt.Sprintf("%v",eoa_eth_balance),
			"wallet_eth": fmt.Sprintf("%v",wallet_eth_balance),
			"eoa_dai": fmt.Sprintf("%v",eoa_dai_balance),
			"wallet_dai": fmt.Sprintf("%v",wallet_dai_balance),
			"eoa_rep": fmt.Sprintf("%v",eoa_rep_balance),
			"wallet_rep": fmt.Sprintf("%v",wallet_rep_balance),
	})
}
func search(c *gin.Context) {

	keyword := c.Query("q")
	if (len(keyword) == 40) || (len(keyword) == 42) { // address
		if len(keyword) == 42 {	// strip 0x prefix
			keyword = keyword[2:]
		}
		addr_bytes,err := hex.DecodeString(keyword)
		if err == nil {
			addr := common.BytesToAddress(addr_bytes)
			addr_str := addr.String()
			_,err:=augur_srv.storage.Nonfatal_lookup_address_id(addr_str)
			if err==nil {
				market_info,err := augur_srv.storage.Get_market_info(addr_str,0,false)
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
		if augur_srv.storage.Tx_exists(hash) {
			serve_tx_info_page(c,hash)
			return
		}
		orders := augur_srv.storage.Get_filling_orders_by_hash(hash)
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
func read_money(c *gin.Context) {
	// this function gets amount of currencies the User holds: ETH, DAI and REP (all in one call)
	addr := c.Param("addr")
	if (len(addr) == 40) || (len(addr) == 42) { // address
		if len(addr) == 42 {	// strip 0x prefix
			addr = addr[2:]
		}
		addr_bytes,err := hex.DecodeString(addr)
		if err == nil {
			addr := common.BytesToAddress(addr_bytes)
			serve_user_funds_v1(c,addr)
		} else {
			c.HTML(http.StatusBadRequest, "error.html", gin.H{
				"title": "Augur Markets: Error",
				"ErrDescr": "Invalid HEX string in address parameter",
			})
			return
		}
	}
}
func output_filling_orders(c *gin.Context,order_hash string,orders []OrderInfo) {
	c.HTML(http.StatusOK, "filling_orders.html", gin.H{
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
		order,err := augur_srv.storage.Get_order_info_by_id(order_id)
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
	orders := augur_srv.storage.Get_filling_orders_by_hash(order_hash)
	output_filling_orders(c,order_hash,orders)
}
func category(c *gin.Context) {

	p_catid:= c.Param("catid")

	cat_id,success := parse_int_from_remote_or_error(c,false,&p_catid)
	if !success {
		return
	}
	cat_markets := augur_srv.storage.Get_category_markets(int64(cat_id))
	c.HTML(http.StatusOK, "category_markets.html", gin.H{
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
		aid,err:=augur_srv.storage.Nonfatal_lookup_address_id(addr_str)
		if err==nil {
			user_info,err := augur_srv.storage.Get_user_info(aid)
			if err!= nil {
				c.HTML(http.StatusBadRequest, "error.html", gin.H{
					"title": "Augur Markets: Error",
					"ErrDescr": fmt.Sprintf("No records found for address: %v",addr_str),
				})
			} else {
				user_reports := augur_srv.storage.Get_user_reports(aid,0)
				c.HTML(http.StatusOK, "full_user_reports.html", gin.H{
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
		aid,err:=augur_srv.storage.Nonfatal_lookup_address_id(addr_str)
		if err==nil {
			user_info,err := augur_srv.storage.Get_user_info(aid)
			if err!= nil {
				c.HTML(http.StatusBadRequest, "error.html", gin.H{
					"title": "Augur Markets: Error",
					"ErrDescr": fmt.Sprintf("No records found for address: %v",addr_str),
				})
			} else {
				user_reports := augur_srv.storage.Get_user_markets(aid)
				c.HTML(http.StatusOK, "user_markets.html", gin.H{
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
		aid,err:=augur_srv.storage.Nonfatal_lookup_address_id(addr_str)
		if err==nil {
			user_info,err := augur_srv.storage.Get_user_info(aid)
			if err!= nil {
				c.HTML(http.StatusBadRequest, "error.html", gin.H{
					"title": "Augur Markets: Error",
					"ErrDescr": fmt.Sprintf("No records found for address: %v",addr_str),
				})
			} else {
				wallet_aid,err := augur_srv.storage.Lookup_wallet_aid(aid)
				if err == nil {
					user_deposits_withdrawals := augur_srv.storage.Get_deposits_withdrawals(wallet_aid)
					c.HTML(http.StatusOK, "user_deposits_withdrawals.html", gin.H{
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
	block_info,err := augur_srv.storage.Get_block_info(block_num)
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

	top_profit_makers := augur_srv.storage.Get_top_profit_makers()
	top_trade_makers := augur_srv.storage.Get_top_trade_makers()
	top_volume_makers := augur_srv.storage.Get_top_volume_makers()
	c.HTML(http.StatusOK, "top_users.html", gin.H{
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

	status,err := augur_srv.storage.Get_mdepth_status(market_aid,int(outcome_idx),last_oo_id)
	if err!=nil {
		respond_error(c,fmt.Sprintf("Error: %v",err))
		return
	}
	c.JSON(200,gin.H{
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

	aid,err:=augur_srv.storage.Nonfatal_lookup_address_id(user_addr_str)
	if err!=nil {
		c.HTML(http.StatusBadRequest, "error.html", gin.H{
			"title": "Augur Markets: Error",
			"ErrDescr": fmt.Sprintf("Such address wasn't found: %v",user_addr_str),
		})
		return
	}
	user_info,err := augur_srv.storage.Get_user_info(aid)
	if err!= nil {
		c.HTML(http.StatusBadRequest, "error.html", gin.H{
			"title": "Augur Markets: Error",
			"ErrDescr": fmt.Sprintf("No records found for address: %v",user_addr_str),
		})
		return
	}

	market_info,err := augur_srv.storage.Get_market_info(mkt_addr_str,0,false)
	if err!= nil {
		show_market_not_found_error(c,false,&mkt_addr_str)
		return
	}
	trades := augur_srv.storage.Get_user_trades_for_market(aid,market_info.MktAid)
	c.HTML(http.StatusOK, "user_trades.html", gin.H{
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
	aid,err := augur_srv.storage.Nonfatal_lookup_address_id(addr)
	if err!= nil {
		c.HTML(http.StatusBadRequest, "error.html", gin.H{
			"title": "Augur Markets: Error",
			"ErrDescr": fmt.Sprintf("Address %v not found",addr),
		})
		return
	}
	transfers := augur_srv.storage.Get_account_statement(aid)
	c.HTML(http.StatusOK, "account_statement.html", gin.H{
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
	aid,err:=augur_srv.storage.Nonfatal_lookup_address_id(user_addr_str)
	if err!=nil {
		c.HTML(http.StatusBadRequest, "error.html", gin.H{
			"title": "Augur Markets: Error",
			"ErrDescr": fmt.Sprintf("Such address wasn't found: %v",user_addr_str),
		})
		return
	}
	user_info,err := augur_srv.storage.Get_user_info(aid)
	if err!= nil {
		c.HTML(http.StatusBadRequest, "error.html", gin.H{
			"title": "Augur Markets: Error",
			"ErrDescr": fmt.Sprintf("No records found for address: %v",user_addr_str),
		})
		return
	}
	oo_history := augur_srv.storage.Get_user_oo_history(aid)
	c.HTML(http.StatusOK, "user_oo_history.html", gin.H{
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
	market_info,err := augur_srv.storage.Get_market_info(market_addr,outcome,true)
	if err != nil {
		show_market_not_found_error(c,false,&market_addr)
		return
	}
	price_estimates := augur_srv.storage.Get_price_estimate_history(market_info.MktAid,outcome)
/*
	mdepth,last_oo_id := augur_srv.storage.Get_mkt_depth(market_info.MktAid,outcome)
	num_orders:=len(mdepth.Bids) + len(mdepth.Asks)
*/
	js_price_estimate_data := build_js_price_estimate_history(&price_estimates)
	js_weighted_price_data := build_js_weighted_price_history(&price_estimates)
	c.HTML(http.StatusOK, "price_estimate.html", gin.H{
		"Market": market_info,
		"OutcomeIdx" : outcome,
		"PriceHistory" : price_estimates ,
		"JSPriceEst" : js_price_estimate_data,
		"JSWeightedPrice" : js_weighted_price_data,
	})
}
func wrapped_tokens(c *gin.Context) {

	market := c.Param("market")
	market_addr,valid := is_address_valid(c,false,market)
	if !valid {
		return
	}
	market_info,err := augur_srv.storage.Get_market_info(market_addr,0,false)
	if err != nil {
		show_market_not_found_error(c,false,&market_addr)
		return
	}
	wrappers := augur_srv.storage.Get_wrapped_tokens_for_market(market_info.MktAid)
	c.HTML(http.StatusOK, "wrapper_contracts.html", gin.H{
			"WrapperContracts" : wrappers,
			"Market": market_info,
	})
}
func wrapped_token_transfers(c *gin.Context) {

	address:= c.Param("address")
	addr,valid := is_address_valid(c,false,address)
	if !valid {
		return
	}
	aid,err := augur_srv.storage.Nonfatal_lookup_address_id(addr)
	if err!=nil {
		c.HTML(http.StatusBadRequest, "error.html", gin.H{
			"title": "Augur Markets: Error",
			"ErrDescr": fmt.Sprintf("Such address wasn't found: %v",address),
		})
		return
	}
	wrapper_info := augur_srv.storage.Get_wrapped_token_info(aid)
	market_info,err := augur_srv.storage.Get_market_info(wrapper_info.MktAddr,wrapper_info.OutcomeIdx,true)
	transfers := augur_srv.storage.Get_wrapped_token_transfers(aid,0,500)
	c.HTML(http.StatusOK, "wrapped_transfers.html", gin.H{
			"MarketInfo" : market_info,
			"TokenInfo" : wrapper_info,
			"WrappedTransfers" : transfers,
	})
}
func pool_swaps(c *gin.Context) {

	address:= c.Param("address")
	addr,valid := is_address_valid(c,false,address)
	if !valid {
		return
	}
	aid,err := augur_srv.storage.Nonfatal_lookup_address_id(addr)
	if err!=nil {
		c.HTML(http.StatusBadRequest, "error.html", gin.H{
			"title": "Augur Markets: Error",
			"ErrDescr": fmt.Sprintf("Such address wasn't found: %v",address),
		})
		return
	}
	pool_info := augur_srv.storage.Get_pool_info(aid)
	swaps := augur_srv.storage.Get_pool_swaps(aid,0,200)
	c.HTML(http.StatusOK, "pool_swaps.html", gin.H{
			"PoolInfo" : pool_info,
			"PoolSwaps" : swaps,
	})
}
func sharetoken_balance_changes(c *gin.Context) {

	market := c.Param("market")
	market_addr,valid:=is_address_valid(c,false,market)
	if !valid {
		return
	}
	minfo,err := augur_srv.storage.Get_market_info(market_addr,0,false)
	if err != nil {
		show_market_not_found_error(c,false,&market_addr)
		return
	}

	outag_sh_bal_chgs := augur_srv.storage.Outside_augur_share_balance_changes(minfo.MktAid,0,500)
	c.HTML(http.StatusOK, "sharetoken_balance_changes.html", gin.H{
			"MarketInfo" : minfo,
			"OutsideAugurBalanceChanges": outag_sh_bal_chgs,
	})
}
func market_uniswap_pairs(c *gin.Context) {

	market := c.Param("market")
	market_addr,valid:=is_address_valid(c,false,market)
	if !valid {
		return
	}
	minfo,err := augur_srv.storage.Get_market_info(market_addr,0,false)
	if err != nil {
		show_market_not_found_error(c,false,&market_addr)
		return
	}
	pairs := augur_srv.storage.Get_market_uniswap_pairs(minfo.MktAid)
	c.HTML(http.StatusOK, "market_upairs.html", gin.H{
			"Market" : minfo,
			"MarketUniswapPairs": pairs,
	})

}
func uniswap_swaps(c *gin.Context) {

	address:= c.Param("address")
	addr,valid := is_address_valid(c,false,address)
	if !valid {
		return
	}
	aid,err := augur_srv.storage.Nonfatal_lookup_address_id(addr)
	if err!=nil {
		c.HTML(http.StatusBadRequest, "error.html", gin.H{
			"title": "Augur Markets: Error",
			"ErrDescr": fmt.Sprintf("Such address wasn't found: %v",address),
		})
		return
	}
	pair_info,_:= augur_srv.storage.Get_uniswap_pair_info(aid)
	swaps := augur_srv.storage.Get_uniswap_swaps(aid,0,200)
	c.HTML(http.StatusOK, "uniswap_swaps.html", gin.H{
			"PairInfo" : pair_info,
			"PairSwaps" : swaps,
	})
}
func do_text_search(c *gin.Context) {

	keywords := c.Query("keywords")
	if len(keywords) == 0 {
		c.HTML(http.StatusBadRequest, "error.html", gin.H{
			"title": "Augur Markets: Error",
			"ErrDescr": fmt.Sprintf("Empty search query"),
		})
		return
	}
	search_results := augur_srv.storage.Search_keywords_in_markets(keywords)
	c.HTML(http.StatusOK, "text_search_results.html", gin.H{
			"SearchResults" : search_results,
	})
}
func show_text_search_form(c *gin.Context) {

	c.HTML(http.StatusOK, "text_search_form.html", gin.H{
	})
}
