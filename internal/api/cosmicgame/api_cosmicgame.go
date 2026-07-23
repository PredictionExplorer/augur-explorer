package cosmicgame

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"math"
	"math/big"
	"net/http"
	"reflect"
	"strings"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	ethcommon "github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"

	cg "github.com/PredictionExplorer/augur-explorer/contracts/cosmicgame"
	"github.com/PredictionExplorer/augur-explorer/internal/api/common"
	"github.com/PredictionExplorer/augur-explorer/internal/api/httpx"
	cgmodel "github.com/PredictionExplorer/augur-explorer/internal/model/cosmicgame"
	"github.com/PredictionExplorer/augur-explorer/internal/store"
	cgdb "github.com/PredictionExplorer/augur-explorer/internal/store/cosmicgame"
)

// Legacy output-mode flags passed to the common.ParseIntFromRemoteOrError
// helpers. Only the JSON branch survives (the HTML explorer is gone), but
// the parameter — and both named values — are frozen with the v1 surface.
const (
	JSON = true
	HTTP = false
)

// respondUserAddrNotIndexedJSON returns 200 with empty user-shaped data when the wallet
// is not in the indexer DB yet (common for new connections).
func respondUserAddrNotIndexedUserInfoJSON(c *httpx.Context, userAddr string) {
	emptyUserInfo := httpx.H{
		"AddressId": int64(0), "Address": userAddr, "NumPrizes": int64(0), "NumBids": int64(0),
		"MaxWinAmount": 0.0, "MaxBidAmount": 0.0, "SumRaffleEthWinnings": 0.0, "SumRaffleEthWithdrawal": 0.0,
		"NumRaffleEthWinnings": int64(0), "RaffleNFTsCount": int64(0), "RewardNFTsCount": int64(0),
		"UnclaimedNFTs": int64(0), "TotalCSTokensWon": int64(0), "CosmicSignatureNumTransfers": int64(0),
		"TotalDonatedCount": int64(0), "TotalDonatedAmountEth": 0.0,
		"StakingStatisticsRWalk": httpx.H{},
	}
	c.JSON(http.StatusOK, httpx.H{
		"status": 1, "error": "", "UserInfo": emptyUserInfo,
		"Bids": []any{}, "MainPrizeClaims": []any{}, "PrizeHistory": []any{},
		"TokenDonationsMade": httpx.H{"NFTDonations": []any{}, "ERC20Donations": []any{}},
		"ETHDonationsMade":   []any{}, "MarketingRewardsAwarded": []any{},
		"DonatedNFTsClaimed": []any{}, "DonatedTokensClaimed": []any{},
		"UnclaimedAssets":       httpx.H{"ETHPrizes": []any{}, "DonatedNFTs": []any{}},
		"CurrentlyStakedTokens": httpx.H{"CST": []any{}, "RWalk": []any{}},
		"StakingActions":        httpx.H{"CST": []any{}, "RWalk": []any{}},
		"ERC20Transfers":        []any{}, "ERC721Transfers": []any{},
		"CosmicSignatureTokensOwned": []any{},
	})
}

// respondStoreError logs a store-layer failure and answers with the legacy
// JSON error envelope and HTTP 500; internal details never reach the client.
// A cancelled request context (client went away) is not worth logging.
// These paths previously killed the whole process via os.Exit(1) inside the
// store methods, so any response at all is an improvement.
func (a *API) respondStoreError(c *httpx.Context, err error) {
	if !errors.Is(err, context.Canceled) {
		errStr := fmt.Sprintf("%s: %v", c.FullPath(), err)
		a.logger.Error(errStr)
	}
	common.RespondInternalErrorJSON(c)
}

// safeFloat64 returns 0 for NaN/±Inf so that encoding/json does not panic.
func safeFloat64(f float64) float64 {
	if math.IsNaN(f) || math.IsInf(f, 0) {
		return 0
	}
	return f
}

// sanitizeMapFloatsForJSON recurses into maps and slices and replaces any float64 NaN/±Inf with 0.
// Used so that c.JSON never sees NaN even if a value was missed earlier.
func sanitizeMapFloatsForJSON(m map[string]any) {
	for k, v := range m {
		if v == nil {
			continue
		}
		switch val := v.(type) {
		case float64:
			if math.IsNaN(val) || math.IsInf(val, 0) {
				m[k] = 0.0
			}
		case map[string]any:
			sanitizeMapFloatsForJSON(val)
		case []any:
			for i, e := range val {
				if f, ok := e.(float64); ok && (math.IsNaN(f) || math.IsInf(f, 0)) {
					val[i] = 0.0
				} else if m2, ok := e.(map[string]any); ok {
					sanitizeMapFloatsForJSON(m2)
				} else if s2, ok := e.([]any); ok {
					sanitizeSliceFloatsForJSON(s2)
				}
			}
		default:
			// Struct values (CurRoundStats, MainStats, etc.) were already sanitized before being put in the map
			break
		}
	}
}

func sanitizeSliceFloatsForJSON(s []any) {
	for i, e := range s {
		if f, ok := e.(float64); ok && (math.IsNaN(f) || math.IsInf(f, 0)) {
			s[i] = 0.0
		} else if m2, ok := e.(map[string]any); ok {
			sanitizeMapFloatsForJSON(m2)
		} else if s2, ok := e.([]any); ok {
			sanitizeSliceFloatsForJSON(s2)
		}
	}
}

// sanitizeFloatsForJSON recursively replaces NaN and ±Inf float64 with 0 in v (in-place).
// Pass a pointer to a struct so fields can be modified. Prevents "json: unsupported value: NaN" panic.
func sanitizeFloatsForJSON(v any) {
	if v == nil {
		return
	}
	val := reflect.ValueOf(v)
	if val.Kind() == reflect.Pointer {
		val = val.Elem()
	}
	if !val.IsValid() {
		return
	}
	switch val.Kind() {
	case reflect.Float64:
		f := val.Float()
		if math.IsNaN(f) || math.IsInf(f, 0) {
			// v must be a settable pointer to float64
			if p := reflect.ValueOf(v); p.Kind() == reflect.Pointer && p.Elem().CanSet() {
				p.Elem().SetFloat(0)
			}
		}
	case reflect.Struct:
		for _, f := range val.Fields() {
			if !f.CanSet() {
				continue
			}
			switch f.Kind() {
			case reflect.Float64:
				x := f.Float()
				if math.IsNaN(x) || math.IsInf(x, 0) {
					f.SetFloat(0)
				}
			case reflect.Struct:
				sanitizeFloatsForJSON(f.Addr().Interface())
			case reflect.Pointer:
				if !f.IsNil() {
					sanitizeFloatsForJSON(f.Interface())
				}
			case reflect.Slice:
				for j := range f.Len() {
					sanitizeFloatsForJSON(f.Index(j).Addr().Interface())
				}
			default:
				// Other field kinds cannot carry float64 values.
			}
		}
	case reflect.Pointer:
		if !val.IsNil() {
			sanitizeFloatsForJSON(val.Interface())
		}
	case reflect.Slice:
		for i := range val.Len() {
			sanitizeFloatsForJSON(val.Index(i).Addr().Interface())
		}
	default:
		// Other kinds cannot carry float64 values.
	}
}

// =============================================================================
// DASHBOARD & STATISTICS API
// =============================================================================

func (a *API) handleDashboard(c *httpx.Context) {
	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")

	snap := a.state.Snapshot()

	// Use live contract price when cache is empty or failed (avoids frontend showing 0 before first refresh or after RPC errors)
	responseBidPrice, responseBidPriceEth := snap.BidPrice, snap.BidPriceEth
	if (responseBidPrice == "" || responseBidPrice == "error" || responseBidPriceEth == 0) && a.ethClient != nil {
		if contract, err := cg.NewCosmicSignatureGame(snap.Addrs.CosmicGame, a.ethClient); err == nil {
			var copts bind.CallOpts
			if tmpVal, err := contract.GetNextEthBidPrice(&copts); err == nil {
				responseBidPrice = tmpVal.String()
				fDivisor := big.NewFloat(0.0).SetInt(big.NewInt(1e18))
				fBidPrice := big.NewFloat(0.0).SetInt(tmpVal)
				fQuo := big.NewFloat(0.0).Quo(fBidPrice, fDivisor)
				responseBidPriceEth, _ = fQuo.Float64()
				a.state.SetBidPrice(responseBidPrice, responseBidPriceEth)
			}
		}
	}

	caddrs, err := a.repo.ContractAddrs(c.Request.Context())
	if err != nil {
		a.respondStoreError(c, err)
		return
	}
	curRoundStats, err := a.repo.CosmicGameRoundStatistics(c.Request.Context(), snap.RoundNum)
	if err != nil {
		a.respondStoreError(c, err)
		return
	}
	// Live contract: delay before round activation (always); activation time when DB has none
	if a.ethClient != nil {
		if contract, err := cg.NewCosmicSignatureGame(snap.Addrs.CosmicGame, a.ethClient); err == nil {
			var copts bind.CallOpts
			if d, err := contract.DelayDurationBeforeRoundActivation(&copts); err == nil {
				curRoundStats.DelayDurationBeforeRoundActivation = d.Int64()
			}
			if curRoundStats.ActivationTime == 0 && snap.RoundNum >= 0 {
				if actTime, err := contract.RoundActivationTime(&copts); err == nil {
					curRoundStats.ActivationTime = actTime.Int64()
				}
			}
		}
	}
	cgBalance := a.state.CosmicGameBalanceEth(c.Request.Context())
	curNumBids := int64(0)
	if snap.RoundNum >= 0 {
		curNumBids, err = a.repo.BidCountForRound(c.Request.Context(), snap.RoundNum)
		if err != nil {
			a.respondStoreError(c, err)
			return
		}
	}
	// Sanitize floats so JSON encoding never sees NaN/Inf (avoids "json: unsupported value: NaN" panic)
	sanitizeFloatsForJSON(&curRoundStats)
	bwStatsCopy := snap.Stats
	sanitizeFloatsForJSON(&bwStatsCopy)
	reqStatus := 1
	errStr := ""
	payload := httpx.H{
		"status":                               reqStatus,
		"error":                                errStr,
		"CosmicGameAddr":                       snap.Addrs.CosmicGame,
		"CosmicGameBalanceEth":                 safeFloat64(cgBalance),
		"CosmicSignatureAddr":                  snap.Addrs.CosmicSignature,
		"CosmicSignatureTokenAddr":             snap.Addrs.CosmicToken,
		"CharityWalletAddr":                    snap.Addrs.CharityWallet,
		"BidPrice":                             responseBidPrice,
		"BidPriceEth":                          safeFloat64(responseBidPriceEth),
		"PrizeClaimDate":                       time.Unix(snap.PrizeClaimTimestamp, 0).Format(time.RFC822),
		"PrizeClaimTs":                         snap.PrizeClaimTimestamp,
		"CurRoundNum":                          snap.RoundNum,
		"CurNumBids":                           curNumBids,
		"PrizeAmount":                          snap.PrizeAmount,
		"PrizeAmountEth":                       safeFloat64(snap.PrizeAmountEth),
		"RaffleAmount":                         snap.RaffleAmount,
		"RaffleAmountEth":                      safeFloat64(snap.RaffleAmountEth),
		"StakingAmount":                        snap.StakingAmount,
		"StakingAmountEth":                     safeFloat64(snap.StakingAmountEth),
		"TotalPrizes":                          bwStatsCopy.TotalPrizes,
		"TotalPrizeAwards":                     bwStatsCopy.TotalPrizeAwards,
		"CgPrizeRowCount":                      bwStatsCopy.CgPrizeRowCount,
		"TotalPrizesPaidAmountEth":             bwStatsCopy.TotalPrizesPaidAmountEth,
		"TotalEthDonatedAmount":                bwStatsCopy.TotalEthDonatedAmount,
		"TotalEthDonatedAmountEth":             bwStatsCopy.TotalEthDonatedAmountEth,
		"LastBidderAddr":                       snap.LastBidder.String(),
		"NumVoluntaryDonations":                bwStatsCopy.NumVoluntaryDonations,
		"SumVoluntaryDonationsEth":             bwStatsCopy.SumVoluntaryDonationsEth,
		"NumRwalkTokensUsed":                   bwStatsCopy.NumRwalkTokensUsed,
		"PriceIncrease":                        snap.PriceIncrease,
		"TimeIncrease":                         snap.TimeIncrease,
		"MainPrizeTimeIncrementInMicroSeconds": snap.MainPrizeTimeIncrement,
		"InitialSecondsUntilPrize":             snap.InitialSecondsUntilPrize,
		"TimeoutClaimPrize":                    snap.TimeoutClaimPrize,
		"RoundStartCSTAuctionLength":           snap.RoundStartAuctionLength,
		"CstDutchAuctionDurationChangeDivisor": snap.CSTAuctionDurationChangeDivisor,
		"ContractMechanicsVersion":             snap.MechanicsVersion,
		"TokenReward":                          snap.TokenReward,
		"PrizePercentage":                      snap.PrizePercentage,
		"RafflePercentage":                     snap.RafflePercentage,
		"ChronoWarriorPercentage":              snap.ChronoPercentage,
		"StakingPercentage":                    snap.StakingPercentage,
		"CharityAddr":                          snap.CharityAddr.String(),
		"CharityPercentage":                    snap.CharityPercentage,
		"CharityBalance":                       snap.CharityBalance,
		"CharityBalanceEth":                    safeFloat64(snap.CharityBalanceEth),
		"NumRaffleEthWinnersBidding":           snap.RaffleEthWinnersBidding,
		"NumRaffleNFTWinnersBidding":           snap.RaffleNFTWinnersBidding,
		"NumRaffleNFTWinnersStakingRWalk":      snap.RaffleNFTWinnersStakingRWalk,
		"NumUniqueBidders":                     bwStatsCopy.NumUniqueBidders,
		"NumUniqueWinners":                     bwStatsCopy.NumUniqueWinners,
		"NumUniqueDonors":                      bwStatsCopy.NumUniqueDonors,
		"NumUniqueStakersCST":                  bwStatsCopy.NumUniqueStakersCST,
		"NumUniqueStakersRWalk":                bwStatsCopy.NumUniqueStakersRWalk,
		"NumUniqueStakersBoth":                 bwStatsCopy.NumUniqueStakersBoth,
		"NumDonatedNFTs":                       bwStatsCopy.NumDonatedNFTs,
		"MainStats":                            bwStatsCopy,
		"CurRoundStats":                        curRoundStats,
		"TsRoundStart":                         snap.RoundStartTimestamp,
		"ContractAddrs":                        caddrs,
	}
	if snap.MechanicsVersion == 3 {
		payload["V3Config"] = httpx.H{
			"IsV3":                                   true,
			"RoundLateBidDurationDivisor":            snap.V3.RoundLateBidDurationDivisor,
			"RoundLateBidDurationSeconds":            snap.V3.RoundLateBidDurationSeconds,
			"RoundLateBidPremiumBaseMultiplier":      snap.V3.RoundLateBidPricePremiumAmountBaseMultiplier,
			"RoundLateBidPremiumExponent":            snap.V3.RoundLateBidPricePremiumAmountExponent,
			"LastBidderBidCstRewardAmountPercentage": snap.V3.LastBidderBidCstRewardAmountPercentage,
			"MainPrizeNumCosmicSignatureNfts":        snap.V3.MainPrizeNumCosmicSignatureNfts,
			"CstAuctionPriceMinLimit":                snap.V3.CstDutchAuctionBeginningBidPriceMinLimit,
			"CstRewardPerTimeIncrement":              snap.V3.BidCstRewardAmountPerMainPrizeTimeIncrement,
		}
	}
	sanitizeMapFloatsForJSON(payload)
	c.JSON(http.StatusOK, payload)
}

