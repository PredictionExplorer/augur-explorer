package main

import (
	"crypto/tls"
	"context"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/ethereum/go-ethereum/ethclient"
	ethrpc "github.com/ethereum/go-ethereum/rpc"
	"golang.org/x/crypto/acme/autocert"

	"github.com/PredictionExplorer/augur-explorer/rwcg/websrv/api/common"
	"github.com/PredictionExplorer/augur-explorer/rwcg/websrv/api/cosmicgame"
	"github.com/PredictionExplorer/augur-explorer/rwcg/websrv/api/randomwalk"
)

var (
	RPC_URL   = os.Getenv("RPC_URL")
	eclient   *ethclient.Client
	rpcclient *ethrpc.Client

	rwcg_srv *RWCGServer

	Error *log.Logger
	Info  *log.Logger
)

func initialize() {
	// Initialize the common context
	common.InitContext(rwcg_srv.db, rwcg_srv.db_arbitrum, eclient, Info, Error)

	// Initialize CosmicGame
	cosmicgame.Init(eclient, rpcclient, Info, Error)

	// Initialize RandomWalk
	randomwalk.Init()
}

func main() {
	if len(RPC_URL) == 0 {
		fmt.Printf("Configuration error: RPC URL of Ethereum node is not set." +
			"Calls to contracts are disabled. " +
			" Please set RPC_URL environment variable\n")
		os.Exit(1)
	}
	var err error
	rpcclient, err = ethrpc.DialContext(context.Background(), RPC_URL)
	if err != nil {
		fmt.Printf("Can't establish connection to RPC service: %v\n", err)
		os.Exit(1)
	}
	eclient = ethclient.NewClient(rpcclient)
	rwcg_srv = create_rwcg_server()

	initialize()

	port_plain := os.Getenv("HTTP_PORT")
	host_secure := os.Getenv("HTTPS_HOSTNAME")

	r := gin.New()
	r.LoadHTMLGlob("templates/*/*html")

	r.Use(gin.Logger())

	// Static files
	r.Static("/black/imgs", "./html/imgs")
	r.Static("/black/res", "./html/res")

	// Main index page
	r.GET("/black/", main_page)
	r.GET("/black/index.html", main_page)

	// Set up all routes
	randomwalk.RegisterHTMLRoutes(r)
	randomwalk.RegisterAPIRoutes(r)
	cosmicgame.RegisterHTMLRoutes(r)
	cosmicgame.RegisterAPIRoutes(r)

	m := autocert.Manager{
		Prompt:     autocert.AcceptTOS,
		HostPolicy: autocert.HostWhitelist(host_secure),
		Cache:      autocert.DirCache(os.Getenv("HOME") + "/.tls-autocert-cache"),
	}
	_ = m
	log.Printf("Listening on port %s", port_plain)

	go func() {
		log.Printf("Listening for TLS on interface %v\n", host_secure)
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
		log.Printf("%v", err)
	}()
	if len(port_plain) > 0 {
		go func() {
			r.Run(":" + port_plain)
		}()
	}
	select {} // infinite suspend for the main go-routine
}

func main_page(c *gin.Context) {
	c.HTML(http.StatusOK, "rw_index.html", gin.H{
		"title": "RandomWalk & CosmicGame",
	})
}
