package main

import (
	"log"
	"fmt"
	"os"
	"flag"
	"math/big"

	. "github.com/PredictionExplorer/augur-explorer/primitives/balancerv2"
	. "github.com/PredictionExplorer/augur-explorer/dbs"
	. "github.com/PredictionExplorer/augur-explorer/dbs/balancerv2"
)
var (
	Info				*log.Logger
	storagew			SQLStorageWrapper
)
func process_block_swaps(swaps []BalV2Swap) {

	if len(swaps) == 0 {
		return
	}
	for i:=0; i<len(swaps); i++ {
		s := swaps[i]
		fee_percentage_str,fee_ts,found := storagew.Get_pool_fee_by_timestamp(
			s.ContractAid,
			s.TimeStamp,
			s.BlockNum,
			s.TxIndex,
		)
		if !found {
			Info.Printf(
				"Fee not found for swap at block %v, tx_id=%v, log_idx=%v",
				s.BlockNum,s.TxIndex,s.LogIndex,
			)
			os.Exit(0)
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
		var rec BalV2SwapHist
		rec.BlockNum = s.BlockNum
		rec.TimeStamp = s.TimeStamp
		rec.TxIndex = s.TxIndex
		rec.LogIndex = s.LogIndex
		rec.ContractAid = s.ContractAid
		rec.PoolId = s.PoolId
		rec.SwapFee = fee.String()
		storagew.Insert_swap_fee_history(&rec)
	}
}
func main() {

	usage_str := fmt.Sprintf("usage: %v --schema [schema_name]\n",os.Args[0])
	if len(os.Args)<2 {
		fmt.Printf("%v",usage_str)
		os.Exit(1)
	}
	schema_name := flag.String("schema", "", "Schema name")
	flag.Parse()
	if len(*schema_name) < 3 {
		fmt.Printf("Schema name must be larger than 2 characters\n")
		fmt.Printf("%v",usage_str)
		os.Exit(1)
	}
	Info = log.New(os.Stdout,"INFO: ",log.Ldate|log.Ltime|log.Lshortfile)
	storagew.S = Connect_to_storage(Info)
	storagew.S.Db_set_schema_name(*schema_name)
	Info.Printf("Schema name: %v\n",schema_name)

	block_num,block_hash,found := storagew.Get_last_block_for_swap_history()
	if !found {
		block_num,block_hash,found = storagew.Get_first_block_for_swap_history()
	}

	for {
		swaps := storagew.Get_swaps_for_block(block_num,block_hash)
		if len(swaps) == 0 {
			bnum,err := storagew.S.Get_block_num_by_hash(block_hash)
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
		process_block_swaps(swaps)
	}
}
