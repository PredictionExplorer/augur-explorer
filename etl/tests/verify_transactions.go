// Will verify event logs for all transactions which have num logs > 0
//	(the matching is done by comparing RLP encoded string of each types.Log object)
package main

import (
	"os"
	"os/signal"
	"syscall"
	"time"
	"strconv"
	"fmt"
	"context"
	"bytes"
	"log"
	"errors"
	"math/big"
	"encoding/hex"
	"encoding/json"

	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/rlp"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/rpc"
	"github.com/ethereum/go-ethereum/common/hexutil"

	. "github.com/PredictionExplorer/augur-explorer/primitives"
	. "github.com/PredictionExplorer/augur-explorer/dbs"
)
var (
	storage *SQLStorage

	RPC_URL = os.Getenv("AUGUR_ETH_NODE_RPC_URL")

	eclient *ethclient.Client
	rpcclient *rpc.Client

	market_order_id int64 = 0
	owner_fld_offset int64 = int64(OWNER_FIELD_OFFSET)	// offset to AugurContract::owner field obtained with eth_getStorage()

	Error   *log.Logger
	Info	*log.Logger
	Mismatch	*log.Logger

)
type rpcTransaction struct {
	tx *types.Transaction
	txExtraInfo
}
func (tx *rpcTransaction) UnmarshalJSON(msg []byte) error {
	if err := json.Unmarshal(msg, &tx.tx); err != nil {
		return err
	}
	return json.Unmarshal(msg, &tx.txExtraInfo)
}
type txExtraInfo struct {
	BlockNumber *string         `json:"blockNumber,omitempty"`
	BlockHash   *common.Hash    `json:"blockHash,omitempty"`
	From        *common.Address `json:"from,omitempty"`
	Hash		*common.Hash	`json:"hash,omitempty"`
}
type rpcBlock struct {
	Hash         common.Hash      `json:"hash"`
	Transactions []rpcTransaction `json:"transactions"`
	UncleHashes  []common.Hash    `json:"uncles"`
}
type receiptCallResult struct {
	receipt		*types.Receipt
	err			error
	idx			int
}
type rpcBlockHash struct {
	Hash		string
}
func get_block_hash(block_num int64) (string,error) {

	// this function is needed because Parity doesn't return the correct block hash over RPC, the hash
	// it returns is re-calculated while some fileds of the types.Header object are unset, giving wrong hash
	ctx := context.Background()
	var raw json.RawMessage
	err := rpcclient.CallContext(ctx, &raw,"eth_getBlockByNumber", hexutil.EncodeBig(big.NewInt(int64(block_num))),false)
	var blockHash rpcBlockHash
	err = json.Unmarshal(raw,&blockHash)
	if err!= nil {
		Error.Printf("Error unmarshalling hash of the block: %v\n",err)
		return "",err
	} else {
		return blockHash.Hash,nil
	}
}
func get_full_block(bnum int64) (common.Hash,*types.Header,[]*AugurTx,error) {

	var head *types.Header
	ctx := context.Background()
	var raw json.RawMessage
	err := rpcclient.CallContext(ctx, &raw,"eth_getBlockByNumber", hexutil.EncodeBig(big.NewInt(int64(bnum))),true)
	if err != nil {
		return common.Hash{},head,make([]*AugurTx,0,0),err
	}
	var body rpcBlock
	err = json.Unmarshal(raw, &body);
	if err != nil {
		Error.Printf("Error unmarshalling transactions of the block: %v\n",err)
		return common.Hash{},head, make([]*AugurTx,0,0),err
	}
	err = json.Unmarshal(raw,&head)
	if err!= nil {
		Error.Printf("Error unmarshalling hash of the block: %v\n",err)
		return body.Hash,head,make([]*AugurTx,0,0),err
	}
	txs := make([]*AugurTx, len(body.Transactions))
	for i, tx := range body.Transactions {
		/*if tx.From != nil {
			setSenderFromServer(tx.tx, *tx.From, body.Hash)
		}*/
		agtx := new(AugurTx)
		agtx.BlockNum = bnum
		agtx.TxHash = tx.txExtraInfo.Hash.String()
		agtx.From = tx.txExtraInfo.From.String()
		if tx.tx.To() != nil {
			agtx.To  = tx.tx.To().String()
		} else {
			agtx.CtrctCreate = true
			agtx.To = "0x0000000000000000000000000000000000000000"
		}
		agtx.Value = tx.tx.Value().String()
		agtx.Input = tx.tx.Data()
		agtx.GasPrice = tx.tx.GasPrice().String()
		txs[i]=agtx
		//txs[i] = tx.tx
	}

	return body.Hash,head,txs,nil
}
func verify_block(bnum int64) error {

	ctx := context.Background()
	block_hash_str,err:=get_block_hash(bnum)
	if err!=nil {
		return err
	}
	block_hash,header,transactions,err := get_full_block(bnum)
	if err!=nil {
		Info.Printf("Can't decode Block object received on RPC: %v. Aborting.\n",err)
		return err
	}
	num_transactions := len(transactions)
	Info.Printf("block %v hash = %v, num_tx=%v\n",bnum,block_hash_str,num_transactions)
	if bnum!=header.Number.Int64() {
		Info.Printf("Retrieved block number %v but Block object contains another number (%v)",bnum,header.Number.Int64())
		Error.Printf("Retrieved block number %v but Block object contains another number (%v)",bnum,header.Number.Int64())
		return errors.New("Block object inconsistency")
	}
	for txidx,tx := range transactions {
		tx_hash := common.HexToHash(tx.TxHash)
		rcpt,err := eclient.TransactionReceipt(ctx,tx_hash)
		if err != nil {
			Error.Printf(
				"Can't get receipt at block %v for tx (%v) %v , getReceipt err: %v\n",
				bnum,txidx,tx.TxHash,err,
			)
			os.Exit(1)
		}
		if rcpt.Status == types.ReceiptStatusFailed {
			continue	// transaction failed (i.e. Out of Gas, etc)
		}
		num_logs := len(rcpt.Logs)
		if num_logs > 0 {
			stored_logs,err := storage.Get_all_event_logs_by_tx_hash_rlp_encoded(tx.TxHash)
			if err!=nil {
				Error.Printf("At block %v ,can't get logs for tx %v from DB: %v\n",bnum,tx.TxHash)
				os.Exit(1)
			}
			if len(stored_logs)!=num_logs {
				Mismatch.Printf(
					"At block %v, tx %v number of logs doesnt match (should=%v, stored=%v)\n",
					num_logs,len(stored_logs),
				)
				continue
			} else {
				for i:=0 ; i < num_logs ; i++ {
					lg := rcpt.Logs[i]
					encoded_rlp_bytes, err := rlp.EncodeToBytes(lg)
					if err != nil {
						Info.Printf("Couldn't RLP-encode log : %v\n",err)
						os.Exit(1)
					}
					if !bytes.Equal(encoded_rlp_bytes,stored_logs[i]) {
						Mismatch.Printf(
							"At block %v for tx %v log #%v does not match (should %v, stored %v)\n",
							bnum,tx.TxHash,i,
							hex.EncodeToString(encoded_rlp_bytes),hex.EncodeToString(stored_logs[i]),
						)
						continue
					} else {
						//Info.Printf("block %v, tx %v (%v), log %v: match!\n",bnum,txidx,tx.TxHash,i)
					}
				}
			}
		}
	}
	Info.Printf("block_verif: %v %v ; %v transactions\n",bnum,block_hash.String(),num_transactions)
	return nil
}

