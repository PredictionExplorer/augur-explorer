// Verifies the owner of each token against the DB by querying directly to RPC node
package main
import (
	"os"
	"log"
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"

	. "github.com/PredictionExplorer/augur-explorer/rwcg/contracts/randomwalk"
	"github.com/PredictionExplorer/augur-explorer/rwcg/dbs"
	rwdb "github.com/PredictionExplorer/augur-explorer/rwcg/dbs/randomwalk"
	rwp "github.com/PredictionExplorer/augur-explorer/rwcg/primitives/randomwalk"
)
var (
	storagew *rwdb.SQLStorageWrapper

	Info    *log.Logger
	RPC_URL string
	caddrs  *rwp.ContractAddresses
)
func main() {
	RPC_URL = os.Getenv("RPC_URL")

	Info = log.New(os.Stdout, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile)
	base := dbs.Connect_to_storage(Info)
	storagew = &rwdb.SQLStorageWrapper{S: base}

	caddrs_obj := storagew.Get_randomwalk_contract_addresses()
	caddrs = &caddrs_obj

	eclient, err := ethclient.Dial(RPC_URL)
	if err != nil {
		fmt.Printf("Can't connect to ETH RPC: %v\n", err)
		os.Exit(1)
	}
	rwalk_addr := common.HexToAddress(caddrs.RandomWalk)
	rwalk_ctrct, err := NewRWalk(rwalk_addr, eclient)
	if err != nil {
		fmt.Printf("Failed to instantiate RWalk contract: %v\n", err)
		os.Exit(1)
	}
	var copts bind.CallOpts
	num_toks_big, err := rwalk_ctrct.NextTokenId(&copts)
	if err != nil {
		fmt.Printf("Error getting num tokens: %v\n", err)
		os.Exit(1)
	}
	num_toks := num_toks_big.Int64()

	rwalk_aid, err := storagew.S.Nonfatal_lookup_address_id(rwalk_addr.String())

	fmt.Printf("num tokens: %v\n",num_toks)

	stats := storagew.Get_random_walk_stats(rwalk_aid)
	if stats.TokensMinted != num_toks {
		fmt.Printf(
			"Error: num tokens doesn't match: real num tokens = %v, db num tokens = %v\n",
			num_toks,stats.TokensMinted,
		)
	} else {
		fmt.Printf("Num tokens in database is set correctly (%v tokens)\n",num_toks)
	}
	fmt.Printf("Starting verification process, will loop %v times\n",num_toks)
	for i:=int64(0); i< num_toks; i++ {
		chain_owner_addr,err := rwalk_ctrct.OwnerOf(&copts,big.NewInt(i))
		if err != nil {
			fmt.Printf("Error during Owner() call: %v\n",err)
			os.Exit(1)
		}

		chain_owner_aid, err := storagew.S.Nonfatal_lookup_address_id(chain_owner_addr.String())
		if err != nil {
			fmt.Printf("Error during addr lookup: %v\n",err)
			os.Exit(1)
		}
		tok_info, err := storagew.Get_rwalk_token_info(rwalk_aid, i)
		if err != nil {
			fmt.Printf("Error getting token info from db: %v",err)
			os.Exit(1)
		}
		if tok_info.CurOwnerAid != chain_owner_aid {
			fmt.Printf(
				"DB invalid: token_id=%v; owner mismatch, real owner %v, owner in db %v\n",
				i,chain_owner_addr.String(),tok_info.CurOwnerAddr,
			)
		}
	}
}
