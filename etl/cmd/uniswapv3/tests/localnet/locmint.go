package main

import (
	"os"
	"fmt"
	"context"
	"math/big"
	"encoding/hex"
	"bytes"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/core/vm"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"

	. "github.com/PredictionExplorer/augur-explorer/uevm"
)
var (
	RPC_URL					string
	chain_id				int64 = 1234
)
func main() {

	if len(os.Args) < 3 {
		fmt.Printf(
			"Usage: \n\t\t%v [tx_hash] [contract_addr]\n\t\t"+
			"executes Mint() transaction on local db\n\n",os.Args[0],
		)
		os.Exit(1)
	}
	tx_hash := common.HexToHash(os.Args[1])
	ctrct_addr := common.HexToAddress(os.Args[2])
	_,db := OpenDB("/var/tmp/evmdb")
	fmt.Printf("db = %+v\n",db)

	RPC_URL = os.Getenv("RPC_URL")
	eclient, err := ethclient.Dial(RPC_URL)
	if err!=nil {
		fmt.Printf("Error: %v\n",err)
		os.Exit(1)
	}
	SetEClient(eclient)
	rcpt,err := eclient.TransactionReceipt(context.Background(),tx_hash)
	if err != nil {
		fmt.Printf("Error: %v\n",err)
		os.Exit(1)
	}
	logs := rcpt.Logs
	if len(logs) < 1 {
		fmt.Printf("Transaction has no logs\n")
		os.Exit(1)
	}
	var pool_addr common.Address
	topic_sig , err := hex.DecodeString("7a53080ba414158be7ec69b987b5fb7d07dee101fe85488f0853ae16239d0bde")
	found := bool(false)
	for i:=0; i<len(logs); i++ {
		elog := logs[i]
		if len(elog.Topics) == 0 { continue }
		if bytes.Equal(logs[i].Topics[0].Bytes(),topic_sig) {
			pool_addr.SetBytes(elog.Address.Bytes())
			found = true
			break;
		}
	}
	if !found {
		fmt.Printf("Mint event log wasn't found in this transaction\n")
		os.Exit(1)
	}
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
	fmt.Printf("pool address: %v\n",pool_addr.String())
	pool_ctrct,err := NewUniswapV3Pool(pool_addr,eclient)
	if err!=nil {
		fmt.Printf("Failed to instantiate Wrapped ETH contract: %v\n",err)
		os.Exit(1)
	}

	var copts = new(bind.CallOpts)
	token0_addr,err := pool_ctrct.Token0(copts)
	if err != nil {
		fmt.Printf("Error getting token0 address: %v\n",err)
		os.Exit(1)
	}
	token1_addr,err := pool_ctrct.Token1(copts)
	if err != nil {
		fmt.Printf("Error getting token1 address: %v\n",err)
		os.Exit(1)
	}

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
	fmt.Printf("Contract addr %v\n",ctrct_addr)
	var rec Record
	rec.BlockNum = rcpt.BlockNumber.Int64()
	rec.BlockNum = MainNetBlockNum // we have to hardcode this block because we are testing MainNet
	rec.BlockHash = rcpt.BlockHash
	rec.TxIndex = int64(rcpt.TransactionIndex)
	rec.TxHash = rcpt.TxHash
	fmt.Printf("calling ExecMint with block %v\n",rec.BlockNum)
	block_ctx := NewDummyBlockContext(big.NewInt(MainNetBlockNum) ,big.NewInt(MainNetTimeStamp))
	tx_ctx := new(vm.TxContext)
	tx_ctx.Origin = tx_msg.From()
	tx_ctx.GasPrice = big.NewInt(TxDefaultGas)
	input := tx_msg.Data()
	end := len(input)
	if end > 64 { end = 64 }
	fmt.Printf("Input: %v\n",hex.EncodeToString(input[0:end]))
	err = mchain.ExecMint(block_ctx,tx_hash,tx_ctx,input,tx_msg.Value(),ctrct_addr,last_line_rec.StateRoot,&rec,token0_addr,token1_addr)
	if err != nil {
		fmt.Printf("Error executing ExecMint(): %v\n",err)
		os.Exit(1)
	}
	fmt.Printf("Exeuction successful")
}
