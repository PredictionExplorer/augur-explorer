package balancerv2

import (
	"fmt"
	"os"
	//"strings"
	//"database/sql"
	_  "github.com/lib/pq"

	p "github.com/PredictionExplorer/augur-explorer/primitives/balancerv2"
)
func (sw *SQLStorageWrapper) Insert_pool_created(evt *p.BalV2PoolCreated) {

	contract_aid := sw.S.Layer1_lookup_or_insert_address_id(evt.ContractAddr)
	pool_aid := sw.S.Layer1_lookup_or_insert_address_id(evt.PoolAddr)
	var query string
	query =  "INSERT INTO "+sw.S.SchemaName()+".pool_created(block_num,time_stamp,tx_index,log_index,contract_aid,pool_aid) " +
				"VALUES($1,TO_TIMESTAMP($2),$3,$4,$5,$6)"
	_,err := sw.S.Db().Exec(query,
		evt.BlockNum,
		evt.TimeStamp,
		evt.TxIndex,
		evt.LogIndex,
		contract_aid,
		pool_aid,
	)
	if err != nil {
		sw.S.Log_msg(fmt.Sprintf("DB error: can't insert into pool_created table: %v\n",err))
		os.Exit(1)
	}

}
func (sw *SQLStorageWrapper) Insert_pool_registered(evt *p.BalV2PoolRegistered) {

	contract_aid := sw.S.Layer1_lookup_or_insert_address_id(evt.ContractAddr)
	pool_aid := sw.S.Layer1_lookup_or_insert_address_id(evt.PoolAddr)
	var query string
	query =  "INSERT INTO "+sw.S.SchemaName()+".pool_reg("+
				"block_num,time_stamp,tx_index,log_index,contract_aid,pool_id,pool_aid,specialization"+
			") VALUES($1,TO_TIMESTAMP($2),$3,$4,$5,$6,$7,$8)"
	_,err := sw.S.Db().Exec(query,
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
		sw.S.Log_msg(fmt.Sprintf("DB error: can't insert into pool_created table: %v\n",err))
		os.Exit(1)
	}
}
func (sw *SQLStorageWrapper) Insert_pool_balance_changed(evt *p.BalV2PoolBalanceChanged) {

	contract_aid := sw.S.Layer1_lookup_or_insert_address_id(evt.ContractAddr)
	liqprov_aid := sw.S.Layer1_lookup_or_insert_address_id(evt.LiqProvAddr)
	var query string
	query =  "INSERT INTO "+sw.S.SchemaName()+".pool_bal("+
				"block_num,time_stamp,tx_index,log_index,contract_aid,"+
				"pool_id,liqprov_aid,tokens,deltas,proto_fee_amounts"+
			") VALUES($1,TO_TIMESTAMP($2),$3,$4,$5,$6,$7,$8,$9,$10)"
	_,err := sw.S.Db().Exec(query,
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
		sw.S.Log_msg(fmt.Sprintf("DB error: can't insert into pool_bal table: %v\n",err))
		os.Exit(1)
	}

}
func (sw *SQLStorageWrapper) Insert_tokens_registered(evt *p.BalV2TokensRegistered) {

	contract_aid := sw.S.Layer1_lookup_or_insert_address_id(evt.ContractAddr)
	var query string
	query =  "INSERT INTO "+sw.S.SchemaName()+".tokens_reg("+
				"block_num,time_stamp,tx_index,log_index,contract_aid,pool_id,tokens,managers"+
			") VALUES($1,TO_TIMESTAMP($2),$3,$4,$5,$6,$7,$8)"
	_,err := sw.S.Db().Exec(query,
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
		sw.S.Log_msg(fmt.Sprintf("DB error: can't insert into pool_registered table: %v\n",err))
		os.Exit(1)
	}

}
func (sw *SQLStorageWrapper) Insert_tokens_deregistered(evt *p.BalV2TokensDeregistered) {

	contract_aid := sw.S.Layer1_lookup_or_insert_address_id(evt.ContractAddr)
	var query string
	query =  "INSERT INTO "+sw.S.SchemaName()+".tokens_dereg("+
				"block_num,time_stamp,tx_index,log_index,contract_aid,pool_id,tokens"+
			") VALUES($1,TO_TIMESTAMP($2),$3,$4,$5,$6,$7)"
	_,err := sw.S.Db().Exec(query,
		evt.BlockNum,
		evt.TimeStamp,
		evt.TxIndex,
		evt.LogIndex,
		contract_aid,
		evt.PoolId,
		evt.Tokens,
	)
	if err != nil {
		sw.S.Log_msg(fmt.Sprintf("DB error: can't insert into tokens_dereg table: %v\n",err))
		os.Exit(1)
	}
}
func (sw *SQLStorageWrapper) Insert_internal_balance_changed(evt *p.BalV2InternalBalanceChanged) {

	contract_aid := sw.S.Layer1_lookup_or_insert_address_id(evt.ContractAddr)
	user_aid := sw.S.Layer1_lookup_or_insert_address_id(evt.UserAddr)
	token_aid := sw.S.Layer1_lookup_or_insert_address_id(evt.TokenAddr)
	var query string
	query =  "INSERT INTO "+sw.S.SchemaName()+".ibalance("+
				"block_num,time_stamp,tx_index,log_index,contract_aid,user_aid,token_aid,delta"+
			") VALUES($1,TO_TIMESTAMP($2),$3,$4,$5,$6,$7,$8)"
	_,err := sw.S.Db().Exec(query,
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
		sw.S.Log_msg(fmt.Sprintf("DB error: can't insert into ibalance table: %v\n",err))
		os.Exit(1)
	}
}
func (sw *SQLStorageWrapper) Insert_external_balance_transfer(evt *p.BalV2ExternalBalanceTransfer) {

	contract_aid := sw.S.Layer1_lookup_or_insert_address_id(evt.ContractAddr)
	sender_aid := sw.S.Layer1_lookup_or_insert_address_id(evt.SenderAddr)
	recipient_aid := sw.S.Layer1_lookup_or_insert_address_id(evt.RecipientAddr)
	token_aid := sw.S.Layer1_lookup_or_insert_address_id(evt.TokenAddr)
	var query string
	query =  "INSERT INTO "+sw.S.SchemaName()+".ebal_transf("+
				"block_num,time_stamp,tx_index,log_index,contract_aid,"+
				"sender_aid,recipient_aid,token_aid,amount"+
			") VALUES($1,TO_TIMESTAMP($2),$3,$4,$5,$6,$7,$8,$9)"
	_,err := sw.S.Db().Exec(query,
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
		sw.S.Log_msg(fmt.Sprintf("DB error: can't insert into ebal_transf table: %v\n",err))
		os.Exit(1)
	}

}
func (sw *SQLStorageWrapper) Insert_swap(evt *p.BalV2Swap) {

	contract_aid := sw.S.Layer1_lookup_or_insert_address_id(evt.ContractAddr)
	token_in_aid := sw.S.Layer1_lookup_or_insert_address_id(evt.TokenInAddr)
	token_out_aid := sw.S.Layer1_lookup_or_insert_address_id(evt.TokenOutAddr)
	var query string
	query =  "INSERT INTO "+sw.S.SchemaName()+".swap("+
				"block_num,time_stamp,tx_index,log_index,contract_aid,"+
				"pool_id,token_in_aid,token_out_aid,amount_in,amount_out"+
			") VALUES($1,TO_TIMESTAMP($2),$3,$4,$5,$6,$7,$8,$9,$10)"
	_,err := sw.S.Db().Exec(query,
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
		sw.S.Log_msg(fmt.Sprintf("DB error: can't insert into swap table: %v\n",err))
		os.Exit(1)
	}

}
func (sw *SQLStorageWrapper) Insert_swap_fee_percentage_changed(evt *p.BalV2SwapFeePercentageChanged) {

	contract_aid := sw.S.Layer1_lookup_or_insert_address_id(evt.ContractAddr)
	var query string
	query =  "INSERT INTO "+sw.S.SchemaName()+".swap_fee("+
				"block_num,time_stamp,tx_index,log_index,contract_aid,"+
				"swap_fee"+
			") VALUES($1,TO_TIMESTAMP($2),$3,$4,$5,$6)"
	_,err := sw.S.Db().Exec(query,
		evt.BlockNum,
		evt.TimeStamp,
		evt.TxIndex,
		evt.LogIndex,
		contract_aid,
		evt.SwapFeePercentage,
	)
	if err != nil {
		sw.S.Log_msg(fmt.Sprintf("DB error: can't insert into swap_fee table: %v\n",err))
		os.Exit(1)
	}

}
func (sw *SQLStorageWrapper) Insert_pool_balance_managed(evt *p.BalV2PoolBalanceManaged) {

	contract_aid := sw.S.Layer1_lookup_or_insert_address_id(evt.ContractAddr)
	asset_manager_aid := sw.S.Layer1_lookup_or_insert_address_id(evt.AssetManagerAddr)
	token_aid := sw.S.Layer1_lookup_or_insert_address_id(evt.TokenAddr)
	var query string
	query =  "INSERT INTO "+sw.S.SchemaName()+".pool_bm("+
				"block_num,time_stamp,tx_index,log_index,contract_aid,"+
				"pool_id,asset_mgr_aid,token_aid,cash_delta,managed_delta"+
			") VALUES($1,TO_TIMESTAMP($2),$3,$4,$5,$6,$7,$8,$9,$10)"
	_,err := sw.S.Db().Exec(query,
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
		sw.S.Log_msg(fmt.Sprintf("DB error: can't insert into swap table: %v\n",err))
		os.Exit(1)
	}
}
func (sw *SQLStorageWrapper) Insert_flash_loan(evt *p.BalV2FlashLoan) {

	contract_aid := sw.S.Layer1_lookup_or_insert_address_id(evt.ContractAddr)
	token_aid := sw.S.Layer1_lookup_or_insert_address_id(evt.TokenAddr)
	recipient_aid := sw.S.Layer1_lookup_or_insert_address_id(evt.RecipientAddr)
	var query string
	query =  "INSERT INTO "+sw.S.SchemaName()+".flash_loan("+
				"block_num,time_stamp,tx_index,log_index,contract_aid,"+
				"recipient_aid,token_aid,amount,fee_amount"+
			") VALUES($1,TO_TIMESTAMP($2),$3,$4,$5,$6,$7,$8,$9)"
	_,err := sw.S.Db().Exec(query,
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
		sw.S.Log_msg(fmt.Sprintf("DB error: can't insert into flash_loan table: %v\n",err))
		os.Exit(1)
	}

}
func (sw *SQLStorageWrapper) Insert_fee_collection(evt *p.BalV2FeeCollection) {

	pool_aid := sw.S.Layer1_lookup_or_insert_address_id(evt.ContractAddr)
	var query string
	query =  "INSERT INTO "+sw.S.SchemaName()+".fee_collection("+
				"block_num,time_stamp,tx_index,log_index,pool_aid,"+
				"col_base,col_bond,rem_base,rem_bond"+
			") VALUES($1,TO_TIMESTAMP($2),$3,$4,$5,$6,$7,$8,$9)"
	_,err := sw.S.Db().Exec(query,
		evt.BlockNum,
		evt.TimeStamp,
		evt.TxIndex,
		evt.LogIndex,
		pool_aid,
		evt.CollectedBase,
		evt.CollectedBond,
		evt.RemainingBase,
		evt.RemainingBond,
	)
	if err != nil {
		sw.S.Log_msg(fmt.Sprintf("DB error: can't insert into fee_collection table: %v\n",err))
		os.Exit(1)
	}

}
func (sw *SQLStorageWrapper) Insert_swap_fee_history(rec *p.BalV2SwapHist) int64 {

	var query string
	query = "INSERT INTO "+sw.S.SchemaName()+".swf_hist("+
				"block_num,block_hash,time_stamp,tx_index,log_index,pool_aid,"+
				"pool_id,swap_fee,protocol_fee,accum_swap_fee,accum_proto_fee"+
			") VALUES($1,$2,TO_TIMESTAMP($3),$4,$5,$6,$7,$8,$9,$10,$11) "+
			"RETURNING id"

	row := sw.S.Db().QueryRow(query,
		rec.BlockNum,
		rec.BlockHash,
		rec.TimeStamp,
		rec.TxIndex,
		rec.LogIndex,
		rec.PoolAid,
		rec.PoolId,
		rec.SwapFee,
		rec.ProtocolFee,
		rec.AccumSwapFee,
		rec.AccumProtoFee,
	)
	var id int64
	err := row.Scan(&id)
	if err != nil {
		sw.S.Log_msg(
			fmt.Sprintf("DB error: can't insert into swf_hist table: %v, q=%v",err,query),
		)
		os.Exit(1)
	}
	return id
}
func (sw *SQLStorageWrapper) Insert_balance_change_history_record(rec *p.BalV2BalChg) {

	var query string
	query = "INSERT INTO "+sw.S.SchemaName()+".tok_bal("+
				"block_num,block_hash,time_stamp,tx_index,log_index,pool_aid,"+
				"pool_id,tok_aid,swf_hist_id,amount"+
			") VALUES($1,$2,TO_TIMESTAMP($3),$4,$5,$6,$7,$8,$9,$10)"
	_,err := sw.S.Db().Exec(query,
		rec.BlockNum,
		rec.BlockHash,
		rec.TimeStamp,
		rec.TxIndex,
		rec.LogIndex,
		rec.PoolAid,
		rec.PoolId,
		rec.TokenAid,
		rec.SwapHistId,
		rec.Amount,
	)
	if err != nil {
		sw.S.Log_msg(fmt.Sprintf("DB error: can't insert into tok_bal table: %v\n",err))
		os.Exit(1)
	}
}
func (sw *SQLStorageWrapper) Insert_bpt_erc20_transfer(evt *p.BalV2BPTTransfer) {

	from_aid := sw.S.Layer1_lookup_or_insert_address_id(evt.From)
	to_aid := sw.S.Layer1_lookup_or_insert_address_id(evt.To)
	var query string
	query =  "INSERT INTO "+sw.S.SchemaName()+".bpt_transf("+
				"block_num,time_stamp,tx_index,log_index,"+
				"pool_aid,from_aid,to_aid,amount"+
			") VALUES($1,TO_TIMESTAMP($2),$3,$4,$5,$6,$7,$8)"
	_,err := sw.S.Db().Exec(query,
		evt.BlockNum,
		evt.TimeStamp,
		evt.TxIndex,
		evt.LogIndex,
		evt.PoolAid,
		from_aid,
		to_aid,
		evt.Amount,
	)
	if err != nil {
		sw.S.Log_msg(fmt.Sprintf("DB error: can't insert into bpt_transf table: %v\n",err))
		os.Exit(1)
	}

}
func (sw *SQLStorageWrapper) Mark_pool_as_unhandled(rec *p.BalV2UnhandledMark) {

	pool_aid,err := sw.Lookup_pool_address_id(rec.PoolId)
	if err != nil {
		sw.S.Log_msg(fmt.Sprintf("Fatal error, can't lookup address for pool %v\n",rec.PoolId))
		os.Exit(1)
	}
	var query string
	query = "INSERT INTO "+sw.S.SchemaName()+".unhandled(pool_id,pool_aid,comments) "+
			"VALUES($1,$2,$3)"
	_,err = sw.S.Db().Exec(query,
		rec.PoolId,
		pool_aid,
		rec.Comments,
	)
	if err != nil {
		sw.S.Log_msg(fmt.Sprintf("DB error: can't insert into unhandled table: %v\n",err))
		os.Exit(1)
	}

}
