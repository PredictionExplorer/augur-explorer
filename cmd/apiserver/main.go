// The RWCG API server: serves the frozen v1 JSON API, the generated v2 API,
// health probes, metrics and env-gated static assets. main wires the
// process-wide dependencies (typed configuration, process logger, RPC
// client, store, the injected API modules) through the shared router
// constructor (internal/api/routes) and runs tracked HTTP listeners with
// coordinated graceful shutdown.
package main

import (
	"context"
	"crypto/tls"
	"errors"
	"fmt"
	"io"
	"log/slog"
	"net"
	"net/http"
	"os"
	"os/signal"
	"path/filepath"
	"sync"
	"syscall"
	"time"

	"github.com/ethereum/go-ethereum/ethclient"
	ethrpc "github.com/ethereum/go-ethereum/rpc"

	"github.com/PredictionExplorer/augur-explorer/internal/api/common"
	"github.com/PredictionExplorer/augur-explorer/internal/api/cosmicgame"
	"github.com/PredictionExplorer/augur-explorer/internal/api/faq"
	"github.com/PredictionExplorer/augur-explorer/internal/api/httpx"
	"github.com/PredictionExplorer/augur-explorer/internal/api/randomwalk"
	"github.com/PredictionExplorer/augur-explorer/internal/api/routes"
	v2 "github.com/PredictionExplorer/augur-explorer/internal/api/v2"
	"github.com/PredictionExplorer/augur-explorer/internal/config"
	"github.com/PredictionExplorer/augur-explorer/internal/version"
)

// shutdownTimeout bounds how long in-flight requests may take to finish once
// SIGTERM/SIGINT arrives before the listeners are closed forcefully.
const shutdownTimeout = 15 * time.Second

// Public-listener timeouts. Request bodies are small JSON documents (the
// largest are the ranking POSTs), so a slow-body client is bounded by
// readTimeout on top of the header bound. idleTimeout reaps abandoned
// keep-alive connections. There is deliberately NO WriteTimeout: several
// frozen v1 endpoints stream unbounded arrays, and a write cap would sever
// legitimate large responses to slow clients — revisit when v1 is removed
// (docs/MODERNIZATION.md §6.3).
const (
	readHeaderTimeout = 10 * time.Second
	readTimeout       = 30 * time.Second
	idleTimeout       = 120 * time.Second
)

// newPublicServer builds one public listener's http.Server with the
// production timeout policy. tlsConfig is cloned per server: ServeTLS
// mutates the config (HTTP/2 NextProtos), so sharing one pointer across
// listeners was a data race.
func newPublicServer(handler http.Handler, tlsConfig *tls.Config) *http.Server {
	srv := &http.Server{
		Handler:           handler,
		ReadHeaderTimeout: readHeaderTimeout,
		ReadTimeout:       readTimeout,
		IdleTimeout:       idleTimeout,
	}
	if tlsConfig != nil {
		srv.TLSConfig = tlsConfig.Clone()
	}
	return srv
}

// buildModules constructs the enabled v1 API modules over the shared
// dependencies. A disabled module stays nil: its routes are not registered
// and the shared /metadata dispatch answers the legacy unavailable envelope.
func buildModules(
	ctx context.Context,
	cfg *config.APIServer,
	deps *serverDeps,
	ethClient *ethclient.Client,
	rpcClient *ethrpc.Client,
) (*cosmicgame.API, *randomwalk.API, *faq.Proxy, error) {
	var cgAPI *cosmicgame.API
	if cfg.EnableCosmicGame {
		var err error
		cgAPI, err = cosmicgame.New(ctx, cosmicgame.Config{
			Store:            deps.store,
			EthClient:        ethClient,
			RPCClient:        rpcClient,
			Logger:           deps.logger,
			AdminAPIKey:      cfg.AdminAPIKey,
			AssetsPublicBase: cfg.NFTAssetsPublicBase,
		})
		if err != nil {
			// Startup cannot proceed without the CosmicGame contract registry.
			errStr := fmt.Sprintf("CosmicGame module init failed: %v", err)
			deps.logger.Error(errStr)
			return nil, nil, nil, fmt.Errorf("FATAL: %s\nHINT: If you don't need CosmicGame, set ENABLE_ROUTES_COSMICGAME=false in websrv .env", errStr)
		}
		deps.logger.Info("CosmicGame HTTP routes: enabled (ENABLE_ROUTES_COSMICGAME unset or true)")
	} else {
		deps.logger.Info("CosmicGame HTTP routes: disabled (ENABLE_ROUTES_COSMICGAME=false)")
	}

	var rwAPI *randomwalk.API
	if cfg.EnableRandomWalk {
		rwAPI = randomwalk.New(deps.store, randomwalk.Options{
			Logger:            deps.logger,
			AdminAPIKey:       cfg.AdminAPIKey,
			RankingAdminKey:   cfg.RankingAdminKey,
			VoteChainIDs:      cfg.RankingVoteChainIDs,
			ExploreMaxTokenID: cfg.ExploreMaxTokenID,
			AssetsPublicBase:  cfg.NFTAssetsPublicBase,
			AssetsFlatPaths:   cfg.NFTAssetsFlatPaths,
		})
		deps.logger.Info("RandomWalk HTTP routes: enabled (ENABLE_ROUTES_RANDOMWALK unset or true)")
	} else {
		deps.logger.Info("RandomWalk HTTP routes: disabled (ENABLE_ROUTES_RANDOMWALK=false)")
	}

	var faqProxy *faq.Proxy
	if cfg.EnableFAQ {
		faqProxy = faq.New(faq.Options{UpstreamURL: cfg.FAQUpstream(), Logger: deps.logger})
	} else {
		deps.logger.Info("FAQ proxy disabled (ENABLE_ROUTES_FAQ=false)")
	}

	return cgAPI, rwAPI, faqProxy, nil
}

