package main

import (
	"log"
	"time"
	//"fmt"
	"math/big"

	//"github.com/ethereum/go-ethereum/ethclient"
	//"github.com/ethereum/go-ethereum/rpc"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"

	. "github.com/PredictionExplorer/augur-explorer/dbs"
	. "github.com/PredictionExplorer/augur-explorer/contracts"
)
/*
var (
	RPC_URL = os.Getenv("AUGUR_ETH_NODE_RPC_URL")

	storage *SQLStorage
	caddrs *ContractAddresses

	eclient *ethclient.Client
	rpcclient *rpc.Client

	market_order_id int64 = 0

	Error   *log.Logger
	Info	*log.Logger
)
)
*/
func update_erc20_balances_backwards(lgr *log.Logger,contract_addr_str string,last_block_num int64,contract_aid,acct_aid int64,addr *common.Address) (int,error) {

	contract_addr := common.HexToAddress(contract_addr_str)
	var copts = new(bind.CallOpts)
	copts.BlockNumber = big.NewInt(int64(last_block_num))
	erc20_token,err := NewOwnedERC20(contract_addr,eclient)
	if err != nil {
		return 0,err
	}
	balance,err := erc20_token.BalanceOf(copts,*addr)
	if err != nil {
		lgr.Printf(
			"XX %v XX %v XX: Failure to update ERC20  token balances backwards. " +
			"last_block_num=%v contract %v",
			contract_aid,acct_aid,last_block_num,contract_addr_str,
		)
		lgr.Printf(
			"XX %v XX %v XX: Failure to update ERC20 token balances backwards. " +
			"last_block_num=%v contract %v",
			contract_aid,acct_aid,last_block_num,contract_addr_str,
		)
		return 0,err
	}
	lgr.Printf(
		"XX %v XX %v XX: Updating balances backwards from block %v. " +
		" addr %v contract %v\n",
		contract_aid,acct_aid,last_block_num,addr.String(),contract_addr_str,
	)
	lgr.Printf(
		"XX %v XX %v XX: Geth node RPC: got last balance = %v for block = %v contract %v\n",
		contract_aid,acct_aid,balance.String(),last_block_num,contract_addr_str,
	)
	return storage.Update_erc20_token_balances_backwards(last_block_num,contract_aid,acct_aid,balance),nil
}
func erc20_bal_sleep() {
	time.Sleep(14 * time.Second)	// Ethereum block time
}
func balance_updater(lgr *log.Logger) {
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
			operations := storage.Get_unprocessed_erc20_balances(last_id)
			if len(operations) > 0 {
				last_id = operations[len(operations)-1].Id
			}
			lgr.Printf("loop begins, got %v operations\n",len(operations))
			for i := 0 ; i<len(operations) ; i++ {
				erc20_bal := &operations[i]
				lgr.Printf("XX %v XX %v XX: balance object: %+v\n",erc20_bal.ContractAid,erc20_bal.Aid,erc20_bal)
				prev_balance_db,err := storage.Get_previous_erc20_balance_from_DB(
					erc20_bal.Id,
					erc20_bal.ContractAid,
					erc20_bal.Aid,
				)
				lgr.Printf(
					"XX %v XX %v XX: prev balance from DB = %v, err=%v\n",
					erc20_bal.ContractAid,erc20_bal.Aid,erc20_bal.Address,prev_balance_db,err,
				)
				if err != nil {
					if err == ErrUnprocessedBalances {
						last_block_num,success := storage.Get_last_block_num()
						if success {
							addr := common.HexToAddress(erc20_bal.Address)
							affected_rows,_:=update_erc20_balances_backwards(
								lgr,erc20_bal.ContractAddr,
								last_block_num,
								erc20_bal.ContractAid,
								erc20_bal.Aid,
								&addr,
							)
							if affected_rows>0 {
								num_changes++
								lgr.Printf(
									"XX %v XX %v XX: restarting loop() do to data invalidation, " +
									"count of rows updated=%v\n",
									erc20_bal.ContractAid,erc20_bal.Aid,num_changes)
								break		// update backards invalidates the 'operations' array
							}
						}
						continue
					}
					// no balance locally (in the DB), get it from RPC
					var copts = new(bind.CallOpts)
					copts.BlockNumber = big.NewInt(int64(erc20_bal.BlockNum)-1)	// previous block is used
					contract_addr := common.HexToAddress(erc20_bal.ContractAddr)
					addr := common.HexToAddress(erc20_bal.Address)
					erc20_token,err := NewOwnedERC20(contract_addr,eclient)
					if err != nil {
						lgr.Printf(
							"XX %v XX %v XX: Can't instantiate ERC20 token contract: %v\n",
							erc20_bal.ContractAid,erc20_bal.Aid,err,
						)
						break
					}
					prev_bal,err := erc20_token.BalanceOf(copts,addr)
					if err != nil {
						lgr.Printf(
							"XX %v XX %v XX: Error on GetBalance call (addr=%v addr=%v): %v\n",
							erc20_bal.ContractAid,erc20_bal.Aid,addr.String(),erc20_bal.Address,err,
						)
						// if error occurs, it probably means the Node has already deleted the State for this block
						// therefore the only way to update balances of this account is calculate changes backwards,
						last_block_num,success := storage.Get_last_block_num()
						if success {
							affected_rows,err:=update_erc20_balances_backwards(
								BalancesLog,
								erc20_bal.ContractAddr,
								last_block_num,
								erc20_bal.ContractAid,
								erc20_bal.Aid,
								&addr,
							)
							if err!=nil {
								//dai_bal_sleep() // RPC service error, go to sleep
								//break;
							}
							if affected_rows>0 {
								num_changes++
								lgr.Printf(
									"XX %v XX %v XX: restarting loop() due to data invalidation, "+
									"affected rows=%v on addr %v\n",
									erc20_bal.ContractAid,erc20_bal.Aid,num_changes,addr.String(),
								)
								break		// update backards invalidates the 'operations' array
							}
						}
					} else {
						amount := new(big.Int)
						amount.SetString(erc20_bal.Amount,10)
						new_bal := new(big.Int)
						new_bal.Add(prev_bal,amount)
						lgr.Printf(
							"XX %v XX %v XX: setting balance of acct %v to %v " +
							"(prev_bal=%v, amount=%v)\n",
							erc20_bal.ContractAid,erc20_bal.Aid,
							addr.String(),new_bal,prev_bal.String(),amount.String(),
						)
						storage.Set_erc20_balance(erc20_bal.Id,erc20_bal.BlockHash,new_bal.String())
						num_changes++
					}
				} else {
					prev_bal := new(big.Int)
					prev_bal.SetString(prev_balance_db,10)
					amount := new(big.Int)
					amount.SetString(erc20_bal.Amount,10)
					new_bal := new(big.Int)
					new_bal.Add(prev_bal,amount)
					lgr.Printf(
						"XX %v XX %v XX: setting balance from db of acct %v to %v "+
						"(prev_bal=%v, amount=%v)\n",
						erc20_bal.ContractAid,erc20_bal.Aid,
						erc20_bal.Address,new_bal,prev_bal.String(),amount.String(),
					)
					storage.Set_erc20_balance(erc20_bal.Id,erc20_bal.BlockHash,new_bal.String())
					num_changes++
				}
			}
			if len(operations) == 0 {
				break;	// we have processed all erc20 records
			}
			if num_changes > 0 {
				break;	// any change in erc20_bal invalidates the query
			}
		}
		if num_changes == 0 {
			erc20_bal_sleep()
		}
	}
}
