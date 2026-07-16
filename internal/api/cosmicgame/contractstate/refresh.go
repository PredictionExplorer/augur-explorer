package contractstate

import (
	"context"
	"errors"
	"fmt"
	"math"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	ethcommon "github.com/ethereum/go-ethereum/common"

	cg "github.com/PredictionExplorer/augur-explorer/contracts/cosmicgame"
)

var weiPerEth = new(big.Float).SetInt(big.NewInt(1e18))

// weiToEth converts a wei amount to a float64 ETH value the way the legacy
// code did (big.Float quotient, rounding ignored).
func weiToEth(wei *big.Int) float64 {
	quo := new(big.Float).Quo(new(big.Float).SetInt(wei), weiPerEth)
	out, _ := quo.Float64()
	return out
}

func (s *State) latestCallOpts(ctx context.Context) (bind.CallOpts, int64, error) {
	header, err := s.client.HeaderByNumber(ctx, nil)
	if err != nil {
		return bind.CallOpts{}, 0, err
	}
	if header.Number == nil || header.Time > math.MaxInt64 {
		return bind.CallOpts{}, 0, errors.New("latest block header is invalid")
	}
	return bind.CallOpts{
		Context:     ctx,
		BlockNumber: new(big.Int).Set(header.Number),
	}, int64(header.Time), nil // #nosec G115 -- checked above
}

func normalizeAuctionProgress(
	durationValue *big.Int,
	elapsedOrStartValue *big.Int,
	blockTimestamp int64,
	secondValueMayBeStart bool,
) (duration int64, elapsed int64, ok bool) {
	if durationValue == nil || elapsedOrStartValue == nil ||
		!durationValue.IsInt64() || !elapsedOrStartValue.IsInt64() {
		return 0, 0, false
	}
	duration = durationValue.Int64()
	raw := elapsedOrStartValue.Int64()
	if duration < 0 {
		return 0, 0, false
	}
	elapsed = raw
	if secondValueMayBeStart && raw > duration && raw > 1_000_000_000 {
		elapsed = blockTimestamp - raw
	}
	if elapsed < 0 {
		elapsed = 0
	}
	if elapsed > duration {
		elapsed = duration
	}
	return duration, elapsed, true
}

func nonNegativeInt64(value *big.Int) (int64, bool) {
	if value == nil || !value.IsInt64() || value.Sign() < 0 {
		return 0, false
	}
	return value.Int64(), true
}

func (s *State) markConstantsUnavailable() {
	s.mu.Lock()
	s.snap.ConstantsReady = false
	s.snap.ConfigurationReady = false
	s.mu.Unlock()
}

func (s *State) markVariablesUnavailable() {
	s.mu.Lock()
	s.snap.BidPrice = "error"
	s.snap.BlockPinnedBidPrice = "error"
	s.snap.NextCSTBidPrice = "error"
	s.snap.TokenReward = "error"
	s.snap.NextCSTBidReward = "error"
	s.snap.ETHAuctionDuration = -1
	s.snap.ETHAuctionElapsed = -1
	s.snap.CSTAuctionDuration = -1
	s.snap.CSTAuctionElapsed = -1
	s.snap.BidPricesReady = false
	s.snap.PrizeClaimTimestamp = -1
	s.snap.PrizeAmount = "error"
	s.snap.RaffleAmount = "error"
	s.snap.StakingAmount = "error"
	s.snap.RoundNum = -1
	s.snap.MainPrizeTimeIncrement = "error"
	s.snap.InitialSecondsUntilPrize = -1
	s.snap.TimeoutClaimPrize = -1
	s.snap.RoundStartAuctionLength = -1
	s.snap.VariablesMechanicsVersion = mechanicsUnknown
	s.snap.ConfigurationReady = false
	s.mu.Unlock()
}

