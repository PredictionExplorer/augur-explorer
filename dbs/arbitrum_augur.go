package dbs

import (
	"fmt"
	"os"
	"database/sql"
	_  "github.com/lib/pq"

	"github.com/ethereum/go-ethereum/common"

	p "github.com/PredictionExplorer/augur-explorer/primitives"
	a "github.com/PredictionExplorer/augur-explorer/amm"

)
func (ss *SQLStorage) Get_arbitrum_augur_contract_addresses() (p.AA_ContractAddrs) {

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
			ss.Log_msg(fmt.Sprintf("No contract addresses in AMM contracts table %v"))
			os.Exit(1)
		} else {
			ss.Log_msg(fmt.Sprintf("Error in Get_arbitrum_augur_contract_addresses(): %v",err))
			os.Exit(1)
		}
	}
	c_addrs.AMM_Factory=common.HexToAddress(amm_factory)
	c_addrs.SportsFactory=common.HexToAddress(sports_factory)
	c_addrs.TrustedFactory=common.HexToAddress(trusted_factory)
	return c_addrs
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
func (ss *SQLStorage) Get_sport_markets(status,sort int64,offset,limit int,constants *p.AMM_Constants,contracts *p.AA_ContractAddrs) (int64,[]p.AMM_SportMarket) {

	var query string

	query = "SELECT address_id FROM address WHERE addr=$1"
	row := ss.db.QueryRow(query,contracts.SportsFactory.String())
	var null_id sql.NullInt64
	err := row.Scan(&null_id)
	if (err != nil) {
		if err == sql.ErrNoRows {
			ss.Log_msg(fmt.Sprintf("Can't find AMM module contract addresses"))
			os.Exit(1)
		}
		ss.Log_msg(fmt.Sprintf("Error in Is_feepot(): %v",err))
		os.Exit(1)
	}
	amm_factory_aid := null_id.Int64

	query = "SELECT count(*) AS total " +
			"FROM aa_sports_market AS m " +
			"WHERE m.contract_aid=$1"
	row = ss.db.QueryRow(query,null_id.Int64)
	var null_counter sql.NullInt64
	err = row.Scan(&null_counter)
	if (err!=nil) {
		if err==sql.ErrNoRows {

		}
		ss.Log_msg(fmt.Sprintf("Error in Get_sport_markets(): %v",err))
		os.Exit(1)
	}
	total_rows := null_counter.Int64

	query = "SELECT " +
				"EXTRACT(EPOCH FROM time_stamp)::BIGINT AS created_ts, " +
				"time_stamp," +
				"m.block_num, " +
				"tx.tx_hash," +
				"m.market_id," +
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
				"m.market_type " +
			"FROM aa_sports_market AS m " +
				"LEFT JOIN address ca ON m.creator_aid=ca.address_id " +
				"LEFT JOIN address fa ON m.contract_aid=fa.address_id " +
				"JOIN transaction tx ON m.tx_id=tx.id " +
			"WHERE m.contract_aid=$1" +
			"ORDER BY m.time_stamp " +
			"OFFSET $2 LIMIT $3"

	rows,err := ss.db.Query(query,amm_factory_aid,offset,limit)
	if (err!=nil) {
		ss.Log_msg(fmt.Sprintf("DB error: %v (query=%v)",err,query))
		os.Exit(1)
	}
	records := make([]p.AMM_SportMarket,0,32)

	defer rows.Close()
	for rows.Next() {
		var rec p.AMM_SportMarket
		err=rows.Scan(
			&rec.CreatedTs,
			&rec.CreatedDate,
			&rec.BlockNum,
			&rec.TxHash,
			&rec.MarketId,
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
		)
		if err!=nil {
			ss.Log_msg(fmt.Sprintf("DB error: %v, q=%v",err,query))
			os.Exit(1)
		}
		team,exists := constants.Teams[rec.HomeTeamId]
		if exists {
			rec.HomeTeam = team.Name
		}
		team,exists = constants.Teams[rec.AwayTeamId]
		if exists {
			rec.AwayTeam = team.Name
		}
		sport_id := a.Get_sport_id_from_team(constants,rec.HomeTeamId)
		title,description := a.Get_market_title(sport_id,rec.HomeTeam,rec.AwayTeam,rec.MarketTypeCode,1)
		/*fmt.Printf(
			"sport_id=%v, Home(id=%v)=%v, Away(id=%v)=%v, title=%v, descr=%v\n",
			sport_id,rec.HomeTeamId,rec.HomeTeam,rec.AwayTeamId,rec.AwayTeam,title,description,
		)*/
		rec.Title = title
		rec.Description = description
		records = append(records,rec)
	}
	return total_rows,records

}
func (ss *SQLStorage) Get_sport_market_info(constants *p.AMM_Constants,contracts *p.AA_ContractAddrs,market_id int64) (p.AMM_SportMarket,error) {

	var query string
	query = "SELECT address_id FROM address WHERE addr=$1"
	row := ss.db.QueryRow(query,contracts.SportsFactory.String())
	var null_id sql.NullInt64
	err := row.Scan(&null_id)
	if (err != nil) {
		if err == sql.ErrNoRows {
			ss.Log_msg(fmt.Sprintf("Can't find AMM module contract addresses"))
			os.Exit(1)
		}
		ss.Log_msg(fmt.Sprintf("Error in Is_feepot(): %v",err))
		os.Exit(1)
	}
	amm_factory_aid := null_id.Int64
	query = "SELECT " +
				"EXTRACT(EPOCH FROM time_stamp)::BIGINT AS created_ts, " +
				"time_stamp," +
				"m.block_num, " +
				"tx.tx_hash," +
				"m.market_id," +
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
				"m.market_type " +
			"FROM aa_sports_market AS m " +
				"LEFT JOIN address ca ON m.creator_aid=ca.address_id " +
				"LEFT JOIN address fa ON m.contract_aid=fa.address_id " +
				"JOIN transaction tx ON m.tx_id=tx.id " +
			"WHERE m.market_id=$1 AND contract_aid=$2"

	row = ss.db.QueryRow(query,market_id,amm_factory_aid)
	if (err!=nil) {
		ss.Log_msg(fmt.Sprintf("DB error: %v (query=%v)",err,query))
		os.Exit(1)
	}

	var rec p.AMM_SportMarket
	err=row.Scan(
			&rec.CreatedTs,
			&rec.CreatedDate,
			&rec.BlockNum,
			&rec.TxHash,
			&rec.MarketId,
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
	)
	if err == sql.ErrNoRows {
		return rec,err
	}
	if err!=nil {
			ss.Log_msg(fmt.Sprintf("DB error: %v, q=%v",err,query))
			os.Exit(1)
	}
	team,exists := constants.Teams[rec.HomeTeamId]
	if exists {
		rec.HomeTeam = team.Name
	}
	team,exists = constants.Teams[rec.AwayTeamId]
	if exists {
		rec.AwayTeam = team.Name
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
func (ss *SQLStorage) Get_liquidity_change_events(factory_addr string,market_id int64,offset,limit int) (int64,[]p.AMM_LiquidityChangedInfo) {

	var query string

	query = "SELECT address_id FROM address WHERE addr=$1"
	row := ss.db.QueryRow(query,factory_addr)
	var null_id sql.NullInt64
	err := row.Scan(&null_id)
	if (err != nil) {
		if err == sql.ErrNoRows {
			ss.Log_msg(fmt.Sprintf("Can't find AMM module contract addresses"))
			os.Exit(1)
		}
		ss.Log_msg(fmt.Sprintf("Error in Get_liquidity_change_events(): %v",err))
		os.Exit(1)
	}
	amm_factory_aid := null_id.Int64

	query = "SELECT count(*) AS total " +
			"FROM aa_liquidity_changed AS l " +
			"WHERE l.market_id=$1 AND contract_aid=$2"
	row = ss.db.QueryRow(query,market_id,null_id.Int64)
	var null_counter sql.NullInt64
	err = row.Scan(&null_counter)
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
			"WHERE l.market_id=$3 AND contract_aid=$4 "+
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
			"WHERE l.market_id=%v AND contract_aid=%v "+
			"ORDER BY l.id DESC ",market_id,amm_factory_aid)
	fmt.Printf("query = %v\n",d_query)
	rows,err := ss.db.Query(query,offset,limit,market_id,amm_factory_aid)
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
func (ss *SQLStorage) Get_shares_swapped(constants *p.AMM_Constants,factory_addr string,market_id int64,offset,limit int) (int64,[]p.AA_SharesSwappedInfo) {

	var query string

	query = "SELECT address_id FROM address WHERE addr=$1"
	row := ss.db.QueryRow(query,factory_addr)
	var null_id sql.NullInt64
	err := row.Scan(&null_id)
	if (err != nil) {
		if err == sql.ErrNoRows {
			ss.Log_msg(fmt.Sprintf("Can't find AMM module contract addresses"))
			os.Exit(1)
		}
		ss.Log_msg(fmt.Sprintf("Error in Is_feepot(): %v",err))
		os.Exit(1)
	}
	amm_factory_aid := null_id.Int64

	query = "SELECT count(*) AS total " +
			"FROM aa_shares_swapped AS l " +
			"WHERE l.market_id=$1 AND factory_aid=$2"
	row = ss.db.QueryRow(query,market_id,null_id.Int64)
	var null_counter sql.NullInt64
	err = row.Scan(&null_counter)
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
			market_id,amm_factory_aid,offset,limit)
		fmt.Printf("q = %v\n",d_query)
	rows,err := ss.db.Query(query,offset,limit,market_id,amm_factory_aid)
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
		fmt.Printf("home_id.Valid=%v mkt_type=%v\n",home_id.Valid,mkt_type.Int64)
		if home_id.Valid {
			h_team,h_exists := constants.Teams[home_id.Int64]
			if h_exists {
				a_team,a_exists := constants.Teams[away_id.Int64]
				if a_exists {
					sport_id := a.Get_sport_id_from_team(constants,home_id.Int64)
					fmt.Printf("sport_id=%v home: %v , away: %v rec.Outcome=%v\n",sport_id,h_team.Name,a_team.Name,rec.Outcome)

					rec.OutcomeStr = a.Get_outcome_name(rec.Outcome,sport_id,h_team.Name,a_team.Name,mkt_type.Int64,"1")
					fmt.Printf("OutcomeStr: %v\n",rec.OutcomeStr)
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
		fmt.Printf("rec = %v\n",rec)
		records = append(records,rec)
	}
	return total_rows,records
}
