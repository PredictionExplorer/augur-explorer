package main

import (
	"os"
	//"fmt"
	"bytes"
	"math/big"
	//"encoding/hex"

	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/accounts/abi"

	. "github.com/PredictionExplorer/augur-explorer/primitives"
	. "github.com/PredictionExplorer/augur-explorer/primitives/uniswapv3"
	. "github.com/PredictionExplorer/augur-explorer/dbs"
	. "github.com/PredictionExplorer/augur-explorer/contracts"
	. "github.com/PredictionExplorer/augur-explorer/layer1"
)
type ETL_Processor struct {
	ETL				*ETL_Layer1
}
func address_array_to_string(addresses []common.Address) string {

	var output string
	for i:=0; i<len(addresses); i++ {
		if len(output) > 0 { output = output + "," }
		output = output + addresses[i].String()
	}
	return output
}
func (this *ETL_Processor) Process_transaction(tx *AugurTx,rcpt *types.Receipt) {


	logs := rcpt.Logs
	for i:=0; i<len(logs); i++ {
		log := logs[i]
		process_event_log(this.ETL.Storage,tx,log,i)
	}
}
func process_pool_created(storage *SQLStorage,tx *AugurTx,log *types.Log,log_index int) {

	Info.Printf(
		"EVENT: PoolCreated. Tx %v TxIndex %v Log %v\n",
		tx.TxHash,tx.TxIndex,log.Index,
	)
	storagew.Delete_pool_created(tx.BlockNum,tx.TxIndex)
	if !bytes.Equal(log.Address.Bytes(),caddrs.FactoryAddr.Bytes()) {
		Info.Printf("Skipping event, address doesn't match our address\n")
		return
	}

	var eth_evt UniswapV3FactoryPoolCreated 
	err := factory_abi.UnpackIntoInterface(&eth_evt,"PoolCreated",log.Data)
	if err != nil {
		Error.Printf("Can't UnpackIntoInterface for PoolCreated: %v\n",err)
		os.Exit(1)
	}
	var evt UniV3PoolCreated
	evt.BlockNum = tx.BlockNum
	evt.TimeStamp = tx.TimeStamp
	evt.TxIndex = int64(tx.TxIndex)
	evt.LogIndex = int64(log.Index)
	evt.ContractAddr = log.Address.String()

	token0_addr := common.BytesToAddress(log.Topics[1][12:])
	token1_addr := common.BytesToAddress(log.Topics[2][12:])
	evt.Token0Addr = token0_addr.String()
	evt.Token1Addr = token1_addr.String()
	evt.Fee = log.Topics[3].Big().String()
	evt.TickSpacing = eth_evt.TickSpacing.String()
	evt.PoolAddr = eth_evt.Pool.String()

	Info.Printf("PoolCreated{\n")
	Info.Printf("\tBlockNum: %v\n",evt.BlockNum)
	Info.Printf("\tTimeStamp: %v\n",evt.TimeStamp)
	Info.Printf("\tTxId: %v\n",evt.TxIndex)
	Info.Printf("\tLogIndex: %v\n",evt.LogIndex)
	Info.Printf("\tContractAddr: %v\n",evt.ContractAddr)
	Info.Printf("\tPoolAddr: %v\n",evt.PoolAddr)
	Info.Printf("\tToken0: %v\n",evt.Token0Addr)
	Info.Printf("\tToken1: %v\n",evt.Token1Addr)
	Info.Printf("\tFee: %v\n",evt.Fee)
	Info.Printf("\tTickSpacing: %v\n",evt.TickSpacing)
	Info.Printf("}\n")
	storagew.Insert_pool_created(&evt)
}
func process_initialize(storage *SQLStorage,tx *AugurTx,log *types.Log,log_index int) {

	Info.Printf(
		"EVENT: Initialize. Tx %v TxIndex %v Log %v\n",
		tx.TxHash,tx.TxIndex,log.Index,
	)
	storagew.Delete_pool_initialize(tx.BlockNum,tx.TxIndex)
	pool_aid := storagew.Get_uniswap_v3_pool_aid(log.Address.String())
	if pool_aid == 0 {
		Info.Printf("Skipping event, address doesn't match our address\n")
		return
	}

	var eth_evt UniswapV3PoolInitialize
	err := pool_abi.UnpackIntoInterface(&eth_evt,"Initialize",log.Data)
	if err != nil {
		Error.Printf("Can't UnpackIntoInterface for Initialize: %v\n",err)
		os.Exit(1)
	}
	var evt UniV3Initialize
	evt.BlockNum = tx.BlockNum
	evt.TimeStamp = tx.TimeStamp
	evt.TxIndex = int64(tx.TxIndex)
	evt.LogIndex = int64(log.Index)
	evt.ContractAddr = log.Address.String()
	evt.PoolAid = pool_aid

	evt.SqrtPriceX96= eth_evt.SqrtPriceX96.String()
	evt.Tick = eth_evt.Tick.String()

	Info.Printf("Initialize{\n")
	Info.Printf("\tBlockNum: %v\n",evt.BlockNum)
	Info.Printf("\tTimeStamp: %v\n",evt.TimeStamp)
	Info.Printf("\tTxId: %v\n",evt.TxIndex)
	Info.Printf("\tLogIndex: %v\n",evt.LogIndex)
	Info.Printf("\tContractAddr: %v\n",evt.ContractAddr)
	Info.Printf("\tSqrtPriceX96: %v\n",evt.SqrtPriceX96)
	Info.Printf("\tTick: %v\n",evt.Tick)
	Info.Printf("}\n")
	storagew.Insert_initialize(&evt)
}
func process_pool_mint(storage *SQLStorage,tx *AugurTx,log *types.Log,log_index int) {

	Info.Printf(
		"EVENT: Mint. Tx %v TxIndex %v Log %v\n",
		tx.TxHash,tx.TxIndex,log.Index,
	)
	storagew.Delete_pool_mint(tx.BlockNum,tx.TxIndex)
	pool_aid := storagew.Get_uniswap_v3_pool_aid(log.Address.String())
	if pool_aid == 0 {
		Info.Printf("Skipping event, address doesn't match our address\n")
		return
	}
	var eth_evt UniswapV3PoolMint
	err := pool_abi.UnpackIntoInterface(&eth_evt,"Mint",log.Data)
	if err != nil {
		Error.Printf("Can't UnpackIntoInterface for Mint: %v\n",err)
		os.Exit(1)
	}

	var evt UniV3Mint
	evt.BlockNum = tx.BlockNum
	evt.TimeStamp = tx.TimeStamp
	evt.TxIndex = int64(tx.TxIndex)
	evt.LogIndex = int64(log_index)
	evt.ContractAddr = log.Address.String()
	evt.PoolAid = pool_aid

	owner_addr := common.BytesToAddress(log.Topics[1][12:])
	evt.OwnerAddr= owner_addr.String()
	var t abi.Type
	t.T = abi.IntTy
	t.Size = 256
	lower := abi.ReadInteger(t,log.Topics[2][:]).(*big.Int)
	upper := abi.ReadInteger(t,log.Topics[3][:]).(*big.Int)
	evt.TickLower = lower.String()
	evt.TickUpper = upper.String()
	evt.SenderAddr = eth_evt.Sender.String()
	evt.Amount = eth_evt.Amount.String()
	evt.Amount0 = eth_evt.Amount0.String()
	evt.Amount1 = eth_evt.Amount1.String()

	Info.Printf("Mint{\n")
	Info.Printf("\tBlockNum: %v\n",evt.BlockNum)
	Info.Printf("\tTimeStamp: %v\n",evt.TimeStamp)
	Info.Printf("\tTxId: %v\n",evt.TxIndex)
	Info.Printf("\tLogIndex: %v\n",evt.LogIndex)
	Info.Printf("\tContractAddr: %v\n",evt.ContractAddr)
	Info.Printf("\tSender: %v\n",evt.SenderAddr)
	Info.Printf("\tOwner: %v\n",evt.OwnerAddr)
	Info.Printf("\tTickLower: %v\n",evt.TickLower)
	Info.Printf("\tTickUpper: %v\n",evt.TickUpper)
	Info.Printf("\tAmount: %v\n",evt.Amount)
	Info.Printf("\tAmount0: %v\n",evt.Amount0)
	Info.Printf("\tAmount1: %v\n",evt.Amount1)
	Info.Printf("}\n")
	storagew.Insert_pool_mint(&evt)
}
func process_pool_collect(storage *SQLStorage,tx *AugurTx,log *types.Log,log_index int) {

	Info.Printf(
		"EVENT: Collect. Tx %v TxIndex %v Log %v\n",
		tx.TxHash,tx.TxIndex,log.Index,
	)
	pool_aid := storagew.Get_uniswap_v3_pool_aid(log.Address.String())
	if pool_aid == 0 {
		Info.Printf("Skipping event, address doesn't match our address\n")
		return
	}
	var eth_evt UniswapV3PoolCollect
	err := pool_abi.UnpackIntoInterface(&eth_evt,"Collect",log.Data)
	if err != nil {
		Error.Printf("Can't UnpackIntoInterface for Collect : %v\n",err)
		os.Exit(1)
	}

	var evt UniV3Collect
	evt.BlockNum = tx.BlockNum
	evt.TimeStamp = tx.TimeStamp
	evt.TxIndex = int64(tx.TxIndex)
	evt.LogIndex = int64(log_index)
	evt.ContractAddr = log.Address.String()
	evt.PoolAid = pool_aid

	owner_addr := common.BytesToAddress(log.Topics[1][12:])
	evt.OwnerAddr= owner_addr.String()
	evt.TickLower = log.Topics[2].Big().String()
	evt.TickUpper= log.Topics[3].Big().String()
	evt.RecipientAddr = eth_evt.Recipient.String()
	evt.Amount0 = eth_evt.Amount0.String()
	evt.Amount1 = eth_evt.Amount1.String()

	Info.Printf("Collect {\n")
	Info.Printf("\tBlockNum: %v\n",evt.BlockNum)
	Info.Printf("\tTimeStamp: %v\n",evt.TimeStamp)
	Info.Printf("\tTxId: %v\n",evt.TxIndex)
	Info.Printf("\tLogIndex: %v\n",evt.LogIndex)
	Info.Printf("\tContractAddr: %v\n",evt.ContractAddr)
	Info.Printf("\tOwner: %v\n",evt.OwnerAddr)
	Info.Printf("\tRecipient: %v\n",evt.RecipientAddr)
	Info.Printf("\tTickLower: %v\n",evt.TickLower)
	Info.Printf("\tTickUpper: %v\n",evt.TickUpper)
	Info.Printf("\tAmount0: %v\n",evt.Amount0)
	Info.Printf("\tAmount1: %v\n",evt.Amount1)
	Info.Printf("}\n")
	storagew.Insert_pool_collect(&evt)
}
func process_pool_burn(storage *SQLStorage,tx *AugurTx,log *types.Log,log_index int) {

	Info.Printf(
		"EVENT: Burn. Tx %v TxIndex %v Log %v\n",
		tx.TxHash,tx.TxIndex,log.Index,
	)
	storagew.Delete_pool_burn(tx.BlockNum,tx.TxIndex)
	pool_aid := storagew.Get_uniswap_v3_pool_aid(log.Address.String())
	if pool_aid == 0 {
		Info.Printf("Skipping event, address doesn't match our address\n")
		return
	}
	var eth_evt UniswapV3PoolBurn
	err := pool_abi.UnpackIntoInterface(&eth_evt,"Burn",log.Data)
	if err != nil {
		Error.Printf("Can't UnpackIntoInterface for Burn: %v\n",err)
		os.Exit(1)
	}

	var evt UniV3Burn
	evt.BlockNum = tx.BlockNum
	evt.TimeStamp = tx.TimeStamp
	evt.TxIndex = int64(tx.TxIndex)
	evt.LogIndex = int64(log_index)
	evt.ContractAddr = log.Address.String()
	evt.PoolAid = pool_aid

	owner_addr := common.BytesToAddress(log.Topics[1][12:])
	evt.OwnerAddr= owner_addr.String()
	evt.TickLower = log.Topics[2].Big().String()
	evt.TickUpper= log.Topics[3].Big().String()
	evt.Amount = eth_evt.Amount.String()
	evt.Amount0 = eth_evt.Amount0.String()
	evt.Amount1 = eth_evt.Amount1.String()

	Info.Printf("Collect {\n")
	Info.Printf("\tBlockNum: %v\n",evt.BlockNum)
	Info.Printf("\tTimeStamp: %v\n",evt.TimeStamp)
	Info.Printf("\tTxId: %v\n",evt.TxIndex)
	Info.Printf("\tLogIndex: %v\n",evt.LogIndex)
	Info.Printf("\tContractAddr: %v\n",evt.ContractAddr)
	Info.Printf("\tOwner: %v\n",evt.OwnerAddr)
	Info.Printf("\tTickLower: %v\n",evt.TickLower)
	Info.Printf("\tTickUpper: %v\n",evt.TickUpper)
	Info.Printf("\tAmount: %v\n",evt.Amount)
	Info.Printf("\tAmount0: %v\n",evt.Amount0)
	Info.Printf("\tAmount1: %v\n",evt.Amount1)
	Info.Printf("}\n")
	storagew.Insert_pool_burn(&evt)
}
func process_pool_swap(storage *SQLStorage,tx *AugurTx,log *types.Log,log_index int) {

	Info.Printf(
		"EVENT: Swap. Tx %v TxIndex %v Log %v\n",
		tx.TxHash,tx.TxIndex,log.Index,
	)
	storagew.Delete_pool_swap(tx.BlockNum,tx.TxIndex)
	pool_aid,fee_str := storagew.Get_uniswap_v3_pool_aid_and_fee(log.Address.String())
	if pool_aid == 0 {
		Info.Printf("Skipping event, address doesn't match our address\n")
		return
	}
	var eth_evt UniswapV3PoolSwap
	err := pool_abi.UnpackIntoInterface(&eth_evt,"Swap",log.Data)
	if err != nil {
		Error.Printf("Can't UnpackIntoInterface for Swap: %v\n",err)
		os.Exit(1)
	}

	var evt UniV3Swap
	evt.BlockNum = tx.BlockNum
	evt.TimeStamp = tx.TimeStamp
	evt.TxIndex = int64(tx.TxIndex)
	evt.LogIndex = int64(log_index)
	evt.ContractAddr = log.Address.String()
	evt.PoolAid = pool_aid

	sender_addr := common.BytesToAddress(log.Topics[1][12:])
	recipient_addr := common.BytesToAddress(log.Topics[2][12:])
	evt.SenderAddr= sender_addr.String()
	evt.RecipientAddr= recipient_addr.String()
	evt.Amount0 = eth_evt.Amount0.String()
	evt.Amount1 = eth_evt.Amount1.String()
	evt.SqrtPriceX96= eth_evt.SqrtPriceX96.String()
	evt.Liquidity=eth_evt.Liquidity.String()
	evt.Tick=eth_evt.Tick.String()
	a := big.NewInt(0)
	if eth_evt.Amount0.Sign() < 0 {
		a.Set(eth_evt.Amount0)
	}
	if eth_evt.Amount1.Sign() < 0 {
		a.Set(eth_evt.Amount1)
	}
	fee_big := big.NewInt(0)
	fee_big.SetString(fee_str,10)
	result_fee := big.NewInt(0)
	result_fee.Mul(a,fee_big)
	result_fee.Quo(a,big.NewInt(1e6))
	result_fee.Abs(result_fee)	// clear negative sign
	evt.Fee = result_fee.String()

	Info.Printf("Swap {\n")
	Info.Printf("\tBlockNum: %v\n",evt.BlockNum)
	Info.Printf("\tTimeStamp: %v\n",evt.TimeStamp)
	Info.Printf("\tTxId: %v\n",evt.TxIndex)
	Info.Printf("\tLogIndex: %v\n",evt.LogIndex)
	Info.Printf("\tContractAddr: %v\n",evt.ContractAddr)
	Info.Printf("\tSender: %v\n",evt.SenderAddr)
	Info.Printf("\tRecipient: %v\n",evt.RecipientAddr)
	Info.Printf("\tAmount0: %v\n",evt.Amount0)
	Info.Printf("\tAmount1: %v\n",evt.Amount1)
	Info.Printf("\tFee: %v\n",evt.Fee)
	Info.Printf("\tSqrtPriceX96: %v\n",evt.SqrtPriceX96)
	Info.Printf("\tLiquidity: %v\n",evt.Liquidity)
	Info.Printf("\tTick: %v\n",evt.Tick)
	Info.Printf("}\n")
	storagew.Insert_pool_swap(&evt)
}
func process_pool_flash(storage *SQLStorage,tx *AugurTx,log *types.Log,log_index int) {

	Info.Printf(
		"EVENT: Flash. Tx %v TxIndex %v Log %v\n",
		tx.TxHash,tx.TxIndex,log.Index,
	)
	pool_aid := storagew.Get_uniswap_v3_pool_aid(log.Address.String())
	if pool_aid == 0 {
		Info.Printf("Skipping event, address doesn't match our address\n")
		return
	}
	var eth_evt UniswapV3PoolFlash
	err := pool_abi.UnpackIntoInterface(&eth_evt,"Flash",log.Data)
	if err != nil {
		Error.Printf("Can't UnpackIntoInterface for Flash: %v\n",err)
		os.Exit(1)
	}

	var evt UniV3Flash
	evt.BlockNum = tx.BlockNum
	evt.TimeStamp = tx.TimeStamp
	evt.TxIndex = int64(tx.TxIndex)
	evt.LogIndex = int64(log_index)
	evt.ContractAddr = log.Address.String()
	evt.PoolAid = pool_aid

	sender_addr := common.BytesToAddress(log.Topics[1][12:])
	recipient_addr := common.BytesToAddress(log.Topics[2][12:])
	evt.SenderAddr= sender_addr.String()
	evt.RecipientAddr= recipient_addr.String()
	evt.Amount0 = eth_evt.Amount0.String()
	evt.Amount1 = eth_evt.Amount1.String()
	evt.Paid0= eth_evt.Paid0.String()
	evt.Paid1=eth_evt.Paid1.String()

	Info.Printf("Flash{\n")
	Info.Printf("\tBlockNum: %v\n",evt.BlockNum)
	Info.Printf("\tTimeStamp: %v\n",evt.TimeStamp)
	Info.Printf("\tTxId: %v\n",evt.TxIndex)
	Info.Printf("\tLogIndex: %v\n",evt.LogIndex)
	Info.Printf("\tContractAddr: %v\n",evt.ContractAddr)
	Info.Printf("\tSender: %v\n",evt.SenderAddr)
	Info.Printf("\tRecipient: %v\n",evt.RecipientAddr)
	Info.Printf("\tAmount0: %v\n",evt.Amount0)
	Info.Printf("\tAmount1: %v\n",evt.Amount1)
	Info.Printf("\tPaid0: %v\n",evt.Paid0)
	Info.Printf("\tPaid1: %v\n",evt.Paid1)
	Info.Printf("}\n")
	storagew.Insert_pool_flash(&evt)
}
func process_pool_increase_observation_cardinality_next(storage *SQLStorage,tx *AugurTx,log *types.Log,log_index int) {

	Info.Printf(
		"EVENT: IncreaseObservationCardinalityNext. Tx %v TxIndex %v Log %v\n",
		tx.TxHash,tx.TxIndex,log.Index,
	)
	pool_aid := storagew.Get_uniswap_v3_pool_aid(log.Address.String())
	if pool_aid == 0 {
		Info.Printf("Skipping event, address doesn't match our address\n")
		return
	}
	var eth_evt UniswapV3PoolIncreaseObservationCardinalityNext
	err := pool_abi.UnpackIntoInterface(&eth_evt,"IncreaseObservationCardinalityNext",log.Data)
	if err != nil {
		Error.Printf("Can't UnpackIntoInterface for UniswapV3PoolIncreaseObservationCardinalityNext: %v\n",err)
		os.Exit(1)
	}

	var evt UniV3IncObservCardinNext
	evt.BlockNum = tx.BlockNum
	evt.TimeStamp = tx.TimeStamp
	evt.TxIndex = int64(tx.TxIndex)
	evt.LogIndex = int64(log_index)
	evt.ContractAddr = log.Address.String()
	evt.PoolAid = pool_aid

	evt.ObservationCardinalityNextOld=eth_evt.ObservationCardinalityNextOld
	evt.ObservationCardinalityNextNew=eth_evt.ObservationCardinalityNextNew


	Info.Printf("Flash{\n")
	Info.Printf("\tBlockNum: %v\n",evt.BlockNum)
	Info.Printf("\tTimeStamp: %v\n",evt.TimeStamp)
	Info.Printf("\tTxId: %v\n",evt.TxIndex)
	Info.Printf("\tLogIndex: %v\n",evt.LogIndex)
	Info.Printf("\tContractAddr: %v\n",evt.ContractAddr)
	Info.Printf("\tObservationCardinalityNextOld: %v\n",evt.ObservationCardinalityNextOld)
	Info.Printf("\tObservationCardinalityNextNew: %v\n",evt.ObservationCardinalityNextNew)
	Info.Printf("}\n")
	storagew.Insert_pool_increase_observation_cardinality_next(&evt)
}
func process_pool_set_fee_protocol(storage *SQLStorage,tx *AugurTx,log *types.Log,log_index int) {

	Info.Printf(
		"EVENT: SetFeeProtocol. Tx %v TxIndex %v Log %v\n",
		tx.TxHash,tx.TxIndex,log.Index,
	)
	pool_aid := storagew.Get_uniswap_v3_pool_aid(log.Address.String())
	if pool_aid == 0 {
		Info.Printf("Skipping event, address doesn't match our address\n")
		return
	}
	var eth_evt UniswapV3PoolSetFeeProtocol
	err := pool_abi.UnpackIntoInterface(&eth_evt,"SetFeeProtocol",log.Data)
	if err != nil {
		Error.Printf("Can't UnpackIntoInterface for SetFeeProtocol: %v\n",err)
		os.Exit(1)
	}

	var evt UniV3SetFeeProtocol
	evt.BlockNum = tx.BlockNum
	evt.TimeStamp = tx.TimeStamp
	evt.TxIndex = int64(tx.TxIndex)
	evt.LogIndex = int64(log_index)
	evt.ContractAddr = log.Address.String()
	evt.PoolAid = pool_aid

	evt.FeeProtocol0Old = eth_evt.FeeProtocol0Old
	evt.FeeProtocol1Old = eth_evt.FeeProtocol1Old
	evt.FeeProtocol0New = eth_evt.FeeProtocol0New
	evt.FeeProtocol1New = eth_evt.FeeProtocol1New

	Info.Printf("SetFeeProtocol{\n")
	Info.Printf("\tBlockNum: %v\n",evt.BlockNum)
	Info.Printf("\tTimeStamp: %v\n",evt.TimeStamp)
	Info.Printf("\tTxId: %v\n",evt.TxIndex)
	Info.Printf("\tLogIndex: %v\n",evt.LogIndex)
	Info.Printf("\tContractAddr: %v\n",evt.ContractAddr)
	Info.Printf("\tFeeProtocol0Old: %v\n",evt.FeeProtocol0Old)
	Info.Printf("\tFeeProtocol1Old: %v\n",evt.FeeProtocol1Old)
	Info.Printf("\tFeeProtocol0New: %v\n",evt.FeeProtocol0New)
	Info.Printf("\tFeeProtocol1New: %v\n",evt.FeeProtocol1New)
	Info.Printf("}\n")
	storagew.Insert_pool_set_fee_protocol(&evt)
}
func process_pool_collect_protocol(storage *SQLStorage,tx *AugurTx,log *types.Log,log_index int) {

	Info.Printf(
		"EVENT: CollectProtocol. Tx %v TxIndex %v Log %v\n",
		tx.TxHash,tx.TxIndex,log.Index,
	)
	pool_aid := storagew.Get_uniswap_v3_pool_aid(log.Address.String())
	if pool_aid == 0 {
		Info.Printf("Skipping event, address doesn't match our address\n")
		return
	}
	var eth_evt UniswapV3PoolCollectProtocol
	err := pool_abi.UnpackIntoInterface(&eth_evt,"CollectProtocol",log.Data)
	if err != nil {
		Error.Printf("Can't UnpackIntoInterface for CollectProtcol: %v\n",err)
		os.Exit(1)
	}

	var evt UniV3PoolCollectProtocol
	evt.BlockNum = tx.BlockNum
	evt.TimeStamp = tx.TimeStamp
	evt.TxIndex = int64(tx.TxIndex)
	evt.LogIndex = int64(log_index)
	evt.ContractAddr = log.Address.String()
	evt.PoolAid = pool_aid

	sender_addr := common.BytesToAddress(log.Topics[1][12:])
	recipient_addr := common.BytesToAddress(log.Topics[2][12:])
	evt.SenderAddr = sender_addr.String()
	evt.RecipientAddr = recipient_addr.String()
	evt.Amount0 = eth_evt.Amount0.String()
	evt.Amount1 = eth_evt.Amount1.String()

	Info.Printf("CollectProtocol{\n")
	Info.Printf("\tBlockNum: %v\n",evt.BlockNum)
	Info.Printf("\tTimeStamp: %v\n",evt.TimeStamp)
	Info.Printf("\tTxId: %v\n",evt.TxIndex)
	Info.Printf("\tLogIndex: %v\n",evt.LogIndex)
	Info.Printf("\tContractAddr: %v\n",evt.ContractAddr)
	Info.Printf("\tSender: %v\n",evt.SenderAddr)
	Info.Printf("\tRecipient: %v\n",evt.RecipientAddr)
	Info.Printf("\tAmount0: %v\n",evt.Amount0)
	Info.Printf("\tAmount1: %v\n",evt.Amount1)
	Info.Printf("}\n")
	storagew.Insert_pool_collect_protocol(&evt)
}
func process_collect_periphery(storage *SQLStorage,tx *AugurTx,log *types.Log,log_index int) {

	Info.Printf(
		"EVENT: Collect (Periphery). Tx %v TxIndex %v Log %v\n",
		tx.TxHash,tx.TxIndex,log.Index,
	)
	if !bytes.Equal(caddrs.NFTPosMgrAddr.Bytes(),log.Address.Bytes()) {
		Info.Printf("Skipping event, address doesn't match our address (%v)\n",caddrs.NFTPosMgrAddr.String())
		return
	}

	var eth_evt NonfungiblePositionManagerCollect
	err := nfpm_abi.UnpackIntoInterface(&eth_evt,"Collect",log.Data)
	if err != nil {
		Error.Printf("Can't UnpackIntoInterface for Collect (Periphery): %v\n",err)
		os.Exit(1)
	}

	var evt UniV3CollectPeriphery
	evt.BlockNum = tx.BlockNum
	evt.TimeStamp = tx.TimeStamp
	evt.TxIndex = int64(tx.TxIndex)
	evt.LogIndex = int64(log_index)
	evt.ContractAddr = log.Address.String()

	token_id := log.Topics[1].Big().String()
	evt.TokenId = token_id
	evt.RecipientAddr = eth_evt.Recipient.String()
	evt.Amount0 = eth_evt.Amount0.String()
	evt.Amount1 = eth_evt.Amount1.String()
	Info.Printf("Collect (Periphery) {\n")
	Info.Printf("\tBlockNum: %v\n",evt.BlockNum)
	Info.Printf("\tTimeStamp: %v\n",evt.TimeStamp)
	Info.Printf("\tTxId: %v\n",evt.TxIndex)
	Info.Printf("\tLogIndex: %v\n",evt.LogIndex)
	Info.Printf("\tContractAddr: %v\n",evt.ContractAddr)
	Info.Printf("\tTokenID: %v\n",evt.TokenId)
	Info.Printf("\tRecipient: %v\n",evt.RecipientAddr)
	Info.Printf("\tAmount0: %v\n",evt.Amount0)
	Info.Printf("\tAmount1: %v\n",evt.Amount1)
	Info.Printf("}\n")
	storagew.Insert_collect_periphery(&evt)
}
func process_increase_liquidity(storage *SQLStorage,tx *AugurTx,log *types.Log,log_index int) {

	Info.Printf(
		"EVENT: IncreaseLiquidity. Tx %v TxIndex %v Log %v\n",
		tx.TxHash,tx.TxIndex,log.Index,
	)
	if !bytes.Equal(caddrs.NFTPosMgrAddr.Bytes(),log.Address.Bytes()) {
		Info.Printf("Skipping event, address doesn't match our address (%v)\n",caddrs.NFTPosMgrAddr.String())
		return
	}

	var eth_evt NonfungiblePositionManagerIncreaseLiquidity
	err := nfpm_abi.UnpackIntoInterface(&eth_evt,"IncreaseLiquidity",log.Data)
	if err != nil {
		Error.Printf("Can't UnpackIntoInterface for IncreaseLiquidity : %v\n",err)
		os.Exit(1)
	}

	var evt UniV3IncreaseLiquidity
	evt.BlockNum = tx.BlockNum
	evt.TimeStamp = tx.TimeStamp
	evt.TxIndex = int64(tx.TxIndex)
	evt.LogIndex = int64(log_index)
	evt.ContractAddr = log.Address.String()

	token_id := log.Topics[1].Big().String()
	evt.TokenId = token_id
	evt.Liquidity = eth_evt.Liquidity.String()
	evt.Amount0 = eth_evt.Amount0.String()
	evt.Amount1 = eth_evt.Amount1.String()

	Info.Printf("IncreaseLiquidity{\n")
	Info.Printf("\tBlockNum: %v\n",evt.BlockNum)
	Info.Printf("\tTimeStamp: %v\n",evt.TimeStamp)
	Info.Printf("\tTxId: %v\n",evt.TxIndex)
	Info.Printf("\tLogIndex: %v\n",evt.LogIndex)
	Info.Printf("\tContractAddr: %v\n",evt.ContractAddr)
	Info.Printf("\tTokenID: %v\n",evt.TokenId)
	Info.Printf("\tLiquidity: %v\n",evt.Liquidity)
	Info.Printf("\tAmount0: %v\n",evt.Amount0)
	Info.Printf("\tAmount1: %v\n",evt.Amount1)
	Info.Printf("}\n")
	storagew.Insert_increase_liquidity(&evt)
}
func process_decrease_liquidity(storage *SQLStorage,tx *AugurTx,log *types.Log,log_index int) {

	Info.Printf(
		"EVENT: DecreaseLiquidity. Tx %v TxIndex %v Log %v\n",
		tx.TxHash,tx.TxIndex,log.Index,
	)
	if !bytes.Equal(caddrs.NFTPosMgrAddr.Bytes(),log.Address.Bytes()) {
		Info.Printf("Skipping event, address doesn't match our address (%v)\n",caddrs.NFTPosMgrAddr.String())
		return
	}

	var eth_evt NonfungiblePositionManagerDecreaseLiquidity
	err := nfpm_abi.UnpackIntoInterface(&eth_evt,"DecreaseLiquidity",log.Data)
	if err != nil {
		Error.Printf("Can't UnpackIntoInterface for DecreaseLiquidity : %v\n",err)
		os.Exit(1)
	}

	var evt UniV3DecreaseLiquidity
	evt.BlockNum = tx.BlockNum
	evt.TimeStamp = tx.TimeStamp
	evt.TxIndex = int64(tx.TxIndex)
	evt.LogIndex = int64(log_index)
	evt.ContractAddr = log.Address.String()

	token_id := log.Topics[1].Big().String()
	evt.TokenId = token_id
	evt.Liquidity = eth_evt.Liquidity.String()
	evt.Amount0 = eth_evt.Amount0.String()
	evt.Amount1 = eth_evt.Amount1.String()

	Info.Printf("DecreaseLiquidity{\n")
	Info.Printf("\tBlockNum: %v\n",evt.BlockNum)
	Info.Printf("\tTimeStamp: %v\n",evt.TimeStamp)
	Info.Printf("\tTxId: %v\n",evt.TxIndex)
	Info.Printf("\tLogIndex: %v\n",evt.LogIndex)
	Info.Printf("\tContractAddr: %v\n",evt.ContractAddr)
	Info.Printf("\tTokenID: %v\n",evt.TokenId)
	Info.Printf("\tLiquidity: %v\n",evt.Liquidity)
	Info.Printf("\tAmount0: %v\n",evt.Amount0)
	Info.Printf("\tAmount1: %v\n",evt.Amount1)
	Info.Printf("}\n")
	storagew.Insert_decrease_liquidity(&evt)
}

