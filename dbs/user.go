// Data Base Storage
package dbs

import (
	"fmt"
	"os"
	"strconv"
	"math/big"
	"database/sql"
	_  "github.com/lib/pq"

	"github.com/ethereum/go-ethereum/common"

	p "augur-extractor/primitives"
)
func (ss *SQLStorage) fill_block_info(ui *p.UserInfo,user_aid int64) {

	var query string
	query = "SELECT address_id,addr,b.block_num, " +
			"FLOOR(EXTRACT(EPOCH FROM b.ts))::BIGINT as ts " +
			"FROM address a,block b " +
			"WHERE (a.address_id=$1) AND (a.block_num=b.block_num) "
	row := ss.db.QueryRow(query,user_aid)
	err := row.Scan(&ui.EOAAid,&ui.EOAAddr,&ui.BlockNum,&ui.TimeStamp)
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

	var query string
	query = "SELECT " +
				"s.wallet_aid," +
				"a.addr as eoa_addr," +
				"w.addr as wallet_addr," +
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
				"r.top_profit " +
			"FROM ustats as s " +
			"LEFT JOIN address AS a ON s.eoa_aid = a.address_id " +
			"LEFT JOIN address AS w ON s.wallet_aid = w.address_id " +
			"LEFT JOIN uranks AS r ON s.eoa_aid = r.eoa_aid " +
			"WHERE s.eoa_aid = $1"

	row := ss.db.QueryRow(query,user_aid)
	var err error
	var (
		eoa_addr		sql.NullString
		wallet_addr		sql.NullString
		top_profits		sql.NullFloat64
		top_trades		sql.NullFloat64
	)
	ui.EOAAid = user_aid
	err=row.Scan(
				&ui.WalletAid,
				&eoa_addr,
				&wallet_addr,
				&ui.TotalTrades,
				&ui.MarketsCreated,
				&ui.MarketsTraded,
				&ui.WithdrawReqs,
				&ui.DepositReqs,
				&ui.TotalReports,
				&ui.TotalDesignated,
				&ui.ProfitLoss,
				&ui.ReportProfits,
				&ui.AffProfits,
				&ui.MoneyAtStake,
				&ui.ValidityBonds,
				&ui.TotalWithdrawn,
				&ui.TotalDeposited,
				&top_trades,
				&top_profits,
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
	if eoa_addr.Valid {
		ui.EOAAddr = eoa_addr.String
		ui.EOAAddrSh = p.Short_address(eoa_addr.String)
	}
	if wallet_addr.Valid {
		ui.WalletAddr = wallet_addr.String
		ui.WalletAddrSh = p.Short_address(wallet_addr.String)
	}
	if top_profits.Valid {
		ui.TopProfit = top_profits.Float64
	}
	if top_trades.Valid {
		ui.TopTrades = top_trades.Float64
	}
	if ui.MoneyAtStake < 0 {
		ui.HedgingProfits = true
	}
	ui.UnclaimedProfit=ss.Get_unclaimed_profit(user_aid)
	return ui,nil
}
func (ss *SQLStorage) Get_ranking_data_for_all_users() []p.RankStats {

	var query string
	query = "SELECT eoa_aid,total_trades,profit_loss,volume_traded FROM ustats"

	rows,err := ss.db.Query(query)
	if (err!=nil) {
		ss.Log_msg(fmt.Sprintf("DB error: %v (query=%v)",err,query))
		os.Exit(1)
	}
	records := make([]p.RankStats,0,8)
	defer rows.Close()
	for rows.Next() {
		var rec p.RankStats
		err=rows.Scan(&rec.EoaAid,&rec.TotalTrades,&rec.ProfitLoss,&rec.VolumeTraded)
		records = append(records,rec)
	}
	return records
}
func (ss *SQLStorage) Update_top_profit_rank(eoa_aid int64,value float64,profit float64) int64 {

	var query string
	query = "UPDATE uranks SET top_profit = $2,profit=$3 WHERE eoa_aid = $1"
	res,err:=ss.db.Exec(query,eoa_aid,value,profit)
	if (err!=nil) {
		ss.Log_msg(fmt.Sprintf("update_top_profit_rank() failed: %v, q=%v",err,query))
		os.Exit(1)
	}
	affected_rows,err:=res.RowsAffected()
	if err!=nil {
		ss.Log_msg(fmt.Sprintf("Error getting RowsAffected in update_top_profit_rank(): %v",err))
	}
	if affected_rows == 0 {
		query = "INSERT INTO uranks(eoa_aid,top_profit,profit) VALUES($1,$2,$3)"
		_,err:=ss.db.Exec(query,eoa_aid,value,profit)
		if (err!=nil) {
			ss.Log_msg(fmt.Sprintf("update_top_profit_rank() failed: %v, q=%v",err,query))
		}

	}
	return affected_rows
}
func (ss *SQLStorage) Update_top_total_trades_rank(eoa_aid int64,value float64,total_trades int64) int64 {

	var query string
	query = "UPDATE uranks SET top_trades = $2,total_trades=$3 WHERE eoa_aid = $1"
	res,err:=ss.db.Exec(query,eoa_aid,value,total_trades)
	if (err!=nil) {
		ss.Log_msg(fmt.Sprintf("update_top_total_trades_rank() failed: %v, q=%v",err,query))
		os.Exit(1)
	}
	affected_rows,err:=res.RowsAffected()
	if err!=nil {
		ss.Log_msg(fmt.Sprintf("Error getting RowsAffected in update_top_total_trades_rank(): %v",err))
	}
	if affected_rows == 0 {
		query = "INSERT INTO uranks(eoa_aid,top_trades,total_trades) VALUES($1,$2,$3)"
		_,err:=ss.db.Exec(query,eoa_aid,value,total_trades)
		if (err!=nil) {
			ss.Log_msg(fmt.Sprintf("update_top_total_trades_rank() failed: value=%v, err: %v, q=%v",value,err,query))
			os.Exit(1)
		}

	}
	return affected_rows
}
func (ss *SQLStorage) Update_top_volume_rank(eoa_aid int64,value float64,volume float64) int64 {

	var query string
	query = "UPDATE uranks SET top_volume = $2,volume=$3 WHERE eoa_aid = $1"
	res,err:=ss.db.Exec(query,eoa_aid,value,volume)
	if (err!=nil) {
		ss.Log_msg(fmt.Sprintf("update_top_volume_rank() failed: %v, q=%v",err,query))
		os.Exit(1)
	}
	affected_rows,err:=res.RowsAffected()
	if err!=nil {
		ss.Log_msg(fmt.Sprintf("Error getting RowsAffected in update_top_volume_rank(): %v",err))
	}
	if affected_rows == 0 {
		query = "INSERT INTO uranks(eoa_aid,top_volume,volume) VALUES($1,$2,$3)"
		_,err:=ss.db.Exec(query,eoa_aid,value,volume)
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
			"LEFT JOIN address AS a ON r.eoa_aid = a.address_id " +
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
		err=rows.Scan(&rec.EOAAddr,&rec.Percentage,&rec.ProfitLoss)
		records = append(records,rec)
	}
	return records
}
func (ss *SQLStorage) Get_top_trade_makers() []p.TradeMaker {

	var query string
	query = "SELECT a.addr,r.top_trades,r.total_trades FROM uranks AS r " +
			"LEFT JOIN address AS a ON r.eoa_aid = a.address_id " +
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
		err=rows.Scan(&rec.EOAAddr,&rec.Percentage,&rec.TotalTrades)
		records = append(records,rec)
	}
	return records
}
func (ss *SQLStorage) Get_top_volume_makers() []p.VolumeMaker {

	var query string
	query = "SELECT a.addr,r.top_volume,r.volume FROM uranks AS r " +
			"LEFT JOIN address AS a ON r.eoa_aid = a.address_id " +
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
		err=rows.Scan(&rec.EOAAddr,&rec.Percentage,&rec.Volume)
		records = append(records,rec)
	}
	return records
}
func (ss *SQLStorage) Get_user_reports(eoa_aid int64,limit int) []p.UserReport {

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
			"WHERE (r.market_aid = m.market_aid) and (r.eoa_aid=$1) " +
			"ORDER BY r.rpt_timestamp"
	if limit > 0 {
		query = query +	" LIMIT " + strconv.Itoa(limit)
	}

	records := make([]p.UserReport,0,8)
	var rows *sql.Rows
	var err error
	rows,err = ss.db.Query(query,eoa_aid)
	if (err!=nil) {
		if err == sql.ErrNoRows {
			return records
		}
		ss.Log_msg(fmt.Sprintf("DB error: %v (query=%v)",err,query))
		os.Exit(1)
	}

	defer rows.Close()
	for rows.Next() {
		var rec p.UserReport
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
func (ss *SQLStorage) Get_user_markets(eoa_aid int64) []p.InfoMarket {

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
			"WHERE eoa_aid = $1 " +
			"ORDER BY " +
				"m.market_aid "

	rows,err := ss.db.Query(query,eoa_aid)
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
		var status_code int
		err=rows.Scan(
					&rec.MktAddr,
					&rec.Signer,
					&rec.MktCreator,
					&rec.EndDate,
					&description,
					&longdesc,
					&categories,
					&rec.Outcomes,
					&rec.MktType,
					&rec.MktTypeStr,
					&status_code,
					&rec.Fee,
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
		rec.MktCreator=p.Short_address(rec.MktCreator)
		rec.Status=get_market_status_str(p.MarketStatus(status_code))
		records = append(records,rec)
	}
	return records
}
func (ss *SQLStorage) Get_unclaimed_profit(eoa_aid int64) float64 {

	var query string
	query = "SELECT sum(final_profit) AS pl FROM claim_funds WHERE eoa_aid=$1 AND claim_status=1"
	row := ss.db.QueryRow(query,eoa_aid)

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
	market_aid := ss.lookup_address_id(mkt_addr.String())

	var query string
	query = "SELECT wa.addr,pl.outcome_idx FROM profit_loss AS pl " +
				"LEFT JOIN address AS wa ON pl.wallet_aid=wa.address_id " +
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
		pchg.Wallet_addr = common.HexToAddress(addr)
		pchg.Outcome = new(big.Int)
		pchg.Outcome.SetInt64(int64(outcome_idx))
		output = append(output,pchg)
	}
	return output
}
func (ss *SQLStorage) Get_active_markets_for_user(eoa_aid int64) []p.InfoMarket {

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
				"cur_volume AS volume, " +
				"s.volume_traded, " +
				"s.total_trades " +
			"FROM market as m " +
				"JOIN trd_mkt_stats AS s ON m.market_aid = s.market_aid " +
				"LEFT JOIN address AS ma ON m.market_aid = ma.address_id " +
				"LEFT JOIN address AS sa ON m.eoa_aid= sa.address_id " +
				"LEFT JOIN address AS ca ON m.wallet_aid = ca.address_id " +
			"WHERE s.eoa_aid = $1 AND m.status < 4" +
			"ORDER BY s.volume_traded DESC"

	rows,err := ss.db.Query(query,eoa_aid)
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
		var status_code int
		err=rows.Scan(
					&rec.MktAddr,
					&rec.Signer,
					&rec.MktCreator,
					&rec.EndDate,
					&description,
					&longdesc,
					&category,
					&rec.Outcomes,
					&rec.MktType,
					&rec.MktTypeStr,
					&status_code,
					&rec.Status,
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
		rec.Status=get_market_status_str(p.MarketStatus(status_code))
		rec.MktAddrSh=p.Short_address(rec.MktAddr)
		rec.MktCreatorSh=p.Short_address(rec.MktCreator)
		records = append(records,rec)
	}
	return records
}
func (ss *SQLStorage) Get_user_trades_for_market(eoa_aid int64,mkt_aid int64) []p.MarketTrade {
	// get market trades with mixed outcomes
	var query string
	query = "SELECT " +
				"o.order_id," +
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
				"o.outcome," +
				"m.market_type AS mtype," +
				"m.outcomes AS outcomes_str " +
			"FROM mktord AS o " +
				"LEFT JOIN address AS a ON o.market_aid=a.address_id " +
				"LEFT JOIN address AS fa ON o.eoa_fill_aid=fa.address_id " +
				"LEFT JOIN address AS ca ON o.eoa_aid=ca.address_id " +
				"LEFT JOIN market AS m ON o.market_aid = m.market_aid " +
			"WHERE o.market_aid = $1 AND ((o.eoa_aid=$2) OR (o.eoa_fill_aid=$2)) " +
			"ORDER BY o.block_num DESC,o.time_stamp DESC"

	ss.Info.Printf("eoa_aid=%v, market_aid=%v, query=%v\n",eoa_aid,mkt_aid,query)
	var rows *sql.Rows
	var err error
	rows,err = ss.db.Query(query,mkt_aid,eoa_aid)
	if (err!=nil) {
		ss.Log_msg(fmt.Sprintf("DB error: %v (query=%v)",err,query))
		os.Exit(1)
	}
	records := make([]p.MarketTrade,0,8)

	defer rows.Close()
	for rows.Next() {
		var rec p.MarketTrade
		var mkt_type int
		var outcomes string
		err=rows.Scan(
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
			&outcomes,
		)
		if err!=nil {
			ss.Log_msg(fmt.Sprintf("DB error: %v, q=%v",err,query))
			os.Exit(1)
		}
		rec.MktAddrSh=p.Short_address(rec.MktAddr)
		rec.CreatorAddrSh=p.Short_address(rec.CreatorAddr)
		rec.FillerAddrSh=p.Short_address(rec.FillerAddr)
		rec.OutcomeStr = get_outcome_str(uint8(mkt_type),rec.Outcome,&outcomes)
		records = append(records,rec)
		ss.Log_msg(fmt.Sprintf("Record appended %+v\n",rec))
	}
	return records
}
func (ss *SQLStorage) Get_user_open_orders(user_aid int64) (p.UserInfo,error) {
	// open orders on 0x Mesh network
	var rec p.UserInfo
	var query string
	_=query
	return rec,nil
}
