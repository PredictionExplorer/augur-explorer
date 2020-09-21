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

	// thes variables should be removed on the next code reorg task
	market_order_id int64 = 0
	fill_order_id int64 = 0

	Error   *log.Logger
	Info    *log.Logger
)
func initialize() {

	caddrs_obj,err := augur_srv.storage.Get_contract_addresses()
	if err!=nil {
		Fatalf("Can't find contract addresses in 'contract_addresses' table")
	}
	caddrs=&caddrs_obj

}
func secure_https(r http.Handler) {
	autotls.Run(r, "localhost")
}
func main() {

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

	augur_srv = create_augur_server(&market_order_id,db_log_file,Info)

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
	port_secure := os.Getenv("AUGUR_HTTPS_PORT")

	if port_plain == "" {
		port_plain = "9090"
		log.Printf("Defaulting plain HTTP to port %s", port_plain)
	}
	if port_secure== "" {
		port_secure= "9443"
		log.Printf("Defaulting secure protocol to port %s", port_secure)
	}


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
	r.GET("/black//deptha/:market_aid/:outcome", market_depth_ajax)
	r.GET("/black/mphist/:market/:outcome", market_price_history)
	r.GET("/black/search", search)
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
	r.GET("/black/statement/:addr",account_statement)
	r.GET("/black/oohist/:addr",open_order_history)

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
	r.GET("/api/user_markets/:user/:active",  a1_user_markets)
	r.GET("/api/utrades/:user/:market",  a1_user_trades_for_market)
	r.GET("/api/user_pl/:user",  a1_user_profit_loss)
	r.GET("/api/user_opos/:user",  a1_user_open_positions)
	r.GET("/api/user_reports/:user",  a1_user_reports)
	r.GET("/api/user_oorders/:user",  a1_user_open_orders)


	m := autocert.Manager{
		Prompt:     autocert.AcceptTOS,
		HostPolicy: autocert.HostWhitelist("predictionexplorer.com","www.predictionexplorer.com"),
		Cache:      autocert.DirCache(os.Getenv("HOME")+".tls-autocert-cache"),
	}
	_ = m
	// Listen and serve on defined port
	log.Printf("Listening on port %s", port_plain)

	go func() {
		log.Printf("%v",autotls.Run(r, "predictionexplorer.com"))
	}()
	r.Run(":" + port_plain)

}
