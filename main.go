// Augur ETL: Converts Augur Data to SQL database
package main

import (
	"os"
	"fmt"
	"context"
	"log"
	"math/big"
	"encoding/hex"

	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/accounts/abi"
)
const (
	// ToDo: get these signatures from the abi files
	MARKET_CREATED = "ea17ae24b0d40ea7962a6d832db46d1f81eaec1562946d0830d1c21d4c000ec1"
	MARKET_OI_CHANGED = "213a05b9ad8567c2f8fa868e7375e5bf30e69add0dbb5913ca8a3e58c815c268"
	MARKET_ORDER = "9bab1368a1ed530afaad9c630ba75e6a5c1efa9f6af0139d6cda2b6af6aa801e"
	MARKET_FINALIZED = "6d39632c2dc10305bf5771cfff4af1851f07c03ea27b821cad382466bdf7a21f"
	INITIAL_REPORT_SUBMITTED = "c3ebb227c22e7644e9bef8822009f746a72c86f239760124d67fdc2c302b3115"
	MARKET_VOLUME_CHANGED = "e9f0af820300e73bae76c8e76943abe7fbb4224b49cb133e2dadc6f71acf6370"
	DISPUTE_CROWDSOURCER_CONTRIBUTION = "e7f47639cdf56ec6c5451df334b73c9ca5cccd20da2c0f4e390e9bb71a6f672a"
	TOKENS_TRANSFERRED = "3c67396e9c55d2fc8ad68875fc5beca1d96ad2a2f23b210ccc1d986551ab6fdf"
	TOKEN_BALANCE_CHANGED = "63fd58f559b73fc4da5511c341ec8a7b31c5c48538ef83c6077712b6edf5f7cb"
	SHARE_TOKEN_BALANCE_CHANGED = "350ea32dc29530b9557420816d743c436f8397086f98c96292138edd69e01cb3"
)
var (
	// these evt_ variables are here for speed to avoid calculation of Keccak256
	//		on each bytes.Compare() operation
	evt_market_created,_ = hex.DecodeString(MARKET_CREATED)
	evt_market_oi_changed,_ = hex.DecodeString(MARKET_OI_CHANGED)
	evt_market_order,_ = hex.DecodeString(MARKET_ORDER)
	evt_market_finalized,_ = hex.DecodeString(MARKET_FINALIZED)
	evt_initial_report_submitted,_ = hex.DecodeString(INITIAL_REPORT_SUBMITTED)
	evt_market_volume_changed,_ = hex.DecodeString(MARKET_VOLUME_CHANGED)
	evt_dispute_crowd_contrib,_ = hex.DecodeString(DISPUTE_CROWDSOURCER_CONTRIBUTION)
	evt_tokens_transferred,_ = hex.DecodeString(TOKENS_TRANSFERRED)
	evt_token_balance_changed,_ = hex.DecodeString(TOKEN_BALANCE_CHANGED)
	evt_share_token_balance_changed,_ = hex.DecodeString(SHARE_TOKEN_BALANCE_CHANGED)

	storage *SQLStorage
	augur_abi *abi.ABI
	trading_abi *abi.ABI
	RPC_URL = os.Getenv("AUGUR_ETH_NODE_RPC_URL")
)
func main() {
	//client, err := ethclient.Dial("http://:::8545")

	if len(RPC_URL) == 0 {
		Fatalf("Configuration error: RPC URL of Ethereum node is not set."+
				" Please set AUGUR_ETH_NODE_RPC environment variable")
	}

	//client, err := ethclient.Dial("http://192.168.1.102:18545")
	client, err := ethclient.Dial(RPC_URL)
	if err != nil {
		log.Fatal(err)
	}

	storage = connect_to_storage()

	augur_init()
	ctx := context.Background()
	latestBlock, err := client.BlockByNumber(ctx, nil)
	if err != nil {
		log.Fatal("oops:", err)
	}
	log.Println("latest block:", latestBlock.Number())

	bnum,exists := storage.get_last_block_num()
	if !exists {
		bnum = 0
	} else {
		bnum = bnum + 1
	}
	fmt.Println("Starting data load from block %v",bnum)
	var bnum_high BlockNumber = BlockNumber(latestBlock.Number().Uint64())
	for ; bnum<bnum_high; bnum++ {
		big_bnum:=big.NewInt(int64(bnum))
		block, _ := client.BlockByNumber(ctx,big_bnum)
		if block != nil {
			num_transactions, err := client.TransactionCount(ctx,block.Hash())
			if err != nil {
				fmt.Printf("block error: %v \n",err)
			} else {
				if num_transactions > 0 {
					fmt.Printf("block: %v %v %v transactions\n",block.Hash().String(),block.Number(),num_transactions)
					for tnum:=0 ; tnum < int(num_transactions) ; tnum++ {
						tx , err := client.TransactionInBlock(ctx,block.Hash(),uint(tnum))
						if err != nil {
							fmt.Printf("Error: %v",err)
						} else {
							fmt.Printf("\ttx: %v\n",tx.Hash().String())
							rcpt,err := client.TransactionReceipt(ctx,tx.Hash())
							if err != nil {
								fmt.Printf("Error: %v",err)
							} else {
								sequencer := new(EventSequencer)
								num_logs := len(rcpt.Logs)
								for i:=0 ; i<num_logs ; i++ {
									fmt.Printf(
										"\t\t\tlog %v\n\t\t\t\t\t\t for contract %v (%v of %v items)\n",
										rcpt.Logs[i].Topics[0].String(),rcpt.Logs[i].Address.String(),(i+1),len(rcpt.Logs))
									sequencer.append_event(rcpt.Logs[i])
								}
								ordered_list := sequencer.get_ordered_event_list()
								num_logs = len(ordered_list)
								for i:=0 ; i < num_logs ; i++ {
									fmt.Printf(
										"\t\t\tprocessing log with sig %v\n\t\t\t\t\t\t for contract %v\n",
										ordered_list[i].Topics[0].String(),
										ordered_list[i].Address.String())
									process_event(ordered_list[i])
								}
							}
						}
					}
				}
			}
		}
		storage.set_last_block_num(bnum)
	}// for block_num
	fmt.Printf("new latest block = %v\n",bnum_high)
}
