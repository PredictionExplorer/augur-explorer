// Comprehensive CosmicGame contract state dump
package main

import (
	"context"
	"fmt"
	"math/big"
	"os"

	"github.com/ethereum/go-ethereum/common"

	. "github.com/PredictionExplorer/augur-explorer/rwcg/contracts/cosmicgame"
	cutils "github.com/PredictionExplorer/augur-explorer/rwcg/etl/cosmicgame/scripts/common"
)

func main() {
	// Usage
	var cgAddr string
	if len(os.Args) < 2 {
		cutils.PrintUsage(os.Args[0],
			"[cosmicgame_contract_addr]",
			"Gets comprehensive CosmicGame contract state",
			map[string]string{"RPC_URL": "Ethereum RPC endpoint (required)"},
		)
		fmt.Printf("Setting default cosmic game contract address to 0x5FbDB2315678afecb367f032d93F642f64180aa3\n")
		cgAddr = "0x5FbDB2315678afecb367f032d93F642f64180aa3"
	} else {
		cgAddr = os.Args[1]
	}

	// Connect to network
	net, err := cutils.ConnectToRPC()
	if err != nil {
		cutils.Fatal("Network connection failed: %v", err)
	}
	cutils.PrintNetworkInfo(net)

	copts := cutils.CreateCallOpts()
	cosmicGameAddr := common.HexToAddress(cgAddr)
	cutils.PrintContractInfo("CosmicSignatureGame", cosmicGameAddr)

	cosmicGame, err := NewCosmicSignatureGame(cosmicGameAddr, net.Client)
	if err != nil {
		cutils.Fatal("Failed to instantiate CosmicGame contract: %v", err)
	}

	// Get current block for time calculations
	blockTime := int64(net.BlockTime)

	// ==================== ROUND STATUS ====================
	roundNum, err := cosmicGame.RoundNum(copts)
	if err != nil {
		cutils.Fatal("Error at RoundNum(): %v", err)
	}
	activationTime, err := cosmicGame.RoundActivationTime(copts)
	if err != nil {
		cutils.Fatal("Error at RoundActivationTime(): %v", err)
	}
	secsToStart := activationTime.Int64() - blockTime
	delayUntilActivation, err := cosmicGame.DelayDurationBeforeRoundActivation(copts)
	if err != nil {
		cutils.Fatal("Error at DelayDurationBeforeRoundActivation(): %v", err)
	}
	totalBids, err := cosmicGame.GetTotalNumBids(copts, roundNum)
	if err != nil {
		cutils.Fatal("Error at GetTotalNumBids(): %v", err)
	}
	numRaffleParticipants, err := cosmicGame.BidderAddresses(copts, big.NewInt(roundNum.Int64()))
	if err != nil {
		cutils.Fatal("Error at BidderAddresses(): %v", err)
	}

	cutils.Section("ROUND STATUS")
	cutils.PrintKeyValue("RoundNum", roundNum.String())
	cutils.PrintKeyValue("Round activation time", activationTime.String())
	cutils.PrintKeyValue("Block time (current)", blockTime)
	cutils.PrintKeyValue("Delay before activation", delayUntilActivation)
	if secsToStart > 0 {
		fmt.Printf("%-28s= INACTIVE - activates in %s\n", "Round status", cutils.FmtDuration(secsToStart))
	} else {
		fmt.Printf("%-28s= ACTIVE - started %s ago\n", "Round status", cutils.FmtDuration(-secsToStart))
	}
	cutils.PrintKeyValue("Total bids this round", totalBids.String())
	cutils.PrintKeyValue("Unique bidders this round", numRaffleParticipants.Int64())

	// ==================== TIMING / COUNTDOWN ====================
	timeUntilPrize, err := cosmicGame.GetDurationUntilMainPrize(copts)
	if err != nil {
		cutils.Fatal("Error at GetDurationUntilMainPrize(): %v", err)
	}
	durationUntilPrizeRaw, err := cosmicGame.GetDurationUntilMainPrizeRaw(copts)
	if err != nil {
		cutils.Fatal("Error at GetDurationUntilMainPrizeRaw(): %v", err)
	}
	prizeTime, err := cosmicGame.MainPrizeTime(copts)
	if err != nil {
		cutils.Fatal("Error at MainPrizeTime(): %v", err)
	}
	timeoutMainPrize, err := cosmicGame.TimeoutDurationToClaimMainPrize(copts)
	if err != nil {
		cutils.Fatal("Error at TimeoutDurationToClaimMainPrize(): %v", err)
	}
	durationUntilAnyoneCanClaim := new(big.Int).Add(durationUntilPrizeRaw, timeoutMainPrize)
	anyoneCanClaim := durationUntilAnyoneCanClaim.Cmp(big.NewInt(0)) <= 0

	timeIncMicrosec, err := cosmicGame.MainPrizeTimeIncrementInMicroSeconds(copts)
	if err != nil {
		cutils.Fatal("Error at MainPrizeTimeIncrementInMicroSeconds(): %v", err)
	}
	timeIncrement := float64(timeIncMicrosec.Int64()) / float64(1000000)
	timeIncOnBid, err := cosmicGame.GetMainPrizeTimeIncrement(copts)
	if err != nil {
		cutils.Fatal("Error at GetMainPrizeTimeIncrement(): %v", err)
	}
	timeIncIncreaseDivisor, err := cosmicGame.MainPrizeTimeIncrementIncreaseDivisor(copts)
	if err != nil {
		cutils.Fatal("Error at MainPrizeTimeIncrementIncreaseDivisor(): %v", err)
	}
	initialDurationDivisor, err := cosmicGame.InitialDurationUntilMainPrizeDivisor(copts)
	if err != nil {
		cutils.Fatal("Error at InitialDurationUntilMainPrizeDivisor(): %v", err)
	}
	initialDurationInc := cutils.ConvertToPercentage(initialDurationDivisor)
	initialDurationSeconds, err := cosmicGame.GetInitialDurationUntilMainPrize(copts)
	if err != nil {
		cutils.Fatal("Error at GetInitialDurationUntilMainPrize(): %v", err)
	}

	cutils.Section("TIMING / COUNTDOWN")
	cutils.PrintKeyValue("MainPrizeTime (timestamp)", prizeTime.String())
	cutils.PrintKeyValueDuration("Duration until prize (clamped)", timeUntilPrize.Int64())
	cutils.PrintKeyValueDuration("Duration until prize (raw)", durationUntilPrizeRaw.Int64())
	cutils.PrintKeyValueDuration("Timeout for last bidder claim", timeoutMainPrize.Int64())
	cutils.PrintKeyValueDuration("Duration until anyone can claim", durationUntilAnyoneCanClaim.Int64())
	if anyoneCanClaim {
		cutils.PrintKeyValue("Can anyone claim now?", "YES")
	} else {
		cutils.PrintKeyValue("Can anyone claim now?", "NO")
	}
	cutils.PrintKeyValue("Time increment (microsec)", timeIncMicrosec)
	cutils.PrintKeyValue("Time increment (seconds)", timeIncrement)
	cutils.PrintKeyValue("Time increment on bid (current)", fmt.Sprintf("%v sec", timeIncOnBid.Int64()))
	cutils.PrintKeyValue("Time increment increase divisor", timeIncIncreaseDivisor)
	cutils.PrintKeyValue("First bid time bump", fmt.Sprintf("%.2f%% (divisor=%v)", initialDurationInc, initialDurationDivisor))
	cutils.PrintKeyValue("First bid time bump (seconds)", initialDurationSeconds)

	// ==================== BIDDING / PRICES ====================
	nextBidPrice, err := cosmicGame.NextEthBidPrice(copts)
	if err != nil {
		cutils.Fatal("Error at NextEthBidPrice(): %v", err)
	}
	bidPriceAuction, err := cosmicGame.GetNextEthBidPrice(copts)
	if err != nil {
		cutils.Fatal("Error at GetNextEthBidPrice(): %v", err)
	}
	cstPrice, err := cosmicGame.GetNextCstBidPrice(copts)
	if err != nil {
		cutils.Fatal("Error at GetNextCstBidPrice(): %v", err)
	}
	ethBidPriceIncreaseDivisor, err := cosmicGame.EthBidPriceIncreaseDivisor(copts)
	if err != nil {
		cutils.Fatal("Error at EthBidPriceIncreaseDivisor(): %v", err)
	}
	priceIncrease := cutils.ConvertToPercentage(ethBidPriceIncreaseDivisor)
	cstReward, err := cosmicGame.CstRewardAmountForBidding(copts)
	if err != nil {
		cutils.Fatal("Error at CstRewardAmountForBidding(): %v", err)
	}

	// Dutch auction info
	cstAuctionDuration, cstAuctionElapsed, err := cosmicGame.GetCstDutchAuctionDurations(copts)
	if err != nil {
		cutils.Fatal("Error at GetCstDutchAuctionDurations(): %v", err)
	}
	ethAuctionDuration, ethAuctionElapsed, err := cosmicGame.GetEthDutchAuctionDurations(copts)
	if err != nil {
		cutils.Fatal("Error at GetEthDutchAuctionDurations(): %v", err)
	}
	cstDutchBeginPrice, err := cosmicGame.CstDutchAuctionBeginningBidPrice(copts)
	if err != nil {
		cutils.Fatal("Error at CstDutchAuctionBeginningBidPrice(): %v", err)
	}
	ethDutchBeginPrice, err := cosmicGame.EthDutchAuctionBeginningBidPrice(copts)
	if err != nil {
		cutils.Fatal("Error at EthDutchAuctionBeginningBidPrice(): %v", err)
	}
	ethDutchEndingDivisor, err := cosmicGame.EthDutchAuctionEndingBidPriceDivisor(copts)
	if err != nil {
		cutils.Fatal("Error at EthDutchAuctionEndingBidPriceDivisor(): %v", err)
	}

	cutils.Section("BIDDING / PRICES")
	cutils.PrintKeyValueEth("NextEthBidPrice (stored)", nextBidPrice)
	cutils.PrintKeyValueEth("NextEthBidPrice (auction)", bidPriceAuction)
	fmt.Printf("%-28s= %s CST\n", "NextCstBidPrice", cutils.WeiToEth(cstPrice))
	cutils.PrintKeyValue("ETH bid price increase div", fmt.Sprintf("%v (%.2f%%)", ethBidPriceIncreaseDivisor, priceIncrease))
	fmt.Printf("%-28s= %s CST\n", "CST reward per bid", cutils.WeiToEth(cstReward))
	cutils.PrintKeyValue("ETH Dutch auction elapsed/total", fmt.Sprintf("%v / %v", ethAuctionElapsed.String(), ethAuctionDuration.String()))
	cutils.PrintKeyValueEth("ETH Dutch auction begin price", ethDutchBeginPrice)
	cutils.PrintKeyValue("ETH Dutch auction end divisor", ethDutchEndingDivisor)
	cutils.PrintKeyValue("CST Dutch auction elapsed/total", fmt.Sprintf("%v / %v", cstAuctionElapsed.String(), cstAuctionDuration.String()))
	fmt.Printf("%-28s= %s CST\n", "CST Dutch auction begin price", cutils.WeiToEth(cstDutchBeginPrice))

	// ==================== CURRENT BIDDERS / CHAMPIONS ====================
	lastBidder, err := cosmicGame.LastBidderAddress(copts)
	if err != nil {
		cutils.Fatal("Error at LastBidderAddress(): %v", err)
	}
	lastCstBidder, err := cosmicGame.LastCstBidderAddress(copts)
	if err != nil {
		cutils.Fatal("Error at LastCstBidderAddress(): %v", err)
	}
	currentChampions, err := cosmicGame.TryGetCurrentChampions(copts)
	if err != nil {
		cutils.Fatal("Error at TryGetCurrentChampions(): %v", err)
	}
	enduranceStartTs, err := cosmicGame.EnduranceChampionStartTimeStamp(copts)
	if err != nil {
		cutils.Fatal("Error at EnduranceChampionStartTimeStamp(): %v", err)
	}
	prevEnduranceDuration, err := cosmicGame.PrevEnduranceChampionDuration(copts)
	if err != nil {
		cutils.Fatal("Error at PrevEnduranceChampionDuration(): %v", err)
	}

	cutils.Section("CURRENT BIDDERS / CHAMPIONS")
	cutils.PrintKeyValue("Last ETH bidder", lastBidder.String())
	cutils.PrintKeyValue("Last CST bidder", lastCstBidder.String())
	cutils.PrintKeyValue("Endurance champion", currentChampions.EnduranceChampionAddress.String())
	cutils.PrintKeyValueDuration("Endurance champion duration", currentChampions.EnduranceChampionDuration.Int64())
	cutils.PrintKeyValue("Endurance champion start ts", enduranceStartTs.String())
	cutils.PrintKeyValueDuration("Prev endurance champion dur", prevEnduranceDuration.Int64())
	cutils.PrintKeyValue("Chrono Warrior", currentChampions.ChronoWarriorAddress.String())
	cutils.PrintKeyValueDuration("Chrono Warrior duration", currentChampions.ChronoWarriorDuration.Int64())

	// ==================== PRIZE DISTRIBUTION ====================
	balance, err := net.Client.BalanceAt(context.Background(), cosmicGameAddr, nil)
	if err != nil {
		cutils.Fatal("Error at BalanceAt(): %v", err)
	}
	prizeAmount, err := cosmicGame.GetMainEthPrizeAmount(copts)
	if err != nil {
		cutils.Fatal("Error at GetMainEthPrizeAmount(): %v", err)
	}
	mainPrizePercentage, err := cosmicGame.MainEthPrizeAmountPercentage(copts)
	if err != nil {
		cutils.Fatal("Error at MainEthPrizeAmountPercentage(): %v", err)
	}
	charityAmount, err := cosmicGame.GetCharityEthDonationAmount(copts)
	if err != nil {
		cutils.Fatal("Error at GetCharityEthDonationAmount(): %v", err)
	}
	charityPercentage, err := cosmicGame.CharityEthDonationAmountPercentage(copts)
	if err != nil {
		cutils.Fatal("Error at CharityEthDonationAmountPercentage(): %v", err)
	}
	raffleAmount, err := cosmicGame.GetRaffleTotalEthPrizeAmountForBidders(copts)
	if err != nil {
		cutils.Fatal("Error at GetRaffleTotalEthPrizeAmountForBidders(): %v", err)
	}
	rafflePercentage, err := cosmicGame.RaffleTotalEthPrizeAmountForBiddersPercentage(copts)
	if err != nil {
		cutils.Fatal("Error at RaffleTotalEthPrizeAmountForBiddersPercentage(): %v", err)
	}
	chronoAmount, err := cosmicGame.GetChronoWarriorEthPrizeAmount(copts)
	if err != nil {
		cutils.Fatal("Error at GetChronoWarriorEthPrizeAmount(): %v", err)
	}
	chronoPercentage, err := cosmicGame.ChronoWarriorEthPrizeAmountPercentage(copts)
	if err != nil {
		cutils.Fatal("Error at ChronoWarriorEthPrizeAmountPercentage(): %v", err)
	}
	stakingAmount, err := cosmicGame.GetCosmicSignatureNftStakingTotalEthRewardAmount(copts)
	if err != nil {
		cutils.Fatal("Error at GetCosmicSignatureNftStakingTotalEthRewardAmount(): %v", err)
	}
	stakingPercentage, err := cosmicGame.CosmicSignatureNftStakingTotalEthRewardAmountPercentage(copts)
	if err != nil {
		cutils.Fatal("Error at CosmicSignatureNftStakingTotalEthRewardAmountPercentage(): %v", err)
	}
	cstPrizeAmount, err := cosmicGame.CstPrizeAmount(copts)
	if err != nil {
		cutils.Fatal("Error at CstPrizeAmount(): %v", err)
	}

	cutils.Section("PRIZE DISTRIBUTION")
	cutils.PrintKeyValueEth("Contract balance", balance)
	fmt.Printf("%-28s= %s ETH (%v%%)\n", "Main prize amount", cutils.WeiToEth(prizeAmount), mainPrizePercentage)
	fmt.Printf("%-28s= %s ETH (%v%%)\n", "Charity amount", cutils.WeiToEth(charityAmount), charityPercentage)
	fmt.Printf("%-28s= %s ETH (%v%%)\n", "Raffle amount (bidders)", cutils.WeiToEth(raffleAmount), rafflePercentage)
	fmt.Printf("%-28s= %s ETH (%v%%)\n", "Chrono Warrior amount", cutils.WeiToEth(chronoAmount), chronoPercentage)
	fmt.Printf("%-28s= %s ETH (%v%%)\n", "Staking reward amount", cutils.WeiToEth(stakingAmount), stakingPercentage)
	fmt.Printf("%-28s= %s CST\n", "CST prize amount", cutils.WeiToEth(cstPrizeAmount))

	// ==================== RAFFLE CONFIG ====================
	numEthWinners, err := cosmicGame.NumRaffleEthPrizesForBidders(copts)
	if err != nil {
		cutils.Fatal("Error at NumRaffleEthPrizesForBidders(): %v", err)
	}
	nftBidders, err := cosmicGame.NumRaffleCosmicSignatureNftsForBidders(copts)
	if err != nil {
		cutils.Fatal("Error at NumRaffleCosmicSignatureNftsForBidders(): %v", err)
	}
	nftStakers, err := cosmicGame.NumRaffleCosmicSignatureNftsForRandomWalkNftStakers(copts)
	if err != nil {
		cutils.Fatal("Error at NumRaffleCosmicSignatureNftsForRandomWalkNftStakers(): %v", err)
	}

	cutils.Section("RAFFLE CONFIG")
	cutils.PrintKeyValue("Num ETH raffle winners", numEthWinners)
	cutils.PrintKeyValue("Num NFT winners (bidders)", nftBidders)
	cutils.PrintKeyValue("Num NFT winners (RW stakers)", nftStakers)

	// ==================== PRIZES WALLET ====================
	prizeWalletAddr, err := cosmicGame.PrizesWallet(copts)
	if err != nil {
		cutils.Fatal("Error at PrizesWallet(): %v", err)
	}
	prizesWallet, err := NewPrizesWallet(prizeWalletAddr, net.Client)
	if err != nil {
		cutils.Fatal("Error at instantiation of PrizesWallet: %v", err)
	}
	numDonatedNfts, err := prizesWallet.NextDonatedNftIndex(copts)
	if err != nil {
		cutils.Fatal("Error at NextDonatedNftIndex(): %v", err)
	}
	timeoutClaim, err := prizesWallet.TimeoutDurationToWithdrawPrizes(copts)
	if err != nil {
		cutils.Fatal("Error at TimeoutDurationToWithdrawPrizes(): %v", err)
	}

	cutils.Section("PRIZES WALLET")
	cutils.PrintKeyValue("PrizesWallet address", prizeWalletAddr.String())
	cutils.PrintKeyValue("Num donated NFTs", numDonatedNfts.String())
	cutils.PrintKeyValueDuration("Timeout to withdraw prizes", timeoutClaim.Int64())

	// ==================== CHARITY ====================
	charityAddr, err := cosmicGame.CharityAddress(copts)
	if err != nil {
		cutils.Fatal("Error at CharityAddress(): %v", err)
	}
	var charityDonationRecipient string
	charityWalletCtrct, err := NewCharityWallet(charityAddr, net.Client)
	if err != nil {
		cutils.Fatal("Failed to instantiate CharityWallet contract: %v", err)
	}
	charityRecvAddr, err := charityWalletCtrct.CharityAddress(copts)
	if err != nil {
		cutils.Fatal("Error calling CharityAddress(): %v", err)
	}
	charityDonationRecipient = charityRecvAddr.String()

	cutils.Section("CHARITY")
	cutils.PrintKeyValue("CharityWallet address", charityAddr.String())
	cutils.PrintKeyValue("Charity donation receiver", charityDonationRecipient)

	// ==================== CONTRACT ADDRESSES ====================
	nftAddr, err := cosmicGame.Nft(copts)
	if err != nil {
		cutils.Fatal("Error at Nft(): %v", err)
	}
	tokenAddr, err := cosmicGame.Token(copts)
	if err != nil {
		cutils.Fatal("Error at Token(): %v", err)
	}
	randomwalkAddr, err := cosmicGame.RandomWalkNft(copts)
	if err != nil {
		cutils.Fatal("Error at RandomWalkNft(): %v", err)
	}
	stakingAddrCst, err := cosmicGame.StakingWalletCosmicSignatureNft(copts)
	if err != nil {
		cutils.Fatal("Error at StakingWalletCosmicSignatureNft(): %v", err)
	}
	stakingAddrRwalk, err := cosmicGame.StakingWalletRandomWalkNft(copts)
	if err != nil {
		cutils.Fatal("Error at StakingWalletRandomWalkNft(): %v", err)
	}
	marketingAddr, err := cosmicGame.MarketingWallet(copts)
	if err != nil {
		cutils.Fatal("Error at MarketingWallet(): %v", err)
	}
	marketingCstAmount, err := cosmicGame.MarketingWalletCstContributionAmount(copts)
	if err != nil {
		cutils.Fatal("Error at MarketingWalletCstContributionAmount(): %v", err)
	}
	ownerAddr, err := cosmicGame.Owner(copts)
	if err != nil {
		cutils.Fatal("Error at Owner(): %v", err)
	}

	cutils.Section("CONTRACT ADDRESSES")
	cutils.PrintKeyValue("CosmicSignatureNft", nftAddr.String())
	cutils.PrintKeyValue("CosmicSignatureToken", tokenAddr.String())
	cutils.PrintKeyValue("RandomWalkNft", randomwalkAddr.String())
	cutils.PrintKeyValue("StakingWallet (CST NFT)", stakingAddrCst.String())
	cutils.PrintKeyValue("StakingWallet (RandomWalk)", stakingAddrRwalk.String())
	cutils.PrintKeyValue("MarketingWallet", marketingAddr.String())
	fmt.Printf("%-28s= %s CST\n", "MarketingWallet CST contrib", cutils.WeiToEth(marketingCstAmount))
	cutils.PrintKeyValue("Owner", ownerAddr.String())

	// ==================== CONFIG PARAMETERS ====================
	bidMsgMaxLen, err := cosmicGame.BidMessageLengthMaxLimit(copts)
	if err != nil {
		cutils.Fatal("Error at BidMessageLengthMaxLimit(): %v", err)
	}

	cutils.Section("CONFIG PARAMETERS")
	cutils.PrintKeyValue("Bid message max length", bidMsgMaxLen)
}
