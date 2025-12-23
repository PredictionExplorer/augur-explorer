package randomwalk

import (
	"database/sql"
	"fmt"
	"os"

	"github.com/PredictionExplorer/augur-explorer/rwcg/dbs"
	rwp "github.com/PredictionExplorer/augur-explorer/rwcg/primitives/randomwalk"
)

// SQLStorageWrapper wraps dbs.SQLStorage to provide RandomWalk-specific database methods
type SQLStorageWrapper struct {
	S *dbs.SQLStorage
}


// =============================================================================
// PROCESSING STATUS
// =============================================================================

func (sw *SQLStorageWrapper) Get_randomwalk_processing_status() rwp.ProcStatus {

	var output rwp.ProcStatus
	var null_id,null_block sql.NullInt64

	var query string
	for {
		query = "SELECT last_evt_id,last_block FROM rw_proc_status"

		res := sw.S.Db().QueryRow(query)
		err := res.Scan(&null_id,&null_block)
		if (err!=nil) {
			if err == sql.ErrNoRows {
				query = "INSERT INTO rw_proc_status DEFAULT VALUES"
				_,err := sw.S.Db().Exec(query)
				if (err!=nil) {
					sw.S.Log_msg(fmt.Sprintf("DB error: %v q=%v",err,query))
					os.Exit(1)
				}
			} else {
				sw.S.Log_msg(fmt.Sprintf("DB error: %v q=%v",err,query))
				os.Exit(1)
			}
		} else {
			break
		}
	}
	if null_id.Valid {
		output.LastIdProcessed = null_id.Int64
	}
	if null_block.Valid {
		output.LastBlockNum = null_block.Int64
	}
	return output
}
func (sw *SQLStorageWrapper) Update_randomwalk_process_status(status *rwp.ProcStatus) {

	var query string
	query = "UPDATE rw_proc_status SET last_evt_id = $1,last_block=$2"

	_,err := sw.S.Db().Exec(query,status.LastIdProcessed,status.LastBlockNum)
	if (err!=nil) {
		sw.S.Log_msg(fmt.Sprintf("DB error: %v q=%v",err,query))
		os.Exit(1)
	}
}

// =============================================================================
// CONTRACT CONFIGURATION
// =============================================================================

func (sw *SQLStorageWrapper) Get_randomwalk_contract_addresses() rwp.ContractAddresses {

	var output rwp.ContractAddresses
	var query string
	query = "SELECT "+
				"marketplace_addr,randomwalk_addr," +
				"mp_a.address_id,rw_a.address_id "+
			"FROM rw_contracts rw " +
				"JOIN address mp_a ON rw.marketplace_addr=mp_a.addr " +
				"JOIN address rw_a ON rw.randomwalk_addr=rw_a.addr "

	res := sw.S.Db().QueryRow(query)
	err := res.Scan(
		&output.MarketPlace,
		&output.RandomWalk,
		&output.MarketPlaceAid,
		&output.RandomWalkAid,
	)
	if (err!=nil) {
		if err == sql.ErrNoRows {
			sw.S.Log_msg("Can't find record in rw_contracts table for contract addresses")
			os.Exit(1)
		}
		sw.S.Log_msg(fmt.Sprintf("Get_randomwalk_contract_addresses() failed: %v, q=%v",err,query))
		os.Exit(1)
	}
	return output
}

// =============================================================================
// OFFER OPERATIONS
// =============================================================================

