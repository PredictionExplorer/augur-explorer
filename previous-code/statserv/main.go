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
	. "github.com/PredictionExplorer/augur-explorer/dbs"
	. "github.com/PredictionExplorer/augur-explorer/dbs/balancerv2"

)
var (
	RPC_URL = os.Getenv("RPC_URL")
	SCHEMA_NAME = os.Getenv("DB_SCHEMA")
	ETHPRICE_SCHEMA = os.Getenv("DB_ETHPRICE_SCHEMA")
	rpcclient *ethclient.Client

	Error   *log.Logger
	Info    *log.Logger

	storagew					SQLStorageWrapper
	ethprice_storage			*SQLStorage
	weth_aid					int64
)
const (
	DEFAULT_DB_LOG              = "db.log"
	FMT_JSON					= true
	FMT_HTML					= false
	PARAM_FORCED				= false
	PARAM_OPTIONAL				= true
)
func initialize() {
	log_dir:=fmt.Sprintf("%v/%v",os.Getenv("HOME"),DEFAULT_LOG_DIR)
	os.MkdirAll(log_dir, os.ModePerm)

	fname:=fmt.Sprintf("%v/%v_%v_db.log",log_dir,"webserv_stats",SCHEMA_NAME)
	logfile, err := os.OpenFile(fname, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err!=nil {
		fmt.Printf("Can't start: %v\n",err)
		os.Exit(1)
	}
	dblog := log.New(logfile,"DB: ",log.Ltime|log.Lshortfile)
	storagew.S = Connect_to_storage_with_schema(dblog,SCHEMA_NAME)

	weth_aid,err = storagew.S.Layer1_lookup_address_id(
		storagew.Get_wrapped_eth_contract_address(),
	)
	ethprice_storage = Connect_to_storage_with_schema(dblog,ETHPRICE_SCHEMA)
}
func secure_https(r http.Handler) {
	autotls.Run(r, "localhost")
}
func main() {

	initialize()


	if len(RPC_URL) == 0 {
		fmt.Printf("Configuration error: RPC URL of Ethereum node is not set."+
			"Calls to contracts are disabled. " +
			" Please set RPC_URL environment variable")
	} else {
		var err error
		rpcclient, err = ethclient.Dial(RPC_URL)
		if err != nil {
			log.Fatal(err)
		}
	}


	port_plain := os.Getenv("HTTP_PORT")
	host_secure := os.Getenv("HTTPS_HOSTNAME")

	r := gin.New()
	//r.RedirectTrailingSlash=false
	//r.RedirectFixedPath = false
	r.LoadHTMLGlob("html/templates/*html")

	r.Use(gin.Logger())

	// Main HTML
	r.GET("/", main_page)
	r.GET("/index.html", main_page)
	r.GET("/index.htm", main_page)
	r.GET("/INDEX.HTM", main_page)

	set_routing_html(r)
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
