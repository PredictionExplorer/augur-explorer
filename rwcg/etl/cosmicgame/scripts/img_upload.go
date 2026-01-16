// Helper to update CosmicSignature NFT images (remotely)
package main

import (
	"fmt"
	"log"
	"os"
	"strconv"

	. "github.com/PredictionExplorer/augur-explorer/rwcg/dbs"
	. "github.com/PredictionExplorer/augur-explorer/rwcg/dbs/cosmicgame"
)

const (
	CMD_STR = "Usage: %v [total_tokens | token_seed {token_id}]\n"
)

var (
	Info     *log.Logger
	storagew SQLStorageWrapper
)

func usageAndExit() {
	fmt.Printf(CMD_STR, os.Args[0])
	os.Exit(1)
}

func main() {
	Info = log.New(os.Stdout, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile)

	storage := Connect_to_storage(Info)
	if storage == nil {
		fmt.Println("failed to connect to storage")
		os.Exit(1)
	}
	storage.Db_set_schema_name("public")
	storagew.S = storage

	if len(os.Args) < 2 {
		usageAndExit()
	}

	switch os.Args[1] {
	case "total_tokens":
		tokenTotal := storagew.Get_erc721_token_total()
		fmt.Printf("%v", tokenTotal)
	case "token_seed":
		if len(os.Args) < 3 {
			usageAndExit()
		}
		tokenID, err := strconv.ParseInt(os.Args[2], 10, 64)
		if err != nil {
			fmt.Printf("Error in integer conversion: %v\n", err)
			os.Exit(1)
		}
		tokenSeed := storagew.Get_erc721_token_seed(tokenID)
		fmt.Printf("%v", tokenSeed)
	default:
		usageAndExit()
	}
}

