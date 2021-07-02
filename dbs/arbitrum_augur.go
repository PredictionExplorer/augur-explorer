package dbs

import (
	"fmt"
	"os"
	//"errors"
	"database/sql"
	_  "github.com/lib/pq"

	"github.com/ethereum/go-ethereum/common"

	p "github.com/PredictionExplorer/augur-explorer/primitives"
	a "github.com/PredictionExplorer/augur-explorer/amm"

)
func (ss *SQLStorage) Get_arbitrum_augur_contract_addresses() (p.AA_ContractAddrs) {

	var query string
	query="SELECT " +
				"amm_factory,sportsball1,sportsball2,mma,trusted_factory "+
			"FROM aa_caddrs";
	row := ss.db.QueryRow(query)
	var c_addrs p.AA_ContractAddrs
	var err error
	var (
		amm_factory string
		sportsball1 string
		sportsball2 string
		mma string
		trusted_factory string
	)
	err=row.Scan(
		&amm_factory,&sportsball1,&sportsball2,&mma,&trusted_factory,
	);
	if (err!=nil) {
		if err == sql.ErrNoRows {
			ss.Log_msg(fmt.Sprintf("No contract addresses in AMM contracts table %v"))
			os.Exit(1)
		} else {
			ss.Log_msg(fmt.Sprintf("Error in Get_arbitrum_augur_contract_addresses(): %v",err))
			os.Exit(1)
		}
	}
	c_addrs.AMM_Factory=common.HexToAddress(amm_factory)
	c_addrs.SportsBall1=common.HexToAddress(sportsball1)
	c_addrs.SportsBall2=common.HexToAddress(sportsball2)
	c_addrs.MMA=common.HexToAddress(mma)
	c_addrs.TrustedFactory=common.HexToAddress(trusted_factory)
	return c_addrs
}
func (ss *SQLStorage) Get_arbitrum_augur_factory_aids(caddrs *p.AA_ContractAddrs) []int64 {

	addresses := "'" + caddrs.AMM_Factory.String() + "'," +
				"'" + caddrs.SportsBall1.String() + "'," +
				"'" + caddrs.SportsBall2.String() + "'," +
				"'" + caddrs.MMA.String() + "'"
	var query string
	query = "SELECT address_id from address WHERE addr in ("+addresses+")"
	rows,err := ss.db.Query(query)
	if (err!=nil) {
		ss.Log_msg(fmt.Sprintf("DB error: %v (query=%v)",err,query))
		os.Exit(1)
	}
	records := make([]int64,0,32)

	defer rows.Close()
	for rows.Next() {
		var aid int64
		err=rows.Scan(&aid)
		if err!=nil {
			ss.Log_msg(fmt.Sprintf("DB error: %v, q=%v",err,query))
			os.Exit(1)
		}
		records = append(records,aid)
	}
	return records
}
func (ss *SQLStorage) Update_arbitrum_augur_process_status(status *p.ArbitrumAugurProcessStatus) {

	var query string
	query = "UPDATE aa_proc_status SET last_evt_id = $1"

	_,err := ss.db.Exec(query,status.LastEvtId)
	if (err!=nil) {
		ss.Log_msg(fmt.Sprintf("DB error: %v q=%v",err,query))
		os.Exit(1)
	}
}
func (ss *SQLStorage) Insert_aa_pool_created_event(evt *p.AA_PoolCreated) {


	contract_aid:=ss.Lookup_or_create_address(evt.Contract,evt.BlockNum,evt.TxId)
	pool_aid:=ss.Lookup_or_create_address(evt.PoolAddr,evt.BlockNum,evt.TxId)
	factory_aid:=ss.Lookup_or_create_address(evt.FactoryAddr,evt.BlockNum,evt.TxId)
	creator_aid:=ss.Lookup_or_create_address(evt.CreatorAddr,evt.BlockNum,evt.TxId)
	tokrcpt_aid:=ss.Lookup_or_create_address(evt.TokenRecipientAddr,evt.BlockNum,evt.TxId)

	var query string
	query = "INSERT INTO aa_pool_created(" +
				"evtlog_id,block_num,tx_id,contract_aid,time_stamp,"+
				"pool_aid,factory_aid,creator_aid,market_id,token_rcpt_aid" +
			") VALUES ($1,$2,$3,$4,TO_TIMESTAMP($5),$6,$7,$8,$9,$10)"

	_,err := ss.db.Exec(query,
			evt.EvtId,
			evt.BlockNum,
			evt.TxId,
			contract_aid,
			evt.TimeStamp,
			pool_aid,
			factory_aid,
			creator_aid,
			evt.MarketId,
			tokrcpt_aid,
	)
	if err != nil {
		ss.Log_msg(fmt.Sprintf("DB error: can't insert into aa_pool_created table: %v; q=%v",err,query))
		os.Exit(1)
	}
}
func (ss *SQLStorage) Insert_aa_liquidity_changed_event(evt *p.AA_LiquidityChanged) {

	contract_aid:=ss.Lookup_or_create_address(evt.Contract,evt.BlockNum,evt.TxId)
	factory_aid:=ss.Lookup_or_create_address(evt.MarketFactoryAddr,evt.BlockNum,evt.TxId)
	user_aid:=ss.Lookup_or_create_address(evt.UserAddr,evt.BlockNum,evt.TxId)
	recipient_aid:=ss.Lookup_or_create_address(evt.RecipientAddr,evt.BlockNum,evt.TxId)
	var query string
	query = "INSERT INTO aa_liquidity_changed (" +
				"evtlog_id,block_num,tx_id,contract_aid,time_stamp,"+
				"market_id,factory_aid,user_aid,recipient_aid,collateral,lp_tokens,shares_returned" +
			") VALUES ($1,$2,$3,$4,TO_TIMESTAMP($5),$6,$7,$8,$9,$10::DECIMAL/1e+6,$11::DECIMAL/1e+18,$12)"

	_,err := ss.db.Exec(query,
			evt.EvtId,
			evt.BlockNum,
			evt.TxId,
			contract_aid,
			evt.TimeStamp,
			evt.MarketId,
			factory_aid,
			user_aid,
			recipient_aid,
			evt.Collateral,
			evt.LpTokens,
			evt.SharesReturned,
	)
	if err != nil {
		ss.Log_msg(fmt.Sprintf("DB error: can't insert into aa_liquidity_changed table: %v; q=%v",err,query))
		os.Exit(1)
	}
}
func (ss *SQLStorage) Get_arbitrum_augur_processing_status() p.ArbitrumAugurProcessStatus {

	var output p.ArbitrumAugurProcessStatus
	var null_id sql.NullInt64

	var query string
	for {
		query = "SELECT last_evt_id FROM aa_proc_status"

		res := ss.db.QueryRow(query)
		err := res.Scan(&null_id)
		if (err!=nil) {
			if err == sql.ErrNoRows {
				query = "INSERT INTO aa_proc_status DEFAULT VALUES"
				_,err := ss.db.Exec(query)
				if (err!=nil) {
					ss.Log_msg(fmt.Sprintf("DB error: %v q=%v",err,query))
					os.Exit(1)
				}
			} else {
				ss.Log_msg(fmt.Sprintf("DB error: %v q=%v",err,query))
				os.Exit(1)
			}
		} else {
			break
		}
	}
	if null_id.Valid {
		output.LastEvtId = null_id.Int64
	}
	return output
}
func (ss *SQLStorage) Insert_aa_price_market_event(evt *p.AA_PriceMarket) {

	contract_aid:=ss.Lookup_or_create_address(evt.Contract,evt.BlockNum,evt.TxId)
	creator_aid:=ss.Lookup_or_create_address(evt.CreatorAddr,evt.BlockNum,evt.TxId)
	var query string
	query = "INSERT INTO aa_price_market (" +
				"evtlog_id,block_num,tx_id,contract_aid,time_stamp,"+
				"creator_aid,end_time,spot_price" +
				") VALUES ($1,$2,$3,$4,TO_TIMESTAMP($5),$6,TO_TIMESTAMP($7),$8::DECIMAL/1e+18)"

	_,err := ss.db.Exec(query,
			evt.EvtId,
			evt.BlockNum,
			evt.TxId,
			contract_aid,
			evt.TimeStamp,
			creator_aid,
			evt.EndTime,
			evt.SpotPrice,
	)
	if err != nil {
		ss.Log_msg(fmt.Sprintf("DB error: can't insert into aa_price_market table: %v; q=%v",err,query))
		os.Exit(1)
	}
}
func (ss *SQLStorage) Insert_aa_sports_market_event(evt *p.AA_SportsMarket) {

	contract_aid:=ss.Lookup_or_create_address(evt.Contract,evt.BlockNum,evt.TxId)
	creator_aid:=ss.Lookup_or_create_address(evt.CreatorAddr,evt.BlockNum,evt.TxId)
	collateral_aid:=ss.Lookup_or_create_address(evt.CollateralAddr,evt.BlockNum,evt.TxId)
	settlement_aid:=ss.Lookup_or_create_address(evt.SettlementAddr,evt.BlockNum,evt.TxId)
	feepot_aid:=ss.Lookup_or_create_address(evt.FeePotAddr,evt.BlockNum,evt.TxId)
	protocol_aid:=ss.Lookup_or_create_address(evt.ProtocolAddr,evt.BlockNum,evt.TxId)
	var query string
	query = "SELECT amm_insert_sports_market("+
				"$1::BIGINT," + // evtlog_id
				"$2::BIGINT," + // block_num
				"$3::BIGINT," + // tx_id
				"$4::BIGINT," + // contract_aid
				"TO_TIMESTAMP($5)," + // time_stamp
				"$6::BIGINT," + // market_id
				"$7::BIGINT," + // creator_aid
				"TO_TIMESTAMP($8)," + // created_time
				"TO_TIMESTAMP($9)," + // end_time
				"$10::DECIMAL," + // settlement_fee
				"$11::DECIMAL,"+ // staker_fee
				"$12::DECIMAL,"+ // protocol fee
				"$13::BIGINT,"+ // settlement_aid
				"$14::BIGINT,"+ // feepot_aid
				"$15::BIGINT,"+ // protocol_aid
				"$16::BIGINT,"+ // collateral_aid
				"$17::DECIMAL,"+// sharefactor
				"$18::TEXT,"+ // sharetokens (comma separated)
				"$19::BIGINT,"+ // event_id (MMA event code)
				"$20::BIGINT,"+ // home_team_id
				"$21::BIGINT,"+ // away_team_id
				"TO_TIMESTAMP($22),"+ // estimated_start
				"$23::INT,"+ // market_type
				"$24::DECIMAL"+ // value0 (score)
			")"

	_,err := ss.db.Exec(query,
			evt.EvtId,
			evt.BlockNum,
			evt.TxId,
			contract_aid,
			evt.TimeStamp,
			evt.MarketId,
			creator_aid,
			evt.EstimatedStarTime,
			evt.EndTime,
			evt.SettlementFee,
			evt.StakerFee,
			evt.ProtocolFee,
			settlement_aid,
			feepot_aid,
			protocol_aid,
			collateral_aid,
			evt.ShareFactor,
			evt.ShareTokens,
			evt.EventId,
			evt.HomeTeamId,
			evt.AwayTeamId,
			evt.EstimatedStarTime,
			evt.MarketType,
			evt.Score,
	)
	if err != nil {
		ss.Log_msg(fmt.Sprintf("DB error: can't insert into aa_sports_market table: %v; q=%v",err,query))
		os.Exit(1)
	}
}
func (ss *SQLStorage) Insert_aa_trusted_market_event(evt *p.AA_TrustedMarket) {

	contract_aid:=ss.Lookup_or_create_address(evt.Contract,evt.BlockNum,evt.TxId)
	creator_aid:=ss.Lookup_or_create_address(evt.CreatorAddr,evt.BlockNum,evt.TxId)
	var query string
	query = "INSERT INTO aa_trusted_market (" +
				"evtlog_id,block_num,tx_id,contract_aid,time_stamp,"+
				"market_id,end_time,creator_aid,description,outcomes" +
				") VALUES ("+
					"$1,$2,$3,$4,TO_TIMESTAMP($5),$6,TO_TIMESTAMP($7),$8,$9,$10)"

	_,err := ss.db.Exec(query,
			evt.EvtId,
			evt.BlockNum,
			evt.TxId,
			contract_aid,
			evt.TimeStamp,
			evt.MarketId,
			evt.EndTime,
			creator_aid,
			evt.Description,
			evt.Outcomes,
	)
	if err != nil {
		ss.Log_msg(fmt.Sprintf("DB error: can't insert into aa_trusted_market table: %v; q=%v",err,query))
		os.Exit(1)
	}
}
func (ss *SQLStorage) Insert_aa_shares_minted_event(evt *p.AA_SharesMinted) {

	contract_aid:=ss.Lookup_or_create_address(evt.Contract,evt.BlockNum,evt.TxId)
	aid:=ss.Lookup_or_create_address(evt.ReceiverAddr,evt.BlockNum,evt.TxId)
	var query string
	query = "INSERT INTO aa_shares_minted (" +
				"evtlog_id,block_num,tx_id,contract_aid,time_stamp,"+
				"aid,market_id,amount" +
			") VALUES ($1,$2,$3,$4,TO_TIMESTAMP($5),$6,$7,$8::DECIMAL/1e+18)"

	_,err := ss.db.Exec(query,
			evt.EvtId,
			evt.BlockNum,
			evt.TxId,
			contract_aid,
			evt.TimeStamp,
			aid,
			evt.MarketId,
			evt.Amount,
	)
	if err != nil {
		ss.Log_msg(fmt.Sprintf("DB error: can't insert into aa_shares_minted table: %v; q=%v",err,query))
		os.Exit(1)
	}
}
func (ss *SQLStorage) Insert_aa_shares_burned_event(evt *p.AA_SharesBurned) {

	contract_aid:=ss.Lookup_or_create_address(evt.Contract,evt.BlockNum,evt.TxId)
	aid:=ss.Lookup_or_create_address(evt.ReceiverAddr,evt.BlockNum,evt.TxId)
	var query string
	query = "INSERT INTO aa_shares_burned(" +
				"evtlog_id,block_num,tx_id,contract_aid,time_stamp,"+
				"aid,market_id,amount" +
			") VALUES ($1,$2,$3,$4,TO_TIMESTAMP($5),$6,$7,$8::DECIMAL/1e+18)"

	_,err := ss.db.Exec(query,
			evt.EvtId,
			evt.BlockNum,
			evt.TxId,
			contract_aid,
			evt.TimeStamp,
			aid,
			evt.MarketId,
			evt.Amount,
	)
	if err != nil {
		ss.Log_msg(fmt.Sprintf("DB error: can't insert into aa_shares_burned table: %v; q=%v",err,query))
		os.Exit(1)
	}
}
func (ss *SQLStorage) Insert_aa_shares_swapped_event(evt *p.AA_SharesSwapped) {

	contract_aid:=ss.Lookup_or_create_address(evt.Contract,evt.BlockNum,evt.TxId)
	user_aid:=ss.Lookup_or_create_address(evt.UserAddr,evt.BlockNum,evt.TxId)
	factory_aid:=ss.Lookup_or_create_address(evt.MarketFactoryAddr,evt.BlockNum,evt.TxId)
	var query string
	query = "INSERT INTO aa_shares_swapped (" +
				"evtlog_id,block_num,tx_id,contract_aid,time_stamp,"+
				"market_id,factory_aid,user_aid,outcome_idx,collateral,shares" +
			") VALUES ($1,$2,$3,$4,TO_TIMESTAMP($5),$6,$7,$8,$9,$10::DECIMAL/1e+6,$11::DECIMAL/1e+18)"

	_,err := ss.db.Exec(query,
			evt.EvtId,
			evt.BlockNum,
			evt.TxId,
			contract_aid,
			evt.TimeStamp,
			evt.MarketId,
			factory_aid,
			user_aid,
			evt.Outcome,
			evt.Collateral,
			evt.Shares,
	)
	if err != nil {
		ss.Log_msg(fmt.Sprintf("DB error: can't insert into aa_shares_swapped table: %v; q=%v",err,query))
		os.Exit(1)
	}
}
func (ss *SQLStorage) Insert_aa_settlement_fee_claimed_event(evt *p.AA_SettlementFeeClaimed) {

	contract_aid:=ss.Lookup_or_create_address(evt.Contract,evt.BlockNum,evt.TxId)
	settlement_aid:=ss.Lookup_or_create_address(evt.SettlementAddr,evt.BlockNum,evt.TxId)
	receiver_aid:=ss.Lookup_or_create_address(evt.ReceiverAddr,evt.BlockNum,evt.TxId)
	var query string
	query = "INSERT INTO aa_sfee_claimed (" +
				"evtlog_id,block_num,tx_id,contract_aid,time_stamp,"+
				"settlement_aid,receiver_aid,amount" +
			") VALUES ($1,$2,$3,$4,TO_TIMESTAMP($5),$6,$7,$8::DECIMAL/1e+6)"

	_,err := ss.db.Exec(query,
			evt.EvtId,
			evt.BlockNum,
			evt.TxId,
			contract_aid,
			evt.TimeStamp,
			settlement_aid,
			receiver_aid,
			evt.Amount,
	)
	if err != nil {
		ss.Log_msg(fmt.Sprintf("DB error: can't insert into aa_sfee_claimed table: %v; q=%v",err,query))
		os.Exit(1)
	}
}
func (ss *SQLStorage) Insert_aa_protocol_fee_claimed_event(evt *p.AA_ProtocolFeeClaimed) {

	contract_aid:=ss.Lookup_or_create_address(evt.Contract,evt.BlockNum,evt.TxId)
	protocol_aid:=ss.Lookup_or_create_address(evt.ProtocolAddr,evt.BlockNum,evt.TxId)
	var query string
	query = "INSERT INTO aa_pfee_claimed (" +
				"evtlog_id,block_num,tx_id,contract_aid,time_stamp,"+
				"protocol_aid,amount" +
			") VALUES ($1,$2,$3,$4,TO_TIMESTAMP($5),$6,$7::DECIMAL/1e+6)"

	_,err := ss.db.Exec(query,
			evt.EvtId,
			evt.BlockNum,
			evt.TxId,
			contract_aid,
			evt.TimeStamp,
			protocol_aid,
			evt.Amount,
	)
	if err != nil {
		ss.Log_msg(fmt.Sprintf("DB error: can't insert into aa_pfee_claimed table: %v; q=%v",err,query))
		os.Exit(1)
	}
}
func (ss *SQLStorage) Insert_aa_protocol_changed_event(evt *p.AA_ProtocolChanged) {

	contract_aid:=ss.Lookup_or_create_address(evt.Contract,evt.BlockNum,evt.TxId)
	protocol_aid:=ss.Lookup_or_create_address(evt.ProtocolAddr,evt.BlockNum,evt.TxId)
	var query string
	query = "INSERT INTO aa_proto_chg (" +
				"evtlog_id,block_num,tx_id,contract_aid,time_stamp,"+
				"protocol_aid" +
			") VALUES ($1,$2,$3,$4,TO_TIMESTAMP($5),$6)"

	_,err := ss.db.Exec(query,
			evt.EvtId,
			evt.BlockNum,
			evt.TxId,
			contract_aid,
			evt.TimeStamp,
			protocol_aid,
	)
	if err != nil {
		ss.Log_msg(fmt.Sprintf("DB error: can't insert into aa_protocol table: %v; q=%v",err,query))
		os.Exit(1)
	}
}
func (ss *SQLStorage) Insert_aa_protocol_fee_changed_event(evt *p.AA_ProtocolFeeChanged) {

	contract_aid:=ss.Lookup_or_create_address(evt.Contract,evt.BlockNum,evt.TxId)
	var query string
	query = "INSERT INTO aa_pfee_chg (" +
				"evtlog_id,block_num,tx_id,contract_aid,time_stamp,"+
				"protocol_fee" +
			") VALUES ($1,$2,$3,$4,TO_TIMESTAMP($5),$6::DECIMAL/1e+18)"

	_,err := ss.db.Exec(query,
			evt.EvtId,
			evt.BlockNum,
			evt.TxId,
			contract_aid,
			evt.TimeStamp,
			evt.Fee,
	)
	if err != nil {
		ss.Log_msg(fmt.Sprintf("DB error: can't insert into aa_pfee_chg table: %v; q=%v",err,query))
		os.Exit(1)
	}
}
func (ss *SQLStorage) Insert_aa_settlement_fee_changed_event(evt *p.AA_SettlementFeeChanged) {

	contract_aid:=ss.Lookup_or_create_address(evt.Contract,evt.BlockNum,evt.TxId)
	var query string
	query = "INSERT INTO aa_sfee_chg (" +
				"evtlog_id,block_num,tx_id,contract_aid,time_stamp,"+
				"settlement_fee" +
			") VALUES ($1,$2,$3,$4,TO_TIMESTAMP($5),$6::DECIMAL/1e+18)"

	_,err := ss.db.Exec(query,
			evt.EvtId,
			evt.BlockNum,
			evt.TxId,
			contract_aid,
			evt.TimeStamp,
			evt.Fee,
	)
	if err != nil {
		ss.Log_msg(fmt.Sprintf("DB error: can't insert into aa_pfee_chg table: %v; q=%v",err,query))
		os.Exit(1)
	}
}
func (ss *SQLStorage) Insert_aa_staker_fee_changed_event(evt *p.AA_StakerFeeChanged) {

	contract_aid:=ss.Lookup_or_create_address(evt.Contract,evt.BlockNum,evt.TxId)
	var query string
	query = "INSERT INTO aa_stk_fee_chg (" +
				"evtlog_id,block_num,tx_id,contract_aid,time_stamp,"+
				"settlement_fee" +
			") VALUES ($1,$2,$3,$4,TO_TIMESTAMP($5),$6::DECIMAL/1e+18)"

	_,err := ss.db.Exec(query,
			evt.EvtId,
			evt.BlockNum,
			evt.TxId,
			contract_aid,
			evt.TimeStamp,
			evt.Fee,
	)
	if err != nil {
		ss.Log_msg(fmt.Sprintf("DB error: can't insert into aa_stk_fee_chg table: %v; q=%v",err,query))
		os.Exit(1)
	}
}
func (ss *SQLStorage) Insert_aa_winnings_claimed_event(evt *p.AA_WinningsClaimed) {

	contract_aid:=ss.Lookup_or_create_address(evt.Contract,evt.BlockNum,evt.TxId)
	win_outc_aid:=ss.Lookup_or_create_address(evt.WinningOutcomeAddr,evt.BlockNum,evt.TxId)
	receiver_aid:=ss.Lookup_or_create_address(evt.ReceiverAddr,evt.BlockNum,evt.TxId)
	var query string
	query = "INSERT INTO aa_winclaim(" +
				"evtlog_id,block_num,tx_id,contract_aid,time_stamp,"+
				"market_id,win_outc_aid,receiver_aid,amount,settlement_fee,payout" +
			") VALUES (" +
				"$1,$2,$3,$4,TO_TIMESTAMP($5),"+
				"$6,$7,$8,$9::DECIMAL/1e+18,$10::DECIMAL/1e+18,$11::DECIMAL/1e+18"+
			")"

	_,err := ss.db.Exec(query,
			evt.EvtId,
			evt.BlockNum,
			evt.TxId,
			contract_aid,
			evt.TimeStamp,
			evt.MarketId,
			win_outc_aid,
			receiver_aid,
			evt.Amount,
			evt.SettlementFee,
			evt.Payout,
	)
	if err != nil {
		ss.Log_msg(fmt.Sprintf("DB error: can't insert into aa_winclaim table: %v; q=%v",err,query))
		os.Exit(1)
	}
}
func (ss *SQLStorage) Insert_aa_market_resolved_event(evt *p.AA_MarketResolved) {

	contract_aid:=ss.Lookup_or_create_address(evt.Contract,evt.BlockNum,evt.TxId)
	winner_aid:=ss.Lookup_or_create_address(evt.WinnerAddr,evt.BlockNum,evt.TxId)
	var query string
	query = "INSERT INTO aa_mkt_resolved (" +
				"evtlog_id,block_num,tx_id,contract_aid,time_stamp,"+
				"market_id,winner_aid" +
			") VALUES (" +
				"$1,$2,$3,$4,TO_TIMESTAMP($5),"+
				"$6,$7"+
			")"

	_,err := ss.db.Exec(query,
			evt.EvtId,
			evt.BlockNum,
			evt.TxId,
			contract_aid,
			evt.TimeStamp,
			evt.MarketId,
			winner_aid,
	)
	if err != nil {
		ss.Log_msg(fmt.Sprintf("DB error: can't insert into aa_mkt_resolved table: %v; q=%v",err,query))
		os.Exit(1)
	}
}
func (ss *SQLStorage) Get_arbitrum_augur_pools() []p.AA_Pool {

	var query string
	query = "SELECT " +
				"EXTRACT(EPOCH FROM time_stamp)::BIGINT AS created_ts, " +
				"time_stamp," +
				"p.block_num, " +
				"tx.tx_hash," +
				"pa.addr," +
				"fa.addr," +
				"ca.addr," +
				"market_id " +
			"FROM aa_pool_created AS p " +
				"LEFT JOIN address pa ON p.pool_aid=pa.address_id " +
				"LEFT JOIN address fa ON p.factory_aid=fa.address_id " +
				"LEFT JOIN address ca ON p.creator_aid=ca.address_id " +
				"JOIN transaction tx ON p.tx_id=tx.id " +
			"ORDER BY p.time_stamp"

	rows,err := ss.db.Query(query)
	if (err!=nil) {
		ss.Log_msg(fmt.Sprintf("DB error: %v (query=%v)",err,query))
		os.Exit(1)
	}
	records := make([]p.AA_Pool,0,32)

	defer rows.Close()
	for rows.Next() {
		var rec p.AA_Pool
		err=rows.Scan(
			&rec.TimeStamp,
			&rec.DateTime,
			&rec.BlockNum,
			&rec.TxHash,
			&rec.PoolAddr,
			&rec.FactoryAddr,
			&rec.CreatorAddr,
			&rec.MarketId,
		)
		if err!=nil {
			ss.Log_msg(fmt.Sprintf("DB error: %v, q=%v",err,query))
			os.Exit(1)
		}
		records = append(records,rec)
	}
	return records

}
func (ss *SQLStorage) Insert_aa_feepot_transfer_event(evt *p.AA_FeePotTransfer) {

	contract_aid:=ss.Lookup_or_create_address(evt.Contract,evt.BlockNum,evt.TxId)
	from_aid:=ss.Lookup_or_create_address(evt.From,evt.BlockNum,evt.TxId)
	to_aid:=ss.Lookup_or_create_address(evt.To,evt.BlockNum,evt.TxId)
	var query string
	query = "INSERT INTO aa_feepot_trsf(" +
				"evtlog_id,block_num,tx_id,contract_aid,time_stamp,"+
				"from_aid,to_aid,value" +
			") VALUES ($1,$2,$3,$4,TO_TIMESTAMP($5),$6,$7,$8::DECIMAL/1e+18)"

	_,err := ss.db.Exec(query,
			evt.EvtId,
			evt.BlockNum,
			evt.TxId,
			contract_aid,
			evt.TimeStamp,
			from_aid,
			to_aid,
			evt.Value,
	)
	if err != nil {
		ss.Log_msg(fmt.Sprintf("DB error: can't insert into aa_sets_burned table: %v; q=%v",err,query))
		os.Exit(1)
	}
}
func (ss *SQLStorage) Is_feepot(addr string) bool {

	var query string
	query = "SELECT feepot_aid FROM aa_new_hatchery h "+
			"JOIN address a ON h.feepot_aid=a.address_id "+
			"WHERE a.addr=$1"+
			"LIMIT 1"
	row := ss.db.QueryRow(query,addr)
	var null_id sql.NullInt64
	err := row.Scan(&null_id)
	if (err!=nil) {
		if err==sql.ErrNoRows {
			return false
		}
		ss.Log_msg(fmt.Sprintf("Error in Is_feepot(): %v",err))
		os.Exit(1)
	}
	_=null_id
	return true
}
func (ss *SQLStorage) Get_markets() {


}
func (ss *SQLStorage) Get_sport_markets(status,sort int64,offset,limit int,constants *p.AMM_Constants,contract_aids []int64) (int64,[]p.AMM_SportMarket) {

	where_condition := " AND r.id IS NULL " // Open market
	if status == 1 {
		where_condition = " AND r.id IS NOT NULL " //Resolved market
	}
	records := make([]p.AMM_SportMarket,0,32)
	if len(contract_aids)==0 {
		return 0,records
	}
	var contract_aids_str string = fmt.Sprintf("%v",contract_aids[0])
	for i:=1 ; i<len(contract_aids); i++ {
		contract_aids_str = contract_aids_str + fmt.Sprintf(",%v",contract_aids[i])
	}

	var query string

	query = "SELECT count(*) AS total " +
			"FROM aa_sports_market AS m " +
			"LEFT JOIN aa_mkt_resolved r ON (m.contract_aid=r.contract_aid AND m.market_id=r.market_id) "+
			"WHERE m.contract_aid IN ("+contract_aids_str+") " + where_condition
	row := ss.db.QueryRow(query)
	var null_counter sql.NullInt64
	err := row.Scan(&null_counter)
	if (err!=nil) {
		if err==sql.ErrNoRows {
			return 0,records
		}
		ss.Log_msg(fmt.Sprintf("Error in Get_sport_markets(): %v",err))
		os.Exit(1)
	}
	total_rows := null_counter.Int64

	query = "SELECT " +
				"EXTRACT(EPOCH FROM m.time_stamp)::BIGINT AS created_ts, " +
				"m.time_stamp," +
				"m.block_num, " +
				"tx.tx_hash," +
				"m.market_id," +
				"m.contract_aid," +
				"m.contract_aid," +
				"ca.addr," +
				"fa.addr," +
				"EXTRACT(EPOCH FROM m.start_time)::BIGINT AS start_time_ts, " +
				"EXTRACT(EPOCH FROM m.end_time)::BIGINT AS end_time_ts, " +
				"m.start_time," +
				"m.end_time," +
				"m.event_id," +
				"m.home_team_id," +
				"m.away_team_id," +
				"m.score," +
				"m.market_type, " +
				"r.id resolved_id, "+
				"r.winner_aid, " +
				"EXTRACT(EPOCH FROM r.time_stamp)::BIGINT AS resolved_ts, " +
				"r.time_stamp resolved_date, " +
				"m.liquidity "+
			"FROM aa_sports_market AS m " +
				"LEFT JOIN address ca ON m.creator_aid=ca.address_id " +
				"LEFT JOIN address fa ON m.contract_aid=fa.address_id " +
				"LEFT JOIN aa_mkt_resolved r ON (m.contract_aid=r.contract_aid AND m.market_id=r.market_id) "+
				"JOIN transaction tx ON m.tx_id=tx.id " +
			"WHERE m.contract_aid IN(" + contract_aids_str + ") "+
			where_condition +
			"ORDER BY m.time_stamp " +
			"OFFSET $1 LIMIT $2"
	fmt.Printf("query = %v\n",query)
	rows,err := ss.db.Query(query,offset,limit)
	if (err!=nil) {
		ss.Log_msg(fmt.Sprintf("DB error: %v (query=%v)",err,query))
		os.Exit(1)
	}

	defer rows.Close()
	for rows.Next() {
		var rec p.AMM_SportMarket
		var null_resolved_id,null_resolved_ts,null_winner_aid sql.NullInt64
		var null_resolved_date sql.NullString
		err=rows.Scan(
			&rec.CreatedTs,
			&rec.CreatedDate,
			&rec.BlockNum,
			&rec.TxHash,
			&rec.MarketId,
			&rec.ContractAid,
			&rec.FactoryAid,
			&rec.CreatorAddr,
			&rec.FactoryAddr,
			&rec.StartTimeTs,
			&rec.EndTimeTs,
			&rec.StartTime,
			&rec.EndTime,
			&rec.EventId,
			&rec.HomeTeamId,
			&rec.AwayTeamId,
			&rec.Score,
			&rec.MarketTypeCode,
			&null_resolved_id,
			&null_winner_aid,
			&null_resolved_ts,
			&null_resolved_date,
			&rec.Liquidity,
		)
		if err!=nil {
			ss.Log_msg(fmt.Sprintf("DB error: %v, q=%v",err,query))
			os.Exit(1)
		}
		team,exists := constants.Teams[rec.HomeTeamId]
		if exists {
			rec.HomeTeam = team.Name + " " +team.Mascot
		}
		team,exists = constants.Teams[rec.AwayTeamId]
		if exists {
			rec.AwayTeam = team.Name + " " + team.Mascot
		}
		sport_id := a.Get_sport_id_from_team(constants,rec.HomeTeamId)
		title,description := a.Get_market_title(sport_id,rec.HomeTeam,rec.AwayTeam,rec.MarketTypeCode,1)
		/*fmt.Printf(
			"sport_id=%v, Home(id=%v)=%v, Away(id=%v)=%v, title=%v, descr=%v\n",
			sport_id,rec.HomeTeamId,rec.HomeTeam,rec.AwayTeamId,rec.AwayTeam,title,description,
		)*/
		rec.Title = title
		rec.Description = description

		if null_resolved_id.Valid {
			rec.ResolvedTs = null_resolved_ts.Int64
			rec.ResolvedDate = null_resolved_date.String
		}
		records = append(records,rec)
	}
	return total_rows,records

}
func (ss *SQLStorage) Get_sport_market_info(constants *p.AMM_Constants,contract_aid,market_id int64) (p.API_AMM_SportsMarket,error) {

	var rec p.API_AMM_SportsMarket
	var query string
	query = "SELECT " +
				"m.block_ts, " +
				"m.block_datetime," +
				"m.block_num, " +
				"tx.tx_hash,"+
				"s.contract_aid,"+
				"m.factory_aid,"+
				"fa.addr," +
				"m.created_ts, "+
				"m.created_time_date," +
				"m.end_time_ts,"+
				"m.end_time," +
				"m.market_id," +
				"m.sharefactor," +
				"m.settlement_fee,"+
				"m.staker_fee,"+
				"m.protocol_fee,"+
				"m.settl_addr,"+
				"m.proto_addr,"+
				"m.feepot_addr,"+
				"s.creator_aid,"+
				"EXTRACT(EPOCH FROM s.est_start_time)::BIGINT AS est_start_time_ts, " +
				"s.est_start_time," +
				"s.event_id," +
				"s.home_team_id," +
				"s.away_team_id," +
				"s.value0," +
				"s.market_type, " +
				"m.liquidity, " +
				"coll_addr, " +
				"win_addr, " +
				"ca.addr " +
			"FROM aa_sports_market AS s " +
				"LEFT JOIN LATERAL (" +
					"SELECT " +
						"EXTRACT(EPOCH FROM m.time_stamp)::BIGINT AS block_ts, " +
						"m.time_stamp as block_datetime," +
						"m.block_num, " +
						"EXTRACT(EPOCH FROM m.created_time)::BIGINT AS created_ts, "+
						"m.created_time as created_time_date," +
						"EXTRACT(EPOCH FROM m.end_time)::BIGINT AS end_time_ts,"+
						"m.end_time," +
						"m.market_id," +
						"m.factory_aid,"+
						"m.sharefactor," +
						"m.settlement_fee/1e+18 AS settlement_fee,"+
						"m.staker_fee/1e+18 AS staker_fee,"+
						"m.protocol_fee/1e+18 AS protocol_fee,"+
						"settl_addr.addr settl_addr,"+
						"proto_addr.addr proto_addr,"+
						"feepot_addr.addr feepot_addr,"+
						"coll_addr.addr coll_addr," +
						"win_addr.addr win_addr, " +
						"liquidity " +
					"FROM aa_market m "+
						"LEFT JOIN address AS settl_addr ON settl_addr.address_id=m.settlement_aid " +
						"LEFT JOIN address AS proto_addr ON proto_addr.address_id=m.protocol_aid " +
						"LEFT JOIN address AS feepot_addr ON feepot_addr.address_id=m.feepot_aid " +
						"LEFT JOIN address AS coll_addr ON coll_addr.address_id=m.collateral_aid " +
						"LEFT JOIN address AS win_addr ON win_addr.address_id=m.winner_aid " +
				") AS m ON m.market_id=s.market_id AND m.factory_aid=s.contract_aid " +
				"LEFT JOIN address ca ON s.creator_aid=ca.address_id " +
				"LEFT JOIN address fa ON s.contract_aid=fa.address_id " +
				"JOIN transaction tx ON s.tx_id=tx.id " +
			"WHERE s.market_id=$1 AND s.contract_aid=$2"

	row := ss.db.QueryRow(query,market_id,contract_aid)
	var win_addr sql.NullString
	err := row.Scan(
			&rec.AbstractMarketInfo.BlockTimeStamp,
			&rec.AbstractMarketInfo.BlockDateTime,
			&rec.AbstractMarketInfo.BlockNum,
			&rec.AbstractMarketInfo.TxHash,
			&rec.AbstractMarketInfo.ContractAid,
			&rec.AbstractMarketInfo.FactoryAid,
			&rec.AbstractMarketInfo.FactoryAddr,
			&rec.AbstractMarketInfo.MarketCreatedTs,
			&rec.AbstractMarketInfo.MarketCreatedDate,
			&rec.AbstractMarketInfo.MarketEndTimeTs,
			&rec.AbstractMarketInfo.MarketEndDate,
			&rec.AbstractMarketInfo.MarketId,
			&rec.AbstractMarketInfo.ShareFactor,
			&rec.AbstractMarketInfo.SettlementFee,
			&rec.AbstractMarketInfo.StakerFee,
			&rec.AbstractMarketInfo.ProtocolFee,
			&rec.AbstractMarketInfo.SettlementAddr,
			&rec.AbstractMarketInfo.ProtocolAddr,
			&rec.AbstractMarketInfo.FeePotAddr,
			&rec.CreatorAid,
			&rec.EstimatedStartTs,
			&rec.EstimatedStartDate,
			&rec.EventId,
			&rec.HomeTeamId,
			&rec.AwayTeamId,
			&rec.Score,
			&rec.MarketTypeCode,
			&rec.Liquidity,
			&rec.AbstractMarketInfo.CollateralAddr,
			&win_addr,
			&rec.AbstractMarketInfo.ContractAddr,
	)
	if win_addr.Valid {
		rec.AbstractMarketInfo.WinnerAddr = win_addr.String
	}
	if err == sql.ErrNoRows {
		return rec,err
	}
	if err!=nil {
			ss.Log_msg(fmt.Sprintf("DB error: %v, q=%v",err,query))
			os.Exit(1)
	}
	team,exists := constants.Teams[rec.HomeTeamId]
	if exists {
		rec.HomeTeam = team.Name + " " + team.Mascot
	}
	team,exists = constants.Teams[rec.AwayTeamId]
	if exists {
		rec.AwayTeam = team.Name + " " + team.Mascot
	}
	sport_id := a.Get_sport_id_from_team(constants,rec.HomeTeamId)
	title,description := a.Get_market_title(sport_id,rec.HomeTeam,rec.AwayTeam,rec.MarketTypeCode,1)
	/*fmt.Printf(
		"sport_id=%v, Home(id=%v)=%v, Away(id=%v)=%v, title=%v, descr=%v\n",
		sport_id,rec.HomeTeamId,rec.HomeTeam,rec.AwayTeamId,rec.AwayTeam,title,description,
	)*/
	rec.Title = title
	rec.Description = description
	rec.MarketRules = a.Get_sports_resolution_rules(sport_id,rec.MarketTypeCode)
	return rec,nil

}
func (ss *SQLStorage) Get_liquidity_change_events(factory_aid,market_id int64,offset,limit int) (int64,[]p.AMM_LiquidityChangedInfo) {

	var query string
	query = "SELECT count(*) AS total " +
			"FROM aa_liquidity_changed AS l " +
			"WHERE l.market_id=$1 AND factory_aid=$2"
			row := ss.db.QueryRow(query,market_id,factory_aid)
	var null_counter sql.NullInt64
	err := row.Scan(&null_counter)
	if (err!=nil) {
		if err==sql.ErrNoRows {

		}
		ss.Log_msg(fmt.Sprintf("Error in Is_feepot(): %v",err))
		os.Exit(1)
	}
	total_rows := null_counter.Int64

	query = "SELECT " +
				"EXTRACT(EPOCH FROM l.time_stamp)::BIGINT AS created_ts, " +
				"l.time_stamp,"+
				"l.block_num,"+
				"tx.tx_hash," +
				"l.user_aid,"+
				"ua.addr," +
				"ra.addr,"+
				"l.collateral," +
				"l.lp_tokens "+
			"FROM aa_liquidity_changed l "+
				"JOIN address ua ON l.user_aid=ua.address_id " +
				"JOIN address ra ON l.recipient_aid=ra.address_id " +
				"JOIN transaction tx ON l.tx_id=tx.id "+
			"WHERE l.market_id=$3 AND factory_aid=$4 "+
			"ORDER BY l.id DESC "+
			"OFFSET $1 LIMIT $2"

	d_query := fmt.Sprintf("SELECT " +
				"EXTRACT(EPOCH FROM l.time_stamp)::BIGINT AS created_ts, " +
				"l.time_stamp,"+
				"l.block_num,"+
				"tx.tx_hash," +
				"l.user_aid,"+
				"ua.addr," +
				"ra.addr,"+
				"l.collateral," +
				"l.lp_tokens "+
			"FROM aa_liquidity_changed l "+
				"JOIN address ua ON l.user_aid=ua.address_id " +
				"JOIN address ra ON l.recipient_aid=ra.address_id " +
				"JOIN transaction tx ON l.tx_id=tx.id "+
			"WHERE l.market_id=%v AND factory_aid=%v "+
			"ORDER BY l.id DESC ",market_id,factory_aid)
	fmt.Printf("query = %v\n",d_query)
	rows,err := ss.db.Query(query,offset,limit,market_id,factory_aid)
	if (err!=nil) {
		ss.Log_msg(fmt.Sprintf("DB error: %v (query=%v)",err,query))
		os.Exit(1)
	}
	records := make([]p.AMM_LiquidityChangedInfo,0,32)

	defer rows.Close()
	for rows.Next() {
		var rec p.AMM_LiquidityChangedInfo
		err=rows.Scan(
			&rec.CreatedTs,
			&rec.CreatedDate,
			&rec.BlockNum,
			&rec.TxHash,
			&rec.UserAid,
			&rec.UserAddr,
			&rec.RecipientAddr,
			&rec.Collateral,
			&rec.Tokens,
		)
		if err!=nil {
			ss.Log_msg(fmt.Sprintf("DB error: %v, q=%v",err,query))
			os.Exit(1)
		}
		rec.MarketId = market_id
		if rec.Collateral < 0 {
			rec.In = true
			rec.Collateral = -rec.Collateral
		}
		if rec.Tokens < 0 {
			rec.Tokens = - rec.Tokens
		}
		records = append(records,rec)
	}
	return total_rows,records
}
func (ss *SQLStorage) Get_shares_swapped(constants *p.AMM_Constants,contract_aid,market_id int64,offset,limit int) (int64,[]p.AA_SharesSwappedInfo) {

	var query string

	query = "SELECT count(*) AS total " +
			"FROM aa_shares_swapped AS l " +
			"WHERE l.market_id=$1 AND factory_aid=$2"
	row := ss.db.QueryRow(query,market_id,contract_aid)
	var null_counter sql.NullInt64
	err := row.Scan(&null_counter)
	if (err!=nil) {
		if err==sql.ErrNoRows {

		}
		ss.Log_msg(fmt.Sprintf("Error in Get_shares_swapped(): %v",err))
		os.Exit(1)
	}
	total_rows := null_counter.Int64

	query = "SELECT " +
				"EXTRACT(EPOCH FROM s.time_stamp)::BIGINT AS created_ts, " +
				"s.time_stamp,"+
				"s.block_num,"+
				"tx.tx_hash," +
				"s.user_aid,"+
				"ua.addr," +
				"s.outcome_idx," +
				"s.collateral," +
				"s.shares, " +
				"sm.home_team_id," +
				"sm.away_team_id," +
				"sm.market_type " +
			"FROM aa_shares_swapped s "+
				"JOIN address ua ON s.user_aid=ua.address_id " +
				"JOIN transaction tx ON s.tx_id=tx.id "+
				"LEFT JOIN aa_sports_market sm ON (sm.contract_aid=$4) AND (sm.market_id=s.market_id) "+
			"WHERE s.market_id=$3 AND factory_aid=$4 "+
			"ORDER BY s.id DESC "+
			"OFFSET $1 LIMIT $2"
		d_query := fmt.Sprintf("SELECT " +
				"EXTRACT(EPOCH FROM s.time_stamp)::BIGINT AS created_ts, " +
				"s.time_stamp,"+
				"s.block_num,"+
				"tx.tx_hash," +
				"s.user_aid,"+
				"ua.addr," +
				"s.outcome_idx," +
				"s.collateral," +
				"s.shares " +
			"FROM aa_shares_swapped s "+
				"JOIN address ua ON s.user_aid=ua.address_id " +
				"JOIN transaction tx ON s.tx_id=tx.id "+
			"WHERE s.market_id=%v AND factory_aid=%v "+
			"ORDER BY s.id DESC "+
			"OFFSET %v LIMIT %v",
			market_id,contract_aid,offset,limit)
		fmt.Printf("q = %v\n",d_query)
	rows,err := ss.db.Query(query,offset,limit,market_id,contract_aid)
	if (err!=nil) {
		ss.Log_msg(fmt.Sprintf("DB error: %v (query=%v)",err,query))
		os.Exit(1)
	}
	records := make([]p.AA_SharesSwappedInfo,0,32)

	defer rows.Close()
	for rows.Next() {
		var rec p.AA_SharesSwappedInfo
		var home_id,away_id,mkt_type sql.NullInt64
		err=rows.Scan(
			&rec.CreatedTs,
			&rec.CreatedDate,
			&rec.BlockNum,
			&rec.TxHash,
			&rec.UserAid,
			&rec.UserAddr,
			&rec.Outcome,
			&rec.Collateral,
			&rec.Shares,
			&home_id,
			&away_id,
			&mkt_type,
		)
		if err!=nil {
			ss.Log_msg(fmt.Sprintf("DB error: %v, q=%v",err,query))
			os.Exit(1)
		}
		if home_id.Valid {
			h_team,h_exists := constants.Teams[home_id.Int64]
			if h_exists {
				a_team,a_exists := constants.Teams[away_id.Int64]
				if a_exists {
					sport_id := a.Get_sport_id_from_team(constants,home_id.Int64)

					rec.OutcomeStr = a.Get_outcome_name(rec.Outcome,sport_id,h_team.Name + " " + h_team.Mascot,a_team.Name + " " + a_team.Mascot,mkt_type.Int64,"1")
				}
			}
		}
		rec.MarketId = market_id
		if rec.Collateral > 0 {
			rec.Buy = true
		} else {
			rec.Collateral = -rec.Collateral
		}
		if rec.Shares < 0 {
			rec.Shares = -rec.Shares
		}
		records = append(records,rec)
	}
	return total_rows,records
}
func (ss *SQLStorage) Get_amm_user_swaps(constants *p.AMM_Constants,user_aid int64,offset,limit int) (int64,[]p.AA_SharesSwappedInfo) {

	var query string

	query = "SELECT count(*) AS total " +
			"FROM aa_shares_swapped AS l " +
			"WHERE l.user_aid=$1"
	row := ss.db.QueryRow(query,user_aid)
	var null_counter sql.NullInt64
	err := row.Scan(&null_counter)
	if (err!=nil) {
		if err==sql.ErrNoRows {

		}
		ss.Log_msg(fmt.Sprintf("Error in Get_amm_user_swaps(): %v",err))
		os.Exit(1)
	}
	total_rows := null_counter.Int64

	query = "SELECT " +
				"EXTRACT(EPOCH FROM s.time_stamp)::BIGINT AS created_ts, " +
				"s.time_stamp,"+
				"s.block_num,"+
				"tx.tx_hash," +
				"s.user_aid,"+
				"s.market_id,"+
				"ua.addr," +
				"s.outcome_idx," +
				"s.collateral," +
				"s.shares, " +
				"sm.home_team_id," +
				"sm.away_team_id," +
				"sm.market_type " +
			"FROM aa_shares_swapped s "+
				"JOIN address ua ON s.user_aid=ua.address_id " +
				"JOIN transaction tx ON s.tx_id=tx.id "+
				"LEFT JOIN aa_sports_market sm ON (sm.contract_aid=s.factory_aid) AND (sm.market_id=s.market_id) "+
			"WHERE s.user_aid=$3 "+
			"ORDER BY s.id DESC "+
			"OFFSET $1 LIMIT $2"
	rows,err := ss.db.Query(query,offset,limit,user_aid)
	if (err!=nil) {
		ss.Log_msg(fmt.Sprintf("DB error: %v (query=%v)",err,query))
		os.Exit(1)
	}
	records := make([]p.AA_SharesSwappedInfo,0,32)

	defer rows.Close()
	for rows.Next() {
		var rec p.AA_SharesSwappedInfo
		var home_id,away_id,mkt_type sql.NullInt64
		err=rows.Scan(
			&rec.CreatedTs,
			&rec.CreatedDate,
			&rec.BlockNum,
			&rec.TxHash,
			&rec.UserAid,
			&rec.MarketId,
			&rec.UserAddr,
			&rec.Outcome,
			&rec.Collateral,
			&rec.Shares,
			&home_id,
			&away_id,
			&mkt_type,
		)
		if err!=nil {
			ss.Log_msg(fmt.Sprintf("DB error: %v, q=%v",err,query))
			os.Exit(1)
		}
		if home_id.Valid {
			h_team,h_exists := constants.Teams[home_id.Int64]
			if h_exists {
				a_team,a_exists := constants.Teams[away_id.Int64]
				if a_exists {
					sport_id := a.Get_sport_id_from_team(constants,home_id.Int64)

					rec.OutcomeStr = a.Get_outcome_name(rec.Outcome,sport_id,h_team.Name + " " +h_team.Mascot,a_team.Name + " " +a_team.Mascot,mkt_type.Int64,"1")
				}
			}
		}
		if rec.Collateral > 0 {
			rec.Buy = true
		} else {
			rec.Collateral = -rec.Collateral
		}
		if rec.Shares < 0 {
			rec.Shares = -rec.Shares
		}
		records = append(records,rec)
	}
	return total_rows,records
}
func (ss *SQLStorage) Get_amm_user_liquidity(constants *p.AMM_Constants,user_aid int64,offset,limit int) (int64,[]p.AMM_LiquidityChangedInfo) {

	var query string
	query = "SELECT count(*) AS total " +
			"FROM aa_liquidity_changed AS l " +
			"WHERE l.user_aid=$1"
			row := ss.db.QueryRow(query,user_aid)
	var null_counter sql.NullInt64
	err := row.Scan(&null_counter)
	if (err!=nil) {
		if err==sql.ErrNoRows {

		}
		ss.Log_msg(fmt.Sprintf("Error in Get_amm_user_liquidity(): %v",err))
		os.Exit(1)
	}
	total_rows := null_counter.Int64

	query = "SELECT " +
				"EXTRACT(EPOCH FROM l.time_stamp)::BIGINT AS created_ts, " +
				"l.time_stamp,"+
				"l.block_num,"+
				"tx.tx_hash," +
				"l.user_aid,"+
				"l.market_id,"+
				"ua.addr," +
				"ra.addr,"+
				"l.collateral," +
				"l.lp_tokens "+
			"FROM aa_liquidity_changed l "+
				"JOIN address ua ON l.user_aid=ua.address_id " +
				"JOIN address ra ON l.recipient_aid=ra.address_id " +
				"JOIN transaction tx ON l.tx_id=tx.id "+
			"WHERE l.user_aid=$3 "+
			"ORDER BY l.id DESC "+
			"OFFSET $1 LIMIT $2"
	fmt.Printf("quuery (user_aid=%v) : %v\n",user_aid,query)
	rows,err := ss.db.Query(query,offset,limit,user_aid)
	if (err!=nil) {
		ss.Log_msg(fmt.Sprintf("DB error: %v (query=%v)",err,query))
		os.Exit(1)
	}
	records := make([]p.AMM_LiquidityChangedInfo,0,32)

	defer rows.Close()
	for rows.Next() {
		var rec p.AMM_LiquidityChangedInfo
		err=rows.Scan(
			&rec.CreatedTs,
			&rec.CreatedDate,
			&rec.BlockNum,
			&rec.TxHash,
			&rec.UserAid,
			&rec.MarketId,
			&rec.UserAddr,
			&rec.RecipientAddr,
			&rec.Collateral,
			&rec.Tokens,
		)
		if err!=nil {
			ss.Log_msg(fmt.Sprintf("DB error: %v, q=%v",err,query))
			os.Exit(1)
		}
		if rec.Collateral < 0 {
			rec.In = true
			rec.Collateral = -rec.Collateral
		}
		if rec.Tokens < 0 {
			rec.Tokens = - rec.Tokens
		}
		records = append(records,rec)
	}
	return total_rows,records
}
func (ss *SQLStorage) Get_market_pool_aid(factory_aid,market_id int64) (int64,error) {

	var query string
	query = "SELECT pool_aid FROM aa_pool_created WHERE factory_aid=$1 AND market_id=$2"

	var null_id sql.NullInt64
	res := ss.db.QueryRow(query,factory_aid,market_id)
	err := res.Scan(&null_id)
	if (err!=nil) {
		if err == sql.ErrNoRows {
			return 0,err
		} else {
			ss.Log_msg(fmt.Sprintf("DB error: %v q=%v",err,query))
			os.Exit(1)
		}
	}
	return null_id.Int64,nil
}
func (ss *SQLStorage) Update_status_not_augur_block_num(block_num int64) {

	var query string
	query = "UPDATE aa_proc_status SET last_block_outgui = $1"

	_,err := ss.db.Exec(query,block_num)
	if (err!=nil) {
		ss.Log_msg(fmt.Sprintf("DB error: %v q=%v",err,query))
		os.Exit(1)
	}
}
func (ss *SQLStorage) Get_status_not_augur_block_num() (int64,int64) {

	var query string
	query = "SELECT " +
				"e.block_num AS last_block_on_chain," +
				"s.last_block_outgui "+
			"FROM aa_proc_status s " +
				"JOIN evt_log e ON s.last_evt_id=e.id "

	var last_block_chain,last_block_processed sql.NullInt64
	res := ss.db.QueryRow(query)
	err := res.Scan(&last_block_chain,&last_block_processed)
	if (err!=nil) {
		if err == sql.ErrNoRows {
			ss.Log_msg(fmt.Sprintf("aa_proc_status' table is empty, insert a record"))
			os.Exit(1)
		} else {
			ss.Log_msg(fmt.Sprintf("DB error: %v q=%v",err,query))
			os.Exit(1)
		}
	}
	return last_block_chain.Int64,last_block_processed.Int64
}
func (ss *SQLStorage) Get_shares_minted_burned_in_block_range(table string,from_block,to_block int64) []p.AMM_TxId_Rec  {

	var query string
	query = "SELECT " +
				"t.id,t.tx_id,ss.id AS shares_swapped_id,liq.id AS liqquidity_id,bs.id as balancer_id " +
			"FROM "+table+" t " +
			"LEFT JOIN aa_shares_swapped ss ON t.tx_id = ss.tx_id " +
			"LEFT JOIN aa_liquidity_changed liq ON t.tx_id=liq.tx_id " +
			"LEFT JOIN bswap bs ON t.tx_id=bs.tx_id " +
			"WHERE (t.block_num >= $1) AND (t.block_num<=$2) ORDER BY t.block_num"

	rows,err := ss.db.Query(query,from_block,to_block)
	if (err!=nil) {
		ss.Log_msg(fmt.Sprintf("DB error: %v (query=%v)",err,query))
		os.Exit(1)
	}
	records := make([]p.AMM_TxId_Rec,0,32)

	defer rows.Close()
	for rows.Next() {
		var rec p.AMM_TxId_Rec
		var null_ss_id,null_liq_id,null_bal_id sql.NullInt64
		err=rows.Scan(&rec.RecordId,&rec.TxId,&null_ss_id,&null_liq_id,&null_bal_id)
		if err!=nil {
			ss.Log_msg(fmt.Sprintf("DB error: %v, q=%v",err,query))
			os.Exit(1)
		}
		if null_ss_id.Valid { rec.SharesSwappedId = null_ss_id.Int64; }
		if null_liq_id.Valid { rec.LiquidityId = null_liq_id.Int64; }
		if null_bal_id.Valid { rec.BalancerId = null_bal_id.Int64 }
		records = append(records,rec)
	}

	return records
}
func (ss *SQLStorage) Get_balancer_swaps_for_augur_markets(from_block,to_block int64) []p.AMM_TxBalSwaps  {

	records := make([]p.AMM_TxBalSwaps,0,32)
	var query string
	query = "SELECT " +
				"bs.id,"+
				"bs.tx_id, " +
				"ss.id AS shares_swapped_id, " +
				"liq.id AS liquidity_id " +
			"FROM bswap bs " +
			"LEFT JOIN aa_shares_swapped ss ON ss.tx_id=bs.tx_id " +
			"LEFT JOIN aa_liquidity_changed liq ON liq.tx_id=bs.tx_id "+
			"WHERE (bs.block_num >= $1) AND (bs.block_num<=$2) ORDER BY bs.block_num"

	rows,err := ss.db.Query(query,from_block,to_block)
	if (err!=nil) {
		ss.Log_msg(fmt.Sprintf("DB error: %v (query=%v)",err,query))
		os.Exit(1)
	}

	defer rows.Close()
	for rows.Next() {
		var rec p.AMM_TxBalSwaps
		var null_ss_id,null_liq_id sql.NullInt64
		err=rows.Scan(&rec.RecordId,&rec.TxId,&null_ss_id,&null_liq_id)
		if err!=nil {
			ss.Log_msg(fmt.Sprintf("DB error: %v, q=%v",err,query))
			os.Exit(1)
		}
		if null_ss_id.Valid { rec.SharesSwappedId = null_ss_id.Int64 }
		if null_liq_id.Valid { rec.LiquidityId = null_liq_id.Int64 }
		records = append(records,rec)
	}

	return records
}
func (ss *SQLStorage) Insert_not_augur_mark(record_id int64,rec_type int) {

	var query string
	query = "INSERT INTO aa_not_augur (rec_id,obj_type)" +
			"VALUES ($1,$2)"

	_,err := ss.db.Exec(query,record_id,rec_type)
	if err != nil {
		ss.Log_msg(
			fmt.Sprintf(
				"DB error: can't insert into 'aa_not_augur' table (record_id=%v,type=%v): %v; q=%v",
				record_id,rec_type,err,query,
			),
		)
		os.Exit(1)
	}
}
