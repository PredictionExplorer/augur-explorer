package contractstate

import (
	"context"
	"math"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
)

var weiPerEth = new(big.Float).SetInt(big.NewInt(1e18))

// weiToEth converts a wei amount to a float64 ETH value the way the legacy
// code did (big.Float quotient, rounding ignored).
func weiToEth(wei *big.Int) float64 {
	quo := new(big.Float).Quo(new(big.Float).SetInt(wei), weiPerEth)
	out, _ := quo.Float64()
	return out
}

// refreshConstants reloads the owner-tunable contract parameters. Each field
// keeps its legacy failure sentinel ("error" strings, -1 counters, 0 charity
// percentage, previous charity address), so a dead RPC node degrades the
// dashboard exactly as before.
func (s *State) refreshConstants(ctx context.Context) {
	copts := bind.CallOpts{Context: ctx}

	code, err := s.client.CodeAt(ctx, s.addrs.CosmicGame, nil)
	if err != nil {
		s.logf("Can't instantiate CosmicGame contract: %v\n", err)
	} else if len(code) == 0 {
		s.logf("Can't instantiate CosmicGame contract: no code at given address\n")
	}

	v1, v2 := s.bindLiveReaders()
	if v1 == nil {
		s.logf("Can't instantiate CosmicGame contract at %v . Contract constants won't be fetched\n", s.addrs.CosmicGame)
		return
	}

	cur := s.Snapshot()

	if val, err := v1.EthBidPriceIncreaseDivisor(&copts); err != nil {
		s.logf("Error at PriceIncrease() call: %v\n", err)
		cur.PriceIncrease = "error"
	} else {
		cur.PriceIncrease = val.String()
	}
	if addr, err := v1.CharityAddress(&copts); err != nil {
		// Keeps the previous address on failure (legacy behavior).
		s.logf("Error at Charity() call: %v\n", err)
	} else {
		cur.CharityAddr = addr
	}
	if val, err := v1.CharityEthDonationAmountPercentage(&copts); err != nil {
		s.logf("Error at Charity() call: %v\n", err)
		cur.CharityPercentage = 0
	} else {
		cur.CharityPercentage = val.Int64()
	}
	if reward, err := s.tokenReward(v1, v2, &copts); err != nil {
		s.logf("Error at TokenReward() call: %v\n", err)
		cur.TokenReward = "error"
	} else {
		cur.TokenReward = reward
	}
	if val, err := v1.MainEthPrizeAmountPercentage(&copts); err != nil {
		s.logf("Error at PrizePercentage() call: %v\n", err)
		cur.PrizePercentage = -1
	} else {
		cur.PrizePercentage = val.Int64()
	}
	if val, err := v1.RaffleTotalEthPrizeAmountForBiddersPercentage(&copts); err != nil {
		s.logf("Error at RafflePercentage() call: %v\n", err)
		cur.RafflePercentage = -1
	} else {
		cur.RafflePercentage = val.Int64()
	}
	if val, err := v1.ChronoWarriorEthPrizeAmountPercentage(&copts); err != nil {
		s.logf("Error at ChronoWarriorEthPrizeAmountPercentage() call: %v\n", err)
		cur.ChronoPercentage = -1
	} else {
		cur.ChronoPercentage = val.Int64()
	}
	if val, err := v1.CosmicSignatureNftStakingTotalEthRewardAmountPercentage(&copts); err != nil {
		s.logf("Error at StakingPercentage() call: %v\n", err)
		cur.StakingPercentage = -1
	} else {
		cur.StakingPercentage = val.Int64()
	}
	if val, err := v1.MainPrizeTimeIncrementIncreaseDivisor(&copts); err != nil {
		s.logf("Error at TimeIncrease() call: %v\n", err)
		cur.TimeIncrease = "error"
	} else {
		cur.TimeIncrease = val.String()
	}
	if val, err := v1.NumRaffleEthPrizesForBidders(&copts); err != nil {
		s.logf("Error at NumRaffleETHWinnersBidding() call: %v\n", err)
		cur.RaffleEthWinnersBidding = -1
	} else {
		cur.RaffleEthWinnersBidding = val.Int64()
	}
	if val, err := v1.NumRaffleCosmicSignatureNftsForBidders(&copts); err != nil {
		s.logf("Error at NumRaffleNFTWinnersBidding() call: %v\n", err)
		cur.RaffleNFTWinnersBidding = -1
	} else {
		cur.RaffleNFTWinnersBidding = val.Int64()
	}
	if val, err := v1.NumRaffleCosmicSignatureNftsForRandomWalkNftStakers(&copts); err != nil {
		s.logf("Error at NumRaffleNFTWinnersStakingRWalk() call: %v\n", err)
		cur.RaffleNFTWinnersStakingRWalk = -1
	} else {
		cur.RaffleNFTWinnersStakingRWalk = val.Int64()
	}
	cur.CSTAuctionDurationChangeDivisor = s.cstAuctionDurationChangeDivisor(v1, v2, &copts)

	s.mu.Lock()
	s.snap.PriceIncrease = cur.PriceIncrease
	s.snap.CharityAddr = cur.CharityAddr
	s.snap.CharityPercentage = cur.CharityPercentage
	s.snap.TokenReward = cur.TokenReward
	s.snap.PrizePercentage = cur.PrizePercentage
	s.snap.RafflePercentage = cur.RafflePercentage
	s.snap.ChronoPercentage = cur.ChronoPercentage
	s.snap.StakingPercentage = cur.StakingPercentage
	s.snap.TimeIncrease = cur.TimeIncrease
	s.snap.RaffleEthWinnersBidding = cur.RaffleEthWinnersBidding
	s.snap.RaffleNFTWinnersBidding = cur.RaffleNFTWinnersBidding
	s.snap.RaffleNFTWinnersStakingRWalk = cur.RaffleNFTWinnersStakingRWalk
	s.snap.CSTAuctionDurationChangeDivisor = cur.CSTAuctionDurationChangeDivisor
	s.mu.Unlock()
}

