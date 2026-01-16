package uevm
import (
	"fmt"
)
func DumpRecord(r *Record) {
	fmt.Printf("BlockNum\t%v\n",r.BlockNum)
	fmt.Printf("BlockHash\t%v\n",r.BlockHash.String())
	fmt.Printf("TxIndex\t%v\n",r.TxIndex)
	fmt.Printf("TxHash\t%v\n",r.TxHash.String())
	fmt.Printf("StateRoot\t%v\n",r.StateRoot)
}
