package main

import (
	"os"
	"fmt"
	"errors"
	"context"
	"math/big"
	"encoding/hex"

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
func deploycode(mchain *MiniChain,code []byte,contract_address common.Address,tx_hash common.Hash,initial_state_hash common.Hash) (common.Address,error) {

	var rec Record
	rec.BlockNum = MainNetBlockNum // we have to hardcode this block because we are testing MainNet
	rec.TxHash = tx_hash
	err,generated_addr,_ := mchain.ExecDeploy(chain_id,tx_hash,common.Address{},0, code,contract_address,initial_state_hash,&rec)
	return generated_addr,err
}
func main() {

	if len(os.Args) < 7 {
		fmt.Printf(
			"Usage: \n\t\t%v [dir] [weth_addr] [factory_addr] [nft_position_descriptor_addr] [nft_manager_addr] [swaprouter_addr]\n\t\t"+
			"Deployes all necesary contracts to the MiniChain (the chain that keeps tcks of local EVM states) using specified addresses. The contract binary deployed is a debug version of Uniswap v3 contracts, they are not the same as Uniswap v3 on the MainNet, they lack callbacks and some of the require() statements. \n\n",os.Args[0],
		)
		os.Exit(1)
	}
	defdir := os.Args[1]
	minichain_fname := fmt.Sprintf("%v/%v",os.Args[1],DEFAULT_MCHAIN_NAME)
	receipts_name := fmt.Sprintf("%v/%v",defdir,DEFAULT_RECEIPTS_NAME)
	evm_dir := fmt.Sprintf("%v/%v",defdir,DEFAULT_EVM_DB_NAME)
	weth_addr := common.HexToAddress(os.Args[2])
	factory_addr := common.HexToAddress(os.Args[3])
	nft_descriptor_addr := common.HexToAddress(os.Args[4])
	nft_manager_addr := common.HexToAddress(os.Args[5])
	swaprouter_addr := common.HexToAddress(os.Args[6])

	fmt.Printf("Using %v for EVM db\n",evm_dir)
	fmt.Printf("Using %v for receipts\n",receipts_name)
	fmt.Printf("Using %v for storing minichain\n",minichain_fname)
	_,db := OpenDB(evm_dir)
	if _, err := os.Stat(minichain_fname); errors.Is(err, os.ErrNotExist) {
			// file does not exist
	} else {
		fmt.Printf("File %v already exists, refusing to overwrite existing file\n",minichain_fname)
		os.Exit(1)
	}
	mchain,err := OpenMiniChain(minichain_fname,receipts_name)
	if err != nil {
		fmt.Printf("Error opening minichain: %v\n",err)
		os.Exit(1)
	}
	mchain.SetStateDB(&db)
	var rec Record
	err = mchain.AppendLine(&rec)
	if err != nil {
		fmt.Printf("Error at AppendLine(): %v\n",err)
		os.Exit(1)
	}

	weth_code,err := hex.DecodeString(WethBin)
	if err != nil {
		fmt.Printf("Error decoding WETH contract binary: %v\n",err)
		os.Exit(1)
	}
	fmt.Printf("Deploying WETH at addr %v\n",weth_addr.String())
	_,err = deploycode(&mchain,weth_code,weth_addr,common.Hash{},rec.StateRoot)
	if err != nil {
		fmt.Printf("Error deploying WETH to local state db : %v\n",err)
		os.Exit(1)
	}

	var caddr common.Address

	rec,err = mchain.ReadLastLine()
	if err != nil { fmt.Printf("Error at ReadLastLine(): %v\n",err); os.Exit(1); }

	factory_code,err := hexutil.Decode(DbgUniswapV3PoolMetaData.Bin)
	if err != nil { fmt.Printf("Error at factory code decode: %v\n",err); os.Exit(1); }
	_,err = deploycode(&mchain,weth_code,weth_addr,common.Hash{},rec.StateRoot)
	if err != nil {
		fmt.Printf("Error deploying WETH to local state db : %v\n",err)
		os.Exit(1)
	}

	caddr,err = redeploy(&mchain,factory_hash,rec.StateRoot)
	if err != nil { fmt.Printf("Error after redeploy(): %v\n",err); os.Exit(1); }
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
