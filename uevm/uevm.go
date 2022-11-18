package uevm				// EVM for Uniswap (v3)

import (
	"fmt"
	"math/big"
	"math"
	//"encoding/hex"

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
	"github.com/ethereum/go-ethereum/crypto"
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
/* DISCONTINUED
func UEVMDeploy(chain_id int64,from common.Address,nonce uint64,contract_code []byte,sdb state.Database,state_root common.Hash) (error,common.Address,common.Hash) {	// deploys contract code

	fmt.Printf("from = %v\n",from.String())
	//state_db,err := state.New(state_root,state.NewDatabase(edb),nil)
	state_db,err := state.New(state_root,sdb,nil)
	if err != nil {
		return err,common.Address{},common.Hash{}
	}
	state_db.Prepare(common.Hash{},1)
	sender_bal := state_db.GetBalance(from)
	min_bal := big.NewInt(math.MaxInt)
	if sender_bal.Cmp(min_bal) < 0 {
		state_db.AddBalance(from,min_bal)	// fix the problem with balances on empty states
	}
	fmt.Printf("Dump after AddBalance\n")
	raw_dump := GetStateDump(state_db)
	DumpStateDB(raw_dump)

	state_db.SetNonce(from,nonce)
	chain_cfg := params.MainnetChainConfig
	//fmt.Printf("chain cfg:\n%+v\n",chain_cfg)
	vm_cfg := vm.Config{}
	block_ctx := NewDummyBlockContext(big.NewInt(12369621) ,big.NewInt(1620131220))
	tx_ctx := new(vm.TxContext)
	tx_ctx.Origin = from
	tx_ctx.GasPrice = big.NewInt(1111111111)
	evm := vm.NewEVM(*block_ctx,*tx_ctx,state_db,chain_cfg,vm_cfg)
	gas := uint64(99999999999)
	value := big.NewInt(0)
	sender := vm.AccountRef(from)
	contract_addr := crypto.CreateAddress(sender.Address(), state_db.GetNonce(sender.Address()))
	ret, _, _, vmerr := evm.Create(sender, contract_code, gas, value)
	_=ret
	fmt.Printf("Create() vmerr = %v\n",vmerr)
	fmt.Printf("Create() contract addr = %v\n",contract_addr.String())
	//fmt.Printf("Create() output: %v\n",hex.EncodeToString(ret))
	delete_empty_objects := false
	iroot_hash := state_db.Int0xb92c5707d43bca67c67a0bc59bdd267d318687f1b3b1cb3b5a852166649c7624ermediateRoot(delete_empty_objects)
	err = state_db.Database().TrieDB().Commit(iroot_hash, true, nil)
	if err != nil {
		fmt.Printf("Error on TrieDB().Commit(): %v\n",err)
		return err,contract_addr,iroot_hash
	}
	out_state,err := state_db.Commit(delete_empty_objects)
	if err!=nil {
		fmt.Printf("Error on state_db.Commit(): %v\n",err)
		return err,contract_addr,iroot_hash
	}
	fmt.Printf("state_hash after commit: %v\n",out_state.String())
	err = state_db.Database().TrieDB().Commit(out_state, true, nil)
	if err != nil {
		fmt.Printf("Error on TrieDB().Commit() for out_state: %v\n",err)
		return err,contract_addr,iroot_hash
	}

	raw_dump = GetStateDump(state_db)
	DumpStateDB(raw_dump)
	return vmerr,contract_addr,iroot_hash
}
*/
func UEVMDeploy2(chain_id int64,from common.Address,nonce uint64,contract_code []byte,sdb *state.Database,state_root common.Hash) (error,common.Address,common.Hash) {	// deploys contract code

	fmt.Printf("from = %v\n",from.String())
	state_db,err := state.New(state_root,*sdb,nil)
	if err != nil {
		return err,common.Address{},common.Hash{}
	}
	state_db.Prepare(common.Hash{},1)
	sender_bal := state_db.GetBalance(from)
	min_bal := big.NewInt(math.MaxInt)
	if sender_bal.Cmp(min_bal) < 0 {
		state_db.AddBalance(from,min_bal)	// fix the problem with balances on empty states
	}
	fmt.Printf("Dump after AddBalance\n")
	raw_dump := GetStateDump(state_db)
	DumpStateDB(raw_dump)

	state_db.SetNonce(from,nonce)
	chain_cfg := params.MainnetChainConfig
	//fmt.Printf("chain cfg:\n%+v\n",chain_cfg)
	vm_cfg := vm.Config{}
	block_ctx := NewDummyBlockContext(big.NewInt(12369621) ,big.NewInt(1620131220))
	tx_ctx := new(vm.TxContext)
	tx_ctx.Origin = from
	tx_ctx.GasPrice = big.NewInt(1111111111)
	evm := vm.NewEVM(*block_ctx,*tx_ctx,state_db,chain_cfg,vm_cfg)
	gas := uint64(99999999999)
	value := big.NewInt(0)
	sender := vm.AccountRef(from)
	contract_addr := crypto.CreateAddress(sender.Address(), state_db.GetNonce(sender.Address()))
	ret, _, _, vmerr := evm.Create(sender, contract_code, gas, value)
	_=ret
	fmt.Printf("Create() vmerr = %v\n",vmerr)
	fmt.Printf("Create() contract addr = %v\n",contract_addr.String())
	//fmt.Printf("Create() output: %v\n",hex.EncodeToString(ret))
	delete_empty_objects := false
	out_state,err := state_db.Commit(delete_empty_objects)
	if err!=nil {
		fmt.Printf("Error on state_db.Commit(): %v\n",err)
		return err,contract_addr,out_state
	}
	fmt.Printf("state_hash after commit: %v\n",out_state.String())
	err = state_db.Database().TrieDB().Commit(out_state, true, nil)
	if err != nil {
		fmt.Printf("Error on TrieDB().Commit() for out_state: %v\n",err)
		return err,contract_addr,out_state
	}
	err = state_db.Database().TrieDB().CommitPreimages()
	if err != nil {
		fmt.Printf("Error on TrieDB().CommitPreimages() for out_state: %v\n",err)
		return err,contract_addr,out_state
	}

	raw_dump = GetStateDump(state_db)
	DumpStateDB(raw_dump)
	return vmerr,contract_addr,out_state
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
func UEVMAcctCreate(chain_id int64,from common.Address,nonce uint64,sdb state.Database,state_root common.Hash) (error,common.Hash) {	// deploys contract code

	fmt.Printf("from = %v\n",from.String())
	//state_db,err := state.New(state_root,state.NewDatabase(edb),nil)
	state_db,err := state.New(state_root,sdb,nil)
	if err != nil {
		return err,common.Hash{}
	}
	state_db.Prepare(common.Hash{},1)
	min_bal := big.NewInt(math.MaxInt)
	_=min_bal
	state_db.AddBalance(from,min_bal)
	state_db.SetNonce(from,66)
	fmt.Printf("Dump after AddBalance\n")
	raw_dump := GetStateDump(state_db)
	DumpStateDB(raw_dump)

	delete_empty_objects := true
	/*
	iroot_hash := state_db.IntermediateRoot(delete_empty_objects)
	err = state_db.Database().TrieDB().Commit(iroot_hash, true, nil)
	if err != nil {
		fmt.Printf("Error on TrieDB().Commit(): %v\n",err)
		return err,iroot_hash
	}*/
	out_state,err := state_db.Commit(delete_empty_objects)
	if err!=nil {
		fmt.Printf("Error on state_db.Commit(): %v\n",err)
		return err,out_state
	}
	fmt.Printf("state_hash after commit: %v\n",out_state.String())
	err = state_db.Database().TrieDB().Commit(out_state, true, nil)
	if err != nil {
		fmt.Printf("Error on TrieDB().Commit() for out_state: %v\n",err)
		return err,out_state
	}

	raw_dump = GetStateDump(state_db)
	DumpStateDB(raw_dump)
	return err,out_state
}
//func OpenDB(file string) ethdb.Database {
func OpenDB(file string) state.Database {
	
	rdb,err := rawdb.NewLevelDBDatabase(file,0 ,0 ,"uniswapcustom",false)
	if err != nil {
		panic("can't open db")
	}
	var cfg trie.Config
	cfg.Preimages = true
	sdb := state.NewDatabaseWithConfig(rdb,&cfg)
	return sdb
}
func GetStateDump(s *state.StateDB) state.Dump {
	dump_config := state.DumpConfig{
		SkipCode:           false,
		SkipStorage:        false,
		OnlyWithAddresses:  false,
		Start:              nil,
		Max:                0,
	}
	raw_dump := s.RawDump(&dump_config)
	return raw_dump
}
func DumpStateDB(raw_dump state.Dump) {
	fmt.Printf("Dump of trie root = %v\n",raw_dump.Root)
	fmt.Printf("Total objects: %v\n",len(raw_dump.Accounts))
	fmt.Printf("Address\t\t\tBalance\t\tNonce\n")
	keys := make([]common.Address,len(raw_dump.Accounts))
	i:=0
	for k := range raw_dump.Accounts {
		keys[i] = k
		i++
	}
	for i:=0;i<len(keys);i++ {
		obj := raw_dump.Accounts[keys[i]]
		astr := ""
		if obj.Address != nil {
			astr = obj.Address.String()
		}
		fmt.Printf("%v (%v)\t%v\t\t%v\n",keys[i].String(),astr,obj.Balance,obj.Nonce)
	}
}
/*
	triedb := trie.NewDatabase(edb)
	regenerate_snapshot := true
	sshot,err := snapshot.New(edb,triedb,256,common.Hash{},false,regenerate_snapshot,false)
	if err != nil {
		return err,common.Address{},common.Hash{}
	}
	state_db,err := state.New(state_root,state.NewDatabase(edb),sshot)
*/
