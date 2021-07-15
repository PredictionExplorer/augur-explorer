package dbs
import (
	"os"
	"fmt"
	"database/sql"
	p "github.com/PredictionExplorer/augur-explorer/primitives"
)
func (ss *SQLStorage) Insert_polymarket_from_web_api(
	evt *p.Pol_Market_API_Record_JSON,
	evt_comp *p.Pol_Market_API_Record_Complementary,
	block_num int64,
	tx_id int64,
) { 

	market_maker_aid:=ss.Lookup_or_create_address(evt.MarketMakerAddr,block_num,tx_id)
	var query string
	query = "INSERT INTO pol_market (" +
				"id," +
				"market_id,"+
				"question,"+
				"condition_id,"+
				"slug," +
				"twitter_card_image,"+
				"resolution_source,"+
				"end_date," +
				"end_date_ts," +
				"category," +
				"amm_type," +
				"sponsor_name,"+
				"sponsor_image,"+
				"start_date,"+
				"start_date_ts,"+
				"x_axis_value,"+
				"y_axis_value,"+
				"denomination_token,"+
				"fee,"+
				"image," +
				"icon,"+
				"lower_bound,"+
				"upper_bound,"+
				"description," +
				"tags,"+
				"outcomes,"+
				"outcome_prices,"+
				"volume," +
				"active,"+
				"market_type,"+
				"format_type,"+
				"lower_bound_date,"+
				"lower_bound_ts,"+
				"upper_bound_date,"+
				"upper_bound_ts,"+
				"closed,"+
				"mkt_mkr_aid,"+
				"created_at_date,"+
				"created_at_ts,"+
				"updated_at_date,"+
				"updated_at_ts,"+
				"closed_time,"+
				"closed_time_ts,"+
				"wide_format," +
				"new," +
				"sent_discord,"+
				"mailchimp_tag,"+
				"featured,"+
				"submitted_by,"+
				"subcategory,"+
				"category_mailchimp_tag,"+
				"use_cases,"+
				"seo,"+
			") VALUES (" +
				"$1,$2,$3,TO_TIMESTAMP($4),$5,$6,$7,$8,$9,$10"+
			")"
	_,err := ss.db.Exec(query,
		evt.MarketId,
	)
	if err != nil {
		ss.Log_msg(fmt.Sprintf("DB error: can't insert into pol_cond_res table: %v\n",err))
		os.Exit(1)
	}
}