func (sw *SQLStorageWrapper) Insert_new_offer(evt *rwp.NewOffer) {

	contract_aid:=sw.S.Lookup_or_create_address(evt.Contract,evt.BlockNum,evt.TxId)
	rwalk_aid:=sw.S.Lookup_or_create_address(evt.RWalkAddr,evt.BlockNum,evt.TxId)
	buyer_aid:=sw.S.Lookup_or_create_address(evt.Buyer,evt.BlockNum,evt.TxId)
	seller_aid:=sw.S.Lookup_or_create_address(evt.Seller,evt.BlockNum,evt.TxId)
	otype:=int(1)
	if evt.Seller == "0x0000000000000000000000000000000000000000" {
		otype=0
	}
	var query string
	query = "INSERT INTO rw_new_offer(" +
				"evtlog_id,block_num,tx_id,time_stamp,contract_aid, "+
				"rwalk_aid,offer_id,otype,token_id,buyer_aid,seller_aid,active,price" +
			") VALUES (" +
				"$1,$2,$3,TO_TIMESTAMP($4),$5,$6,$7,$8,$9,$10,$11,$12,$13"+
			")"
	_,err := sw.S.Db().Exec(query,
		evt.EvtId,
		evt.BlockNum,
		evt.TxId,
		evt.TimeStamp,
		contract_aid,
		rwalk_aid,
		evt.OfferId,
		otype,
		evt.TokenId,
		buyer_aid,
		seller_aid,
		true,
		evt.Price,
	)
	if err != nil {
		sw.S.Log_msg(fmt.Sprintf("DB error: can't insert into new_offer table: %v\n",err))
		os.Exit(1)
	}
}
func (sw *SQLStorageWrapper) Insert_item_bought(evt *rwp.ItemBought) {

	contract_aid:=sw.S.Lookup_or_create_address(evt.Contract,evt.BlockNum,evt.TxId)
	seller_aid:=sw.S.Lookup_or_create_address(evt.SellerAddr,evt.BlockNum,evt.TxId)
	buyer_aid:=sw.S.Lookup_or_create_address(evt.BuyerAddr,evt.BlockNum,evt.TxId)
	var query string
	query = "INSERT INTO rw_item_bought(" +
				"evtlog_id,block_num,tx_id,time_stamp,contract_aid, "+
				"offer_id,seller_aid,buyer_aid" +
			") VALUES (" +
				"$1,$2,$3,TO_TIMESTAMP($4),$5,$6,$7,$8"+
			")"
	_,err := sw.S.Db().Exec(query,
		evt.EvtId,
		evt.BlockNum,
		evt.TxId,
		evt.TimeStamp,
		contract_aid,
		evt.OfferId,
		seller_aid,
		buyer_aid,
	)
	if err != nil {
		sw.S.Log_msg(fmt.Sprintf("DB error: can't insert into item_bought table: %v\n",err))
		os.Exit(1)
	}

}
func (sw *SQLStorageWrapper) Insert_offer_canceled(evt *rwp.OfferCanceled) {

	contract_aid:=sw.S.Lookup_or_create_address(evt.Contract,evt.BlockNum,evt.TxId)
	var query string
	query = "INSERT INTO rw_offer_canceled(" +
				"evtlog_id,block_num,tx_id,time_stamp,contract_aid, "+
				"offer_id" +
			") VALUES (" +
				"$1,$2,$3,TO_TIMESTAMP($4),$5,$6"+
			")"
	_,err := sw.S.Db().Exec(query,
		evt.EvtId,
		evt.BlockNum,
		evt.TxId,
		evt.TimeStamp,
		contract_aid,
		evt.OfferId,
	)
	if err != nil {
		sw.S.Log_msg(fmt.Sprintf("DB error: can't insert into offer_canceled table: %v\n",err))
		os.Exit(1)
	}

}

// =============================================================================
// TOKEN OPERATIONS
// =============================================================================

