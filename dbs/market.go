package dbs

import (
	"fmt"
	"os"
	"math/big"
	"strings"
	"database/sql"
	"encoding/json"
	"strconv"
	_  "github.com/lib/pq"

	//"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"

	p "github.com/PredictionExplorer/augur-explorer/primitives"
)

func make_subcategories(cat_str *string) []string {
	subcategories := strings.Split(*cat_str,",")
	for i := 0 ; i< len(subcategories); i++ {
		subcategories[i] = strings.Title(subcategories[i])
	}
	if len(subcategories) > 0 {	// sometimes last category is empty, delete it
		if len(subcategories[len(subcategories)-1]) == 0 {
			subcategories = subcategories[:len(subcategories)-1]
		}
	}
	return subcategories
}
func get_market_status_str(status_code p.MarketStatus) string {

	switch p.MarketStatus(status_code) {
		case p.MktStatusReporting:
			return "Reporting"
		case p.MktStatusReported:
			return "Reported"
		case p.MktStatusDisputing:
			return "Disputing"
		case p.MktStatusFinalized:
			return "Finalized"
		case p.MktStatusFinInvalid:
			return "Finalized Invalid"
		default:
			return "Traded"
	}
	return "undefined"
}
func (ss *SQLStorage) Insert_market_created_evt(agtx *p.AugurTx,eoa_aid int64,validity_bond string,evt *p.EMarketCreated) {

	var query string
	var market_aid int64;
	market_aid = ss.Lookup_or_create_address(evt.Market.String(),agtx.BlockNum,agtx.TxId)
	// check if Market is already registered
	query = "SELECT market_aid FROM market WHERE market_aid=$1"
	err:=ss.db.QueryRow(query,market_aid).Scan(&market_aid);
	if (err!=nil) {
		if (err==sql.ErrNoRows) {
			// break
		} else {
			ss.Log_msg(fmt.Sprintf("DB error for market_aid=%v: %v, q=%v",market_aid,err,query))
			os.Exit(1)
		}
	} else {
		// market already registered, sliently exit
		return
	}
	wallet_aid := ss.Lookup_or_create_address(evt.MarketCreator.String(),agtx.BlockNum,agtx.TxId)
	universe_id,err := ss.lookup_universe_id(evt.Universe.String())
	if err != nil {
		ss.Log_msg(
			fmt.Sprintf(
				"Universe %v not found when trying to insert MarketCreated evt at block %v: %v\n",
				evt.Universe.String(),agtx.BlockNum,err,
			),
		)
		os.Exit(1)
	}
	reporter_aid := ss.Lookup_or_create_address(evt.DesignatedReporter.String(),agtx.BlockNum,agtx.TxId)
	ss.Info.Printf(
		"create_market: eoa_aid = %v, wallet_aid=%v (%v), reporter_id=%v (%v)\n",
		eoa_aid,wallet_aid,evt.MarketCreator.String(),reporter_aid,evt.DesignatedReporter.String(),
	)
	if eoa_aid==0 {
		eoa_aid = wallet_aid	// WarpSync contract creates markets with eoa_aid=0
	}
	if wallet_aid == 0 {
		ss.Log_msg(
			fmt.Sprintf(
				"insert_market_created_evt(): creator addr = %v, wallet_aid = 0, can't continue, exiting\n",
				evt.MarketCreator.String(),
			),
		)
		os.Exit(1)
	}
	prices := p.Bigint_ptr_slice_to_str(&evt.Prices,",")
	outcomes := p.Outcomes_to_str(&evt.Outcomes,",")

	var extra_info p.ExtraInfo
	json.Unmarshal([]byte(evt.ExtraInfo), &extra_info)
	categories := strings.Join(extra_info.Categories,",")
	cat_id := ss.lookup_or_create_category(categories)

	query = `
		INSERT INTO market(
			block_num,
			tx_id,
			cat_id,
			universe_id,
			market_aid,
			wallet_aid,
			eoa_aid,
			reporter_aid,
			end_time,
			num_ticks,
			create_timestamp,
			fee,
			prices,
			market_type,
			extra_info,
			outcomes,
			no_show_bond,
			validity_bond
		) VALUES ($1,$2,$3,$4,$5,$6,$7,$8,TO_TIMESTAMP($9),$10,TO_TIMESTAMP($11),` +
						evt.FeePerCashInAttoCash.String() +
						"/1e+16,$12,$13,$14,$15,(" + evt.NoShowBond.String() + "/1e+18)," +
						"("+ validity_bond + "/1e+18))";

	d_query := fmt.Sprintf( `
		INSERT INTO market(
			block_num,
			tx_id,
			cat_id,
			universe_id,
			market_aid,
			wallet_aid,
			eoa_aid,
			reporter_aid,
			end_time,
			num_ticks,
			create_timestamp,
			fee,
			prices,
			market_type,
			extra_info,
			outcomes,
			no_show_bond,
			validity_bond
		) VALUES (%v,%v,%v,%v,%v,%v,%v,%v,TO_TIMESTAMP(%v),%v,TO_TIMESTAMP(%v),` +
						evt.FeePerCashInAttoCash.String() +
						"/1e+16,'%v',%v,'%v','%v',(" + evt.NoShowBond.String() + "/1e+18)," +
						"("+ validity_bond + "/1e+18))",
			agtx.BlockNum,
			agtx.TxId,
			cat_id,
			universe_id,
			market_aid,
			wallet_aid,
			eoa_aid,
			reporter_aid,
			evt.EndTime.Int64(),
			evt.NumTicks.Int64(),
			evt.Timestamp.Int64(),
			prices,
			evt.MarketType,
			evt.ExtraInfo,
			outcomes,
	)
	_ = d_query
	result,err := ss.db.Exec(query,
			agtx.BlockNum,
			agtx.TxId,
			cat_id,
			universe_id,
			market_aid,
			wallet_aid,
			eoa_aid,
			reporter_aid,
			evt.EndTime.Int64(),
			evt.NumTicks.Int64(),
			evt.Timestamp.Int64(),
			prices,
			evt.MarketType,
			evt.ExtraInfo,
			outcomes,
	)
	if err != nil {
		ss.Log_msg(
			fmt.Sprintf(
				"DB error: can't insert into market table at block %v : %v: q=%v",
				agtx.BlockNum,err,query,
			),
		)
		os.Exit(1)
	}
	rows_affected,err:=result.RowsAffected()
	if err != nil {
		ss.Log_msg(fmt.Sprintf("DB error: %v, %q",err,query))
		os.Exit(1)
	}
	if rows_affected > 0 {
	} else {
		ss.Log_msg(fmt.Sprintf("DB error: couldn't insert into Market table. Rows affeced = 0"))
		os.Exit(1)
	}
	switch evt.MarketType {
		case 0:
			outcomes = "Invalid,No,Yes"
		case 1:
			outcomes = "Invalid," + outcomes	// Categorical
		case 2:
			outcomes = "Invalid,,Scalar"
		default:
			ss.Log_msg(
				fmt.Sprintf("Invalid market type = % for market %v",evt.MarketType,evt.Market.String()),
			)
	}
	ss.init_market_outcome_volumes(market_aid,outcomes)
}
func (ss *SQLStorage) init_market_outcome_volumes(market_aid int64,outcomes string) {

	var query string
	outcomes_list := strings.Split(outcomes,",")
	for outcome_idx:=0 ; outcome_idx < len(outcomes_list) ; outcome_idx ++ {
		if len(outcomes_list[outcome_idx])>0 {
			query = "INSERT INTO outcome_vol(" +
						"market_aid," +
						"outcome_idx" +
					") VALUES(" +
						"$1," +
						"$2" +
					")"
			d_query := fmt.Sprintf("INSERT INTO outcome_vol(" +
						"market_aid," +
						"outcome_idx" +
					") VALUES(" +
						"%v," +
						"%v" +
					")",market_aid,outcome_idx)
			ss.Info.Printf("insert into outcome volumes query: %v\n",d_query)
			_,err := ss.db.Exec(query,market_aid,outcome_idx)
			if (err!=nil) {
				ss.Log_msg(fmt.Sprintf("DB error: %v; q=%v",err,query))
				os.Exit(1)
			}
		}
	}
}
func (ss *SQLStorage) Insert_market_oi_changed_evt(block *types.Header,agtx *p.AugurTx,evt *p.EMarketOIChanged) {
	// Note: this event arrives with evt.Market set to 0x0000000000000000000000000 (a contract bug?) ,
	//			so we pass the market address as parameter ('market_addr') to the function
	var query string
	market_aid := ss.Lookup_address_id(evt.Market.String())
	ts_inserted := int64(block.Time)
	query = "INSERT INTO oi_chg(block_num,tx_id,market_aid,ts_inserted,oi) " +
			"VALUES($1,$2,$3,TO_TIMESTAMP($4),(" +
			evt.MarketOI.String() +
			"/1e+18))"
	result,err := ss.db.Exec(query,agtx.BlockNum,agtx.TxId,market_aid,ts_inserted)
	if err != nil {
		ss.Log_msg(fmt.Sprintf("DB error: can't insert into oi_chg table: %v; q=%v",err,query))
		os.Exit(1)
	}
	rows_affected,err:=result.RowsAffected()
	if err != nil {
		ss.Log_msg(fmt.Sprintf("DB error: %v, q=%v",err,query))
		os.Exit(1)
	}
	if rows_affected > 0 {
		return
	} else {
		ss.Log_msg(fmt.Sprintf("DB error: couldn't insert into oi_chg table. Rows affeced = 0"))
		os.Exit(1)
	}

	ss.Info.Printf("Set market (id=%v) open interest to %v",market_aid,evt.MarketOI.String())
}
func get_outcome_idx_from_numerators(mkt_type int,num_ticks int64,numerators []*big.Int) int {

	if mkt_type == 2 {
		return 2
	}
	big_num_ticks := big.NewInt(num_ticks)
	for i:=0 ; i < len(numerators) ; i++ {
		if big_num_ticks.Cmp(numerators[i]) == 0 {
			return i
		}
	}
	return -1
}
func (ss *SQLStorage) Insert_market_finalized_evt(agtx *p.AugurTx,timestamp int64,evt *p.EMarketFinalized) {

	var query string

	universe_id,err := ss.lookup_universe_id(evt.Universe.String())
	if err!=nil {
		ss.Log_msg(
			fmt.Sprintf(
				"Universe %v not found on insert of MarketFinalized event at block %v: %v",
				evt.Universe.String(),agtx.BlockNum,err,
			),
		)
		os.Exit(1)
	}
	_ = universe_id	// ToDo: add universe_id match condition (for market)
	market_aid := ss.Lookup_address_id(evt.Market.String())
	fin_timestamp := evt.Timestamp.Int64()
	winning_payouts := p.Bigint_ptr_slice_to_str(&evt.WinningPayoutNumerators,",")

	market_type,mticks := ss.get_market_type_and_ticks(market_aid)
	winning_outcome := get_outcome_idx_from_numerators(market_type,mticks,evt.WinningPayoutNumerators)

	query = "INSERT INTO mkt_fin(block_num,tx_id,market_aid,fin_timestamp,winning_payouts,winning_outcome)" +
			"VALUES($1,$2,$3,TO_TIMESTAMP($4),$5,$6)"
	_,err = ss.db.Exec(query,agtx.BlockNum,agtx.TxId,market_aid,fin_timestamp,winning_payouts,winning_outcome)
	if err != nil {
		ss.Log_msg(
			fmt.Sprintf(
				"DB error: can't update market finalization of market %v at block %v : %v, q=%v",
				agtx.BlockNum,market_aid,err,query,
			),
		)
		os.Exit(1)
	}
	mkt_status:=p.MktStatusFinalized
	if winning_outcome == 0 {
		mkt_status = p.MktStatusFinInvalid
	}
	ss.update_market_status(market_aid,mkt_status)
	ss.calculate_profit_loss_for_all_users(market_aid,agtx.BlockNum,agtx.TxId,timestamp,evt)
	ss.close_all_open_positions_for_market(market_aid)
}
func (ss *SQLStorage) get_market_type_and_ticks(market_aid int64) (int,int64) {

	var query string
	query = "SELECT market_type,num_ticks FROM market WHERE market_aid=$1"

	var market_type int
	var num_ticks int64
	err:=ss.db.QueryRow(query,market_aid).Scan(&market_type,&num_ticks);
	if (err!=nil) {
		d_query:=strings.ReplaceAll(query,"$1",fmt.Sprintf("%v",market_aid))
		ss.Log_msg(fmt.Sprintf("DB Error: %v, q=%v market_aid=%v\n",err,d_query,market_aid))
		os.Exit(1)
	}
	return market_type,num_ticks
}
func (ss *SQLStorage) update_market_status(market_aid int64,status p.MarketStatus) {
	var query string
	query = "UPDATE market SET status=$2 WHERE market_aid = $1"

	_,err:=ss.db.Exec(query,market_aid,status)
	if (err!=nil) {
		d_query:=strings.ReplaceAll(query,"$1",fmt.Sprintf("%v",market_aid))
		ss.Log_msg(fmt.Sprintf("DB error: %v ; q=%v",err,d_query))
		os.Exit(1)
	}
}
func (ss *SQLStorage) Insert_market_volume_changed_evt_v1(agtx *p.AugurTx,evt *p.EMarketVolumeChanged_v1) {
	// Note: this function will be discontinued after Augur is released on 28 Jul
	market_aid := ss.Lookup_address_id(evt.Market.String())

	volume := evt.Volume.String()
	outcome_vols := p.Bigint_ptr_slice_to_str(&evt.OutcomeVolumes,",")
	timestamp := evt.Timestamp.Int64()

	var query string
	query = `
		INSERT INTO volume (
			block_num,
			tx_id,
			market_aid,
			volume,
			outcome_vols,
			ins_timestamp
		) VALUES ($1,$2,$3,`+volume+`/1e+18,$4,TO_TIMESTAMP($5))`
	result,err := ss.db.Exec(query,
			agtx.BlockNum,
			agtx.TxId,
			market_aid,
			outcome_vols,
			timestamp)
	if err != nil {
		ss.Log_msg(
			fmt.Sprintf(
				"DB error: can't insert into volume table at block %v : %v, q=%v",
				agtx.BlockNum,err,query,
			),
		)
		os.Exit(1)
	}
	rows_affected,err:=result.RowsAffected()
	if err != nil {
		ss.Log_msg(fmt.Sprintf("DB error: %v, q=%v",err,query))
	}
	if rows_affected > 0 {
		//break
	} else {
		ss.Log_msg(fmt.Sprintf("DB error: couldn't insert into InitialReport table. Rows affeced = 0"))
	}

	// Updates volume per outcome in an indexed table for querying market info
	for outcome_idx := 0; outcome_idx < len(evt.OutcomeVolumes) ; outcome_idx++ {
		query = "UPDATE outcome_vol " +
				"SET " +
					"volume = "+evt.OutcomeVolumes[outcome_idx].String()+"/1e+18 " +
				"WHERE " +
					"market_aid = $1 AND outcome_idx = $2"
		_,err=ss.db.Exec(query,market_aid,outcome_idx)
		if (err!=nil) {
			ss.Log_msg(fmt.Sprintf("DB error: %v ; q=%v",err,query))
			os.Exit(1)
		}
	}
}
func (ss *SQLStorage) Insert_market_volume_changed_evt_v2(agtx *p.AugurTx,evt *p.EMarketVolumeChanged_v2) {

	market_aid := ss.Lookup_address_id(evt.Market.String())

	volume := evt.Volume.String()
	outcome_vols := p.Bigint_ptr_slice_to_str(&evt.OutcomeVolumes,",")
	var total_trades int64 = 0
	if evt.TotalTrades != nil {
		total_trades = evt.TotalTrades.Int64()
	}
	timestamp := evt.Timestamp.Int64()

	var query string
	query = `
		INSERT INTO volume (
			block_num,
			tx_id,
			market_aid,
			total_trades,
			volume,
			outcome_vols,
			ins_timestamp
		) VALUES ($1,$2,$3,$4,`+volume+`/1e+18,$5,TO_TIMESTAMP($6))`
	result,err := ss.db.Exec(query,
			agtx.BlockNum,
			agtx.TxId,
			market_aid,
			total_trades,
			outcome_vols,
			timestamp)
	if err != nil {
		ss.Log_msg(
			fmt.Sprintf(
				"DB error: can't insert into volume table at block %v : %v, q=%v",
				agtx.BlockNum,err,query,
				),
			)
		os.Exit(1)
	}
	rows_affected,err:=result.RowsAffected()
	if err != nil {
		ss.Log_msg(fmt.Sprintf("DB error: %v, q=%v",err,query))
	}
	if rows_affected > 0 {
		//break
	} else {
		ss.Log_msg(fmt.Sprintf("DB error: couldn't insert into InitialReport table. Rows affeced = 0"))
	}

	// Updates volume per outcome in an indexed table for querying market info
	for outcome_idx := 0; outcome_idx < len(evt.OutcomeVolumes) ; outcome_idx++ {
		query = "UPDATE outcome_vol " +
				"SET " +
					"volume = "+evt.OutcomeVolumes[outcome_idx].String()+"/1e+18 " +
				"WHERE " +
					"market_aid = $1 AND outcome_idx = $2"
		_,err=ss.db.Exec(query,market_aid,outcome_idx)
		if (err!=nil) {
			ss.Log_msg(fmt.Sprintf("DB error: %v ; q=%v",err,query))
			os.Exit(1)
		}
	}
}
func (ss *SQLStorage) Insert_share_balance_changed_evt(agtx *p.AugurTx,evt *p.EShareTokenBalanceChanged) {

	market_aid := ss.Lookup_address_id(evt.Market.String())
	account_aid := ss.Lookup_or_create_address(evt.Account.String(),agtx.BlockNum,agtx.TxId)

	outcome := evt.Outcome.Int64()
	balance := evt.Balance.String()

	var query string
/* DISCONTINUED , in favor of triggers
	query = "UPDATE sbalances SET balance = (" + balance + "/1e+18) " +
				"WHERE " +
					"market_aid = $1 AND " +
					"account_aid = $2 AND " +
					"outcome_idx = $3"
	result,err := ss.db.Exec(query,	market_aid,account_aid,outcome)
	if err != nil {
		ss.Log_msg(
			fmt.Sprintf(
				"DB error: can't update 'sbalances' at block %v for account %v, market %v : %v; q=%v",
					agtx.BlockNum,evt.Account.String(),evt.Market.String(),err,query,
			),
		)
		os.Exit(1)
	}
*/
/* DISCONTINUED, now inserts go into 'stbc' table
	rows_affected,err:=result.RowsAffected()
	if err != nil {
		ss.Log_msg(fmt.Sprintf("DB error: %v",err))
	}
	if rows_affected > 0 {
		//break
	} else {
		query = "INSERT INTO sbalances (" +
					"block_num," + 
					"tx_id," +
					"account_aid," +
					"market_aid," +
					"outcome_idx," +
					"balance" +
				") VALUES(" +
					"$1," +
					"$2," +
					"$3," +
					"$4," +
					"$5," +
					balance + "/1e+18" +
				")"
		result,err := ss.db.Exec(query,agtx.BlockNum,agtx.TxId,account_aid,market_aid,outcome)
		if err != nil {
			ss.Log_msg(fmt.Sprintf("DB error: can't insert into sbalances table: %v, q=%v",err,query))
			os.Exit(1)
		}
		rows_affected,err:=result.RowsAffected()
		if err != nil {
			ss.Log_msg(fmt.Sprintf("DB error: %v, query=%v",err,query))
		}
		if rows_affected > 0 {
			return
		} else {
			ss.Log_msg(fmt.Sprintf("DB error: couldn't insert into 'sbalances' table. Rows affeced = 0"))
		}
	}
	*/
	query = "INSERT INTO stbc (" +
				"block_num," + 
				"tx_id," +
				"account_aid," +
				"market_aid," +
				"outcome_idx," +
				"balance" +
			") VALUES(" +
				"$1," +
				"$2," +
				"$3," +
				"$4," +
				"$5," +
				balance + "/1e+18" +
			")"
	_,err := ss.db.Exec(query,agtx.BlockNum,agtx.TxId,account_aid,market_aid,outcome)
	if err != nil {
		ss.Log_msg(fmt.Sprintf("DB error: can't insert into sbalances table at block %v: %v, q=%v",agtx.BlockNum,err,query))
		os.Exit(1)
	}
}
func (ss *SQLStorage) Get_active_market_list(off int, lim int) []p.InfoMarket {

	var query string
	query = "SELECT " +
				"m.market_aid,"+
				"ma.addr as mkt_addr," +
				"sa.addr AS signer," +
				"ca.addr as mcreator," +
				"TO_CHAR(end_time,'dd/mm/yyyy HH24:SS UTC') as end_date," + 
				"extra_info::json->>'description' AS descr," +
				"extra_info::json->>'longDescription' AS long_desc," +
				"extra_info::json->>'categories' AS categories," +
				"outcomes," +
				"num_ticks," +
				"m.market_type, " +
				"CASE m.market_type " +
					"WHEN 0 THEN 'YES/NO' " +
					"WHEN 1 THEN 'CATEGORICAL' " +
					"WHEN 2 THEN 'SCALAR' " +
				"END AS mtype," +
				"status,"+
				"CASE WHEN EXTRACT(epoch from (fin_timestamp-now())) < 0 " +
					"THEN 'Trading' ELSE 'Reporting' END AS status_desc," +
				"fee," +
				"money_at_stake, " +
				"open_interest AS OI," +
				"cur_volume AS volume " +
			"FROM market as m " +
				"LEFT JOIN " +
					"address AS ma ON m.market_aid = ma.address_id " +
				"LEFT JOIN " +
					"address AS sa ON m.eoa_aid= sa.address_id " +
				"LEFT JOIN " +
					"address AS ca ON m.wallet_aid = ca.address_id " +
			"WHERE m.status < 4 " +
			"ORDER BY " +
				"m.fin_timestamp DESC " +
			"OFFSET $1 LIMIT $2";

	rows,err := ss.db.Query(query,off,lim)
	if (err!=nil) {
		ss.Log_msg(fmt.Sprintf("DB error: %v (query=%v)",err,query))
		os.Exit(1)
	}
	var rec p.InfoMarket
	records := make([]p.InfoMarket,0,8)

	defer rows.Close()
	for rows.Next() {
		var description sql.NullString
		var longdesc sql.NullString
		var category sql.NullString
		err=rows.Scan(
					&rec.MktAid,
					&rec.MktAddr,
					&rec.Signer,
					&rec.MktCreator,
					&rec.EndDate,
					&description,
					&longdesc,
					&category,
					&rec.Outcomes,
					&rec.NumTicks,
					&rec.MktType,
					&rec.MktTypeStr,
					&rec.MktStatus,
					&rec.Status,
					&rec.Fee,
					&rec.MoneyAtStake,
					&rec.OpenInterest,
					&rec.CurVolume,
		)
		if err!=nil {
			ss.Log_msg(fmt.Sprintf("DB error: %v, q=%v",err,query))
			os.Exit(1)
		}
		rec.MktAddrSh=p.Short_address(rec.MktAddr)
		rec.MktCreatorSh=p.Short_address(rec.MktCreator)
		if description.Valid {
			rec.Description = description.String
		}
		if longdesc.Valid {
			rec.LongDesc = longdesc.String
		}
		if category.Valid {
			rec.CategoryStr=category.String
		}
		rec.Status = get_market_status_str(p.MarketStatus(rec.MktStatus))
		records = append(records,rec)
	}
	return records
}
func (ss *SQLStorage) Get_active_market_ids(sort int,all int,fin int,off int, lim int) []int64 {

	var order_condition string
	switch sort {
		case 1: order_condition = "m.money_at_stake DESC,m.market_aid DESC";
		case 2: order_condition = "m.cur_volume DESC, m.market_aid DESC";
		case 3: order_condition = "m.create_timestamp DESC,m.market_aid DESC";
		case 4: order_condition = "m.end_time DESC,m.market_aid DESC";
		case 5: order_condition = "m.fin_timestamp DESC,m.market_aid DESC";
		default:
			order_condition = "m.market_aid DESC";
	}
	var where_condition string
	if fin == 0 {
		where_condition = "(m.status < 4) "
	} else {
		where_condition = "(m.status > 3) "
	}
	if all == 0 {
		where_condition = where_condition + " AND (m.cur_volume > 0) AND (m.money_at_stake > 0) "
	}
	var query string
	query = "SELECT " +
				"market_aid,substring(extra_info::json->>'description',1,43) as descr " +
			"FROM market as m " +
			"WHERE " + where_condition +
			"ORDER BY " + order_condition + " " +
			"OFFSET $1 LIMIT $2";

	rows,err := ss.db.Query(query,off,lim)
	if (err!=nil) {
		ss.Log_msg(fmt.Sprintf("DB error: %v (query=%v)",err,query))
		os.Exit(1)
	}
	records := make([]int64,0,32)

	defer rows.Close()
	for rows.Next() {
		var market_aid int64
		var null_descr sql.NullString
		err=rows.Scan(&market_aid,&null_descr)
		if err!=nil {
			ss.Log_msg(fmt.Sprintf("DB error: %v, q=%v",err,query))
			os.Exit(1)
		}
		if null_descr.Valid {
			if null_descr.String == "What will the next Augur Warp Sync hash be?"  {
//				continue			// skipping internal Augur Markets
			}
		}
		records = append(records,market_aid)
	}
	return records
}
func (ss *SQLStorage) Get_market_card_data(id int64) (p.InfoMarket,error) {
	var rec p.InfoMarket

	var query string
	query = "SELECT " +
				"m.market_aid," +
				"ma.addr as mkt_addr," +
				"sa.addr AS signer," +
				"ca.addr as mcreator," +
				"TO_CHAR(end_time,'dd/mm/yyyy HH24:SS UTC') as end_date," + 
				"extra_info::json->>'description' AS descr," +
				"extra_info::json->>'longDescription' AS long_desc," +
				"extra_info::json->>'categories' AS categories," +
				"outcomes," +
				"num_ticks," +
				"m.market_type, " +
				"CASE m.market_type " +
					"WHEN 0 THEN 'YES/NO' " +
					"WHEN 1 THEN 'CATEGORICAL' " +
					"WHEN 2 THEN 'SCALAR' " +
				"END AS mtype," +
				"status,"+
				"CASE WHEN EXTRACT(epoch from (fin_timestamp-now())) < 0 " +
					"THEN 'Trading' ELSE 'Reporting' END AS status_desc," +
				"fee," +
				"money_at_stake, " +
				"open_interest AS OI," +
				"cur_volume AS volume " +
			"FROM market as m " +
				"LEFT JOIN " +
					"address AS ma ON m.market_aid = ma.address_id " +
				"LEFT JOIN " +
					"address AS sa ON m.eoa_aid= sa.address_id " +
				"LEFT JOIN " +
					"address AS ca ON m.wallet_aid = ca.address_id " +
			"WHERE m.market_aid=$1 " +
			"ORDER BY " +
				"m.fin_timestamp DESC "


	var description sql.NullString
	var longdesc sql.NullString
	var category sql.NullString
	res := ss.db.QueryRow(query,id)
	err := res.Scan(
			&rec.MktAid,
			&rec.MktAddr,
			&rec.Signer,
			&rec.MktCreator,
			&rec.EndDate,
			&description,
			&longdesc,
			&category,
			&rec.Outcomes,
			&rec.NumTicks,
			&rec.MktType,
			&rec.MktTypeStr,
			&rec.MktStatus,
			&rec.Status,
			&rec.Fee,
			&rec.MoneyAtStake,
			&rec.OpenInterest,
			&rec.CurVolume,
	)
	if (err!=nil) {
		if err!=sql.ErrNoRows {
			ss.Log_msg(fmt.Sprintf("DB error querying card for id=%v : %v (query=%v)",id,err,query))
		}
		return rec,err
	}
	rec.MktAddrSh=p.Short_address(rec.MktAddr)
	rec.MktCreatorSh=p.Short_address(rec.MktCreator)
	if description.Valid {
		rec.Description = description.String
	}
	if longdesc.Valid {
		rec.LongDesc = longdesc.String
	}
	if category.Valid {
		rec.CategoryStr=category.String
	}
	rec.Status = get_market_status_str(p.MarketStatus(rec.MktStatus))


	volumes,err := ss.Get_outcome_volumes(rec.MktAddr,id,1)
	if err!=nil {
		ss.Log_msg(fmt.Sprintf("DB error querying card for id=%v : %v (query=%v)",id,err,query))
		return rec,err
	}
	rec.OutcomeVolumes = volumes

	return rec,nil
}
func (ss *SQLStorage) Get_categories() []p.InfoCategories {

	var query string
	query = "SELECT " +
				"cat_id," +
				"total_markets," +
				"category " +
			"FROM category " +
			"ORDER BY " +
				"category";

	rows,err:=ss.db.Query(query)
	if err!=nil {
		if err!=sql.ErrNoRows {
			ss.Log_msg(fmt.Sprintf("Error for query %v: %v",query,err))
			os.Exit(1)
		}
	}
	var rec p.InfoCategories
	records := make([]p.InfoCategories,0,8)

	defer rows.Close()
	for rows.Next() {
		err=rows.Scan(&rec.CatId,&rec.TotalMarkets,&rec.Category)
		if err!=nil {
			ss.Log_msg(fmt.Sprintf("DB error: %v",err))
			os.Exit(1)
		}
		rec.Subcategories = make_subcategories(&rec.Category)
		records = append(records,rec)
	}
	return records
}
func get_outcome_str(mkt_type uint8,outcome_idx int,outcomes_str *string) string {

		var output string
		if mkt_type == 0 { // Yes/No
			switch outcome_idx {
				case 0:
					output = "Invalid"
				case 1:
					output = "No"
				case 2:
					output = "Yes"
			}
		}
		if mkt_type == 1 { // Categorical
			outcomes_list := strings.Split(*outcomes_str,",")
			if outcome_idx == 0 {
				output = "Invalid"
			} else {
				output = outcomes_list[outcome_idx-1]
			}
		}
		if mkt_type == 2 {
			if outcome_idx == 0 {
				output = "Invalid"
			}
			if outcome_idx == 2 {
				output = "Scalar"
			}
			if outcome_idx == 1 {
				output = "-"
			}
		}
		return output
}
func (ss *SQLStorage) Get_market_info(mkt_addr string,outcome_idx int,oc bool) (p.InfoMarket,error) {
	// Inputs: 
	//		mkt_addr			address of the market to get the data from
	//		outcome_idx			narrow search by specific outcome
	//		oc					format outcome as string (from the integer parameter in the args)
	var rec p.InfoMarket
	market_aid,err := ss.Nonfatal_lookup_address_id(mkt_addr)
	if err != nil {
		return rec,err
	}
	rec.MktAid=market_aid
	var reporter_aid int64
	var query string
	query = "SELECT " +
				"m.market_type," +
				"ma.addr as mkt_addr," +
				"sa.addr AS signer," +
				"ca.addr as mcreator," +
				"ra.addr AS reporter,"+
				"reporter_aid," +
				"TO_CHAR(end_time,'dd/mm/yyyy HH24:SS UTC') AS end_date," + 
				"extra_info::json->>'description' AS descr," +
				"extra_info::json->>'longDescription' AS long_desc," +
				"cat.category," +
				"outcomes," +
				"m.market_type, " +
				"CASE m.market_type " +
					"WHEN 0 THEN 'YES/NO' " +
					"WHEN 1 THEN 'CATEGORICAL' " +
					"WHEN 2 THEN 'SCALAR' " +
				"END AS mtype, " +
				"m.status," +
				"round(fee,2) as fee," +
				"open_interest AS OI," +
				"cur_volume AS volume, " +
				"total_trades," +
				"money_at_stake " +
			"FROM market as m " +
				"LEFT JOIN address AS ma ON m.market_aid = ma.address_id " +
				"LEFT JOIN address AS sa ON m.eoa_aid = sa.address_id " +
				"LEFT JOIN address AS ca ON m.wallet_aid = ca.address_id " +
				"LEFT JOIN address AS ra ON m.reporter_aid = ra.address_id " +
				"LEFT JOIN category AS cat On m.cat_id = cat.cat_id " +
			"WHERE market_aid = $1"

	row := ss.db.QueryRow(query,market_aid)
	var mkt_type int
	var description sql.NullString
	var long_desc sql.NullString
	var category sql.NullString
	err=row.Scan(
				&mkt_type,
				&rec.MktAddr,
				&rec.Signer,
				&rec.MktCreator,
				&rec.Reporter,
				&reporter_aid,
				&rec.EndDate,
				&description,
				&long_desc,
				&category,
				&rec.Outcomes,
				&rec.MktType,
				&rec.MktTypeStr,
				&rec.MktStatus,
				&rec.Fee,
				&rec.OpenInterest,
				&rec.CurVolume,
				&rec.TotalTrades,
				&rec.MoneyAtStake,
	)
	rec.MktAddrSh=p.Short_address(rec.MktAddr)
	rec.SignerSh=p.Short_address(rec.Signer)
	rec.MktCreatorSh=p.Short_address(rec.MktCreator)
	rec.ReporterSh=p.Short_address(rec.Reporter)
	if description.Valid {
		rec.Description=description.String
	}
	if long_desc.Valid {
		rec.LongDesc = long_desc.String
	}
	if category.Valid {
		rec.CategoryStr = category.String
	}
	if oc { // get outcome string
		rec.CurOutcome = get_outcome_str(uint8(mkt_type),outcome_idx,&rec.Outcomes)
	}
	if err!=nil {
		if err == sql.ErrNoRows {
			return rec,err
		}
		ss.Log_msg(fmt.Sprintf("DB error: %v, q=%v\n",err,query))
		os.Exit(1)
	}
	rec.Status=get_market_status_str(p.MarketStatus(rec.MktStatus))
	reporter_eoa_aid,err := ss.Lookup_eoa_aid(reporter_aid)
	if err == nil {
		rep_addr,err := ss.Lookup_address(reporter_eoa_aid)
		if err == nil {
			rec.Reporter = rep_addr
			rec.ReporterSh = string(rep_addr[0:6]+string('…')+rep_addr[26:32])
		}
	}
	subcategories := make_subcategories(&rec.CategoryStr)
	rec.Subcategories = subcategories

	return rec,nil
}
func (ss *SQLStorage) Get_outcome_volumes(mkt_addr string,market_aid int64,orderby int) ([]p.OutcomeVol,error) {


	var rec p.OutcomeVol
	records := make([]p.OutcomeVol,0,8)

	var orderby_str = "ORDER BY o.outcome_idx"
	if orderby == 1 {
		orderby_str = "ORDER BY o.last_price DESC"
	}
	var query string
	query = "SELECT " +
				"o.outcome_idx, " +
				"o.volume," +
				"o.last_price, " +
				"m.market_type, " +
				"m.outcomes " +
			"FROM outcome_vol AS o " +
				"LEFT JOIN " +
					"market AS m ON o.market_aid = m.market_aid " +
			"WHERE o.market_aid = $1 " +
			orderby_str

	var rows *sql.Rows
	rows,err := ss.db.Query(query,market_aid)
	if (err!=nil) {
		ss.Log_msg(fmt.Sprintf("DB error: %v (query=%v)",err,query))
		os.Exit(1)
	}

	defer rows.Close()
	for rows.Next() {
		var outcomes string
		err=rows.Scan(
			&rec.Outcome,
			&rec.Volume,
			&rec.LastPrice,
			&rec.MktType,
			&outcomes,
		)
		rec.MktAddr = mkt_addr
		if err!=nil {
			ss.Log_msg(fmt.Sprintf("DB error: %v, q=%v",err,query))
			os.Exit(1)
		}
		rec.OutcomeStr = get_outcome_str(uint8(rec.MktType),rec.Outcome,&outcomes)
		records = append(records,rec)
	}
	return records,nil
}
func (ss *SQLStorage) Get_category_markets(cat_id int64) []p.InfoMarket {

	var query string
	query = "SELECT " +
				"ma.addr as mkt_addr," +
				"sa.addr AS signer," +
				"ca.addr as mcreator," +
				"TO_CHAR(end_time,'dd/mm/yyyy HH24:SS UTC') as end_date," + 
				"extra_info::json->>'description' AS descr," +
				"extra_info::json->>'longDescription' AS long_desc," +
				"extra_info::json->>'categories' AS categories," +
				"outcomes," +
				"m.market_type, " +
				"CASE m.market_type " +
					"WHEN 0 THEN 'YES/NO' " +
					"WHEN 1 THEN 'CATEGORICAL' " +
					"WHEN 2 THEN 'SCALAR' " +
				"END AS mtype," +
				"status,"+
				"CASE WHEN EXTRACT(epoch from (fin_timestamp-now())) < 0 " +
					"THEN 'Trading' ELSE 'Reporting' END AS status_desc," +
				"fee," +
				"open_interest AS OI," +
				"cur_volume AS volume " +
			"FROM market as m " +
				"LEFT JOIN " +
					"address AS ma ON m.market_aid = ma.address_id " +
				"LEFT JOIN " +
					"address AS sa ON m.eoa_aid= sa.address_id " +
				"LEFT JOIN " +
					"address AS ca ON m.wallet_aid = ca.address_id " +
			"WHERE cat_id = $1 " +
			"ORDER BY m.market_aid "

	rows,err := ss.db.Query(query,cat_id)
	if (err!=nil) {
		ss.Log_msg(fmt.Sprintf("DB error: %v (query=%v)",err,query))
		os.Exit(1)
	}
	var rec p.InfoMarket
	records := make([]p.InfoMarket,0,8)

	defer rows.Close()
	for rows.Next() {
		var description sql.NullString
		var longdesc sql.NullString
		var category_str sql.NullString
		err=rows.Scan(
					&rec.MktAddr,
					&rec.Signer,
					&rec.MktCreator,
					&rec.EndDate,
					&description,
					&longdesc,
					&category_str,
					&rec.Outcomes,
					&rec.MktType,
					&rec.MktTypeStr,
					&rec.MktStatus,
					&rec.Status,
					&rec.Fee,
					&rec.OpenInterest,
					&rec.CurVolume,
		)
		if err!=nil {
			ss.Log_msg(fmt.Sprintf("DB error: %v, q=%v",err,query))
			os.Exit(1)
		}
		if description.Valid {
			rec.Description = description.String
		}
		if longdesc.Valid {
			rec.LongDesc = longdesc.String
		}
		if category_str.Valid {
			rec.CategoryStr = category_str.String
		}
		rec.Status=get_market_status_str(p.MarketStatus(rec.MktStatus))
		rec.MktAddrSh=p.Short_address(rec.MktAddr)
		rec.MktCreatorSh=p.Short_address(rec.MktCreator)
		records = append(records,rec)
	}
	return records
}
func (ss *SQLStorage) Get_market_reports(market_aid int64,limit int) []p.Report {

	var query string
	query = "SELECT " +
				"r.rpt_timestamp::date," +
				"ma.addr as mkt_addr," +
				"r.is_initial," +
				"r.is_designated," +
				"round(r.amount_staked,2),"+
				"r.outcome_idx," +
				"r.next_win_start," +
				"r.next_win_end," +
				"m.initial_outcome," +
				"m.designated_outcome," +
				"m.winning_outcome," +
				"m.market_type AS mtype," +
				"m.outcomes AS outcomes_str " +
			"FROM " +
					"report AS r, " +
					"market AS m " +
						"LEFT JOIN address AS ma ON m.market_aid = ma.address_id " +
			"WHERE (r.market_aid = m.market_aid) and (r.market_aid=$1) " +
			"ORDER BY r.rpt_timestamp"
	if limit > 0 {
		query = query +	" LIMIT " + strconv.Itoa(limit)
	}

	records := make([]p.Report,0,8)
	var rows *sql.Rows
	var err error
	rows,err = ss.db.Query(query,market_aid)
	if (err!=nil) {
		if err == sql.ErrNoRows {
			return records
		}
		ss.Log_msg(fmt.Sprintf("DB error: %v (query=%v) market_aid=%v",err,query,market_aid))
		os.Exit(1)
	}

	defer rows.Close()
	for rows.Next() {
		var rec p.Report
		var mkt_type int
		var designated_outcome int
		var winning_outcome int
		var initial_outcome int
		var outcomes string
		err=rows.Scan(
			&rec.Date,
			&rec.MktAddr,
			&rec.IsInitial,
			&rec.IsDesignated,
			&rec.RepStake,
			&rec.OutcomeIdx,
			&rec.WinStart,
			&rec.WinEnd,
			&initial_outcome,
			&designated_outcome,
			&winning_outcome,
			&rec.MktType,
			&outcomes,
		)
		if err!=nil {
			ss.Log_msg(fmt.Sprintf("DB error: %v, q=%v",err,query))
			os.Exit(1)
		}
		rec.MktAddrSh=p.Short_address(rec.MktAddr)
		if winning_outcome == -1 {	// market wasn't finalized yet
			if designated_outcome == -1 {
				rec.ReportType="CROWDSOURCED"
			} else {
				if designated_outcome == rec.OutcomeIdx {
					rec.ReportType = "SUPPORTING"
				} else {
					rec.ReportType = "DISPUTING"
				}
			}
		} else {					// market was finalized
			if designated_outcome == -1 {	// designated reporter never showed up
				if initial_outcome == rec.OutcomeIdx {
					rec.ReportType = "SUPPORTING"
				} else {
					rec.ReportType = "DISPUTING"
				}
			} else {
				if designated_outcome == rec.OutcomeIdx {
					rec.ReportType = "SUPPORTING"
				} else {
					rec.ReportType = "DISPUTING"
				}
			}
		}
		rec.OutcomeStr = get_outcome_str(uint8(mkt_type),rec.OutcomeIdx,&outcomes)
		records = append(records,rec)
	}
	return records
}
