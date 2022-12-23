package main

import (
	"os"
	"fmt"
	"errors"
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
	DEFAULT_MCHAIN_NAME		string = "minichain.dat"
	DEFAULT_RECEIPTS_NAME	string = "receipts"
	DEFAULT_EVM_DB_NAME		string = "evmdb"
)
func redeploy(mchain *MiniChain,tx_hash common.Hash,initial_state_hash common.Hash) (common.Address,error) {

	RPC_URL = os.Getenv("RPC_URL")
	eclient, err := ethclient.Dial(RPC_URL)
	if err!=nil {
		fmt.Printf("Error: %v\n",err)
		os.Exit(1)
	}
	tx,_,err := eclient.TransactionByHash(context.Background(),tx_hash)
	if err != nil { return common.Address{},err }
	rcpt,err := eclient.TransactionReceipt(context.Background(),tx_hash)
	if err != nil { return common.Address{},err }
	contract_address := rcpt.ContractAddress
	tx_msg,err := tx.AsMessage(types.LatestSignerForChainID(big.NewInt(chain_id)),tx.GasPrice())
	if err != nil { return common.Address{},err }
	
	var rec Record
	rec.BlockNum = MainNetBlockNum // we have to hardcode this block because we are testing MainNet
	rec.BlockHash = rcpt.BlockHash
	rec.TxIndex = int64(rcpt.TransactionIndex)
	rec.TxHash = rcpt.TxHash
	input := tx_msg.Data()
	err,generated_addr,_ := mchain.ExecDeploy(chain_id,tx_hash,tx_msg.From(),tx.Nonce(),input,contract_address,initial_state_hash,&rec)
	return generated_addr,err
}
func main() {

	if len(os.Args) < 2 {
		fmt.Printf(
			"Usage: \n\t\t%v [dir] [factory_tx] [nft_descriptor_tx] [nft_manager_tx] [swaprouter_tx]\n\t\t"+
			"Creates required files/databases in [dir] using transaction hashes to main contract deployments (note: transaction hashes, not contract addresses)\"\n\n",os.Args[0],
		)
		os.Exit(1)
	}
	minichain_fname := fmt.Sprintf("%v/%v",os.Args[1],DEFAULT_MCHAIN_NAME)
	receipts_dir := fmt.Sprintf("%v/%v",os.Args[1],DEFAULT_RECEIPTS_NAME)
	evm_dir := fmt.Sprintf("%v/%v",os.Args[1],DEFAULT_EVM_DB_NAME)
	factory_hash := common.HexToHash(os.Args[2])
	nft_descriptor_hash := common.HexToHash(os.Args[3])
	nft_manager_hash := common.HexToHash(os.Args[4])
	swaprouter_hash := common.HexToHash(os.Args[5])

	fmt.Printf("Using %v for EVM db\n",evm_dir)
	fmt.Printf("Using %v for receipts\n",receipts_dir)
	fmt.Printf("Using %v for storing minichain\n",minichain_fname)
	_,db := OpenDB(evm_dir)
	if _, err := os.Stat(minichain_fname); errors.Is(err, os.ErrNotExist) {
			// file does not exist
	} else {
		fmt.Printf("File %v already exists, refusing to overwrite existing file\n",minichain_fname)
		os.Exit(1)
	}
	mchain,err := OpenMiniChain(minichain_fname,receipts_dir)
	if err != nil {
		fmt.Printf("Error opening minichain: %v\n",err)
		os.Exit(1)
	}
	mchain.SetStateDB(&db)
	var rec Record
	err = mchain.AppendLine(&rec)
	if err != nil {
		fmt.Printf("Error: %v\n",err)
		os.Exit(1)
	}
	var caddr common.Address

	caddr,err = redeploy(&mchain,factory_hash,rec.StateRoot)
	if err != nil { fmt.Printf("Error: %v\n",err); os.Exit(1); }
	fmt.Printf("Deployed Factory contract: %v\n",caddr.String())

	rec,err = mchain.ReadLastLine()
	if err != nil { fmt.Printf("Error: %v\n",err); os.Exit(1); }
	caddr,err = redeploy(&mchain,nft_descriptor_hash,rec.StateRoot)
	if err != nil { fmt.Printf("Error: %v\n",err); os.Exit(1); }
	fmt.Printf("Deployed NFT descriptor contract: %v\n",caddr.String())

	rec,err = mchain.ReadLastLine()
	if err != nil { fmt.Printf("Error: %v\n",err); os.Exit(1); }
	caddr,err = redeploy(&mchain,nft_manager_hash,rec.StateRoot)
	if err != nil { fmt.Printf("Error: %v\n",err); os.Exit(1); }
	fmt.Printf("Deployed NFT Manager contract: %v\n",caddr.String())

	rec,err = mchain.ReadLastLine()
	if err != nil { fmt.Printf("Error: %v\n",err); os.Exit(1); }
	caddr,err = redeploy(&mchain,swaprouter_hash,rec.StateRoot)
	if err != nil { fmt.Printf("Error: %v\n",err); os.Exit(1); }
	fmt.Printf("Deployed SwapRouter contract: %v\n",caddr.String())

	fmt.Printf("Done\n")
}