// =============================================================================
// ROUNDS & PRIZES API
// =============================================================================

func (a *API) handlePrizeList(c *httpx.Context) {
	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	if !a.dbInitialized() {
		common.RespondErrorJSON(c, "Database link wasn't configured")
		return
	}
	success, offset, limit := common.ParseOffsetLimitParamsJSON(c)
	if !success {
		return
	}
	prizes, err := a.repo.PrizeClaims(c.Request.Context(), offset, limit)
	if err != nil {
		a.respondStoreError(c, err)
		return
	}

	reqStatus := 1
	errStr := ""
	c.JSON(http.StatusOK, httpx.H{
		"status": reqStatus,
		"error":  errStr,
		"Rounds": prizes,
	})
}

// =============================================================================
// BIDS API
// =============================================================================

func (a *API) handleBidList(c *httpx.Context) {
	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	if !a.dbInitialized() {
		common.RespondErrorJSON(c, "Database link wasn't configured")
		return
	}
	success, offset, limit := common.ParseOffsetLimitParamsJSON(c)
	if !success {
		return
	}
	bids, err := a.repo.Bids(c.Request.Context(), offset, limit)
	if err != nil {
		a.respondStoreError(c, err)
		return
	}

	reqStatus := 1
	errStr := ""
	c.JSON(http.StatusOK, httpx.H{
		"status": reqStatus,
		"error":  errStr,
		"Bids":   bids,
		"Offset": offset,
		"Limit":  limit,
	})
}

func (a *API) handleBidListByRound(c *httpx.Context) {
	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	if !a.dbInitialized() {
		common.RespondErrorJSON(c, "Database link wasn't configured")
		return
	}

	pRoundNum := c.Param("round_num")
	var roundNum int64
	if len(pRoundNum) > 0 {
		var success bool
		roundNum, success = common.ParseIntFromRemoteOrError(c, JSON, &pRoundNum)
		if !success {
			return
		}
	}
	pSort := c.Param("sort")
	var sort int64
	if len(pSort) > 0 {
		var success bool
		sort, success = common.ParseIntFromRemoteOrError(c, JSON, &pSort)
		if !success {
			return
		}
	}
	success, offset, limit := common.ParseOffsetLimitParamsJSON(c)
	if !success {
		return
	}

	bids, totalRows, err := a.repo.BidsByRound(c.Request.Context(), roundNum, int(sort), offset, limit)
	if err != nil {
		a.respondStoreError(c, err)
		return
	}
	reqStatus := 1
	errStr := ""
	c.JSON(http.StatusOK, httpx.H{
		"status":      reqStatus,
		"error":       errStr,
		"RoundNum":    roundNum,
		"Offset":      offset,
		"Limit":       limit,
		"Sort":        sort,
		"BidsByRound": bids,
		"TotalRows":   totalRows,
	})
}

type bidMessageTxJSON struct {
	EvtLogID  int64  `json:"evtLogId"`
	BlockNum  int64  `json:"blockNum"`
	TxID      int64  `json:"txId"`
	TxHash    string `json:"txHash"`
	TimeStamp int64  `json:"timeStamp"`
	DateTime  string `json:"dateTime"`
}

type bidWithMessageJSON struct {
	EvtLogID                   int64            `json:"evtLogId"`
	GesturePosition            int64            `json:"gesturePosition"`
	RoundNum                   int64            `json:"roundNum"`
	ParticipantAddr            string           `json:"participantAddr"`
	TimeStamp                  int64            `json:"timeStamp"`
	DateTime                   string           `json:"dateTime"`
	Message                    string           `json:"message"`
	BidCstRewardAmount         string           `json:"BidCstRewardAmount"`
	BidCstRewardAmountEth      float64          `json:"BidCstRewardAmountEth"`
	CstDutchAuctionDuration    string           `json:"CstDutchAuctionDuration"`
	CstDutchAuctionDurationInt int64            `json:"CstDutchAuctionDurationInt"`
	Tx                         bidMessageTxJSON `json:"tx"`
}

func bidRecToWithMessageJSON(rec cgmodel.CGBidRec) bidWithMessageJSON {
	return bidWithMessageJSON{
		EvtLogID:                   rec.Tx.EvtLogId,
		GesturePosition:            rec.BidPosition,
		RoundNum:                   rec.RoundNum,
		ParticipantAddr:            rec.BidderAddr,
		TimeStamp:                  rec.Tx.TimeStamp,
		DateTime:                   rec.Tx.DateTime,
		Message:                    rec.Message,
		BidCstRewardAmount:         rec.BidCstRewardAmount,
		BidCstRewardAmountEth:      rec.BidCstRewardAmountEth,
		CstDutchAuctionDuration:    rec.CstDutchAuctionDuration,
		CstDutchAuctionDurationInt: rec.CstDutchAuctionDurationInt,
		Tx: bidMessageTxJSON{
			EvtLogID:  rec.Tx.EvtLogId,
			BlockNum:  rec.Tx.BlockNum,
			TxID:      rec.Tx.TxId,
			TxHash:    rec.Tx.TxHash,
			TimeStamp: rec.Tx.TimeStamp,
			DateTime:  rec.Tx.DateTime,
		},
	}
}

func (a *API) respondBidInfoJSON(c *httpx.Context, evtlogID int64) {
	bidInfo, err := a.repo.BidInfo(c.Request.Context(), evtlogID)
	if err != nil {
		if errors.Is(err, store.ErrNotFound) {
			common.RespondErrorJSON(c, "record not found")
			return
		}
		a.respondStoreError(c, err)
		return
	}
	c.JSON(http.StatusOK, httpx.H{
		"status":  1,
		"error":   "",
		"BidInfo": bidInfo,
	})
}

func (a *API) handleBidInfo(c *httpx.Context) {
	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	if !a.dbInitialized() {
		common.RespondErrorJSON(c, "Database link wasn't configured")
		return
	}

	pEvtlogID := c.Param("evtlog_id")
	var evtlogID int64
	if len(pEvtlogID) > 0 {
		var success bool
		evtlogID, success = common.ParseIntFromRemoteOrError(c, JSON, &pEvtlogID)
		if !success {
			return
		}
	}
	a.respondBidInfoJSON(c, evtlogID)
}

func (a *API) handleBidInfoByPos(c *httpx.Context) {
	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	if !a.dbInitialized() {
		common.RespondErrorJSON(c, "Database link wasn't configured")
		return
	}

	pRoundNum := c.Param("round_num")
	var roundNum int64
	if len(pRoundNum) > 0 {
		var success bool
		roundNum, success = common.ParseIntFromRemoteOrError(c, JSON, &pRoundNum)
		if !success {
			return
		}
	}

	pBidPosition := c.Param("bid_position")
	var bidPosition int64
	if len(pBidPosition) > 0 {
		var success bool
		bidPosition, success = common.ParseIntFromRemoteOrError(c, JSON, &pBidPosition)
		if !success {
			return
		}
	}

	evtlogID, err := a.repo.EvtlogIDByRoundAndBidPosition(c.Request.Context(), roundNum, bidPosition)
	if err != nil {
		if errors.Is(err, store.ErrNotFound) {
			common.RespondErrorJSON(c, "record not found")
			return
		}
		a.respondStoreError(c, err)
		return
	}
	a.respondBidInfoJSON(c, evtlogID)
}

func (a *API) handleBidWithMessageByRound(c *httpx.Context) {
	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	if !a.dbInitialized() {
		common.RespondErrorJSON(c, "Database link wasn't configured")
		return
	}

	pRound := c.Param("round")
	var roundNum int64
	if len(pRound) > 0 {
		var success bool
		roundNum, success = common.ParseIntFromRemoteOrError(c, JSON, &pRound)
		if !success {
			return
		}
	}

	sortDesc := strings.EqualFold(c.Query("sort"), "desc")
	offset := cgdb.ParseOptionalIntQuery(c.Query("offset"), 0)
	limit := cgdb.ParseOptionalIntQuery(c.Query("limit"), 1000)
	if limit <= 0 {
		limit = 1000
	}

	bids, err := a.repo.BidsWithMessageByRound(c.Request.Context(), roundNum, sortDesc, offset, limit)
	if err != nil {
		a.respondStoreError(c, err)
		return
	}
	messages := make([]bidWithMessageJSON, 0, len(bids))
	for _, bid := range bids {
		messages = append(messages, bidRecToWithMessageJSON(bid))
	}

	c.JSON(http.StatusOK, httpx.H{
		"status":   1,
		"error":    "",
		"Bids":     bids,
		"messages": messages,
		"offset":   offset,
		"limit":    limit,
	})
}

func (a *API) handleRoundInfo(c *httpx.Context) {
	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	if !a.dbInitialized() {
		common.RespondErrorJSON(c, "Database link wasn't configured")
		return
	}

	pPrizeNum := c.Param("prize_num")
	var prizeNum int64
	if len(pPrizeNum) > 0 {
		var success bool
		prizeNum, success = common.ParseIntFromRemoteOrError(c, JSON, &pPrizeNum)
		if !success {
			return
		}
	}
	prizeInfo, err := a.repo.PrizeInfo(c.Request.Context(), prizeNum)
	if errors.Is(err, store.ErrNotFound) {
		common.RespondErrorJSON(c, "record not found")
		return
	}
	if err != nil {
		a.respondStoreError(c, err)
		return
	}
	reqStatus := 1
	errStr := ""
	c.JSON(http.StatusOK, httpx.H{
		"status":    reqStatus,
		"error":     errStr,
		"RoundInfo": prizeInfo,
	})
}