func main() {
	//client, err := ethclient.Dial("http://:::8545")

	if len(RPC_URL) == 0 {
		Fatalf("Configuration error: RPC URL of Ethereum node is not set."+
				" Please set AUGUR_ETH_NODE_RPC environment variable")
	}
	var starting_block_num int64 = 0;
	if len(os.Args) > 2 {
		fmt.Printf("usage: %v [starting_block_num]\n",os.Args[0])
		os.Exit(1)
	} else {
		var er error
		if len(os.Args)==2 {
			starting_block_num,er = strconv.ParseInt(os.Args[1],10,64)
			if er!=nil {
				fmt.Printf("Error parsing starting block num: %v\n",er)
				os.Exit(1)
			}
		}
	}

	log_dir:=fmt.Sprintf("%v/%v",os.Getenv("HOME"),DEFAULT_LOG_DIR)
	os.MkdirAll(log_dir, os.ModePerm)
	db_log_file:=fmt.Sprintf("%v/l1_verify_db.log",log_dir)

	fname:=fmt.Sprintf("%v/l1_verify_info.log",log_dir)
	logfile, err := os.OpenFile(fname, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err!=nil {
		fmt.Printf("Can't start: %v\n",err)
		os.Exit(1)
	}
	Info = log.New(logfile,"INFO: ",log.Ltime|log.Lshortfile)

	fname=fmt.Sprintf("%v/l1_verify_error.log",log_dir)
	if err!=nil {
		fmt.Printf("Can't start: %v\n",err)
		os.Exit(1)
	}
	logfile, err = os.OpenFile(fname, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	Error = log.New(logfile,"ERROR: ",log.Ltime|log.Lshortfile)

	fname=fmt.Sprintf("%v/l1_verify_mismatch.log",log_dir)
	if err!=nil {
		fmt.Printf("Can't start: %v\n",err)
		os.Exit(1)
	}
	logfile, err = os.OpenFile(fname, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	Mismatch = log.New(logfile,"ERROR: ",log.Ltime|log.Lshortfile)

	rpcclient, err=rpc.DialContext(context.Background(), RPC_URL)
	if err != nil {
		log.Fatal(err)
	}
	Info.Printf("Connected to ETH node: %v\n",RPC_URL)
	eclient = ethclient.NewClient(rpcclient)

	storage = Connect_to_storage(&market_order_id,Info)
	storage.Init_log(db_log_file)
	storage.Log_msg("Log initialized\n")

	ctx := context.Background()
	stored_chain_id := storage.Get_stored_chain_id()
	network_chain_id,err :=eclient.NetworkID(ctx)
	if err != nil {
		Fatalf("Can't get Network ID: %v\n",err)
	}
	if stored_chain_id != network_chain_id.Int64() {
		if stored_chain_id == 0 {
			// not initialized yet
			Fatalf("DB not initialized chain id = 0")
		} else {
			Fatalf(
				"Network chain_id = %v , my chain_id = %v. Mismatch, exiting",
				network_chain_id.Int64(),stored_chain_id,
			)
		}
	}

	c := make(chan os.Signal)
	exit_chan := make(chan bool)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-c
		Info.Printf("Got SIGINT signal, will exit after block processing is over." +
					" To interrupt abruptly send SIGKILL (9) to the kernel.\n")
		exit_chan <- true
	}()

	latestBlock, err := eclient.BlockByNumber(ctx, nil)
	if err != nil {
		log.Fatal("oops:", err)
	}

	bnum,exists := storage.Get_last_block_num()
	if !exists {
		bnum = 0
	} else {
		bnum = bnum + 1
	}
	var bnum_high int64 = latestBlock.Number().Int64()
	if bnum_high < bnum {
		Info.Printf("Database has more blocks than the blockchain, aborting. Fix last_block table.\n")
		os.Exit(1)
	}
	if starting_block_num == 0 {
		starting_block_num = storage.Get_first_block_num()
	}
	bnum=starting_block_num
	Info.Printf("Integrity check started from block %v\n",starting_block_num)

	latestBlock, err = eclient.BlockByNumber(ctx, nil)
	if err != nil {
		log.Fatal("oops:", err)
	}

	bnum_high = latestBlock.Number().Int64()
	Info.Printf("Integrity check will stop at block %v\n",bnum_high)
	for ; bnum<bnum_high; bnum++ {
		select {
			case exit_flag := <-exit_chan:
				if exit_flag {
					Info.Printf("Exiting by user request at block %v.",bnum)
					os.Exit(0)
				}
			default:
		}
		err := verify_block(bnum)
		if err!=nil {
			// this is probably happening due to RPC unavailability, so we use a delay
			time.Sleep(1 * time.Second)
			Error.Printf("Block processing error: %v\n",err)
			break
		}
	}// for block_num
}
