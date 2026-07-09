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
	"github.com/gin-gonic/gin"

	. "github.com/PredictionExplorer/augur-explorer/contracts/cosmicgame"
	"github.com/PredictionExplorer/augur-explorer/internal/api/common"
	p "github.com/PredictionExplorer/augur-explorer/internal/primitives/cosmicgame"
	"github.com/PredictionExplorer/augur-explorer/internal/store"
	cgdb "github.com/PredictionExplorer/augur-explorer/internal/store/cosmicgame"
)

const (
	JSON = true
	HTTP = false
)

// respondUserAddrNotIndexedJSON returns 200 with empty user-shaped data when the wallet
// is not in the indexer DB yet (common for new connections).
func respondUserAddrNotIndexedUserInfoJSON(c *gin.Context, userAddr string) {
	emptyUserInfo := gin.H{
		"AddressId": int64(0), "Address": userAddr, "NumPrizes": int64(0), "NumBids": int64(0),
		"MaxWinAmount": 0.0, "MaxBidAmount": 0.0, "SumRaffleEthWinnings": 0.0, "SumRaffleEthWithdrawal": 0.0,
		"NumRaffleEthWinnings": int64(0), "RaffleNFTsCount": int64(0), "RewardNFTsCount": int64(0),
		"UnclaimedNFTs": int64(0), "TotalCSTokensWon": int64(0), "CosmicSignatureNumTransfers": int64(0),
		"TotalDonatedCount": int64(0), "TotalDonatedAmountEth": 0.0,
		"StakingStatisticsRWalk": gin.H{},
	}
	c.JSON(http.StatusOK, gin.H{
		"status": 1, "error": "", "UserInfo": emptyUserInfo,
		"Bids": []interface{}{}, "MainPrizeClaims": []interface{}{}, "PrizeHistory": []interface{}{},
		"TokenDonationsMade": gin.H{"NFTDonations": []interface{}{}, "ERC20Donations": []interface{}{}},
		"ETHDonationsMade":   []interface{}{}, "MarketingRewardsAwarded": []interface{}{},
		"DonatedNFTsClaimed": []interface{}{}, "DonatedTokensClaimed": []interface{}{},
		"UnclaimedAssets":       gin.H{"ETHPrizes": []interface{}{}, "DonatedNFTs": []interface{}{}},
		"CurrentlyStakedTokens": gin.H{"CST": []interface{}{}, "RWalk": []interface{}{}},
		"StakingActions":        gin.H{"CST": []interface{}{}, "RWalk": []interface{}{}},
		"ERC20Transfers":        []interface{}{}, "ERC721Transfers": []interface{}{},
		"CosmicSignatureTokensOwned": []interface{}{},
	})
}

// respondStoreError logs a store-layer failure and answers with the legacy
// JSON error envelope and HTTP 500; internal details never reach the client.
// A cancelled request context (client went away) is not worth logging.
// These paths previously killed the whole process via os.Exit(1) inside the
// store methods, so any response at all is an improvement.
func respondStoreError(c *gin.Context, err error) {
	if !errors.Is(err, context.Canceled) {
		err_str := fmt.Sprintf("%s: %v", c.FullPath(), err)
		Error.Print(err_str)
		Info.Print(err_str)
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
func sanitizeMapFloatsForJSON(m map[string]interface{}) {
	for k, v := range m {
		if v == nil {
			continue
		}
		switch val := v.(type) {
		case float64:
			if math.IsNaN(val) || math.IsInf(val, 0) {
				m[k] = 0.0
			}
		case map[string]interface{}:
			sanitizeMapFloatsForJSON(val)
		case []interface{}:
			for i, e := range val {
				if f, ok := e.(float64); ok && (math.IsNaN(f) || math.IsInf(f, 0)) {
					val[i] = 0.0
				} else if m2, ok := e.(map[string]interface{}); ok {
					sanitizeMapFloatsForJSON(m2)
				} else if s2, ok := e.([]interface{}); ok {
					sanitizeSliceFloatsForJSON(s2)
				}
			}
		default:
			// Struct values (CurRoundStats, MainStats, etc.) were already sanitized before being put in the map
			break
		}
	}
}

func sanitizeSliceFloatsForJSON(s []interface{}) {
	for i, e := range s {
		if f, ok := e.(float64); ok && (math.IsNaN(f) || math.IsInf(f, 0)) {
			s[i] = 0.0
		} else if m2, ok := e.(map[string]interface{}); ok {
			sanitizeMapFloatsForJSON(m2)
		} else if s2, ok := e.([]interface{}); ok {
			sanitizeSliceFloatsForJSON(s2)
		}
	}
}

// sanitizeFloatsForJSON recursively replaces NaN and ±Inf float64 with 0 in v (in-place).
// Pass a pointer to a struct so fields can be modified. Prevents "json: unsupported value: NaN" panic.
func sanitizeFloatsForJSON(v interface{}) {
	if v == nil {
		return
	}
	val := reflect.ValueOf(v)
	if val.Kind() == reflect.Ptr {
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
			if p := reflect.ValueOf(v); p.Kind() == reflect.Ptr && p.Elem().CanSet() {
				p.Elem().SetFloat(0)
			}
		}
	case reflect.Struct:
		for i := 0; i < val.NumField(); i++ {
			f := val.Field(i)
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
			case reflect.Ptr:
				if !f.IsNil() {
					sanitizeFloatsForJSON(f.Interface())
				}
			case reflect.Slice:
				for j := 0; j < f.Len(); j++ {
					sanitizeFloatsForJSON(f.Index(j).Addr().Interface())
				}
			}
		}
	case reflect.Ptr:
		if !val.IsNil() {
			sanitizeFloatsForJSON(val.Interface())
		}
	case reflect.Slice:
		for i := 0; i < val.Len(); i++ {
			sanitizeFloatsForJSON(val.Index(i).Addr().Interface())
		}
	}
}

// =============================================================================
// DASHBOARD & STATISTICS API
// =============================================================================

func api_cosmic_game_dashboard(c *gin.Context) {

	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")

	snap := contractState.Snapshot()

	// Use live contract price when cache is empty or failed (avoids frontend showing 0 before first refresh or after RPC errors)
	responseBidPrice, responseBidPriceEth := snap.BidPrice, snap.BidPriceEth
	if (responseBidPrice == "" || responseBidPrice == "error" || responseBidPriceEth == 0) && EthClient != nil {
		if contract, err := NewCosmicSignatureGame(snap.Addrs.CosmicGame, EthClient); err == nil {
			var copts bind.CallOpts
			if tmpVal, err := contract.GetNextEthBidPrice(&copts); err == nil {
				responseBidPrice = tmpVal.String()
				fDivisor := big.NewFloat(0.0).SetInt(big.NewInt(1e18))
				fBidPrice := big.NewFloat(0.0).SetInt(tmpVal)
				fQuo := big.NewFloat(0.0).Quo(fBidPrice, fDivisor)
				responseBidPriceEth, _ = fQuo.Float64()
				contractState.SetBidPrice(responseBidPrice, responseBidPriceEth)
			}
		}
	}

	caddrs, err := arbRepo.ContractAddrs(c.Request.Context())
	if err != nil {
		respondStoreError(c, err)
		return
	}
	cur_round_stats, err := arbRepo.CosmicGameRoundStatistics(c.Request.Context(), snap.RoundNum)
	if err != nil {
		respondStoreError(c, err)
		return
	}
	// Live contract: delay before round activation (always); activation time when DB has none
	if EthClient != nil {
		if contract, err := NewCosmicSignatureGame(snap.Addrs.CosmicGame, EthClient); err == nil {
			var copts bind.CallOpts
			if d, err := contract.DelayDurationBeforeRoundActivation(&copts); err == nil {
				cur_round_stats.DelayDurationBeforeRoundActivation = d.Int64()
			}
			if cur_round_stats.ActivationTime == 0 && snap.RoundNum >= 0 {
				if actTime, err := contract.RoundActivationTime(&copts); err == nil {
					cur_round_stats.ActivationTime = actTime.Int64()
				}
			}
		}
	}
	cg_balance := contractState.CosmicGameBalanceEth(c.Request.Context())
	curNumBids := int64(0)
	if snap.RoundNum >= 0 {
		curNumBids, err = arbRepo.BidCountForRound(c.Request.Context(), snap.RoundNum)
		if err != nil {
			respondStoreError(c, err)
			return
		}
	}
	// Sanitize floats so JSON encoding never sees NaN/Inf (avoids "json: unsupported value: NaN" panic)
	sanitizeFloatsForJSON(&cur_round_stats)
	bw_stats_copy := snap.Stats
	sanitizeFloatsForJSON(&bw_stats_copy)
	var req_status int = 1
	var err_str string = ""
	payload := gin.H{
		"status":                               req_status,
		"error":                                err_str,
		"CosmicGameAddr":                       snap.Addrs.CosmicGame,
		"CosmicGameBalanceEth":                 safeFloat64(cg_balance),
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
		"TotalPrizes":                          bw_stats_copy.TotalPrizes,
		"TotalPrizeAwards":                     bw_stats_copy.TotalPrizeAwards,
		"CgPrizeRowCount":                      bw_stats_copy.CgPrizeRowCount,
		"TotalPrizesPaidAmountEth":             bw_stats_copy.TotalPrizesPaidAmountEth,
		"TotalEthDonatedAmount":                bw_stats_copy.TotalEthDonatedAmount,
		"TotalEthDonatedAmountEth":             bw_stats_copy.TotalEthDonatedAmountEth,
		"LastBidderAddr":                       snap.LastBidder.String(),
		"NumVoluntaryDonations":                bw_stats_copy.NumVoluntaryDonations,
		"SumVoluntaryDonationsEth":             bw_stats_copy.SumVoluntaryDonationsEth,
		"NumRwalkTokensUsed":                   bw_stats_copy.NumRwalkTokensUsed,
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
		"NumUniqueBidders":                     bw_stats_copy.NumUniqueBidders,
		"NumUniqueWinners":                     bw_stats_copy.NumUniqueWinners,
		"NumUniqueDonors":                      bw_stats_copy.NumUniqueDonors,
		"NumUniqueStakersCST":                  bw_stats_copy.NumUniqueStakersCST,
		"NumUniqueStakersRWalk":                bw_stats_copy.NumUniqueStakersRWalk,
		"NumUniqueStakersBoth":                 bw_stats_copy.NumUniqueStakersBoth,
		"NumDonatedNFTs":                       bw_stats_copy.NumDonatedNFTs,
		"MainStats":                            bw_stats_copy,
		"CurRoundStats":                        cur_round_stats,
		"TsRoundStart":                         snap.RoundStartTimestamp,
		"ContractAddrs":                        caddrs,
	}
	sanitizeMapFloatsForJSON(payload)
	c.JSON(http.StatusOK, payload)
}

// =============================================================================
// ROUNDS & PRIZES API
// =============================================================================

func api_cosmic_game_prize_list(c *gin.Context) {

	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	if !dbInitialized() {
		common.RespondErrorJSON(c, "Database link wasn't configured")
		return
	}
	success, offset, limit := common.ParseOffsetLimitParamsJSON(c)
	if !success {
		return
	}
	prizes, err := arbRepo.PrizeClaims(c.Request.Context(), offset, limit)
	if err != nil {
		respondStoreError(c, err)
		return
	}

	var req_status int = 1
	var err_str string = ""
	c.JSON(http.StatusOK, gin.H{
		"status": req_status,
		"error":  err_str,
		"Rounds": prizes,
	})
}

// =============================================================================
// BIDS API
// =============================================================================

