// dumps StateDB objects to verify all accounts have correct balances
package main
import (
	"fmt"
	"os"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/state"
	//"github.com/ethereum/go-ethereum/core/state/snapshot"
	//"github.com/ethereum/go-ethereum/trie"

	 . "github.com/PredictionExplorer/augur-explorer/uevm"
)


func main() {

	if len(os.Args) < 3 {
		fmt.Printf(
			"Usage: \n\t\t%v [datadir] [state_root_hash]\n\t\t"+
			"dumps state accounts provided state root hash\")\n\n",os.Args[0],
		)
		os.Exit(1)
	}
	datadir := os.Args[1]
	state_root := common.HexToHash(os.Args[2])
	fmt.Printf("Trying to get state with hash %v\n",state_root.String())
	_,edb := OpenDB(datadir)
	state_db,err := state.New(state_root,edb,nil)
	if err != nil {
		fmt.Printf("Error creating StateDB object: %v\n",err)
		os.Exit(1)
	}
	raw_dump := GetStateDump(state_db)
	DumpStateDB(raw_dump)
	addr := common.HexToAddress("0x913dA4198E6bE1D5f5E4a40D0667f70C0B5430Eb")
	balance := state_db.GetBalance(addr)
	fmt.Printf("Addr %v has balance %v\n",addr.String(),balance.String())
}
