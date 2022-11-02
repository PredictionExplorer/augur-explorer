package uevm				// EVM for Uniswap (v3)

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/core/vm"
	"github.com/ethereum/go-ethereum/core"
	"github.com/ethereum/go-ethereum/core/rawdb"
	"github.com/ethereum/go-ethereum/core/ethdb"
	"github.com/ethereum/go-ethereum/core/state"
	"github.com/ethereum/go-ethereum/core/state/snapshot"
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
func UEVMCall(tx *types.Transaction,state_root common.Hash,edb ethdb.Database) {

	sshot := snapshot.New(edb)
	statedb := state.New(state.NewDatabase(edb),sshot)
	evm := core.NewEVM(blockCtx,txCtx,stateDB,chain_cfg,cfg)
	gp := new(core.GasPool)
	tx_msg := tx.AsMessage()

	st := evm.NewStateTransition(evm,tx_msg,gp)
	ret, st.gas, vmerr = st.evm.Call(sender, st.to(), st.data, st.gas, st.value)
}
func OpenDB(file string) ethdb.Database {
	return rawdb.NewLevelDBDatabase(file,0 ,0 ,"uniswapcustom",false)
}