/* Possible removal pending
func process_bpt_erc20_transfer(storage *SQLStorage,tx *AugurTx,log *types.Log,log_index int) {
	// Only processes BPT (Balancer Pool Token) events (the token indicating share of the funder)

	pool_aid := storagew.Is_balancer_pool_address(log.Address.String())
	if pool_aid == 0 {
		return	// not a Pool contract
	}
	Info.Printf(
		"EVENT: BPT ERC20 Transfer. Tx %v TxIndex %v Log %v\n",
		tx.TxHash,tx.TxIndex,log.Index,
	)
	var eth_evt ETransfer
	eth_evt.From= common.BytesToAddress(log.Topics[1][12:])
	eth_evt.To= common.BytesToAddress(log.Topics[2][12:])
	err := erc20_abi.UnpackIntoInterface(&eth_evt,"Transfer",log.Data)
	if err != nil {
		Error.Printf("Can't UnpackIntoInterface for BPT ERC20 Transfer: %v\n",err)
		os.Exit(1)
	}

	var evt BalV2BPTTransfer
	evt.BlockNum = tx.BlockNum
	evt.TimeStamp = tx.TimeStamp
	evt.TxIndex = int64(tx.TxIndex)
	evt.LogIndex = int64(log_index)
	evt.PoolAid = pool_aid
	evt.From = eth_evt.From.String()
	evt.To = eth_evt.To.String()
	evt.Amount = eth_evt.Value.String()

	Info.Printf("BPT ERC20 Transfer{\n")
	Info.Printf("\tBlockNum: %v\n",evt.BlockNum)
	Info.Printf("\tTimeStamp: %v\n",evt.TimeStamp)
	Info.Printf("\tTxId: %v\n",evt.TxIndex)
	Info.Printf("\tLogIndex: %v\n",evt.LogIndex)
	Info.Printf("\tContractAddr: %v\n",log.Address.String())
	Info.Printf("\tPoolAid: %v\n",evt.PoolAid)
	Info.Printf("\tFrom: %v\n",evt.From)
	Info.Printf("\tTo: %v\n",evt.To)
	Info.Printf("\tAmount: %v\n",evt.Amount)
	Info.Printf("}\n")
	storagew.Insert_bpt_erc20_transfer(&evt)
}
*/
func process_event_log(storage *SQLStorage,tx *AugurTx,log *types.Log,log_index int) {

	if len(log.Topics) == 0 { return }
	topic0 := log.Topics[0].Bytes()
	//Info.Printf("topic0=%v\n",hex.EncodeToString(topic0[:]))
	if bytes.Equal(topic0,evt_pool_created) {
		process_pool_created(storage,tx,log,log_index)
	}
	if bytes.Equal(topic0,evt_initialize) {
		process_initialize(storage,tx,log,log_index)
	}
	if bytes.Equal(topic0,evt_mint) {
		process_pool_mint(storage,tx,log,log_index)
	}
	if bytes.Equal(topic0,evt_collect) {
		process_pool_collect(storage,tx,log,log_index)
	}
	if bytes.Equal(topic0,evt_collect_periphery) {
		process_collect_periphery(storage,tx,log,log_index)
	}
	if bytes.Equal(topic0,evt_burn) {
		process_pool_burn(storage,tx,log,log_index)
	}
	if bytes.Equal(topic0,evt_swap) {
		process_pool_swap(storage,tx,log,log_index)
	}
	if bytes.Equal(topic0,evt_flash) {
		process_pool_flash(storage,tx,log,log_index)
	}
	if bytes.Equal(topic0,evt_inc_obs_cardin_next) {
		process_pool_increase_observation_cardinality_next(storage,tx,log,log_index)
	}
	if bytes.Equal(topic0,evt_set_fee_protocol) {
		process_pool_set_fee_protocol(storage,tx,log,log_index)
	}
	if bytes.Equal(topic0,evt_increase_liquidity) {
		process_increase_liquidity(storage,tx,log,log_index)
	}
	if bytes.Equal(topic0,evt_decrease_liquidity) {
		process_decrease_liquidity(storage,tx,log,log_index)
	}
}
