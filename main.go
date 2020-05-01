// Augur ETL: Converts Augur Data to SQL database
package main

import (
	"os"
	"fmt"
	"context"
	"log"
	"math/big"
	"encoding/hex"

	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/accounts/abi"
)
const (
	// ToDo: get these signatures from the abi files (after our code stabilizes, currently we will
	//	leave these constants visible to aid debugging processes)
	MARKET_CREATED = "ea17ae24b0d40ea7962a6d832db46d1f81eaec1562946d0830d1c21d4c000ec1"
	MARKET_OI_CHANGED = "213a05b9ad8567c2f8fa868e7375e5bf30e69add0dbb5913ca8a3e58c815c268"
	MARKET_ORDER = "9bab1368a1ed530afaad9c630ba75e6a5c1efa9f6af0139d6cda2b6af6aa801e"
	MARKET_FINALIZED = "6d39632c2dc10305bf5771cfff4af1851f07c03ea27b821cad382466bdf7a21f"
	INITIAL_REPORT_SUBMITTED = "c3ebb227c22e7644e9bef8822009f746a72c86f239760124d67fdc2c302b3115"
	MARKET_VOLUME_CHANGED = "e9f0af820300e73bae76c8e76943abe7fbb4224b49cb133e2dadc6f71acf6370"
	DISPUTE_CROWDSOURCER_CONTRIBUTION = "e7f47639cdf56ec6c5451df334b73c9ca5cccd20da2c0f4e390e9bb71a6f672a"
	TOKENS_TRANSFERRED = "3c67396e9c55d2fc8ad68875fc5beca1d96ad2a2f23b210ccc1d986551ab6fdf"
	TOKEN_BALANCE_CHANGED = "63fd58f559b73fc4da5511c341ec8a7b31c5c48538ef83c6077712b6edf5f7cb"
	SHARE_TOKEN_BALANCE_CHANGED = "350ea32dc29530b9557420816d743c436f8397086f98c96292138edd69e01cb3"
	CANCEL_0X_ORDER = "be80e5687d7095071b7c4e7a56e0e67bfb9e8a39352f1690fdf74c1ee935c75e"
	TRANSFER_BATCH = "4a39dc06d4c0dbc64b70af90fd698a233a518aa5d07e595d983b8c0526c8f7fb"
	TRANSFER_SINGLE = "c3d58168c5ae7397731d063d5bbf3d657854427343f4c083240f7aacaa2d0f62"
	PROFIT_LOSS_CHANGED = "59543b7f82735782aa5bdb97dff40ff288d4548a5865da513b40e4088e2ee77e"
	ERC20_TRANSFER = "ddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef"
	EXCHANGE_FILL = "6869791f0a34781b29882982cc39e882768cf2c96995c2a110c577c53bc932d5"
	TRADING_PROCEEDS_CLAIMED = "95366b7f64c6bb45149f9f7c522403fceebe5170ff76b8ffde2b0ab943ac11ce"
	ZEROX_APPROVAL_FOR_ALL = "17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31"
	ERC20_APPROVAL = "8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925"

	MAX_BLOCKS_CHAIN_SPLIT = 128
)
var (
	// these evt_ variables are here for speed to avoid calculation of Keccak256
	//		on each bytes.Compare() operation
	evt_market_created,_ = hex.DecodeString(MARKET_CREATED)
	evt_market_oi_changed,_ = hex.DecodeString(MARKET_OI_CHANGED)
	evt_market_order,_ = hex.DecodeString(MARKET_ORDER)
	evt_market_finalized,_ = hex.DecodeString(MARKET_FINALIZED)
	evt_initial_report_submitted,_ = hex.DecodeString(INITIAL_REPORT_SUBMITTED)
	evt_market_volume_changed,_ = hex.DecodeString(MARKET_VOLUME_CHANGED)
	evt_dispute_crowd_contrib,_ = hex.DecodeString(DISPUTE_CROWDSOURCER_CONTRIBUTION)
	evt_tokens_transferred,_ = hex.DecodeString(TOKENS_TRANSFERRED)
	evt_token_balance_changed,_ = hex.DecodeString(TOKEN_BALANCE_CHANGED)
	evt_share_token_balance_changed,_ = hex.DecodeString(SHARE_TOKEN_BALANCE_CHANGED)
	evt_cancel_0x_order,_ = hex.DecodeString(CANCEL_0X_ORDER)
	evt_transfer_batch,_ = hex.DecodeString(TRANSFER_BATCH)
	evt_transfer_single,_ = hex.DecodeString(TRANSFER_SINGLE)
	evt_profit_loss_changed,_ = hex.DecodeString(PROFIT_LOSS_CHANGED)
	evt_erc20_transfer,_ = hex.DecodeString(ERC20_TRANSFER)
	evt_exchange_fill,_ = hex.DecodeString(EXCHANGE_FILL)
	evt_trading_proceeds_claimed,_ = hex.DecodeString(TRADING_PROCEEDS_CLAIMED)
	evt_zerox_approval_for_all,_ = hex.DecodeString(ZEROX_APPROVAL_FOR_ALL)
	evt_erc20_approval,_ = hex.DecodeString(ERC20_APPROVAL)

	storage *SQLStorage

	all_contracts map[string]interface{}
	inspected_events [][]byte

	augur_abi *abi.ABI
	trading_abi *abi.ABI
	zerox_abi *abi.ABI
	cash_abi *abi.ABI
	exchange_abi *abi.ABI
	wallet_abi *abi.ABI

	RPC_URL = os.Getenv("AUGUR_ETH_NODE_RPC_URL")
)
func main() {
	//client, err := ethclient.Dial("http://:::8545")

	if len(RPC_URL) == 0 {
		Fatalf("Configuration error: RPC URL of Ethereum node is not set."+
				" Please set AUGUR_ETH_NODE_RPC environment variable")
	}

	augur_init()
	//client, err := ethclient.Dial("http://192.168.1.102:18545")
	client, err := ethclient.Dial(RPC_URL)
	if err != nil {
		log.Fatal(err)
	}

	storage = connect_to_storage()



	ctx := context.Background()
	latestBlock, err := client.BlockByNumber(ctx, nil)
	if err != nil {
		log.Fatal("oops:", err)
	}
	log.Printf("latest block: %v\n", latestBlock.Number())

	bnum,exists := storage.get_last_block_num()
	if !exists {
		bnum = 0
	} else {
		bnum = bnum + 1
	}
	split_simulated := false
	fmt.Printf("Starting data load from block %v\n",bnum)
	var bnum_high BlockNumber = BlockNumber(latestBlock.Number().Uint64())
	for ; bnum<bnum_high; bnum++ {
		big_bnum:=big.NewInt(int64(bnum))
		block, _ := client.BlockByNumber(ctx,big_bnum)
		if block != nil {
			storage.block_delete_with_everything(BlockNumber(big_bnum.Int64()))
			num_transactions, err := client.TransactionCount(ctx,block.Hash())
			if err != nil {
				fmt.Printf("block error: %v \n",err)
			} else {
				header := block.Header()
				back_block_num := new(big.Int).SetUint64(header.Number.Uint64() - 20)
				if (back_block_num.Uint64() == 99999999999999) && !split_simulated {//simulation disabled
					// code to simulate chain split (naive) , this block should be removed when stable
					block,_ = client.BlockByNumber(ctx,back_block_num)
					header = block.Header()
					big_bnum = big.NewInt(int64(header.Number.Int64()))
					bnum = BlockNumber(big_bnum.Uint64())
					storage.block_delete_with_everything(BlockNumber(header.Number.Int64()))
					split_simulated = true
					fmt.Println("Chain split simulation in action");
				}
				if !storage.insert_block(header,int64(num_transactions)) {
					// chainsplit detected
					set_back_block_num := storage.fix_chainsplit(header)
					fmt.Printf("Chain rewind to block %v. Restarting.",set_back_block_num)
					bnum = set_back_block_num
					continue
				}
				if num_transactions > 0 {
					fmt.Printf("block: %v %v transactions\n",block.Number(),num_transactions)
					for tnum:=0 ; tnum < int(num_transactions) ; tnum++ {
						tx , err := client.TransactionInBlock(ctx,block.Hash(),uint(tnum))
						if err != nil {
							fmt.Printf("Error: %v",err)
						} else {
							tx_id := storage.insert_transaction(BlockNumber(block.Number().Uint64()),tx)
							tx_msg, err := tx.AsMessage(types.NewEIP155Signer(tx.ChainId()))
							if err != nil {
								Fatalf("Error in tx signature validation (shoudln't happen): %v",err)
							}
							from := tx_msg.From()
							dump_tx_input_if_known(tx)
							to:=""
							if tx.To() != nil {
								to = tx.To().String()
							}
							fmt.Printf("\ttx: %v\n",tx.Hash().String())
							fmt.Printf("\t from=%v\n",tx_msg.From().String())
							fmt.Printf("\t to=%v for $%v (%v bytes data)\n",
											to,tx.Value().String(),len(tx.Data()))
							fmt.Printf("\t input: \n%v\n",hex.EncodeToString(tx.Data()[:]))
							rcpt,err := client.TransactionReceipt(ctx,tx.Hash())
							if err != nil {
								fmt.Printf("Error: %v",err)
							} else {
								sequencer := new(EventSequencer)
								num_logs := len(rcpt.Logs)
								for i:=0 ; i<num_logs ; i++ {
									fmt.Printf(
										"\t\t\tlog %v\n\t\t\t\t\t\t for contract %v (%v of %v items)\n",
										rcpt.Logs[i].Topics[0].String(),rcpt.Logs[i].Address.String(),(i+1),len(rcpt.Logs))
									sequencer.append_event(rcpt.Logs[i])
								}
								ordered_list := sequencer.get_ordered_event_list()
								num_logs = len(ordered_list)
								for i:=0 ; i < num_logs ; i++ {
									fmt.Printf(
										"\t\t\tchecking log with sig %v\n\t\t\t\t\t\t for contract %v\n",
										ordered_list[i].Topics[0].String(),
										ordered_list[i].Address.String())
									process_event(block.Header(),tx_id,from,ordered_list[i])
								}
							}
						}
					}
				} else {
					fmt.Printf("block: %v EMPTY\n",block.Number())
				}
			}
		}
		storage.set_last_block_num(bnum)
	}// for block_num
	fmt.Printf("new latest block = %v\n",bnum_high)
}
