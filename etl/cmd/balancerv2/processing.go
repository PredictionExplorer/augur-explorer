package main

import (
	"os"
	"bytes"
	"encoding/hex"

	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/common"

	. "github.com/PredictionExplorer/augur-explorer/primitives"
	. "github.com/PredictionExplorer/augur-explorer/primitives/balancerv2"
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
		process_event_log(this.ETL.Storage,tx,log)
	}
}
func process_pool_created(storage *SQLStorage,tx *AugurTx,log *types.Log) {

	Info.Printf(
		"EVENT: PoolCreated. Tx %v TxIndex %v Log %v\n",
		tx.TxHash,tx.TxIndex,log.Index,
	)
	if !bytes.Equal(log.Address.Bytes(),caddrs.FactoryAddr.Bytes()) {
		Info.Printf("Skipping event, address doesn't match our address\n")
		return
	}

	var evt BalV2PoolCreated
	evt.BlockNum = tx.BlockNum
	evt.TimeStamp = tx.TimeStamp
	evt.TxIndex = int64(tx.TxIndex)
	evt.LogIndex = int64(log.Index)
	evt.ContractAddr = log.Address.String()

	pool_addr := common.BytesToAddress(log.Topics[1][12:])
	evt.PoolAddr = pool_addr.String()

	Info.Printf("PoolCreated{\n")
	Info.Printf("\tBlockNum: %v\n",evt.BlockNum)
	Info.Printf("\tTimeStamp: %v\n",evt.TimeStamp)
	Info.Printf("\tTxId: %v\n",evt.TxIndex)
	Info.Printf("\tLogIndex: %v\n",evt.LogIndex)
	Info.Printf("\tContractAddr: %v\n",evt.ContractAddr)
	Info.Printf("\tPoolAddr: %v\n",evt.PoolAddr)
	Info.Printf("}\n")
	storage.Insert_pool_created(&evt)
}
func process_pool_balance_changed(storage *SQLStorage,tx *AugurTx,log *types.Log) {

	Info.Printf(
		"EVENT: PoolBalanceChanged. Tx %v TxIndex %v Log %v\n",
		tx.TxHash,tx.TxIndex,log.Index,
	)
	if !bytes.Equal(log.Address.Bytes(),caddrs.VaultAddr.Bytes()) {
		Info.Printf("Skipping event, address doesn't match our address\n")
		return
	}
	var eth_evt BalancerV2VaultPoolBalanceChanged
	err := vault_abi.UnpackIntoInterface(&eth_evt,"PoolBalanceChanged",log.Data)
	if err != nil {
		Error.Printf("Can't UnpackIntoInterface for PoolBalanceChanged: %v\n",err)
		os.Exit(1)
	}

	var evt BalV2PoolBalanceChanged
	evt.BlockNum = tx.BlockNum
	evt.TimeStamp = tx.TimeStamp
	evt.TxIndex = int64(tx.TxIndex)
	evt.LogIndex = int64(log.Index)
	evt.ContractAddr = log.Address.String()

	evt.PoolId = hex.EncodeToString(log.Topics[1][:])
	liqprov_addr := common.BytesToAddress(log.Topics[2][12:])
	evt.LiqProvAddr = liqprov_addr.String()
	evt.Tokens = address_array_to_string(eth_evt.Tokens)
	evt.Deltas = Bigint_ptr_slice_to_str(&eth_evt.Deltas,",")
	evt.ProtocolFeeAmounts = Bigint_ptr_slice_to_str(&eth_evt.ProtocolFeeAmounts,",")

	Info.Printf("PoolBalanceChanged{\n")
	Info.Printf("\tBlockNum: %v\n",evt.BlockNum)
	Info.Printf("\tTimeStamp: %v\n",evt.TimeStamp)
	Info.Printf("\tTxId: %v\n",evt.TxIndex)
	Info.Printf("\tLogIndex: %v\n",evt.LogIndex)
	Info.Printf("\tContractAddr: %v\n",evt.ContractAddr)
	Info.Printf("\tLiquidityProviderAddr: %v\n",evt.LiqProvAddr)
	Info.Printf("\tTokens: %v\n",evt.Tokens)
	Info.Printf("\tDeltas: %v\n",evt.Deltas)
	Info.Printf("\tProtocolFeeAmounts: %v\n",evt.ProtocolFeeAmounts)
	Info.Printf("}\n")
	storage.Insert_pool_balance_changed(&evt)
}
func process_pool_registered(storage *SQLStorage,tx *AugurTx,log *types.Log) {

	Info.Printf(
		"EVENT: PoolRegistered. Tx %v TxIndex %v Log %v\n",
		tx.TxHash,tx.TxIndex,log.Index,
	)
	if !bytes.Equal(log.Address.Bytes(),caddrs.VaultAddr.Bytes()) {
		Info.Printf("Skipping event, address doesn't match our address\n")
		return
	}
	var eth_evt BalancerV2VaultPoolRegistered
	err := vault_abi.UnpackIntoInterface(&eth_evt,"PoolRegistered",log.Data)
	if err != nil {
		Error.Printf("Can't UnpackIntoInterface for PoolRegistered: %v\n",err)
		os.Exit(1)
	}

	var evt BalV2PoolRegistered
	evt.BlockNum = tx.BlockNum
	evt.TimeStamp = tx.TimeStamp
	evt.TxIndex = int64(tx.TxIndex)
	evt.LogIndex = int64(log.Index)
	evt.ContractAddr = log.Address.String()

	evt.PoolId = hex.EncodeToString(log.Topics[1][:])
	pool_addr := common.BytesToAddress(log.Topics[2][12:])
	evt.PoolAddr = pool_addr.String()
	evt.Specialization= int64(eth_evt.Specialization)
	Info.Printf("PoolRegistered {\n")
	Info.Printf("\tBlockNum: %v\n",evt.BlockNum)
	Info.Printf("\tTimeStamp: %v\n",evt.TimeStamp)
	Info.Printf("\tTxId: %v\n",evt.TxIndex)
	Info.Printf("\tLogIndex: %v\n",evt.LogIndex)
	Info.Printf("\tContractAddr: %v\n",evt.ContractAddr)
	Info.Printf("\tPoolId: %v\n",evt.PoolId)
	Info.Printf("\tPoolAddr: %v\n",evt.PoolAddr)
	Info.Printf("\tSpecialization: %v\n",evt.ContractAddr)
	Info.Printf("}\n")
	storage.Insert_pool_registered(&evt)
}
func process_internal_balance_changed(storage *SQLStorage,tx *AugurTx,log *types.Log) {

	Info.Printf(
		"EVENT: InternalBalanceChanged. Tx %v TxIndex %v Log %v\n",
		tx.TxHash,tx.TxIndex,log.Index,
	)
	if !bytes.Equal(log.Address.Bytes(),caddrs.VaultAddr.Bytes()) {
		Info.Printf("Skipping event, address doesn't match our address\n")
		return
	}
	var eth_evt BalancerV2VaultInternalBalanceChanged
	err := vault_abi.UnpackIntoInterface(&eth_evt,"InternalBalanceChanged",log.Data)
	if err != nil {
		Error.Printf("Can't UnpackIntoInterface for InternalBalanceChanged : %v\n",err)
		os.Exit(1)
	}

	var evt BalV2InternalBalanceChanged
	evt.BlockNum = tx.BlockNum
	evt.TimeStamp = tx.TimeStamp
	evt.TxIndex = int64(tx.TxIndex)
	evt.LogIndex = int64(log.Index)
	evt.ContractAddr = log.Address.String()

	user_addr := common.BytesToAddress(log.Topics[1][12:])
	token_addr := common.BytesToAddress(log.Topics[2][12:])
	evt.UserAddr = user_addr.String()
	evt.TokenAddr = token_addr.String()
	evt.Delta = eth_evt.Delta.String()

	Info.Printf("InternalBalanceChanged {\n")
	Info.Printf("\tBlockNum: %v\n",evt.BlockNum)
	Info.Printf("\tTimeStamp: %v\n",evt.TimeStamp)
	Info.Printf("\tTxId: %v\n",evt.TxIndex)
	Info.Printf("\tLogIndex: %v\n",evt.LogIndex)
	Info.Printf("\tContractAddr: %v\n",evt.ContractAddr)
	Info.Printf("\tUserAddr: %v\n",evt.UserAddr)
	Info.Printf("\tTokenAddr: %v\n",evt.TokenAddr)
	Info.Printf("\tDelta: %v\n",evt.Delta)
	Info.Printf("}\n")
	storage.Insert_internal_balance_changed(&evt)
}
func process_external_balance_transfer(storage *SQLStorage,tx *AugurTx,log *types.Log) {

	Info.Printf(
		"EVENT: ExternalBalanceTransfer. Tx %v TxIndex %v Log %v\n",
		tx.TxHash,tx.TxIndex,log.Index,
	)
	if !bytes.Equal(log.Address.Bytes(),caddrs.VaultAddr.Bytes()) {
		Info.Printf("Skipping event, address doesn't match our address\n")
		return
	}
	var eth_evt BalancerV2VaultExternalBalanceTransfer 
	err := vault_abi.UnpackIntoInterface(&eth_evt,"ExternalBalanceTransfer",log.Data)
	if err != nil {
		Error.Printf("Can't UnpackIntoInterface for ExternalBalanceTransfer: %v\n",err)
		os.Exit(1)
	}

	var evt BalV2ExternalBalanceTransfer
	evt.BlockNum = tx.BlockNum
	evt.TimeStamp = tx.TimeStamp
	evt.TxIndex = int64(tx.TxIndex)
	evt.LogIndex = int64(log.Index)
	evt.ContractAddr = log.Address.String()

	token_addr := common.BytesToAddress(log.Topics[1][12:])
	sender_addr := common.BytesToAddress(log.Topics[2][12:])
	evt.TokenAddr= token_addr.String()
	evt.SenderAddr = sender_addr.String()
	evt.RecipientAddr = eth_evt.Recipient.String()
	evt.Amount = eth_evt.Amount.String()

	Info.Printf("ExternalBalanceTransfer{\n")
	Info.Printf("\tBlockNum: %v\n",evt.BlockNum)
	Info.Printf("\tTimeStamp: %v\n",evt.TimeStamp)
	Info.Printf("\tTxId: %v\n",evt.TxIndex)
	Info.Printf("\tLogIndex: %v\n",evt.LogIndex)
	Info.Printf("\tContractAddr: %v\n",evt.ContractAddr)
	Info.Printf("\tTokenAddr: %v\n",evt.TokenAddr)
	Info.Printf("\tRecipientAddr: %v\n",evt.RecipientAddr)
	Info.Printf("\tAmount: %v\n",evt.Amount)
	Info.Printf("}\n")
	storage.Insert_external_balance_transfer(&evt)
}
func process_swap(storage *SQLStorage,tx *AugurTx,log *types.Log) {

	Info.Printf(
		"EVENT: Swap. Tx %v TxIndex %v Log %v\n",
		tx.TxHash,tx.TxIndex,log.Index,
	)
	if !bytes.Equal(log.Address.Bytes(),caddrs.VaultAddr.Bytes()) {
		Info.Printf("Skipping event, address doesn't match our address\n")
		return
	}
	var eth_evt BalancerV2VaultSwap
	err := vault_abi.UnpackIntoInterface(&eth_evt,"Swap",log.Data)
	if err != nil {
		Error.Printf("Can't UnpackIntoInterface for Swap: %v\n",err)
		os.Exit(1)
	}

	var evt BalV2Swap
	evt.BlockNum = tx.BlockNum
	evt.TimeStamp = tx.TimeStamp
	evt.TxIndex = int64(tx.TxIndex)
	evt.LogIndex = int64(log.Index)
	evt.ContractAddr = log.Address.String()

	evt.PoolId = hex.EncodeToString(log.Topics[1][:])
	tokenin_addr := common.BytesToAddress(log.Topics[2][12:])
	tokenout_addr := common.BytesToAddress(log.Topics[3][12:])
	evt.TokenInAddr = tokenin_addr.String()
	evt.TokenOutAddr = tokenout_addr.String()
	evt.AmountIn = eth_evt.AmountIn.String()
	evt.AmountOut = eth_evt.AmountOut.String()
	Info.Printf("Swap {\n")
	Info.Printf("\tBlockNum: %v\n",evt.BlockNum)
	Info.Printf("\tTimeStamp: %v\n",evt.TimeStamp)
	Info.Printf("\tTxId: %v\n",evt.TxIndex)
	Info.Printf("\tLogIndex: %v\n",evt.LogIndex)
	Info.Printf("\tContractAddr: %v\n",evt.ContractAddr)
	Info.Printf("\tPoolId: %v\n",evt.PoolId)
	Info.Printf("\tTokenIn: %v\n",evt.TokenInAddr)
	Info.Printf("\tTokenOut: %v\n",evt.TokenOutAddr)
	Info.Printf("\tAmountIn: %v\n",evt.AmountIn)
	Info.Printf("\tAmountOut: %v\n",evt.AmountOut)
	storage.Insert_swap(&evt)
}
func process_pool_balance_managed(storage *SQLStorage,tx *AugurTx,log *types.Log) {

	Info.Printf(
		"EVENT: PoolBalanceManaged. Tx %v TxIndex %v Log %v\n",
		tx.TxHash,tx.TxIndex,log.Index,
	)
	if !bytes.Equal(log.Address.Bytes(),caddrs.VaultAddr.Bytes()) {
		Info.Printf("Skipping event, address doesn't match our address\n")
		return
	}
	var eth_evt BalancerV2VaultPoolBalanceManaged
	err := vault_abi.UnpackIntoInterface(&eth_evt,"PoolBalanceManaged",log.Data)
	if err != nil {
		Error.Printf("Can't UnpackIntoInterface for PoolBalanceManaged: %v\n",err)
		os.Exit(1)
	}

	var evt BalV2PoolBalanceManaged

	evt.BlockNum = tx.BlockNum
	evt.TimeStamp = tx.TimeStamp
	evt.TxIndex = int64(tx.TxIndex)
	evt.LogIndex = int64(log.Index)
	evt.ContractAddr = log.Address.String()

	evt.PoolId = hex.EncodeToString(log.Topics[1][:])
	asset_manager_addr := common.BytesToAddress(log.Topics[2][12:])
	evt.AssetManagerAddr = asset_manager_addr.String()
	evt.TokenAddr = eth_evt.Token.String()
	evt.CashDelta = eth_evt.CashDelta.String()
	evt.ManagedDelta = eth_evt.ManagedDelta.String()
	Info.Printf("PoolBalanceManaged{\n")
	Info.Printf("\tBlockNum: %v\n",evt.BlockNum)
	Info.Printf("\tTimeStamp: %v\n",evt.TimeStamp)
	Info.Printf("\tTxId: %v\n",evt.TxIndex)
	Info.Printf("\tLogIndex: %v\n",evt.LogIndex)
	Info.Printf("\tContractAddr: %v\n",evt.ContractAddr)
	Info.Printf("\tPoolId: %v\n",evt.PoolId)
	Info.Printf("\tAssetManagerAddr: %v\n",evt.AssetManagerAddr)
	Info.Printf("\tTokenAddr: %v\n",evt.TokenAddr)
	Info.Printf("\tCashDelta: %v\n",evt.CashDelta)
	Info.Printf("\tManagedDelta: %v\n",evt.ManagedDelta)
	storage.Insert_pool_balance_managed(&evt)
}
func process_swap_fee_changed(storage *SQLStorage,tx *AugurTx,log *types.Log) {

	Info.Printf(
		"EVENT: SwapFeePercentageChanged. Tx %v TxIndex %v Log %v\n",
		tx.TxHash,tx.TxIndex,log.Index,
	)
	var eth_evt BalancerV2SwapFeePercentageChangedSwapFeePercentageChanged
	err := swapfee_abi.UnpackIntoInterface(&eth_evt,"SwapFeePercentageChanged",log.Data)
	if err != nil {
		Error.Printf("Can't UnpackIntoInterface for SwapFeePercentageChanged: %v\n",err)
		os.Exit(1)
	}

	var evt BalV2SwapFeePercentageChanged
	evt.BlockNum = tx.BlockNum
	evt.TimeStamp = tx.TimeStamp
	evt.TxIndex = int64(tx.TxIndex)
	evt.LogIndex = int64(log.Index)
	evt.ContractAddr = log.Address.String()
	evt.SwapFeePercentage = eth_evt.SwapFeePercentage.String()

	Info.Printf("SwapFeePercentageChanged {\n")
	Info.Printf("\tBlockNum: %v\n",evt.BlockNum)
	Info.Printf("\tTimeStamp: %v\n",evt.TimeStamp)
	Info.Printf("\tTxId: %v\n",evt.TxIndex)
	Info.Printf("\tLogIndex: %v\n",evt.LogIndex)
	Info.Printf("\tContractAddr: %v\n",evt.ContractAddr)
	Info.Printf("\tSwapFeePercentage: %v\n",evt.SwapFeePercentage)
	Info.Printf("\n")
	storage.Insert_swap_fee_percentage_changed(&evt)
}
func process_tokens_registered(storage *SQLStorage,tx *AugurTx,log *types.Log) {

	Info.Printf(
		"EVENT: TokensRegistered. Tx %v TxIndex %v Log %v\n",
		tx.TxHash,tx.TxIndex,log.Index,
	)
	if !bytes.Equal(log.Address.Bytes(),caddrs.VaultAddr.Bytes()) {
		Info.Printf("Skipping event, address doesn't match our address\n")
		return
	}
	var eth_evt BalancerV2VaultTokensRegistered
	err := vault_abi.UnpackIntoInterface(&eth_evt,"TokensRegistered",log.Data)
	if err != nil {
		Error.Printf("Can't UnpackIntoInterface for TokenRegistered: %v\n",err)
		os.Exit(1)
	}

	var evt BalV2TokensRegistered
	evt.BlockNum = tx.BlockNum
	evt.TimeStamp = tx.TimeStamp
	evt.TxIndex = int64(tx.TxIndex)
	evt.LogIndex = int64(log.Index)
	evt.ContractAddr = log.Address.String()

	evt.Tokens = address_array_to_string(eth_evt.Tokens)
	evt.AssetManagers = address_array_to_string(eth_evt.AssetManagers)
	evt.PoolId = hex.EncodeToString(log.Topics[1][:])

	Info.Printf("TokenRegistered{\n")
	Info.Printf("\tBlockNum: %v\n",evt.BlockNum)
	Info.Printf("\tTimeStamp: %v\n",evt.TimeStamp)
	Info.Printf("\tTxId: %v\n",evt.TxIndex)
	Info.Printf("\tLogIndex: %v\n",evt.LogIndex)
	Info.Printf("\tContractAddr: %v\n",evt.ContractAddr)
	Info.Printf("\tTokens: %v\n",evt.Tokens)
	Info.Printf("\tAssetManagers: %v\n",evt.AssetManagers)
	Info.Printf("\tPoolId: %v\n",evt.PoolId)
	Info.Printf("}\n")
	storage.Insert_tokens_registered(&evt)
}
func process_tokens_deregistered(storage *SQLStorage,tx *AugurTx,log *types.Log) {

	Info.Printf(
		"EVENT: TokensDeregistered. Tx %v TxIndex %v Log %v\n",
		tx.TxHash,tx.TxIndex,log.Index,
	)
	if !bytes.Equal(log.Address.Bytes(),caddrs.VaultAddr.Bytes()) {
		Info.Printf("Skipping event, address doesn't match our address\n")
		return
	}
	var eth_evt BalancerV2VaultTokensRegistered
	err := vault_abi.UnpackIntoInterface(&eth_evt,"TokensDeregistered",log.Data)
	if err != nil {
		Error.Printf("Can't UnpackIntoInterface for TokenDeregistered: %v\n",err)
		os.Exit(1)
	}

	var evt BalV2TokensDeregistered
	evt.BlockNum = tx.BlockNum
	evt.TimeStamp = tx.TimeStamp
	evt.TxIndex = int64(tx.TxIndex)
	evt.LogIndex = int64(log.Index)
	evt.ContractAddr = log.Address.String()

	evt.Tokens = address_array_to_string(eth_evt.Tokens)
	evt.PoolId = hex.EncodeToString(log.Topics[1][:])

	Info.Printf("TokenDeregistered{\n")
	Info.Printf("\tBlockNum: %v\n",evt.BlockNum)
	Info.Printf("\tTimeStamp: %v\n",evt.TimeStamp)
	Info.Printf("\tTxId: %v\n",evt.TxIndex)
	Info.Printf("\tLogIndex: %v\n",evt.LogIndex)
	Info.Printf("\tContractAddr: %v\n",evt.ContractAddr)
	Info.Printf("\tTokens: %v\n",evt.Tokens)
	Info.Printf("\tPoolId: %v\n",evt.PoolId)
	Info.Printf("}\n")
	storage.Insert_tokens_deregistered(&evt)
}
func process_flash_loan(storage *SQLStorage,tx *AugurTx,log *types.Log) {

	Info.Printf(
		"EVENT: FlashLoan. Tx %v TxIndex %v Log %v\n",
		tx.TxHash,tx.TxIndex,log.Index,
	)
	if !bytes.Equal(log.Address.Bytes(),caddrs.VaultAddr.Bytes()) {
		Info.Printf("Skipping event, address doesn't match our address\n")
		return
	}
	var eth_evt BalancerV2VaultFlashLoan
	err := vault_abi.UnpackIntoInterface(&eth_evt,"FlashLoan",log.Data)
	if err != nil {
		Error.Printf("Can't UnpackIntoInterface for FlashLoan: %v\n",err)
		os.Exit(1)
	}

	var evt BalV2FlashLoan

	evt.BlockNum = tx.BlockNum
	evt.TimeStamp = tx.TimeStamp
	evt.TxIndex = int64(tx.TxIndex)
	evt.LogIndex = int64(log.Index)
	evt.ContractAddr = log.Address.String()

	//recipient_addr := common.BytesToAddress(log.Topics[1][12:])
	//evt.RecipientAddr = recipient_addr.String()
	evt.RecipientAddr = eth_evt.Recipient.String()
	//token_addr := common.BytesToAddress(log.Topics[2][12:])
	//evt.TokenAddr = token_addr.String()
	evt.TokenAddr = eth_evt.Token.String()
	evt.Amount= eth_evt.Amount.String()
	evt.FeeAmount = eth_evt.FeeAmount.String()
	Info.Printf("FlashLoan{\n")
	Info.Printf("\tBlockNum: %v\n",evt.BlockNum)
	Info.Printf("\tTimeStamp: %v\n",evt.TimeStamp)
	Info.Printf("\tTxId: %v\n",evt.TxIndex)
	Info.Printf("\tLogIndex: %v\n",evt.LogIndex)
	Info.Printf("\tContractAddr: %v\n",evt.ContractAddr)
	Info.Printf("\tRecipientAddr: %v\n",evt.RecipientAddr)
	Info.Printf("\tTokenAddr: %v\n",evt.TokenAddr)
	Info.Printf("\tAmount: %v\n",evt.Amount)
	Info.Printf("\tFeeAmount: %v\n",evt.FeeAmount)
	if len(evt.TokenAddr) == 0 {
		Error.Printf("Token address is empty\n")
		os.Exit(1)
	}
	if len(evt.RecipientAddr) == 0 {
		Error.Printf("Recipient address is empty\n")
		os.Exit(1)
	}
	storage.Insert_flash_loan(&evt)
}
func process_event_log(storage *SQLStorage,tx *AugurTx,log *types.Log) {

	if len(log.Topics) == 0 { return }
	topic0 := log.Topics[0].Bytes()
	//Info.Printf("topic0=%v\n",hex.EncodeToString(topic0[:]))
	if bytes.Equal(topic0,evt_pool_created) {
		process_pool_created(storage,tx,log)
	}
	if bytes.Equal(topic0,evt_pool_balance_changed) {
		process_pool_balance_changed(storage,tx,log)
	}
	if bytes.Equal(topic0,evt_pool_registered) {
		process_pool_registered(storage,tx,log)
	}
	if bytes.Equal(topic0,evt_external_balance_transfer) {
		process_external_balance_transfer(storage,tx,log)
	}
	if bytes.Equal(topic0,evt_internal_balance_changed) {
		process_internal_balance_changed(storage,tx,log)
	}
	if bytes.Equal(topic0,evt_token_deregistered) {
		process_tokens_deregistered(storage,tx,log)
	}
	if bytes.Equal(topic0,evt_token_registered) {
		process_tokens_registered(storage,tx,log)
	}
	if bytes.Equal(topic0,evt_swap) {
		process_swap(storage,tx,log)
	}
	if bytes.Equal(topic0,evt_swap_fee_changed) {
		process_swap_fee_changed(storage,tx,log)
	}
	if bytes.Equal(topic0,evt_pool_balance_managed) {
		process_pool_balance_managed(storage,tx,log)
	}
	if bytes.Equal(topic0,evt_flash_loan) {
		process_flash_loan(storage,tx,log)
	}
}
