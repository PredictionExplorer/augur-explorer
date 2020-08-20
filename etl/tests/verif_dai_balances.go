///Verifies block.cash flow value with optional fix in case of error
package main
import (
	"os"
	"log"
	"math/big"

	. "github.com/PredictionExplorer/augur-explorer/dbs"
)
var (
	storage *SQLStorage
	market_order_id int64 = 0
	Info    *log.Logger
)
func verify_dai_balances_for_address(aid int64,fix_flag bool) bool {

	dai_balances := storage.Get_dai_balances_by_aid(aid)
	prev_bal := big.NewInt(0)
	for i:=0 ; i<len(dai_balances) ; i++ {
		daib := &dai_balances[i]
		if i==0 {
			prev_bal.SetString(daib.Balance,10)
			continue
		}
		if daib.Processed == false {
			return true	// reached unprocessed balance, assume success
		}
		correct_balance := big.NewInt(0)
		amount := big.NewInt(0)
		amount.SetString(daib.Amount,10)
		correct_balance.Add(prev_bal,amount)
		dai_balance:=big.NewInt(0)
		dai_balance.SetString(daib.Balance,10)
		cmp_res := correct_balance.Cmp(dai_balance)
		prev_bal.SetString(daib.Balance,10)
		if cmp_res != 0 {
			difference := big.NewInt(0)
			difference.Sub(correct_balance,dai_balance)
			bnum,err := storage.Get_block_num_by_hash(daib.BlockHash)
			if (err!=nil) || (bnum!=daib.BlockNum) {
				// Chain have changed, the verification process is no longer valid, assume success
				Info.Printf("Chain have changed during verification of aid=%v. Aborting.\n",aid)
				return true
			}
			prev_bnum,err := storage.Get_block_num_by_hash(dai_balances[i-1].BlockHash)
			if (err!=nil) || (prev_bnum!=dai_balances[i-1].BlockNum) {
				// Chain have changed, the verification process is no longer valid, assume success
				Info.Printf("Chain have changed during verification of aid=%v. Aborting.\n",aid)
				return true
			}
			Info.Printf(
				"Incorrect balance on dai_bal.id=%v (block %v, aid=%v)  have %v should have %v (%v dif)\n",
				daib.Id,daib.BlockNum,aid,daib.Balance,correct_balance.String(),difference.String(),
			)
			if fix_flag {
			}
			return false
		}
	}
	if len(dai_balances) > 1 {
		Info.Printf("aid %v with %v records is fine\n",aid,len(dai_balances))
	}
	return true
}
func main() {
	Info = log.New(os.Stdout,"INFO: ",log.Ltime)        //|log.Lshortfile)

	var fix_flag bool = false
	if len(os.Args) == 2 {
		if os.Args[1] == "-f" {
			fix_flag = true
		}
	}
	Info.Printf("Fix bad records flag is set to : %v\n",fix_flag)

	storage = Connect_to_storage(&market_order_id,Info)

	addresses := storage.Get_all_address_ids()
	var counter_bad int64 = 0
	for _,aid := range addresses {
		success := verify_dai_balances_for_address(aid,fix_flag)
		if !success {
			counter_bad++
		}
	}
	Info.Printf("Results: %v bad records out of total %v records\n",counter_bad,len(addresses))
}
