package cosmicgame

import (
	"os"
	"fmt"
	"database/sql"


	p "github.com/PredictionExplorer/augur-explorer/primitives/cosmicgame"
)

// buildBidSelectQuery returns the unified SELECT query for bids
func (sw *SQLStorageWrapper) buildBidSelectQuery(whereClause, orderBy, limitOffset string) string {
	query := "SELECT b.evtlog_id,b.block_num,t.id,t.tx_hash," +
		"EXTRACT(EPOCH FROM b.time_stamp)::BIGINT,b.time_stamp," +
		"b.bidder_aid,ba.addr,b.eth_price,b.eth_price/1e18 AS eth_price_eth, " +
		"b.cst_price, b.cst_price/1e18 AS cst_price_eth, b.rwalk_nft_id," +
		"d.token_id,d.tok_addr, d.token_uri, b.msg, b.round_num, " +
		"b.cst_reward, b.cst_reward/1e18, b.bid_type, " +
		"EXTRACT(EPOCH FROM b.prize_time)::BIGINT AS prize_time_ts, b.prize_time, " +
		"GREATEST(0, EXTRACT(EPOCH FROM b.prize_time)::BIGINT - EXTRACT(EPOCH FROM NOW())::BIGINT) AS time_until_prize, " +
		"b.bid_position, d2.tok_addr, d2.amount, d2.amount/1e18 " +
		"FROM " + sw.S.SchemaName() + ".cg_bid b " +
		"LEFT JOIN " + sw.S.SchemaName() + ".transaction t ON t.id=tx_id " +
		"LEFT JOIN " + sw.S.SchemaName() + ".address ba ON b.bidder_aid=ba.address_id " +
		"LEFT JOIN LATERAL (SELECT d.bid_id,token_id,token_aid,ta.addr tok_addr,d.token_uri " +
		"FROM " + sw.S.SchemaName() + ".cg_nft_donation d " +
		"JOIN " + sw.S.SchemaName() + ".address ta ON d.token_aid=ta.address_id) d ON b.id=d.bid_id " +
		"LEFT JOIN LATERAL (SELECT d.bid_id,token_id,token_aid,ta.addr tok_addr,d.amount " +
		"FROM " + sw.S.SchemaName() + ".cg_erc20_donation d " +
		"JOIN " + sw.S.SchemaName() + ".address ta ON d.token_aid=ta.address_id) d2 ON b.id=d2.bid_id "
	
	if whereClause != "" {
		query += "WHERE " + whereClause + " "
	}
	if orderBy != "" {
		query += "ORDER BY " + orderBy + " "
	}
	if limitOffset != "" {
		query += limitOffset
	}
	return query
}

// scanBidRecord scans a single bid record from a row
// Returns the populated CGBidRec and any error
func scanBidRecord(rows *sql.Rows) (p.CGBidRec, error) {
	var rec p.CGBidRec
	var null_token_id sql.NullInt64
	var null_tok_addr, null_token_uri sql.NullString
	var null_donated_erc20_addr, null_donated_erc20_amount sql.NullString
	var null_donated_erc20_amount_eth sql.NullFloat64
	
	err := rows.Scan(
		&rec.EvtLogId,
		&rec.BlockNum,
		&rec.TxId,
		&rec.TxHash,
		&rec.TimeStamp,
		&rec.DateTime,
		&rec.BidderAid,
		&rec.BidderAddr,
		&rec.EthPrice,
		&rec.EthPriceEth,
		&rec.CstPrice,
		&rec.CstPriceEth,
		&rec.RWalkNFTId,
		&null_token_id,
		&null_tok_addr,
		&null_token_uri,
		&rec.Message,
		&rec.RoundNum,
		&rec.CSTReward,
		&rec.CSTRewardEth,
		&rec.BidType,
		&rec.PrizeTime,
		&rec.PrizeTimeDate,
		&rec.TimeUntilPrize,
		&rec.BidPosition,
		&null_donated_erc20_addr,
		&null_donated_erc20_amount,
		&null_donated_erc20_amount_eth,
	)
	
	if err != nil {
		return rec, err
	}
	
	// Handle null fields
	rec.NFTDonationTokenId = -1
	if null_token_id.Valid {
		rec.NFTDonationTokenId = null_token_id.Int64
	}
	if null_tok_addr.Valid {
		rec.NFTDonationTokenAddr = null_tok_addr.String
	}
	if null_token_uri.Valid {
		rec.NFTTokenURI = null_token_uri.String
	}
	if null_donated_erc20_addr.Valid {
		rec.DonatedERC20TokenAddr = null_donated_erc20_addr.String
	}
	if null_donated_erc20_amount.Valid {
		rec.DonatedERC20TokenAmount = null_donated_erc20_amount.String
	}
	if null_donated_erc20_amount_eth.Valid {
		rec.DonatedERC20TokenAmountEth = null_donated_erc20_amount_eth.Float64
	}
	
	return rec, nil
}

