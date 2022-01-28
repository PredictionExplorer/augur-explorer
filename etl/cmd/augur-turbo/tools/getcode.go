// Gets contracts code and saves to file
package main

import (
	"os"
	"fmt"
	"context"
	"io/ioutil"
	"encoding/hex"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"

	. "github.com/PredictionExplorer/augur-explorer/primitives"
)
var (
	RPC_URL string
)
func main() {

	if len(os.Args) !=3 {
		fmt.Printf("Usage: %v [contract_addr] [output_file]\n",os.Args[0])
		os.Exit(1)
	}

	RPC_URL = os.Getenv("AUGUR_ETH_NODE_RPC_URL")
	eclient, err := ethclient.Dial(RPC_URL)
	if err!=nil {
		Fatalf("Can't connect to ETH RPC, please set AUGUR_ETH_NODE_RPC_URL env variable : %v\n",err)
	}

	contract_addr := common.HexToAddress(os.Args[1])
	file_name := os.Args[2]

	data,err := eclient.CodeAt(context.Background(),contract_addr,nil)
	if err == nil {
		hex_data := hex.EncodeToString(data[:])
		data_out := []byte(hex_data)
		ioutil.WriteFile(file_name,data_out,0644)
		if err != nil {
			fmt.Printf("Error writing file: %v\n",err)
		} else {
			fmt.Printf("Data saved in %v\n",file_name)
		}
	} else {
		Fatalf("Error at RPC: %v\n",err)
	}
}
