package main
import (
	"os"
	"fmt"
	"log"
	"time"
	"strconv"
	"net/http"
	"encoding/hex"
	"github.com/gin-gonic/gin"
	"math/big"
	"context"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"

	. "github.com/PredictionExplorer/augur-explorer/primitives"
	. "github.com/PredictionExplorer/augur-explorer/dbs"
	. "github.com/PredictionExplorer/augur-explorer/contracts"
)
const (
	DEFAULT_DB_LOG_FILE_NAME = "/var/tmp/backend-db.log"
	DEFAULT_AMM_LOG_FILE_NAME = ""
	DEFAULT_MARKET_ROWS_LIMIT int	= 500
	DEFAILT_MARKET_TRADES_LIMIT int = 20
	DEFAULT_USER_REPORTS_LIMIT int = 30
	DEFAULT_MARKET_REPORTS_LIMIT int = 40

	JSON				bool = true
	HTTP				bool = false
)
type AugurServer struct {
	db_augur		*SQLStorage
	db_matic		*SQLStorage
	db_arbitrum		*SQLStorage
	db_main_net		*SQLStorage
}
func (self *AugurServer) matic_initialized() bool {

	if self.db_matic == nil {
		return false
	}
	return true
}
func (self *AugurServer) arbitrum_initialized() bool {

	if self.db_arbitrum == nil {
		return false
	}
	return true
}
func (self *AugurServer) main_net_initialized() bool {

	if self.db_main_net == nil {
		return false
	}
	return true
}
func connect_to_main_net(srv *AugurServer) {
	eth_user := os.Getenv("ETH_USERNAME")
	eth_passwd := os.Getenv("ETH_PASSWORD")
	eth_db_name := os.Getenv("ETH_DATABASE")
	eth_host_port := os.Getenv("ETH_HOST")
	if len(eth_user) > 0 {
		log_dir:=fmt.Sprintf("%v/%v",os.Getenv("HOME"),DEFAULT_LOG_DIR)
		db_log_file:=fmt.Sprintf("%v/%v",log_dir,"mainnet-db.log")
		eth_db_logfile, err := os.OpenFile(db_log_file, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
		if err != nil {
			fmt.Printf("Can't start: %v\n",err)
			os.Exit(1)
		}
		ETH_DB := log.New(eth_db_logfile,"INFO: ",log.Ldate|log.Ltime|log.Lshortfile)

		srv.db_main_net = New_sql_storage(
			Info,
			ETH_DB,
			eth_host_port,
			eth_db_name,
			eth_user,
			eth_passwd,
		)
	}
}
func connect_to_amm(srv *AugurServer) {
	amm_user := os.Getenv("AMM_USERNAME")
	amm_passwd := os.Getenv("AMM_PASSWORD")
	amm_db_name := os.Getenv("AMM_DATABASE")
	amm_host_port := os.Getenv("AMM_HOST")
	if len(amm_user) > 0 {
		fmt.Printf("Amm=%v\n",amm_user)
		log_dir:=fmt.Sprintf("%v/%v",os.Getenv("HOME"),DEFAULT_LOG_DIR)
		db_log_file:=fmt.Sprintf("%v/%v",log_dir,"amm-db.log")
		amm_db_logfile, err := os.OpenFile(db_log_file, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
		if err != nil {
			fmt.Printf("Can't start: %v\n",err)
			os.Exit(1)
		}
		AMM_DB := log.New(amm_db_logfile,"INFO: ",log.Ldate|log.Ltime|log.Lshortfile)

		srv.db_matic = New_sql_storage(
			Info,
			AMM_DB,
			amm_host_port,
			amm_db_name,
			amm_user,
			amm_passwd,
		)
	}
}
func connect_to_arbitrum(srv *AugurServer ) {

	arb_user := os.Getenv("ARB_USERNAME")
	arb_passwd := os.Getenv("ARB_PASSWORD")
	arb_db_name := os.Getenv("ARB_DATABASE")
	arb_host_port := os.Getenv("ARB_HOST")
	if len(arb_user) > 0 {
		log_dir:=fmt.Sprintf("%v/%v",os.Getenv("HOME"),DEFAULT_LOG_DIR)
		db_log_file:=fmt.Sprintf("%v/%v",log_dir,"arbitrum-db.log")
		arbitrum_db_logfile, err := os.OpenFile(db_log_file, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
		if err != nil {
			fmt.Printf("Can't start: %v\n",err)
			os.Exit(1)
		}
		arb_DB := log.New(arbitrum_db_logfile,"INFO: ",log.Ldate|log.Ltime|log.Lshortfile)

		srv.db_arbitrum= New_sql_storage(
			Info,
			arb_DB,
			arb_host_port,
			arb_db_name,
			arb_user,
			arb_passwd,
		)
	}
}
func create_augur_server() *AugurServer {

	log_dir:=fmt.Sprintf("%v/%v",os.Getenv("HOME"),DEFAULT_LOG_DIR)
	os.MkdirAll(log_dir, os.ModePerm)
	web_db_log_file:=fmt.Sprintf("%v/%v",log_dir,"webserver-db.log")

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
	srv.db_augur= Connect_to_storage(Info)
	srv.db_augur.Init_log(web_db_log_file)
	connect_to_amm(srv)
	connect_to_arbitrum(srv)
	connect_to_main_net(srv)

	return srv
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
func get_REP_token_price_in_ETH() (float64,error) {

	// token0 - REP (0x221657776846890989a759BA2973e427DfF5C9bB)
	// token1 - Wrapped ETH (0xC02aaA39b223FE8D0A0e5C4F27eAD9083C756Cc2)
	addr := common.HexToAddress(REP_ETH_UNISWAP_PAIR_ADDR)
	ctrct_pair,err := NewUniswapV2Pair(addr,eclient)
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
	big_eth_balance,err := eclient.BalanceAt(ctx,*addr,nil)
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
func is_bidding_war_enabled() {
	bw_enabled := os.Getenv("BIDDING_WAR_ENABLED")
	if len(bw_enabled) > 0 {
		cosmic_game_init()
	}
}