// refreshVariables reloads the per-round live contract state. As with the
// constants, every field keeps its legacy failure sentinel; the last-bidder
// address keeps its previous value when the read fails.
func (s *State) refreshVariables(ctx context.Context) {
	copts := bind.CallOpts{Context: ctx}

	v1, v2 := s.bindLiveReaders()
	if v1 == nil {
		s.logf("Can't instantiate CosmicGame contract at %v . Contract variables won't be fetched\n", s.addrs.CosmicGame)
		return
	}

	cur := s.Snapshot()

	if val, err := v1.GetNextEthBidPrice(&copts); err != nil {
		s.logf("Error at GetBidPrice() call: %v\n", err)
		cur.BidPrice = "error"
	} else {
		cur.BidPrice = val.String()
		cur.BidPriceEth = weiToEth(val)
	}
	if val, err := v1.GetDurationUntilMainPrize(&copts); err != nil {
		s.logf("Error at PrizeTime() call: %v\n", err)
		cur.PrizeClaimTimestamp = -1
	} else {
		cur.PrizeClaimTimestamp = val.Int64()
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
	} else {
		cur.RoundNum = val.Int64()
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
	} else {
		cur.InitialSecondsUntilPrize = val.Int64()
	}
	if val, err := v1.TimeoutDurationToClaimMainPrize(&copts); err != nil {
		s.logf("Error at TimeoutClaimPrize() call: %v\n", err)
		cur.TimeoutClaimPrize = -1
	} else {
		cur.TimeoutClaimPrize = val.Int64()
	}
	cur.RoundStartAuctionLength = s.roundStartCSTAuctionSetting(v1, v2, &copts)
	if cur.RoundStartAuctionLength == -1 {
		s.logf("Error reading CST round-start auction setting (V1 divisor / V2 duration)\n")
	}
	if val, err := s.client.BalanceAt(ctx, cur.CharityAddr, nil); err != nil {
		s.logf("Error at BalanceAt() call for charity addr: %v\n", err)
		cur.CharityBalance = "error"
	} else {
		cur.CharityBalance = val.String()
		cur.CharityBalanceEth = weiToEth(val)
	}

	s.mu.Lock()
	s.snap.BidPrice = cur.BidPrice
	s.snap.BidPriceEth = cur.BidPriceEth
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
	s.snap.CharityBalance = cur.CharityBalance
	s.snap.CharityBalanceEth = cur.CharityBalanceEth
	s.mu.Unlock()
}

// refreshDBStats reloads the database aggregates. A failed statistics read
// keeps both previous values; a failed round-start read keeps the previous
// timestamp while the fresh statistics stand (legacy behavior).
func (s *State) refreshDBStats(ctx context.Context) {
	stats, err := s.db.CosmicGameStatistics(ctx)
	if err != nil {
		s.errlog.Printf("state refresh: cosmic game statistics: %v", err)
		return
	}
	s.mu.Lock()
	s.snap.Stats = stats
	s.mu.Unlock()

	ts, err := s.db.RoundStartTimestamp(ctx, stats.TotalPrizes)
	if err != nil {
		s.errlog.Printf("state refresh: round start timestamp: %v", err)
		return
	}
	s.mu.Lock()
	s.snap.RoundStartTimestamp = ts
	s.mu.Unlock()
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