// refreshConstants reloads the owner-tunable contract parameters. Each field
// keeps its legacy failure sentinel ("error" strings, -1 counters, 0 charity
// percentage, previous charity address), so a dead RPC node degrades the
// dashboard exactly as before.
func (s *State) refreshConstants(ctx context.Context) {
	s.contractRefreshMu.Lock()
	defer s.contractRefreshMu.Unlock()
	ctx, cancel := context.WithTimeout(ctx, s.rpcReadTimeout)
	defer cancel()

	copts, _, err := s.latestCallOpts(ctx)
	if err != nil {
		s.logf("Error fetching block for contract constants: %v\n", err)
		s.markConstantsUnavailable()
		return
	}

	code, err := s.client.CodeAt(ctx, s.addrs.CosmicGame, copts.BlockNumber)
	if err != nil {
		s.logf("Can't instantiate CosmicGame contract: %v\n", err)
	} else if len(code) == 0 {
		s.logf("Can't instantiate CosmicGame contract: no code at given address\n")
	}

	v1, v2 := s.bindLiveReaders()
	if v1 == nil {
		s.logf("Can't instantiate CosmicGame contract at %v . Contract constants won't be fetched\n", s.addrs.CosmicGame)
		s.mu.Lock()
		s.snap.ConstantsReady = false
		s.snap.ConfigurationReady = false
		s.mu.Unlock()
		return
	}

	cur := s.Snapshot()
	ready := true
	charityAddressReady := true

	if val, err := v1.EthBidPriceIncreaseDivisor(&copts); err != nil {
		s.logf("Error at PriceIncrease() call: %v\n", err)
		cur.PriceIncrease = "error"
		ready = false
	} else {
		cur.PriceIncrease = val.String()
	}
	if addr, err := v1.CharityAddress(&copts); err != nil {
		// Keeps the previous address on failure (legacy behavior).
		s.logf("Error at Charity() call: %v\n", err)
		ready = false
		charityAddressReady = false
	} else {
		cur.CharityAddr = addr
	}
	if val, err := v1.CharityEthDonationAmountPercentage(&copts); err != nil {
		s.logf("Error at Charity() call: %v\n", err)
		cur.CharityPercentage = 0
		ready = false
	} else if parsed, ok := nonNegativeInt64(val); !ok {
		s.logf("Error at Charity() call: value exceeds int64\n")
		cur.CharityPercentage = 0
		ready = false
	} else {
		cur.CharityPercentage = parsed
	}
	if fixedReward, multiplier, err := s.bidCSTRewardConfiguration(v1, v2, &copts); err != nil {
		s.logf("Error at CST bid reward configuration call: %v\n", err)
		cur.TokenReward = "error"
		ready = false
	} else {
		cur.FixedCSTBidReward = fixedReward
		if fixedReward != "" {
			cur.TokenReward = fixedReward
			cur.NextCSTBidReward = fixedReward
		}
		cur.BidCSTRewardMultiplier = multiplier
	}
	if val, err := v1.MainEthPrizeAmountPercentage(&copts); err != nil {
		s.logf("Error at PrizePercentage() call: %v\n", err)
		cur.PrizePercentage = -1
		ready = false
	} else if parsed, ok := nonNegativeInt64(val); !ok {
		s.logf("Error at PrizePercentage() call: value exceeds int64\n")
		cur.PrizePercentage = -1
		ready = false
	} else {
		cur.PrizePercentage = parsed
	}
	if val, err := v1.RaffleTotalEthPrizeAmountForBiddersPercentage(&copts); err != nil {
		s.logf("Error at RafflePercentage() call: %v\n", err)
		cur.RafflePercentage = -1
		ready = false
	} else if parsed, ok := nonNegativeInt64(val); !ok {
		s.logf("Error at RafflePercentage() call: value exceeds int64\n")
		cur.RafflePercentage = -1
		ready = false
	} else {
		cur.RafflePercentage = parsed
	}
	if val, err := v1.ChronoWarriorEthPrizeAmountPercentage(&copts); err != nil {
		s.logf("Error at ChronoWarriorEthPrizeAmountPercentage() call: %v\n", err)
		cur.ChronoPercentage = -1
		ready = false
	} else if parsed, ok := nonNegativeInt64(val); !ok {
		s.logf("Error at ChronoWarriorEthPrizeAmountPercentage() call: value exceeds int64\n")
		cur.ChronoPercentage = -1
		ready = false
	} else {
		cur.ChronoPercentage = parsed
	}
	if val, err := v1.CosmicSignatureNftStakingTotalEthRewardAmountPercentage(&copts); err != nil {
		s.logf("Error at StakingPercentage() call: %v\n", err)
		cur.StakingPercentage = -1
		ready = false
	} else if parsed, ok := nonNegativeInt64(val); !ok {
		s.logf("Error at StakingPercentage() call: value exceeds int64\n")
		cur.StakingPercentage = -1
		ready = false
	} else {
		cur.StakingPercentage = parsed
	}
	if val, err := v1.MainPrizeTimeIncrementIncreaseDivisor(&copts); err != nil {
		s.logf("Error at TimeIncrease() call: %v\n", err)
		cur.TimeIncrease = "error"
		ready = false
	} else {
		cur.TimeIncrease = val.String()
	}
	if val, err := v1.NumRaffleEthPrizesForBidders(&copts); err != nil {
		s.logf("Error at NumRaffleETHWinnersBidding() call: %v\n", err)
		cur.RaffleEthWinnersBidding = -1
		ready = false
	} else if parsed, ok := nonNegativeInt64(val); !ok {
		s.logf("Error at NumRaffleETHWinnersBidding() call: value exceeds int64\n")
		cur.RaffleEthWinnersBidding = -1
		ready = false
	} else {
		cur.RaffleEthWinnersBidding = parsed
	}
	if val, err := v1.NumRaffleCosmicSignatureNftsForBidders(&copts); err != nil {
		s.logf("Error at NumRaffleNFTWinnersBidding() call: %v\n", err)
		cur.RaffleNFTWinnersBidding = -1
		ready = false
	} else if parsed, ok := nonNegativeInt64(val); !ok {
		s.logf("Error at NumRaffleNFTWinnersBidding() call: value exceeds int64\n")
		cur.RaffleNFTWinnersBidding = -1
		ready = false
	} else {
		cur.RaffleNFTWinnersBidding = parsed
	}
	if val, err := v1.NumRaffleCosmicSignatureNftsForRandomWalkNftStakers(&copts); err != nil {
		s.logf("Error at NumRaffleNFTWinnersStakingRWalk() call: %v\n", err)
		cur.RaffleNFTWinnersStakingRWalk = -1
		ready = false
	} else if parsed, ok := nonNegativeInt64(val); !ok {
		s.logf("Error at NumRaffleNFTWinnersStakingRWalk() call: value exceeds int64\n")
		cur.RaffleNFTWinnersStakingRWalk = -1
		ready = false
	} else {
		cur.RaffleNFTWinnersStakingRWalk = parsed
	}
	var divisorReady bool
	cur.CSTAuctionDurationChangeDivisor, divisorReady = s.cstAuctionDurationChangeDivisor(v1, v2, &copts)
	if !divisorReady {
		s.logf("Error reading CST auction duration-change divisor\n")
		ready = false
	}
	// The treasurer is a MarketingWallet constant; it refreshes with this
	// group and keeps the previous address on failure like the charity
	// address does.
	marketing, _ := cg.NewMarketingWallet(s.addrs.MarketingWallet, s.client)
	if addr, err := marketing.TreasurerAddress(&copts); err != nil {
		s.logf("Error at TreasurerAddress() call: %v\n", err)
		ready = false
	} else {
		cur.TreasurerAddr = addr
	}
	cur.MechanicsVersion = s.mechanicsVersion()

	s.mu.Lock()
	charityAddressChanged := s.snap.CharityAddr != cur.CharityAddr
	s.snap.PriceIncrease = cur.PriceIncrease
	s.snap.CharityAddr = cur.CharityAddr
	s.snap.CharityPercentage = cur.CharityPercentage
	s.snap.TokenReward = cur.TokenReward
	s.snap.FixedCSTBidReward = cur.FixedCSTBidReward
	s.snap.BidCSTRewardMultiplier = cur.BidCSTRewardMultiplier
	s.snap.PrizePercentage = cur.PrizePercentage
	s.snap.RafflePercentage = cur.RafflePercentage
	s.snap.ChronoPercentage = cur.ChronoPercentage
	s.snap.StakingPercentage = cur.StakingPercentage
	s.snap.TimeIncrease = cur.TimeIncrease
	s.snap.RaffleEthWinnersBidding = cur.RaffleEthWinnersBidding
	s.snap.RaffleNFTWinnersBidding = cur.RaffleNFTWinnersBidding
	s.snap.RaffleNFTWinnersStakingRWalk = cur.RaffleNFTWinnersStakingRWalk
	s.snap.CSTAuctionDurationChangeDivisor = cur.CSTAuctionDurationChangeDivisor
	s.snap.TreasurerAddr = cur.TreasurerAddr
	s.snap.MechanicsVersion = cur.MechanicsVersion
	s.snap.ConstantsMechanicsVersion = cur.MechanicsVersion
	s.snap.ConstantsReady = ready
	if !charityAddressReady || charityAddressChanged {
		s.snap.BalancesReady = false
	}
	s.snap.NextCSTBidReward = cur.NextCSTBidReward
	s.snap.BidPricesReady = bidPricesReady(s.snap)
	s.snap.ConfigurationReady = ready && configurationReady(s.snap)
	s.mu.Unlock()
}

