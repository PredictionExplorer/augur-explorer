package uevm				// EVM for Uniswap (v3)
import (
	"os"
	"fmt"
	"errors"
	"strconv"
	"math/big"
	"encoding/hex"

	//"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/state"
	"github.com/ethereum/go-ethereum/core/vm"
	//"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/ethdb/leveldb"

	contracts "github.com/PredictionExplorer/augur-explorer/contracts"
)
// TechNotes:
//		the file stores one record at one line, each line is terminated by '\n' character
//		Field list:
//		1			Block number in decimal, padded with 0s, length = 9 characters
//		2			Block's hash in hexadecimal, 0x prepended, length = 66 characters
//		3			Transaction index in decimal, padded with 0s, legnth = 5 characters
//		5			Transaction hash in hexadecimal, 0x prepended, length = 66 characters
//		6			State Root hash in hexadecimal, 0x prepended, length = 66 characters
//
//		Field separator is 'space' character
const (
	RECORD_LENGTH			int = 217
)
type (
	MiniChain struct {
		StatesFileName		string		// name of the file which stores state root hashes
		F					*os.File
		sdb					*state.Database
		receipts_db			*leveldb.Database
	}
	Record struct {
		BlockNum			int64
		BlockHash			common.Hash
		TxIndex				int64
		TxHash				common.Hash
		StateRoot			common.Hash
	}
	DummyERC_Spec struct {
		Address				common.Address
		Symbol				string
		Name				string
		Decimals			uint8
	}
)
func OpenMiniChain(states_fname,receipts_datadir string) (MiniChain,error) {

	var err error
	var mc MiniChain
	mc.StatesFileName = states_fname
	mc.F,err = os.OpenFile(states_fname, os.O_RDWR|os.O_CREATE, 0755)
	if err != nil {
		return mc,err
	}
	mc.receipts_db,err = leveldb.New(receipts_datadir,0 ,0 ,"receipts",false)
	return mc,err
}
func (self *MiniChain) CloseMiniChain() {
	self.F.Close()
}
func (self *MiniChain) SetStateDB(state *state.Database) {

	self.sdb = state
}
func (self *MiniChain) NumRecords() int64 {
	fileinfo,err := os.Stat(self.StatesFileName)
	if err != nil {
		panic(fmt.Sprintf("NumRecords(): error on os.Stat() call: %v\n",err))
	}
	mod_flen := fileinfo.Size() % int64(RECORD_LENGTH) // safety check
	if mod_flen != 0 {
		panic(fmt.Sprintf("Size of minichain file is not divisible by RECORD_LEN (mod=%v)",mod_flen))
	}
	flen := fileinfo.Size() / int64(RECORD_LENGTH)
	return flen
}
func (self *MiniChain) ReadLastLine() (Record,error) {

	var r Record
	_,err := self.F.Seek(int64(-RECORD_LENGTH),2)
	if err != nil {
		return r,err
	}

	var count int
	rline := make([]byte,RECORD_LENGTH)
	count,err = self.F.Read(rline)
	if err != nil {
		return r,err
	}
	if count != RECORD_LENGTH {
		return r,errors.New("Didn't read the amount of bytes required to match record length")
	}
	rec := rline[:RECORD_LENGTH]
	block_num,err := strconv.ParseInt(string(rec[0:9]),10,64)
	if err != nil { return r,err }
	r.BlockNum = block_num
	tx_index,err := strconv.ParseInt(string(rec[77:77+5]),10,64)
	if err != nil { return r,err }
	r.TxIndex = tx_index
	r.BlockHash = common.HexToHash(string(rec[10:10+66]))
	r.TxHash= common.HexToHash(string(rec[83:83+66]))
	r.StateRoot = common.HexToHash(string(rec[150:150+66]))
	fmt.Printf("state root = %v\n",r.StateRoot.String())

	return r,err
}
func (self *MiniChain) AppendLine(r *Record) error {

	var output string
	output = output + fmt.Sprintf("%09d ",r.BlockNum)
	output = output + r.BlockHash.String()+ " "
	output = output + fmt.Sprintf("%05d ",r.TxIndex)
	output = output + r.TxHash.String() + " "
	output = output + r.StateRoot.String()+"\n"
	fmt.Printf("AppendLine, output len=%v\n",len(output))
	_,err := self.F.Seek(0 ,2)
	if err != nil {
		return err
	}
	bytes := []byte(output)
	count,err := self.F.Write(bytes)
	if err != nil {
		return err
	}
	if count != len(bytes) {
		panic("AppendLine() wrote less bytes than expected")
	}
	return err
}
func (self *MiniChain) ExecDeploy(chain_id int64,tx_hash common.Hash,from common.Address,nonce uint64,contract_code []byte,contract_addr common.Address,initial_state_root common.Hash,r *Record) (error,common.Address,common.Hash) {

	fmt.Printf("ExecDeploy(): tx_hash=%v\n",tx_hash.String())
	err,addr,root,encoded_logs := UEVMDeploy(chain_id,tx_hash,from,nonce,contract_code,contract_addr,self.sdb,initial_state_root)
	if err != nil {
		return err,addr,common.Hash{}
	}
	r.StateRoot = root
	err = self.AppendLine(r)
	if err != nil {
		return err,common.Address{},common.Hash{}
	}
	lenlogs := 0
	if encoded_logs != nil { lenlogs=len(encoded_logs) }
	fmt.Printf("Storing %v log bytes for tx hash %v\n",lenlogs,tx_hash.String())
	err = self.receipts_db.Put(tx_hash.Bytes(),encoded_logs)
	return err,addr,root
}
func (self *MiniChain) ExecCall(block_ctx *vm.BlockContext,tx_hash common.Hash,tx_ctx *vm.TxContext,input []byte,value *big.Int,contract_addr common.Address,initial_state_root common.Hash,r *Record) (error,common.Hash) {

	err,root,encoded_logs := UEVMCall(block_ctx,tx_hash,tx_ctx,input,value,contract_addr,initial_state_root,self.sdb)
	if err != nil {
		return err,common.Hash{}
	}
	r.StateRoot = root
	err = self.AppendLine(r)
	self.receipts_db.Put(tx_hash.Bytes(),encoded_logs)
	return err,root
}
func (self *MiniChain) ExecMint(block_ctx *vm.BlockContext,tx_hash common.Hash,tx_ctx *vm.TxContext,input []byte,value *big.Int,contract_addr common.Address,initial_state_root common.Hash,r *Record,token0_addr,token1_addr common.Address) error {

	err,tmp_state_hash := self.MaybeAddDummyTokens(block_ctx,tx_hash,tx_ctx,initial_state_root,token0_addr,token1_addr,r)
	if err != nil {
		return err
	}
	initial_state_root.SetBytes(tmp_state_hash.Bytes())
	err,new_state_root := self.ExecCall(block_ctx,tx_hash,tx_ctx,input,value,contract_addr,initial_state_root,r)
	if err != nil {
		return err
	}
	r.StateRoot=new_state_root
	self.AppendLine(r)
	return err
}
func get_erc20_specs(addr common.Address) (DummyERC_Spec,error) {

	var ERC20Spec DummyERC_Spec
	var copts = new(bind.CallOpts)

	erc20_ctrct,err := contracts.NewERC20Unlimited(addr,eclient)
	if err != nil {
		return ERC20Spec,errors.New(fmt.Sprintf("Error creating ERC20 instance: %v",err.Error()))
	}
	symbol,err := erc20_ctrct.Symbol(copts)
	if err != nil {
		return ERC20Spec,errors.New(fmt.Sprintf("Error on Symbol(): %v",err.Error()))
	}
	name,err := erc20_ctrct.Name(copts)
	if err != nil {
		return ERC20Spec,errors.New(fmt.Sprintf("Error on Name(): %v",err.Error()))
	}
	decimals,err := erc20_ctrct.Decimals(copts)
	if err != nil {
		return ERC20Spec,errors.New(fmt.Sprintf("Error on Decimals(): %v",err.Error()))
	}
	ERC20Spec.Symbol = symbol
	ERC20Spec.Name = name
	ERC20Spec.Decimals = decimals
	return ERC20Spec,nil
}
func (self *MiniChain) MaybeAddDummyTokens(block_ctx *vm.BlockContext,tx_hash common.Hash,tx_ctx *vm.TxContext,state_root common.Hash,t0_addr,t1_addr common.Address,r *Record) (error,common.Hash){
	// adds dummy ERC20 token contracts with addresses of real tokens
	var add_token0 bool = false
	var add_token1 bool = false
	clone_rec := *r
	clone_rec.TxHash = common.BytesToHash(r.TxHash.Bytes())
	new_root := state_root
	if !self.AccountExists(state_root,t0_addr) {
		add_token0 = true
	}
	if !self.AccountExists(state_root,t1_addr) {
		add_token1 = true
	}
	if add_token0 {
		erc20_specs,err := get_erc20_specs(t0_addr)
		if err != nil {
			return err,common.Hash{}
		}
		fmt.Printf("Deploying token0 contract (addr %v) (%v %v)\n",t1_addr.String(),erc20_specs.Symbol,erc20_specs.Name)
		tmphash := common.BytesToHash(clone_rec.TxHash.Bytes())
		hbytes := tmphash.Bytes()
		hbytes[0]=116; hbytes[1]=111; hbytes[2]=107; hbytes[3]=101; hbytes[4]=110; hbytes[5]=48; // 'token0'
		newtxhash:=common.BytesToHash(hbytes)
		fmt.Printf("altered bytes hash %v\n",hex.EncodeToString(hbytes))
		fmt.Printf("new hash = %v\n",newtxhash.String())
		err,tmp_new_root,encoded_logs := UEVMDeployDummyToken(block_ctx,newtxhash,tx_ctx,t0_addr,erc20_specs.Symbol,erc20_specs.Name,erc20_specs.Decimals,new_root,self.sdb)
		if err != nil {
			err_str := fmt.Sprintf("Deployment of ERC20 token1 failed: %v",err.Error())
			return errors.New(err_str),common.Hash{}
		}
		fmt.Printf("MaybeAddDummyToken() token 0 root hash %v\n",tmp_new_root.String())
		clone_rec.StateRoot.SetBytes(tmp_new_root.Bytes())
		clone_rec.TxHash.SetBytes(newtxhash.Bytes())
		err = self.AppendLine(&clone_rec)
		self.receipts_db.Put(newtxhash.Bytes(),encoded_logs)
		new_root.SetBytes(tmp_new_root.Bytes())
	}
	if add_token1 {
		erc20_specs,err := get_erc20_specs(t1_addr)
		if err != nil {
			return err,common.Hash{}
		}
		fmt.Printf("Deploying token0 contract (addr %v) (%v %v)\n",t1_addr.String(),erc20_specs.Symbol,erc20_specs.Name)
		tmphash := common.BytesToHash(clone_rec.TxHash.Bytes())
		hbytes := tmphash.Bytes()
		hbytes[0]=116; hbytes[1]=111; hbytes[2]=107; hbytes[3]=101; hbytes[4]=110; hbytes[5]=49; // 'token1'
		fmt.Printf("altered bytes hash %v\n",hex.EncodeToString(hbytes))
		newtxhash:=common.BytesToHash(hbytes)
		fmt.Printf("new hash = %v\n",newtxhash.String())
		err,tmp_new_root,encoded_logs := UEVMDeployDummyToken(block_ctx,newtxhash,tx_ctx,t1_addr,erc20_specs.Symbol,erc20_specs.Name,erc20_specs.Decimals,new_root,self.sdb)
		if err != nil {
			err_str := fmt.Sprintf("Deployment of ERC20 token1 failed: %v",err.Error())
			return errors.New(err_str),common.Hash{}
		}
		fmt.Printf("MaybeAddDummyToken() token 1 root hash %v\n",tmp_new_root.String())
		clone_rec.StateRoot.SetBytes(tmp_new_root.Bytes())
		clone_rec.TxHash.SetBytes(newtxhash.Bytes())
		err = self.AppendLine(&clone_rec)
		self.receipts_db.Put(newtxhash.Bytes(),encoded_logs)
		new_root.SetBytes(tmp_new_root.Bytes())
	}
	fmt.Printf("MaybeAddDummyTokens() returns new root hash %v\n",new_root.String())
	return nil,new_root
}
func (self *MiniChain) AccountExists(state_root common.Hash,addr common.Address) bool {

	state_db,err := state.New(state_root,*self.sdb,nil)
	if err != nil {
		panic(fmt.Sprintf("Cant create new StateDB object: %v",err))
	}

	return state_db.Exist(addr)
}
