package main

import (
	"crypto/tls"
	"context"
	"errors"
	"fmt"
	"log"
	"net"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/ethereum/go-ethereum/ethclient"
	ethrpc "github.com/ethereum/go-ethereum/rpc"
	"golang.org/x/crypto/acme/autocert"

	"github.com/PredictionExplorer/augur-explorer/rwcg/dbs"
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

	enableCosmicGame bool
)

// cosmicgameSchemaPresent returns true if CosmicGame tables (cg_contracts) exist in public.
func cosmicgameSchemaPresent(db *dbs.SQLStorage) bool {
	if db == nil {
		return false
	}
	var ok bool
	err := db.Db().QueryRow(`
		SELECT EXISTS (
			SELECT 1 FROM information_schema.tables
			WHERE table_schema = 'public' AND table_name = 'cg_contracts'
		)`).Scan(&ok)
	if err != nil {
		log.Printf("WARN: could not probe for public.cg_contracts: %v — CosmicGame will be disabled", err)
		return false
	}
	return ok
}

// cosmicgameContractsHasRows returns true if public.cg_contracts exists and has at least one row.
// An empty table used to still flip auto mode "on" (table exists), then cosmic_game_init would
// call Get_cosmic_game_contract_addrs → ErrNoRows → os.Exit(1), or periodic RPC refresh could run
// with bogus addresses against the node shared with RandomWalk local dev.
func cosmicgameContractsHasRows(db *dbs.SQLStorage) bool {
	if db == nil || !cosmicgameSchemaPresent(db) {
		return false
	}
	var n int64
	err := db.Db().QueryRow(`SELECT COUNT(*) FROM public.cg_contracts`).Scan(&n)
	if err != nil {
		log.Printf("WARN: could not COUNT public.cg_contracts: %v — CosmicGame will be disabled", err)
		return false
	}
	return n > 0
}

