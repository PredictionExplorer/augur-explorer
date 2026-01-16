package main

import (
	"os"
	"fmt"
	"errors"
	//"context"
	//"math/big"

	//"github.com/ethereum/go-ethereum/common"
	//"github.com/ethereum/go-ethereum/core/types"
	//"github.com/ethereum/go-ethereum/ethclient"

	 . "github.com/PredictionExplorer/augur-explorer/uevm"
)
var (
	RPC_URL					string
	chain_id				int64 = 1234
)
func main() {

	if len(os.Args) < 3 {
		fmt.Printf(
			"Usage: \n\t\t%v [filename] [receipts_dir]\n\t\t"+
			"initializes minichain file\"\n\n",os.Args[0],
		)
		os.Exit(1)
	}
	minichain_filename := os.Args[1]
	receipts_dir := os.Args[2]

	if _, err := os.Stat(minichain_filename); errors.Is(err, os.ErrNotExist) {
			// file does not exist
	} else {
		fmt.Printf("File %v already exists, refusing to overwrite existing file\n",minichain_filename)
		os.Exit(1)
	}
	mchain,err := OpenMiniChain(minichain_filename,receipts_dir)
	if err != nil {
		fmt.Printf("Error opening minichain: %v\n",err)
		os.Exit(1)
	}
	var rec Record
	err = mchain.AppendLine(&rec)
	if err != nil {
		fmt.Printf("Error: %v\n",err)
		os.Exit(1)
	}
	fmt.Printf("Done\n")
}
