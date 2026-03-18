package cosmicgame

import (
	"fmt"
	"time"

	p "github.com/PredictionExplorer/augur-explorer/rwcg/primitives/cosmicgame"
)

// Get_banned_bids returns all rows from cg_banned_bids ordered by id (same as FastAPI).
func (sw *SQLStorageWrapper) Get_banned_bids() []p.CGBannedBidRec {
	query := "SELECT id, bid_id, user_addr, created_at FROM " + sw.S.SchemaName() + ".cg_banned_bids ORDER BY id"
	rows, err := sw.S.Db().Query(query)
	if err != nil {
		sw.S.Log_msg(fmt.Sprintf("DB error: %v (query=%v)", err, query))
		return nil
	}
	defer rows.Close()
	var list []p.CGBannedBidRec
	for rows.Next() {
		var rec p.CGBannedBidRec
		if err := rows.Scan(&rec.Id, &rec.BidId, &rec.UserAddr, &rec.CreatedAt); err != nil {
			sw.S.Log_msg(fmt.Sprintf("DB error scanning cg_banned_bids: %v", err))
			return list
		}
		list = append(list, rec)
	}
	return list
}

// Insert_banned_bid adds a row to cg_banned_bids. created_at is set to Unix timestamp.
func (sw *SQLStorageWrapper) Insert_banned_bid(bid_id int64, user_addr string) error {
	created_at := time.Now().Unix()
	query := "INSERT INTO " + sw.S.SchemaName() + ".cg_banned_bids (bid_id, user_addr, created_at) VALUES ($1, $2, $3)"
	_, err := sw.S.Db().Exec(query, bid_id, user_addr, created_at)
	return err
}

// Delete_banned_bid_by_bid_id removes all rows with the given bid_id (FastAPI unban_bid deletes by bid_id).
func (sw *SQLStorageWrapper) Delete_banned_bid_by_bid_id(bid_id int64) error {
	query := "DELETE FROM " + sw.S.SchemaName() + ".cg_banned_bids WHERE bid_id = $1"
	_, err := sw.S.Db().Exec(query, bid_id)
	return err
}
