package main

import (
	"os"
	"fmt"
	"context"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"

	 . "github.com/PredictionExplorer/augur-explorer/uevm"
)
var (
	RPC_URL					string
	chain_id				int64 = 1234
)
func main() {

	if len(os.Args) < 2 {
		fmt.Printf(
			"Usage: \n\t\t%v [tx_hash]\n\t\t"+
			"replicates contract from main net on local db (does a deployment)\")\n\n",os.Args[0],
		)
		os.Exit(1)
	}
	tx_hash := common.HexToHash(os.Args[1])
	db := OpenDB("/tmp/evmdb")
	fmt.Printf("db = %+v\n",db)

	RPC_URL = os.Getenv("RPC_URL")
	eclient, err := ethclient.Dial(RPC_URL)
	if err!=nil {
		fmt.Printf("Error: %v\n",err)
		os.Exit(1)
	}
	r,err := eclient.TransactionReceipt(context.Background(),tx_hash)
	if err != nil {
		fmt.Printf("Error: %v\n",err)
		os.Exit(1)
	}
	contract_address := r.ContractAddress
	tx,_,err := eclient.TransactionByHash(context.Background(),tx_hash)
	if err != nil {
		fmt.Printf("Error: %v\n",err)
		os.Exit(1)
	}
	tx_msg,err := tx.AsMessage(types.LatestSignerForChainID(big.NewInt(chain_id)),tx.GasPrice())
	if err != nil {
		fmt.Printf("Error: %v\n",err)
		os.Exit(1)
	}
	code, err := eclient.CodeAt(context.Background(),contract_address,nil)
	if err != nil {
		fmt.Printf("Error: %v\n",err)
		os.Exit(1)
	}
	fmt.Printf("code size=%v\n",len(code))
	fmt.Printf("input size=%v\n",len(tx.Data()))
	err,generated_addr,state_root := UEVMDeploy2(chain_id,tx_msg.From(),tx.Nonce(),tx.Data(),db,common.Hash{})
	//err,state_root := UEVMAcctCreate(chain_id,tx_msg.From(),tx.Nonce(),db,common.Hash{})
	//generated_addr := common.Address{}
	fmt.Printf("Deploy result: %v\n",err)
	fmt.Printf("Contract address: %v\n",generated_addr.String())
	fmt.Printf("Intermediate state hash: %v\n",state_root.String())
}
