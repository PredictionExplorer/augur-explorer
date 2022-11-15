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
	edb := OpenDB(datadir)
	//triedb := trie.NewDatabase(edb)
	/*
	regenerate_snapshot := true
	sshot,err := snapshot.New(edb,triedb,256,common.Hash{},false,regenerate_snapshot,false)
	if err != nil {
		fmt.Printf("Error creating snapshot object: %v\n",err)
		os.Exit(1)
	}
	state_db,err := state.New(state_root,state.NewDatabase(edb),sshot)
	*/
	state_db,err := state.New(state_root,state.NewDatabase(edb),nil)
	if err != nil {
		fmt.Printf("Error creating StateDB object: %v\n",err)
		os.Exit(1)
	}/*
	dump_config := state.DumpConfig{
		SkipCode:			false,
		SkipStorage:		false,
		OnlyWithAddresses:	false,
		Start:				nil,
		Max:				0,
	}
	raw_dump := state_db.RawDump(&dump_config)
	fmt.Printf("Total objects: %v\n",len(raw_dump.Accounts))
	fmt.Printf("Address\t\t\tBalance\n")
	keys := make([]common.Address,len(raw_dump.Accounts))
	for i:=0;i<len(keys);i++ {
		obj := raw_dump.Accounts[keys[i]]
		fmt.Printf("Address\t%v\n",obj.Address.String(),obj.Balance)
	}
	*/
	raw_dump := GetStateDump(state_db)
	DumpStateDB(raw_dump)
}
