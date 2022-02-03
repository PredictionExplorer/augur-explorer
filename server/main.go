package main

import (
	"log"
	"os"
	"fmt"
	"net/http"
	"crypto/tls"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/autotls"
	"github.com/ethereum/go-ethereum/ethclient"

	"golang.org/x/crypto/acme/autocert"

	. "github.com/PredictionExplorer/augur-explorer/primitives"
	. "github.com/PredictionExplorer/augur-explorer/contracts"
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
		Info.Printf("Warning! Can't find contract addresses in 'contract_addresses' table, using null values: error text: %v",err)
	}
	caddrs=&caddrs_obj

	amm_constants = Load_amm_constants("./amm_constants")
	if augur_srv.db_matic != nil {
		amm_contracts = augur_srv.db_matic.Get_augur_amm_contract_addresses()
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
	host_secure := os.Getenv("AUGUR_HTTPS_HOSTNAME")

	r := gin.New()
	//r.RedirectTrailingSlash=false
	//r.RedirectFixedPath = false
	r.LoadHTMLGlob("html/templates/*/*html")

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

	set_routing_black_templates(r)
	set_routing_api(r)

	r.Static("/black/imgs", "./html/imgs")
	r.Static("/black/res", "./html/res")			// resources (static)


	m := autocert.Manager{
		Prompt:     autocert.AcceptTOS,
		//HostPolicy: autocert.HostWhitelist("api.predictionexplorer.com"),
		HostPolicy: autocert.HostWhitelist(host_secure),
		Cache:      autocert.DirCache(os.Getenv("HOME")+".tls-autocert-cache"),
	}
	_ = m
	// Listen and serve on defined port
	log.Printf("Listening on port %s", port_plain)

	go func() {
		log.Printf("Listening for TLS on interface %v\n",host_secure)
		cer, err := tls.LoadX509KeyPair(
				os.Getenv("HOME")+"/configs/server.crt",
				os.Getenv("HOME")+"/configs/server.key",
		)
		if err != nil {
			log.Println(err)
			return
		}
		tlsConfig := &tls.Config{Certificates: []tls.Certificate{cer}}
		server := http.Server{
			Addr:      host_secure,
			Handler:   r,
			TLSConfig: tlsConfig,
		}
		err = server.ListenAndServeTLS("", "")
		log.Printf("%v",err)
	}()
	if len(port_plain) > 0 {
		go func() {
			r.Run(":" + port_plain)
		}()
	}
	select{} // infinite suspend for the main go-routine
}
