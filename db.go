package main

import (
	"os"
	"fmt"
	"net"
	"database/sql"
	_  "github.com/lib/pq"
)

// db const
const (
	host     = "localhost"
	port     = 5432
	user     = "aedev"
	password = "123"
	dbname   = "dev"
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
	row := db.QueryRow("SELECT now()")
	var now string
	err=row.Scan(&now);
	if (err!=nil) {
		show_connect_error()
		os.Exit(1)
	} else {
	}

	ss := new(SQLStorage)
	ss.db = db
	return ss
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
func (ss *SQLStorage) lookup_address(addr string) int64 {

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
func (ss *SQLStorage) insert_market_created_evt(evt *MarketCreatedEvt) {

	var query string
	var market_aid int64;
	market_aid = ss.lookup_or_create_address(evt.Market.String())
	// check if Market is already registered
	query = "SELECT market_aid FROM market WHERE market_aid=$1"
	err:=ss.db.QueryRow(query,market_aid).Scan(&market_aid);
	if (err!=nil) {
		if (err==sql.ErrNoRows) {
			// break
		} else {
			Fatalf("DB error: %v",err)
		}
	} else {
		// market already registered, sliently exit
		return
	}
	universe_id := ss.lookup_universe_id(evt.Universe.String())
	creator_aid := ss.lookup_or_create_address(evt.MarketCreator.String())
	reporter_aid := ss.lookup_or_create_address(evt.DesignatedReporter.String())

	prices := bigint_ptr_slice_to_str(&evt.Prices,",")
	outcomes := outcomes_to_str(&evt.Outcomes,",")
	query = `
		INSERT INTO market(
			universe_id,
			market_aid,
			creator_aid,
			reporter_aid,
			end_time,
			max_ticks,
			time_stamp,
			fee,
			prices,
			market_type,
			extra_info,
			outcomes,
			no_show_bond
		) VALUES ($1,$2,$3,$4,$5,$6,$7,$8,$9,$10,$11,$12,$13)`
	result,err := ss.db.Exec(query,
			universe_id,
			market_aid,
			creator_aid,
			reporter_aid,
			evt.EndTime.Int64(),
			evt.NumTicks.Int64(),
			evt.Timestamp.Int64(),
			evt.FeePerCashInAttoCash.String(),
			prices,
			evt.MarketType,
			evt.ExtraInfo,
			outcomes,
			evt.NoShowBond.String())
	if err != nil {
		Fatalf("DB error: can't insert into market table: %v",err)
	}
	rows_affected,err:=result.RowsAffected()
	if err != nil {
		Fatalf("DB error: %v",err)
	}
	if rows_affected > 0 {
		return
	} else {
		Fatalf("DB error: couldn't insert into Market table. Rows affeced = 0")
	}
}
func (ss *SQLStorage) insert_market_oi_changed_evt(evt *MarketOIChangedEvt,market_addr string) {
	// Note: this event arrives with evt.Market set to 0x0000000000000000000000000 (a contract bug?) ,
	//			so we pass the market address as parameter ('market_addr') to the function
	var query string
	market_aid := ss.lookup_address(market_addr)
	universe_id := ss.lookup_universe_id(evt.Universe.String())

	query = "UPDATE market SET open_interest = $3 WHERE universe_id = $1 AND market_aid = $2"
	_,err := ss.db.Exec(query,universe_id,market_aid,evt.MarketOI.String())
	if err != nil {
		Fatalf("DB error: can't update open interest of market %v : %v",market_aid,err)
	}
	fmt.Printf("Set market %v open interst to %v",market_aid,evt.MarketOI.String())
}
func (ss *SQLStorage) insert_market_order_evt(evt *MktOrderEvt) {
/*
	var query string
	var creator_aid int64;
	creator_aid = ss.lookup_or_create_address(evt.AddressData[0].String())
	var filler_aid int64 = 0;
	if len(evt.AddressData) > 1 {
		filler_aid = ss.lookup_or_create_address(evt.AddressData[1].String())
	}
	universe_id := ss.lookup_universe_id(evt.Universe.String())
	creator_aid := ss.lookup_or_create_address(evt.MarketCreator.String())
	reporter_aid := ss.lookup_or_create_address(evt.DesignatedReporter.String())

	prices := bigint_ptr_slice_to_str(&evt.Prices,",")
	outcomes := outcomes_to_str(&evt.Outcomes,",")
	query = `
		INSERT INTO market(
			universe_id,
			market_aid,
			creator_aid,
			reporter_aid,
			end_time,
			max_ticks,
			time_stamp,
			fee,
			prices,
			market_type,
			extra_info,
			outcomes,
			no_show_bond
		) VALUES ($1,$2,$3,$4,$5,$6,$7,$8,$9,$10,$11,$12,$13)`
	result,err := ss.db.Exec(query,
			universe_id,
			market_aid,
			creator_aid,
			reporter_aid,
			evt.EndTime.Int64(),
			evt.NumTicks.Int64(),
			evt.Timestamp.Int64(),
			evt.FeePerCashInAttoCash.String(),
			prices,
			evt.MarketType,
			evt.ExtraInfo,
			outcomes,
			evt.NoShowBond.String())
	if err != nil {
		Fatalf("DB error: can't insert into market table: %v",err)
	}
	rows_affected,err:=result.RowsAffected()
	if err != nil {
		Fatalf("DB error: %v",err)
	}
	if rows_affected > 0 {
		return
	} else {
		Fatalf("DB error: couldn't insert into Market table. Rows affeced = 0")
	}
*/
}
