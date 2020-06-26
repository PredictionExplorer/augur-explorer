package main
/// Updates number of unique addresses per day
import (
	"os"
	"log"
	"time"
	//"sort"
	//"fmt"

	. "augur-extractor/dbs"
	//. "augur-extractor/primitives"
)
const ONE_DAY_SECS int64 = 24 * 60 * 60	// add 1 day of seconds
var (
	storage *SQLStorage

	fill_order_id int64 = 0			// during event processing, holds id of record in mktord from Fill evt
	market_order_id int64 = 0

	Info    *log.Logger
)
func main() {	// returns 0 - no changes, 2 - day was added

	Info = log.New(os.Stdout,"INFO: ",log.Ldate|log.Ltime|log.Lshortfile)
	storage = Connect_to_storage(&market_order_id,Info)
	last_block_ts:=storage.Get_last_block_timestamp()
	last_block_ts = last_block_ts - 1 * (60*60)	// discard possibly non-finalized blocks
	day_ts := storage.Get_last_unique_addr_day()
	if day_ts == 0 {
		day_ts = storage.Get_first_block_timestamp()
		modulus := day_ts % ONE_DAY_SECS
		day_ts = day_ts - modulus
	} else {
		day_ts = day_ts + ONE_DAY_SECS
		if day_ts > last_block_ts {
			os.Exit(2)	// the day is not yet ready (didn't accumulate all the blocks)
		}
	}
	next_day_ts := day_ts + ONE_DAY_SECS
	num_addrs := storage.Calc_unique_addresses(day_ts,next_day_ts)
	storage.Insert_unique_addresses_entry(day_ts,num_addrs)
	tm := time.Unix(day_ts, 0)
	Info.Printf("Day %v processed: from %v to %v\n",tm,day_ts,next_day_ts)
	os.Exit(0)
}
