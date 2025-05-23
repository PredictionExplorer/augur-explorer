package main

import (
	"os"
	"fmt"
	"context"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/ethclient"

	. "github.com/PredictionExplorer/augur-explorer/contracts"
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
	fmt.Printf("Calling to contract at %v\n",cosmic_game_addr.String())

	cosmic_game_ctrct,err := NewCosmicSignatureGame(cosmic_game_addr,eclient)
	if err!=nil {
		fmt.Printf("Failed to instantiate CosmicGame contract: %v\n",err)
		os.Exit(1)
	}
	prize_wallet_addr,err := cosmic_game_ctrct.PrizesWallet(&copts);
	if err != nil {
		fmt.Printf("Error at PrizesWallet()(): %v\n",err)
		fmt.Printf("Aborting\n")
		os.Exit(1)
	}
	prizes_wallet,err := NewPrizesWallet(prize_wallet_addr,eclient);
	if err != nil {
		fmt.Printf("Error at instantiation of PrizesWallet()(): %v\n",err)
		fmt.Printf("Aborting\n")
		os.Exit(1)
	}
	cst_reward,err := cosmic_game_ctrct.CstRewardAmountForBidding(&copts)
	if err != nil {
		fmt.Printf("Error at CstRewardAmountForBidding()(): %v\n",err)
		fmt.Printf("Aborting\n")
		os.Exit(1)
	}

	time_until_prize,err := cosmic_game_ctrct.GetDurationUntilMainPrize(&copts)
	if err != nil {
		fmt.Printf("Error at TimeUntilPrize()(): %v\n",err)
		fmt.Printf("Aborting\n")
		os.Exit(1)
	}
	next_bid_price,err := cosmic_game_ctrct.NextEthBidPrice(&copts)
	if err != nil {
		fmt.Printf("Error at (next) bidprice()(): %v\n",err)
		fmt.Printf("Aborting\n")
		os.Exit(1)
	}

	bid_price_auction,err := cosmic_game_ctrct.GetNextEthBidPrice(&copts,big.NewInt(0))
	if err != nil {
		fmt.Printf("Error at BidPrice()(): %v\n",err)
		fmt.Printf("Aborting\n")
		os.Exit(1)
	}
	round_num,err := cosmic_game_ctrct.RoundNum(&copts)
	if err != nil {
		fmt.Printf("Error at RoundNum()(): %v\n",err)
		fmt.Printf("Aborting\n")
		os.Exit(1)
	}
	prize_amount,err := cosmic_game_ctrct.GetMainEthPrizeAmount(&copts)
	if err != nil {
		fmt.Printf("Error at PrizeAmount()(): %v\n",err)
		fmt.Printf("Aborting\n")
		os.Exit(1)
	}
	prize_percentage,err := cosmic_game_ctrct.GetMainEthPrizeAmount(&copts)
	if err != nil {
		fmt.Printf("Error at PrizePercentage()(): %v\n",err)
		fmt.Printf("Aborting\n")
		os.Exit(1)
	}
	raffle_percentage,err := cosmic_game_ctrct.RaffleTotalEthPrizeAmountForBiddersPercentage(&copts)
	if err != nil {
		fmt.Printf("Error at RafflePercentage()(): %v\n",err)
		fmt.Printf("Aborting\n")
		os.Exit(1)
	}
	chrono_percentage,err := cosmic_game_ctrct.ChronoWarriorEthPrizeAmountPercentage(&copts)
	if err != nil {
		fmt.Printf("Error at ChronoWarriorEthPrizeAmountPercentage()(): %v\n",err)
		fmt.Printf("Aborting\n")
		os.Exit(1)
	}
	eth_bidders,err := cosmic_game_ctrct.RaffleTotalEthPrizeAmountForBiddersPercentage(&copts)
	if err != nil {
		fmt.Printf("Error at NumRaffleETHWinnersBidding()(): %v\n",err)
		fmt.Printf("Aborting\n")
		os.Exit(1)
	}
	nft_bidders,err := cosmic_game_ctrct.NumRaffleCosmicSignatureNftsForBidders(&copts)
	if err != nil {
		fmt.Printf("Error at NumRaffleNFTWinnersBidding()(): %v\n",err)
		fmt.Printf("Aborting\n")
		os.Exit(1)
	}
	nft_stakers,err := cosmic_game_ctrct.NumRaffleCosmicSignatureNftsForRandomWalkNftStakers(&copts)
	if err != nil {
		fmt.Printf("Error at NumRaffleNFTWinnersBidding()(): %v\n",err)
		fmt.Printf("Aborting\n")
		os.Exit(1)
	}
	num_raffle_participants,err := cosmic_game_ctrct.BidderAddresses(&copts,big.NewInt(round_num.Int64()))
	if err != nil {
		fmt.Printf("Error at NumRaffleParticipants()(): %v\n",err)
		fmt.Printf("Aborting\n")
		os.Exit(1)
	}
	var charity_donation_recipient string
	charity_addr,err := cosmic_game_ctrct.CharityAddress(&copts)
	if err != nil {
		fmt.Printf("Error at Charity()(): %v\n",err)
		fmt.Printf("Aborting\n")
		os.Exit(1)
	} else {
		charity_wallet_ctrct,err := NewCharityWallet(charity_addr,eclient)
		if err!=nil {
			fmt.Printf("Failed to instantiate CharityWallet contract: %v\n",err)
			os.Exit(1)
		}
		addr,err := charity_wallet_ctrct.CharityAddress(&copts)
		if err != nil {
			fmt.Printf("Error calling CharityAddress() : %v\n",err)
			os.Exit(1)
		}
		charity_donation_recipient= addr.String()
	}
	charity_amount,err := cosmic_game_ctrct.GetCharityEthDonationAmount(&copts)
	if err != nil {
		fmt.Printf("Error at CharityAmount()(): %v\n",err)
		fmt.Printf("Aborting\n")
		os.Exit(1)
	}
	raffle_amount,err := cosmic_game_ctrct.GetRaffleTotalEthPrizeAmountForBidders(&copts)
	if err != nil {
		fmt.Printf("Error at RaffleAmount()(): %v\n",err)
		fmt.Printf("Aborting\n")
		os.Exit(1)
	}
	last_bidder,err := cosmic_game_ctrct.LastBidderAddress(&copts)
	if err != nil {
		fmt.Printf("Error at LastBidder()(): %v\n",err)
		fmt.Printf("Aborting\n")
		os.Exit(1)
	}

	prize_time,err := cosmic_game_ctrct.MainPrizeTime(&copts)
	if err != nil {
		fmt.Printf("Error at prizeTime()(): %v\n",err)
		fmt.Printf("Aborting\n")
		os.Exit(1)
	}
	balance,err := eclient.BalanceAt(context.Background(),cosmic_game_addr,nil)
	if err != nil {
		fmt.Printf("Error at balanceAt()(): %v\n",err)
		fmt.Printf("Aborting\n")
		os.Exit(1)
	}
	num_donated_nfts,err := prizes_wallet.NumDonatedNfts(&copts)
	if err != nil {
		fmt.Printf("Error at numDonatedNFTs()(): %v\n",err)
		fmt.Printf("Aborting\n")
		os.Exit(1)
	}
	timeout,err := cosmic_game_ctrct.TimeoutDurationToClaimMainPrize(&copts)
	if err != nil {
		fmt.Printf("Error at timeoutClaimPrize(): %v\n",err)
		fmt.Printf("Aborting\n")
		//os.Exit(1)
	}
	owneraddr,err := cosmic_game_ctrct.Owner(&copts)
	if err != nil {
		fmt.Printf("Error at Owner(): %v\n",err)
		fmt.Printf("Aborting\n")
		os.Exit(1)
	}
	/*
	staking_addr_cst,err := cosmic_game_ctrct.StakingWalletCosmicSignatureNft(&copts)
	if err != nil {
		fmt.Printf("Error at StakingWalletCST(): %v\n",err)
		fmt.Printf("Aborting\n")
		os.Exit(1)
	}
	staking_addr_rwalk,err := cosmic_game_ctrct.StakingWalletRandomWalkNft(&copts)
	if err != nil {
		fmt.Printf("Error at StakingWalletRWalk(): %v\n",err)
		fmt.Printf("Aborting\n")
		os.Exit(1)
	}*/
	randomwalk_addr,err := cosmic_game_ctrct.RandomWalkNft(&copts)
	if err != nil {
		fmt.Printf("Error at RandomWalk(): %v\n",err)
		fmt.Printf("Aborting\n")
		os.Exit(1)
	}
	delay_until_activation,err := cosmic_game_ctrct.DelayDurationBeforeRoundActivation(&copts)
	if err != nil {
		fmt.Printf("Error at ActivationTime()(): %v\n",err)
		fmt.Printf("Aborting\n")
		os.Exit(1)
	}
	current_champions,err := cosmic_game_ctrct.TryGetCurrentChampions(&copts)
	if err != nil {
		fmt.Printf("Error at TryGetCurrentChampions(): %v\n",err)
		os.Exit(1)
	}
	cst_auction_duration,cst_auction_elapsed,err := cosmic_game_ctrct.GetCstDutchAuctionDurations(&copts)
	if err != nil {
		fmt.Printf("Error at GetCstDutchAuctionDuration(): %v\n",err)
		os.Exit(1)
	}
	eth_auction_duration,eth_auction_elapsed,err := cosmic_game_ctrct.GetEthDutchAuctionDurations(&copts)
	if err != nil {
		fmt.Printf("Error at GetEthDutchAuctionDuration(): %v\n",err)
		os.Exit(1)
	}
	eth_bid_price_increase_divisor,err := cosmic_game_ctrct.EthBidPriceIncreaseDivisor(&copts)
	if err != nil {
		fmt.Printf("Error at EthBidPriceIncreaseDivisor(): %v\n",err)
		os.Exit(1)
	}
	fmt.Printf("eth_bid_price_increase_divisor = %v\n",eth_bid_price_increase_divisor)
	price_increase := convert_to_percentage(eth_bid_price_increase_divisor)

	initial_duration_divisor,err := cosmic_game_ctrct.InitialDurationUntilMainPrizeDivisor(&copts)
	if err != nil {
		fmt.Printf("Error at initialDurationUntilMainPrizeDivisor(): %v\n",err)
		os.Exit(1)
	}
	initial_duration_inc := convert_to_percentage(initial_duration_divisor)
	initial_duration_seconds,err := cosmic_game_ctrct.GetInitialDurationUntilMainPrize(&copts)
	if err != nil {
		fmt.Printf("Error at getInitialDurationUntilMainPrize(): %v\n",err)
		os.Exit(1)
	}

	time_inc_microsec,err :=  cosmic_game_ctrct.MainPrizeTimeIncrementInMicroSeconds(&copts)
	if err != nil {
		fmt.Printf("Error at MainPrizeTimeIncrementInMicroSeconds(): %v\n",err)
		os.Exit(1)
	}
	time_increment := float64(time_inc_microsec.Int64())/float64(1000000);
	time_inc_on_bid,err := cosmic_game_ctrct.GetMainPrizeTimeIncrement(&copts)
	if err != nil {
		fmt.Printf("Error at GetMainPrizeTimeIncrement(): %v\n",err)
		os.Exit(1)
	}


	activation_time,err := cosmic_game_ctrct.RoundActivationTime(&copts)
	if err != nil {
		fmt.Printf("Error at RoundActivationTime(): %v\n",err)
		os.Exit(1)
	}
	lblock, err := eclient.BlockByNumber(context.Background(), nil)
	if err != nil {
		fmt.Printf("Error at BlockByNumber(latest): %v\n",err)
		os.Exit(1)
	}
	block_time := int64(lblock.Time())
	secs_to_start := activation_time.Int64() - block_time;

	fmt.Printf("Time until prize = %v\n",time_until_prize.Int64())
	fmt.Printf("Next Bid Price = %v\n",fmt_eth(next_bid_price))
	fmt.Printf("Next bid price auction %v\n",fmt_eth(bid_price_auction))
	fmt.Printf("RoundNum = %v\n",round_num.String())
	fmt.Printf("PrizeAmount = %v\n",fmt_eth(prize_amount))
	fmt.Printf("PrizePercentage = %v\n",prize_percentage.String())
	fmt.Printf("RafflePercentage = %v\n",raffle_percentage.String())
	fmt.Printf("ChronoPercentage = %v\n",chrono_percentage.String())
	fmt.Printf("ETHWinnersBidding = %v\n",eth_bidders);
	fmt.Printf("NFTWinnersBidding = %v\n",nft_bidders);
	fmt.Printf("NFTWinnersStaking = %v\n",nft_stakers);
	fmt.Printf("PrizeTime = %v\n",prize_time.String());
	fmt.Printf("CharityWallet addr = %v\n",charity_addr.String());
	fmt.Printf("Charity donation receiver = %v\n",charity_donation_recipient);
	fmt.Printf("CharityAmount = %v\n",charity_amount.String())
	fmt.Printf("RaffleAmount = %v\n",raffle_amount.String())
	fmt.Printf("Last bidder = %v\n",last_bidder.String())
	fmt.Printf("Contract balance = %v\n",balance.String())
	fmt.Printf("Num donated NFTs = %v\n",num_donated_nfts.String());
	fmt.Printf("Num raffle participants = %v\n",num_raffle_participants.Int64())
	fmt.Printf("Claimprize timeout = %v\n",timeout.String())
	fmt.Printf("Owner = %v\n",owneraddr.String())
	fmt.Printf("Endurance champion = %v\n",current_champions.EnduranceChampionAddress.String())
	fmt.Printf("Endurance champion duration = %v\n",current_champions.EnduranceChampionDuration.Int64())
	fmt.Printf("ChronoWarrior = %v\n",current_champions.ChronoWarriorAddress.String())
	fmt.Printf("ChronoWarrior duration = %v\n",current_champions.ChronoWarriorDuration.Int64())
	fmt.Printf("RandomWalk addr = %v\n",randomwalk_addr.String())
	fmt.Printf("Delay until Activation = %v\n",delay_until_activation)
	fmt.Printf("CST Auction duration %v of %v\n",cst_auction_elapsed.String(),cst_auction_duration.String())
	fmt.Printf("ETH Auction duration %v of %v\n",eth_auction_elapsed.String(),eth_auction_duration.String())
	fmt.Printf("Price increase (on bid) = %v%%\n",price_increase)
	fmt.Printf("CST Reward = %v\n",fmt_eth(cst_reward))
	fmt.Printf("First bid time bump %v%% (divisor=%v)\n",initial_duration_inc,initial_duration_divisor)
	fmt.Printf("First bid time bump %v sseconds\n",initial_duration_seconds)
	fmt.Printf("Time increment (on claimPrize()): %v\n",time_increment)
	fmt.Printf("Time increment on bid: %v\n",time_inc_on_bid.Int64())
	fmt.Printf("Time increment in microseconds: %v\n",time_inc_microsec)
	fmt.Printf("Round activation time = %v\n",activation_time.String())
	fmt.Printf("Seconds to round start= %v (%v - %v)\n",secs_to_start,block_time,activation_time.Int64())
}
