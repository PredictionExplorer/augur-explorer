package main

import (
	"log"
	"fmt"
	"os"
	"flag"
	"math/big"

	//. "github.com/PredictionExplorer/augur-explorer/primitives"
	. "github.com/PredictionExplorer/augur-explorer/dbs"
)
var (
	Info				*log.Logger
	storage 			*SQLStorage
)
func process_block_swaps(swaps []BalV2Swap,last_fee_ts int64) {

	if len(swaps) == 0 {
		return
	}
	for i:=0; i<len(swaps); i++ {
		s := swaps[i]
		fee_str := storage.Get_pool_fee_in_timeframe(last_fee_ts,swap.TimeStamp)
		amount_in := big.NewInt(0)
		amount_in.SetString(s.AmountIn,10)
		fee := big.NewInt(0)
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
	storage = Connect_to_storage(Info)
	storage.Db_set_schema_name(*schema_name)
	Info.Printf("Schema name: %v\n",schema_name)

	block_num,block_hash,found := storage.Get_last_block_for_swap_history()
	if !found {
		block_num,block_hash,found = storage.Get_first_block_for_swap_history()
	}

	for {
		swaps := storage.Get_swaps_for_block(block_num,block_hash)
		if len(swaps) == 0 {
			bnum,err := storage.Get_block_num_by_hash(block_hash)
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
