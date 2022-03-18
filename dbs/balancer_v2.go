package dbs

import (
	"fmt"
	"os"
	//"strings"
	//"database/sql"
	_  "github.com/lib/pq"

	p "github.com/PredictionExplorer/augur-explorer/primitives/balancerv2"
)

func (ss *SQLStorage) Insert_pool_created(evt *p.BalV2PoolCreated) {

	contract_aid := ss.Layer1_lookup_or_insert_address_id(evt.ContractAddr)
	pool_aid := ss.Layer1_lookup_or_insert_address_id(evt.PoolAddr)
	var query string
	query =  "INSERT INTO "+ss.schema_name+".pool_created(block_num,time_stamp,tx_index,log_index,contract_aid,pool_aid) " +
				"VALUES($1,TO_TIMESTAMP($2),$3,$4,$5,$6)"
	_,err := ss.db.Exec(query,
		evt.BlockNum,
		evt.TimeStamp,
		evt.TxIndex,
		evt.LogIndex,
		contract_aid,
		pool_aid,
	)
	if err != nil {
		ss.Log_msg(fmt.Sprintf("DB error: can't insert into pool_created table: %v\n",err))
		os.Exit(1)
	}

}
func (ss *SQLStorage) Insert_pool_registered(evt *p.BalV2PoolRegistered) {

	contract_aid := ss.Layer1_lookup_or_insert_address_id(evt.ContractAddr)
	pool_aid := ss.Layer1_lookup_or_insert_address_id(evt.PoolAddr)
	var query string
	query =  "INSERT INTO "+ss.schema_name+".pool_reg("+
				"block_num,time_stamp,tx_index,log_index,contract_aid,pool_id,pool_aid,specialization"+
			") VALUES($1,TO_TIMESTAMP($2),$3,$4,$5,$6,$7,$8)"
	_,err := ss.db.Exec(query,
		evt.BlockNum,
		evt.TimeStamp,
		evt.TxIndex,
		evt.LogIndex,
		contract_aid,
		evt.PoolId,
		pool_aid,
		evt.Specialization,
	)
	if err != nil {
		ss.Log_msg(fmt.Sprintf("DB error: can't insert into pool_created table: %v\n",err))
		os.Exit(1)
	}
}
func (ss *SQLStorage) Insert_pool_balance_changed(evt *p.BalV2PoolBalanceChanged) {

	contract_aid := ss.Layer1_lookup_or_insert_address_id(evt.ContractAddr)
	liqprov_aid := ss.Layer1_lookup_or_insert_address_id(evt.LiqProvAddr)
	var query string
	query =  "INSERT INTO "+ss.schema_name+".pool_bal("+
				"block_num,time_stamp,tx_index,log_index,contract_aid,"+
				"pool_id,liqprov_aid,tokens,deltas,fee_amounts"+
			") VALUES($1,TO_TIMESTAMP($2),$3,$4,$5,$6,$7,$8,$9,$10)"
	_,err := ss.db.Exec(query,
		evt.BlockNum,
		evt.TimeStamp,
		evt.TxIndex,
		evt.LogIndex,
		contract_aid,
		evt.PoolId,
		liqprov_aid,
		evt.Tokens,
		evt.Deltas,
		evt.ProtocolFeeAmounts,
	)
	if err != nil {
		ss.Log_msg(fmt.Sprintf("DB error: can't insert into pool_bal table: %v\n",err))
		os.Exit(1)
	}

}
func (ss *SQLStorage) Insert_tokens_registered(evt *p.BalV2TokensRegistered) {

	contract_aid := ss.Layer1_lookup_or_insert_address_id(evt.ContractAddr)
	var query string
	query =  "INSERT INTO "+ss.schema_name+".tokens_reg("+
				"block_num,time_stamp,tx_index,log_index,contract_aid,pool_id,tokens,managers"+
			") VALUES($1,TO_TIMESTAMP($2),$3,$4,$5,$6,$7,$8)"
	_,err := ss.db.Exec(query,
		evt.BlockNum,
		evt.TimeStamp,
		evt.TxIndex,
		evt.LogIndex,
		contract_aid,
		evt.PoolId,
		evt.Tokens,
		evt.AssetManagers,
	)
	if err != nil {
		ss.Log_msg(fmt.Sprintf("DB error: can't insert into pool_registered table: %v\n",err))
		os.Exit(1)
	}

}
func (ss *SQLStorage) Insert_tokens_deregistered(evt *p.BalV2TokensDeregistered) {

	contract_aid := ss.Layer1_lookup_or_insert_address_id(evt.ContractAddr)
	var query string
	query =  "INSERT INTO "+ss.schema_name+".tokens_dereg("+
				"block_num,time_stamp,tx_index,log_index,contract_aid,pool_id,tokens"+
			") VALUES($1,TO_TIMESTAMP($2),$3,$4,$5,$6,$7)"
	_,err := ss.db.Exec(query,
		evt.BlockNum,
		evt.TimeStamp,
		evt.TxIndex,
		evt.LogIndex,
		contract_aid,
		evt.PoolId,
		evt.Tokens,
	)
	if err != nil {
		ss.Log_msg(fmt.Sprintf("DB error: can't insert into tokens_dereg table: %v\n",err))
		os.Exit(1)
	}
}
func (ss *SQLStorage) Insert_internal_balance_changed(evt *p.BalV2InternalBalanceChanged) {

	contract_aid := ss.Layer1_lookup_or_insert_address_id(evt.ContractAddr)
	user_aid := ss.Layer1_lookup_or_insert_address_id(evt.UserAddr)
	token_aid := ss.Layer1_lookup_or_insert_address_id(evt.TokenAddr)
	var query string
	query =  "INSERT INTO "+ss.schema_name+".ibalance("+
				"block_num,time_stamp,tx_index,log_index,contract_aid,user_aid,token_aid,delta"+
			") VALUES($1,TO_TIMESTAMP($2),$3,$4,$5,$6,$7,$8)"
	_,err := ss.db.Exec(query,
		evt.BlockNum,
		evt.TimeStamp,
		evt.TxIndex,
		evt.LogIndex,
		contract_aid,
		user_aid,
		token_aid,
		evt.Delta,
	)
	if err != nil {
		ss.Log_msg(fmt.Sprintf("DB error: can't insert into ibalance table: %v\n",err))
		os.Exit(1)
	}
}
func (ss *SQLStorage) Insert_external_balance_transfer(evt *p.BalV2ExternalBalanceTransfer) {

	contract_aid := ss.Layer1_lookup_or_insert_address_id(evt.ContractAddr)
	sender_aid := ss.Layer1_lookup_or_insert_address_id(evt.SenderAddr)
	recipient_aid := ss.Layer1_lookup_or_insert_address_id(evt.RecipientAddr)
	token_aid := ss.Layer1_lookup_or_insert_address_id(evt.TokenAddr)
	var query string
	query =  "INSERT INTO "+ss.schema_name+".ebal_transf("+
				"block_num,time_stamp,tx_index,log_index,contract_aid,"+
				"sender_aid,recipient_aid,token_aid,amount"+
			") VALUES($1,TO_TIMESTAMP($2),$3,$4,$5,$6,$7,$8,$9)"
	_,err := ss.db.Exec(query,
		evt.BlockNum,
		evt.TimeStamp,
		evt.TxIndex,
		evt.LogIndex,
		contract_aid,
		sender_aid,
		recipient_aid,
		token_aid,
		evt.Amount,
	)
	if err != nil {
		ss.Log_msg(fmt.Sprintf("DB error: can't insert into ebal_transf table: %v\n",err))
		os.Exit(1)
	}

}
func (ss *SQLStorage) Insert_swap(evt *p.BalV2Swap) {

	contract_aid := ss.Layer1_lookup_or_insert_address_id(evt.ContractAddr)
	token_in_aid := ss.Layer1_lookup_or_insert_address_id(evt.TokenInAddr)
	token_out_aid := ss.Layer1_lookup_or_insert_address_id(evt.TokenOutAddr)
	var query string
	query =  "INSERT INTO "+ss.schema_name+".swap("+
				"block_num,time_stamp,tx_index,log_index,contract_aid,"+
				"pool_id,token_in_aid,token_out_aid,amount_in,amount_out"+
			") VALUES($1,TO_TIMESTAMP($2),$3,$4,$5,$6,$7,$8,$9,$10)"
	_,err := ss.db.Exec(query,
		evt.BlockNum,
		evt.TimeStamp,
		evt.TxIndex,
		evt.LogIndex,
		contract_aid,
		evt.PoolId,
		token_in_aid,
		token_out_aid,
		evt.AmountIn,
		evt.AmountOut,
	)
	if err != nil {
		ss.Log_msg(fmt.Sprintf("DB error: can't insert into swap table: %v\n",err))
		os.Exit(1)
	}

}
func (ss *SQLStorage) Insert_swap_fee_percentage_changed(evt *p.BalV2SwapFeePercentageChanged) {

	contract_aid := ss.Layer1_lookup_or_insert_address_id(evt.ContractAddr)
	var query string
	query =  "INSERT INTO "+ss.schema_name+".swap_fee("+
				"block_num,time_stamp,tx_index,log_index,contract_aid,"+
				"swap_fee"+
			") VALUES($1,TO_TIMESTAMP($2),$3,$4,$5,$6)"
	_,err := ss.db.Exec(query,
		evt.BlockNum,
		evt.TimeStamp,
		evt.TxIndex,
		evt.LogIndex,
		contract_aid,
		evt.SwapFeePercentage,
	)
	if err != nil {
		ss.Log_msg(fmt.Sprintf("DB error: can't insert into swap_fee table: %v\n",err))
		os.Exit(1)
	}

}
func (ss *SQLStorage) Insert_pool_balance_managed(evt *p.BalV2PoolBalanceManaged) {

	contract_aid := ss.Layer1_lookup_or_insert_address_id(evt.ContractAddr)
	asset_manager_aid := ss.Layer1_lookup_or_insert_address_id(evt.AssetManagerAddr)
	token_aid := ss.Layer1_lookup_or_insert_address_id(evt.TokenAddr)
	var query string
	query =  "INSERT INTO "+ss.schema_name+".pool_bm("+
				"block_num,time_stamp,tx_index,log_index,contract_aid,"+
				"pool_id,asset_mgr_aid,token_aid,cash_delta,managed_delta"+
			") VALUES($1,TO_TIMESTAMP($2),$3,$4,$5,$6,$7,$8,$9,$10)"
	_,err := ss.db.Exec(query,
		evt.BlockNum,
		evt.TimeStamp,
		evt.TxIndex,
		evt.LogIndex,
		contract_aid,
		evt.PoolId,
		asset_manager_aid,
		token_aid,
		evt.CashDelta,
		evt.ManagedDelta,
	)
	if err != nil {
		ss.Log_msg(fmt.Sprintf("DB error: can't insert into swap table: %v\n",err))
		os.Exit(1)
	}
}
func (ss *SQLStorage) Insert_flash_loan(evt *p.BalV2FlashLoan) {

	contract_aid := ss.Layer1_lookup_or_insert_address_id(evt.ContractAddr)
	token_aid := ss.Layer1_lookup_or_insert_address_id(evt.TokenAddr)
	recipient_aid := ss.Layer1_lookup_or_insert_address_id(evt.RecipientAddr)
	var query string
	query =  "INSERT INTO "+ss.schema_name+".flash_loan("+
				"block_num,time_stamp,tx_index,log_index,contract_aid,"+
				"recipient_aid,token_aid,amount,fee_amount"+
			") VALUES($1,TO_TIMESTAMP($2),$3,$4,$5,$6,$7,$8,$9)"
	_,err := ss.db.Exec(query,
		evt.BlockNum,
		evt.TimeStamp,
		evt.TxIndex,
		evt.LogIndex,
		contract_aid,
		recipient_aid,
		token_aid,
		evt.Amount,
		evt.FeeAmount,
	)
	if err != nil {
		ss.Log_msg(fmt.Sprintf("DB error: can't insert into flash_loan table: %v\n",err))
		os.Exit(1)
	}

}