func api_cosmic_game_bid_list(c *gin.Context) {

	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	if !dbInitialized() {
		common.RespondErrorJSON(c, "Database link wasn't configured")
		return
	}
	success, offset, limit := common.ParseOffsetLimitParamsJSON(c)
	if !success {
		return
	}
	bids, err := arbRepo.Bids(c.Request.Context(), offset, limit)
	if err != nil {
		respondStoreError(c, err)
		return
	}

	var req_status int = 1
	var err_str string = ""
	c.JSON(http.StatusOK, gin.H{
		"status": req_status,
		"error":  err_str,
		"Bids":   bids,
		"Offset": offset,
		"Limit":  limit,
	})
}
func api_cosmic_game_bid_list_by_round(c *gin.Context) {

	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	if !dbInitialized() {
		common.RespondErrorJSON(c, "Database link wasn't configured")
		return
	}

	p_round_num := c.Param("round_num")
	var round_num int64
	if len(p_round_num) > 0 {
		var success bool
		round_num, success = common.ParseIntFromRemoteOrError(c, JSON, &p_round_num)
		if !success {
			return
		}
	} else {
		common.RespondErrorJSON(c, "'round_num' parameter is not set")
		return
	}
	p_sort := c.Param("sort")
	var sort int64
	if len(p_sort) > 0 {
		var success bool
		sort, success = common.ParseIntFromRemoteOrError(c, JSON, &p_sort)
		if !success {
			return
		}
	} else {
		common.RespondErrorJSON(c, "'sort' parameter is not set")
		return
	}
	success, offset, limit := common.ParseOffsetLimitParamsJSON(c)
	if !success {
		return
	}

	bids, total_rows, err := arbRepo.BidsByRound(c.Request.Context(), round_num, int(sort), offset, limit)
	if err != nil {
		respondStoreError(c, err)
		return
	}
	var req_status int = 1
	var err_str string = ""
	c.JSON(http.StatusOK, gin.H{
		"status":      req_status,
		"error":       err_str,
		"RoundNum":    round_num,
		"Offset":      offset,
		"Limit":       limit,
		"Sort":        sort,
		"BidsByRound": bids,
		"TotalRows":   total_rows,
	})
}

type bidMessageTxJSON struct {
	EvtLogId  int64  `json:"evtLogId"`
	BlockNum  int64  `json:"blockNum"`
	TxId      int64  `json:"txId"`
	TxHash    string `json:"txHash"`
	TimeStamp int64  `json:"timeStamp"`
	DateTime  string `json:"dateTime"`
}

type bidWithMessageJSON struct {
	EvtLogId                   int64            `json:"evtLogId"`
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

func bidRecToWithMessageJSON(rec p.CGBidRec) bidWithMessageJSON {
	return bidWithMessageJSON{
		EvtLogId:                   rec.Tx.EvtLogId,
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
			EvtLogId:  rec.Tx.EvtLogId,
			BlockNum:  rec.Tx.BlockNum,
			TxId:      rec.Tx.TxId,
			TxHash:    rec.Tx.TxHash,
			TimeStamp: rec.Tx.TimeStamp,
			DateTime:  rec.Tx.DateTime,
		},
	}
}

func respondBidInfoJSON(c *gin.Context, evtlog_id int64) {
	bid_info, err := arbRepo.BidInfo(c.Request.Context(), evtlog_id)
	if err != nil {
		if errors.Is(err, store.ErrNotFound) {
			common.RespondErrorJSON(c, "record not found")
			return
		}
		respondStoreError(c, err)
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"status":  1,
		"error":   "",
		"BidInfo": bid_info,
	})
}

func api_cosmic_game_bid_info(c *gin.Context) {

	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	if !dbInitialized() {
		common.RespondErrorJSON(c, "Database link wasn't configured")
		return
	}

	p_evtlog_id := c.Param("evtlog_id")
	var evtlog_id int64
	if len(p_evtlog_id) > 0 {
		var success bool
		evtlog_id, success = common.ParseIntFromRemoteOrError(c, JSON, &p_evtlog_id)
		if !success {
			return
		}
	} else {
		common.RespondErrorJSON(c, "'evtlog_id' parameter is not set")
		return
	}
	respondBidInfoJSON(c, evtlog_id)
}

func api_cosmic_game_bid_info_by_pos(c *gin.Context) {

	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	if !dbInitialized() {
		common.RespondErrorJSON(c, "Database link wasn't configured")
		return
	}

	p_round_num := c.Param("round_num")
	var round_num int64
	if len(p_round_num) > 0 {
		var success bool
		round_num, success = common.ParseIntFromRemoteOrError(c, JSON, &p_round_num)
		if !success {
			return
		}
	} else {
		common.RespondErrorJSON(c, "'round_num' parameter is not set")
		return
	}

	p_bid_position := c.Param("bid_position")
	var bid_position int64
	if len(p_bid_position) > 0 {
		var success bool
		bid_position, success = common.ParseIntFromRemoteOrError(c, JSON, &p_bid_position)
		if !success {
			return
		}
	} else {
		common.RespondErrorJSON(c, "'bid_position' parameter is not set")
		return
	}

	evtlog_id, err := arbRepo.EvtlogIDByRoundAndBidPosition(c.Request.Context(), round_num, bid_position)
	if err != nil {
		if errors.Is(err, store.ErrNotFound) {
			common.RespondErrorJSON(c, "record not found")
			return
		}
		respondStoreError(c, err)
		return
	}
	respondBidInfoJSON(c, evtlog_id)
}

func api_cosmic_game_bid_with_message_by_round(c *gin.Context) {

	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	if !dbInitialized() {
		common.RespondErrorJSON(c, "Database link wasn't configured")
		return
	}

	p_round := c.Param("round")
	var round_num int64
	if len(p_round) > 0 {
		var success bool
		round_num, success = common.ParseIntFromRemoteOrError(c, JSON, &p_round)
		if !success {
			return
		}
	} else {
		common.RespondErrorJSON(c, "'round' parameter is not set")
		return
	}

	sortDesc := strings.EqualFold(c.Query("sort"), "desc")
	offset := cgdb.ParseOptionalIntQuery(c.Query("offset"), 0)
	limit := cgdb.ParseOptionalIntQuery(c.Query("limit"), 1000)
	if limit <= 0 {
		limit = 1000
	}

	bids, err := arbRepo.BidsWithMessageByRound(c.Request.Context(), round_num, sortDesc, offset, limit)
	if err != nil {
		respondStoreError(c, err)
		return
	}
	messages := make([]bidWithMessageJSON, 0, len(bids))
	for _, bid := range bids {
		messages = append(messages, bidRecToWithMessageJSON(bid))
	}

	c.JSON(http.StatusOK, gin.H{
		"status":   1,
		"error":    "",
		"Bids":     bids,
		"messages": messages,
		"offset":   offset,
		"limit":    limit,
	})
}
func api_cosmic_game_round_info(c *gin.Context) {

	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	if !dbInitialized() {
		common.RespondErrorJSON(c, "Database link wasn't configured")
		return
	}

	p_prize_num := c.Param("prize_num")
	var prize_num int64
	if len(p_prize_num) > 0 {
		var success bool
		prize_num, success = common.ParseIntFromRemoteOrError(c, JSON, &p_prize_num)
		if !success {
			return
		}
	} else {
		common.RespondErrorJSON(c, "'prize_num' parameter is not set")
		return
	}
	prize_info, err := arbRepo.PrizeInfo(c.Request.Context(), prize_num)
	if errors.Is(err, store.ErrNotFound) {
		common.RespondErrorJSON(c, "record not found")
		return
	}
	if err != nil {
		respondStoreError(c, err)
		return
	}
	var req_status int = 1
	var err_str string = ""
	c.JSON(http.StatusOK, gin.H{
		"status":    req_status,
		"error":     err_str,
		"RoundInfo": prize_info,
	})
}
func api_cosmic_game_user_info(c *gin.Context) {

	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	if !dbInitialized() {
		common.RespondErrorJSON(c, "Database link wasn't configured")
		return
	}

	p_user_addr := c.Param("user_addr")
	if len(p_user_addr) == 0 {
		common.RespondErrorJSON(c, "'user_addr' parameter is not set")
		return
	}
	user_aid, err := arbStore.LookupAddressID(c.Request.Context(), p_user_addr)
	if err != nil {
		respondUserAddrNotIndexedUserInfoJSON(c, p_user_addr)
		return
	}

	user_info, err := arbRepo.UserInfo(c.Request.Context(), user_aid)
	if errors.Is(err, store.ErrNotFound) {
		respondUserAddrNotIndexedUserInfoJSON(c, p_user_addr)
		return
	}
	if err != nil {
		respondStoreError(c, err)
		return
	}
	bids, err := arbRepo.BidsByUser(c.Request.Context(), user_aid)
	if err != nil {
		respondStoreError(c, err)
		return
	}
	main_prize_claims, err := arbRepo.PrizeClaimsByUser(c.Request.Context(), user_aid)
	if err != nil {
		respondStoreError(c, err)
		return
	}
	prize_history, err := arbRepo.PrizeHistoryByUser(c.Request.Context(), user_aid, 0, 1000)
	if err != nil {
		respondStoreError(c, err)
		return
	}

	// Get donations made by user
	nft_donations, err := arbRepo.NFTDonationsByUser(c.Request.Context(), user_aid)
	if err != nil {
		respondStoreError(c, err)
		return
	}
	erc20_donations, err := arbRepo.ERC20DonationsByUser(c.Request.Context(), user_aid)
	if err != nil {
		respondStoreError(c, err)
		return
	}
	eth_donations, err := arbRepo.EthDonationsByUser(c.Request.Context(), user_aid)
	if err != nil {
		respondStoreError(c, err)
		return
	}

	// Get marketing rewards awarded to user
	marketing_rewards, err := arbRepo.MarketingRewardsByUser(c.Request.Context(), user_aid)
	if err != nil {
		respondStoreError(c, err)
		return
	}

	// Get unclaimed assets from PrizesWallet
	unclaimed_eth, err := arbRepo.UnclaimedPrizeEthDeposits(c.Request.Context(), user_aid, 0, 1000)
	if err != nil {
		respondStoreError(c, err)
		return
	}
	unclaimed_nfts, err := arbRepo.UnclaimedDonatedNFTsByUser(c.Request.Context(), user_aid)
	if err != nil {
		respondStoreError(c, err)
		return
	}

	// Get claimed donated assets
	claimed_nfts, err := arbRepo.DonatedNFTClaimsByUser(c.Request.Context(), user_aid)
	if err != nil {
		respondStoreError(c, err)
		return
	}
	claimed_tokens, err := arbRepo.ERC20DonatedPrizesByWinner(c.Request.Context(), user_aid)
	if err != nil {
		respondStoreError(c, err)
		return
	}

	// Get staking data
	staked_cst, err := arbRepo.StakedTokensCstByUser(c.Request.Context(), user_aid)
	if err != nil {
		respondStoreError(c, err)
		return
	}
	staked_rwalk, err := arbRepo.StakedTokensRwalkByUser(c.Request.Context(), user_aid)
	if err != nil {
		respondStoreError(c, err)
		return
	}
	stake_actions_cst, err := arbRepo.StakingActionsCstByUser(c.Request.Context(), user_aid, 0, 1000)
	if err != nil {
		respondStoreError(c, err)
		return
	}
	stake_actions_rwalk, err := arbRepo.StakingActionsRwalkByUser(c.Request.Context(), user_aid, 0, 1000)
	if err != nil {
		respondStoreError(c, err)
		return
	}

	// Get token transfer history
	erc20_transfers, err := arbRepo.CosmicTokenTransfersByUser(c.Request.Context(), user_aid, 0, 1000)
	if err != nil {
		respondStoreError(c, err)
		return
	}
	erc721_transfers, err := arbRepo.CosmicSignatureTransfersByUser(c.Request.Context(), user_aid, 0, 1000)
	if err != nil {
		respondStoreError(c, err)
		return
	}

	// Get owned Cosmic Signature NFTs (where user is current owner)
	owned_cst_nfts, err := arbRepo.CosmicSignatureTokensByUser(c.Request.Context(), user_aid, 0, 1000)
	if err != nil {
		respondStoreError(c, err)
		return
	}

	var req_status int = 1
	var err_str string = ""
	c.JSON(http.StatusOK, gin.H{
		"status":          req_status,
		"error":           err_str,
		"UserInfo":        user_info,
		"Bids":            bids,
		"MainPrizeClaims": main_prize_claims,
		"PrizeHistory":    prize_history,
		"TokenDonationsMade": gin.H{
			"NFTDonations":   nft_donations,
			"ERC20Donations": erc20_donations,
		},
		"ETHDonationsMade":        eth_donations,
		"MarketingRewardsAwarded": marketing_rewards,
		"DonatedNFTsClaimed":      claimed_nfts,
		"DonatedTokensClaimed":    claimed_tokens,
		"UnclaimedAssets": gin.H{
			"ETHPrizes":   unclaimed_eth,
			"DonatedNFTs": unclaimed_nfts,
		},
		"CurrentlyStakedTokens": gin.H{
			"CST":   staked_cst,
			"RWalk": staked_rwalk,
		},
		"StakingActions": gin.H{
			"CST":   stake_actions_cst,
			"RWalk": stake_actions_rwalk,
		},
		"ERC20Transfers":             erc20_transfers,
		"ERC721Transfers":            erc721_transfers,
		"CosmicSignatureTokensOwned": owned_cst_nfts,
	})
}