// osExit is stubbed by tests that drive main through its failure arm.
var osExit = os.Exit

func main() {
	if version.HandleFlag(os.Args[1:], os.Stdout) {
		return
	}
	if err := runMain(); err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		osExit(1)
	}
}

// runMain owns the signal-scoped context so its deferred cleanup always runs
// before main decides the exit code (os.Exit skips deferred calls).
func runMain() error {
	// Root context: cancelled on SIGINT/SIGTERM, which starts the drain.
	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()
	return run(ctx, os.Getenv, os.Stdout)
}

// run wires every dependency and serves until ctx is cancelled (returning
// nil after the graceful drain) or a startup error occurs. Environment
// access goes through getenv and structured logs go to logOut so tests can
// inject configuration and inspect records.
func run(ctx context.Context, getenv func(string) string, logOut io.Writer) error {
	cfg, err := config.LoadAPIServer(getenv)
	if err != nil {
		return err
	}
	logger := cfg.Log.NewLogger(logOut)
	logger.LogAttrs(ctx, slog.LevelInfo, "build info", version.LogAttrs()...)
	logger.LogAttrs(ctx, slog.LevelInfo, "effective configuration", config.Attrs(cfg)...)

	rpcClient, err := ethrpc.DialContext(ctx, cfg.RPCURL)
	if err != nil {
		return fmt.Errorf("can't establish connection to RPC service: %w", err)
	}
	ethClient := ethclient.NewClient(rpcClient)
	deps, err := newServerDeps(ctx, cfg, logger)
	if err != nil {
		return err
	}
	defer deps.store.Close()

	cgAPI, rwAPI, faqProxy, err := buildModules(ctx, cfg, deps, ethClient, rpcClient)
	if err != nil {
		return err
	}
	if cgAPI != nil {
		cgAPI.StartBackgroundRefresh(ctx)
	}

	// The shared constructor (internal/api/routes) builds the middleware
	// chain — CORS, panic recovery, access log, per-IP rate limiting,
	// Prometheus metrics — and registers the injected v1 modules, generated
	// v2 server, health probes, static assets (config-gated, see
	// static_assets.go), and host-dispatched bare /metadata route. The
	// integration suites build through the same constructor.
	var v2Server *v2.Server
	if cgAPI != nil {
		v2Server, err = v2.NewServer(deps.store, cgAPI.ContractState(), logger,
			v2.WithAdmin(v2.AdminConfig{
				AdminKeys: []common.AdminKey{
					{Name: "ADMIN_API_KEY", Value: cfg.AdminAPIKey},
				},
			}),
			v2.WithRanking(v2.RankingConfig{
				AdminKeys: []common.AdminKey{
					{Name: "RANKING_ADMIN_KEY", Value: cfg.RankingAdminKey},
					{Name: "ADMIN_API_KEY", Value: cfg.AdminAPIKey},
				},
				VoteChainIDs:      cfg.RankingVoteChainIDs,
				ExploreMaxTokenID: cfg.ExploreMaxTokenID,
			}))
		if err != nil {
			return fmt.Errorf("API v2 initialization failed: %w", err)
		}
	}
	r := routes.New(deps.store, routes.Options{
		AccessLog:     logger,
		PanicLog:      logger,
		CosmicGame:    cgAPI,
		RandomWalk:    rwAPI,
		FAQ:           faqProxy,
		V2:            v2Server,
		V1SunsetAt:    cfg.V1SunsetAt,
		Extra:         []httpx.Middleware{metricsMiddleware()},
		RegisterExtra: registerStaticAssetRoutes(staticAssetsConfig(cfg, logger)),
	})

	// Internal metrics/pprof listener (METRICS_ADDR).
	internalSrv := startInternalServer(cfg.MetricsAddr, logger)

	// Public listeners. Each runs in its own goroutine and is tracked for
	// the coordinated Shutdown below.
	var servers []*http.Server
	var wg sync.WaitGroup

	serve := func(srv *http.Server, ln net.Listener, useTLS bool) {
		servers = append(servers, srv)
		wg.Go(func() {
			var err error
			if useTLS {
				err = srv.ServeTLS(ln, "", "")
			} else {
				err = srv.Serve(ln)
			}
			if err != nil && !errors.Is(err, http.ErrServerClosed) {
				logger.Error(fmt.Sprintf("HTTP server %s: %v", ln.Addr(), err))
			}
		})
	}

	// TLS: HTTPS_HOSTNAME is the primary bind address (e.g. :443 or 0.0.0.0:443).
	// Optional HTTPS_EXTRA_LISTEN_ADDR starts a second TLS listener (e.g. :1443) with the same routes and certs.
	if len(cfg.HTTPSHostname) > 0 {
		certs := loadHTTPSCertificates(cfg, getenv("HOME"), logger)
		if len(certs) == 0 {
			logger.Warn("TLS: no certificates loaded; HTTPS listeners not started")
		} else {
			tlsConfig := &tls.Config{Certificates: certs}
			startTLS := func(addr string) {
				ln, err := new(net.ListenConfig).Listen(ctx, "tcp", addr)
				if err != nil {
					logger.Error(fmt.Sprintf("HTTPS bind failed on %q: %v", addr, err))
					return
				}
				logger.Info("HTTPS bound and listening", "addr", ln.Addr().String())
				serve(newPublicServer(r, tlsConfig), ln, true)
			}
			startTLS(cfg.HTTPSHostname)
			if cfg.HTTPSExtraListenAddr != "" {
				startTLS(cfg.HTTPSExtraListenAddr)
			}
		}
	}
	if len(cfg.HTTPPort) > 0 {
		ln, err := new(net.ListenConfig).Listen(ctx, "tcp", ":"+cfg.HTTPPort)
		if err != nil {
			return fmt.Errorf("HTTP bind failed on port %s: %w", cfg.HTTPPort, err)
		}
		logger.Info("listening", "port", cfg.HTTPPort)
		serve(newPublicServer(r, nil), ln, false)
	}
	if len(servers) == 0 {
		return errors.New("configuration error: no listeners started — check the TLS certificate paths")
	}

	// Block until SIGINT/SIGTERM, then drain: readiness flips false so load
	// balancers stop sending traffic, in-flight requests get shutdownTimeout
	// to finish, background refresh loops stop via ctx, and the store pool
	// closes last (deferred).
	<-ctx.Done()
	logger.Info("shutdown: signal received, draining", "timeout", shutdownTimeout.String())
	common.SetDraining()

	shutdownCtx, cancel := context.WithTimeout(context.Background(), shutdownTimeout)
	defer cancel()
	for _, srv := range servers {
		if err := srv.Shutdown(shutdownCtx); err != nil {
			logger.Error(fmt.Sprintf("shutdown: %v", err))
		}
	}
	wg.Wait()
	if internalSrv != nil {
		if err := internalSrv.Shutdown(shutdownCtx); err != nil {
			logger.Error(fmt.Sprintf("shutdown internal server: %v", err))
		}
	}
	logger.Info("shutdown: complete")
	return nil
}

