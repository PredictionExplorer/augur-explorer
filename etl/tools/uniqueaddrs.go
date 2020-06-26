package main
/// Updates number of unique addresses per day
import (
	"os"
	"log"
	"sort"
	//"fmt"

	. "augur-extractor/dbs"
	. "augur-extractor/primitives"
)
var (
	storage *SQLStorage

	fill_order_id int64 = 0			// during event processing, holds id of record in mktord from Fill evt
	market_order_id int64 = 0

	Info    *log.Logger
)
func main() int {	// returns 0 - no changes, 2 - day was added

	Info = log.New(os.Stdout,"INFO: ",log.Ldate|log.Ltime|log.Lshortfile)
	storage = Connect_to_storage(&market_order_id,Info)

	ts:=storage.Get_last_block_timestamp()
	ts = ts - 1 * (60*60)		// substract time for 1 hour for any possible network split (we want finalized blocks only)
	day_ts := storage.Get_last_unique_addr_day()
	new_day_ts := day_ts + 24 * 60 * 60	// add 1 day of seconds
	Info.Printf("from %v to %v\n",day_ts,new_day_ts)
}