func (a *API) handleUserInfo(c *httpx.Context) {
	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	if !a.dbInitialized() {
		common.RespondErrorJSON(c, "Database link wasn't configured")
		return
	}

	pUserAddr := c.Param("user_addr")
	userAid, err := a.store.LookupAddressID(c.Request.Context(), pUserAddr)
	if err != nil {
		respondUserAddrNotIndexedUserInfoJSON(c, pUserAddr)
		return
	}

	userInfo, err := a.repo.UserInfo(c.Request.Context(), userAid)
	if errors.Is(err, store.ErrNotFound) {
		respondUserAddrNotIndexedUserInfoJSON(c, pUserAddr)
		return
	}
	if err != nil {
		a.respondStoreError(c, err)
		return
	}
	bids, err := a.repo.BidsByUser(c.Request.Context(), userAid)
	if err != nil {
		a.respondStoreError(c, err)
		return
	}
	mainPrizeClaims, err := a.repo.PrizeClaimsByUser(c.Request.Context(), userAid)
	if err != nil {
		a.respondStoreError(c, err)
		return
	}
	prizeHistory, err := a.repo.PrizeHistoryByUser(c.Request.Context(), userAid, 0, 1000)
	if err != nil {
		a.respondStoreError(c, err)
		return
	}

	// Get donations made by user
	nftDonations, err := a.repo.NFTDonationsByUser(c.Request.Context(), userAid)
	if err != nil {
		a.respondStoreError(c, err)
		return
	}
	erc20Donations, err := a.repo.ERC20DonationsByUser(c.Request.Context(), userAid)
	if err != nil {
		a.respondStoreError(c, err)
		return
	}
	ethDonations, err := a.repo.EthDonationsByUser(c.Request.Context(), userAid)
	if err != nil {
		a.respondStoreError(c, err)
		return
	}

	// Get marketing rewards awarded to user
	marketingRewards, err := a.repo.MarketingRewardsByUser(c.Request.Context(), userAid)
	if err != nil {
		a.respondStoreError(c, err)
		return
	}

	// Get unclaimed assets from PrizesWallet
	unclaimedEth, err := a.repo.UnclaimedPrizeEthDeposits(c.Request.Context(), userAid, 0, 1000)
	if err != nil {
		a.respondStoreError(c, err)
		return
	}
	unclaimedNfts, err := a.repo.UnclaimedDonatedNFTsByUser(c.Request.Context(), userAid)
	if err != nil {
		a.respondStoreError(c, err)
		return
	}

	// Get claimed donated assets
	claimedNfts, err := a.repo.DonatedNFTClaimsByUser(c.Request.Context(), userAid)
	if err != nil {
		a.respondStoreError(c, err)
		return
	}
	claimedTokens, err := a.repo.ERC20DonatedPrizesByWinner(c.Request.Context(), userAid)
	if err != nil {
		a.respondStoreError(c, err)
		return
	}

	// Get staking data
	stakedCst, err := a.repo.StakedTokensCstByUser(c.Request.Context(), userAid)
	if err != nil {
		a.respondStoreError(c, err)
		return
	}
	stakedRwalk, err := a.repo.StakedTokensRwalkByUser(c.Request.Context(), userAid)
	if err != nil {
		a.respondStoreError(c, err)
		return
	}
	stakeActionsCst, err := a.repo.StakingActionsCstByUser(c.Request.Context(), userAid, 0, 1000)
	if err != nil {
		a.respondStoreError(c, err)
		return
	}
	stakeActionsRwalk, err := a.repo.StakingActionsRwalkByUser(c.Request.Context(), userAid, 0, 1000)
	if err != nil {
		a.respondStoreError(c, err)
		return
	}

	// Get token transfer history
	erc20Transfers, err := a.repo.CosmicTokenTransfersByUser(c.Request.Context(), userAid, 0, 1000)
	if err != nil {
		a.respondStoreError(c, err)
		return
	}
	erc721Transfers, err := a.repo.CosmicSignatureTransfersByUser(c.Request.Context(), userAid, 0, 1000)
	if err != nil {
		a.respondStoreError(c, err)
		return
	}

	// Get owned Cosmic Signature NFTs (where user is current owner)
	ownedCstNfts, err := a.repo.CosmicSignatureTokensByUser(c.Request.Context(), userAid, 0, 1000)
	if err != nil {
		a.respondStoreError(c, err)
		return
	}

	reqStatus := 1
	errStr := ""
	c.JSON(http.StatusOK, httpx.H{
		"status":          reqStatus,
		"error":           errStr,
		"UserInfo":        userInfo,
		"Bids":            bids,
		"MainPrizeClaims": mainPrizeClaims,
		"PrizeHistory":    prizeHistory,
		"TokenDonationsMade": httpx.H{
			"NFTDonations":   nftDonations,
			"ERC20Donations": erc20Donations,
		},
		"ETHDonationsMade":        ethDonations,
		"MarketingRewardsAwarded": marketingRewards,
		"DonatedNFTsClaimed":      claimedNfts,
		"DonatedTokensClaimed":    claimedTokens,
		"UnclaimedAssets": httpx.H{
			"ETHPrizes":   unclaimedEth,
			"DonatedNFTs": unclaimedNfts,
		},
		"CurrentlyStakedTokens": httpx.H{
			"CST":   stakedCst,
			"RWalk": stakedRwalk,
		},
		"StakingActions": httpx.H{
			"CST":   stakeActionsCst,
			"RWalk": stakeActionsRwalk,
		},
		"ERC20Transfers":             erc20Transfers,
		"ERC721Transfers":            erc721Transfers,
		"CosmicSignatureTokensOwned": ownedCstNfts,
	})
}

// =============================================================================
// CHARITY DONATIONS API
// =============================================================================

func (a *API) handleCharityCosmicgameDeposits(c *httpx.Context) {
	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	if !a.dbInitialized() {
		common.RespondErrorJSON(c, "Database link wasn't configured")
		return
	}
	cosmicgameAid, err := a.store.LookupAddressID(c.Request.Context(), a.state.Snapshot().Addrs.CosmicGame.String())
	if err != nil {
		// Previously this path killed the whole API server (os.Exit); a
		// missing registry row or DB blip now answers HTTP 500 instead.
		a.respondStoreError(c, err)
		return
	}

	donations, err := a.repo.CharityDonationsFromCosmicGame(c.Request.Context(), cosmicgameAid)
	if err != nil {
		a.respondStoreError(c, err)
		return
	}
	reqStatus := 1
	errStr := ""
	c.JSON(http.StatusOK, httpx.H{
		"status":           reqStatus,
		"error":            errStr,
		"CharityDonations": donations,
	})
}

func (a *API) handleCharityVoluntaryDeposits(c *httpx.Context) {
	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	if !a.dbInitialized() {
		common.RespondErrorJSON(c, "Database link wasn't configured")
		return
	}
	cosmicgameAid, err := a.store.LookupAddressID(c.Request.Context(), a.state.Snapshot().Addrs.CosmicGame.String())
	if err != nil {
		// Previously this path killed the whole API server (os.Exit); a
		// missing registry row or DB blip now answers HTTP 500 instead.
		a.respondStoreError(c, err)
		return
	}

	donations, err := a.repo.CharityDonationsVoluntary(c.Request.Context(), cosmicgameAid)
	if err != nil {
		a.respondStoreError(c, err)
		return
	}
	reqStatus := 1
	errStr := ""
	c.JSON(http.StatusOK, httpx.H{
		"status":           reqStatus,
		"error":            errStr,
		"CharityDonations": donations,
	})
}

func (a *API) handleCharityDonationsDeposits(c *httpx.Context) {
	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	if !a.dbInitialized() {
		common.RespondErrorJSON(c, "Database link wasn't configured")
		return
	}
	cosmicgameAid, err := a.store.LookupAddressID(c.Request.Context(), a.state.Snapshot().Addrs.CosmicGame.String())
	if err != nil {
		// Previously this path killed the whole API server (os.Exit); a
		// missing registry row or DB blip now answers HTTP 500 instead.
		a.respondStoreError(c, err)
		return
	}

	donations, err := a.repo.CharityDonations(c.Request.Context(), cosmicgameAid)
	if err != nil {
		a.respondStoreError(c, err)
		return
	}
	reqStatus := 1
	errStr := ""
	c.JSON(http.StatusOK, httpx.H{
		"status":           reqStatus,
		"error":            errStr,
		"CharityDonations": donations,
	})
}

func (a *API) handleCharityDonationsWithdrawals(c *httpx.Context) {
	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	if !a.dbInitialized() {
		common.RespondErrorJSON(c, "Database link wasn't configured")
		return
	}

	withdrawals, err := a.repo.CharityWalletWithdrawals(c.Request.Context())
	if err != nil {
		a.respondStoreError(c, err)
		return
	}
	reqStatus := 1
	errStr := ""
	c.JSON(http.StatusOK, httpx.H{
		"status":             reqStatus,
		"error":              errStr,
		"CharityWithdrawals": withdrawals,
	})
}

// =============================================================================
// UNIQUE PARTICIPANTS API
// =============================================================================

func (a *API) handleUserUniqueBidders(c *httpx.Context) {
	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	if !a.dbInitialized() {
		common.RespondErrorJSON(c, "Database link wasn't configured")
		return
	}

	uniqueBidders, err := a.repo.UniqueBidders(c.Request.Context())
	if err != nil {
		a.respondStoreError(c, err)
		return
	}

	reqStatus := 1
	errStr := ""
	c.JSON(http.StatusOK, httpx.H{
		"status":        reqStatus,
		"error":         errStr,
		"UniqueBidders": uniqueBidders,
	})
}

func (a *API) handleUserUniqueWinners(c *httpx.Context) {
	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	if !a.dbInitialized() {
		common.RespondErrorJSON(c, "Database link wasn't configured")
		return
	}

	uniqueWinners, err := a.repo.UniqueWinners(c.Request.Context())
	if err != nil {
		a.respondStoreError(c, err)
		return
	}

	reqStatus := 1
	errStr := ""
	c.JSON(http.StatusOK, httpx.H{
		"status":        reqStatus,
		"error":         errStr,
		"UniqueWinners": uniqueWinners,
	})
}

func (a *API) handleRoiLeaderboard(c *httpx.Context) {
	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	if !a.dbInitialized() {
		common.RespondErrorJSON(c, "Database link wasn't configured")
		return
	}

	minBids := max(cgdb.ParseOptionalIntQuery(c.Query("min_bids"), 5), 0)
	sortBy := c.Query("sort") // one of: net_pl(default), roi, winrate, spent, nfts, bids
	offset := max(cgdb.ParseOptionalIntQuery(c.Query("offset"), 0), 0)
	limit := cgdb.ParseOptionalIntQuery(c.Query("limit"), 100)
	if limit <= 0 || limit > 1000 {
		limit = 100
	}

	leaderboard, err := a.repo.RoiLeaderboard(c.Request.Context(), minBids, sortBy, offset, limit)
	if err != nil {
		a.respondStoreError(c, err)
		return
	}

	c.JSON(http.StatusOK, httpx.H{
		"status":         1,
		"error":          "",
		"MinBids":        minBids,
		"Sort":           sortBy,
		"Offset":         offset,
		"Limit":          limit,
		"RoiLeaderboard": leaderboard,
	})
}

func (a *API) handleClaimsByRound(c *httpx.Context) {
	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	if !a.dbInitialized() {
		common.RespondErrorJSON(c, "Database link wasn't configured")
		return
	}

	claims, err := a.repo.ClaimsByRound(c.Request.Context())
	if err != nil {
		a.respondStoreError(c, err)
		return
	}

	c.JSON(http.StatusOK, httpx.H{
		"status":        1,
		"error":         "",
		"ClaimsByRound": claims,
	})
}

func (a *API) handleClaimDetailByRound(c *httpx.Context) {
	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	if !a.dbInitialized() {
		common.RespondErrorJSON(c, "Database link wasn't configured")
		return
	}

	round := cgdb.ParseOptionalIntQuery(c.Param("round_num"), -1)
	if round < 0 {
		common.RespondErrorJSON(c, "Invalid round_num")
		return
	}

	detail, err := a.repo.ClaimDetailByRound(c.Request.Context(), int64(round))
	if err != nil {
		a.respondStoreError(c, err)
		return
	}

	c.JSON(http.StatusOK, httpx.H{
		"status":            1,
		"error":             "",
		"RoundNum":          detail.RoundNum,
		"ClaimTransactions": detail.ClaimTransactions,
		"AttachedTokens":    detail.AttachedTokens,
	})
}

