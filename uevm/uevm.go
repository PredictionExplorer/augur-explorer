package uevm				// EVM for Uniswap (v3)

import (
	"fmt"
	"math/big"
	"math"
	"encoding/hex"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/rlp"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/core/vm"
	"github.com/ethereum/go-ethereum/core"
	"github.com/ethereum/go-ethereum/core/rawdb"
	"github.com/ethereum/go-ethereum/core/state"
	"github.com/ethereum/go-ethereum/trie"
	"github.com/ethereum/go-ethereum/params"
	"github.com/ethereum/go-ethereum/crypto"
)
var (
	ChainID			int64	= 1234
	MainNetBlockNum		int64 = 12369621
	MainNetTimeStamp	int64 = 1620131220
	TxDefaultGas		int64 = 1111111111
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
func UEVMDeploy2(chain_id int64,tx_hash common.Hash,from common.Address,nonce uint64,contract_code []byte,sdb *state.Database,state_root common.Hash) (error,common.Address,common.Hash,[]byte) {	// deploys contract code

	fmt.Printf("from = %v\n",from.String())
	state_db,err := state.New(state_root,*sdb,nil)
	if err != nil {
		return err,common.Address{},common.Hash{},nil
	}
	state_db.Prepare(tx_hash,1)
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
	logs := state_db.GetLogs(tx_hash,common.Hash{})
	logs_encoded_bytes,err := rlp.EncodeToBytes(logs)
	lenlogs := 0
	if logs!=nil { lenlogs=len(logs) }
	fmt.Printf("Deploy for tx_hash %v generated %v logs\n",tx_hash.String(),lenlogs)
	fmt.Printf("Create() vmerr = %v\n",vmerr)
	fmt.Printf("Create() contract addr = %v\n",contract_addr.String())
	//fmt.Printf("Create() output: %v\n",hex.EncodeToString(ret))
	delete_empty_objects := false
	out_state,err := state_db.Commit(delete_empty_objects)
	if err != nil {
		fmt.Printf("Error on state_db.Commit(): %v\n",err)
		return err,contract_addr,out_state,nil
	}
	fmt.Printf("state_hash after commit: %v\n",out_state.String())
	err = state_db.Database().TrieDB().Commit(out_state, true, nil)
	if err != nil {
		fmt.Printf("Error on TrieDB().Commit() for out_state: %v\n",err)
		return err,contract_addr,out_state,nil
	}
	err = state_db.Database().TrieDB().CommitPreimages()
	if err != nil {
		fmt.Printf("Error on TrieDB().CommitPreimages() for out_state: %v\n",err)
		return err,contract_addr,out_state,nil
	}

	raw_dump = GetStateDump(state_db)
	DumpStateDB(raw_dump)
	return vmerr,contract_addr,out_state,logs_encoded_bytes
}
func UEVMDeployDummyToken(block_ctx *vm.BlockContext,tx_hash common.Hash,tx_ctx *vm.TxContext,to common.Address,state_root common.Hash,sdb *state.Database )	(error,common.Hash,[]byte) {
	// Note: tx_hash has to be altered from original Mint tx hash by inserting 'token0' or 'token1' string converted to bytes whithin the first 6 bytes of the transaction hash (so the hash doesn't collide with Mint's hash)
	// (this is required because we have to insert token accounts before executing Mint call)
	state_db,err := state.New(state_root,*sdb,nil)
	if err != nil {
		return err,common.Hash{},nil
	}
	state_db.Prepare(tx_hash,1)
	sender_bal := state_db.GetBalance(tx_ctx.Origin)
	min_bal := big.NewInt(math.MaxInt)
	if sender_bal.Cmp(min_bal) < 0 {
		state_db.AddBalance(tx_ctx.Origin,min_bal)	// fix the problem with balances on empty states
	}
	chain_cfg := params.MainnetChainConfig
	vm_cfg := vm.Config{}
	evm := vm.NewEVM(*block_ctx,*tx_ctx,state_db,chain_cfg,vm_cfg)
	gas := uint64(99999999999)
	value := big.NewInt(0)
	sender := vm.AccountRef(tx_ctx.Origin)
	contract_code,err := hex.DecodeString(DummyERC20CodeStr)
	if err != nil {
		return err,common.Hash{},nil
	}
	ret, _, _, vmerr := evm.Create3(sender, contract_code, gas, value,to)
	_=ret
	logs := state_db.GetLogs(tx_hash,common.Hash{})
	logs_encoded_bytes,err := rlp.EncodeToBytes(logs)
	delete_empty_objects := false
	out_state,err := state_db.Commit(delete_empty_objects)
	if err != nil {
		return err,out_state,nil
	}
	fmt.Printf("state_hash after commit: %v\n",out_state.String())
	err = state_db.Database().TrieDB().Commit(out_state, true, nil)
	if err != nil {
		return err,out_state,nil
	}
	err = state_db.Database().TrieDB().CommitPreimages()
	if err != nil {
		return err,out_state,nil
	}

	raw_dump := GetStateDump(state_db)
	DumpStateDB(raw_dump)
	return vmerr,out_state,logs_encoded_bytes
}
func UEVMCall(chain_id int64,tx *types.Transaction,block_num,time_stamp int64,state_root common.Hash,sdb *state.Database) (error,common.Hash,[]byte) {

	state_db,err := state.New(state_root,*sdb,nil)
	if err != nil {
		return err,common.Hash{},nil
	}
	state_db.Prepare(tx.Hash(),1)
	block_ctx := NewDummyBlockContext(big.NewInt(block_num) ,big.NewInt(time_stamp))
	tx_ctx := new(vm.TxContext)
	chain_cfg := params.MainnetChainConfig
	vm_cfg := vm.Config{}
	fmt.Printf("Block num = %v\n",block_num)
	fmt.Printf("chain config: %+v\n",chain_cfg)
	evm := vm.NewEVM(*block_ctx,*tx_ctx,state_db,chain_cfg,vm_cfg)
	gp := new(core.GasPool)
	gas := uint64(99999999999)
	tx_msg,err := tx.AsMessage(types.LatestSignerForChainID(big.NewInt(chain_id)),tx.GasPrice())
	if err != nil {
		return err,common.Hash{},nil
	}
	from := tx_msg.From()
	sender := vm.AccountRef(tx_msg.From())
	tx_ctx.Origin = from
	tx_ctx.GasPrice = big.NewInt(1111111111)
	//st := core.NewStateTransition(evm,tx_msg,gp)
	to := tx_msg.To()
	if to == nil {
		panic("call to nil contract address")
	}
	ret, gas, vmerr := evm.Call(sender, *to , tx_msg.Data(), gas, tx_msg.Value())
	_=ret
	_=gas
	_=gp
	logs := state_db.GetLogs(tx.Hash(),common.Hash{})
	fmt.Printf("num logs = %v\n",len(logs))
	logs_encoded_bytes,err := rlp.EncodeToBytes(logs)
	delete_empty_objects := false
	out_state,err := state_db.Commit(delete_empty_objects)
	if err!=nil {
		fmt.Printf("Error on state_db.Commit(): %v\n",err)
		return err,common.Hash{},nil
	}
	fmt.Printf("state_hash after commit: %v\n",out_state.String())
	err = state_db.Database().TrieDB().Commit(out_state, true, nil)
	if err != nil {
		fmt.Printf("Error on TrieDB().Commit() for out_state: %v\n",err)
		return err,common.Hash{},nil
	}
	err = state_db.Database().TrieDB().CommitPreimages()
	if err != nil {
		fmt.Printf("Error on TrieDB().CommitPreimages() for out_state: %v\n",err)
		return err,common.Hash{},nil
	}
	raw_dump := GetStateDump(state_db)
	DumpStateDB(raw_dump)
	return vmerr,out_state,logs_encoded_bytes
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
func UEVMCall2(block_ctx *vm.BlockContext,tx_hash common.Hash,tx_ctx *vm.TxContext,input []byte,value *big.Int,to common.Address,state_root common.Hash,sdb *state.Database) (error,common.Hash,[]byte) {

	state_db,err := state.New(state_root,*sdb,nil)
	if err != nil {
		return err,common.Hash{},nil
	}
	state_db.Prepare(tx_hash,1)
	chain_cfg := params.MainnetChainConfig
	vm_cfg := vm.Config{}
	evm := vm.NewEVM(*block_ctx,*tx_ctx,state_db,chain_cfg,vm_cfg)
	gp := new(core.GasPool)
	gas := uint64(99999999999)
	sender := vm.AccountRef(tx_ctx.Origin)
	ret, gas, vmerr := evm.Call(sender, to , input, gas,value) 
	_=ret; _=gas;_=gp;
	logs := state_db.GetLogs(tx_hash,common.Hash{})
	logs_encoded_bytes,err := rlp.EncodeToBytes(logs)
	delete_empty_objects := false
	out_state,err := state_db.Commit(delete_empty_objects)
	if err!=nil {
		return err,common.Hash{},nil
	}
	fmt.Printf("state_hash after commit: %v\n",out_state.String())
	err = state_db.Database().TrieDB().Commit(out_state, true, nil)
	if err != nil {
		return err,common.Hash{},nil
	}
	err = state_db.Database().TrieDB().CommitPreimages()
	if err != nil {
		return err,common.Hash{},nil
	}
	raw_dump := GetStateDump(state_db)
	DumpStateDB(raw_dump)
	return vmerr,out_state,logs_encoded_bytes
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
