package uevm				// EVM for Uniswap (v3)
import (
	"os"
	"fmt"
	"errors"
	"strconv"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/state"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethdb/leveldb"
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
func (self *MiniChain) ReadLastLine() (Record,error) {

	var r Record
	/*flen,err := os.Stat(self.StatesFileName)
	if err != nil {
		return r,err
	}
	flen = flen - RECORD_LENGTH*/
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
func (self *MiniChain) ExecDeploy(chain_id int64,tx_hash common.Hash,from common.Address,nonce uint64,contract_code []byte,initial_state_root common.Hash,r *Record) (error,common.Address,common.Hash) {

	fmt.Printf("ExecDeploy(): tx_hash=%v\n",tx_hash.String())
	err,addr,root,encoded_logs := UEVMDeploy2(chain_id,tx_hash,from,nonce,contract_code,self.sdb,initial_state_root)
	if err != nil {
		return err,addr,common.Hash{}
	}
	r.StateRoot = root
	err = self.AppendLine(r)
	lenlogs := 0
	if encoded_logs != nil { lenlogs=len(encoded_logs) }
	fmt.Printf("Storing %v log bytes for tx hash %v\n",lenlogs,tx_hash.String())
	self.receipts_db.Put(tx_hash.Bytes(),encoded_logs)
	return err,addr,root
}
func (self *MiniChain) ExecCall(chain_id int64,tx *types.Transaction,block_num,time_stamp int64,initial_state_root common.Hash,r *Record) (error,common.Hash) {

	err,root,encoded_logs := UEVMCall(chain_id,tx,block_num,time_stamp,initial_state_root,self.sdb)
	if err != nil {
		return err,common.Hash{}
	}
	r.StateRoot = root
	err = self.AppendLine(r)
	self.receipts_db.Put(tx.Hash().Bytes(),encoded_logs)
	return err,root
}
func (self *MiniChain) ExecMint(chain_id int64,tx *types.Transaction,block_num,time_stamp int64,initial_state_root common.Hash,r *Record,token0_addr,token1_addr common.Address) error {

	if !self.AccountExists(initial_state_root,token0_addr) {

	}
	if !self.AccountExists(initial_state_root,token1_addr) {

	}
	err,_ := self.ExecCall(chain_id,tx,block_num,time_stamp,initial_state_root,r)
	if err != nil {
		return err
	}
	return err
}
func (self *MiniChain) AddDummyTokens(state_root common.Hash,addr1,addr2 common.Address) (error,common.Hash){
	// adds dummy ERC20 token contracts with addresses of real tokens
	return nil,common.Hash{}
}
func (self *MiniChain) AccountExists(state_root common.Hash,addr common.Address) bool {

	state_db,err := state.New(state_root,*self.sdb,nil)
	if err != nil {
		panic(fmt.Sprintf("Cant create new StateDB object: %v",err))
	}

	return state_db.Exist(addr)
}