func initialize() {
	// Initialize the common context
	common.InitContext(rwcg_srv.db, eclient, Info, Error)

	// CosmicGame: ENABLE_COSMICGAME=false|0 → off; true|1 → on (fatal if no contract row);
	// unset → auto: on only if public.cg_contracts exists **and has ≥1 row** (RandomWalk-only / empty CG).
	ev := strings.TrimSpace(os.Getenv("ENABLE_COSMICGAME"))
	switch {
	case ev == "false" || ev == "0":
		enableCosmicGame = false
	case ev == "true" || ev == "1":
		if !cosmicgameContractsHasRows(rwcg_srv.db) {
			fmt.Printf("FATAL: ENABLE_COSMICGAME=true requires public.cg_contracts with at least one row.\n")
			if !cosmicgameSchemaPresent(rwcg_srv.db) {
				fmt.Printf("  The table is missing. Create CosmicGame schema or set ENABLE_COSMICGAME=false.\n")
			} else {
				fmt.Printf("  The table exists but is empty. Seed cg_contracts or set ENABLE_COSMICGAME=false.\n")
			}
			os.Exit(1)
		}
		enableCosmicGame = true
	default:
		enableCosmicGame = cosmicgameContractsHasRows(rwcg_srv.db)
		if !enableCosmicGame {
			if !cosmicgameSchemaPresent(rwcg_srv.db) {
				fmt.Printf("INFO: CosmicGame disabled: public.cg_contracts not found (RandomWalk-only DB). Set ENABLE_COSMICGAME=true to require CosmicGame.\n")
				Info.Printf("CosmicGame auto-disabled: cg_contracts not in database")
			} else {
				fmt.Printf("INFO: CosmicGame disabled: public.cg_contracts has no rows (seed contract addresses to enable). Set ENABLE_COSMICGAME=false to silence this.\n")
				Info.Printf("CosmicGame auto-disabled: cg_contracts table is empty")
			}
		}
	}

	cosmicgame.Init(eclient, rpcclient, Info, Error, enableCosmicGame)
	if enableCosmicGame {
		Info.Printf("CosmicGame module enabled")
	} else {
		if ev == "false" || ev == "0" {
			Info.Printf("CosmicGame module disabled (ENABLE_COSMICGAME=false)")
			fmt.Printf("INFO: CosmicGame module disabled by ENABLE_COSMICGAME=false.\n")
		} else {
			Info.Printf("CosmicGame module disabled")
		}
	}

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
	httpsExtra := strings.TrimSpace(os.Getenv("HTTPS_EXTRA_LISTEN_ADDR"))

	r := gin.New()
	r.LoadHTMLGlob("templates/*/*html")

	// Don't trust all proxies (avoids GIN-debug warning; set trusted proxies explicitly if behind one)
	r.SetTrustedProxies(nil)

	// CORS: allow cross-origin requests from browsers (frontend on different origin)
	// Handles OPTIONS preflight and adds CORS headers to all responses
	r.Use(func(c *gin.Context) {
		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Methods", "GET, POST, PUT, PATCH, DELETE, OPTIONS")
		c.Header("Access-Control-Allow-Headers", "Origin, Content-Type, Accept, Authorization, X-Requested-With")
		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(http.StatusNoContent)
			return
		}
		c.Next()
	})

	r.Use(gin.Recovery()) // recover from panics (e.g. broken pipe when client disconnects)
	r.Use(gin.Logger())

	// NFT asset files (nft-assets mirror) and optional /static ABI JSON; see static_assets.go
	registerStaticAssetRoutes(r)

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

	// TLS: HTTPS_HOSTNAME is the primary bind address (e.g. :443 or 0.0.0.0:443).
	// Optional HTTPS_EXTRA_LISTEN_ADDR starts a second TLS listener (e.g. :1443) with the same routes and certs.
	if len(host_secure) > 0 {
		certs := loadHTTPSCertificates()
		if len(certs) == 0 {
			log.Println("TLS: no certificates loaded; HTTPS listeners not started")
		} else {
			tlsConfig := &tls.Config{Certificates: certs}
			startTLS := func(addr string) {
				ln, err := net.Listen("tcp", addr)
				if err != nil {
					log.Printf("HTTPS bind failed on %q: %v", addr, err)
					return
				}
				log.Printf("HTTPS bound and listening on %s", ln.Addr().String())
				server := http.Server{
					Handler:   r,
					TLSConfig: tlsConfig,
				}
				if err := server.ServeTLS(ln, "", ""); err != nil && !errors.Is(err, http.ErrServerClosed) {
					log.Printf("HTTPS server %s: %v", addr, err)
				}
			}
			go startTLS(host_secure)
			if httpsExtra != "" {
				go startTLS(httpsExtra)
			}
		}
	}
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

// loadHTTPSCertificates returns TLS cert/key pairs for one listener. Go picks the right leaf cert per TLS SNI
// when multiple entries are present (e.g. api1.randomwalknft.com + randomwalknft.com on :443).
//
// Primary pair: TLS_CERT_FILE + TLS_KEY_FILE, or $HOME/configs/server.crt + server.key.
// Optional second pair (another domain): TLS_CERT_FILE_2 + TLS_KEY_FILE_2.
func loadHTTPSCertificates() []tls.Certificate {
	home := os.Getenv("HOME")
	cert1 := strings.TrimSpace(os.Getenv("TLS_CERT_FILE"))
	key1 := strings.TrimSpace(os.Getenv("TLS_KEY_FILE"))
	if cert1 == "" {
		cert1 = filepath.Join(home, "configs", "server.crt")
	}
	if key1 == "" {
		key1 = filepath.Join(home, "configs", "server.key")
	}
	var out []tls.Certificate
	if cer, err := tls.LoadX509KeyPair(cert1, key1); err == nil {
		out = append(out, cer)
		log.Printf("TLS: loaded certificate %q", cert1)
	} else {
		log.Printf("TLS: primary cert load failed (%q + %q): %v", cert1, key1, err)
	}
	cert2 := strings.TrimSpace(os.Getenv("TLS_CERT_FILE_2"))
	key2 := strings.TrimSpace(os.Getenv("TLS_KEY_FILE_2"))
	if cert2 != "" && key2 != "" {
		if cer, err := tls.LoadX509KeyPair(cert2, key2); err == nil {
			out = append(out, cer)
			log.Printf("TLS: loaded additional certificate %q (SNI will choose between %d certs)", cert2, len(out))
		} else {
			log.Printf("TLS: TLS_CERT_FILE_2 load failed: %v", err)
		}
	}
	return out
}

