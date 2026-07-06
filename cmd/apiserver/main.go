package main

import (
	"context"
	"crypto/tls"
	"errors"
	"fmt"
	"log"
	"net"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"github.com/ethereum/go-ethereum/ethclient"
	ethrpc "github.com/ethereum/go-ethereum/rpc"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/acme/autocert"

	"github.com/PredictionExplorer/augur-explorer/internal/api/common"
	"github.com/PredictionExplorer/augur-explorer/internal/api/cosmicgame"
	"github.com/PredictionExplorer/augur-explorer/internal/api/faq"
	"github.com/PredictionExplorer/augur-explorer/internal/api/randomwalk"
)

var (
	RPC_URL   = os.Getenv("RPC_URL")
	eclient   *ethclient.Client
	rpcclient *ethrpc.Client

	rwcg_srv *RWCGServer

	Error *log.Logger
	Info  *log.Logger
)

// envBoolDefaultTrue returns true if the variable is unset or empty (default on).
// Returns false for "false", "0", "no", "off" (case-insensitive). Any other non-empty value is true.
func envBoolDefaultTrue(key string) bool {
	v := strings.ToLower(strings.TrimSpace(os.Getenv(key)))
	if v == "" {
		return true
	}
	if v == "false" || v == "0" || v == "no" || v == "off" {
		return false
	}
	return true
}

func initialize() {
	// Initialize the common context
	common.InitContext(rwcg_srv.db, eclient, Info, Error)

	enableRWRoutes := envBoolDefaultTrue("ENABLE_ROUTES_RANDOMWALK")
	enableCGRoutes := envBoolDefaultTrue("ENABLE_ROUTES_COSMICGAME")
	enableFAQRoutes := envBoolDefaultTrue("ENABLE_ROUTES_FAQ")

	cosmicgame.Init(eclient, rpcclient, Info, Error, enableCGRoutes)
	if enableCGRoutes {
		Info.Printf("CosmicGame HTTP routes: enabled (ENABLE_ROUTES_COSMICGAME unset or true)")
	} else {
		Info.Printf("CosmicGame HTTP routes: disabled (ENABLE_ROUTES_COSMICGAME=false)")
		fmt.Printf("INFO: CosmicGame API routes are not registered (ENABLE_ROUTES_COSMICGAME=false).\n")
	}

	randomwalk.Init(enableRWRoutes)
	if enableRWRoutes {
		Info.Printf("RandomWalk HTTP routes: enabled (ENABLE_ROUTES_RANDOMWALK unset or true)")
	} else {
		Info.Printf("RandomWalk HTTP routes: disabled (ENABLE_ROUTES_RANDOMWALK=false)")
		fmt.Printf("INFO: RandomWalk API routes are not registered (ENABLE_ROUTES_RANDOMWALK=false).\n")
	}

	faq.Init(Info, Error, enableFAQRoutes)
	if !enableFAQRoutes {
		fmt.Printf("INFO: FAQ bot proxy routes are not registered (ENABLE_ROUTES_FAQ=false).\n")
	}
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

	// Baseline abuse protection: per-IP token bucket across the whole API.
	// Generous enough for legitimate frontends; mutating endpoints add their
	// own stricter limits at registration.
	r.Use(common.RateLimit(50, 100))

	// Prometheus request metrics; /metrics itself is served on METRICS_ADDR.
	r.Use(metricsMiddleware())

	// Liveness/readiness probes and the internal metrics/pprof listener.
	registerHealthRoutes(r)
	startInternalServer()

	// NFT asset files (nft-assets mirror) and optional /static ABI JSON; see static_assets.go
	registerStaticAssetRoutes(r)

	// Set up all routes
	randomwalk.RegisterAPIRoutes(r)
	cosmicgame.RegisterAPIRoutes(r)
	faq.RegisterAPIRoutes(r)

	// Bare ERC-721 tokenURI route. Both projects' on-chain baseURI is
	// https://<host>/metadata/, and this single webserv serves both the
	// RandomWalk and Cosmic Signature hosts, so dispatch by request Host:
	// a Cosmic Signature host serves Cosmic Signature metadata, anything else
	// (RandomWalk hosts) serves RandomWalk metadata.
	r.GET("/metadata/:token_id", func(c *gin.Context) {
		host := strings.ToLower(c.Request.Host)
		if xfh := c.Request.Header.Get("X-Forwarded-Host"); xfh != "" {
			if i := strings.IndexByte(xfh, ','); i >= 0 {
				xfh = xfh[:i]
			}
			host = strings.ToLower(strings.TrimSpace(xfh))
		}
		if strings.Contains(host, "cosmicsignature") {
			cosmicgame.TokenMetadataHandler(c)
			return
		}
		randomwalk.TokenMetadataHandler(c)
	})

	m := &autocert.Manager{
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