func (a *API) handleUserUniqueDonors(c *httpx.Context) {
	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	if !a.dbInitialized() {
		common.RespondErrorJSON(c, "Database link wasn't configured")
		return
	}

	uniqueDonors, err := a.repo.UniqueDonors(c.Request.Context())
	if err != nil {
		a.respondStoreError(c, err)
		return
	}

	reqStatus := 1
	errStr := ""
	c.JSON(http.StatusOK, httpx.H{
		"status":       reqStatus,
		"error":        errStr,
		"UniqueDonors": uniqueDonors,
	})
}

// =============================================================================
// NFT DONATIONS API
// =============================================================================

func (a *API) handleDonationsNftList(c *httpx.Context) {
	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	if !a.dbInitialized() {
		common.RespondErrorJSON(c, "Database link wasn't configured")
		return
	}

	success, offset, limit := common.ParseOffsetLimitParamsJSON(c)
	if !success {
		return
	}
	nftDonations, err := a.repo.NFTDonations(c.Request.Context(), offset, limit)
	if err != nil {
		a.respondStoreError(c, err)
		return
	}

	reqStatus := 1
	errStr := ""
	c.JSON(http.StatusOK, httpx.H{
		"status":       reqStatus,
		"error":        errStr,
		"NFTDonations": nftDonations,
		"Offset":       offset,
		"Limit":        limit,
	})
}

func (a *API) handleNftDonationStats(c *httpx.Context) {
	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	if !a.dbInitialized() {
		common.RespondErrorJSON(c, "Database link wasn't configured")
		return
	}

	nftDonationStats, err := a.repo.NFTDonationStats(c.Request.Context())
	if err != nil {
		a.respondStoreError(c, err)
		return
	}

	reqStatus := 1
	errStr := ""
	c.JSON(http.StatusOK, httpx.H{
		"status":           reqStatus,
		"error":            errStr,
		"NFTDonationStats": nftDonationStats,
	})
}

func (a *API) handleNftDonationsByUser(c *httpx.Context) {
	// DONOR PERSPECTIVE: Returns NFTs this user DONATED

	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	if !a.dbInitialized() {
		common.RespondErrorJSON(c, "Database link wasn't configured")
		return
	}
	pUserAddr := c.Param("user_addr")
	userAid, err := a.store.LookupAddressID(c.Request.Context(), pUserAddr)
	if err != nil {
		c.JSON(http.StatusOK, httpx.H{
			"status": 1, "error": "", "NFTDonationsByDonor": []any{},
			"UserAddr": pUserAddr, "UserAid": int64(0),
		})
		return
	}

	donations, err := a.repo.NFTDonationsByUser(c.Request.Context(), userAid)
	if err != nil {
		a.respondStoreError(c, err)
		return
	}
	reqStatus := 1
	errStr := ""
	c.JSON(http.StatusOK, httpx.H{
		"status":              reqStatus,
		"error":               errStr,
		"NFTDonationsByDonor": donations,
		"UserAddr":            pUserAddr,
		"UserAid":             userAid,
	})
}

func (a *API) handleRecordCounters(c *httpx.Context) {
	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	if !a.dbInitialized() {
		common.RespondErrorJSON(c, "Database link wasn't configured")
		return
	}

	recordCounters, err := a.repo.RecordCounters(c.Request.Context())
	if err != nil {
		a.respondStoreError(c, err)
		return
	}

	reqStatus := 1
	errStr := ""
	c.JSON(http.StatusOK, httpx.H{
		"status":         reqStatus,
		"error":          errStr,
		"RecordCounters": recordCounters,
	})
}

func (a *API) handleDonatedNftInfo(c *httpx.Context) {
	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	if !a.dbInitialized() {
		common.RespondErrorJSON(c, "Database link wasn't configured")
		return
	}

	pRecordID := c.Param("record_id")
	var recordID int64
	if len(pRecordID) > 0 {
		var success bool
		recordID, success = common.ParseIntFromRemoteOrError(c, JSON, &pRecordID)
		if !success {
			return
		}
	}
	nftdonation, err := a.repo.NFTDonationInfo(c.Request.Context(), recordID)
	if err != nil {
		if errors.Is(err, store.ErrNotFound) {
			common.RespondErrorJSON(c, "Record not found")
			return
		}
		a.respondStoreError(c, err)
		return
	}
	c.JSON(http.StatusOK, httpx.H{
		"status":      1,
		"error":       "",
		"NFTDonation": nftdonation,
	})
}

// =============================================================================
// PRIZE & RAFFLE DEPOSITS API
// =============================================================================

func (a *API) handlePrizeDepositsList(c *httpx.Context) {
	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	if !a.dbInitialized() {
		common.RespondErrorJSON(c, "Database link wasn't configured")
		return
	}

	var offset, limit int
	pOffset := c.Param("offset")
	pLimit := c.Param("limit")

	if len(pOffset) == 0 || len(pLimit) == 0 {
		offset = 0
		limit = 10000
	} else {
		var success bool
		success, offset, limit = common.ParseOffsetLimitParamsJSON(c)
		if !success {
			return
		}
	}

	deposits, err := a.repo.PrizeEthDeposits(c.Request.Context(), offset, limit)
	if err != nil {
		a.respondStoreError(c, err)
		return
	}

	reqStatus := 1
	errStr := ""
	c.JSON(http.StatusOK, httpx.H{
		"status":         reqStatus,
		"error":          errStr,
		"RaffleDeposits": deposits,
		"Offset":         offset,
		"Limit":          limit,
	})
}

func (a *API) handleAllEthDepositsList(c *httpx.Context) {
	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	if !a.dbInitialized() {
		common.RespondErrorJSON(c, "Database link wasn't configured")
		return
	}

	var offset, limit int
	pOffset := c.Param("offset")
	pLimit := c.Param("limit")

	if len(pOffset) == 0 || len(pLimit) == 0 {
		offset = 0
		limit = 10000
	} else {
		var success bool
		success, offset, limit = common.ParseOffsetLimitParamsJSON(c)
		if !success {
			return
		}
	}

	deposits, err := a.repo.PrizeEthDeposits(c.Request.Context(), offset, limit)
	if err != nil {
		a.respondStoreError(c, err)
		return
	}

	reqStatus := 1
	errStr := ""
	c.JSON(http.StatusOK, httpx.H{
		"status":      reqStatus,
		"error":       errStr,
		"AllDeposits": deposits,
		"Offset":      offset,
		"Limit":       limit,
	})
}

func (a *API) handleRaffleEthDepositsList(c *httpx.Context) {
	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	if !a.dbInitialized() {
		common.RespondErrorJSON(c, "Database link wasn't configured")
		return
	}

	var offset, limit int
	pOffset := c.Param("offset")
	pLimit := c.Param("limit")

	if len(pOffset) == 0 || len(pLimit) == 0 {
		offset = 0
		limit = 10000
	} else {
		var success bool
		success, offset, limit = common.ParseOffsetLimitParamsJSON(c)
		if !success {
			return
		}
	}

	deposits, err := a.repo.RaffleEthDeposits(c.Request.Context(), offset, limit)
	if err != nil {
		a.respondStoreError(c, err)
		return
	}

	reqStatus := 1
	errStr := ""
	c.JSON(http.StatusOK, httpx.H{
		"status":         reqStatus,
		"error":          errStr,
		"RaffleDeposits": deposits,
		"Offset":         offset,
		"Limit":          limit,
	})
}

func (a *API) handleChronowarriorEthDepositsList(c *httpx.Context) {
	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	if !a.dbInitialized() {
		common.RespondErrorJSON(c, "Database link wasn't configured")
		return
	}

	var offset, limit int
	pOffset := c.Param("offset")
	pLimit := c.Param("limit")

	if len(pOffset) == 0 || len(pLimit) == 0 {
		offset = 0
		limit = 10000
	} else {
		var success bool
		success, offset, limit = common.ParseOffsetLimitParamsJSON(c)
		if !success {
			return
		}
	}

	deposits, err := a.repo.ChronoWarriorEthDeposits(c.Request.Context(), offset, limit)
	if err != nil {
		a.respondStoreError(c, err)
		return
	}

	reqStatus := 1
	errStr := ""
	c.JSON(http.StatusOK, httpx.H{
		"status":                reqStatus,
		"error":                 errStr,
		"ChronoWarriorDeposits": deposits,
		"Offset":                offset,
		"Limit":                 limit,
	})
}

// Unified URI scheme handlers - per-user.
func (a *API) handleUnifiedEthAllByUser(c *httpx.Context) {
	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	if !a.dbInitialized() {
		common.RespondErrorJSON(c, "Database link wasn't configured")
		return
	}
	pUserAddr := c.Param("user_addr")
	userAid, err := a.store.LookupAddressID(c.Request.Context(), pUserAddr)
	if err != nil {
		common.RespondErrorJSON(c, "Provided address wasn't found")
		return
	}

	deposits, err := a.repo.EthDepositsByUser(c.Request.Context(), userAid)
	if err != nil {
		a.respondStoreError(c, err)
		return
	}

	reqStatus := 1
	errStr := ""
	c.JSON(http.StatusOK, httpx.H{
		"status":      reqStatus,
		"error":       errStr,
		"UserAddr":    pUserAddr,
		"UserAid":     userAid,
		"AllDeposits": deposits,
	})
}

func (a *API) handleUnifiedEthRaffleByUser(c *httpx.Context) {
	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	if !a.dbInitialized() {
		common.RespondErrorJSON(c, "Database link wasn't configured")
		return
	}
	pUserAddr := c.Param("user_addr")
	userAid, err := a.store.LookupAddressID(c.Request.Context(), pUserAddr)
	if err != nil {
		common.RespondErrorJSON(c, "Provided address wasn't found")
		return
	}

	deposits, err := a.repo.RaffleEthDepositsByUser(c.Request.Context(), userAid)
	if err != nil {
		a.respondStoreError(c, err)
		return
	}

	reqStatus := 1
	errStr := ""
	c.JSON(http.StatusOK, httpx.H{
		"status":         reqStatus,
		"error":          errStr,
		"UserAddr":       pUserAddr,
		"UserAid":        userAid,
		"RaffleDeposits": deposits,
	})
}

func (a *API) handleUnifiedEthChronowarriorByUser(c *httpx.Context) {
	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	if !a.dbInitialized() {
		common.RespondErrorJSON(c, "Database link wasn't configured")
		return
	}
	pUserAddr := c.Param("user_addr")
	userAid, err := a.store.LookupAddressID(c.Request.Context(), pUserAddr)
	if err != nil {
		common.RespondErrorJSON(c, "Provided address wasn't found")
		return
	}

	deposits, err := a.repo.ChronoWarriorEthDepositsByUser(c.Request.Context(), userAid)
	if err != nil {
		a.respondStoreError(c, err)
		return
	}

	reqStatus := 1
	errStr := ""
	c.JSON(http.StatusOK, httpx.H{
		"status":                reqStatus,
		"error":                 errStr,
		"UserAddr":              pUserAddr,
		"UserAid":               userAid,
		"ChronoWarriorDeposits": deposits,
	})
}

func (a *API) handlePrizeDepositsByRound(c *httpx.Context) {
	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	if !a.dbInitialized() {
		common.RespondErrorJSON(c, "Database link wasn't configured")
		return
	}

	pRoundNum := c.Param("round_num")
	var roundNum int64
	if len(pRoundNum) > 0 {
		var success bool
		roundNum, success = common.ParseIntFromRemoteOrError(c, JSON, &pRoundNum)
		if !success {
			return
		}
	}

	deposits, err := a.repo.PrizeDepositsByRound(c.Request.Context(), roundNum)
	if err != nil {
		a.respondStoreError(c, err)
		return
	}

	reqStatus := 1
	errStr := ""
	c.JSON(http.StatusOK, httpx.H{
		"status":         reqStatus,
		"error":          errStr,
		"RaffleDeposits": deposits,
		"RoundNum":       roundNum,
	})
}

func (a *API) handleRaffleNftWinnersList(c *httpx.Context) {
	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	if !a.dbInitialized() {
		common.RespondErrorJSON(c, "Database link wasn't configured")
		return
	}

	var offset, limit int
	pOffset := c.Param("offset")
	pLimit := c.Param("limit")

	if len(pOffset) == 0 || len(pLimit) == 0 {
		offset = 0
		limit = 10000
	} else {
		var success bool
		success, offset, limit = common.ParseOffsetLimitParamsJSON(c)
		if !success {
			return
		}
	}

	winners, err := a.repo.RaffleNFTWinners(c.Request.Context(), offset, limit)
	if err != nil {
		a.respondStoreError(c, err)
		return
	}

	reqStatus := 1
	errStr := ""
	c.JSON(http.StatusOK, httpx.H{
		"status":           reqStatus,
		"error":            errStr,
		"RaffleNFTWinners": winners,
		"Offset":           offset,
		"Limit":            limit,
	})
}

