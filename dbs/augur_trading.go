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
	"github.com/0xProject/0x-mesh/zeroex"
	"github.com/ethereum/go-ethereum/common"

	p "github.com/PredictionExplorer/augur-explorer/primitives"
)
func (ss *SQLStorage) Insert_market_order_evt(agtx *p.AugurTx,timestamp int64,evt *p.EOrderEvent,submitted_orders map[string]*ztypes.OrderInfo,order_specs map[string]*p.ZxMeshOrderSpec) {

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

	var aid int64;
	aid = ss.Lookup_or_create_address(evt.AddressData[0].String(),agtx.BlockNum,agtx.TxId)
	var fill_aid int64 = 0;
	if len(evt.AddressData) > 1 {
		fill_aid = ss.Lookup_or_create_address(evt.AddressData[1].String(),agtx.BlockNum,agtx.TxId)
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

	lo_price,err := ss.Get_market_lo_price(market_aid)
	if err != nil {
		ss.Log_msg(
			fmt.Sprintf("DB error: can't get lo price range for market %v : %v",market_aid,err),
		)
		os.Exit(1)
	}
	var query string
	var opcode int = p.OOOpCodeFill
	ss.Info.Printf("amount = %v, amount_filled = %v, opcode=%v\n",amount,amount_filled,opcode)
	ss.Info.Printf("OrderAction = %v, otype=%v, order_hash=%v\n",oaction,otype,order_hash)
	ss.Info.Printf("Filling existing order %v\n",order_hash)
	query = `
		INSERT INTO mktord(
			tx_id,
			market_aid,
			aid,
			fill_aid,
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
				$1,$2,$3,$4,$5,$6,$7,
				(` + fmt.Sprintf("%v",lo_price) +" + "+ price+ ")," +
				"(" + amount.String() + "/1e+18)," +
				"$8," +
				"(" + token_refund + "/1e+18)," +
				"(" + shares_refund + "/1e+18)," +
				"(" + fees + "/1e+18)," +
				"(" + amount_filled.String() + "/1e+18)," +
				"TO_TIMESTAMP($9)," +
				"$10,$11,$12,$13) RETURNING id"

	var null_id sql.NullInt64
	err=ss.db.QueryRow(query,
			agtx.TxId,
			market_aid,
			aid,
			fill_aid,
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
	ss.Insert_0x_mesh_order_event(null_id.Int64,fill_aid,timestamp,zorder,order_specs[order_hash],amount_filled,mesh_evt_code)
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
}
func (ss *SQLStorage) Delete_market_order_evt(tx_id int64) {

	var query string
	query = "DELETE FROM mktord WHERE tx_id=$1"
	_,err := ss.db.Exec(query,tx_id)
	if (err!=nil) {
		ss.Log_msg(fmt.Sprintf("DB error: %v q=%v",err,query))
		os.Exit(1)
	}
}
func (ss *SQLStorage) Update_open_order_history(mktord_id int64,order_hash string,timestamp int64,amount_filled string,opcode int) {
}
func (ss *SQLStorage) Insert_open_order(ohash *string,order *zeroex.SignedOrder,fillable_amount *big.Int,acct_addr *common.Address,ospec *p.ZxMeshOrderSpec,opcode int,evt_timestamp int64) error {
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

	aid:= ss.Lookup_or_create_address(acct_addr.String(),0,0)

	ss.Info.Printf(
		"creating open order made by %v : market=%v, price=%v, Outcome=%v, Type=%v\n",
		acct_addr.String(),ospec.Market.String(),ospec.Price.String(),ospec.Outcome,ospec.Type,
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
	query = "INSERT INTO oostats(market_aid,aid,outcome_idx) VALUES($1,$2,$3)"
	_,err = ss.db.Exec(query,market_aid,aid,ospec.Outcome)
	if err != nil {
		if !strings.Contains(err.Error(),"duplicate key value") {
			ss.Log_msg(fmt.Sprintf("DB error: %v, q=%v\n",err,query))
		}
	}
	query = "INSERT INTO oorders(" +
				"market_aid,otype,aid,price,initial_amount,amount,outcome_idx,opcode," +
				"evt_timestamp,srv_timestamp,expiration,order_hash" +
			") VALUES("+
				"$1,$2,$3,$4,"+
				initial_amount+"/1e+18," + amount+"/1e+18,"+
				"$5,$6," +
				"TO_TIMESTAMP($7),NOW(),TO_TIMESTAMP($8),$9"+
			") " +
			"ON CONFLICT DO NOTHING"
	d_query := fmt.Sprintf("INSERT INTO oorders(" +
				"market_aid,otype,aid,price,initial_amount,amount,outcome_idx,opcode," +
				"evt_timestamp,srv_timestamp,expiration,order_hash" +
			") VALUES("+
				"%v,%v,%v,%v,"+
				initial_amount+"/1e+18," + amount+"/1e+18,"+
				"%v,%v," +
				"TO_TIMESTAMP(%v),NOW(),TO_TIMESTAMP(%v),'%v'"+
			") " +
			"ON CONFLICT DO NOTHING",
			market_aid,
			otype,
			aid,
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
			aid,
			price,
			ospec.Outcome,
			opcode,
			evt_timestamp,
			expiration,
			*order_hash)
	ss.Info.Printf("Afeter oorders INSERT, err=%v\n",err)
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
		ss.Log_msg(fmt.Sprintf("DB error: couldn't insert into Open Orders table. Rows affeced = 0, order_hash=%v",*order_hash))
	}
	return errors.New("Affected rows=0")
}
func (ss *SQLStorage) Insert_cancel_open_order_evt(agtx *p.AugurTx,evt *p.ECancelZeroXOrder) {

	market_aid := ss.Lookup_address_id(evt.Market.String())
	aid := ss.Lookup_or_create_address(evt.Account.String(),agtx.BlockNum,agtx.TxId)
	ohash := common.BytesToHash(evt.OrderHash[:])
	ohash_str := ohash.String()

	var query string
	query = "INSERT INTO cancel_0x(block_num,tx_id,market_aid,aid,outcome_idx,otype,price,order_hash) " +
			"VALUES($1,$2,$3,$4,$5,($6::DECIMAL/1e+18),$7)"
	_,err := ss.db.Exec(query,
		agtx.BlockNum,
		agtx.TxId,
		market_aid,
		aid,
		evt.Outcome.Int64(),
		evt.OrderType,
		evt.Price.String(),
		ohash_str,
	)
	if err != nil {
		ss.Log_msg(fmt.Sprintf("DB error: %v; q=%v",err,query))
		os.Exit(1)
	}
}
func (ss *SQLStorage) Delete_cancel_open_order_evt(tx_id int64) {

	var query string
	query = "DELETE FROM cancel_0x WHERE tx_id=$1"
	_,err := ss.db.Exec(query,tx_id)
	if (err!=nil) {
		ss.Log_msg(fmt.Sprintf("DB error: %v q=%v",err,query))
		os.Exit(1)
	}
}
func (ss *SQLStorage) Cancel_open_order(aid int64,orders map[string]*ztypes.OrderInfo,order_specs map[string]*p.ZxMeshOrderSpec,order_hash string,timestamp int64) {

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
	var query string
	ss.Insert_0x_mesh_order_event(0,aid,timestamp,oinfo,ospec,nil,p.MeshEvtCancelled)

	query = "DELETE FROM oorders WHERE order_hash = $1"
	result,err := ss.db.Exec(query,order_hash)
	if err!=nil {
		ss.Info.Printf(fmt.Sprintf("DB error: couldn't delete open order with order_hash = %v: %v;  q=%v\n",order_hash,err,query))
		if result == nil {
			ss.Log_msg(
				fmt.Sprintf(
					"DB error: couldn't delete open order with order_hash = %v, q=%v\n",
					order_hash,query,
				),
			)
			os.Exit(1)
		}
	}
	rows_affected,err:=result.RowsAffected()
	if rows_affected == 0  {
		ss.Info.Printf(fmt.Sprintf("DB error: couldn't delete open order with order_hash = %v (not found)\n",order_hash))
	}
}
func (ss *SQLStorage) Delete_open_0x_order(order_hash string,timestamp int64,opcode int) {

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
func (ss *SQLStorage) Open_order_exists(oo_hash string) bool {

	var query string
	query = "SELECT id FROM oorders WHERE order_hash=$1"
	row := ss.db.QueryRow(query,oo_hash)
	var null_id sql.NullInt64
	err := row.Scan(&null_id);
	if (err!=nil) {
		if err == sql.ErrNoRows {
			return false
		} else {
			ss.Log_msg(fmt.Sprintf("DB error: %v, q=%v",err,query))
			os.Exit(1)
		}
	}
	return true
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
				"m.decimals," +
				"m.outcomes AS outcomes_str " +
			"FROM mktord AS o " +
				"LEFT JOIN address AS a ON o.market_aid=a.address_id " +
				"LEFT JOIN address AS fa ON o.fill_aid=fa.address_id " +
				"LEFT JOIN address AS ca ON o.aid=ca.address_id " +
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
				"m.decimals," +
				"ua.addr AS user_addr," +
				"o.srv_timestamp::date AS date_created," +
				"o.expiration::date AS expires," +
				"FLOOR(EXTRACT(EPOCH FROM o.expiration))::BIGINT as expires_ts," +
				"o.price AS price, " +
				"o.amount AS amount," +
				"o.order_hash," +
				"s.num_bids," +
				"s.num_asks," +
				"s.num_cancel " +
			"FROM oorders AS o " +
				"JOIN market AS m ON m.market_aid=o.market_aid " +
				"LEFT JOIN oostats AS s ON (" +
						"o.market_aid=s.market_aid AND " +
						"o.aid=s.aid AND " +
						"o.outcome_idx=s.outcome_idx" +
				") " +
				"LEFT JOIN address AS a ON o.market_aid=a.address_id " +
				"LEFT JOIN address AS ua ON o.aid=ua.address_id " +
			"WHERE o.market_aid = $1 AND o.outcome_idx=$2 AND o.otype = $3 " +
			"ORDER BY "
	if otype == p.OrderTypeBid {
				query = query + "o.price DESC,o.evt_timestamp DESC";
	} else {
				query = query + "o.price ASC,o.evt_timestamp DESC";
	}
	//ss.Info.Printf("q=%v\n",query)
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
		var mkt_type,decimals int
		var rec p.DepthEntry
		var num_bids sql.NullInt64
		var num_asks sql.NullInt64
		var num_cancels sql.NullInt64
		err=rows.Scan(
			&oo_id,
			&rec.MktAid,
			&rec.OutcomeIdx,
			&mkt_type,
			&decimals,
			&rec.Addr,
			&rec.DateCreated,
			&rec.Expires,
			&rec.ExpiresTs,
			&rec.Price,
			&rec.Volume,
			&rec.OrderHash,
			&num_bids,
			&num_asks,
			&num_cancels,
		)
		ss.Info.Printf("Addr=%v (ooid=%v, hash=%v)\n",rec.Addr,oo_id,rec.OrderHash)
		if err!=nil {
			ss.Log_msg(fmt.Sprintf("DB error: %v, q=%v",err,query))
			os.Exit(1)
		}
		p.Augur_UI_price_adjustments(&rec.Price,&rec.Volume,mkt_type,decimals)
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
		rec.AddrSh=p.Short_address(rec.Addr)
		records = append(records,rec)
		if max_id < oo_id {
			max_id = oo_id
		}
	}
	return records,max_id
}
func (ss *SQLStorage) Get_price_history_for_outcome(market_aid int64,outc int,low_price_limit float64) []p.OrderInfo {

	var query string
	query = "SELECT " +
				"o.id,"+
				"o.order_hash," +
				"o.market_aid," +
				"m.market_type," +
				"m.decimals," +
				"ca.addr AS creator_addr," +
				"fa.addr AS filler_addr," +
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
				"LEFT JOIN address AS ca ON o.aid=ca.address_id " +
				"LEFT JOIN address AS fa ON o.fill_aid=fa.address_id " +
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
		var mkt_type,decimals int
		err=rows.Scan(
			&rec.OrderId,
			&rec.OrderHash,
			&rec.MktAid,
			&mkt_type,
			&decimals,
			&rec.CreatorAddr,
			&rec.FillerAddr,
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
		p.Augur_UI_price_adjustments(&rec.Price,&rec.Amount,mkt_type,decimals)
		if rec.OutcomeIdx != 0 {
			rec.Price = rec.Price + low_price_limit
		}
		rec.CreatorBuyer = true
		rec.FillerBuyer = false
		if rec.OType == 1 {
			rec.CreatorBuyer = false
			rec.FillerBuyer = true
		}
		rec.CreatorAddrSh=p.Short_address(rec.CreatorAddr)
		rec.FillerAddrSh=p.Short_address(rec.FillerAddr)
		accumulated_volume = accumulated_volume + rec.Amount
		rec.AccumVol = accumulated_volume
		records = append(records,rec)
	}
	return records
}
func (ss *SQLStorage) Get_full_price_history(mkt_addr string,market_aid int64,low_price_limit float64) p.FullPriceHistory {

	var output p.FullPriceHistory
	outcomes,_ := ss.Get_outcome_volumes(mkt_addr,market_aid,0,low_price_limit);
	for _,outc := range outcomes {
		var ph p.PriceHistory
		ph.OutcomeIdx = outc.Outcome
		ph.OutcomeStr = outc.OutcomeStr
		ph.Trades = ss.Get_price_history_for_outcome(market_aid,outc.Outcome,low_price_limit)
		output.Outcomes = append(output.Outcomes,ph)
	}
	return output
}
func (ss *SQLStorage) Get_zoomed_t1_price_history_for_outcome(market_aid int64,mkt_type,decimals int,outc int,init_ts int,fin_ts int) []p.ZHistT1Entry {

	var query string
	query = "SELECT " +
				"p.id,"+
				"e.order_hash," +
				"p.market_aid," +
				"e.maker_addr AS maker_addr," +
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
			"WHERE " +
				"p.market_aid = $1 AND " +
				"p.outcome_idx=$2 AND " +
				"p.time_stamp >= TO_TIMESTAMP($3) AND "+
				"p.time_stamp < TO_TIMESTAMP($4) " +
			"ORDER BY p.time_stamp"
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
			&rec.MakerAddr,
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
		rec.MakerAddrSh=p.Short_address(rec.MakerAddr)
		p.Augur_UI_price_adjustments(&rec.Price,&rec.FillableAmount,mkt_type,decimals)
		p.Augur_UI_price_adjustments(&rec.PriceEstimate,&rec.Amount,mkt_type,decimals)
		p.Augur_UI_price_adjustments(&rec.WeightedPriceEst,nil,mkt_type,decimals)
		p.Augur_UI_price_adjustments(&rec.MaxBid,nil,mkt_type,decimals)
		p.Augur_UI_price_adjustments(&rec.MinAsk,nil,mkt_type,decimals)
		p.Augur_UI_price_adjustments(&rec.WMaxBid,nil,mkt_type,decimals)
		p.Augur_UI_price_adjustments(&rec.WMinAsk,nil,mkt_type,decimals)
		p.Augur_UI_price_adjustments(&rec.Spread,nil,mkt_type,decimals)
		records = append(records,rec)
	}
	return records
}
func (ss *SQLStorage) Get_zoomed_t2_price_history_for_outcome(market_aid int64,mkt_type,decimals int,outc int,init_ts int,fin_ts int,interval int,low_price_limit float64) []p.ZHistT2Entry {

	var query string
	query = "WITH periods AS (" +
				"SELECT * FROM (" +
					"SELECT " +
						"generate_series AS start_ts,"+
						"TO_TIMESTAMP(EXTRACT(EPOCH FROM generate_series) + $3) AS end_ts "+
					"FROM (" +
						"SELECT * " +
							"FROM generate_series(" +
								"TO_TIMESTAMP($1)," +
								"TO_TIMESTAMP($2)," +
								"TO_TIMESTAMP($3)-TO_TIMESTAMP(0)) " +
					") AS i" +
				") AS data " +
			") " +
			"SELECT " +
				"COALESCE(COUNT(pr.id),0) as num_rows, " +
				"ROUND(FLOOR(EXTRACT(EPOCH FROM start_ts)))::BIGINT as start_ts," +
				"AVG(price_est) AS avg_price_estimate," +
				"AVG(wprice_est) AS avg_wpe " +
			"FROM periods AS p " +
				"LEFT JOIN price_estimate AS pr ON (" +
					"p.start_ts <= pr.time_stamp AND "+
					"pr.time_stamp < p.end_ts AND " +
					"market_aid=$4 AND " +
					"outcome_idx=$5" +
			") " +
			"GROUP BY start_ts " +
			"ORDER BY start_ts"

	rows,err := ss.db.Query(query,init_ts,fin_ts,interval,market_aid,outc)
	if (err!=nil) {
		ss.Log_msg(fmt.Sprintf("DB error: %v (query=%v)",err,query))
		os.Exit(1)
	}
	records := make([]p.ZHistT2Entry,0,8)
	var last_rec p.ZHistT2Entry
	defer rows.Close()
	for rows.Next() {
		var rec p.ZHistT2Entry
		var null_pest,null_wpest sql.NullFloat64
		var num_rows int
		err=rows.Scan(
			&num_rows,
			&rec.Timestamp,
			&null_pest,
			&null_wpest,
		)
		if err!=nil {
			ss.Log_msg(fmt.Sprintf("DB error: %v",err))
			os.Exit(1)
		}
		if null_pest.Valid {
			rec.PriceEstimate = null_pest.Float64
		}
		if null_wpest.Valid {
			rec.WeightedPriceEstimate = null_wpest.Float64
		}
		p.Augur_UI_price_adjustments(&rec.PriceEstimate,nil,mkt_type,decimals)
		p.Augur_UI_price_adjustments(&rec.WeightedPriceEstimate,nil,mkt_type,decimals)
		if num_rows == 0 {
			rec.PriceEstimate = last_rec.PriceEstimate
			rec.WeightedPriceEstimate = last_rec.WeightedPriceEstimate
		} else {
			last_rec.PriceEstimate = rec.PriceEstimate
			last_rec.WeightedPriceEstimate = rec.WeightedPriceEstimate
		}
		if outc != 0 {
			rec.PriceEstimate = rec.PriceEstimate + low_price_limit
			rec.WeightedPriceEstimate = rec.WeightedPriceEstimate + low_price_limit
		}
		records = append(records,rec)
	}
	return records
}
func (ss *SQLStorage) Get_zoomed_price_history(mkt_addr string,market_aid int64,init_ts int,fin_ts int,interval int) p.FullZoomedPriceHist {

	var output p.FullZoomedPriceHist
	var query string
	if fin_ts == 2147483647 {
		query = "SELECT  EXTRACT(EPOCH FROM time_stamp)::BIGINT AS ending_ts "+
				"FROM price_estimate " +
				"WHERE market_aid=$1" +
				"ORDER BY ending_ts DESC LIMIT 1"

		var null_ts sql.NullInt64
		err := ss.db.QueryRow(query,market_aid).Scan(&null_ts)
		ss.adjust_ts(&fin_ts,err,&null_ts)
		fin_ts++
	}
	if init_ts == 0 {
		query = "SELECT  EXTRACT(EPOCH FROM time_stamp)::BIGINT AS starting_ts "+
				"FROM price_estimate " +
				"WHERE market_aid=$1" +
				"ORDER BY starting_ts LIMIT 1"
		var null_ts sql.NullInt64
		err := ss.db.QueryRow(query,market_aid).Scan(&null_ts)
		ss.adjust_ts(&init_ts,err,&null_ts)

	}
	query = "SELECT " +
				"lo_price " +
			"FROM market " +
			"WHERE market_aid = $1"

	row := ss.db.QueryRow(query,market_aid)
	var low_price_limit float64
	err := row.Scan(&low_price_limit);
	if (err!=nil) {
		if err == sql.ErrNoRows {
			return output
		} else {
			ss.Log_msg(fmt.Sprintf("DB error: %v, q=%v",err,query))
			os.Exit(1)
		}
	}

	init_ts = init_ts / interval
	init_ts = init_ts * interval
	mkt_type,_,decimals,err := ss.get_market_type_and_ticks(market_aid)
	if err != nil {
		ss.Log_msg(fmt.Sprintf("Aborting Get_zoomed_price_history() call, market %v not found\n",market_aid))
		return output
	}
	outcomes,_ := ss.Get_outcome_volumes(mkt_addr,market_aid,0,low_price_limit);
	for _,outc := range outcomes {
		var ph p.ZoomedPriceHist
		ph.OutcomeIdx = outc.Outcome
		ph.OutcomeStr = outc.OutcomeStr
		ph.InitTs = init_ts
		ph.FinTs = fin_ts
		ph.IntervalSecs = interval
		ph.Type2Entries = ss.Get_zoomed_t2_price_history_for_outcome(
			market_aid,
			mkt_type,
			decimals,
			outc.Outcome,
			init_ts,
			fin_ts,
			interval,
			low_price_limit,
		)
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
func (ss *SQLStorage) Get_open_positions(aid int64) []p.PLEntry {
	return ss.Get_trade_data(aid,true)
}
func (ss *SQLStorage) close_previous_positions(market_aid int64,aid int64,outcome_idx int,profit_loss string) string {

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
						"(aid = $2) AND " +
						"(outcome_idx = $3) AND " +
						"(closed_position = 0) " +
				"RETURNING round(realized_profit*1e+36)::text"
	var previous_profit string
	row:=ss.db.QueryRow(query,market_aid,aid,outcome_idx)
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
func (ss *SQLStorage) Get_trade_data(aid int64,open_positions bool) []p.PLEntry {

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
				"m.decimals," +
				"pl.outcome_idx," +
				"m.outcomes," +
				"substring(extra_info::json->>'description',1,100) as descr," +
				"ma.addr as mkt_addr," +
				"aa.addr AS act_addr," +
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
				"o.aid," +
				"o.fill_aid ," +
				"o.creator_addr," +
				"o.filler_addr," +
				"cf.id as cf_id," +
				"cf.final_profit, " +
				"cf.claim_status " +
			"FROM " +
				"profit_loss AS pl " +
					"LEFT JOIN address AS ma ON pl.market_aid=ma.address_id " +
					"LEFT JOIN address AS aa ON pl.aid=aa.address_id " +
					"LEFT JOIN market AS m ON pl.market_aid = m.market_aid " +
					"LEFT JOIN claim_funds AS cf ON (pl.market_aid=cf.market_aid AND pl.outcome_idx=cf.outcome_idx AND pl.aid=cf.aid AND pl.id=cf.last_pl_id) " +
					"LEFT JOIN LATERAL ( " +
						"SELECT mo.id,mo.order_hash,mo.otype,mo.block_num,mo.aid,mo.fill_aid," +
							"cr_a.addr AS creator_addr," +
							"fil_a.addr AS filler_addr " +
						"FROM mktord AS mo " +
							"LEFT JOIN address AS cr_a ON mo.aid = cr_a.address_id " +
							"LEFT JOIN address AS fil_a ON mo.fill_aid = fil_a.address_id " +
					") AS o ON pl.mktord_id=o.id " +
			"WHERE (pl.aid = $1) AND (pl.mktord_id>0) AND " +
			extra_condition +
			" ORDER BY pl.time_stamp"


	rows,err := ss.db.Query(query,aid)
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
			decimals int
			cf_id sql.NullInt64
			claim_status sql.NullInt32
			cf_final_profit sql.NullFloat64
			order_hash sql.NullString
			block_num sql.NullInt64
			creator_aid sql.NullInt64
			filler_aid sql.NullInt64
			creator_addr sql.NullString
			filler_addr sql.NullString
		)
		err=rows.Scan(
			&rec.Id,
			&rec.MktAid,
			&rec.MktType,
			&decimals,
			&rec.OutcomeIdx,
			&outcomes,
			&rec.MktDescr,
			&rec.MktAddr,
			&rec.Addr,
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
			&creator_aid,
			&filler_aid,
			&creator_addr,
			&filler_addr,
			&cf_id,
			&cf_final_profit,
			&claim_status,
		)
		if err!=nil {
			ss.Log_msg(fmt.Sprintf("DB error: %v aid=%v q=%v",err,aid,query))
			os.Exit(1)
		}
		p.Augur_UI_price_adjustments(&rec.AvgPrice,&rec.NetPosition,rec.MktType,decimals)
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
		rec.AddrSh=p.Short_address(rec.Addr)
		if creator_aid.Valid {
			if aid == creator_aid.Int64 {
				if filler_addr.Valid {
					rec.CounterPAddr = filler_addr.String
					rec.CounterPAddrSh = p.Short_address(filler_addr.String)
				}
			}
		}
		if filler_aid.Valid {
			if aid == filler_aid.Int64 {
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
				"sa.addr AS seller_eoa_addr," +
				"ba.addr AS byer_eoa_addr," +
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
				"m.decimals," +
				"m.outcomes AS outcomes_str, " +
				"ma.addr " +
			"FROM " +
				"mktord AS o " +
					"LEFT JOIN address AS a ON o.market_aid=a.address_id " +
					"LEFT JOIN address AS sa ON o.aid=sa.address_id " +
					"LEFT JOIN address AS ba ON o.fill_aid=ba.address_id, " +
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
	var decimals int
	err:=ss.db.QueryRow(query,order_id).Scan(
		&order.OrderId,
		&order.OrderHash,
		&order.OType,
		&order.CreatorAddr,
		&order.FillerAddr,
		&order.OTypeStr,
		&order.Date,
		&order.CreatedTs,
		&order.OutcomeIdx,
		&order.Price,
		&order.Amount,
		&order.MktType,
		&decimals,
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
	p.Augur_UI_price_adjustments(&order.Price,&order.Amount,int(order.MktType),decimals)
	order.CreatorBuyer = true
	order.FillerBuyer = false
	if order.OType == 1 {
		order.CreatorBuyer = false
		order.FillerBuyer = true
	}
	order.OutcomeStr = get_outcome_str(uint8(order.MktType),int(order.OutcomeIdx),&outcomes)
	order.OrderHashSh=p.Short_hash(order.OrderHash)
	order.CreatorAddrSh=p.Short_address(order.CreatorAddr)
	order.FillerAddrSh=p.Short_address(order.FillerAddr)
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
		var decimals int
		err=rows.Scan(
			&rec.OrderId,
			&rec.OrderHash,
			&rec.OType,
			&rec.CreatorAddr,
			&rec.FillerAddr,
			&rec.OTypeStr,
			&rec.Date,
			&rec.CreatedTs,
			&rec.OutcomeIdx,
			&rec.Price,
			&rec.Amount,
			&rec.MktType,
			&decimals,
			&outcomes,
			&rec.MarketAddr,
		)
		if err!=nil {
			ss.Log_msg(fmt.Sprintf("DB error: %v",err))
			os.Exit(1)
		}
		p.Augur_UI_price_adjustments(&rec.Price,&rec.Amount,int(rec.MktType),decimals)
		rec.CreatorBuyer = true
		rec.FillerBuyer = false
		if rec.OType == 1 {
			rec.CreatorBuyer = false
			rec.FillerBuyer = true
		}
		rec.OutcomeStr = get_outcome_str(uint8(rec.MktType),int(rec.OutcomeIdx),&outcomes)
		rec.OrderHashSh=p.Short_hash(rec.OrderHash)
		rec.CreatorAddrSh=p.Short_address(rec.CreatorAddr)
		rec.FillerAddrSh=p.Short_address(rec.FillerAddr)
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
func (ss *SQLStorage) Update_oo_fillable_amount(order_hash string,order *zeroex.SignedOrder) {
	// Return value: 0 - no need to update, 1 - updated incorrect amount, 2 - order doesn't exist
	var order_amount string
	var query string
	ss.Info.Printf("Updating fillable amount for ohash %v\n",order_hash)
	query = "SELECT ROUND(amount_fill*1e+18)::text AS amount FROM mesh_evt AS e " +
				"WHERE order_hash=$1 AND " +
					"((e.evt_code=3) OR (e.evt_code=4))"
	row := ss.db.QueryRow(query,order_hash)
	err := row.Scan(&order_amount)
	if err != nil {
		if err != sql.ErrNoRows {
			ss.Log_msg(fmt.Sprintf("Error in Scan(): %v: q=%v\n",err,query))
			os.Exit(1)
		}
		return
	}
	query = "UPDATE oorders SET amount = initial_amount - ("+order_amount+"/1e+18) WHERE order_hash=$1"
	ss.Info.Printf("Fillable amount for %v is now %v (query=%v)\n",order_hash,order_amount,query)
	_,err = ss.db.Exec(query,order_hash)
	if (err!=nil) {
		ss.Log_msg(fmt.Sprintf("DB error: %v ; q=%v",err,query))
		os.Exit(1)
	}
}
func (ss *SQLStorage) Get_all_open_order_hashes() ([]string,[]int64) {
	// Used in 0x Mesh listener to delete orders that no longer present in 0x Mesh Network

	records_hashes := make([]string,0,512)
	records_expirations := make([]int64,0,512)
	// open orders on 0x Mesh network
	var query string
	query = "SELECT order_hash,FLOOR(EXTRACT(EPOCH FROM expiration))::BIGINT  FROM oorders"
	rows,err := ss.db.Query(query)
	if err!=nil {
		ss.Log_msg(fmt.Sprintf("DB error: %v, q=%v",err,query))
		os.Exit(1)
	}
	defer rows.Close()
	for rows.Next() {
		var order_hash string
		var order_expiration int64
		err=rows.Scan(&order_hash,&order_expiration)
		if err!=nil {
			ss.Log_msg(fmt.Sprintf("DB error: %v, q=%v",err,query))
			os.Exit(1)
		}
		records_hashes = append(records_hashes,order_hash)
		records_expirations = append(records_expirations,order_expiration)
	}
	return records_hashes,records_expirations
}
func (ss *SQLStorage) Get_depth_states(market_aid int64,mkt_type,decimals int,outcome_idx int,otype int,ts int64) []p.DepthState {

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
				"ORDER BY d.price DESC, ini_ts"
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
		p.Augur_UI_price_adjustments(&ds.Price,&ds.Amount,mkt_type,decimals)
		records = append(records,ds)
	}
	return records
}
func (ss *SQLStorage) Get_price_estimate_history(market_aid int64,outcome_idx int) []p.PriceEstimate {

	records := make([]p.PriceEstimate,0,32)

	mkt_type,_,decimals,err := ss.get_market_type_and_ticks(market_aid)
	if err != nil {
		ss.Log_msg(fmt.Sprintf("Get_price_estimate_history() aborted: market_aid=%v not found",market_aid))
		return records
	}

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
		p.Augur_UI_price_adjustments(&pe.PriceEst,nil,mkt_type,decimals)
		p.Augur_UI_price_adjustments(&pe.MaxBid,nil,mkt_type,decimals)
		p.Augur_UI_price_adjustments(&pe.MinAsk,nil,mkt_type,decimals)
		p.Augur_UI_price_adjustments(&pe.WeightedPriceEst,nil,mkt_type,decimals)
		p.Augur_UI_price_adjustments(&pe.WMaxBid,nil,mkt_type,decimals)
		p.Augur_UI_price_adjustments(&pe.WMinAsk,nil,mkt_type,decimals)
		pe.MatchingBids= ss.Get_depth_states(market_aid,mkt_type,decimals,outcome_idx,int(p.OrderTypeBid),pe.TimeStamp)
		pe.MatchingAsks= ss.Get_depth_states(market_aid,mkt_type,decimals,outcome_idx,int(p.OrderTypeAsk),pe.TimeStamp)
		records = append(records,pe)
	}
	return records
}
func calc_oi_total(oi_map map[int64]float64) float64 {

	var sum float64 = 0.0
	for _,oi_value := range oi_map {
		sum = sum + oi_value
	}
	return sum
}
func (ss *SQLStorage) Get_accumulated_open_interest_all_markets_v3(init_ts int,fin_ts int,interval int) []p.OIAccum {
	// Return value: given initial and finel timestamps and the interval, this function splits the
	//		open interest data in equal intervals and for each interval it returns open interst
	//		for all markets althogether (the sum of open intersest) in an array
	// Tech Notes:
	//		Open Interest is a value that is already being accumulated
	//		The process is made of 2 subprocesses:
	//			1. First. We select the last open interest records that exist before 'init_ts' and these 
	//					values are going to be used as initial values, in case data point for the 
	//					corresponding period isn't returned by the query
	//			2. After that we query records within the period (init_ts-fin_ts) tobuild the data. 
	//				If no data exist for a particular market we pick it from the query we did on step 1
	//		The overall process takes the open interest value at the end of the interval (it is not AVG)
	//		Autofills intervals if there is no data by copying open interest from previous interval
	//			using the 'for' loop
	// 		This process could be implemented with genetate_series() function and remove the 'for' loop
	//			that does the autofill, but for loop is faster
	//		The hashmap is required because open interest stored in the database is an accumulted value
	//			(it is being accumulated by Augur) so we can't replace it with just one variable where
	//			we could track total open interest for all markets

	output := make([]p.OIAccum,0,512)
	var query string

	if fin_ts == 2147483647 {
		query = "SELECT  EXTRACT(EPOCH FROM ts_inserted)::BIGINT AS ending_ts "+
				"FROM oi_chg " +
				"ORDER BY ending_ts DESC LIMIT 1"

		var null_ts sql.NullInt64
		err := ss.db.QueryRow(query).Scan(&null_ts)
		ss.adjust_ts(&fin_ts,err,&null_ts)
		fin_ts++ // +1 second because we have our comparison operator '<', not '<=' in SELECT
	}
	if init_ts == 0 {
		query = "SELECT  EXTRACT(EPOCH FROM ts_inserted)::BIGINT AS starting_ts "+
				"FROM oi_chg " +
				"ORDER BY starting_ts LIMIT 1"
		var null_ts sql.NullInt64
		err := ss.db.QueryRow(query).Scan(&null_ts)
		ss.adjust_ts(&init_ts,err,&null_ts)
		if null_ts.Valid {
			init_ts = int(null_ts.Int64)
		} else {
			init_ts = 1595894400 // 28 July 2020 GMT-0  (Augur Release date)
		}
	}
	// align timestamp to interval boundaries:
	init_ts = init_ts / interval
	init_ts = init_ts * interval

	query = "SELECT oi,market_aid,id,EXTRACT(EPOCH FROM ts_inserted)::BIGINT " +
			"FROM oi_chg " +
			"WHERE id IN (" +
				"SELECT MAX(id) " +
				"FROM oi_chg " +
				"WHERE ts_inserted < TO_TIMESTAMP($1) " + // Note the '<' rows fetched must not intersect with the query in Step 2
				"GROUP BY market_aid " +
			")"
	rows,err := ss.db.Query(query,init_ts)
	if (err!=nil) {
		ss.Log_msg(fmt.Sprintf("DB error: %v (query=%v)",err,query))
		os.Exit(1)
	}

	oi_map := make(map[int64]float64)
	
	// build map of initial open interest values , these values are the last entries found
	// before our 'init_ts'.
	defer rows.Close()
	for rows.Next() {
		var market_aid int64
		var oi float64
		var id,ts int64
		err = rows.Scan(&oi,&market_aid,&id,&ts)
		if err != nil {
			ss.Log_msg(fmt.Sprintf("DB error: %v (query=%v)",err,query))
			os.Exit(1)
		}
		oi_map[market_aid] = oi
	}

	// Step 2 . Fetch available data points whitin interval length
	query = "SELECT " +
				"id," +
				"market_aid," +
				"ROUND(FLOOR(EXTRACT(EPOCH FROM ts_inserted)))::BIGINT AS ts," +
				"oi " +
			"FROM oi_chg " +
			"WHERE " +
				"ts_inserted >= TO_TIMESTAMP($1) AND "+
				"ts_inserted < TO_TIMESTAMP($2) " +
			"ORDER BY ts"

	rows,err = ss.db.Query(query,init_ts,fin_ts)
	if (err!=nil) {
		ss.Log_msg(fmt.Sprintf("DB error: %v (query=%v)",err,query))
		os.Exit(1)
	}
	var cur_ts int = init_ts
	defer rows.Close()
	for rows.Next() {
		var row p.OIAccum
		var ts int
		var market_aid int64
		var oi float64
		var id int64
		err = rows.Scan(&id,&market_aid,&ts,&oi)
		if err != nil {
			ss.Log_msg(fmt.Sprintf("DB error: %v (query=%v)",err,query))
			os.Exit(1)
		}
		for {
			if ts >= (cur_ts + interval)  {
				sum_oi := calc_oi_total(oi_map)
				row.AccumOpenInterest = sum_oi
				row.TimeStamp = cur_ts
				output = append(output,row)
				cur_ts = cur_ts + interval
			} else {
				break
			}
		}
		oi_map[market_aid] = oi
	}
	var row p.OIAccum
	sum_oi := calc_oi_total(oi_map)
	row.AccumOpenInterest = sum_oi
	row.TimeStamp = cur_ts
	output = append(output,row)
	return output
}
func (ss *SQLStorage) adjust_ts(ts_ptr *int,err error,field *sql.NullInt64) {
	// sets ts_ptr to value stored in the DB or, sets NOW()
	var query string
	if (err!=nil) {
		if err==sql.ErrNoRows {
			query = "SELECT EXTRACT(EPOCH FROM NOW())::BIGINT as cur_ts "
			var null_ts sql.NullInt64
			err=ss.db.QueryRow(query).Scan(&null_ts)
			if (err!=nil) {
				ss.Log_msg(fmt.Sprintf("DB error: %v",err))
				os.Exit(1)
			} else {
				*ts_ptr = int(null_ts.Int64)
			}
		} else {
			ss.Log_msg(fmt.Sprintf("DB error: %v",err))
			os.Exit(1)
		}
	} else {
		*ts_ptr = int(field.Int64)
	}
}
func (ss *SQLStorage) Get_accumulated_trades_all_markets(init_ts int,fin_ts int,interval int) []p.TradesByInterval {

	output := make([]p.TradesByInterval,0,512)

	var query string
	if fin_ts == 2147483647 {
		query = "SELECT  EXTRACT(EPOCH FROM time_stamp)::BIGINT AS ending_ts "+
				"FROM mktord " +
				"ORDER BY ending_ts DESC LIMIT 1"

		var null_ts sql.NullInt64
		err := ss.db.QueryRow(query).Scan(&null_ts)
		ss.adjust_ts(&fin_ts,err,&null_ts)
		fin_ts++
	}
	if init_ts == 0 {
		query = "SELECT  EXTRACT(EPOCH FROM time_stamp)::BIGINT AS starting_ts "+
				"FROM mktord " +
				"ORDER BY starting_ts LIMIT 1"
		var null_ts sql.NullInt64
		err := ss.db.QueryRow(query).Scan(&null_ts)
		ss.adjust_ts(&init_ts,err,&null_ts)
	}
	init_ts = init_ts / interval
	init_ts = init_ts * interval

	var accum_num_trades int64 = 0
	var accum_volume float64 = 0.0

	query = "SELECT count(*) as num_trades,SUM(amount_filled*price) AS volume " +
			"FROM mktord WHERE time_stamp < TO_TIMESTAMP($1)"
	var null_num_trades sql.NullInt64
	var null_volume sql.NullFloat64
	err := ss.db.QueryRow(query,init_ts).Scan(&null_num_trades,&null_volume);
	if (err!=nil) {
		ss.Log_msg(	fmt.Sprintf("DB error: %v, q=%v",err,query))
		os.Exit(1)
	}
	if null_num_trades.Valid {
		accum_num_trades = null_num_trades.Int64
	}
	if null_volume.Valid {
		accum_volume = null_volume.Float64
	}

	query = "WITH periods AS (" +
				"SELECT * FROM (" +
					"SELECT " +
						"generate_series AS start_ts," +
						"TO_TIMESTAMP(EXTRACT(EPOCH FROM generate_series) + $3) AS end_ts " +
					"FROM (" +
						"SELECT * " +
							"FROM generate_series(" +
								"TO_TIMESTAMP($1)," +
								"TO_TIMESTAMP($2)," +
								"TO_TIMESTAMP($3)-TO_TIMESTAMP(0)" +
							") " +
					") AS i" +
				") AS data " +
			")" +
			"SELECT " +
				"ROUND(FLOOR(EXTRACT(EPOCH FROM start_ts)))::BIGINT as start_ts," +
				"SUM(amount_filled*price) AS volume," +
				"COALESCE(COUNT(o.id),0) as num_trades " +
			"FROM periods AS p " +
				"LEFT JOIN mktord AS o ON (" +
					"p.start_ts <= o.time_stamp AND "+
					"o.time_stamp < p.end_ts " +
			") " +
			"GROUP BY start_ts " +
			"ORDER BY start_ts"
	d_query := fmt.Sprintf("WITH periods AS (" +
				"SELECT * FROM (" +
					"SELECT " +
						"generate_series AS start_ts," +
						"TO_TIMESTAMP(EXTRACT(EPOCH FROM generate_series) + %v) AS end_ts " +
					"FROM (" +
						"SELECT * " +
							"FROM generate_series(" +
								"TO_TIMESTAMP(%v)," +
								"TO_TIMESTAMP(%v)," +
								"TO_TIMESTAMP(%v)-TO_TIMESTAMP(0)" +
							") " +
					") AS i" +
				") AS data " +
			")" +
			"SELECT " +
				"ROUND(FLOOR(EXTRACT(EPOCH FROM start_ts)))::BIGINT as start_ts," +
				"SUM(amount_filled*price) AS volume," +
				"count(*) as num_trades " +
			"FROM periods AS p " +
				"LEFT JOIN mktord AS o ON (" +
					"p.start_ts <= o.time_stamp AND "+
					"o.time_stamp < p.end_ts " +
			") " +
			"GROUP BY start_ts " +
			"ORDER BY start_ts",
			interval,init_ts,fin_ts,interval)
	ss.Info.Printf("query=%v\n",d_query)
	rows,err := ss.db.Query(query,init_ts,fin_ts,interval)
	if (err!=nil) {
		ss.Log_msg(fmt.Sprintf("DB error: %v (query=%v)",err,query))
		os.Exit(1)
	}
	defer rows.Close()
	for rows.Next() {
		var rec p.TradesByInterval
		var null_vol sql.NullFloat64
		var null_trades sql.NullInt64
		err = rows.Scan(&rec.TimeStamp,&null_vol,&null_trades)
		if err != nil {
			ss.Log_msg(fmt.Sprintf("DB error: %v (query=%v)",err,query))
			os.Exit(1)
		}
		if null_vol.Valid {
			rec.Volume = null_vol.Float64
		}
		if null_trades.Valid {
			rec.NumTrades = null_trades.Int64
		}
		accum_num_trades = accum_num_trades + rec.NumTrades
		accum_volume = accum_volume + rec.Volume
		rec.AccumNumTrades = accum_num_trades
		rec.AccumVolume = accum_volume
		output = append(output,rec)
	}
	return output
}
