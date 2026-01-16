///Verifies block.cash flow value with optional fix in case of error
package main
import (
	"os"
	"log"

	. "github.com/PredictionExplorer/augur-explorer/dbs"
)
var (
	storage *SQLStorage
	market_order_id int64 = 0
	Info    *log.Logger
)
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

	cash_flow := storage.Get_cash_flow()
	var counter_bad int64 = 0
	for _,rec := range cash_flow {
		total_amount := storage.Get_dai_bal_total_amount(rec.BlockNum)
		if total_amount != rec.CashFlow {
			difference := total_amount - rec.CashFlow
			Info.Printf(
				"%v block: cash flow difference of %v (have %v, should have %v)\n",
				rec.BlockNum,difference,rec.CashFlow,total_amount,
			)
			counter_bad++
			if fix_flag {
				storage.Set_cash_flow_value(rec.BlockNum,total_amount)
				Info.Printf("fix: set cash_flow to %v for block %v\n",total_amount,rec.BlockNum)
			}
		}
	}
	Info.Printf("Results: %v bad records out of total %v records\n",counter_bad,len(cash_flow))
}
