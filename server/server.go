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
const (
	DEFAULT_DB_LOG_FILE_NAME = "/var/tmp/backend-db.log"
	DEFAULT_AMM_LOG_FILE_NAME = ""
	DEFAULT_MARKET_ROWS_LIMIT int	= 500
	DEFAILT_MARKET_TRADES_LIMIT int = 20
	DEFAULT_USER_REPORTS_LIMIT int = 30
	DEFAULT_MARKET_REPORTS_LIMIT int = 40
)
type AugurServer struct {
	db_augur		*SQLStorage
	db_matic		*SQLStorage
}
func (self *AugurServer) matic_initialized() bool {

	if self.db_matic == nil {
		return false
	}
	return true
}
func create_augur_server() *AugurServer {

	log_dir:=fmt.Sprintf("%v/%v",os.Getenv("HOME"),DEFAULT_LOG_DIR)
	os.MkdirAll(log_dir, os.ModePerm)
	db_log_file:=fmt.Sprintf("%v/%v",log_dir,"webserver-db.log")

	fname:=fmt.Sprintf("%v/webserver_info.log",log_dir)
	logfile, err := os.OpenFile(fname, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err!=nil {
		fmt.Printf("Can't start: %v\n",err)
		os.Exit(1)
	}
	Info = log.New(logfile,"INFO: ",log.Ldate|log.Ltime|log.Lshortfile)

	fname=fmt.Sprintf("%v/webserver_error.log",log_dir)
	logfile, err = os.OpenFile(fname, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err!=nil {
		fmt.Printf("Can't start: %v\n",err)
		os.Exit(1)
	}
	Error = log.New(logfile,"ERROR: ",log.Ldate|log.Ltime|log.Lshortfile)
	srv := new(AugurServer)
	srv.db_augur= Connect_to_storage(&market_order_id,Info)
	srv.db_augur.Init_log(db_log_file)
	amm_user := os.Getenv("AMM_USERNAME")
	amm_passwd := os.Getenv("AMM_PASSWORD")
	amm_db_name := os.Getenv("AMM_DATABASE")
	amm_host_port := os.Getenv("AMM_HOST")
	if len(amm_user) > 0 {
		srv.db_augur.Init_log(db_log_file)

		db_log_file:=fmt.Sprintf("%v/%v",log_dir,"amm-db.log")
		amm_db_logfile, err := os.OpenFile(db_log_file, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
		if err != nil {
			fmt.Printf("Can't start: %v\n",err)
			os.Exit(1)
		}
		AMM_DB := log.New(amm_db_logfile,"INFO: ",log.Ldate|log.Ltime|log.Lshortfile)

		srv.db_matic = New_sql_storage(
			&market_order_id,
			Info,
			AMM_DB,
			amm_host_port,
			amm_db_name,
			amm_user,
			amm_passwd,
		)
	}

	return srv
}
func respond_error(c *gin.Context,error_text string) {

	c.HTML(http.StatusBadRequest, "error.html", gin.H{
		"title": "Augur Markets: Error",
		"ErrDescr": error_text,
	})
}
func respond_error_json(c *gin.Context,error_text string) {

	c.JSON(http.StatusBadRequest, gin.H{
		"status": 0,
		"error": error_text,
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
	blknum,_:= augur_srv.db_augur.Get_last_block_num()
	blknum_thousand_separated := ThousandsFormat(int64(blknum))
	stats := augur_srv.db_augur.Get_front_page_stats()
	c.HTML(http.StatusOK, "augur_v2/index.html", gin.H{
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
	markets := augur_srv.db_augur.Get_active_market_list(off,DEFAULT_MARKET_ROWS_LIMIT)
	c.HTML(http.StatusOK, "markets.html", gin.H{
			"title": "Augur Markets",
			"Markets" : markets,
	})
}
func is_address_valid(c *gin.Context,json_output bool,addr string) (string,bool) {

	if (len(addr) != 40) && (len(addr)!=42) {
		var err_msg = fmt.Sprintf("Provided address has invalid length (len=%v)",len(addr))
		if json_output {
			c.JSON(http.StatusOK,gin.H{
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
			c.JSON(http.StatusOK,gin.H{
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
			c.JSON(http.StatusOK,gin.H{
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
			addr,err = augur_srv.db_augur.Lookup_address(aid)
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
	aid,err := augur_srv.db_augur.Nonfatal_lookup_address_id(address)
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
func get_REP_token_price_in_ETH() (float64,error) {

	// token0 - REP (0x221657776846890989a759BA2973e427DfF5C9bB)
	// token1 - Wrapped ETH (0xC02aaA39b223FE8D0A0e5C4F27eAD9083C756Cc2)
	addr := common.HexToAddress(REP_ETH_UNISWAP_PAIR_ADDR)
	ctrct_pair,err := NewUniswapV2Pair(addr,rpcclient)
	if err != nil {
		return 0.0,err
	}
	var copts = new(bind.CallOpts)
	reserves,err := ctrct_pair.GetReserves(copts)
	if err != nil {
		return 0.0,err
	}
	price_f := big.NewFloat(0.0)
	price_f.SetString(reserves.Reserve0.String()) // REP
	eth_f := big.NewFloat(0.0)
	eth_f.SetString(reserves.Reserve1.String()) // WETH
	price_f.Quo(price_f,eth_f)
	output,_ := price_f.Float64()
	return output,nil
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
func has0xPrefix(input string) bool {
	return len(input) >= 2 && input[0] == '0' && (input[1] == 'x' || input[1] == 'X')
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
			addr_str := addr.String()
			serve_user_funds_v2(c,&addr_str)
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
			"ErrDescr": fmt.Sprintf("Invalid address len (must be 40 or 42 chars) len=%v",len(addr)),
		})
		return
	}
}
func wrapped_tokens(c *gin.Context) {

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
	wrappers := augur_srv.db_augur.Get_wrapped_tokens_for_market(market_info.MktAid)
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
	aid,err := augur_srv.db_augur.Nonfatal_lookup_address_id(addr)
	if err!=nil {
		c.HTML(http.StatusBadRequest, "error.html", gin.H{
			"title": "Augur Markets: Error",
			"ErrDescr": fmt.Sprintf("Such address wasn't found: %v",address),
		})
		return
	}
	wrapper_info,_ := augur_srv.db_augur.Get_wrapped_token_info(aid)
	market_info,err := augur_srv.db_augur.Get_market_info(wrapper_info.MktAddr,wrapper_info.OutcomeIdx,true)
	transfers,total_rows := augur_srv.db_augur.Get_wrapped_token_transfers(aid,0,500)
	c.HTML(http.StatusOK, "wrapped_sharetokens/transfers.html", gin.H{
			"MarketInfo" : market_info,
			"TokenInfo" : wrapper_info,
			"TotalRows" : total_rows,
			"WrappedTransfers" : transfers,
	})
}
func user_wrapped_token_transfers(c *gin.Context) {

	p_user:= c.Param("user")
	user_addr,valid := is_address_valid(c,false,p_user)
	if !valid {
		return
	}
	p_wrapper:= c.Param("wrapper")
	wrapper_addr,valid := is_address_valid(c,false,p_wrapper)
	if !valid {
		return
	}
	user_aid,err := augur_srv.db_augur.Nonfatal_lookup_address_id(user_addr)
	if err!=nil {
		c.HTML(http.StatusBadRequest, "error.html", gin.H{
			"title": "Augur Markets: Error",
			"ErrDescr": fmt.Sprintf("Such address wasn't found: %v",user_addr),
		})
		return
	}
	wrapper_aid,err := augur_srv.db_augur.Nonfatal_lookup_address_id(wrapper_addr)
	if err!=nil {
		c.HTML(http.StatusBadRequest, "error.html", gin.H{
			"title": "Augur Markets: Error",
			"ErrDescr": fmt.Sprintf("Such address wasn't found: %v",wrapper_addr),
		})
		return
	}
	wrapper_info,_ := augur_srv.db_augur.Get_wrapped_token_info(wrapper_aid)
	market_info,err := augur_srv.db_augur.Get_market_info(wrapper_info.MktAddr,wrapper_info.OutcomeIdx,true)
	total_rows,transfers:= augur_srv.db_augur.Get_user_wrapped_shtoken_transfers(user_aid,wrapper_aid,0,10000)
	c.HTML(http.StatusOK, "user_wrapped_transfers.html", gin.H{
			"UserAddr" : user_addr,
			"MarketInfo" : market_info,
			"TokenInfo" : wrapper_info,
			"Transfers" : transfers,
			"TotalRows" : total_rows,
	})
}
func pool_swaps(c *gin.Context) {

	address:= c.Param("address")
	addr,valid := is_address_valid(c,false,address)
	if !valid {
		return
	}
	aid,err := augur_srv.db_augur.Nonfatal_lookup_address_id(addr)
	if err!=nil {
		c.HTML(http.StatusBadRequest, "error.html", gin.H{
			"title": "Augur Markets: Error",
			"ErrDescr": fmt.Sprintf("Such address wasn't found: %v",address),
		})
		return
	}
	pool_info,_ := augur_srv.db_augur.Get_pool_info(aid)
	swaps := augur_srv.db_augur.Get_pool_swaps(aid,0,200)
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
	minfo,err := augur_srv.db_augur.Get_market_info(market_addr,0,false)
	if err != nil {
		show_market_not_found_error(c,false,&market_addr)
		return
	}

	outag_sh_bal_chgs,total_rows := augur_srv.db_augur.Outside_augur_share_balance_changes(minfo.MktAid,0,500)
	c.HTML(http.StatusOK, "sharetoken_balance_changes.html", gin.H{
			"MarketInfo" : minfo,
			"TotalRows" : total_rows,
			"OutsideAugurBalanceChanges": outag_sh_bal_chgs,
	})
}
func market_uniswap_pairs(c *gin.Context) {

	market := c.Param("market")
	market_addr,valid:=is_address_valid(c,false,market)
	if !valid {
		return
	}
	minfo,err := augur_srv.db_augur.Get_market_info(market_addr,0,false)
	if err != nil {
		show_market_not_found_error(c,false,&market_addr)
		return
	}
	pairs := augur_srv.db_augur.Get_market_uniswap_pairs(minfo.MktAid)
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
	aid,err := augur_srv.db_augur.Nonfatal_lookup_address_id(addr)
	if err!=nil {
		c.HTML(http.StatusBadRequest, "error.html", gin.H{
			"title": "Augur Markets: Error",
			"ErrDescr": fmt.Sprintf("Such address wasn't found: %v",address),
		})
		return
	}
	now_ts := time.Now().Unix()
	past_ts := now_ts - 100 * 3600 * 24
	pair_info,_:= augur_srv.db_augur.Get_uniswap_pair_info(aid)
	swaps := augur_srv.db_augur.Get_uniswap_swaps(aid,0,200)
	c.HTML(http.StatusOK, "uniswap_swaps.html", gin.H{
			"PairInfo" : pair_info,
			"PairSwaps" : swaps,
			"SampleFinTs" : now_ts,
			"SampleInitTs" : past_ts,
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
	search_results := augur_srv.db_augur.Search_keywords_in_markets(keywords)
	c.HTML(http.StatusOK, "text_search_results.html", gin.H{
			"SearchResults" : search_results,
	})
}
func show_text_search_form(c *gin.Context) {

	c.HTML(http.StatusOK, "text_search_form.html", gin.H{
	})
}
func show_pool_swap_prices(c *gin.Context) {

	p_pool_aid := c.Param("pool_aid")
	var pool_aid int64
	if len(p_pool_aid) > 0 {
		var success bool
		pool_aid,success = parse_int_from_remote_or_error(c,false,&p_pool_aid)
		if !success {
			return
		}
	} else {
		respond_error(c,"Pool ID is not set")
		return
	}
	p_token1_aid := c.Param("token1_aid")
	var token1_aid int64
	if len(p_token1_aid) > 0 {
		var success bool
		token1_aid,success = parse_int_from_remote_or_error(c,false,&p_token1_aid)
		if !success {
			return
		}
	} else {
		respond_error(c,"Token1 ID is not set")
		return
	}
	p_token2_aid := c.Param("token2_aid")
	var token2_aid int64
	if len(p_token2_aid) > 0 {
		var success bool
		token2_aid,success = parse_int_from_remote_or_error(c,false,&p_token2_aid)
		if !success {
			return
		}
	} else {
		respond_error(c,"Token2 ID is not set")
		return
	}
	var err error
	p_init_ts := c.Param("init_ts")
	var init_ts int
	if len(p_init_ts) > 0 {
		init_ts, err = strconv.Atoi(p_init_ts)
		if err != nil {
			c.HTML(http.StatusBadRequest, "error.html", gin.H{
				"title": "Augur Markets: Error",
				"ErrDescr": "Can't parse init_ts",
			})
			return
		}
	}
	p_fin_ts := c.Param("fin_ts")
	var fin_ts int
	if len(p_fin_ts) > 0 {
		fin_ts, err = strconv.Atoi(p_fin_ts)
		if err != nil {
			c.HTML(http.StatusBadRequest, "error.html", gin.H{
				"title": "Augur Markets: Error",
				"ErrDescr": "Can't parse fin_ts",
			})
			return
		}
	}
	if fin_ts == 0 {
		fin_ts = 2147483647
	}
	p_interval_secs := c.Param("interval_secs")
	var interval_secs int = 0
	if len(p_interval_secs) > 0 {
		interval_secs, err = strconv.Atoi(p_interval_secs)
		if err != nil {
			c.HTML(http.StatusBadRequest, "error.html", gin.H{
				"title": "Augur Markets: Error",
				"ErrDescr": "Can't parse 'interval_secs' param",
			})
			return
		}
	}
	if interval_secs == 0 {
		interval_secs = 60*60
	}

	pool_info,_ := augur_srv.db_augur.Get_pool_info(pool_aid)
	token1_info,_ := augur_srv.db_augur.Get_bpool_token_info(pool_aid,token1_aid)
	token2_info,_ := augur_srv.db_augur.Get_bpool_token_info(pool_aid,token2_aid)
	prices := augur_srv.db_augur.Get_balancer_token_prices(pool_aid,token1_aid,token2_aid,init_ts,fin_ts,interval_secs)
	js_prices := build_js_bpool_swap_prices(&prices)
	c.HTML(http.StatusOK, "bswap_prices.html", gin.H{
			"PoolInfo" : pool_info,
			"Token1Info" : token1_info,
			"Token2Info" : token2_info,
			"Prices" : prices,
			"JSPriceData" :js_prices,
			"InitTimeStamp": init_ts,
			"FinTimeSTamp": fin_ts,
			"IntervalSecs": interval_secs,
	})
}
func show_upair_swap_prices(c *gin.Context) {

	p_pair_aid := c.Param("pair_aid")
	var pair_aid int64
	if len(p_pair_aid) > 0 {
		var success bool
		pair_aid,success = parse_int_from_remote_or_error(c,false,&p_pair_aid)
		if !success {
			return
		}
	} else {
		respond_error(c,"Pair ID is not set")
		return
	}
	p_inverse := c.Param("inverse")
	var inverse int64
	if len(p_inverse) > 0 {
		var success bool
		inverse,success = parse_int_from_remote_or_error(c,false,&p_inverse)
		if !success {
			return
		}
	} else {
		respond_error(c,"'inverse' parameter is not set")
		return
	}
	var err error
	p_init_ts := c.Param("init_ts")
	var init_ts int
	if len(p_init_ts) > 0 {
		init_ts, err = strconv.Atoi(p_init_ts)
		if err != nil {
			c.HTML(http.StatusBadRequest, "error.html", gin.H{
				"title": "Augur Markets: Error",
				"ErrDescr": "Can't parse init_ts",
			})
			return
		}
	}
	p_fin_ts := c.Param("fin_ts")
	var fin_ts int
	if len(p_fin_ts) > 0 {
		fin_ts, err = strconv.Atoi(p_fin_ts)
		if err != nil {
			c.HTML(http.StatusBadRequest, "error.html", gin.H{
				"title": "Augur Markets: Error",
				"ErrDescr": "Can't parse fin_ts",
			})
			return
		}
	}
	if fin_ts == 0 {
		fin_ts = 2147483647
	}

	p_interval_secs := c.Param("interval_secs")
	var interval_secs int = 0
	if len(p_interval_secs) > 0 {
		interval_secs, err = strconv.Atoi(p_interval_secs)
		if err != nil {
			c.HTML(http.StatusBadRequest, "error.html", gin.H{
				"title": "Augur Markets: Error",
				"ErrDescr": "Can't parse 'interval_secs' param",
			})
			return
		}
	}
	if interval_secs == 0 {
		interval_secs = 60*60
	}

	bool_inverse := false
	if inverse > 0 {
		bool_inverse = true
	}
	pair_info,_:= augur_srv.db_augur.Get_uniswap_pair_info(pair_aid)
	prices := augur_srv.db_augur.Get_uniswap_token_prices(pair_aid,bool_inverse,init_ts,fin_ts,interval_secs)
	js_prices := build_js_upair_swap_prices(&prices)
	c.HTML(http.StatusOK, "upair_prices.html", gin.H{
			"PairInfo" : pair_info,
			"Prices" : prices,
			"JSPriceData" :js_prices,
			"InitTimeStamp": init_ts,
			"FinTimeSTamp": fin_ts,
	})
}
func show_single_uniswap_swap(c *gin.Context) {

	p_id:= c.Param("id")
	var id int64
	if len(p_id) > 0 {
		var success bool
		id,success = parse_int_from_remote_or_error(c,false,&p_id)
		if !success {
			return
		}
	} else {
		respond_error(c,"'id' parameter is not set")
		return
	}

	swap,err := augur_srv.db_augur.Get_uniswap_swap_by_id(id)
	if err!=nil {
		c.HTML(http.StatusBadRequest, "error.html", gin.H{
			"title": "Augur Markets: Error",
			"ErrDescr": fmt.Sprintf("Error getting swap with id=%v : %v",id,err),
		})
		return
	}

	c.HTML(http.StatusOK, "single_uniswap_swap.html", gin.H{
			"UniswapSwap" : swap,
			"Id": id,
	})
}
func show_single_balancer_swap(c *gin.Context) {

	p_id:= c.Param("id")
	var id int64
	if len(p_id) > 0 {
		var success bool
		id,success = parse_int_from_remote_or_error(c,false,&p_id)
		if !success {
			return
		}
	} else {
		respond_error(c,"'id' parameter is not set")
		return
	}

	swap,err := augur_srv.db_augur.Get_balancer_swap_by_id(id)
	if err!=nil {
		c.HTML(http.StatusBadRequest, "error.html", gin.H{
			"title": "Augur Markets: Error",
			"ErrDescr": fmt.Sprintf("Error getting swap with id=%v : %v",id,err),
		})
		return
	}

	c.HTML(http.StatusOK, "single_balancer_swap.html", gin.H{
			"BalancerSwap" : swap,
			"Id": id,
	})
}
func wrapped_token_info(c *gin.Context) {

	p_address := c.Param("address")
	wrapper_addr,valid:=is_address_valid(c,false,p_address)
	if !valid {
		return
	}
	wrapper_aid,err := augur_srv.db_augur.Nonfatal_lookup_address_id(wrapper_addr)
	if err == nil {
		respond_error(c,fmt.Sprintf("Address %v not found",p_address))
		return
	}
	winfo,err := augur_srv.db_augur.Get_wrapped_token_info(wrapper_aid)
	if err != nil {
		respond_error(c,fmt.Sprintf("ShareToken wrapper with address %v not found",p_address))
		return
	}
	c.HTML(http.StatusOK, "wrapped_sharetokens/token_info.html", gin.H{
		"WrapperInfo" : winfo,
	})
}
func balancer_calc_slippage(addr_str string,token_in_str string,token_out_str string,amount_str string) (*big.Int,*big.Int,error) {

	addr := common.HexToAddress(addr_str)
	token_in := common.HexToAddress(token_in_str)
	token_out := common.HexToAddress(token_out_str)
	ctrct_bpool,err := NewBPool(addr,rpcclient)
	if err != nil {
		return nil,nil,err
	}
	var copts = new(bind.CallOpts)
	ten := big.NewInt(10)
	max_price := big.NewInt(0)


	token_in_balance,err := ctrct_bpool.GetBalance(copts,token_in)
	if err != nil {
		return nil,nil,err
	}
	token_out_balance,err := ctrct_bpool.GetBalance(copts,token_out)
	if err != nil {
		return nil,nil,err
	}
	token_in_weight,err := ctrct_bpool.GetDenormalizedWeight(copts,token_in)
	if err != nil {
		return nil,nil,err
	}
	token_out_weight,err := ctrct_bpool.GetDenormalizedWeight(copts,token_out)
	if err != nil {
		return nil,nil,err
	}
	swap_fee,err := ctrct_bpool.GetSwapFee(copts)
	if err != nil {
		return nil,nil,err
	}
	spot_price,err := ctrct_bpool.CalcSpotPrice(copts,token_in_balance,token_in_weight,token_out_balance,token_out_weight,swap_fee)
	max_price.Mul(spot_price,ten)

	amount := big.NewInt(0)
	amount.SetString(amount_str,10)
	token_amount_out,err := ctrct_bpool.CalcOutGivenIn(copts,token_in_balance,token_in_weight,token_out_balance,token_out_weight,amount,swap_fee)
	if err != nil {
		return nil,nil,err
	}
	new_in_balance := big.NewInt(0)
	new_in_balance.Set(token_in_balance)
	new_in_balance.Add(new_in_balance,amount)
	new_out_balance := big.NewInt(0)
	new_out_balance.Set(token_out_balance)
	new_out_balance.Add(new_out_balance,token_amount_out)
	spot_price_after,err := ctrct_bpool.CalcSpotPrice(copts,new_in_balance,token_in_weight,new_out_balance,token_out_weight,swap_fee)
	if err != nil {
		return nil,nil,err
	}
	slippage := big.NewInt(0)
	slippage.Sub(spot_price,spot_price_after)
	return slippage,token_amount_out,nil
}
func produce_pool_slippages(amount_to_trade string,pool_aid int64) []TokenSlippage {

	tokens := augur_srv.db_augur.Get_balancer_pool_tokens_for_slippage(pool_aid)
	for i:=0; i < len(tokens) ; i++ {
		t := &tokens[i]
		amount := fmt.Sprintf("%v%0*d",amount_to_trade,t.Decimals1, 0)
		slippage,amount_token_out,_:= balancer_calc_slippage(
			t.PoolAddr,
			t.Token1Addr,
			t.Token2Addr,
			amount,
		)
		if slippage != nil {
			fslippage := big.NewFloat(0.0)
			fslippage.SetString(slippage.String())
			divisor1_str := fmt.Sprintf("1%0*d", t.Decimals1, 0)
			divisor2_str := fmt.Sprintf("1%0*d", t.Decimals2, 0)
			divisor1 := big.NewFloat(0.0)
			divisor1.SetString(divisor1_str)
			divisor2 := big.NewFloat(0.0)
			divisor2.SetString(divisor2_str)
			quo := big.NewFloat(0.0)
			quo.Quo(fslippage,divisor1)
			resulting_slippage,_ := quo.Float64()
			t.Slippage = resulting_slippage
			famount := big.NewFloat(0.0)
			famount.SetString(amount)
			famount.Quo(famount,divisor1)
			t.AmountIn,_ = famount.Float64()
			famount.SetString(amount_token_out.String())
			famount.Quo(famount,divisor2)
			t.AmountOut,_ = famount.Float64()
		}
	}
	return tokens
}
func show_pool_slippage(c *gin.Context) {

	p_pool:= c.Param("pool")
	pool_addr,valid:=is_address_valid(c,false,p_pool)
	if !valid {
		return
	}
	pool_aid,err := augur_srv.db_augur.Nonfatal_lookup_address_id(pool_addr)
	if err != nil {
		respond_error(c,fmt.Sprintf("Address %v not found",))
		return
	}
	pool_info,_ := augur_srv.db_augur.Get_pool_info(pool_aid)
	//amount_to_trade := "100";
	//tokens := produce_pool_slippages(amount_to_trade,pool_aid)
	tokens := augur_srv.db_augur.Get_balancer_latest_slippages(pool_aid)
	var amount_to_trade float64 = 0.0
	if len(tokens) > 0 {
		amount_to_trade = tokens[0].AmountIn
	}
	c.HTML(http.StatusOK, "pool_slippage.html", gin.H{
		"PoolInfo" : pool_info,
		"TokenSlippage" : tokens,
		"AmountToTrade" : amount_to_trade,
	})
}
func uniswap_correct_for_difference_in_decimals(value *big.Float,decimals1,decimals2 int) {
	if decimals1 != decimals2 {
		var dec_diff int = 0
		if decimals1 < decimals2 {
			dec_diff = decimals2 - decimals1;
			divisor_str := fmt.Sprintf("1%0*d",dec_diff, 0)
			divisor_big := big.NewFloat(0.0)
			divisor_big.SetString(divisor_str)
			value.Quo(value,divisor_big)
		} else {
			dec_diff = decimals1 - decimals2;
			multiplier_str := fmt.Sprintf("1%0*d",dec_diff, 0)
			multiplier_big := big.NewFloat(0.0)
			multiplier_big.SetString(multiplier_str)
			value.Mul(value,multiplier_big)
		}
	}
}
func uniswap_calc_slippage(pair_addr_str string,token_str string,amount_str string) (*big.Float,*big.Int,error) {
	// note: we are receiving decimals as parameter because the fetch porcess to get decimals from the
	//		contract is more complicated than just calling Decimals() on the contract. The code to
	//		fetch all token info is at primitives/augur_utils.go 
	//		the decimals should be provided from the DB

	addr := common.HexToAddress(pair_addr_str)
	qtoken := common.HexToAddress(token_str)

	ctrct_pair,err := NewUniswapV2Pair(addr,rpcclient)
	if err != nil {
		return nil,nil,err
	}
	var copts = new(bind.CallOpts)
	reserves,err := ctrct_pair.GetReserves(copts)
	if err != nil {
		return nil,nil,err
	}
	token0,err := ctrct_pair.Token0(copts)
	if err != nil {
		return nil,nil,err
	}
	token1,err := ctrct_pair.Token1(copts)
	if err != nil {
		return nil,nil,err
	}
	_=token1
	var r1,r2 *big.Int
	if bytes.Equal(qtoken.Bytes(),token0.Bytes()) {
		r1=reserves.Reserve0
		r2=reserves.Reserve1
	} else {
		r1=reserves.Reserve1
		r2=reserves.Reserve0
	}
	_,_,router02_addr_str := augur_srv.db_augur.Get_uniswap_contracts()
	router02_addr := common.HexToAddress(router02_addr_str)
	ctrct_router,err := NewUniswapV2Router02(router02_addr,rpcclient)
	amount := big.NewInt(0)
	amount.SetString(amount_str,10)
	token_amount_out,err := ctrct_router.GetAmountOut(copts,amount,r1,r2)

	// calculate spot price before swap
	spot_price_before := big.NewFloat(0.0)
	r1_float := big.NewFloat(0.0)
	r1_float.SetString(r1.String())
	r2_float := big.NewFloat(0.0)
	r2_float.SetString(r2.String())
	spot_price_before.Quo(r1_float,r2_float)

	r1big := big.NewInt(0)
	r2big := big.NewInt(0)
	r1big.Set(r1)
	r1big.Add(r1big,amount)
	r2big.Sub(r2,token_amount_out)
	spot_price_after := big.NewFloat(0.0)
	r1_float = big.NewFloat(0.0)
	r1_float.SetString(r1.String())
	amount_float := big.NewFloat(0.0)
	amount_float.SetString(amount.String())
	r1_float.Add(r1_float,amount_float)
	r2_float = big.NewFloat(0.0)
	r2_float.SetString(r2.String())
	token_out_float := big.NewFloat(0.0)
	r2_float.Sub(r2_float,token_out_float)
	spot_price_after.Quo(r1_float,r2_float)

	slippage_float:= big.NewFloat(0.0)
	slippage_float.Sub(spot_price_after,spot_price_before)
	return slippage_float,token_amount_out,nil
}
func produce_uniswap_slippages(pi *MarketUPair,amount_str string) ([]TokenSlippage,error) {

	output := make([]TokenSlippage,0,2)
	{
		var ts TokenSlippage
		ts.Token1Addr = pi.Token0Addr
		ts.Token2Addr = pi.Token1Addr
		ts.Token1Symbol = pi.Token0Symbol
		ts.Token2Symbol = pi.Token1Symbol
		ts.Decimals1 = pi.Token0Decimals
		ts.Decimals2 = pi.Token1Decimals
		ts.PoolAddr = pi.PairAddr
		ts.NumSwaps = pi.TotalSwaps
		in_float := big.NewFloat(0.0)
		in_float.SetString(amount_str)
		ts.AmountIn,_ = in_float.Float64()

		amount := fmt.Sprintf("%v%0*d",amount_str,ts.Decimals1,0)
		slippage,token_amount_out,err := uniswap_calc_slippage(pi.PairAddr,ts.Token1Addr,amount)
		if err != nil {
			return output,err
		}
		uniswap_correct_for_difference_in_decimals(slippage,ts.Decimals2,ts.Decimals1)
		ts.Slippage,_ = slippage.Float64()

		famount := big.NewFloat(0.0)
		famount.SetString(token_amount_out.String())
		divisor := fmt.Sprintf("%v%0*d",1,ts.Decimals2,0)
		fdivisor := big.NewFloat(0.0)
		fdivisor.SetString(divisor)
		famount.Quo(famount,fdivisor)
		ts.AmountOut,_ = famount.Float64()

		output = append(output,ts)
	}
	{
		var ts TokenSlippage
		ts.Token1Addr = pi.Token1Addr
		ts.Token2Addr = pi.Token0Addr
		ts.Token1Symbol = pi.Token1Symbol
		ts.Token2Symbol = pi.Token0Symbol
		ts.Decimals1 = pi.Token1Decimals
		ts.Decimals2 = pi.Token0Decimals
		ts.PoolAddr = pi.PairAddr
		ts.NumSwaps = pi.TotalSwaps
		in_float := big.NewFloat(0.0)
		in_float.SetString(amount_str)
		ts.AmountIn,_ = in_float.Float64()

		amount := fmt.Sprintf("%v%0*d",amount_str,ts.Decimals1,0)
		slippage,token_amount_out,err := uniswap_calc_slippage(pi.PairAddr,ts.Token1Addr,amount)
		if err != nil {
			return output,err
		}
		uniswap_correct_for_difference_in_decimals(slippage,ts.Decimals2,ts.Decimals1)
		ts.Slippage,_ = slippage.Float64()

		famount := big.NewFloat(0.0)
		famount.SetString(token_amount_out.String())
		divisor := fmt.Sprintf("%v%0*d",1,ts.Decimals2,0)
		fdivisor := big.NewFloat(0.0)
		fdivisor.SetString(divisor)
		famount.Quo(famount,fdivisor)
		ts.AmountOut,_ = famount.Float64()

		output = append(output,ts)
	}
	return output,nil
}
func show_uniswap_slippage(c *gin.Context) {

	p_pair:= c.Param("pair")
	pair_addr,valid:=is_address_valid(c,false,p_pair)
	if !valid {
		return
	}
	pair_aid,err := augur_srv.db_augur.Nonfatal_lookup_address_id(pair_addr)
	if err != nil {
		respond_error(c,fmt.Sprintf("Address %v not found",))
		return
	}
	pair_info,err := augur_srv.db_augur.Get_uniswap_pair_info(pair_aid)
	if err != nil {
		respond_error(c,err.Error())
		return
	}
	//slippages,err := produce_uniswap_slippages(&pair_info,amount_to_trade)
	slippages := augur_srv.db_augur.Get_uniswap_latest_slippages(pair_aid)
	//amount_to_trade := "100";
	var amount_to_trade float64 = 0.0
	if len(slippages) > 0 {
		amount_to_trade = slippages[0].AmountIn
	}
	if err!=nil {
		c.JSON(http.StatusBadRequest,gin.H{
			"status":0,
			"error": err.Error(),
		})
		return
	}

	c.HTML(http.StatusOK, "uniswap/slippages.html", gin.H{
		"PairInfo" : pair_info,
		"TokenSlippage" : slippages,
		"AmountToTrade" : amount_to_trade,
	})
}
func rt_show_uniswap_slippage(c *gin.Context) {

	p_pair:= c.Param("pair")
	pair_addr,valid:=is_address_valid(c,false,p_pair)
	if !valid {
		return
	}
	pair_aid,err := augur_srv.db_augur.Nonfatal_lookup_address_id(pair_addr)
	if err != nil {
		respond_error(c,fmt.Sprintf("Address %v not found",))
		return
	}
	pair_info,err := augur_srv.db_augur.Get_uniswap_pair_info(pair_aid)
	if err != nil {
		respond_error(c,err.Error())
		return
	}
	amount_to_trade := "100";
	slippages,err := produce_uniswap_slippages(&pair_info,amount_to_trade)
	if err!=nil {
		c.JSON(http.StatusBadRequest,gin.H{
			"status":0,
			"error": err.Error(),
		})
		return
	}

	c.HTML(http.StatusOK, "uniswap_slippages.html", gin.H{
		"PairInfo" : pair_info,
		"TokenSlippage" : slippages,
		"AmountToTrade" : amount_to_trade,
	})
}
func show_ethusd_price(c *gin.Context) {

	var err error
	p_init_ts := c.Query("init_ts")
	var init_ts int
	if len(p_init_ts) > 0 {
		init_ts, err = strconv.Atoi(p_init_ts)
		if err != nil {
			c.HTML(http.StatusBadRequest, "error.html", gin.H{
				"title": "Augur Markets: Error",
				"ErrDescr": "Can't parse init_ts",
			})
			return
		}
	}
	if init_ts == 0 {
		init_ts = int(time.Now().Unix())
		init_ts = init_ts - 30 * 24 * 60* 60
	}
	p_fin_ts := c.Query("fin_ts")
	var fin_ts int
	if len(p_fin_ts) > 0 {
		fin_ts, err = strconv.Atoi(p_fin_ts)
		if err != nil {
			c.HTML(http.StatusBadRequest, "error.html", gin.H{
				"title": "Augur Markets: Error",
				"ErrDescr": "Can't parse fin_ts",
			})
			return
		}
	}
	if fin_ts == 0 {
		fin_ts = 2147483647
	}
	ini,fin,prices:= augur_srv.db_augur.Get_ethusd_price_history(init_ts,fin_ts)
	ts := time.Unix(int64(ini),0)
	start_date := ts.String()
	ts = time.Unix(int64(fin),0)
	end_date := ts.String()
	js_prices := build_js_ethusd_price_history(&prices)
	c.HTML(http.StatusOK, "ethusd_price.html", gin.H{
			"Prices" : prices,
			"JSPriceData" :js_prices,
			"InitTimeStamp": init_ts,
			"FinTimeSTamp": fin_ts,
			"InitDate" : start_date,
			"FinDate" : end_date,
	})
}
func whats_new_in_augur(c *gin.Context) {

	var err error
	var p_code string
	p_code = c.Param("code")
	if len(c.Query("code"))>0 {
		p_code = c.Query("code")
	}
	var code int = 0
	if len(p_code) > 0 {
		code , err = strconv.Atoi(p_code)
		if err != nil {
			c.HTML(http.StatusBadRequest, "error.html", gin.H{
				"title": "Augur Markets: Error",
				"ErrDescr": "Can't parse code",
			})
			return
		}
	}
	block_num_from,block_num_to,err := augur_srv.db_augur.Get_block_range_for_whats_new(WhatsNewAugurCode(code))
	if err != nil {
		c.HTML(http.StatusBadRequest, "error.html", gin.H{
			"title": "Augur Markets: Error",
			"ErrDescr": fmt.Sprintf("%v",err.Error()),
		})
		return
	}
	block_info,err := augur_srv.db_augur.Get_block_info(int64(block_num_from),int64(block_num_to))
	if err != nil {
		c.HTML(http.StatusBadRequest, "error.html", gin.H{
			"title": "Augur Markets: Error",
			"ErrDescr": fmt.Sprintf("%v",err.Error()),
		})
		return
	}
	c.HTML(http.StatusOK, "block_info.html", gin.H{
		"BlockInfo" : block_info,
	})
}
func user_uniswap_swaps(c *gin.Context) {

	user := c.Param("user")
	user_addr,valid := is_address_valid(c,false,user)
	if !valid {
		return
	}
	user_aid,err := augur_srv.db_augur.Nonfatal_lookup_address_id(user_addr)
	if err!=nil {
		c.HTML(http.StatusBadRequest, "error.html", gin.H{
			"title": "Augur Markets: Error",
			"ErrDescr": fmt.Sprintf("Such address wasn't found: %v",user_addr),
		})
		return
	}
	user_info,err := augur_srv.db_augur.Get_user_info(user_aid)
	swaps,total_rows := augur_srv.db_augur.Get_user_uniswap_swaps(user_aid,0,200)
	c.HTML(http.StatusOK, "user_uniswap_swaps.html", gin.H{
		"UserInfo" : user_info,
		"UserSwaps" : swaps,
		"TotalRows" : total_rows,
	})
}
func user_balancer_swaps(c *gin.Context) {

	user := c.Param("user")
	user_addr,valid := is_address_valid(c,false,user)
	if !valid {
		return
	}
	user_aid,err := augur_srv.db_augur.Nonfatal_lookup_address_id(user_addr)
	if err!=nil {
		c.HTML(http.StatusBadRequest, "error.html", gin.H{
			"title": "Augur Markets: Error",
			"ErrDescr": fmt.Sprintf("Such address wasn't found: %v",user_addr),
		})
		return
	}
	user_info,err := augur_srv.db_augur.Get_user_info(user_aid)
	swaps,total_rows := augur_srv.db_augur.Get_user_balancer_swaps(user_aid,0,200)
	c.HTML(http.StatusOK, "user_balancer_swaps.html", gin.H{
		"UserInfo" : user_info,
		"UserSwaps" : swaps,
		"TotalRows" : total_rows,
	})
}
func user_ens_names(c *gin.Context) {

	user := c.Param("user")
	user_addr,valid := is_address_valid(c,false,user)
	if !valid {
		return
	}
	user_aid,err := augur_srv.db_augur.Nonfatal_lookup_address_id(user_addr)
	if err!=nil {
		c.HTML(http.StatusBadRequest, "error.html", gin.H{
			"title": "Augur Markets: Error",
			"ErrDescr": fmt.Sprintf("Such address wasn't found: %v",user_addr),
		})
		return
	}
	user_info,err := augur_srv.db_augur.Get_user_info(user_aid)
	active_names,active_total_rows := augur_srv.db_augur.Get_user_ens_names_active(user_aid,0,1000000)
	inactive_names,inactive_total_rows := augur_srv.db_augur.Get_user_ens_names_inactive(user_aid,0,1000000)
	addr_changes,achanges_total_rows := augur_srv.db_augur.Get_user_address_change_history(user_aid,0,1000000)
	ownership_changes,own_changes_total_rows := augur_srv.db_augur.Get_user_ownership_change_history(user_aid,0,1000000)
	c.HTML(http.StatusOK, "user_ens_names.html", gin.H{
		"UserInfo" : user_info,
		"ENS_Names_Active" : active_names,
		"ENS_Names_Inactive" : inactive_names,
		"ENS_OwnershipChanges" : ownership_changes,
		"ENS_AddrChanges" : addr_changes,
		"TotalRowsActive" : active_total_rows,
		"TotalRowsInactive" :inactive_total_rows,
		"TotalRowsAddrChanges" : achanges_total_rows,
		"TotalRowsOwnershipChanges" : own_changes_total_rows,
	})
}
func show_node_text_data(c *gin.Context) {

	node := c.Param("node")
	fqdn,key_value_pairs:= augur_srv.db_augur.Get_node_text_key_values(node)
	c.HTML(http.StatusOK, "user_text_kv_pairs.html", gin.H{
		"Node" : node,
		"FullName" : fqdn,
		"KeyValuePairs" : key_value_pairs,
	})
}
func show_augur_foundry_contracts(c *gin.Context) {

	wrappers:= augur_srv.db_augur.Get_augur_foundry_wrapper_list()
	c.HTML(http.StatusOK, "augur_foundry_wrappers.html", gin.H{
		"ERC20MarketOutcomeWrappers" : wrappers,
	})
}
func show_reporting_table(c *gin.Context) {

	market := c.Param("market")
	market_addr,valid := is_address_valid(c,false,market)
	if !valid {
		return
	}
	market_aid,err := augur_srv.db_augur.Nonfatal_lookup_address_id(market_addr)
	if err!=nil {
		c.HTML(http.StatusBadRequest, "error.html", gin.H{
			"title": "Augur Markets: Error",
			"ErrDescr": fmt.Sprintf("Such address wasn't found: %v",market_addr),
		})
		return
	}
	market_info,err := augur_srv.db_augur.Get_market_info(market_addr,0,false)
	if err != nil {
		show_market_not_found_error(c,false,&market_addr)
		return
	}
	reporting_status,err := augur_srv.db_augur.Get_reporting_table(market_aid)
	if err!=nil {
		c.HTML(http.StatusBadRequest, "error.html", gin.H{
			"title": "Augur Markets: Error",
			"ErrDescr": fmt.Sprintf("Error: %v",err.Error()),
		})
		return
	}
	round_table,num_outcomes,outcomes,scalar_vals := augur_srv.db_augur.Get_round_table(market_aid)
	outcomes_split := strings.Split(outcomes,",")
	initial_report_redemption := augur_srv.db_augur.Get_initial_report_redeemed_record(market_aid)
	redeemed_participants := augur_srv.db_augur.Get_redeemed_participants(market_aid)
	losing_reports := augur_srv.db_augur.Get_losing_rep_participants(market_aid)

	c.HTML(http.StatusOK, "reporting_table.html", gin.H{
		"MarketInfo" : market_info,
		"ReportingTable" : reporting_status,
		"RoundTable" : round_table,
		"NumOutcomes" : num_outcomes,
		"Outcomes" : outcomes,
		"OutcomesSplit" : outcomes_split,
		"ScalarValues" : scalar_vals,
		"RedeemIniReporter" : initial_report_redemption,
		"WinningReports" : redeemed_participants,
		"LosingReports" : losing_reports,
	})
}
func user_rep_profit_loss(c *gin.Context) {

	user := c.Param("user")
	user_addr,valid := is_address_valid(c,false,user)
	if !valid {
		return
	}
	user_aid,err := augur_srv.db_augur.Nonfatal_lookup_address_id(user_addr)
	if err!=nil {
		c.HTML(http.StatusBadRequest, "error.html", gin.H{
			"title": "Augur Markets: Error",
			"ErrDescr": fmt.Sprintf("Such address wasn't found: %v",user_addr),
		})
		return
	}
	user_info,err := augur_srv.db_augur.Get_user_info(user_aid)
	if err!=nil {
		c.HTML(http.StatusBadRequest, "error.html", gin.H{
			"title": "Augur Markets: Error",
			"ErrDescr": fmt.Sprintf("Such address wasn't found: %v",user_addr),
		})
		return
	}
	rep_profits := augur_srv.db_augur.Get_user_report_profits(user_aid)
	c.HTML(http.StatusOK, "user_rep_pl.html", gin.H{
		"UserInfo" : user_info,
		"RepProfits" : rep_profits,
	})
}
func augur_noshow_bond_prices(c *gin.Context) {

	bond_prices := augur_srv.db_augur.Get_noshow_bond_price_history()
	js_prices := build_js_noshow_bond_price_history(&bond_prices)
	c.HTML(http.StatusOK, "noshow_bond_prices.html", gin.H{
		"Prices" : bond_prices,
		"JSPriceData" :js_prices,
	})
}
func augur_validity_bond_prices(c *gin.Context) {

	bond_prices := augur_srv.db_augur.Get_validity_bond_price_history()
	js_prices := build_js_validity_bond_price_history(&bond_prices)
	c.HTML(http.StatusOK, "validity_bond_prices.html", gin.H{
		"Prices" : bond_prices,
		"JSPriceData" :js_prices,
	})
}
func ens_name_info(c *gin.Context) {

	fqdn := c.Param("fqdn")
	ens_info,err := augur_srv.db_augur.Get_ens_record_info(fqdn)
	if err!=nil {
		c.HTML(http.StatusBadRequest, "error.html", gin.H{
			"title": "Augur Markets: Error",
			"ErrDescr": fmt.Sprintf("%v",err),
		})
		return
	}
	if len(ens_info.ContentHash) > 0 {
		data,err := hex.DecodeString(ens_info.ContentHash)
		if err==nil {
			ens_info.ContentHash,err = ens.ContenthashToString(data[:])
			Error.Printf(
				"Content hash bianry string for node %v  has invalid bin fmt : %v\n",
				ens_info.FQDN,err,
			)
		} else {
			Error.Printf(
				"Content hash bianry string couldn't be decoded for node %v : %v\n",
				ens_info.FQDN,err,
			)
		}

	}
	c.HTML(http.StatusOK, "ens_info.html", gin.H{
		"ENSInfo" : ens_info,
	})
}
func arbitrum_augur_pools(c *gin.Context) {

	if  !augur_srv.matic_initialized() {
		respond_error(c,"Database link wasn't configured")
		return 
	}
	pools := augur_srv.db_matic.Get_arbitrum_augur_pools()
	c.HTML(http.StatusOK, "arbitrum_augur_pools.html", gin.H{
		"ArbitrumAugurPools" : pools,
	})
}
func arbitrum_markets_sports(c *gin.Context) {

	if  !augur_srv.matic_initialized() {
		respond_error(c,"Database link wasn't configured")
		return 
	}
	p_status := c.Param("status")
	var status int64
	if len(p_status) > 0 {
		var success bool
		status,success = parse_int_from_remote_or_error(c,false,&p_status)
		if !success {
			return
		}
	} else {
		respond_error(c,"'status' parameter is not set")
		return
	}
	p_sort := c.Param("sort")
	var sort int64
	if len(p_sort) > 0 {
		var success bool
		sort ,success = parse_int_from_remote_or_error(c,false,&p_sort)
		if !success {
			return
		}
	} else {
		respond_error(c,"'sort' parameter is not set")
		return
	}
	contract_addrs := augur_srv.db_matic.Get_arbitrum_augur_factory_aids(&amm_contracts)
	fmt.Printf("contract_addrs = %+v\n",contract_addrs)
	total_rows,markets := augur_srv.db_matic.Get_sport_markets(status,sort,0,10000000,&amm_constants,contract_addrs)
	c.HTML(http.StatusOK, "arbitrum_markets_sports.html", gin.H{
		"Status" : status,
		"Markets" : markets,
		"TotalRows" : total_rows,
	})
}
func arbitrum_liquidity_changed(c *gin.Context) {

	if  !augur_srv.matic_initialized() {
		respond_error(c,"Database link wasn't configured")
		return 
	}
	p_market_id := c.Param("market_id")
	var market_id int64
	if len(p_market_id) > 0 {
		var success bool
		market_id,success = parse_int_from_remote_or_error(c,false,&p_market_id)
		if !success {
			return
		}
	} else {
		respond_error(c,"'market_id' parameter is not set")
		return
	}
	p_factory_aid:= c.Param("factory_aid")
	var factory_aid int64
	if len(p_factory_aid) > 0 {
		var success bool
		factory_aid,success = parse_int_from_remote_or_error(c,false,&p_factory_aid)
		if !success {
			return
		}
	} else {
		respond_error(c,"'factory_aid' parameter is not set")
		return
	}
	market,err := augur_srv.db_matic.Get_sport_market_info(&amm_constants,factory_aid,market_id)
	if err!=nil {
		respond_error(c,fmt.Sprintf("Market with market_id=%v has error: %v",market_id,err))
		return
	}
	total_rows,lchanges := augur_srv.db_matic.Get_liquidity_change_events(
		factory_aid,market_id,0,10000000,
	)
	c.HTML(http.StatusOK, "augur_amm/liquidity_changed.html", gin.H{
		"MarketId":market_id,
		"MarketInfo" : market,
		"LiquidityChanges" : lchanges,
		"TotalRows" : total_rows,
	})
}
func arbitrum_shares_swapped(c *gin.Context) {

	if  !augur_srv.matic_initialized() {
		respond_error(c,"Database link wasn't configured")
		return 
	}
	p_market_id := c.Param("market_id")
	var market_id int64
	if len(p_market_id) > 0 {
		var success bool
		market_id,success = parse_int_from_remote_or_error(c,false,&p_market_id)
		if !success {
			return
		}
	} else {
		respond_error(c,"'market_id' parameter is not set")
		return
	}
	p_contract_aid := c.Param("contract_aid")
	var contract_aid int64
	if len(p_contract_aid) > 0 {
		var success bool
		contract_aid,success = parse_int_from_remote_or_error(c,false,&p_contract_aid)
		if !success {
			return
		}
	} else {
		respond_error(c,"'contract_aid' parameter is not set")
		return
	}
	market,err := augur_srv.db_matic.Get_sport_market_info(&amm_constants,contract_aid,market_id)
	if err!=nil {
		respond_error(c,fmt.Sprintf("Market with market_id=%v has error: %v",market_id,err))
		return
	}

	total_rows,swaps:= augur_srv.db_matic.Get_shares_swapped(
		&amm_constants,contract_aid,market_id,0,10000000,
	)
	c.HTML(http.StatusOK, "augur_amm/shares_swapped.html", gin.H{
		"MarketId":market_id,
		"MarketInfo" : market,
		"Swaps" : swaps,
		"TotalRows" : total_rows,
	})
}
func amm_user_swaps(c *gin.Context) {

	if  !augur_srv.matic_initialized() {
		respond_error(c,"Database link wasn't configured")
		return
	}
	p_user := c.Param("user")
	user_addr,valid := is_address_valid(c,false,p_user)
	if !valid {
		return
	}
	success,offset,limit := parse_offset_limit_params(c)
	if !success {
		return
	}
	aid,err := augur_srv.db_matic.Nonfatal_lookup_address_id(user_addr)
	if err != nil {
		aid = 0
	}
	total_rows,swaps := augur_srv.db_matic.Get_amm_user_swaps(&amm_constants,aid,offset,limit)

	c.HTML(http.StatusOK, "augur_amm/user_swaps.html", gin.H{
		"Swaps" : swaps,
		"TotalRows" : total_rows,
		"User":p_user,
		"UserAid":aid,
	})
}
func amm_user_liquidity(c *gin.Context) {

	if  !augur_srv.matic_initialized() {
		respond_error(c,"Database link wasn't configured")
		return
	}
	p_user := c.Param("user")
	user_addr,valid := is_address_valid(c,false,p_user)
	if !valid {
		return
	}
	success,offset,limit := parse_offset_limit_params(c)
	if !success {
		return
	}
	aid,err := augur_srv.db_matic.Nonfatal_lookup_address_id(user_addr)
	if err != nil {
		aid = 0
	}
	total_rows,liquidity := augur_srv.db_matic.Get_amm_user_liquidity(&amm_constants,aid,offset,limit)

	c.HTML(http.StatusOK, "augur_amm/user_liquidity.html", gin.H{
		"Liquidity" : liquidity,
		"TotalRows" : total_rows,
		"User": p_user,
		"UserAid": aid,
	})
}
func arbitrum_market_info(c *gin.Context) {

	if  !augur_srv.matic_initialized() {
		respond_error(c,"Database link wasn't configured")
		return 
	}
	p_market_id := c.Param("market_id")
	var market_id int64
	if len(p_market_id) > 0 {
		var success bool
		market_id,success = parse_int_from_remote_or_error(c,false,&p_market_id)
		if !success {
			return
		}
	} else {
		respond_error(c,"'market_id' parameter is not set")
		return
	}
	p_contract_aid := c.Param("contract_aid")
	var contract_aid int64
	if len(p_contract_aid) > 0 {
		var success bool
		contract_aid,success = parse_int_from_remote_or_error(c,false,&p_contract_aid)
		if !success {
			return
		}
	} else {
		respond_error(c,"'contract_aid' parameter is not set")
		return
	}
	market,err := augur_srv.db_matic.Get_sport_market_info(&amm_constants,contract_aid,market_id)
	if err!=nil {
		respond_error(c,fmt.Sprintf("Market with market_id=%v has error: %v",market_id,err))
		return
	}

	c.HTML(http.StatusOK, "augur_amm/market_info.html", gin.H{
		"MarketId":market_id,
		"MarketInfo" : market,
	})
}
func arbitrum_market_liquidity_providers(c *gin.Context) {

	if  !augur_srv.matic_initialized() {
		respond_error(c,"Database link wasn't configured")
		return 
	}
	p_market_id := c.Param("market_id")
	var market_id int64
	if len(p_market_id) > 0 {
		var success bool
		market_id,success = parse_int_from_remote_or_error(c,false,&p_market_id)
		if !success {
			return
		}
	} else {
		respond_error(c,"'market_id' parameter is not set")
		return
	}
	p_contract_aid := c.Param("contract_aid")
	var contract_aid int64
	if len(p_contract_aid) > 0 {
		var success bool
		contract_aid,success = parse_int_from_remote_or_error(c,false,&p_contract_aid)
		if !success {
			return
		}
	} else {
		respond_error(c,"'contract_aid' parameter is not set")
		return
	}
	market,err := augur_srv.db_matic.Get_sport_market_info(&amm_constants,contract_aid,market_id)
	if err!=nil {
		respond_error(c,fmt.Sprintf("Market with market_id=%v has error: %v",market_id,err))
		return
	}
	pool_aid,err := augur_srv.db_matic.Get_market_pool_aid(contract_aid,market_id)
	if err!=nil {
		respond_error(c,fmt.Sprintf("Pool wasn't found in the database for this market: %v",err))
		return
	}
	providers:= augur_srv.db_matic.Get_pool_holder_distribution(pool_aid)
	if err!=nil {
		respond_error(c,fmt.Sprintf("Market with market_id=%v has error: %v",market_id,err))
		return
	}
	js_tok_distr := build_js_token_holder_distribution(&providers)

	c.HTML(http.StatusOK, "augur_amm/liquidity_providers_distrib.html", gin.H{
		"MarketId":market_id,
		"MarketInfo" : market,
		"PoolTokenHolderDistribution" : providers,
		"JSTokenHolderDistribution" : js_tok_distr,
	})
}
func arbitrum_market_outside_augur_shares_burned(c *gin.Context) {

	if  !augur_srv.matic_initialized() {
		respond_error(c,"Database link wasn't configured")
		return 
	}
	p_market_id := c.Param("market_id")
	var market_id int64
	if len(p_market_id) > 0 {
		var success bool
		market_id,success = parse_int_from_remote_or_error(c,false,&p_market_id)
		if !success {
			return
		}
	} else {
		respond_error(c,"'market_id' parameter is not set")
		return
	}
	p_contract_aid := c.Param("contract_aid")
	var contract_aid int64
	if len(p_contract_aid) > 0 {
		var success bool
		contract_aid,success = parse_int_from_remote_or_error(c,false,&p_contract_aid)
		if !success {
			return
		}
	} else {
		respond_error(c,"'contract_aid' parameter is not set")
		return
	}
	market,err := augur_srv.db_matic.Get_sport_market_info(&amm_constants,contract_aid,market_id)
	if err!=nil {
		respond_error(c,fmt.Sprintf("Market with market_id=%v has error: %v",market_id,err))
		return
	}
	offset := int(0) ; limit:= int(100000)
	operations := augur_srv.db_matic.Get_outside_augur_shares_burned(contract_aid,market_id,offset,limit)

	c.HTML(http.StatusOK, "augur_amm/outside_augur_shares_bruned.html", gin.H{
		"MarketId":market_id,
		"MarketInfo" : market,
		"SharesBurnedOperations" : operations,
	})

}
func arbitrum_market_outside_augur_shares_minted(c *gin.Context) {

	if  !augur_srv.matic_initialized() {
		respond_error(c,"Database link wasn't configured")
		return 
	}
	p_market_id := c.Param("market_id")
	var market_id int64
	if len(p_market_id) > 0 {
		var success bool
		market_id,success = parse_int_from_remote_or_error(c,false,&p_market_id)
		if !success {
			return
		}
	} else {
		respond_error(c,"'market_id' parameter is not set")
		return
	}
	p_contract_aid := c.Param("contract_aid")
	var contract_aid int64
	if len(p_contract_aid) > 0 {
		var success bool
		contract_aid,success = parse_int_from_remote_or_error(c,false,&p_contract_aid)
		if !success {
			return
		}
	} else {
		respond_error(c,"'contract_aid' parameter is not set")
		return
	}
	market,err := augur_srv.db_matic.Get_sport_market_info(&amm_constants,contract_aid,market_id)
	if err!=nil {
		respond_error(c,fmt.Sprintf("Market with market_id=%v has error: %v",market_id,err))
		return
	}
	offset := int(0) ; limit:= int(100000)
	operations := augur_srv.db_matic.Get_outside_augur_shares_minted(contract_aid,market_id,offset,limit)

	c.HTML(http.StatusOK, "augur_amm/outside_augur_shares_minted.html", gin.H{
		"MarketId":market_id,
		"MarketInfo" : market,
		"SharesMintedOperations" : operations,
	})

}
func arbitrum_market_outside_augur_balancer_swaps(c *gin.Context) {

	if  !augur_srv.matic_initialized() {
		respond_error(c,"Database link wasn't configured")
		return 
	}
	p_market_id := c.Param("market_id")
	var market_id int64
	if len(p_market_id) > 0 {
		var success bool
		market_id,success = parse_int_from_remote_or_error(c,false,&p_market_id)
		if !success {
			return
		}
	} else {
		respond_error(c,"'market_id' parameter is not set")
		return
	}
	p_contract_aid := c.Param("contract_aid")
	var contract_aid int64
	if len(p_contract_aid) > 0 {
		var success bool
		contract_aid,success = parse_int_from_remote_or_error(c,false,&p_contract_aid)
		if !success {
			return
		}
	} else {
		respond_error(c,"'contract_aid' parameter is not set")
		return
	}
	market,err := augur_srv.db_matic.Get_sport_market_info(&amm_constants,contract_aid,market_id)
	if err!=nil {
		respond_error(c,fmt.Sprintf("Market with market_id=%v has error: %v",market_id,err))
		return
	}
	pool_aid,err := augur_srv.db_matic.Get_market_pool_aid(contract_aid,market_id)
	if err!=nil {
		respond_error(c,fmt.Sprintf("Pool wasn't found in the database for this market: %v",err))
		return
	}
	pool_addr,_ := augur_srv.db_matic.Lookup_address(pool_aid)
	offset:=int(0);limit:=int(1000000000)
	balancer_swaps := augur_srv.db_matic.Get_outside_augur_balancer_swaps(pool_aid,offset,limit)

	c.HTML(http.StatusOK, "augur_amm/balancer_swaps_outside_augur.html", gin.H{
		"MarketId":market_id,
		"MarketInfo" : market,
		"PoolAid": pool_aid,
		"PoolAddr" : pool_addr,
		"BalancerSwaps" : balancer_swaps,
	})

}
func arbitrum_market_outside_augur_erc20_transfers(c *gin.Context) {

	if  !augur_srv.matic_initialized() {
		respond_error(c,"Database link wasn't configured")
		return 
	}
	p_market_id := c.Param("market_id")
	var market_id int64
	if len(p_market_id) > 0 {
		var success bool
		market_id,success = parse_int_from_remote_or_error(c,false,&p_market_id)
		if !success {
			return
		}
	} else {
		respond_error(c,"'market_id' parameter is not set")
		return
	}
	p_contract_aid := c.Param("contract_aid")
	var contract_aid int64
	if len(p_contract_aid) > 0 {
		var success bool
		contract_aid,success = parse_int_from_remote_or_error(c,false,&p_contract_aid)
		if !success {
			return
		}
	} else {
		respond_error(c,"'contract_aid' parameter is not set")
		return
	}
	market,err := augur_srv.db_matic.Get_sport_market_info(&amm_constants,contract_aid,market_id)
	if err!=nil {
		respond_error(c,fmt.Sprintf("Market with market_id=%v couldn't be located, error: %v",market_id,err))
		return
	}
	offset:=int(0);limit:=int(1000000000)
	transfers := augur_srv.db_matic.Get_erc20_transfers_outside_augur(contract_aid,market_id,offset,limit)

	c.HTML(http.StatusOK, "augur_amm/erc20_transfers_outside_augur.html", gin.H{
		"MarketId":market_id,
		"MarketInfo" : market,
		"ERC20Transfers" : transfers,
	})
}
func poly_buysell_operations(c *gin.Context) {

	p_market_id := c.Param("market_id")
	var market_id int64
	if len(p_market_id) > 0 {
		var success bool
		market_id,success = parse_int_from_remote_or_error(c,false,&p_market_id)
		if !success {
			return
		}
	} else {
		respond_error(c,"'market_id' parameter is not set")
		return
	}

	fpmm_aid := augur_srv.db_matic.Get_fpmm_contract_aid(market_id)
	if fpmm_aid == 0 {
		respond_error(c,"Polymarket with this ID wasn't found")
		return
	}

	market_info,err := augur_srv.db_matic.Get_poly_market_info(market_id)
	if err != nil {
		respond_error(c,"Market not found")
		return
	}
	operations := augur_srv.db_matic.Get_polymarkets_buysell_operations(fpmm_aid,0,1000000)

	var js_outcomes_history JSOutcomes
	for outc:=0; outc<int(market_info.OutcomeSlotCount); outc++ {
		prices:= augur_srv.db_matic.Get_poly_market_outcome_price_history(fpmm_aid,int32(outc))
		js_prices := build_js_polymarkets_outcome_price_history(&prices)
		js_outcomes_history.OutcomesDataJS  = append(js_outcomes_history.OutcomesDataJS,js_prices)
	}
	prices:= augur_srv.db_matic.Get_poly_market_outcome_price_history(fpmm_aid,0)
	price0 := build_js_polymarkets_outcome_price_history(&prices)
	prices= augur_srv.db_matic.Get_poly_market_outcome_price_history(fpmm_aid,1)
	price1 := build_js_polymarkets_outcome_price_history(&prices)

	c.HTML(http.StatusOK, "polymarket/buysell_operations.html", gin.H{
		"BuySellOperations" : operations,
		"MarketId" : market_id,
		"ContractAid" : fpmm_aid,
		"Prices" : js_outcomes_history,
		"Price0" : price0,
		"Price1" : price1,
	})
}
func poly_liquidity_operations(c *gin.Context) {

	p_market_id := c.Param("market_id")
	var market_id int64
	if len(p_market_id) > 0 {
		var success bool
		market_id,success = parse_int_from_remote_or_error(c,false,&p_market_id)
		if !success {
			return
		}
	} else {
		respond_error(c,"'market_id' parameter is not set")
		return
	}

	fpmm_aid := augur_srv.db_matic.Get_fpmm_contract_aid(market_id)
	if fpmm_aid == 0 {
		respond_error(c,"Polymarket with this ID wasn't found")
		return
	}

	operations := augur_srv.db_matic.Get_polymarkets_liquidity_operations(fpmm_aid,0,1000000)

	c.HTML(http.StatusOK, "polymarket/liquidity_operations.html", gin.H{
		"LiquidityOperations" : operations,
		"MarketId" : market_id,
		"ContractAid" : fpmm_aid,
	})
}
func poly_market_info(c *gin.Context) {

	p_market_id := c.Param("market_id")
	var market_id int64
	if len(p_market_id) > 0 {
		var success bool
		market_id,success = parse_int_from_remote_or_error(c,false,&p_market_id)
		if !success {
			return
		}
	} else {
		respond_error(c,"'market_id' parameter is not set")
		return
	}
	info,err := augur_srv.db_matic.Get_poly_market_info(market_id)
	if err != nil {
		respond_error(c,"Market not found")
		return
	}
	c.HTML(http.StatusOK, "polymarket/market_info.html", gin.H{
		"MarketInfo" : info,
		"MarketId" : market_id,
	})
}
func poly_market_stats(c *gin.Context) {

	p_market_id := c.Param("market_id")
	var market_id int64
	if len(p_market_id) > 0 {
		var success bool
		market_id,success = parse_int_from_remote_or_error(c,false,&p_market_id)
		if !success {
			return
		}
	} else {
		respond_error(c,"'market_id' parameter is not set")
		return
	}
	fpmm_aid := augur_srv.db_matic.Get_fpmm_contract_aid(market_id)
	if fpmm_aid == 0 {
		respond_error(c,"Polymarket with this ID wasn't found")
		return
	}
	stats,_:= augur_srv.db_matic.Get_poly_market_stats(fpmm_aid)
	c.HTML(http.StatusOK, "polymarket/market_stats.html", gin.H{
		"MarketStats" : stats,
		"MarketId" : market_id,
		"ContractAid" : fpmm_aid,
	})
}
func poly_liq_hist_global(c *gin.Context) {

	success,init_ts,fin_ts,interval_secs := parse_timeframe_params(c)
	if !success {
		return
	}

	liq_hist := augur_srv.db_matic.Get_polymarket_global_liquidity_history(init_ts,fin_ts,interval_secs)

	c.HTML(http.StatusOK, "polymarket/global_liquidity.html", gin.H{
		"GlobalLiquidityHistory" : liq_hist,
		"InitTs" : init_ts,
		"FinTs" : fin_ts,
		"Interval" : interval_secs,
	})
}
func poly_market_liquidity_periods(c *gin.Context) {

	success,init_ts,fin_ts,interval_secs := parse_timeframe_params(c)
	if !success {
		return
	}

	p_market_id := c.Param("market_id")
	var market_id int64
	if len(p_market_id) > 0 {
		var success bool
		market_id,success = parse_int_from_remote_or_error(c,false,&p_market_id)
		if !success {
			return
		}
	} else {
		respond_error(c,"'market_id' parameter is not set")
		return
	}

	fpmm_aid := augur_srv.db_matic.Get_fpmm_contract_aid(market_id)
	if fpmm_aid == 0 {
		respond_error(c,"Polymarket with this ID wasn't found")
		return
	}

	liq_hist := augur_srv.db_matic.Get_polymarket_market_liquidity_history(fpmm_aid,init_ts,fin_ts,interval_secs)

	c.HTML(http.StatusOK, "polymarket/market_liquidity_by_periods.html", gin.H{
		"MarketId" : market_id,
		"ContractAid" : fpmm_aid,
		"MarketLiquidityHistory" : liq_hist,
		"InitTs" : init_ts,
		"FinTs" : fin_ts,
		"Interval" : interval_secs,
	})
}
func poly_user_list(c *gin.Context) {

	p_market_id := c.Param("market_id")
	var market_id int64
	if len(p_market_id) > 0 {
		var success bool
		market_id,success = parse_int_from_remote_or_error(c,false,&p_market_id)
		if !success {
			return
		}
	} else {
		respond_error(c,"'market_id' parameter is not set")
		return
	}

	fpmm_aid := augur_srv.db_matic.Get_fpmm_contract_aid(market_id)
	if fpmm_aid == 0 {
		respond_error(c,"Polymarket with this ID wasn't found")
		return
	}

	users_list := augur_srv.db_matic.Get_polymarkets_market_user_list(fpmm_aid)

	c.HTML(http.StatusOK, "polymarket/market_userlist.html", gin.H{
		"MarketId" : market_id,
		"ContractAid" : fpmm_aid,
		"Users" : users_list,
	})

}
func poly_market_trader_operations(c *gin.Context) {

	p_market_id := c.Param("market_id")
	var market_id int64
	if len(p_market_id) > 0 {
		var success bool
		market_id,success = parse_int_from_remote_or_error(c,false,&p_market_id)
		if !success {
			return
		}
	} else {
		respond_error(c,"'market_id' parameter is not set")
		return
	}
	p_user_aid := c.Param("user_aid")
	var user_aid int64
	if len(p_user_aid) > 0 {
		var success bool
		user_aid,success = parse_int_from_remote_or_error(c,false,&p_user_aid)
		if !success {
			return
		}
	} else {
		respond_error(c,"'user_aid' parameter is not set")
		return
	}
	success,offset,limit := parse_offset_limit_params(c)
	if !success {
		return
	}

	fpmm_aid := augur_srv.db_matic.Get_fpmm_contract_aid(market_id)
	if fpmm_aid == 0 {
		respond_error(c,"Polymarket with this ID wasn't found")
		return
	}

	trade_list := augur_srv.db_matic.Get_poly_market_trader_operations(fpmm_aid,user_aid,offset,limit)

	c.HTML(http.StatusOK, "polymarket/market_trader_operations.html", gin.H{
		"MarketId" : market_id,
		"UserAid" : user_aid,
		"ContractAid" : fpmm_aid,
		"TraderOperations" : trade_list,
	})
}
func poly_market_funder_operations(c *gin.Context) {

	p_market_id := c.Param("market_id")
	var market_id int64
	if len(p_market_id) > 0 {
		var success bool
		market_id,success = parse_int_from_remote_or_error(c,false,&p_market_id)
		if !success {
			return
		}
	} else {
		respond_error(c,"'market_id' parameter is not set")
		return
	}
	p_user_aid := c.Param("user_aid")
	var user_aid int64
	if len(p_user_aid) > 0 {
		var success bool
		user_aid,success = parse_int_from_remote_or_error(c,false,&p_user_aid)
		if !success {
			return
		}
	} else {
		respond_error(c,"'user_aid' parameter is not set")
		return
	}
	success,offset,limit := parse_offset_limit_params(c)
	if !success {
		return
	}

	fpmm_aid := augur_srv.db_matic.Get_fpmm_contract_aid(market_id)
	if fpmm_aid == 0 {
		respond_error(c,"Polymarket with this ID wasn't found")
		return
	}

	liq_operation_list := augur_srv.db_matic.Get_poly_market_funder_operations(fpmm_aid,user_aid,offset,limit)

	c.HTML(http.StatusOK, "polymarket/market_funder_operations.html", gin.H{
		"MarketId" : market_id,
		"UserAid" : user_aid,
		"ContractAid" : fpmm_aid,
		"FunderOperations" : liq_operation_list,
	})
}
func poly_market_open_positions(c *gin.Context) {

	p_market_id := c.Param("market_id")
	var market_id int64
	if len(p_market_id) > 0 {
		var success bool
		market_id,success = parse_int_from_remote_or_error(c,false,&p_market_id)
		if !success {
			return
		}
	} else {
		respond_error(c,"'market_id' parameter is not set")
		return
	}

	fpmm_aid := augur_srv.db_matic.Get_fpmm_contract_aid(market_id)
	if fpmm_aid == 0 {
		respond_error(c,"Polymarket with this ID wasn't found")
		return
	}

	open_positions,prices := augur_srv.db_matic.Get_poly_market_open_positions(fpmm_aid)

	c.HTML(http.StatusOK, "polymarket/market_open_positions.html", gin.H{
		"MarketId" : market_id,
		"ContractAid" : fpmm_aid,
		"OpenPositions" : open_positions,
		"Prices" : prices,
	})
}
func poly_market_user_open_positions(c *gin.Context) {

	p_user_aid := c.Param("user_aid")
	var user_aid int64
	if len(p_user_aid) > 0 {
		var success bool
		user_aid,success = parse_int_from_remote_or_error(c,false,&p_user_aid)
		if !success {
			return
		}
	} else {
		respond_error(c,"'user_aid' parameter is not set")
		return
	}

	user_open_positions := augur_srv.db_matic.Get_poly_market_user_open_positions(user_aid)

	c.HTML(http.StatusOK, "polymarket/market_user_open_positions.html", gin.H{
		"UserAid": user_aid,
		"UserOpenPositions" :user_open_positions,
	})
}
func poly_market_funder_share_ratio(c *gin.Context) {

	p_market_id := c.Param("market_id")
	var market_id int64
	if len(p_market_id) > 0 {
		var success bool
		market_id,success = parse_int_from_remote_or_error(c,false,&p_market_id)
		if !success {
			return
		}
	} else {
		respond_error(c,"'market_id' parameter is not set")
		return
	}

	fpmm_aid := augur_srv.db_matic.Get_fpmm_contract_aid(market_id)
	if fpmm_aid == 0 {
		respond_error(c,"Polymarket with this ID wasn't found")
		return
	}

	share_ratios := augur_srv.db_matic.Get_poly_liquidity_provider_share_ratio(fpmm_aid)

	c.HTML(http.StatusOK, "polymarket/market_funders_share_ratio.html", gin.H{
		"MarketId" : market_id,
		"ContractAid" : fpmm_aid,
		"ShareRatios" : share_ratios,
	})
}
func poly_markets_listing(c *gin.Context) {

	p_status := c.Param("status")
	var status int64
	if len(p_status) > 0 {
		var success bool
		status,success = parse_int_from_remote_or_error(c,HTTP,&p_status)
		if !success {
			return
		}
	} else {
		// the default is status = 0
	}
	p_sort:= c.Param("sort")
	var sort int64
	if len(p_sort) > 0 {
		var success bool
		sort,success = parse_int_from_remote_or_error(c,HTTP,&p_sort)
		if !success {
			return
		}
	} else {
		// the default is sort = 0
	}

	category := c.Query("c")

	markets_listing := augur_srv.db_matic.Get_polymarkets_markets(int(status),int(sort),category)
	num_elts := len(markets_listing)
	c.HTML(http.StatusOK, "polymarket/market_listing.html", gin.H{
		"Markets" : markets_listing,
		"QueryingStatus" : status,
		"NumElts" : num_elts,
	})
}
func poly_top_users(c *gin.Context) {

	top_profit_makers := augur_srv.db_matic.Get_polymarket_top_profit_makers()
	top_trade_makers := augur_srv.db_matic.Get_polymarket_top_trade_makers()
	top_volume_makers := augur_srv.db_matic.Get_polymarket_top_volume_makers()
	c.HTML(http.StatusOK, "poly_top_users.html", gin.H{
			"title": "Top 100 Users of Polymarket Markets",
			"ProfitMakers" : top_profit_makers,
			"TradeMakers" : top_trade_makers,
			"VolumeMakers" : top_volume_makers,
	})
}
func poly_market_payout_redemptions(c *gin.Context) {

	p_market_id := c.Param("market_id")
	var market_id int64
	if len(p_market_id) > 0 {
		var success bool
		market_id,success = parse_int_from_remote_or_error(c,false,&p_market_id)
		if !success {
			return
		}
	} else {
		respond_error(c,"'market_id' parameter is not set")
		return
	}

	condition_id := augur_srv.db_matic.Get_condition_id(market_id)
	if len(condition_id) == 0 {
		respond_error(c,"Polymarket with this ID wasn't found")
		return
	}

	payout_redemptions := augur_srv.db_matic.Get_polymarket_market_redemptions(condition_id,0,1000000)

	c.HTML(http.StatusOK, "polymarket/market_redemptions.html", gin.H{
		"MarketId" : market_id,
		"PayoutRedemptions" : payout_redemptions,
	})
}
func poly_market_categories(c *gin.Context) {

	categories := augur_srv.db_matic.Get_polymarket_categories()
	c.HTML(http.StatusOK, "polymarket/categories.html", gin.H{
		"MarketCategories" : categories,
	})
}
func poly_market_erc1155_transfers(c *gin.Context) {

	p_market_id := c.Param("market_id")
	var market_id int64
	if len(p_market_id) > 0 {
		var success bool
		market_id,success = parse_int_from_remote_or_error(c,HTTP,&p_market_id)
		if !success {
			return
		}
	} else {
		respond_error(c,"'market_id' parameter is not set")
		return
	}
	success,offset,limit := parse_offset_limit_params(c)
	if !success {
		return
	}

	fpmm_aid := augur_srv.db_matic.Get_fpmm_contract_aid(market_id)
	if fpmm_aid == 0 {
		respond_error(c,"Polymarket with this ID wasn't found")
		return
	}

	erc1155_transfers := augur_srv.db_matic.Get_polymarket_erc1155_transfers(fpmm_aid,offset,limit)

	c.HTML(http.StatusOK, "polymarket/erc1155_transfers.html", gin.H{
		"MarketId" : market_id,
		"ContractAid" : fpmm_aid,
		"ERC1155Transfers" : erc1155_transfers,
	})
}
func poly_market_open_interest_history(c *gin.Context) {

	p_market_id := c.Param("market_id")
	var market_id int64
	if len(p_market_id) > 0 {
		var success bool
		market_id,success = parse_int_from_remote_or_error(c,HTTP,&p_market_id)
		if !success {
			return
		}
	} else {
		respond_error(c,"'market_id' parameter is not set")
		return
	}
	success,offset,limit := parse_offset_limit_params(c)
	if !success {
		return
	}

	fpmm_aid := augur_srv.db_matic.Get_fpmm_contract_aid(market_id)
	if fpmm_aid == 0 {
		respond_error(c,"Polymarket with this ID wasn't found")
		return
	}
	condition_id := augur_srv.db_matic.Get_condition_id(market_id)
	caddrs := augur_srv.db_matic.Get_polymarket_contract_addresses()
	oi_hist := augur_srv.db_matic.Get_polymarket_open_interst_history(
		caddrs.USDCAid,
		caddrs.CondTokAid,
		fpmm_aid,
		condition_id,
		offset,
		limit,
	)

	c.HTML(http.StatusOK, "polymarket/open_interest_history.html", gin.H{
		"MarketId" : market_id,
		"ContractAid" : fpmm_aid,
		"OIHistory" : oi_hist,
		"CondTokAid" : caddrs.CondTokAid,
		"USDCAid" : caddrs.USDCAid,
	})
}
func poly_market_search(c *gin.Context) {

	p_keyword:= c.Query("q")
	if len(p_keyword) == 0 {
		respond_error(c,"'q' parameter is not set")
		return
	}

	results := augur_srv.db_matic.Search_polymarket_keywords(p_keyword)

	c.HTML(http.StatusOK, "polymarket/search_results.html", gin.H{
		"Keywords" : p_keyword,
		"SearchResults" : results,
	})
}