// =============================================================================
// CHARITY DONATIONS API
// =============================================================================

func api_cosmic_game_charity_cosmicgame_deposits(c *gin.Context) {

	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	if !dbInitialized() {
		common.RespondErrorJSON(c, "Database link wasn't configured")
		return
	}
	cosmicgame_aid, err := arbStore.LookupAddressID(c.Request.Context(), contractState.Snapshot().Addrs.CosmicGame.String())
	if err != nil {
		// Previously this path killed the whole API server (os.Exit); a
		// missing registry row or DB blip now answers HTTP 500 instead.
		respondStoreError(c, err)
		return
	}

	donations, err := arbRepo.CharityDonationsFromCosmicGame(c.Request.Context(), cosmicgame_aid)
	if err != nil {
		respondStoreError(c, err)
		return
	}
	var req_status int = 1
	var err_str string = ""
	c.JSON(http.StatusOK, gin.H{
		"status":           req_status,
		"error":            err_str,
		"CharityDonations": donations,
	})
}
func api_cosmic_game_charity_voluntary_deposits(c *gin.Context) {

	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	if !dbInitialized() {
		common.RespondErrorJSON(c, "Database link wasn't configured")
		return
	}
	cosmicgame_aid, err := arbStore.LookupAddressID(c.Request.Context(), contractState.Snapshot().Addrs.CosmicGame.String())
	if err != nil {
		// Previously this path killed the whole API server (os.Exit); a
		// missing registry row or DB blip now answers HTTP 500 instead.
		respondStoreError(c, err)
		return
	}

	donations, err := arbRepo.CharityDonationsVoluntary(c.Request.Context(), cosmicgame_aid)
	if err != nil {
		respondStoreError(c, err)
		return
	}
	var req_status int = 1
	var err_str string = ""
	c.JSON(http.StatusOK, gin.H{
		"status":           req_status,
		"error":            err_str,
		"CharityDonations": donations,
	})
}
func api_cosmic_game_charity_donations_deposits(c *gin.Context) {

	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	if !dbInitialized() {
		common.RespondErrorJSON(c, "Database link wasn't configured")
		return
	}
	cosmicgame_aid, err := arbStore.LookupAddressID(c.Request.Context(), contractState.Snapshot().Addrs.CosmicGame.String())
	if err != nil {
		// Previously this path killed the whole API server (os.Exit); a
		// missing registry row or DB blip now answers HTTP 500 instead.
		respondStoreError(c, err)
		return
	}

	donations, err := arbRepo.CharityDonations(c.Request.Context(), cosmicgame_aid)
	if err != nil {
		respondStoreError(c, err)
		return
	}
	var req_status int = 1
	var err_str string = ""
	c.JSON(http.StatusOK, gin.H{
		"status":           req_status,
		"error":            err_str,
		"CharityDonations": donations,
	})
}
func api_cosmic_game_charity_donations_withdrawals(c *gin.Context) {

	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	if !dbInitialized() {
		common.RespondErrorJSON(c, "Database link wasn't configured")
		return
	}

	withdrawals, err := arbRepo.CharityWalletWithdrawals(c.Request.Context())
	if err != nil {
		respondStoreError(c, err)
		return
	}
	var req_status int = 1
	var err_str string = ""
	c.JSON(http.StatusOK, gin.H{
		"status":             req_status,
		"error":              err_str,
		"CharityWithdrawals": withdrawals,
	})
}

// =============================================================================
// UNIQUE PARTICIPANTS API
// =============================================================================

func api_cosmic_game_user_unique_bidders(c *gin.Context) {

	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	if !dbInitialized() {
		common.RespondErrorJSON(c, "Database link wasn't configured")
		return
	}

	unique_bidders, err := arbRepo.UniqueBidders(c.Request.Context())
	if err != nil {
		respondStoreError(c, err)
		return
	}

	var req_status int = 1
	var err_str string = ""
	c.JSON(http.StatusOK, gin.H{
		"status":        req_status,
		"error":         err_str,
		"UniqueBidders": unique_bidders,
	})
}
func api_cosmic_game_user_unique_winners(c *gin.Context) {

	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	if !dbInitialized() {
		common.RespondErrorJSON(c, "Database link wasn't configured")
		return
	}

	unique_winners, err := arbRepo.UniqueWinners(c.Request.Context())
	if err != nil {
		respondStoreError(c, err)
		return
	}

	var req_status int = 1
	var err_str string = ""
	c.JSON(http.StatusOK, gin.H{
		"status":        req_status,
		"error":         err_str,
		"UniqueWinners": unique_winners,
	})
}
func api_cosmic_game_roi_leaderboard(c *gin.Context) {

	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	if !dbInitialized() {
		common.RespondErrorJSON(c, "Database link wasn't configured")
		return
	}

	min_bids := cgdb.ParseOptionalIntQuery(c.Query("min_bids"), 5)
	if min_bids < 0 {
		min_bids = 0
	}
	sort_by := c.Query("sort") // one of: net_pl(default), roi, winrate, spent, nfts, bids
	offset := cgdb.ParseOptionalIntQuery(c.Query("offset"), 0)
	if offset < 0 {
		offset = 0
	}
	limit := cgdb.ParseOptionalIntQuery(c.Query("limit"), 100)
	if limit <= 0 || limit > 1000 {
		limit = 100
	}

	leaderboard, err := arbRepo.RoiLeaderboard(c.Request.Context(), min_bids, sort_by, offset, limit)
	if err != nil {
		respondStoreError(c, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":         1,
		"error":          "",
		"MinBids":        min_bids,
		"Sort":           sort_by,
		"Offset":         offset,
		"Limit":          limit,
		"RoiLeaderboard": leaderboard,
	})
}
func api_cosmic_game_claims_by_round(c *gin.Context) {

	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	if !dbInitialized() {
		common.RespondErrorJSON(c, "Database link wasn't configured")
		return
	}

	claims, err := arbRepo.ClaimsByRound(c.Request.Context())
	if err != nil {
		respondStoreError(c, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":        1,
		"error":         "",
		"ClaimsByRound": claims,
	})
}
func api_cosmic_game_claim_detail_by_round(c *gin.Context) {

	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	if !dbInitialized() {
		common.RespondErrorJSON(c, "Database link wasn't configured")
		return
	}

	round := cgdb.ParseOptionalIntQuery(c.Param("round_num"), -1)
	if round < 0 {
		common.RespondErrorJSON(c, "Invalid round_num")
		return
	}

	detail, err := arbRepo.ClaimDetailByRound(c.Request.Context(), int64(round))
	if err != nil {
		respondStoreError(c, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":            1,
		"error":             "",
		"RoundNum":          detail.RoundNum,
		"ClaimTransactions": detail.ClaimTransactions,
		"AttachedTokens":    detail.AttachedTokens,
	})
}
func api_cosmic_game_user_unique_donors(c *gin.Context) {

	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	if !dbInitialized() {
		common.RespondErrorJSON(c, "Database link wasn't configured")
		return
	}

	unique_donors, err := arbRepo.UniqueDonors(c.Request.Context())
	if err != nil {
		respondStoreError(c, err)
		return
	}

	var req_status int = 1
	var err_str string = ""
	c.JSON(http.StatusOK, gin.H{
		"status":       req_status,
		"error":        err_str,
		"UniqueDonors": unique_donors,
	})
}

// =============================================================================
// NFT DONATIONS API
// =============================================================================

func api_cosmic_game_donations_nft_list(c *gin.Context) {

	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	if !dbInitialized() {
		common.RespondErrorJSON(c, "Database link wasn't configured")
		return
	}

	success, offset, limit := common.ParseOffsetLimitParamsJSON(c)
	if !success {
		return
	}
	nft_donations, err := arbRepo.NFTDonations(c.Request.Context(), offset, limit)
	if err != nil {
		respondStoreError(c, err)
		return
	}

	var req_status int = 1
	var err_str string = ""
	c.JSON(http.StatusOK, gin.H{
		"status":       req_status,
		"error":        err_str,
		"NFTDonations": nft_donations,
		"Offset":       offset,
		"Limit":        limit,
	})
}
func api_cosmic_game_nft_donation_stats(c *gin.Context) {

	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	if !dbInitialized() {
		common.RespondErrorJSON(c, "Database link wasn't configured")
		return
	}

	nft_donation_stats, err := arbRepo.NFTDonationStats(c.Request.Context())
	if err != nil {
		respondStoreError(c, err)
		return
	}

	var req_status int = 1
	var err_str string = ""
	c.JSON(http.StatusOK, gin.H{
		"status":           req_status,
		"error":            err_str,
		"NFTDonationStats": nft_donation_stats,
	})
}
func api_cosmic_game_nft_donations_by_user(c *gin.Context) {
	// DONOR PERSPECTIVE: Returns NFTs this user DONATED

	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	if !dbInitialized() {
		common.RespondErrorJSON(c, "Database link wasn't configured")
		return
	}
	p_user_addr := c.Param("user_addr")
	if len(p_user_addr) == 0 {
		common.RespondErrorJSON(c, "'user_addr' parameter is not set")
		return
	}
	user_aid, err := arbStore.LookupAddressID(c.Request.Context(), p_user_addr)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"status": 1, "error": "", "NFTDonationsByDonor": []interface{}{},
			"UserAddr": p_user_addr, "UserAid": int64(0),
		})
		return
	}

	donations, err := arbRepo.NFTDonationsByUser(c.Request.Context(), user_aid)
	if err != nil {
		respondStoreError(c, err)
		return
	}
	var req_status int = 1
	var err_str string = ""
	c.JSON(http.StatusOK, gin.H{
		"status":              req_status,
		"error":               err_str,
		"NFTDonationsByDonor": donations,
		"UserAddr":            p_user_addr,
		"UserAid":             user_aid,
	})
}
func api_cosmic_game_record_counters(c *gin.Context) {

	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	if !dbInitialized() {
		common.RespondErrorJSON(c, "Database link wasn't configured")
		return
	}

	record_counters, err := arbRepo.RecordCounters(c.Request.Context())
	if err != nil {
		respondStoreError(c, err)
		return
	}

	var req_status int = 1
	var err_str string = ""
	c.JSON(http.StatusOK, gin.H{
		"status":         req_status,
		"error":          err_str,
		"RecordCounters": record_counters,
	})
}
func api_cosmic_game_donated_nft_info(c *gin.Context) {

	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	if !dbInitialized() {
		common.RespondErrorJSON(c, "Database link wasn't configured")
		return
	}

	p_record_id := c.Param("record_id")
	var record_id int64
	if len(p_record_id) > 0 {
		var success bool
		record_id, success = common.ParseIntFromRemoteOrError(c, JSON, &p_record_id)
		if !success {
			return
		}
	} else {
		common.RespondErrorJSON(c, "'record_id' parameter is not set")
		return
	}
	nftdonation, err := arbRepo.NFTDonationInfo(c.Request.Context(), record_id)
	if err != nil {
		if errors.Is(err, store.ErrNotFound) {
			common.RespondErrorJSON(c, "Record not found")
			return
		}
		respondStoreError(c, err)
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"status":      1,
		"error":       "",
		"NFTDonation": nftdonation,
	})
}

