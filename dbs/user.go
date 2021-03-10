package dbs

import (
	"fmt"
	"os"
	"strings"
	"strconv"
	"math/big"
	"database/sql"
	_  "github.com/lib/pq"

	"github.com/ethereum/go-ethereum/common"

	p "github.com/PredictionExplorer/augur-explorer/primitives"
)
func (ss *SQLStorage) fill_block_info(ui *p.UserInfo,user_aid int64) {

	var query string
	query = "SELECT address_id,addr,b.block_num, " +
			"FLOOR(EXTRACT(EPOCH FROM b.ts))::BIGINT as ts " +
			"FROM address a,block b " +
			"WHERE (a.address_id=$1) AND (a.block_num=b.block_num) "
	row := ss.db.QueryRow(query,user_aid)
	err := row.Scan(&ui.Aid,&ui.Addr,&ui.BlockNum,&ui.TimeStamp)
	if err != nil {
		if err!=sql.ErrNoRows {
			ss.Log_msg(fmt.Sprintf("DB error: %v, q=%v\n",err,query))
			os.Exit(1)
		}
	}
}
func (ss *SQLStorage) Get_user_info(user_aid int64) (p.UserInfo,error) {


	var ui p.UserInfo
	ss.fill_block_info(&ui,user_aid)
	ui.AugurFlags = ss.Get_augur_flags(user_aid)

	var query string
	query = "SELECT " +
				"act.address_id, " +
				"ea.eoa_aid AS eoa_aid," +
				"ea.addr AS eoa_addr," +
				"ew.wallet_aid AS wallet_aid," +
				"ew.addr AS wallet_addr," +
				"s.total_trades," +
				"s.markets_created," +
				"s.markets_traded," +
				"s.withdraw_reqs," +
				"s.deposit_reqs," +
				"s.total_reports," +
				"s.total_designated," +
				"s.profit_loss," +
				"s.report_profits," +
				"s.aff_profits," +
				"s.money_at_stake," +
				"s.validity_bonds," +
				"s.total_withdrawn," +
				"s.total_deposited, " +
				"r.top_trades, " +
				"r.top_profit, " +
				"ds.balancer_swaps," +
				"ds.uniswap_swaps " +
			"FROM address AS act " +
				"LEFT JOIN ustats AS s ON act.address_id = s.aid "+
				"LEFT JOIN uranks AS r ON act.address_id = r.aid " +
				"LEFT JOIN LATERAL (" +
					"SELECT eoa_aid,wallet_aid,a.addr FROM eoa_wallet ea " +
					"JOIN address a ON ea.eoa_aid=a.address_id " +
				") AS ea ON act.address_id=ea.wallet_aid " +
				"LEFT JOIN LATERAL (" +
					"SELECT eoa_aid,wallet_aid,a.addr FROM eoa_wallet ew " +
					"JOIN address a ON ew.wallet_aid=a.address_id " +
				") AS ew ON act.address_id=ew.eoa_aid " +
				"LEFT JOIN defi_stats AS ds ON act.address_id=ds.aid " +
				"WHERE act.address_id=$1"
	d_query := strings.ReplaceAll(query,`$1`,fmt.Sprintf("%v",user_aid))
	ss.Info.Printf("query = %v\n",d_query)
	row := ss.db.QueryRow(query,user_aid)
	var err error
	var (
		top_profits					sql.NullFloat64
		top_trades				sql.NullFloat64
		wallet_addr				sql.NullString
		wallet_aid				sql.NullInt64
		eoa_addr				sql.NullString
		eoa_aid					sql.NullInt64
		null_total_trades		sql.NullInt32
		null_markets_created	sql.NullInt32
		null_markets_traded		sql.NullInt32
		null_withdraw_reqs		sql.NullInt32
		null_deposit_reqs		sql.NullInt32
		null_total_reports		sql.NullInt32
		null_total_designated	sql.NullInt32
		null_profit_loss		sql.NullFloat64
		null_report_profits		sql.NullFloat64
		null_aff_profits		sql.NullFloat64
		null_money_at_stake		sql.NullFloat64
		null_validity_bonds		sql.NullFloat64
		null_total_withdrawn	sql.NullFloat64
		null_total_deposited	sql.NullFloat64
		null_balancer_swaps		sql.NullInt64
		null_uniswap_swaps		sql.NullInt64
	)
	ui.Aid = user_aid
	err=row.Scan(
				&ui.Aid,
				&eoa_aid,
				&eoa_addr,
				&wallet_aid,
				&wallet_addr,
				&null_total_trades,
				&null_markets_created,
				&null_markets_traded,
				&null_withdraw_reqs,
				&null_deposit_reqs,
				&null_total_reports,
				&null_total_designated,
				&null_profit_loss,
				&null_report_profits,
				&null_aff_profits,
				&null_money_at_stake,
				&null_validity_bonds,
				&null_total_withdrawn,
				&null_total_deposited,
				&top_trades,
				&top_profits,
				&null_balancer_swaps,
				&null_uniswap_swaps,
	);
	if (err!=nil) {
		if err == sql.ErrNoRows {
			if ui.BlockNum > 0 {
				return ui,nil
			} else {
				return ui,err
			}
		} else {
			ss.Log_msg(fmt.Sprintf("DB error: %v, q=%v\n",err,query))
		}
		os.Exit(1)
	}
	ui.Addr,err = ss.Lookup_address(user_aid)
	if err == nil {
		ui.AddrSh = p.Short_address(ui.Addr)
	}
	if null_total_trades.Valid { ui.TotalTrades = null_total_trades.Int32 }
	if null_markets_created.Valid { ui.MarketsCreated = null_markets_created.Int32 }
	if null_markets_traded.Valid { ui.MarketsTraded = null_markets_traded.Int32 }
	if null_withdraw_reqs.Valid { ui.WithdrawReqs = null_withdraw_reqs.Int32 }
	if null_deposit_reqs.Valid { ui.DepositReqs = null_deposit_reqs.Int32 }
	if null_total_reports.Valid { ui.TotalReports = null_total_reports.Int32 }
	if null_total_designated.Valid { ui.TotalDesignated = null_total_designated.Int32 }
	if null_profit_loss.Valid { ui.ProfitLoss = null_profit_loss.Float64 }
	if null_report_profits.Valid { ui.ReportProfits = null_report_profits.Float64 }
	if null_aff_profits.Valid { ui.AffProfits = null_aff_profits.Float64 }
	if null_money_at_stake.Valid { ui.MoneyAtStake = null_money_at_stake.Float64 }
	if null_validity_bonds.Valid { ui.ValidityBonds = null_validity_bonds.Float64 }
	if null_total_withdrawn.Valid { ui.TotalWithdrawn = null_total_withdrawn.Float64 }
	if null_total_deposited.Valid { ui.TotalDeposited = null_total_deposited.Float64 }

	if top_profits.Valid { ui.TopProfit = top_profits.Float64 }
	if top_trades.Valid { ui.TopTrades = top_trades.Float64 }
	if ui.MoneyAtStake < 0 { ui.HedgingProfits = true }
	ui.UnclaimedProfit=ss.Get_unclaimed_profit(user_aid)
	if wallet_aid.Valid { ui.WalletAid = wallet_aid.Int64; ui.WalletAddr = wallet_addr.String }
	if eoa_aid.Valid { ui.EOAAid = eoa_aid.Int64; ui.EOAAddr = eoa_addr.String }
	if null_balancer_swaps.Valid { ui.BalancerNumSwaps = null_balancer_swaps.Int64 }
	if null_uniswap_swaps.Valid { ui.UniswapNumSwaps = null_uniswap_swaps.Int64 }

	return ui,nil
}
func (ss *SQLStorage) Get_ranking_data_for_all_users() []p.RankStats {

	var query string
	query = "SELECT aid,total_trades,profit_loss,volume_traded FROM ustats"

	rows,err := ss.db.Query(query)
	if (err!=nil) {
		ss.Log_msg(fmt.Sprintf("DB error: %v (query=%v)",err,query))
		os.Exit(1)
	}
	records := make([]p.RankStats,0,8)
	defer rows.Close()
	for rows.Next() {
		var rec p.RankStats
		err=rows.Scan(&rec.Aid,&rec.TotalTrades,&rec.ProfitLoss,&rec.VolumeTraded)
		records = append(records,rec)
	}
	return records
}
func (ss *SQLStorage) Update_top_profit_rank(aid int64,value float64,profit float64) int64 {

	var query string
	query = "UPDATE uranks SET top_profit = $2,profit=$3 WHERE aid = $1"
	res,err:=ss.db.Exec(query,aid,value,profit)
	if (err!=nil) {
		ss.Log_msg(fmt.Sprintf("update_top_profit_rank() failed: %v, q=%v",err,query))
		os.Exit(1)
	}
	affected_rows,err:=res.RowsAffected()
	if err!=nil {
		ss.Log_msg(fmt.Sprintf("Error getting RowsAffected in update_top_profit_rank(): %v",err))
	}
	if affected_rows == 0 {
		query = "INSERT INTO uranks(aid,top_profit,profit) VALUES($1,$2,$3)"
		_,err:=ss.db.Exec(query,aid,value,profit)
		if (err!=nil) {
			ss.Log_msg(fmt.Sprintf("update_top_profit_rank() failed: %v, q=%v",err,query))
		}

	}
	return affected_rows
}
func (ss *SQLStorage) Update_top_total_trades_rank(aid int64,value float64,total_trades int64) int64 {

	var query string
	query = "UPDATE uranks SET top_trades = $2,total_trades=$3 WHERE aid = $1"
	res,err:=ss.db.Exec(query,aid,value,total_trades)
	if (err!=nil) {
		ss.Log_msg(fmt.Sprintf("update_top_total_trades_rank() failed: %v, q=%v",err,query))
		os.Exit(1)
	}
	affected_rows,err:=res.RowsAffected()
	if err!=nil {
		ss.Log_msg(fmt.Sprintf("Error getting RowsAffected in update_top_total_trades_rank(): %v",err))
	}
	if affected_rows == 0 {
		query = "INSERT INTO uranks(aid,top_trades,total_trades) VALUES($1,$2,$3)"
		_,err:=ss.db.Exec(query,aid,value,total_trades)
		if (err!=nil) {
			ss.Log_msg(fmt.Sprintf("update_top_total_trades_rank() failed: value=%v, err: %v, q=%v",value,err,query))
			os.Exit(1)
		}

	}
	return affected_rows
}
func (ss *SQLStorage) Update_top_volume_rank(aid int64,value float64,volume float64) int64 {

	var query string
	query = "UPDATE uranks SET top_volume = $2,volume=$3 WHERE aid = $1"
	res,err:=ss.db.Exec(query,aid,value,volume)
	if (err!=nil) {
		ss.Log_msg(fmt.Sprintf("update_top_volume_rank() failed: %v, q=%v",err,query))
		os.Exit(1)
	}
	affected_rows,err:=res.RowsAffected()
	if err!=nil {
		ss.Log_msg(fmt.Sprintf("Error getting RowsAffected in update_top_volume_rank(): %v",err))
	}
	if affected_rows == 0 {
		query = "INSERT INTO uranks(aid,top_volume,volume) VALUES($1,$2,$3)"
		_,err:=ss.db.Exec(query,aid,value,volume)
		if (err!=nil) {
			ss.Log_msg(fmt.Sprintf("update_top_volume_rank() failed: value=%v, err: %v, q=%v",value,err,query))
			os.Exit(1)
		}

	}
	return affected_rows
}
func (ss *SQLStorage) Get_top_profit_makers() []p.ProfitMaker {

	var query string
	query = "SELECT a.addr,r.top_profit,r.profit FROM uranks AS r " +
			"LEFT JOIN address AS a ON r.aid = a.address_id " +
			"ORDER BY r.top_profit ASC,r.profit DESC LIMIT 100"

	rows,err := ss.db.Query(query)
	if (err!=nil) {
		ss.Log_msg(fmt.Sprintf("DB error: %v (query=%v)",err,query))
		os.Exit(1)
	}
	records := make([]p.ProfitMaker,0,101)
	defer rows.Close()
	for rows.Next() {
		var rec p.ProfitMaker
		err=rows.Scan(&rec.Addr,&rec.Percentage,&rec.ProfitLoss)
		records = append(records,rec)
	}
	return records
}
func (ss *SQLStorage) Get_top_trade_makers() []p.TradeMaker {

	var query string
	query = "SELECT a.addr,r.top_trades,r.total_trades FROM uranks AS r " +
			"LEFT JOIN address AS a ON r.aid = a.address_id " +
			"ORDER BY r.top_trades ASC,r.total_trades DESC LIMIT 100"

	rows,err := ss.db.Query(query)
	if (err!=nil) {
		ss.Log_msg(fmt.Sprintf("DB error: %v (query=%v)",err,query))
		os.Exit(1)
	}
	records := make([]p.TradeMaker,0,101)
	defer rows.Close()
	for rows.Next() {
		var rec p.TradeMaker
		err=rows.Scan(&rec.Addr,&rec.Percentage,&rec.TotalTrades)
		records = append(records,rec)
	}
	return records
}
func (ss *SQLStorage) Get_top_volume_makers() []p.VolumeMaker {

	var query string
	query = "SELECT a.addr,r.top_volume,r.volume FROM uranks AS r " +
			"LEFT JOIN address AS a ON r.aid = a.address_id " +
			"ORDER BY r.top_volume ASC,r.volume DESC LIMIT 100"

	rows,err := ss.db.Query(query)
	if (err!=nil) {
		ss.Log_msg(fmt.Sprintf("DB error: %v (query=%v)",err,query))
		os.Exit(1)
	}
	records := make([]p.VolumeMaker,0,101)
	defer rows.Close()
	for rows.Next() {
		var rec p.VolumeMaker
		err=rows.Scan(&rec.Addr,&rec.Percentage,&rec.Volume)
		records = append(records,rec)
	}
	return records
}
func (ss *SQLStorage) Get_user_ranks(sort int,order int) []p.UserRank {

	records := make([]p.UserRank,0,256)
	var query string
	var order_field string
	var order_dir string = "DESC"

	switch (sort) {
	case 0: order_field = "r.profit"
	case 1: order_field = "r.volume"
	case 2: order_field = "r.total_trades"
	case 3: order_field = "s.markets_created"
	default:
		return records
	}
	if order!=0 {
		order_dir="ASC"
	}

	query = "SELECT " +
				"r.aid,a.addr,r.profit,r.total_trades,r.volume,s.markets_created " +
				"FROM uranks AS r " +
					"JOIN  ustats AS s ON r.aid=s.aid " +
			"LEFT JOIN address AS a ON r.aid = a.address_id " +
			"ORDER BY "+order_field+" "+order_dir

	rows,err := ss.db.Query(query)
	if (err!=nil) {
		ss.Log_msg(fmt.Sprintf("DB error: %v (query=%v)",err,query))
		os.Exit(1)
	}
	defer rows.Close()
	for rows.Next() {
		var rec p.UserRank
		err=rows.Scan(
			&rec.Aid,
			&rec.Addr,
			&rec.ProfitLoss,
			&rec.TotalTrades,
			&rec.VolumeTraded,
			&rec.NumMktCreated,
		)
		records = append(records,rec)
	}
	return records
}
func (ss *SQLStorage) Get_user_reports(aid int64,limit int) []p.Report {

	var query string
	/* OLD, removal pending
	query = "SELECT " +
				"m.market_aid,"+
				"r.time_stamp::date," +
				"ma.addr as mkt_addr," +
				"r.is_designated," +
				"r.amount_staked,"+
				"r.outcome_idx," +
				"m.extra_info::json->>'description' AS descr," +
				"m.initial_outcome," +
				"m.designated_outcome," +
				"m.winning_outcome," +
				"m.market_type AS mtype," +
				"m.outcomes AS outcomes_str " +
			"FROM " +
					"initial_report AS r, " +
					"market AS m " +
						"LEFT JOIN address AS ma ON m.market_aid = ma.address_id " +
			"WHERE (r.market_aid = m.market_aid) and (r.reporter_aid=$1) " +
			"ORDER BY r.time_stamp"
	*/
	query = "SELECT " +
				"m.market_aid AS maid,"+
				"EXTRACT(EPOCH FROM r.time_stamp)::BIGINT AS ts," +
				"r.time_stamp::date AS datetime," +
				"ma.addr AS mkt_addr," +
				"r.is_designated," +
				"r.amount_staked,"+
				"r.outcome_idx," +
				"m.extra_info::json->>'description' AS descr," +
				"m.initial_outcome," +
				"m.designated_outcome," +
				"m.winning_outcome," +
				"m.market_type AS mtype," +
				"1 AS dispute_round," +
				"0 AS rtype," +
				"m.outcomes AS outcomes_str " +
			"FROM initial_report AS r " +
				"JOIN market AS m ON r.market_aid=m.market_aid " +
				"JOIN address AS ma ON r.market_aid = ma.address_id " +
			"WHERE r.reporter_aid=$1 " +
			"UNION ALL " +
			"SELECT " +
				"m.market_aid as maid," +
				"EXTRACT(EPOCH FROM c.time_stamp)::BIGINT AS ts," +
				"c.time_stamp::date AS datetime," +
				"ma.addr AS mkt_addr," +
				"FALSE as is_designated," +
				"c.amount_staked," +
				"c.outcome_idx," +
				"m.extra_info::json->>'description' AS descr," +
				"m.initial_outcome," +
				"m.designated_outcome," +
				"m.winning_outcome," +
				"m.market_type AS mtype," +
				"c.dispute_round," +
				"1 AS rtype," +
				"m.outcomes AS outcomes_str " +
			"FROM crowdsourcer_contrib AS c " +
				"JOIN market AS m ON c.market_aid=m.market_aid " +
				"JOIN address AS ma ON c.market_aid = ma.address_id " +
			"WHERE c.reporter_aid=$1 " +
			"ORDER BY ts"
	if limit > 0 {
		query = query +	" LIMIT " + strconv.Itoa(limit)
	}

	records := make([]p.Report,0,8)
	var rows *sql.Rows
	var err error
	rows,err = ss.db.Query(query,aid)
	if (err!=nil) {
		if err == sql.ErrNoRows {
			return records
		}
		ss.Log_msg(fmt.Sprintf("DB error: %v (query=%v)",err,query))
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
		var mkt_descr sql.NullString
		err=rows.Scan(
			&rec.MktAid,
			&rec.TimeStamp,
			&rec.Date,
			&rec.MktAddr,
			&rec.IsDesignated,
			&rec.RepStake,
			&rec.OutcomeIdx,
			&mkt_descr,
			&initial_outcome,
			&designated_outcome,
			&winning_outcome,
			&rec.MktType,
			&rec.Round,
			&rec.RType,
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
		if mkt_descr.Valid {
			rec.MktDescription = mkt_descr.String
		}
		records = append(records,rec)
	}
	return records
}
func (ss *SQLStorage) Get_user_markets(aid int64) []p.InfoMarket {

	var query string
	query = "SELECT " +
				"ma.addr as mkt_addr," +
				"ca.addr as mcreator," +
				"TO_CHAR(end_time,'dd/mm/yyyy HH24:SS UTC') as end_date," + 
				"FLOOR(EXTRACT(EPOCH FROM m.create_timestamp))::BIGINT as created_ts," +
				"FLOOR(EXTRACT(EPOCH FROM m.end_time))::BIGINT as end_ts," + 
				"FLOOR(EXTRACT(EPOCH FROM m.fin_timestamp))::BIGINT as fin_ts," +
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
				"num_ticks,"+
				"fee," +
				"total_trades,"+
				"open_interest AS OI," +
				"cur_volume AS volume " +
			"FROM market as m " +
				"LEFT JOIN address AS ma ON m.market_aid = ma.address_id " +
				"LEFT JOIN address AS ca ON m.creator_aid = ca.address_id " +
			"WHERE aid = $1 " +
			"ORDER BY " +
				"m.market_aid "

	rows,err := ss.db.Query(query,aid)
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
		var categories sql.NullString
		err=rows.Scan(
					&rec.MktAddr,
					&rec.MktCreator,
					&rec.EndDate,
					&rec.CreatedTs,
					&rec.EndTs,
					&rec.FinTs,
					&description,
					&longdesc,
					&categories,
					&rec.Outcomes,
					&rec.MktType,
					&rec.MktTypeStr,
					&rec.MktStatus,
					&rec.NumTicks,
					&rec.Fee,
					&rec.TotalTrades,
					&rec.OpenInterest,
					&rec.CurVolume,
		)
		if err!=nil {
			ss.Log_msg(fmt.Sprintf("DB error: %v, q=%v",err,query))
			os.Exit(1)
		}
		fmt.Printf("type=%v, typestr=%v, addr=%v\n",rec.MktType,rec.MktTypeStr,rec.MktAddr)
		if description.Valid {
			rec.Description = description.String
		}
		if longdesc.Valid {
			rec.LongDesc = longdesc.String
		}
		if categories.Valid {
			rec.CategoryStr = categories.String
		}
		rec.MktAddrSh=p.Short_address(rec.MktAddr)
		rec.MktCreatorSh=p.Short_address(rec.MktCreator)
		rec.Status=get_market_status_str(p.MarketStatus(rec.MktStatus))
		records = append(records,rec)
	}
	return records
}
func (ss *SQLStorage) Get_unclaimed_profit(aid int64) float64 {

	var query string
	query = "SELECT sum(final_profit) AS pl FROM claim_funds WHERE aid=$1 AND claim_status=1"
	row := ss.db.QueryRow(query,aid)

	var profit sql.NullFloat64
	var err error
	err=row.Scan(&profit);
	if (err!=nil) {
		if err == sql.ErrNoRows {
			return 0.0
		} else {
			ss.Log_msg(fmt.Sprintf("Error in Get_unclaimed_amount(): %v, q=%v",err,query))
			os.Exit(1)
		}
	}
	if profit.Valid {
		return profit.Float64
	} else {
		return 0.0
	}
}
func (ss *SQLStorage) Get_mkt_participant_outcomes(mkt_addr *common.Address) []*p.PosChg {

	output := make([]*p.PosChg,0,16)
	market_aid := ss.Lookup_address_id(mkt_addr.String())

	var query string
	query = "SELECT wa.addr,pl.outcome_idx FROM profit_loss AS pl " +
				"LEFT JOIN address AS a ON pl.aid=a.address_id " +
				"WHERE pl.market_aid=$1 "
	rows,err := ss.db.Query(query,market_aid)
	if (err!=nil) {
		ss.Log_msg(fmt.Sprintf("DB error: %v (query=%v)",err,query))
		os.Exit(1)
	}
	defer rows.Close()
	for rows.Next() {
		var (
			addr string
			outcome_idx int
		)
		err=rows.Scan(&addr,&outcome_idx)
		if err!=nil {
			ss.Log_msg(fmt.Sprintf("Scan error at Get_mkt_participant_outcomes: %v \n",err))
			os.Exit(1)
		}
		pchg := new(p.PosChg)
		pchg.Mkt_addr = *mkt_addr
		pchg.Addr = common.HexToAddress(addr)
		pchg.Outcome = new(big.Int)
		pchg.Outcome.SetInt64(int64(outcome_idx))
		output = append(output,pchg)
	}
	return output
}
func (ss *SQLStorage) Get_traded_markets_for_user(aid int64,active_flag int) []p.InfoMarket {

	var where_condition string
	if active_flag == 1 {
		where_condition = "(m.status < 4) "
	} else {
		where_condition = "(m.status > 3) "
	}
	var query string
	query = "SELECT " +
				"m.market_aid," +
				"ma.addr as mkt_addr," +
				"ca.addr as mcreator," +
				"TO_CHAR(end_time,'dd/mm/yyyy HH24:SS UTC') as end_date," + 
				"FLOOR(EXTRACT(EPOCH FROM m.create_timestamp))::BIGINT as created_ts," +
				"FLOOR(EXTRACT(EPOCH FROM m.end_time))::BIGINT as end_ts," + 
				"FLOOR(EXTRACT(EPOCH FROM m.fin_timestamp))::BIGINT as fin_ts," +
				"extra_info::json->>'description' AS descr," +
				"extra_info::json->>'longDescription' AS long_desc," +
				"extra_info::json->>'categories' AS categories," +
				"outcomes," +
				"m.winning_outcome," +
				"m.market_type, " +
				"CASE m.market_type " +
					"WHEN 0 THEN 'YES/NO' " +
					"WHEN 1 THEN 'CATEGORICAL' " +
					"WHEN 2 THEN 'SCALAR' " +
				"END AS mtype," +
				"status,"+
				"CASE WHEN EXTRACT(epoch from (fin_timestamp-now())) < 0 " +
					"THEN 'Trading' ELSE 'Reporting' END AS status_desc," +
				"num_ticks,"+
				"fee," +
				"open_interest AS OI," +
				"cur_volume AS volume, " +
				"s.volume_traded, " +
				"s.total_trades " +
			"FROM market as m " +
				"JOIN trd_mkt_stats AS s ON m.market_aid = s.market_aid " +
				"LEFT JOIN address AS ma ON m.market_aid = ma.address_id " +
				"LEFT JOIN address AS ca ON m.creator_aid = ca.address_id " +
			"WHERE s.aid = $1 AND " + where_condition +
			"ORDER BY s.volume_traded DESC"

	rows,err := ss.db.Query(query,aid)
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
					&rec.MktCreator,
					&rec.EndDate,
					&rec.CreatedTs,
					&rec.EndTs,
					&rec.FinTs,
					&description,
					&longdesc,
					&category,
					&rec.Outcomes,
					&rec.WinOutcomeIdx,
					&rec.MktType,
					&rec.MktTypeStr,
					&rec.MktStatus,
					&rec.Status,
					&rec.NumTicks,
					&rec.Fee,
					&rec.OpenInterest,
					&rec.CurVolume,
					&rec.VolTraded,
					&rec.TotalTrades,
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
		if category.Valid {
			rec.CategoryStr = category.String
		}
		rec.Status=get_market_status_str(p.MarketStatus(rec.MktStatus))
		rec.MktAddrSh=p.Short_address(rec.MktAddr)
		rec.MktCreatorSh=p.Short_address(rec.MktCreator)
		records = append(records,rec)
	}
	return records
}
func (ss *SQLStorage) Get_created_markets_for_user(aid int64) []p.InfoMarket {

	var query string
	query = "SELECT " +
				"m.market_aid," +
				"ma.addr as mkt_addr," +
				"ca.addr as mcreator," +
				"TO_CHAR(end_time,'dd/mm/yyyy HH24:SS UTC') as end_date," + 
				"FLOOR(EXTRACT(EPOCH FROM m.create_timestamp))::BIGINT as created_ts," +
				"FLOOR(EXTRACT(EPOCH FROM m.end_time))::BIGINT as end_ts," + 
				"FLOOR(EXTRACT(EPOCH FROM m.fin_timestamp))::BIGINT as fin_ts," +
				"extra_info::json->>'description' AS descr," +
				"extra_info::json->>'longDescription' AS long_desc," +
				"extra_info::json->>'categories' AS categories," +
				"outcomes," +
				"m.winning_outcome," +
				"m.market_type, " +
				"CASE m.market_type " +
					"WHEN 0 THEN 'YES/NO' " +
					"WHEN 1 THEN 'CATEGORICAL' " +
					"WHEN 2 THEN 'SCALAR' " +
				"END AS mtype," +
				"status,"+
				"CASE WHEN EXTRACT(epoch from (fin_timestamp-now())) < 0 " +
					"THEN 'Trading' ELSE 'Reporting' END AS status_desc," +
				"num_ticks,"+
				"fee," +
				"open_interest AS OI," +
				"cur_volume," +
				"money_at_stake " +
			"FROM market as m " +
				"LEFT JOIN address AS ma ON m.market_aid = ma.address_id " +
				"LEFT JOIN address AS ca ON m.creator_aid = ca.address_id " +
			"WHERE m.creator_aid=$1 " +
			"ORDER BY m.create_timestamp"

	rows,err := ss.db.Query(query,aid)
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
					&rec.MktCreator,
					&rec.EndDate,
					&rec.CreatedTs,
					&rec.EndTs,
					&rec.FinTs,
					&description,
					&longdesc,
					&category,
					&rec.Outcomes,
					&rec.WinOutcomeIdx,
					&rec.MktType,
					&rec.MktTypeStr,
					&rec.MktStatus,
					&rec.Status,
					&rec.NumTicks,
					&rec.Fee,
					&rec.OpenInterest,
					&rec.CurVolume,
					&rec.MoneyAtStake,
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
		if category.Valid {
			rec.CategoryStr = category.String
		}
		rec.Status=get_market_status_str(p.MarketStatus(rec.MktStatus))
		rec.MktAddrSh=p.Short_address(rec.MktAddr)
		rec.MktCreatorSh=p.Short_address(rec.MktCreator)
		records = append(records,rec)
	}
	return records
}
func (ss *SQLStorage) Get_user_trades_for_market(aid int64,mkt_aid int64) []p.MarketTrade {
	// get market trades with mixed outcomes
	var query string
	query = "SELECT " +
				"o.id," +
				"o.order_hash," +
				"a.addr as mkt_addr," +
				"ca.addr as creator_addr," +
				"fa.addr as filler_addr," +
				"CASE oaction " +
					"WHEN 0 THEN 'CREATE' " +
					"WHEN 1 THEN 'CANCEL' " +
					"WHEN 2 THEN 'FILL' " +
				"END AS type, " +
				"CASE o.otype " +
					"WHEN 0 THEN 'BID' " +
					"ELSE 'ASK' " +
				"END AS dir, " +
				"o.time_stamp::text AS date," +
				"o.price, " +
				"o.amount_filled AS amount," +
				"o.outcome_idx," +
				"m.market_type AS mtype," +
				"m.outcomes AS outcomes_str " +
			"FROM mktord AS o " +
				"JOIN market AS m ON o.market_aid=m.market_aid " +
				"LEFT JOIN address AS a ON o.market_aid=a.address_id " +
				"LEFT JOIN address AS fa ON o.fill_aid=fa.address_id " +
				"LEFT JOIN address AS ca ON o.aid=ca.address_id " +
			"WHERE o.market_aid = $1 AND ((o.aid=$2) OR (o.fill_aid=$2)) " +
			"ORDER BY o.block_num DESC,o.time_stamp DESC"

	ss.Info.Printf("aid=%v, market_aid=%v, query=%v\n",aid,mkt_aid,query)
	var rows *sql.Rows
	var err error
	rows,err = ss.db.Query(query,mkt_aid,aid)
	if (err!=nil) {
		ss.Log_msg(fmt.Sprintf("DB error: %v (query=%v)",err,query))
		os.Exit(1)
	}
	records := make([]p.MarketTrade,0,8)

	defer rows.Close()
	for rows.Next() {
		var rec p.MarketTrade
		var mkt_type,decimals int
		var outcomes string
		err=rows.Scan(
			&rec.OrderId,
			&rec.OrderHash,
			&rec.MktAddr,
			&rec.CreatorAddr,
			&rec.FillerAddr,
			&rec.Type,
			&rec.Direction,
			&rec.Date,
			&rec.Price,
			&rec.Amount,
			&rec.Outcome,
			&mkt_type,
			&decimals,
			&outcomes,
		)
		if err!=nil {
			ss.Log_msg(fmt.Sprintf("DB error: %v, q=%v",err,query))
			os.Exit(1)
		}
		p.Augur_UI_price_adjustments(&rec.Price,&rec.Amount,mkt_type,decimals)
		rec.MktAddrSh=p.Short_address(rec.MktAddr)
		rec.CreatorAddrSh=p.Short_address(rec.CreatorAddr)
		rec.FillerAddrSh=p.Short_address(rec.FillerAddr)
		rec.OutcomeStr = get_outcome_str(uint8(mkt_type),rec.Outcome,&outcomes)
		records = append(records,rec)
		ss.Log_msg(fmt.Sprintf("Record appended %+v\n",rec))
	}
	return records
}
func (ss *SQLStorage) Get_user_open_orders(user_aid int64) []p.OpenOrder {

	records := make([]p.OpenOrder,0,8)
	// open orders on 0x Mesh network
	var query string
	query = "SELECT " +
				"ca.addr AS creator_addr," +
				"ma.addr AS mkt_addr," +
				"m.market_type," +
				"m.decimals," +
				"CASE m.market_type " +
					"WHEN 0 THEN 'YES/NO' " +
					"WHEN 1 THEN 'CATEGORICAL' " +
					"WHEN 2 THEN 'SCALAR' " +
				"END AS market_type_str," +
				"m.status,"+
				"FLOOR(EXTRACT(EPOCH FROM m.end_time))::BIGINT AS ts," +
				"TO_CHAR(m.end_time,'dd/mm/yyyy HH24:SS UTC') as order_date," + 
				"m.extra_info::json->>'description' AS descr," +
				"m.outcomes," +
				"o.id," +
				"o.otype," +
				"CASE o.otype " +
					"WHEN 0 THEN 'BID' " +
					"ELSE 'ASK' " +
				"END AS dir, " +
				"o.outcome_idx," +
				"ROUND(o.price,3), "+
				"o.amount," +
				"FLOOR(EXTRACT(EPOCH FROM o.evt_timestamp))::BIGINT AS ts," +
				"o.order_hash " +
			"FROM oorders AS o " +
				"LEFT JOIN market AS m ON o.market_aid = m.market_aid " +
				"LEFT JOIN address AS ma ON o.market_aid = ma.address_id " +
				"LEFT JOIN address AS ca ON o.aid = ca.address_id " +
			"WHERE o.aid = $1 " +
			"ORDER BY o.id DESC"

	rows,err := ss.db.Query(query,user_aid)
	if err!=nil {
		ss.Log_msg(fmt.Sprintf("DB error: %v, q=%v",err,query))
		os.Exit(1)
	}
	defer rows.Close()
	for rows.Next() {
		var rec p.OpenOrder
		var outcomes string
		var decimals int
		var descr sql.NullString
		err=rows.Scan(
			&rec.CreatorAddr,
			&rec.MktAddr,
			&rec.MktType,
			&decimals,
			&rec.MktTypeStr,
			&rec.MktStatus,
			&rec.MktExpirationTs,
			&rec.OrderDate,
			&descr,
			&outcomes,
			&rec.Id,
			&rec.OrderType,
			&rec.Direction,
			&rec.Outcome,
			&rec.Price,
			&rec.Amount,
			&rec.Timestamp,
			&rec.OrderHash,
		)
		if err!=nil {
			ss.Log_msg(fmt.Sprintf("DB error: %v, q=%v",err,query))
			os.Exit(1)
		}
		p.Augur_UI_price_adjustments(&rec.Price,&rec.Amount,rec.MktType,decimals)
		rec.MktStatusStr = get_market_status_str(p.MarketStatus(rec.MktStatus))
		rec.MktAddrSh=p.Short_address(rec.MktAddr)
		rec.CreatorAddrSh=p.Short_address(rec.CreatorAddr)
		rec.OutcomeStr = get_outcome_str(uint8(rec.MktType),rec.Outcome,&outcomes)
		records = append(records,rec)
	}
	return records
}
func  (ss *SQLStorage) Get_gas_spent_for_user(aid int64) (p.GasSpent,error) {

	var output p.GasSpent
	var query string
	query =
		"SELECT " +
			"gtrading,greporting,gmarkets," +
			"geth_trading,geth_reporting,geth_markets "+
		"FROM ustats "+
		"WHERE aid=$1"

	row := ss.db.QueryRow(query,aid)
	err := row.Scan(
		&output.Trading,
		&output.Reporting,
		&output.Markets,
		&output.EthTrading,
		&output.EthReporting,
		&output.EthMarkets,
	)
	if err != nil {
		if err!=sql.ErrNoRows {
			ss.Log_msg(fmt.Sprintf("DB error: %v, q=%v\n",err,query))
			os.Exit(1)
		}
	}
	return output,nil
}
func (ss *SQLStorage) Get_user_oo_history(aid int64) []p.OpenOrder {

	records := make([]p.OpenOrder,0,8)
	// open orders on 0x Mesh network
	var query string
	query = "SELECT " +
				"o.mktord_id," +
				"ca.addr AS creator_addr," +
				"ma.addr AS mkt_addr," +
				"m.market_type," +
				"CASE m.market_type " +
					"WHEN 0 THEN 'YES/NO' " +
					"WHEN 1 THEN 'CATEGORICAL' " +
					"WHEN 2 THEN 'SCALAR' " +
				"END AS market_type_str," +
				"m.status,"+
				"FLOOR(EXTRACT(EPOCH FROM m.end_time))::BIGINT AS mkt_end_ts," +
				"FLOOR(EXTRACT(EPOCH FROM o.expiration))::BIGINT AS order_exp_ts," +
				"TO_CHAR(o.srv_timestamp,'dd/mm/yyyy HH24:SS UTC') as order_date," + 
				"m.extra_info::json->>'description' AS descr," +
				"m.outcomes," +
				"o.id," +
				"o.otype," +
				"o.opcode," +
				"CASE o.otype " +
					"WHEN 0 THEN 'BID' " +
					"ELSE 'ASK' " +
				"END AS dir, " +
				"o.outcome_idx," +
				"ROUND(o.price,3), "+
				"o.initial_amount,"+
				"o.amount," +
				"FLOOR(EXTRACT(EPOCH FROM o.evt_timestamp))::BIGINT AS ts," +
				"o.order_hash " +
			"FROM oohist AS o " +
				"LEFT JOIN market AS m ON o.market_aid = m.market_aid " +
				"LEFT JOIN address AS ma ON o.market_aid = ma.address_id " +
				"LEFT JOIN address AS ca ON o.aid = ca.address_id " +
			"WHERE o.aid = $1 " +
			"ORDER BY o.id DESC"

	rows,err := ss.db.Query(query,aid)
	if err!=nil {
		ss.Log_msg(fmt.Sprintf("DB error: %v, q=%v",err,query))
		os.Exit(1)
	}
	defer rows.Close()
	for rows.Next() {
		var rec p.OpenOrder
		var outcomes string
		var decimals int
		var descr sql.NullString
		err=rows.Scan(
			&rec.MktOrderId,
			&rec.CreatorAddr,
			&rec.MktAddr,
			&rec.MktType,
			&decimals,
			&rec.MktTypeStr,
			&rec.MktStatus,
			&rec.MktExpirationTs,
			&rec.OrderExpirationTs,
			&rec.OrderDate,
			&descr,
			&outcomes,
			&rec.Id,
			&rec.OrderType,
			&rec.OpCode,
			&rec.Direction,
			&rec.Outcome,
			&rec.Price,
			&rec.InitialAmount,
			&rec.Amount,
			&rec.Timestamp,
			&rec.OrderHash,
		)
		if err!=nil {
			ss.Log_msg(fmt.Sprintf("DB error: %v, q=%v",err,query))
			os.Exit(1)
		}
		p.Augur_UI_price_adjustments(&rec.Price,&rec.Amount,rec.MktType,decimals)
		p.Augur_UI_price_adjustments(nil,&rec.InitialAmount,rec.MktType,decimals)
		rec.MktStatusStr = get_market_status_str(p.MarketStatus(rec.MktStatus))
		rec.MktAddrSh=p.Short_address(rec.MktAddr)
		rec.CreatorAddrSh=p.Short_address(rec.CreatorAddr)
		rec.OrderHashSh=p.Short_hash(rec.OrderHash)
		rec.OutcomeStr = get_outcome_str(uint8(rec.MktType),rec.Outcome,&outcomes)
		records = append(records,rec)
	}
	return records
}
func (ss *SQLStorage) Get_wrapped_shtoken_balances(aid int64) []p.UserShTokens {

	var query string
	query = "WITH idrecs AS (" +
				"SELECT MAX(id) AS id,count(*) AS num_recs,wrapper_aid " +
				"FROM wstok_bal WHERE aid=$1 GROUP BY wrapper_aid" +
			") " +
			"SELECT " +
				"balance + amount AS balance," +
				"num_recs, " +
				"wa.addr," +
				"inf.symbol,"+
				"inf.name, " +
				"w.outcome_idx," +
				"ma.addr " +
			"FROM idrecs AS i " +
				"JOIN wstok_bal AS b ON i.id=b.id " +
				"JOIN af_wrapper AS w ON i.wrapper_aid=w.wrapper_aid " +
				"JOIN erc20_info AS inf ON inf.aid=i.wrapper_aid " +
				"JOIN address AS wa ON i.wrapper_aid=wa.address_id " +
				"JOIN address AS ma ON w.market_aid = ma.address_id " +
			"ORDER BY num_recs DESC "
	rows,err := ss.db.Query(query,aid)
	if (err!=nil) {
		ss.Log_msg(fmt.Sprintf("DB error: %v (query=%v)",err,query))
		os.Exit(1)
	}
	records := make([]p.UserShTokens,0,8)
	defer rows.Close()
	for rows.Next() {
		var rec p.UserShTokens
		err=rows.Scan(
			&rec.Balance,
			&rec.NumTransfers,
			&rec.WrapperAddr,
			&rec.Symbol,
			&rec.Name,
			&rec.OutcomeIdx,
			&rec.MarketAddr,
		)
		records = append(records,rec)
	}
	return records
}
func (ss *SQLStorage) Get_user_wrapped_shtoken_transfers(aid,wrapper_aid int64,offset,limit int) (int64,[]p.UserShtokTransfer) {

	var query string

	var total_rows int64
	query = "SELECT COUNT(*) AS total_rows FROM wstok_bal b " +
			"WHERE b.aid=$1 AND b.wrapper_aid=$2"

	var null_counter sql.NullInt64
	var err error
	err=ss.db.QueryRow(query,aid,wrapper_aid).Scan(&null_counter);
	if (err!=nil) {
		if err != sql.ErrNoRows {
			ss.Log_msg(fmt.Sprintf("DB error: %v , q=%v",err,query))
			os.Exit(1)
		}
	}
	total_rows=null_counter.Int64

	query = "SELECT " +
				"fa.addr," +
				"ta.addr," +
				"EXTRACT(EPOCH FROM tr.time_stamp)::BIGINT AS ts," +
				"tr.time_stamp," +
				"tr.from_aid," +
				"tr.to_aid," +
				"tr.amount, " +
				"tb.balance " +
			"FROM wstok_bal AS tb " +
			"JOIN wstok_transf AS tr ON tb.tr_id=tr.id " +
			"LEFT JOIN address AS fa ON tr.from_aid=fa.address_id " +
			"LEFT JOIN address AS ta ON tr.to_aid=ta.address_id " +
			"WHERE tb.aid=$1 AND tb.wrapper_aid=$2 " +
			"ORDER BY tr.time_stamp DESC,tb.id DESC " +
			"OFFSET $3 LIMIT $4"
	
	rows,err := ss.db.Query(query,aid,wrapper_aid,offset,limit)
	if (err!=nil) {
		ss.Log_msg(fmt.Sprintf("DB error: %v (query=%v)",err,query))
		os.Exit(1)
	}
	records := make([]p.UserShtokTransfer,0,8)
	defer rows.Close()
	for rows.Next() {
		var rec p.UserShtokTransfer
		err=rows.Scan(
			&rec.From,
			&rec.To,
			&rec.TimeStamp,
			&rec.Date,
			&rec.FromAid,
			&rec.ToAid,
			&rec.Amount,
			&rec.Balance,
		)
		records = append(records,rec)
	}
	return total_rows,records
}
func (ss *SQLStorage) Get_user_uniswap_swaps(user_aid int64,offset int,limit int) ([]p.UserUniswapSwap , int64) {

	records := make([]p.UserUniswapSwap,0,128)
	var query string

	query = "SELECT COUNT(*) AS total_recs FROM uswap1 sw " +
			"JOIN upair AS p ON sw.pair_aid = p.pair_aid " +
			"JOIN af_wrapper AS w ON (w.wrapper_aid=p.token0_aid OR w.wrapper_aid=p.token1_aid) " +
			"WHERE sw.recipient_aid=$1"
	var total_rows int64 = 0
	var null_recs sql.NullInt64
	err := ss.db.QueryRow(query,user_aid).Scan(&null_recs)
	if err != nil {
		if err != sql.ErrNoRows {
			ss.Log_msg(fmt.Sprintf("DB error: %v (query=%v)",err,query))
			os.Exit(1)
		}
	}
	total_rows = null_recs.Int64

	query = "SELECT " +
				"sw.id,"+
				"sw.block_num," +
				"sw.amount0_in, " +
				"sw.amount1_in," +
				"sw.amount0_out," +
				"sw.amount1_out," +
				"sw.time_stamp," +
				"EXTRACT(EPOCH FROM sw.time_stamp)::BIGINT AS created_ts, "+
				"ra.addr AS recipient_addr," +
				"sw.recipient_aid, " +
				"p.pair_aid, " +
				"pa.addr, " +
				"inf1.symbol," +
				"inf1.name," +
				"inf2.symbol," +
				"inf2.name, " +
				"w.outcome_idx, " +
				"m.market_aid,"+
				"ma.addr," +
				"m.market_type, " +
				"m.extra_info::json->>'description' AS descr," +
				"m.outcomes " +
			"FROM uswap1 AS sw " +
			"JOIN upair AS p ON sw.pair_aid = p.pair_aid " +
			"JOIN af_wrapper AS w ON (w.wrapper_aid=p.token0_aid OR w.wrapper_aid=p.token1_aid) " +
			"JOIN market AS m ON w.market_aid=m.market_aid " +
			"LEFT JOIN address AS ra ON sw.recipient_aid=ra.address_id " +
			"LEFT JOIN address AS pa ON sw.pair_aid=pa.address_id " +
			"LEFT JOIN address AS ma ON m.market_aid=ma.address_id " +
			"LEFT JOIN erc20_info inf1 ON p.token0_aid=inf1.aid " +
			"LEFT JOIN erc20_info inf2 ON p.token1_aid=inf2.aid " +
			"WHERE sw.recipient_aid=$1 " +
			"ORDER BY sw.time_stamp DESC "+
			"OFFSET $2 LIMIT $3"

	rows,err := ss.db.Query(query,user_aid,offset,limit)
	if (err!=nil) {
		ss.Log_msg(fmt.Sprintf("DB error: %v (query=%v)",err,query))
		os.Exit(1)
	}
	defer rows.Close()
	for rows.Next() {
		var rec p.UserUniswapSwap
		var mkt_type int
		var null_desc sql.NullString
		var outcomes string
		err=rows.Scan(
			&rec.Id,
			&rec.BlockNum,
			&rec.Amount0_In,
			&rec.Amount1_In,
			&rec.Amount0_Out,
			&rec.Amount1_Out,
			&rec.CreatedDate,
			&rec.CreatedTs,
			&rec.RequesterAddr,
			&rec.RequesterAid,
			&rec.PairAid,
			&rec.PairAddr,
			&rec.Symbol0,
			&rec.Name0,
			&rec.Symbol1,
			&rec.Name1,
			&rec.OutcomeIdx,
			&rec.MktAid,
			&rec.MktAddr,
			&mkt_type,
			&null_desc,
			&outcomes,
		)
		if err!=nil {
			ss.Log_msg(fmt.Sprintf("DB error: %v, q=%v",err,query))
			os.Exit(1)
		}
		if null_desc.Valid {
			rec.MktDescription = null_desc.String
		}
		rec.Outcome = get_outcome_str(uint8(mkt_type),rec.OutcomeIdx,&outcomes)
		records = append(records,rec)
	}
	return records,total_rows
}
func (ss *SQLStorage) Get_user_balancer_swaps(user_aid int64,offset int,limit int) ([]p.UserBalancerSwap,int64) {

	records := make([]p.UserBalancerSwap,0,64)
	var query string
	query = "SELECT count(*) as total " +
			"FROM bswap AS s " +
				"JOIN af_wrapper AS w ON (w.wrapper_aid=s.token_in_aid OR w.wrapper_aid=s.token_out_aid) "+
			"WHERE s.caller_aid=$1"
	var total_rows int64 = 0
	var null_recs sql.NullInt64
	err := ss.db.QueryRow(query,user_aid).Scan(&null_recs)
	if err != nil {
		if err != sql.ErrNoRows {
			ss.Log_msg(fmt.Sprintf("DB error: %v (query=%v)",err,query))
			os.Exit(1)
		}
	}
	total_rows = null_recs.Int64

	query = "SELECT " +
				"s.id,"+
				"EXTRACT(EPOCH FROM s.time_stamp)::BIGINT AS ts, " +
				"s.time_stamp as datetime,"+
				"s.block_num," +
				"s.tx_id,"+
				"ca.addr,"+
				"tia.addr," +
				"toa.addr," +
				"s.token_in_aid," +
				"s.token_out_aid," +
				"e_in.symbol,"+
				"e_out.symbol," +
				"e_in.name," +
				"e_out.name," +
				"s.amount_in, " +
				"s.amount_out, " +
				"s.pool_aid," +
				"pa.addr," +
				"w.outcome_idx, "+
				"m.market_aid," +
				"ma.addr," +
				"m.market_type, " +
				"m.extra_info::json->>'description' AS descr," +
				"m.outcomes " +
			"FROM bswap AS s " +
				"JOIN af_wrapper AS w ON (w.wrapper_aid=s.token_in_aid OR w.wrapper_aid=s.token_out_aid) "+
				"JOIN market AS m ON w.market_aid=m.market_aid " +
				"LEFT JOIN address AS ca ON s.caller_aid=ca.address_id " +
				"LEFT JOIN address AS tia ON s.token_in_aid=tia.address_id " +
				"LEFT JOIN address AS toa ON s.token_out_aid=toa.address_id " +
				"LEFT JOIN address AS ma ON m.market_aid=ma.address_id " +
				"LEFT JOIN address AS pa ON s.pool_aid=pa.address_id "+
				"LEFT JOIN erc20_info AS e_in ON s.token_in_aid=e_in.aid " +
				"LEFT JOIN erc20_info AS e_out ON s.token_out_aid=e_out.aid " +
			"WHERE s.caller_aid=$1 " +
			"ORDER BY ts DESC OFFSET $2 LIMIT $3"
	rows,err := ss.db.Query(query,user_aid,offset,limit)
	if (err!=nil) {
		ss.Log_msg(fmt.Sprintf("DB error: %v (query=%v)",err,query))
		os.Exit(1)
	}

	defer rows.Close()
	for rows.Next() {
		var rec p.UserBalancerSwap
		var symbol_in,symbol_out,name_in,name_out sql.NullString
		var mkt_type int
		var null_desc sql.NullString
		var outcomes string
		err=rows.Scan(
			&rec.Id,
			&rec.TimeStamp,
			&rec.Date,
			&rec.BlockNum,
			&rec.TxId,
			&rec.CallerAddr,
			&rec.TokenInAddr,
			&rec.TokenOutAddr,
			&rec.TokenInAid,
			&rec.TokenOutAid,
			&symbol_in,
			&symbol_out,
			&name_in,
			&name_out,
			&rec.AmountInF,
			&rec.AmountOutF,
			&rec.PoolAid,
			&rec.PoolAddr,
			&rec.OutcomeIdx,
			&rec.MktAid,
			&rec.MktAddr,
			&mkt_type,
			&null_desc,
			&outcomes,
		)
		if err!=nil {
			ss.Log_msg(fmt.Sprintf("DB error: %v, q=%v",err,query))
			os.Exit(1)
		}
		rec.CallerAid = user_aid
		rec.SymbolIn = symbol_in.String
		rec.SymbolOut = symbol_out.String
		rec.NameIn = name_in.String
		rec.NameOut = name_out.String
		if null_desc.Valid {
			rec.MktDescription = null_desc.String
		}
		rec.Outcome = get_outcome_str(uint8(mkt_type),rec.OutcomeIdx,&outcomes)
		records = append(records,rec)
	}
	return records,total_rows
}
func (ss *SQLStorage) Get_user_ens_names_active(user_aid int64,offset int,limit int) ([]p.UserENS,int64) {

	records := make([]p.UserENS,0,4)

	query_last_owner :=
		"SELECT DISTINCT ON (fqdn) fqdn,id,EXTRACT(EPOCH FROM o.time_stamp)::BIGINT as ts,o.time_stamp " +
			"FROM ens_new_owner AS o " +
			"WHERE o.owner_aid=$1 " +
			"ORDER BY o.fqdn,o.time_stamp DESC"

	var query string
	var total_rows int64 = 0
	var null_recs sql.NullInt64
	query = "SELECT COUNT(*) FROM ("+ query_last_owner + ") AS names"
	err := ss.db.QueryRow(query,user_aid).Scan(&null_recs)
	if err != nil {
		if err != sql.ErrNoRows {
			ss.Log_msg(fmt.Sprintf("DB error: %v (query=%v)",err,query))
			os.Exit(1)
		}
	}
	total_rows = null_recs.Int64

	query  = 
		"WITH " +
			"last_owner_user AS (" + query_last_owner + "), " + // this is User's ownership
			"ens_names_to_check AS (SELECT DISTINCT ON (fqdn) fqdn,id FROM ens_new_owner WHERE owner_aid=$1)," +
			"last_name_owner AS (" + // this is the last owner of the name which might not be the User's
				"SELECT " +
					"DISTINCT ON (o.fqdn) o.fqdn, "+
					"o.id," +
					"EXTRACT(EPOCH FROM o.time_stamp)::BIGINT as ts," +
					"o.time_stamp " +
				"FROM ens_new_owner AS o " +
					"JOIN ens_names_to_check AS chk ON chk.fqdn=o.fqdn " +
					// Note: here 'WHERE' is omitted and we will get all the owners for this name list
			") " +
		"SELECT " +
				"EXTRACT(EPOCH FROM o.time_stamp)::BIGINT AS ts," +
				"o.time_stamp," +
				"o.fqdn," +
				"n.fqdn_words, " +
				"txt.num_keys " +
			"FROM ens_new_owner AS o " +
				"JOIN ens_node n ON o.fqdn=n.fqdn " +
				"JOIN last_owner_user lown ON lown.fqdn=o.fqdn " +
				"JOIN last_name_owner nown ON nown.fqdn=o.fqdn " +
				"LEFT JOIN ens_text AS txt ON n.fqdn=txt.node " +
			"WHERE o.id=lown.id AND o.id=nown.id AND lown.id=nown.id " +
			"ORDER BY ts DESC OFFSET $2 LIMIT $3"

	ss.Info.Printf("getens query (owner_aid=%v) : %v\n",user_aid,query)

	rows,err := ss.db.Query(query,user_aid,offset,limit)
	if (err!=nil) {
		ss.Log_msg(fmt.Sprintf("DB error: %v (query=%v)",err,query))
		os.Exit(1)
	}
	defer rows.Close()
	for rows.Next() {
		var rec p.UserENS
		var null_num_keys sql.NullInt64
		var null_ens_name sql.NullString
		err := rows.Scan(
			&rec.TsNameAcquired,
			&rec.DateNameAcquired,
			&rec.NodeHash,
			&null_ens_name,
			&null_num_keys,
		)
		if err!=nil {
			ss.Log_msg(fmt.Sprintf("DB error: %v, q=%v",err,query))
			os.Exit(1)
		}
		ss.Info.Printf("fqdn = %v valid=%v, ensname=%v\n",rec.NodeHash,null_ens_name.Valid,null_ens_name.String)
		if null_ens_name.Valid {
			rec.ENS_Name = null_ens_name.String
		}
		if len(rec.ENS_Name) == 0 {
			rec.ENS_Name = "ENS Name is not public"
		}
		rec.NumTextKeyValuePairs = null_num_keys.Int64
		node_addresses := ss.Get_ens_node_addresses(rec.NodeHash)
		if len(node_addresses) > 0 {
			rec.CurAddr = node_addresses[0].Address
			rec.CurAddrAid = node_addresses[0].Aid
		}
		rec.NodeAddressHistory = node_addresses
		records = append(records,rec)
	}
	return records,total_rows
}
func (ss *SQLStorage) Get_user_ens_names_history(user_aid int64,offset int,limit int) ([]p.UserENS,int64) {

	records := make([]p.UserENS,0,4)
	var query string
	query = "WITH last_owner AS (" +
				"SELECT DISTINCT fqdn " +
					"FROM ens_new_owner " +
					"WHERE owner_aid=$1 " +
			") " +
			"SELECT " +
				"count(n.fqdn) "+
			"FROM last_owner AS last " +
				"JOIN ens_node AS n ON last.fqdn=n.fqdn "
//			"WHERE " +
//				"LENGTH(n.fqdn_words) > 0"
	var total_rows int64 = 0
	var null_recs sql.NullInt64
	err := ss.db.QueryRow(query,user_aid).Scan(&null_recs)
	if err != nil {
		if err != sql.ErrNoRows {
			ss.Log_msg(fmt.Sprintf("DB error: %v (query=%v)",err,query))
			os.Exit(1)
		}
	}
	total_rows = null_recs.Int64

	query  = "SELECT " +
				"DISTINCT ON (fqdn) " +
				"EXTRACT(EPOCH FROM o.time_stamp)::BIGINT AS ts," +
				"o.time_stamp," +
				"o.fqdn," +
				"n.fqdn_words, " +
				"txt.num_keys " +
			"FROM ens_new_owner AS o " +
				"JOIN ens_node n ON o.fqdn=n.fqdn " +
				"LEFT JOIN ens_text AS txt ON n.fqdn=txt.node " +
			"WHERE o.owner_aid=$1 " +
			"ORDER BY fqdn,ts DESC OFFSET $2 LIMIT $3"

	ss.Info.Printf("getens query (owner_aid=%v) : %v\n",user_aid,query)

	rows,err := ss.db.Query(query,user_aid,offset,limit)
	if (err!=nil) {
		ss.Log_msg(fmt.Sprintf("DB error: %v (query=%v)",err,query))
		os.Exit(1)
	}
	defer rows.Close()
	for rows.Next() {
		var rec p.UserENS
		var null_num_keys sql.NullInt64
		var null_ens_name sql.NullString
		err := rows.Scan(
			&rec.TsNameAcquired,
			&rec.DateNameAcquired,
			&rec.NodeHash,
			&null_ens_name,
			&null_num_keys,
		)
		if err!=nil {
			ss.Log_msg(fmt.Sprintf("DB error: %v, q=%v",err,query))
			os.Exit(1)
		}
		ss.Info.Printf("fqdn = %v valid=%v, ensname=%v\n",rec.NodeHash,null_ens_name.Valid,null_ens_name.String)
		if null_ens_name.Valid {
			rec.ENS_Name = null_ens_name.String
		}
		if len(rec.ENS_Name) == 0 {
			rec.ENS_Name = "ENS Name is not public"
		}
		rec.NumTextKeyValuePairs = null_num_keys.Int64
		node_addresses := ss.Get_ens_node_addresses(rec.NodeHash)
		if len(node_addresses) > 0 {
			rec.CurAddr = node_addresses[0].Address
			rec.CurAddrAid = node_addresses[0].Aid
		}
		rec.NodeAddressHistory = node_addresses
		records = append(records,rec)
	}
	return records,total_rows
}
func (ss *SQLStorage) Get_user_report_profits(user_aid int64) []p.UserRepProfit {

	records := make([]p.UserRepProfit,0,8)

	var query string
	query = "SELECT " +
				"EXTRACT(EPOCH FROM time_stamp)::BIGINT ts," +
				"TO_CHAR(c.time_stamp, 'dd/mm/yyyy HH:ii')," +
				"m.market_aid," +
				"ma.addr," +
				"m.extra_info::json->>'description' AS descr," +
				"m.market_type," +
				"m.outcomes," +
				"outcome_idx," +
				"amount, " +
				"rep," +
				"1 AS rtype," +
				"tx.tx_hash " +
			"FROM crowdsourcer_redeemed c " +
				"JOIN transaction tx ON c.tx_id=tx.id " +
				"JOIN market m ON c.market_aid=m.market_aid " +
				"JOIN address ma ON c.market_aid=ma.address_id " +
			"WHERE (c.reporter_aid=$1) AND (amount<rep) " + // amount>rep because otherwise it is a return of unused stake (pre loaded in case somebody disputes the outcome)
			"UNION ALL " +
			"SELECT  " +
				"EXTRACT(EPOCH FROM time_stamp)::BIGINT ts," +
				"TO_CHAR(r.time_stamp, 'dd/mm/yyyy HH:ii')," +
				"m.market_aid," +
				"ma.addr," +
				"m.extra_info::json->>'description' AS descr," +
				"m.market_type," +
				"m.outcomes," +
				"outcome_idx," +
				"amount, " +
				"rep," +
				"0 AS rtype," +
				"tx.tx_hash " +
			"FROM irep_redeem r " +
				"JOIN transaction tx ON r.tx_id=tx.id " +
				"JOIN market m ON r.market_aid=m.market_aid " +
				"JOIN address ma ON r.market_aid=ma.address_id " +
			"WHERE (r.reporter_aid=$1) AND (amount<rep) " + // amount>rep because otherwise it is a return of unused stake (pre loaded in case somebody disputes the outcome)
			"ORDER BY ts"

	rows,err := ss.db.Query(query,user_aid)
	if (err!=nil) {
		ss.Log_msg(fmt.Sprintf("DB error: %v (query=%v)",err,query))
		os.Exit(1)
	}

	defer rows.Close()
	for rows.Next() {
		var rec p.UserRepProfit
		var mkt_type int
		var outcomes string
		err := rows.Scan(
			&rec.TimeStamp,
			&rec.DateTime,
			&rec.MarketAid,
			&rec.MarketAddr,
			&rec.MarketDescr,
			&mkt_type,
			&outcomes,
			&rec.OutcomeIdx,
			&rec.RepInvested,
			&rec.RepReturned,
			&rec.RType,
			&rec.TxHash,
		)
		if (err!=nil) {
			if err != sql.ErrNoRows {
				ss.Log_msg(fmt.Sprintf("Error in Get_reporting_table(): %v : %v",err,query))
				os.Exit(1)
			}
			return records
		}
		rec.OutcomeStr = get_outcome_str(uint8(mkt_type),rec.OutcomeIdx,&outcomes)
		rec.TxHashSh = p.Short_hash(rec.TxHash)
		rec.Profit = rec.RepReturned - rec.RepInvested
		rec.ROI = 100*(rec.RepReturned - rec.RepInvested)/rec.RepInvested
		records = append(records,rec)
	}
	return records
}
