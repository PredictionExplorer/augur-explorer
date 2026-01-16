package uevm				// EVM for Uniswap (v3)

import (
	"fmt"
	"math/big"
	"math"
	"encoding/hex"
	"strings"
	//"errors"

	"github.com/ethereum/go-ethereum/ethdb"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/rlp"
	//"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/core/vm"
	"github.com/ethereum/go-ethereum/core"
	"github.com/ethereum/go-ethereum/core/rawdb"
	"github.com/ethereum/go-ethereum/core/state"
	"github.com/ethereum/go-ethereum/trie"
	"github.com/ethereum/go-ethereum/params"
	//"github.com/ethereum/go-ethereum/crypto"

	contracts "github.com/PredictionExplorer/augur-explorer/contracts"
)
var (
	ChainID			int64	= 1234
	MainNetBlockNum		int64 = 12369621
	MainNetTimeStamp	int64 = 1620131220
	TxDefaultGas		int64 = 1111111111
	eclient				*ethclient.Client
)
func SetEClient(c *ethclient.Client) {
	eclient = c
}
func UEVMDeploy(chain_id int64,tx_hash common.Hash,from common.Address,nonce uint64,contract_code []byte,contract_addr common.Address,sdb *state.Database,state_root common.Hash) (error,common.Address,common.Hash,[]byte) {	// deploys contract code

	fmt.Printf("from = %v\n",from.String())
	state_db,err := state.New(state_root,*sdb,nil)
	if err != nil {
		fmt.Printf("Error at state.New(): %v\n",err)
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
	vm_cfg := vm.Config{}
	block_ctx := NewDummyBlockContext(big.NewInt(MainNetBlockNum) ,big.NewInt(MainNetTimeStamp))
	tx_ctx := new(vm.TxContext)
	tx_ctx.Origin = from
	tx_ctx.GasPrice = big.NewInt(TxDefaultGas)
	evm := vm.NewEVM(*block_ctx,*tx_ctx,state_db,chain_cfg,vm_cfg)
	gas := uint64(99999999999)
	value := big.NewInt(0)
	sender := vm.AccountRef(from)
	ret, _, _, vmerr := evm.Create3(sender, contract_code, gas, value,contract_addr)
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
		fmt.Printf("Errr on state_db.Commit(): %v\n",err)
		return err,contract_addr,out_state,nil
	}
	fmt.Printf("state_hash after commit: %v\n",out_state.String())
	/*
	_,err = (*sdb).TrieDB().Node(out_state)
	if err == nil {
		return errors.New(fmt.Sprintf("State %v already exists in the Trie DB, aborting",out_state.String())),common.Address{},common.Hash{},nil
	}*/
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
func UEVMDeployDummyToken(block_ctx *vm.BlockContext,tx_hash common.Hash,tx_ctx *vm.TxContext,to common.Address,symbol,name string, decimals uint8,state_root common.Hash,sdb *state.Database )	(error,common.Hash,[]byte) {
	// Note: tx_hash has to be altered from original Mint tx hash by inserting 'token0' or 'token1' string converted to bytes whithin the first 6 bytes of the transaction hash (so the hash doesn't collide with Mint's hash)
	// (this is required because we have to insert token accounts before executing Mint call)
	fmt.Printf("UEVMDEployDummyToken(): contract_addr %v , sym %v name %v decimals %v, state root hash is %v\n",to.String(),symbol,name,decimals,state_root.String())
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
	block_ctx.BlockNumber.SetInt64(MainNetBlockNum)
	block_ctx.Time.SetInt64(MainNetTimeStamp)
	chain_cfg := params.MainnetChainConfig
	vm_cfg := vm.Config{}
	evm := vm.NewEVM(*block_ctx,*tx_ctx,state_db,chain_cfg,vm_cfg)
	gas := uint64(99999999999)
	value := big.NewInt(0)
	sender := vm.AccountRef(tx_ctx.Origin)
	contract_code := common.FromHex(contracts.ERC20UnlimitedMetaData.Bin)
	abi_parsed := strings.NewReader(contracts.ERC20UnlimitedMetaData.ABI)
	dummyerc20_abi,err := abi.JSON(abi_parsed)
	if err != nil {
		fmt.Printf("Error parsing abi: %v\n",err)
		return err,common.Hash{},nil
	}
	input, err := dummyerc20_abi.Pack("",name,symbol,decimals) // "" - means constructor args
	if err != nil {
		fmt.Printf("Error in packing input arguments of DummyERC token: %v\n",err)
		return err,common.Hash{},nil
	}
	fmt.Printf("input hex %v\n",hex.EncodeToString(input))
	input = append(contract_code, input...)
	ret, _, _, vmerr := evm.Create3(sender, input, gas, value,to)
	maxlen := len(ret)
	if maxlen > 128 { maxlen = 128}
	fmt.Printf("evm.Create3() returns %v\n",hex.EncodeToString(ret[0:maxlen]))
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
func UEVMCall(block_ctx *vm.BlockContext,tx_hash common.Hash,tx_ctx *vm.TxContext,input []byte,value *big.Int,to common.Address,state_root common.Hash,sdb *state.Database) (error,common.Hash,[]byte) {

	fmt.Printf("Executing call2(tx=%v) to contract %v on state %v\n",tx_hash.String(),to.String(),state_root.String())
	state_db,err := state.New(state_root,*sdb,nil)
	if err != nil {
		return err,common.Hash{},nil
	}
	inputlen := len(input)
	if inputlen>128 { inputlen=123 }
	fmt.Printf("input: %v\n",hex.EncodeToString(input[0:inputlen]))
	state_db.Prepare(tx_hash,1)
	chain_cfg := params.MainnetChainConfig
	vm_cfg := vm.Config{}
	block_ctx.BlockNumber.SetInt64(MainNetBlockNum)
	block_ctx.Time.SetInt64(MainNetTimeStamp)
	evm := vm.NewEVM(*block_ctx,*tx_ctx,state_db,chain_cfg,vm_cfg)
	gp := new(core.GasPool)
	gas := uint64(99999999999)
	sender := vm.AccountRef(tx_ctx.Origin)
	ret, gas, vmerr := evm.Call(sender, to , input, gas,value) 
	fmt.Printf("Call() ended , return value length: %v bytes\n",len(ret))
	_=ret; _=gas;_=gp;
	if len(ret)>0 {
		end := len(ret)
		if end > 128 { end = 128 }
		fmt.Printf("Call return value: %v\n",hex.EncodeToString(ret[0:end]))
	} else {
		fmt.Printf("Call return value is nil\n")
	}
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
func OpenDB(file string) (ethdb.Database,state.Database) {
	
	fmt.Printf("OpenDB(file = %v)\n",file)
	rdb,err := rawdb.NewLevelDBDatabase(file,0 ,0 ,"uniswapcustom",false)
	if err != nil {
		panic("can't open db")
	}
	var cfg trie.Config
	cfg.Preimages = true
	sdb := state.NewDatabaseWithConfig(rdb,&cfg)
	return rdb,sdb
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
