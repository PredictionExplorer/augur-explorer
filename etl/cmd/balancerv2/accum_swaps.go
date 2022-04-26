// Accumulates swap fees per different timeframe
// Timeframe codes:
//		0 - Hourly
//		1 - Daily
//		2 - Weekly
//		3 - Monthly

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
var (
	timeframe_code			int64 = -1
)
func get_timestamp_interval_secs(tf_code int64) int64 {

	switch tf_code {
	case 0: return 60*60
	case 1: return 60*60*24
	case 2: return 60*60*24*7
	case 3: return 60*60*24*30
}
func main() {

	usage_str := fmt.Sprintf("usage: %v --schema [schema_name]\n",os.Args[0])
	if len(os.Args)<2 {
		fmt.Printf("%v",usage_str)
		os.Exit(1)
	}
	schema_name := flag.String("schema", "", "Schema name")
	tf_code_ptr := flag.Int64("tf_code",0,"Timeframe code (0-hourly(default),1-daily,2-weekly,3-monthly)")
	flag.Parse()
	if len(*schema_name) < 3 {
		fmt.Printf("Schema name must be larger than 2 characters\n")
		fmt.Printf("%v",usage_str)
		os.Exit(1)
	}
	timefame_code = *tf_code_ptr
	var err error
	RPC_URL = os.Getenv("RPC_URL")
	eclient, err = ethclient.Dial(RPC_URL)
	if err!=nil {
		fmt.Printf("Can't connect to ETH RPC: %v\n",err)
		os.Exit(1)
	}

	Info = log.New(os.Stdout,"INFO: ",log.Ldate|log.Ltime|log.Lshortfile)
	storagew.S = Connect_to_storage_with_schema(Info,*schema_name)
	Info.Printf("Schema name: %v\n",schema_name)

	pool_aid := storagew.Get_lowest_pool_aid()
	if pool_aid == 0 {
		Info.Printf("No pools found in the database, aborting\n")
		os.Exit(1)
	}
	var first_ts,cur_ts int64
	for {
		last_rec,err := storagew.Get_last_swap_accum_record(pool_aid,timeframe_code)
		if err != nil {
			first_ts = storagew.Get_timestamp_of_first_swap_fee_hist_record(pool_aid)
			if first_ts == 0 { // no fees yet
				swap_fee_total,last_id,err := storagew.Get_swaps_for_period(pool_aid,ini_ts,fin_ts)
				pool_aid = storagew.Get_greater_pool_aid(pool_aid)
				continue	// process next pool
			}
			if pool_aid == 0 {
				Info.Printf("All pools were processed, exiting.\n")
				os.Exit(0)
			}
			ts = firtst_ts
		} else {
			ts = last_rec.TimeStamp
		}
		interv := get_timeframe_interval_secs(timeframe_code)
		ts = ts / interv
		ts = ts * interv	// make it divisible without reminder

		// We always update the last record first, and then try the subsequent period.
		// This is needed because we aren't keeping track of chain reorg in the table,
		//	so , data can change. Since each record of 'swap_accum' is linked by ID to
		//	the latest record of swf_hist upon DELETE of the block it will disappear since
		//	actions are CASCADEd. So, the update of latest record will update the data always
		//	if the dat has changed
		fin_ts := ts + interv
	}	
}