func (a *API) handleRaffleNftWinnersByRound(c *httpx.Context) {
	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	if !a.dbInitialized() {
		common.RespondErrorJSON(c, "Database link wasn't configured")
		return
	}

	pRoundNum := c.Param("round_num")
	var roundNum int64
	if len(pRoundNum) > 0 {
		var success bool
		roundNum, success = common.ParseIntFromRemoteOrError(c, JSON, &pRoundNum)
		if !success {
			return
		}
	}

	winnersRaffle, err := a.repo.RaffleNFTWinnersByRound(c.Request.Context(), roundNum, false)
	if err != nil {
		a.respondStoreError(c, err)
		return
	}
	winnersStaking, err := a.repo.RaffleNFTWinnersByRound(c.Request.Context(), roundNum, true)
	if err != nil {
		a.respondStoreError(c, err)
		return
	}

	reqStatus := 1
	errStr := ""
	c.JSON(http.StatusOK, httpx.H{
		"status":            reqStatus,
		"error":             errStr,
		"RaffleNFTWinners":  winnersRaffle,
		"StakingNFTWinners": winnersStaking,
		"RoundNum":          roundNum,
	})
}

func (a *API) handleUserRaffleNftWinnings(c *httpx.Context) {
	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	if !a.dbInitialized() {
		common.RespondErrorJSON(c, "Database link wasn't configured")
		return
	}

	pUserAddr := c.Param("user_addr")
	userAid, err := a.store.LookupAddressID(c.Request.Context(), pUserAddr)
	if err != nil {
		common.RespondErrorJSON(c, "Provided address wasn't found")
		return
	}

	userInfo, err := a.repo.UserInfo(c.Request.Context(), userAid)
	if errors.Is(err, store.ErrNotFound) {
		common.RespondErrorJSON(c, "Provided address wasn't found")
		return
	}
	if err != nil {
		a.respondStoreError(c, err)
		return
	}

	winnings, err := a.repo.RaffleNFTWinningsByUser(c.Request.Context(), userAid)
	if err != nil {
		a.respondStoreError(c, err)
		return
	}

	reqStatus := 1
	errStr := ""
	c.JSON(http.StatusOK, httpx.H{
		"status":                reqStatus,
		"error":                 errStr,
		"UserRaffleNFTWinnings": winnings,
		"UserInfo":              userInfo,
	})
}

func (a *API) handlePrizeDepositsRaffleEthByUser(c *httpx.Context) {
	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	if !a.dbInitialized() {
		common.RespondErrorJSON(c, "Database link wasn't configured")
		return
	}

	pUserAddr := c.Param("user_addr")
	userAid, err := a.store.LookupAddressID(c.Request.Context(), pUserAddr)
	if err != nil {
		common.RespondErrorJSON(c, "Provided address wasn't found")
		return
	}
	userInfo, err := a.repo.UserInfo(c.Request.Context(), userAid)
	if errors.Is(err, store.ErrNotFound) {
		common.RespondErrorJSON(c, "Provided address wasn't found")
		return
	}
	if err != nil {
		a.respondStoreError(c, err)
		return
	}

	deposits, err := a.repo.PrizeDepositsRaffleEthByUser(c.Request.Context(), userAid)
	if err != nil {
		a.respondStoreError(c, err)
		return
	}
	reqStatus := 1
	errStr := ""
	c.JSON(http.StatusOK, httpx.H{
		"status":             reqStatus,
		"error":              errStr,
		"UserRaffleDeposits": deposits,
		"UserInfo":           userInfo,
	})
}

func (a *API) handlePrizeDepositsChronoWarriorByUser(c *httpx.Context) {
	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	if !a.dbInitialized() {
		common.RespondErrorJSON(c, "Database link wasn't configured")
		return
	}

	pUserAddr := c.Param("user_addr")
	userAid, err := a.store.LookupAddressID(c.Request.Context(), pUserAddr)
	if err != nil {
		common.RespondErrorJSON(c, "Provided address wasn't found")
		return
	}
	userInfo, err := a.repo.UserInfo(c.Request.Context(), userAid)
	if errors.Is(err, store.ErrNotFound) {
		common.RespondErrorJSON(c, "Provided address wasn't found")
		return
	}
	if err != nil {
		a.respondStoreError(c, err)
		return
	}

	deposits, err := a.repo.PrizeDepositsChronoWarriorByUser(c.Request.Context(), userAid)
	if err != nil {
		a.respondStoreError(c, err)
		return
	}
	reqStatus := 1
	errStr := ""
	c.JSON(http.StatusOK, httpx.H{
		"status":                    reqStatus,
		"error":                     errStr,
		"UserChronoWarriorDeposits": deposits,
		"UserInfo":                  userInfo,
	})
}

func (a *API) handleNftDonationsByPrize(c *httpx.Context) {
	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	if !a.dbInitialized() {
		common.RespondErrorJSON(c, "Database link wasn't configured")
		return
	}

	pPrizeNum := c.Param("prize_num")
	var prizeNum int64
	if len(pPrizeNum) > 0 {
		var success bool
		prizeNum, success = common.ParseIntFromRemoteOrError(c, JSON, &pPrizeNum)
		if !success {
			return
		}
	}
	nftDonations, err := a.repo.NFTDonationsByRound(c.Request.Context(), prizeNum)
	if err != nil {
		a.respondStoreError(c, err)
		return
	}
	reqStatus := 1
	errStr := ""
	c.JSON(http.StatusOK, httpx.H{
		"status":       reqStatus,
		"error":        errStr,
		"NFTDonations": nftDonations,
		"RoundNum":     prizeNum,
	})
}

func (a *API) handleNftDonationsByToken(c *httpx.Context) {
	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	if !a.dbInitialized() {
		common.RespondErrorJSON(c, "Database link wasn't configured")
		return
	}
	pTokenAddr := c.Param("token_addr")
	tokenAid, err := a.store.LookupAddressID(c.Request.Context(), pTokenAddr)
	if err != nil {
		common.RespondErrorJSON(c, "Token address not found")
		return
	}
	nftDonations, err := a.repo.NFTDonationsByToken(c.Request.Context(), tokenAid)
	if err != nil {
		a.respondStoreError(c, err)
		return
	}
	c.JSON(http.StatusOK, httpx.H{
		"status":       1,
		"error":        "",
		"NFTDonations": nftDonations,
		"TokenAddr":    pTokenAddr,
	})
}

func (a *API) handleCosmicSignatureTokenList(c *httpx.Context) {
	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	if !a.dbInitialized() {
		common.RespondErrorJSON(c, "Database link wasn't configured")
		return
	}

	success, offset, limit := common.ParseOffsetLimitParamsJSON(c)
	if !success {
		return
	}
	tokens, err := a.repo.CosmicSignatureTokens(c.Request.Context(), offset, limit)
	if err != nil {
		a.respondStoreError(c, err)
		return
	}

	reqStatus := 1
	errStr := ""
	c.JSON(http.StatusOK, httpx.H{
		"status":                   reqStatus,
		"error":                    errStr,
		"CosmicSignatureTokenList": tokens,
		"Offset":                   offset,
		"Limit":                    limit,
	})
}

func (a *API) handleCosmicSignatureTokenInfo(c *httpx.Context) {
	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	if !a.dbInitialized() {
		common.RespondErrorJSON(c, "Database link wasn't configured")
		return
	}

	pTokenID := c.Param("token_id")
	var tokenID int64
	if len(pTokenID) > 0 {
		var success bool
		tokenID, success = common.ParseIntFromRemoteOrError(c, JSON, &pTokenID)
		if !success {
			return
		}
	}

	tokenInfo, err := a.repo.CosmicSignatureTokenInfo(c.Request.Context(), tokenID)
	if err != nil {
		if errors.Is(err, store.ErrNotFound) {
			common.RespondErrorJSON(c, "record not found")
			return
		}
		a.respondStoreError(c, err)
		return
	}

	reqStatus := 1
	errStr := ""

	if tokenInfo.RecordType == 3 {
		// The legacy code ignored the found/not-found flag and rendered the
		// zero-value record when the round has no prize claim; ErrNotFound
		// keeps that exact behavior.
		prizeInfo, err := a.repo.PrizeInfo(c.Request.Context(), tokenInfo.RoundNum)
		if err != nil && !errors.Is(err, store.ErrNotFound) {
			a.respondStoreError(c, err)
			return
		}
		c.JSON(http.StatusOK, httpx.H{
			"status":    reqStatus,
			"error":     errStr,
			"TokenInfo": tokenInfo,
			"PrizeInfo": prizeInfo,
		})
	} else {
		c.JSON(http.StatusOK, httpx.H{
			"status":    reqStatus,
			"error":     errStr,
			"TokenInfo": tokenInfo,
		})
	}
}

func (a *API) handleDonatedNftClaimsAll(c *httpx.Context) {
	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	if !a.dbInitialized() {
		common.RespondErrorJSON(c, "Database link wasn't configured")
		return
	}

	var offset, limit int
	pOffset := c.Param("offset")
	pLimit := c.Param("limit")

	if len(pOffset) == 0 || len(pLimit) == 0 {
		offset = 0
		limit = 10000
	} else {
		var success bool
		success, offset, limit = common.ParseOffsetLimitParamsJSON(c)
		if !success {
			return
		}
	}

	claims, err := a.repo.DonatedNFTClaims(c.Request.Context(), offset, limit)
	if err != nil {
		a.respondStoreError(c, err)
		return
	}

	reqStatus := 1
	errStr := ""
	c.JSON(http.StatusOK, httpx.H{
		"status":           reqStatus,
		"error":            errStr,
		"DonatedNFTClaims": claims,
		"Offset":           offset,
		"Limit":            limit,
	})
}

func (a *API) handleDonatedNftClaimsByUser(c *httpx.Context) {
	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	if !a.dbInitialized() {
		common.RespondErrorJSON(c, "Database link wasn't configured")
		return
	}
	pUserAddr := c.Param("user_addr")
	userAid, err := a.store.LookupAddressID(c.Request.Context(), pUserAddr)
	if err != nil {
		c.JSON(http.StatusOK, httpx.H{
			"status": 1, "error": "", "DonatedNFTClaims": []any{},
			"UserInfo": cgmodel.CGUserInfo{Address: pUserAddr},
		})
		return
	}
	userInfo, err := a.repo.UserInfo(c.Request.Context(), userAid)
	if errors.Is(err, store.ErrNotFound) {
		c.JSON(http.StatusOK, httpx.H{
			"status": 1, "error": "", "DonatedNFTClaims": []any{},
			"UserInfo": cgmodel.CGUserInfo{Address: pUserAddr},
		})
		return
	}
	if err != nil {
		a.respondStoreError(c, err)
		return
	}
	claims, err := a.repo.DonatedNFTClaimsByUser(c.Request.Context(), userAid)
	if err != nil {
		a.respondStoreError(c, err)
		return
	}
	reqStatus := 1
	errStr := ""
	c.JSON(http.StatusOK, httpx.H{
		"status":           reqStatus,
		"error":            errStr,
		"DonatedNFTClaims": claims,
		"UserInfo":         userInfo,
	})
}

func (a *API) handleTimeCurrent(c *httpx.Context) {
	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	if !a.dbInitialized() {
		common.RespondErrorJSON(c, "Database link wasn't configured")
		return
	}
	var raw json.RawMessage
	err := a.rpc.CallContext(context.Background(), &raw, "eth_getBlockByNumber", "pending", true)
	if err != nil {
		common.RespondErrorJSON(c, fmt.Sprintf("%v", err))
		return
	}
	rpcobj := make(map[string]any)
	err = json.Unmarshal(raw, &rpcobj)
	if err != nil {
		common.RespondErrorJSON(c, fmt.Sprintf("Error decoding JSON: %v", err))
		return
	}

	tsHex := rpcobj["timestamp"].(string)
	ts, err := hexutil.DecodeUint64(tsHex)
	if err != nil {
		common.RespondErrorJSON(c, fmt.Sprintf("Error decoding timestamp from hex: %v", err))
		return
	}
	reqStatus := 1
	errStr := ""
	c.JSON(http.StatusOK, httpx.H{
		"status":           reqStatus,
		"error":            errStr,
		"CurrentTimeStamp": ts,
	})
}

