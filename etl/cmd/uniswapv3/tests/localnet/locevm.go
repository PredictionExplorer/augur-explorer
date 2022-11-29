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
	db := OpenDB("/var/tmp/evmdb")
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

	mchain,err := OpenMiniChain("/var/tmp/minichain.dat","/var/tmp/receipts")
	if err != nil {
		fmt.Printf("Error opening minichain: %v\n",err)
		os.Exit(1)
	}
	mchain.SetStateDB(&db)
	num_recs := mchain.NumRecords()
	var state_hash common.Hash
	fmt.Printf("Records in minichain database: %v recs\n",num_recs)
	if num_recs != 0 {
		last_line_rec,err := mchain.ReadLastLine()
		if err != nil {
			fmt.Printf("Error on ReadLastLine(): %v\n",err)
			os.Exit(1)
		}
		state_hash = last_line_rec.StateRoot
		fmt.Printf("Using state root %v\n",state_hash.String())
	} else {
		fmt.Printf("Creating first entry in minichain record with hash %v\n",state_hash.String())
	}
	var rec Record
	rec.BlockNum = 111
	rec.BlockHash = common.Hash{}
	rec.TxIndex = 222
	rec.TxHash = tx_hash
	err,generated_addr,state_root := mchain.ExecDeploy(chain_id,tx_hash,tx_msg.From(),tx.Nonce(),tx.Data(),contract_address,state_hash,&rec)

	fmt.Printf("Deploy result: %v\n",err)
	fmt.Printf("Contract address: %v\n",generated_addr.String())
	fmt.Printf("Intermediate state hash: %v\n",state_root.String())
}
