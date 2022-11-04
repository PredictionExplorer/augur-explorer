package uevm				// EVM for Uniswap (v3)

import (
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/core/vm"
	"github.com/ethereum/go-ethereum/core"
	"github.com/ethereum/go-ethereum/core/rawdb"
	"github.com/ethereum/go-ethereum/ethdb"
	"github.com/ethereum/go-ethereum/core/state"
	"github.com/ethereum/go-ethereum/core/state/snapshot"
	"github.com/ethereum/go-ethereum/trie"
	"github.com/ethereum/go-ethereum/params"
)
func DeployFactory() {

}
func DeployPool() {
	
}
func CallInitialize() {

}
func CallMint() {

}
func CallBurn() {

}
func CallSwapFn() {	// calls swap() function

}
func UEVMCall(chain_id int64,tx *types.Transaction,state_root common.Hash,edb ethdb.Database) error {

	sshot,err := snapshot.New(edb,trie.NewDatabase(edb),256,common.Hash{},false,false,false)
	if err != nil {
		return err
	}
	state_db,err := state.New(state_root,state.NewDatabase(edb),sshot)
	if err != nil {
		return err
	}
	block_ctx := new(vm.BlockContext)
	tx_ctx := new(vm.TxContext)
	chain_cfg := params.MainnetChainConfig
	vm_cfg := vm.Config{}
	evm := vm.NewEVM(*block_ctx,*tx_ctx,state_db,chain_cfg,vm_cfg)
	gp := new(core.GasPool)
	tx_msg,err := tx.AsMessage(types.LatestSignerForChainID(big.NewInt(chain_id)),tx.GasPrice())
	if err != nil {
		return err
	}

	sender := vm.AccountRef(tx_msg.From())
	//st := core.NewStateTransition(evm,tx_msg,gp)
	to := tx_msg.To()
	if to == nil {
		panic("call to nil contract address")
	}
	ret, gas, vmerr := evm.Call(sender, *to , tx_msg.Data(), tx_msg.Gas(), tx_msg.Value())
	_=ret
	_=gas
	_=gp
	return vmerr
}
func OpenDB(file string) ethdb.Database {
	rdb,err := rawdb.NewLevelDBDatabase(file,0 ,0 ,"uniswapcustom",false)
	if err != nil {
		panic("can't open db")
	}
	return rdb
}
