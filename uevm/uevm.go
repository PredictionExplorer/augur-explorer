package uevm				// EVM for Uniswap (v3)

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/core/vm"
	"github.com/ethereum/go-ethereum/core"
	"github.com/ethereum/go-ethereum/core/rawdb"
	"github.com/ethereum/go-ethereum/ethdb"
	"github.com/ethereum/go-ethereum/core/state"
	"github.com/ethereum/go-ethereum/core/state/snapshot"
	"github.com/ethereum/go-ethereum/trie"
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
func UEVMCall(tx *types.Transaction,state_root common.Hash,edb ethdb.Database) error {

	sshot,err := snapshot.New(edb,trie.NewDatabase(edb),256,common.Hash{},false,false,false)
	if err != nil {
		return err
	}
	statedb,err := state.New(state_root,state.NewDatabase(edb),sshot)
	if err != nil {
		return err
	}
	block_ctx := new(vm.BlockContext)
	block_ctx.
	evm := vm.NewEVM(blockCtx,txCtx,stateDB,chain_cfg,cfg)
	gp := new(core.GasPool)
	tx_msg := tx.AsMessage()

	st := evm.NewStateTransition(evm,tx_msg,gp)
	ret, st.gas, vmerr = st.evm.Call(sender, st.to(), st.data, st.gas, st.value)
	return vmerr
}
func OpenDB(file string) ethdb.Database {
	return rawdb.NewLevelDBDatabase(file,0 ,0 ,"uniswapcustom",false)
}
