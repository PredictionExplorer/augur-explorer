package main

import (
	"os"
	"fmt"
	"context"
	//"math/big"

	"github.com/ethereum/go-ethereum/common"
	//"github.com/ethereum/go-ethereum/core/types"
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
			"executes new pool transaction on local db\n\n",os.Args[0],
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
	tx,_,err := eclient.TransactionByHash(context.Background(),tx_hash)
	if err != nil {
		fmt.Printf("Error: %v\n",err)
		os.Exit(1)
	}
	/*tx_msg,err := tx.AsMessage(types.LatestSignerForChainID(big.NewInt(chain_id)),tx.GasPrice())
	if err != nil {
		fmt.Printf("Error: %v\n",err)
		os.Exit(1)
	}*/

	mchain,err := OpenMiniChain("/var/tmp/minichain.dat","/var/tmp/receipts")
	if err != nil {
		fmt.Printf("Error opening minichain: %v\n",err)
		os.Exit(1)
	}
	mchain.SetStateDB(&db)

	last_line_rec,err := mchain.ReadLastLine()
	if err != nil {
		fmt.Printf("Error getting last record: %v\n",err)
		os.Exit(1)
	}
	fmt.Printf("Executing call on state root %v\n",last_line_rec.StateRoot.String())
	var rec Record
	rec.BlockNum = r.BlockNumber.Int64()
	rec.BlockNum = 12369621 // we have to hardcode this block because we are testing MainNet
	time_stamp := int64(1620131220)
	rec.BlockHash = r.BlockHash
	rec.TxIndex = int64(r.TransactionIndex)
	rec.TxHash = r.TxHash
	fmt.Printf("calling ExecCall with block %v\n",rec.BlockNum)
	err,state_root := mchain.ExecCall(chain_id,tx,rec.BlockNum,time_stamp ,last_line_rec.StateRoot,&rec)

	//err,state_root := UEVMAcctCreate(chain_id,tx_msg.From(),tx.Nonce(),db,common.Hash{})
	//generated_addr := common.Address{}
	fmt.Printf("Deploy result: %v\n",err)
	fmt.Printf("Intermediate state hash: %v\n",state_root.String())
}
