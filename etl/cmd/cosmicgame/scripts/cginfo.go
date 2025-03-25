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

	time_until_prize,err := cosmic_game_ctrct.GetDurationUntilMainPrize(&copts)
	if err != nil {
		fmt.Printf("Error at TimeUntilPrize()(): %v\n",err)
		fmt.Printf("Aborting\n")
		os.Exit(1)
	}
	bid_price,err := cosmic_game_ctrct.GetNextEthBidPrice(&copts,big.NewInt(0))
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
	price_increase := convert_to_percentage(eth_bid_price_increase_divisor)
	if err != nil {
		fmt.Printf("Error at EthBidPriceIncreaseDivisor(): %v\n",err)
		os.Exit(1)
	}

	time_inc_microsec,err :=  cosmic_game_ctrct.MainPrizeTimeIncrementInMicroSeconds(&copts)
	if err != nil {
		fmt.Printf("Error at MainPrizeTimeIncrementInMicroSeconds(): %v\n",err)
		os.Exit(1)
	}
	time_increment := float64(time_inc_microsec.Int64())/float64(1000000);
	first_bid_time_increment,err := cosmic_game_ctrct.GetInitialDurationUntilMainPrize(&copts)
	if err != nil {
		fmt.Printf("Error at GetInitialDurationUntilMainPrize(): %v\n",err)
		os.Exit(1)
	}

	fmt.Printf("Time until prize = %v\n",time_until_prize.Int64())
	fmt.Printf("Bid Price = %v\n",bid_price.String())
	fmt.Printf("RoundNum = %v\n",round_num.String())
	fmt.Printf("PrizeAmount = %v\n",prize_amount.String())
	fmt.Printf("PrizePercentage = %v\n",prize_percentage.String())
	fmt.Printf("RafflePercentage = %v\n",raffle_percentage.String())
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
	fmt.Printf("Price increase (on bid) = %v\n",price_increase)
	fmt.Printf("Time increment (on claimPrize()): %v\n",time_increment)
	fmt.Printf("First bid time increment = %v\n",first_bid_time_increment.String())
}