// =============================================================================
// PRIZE & RAFFLE DEPOSITS API
// =============================================================================

func api_cosmic_game_prize_deposits_list(c *gin.Context) {

	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	if !dbInitialized() {
		common.RespondErrorJSON(c, "Database link wasn't configured")
		return
	}

	var offset, limit int
	p_offset := c.Param("offset")
	p_limit := c.Param("limit")

	if len(p_offset) == 0 || len(p_limit) == 0 {
		offset = 0
		limit = 10000
	} else {
		var success bool
		success, offset, limit = common.ParseOffsetLimitParamsJSON(c)
		if !success {
			return
		}
	}

	deposits, err := arbRepo.PrizeEthDeposits(c.Request.Context(), offset, limit)
	if err != nil {
		respondStoreError(c, err)
		return
	}

	var req_status int = 1
	var err_str string = ""
	c.JSON(http.StatusOK, gin.H{
		"status":         req_status,
		"error":          err_str,
		"RaffleDeposits": deposits,
		"Offset":         offset,
		"Limit":          limit,
	})
}
func api_cosmic_game_all_eth_deposits_list(c *gin.Context) {

	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	if !dbInitialized() {
		common.RespondErrorJSON(c, "Database link wasn't configured")
		return
	}

	var offset, limit int
	p_offset := c.Param("offset")
	p_limit := c.Param("limit")

	if len(p_offset) == 0 || len(p_limit) == 0 {
		offset = 0
		limit = 10000
	} else {
		var success bool
		success, offset, limit = common.ParseOffsetLimitParamsJSON(c)
		if !success {
			return
		}
	}

	deposits, err := arbRepo.PrizeEthDeposits(c.Request.Context(), offset, limit)
	if err != nil {
		respondStoreError(c, err)
		return
	}

	var req_status int = 1
	var err_str string = ""
	c.JSON(http.StatusOK, gin.H{
		"status":      req_status,
		"error":       err_str,
		"AllDeposits": deposits,
		"Offset":      offset,
		"Limit":       limit,
	})
}
func api_cosmic_game_raffle_eth_deposits_list(c *gin.Context) {

	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	if !dbInitialized() {
		common.RespondErrorJSON(c, "Database link wasn't configured")
		return
	}

	var offset, limit int
	p_offset := c.Param("offset")
	p_limit := c.Param("limit")

	if len(p_offset) == 0 || len(p_limit) == 0 {
		offset = 0
		limit = 10000
	} else {
		var success bool
		success, offset, limit = common.ParseOffsetLimitParamsJSON(c)
		if !success {
			return
		}
	}

	deposits, err := arbRepo.RaffleEthDeposits(c.Request.Context(), offset, limit)
	if err != nil {
		respondStoreError(c, err)
		return
	}

	var req_status int = 1
	var err_str string = ""
	c.JSON(http.StatusOK, gin.H{
		"status":         req_status,
		"error":          err_str,
		"RaffleDeposits": deposits,
		"Offset":         offset,
		"Limit":          limit,
	})
}
func api_cosmic_game_chronowarrior_eth_deposits_list(c *gin.Context) {

	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	if !dbInitialized() {
		common.RespondErrorJSON(c, "Database link wasn't configured")
		return
	}

	var offset, limit int
	p_offset := c.Param("offset")
	p_limit := c.Param("limit")

	if len(p_offset) == 0 || len(p_limit) == 0 {
		offset = 0
		limit = 10000
	} else {
		var success bool
		success, offset, limit = common.ParseOffsetLimitParamsJSON(c)
		if !success {
			return
		}
	}

	deposits, err := arbRepo.ChronoWarriorEthDeposits(c.Request.Context(), offset, limit)
	if err != nil {
		respondStoreError(c, err)
		return
	}

	var req_status int = 1
	var err_str string = ""
	c.JSON(http.StatusOK, gin.H{
		"status":                req_status,
		"error":                 err_str,
		"ChronoWarriorDeposits": deposits,
		"Offset":                offset,
		"Limit":                 limit,
	})
}

