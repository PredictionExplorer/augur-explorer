package main

import (
	"os"
	"log"
	"time"
	//"fmt"
	"math/big"
	"context"

	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/rpc"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"

	. "github.com/PredictionExplorer/augur-explorer/dbs"
	. "github.com/PredictionExplorer/augur-explorer/primitives"
)
var (
	RPC_URL = os.Getenv("AUGUR_ETH_NODE_RPC_URL")

	storage *SQLStorage
	caddrs *ContractAddresses
	ctrct_dai_token *DAICash

	eclient *ethclient.Client
	rpcclient *rpc.Client

	market_order_id int64 = 0

	Error   *log.Logger
	Info	*log.Logger
)

func update_dai_balances_backwards(last_block_num int64,aid int64,addr *common.Address) (int,error) {

	var copts = new(bind.CallOpts)
	copts.BlockNumber = big.NewInt(int64(last_block_num))
	balance,err := ctrct_dai_token.BalanceOf(copts,*addr)
	if err != nil {
		Info.Printf("Failure to update DAI token balances backwards for eoa_aid=%v,last_block_num=%v",
							aid,last_block_num)
		Error.Printf("Failure to update DAI token balances backwards for eoa_aid=%v,last_block_num=%v",
							aid,last_block_num)
		return 0,err
	}
	Info.Printf("balance_updater(): updating balances backwards from block %v , addr %v (aid=%v)\n",
			last_block_num,addr.String(),aid)
	Info.Printf("balance_updater(): got last balance = %v for block = %v\n",balance.String(),last_block_num)
	return storage.Update_dai_token_balances_backwards(last_block_num,aid,balance),nil
}
func dai_bal_sleep() {
	time.Sleep(100 * time.Second)
}
func balance_updater() {
	// go-routine that updates balances of DAI tokens
	// when the record is inserted into dai_transf table it is inserted with balance = 0 because
	// we don't have the previous balance (and we can't get it during the processing because we are
	// processing finalized  blocks and at this stage of the process the order of transfers was lost)
	// Therefore the only way to calculate valid balances for all the accounts involved is to get the
	// balance on the previous block, and run the sequence of balance changes
	// The order of insertion into dai_transf table is valid and we can use it to reproduce the history

	// in order to avoid being a bottleneck this process must run as an independent thread

	var num_changes int;
	for {
		var last_id int64 = 0
		num_changes = 0
		for {	// while we do have dai_balances available
			Info.Printf("balance_updater() running. last_id=%v\n",last_id)
			operations := storage.Get_unprocessed_dai_balances(last_id)
			if len(operations) > 0 {
				last_id = operations[len(operations)-1].Id
			}
			Info.Printf("balance_updater(): got %v operations\n",len(operations))
			for i := 0 ; i<len(operations) ; i++ {
				dai_bal := &operations[i]
				prev_balance_db,err := storage.Get_previous_balance_from_DB(dai_bal.Id,dai_bal.Aid)
				Info.Printf("balance_updater(): acct %v: prev_balance=%v, err=%v\n",dai_bal.Address,prev_balance_db,err)
				if err != nil {
					if err == ErrUnprocessedBalances {
						last_block_num,success := storage.Get_last_block_num()
						if success {
							addr := common.HexToAddress(dai_bal.Address)
							affected_rows,err:=update_dai_balances_backwards(last_block_num,dai_bal.Aid,&addr)
							if err!=nil {
							//	dai_bal_sleep() // RPC service error, go to sleep
							//	break;
							}
							if affected_rows>0 {
								num_changes++
								Info.Printf("balance_updater(): restarting loop() affected rows=%v on addr %v\n",num_changes,addr.String())
								break		// update backards invalidates the 'operations' array
							}
						}
						continue
					}
					// no balance locally (in the DB), get it from RPC
					var copts = new(bind.CallOpts)
					copts.BlockNumber = big.NewInt(int64(dai_bal.BlockNum)-1)	// previous block is used
					addr := common.HexToAddress(dai_bal.Address)
					prev_bal,err := ctrct_dai_token.BalanceOf(copts,addr)
					if err != nil {
						Error.Printf("Error on GetBalance call: %v\n",err)
						// if error occurs, it probably means the Node has already deleted the State for this block
						// therefore the only way to update balances of this account is calculate changes backwards,
						last_block_num,success := storage.Get_last_block_num()
						if success {
							affected_rows,err:=update_dai_balances_backwards(last_block_num,dai_bal.Aid,&addr)
							if err!=nil {
								//dai_bal_sleep() // RPC service error, go to sleep
								//break;
							}
							if affected_rows>0 {
								num_changes++
								Info.Printf("balance_updater(): restarting loop() affected rows=%v on addr %v\n",num_changes,addr.String())
								break		// update backards invalidates the 'operations' array
							}
						}
					} else {
						amount := new(big.Int)
						amount.SetString(dai_bal.Amount,10)
						new_bal := new(big.Int)
						new_bal.Add(prev_bal,amount)
						Info.Printf("balance_updater(): setting balance of acct %v (id=%v) to %v (prev_bal=%v, amount=%v\n",
									addr.String(),dai_bal.Aid,new_bal,prev_bal.String(),amount.String())
						storage.Set_dai_balance(dai_bal.Id,new_bal.String())
						num_changes++
					}
				} else {
					prev_bal := new(big.Int)
					prev_bal.SetString(prev_balance_db,10)
					amount := new(big.Int)
					amount.SetString(dai_bal.Amount,10)
					new_bal := new(big.Int)
					new_bal.Add(prev_bal,amount)
					Info.Printf("balance_updater(): got balance from db of acct %v (id=%v) to %v (prev_bal=%v, amount=%v\n",
									dai_bal.Address,dai_bal.Aid,new_bal,prev_bal.String(),amount.String())
					storage.Set_dai_balance(dai_bal.Id,new_bal.String())
					num_changes++
				}
			}
			if len(operations) == 0 {
				break;	// we have processed all dai_bal records
			}
			if num_changes > 0 {
				break;	// any change in dai_bal invalidates the query
			}
		}
		if num_changes == 0 {
			dai_bal_sleep()
		}
	}
}

func main() {
	//client, err := ethclient.Dial("http://:::8545")

	if len(RPC_URL) == 0 {
		Fatalf("Configuration error: RPC URL of Ethereum node is not set."+
				" Please set AUGUR_ETH_NODE_RPC_URL environment variable")
	}

	Info = log.New(os.Stdout,"INFO: ",log.Ltime)		//|log.Lshortfile)
	Error = log.New(os.Stderr,"ERROR: ",log.Ltime)		//|log.Lshortfile)

	storage = Connect_to_storage(&market_order_id,Info)
	storage.Log_msg("Log initialized\n")

	//client, err := ethclient.Dial("http://192.168.1.102:18545")
	var err error
	rpcclient, err=rpc.DialContext(context.Background(), RPC_URL)

	if err != nil {
		log.Fatal(err)
	}
	Info.Printf("Connected to ETH node: %v\n",RPC_URL)
	eclient = ethclient.NewClient(rpcclient)
//	eclient, err = ethclient.Dial(RPC_URL)

	caddrs_obj,err := storage.Get_contract_addresses()
	if err!=nil {
		Fatalf("Can't find contract addresses in 'contract_addresses' table")
	}
	caddrs=&caddrs_obj
	ctrct_dai_token,err = NewDAICash(caddrs.Dai,eclient)
	if err != nil {
		Fatalf("Couldn't initialize DAI Cash contract: %v\n",err)
	}

	balance_updater()	// updates DAI token balances very 10 seconds
}
