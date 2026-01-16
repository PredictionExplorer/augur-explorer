package contracts

import (
	"log"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	//"github.com/0xProject/0x-mesh/common/types"
	p "github.com/PredictionExplorer/augur-explorer/primitives"

)

func DumpETransferBatch(evt *p.ETransferBatch,zc *ZeroX,l *log.Logger) {
	l.Printf("TransferBatch {\n")
	l.Printf("\tOperator: %v\n",evt.Operator.String())
	l.Printf("\tFrom: %v\n",evt.From.String())
	l.Printf("\tTo: %v\n",evt.To.String())
	ids := p.Bigint_ptr_slice_to_str(&evt.Ids,",")
	l.Printf("\tIds: %v\n",ids)
	l.Printf("\tDecoded token IDs:\n")
	var copts = new(bind.CallOpts)
	copts.Pending = true
	for i:=0 ; i<len(evt.Ids); i++ {
		if false {
			l.Printf("\t\tcan't decode token info hex string: \n")
		} else {
			tok_info,err := zc.UnpackTokenId(copts,evt.Ids[i])
			if err == nil {
				l.Printf("\t\tMarket: %v\n",tok_info.Market.String())
				l.Printf("\t\tPrice: %v\n",tok_info.Price)
				l.Printf("\t\tOutcome: %v\n",tok_info.Outcome)
				l.Printf("\t\tType: %v\n",tok_info.Type)
			} else {
				l.Printf("\t\ttoken decode error: %v\n",err)
			}
		}
	}
	values := p.Bigint_ptr_slice_to_str(&evt.Values,",")
	l.Printf("\tValues: %v\n",values)
	l.Printf("}\n")
}