func (a *API) handleTimeUntilPrize(c *httpx.Context) {
	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	if !a.dbInitialized() {
		common.RespondErrorJSON(c, "Database link wasn't configured")
		return
	}
	// Function: getDurationUntilMainPrize() - returns time until prize can be claimed
	const timeUntilPrizeSig string = "0x36750d2c"
	var raw json.RawMessage
	args := map[string]any{
		"to":   a.state.Snapshot().Addrs.CosmicGame.String(),
		"data": timeUntilPrizeSig,
	}
	err := a.rpc.CallContext(context.Background(), &raw, "eth_call", args, "pending")
	if err != nil {
		common.RespondErrorJSON(c, fmt.Sprintf("%v", err))
		return
	}
	var tsHex string
	err = json.Unmarshal(raw, &tsHex)
	if err != nil {
		common.RespondErrorJSON(c, fmt.Sprintf("Error decoding JSON: %v", err))
		return
	}
	tsBig := ethcommon.HexToHash(tsHex).Big()
	reqStatus := 1
	errStr := ""
	c.JSON(http.StatusOK, httpx.H{
		"status":         reqStatus,
		"error":          errStr,
		"TimeUntilPrize": tsBig.Int64(),
	})
}

func (a *API) handlePrizeCurRoundTime(c *httpx.Context) {
	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	if !a.dbInitialized() {
		common.RespondErrorJSON(c, "Database link wasn't configured")
		return
	}

	// Return 200 with status/error in body when contract is unavailable, so frontend gets consistent response shape
	reqStatus := 1
	var errStr string
	var prizeTime *big.Int

	// abigen constructors only fail parsing their embedded ABI constant, so
	// the error is statically impossible; a constructed module always has an
	// Ethereum client (contractstate requires one).
	var copts bind.CallOpts
	bwcontract, _ := cg.NewCosmicSignatureGame(a.state.Snapshot().Addrs.CosmicGame, a.ethClient)
	pt, err := bwcontract.MainPrizeTime(&copts)
	if err != nil {
		reqStatus = 0
		errStr = fmt.Sprintf("MainPrizeTime call failed: %v", err)
		prizeTime = big.NewInt(0)
	} else {
		prizeTime = pt
	}

	c.JSON(http.StatusOK, httpx.H{
		"status":            reqStatus,
		"error":             errStr,
		"CurRoundPrizeTime": prizeTime,
	})
}

func (a *API) handleUserGlobalWinnings(c *httpx.Context) {
	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	if !a.dbInitialized() {
		common.RespondErrorJSON(c, "Database link wasn't configured")
		return
	}
	pUserAddr := c.Param("user_addr")
	userAid, err := a.store.LookupAddressID(c.Request.Context(), pUserAddr)
	if err != nil {
		// Address not in DB yet — return 200 with empty winnings so UI works
		c.JSON(http.StatusOK, httpx.H{
			"status": 1, "error": "", "Winnings": []any{}, "UserAddr": pUserAddr, "UserAid": int64(0),
		})
		return
	}

	claimInfo, err := a.repo.UserNotifRedBoxRewards(c.Request.Context(), userAid)
	if err != nil {
		a.respondStoreError(c, err)
		return
	}

	reqStatus := 1
	errStr := ""
	c.JSON(http.StatusOK, httpx.H{
		"status":   reqStatus,
		"error":    errStr,
		"Winnings": claimInfo,
		"UserAddr": pUserAddr,
		"UserAid":  userAid,
	})
}

func (a *API) handlePrizeHistoryDetailByUser(c *httpx.Context) {
	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	if !a.dbInitialized() {
		common.RespondErrorJSON(c, "Database link wasn't configured")
		return
	}
	pUserAddr := c.Param("user_addr")
	success, offset, limit := common.ParseOffsetLimitParamsJSON(c)
	if !success {
		return
	}
	userAid, err := a.store.LookupAddressID(c.Request.Context(), pUserAddr)
	if err != nil {
		c.JSON(http.StatusOK, httpx.H{
			"status": 1, "error": "", "UserAddr": pUserAddr, "UserAid": int64(0),
			"UserPrizeHistory": []any{},
		})
		return
	}
	_, err = a.repo.UserInfo(c.Request.Context(), userAid)
	if errors.Is(err, store.ErrNotFound) {
		c.JSON(http.StatusOK, httpx.H{
			"status": 1, "error": "", "UserAddr": pUserAddr, "UserAid": userAid,
			"UserPrizeHistory": []any{},
		})
		return
	}
	if err != nil {
		a.respondStoreError(c, err)
		return
	}

	claimHistory, err := a.repo.PrizeHistoryByUser(c.Request.Context(), userAid, offset, limit)
	if err != nil {
		a.respondStoreError(c, err)
		return
	}

	reqStatus := 1
	errStr := ""
	c.JSON(http.StatusOK, httpx.H{
		"status":           reqStatus,
		"error":            errStr,
		"UserAddr":         pUserAddr,
		"UserAid":          userAid,
		"UserPrizeHistory": claimHistory,
	})
}

func (a *API) handleGlobalClaimHistoryDetail(c *httpx.Context) {
	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	if !a.dbInitialized() {
		common.RespondErrorJSON(c, "Database link wasn't configured")
		return
	}

	success, offset, limit := common.ParseOffsetLimitParamsJSON(c)
	if !success {
		return
	}

	claimHistory, err := a.repo.ClaimHistoryGlobal(c.Request.Context(), offset, limit)
	if err != nil {
		a.respondStoreError(c, err)
		return
	}
	reqStatus := 1
	errStr := ""
	c.JSON(http.StatusOK, httpx.H{
		"status":             reqStatus,
		"error":              errStr,
		"GlobalPrizeHistory": claimHistory,
	})
}

func (a *API) handleUnclaimedDonatedNftsByUser(c *httpx.Context) {
	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	if !a.dbInitialized() {
		common.RespondErrorJSON(c, "Database link wasn't configured")
		return
	}
	pUserAddr := c.Param("user_addr")
	userAid, err := a.store.LookupAddressID(c.Request.Context(), pUserAddr)
	if err != nil {
		c.JSON(http.StatusOK, httpx.H{
			"status": 1, "error": "", "UnclaimedDonatedNFTs": []any{},
			"UserAddr": pUserAddr, "UserAid": int64(0),
		})
		return
	}
	_, err = a.repo.UserInfo(c.Request.Context(), userAid)
	if errors.Is(err, store.ErrNotFound) {
		c.JSON(http.StatusOK, httpx.H{
			"status": 1, "error": "", "UnclaimedDonatedNFTs": []any{},
			"UserAddr": pUserAddr, "UserAid": userAid,
		})
		return
	}
	if err != nil {
		a.respondStoreError(c, err)
		return
	}

	nfts, err := a.repo.UnclaimedDonatedNFTsByUser(c.Request.Context(), userAid)
	if err != nil {
		a.respondStoreError(c, err)
		return
	}

	reqStatus := 1
	errStr := ""
	c.JSON(http.StatusOK, httpx.H{
		"status":               reqStatus,
		"error":                errStr,
		"UnclaimedDonatedNFTs": nfts,
		"UserAddr":             pUserAddr,
		"UserAid":              userAid,
	})
}

func (a *API) handleUnclaimedDonatedNftsByPrize(c *httpx.Context) {
	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	if !a.dbInitialized() {
		common.RespondErrorJSON(c, "Database link wasn't configured")
		return
	}

	pPrizeNum := c.Param("prize_num")
	var prizeNum int64
	if len(pPrizeNum) > 0 {
		var success bool
		prizeNum, success = common.ParseIntFromRemoteOrError(c, JSON, &pPrizeNum)
		if !success {
			return
		}
	}

	nftDonations, err := a.repo.UnclaimedDonatedNFTsByRound(c.Request.Context(), prizeNum)
	if err != nil {
		a.respondStoreError(c, err)
		return
	}
	reqStatus := 1
	errStr := ""
	c.JSON(http.StatusOK, httpx.H{
		"status":       reqStatus,
		"error":        errStr,
		"NFTDonations": nftDonations,
		"RoundNum":     prizeNum,
	})
}

func (a *API) handleUnclaimedPrizeDepositsByUser(c *httpx.Context) {
	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	if !a.dbInitialized() {
		common.RespondErrorJSON(c, "Database link wasn't configured")
		return
	}
	pUserAddr := c.Param("user_addr")
	success, offset, limit := common.ParseOffsetLimitParamsJSON(c)
	if !success {
		return
	}
	userAid, err := a.store.LookupAddressID(c.Request.Context(), pUserAddr)
	if err != nil {
		c.JSON(http.StatusOK, httpx.H{
			"status": 1, "error": "", "UserAddr": pUserAddr, "UserAid": int64(0),
			"UnclaimedDeposits": []any{},
		})
		return
	}
	_, err = a.repo.UserInfo(c.Request.Context(), userAid)
	if errors.Is(err, store.ErrNotFound) {
		c.JSON(http.StatusOK, httpx.H{
			"status": 1, "error": "", "UserAddr": pUserAddr, "UserAid": userAid,
			"UnclaimedDeposits": []any{},
		})
		return
	}
	if err != nil {
		a.respondStoreError(c, err)
		return
	}

	deposits, err := a.repo.UnclaimedPrizeEthDeposits(c.Request.Context(), userAid, offset, limit)
	if err != nil {
		a.respondStoreError(c, err)
		return
	}

	reqStatus := 1
	errStr := ""
	c.JSON(http.StatusOK, httpx.H{
		"status":            reqStatus,
		"error":             errStr,
		"UserAddr":          pUserAddr,
		"UserAid":           userAid,
		"UnclaimedDeposits": deposits,
	})
}

func (a *API) handleCosmicSignatureTokenListByUser(c *httpx.Context) {
	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	if !a.dbInitialized() {
		common.RespondErrorJSON(c, "Database link wasn't configured")
		return
	}
	pUserAddr := c.Param("user_addr")
	userAid, err := a.store.LookupAddressID(c.Request.Context(), pUserAddr)
	if err != nil {
		// Address not in DB yet — return 200 with empty list so UI works
		success, offset, limit := common.ParseOffsetLimitParamsJSON(c)
		if !success {
			return
		}
		c.JSON(http.StatusOK, httpx.H{
			"status": 1, "error": "", "UserAddr": pUserAddr, "UserAid": int64(0),
			"UserTokens": []any{}, "Offset": offset, "Limit": limit,
		})
		return
	}
	_, err = a.repo.UserInfo(c.Request.Context(), userAid)
	if errors.Is(err, store.ErrNotFound) {
		success, offset, limit := common.ParseOffsetLimitParamsJSON(c)
		if !success {
			return
		}
		c.JSON(http.StatusOK, httpx.H{
			"status": 1, "error": "", "UserAddr": pUserAddr, "UserAid": userAid,
			"UserTokens": []any{}, "Offset": offset, "Limit": limit,
		})
		return
	}
	if err != nil {
		a.respondStoreError(c, err)
		return
	}
	success, offset, limit := common.ParseOffsetLimitParamsJSON(c)
	if !success {
		return
	}

	userTokens, err := a.repo.CosmicSignatureTokensByUser(c.Request.Context(), userAid, offset, limit)
	if err != nil {
		a.respondStoreError(c, err)
		return
	}

	reqStatus := 1
	errStr := ""
	c.JSON(http.StatusOK, httpx.H{
		"status":     reqStatus,
		"error":      errStr,
		"UserAddr":   pUserAddr,
		"UserAid":    userAid,
		"UserTokens": userTokens,
	})
}

func (a *API) handleTokenNameHistory(c *httpx.Context) {
	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	if !a.dbInitialized() {
		common.RespondErrorJSON(c, "Database link wasn't configured")
		return
	}

	pTokenID := c.Param("token_id")
	var tokenID int64
	if len(pTokenID) > 0 {
		var success bool
		tokenID, success = common.ParseIntFromRemoteOrError(c, JSON, &pTokenID)
		if !success {
			return
		}
	}

	reqStatus := 1
	errStr := ""

	toknameHistory, err := a.repo.TokenNameHistory(c.Request.Context(), tokenID)
	if err != nil {
		a.respondStoreError(c, err)
		return
	}
	c.JSON(http.StatusOK, httpx.H{
		"status":           reqStatus,
		"error":            errStr,
		"TokenId":          tokenID,
		"TokenNameHistory": toknameHistory,
	})
}

