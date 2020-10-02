package dbs

import (
	"fmt"
	"os"
	"math/big"
	"strings"
	"strconv"
	"errors"
	"database/sql"
	"encoding/hex"
	_  "github.com/lib/pq"

	ztypes "github.com/0xProject/0x-mesh/common/types"
	"github.com/ethereum/go-ethereum/common"
	"github.com/0xProject/0x-mesh/zeroex"

	p "github.com/PredictionExplorer/augur-explorer/primitives"
)
func (ss *SQLStorage) Insert_market_order_evt(agtx *p.AugurTx,timestamp int64,p_eoa_aid int64,p_eoa_fill_aid int64,	evt *p.EOrderEvent,submitted_orders map[string]*ztypes.OrderInfo,order_specs map[string]*p.ZxMeshOrderSpec) {

	// depending on the order action (Create/Cancel/Fill) different table is used for storage
	//		Create/Cancel order actions go to 'oorders' (Open Orders) table because these orders
	//		do not alter market's open interest.
	//		Fill order goes to 'mktord' table because the share has been created and now
	//		open interest increased

	var order_hash_obj = common.BytesToHash(evt.OrderId[:])
	var order_hash = order_hash_obj.String()

	zorder := submitted_orders[order_hash]
	initial_amount := zorder.SignedOrder.MakerAssetAmount
	if initial_amount == nil {
		ss.Log_msg(
			fmt.Sprintf(
				"Initial amount for order %v (block %v) not found",
				order_hash,agtx.BlockNum,
			),
		)
		os.Exit(1)
	}

	var wallet_aid int64;
	wallet_aid = ss.Lookup_or_create_address(evt.AddressData[0].String(),agtx.BlockNum,agtx.TxId)
	var wallet_fill_aid int64 = 0;
	if len(evt.AddressData) > 1 {
		wallet_fill_aid = ss.Lookup_or_create_address(evt.AddressData[1].String(),agtx.BlockNum,agtx.TxId)
	}
	universe_id,err := ss.lookup_universe_id(evt.Universe.String())
	if err!=nil {
		ss.Log_msg(
			fmt.Sprintf(
				"Universe %v wasn't found when trying toinsert MarketOrder event at block %v: %v",
				evt.Universe.String(),agtx.BlockNum,err,
			),
		)
		os.Exit(1)
	}
	_ = universe_id	// ToDo: add universe_id match condition (for market)
	market_aid,_:= ss.lookup_market_id(evt.Market.String())
	eoa_aid,err := ss.Lookup_eoa_aid(wallet_aid);
	if err!=nil {
		// sometimes creator can be an EOA, so we set eoa_aid to wallet_aid
		eoa_aid = wallet_aid
	}

	eoa_fill_aid,err := ss.Lookup_eoa_aid(wallet_fill_aid)
	if err != nil {
		// sometimes creator can be an EOA, so we set eoa_aid to wallet_aid
		eoa_fill_aid = wallet_fill_aid
	}

	var oaction p.OrderAction = p.OrderAction(evt.EventType)
	var otype p.OrderType = p.OrderType(evt.OrderType)
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
	amount := initial_amount
	outcome_idx := evt.Uint256Data[2].Int64()
	token_refund := evt.Uint256Data[3].String()
	shares_refund := evt.Uint256Data[4].String()
	fees := evt.Uint256Data[5].String()
	amount_filled := evt.Uint256Data[6]
	shares_escrowed := evt.Uint256Data[8].String()
	tokens_escrowed := evt.Uint256Data[9].String()

	mesh_evt_code := p.MeshEvtFullyFilled
	if 0 != initial_amount.Cmp(amount_filled) {
		mesh_evt_code = p.MeshEvtFilled
	}
	ss.Insert_0x_mesh_order_event(timestamp,zorder,order_specs[order_hash],amount_filled,mesh_evt_code)

	var query string
	var opcode int = p.OOOpCodeFill
	ss.Info.Printf("amount = %v, amount_filled = %v, opcode=%v\n",amount,amount_filled,opcode)
/* Discontinued
	query = "DELETE FROM oorders WHERE order_hash = $1"
	_,err = ss.db.Exec(query,market_aid)
	if err!=nil {
		msg:=fmt.Sprintf("DB error: couldn't delete open order with hash = %v\n",order_hash)
		ss.Info.Printf(msg)
		ss.Log_msg(msg)
		os.Exit(1)
	}
*/
	ss.Info.Printf("OrderAction = %v, otype=%v, order_hash=%v\n",oaction,otype,order_hash)
	ss.Info.Printf("Filling existing order %v\n",order_hash)
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
			outcome_idx,
			token_refund,
			shares_refund,
			fees,
			amount_filled,
			time_stamp,
			shares_escrowed,
			tokens_escrowed,
			trade_group,
			order_hash
		) VALUES (
				$1,$2,$3,$4,$5,$6,$7,$8,$9,
				` + price + "," +
				"(" + amount.String() + "/1e+18)," +
				"$10," +
				"(" + token_refund + "/1e+18)," +
				"(" + shares_refund + "/1e+18)," +
				"(" + fees + "/1e+18)," +
				"(" + amount_filled.String() + "/1e+18)," +
				"TO_TIMESTAMP($11)," +
				"$12,$13,$14,$15) RETURNING id"

	var null_id sql.NullInt64
	err=ss.db.QueryRow(query,
			agtx.TxId,
			market_aid,
			eoa_aid,
			wallet_aid,
			eoa_fill_aid,
			wallet_fill_aid,
			agtx.BlockNum,
			oaction,
			otype,
			outcome_idx,
			timestamp,
			shares_escrowed,
			tokens_escrowed,
			hex.EncodeToString(evt.TradeGroupId[:]),
			order_hash,
	).Scan(&null_id);
	if (err!=nil) {
		ss.Log_msg(
			fmt.Sprintf(
				"DB error: can't insert into mktord table at block: %v : %v, q=%v",
				agtx.BlockNum,err,query,
			),
		)
		os.Exit(1)
	}
	if null_id.Valid {
		*(ss.mkt_order_id_ptr) = null_id.Int64
	} else {
		*(ss.mkt_order_id_ptr) = 0
	}
	query = "UPDATE outcome_vol " +
			"SET " +
				"last_price = "+price+ " " +
			"WHERE " +
				"market_aid = $1 AND outcome_idx = $2"
	_,err = ss.db.Exec(query,market_aid,outcome_idx)
	if (err!=nil) {
		ss.Log_msg(fmt.Sprintf("DB error at block %v : %v ; q=%v",agtx.BlockNum,err,query))
		os.Exit(1)
	}
	ss.Update_open_order_history(*ss.mkt_order_id_ptr,order_hash,timestamp,amount_filled.String(),opcode)
}
func (ss *SQLStorage) Update_open_order_history(mktord_id int64,order_hash string,timestamp int64,amount_filled string,opcode int) {
/* Discontinued
	// Note: the following function may have null effect if the corresponding record already exists
	//			the unique-key for the record is orderhash+opcode, so multiple calls can be made
	//			to this function without any problem. In fact multiple calls will come from the
	//			order insertion process, and from 0x Mesh daemon at the same time. We can't say which is
	//			going to be the first to insert the record, but the second call will result in
	//			no effect at all
	var err error
	var query string
	var d_query string
	if mktord_id == 0 {
		query = "SELECT update_oo_hist(NULL,$1,$2,$3,$4)"
		d_query = fmt.Sprintf("SELECT update_oo_hist(NULL,'%v',%v,'%v',%v)",order_hash,timestamp,amount_filled,opcode)
		_,err = ss.db.Exec(query,order_hash,timestamp,amount_filled,opcode)
	} else {
		query = "SELECT update_oo_hist($1,$2,$3,$4,$5)"
		d_query = fmt.Sprintf("SELECT update_oo_hist(%v,'%v',%v,'%v',%v)",mktord_id,order_hash,timestamp,amount_filled,opcode)
		_,err = ss.db.Exec(query,mktord_id,order_hash,timestamp,amount_filled,opcode)
	}
	ss.Info.Printf("Update_open_order_history query: %v\n",d_query)
	if err!=nil {
		msg:=fmt.Sprintf("DB error: couldn't update history of order with hash = %v mktord=%v: %v\n",order_hash,mktord_id,err)
		ss.Info.Printf(msg)
		ss.Log_msg(msg)
		os.Exit(1)
	}
	*/
}
func (ss *SQLStorage) Insert_open_order(ohash *string,order *zeroex.SignedOrder,fillable_amount *big.Int,eoa_addr string,ospec *p.ZxMeshOrderSpec,opcode int,evt_timestamp int64) error {
	// Insert an open order, this order needs to be Filled by another market participant
	// It also can be canceled by its creator (with another transaction)
	var err error
	order_hash := ohash
	expiration := order.ExpirationTimeSeconds.Int64()
	// note: we don't have block number/tx hash for activity from 0x Mesh, so we insert with 0s
	ss.Info.Printf(
		"Open Order: Market %v, Price %v, Otcome %v\n",
		ospec.Market.String(),ospec.Price.String(),ospec.Outcome,
	)
	initial_amount := order.MakerAssetAmount.String()
	// Note: the MakerAddress can be either EOA of the User or Wallet contract of the User
	//			we need to figure out which one we have been given
	var wallet_aid int64 = 0
	var eoa_aid int64 = 0
	zero_addr := common.BigToAddress(zero)
	if zero_addr.String()==eoa_addr {	// MakerAddress must be an EOA address
		eoa_aid,err = ss.Nonfatal_lookup_address_id(order.MakerAddress.String())
		if err != nil {
			// Note: we can't INSERT an address from here because we need Transaction Hash (for reference)
			ss.Info.Printf(
				"MakerAddress %v is unregistered in the DB as EOA ",
				order.MakerAddress.String(),
			)
			return errors.New(
				fmt.Sprintf("MakerAddress %v is unregistered in the DB as EOA.",order.MakerAddress.String()),
			)
		}
		// Now we need to validate this this address is indeed an EOA
		wallet_aid,err = ss.Lookup_wallet_aid(eoa_aid)
		if err != nil {
			ss.Info.Printf(
				"MakerAddress %v doesn't have an associated Wallet contract",
				order.MakerAddress.String(),
			)
			return errors.New("EOA address provided from Mesh listener is zero address and un-registered.")
		}
	} else { // MakerAddress has an EOA address, means MakerAddress is a Wallet contract
		wallet_aid,err = ss.Nonfatal_lookup_address_id(order.MakerAddress.String())
		if err != nil {
			// Maker is not Wallet Contract, then it must be EOA
			ss.Info.Printf(
				"MakerAddress %v is unregistered in the DB as wallet : %v",
				order.MakerAddress.String(),err,
			)
			return errors.New(
				fmt.Sprintf("MakerAddress %v is unregistered in the DB.",order.MakerAddress.String()),
			)
		}
		eoa_aid,err = ss.Lookup_eoa_aid(wallet_aid)
		if err != nil {
			ss.Info.Printf(
				"MakerAddress %v is a Wallet contract that doesn't have an associated EOA address: %v",
				order.MakerAddress.String,err,
			)
			return errors.New(
				fmt.Sprintf("MakerAddress %v doesn't have associated EOA",order.MakerAddress.String()),
			)
		}
	}
	ss.Link_eoa_and_wallet_contract(eoa_aid,wallet_aid) // enforce EOA-Wallet link (though it may exist)

	ss.Info.Printf(
		"creating open order made by %v : market=%v, price=%v, Outcome=%v, Type=%v\n",
		eoa_addr,ospec.Market.String(),ospec.Price.String(),ospec.Outcome,ospec.Type,
	)
	market_aid,err := ss.Nonfatal_lookup_address_id(ospec.Market.String())
	if err != nil {
		ss.Info.Printf(
			"Cant find market %v, probably 0x Mesh is ahead of the main chain",
			ospec.Market.String(),
		)
		return errors.New("Market not yet registered")
	}
	price := float64(ospec.Price.Int64())
	otype := ospec.Type	// Bid/Ask
	amount := fillable_amount.String()

	var query string
	query = "INSERT INTO oostats(market_aid,eoa_aid,outcome_idx) VALUES($1,$2,$3)"
	_,err = ss.db.Exec(query,market_aid,eoa_aid,ospec.Outcome)
	if err != nil {
		if !strings.Contains(err.Error(),"duplicate key value") {
			ss.Log_msg(fmt.Sprintf("DB error: %v, q=%v\n",err,query))
		}
	}
	query = "INSERT INTO oorders(" +
				"market_aid,otype,wallet_aid,eoa_aid,price,initial_amount,amount,outcome_idx,opcode," +
				"evt_timestamp,srv_timestamp,expiration,order_hash" +
			") VALUES("+
				"$1,$2,$3,$4,$5,"+
				initial_amount+"/1e+18," + amount+"/1e+18,"+
				"$6,$7," +
				"TO_TIMESTAMP($8),NOW(),TO_TIMESTAMP($9),$10"+
			") " +
			"ON CONFLICT DO NOTHING"
	d_query := fmt.Sprintf("INSERT INTO oorders(" +
				"market_aid,otype,wallet_aid,eoa_aid,price,initial_amount,amount,outcome_idx,opcode," +
				"evt_timestamp,srv_timestamp,expiration,order_hash" +
			") VALUES("+
				"%v,%v,%v,%v,%v,"+
				initial_amount+"/1e+18," + amount+"/1e+18,"+
				"%v,%v," +
				"TO_TIMESTAMP(%v),NOW(),TO_TIMESTAMP(%v),'%v'"+
			") " +
			"ON CONFLICT DO NOTHING",
			market_aid,
			otype,
			wallet_aid,
			eoa_aid,
			price,
			ospec.Outcome,
			opcode,
			evt_timestamp,
			expiration,
			*order_hash,
	)
	ss.Info.Printf("query = %v\n",d_query)
	result,err := ss.db.Exec(query,
			market_aid,
			otype,
			wallet_aid,
			eoa_aid,
			price,
			ospec.Outcome,
			opcode,
			evt_timestamp,
			expiration,
			*order_hash)
	if err != nil {
		ss.Log_msg(fmt.Sprintf("DB error: can't insert into open orders table: %v, q=%v",err,query))
		return err
	}
	rows_affected,err:=result.RowsAffected()
	if err != nil {
		ss.Log_msg(fmt.Sprintf("DB error: %v, q=%v",err,query))
	}
	if rows_affected > 0 {
		return nil
	} else {
		ss.Log_msg(fmt.Sprintf("DB error: couldn't insert into Open Orders table. Rows affeced = 0"))
	}
	return errors.New("Affected rows=0")
}
func (ss *SQLStorage) Cancel_open_order(orders map[string]*ztypes.OrderInfo,order_specs map[string]*p.ZxMeshOrderSpec,order_hash string,timestamp int64) {

	oinfo := orders[order_hash]
	if oinfo == nil {
		ss.Log_msg(
			fmt.Sprintf(
				"DB error: can't find order on CancelZOrder event, hash=%v\n",order_hash,
			),
		)
		os.Exit(1)
	}
	ospec := order_specs[order_hash]
	if ospec == nil {
		ss.Log_msg(
			fmt.Sprintf(
				"No Augur Trading data for order hash=%v in CancelOrder event\n",order_hash,
			),
		)
		os.Exit(1)
	}

	ss.Insert_0x_mesh_order_event(timestamp,oinfo,ospec,nil,p.MeshEvtCancelled)
	ss.Update_open_order_history(0,order_hash,timestamp,"0",p.OOOpCodeCancelledByUser)

	var query string
	query = "DELETE FROM oorders WHERE order_hash = $1"
	result,err := ss.db.Exec(query,order_hash)
	if err!=nil {
		ss.Info.Printf(fmt.Sprintf("DB error: couldn't delete open order with order_hash = %v, q=%v\n",order_hash,query))
	}
	rows_affected,err:=result.RowsAffected()
	if rows_affected == 0  {
		ss.Info.Printf(fmt.Sprintf("DB error: couldn't delete open order with order_hash = %v (not found)\n",order_hash))
	}
}
func (ss *SQLStorage) Delete_open_0x_order(order_hash string,timestamp int64,opcode int) {

	if opcode != p.OOOpCodeNone {
		ss.Update_open_order_history(0,order_hash,timestamp,"0",opcode)
	}

	var query string
	query = "DELETE FROM oorders WHERE order_hash = $1"
	result,err := ss.db.Exec(query,order_hash)
	if err!=nil {
		ss.Info.Printf(fmt.Sprintf("DB error: couldn't delete open order with order_hash = %v, q=%v\n",order_hash,query))
	}
	rows_affected,err:=result.RowsAffected()
	if rows_affected == 0  {
		ss.Info.Printf(fmt.Sprintf("DB error: couldn't delete open order with order_hash = %v (not found)\n",order_hash))
	}
}
func (ss *SQLStorage) Update_0x_order_on_partial_fill(oinfo *ztypes.OrderInfo) {

	order_hash := oinfo.OrderHash.String()
	amount := oinfo.FillableTakerAssetAmount.String()

	var query string
	query = "UPDATE oorders "+
			"SET id=DEFAULT,amount=("+amount+"/1e+18) "+
			"WHERE order_hash = $1"
			// Note: we update the 'id' field to the next sequential number too because
			//		the UI needs to receive a notification to refresh the Market Depth chart
	result,err := ss.db.Exec(query,order_hash)
	if err!=nil {
		ss.Info.Printf(fmt.Sprintf("DB error: couldn't delete open order with order_hash = %v, q=%v\n",order_hash,query))
	}
	rows_affected,err:=result.RowsAffected()
	if rows_affected == 0  {
		ss.Info.Printf(fmt.Sprintf("DB error: couldn't delete open order with order_hash = %v (not found)\n",order_hash))
	}
}
func (ss *SQLStorage) close_all_open_positions_for_market(market_aid int64) {

	var query string
	// close existing positions
	query = "UPDATE profit_loss " +
					"SET closed_position = 1 WHERE market_aid = $1 and closed_position = 0"
	_,err:=ss.db.Exec(query,market_aid)
	if (err!=nil) {
		ss.Log_msg(fmt.Sprintf("DB error for market_aid=%v: %v ; q=%v",market_aid,err,query))
		os.Exit(1)
	}
}
func (ss *SQLStorage) Get_mkt_trades(mkt_addr string,limit int) []p.MarketTrade {
	// get market trades with mixed outcomes
	var where string = ""
	var market_aid int64 = 0;
	if len(mkt_addr) > 0 {
		market_aid = ss.Lookup_address_id(mkt_addr)
		where = " WHERE o.market_aid = $1 "
	}
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
				"LEFT JOIN address AS a ON o.market_aid=a.address_id " +
				"LEFT JOIN address AS fa ON o.eoa_fill_aid=fa.address_id " +
				"LEFT JOIN address AS ca ON o.eoa_aid=ca.address_id " +
				"LEFT JOIN market AS m ON o.market_aid = m.market_aid " +
			where +
			"ORDER BY o.block_num DESC,o.time_stamp DESC"
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
			&outcomes,
		)
		if err!=nil {
			ss.Log_msg(fmt.Sprintf("DB error: %v, q=%v",err,query))
			os.Exit(1)
		}
		p.Augur_UI_price_adjustments(&rec.Price,&rec.Amount,mkt_type)
		rec.OrderHashSh=p.Short_hash(rec.OrderHash)
		rec.MktAddrSh=p.Short_address(rec.MktAddr)
		rec.CreatorAddrSh=p.Short_address(rec.CreatorAddr)
		rec.FillerAddrSh=p.Short_address(rec.FillerAddr)
		rec.OutcomeStr = get_outcome_str(uint8(mkt_type),rec.Outcome,&outcomes)
		records = append(records,rec)
	}
	return records
}
func (ss *SQLStorage) build_depth_by_otype(market_aid int64,outc int,otype p.OrderType) ([]p.DepthEntry,int64) {

	var query string
	query = "SELECT " +
				"o.id," +
				"o.market_aid," +
				"o.outcome_idx," +
				"m.market_type," +
				"wa.addr AS wallet_addr," +
				"ua.addr AS user_addr," +
				"o.srv_timestamp::date AS date_created," +
				"o.expiration::date AS expires," +
				"FLOOR(EXTRACT(EPOCH FROM o.expiration))::BIGINT as expires_ts," +
				"o.price AS price, " +
				"o.amount AS amount," +
				"s.num_bids," +
				"s.num_asks," +
				"s.num_cancel " +
			"FROM oorders AS o " +
				"JOIN market AS m ON m.market_aid=o.market_aid " +
				"LEFT JOIN oostats AS s ON (" +
						"o.market_aid=s.market_aid AND " +
						"o.eoa_aid=s.eoa_aid AND " +
						"o.outcome_idx=s.outcome_idx" +
				") " +
				"LEFT JOIN address AS a ON o.market_aid=a.address_id " +
				"LEFT JOIN address AS wa ON o.wallet_aid=wa.address_id " +
				"LEFT JOIN address AS ua ON o.eoa_aid=ua.address_id " +
			"WHERE o.market_aid = $1 AND o.outcome_idx=$2 AND o.otype = $3 " +
			"ORDER BY "
	if otype == p.OrderTypeBid {
				query = query + "o.price DESC,o.evt_timestamp DESC";
	} else {
				query = query + "o.price ASC,o.evt_timestamp DESC";
	}
	ss.Info.Printf("q=%v\n",query)
	var accumulated_volume = 0.0
	rows,err := ss.db.Query(query,market_aid,outc,otype)
	if (err!=nil) {
		ss.Log_msg(fmt.Sprintf("DB error: %v (query=%v)",err,query))
		os.Exit(1)
	}
	records := make([]p.DepthEntry,0,8)
	var max_id int64 = 0
	var oo_id int64 = 0
	defer rows.Close()
	for rows.Next() {
		var mkt_type int
		var rec p.DepthEntry
		var num_bids sql.NullInt64
		var num_asks sql.NullInt64
		var num_cancels sql.NullInt64
		err=rows.Scan(
			&oo_id,
			&rec.MktAid,
			&rec.OutcomeIdx,
			&mkt_type,
			&rec.WalletAddr,
			&rec.EOAAddr,
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
			ss.Log_msg(fmt.Sprintf("DB error: %v, q=%v",err,query))
			os.Exit(1)
		}
		p.Augur_UI_price_adjustments(&rec.Price,&rec.Volume,mkt_type)
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
		rec.WalletAddrSh=p.Short_address(rec.WalletAddr)
		rec.EOAAddrSh=p.Short_address(rec.EOAAddr)
		records = append(records,rec)
		if max_id < oo_id {
			max_id = oo_id
		}
	}
	return records,max_id
}
func (ss *SQLStorage) Get_price_history_for_outcome(market_aid int64,outc int) []p.OrderInfo {

	var query string
	query = "SELECT " +
				"o.id,"+
				"o.order_hash," +
				"o.market_aid," +
				"m.market_type," +
				"c_w_a.addr AS c_w_a_addr," +
				"c_e_a.addr AS filler_eoa_addr," +
				"f_w_a.addr AS f_w_a_addr," +
				"f_e_a.addr AS filler_eoa_addr," +
				"o.otype, " +
				"CASE o.otype " +
					"WHEN 0 THEN 'BID' " +
					"ELSE 'ASK' " +
				"END AS dir, " +
				"o.time_stamp::date AS date," +
				"FLOOR(EXTRACT(EPOCH FROM o.time_stamp))::BIGINT as created_ts," +
				"o.outcome_idx," +
				"o.price AS price, " +
				"o.amount_filled AS volume " +
			"FROM mktord AS o " +
				"JOIN market AS m on o.market_aid=m.market_aid " +
				"LEFT JOIN address AS a ON o.market_aid=a.address_id " +
				"LEFT JOIN address AS c_w_a ON o.wallet_aid=c_w_a.address_id " +
				"LEFT JOIN address AS c_e_a ON o.eoa_aid=c_e_a.address_id " +
				"LEFT JOIN address AS f_w_a ON o.wallet_fill_aid=f_w_a.address_id " +
				"LEFT JOIN address AS f_e_a ON o.eoa_fill_aid=f_e_a.address_id " +
			"WHERE o.market_aid = $1 AND o.outcome_idx=$2 " +
			"ORDER BY o.time_stamp"
	var accumulated_volume = 0.0
	rows,err := ss.db.Query(query,market_aid,outc)
	if (err!=nil) {
		ss.Log_msg(fmt.Sprintf("DB error: %v (query=%v)",err,query))
		os.Exit(1)
	}
	records := make([]p.OrderInfo,0,8)

	defer rows.Close()
	for rows.Next() {
		var rec p.OrderInfo
		var mkt_type int
		err=rows.Scan(
			&rec.OrderId,
			&rec.OrderHash,
			&rec.MktAid,
			&mkt_type,
			&rec.CreatorWalletAddr,
			&rec.CreatorEOAAddr,
			&rec.FillerWalletAddr,
			&rec.FillerEOAAddr,
			&rec.OType,
			&rec.Direction,
			&rec.Date,
			&rec.CreatedTs,
			&rec.OutcomeIdx,
			&rec.Price,
			&rec.Amount,
		)
		if err!=nil {
			ss.Log_msg(fmt.Sprintf("DB error: %v",err))
			os.Exit(1)
		}
		p.Augur_UI_price_adjustments(&rec.Price,&rec.Amount,mkt_type)
		rec.CreatorBuyer = true
		rec.FillerBuyer = false
		if rec.OType == 1 {
			rec.CreatorBuyer = false
			rec.FillerBuyer = true
		}
		rec.CreatorWalletAddrSh=p.Short_address(rec.CreatorWalletAddr)
		rec.CreatorEOAAddrSh=p.Short_address(rec.CreatorEOAAddr)
		rec.FillerWalletAddrSh=p.Short_address(rec.FillerWalletAddr)
		rec.FillerEOAAddrSh=p.Short_address(rec.FillerEOAAddr)
		accumulated_volume = accumulated_volume + rec.Amount
		rec.AccumVol = accumulated_volume
		records = append(records,rec)
	}
	return records
}
func (ss *SQLStorage) Get_full_price_history(mkt_addr string,market_aid int64) p.FullPriceHistory {

	var output p.FullPriceHistory
	outcomes,_ := ss.Get_outcome_volumes(mkt_addr,market_aid,0);
	for _,outc := range outcomes {
		var ph p.PriceHistory
		ph.OutcomeIdx = outc.Outcome
		ph.OutcomeStr = outc.OutcomeStr
		ph.Trades = ss.Get_price_history_for_outcome(market_aid,outc.Outcome)
		output.Outcomes = append(output.Outcomes,ph)
	}
	return output
}
func (ss *SQLStorage) Get_zoomed_t1_price_history_for_outcome(market_aid int64,outc int,init_ts int,fin_ts int) []p.ZHistT1Entry {

	var query string
	/* DISCONTINUED
	query = "SELECT " +
				"o.id,"+
				"o.order_hash," +
				"o.market_aid," +
				"c_e_a.addr AS creator_addr," +
				"o.otype, " +
				"CASE o.otype " +
					"WHEN 0 THEN 'BID' " +
					"ELSE 'ASK' " +
				"END AS dir, " +
				"o.evt_timestamp::date AS date," +
				"FLOOR(EXTRACT(EPOCH FROM o.evt_timestamp))::BIGINT as created_ts," +
				"FLOOR(EXTRACT(EPOCH FROM o.expiration))::BIGINT as expiration_ts," +
				"o.opcode," +
				"o.outcome_idx," +
				"o.price AS price, " +
				"o.price_estimate, " +
				"o.initial_amount, " +
				"o.amount " +
			"FROM oohist AS o " +
				"LEFT JOIN " +
					"address AS a ON o.market_aid=a.address_id " +
				"LEFT JOIN address AS c_e_a ON o.eoa_aid=c_e_a.address_id " +
			"WHERE " +
				"o.market_aid = $1 AND " +
				"o.outcome_idx=$2 AND " +
				"o.evt_timestamp >= TO_TIMESTAMP($3) AND "+
				"o.evt_timestamp < TO_TIMESTAMP($4) " +
			"ORDER BY o.evt_timestamp"
	*/
	query = "SELECT " +
				"p.id,"+
				"e.order_hash," +
				"p.market_aid," +
				"'PENDING' AS creator_addr," +
				"e.otype, " +
				"CASE e.otype " +
					"WHEN 0 THEN 'BID' " +
					"ELSE 'ASK' " +
				"END AS dir, " +
				"p.time_stamp::date AS date," +
				"FLOOR(EXTRACT(EPOCH FROM p.time_stamp))::BIGINT as time_stamp," +
				"FLOOR(EXTRACT(EPOCH FROM e.expiration_time))::BIGINT as expiration_ts," +
				"e.evt_code," +
				"p.outcome_idx," +
				"e.price AS price, " +
				"p.price_est, " +
				"p.wprice_est, " +
				"e.fillable_amount, " +
				"e.amount_fill, " +
				"p.max_bid," +
				"p.min_ask," +
				"p.wmax_bid," +
				"p.wmin_ask," +
				"p.spread " +
			"FROM price_estimate AS p " +
				"JOIN mesh_evt AS e ON p.meshevt_id=e.id " +
				"LEFT JOIN " +
					"address AS a ON p.market_aid=a.address_id " +
//				"LEFT JOIN address AS c_e_a ON o.eoa_aid=c_e_a.address_id " +
			"WHERE " +
				"p.market_aid = $1 AND " +
				"p.outcome_idx=$2 AND " +
				"p.time_stamp >= TO_TIMESTAMP($3) AND "+
				"p.time_stamp < TO_TIMESTAMP($4) " +
			"ORDER BY p.time_stamp"
	ss.Info.Printf(
		"market_aid=%v, outcome=%v, init_ts=%v, fin_ts=%v, query=%v\n",
		market_aid,outc,init_ts,fin_ts,query,
	)
	rows,err := ss.db.Query(query,market_aid,outc,init_ts,fin_ts)
	if (err!=nil) {
		ss.Log_msg(fmt.Sprintf("DB error: %v (query=%v)",err,query))
		os.Exit(1)
	}
	records := make([]p.ZHistT1Entry,0,8)

	defer rows.Close()
	for rows.Next() {
		var rec p.ZHistT1Entry
		err=rows.Scan(
			&rec.Id,
			&rec.OrderHash,
			&rec.MktAid,
			&rec.CreatorAddr,
			&rec.OrderType,
			&rec.Direction,
			&rec.OrderDate,
			&rec.Timestamp,
			&rec.OrderExpirationTs,
			&rec.EvtCode,
			&rec.OutcomeIdx,
			&rec.Price,
			&rec.PriceEstimate,
			&rec.WeightedPriceEst,
			&rec.FillableAmount,
			&rec.Amount,
			&rec.MaxBid,
			&rec.MinAsk,
			&rec.WMaxBid,
			&rec.WMinAsk,
			&rec.Spread,
		)
		if err!=nil {
			ss.Log_msg(fmt.Sprintf("DB error: %v",err))
			os.Exit(1)
		}
		//rec.CreatorAddrSh=p.Short_address(rec.CreatorAddr)
		records = append(records,rec)
	}
	return records
}
func (ss *SQLStorage) Get_zoomed_t2_price_history_for_outcome(market_aid int64,outc int,init_ts int,fin_ts int,interval int) []p.ZHistT2Entry {

	var query string
	/*
	query = "SELECT " +
				"ROUND(FLOOR(EXTRACT(EPOCH FROM o.evt_timestamp))/$5)::BIGINT*$5::BIGINT AS start_ts,"+
				"AVG(price_estimate) AS avg_price_estimate " +
			"FROM oohist AS o " +
			"WHERE " +
				"o.market_aid = $1  AND " +
				"o.outcome_idx = $2 AND " +
				"o.evt_timestamp >= TO_TIMESTAMP($3) AND "+
				"o.evt_timestamp < TO_TIMESTAMP($4) " +
			"GROUP BY start_ts " +
			"ORDER BY start_ts"
	*/
	query = "SELECT " +
				"ROUND(FLOOR(EXTRACT(EPOCH FROM p.time_stamp))/$5)::BIGINT*$5::BIGINT AS start_ts,"+
				"AVG(price_est) AS avg_price_estimate, " +
				"AVG(wprice_est) AS avg_weighted_price_est " +
			"FROM price_estimate as p " +
			"WHERE " +
				"p.market_aid = $1  AND " +
				"p.outcome_idx = $2 AND " +
				"p.time_stamp >= TO_TIMESTAMP($3) AND "+
				"p.time_stamp < TO_TIMESTAMP($4) " +
			"GROUP BY start_ts " +
			"ORDER BY start_ts"
	rows,err := ss.db.Query(query,market_aid,outc,init_ts,fin_ts,interval)
	if (err!=nil) {
		ss.Log_msg(fmt.Sprintf("DB error: %v (query=%v)",err,query))
		os.Exit(1)
	}
	records := make([]p.ZHistT2Entry,0,8)

	defer rows.Close()
	for rows.Next() {
		var rec p.ZHistT2Entry
		err=rows.Scan(
			&rec.Timestamp,
			&rec.PriceEstimate,
			&rec.WeightedPriceEstimate,
		)
		if err!=nil {
			ss.Log_msg(fmt.Sprintf("DB error: %v",err))
			os.Exit(1)
		}
		records = append(records,rec)
	}
	return records
}
func (ss *SQLStorage) Get_zoomed_price_history(mkt_addr string,market_aid int64,zoom int,init_ts int,fin_ts int,interval int) p.FullZoomedPriceHist {

	var output p.FullZoomedPriceHist
	outcomes,_ := ss.Get_outcome_volumes(mkt_addr,market_aid,0);
	for _,outc := range outcomes {
		var ph p.ZoomedPriceHist
		ph.OutcomeIdx = outc.Outcome
		ph.OutcomeStr = outc.OutcomeStr
		ph.Zoom = zoom
		ph.InitTs = init_ts
		ph.FinTs = fin_ts
		ph.IntervalSecs = interval
		if zoom == 0 {
			ph.Type1Entries= ss.Get_zoomed_t1_price_history_for_outcome(market_aid,outc.Outcome,init_ts,fin_ts)
		}
		if zoom == 1 {
			ph.Type2Entries = ss.Get_zoomed_t2_price_history_for_outcome(market_aid,outc.Outcome,init_ts,fin_ts,interval)
		}
		output.Outcomes = append(output.Outcomes,ph)
	}
	return output
}
func (ss *SQLStorage) Get_last_open_order_id() int64 {

	var query string
	query = "SELECT id FROM oorders ORDER BY id DESC LIMIT 1"

	var null_id sql.NullInt64
	var err error
	row := ss.db.QueryRow(query)
	err=row.Scan(&null_id);
	if (err!=nil) {
		if err == sql.ErrNoRows {
			return 0
		} else {
			ss.Log_msg(fmt.Sprintf("Error in Get_last_open_order_id(): %v, q=%v",err,query))
			os.Exit(1)
		}
	}
	return null_id.Int64
}
func (ss *SQLStorage) Get_mkt_depth(market_aid int64,outcome_idx int) (*p.MarketDepth,int64) {

	market_depth := new(p.MarketDepth)
	var max_buys,max_sells int64	// max_id is required for polling new open orders on the Client side
	market_depth.Bids,max_buys = ss.build_depth_by_otype(market_aid,outcome_idx,p.OrderTypeBid)
	market_depth.Asks,max_sells = ss.build_depth_by_otype(market_aid,outcome_idx,p.OrderTypeAsk)
	var max_id int64
	if max_buys > max_sells {
		max_id = max_buys
	} else {
		max_id = max_sells
	}
	return market_depth,max_id
}
func (ss *SQLStorage) Get_open_positions(eoa_aid int64) []p.PLEntry {
	return ss.Get_trade_data(eoa_aid,true)
}
func (ss *SQLStorage) close_previous_positions(market_aid int64,eoa_aid int64,outcome_idx int,profit_loss string) string {

	var err error
	var query string

	var pl_update string
	if len(profit_loss) > 0  {
		pl_update = ",immediate_profit=(" + profit_loss + "/1e+36) "
	}
	query = "UPDATE profit_loss " +
				"SET closed_position = 1 " +
				pl_update +
				"WHERE " +
						"(market_aid = $1) AND " +
						"(eoa_aid = $2) AND " +
						"(outcome_idx = $3) AND " +
						"(closed_position = 0) " +
				"RETURNING round(realized_profit*1e+36)::text"
	var previous_profit string
	row:=ss.db.QueryRow(query,market_aid,eoa_aid,outcome_idx)
	err=row.Scan(&previous_profit);
	if err != nil {
		if err == sql.ErrNoRows {
		} else {
			ss.Log_msg(fmt.Sprintf("DB error: %v, (on Scan of previous profit) q=%v",err,query))
			os.Exit(1)
		}
	}
	ss.Info.Printf("Position update query returned profit=%v\n",previous_profit)
	return previous_profit
}
func (ss *SQLStorage) Get_trade_data(eoa_aid int64,open_positions bool) []p.PLEntry {

	var extra_condition string
	if open_positions {
		extra_condition = "(pl.closed_position=0)"
	} else {
		extra_condition = "(pl.closed_position=1)"
	}
	var query string
	query = "SELECT " +
				"pl.id," +
				"pl.market_aid," +
				"m.market_type, " +
				"pl.outcome_idx," +
				"m.outcomes," +
				"substring(extra_info::json->>'description',1,100) as descr," +
				"a.addr as mkt_addr," +
				"w_a.addr AS w_a_addr," +
				"e_a.addr AS e_a_addr," +
				"pl.time_stamp::date AS date," +
				"FLOOR(EXTRACT(EPOCH FROM pl.time_stamp))::BIGINT as created_ts," +
				"pl.net_position," +
				"pl.avg_price," +
				"pl.frozen_funds," +
				"pl.realized_profit," +
				"pl.realized_cost," +
				"pl.immediate_profit," +
				"o.order_hash," +
				"o.otype,"+
				"o.block_num," +
				"o.eoa_aid," +
				"o.eoa_fill_aid ," +
				"o.creator_eoa_addr," +
				"o.filler_eoa_addr," +
				"cf.id as cf_id," +
				"cf.final_profit, " +
				"cf.claim_status " +
			"FROM " +
				"profit_loss AS pl " +
					"LEFT JOIN address AS a ON pl.market_aid=a.address_id " +
					"LEFT JOIN address AS w_a ON pl.wallet_aid=w_a.address_id " +
					"LEFT JOIN address AS e_a ON pl.eoa_aid=e_a.address_id " +
					"LEFT JOIN market AS m ON pl.market_aid = m.market_aid " +
					"LEFT JOIN claim_funds AS cf ON (pl.market_aid=cf.market_aid AND pl.outcome_idx=cf.outcome_idx AND pl.eoa_aid=cf.eoa_aid AND pl.id=cf.last_pl_id) " +
					"LEFT JOIN LATERAL ( " +
						"SELECT mo.id,mo.order_hash,mo.otype,mo.block_num,mo.eoa_aid,mo.eoa_fill_aid," +
							"cr_a.addr AS creator_eoa_addr," +
							"fil_a.addr AS filler_eoa_addr " +
						"FROM mktord AS mo " +
							"LEFT JOIN address AS cr_a ON mo.eoa_aid = cr_a.address_id " +
							"LEFT JOIN address AS fil_a ON mo.eoa_fill_aid = fil_a.address_id " +
					") AS o ON pl.mktord_id=o.id " +
			"WHERE (pl.eoa_aid = $1) AND (pl.mktord_id>0) AND " +
			extra_condition +
			" ORDER BY pl.time_stamp"


	d_query:=strings.ReplaceAll(query,"$1",fmt.Sprintf("%v",eoa_aid))
	ss.Info.Printf("get_market_data(eoa_aid=%v,open_pos=%v): %v\n",eoa_aid,open_positions,d_query)
	rows,err := ss.db.Query(query,eoa_aid)
	if (err!=nil) {
		ss.Log_msg(fmt.Sprintf("DB error: %v (query=%v)",err,query))
		os.Exit(1)
	}
	records := make([]p.PLEntry,0,8)
	var starting_point p.PLEntry
	records = append(records,starting_point)
	var accumulator float64 = 0.0
	var otype int
	defer rows.Close()
	for rows.Next() {
		var  (
			rec p.PLEntry
			outcomes string
			cf_id sql.NullInt64
			claim_status sql.NullInt32
			cf_final_profit sql.NullFloat64
			order_hash sql.NullString
			block_num sql.NullInt64
			creator_eoa_aid sql.NullInt64
			filler_eoa_aid sql.NullInt64
			creator_addr sql.NullString
			filler_addr sql.NullString
		)
		err=rows.Scan(
			&rec.Id,
			&rec.MktAid,
			&rec.MktType,
			&rec.OutcomeIdx,
			&outcomes,
			&rec.MktDescr,
			&rec.MktAddr,
			&rec.WalletAddr,
			&rec.EOAAddr,
			&rec.Date,
			&rec.Timestamp,
			&rec.NetPosition,
			&rec.AvgPrice,
			&rec.FrozenFunds,
			&rec.RealizedProfit,
			&rec.RealizedCost,
			&rec.ImmediateProfit,
			&order_hash,
			&otype,
			&block_num,
			&creator_eoa_aid,
			&filler_eoa_aid,
			&creator_addr,
			&filler_addr,
			&cf_id,
			&cf_final_profit,
			&claim_status,
		)
		if err!=nil {
			ss.Log_msg(fmt.Sprintf("DB error: %v eoa_aid=%v q=%v",err,eoa_aid,query))
			os.Exit(1)
		}
		p.Augur_UI_price_adjustments(&rec.AvgPrice,&rec.NetPosition,rec.MktType)
		rec.CreatorBuyer = true
		rec.FillerBuyer = false
		if otype == 1 {
			rec.CreatorBuyer = false
			rec.FillerBuyer = true
		}
		rec.OutcomeStr = get_outcome_str(uint8(rec.MktType),rec.OutcomeIdx,&outcomes)
		if claim_status.Valid {
			rec.ClaimStatus = int(claim_status.Int32)
		} else {
			rec.ClaimStatus = -1	// claim funds record doesn't exist (market not finalized or simply not claimed)
		}
		if open_positions {
			accumulator = accumulator + rec.FrozenFunds
			rec.AccumFrozen = accumulator
		} else {
			if rec.ImmediateProfit == 0 {
				if cf_id.Valid {
					rec.ImmediateProfit = cf_final_profit.Float64
				}
			}
			accumulator = accumulator + rec.ImmediateProfit
			rec.AccumPl = accumulator
		}

		if order_hash.Valid { rec.OrderHash = order_hash.String }
		if block_num.Valid { rec.BlockNum = block_num.Int64 }

		rec.MktAddrSh=p.Short_address(rec.MktAddr)
		rec.EOAAddrSh=p.Short_address(rec.EOAAddr)
		rec.WalletAddrSh=p.Short_address(rec.WalletAddr)
		if creator_eoa_aid.Valid {
			if eoa_aid == creator_eoa_aid.Int64 {
				if filler_addr.Valid {
					rec.CounterPAddr = filler_addr.String
					rec.CounterPAddrSh = p.Short_address(filler_addr.String)
				}
			}
		}
		if filler_eoa_aid.Valid {
			if eoa_aid == filler_eoa_aid.Int64 {
				if creator_addr.Valid {
					rec.CounterPAddr = creator_addr.String
					rec.CounterPAddrSh = p.Short_address(creator_addr.String)
				}
			}
		}
		records = append(records,rec)
	}
	return records
}
func (ss *SQLStorage) Locate_fill_event_order(evt *p.EFill) int64 {

	var id int64 = 0
	var query string
	query = "SELECT id FROM mktord WHERE order_hash = $1"

	order_hash := common.BytesToHash(evt.OrderHash[:])
	h:= order_hash.String()
	row := ss.db.QueryRow(query,h)
	var null_id sql.NullInt64
	var err error
	err=row.Scan(&null_id);
	if (err!=nil) {
		if err == sql.ErrNoRows {
			// break
		} else {
			ss.Log_msg(fmt.Sprintf("DB Error: %v, q=%v\n",err,query))
			os.Exit(1)
		}
	} else {
		if null_id.Valid {
			id = null_id.Int64
		}
	}
	return id
}
func order_info_query() string {

	var query string
	query = "SELECT " +
				"o.id," +
				"o.order_hash," +
				"o.otype," +
				"s_w_a.addr AS s_w_a_addr," +
				"s_e_a.addr AS seller_eoa_addr," +
				"b_w_a.addr AS b_w_a_addr," +
				"b_e_a.addr AS byer_eoa_addr," +
				"CASE o.otype " +
					"WHEN 0 THEN 'BID' " +
					"ELSE 'ASK' " +
				"END AS dir, " +
				"o.time_stamp::date AS date," +
				"FLOOR(EXTRACT(EPOCH FROM o.time_stamp))::BIGINT as created_ts," +
				"o.outcome_idx," +
				"o.price AS price, " +
				"o.amount_filled AS volume, " +
				"m.market_type," +
				"m.outcomes AS outcomes_str, " +
				"ma.addr " +
			"FROM " +
				"mktord AS o " +
					"LEFT JOIN address AS a ON o.market_aid=a.address_id " +
					"LEFT JOIN address AS s_w_a ON o.wallet_aid=s_w_a.address_id " +
					"LEFT JOIN address AS s_e_a ON o.eoa_aid=s_e_a.address_id " +
					"LEFT JOIN address AS b_w_a ON o.wallet_fill_aid=b_w_a.address_id " +
					"LEFT JOIN address AS b_e_a ON o.eoa_fill_aid=b_e_a.address_id, " +
				"market AS m " +
					"LEFT JOIN address AS ma ON m.market_aid  = ma.address_id " +
			"WHERE (m.market_aid=o.market_aid)"
	return query
}
func (ss *SQLStorage) Get_order_info_by_id(order_id int64) (p.OrderInfo,error) {

	var order p.OrderInfo
	order.OrderId=order_id
	var query string
	query =  order_info_query() + " AND (o.id=$1)"
	var outcomes string
	err:=ss.db.QueryRow(query,order_id).Scan(
		&order.OrderId,
		&order.OrderHash,
		&order.OType,
		&order.CreatorWalletAddr,
		&order.CreatorEOAAddr,
		&order.FillerWalletAddr,
		&order.FillerEOAAddr,
		&order.OTypeStr,
		&order.Date,
		&order.CreatedTs,
		&order.OutcomeIdx,
		&order.Price,
		&order.Amount,
		&order.MktType,
		&outcomes,
		&order.MarketAddr,
	);
	if (err!=nil) {
		if (err==sql.ErrNoRows) {
			return order,err
		} else {
			ss.Log_msg(fmt.Sprintf("DB error looking up for Order record: %v (order_id=%v)",err,order_id))
			os.Exit(1)
		}
	}
	p.Augur_UI_price_adjustments(&order.Price,&order.Amount,int(order.MktType))
	order.CreatorBuyer = true
	order.FillerBuyer = false
	if order.OType == 1 {
		order.CreatorBuyer = false
		order.FillerBuyer = true
	}
	order.OutcomeStr = get_outcome_str(uint8(order.MktType),int(order.OutcomeIdx),&outcomes)
	order.OrderHashSh=p.Short_hash(order.OrderHash)
	order.CreatorWalletAddrSh=p.Short_address(order.CreatorWalletAddr)
	order.CreatorEOAAddrSh=p.Short_address(order.CreatorEOAAddr)
	order.FillerWalletAddrSh=p.Short_address(order.FillerWalletAddr)
	order.FillerEOAAddrSh=p.Short_address(order.FillerEOAAddr)
	order.MarketAddrSh=p.Short_address(order.MarketAddr)
	return order,nil
}
func (ss *SQLStorage) Get_filling_orders_by_hash(order_hash string) []p.OrderInfo {

	// Since 'mktord' table can contain many records with the same order hash
	//		we are sending them in an array
	var query string
	query =  order_info_query() + " AND (o.order_hash = $1) ORDER BY o.id"
	var outcomes string

	rows,err := ss.db.Query(query,order_hash)
	if (err!=nil) {
		ss.Log_msg(fmt.Sprintf("DB error: %v (query=%v) order_hash=%v",err,query,order_hash))
		os.Exit(1)
	}
	records := make([]p.OrderInfo,0,8)

	defer rows.Close()
	for rows.Next() {
		var rec p.OrderInfo
		err=rows.Scan(
			&rec.OrderId,
			&rec.OrderHash,
			&rec.OType,
			&rec.CreatorWalletAddr,
			&rec.CreatorEOAAddr,
			&rec.FillerWalletAddr,
			&rec.FillerEOAAddr,
			&rec.OTypeStr,
			&rec.Date,
			&rec.CreatedTs,
			&rec.OutcomeIdx,
			&rec.Price,
			&rec.Amount,
			&rec.MktType,
			&outcomes,
			&rec.MarketAddr,
		)
		if err!=nil {
			ss.Log_msg(fmt.Sprintf("DB error: %v",err))
			os.Exit(1)
		}
		p.Augur_UI_price_adjustments(&rec.Price,&rec.Amount,int(rec.MktType))
		rec.CreatorBuyer = true
		rec.FillerBuyer = false
		if rec.OType == 1 {
			rec.CreatorBuyer = false
			rec.FillerBuyer = true
		}
		rec.OutcomeStr = get_outcome_str(uint8(rec.MktType),int(rec.OutcomeIdx),&outcomes)
		rec.OrderHashSh=p.Short_hash(rec.OrderHash)
		rec.CreatorWalletAddrSh=p.Short_address(rec.CreatorWalletAddr)
		rec.CreatorEOAAddrSh=p.Short_address(rec.CreatorEOAAddr)
		rec.FillerWalletAddrSh=p.Short_address(rec.FillerWalletAddr)
		rec.FillerEOAAddrSh=p.Short_address(rec.FillerEOAAddr)
		records = append(records,rec)
	}
	return records
}
func (ss *SQLStorage) Get_cash_flow() []p.BlockCash {

	records := make([]p.BlockCash,0,256)
	var query string
	query = "SELECT block_num,cash_flow," +
			"EXTRACT(EPOCH FROM block.ts::TIMESTAMP)::BIGINT AS ts " +
			"FROM block WHERE cash_flow != 0 ORDER BY block_num"
	rows,err := ss.db.Query(query)
	if (err!=nil) {
		ss.Log_msg(fmt.Sprintf("DB error: %v (query=%v)",err,query))
		os.Exit(1)
	}
	defer rows.Close()
	var accumulator float64 = 0.0
	for rows.Next() {
		var bc p.BlockCash
		err = rows.Scan(&bc.BlockNum,&bc.CashFlow,&bc.Ts)
		if err != nil {
			ss.Log_msg(fmt.Sprintf("DB error: %v (query=%v)",err,query))
			os.Exit(1)
		}
		accumulator = accumulator + bc.CashFlow
		bc.AccumCashFlow = accumulator
		records = append(records,bc)
	}
	return records
}
func (ss *SQLStorage) Get_mdepth_status(market_aid int64,outcome_idx int,last_oo_id int64) (p.MktDepthStatus,error) {

	var status p.MktDepthStatus
	var query string
	query = "SELECT id FROM oorders WHERE market_aid=$1 AND outcome_idx=$2 AND id>$3"

	row := ss.db.QueryRow(query,market_aid,outcome_idx,last_oo_id)
	var null_id sql.NullInt64
	var err error
	err=row.Scan(&null_id);
	if (err!=nil) {
		if err == sql.ErrNoRows {
		} else {
			ss.Log_msg(fmt.Sprintf("Error in Get_mdepth_status() q1: %v, q=%v",err,query))
			os.Exit(1)
		}
	}
	if null_id.Valid {
		status.LastOOID=null_id.Int64
	}

	query = "SELECT count(*) AS num_rows FROM oorders WHERE market_aid=$1 AND outcome_idx=$2"
	var null_counter sql.NullInt64
	row = ss.db.QueryRow(query,market_aid,outcome_idx)
	err=row.Scan(&null_counter);
	if (err!=nil) {
		if err == sql.ErrNoRows {
			ss.Info.Printf("no rows for mdepth status for num_rows query\n")
			return status,nil
		} else {
			ss.Log_msg(fmt.Sprintf("Error in Get_mdepth_status() q2: %v, q=%v",err,query))
			os.Exit(1)
		}
	}
	if null_counter.Valid {
		status.NumOrders = null_counter.Int64
	}
	ss.Info.Printf("num_orders=%v, last_oo_id=%v\n",status.NumOrders,status.LastOOID)
	return status,nil
}
func (ss *SQLStorage) Update_oo_fillable_amount(order_hash string,cur_amount *big.Int,order *zeroex.SignedOrder) (int,string) {
	// Return value: 0 - no need to update, 1 - updated incorrect amount, 2 - order doesn't exist
	var id int64 = 0
	var order_amount string
	var query string
	query = "SELECT id,ROUND(amount*1e+18)::text AS amount FROM oorders WHERE order_hash=$1"
	row := ss.db.QueryRow(query,order_hash)
	err := row.Scan(&id,&order_amount)
	if err != nil {
		if err != sql.ErrNoRows {
			ss.Log_msg(fmt.Sprintf("Error in Scan(): %v: q=%v\n",err,query))
			os.Exit(1)
		}
		return 2,""
	}
	if order_amount != cur_amount.String() {
		query = "UPDATE oorders SET amount = ("+cur_amount.String()+"/1e+18) WHERE id=$1"
		_,err:=ss.db.Exec(query,id)
		if (err!=nil) {
			ss.Log_msg(fmt.Sprintf("DB error: %v ; q=%v",err,query))
			os.Exit(1)
		}
		return 1,order_amount
	}
	return 0,""
}
func (ss *SQLStorage) Get_all_open_order_hashes() []string {
	// Used in 0x Mesh listener to delete orders that no longer present in 0x Mesh Network

	records := make([]string,0,512)
	// open orders on 0x Mesh network
	var query string
	query = "SELECT order_hash FROM oorders"
	rows,err := ss.db.Query(query)
	if err!=nil {
		ss.Log_msg(fmt.Sprintf("DB error: %v, q=%v",err,query))
		os.Exit(1)
	}
	defer rows.Close()
	for rows.Next() {
		var order_hash string
		err=rows.Scan(&order_hash)
		if err!=nil {
			ss.Log_msg(fmt.Sprintf("DB error: %v, q=%v",err,query))
			os.Exit(1)
		}
		records = append(records,order_hash)
	}
	return records
}
func (ss *SQLStorage) Get_depth_states(market_aid int64,outcome_idx int,otype int,ts int64) []p.DepthState {

	records := make([]p.DepthState,0,32)
	var query string
		query = "SELECT " +
					"d.id,d.meshevt_id,d.market_aid,d.outcome_idx,d.otype,d.order_hash," +
					"d.price,d.amount,"+
					"FLOOR(EXTRACT(EPOCH FROM d.ini_ts))::BIGINT, " +
					"FLOOR(EXTRACT(EPOCH FROM d.fin_ts))::BIGINT, " +
					"d.ini_ts,d.fin_ts " +
				"FROM depth_state AS d " +
				"WHERE d.market_aid=$1 AND d.outcome_idx=$2 AND otype=$3 AND "+
						"(d.ini_ts <= TO_TIMESTAMP($4)) AND (TO_TIMESTAMP($4) < fin_ts) "+
				"ORDER BY ini_ts"
	rows,err := ss.db.Query(query,market_aid,outcome_idx,otype,ts)
	if (err!=nil) {
		ss.Log_msg(fmt.Sprintf("DB error: %v (query=%v)",err,query))
		os.Exit(1)
	}
	defer rows.Close()
	for rows.Next() {
		var ds p.DepthState
		err = rows.Scan(
			&ds.Id,
			&ds.MeshEvtId,
			&ds.MarketAid,
			&ds.OutcomeIdx,
			&ds.OrderType,
			&ds.OrderHash,
			&ds.Price,
			&ds.Amount,
			&ds.IniTs,
			&ds.FinTs,
			&ds.IniDate,
			&ds.FinDate,
		)
		records = append(records,ds)
	}
	return records
}
func (ss *SQLStorage) Get_price_estimate_history(market_aid int64,outcome_idx int) []p.PriceEstimate {

	records := make([]p.PriceEstimate,0,32)
	var query string
	query = "SELECT " +
				"p.id,p.market_aid,p.meshevt_id,"+
				"FLOOR(EXTRACT(EPOCH FROM p.time_stamp))::BIGINT, " +
				"p.time_stamp, " +
				"p.bid_state_id,p.ask_state_id,p.outcome_idx,"+
				"p.spread,"+
				"ROUND(p.price_est,1),ROUND(p.wprice_est,1),"+
				"p.max_bid,p.min_ask,"+
				"ROUND(wmax_bid,1),ROUND(wmin_ask,1), "+
				"e.evt_code " +
			"FROM price_estimate AS p " +
				"JOIN mesh_evt AS e ON p.meshevt_id=e.id " +
			"WHERE p.market_aid=$1 AND p.outcome_idx=$2 " +
			"ORDER BY p.time_stamp"

	rows,err := ss.db.Query(query,market_aid,outcome_idx)
	if (err!=nil) {
		ss.Log_msg(fmt.Sprintf("DB error: %v (query=%v)",err,query))
		os.Exit(1)
	}
	defer rows.Close()
	for rows.Next() {
		var pe p.PriceEstimate
		var null_bid_state,null_ask_state sql.NullInt64
		var null_wprice,null_wmax_bid,null_wmin_ask sql.NullFloat64
		err = rows.Scan(
			&pe.Id,
			&pe.MarketAid,
			&pe.MeshEvtId,
			&pe.TimeStamp,
			&pe.Date,
			&null_bid_state,
			&null_ask_state,
			&pe.OutcomeIdx,
			&pe.Spread,
			&pe.PriceEst,
			&null_wprice,
			&pe.MaxBid,
			&pe.MinAsk,
			&null_wmax_bid,
			&null_wmin_ask,
			&pe.EvtCode,
		)
		if err != nil {
			ss.Log_msg(fmt.Sprintf("DB error: %v (query=%v)",err,query))
			os.Exit(1)
		}
		if null_bid_state.Valid {
			pe.BidStateId = null_bid_state.Int64
		}
		if null_ask_state.Valid {
			pe.AskStateId = null_ask_state.Int64
		}
		if null_wprice.Valid {
			pe.WeightedPriceEst = null_wprice.Float64
		}
		if null_wmax_bid.Valid {
			pe.WMaxBid = null_wmax_bid.Float64
		}
		if null_wmin_ask.Valid {
			pe.WMinAsk = null_wmin_ask.Float64
		}
		pe.MatchingBids= ss.Get_depth_states(market_aid,outcome_idx,int(p.OrderTypeBid),pe.TimeStamp)
		pe.MatchingAsks= ss.Get_depth_states(market_aid,outcome_idx,int(p.OrderTypeAsk),pe.TimeStamp)
		records = append(records,pe)
	}
	return records
}
