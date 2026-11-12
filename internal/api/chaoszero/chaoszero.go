// Package chaoszero serves the ChaosZero prediction-market JSON API over
// the cz_ tables that cmd/cz-etl maintains. The module is an injected API
// value (no package-level state) mirroring the RandomWalk module shape;
// responses use the v1 {"status":1,"error":"",...} envelope.
package chaoszero

import (
	"context"
	"errors"
	"fmt"
	"log/slog"

	"github.com/PredictionExplorer/augur-explorer/internal/api/common"
	"github.com/PredictionExplorer/augur-explorer/internal/api/httpx"
	"github.com/PredictionExplorer/augur-explorer/internal/store"
	czdb "github.com/PredictionExplorer/augur-explorer/internal/store/chaoszero"
)

// API carries the injected dependencies of the ChaosZero module.
type API struct {
	repo   *czdb.Repo
	store  *store.Store
	logger *slog.Logger
}

// Options carries the module configuration.
type Options struct {
	// Logger receives module diagnostics; nil discards.
	Logger *slog.Logger
}

// New builds the module over the shared store. A nil store yields a bare
// module whose handlers answer the "Database link wasn't configured"
// envelope.
func New(st *store.Store, opts Options) *API {
	logger := opts.Logger
	if logger == nil {
		logger = slog.New(slog.DiscardHandler)
	}
	a := &API{store: st, logger: logger}
	if st != nil {
		a.repo = czdb.NewRepo(st)
	}
	return a
}

// NewBare returns an unloaded module: handlers answer the guard envelopes.
func NewBare() *API {
	return New(nil, Options{})
}

// respondStoreError logs a database failure (unless the client canceled the
// request) and answers HTTP 500 in the {"status":0,...} envelope without
// leaking internal detail.
func (a *API) respondStoreError(c *httpx.Context, err error) {
	if !errors.Is(err, context.Canceled) {
		a.logger.Error(fmt.Sprintf("%s: %v", c.FullPath(), err))
	}
	common.RespondInternalErrorJSON(c)
}

// dbInitialized reports whether the module was constructed with a database
// link (bare modules answer the guard envelopes).
func (a *API) dbInitialized() bool {
	return a.repo != nil && a.store != nil
}

// RegisterRoutes registers all ChaosZero JSON API routes.
func (a *API) RegisterRoutes(r *httpx.Router) {
	r.GET("/api/chaoszero/statistics", a.handleStatistics)
	r.GET("/api/chaoszero/contracts", a.handleContracts)
	r.GET("/api/chaoszero/rounds/list/{offset}/{limit}", a.handleRoundsList)
	r.GET("/api/chaoszero/rounds/info/{round_num}", a.handleRoundInfo)
	r.GET("/api/chaoszero/trades/list/{offset}/{limit}", a.handleTradesList)
	r.GET("/api/chaoszero/trades/by_round/{round_num}/{offset}/{limit}", a.handleTradesByRound)
	r.GET("/api/chaoszero/trades/by_user/{user_addr}/{offset}/{limit}", a.handleTradesByUser)
	r.GET("/api/chaoszero/traders/list/{offset}/{limit}", a.handleTradersList)
	r.GET("/api/chaoszero/funding/list/{offset}/{limit}", a.handleFundingList)
	r.GET("/api/chaoszero/splits/list/{offset}/{limit}", a.handleSplitsList)
	r.GET("/api/chaoszero/merges/list/{offset}/{limit}", a.handleMergesList)
	r.GET("/api/chaoszero/settlements/list/{offset}/{limit}", a.handleSettlementsList)
	r.GET("/api/chaoszero/claims/list/{offset}/{limit}", a.handleClaimsList)
	r.GET("/api/chaoszero/claims/by_user/{user_addr}/{offset}/{limit}", a.handleClaimsByUser)
	r.GET("/api/chaoszero/answers/list/{offset}/{limit}", a.handleAnswersList)
	r.GET("/api/chaoszero/escrows/list/{offset}/{limit}", a.handleEscrowsList)
	r.GET("/api/chaoszero/payouts/list/{offset}/{limit}", a.handlePayoutsList)
	r.GET("/api/chaoszero/transfers/list/{offset}/{limit}", a.handleTransfersList)
	r.GET("/api/chaoszero/transfers/by_user/{user_addr}/{offset}/{limit}", a.handleTransfersByUser)
	r.GET("/api/chaoszero/approvals/list/{offset}/{limit}", a.handleApprovalsList)
}
