package main

import (
	"log"
	"fmt"
	"os"
	"flag"

	//. "github.com/PredictionExplorer/augur-explorer/primitives"
	. "github.com/PredictionExplorer/augur-explorer/dbs"
)
var (
	Info				*log.Logger
	storage 			*SQLStorage
)
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

		
	}
	
}
