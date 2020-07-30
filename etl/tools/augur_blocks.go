package main
/// Exttracts only blocks that belong to Augur trading, that is blocks in the tables:
//			'mktord','market','mktfin','claim_funds','report'
//			(without including DAI/REP token transfer blocks)
// if market address isn't provided as parameter exports blocks for all markets, otherwise only for that market
// This tool is meant for testing purposes only to work with the Main Net on specific portions of data.
// After extracting block numbers you can feed the data into the database only for the list of extracted blocks
import (
	"os"
	"log"
	"fmt"

	"github.com/ethereum/go-ethereum/common"
	. "github.com/PredictionExplorer/augur-explorer/dbs"
)
var (
	storage *SQLStorage

	fill_order_id int64 = 0			// during event processing, holds id of record in mktord from Fill evt
	market_order_id int64 = 0

	Info    *log.Logger
)
func main() {	// returns 0 - no changes, 2 - day was added

	var specific_market bool = false
	var market_addr common.Address
	if len(os.Args) == 2 {
		specific_market = true
		market_addr = common.HexToAddress(os.Args[1])
	}
	logfile, err := os.OpenFile("/tmp/augur_blocks.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err!=nil {
		fmt.Printf("Can't start: %v\n",err)
		os.Exit(1)
	}
	Info = log.New(logfile,"INFO: ",log.Ldate|log.Ltime|log.Lshortfile)
	storage = Connect_to_storage(&market_order_id,Info)

	var market_aid int64 = 0
	if specific_market {
		market_aid,err = storage.Nonfatal_lookup_address_id(market_addr.String())
	}
	blocks := storage.Get_augur_blocks(market_aid)
	for i := 0; i < len(blocks) ; i++ {
		if i>0 {
			fmt.Printf(",")
		}
		fmt.Printf("%v",blocks[i])
	}
}