// refreshVariables reloads the per-round live contract state. As with the
// constants, every field keeps its legacy failure sentinel; the last-bidder
// address keeps its previous value when the read fails.
func (s *State) refreshVariables(ctx context.Context) {
	s.contractRefreshMu.Lock()
	defer s.contractRefreshMu.Unlock()
	ctx, cancel := context.WithTimeout(ctx, s.rpcReadTimeout)
	defer cancel()

	copts, blockTimestamp, err := s.latestCallOpts(ctx)
	if err != nil {
		s.logf("Error fetching block for contract variables: %v\n", err)
		s.markVariablesUnavailable()
		return
	}

	v1, v2 := s.bindLiveReaders()
	if v1 == nil {
		s.logf("Can't instantiate CosmicGame contract at %v . Contract variables won't be fetched\n", s.addrs.CosmicGame)
		return
	}

	cur := s.Snapshot()

	cur.RoundStartAuctionLength = s.roundStartCSTAuctionSetting(v1, v2, &copts)
	if cur.RoundStartAuctionLength == -1 {
		s.logf("Error reading CST round-start auction setting (V1 divisor / V2 duration)\n")
	}
	cur.MechanicsVersion = s.mechanicsVersion()

	if val, err := v1.GetNextEthBidPrice(&copts); err != nil {
		s.logf("Error at GetBidPrice() call: %v\n", err)
		cur.BidPrice = "error"
		cur.BlockPinnedBidPrice = "error"
	} else {
		cur.BidPrice = val.String()
		cur.BlockPinnedBidPrice = val.String()
		cur.BidPriceEth = weiToEth(val)
	}
	if val, err := v1.GetNextCstBidPrice(&copts); err != nil {
		s.logf("Error at GetNextCstBidPrice() call: %v\n", err)
		cur.NextCSTBidPrice = "error"
	} else {
		cur.NextCSTBidPrice = val.String()
	}
	if duration, elapsed, err := v1.GetEthDutchAuctionDurations(&copts); err != nil {
		s.logf("Error at GetEthDutchAuctionDurations() call: %v\n", err)
		cur.ETHAuctionDuration = -1
		cur.ETHAuctionElapsed = -1
	} else if durationSeconds, elapsedSeconds, ok := normalizeAuctionProgress(
		duration,
		elapsed,
		blockTimestamp,
		false,
	); !ok {
		s.logf("Error at GetEthDutchAuctionDurations() call: values exceed int64\n")
		cur.ETHAuctionDuration = -1
		cur.ETHAuctionElapsed = -1
	} else {
		cur.ETHAuctionDuration = durationSeconds
		cur.ETHAuctionElapsed = elapsedSeconds
	}
	if duration, elapsed, err := v1.GetCstDutchAuctionDurations(&copts); err != nil {
		s.logf("Error at GetCstDutchAuctionDurations() call: %v\n", err)
		cur.CSTAuctionDuration = -1
		cur.CSTAuctionElapsed = -1
	} else if durationSeconds, elapsedSeconds, ok := normalizeAuctionProgress(
		duration,
		elapsed,
		blockTimestamp,
		true,
	); !ok {
		s.logf("Error at GetCstDutchAuctionDurations() call: values exceed int64\n")
		cur.CSTAuctionDuration = -1
		cur.CSTAuctionElapsed = -1
	} else {
		cur.CSTAuctionDuration = durationSeconds
		cur.CSTAuctionElapsed = elapsedSeconds
	}
	if cur.MechanicsVersion == mechanicsV1 {
		cur.NextCSTBidReward = cur.FixedCSTBidReward
	} else if reward, err := s.tokenReward(v1, v2, &copts); err != nil {
		s.logf("Error at TokenReward() call: %v\n", err)
		cur.TokenReward = "error"
		cur.NextCSTBidReward = "error"
	} else {
		cur.TokenReward = reward
		cur.NextCSTBidReward = reward
	}
	if val, err := v1.GetDurationUntilMainPrize(&copts); err != nil {
		s.logf("Error at PrizeTime() call: %v\n", err)
		cur.PrizeClaimTimestamp = -1
	} else if parsed, ok := nonNegativeInt64(val); !ok {
		s.logf("Error at PrizeTime() call: value exceeds int64\n")
		cur.PrizeClaimTimestamp = -1
	} else {
		cur.PrizeClaimTimestamp = parsed
	}
	if val, err := v1.GetMainEthPrizeAmount(&copts); err != nil {
		s.logf("Error at PrizeAmount() call: %v\n", err)
		cur.PrizeAmount = "error"
	} else {
		cur.PrizeAmount = val.String()
		cur.PrizeAmountEth = weiToEth(val)
	}
	if val, err := v1.GetCosmicSignatureNftStakingTotalEthRewardAmount(&copts); err != nil {
		s.logf("Error at GetCosmicSignatureNftStakingTotalEthRewardAmount() call: %v\n", err)
		cur.StakingAmount = "error"
	} else {
		cur.StakingAmount = val.String()
		cur.StakingAmountEth = weiToEth(val)
	}
	if val, err := v1.GetRaffleTotalEthPrizeAmountForBidders(&copts); err != nil {
		s.logf("Error at RaffleAmount() call: %v\n", err)
		cur.RaffleAmount = "error"
	} else {
		cur.RaffleAmount = val.String()
		cur.RaffleAmountEth = weiToEth(val)
	}
	if val, err := v1.RoundNum(&copts); err != nil {
		s.logf("Error at RoundNum() call: %v\n", err)
		cur.RoundNum = -1
	} else if parsed, ok := nonNegativeInt64(val); !ok {
		s.logf("Error at RoundNum() call: value exceeds int64\n")
		cur.RoundNum = -1
	} else {
		cur.RoundNum = parsed
	}
	if val, err := v1.MainPrizeTimeIncrementInMicroSeconds(&copts); err != nil {
		s.logf("Error at MainPrizeTimeIncrementInMicroseconds() call: %v\n", err)
		cur.MainPrizeTimeIncrement = "error"
	} else {
		cur.MainPrizeTimeIncrement = val.String()
	}
	if addr, err := v1.LastBidderAddress(&copts); err != nil {
		// Keeps the previous address on failure (legacy behavior).
		s.logf("Error at LastBidder() call: %v\n", err)
	} else {
		cur.LastBidder = addr
	}
	if val, err := v1.InitialDurationUntilMainPrizeDivisor(&copts); err != nil {
		s.logf("Error at InitialDurationUntilMainPrizeDivisor() call: %v\n", err)
		cur.InitialSecondsUntilPrize = -1
	} else if parsed, ok := nonNegativeInt64(val); !ok {
		s.logf("Error at InitialDurationUntilMainPrizeDivisor() call: value exceeds int64\n")
		cur.InitialSecondsUntilPrize = -1
	} else {
		cur.InitialSecondsUntilPrize = parsed
	}
	if val, err := v1.TimeoutDurationToClaimMainPrize(&copts); err != nil {
		s.logf("Error at TimeoutClaimPrize() call: %v\n", err)
		cur.TimeoutClaimPrize = -1
	} else if parsed, ok := nonNegativeInt64(val); !ok {
		s.logf("Error at TimeoutClaimPrize() call: value exceeds int64\n")
		cur.TimeoutClaimPrize = -1
	} else {
		cur.TimeoutClaimPrize = parsed
	}
	s.mu.Lock()
	s.snap.BidPrice = cur.BidPrice
	s.snap.BidPriceEth = cur.BidPriceEth
	s.snap.BlockPinnedBidPrice = cur.BlockPinnedBidPrice
	s.snap.PrizeClaimTimestamp = cur.PrizeClaimTimestamp
	s.snap.PrizeAmount = cur.PrizeAmount
	s.snap.PrizeAmountEth = cur.PrizeAmountEth
	s.snap.RaffleAmount = cur.RaffleAmount
	s.snap.RaffleAmountEth = cur.RaffleAmountEth
	s.snap.StakingAmount = cur.StakingAmount
	s.snap.StakingAmountEth = cur.StakingAmountEth
	s.snap.RoundNum = cur.RoundNum
	s.snap.MainPrizeTimeIncrement = cur.MainPrizeTimeIncrement
	s.snap.LastBidder = cur.LastBidder
	s.snap.InitialSecondsUntilPrize = cur.InitialSecondsUntilPrize
	s.snap.TimeoutClaimPrize = cur.TimeoutClaimPrize
	s.snap.RoundStartAuctionLength = cur.RoundStartAuctionLength
	s.snap.TokenReward = cur.TokenReward
	s.snap.NextCSTBidPrice = cur.NextCSTBidPrice
	s.snap.NextCSTBidReward = cur.NextCSTBidReward
	s.snap.ETHAuctionDuration = cur.ETHAuctionDuration
	s.snap.ETHAuctionElapsed = cur.ETHAuctionElapsed
	s.snap.CSTAuctionDuration = cur.CSTAuctionDuration
	s.snap.CSTAuctionElapsed = cur.CSTAuctionElapsed
	s.snap.MechanicsVersion = cur.MechanicsVersion
	s.snap.VariablesMechanicsVersion = cur.MechanicsVersion
	s.snap.BidPricesReady = bidPricesReady(s.snap)
	s.snap.ConfigurationReady = s.snap.ConstantsReady && configurationReady(s.snap)
	s.mu.Unlock()
}

