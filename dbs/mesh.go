// 0x Mesh events
//	A layer separating Augur trading from 0x Mesh orders
//		Augur trading events are built in 'oohist' table
package dbs

import (
	"fmt"
	"os"
	"encoding/hex"
	"database/sql"
	_  "github.com/lib/pq"

//	"github.com/0xProject/0x-mesh/zeroex"
	ztypes "github.com/0xProject/0x-mesh/common/types"

	p "github.com/PredictionExplorer/augur-explorer/primitives"
)
func (ss *SQLStorage) Insert_0x_mesh_order_event(timestamp int64,order_info *ztypes.OrderInfo,ospec *ZxMeshOrderSpec,event_code p.MeshEvtCode) {

	var query string

	query = "SELECT id FROM mesh_evt WHERE order_hash=$1 AND evt_code=$2"

	var null_id sql.NullInt64
	err:=ss.db.QueryRow(query,order_info.OrderHash.String(),event_code).Scan(&null_id);
	if (err!=nil) {
		if (err==sql.ErrNoRows) {
			// break
		} else {
			ss.Log_msg(fmt.Sprintf("DB error : %v, q=%v",err,query))
			os.Exit(1)
		}
	} else {
		// market already exists
		return
	}

	query = "INSERT INTO mesh_evt (" +
				"time_stamp," +
				"fillable_amount," +
				"evt_code," +
				"market_aid," +
				"outcome_idx," +
				"otype," +
				"price," +
				"order_hash," +
				"chain_id," +
				"exchange_addr," +
				"maker_addr," +
				"maker_asset_data," +
				"maker_fee_asset_data," +
				"maker_asset_amount," +
				"maker_fee," +
				"taker_address," +
				"taker_asset_data," +
				"taker_asset_amount," +
				"taker_fee," +
				"taker_fee_asset_data," +
				"sender_address," +
				"fee_recipient_address," +
				"expiration_time," +
				"salt," +
				"signature" +
			") VALUES (" +
					"TO_TIMESTAMP($1)," +
					"($2::decimal/1e+18),$3,"+
					"$4,$5,$6,($7/1e+18)," +
					"$8,$9,$10,$12,$13,$14," +
					"($15::decimal/1e+18),($16::decimal/1e+18),"+
					"$17,$18,"+
					"($19::decimal/1e+18),($20::decimal/1e+18),"+
					"$21,$22,$23,TO_TIMESTAMP($24),$25,$26"+
			") ON CONFLICT DO NOTHING"

	_,err = ss.db.Exec(query,
		timestamp,
		order_info.FillableTakerAssetAmount.String(),
		event_code,
		ospec.Market.String(),
		ospec.Outcome,
		ospec.Type,
		ospec.Price.String(),
		order_info.OrderHash.String(),
		order_info.SignedOrder.Order.ChainID.Int64(),
		order_info.SignedOrder.Order.ExchangeAddress.String(),
		order_info.SignedOrder.Order.MakerAddress.String(),
		hex.EncodeToString(order_info.SignedOrder.Order.MakerAssetData),
		hex.EncodeToString(order_info.SignedOrder.Order.MakerFeeAssetData),
		order_info.SignedOrder.Order.MakerAssetAmount.String(),
		order_info.SignedOrder.Order.MakerFee.String(),
		order_info.SignedOrder.Order.TakerAddress.String(),
		hex.EncodeToString(order_info.SignedOrder.Order.TakerAssetData),
		order_info.SignedOrder.Order.TakerAssetAmount.String(),
		order_info.SignedOrder.Order.TakerFee.String(),
		hex.EncodeToString(order_info.SignedOrder.Order.TakerFeeAssetData),
		order_info.SignedOrder.Order.SenderAddress.String(),
		order_info.SignedOrder.Order.FeeRecipientAddress.String(),
		order_info.SignedOrder.Order.ExpirationTimeSeconds.Int64(),
		order_info.SignedOrder.Order.Salt.String(),
		hex.EncodeToString(order_info.SignedOrder.Signature),
	)
	if (err!=nil) {
		ss.Log_msg(fmt.Sprintf("Insert_0x_mesh_order_event() failed: %v, q=%v",err,query))
		os.Exit(1)
	}
}
func (ss *SQLStorage) Get_mesh_proc_status() p.MeshProcStatus {

	var status p.MeshProcStatus
	var query string
	var null_id sql.NullInt64

	for {
		query = "SELECT last_id_processed FROM mesh_status"

		res := ss.db.QueryRow(query)
		err := res.Scan(&null_id)
		if (err!=nil) {
			if err == sql.ErrNoRows {
				query = "INSERT INTO mesh_status DEFAULT VALUES"
				_,err := ss.db.Exec(query)
				if (err!=nil) {
					ss.Log_msg(fmt.Sprintf("DB error: %v q=%v",err,query))
					os.Exit(1)
				}
			} else {
				if (err!=nil) {
					ss.Log_msg(fmt.Sprintf("DB error: %v q=%v",err,query))
					os.Exit(1)
				}
			}
		} else {
			break
		}
	}
	if null_id.Valid {
		status.LastIdProcessed = null_id.Int64
	}
	return status
}
func (ss *SQLStorage) Set_mesh_proc_status(last_id int64) {

	var query string = "UPDATE mesh_status SET last_id_processed=$1"
	res,err:=ss.db.Exec(query,last_id)
	if (err!=nil) {
		ss.Log_msg(fmt.Sprintf("Set_mesh_proc_status() failed: %v",err))
		os.Exit(1)
	}
	affected_rows,err:=res.RowsAffected()
	if err!=nil {
		ss.Log_msg(fmt.Sprintf("Error getting RowsAffected in Set_mesh_proc_status(): %v",err))
		os.Exit(1)
	}
	if affected_rows>0 {
		// break
	} else {
		query = "INSERT INTO mesh_status VALUES($1)"
		_,err := ss.db.Exec(query,last_id)
		if (err!=nil) {
			ss.Log_msg(fmt.Sprintf("Set_mesh_proc_status() failed on INSERT: %v",err));
			os.Exit(1)
		}
	}
}
func (ss *SQLStorage) Get_mesh_events_from_id(id int64) []p.MeshEvent {

	records := make([]p.MeshEvent,0,16)

	var query string
	query = "SELECT " +
				"id," +
				"FLOOR(EXTRACT(EPOCH FROM time_stamp))::BIGINT as time_stamp," +
				"ROUND(fillable_amount*1e+18) AS fillable_amount,evt_code,order_hash," +
				"chain_id,exchange_addr,maker_addr,maker_asset_data,maker_fee_asset_data," +
				"ROUND(maker_asset_amount*1e+18) AS maker_asset_amount,"+
				"ROUND(maker_fee*1e+18) AS maker_fee,taker_address," +
				"taker_asset_data,taker_fee_asset_data," +
				"ROUND(taker_asset_amount*1e+18) AS taker_asset_amount," +
				"ROUND(taker_fee*1e+18) AS taker_fee,sender_address,fee_recipient_address," +
				"FLOOR(EXTRACT(EPOCH FROM expiration_time))::BIGINT as expiration_ts," +
				"salt,signature " +
			"FROM mesh_evt " +
			"WHERE id > $1 " +
			"ORDER BY ID"
//	ss.Info.Println("(id=%v) q=%v\n",id,query)
	rows,err := ss.db.Query(query,id)
	if (err!=nil) {
		ss.Log_msg(fmt.Sprintf("DB error: %v (query=%v)",err,query))
		os.Exit(1)
	}
	var rec p.MeshEvent

	defer rows.Close()
	for rows.Next() {
		err=rows.Scan(
			&rec.Id,
			&rec.Timestamp,
			&rec.FillableAmount,
			&rec.EvtCode,
			&rec.OrderHash,
			&rec.ChainId,
			&rec.ExchangeAddress,
			&rec.MakerAddress,
			&rec.MakerAssetData,
			&rec.MakerFeeAssetData,
			&rec.MakerAssetAmount,
			&rec.MakerFee,
			&rec.TakerAddress,
			&rec.TakerAssetData,
			&rec.TakerFeeAssetData,
			&rec.TakerAssetAmount,
			&rec.TakerFee,
			&rec.SenderAddress,
			&rec.FeeRecipientAddress,
			&rec.ExpirationTime,
			&rec.Salt,
			&rec.Signature,
		)
		if err!=nil {
			ss.Log_msg(fmt.Sprintf("DB error: %v, q=%v",err,query))
			os.Exit(1)
		}
		records = append(records,rec)
	}
	return records
}
func (ss *SQLStorage) Get_mesh_event_by_order_hash(order_hash string) (p.MeshEvent,error) {

	var order p.MeshEvent
	order.OrderHash = order_hash
	var query string
	query = "SELECT " +
				"FLOOR(EXTRACT(EPOCH FROM time_stamp) AS time_stamp," +
				"fillable_amount," +
				"evt_code," +
				"chain_id," +
				"exchange_addr," +
				"maker_addr," +
				"maker_asset_data," +
				"maker_fee_asset_data," +
				"maker_asset_amount," +
				"maker_fee," +
				"taker_address," +
				"taker_asset_data," +
				"taker_asset_amount," +
				"taker_fee," +
				"taker_fee_asset_data," +
				"sender_address," +
				"fee_recipient_address," +
				"FLOOR(EXTRACT(EPOCH FROM expiration_time))::BIGINT AS expiration_time," +
				"salt," +
				"signature" +
			"FROM mesh_evt " +
			"WHERE order_hash = $1 " +
			"LIMIT 1"
	row := ss.db.QueryRow(query,order_hash)
	err := row.Scan(
		&order.Timestamp,
		&order.FillableAmount,
		&order.EvtCode,
		&order.ChainId,
		&order.ExchangeAddress,
		&order.MakerAddress,
		&order.MakerAssetData,
		&order.MakerFeeAssetData,
		&order.MakerAssetAmount,
		&order.MakerFee,
		&order.TakerAddress,
		&order.TakerAssetData,
		&order.TakerFeeAssetData,
		&order.TakerAssetAmount,
		&order.TakerFee,
		&order.SenderAddress,
		&order.FeeRecipientAddress,
		&order.ExpirationTime,
		&order.Salt,
		&order.Signature,
	)
	if err!=nil {
		ss.Log_msg(fmt.Sprintf("DB error: %v, q=%v",err,query))
		os.Exit(1)
	}
	return order,nil
}
