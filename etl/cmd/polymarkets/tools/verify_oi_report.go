package main

import (
	"os"
	"log"
	"sort"
	"fmt"
	"strconv"

	. "github.com/PredictionExplorer/augur-explorer/dbs"
	. "github.com/PredictionExplorer/augur-explorer/primitives"
)
var (
	storage *SQLStorage

	Info    *log.Logger

	condition_id			string
	market_id				int64
	contract_aid			int64
)
func main() {

	if len(os.Args) < 2 {
		fmt.Printf("Usage: %v [market_id]\n")
		os.Exit(1)
	}
	var err error
	market_id,err = strconv.ParseInt(os.Args[1],10,64)
	if err != nil {
		fmt.Printf("Error parsing market id: %v\n",err)
		os.Exit(1)
	}

	Info = log.New(os.Stdout,"INFO: ",log.Ldate|log.Ltime|log.Lshortfile)
	storage = Connect_to_storage(Info)

	condition_id = storage.Get_condition_id(market_id)
	contract_aid = storage.Get_fpmm_contract_aid(market_id)

	tokens := storage.Get_tokens_by_condition_id(condition_id)

	

}
