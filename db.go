package main

import (
	"os"
	"fmt"
	"net"
	"math/big"
	"strings"
//	"context"
	"database/sql"
	"encoding/hex"
	"encoding/json"
	_  "github.com/lib/pq"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/0xProject/0x-mesh/zeroex"
)
type SQLStorage struct {
	db		*sql.DB
}
func show_connect_error() {
	fmt.Printf(`AugurExtractor: can't connect to PostgreSQL database.
				Check that you have set AUGUR_EXTRACTOR_USERNAME,AUGUR_EXTRACTOR_PASSWORD,AUGUR_EXTRACTOR_DATABASE
				and AUGUR_EXTRACTOR_HOST environment variables`);
}
func connect_to_storage() *SQLStorage {
	var err error
	host,port,err:=net.SplitHostPort(os.Getenv("AUGUR_EXTRACTOR_HOST"))
	if (err!=nil) {
		host=os.Getenv("AUGUR_EXTRACTOR_HOST")
		port="5432"
	}
	conn_str := "user='"+
				os.Getenv("AUGUR_EXTRACTOR_USERNAME") +
				"' dbname='" +
				os.Getenv("AUGUR_EXTRACTOR_DATABASE") +
				"' password='" +
				os.Getenv("AUGUR_EXTRACTOR_PASSWORD") +
				"' host='" +
				host +
				"' port='" +
				port +
				"'";
	db,err := sql.Open("postgres",conn_str);
	if (err!=nil) {
		show_connect_error()
	} else {

	}
	_,err = db.Exec("SET timezone TO 0")		// Setting timezone to UTC (which Augur uses)
	if (err!=nil) {
		Fatalf("DB Error: %v",err);
	}

	ss := new(SQLStorage)
	ss.db = db
	return ss
}
func (ss *SQLStorage) check_main_stats() {

	var query string
	query="SELECT id FROM main_stats LIMIT 1";
	row := ss.db.QueryRow(query)
	var null_id sql.NullInt64
	var err error
	err=row.Scan(&null_id);
	if (err!=nil) {
		if err == sql.ErrNoRows {
			query="INSERT INTO main_stats(universe_id) VALUES(1)";
			_,_ =ss.db.Exec(query)
		} else {
			Fatalf("Error in check_main_stats(): %v, q=%v",err,query)
		}
	}
}
func (ss *SQLStorage) get_last_block_num() (BlockNumber,bool) {

	var query string
	query="SELECT block_num FROM last_block LIMIT 1";
	row := ss.db.QueryRow(query)
	var null_block_num sql.NullInt64
	var err error
	err=row.Scan(&null_block_num);
	if (err!=nil) {
		if err == sql.ErrNoRows {
			return -1,false
		} else {
			Fatalf("Error in get_last_block_num(): %v",err)
		}
	}
	if (null_block_num.Valid) {
		return BlockNumber(null_block_num.Int64),true
	} else {
		return -1,false
	}
}
func (ss *SQLStorage) set_last_block_num(block_num BlockNumber) {

	bnum := int64(block_num)
	var query string = "UPDATE last_block SET block_num=$1 WHERE block_num < $1"
	res,err:=ss.db.Exec(query,bnum)
	if (err!=nil) {
		Fatalf("set_last_block_num() failed: %v",err);
	}
	affected_rows,err:=res.RowsAffected()
	if err!=nil {
		Fatalf("Error getting RowsAffected in set_last_block(): %v",err)
	}
	if affected_rows>0 {
		// break
	} else {
		query = "INSERT INTO last_block VALUES($1)"
		_,err := ss.db.Exec(query,bnum)
		if (err!=nil) {
			Fatalf("set_last_block_num() failed on INSERT: %v",err);
		}
	}
}
func (ss *SQLStorage) lookup_universe_id(addr string) int64 {

	var query string
	query="SELECT universe_id FROM universe WHERE universe_addr=$1"
	var universe_id int64 = 0
	err:=ss.db.QueryRow(query,addr).Scan(&universe_id);
	if (err!=nil) {
		if (err==sql.ErrNoRows) {
			Fatalf("DB error: Universe doesn't exist (addr=%v). Database wasn't initialized correctly",addr)
		} else {
			Fatalf("DB error looking up for Universe record: %v",err);
		}
	}
	return universe_id
}
func (ss *SQLStorage) lookup_eoa_aid(wallet_aid int64) (int64,error) {

	var addr_id int64;
	var query string
	query="SELECT eoa_id FROM ustats WHERE wallet_aid=$1"
	err:=ss.db.QueryRow(query,wallet_aid).Scan(&addr_id);
	if (err!=nil) {
		return 0,err
	}

	return addr_id,nil
}
func (ss *SQLStorage) lookup_wallet_aid(eoa_aid int64) (int64,error) {

	var addr_id int64;
	var query string
	query="SELECT wallet_aid FROM ustats WHERE eoa_aid=$1"
	err:=ss.db.QueryRow(query,eoa_aid).Scan(&addr_id);
	if (err!=nil) {
		if err!=sql.ErrNoRows {
			Fatalf("DB error: %v, q=%v",err,query)
		} else {
			fmt.Printf("lookup_wallet_aid(%v) error: %v\n",eoa_aid,err)
		}
		return 0,err
	}

	return addr_id,nil
}
func (ss *SQLStorage) nonfatal_lookup_address_id(addr string) (int64,error) {

	var addr_id int64;
	var query string
	query="SELECT address_id FROM address WHERE addr=$1"
	err:=ss.db.QueryRow(query,addr).Scan(&addr_id);
	if (err!=nil) {
		return 0,err
	}

	return addr_id,nil
}
func (ss *SQLStorage) lookup_address(eoa_aid int64) (string,error) {

	var addr string;
	var query string
	query="SELECT addr FROM address WHERE address_id=$1"
	err:=ss.db.QueryRow(query,eoa_aid).Scan(&addr);
	return addr,err
}
func (ss *SQLStorage) lookup_address_id(addr string) int64 {

	var addr_id int64;
	var query string
	query="SELECT address_id FROM address WHERE addr=$1"
	err:=ss.db.QueryRow(query,addr).Scan(&addr_id);
	if (err!=nil) {
		if (err==sql.ErrNoRows) {
			Fatalf("DB error: address %v does not exist",addr)
		} else {
			Fatalf("DB error upon address lookup: %v",err)
		}
	}

	return addr_id
}
func (ss *SQLStorage) lookup_or_create_address(addr string) int64 {

	var addr_id int64;
	var query string
	query="SELECT address_id FROM address WHERE addr=$1"
	err:=ss.db.QueryRow(query,addr).Scan(&addr_id);
	if (err!=nil) {
		if (err==sql.ErrNoRows) {
			query = "INSERT INTO address(addr) VALUES($1) RETURNING address_id"
			row:=ss.db.QueryRow(query,addr);
			err:=row.Scan(&addr_id)
			if err!=nil {
				Fatalf(fmt.Sprintf("DB error in address insertion: %v : %v",query,err))
			}
			if addr_id==0 {
				Fatalf(fmt.Sprintf("DB error, addr_id after INSERT is 0"))
			}
			return addr_id
		}
	}
	if (err!=nil) {
		Fatalf(fmt.Sprintf("DB error in getting address id : %v",err))
	}

	return addr_id
}
func (ss *SQLStorage) lookup_or_create_category(categories string) int64 {

	var cat_id int64
	var query string

	query="SELECT cat_id FROM category WHERE category=$1"
	err:=ss.db.QueryRow(query,categories).Scan(&cat_id);
	if (err!=nil) {
		if (err==sql.ErrNoRows) {
			query = "INSERT INTO category(category) VALUES($1) RETURNING cat_id"
			row:=ss.db.QueryRow(query,categories);
			err:=row.Scan(&cat_id)
			if err!=nil {
				Fatalf(fmt.Sprintf("DB error in category insertion: %v : %v",query,err))
			}
			if cat_id==0 {
				Fatalf(fmt.Sprintf("DB error, cat_id after INSERT is 0"))
			}
			return cat_id
		}
	}
	if (err!=nil) {
		Fatalf(fmt.Sprintf("DB error in getting category id : %v",err))
	}

	return cat_id
}
func (ss *SQLStorage) insert_market_created_evt(block_num BlockNumber,tx_id int64,signer common.Address,wallet_aid int64,evt *MarketCreatedEvt) {

	var query string
	var market_aid int64;
	market_aid = ss.lookup_or_create_address(evt.Market.String())
	signer_aid := ss.lookup_or_create_address(signer.String())
	// check if Market is already registered
	query = "SELECT market_aid FROM market WHERE market_aid=$1"
	err:=ss.db.QueryRow(query,market_aid).Scan(&market_aid);
	if (err!=nil) {
		if (err==sql.ErrNoRows) {
			// break
		} else {
			Fatalf("DB error: %v, q=%v",err,query)
		}
	} else {
		// market already registered, sliently exit
		return
	}
	creator_aid := ss.lookup_or_create_address(evt.MarketCreator.String())
	universe_id := ss.lookup_universe_id(evt.Universe.String())
	eoa_aid := ss.lookup_or_create_address(evt.MarketCreator.String())
	reporter_aid := ss.lookup_or_create_address(evt.DesignatedReporter.String())
	fmt.Printf("create_market: signer_aid = %v (%v), creator_aid=%v (%v), reporter_id=%v (%v) , wallet_aid =%v\n",
				signer_aid,signer.String(),
				creator_aid,evt.MarketCreator.String(),
				reporter_aid,evt.DesignatedReporter.String(),
				wallet_aid,
			)
	if signer_aid == creator_aid { // a case only seen in Test environment, production check pending
		// Normally signer != creator, but this happens only in Dev (local testnet), so we have to fix it
		//creator_aid = wallet_aid // this doesn't work, if starting blockchain from block 0, wallt isn't created yet
		wallet_aid = creator_aid
		fmt.Printf("create_market: fixed creator id to contract address %v (wallet_aid)\n",wallet_aid)
	} else {
		eoa_aid = signer_aid
		wallet_aid = creator_aid
	}
	if wallet_aid == 0 {
		Fatalf("insert_market_created_evt(): creator addr = %v, wallet_aid = 0, can't continue, exiting\n",
					evt.MarketCreator.String())
	}
	prices := bigint_ptr_slice_to_str(&evt.Prices,",")
	outcomes := outcomes_to_str(&evt.Outcomes,",")

	var extra_info ExtraInfo
	json.Unmarshal([]byte(evt.ExtraInfo), &extra_info)
	fmt.Printf("extra_info unmarshalled: %+v\n",extra_info)
	categories := strings.Join(extra_info.Categories,",")
	fmt.Printf("market_categories: %v\n",categories)
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
			max_ticks,
			create_timestamp,
			fee,
			prices,
			market_type,
			extra_info,
			outcomes,
			no_show_bond
		) VALUES ($1,$2,$3,$4,$5,$6,$7,$8,TO_TIMESTAMP($9),$10,TO_TIMESTAMP($11),` +
						evt.FeePerCashInAttoCash.String() +
						"/1e+18,$12,$13,$14,$15,$16)";

	result,err := ss.db.Exec(query,
			block_num,
			tx_id,
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
			evt.NoShowBond.String())
	if err != nil {
		Fatalf("DB error: can't insert into market table: %v: q=%v",err,query)
	}
	rows_affected,err:=result.RowsAffected()
	if err != nil {
		Fatalf("DB error: %v, %q",err,query)
	}
	if rows_affected > 0 {
	} else {
		Fatalf("DB error: couldn't insert into Market table. Rows affeced = 0")
	}
	if len(outcomes) == 0 {
		if evt.MarketType == 0 {	// Yes / No
			outcomes = "Invalid,No,Yes"
		}
	}
	ss.init_market_outcome_volumes(market_aid,outcomes)
}
func (ss *SQLStorage) init_market_outcome_volumes(market_aid int64,outcomes string) {

	var query string
	outcomes_list := strings.Split(outcomes,",")
	for outcome_idx:=0 ; outcome_idx < len(outcomes_list) ; outcome_idx ++ {
		query = "INSERT INTO outcome_vol(" +
					"market_aid," +
					"outcome_idx" +
				") VALUES(" +
					"$1," +
					"$2" +
				")"
		_,err := ss.db.Exec(query,market_aid,outcome_idx)
		if (err!=nil) {
			Fatalf("DB error: %v; q=%v",err,query);
		}
	}
}
func (ss *SQLStorage) insert_market_oi_changed_evt(block *types.Header,evt *MarketOIChangedEvt) {
	// Note: this event arrives with evt.Market set to 0x0000000000000000000000000 (a contract bug?) ,
	//			so we pass the market address as parameter ('market_addr') to the function
	var query string
	market_aid := ss.lookup_address_id(evt.Market.String())
	ts_inserted := int64(block.Time)
	query = "INSERT INTO oi_chg(market_aid,ts_inserted,oi) VALUES($1,TO_TIMESTAMP($2),(" +
			evt.MarketOI.String() +
			"/1e+18))"
	result,err := ss.db.Exec(query,market_aid,ts_inserted)
	if err != nil {
		Fatalf("DB error: can't insert into oi_chg table: %v; q=%v",err,query)
	}
	rows_affected,err:=result.RowsAffected()
	if err != nil {
		Fatalf("DB error: %v, q=%v",err,query)
	}
	if rows_affected > 0 {
		return
	} else {
		Fatalf("DB error: couldn't insert into oi_chg table. Rows affeced = 0")
	}

	fmt.Printf("Set market (id=%v) open interest to %v",market_aid,evt.MarketOI.String())
}
func (ss *SQLStorage) insert_market_order_evt(block_num BlockNumber,tx_id int64,signer common.Address,eoa_aid int64,evt *MktOrderEvt) {

	// depending on the order action (Create/Cancel/Fill) different table is used for storage
	//		Create/Cancel order actions go to 'oorders' (Open Orders) table because these orders
	//		do not alter market's open interest.
	//		Fill order goes to 'mktord' table because the share has been created and now
	//		open interest increased
	var wallet_aid int64;
	wallet_aid = ss.lookup_or_create_address(evt.AddressData[0].String())
	var wallet_fill_aid int64 = 0;
	if len(evt.AddressData) > 1 {
		wallet_fill_aid = ss.lookup_or_create_address(evt.AddressData[1].String())
	}
	universe_id := ss.lookup_universe_id(evt.Universe.String())
	_ = universe_id	// ToDo: add universe_id match condition (for market)
	market_aid := ss.lookup_address_id(evt.Market.String())
	eoa_fill_aid := ss.lookup_or_create_address(signer.String())

	var oaction OrderAction = OrderAction(evt.EventType)
	var otype OrderType = OrderType(evt.OrderType)
	var order_id = hex.EncodeToString(evt.OrderId[:])
	// uint256data legend
	// 0:  price
	// 1:  amount
	// 2:  outcome
	// 3:  tokenRefund (Cancel)
	// 4:  sharesRefund (Cancel)
	// 5:  fees (Fill)
	// 6:  amountFilled (Fill)
	// 7:  timestamp
	// 8:  sharesEscrowed
	// 9:  tokensEscrowed
	price := evt.Uint256Data[0].String()
	amount := evt.Uint256Data[1].String()
	outcome_idx := evt.Uint256Data[2].Int64()
	token_refund := evt.Uint256Data[3].String()
	shares_refund := evt.Uint256Data[4].String()
	fees := evt.Uint256Data[5].String()
	amount_filled := evt.Uint256Data[6].String()
	time_stamp := evt.Uint256Data[7].Int64()
	shares_escrowed := evt.Uint256Data[8].String()
	tokens_escrowed := evt.Uint256Data[9].String()

	var query string

	query = "DELETE FROM oorders WHERE order_id = $1"
	_,err := ss.db.Exec(query,market_aid)
	if err!=nil {
		fmt.Printf("DB error: couldn't delete open order with order_id = %v\n",order_id)
	}

	fmt.Printf("OrderAction = %v, otype=%v, order_id=%v\n",oaction,otype,order_id)
	fmt.Printf("Filling existing order %v\n",order_id)
	query = `
		INSERT INTO mktord(
			tx_id,
			market_aid,
			eoa_aid,
			wallet_aid,
			eoa_fill_aid,
			wallet_fill_aid,
			block_num,
			oaction,
			otype,
			price,
			amount,
			outcome,
			token_refund,
			shares_refund,
			fees,
			amount_filled,
			time_stamp,
			shares_escrowed,
			tokens_escrowed,
			trade_group,
			order_id
		) VALUES (
				$1,$2,$3,$4,$5,$6,$7,$8,$9,
				(` + price + "/1e+2)," +
				"(" + amount + "/1e+18)," +
				"$10," +
				"(" + token_refund + "/1e+18)," +
				"(" + shares_refund + "/1e18)," +
				"(" + fees + "/1e18)," +
				"(" + amount_filled + "/1e16)," +
				"TO_TIMESTAMP($11)," +
				"$12,$13,$14,$15) RETURNING id"

	var null_id sql.NullInt64
	err=ss.db.QueryRow(query,
			tx_id,
			market_aid,
			eoa_aid,
			wallet_aid,
			eoa_fill_aid,
			wallet_fill_aid,
			block_num,
			oaction,
			otype,
			outcome_idx,
			time_stamp,
			shares_escrowed,
			tokens_escrowed,
			hex.EncodeToString(evt.TradeGroupId[:]),
			order_id,
	).Scan(&null_id);
	if (err!=nil) {
		Fatalf("DB error: can't insert into mktord table: %v, q=%v",err,query)
	}
	if null_id.Valid {
		market_order_id = null_id.Int64
	} else {
		market_order_id = 0
	}
	query = "UPDATE " +
				"outcome_vol " +
			"SET " +
				"last_price = "+price+ " " +
			"WHERE " +
				"market_aid = $1 AND outcome_idx = $2"
	res,err:=ss.db.Exec(query,market_aid,outcome_idx)
	if (err!=nil) {
		Fatalf("DB error: %v ; q=%v",err,query);
	}
	affected_rows,err:=res.RowsAffected()
	if err!=nil {
		Fatalf("DB error in rows affected: %v, q=%v",err,query)
	}
	if affected_rows == 0 {
		fmt.Printf("Last price for market_aid = %v and outcome_idx = %v wasn't updated",
					market_aid,outcome_idx)
	}
}
func (ss *SQLStorage) insert_open_order(evt *zeroex.OrderEvent,eoa_addr string,ospec *ZxMeshOrderSpec) {
	// Insert an open order, this order needs to be Filled by another market participant
	// It also can be canceled by its creator (with another transaction)
	order := evt.SignedOrder.Order
	ohash,err := order.ComputeOrderHash()
	if err != nil {
		fmt.Printf("Error at computing 0x Mesh order: %v",err)
	}
	order_id := ohash.String()
	evt_timestamp := evt.Timestamp.Unix()
	expiration := order.ExpirationTimeSeconds.Int64()
	wallet_aid := ss.lookup_or_create_address(order.MakerAddress.String())
	eoa_aid := ss.lookup_or_create_address(eoa_addr)
	fmt.Printf("creating open order made by %v : %+v\n",eoa_addr,ospec)
	market_aid := ss.lookup_address_id(ospec.Market.String())
	price := float64(ospec.Price.Int64())/100
	otype := ospec.Type	// Bid/Ask
	amount := order.MakerAssetAmount.String()

	var query string
	query = "INSERT INTO oostats(market_aid,eoa_aid,outcome_idx) VALUES($1,$2,$3)"
	_,err = ss.db.Exec(query,market_aid,eoa_aid,ospec.Outcome)
	if err != nil {
		fmt.Printf("DB error: %v, q=%v\n",err,query)
	}
	query = "INSERT INTO oorders(" +
				"market_aid,otype,wallet_aid,eoa_aid,price,amount,outcome_idx," +
				"evt_timestamp,srv_timestamp,expiration,order_id" +
			") VALUES($1,$2,$3,$4,$5,"+amount+"/1e+18,$6,TO_TIMESTAMP($7),TO_TIMESTAMP($8),NOW(),$9)"
	result,err := ss.db.Exec(query,
			market_aid,
			otype,
			wallet_aid,
			eoa_aid,
			price,
			ospec.Outcome,
			evt_timestamp,
			expiration,
			order_id)
	if err != nil {
		fmt.Printf("DB error: can't insert into open orders table: %v, q=%v",err,query)
		return
	}
	rows_affected,err:=result.RowsAffected()
	if err != nil {
		Fatalf("DB error: %v, q=%v",err,query)
	}
	if rows_affected > 0 {
		return
	} else {
		Fatalf("DB error: couldn't insert into Open Orders table. Rows affeced = 0")
	}
}
func (ss *SQLStorage) delete_open_0x_order(order_hash string) {

	var query string
	query = "DELETE FROM oorders WHERE order_id = $1"
	result,err := ss.db.Exec(query,order_hash)
	if err!=nil {
		fmt.Printf("DB error: couldn't delete open order with order_id = %v, q=%v\n",order_hash,query)
	}
	rows_affected,err:=result.RowsAffected()
	if rows_affected == 0  {
		fmt.Printf("DB error: couldn't delete open order with order_id = %v (not found)\n",order_hash)
	}
}
func (ss *SQLStorage) insert_market_finalized_evt(evt *MktFinalizedEvt) {

	universe_id := ss.lookup_universe_id(evt.Universe.String())
	_ = universe_id	// ToDo: add universe_id match condition (for market)
	market_aid := ss.lookup_address_id(evt.Market.String())
	fin_timestamp := evt.Timestamp.Int64()
	winning_payouts := bigint_ptr_slice_to_str(&evt.WinningPayoutNumerators,",")

	var query string
	query = "INSERT INTO mkt_fin(market_aid,fin_timestamp,winning_payouts) VALUES($1,TO_TIMESTAMP($2),$3)"
	_,err := ss.db.Exec(query,market_aid,fin_timestamp,winning_payouts)
	if err != nil {
		Fatalf("DB error: can't update market finalization of market %v : %v, q=%v",market_aid,err,query)
	}
	ss.update_market_status(market_aid,MktStatusFinalized)
	ss.update_losing_positions(market_aid,evt)
}
func (ss *SQLStorage) update_market_status(market_aid int64,status MarketStatus) {
	var query string
	query = "UPDATE " +
				"market " +
			"SET " +
				"status=$2" +
			"WHERE " +
				"market_aid = $1"

	res,err:=ss.db.Exec(query,market_aid,status)
	if (err!=nil) {
		Fatalf("DB error: %v ; q=%v",err,query);
	}
	affected_rows,err:=res.RowsAffected()
	if err!=nil {
		Fatalf("DB error: %v ; q=%v",err,query);
	}
	if affected_rows == 0 {
		fmt.Printf("Couldn't update market status = %v for market %v",status,market_aid)
	} else {
		fmt.Printf("MKTSTATUS: market_aid = %v, status = %v\n",market_aid,status)
	}
}
func (ss *SQLStorage) update_losing_positions(market_aid int64,evt *MktFinalizedEvt) {

	// this function marks losing positions as closed (because we don't have ProfitLoss event
	//			on a losing position (position with wrong outcome)
	var query string
	query = "SELECT market_type FROM market WHERE market_aid=$1"

	var market_type int
	err:=ss.db.QueryRow(query,market_aid).Scan(&market_type);
	if (err!=nil) {
		if (err==sql.ErrNoRows) {
			return
		}
		Fatalf("DB Error: %v, q=%v\n",err,query)
	}

	var where_condition string
	switch market_type {
		case 0:		// Yes/No
			hundred := big.NewInt(100)
			if hundred.Cmp(evt.WinningPayoutNumerators[0]) == 0 { // Invalid
				where_condition = "implmenentation pending"
			}
			if hundred.Cmp(evt.WinningPayoutNumerators[1]) ==0 { // No wins
				where_condition = " (((outcome_idx = 2) AND (net_position > 0)) OR " +
								  "  ((outcome_idx = 1) AND (net_position < 0))) "
			}
			if hundred.Cmp(evt.WinningPayoutNumerators[2]) ==0 { // Yes wins
				where_condition = " (((outcome_idx = 2) AND (net_position < 0)) OR " +
								  "  ((outcome_idx = 1) AND (net_position > 0))) "
			}
			query = "UPDATE profit_loss " +
						"SET closed_position = 1, " +
							"final_profit = frozen_funds " +
						"WHERE (market_aid = $1) AND "+
						where_condition

		default:
	}
	res,err:=ss.db.Exec(query,market_aid)
	if (err!=nil) {
		Fatalf("DB error: %v ; q=%v",err,query);
	}
	affected_rows,err:=res.RowsAffected()
	if err!=nil {
		Fatalf("DB error in rows affected: %v",err)
	}
	fmt.Printf("Market finalized. amount of closed losing positions: %v\n",affected_rows)
}
func (ss *SQLStorage) insert_initial_report_evt(block_num BlockNumber,tx_id int64,signer common.Address,evt *InitialReportSubmittedEvt) {

	_= ss.lookup_universe_id(evt.Universe.String())
	market_aid := ss.lookup_address_id(evt.Market.String())
	reporter_aid := ss.lookup_or_create_address(evt.Reporter.String())
	signer_aid := ss.lookup_or_create_address(signer.String())
	ini_reporter_aid := ss.lookup_or_create_address(evt.InitialReporter.String())

	amount_staked := evt.AmountStaked.String()
	payout_numerators := bigint_ptr_slice_to_str(&evt.PayoutNumerators,",")
	next_win_start := evt.NextWindowStartTime.Int64()
	next_win_end := evt.NextWindowEndTime.Int64()
	rpt_timestamp := evt.Timestamp.Int64()

	fmt.Printf("insert_initial_report_evt(): market_aid=%v, reporter_id=%v, signer_aid=%v\n",
					market_aid,reporter_aid,signer_aid)
	var query string
	query = `
		INSERT INTO report (
			block_num,
			tx_id,
			market_aid,
			wallet_aid,
			eoa_aid,
			ini_reporter_aid,
			is_initial,
			is_designated,
			amount_staked,
			pnumerators,
			description,
			next_win_start,
			next_win_end,
			rpt_timestamp
		) VALUES (
			$1,$2,$3,$4,$5,$6,$7,$8,(` + amount_staked + `/1e+18),$9,$10,
			TO_TIMESTAMP($11),
			TO_TIMESTAMP($12),
			TO_TIMESTAMP($13)
		)`
	result,err := ss.db.Exec(query,
			block_num,
			tx_id,
			market_aid,
			reporter_aid,
			signer_aid,
			ini_reporter_aid,
			true,
			evt.IsDesignatedReporter,
			payout_numerators,
			evt.Description,
			next_win_start,
			next_win_end,
			rpt_timestamp)
	if err != nil {
		Fatalf("DB error: can't insert into report table: %v,q=%v",err,query)
	}
	rows_affected,err:=result.RowsAffected()
	if err != nil {
		Fatalf("DB error: %v, q=%v",err,query)
	}
	if rows_affected > 0 {
		//break
	} else {
		Fatalf("DB error: couldn't insert into InitialReport table. Rows affeced = 0")
	}
	// set 'Reporting' status
	// ToDo: possibly migrate to triggers (or maybe not)
	ss.update_market_status(market_aid,MktStatusReported)
}
func (ss *SQLStorage) insert_market_volume_changed_evt(block_num BlockNumber,tx_id int64,evt *MktVolumeChangedEvt) {

	market_aid := ss.lookup_address_id(evt.Market.String())

	volume := evt.Volume.String()
	outcome_vols := bigint_ptr_slice_to_str(&evt.OutcomeVolumes,",")
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
			block_num,
			tx_id,
			market_aid,
			outcome_vols,
			timestamp)
	if err != nil {
		Fatalf("DB error: can't insert into volume table: %v, q=%v",err,query)
	}
	rows_affected,err:=result.RowsAffected()
	if err != nil {
		Fatalf("DB error: %v, q=%v",err,query)
	}
	if rows_affected > 0 {
		//break
	} else {
		Fatalf("DB error: couldn't insert into InitialReport table. Rows affeced = 0")
	}

	// Updates volume per outcome in an indexed table for querying market info
	for outcome_idx := 0; outcome_idx < len(evt.OutcomeVolumes) ; outcome_idx++ {
		query = "UPDATE " +
					"outcome_vol " +
				"SET " +
					"volume = "+evt.OutcomeVolumes[outcome_idx].String()+"/1e+18 " +
				"WHERE " +
					"market_aid = $1 AND outcome_idx = $2"
		res,err:=ss.db.Exec(query,market_aid,outcome_idx)
		if (err!=nil) {
			Fatalf("DB error: %v ; q=%v",err,query);
		}
		affected_rows,err:=res.RowsAffected()
		if err!=nil {
			Fatalf("DB error in rows affected: %v",err)
		}
		if affected_rows>0 {
			// break
		} else {
			query = "INSERT INTO outcome_vol(" +
						"market_aid," +
						"outcome_idx," +
						"volume" +
					") VALUES(" +
						"$1," +
						"$2," +
						evt.OutcomeVolumes[outcome_idx].String() + "/1e+18" +
					")"
			_,err := ss.db.Exec(query,market_aid,outcome_idx)
			if (err!=nil) {
				Fatalf("DB error: %v; q=%v",err,query);
			}
		}
	}
}
func (ss *SQLStorage) insert_dispute_crowd_contrib(block_num BlockNumber,tx_id int64,signer common.Address,evt *DisputeCrowdsourcerContributionEvt) {

	_= ss.lookup_universe_id(evt.Universe.String())
	market_aid := ss.lookup_address_id(evt.Market.String())
	reporter_aid := ss.lookup_or_create_address(evt.Reporter.String())
	signer_aid := ss.lookup_or_create_address(signer.String())
	disputed_aid := ss.lookup_or_create_address(evt.DisputeCrowdsourcer.String())

	amount_staked := evt.AmountStaked.String()
	payout_numerators := bigint_ptr_slice_to_str(&evt.PayoutNumerators,",")
	cur_stake := evt.CurrentStake.String()
	stake_remaining := evt.StakeRemaining.String()
	dispute_round := evt.DisputeRound.Int64()
	rpt_timestamp := evt.Timestamp.Int64()

	fmt.Printf("insert_dispute_crows_contrib(): market_aid=%v, reporter_id=%v, signer_aid=%v",
					market_aid,reporter_aid,signer_aid)
	var query string
	query = `
		INSERT INTO report (
			block_num,
			tx_id,
			market_aid,
			wallet_aid,
			eoa_aid,
			disputed_aid,
			dispute_round,
			amount_staked,
			description,
			pnumerators,
			current_stake,
			stake_remaining,
			rpt_timestamp
		) VALUES ($1,$2,$3,$4,$5,$6,$7,`+amount_staked+`/1e+18,$8,$9,
				`+cur_stake+`/1e+18,`+stake_remaining+`/1e+18,TO_TIMESTAMP($10))`
	result,err := ss.db.Exec(query,
			block_num,
			tx_id,
			market_aid,
			reporter_aid,
			signer_aid,
			disputed_aid,
			dispute_round,
			evt.Description,
			payout_numerators,
			rpt_timestamp)
	if err != nil {
		Fatalf("DB error: can't insert dispute into report table: %v; q=%v",err,query)
	}
	rows_affected,err:=result.RowsAffected()
	if err != nil {
		Fatalf("DB error: %v, q=%v",err,query)
	}
	if rows_affected == 0 {
		Fatalf("DB error: couldn't insert dispute into Report table. Rows affeced = 0")
	}
	ss.update_market_status(market_aid,MktStatusDisputing)
}
func (ss *SQLStorage) insert_share_balance_changed_evt(block_num BlockNumber,tx_id int64,evt *ShareTokenBalanceChanged) {

	_= ss.lookup_universe_id(evt.Universe.String())
	market_aid := ss.lookup_address_id(evt.Market.String())
	account_aid := ss.lookup_or_create_address(evt.Account.String())

	outcome := evt.Outcome.Int64()
	balance := evt.Balance.String()

	var query string

	query = "UPDATE sbalances SET balance = (" + balance + "/1e+18) " +
				"WHERE " +
					"market_aid = $1 AND " +
					"account_aid = $2 AND " +
					"outcome_idx = $3"
	result,err := ss.db.Exec(query,	market_aid,account_aid,outcome)
	if err != nil {
		Fatalf("DB error: can't update 'sbalances' for account %v, market %v : %v; q=%v",
					evt.Account.String(),evt.Market.String(),err,query)
	}
	rows_affected,err:=result.RowsAffected()
	if err != nil {
		Fatalf("DB error: %v",err)
	}
	fmt.Printf("No error, rows affected = %v\n",rows_affected)
	if rows_affected > 0 {
		fmt.Printf("Update to sbalances %v , outcome %v holds %v \n",evt.Account.String(),outcome,balance);
		//break
	} else {
		fmt.Printf("Insert to sbalances (%v outcome %v bal=%v\n",evt.Account.String(),outcome,balance);
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
		result,err := ss.db.Exec(query,block_num,tx_id,account_aid,market_aid,outcome)
		if err != nil {
			Fatalf("DB error: can't insert into sbalances table: %v, q=%v",err,query)
		}
		rows_affected,err:=result.RowsAffected()
		if err != nil {
			Fatalf("DB error: %v, query=%v",err,query)
		}
		if rows_affected > 0 {
			return
		} else {
			Fatalf("DB error: couldn't insert into 'sbalances' table. Rows affeced = 0")
		}
	}
}
func (ss *SQLStorage) insert_block(block *types.Header,num_tx int64)  bool {

	var query string
	var parent_block_num int64
	parent_hash := block.ParentHash.String()

	query="SELECT block_num,parent_hash FROM block WHERE hash=$1"
	err:=ss.db.QueryRow(query,parent_hash).Scan(&parent_block_num);
	if (err!=nil) {
		if (err==sql.ErrNoRows) {
			if block.Number.Uint64() == 0 {
				// Genesis. Allow.
			} else {
				if (parent_block_num + 1) != int64(block.Number.Uint64()) {
					fmt.Printf("Block sequence broken after block %v\n",parent_block_num)
					return false
				}
			}
		}
	}

	block_num := int64(block.Number.Uint64())
	block_hash := block.Hash().String()
	query = `
		INSERT INTO block(
			block_num,
			num_tx,
			block_hash,
			parent_hash
		) VALUES ($1,$2,$3,$4)`

	result,err := ss.db.Exec(query,
			block_num,
			num_tx,
			block_hash,
			parent_hash)
	if err != nil {
		Fatalf("DB error: can't insert into block  table: %v, q=%v",err,query)
	}
	rows_affected,err:=result.RowsAffected()
	if err != nil {
		Fatalf("DB error: %v, q=%v",err,query)
	}
	if rows_affected > 0 {
		return true
	}
	Fatalf("DB error: couldn't insert into block table. Rows affeced = 0")
	return false
}
func (ss *SQLStorage) insert_transaction(block_num BlockNumber,tx *types.Transaction) int64 {

	var query string
	var tx_id int64

	tx_hash := tx.Hash().String()
	query = `
		INSERT INTO transaction (
			block_num,
			tx_hash
		) VALUES ($1,$2) RETURNING id`

	row := ss.db.QueryRow(query,
			block_num,
			tx_hash)
	err := row.Scan(&tx_id)
	if err != nil {
		Fatalf("DB error: can't insert into transactions table: %v, q=%v",err,query)
	}
	return tx_id
}
func (ss *SQLStorage) fix_chainsplit(block *types.Header) BlockNumber {

	var query string
	var my_block_num int64
	parent_hash := block.ParentHash.String()
	query = "SELECT block_num FROM block WHERE block_hash = $1"
	row := ss.db.QueryRow(query,parent_hash)
	err := row.Scan(&my_block_num);
	if (err!=nil) {
		if err==sql.ErrNoRows {
			Fatalf("Chainsplit detected, I don't have the parent hash %v, exiting. ",parent_hash)
		} else {
			Fatalf("DB error: %v, q=%v",err,query)
		}
	}
	cur_block_num := int64(block.Number.Uint64())
	if cur_block_num > (my_block_num + MAX_BLOCKS_CHAIN_SPLIT) {
		Fatalf("Chainsplit detected, and it is more than %v blocks, aborting.",MAX_BLOCKS_CHAIN_SPLIT)
	}
	query = "DELETE FROM block WHERE block_num > $1 CASCADE"
	_,err = ss.db.Exec(query,my_block_num)
	if (err!=nil) {
		Fatalf("DB error: %v, q=%v, block_num=%v",err,query,my_block_num);
	}
	return BlockNumber(my_block_num + 1)	// parent + 1 = current
}
func (ss *SQLStorage) block_delete_with_everything(block_num BlockNumber) {

	// deletes block table and all the other tables receieve cascaded DELETEs also
	var query string
	query = "DELETE FROM block WHERE block_num = $1"
	_,err := ss.db.Exec(query,block_num)
	if (err!=nil) {
		Fatalf("DB error: %v (block_num=%v, %v)",err,block_num,query);
	}
}
func (ss *SQLStorage) get_active_market_list(off int, lim int) []InfoMarket {

	var query string
	query = "SELECT " +
				"ma.addr as mkt_addr," +
				"CONCAT(LEFT(ma.addr,6),'…',RIGHT(ma.addr,6)) AS mkt_addr_sh," +
				"CONCAT(LEFT(sa.addr,6),'…',RIGHT(sa.addr,6)) AS signer," +
				"ca.addr as mcreator," +
				"CONCAT(LEFT(ca.addr,6),'…',RIGHT(ca.addr,6)) AS mcreator_sh, " +
				"TO_CHAR(end_time,'dd/mm/yyyy HH24:SS UTC') as end_date," + 
				"extra_info::json->>'description' AS descr," +
				"extra_info::json->>'longDescription' AS long_desc," +
				"extra_info::json->>'categories' AS categories," +
				"outcomes," +
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
			"ORDER BY " +
				"m.market_aid " +
			"OFFSET $1 LIMIT $2";

	rows,err := ss.db.Query(query,off,lim)
	if (err!=nil) {
		Fatalf("DB error: %v (query=%v)",err,query);
	}
	var rec InfoMarket
	records := make([]InfoMarket,0,8)

	defer rows.Close()
	for rows.Next() {
		var longdesc sql.NullString
		var status_code int
		err=rows.Scan(
					&rec.MktAddr,
					&rec.MktAddrSh,
					&rec.Signer,
					&rec.MktCreator,
					&rec.MktCreatorSh,
					&rec.EndDate,
					&rec.Description,
					&longdesc,
					&rec.Categories,
					&rec.Outcomes,
					&rec.MktType,
					&status_code,
					&rec.Status,
					&rec.Fee,
					&rec.OpenInterest,
					&rec.CurVolume,
		)
		if err!=nil {
			Fatalf(fmt.Sprintf("DB error: %v, q=%v",err,query))
		}
		if longdesc.Valid {
			rec.LongDesc = longdesc.String
		}
		if status_code == 0 {
			// nothing
		} else {
			switch MarketStatus(status_code) {
				case MktStatusReported:
					rec.Status = "Reported"
				case MktStatusDisputing:
					rec.Status = "Disputing"
				case MktStatusFinalized:
					rec.Status = "Finalized"
				case MktStatusFinInvalid:
					rec.Status = "Finalized Invalid"
				default:
			}
		}
		records = append(records,rec)
	}
	return records
}
func (ss *SQLStorage) get_categories() []InfoCategories {

	var query string
	query = "SELECT " +
				"cat_id," +
				"category " +
			"FROM category " +
			"ORDER BY " +
				"category";

	_,err := ss.db.Exec(query)
	if (err!=nil) {
		Fatalf("DB error: %v ,q=%v",err,query);
	}
	rows,err:=ss.db.Query(query)
	if err!=nil {
		if err!=sql.ErrNoRows {
			Fatalf(fmt.Sprintf("Error for query %v: %v",query,err))
		}
	}
	var rec InfoCategories
	records := make([]InfoCategories,0,8)

	defer rows.Close()
	for rows.Next() {
		err=rows.Scan(&rec.CatId,&rec.Category)
		if err!=nil {
			Fatalf(fmt.Sprintf("DB error: %v",err))
		}
		fmt.Printf("going to do split of: %+v\n",rec.Category)
		subcategories := strings.Split(rec.Category,",")
		for i := 0 ; i< len(subcategories); i++ {
			subcategories[i] = strings.Title(subcategories[i])
			fmt.Printf("added subcategory i=%v, subcat = %v\n",i,subcategories[i])
		}
		if len(subcategories) > 0 {	// sometimes last category is empty, delete it
			if len(subcategories[len(subcategories)-1]) == 0 {
				subcategories = subcategories[:len(subcategories)-1]
			}
		}
		rec.Subcategories = subcategories
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
			if len(outcomes_list) > outcome_idx {
				output = outcomes_list[outcome_idx]
			} else {
				output = "??????"
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
func (ss *SQLStorage) get_mkt_trades(mkt_addr string) []MarketTrade {
	// get market trades with mixed outcomes
	fmt.Printf("get_mkt_trades() mkt_addr=%v\n",mkt_addr)
	var where string = ""
	var market_aid int64 = 0;
	if len(mkt_addr) > 0 {
		market_aid = ss.lookup_address_id(mkt_addr)
		where = " WHERE o.market_aid = $1 "
	}
	var query string
	query = "SELECT " +
				"o.order_id," +
				"a.addr as mkt_addr," +
				"CONCAT(LEFT(a.addr,6),'…',RIGHT(a.addr,6)) AS mkt_addr_sh, " +
				"fa.addr as signer_addr," +
				"CONCAT(LEFT(fa.addr,6),'…',RIGHT(fa.addr,6)) AS signer_addr_sh," +
				"CASE oaction " +
					"WHEN 0 THEN 'CREATE' " +
					"WHEN 1 THEN 'CANCEL' " +
					"WHEN 2 THEN 'FILL' " +
				"END AS type, " +
				"CASE o.otype " +
					"WHEN 0 THEN 'BID' " +
					"ELSE 'ASK' " +
				"END AS dir, " +
				"o.time_stamp::date AS date," +
				"o.price, " +
				"o.amount_filled AS amount," +
				"o.outcome," +
				"m.market_type AS mtype," +
				"m.outcomes AS outcomes_str " +
			"FROM mktord AS o " +
				"LEFT JOIN " +
					"address AS a ON o.market_aid=a.address_id " +
				"LEFT JOIN " +
					"address AS fa ON o.eoa_aid=fa.address_id " +
				"LEFT JOIN " +
					"market AS m ON o.market_aid = m.market_aid " +
			where +
			"ORDER BY " +
				"o.time_stamp";

	var rows *sql.Rows
	var err error
	if market_aid > 0 {
		rows,err = ss.db.Query(query,market_aid)
	} else {
		rows,err = ss.db.Query(query)
	}
	if (err!=nil) {
		Fatalf("DB error: %v (query=%v)",err,query);
	}
	records := make([]MarketTrade,0,8)

	defer rows.Close()
	for rows.Next() {
		var rec MarketTrade
		var mkt_type int
		var outcomes string
		err=rows.Scan(
			&rec.OrderHash,
			&rec.MktAddr,
			&rec.MktAddrSh,
			&rec.FillerAddr,
			&rec.FillerAddrSh,
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
			Fatalf(fmt.Sprintf("DB error: %v, q=%v",err,query))
		}
		rec.OutcomeStr = get_outcome_str(uint8(mkt_type),rec.Outcome,&outcomes)
		records = append(records,rec)
	}
	fmt.Printf("get_mkt_trades(): returning %v rows\n",len(records))
	if len(records) == 0 {
		fmt.Printf("null records, q: %v",query)
	}
	return records
}
func (ss *SQLStorage) get_market_info(mkt_addr string,outcome_idx int,oc bool) (InfoMarket,error) {

	var rec InfoMarket
	market_aid,err := ss.nonfatal_lookup_address_id(mkt_addr)
	if err != nil {
		fmt.Printf("market %v not found, returning empty data\n",mkt_addr)
		return rec,err
	}
	rec.MktAid=market_aid
	fmt.Printf("querying info for market aid = %v\n",market_aid)
	var query string
	query = "SELECT " +
				"m.market_type," +
				"ma.addr as mkt_addr," +
				"CONCAT(LEFT(ma.addr,6),'…',RIGHT(ma.addr,6)) AS mkt_addr_sh," +
				"sa.addr AS signer," +
				"CONCAT(LEFT(sa.addr,6),'…',RIGHT(sa.addr,6)) AS signer_sh," +
				"ca.addr as mcreator," +
				"CONCAT(LEFT(ca.addr,6),'…',RIGHT(ca.addr,6)) AS mcreator_sh, " +
				"TO_CHAR(end_time,'dd/mm/yyyy HH24:SS UTC') as end_date," + 
				"extra_info::json->>'description' AS descr," +
				"extra_info::json->>'longDescription' AS long_desc," +
				"extra_info::json->>'categories' AS categories," +
				"outcomes," +
				"CASE m.market_type " +
					"WHEN 0 THEN 'YES/NO' " +
					"WHEN 1 THEN 'CATEGORICAL' " +
					"WHEN 2 THEN 'SCALAR' " +
				"END AS mtype, " +
				"fee," +
				"open_interest AS OI," +
				"cur_volume AS volume " +
			"FROM market as m " +
				"LEFT JOIN " +
					"address AS ma ON m.market_aid = ma.address_id " +
				"LEFT JOIN " +
					"address AS sa ON m.eoa_aid = sa.address_id " +
				"LEFT JOIN " +
					"address AS ca ON m.wallet_aid = ca.address_id " +
			"WHERE market_aid = $1"

	row := ss.db.QueryRow(query,market_aid)
	var mkt_type int
	err=row.Scan(
				&mkt_type,
				&rec.MktAddr,
				&rec.MktAddrSh,
				&rec.Signer,
				&rec.SignerSh,
				&rec.MktCreator,
				&rec.MktCreatorSh,
				&rec.EndDate,
				&rec.Description,
				&rec.LongDesc,
				&rec.Categories,
				&rec.Outcomes,
				&rec.MktType,
				&rec.Fee,
				&rec.OpenInterest,
				&rec.CurVolume,
	)
	if oc { // get outcome string
		rec.CurOutcome = get_outcome_str(uint8(mkt_type),outcome_idx,&rec.Outcomes)
	}
	if err!=nil {
		fmt.Printf("DB error: %v, q=%v\n",err,query)
		return rec,err
	}
	return rec,nil
}
func (ss *SQLStorage) get_outcome_volumes(mkt_addr string) ([]OutcomeVol,error) {

	var rec OutcomeVol
	records := make([]OutcomeVol,0,8)
	market_aid,err := ss.nonfatal_lookup_address_id(mkt_addr)
	if err != nil {
		return records,err
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
			"WHERE o.market_aid = $1"

	var rows *sql.Rows
	rows,err = ss.db.Query(query,market_aid)
	if (err!=nil) {
		Fatalf("DB error: %v (query=%v)",err,query);
	}

	defer rows.Close()
	for rows.Next() {
		var outcomes string
		rec.MktAddr = mkt_addr
		err=rows.Scan(
			&rec.Outcome,
			&rec.Volume,
			&rec.LastPrice,
			&rec.MktType,
			&outcomes,
		)
		if err!=nil {
			Fatalf(fmt.Sprintf("DB error: %v, q=%v",err,query))
		}
		/* to be deleted
		if rec.MktType == 0 { // Yes/No
			switch rec.Outcome {
				case 0:
					rec.OutcomeStr = "Invalid"
				case 1:
					rec.OutcomeStr = "No"
				case 2:
					rec.OutcomeStr = "Yes"
			}
		}
		if rec.MktType == 1 { // Categorical
			outcomes_list := strings.Split(outcomes,",")
			if len(outcomes) > rec.Outcome {
				rec.OutcomeStr = outcomes_list[rec.Outcome]	// ToDo: possibly move to INSERT stage
			}
		}
		if rec.MktType == 2 {
			if rec.Outcome == 0 {
				rec.OutcomeStr = "Invalid"
			} 
			if rec.Outcome == 2 {
				rec.OutcomeStr="Scalar"
			}
			if rec.Outcome == 1 {
				rec.OutcomeStr="-"
			}
		}
		*/
		rec.OutcomeStr = get_outcome_str(uint8(rec.MktType),rec.Outcome,&outcomes)
		fmt.Printf("get_outcome_volumes(): rec.OutcomeStr=%v (extracted from %v)\n",rec.OutcomeStr,outcomes)
		records = append(records,rec)
	}
	return records,nil
}
func (ss *SQLStorage) build_depth_by_otype(market_aid int64,outc int,otype OrderType) []DepthEntry {

	var query string
	query = "SELECT " +
				"o.market_aid," +
				"o.outcome_idx," +
//				"a.addr as mkt_addr," +
//				"CONCAT(LEFT(a.addr,6),'…',RIGHT(a.addr,6)) AS mkt_addr_sh, " +
				"wa.addr AS wallet_addr," +
				"CONCAT(LEFT(wa.addr,6),'…',RIGHT(wa.addr,6)) AS wallet_addr_sh," +
				"ua.addr AS user_addr," +
				"CONCAT(LEFT(ua.addr,6),'…',RIGHT(ua.addr,6)) AS user_addr_sh," +
//				"CASE o.otype " +
//					"WHEN 0 THEN 'BID' " +
//					"ELSE 'ASK' " +
//				"END AS dir, " +
				"o.srv_timestamp::date AS date_created," +
				"o.expiration::date AS expires," +
				"FLOOR(EXTRACT(EPOCH FROM o.expiration))::BIGINT as expires_ts," +
				"o.price AS price, " +
				"o.amount AS volume," +
				"s.num_bids," +
				"s.num_asks," +
				"s.num_cancel " +
//				"m.market_type AS mtype," +
//				"m.outcomes AS outcomes_str " +
			"FROM oorders AS o " +
				"LEFT JOIN " +
					"address AS a ON o.market_aid=a.address_id " +
				"LEFT JOIN " +
					"address AS wa ON o.wallet_aid=wa.address_id " +
				"LEFT JOIN " +
					"address AS ua ON o.eoa_aid=ua.address_id " +
//				"LEFT JOIN " +
//					"market AS m ON o.market_aid = m.market_aid " +
				"LEFT JOIN " +
					"oostats AS s ON (" +
						"o.market_aid=s.market_aid AND " +
						"o.eoa_aid=s.eoa_aid AND " +
						"o.outcome_idx=$2) " +
			"WHERE o.market_aid = $1 AND o.outcome_idx=$2 AND o.otype = $3 " +
			"ORDER BY "
	if otype == OrderTypeBid {
				query = query + "o.price DESC,o.evt_timestamp DESC";
	} else {
				query = query + "o.price ASC,o.evt_timestamp DESC";
	}
	fmt.Printf("q=%v\n",query)
	var accumulated_volume = 0.0
	rows,err := ss.db.Query(query,market_aid,outc,otype)
	if (err!=nil) {
		Fatalf("DB error: %v (query=%v)",err,query);
	}
	records := make([]DepthEntry,0,8)

	defer rows.Close()
	for rows.Next() {
		var rec DepthEntry
		var num_bids sql.NullInt64
		var num_asks sql.NullInt64
		var num_cancels sql.NullInt64
		err=rows.Scan(
			&rec.MktAid,
			&rec.OutcomeIdx,
			&rec.WalletAddr,
			&rec.WalletAddrSh,
			&rec.EOAAddr,
			&rec.EOAAddrSh,
			&rec.DateCreated,
			&rec.Expires,
			&rec.ExpiresTs,
			&rec.Price,
			&rec.Volume,
			&num_bids,
			&num_asks,
			&num_cancels,
		)
		if err!=nil {
			Fatalf(fmt.Sprintf("DB error: %v, q=%v",err,query))
		}
		if num_bids.Valid {
			rec.TotalBids = int32(num_bids.Int64)
		}
		if num_asks.Valid {
			rec.TotalAsks = int32(num_asks.Int64)
		}
		if num_cancels.Valid {
			rec.TotalCancel = int32(num_cancels.Int64)
		}
		accumulated_volume = accumulated_volume + rec.Volume
		rec.AccumVol = accumulated_volume
		records = append(records,rec)
	}
	return records
}
func (ss *SQLStorage) get_price_history_for_outcome(market_aid int64,outc int) []MarketOrder{

	var query string
	query = "SELECT " +
				"o.order_id," +
				"o.market_aid," +
				"s_w_a.addr AS s_w_a_addr," +
				"CONCAT(LEFT(s_w_a.addr,6),'…',RIGHT(s_w_a.addr,6)) AS seller_wallet_addr_sh," +
				"s_e_a.addr AS seller_eoa_addr," +
				"CONCAT(LEFT(s_e_a.addr,6),'…',RIGHT(s_e_a.addr,6)) AS seller_eoa_addr_sh," +
				"b_w_a.addr AS b_w_a_addr," +
				"CONCAT(LEFT(b_w_a.addr,6),'…',RIGHT(b_w_a.addr,6)) AS buyer_wallet_addr_sh," +
				"b_e_a.addr AS byer_eoa_addr," +
				"CONCAT(LEFT(b_e_a.addr,6),'…',RIGHT(b_e_a.addr,6)) AS buyer_eoa_addr_sh," +
				"o.otype, " +
				"CASE o.otype " +
					"WHEN 0 THEN 'BID' " +
					"ELSE 'ASK' " +
				"END AS dir, " +
				"o.time_stamp::date AS date," +
				"FLOOR(EXTRACT(EPOCH FROM o.time_stamp))::BIGINT as created_ts," +
				"o.outcome," +
				"o.price AS price, " +
				"o.amount_filled AS volume " +
//				"m.market_type AS mtype," +
//				"m.outcomes AS outcomes_str " +
			"FROM mktord AS o " +
				"LEFT JOIN " +
					"address AS a ON o.market_aid=a.address_id " +
				"LEFT JOIN address AS s_w_a ON o.wallet_aid=s_w_a.address_id " +
				"LEFT JOIN address AS s_e_a ON o.eoa_aid=s_e_a.address_id " +
				"LEFT JOIN address AS b_w_a ON o.wallet_fill_aid=b_w_a.address_id " +
				"LEFT JOIN address AS b_e_a ON o.eoa_fill_aid=b_e_a.address_id " +
//				"LEFT JOIN " +
//					"market AS m ON o.market_aid = m.market_aid " +
			"WHERE o.market_aid = $1 AND o.outcome=$2 " +
			"ORDER BY o.time_stamp"
//	fmt.Printf("q=%v\n",query)
	var accumulated_volume = 0.0
	rows,err := ss.db.Query(query,market_aid,outc)
	if (err!=nil) {
		Fatalf("DB error: %v (query=%v)",err,query);
	}
	records := make([]MarketOrder,0,8)

	defer rows.Close()
	for rows.Next() {
		var rec MarketOrder
		err=rows.Scan(
			&rec.OrderHash,
			&rec.MktAid,
			&rec.SellerWalletAddr,
			&rec.SellerWalletAddrSh,
			&rec.SellerEOAAddr,
			&rec.SellerEOAAddrSh,
			&rec.BuyerWalletAddr,
			&rec.BuyerWalletAddrSh,
			&rec.BuyerEOAAddr,
			&rec.BuyerEOAAddrSh,
			&rec.OType,
			&rec.Direction,
			&rec.Date,
			&rec.CreatedTs,
			&rec.OutcomeIdx,
			&rec.Price,
			&rec.Volume,
		)
		fmt.Printf("record: price = %v, volume = %v\n",rec.Price,rec.Volume)
		if err!=nil {
			Fatalf(fmt.Sprintf("DB error: %v",err))
		}
		accumulated_volume = accumulated_volume + rec.Volume
		rec.AccumVol = accumulated_volume
		records = append(records,rec)
	}
	return records
}
func (ss *SQLStorage) get_mkt_depth(mkt_addr string,outcome_idx int) *MarketDepth {

	fmt.Printf("get_mkt_depth(mkt=%v,outcome_idx=%v)\n",mkt_addr,outcome_idx)
	market_aid := ss.lookup_address_id(mkt_addr)
	market_depth := new(MarketDepth)
	market_depth.Bids = ss.build_depth_by_otype(market_aid,outcome_idx,OrderTypeBid)
	market_depth.Asks = ss.build_depth_by_otype(market_aid,outcome_idx,OrderTypeAsk)
	return market_depth
}
func (ss *SQLStorage) get_user_info(user_aid int64) UserInfo {

	var query string
	query = "SELECT " +
				"s.wallet_aid," +
				"a.addr as eoa_addr," +
				"CONCAT(LEFT(a.addr,6),'…',RIGHT(a.addr,6)) AS eoa_addr_sh," +
				"w.addr as wallet_addr," +
				"CONCAT(LEFT(w.addr,6),'…',RIGHT(w.addr,6)) AS wallet_addr_sh," +
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
	var ui UserInfo
	var (
		eoa_addr		sql.NullString
		eoa_addr_sh		sql.NullString
		wallet_addr		sql.NullString
		wallet_addr_sh	sql.NullString
		top_profits		sql.NullFloat64
		top_trades		sql.NullFloat64
	)
	ui.EOAAid = user_aid
	err=row.Scan(
				&ui.WalletAid,
				&eoa_addr,
				&eoa_addr_sh,
				&wallet_addr,
				&wallet_addr_sh,
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
				&ui.TotalWithdrawn,
				&ui.TotalDeposited,
				&top_trades,
				&top_profits,
	);
	if (err!=nil) {
		if err == sql.ErrNoRows {
			fmt.Printf("No rows for User Info on eoa_aid = %v\n",user_aid)
		} else {
			Fatalf("DB error: %v, q=%v\n",err,query)
		}
	}
	if eoa_addr.Valid {
		ui.EOAAddr = eoa_addr.String
	}
	if eoa_addr_sh.Valid {
		ui.EOAAddrSh = eoa_addr_sh.String
	}
	if wallet_addr.Valid {
		ui.WalletAddr = wallet_addr.String
	}
	if wallet_addr_sh.Valid {
		ui.WalletAddrSh = wallet_addr_sh.String
	}
	if top_profits.Valid {
		ui.TopProfit = top_profits.Float64
	}
	if top_trades.Valid {
		ui.TopTrades = top_trades.Float64
	}
	return ui
}
func (ss *SQLStorage) get_main_stats() MainStats {

	var query string
	query = "SELECT " +
				"markets_count," +
				"yesno_count," +
				"categ_count," +
				"scalar_count," +
				"active_count," +
				"money_at_stake," +
				"trades_count " +
			"FROM main_stats "

	row := ss.db.QueryRow(query)
	var err error
	var s MainStats
	err=row.Scan(
				&s.MarketsCount,
				&s.YesNoCount,
				&s.CategCount,
				&s.ScalarCount,
				&s.ActiveCount,
				&s.MoneyAtStake,
				&s.TradesCount,
	);
	fmt.Printf("main stats = %+v\n",s)
	if (err!=nil) {
		if err == sql.ErrNoRows {
			fmt.Printf("No rows for market statistics query\n")
		} else {
			Fatalf("DB error: %v, q=%v\n",err,query)
		}
	}
	s.FinalizedCount = (s.YesNoCount + s.CategCount + s.ScalarCount) - s.ActiveCount
	return s
}
func (ss *SQLStorage) process_DAI_token_transfer(evt *Transfer) {

	from_aid := ss.lookup_or_create_address(evt.From.String())
	to_aid := ss.lookup_or_create_address(evt.To.String())
	amount := evt.Value.String()


	var query string
/*
	// pending for removal
	query = "SELECT count(*) AS num_recs FROM market WHERE market_aid = $1"
	row := ss.db.QueryRow(query,to_wallet_aid)
	var null_num sql.NullInt64
	var err error
	err=row.Scan(&null_num);
	if (err!=nil) {
		if err == sql.ErrNoRows {
			row := ss.db.QueryRow(query,from_wallet_aid)
			var null_num sql.NullInt64
			var err error
			err=row.Scan(&null_num);
			if (err!=nil) {
				if err == sql.ErrNoRows {

				} else {
					Fatalf("DB error: %v",err)
				}
			} else {
				transf_type = 2		// From market
			}
		} else {
			Fatalf("DB error: %v",err)
		}
	} else {
		transf_type = 2		// To market
	}
*/
	query = "INSERT INTO dai_transf(" +
				"from_aid,to_aid,amount" +
			") VALUES($1,$2,$3/1e+18)"

	_,err := ss.db.Exec(query,
			from_aid,
			to_aid,
			amount,
	)

	if (err!=nil) {
		Fatalf("DB error: %v q=%v",err,query);
	}

}
func (ss *SQLStorage) insert_profit_loss_evt(block_num BlockNumber,tx_id int64,eoa_aid int64,evt *ProfitLossChanged) int64  {

	_= ss.lookup_universe_id(evt.Universe.String())
	market_aid := ss.lookup_address_id(evt.Market.String())
	wallet_aid := ss.lookup_or_create_address(evt.Account.String())

	outcome_idx := evt.Outcome.Int64()
	net_position := evt.NetPosition.String()
	avg_price := evt.AvgPrice.String()
	realized_profit := evt.RealizedProfit.String()
	frozen_funds := evt.FrozenFunds.String()
	realized_cost := evt.RealizedCost.String()
	time_stamp := evt.Timestamp.Int64()

	var query string

	// Update previous position status on this outcome
	query = "UPDATE profit_loss " +
				"SET closed_position = 1, " +
					"final_profit = ($4/1e+36) " +
				"WHERE " +
						"(market_aid = $1) AND" +
						"(eoa_aid = $2) AND " +
						"(outcome_idx = $3) AND " +
						"(closed_position = 0)"
	res,err:=ss.db.Exec(query,market_aid,eoa_aid,outcome_idx,realized_profit)
	if (err!=nil) {
		Fatalf("DB error: %v ; q=%v",err,query);
	}
	affected_rows,err:=res.RowsAffected()
	if affected_rows == 0 {
		fmt.Printf("pl_calc: notice: no previous trades were closed for this market & outcome\n")
	} else {
		fmt.Printf("pl_calc: notice: %v trades were closed for this market & outcome\n",affected_rows)
		if affected_rows != 1 {
			fmt.Printf("pl_calc: WARNING! 'closed_position' was set for more than 1 trade\n")
			Fatalf("pl_calc: undefined behaviour. please implement this use case\n")
		}
	}

	fmt.Printf("Insert to profitloss (wallet %v outcome %v)\n",evt.Account.String(),outcome_idx);
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
				"time_stamp" +
			") VALUES($1,$2,$3,$4,$5,$6,$7," +
				"(" +net_position+ "/1e+16)," +
				"(" +avg_price+ "/1e+20)," +
				"(" +frozen_funds+ "/1e+36)," +
				"(" +realized_profit+ "/1e+36)," +
				"(" +realized_cost+ "/1e+36)," +
				"TO_TIMESTAMP($8)" +
			") RETURNING id,realized_profit,realized_cost"

	var null_pl_id sql.NullInt64
	var null_profit sql.NullFloat64
	var null_rcost sql.NullFloat64
	var pl_id int64 = 0
	row := ss.db.QueryRow(query,
								block_num,
								tx_id,
								market_aid,
								eoa_aid,
								wallet_aid,
								outcome_idx,
								market_order_id,// note, this contains meaningful value only because we reverse event processing order
								time_stamp,
	)
	err=row.Scan(&null_pl_id,&null_profit,&null_rcost);
	if (err!=nil) {
		if err == sql.ErrNoRows {
			//
		} else {
			Fatalf("DB error: %v; q=%v",err,query)
		}
	} else {
		pl_id = null_pl_id.Int64
	}
	if null_profit.Valid {
		if null_profit.Float64 > 0 {
			/*
			query = "UPDATE profit_loss SET mktord_id = $4 " +
					"WHERE (market_aid=$1) AND (outcome=$2) AND (eoa_aid=$3) AND ("
			_,err:=ss.db.Exec(query,market_aid,outcome_idx,eoa_aid)
			if (err!=nil) {
				Fatalf("DB error: %v ; q=%v",err,query);
			}
			*/
		}
	}
	/*
	if null_rcost.Valid {
		if null_rcost.Float64 > 0 {
			query = "SELECT id FROM profit_loss " +
						"WHERE  (market_aid=$1) AND " +
								"(eoa_aid=$2) AND " +
								"(outcome_idx=$3) AND" +
								"(frozen_funds=$4) AND " +
								"(closed_position=0) " +
						"ORDER by id DESC LIMIT 1"

			d_query := fmt.Sprintf("SELECT id FROM profit_loss " +
						"WHERE  (market_aid=%v) AND " +
								"(eoa_aid=%v) AND " +
								"(outcome_idx=%v) AND" +
								"(frozen_funds=%v) AND " +
								"(closed_position=0) " +
						"ORDER by id DESC LIMIT 1",
								market_aid,eoa_aid,outcome_idx,null_rcost.Float64)
			fmt.Printf("pl_calc: query = %v\n",d_query)
			row := ss.db.QueryRow(query,market_aid,eoa_aid,outcome_idx,null_rcost.Float64)
			var opened_pos_pl_id sql.NullInt64
			var err error
			err=row.Scan(&opened_pos_pl_id);
			if (err!=nil) {
				if err == sql.ErrNoRows {
					fmt.Printf("Error: realized cost > 0 but setting position flag " +
									"to Closed was not possible, order id not found: %+v\n",evt);
				} else {
					Fatalf("DB error: %v ; q=%v",err,query);
				}
			} else {
				fmt.Printf("pl_calc: profit loss record id =%v, setting 'closed' to 1\n",opened_pos_pl_id.Int64)
				query = "UPDATE profit_loss SET closed_position = 1 WHERE id=$1"
				res,err:=ss.db.Exec(query,opened_pos_pl_id.Int64)
				if (err!=nil) {
					Fatalf("DB error: %v ; q=%v",err,query);
				}
				affected_rows,err:=res.RowsAffected()
				if affected_rows == 0 {
					fmt.Printf("Error: realized cost > 0 but setting position flag " +
									"to Closed was not possible" +
									"affected rows=0, pl_id =%v evt: %+v\n",evt);
				}
			}
		}
	}
	*/

	return pl_id
}
func (ss *SQLStorage) get_profit_loss(eoa_aid int64) []PLEntry {
	return ss.get_trade_data(eoa_aid,false)
}
func (ss *SQLStorage) get_open_positions(eoa_aid int64) []PLEntry {
	return ss.get_trade_data(eoa_aid,true)
}
func (ss *SQLStorage) get_trade_data(eoa_aid int64,open_positions bool) []PLEntry {

	var extra_condition string
	if open_positions {
		extra_condition = "(pl.closed_position=0)"
	} else {
		extra_condition = "(pl.closed_position=1)"
	}
	var query string
	query = "SELECT " +
				"pl.market_aid," +
				"m.market_type, " +
				"pl.outcome_idx," +
				"m.outcomes," +
				"substring(extra_info::json->>'description',1,100) as descr," +
				"a.addr as mkt_addr," +
				"CONCAT(LEFT(a.addr,6),'…',RIGHT(a.addr,6)) AS mkt_addr_sh," +
				"w_a.addr AS w_a_addr," +
				"CONCAT(LEFT(w_a.addr,6),'…',RIGHT(w_a.addr,6)) AS wallet_addr_sh," +
				"e_a.addr AS eoa_addr," +
				"CONCAT(LEFT(e_a.addr,6),'…',RIGHT(e_a.addr,6)) AS eoa_addr_sh," +
				"pl.time_stamp::date AS date," +
				"FLOOR(EXTRACT(EPOCH FROM pl.time_stamp))::BIGINT as created_ts," +
				"pl.net_position," +
				"pl.avg_price," +
				"pl.frozen_funds," +
				"pl.realized_profit," +
				"pl.realized_cost," +
				"pl.final_profit," +
				"o.order_id," +
				"o.block_num," +
				"o.eoa_aid," +
				"o.eoa_fill_aid ," +
				"cr_a.addr AS creator_eoa_addr," +
				"CONCAT(LEFT(cr_a.addr,6),'…',RIGHT(cr_a.addr,6)) AS creator_eoa_addr_sh," +
				"fil_a.addr AS filler_eoa_addr," +
				"CONCAT(LEFT(fil_a.addr,6),'…',RIGHT(fil_a.addr,6)) AS filler_eoa_addr_sh " +
			"FROM " +
				"profit_loss AS pl " +
					"LEFT JOIN address AS a ON pl.market_aid=a.address_id " +
					"LEFT JOIN address AS w_a ON pl.wallet_aid=w_a.address_id " +
					"LEFT JOIN address AS e_a ON pl.eoa_aid=e_a.address_id " +
					"LEFT JOIN market AS m ON pl.market_aid = m.market_aid," +
				"mktord AS o " +
					"LEFT JOIN address AS cr_a ON o.eoa_aid = cr_a.address_id " +
					"LEFT JOIN address AS fil_a ON o.eoa_fill_aid = fil_a.address_id " +
			"WHERE (pl.mktord_id=o.id) AND (pl.eoa_aid = $1) AND " +
			extra_condition +
			" ORDER BY pl.time_stamp"
	fmt.Printf("pl.eoa_aid=%v; q=%v\n",eoa_aid,query)
	rows,err := ss.db.Query(query,eoa_aid)
	if (err!=nil) {
		Fatalf("DB error: %v (query=%v)",err,query);
	}
	records := make([]PLEntry,0,8)
	var starting_point PLEntry
	records = append(records,starting_point)
	var accumulator float64 = 0.0
	defer rows.Close()
	for rows.Next() {
		var  (
			rec PLEntry
			outcomes string
			order_hash sql.NullString
			block_num sql.NullInt64
			creator_eoa_aid int64
			filler_eoa_aid int64
			creator_addr string
			creator_addr_sh string
			filler_addr string
			filler_addr_sh string
		)
		err=rows.Scan(
			&rec.MktAid,
			&rec.MktType,
			&rec.OutcomeIdx,
			&outcomes,
			&rec.MktDescr,
			&rec.MktAddr,
			&rec.MktAddrSh,
			&rec.WalletAddr,
			&rec.WalletAddrSh,
			&rec.EOAAddr,
			&rec.EOAAddrSh,
			&rec.Date,
			&rec.Timestamp,
			&rec.NetPosition,
			&rec.AvgPrice,
			&rec.FrozenFunds,
			&rec.RealizedProfit,
			&rec.RealizedCost,
			&rec.FinalProfit,
			&order_hash,
			&block_num,
			&creator_eoa_aid,
			&filler_eoa_aid,
			&creator_addr,
			&creator_addr_sh,
			&filler_addr,
			&filler_addr_sh,
		)
		if err!=nil {
			Fatalf(fmt.Sprintf("DB error: %v eoa_aid=%v q=%v",err,eoa_aid,query))
		}
		rec.OutcomeStr = get_outcome_str(uint8(rec.MktType),rec.OutcomeIdx,&outcomes)
		if open_positions {
			rec.FinalProfit = rec.FinalProfit - rec.RealizedProfit
			accumulator = accumulator + rec.FrozenFunds
			rec.AccumFrozen = accumulator
		} else {
			accumulator = accumulator + rec.FinalProfit
			rec.AccumPl = accumulator
		}

		if order_hash.Valid { rec.OrderHash = order_hash.String }
		if block_num.Valid { rec.BlockNum = block_num.Int64 }

		if eoa_aid == creator_eoa_aid {
			rec.CounterPAddr = filler_addr
			rec.CounterPAddrSh = filler_addr_sh
		}
		if eoa_aid == filler_eoa_aid {
			rec.CounterPAddr = creator_addr
			rec.CounterPAddrSh = creator_addr_sh
		}
		fmt.Printf("rec = %+v\n",rec)
		records = append(records,rec)
	}
	return records
}
func (ss *SQLStorage) locate_fill_event_order(evt *FillEvt) int64 {

	var id int64 = 0
	var query string
	query = "SELECT id FROM mktord WHERE order_id = $1"

	h:=hex.EncodeToString(evt.OrderHash[:])
	row := ss.db.QueryRow(query,h)
	var null_id sql.NullInt64
	var err error
	err=row.Scan(&null_id);
	if (err!=nil) {
		if err == sql.ErrNoRows {
			fmt.Printf("pl_calc: order with hash %v wasn't found\n",h)
			// break
		} else {

			Fatalf("DB Error: %v, q=%v\n",err,query);
		}
	} else {
		if null_id.Valid {
			id = null_id.Int64
		}
	}
	return id
}
func (ss *SQLStorage) get_ranking_data_for_all_users() []RankStats {

	var query string
	query = "SELECT eoa_aid,total_trades,profit_loss FROM ustats"

	rows,err := ss.db.Query(query)
	if (err!=nil) {
		Fatalf("DB error: %v (query=%v)",err,query);
	}
	records := make([]RankStats,0,8)
	defer rows.Close()
	for rows.Next() {
		var rec RankStats
		err=rows.Scan(&rec.EoaAid,&rec.TotalTrades,&rec.ProfitLoss)
		records = append(records,rec)
	}
	return records
}
func (ss *SQLStorage) update_top_profit_rank(eoa_aid int64,value float64) int64 {

	var query string
	query = "UPDATE uranks SET top_profit = $2 WHERE eoa_aid = $1"
	res,err:=ss.db.Exec(query,eoa_aid,value)
	if (err!=nil) {
		Fatalf("update_top_profit_rank() failed: %v, q=%v",err,query);
	}
	affected_rows,err:=res.RowsAffected()
	if err!=nil {
		Fatalf("Error getting RowsAffected in update_top_profit(): %v",err)
	}
	if affected_rows == 0 {
		query = "INSERT INTO uranks(eoa_aid,top_profit) VALUES($1,$2)"
		_,err:=ss.db.Exec(query,eoa_aid,value)
		if (err!=nil) {
			Fatalf("update_top_profit_rank() failed: %v, q=%v",err,query);
		}

	}
	return affected_rows
}
func (ss *SQLStorage) update_top_total_trades_rank(eoa_aid int64,value float64) int64 {

	var query string
	query = "UPDATE uranks SET top_trades = $2 WHERE eoa_aid = $1"
	res,err:=ss.db.Exec(query,eoa_aid,value)
	if (err!=nil) {
		Fatalf("update_top_total_trades_rank() failed: %v, q=%v",err,query);
	}
	affected_rows,err:=res.RowsAffected()
	if err!=nil {
		Fatalf("Error getting RowsAffected in update_top_tra(): %v",err)
	}
	if affected_rows == 0 {
		query = "INSERT INTO uranks(eoa_aid,top_trades) VALUES($1,$2)"
		_,err:=ss.db.Exec(query,eoa_aid,value)
		if (err!=nil) {
			Fatalf("update_top_total_trades_rank() failed: value=%v, err: %v, q=%v",value,err,query);
		}

	}
	return affected_rows
}
func (ss *SQLStorage) get_order_info(order_hash string) (OrderInfo,error) {

	var order OrderInfo
	var query string
	query = "SELECT " +
				"o.order_id," +
				"CONCAT(LEFT(o.order_id,6),'…',RIGHT(o.order_id,6)) as order_hash_sh," +
				"s_w_a.addr AS s_w_a_addr," +
				"CONCAT(LEFT(s_w_a.addr,6),'…',RIGHT(s_w_a.addr,6)) AS seller_wallet_addr_sh," +
				"s_e_a.addr AS seller_eoa_addr," +
				"CONCAT(LEFT(s_e_a.addr,6),'…',RIGHT(s_e_a.addr,6)) AS seller_eoa_addr_sh," +
				"b_w_a.addr AS b_w_a_addr," +
				"CONCAT(LEFT(b_w_a.addr,6),'…',RIGHT(b_w_a.addr,6)) AS buyer_wallet_addr_sh," +
				"b_e_a.addr AS byer_eoa_addr," +
				"CONCAT(LEFT(b_e_a.addr,6),'…',RIGHT(b_e_a.addr,6)) AS buyer_eoa_addr_sh," +
				"CASE o.otype " +
					"WHEN 0 THEN 'BID' " +
					"ELSE 'ASK' " +
				"END AS dir, " +
				"o.time_stamp::date AS date," +
				"FLOOR(EXTRACT(EPOCH FROM o.time_stamp))::BIGINT as created_ts," +
				"o.outcome," +
				"o.price AS price, " +
				"o.amount_filled AS volume, " +
				"m.outcomes AS outcomes_str, " +
				"ma.addr, " +
				"CONCAT(LEFT(ma.addr,6),'…',RIGHT(ma.addr,6)) as market_addr_sh " +
			"FROM " +
				"mktord AS o " +
					"LEFT JOIN address AS a ON o.market_aid=a.address_id " +
					"LEFT JOIN address AS s_w_a ON o.wallet_aid=s_w_a.address_id " +
					"LEFT JOIN address AS s_e_a ON o.eoa_aid=s_e_a.address_id " +
					"LEFT JOIN address AS b_w_a ON o.wallet_fill_aid=b_w_a.address_id " +
					"LEFT JOIN address AS b_e_a ON o.eoa_fill_aid=b_e_a.address_id, " +
				"market AS m " +
					"LEFT JOIN address AS ma ON m.market_aid  = ma.address_id " +
			"WHERE (m.market_aid=o.market_aid) AND (o.order_id = $1)"

	var outcomes string
	err:=ss.db.QueryRow(query,order_hash).Scan(
		&order.OrderHash,
		&order.OrderHashSh,
		&order.CreatorrWalletAddr,
		&order.CreatorWalletAddrSh,
		&order.CreatorEOAAddr,
		&order.CreatorEOAAddrSh,
		&order.FillerWalletAddr,
		&order.FillerWalletAddrSh,
		&order.FillerEOAAddr,
		&order.FillerEOAAddrSh,
		&order.OType,
		&order.Date,
		&order.CreatedTs,
		&order.OutcomeIdx,
		&order.Price,
		&order.Volume,
		&outcomes,
		&order.MarketAddr,
		&order.MarketAddrSh,
	);
	if (err!=nil) {
		if (err==sql.ErrNoRows) {
			return order,err
		} else {
			Fatalf("DB error looking up for Order record: %v",err);
		}
	}
	return order,nil
}
/* DISCONTINUED function. ToDo: remove in 2 weeks (29 May)
func (ss *SQLStorage) link_pl_to_order(mktord_id int64, profit_loss_ids *[]int64) {

	var query string
	var pl_ids_str string
	for i:=0 ; i<len(*profit_loss_ids); i++  {
		if len(pl_ids_str) > 0 {
			pl_ids_str = pl_ids_str + ","
		}
		pl_ids_str = pl_ids_str + fmt.Sprintf("%v",(*profit_loss_ids)[i])
	}

	fmt.Printf("q=UPDATE profit_loss SET mktord_id = %v WHERE id IN(%v)\n",mktord_id,pl_ids_str)
	query = "UPDATE profit_loss SET mktord_id = $1 WHERE id IN("+pl_ids_str+")"
	_,err:=ss.db.Exec(query,mktord_id)
	if (err!=nil) {
		Fatalf("DB error: %v ; q=%v",err,query);
	}
}
*/
/* PENDING FOR REMOVAL
	var query string
	query = "SELECT " +
				"pl.market_aid," +
				"m.market_type, " +
				"pl.outcome_idx," +
				"m.outcomes," +
				"substring(extra_info::json->>'description',1,100) as descr," +
				"a.addr as mkt_addr," +
				"CONCAT(LEFT(a.addr,6),'…',RIGHT(a.addr,6)) AS mkt_addr_sh," +
				"w_a.addr AS w_a_addr," +
				"CONCAT(LEFT(w_a.addr,6),'…',RIGHT(w_a.addr,6)) AS wallet_addr_sh," +
				"e_a.addr AS eoa_addr," +
				"CONCAT(LEFT(e_a.addr,6),'…',RIGHT(e_a.addr,6)) AS eoa_addr_sh," +
				"pl.time_stamp::date AS date," +
				"FLOOR(EXTRACT(EPOCH FROM pl.time_stamp))::BIGINT as created_ts," +
				"pl.net_position," +
				"pl.avg_price," +
				"pl.frozen_funds," +
				"pl.realized_profit," +	
				"pl.final_profit," +
				"pl.realized_cost," +
				"o.order_id," +
				"o.block_num " +
			"FROM profit_loss AS pl " +
				"LEFT JOIN address AS a ON pl.market_aid=a.address_id " +
				"LEFT JOIN address AS w_a ON pl.wallet_aid=w_a.address_id " +
				"LEFT JOIN address AS e_a ON pl.eoa_aid=e_a.address_id " +
				"LEFT JOIN market AS m ON pl.market_aid = m.market_aid " +
				"LEFT JOIN mktord AS o ON pl.mktord_id = o.id "+
			"WHERE (pl.eoa_aid = $1) AND (pl.closed_position = 1) " +
			"ORDER BY pl.time_stamp"
//	fmt.Printf("q=%v\n",query)
	rows,err := ss.db.Query(query,eoa_aid)
	if (err!=nil) {
		Fatalf("DB error: %v (query=%v)",err,query);
	}
	records := make([]PLEntry,0,8)
	var starting_point PLEntry
	records = append(records,starting_point)
	var accum_pl float64 = 0.0
	defer rows.Close()
	for rows.Next() {
		var rec PLEntry
		var outcomes string
		var order_hash sql.NullString
		var block_num sql.NullInt64
		err=rows.Scan(
			&rec.MktAid,
			&rec.MktType,
			&rec.OutcomeIdx,
			&outcomes,
			&rec.MktDescr,
			&rec.MktAddr,
			&rec.MktAddrSh,
			&rec.WalletAddr,
			&rec.WalletAddrSh,
			&rec.EOAAddr,
			&rec.EOAAddrSh,
			&rec.Date,
			&rec.Timestamp,
			&rec.NetPosition,
			&rec.AvgPrice,
			&rec.FrozenFunds,
			&rec.RealizedProfit,
			&rec.FinalProfit,
			&rec.RealizedCost,
			&order_hash,
			&block_num,
		)
		if err!=nil {
			Fatalf(fmt.Sprintf("DB error: %v q=%v",err,query))
		}
		rec.OutcomeStr = get_outcome_str(uint8(rec.MktType),rec.OutcomeIdx,&outcomes)
		rec.FinalProfit = rec.FinalProfit - rec.RealizedProfit
		accum_pl = accum_pl + rec.FinalProfit
		rec.AccumPl = accum_pl
		if order_hash.Valid { rec.OrderHash = order_hash.String }
		if block_num.Valid { rec.BlockNum = block_num.Int64 }
		records = append(records,rec)
	}
	return records
}
*/
