package uniswapv3

import (
	"fmt"
	"os"
	//"strings"
	//"database/sql"
	_  "github.com/lib/pq"

	p "github.com/PredictionExplorer/augur-explorer/primitives/uniswapv3"
)
func (sw *SQLStorageWrapper) Insert_pool_created(evt *p.UniV3PoolCreated) {

	contract_aid := sw.S.Layer1_lookup_or_insert_address_id(evt.ContractAddr)
	token0_aid := sw.S.Layer1_lookup_or_insert_address_id(evt.Token0Addr)
	token1_aid := sw.S.Layer1_lookup_or_insert_address_id(evt.Token1Addr)
	pool_aid := sw.S.Layer1_lookup_or_insert_address_id(evt.PoolAddr)
	var query string
	query =  "INSERT INTO "+sw.S.SchemaName()+".pool_created("+
					"block_num,time_stamp,tx_index,log_index,contract_aid,"+
					"pool_aid,token0_aid,token1_aid,fee,tick_spacing"+
					") VALUES($1,TO_TIMESTAMP($2),$3,$4,$5,$6,$7,$8,$9,$10)"
	_,err := sw.S.Db().Exec(query,
		evt.BlockNum,
		evt.TimeStamp,
		evt.TxIndex,
		evt.LogIndex,
		contract_aid,
		pool_aid,
		token0_aid,
		token1_aid,
		evt.Fee,
		evt.TickSpacing,
	)
	if err != nil {
		sw.S.Log_msg(fmt.Sprintf("DB error: can't insert into pool_created table: %v\n",err))
		os.Exit(1)
	}
}
func (sw *SQLStorageWrapper) Insert_initialize(evt *p.UniV3Initialize) {

	contract_aid := sw.S.Layer1_lookup_or_insert_address_id(evt.ContractAddr)
	var query string
	query =  "INSERT INTO "+sw.S.SchemaName()+".initialize("+
					"block_num,time_stamp,tx_index,log_index,contract_aid,"+
					"pool_aid,sqrt_pricex96,tick"+
					") VALUES($1,TO_TIMESTAMP($2),$3,$4,$5,$6,$7,$8)"
	_,err := sw.S.Db().Exec(query,
		evt.BlockNum,
		evt.TimeStamp,
		evt.TxIndex,
		evt.LogIndex,
		contract_aid,
		evt.PoolAid,
		evt.SqrtPriceX96,
		evt.Tick,
	)
	if err != nil {
		sw.S.Log_msg(fmt.Sprintf("DB error: can't insert into initialize table: %v\n",err))
		os.Exit(1)
	}
}
func (sw *SQLStorageWrapper) Insert_pool_mint(evt *p.UniV3Mint) {

	contract_aid := sw.S.Layer1_lookup_or_insert_address_id(evt.ContractAddr)
	sender_aid := sw.S.Layer1_lookup_or_insert_address_id(evt.SenderAddr)
	owner_aid := sw.S.Layer1_lookup_or_insert_address_id(evt.OwnerAddr)
	var query string
	query =  "INSERT INTO "+sw.S.SchemaName()+".mint("+
					"block_num,time_stamp,tx_index,log_index,contract_aid,"+
					"pool_aid,sender_aid,owner_aid,tick_lower,tick_upper,"+
					"amount,amount0,amount1"+
					") VALUES($1,TO_TIMESTAMP($2),$3,$4,$5,$6,$7,$8,$9,$10,$11,$12,$13)"
	_,err := sw.S.Db().Exec(query,
		evt.BlockNum,
		evt.TimeStamp,
		evt.TxIndex,
		evt.LogIndex,
		contract_aid,
		evt.PoolAid,
		sender_aid,
		owner_aid,
		evt.TickLower,
		evt.TickUpper,
		evt.Amount,
		evt.Amount0,
		evt.Amount1,
	)
	if err != nil {
		sw.S.Log_msg(fmt.Sprintf("DB error: can't insert into mint table: %v\n",err))
		os.Exit(1)
	}
}
func (sw *SQLStorageWrapper) Insert_pool_collect(evt *p.UniV3Collect) {

	contract_aid := sw.S.Layer1_lookup_or_insert_address_id(evt.ContractAddr)
	recipient_aid := sw.S.Layer1_lookup_or_insert_address_id(evt.RecipientAddr)
	owner_aid := sw.S.Layer1_lookup_or_insert_address_id(evt.OwnerAddr)
	var query string
	query =  "INSERT INTO "+sw.S.SchemaName()+".collect("+
					"block_num,time_stamp,tx_index,log_index,contract_aid,"+
					"pool_aid,recipient_aid,owner_aid,tick_lower,tick_upper,"+
					"amount0,amount1"+
					") VALUES($1,TO_TIMESTAMP($2),$3,$4,$5,$6,$7,$8,$9,$10,$11,$12)"
	_,err := sw.S.Db().Exec(query,
		evt.BlockNum,
		evt.TimeStamp,
		evt.TxIndex,
		evt.LogIndex,
		contract_aid,
		evt.PoolAid,
		recipient_aid,
		owner_aid,
		evt.TickLower,
		evt.TickUpper,
		evt.Amount0,
		evt.Amount1,
	)
	if err != nil {
		sw.S.Log_msg(fmt.Sprintf("DB error: can't insert into collect table: %v\n",err))
		os.Exit(1)
	}
}
func (sw *SQLStorageWrapper) Insert_collect_periphery(evt *p.UniV3CollectPeriphery) {

	contract_aid := sw.S.Layer1_lookup_or_insert_address_id(evt.ContractAddr)
	recipient_aid := sw.S.Layer1_lookup_or_insert_address_id(evt.RecipientAddr)
	var query string
	query =  "INSERT INTO "+sw.S.SchemaName()+".collect_nfpm("+
					"block_num,time_stamp,tx_index,log_index,contract_aid,"+
					"recipient_aid,token_id,amount0,amount1"+
					") VALUES($1,TO_TIMESTAMP($2),$3,$4,$5,$6,$7,$8,$9)"
	_,err := sw.S.Db().Exec(query,
		evt.BlockNum,
		evt.TimeStamp,
		evt.TxIndex,
		evt.LogIndex,
		contract_aid,
		recipient_aid,
		evt.TokenId,
		evt.Amount0,
		evt.Amount1,
	)
	if err != nil {
		sw.S.Log_msg(fmt.Sprintf("DB error: can't insert into collect_nfpm table: %v\n",err))
		os.Exit(1)
	}
}
func (sw *SQLStorageWrapper) Insert_pool_burn(evt *p.UniV3Burn) {

	contract_aid := sw.S.Layer1_lookup_or_insert_address_id(evt.ContractAddr)
	owner_aid := sw.S.Layer1_lookup_or_insert_address_id(evt.OwnerAddr)
	var query string
	query =  "INSERT INTO "+sw.S.SchemaName()+".burn("+
					"block_num,time_stamp,tx_index,log_index,contract_aid,"+
					"pool_aid,owner_aid,tick_lower,tick_upper,"+
					"amount,amount0,amount1"+
					") VALUES($1,TO_TIMESTAMP($2),$3,$4,$5,$6,$7,$8,$9,$10,$11,$12)"
	_,err := sw.S.Db().Exec(query,
		evt.BlockNum,
		evt.TimeStamp,
		evt.TxIndex,
		evt.LogIndex,
		contract_aid,
		evt.PoolAid,
		owner_aid,
		evt.TickLower,
		evt.TickUpper,
		evt.Amount,
		evt.Amount0,
		evt.Amount1,
	)
	if err != nil {
		sw.S.Log_msg(fmt.Sprintf("DB error: can't insert into burn table: %v\n",err))
		os.Exit(1)
	}
}
func (sw *SQLStorageWrapper) Insert_pool_swap(evt *p.UniV3Swap) {

	contract_aid := sw.S.Layer1_lookup_or_insert_address_id(evt.ContractAddr)
	sender_aid := sw.S.Layer1_lookup_or_insert_address_id(evt.SenderAddr)
	recipient_aid := sw.S.Layer1_lookup_or_insert_address_id(evt.RecipientAddr)
	var query string
	query =  "INSERT INTO "+sw.S.SchemaName()+".swap("+
					"block_num,time_stamp,tx_index,log_index,contract_aid,"+
					"pool_aid,sender_aid,recipient_aid,"+
					"amount0,amount1,fee,"+
					"sqrt_pricex96,liquidity,tick"+
					") VALUES($1,TO_TIMESTAMP($2),$3,$4,$5,$6,$7,$8,$9,$10,$11,$12,$13,$14)"
	_,err := sw.S.Db().Exec(query,
		evt.BlockNum,
		evt.TimeStamp,
		evt.TxIndex,
		evt.LogIndex,
		contract_aid,
		evt.PoolAid,
		sender_aid,
		recipient_aid,
		evt.Amount0,
		evt.Amount1,
		evt.Fee,
		evt.SqrtPriceX96,
		evt.Liquidity,
		evt.Tick,
	)
	if err != nil {
		sw.S.Log_msg(fmt.Sprintf("DB error: can't insert into swap table: %v\n",err))
		os.Exit(1)
	}
}
func (sw *SQLStorageWrapper) Insert_pool_flash(evt *p.UniV3Flash) {

	contract_aid := sw.S.Layer1_lookup_or_insert_address_id(evt.ContractAddr)
	sender_aid := sw.S.Layer1_lookup_or_insert_address_id(evt.SenderAddr)
	recipient_aid := sw.S.Layer1_lookup_or_insert_address_id(evt.RecipientAddr)
	var query string
	query =  "INSERT INTO "+sw.S.SchemaName()+".flash("+
					"block_num,time_stamp,tx_index,log_index,contract_aid,"+
					"pool_aid,sender_aid,recipient_aid,"+
					"amount0,amount1,"+
					"paid0,paid1"+
					") VALUES($1,TO_TIMESTAMP($2),$3,$4,$5,$6,$7,$8,$9,$10,$11,$12)"
	_,err := sw.S.Db().Exec(query,
		evt.BlockNum,
		evt.TimeStamp,
		evt.TxIndex,
		evt.LogIndex,
		contract_aid,
		evt.PoolAid,
		sender_aid,
		recipient_aid,
		evt.Amount0,
		evt.Amount1,
		evt.Paid0,
		evt.Paid1,
	)
	if err != nil {
		sw.S.Log_msg(fmt.Sprintf("DB error: can't insert into flash table: %v\n",err))
		os.Exit(1)
	}
}
func (sw *SQLStorageWrapper) Insert_pool_increase_observation_cardinality_next(evt *p.UniV3IncObservCardinNext) {

	contract_aid := sw.S.Layer1_lookup_or_insert_address_id(evt.ContractAddr)
	var query string
	query =  "INSERT INTO "+sw.S.SchemaName()+".iocn("+
					"block_num,time_stamp,tx_index,log_index,contract_aid,"+
					"pool_aid,next_old,next_new"+
					") VALUES($1,TO_TIMESTAMP($2),$3,$4,$5,$6,$7,$8)"
	_,err := sw.S.Db().Exec(query,
		evt.BlockNum,
		evt.TimeStamp,
		evt.TxIndex,
		evt.LogIndex,
		contract_aid,
		evt.PoolAid,
		evt.ObservationCardinalityNextOld,
		evt.ObservationCardinalityNextNew,
	)
	if err != nil {
		sw.S.Log_msg(fmt.Sprintf("DB error: can't insert into iocn table: %v\n",err))
		os.Exit(1)
	}
}
func (sw *SQLStorageWrapper) Insert_pool_set_fee_protocol(evt *p.UniV3SetFeeProtocol) {

	contract_aid := sw.S.Layer1_lookup_or_insert_address_id(evt.ContractAddr)
	var query string
	query =  "INSERT INTO "+sw.S.SchemaName()+".set_fee_proto("+
					"block_num,time_stamp,tx_index,log_index,contract_aid,"+
					"pool_aid,fee_protocol0_old,fee_protocol1_old,fee_protocol0_new,fee_protocol1_new"+
					") VALUES($1,TO_TIMESTAMP($2),$3,$4,$5,$6,$7,$8,$9,$10)"
	_,err := sw.S.Db().Exec(query,
		evt.BlockNum,
		evt.TimeStamp,
		evt.TxIndex,
		evt.LogIndex,
		contract_aid,
		evt.PoolAid,
		evt.FeeProtocol0Old,
		evt.FeeProtocol1Old,
		evt.FeeProtocol0New,
		evt.FeeProtocol1New,
	)
	if err != nil {
		sw.S.Log_msg(fmt.Sprintf("DB error: can't insert into set_fee_proto table: %v\n",err))
		os.Exit(1)
	}
}
func (sw *SQLStorageWrapper) Insert_pool_collect_protocol(evt *p.UniV3PoolCollectProtocol) {

	contract_aid := sw.S.Layer1_lookup_or_insert_address_id(evt.ContractAddr)
	sender_aid := sw.S.Layer1_lookup_or_insert_address_id(evt.SenderAddr)
	recipient_aid := sw.S.Layer1_lookup_or_insert_address_id(evt.RecipientAddr)
	var query string
	query =  "INSERT INTO "+sw.S.SchemaName()+".collect_proto("+
					"block_num,time_stamp,tx_index,log_index,contract_aid,"+
					"pool_aid,sender_aid,recipient_aid,"+
					"amount0,amount1,"+
					") VALUES($1,TO_TIMESTAMP($2),$3,$4,$5,$6,$7,$8,$9,$10)"
	_,err := sw.S.Db().Exec(query,
		evt.BlockNum,
		evt.TimeStamp,
		evt.TxIndex,
		evt.LogIndex,
		contract_aid,
		evt.PoolAid,
		sender_aid,
		recipient_aid,
		evt.Amount0,
		evt.Amount1,
	)
	if err != nil {
		sw.S.Log_msg(fmt.Sprintf("DB error: can't insert into collect_proto table: %v\n",err))
		os.Exit(1)
	}
}
func (sw *SQLStorageWrapper) Insert_increase_liquidity(evt *p.UniV3IncreaseLiquidity) {

	contract_aid := sw.S.Layer1_lookup_or_insert_address_id(evt.ContractAddr)
	var query string
	query =  "INSERT INTO "+sw.S.SchemaName()+".inc_liq("+
					"block_num,time_stamp,tx_index,log_index,contract_aid,"+
					"token_id,liquidity,amount0,amount1"+
					") VALUES($1,TO_TIMESTAMP($2),$3,$4,$5,$6,$7,$8,$9)"
	_,err := sw.S.Db().Exec(query,
		evt.BlockNum,
		evt.TimeStamp,
		evt.TxIndex,
		evt.LogIndex,
		contract_aid,
		evt.TokenId,
		evt.Liquidity,
		evt.Amount0,
		evt.Amount1,
	)
	if err != nil {
		sw.S.Log_msg(fmt.Sprintf("DB error: can't insert into inc_liq table: %v\n",err))
		os.Exit(1)
	}
}
func (sw *SQLStorageWrapper) Insert_decrease_liquidity(evt *p.UniV3DecreaseLiquidity) {

	contract_aid := sw.S.Layer1_lookup_or_insert_address_id(evt.ContractAddr)
	var query string
	query =  "INSERT INTO "+sw.S.SchemaName()+".dec_liq("+
					"block_num,time_stamp,tx_index,log_index,contract_aid,"+
					"token_id,liquidity,amount0,amount1"+
					") VALUES($1,TO_TIMESTAMP($2),$3,$4,$5,$6,$7,$8,$9)"
	_,err := sw.S.Db().Exec(query,
		evt.BlockNum,
		evt.TimeStamp,
		evt.TxIndex,
		evt.LogIndex,
		contract_aid,
		evt.TokenId,
		evt.Liquidity,
		evt.Amount0,
		evt.Amount1,
	)
	if err != nil {
		sw.S.Log_msg(fmt.Sprintf("DB error: can't insert into dec_liq table: %v\n",err))
		os.Exit(1)
	}
}
func (sw *SQLStorageWrapper) Insert_dbg_swap_loop(evt *p.UniV3DbgSwapLoop) {

	contract_aid := sw.S.Layer1_lookup_or_insert_address_id(evt.ContractAddr)
	var query string
	query =  "INSERT INTO "+sw.S.SchemaName()+".dbg_swap_loop("+
					"block_num,time_stamp,tx_index,log_index,contract_aid,"+
					"pool_aid,tick,sqrt_price,liquidity,liquidity_diff,"+
					"fee_grow_0,fee_grow_1"+
					") VALUES($1,TO_TIMESTAMP($2),$3,$4,$5,$6,$7,$8,$9,$10,$11,$12)"
	_,err := sw.S.Db().Exec(query,
		evt.BlockNum,
		evt.TimeStamp,
		evt.TxIndex,
		evt.LogIndex,
		contract_aid,
		evt.PoolAid,
		evt.Tick,
		evt.SqrtPrice,
		evt.Liquidity,
		evt.LiquidityDiff,
		evt.FeeGrow0,
		evt.FeeGrow1,
	)
	if err != nil {
		sw.S.Log_msg(fmt.Sprintf("DB error: can't insert into dbg_swap_loop table: %v\n",err))
		os.Exit(1)
	}
}
