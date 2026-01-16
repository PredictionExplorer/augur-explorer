/// Recalculate price estimate for all markets
package main
import (
	"os"
	"log"
//	"strconv"
//	"fmt"
//	"encoding/hex"
//	"io/ioutil"
	"strings"
//	"math/big"


	. "github.com/PredictionExplorer/augur-explorer/dbs"
	//. "github.com/PredictionExplorer/augur-explorer/primitives"
)
var (
	storage *SQLStorage

	fill_order_id int64 = 0			// during event processing, holds id of record in mktord from Fill evt
	market_order_id int64 = 0

	Info    *log.Logger

)

func main() {	// returns 0 - no changes, 2 - day was added

	Info = log.New(os.Stdout,"INFO: ",log.Ldate|log.Ltime|log.Lshortfile)
	storage = Connect_to_storage(&market_order_id,Info)

	market_list := storage.Get_price_estimate_update_market_list()
	for _,entry := range market_list {
		Info.Printf("Calculating price estimate for market_aid=%v\n",entry.MktAid)
		outcomes_list := strings.Split(entry.Outcomes,",")
		for outcome_idx := 0; outcome_idx < len(outcomes_list); outcome_idx++ {
			storage.Update_future_price_estimates(entry.MktAid,outcome_idx,entry.TimeStamp)
		}
	}
}