func (sw *SQLStorageWrapper) Get_bid_id_by_evtlog(evtlog_id int64) (int64) {

	var query string
	query = "SELECT b.id FROM cg_bid b WHERE b.evtlog_id=$1"

	row := sw.S.Db().QueryRow(query,evtlog_id)
	var err error
	var output int64
	err=row.Scan(&output)
	if (err!=nil) {
		if err == sql.ErrNoRows {
			return -1
		}
		sw.S.Log_msg(fmt.Sprintf("DB Error: %v, q=%v",err,query))
		os.Exit(1)
	}
	return output
}
func (sw *SQLStorageWrapper) Get_bids(offset,limit int) []p.CGBidRec {

	if limit == 0 { limit = 1000000 }
	query := sw.buildBidSelectQuery("", "b.id DESC", "OFFSET $1 LIMIT $2")
	rows,err := sw.S.Db().Query(query,offset,limit)
	if (err!=nil) {
		sw.S.Log_msg(fmt.Sprintf("DB error: %v (query=%v)",err,query))
		os.Exit(1)
	}
	records := make([]p.CGBidRec, 0, 256)
	defer rows.Close()
	for rows.Next() {
		rec, err := scanBidRecord(rows)
		if err != nil {
			sw.S.Log_msg(fmt.Sprintf("DB error: %v (query=%v)", err, query))
			os.Exit(1)
		}
		records = append(records, rec)
	}
	return records
}
func (sw *SQLStorageWrapper) Get_bid_info(evtlog_id int64) (bool, p.CGBidRec) {

	query := sw.buildBidSelectQuery("b.evtlog_id=$1", "", "")
	rows, err := sw.S.Db().Query(query, evtlog_id)
	if err != nil {
		sw.S.Log_msg(fmt.Sprintf("DB error: %v (query=%v)", err, query))
		os.Exit(1)
	}
	defer rows.Close()
	
	if !rows.Next() {
		return false, p.CGBidRec{} // No rows found
	}
	
	rec, err := scanBidRecord(rows)
	if err != nil {
		sw.S.Log_msg(fmt.Sprintf("DB Error: %v, q=%v", err, query))
		os.Exit(1)
	}
	
	return true, rec
}