// refreshDBStats reloads the database aggregates. A failed statistics read
// keeps both previous values; a failed round-start read keeps the previous
// timestamp while the fresh statistics stand (legacy behavior).
func (s *State) refreshDBStats(ctx context.Context) {
	s.dbStatsRefreshMu.Lock()
	defer s.dbStatsRefreshMu.Unlock()
	ctx, cancel := context.WithTimeout(ctx, s.dbReadTimeout)
	defer cancel()

	stats, err := s.db.CosmicGameStatistics(ctx)
	if err != nil {
		s.logger.Error(fmt.Sprintf("state refresh: cosmic game statistics: %v", err))
		return
	}
	roundStartTimestamp := s.Snapshot().RoundStartTimestamp
	ts, err := s.db.RoundStartTimestamp(ctx, stats.TotalPrizes)
	if err != nil {
		s.logger.Error(fmt.Sprintf("state refresh: round start timestamp: %v", err))
	} else {
		roundStartTimestamp = ts
	}
	s.mu.Lock()
	s.snap.Stats = stats
	s.snap.RoundStartTimestamp = roundStartTimestamp
	s.mu.Unlock()
}

func (s *State) refreshBalances(ctx context.Context) {
	s.contractRefreshMu.Lock()
	defer s.contractRefreshMu.Unlock()
	ctx, cancel := context.WithTimeout(ctx, s.rpcReadTimeout)
	defer cancel()

	cur := s.Snapshot()
	copts, _, headerErr := s.latestCallOpts(ctx)
	if headerErr != nil {
		s.logf("Error fetching block for contract balances: %v\n", headerErr)
		cur.CosmicGameBalance = "error"
		cur.CharityBalance = "error"
		cur.CharityBalanceEth = 0
	} else if s.addrs.CosmicGame == (ethcommon.Address{}) {
		cur.CosmicGameBalance = "error"
	} else if val, err := s.client.BalanceAt(ctx, s.addrs.CosmicGame, copts.BlockNumber); err != nil {
		s.logf("Error at BalanceAt() call for cosmic game: %v\n", err)
		cur.CosmicGameBalance = "error"
	} else {
		cur.CosmicGameBalance = val.String()
	}
	if headerErr == nil {
		val, err := s.client.BalanceAt(ctx, cur.CharityAddr, copts.BlockNumber)
		if err != nil {
			s.logf("Error at BalanceAt() call for charity addr: %v\n", err)
			cur.CharityBalance = "error"
			cur.CharityBalanceEth = 0
		} else {
			cur.CharityBalance = val.String()
			cur.CharityBalanceEth = weiToEth(val)
		}
	}
	cur.BalanceCharityAddr = cur.CharityAddr
	cur.BalancesReady = cachedDecimalReady(cur.CosmicGameBalance) &&
		cachedDecimalReady(cur.CharityBalance) &&
		cur.CharityAddr != (ethcommon.Address{}) &&
		cur.BalanceCharityAddr == cur.CharityAddr

	s.mu.Lock()
	s.snap.CosmicGameBalance = cur.CosmicGameBalance
	s.snap.CharityBalance = cur.CharityBalance
	s.snap.CharityBalanceEth = cur.CharityBalanceEth
	s.snap.BalanceCharityAddr = cur.BalanceCharityAddr
	s.snap.BalancesReady = cur.BalancesReady
	s.mu.Unlock()
}

