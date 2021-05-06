package dbs

import (
	"fmt"
	"os"
	"database/sql"
	_  "github.com/lib/pq"

	"github.com/ethereum/go-ethereum/common"

	p "github.com/PredictionExplorer/augur-explorer/primitives"
)
func (ss *SQLStorage) Get_arbitrum_augur_contract_addresses() (p.AA_ContractAddrs,error) {

	var query string
	query="SELECT " +
				"amm_factory,sports_factory,trusted_factory "+
			"FROM aa_caddrs";
	row := ss.db.QueryRow(query)
	var c_addrs p.AA_ContractAddrs
	var err error
	var (
		amm_factory string
		sports_factory string
		trusted_factory string
	)
	err=row.Scan(
		&amm_factory,&sports_factory,&trusted_factory,
	);
	if (err!=nil) {
		if err == sql.ErrNoRows {
		} else {
			ss.Log_msg(fmt.Sprintf("Error in Get_arbitrum_augur_contract_addresses(): %v",err))
			os.Exit(1)
		}
		return c_addrs,err
	}
	c_addrs.AMM_Factory=common.HexToAddress(amm_factory)
	c_addrs.SportsFactory=common.HexToAddress(sports_factory)
	c_addrs.TrustedFactory=common.HexToAddress(trusted_factory)
	return c_addrs,nil
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
				"market_id,factory_aid,user_aid,recipient_aid,collateral,lp_tokens" +
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
			recipient_aid,
			evt.Collateral,
			evt.LpTokens,
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
	var query string
	query = "INSERT INTO aa_sports_market (" +
				"evtlog_id,block_num,tx_id,contract_aid,time_stamp,"+
				"market_id,start_time,end_time,market_type,creator_aid,"+
				"event_id,home_team_id,away_team_id,score" +
			") VALUES ("+
				"$1,$2,$3,$4,TO_TIMESTAMP($5)"+
				",$6,TO_TIMESTAMP($7),TO_TIMESTAMP($8),$9,$10,"+
				"$11,$12,$13,$14"+
			")"

	_,err := ss.db.Exec(query,
			evt.EvtId,
			evt.BlockNum,
			evt.TxId,
			contract_aid,
			evt.TimeStamp,
			evt.MarketId,
			evt.EstimatedStarTime,
			evt.EndTime,
			evt.MarketType,
			creator_aid,
			evt.EventId,
			evt.HomeTeamId,
			evt.AwayTeamId,
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
				"market_id,factory_aid,user_aid,collateral,shares" +
			") VALUES ($1,$2,$3,$4,TO_TIMESTAMP($5),$6,$7,$8,$9::DECIMAL/1e+6,$10::DECIMAL/1e+18)"

	_,err := ss.db.Exec(query,
			evt.EvtId,
			evt.BlockNum,
			evt.TxId,
			contract_aid,
			evt.TimeStamp,
			evt.MarketId,
			factory_aid,
			user_aid,
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
	err := row.Scan(&addr)
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