func (a *API) handleTokenNameSearch(c *httpx.Context) {
	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	if !a.dbInitialized() {
		common.RespondErrorJSON(c, "Database link wasn't configured")
		return
	}

	pName := c.Param("name")

	results, err := a.repo.SearchTokensByName(c.Request.Context(), pName)
	if err != nil {
		a.respondStoreError(c, err)
		return
	}
	reqStatus := 1
	errStr := ""

	c.JSON(http.StatusOK, httpx.H{
		"status":                 reqStatus,
		"error":                  errStr,
		"SearchText":             pName,
		"TokenNameSearchResults": results,
	})
}

func (a *API) handleNamedTokensOnly(c *httpx.Context) {
	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	if !a.dbInitialized() {
		common.RespondErrorJSON(c, "Database link wasn't configured")
		return
	}

	reqStatus := 1
	errStr := ""

	results, err := a.repo.NamedTokens(c.Request.Context())
	if err != nil {
		a.respondStoreError(c, err)
		return
	}
	c.JSON(http.StatusOK, httpx.H{
		"status":      reqStatus,
		"error":       errStr,
		"NamedTokens": results,
	})
}

func (a *API) handleTokenOwnershipTransfers(c *httpx.Context) {
	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	if !a.dbInitialized() {
		common.RespondErrorJSON(c, "Database link wasn't configured")
		return
	}

	pTokenID := c.Param("token_id")
	var tokenID int64
	if len(pTokenID) > 0 {
		var success bool
		tokenID, success = common.ParseIntFromRemoteOrError(c, JSON, &pTokenID)
		if !success {
			return
		}
	}
	success, offset, limit := common.ParseOffsetLimitParamsJSON(c)
	if !success {
		return
	}

	reqStatus := 1
	errStr := ""

	transfers, err := a.repo.TokenOwnershipTransfers(c.Request.Context(), tokenID, offset, limit)
	if err != nil {
		a.respondStoreError(c, err)
		return
	}
	c.JSON(http.StatusOK, httpx.H{
		"status":         reqStatus,
		"error":          errStr,
		"Offset":         offset,
		"Limit":          limit,
		"TokenId":        tokenID,
		"TokenTransfers": transfers,
	})
}

func (a *API) handleCsTokenDistribution(c *httpx.Context) {
	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")

	if !a.dbInitialized() {
		common.RespondErrorJSON(c, "Database link wasn't configured")
		return
	}
	distribution, err := a.repo.CosmicSignatureTokenDistribution(c.Request.Context())
	if err != nil {
		a.respondStoreError(c, err)
		return
	}

	reqStatus := 1
	errStr := ""

	c.JSON(http.StatusOK, httpx.H{
		"status":                           reqStatus,
		"error":                            errStr,
		"CosmicSignatureTokenDistribution": distribution,
	})
}

func (a *API) handleUserBalances(c *httpx.Context) {
	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	if !a.dbInitialized() {
		common.RespondErrorJSON(c, "Database link wasn't configured")
		return
	}

	pUserAddr := c.Param("user_addr")
	userAid, addrLookupErr := a.store.LookupAddressID(c.Request.Context(), pUserAddr)
	if addrLookupErr != nil {
		userAid = 0 // address not in DB yet; still return 200 with on-chain balances below
	}

	addr := ethcommon.HexToAddress(pUserAddr)
	userEthBal, err := a.ethClient.BalanceAt(context.Background(), addr, nil)
	if err != nil {
		errStr := fmt.Sprintf("Error at BalanceAt() call for addr: %v\n", err)
		common.RespondErrorJSON(c, errStr)
		return
	}
	// The abigen constructor only fails parsing its embedded ABI constant —
	// statically impossible.
	ctContract, _ := cg.NewERC20(a.state.Snapshot().Addrs.CosmicToken, a.ethClient)
	var copts bind.CallOpts
	ctBalance, err := ctContract.BalanceOf(&copts, addr)
	if err != nil {
		errStr := fmt.Sprintf("Error at BalanceOf() call: %v\n", err)
		common.RespondErrorJSON(c, errStr)
		return
	}

	reqStatus := 1
	errStr := ""

	c.JSON(http.StatusOK, httpx.H{
		"status":             reqStatus,
		"error":              errStr,
		"UserAddr":           pUserAddr,
		"UserAid":            userAid,
		"ETH_Balance":        userEthBal.String(),
		"CosmicTokenBalance": ctBalance.String(),
	})
}

func (a *API) handleCosmicTokenBalances(c *httpx.Context) {
	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	if !a.dbInitialized() {
		common.RespondErrorJSON(c, "Database link wasn't configured")
		return
	}

	balances, err := a.repo.CosmicTokenHolders(c.Request.Context())
	if err != nil {
		a.respondStoreError(c, err)
		return
	}
	reqStatus := 1
	errStr := ""

	c.JSON(http.StatusOK, httpx.H{
		"status":              reqStatus,
		"error":               errStr,
		"CosmicTokenBalances": balances,
	})
}

func (a *API) handleCosmicTokenStatistics(c *httpx.Context) {
	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	if !a.dbInitialized() {
		common.RespondErrorJSON(c, "Database link wasn't configured")
		return
	}

	stats, err := a.repo.CosmicTokenStatistics(c.Request.Context())
	if err != nil {
		a.respondStoreError(c, err)
		return
	}

	cosmicTokenAddr := a.state.Snapshot().Addrs.CosmicToken

	// Read contract info from blockchain. The abigen constructor only fails
	// parsing its embedded ABI constant — statically impossible.
	var copts bind.CallOpts
	contract, _ := cg.NewCosmicSignatureToken(cosmicTokenAddr, a.ethClient)

	tokenName, err := contract.Name(&copts)
	if err != nil {
		errStr := fmt.Sprintf("Error reading token name: %v", err)
		a.logger.Error(errStr)
		common.RespondErrorJSON(c, errStr)
		return
	}

	tokenSymbol, err := contract.Symbol(&copts)
	if err != nil {
		errStr := fmt.Sprintf("Error reading token symbol: %v", err)
		a.logger.Error(errStr)
		common.RespondErrorJSON(c, errStr)
		return
	}

	decimals, err := contract.Decimals(&copts)
	if err != nil {
		errStr := fmt.Sprintf("Error reading decimals: %v", err)
		a.logger.Error(errStr)
		common.RespondErrorJSON(c, errStr)
		return
	}

	gameAddr, err := contract.Game(&copts)
	if err != nil {
		errStr := fmt.Sprintf("Error reading game address: %v", err)
		a.logger.Error(errStr)
		common.RespondErrorJSON(c, errStr)
		return
	}

	reqStatus := 1
	errStr := ""
	c.JSON(http.StatusOK, httpx.H{
		"status":           reqStatus,
		"error":            errStr,
		"TokenName":        tokenName,
		"TokenSymbol":      tokenSymbol,
		"TokenDecimals":    decimals,
		"GameContractAddr": gameAddr.String(),
		"CosmicTokenAddr":  cosmicTokenAddr.String(),
		"Statistics":       stats,
	})
}

func (a *API) handleCosmicTokenSummaryByUser(c *httpx.Context) {
	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	if !a.dbInitialized() {
		common.RespondErrorJSON(c, "Database link wasn't configured")
		return
	}
	pUserAddr := c.Param("user_addr")
	userAid, err := a.store.LookupAddressID(c.Request.Context(), pUserAddr)
	if err != nil {
		common.RespondErrorJSON(c, "Provided address wasn't found")
		return
	}

	summary, err := a.repo.UserCosmicTokenSummary(c.Request.Context(), userAid)
	if err != nil {
		a.respondStoreError(c, err)
		return
	}
	summary.UserAddr = pUserAddr

	reqStatus := 1
	errStr := ""
	c.JSON(http.StatusOK, httpx.H{
		"status":  reqStatus,
		"error":   errStr,
		"Summary": summary,
	})
}

func parseCosmicTokenHistoryDateParam(c *httpx.Context, paramName string) (string, bool) {
	val := c.Param(paramName)
	if len(val) != 8 {
		common.RespondErrorJSON(c, fmt.Sprintf("'%s' must be YYYYMMDD (8 digits)", paramName))
		return "", false
	}
	if _, err := time.Parse("20060102", val); err != nil {
		common.RespondErrorJSON(c, fmt.Sprintf("Bad '%s' parameter: %v", paramName, err))
		return "", false
	}
	return val, true
}

func (a *API) handleCosmicTokenTotalSupplyHistoryByBid(c *httpx.Context) {
	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	if !a.dbInitialized() {
		common.RespondErrorJSON(c, "Database link wasn't configured")
		return
	}

	history, err := a.repo.CosmicTokenSupplyHistoryByBid(c.Request.Context())
	if err != nil {
		a.respondStoreError(c, err)
		return
	}
	c.JSON(http.StatusOK, httpx.H{
		"status":             1,
		"error":              "",
		"TotalSupplyHistory": history,
	})
}

func (a *API) handleCosmicTokenTotalSupplyHistoryByDate(c *httpx.Context) {
	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	if !a.dbInitialized() {
		common.RespondErrorJSON(c, "Database link wasn't configured")
		return
	}

	fromDate, ok := parseCosmicTokenHistoryDateParam(c, "from_date")
	if !ok {
		return
	}
	toDate, ok := parseCosmicTokenHistoryDateParam(c, "to_date")
	if !ok {
		return
	}
	if fromDate > toDate {
		common.RespondErrorJSON(c, "'from_date' must be on or before 'to_date'")
		return
	}

	history, err := a.repo.CosmicTokenSupplyHistoryByDate(c.Request.Context(), fromDate, toDate)
	if err != nil {
		a.respondStoreError(c, err)
		return
	}
	c.JSON(http.StatusOK, httpx.H{
		"status":             1,
		"error":              "",
		"DateFrom":           fromDate,
		"DateTo":             toDate,
		"TotalSupplyHistory": history,
	})
}

func (a *API) handleCosmicTokenTransfersByUser(c *httpx.Context) {
	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	if !a.dbInitialized() {
		common.RespondErrorJSON(c, "Database link wasn't configured")
		return
	}
	pUserAddr := c.Param("user_addr")
	userAid, err := a.store.LookupAddressID(c.Request.Context(), pUserAddr)
	if err != nil {
		common.RespondErrorJSON(c, "Provided address wasn't found")
		return
	}
	success, offset, limit := common.ParseOffsetLimitParamsJSON(c)
	if !success {
		return
	}
	transfers, err := a.repo.CosmicTokenTransfersByUser(c.Request.Context(), userAid, offset, limit)
	if err != nil {
		a.respondStoreError(c, err)
		return
	}
	reqStatus := 1
	errStr := ""

	c.JSON(http.StatusOK, httpx.H{
		"status":               reqStatus,
		"error":                errStr,
		"UserAddr":             pUserAddr,
		"UserAid":              userAid,
		"Offset":               offset,
		"Limit":                limit,
		"CosmicTokenTransfers": transfers,
	})
}

func (a *API) handleCosmicSignatureTransfersByUser(c *httpx.Context) {
	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	if !a.dbInitialized() {
		common.RespondErrorJSON(c, "Database link wasn't configured")
		return
	}
	pUserAddr := c.Param("user_addr")
	userAid, err := a.store.LookupAddressID(c.Request.Context(), pUserAddr)
	if err != nil {
		common.RespondErrorJSON(c, "Provided address wasn't found")
		return
	}
	success, offset, limit := common.ParseOffsetLimitParamsJSON(c)
	if !success {
		return
	}
	transfers, err := a.repo.CosmicSignatureTransfersByUser(c.Request.Context(), userAid, offset, limit)
	if err != nil {
		a.respondStoreError(c, err)
		return
	}
	reqStatus := 1
	errStr := ""

	c.JSON(http.StatusOK, httpx.H{
		"status":                   reqStatus,
		"error":                    errStr,
		"UserAddr":                 pUserAddr,
		"UserAid":                  userAid,
		"Offset":                   offset,
		"Limit":                    limit,
		"CosmicSignatureTransfers": transfers,
	})
}

func (a *API) handleUsedRwalkNfts(c *httpx.Context) {
	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	if !a.dbInitialized() {
		common.RespondErrorJSON(c, "Database link wasn't configured")
		return
	}

	usedNfts, err := a.repo.RandomWalkTokensUsedInBids(c.Request.Context())
	if err != nil {
		a.respondStoreError(c, err)
		return
	}
	reqStatus := 1
	errStr := ""

	c.JSON(http.StatusOK, httpx.H{
		"status":        reqStatus,
		"error":         errStr,
		"UsedRwalkNFTs": usedNfts,
	})
}