func cachedDecimalReady(value string) bool {
	if value == "" || value == "error" {
		return false
	}
	parsed, ok := new(big.Int).SetString(value, 10)
	return ok && parsed.Sign() >= 0
}

func bidPricesReady(snapshot Snapshot) bool {
	return cachedDecimalReady(snapshot.BlockPinnedBidPrice) &&
		cachedDecimalReady(snapshot.NextCSTBidPrice) &&
		cachedDecimalReady(snapshot.NextCSTBidReward) &&
		snapshot.ETHAuctionDuration >= 0 &&
		snapshot.ETHAuctionElapsed >= 0 &&
		snapshot.ETHAuctionElapsed <= snapshot.ETHAuctionDuration &&
		snapshot.CSTAuctionDuration >= 0 &&
		snapshot.CSTAuctionElapsed >= 0 &&
		snapshot.CSTAuctionElapsed <= snapshot.CSTAuctionDuration
}

func configurationReady(snapshot Snapshot) bool {
	percentageReady := func(value int64) bool { return value >= 0 && value <= 100 }
	if !cachedDecimalReady(snapshot.PriceIncrease) ||
		!cachedDecimalReady(snapshot.TimeIncrease) ||
		snapshot.CharityAddr == (ethcommon.Address{}) ||
		snapshot.TreasurerAddr == (ethcommon.Address{}) ||
		!percentageReady(snapshot.CharityPercentage) ||
		!percentageReady(snapshot.PrizePercentage) ||
		!percentageReady(snapshot.RafflePercentage) ||
		!percentageReady(snapshot.ChronoPercentage) ||
		!percentageReady(snapshot.StakingPercentage) ||
		snapshot.RaffleEthWinnersBidding < 0 ||
		snapshot.RaffleNFTWinnersBidding < 0 ||
		snapshot.RaffleNFTWinnersStakingRWalk < 0 ||
		snapshot.InitialSecondsUntilPrize <= 0 ||
		snapshot.TimeoutClaimPrize < 0 ||
		snapshot.RoundStartAuctionLength <= 0 ||
		snapshot.ConstantsMechanicsVersion != snapshot.MechanicsVersion ||
		snapshot.VariablesMechanicsVersion != snapshot.MechanicsVersion {
		return false
	}
	switch snapshot.MechanicsVersion {
	case mechanicsV1:
		return snapshot.CSTAuctionDurationChangeDivisor == -1 &&
			cachedDecimalReady(snapshot.FixedCSTBidReward) &&
			snapshot.BidCSTRewardMultiplier == ""
	case mechanicsV2:
		return snapshot.CSTAuctionDurationChangeDivisor > 0 &&
			cachedDecimalReady(snapshot.BidCSTRewardMultiplier)
	default:
		return false
	}
}

// CosmicGameBalanceEth reads the game contract's live ETH balance. It
// returns NaN on failure; the dashboard maps that to 0 (legacy behavior).
func (s *State) CosmicGameBalanceEth(ctx context.Context) float64 {
	bal, err := s.client.BalanceAt(ctx, s.addrs.CosmicGame, nil)
	if err != nil {
		s.logf("Error at BalanceAt() call for cosmic game: %v\n", err)
		return math.NaN()
	}
	return weiToEth(bal)
}