func (sw *SQLStorageWrapper) Get_bids_by_round_num(round_num int64,sort,offset,limit int) ([]p.CGBidRec,int64) {

	var order_by string = "ASC"
	if sort == 1 {
		order_by = "DESC"
	}
	query := sw.buildBidSelectQuery("b.round_num=$1", "b.id "+order_by, "OFFSET $2 LIMIT $3")
	rows,err := sw.S.Db().Query(query,round_num,offset,limit)
	if (err!=nil) {
		sw.S.Log_msg(fmt.Sprintf("DB error: %v (query=%v)",err,query))
		os.Exit(1)
	}
	records := make([]p.CGBidRec, 0, 32)
	defer rows.Close()
	for rows.Next() {
		rec, err := scanBidRecord(rows)
		if err != nil {
			sw.S.Log_msg(fmt.Sprintf("DB error: %v (query=%v)", err, query))
			os.Exit(1)
		}
		records = append(records, rec)
	}

	query = "SELECT count(*) AS total_rows FROM cg_bid b WHERE round_num=$1"
	var null_total_rows sql.NullInt64
	row := sw.S.Db().QueryRow(query,round_num)
	err=row.Scan(&null_total_rows);
	if (err!=nil) {
		if err == sql.ErrNoRows {
			return records,-1
		}
		sw.S.Log_msg(fmt.Sprintf("Error in Get_bids_by_round_num(): %v, q=%v",err,query))
		os.Exit(1)
	}
	return records,null_total_rows.Int64

}
func (sw *SQLStorageWrapper) Get_round_start_timestamp(round_num uint64) int64  {

	var query string
	query = "SELECT "+
				"EXTRACT(EPOCH FROM b.time_stamp)::BIGINT "+
			"FROM cg_bid b "+
			"WHERE round_num=$1 "+
			"ORDER BY b.id LIMIT 1"

	var null_ts sql.NullInt64
	row := sw.S.Db().QueryRow(query,round_num)
	var err error
	err=row.Scan(&null_ts);
	if (err!=nil) {
		if err == sql.ErrNoRows {
			return 0
		}
		sw.S.Log_msg(fmt.Sprintf("Error in Get_round_start_timestamp(): %v, q=%v",err,query))
		os.Exit(1)
	}
	return null_ts.Int64
}
func (sw *SQLStorageWrapper) Get_random_walk_tokens_in_bids() []p.CGRWalkUsed {

	var query string
	query = "SELECT "+
				"b.id,"+
				"b.evtlog_id,"+
				"b.block_num,"+
				"tx.id,"+
				"tx.tx_hash,"+
				"EXTRACT(EPOCH FROM b.time_stamp)::BIGINT,"+
				"b.time_stamp,"+
				"b.round_num,"+
				"b.bidder_aid,"+
				"ba.addr,"+
				"b.rwalk_nft_id "+
			"FROM "+sw.S.SchemaName()+".cg_bid b "+
				"LEFT JOIN transaction tx ON tx.id=b.tx_id "+
				"LEFT JOIN address ba ON b.bidder_aid=ba.address_id "+
			"WHERE b.rwalk_nft_id != -1 "+
			"ORDER BY b.id DESC "

	rows,err := sw.S.Db().Query(query)
	if (err!=nil) {
		sw.S.Log_msg(fmt.Sprintf("DB error: %v (query=%v)",err,query))
		os.Exit(1)
	}
	records := make([]p.CGRWalkUsed,0, 16)
	defer rows.Close()
	for rows.Next() {
		var rec p.CGRWalkUsed
		err=rows.Scan(
			&rec.RecordId,
			&rec.EvtLogId,
			&rec.BlockNum,
			&rec.TxId,
			&rec.TxHash,
			&rec.TimeStamp,
			&rec.DateTime,
			&rec.RoundNum,
			&rec.BidderAid,
			&rec.BidderAddr,
			&rec.RWalkTokenId,
		)
		if err != nil {
			sw.S.Log_msg(fmt.Sprintf("DB error: %v (query=%v)",err,query))
			os.Exit(1)
		}
		records = append(records,rec)
	}
	return records
}
func (sw *SQLStorageWrapper) Get_cosmic_game_bid_by_evtlog_id(bid_evtlog_id int64) int64 {

	var query string
	query = "SELECT id FROM "+sw.S.SchemaName()+".cg_bid WHERE evtlog_id=$1"
	res := sw.S.Db().QueryRow(query,bid_evtlog_id)
	var null_id sql.NullInt64
	err := res.Scan(&null_id)
	if (err!=nil) {
		if err == sql.ErrNoRows {
			return 0	// if bid wasn't found there wasn't any bid but pure Donate() instead,
						//	so we return 0 as Id
		}
		sw.S.Log_msg(fmt.Sprintf("DB error: %v q=%v",err,query))
		os.Exit(1)
	}
	return null_id.Int64
}
