package main

import (
	"os"
	"fmt"
	"net"
	"database/sql"
	"encoding/hex"
	_  "github.com/lib/pq"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
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
func (ss *SQLStorage) insert_market_created_evt(
		block_num BlockNumber,
		tx_id int64,
		signer common.Address,
		evt *MarketCreatedEvt,
) {

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
	signer_aid := ss.lookup_or_create_address(signer.String())
	fmt.Printf("create_market: creator_aid=%v (%v), reporter_id=%v (%v)\n",
				creator_aid,evt.MarketCreator.String(),reporter_aid,evt.DesignatedReporter.String())

	prices := bigint_ptr_slice_to_str(&evt.Prices,",")
	outcomes := outcomes_to_str(&evt.Outcomes,",")
	query = `
		INSERT INTO market(
			block_num,
			tx_id,
			universe_id,
			market_aid,
			creator_aid,
			signer_aid,
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
		) VALUES ($1,$2,$3,$4,$5,$6,$7,$8,$9,$10,$11,$12,$13,$14,$15,$16)`
	result,err := ss.db.Exec(query,
			block_num,
			tx_id,
			universe_id,
			market_aid,
			creator_aid,
			signer_aid,
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
func (ss *SQLStorage) insert_market_oi_changed_evt(block *types.Header,evt *MarketOIChangedEvt) {
	// Note: this event arrives with evt.Market set to 0x0000000000000000000000000 (a contract bug?) ,
	//			so we pass the market address as parameter ('market_addr') to the function
	var query string
	market_aid := ss.lookup_address(evt.Market.String())
	ts_inserted := int64(block.Time)
	query = "INSERT INTO oi_chg(market_aid,ts_inserted,oi) VALUES($1,$2,$3)"
	result,err := ss.db.Exec(query,market_aid,ts_inserted,evt.MarketOI.String())
	if err != nil {
		Fatalf("DB error: can't insert into oi_chg table: %v",err)
	}
	rows_affected,err:=result.RowsAffected()
	if err != nil {
		Fatalf("DB error: %v",err)
	}
	if rows_affected > 0 {
		return
	} else {
		Fatalf("DB error: couldn't insert into oi_chg table. Rows affeced = 0")
	}

	fmt.Printf("Set market (id=%v) open interest to %v",market_aid,evt.MarketOI.String())
}
func (ss *SQLStorage) insert_market_order_evt(
	block_num BlockNumber,
	tx_id int64,
	signer common.Address,
	evt *MktOrderEvt,
) {

	// depending on the order action (Create/Cancel/Fill) different table is used for storage
	//		Create/Cancel order actions go to 'oorders' (Open Orders) table because these orders
	//		do not alter market's open interest.
	//		Fill order goes to 'mktord' table because the share has been created and now
	//		open interest increased
	var creator_aid int64;
	creator_aid = ss.lookup_or_create_address(evt.AddressData[0].String())
	var filler_aid int64 = 0;
	if len(evt.AddressData) > 1 {
		filler_aid = ss.lookup_or_create_address(evt.AddressData[1].String())
	}
	universe_id := ss.lookup_universe_id(evt.Universe.String())
	_ = universe_id	// ToDo: add universe_id match condition (for market)
	market_aid := ss.lookup_address(evt.Market.String())
	signer_aid := ss.lookup_or_create_address(signer.String())

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
	price := evt.Uint256Data[0].Int64()
	amount := evt.Uint256Data[1].Int64()
	outcome := evt.Uint256Data[2].Int64()
	token_refund := evt.Uint256Data[3].String()
	shares_refund := evt.Uint256Data[4].String()
	fees := evt.Uint256Data[5].String()
	amount_filled := evt.Uint256Data[6].String()
	time_stamp := evt.Uint256Data[7].Int64()
	shares_escrowed := evt.Uint256Data[8].String()
	tokens_escrowed := evt.Uint256Data[9].String()

	var query string
	fmt.Printf("OrderAction = %v, otype=%v, order_id=%v\n",oaction,otype,order_id)
	switch oaction {
		case OrderActionFill:
			fmt.Printf("Filling existing order %v\n",order_id)
			query = `
				INSERT INTO mktord(
					market_aid,
					signer_aid,
					block_num,
					tx_id,
					oaction,
					otype,
					creator_aid,
					filler_aid,
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
				) VALUES ($1,$2,$3,$4,$5,$6,$7,$8,$9,$10,$11,$12,$13,$14,$15,$16,$17,$18,$19,$20)`
			result,err := ss.db.Exec(query,
					market_aid,
					signer_aid,
					block_num,
					tx_id,
					oaction,
					otype,
					creator_aid,
					filler_aid,
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
					hex.EncodeToString(evt.TradeGroupId[:]),
					order_id)
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
		// end of case OrderActionFill
		case OrderActionCreate:
			fmt.Printf("creating open order: %v\n",order_id)
			query = "INSERT INTO oorders(" +
						"market_aid,otype,creator_aid,price,amount,outcome,time_stamp,order_id" +
					") VALUES($1,$2,$3,$4,$5,$6,$7,$8,$9)"
			result,err := ss.db.Exec(query,
					tx_id,
					market_aid,
					otype,
					creator_aid,
					price,
					amount,
					outcome,
					time_stamp,
					order_id)
			if err != nil {
				Fatalf("DB error: can't insert into open orders table: %v",err)
			}
			rows_affected,err:=result.RowsAffected()
			if err != nil {
				Fatalf("DB error: %v",err)
			}
			if rows_affected > 0 {
				return
			} else {
				Fatalf("DB error: couldn't insert into Open Orders table. Rows affeced = 0")
			}
		// end of case OrderActionCreate
		case OrderActionCancel:
			fmt.Printf("deleting open order %v\n",order_id)
			query = "DELETE FROM oorders WHERE order_id = $1"
			result,err := ss.db.Exec(query,order_id)
			if err != nil {
				Fatalf("DB error: can't delete open order %v : %v",order_id,err)
			}
			rows_affected,err:=result.RowsAffected()
			if err != nil {
				Fatalf("DB error: %v",err)
			}
			if rows_affected > 0 {
				return
			} else {
				Fatalf("DB error: delete of open order %v failed, rows affected = 0",order_id)
			}
		// end of case OrderActionCancel
	} // end of switch order action
}
func (ss *SQLStorage) insert_cancel_0x_order_evt(evt *CancelZeroXOrder) {

/*
	Note:This code is currently disabled because we don't have data feed from
		0x exchange for 'Create' type orders

	universe_id := ss.lookup_universe_id(evt.Universe.String())
	_ = universe_id	// ToDo: add universe_id match condition (for market)
	market_aid := ss.lookup_address(evt.Market.String())
	var order_id = hex.EncodeToString(evt.OrderHash[:])

//	var oaction OrderAction = OrderActionCancel
//	var otype OrderType = OrderType(evt.OrderType)
//	price := evt.Price.Int64()
//	amount := evt.Amount.Int64()
//	outcome := evt.Outcome.Int64()
	var query string
	query = "DELETE FROM oorders WHERE order_id = $1"
	result,err := ss.db.Exec(query,market_aid)
	if err!=nil {
		Fatalf("DB error: couldn't delete open order with order_id = %v",order_id)
	}
	rows_affected,err:=result.RowsAffected()
	if rows_affected == 0  {
		Fatalf("DB error: couldn't delete open order with order_id = %v (not found)",order_id)
	}
*/
}

func (ss *SQLStorage) insert_market_finalized_evt(evt *MktFinalizedEvt) {

	universe_id := ss.lookup_universe_id(evt.Universe.String())
	_ = universe_id	// ToDo: add universe_id match condition (for market)
	market_aid := ss.lookup_address(evt.Market.String())
	fin_timestamp := evt.Timestamp.Int64()
	winning_payouts := bigint_ptr_slice_to_str(&evt.WinningPayoutNumerators,",")

	var query string
	query = "INSERT INTO mkt_fin(market_aid,fin_timestamp,winning_payouts) VALUES($1,$2,$3)"
	_,err := ss.db.Exec(query,market_aid,fin_timestamp,winning_payouts)
	if err != nil {
		Fatalf("DB error: can't update market finalization of market %v : %v",market_aid,err)
	}
}
func (ss *SQLStorage) insert_initial_report_evt(
	block_num BlockNumber,
	tx_id int64,
	signer common.Address,
	evt *InitialReportSubmittedEvt,
) {

	_= ss.lookup_universe_id(evt.Universe.String())
	market_aid := ss.lookup_address(evt.Market.String())
	reporter_aid := ss.lookup_or_create_address(evt.Reporter.String())
	signer_aid := ss.lookup_or_create_address(signer.String())
	ini_reporter_aid := ss.lookup_or_create_address(evt.InitialReporter.String())

	amount_staked := evt.AmountStaked.Int64()
	payout_numerators := bigint_ptr_slice_to_str(&evt.PayoutNumerators,",")
	next_win_start := evt.NextWindowStartTime.Int64()
	next_win_end := evt.NextWindowEndTime.Int64()
	rpt_timestamp := evt.Timestamp.Int64()

	fmt.Printf("insert_initial_report_evt(): market_aid=%v, reporter_id=%v, signer_aid=%v",
					market_aid,reporter_aid,signer_aid)
	var query string
	query = `
		INSERT INTO report (
			block_num,
			tx_id,
			market_aid,
			reporter_aid,
			signer_aid,
			ini_reporter_aid,
			is_initial,
			is_designated,
			amount_staked,
			pnumerators,
			description,
			next_win_start,
			next_win_end,
			rpt_timestamp
		) VALUES ($1,$2,$3,$4,$5,$6,$7,$8,$9,$10,$11,$12,$13,$14)`
	result,err := ss.db.Exec(query,
			block_num,
			tx_id,
			market_aid,
			reporter_aid,
			signer_aid,
			ini_reporter_aid,
			true,
			evt.IsDesignatedReporter,
			amount_staked,
			payout_numerators,
			evt.Description,
			next_win_start,
			next_win_end,
			rpt_timestamp)
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
		Fatalf("DB error: couldn't insert into InitialReport table. Rows affeced = 0")
	}
}
func (ss *SQLStorage) insert_market_volume_changed_evt(
	block_num BlockNumber,
	tx_id int64,
	evt *MktVolumeChangedEvt,
) {

	market_aid := ss.lookup_address(evt.Market.String())

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
		) VALUES ($1,$2,$3,$4,$5,$6)`
	result,err := ss.db.Exec(query,
			block_num,
			tx_id,
			market_aid,
			volume,
			outcome_vols,
			timestamp)
	if err != nil {
		Fatalf("DB error: can't insert into volume table: %v",err)
	}
	rows_affected,err:=result.RowsAffected()
	if err != nil {
		Fatalf("DB error: %v",err)
	}
	if rows_affected > 0 {
		return
	} else {
		Fatalf("DB error: couldn't insert into InitialReport table. Rows affeced = 0")
	}
}
func (ss *SQLStorage) insert_dispute_crowd_contrib(
	block_num BlockNumber,
	tx_id int64,
	signer common.Address,
	evt *DisputeCrowdsourcerContributionEvt,
) {

	_= ss.lookup_universe_id(evt.Universe.String())
	market_aid := ss.lookup_address(evt.Market.String())
	reporter_aid := ss.lookup_or_create_address(evt.Reporter.String())
	signer_aid := ss.lookup_or_create_address(signer.String())
	disputed_aid := ss.lookup_or_create_address(evt.DisputeCrowdsourcer.String())

	amount_staked := evt.AmountStaked.Int64()
	payout_numerators := bigint_ptr_slice_to_str(&evt.PayoutNumerators,",")
	cur_stake := evt.CurrentStake.Int64()
	stake_remaining := evt.StakeRemaining.Int64()
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
			reporter_aid,
			signer_aid,
			disputed_aid,
			dispute_round,
			amount_staked,
			description,
			pnumerators,
			current_stake,
			stake_remaining,
			rpt_timestamp
		) VALUES ($1,$2,$3,$4,$5,$6,$7,$8,$9,$10,$11,$12,$13)`
	result,err := ss.db.Exec(query,
			block_num,
			tx_id,
			market_aid,
			reporter_aid,
			signer_aid,
			disputed_aid,
			dispute_round,
			amount_staked,
			evt.Description,
			payout_numerators,
			cur_stake,
			stake_remaining,
			rpt_timestamp)
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
		Fatalf("DB error: couldn't insert into InitialReport table. Rows affeced = 0")
	}
}
func (ss *SQLStorage) insert_share_balance_changed_evt(
	block_num BlockNumber,
	tx_id int64,
	evt *ShareTokenBalanceChanged,
) {

	_= ss.lookup_universe_id(evt.Universe.String())
	market_aid := ss.lookup_address(evt.Market.String())
	account_aid := ss.lookup_or_create_address(evt.Account.String())

	outcome := evt.Outcome.String()
	balance := evt.Balance.String()

	var query string

	query = "UPDATE sbalances SET balance = $4" +
				"WHERE " +
					"market_aid = $1 AND " +
					"account_aid = $2 AND " +
					"outcome = $3"
	result,err := ss.db.Exec(query,	market_aid,account_aid,outcome,balance)
	if err != nil {
		Fatalf("DB error: can't update 'sbalances' for account %v, market %v : %v",
					evt.Account.String(),evt.Market.String(),err)
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
		query = "INSERT INTO sbalances (block_num,tx_id,account_aid,market_aid,outcome,balance)" +
				"VALUES($1,$2,$3,$4,$5,$6)"
		result,err := ss.db.Exec(query,block_num,tx_id,account_aid,market_aid,outcome,balance)
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
		Fatalf("DB error: can't insert into block  table: %v",err)
	}
	rows_affected,err:=result.RowsAffected()
	if err != nil {
		Fatalf("DB error: %v",err)
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
		Fatalf("DB error: can't insert into transactions table: %v",err)
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
			Fatalf("DB error: %v",err)
		}
	}
	cur_block_num := int64(block.Number.Uint64())
	if cur_block_num > (my_block_num + MAX_BLOCKS_CHAIN_SPLIT) {
		Fatalf("Chainsplit detected, and it is more than %v blocks, aborting.",MAX_BLOCKS_CHAIN_SPLIT)
	}
	query = "DELETE FROM block WHERE block_num > $1 CASCADE"
	_,err = ss.db.Exec(query,my_block_num)
	if (err!=nil) {
		Fatalf("DB error: %v",err);
	}
	return BlockNumber(my_block_num + 1)	// parent + 1 = current
}
func (ss *SQLStorage) block_delete_with_everything(block_num BlockNumber) {

	// deletes block table and all the other tables receieve cascaded DELETEs also
	var query string
	query = "DELETE FROM block WHERE block_num = $1"
	_,err := ss.db.Exec(query,block_num)
	if (err!=nil) {
		Fatalf("DB error: %v (block_num=%v)",err,block_num);
	}
}
