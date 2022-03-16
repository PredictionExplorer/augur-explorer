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
)

func address_array_to_string(addresses []common.Address) string {

	var output string
	for i:=0; i<len(addresses); i++ {
		if len(output) > 0 { output = output + "," }
		output = output + addresses[i].String()
	}
	return output
}
func process_transaction(storage *SQLStorage,tx *AugurTx,rcpt *types.Receipt) {


	logs := rcpt.Logs
	for i:=0; i<len(logs); i++ {
		log := logs[i]
		process_event_log(storage,tx,log)
	}
}
func process_pool_created(storage *SQLStorage,tx *AugurTx,log *types.Log) {

	var eth_evt BalancerV2WeightedPoolFactoryPoolCreated
	err := pool_factory_abi.UnpackIntoInterface(&eth_evt,"PoolCreated",log.Data)
	if err != nil {
		Error.Printf("Can't UnpackIntoInterface for PoolCreated: %v\n",err)
		os.Exit(1)
	}

	var evt BalV2PoolCreated
	evt.BlockNum = tx.BlockNum
	evt.TimeStamp = tx.TimeStamp
	evt.TxIndex = int64(tx.TxIndex)
	evt.LogIndex = int64(log.Index)
	evt.ContractAddr = log.Address.String()

	pool_addr := common.BytesToAddress(log.Topics[2][12:])
	evt.PoolAddr = pool_addr.String()
	storage.Insert_pool_created(&evt)
}
func process_pool_balance_changed(storage *SQLStorage,tx *AugurTx,log *types.Log) {

}
func process_pool_registered(storage *SQLStorage,tx *AugurTx,log *types.Log) {

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

	pool_addr := common.BytesToAddress(log.Topics[2][12:])
	evt.PoolAddr = pool_addr.String()
	evt.Specialization= int64(eth_evt.Specialization)
	storage.Insert_pool_registered(&evt)
}
func process_internal_balance_changed(storage *SQLStorage,tx *AugurTx,log *types.Log) {

	var eth_evt BalancerV2VaultInternalBalanceChanged
	err := pool_factory_abi.UnpackIntoInterface(&eth_evt,"InternalBalanceChanged",log.Data)
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
	storage.Insert_internal_balance_changed(&evt)
}
func process_external_balance_transfer(storage *SQLStorage,tx *AugurTx,log *types.Log) {

	var eth_evt BalancerV2VaultExternalBalanceTransfer 
	err := pool_factory_abi.UnpackIntoInterface(&eth_evt,"ExternalBalanceTransfer",log.Data)
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

	storage.Insert_external_balance_transfer(&evt)
}
func process_swap(storage *SQLStorage,tx *AugurTx,log *types.Log) {

	var eth_evt BalancerV2VaultSwap
	err := pool_factory_abi.UnpackIntoInterface(&eth_evt,"Swap",log.Data)
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
	storage.Insert_swap(&evt)
}
func process_tokens_registered(storage *SQLStorage,tx *AugurTx,log *types.Log) {

	var eth_evt BalancerV2VaultTokensRegistered
	err := pool_factory_abi.UnpackIntoInterface(&eth_evt,"TokensRegistered",log.Data)
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

	evt.PoolId = hex.EncodeToString(eth_evt.PoolId[:])
	storage.Insert_tokens_registered(&evt)
}
func process_tokens_deregistered(storage *SQLStorage,tx *AugurTx,log *types.Log) {

	var eth_evt BalancerV2VaultTokensRegistered
	err := pool_factory_abi.UnpackIntoInterface(&eth_evt,"TokensDeregistered",log.Data)
	if err != nil {
		Error.Printf("Can't UnpackIntoInterface for TokenDeegistered: %v\n",err)
		os.Exit(1)
	}

	var evt BalV2TokensDeregistered
	evt.BlockNum = tx.BlockNum
	evt.TimeStamp = tx.TimeStamp
	evt.TxIndex = int64(tx.TxIndex)
	evt.LogIndex = int64(log.Index)
	evt.ContractAddr = log.Address.String()

	evt.Tokens = address_array_to_string(eth_evt.Tokens)

	evt.PoolId = hex.EncodeToString(eth_evt.PoolId[:])
	storage.Insert_tokens_deregistered(&evt)
}
func process_event_log(storage *SQLStorage,tx *AugurTx,log *types.Log) {

	if len(log.Topics) > 0 { return }
	topic0 := log.Topics[0].Bytes()
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
}
