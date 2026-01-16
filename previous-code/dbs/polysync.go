package dbs
import (
	"os"
	"fmt"
	"database/sql"
	p "github.com/PredictionExplorer/augur-explorer/primitives"
)
func (ss *SQLStorage) Insert_polymarket_from_web_api(
	evt *p.Pol_Market_API_Record_JSON_v2,
	evt_comp *p.Pol_Market_API_Record_Complementary,
	block_num int64,
	tx_id int64,
) {

	var outcomes string
	for i:=0;i<len(evt.Outcomes);i++ {
		if len(outcomes)>0 { outcomes = outcomes + "," }
		outcomes = outcomes + evt.Outcomes[i]
	}
	var tags string
	for i:=0;i<len(evt.Tags);i++ {
		if len(tags)>0 { tags = tags + "," }
		tags = tags + evt.Tags[i]
	}
	var outcome_prices string
	for i:=0;i<len(evt.OutcomePrices);i++ {
		if len(outcome_prices)>0 { outcome_prices = outcome_prices + "," }
		outcome_prices = outcome_prices + evt.OutcomePrices[i]
	}
	var use_cases string
	for i:=0;i<len(evt.UseCases);i++ {
		if len(use_cases) > 0 { use_cases = use_cases + "," }
		use_cases = use_cases + "{"+
			fmt.Sprintf("%v",evt.UseCases[i].Id)+","+
			evt.UseCases[i].Title+","+
			evt.UseCases[i].Content + 
		"}"
	}
	var group_item_title string
	var group_item_threshold string
	for i:=0;i<len(evt.MarketGroup.Markets);i++ {
		market := evt.MarketGroup.Markets[i]
		if market.Id == evt.MarketId {
			group_item_title = market.GroupItemTitle
			group_item_threshold = market.GroupItemThreshold
		}
	}
	market_maker_aid:=ss.Lookup_or_create_address(evt.MarketMakerAddr,block_num,tx_id)
	var query string
	query = "INSERT INTO pol_market (" +
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
				"market_type_code,"+
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
				"liquidity," +
				"mkt_group_question,"+
				"group_item_title,"+
				"group_item_threshold"+
			") VALUES (" +
				"$1,$2,$3,$4,$5,$6,$7,"+
				"TO_TIMESTAMP($8)," + // end_date_ts
				"$9,$10,$11,$12,$13,"+
				"TO_TIMESTAMP($14)," + // start_date_ts
				"$15,$16,$17,"+
				"CAST(COALESCE(NULLIF($18,''),'0') AS DECIMAL),"+	//fee
				"$19,$20,"+
				"CAST(COALESCE(NULLIF($21,''),'0') AS DECIMAL),"+ // lower_bound
				"CAST(COALESCE(NULLIF($22,''),'0') AS DECIMAL),"+ // upper bound
				"$23,$24,$25,$26," +
				"CAST(COALESCE(NULLIF($27,''),'0') AS DECIMAL),"+	//volume
				"$28,$29,$30,$31,$32,"+
				"TO_TIMESTAMP($33),"+ //lower_bound_ts
				"$34,"+
				"TO_TIMESTAMP($35),"+ //upper_bound_ts
				"$36,$37,$38,"+
				"TO_TIMESTAMP($39),"+ //created_at_ts
				"$40,"+
				"TO_TIMESTAMP($41),"+ //updated_at_ts
				"$42," +
				"TO_TIMESTAMP($43),"+ // closed_time_ts
				"$44,$45,$46,$47,$48,$49,$50,$51,$52,"+
				"CAST(COALESCE(NULLIF($53,''),'0') AS DECIMAL),"+	// liquidity
				"$54,$55,$56"+
			")"
	_,err := ss.db.Exec(query,
		evt.MarketId,
		evt.Question,
		evt.ConditionId,
		evt.Slug,
		evt.TwitterCardImage,
		evt.ResolutionSource,
		evt.EndDate,
		evt_comp.EndDateTs,
		evt.Category,
		evt.AmmType,
		evt.SponsorName,
		evt.SponsorImage,
		evt.StartDate,
		evt_comp.StartDateTs,
		evt.XAxisValue,
		evt.YAxisValue,
		evt.DenominationToken,
		evt.Fee,
		evt.Image,
		evt.Icon,
		evt.LowerBound,
		evt.UpperBound,
		evt.Description,
		tags,
		outcomes,
		outcome_prices,
		evt.Volume,
		evt.Active,
		evt.MarketType,
		evt_comp.MarketTypeCode,
		evt.FormatType,
		evt.LowerBoundDate,
		evt_comp.LowerBoundDateTs,
		evt.UpperBoundDate,
		evt_comp.UpperBoundDateTs,
		evt.Closed,
		market_maker_aid,
		evt.CreatedAtDate,
		evt_comp.CreatedAtDateTs,
		evt.UpdatedAt,
		evt_comp.UpdatedAtTs,
		evt.ClosedTimeDate,
		evt_comp.ClosedTimeDateTs,
		evt.WideFormat,
		evt.New,
		evt.SentDiscord,
		evt.MailChimpTag,
		evt.Featured,
		evt.SubmittedBy,
		evt.Subcategory,
		evt.CategoryMailChimpTag,
		use_cases,
		evt.Liquidity,
		evt.MarketGroup.Question,
		group_item_title,
		group_item_threshold,
	)
	if err != nil {
		ss.Log_msg(fmt.Sprintf("DB error: can't insert into 'pol_market' table: %v\n",err))
		os.Exit(1)
	}
}
func (ss *SQLStorage) Market_exists(market_id int64) bool {

	var query string
	query = "SELECT market_id FROM pol_market WHERE market_id=$1"
	
	row := ss.db.QueryRow(query,market_id)
	var null_id sql.NullInt64
	var err error
	err=row.Scan(&null_id);
	if (err!=nil) {
		if err == sql.ErrNoRows {
			return false
		} else {
			ss.Log_msg(fmt.Sprintf("Error in Market_exists()(): %v",err))
			os.Exit(1)
		}
	}
	if (null_id.Valid) {
		return true
	} else {
		return false
	}

}