// Unified URI scheme handlers - per-user
func api_cosmic_game_unified_eth_all_by_user(c *gin.Context) {

	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	if !dbInitialized() {
		common.RespondErrorJSON(c, "Database link wasn't configured")
		return
	}
	p_user_addr := c.Param("user_addr")
	if len(p_user_addr) == 0 {
		common.RespondErrorJSON(c, "'user_addr' parameter is not set")
		return
	}
	user_aid, err := arbStore.LookupAddressID(c.Request.Context(), p_user_addr)
	if err != nil {
		common.RespondErrorJSON(c, "Provided address wasn't found")
		return
	}

	deposits, err := arbRepo.EthDepositsByUser(c.Request.Context(), user_aid)
	if err != nil {
		respondStoreError(c, err)
		return
	}

	var req_status int = 1
	var err_str string = ""
	c.JSON(http.StatusOK, gin.H{
		"status":      req_status,
		"error":       err_str,
		"UserAddr":    p_user_addr,
		"UserAid":     user_aid,
		"AllDeposits": deposits,
	})
}
func api_cosmic_game_unified_eth_raffle_by_user(c *gin.Context) {

	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	if !dbInitialized() {
		common.RespondErrorJSON(c, "Database link wasn't configured")
		return
	}
	p_user_addr := c.Param("user_addr")
	if len(p_user_addr) == 0 {
		common.RespondErrorJSON(c, "'user_addr' parameter is not set")
		return
	}
	user_aid, err := arbStore.LookupAddressID(c.Request.Context(), p_user_addr)
	if err != nil {
		common.RespondErrorJSON(c, "Provided address wasn't found")
		return
	}

	deposits, err := arbRepo.RaffleEthDepositsByUser(c.Request.Context(), user_aid)
	if err != nil {
		respondStoreError(c, err)
		return
	}

	var req_status int = 1
	var err_str string = ""
	c.JSON(http.StatusOK, gin.H{
		"status":         req_status,
		"error":          err_str,
		"UserAddr":       p_user_addr,
		"UserAid":        user_aid,
		"RaffleDeposits": deposits,
	})
}
func api_cosmic_game_unified_eth_chronowarrior_by_user(c *gin.Context) {

	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	if !dbInitialized() {
		common.RespondErrorJSON(c, "Database link wasn't configured")
		return
	}
	p_user_addr := c.Param("user_addr")
	if len(p_user_addr) == 0 {
		common.RespondErrorJSON(c, "'user_addr' parameter is not set")
		return
	}
	user_aid, err := arbStore.LookupAddressID(c.Request.Context(), p_user_addr)
	if err != nil {
		common.RespondErrorJSON(c, "Provided address wasn't found")
		return
	}

	deposits, err := arbRepo.ChronoWarriorEthDepositsByUser(c.Request.Context(), user_aid)
	if err != nil {
		respondStoreError(c, err)
		return
	}

	var req_status int = 1
	var err_str string = ""
	c.JSON(http.StatusOK, gin.H{
		"status":                req_status,
		"error":                 err_str,
		"UserAddr":              p_user_addr,
		"UserAid":               user_aid,
		"ChronoWarriorDeposits": deposits,
	})
}
func api_cosmic_game_prize_deposits_by_round(c *gin.Context) {

	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	if !dbInitialized() {
		common.RespondErrorJSON(c, "Database link wasn't configured")
		return
	}

	p_round_num := c.Param("round_num")
	var round_num int64
	if len(p_round_num) > 0 {
		var success bool
		round_num, success = common.ParseIntFromRemoteOrError(c, JSON, &p_round_num)
		if !success {
			return
		}
	} else {
		common.RespondErrorJSON(c, "'round_num' parameter is not set")
		return
	}

	deposits, err := arbRepo.PrizeDepositsByRound(c.Request.Context(), round_num)
	if err != nil {
		respondStoreError(c, err)
		return
	}

	var req_status int = 1
	var err_str string = ""
	c.JSON(http.StatusOK, gin.H{
		"status":         req_status,
		"error":          err_str,
		"RaffleDeposits": deposits,
		"RoundNum":       round_num,
	})
}
func api_cosmic_game_raffle_nft_winners_list(c *gin.Context) {

	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	if !dbInitialized() {
		common.RespondErrorJSON(c, "Database link wasn't configured")
		return
	}

	var offset, limit int
	p_offset := c.Param("offset")
	p_limit := c.Param("limit")

	if len(p_offset) == 0 || len(p_limit) == 0 {
		offset = 0
		limit = 10000
	} else {
		var success bool
		success, offset, limit = common.ParseOffsetLimitParamsJSON(c)
		if !success {
			return
		}
	}

	winners, err := arbRepo.RaffleNFTWinners(c.Request.Context(), offset, limit)
	if err != nil {
		respondStoreError(c, err)
		return
	}

	var req_status int = 1
	var err_str string = ""
	c.JSON(http.StatusOK, gin.H{
		"status":           req_status,
		"error":            err_str,
		"RaffleNFTWinners": winners,
		"Offset":           offset,
		"Limit":            limit,
	})
}
func api_cosmic_game_raffle_nft_winners_by_round(c *gin.Context) {

	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	if !dbInitialized() {
		common.RespondErrorJSON(c, "Database link wasn't configured")
		return
	}

	p_round_num := c.Param("round_num")
	var round_num int64
	if len(p_round_num) > 0 {
		var success bool
		round_num, success = common.ParseIntFromRemoteOrError(c, JSON, &p_round_num)
		if !success {
			return
		}
	} else {
		common.RespondErrorJSON(c, "'round_num' parameter is not set")
		return
	}

	winners_raffle, err := arbRepo.RaffleNFTWinnersByRound(c.Request.Context(), round_num, false)
	if err != nil {
		respondStoreError(c, err)
		return
	}
	winners_staking, err := arbRepo.RaffleNFTWinnersByRound(c.Request.Context(), round_num, true)
	if err != nil {
		respondStoreError(c, err)
		return
	}

	var req_status int = 1
	var err_str string = ""
	c.JSON(http.StatusOK, gin.H{
		"status":            req_status,
		"error":             err_str,
		"RaffleNFTWinners":  winners_raffle,
		"StakingNFTWinners": winners_staking,
		"RoundNum":          round_num,
	})
}
func api_cosmic_game_user_raffle_nft_winnings(c *gin.Context) {

	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	if !dbInitialized() {
		common.RespondErrorJSON(c, "Database link wasn't configured")
		return
	}

	p_user_addr := c.Param("user_addr")
	if len(p_user_addr) == 0 {
		common.RespondErrorJSON(c, "'user_addr' parameter is not set")
		return
	}
	user_aid, err := arbStore.LookupAddressID(c.Request.Context(), p_user_addr)
	if err != nil {
		common.RespondErrorJSON(c, "Provided address wasn't found")
		return
	}

	user_info, err := arbRepo.UserInfo(c.Request.Context(), user_aid)
	if errors.Is(err, store.ErrNotFound) {
		common.RespondErrorJSON(c, "Provided address wasn't found")
		return
	}
	if err != nil {
		respondStoreError(c, err)
		return
	}

	winnings, err := arbRepo.RaffleNFTWinningsByUser(c.Request.Context(), user_aid)
	if err != nil {
		respondStoreError(c, err)
		return
	}

	var req_status int = 1
	var err_str string = ""
	c.JSON(http.StatusOK, gin.H{
		"status":                req_status,
		"error":                 err_str,
		"UserRaffleNFTWinnings": winnings,
		"UserInfo":              user_info,
	})
}
func api_cosmic_game_prize_deposits_raffle_eth_by_user(c *gin.Context) {

	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	if !dbInitialized() {
		common.RespondErrorJSON(c, "Database link wasn't configured")
		return
	}

	p_user_addr := c.Param("user_addr")
	if len(p_user_addr) == 0 {
		common.RespondErrorJSON(c, "'user_addr' parameter is not set")
		return
	}
	user_aid, err := arbStore.LookupAddressID(c.Request.Context(), p_user_addr)
	if err != nil {
		common.RespondErrorJSON(c, "Provided address wasn't found")
		return
	}
	user_info, err := arbRepo.UserInfo(c.Request.Context(), user_aid)
	if errors.Is(err, store.ErrNotFound) {
		common.RespondErrorJSON(c, "Provided address wasn't found")
		return
	}
	if err != nil {
		respondStoreError(c, err)
		return
	}

	deposits, err := arbRepo.PrizeDepositsRaffleEthByUser(c.Request.Context(), user_aid)
	if err != nil {
		respondStoreError(c, err)
		return
	}
	var req_status int = 1
	var err_str string = ""
	c.JSON(http.StatusOK, gin.H{
		"status":             req_status,
		"error":              err_str,
		"UserRaffleDeposits": deposits,
		"UserInfo":           user_info,
	})
}
func api_cosmic_game_prize_deposits_chrono_warrior_by_user(c *gin.Context) {

	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	if !dbInitialized() {
		common.RespondErrorJSON(c, "Database link wasn't configured")
		return
	}

	p_user_addr := c.Param("user_addr")
	if len(p_user_addr) == 0 {
		common.RespondErrorJSON(c, "'user_addr' parameter is not set")
		return
	}
	user_aid, err := arbStore.LookupAddressID(c.Request.Context(), p_user_addr)
	if err != nil {
		common.RespondErrorJSON(c, "Provided address wasn't found")
		return
	}
	user_info, err := arbRepo.UserInfo(c.Request.Context(), user_aid)
	if errors.Is(err, store.ErrNotFound) {
		common.RespondErrorJSON(c, "Provided address wasn't found")
		return
	}
	if err != nil {
		respondStoreError(c, err)
		return
	}

	deposits, err := arbRepo.PrizeDepositsChronoWarriorByUser(c.Request.Context(), user_aid)
	if err != nil {
		respondStoreError(c, err)
		return
	}
	var req_status int = 1
	var err_str string = ""
	c.JSON(http.StatusOK, gin.H{
		"status":                    req_status,
		"error":                     err_str,
		"UserChronoWarriorDeposits": deposits,
		"UserInfo":                  user_info,
	})
}
func api_cosmic_game_nft_donations_by_prize(c *gin.Context) {

	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	if !dbInitialized() {
		common.RespondErrorJSON(c, "Database link wasn't configured")
		return
	}

	p_prize_num := c.Param("prize_num")
	var prize_num int64
	if len(p_prize_num) > 0 {
		var success bool
		prize_num, success = common.ParseIntFromRemoteOrError(c, JSON, &p_prize_num)
		if !success {
			return
		}
	} else {
		common.RespondErrorJSON(c, "'prize_num' parameter is not set")
		return
	}
	nft_donations, err := arbRepo.NFTDonationsByRound(c.Request.Context(), prize_num)
	if err != nil {
		respondStoreError(c, err)
		return
	}
	var req_status int = 1
	var err_str string = ""
	c.JSON(http.StatusOK, gin.H{
		"status":       req_status,
		"error":        err_str,
		"NFTDonations": nft_donations,
		"RoundNum":     prize_num,
	})
}
func api_cosmic_game_nft_donations_by_token(c *gin.Context) {
	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	if !dbInitialized() {
		common.RespondErrorJSON(c, "Database link wasn't configured")
		return
	}
	p_token_addr := c.Param("token_addr")
	if len(p_token_addr) == 0 {
		common.RespondErrorJSON(c, "'token_addr' parameter is not set")
		return
	}
	token_aid, err := arbStore.LookupAddressID(c.Request.Context(), p_token_addr)
	if err != nil {
		common.RespondErrorJSON(c, "Token address not found")
		return
	}
	nft_donations, err := arbRepo.NFTDonationsByToken(c.Request.Context(), token_aid)
	if err != nil {
		respondStoreError(c, err)
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"status":       1,
		"error":        "",
		"NFTDonations": nft_donations,
		"TokenAddr":    p_token_addr,
	})
}
func api_cosmic_game_cosmic_signature_token_list(c *gin.Context) {

	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	if !dbInitialized() {
		common.RespondErrorJSON(c, "Database link wasn't configured")
		return
	}

	success, offset, limit := common.ParseOffsetLimitParamsJSON(c)
	if !success {
		return
	}
	tokens, err := arbRepo.CosmicSignatureTokens(c.Request.Context(), offset, limit)
	if err != nil {
		respondStoreError(c, err)
		return
	}

	var req_status int = 1
	var err_str string = ""
	c.JSON(http.StatusOK, gin.H{
		"status":                   req_status,
		"error":                    err_str,
		"CosmicSignatureTokenList": tokens,
		"Offset":                   offset,
		"Limit":                    limit,
	})
}
func api_cosmic_game_cosmic_signature_token_info(c *gin.Context) {

	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	if !dbInitialized() {
		common.RespondErrorJSON(c, "Database link wasn't configured")
		return
	}

	p_token_id := c.Param("token_id")
	var token_id int64
	if len(p_token_id) > 0 {
		var success bool
		token_id, success = common.ParseIntFromRemoteOrError(c, JSON, &p_token_id)
		if !success {
			return
		}
	} else {
		common.RespondErrorJSON(c, "'token_id' parameter is not set")
		return
	}

	token_info, err := arbRepo.CosmicSignatureTokenInfo(c.Request.Context(), token_id)
	if err != nil {
		if errors.Is(err, store.ErrNotFound) {
			common.RespondErrorJSON(c, "record not found")
			return
		}
		respondStoreError(c, err)
		return
	}

	var req_status int = 1
	var err_str string = ""

	if token_info.RecordType == 3 {
		// The legacy code ignored the found/not-found flag and rendered the
		// zero-value record when the round has no prize claim; ErrNotFound
		// keeps that exact behavior.
		prize_info, err := arbRepo.PrizeInfo(c.Request.Context(), token_info.RoundNum)
		if err != nil && !errors.Is(err, store.ErrNotFound) {
			respondStoreError(c, err)
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"status":    req_status,
			"error":     err_str,
			"TokenInfo": token_info,
			"PrizeInfo": prize_info,
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"status":    req_status,
			"error":     err_str,
			"TokenInfo": token_info,
		})
	}
}
func api_cosmic_game_donated_nft_claims_all(c *gin.Context) {

	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	if !dbInitialized() {
		common.RespondErrorJSON(c, "Database link wasn't configured")
		return
	}

	var offset, limit int
	p_offset := c.Param("offset")
	p_limit := c.Param("limit")

	if len(p_offset) == 0 || len(p_limit) == 0 {
		offset = 0
		limit = 10000
	} else {
		var success bool
		success, offset, limit = common.ParseOffsetLimitParamsJSON(c)
		if !success {
			return
		}
	}

	claims, err := arbRepo.DonatedNFTClaims(c.Request.Context(), offset, limit)
	if err != nil {
		respondStoreError(c, err)
		return
	}

	var req_status int = 1
	var err_str string = ""
	c.JSON(http.StatusOK, gin.H{
		"status":           req_status,
		"error":            err_str,
		"DonatedNFTClaims": claims,
		"Offset":           offset,
		"Limit":            limit,
	})
}
func api_cosmic_game_donated_nft_claims_by_user(c *gin.Context) {
	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	if !dbInitialized() {
		common.RespondErrorJSON(c, "Database link wasn't configured")
		return
	}
	p_user_addr := c.Param("user_addr")
	if len(p_user_addr) == 0 {
		common.RespondErrorJSON(c, "'user_addr' parameter is not set")
		return
	}
	user_aid, err := arbStore.LookupAddressID(c.Request.Context(), p_user_addr)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"status": 1, "error": "", "DonatedNFTClaims": []interface{}{},
			"UserInfo": p.CGUserInfo{Address: p_user_addr},
		})
		return
	}
	user_info, err := arbRepo.UserInfo(c.Request.Context(), user_aid)
	if errors.Is(err, store.ErrNotFound) {
		c.JSON(http.StatusOK, gin.H{
			"status": 1, "error": "", "DonatedNFTClaims": []interface{}{},
			"UserInfo": p.CGUserInfo{Address: p_user_addr},
		})
		return
	}
	if err != nil {
		respondStoreError(c, err)
		return
	}
	claims, err := arbRepo.DonatedNFTClaimsByUser(c.Request.Context(), user_aid)
	if err != nil {
		respondStoreError(c, err)
		return
	}
	var req_status int = 1
	var err_str string = ""
	c.JSON(http.StatusOK, gin.H{
		"status":           req_status,
		"error":            err_str,
		"DonatedNFTClaims": claims,
		"UserInfo":         user_info,
	})
}
func api_cosmic_game_time_current(c *gin.Context) {

	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	if !dbInitialized() {
		common.RespondErrorJSON(c, "Database link wasn't configured")
		return
	}
	var raw json.RawMessage
	err := rpcclient.CallContext(context.Background(), &raw, "eth_getBlockByNumber", "pending", true)
	if err != nil {
		common.RespondErrorJSON(c, fmt.Sprintf("%v", err))
		return
	}
	var rpcobj map[string]interface{}
	rpcobj = make(map[string]interface{})
	err = json.Unmarshal(raw, &rpcobj)
	if err != nil {
		common.RespondErrorJSON(c, fmt.Sprintf("Error decoding JSON: %v", err))
		return
	}

	ts_hex := rpcobj["timestamp"].(string)
	ts, err := hexutil.DecodeUint64(ts_hex)
	if err != nil {
		common.RespondErrorJSON(c, fmt.Sprintf("Error decoding timestamp from hex: %v", err))
		return
	}
	var req_status int = 1
	var err_str string = ""
	c.JSON(http.StatusOK, gin.H{
		"status":           req_status,
		"error":            err_str,
		"CurrentTimeStamp": ts,
	})
}
func api_cosmic_game_time_until_prize(c *gin.Context) {

	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	if !dbInitialized() {
		common.RespondErrorJSON(c, "Database link wasn't configured")
		return
	}
	// Function: getDurationUntilMainPrize() - returns time until prize can be claimed
	const time_until_prize_sig string = "0x36750d2c"
	var raw json.RawMessage
	args := map[string]interface{}{
		"to":   contractState.Snapshot().Addrs.CosmicGame.String(),
		"data": time_until_prize_sig,
	}
	err := rpcclient.CallContext(context.Background(), &raw, "eth_call", args, "pending")
	if err != nil {
		common.RespondErrorJSON(c, fmt.Sprintf("%v", err))
		return
	}
	var ts_hex string
	err = json.Unmarshal(raw, &ts_hex)
	if err != nil {
		common.RespondErrorJSON(c, fmt.Sprintf("Error decoding JSON: %v", err))
		return
	}
	ts_big := ethcommon.HexToHash(ts_hex).Big()
	var req_status int = 1
	var err_str string = ""
	c.JSON(http.StatusOK, gin.H{
		"status":         req_status,
		"error":          err_str,
		"TimeUntilPrize": ts_big.Int64(),
	})
}
func api_cosmic_game_prize_cur_round_time(c *gin.Context) {

	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	if !dbInitialized() {
		common.RespondErrorJSON(c, "Database link wasn't configured")
		return
	}

	// Return 200 with status/error in body when contract is unavailable, so frontend gets consistent response shape
	var req_status int = 1
	var err_str string
	var prize_time *big.Int

	if EthClient == nil {
		req_status = 0
		err_str = "Ethereum client not configured"
		prize_time = big.NewInt(0)
	} else {
		var copts bind.CallOpts
		bwcontract, err := NewCosmicSignatureGame(contractState.Snapshot().Addrs.CosmicGame, EthClient)
		if err != nil {
			req_status = 0
			err_str = fmt.Sprintf("Can't instantiate CG contract: %v", err)
			prize_time = big.NewInt(0)
		} else {
			pt, err := bwcontract.MainPrizeTime(&copts)
			if err != nil {
				req_status = 0
				err_str = fmt.Sprintf("MainPrizeTime call failed: %v", err)
				prize_time = big.NewInt(0)
			} else {
				prize_time = pt
			}
		}
	}

	c.JSON(http.StatusOK, gin.H{
		"status":            req_status,
		"error":             err_str,
		"CurRoundPrizeTime": prize_time,
	})
}
func api_cosmic_game_user_global_winnings(c *gin.Context) {

	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	if !dbInitialized() {
		common.RespondErrorJSON(c, "Database link wasn't configured")
		return
	}
	p_user_addr := c.Param("user_addr")
	if len(p_user_addr) == 0 {
		common.RespondErrorJSON(c, "'user_addr' parameter is not set")
		return
	}
	user_aid, err := arbStore.LookupAddressID(c.Request.Context(), p_user_addr)
	if err != nil {
		// Address not in DB yet — return 200 with empty winnings so UI works
		c.JSON(http.StatusOK, gin.H{
			"status": 1, "error": "", "Winnings": []interface{}{}, "UserAddr": p_user_addr, "UserAid": int64(0),
		})
		return
	}

	claim_info, err := arbRepo.UserNotifRedBoxRewards(c.Request.Context(), user_aid)
	if err != nil {
		respondStoreError(c, err)
		return
	}

	var req_status int = 1
	var err_str string = ""
	c.JSON(http.StatusOK, gin.H{
		"status":   req_status,
		"error":    err_str,
		"Winnings": claim_info,
		"UserAddr": p_user_addr,
		"UserAid":  user_aid,
	})
}
func api_cosmic_game_prize_history_detail_by_user(c *gin.Context) {

	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	if !dbInitialized() {
		common.RespondErrorJSON(c, "Database link wasn't configured")
		return
	}
	p_user_addr := c.Param("user_addr")
	if len(p_user_addr) == 0 {
		common.RespondErrorJSON(c, "'user_addr' parameter is not set")
		return
	}
	success, offset, limit := common.ParseOffsetLimitParamsJSON(c)
	if !success {
		return
	}
	user_aid, err := arbStore.LookupAddressID(c.Request.Context(), p_user_addr)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"status": 1, "error": "", "UserAddr": p_user_addr, "UserAid": int64(0),
			"UserPrizeHistory": []interface{}{},
		})
		return
	}
	_, err = arbRepo.UserInfo(c.Request.Context(), user_aid)
	if errors.Is(err, store.ErrNotFound) {
		c.JSON(http.StatusOK, gin.H{
			"status": 1, "error": "", "UserAddr": p_user_addr, "UserAid": user_aid,
			"UserPrizeHistory": []interface{}{},
		})
		return
	}
	if err != nil {
		respondStoreError(c, err)
		return
	}

	claim_history, err := arbRepo.PrizeHistoryByUser(c.Request.Context(), user_aid, offset, limit)
	if err != nil {
		respondStoreError(c, err)
		return
	}

	var req_status int = 1
	var err_str string = ""
	c.JSON(http.StatusOK, gin.H{
		"status":           req_status,
		"error":            err_str,
		"UserAddr":         p_user_addr,
		"UserAid":          user_aid,
		"UserPrizeHistory": claim_history,
	})
}
func api_cosmic_game_global_claim_history_detail(c *gin.Context) {

	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	if !dbInitialized() {
		common.RespondErrorJSON(c, "Database link wasn't configured")
		return
	}

	success, offset, limit := common.ParseOffsetLimitParamsJSON(c)
	if !success {
		return
	}

	claim_history, err := arbRepo.ClaimHistoryGlobal(c.Request.Context(), offset, limit)
	if err != nil {
		respondStoreError(c, err)
		return
	}
	var req_status int = 1
	var err_str string = ""
	c.JSON(http.StatusOK, gin.H{
		"status":             req_status,
		"error":              err_str,
		"GlobalPrizeHistory": claim_history,
	})
}
func api_cosmic_game_unclaimed_donated_nfts_by_user(c *gin.Context) {

	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	if !dbInitialized() {
		common.RespondErrorJSON(c, "Database link wasn't configured")
		return
	}
	p_user_addr := c.Param("user_addr")
	if len(p_user_addr) == 0 {
		common.RespondErrorJSON(c, "'user_addr' parameter is not set")
		return
	}
	user_aid, err := arbStore.LookupAddressID(c.Request.Context(), p_user_addr)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"status": 1, "error": "", "UnclaimedDonatedNFTs": []interface{}{},
			"UserAddr": p_user_addr, "UserAid": int64(0),
		})
		return
	}
	_, err = arbRepo.UserInfo(c.Request.Context(), user_aid)
	if errors.Is(err, store.ErrNotFound) {
		c.JSON(http.StatusOK, gin.H{
			"status": 1, "error": "", "UnclaimedDonatedNFTs": []interface{}{},
			"UserAddr": p_user_addr, "UserAid": user_aid,
		})
		return
	}
	if err != nil {
		respondStoreError(c, err)
		return
	}

	nfts, err := arbRepo.UnclaimedDonatedNFTsByUser(c.Request.Context(), user_aid)
	if err != nil {
		respondStoreError(c, err)
		return
	}

	var req_status int = 1
	var err_str string = ""
	c.JSON(http.StatusOK, gin.H{
		"status":               req_status,
		"error":                err_str,
		"UnclaimedDonatedNFTs": nfts,
		"UserAddr":             p_user_addr,
		"UserAid":              user_aid,
	})
}
func api_cosmic_game_unclaimed_donated_nfts_by_prize(c *gin.Context) {

	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	if !dbInitialized() {
		common.RespondErrorJSON(c, "Database link wasn't configured")
		return
	}

	p_prize_num := c.Param("prize_num")
	var prize_num int64
	if len(p_prize_num) > 0 {
		var success bool
		prize_num, success = common.ParseIntFromRemoteOrError(c, JSON, &p_prize_num)
		if !success {
			return
		}
	} else {
		common.RespondErrorJSON(c, "'prize_num' parameter is not set")
		return
	}

	nft_donations, err := arbRepo.UnclaimedDonatedNFTsByRound(c.Request.Context(), prize_num)
	if err != nil {
		respondStoreError(c, err)
		return
	}
	var req_status int = 1
	var err_str string = ""
	c.JSON(http.StatusOK, gin.H{
		"status":       req_status,
		"error":        err_str,
		"NFTDonations": nft_donations,
		"RoundNum":     prize_num,
	})
}
func api_cosmic_game_unclaimed_prize_deposits_by_user(c *gin.Context) {

	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	if !dbInitialized() {
		common.RespondErrorJSON(c, "Database link wasn't configured")
		return
	}
	p_user_addr := c.Param("user_addr")
	if len(p_user_addr) == 0 {
		common.RespondErrorJSON(c, "'user_addr' parameter is not set")
		return
	}
	success, offset, limit := common.ParseOffsetLimitParamsJSON(c)
	if !success {
		return
	}
	user_aid, err := arbStore.LookupAddressID(c.Request.Context(), p_user_addr)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"status": 1, "error": "", "UserAddr": p_user_addr, "UserAid": int64(0),
			"UnclaimedDeposits": []interface{}{},
		})
		return
	}
	_, err = arbRepo.UserInfo(c.Request.Context(), user_aid)
	if errors.Is(err, store.ErrNotFound) {
		c.JSON(http.StatusOK, gin.H{
			"status": 1, "error": "", "UserAddr": p_user_addr, "UserAid": user_aid,
			"UnclaimedDeposits": []interface{}{},
		})
		return
	}
	if err != nil {
		respondStoreError(c, err)
		return
	}

	deposits, err := arbRepo.UnclaimedPrizeEthDeposits(c.Request.Context(), user_aid, offset, limit)
	if err != nil {
		respondStoreError(c, err)
		return
	}

	var req_status int = 1
	var err_str string = ""
	c.JSON(http.StatusOK, gin.H{
		"status":            req_status,
		"error":             err_str,
		"UserAddr":          p_user_addr,
		"UserAid":           user_aid,
		"UnclaimedDeposits": deposits,
	})
}
func api_cosmic_game_cosmic_signature_token_list_by_user(c *gin.Context) {

	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	if !dbInitialized() {
		common.RespondErrorJSON(c, "Database link wasn't configured")
		return
	}
	p_user_addr := c.Param("user_addr")
	if len(p_user_addr) == 0 {
		common.RespondErrorJSON(c, "'user_addr' parameter is not set")
		return
	}
	user_aid, err := arbStore.LookupAddressID(c.Request.Context(), p_user_addr)
	if err != nil {
		// Address not in DB yet — return 200 with empty list so UI works
		success, offset, limit := common.ParseOffsetLimitParamsJSON(c)
		if !success {
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"status": 1, "error": "", "UserAddr": p_user_addr, "UserAid": int64(0),
			"UserTokens": []interface{}{}, "Offset": offset, "Limit": limit,
		})
		return
	}
	_, err = arbRepo.UserInfo(c.Request.Context(), user_aid)
	if errors.Is(err, store.ErrNotFound) {
		success, offset, limit := common.ParseOffsetLimitParamsJSON(c)
		if !success {
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"status": 1, "error": "", "UserAddr": p_user_addr, "UserAid": user_aid,
			"UserTokens": []interface{}{}, "Offset": offset, "Limit": limit,
		})
		return
	}
	if err != nil {
		respondStoreError(c, err)
		return
	}
	success, offset, limit := common.ParseOffsetLimitParamsJSON(c)
	if !success {
		return
	}

	user_tokens, err := arbRepo.CosmicSignatureTokensByUser(c.Request.Context(), user_aid, offset, limit)
	if err != nil {
		respondStoreError(c, err)
		return
	}

	var req_status int = 1
	var err_str string = ""
	c.JSON(http.StatusOK, gin.H{
		"status":     req_status,
		"error":      err_str,
		"UserAddr":   p_user_addr,
		"UserAid":    user_aid,
		"UserTokens": user_tokens,
	})
}
func api_cosmic_game_token_name_history(c *gin.Context) {

	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	if !dbInitialized() {
		common.RespondErrorJSON(c, "Database link wasn't configured")
		return
	}

	p_token_id := c.Param("token_id")
	var token_id int64
	if len(p_token_id) > 0 {
		var success bool
		token_id, success = common.ParseIntFromRemoteOrError(c, JSON, &p_token_id)
		if !success {
			return
		}
	} else {
		common.RespondErrorJSON(c, "'token_id' parameter is not set")
		return
	}

	var req_status int = 1
	var err_str string = ""

	tokname_history, err := arbRepo.TokenNameHistory(c.Request.Context(), token_id)
	if err != nil {
		respondStoreError(c, err)
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"status":           req_status,
		"error":            err_str,
		"TokenId":          token_id,
		"TokenNameHistory": tokname_history,
	})
}
func api_cosmic_game_token_name_search(c *gin.Context) {

	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	if !dbInitialized() {
		common.RespondErrorJSON(c, "Database link wasn't configured")
		return
	}

	p_name := c.Param("name")
	if len(p_name) > 0 {
	} else {
		common.RespondErrorJSON(c, "'search_text' parameter is not set")
		return
	}

	results, err := arbRepo.SearchTokensByName(c.Request.Context(), p_name)
	if err != nil {
		respondStoreError(c, err)
		return
	}
	var req_status int = 1
	var err_str string = ""

	c.JSON(http.StatusOK, gin.H{
		"status":                 req_status,
		"error":                  err_str,
		"SearchText":             p_name,
		"TokenNameSearchResults": results,
	})
}
func api_cosmic_game_named_tokens_only(c *gin.Context) {

	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	if !dbInitialized() {
		common.RespondErrorJSON(c, "Database link wasn't configured")
		return
	}

	var req_status int = 1
	var err_str string = ""

	results, err := arbRepo.NamedTokens(c.Request.Context())
	if err != nil {
		respondStoreError(c, err)
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"status":      req_status,
		"error":       err_str,
		"NamedTokens": results,
	})
}
func api_cosmic_game_token_ownership_transfers(c *gin.Context) {

	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	if !dbInitialized() {
		common.RespondErrorJSON(c, "Database link wasn't configured")
		return
	}

	p_token_id := c.Param("token_id")
	var token_id int64
	if len(p_token_id) > 0 {
		var success bool
		token_id, success = common.ParseIntFromRemoteOrError(c, JSON, &p_token_id)
		if !success {
			return
		}
	} else {
		common.RespondErrorJSON(c, "'token_id' parameter is not set")
		return
	}
	success, offset, limit := common.ParseOffsetLimitParamsJSON(c)
	if !success {
		return
	}

	var req_status int = 1
	var err_str string = ""

	transfers, err := arbRepo.TokenOwnershipTransfers(c.Request.Context(), token_id, offset, limit)
	if err != nil {
		respondStoreError(c, err)
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"status":         req_status,
		"error":          err_str,
		"Offset":         offset,
		"Limit":          limit,
		"TokenId":        token_id,
		"TokenTransfers": transfers,
	})
}
func api_cosmic_game_cs_token_distribution(c *gin.Context) {
	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")

	if !dbInitialized() {
		common.RespondErrorJSON(c, "Database link wasn't configured")
		return
	}
	distribution, err := arbRepo.CosmicSignatureTokenDistribution(c.Request.Context())
	if err != nil {
		respondStoreError(c, err)
		return
	}

	var req_status int = 1
	var err_str string = ""

	c.JSON(http.StatusOK, gin.H{
		"status":                           req_status,
		"error":                            err_str,
		"CosmicSignatureTokenDistribution": distribution,
	})
}
func api_cosmic_game_user_balances(c *gin.Context) {

	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	if !dbInitialized() {
		common.RespondErrorJSON(c, "Database link wasn't configured")
		return
	}

	p_user_addr := c.Param("user_addr")
	if len(p_user_addr) == 0 {
		common.RespondErrorJSON(c, "'user_addr' parameter is not set")
		return
	}
	user_aid, addrLookupErr := arbStore.LookupAddressID(c.Request.Context(), p_user_addr)
	if addrLookupErr != nil {
		user_aid = 0 // address not in DB yet; still return 200 with on-chain balances below
	}

	addr := ethcommon.HexToAddress(p_user_addr)
	user_eth_bal, err := EthClient.BalanceAt(context.Background(), addr, nil)
	if err != nil {
		err_str := fmt.Sprintf("Error at BalanceAt() call for addr: %v\n", err)
		common.RespondErrorJSON(c, err_str)
		return
	}
	ct_contract, err := NewERC20(contractState.Snapshot().Addrs.CosmicToken, EthClient)
	if err != nil {
		err_str := fmt.Sprintf("Error at instantiation of ERC20 contract: %v\n", err)
		common.RespondErrorJSON(c, err_str)
		return
	}
	var copts bind.CallOpts
	ct_balance, err := ct_contract.BalanceOf(&copts, addr)
	if err != nil {
		err_str := fmt.Sprintf("Error at BalanceOf() call: %v\n", err)
		common.RespondErrorJSON(c, err_str)
		return
	}

	var req_status int = 1
	var err_str string = ""

	c.JSON(http.StatusOK, gin.H{
		"status":             req_status,
		"error":              err_str,
		"UserAddr":           p_user_addr,
		"UserAid":            user_aid,
		"ETH_Balance":        user_eth_bal.String(),
		"CosmicTokenBalance": ct_balance.String(),
	})
}
func api_cosmic_game_cosmic_token_balances(c *gin.Context) {

	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	if !dbInitialized() {
		common.RespondErrorJSON(c, "Database link wasn't configured")
		return
	}

	balances, err := arbRepo.CosmicTokenHolders(c.Request.Context())
	if err != nil {
		respondStoreError(c, err)
		return
	}
	var req_status int = 1
	var err_str string = ""

	c.JSON(http.StatusOK, gin.H{
		"status":              req_status,
		"error":               err_str,
		"CosmicTokenBalances": balances,
	})
}
func api_cosmic_game_cosmic_token_statistics(c *gin.Context) {

	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	if !dbInitialized() {
		common.RespondErrorJSON(c, "Database link wasn't configured")
		return
	}

	stats, err := arbRepo.CosmicTokenStatistics(c.Request.Context())
	if err != nil {
		respondStoreError(c, err)
		return
	}

	cosmicTokenAddr := contractState.Snapshot().Addrs.CosmicToken

	// Read contract info from blockchain
	var copts bind.CallOpts
	contract, err := NewCosmicSignatureToken(cosmicTokenAddr, EthClient)
	if err != nil {
		err_str := fmt.Sprintf("Can't instantiate CosmicSignatureToken contract: %v", err)
		Error.Print(err_str)
		Info.Print(err_str)
		common.RespondErrorJSON(c, err_str)
		return
	}

	token_name, err := contract.Name(&copts)
	if err != nil {
		err_str := fmt.Sprintf("Error reading token name: %v", err)
		Error.Print(err_str)
		Info.Print(err_str)
		common.RespondErrorJSON(c, err_str)
		return
	}

	token_symbol, err := contract.Symbol(&copts)
	if err != nil {
		err_str := fmt.Sprintf("Error reading token symbol: %v", err)
		Error.Print(err_str)
		Info.Print(err_str)
		common.RespondErrorJSON(c, err_str)
		return
	}

	decimals, err := contract.Decimals(&copts)
	if err != nil {
		err_str := fmt.Sprintf("Error reading decimals: %v", err)
		Error.Print(err_str)
		Info.Print(err_str)
		common.RespondErrorJSON(c, err_str)
		return
	}

	game_addr, err := contract.Game(&copts)
	if err != nil {
		err_str := fmt.Sprintf("Error reading game address: %v", err)
		Error.Print(err_str)
		Info.Print(err_str)
		common.RespondErrorJSON(c, err_str)
		return
	}

	var req_status int = 1
	var err_str string = ""
	c.JSON(http.StatusOK, gin.H{
		"status":           req_status,
		"error":            err_str,
		"TokenName":        token_name,
		"TokenSymbol":      token_symbol,
		"TokenDecimals":    decimals,
		"GameContractAddr": game_addr.String(),
		"CosmicTokenAddr":  cosmicTokenAddr.String(),
		"Statistics":       stats,
	})
}
func api_cosmic_game_cosmic_token_summary_by_user(c *gin.Context) {

	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	if !dbInitialized() {
		common.RespondErrorJSON(c, "Database link wasn't configured")
		return
	}
	p_user_addr := c.Param("user_addr")
	if len(p_user_addr) == 0 {
		common.RespondErrorJSON(c, "'user_addr' parameter is not set")
		return
	}
	user_aid, err := arbStore.LookupAddressID(c.Request.Context(), p_user_addr)
	if err != nil {
		common.RespondErrorJSON(c, "Provided address wasn't found")
		return
	}

	summary, err := arbRepo.UserCosmicTokenSummary(c.Request.Context(), user_aid)
	if err != nil {
		respondStoreError(c, err)
		return
	}
	summary.UserAddr = p_user_addr

	var req_status int = 1
	var err_str string = ""
	c.JSON(http.StatusOK, gin.H{
		"status":  req_status,
		"error":   err_str,
		"Summary": summary,
	})
}
func parseCosmicTokenHistoryDateParam(c *gin.Context, paramName string) (string, bool) {
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

func api_cosmic_game_cosmic_token_total_supply_history_by_bid(c *gin.Context) {

	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	if !dbInitialized() {
		common.RespondErrorJSON(c, "Database link wasn't configured")
		return
	}

	history, err := arbRepo.CosmicTokenSupplyHistoryByBid(c.Request.Context())
	if err != nil {
		respondStoreError(c, err)
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"status":             1,
		"error":              "",
		"TotalSupplyHistory": history,
	})
}
func api_cosmic_game_cosmic_token_total_supply_history_by_date(c *gin.Context) {

	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	if !dbInitialized() {
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

	history, err := arbRepo.CosmicTokenSupplyHistoryByDate(c.Request.Context(), fromDate, toDate)
	if err != nil {
		respondStoreError(c, err)
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"status":             1,
		"error":              "",
		"DateFrom":           fromDate,
		"DateTo":             toDate,
		"TotalSupplyHistory": history,
	})
}
func api_cosmic_game_cosmic_token_transfers_by_user(c *gin.Context) {

	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	if !dbInitialized() {
		common.RespondErrorJSON(c, "Database link wasn't configured")
		return
	}
	p_user_addr := c.Param("user_addr")
	if len(p_user_addr) == 0 {
		common.RespondErrorJSON(c, "'user_addr' parameter is not set")
		return
	}
	user_aid, err := arbStore.LookupAddressID(c.Request.Context(), p_user_addr)
	if err != nil {
		common.RespondErrorJSON(c, "Provided address wasn't found")
		return
	}
	success, offset, limit := common.ParseOffsetLimitParamsJSON(c)
	if !success {
		return
	}
	transfers, err := arbRepo.CosmicTokenTransfersByUser(c.Request.Context(), user_aid, offset, limit)
	if err != nil {
		respondStoreError(c, err)
		return
	}
	var req_status int = 1
	var err_str string = ""

	c.JSON(http.StatusOK, gin.H{
		"status":               req_status,
		"error":                err_str,
		"UserAddr":             p_user_addr,
		"UserAid":              user_aid,
		"Offset":               offset,
		"Limit":                limit,
		"CosmicTokenTransfers": transfers,
	})
}
func api_cosmic_game_cosmic_signature_transfers_by_user(c *gin.Context) {

	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	if !dbInitialized() {
		common.RespondErrorJSON(c, "Database link wasn't configured")
		return
	}
	p_user_addr := c.Param("user_addr")
	if len(p_user_addr) == 0 {
		common.RespondErrorJSON(c, "'user_addr' parameter is not set")
		return
	}
	user_aid, err := arbStore.LookupAddressID(c.Request.Context(), p_user_addr)
	if err != nil {
		common.RespondErrorJSON(c, "Provided address wasn't found")
		return
	}
	success, offset, limit := common.ParseOffsetLimitParamsJSON(c)
	if !success {
		return
	}
	transfers, err := arbRepo.CosmicSignatureTransfersByUser(c.Request.Context(), user_aid, offset, limit)
	if err != nil {
		respondStoreError(c, err)
		return
	}
	var req_status int = 1
	var err_str string = ""

	c.JSON(http.StatusOK, gin.H{
		"status":                   req_status,
		"error":                    err_str,
		"UserAddr":                 p_user_addr,
		"UserAid":                  user_aid,
		"Offset":                   offset,
		"Limit":                    limit,
		"CosmicSignatureTransfers": transfers,
	})
}
func api_cosmic_game_used_rwalk_nfts(c *gin.Context) {

	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	if !dbInitialized() {
		common.RespondErrorJSON(c, "Database link wasn't configured")
		return
	}

	used_nfts, err := arbRepo.RandomWalkTokensUsedInBids(c.Request.Context())
	if err != nil {
		respondStoreError(c, err)
		return
	}
	var req_status int = 1
	var err_str string = ""

	c.JSON(http.StatusOK, gin.H{
		"status":        req_status,
		"error":         err_str,
		"UsedRwalkNFTs": used_nfts,
	})
}
func api_cosmic_game_marketing_rewards_global(c *gin.Context) {
	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	if !dbInitialized() {
		common.RespondErrorJSON(c, "Database link wasn't configured")
		return
	}
	success, offset, limit := common.ParseOffsetLimitParamsJSON(c)
	if !success {
		return
	}
	rewards, err := arbRepo.MarketingRewardHistoryGlobal(c.Request.Context(), offset, limit)
	if err != nil {
		respondStoreError(c, err)
		return
	}
	var req_status int = 1
	var err_str string = ""
	c.JSON(http.StatusOK, gin.H{
		"status":           req_status,
		"error":            err_str,
		"Offset":           offset,
		"Limit":            limit,
		"MarketingRewards": rewards,
	})
}
func api_cosmic_game_marketing_rewards_by_user(c *gin.Context) {
	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	if !dbInitialized() {
		common.RespondErrorJSON(c, "Database link wasn't configured")
		return
	}
	p_user_addr := c.Param("user_addr")
	if len(p_user_addr) == 0 {
		common.RespondErrorJSON(c, "'user_addr' parameter is not set")
		return
	}
	user_aid, err := arbStore.LookupAddressID(c.Request.Context(), p_user_addr)
	if err != nil {
		// Address not in DB yet — return 200 with empty list so UI works
		success, offset, limit := common.ParseOffsetLimitParamsJSON(c)
		if !success {
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"status": 1, "error": "", "Offset": offset, "Limit": limit,
			"UserAddr": p_user_addr, "UserAid": int64(0), "UserMarketingRewards": []interface{}{},
		})
		return
	}
	success, offset, limit := common.ParseOffsetLimitParamsJSON(c)
	if !success {
		return
	}
	rewards, err := arbRepo.MarketingRewardHistoryByUser(c.Request.Context(), user_aid, offset, limit)
	if err != nil {
		respondStoreError(c, err)
		return
	}
	var req_status int = 1
	var err_str string = ""
	c.JSON(http.StatusOK, gin.H{
		"status":               req_status,
		"error":                err_str,
		"Offset":               offset,
		"Limit":                limit,
		"UserAddr":             p_user_addr,
		"UserAid":              user_aid,
		"UserMarketingRewards": rewards,
	})
}
func api_cosmic_game_marketing_config_current(c *gin.Context) {
	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	if !dbInitialized() {
		common.RespondErrorJSON(c, "Database link wasn't configured")
		return
	}

	marketingWalletAddr := contractState.Snapshot().Addrs.MarketingWallet

	var copts bind.CallOpts
	contract, err := NewMarketingWallet(marketingWalletAddr, EthClient)
	if err != nil {
		err_str := fmt.Sprintf("Can't instantiate MarketingWallet contract: %v", err)
		Error.Print(err_str)
		Info.Print(err_str)
		common.RespondErrorJSON(c, err_str)
		return
	}

	// Read current treasurer address
	treasurer_addr, err := contract.TreasurerAddress(&copts)
	if err != nil {
		err_str := fmt.Sprintf("Error reading TreasurerAddress: %v", err)
		Error.Print(err_str)
		Info.Print(err_str)
		common.RespondErrorJSON(c, err_str)
		return
	}

	// Read token contract address
	token_addr, err := contract.Token(&copts)
	if err != nil {
		err_str := fmt.Sprintf("Error reading Token address: %v", err)
		Error.Print(err_str)
		Info.Print(err_str)
		common.RespondErrorJSON(c, err_str)
		return
	}

	// Read owner address
	owner_addr, err := contract.Owner(&copts)
	if err != nil {
		err_str := fmt.Sprintf("Error reading Owner address: %v", err)
		Error.Print(err_str)
		Info.Print(err_str)
		common.RespondErrorJSON(c, err_str)
		return
	}

	var req_status int = 1
	var err_str string = ""
	c.JSON(http.StatusOK, gin.H{
		"status":              req_status,
		"error":               err_str,
		"MarketingWalletAddr": marketingWalletAddr.String(),
		"TreasurerAddr":       treasurer_addr.String(),
		"TokenAddr":           token_addr.String(),
		"OwnerAddr":           owner_addr.String(),
	})
}
func api_cosmic_game_get_cst_price(c *gin.Context) {

	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	var copts bind.CallOpts
	contract, err := NewCosmicSignatureGame(contractState.Snapshot().Addrs.CosmicGame, EthClient)
	if err != nil {
		err_str := fmt.Sprintf("Can't instantiate CosmicGame contract: %v . Contract constants won't be fetched\n", err)
		Error.Print(err_str)
		Info.Print(err_str)
		common.RespondErrorJSON(c, err_str)
	} else {
		cst_price, err := contract.GetNextCstBidPrice(&copts)
		if err != nil {
			Error.Print(err.Error())
			Info.Print(err.Error())
			common.RespondError(c, err.Error())
		} else {
			auction_duration, seconds_elapsed, err := contract.GetCstDutchAuctionDurations(&copts)
			if err != nil {
				Error.Print(err.Error())
				Info.Print(err.Error())
				common.RespondError(c, err.Error())
			} else {
				var req_status int = 1
				var err_str string = ""
				c.JSON(http.StatusOK, gin.H{
					"status":          req_status,
					"error":           err_str,
					"CSTPrice":        cst_price.String(),
					"SecondsElapsed":  seconds_elapsed.String(),
					"AuctionDuration": auction_duration.String(),
				})
			}
		}
	}
}
func api_cosmic_game_get_eth_price(c *gin.Context) {

	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	var copts bind.CallOpts
	contract, err := NewCosmicSignatureGame(contractState.Snapshot().Addrs.CosmicGame, EthClient)
	if err != nil {
		err_str := fmt.Sprintf("Can't instantiate CosmicGame contract: %v . Contract constants won't be fetched\n", err)
		Error.Print(err_str)
		Info.Print(err_str)
		common.RespondErrorJSON(c, err_str)
	} else {
		eth_price, err := contract.GetNextEthBidPrice(&copts)
		if err != nil {
			Error.Print(err.Error())
			Info.Print(err.Error())
			common.RespondError(c, err.Error())
		} else {
			auction_duration, seconds_elapsed, err := contract.GetEthDutchAuctionDurations(&copts)
			if err != nil {
				Error.Print(err.Error())
				Info.Print(err.Error())
				common.RespondError(c, err.Error())
			} else {
				var req_status int = 1
				var err_str string = ""
				c.JSON(http.StatusOK, gin.H{
					"status":          req_status,
					"error":           err_str,
					"ETHPrice":        eth_price.String(),
					"SecondsElapsed":  seconds_elapsed.String(),
					"AuctionDuration": auction_duration.String(),
				})
			}
		}
	}
}
func api_cosmic_game_sysmode_changes(c *gin.Context) {

	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	if !dbInitialized() {
		common.RespondErrorJSON(c, "Database link wasn't configured")
		return
	}

	var offset, limit int
	p_offset := c.Param("offset")
	p_limit := c.Param("limit")

	if len(p_offset) == 0 || len(p_limit) == 0 {
		offset = 0
		limit = 10000
	} else {
		var success bool
		success, offset, limit = common.ParseOffsetLimitParamsJSON(c)
		if !success {
			return
		}
	}

	system_mode_changes, err := arbRepo.SystemModeChanges(c.Request.Context(), offset, limit)
	if err != nil {
		respondStoreError(c, err)
		return
	}

	// When offset=-1 (events from deployment), ensure at least the pre-round-0 row is shown even if no bids yet.
	if offset == -1 && len(system_mode_changes) == 0 {
		system_mode_changes = []p.CGSystemModeRec{{
			EvtLogId:     -1,
			BlockNum:     -1,
			RoundNum:     0,
			NextEvtLogId: math.MaxInt64,
		}}
	}

	var req_status int = 1
	var err_str string = ""
	c.JSON(http.StatusOK, gin.H{
		"status":            req_status,
		"error":             err_str,
		"Offset":            offset,
		"Limit":             limit,
		"SystemModeChanges": system_mode_changes,
	})
}
func api_cosmic_game_admin_events_in_range(c *gin.Context) {

	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	if !dbInitialized() {
		common.RespondErrorJSON(c, "Database link wasn't configured")
		return
	}

	p_evtlog_start := c.Param("evtlog_start")
	var evtlog_start int64
	if len(p_evtlog_start) > 0 {
		var success bool
		evtlog_start, success = common.ParseIntFromRemoteOrError(c, JSON, &p_evtlog_start)
		if !success {
			return
		}
	} else {
		common.RespondErrorJSON(c, "'evtlog_start' parameter is not set")
		return
	}
	p_evtlog_end := c.Param("evtlog_end")
	var evtlog_end int64
	if len(p_evtlog_end) > 0 {
		var success bool
		evtlog_end, success = common.ParseIntFromRemoteOrError(c, JSON, &p_evtlog_end)
		if !success {
			return
		}
	} else {
		common.RespondErrorJSON(c, "'evtlog_end' parameter is not set")
		return
	}
	event_list, err := arbRepo.AdminEventsInRange(c.Request.Context(), evtlog_start, evtlog_end)
	if err != nil {
		respondStoreError(c, err)
		return
	}
	if err := arbRepo.ResolveAdminEventValues(c.Request.Context(), event_list); err != nil {
		respondStoreError(c, err)
		return
	}
	enrichAdminEventsResolvedValues(event_list)

	var req_status int = 1
	var err_str string = ""
	c.JSON(http.StatusOK, gin.H{
		"status":        req_status,
		"error":         err_str,
		"AdminEvents":   event_list,
		"EvtLogIdStart": evtlog_start,
		"EvtLogIdEnd":   evtlog_end,
	})
}
func api_cosmic_game_bid_special_winners(c *gin.Context) {

	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	if !dbInitialized() {
		common.RespondErrorJSON(c, "Database link wasn't configured")
		return
	}
	if EthClient == nil {
		common.RespondErrorJSON(c, "Ethereum client is not configured")
		return
	}

	state := contractState.FetchLiveSpecialWinners(c.Request.Context())
	if state.Err != nil {
		common.RespondErrorJSON(c, state.Err.Error())
		return
	}

	resp := gin.H{
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
	if state.HasLastCstBidEventLogId {
		resp["LastCstBidEventLogId"] = state.LastCstBidEventLogId
	}
	c.JSON(http.StatusOK, resp)
}

// =============================================================================
// BANNED BIDS API (mirrors FastAPI get_banned_bids / ban_bid / unban_bid)
// =============================================================================

func api_cosmic_game_get_banned_bids(c *gin.Context) {
	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	if !dbInitialized() {
		common.RespondErrorJSON(c, "Database link wasn't configured")
		return
	}
	list, err := arbRepo.BannedBids(c.Request.Context())
	if err != nil {
		respondStoreError(c, err)
		return
	}
	// Return raw JSON array like FastAPI so frontend gets same shape
	c.JSON(http.StatusOK, list)
}

type banBidPayload struct {
	BidId    int64  `json:"bid_id" binding:"required"`
	UserAddr string `json:"user_addr" binding:"required"`
}

func api_cosmic_game_ban_bid(c *gin.Context) {
	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	if !dbInitialized() {
		common.RespondErrorJSON(c, "Database link wasn't configured")
		return
	}
	var payload banBidPayload
	if err := c.ShouldBindJSON(&payload); err != nil {
		common.RespondErrorJSON(c, "Invalid JSON: bid_id and user_addr required")
		return
	}
	if err := arbRepo.InsertBannedBid(c.Request.Context(), payload.BidId, payload.UserAddr); err != nil {
		common.RespondErrorJSON(c, fmt.Sprintf("Failed to insert banned bid: %v", err))
		return
	}
	c.JSON(http.StatusCreated, gin.H{"result": "success"})
}

type unbanBidPayload struct {
	BidId int64 `json:"bid_id" binding:"required"`
}

func api_cosmic_game_unban_bid(c *gin.Context) {
	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	if !dbInitialized() {
		common.RespondErrorJSON(c, "Database link wasn't configured")
		return
	}
	var payload unbanBidPayload
	if err := c.ShouldBindJSON(&payload); err != nil {
		common.RespondErrorJSON(c, "Invalid JSON: bid_id required")
		return
	}
	if err := arbRepo.DeleteBannedBid(c.Request.Context(), payload.BidId); err != nil {
		common.RespondErrorJSON(c, fmt.Sprintf("Failed to unban bid: %v", err))
		return
	}
	c.JSON(http.StatusCreated, gin.H{"result": "success"})
}
