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
					"WHERE token_id=$1 "+
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
					"WHERE token_id=$1 " +
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
						"JOIN rw_new_offer o ON c.offer_id=o.offer_id "+
						"LEFT JOIN address ca ON c.contract_aid=ca.address_id " +
						"LEFT JOIN address sa ON o.seller_aid=sa.address_id "+
						"LEFT JOIN address ba ON o.buyer_aid=ba.address_id " +
					"WHERE o.token_id=$1 "+
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
						"JOIN rw_new_offer o ON b.offer_id=o.offer_id "+
						"LEFT JOIN address ca ON b.contract_aid=ca.address_id " +
						"LEFT JOIN address sa ON b.seller_aid=sa.address_id "+
						"LEFT JOIN address ba ON b.buyer_aid=ba.address_id " +
					"WHERE o.token_id=$1 "+
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
					"WHERE n.token_id=$1 "+
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
					"WHERE tr.token_id=$1 "+
						"AND (off.id IS NULL) "+
						"AND (mint.id IS NULL) " +
						"AND (item.id IS NULL) " +
						"AND (cancel.id IS NULL) "+
						"AND (name.id IS NULL) "+
					"ORDER BY tr.id " +
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
