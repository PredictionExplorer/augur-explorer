package dbs
import (
	"os"
	"fmt"
	"time"
	"math"
	"database/sql"
	p "github.com/PredictionExplorer/augur-explorer/primitives"
)
func (ss *SQLStorage) Get_active_offers(rwalk_aid int64,market_aid int64, order_by int) []p.RW_API_Offer {

	records := make([]p.RW_API_Offer,0,16)

	var order_by_mod string =" ORDER BY o.id"
	if order_by == 1 {
		order_by_mod = " ORDER BY o.price DESC"
	}
	if order_by == 2 {
		order_by_mod = " ORDER BY o.price ASC"
	}
	var where_cond string =""
	where_cond = fmt.Sprintf(" AND (o.rwalk_aid=%v) ",rwalk_aid)
	where_cond += fmt.Sprintf(" AND (o.contract_aid=%v) ",market_aid)
	var query string
	query = "SELECT " +
				"o.id,"+
				"o.evtlog_id,"+
				"o.block_num,"+
				"o.tx_id, "+
				"tx.tx_hash,"+
				"EXTRACT(EPOCH FROM o.time_stamp)::BIGINT as ts,"+
				"o.time_stamp," +
				"o.offer_id,"+
				"o.otype,"+
				"o.seller_aid,"+
				"sa.addr seller_addr,"+
				"o.buyer_aid,"+
				"ba.addr buyer_addr,"+
				"o.token_id,"+
				"o.active,"+
				"o.price/1e+18 price, "+
				"o.rwalk_aid " +
			"FROM "+
				"rw_new_offer o "+
				"JOIN transaction tx ON o.tx_id=tx.id "+
				"JOIN address sa ON o.seller_aid=sa.address_id "+
				"JOIN address ba ON o.buyer_aid=ba.address_id "+
			"WHERE (active = 't')" + where_cond +
			order_by_mod

	rows,err := ss.db.Query(query,)
	if (err!=nil) {
		ss.Log_msg(fmt.Sprintf("DB error: %v (query=%v)",err,query))
		os.Exit(1)
	}

	defer rows.Close()
	for rows.Next() {
		var rec p.RW_API_Offer
		err=rows.Scan(
			&rec.Id,
			&rec.EvtLogId,
			&rec.BlockNum,
			&rec.TxId,
			&rec.TxHash,
			&rec.TimeStamp,
			&rec.DateTime,
			&rec.OfferId,
			&rec.OfferType,
			&rec.SellerAid,
			&rec.SellerAddr,
			&rec.BuyerAid,
			&rec.BuyerAddr,
			&rec.TokenId,
			&rec.Active,
			&rec.Price,
			&rec.RWalkAid,
		)
		if err!=nil {
			ss.Log_msg(fmt.Sprintf("DB error: %v, q=%v",err,query))
			os.Exit(1)
		}
		records = append(records,rec)
	}

	return records
}
func (ss *SQLStorage) Get_minted_tokens_by_period(rwalk_aid int64,ini_ts,fin_ts int) []p.RW_API_TokenMint {

	records := make([]p.RW_API_TokenMint,0,32)
	var query string
	query = "SELECT "+
				"t.block_num,"+
				"EXTRACT(EPOCH FROM t.time_stamp)::BIGINT as ts,"+
				"t.time_stamp," +
				"t.contract_aid,"+
				"ca.addr,"+
				"t.token_id,"+
				"t.owner_aid,"+
				"oa.addr minter_addr,"+
				"seed seed_hex,"+
				"seed_num,"+
				"price/1e+18,"+
				"tx.tx_hash "+
			"FROM rw_mint_evt t "+
				"LEFT JOIN address ca ON t.contract_aid=ca.address_id "+
				"LEFT JOIN address oa ON t.owner_aid=oa.address_id "+
				"LEFT JOIN transaction tx ON t.tx_id=tx.id "+
			"WHERE (t.time_stamp >= TO_TIMESTAMP($1)) AND (t.time_stamp<TO_TIMESTAMP($2)) " +
				"AND t.contract_aid=$3"
	ss.Info.Printf("rwalk_aid=%v ini=%v,fin=%v, q=%v\n",rwalk_aid,ini_ts,fin_ts,query)
	rows,err := ss.db.Query(query,ini_ts,fin_ts,rwalk_aid)
	if (err!=nil) {
		ss.Log_msg(fmt.Sprintf("DB error: %v (query=%v)",err,query))
		os.Exit(1)
	}

	defer rows.Close()
	for rows.Next() {
		var rec p.RW_API_TokenMint
		err=rows.Scan(
			&rec.BlockNum,
			&rec.TimeStamp,
			&rec.DateTime,
			&rec.ContractAid,
			&rec.ContractAddr,
			&rec.TokenId,
			&rec.MinterAid,
			&rec.MinterAddr,
			&rec.Seed,
			&rec.SeedNum,
			&rec.Price,
			&rec.TxHash,
		)
		if err!=nil {
			ss.Log_msg(fmt.Sprintf("DB error: %v, q=%v",err,query))
			os.Exit(1)
		}
		records = append(records,rec)
	}

	return records
}
func (ss *SQLStorage) Get_minted_tokens_sequentially(rwalk_aid int64,offset,limit int) []p.RW_API_TokenMint {

	records := make([]p.RW_API_TokenMint,0,32)
	var query string
	query = "SELECT "+
				"t.block_num,"+
				"EXTRACT(EPOCH FROM t.time_stamp)::BIGINT as ts,"+
				"t.time_stamp," +
				"t.contract_aid,"+
				"ca.addr,"+
				"t.token_id,"+
				"t.owner_aid,"+
				"oa.addr minter_addr,"+
				"seed seed_hex,"+
				"seed_num,"+
				"price/1e+18,"+
				"tx.tx_hash "+
			"FROM rw_mint_evt t "+
				"LEFT JOIN address ca ON t.contract_aid=ca.address_id "+
				"LEFT JOIN address oa ON t.owner_aid=oa.address_id "+
				"LEFT JOIN transaction tx ON t.tx_id=tx.id "+
			"WHERE contract_aid=$1 "+
			"ORDER by t.id DESC "+
			"OFFSET $2 LIMIT $3"

	rows,err := ss.db.Query(query,rwalk_aid,offset,limit)
	if (err!=nil) {
		ss.Log_msg(fmt.Sprintf("DB error: %v (query=%v)",err,query))
		os.Exit(1)
	}

	defer rows.Close()
	for rows.Next() {
		var rec p.RW_API_TokenMint
		err=rows.Scan(
			&rec.BlockNum,
			&rec.TimeStamp,
			&rec.DateTime,
			&rec.ContractAid,
			&rec.ContractAddr,
			&rec.TokenId,
			&rec.MinterAid,
			&rec.MinterAddr,
			&rec.Seed,
			&rec.SeedNum,
			&rec.Price,
			&rec.TxHash,
		)
		if err!=nil {
			ss.Log_msg(fmt.Sprintf("DB error: %v, q=%v",err,query))
			os.Exit(1)
		}
		records = append(records,rec)
	}

	return records
}
func (ss *SQLStorage) Get_trading_history(contract_aid int64,offset,limit int) []p.RW_API_TradingHistoryLog {

	records := make([]p.RW_API_TradingHistoryLog,0,16)
	/*var where_condition string
	where_condition =  fmt.Sprintf(" WHERE contract_aid=%v ",contract_aid)
	if contract_aid != 0 {
		where_condition = ""
	}*/

	var query string
	query = "SELECT " +
				"record_id,"+"evtlog_id,"+"block_num,"+"tx_id, "+"offer_ts,"+"offer_date," +
				"offer_id,"+"otype,"+"seller_aid,"+"seller_addr,"+"buyer_aid,"+"buyer_addr,"+
				"token_id,"+"active,"+"cancel_id,"+"price, "+"profit, "+"contract_aid," +"contract_addr," +
				"rwalk_aid,"+"rwalk_addr, "+"itembought_ts,"+"itembought_date, "+
				"canceled_ts,"+	"canceled_date, "+	"real_ts, "+"real_date "+
			"FROM (" +
				"(" +
					"SELECT " +
						"o.id AS record_id,"+
						"o.evtlog_id AS evtlog_id,"+
						"o.block_num AS block_num,"+
						"o.tx_id AS tx_id, "+
						"EXTRACT(EPOCH FROM o.time_stamp)::BIGINT as offer_ts,"+
						"o.time_stamp offer_date," +"o.offer_id offer_id,"+	"o.otype,"+
						"o.seller_aid,"+"sa.addr seller_addr,"+	"o.buyer_aid,"+	"ba.addr buyer_addr,"+
						"o.token_id,"+"o.active,"+"NULL AS cancel_id,"+"o.price/1e+18 price, "+
						"o.profit/1e+18 profit, "+	"o.contract_aid," +	"ca.addr contract_addr," +
						"o.rwalk_aid,"+	"rwa.addr rwalk_addr, "+
						"NULL AS itembought_ts,"+
						"NULL AS itembought_date, "+
						"NULL AS canceled_ts,"+
						"NULL AS canceled_date, "+
						"EXTRACT(EPOCH FROM o.time_stamp)::BIGINT  AS real_ts, "+
						"o.time_stamp AS real_date "+
					"FROM "+
						"rw_new_offer o "+
						"JOIN transaction tx ON o.tx_id=tx.id "+
						"JOIN address sa ON o.seller_aid=sa.address_id "+
						"JOIN address ba ON o.buyer_aid=ba.address_id "+
						"JOIN address ca ON o.contract_aid=ca.address_id "+
						"JOIN address rwa ON o.rwalk_aid=rwa.address_id "+
					"WHERE o.contract_aid=$3 "+
				") UNION ALL ("+
					"SELECT " +
						"COALESCE(ib.id,can.id,o.id) AS record_id,"+
						"COALESCE(ib.evtlog_id,can.evtlog_id,o.evtlog_id) AS evtlog_id,"+
						"COALESCE(ib.block_num,can.block_num,o.block_num) AS block_num,"+
						"COALESCE(ib.tx_id,can.tx_id,o.tx_id) AS tx_id, "+
						"EXTRACT(EPOCH FROM o.time_stamp)::BIGINT as offer_ts,"+
						"o.time_stamp offer_date," +"o.offer_id offer_id,"+	"o.otype,"+
						"o.seller_aid,"+"sa.addr seller_addr,"+	"o.buyer_aid,"+	"ba.addr buyer_addr,"+
						"o.token_id,"+"o.active,"+	"can.id cancel_id,"+"o.price/1e+18 price, "+
						"o.profit/1e+18 profit, "+"o.contract_aid," +"ca.addr contract_addr," +
						"o.rwalk_aid,"+	"rwa.addr rwalk_addr, "+
						"EXTRACT(EPOCH FROM ib.time_stamp)::BIGINT as itembought_ts,"+
						"ib.time_stamp itembought_date, "+
						"EXTRACT(EPOCH FROM can.time_stamp)::BIGINT as canceled_ts,"+
						"can.time_stamp canceled_date, "+
						"COALESCE(" +
							"EXTRACT(EPOCH FROM ib.time_stamp)::BIGINT,"+
							"EXTRACT(EPOCH FROM can.time_stamp)::BIGINT,"+
							"EXTRACT(EPOCH FROM o.time_stamp)::BIGINT"+
						") AS real_ts, "+
						"COALESCE(" +
							"ib.time_stamp,"+
							"can.time_stamp,"+
							"o.time_stamp"+
						") AS real_date "+
					"FROM "+
						"rw_new_offer o "+
						"JOIN transaction tx ON o.tx_id=tx.id "+
						"JOIN address sa ON o.seller_aid=sa.address_id "+
						"JOIN address ba ON o.buyer_aid=ba.address_id "+
						"JOIN address ca ON o.contract_aid=ca.address_id "+
						"JOIN address rwa ON o.rwalk_aid=rwa.address_id "+
						"LEFT JOIN rw_offer_canceled can ON (can.contract_aid=o.contract_aid) AND (can.offer_id=o.offer_id) "+
						"LEFT JOIN rw_item_bought ib ON (ib.contract_aid=o.contract_aid) AND (ib.offer_id=o.offer_id) "+
					 "WHERE o.contract_aid=$3 AND ((can.id IS NOT NULL) OR (ib.id IS NOT NULL)) "+
				")" +
			") recs "+
			"ORDER BY real_ts " +
			"OFFSET $1 LIMIT $2"

	rows,err := ss.db.Query(query,offset,limit,contract_aid)
	if (err!=nil) {
		ss.Log_msg(fmt.Sprintf("DB error: %v (query=%v)",err,query))
		os.Exit(1)
	}

	defer rows.Close()
	for rows.Next() {
		var rec p.RW_API_TradingHistoryLog
		var null_profit sql.NullFloat64
		var null_can_id sql.NullInt64
		var null_bought_ts,null_cancel_ts sql.NullInt64
		var null_bought_date,null_cancel_date sql.NullString
		err=rows.Scan(
			&rec.Id,
			&rec.EvtLogId,
			&rec.BlockNum,
			&rec.TxId,
			&rec.TimeStamp,
			&rec.DateTime,
			&rec.OfferId,
			&rec.OfferType,
			&rec.SellerAid,
			&rec.SellerAddr,
			&rec.BuyerAid,
			&rec.BuyerAddr,
			&rec.TokenId,
			&rec.Active,
			&null_can_id,
			&rec.Price,
			&null_profit,
			&rec.ContractAid,
			&rec.ContractAddr,
			&rec.RWalkAid,
			&rec.RWalkAddr,
			&null_bought_ts,
			&null_bought_date,
			&null_cancel_ts,
			&null_cancel_date,
			&rec.RealTs,
			&rec.RealDate,
		)
		if err!=nil {
			ss.Log_msg(fmt.Sprintf("DB error: %v, q=%v",err,query))
			os.Exit(1)
		}
		if null_profit.Valid {
			rec.Profit = null_profit.Float64
		} else {
			rec.Profit = math.NaN()
		}
		if null_can_id.Valid {rec.WasCanceled = true}
		if null_cancel_ts.Valid {
			rec.CanceledTs = null_cancel_ts.Int64
			time_canceled := time.Unix(int64(rec.CanceledTs),0)
			time_offered := time.Unix(int64(rec.TimeStamp),0)
			rec.CanceledDuration = p.DurationToString(p.TimeDifference(time_offered,time_canceled))
			fmt.Printf(
				"id=%v: offer_id=%v , canceled ts = %v , offered_Ts=%v rec.CanceledDuration=%v\n",
				rec.Id,rec.OfferId,rec.CanceledTs,rec.TimeStamp,rec.CanceledDuration,
			)
		}
		if null_cancel_date.Valid { rec.CanceledDate = null_cancel_date.String }
		if null_bought_ts.Valid {
			rec.WasBought = true
			rec.ItemBoughtTs = null_bought_ts.Int64
			time_bought := time.Unix(int64(rec.ItemBoughtTs),0)
			time_offered := time.Unix(int64(rec.TimeStamp),0)
			rec.BoughtDuration = p.DurationToString(p.TimeDifference(time_offered,time_bought))
			fmt.Printf(
				"id=%v: offer_id=%v , bought_ts = %v , offered_Ts=%v rec.BoughtDuration=%v\n",
				rec.Id,rec.OfferId,rec.ItemBoughtTs,rec.TimeStamp,rec.BoughtDuration,
			)
		}
		if null_bought_date.Valid { rec.ItemBoughtDate = null_bought_date.String }
		records = append(records,rec)
	}

	return records
}
func (ss *SQLStorage) Get_random_walk_stats(rwalk_aid int64) p.RW_API_RWalkStats {

	var output p.RW_API_RWalkStats
	var query string
	query = "SELECT " +
				"total_vol/1e+18,"+
				"total_num_trades,"+
				"total_num_toks,"+
				"total_withdrawals, "+
				"money_accumulated/2e+18 withdrawal_amount "+
			"FROM "+
				"rw_stats " +
			"WHERE rwalk_aid = $1"

	res := ss.db.QueryRow(query,rwalk_aid)
	err := res.Scan(
		&output.TradingVol,
		&output.NumTrades,
		&output.TokensMinted,
		&output.NumWithdrawals,
		&output.WithdrawalAmount,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return output
		} else {
			ss.Log_msg(fmt.Sprintf("DB error: %v, q=%v",err,query))
			os.Exit(1)
		}
	}
	return output
}
func (ss *SQLStorage) Get_market_stats(market_aid int64) p.RW_API_MarketStats {

	var output p.RW_API_MarketStats
	var query string
	query = "SELECT " +
				"total_vol/1e+18,"+
				"total_num_trades "+
			"FROM "+
				"rw_mkt_stats " +
			"WHERE contract_aid = $1"

	res := ss.db.QueryRow(query,market_aid)
	err := res.Scan(
		&output.TradingVol,
		&output.NumTrades,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return output
		} else {
			ss.Log_msg(fmt.Sprintf("DB error: %v, q=%v",err,query))
			os.Exit(1)
		}
	}
	return output
}
func (ss *SQLStorage) Get_token_full_history(rwalk_aid,token_id int64,offset,limit int) []p.RW_API_FullHistoryEntry {

	records := make([]p.RW_API_FullHistoryEntry,0,32)

	var query string
	query = "SELECT " +
				"block_num," +
				"EXTRACT(EPOCH FROM time_stamp)::BIGINT as ts,"+
				"time_stamp," +
				"contract_aid,"+
				"contract_addr," +
				//--------Mint
				"token_id,"+
				"owner_aid,"+
				"owner_addr,"+
				"seed,"+
				"seed_num::TEXT,"+
				//--------NewOffer
				"seller_aid,"+
				"seller_addr,"+
				"buyer_aid,"+
				"buyer_addr,"+
				"otype,"+
				"offer_id,"+
				"active, "+
				"price, " +
				//--------Offer Canceled
				"offer_canceled_id,"+
				//--------Item Bought
				"item_bought_id," +
				//--------Token Name
				"token_name, "+
				//--------Transfer
				"transfer_id " +
			"FROM (" +
				"(" +
					"SELECT " +
						"t.block_num," +
						"EXTRACT(EPOCH FROM time_stamp)::BIGINT as ts,"+
						"time_stamp," +
						"contract_aid,"+
						//---------Mint
						"ca.addr contract_addr," +
						"token_id,"+
						"owner_aid,"+
						"oa.addr owner_addr,"+
						"seed,"+
						"seed_num,"+
						//--------NewOffer
						"NULL AS seller_aid,"+
						"NULL AS seller_addr,"+
						"NULL AS buyer_aid,"+
						"NULL AS buyer_addr,"+
						"NULL as otype,"+
						"CAST(NULL AS BIGINT) AS offer_id,"+
						"NULL AS active,"+
						"price/1e+18 AS price, " +
						//---------Offer Canceled
						"CAST(NULL AS BIGINT) AS offer_canceled_id,"+
						//---------Item Bought
						"CAST(NULL AS BIGINT) AS item_bought_id,"+
						//---------Token Name
						"CAST(NULL AS TEXT ) as token_name,"+
						//---------TransferId
						"CAST(NULL AS BIGINT) as transfer_id "+
					"FROM rw_mint_evt t " +
						"LEFT JOIN address ca ON contract_aid=ca.address_id "+
						"LEFT JOIN address oa ON owner_aid=oa.address_id "+
					"WHERE (token_id=$1) AND (contract_aid=$2) "+
					"ORDER BY id"+
				") " +
				"UNION ALL" +
				"(" +
					"SELECT "+
						"t.block_num," +
						"EXTRACT(EPOCH FROM time_stamp)::BIGINT as ts,"+
						"time_stamp," +
						"contract_aid,"+
						"ca.addr contract_addr," +
						//---------Mint
						"token_id,"+
						"NULL AS owner_aid,"+
						"NULL AS owner_addr,"+
						"NULL AS seed,"+
						"NULL AS seed_num,"+
						//---------NewOffer
						"seller_aid," +
						"sa.addr seller_addr,"+
						"buyer_aid,"+
						"ba.addr buyer_addr,"+
						"otype," +
						"t.offer_id,"+
						"active,"+
						"price/1e+18 price, "+
						//---------Offer Canceled
						"CAST(NULL AS BIGINT) AS offer_canceled_id,"+
						//---------Item Bought
						"CAST(NULL AS BIGINT) AS item_bought_id,"+
						//---------Token Name
						"CAST(NULL AS TEXT) as token_name,"+
						//---------Transfer
						"CAST(NULL AS BIGINT) as transfer_id " +
					"FROM rw_new_offer t " +
						"LEFT JOIN address ca ON contract_aid=ca.address_id "+
						"LEFT JOIN address sa ON seller_aid=sa.address_id "+
						"LEFT JOIN address ba ON buyer_aid=ba.address_id "+
					"WHERE (token_id=$1) AND (rwalk_aid=$2) " +
					"ORDER BY id" +
				")"+
				"UNION ALL" +
				"(" +
					"SELECT "+
						"c.block_num," +
						"EXTRACT(EPOCH FROM c.time_stamp)::BIGINT as ts,"+
						"c.time_stamp," +
						"c.contract_aid,"+
						"ca.addr contract_addr," +
						//---------Mint
						"token_id,"+
						"NULL AS owner_aid,"+
						"NULL AS owner_addr,"+
						"NULL AS seed,"+
						"NULL AS seed_num,"+
						//---------NewOffer
						"seller_aid," +
						"sa.addr seller_addr,"+
						"buyer_aid,"+
						"ba.addr buyer_addr,"+
						"o.otype as otype,"+
						"o.offer_id,"+
						"NULL AS active,"+
						"o.price/1e+18 AS price,"+
						//---------Offer Canceled
						"c.id offer_canceled_id,"+
						//---------Item Bought
						"CAST(NULL AS BIGINT) AS item_bought_id,"+
						//---------Token Name
						"CAST(NULL AS TEXT) as token_name,"+
						//---------Transfer
						"CAST(NULL AS BIGINT) as transfer_id "+
					"FROM rw_offer_canceled c "+
						"JOIN rw_new_offer o ON (c.offer_id=o.offer_id AND c.contract_aid=o.contract_aid) "+
						"LEFT JOIN address ca ON c.contract_aid=ca.address_id " +
						"LEFT JOIN address sa ON o.seller_aid=sa.address_id "+
						"LEFT JOIN address ba ON o.buyer_aid=ba.address_id " +
					"WHERE (o.token_id=$1) AND (o.rwalk_aid=$2) "+
					"ORDER BY c.id " +
				")"+
				"UNION ALL" +
				"(" +
					"SELECT "+
						"b.block_num," +
						"EXTRACT(EPOCH FROM b.time_stamp)::BIGINT as ts,"+
						"b.time_stamp," +
						"b.contract_aid,"+
						"ca.addr contract_addr," +
						//---------Mint
						"token_id,"+
						"NULL AS owner_aid,"+
						"NULL AS owner_addr,"+
						"NULL AS seed,"+
						"NULL AS seed_num,"+
						//---------NewOffer
						"b.seller_aid," +
						"sa.addr seller_addr,"+
						"b.buyer_aid,"+
						"ba.addr buyer_addr,"+
						"o.otype as otype,"+
						"o.offer_id,"+
						"NULL AS active,"+
						"o.price/1e+18 AS price,"+
						//---------Offer Canceled
						"CAST(NULL AS BIGINT) offer_canceled_id,"+
						//---------Item Bought
						"b.id AS item_bought_id,"+
						//---------Token Name
						"CAST(NULL AS TEXT) as token_name,"+
						//---------Transfer
						"CAST(NULL AS BIGINT) as transfer_id "+
					"FROM rw_item_bought b "+
						"JOIN rw_new_offer o ON (b.offer_id=o.offer_id AND b.contract_aid=o.contract_aid) "+
						"LEFT JOIN address ca ON b.contract_aid=ca.address_id " +
						"LEFT JOIN address sa ON b.seller_aid=sa.address_id "+
						"LEFT JOIN address ba ON b.buyer_aid=ba.address_id " +
					"WHERE (o.token_id=$1) AND (o.rwalk_aid=$2) "+
					"ORDER BY b.id " +
				")"+
				"UNION ALL" +
				"(" +
					"SELECT "+
						"n.block_num," +
						"EXTRACT(EPOCH FROM n.time_stamp)::BIGINT as ts,"+
						"n.time_stamp," +
						"n.contract_aid,"+
						"ca.addr contract_addr," +
						//---------Mint
						"token_id,"+
						"NULL AS owner_aid,"+
						"NULL AS owner_addr,"+
						"NULL AS seed,"+
						"NULL AS seed_num,"+
						//---------NewOffer
						"CAST(NULL AS BIGINT) AS seller_aid," +
						"NULL AS seller_addr,"+
						"CAST(NULL AS BIGINT) AS buyer_aid,"+
						"NULL AS buyer_addr,"+
						"NULL AS otype,"+
						"NULL AS offer_id,"+
						"NULL AS active,"+
						"NULL AS price,"+
						//---------Offer Canceled
						"CAST(NULL AS BIGINT) offer_canceled_id,"+
						//---------Item Bought
						"NULL AS item_bought_id,"+
						//---------Token Name
						"n.new_name token_name,"+
						//---------TransferID
						"CAST(NULL AS BIGINT) transfer_id " +
					"FROM rw_token_name n "+
						"LEFT JOIN address ca ON n.contract_aid=ca.address_id " +
					"WHERE (n.token_id=$1) AND (n.contract_aid=$2) "+
					"ORDER BY n.id " +
				")"+
				"UNION ALL" +
				"(" +
					"SELECT "+
						"tr.block_num," +
						"EXTRACT(EPOCH FROM tr.time_stamp)::BIGINT as ts,"+
						"tr.time_stamp," +
						"tr.contract_aid,"+
						"ca.addr contract_addr," +
						//---------Mint
						"tr.token_id,"+
						"NULL AS owner_aid,"+
						"NULL AS owner_addr,"+
						"NULL AS seed,"+
						"NULL AS seed_num,"+
						//---------NewOffer
						"tr.from_aid AS seller_aid," +
						"fa.addr AS seller_addr,"+
						"tr.to_aid AS buyer_aid,"+
						"ta.addr AS buyer_addr,"+
						"NULL AS otype,"+
						"NULL AS offer_id,"+
						"NULL AS active,"+
						"NULL AS price,"+
						//---------Offer Canceled
						"CAST(NULL AS BIGINT) offer_canceled_id,"+
						//---------Item Bought
						"NULL AS item_bought_id,"+
						//---------Token Name
						"NULL as token_name,"+
						//---------Transfer
						"tr.id AS transfer_id "+
					"FROM rw_transfer tr "+
						"LEFT JOIN address ca ON tr.contract_aid=ca.address_id " +
						"LEFT JOIN address fa ON tr.from_aid=fa.address_id " +
						"LEFT JOIN address ta ON tr.to_aid=ta.address_id " +
						"LEFT JOIN rw_new_offer off ON tr.tx_id=off.tx_id "+
						"LEFT JOIN rw_mint_evt mint ON tr.tx_id=mint.tx_id " +
						"LEFT JOIN rw_item_bought item ON tr.tx_id=item.tx_id " +
						"LEFT JOIN rw_offer_canceled cancel ON tr.tx_id=cancel.tx_id "+
						"LEFT JOIN rw_token_name name ON tr.tx_id=name.tx_id "+
						"LEFT JOIN evt_log elog ON ((tr.tx_id=elog.tx_id) AND ("+
							"elog.topic0_sig='55076e90' OR "+	// new offer
							"elog.topic0_sig='caacc56f' OR "+	// item bought
							"elog.topic0_sig='0ff09947' OR "+	// offer canceled
							"elog.topic0_sig='8ad5e159' OR "+	// token name
							"elog.topic0_sig='ad2bc79f' "+	// mint event
						"))"+
/*					"FROM rw_transfer tr "+
						"LEFT JOIN evt_log l ON (tr.tx_id=l.tx_id) AND "+
							"(l.topic0_sig=
*/
					"WHERE (tr.token_id=$1) AND (tr.contract_aid=$2) "+
						"AND (off.id IS NULL) "+
						"AND (mint.id IS NULL) " +
						"AND (item.id IS NULL) " +
						"AND (cancel.id IS NULL) "+
						"AND (name.id IS NULL) "+
						"AND (elog.id IS NULL) " +
					"ORDER BY tr.id " +
				")"+
			") AS data " +		// FROM
		"ORDER BY ts " +
		"OFFSET $3 LIMIT $4"

	ss.Info.Printf("rwalk_aid=%v token_id=%v, query = %v\n",rwalk_aid,token_id,query)
	rows,err := ss.db.Query(query,token_id,rwalk_aid,offset,limit)
	if (err!=nil) {
		ss.Log_msg(fmt.Sprintf("DB error: %v (query=%v)",err,query))
		os.Exit(1)
	}

	defer rows.Close()
	for rows.Next() {
		var (
			block_num			sql.NullInt64
			timestamp			sql.NullInt64
			datetime			sql.NullString
			contract_aid		sql.NullInt64
			contract_addr		sql.NullString
			token_id			sql.NullInt64
			owner_aid			sql.NullInt64
			owner_addr			sql.NullString
			seed				sql.NullString
			seed_num			sql.NullString
			price				sql.NullFloat64
			seller_aid			sql.NullInt64
			seller_addr			sql.NullString
			buyer_aid			sql.NullInt64
			buyer_addr			sql.NullString
			otype				sql.NullInt64
			offer_id			sql.NullInt64
			active				sql.NullBool
			offer_canceled_id	sql.NullInt64
			item_bought_id		sql.NullInt64
			token_name			sql.NullString
			transfer_id			sql.NullInt64
		)
		err=rows.Scan(
			&block_num,
			&timestamp,
			&datetime,
			&contract_aid,
			&contract_addr,
			&token_id,
			&owner_aid,
			&owner_addr,
			&seed,
			&seed_num,
			&seller_aid,
			&seller_addr,
			&buyer_aid,
			&buyer_addr,
			&otype,
			&offer_id,
			&active,
			&price,
			&offer_canceled_id,
			&item_bought_id,
			&token_name,
			&transfer_id,
		)
		if err!=nil {
			ss.Log_msg(fmt.Sprintf("DB error: %v, q=%v",err,query))
			os.Exit(1)
		}
		var rec p.RW_API_FullHistoryEntry
		if seed.Valid {
			iface := p.RW_API_HistEntry_Mint{
				BlockNum:			block_num.Int64,
				TimeStamp:			timestamp.Int64,
				DateTime:			datetime.String,
				ContractAid:		contract_aid.Int64,
				ContractAddr:		contract_addr.String,
				OwnerAid:			owner_aid.Int64,
				OwnerAddr:			owner_addr.String,
				TokenId:			token_id.Int64,
				SeedHex:			seed.String,
				SeedNum:			seed_num.String,
				Price:				price.Float64,
			}
			rec.Record = iface
			rec.RecordType = 1
		}
		if otype.Valid {
			iface := p.RW_API_HistEntry_Offer{
				BlockNum:			block_num.Int64,
				TimeStamp:			timestamp.Int64,
				DateTime:			datetime.String,
				ContractAid:		contract_aid.Int64,
				ContractAddr:		contract_addr.String,
				TokenId:			token_id.Int64,
				SellerAid:			seller_aid.Int64,
				SellerAddr:			seller_addr.String,
				BuyerAid:			buyer_aid.Int64,
				BuyerAddr:			buyer_addr.String,
				OfferType:			int(otype.Int64),
				OfferId:			offer_id.Int64,
				Active:				active.Bool,
				Price:				price.Float64,
			}
			rec.Record = iface
			rec.RecordType = 2
		}
		if offer_canceled_id.Valid {
			iface := p.RW_API_HistEntry_OfferCanceled{
				BlockNum:			block_num.Int64,
				TimeStamp:			timestamp.Int64,
				DateTime:			datetime.String,
				ContractAid:		contract_aid.Int64,
				ContractAddr:		contract_addr.String,
				TokenId:			token_id.Int64,
				OfferCanceledId:	offer_canceled_id.Int64,
				OfferType:			int(otype.Int64),
				OfferId:			offer_id.Int64,
				SellerAid:			seller_aid.Int64,
				SellerAddr:			seller_addr.String,
				BuyerAid:			buyer_aid.Int64,
				BuyerAddr:			buyer_addr.String,
				Price:				price.Float64,
			};
			if iface.OfferType == 0 { // BUY
				iface.Aid = seller_aid.Int64
				iface.Address = seller_addr.String
			} else {
				iface.Aid = buyer_aid.Int64
				iface.Address = buyer_addr.String
			}
			rec.Record = iface
			rec.RecordType = 3
		}
		if item_bought_id.Valid {
			iface := p.RW_API_HistEntry_ItemBought{
				BlockNum:			block_num.Int64,
				TimeStamp:			timestamp.Int64,
				DateTime:			datetime.String,
				ContractAid:		contract_aid.Int64,
				ContractAddr:		contract_addr.String,
				TokenId:			token_id.Int64,
				ItemBoughtId:		item_bought_id.Int64,
				OfferType:			int(otype.Int64),
				OfferId:			offer_id.Int64,
				SellerAid:			seller_aid.Int64,
				SellerAddr:			seller_addr.String,
				BuyerAid:			buyer_aid.Int64,
				BuyerAddr:			buyer_addr.String,
				Price:				price.Float64,
			};
			if iface.OfferType == 0 { // BUY
				iface.Aid = seller_aid.Int64
				iface.Address = seller_addr.String
			} else {
				iface.Aid = buyer_aid.Int64
				iface.Address = buyer_addr.String
			}
			rec.Record = iface
			rec.RecordType = 4
		}
		if token_name.Valid {
			iface := p.RW_API_HistEntry_TokenName {
				BlockNum:			block_num.Int64,
				TimeStamp:			timestamp.Int64,
				DateTime:			datetime.String,
				ContractAid:		contract_aid.Int64,
				ContractAddr:		contract_addr.String,
				TokenId:			token_id.Int64,
				TokenName:			token_name.String,
			}
			rec.Record = iface
			rec.RecordType = 5
		}
		if transfer_id.Valid {
			iface := p.RW_API_HistEntry_Transfer {
				BlockNum:			block_num.Int64,
				TimeStamp:			timestamp.Int64,
				DateTime:			datetime.String,
				ContractAid:		contract_aid.Int64,
				ContractAddr:		contract_addr.String,
				TokenId:			token_id.Int64,
				TransferId:			transfer_id.Int64,
				FromAid:			seller_aid.Int64,
				FromAddr:			seller_addr.String,
				ToAid:				buyer_aid.Int64,
				ToAddr:				buyer_addr.String,
			}
			rec.Record = iface
			rec.RecordType = 6
		}
		records = append(records,rec)
	}

	return records
}
func (ss *SQLStorage) Get_market_trading_volume_by_period(contract_aid int64,init_ts int,fin_ts int,interval int) []p.RW_API_RandomWalkVolumeHistory {

	var query string
	query = "SELECT sum(price)/1e+18 AS accum_vol FROM rw_item_bought b " +
				"JOIN rw_new_offer o ON o.offer_id=b.offer_id " +
				"WHERE (b.time_stamp < TO_TIMESTAMP($1)i) AND (o.contract_aid=$2)"
	var initial_volume sql.NullFloat64
	err := ss.db.QueryRow(query,init_ts,contract_aid).Scan(&initial_volume)
	if (err!=nil) {
		if err != sql.ErrNoRows {
			ss.Log_msg(fmt.Sprintf("Error in Get_randomwalk_trading_volume_by_period(): %v, q=%v",err,query))
			os.Exit(1)
		}
	}
	query = "WITH periods AS (" +
				"SELECT * FROM (" +
					"SELECT " +
						"generate_series AS start_ts,"+
						"TO_TIMESTAMP(EXTRACT(EPOCH FROM generate_series) + $3) AS end_ts "+
					"FROM (" +
						"SELECT * " +
							"FROM generate_series(" +
								"TO_TIMESTAMP($1)," +
								"TO_TIMESTAMP($2)," +
								"TO_TIMESTAMP($3)-TO_TIMESTAMP(0)) " +
					") AS i" +
				") AS data " +
			") " +
			"SELECT " +
				"COALESCE(COUNT(b.id),0) as num_rows, " +
				"ROUND(FLOOR(EXTRACT(EPOCH FROM start_ts)))::BIGINT as start_ts," +
				"SUM(b.price)/1e+18 as volume "+
			"FROM periods AS p " +
				"LEFT JOIN LATERAL ( "+
					"SELECT b.id,b.time_stamp,o.price "+
						"FROM rw_item_bought b "+
						"JOIN rw_new_offer o ON b.offer_id=o.offer_id "+
						"WHERE b.contract_aid=$4" +
				") b ON " +
					"(p.start_ts <= b.time_stamp) AND "+
					"(b.time_stamp < p.end_ts) " +
			"GROUP BY start_ts " +
			"ORDER BY start_ts"

	ss.Info.Printf("contract_aid=%v init_ts= %v , fin_ts= %v , interval = %v\n",contract_aid,init_ts,fin_ts,interval)
	ss.Info.Printf("query = %v\n",query)
	rows,err := ss.db.Query(query,init_ts,fin_ts,interval,contract_aid)
	if (err!=nil) {
		ss.Log_msg(fmt.Sprintf("DB error: %v (query=%v)",err,query))
		os.Exit(1)
	}
	records := make([]p.RW_API_RandomWalkVolumeHistory,0,8)
	var accum_vol float64 = 0.0
	if initial_volume.Valid {
		accum_vol = initial_volume.Float64
	}
	defer rows.Close()
	for rows.Next() {
		var rec p.RW_API_RandomWalkVolumeHistory
		var sum_volume sql.NullFloat64
		var num_rows int
		err=rows.Scan(
			&num_rows,
			&rec.StartTs,
			&sum_volume,
		)
		if err!=nil {
			ss.Log_msg(fmt.Sprintf("DB error: %v",err))
			os.Exit(1)
		}
		if sum_volume.Valid {
			rec.Volume= sum_volume.Float64
			rec.NumOperations = int64(num_rows)
			accum_vol = accum_vol + rec.Volume
		}
		rec.VolumeAccum = accum_vol
//		fmt.Printf("rec.Vol = %v Accum=%v\n",rec.Volume,rec.VolumeAccum)

		records = append(records,rec)
	}
	return records
}
func (ss *SQLStorage) Get_name_changes_for_token(token_id int64) []p.RW_API_TokenName{

	records := make([]p.RW_API_TokenName,0,32)
	var query string
	query = "SELECT "+
				"t.block_num,"+
				"EXTRACT(EPOCH FROM t.time_stamp)::BIGINT as ts,"+
				"t.time_stamp," +
				"t.contract_aid,"+
				"ca.addr,"+
				"t.new_name,"+
				"tx.tx_hash,"+
				"oa.address_id,"+
				"oa.addr "+
			"FROM rw_token_name t "+
				"LEFT JOIN address ca ON t.contract_aid=ca.address_id "+
				"LEFT JOIN transaction tx ON t.tx_id=tx.id "+
				"LEFT JOIN address oa ON tx.from_aid=oa.address_id "+
			"WHERE token_id = $1 " +
			"ORDER by t.id DESC "

	rows,err := ss.db.Query(query,token_id)
	if (err!=nil) {
		ss.Log_msg(fmt.Sprintf("DB error: %v (query=%v)",err,query))
		os.Exit(1)
	}

	defer rows.Close()
	for rows.Next() {
		var rec p.RW_API_TokenName
		err=rows.Scan(
			&rec.BlockNum,
			&rec.TimeStamp,
			&rec.DateTime,
			&rec.ContractAid,
			&rec.ContractAddr,
			&rec.TokenName,
			&rec.TxHash,
			&rec.OwnerAid,
			&rec.OwnerAddr,
		)
		if err!=nil {
			ss.Log_msg(fmt.Sprintf("DB error: %v, q=%v",err,query))
			os.Exit(1)
		}
		rec.TokenId= token_id
		records = append(records,rec)
	}

	return records

}
func (ss *SQLStorage) Get_random_walk_tokens_by_user(user_aid int64) []p.RW_API_UserToken {

	records := make([]p.RW_API_UserToken,0,32)
	var query string
	query = "SELECT "+
				"t.token_id,"+
				"seed_hex,"+
				"seed_num,"+
				"last_price/1e+18 "+
			"FROM rw_token t "+
			"WHERE cur_owner_aid=$1 "+
			"ORDER BY token_id"
	rows,err := ss.db.Query(query,user_aid)
	if (err!=nil) {
		ss.Log_msg(fmt.Sprintf("DB error: %v (query=%v)",err,query))
		os.Exit(1)
	}

	defer rows.Close()
	for rows.Next() {
		var rec p.RW_API_UserToken
		err=rows.Scan(
			&rec.TokenId,
			&rec.Seed,
			&rec.SeedNum,
			&rec.Price,
		)
		if err!=nil {
			ss.Log_msg(fmt.Sprintf("DB error: %v, q=%v",err,query))
			os.Exit(1)
		}
		records = append(records,rec)
	}

	return records
}
func (ss *SQLStorage) Get_floor_price(rwalk_aid int64,market_aid int64) (float64,error) {

	var output sql.NullFloat64
	var where_cond string =""
	where_cond = fmt.Sprintf(" AND (o.rwalk_aid=%v) ",rwalk_aid)
	where_cond += fmt.Sprintf(" AND (o.contract_aid=%v) ",market_aid)
	var query string
	query = "SELECT " +
				"o.price/1e+18 price "+
			"FROM "+
				"rw_new_offer o "+
			"WHERE (active = 't') AND (otype=1) " + where_cond +
			"ORDER BY o.price ASC "+
			"LIMIT 1"

	res := ss.db.QueryRow(query)
	err := res.Scan(&output)
	if err != nil {
		if err == sql.ErrNoRows {
			return output.Float64,err
		} else {
			ss.Log_msg(fmt.Sprintf("DB error: %v, q=%v",err,query))
			os.Exit(1)
		}
	}
	return output.Float64,nil
}
func (ss *SQLStorage) Get_trading_history_by_user(user_aid int64) []p.RW_API_Offer {

	records := make([]p.RW_API_Offer,0,16)

	var query string
	query = "SELECT " +
				"o.id,"+
				"o.evtlog_id,"+
				"o.block_num,"+
				"o.tx_id, "+
				"tx.tx_hash,"+
				"EXTRACT(EPOCH FROM o.time_stamp)::BIGINT as ts,"+
				"o.time_stamp," +
				"o.offer_id,"+
				"o.otype,"+
				"o.seller_aid,"+
				"sa.addr seller_addr,"+
				"o.buyer_aid,"+
				"ba.addr buyer_addr,"+
				"o.token_id,"+
				"o.active,"+
				"can.id,"+
				"o.price/1e+18 price,"+
				"o.profit/1e+18 profit,"+
				"o.contract_aid,"+
				"ca.addr, " +
				"o.rwalk_aid,"+
				"rwa.addr "+
			"FROM "+
				"rw_new_offer o "+
				"JOIN transaction tx ON o.tx_id=tx.id "+
				"JOIN address sa ON o.seller_aid=sa.address_id "+
				"JOIN address ba ON o.buyer_aid=ba.address_id "+
				"JOIN address ca ON o.contract_aid=ca.address_id "+
				"JOIN address rwa ON o.rwalk_aid=rwa.address_id "+
				"LEFT JOIN rw_offer_canceled can ON (can.contract_aid=o.contract_aid) AND (can.offer_id=o.offer_id) "+
			"WHERE (active = 'f') AND ((o.buyer_aid=$1) OR (o.seller_aid=$1)) " +
			"ORDER BY o.id "

	ss.Info.Printf("user_aid=%v q=%v\n",user_aid,query)
	rows,err := ss.db.Query(query,user_aid)
	if (err!=nil) {
		ss.Log_msg(fmt.Sprintf("DB error: %v (query=%v)",err,query))
		os.Exit(1)
	}

	defer rows.Close()
	for rows.Next() {
		var rec p.RW_API_Offer
		var null_profit sql.NullFloat64
		var null_can_id sql.NullInt64
		err=rows.Scan(
			&rec.Id,
			&rec.EvtLogId,
			&rec.BlockNum,
			&rec.TxId,
			&rec.TxHash,
			&rec.TimeStamp,
			&rec.DateTime,
			&rec.OfferId,
			&rec.OfferType,
			&rec.SellerAid,
			&rec.SellerAddr,
			&rec.BuyerAid,
			&rec.BuyerAddr,
			&rec.TokenId,
			&rec.Active,
			&null_can_id,
			&rec.Price,
			&null_profit,
			&rec.ContractAid,
			&rec.ContractAddr,
			&rec.RWalkAid,
			&rec.RWalkAddr,
		)
		if err!=nil {
			ss.Log_msg(fmt.Sprintf("DB error: %v, q=%v",err,query))
			os.Exit(1)
		}
		if null_profit.Valid {
			rec.Profit = null_profit.Float64
		} else {
			rec.Profit = math.NaN()
		}
		if null_can_id.Valid {
			rec.WasCanceled = true
		}
		records = append(records,rec)
	}

	return records
}
func (ss *SQLStorage) Get_rwalk_user_info(user_aid int64,rwalk_aid int64) (p.RW_API_UserInfo,error) {

	var output p.RW_API_UserInfo
	var query string
	query = "SELECT contract_aid FROM rw_new_offer WHERE contract_aid=$1 LIMIT 1"
	var null_aid sql.NullInt64
	res := ss.db.QueryRow(query,user_aid)
	err := res.Scan(&null_aid)
	if err == nil {
		output.IsMarketPlaceContract = true
	}
	query = "SELECT " +
				"us.total_vol/1e+18, "+
				"us.total_num_trades,"+
				"us.total_num_toks,"+
				"us.total_withdrawals "+
			"FROM "+
				"rw_user_stats us "+
			"WHERE user_aid=$1 AND rwalk_aid=$2"

	output.UserAid=user_aid
	res = ss.db.QueryRow(query,user_aid,rwalk_aid)
	err = res.Scan(
			&output.TotalVolume,
			&output.TotalNumTrades,
			&output.TotalMintedToks,
			&output.TotalNumWithdrawals,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return output,err
		} else {
			ss.Log_msg(fmt.Sprintf("DB error: %v, q=%v",err,query))
			os.Exit(1)
		}
	}
	return output,nil
}
func (ss *SQLStorage) Get_top5_traded_tokens() []p.RW_API_Top5Toks {

	records := make([]p.RW_API_Top5Toks,0,16)

	var query string
	query = "SELECT " +
				"token_id, "+
				"num_trades,"+
				"seed_hex "+
			"FROM "+
				"rw_token t "+
			"ORDER BY num_trades DESC " +
			"LIMIT 5"

	rows,err := ss.db.Query(query)
	if (err!=nil) {
		ss.Log_msg(fmt.Sprintf("DB error: %v (query=%v)",err,query))
		os.Exit(1)
	}

	defer rows.Close()
	for rows.Next() {
		var rec p.RW_API_Top5Toks
		err=rows.Scan(
			&rec.TokenId,
			&rec.TotalTrades,
			&rec.SeedHex,
		)
		if err!=nil {
			ss.Log_msg(fmt.Sprintf("DB error: %v, q=%v",err,query))
			os.Exit(1)
		}
		records = append(records,rec)
	}

	return records
}
func (ss *SQLStorage) Get_rwalk_token_info(rwalk_aid int64,token_id int64) (p.RW_API_TokenInfo,error) {

	var query string
	query = "SELECT " +
				"t.cur_owner_aid,"+
				"oa.addr,"+
				"seed_hex,"+
				"seed_num,"+
				"last_name,"+
				"last_price/1e+18,"+
				"t.total_vol/1e+18, "+
				"t.num_trades "+
			"FROM "+
				"rw_token t "+
				"LEFT JOIN address oa ON oa.address_id=t.cur_owner_aid "+
			"WHERE t.rwalk_aid=$1 AND token_id=$2"

	var output p.RW_API_TokenInfo
	output.TokenId=token_id
	res := ss.db.QueryRow(query,rwalk_aid,token_id)
	err := res.Scan(
			&output.CurOwnerAid,
			&output.CurOwnerAddr,
			&output.SeedHex,
			&output.SeedNum,
			&output.CurName,
			&output.LastPrice,
			&output.TotalVolume,
			&output.NumTrades,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return output,err
		} else {
			ss.Log_msg(fmt.Sprintf("DB error: %v, q=%v",err,query))
			os.Exit(1)
		}
	}
	query = "SELECT "+
				"EXTRACT(EPOCH FROM time_stamp)::BIGINT as ts,"+
				"time_stamp " +
			"FROM rw_token_name "+
			"WHERE contract_aid=$1 AND token_id=$2 " +
			"ORDER BY id DESC "+
			"LIMIT 1"
	res = ss.db.QueryRow(query,rwalk_aid,token_id)
	err = res.Scan(&output.LastNameUpdateTs,&output.LastNameUpdateDate)
	if err != nil {
		if err == sql.ErrNoRows {
		} else {
			ss.Log_msg(fmt.Sprintf("DB error: %v, q=%v",err,query))
			os.Exit(1)
		}
	}
	return output,nil
}
func (ss *SQLStorage) Get_rwalk_mint_intervals(rwalk_aid int64) []p.RW_API_MintInterval {
	// gets mint number and time elapsed between mints (for scatter plot)

	records := make([]p.RW_API_MintInterval,0,256)

	var query string
	query = "SELECT "+
				"EXTRACT(EPOCH FROM m.time_stamp)::BIGINT as ts,"+
				"token_id " +
			"FROM rw_mint_evt m " +
				"WHERE contract_aid = $1 "+
			"ORDER BY token_id"

	rows,err := ss.db.Query(query,rwalk_aid)
	if (err!=nil) {
		ss.Log_msg(fmt.Sprintf("DB error: %v (query=%v)",err,query))
		os.Exit(1)
	}

	defer rows.Close()
	for rows.Next() {
		var rec p.RW_API_MintInterval
		err=rows.Scan(
			&rec.TimeStamp,
			&rec.TokenId,
		)
		if err!=nil {
			ss.Log_msg(fmt.Sprintf("DB error: %v, q=%v",err,query))
			os.Exit(1)
		}
		rec.MintNumber = rec.TokenId // because tokenID is sequential also
		records = append(records,rec)
	}

	var prev_ts int64 = 0
	for i:=0; i<len(records); i++ {
		rec := &records[i]
		if prev_ts > 0 {
			rec.Interval = rec.TimeStamp - prev_ts
		}
		prev_ts = rec.TimeStamp
	}
	return records
}
func (ss *SQLStorage) Get_rwalk_withdrawal_chart(rwalk_aid int64) []p.RW_API_WithdrawalChartEntry {

	records := make([]p.RW_API_WithdrawalChartEntry,0,256)
	var query string
	query = "SELECT "+
				"EXTRACT(EPOCH FROM m.time_stamp)::BIGINT as ts,"+
				"time_stamp," +
				"price/2e+18 price "+	// we use 2e+18 because we divide by 2
			"FROM rw_mint_evt m "+
			"WHERE contract_aid=$1"


	rows,err := ss.db.Query(query,rwalk_aid)
	if (err!=nil) {
		ss.Log_msg(fmt.Sprintf("DB error: %v (query=%v)",err,query))
		os.Exit(1)
	}

	var withdrawal_amount float64
	defer rows.Close()
	for rows.Next() {
		var rec p.RW_API_WithdrawalChartEntry
		err=rows.Scan(
			&rec.TimeStamp,
			&rec.DateTime,
			&rec.WithdrawalAmount,
		)
		if err!=nil {
			ss.Log_msg(fmt.Sprintf("DB error: %v, q=%v",err,query))
			os.Exit(1)
		}
		withdrawal_amount += rec.WithdrawalAmount
		rec.WithdrawalAmount = withdrawal_amount
		records = append(records,rec)
	}

	return records
}
func (ss *SQLStorage) Get_sale_history(contract_aid int64,offset,limit int) []p.RW_API_Offer {

	records := make([]p.RW_API_Offer,0,16)

	var query string
	query = "SELECT " +
				"o.id,"+
				"o.evtlog_id,"+
				"o.block_num,"+
				"o.tx_id, "+
				"tx.tx_hash,"+
				"EXTRACT(EPOCH FROM o.time_stamp)::BIGINT as ts,"+
				"o.time_stamp," +
				"o.offer_id,"+
				"o.otype,"+
				"o.seller_aid,"+
				"sa.addr seller_addr,"+
				"o.buyer_aid,"+
				"ba.addr buyer_addr,"+
				"o.token_id,"+
				"o.active,"+
				"can.id,"+
				"o.price/1e+18 price, "+
				"o.profit/1e+18 profit, "+
				"o.contract_aid," +
				"ca.addr," +
				"o.rwalk_aid,"+
				"rwa.addr "+
			"FROM "+
				"rw_new_offer o "+
				"JOIN transaction tx ON o.tx_id=tx.id "+
				"JOIN address sa ON o.seller_aid=sa.address_id "+
				"JOIN address ba ON o.buyer_aid=ba.address_id "+
				"JOIN address ca ON o.contract_aid=ca.address_id "+
				"JOIN address rwa ON o.rwalk_aid=rwa.address_id "+
				"LEFT JOIN rw_offer_canceled can ON (can.contract_aid=o.contract_aid) AND (can.offer_id=o.offer_id) "+
			"WHERE (active = 'f') AND (o.contract_aid=$3) AND (can.id IS NULL) "+
			"ORDER BY o.id " +
			"OFFSET $1 LIMIT $2"

	rows,err := ss.db.Query(query,offset,limit,contract_aid)
	if (err!=nil) {
		ss.Log_msg(fmt.Sprintf("DB error: %v (query=%v)",err,query))
		os.Exit(1)
	}

	defer rows.Close()
	for rows.Next() {
		var rec p.RW_API_Offer
		var null_profit sql.NullFloat64
		var null_can_id sql.NullInt64
		err=rows.Scan(
			&rec.Id,
			&rec.EvtLogId,
			&rec.BlockNum,
			&rec.TxId,
			&rec.TxHash,
			&rec.TimeStamp,
			&rec.DateTime,
			&rec.OfferId,
			&rec.OfferType,
			&rec.SellerAid,
			&rec.SellerAddr,
			&rec.BuyerAid,
			&rec.BuyerAddr,
			&rec.TokenId,
			&rec.Active,
			&null_can_id,
			&rec.Price,
			&null_profit,
			&rec.ContractAid,
			&rec.ContractAddr,
			&rec.RWalkAid,
			&rec.RWalkAddr,
		)
		if err!=nil {
			ss.Log_msg(fmt.Sprintf("DB error: %v, q=%v",err,query))
			os.Exit(1)
		}
		if null_profit.Valid {
			rec.Profit = null_profit.Float64
		} else {
			rec.Profit = math.NaN()
		}
		if null_can_id.Valid {
			rec.WasCanceled = true
		}
		records = append(records,rec)
	}

	return records
}
func (ss *SQLStorage) Get_rwalk_floor_price_for_periods(rwalk_aid,market_aid int64,init_ts int,fin_ts int,interval int) []p.RW_API_RWalkFloorPrice {

	var query string
	query = "WITH periods AS (" +
				"SELECT * FROM (" +
					"SELECT " +
						"generate_series AS start_ts,"+
						"TO_TIMESTAMP(EXTRACT(EPOCH FROM generate_series) + $3) AS end_ts "+
					"FROM (" +
						"SELECT * " +
							"FROM generate_series(" +
								"TO_TIMESTAMP($1)," +
								"TO_TIMESTAMP($2)," +
								"TO_TIMESTAMP($3)-TO_TIMESTAMP(0)) " +
					") AS i" +
				") AS data " +
			") " +
			"SELECT " +
			//	"COALESCE(MIN(o.id),0) as num_rows, " +
				"ROUND(FLOOR(EXTRACT(EPOCH FROM start_ts)))::BIGINT as start_ts," +
				"MIN(o.price)/1e+18 as floor_price "+
			"FROM periods AS p " +
				"LEFT JOIN LATERAL ( "+
					"SELECT o.id,o.time_stamp,o.price "+
						"FROM rw_new_offer o "+
						"WHERE o.contract_aid=$4 AND o.rwalk_aid=$5 " +
							"AND otype = 1 "+
				") o ON " +
					"(p.start_ts <= o.time_stamp) AND "+
					"(o.time_stamp < p.end_ts) " +
			"GROUP BY start_ts " +
			"ORDER BY start_ts"

	ss.Info.Printf("market_aid=%v rwalk_aid=%v init_ts= %v , fin_ts= %v , interval = %v\n",market_aid,rwalk_aid,init_ts,fin_ts,interval)
	ss.Info.Printf("query = %v\n",query)
	rows,err := ss.db.Query(query,init_ts,fin_ts,interval,market_aid,rwalk_aid)
	if (err!=nil) {
		ss.Log_msg(fmt.Sprintf("DB error: %v (query=%v)",err,query))
		os.Exit(1)
	}
	records := make([]p.RW_API_RWalkFloorPrice,0,8)
	defer rows.Close()
	for rows.Next() {
		var rec p.RW_API_RWalkFloorPrice
		var null_float sql.NullFloat64
		err=rows.Scan(
			&rec.TimeStamp,
			&null_float,
		)
		if err!=nil {
			ss.Log_msg(fmt.Sprintf("DB error: %v",err))
			os.Exit(1)
		}
		if null_float.Valid {
			rec.Price = null_float.Float64
		} else {
			continue
		}
		records = append(records,rec)
	}
	return records
}
