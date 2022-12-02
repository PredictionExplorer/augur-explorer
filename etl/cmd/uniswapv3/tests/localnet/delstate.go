// deletes StateDB object from Trie Database
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
			"deletes statedb object from TrieDB\n\n",os.Args[0],
		)
		os.Exit(1)
	}
	datadir := os.Args[1]
	state_root := common.HexToHash(os.Args[2])
	fmt.Printf("Trying to get state with hash %v\n",state_root.String())
	edb := OpenDB(datadir)
	_,err := state.New(state_root,edb,nil)
	if err != nil {
		fmt.Printf("Error. state doesn't exist: %v\n",err)
		os.Exit(1)
	}
	trie,err:=edb.OpenTrie(state_root)
	if err != nil {
		fmt.Printf("Error opening trie: %v\n",err)
		os.Exit(1)
	}
	err = trie.TryDelete(state_root.Bytes())
	if err == nil {
		fmt.Printf("object deleted\n")
	} else {
		fmt.Printf("Error: %v\n",err)
	}
	/*
	err = edb.TrieDB().Commit(state_root, true, nil)
	if err != nil {
		fmt.Printf("Error at TrieDB()).Commit(): %v\n",err)
		os.Exit(1)
	}
	*/
	newhash,nodeset,err := trie.Commit(true)
	if err != nil {
		fmt.Printf("Error at Commit(): %v\n",err)
		os.Exit(1)
	}
	fmt.Printf("Commit: new hash %v\n",newhash.String())
	fmt.Printf("nodeset: %+v\n",nodeset)
}
