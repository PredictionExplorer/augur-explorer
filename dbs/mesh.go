package dbs

import (
	"fmt"
	"os"
	"encoding/hex"
	_  "github.com/lib/pq"

//	"github.com/0xProject/0x-mesh/zeroex"
	"github.com/0xProject/0x-mesh/common/types"

	p "github.com/PredictionExplorer/augur-explorer/primitives"
)
func (ss *SQLStorage) Insert_0x_mesh_order_event(timestamp int64,order_info *types.OrderInfo,event_code p.MeshEvtCode) {

	var query string
	query = "INSERT INTO mesh_evt (" +
				"time_stamp" +
				"fillable_amount," +
				"evt_code," +
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
				"sender_address," +
				"fee_recipient_address," +
				"salt," +
				"signature" +
				") VALUES (" +
					"$1," +
					"($2::decimal/1e+18),"+
					"$3,$4,$5,$6,$7,$8,$9" +
					"($10::decimal/1e+18),($11::decimal/1e+18),"+
					"$12,$13,"+
					"($14::decimal/1e+18),($15::decimal/1e+18),"+
					"$16,$17,$18,$19"+
				")"

	_,err:=ss.db.Exec(query,
		timestamp,
		order_info.FillableTakerAssetAmount.String(),
		event_code,
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
		hex.EncodeToString(order_info.SignedOrder.Order.TakerFeeAssetData),
		order_info.SignedOrder.Order.TakerAssetAmount.String(),
		order_info.SignedOrder.Order.SenderAddress.String(),
		order_info.SignedOrder.Order.FeeRecipientAddress.String(),
		order_info.SignedOrder.Order.ExpirationTimeSeconds.Int64(),
		order_info.SignedOrder.Order.Salt.String(),
		hex.EncodeToString(order_info.SignedOrder.Signature),
	)
	if (err!=nil) {
		ss.Log_msg(fmt.Sprintf("update_top_profit_rank() failed: %v, q=%v",err,query))
		os.Exit(1)
	}
}
func (ss *SQLStorage) Get_mesh_proc_status() MeshProcStatus {

	var status p.MeshProcStatus
	var query string
	var null_id sql.NullInt64

	for {
		query = "SELECT last_id_processed FROM mesh_status"

		res := ss.db.QueryRow(query)
		err := res.Scan(&null_id)
		if (err!=nil) {
			if err == sql.ErrNoRows {
				query = "INSERT INTO last_id_processed DEFAULT VALUES"
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
	if null_last_block.Valid {
		output.LastBlock = null_last_block.Int64
	}
	return output
}
