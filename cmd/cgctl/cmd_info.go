package main

import (
	"context"
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/spf13/cobra"

	"github.com/PredictionExplorer/augur-explorer/cmd/cgctl/internal/ethtx"
	cgcontracts "github.com/PredictionExplorer/augur-explorer/contracts/cosmicgame"
)

func init() {
	c := &cobra.Command{
		Use:   "info [cosmicgame-addr]",
		Short: "Dump comprehensive CosmicGame contract state",
		Long: `Dump comprehensive CosmicGame contract state: round status, timing,
bid prices, champions, prize distribution, raffle config, wallet and contract
addresses.

If no address is given, the default local Hardhat deployment address
` + defaultLocalGameAddr + ` is used.

Environment:
  RPC_URL  Ethereum RPC endpoint (required)`,
		Args: cobra.MaximumNArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			addr := defaultLocalGameAddr
			if len(args) == 1 {
				addr = args[0]
			} else {
				fmt.Printf("Setting default cosmic game contract address to %s\n", defaultLocalGameAddr)
			}
			return runInfo(cmd.Context(), ethtx.NewPrinter(true), addr)
		},
	}
	register(c)
}

func runInfo(ctx context.Context, out *ethtx.Printer, addrArg string) error {
	gameAddr, err := parseAddress("cosmicgame-addr", addrArg)
	if err != nil {
		return err
	}

	net, err := ethtx.Connect(ctx)
	if err != nil {
		return fmt.Errorf("network connection failed: %w", err)
	}
	out.NetworkInfo(net)

	copts := ethtx.CallOpts()
	out.ContractInfo("CosmicSignatureGame", gameAddr)

	game, err := cgcontracts.NewCosmicSignatureGame(gameAddr, net.Client)
	if err != nil {
		return fmt.Errorf("failed to instantiate CosmicGame contract: %w", err)
	}
	gameV2, _ := cgcontracts.NewCosmicSignatureGameV2(gameAddr, net.Client)

	blockTime := int64(net.BlockTime)

	// Round status.
	roundNum, err := game.RoundNum(copts)
	if err != nil {
		return fmt.Errorf("RoundNum(): %w", err)
	}
	activationTime, err := game.RoundActivationTime(copts)
	if err != nil {
		return fmt.Errorf("RoundActivationTime(): %w", err)
	}
	secsToStart := activationTime.Int64() - blockTime
	delayUntilActivation, err := game.DelayDurationBeforeRoundActivation(copts)
	if err != nil {
		return fmt.Errorf("DelayDurationBeforeRoundActivation(): %w", err)
	}
	totalBids, err := game.GetTotalNumBids(copts, roundNum)
	if err != nil {
		return fmt.Errorf("GetTotalNumBids(): %w", err)
	}
	numRaffleParticipants, err := game.BidderAddresses(copts, big.NewInt(roundNum.Int64()))
	if err != nil {
		return fmt.Errorf("BidderAddresses(): %w", err)
	}

	out.Section("ROUND STATUS")
	out.KeyValue("RoundNum", roundNum.String())
	out.KeyValue("Round activation time", activationTime.String())
	out.KeyValue("Block time (current)", blockTime)
	out.KeyValue("Delay before activation", delayUntilActivation)
	if secsToStart > 0 {
		fmt.Printf("%-28s= INACTIVE - activates in %s\n", "Round status", ethtx.FmtDuration(secsToStart))
	} else {
		fmt.Printf("%-28s= ACTIVE - started %s ago\n", "Round status", ethtx.FmtDuration(-secsToStart))
	}
	out.KeyValue("Total bids this round", totalBids.String())
	out.KeyValue("Unique bidders this round", numRaffleParticipants.Int64())

	// Timing / countdown.
	timeUntilPrize, err := game.GetDurationUntilMainPrize(copts)
	if err != nil {
		return fmt.Errorf("GetDurationUntilMainPrize(): %w", err)
	}
	durationUntilPrizeRaw, err := game.GetDurationUntilMainPrizeRaw(copts)
	if err != nil {
		return fmt.Errorf("GetDurationUntilMainPrizeRaw(): %w", err)
	}
	prizeTime, err := game.MainPrizeTime(copts)
	if err != nil {
		return fmt.Errorf("MainPrizeTime(): %w", err)
	}
	timeoutMainPrize, err := game.TimeoutDurationToClaimMainPrize(copts)
	if err != nil {
		return fmt.Errorf("TimeoutDurationToClaimMainPrize(): %w", err)
	}
	durationUntilAnyoneCanClaim := new(big.Int).Add(durationUntilPrizeRaw, timeoutMainPrize)
	anyoneCanClaim := durationUntilAnyoneCanClaim.Sign() <= 0

	timeIncMicrosec, err := game.MainPrizeTimeIncrementInMicroSeconds(copts)
	if err != nil {
		return fmt.Errorf("MainPrizeTimeIncrementInMicroSeconds(): %w", err)
	}
	timeIncrement := float64(timeIncMicrosec.Int64()) / float64(1000000)
	timeIncOnBid, err := game.GetMainPrizeTimeIncrement(copts)
	if err != nil {
		return fmt.Errorf("GetMainPrizeTimeIncrement(): %w", err)
	}
	timeIncIncreaseDivisor, err := game.MainPrizeTimeIncrementIncreaseDivisor(copts)
	if err != nil {
		return fmt.Errorf("MainPrizeTimeIncrementIncreaseDivisor(): %w", err)
	}
	initialDurationDivisor, err := game.InitialDurationUntilMainPrizeDivisor(copts)
	if err != nil {
		return fmt.Errorf("InitialDurationUntilMainPrizeDivisor(): %w", err)
	}
	initialDurationInc := ethtx.ConvertToPercentage(initialDurationDivisor)
	initialDurationSeconds, err := game.GetInitialDurationUntilMainPrize(copts)
	if err != nil {
		return fmt.Errorf("GetInitialDurationUntilMainPrize(): %w", err)
	}

	out.Section("TIMING / COUNTDOWN")
	out.KeyValue("MainPrizeTime (timestamp)", prizeTime.String())
	out.KeyValueDuration("Duration until prize (clamped)", timeUntilPrize.Int64())
	out.KeyValueDuration("Duration until prize (raw)", durationUntilPrizeRaw.Int64())
	out.KeyValueDuration("Timeout for last bidder claim", timeoutMainPrize.Int64())
	out.KeyValueDuration("Duration until anyone can claim", durationUntilAnyoneCanClaim.Int64())
	if anyoneCanClaim {
		out.KeyValue("Can anyone claim now?", "YES")
	} else {
		out.KeyValue("Can anyone claim now?", "NO")
	}
	out.KeyValue("Time increment (microsec)", timeIncMicrosec)
	out.KeyValue("Time increment (seconds)", timeIncrement)
	out.KeyValue("Time increment on bid (current)", fmt.Sprintf("%v sec", timeIncOnBid.Int64()))
	out.KeyValue("Time increment increase divisor", timeIncIncreaseDivisor)
	out.KeyValue("First bid time bump", fmt.Sprintf("%.2f%% (divisor=%v)", initialDurationInc, initialDurationDivisor))
	out.KeyValue("First bid time bump (seconds)", initialDurationSeconds)

	// Bidding / prices.
	nextBidPrice, err := game.NextEthBidPrice(copts)
	if err != nil {
		return fmt.Errorf("NextEthBidPrice(): %w", err)
	}
	bidPriceAuction, err := game.GetNextEthBidPrice(copts)
	if err != nil {
		return fmt.Errorf("GetNextEthBidPrice(): %w", err)
	}
	cstPrice, err := game.GetNextCstBidPrice(copts)
	if err != nil {
		return fmt.Errorf("GetNextCstBidPrice(): %w", err)
	}
	ethBidPriceIncreaseDivisor, err := game.EthBidPriceIncreaseDivisor(copts)
	if err != nil {
		return fmt.Errorf("EthBidPriceIncreaseDivisor(): %w", err)
	}
	priceIncrease := ethtx.ConvertToPercentage(ethBidPriceIncreaseDivisor)

	var cstRewardFixed *big.Int
	var bidCstRewardAmount *big.Int
	var bidCstRewardMultiplier *big.Int
	isV2Reward := false
	if gameV2 != nil {
		if v, errV2 := gameV2.GetBidCstRewardAmount(copts); errV2 == nil {
			bidCstRewardAmount = v
			isV2Reward = true
			bidCstRewardMultiplier, _ = gameV2.BidCstRewardAmountMultiplier(copts)
		}
	}
	if !isV2Reward {
		cstRewardFixed, err = game.CstRewardAmountForBidding(copts)
		if err != nil {
			return fmt.Errorf("CstRewardAmountForBidding(): %w", err)
		}
	}

	cstAuctionDuration, cstAuctionElapsed, err := game.GetCstDutchAuctionDurations(copts)
	if err != nil {
		return fmt.Errorf("GetCstDutchAuctionDurations(): %w", err)
	}
	ethAuctionDuration, ethAuctionElapsed, err := game.GetEthDutchAuctionDurations(copts)
	if err != nil {
		return fmt.Errorf("GetEthDutchAuctionDurations(): %w", err)
	}
	cstDutchBeginPrice, err := game.CstDutchAuctionBeginningBidPrice(copts)
	if err != nil {
		return fmt.Errorf("CstDutchAuctionBeginningBidPrice(): %w", err)
	}
	ethDutchBeginPrice, err := game.EthDutchAuctionBeginningBidPrice(copts)
	if err != nil {
		return fmt.Errorf("EthDutchAuctionBeginningBidPrice(): %w", err)
	}
	ethDutchEndingDivisor, err := game.EthDutchAuctionEndingBidPriceDivisor(copts)
	if err != nil {
		return fmt.Errorf("EthDutchAuctionEndingBidPriceDivisor(): %w", err)
	}

	out.Section("BIDDING / PRICES")
	out.KeyValueEth("storage::nextEthBidPrice", nextBidPrice)
	out.KeyValueEth("NextEthBidPrice (auction)", bidPriceAuction)
	fmt.Printf("%-28s  ^ effective price to pay (storage=0 until first bid)\n", "")
	fmt.Printf("%-28s= %s CST\n", "NextCstBidPrice", ethtx.WeiToEth(cstPrice))
	out.KeyValue("ETH bid price increase div", fmt.Sprintf("%v (%.2f%%)", ethBidPriceIncreaseDivisor, priceIncrease))
	if isV2Reward {
		fmt.Printf("%-28s= %s CST\n", "getBidCstRewardAmount (next bid)", ethtx.WeiToEth(bidCstRewardAmount))
		fmt.Printf("%-28s  ^ V2 dynamic CST reward (dashboard TokenReward); 0 right after a bid\n", "")
		if bidCstRewardMultiplier != nil {
			out.KeyValue("bidCstRewardAmountMultiplier", bidCstRewardMultiplier)
		}
	} else {
		fmt.Printf("%-28s= %s CST\n", "CST reward per bid (fixed)", ethtx.WeiToEth(cstRewardFixed))
	}
	out.KeyValue("ETH Dutch auction elapsed/total", fmt.Sprintf("%v / %v", ethAuctionElapsed.String(), ethAuctionDuration.String()))
	out.KeyValueEth("ETH Dutch auction begin price", ethDutchBeginPrice)
	if ethDutchBeginPrice.Sign() == 0 {
		fmt.Printf("%-28s  ^ when 0, contract uses a minimum floor (e.g. 0.0001 ETH); effective price = NextEthBidPrice (auction) above\n", "")
		fmt.Printf("%-28s    (stored begin price is only set after first bid; getter returns min until then)\n", "")
	}
	out.KeyValue("ETH Dutch auction end divisor", ethDutchEndingDivisor)
	// CST auction: contract may return (duration, startTimestamp) not
	// (duration, secondsElapsed); startTimestamp is Unix seconds.
	cstElapsedSec := cstAuctionElapsed
	if cstAuctionElapsed.Cmp(cstAuctionDuration) > 0 && cstAuctionElapsed.Cmp(big.NewInt(1e9)) > 0 {
		elapsed := blockTime - cstAuctionElapsed.Int64()
		if elapsed < 0 {
			elapsed = 0
		}
		cstElapsedSec = big.NewInt(elapsed)
		fmt.Printf("%-28s= %v / %v (elapsed from start_ts; raw 2nd value was start timestamp)\n", "CST Dutch auction elapsed/total", cstElapsedSec.String(), cstAuctionDuration.String())
	} else {
		out.KeyValue("CST Dutch auction elapsed/total", fmt.Sprintf("%v / %v", cstAuctionElapsed.String(), cstAuctionDuration.String()))
	}
	fmt.Printf("%-28s= %s CST\n", "storage::cstDutchAuctionBeginningBidPrice", ethtx.WeiToEth(cstDutchBeginPrice))
	if cstDutchBeginPrice.Sign() == 0 {
		fmt.Printf("%-28s  ^ when 0, effective CST price = getNextCstBidPrice() (min limit / auction)\n", "")
	}

	// V2 initialized parameters (initializeV2). Only printed when the proxy is
	// on V2 (the V2-only getters succeed); on a V1 contract they revert.
	if gameV2 != nil {
		if cstAucChgDiv, errV2 := gameV2.CstDutchAuctionDurationChangeDivisor(copts); errV2 == nil {
			cstAucDurationV2, _ := gameV2.CstDutchAuctionDuration(copts)
			bidCstRewardMultiplierV2, _ := gameV2.BidCstRewardAmountMultiplier(copts)

			out.Section("V2 INITIALIZED PARAMETERS (initializeV2)")
			out.KeyValueDuration("cstDutchAuctionDuration", cstAucDurationV2.Int64())
			out.KeyValue("cstDutchAuctionDurationChangeDivisor", cstAucChgDiv)
			out.KeyValue("bidCstRewardAmountMultiplier", bidCstRewardMultiplierV2)
			out.KeyValueDuration("timeoutDurationToClaimMainPrize", timeoutMainPrize.Int64())
		}
	}

	// Current bidders / champions.
	lastBidder, err := game.LastBidderAddress(copts)
	if err != nil {
		return fmt.Errorf("LastBidderAddress(): %w", err)
	}
	lastCstBidder, err := game.LastCstBidderAddress(copts)
	if err != nil {
		return fmt.Errorf("LastCstBidderAddress(): %w", err)
	}
	currentChampions, err := game.TryGetCurrentChampions(copts)
	if err != nil {
		return fmt.Errorf("TryGetCurrentChampions(): %w", err)
	}
	enduranceStartTs, err := game.EnduranceChampionStartTimeStamp(copts)
	if err != nil {
		return fmt.Errorf("EnduranceChampionStartTimeStamp(): %w", err)
	}
	prevEnduranceDuration, err := game.PrevEnduranceChampionDuration(copts)
	if err != nil {
		return fmt.Errorf("PrevEnduranceChampionDuration(): %w", err)
	}

	out.Section("CURRENT BIDDERS / CHAMPIONS")
	out.KeyValue("Last ETH bidder", lastBidder.String())
	out.KeyValue("Last CST bidder", lastCstBidder.String())
	var zeroAddr common.Address
	lastCstBidderPrizeMinted := "NO"
	if lastCstBidder != zeroAddr {
		lastCstBidderPrizeMinted = "YES"
	}
	out.KeyValue("LastCSTBidder prize (this round)", lastCstBidderPrizeMinted)
	out.KeyValue("Endurance champion", currentChampions.EnduranceChampionAddress.String())
	out.KeyValueDuration("Endurance champion duration", currentChampions.EnduranceChampionDuration.Int64())
	out.KeyValue("Endurance champion start ts", enduranceStartTs.String())
	out.KeyValueDuration("Prev endurance champion dur", prevEnduranceDuration.Int64())
	out.KeyValue("Chrono Warrior", currentChampions.ChronoWarriorAddress.String())
	out.KeyValueDuration("Chrono Warrior duration", currentChampions.ChronoWarriorDuration.Int64())

	// Prize distribution.
	balance, err := net.Client.BalanceAt(ctx, gameAddr, nil)
	if err != nil {
		return fmt.Errorf("BalanceAt(): %w", err)
	}
	prizeAmount, err := game.GetMainEthPrizeAmount(copts)
	if err != nil {
		return fmt.Errorf("GetMainEthPrizeAmount(): %w", err)
	}
	mainPrizePercentage, err := game.MainEthPrizeAmountPercentage(copts)
	if err != nil {
		return fmt.Errorf("MainEthPrizeAmountPercentage(): %w", err)
	}
	charityAmount, err := game.GetCharityEthDonationAmount(copts)
	if err != nil {
		return fmt.Errorf("GetCharityEthDonationAmount(): %w", err)
	}
	charityPercentage, err := game.CharityEthDonationAmountPercentage(copts)
	if err != nil {
		return fmt.Errorf("CharityEthDonationAmountPercentage(): %w", err)
	}
	raffleAmount, err := game.GetRaffleTotalEthPrizeAmountForBidders(copts)
	if err != nil {
		return fmt.Errorf("GetRaffleTotalEthPrizeAmountForBidders(): %w", err)
	}
	rafflePercentage, err := game.RaffleTotalEthPrizeAmountForBiddersPercentage(copts)
	if err != nil {
		return fmt.Errorf("RaffleTotalEthPrizeAmountForBiddersPercentage(): %w", err)
	}
	chronoAmount, err := game.GetChronoWarriorEthPrizeAmount(copts)
	if err != nil {
		return fmt.Errorf("GetChronoWarriorEthPrizeAmount(): %w", err)
	}
	chronoPercentage, err := game.ChronoWarriorEthPrizeAmountPercentage(copts)
	if err != nil {
		return fmt.Errorf("ChronoWarriorEthPrizeAmountPercentage(): %w", err)
	}
	stakingAmount, err := game.GetCosmicSignatureNftStakingTotalEthRewardAmount(copts)
	if err != nil {
		return fmt.Errorf("GetCosmicSignatureNftStakingTotalEthRewardAmount(): %w", err)
	}
	stakingPercentage, err := game.CosmicSignatureNftStakingTotalEthRewardAmountPercentage(copts)
	if err != nil {
		return fmt.Errorf("CosmicSignatureNftStakingTotalEthRewardAmountPercentage(): %w", err)
	}
	cstPrizeAmount, err := game.CstPrizeAmount(copts)
	if err != nil {
		return fmt.Errorf("CstPrizeAmount(): %w", err)
	}

	out.Section("PRIZE DISTRIBUTION")
	out.KeyValueEth("Contract balance", balance)
	fmt.Printf("%-28s= %s ETH (%v%%)\n", "Main prize amount", ethtx.WeiToEth(prizeAmount), mainPrizePercentage)
	fmt.Printf("%-28s= %s ETH (%v%%)\n", "Charity amount", ethtx.WeiToEth(charityAmount), charityPercentage)
	fmt.Printf("%-28s= %s ETH (%v%%)\n", "Raffle amount (bidders)", ethtx.WeiToEth(raffleAmount), rafflePercentage)
	fmt.Printf("%-28s= %s ETH (%v%%)\n", "Chrono Warrior amount", ethtx.WeiToEth(chronoAmount), chronoPercentage)
	fmt.Printf("%-28s= %s ETH (%v%%)\n", "Staking reward amount", ethtx.WeiToEth(stakingAmount), stakingPercentage)
	fmt.Printf("%-28s= %s CST\n", "CST prize amount", ethtx.WeiToEth(cstPrizeAmount))

	// Raffle config.
	numEthWinners, err := game.NumRaffleEthPrizesForBidders(copts)
	if err != nil {
		return fmt.Errorf("NumRaffleEthPrizesForBidders(): %w", err)
	}
	nftBidders, err := game.NumRaffleCosmicSignatureNftsForBidders(copts)
	if err != nil {
		return fmt.Errorf("NumRaffleCosmicSignatureNftsForBidders(): %w", err)
	}
	nftStakers, err := game.NumRaffleCosmicSignatureNftsForRandomWalkNftStakers(copts)
	if err != nil {
		return fmt.Errorf("NumRaffleCosmicSignatureNftsForRandomWalkNftStakers(): %w", err)
	}

	out.Section("RAFFLE CONFIG")
	out.KeyValue("Num ETH raffle winners", numEthWinners)
	out.KeyValue("Num NFT winners (bidders)", nftBidders)
	out.KeyValue("Num NFT winners (RW stakers)", nftStakers)

	// Prizes wallet.
	prizeWalletAddr, err := game.PrizesWallet(copts)
	if err != nil {
		return fmt.Errorf("PrizesWallet(): %w", err)
	}
	prizesWallet, err := cgcontracts.NewPrizesWallet(prizeWalletAddr, net.Client)
	if err != nil {
		return fmt.Errorf("instantiating PrizesWallet: %w", err)
	}
	numDonatedNfts, err := prizesWallet.NextDonatedNftIndex(copts)
	if err != nil {
		return fmt.Errorf("NextDonatedNftIndex(): %w", err)
	}
	timeoutClaim, err := prizesWallet.TimeoutDurationToWithdrawPrizes(copts)
	if err != nil {
		return fmt.Errorf("TimeoutDurationToWithdrawPrizes(): %w", err)
	}

	out.Section("PRIZES WALLET")
	out.KeyValue("PrizesWallet address", prizeWalletAddr.String())
	out.KeyValue("Num donated NFTs", numDonatedNfts.String())
	out.KeyValueDuration("Timeout to withdraw prizes", timeoutClaim.Int64())

	// Charity.
	charityAddr, err := game.CharityAddress(copts)
	if err != nil {
		return fmt.Errorf("CharityAddress(): %w", err)
	}
	charityWallet, err := cgcontracts.NewCharityWallet(charityAddr, net.Client)
	if err != nil {
		return fmt.Errorf("failed to instantiate CharityWallet contract: %w", err)
	}
	charityRecvAddr, err := charityWallet.CharityAddress(copts)
	if err != nil {
		return fmt.Errorf("calling CharityAddress(): %w", err)
	}

	out.Section("CHARITY")
	out.KeyValue("CharityWallet address", charityAddr.String())
	out.KeyValue("Charity donation receiver", charityRecvAddr.String())

	// Contract addresses.
	nftAddr, err := game.Nft(copts)
	if err != nil {
		return fmt.Errorf("Nft(): %w", err)
	}
	tokenAddr, err := game.Token(copts)
	if err != nil {
		return fmt.Errorf("Token(): %w", err)
	}
	randomwalkAddr, err := game.RandomWalkNft(copts)
	if err != nil {
		return fmt.Errorf("RandomWalkNft(): %w", err)
	}
	stakingAddrCst, err := game.StakingWalletCosmicSignatureNft(copts)
	if err != nil {
		return fmt.Errorf("StakingWalletCosmicSignatureNft(): %w", err)
	}
	stakingAddrRwalk, err := game.StakingWalletRandomWalkNft(copts)
	if err != nil {
		return fmt.Errorf("StakingWalletRandomWalkNft(): %w", err)
	}
	marketingAddr, err := game.MarketingWallet(copts)
	if err != nil {
		return fmt.Errorf("MarketingWallet(): %w", err)
	}
	marketingCstAmount, err := game.MarketingWalletCstContributionAmount(copts)
	if err != nil {
		return fmt.Errorf("MarketingWalletCstContributionAmount(): %w", err)
	}
	ownerAddr, err := game.Owner(copts)
	if err != nil {
		return fmt.Errorf("Owner(): %w", err)
	}

	out.Section("CONTRACT ADDRESSES")
	out.KeyValue("CosmicSignatureNft", nftAddr.String())
	out.KeyValue("CosmicSignatureToken", tokenAddr.String())
	out.KeyValue("RandomWalkNft", randomwalkAddr.String())
	out.KeyValue("StakingWallet (CST NFT)", stakingAddrCst.String())
	out.KeyValue("StakingWallet (RandomWalk)", stakingAddrRwalk.String())
	out.KeyValue("MarketingWallet", marketingAddr.String())
	fmt.Printf("%-28s= %s CST\n", "MarketingWallet CST contrib", ethtx.WeiToEth(marketingCstAmount))
	out.KeyValue("Owner", ownerAddr.String())

	// Config parameters.
	bidMsgMaxLen, err := game.BidMessageLengthMaxLimit(copts)
	if err != nil {
		return fmt.Errorf("BidMessageLengthMaxLimit(): %w", err)
	}

	out.Section("CONFIG PARAMETERS")
	out.KeyValue("Bid message max length", bidMsgMaxLen)
	return nil
}
