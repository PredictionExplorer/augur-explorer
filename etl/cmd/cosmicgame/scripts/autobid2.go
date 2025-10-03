// A basic code for a bot that bids and claims main prize when it is the lastBidder
// Design note: the event programming model is a must, to properly handle each case differntly when sequences are reusable (this bot is very draft yet, not for productive use, still has many points for improvement)
package main

import (
	"os"
	"fmt"
	"time"
	"bytes"
	"math/big"
	"context"
	"crypto/ecdsa"
	"encoding/hex"

	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	ethrpc "github.com/ethereum/go-ethereum/rpc"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/crypto"

	. "github.com/PredictionExplorer/augur-explorer/contracts"
)
const (
	RWALK_MINT_EVENT		= "ad2bc79f659de022c64ef55c71f16d0cf125452ed5fc5757b2edc331f58565ec"
	BID_WITH_CST	bool = true
	TIME_DELAY_SEC		= 1
	TIME_DELAY_ON_ERROR   = 500 // milliseconds
	DELAY_AFTER_TX		= 2 // seconds
	DELAY_NO_ACTION		= 5	// seconds
	TIME_UNTIL_PRIZE_LIMIT = 300 // seconds
	BID_GAS_LIMIT		int = 1000000
	CLAIM_GAS_LIMIT		int = 5000000
	MAX_RETRIES			int64 = 3		// if something doesn't work for 3 times, then we abandon the task
	CST_BID_ANYWAY		bool = true
)
const (
	FLOW_ID_UNINITIALIZED = iota
	FLOW_ID_NOT_THE_LAST_BIDDER
	FLOW_ID_I_AM_THE_LAST_BIDDER
	FLOW_ID_NEED_TO_BID_WITH_CST
	FLOW_ID_NEED_TO_WAIT_FOR_CST_BID_TX
	FLOW_ID_NEED_TO_BID_WITH_PLAIN_ETH
	FLOW_ID_NEED_TO_WAIT_FOR_ETH_BID_TX
	FLOW_ID_NEED_TO_BID_WITH_RWALK
	FLOW_ID_NEED_TO_CLAIM_PRIZE
	FLOW_ID_WAITING_FOR_CLAIM_PRIZE_TX
	FLOW_ID_NEED_TO_WAIT_FOR_RWALK_BID_TX
	FLOW_ID_NEED_TO_WAIT_FOR_RWALK_MINT
	FLOW_ID_NEED_TO_SEND_RWALK_BID_TX

)
var (
	PKEY_HEX		string = os.Getenv("PKEY_HEX")
	CGAME_ADDR		string = os.Getenv("CGAME_ADDR")
	MAX_ETH_BID_ETHER_AMOUNT *big.Int = big.NewInt(8e17)	// maximum amount of eth to spend on bidding (including 50% rwalk discount + rwalk mint price)
	MAX_BID_CST_AMOUNT		*big.Int = big.NewInt(0).Mul(big.NewInt(1e18),big.NewInt(8))	// below this bids will be made using CST
	RWALK_BID_START_PRICE	*big.Int = big.NewInt(1e17)		// only bid with RWALK if current ETH bid price is above this value

	RPC_URL string
	token_addr		common.Address
	chain_id		*big.Int
	round_num_played	int64 = -1
	flow_id				int	= 0	// event programming (event code identifier) Codes: 0: uninitialized: 1: I am not last bidder; 2: I am last bidder; waiting for claim prize; 3: I claimed prize, job done, end of work; 4: I sent bid transaction (CST / ETH) ; 5: I sent rwalk mint transaction; 

	bidParamType, _ = abi.NewType("tuple","BidParams",[]abi.ArgumentMarshaling{
		{Name: "message", Type: "string"},
		{Name: "randomWalkNFTId", Type: "int256"},
	})
	params = abi.Arguments{
		{Type: bidParamType, Name: "bp"},
	}
	eth_bid_tx_hash			common.Hash
	cst_bid_tx_hash			common.Hash
	rwalk_mint_tx_hash		common.Hash
	rwalk_bid_tx_hash		common.Hash
	claim_prize_tx_hash		common.Hash
	max_retries_counter		int64 = 0
	next_rwalk_token_id		int64 = -1
	prev_rwalk_token_id		int64 = -1

	rpcclient				*ethrpc.Client
	eclient 				*ethclient.Client
	copts 					bind.CallOpts
)
type RW_MintEvent struct {
    TokenId *big.Int
    Owner   common.Address
    Seed    [32]byte
    Price   *big.Int
    Raw     types.Log // Blockchain specific contextual infos
}
func fmt_eth(wei *big.Int) string {
	ether := new(big.Float).SetInt(wei)
	eth_value := new(big.Float).Quo(ether, big.NewFloat(1e18))
	return eth_value.Text('f', 18) // 18 decimal places to match Ethereum precision
}
func find_rwalk_token_id(cg *CosmicSignatureGame,rw *RWalk,me common.Address) {
	
	if next_rwalk_token_id > -1 {
		// we will just verify that next_rwalk_token_id hasn't been used to fix any potential bugs that might arise in event loop logic
		is_used,err := cg.UsedRandomWalkNfts(&copts,big.NewInt(next_rwalk_token_id))
		if err == nil {
			if is_used.Cmp(big.NewInt(1)) == 0 {
				fmt.Printf("Resetting next_rwalk_token_id (%v) to -1 because it was used already\n",next_rwalk_token_id)
				next_rwalk_token_id=-1
			} else  {
				fmt.Printf("returning on find_rwalk_token_id(), next_rwalk_token_id=%v\n",next_rwalk_token_id)
				return	// rwalk token id for next mint is already set
			}
		}
	}
	target_token_id := prev_rwalk_token_id + 1
	last_token_id,err := rw.NextTokenId(&copts)
	if err != nil {
		fmt.Printf("Error calling NextTokenId(): %v\n",err)
		time.Sleep(TIME_DELAY_ON_ERROR * 2 * time.Millisecond)
		return
	}
	for {
		if last_token_id.Int64() <= target_token_id {
			return	// we have used all our rwalk tokens , new mint is required
		}
		owner,err := rw.OwnerOf(&copts,big.NewInt(target_token_id))
		if err != nil {
			fmt.Printf("Error getting OwnerOf() random walk token: %v\n",err)
			time.Sleep(TIME_DELAY_ON_ERROR * 2 * time.Millisecond)
			return
		}
		prev_rwalk_token_id=target_token_id
		if bytes.Compare(owner.Bytes(),me.Bytes()) == 0 {
			is_used,err := cg.UsedRandomWalkNfts(&copts,big.NewInt(target_token_id))
			if err == nil {
				if is_used.Cmp(big.NewInt(1)) != 0 {
					next_rwalk_token_id=target_token_id
					fmt.Printf("Next token ID to use for RWalk bid will be %v\n",next_rwalk_token_id)
					return
				}
			} else {
				time.Sleep(TIME_DELAY_ON_ERROR * 2 * time.Millisecond)
			}
		}
		target_token_id = target_token_id + 1
		if target_token_id >= last_token_id.Int64() {
			break
		}
	}
}
func main() {
	var err error 
	RPC_URL = os.Getenv("RPC_URL")

	var zero_addr common.Address
	rpcclient,err = ethrpc.DialContext(context.Background(),RPC_URL)
	if err!=nil {
		fmt.Printf("Can't connect to ETH RPC: %v\n",err)
		os.Exit(1)
	}
	eclient = ethclient.NewClient(rpcclient)

	chain_id,err = eclient.ChainID(context.Background())
	if err != nil {
		fmt.Printf("Error getting chain id : %v\n",err)
		os.Exit(1)
	}

	from_pkey_str := PKEY_HEX
	if len(from_pkey_str) != 64 {
		fmt.Printf("Sender's private key is not 64 characters long (no 0x prefix please)\n")
		os.Exit(1)
	}

	cosmic_game_addr_hex := CGAME_ADDR
	if len(cosmic_game_addr_hex) != 40 {
		fmt.Printf("%v\n",cosmic_game_addr_hex)
		fmt.Printf("CosmicGame contract address doesn't have 40 characters long (no 0x prefix please)\n")
		os.Exit(1)
	}
	cosmic_game_addr := common.HexToAddress(cosmic_game_addr_hex)

	cosmic_game_ctrct,err := NewCosmicSignatureGame(cosmic_game_addr,eclient)
	if err!=nil {
		fmt.Printf("Failed to instantiate CosmicGame contract: %v\n",err)
		os.Exit(1)
	}
	prizes_wallet_addr,err := cosmic_game_ctrct.PrizesWallet(&copts)
	if err != nil {
		fmt.Printf("Failed to get PrizesWallet address: %v\n",err)
		os.Exit(1)
	}
	prizes_wallet,err := NewPrizesWallet(prizes_wallet_addr,eclient)
	if err!=nil {
		fmt.Printf("Failed to instantiate PrizesWallet contract: %v\n",err)
		os.Exit(1)
	}
	rwalk_addr,err := cosmic_game_ctrct.RandomWalkNft(&copts)
	if err != nil {
		fmt.Printf("Error getting RWalk addr: %v\n",err)
		os.Exit(1)
	}
	rwalk_ctrct,err := NewRWalk(rwalk_addr,eclient)
	if err != nil {
		fmt.Printf("Error creating RWalk() instance: %v\n",err)
		os.Exit(1)
	}
	round_num_played_big,err := cosmic_game_ctrct.RoundNum(&copts)
	if err != nil {
		fmt.Printf("Error getting round_num: %v\n",err)
		os.Exit(1)
	}
	round_num_played = round_num_played_big.Int64()
	fmt.Printf("Playing round %v\n",round_num_played)

	from_pkey, err := crypto.HexToECDSA(from_pkey_str)
	if err != nil {
		fmt.Printf("Error making private key: %v\n",err)
		os.Exit(1)
	}
	from_publicKey := from_pkey.Public()
	from_publicKeyECDSA, ok := from_publicKey.(*ecdsa.PublicKey)
	if !ok {
		fmt.Printf("Couldn't derive public key for Sender")
		os.Exit(1)
	}
	from_address := crypto.PubkeyToAddress(*from_publicKeyECDSA)
	fmt.Printf("Config params:\n\tmax bid amount: %v ETH\n\tRwalk bidding begins from %v ETH\n\tCST bidding starts below %v CST\n\tSend bid %v secs before claimPrize() event",fmt_eth(MAX_ETH_BID_ETHER_AMOUNT),fmt_eth(RWALK_BID_START_PRICE),fmt_eth(MAX_BID_CST_AMOUNT),TIME_UNTIL_PRIZE_LIMIT)
	// Design note: we use event programming model, the business logic is triggered by event codes
	//				we use single event loop and lots of 'case's in a switch{} using flow_id variable, this simplifies 
	//				error handling since transactions have asynchronous execution
	//				if the action has more following actions, the inner loop should issue 'continue' instruction, but
	//				if the action is final, the inner loop should issue 'continue maineventloop' instruction
	im_last_bidder_notified := false
	maineventloop:
	for {
		gasPrice, err := eclient.SuggestGasPrice(context.Background())
		if err != nil {
			fmt.Printf("Error getting suggested gas price: %v\n",err)
			os.Exit(1)
		}
		// get latest variable values
		fmt.Printf("Main event loop starts (rwnext=%v, rwprev=%v)\n",next_rwalk_token_id,prev_rwalk_token_id)
		if next_rwalk_token_id == -1 {
			go find_rwalk_token_id(cosmic_game_ctrct,rwalk_ctrct,from_address)
		}
		var cst_price *big.Int
		cst_price,err = cosmic_game_ctrct.GetNextCstBidPrice(&copts)
		if err != nil {
			fmt.Printf("Error getting CST bid price: %v\n",err)
			time.Sleep(TIME_DELAY_ON_ERROR * time.Millisecond)
			continue
		}
		bid_price,err := cosmic_game_ctrct.GetNextEthBidPrice(&copts)
		if err != nil {
			fmt.Printf("Error getting ETH bid price: %v\n",err)
			time.Sleep(TIME_DELAY_ON_ERROR * time.Millisecond)
			continue
		}
		time_until_prize,err := cosmic_game_ctrct.GetDurationUntilMainPrize(&copts)
		if err != nil {
			fmt.Printf("Error at getDurationUntilMainPrize(): %v\n",err)
			time.Sleep(TIME_DELAY_ON_ERROR * time.Millisecond)
			continue
		}
		token_addr,err := cosmic_game_ctrct.Token(&copts)
		if err != nil {
			fmt.Printf("Error at getting ERC20 token addr(): %v\n",err)
			time.Sleep(TIME_DELAY_ON_ERROR * time.Millisecond)
			continue
		}
		token_ctrct,err := NewERC20(token_addr,eclient)
		if err != nil {
			fmt.Printf("Error instantiating ERC20 token contract: %v\n",err)
			time.Sleep(TIME_DELAY_ON_ERROR * time.Millisecond)
			continue
		}
		cst_balance,err:=token_ctrct.BalanceOf(&copts,from_address)
		if err!=nil {
			fmt.Printf("Error during BalanceOf() call: %v\n",err)
			time.Sleep(TIME_DELAY_ON_ERROR * time.Millisecond)
			continue
		}

		last_bidder,err := cosmic_game_ctrct.LastBidderAddress(&copts)
		if err != nil {
			fmt.Printf("Error at LastBidder(): %v\n",err)
			time.Sleep(TIME_DELAY_ON_ERROR * time.Millisecond)
			continue
		}
		rnum,err := cosmic_game_ctrct.RoundNum(&copts)
		if err != nil {
			fmt.Printf("Error getting round_num: %v\n",err)
			time.Sleep(TIME_DELAY_ON_ERROR * time.Millisecond)
			continue
		}
		eth_balance,err := eclient.BalanceAt(context.Background(),from_address,nil)
		if err != nil {
			fmt.Printf("Error getting my ETH balance: %v\n",err)
			time.Sleep(TIME_DELAY_ON_ERROR * time.Millisecond)
			continue
		}
		rwalk_mint_price,err := rwalk_ctrct.GetMintPrice(&copts)
		if err != nil {
			fmt.Printf("Error getting RWalk mint price: %v\n",err)
			continue
		}
		timeout_claimprize,err := cosmic_game_ctrct.TimeoutDurationToClaimMainPrize(&copts)
		if err != nil {
			fmt.Printf("Error at TimeoutClaimPrize(): %v\n",err)
			continue
		}
		prize_time,err := cosmic_game_ctrct.MainPrizeTime(&copts)
		if err != nil {
			fmt.Printf("Error at PrizeTime(): %v\n",err)
			continue
		}
		if bytes.Compare(last_bidder.Bytes(),zero_addr.Bytes()) != 0 {
			var result map[string]interface{}
			err = rpcclient.CallContext(context.Background(), &result, "eth_getBlockByNumber", "pending", false)
			if err != nil {
				fmt.Printf("Failed to get pending block: %v", err)
			} else {
				if timestamp_hex, ok := result["timestamp"].(string); ok {
					block_timestamp, err := hexutil.DecodeBig(timestamp_hex)
					if err != nil {
						fmt.Printf("Failed to decode timestamp: %v", err)
					} else {
						time_out_expired := big.NewInt(0).Add(prize_time,timeout_claimprize);
						if time_out_expired.Cmp(block_timestamp) < 0 {
							fmt.Printf("Winner of the main prize didn't claim it during claim-prize window, I am going to claim it\n")
							flow_id=FLOW_ID_NEED_TO_CLAIM_PRIZE
						}
					}
				}
			}
		}

		if rnum.Int64() != round_num_played {
			fmt.Printf("Round num changed (was playing %v, but current round is %v) , exiting\n",round_num_played,rnum)
			rwinner,err := prizes_wallet.MainPrizeBeneficiaryAddresses(&copts,big.NewInt(round_num_played))
			if err != nil  {
				fmt.Printf("Error getting winner for the round played: %v\n",err)
			}
			if bytes.Compare(rwinner.Bytes(),from_address.Bytes()) == 0 {
				fmt.Printf("I am the winner of the round %v, execution successful\n",round_num_played)
			} else {
				fmt.Printf("I am not the winner of round %v, execution failed\n",round_num_played)
			}
			os.Exit(0)
		}
		cur_flow_id := flow_id
		for {	// in this for we reuse contract variables fetched earlier, i.e. run the flow with the same vars
			fmt.Printf("flow_id=%v (cur_flow_id=%v) time_u_p=%v\n",flow_id,cur_flow_id,time_until_prize.Int64())
			switch flow_id {
				case FLOW_ID_UNINITIALIZED:
					if bytes.Equal(last_bidder.Bytes(),from_address.Bytes()) {
						// I am the last bidder
						if !im_last_bidder_notified {
							fmt.Printf("I am last bidder, continuing\n")
							im_last_bidder_notified = true
						}
						flow_id = FLOW_ID_I_AM_THE_LAST_BIDDER
						time.Sleep(DELAY_NO_ACTION * time.Second)
						continue maineventloop	// im last bidder, do nothing
					} else {
						fmt.Printf("I am not the last bidder (time until prize = %v)\n",time_until_prize.String())
						im_last_bidder_notified = false
						flow_id = FLOW_ID_NOT_THE_LAST_BIDDER
					}
				case FLOW_ID_NOT_THE_LAST_BIDDER:
					//fmt.Printf("CST_BID_ANYWAY=%v\n",CST_BID_ANYWAY)
					if CST_BID_ANYWAY {
						//fmt.Printf("cst_balance = %v, cst_price=%v, MAX_BID_CST_AMOUNT=%v\n",fmt_eth(cst_balance),fmt_eth(cst_price),fmt_eth(MAX_BID_CST_AMOUNT))
						if (MAX_BID_CST_AMOUNT.Cmp(cst_balance)<=0) && (cst_price.Cmp(MAX_BID_CST_AMOUNT) <= 0) {
							fmt.Printf("CST_BID_ANYWAY = true , so I will keep bidding if CST price is low\n")
							fmt.Printf("CST price (%v) below my limit (%v), bidding with CST\n",fmt_eth(cst_price),fmt_eth(MAX_BID_CST_AMOUNT))
							flow_id = FLOW_ID_NEED_TO_BID_WITH_CST
							continue
						}
					}
					if bytes.Equal(last_bidder.Bytes(),from_address.Bytes()) {	// verify the information
						// I am the last bidder
						if !im_last_bidder_notified {
							fmt.Printf("I am last bidder, continuing\n")
							im_last_bidder_notified = true
						}
						flow_id = FLOW_ID_I_AM_THE_LAST_BIDDER
						time.Sleep(DELAY_NO_ACTION * time.Second)
						continue maineventloop	// im last bidder, do nothing
					}
					if time_until_prize.Cmp(big.NewInt(0)) == 0 { // check if I can claim prize due to timeout
						// pending
						_=timeout_claimprize
					}
					if time_until_prize.Cmp(big.NewInt(TIME_UNTIL_PRIZE_LIMIT)) <= 0 {
						fmt.Printf("%v sec before claim prize event, my time to bid has came\n",time_until_prize.Int64())
						// time is running out, I'm not the last bidder and  I have to make a bid quickly
						if (MAX_BID_CST_AMOUNT.Cmp(cst_balance)<=0) && (cst_price.Cmp(MAX_BID_CST_AMOUNT) <= 0) {
							fmt.Printf("CST price (%v) below my limit (%v), bidding with CST\n",fmt_eth(cst_price),fmt_eth(MAX_BID_CST_AMOUNT))
							flow_id = FLOW_ID_NEED_TO_BID_WITH_CST
						} else {
							// can't bid with CST, 2 options left: 1) ETH + RWALK or 2) plain ETH
							if MAX_BID_CST_AMOUNT.Cmp(cst_balance)>0 {
								fmt.Printf("Can't make a bid with CST ,not enough CST balance for max bid of %v CST\n",fmt_eth(MAX_BID_CST_AMOUNT))
							} else {
								fmt.Printf("Can't make a bid with CST price is above %v CST\n",fmt_eth(MAX_BID_CST_AMOUNT))
							}
							if RWALK_BID_START_PRICE.Cmp(bid_price) < 0 {
								// bidding with RWALK is allowed
								rwalk_discounted_price := big.NewInt(0).Quo(bid_price,big.NewInt(2))
								bid_with_rwalk_price := big.NewInt(0).Add(rwalk_mint_price,rwalk_discounted_price)
								// is bidding with RWALK cheaper?
								fmt.Printf("Bidding with RWALK + ETH costs %v , with pure ETH costs %v\n",fmt_eth(bid_with_rwalk_price),fmt_eth(bid_price))
								if bid_price.Cmp(bid_with_rwalk_price) <= 0 {
									// plain ETH is cheaper
									fmt.Printf("Deciding to bid with plain ETH (it is cheaper)\n")
									if bid_price.Cmp(MAX_ETH_BID_ETHER_AMOUNT)<0 {	// is ETH price below my limit ?
										if bid_price.Cmp(eth_balance) < 0 {
											flow_id = FLOW_ID_NEED_TO_BID_WITH_PLAIN_ETH		// yes, lets bid with plain ETH
										} else {
											fmt.Printf("Im out of funds for bidding (my limit = %v, bid_price=%v, my balance = %v)\n",fmt_eth(MAX_ETH_BID_ETHER_AMOUNT),fmt_eth(bid_price),fmt_eth(eth_balance))
											time.Sleep(DELAY_NO_ACTION * time.Second)
											continue maineventloop
										}
									} else {
										// no ETH bid price too high
										fmt.Printf("ETH bid price rose above my limit (cur price = %v, limit = %v), skipping\n",fmt_eth(bid_price),fmt_eth(MAX_ETH_BID_ETHER_AMOUNT))
										time.Sleep(DELAY_NO_ACTION * time.Second)
										continue maineventloop
									}
								} else {
									if bid_with_rwalk_price.Cmp(MAX_ETH_BID_ETHER_AMOUNT)<0 {	// is RWALK mint + 50% of bid_price is below my limit ?
										fmt.Printf("Bid with RWalk (%v is cheaper than plain eth (%v)\n",fmt_eth(bid_with_rwalk_price),fmt_eth(bid_price))
										flow_id = FLOW_ID_NEED_TO_BID_WITH_RWALK		// ETH + RWALK is cheaper
									} else {
										fmt.Printf("Im out of funds for bidding even with RWalk (my balance = %v, bid_cost=%v)\n",fmt_eth(eth_balance),fmt_eth(bid_with_rwalk_price))
										time.Sleep(DELAY_NO_ACTION * time.Second)
										continue maineventloop
									}
								}
							} else { 
								// ETH bid only (rwalk bid not allowed)
								if MAX_ETH_BID_ETHER_AMOUNT.Cmp(bid_price) < 0 {	// 
									fmt.Printf("Bid price is too high for my config (%v), skipping\n",fmt_eth(bid_price))
									time.Sleep(DELAY_NO_ACTION * time.Second)
									continue maineventloop
								} else {
									if bid_price.Cmp(eth_balance) < 0 {
										flow_id = FLOW_ID_NEED_TO_BID_WITH_PLAIN_ETH
									} else {
										fmt.Printf("Im out of funds for bidding (my balance = %v, bid_price=%v)",fmt_eth(eth_balance),fmt_eth(bid_price))
										time.Sleep(DELAY_NO_ACTION * time.Second)
										continue maineventloop
									}
								}
							}
							continue
						}
					} else {
						fmt.Printf("Not my time to bid yet (time_until_prize = %v)\n",time_until_prize.Int64())
						time.Sleep(DELAY_NO_ACTION * time.Second)
						continue maineventloop
					}
				case FLOW_ID_I_AM_THE_LAST_BIDDER:
					last_bidder,err := cosmic_game_ctrct.LastBidderAddress(&copts)	// just verify once again
					if err != nil {
						fmt.Printf("Error at LastBidder(): %v\n",err)
						time.Sleep(TIME_DELAY_ON_ERROR * time.Millisecond)
						continue maineventloop
					}
					if bytes.Equal(last_bidder.Bytes(),from_address.Bytes()) {  //verify it
						if CST_BID_ANYWAY {
							//fmt.Printf("cst_balance = %v, cst_price=%v, MAX_BID_CST_AMOUNT=%v\n",fmt_eth(cst_balance),fmt_eth(cst_price),fmt_eth(MAX_BID_CST_AMOUNT))
							if (MAX_BID_CST_AMOUNT.Cmp(cst_balance)<=0) && (cst_price.Cmp(MAX_BID_CST_AMOUNT) <= 0) {
								fmt.Printf("CST_BID_ANYWAY = true , so I will keep bidding if CST price is low\n")
								fmt.Printf("CST price (%v) below my limit (%v), bidding with CST\n",fmt_eth(cst_price),fmt_eth(MAX_BID_CST_AMOUNT))
								flow_id = FLOW_ID_NEED_TO_BID_WITH_CST
								continue
							}
						}
						// I am the last bidder
						if !im_last_bidder_notified {
							fmt.Printf("I am last bidder, continuing\n")
							im_last_bidder_notified = true
						}
						if time_until_prize.Cmp(big.NewInt(0)) == 0 {
							if bytes.Equal(last_bidder.Bytes(),from_address.Bytes()) {  //verify it
								flow_id = FLOW_ID_NEED_TO_CLAIM_PRIZE
								continue
							} else {
								continue maineventloop
							}
						} else {
							if (time_until_prize.Cmp(big.NewInt(DELAY_NO_ACTION)) < 0) {
								continue maineventloop// when we get closer to claim prize time, do not sleep
							} else {
								time.Sleep(DELAY_NO_ACTION * time.Second)
								continue maineventloop
							}
						}
						flow_id = FLOW_ID_I_AM_THE_LAST_BIDDER
						time.Sleep(TIME_DELAY_SEC * time.Second)
						continue maineventloop	// im last bidder, do nothing
					} else {
						fmt.Printf("I am not the last bidder (time until prize = %v)\n",time_until_prize.String())
						im_last_bidder_notified = false
						flow_id = FLOW_ID_NOT_THE_LAST_BIDDER
					}
				case FLOW_ID_NEED_TO_BID_WITH_CST:
					from_nonce, err := eclient.PendingNonceAt(context.Background(), from_address)
					if err != nil {
						fmt.Printf("Error getting account's nonce: %v\n",err)
						time.Sleep(TIME_DELAY_ON_ERROR * time.Millisecond)
						continue maineventloop
					}
					txopts := bind.NewKeyedTransactor(from_pkey)
					txopts.GasLimit = uint64(BID_GAS_LIMIT)
					txopts.GasPrice = big.NewInt(0).Mul(gasPrice,big.NewInt(2))
					txopts.Nonce = big.NewInt(int64(from_nonce))
					signfunc := func(address common.Address, tx *types.Transaction) (*types.Transaction, error) {
						signer := types.NewEIP155Signer(chain_id)
						signature, err := crypto.Sign(signer.Hash(tx).Bytes(), from_pkey)
						if err != nil {
							fmt.Printf("Error signing: %v\n",err)
							return nil,nil
						}
						return tx.WithSignature(signer, signature)
					}
					if signfunc == nil {
						time.Sleep(TIME_DELAY_ON_ERROR * time.Millisecond)
						continue maineventloop
					}
					txopts.Signer = signfunc
					tx,err := cosmic_game_ctrct.BidWithCst(txopts,cst_price,"")
					if err!=nil {
						fmt.Printf("BidWithCST(): error sending tx: %v\n",err)
						time.Sleep(TIME_DELAY_ON_ERROR * time.Millisecond)
						continue maineventloop
					}
					if tx != nil {
						fmt.Printf("CST bid tx hash: %v\n",tx.Hash().String())
						flow_id = FLOW_ID_NEED_TO_WAIT_FOR_CST_BID_TX
						max_retries_counter = 0
						cst_bid_tx_hash = tx.Hash()
						time.Sleep(DELAY_AFTER_TX * time.Second)
						continue maineventloop
					}
				case FLOW_ID_NEED_TO_WAIT_FOR_CST_BID_TX:
					cst_receipt,err := eclient.TransactionReceipt(context.Background(),cst_bid_tx_hash)
					if err != nil {
						fmt.Printf("Error getting receipt for cst bid tx: %v\n",err)
						if max_retries_counter >= MAX_RETRIES {
							flow_id=FLOW_ID_UNINITIALIZED
							time.Sleep(TIME_DELAY_ON_ERROR * time.Millisecond)
							continue maineventloop
						}
						time.Sleep(TIME_DELAY_ON_ERROR * time.Millisecond)
						max_retries_counter++
						continue maineventloop
					}
					if cst_receipt.Status != types.ReceiptStatusSuccessful {
						fmt.Printf("Error at receipt status on CST bid tx(), resetting status\n")
						flow_id=FLOW_ID_UNINITIALIZED
						time.Sleep(TIME_DELAY_ON_ERROR * time.Millisecond)
						continue maineventloop
					}
					last_bidder,err := cosmic_game_ctrct.LastBidderAddress(&copts)
					if err != nil {
						fmt.Printf("Error at LastBidder(): %v\n",err)
						time.Sleep(TIME_DELAY_ON_ERROR * time.Millisecond)
						continue maineventloop
					}
					if bytes.Equal(last_bidder.Bytes(),from_address.Bytes()) {
						// I am the last bidder
						if !im_last_bidder_notified {
							fmt.Printf("I am last bidder after bid with CST, continuing\n")
							im_last_bidder_notified = true
						}
						flow_id = FLOW_ID_I_AM_THE_LAST_BIDDER
						time.Sleep(DELAY_NO_ACTION * time.Second)
						continue maineventloop	// im last bidder, do nothing
					} else {
						fmt.Printf("I am not the last bidder after cst bid (time until prize = %v)\n",time_until_prize.String())
						im_last_bidder_notified = false
						flow_id = FLOW_ID_NOT_THE_LAST_BIDDER
					}
				case FLOW_ID_NEED_TO_BID_WITH_PLAIN_ETH:
					from_nonce, err := eclient.PendingNonceAt(context.Background(), from_address)
					if err != nil {
						fmt.Printf("Error getting account's nonce: %v\n",err)
						time.Sleep(TIME_DELAY_ON_ERROR * time.Millisecond)
						continue maineventloop
					}
					txopts := bind.NewKeyedTransactor(from_pkey)
					txopts.GasLimit = uint64(BID_GAS_LIMIT)
					txopts.GasPrice = big.NewInt(0).Mul(gasPrice,big.NewInt(2))
					txopts.Nonce = big.NewInt(int64(from_nonce))
					txopts.Value = big.NewInt(0).Set(bid_price)
					signfunc := func(address common.Address, tx *types.Transaction) (*types.Transaction, error) {
						signer := types.NewEIP155Signer(chain_id)
						signature, err := crypto.Sign(signer.Hash(tx).Bytes(), from_pkey)
						if err != nil {
							fmt.Printf("Error signing: %v\n",err)
							return nil,nil
						}
						return tx.WithSignature(signer, signature)
					} 
					if signfunc == nil {
						time.Sleep(TIME_DELAY_ON_ERROR * time.Millisecond)
						continue maineventloop
					}
					txopts.Signer = signfunc
					tx,err := cosmic_game_ctrct.BidWithEth(txopts,big.NewInt(-1),"")
					if err!=nil {
						fmt.Printf("Bid() (with plain ETH) error during sending tx: %v\n",err)
						time.Sleep(TIME_DELAY_ON_ERROR * time.Millisecond)
						continue maineventloop
					}
					if tx != nil {
						fmt.Printf("ETH bid (price = %v ETH)  tx hash: %v\n",fmt_eth(bid_price),tx.Hash().String())
						flow_id = FLOW_ID_NEED_TO_WAIT_FOR_ETH_BID_TX
						max_retries_counter = 0
						eth_bid_tx_hash = tx.Hash()
						time.Sleep(DELAY_AFTER_TX * time.Second)
						continue maineventloop
					}
				case FLOW_ID_NEED_TO_WAIT_FOR_ETH_BID_TX:
					eth_receipt,err := eclient.TransactionReceipt(context.Background(),eth_bid_tx_hash)
					if err != nil {
						fmt.Printf("Error getting receipt for plain eth bid tx: %v \n",err)
						if max_retries_counter >= MAX_RETRIES {
							flow_id=FLOW_ID_UNINITIALIZED
							time.Sleep(TIME_DELAY_ON_ERROR * time.Millisecond)
							continue maineventloop
						}
						time.Sleep(TIME_DELAY_ON_ERROR * time.Millisecond)
						max_retries_counter++
						continue maineventloop
					}
					if eth_receipt.Status != types.ReceiptStatusSuccessful {
						fmt.Printf("Error at receipt status on ETH bid tx(), resetting status\n")
						flow_id=FLOW_ID_UNINITIALIZED
						time.Sleep(TIME_DELAY_ON_ERROR * time.Millisecond)
						continue maineventloop
					}
					last_bidder,err := cosmic_game_ctrct.LastBidderAddress(&copts)
					if err != nil {
						fmt.Printf("Error at LastBidder(): %v\n",err)
						time.Sleep(TIME_DELAY_ON_ERROR * time.Millisecond)
						continue maineventloop
					}
					if bytes.Equal(last_bidder.Bytes(),from_address.Bytes()) {
						// I am the last bidder
						if !im_last_bidder_notified {
							fmt.Printf("I am last bidder after plain ETH bid, continuing\n")
							im_last_bidder_notified = true
						}
						flow_id = FLOW_ID_I_AM_THE_LAST_BIDDER
						time.Sleep(DELAY_NO_ACTION * time.Second)
						continue maineventloop	// im last bidder, do nothing
					} else {
						fmt.Printf("I am not the last bidder after eth bid (time until prize = %v)\n",time_until_prize.String())
						im_last_bidder_notified = false
						flow_id = FLOW_ID_NOT_THE_LAST_BIDDER
						continue maineventloop
					}
				case FLOW_ID_NEED_TO_BID_WITH_RWALK:
					if (next_rwalk_token_id > -1) {
						fmt.Printf("We have pre-minted RWalk tokens , going to use token_id = %v\n",next_rwalk_token_id)
						flow_id = FLOW_ID_NEED_TO_SEND_RWALK_BID_TX
						continue
					} else {
						fmt.Printf("We don't own any RWalk, need to mint one\n")
					}
					rwalk_from_nonce, err := eclient.PendingNonceAt(context.Background(), from_address)
					if err != nil {
						fmt.Printf("Error getting account's nonce: %v\n",err)
						time.Sleep(TIME_DELAY_ON_ERROR * time.Millisecond)
						continue maineventloop
					}
					rwopts := bind.NewKeyedTransactor(from_pkey)
					rwopts.Nonce = big.NewInt(int64(rwalk_from_nonce))
					rwopts.Value = big.NewInt(0).Set(rwalk_mint_price)
					rwopts.GasPrice = big.NewInt(0).Mul(gasPrice,big.NewInt(2))	// aribtrum gas glitches fix
					rwopts.GasLimit = uint64(BID_GAS_LIMIT)
					// mint rwalk
					signfunc := func(address common.Address, tx *types.Transaction) (*types.Transaction, error) {
						signer := types.NewEIP155Signer(chain_id)
						signature, err := crypto.Sign(signer.Hash(tx).Bytes(), from_pkey)
						if err != nil {
							fmt.Printf("Error signing: %v\n",err)
							return nil,nil
						}
						return tx.WithSignature(signer, signature)
					}
					if signfunc == nil {
						time.Sleep(TIME_DELAY_ON_ERROR * time.Millisecond)
						continue maineventloop
					}
					rwopts.Signer = signfunc
					rwtx,err := rwalk_ctrct.Mint(rwopts)
					if err!=nil {
						fmt.Printf("RWalk mint, error sending tx: %v\n",err)
						time.Sleep(TIME_DELAY_ON_ERROR * time.Millisecond)
						continue maineventloop
					}
					if rwtx != nil {
						fmt.Printf("RWalk mint tx hash: %v\n",rwtx.Hash().String())
						flow_id = FLOW_ID_NEED_TO_WAIT_FOR_RWALK_MINT
						rwalk_mint_tx_hash = rwtx.Hash()
						max_retries_counter = 0
						time.Sleep(DELAY_AFTER_TX * time.Second)
						continue maineventloop
					}
				case FLOW_ID_NEED_TO_WAIT_FOR_RWALK_MINT:
					rw_receipt,err := eclient.TransactionReceipt(context.Background(),rwalk_mint_tx_hash)
					if err != nil {
						fmt.Printf("Error getting receipt for rwalk mint tx: %v\n",err)
						if max_retries_counter >= MAX_RETRIES {
							flow_id=FLOW_ID_UNINITIALIZED
							time.Sleep(TIME_DELAY_ON_ERROR * time.Millisecond)
							continue maineventloop
						}
						time.Sleep(TIME_DELAY_ON_ERROR * time.Millisecond)
						max_retries_counter++
						continue
					}
					evt_mint_event,_ := hex.DecodeString(RWALK_MINT_EVENT)
					var log_found bool = false
					var rwalk_token_id int64 = -1
					for _,elog := range rw_receipt.Logs {
						if len(elog.Topics)>0 {
							if bytes.Compare(elog.Topics[0].Bytes(),evt_mint_event) == 0 {
								log_found = true
								rwalk_token_id = elog.Topics[1].Big().Int64()
								next_rwalk_token_id = rwalk_token_id
								flow_id = FLOW_ID_NEED_TO_SEND_RWALK_BID_TX
								continue
							}
						}
					}
					if !log_found {
						fmt.Printf("RandomWalk mint event not found in receipt logs\n")
						time.Sleep(TIME_DELAY_ON_ERROR * time.Millisecond)
						continue maineventloop
					}
				case FLOW_ID_NEED_TO_SEND_RWALK_BID_TX:
					from_nonce, err := eclient.PendingNonceAt(context.Background(), from_address)
					if err != nil {
						fmt.Printf("Error getting account's nonce: %v\n",err)
						time.Sleep(TIME_DELAY_ON_ERROR * time.Millisecond)
						continue maineventloop
					}

					txopts := bind.NewKeyedTransactor(from_pkey)
					txopts.Nonce = big.NewInt(int64(from_nonce))
					txopts.GasPrice = big.NewInt(0).Mul(gasPrice,big.NewInt(2))
					txopts.GasLimit = uint64(BID_GAS_LIMIT)
					txopts.Value = big.NewInt(0).Set(bid_price)

					signfunc := func(address common.Address, tx *types.Transaction) (*types.Transaction, error) {
						signer := types.NewEIP155Signer(chain_id)
						signature, err := crypto.Sign(signer.Hash(tx).Bytes(), from_pkey)
						if err != nil {
							fmt.Printf("Error signing tx for bid with rwalk: %v\n",err)
							return nil,nil
						}
						return tx.WithSignature(signer, signature)
					} 
					if signfunc == nil {
						time.Sleep(TIME_DELAY_ON_ERROR * time.Millisecond)
						continue maineventloop
					}
					txopts.Signer = signfunc
					tx,err := cosmic_game_ctrct.BidWithEth(txopts,big.NewInt(next_rwalk_token_id),"")
					if err!=nil {
						fmt.Printf("Bid() (with RWalk), error sending tx: %v\n",err)
						time.Sleep(TIME_DELAY_ON_ERROR * time.Millisecond)
						continue maineventloop
					}
					if tx != nil {
						fmt.Printf("Bid tx hash: %v, rwalk token id used = %v\n",tx.Hash().String(),next_rwalk_token_id)
						flow_id = FLOW_ID_NEED_TO_WAIT_FOR_RWALK_BID_TX
						max_retries_counter = 0
						rwalk_bid_tx_hash = tx.Hash()
						time.Sleep(DELAY_AFTER_TX * time.Second)
					}

				case FLOW_ID_NEED_TO_WAIT_FOR_RWALK_BID_TX:
					rw_receipt,err := eclient.TransactionReceipt(context.Background(),rwalk_bid_tx_hash)
					if err != nil {
						fmt.Printf("Error getting receipt for rwalk bid tx: %v\n",err)
						if max_retries_counter >= MAX_RETRIES {
							flow_id=FLOW_ID_UNINITIALIZED
							time.Sleep(TIME_DELAY_ON_ERROR * time.Millisecond)
							continue maineventloop
						}
						time.Sleep(TIME_DELAY_ON_ERROR * time.Millisecond)
						max_retries_counter++
						continue
					}
					if rw_receipt.Status != types.ReceiptStatusSuccessful {
						fmt.Printf("Error at receipt status on RWALK bid tx(), resetting status\n")
						flow_id=FLOW_ID_UNINITIALIZED
						time.Sleep(TIME_DELAY_ON_ERROR * time.Millisecond)
						continue maineventloop
					}
					next_rwalk_token_id = -1
					// rwalk bit tx was processed, check last bidder status
					last_bidder,err := cosmic_game_ctrct.LastBidderAddress(&copts)
					if err != nil {
						fmt.Printf("Error at LastBidder(): %v\n",err)
						time.Sleep(TIME_DELAY_ON_ERROR * time.Millisecond)
						continue maineventloop
					}
					if bytes.Equal(last_bidder.Bytes(),from_address.Bytes()) {
						// I am the last bidder
						if !im_last_bidder_notified {
							fmt.Printf("I am last bidder after rwalk bid, continuing\n")
							im_last_bidder_notified = true
						}
						flow_id = FLOW_ID_I_AM_THE_LAST_BIDDER
						time.Sleep(DELAY_NO_ACTION * time.Second)
						continue maineventloop	// im last bidder, do nothing
					} else {
						fmt.Printf("I am not the last bidder after rwalk bid (time until prize = %v)\n",time_until_prize.String())
						im_last_bidder_notified = false
						flow_id = FLOW_ID_NOT_THE_LAST_BIDDER
						continue maineventloop
					}
				case FLOW_ID_NEED_TO_CLAIM_PRIZE:
					from_nonce, err := eclient.PendingNonceAt(context.Background(), from_address)
					if err != nil {
						fmt.Printf("Error getting account's nonce at claimPrize(): %v\n",err)
						time.Sleep(TIME_DELAY_ON_ERROR * time.Millisecond)
						continue maineventloop
					}
					txopts := bind.NewKeyedTransactor(from_pkey)
					txopts.GasLimit = uint64(CLAIM_GAS_LIMIT)
					txopts.GasPrice = big.NewInt(0).Mul(gasPrice,big.NewInt(2))
					txopts.Nonce = big.NewInt(int64(from_nonce))
					signfunc := func(address common.Address, tx *types.Transaction) (*types.Transaction, error) {
						signer := types.NewEIP155Signer(chain_id)
						signature, err := crypto.Sign(signer.Hash(tx).Bytes(), from_pkey)
						if err != nil {
							fmt.Printf("Error signing tx for claim prize: %v\n",err)
							return nil,nil
						}
						return tx.WithSignature(signer, signature)
					} 
					if signfunc == nil {
						time.Sleep(TIME_DELAY_ON_ERROR * time.Millisecond)
						continue maineventloop
					}
					txopts.Signer = signfunc
					tx,err := cosmic_game_ctrct.ClaimMainPrize(txopts)
					if err!=nil {
						fmt.Printf("ClaimPrize, error sending tx: %v\n",err)
						time.Sleep(TIME_DELAY_ON_ERROR * time.Millisecond)
						continue maineventloop
					}
					if tx != nil {
						fmt.Printf("ClaimPrize tx hash: %v\n",tx.Hash().String())
						claim_prize_tx_hash = tx.Hash()
						flow_id=FLOW_ID_WAITING_FOR_CLAIM_PRIZE_TX
						max_retries_counter = 0
						time.Sleep(DELAY_AFTER_TX * time.Second)
					}

				case FLOW_ID_WAITING_FOR_CLAIM_PRIZE_TX:
					receipt,err := eclient.TransactionReceipt(context.Background(),claim_prize_tx_hash)
					if err != nil {
						fmt.Printf("Error getting receipt for rwalk bid tx: %v\n",err)
						if max_retries_counter >= MAX_RETRIES {
							flow_id=FLOW_ID_UNINITIALIZED
							time.Sleep(TIME_DELAY_ON_ERROR * time.Millisecond)
							continue maineventloop
						}
						max_retries_counter++
						time.Sleep(TIME_DELAY_ON_ERROR * time.Millisecond)
						continue
					}
					if receipt.Status == types.ReceiptStatusSuccessful {
						fmt.Printf("ClaimPrize tx processed, I should be the winner\n")
						flow_id=FLOW_ID_UNINITIALIZED
						time.Sleep(TIME_DELAY_ON_ERROR * time.Millisecond)
						continue maineventloop	// main event loop should exit on next iteration because round num has changed
					} else {
						// something went wrong, lets restart things
						flow_id=FLOW_ID_UNINITIALIZED
						time.Sleep(TIME_DELAY_ON_ERROR * time.Millisecond)
						continue maineventloop
					}
				default:
					fmt.Printf("Unknown flow id code, bug?\n")
					time.Sleep(DELAY_NO_ACTION * time.Second)
			}
			if flow_id == cur_flow_id {
				break;	// exit loop if flow_id doesn't change
			}
			cur_flow_id = flow_id
		}// for 2 (flow processing)
		gasPrice, err = eclient.SuggestGasPrice(context.Background())
		if err != nil {
			fmt.Printf("Error getting suggested gas price: %v\n",err)
			time.Sleep(TIME_DELAY_ON_ERROR * time.Millisecond)
			continue maineventloop
		}
		time.Sleep(TIME_DELAY_SEC * time.Second)
	} // for 1 (main event loop)
}
