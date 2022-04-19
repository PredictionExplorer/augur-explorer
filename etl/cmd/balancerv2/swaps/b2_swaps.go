package main

import (
	"log"
	"fmt"
	"os"
	"flag"
	"strings"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/common"

	. "github.com/PredictionExplorer/augur-explorer/primitives/balancerv2"
	. "github.com/PredictionExplorer/augur-explorer/dbs"
	. "github.com/PredictionExplorer/augur-explorer/dbs/balancerv2"
	. "github.com/PredictionExplorer/augur-explorer/contracts"
)
// Notes:
//		first block with swaps on main net: 
//			swap				12293069
//			balanc change		12286258
var (
	Info				*log.Logger
	storagew			SQLStorageWrapper
	pool_map			map[string]int64
	RPC_URL string

	eclient				 *ethclient.Client
	getswapfee_abi		*abi.ABI
)
func process_block_balance_changes(block_num int64,block_hash string,bchanges []BalV2PoolBalanceChanged) {

	if len(bchanges) == 0 {
		return
	}
	for i:=0; i<len(bchanges); i++ {
		rec_in := bchanges[i]
		Info.Printf(
			"BalanceChange: block %v , tx_index %v , logidx %v timestamp %v\n",
			rec_in.BlockNum,rec_in.TxIndex,rec_in.LogIndex,rec_in.TimeStamp,
		)
		pool_aid,exists := pool_map[rec_in.PoolId]
		if !exists {
			pool_aid,_ = storagew.Lookup_pool_address_id(rec_in.PoolId)
			if pool_aid == 0 {
				Info.Printf("Error looking up for pool id %v : not found\n",rec_in.PoolId)
				os.Exit(1)
			}
			pool_map[rec_in.PoolId]=pool_aid
		}
		Info.Printf("\tPool %v  pool_aid = %v\n",rec_in.PoolId,pool_aid)
		Info.Printf("\tblock num = %v,contract_aid=%v\n",rec_in.BlockNum,pool_aid)
		var rec_out BalV2BalChg
		rec_out.BlockNum = rec_in.BlockNum
		rec_out.BlockHash = block_hash
		rec_out.TimeStamp = rec_in.TimeStamp
		rec_out.TxIndex = rec_in.TxIndex
		rec_out.LogIndex = rec_in.LogIndex
		rec_out.PoolAid = pool_aid
		rec_out.PoolId = rec_in.PoolId
		toks := strings.Split(rec_in.Tokens,",")
		amounts := strings.Split(rec_in.Deltas,",")
		for i:=0; i<len(toks); i++ {
			tok_addr := toks[i]
			tok_aid,err := storagew.S.Layer1_lookup_address_id(tok_addr)
			if err != nil {
				Info.Printf("Error looking up token address: %v\n",err)
				os.Exit(1)
			}
			rec_out.TokenAid = tok_aid
			amount := amounts[i]
			rec_out.Amount = amount
			storagew.Insert_balance_change_history_record(&rec_out)
		}
	}
}
func process_block_swaps(block_num int64,block_hash string,swaps []BalV2Swap) {

	if len(swaps) == 0 {
		return
	}
	for i:=0; i<len(swaps); i++ {
		s := swaps[i]
		Info.Printf(
			"Swap: block %v , tx_index %v , logidx %v timestamp %v\n",
			s.BlockNum,s.TxIndex,s.LogIndex,s.TimeStamp,
		)
		pool_aid,exists := pool_map[s.PoolId]
		if !exists {
			pool_aid,_ = storagew.Lookup_pool_address_id(s.PoolId)
			if pool_aid == 0 {
				Info.Printf("Error looking up for pool id %v : not found\n",s.PoolId)
				os.Exit(1)
			}
			pool_map[s.PoolId]=pool_aid
		}
		Info.Printf("\tPool %v  pool_aid = %v\n",s.PoolId,pool_aid)
		Info.Printf("\tIn: %v \t Out: %v\n",s.AmountIn,s.AmountOut)
		Info.Printf("\tblock num = %v,contract_aid=%v\n",s.BlockNum,pool_aid)
		fee_percentage_str,_,found := storagew.Get_pool_fee_by_block_num(
			pool_aid,
			s.BlockNum,
			s.TxIndex,
		)
		if !found {
			Info.Printf(
				"Fee not found for swap at block %v, tx_id=%v, log_idx=%v",
				s.BlockNum,s.TxIndex,s.LogIndex,
			)
			Info.Printf(
				"Trying to fetch the fee from the chain, by direct call to contract of pool %v\n",
				s.PoolId,
			)
			pool_aid,err := storagew.Lookup_pool_address_id(s.PoolId)
			if err != nil {
				Info.Printf("Can't find pool address id, exiting\n")
				os.Exit(1)
			}
			pool_addr_str,err := storagew.S.Layer1_lookup_address(pool_aid)
			if err != nil {
				Info.Printf("Can't lookup address string: %v\n",err)
				os.Exit(1)
			}
			var copts bind.CallOpts
			pool_addr := common.HexToAddress(pool_addr_str)
			Info.Printf("Calling GetSwapFeePercentage at %v\n",pool_addr.String())
			pool_ctrct,err := NewGetSwapFee(pool_addr,eclient)
			if err != nil {
				Info.Printf("Can't instantiate pool contract:%v\n",err)
				os.Exit(1)
			}
			result,err := pool_ctrct.GetSwapFeePercentage(&copts)
			if err != nil {
				Info.Printf("Call to GetSwapFeePercentage() failed: %v\n",err)
				os.Exit(1)
			}
			fee_percentage_str = result.String()
			Info.Printf("The fee for this pool (addr=%v) is %v, continuing...\n",pool_addr.String(),fee_percentage_str)
		}
		one := big.NewInt(1e18)
		amount_in := big.NewInt(0)
		amount_in.SetString(s.AmountIn,10)
		fee_percentage := big.NewInt(0)
		fee_percentage.SetString(fee_percentage_str,10)
		fee := big.NewInt(0)
		product := big.NewInt(0)
		product.Mul(fee_percentage,amount_in)
		uno := big.NewInt(1)
		product.Sub(product,uno)
		product.Quo(product,one)
		fee.Add(product,uno)
		in_plus_fee := big.NewInt(0)
		in_plus_fee.Set(amount_in)
		in_plus_fee.Add(in_plus_fee,fee)
		var rec BalV2SwapHist
		rec.BlockNum = s.BlockNum
		rec.BlockHash = block_hash
		rec.TimeStamp = s.TimeStamp
		rec.TxIndex = s.TxIndex
		rec.LogIndex = s.LogIndex
		rec.PoolAid = pool_aid
		rec.PoolId = s.PoolId
		rec.SwapFee = fee.String()
		rec.ProtocolFee = "0"
		rec.AccumSwapFee = "0"
		rec.AccumProtoFee = "0"
		id := storagew.Insert_swap_fee_history(&rec)
		Info.Printf("After insert swap fee id=%v\n",id)

		// Update token balance table
		var rec_bal BalV2BalChg
		rec_bal.SwapHistId = id
		rec_bal.BlockNum = rec.BlockNum
		rec_bal.BlockHash = block_hash
		rec_bal.TimeStamp = rec.TimeStamp
		rec_bal.TxIndex = rec.TxIndex
		rec_bal.LogIndex = rec.LogIndex
		rec_bal.PoolAid = pool_aid
		rec_bal.PoolId = rec.PoolId
		rec_bal.TokenAid = s.TokenInAid
		rec_bal.Amount = in_plus_fee.String()
		storagew.Insert_balance_change_history_record(&rec_bal) // incoming token

		rec_bal.TokenAid = s.TokenOutAid
		rec_bal.Amount = s.AmountOut
		storagew.Insert_balance_change_history_record(&rec_bal) // outgoing token
	}
}
func main() {

	usage_str := fmt.Sprintf("usage: %v --schema [schema_name]\n",os.Args[0])
	if len(os.Args)<2 {
		fmt.Printf("%v",usage_str)
		os.Exit(1)
	}
	schema_name := flag.String("schema", "", "Schema name")
	starting_block := flag.Int64("startblock",0,"Single block number to process")
	flag.Parse()
	if len(*schema_name) < 3 {
		fmt.Printf("Schema name must be larger than 2 characters\n")
		fmt.Printf("%v",usage_str)
		os.Exit(1)
	}
	var err error
	RPC_URL = os.Getenv("RPC_URL")
	eclient, err = ethclient.Dial(RPC_URL)
	if err!=nil {
		fmt.Printf("Can't connect to ETH RPC: %v\n",err)
		os.Exit(1)
	}

	Info = log.New(os.Stdout,"INFO: ",log.Ldate|log.Ltime|log.Lshortfile)
	storagew.S = Connect_to_storage_with_schema(Info,*schema_name)
	//storagew.S.Db_set_schema_name(*schema_name)
	Info.Printf("Schema name: %v\n",schema_name)
	//storagew.S.Set_search_path_to_schema_name()

	abi_parsed1 := strings.NewReader(GetSwapFeeABI)
	abi1,err := abi.JSON(abi_parsed1)
	if err!= nil {
		Info.Printf("Can't parse PoolFactory ABI: %v\n",err)
		os.Exit(1)
	}
	getswapfee_abi = &abi1

	var block_num int64
	var block_hash string
	var found bool
	if *starting_block != 0 {
		Info.Printf("Processing from block %v\n",*starting_block)
		block_hash_found,success := storagew.S.Layer1_get_hash_by_block_num(*starting_block)
		if !success {
			Info.Printf("Starting block %v wasn't found\n",*starting_block)
			os.Exit(1)
		}
		found = true
		block_hash = block_hash_found
		block_num = *starting_block
	} else {
		block_num,block_hash,found = storagew.Get_last_block_for_swap_history()
	}
	Info.Printf("First call: block_num=%v, hash = %v\n",block_num,block_hash)
	if !found {
		block_num,block_hash,found = storagew.Get_first_block_for_swap_history()
		Info.Printf("Second call: block_num=%v, hash = %v\n",block_num,block_hash)
	}

	pool_map = make(map[string]int64)
	for {
		storagew.Delete_balance_changes(block_num)
		storagew.Delete_swap_history(block_num)
		swaps := storagew.Get_swaps_for_block(block_num,block_hash)
		Info.Printf("swaps: block_num=%v hash %v , len(swaps) = %v\n",block_num,block_hash,len(swaps))
		if len(swaps) == 0 {
			bnum,err := storagew.S.Layer1_get_block_num_by_hash(block_hash)
			if err != nil {
				Info.Printf(
					"No more blocks in the DB, exiting at block_num=%v (hash=%v)\n",
					block_num,block_hash,
				)
				os.Exit(0)
			}
			if bnum != block_num {
				Info.Printf(
					"Chain split detected at block %v (block hash=%v), exiting\n",
					block_num,block_hash,
				)
				os.Exit(1)
			}
		}
		process_block_swaps(block_num,block_hash,swaps)

		bchanges := storagew.Get_balance_changes_for_block(block_num,block_hash)
		Info.Printf("bchanges: block_num=%v hash %v , len(bchanges) = %v\n",block_num,block_hash,len(bchanges))
		if len(bchanges) == 0 {
			bnum,err := storagew.S.Layer1_get_block_num_by_hash(block_hash)
			if err != nil {
				Info.Printf(
					"No more blocks in the DB, exiting at block_num=%v (hash=%v)\n",
					block_num,block_hash,
				)
				os.Exit(0)
			}
			if bnum != block_num {
				Info.Printf(
					"Chain split detected at block %v (block hash=%v), exiting\n",
					block_num,block_hash,
				)
				os.Exit(1)
			}
		}
		process_block_balance_changes(block_num,block_hash,bchanges)

		saved_block_num := block_num
		block_num,block_hash,found = storagew.S.Layer1_get_next_block_by_hash(block_hash)
		if !found {
			Info.Printf("No more blocks found, exiting at block %v\n",saved_block_num)
			os.Exit(0)
		}
	}
}
