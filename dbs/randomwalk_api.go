package dbs
import (
	"os"
	"fmt"
//	"database/sql"
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