func (sw *SQLStorageWrapper) Insert_withdrawal(evt *rwp.Withdrawal) {

	contract_aid:=sw.S.Lookup_or_create_address(evt.Contract,evt.BlockNum,evt.TxId)
	aid:=sw.S.Lookup_or_create_address(evt.Destination,evt.BlockNum,evt.TxId)
	var query string
	query = "INSERT INTO rw_withdrawal(" +
				"evtlog_id,block_num,tx_id,time_stamp,contract_aid, "+
				"aid,token_id,amount" +
			") VALUES (" +
				"$1,$2,$3,TO_TIMESTAMP($4),$5,$6,$7,$8"+
			")"
	_,err := sw.S.Db().Exec(query,
		evt.EvtId,
		evt.BlockNum,
		evt.TxId,
		evt.TimeStamp,
		contract_aid,
		aid,
		evt.TokenId,
		evt.Amount,
	)
	if err != nil {
		sw.S.Log_msg(fmt.Sprintf("DB error: can't insert into rw_withdrawal table: %v\n",err))
		os.Exit(1)
	}

}
func (sw *SQLStorageWrapper) Insert_token_name(evt *rwp.TokenName) {

	contract_aid:=sw.S.Lookup_or_create_address(evt.Contract,evt.BlockNum,evt.TxId)
	var query string
	query = "INSERT INTO rw_token_name(" +
				"evtlog_id,block_num,tx_id,time_stamp,contract_aid, "+
				"token_id,new_name" +
			") VALUES (" +
				"$1,$2,$3,TO_TIMESTAMP($4),$5,$6,$7"+
			")"
	_,err := sw.S.Db().Exec(query,
		evt.EvtId,
		evt.BlockNum,
		evt.TxId,
		evt.TimeStamp,
		contract_aid,
		evt.TokenId,
		evt.NewName,
	)
	if err != nil {
		sw.S.Log_msg(fmt.Sprintf("DB error: can't insert into rw_token_name table: %v\n",err))
		os.Exit(1)
	}

}
func (sw *SQLStorageWrapper) Insert_mint_event(evt *rwp.MintEvent) {

	contract_aid:=sw.S.Lookup_or_create_address(evt.Contract,evt.BlockNum,evt.TxId)
	owner_aid:=sw.S.Lookup_or_create_address(evt.Owner,evt.BlockNum,evt.TxId)
	var query string
	query = "INSERT INTO rw_mint_evt(" +
				"evtlog_id,block_num,tx_id,time_stamp,contract_aid, "+
				"token_id,owner_aid,seed,seed_num,price" +
			") VALUES (" +
				"$1,$2,$3,TO_TIMESTAMP($4),$5,$6,$7,$8,$9,$10"+
			")"
	_,err := sw.S.Db().Exec(query,
		evt.EvtId,
		evt.BlockNum,
		evt.TxId,
		evt.TimeStamp,
		contract_aid,
		evt.TokenId,
		owner_aid,
		evt.Seed,
		evt.SeedNum,
		evt.Price,
	)
	if err != nil {
		sw.S.Log_msg(fmt.Sprintf("DB error: can't insert into rw_mint_evt table: %v\n",err))
		os.Exit(1)
	}
}
func (sw *SQLStorageWrapper) Insert_token_transfer_event(evt *rwp.Transfer) {

	contract_aid:=sw.S.Lookup_or_create_address(evt.Contract,evt.BlockNum,evt.TxId)
	from_aid:=sw.S.Lookup_or_create_address(evt.From,evt.BlockNum,evt.TxId)
	to_aid:=sw.S.Lookup_or_create_address(evt.To,evt.BlockNum,evt.TxId)
	otype := int(0)
	if evt.From == "0x0000000000000000000000000000000000000000" {
		otype = 1
	}
	if evt.To == "0x0000000000000000000000000000000000000000" {
		otype = 2
	}
	var query string
	query = "INSERT INTO rw_transfer(" +
				"evtlog_id,block_num,tx_id,time_stamp,contract_aid, "+
				"token_id,from_aid,to_aid,otype" +
			") VALUES (" +
				"$1,$2,$3,TO_TIMESTAMP($4),$5,$6,$7,$8,$9"+
			")"
	_,err := sw.S.Db().Exec(query,
		evt.EvtId,
		evt.BlockNum,
		evt.TxId,
		evt.TimeStamp,
		contract_aid,
		evt.TokenId,
		from_aid,
		to_aid,
		otype,
	)
	if err != nil {
		sw.S.Log_msg(fmt.Sprintf("DB error: can't insert into rw_transfer table: %v\n",err))
		os.Exit(1)
	}
}
func (sw *SQLStorageWrapper) Offer_exists(contract_addr string,offer_id int64) bool {

	contract_aid,err := sw.S.Nonfatal_lookup_address_id(contract_addr)
	if err != nil {
		return false
	}
	var query string
	query = "SELECT id FROM rw_new_offer WHERE contract_aid=$1 AND offer_id=$2"
	var null_offer_id sql.NullInt64
	res := sw.S.Db().QueryRow(query,contract_aid,offer_id)
	err = res.Scan(&null_offer_id)
	if (err!=nil) {
		if err == sql.ErrNoRows {
			return false
		} else {
			sw.S.Log_msg(fmt.Sprintf("DB error: %v q=%v",err,query))
			os.Exit(1)
		}
	}
	return true
}
func (sw *SQLStorageWrapper) RWalk_token_exists(contract_addr string,token_id int64) bool {

	contract_aid,err := sw.S.Nonfatal_lookup_address_id(contract_addr)
	if err != nil {
		return false
	}
	var query string
	query = "SELECT id FROM rw_mint_evt WHERE contract_aid=$1 AND token_id=$2"
	var null_token_id sql.NullInt64
	res := sw.S.Db().QueryRow(query,contract_aid,token_id)
	err = res.Scan(&null_token_id)
	if (err!=nil) {
		if err == sql.ErrNoRows {
			return false
		} else {
			sw.S.Log_msg(fmt.Sprintf("DB error: %v q=%v",err,query))
			os.Exit(1)
		}
	}
	return true
}

