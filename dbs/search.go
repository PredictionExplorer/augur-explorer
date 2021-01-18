package dbs

import (
	"fmt"
	"os"
	"database/sql"
	_  "github.com/lib/pq"

	//"github.com/ethereum/go-ethereum/common"
//	"github.com/ethereum/go-ethereum/core/types"

	p "github.com/PredictionExplorer/augur-explorer/primitives"
)
func (ss *SQLStorage) Search_keywords_in_markets(keywords string) []p.TextSearchResult {

	var query string
	query = "WITH search_tokens AS (" +
				"SELECT market_aid,cat_id,tok_type "+
				"FROM mkt_words,plainto_tsquery($1) AS q " +
				"WHERE tokens @@ q" +
			") " +
			"SELECT " + 
				"st.market_aid," +
				"st.tok_type, " +
				"st.cat_id," +
				"ma.addr," +
				"m.extra_info::json->>'description'," +
				"m.cur_volume," +
				"c.category, " +
				"c.total_markets " +
			"FROM search_tokens AS st " +
			"LEFT JOIN market AS m ON st.market_aid=m.market_aid " +
			"LEFT JOIN category AS c ON st.cat_id=c.cat_id " +
			"LEFT JOIN address AS ma ON m.market_aid=ma.address_id " +
			"ORDER BY st.tok_type DESC,c.total_markets DESC,cur_volume DESC"

	rows,err := ss.db.Query(query,keywords)
	if (err!=nil) {
		ss.Log_msg(fmt.Sprintf("DB error: %v (query=%v)",err,query))
		os.Exit(1)
	}
	records := make([]p.TextSearchResult,0,8)

	defer rows.Close()
	for rows.Next() {
		var description sql.NullString
		var category sql.NullString
		var mkt_addr sql.NullString
		var market_aid sql.NullInt64
		var cur_vol sql.NullFloat64
		var cat_id,total_markets sql.NullInt64
		var rec p.TextSearchResult
		err=rows.Scan(
			&market_aid,
			&rec.ObjType,
			&cat_id,
			&mkt_addr,
			&description,
			&cur_vol,
			&category,
			&total_markets,
		)
		if err!=nil {
			ss.Log_msg(fmt.Sprintf("DB error: %v, q=%v",err,query))
			os.Exit(1)
		}
		if market_aid.Valid { rec.MktAid = market_aid.Int64 }
		if cat_id.Valid { rec.CatId = cat_id.Int64 }
		if total_markets.Valid { rec.TotalMarkets = total_markets.Int64 }
		if mkt_addr.Valid { rec.MktAddr = mkt_addr.String }
		if cur_vol.Valid { rec.Volume = cur_vol.Float64 }
		if description.Valid { rec.MktDescription = description.String }
		if category.Valid { rec.Category = category.String }
		records = append(records,rec)
	}
	return records
}
