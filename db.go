package main

import (
	"os"
	"fmt"
	"net"
	"math/big"
	"strings"
	"strconv"
	"log"
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
	db			*sql.DB
	db_logger	*log.Logger
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
		Fatalf(fmt.Sprintf("DB Error: %v",err));
	}

	ss := new(SQLStorage)
	ss.db = db
	return ss
}
func (ss *SQLStorage) init_log(fname string) {

	f, err := os.OpenFile(fname,os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	check(err)
	ss.db_logger = log.New(f,"DB: ",log.LstdFlags)
}
func (ss *SQLStorage) log_msg(msg string) {
	if ss.db_logger !=nil {
		ss.db_logger.Printf(msg)
	} else {
		fmt.Printf(msg)
	}
}
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
			ss.log_msg(fmt.Sprintf("Error in check_main_stats(): %v, q=%v",err,query))
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
			ss.log_msg(fmt.Sprintf("Error in get_last_block_num(): %v",err))
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
		ss.log_msg(fmt.Sprintf("set_last_block_num() failed: %v",err))
	}
	affected_rows,err:=res.RowsAffected()
	if err!=nil {
		ss.log_msg(fmt.Sprintf("Error getting RowsAffected in set_last_block(): %v",err))
	}
	if affected_rows>0 {
		// break
	} else {
		query = "INSERT INTO last_block VALUES($1)"
		_,err := ss.db.Exec(query,bnum)
		if (err!=nil) {
			ss.log_msg(fmt.Sprintf("set_last_block_num() failed on INSERT: %v",err));
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
			ss.log_msg(fmt.Sprintf("DB error: Universe doesn't exist (addr=%v). Database wasn't initialized correctly",addr))
		} else {
			ss.log_msg(fmt.Sprintf("DB error looking up for Universe record: %v",err))
		}
	}
	return universe_id
}
func (ss *SQLStorage) lookup_eoa_aid(wallet_aid int64) (int64,error) {

	var addr_id int64;
	var query string
	query="SELECT eoa_aid FROM ustats WHERE wallet_aid=$1"
	err:=ss.db.QueryRow(query,wallet_aid).Scan(&addr_id);
	if (err!=nil) {
		if err != sql.ErrNoRows {
			Info.Printf("lookup_eoa_aid(wallet_aid=%v) sql error=%v\n",wallet_aid,err)
		}
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
			ss.log_msg(fmt.Sprintf("DB error: %v, q=%v",err,query))
		} else {
			Info.Printf("lookup_wallet_aid(%v) error: %v\n",eoa_aid,err)
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
			ss.log_msg(fmt.Sprintf("DB error: address %v does not exist",addr))
		} else {
			ss.log_msg(fmt.Sprintf("DB error upon address lookup: %v",err))
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
				ss.log_msg(fmt.Sprintf("DB error in address insertion: %v : %v",query,err))
			}
			if addr_id==0 {
				ss.log_msg(fmt.Sprintf("DB error, addr_id after INSERT is 0"))
			}
			return addr_id
		}
	}
	if (err!=nil) {
		ss.log_msg(fmt.Sprintf("DB error in getting address id : %v",err))
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
				ss.log_msg(fmt.Sprintf("DB error in category insertion: %v : %v",query,err))
			}
			if cat_id==0 {
				ss.log_msg(fmt.Sprintf("DB error, cat_id after INSERT is 0"))
			}
			return cat_id
		}
	}
	if (err!=nil) {
		ss.log_msg(fmt.Sprintf("DB error in getting category id : %v",err))
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
			ss.log_msg(fmt.Sprintf("DB error: %v, q=%v",err,query))
		}
	} else {
		// market already registered, sliently exit
		return
	}
	creator_aid := ss.lookup_or_create_address(evt.MarketCreator.String())
	universe_id := ss.lookup_universe_id(evt.Universe.String())
	eoa_aid := ss.lookup_or_create_address(evt.MarketCreator.String())
	reporter_aid := ss.lookup_or_create_address(evt.DesignatedReporter.String())
	Info.Printf("create_market: signer_aid = %v (%v), creator_aid=%v (%v), reporter_id=%v (%v) , wallet_aid =%v\n",
				signer_aid,signer.String(),
				creator_aid,evt.MarketCreator.String(),
				reporter_aid,evt.DesignatedReporter.String(),
				wallet_aid,
			)
	if signer_aid == creator_aid { // a case only seen in Test environment, production check pending
		// Normally signer != creator, but this happens only in Dev (local testnet), so we have to fix it
		//creator_aid = wallet_aid // this doesn't work, if starting blockchain from block 0, wallt isn't created yet
		wallet_aid = creator_aid
		Info.Printf("create_market: fixed creator id to contract address %v (wallet_aid)\n",wallet_aid)
	} else {
		eoa_aid = signer_aid
		wallet_aid = creator_aid
	}
	if wallet_aid == 0 {
		ss.log_msg(fmt.Sprintf("insert_market_created_evt(): creator addr = %v, wallet_aid = 0, can't continue, exiting\n",
					evt.MarketCreator.String()))
	}
	prices := bigint_ptr_slice_to_str(&evt.Prices,",")
	outcomes := outcomes_to_str(&evt.Outcomes,",")

	var extra_info ExtraInfo
	json.Unmarshal([]byte(evt.ExtraInfo), &extra_info)
	Info.Printf("extra_info unmarshalled: %+v\n",extra_info)
	categories := strings.Join(extra_info.Categories,",")
	Info.Printf("market_categories: %v\n",categories)
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
						"/1e+16,$12,$13,$14,$15,$16)";

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
		ss.log_msg(fmt.Sprintf("DB error: can't insert into market table: %v: q=%v",err,query))
	}
	rows_affected,err:=result.RowsAffected()
	if err != nil {
		ss.log_msg(fmt.Sprintf("DB error: %v, %q",err,query))
	}
	if rows_affected > 0 {
	} else {
		ss.log_msg(fmt.Sprintf("DB error: couldn't insert into Market table. Rows affeced = 0"))
	}
	if len(outcomes) == 0 {
		Info.Printf("len(outcomes)=0\n")
		if evt.MarketType == 0 {	// Yes / No
			outcomes = "Invalid,No,Yes"
		}
		if evt.MarketType == 2 {	// Scalar
			outcomes = "Invalid,,Scalar"
		}
	}
	Info.Printf("init_market_outcome_volumes() outcomes=%v, mkt type = %v\n",outcomes,evt.MarketType)
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
			Info.Printf("insert into outcome volumes query: %v\n",d_query)
			_,err := ss.db.Exec(query,market_aid,outcome_idx)
			if (err!=nil) {
				ss.log_msg(fmt.Sprintf("DB error: %v; q=%v",err,query))
			}
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
		ss.log_msg(fmt.Sprintf("DB error: can't insert into oi_chg table: %v; q=%v",err,query))
	}
	rows_affected,err:=result.RowsAffected()
	if err != nil {
		ss.log_msg(fmt.Sprintf("DB error: %v, q=%v",err,query))
	}
	if rows_affected > 0 {
		return
	} else {
		ss.log_msg(fmt.Sprintf("DB error: couldn't insert into oi_chg table. Rows affeced = 0"))
	}

	Info.Printf("Set market (id=%v) open interest to %v",market_aid,evt.MarketOI.String())
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
		Info.Printf("DB error: couldn't delete open order with order_id = %v\n",order_id)
	}

	Info.Printf("OrderAction = %v, otype=%v, order_id=%v\n",oaction,otype,order_id)
	Info.Printf("Filling existing order %v\n",order_id)
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
		ss.log_msg(fmt.Sprintf("DB error: can't insert into mktord table: %v, q=%v",err,query))
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
		ss.log_msg(fmt.Sprintf("DB error: %v ; q=%v",err,query))
	}
	affected_rows,err:=res.RowsAffected()
	if err!=nil {
		ss.log_msg(fmt.Sprintf("DB error in rows affected: %v, q=%v",err,query))
	}
	if affected_rows == 0 {
		Error.Printf("Last price for market_aid = %v and outcome_idx = %v wasn't updated",
					market_aid,outcome_idx)
	}
}
func (ss *SQLStorage) insert_open_order(evt *zeroex.OrderEvent,eoa_addr string,ospec *ZxMeshOrderSpec) {
	// Insert an open order, this order needs to be Filled by another market participant
	// It also can be canceled by its creator (with another transaction)
	order := evt.SignedOrder.Order
	ohash,err := order.ComputeOrderHash()
	if err != nil {
		Error.Printf("Error at computing 0x Mesh order: %v",err)
	}
	order_id := ohash.String()
	evt_timestamp := evt.Timestamp.Unix()
	expiration := order.ExpirationTimeSeconds.Int64()
	wallet_aid := ss.lookup_or_create_address(order.MakerAddress.String())
	eoa_aid := ss.lookup_or_create_address(eoa_addr)
	Info.Printf("creating open order made by %v : %+v\n",eoa_addr,ospec)
	market_aid := ss.lookup_address_id(ospec.Market.String())
	price := float64(ospec.Price.Int64())/100
	otype := ospec.Type	// Bid/Ask
	amount := order.MakerAssetAmount.String()

	var query string
	query = "INSERT INTO oostats(market_aid,eoa_aid,outcome_idx) VALUES($1,$2,$3)"
	_,err = ss.db.Exec(query,market_aid,eoa_aid,ospec.Outcome)
	if err != nil {
		ss.log_msg(fmt.Sprintf("DB error: %v, q=%v\n",err,query))
	}
	query = "INSERT INTO oorders(" +
				"market_aid,otype,wallet_aid,eoa_aid,price,amount,outcome_idx," +
				"evt_timestamp,srv_timestamp,expiration,order_id" +
			") VALUES($1,$2,$3,$4,$5,"+amount+"/1e+16,$6,TO_TIMESTAMP($7),TO_TIMESTAMP($8),NOW(),$9)"
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
		ss.log_msg(fmt.Sprintf("DB error: can't insert into open orders table: %v, q=%v",err,query))
		return
	}
	rows_affected,err:=result.RowsAffected()
	if err != nil {
		ss.log_msg(fmt.Sprintf("DB error: %v, q=%v",err,query))
	}
	if rows_affected > 0 {
		return
	} else {
		ss.log_msg(fmt.Sprintf("DB error: couldn't insert into Open Orders table. Rows affeced = 0"))
	}
}
func (ss *SQLStorage) delete_open_0x_order(order_hash string) {

	var query string
	query = "DELETE FROM oorders WHERE order_id = $1"
	result,err := ss.db.Exec(query,order_hash)
	if err!=nil {
		ss.log_msg(fmt.Sprintf("DB error: couldn't delete open order with order_id = %v, q=%v\n",order_hash,query))
	}
	rows_affected,err:=result.RowsAffected()
	if rows_affected == 0  {
		ss.log_msg(fmt.Sprintf("DB error: couldn't delete open order with order_id = %v (not found)\n",order_hash))
	}
}
func get_outcome_idx_from_numerators(mkt_type int,numerators []*big.Int) int {

	if mkt_type == 2 {
		return 2
	}
	hundred := big.NewInt(100)
	for i:=0 ; i < len(numerators) ; i++ {
		if hundred.Cmp(numerators[i]) == 0 {
			return i
		}
	}
	return -1
}
func (ss *SQLStorage) insert_market_finalized_evt(evt *MktFinalizedEvt) {

	var query string

	universe_id := ss.lookup_universe_id(evt.Universe.String())
	_ = universe_id	// ToDo: add universe_id match condition (for market)
	market_aid := ss.lookup_address_id(evt.Market.String())
	fin_timestamp := evt.Timestamp.Int64()
	winning_payouts := bigint_ptr_slice_to_str(&evt.WinningPayoutNumerators,",")

	market_type := ss.get_market_type(market_aid)
	winning_outcome := get_outcome_idx_from_numerators(market_type,evt.WinningPayoutNumerators)

	query = "INSERT INTO mkt_fin(market_aid,fin_timestamp,winning_payouts,winning_outcome)" +
			"VALUES($1,TO_TIMESTAMP($2),$3,$4)"
	_,err := ss.db.Exec(query,market_aid,fin_timestamp,winning_payouts,winning_outcome)
	if err != nil {
		ss.log_msg(fmt.Sprintf("DB error: can't update market finalization of market %v : %v, q=%v",market_aid,err,query))
	}
	ss.update_market_status(market_aid,MktStatusFinalized)
	ss.update_losing_positions(market_aid,evt)
}
func (ss *SQLStorage) get_market_type(market_aid int64) int {

	var query string
	query = "SELECT market_type FROM market WHERE market_aid=$1"

	var market_type int
	err:=ss.db.QueryRow(query,market_aid).Scan(&market_type);
	if (err!=nil) {
		ss.log_msg(fmt.Sprintf("DB Error: %v, q=%v\n",err,query))
		Fatalf("get_market_type() failed, market not found")
	}
	return market_type
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
		ss.log_msg(fmt.Sprintf("DB error: %v ; q=%v",err,query))
	}
	affected_rows,err:=res.RowsAffected()
	if err!=nil {
		ss.log_msg(fmt.Sprintf("DB error: %v ; q=%v",err,query))
	}
	if affected_rows == 0 {
		Error.Printf("Couldn't update market status = %v for market %v",status,market_aid)
	} else {
		Info.Printf("MKTSTATUS: market_aid = %v, status = %v\n",market_aid,status)
	}
}
func (ss *SQLStorage) update_losing_positions(market_aid int64,evt *MktFinalizedEvt) {

	// this function marks losing positions as closed (because we don't have ProfitLoss event
	//			on a losing position (position with wrong outcome)
	var query string
	query = "SELECT market_type FROM market WHERE market_aid=$1"

	/*
	discontinued, to be delted
	var market_type int
	err:=ss.db.QueryRow(query,market_aid).Scan(&market_type);
	if (err!=nil) {
		if (err==sql.ErrNoRows) {
			return
		}
		ss.log_msg(fmt.Sprintf("DB Error: %v, q=%v\n",err,query))
	}
	*/
	market_type:=ss.get_market_type(market_aid)

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
		ss.log_msg(fmt.Sprintf("DB error: %v ; q=%v",err,query))
	}
	affected_rows,err:=res.RowsAffected()
	if err!=nil {
		ss.log_msg(fmt.Sprintf("DB error in rows affected: %v",err))
	}
	Info.Printf("Market finalized. amount of closed losing positions: %v\n",affected_rows)
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

	Info.Printf("insert_initial_report_evt(): market_aid=%v, reporter_id=%v, signer_aid=%v\n",
					market_aid,reporter_aid,signer_aid)

	market_type := ss.get_market_type(market_aid)
	reported_outcome := get_outcome_idx_from_numerators(market_type,evt.PayoutNumerators)

	var query string
	query = `
		INSERT INTO report (
			block_num,
			tx_id,
			market_aid,
			wallet_aid,
			eoa_aid,
			ini_reporter_aid,
			outcome_idx,
			is_initial,
			is_designated,
			amount_staked,
			pnumerators,
			description,
			next_win_start,
			next_win_end,
			rpt_timestamp
		) VALUES (
			$1,$2,$3,$4,$5,$6,$7,$8,$9,(` + amount_staked + `/1e+18),$10,$11,
			TO_TIMESTAMP($12),
			TO_TIMESTAMP($13),
			TO_TIMESTAMP($14)
		)`
	result,err := ss.db.Exec(query,
			block_num,
			tx_id,
			market_aid,
			reporter_aid,
			signer_aid,
			ini_reporter_aid,
			reported_outcome,
			true,
			evt.IsDesignatedReporter,
			payout_numerators,
			evt.Description,
			next_win_start,
			next_win_end,
			rpt_timestamp)
	if err != nil {
		ss.log_msg(fmt.Sprintf("DB error: can't insert into report table: %v,q=%v",err,query))
	}
	rows_affected,err:=result.RowsAffected()
	if err != nil {
		ss.log_msg(fmt.Sprintf("DB error: %v, q=%v",err,query))
	}
	if rows_affected > 0 {
		//break
	} else {
		ss.log_msg(fmt.Sprintf("DB error: couldn't insert into InitialReport table. Rows affeced = 0"))
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
		ss.log_msg(fmt.Sprintf("DB error: can't insert into volume table: %v, q=%v",err,query))
	}
	rows_affected,err:=result.RowsAffected()
	if err != nil {
		ss.log_msg(fmt.Sprintf("DB error: %v, q=%v",err,query))
	}
	if rows_affected > 0 {
		//break
	} else {
		ss.log_msg(fmt.Sprintf("DB error: couldn't insert into InitialReport table. Rows affeced = 0"))
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
			ss.log_msg(fmt.Sprintf("DB error: %v ; q=%v",err,query))
		}
		affected_rows,err:=res.RowsAffected()
		if err!=nil {
			ss.log_msg(fmt.Sprintf("DB error in rows affected: %v",err))
		}
		if affected_rows>0 {
			// break
		} else {
/* no longer required. to be deleted later
			query = "INSERT INTO outcome_vol(" +
						"market_aid," +
						"outcome_idx," +
						"volume" +
					") VALUES(" +
						"$1," +
						"$2," +
						evt.OutcomeVolumes[outcome_idx].String() + "/1e+18" +
					")"
			d_query := fmt.Sprintf("INSERT INTO outcome_vol(" +
						"market_aid," +
						"outcome_idx," +
						"volume" +
					") VALUES(" +
						"%v," +
						"%v,",market_aid,outcome_idx) +
						evt.OutcomeVolumes[outcome_idx].String() + "/1e+18" +
					")"
					fmt.Printf("insert_market_volume_changed_evt(): query = %v\n",d_query)
			_,err := ss.db.Exec(query,market_aid,outcome_idx)
			if (err!=nil) {
				Fatalf("DB error: %v; q=%v",err,query);
			}
*/
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

	Info.Printf("insert_dispute_crows_contrib(): market_aid=%v, reporter_id=%v, signer_aid=%v",
					market_aid,reporter_aid,signer_aid)

	market_type := ss.get_market_type(market_aid)
	reported_outcome := get_outcome_idx_from_numerators(market_type,evt.PayoutNumerators)

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
			outcome_idx,
			amount_staked,
			description,
			pnumerators,
			current_stake,
			stake_remaining,
			rpt_timestamp
		) VALUES ($1,$2,$3,$4,$5,$6,$7,$8,`+amount_staked+`/1e+18,$9,$10,
				`+cur_stake+`/1e+18,`+stake_remaining+`/1e+18,TO_TIMESTAMP($11))`
	result,err := ss.db.Exec(query,
			block_num,
			tx_id,
			market_aid,
			reporter_aid,
			signer_aid,
			disputed_aid,
			dispute_round,
			reported_outcome,
			evt.Description,
			payout_numerators,
			rpt_timestamp)
	if err != nil {
		ss.log_msg(fmt.Sprintf("DB error: can't insert dispute into report table: %v; q=%v",err,query))
	}
	rows_affected,err:=result.RowsAffected()
	if err != nil {
		ss.log_msg(fmt.Sprintf("DB error: %v, q=%v",err,query))
	}
	if rows_affected == 0 {
		ss.log_msg(fmt.Sprintf("DB error: couldn't insert dispute into Report table. Rows affeced = 0"))
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
		ss.log_msg(fmt.Sprintf("DB error: can't update 'sbalances' for account %v, market %v : %v; q=%v",
					evt.Account.String(),evt.Market.String(),err,query))
	}
	rows_affected,err:=result.RowsAffected()
	if err != nil {
		ss.log_msg(fmt.Sprintf("DB error: %v",err))
	}
	Error.Printf("No error, rows affected = %v\n",rows_affected)
	if rows_affected > 0 {
		Info.Printf("Update to sbalances %v , outcome %v holds %v \n",evt.Account.String(),outcome,balance);
		//break
	} else {
		Info.Printf("Insert to sbalances (%v outcome %v bal=%v\n",evt.Account.String(),outcome,balance);
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
			ss.log_msg(fmt.Sprintf("DB error: can't insert into sbalances table: %v, q=%v",err,query))
		}
		rows_affected,err:=result.RowsAffected()
		if err != nil {
			ss.log_msg(fmt.Sprintf("DB error: %v, query=%v",err,query))
		}
		if rows_affected > 0 {
			return
		} else {
			ss.log_msg(fmt.Sprintf("DB error: couldn't insert into 'sbalances' table. Rows affeced = 0"))
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
					Error.Printf("Block sequence broken after block %v\n",parent_block_num)
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
		ss.log_msg(fmt.Sprintf("DB error: can't insert into block  table: %v, q=%v",err,query))
	}
	rows_affected,err:=result.RowsAffected()
	if err != nil {
		ss.log_msg(fmt.Sprintf("DB error: %v, q=%v",err,query))
	}
	if rows_affected > 0 {
		return true
	}
	ss.log_msg(fmt.Sprintf("DB error: couldn't insert into block table. Rows affeced = 0"))
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
		ss.log_msg(fmt.Sprintf("DB error: can't insert into transactions table: %v, q=%v",err,query))
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
			ss.log_msg(fmt.Sprintf("Chainsplit detected, I don't have the parent hash %v, exiting. ",parent_hash))
		} else {
			ss.log_msg(fmt.Sprintf("DB error: %v, q=%v",err,query))
		}
	}
	cur_block_num := int64(block.Number.Uint64())
	if cur_block_num > (my_block_num + MAX_BLOCKS_CHAIN_SPLIT) {
		ss.log_msg(fmt.Sprintf("Chainsplit detected, and it is more than %v blocks, aborting.",MAX_BLOCKS_CHAIN_SPLIT))
	}
	query = "DELETE FROM block WHERE block_num > $1 CASCADE"
	_,err = ss.db.Exec(query,my_block_num)
	if (err!=nil) {
		ss.log_msg(fmt.Sprintf("DB error: %v, q=%v, block_num=%v",err,query,my_block_num))
	}
	return BlockNumber(my_block_num + 1)	// parent + 1 = current
}
func (ss *SQLStorage) block_delete_with_everything(block_num BlockNumber) {

	// deletes block table and all the other tables receieve cascaded DELETEs also
	var query string
	query = "DELETE FROM block WHERE block_num = $1"
	_,err := ss.db.Exec(query,block_num)
	if (err!=nil) {
		ss.log_msg(fmt.Sprintf("DB error: %v (block_num=%v, %v)",err,block_num,query))
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
			"ORDER BY " +
				"m.market_aid " +
			"OFFSET $1 LIMIT $2";

	rows,err := ss.db.Query(query,off,lim)
	if (err!=nil) {
		ss.log_msg(fmt.Sprintf("DB error: %v (query=%v)",err,query))
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
					&rec.CategoryStr,
					&rec.Outcomes,
					&rec.MktType,
					&rec.MktTypeStr,
					&status_code,
					&rec.Status,
					&rec.Fee,
					&rec.OpenInterest,
					&rec.CurVolume,
		)
		if err!=nil {
			ss.log_msg(fmt.Sprintf("DB error: %v, q=%v",err,query))
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
				"total_markets," +
				"category " +
			"FROM category " +
			"ORDER BY " +
				"category";

	_,err := ss.db.Exec(query)
	if (err!=nil) {
		ss.log_msg(fmt.Sprintf("DB error: %v ,q=%v",err,query))
	}
	rows,err:=ss.db.Query(query)
	if err!=nil {
		if err!=sql.ErrNoRows {
			ss.log_msg(fmt.Sprintf("Error for query %v: %v",query,err))
		}
	}
	var rec InfoCategories
	records := make([]InfoCategories,0,8)

	defer rows.Close()
	for rows.Next() {
		err=rows.Scan(&rec.CatId,&rec.TotalMarkets,&rec.Category)
		if err!=nil {
			ss.log_msg(fmt.Sprintf("DB error: %v",err))
		}
//		fmt.Printf("going to do split of: %+v\n",rec.Category)
/* disabled, DELETION pending
		subcategories := strings.Split(rec.Category,",")
		for i := 0 ; i< len(subcategories); i++ {
			subcategories[i] = strings.Title(subcategories[i])
//			fmt.Printf("added subcategory i=%v, subcat = %v\n",i,subcategories[i])
		}
		if len(subcategories) > 0 {	// sometimes last category is empty, delete it
			if len(subcategories[len(subcategories)-1]) == 0 {
				subcategories = subcategories[:len(subcategories)-1]
			}
		}
		rec.Subcategories = subcategories
*/
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
func (ss *SQLStorage) get_mkt_trades(mkt_addr string,limit int) []MarketTrade {
	// get market trades with mixed outcomes
	Info.Printf("get_mkt_trades() mkt_addr=%v\n",mkt_addr)
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
				"ca.addr as creator_addr," +
				"CONCAT(LEFT(ca.addr,6),'…',RIGHT(ca.addr,6)) AS creator_addr_sh," +
				"fa.addr as filler_addr," +
				"CONCAT(LEFT(fa.addr,6),'…',RIGHT(fa.addr,6)) AS filler_addr_sh," +
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
				"LEFT JOIN address AS a ON o.market_aid=a.address_id " +
				"LEFT JOIN address AS fa ON o.eoa_fill_aid=fa.address_id " +
				"LEFT JOIN address AS ca ON o.eoa_aid=ca.address_id " +
				"LEFT JOIN market AS m ON o.market_aid = m.market_aid " +
			where +
			"ORDER BY o.time_stamp"
	if limit > 0 {
		query = query +	" LIMIT " + strconv.Itoa(limit)
	}

	var rows *sql.Rows
	var err error
	if market_aid > 0 {
		rows,err = ss.db.Query(query,market_aid)
	} else {
		rows,err = ss.db.Query(query)
	}
	if (err!=nil) {
		ss.log_msg(fmt.Sprintf("DB error: %v (query=%v)",err,query))
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
			&rec.CreatorAddr,
			&rec.CreatorAddrSh,
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
			ss.log_msg(fmt.Sprintf("DB error: %v, q=%v",err,query))
		}
		rec.OutcomeStr = get_outcome_str(uint8(mkt_type),rec.Outcome,&outcomes)
		records = append(records,rec)
	}
	return records
}
func (ss *SQLStorage) get_market_info(mkt_addr string,outcome_idx int,oc bool) (InfoMarket,error) {
	// Inputs: 
	//		mkt_addr			address of the market to get the data from
	//		outcome_idx			narrow search by specific outcome
	//		oc					format outcome as string (from the integer parameter in the args)
	var rec InfoMarket
	market_aid,err := ss.nonfatal_lookup_address_id(mkt_addr)
	if err != nil {
		Info.Printf("market %v not found, returning empty data\n",mkt_addr)
		return rec,err
	}
	rec.MktAid=market_aid
	Info.Printf("querying info for market aid = %v\n",market_aid)
	var reporter_aid int64
	var query string
	query = "SELECT " +
				"m.market_type," +
				"ma.addr as mkt_addr," +
				"CONCAT(LEFT(ma.addr,6),'…',RIGHT(ma.addr,6)) AS mkt_addr_sh," +
				"sa.addr AS signer," +
				"CONCAT(LEFT(sa.addr,6),'…',RIGHT(sa.addr,6)) AS signer_sh," +
				"ca.addr as mcreator," +
				"CONCAT(LEFT(ca.addr,6),'…',RIGHT(ca.addr,6)) AS mcreator_sh, " +
				"ra.addr AS reporter,"+
				"CONCAT(LEFT(ra.addr,6),'…',RIGHT(ra.addr,6)) AS reporter_sh," +
				"reporter_aid," +
				"TO_CHAR(end_time,'dd/mm/yyyy HH24:SS UTC') AS end_date," + 
				"extra_info::json->>'description' AS descr," +
				"extra_info::json->>'longDescription' AS long_desc," +
// old version	"extra_info::json->>'categories' AS categories," +
				"cat.category," +
				"outcomes," +
				"m.market_type, " +
				"CASE m.market_type " +
					"WHEN 0 THEN 'YES/NO' " +
					"WHEN 1 THEN 'CATEGORICAL' " +
					"WHEN 2 THEN 'SCALAR' " +
				"END AS mtype, " +
				"CASE m.status " +
					"WHEN 0 THEN 'TRADED' " +
					"WHEN 1 THEN 'REPORTING' " +
					"WHEN 2 THEN 'REPORTED' " +
					"WHEN 3 THEN 'DISPUTING' " +
					"WHEN 4 THEN 'FINALIZED SUCCESSULY' " +
					"WHEN 5 THEN 'FINALIZED INVALID' " +
				"END AS mstatus," +
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
	err=row.Scan(
				&mkt_type,
				&rec.MktAddr,
				&rec.MktAddrSh,
				&rec.Signer,
				&rec.SignerSh,
				&rec.MktCreator,
				&rec.MktCreatorSh,
				&rec.Reporter,
				&rec.ReporterSh,
				&reporter_aid,
				&rec.EndDate,
				&rec.Description,
				&rec.LongDesc,
				&rec.CategoryStr,
				&rec.Outcomes,
				&rec.MktType,
				&rec.MktTypeStr,
				&rec.Status,
				&rec.Fee,
				&rec.OpenInterest,
				&rec.CurVolume,
				&rec.TotalTrades,
				&rec.MoneyAtStake,
	)
	if oc { // get outcome string
		rec.CurOutcome = get_outcome_str(uint8(mkt_type),outcome_idx,&rec.Outcomes)
	}
	if err!=nil {
		if err == sql.ErrNoRows {
			return rec,err
		}
		ss.log_msg(fmt.Sprintf("DB error: %v, q=%v\n",err,query))
		os.Exit(1)
	}
	reporter_eoa_aid,err := ss.lookup_eoa_aid(reporter_aid)
	Info.Printf("reporter_aid = %v,       reporter_eoa_aid=%v\n",reporter_aid,reporter_eoa_aid)
	if err == nil {
		rep_addr,err := ss.lookup_address(reporter_eoa_aid)
		if err == nil {
			Info.Printf("looked up reporter addr = %v\n",rep_addr)
			rec.Reporter = rep_addr
			rec.ReporterSh = string(rep_addr[0:6]+string('…')+rep_addr[26:32])
		}
	}
	subcategories := make_subcategories(&rec.CategoryStr)
	rec.Subcategories = subcategories

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
	d_query := fmt.Sprintf("SELECT " +
				"o.outcome_idx, " +
				"o.volume," +
				"o.last_price, " +
				"m.market_type, " +
				"m.outcomes " +
			"FROM outcome_vol AS o " +
				"LEFT JOIN " +
					"market AS m ON o.market_aid = m.market_aid " +
			"WHERE o.market_aid = %v",market_aid)
	Info.Printf("outcome volumes query: %v\n",d_query)

	var rows *sql.Rows
	rows,err = ss.db.Query(query,market_aid)
	if (err!=nil) {
		ss.log_msg(fmt.Sprintf("DB error: %v (query=%v)",err,query))
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
			ss.log_msg(fmt.Sprintf("DB error: %v, q=%v",err,query))
		}
		rec.OutcomeStr = get_outcome_str(uint8(rec.MktType),rec.Outcome,&outcomes)
		Info.Printf("get_outcome_volumes(): rec.OutcomeStr=%v (extracted from %v)\n",rec.OutcomeStr,outcomes)
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
	Info.Printf("q=%v\n",query)
	var accumulated_volume = 0.0
	rows,err := ss.db.Query(query,market_aid,outc,otype)
	if (err!=nil) {
		ss.log_msg(fmt.Sprintf("DB error: %v (query=%v)",err,query))
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
			ss.log_msg(fmt.Sprintf("DB error: %v, q=%v",err,query))
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
		ss.log_msg(fmt.Sprintf("DB error: %v (query=%v)",err,query))
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
		Info.Printf("record: price = %v, volume = %v\n",rec.Price,rec.Volume)
		if err!=nil {
			ss.log_msg(fmt.Sprintf("DB error: %v",err))
		}
		accumulated_volume = accumulated_volume + rec.Volume
		rec.AccumVol = accumulated_volume
		records = append(records,rec)
	}
	return records
}
func (ss *SQLStorage) get_mkt_depth(mkt_addr string,outcome_idx int) *MarketDepth {

	Info.Printf("get_mkt_depth(mkt=%v,outcome_idx=%v)\n",mkt_addr,outcome_idx)
	market_aid := ss.lookup_address_id(mkt_addr)
	market_depth := new(MarketDepth)
	market_depth.Bids = ss.build_depth_by_otype(market_aid,outcome_idx,OrderTypeBid)
	market_depth.Asks = ss.build_depth_by_otype(market_aid,outcome_idx,OrderTypeAsk)
	return market_depth
}
func (ss *SQLStorage) get_user_info(user_aid int64) (UserInfo,error) {

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
			Info.Printf("No rows for User Info on eoa_aid = %v\n",user_aid)
			return ui,err
		} else {
			ss.log_msg(fmt.Sprintf("DB error: %v, q=%v\n",err,query))
		}
		os.Exit(1)
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
	return ui,nil
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
//	Info.Printf(fmt.Sprintf("main stats = %+v\n",s))
	if (err!=nil) {
		if err == sql.ErrNoRows {
			Info.Printf("No rows for market statistics query\n")
		} else {
			ss.log_msg(fmt.Sprintf("DB error: %v, q=%v\n",err,query))
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
		ss.log_msg(fmt.Sprintf("DB error: %v q=%v",err,query))
	}

}
func (ss *SQLStorage) update_users_profit_loss(market_aid int64,eoa_aid int64,outcome_idx int,realized_profit string) string {

	var err error
	var query string

	// Update previous position status on this outcome
	query = "UPDATE profit_loss " +
				"SET closed_position = 1 " +
//					"final_profit = ($4/1e+36) " +
				"WHERE " +
						"(market_aid = $1) AND " +
						"(eoa_aid = $2) AND " +
						"(outcome_idx = $3) AND " +
						"(closed_position = 0) " +
				"RETURNING realized_profit::text"
	d_query := fmt.Sprintf("UPDATE profit_loss " +
				"SET closed_position = 1 " +
//					"final_profit = (%v/1e+36) " +
				"WHERE " +
						"(market_aid = %v) AND " +
						"(eoa_aid = %v) AND " +
						"(outcome_idx = %v) AND " +
						"(closed_position = 0)",
						realized_profit,market_aid,eoa_aid,outcome_idx)
	Info.Printf("Position update query: %v\n",d_query)

	var previous_profit string
	row:=ss.db.QueryRow(query,market_aid,eoa_aid,outcome_idx)
	err=row.Scan(&previous_profit);
	if err != nil {
		if err == sql.ErrNoRows {
			Info.Printf("pl_calc: notice: no previous trades were closed for this market & outcome\n")
		} else {
			ss.log_msg(fmt.Sprintf("DB error: %v, (on Scan of previous profit) q=%v",err,query))
		}
	}
	return previous_profit
}
func (ss *SQLStorage) insert_profit_loss_evt(block_num BlockNumber,tx_id int64,eoa_aid int64,evt *ProfitLossChanged) int64  {

	var query string
	var err error

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

	prev_profit:=ss.update_users_profit_loss(market_aid,eoa_aid,int(outcome_idx),realized_profit)

	var final_profit string
	if len(prev_profit) > 0 {
		final_profit="((" + realized_profit + "/1e+36)-" + prev_profit + ")"
	} else {
		//final_profit="(" + realized_profit + "/1e+36)"
		final_profit="(0)"
	}
	Info.Printf("Insert to profitloss (wallet %v outcome %v)\n",evt.Account.String(),outcome_idx);
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
				"final_profit," +
				"time_stamp" +
			") VALUES($1,$2,$3,$4,$5,$6,$7," +
				"(" +net_position+ "/1e+16)," +
				"(" +avg_price+ "/1e+20)," +
				"(" +frozen_funds+ "/1e+36)," +
				"(" +realized_profit+ "/1e+36)," +
				"(" +realized_cost+ "/1e+36)," +
				"(" +final_profit+ ")," +
				"TO_TIMESTAMP($8)" +
			") RETURNING id,realized_profit,realized_cost,net_position"

	var null_pl_id sql.NullInt64
	var null_profit sql.NullFloat64
	var null_rcost sql.NullFloat64
	var null_volume sql.NullFloat64
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
	err=row.Scan(&null_pl_id,&null_profit,&null_rcost,&null_volume);
	if (err!=nil) {
		if err == sql.ErrNoRows {
			//
		} else {
			ss.log_msg(fmt.Sprintf("DB error: %v; q=%v",err,query))
		}
	} else {
		pl_id = null_pl_id.Int64
	}
	if null_volume.Valid {
		if null_volume.Float64 == 0 {
			// Volume = 0 means the User has closed all his positions,
			// therefore we must mark position as closed in the DB too
			Info.Printf("Closing position due to 0 volume, realized profit = %v\n")
			ss.update_users_profit_loss(market_aid,eoa_aid,int(outcome_idx),realized_profit)
		}
	}

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
	Info.Printf("pl.eoa_aid=%v; q=%v\n",eoa_aid,query)
	rows,err := ss.db.Query(query,eoa_aid)
	if (err!=nil) {
		ss.log_msg(fmt.Sprintf("DB error: %v (query=%v)",err,query))
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
			ss.log_msg(fmt.Sprintf("DB error: %v eoa_aid=%v q=%v",err,eoa_aid,query))
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
		Info.Printf("rec = %+v\n",rec)
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
			Info.Printf("pl_calc: order with hash %v wasn't found\n",h)
			// break
		} else {

			ss.log_msg(fmt.Sprintf("DB Error: %v, q=%v\n",err,query))
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
		ss.log_msg(fmt.Sprintf("DB error: %v (query=%v)",err,query))
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
		ss.log_msg(fmt.Sprintf("update_top_profit_rank() failed: %v, q=%v",err,query))
	}
	affected_rows,err:=res.RowsAffected()
	if err!=nil {
		ss.log_msg(fmt.Sprintf("Error getting RowsAffected in update_top_profit(): %v",err))
	}
	if affected_rows == 0 {
		query = "INSERT INTO uranks(eoa_aid,top_profit) VALUES($1,$2)"
		_,err:=ss.db.Exec(query,eoa_aid,value)
		if (err!=nil) {
			ss.log_msg(fmt.Sprintf("update_top_profit_rank() failed: %v, q=%v",err,query))
		}

	}
	return affected_rows
}
func (ss *SQLStorage) update_top_total_trades_rank(eoa_aid int64,value float64) int64 {

	var query string
	query = "UPDATE uranks SET top_trades = $2 WHERE eoa_aid = $1"
	res,err:=ss.db.Exec(query,eoa_aid,value)
	if (err!=nil) {
		ss.log_msg(fmt.Sprintf("update_top_total_trades_rank() failed: %v, q=%v",err,query))
	}
	affected_rows,err:=res.RowsAffected()
	if err!=nil {
		ss.log_msg(fmt.Sprintf("Error getting RowsAffected in update_top_tra(): %v",err))
	}
	if affected_rows == 0 {
		query = "INSERT INTO uranks(eoa_aid,top_trades) VALUES($1,$2)"
		_,err:=ss.db.Exec(query,eoa_aid,value)
		if (err!=nil) {
			ss.log_msg(fmt.Sprintf("update_top_total_trades_rank() failed: value=%v, err: %v, q=%v",value,err,query))
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
			ss.log_msg(fmt.Sprintf("DB error looking up for Order record: %v",err))
		}
	}
	return order,nil
}
func (ss *SQLStorage) get_category_markets(cat_id int64) []InfoMarket {

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
		ss.log_msg(fmt.Sprintf("DB error: %v (query=%v)",err,query))
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
					&rec.CategoryStr,
					&rec.Outcomes,
					&rec.MktType,
					&rec.MktTypeStr,
					&status_code,
					&rec.Status,
					&rec.Fee,
					&rec.OpenInterest,
					&rec.CurVolume,
		)
		if err!=nil {
			ss.log_msg(fmt.Sprintf("DB error: %v, q=%v",err,query))
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
func (ss *SQLStorage) get_user_reports(eoa_aid int64,limit int) []UserReport {

	var query string
	query = "SELECT " +
				"r.rpt_timestamp::date," +
				"ma.addr as mkt_addr," +
				"CONCAT(LEFT(ma.addr,6),'…',RIGHT(ma.addr,6)) AS mkt_addr_sh, " +
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

	records := make([]UserReport,0,8)
	var rows *sql.Rows
	var err error
	rows,err = ss.db.Query(query,eoa_aid)
	if (err!=nil) {
		if err == sql.ErrNoRows {
			return records
		}
		ss.log_msg(fmt.Sprintf("DB error: %v (query=%v)",err,query))
		os.Exit(1)
	}

	defer rows.Close()
	for rows.Next() {
		var rec UserReport
		var mkt_type int
		var designated_outcome int
		var winning_outcome int
		var initial_outcome int
		var outcomes string
		err=rows.Scan(
			&rec.Date,
			&rec.MktAddr,
			&rec.MktAddrSh,
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
			ss.log_msg(fmt.Sprintf("DB error: %v, q=%v",err,query))
		}
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
		fmt.Printf("adding record %+v\n",rec)
		records = append(records,rec)
	}
	return records
}
