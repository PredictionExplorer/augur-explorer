// Package randomwalk serves the frozen v1 RandomWalk NFT JSON API. The
// module is an injected API value (no package-level state); route
// registration is a method the shared router constructor calls.
package randomwalk

import (
	"context"
	"errors"
	"fmt"
	"log/slog"

	"github.com/PredictionExplorer/augur-explorer/internal/api/common"
	"github.com/PredictionExplorer/augur-explorer/internal/api/httpx"
	rwmodel "github.com/PredictionExplorer/augur-explorer/internal/model/randomwalk"
	"github.com/PredictionExplorer/augur-explorer/internal/store"
	rwdb "github.com/PredictionExplorer/augur-explorer/internal/store/randomwalk"
)

// Response format flags for handlers.
const (
	JSON = true
	HTTP = false
)

// legacyNoRowsText is what database/sql produced for a missing row; the
// pre-rewrite handlers surfaced it verbatim in DBError/error fields, so the
// converted not-found paths render the same string to keep the wire format
// byte-identical.
const legacyNoRowsText = "sql: no rows in result set"

// defaultExploreMaxTokenID bounds explore/random token ids when the
// configuration does not override it (RANDOMWALK_EXPLORE_MAX_TOKEN_ID).
const defaultExploreMaxTokenID = 3766

// API carries the injected dependencies of the v1 RandomWalk module: the
// context-first repository, the base store for address lookups, the process
// logger and the module configuration values.
type API struct {
	repo              *rwdb.Repo
	store             *store.Store
	logger            *slog.Logger
	adminAPIKey       string
	rankingAdminKey   string
	voteChainIDs      []int64
	exploreMaxTokenID int64
	assetsPublicBase  string
	assetsFlatPaths   bool
}

// Options carries the module configuration (typed config values that were
// scattered os.Getenv reads before §8.3). The zero value is a working
// default: fail-closed admin routes, any vote chain id, the standard
// explore range and request-derived asset URLs.
type Options struct {
	// Logger receives module diagnostics; nil discards.
	Logger *slog.Logger
	// RankingAdminKey guards the direct Elo-match route
	// (RANKING_ADMIN_KEY); AdminAPIKey (ADMIN_API_KEY) is its fallback.
	// Both empty fails closed: the route answers 503.
	RankingAdminKey string
	AdminAPIKey     string
	// VoteChainIDs allowlists the wallet-signed beauty-vote chain ids
	// (RANKING_VOTE_CHAIN_IDS); empty allows any chain id.
	VoteChainIDs []int64
	// ExploreMaxTokenID bounds explore/random token ids
	// (RANDOMWALK_EXPLORE_MAX_TOKEN_ID); zero applies the 3766 default.
	ExploreMaxTokenID int64
	// AssetsPublicBase overrides the public /images URL base
	// (NFT_ASSETS_PUBLIC_BASE); empty derives it per request or applies
	// the documented production default for metadata.
	AssetsPublicBase string
	// AssetsFlatPaths selects the flat /images/<file> URL layout
	// (NFT_ASSETS_FLAT_PATHS).
	AssetsFlatPaths bool
}

// New builds the module over the shared store. A nil store yields a bare
// module whose handlers answer the legacy "Database link wasn't configured"
// envelope.
func New(st *store.Store, opts Options) *API {
	logger := opts.Logger
	if logger == nil {
		logger = slog.New(slog.DiscardHandler)
	}
	maxTokenID := opts.ExploreMaxTokenID
	if maxTokenID <= 0 {
		maxTokenID = defaultExploreMaxTokenID
	}
	a := &API{
		store:             st,
		logger:            logger,
		adminAPIKey:       opts.AdminAPIKey,
		rankingAdminKey:   opts.RankingAdminKey,
		voteChainIDs:      opts.VoteChainIDs,
		exploreMaxTokenID: maxTokenID,
		assetsPublicBase:  opts.AssetsPublicBase,
		assetsFlatPaths:   opts.AssetsFlatPaths,
	}
	if st != nil {
		a.repo = rwdb.NewRepo(st)
	}
	return a
}

// NewBare returns an unloaded module: handlers answer the legacy guard
// envelopes. The route-drift test uses it to enumerate the route table
// without a database, and the guard tests pin the failure shapes.
func NewBare() *API {
	return New(nil, Options{})
}

// respondStoreError logs a database failure (unless the client canceled the
// request) and answers with HTTP 500 in the legacy {"status":0,...} envelope
// without leaking internal detail. These paths previously terminated the
// whole process from inside the store layer, so no golden constrains them.
func (a *API) respondStoreError(c *httpx.Context, err error) {
	if !errors.Is(err, context.Canceled) {
		a.logger.Error(fmt.Sprintf("%s: %v", c.FullPath(), err))
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
		common.RequireAdminKey("X-Ranking-Admin-Key",
			common.AdminKey{Name: "RANKING_ADMIN_KEY", Value: a.rankingAdminKey},
			common.AdminKey{Name: "ADMIN_API_KEY", Value: a.adminAPIKey}))
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
