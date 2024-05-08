package main

import (
	"os"
	"fmt"
	"context"
	"math/big"
	"encoding/hex"

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

	cosmic_game_ctrct,err := NewCosmicGame(cosmic_game_addr,eclient)
	if err!=nil {
		fmt.Printf("Failed to instantiate CosmicGame contract: %v\n",err)
		os.Exit(1)
	}

	time_until_prize,err := cosmic_game_ctrct.TimeUntilPrize(&copts)
	if err != nil {
		fmt.Printf("Error at TimeUntilPrize()(): %v\n",err)
		fmt.Printf("Aborting\n")
		os.Exit(1)
	}
	bid_price,err := cosmic_game_ctrct.GetBidPrice(&copts)
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
	prize_amount,err := cosmic_game_ctrct.PrizeAmount(&copts)
	if err != nil {
		fmt.Printf("Error at PrizeAmount()(): %v\n",err)
		fmt.Printf("Aborting\n")
		os.Exit(1)
	}
	prize_percentage,err := cosmic_game_ctrct.PrizePercentage(&copts)
	if err != nil {
		fmt.Printf("Error at PrizePercentage()(): %v\n",err)
		fmt.Printf("Aborting\n")
		os.Exit(1)
	}
	var charity_donation_recipient string
	charity_addr,err := cosmic_game_ctrct.Charity(&copts)
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
	charity_amount,err := cosmic_game_ctrct.CharityAmount(&copts)
	if err != nil {
		fmt.Printf("Error at CharityAmount()(): %v\n",err)
		fmt.Printf("Aborting\n")
		os.Exit(1)
	}
	raffle_amount,err := cosmic_game_ctrct.RaffleAmount(&copts)
	if err != nil {
		fmt.Printf("Error at RaffleAmount()(): %v\n",err)
		fmt.Printf("Aborting\n")
		os.Exit(1)
	}
	last_bidder,err := cosmic_game_ctrct.LastBidder(&copts)
	if err != nil {
		fmt.Printf("Error at LastBidder()(): %v\n",err)
		fmt.Printf("Aborting\n")
		os.Exit(1)
	}
	prize_time,err := cosmic_game_ctrct.PrizeTime(&copts)
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
	num_donated_nfts,err := cosmic_game_ctrct.NumDonatedNFTs(&copts)
	if err != nil {
		fmt.Printf("Error at numDonatedNFTs()(): %v\n",err)
		fmt.Printf("Aborting\n")
		os.Exit(1)
	}
	nanoseconds,err := cosmic_game_ctrct.NanoSecondsExtra(&copts)
	if err != nil {
		fmt.Printf("Error at nanoSecondsExtra()(): %v\n",err)
		fmt.Printf("Aborting\n")
		os.Exit(1)
	}
	initialseconds,err := cosmic_game_ctrct.InitialSecondsUntilPrize(&copts)
	if err != nil {
		fmt.Printf("Error at initialSecondsUntilPrize()(): %v\n",err)
		fmt.Printf("Aborting\n")
		os.Exit(1)
	}
	timeout,err := cosmic_game_ctrct.TimeoutClaimPrize(&copts)
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
	blogic_addr,err := cosmic_game_ctrct.BLogic(&copts)
	if err != nil {
		fmt.Printf("Error at bLogic(): %v\n",err)
		fmt.Printf("Aborting\n")
		os.Exit(1)
	}
	staking_addr,err := cosmic_game_ctrct.StakingWallet(&copts)
	if err != nil {
		fmt.Printf("Error at StakingWallet(): %v\n",err)
		fmt.Printf("Aborting\n")
		os.Exit(1)
	}
	last_bid_type,err := cosmic_game_ctrct.LastBidType(&copts)
	if err != nil {
		fmt.Printf("Error at LastBidType()(): %v\n",err)
		fmt.Printf("Aborting\n")
		os.Exit(1)
	}
	activation_time,err := cosmic_game_ctrct.ActivationTime(&copts)
	if err != nil {
		fmt.Printf("Error at ActivationTime()(): %v\n",err)
		fmt.Printf("Aborting\n")
		os.Exit(1)
	}
	num_eth_bids,err := cosmic_game_ctrct.NumETHBids(&copts)
	if err != nil {
		fmt.Printf("Error at NumETHBids()(): %v\n",err)
		fmt.Printf("Aborting\n")
		os.Exit(1)
	}
	num_cst_bids,err := cosmic_game_ctrct.NumCSTBids(&copts)
	if err != nil {
		fmt.Printf("Error at NumCSTBids()(): %v\n",err)
		fmt.Printf("Aborting\n")
		os.Exit(1)
	}
	cst_auction_length,err := cosmic_game_ctrct.CSTAuctionLength(&copts)
	if err != nil {
		fmt.Printf("Error at CSTAuctionLength()(): %v\n",err)
		fmt.Printf("Aborting\n")
		os.Exit(1)
	}
	system_mode,err := cosmic_game_ctrct.SystemMode(&copts)
	if err != nil {
		fmt.Printf("Error at SystemMode()(): %v\n",err)
		fmt.Printf("Aborting\n")
		os.Exit(1)
	}


	fmt.Printf("Time until prize = %v\n",time_until_prize.Int64())
	fmt.Printf("Bid Price = %v\n",bid_price.String())
	fmt.Printf("RoundNum = %v\n",round_num.String())
	fmt.Printf("PrizeAmount = %v\n",prize_amount.String())
	fmt.Printf("PrizePercentage = %v\n",prize_percentage.String())
	fmt.Printf("PrizeTime = %v\n",prize_time.String());
	fmt.Printf("CharityWallet addr = %v\n",charity_addr.String());
	fmt.Printf("Charity donation receiver = %v\n",charity_donation_recipient);
	fmt.Printf("CharityAmount = %v\n",charity_amount.String())
	fmt.Printf("RaffleAmount = %v\n",raffle_amount.String())
	fmt.Printf("Last bidder = %v\n",last_bidder.String())
	fmt.Printf("Contract balance = %v\n",balance.String())
	fmt.Printf("Num donated NFTs = %v\n",num_donated_nfts.String());
	fmt.Printf("Nanoseconds = %v\n",nanoseconds.String())
	fmt.Printf("Initial seconds = %v\n",initialseconds.String())
	fmt.Printf("Claimprize timeout = %v\n",timeout.String())
	fmt.Printf("Owner = %v\n",owneraddr.String())
	fmt.Printf("BusinessLogic = %v\n",blogic_addr.String())
	fmt.Printf("LastBidType = %v\n",last_bid_type)
	fmt.Printf("ActivationTime= %v\n",activation_time)
	fmt.Printf("NumETHBids = %v\n",num_eth_bids);
	fmt.Printf("NumCSTBids = %v\n",num_cst_bids);
	fmt.Printf("CSTAuctionLength = %v\n",cst_auction_length);
	fmt.Printf("SystemMode = %v (0-Runtime, 1-Prepare maintenance, 2-Maintenance)\n",system_mode.String());

	swallet,err := NewStakingWallet(staking_addr,eclient);
	if err != nil {
		fmt.Printf("Failed to instantiate StakingWallet contract: %v\n",err)
		os.Exit(1)
	}
	min_stake_period,err := swallet.MinStakePeriod(&copts);
	if err != nil {
		os.Exit(1)
	} else {
		fmt.Printf("StakingWallet = %v\n",staking_addr.String());
		fmt.Printf("MinStakePeriod = %v\n",min_stake_period);
	}
	blogic_ctrct,err := NewBusinessLogic(cosmic_game_addr,eclient)
	if err!=nil {
		fmt.Printf("Failed to instantiate BusinessLogic contract: %v\n",err)
		os.Exit(1)
	}
	cst_bytes,err := blogic_ctrct.CurrentCSTPrice(&copts);
	if err != nil {
		fmt.Printf("Error at CurrentCSTPrice()(): %v\n",err)
	} else {
		slice_len := len(cst_bytes);
		fmt.Printf("currentCSTPrice output length: %v\n",slice_len);
		if slice_len > 0 {
			fmt.Printf("currentCSTPrice hex value: %v\n",hex.EncodeToString(cst_bytes));
			price_hash := common.BytesToHash(cst_bytes[64:])
			cst_price := price_hash.Big()
			f_divisor := big.NewFloat(0.0).SetInt(big.NewInt(1e18))
			f_price := big.NewFloat(0.0).SetInt(cst_price)
			f_quo := big.NewFloat(0.0).Quo(f_price,f_divisor)
			bid_price_eth,_ := f_quo.Float64()
			fmt.Printf("CST Bid Price (Ether) = %.6f\n",bid_price_eth)
			fmt.Printf("CST Bid Price (Wei) = %v\n",cst_price.String())
		}
	}
	ad_bytes,err := blogic_ctrct.AuctionDuration(&copts);
	if err != nil {
		fmt.Printf("Error at AuctionDuration()(): %v\n",err)
		fmt.Printf("Aborting\n")
		os.Exit(1)
	} else {
		slice_len := len(cst_bytes);
		fmt.Printf("AuctionDuration() output length: %v\n",slice_len);
		if slice_len > 0 {
			seconds_hash := common.BytesToHash(ad_bytes[64:96])
			seconds := seconds_hash.Big().Int64()
			duration_hash := common.BytesToHash(ad_bytes[96:])
			duration := duration_hash.Big().Int64()
			fmt.Printf("CST auction elapsed time (sec): %v\n",seconds)
			fmt.Printf("CST auction duration: %v\n",duration)
		}
	}
}
