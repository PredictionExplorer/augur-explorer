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
	r.LoadHTMLGlob("html/templates/*html")
	r.Use(gin.Logger())

	// Main HTML
	r.GET("/", main_page)
	r.GET("/index.html", main_page)
	r.GET("/index.htm", main_page)
	r.GET("/INDEX.HTM", main_page)

	// All the other dynamic HTML pages
	r.GET("/markets.html",markets)
	r.GET("/statistics.html",statistics)
	r.GET("/categories.html",categories)
	r.GET("/explorer.html",explorer)
	r.GET("/market/:market",  market_info)
	r.GET("/fulltradelist/:market",  full_trade_list)
	r.GET("/mdepth/:market/:outcome", market_depth)
	r.GET("/deptha/:market_aid/:outcome", market_depth_ajax)
	r.GET("/mphist/:market/:outcome", market_price_history)
	r.GET("/search", search)
	r.GET("/money/:addr",  read_money)
	r.GET("/order/:order",  order)
	r.GET("/category/:catid",  category)
	r.GET("/user/:addr",  user_info)
	r.GET("/fullreports/:addr",  full_reports)
	r.GET("/umarkets/:addr",  user_markets)
	r.GET("/udw/:addr",  user_deposits_withdrawals)
	r.GET("/block/:block_num",  block_info)
	r.GET("/topusers.html",top_users)
	r.GET("/mdstat/:market_aid/:outcome_idx/:last_oo_id",market_depth_status)
	r.GET("/umtrades.html",user_trades_for_market)
	r.GET("/statement/:addr",account_statement)

	// API calls for the new FrontEnd
	r.GET("/api/active_market_ids",a1_active_market_ids)
	r.GET("/api/active_markets/:start/:num_rows",a1_active_markets)
	r.GET("/api/mkt_card/:market_aid",a1_market_card)
	r.GET("/api/mkt_cards/:market_aid_list",a1_multiple_market_cards)
	r.GET("/api/market/:market",  a1_market_info)
	r.GET("/api/user/:user",  a1_user_info)


	r.Static("/imgs", "./html/imgs")
	r.Static("/res", "./html/res")			// resources (static)
	r.StaticFile("/favicon.ico", "./html/res/favicon.ico")

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
