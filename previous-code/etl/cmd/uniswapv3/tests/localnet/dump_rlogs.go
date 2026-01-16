// dumps receipt logs from local state database using tx hash
package main
import (
	"fmt"
	"os"
	"encoding/hex"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/rlp"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethdb/leveldb"

	//. "github.com/PredictionExplorer/augur-explorer/uevm"
)


func main() {

	if len(os.Args) < 3 {
		fmt.Printf(
			"Usage: \n\t\t%v [datadir] [tx_hash]\n\t\t"+
			"dumps receipt logs provided transaction hash\")\n\n",os.Args[0],
		)
		os.Exit(1)
	}
	datadir := os.Args[1]
	tx_hash := common.HexToHash(os.Args[2])
	db,err := leveldb.New(datadir,0 ,0 ,"receipts",true)
	if err != nil {
		fmt.Printf("Error on New() database: %v\n",err)
		os.Exit(1)
	}
	logs_encoded,err := db.Get(tx_hash.Bytes())
	if err != nil {
		fmt.Printf("Error on Get(): %v\n",err)
		os.Exit(1)
	}
	var logs []*types.Log
	err = rlp.DecodeBytes(logs_encoded,&logs)
	if err != nil {
		fmt.Printf("Error in RLP decode: %v\n",err)
		os.Exit(1)
	}
	if len(logs) == 0 {
		fmt.Printf("Transaction exists, but it doesn't have any logs\n")
		os.Exit(1)
	}
	fmt.Printf("%v entries:\n",len(logs))
	fmt.Printf("Topics\t\t\t\tData\n")
	for i:=0;i<len(logs);i++ {
		log := logs[i]
		if len(log.Topics) > 0 {
			fmt.Printf("0[%v]\t",log.Topics[0])
		}
		if len(log.Topics) > 1 {
			fmt.Printf("1[%v]\t",log.Topics[1])
		}
		if len(log.Topics) >2 {
			fmt.Printf("2[%v]\t",log.Topics[2])
		}
		fmt.Printf("%v\n",hex.EncodeToString(log.Data))
	}

}
