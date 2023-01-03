package main

import (
	"os"
	//"fmt"
	"bytes"
	"math/big"
	//"encoding/hex"

	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/vm"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/rlp"

	. "github.com/PredictionExplorer/augur-explorer/primitives"
	. "github.com/PredictionExplorer/augur-explorer/primitives/uniswapv3"
	. "github.com/PredictionExplorer/augur-explorer/dbs"
	. "github.com/PredictionExplorer/augur-explorer/contracts"
	. "github.com/PredictionExplorer/augur-explorer/layer1"
	"github.com/PredictionExplorer/augur-explorer/uevm"
)
type ETL_Processor struct {
	ETL				*ETL_Layer1
}
type IUniswapV3PoolEventsDBGSWAPLOOP struct {
	Pool                common.Address
	SqrtPriceX96        *big.Int
	SqrtPriceStartX96   *big.Int
	SqrtPriceNextX96    *big.Int
	Tick                *big.Int
	TickCumulative      *big.Int
	Initialized         bool
	StepAmountIn        *big.Int
	StepAmountOut       *big.Int
	AmountProcessed     *big.Int
	FeeAmount           *big.Int
	Liquidity           *big.Int
	ExactInput          bool
	FeeGrowthGlobalX128 *big.Int
	Raw                 types.Log // Blockchain specific contextual infos
}
type IUniswapV3PoolEventsDBGUPDPOS struct {
	Owner                      common.Address
	TickLower                  *big.Int
	TickUpper                  *big.Int
	Tick                       *big.Int
	LiquidityDelta             *big.Int
	FeeGrowthGlobal0X128Before *big.Int
	FeeGrowthGlobal1X128Before *big.Int
	FeeGrowthInside0X128       *big.Int
	FeeGrowthInside1X128       *big.Int
	FlippedLower               bool
	FlippedUpper               bool
	Raw                        types.Log // Blockchain specific contextual infos
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

	var rec uevm.Record
	rec.BlockNum = tx.BlockNum 
	rec.BlockHash = common.HexToHash(tx.BlockHash)
	rec.TxIndex = int64(tx.TxIndex)
	rec.TxHash = common.HexToHash(tx.TxHash)

	ctrct_addr := common.HexToAddress(tx.To)
	block_ctx := uevm.NewDummyBlockContext(big.NewInt(uevm.MainNetBlockNum) ,big.NewInt(uevm.MainNetTimeStamp))
	tx_ctx := new(vm.TxContext)
	tx_ctx.Origin = common.HexToAddress(tx.From) 
	tx_ctx.GasPrice = big.NewInt(uevm.TxDefaultGas)
	value := big.NewInt(0)
	value.SetString(tx.Value,10)
	input := tx.Input

	last_line_rec,err := mchain.ReadLastLine()
	if err != nil {
		Info.Printf("Error getting last record: %v\n",err)
		os.Exit(1)
	}

	err,_ = mchain.ExecCall(block_ctx,rec.TxHash,tx_ctx,input,value,ctrct_addr,last_line_rec.StateRoot,&rec)
	if err != nil {
		Info.Printf("Error in ExecCall(): %v\n",err)
		os.Exit(1)
	}
}
func process_initialize(storage *SQLStorage,tx *AugurTx,log *types.Log,log_index int) {

	Info.Printf(
		"EVENT: Initialize. Tx %v TxIndex %v Log %v\n",
	tx.TxHash,tx.TxIndex,log.Index,
	)
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

	var rec uevm.Record
	rec.BlockNum = tx.BlockNum 
	rec.BlockHash = common.HexToHash(tx.BlockHash)
	rec.TxIndex = int64(tx.TxIndex)
	rec.TxHash = common.HexToHash(tx.TxHash)

	ctrct_addr := common.HexToAddress(tx.To)
	block_ctx := uevm.NewDummyBlockContext(big.NewInt(uevm.MainNetBlockNum) ,big.NewInt(uevm.MainNetTimeStamp))
	tx_ctx := new(vm.TxContext)
	tx_ctx.Origin = common.HexToAddress(tx.From) 
	tx_ctx.GasPrice = big.NewInt(uevm.TxDefaultGas)
	value := big.NewInt(0)
	value.SetString(tx.Value,10)
	input := tx.Input

	last_line_rec,err := mchain.ReadLastLine()
	if err != nil {
		Info.Printf("Error getting last record: %v\n",err)
		os.Exit(1)
	}

	err,_ = mchain.ExecCall(block_ctx,rec.TxHash,tx_ctx,input,value,ctrct_addr,last_line_rec.StateRoot,&rec)
	if err != nil {
		Info.Printf("Error in ExecCall(): %v\n",err)
		os.Exit(1)
	}
}
func process_pool_mint(storage *SQLStorage,tx *AugurTx,log *types.Log,log_index int) {

	Info.Printf(
		"EVENT: Mint. Tx %v TxIndex %v Log %v\n",
		tx.TxHash,tx.TxIndex,log.Index,
	)
	pool_aid := storagew.Get_uniswap_v3_pool_aid(log.Address.String())
	if pool_aid == 0 {
		Info.Printf("Skipping event, address doesn't match our address\n")
		return
	}
	t0_addr_str,_,t1_addr_str,_ := storagew.Get_uniswap_v3_pool_token_addresses(pool_aid)
	var eth_evt UniswapV3PoolMint
	err := pool_abi.UnpackIntoInterface(&eth_evt,"Mint",log.Data)
	if err != nil {
		Error.Printf("Can't UnpackIntoInterface for Mint: %v\n",err)
		os.Exit(1)
	}

	t0_addr := common.HexToAddress(t0_addr_str)
	t1_addr := common.HexToAddress(t1_addr_str)
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

	var rec uevm.Record
	rec.BlockNum = tx.BlockNum 
	rec.BlockHash = common.HexToHash(tx.BlockHash)
	rec.TxIndex = int64(tx.TxIndex)
	rec.TxHash = common.HexToHash(tx.TxHash)
	ctrct_addr := common.HexToAddress(tx.To)
	last_line_rec,err := mchain.ReadLastLine()
	if err != nil {
		Info.Printf("Error getting last record: %v\n",err)
		os.Exit(1)
	}
	block_ctx := uevm.NewDummyBlockContext(big.NewInt(uevm.MainNetBlockNum) ,big.NewInt(uevm.MainNetTimeStamp))
	tx_ctx := new(vm.TxContext)
	tx_ctx.Origin = common.HexToAddress(tx.From) 
	tx_ctx.GasPrice = big.NewInt(uevm.TxDefaultGas)
	value := big.NewInt(0)
	value.SetString(tx.Value,10)
	input := tx.Input
	err = mchain.ExecMint(block_ctx,rec.TxHash,tx_ctx,input,value,ctrct_addr,last_line_rec.StateRoot,&rec,t0_addr,t1_addr)
	if err != nil {
		Info.Printf("Error executing ExecMint(): %v\n",err)
		os.Exit(1)
	}

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

	var rec uevm.Record
	rec.BlockNum = tx.BlockNum 
	rec.BlockHash = common.HexToHash(tx.BlockHash)
	rec.TxIndex = int64(tx.TxIndex)
	rec.TxHash = common.HexToHash(tx.TxHash)
	ctrct_addr := common.HexToAddress(tx.To)
	last_line_rec,err := mchain.ReadLastLine()
	if err != nil {
		Info.Printf("Error getting last record: %v\n",err)
		os.Exit(1)
	}
	block_ctx := uevm.NewDummyBlockContext(big.NewInt(uevm.MainNetBlockNum) ,big.NewInt(uevm.MainNetTimeStamp))
	tx_ctx := new(vm.TxContext)
	tx_ctx.Origin = common.HexToAddress(tx.From) 
	tx_ctx.GasPrice = big.NewInt(uevm.TxDefaultGas)
	value := big.NewInt(0)
	value.SetString(tx.Value,10)
	input := tx.Input
	err = mchain.ExecBurn(block_ctx,rec.TxHash,tx_ctx,input,value,ctrct_addr,last_line_rec.StateRoot,&rec)
	if err != nil {
		Info.Printf("Error executing ExecMint(): %v\n",err)
		os.Exit(1)
	}
}
func process_pool_debug_swap_events(tx *AugurTx,tx_hash common.Hash,pool_aid int64) error {

	encoded_logs,err := mchain.GetReceiptLogs(tx_hash)
	if err != nil { return err }

	var decoded_logs []*types.Log
	err = rlp.DecodeBytes(encoded_logs,&decoded_logs)
	if err != nil { return err }

	count := len(decoded_logs)
	for i:=0;i<count;i++ {
		lg := decoded_logs[i]
		if len(lg.Topics) > 0 {
			topic0 := lg.Topics[0].Bytes()
			if bytes.Equal(topic0,evt_dbg_swap_loop) {
				var evt UniV3DBGSwapLoop
				evt.BlockNum = tx.BlockNum
				evt.TimeStamp = tx.TimeStamp
				evt.TxIndex = int64(tx.TxIndex)
				evt.LogIndex = int64(i)
				evt.ContractAddr = lg.Address.String()
				evt.PoolAid = pool_aid

				var eth_evt IUniswapV3PoolEventsDBGSWAPLOOP
				err := dbg_abi.UnpackIntoInterface(&eth_evt,"DBG_SWAP_LOOP",lg.Data)
				if err != nil { return err }
				eth_evt.Pool = common.BytesToAddress(lg.Topics[1][12:])
				evt.Tick = eth_evt.Tick.Int64()
				evt.SqrtPrice = eth_evt.SqrtPriceX96.String()
				float_price := uevm.BinaryFixedToBigFloat(96,evt.SqrtPrice)
				evt.Price = float_price.String()
				evt.Liquidity = eth_evt.Liquidity.String()
				evt.StepAmountIn = eth_evt.StepAmountIn.String()
				evt.StepAmountOut = eth_evt.StepAmountOut.String()
				evt.FeeGrowthGlobalX128 = eth_evt.FeeGrowthGlobalX128.String()
				evt.FeeAmount = eth_evt.FeeAmount.String()
				float_fees := uevm.BinaryFixedToBigFloat(128,evt.FeeGrowthGlobalX128)
				evt.FeeGrowthDecoded = float_fees.String()
				storagew.Insert_dbg_swap_loop(&evt)
			}
		}
	}

	return nil
}
func process_pool_upd_pos_events(tx *AugurTx,tx_hash common.Hash,pool_aid int64) error {

	encoded_logs,err := mchain.GetReceiptLogs(tx_hash)
	if err != nil { return err }

	var decoded_logs []*types.Log
	err = rlp.DecodeBytes(encoded_logs,&decoded_logs)
	if err != nil { return err }

	count := len(decoded_logs)
	for i:=0;i<count;i++ {
		lg := decoded_logs[i]
		if len(lg.Topics) > 0 {
			topic0 := lg.Topics[0].Bytes()
			if bytes.Equal(topic0,evt_dbg_swap_loop) {
				var evt UniV3DBGUpdPos
				evt.BlockNum = tx.BlockNum
				evt.TimeStamp = tx.TimeStamp
				evt.TxIndex = int64(tx.TxIndex)
				evt.LogIndex = int64(i)
				evt.ContractAddr = lg.Address.String()
				evt.PoolAid = pool_aid

				var eth_evt IUniswapV3PoolEventsDBGUPDPOS
				err := dbg_abi.UnpackIntoInterface(&eth_evt,"DBG_UPD_POS",lg.Data)
				if err != nil { return err }
				evt.Tick = eth_evt.Tick.Int64()
				evt.LiquidityDelta = eth_evt.LiquidityDelta.String()
				evt.FeeGrowth0Before=uevm.BinaryFixedToBigFloat(128,eth_evt.FeeGrowthGlobal0X128Before.String()).String()
				evt.FeeGrowth1Before=uevm.BinaryFixedToBigFloat(128,eth_evt.FeeGrowthGlobal1X128Before.String()).String()
				evt.FeeGrowth0Inside=uevm.BinaryFixedToBigFloat(128,eth_evt.FeeGrowthInside0X128.String()).String()
				evt.FeeGrowth1Inside=uevm.BinaryFixedToBigFloat(128,eth_evt.FeeGrowthInside1X128.String()).String()
				evt.FlippedLower = eth_evt.FlippedLower
				evt.FlippedUpper= eth_evt.FlippedUpper
				storagew.Insert_dbg_upd_pos_event(&evt)
			}
		}
	}

	return nil
}
func process_pool_swap(storage *SQLStorage,tx *AugurTx,log *types.Log,log_index int) {

	Info.Printf(
		"EVENT: Swap. Tx %v TxIndex %v Log %v\n",
		tx.TxHash,tx.TxIndex,log.Index,
	)
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

	var rec uevm.Record
	rec.BlockNum = tx.BlockNum 
	rec.BlockHash = common.HexToHash(tx.BlockHash)
	rec.TxIndex = int64(tx.TxIndex)
	rec.TxHash = common.HexToHash(tx.TxHash)
	ctrct_addr := common.HexToAddress(tx.To)
	last_line_rec,err := mchain.ReadLastLine()
	if err != nil {
		Info.Printf("Error getting last record: %v\n",err)
		os.Exit(1)
	}
	block_ctx := uevm.NewDummyBlockContext(big.NewInt(uevm.MainNetBlockNum) ,big.NewInt(uevm.MainNetTimeStamp))
	tx_ctx := new(vm.TxContext)
	tx_ctx.Origin = common.HexToAddress(tx.From) 
	tx_ctx.GasPrice = big.NewInt(uevm.TxDefaultGas)
	value := big.NewInt(0)
	value.SetString(tx.Value,10)
	input := tx.Input
	err,_ = mchain.ExecCall(block_ctx,rec.TxHash,tx_ctx,input,value,ctrct_addr,last_line_rec.StateRoot,&rec)
	if err != nil {
		Info.Printf("Error executing ExecCall() for swap tx: %v\n",err)
		os.Exit(1)
	}

	err = process_pool_debug_swap_events(tx,rec.TxHash,pool_aid)
	if err != nil {
		Info.Printf("Error processing debug swap events: %v\n",err)
		Error.Printf("Error processing debug swap events: %v\n",err)
		os.Exit(1)
	}

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
