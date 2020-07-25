package main
import (
	"fmt"
	"log"
	"strconv"
	"net/http"
	"encoding/hex"
	"github.com/gin-gonic/gin"
	"html/template"
	"math/big"
	"context"
	"strings"

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
func parse_int_from_remote_or_error(c *gin.Context,ascii_int *string) (int64,bool) {
	p, err := strconv.ParseInt(*ascii_int,10,64)
	if err != nil {
		c.HTML(http.StatusBadRequest, "error.html", gin.H{
			"title": "Augur Markets: Error",
			"ErrDescr": "Can't parse integer parameter (non-numeric characters detected)",
		})
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
				"addr: \"" + de.EOAAddrSh + "\"," +
				"expires: \"" + de.Expires + "\"," +
				"volume: " + fmt.Sprintf("%v",de.Volume) + "," +
				"click: function() {load_order_data(\"" +
					de.EOAAddrSh +"\",\"" +
					de.WalletAddrSh + "\"," +
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
func build_javascript_data_obj(mdepth *MarketDepth) (template.JS,template.JS) {
	var asks_str string = "["
	var bids_str string = "["

	var last_price float64 = 0.0
	for i:=0 ; i < len(mdepth.Asks) ; i++ {
		if len(asks_str) > 1 {
			asks_str = asks_str + ","
		}
		var entry string
		entry = mkt_depth_entry_to_js_obj(&mdepth.Asks[i])
		asks_str= asks_str + entry
		last_price = mdepth.Asks[i].Price
	}
/*	// Possibly replace this with a line indicating the spread, as another Serie
	if len(mdepth.Asks) > 0 {
		if len(mdepth.Bids) > 0 {
			// add fake BID entry to fill the hole for the spread
			last_elt:=len(mdepth.Asks)-1
			fake_entry := mdepth.Asks[last_elt]
			fake_entry.Price = mdepth.Bids[0].Price*10
			bids_str = "[" + mkt_depth_entry_to_js_obj(&fake_entry)
		}
	}
*/
	for i:=0 ; i < len(mdepth.Bids) ; i++ {
		if len(bids_str) > 1 {
			bids_str = bids_str + ","
		}
		var entry string
		entry = mkt_depth_entry_to_js_obj(&mdepth.Bids[i])
		bids_str= bids_str + entry
		last_price = mdepth.Bids[i].Price
	}

	asks_str = asks_str + "]"
	bids_str = bids_str + "]"
	_ = last_price
	return template.JS(bids_str),template.JS(asks_str)
}
func build_javascript_price_history(orders *[]MarketOrder) template.JS {
	var data_str string = "["

	for i:=0 ; i < len(*orders) ; i++ {
		if len(data_str) > 1 {
			data_str = data_str + ","
		}
		var e = &(*orders)[i];
		var entry string
		entry = "{" +
//				"x:" + fmt.Sprintf("\"%v\"",e.Date)  + "," +
				"x:" + fmt.Sprintf("%v",i)  + "," +
				"y:"  + fmt.Sprintf("%v",e.Price) + "," +
				"price: " + fmt.Sprintf("%v",e.Price) + "," +
				"volume: " + fmt.Sprintf("%v",e.Volume) + "," +
				"click: function() {load_order_data(\"" +
					e.CreatorEOAAddr +"\",\"" +
					e.FillerEOAAddr+ "\"," +
					fmt.Sprintf("%v,%v,%v,\"%v\"",e.MktAid,e.Price,e.Volume,e.Date) +
				")}" +
				"}"
		data_str= data_str + entry
	}
	data_str = data_str + "]"
	fmt.Printf("JS price history string: %v\n",data_str)
	return template.JS(data_str)
}
func build_javascript_profit_loss_history(entries *[]PLEntry) template.JS {
	var data_str string = "["

	for i:=0 ; i < len(*entries) ; i++ {
		if len(data_str) > 1 {
			data_str = data_str + ","
		}
		var e = &(*entries)[i];
		outcome_escaped := strings.ReplaceAll(e.OutcomeStr,"\"","\\\"")
		descr_escaped := strings.ReplaceAll(e.MktDescr,"\"","\\\"")
		var entry string
		entry = "{" +
				"x:" + fmt.Sprintf("%v",i)  + "," +
				"y:"  + fmt.Sprintf("%v",e.AccumPl) + "," +
				"pl: " + fmt.Sprintf("%v",e.ImmediateProfit) + "," +
				"pl_accum: " + fmt.Sprintf("%v",e.AccumPl) + "," +
				"date: \"" + fmt.Sprintf("%v",e.Date) + "\"," +
				"click: function() {load_pl_data(" +
					fmt.Sprintf("%v,%v,%v,%v,%v,%v,\"%v\",\"%v\",\"%v\",\"%v\",\"%v\",\"%v\",\"%v\",\"%v\",%v",
							e.Id,e.ClaimStatus,e.NetPosition,e.AvgPrice,e.ImmediateProfit,e.AccumPl,e.MktAddr,e.MktAddrSh,outcome_escaped,
							descr_escaped,e.Date,e.CounterPAddr,e.CounterPAddrSh,e.OrderHash,e.BlockNum) +
				")}" +
				"}"
		data_str= data_str + entry
	}
	data_str = data_str + "]"
	return template.JS(data_str)
}
func build_javascript_open_positions(entries *[]PLEntry) template.JS {
	var data_str string = "["

	for i:=0 ; i < len(*entries) ; i++ {
		if len(data_str) > 1 {
			data_str = data_str + ","
		}
		var e = &(*entries)[i];
		outcome_escaped := strings.ReplaceAll(e.OutcomeStr,"\"","\\\"")
		descr_escaped := strings.ReplaceAll(e.MktDescr,"\"","\\\"")
		var entry string
		entry = "{" +
				"x:" + fmt.Sprintf("%v",i)  + "," +
				"y:"  + fmt.Sprintf("%v",e.AccumFrozen) + "," +
				"frozen: " + fmt.Sprintf("%v",e.FrozenFunds) + "," +
				"frozen_accum: " + fmt.Sprintf("%v",e.AccumFrozen) + "," +
				"date: \"" + fmt.Sprintf("%v",e.Date) + "\"," +
				"click: function() {load_open_pos_data(" +
					fmt.Sprintf("%v,%v,%v,\"%v\",\"%v\",\"%v\",\"%v\",\"%v\",\"%v\",\"%v\",\"%v\",%v",
							e.AvgPrice,e.FrozenFunds,e.NetPosition,e.MktAddr,e.MktAddrSh,outcome_escaped,
							descr_escaped,e.Date,e.CounterPAddr,e.CounterPAddrSh,e.OrderHash,e.BlockNum) +
				")}" +
				"}"
		data_str= data_str + entry
	}
	data_str = data_str + "]"
	return template.JS(data_str)
}
func build_javascript_cash_flow_data(entries *[]BlockCash) template.JS {
	var data_str string = "["

	for i:=0 ; i < len(*entries) ; i++ {
		if len(data_str) > 1 {
			data_str = data_str + ","
		}
		var e = &(*entries)[i];
		var entry string
		entry = "{" +
				//"x:" + fmt.Sprintf("%v",i)  + "," +
				"x:" + fmt.Sprintf("new Date(%v * 1000)",e.Ts)  + "," +
				"y:"  + fmt.Sprintf("%.2f",e.CashFlow) + "," +
				"block_num: " + fmt.Sprintf("%v",e.BlockNum) + "," +
				"cash: " + fmt.Sprintf("%v",e.CashFlow) + "" +
				"}"
		data_str= data_str + entry
	}
	data_str = data_str + "]"
	return template.JS(data_str)
}
func build_javascript_uniq_addrs(entries *[]UniqueAddrEntry) template.JS {
	var data_str string = "["

	for i:=0 ; i < len(*entries) ; i++ {
		if len(data_str) > 1 {
			data_str = data_str + ","
		}
		var e = &(*entries)[i];
		var entry string
		entry = "{" +
//				"x:" + fmt.Sprintf("\"%v\"",e.Day)  + "," +
				"x:" + fmt.Sprintf("new Date(%v * 1000)",e.Ts)  + "," +
				"y:"  + fmt.Sprintf("%v",e.NumAddrsAccum) + "," +
				"num_addrs: " + fmt.Sprintf("%v",e.NumAddrs) + "," +
				"num_addrs_accum: " + fmt.Sprintf("%v",e.NumAddrsAccum) + "," +
				"date_str: " + fmt.Sprintf("\"%v\"",e.Day) + "" +
				"}"
		data_str= data_str + entry
	}
	data_str = data_str + "]"
	return template.JS(data_str)
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
	cash_flow_data := build_javascript_cash_flow_data(&cash_flow_entries)
	uniq_addr_entries := augur_srv.storage.Get_unique_addresses()
	uniq_addrs_data := build_javascript_uniq_addrs(&uniq_addr_entries)
	c.HTML(http.StatusOK, "statistics.html", gin.H{
			"title": "Augur Market Statistics",
			"MainStats" : stats,
			"CashFlowData" : cash_flow_data,
			"UniqAddrsData" : uniq_addrs_data,
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

func complete_and_output_market_info(c *gin.Context,minfo InfoMarket) {

	var limit int64 = int64(DEFAILT_MARKET_TRADES_LIMIT);
	p_limit := c.Query("limit")
	if len(p_limit) > 0 {
		var success bool
		limit,success = parse_int_from_remote_or_error(c,&p_limit)
		if !success {
			return
		}
	}
	trades := augur_srv.storage.Get_mkt_trades(minfo.MktAddr,int(limit))
	outcome_vols,_ := augur_srv.storage.Get_outcome_volumes(minfo.MktAddr,minfo.MktAid,0)
	c.HTML(http.StatusOK, "market_info.html", gin.H{
			"title": "Trades for market",
			"Trades" : trades,
			"Market": minfo ,
			"OutcomeVols" : outcome_vols,
	})
}
func market_info(c *gin.Context) {

	market := c.Param("market")

	market_info,_ := augur_srv.storage.Get_market_info(market,0,false)
	complete_and_output_market_info(c,market_info)
}
func full_trade_list(c *gin.Context) {

	market := c.Param("market")

	market_info,_ := augur_srv.storage.Get_market_info(market,0,false)
	trades := augur_srv.storage.Get_mkt_trades(market,0)
	c.HTML(http.StatusOK, "full_trade_list.html", gin.H{
			"title": "Trades for market",
			"Trades" : trades,
			"Market": market_info,
	})
}
func market_depth(c *gin.Context) {

	// Market Depth Info: https://en.wikipedia.org/wiki/Order_book_(trading)
	market := c.Param("market")
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

	market_info,err := augur_srv.storage.Get_market_info(market,outcome,true)
	if err != nil {
		c.HTML(http.StatusBadRequest, "error.html", gin.H{
			"title": "Augur Markets: Error",
			"ErrDescr": "Can't find this market, address not registered",
		})
		return
	}
	mdepth,last_oo_id := augur_srv.storage.Get_mkt_depth(market_info.MktAid,outcome)
	num_orders:=len(mdepth.Bids) + len(mdepth.Asks)
	js_bid_data,js_ask_data := build_javascript_data_obj(mdepth)
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
		outcome,success = parse_int_from_remote_or_error(c,&p_outcome)
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
		market_aid,success = parse_int_from_remote_or_error(c,&p_market_aid)
		if !success {
			return
		}
	} else {
		respond_error(c,"No outcome provided")
		return
	}
	mdepth,last_oo_id := augur_srv.storage.Get_mkt_depth(market_aid,int(outcome))
	js_bid_data,js_ask_data := build_javascript_data_obj(mdepth)
	c.JSON(200,gin.H{
		"bids":js_bid_data,
		"asks":js_ask_data,
		"LastOOID":last_oo_id,
	})
}
func market_price_history(c *gin.Context) {

	market := c.Param("market")
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

	market_info,err := augur_srv.storage.Get_market_info(market,outcome,true)
	if err != nil {
		c.HTML(http.StatusBadRequest, "error.html", gin.H{
			"title": "Augur Markets: Error",
			"ErrDescr": "Can't find this market, address not registered",
		})
		return
	}
	mkt_price_hist := augur_srv.storage.Get_price_history_for_outcome(market_info.MktAid,outcome)
	js_price_history := build_javascript_price_history(&mkt_price_hist)
	fmt.Printf("js price history = %v\n",js_price_history)
	c.HTML(http.StatusOK, "price_history.html", gin.H{
			"title": "Market Price History",
			"Market": market_info,
			"Prices": mkt_price_hist,
			"JSPriceData": js_price_history,
	})
}
func serve_user_info_page(c *gin.Context,addr string,from_wallet bool) {

	eoa_aid,err := augur_srv.storage.Nonfatal_lookup_address_id(addr)
	if err == nil {
		user_info,err := augur_srv.storage.Get_user_info(eoa_aid)
		if err == nil {
			pl_entries := augur_srv.storage.Get_profit_loss(eoa_aid)
			open_pos_entries := augur_srv.storage.Get_open_positions(eoa_aid)
			js_pl_data := build_javascript_profit_loss_history(&pl_entries)
			js_open_pos_data := build_javascript_open_positions(&open_pos_entries)
			user_reports := augur_srv.storage.Get_user_reports(eoa_aid,DEFAULT_USER_REPORTS_LIMIT)
			user_active_markets := augur_srv.storage.Get_active_markets_for_user(eoa_aid)
			var has_active_markets bool = false
			if len(user_active_markets) > 0 {
				has_active_markets = true
			}
			open_orders := augur_srv.storage.Get_user_open_orders(user_info.EOAAid)

			c.HTML(http.StatusOK, "user_info.html", gin.H{
				"title": "User "+addr,
				"user_addr": addr,
				"UserInfo" : user_info,
				"QueriedWallet":from_wallet,
				"PLEntries" : pl_entries,
				"JSPLData" : js_pl_data,
				"JSOpenPosData" : js_open_pos_data,
				"OpenOrders": open_orders,
				"UserReports" : user_reports,
				"UserActiveMarkets" : user_active_markets,
				"HasActiveMarkets" : has_active_markets,
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
func serve_money(c *gin.Context,addr common.Address) {
	// the input address must be EOA, from that we can get Wallet addr

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
			aid,err:=augur_srv.storage.Nonfatal_lookup_address_id(addr_str)
			if err==nil {
				market_info,err := augur_srv.storage.Get_market_info(addr_str,0,false)
				if err == nil {
					complete_and_output_market_info(c,market_info)
					return
				}
				eoa_aid,err:=augur_srv.storage.Lookup_eoa_aid(aid)
				if err==nil {
					// the input was - user's wallet
					eoa_addr,_:=augur_srv.storage.Lookup_address(eoa_aid)
					serve_user_info_page(c,eoa_addr,true)
				} else {
					serve_user_info_page(c,addr_str,false)
				}
				return
			} else {
				c.HTML(http.StatusBadRequest, "error.html", gin.H{
					"title": "Augur Markets: Error",
					"ErrDescr": "Invalid HEX string in address parameter",
				})
				return
			}
		} else {
			c.HTML(http.StatusBadRequest, "error.html", gin.H{
				"title": "Augur Markets: Error",
				"ErrDescr": "Invalid HEX string in address parameter",
			})
			return
		}
	}
	if (len(keyword) == 64) || (len(keyword) == 66) { // Hash (Tx hash)
		if len(keyword) == 66 {	// strip 0x prefix
			keyword = keyword[2:]
		}
		hash_bytes,err := hex.DecodeString(keyword)
		if err == nil {
			hash := common.BytesToHash(hash_bytes)
			hash_str := hash.String()
			serve_tx_info_page(c,hash_str)
		} else {
			c.HTML(http.StatusBadRequest, "error.html", gin.H{
				"title": "Augur Markets: Error",
				"ErrDescr": "Invalid HEX string in hash parameter",
			})
			return
		}
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
			serve_money(c,addr)
		} else {
			c.HTML(http.StatusBadRequest, "error.html", gin.H{
				"title": "Augur Markets: Error",
				"ErrDescr": "Invalid HEX string in address parameter",
			})
			return
		}
	}
}
func order(c *gin.Context) {

	order_hash:= c.Param("order")
	order,err := augur_srv.storage.Get_order_info(order_hash)
	if err!=nil {
		c.HTML(http.StatusBadRequest, "error.html", gin.H{
			"title": "Augur Markets: Error",
			"ErrDescr": "Can't find the order",
		})
		return
	}

	c.HTML(http.StatusOK, "order_info.html", gin.H{
			"title": "Order "+order_hash,
			"OrderInfo" : order,
	})
}
func category(c *gin.Context) {

	p_catid:= c.Param("catid")

	cat_id,success := parse_int_from_remote_or_error(c,&p_catid)
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
		serve_user_info_page(c,addr_str,false)
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
			"ErrDescr": "Invalid HEX string in address parameter",
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
			"ErrDescr": "Invalid HEX string in address parameter",
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
			"ErrDescr": "Invalid HEX string in address parameter",
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
		block_num,success = parse_int_from_remote_or_error(c,&p_block_num)
		if !success {
			return
		}
	}
	block_info,err := augur_srv.storage.Get_block_info(BlockNumber(block_num))
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
		market_aid,success = parse_int_from_remote_or_error(c,&p_market_aid)
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
		outcome_idx,success = parse_int_from_remote_or_error(c,&p_outcome_idx)
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
		last_oo_id,success = parse_int_from_remote_or_error(c,&p_last_oo_id)
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
	if (len(p_addr) == 40) || (len(p_addr) == 42) { // address
		if len(p_addr) == 42 {	// strip 0x prefix
			p_addr = p_addr[2:]
		}
	} else {
		c.HTML(http.StatusBadRequest, "error.html", gin.H{
			"title": "Augur Markets: Error",
			"ErrDescr": "Invalid length of address parameter",
		})
		return
	}
	addr_bytes,err := hex.DecodeString(p_addr)
	if err != nil {
		c.HTML(http.StatusBadRequest, "error.html", gin.H{
			"title": "Augur Markets: Error",
			"ErrDescr": "Invalid HEX string in address parameter for User Account",
		})
		return
	}
	user_addr:= common.BytesToAddress(addr_bytes)
	user_addr_str := user_addr.String()

	p_addr = c.Query("market")
	if (len(p_addr) == 40) || (len(p_addr) == 42) { // address
		if len(p_addr) == 42 {	// strip 0x prefix
			p_addr = p_addr[2:]
		}
	} else {
		c.HTML(http.StatusBadRequest, "error.html", gin.H{
			"title": "Augur Markets: Error",
			"ErrDescr": "Invalid length of address parameter for Market",
		})
		return
	}
	addr_bytes,err = hex.DecodeString(p_addr)
	if err != nil {
		c.HTML(http.StatusBadRequest, "error.html", gin.H{
			"title": "Augur Markets: Error",
			"ErrDescr": "Invalid HEX string in address parameter",
		})
		return
	}
	mkt_addr:= common.BytesToAddress(addr_bytes)
	mkt_addr_str := mkt_addr.String()

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
	}

	market_info,err := augur_srv.storage.Get_market_info(mkt_addr_str,0,false)
	if err!= nil {
		c.HTML(http.StatusBadRequest, "error.html", gin.H{
			"title": "Augur Markets: Error",
			"ErrDescr": fmt.Sprintf("No Market was found with address: %v",mkt_addr_str),
		})
	}
	trades := augur_srv.storage.Get_user_trades_for_market(aid,market_info.MktAid)
	c.HTML(http.StatusOK, "user_trades.html", gin.H{
		"title": fmt.Sprintf("Trades for User %v",user_addr_str),
		"UTrades" : trades,
		"UserInfo" : user_info,
		"Market" : market_info,
	})
}
