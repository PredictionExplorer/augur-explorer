package main

import (
	"os"
	"fmt"
	"context"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/ethclient"

	. "github.com/PredictionExplorer/augur-explorer/rwcg/contracts/cosmicgame"
)
const (
)
var (
	RPC_URL string
)
func fmt_eth(wei *big.Int) string {
    ether := new(big.Float).SetInt(wei)
    eth_value := new(big.Float).Quo(ether, big.NewFloat(1e18))
    return eth_value.Text('f', 18) // 18 decimal places to match Ethereum precision
}
func convert_to_percentage(in *big.Int) (float64) {

	one := big.NewFloat(1)
	hundred := big.NewFloat(100)
	divisor_float := new(big.Float).SetInt(in)
	increase_fraction := new(big.Float).Quo(one,divisor_float)
	increase_percent := new(big.Float).Mul(increase_fraction, hundred)
	out,_ := increase_percent.Float64()
	return out
}
func fmt_duration(secs int64) string {
	if secs < 0 {
		return fmt.Sprintf("%d sec (negative)", secs)
	}
	if secs < 60 {
		return fmt.Sprintf("%d sec", secs)
	}
	if secs < 3600 {
		return fmt.Sprintf("%d min %d sec", secs/60, secs%60)
	}
	hours := secs / 3600
	mins := (secs % 3600) / 60
	sec := secs % 60
	return fmt.Sprintf("%dh %dm %ds", hours, mins, sec)
}
func section(title string) {
	fmt.Printf("\n==================== %s ====================\n", title)
}
func main() {

	RPC_URL = os.Getenv("RPC_URL")
	eclient, err := ethclient.Dial(RPC_URL)
	if err!=nil {
		fmt.Printf("Can't connect to ETH RPC: %v\n",err)
		os.Exit(1)
	}

	var cg_addr string
	if len(os.Args) < 2 {
		fmt.Printf(
			"Usage: \n\t\t%v [cosmic_game_addr]\n\t\t"+
			"Gets CosmicGame read only variables\n\n",os.Args[0],
		)
		fmt.Printf("Setting default cosmic game contract address to 0x5FbDB2315678afecb367f032d93F642f64180aa3\n")
		cg_addr = "0x5FbDB2315678afecb367f032d93F642f64180aa3";
	} else {
		cg_addr = os.Args[1]
	}
	var copts bind.CallOpts
	cosmic_game_addr := common.HexToAddress(cg_addr)
	fmt.Printf("CosmicSignatureGame contract: %v\n",cosmic_game_addr.String())

	cosmic_game_ctrct,err := NewCosmicSignatureGame(cosmic_game_addr,eclient)
	if err!=nil {
		fmt.Printf("Failed to instantiate CosmicGame contract: %v\n",err)
		os.Exit(1)
	}

	// Get current block for time calculations
	lblock, err := eclient.HeaderByNumber(context.Background(), nil)
	if err != nil {
		fmt.Printf("Error at BlockByNumber(latest): %v\n",err)
		os.Exit(1)
	}
	block_time := int64(lblock.Time)

	// ==================== ROUND STATUS ====================
	round_num,err := cosmic_game_ctrct.RoundNum(&copts)
	if err != nil {
		fmt.Printf("Error at RoundNum(): %v\n",err)
		os.Exit(1)
	}
	activation_time,err := cosmic_game_ctrct.RoundActivationTime(&copts)
	if err != nil {
		fmt.Printf("Error at RoundActivationTime(): %v\n",err)
		os.Exit(1)
	}
	secs_to_start := activation_time.Int64() - block_time
	delay_until_activation,err := cosmic_game_ctrct.DelayDurationBeforeRoundActivation(&copts)
	if err != nil {
		fmt.Printf("Error at DelayDurationBeforeRoundActivation(): %v\n",err)
		os.Exit(1)
	}
	total_bids,err := cosmic_game_ctrct.GetTotalNumBids(&copts, round_num)
	if err != nil {
		fmt.Printf("Error at GetTotalNumBids(): %v\n",err)
		os.Exit(1)
	}
	num_raffle_participants,err := cosmic_game_ctrct.BidderAddresses(&copts,big.NewInt(round_num.Int64()))
	if err != nil {
		fmt.Printf("Error at BidderAddresses(): %v\n",err)
		os.Exit(1)
	}

	section("ROUND STATUS")
	fmt.Printf("RoundNum                       = %v\n",round_num.String())
	fmt.Printf("Round activation time          = %v\n",activation_time.String())
	fmt.Printf("Block time (current)           = %v\n",block_time)
	fmt.Printf("Delay before activation        = %v\n",delay_until_activation)
	if secs_to_start > 0 {
		fmt.Printf("Round status                   = INACTIVE - activates in %s\n", fmt_duration(secs_to_start))
	} else {
		fmt.Printf("Round status                   = ACTIVE - started %s ago\n", fmt_duration(-secs_to_start))
	}
	fmt.Printf("Total bids this round          = %v\n",total_bids.String())
	fmt.Printf("Unique bidders this round      = %v\n",num_raffle_participants.Int64())

	// ==================== TIMING / COUNTDOWN ====================
	time_until_prize,err := cosmic_game_ctrct.GetDurationUntilMainPrize(&copts)
	if err != nil {
		fmt.Printf("Error at GetDurationUntilMainPrize(): %v\n",err)
		os.Exit(1)
	}
	duration_until_prize_raw,err := cosmic_game_ctrct.GetDurationUntilMainPrizeRaw(&copts)
	if err != nil {
		fmt.Printf("Error at GetDurationUntilMainPrizeRaw(): %v\n",err)
		os.Exit(1)
	}
	prize_time,err := cosmic_game_ctrct.MainPrizeTime(&copts)
	if err != nil {
		fmt.Printf("Error at MainPrizeTime(): %v\n",err)
		os.Exit(1)
	}
	timeout_main_prize,err := cosmic_game_ctrct.TimeoutDurationToClaimMainPrize(&copts)
	if err != nil {
		fmt.Printf("Error at TimeoutDurationToClaimMainPrize(): %v\n",err)
		os.Exit(1)
	}
	duration_until_anyone_can_claim := new(big.Int).Add(duration_until_prize_raw, timeout_main_prize)
	anyone_can_claim := duration_until_anyone_can_claim.Cmp(big.NewInt(0)) <= 0

	time_inc_microsec,err := cosmic_game_ctrct.MainPrizeTimeIncrementInMicroSeconds(&copts)
	if err != nil {
		fmt.Printf("Error at MainPrizeTimeIncrementInMicroSeconds(): %v\n",err)
		os.Exit(1)
	}
	time_increment := float64(time_inc_microsec.Int64())/float64(1000000)
	time_inc_on_bid,err := cosmic_game_ctrct.GetMainPrizeTimeIncrement(&copts)
	if err != nil {
		fmt.Printf("Error at GetMainPrizeTimeIncrement(): %v\n",err)
		os.Exit(1)
	}
	time_inc_increase_divisor,err := cosmic_game_ctrct.MainPrizeTimeIncrementIncreaseDivisor(&copts)
	if err != nil {
		fmt.Printf("Error at MainPrizeTimeIncrementIncreaseDivisor(): %v\n",err)
		os.Exit(1)
	}
	initial_duration_divisor,err := cosmic_game_ctrct.InitialDurationUntilMainPrizeDivisor(&copts)
	if err != nil {
		fmt.Printf("Error at InitialDurationUntilMainPrizeDivisor(): %v\n",err)
		os.Exit(1)
	}
	initial_duration_inc := convert_to_percentage(initial_duration_divisor)
	initial_duration_seconds,err := cosmic_game_ctrct.GetInitialDurationUntilMainPrize(&copts)
	if err != nil {
		fmt.Printf("Error at GetInitialDurationUntilMainPrize(): %v\n",err)
		os.Exit(1)
	}

	section("TIMING / COUNTDOWN")
	fmt.Printf("MainPrizeTime (timestamp)      = %v\n",prize_time.String())
	fmt.Printf("Duration until prize (clamped) = %v (%s)\n",time_until_prize.Int64(), fmt_duration(time_until_prize.Int64()))
	fmt.Printf("Duration until prize (raw)     = %v (%s)\n",duration_until_prize_raw.Int64(), fmt_duration(duration_until_prize_raw.Int64()))
	fmt.Printf("Timeout for last bidder claim  = %v (%s)\n",timeout_main_prize.String(), fmt_duration(timeout_main_prize.Int64()))
	fmt.Printf("Duration until anyone can claim= %v (%s)\n",duration_until_anyone_can_claim.Int64(), fmt_duration(duration_until_anyone_can_claim.Int64()))
	if anyone_can_claim {
		fmt.Printf("Can anyone claim now?          = YES\n")
	} else {
		fmt.Printf("Can anyone claim now?          = NO\n")
	}
	fmt.Printf("Time increment (microsec)      = %v\n",time_inc_microsec)
	fmt.Printf("Time increment (seconds)       = %v\n",time_increment)
	fmt.Printf("Time increment on bid (current)= %v sec\n",time_inc_on_bid.Int64())
	fmt.Printf("Time increment increase divisor= %v\n",time_inc_increase_divisor)
	fmt.Printf("First bid time bump            = %v%% (divisor=%v)\n",initial_duration_inc,initial_duration_divisor)
	fmt.Printf("First bid time bump (seconds)  = %v\n",initial_duration_seconds)

	// ==================== BIDDING / PRICES ====================
	next_bid_price,err := cosmic_game_ctrct.NextEthBidPrice(&copts)
	if err != nil {
		fmt.Printf("Error at NextEthBidPrice(): %v\n",err)
		os.Exit(1)
	}
	bid_price_auction,err := cosmic_game_ctrct.GetNextEthBidPrice(&copts)
	if err != nil {
		fmt.Printf("Error at GetNextEthBidPrice(): %v\n",err)
		os.Exit(1)
	}
	cst_price,err := cosmic_game_ctrct.GetNextCstBidPrice(&copts)
	if err != nil {
		fmt.Printf("Error at GetNextCstBidPrice(): %v\n",err)
		os.Exit(1)
	}
	eth_bid_price_increase_divisor,err := cosmic_game_ctrct.EthBidPriceIncreaseDivisor(&copts)
	if err != nil {
		fmt.Printf("Error at EthBidPriceIncreaseDivisor(): %v\n",err)
		os.Exit(1)
	}
	price_increase := convert_to_percentage(eth_bid_price_increase_divisor)
	cst_reward,err := cosmic_game_ctrct.CstRewardAmountForBidding(&copts)
	if err != nil {
		fmt.Printf("Error at CstRewardAmountForBidding(): %v\n",err)
		os.Exit(1)
	}

	// Dutch auction info
	cst_auction_duration,cst_auction_elapsed,err := cosmic_game_ctrct.GetCstDutchAuctionDurations(&copts)
	if err != nil {
		fmt.Printf("Error at GetCstDutchAuctionDurations(): %v\n",err)
		os.Exit(1)
	}
	eth_auction_duration,eth_auction_elapsed,err := cosmic_game_ctrct.GetEthDutchAuctionDurations(&copts)
	if err != nil {
		fmt.Printf("Error at GetEthDutchAuctionDurations(): %v\n",err)
		os.Exit(1)
	}
	cst_dutch_begin_price,err := cosmic_game_ctrct.CstDutchAuctionBeginningBidPrice(&copts)
	if err != nil {
		fmt.Printf("Error at CstDutchAuctionBeginningBidPrice(): %v\n",err)
		os.Exit(1)
	}
	eth_dutch_begin_price,err := cosmic_game_ctrct.EthDutchAuctionBeginningBidPrice(&copts)
	if err != nil {
		fmt.Printf("Error at EthDutchAuctionBeginningBidPrice(): %v\n",err)
		os.Exit(1)
	}
	eth_dutch_ending_divisor,err := cosmic_game_ctrct.EthDutchAuctionEndingBidPriceDivisor(&copts)
	if err != nil {
		fmt.Printf("Error at EthDutchAuctionEndingBidPriceDivisor(): %v\n",err)
		os.Exit(1)
	}

	section("BIDDING / PRICES")
	fmt.Printf("NextEthBidPrice (stored)       = %v ETH\n",fmt_eth(next_bid_price))
	fmt.Printf("NextEthBidPrice (auction)      = %v ETH\n",fmt_eth(bid_price_auction))
	fmt.Printf("NextCstBidPrice                = %v CST\n",fmt_eth(cst_price))
	fmt.Printf("ETH bid price increase divisor = %v (%.2f%%)\n",eth_bid_price_increase_divisor,price_increase)
	fmt.Printf("CST reward per bid             = %v CST\n",fmt_eth(cst_reward))
	fmt.Printf("ETH Dutch auction elapsed/total= %v / %v\n",eth_auction_elapsed.String(),eth_auction_duration.String())
	fmt.Printf("ETH Dutch auction begin price  = %v ETH\n",fmt_eth(eth_dutch_begin_price))
	fmt.Printf("ETH Dutch auction end divisor  = %v\n",eth_dutch_ending_divisor)
	fmt.Printf("CST Dutch auction elapsed/total= %v / %v\n",cst_auction_elapsed.String(),cst_auction_duration.String())
	fmt.Printf("CST Dutch auction begin price  = %v CST\n",fmt_eth(cst_dutch_begin_price))

	// ==================== CURRENT BIDDERS / CHAMPIONS ====================
	last_bidder,err := cosmic_game_ctrct.LastBidderAddress(&copts)
	if err != nil {
		fmt.Printf("Error at LastBidderAddress(): %v\n",err)
		os.Exit(1)
	}
	last_cst_bidder,err := cosmic_game_ctrct.LastCstBidderAddress(&copts)
	if err != nil {
		fmt.Printf("Error at LastCstBidderAddress(): %v\n",err)
		os.Exit(1)
	}
	current_champions,err := cosmic_game_ctrct.TryGetCurrentChampions(&copts)
	if err != nil {
		fmt.Printf("Error at TryGetCurrentChampions(): %v\n",err)
		os.Exit(1)
	}
	endurance_start_ts,err := cosmic_game_ctrct.EnduranceChampionStartTimeStamp(&copts)
	if err != nil {
		fmt.Printf("Error at EnduranceChampionStartTimeStamp(): %v\n",err)
		os.Exit(1)
	}
	prev_endurance_duration,err := cosmic_game_ctrct.PrevEnduranceChampionDuration(&copts)
	if err != nil {
		fmt.Printf("Error at PrevEnduranceChampionDuration(): %v\n",err)
		os.Exit(1)
	}

	section("CURRENT BIDDERS / CHAMPIONS")
	fmt.Printf("Last ETH bidder                = %v\n",last_bidder.String())
	fmt.Printf("Last CST bidder                = %v\n",last_cst_bidder.String())
	fmt.Printf("Endurance champion             = %v\n",current_champions.EnduranceChampionAddress.String())
	fmt.Printf("Endurance champion duration    = %v (%s)\n",current_champions.EnduranceChampionDuration.Int64(), fmt_duration(current_champions.EnduranceChampionDuration.Int64()))
	fmt.Printf("Endurance champion start ts    = %v\n",endurance_start_ts.String())
	fmt.Printf("Prev endurance champion dur    = %v (%s)\n",prev_endurance_duration.Int64(), fmt_duration(prev_endurance_duration.Int64()))
	fmt.Printf("Chrono Warrior                 = %v\n",current_champions.ChronoWarriorAddress.String())
	fmt.Printf("Chrono Warrior duration        = %v (%s)\n",current_champions.ChronoWarriorDuration.Int64(), fmt_duration(current_champions.ChronoWarriorDuration.Int64()))

	// ==================== PRIZE DISTRIBUTION ====================
	balance,err := eclient.BalanceAt(context.Background(),cosmic_game_addr,nil)
	if err != nil {
		fmt.Printf("Error at BalanceAt(): %v\n",err)
		os.Exit(1)
	}
	prize_amount,err := cosmic_game_ctrct.GetMainEthPrizeAmount(&copts)
	if err != nil {
		fmt.Printf("Error at GetMainEthPrizeAmount(): %v\n",err)
		os.Exit(1)
	}
	main_prize_percentage,err := cosmic_game_ctrct.MainEthPrizeAmountPercentage(&copts)
	if err != nil {
		fmt.Printf("Error at MainEthPrizeAmountPercentage(): %v\n",err)
		os.Exit(1)
	}
	charity_amount,err := cosmic_game_ctrct.GetCharityEthDonationAmount(&copts)
	if err != nil {
		fmt.Printf("Error at GetCharityEthDonationAmount(): %v\n",err)
		os.Exit(1)
	}
	charity_percentage,err := cosmic_game_ctrct.CharityEthDonationAmountPercentage(&copts)
	if err != nil {
		fmt.Printf("Error at CharityEthDonationAmountPercentage(): %v\n",err)
		os.Exit(1)
	}
	raffle_amount,err := cosmic_game_ctrct.GetRaffleTotalEthPrizeAmountForBidders(&copts)
	if err != nil {
		fmt.Printf("Error at GetRaffleTotalEthPrizeAmountForBidders(): %v\n",err)
		os.Exit(1)
	}
	raffle_percentage,err := cosmic_game_ctrct.RaffleTotalEthPrizeAmountForBiddersPercentage(&copts)
	if err != nil {
		fmt.Printf("Error at RaffleTotalEthPrizeAmountForBiddersPercentage(): %v\n",err)
		os.Exit(1)
	}
	chrono_amount,err := cosmic_game_ctrct.GetChronoWarriorEthPrizeAmount(&copts)
	if err != nil {
		fmt.Printf("Error at GetChronoWarriorEthPrizeAmount(): %v\n",err)
		os.Exit(1)
	}
	chrono_percentage,err := cosmic_game_ctrct.ChronoWarriorEthPrizeAmountPercentage(&copts)
	if err != nil {
		fmt.Printf("Error at ChronoWarriorEthPrizeAmountPercentage(): %v\n",err)
		os.Exit(1)
	}
	staking_amount,err := cosmic_game_ctrct.GetCosmicSignatureNftStakingTotalEthRewardAmount(&copts)
	if err != nil {
		fmt.Printf("Error at GetCosmicSignatureNftStakingTotalEthRewardAmount(): %v\n",err)
		os.Exit(1)
	}
	staking_percentage,err := cosmic_game_ctrct.CosmicSignatureNftStakingTotalEthRewardAmountPercentage(&copts)
	if err != nil {
		fmt.Printf("Error at CosmicSignatureNftStakingTotalEthRewardAmountPercentage(): %v\n",err)
		os.Exit(1)
	}
	cst_prize_amount,err := cosmic_game_ctrct.CstPrizeAmount(&copts)
	if err != nil {
		fmt.Printf("Error at CstPrizeAmount(): %v\n",err)
		os.Exit(1)
	}

	section("PRIZE DISTRIBUTION")
	fmt.Printf("Contract balance               = %v ETH\n",fmt_eth(balance))
	fmt.Printf("Main prize amount              = %v ETH (%v%%)\n",fmt_eth(prize_amount),main_prize_percentage)
	fmt.Printf("Charity amount                 = %v ETH (%v%%)\n",fmt_eth(charity_amount),charity_percentage)
	fmt.Printf("Raffle amount (bidders)        = %v ETH (%v%%)\n",fmt_eth(raffle_amount),raffle_percentage)
	fmt.Printf("Chrono Warrior amount          = %v ETH (%v%%)\n",fmt_eth(chrono_amount),chrono_percentage)
	fmt.Printf("Staking reward amount          = %v ETH (%v%%)\n",fmt_eth(staking_amount),staking_percentage)
	fmt.Printf("CST prize amount               = %v CST\n",fmt_eth(cst_prize_amount))

	// ==================== RAFFLE CONFIG ====================
	num_eth_winners,err := cosmic_game_ctrct.NumRaffleEthPrizesForBidders(&copts)
	if err != nil {
		fmt.Printf("Error at NumRaffleEthPrizesForBidders(): %v\n",err)
		os.Exit(1)
	}
	nft_bidders,err := cosmic_game_ctrct.NumRaffleCosmicSignatureNftsForBidders(&copts)
	if err != nil {
		fmt.Printf("Error at NumRaffleCosmicSignatureNftsForBidders(): %v\n",err)
		os.Exit(1)
	}
	nft_stakers,err := cosmic_game_ctrct.NumRaffleCosmicSignatureNftsForRandomWalkNftStakers(&copts)
	if err != nil {
		fmt.Printf("Error at NumRaffleCosmicSignatureNftsForRandomWalkNftStakers(): %v\n",err)
		os.Exit(1)
	}

	section("RAFFLE CONFIG")
	fmt.Printf("Num ETH raffle winners         = %v\n",num_eth_winners)
	fmt.Printf("Num NFT winners (bidders)      = %v\n",nft_bidders)
	fmt.Printf("Num NFT winners (RW stakers)   = %v\n",nft_stakers)

	// ==================== PRIZES WALLET ====================
	prize_wallet_addr,err := cosmic_game_ctrct.PrizesWallet(&copts)
	if err != nil {
		fmt.Printf("Error at PrizesWallet(): %v\n",err)
		os.Exit(1)
	}
	prizes_wallet,err := NewPrizesWallet(prize_wallet_addr,eclient)
	if err != nil {
		fmt.Printf("Error at instantiation of PrizesWallet: %v\n",err)
		os.Exit(1)
	}
	num_donated_nfts,err := prizes_wallet.NextDonatedNftIndex(&copts)
	if err != nil {
		fmt.Printf("Error at NextDonatedNftIndex(): %v\n",err)
		os.Exit(1)
	}
	timeout_claim,err := prizes_wallet.TimeoutDurationToWithdrawPrizes(&copts)
	if err != nil {
		fmt.Printf("Error at TimeoutDurationToWithdrawPrizes(): %v\n",err)
		os.Exit(1)
	}

	section("PRIZES WALLET")
	fmt.Printf("PrizesWallet address           = %v\n",prize_wallet_addr.String())
	fmt.Printf("Num donated NFTs               = %v\n",num_donated_nfts.String())
	fmt.Printf("Timeout to withdraw prizes     = %v (%s)\n",timeout_claim.String(), fmt_duration(timeout_claim.Int64()))

	// ==================== CHARITY ====================
	charity_addr,err := cosmic_game_ctrct.CharityAddress(&copts)
	if err != nil {
		fmt.Printf("Error at CharityAddress(): %v\n",err)
		os.Exit(1)
	}
	var charity_donation_recipient string
	charity_wallet_ctrct,err := NewCharityWallet(charity_addr,eclient)
	if err!=nil {
		fmt.Printf("Failed to instantiate CharityWallet contract: %v\n",err)
		os.Exit(1)
	}
	charity_recv_addr,err := charity_wallet_ctrct.CharityAddress(&copts)
	if err != nil {
		fmt.Printf("Error calling CharityAddress(): %v\n",err)
		os.Exit(1)
	}
	charity_donation_recipient = charity_recv_addr.String()

	section("CHARITY")
	fmt.Printf("CharityWallet address          = %v\n",charity_addr.String())
	fmt.Printf("Charity donation receiver      = %v\n",charity_donation_recipient)

	// ==================== CONTRACT ADDRESSES ====================
	nft_addr,err := cosmic_game_ctrct.Nft(&copts)
	if err != nil {
		fmt.Printf("Error at Nft(): %v\n",err)
		os.Exit(1)
	}
	token_addr,err := cosmic_game_ctrct.Token(&copts)
	if err != nil {
		fmt.Printf("Error at Token(): %v\n",err)
		os.Exit(1)
	}
	randomwalk_addr,err := cosmic_game_ctrct.RandomWalkNft(&copts)
	if err != nil {
		fmt.Printf("Error at RandomWalkNft(): %v\n",err)
		os.Exit(1)
	}
	staking_addr_cst,err := cosmic_game_ctrct.StakingWalletCosmicSignatureNft(&copts)
	if err != nil {
		fmt.Printf("Error at StakingWalletCosmicSignatureNft(): %v\n",err)
		os.Exit(1)
	}
	staking_addr_rwalk,err := cosmic_game_ctrct.StakingWalletRandomWalkNft(&copts)
	if err != nil {
		fmt.Printf("Error at StakingWalletRandomWalkNft(): %v\n",err)
		os.Exit(1)
	}
	marketing_addr,err := cosmic_game_ctrct.MarketingWallet(&copts)
	if err != nil {
		fmt.Printf("Error at MarketingWallet(): %v\n",err)
		os.Exit(1)
	}
	marketing_cst_amount,err := cosmic_game_ctrct.MarketingWalletCstContributionAmount(&copts)
	if err != nil {
		fmt.Printf("Error at MarketingWalletCstContributionAmount(): %v\n",err)
		os.Exit(1)
	}
	owneraddr,err := cosmic_game_ctrct.Owner(&copts)
	if err != nil {
		fmt.Printf("Error at Owner(): %v\n",err)
		os.Exit(1)
	}

	section("CONTRACT ADDRESSES")
	fmt.Printf("CosmicSignatureNft             = %v\n",nft_addr.String())
	fmt.Printf("CosmicSignatureToken           = %v\n",token_addr.String())
	fmt.Printf("RandomWalkNft                  = %v\n",randomwalk_addr.String())
	fmt.Printf("StakingWallet (CST NFT)        = %v\n",staking_addr_cst.String())
	fmt.Printf("StakingWallet (RandomWalk)     = %v\n",staking_addr_rwalk.String())
	fmt.Printf("MarketingWallet                = %v\n",marketing_addr.String())
	fmt.Printf("MarketingWallet CST contrib    = %v CST\n",fmt_eth(marketing_cst_amount))
	fmt.Printf("Owner                          = %v\n",owneraddr.String())

	// ==================== CONFIG PARAMETERS ====================
	bid_msg_max_len,err := cosmic_game_ctrct.BidMessageLengthMaxLimit(&copts)
	if err != nil {
		fmt.Printf("Error at BidMessageLengthMaxLimit(): %v\n",err)
		os.Exit(1)
	}

	section("CONFIG PARAMETERS")
	fmt.Printf("Bid message max length         = %v\n",bid_msg_max_len)
}