// =============================================================================
// RANKING OPERATIONS
// =============================================================================

func (sw *SQLStorageWrapper) Update_randomwalk_top_profit_rank(aid int64,value float64,profit float64) int64 {

	var query string
	query = "UPDATE rw_uranks SET top_profit = $2,profit=$3 WHERE aid = $1"
	res,err:=sw.S.Db().Exec(query,aid,value,profit)
	if (err!=nil) {
		sw.S.Log_msg(fmt.Sprintf("Update_randomwalk_top_profit_rank() failed: %v, q=%v",err,query))
		os.Exit(1)
	}
	affected_rows,err:=res.RowsAffected()
	if err!=nil {
		sw.S.Log_msg(fmt.Sprintf("Error getting RowsAffected in update_top_profit_rank(): %v",err))
	}
	if affected_rows == 0 {
		query = "INSERT INTO rw_uranks(aid,top_profit,profit) VALUES($1,$2,$3)"
		_,err:=sw.S.Db().Exec(query,aid,value,profit)
		if (err!=nil) {
			sw.S.Log_msg(fmt.Sprintf("Update_randomwalk_top_profit_rank() failed: %v, q=%v",err,query))
		}
	}
	return affected_rows
}
func (sw *SQLStorageWrapper) Update_randomwalk_top_total_trades_rank(aid int64,value float64,total_trades int64) int64 {

	var query string
	query = "UPDATE rw_uranks SET top_trades = $2,total_trades=$3 WHERE aid = $1"
	res,err:=sw.S.Db().Exec(query,aid,value,total_trades)
	if (err!=nil) {
		sw.S.Log_msg(fmt.Sprintf("Uppdate_randomwalk_top_total_trades_rank() failed: %v, q=%v",err,query))
		os.Exit(1)
	}
	affected_rows,err:=res.RowsAffected()
	if err!=nil {
		sw.S.Log_msg(fmt.Sprintf("Error getting RowsAffected in update_top_total_trades_rank(): %v",err))
	}
	if affected_rows == 0 {
		query = "INSERT INTO rw_uranks(aid,top_trades,total_trades) VALUES($1,$2,$3)"
		_,err:=sw.S.Db().Exec(query,aid,value,total_trades)
		if (err!=nil) {
			sw.S.Log_msg(fmt.Sprintf("Update_randomwalk_top_total_trades_rank() failed: value=%v, err: %v, q=%v",value,err,query))
			os.Exit(1)
		}

	}
	return affected_rows
}
func (sw *SQLStorageWrapper) Update_randomwalk_top_volume_rank(aid int64,value float64,volume float64) int64 {

	var query string
	query = "UPDATE rw_uranks SET top_volume = $2,volume=$3 WHERE aid = $1"
	res,err:=sw.S.Db().Exec(query,aid,value,volume)
	if (err!=nil) {
		sw.S.Log_msg(fmt.Sprintf("Update_randomwalk_top_volume_rank() failed: %v, q=%v",err,query))
		os.Exit(1)
	}
	affected_rows,err:=res.RowsAffected()
	if err!=nil {
		sw.S.Log_msg(fmt.Sprintf("Error getting RowsAffected in Update_randomwalk_top_volume_rank(): %v",err))
	}
	if affected_rows == 0 {
		query = "INSERT INTO rw_uranks(aid,top_volume,volume) VALUES($1,$2,$3)"
		_,err:=sw.S.Db().Exec(query,aid,value,volume)
		if (err!=nil) {
			sw.S.Log_msg(fmt.Sprintf("Update_randomwalk_top_volume_rank() failed: value=%v, err: %v, q=%v",value,err,query))
			os.Exit(1)
		}

	}
	return affected_rows
}
func (sw *SQLStorageWrapper) Get_randomwalk_ranking_data_for_all_users() []rwp.RankStats {

	var query string
	query = "SELECT "+
				"user_aid," +
				"SUM(total_num_trades) AS tot_trades,"+
				"SUM(total_profit) AS profit,"+
				"SUM(total_vol) AS tot_volume "+
			"FROM rw_user_stats " +
			"GROUP BY user_aid"

	rows,err := sw.S.Db().Query(query)
	if (err!=nil) {
		sw.S.Log_msg(fmt.Sprintf("DB error: %v (query=%v)",err,query))
		os.Exit(1)
	}
	records := make([]rwp.RankStats,0,8)
	defer rows.Close()
	for rows.Next() {
		var rec rwp.RankStats
		err=rows.Scan(&rec.Aid,&rec.TotalTrades,&rec.ProfitLoss,&rec.VolumeTraded)
		records = append(records,rec)
	}
	return records
}
func (sw *SQLStorageWrapper) Get_randomwalk_top_profit_makers() []rwp.ProfitMaker {

	var query string
	query = "SELECT "+
				"a.addr,"+
				"r.top_profit,"+
				"r.profit/1e+18 " +
			"FROM rw_uranks AS r " +
				"LEFT JOIN address AS a ON r.aid = a.address_id " +
			"ORDER BY r.top_profit ASC,r.profit DESC LIMIT 100"

	rows,err := sw.S.Db().Query(query)
	if (err!=nil) {
		sw.S.Log_msg(fmt.Sprintf("DB error: %v (query=%v)",err,query))
		os.Exit(1)
	}
	records := make([]rwp.ProfitMaker,0,101)
	defer rows.Close()
	for rows.Next() {
		var rec rwp.ProfitMaker
		err=rows.Scan(&rec.Addr,&rec.Percentage,&rec.ProfitLoss)
		records = append(records,rec)
	}
	return records
}
func (sw *SQLStorageWrapper) Get_randomwalk_top_trade_makers() []rwp.TradeMaker {

	var query string
	query = "SELECT " +
				"a.addr," +
				"r.top_trades,"+
				"r.total_trades " +
			"FROM rw_uranks AS r " +
				"LEFT JOIN address AS a ON r.aid = a.address_id " +
			"ORDER BY r.top_trades ASC,r.total_trades DESC LIMIT 100"

	rows,err := sw.S.Db().Query(query)
	if (err!=nil) {
		sw.S.Log_msg(fmt.Sprintf("DB error: %v (query=%v)",err,query))
		os.Exit(1)
	}
	records := make([]rwp.TradeMaker,0,101)
	defer rows.Close()
	for rows.Next() {
		var rec rwp.TradeMaker
		err=rows.Scan(&rec.Addr,&rec.Percentage,&rec.TotalTrades)
		records = append(records,rec)
	}
	return records
}
func (sw *SQLStorageWrapper) Get_randomwalk_top_volume_makers() []rwp.VolumeMaker {

	var query string
	query = "SELECT "+
				"a.addr," +
				"r.top_volume,"+
				"r.volume/1e+18 "+
			"FROM rw_uranks AS r " +
				"LEFT JOIN address AS a ON r.aid = a.address_id " +
			"ORDER BY r.top_volume ASC,r.volume DESC LIMIT 100"

	rows,err := sw.S.Db().Query(query)
	if (err!=nil) {
		sw.S.Log_msg(fmt.Sprintf("DB error: %v (query=%v)",err,query))
		os.Exit(1)
	}
	records := make([]rwp.VolumeMaker,0,101)
	defer rows.Close()
	for rows.Next() {
		var rec rwp.VolumeMaker
		err=rows.Scan(&rec.Addr,&rec.Percentage,&rec.Volume)
		records = append(records,rec)
	}
	return records
}

