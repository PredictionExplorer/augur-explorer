// HTTP handlers for the ChaosZero JSON API. Every handler answers the v1
// envelope {"status":1,"error":"",...}; list endpoints take {offset}/{limit}
// path parameters and by-user endpoints take a hex {user_addr} that is
// resolved through the shared address table (an unknown address answers an
// empty list, matching the Cosmic Signature by-user convention).
package chaoszero

import (
	"context"
	"errors"
	"net/http"

	"github.com/PredictionExplorer/augur-explorer/internal/api/common"
	"github.com/PredictionExplorer/augur-explorer/internal/api/httpx"
	"github.com/PredictionExplorer/augur-explorer/internal/store"
)

// listHandler answers one offset/limit list endpoint: fetch runs the store
// read and the result is placed under key in the envelope.
func (a *API) listHandler(c *httpx.Context, key string, fetch func(ctx context.Context, offset, limit int) (any, error)) {
	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	if !a.dbInitialized() {
		common.RespondErrorJSON(c, "Database link wasn't configured")
		return
	}
	success, offset, limit := common.ParseOffsetLimitParamsJSON(c)
	if !success {
		return
	}
	rows, err := fetch(c.Request.Context(), offset, limit)
	if err != nil {
		a.respondStoreError(c, err)
		return
	}
	c.JSON(http.StatusOK, httpx.H{
		"status": 1,
		"error":  "",
		"Offset": offset,
		"Limit":  limit,
		key:      rows,
	})
}

// byUserHandler answers one by-user list endpoint. An address absent from
// the address table yields an empty list with UserAid 0.
func (a *API) byUserHandler(c *httpx.Context, key string, fetch func(ctx context.Context, userAid int64, offset, limit int) (any, error)) {
	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	if !a.dbInitialized() {
		common.RespondErrorJSON(c, "Database link wasn't configured")
		return
	}
	pUserAddr := c.Param("user_addr")
	if len(pUserAddr) == 0 {
		common.RespondErrorJSON(c, "'user_addr' parameter is not set")
		return
	}
	success, offset, limit := common.ParseOffsetLimitParamsJSON(c)
	if !success {
		return
	}
	userAid, err := a.store.LookupAddressID(c.Request.Context(), pUserAddr)
	if err != nil {
		c.JSON(http.StatusOK, httpx.H{
			"status": 1, "error": "", "UserAid": int64(0), "UserAddr": pUserAddr,
			"Offset": offset, "Limit": limit, key: []any{},
		})
		return
	}
	rows, err := fetch(c.Request.Context(), userAid, offset, limit)
	if err != nil {
		a.respondStoreError(c, err)
		return
	}
	c.JSON(http.StatusOK, httpx.H{
		"status":   1,
		"error":    "",
		"UserAid":  userAid,
		"UserAddr": pUserAddr,
		"Offset":   offset,
		"Limit":    limit,
		key:        rows,
	})
}

// roundParam parses the {round_num} path parameter.
func (a *API) roundParam(c *httpx.Context) (int64, bool) {
	pRound := c.Param("round_num")
	if len(pRound) == 0 {
		common.RespondErrorJSON(c, "'round_num' parameter is not set")
		return 0, false
	}
	return common.ParseIntFromRemoteOrError(c, JSONResponse, &pRound)
}

// JSONResponse selects the JSON error format of the shared parsers.
const JSONResponse = true

func (a *API) handleStatistics(c *httpx.Context) {
	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	if !a.dbInitialized() {
		common.RespondErrorJSON(c, "Database link wasn't configured")
		return
	}
	stats, err := a.repo.Statistics(c.Request.Context())
	if err != nil {
		a.respondStoreError(c, err)
		return
	}
	c.JSON(http.StatusOK, httpx.H{
		"status":     1,
		"error":      "",
		"Statistics": stats,
	})
}

func (a *API) handleContracts(c *httpx.Context) {
	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	if !a.dbInitialized() {
		common.RespondErrorJSON(c, "Database link wasn't configured")
		return
	}
	market, source, vault, err := a.repo.RawContractAddrs(c.Request.Context())
	if err != nil {
		a.respondStoreError(c, err)
		return
	}
	c.JSON(http.StatusOK, httpx.H{
		"status":     1,
		"error":      "",
		"MarketAddr": market,
		"SourceAddr": source,
		"VaultAddr":  vault,
	})
}

func (a *API) handleRoundsList(c *httpx.Context) {
	a.listHandler(c, "Rounds", func(ctx context.Context, offset, limit int) (any, error) {
		return a.repo.RoundsList(ctx, offset, limit)
	})
}