// loadHTTPSCertificates returns TLS cert/key pairs for one listener. Go picks the right leaf cert per TLS SNI
// when multiple entries are present (e.g. api1.randomwalknft.com + randomwalknft.com on :443).
//
// Primary pair: TLS_CERT_FILE + TLS_KEY_FILE, or $HOME/configs/server.crt + server.key.
// Optional second pair (another domain): TLS_CERT_FILE_2 + TLS_KEY_FILE_2.
func loadHTTPSCertificates(cfg *config.APIServer, home string, logger *slog.Logger) []tls.Certificate {
	cert1 := cfg.TLSCertFile
	key1 := cfg.TLSKeyFile
	if cert1 == "" {
		cert1 = filepath.Join(home, "configs", "server.crt")
	}
	if key1 == "" {
		key1 = filepath.Join(home, "configs", "server.key")
	}
	var out []tls.Certificate
	if cer, err := tls.LoadX509KeyPair(cert1, key1); err == nil {
		out = append(out, cer)
		logger.Info("TLS: loaded certificate", "cert", cert1)
	} else {
		logger.Error(fmt.Sprintf("TLS: primary cert load failed (%q + %q): %v", cert1, key1, err))
	}
	if cfg.TLSCertFile2 != "" && cfg.TLSKeyFile2 != "" {
		if cer, err := tls.LoadX509KeyPair(cfg.TLSCertFile2, cfg.TLSKeyFile2); err == nil {
			out = append(out, cer)
			logger.Info("TLS: loaded additional certificate (SNI selects per hostname)",
				"cert", cfg.TLSCertFile2, "total", len(out))
		} else {
			logger.Error(fmt.Sprintf("TLS: TLS_CERT_FILE_2 load failed: %v", err))
		}
	}
	return out
}
