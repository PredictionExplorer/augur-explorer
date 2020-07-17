package main
/// Updates number of unique addresses per day
import (
	"os"
	"log"
	"time"
	"strconv"
	"fmt"

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

	if len(os.Args) < 2 {
		fmt.Printf("usage: %v [days_back]\n",os.Args[0])
		os.Exit(1)
	}
	days_back,err := strconv.ParseInt(os.Args[1],10,64)
	if err!=nil {
		fmt.Printf("Bad number for days_back parameter: %v\n",err)
		os.Exit(1)
	}

	Info = log.New(os.Stdout,"INFO: ",log.Ldate|log.Ltime|log.Lshortfile)
	storage = Connect_to_storage(&market_order_id,Info)
	last_block_ts:=storage.Get_last_block_timestamp()
	modulus := last_block_ts % ONE_DAY_SECS
	day_ts := last_block_ts - modulus
	next_day_ts := day_ts + ONE_DAY_SECS
	for days_back > 0 {
		num_addrs,no_rows := storage.Calc_unique_addresses(day_ts,next_day_ts)
		if !no_rows {
			storage.Update_unique_addresses_entry(day_ts,num_addrs)
		}
		tm := time.Unix(day_ts, 0)
		Info.Printf("Day %v processed: from %v to %v\n",tm,day_ts,next_day_ts)
		day_ts = day_ts - ONE_DAY_SECS
		next_day_ts = next_day_ts - ONE_DAY_SECS
		days_back--
	}
}