// =============================================================================
// NOTIFICATION & MESSAGING
// =============================================================================

func (sw *SQLStorageWrapper) Get_mint_events_for_notification(rwalk_aid int64,start_ts int64) []rwp.NotificationEvent {

	records := make([]rwp.NotificationEvent,0,101)
	var query string
	query = "SELECT "+
				"EXTRACT(EPOCH FROM m.time_stamp)::BIGINT as ts,"+
				"token_id,"+
				"price/1e+18 AS price, "+
				"seed "+
			"FROM rw_mint_evt m " +
			"WHERE (contract_aid=$1) AND (time_stamp > TO_TIMESTAMP($2))  "
	rows,err := sw.S.Db().Query(query,rwalk_aid,start_ts)
	if (err!=nil) {
		sw.S.Log_msg(fmt.Sprintf("DB error: %v (query=%v)",err,query))
		os.Exit(1)
	}
	defer rows.Close()
	for rows.Next() {
		var rec rwp.NotificationEvent
		err=rows.Scan(
			&rec.TimeStampMinted,
			&rec.TokenId,
			&rec.Price,
			&rec.SeedHex,
		)
		records = append(records,rec)
	}
	return records
}
func (sw *SQLStorageWrapper) Get_messaging_status() rwp.MsgStatus  {

	var query string
	query = "SELECT last_tx_id,last_evtlog_id,last_block_num,last_timestamp "+
				"FROM rw_messaging_status"
	res := sw.S.Db().QueryRow(query)
	var output rwp.MsgStatus
	err := res.Scan(
		&output.TxId,
		&output.EvtLogId,
		&output.BlockNum,
		&output.TimeStamp,
	)
	if (err!=nil) {
		if err == sql.ErrNoRows {
			return output
		}
		sw.S.Log_msg(fmt.Sprintf("DB error: %v q=%v",err,query))
		os.Exit(1)
	}
	return output

}
func (sw *SQLStorageWrapper) Update_messaging_status(status *rwp.MsgStatus) {

	var query string
	query = "UPDATE rw_messaging_status SET "+
				"last_tx_id = $1, "+
				"last_evtlog_id = $2,"+
				"last_block_num = $3, "+
				"last_timestamp = $4 "

	_,err := sw.S.Db().Exec(query,status.TxId,status.EvtLogId,status.BlockNum,status.TimeStamp)
	if (err!=nil) {
		sw.S.Log_msg(fmt.Sprintf("DB error: %v q=%v",err,query))
		os.Exit(1)
	}
}
func (sw *SQLStorageWrapper) Get_all_events_for_notification(rwalk_aid int64,start_ts int64) []rwp.NotificationEvent {

	records := make([]rwp.NotificationEvent,0,101)
	var query string
	query = "SELECT "+
				"ts,"+
				"token_id,"+
				"price, "+
				"evt_type "+
			"FROM (" +
				"("+
					"SELECT "+
						"EXTRACT(EPOCH FROM m.time_stamp)::BIGINT as ts,"+
						"token_id,"+
						"price/1e+18 AS price, "+
						"1 AS evt_type " +
					"FROM rw_mint_evt m " +
						"WHERE (contract_aid=$1) AND (time_stamp > TO_TIMESTAMP($2))  "+
				") UNION ALL( "+
					"SELECT "+
						"EXTRACT(EPOCH FROM o.time_stamp)::BIGINT as ts,"+
						"token_id,"+
						"price/1e+18 AS price, "+
						"2 AS evt_type "+
					"FROM rw_new_offer o " +
						"WHERE (rwalk_aid=$1) AND (time_stamp > TO_TIMESTAMP($2)) AND (otype=1) " +
				") UNION ALL( "+
					"SELECT "+
						"EXTRACT(EPOCH FROM o.time_stamp)::BIGINT as ts,"+
						"token_id,"+
						"price/1e+18 AS price, "+
						"5 AS evt_type "+
					"FROM rw_new_offer o " +
						"WHERE (rwalk_aid=$1) AND (time_stamp > TO_TIMESTAMP($2)) AND (otype=0) " +
				") UNION ALL ( "+
					"SELECT "+
						"EXTRACT(EPOCH FROM b.time_stamp)::BIGINT as ts,"+
						"o.token_id,"+
						"o.price/1e+18 AS price, "+
						"3 AS evt_type " +
					"FROM rw_item_bought b " +
						"JOIN rw_new_offer o ON (b.contract_aid=o.contract_aid) AND (b.offer_id=o.offer_id) " +
					"WHERE (o.rwalk_aid=$1) AND (b.time_stamp > TO_TIMESTAMP($2))  " +
				") " +
			") data " +
			"ORDER BY ts"

	sw.S.Info.Printf("rwalk_aid=%v start_ts=%v q=%v\n",rwalk_aid,start_ts,query)
	rows,err := sw.S.Db().Query(query,rwalk_aid,start_ts)
	if (err!=nil) {
		sw.S.Log_msg(fmt.Sprintf("DB error: %v (query=%v)",err,query))
		os.Exit(1)
	}
	defer rows.Close()
	for rows.Next() {
		var rec rwp.NotificationEvent
		err=rows.Scan(
			&rec.TimeStampMinted,
			&rec.TokenId,
			&rec.Price,
			&rec.EvtType,
		)
		records = append(records,rec)
	}
	return records
}
func (sw *SQLStorageWrapper) Get_all_events_for_notification2(rwalk_aid int64,start_evtlog_id int64) []rwp.NotificationEvent2 {

	records := make([]rwp.NotificationEvent2,0,101)
	var query string
	query = "SELECT "+
				"ts,"+
				"tx_id,"+
				"evtlog_id,"+
				"token_id,"+
				"price, "+
				"evt_type "+
			"FROM (" +
				"("+
					"SELECT "+
						"EXTRACT(EPOCH FROM m.time_stamp)::BIGINT as ts,"+
						"m.tx_id,"+
						"m.evtlog_id,"+
						"token_id,"+
						"price/1e+18 AS price, "+
						"1 AS evt_type " +
					"FROM rw_mint_evt m " +
						"WHERE (contract_aid=$1) AND (m.evtlog_id>$2)  "+
				") UNION ALL( "+
					"SELECT "+
						"EXTRACT(EPOCH FROM o.time_stamp)::BIGINT as ts,"+
						"o.tx_id,"+
						"o.evtlog_id,"+
						"token_id,"+
						"price/1e+18 AS price, "+
						"2 AS evt_type "+
					"FROM rw_new_offer o " +
						"WHERE (rwalk_aid=$1) AND (o.evtlog_id>$2) AND (otype=1) " +
				") UNION ALL( "+
					"SELECT "+
						"EXTRACT(EPOCH FROM o.time_stamp)::BIGINT as ts,"+
						"o.tx_id,"+
						"o.evtlog_id,"+
						"token_id,"+
						"price/1e+18 AS price, "+
						"5 AS evt_type "+
					"FROM rw_new_offer o " +
						"WHERE (rwalk_aid=$1) AND (o.evtlog_id>$2) AND (otype=0) " +
				") UNION ALL ( "+
					"SELECT "+
						"EXTRACT(EPOCH FROM b.time_stamp)::BIGINT as ts,"+
						"b.tx_id,"+
						"b.evtlog_id,"+
						"o.token_id,"+
						"o.price/1e+18 AS price, "+
						"3 AS evt_type " +
					"FROM rw_item_bought b " +
						"JOIN rw_new_offer o ON (b.contract_aid=o.contract_aid) AND (b.offer_id=o.offer_id) " +
					"WHERE (o.rwalk_aid=$1) AND (b.evtlog_id>$2)  " +
				") " +
			") data " +
			"ORDER BY evtlog_id"
	sw.S.Info.Printf("rwalk_aid=%v start_evtlog_id=%v q=%v\n",rwalk_aid,start_evtlog_id,query)
	rows,err := sw.S.Db().Query(query,rwalk_aid,start_evtlog_id)
	if (err!=nil) {
		sw.S.Log_msg(fmt.Sprintf("DB error: %v (query=%v)",err,query))
		os.Exit(1)
	}
	defer rows.Close()
	for rows.Next() {
		var rec rwp.NotificationEvent2
		err=rows.Scan(
			&rec.TimeStampMinted,
			&rec.TxId,
			&rec.EvtLogId,
			&rec.TokenId,
			&rec.Price,
			&rec.EvtType,
		)
		records = append(records,rec)
	}
	return records
}
func (sw *SQLStorageWrapper) Get_all_events_for_notification_test(rwalk_aid int64,start_ts int64) []rwp.NotificationEvent {

	records := make([]rwp.NotificationEvent,0,101)
	var query string
	query = "SELECT "+
				"ts,"+
				"token_id,"+
				"price, "+
				"evt_type "+
			"FROM (" +
				"("+
					"SELECT "+
						"EXTRACT(EPOCH FROM m.time_stamp)::BIGINT as ts,"+
						"token_id,"+
						"price/1e+18 AS price, "+
						"1 AS evt_type " +
					"FROM rw_mint_evt m " +
						"WHERE (contract_aid=$1) AND (time_stamp > TO_TIMESTAMP($2))  "+
				") "+
				/*UNION ALL( "+
					"SELECT "+
						"EXTRACT(EPOCH FROM o.time_stamp)::BIGINT as ts,"+
						"token_id,"+
						"price/1e+18 AS price, "+
						"2 AS evt_type "+
					"FROM rw_new_offer o " +
						"WHERE (rwalk_aid=$1) AND (time_stamp > TO_TIMESTAMP($2))  " +
				") UNION ALL ( "+
					"SELECT "+
						"EXTRACT(EPOCH FROM b.time_stamp)::BIGINT as ts,"+
						"o.token_id,"+
						"o.price/1e+18 AS price, "+
						"3 AS evt_type " +
					"FROM rw_item_bought b " +
						"JOIN rw_new_offer o ON (b.contract_aid=o.contract_aid) AND (b.offer_id=o.offer_id) " +
					"WHERE (o.rwalk_aid=$1) AND (b.time_stamp > TO_TIMESTAMP($2))  " +
				") " +
				*/
			") data " +
			"ORDER BY ts"
	rows,err := sw.S.Db().Query(query,rwalk_aid,start_ts)
	if (err!=nil) {
		sw.S.Log_msg(fmt.Sprintf("DB error: %v (query=%v)",err,query))
		os.Exit(1)
	}
	defer rows.Close()
	for rows.Next() {
		var rec rwp.NotificationEvent
		err=rows.Scan(
			&rec.TimeStampMinted,
			&rec.TokenId,
			&rec.Price,
			&rec.EvtType,
		)
		records = append(records,rec)
	}
	return records
}
func (sw *SQLStorageWrapper) Get_server_timestamp() int64 {

	var query string
	query = "SELECT EXTRACT(EPOCH FROM now())::BIGINT AS ts"
	res := sw.S.Db().QueryRow(query)
	var null_ts sql.NullInt64
	err := res.Scan(&null_ts)
	if (err!=nil) {
		sw.S.Log_msg(fmt.Sprintf("DB error: %v q=%v",err,query))
		os.Exit(1)
	}
	return null_ts.Int64
}
func (sw *SQLStorageWrapper) Get_last_mint_timestamp() int64 {

	var query string
	query = "SELECT EXTRACT(EPOCH FROM time_stamp)::BIGINT AS ts FROM rw_mint_evt "+
			"ORDER BY id DESC LIMIT 1"
	res := sw.S.Db().QueryRow(query)
	var null_ts sql.NullInt64
	err := res.Scan(&null_ts)
	if (err!=nil) {
		if err == sql.ErrNoRows {
			return 0
		}
		sw.S.Log_msg(fmt.Sprintf("DB error: %v q=%v",err,query))
		os.Exit(1)
	}
	return null_ts.Int64
}
func (sw *SQLStorageWrapper) Get_rw_token_transfers_by_tx_hash(tx_hash string) []rwp.TransferEntry  {

	var query string
	query = "SELECT " +
				"fa.addr,ta.addr,token_id "+
			"FROM rw_transfer tr "+
				"JOIN address fa ON tr.from_aid=fa.address_id "+
				"JOIN address ta ON tr.to_aid=ta.address_id "+
				"JOIN transaction tx ON tx.id=tr.tx_id "+
			"WHERE tx.tx_hash=$1"

	rows,err := sw.S.Db().Query(query,tx_hash)
	if (err!=nil) {
		sw.S.Log_msg(fmt.Sprintf("DB error: %v (query=%v)",err,query))
		os.Exit(1)
	}
	records := make([]rwp.TransferEntry,0,8)
	defer rows.Close()
	for rows.Next() {
		var rec rwp.TransferEntry
		err=rows.Scan(&rec.From,&rec.To,&rec.TokenId)
		if (err!=nil) {
			sw.S.Log_msg(fmt.Sprintf("DB error: %v q=%v",err,query))
			os.Exit(1)
		}
		records = append(records,rec)
	}
	return records
}
