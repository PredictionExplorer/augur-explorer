package dbs
import (
	"os"
	"fmt"
	"database/sql"
	p "github.com/PredictionExplorer/augur-explorer/primitives"
)
func (ss *SQLStorage) Get_active_offers(order_by int) []p.RW_API_Offer {

	records := make([]p.RW_API_Offer,0,16)

	var order_by_mod string ="ORDER BY o.id"
	if order_by == 1 {
		order_by_mod = "ORDER BY o.price DESC"
	}
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
				"o.price/1e+18 price "+
			"FROM "+
				"rw_new_offer o "+
				"JOIN transaction tx ON o.tx_id=tx.id "+
				"JOIN address sa ON o.seller_aid=sa.address_id "+
				"JOIN address ba ON o.buyer_aid=ba.address_id "+
			"WHERE active = 't' " +
			order_by_mod

	rows,err := ss.db.Query(query)
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
		)
		if err!=nil {
			ss.Log_msg(fmt.Sprintf("DB error: %v, q=%v",err,query))
			os.Exit(1)
		}
		records = append(records,rec)
	}

	return records
}
func (ss *SQLStorage) Get_minted_tokens_by_period(ini_ts,fin_ts int) []p.RW_API_Token {

	records := make([]p.RW_API_Token,0,32)
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
			"WHERE (t.time_stamp >= TO_TIMESTAMP($1)) AND (t.time_stamp<TO_TIMESTAMP($2))"
	ss.Info.Printf("ini=%v,fin=%v, q=%v\n",ini_ts,fin_ts,query)
	rows,err := ss.db.Query(query,ini_ts,fin_ts)
	if (err!=nil) {
		ss.Log_msg(fmt.Sprintf("DB error: %v (query=%v)",err,query))
		os.Exit(1)
	}

	defer rows.Close()
	for rows.Next() {
		var rec p.RW_API_Token
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
func (ss *SQLStorage) Get_minted_tokens_sequentially(offset,limit int) []p.RW_API_Token {

	records := make([]p.RW_API_Token,0,32)
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
			"ORDER by t.id DESC "+
			"OFFSET $1 LIMIT $2"

	rows,err := ss.db.Query(query,offset,limit)
	if (err!=nil) {
		ss.Log_msg(fmt.Sprintf("DB error: %v (query=%v)",err,query))
		os.Exit(1)
	}

	defer rows.Close()
	for rows.Next() {
		var rec p.RW_API_Token
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
func (ss *SQLStorage) Get_sale_history(offset,limit int) []p.RW_API_Offer {

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
				"o.price/1e+18 price "+
			"FROM "+
				"rw_new_offer o "+
				"JOIN transaction tx ON o.tx_id=tx.id "+
				"JOIN address sa ON o.seller_aid=sa.address_id "+
				"JOIN address ba ON o.buyer_aid=ba.address_id "+
			"WHERE active = 'f' " +
			"ORDER BY o.id " +
			"OFFSET $1 LIMIT $2"

	rows,err := ss.db.Query(query,offset,limit)
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
		)
		if err!=nil {
			ss.Log_msg(fmt.Sprintf("DB error: %v, q=%v",err,query))
			os.Exit(1)
		}
		records = append(records,rec)
	}

	return records
}
func (ss *SQLStorage) Get_global_stats() p.RW_API_GlobalStats {

	var output p.RW_API_GlobalStats
	var query string
	query = "SELECT " +
				"total_vol/1e+18,"+
				"total_num_trades,"+
				"total_num_toks,"+
				"total_withdrawals "+
			"FROM "+
				"rw_stats "

	res := ss.db.QueryRow(query)
	err := res.Scan(
		&output.TradingVol,
		&output.NumTrades,
		&output.TokensMinted,
		&output.NumWithdrawals,
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
func (ss *SQLStorage) Get_token_full_history(token_id int64,offset,limit int) []p.RW_API_FullHistoryEntry {

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
				"seed,"+
				"seed_num::TEXT,"+
				"price, " +
				//--------NewOffer
				"seller_aid,"+
				"seller_addr,"+
				"buyer_aid,"+
				"buyer_addr,"+
				"otype,"+
				"active "+
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
						"seed,"+
						"seed_num,"+
						//--------NewOffer
						"NULL AS seller_aid,"+
						"NULL AS seller_addr,"+
						"NULL AS buyer_aid,"+
						"NULL AS buyer_addr,"+
						"NULL as otype,"+
						"NULL AS active,"+
						"price/1e+18 AS price " +
					"FROM rw_mint_evt t " +
						"LEFT JOIN address ca ON contract_aid=ca.address_id "+
					"WHERE token_id=$1 "+
					"ORDER BY id"+
				")" +
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
						"NULL AS seed,"+
						"NULL AS seed_num,"+
						//---------NewOffer
						"seller_aid," +
						"sa.addr seller_addr,"+
						"buyer_aid,"+
						"ba.addr buyer_addr,"+
						"otype," +
						"active,"+
						"price/1e+18 price "+
					"FROM rw_new_offer t " +
						"LEFT JOIN address ca ON contract_aid=ca.address_id "+
						"LEFT JOIN address sa ON seller_aid=sa.address_id "+
						"LEFT JOIN address ba ON buyer_aid=ba.address_id "+
					"WHERE token_id=$1 " +
					"ORDER BY id" +
				")"+
			") AS data " +		// FROM
		"ORDER BY ts " +
		"OFFSET $2 LIMIT $3"

	ss.Info.Printf("token_id=%v, query = %v\n",token_id,query)
	rows,err := ss.db.Query(query,token_id,offset,limit)
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
			seed				sql.NullString
			seed_num			sql.NullString
			price				sql.NullFloat64
			seller_aid			sql.NullInt64
			seller_addr			sql.NullString
			buyer_aid			sql.NullInt64
			buyer_addr			sql.NullString
			otype				sql.NullInt64
			active				sql.NullBool
		)
		err=rows.Scan(
			&block_num,
			&timestamp,
			&datetime,
			&contract_aid,
			&contract_addr,
			&token_id,
			&seed,
			&seed_num,
			&price,
			&seller_aid,
			&seller_addr,
			&buyer_aid,
			&buyer_addr,
			&otype,
			&active,
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
				Active:				active.Bool,
				Price:				price.Float64,
			}
			rec.Record = iface
			rec.RecordType = 2
		}
		records = append(records,rec)
	}

	return records
}
