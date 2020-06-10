package main
import (
	"fmt"
	"strconv"
	"net/http"
	"encoding/hex"
	"github.com/gin-gonic/gin"
	"html/template"
	"math/big"
	"context"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
/*
	. "github.com/afterether/augur-extractor/primitives"
	. "github.com/afterether/augur-extractor/dbs"
	*/
	. "augur-extractor/primitives"
	. "augur-extractor/dbs"
)
const (
	DEFAULT_DB_LOG_FILE_NAME = "/var/tmp/backend-db.log"
	DEFAULT_MARKET_ROWS_LIMIT int	= 10
	DEFAILT_MARKET_TRADES_LIMIT int = 20
	DEFAULT_USER_REPORTS_LIMIT int = 10
)
type AugurServer struct {
	storage		*SQLStorage
}

func create_augur_server(mkt_order_id_ptr *int64) *AugurServer {

	srv := new(AugurServer)
	srv.storage = Connect_to_storage(mkt_order_id_ptr)
	srv.storage.Init_log(DEFAULT_DB_LOG_FILE_NAME)
	return srv
}
func parse_int_from_remote_or_error(c *gin.Context,ascii_int *string) (int,bool) {
	p, err := strconv.Atoi(*ascii_int)
	if err != nil {
		c.HTML(http.StatusBadRequest, "error.html", gin.H{
			"title": "Augur Markets: Error",
			"ErrDescr": "Can't parse integer parameter (non-numeric characters detected)",
		})
		return 0,false
	}
	result := int(p)
	return result,true
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
	fmt.Printf("asks JS string: %v\n",asks_str)
	fmt.Printf("bids JS string: %v\n",bids_str)
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
					e.SellerEOAAddrSh +"\",\"" +
					e.SellerWalletAddrSh + "\",\"" +
					e.BuyerEOAAddrSh + "\",\"" +
					e.BuyerWalletAddrSh + "\"," +
					fmt.Sprintf("%v,%v,%v,\"%v\"",e.MktAid,e.Price,e.Volume,e.Date) +
				")}" +
				"}"
		fmt.Printf("\nentry = %v\n",entry)
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
		var entry string
		entry = "{" +
				"x:" + fmt.Sprintf("%v",i)  + "," +
				"y:"  + fmt.Sprintf("%v",e.AccumPl) + "," +
				"pl: " + fmt.Sprintf("%v",e.FinalProfit) + "," +
				"pl_accum: " + fmt.Sprintf("%v",e.AccumPl) + "," +
				"date: \"" + fmt.Sprintf("%v",e.Date) + "\"," +
				"click: function() {load_pl_data(" +
					fmt.Sprintf("%v,%v,\"%v\",\"%v\",\"%v\",\"%v\",\"%v\",\"%v\",\"%v\",\"%v\",%v",
								e.FinalProfit,e.AccumPl,e.MktAddr,e.MktAddrSh,e.OutcomeStr,e.MktDescr,
								e.Date,e.CounterPAddr,e.CounterPAddrSh,e.OrderHash,e.BlockNum) +
				")}" +
				"}"
		fmt.Printf("\nentry = %v\n",entry)
		data_str= data_str + entry
	}
	data_str = data_str + "]"
	fmt.Printf("JS profit loss hist string: %v\n",data_str)
	return template.JS(data_str)
}
func build_javascript_open_positions(entries *[]PLEntry) template.JS {
	var data_str string = "["

	for i:=0 ; i < len(*entries) ; i++ {
		if len(data_str) > 1 {
			data_str = data_str + ","
		}
		var e = &(*entries)[i];
		var entry string
		entry = "{" +
				"x:" + fmt.Sprintf("%v",i)  + "," +
				"y:"  + fmt.Sprintf("%v",e.AccumFrozen) + "," +
				"frozen: " + fmt.Sprintf("%v",e.FrozenFunds) + "," +
				"frozen_accum: " + fmt.Sprintf("%v",e.AccumFrozen) + "," +
				"date: \"" + fmt.Sprintf("%v",e.Date) + "\"," +
				"click: function() {load_open_pos_data(" +
					fmt.Sprintf("%v,\"%v\",\"%v\",\"%v\",\"%v\",\"%v\",\"%v\",\"%v\",\"%v\",%v",
							e.FrozenFunds,e.MktAddr,e.MktAddrSh,e.OutcomeStr,e.MktDescr,e.Date,
							e.CounterPAddr,e.CounterPAddrSh,e.OrderHash,e.BlockNum) +
				")}" +
				"}"
		fmt.Printf("\nentry = %v\n",entry)
		data_str= data_str + entry
	}
	data_str = data_str + "]"
	fmt.Printf("JS profit loss hist string: %v\n",data_str)
	return template.JS(data_str)
}
func main_page(c *gin.Context) {
	blknum,_:= augur_srv.storage.Get_last_block_num()
	c.HTML(http.StatusOK, "index.html", gin.H{
			"title": "Augur Prediction Markets",
			"block_num" : blknum,
	})
}
func markets(c *gin.Context) {
	off_str := c.Query("off")
	var off int = 0
	var err error
	fmt.Printf("off_str = %v\n",off_str)
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
	fmt.Printf("off = %v\n",off)
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
	c.HTML(http.StatusOK, "statistics.html", gin.H{
			"title": "Augur Market Statistics",
			"MainStats" : stats,
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

	var limit int = DEFAILT_MARKET_TRADES_LIMIT;
	p_limit := c.Query("limit")
	if len(p_limit) > 0 {
		var success bool
		limit,success = parse_int_from_remote_or_error(c,&p_limit)
		if !success {
			return
		}
	}
	fmt.Printf("markte info addr=%v, full var = %+v\n",minfo.MktAddr,minfo)
	Info.Printf("markte info addr=%v, full var = %+v\n",minfo.MktAddr,minfo)
	trades := augur_srv.storage.Get_mkt_trades(minfo.MktAddr,limit)
	outcome_vols,_ := augur_srv.storage.Get_outcome_volumes(minfo.MktAddr)
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
	fmt.Printf("market info = %+v",market_info)
	complete_and_output_market_info(c,market_info)
}
func full_trade_list(c *gin.Context) {

	market := c.Param("market")

	fmt.Printf("getting trades for market %vi for a full trade listing",market)
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
	mdepth := augur_srv.storage.Get_mkt_depth(market,outcome)
	js_bid_data,js_ask_data := build_javascript_data_obj(mdepth)
	fmt.Printf("js ask_data = %v\n",js_ask_data)
	fmt.Printf("js bid_data = %v\n",js_bid_data)
	c.HTML(http.StatusOK, "market_depth.html", gin.H{
			"title": "Market Depth",
			"Market": market_info,
			"Bids": mdepth.Bids,
			"Asks": mdepth.Asks,
			"JSAskData": js_ask_data,
			"JSBidData": js_bid_data,
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
func serve_user_info_page(c *gin.Context,addr string) {

	eoa_aid,err := augur_srv.storage.Nonfatal_lookup_address_id(addr)
	if err == nil {
		user_info,err := augur_srv.storage.Get_user_info(eoa_aid)
		if err == nil {
			pl_entries := augur_srv.storage.Get_profit_loss(eoa_aid)
			open_pos_entries := augur_srv.storage.Get_open_positions(eoa_aid)
			js_pl_data := build_javascript_profit_loss_history(&pl_entries)
			js_open_pos_data := build_javascript_open_positions(&open_pos_entries)
			user_reports := augur_srv.storage.Get_user_reports(eoa_aid,DEFAULT_USER_REPORTS_LIMIT)
			c.HTML(http.StatusOK, "user_info.html", gin.H{
				"title": "User "+addr,
				"user_addr": addr,
				"UserInfo" : user_info,
				"PLEntries" : pl_entries,
				"JSPLData" : js_pl_data,
				"JSOpenPosData" : js_open_pos_data,
				"UserReports" : user_reports,
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

	c.HTML(http.StatusOK, "tx_info.html", gin.H{
			"title": "Transaction " + tx_hash,
			"tx_hash" : tx_hash,
	})
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
	fmt.Printf("token.BalanceOf() returns: %v\n",int_balance.String())
	if err == nil {
		f_balance :=big.NewFloat(0.0)
		f_balance.SetInt(int_balance)
		divisor:=big.NewFloat(0.0)
		divisor.SetString("1000000000000000000.0")
		div_result:=new(big.Float).Quo(f_balance,divisor)
		fmt.Printf("div_result=%v\n",div_result.String())
		big_float_balance.Set(div_result)
		fmt.Printf("big_float_balance=%v\n",big_float_balance)
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
		fmt.Printf("eth_balance big int: %v\n",big_eth_balance.String())
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
		fmt.Printf("wallet_aid lookup failed for eoa_aid=%v\n",eoa_aid)
		c.JSON(200,gin.H{
			"eoa_eth":0,"wallet_eth":0,"eoa_dai":0,"wallet_dai":0,"eoa_rep":0,"wallet_rep":0,
		})
		return
	}
	fmt.Printf("money for addr %v, eoa_aid=%v\n",addr.String(),eoa_aid)
	eoa_dai_balance := get_token_balance(0,&addr)
	eoa_rep_balance := get_token_balance(1,&addr)
	eoa_eth_balance := get_eth_balance(&addr)

	var wallet_dai_balance float64 = 0.0
	var wallet_rep_balance float64 = 0.0
	var wallet_eth_balance float64 = 0.0

	fmt.Printf("wallet_aid=%v\n",wallet_aid)
	if wallet_aid != 0 {
		wallet_addr,err := augur_srv.storage.Lookup_address(wallet_aid)
		if err == nil {
			waddr := common.HexToAddress(wallet_addr)
			fmt.Printf("wallet addr = %v\n",wallet_addr)
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
	fmt.Printf("Searching for %v\n",keyword)
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
				fmt.Printf("checking address in market info: %v\n",addr_str)
				market_info,err := augur_srv.storage.Get_market_info(addr_str,0,false)
				if err == nil {
					complete_and_output_market_info(c,market_info)
					return
				}
				eoa_aid,err:=augur_srv.storage.Lookup_eoa_aid(aid)
				if err==nil {
					// the input was - user's wallet
					eoa_addr,_:=augur_srv.storage.Lookup_address(eoa_aid)
					serve_user_info_page(c,eoa_addr)
				} else {
					serve_user_info_page(c,addr_str)
				}
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
		if err != nil {
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
		fmt.Printf("addr=%v, aid=%v\n",addr_str,aid)
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
