// Data Base Storage
package dbs

import (
	"fmt"
	"os"
	"math/big"
	"database/sql"
	_  "github.com/lib/pq"

//	"github.com/ethereum/go-ethereum/common"

	p "github.com/PredictionExplorer/augur-explorer/primitives"
)
func (ss *SQLStorage) calculate_profit(num_ticks int64,win_tick int64,amount *big.Int,price *big.Int) *big.Int {
	// Calculates the profit for a position
	// Source:
	//	https://github.com/AugurProject/augur/blob/master/packages/augur-core/src/contracts/trading/ProfitLoss.sol#L82a
	// this function executes only when we know the User has a profit, losses are calculated as -frozen_funds
	tick_adjustment:=big.NewInt(0)
	tick_adjustment.SetString("10000000000000000",10)

	ether_in_weis:= new(big.Int)
	ether_in_weis.SetString("1000000000000000000",10)	// 10 ^ 18

	win_price := big.NewInt(win_tick)
	win_price.Mul(win_price,ether_in_weis)

	ticks:=big.NewInt(num_ticks)
	adjusted_ticks:=big.NewInt(0)
	adjusted_ticks.Mul(ticks,ether_in_weis)

	ss.Info.Printf("adjusted ticks = %v, win_price=%v , price=%v\n",adjusted_ticks.String(),win_price.String(),price.String())
	result := big.NewInt(0)
	if amount.Cmp(zero) < 0 {	// Short
		result.Sub(win_price,price)
		ss.Info.Printf("substracted price =%v, new price=%v\n",price.String(),result.String())
		result.Mul(result,amount)
		ss.Info.Printf("multiplication = %v\n",result.String())
//		result.Quo(result,ticks) part of original formula of Augur, but this op doesn't work
//		ss.Info.Printf("division by %v : %v\n",ticks.String(),result.String())
	}
	if amount.Cmp(zero) > 0 {	// Long
		result.Sub(win_price,price)
		ss.Info.Printf("substracted price = %v\n",result.String())
		result.Mul(result,amount)
		ss.Info.Printf("multiplication = %v\n",result.String())
//		result.Quo(result,ticks)
//		ss.Info.Printf("division by %v : %v\n",ticks.String(),result.String())
	}
	// Note: if amount == 0 , returns 0
	return result
}
func (ss *SQLStorage) set_all_unclaimed_to_claimed(market_aid int64,eoa_aid int64,timestamp int64) {
	var query string
	query = "UPDATE claim_funds SET claim_status=2,claim_ts=TO_TIMESTAMP($3) " +
			"WHERE claim_status=1 AND market_aid=$1 AND eoa_aid=$2"
	ss.Info.Printf("update_claimed: market=%v, aid=%v, query=%v",market_aid,eoa_aid,query)
	_,err := ss.db.Exec(query,market_aid,eoa_aid,timestamp)
	if err!=nil {
		ss.Log_msg(fmt.Sprintf("DB error: %v (query=%v)",err,query))
		os.Exit(1)
	}
}
func (ss *SQLStorage) calculate_profit_loss_for_all_users(market_aid int64,block_num p.BlockNumber,tx_id int64,timestamp int64,evt *p.MktFinalizedEvt) {

	var query string

	market_type,num_ticks:=ss.get_market_type_and_ticks(market_aid)
	_ = market_type

	query = "SELECT " +
				"pl.id," +
				"pl.eoa_aid," +
				"pl.outcome_idx," +
				"round(pl.net_position*1e+18)::text," +
				"round(pl.avg_price*1e+18)::text, " +
				"pl.frozen_funds::text," +
				"round(pl.frozen_funds*1e+18)::text AS frozen_funds_big " +
			"FROM profit_loss AS pl " +
			"WHERE (market_aid = $1) AND (closed_position=0) " +
			" ORDER BY pl.id"

//	d_query:=strings.ReplaceAll(query,"$1",fmt.Sprintf("%v",market_aid))
//	ss.Info.Printf("update_losing_positions(): CLAIM FUNDS: query=%v\n",d_query)
	rows,err := ss.db.Query(query,market_aid)
	if (err!=nil) {
		ss.Log_msg(fmt.Sprintf("DB error: %v (query=%v)",err,query))
		os.Exit(1)
	}
	defer rows.Close()
	for rows.Next() {
		var (
			claim_status int = 0
			pl_id int64
			eoa_aid int64
			outcome_idx int
			str_net_position string
			str_price string
			frozen_funds string
			frozen_funds_big string
		)
		err=rows.Scan(&pl_id,&eoa_aid,&outcome_idx,&str_net_position,&str_price,&frozen_funds,&frozen_funds_big)
		if err!=nil {
			ss.Log_msg(fmt.Sprintf("DB error: %v ; q=%v",err,query))
			os.Exit(1)
		}
		ff_big:=big.NewInt(0)
		ff_big.SetString(frozen_funds_big,10)
		// frozen funds negative, calculate FF using Augur's business logic
		net_position := big.NewInt(0)
		net_position.SetString(str_net_position,10)
		price := big.NewInt(0)
		price.SetString(str_price,10)
		winning_tick := evt.WinningPayoutNumerators[outcome_idx].Int64()
		profit := ss.calculate_profit(num_ticks,winning_tick,net_position,price)
		ss.Info.Printf("loss = %v\n",profit)
		profit_str := profit.String()
		query = "INSERT INTO claim_funds(" +
						"block_num,tx_id,eoa_aid,market_aid,outcome_idx,last_pl_id,"+
						"claim_status,autocalculated,final_profit,unfrozen_funds" +
					") VALUES (" +
						"$1,$2,$3,$4,$5,$6,$7,$8,(("+profit_str+"/1e+36)),("+frozen_funds+")" +
					")"
		if profit.Cmp(zero) > 0 {
			claim_status=1
		}
		if ff_big.Cmp(zero) < 0 {
			claim_status=2 // if we have negative frozen funds, then this position is considered claimed
		}
		ss.Info.Printf("update_losing: pl_id=%v eoa_aid=%v frozen=%v profit=%v\n",pl_id,eoa_aid,frozen_funds,profit.String())
		ss.Info.Printf("update_losing: INSERT: %v\n",query)
		_,err:=ss.db.Exec(query,
			block_num,
			tx_id,
			eoa_aid,
			market_aid,
			outcome_idx,
			pl_id,
			claim_status,
			true,
		)
		if (err!=nil) {
			ss.Log_msg(fmt.Sprintf("DB error in update_losing_positions(): %v ; q=%v",err,query))
			os.Exit(1)
		}
	}
}
func (ss *SQLStorage) Update_claim_status(agtx *p.AugurTx,evt *p.TradingProceedsClaimed,timestamp int64) {
	// Note: we don't use outcome in WHERE clause because Proceeds aren't reported for all outcomes,
	//		however just knowing that proceeds where claimed is enough to update all the outcomes
	//		This function will be executed multiple times in a transaction, but that's ok
	market_aid := ss.Lookup_address_id(evt.Market.String())
	signer_aid := ss.Lookup_address_id(agtx.TxMsg.From().String())
	//outcome_idx := evt.Outcome.Int64()

	var query string
	query = "UPDATE claim_funds SET claim_status=2,autocalculated=FALSE,claim_ts=TO_TIMESTAMP($3) " +
			"WHERE market_aid=$1 AND eoa_aid=$2 AND claim_status=1"
	_,err := ss.db.Exec(query,	market_aid,signer_aid,timestamp)
	if err != nil {
		ss.Log_msg(fmt.Sprintf("DB error: %v:q=%v",err,query))
		os.Exit(1)
	}
}
func (ss *SQLStorage) Insert_profit_loss_evt(agtx *p.AugurTx,eoa_aid int64,evt *p.ProfitLossChanged) int64  {

	var query string
	var err error

	_,err = ss.lookup_universe_id(evt.Universe.String())
	if err != nil {
		ss.Log_msg(fmt.Sprintf("Universe %v not found on ProfitLossChanged event\n",evt.Universe.String()))
		os.Exit(1)
	}
	market_aid,market_type := ss.lookup_market_id(evt.Market.String())
	wallet_aid := ss.Lookup_or_create_address(evt.Account.String(),agtx.BlockNum,agtx.TxId)

	var qty_divisor string = "18" //16
	var price_divisor string = "18" //20
	if market_type == p.MktTypeScalar {
		qty_divisor = "18"
		price_divisor = "18"
	}

	outcome_idx := evt.Outcome.Int64()
	net_position := evt.NetPosition.String()
	if len(net_position)==0 {
		net_position = "0"
	}
	avg_price := evt.AvgPrice.String()
	if len(avg_price) == 0 {
		avg_price = "0"
	}
	realized_profit := evt.RealizedProfit.String()
	if len(realized_profit) == 0 {
		realized_profit = "0"
	}
	frozen_funds := evt.FrozenFunds.String()
	if len(frozen_funds) == 0 {
		frozen_funds = "0"
	}
	realized_cost := evt.RealizedCost.String()
	if len(realized_cost) == 0 {
		realized_cost = "0"
	}
	time_stamp := evt.Timestamp.Int64()
	ss.close_previous_positions(market_aid,eoa_aid,int(outcome_idx),"")
	var immed_profit_str string
	previous_rprofit_str,previous_ff_str:=ss.get_previous_profit_and_ff(market_aid,eoa_aid,int(outcome_idx))
	prev_profit:=new(big.Int)
	if len(previous_rprofit_str) > 0 {
		prev_profit.SetString(previous_rprofit_str,10)
	}
	immed_pl:=new(big.Int)
	immed_pl.Sub(evt.RealizedProfit,prev_profit)
	immed_profit_str=immed_pl.String()

	var immed_ff_str string = "0"
	var prev_ff *big.Int
	if len(previous_ff_str) == 0 {
		prev_ff=big.NewInt(0)
	} else {
		prev_ff = new(big.Int)
	}
	prev_ff.SetString(previous_ff_str,10)
	immed_ff:=new(big.Int)
	immed_ff.Sub(evt.FrozenFunds,prev_ff)
	immed_ff_str=immed_ff.String()
	ss.Info.Printf("previous_realized_profit = %v, current realized_profit=%v, immediate=%v\n",previous_rprofit_str,evt.RealizedProfit.String(),immed_profit_str)
	ss.Info.Printf("previous_frozen_funds = %v, current_frozen_funds=%v, immediate=%v\n",previous_ff_str,evt.FrozenFunds.String(),immed_ff_str)
	query = "INSERT INTO profit_loss (" +
				"block_num," + 
				"tx_id," +
				"market_aid," +
				"eoa_aid," +
				"wallet_aid," +
				"outcome_idx," +
				"mktord_id," +
				"net_position," +
				"avg_price," +
				"frozen_funds," +
				"realized_profit," +
				"realized_cost," +
				"immediate_profit," +
				"immediate_ff," +
				"time_stamp" +
			") VALUES($1,$2,$3,$4,$5,$6,$7," +
				"(" +net_position+ "/1e+"+qty_divisor+")," +
				"(" +avg_price+ "/1e+"+price_divisor+")," +
				"(" +frozen_funds+ "/1e+36)," +
				"(" +realized_profit+ "/1e+36)," +
				"(" +realized_cost + "/1e+36)," +
				"(" +immed_profit_str + "/1e+36)," +
				"(" +immed_ff_str + "/1e+36)," +
				"TO_TIMESTAMP($8)" +
			") RETURNING id,realized_profit,realized_cost,net_position"

	var null_pl_id sql.NullInt64
	var null_profit sql.NullFloat64
	var null_rcost sql.NullFloat64
	var null_volume sql.NullFloat64
	var pl_id int64 = 0
	row := ss.db.QueryRow(query,
								agtx.BlockNum,
								agtx.TxId,
								market_aid,
								eoa_aid,
								wallet_aid,
								outcome_idx,
								*ss.mkt_order_id_ptr,// note, this contains meaningful value only because we reverse event processing order
								time_stamp,
	)
	err=row.Scan(&null_pl_id,&null_profit,&null_rcost,&null_volume);
	if (err!=nil) {
		if err == sql.ErrNoRows {
			//
		} else {
			ss.Log_msg(fmt.Sprintf("DB error: %v; q=%v VALUES: block_num=%v,tx_id=%v,market_aid=%v, eoa_aid=%v, wallet_aid=%v, outcome_idx=%v, order_id=%v, time_stamp=%v",err,query,agtx.BlockNum,agtx.TxId,market_aid,eoa_aid,wallet_aid,outcome_idx,*ss.mkt_order_id_ptr,time_stamp))
			os.Exit(1)
		}
	} else {
		pl_id = null_pl_id.Int64
	}
	if null_volume.Valid {
		if null_volume.Float64 == 0 {
			// Volume = 0 means the User has closed all his positions,
			// therefore we must mark position as closed in the DB too
			ss.close_previous_positions(market_aid,eoa_aid,int(outcome_idx),"")
		}
	}

	return pl_id
}
func (ss *SQLStorage) Get_profit_loss(eoa_aid int64) []p.PLEntry {
	return ss.Get_trade_data(eoa_aid,false)
}
func (ss *SQLStorage) Insert_profit_loss_debug_rec(pchg *p.PosChg) {

	market_aid := ss.Lookup_address_id(pchg.Mkt_addr.String())
	wallet_aid := ss.Lookup_address_id(pchg.Wallet_addr.String())

	var query string

	query = "INSERT INTO pl_debug(block_num,market_aid,wallet_aid,outcome_idx," +
									"profit_loss,frozen_funds,net_position,avg_price) " +
				" VALUES(" +
					"$1,$2,$3,$4,"+
					"(" + pchg.ProfitLoss.String() + "/1e+36)," +
					"(" + pchg.FrozenFunds.String()+ "/1e+36)," +
					"(" + pchg.NetPos.String() + "/1e+16), " +
					"(" + pchg.AvgPrice.String() + "/1e+20) " +
				") " +
				"ON CONFLICT DO NOTHING"

	_,err := ss.db.Exec(query,pchg.BlockNum,market_aid,wallet_aid,pchg.Outcome.Int64())
	if (err!=nil) {
		ss.Log_msg(fmt.Sprintf("DB Error: %v q=%v",err,query));
		os.Exit(1)
	}
}
