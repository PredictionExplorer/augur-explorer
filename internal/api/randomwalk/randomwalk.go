// Package randomwalk serves the frozen v1 RandomWalk NFT JSON API. The
// module is an injected API value (no package-level state); route
// registration is a method the shared router constructor calls.
package randomwalk

import (
	"context"
	"errors"
	"fmt"
	"io"
	"log"

	"github.com/PredictionExplorer/augur-explorer/internal/api/httpx"

	"github.com/PredictionExplorer/augur-explorer/internal/api/common"
	rwmodel "github.com/PredictionExplorer/augur-explorer/internal/model/randomwalk"
	"github.com/PredictionExplorer/augur-explorer/internal/store"
	rwdb "github.com/PredictionExplorer/augur-explorer/internal/store/randomwalk"
)

// Response format flags for handlers
const (
	JSON = true
	HTTP = false
)

// legacyNoRowsText is what database/sql produced for a missing row; the
// pre-rewrite handlers surfaced it verbatim in DBError/error fields, so the
// converted not-found paths render the same string to keep the wire format
// byte-identical.
const legacyNoRowsText = "sql: no rows in result set"

// API carries the injected dependencies of the v1 RandomWalk module: the
// context-first repository, the base store for address lookups, and the
// legacy file loggers.
type API struct {
	repo   *rwdb.Repo
	store  *store.Store
	info   *log.Logger
	errlog *log.Logger
}

// New builds the module over the shared store. info and errlog may be nil
// (discard). A nil store yields a bare module whose handlers answer the
// legacy "Database link wasn't configured" envelope.
func New(st *store.Store, info, errlog *log.Logger) *API {
	discard := log.New(io.Discard, "", 0)
	if info == nil {
		info = discard
	}
	if errlog == nil {
		errlog = discard
	}
	a := &API{store: st, info: info, errlog: errlog}
	if st != nil {
		a.repo = rwdb.NewRepo(st)
	}
	return a
}

// NewBare returns an unloaded module: handlers answer the legacy guard
// envelopes. The route-drift test uses it to enumerate the route table
// without a database, and the guard tests pin the failure shapes.
func NewBare() *API {
	return New(nil, nil, nil)
}

// respondStoreError logs a database failure (unless the client canceled the
// request) and answers with HTTP 500 in the legacy {"status":0,...} envelope
// without leaking internal detail. These paths previously terminated the
// whole process from inside the store layer, so no golden constrains them.
func (a *API) respondStoreError(c *httpx.Context, err error) {
	if !errors.Is(err, context.Canceled) {
		errStr := fmt.Sprintf("%s: %v", c.FullPath(), err)
		a.errlog.Print(errStr)
		a.info.Print(errStr)
	}
	common.RespondInternalErrorJSON(c)
}

// rwContractAddrs resolves marketplace + RandomWalk addresses and AIDs from
// rw_contracts (same source as the ETL), answering the request with an
// internal error when the registry is unreadable.
func (a *API) rwContractAddrs(c *httpx.Context) (rwmodel.ContractAddresses, bool) {
	addrs, err := a.repo.ContractAddrs(c.Request.Context())
	if err != nil {
		a.respondStoreError(c, err)
		return rwmodel.ContractAddresses{}, false
	}
	return addrs, true
}

// RegisterRoutes registers all RandomWalk JSON API routes.
func (a *API) RegisterRoutes(r *httpx.Router) {
	r.GET("/api/randomwalk/current_offers/{order_by}", a.handleCurrentOffers)
	r.GET("/api/randomwalk/floor_price", a.handleFloorPrice)
	r.GET("/api/randomwalk/tokens/list/sequential", a.handleTokenListSeq)
	r.GET("/api/randomwalk/tokens/list/by_period/{init_ts}/{fin_ts}", a.handleTokenListPeriod)
	r.GET("/api/randomwalk/tokens/info/{token_id}", a.handleTokenInfo)
	// Same JSON as above; use this path when a reverse proxy only forwards /api/cosmicgame/* to the app.
	r.GET("/api/cosmicgame/randomwalk/tokens/info/{token_id}", a.handleTokenInfo)
	r.GET("/api/randomwalk/tokens/name_changes/{token_id}", a.handleTokenNameHistory)
	r.GET("/api/randomwalk/trading/history/{offset}/{limit}", a.handleTradingHistory)
	r.GET("/api/randomwalk/trading/by_user/{user_aid}/{offset}/{limit}", a.handleTradingHistoryByUser)
	r.GET("/api/randomwalk/trading/sales/{offset}/{limit}", a.handleSaleHistory)
	r.GET("/api/randomwalk/tokens/history/{token_id}/{offset}/{limit}", a.handleTokenHistory)
	r.GET("/api/randomwalk/tokens/by_user/{user_aid}", a.handleTokensByUser)
	r.GET("/api/randomwalk/statistics/by_token", a.handleTokenStats)
	r.GET("/api/randomwalk/statistics/by_market", a.handleMarketStats)
	r.GET("/api/randomwalk/statistics/trading_volume/{init_ts}/{fin_ts}/{interval_secs}", a.handleTradingVolumeByPeriod)
	r.GET("/api/randomwalk/statistics/mint_intervals", a.handleMintIntervals)
	r.GET("/api/randomwalk/statistics/floor_price/{init_ts}/{fin_ts}/{interval_secs}", a.handleFloorPriceOverTime)
	r.GET("/api/randomwalk/statistics/withdrawal_chart", a.handleWithdrawalChart)
	r.GET("/api/randomwalk/user/info/{user_aid}", a.handleUserInfo)
	r.GET("/api/randomwalk/top5tokens", a.handleTop5TradedTokens)
	r.GET("/api/randomwalk/mint_report", a.handleMintReport)
	r.GET("/api/randomwalk/contracts", a.handleContracts)

	// NFT metadata + explore (legacy Python backend parity)
	r.GET("/api/randomwalk/explore/random", a.handleExploreRandom)
	r.GET("/api/randomwalk/random", a.handleExploreRandom)
	r.GET("/api/randomwalk/token-ranking/order", a.handleTokenRankingOrder)
	r.GET("/api/randomwalk/rating_order", a.handleTokenRankingOrder)
	r.GET("/api/randomwalk/vote_count", a.handleVoteCount)
	r.GET("/api/randomwalk/ranking/sign-challenge", a.handleRankingSignChallenge)
	r.GET("/api/randomwalk/ranking/beauty-pair-ids", a.handleRankingBeautyPairIDs)
	// Direct match recording is admin-only (fails closed when no key is set).
	// Public voting goes through /add_game, which is wallet-signature-verified
	// and rate limited instead.
	r.POST("/api/randomwalk/token-ranking/match", a.handleTokenRankingMatch,
		common.RateLimit(2, 5),
		common.RequireAdminKey("X-Ranking-Admin-Key", "RANKING_ADMIN_KEY", "ADMIN_API_KEY"))
	r.POST("/api/randomwalk/add_game", a.handleAddGameLegacy, common.RateLimit(1, 10))
	r.GET("/api/randomwalk/metadata/{token_id}", a.handleTokenMetadata)

	// NOTE: the bare ERC-721 tokenURI route GET /metadata/{tokenID} is registered
	// centrally in main.go with host-aware dispatch, because the same webserv
	// serves both RandomWalk and Cosmic Signature hosts (and the Cosmic Signature
	// contract's tokenURI base is https://nfts.cosmicsignature.com/metadata/).
}

// dbInitialized reports whether the module was constructed with a database
// link (bare modules answer the legacy guard envelopes).
func (a *API) dbInitialized() bool {
	return a.repo != nil && a.store != nil
}
