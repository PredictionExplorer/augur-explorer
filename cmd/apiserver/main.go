package main

import (
	"context"
	"crypto/tls"
	"errors"
	"fmt"
	"log"
	"log/slog"
	"net"
	"net/http"
	"os"
	"os/signal"
	"path/filepath"
	"strings"
	"sync"
	"syscall"
	"time"

	"github.com/ethereum/go-ethereum/ethclient"
	ethrpc "github.com/ethereum/go-ethereum/rpc"

	"github.com/PredictionExplorer/augur-explorer/internal/api/common"
	"github.com/PredictionExplorer/augur-explorer/internal/api/cosmicgame"
	"github.com/PredictionExplorer/augur-explorer/internal/api/cosmicgame/contractstate"
	"github.com/PredictionExplorer/augur-explorer/internal/api/faq"
	"github.com/PredictionExplorer/augur-explorer/internal/api/httpx"
	"github.com/PredictionExplorer/augur-explorer/internal/api/randomwalk"
	"github.com/PredictionExplorer/augur-explorer/internal/api/routes"
	"github.com/PredictionExplorer/augur-explorer/internal/api/v2"
)

// shutdownTimeout bounds how long in-flight requests may take to finish once
// SIGTERM/SIGINT arrives before the listeners are closed forcefully.
const shutdownTimeout = 15 * time.Second

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

func initialize(ctx context.Context) *contractstate.State {
	// Initialize the common context
	common.InitContext(rwcg_srv.store, eclient, Info, Error)

	enableRWRoutes := envBoolDefaultTrue("ENABLE_ROUTES_RANDOMWALK")
	enableCGRoutes := envBoolDefaultTrue("ENABLE_ROUTES_COSMICGAME")
	enableFAQRoutes := envBoolDefaultTrue("ENABLE_ROUTES_FAQ")

	cgState, err := cosmicgame.Init(ctx, eclient, rpcclient, Info, Error, enableCGRoutes)
	if err != nil {
		// Startup cannot proceed without the CosmicGame contract registry.
		err_str := fmt.Sprintf("CosmicGame module init failed: %v", err)
		Error.Print(err_str)
		Info.Print(err_str)
		fmt.Printf("\nFATAL: %s\n", err_str)
		fmt.Printf("HINT: If you don't need CosmicGame, set ENABLE_ROUTES_COSMICGAME=false in websrv .env\n\n")
		os.Exit(1)
	}
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
	return cgState
}

func main() {
	if len(RPC_URL) == 0 {
		fmt.Printf("Configuration error: RPC URL of Ethereum node is not set." +
			"Calls to contracts are disabled. " +
			" Please set RPC_URL environment variable\n")
		os.Exit(1)
	}

	// Root context: cancelled on SIGINT/SIGTERM, which starts the drain.
	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	var err error
	rpcclient, err = ethrpc.DialContext(ctx, RPC_URL)
	if err != nil {
		fmt.Printf("Can't establish connection to RPC service: %v\n", err)
		os.Exit(1)
	}
	eclient = ethclient.NewClient(rpcclient)
	rwcg_srv = create_rwcg_server()

	cgState := initialize(ctx)
	cosmicgame.StartBackgroundRefresh(ctx)

	port_plain := os.Getenv("HTTP_PORT")
	host_secure := os.Getenv("HTTPS_HOSTNAME")
	httpsExtra := strings.TrimSpace(os.Getenv("HTTPS_EXTRA_LISTEN_ADDR"))

	// The shared constructor (internal/api/routes) builds the middleware
	// chain — CORS, panic recovery, access log, per-IP rate limiting,
	// Prometheus metrics — and registers the frozen v1 table, generated v2
	// server, health probes, static assets (env-gated, see static_assets.go),
	// and host-dispatched bare /metadata route. The integration suites build
	// through the same constructor.
	accessLogger := slog.New(slog.NewTextHandler(os.Stdout, nil))
	errorLogger := slog.New(slog.NewTextHandler(os.Stderr, nil))
	var v2Server *v2.Server
	if cgState != nil {
		v2Server, err = v2.NewServer(rwcg_srv.store, cgState, errorLogger)
		if err != nil {
			fmt.Printf("API v2 initialization failed: %v\n", err)
			os.Exit(1)
		}
	}
	r := routes.New(rwcg_srv.store, routes.Options{
		AccessLog:     accessLogger,
		PanicLog:      errorLogger,
		V2:            v2Server,
		Extra:         []httpx.Middleware{metricsMiddleware()},
		RegisterExtra: registerStaticAssetRoutes,
	})

	// Internal metrics/pprof listener (METRICS_ADDR).
	internalSrv := startInternalServer()

	// Public listeners. Each runs in its own goroutine and is tracked for
	// the coordinated Shutdown below.
	var servers []*http.Server
	var wg sync.WaitGroup

	serve := func(srv *http.Server, ln net.Listener, useTLS bool) {
		servers = append(servers, srv)
		wg.Add(1)
		go func() {
			defer wg.Done()
			var err error
			if useTLS {
				err = srv.ServeTLS(ln, "", "")
			} else {
				err = srv.Serve(ln)
			}
			if err != nil && !errors.Is(err, http.ErrServerClosed) {
				Error.Printf("HTTP server %s: %v", ln.Addr(), err)
				log.Printf("HTTP server %s: %v", ln.Addr(), err)
			}
		}()
	}

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
				serve(&http.Server{Handler: r, TLSConfig: tlsConfig, ReadHeaderTimeout: 10 * time.Second}, ln, true)
			}
			startTLS(host_secure)
			if httpsExtra != "" {
				startTLS(httpsExtra)
			}
		}
	}
	if len(port_plain) > 0 {
		ln, err := net.Listen("tcp", ":"+port_plain)
		if err != nil {
			fmt.Printf("HTTP bind failed on port %s: %v\n", port_plain, err)
			os.Exit(1)
		}
		log.Printf("Listening on port %s", port_plain)
		serve(&http.Server{Handler: r, ReadHeaderTimeout: 10 * time.Second}, ln, false)
	}
	if len(servers) == 0 {
		fmt.Printf("Configuration error: no listeners configured. Set HTTP_PORT and/or HTTPS_HOSTNAME.\n")
		os.Exit(1)
	}

	// Block until SIGINT/SIGTERM, then drain: readiness flips false so load
	// balancers stop sending traffic, in-flight requests get shutdownTimeout
	// to finish, background refresh loops stop via ctx, and the store pool
	// closes last.
	<-ctx.Done()
	log.Printf("shutdown: signal received, draining (timeout %s)", shutdownTimeout)
	common.SetDraining()

	shutdownCtx, cancel := context.WithTimeout(context.Background(), shutdownTimeout)
	defer cancel()
	for _, srv := range servers {
		if err := srv.Shutdown(shutdownCtx); err != nil {
			Error.Printf("shutdown: %v", err)
			log.Printf("shutdown: %v", err)
		}
	}
	wg.Wait()
	if internalSrv != nil {
		if err := internalSrv.Shutdown(shutdownCtx); err != nil {
			Error.Printf("shutdown internal server: %v", err)
		}
	}
	rwcg_srv.store.Close()
	log.Printf("shutdown: complete")
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