func (a *API) handleRoundInfo(c *httpx.Context) {
	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	if !a.dbInitialized() {
		common.RespondErrorJSON(c, "Database link wasn't configured")
		return
	}
	roundID, ok := a.roundParam(c)
	if !ok {
		return
	}
	detail, err := a.repo.RoundDetail(c.Request.Context(), roundID)
	if err != nil {
		if errors.Is(err, store.ErrNotFound) {
			common.RespondErrorJSON(c, "Round not found")
			return
		}
		a.respondStoreError(c, err)
		return
	}
	c.JSON(http.StatusOK, httpx.H{
		"status":   1,
		"error":    "",
		"RoundNum": roundID,
		"Round":    detail,
	})
}

func (a *API) handleTradesList(c *httpx.Context) {
	a.listHandler(c, "Trades", func(ctx context.Context, offset, limit int) (any, error) {
		return a.repo.TradesList(ctx, offset, limit)
	})
}

func (a *API) handleTradesByRound(c *httpx.Context) {
	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	if !a.dbInitialized() {
		common.RespondErrorJSON(c, "Database link wasn't configured")
		return
	}
	roundID, ok := a.roundParam(c)
	if !ok {
		return
	}
	success, offset, limit := common.ParseOffsetLimitParamsJSON(c)
	if !success {
		return
	}
	trades, err := a.repo.TradesByRound(c.Request.Context(), roundID, offset, limit)
	if err != nil {
		a.respondStoreError(c, err)
		return
	}
	c.JSON(http.StatusOK, httpx.H{
		"status":   1,
		"error":    "",
		"RoundNum": roundID,
		"Offset":   offset,
		"Limit":    limit,
		"Trades":   trades,
	})
}

func (a *API) handleTradesByUser(c *httpx.Context) {
	a.byUserHandler(c, "Trades", func(ctx context.Context, userAid int64, offset, limit int) (any, error) {
		return a.repo.TradesByUser(ctx, userAid, offset, limit)
	})
}

func (a *API) handleTradersList(c *httpx.Context) {
	a.listHandler(c, "Traders", func(ctx context.Context, offset, limit int) (any, error) {
		return a.repo.TradersList(ctx, offset, limit)
	})
}

func (a *API) handleFundingList(c *httpx.Context) {
	a.listHandler(c, "Fundings", func(ctx context.Context, offset, limit int) (any, error) {
		return a.repo.FundingList(ctx, offset, limit)
	})
}

func (a *API) handleSplitsList(c *httpx.Context) {
	a.listHandler(c, "Splits", func(ctx context.Context, offset, limit int) (any, error) {
		return a.repo.SplitsList(ctx, offset, limit)
	})
}

func (a *API) handleMergesList(c *httpx.Context) {
	a.listHandler(c, "Merges", func(ctx context.Context, offset, limit int) (any, error) {
		return a.repo.MergesList(ctx, offset, limit)
	})
}

func (a *API) handleSettlementsList(c *httpx.Context) {
	a.listHandler(c, "Settlements", func(ctx context.Context, offset, limit int) (any, error) {
		return a.repo.SettlementsList(ctx, offset, limit)
	})
}

func (a *API) handleClaimsList(c *httpx.Context) {
	a.listHandler(c, "Claims", func(ctx context.Context, offset, limit int) (any, error) {
		return a.repo.ClaimsList(ctx, offset, limit)
	})
}

func (a *API) handleClaimsByUser(c *httpx.Context) {
	a.byUserHandler(c, "Claims", func(ctx context.Context, userAid int64, offset, limit int) (any, error) {
		return a.repo.ClaimsByUser(ctx, userAid, offset, limit)
	})
}

func (a *API) handleAnswersList(c *httpx.Context) {
	a.listHandler(c, "Answers", func(ctx context.Context, offset, limit int) (any, error) {
		return a.repo.AnswersList(ctx, offset, limit)
	})
}

func (a *API) handleEscrowsList(c *httpx.Context) {
	a.listHandler(c, "Escrows", func(ctx context.Context, offset, limit int) (any, error) {
		return a.repo.EscrowsList(ctx, offset, limit)
	})
}

func (a *API) handlePayoutsList(c *httpx.Context) {
	a.listHandler(c, "Payouts", func(ctx context.Context, offset, limit int) (any, error) {
		return a.repo.PayoutsList(ctx, offset, limit)
	})
}

func (a *API) handleTransfersList(c *httpx.Context) {
	a.listHandler(c, "Transfers", func(ctx context.Context, offset, limit int) (any, error) {
		return a.repo.TransfersList(ctx, offset, limit)
	})
}

func (a *API) handleTransfersByUser(c *httpx.Context) {
	a.byUserHandler(c, "Transfers", func(ctx context.Context, userAid int64, offset, limit int) (any, error) {
		return a.repo.TransfersByUser(ctx, userAid, offset, limit)
	})
}

func (a *API) handleApprovalsList(c *httpx.Context) {
	a.listHandler(c, "Approvals", func(ctx context.Context, offset, limit int) (any, error) {
		return a.repo.ApprovalsList(ctx, offset, limit)
	})
}
