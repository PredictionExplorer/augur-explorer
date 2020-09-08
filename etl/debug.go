package main

import (
//	"time"
	"bytes"
	"encoding/hex"
//	"math/big"
//	"context"
//	"os"
//	"errors"
//	"fmt"

//	"github.com/ethereum/go-ethereum/common"
//	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
//	"github.com/0xProject/0x-mesh/zeroex"

	. "github.com/PredictionExplorer/augur-explorer/primitives"
)
func dump_tx_input_if_known(tx_data []byte) {

	if len(tx_data) < 32 {
		Info.Printf("dump_tx_input: input sig: %v\n",hex.EncodeToString(tx_data[:]))
		return
	}
	input_sig := tx_data[:4]
	set_timestamp_sig ,_ := hex.DecodeString("a0a2b573")
	if 0 == bytes.Compare(input_sig,set_timestamp_sig) {
		Info.Printf("augur_wallet_call: Skipping setTimestamp() transaction\n")
		return
	}
	decoded_sig ,_ := hex.DecodeString("78dc0eed")
	if 0 == bytes.Compare(input_sig,decoded_sig) {
		input_data_raw:= tx_data[4:]
		var input_data ExecWalletTxInputStruct
		method, err := wallet_abi.MethodById(decoded_sig)
		if err != nil {
			Fatalf("Method not found")
		}
		err = method.Inputs.Unpack(&input_data, input_data_raw)
		if err != nil {
			Fatalf("Couldn't decode input of tx %v",err)
		}
		Info.Printf("ExecuteWalletTransaction {\n")
		Info.Printf("\tto: %v\n",input_data.To.String())
		Info.Printf("\tdata: %v\n",hex.EncodeToString(input_data.Data[:]))
		Info.Printf("\tvalue: %v\n",input_data.Value.String())
		Info.Printf("\tpayment: %v\n",input_data.Payment.String())
		Info.Printf("\treferralAddress:  %v\n",input_data.ReferralAddress.String())
		Info.Printf("\tfingerprint: %v\n",hex.EncodeToString(input_data.Fingerprint[:]))
		Info.Printf("\tdesiredSignerBalance: %v\n",input_data.DesiredSignerBalance.String())
		Info.Printf("\tmaxExchangeRateInDai: %v\n",input_data.MaxExchangeRateInDai.String())
		Info.Printf("\trevertOnFaliure: %v\n",input_data.RevertOnFailure)
		Info.Printf("}\n")

		// check for internal transactions for the Wallet Registry contract
		if len(input_data.Data) >= 4 {
			input_sig := input_data.Data[:4]
			market_proceeds_sig ,_ := hex.DecodeString("db754422")
			if 0 == bytes.Compare(input_sig,market_proceeds_sig) {
				Info.Printf("augur_wallet_call: claimMarketProceeds()\n")
				return
			}
			trading_proceeds_sig ,_ := hex.DecodeString("efd342c1")
			if 0 == bytes.Compare(input_sig,trading_proceeds_sig) {
				Info.Printf("augur_wallet_call: claimTradingProceeds()\n")
				return
			}
			zeroex_trade_sig ,_ := hex.DecodeString("2f562016")
			if 0 == bytes.Compare(input_sig,zeroex_trade_sig) {
				Info.Printf("augur_wallet_call: ZeroEx::trade()\n")
				amounts := decode_original_fill_amount(input_data.Data[4:],zeroex_trade_sig)
				for h,a := range amounts {
					if a == nil {
						Fatalf("amounts map contains null initial_order bigint")
					}
					Info.Printf("o %v, amount = %v\n",h,a.String())
				}
				return
			}
		}
	} else {
		Info.Printf("dump_tx_input: input sig: %v\n",hex.EncodeToString(input_sig[:]))
		if len(input_sig) >= 4 {
			input_data_raw:= tx_data[4:]
			Info.Printf("tx input= %v\n",hex.EncodeToString(input_data_raw))
			zeroex_trade_sig ,_ := hex.DecodeString("2f562016")
			if 0 == bytes.Compare(input_sig,zeroex_trade_sig) {
				Info.Printf("direct call to ZeroEx::trade()\n")
				amounts := decode_original_fill_amount(input_data_raw,zeroex_trade_sig)
				for h,a := range amounts {
					if a == nil {
						Fatalf("amounts map contains null initial_order bigint")
					}
					Info.Printf("o %v, amount = %v\n",h,a.String())
				}
				return
			}
		}
	}
}
func scan_profit_loss_data_for_debugging(block_num int64,position_changes *[]*PosChg) {
	// this function makes direct calls to contracts to get changes in profit loss and record them
	// right after each block is processed (developed for debugging purposes)

	var copts = new(bind.CallOpts)
	for i:=0 ; i<len(*position_changes) ; i++ {
		pchg := (*position_changes)[i]
		Info.Printf("profit_loss debug: processing pl for %v\n",pchg.Wallet_addr.String())
		pl,err := ctrct_pl.GetRealizedProfit(copts,pchg.Mkt_addr,pchg.Wallet_addr,pchg.Outcome)
		if err != nil {
			Error.Printf("Failure to GetRealizedProfit for addr %v outcome %v: %v\n",
							pchg.Wallet_addr.String(),pchg.Outcome.String(),err)
			continue
		} else {
			pchg.ProfitLoss = pl
		}
		ff,err := ctrct_pl.GetFrozenFunds(copts,pchg.Mkt_addr,pchg.Wallet_addr,pchg.Outcome)
		if err != nil {
			Error.Printf("Failure to GetFrozenFunds for addr %v outcome %v: %v\n",
							pchg.Wallet_addr.String(),pchg.Outcome.String(),err)
			continue
		} else {
			pchg.FrozenFunds = ff
		}
		psize,err := ctrct_pl.GetNetPosition(copts,pchg.Mkt_addr,pchg.Wallet_addr,pchg.Outcome)
		if err != nil {
			Error.Printf("Failure to GetNetPosition for addr %v outcome %v: %v\n",
							pchg.Wallet_addr.String(),pchg.Outcome.String(),err)
			continue
		} else {
			pchg.NetPos = psize
		}
		price,err := ctrct_pl.GetAvgPrice(copts,pchg.Mkt_addr,pchg.Wallet_addr,pchg.Outcome)
		if err != nil {
			Error.Printf("Failure to GetAvgPricefor addr %v outcome %v: %v\n",
							pchg.Wallet_addr.String(),pchg.Outcome.String(),err)
			continue
		} else {
			pchg.AvgPrice = price
		}
		pchg.BlockNum=int64(block_num)
		Info.Printf("inserting pchg %+v\n",*pchg)
		storage.Insert_profit_loss_debug_rec(pchg)
	}
}
