package dbs
import (
	"os"
	"fmt"
	"database/sql"
	p "github.com/PredictionExplorer/augur-explorer/primitives"
)
func (ss *SQLStorage) Get_randomwalk_processing_status() p.RandomWalkProcStatus {

	var output p.RandomWalkProcStatus
	var null_id,null_block sql.NullInt64

	var query string
	for {
		query = "SELECT last_evt_id,last_block FROM rw_proc_status"

		res := ss.db.QueryRow(query)
		err := res.Scan(&null_id,&null_block)
		if (err!=nil) {
			if err == sql.ErrNoRows {
				query = "INSERT INTO rw_proc_status DEFAULT VALUES"
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
		output.LastIdProcessed = null_id.Int64
	}
	if null_block.Valid {
		output.LastBlockNum = null_block.Int64
	}
	return output
}
func (ss *SQLStorage) Update_randomwalk_process_status(status *p.RandomWalkProcStatus) {

	var query string
	query = "UPDATE rw_proc_status SET last_evt_id = $1,last_block=$2"

	_,err := ss.db.Exec(query,status.LastIdProcessed,status.LastBlockNum)
	if (err!=nil) {
		ss.Log_msg(fmt.Sprintf("DB error: %v q=%v",err,query))
		os.Exit(1)
	}
}
func (ss *SQLStorage) Get_randomwalk_contract_addresses() p.RW_ContractAddresses {

	var output p.RW_ContractAddresses
	var query string
	query = "SELECT "+
				"marketplace_addr,randomwalk_addr," +
				"mp_a.address_id,rw_a.address_id "+
			"FROM rw_contracts rw " +
				"JOIN address mp_a ON rw.marketplace_addr=mp_a.addr " +
				"JOIN address rw_a ON rw.randomwalk_addr=rw_a.addr "

	res := ss.db.QueryRow(query)
	err := res.Scan(
		&output.MarketPlace,
		&output.RandomWalk,
		&output.MarketPlaceAid,
		&output.RandomWalkAid,
	)
	if (err!=nil) {
		if err == sql.ErrNoRows {
			ss.Log_msg("Can't find record in rw_contracts table for contract addresses")
			os.Exit(1)
		}
		ss.Log_msg(fmt.Sprintf("Get_randomwalk_contract_addresses() failed: %v, q=%v",err,query))
		os.Exit(1)
	}
	return output
}
func (ss *SQLStorage) Insert_new_offer(evt *p.RW_NewOffer) {

	contract_aid:=ss.Lookup_or_create_address(evt.Contract,evt.BlockNum,evt.TxId)
	rwalk_aid:=ss.Lookup_or_create_address(evt.RWalkAddr,evt.BlockNum,evt.TxId)
	buyer_aid:=ss.Lookup_or_create_address(evt.Buyer,evt.BlockNum,evt.TxId)
	seller_aid:=ss.Lookup_or_create_address(evt.Seller,evt.BlockNum,evt.TxId)
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
	_,err := ss.db.Exec(query,
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
		ss.Log_msg(fmt.Sprintf("DB error: can't insert into new_offer table: %v\n",err))
		os.Exit(1)
	}
}
func (ss *SQLStorage) Insert_item_bought(evt *p.RW_ItemBought) {

	contract_aid:=ss.Lookup_or_create_address(evt.Contract,evt.BlockNum,evt.TxId)
	seller_aid:=ss.Lookup_or_create_address(evt.SellerAddr,evt.BlockNum,evt.TxId)
	buyer_aid:=ss.Lookup_or_create_address(evt.BuyerAddr,evt.BlockNum,evt.TxId)
	var query string
	query = "INSERT INTO rw_item_bought(" +
				"evtlog_id,block_num,tx_id,time_stamp,contract_aid, "+
				"offer_id,seller_aid,buyer_aid" +
			") VALUES (" +
				"$1,$2,$3,TO_TIMESTAMP($4),$5,$6,$7,$8"+
			")"
	_,err := ss.db.Exec(query,
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
		ss.Log_msg(fmt.Sprintf("DB error: can't insert into item_bought table: %v\n",err))
		os.Exit(1)
	}

}
func (ss *SQLStorage) Insert_offer_canceled(evt *p.RW_OfferCanceled) {

	contract_aid:=ss.Lookup_or_create_address(evt.Contract,evt.BlockNum,evt.TxId)
	var query string
	query = "INSERT INTO rw_offer_canceled(" +
				"evtlog_id,block_num,tx_id,time_stamp,contract_aid, "+
				"offer_id" +
			") VALUES (" +
				"$1,$2,$3,TO_TIMESTAMP($4),$5,$6"+
			")"
	_,err := ss.db.Exec(query,
		evt.EvtId,
		evt.BlockNum,
		evt.TxId,
		evt.TimeStamp,
		contract_aid,
		evt.OfferId,
	)
	if err != nil {
		ss.Log_msg(fmt.Sprintf("DB error: can't insert into offer_canceled table: %v\n",err))
		os.Exit(1)
	}

}
func (ss *SQLStorage) Insert_withdrawal(evt *p.RW_Withdrawal) {

	contract_aid:=ss.Lookup_or_create_address(evt.Contract,evt.BlockNum,evt.TxId)
	aid:=ss.Lookup_or_create_address(evt.Destination,evt.BlockNum,evt.TxId)
	var query string
	query = "INSERT INTO rw_withdrawal(" +
				"evtlog_id,block_num,tx_id,time_stamp,contract_aid, "+
				"aid,token_id,amount" +
			") VALUES (" +
				"$1,$2,$3,TO_TIMESTAMP($4),$5,$6,$7,$8"+
			")"
	_,err := ss.db.Exec(query,
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
		ss.Log_msg(fmt.Sprintf("DB error: can't insert into rw_withdrawal table: %v\n",err))
		os.Exit(1)
	}

}
func (ss *SQLStorage) Insert_token_name(evt *p.RW_TokenName) {

	contract_aid:=ss.Lookup_or_create_address(evt.Contract,evt.BlockNum,evt.TxId)
	var query string
	query = "INSERT INTO rw_token_name(" +
				"evtlog_id,block_num,tx_id,time_stamp,contract_aid, "+
				"token_id,new_name" +
			") VALUES (" +
				"$1,$2,$3,TO_TIMESTAMP($4),$5,$6,$7"+
			")"
	_,err := ss.db.Exec(query,
		evt.EvtId,
		evt.BlockNum,
		evt.TxId,
		evt.TimeStamp,
		contract_aid,
		evt.TokenId,
		evt.NewName,
	)
	if err != nil {
		ss.Log_msg(fmt.Sprintf("DB error: can't insert into rw_token_name table: %v\n",err))
		os.Exit(1)
	}

}
func (ss *SQLStorage) Insert_mint_event(evt *p.RW_MintEvent) {

	contract_aid:=ss.Lookup_or_create_address(evt.Contract,evt.BlockNum,evt.TxId)
	owner_aid:=ss.Lookup_or_create_address(evt.Owner,evt.BlockNum,evt.TxId)
	var query string
	query = "INSERT INTO rw_mint_evt(" +
				"evtlog_id,block_num,tx_id,time_stamp,contract_aid, "+
				"token_id,owner_aid,seed,seed_num,price" +
			") VALUES (" +
				"$1,$2,$3,TO_TIMESTAMP($4),$5,$6,$7,$8,$9,$10"+
			")"
	_,err := ss.db.Exec(query,
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
		ss.Log_msg(fmt.Sprintf("DB error: can't insert into rw_mint_evt table: %v\n",err))
		os.Exit(1)
	}
}
func (ss *SQLStorage) Insert_token_transfer_event(evt *p.RW_Transfer) {

	contract_aid:=ss.Lookup_or_create_address(evt.Contract,evt.BlockNum,evt.TxId)
	from_aid:=ss.Lookup_or_create_address(evt.From,evt.BlockNum,evt.TxId)
	to_aid:=ss.Lookup_or_create_address(evt.To,evt.BlockNum,evt.TxId)
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
	_,err := ss.db.Exec(query,
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
		ss.Log_msg(fmt.Sprintf("DB error: can't insert into rw_transfer table: %v\n",err))
		os.Exit(1)
	}
}
func (ss *SQLStorage) Offer_exists(contract_addr string,offer_id int64) bool {

	contract_aid,err := ss.Nonfatal_lookup_address_id(contract_addr)
	if err != nil {
		return false
	}
	var query string
	query = "SELECT id FROM rw_new_offer WHERE contract_aid=$1 AND offer_id=$2"
	var null_offer_id sql.NullInt64
	res := ss.db.QueryRow(query,contract_aid,offer_id)
	err = res.Scan(&null_offer_id)
	if (err!=nil) {
		if err == sql.ErrNoRows {
			return false
		} else {
			ss.Log_msg(fmt.Sprintf("DB error: %v q=%v",err,query))
			os.Exit(1)
		}
	}
	return true
}
func (ss *SQLStorage) RWalk_token_exists(contract_addr string,token_id int64) bool {

	contract_aid,err := ss.Nonfatal_lookup_address_id(contract_addr)
	if err != nil {
		return false
	}
	var query string
	query = "SELECT id FROM rw_mint_evt WHERE contract_aid=$1 AND token_id=$2"
	var null_token_id sql.NullInt64
	res := ss.db.QueryRow(query,contract_aid,token_id)
	err = res.Scan(&null_token_id)
	if (err!=nil) {
		if err == sql.ErrNoRows {
			return false
		} else {
			ss.Log_msg(fmt.Sprintf("DB error: %v q=%v",err,query))
			os.Exit(1)
		}
	}
	return true
}
func (ss *SQLStorage) Update_randomwalk_top_profit_rank(aid int64,value float64,profit float64) int64 {

	var query string
	query = "UPDATE rw_uranks SET top_profit = $2,profit=$3 WHERE aid = $1"
	res,err:=ss.db.Exec(query,aid,value,profit)
	if (err!=nil) {
		ss.Log_msg(fmt.Sprintf("Update_randomwalk_top_profit_rank() failed: %v, q=%v",err,query))
		os.Exit(1)
	}
	affected_rows,err:=res.RowsAffected()
	if err!=nil {
		ss.Log_msg(fmt.Sprintf("Error getting RowsAffected in update_top_profit_rank(): %v",err))
	}
	if affected_rows == 0 {
		query = "INSERT INTO rw_uranks(aid,top_profit,profit) VALUES($1,$2,$3)"
		_,err:=ss.db.Exec(query,aid,value,profit)
		if (err!=nil) {
			ss.Log_msg(fmt.Sprintf("Update_randomwalk_top_profit_rank() failed: %v, q=%v",err,query))
		}
	}
	return affected_rows
}
func (ss *SQLStorage) Update_randomwalk_top_total_trades_rank(aid int64,value float64,total_trades int64) int64 {

	var query string
	query = "UPDATE rw_uranks SET top_trades = $2,total_trades=$3 WHERE aid = $1"
	res,err:=ss.db.Exec(query,aid,value,total_trades)
	if (err!=nil) {
		ss.Log_msg(fmt.Sprintf("Uppdate_randomwalk_top_total_trades_rank() failed: %v, q=%v",err,query))
		os.Exit(1)
	}
	affected_rows,err:=res.RowsAffected()
	if err!=nil {
		ss.Log_msg(fmt.Sprintf("Error getting RowsAffected in update_top_total_trades_rank(): %v",err))
	}
	if affected_rows == 0 {
		query = "INSERT INTO rw_uranks(aid,top_trades,total_trades) VALUES($1,$2,$3)"
		_,err:=ss.db.Exec(query,aid,value,total_trades)
		if (err!=nil) {
			ss.Log_msg(fmt.Sprintf("Update_randomwalk_top_total_trades_rank() failed: value=%v, err: %v, q=%v",value,err,query))
			os.Exit(1)
		}

	}
	return affected_rows
}
func (ss *SQLStorage) Update_randomwalk_top_volume_rank(aid int64,value float64,volume float64) int64 {

	var query string
	query = "UPDATE rw_uranks SET top_volume = $2,volume=$3 WHERE aid = $1"
	res,err:=ss.db.Exec(query,aid,value,volume)
	if (err!=nil) {
		ss.Log_msg(fmt.Sprintf("Update_randomwalk_top_volume_rank() failed: %v, q=%v",err,query))
		os.Exit(1)
	}
	affected_rows,err:=res.RowsAffected()
	if err!=nil {
		ss.Log_msg(fmt.Sprintf("Error getting RowsAffected in Update_randomwalk_top_volume_rank(): %v",err))
	}
	if affected_rows == 0 {
		query = "INSERT INTO rw_uranks(aid,top_volume,volume) VALUES($1,$2,$3)"
		_,err:=ss.db.Exec(query,aid,value,volume)
		if (err!=nil) {
			ss.Log_msg(fmt.Sprintf("Update_randomwalk_top_volume_rank() failed: value=%v, err: %v, q=%v",value,err,query))
			os.Exit(1)
		}

	}
	return affected_rows
}
func (ss *SQLStorage) Get_randomwalk_ranking_data_for_all_users() []p.RankStats {

	var query string
	query = "SELECT "+
				"user_aid," +
				"SUM(total_num_trades) AS tot_trades,"+
				"SUM(total_profit) AS profit,"+
				"SUM(total_vol) AS tot_volume "+
			"FROM rw_user_stats " +
			"GROUP BY user_aid"

	rows,err := ss.db.Query(query)
	if (err!=nil) {
		ss.Log_msg(fmt.Sprintf("DB error: %v (query=%v)",err,query))
		os.Exit(1)
	}
	records := make([]p.RankStats,0,8)
	defer rows.Close()
	for rows.Next() {
		var rec p.RankStats
		err=rows.Scan(&rec.Aid,&rec.TotalTrades,&rec.ProfitLoss,&rec.VolumeTraded)
		records = append(records,rec)
	}
	return records
}
func (ss *SQLStorage) Get_randomwalk_top_profit_makers() []p.ProfitMaker {

	var query string
	query = "SELECT "+
				"a.addr,"+
				"r.top_profit,"+
				"r.profit/1e+18 " +
			"FROM rw_uranks AS r " +
				"LEFT JOIN address AS a ON r.aid = a.address_id " +
			"ORDER BY r.top_profit ASC,r.profit DESC LIMIT 100"

	rows,err := ss.db.Query(query)
	if (err!=nil) {
		ss.Log_msg(fmt.Sprintf("DB error: %v (query=%v)",err,query))
		os.Exit(1)
	}
	records := make([]p.ProfitMaker,0,101)
	defer rows.Close()
	for rows.Next() {
		var rec p.ProfitMaker
		err=rows.Scan(&rec.Addr,&rec.Percentage,&rec.ProfitLoss)
		records = append(records,rec)
	}
	return records
}
func (ss *SQLStorage) Get_randomwalk_top_trade_makers() []p.TradeMaker {

	var query string
	query = "SELECT " +
				"a.addr," +
				"r.top_trades,"+
				"r.total_trades " +
			"FROM rw_uranks AS r " +
				"LEFT JOIN address AS a ON r.aid = a.address_id " +
			"ORDER BY r.top_trades ASC,r.total_trades DESC LIMIT 100"

	rows,err := ss.db.Query(query)
	if (err!=nil) {
		ss.Log_msg(fmt.Sprintf("DB error: %v (query=%v)",err,query))
		os.Exit(1)
	}
	records := make([]p.TradeMaker,0,101)
	defer rows.Close()
	for rows.Next() {
		var rec p.TradeMaker
		err=rows.Scan(&rec.Addr,&rec.Percentage,&rec.TotalTrades)
		records = append(records,rec)
	}
	return records
}
func (ss *SQLStorage) Get_randomwalk_top_volume_makers() []p.VolumeMaker {

	var query string
	query = "SELECT "+
				"a.addr," +
				"r.top_volume,"+
				"r.volume/1e+18 "+
			"FROM rw_uranks AS r " +
				"LEFT JOIN address AS a ON r.aid = a.address_id " +
			"ORDER BY r.top_volume ASC,r.volume DESC LIMIT 100"

	rows,err := ss.db.Query(query)
	if (err!=nil) {
		ss.Log_msg(fmt.Sprintf("DB error: %v (query=%v)",err,query))
		os.Exit(1)
	}
	records := make([]p.VolumeMaker,0,101)
	defer rows.Close()
	for rows.Next() {
		var rec p.VolumeMaker
		err=rows.Scan(&rec.Addr,&rec.Percentage,&rec.Volume)
		records = append(records,rec)
	}
	return records
}
func (ss *SQLStorage) Get_mint_events_for_notification(rwalk_aid int64,start_ts int64) []p.RW_NotificationEvent {

	records := make([]p.RW_NotificationEvent,0,101)
	var query string
	query = "SELECT "+
				"EXTRACT(EPOCH FROM m.time_stamp)::BIGINT as ts,"+
				"token_id,"+
				"price/1e+18 AS price, "+
				"seed "+
			"FROM rw_mint_evt m " +
			"WHERE (contract_aid=$1) AND (time_stamp > TO_TIMESTAMP($2))  "
	rows,err := ss.db.Query(query,rwalk_aid,start_ts)
	if (err!=nil) {
		ss.Log_msg(fmt.Sprintf("DB error: %v (query=%v)",err,query))
		os.Exit(1)
	}
	defer rows.Close()
	for rows.Next() {
		var rec p.RW_NotificationEvent
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
func (ss *SQLStorage) Get_messaging_status() p.RW_MsgStatus  {

	var query string
	query = "SELECT last_tx_id,last_evtlog_id,last_block_num,last_timestamp "+
				"FROM rw_messaging_status"
	res := ss.db.QueryRow(query)
	var output p.RW_MsgStatus
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
		ss.Log_msg(fmt.Sprintf("DB error: %v q=%v",err,query))
		os.Exit(1)
	}
	return output

}
func (ss *SQLStorage) Update_messaging_status(status *p.RW_MsgStatus) {

	var query string
	query = "UPDATE rw_messaging_status SET "+
				"last_tx_id = $1, "+
				"last_evtlog_id = $2,"+
				"last_block_num = $3, "+
				"last_timestamp = $4 "

	_,err := ss.db.Exec(query,status.TxId,status.EvtLogId,status.BlockNum,status.TimeStamp)
	if (err!=nil) {
		ss.Log_msg(fmt.Sprintf("DB error: %v q=%v",err,query))
		os.Exit(1)
	}
}
func (ss *SQLStorage) Get_all_events_for_notification(rwalk_aid int64,start_ts int64) []p.RW_NotificationEvent {

	records := make([]p.RW_NotificationEvent,0,101)
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

	ss.Info.Printf("rwalk_aid=%v start_ts=%v q=%v\n",rwalk_aid,start_ts,query)
	rows,err := ss.db.Query(query,rwalk_aid,start_ts)
	if (err!=nil) {
		ss.Log_msg(fmt.Sprintf("DB error: %v (query=%v)",err,query))
		os.Exit(1)
	}
	defer rows.Close()
	for rows.Next() {
		var rec p.RW_NotificationEvent
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
func (ss *SQLStorage) Get_all_events_for_notification2(rwalk_aid int64,start_evtlog_id int64) []p.RW_NotificationEvent2 {

	records := make([]p.RW_NotificationEvent2,0,101)
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
	ss.Info.Printf("rwalk_aid=%v start_evtlog_id=%v q=%v\n",rwalk_aid,start_evtlog_id,query)
	rows,err := ss.db.Query(query,rwalk_aid,start_evtlog_id)
	if (err!=nil) {
		ss.Log_msg(fmt.Sprintf("DB error: %v (query=%v)",err,query))
		os.Exit(1)
	}
	defer rows.Close()
	for rows.Next() {
		var rec p.RW_NotificationEvent2
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
func (ss *SQLStorage) Get_all_events_for_notification_test(rwalk_aid int64,start_ts int64) []p.RW_NotificationEvent {

	records := make([]p.RW_NotificationEvent,0,101)
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
	rows,err := ss.db.Query(query,rwalk_aid,start_ts)
	if (err!=nil) {
		ss.Log_msg(fmt.Sprintf("DB error: %v (query=%v)",err,query))
		os.Exit(1)
	}
	defer rows.Close()
	for rows.Next() {
		var rec p.RW_NotificationEvent
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
func (ss *SQLStorage) Get_server_timestamp() int64 {

	var query string
	query = "SELECT EXTRACT(EPOCH FROM now())::BIGINT AS ts"
	res := ss.db.QueryRow(query)
	var null_ts sql.NullInt64
	err := res.Scan(&null_ts)
	if (err!=nil) {
		ss.Log_msg(fmt.Sprintf("DB error: %v q=%v",err,query))
		os.Exit(1)
	}
	return null_ts.Int64
}
func (ss *SQLStorage) Get_last_mint_timestamp() int64 {

	var query string
	query = "SELECT EXTRACT(EPOCH FROM time_stamp)::BIGINT AS ts FROM rw_mint_evt "+
			"ORDER BY id DESC LIMIT 1"
	res := ss.db.QueryRow(query)
	var null_ts sql.NullInt64
	err := res.Scan(&null_ts)
	if (err!=nil) {
		if err == sql.ErrNoRows {
			return 0
		}
		ss.Log_msg(fmt.Sprintf("DB error: %v q=%v",err,query))
		os.Exit(1)
	}
	return null_ts.Int64
}
func (ss *SQLStorage) Get_rw_token_transfers_by_tx_hash(tx_hash string) []p.RW_TransferEntry  {

	var query string
	query = "SELECT " +
				"fa.addr,ta.addr,token_id "+
			"FROM rw_transfer tr "+
				"JOIN address fa ON tr.from_aid=fa.address_id "+
				"JOIN address ta ON tr.to_aid=ta.address_id "+
				"JOIN transaction tx ON tx.id=tr.tx_id "+
			"WHERE tx.tx_hash=$1"

	rows,err := ss.db.Query(query,tx_hash)
	if (err!=nil) {
		ss.Log_msg(fmt.Sprintf("DB error: %v (query=%v)",err,query))
		os.Exit(1)
	}
	records := make([]p.RW_TransferEntry,0,8)
	defer rows.Close()
	for rows.Next() {
		var rec p.RW_TransferEntry
		err=rows.Scan(&rec.From,&rec.To,&rec.TokenId)
		if (err!=nil) {
			ss.Log_msg(fmt.Sprintf("DB error: %v q=%v",err,query))
			os.Exit(1)
		}
		records = append(records,rec)
	}
	return records
}