func (a *API) handleMarketingRewardsGlobal(c *httpx.Context) {
	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	if !a.dbInitialized() {
		common.RespondErrorJSON(c, "Database link wasn't configured")
		return
	}
	success, offset, limit := common.ParseOffsetLimitParamsJSON(c)
	if !success {
		return
	}
	rewards, err := a.repo.MarketingRewardHistoryGlobal(c.Request.Context(), offset, limit)
	if err != nil {
		a.respondStoreError(c, err)
		return
	}
	reqStatus := 1
	errStr := ""
	c.JSON(http.StatusOK, httpx.H{
		"status":           reqStatus,
		"error":            errStr,
		"Offset":           offset,
		"Limit":            limit,
		"MarketingRewards": rewards,
	})
}

func (a *API) handleMarketingRewardsByUser(c *httpx.Context) {
	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	if !a.dbInitialized() {
		common.RespondErrorJSON(c, "Database link wasn't configured")
		return
	}
	pUserAddr := c.Param("user_addr")
	userAid, err := a.store.LookupAddressID(c.Request.Context(), pUserAddr)
	if err != nil {
		// Address not in DB yet — return 200 with empty list so UI works
		success, offset, limit := common.ParseOffsetLimitParamsJSON(c)
		if !success {
			return
		}
		c.JSON(http.StatusOK, httpx.H{
			"status": 1, "error": "", "Offset": offset, "Limit": limit,
			"UserAddr": pUserAddr, "UserAid": int64(0), "UserMarketingRewards": []any{},
		})
		return
	}
	success, offset, limit := common.ParseOffsetLimitParamsJSON(c)
	if !success {
		return
	}
	rewards, err := a.repo.MarketingRewardHistoryByUser(c.Request.Context(), userAid, offset, limit)
	if err != nil {
		a.respondStoreError(c, err)
		return
	}
	reqStatus := 1
	errStr := ""
	c.JSON(http.StatusOK, httpx.H{
		"status":               reqStatus,
		"error":                errStr,
		"Offset":               offset,
		"Limit":                limit,
		"UserAddr":             pUserAddr,
		"UserAid":              userAid,
		"UserMarketingRewards": rewards,
	})
}

func (a *API) handleMarketingConfigCurrent(c *httpx.Context) {
	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	if !a.dbInitialized() {
		common.RespondErrorJSON(c, "Database link wasn't configured")
		return
	}

	marketingWalletAddr := a.state.Snapshot().Addrs.MarketingWallet

	// The abigen constructor only fails parsing its embedded ABI constant —
	// statically impossible.
	var copts bind.CallOpts
	contract, _ := cg.NewMarketingWallet(marketingWalletAddr, a.ethClient)

	// Read current treasurer address
	treasurerAddr, err := contract.TreasurerAddress(&copts)
	if err != nil {
		errStr := fmt.Sprintf("Error reading TreasurerAddress: %v", err)
		a.logger.Error(errStr)
		common.RespondErrorJSON(c, errStr)
		return
	}

	// Read token contract address
	tokenAddr, err := contract.Token(&copts)
	if err != nil {
		errStr := fmt.Sprintf("Error reading Token address: %v", err)
		a.logger.Error(errStr)
		common.RespondErrorJSON(c, errStr)
		return
	}

	// Read owner address
	ownerAddr, err := contract.Owner(&copts)
	if err != nil {
		errStr := fmt.Sprintf("Error reading Owner address: %v", err)
		a.logger.Error(errStr)
		common.RespondErrorJSON(c, errStr)
		return
	}

	reqStatus := 1
	errStr := ""
	c.JSON(http.StatusOK, httpx.H{
		"status":              reqStatus,
		"error":               errStr,
		"MarketingWalletAddr": marketingWalletAddr.String(),
		"TreasurerAddr":       treasurerAddr.String(),
		"TokenAddr":           tokenAddr.String(),
		"OwnerAddr":           ownerAddr.String(),
	})
}

func (a *API) handleGetCstPrice(c *httpx.Context) {
	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	// The abigen constructor only fails parsing its embedded ABI constant —
	// statically impossible.
	var copts bind.CallOpts
	contract, _ := cg.NewCosmicSignatureGame(a.state.Snapshot().Addrs.CosmicGame, a.ethClient)
	cstPrice, err := contract.GetNextCstBidPrice(&copts)
	if err != nil {
		a.logger.Error(err.Error())
		common.RespondError(c, err.Error())
		return
	}
	auctionDuration, secondsElapsed, err := contract.GetCstDutchAuctionDurations(&copts)
	if err != nil {
		a.logger.Error(err.Error())
		common.RespondError(c, err.Error())
		return
	}
	reqStatus := 1
	errStr := ""
	c.JSON(http.StatusOK, httpx.H{
		"status":          reqStatus,
		"error":           errStr,
		"CSTPrice":        cstPrice.String(),
		"SecondsElapsed":  secondsElapsed.String(),
		"AuctionDuration": auctionDuration.String(),
	})
}

func (a *API) handleGetEthPrice(c *httpx.Context) {
	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	// The abigen constructor only fails parsing its embedded ABI constant —
	// statically impossible.
	var copts bind.CallOpts
	contract, _ := cg.NewCosmicSignatureGame(a.state.Snapshot().Addrs.CosmicGame, a.ethClient)
	ethPrice, err := contract.GetNextEthBidPrice(&copts)
	if err != nil {
		a.logger.Error(err.Error())
		common.RespondError(c, err.Error())
		return
	}
	auctionDuration, secondsElapsed, err := contract.GetEthDutchAuctionDurations(&copts)
	if err != nil {
		a.logger.Error(err.Error())
		common.RespondError(c, err.Error())
		return
	}
	reqStatus := 1
	errStr := ""
	c.JSON(http.StatusOK, httpx.H{
		"status":          reqStatus,
		"error":           errStr,
		"ETHPrice":        ethPrice.String(),
		"SecondsElapsed":  secondsElapsed.String(),
		"AuctionDuration": auctionDuration.String(),
	})
}

func (a *API) handleSysmodeChanges(c *httpx.Context) {
	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	if !a.dbInitialized() {
		common.RespondErrorJSON(c, "Database link wasn't configured")
		return
	}

	var offset, limit int
	pOffset := c.Param("offset")
	pLimit := c.Param("limit")

	if len(pOffset) == 0 || len(pLimit) == 0 {
		offset = 0
		limit = 10000
	} else {
		var success bool
		success, offset, limit = common.ParseOffsetLimitParamsJSON(c)
		if !success {
			return
		}
	}

	systemModeChanges, err := a.repo.SystemModeChanges(c.Request.Context(), offset, limit)
	if err != nil {
		a.respondStoreError(c, err)
		return
	}

	// When offset=-1 (events from deployment), ensure at least the pre-round-0 row is shown even if no bids yet.
	if offset == -1 && len(systemModeChanges) == 0 {
		systemModeChanges = []cgmodel.CGSystemModeRec{{
			EvtLogId:     -1,
			BlockNum:     -1,
			RoundNum:     0,
			NextEvtLogId: math.MaxInt64,
		}}
	}

	reqStatus := 1
	errStr := ""
	c.JSON(http.StatusOK, httpx.H{
		"status":            reqStatus,
		"error":             errStr,
		"Offset":            offset,
		"Limit":             limit,
		"SystemModeChanges": systemModeChanges,
	})
}

func (a *API) handleAdminEventsInRange(c *httpx.Context) {
	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	if !a.dbInitialized() {
		common.RespondErrorJSON(c, "Database link wasn't configured")
		return
	}

	pEvtlogStart := c.Param("evtlog_start")
	var evtlogStart int64
	if len(pEvtlogStart) > 0 {
		var success bool
		evtlogStart, success = common.ParseIntFromRemoteOrError(c, JSON, &pEvtlogStart)
		if !success {
			return
		}
	}
	pEvtlogEnd := c.Param("evtlog_end")
	var evtlogEnd int64
	if len(pEvtlogEnd) > 0 {
		var success bool
		evtlogEnd, success = common.ParseIntFromRemoteOrError(c, JSON, &pEvtlogEnd)
		if !success {
			return
		}
	}
	eventList, err := a.repo.AdminEventsInRange(c.Request.Context(), evtlogStart, evtlogEnd)
	if err != nil {
		a.respondStoreError(c, err)
		return
	}
	if err := a.repo.ResolveAdminEventValues(c.Request.Context(), eventList); err != nil {
		a.respondStoreError(c, err)
		return
	}
	a.enrichAdminEventsResolvedValues(eventList)

	reqStatus := 1
	errStr := ""
	c.JSON(http.StatusOK, httpx.H{
		"status":        reqStatus,
		"error":         errStr,
		"AdminEvents":   eventList,
		"EvtLogIdStart": evtlogStart,
		"EvtLogIdEnd":   evtlogEnd,
	})
}

func (a *API) handleBidSpecialWinners(c *httpx.Context) {
	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	if !a.dbInitialized() {
		common.RespondErrorJSON(c, "Database link wasn't configured")
		return
	}

	state := a.state.FetchLiveSpecialWinners(c.Request.Context())
	if state.Err != nil {
		common.RespondErrorJSON(c, state.Err.Error())
		return
	}

	resp := httpx.H{
		"status":                          1,
		"error":                           "",
		"LastBidderAddress":               state.LastBidderAddress,
		"LastBidderLastBidTime":           state.LastBidderLastBidTime,
		"EnduranceChampionAddress":        state.EnduranceChampionAddress,
		"EnduranceChampionDuration":       state.EnduranceChampionDuration,
		"EnduranceChampionStartTimeStamp": state.EnduranceChampionStartTimeStamp,
		"PrevEnduranceChampionDuration":   state.PrevEnduranceChampionDuration,
		"ChronoWarriorAddress":            state.ChronoWarriorAddress,
		"ChronoWarriorDuration":           state.ChronoWarriorDuration,
		"ChronoWarriorIsLive":             state.ChronoWarriorIsLive,
		"LastCstBidderAddress":            state.LastCstBidderAddress,
		"RoundNum":                        state.RoundNum,
		"SourceBlockNumber":               state.SourceBlockNumber,
		"SourceBlockTimeStamp":            state.SourceBlockTimeStamp,
	}
	if state.HasLastCstBidderLastBidTime {
		resp["LastCstBidderLastBidTime"] = state.LastCstBidderLastBidTime
	}
	if state.HasLastCstBidEventLogID {
		resp["LastCstBidEventLogId"] = state.LastCstBidEventLogID
	}
	c.JSON(http.StatusOK, resp)
}

// =============================================================================
// BANNED BIDS API (mirrors FastAPI get_banned_bids / ban_bid / unban_bid)
// =============================================================================

func (a *API) handleGetBannedBids(c *httpx.Context) {
	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	if !a.dbInitialized() {
		common.RespondErrorJSON(c, "Database link wasn't configured")
		return
	}
	list, err := a.repo.BannedBids(c.Request.Context())
	if err != nil {
		a.respondStoreError(c, err)
		return
	}
	// Return raw JSON array like FastAPI so frontend gets same shape
	c.JSON(http.StatusOK, list)
}

type banBidPayload struct {
	BidID    int64  `json:"bid_id"`
	UserAddr string `json:"user_addr"`
}

func (a *API) handleBanBid(c *httpx.Context) {
	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	if !a.dbInitialized() {
		common.RespondErrorJSON(c, "Database link wasn't configured")
		return
	}
	// Zero values are rejected like the legacy required-field binding did.
	var payload banBidPayload
	if err := c.ShouldBindJSON(&payload); err != nil || payload.BidID == 0 || payload.UserAddr == "" {
		common.RespondErrorJSON(c, "Invalid JSON: bid_id and user_addr required")
		return
	}
	if err := a.repo.InsertBannedBid(c.Request.Context(), payload.BidID, payload.UserAddr); err != nil &&
		!errors.Is(err, store.ErrConflict) {
		common.RespondErrorJSON(c, fmt.Sprintf("Failed to insert banned bid: %v", err))
		return
	}
	c.JSON(http.StatusCreated, httpx.H{"result": "success"})
}

type unbanBidPayload struct {
	BidID int64 `json:"bid_id"`
}

func (a *API) handleUnbanBid(c *httpx.Context) {
	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	if !a.dbInitialized() {
		common.RespondErrorJSON(c, "Database link wasn't configured")
		return
	}
	// A zero bid_id is rejected like the legacy required-field binding did.
	var payload unbanBidPayload
	if err := c.ShouldBindJSON(&payload); err != nil || payload.BidID == 0 {
		common.RespondErrorJSON(c, "Invalid JSON: bid_id required")
		return
	}
	if err := a.repo.DeleteBannedBid(c.Request.Context(), payload.BidID); err != nil {
		common.RespondErrorJSON(c, fmt.Sprintf("Failed to unban bid: %v", err))
		return
	}
	c.JSON(http.StatusCreated, httpx.H{"result": "success"})
}
