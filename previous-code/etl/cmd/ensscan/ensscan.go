
package main

import (
	"os"
	"os/signal"
	"syscall"
	"context"
	"math/big"
	"fmt"
	"log"
	"bytes"
	"strings"
	"encoding/hex"

//	"golang.org/x/crypto/sha3"
	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/rpc"
	"github.com/ethereum/go-ethereum/accounts/abi"
//	"github.com/wealdtech/go-ens/v3"
//	"github.com/wealdtech/go-ens/v3/contracts/resolver"

	. "github.com/PredictionExplorer/augur-explorer/primitives"
	. "github.com/PredictionExplorer/augur-explorer/dbs"
	. "github.com/PredictionExplorer/augur-explorer/contracts"
)
const (
	DEFAULT_DB_LOG				= "db.log"
	GRACE_PERIOD			int = 90 * 24 * 60 * 60 // 90 days
	ENS_NEWOWNER				= "ce0457fe73731f824cc272376169235128c118b49d344817417c6d108d155e82"

	ENS_NAME_REGISTERED1		= "ca6abbe9d7f11422cb6ca7629fbf6fe9efb1c621f71ce8f02b9f2a230097404f"
	//	NameRegistered (	// this was introduced in 2019
	//		string name,
	//		index_topic_1 bytes32 label,
	//		index_topic_2 address owner,
	//		uint256 cost,
	//		uint256 expires
	//	)

	ENS_NAME_REGISTERED2		= "b3d987963d01b2f68493b4bdb130988f157ea43070d4ad840fee0466ed9370d9"
	// NameRegistered (
	//		index_topic_1 uint256 id,
	//		index_topic_2 address owner,
	//		uint256 expires
	//	)

	ENS_NAME_REGISTERED3		= "570313dae523ecb48b1176a4b60272e5ea7ec637f5b2d09983cbc4bf25e7e9e3"
	//	NameRegistered (
	//		index_topic_1 address _caller,
	//		index_topic_2 address _beneficiary,
	//		index_topic_3 bytes32 _labelHash,
	//		string _subdomain,
	//		uint256 _createdDate
	//	)

	ENS_NAME_MIGRATED			= "ea3d7e1195a15d2ddcd859b01abd4c6b960fa9f9264e499a70a90c7f0c64b717"
	//	NameMigrated(
	//		uint256 indexed hash,
	//		address indexed owner,
	//		uint expires
	//  )

	ENS_ADDR_CHANGED			= "52d7d861f09ab3d26239d492e8968629f95e9e318cf0b73bfddc441522a15fd2"
	ENS_ADDRESS_CHANGED			= "65412581168e88a1e60c6459d7f44ae83ad0832e670826c05a4e2476b57af752"
	NEW_RESOLVER				= "335721b01866dc23fbee8b6b2c7b1e14d6f05c28cd35a2c934239f94095602a0"

	ENS_REGISTRAR_TRANSFER		= "ddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef"
	// transfer at the Registrar level
	//	event Transfer(
	//		address indexed from,
	//		address indexed to,
	//		uint256 indexed tokenId
	//	);

	ENS_REGISTRY_TRANSFER		= "d4735d920b0f87494915f556dd9b54c8f309026070caea5c737245152564d266"
	// transfer at Registry level
	// event Transfer(
	//		bytes32 indexed node,
	//		address owner
	//	);

	HASH_REGISTERED				= "0f0c27adfd84b60b6f456b0e87cdccb1e5fb9603991588d87fa99f5b6b61e670"
	//  HashRegistered (	Sample tx: 0x216d8103a59c3f3921210b3a4c6aca32c21724dac8451d32ed0d88103c20b802
	//		index_topic_1 bytes32 hash,
	//		index_topic_2 address owner,
	//		uint256 value,
	//		uint256 registrationDate
	//	)

	NAME_RENEWED				= "3da24c024582931cfaf8267d8ed24d13a82a8068d5bd337d30ec45cea4e506ae"
	//	event NameRenewed(
	//		uint256 indexed id,
	//		uint expires
	//	);

	OWNER_CHANGED				= "06e9c07310f63759634ddbb7257dbb19ca404f90bd6bdef1d3386fab033cebce"
	// event OwnerChanged(bytes32 indexed label, address indexed oldOwner, address indexed newOwner);
	//  from ENS migration contract

	CONTROLLER_ADDED			= "0a8bb31534c0ed46f380cb867bd5c803a189ced9a764e30b3a4991a9901d7474"
	//	event ControllerAdded(address indexed controller);

	CONTROLLER_REMOVED			= "33d83959be2573f5453b12eb9d43b3499bc57d96bd2f067ba44803c859e81113"
	//	event ControllerRemoved(address indexed controller);

	HASH_INVALIDATED			= "1f9c649fe47e58bb60f4e52f0d90e4c47a526c9f90c5113df842c025970b66ad"
	NEW_TTL						= "1d4f9bbfc9cab89d66e1a1562f2233ccbf1308cb4f63de2ead5787adddb8fa68."
	ENS_TEXT_CHANGED			= "d8c9334b1a9c2f9da342a0a2b32629c1a229b6445dad78947f674b44444a7550"
	NAME_BOUGHT					= "b8c56202a5ae8b00edfcd57a54ec6c3fb8d2f6deb3067a7ba11408a7bd014a3e"
	NAME_CHANGED				= "b7d29e911041e8d9b843369e890bcb72c9388692ba48b65ac54e7214c4c348f7"

	PUBKEY_CHANGED				= "1d6f5e03d3f63eb58751986629a5439baee5079ff04f345becb66e23eb154e46"
	CONTENT_HASH_CHANGED		= "e379c1624ed7e714cc0937528a32359d69d5281337765313dba4e081b72d7578"

	ENS_V1_REGISTRY_ADDR		= "0x00000000000C2E074eC69A0dFb2997BA6C7d2e1e"	// 10 Mar 2017
	ENS_V2_REGISTRY_ADDR		= "0x314159265dD8dbb310642f98f50C066173C1259b"	// 30 Jan 2020

	//[common.Address].addr.reverse
	ADDRIN_REVERSE_NODE			= "91d1777781884d03a6757a803996e38de2a42967fb37eeaca72729271025a9e2"
	REVERSE_REG_V2_ADDR			= "0x084b1c3C81545d370f3634392De611CaaBFf8148"
	REVERSE_REG_V1_ADDR			= "0x9062C0A6Dbd6108336BcBe4593a3D1cE05512069"
)
var (
	evt_newowner,_ = hex.DecodeString(ENS_NEWOWNER)
	evt_name_registered1,_ = hex.DecodeString(ENS_NAME_REGISTERED1)
	evt_name_registered2,_ = hex.DecodeString(ENS_NAME_REGISTERED2)
	evt_name_registered3,_ = hex.DecodeString(ENS_NAME_REGISTERED3)
	evt_name_renewed,_ = hex.DecodeString(NAME_RENEWED)
	evt_name_migrated,_ = hex.DecodeString(ENS_NAME_MIGRATED)
	evt_hash_invalidated,_ = hex.DecodeString(HASH_INVALIDATED)
	evt_hash_registered,_ = hex.DecodeString(HASH_REGISTERED)
	evt_new_resolver,_ = hex.DecodeString(NEW_RESOLVER)
	evt_registry_transfer,_ = hex.DecodeString(ENS_REGISTRY_TRANSFER)
	evt_registrar_transfer,_ = hex.DecodeString(ENS_REGISTRAR_TRANSFER)
	evt_text_changed,_		= hex.DecodeString(ENS_TEXT_CHANGED)
	evt_name_bought,_ = hex.DecodeString(NAME_BOUGHT)
	evt_addrchanged1,_ = hex.DecodeString(ENS_ADDR_CHANGED)
	evt_addresschanged2,_ = hex.DecodeString(ENS_ADDRESS_CHANGED)
	evt_pubkey_changed,_ = hex.DecodeString(PUBKEY_CHANGED)
	evt_contenthash_changed,_ = hex.DecodeString(CONTENT_HASH_CHANGED)
	evt_name_changed,_ = hex.DecodeString(NAME_CHANGED)
	evt_owner_changed,_ = hex.DecodeString(OWNER_CHANGED)
	evt_controller_added,_ = hex.DecodeString(CONTROLLER_ADDED)
	evt_controller_removed,_ = hex.DecodeString(CONTROLLER_REMOVED)

	storage *SQLStorage
	RPC_URL = os.Getenv("AUGUR_ETH_NODE_RPC_URL")
	eclient *ethclient.Client
	rpcclient *rpc.Client
	Error   *log.Logger
	Info	*log.Logger

	ens_abi abi.ABI

	ens1_addr			= common.HexToAddress(ENS_V1_REGISTRY_ADDR)
	ens2_addr			= common.HexToAddress(ENS_V2_REGISTRY_ADDR)

	blacklisted				map[common.Address]struct{}	// contracts we need to ignore
)
func do_initial_load_new_owner(block_num_from,block_num_to int64) {

	filter := ethereum.FilterQuery{}
	filter.FromBlock = big.NewInt(block_num_from)
	filter.ToBlock = big.NewInt(block_num_to)
	topics := make([]common.Hash,0,1)
	signature := common.BytesToHash(evt_newowner)
	topics = append(topics,signature)
	filter.Topics= append(filter.Topics,topics)
	filter.Addresses = nil
	Info.Printf("Submitting filter logs query with signature %v\n",hex.EncodeToString(signature.Bytes()))
	Info.Printf("filter query = %+v\n",filter)
	Info.Printf("NewOwner: block range: %v - %v\n",block_num_from,block_num_to)
	logs,err := eclient.FilterLogs(context.Background(),filter)
	if err!= nil {
		Error.Printf("Error: %v\n",err)
		Info.Printf("Error: %v\n",err)
		os.Exit(1)
	}
	for _,log := range logs {
		if log.Removed {
			continue
		}
		if	ens1,ens2 :=
				bytes.Equal(ens1_addr.Bytes(),log.Address.Bytes()),
				bytes.Equal(ens2_addr.Bytes(),log.Address.Bytes());
			!(ens1 || ens2) {
				continue
		}
		proc_newowner(&log,0,0,0)
	}
}
func range_initial_load_new_owner(exit_chan chan bool,block_num_limit int64) {

	var block_num int64 = 2933000 // found empirically
	for ; block_num <= block_num_limit ; {
		select {
			case exit_flag := <-exit_chan:
				if exit_flag {
					Info.Println("Exiting by user request.\n")
					os.Exit(0)
				}
			default:
		}

		next_block_num := block_num + 1000 - 1
		if next_block_num > block_num_limit {
			do_initial_load_new_owner(block_num,block_num_limit)
			break
		} else {
			do_initial_load_new_owner(block_num,next_block_num)
			block_num = next_block_num + 1
		}
	}
}
func do_std_initial_load(block_num_from,block_num_to int64,f std_proc_func,evtname string,sig []byte) {

	filter := ethereum.FilterQuery{}
	filter.FromBlock = big.NewInt(block_num_from)
	filter.ToBlock = big.NewInt(block_num_to)
	topics := make([]common.Hash,0,1)
	signature := common.BytesToHash(sig)
	topics = append(topics,signature)
	filter.Topics= append(filter.Topics,topics)
	filter.Addresses = nil
	Info.Printf("Submitting filter logs query with signature %v\n",hex.EncodeToString(signature.Bytes()))
	Info.Printf("filter query = %+v\n",filter)
	Info.Printf("%v: block range: %v - %v\n",evtname,block_num_from,block_num_to)

	// FilterLogs sometimes returns EOF error (possibly too much resources used and connection is aborted)
	// so we do it in a loop for 10 times, to avoid our load process to fail due to random resource problem
	var logs []types.Log
	num_tries := 10
	for {
		var err error
		logs,err = eclient.FilterLogs(context.Background(),filter)
		if err!= nil {
			Error.Printf("Error: %v\n",err)
			Info.Printf("Error: %v\n",err)
		} else {
			break;
		}
		num_tries = num_tries - 1
		if num_tries == 0 {
			Error.Printf("Aborting process. Error: %v\n",err)
			Info.Printf("Aborting processs. Error: %v\n",err)
			os.Exit(1)
		}
	}
	for _,log := range logs {
		if log.Removed {
			continue
		}
		f(&log,0,0,0)
	}
}
func std_initial_load(exit_chan chan bool,block_num_limit int64,f std_proc_func,evtname string,sig []byte) {

	var block_num int64 = 0
	var interval int64 = 1000
	if evtname == "RegistrarTransfer" {
		block_num = 7691000 // otherwise it takes very long (first block found empirically)
		interval = 250	// load is to high on 1000, we decrease by 50%
	}
	for ; block_num <= block_num_limit ; {
		select {
			case exit_flag := <-exit_chan:
				if exit_flag {
					Info.Println("Exiting by user request.\n")
					os.Exit(0)
				}
			default:
		}

		next_block_num := block_num + interval - 1
		if next_block_num > block_num_limit {
			do_std_initial_load(block_num,block_num_limit,f,evtname,sig)
			break
		} else {
			do_std_initial_load(block_num,next_block_num,f,evtname,sig)
			block_num = next_block_num + 1
		}
	}
}
func initial_load(exit_chan chan bool,bnum_lim int64) {
/*
	std_initial_load(exit_chan,bnum_lim,proc_name_registered1,"NameRegistered1",evt_name_registered1)
	std_initial_load(exit_chan,bnum_lim,proc_name_registered2,"NameRegistered2",evt_name_registered2)
	std_initial_load(exit_chan,bnum_lim,proc_name_registered3,"NameRegistered3",evt_name_registered3)
	range_initial_load_new_owner(exit_chan,bnum_lim)
	std_initial_load(exit_chan,bnum_lim,proc_addr_changed1,"AddrChanged_1",evt_addrchanged1)
	std_initial_load(exit_chan,bnum_lim,proc_address_changed2,"AddressChanged_2",evt_addresschanged2)
	std_initial_load(exit_chan,bnum_lim,proc_name_changed,"NameChanged",evt_name_changed)
	std_initial_load(exit_chan,bnum_lim,proc_hash_invalidated,"HashInvalidated",evt_hash_invalidated)
	std_initial_load(exit_chan,bnum_lim,proc_new_resolver,"NewResolver",evt_new_resolver)
	std_initial_load(exit_chan,bnum_lim,proc_name_migrated,"NameMigrated",evt_name_migrated)
	std_initial_load(exit_chan,bnum_lim,proc_registry_transfer,"RegistryTransfer",evt_registry_transfer)
	std_initial_load(exit_chan,bnum_lim,proc_name_renewed,"NameRenewed",evt_name_renewed)
	std_initial_load(exit_chan,bnum_lim,proc_text_changed,"TextChanged",evt_text_changed)
	std_initial_load(exit_chan,bnum_lim,proc_hash_registered,"HashRegistered",evt_hash_registered)
	std_initial_load(exit_chan,bnum_lim,proc_pubkey_changed,"PubkeyChanged",evt_pubkey_changed[:])
	std_initial_load(exit_chan,bnum_lim,proc_contenthash_changed,"ContenthashChanged",evt_contenthash_changed[:])
	std_initial_load(exit_chan,bnum_lim,proc_owner_changed,"OwnerChanged",evt_owner_changed[:])
	std_initial_load(exit_chan,bnum_lim,proc_registrar_transfer,"RegistrarTransfer",evt_registrar_transfer)
	*/
}
func check_initial_load_completness() bool {

	const correct_num_active_names int64 = 14040; // empirically found
	num_active_names := storage.Get_count_of_active_names()
	if num_active_names != correct_num_active_names {
		msg := fmt.Sprintf(
			"Number of active names is %v but should be %v\n",
			num_active_names,
			correct_num_active_names,
		)
		Error.Printf(msg)
		Info.Printf(msg)
		return false
	}
	return true
}
func main() {

	var do_initial_load bool = false
	if len(os.Args) > 1 {
		if os.Args[1] != "-i" {
			fmt.Printf("Unknown option %v\n",os.Args[1])
			fmt.Printf("Usage: %v [-i]\n\t -i = initial load\n",os.Args[0])
			os.Exit(1)
		}
		do_initial_load = true
	}
	log_dir:=fmt.Sprintf("%v/%v",os.Getenv("HOME"),DEFAULT_LOG_DIR)
	os.MkdirAll(log_dir, os.ModePerm)
	db_log_file:=fmt.Sprintf("%v/ensscan_%v",log_dir,DEFAULT_DB_LOG)

	fname:=fmt.Sprintf("%v/ensscan_info.log",log_dir)
	logfile, err := os.OpenFile(fname, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err!=nil {
		fmt.Printf("Can't start: %v\n",err)
		os.Exit(1)
	}
	Info = log.New(logfile,"INFO: ",log.Ltime|log.Lshortfile)

	fname=fmt.Sprintf("%v/ensscan_error.log",log_dir)
	if err!=nil {
		fmt.Printf("Can't start: %v\n",err)
		os.Exit(1)
	}
	logfile, err = os.OpenFile(fname, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	Error = log.New(logfile,"ERROR: ",log.Ltime|log.Lshortfile)

	rpcclient, err := rpc.DialContext(context.Background(), RPC_URL)
	if err != nil {
		log.Fatal(err)
	}
	Info.Printf("Connected to ETH node: %v\n",RPC_URL)
	eclient = ethclient.NewClient(rpcclient)
	storage = Connect_to_storage(Info)
	storage.Init_log(db_log_file)
	storage.Log_msg("Log initialized\n")

	c := make(chan os.Signal)
	exit_chan := make(chan bool)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-c
		Info.Printf("Got SIGINT signal, will exit after processing is over." +
					" To interrupt abruptly send SIGKILL (9) to the kernel.\n")
		exit_chan <- true
	}()
	abi_parsed := strings.NewReader(ENSABI)
	ens_abi,err = abi.JSON(abi_parsed)
	if err!= nil {
		Info.Printf("Can't parse ENS ABI: %v\n",err)
		os.Exit(1)
	}
	blacklisted = make(map[common.Address]struct{})
	status := storage.Get_ens_processing_status()
	if do_initial_load {
		ens_status := storage.Get_ens_proc_status()
		if ens_status.LastEvtId != 0 {
			fmt.Printf("ENS status record shows prior usage: last_evt_id=%v\n",ens_status.LastEvtId)
			fmt.Printf("Will refuse to run initialization process\n")
			os.Exit(1)
		}
		initial_load(exit_chan,status.IniLoadBlockNumLimit)
		fmt.Printf("Initial load finished.")
		return
	} else {
		init_ens_processing()
		process_ens_events(exit_chan)
	}

}
