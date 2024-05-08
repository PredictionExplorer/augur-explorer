package main

import (
	"os"
	"log"
	"fmt"
	"sort"
	"time"
	"context"
	"strings"
	"encoding/hex"
	"os/signal"
	"syscall"

	"github.com/ethereum/go-ethereum/rpc"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"

	. "github.com/PredictionExplorer/augur-explorer/contracts"
	. "github.com/PredictionExplorer/augur-explorer/primitives"
	. "github.com/PredictionExplorer/augur-explorer/primitives/cosmicgame"
	. "github.com/PredictionExplorer/augur-explorer/dbs"
	. "github.com/PredictionExplorer/augur-explorer/dbs/cosmicgame"
)
const (
	DEFAULT_DB_LOG			= "db.log"
	IMGGEN_PATH				= "v2/etl/cmd/cosmicgame/imggen_monitor/imggen_exec" // relative to $HOME

	PRIZE_CLAIM_EVENT		= "27bc828c399c2947fea27bca8a75ced2e94ff2651d607271f051e39db52286ce"
	//BID_EVENT				= "c7beef8f8d25068377c10e7b7f30ef5622f443eb674e07835df1daf5fe84acbf"
	BID_EVENT				= "3ebe28e9be13fedb71392c114461386e763acb563218b28db3690553055cd5a5"
	DONATION_EVENT			= "8b7fe5be5699654fd637d2250cb0d47e88205730710745e78e9d8bcaf8aad8f1"
	DONATION_RECEIVED_EVENT	= "46ff3d75d4645bdbbae4cd6109ba42e6e1b80ea25e69d10472b357b733300368"
	DONATION_SENT_EVENT		= "44d398d152fa0735a428b13ebc78f79fe4cb1b4722292cd233e278552fa36d32"
	CHARITY_UPDATED			= "a0bd6b2fdbf082ae2356710c23fc8d76d56d418cecb4514d119c77a8617b4ffe"
	TOKEN_NAME_EVENT		= "8ad5e159ff95649c8a9f323ac5a457e741897cf44ce07dfce0e98b84ef9d5f12"
	MINT_EVENT				= "c646da88dc2b2526461a0ebb4326e2418ec0bea89496b632b7c9ee42fbfe1d4d"
	NFT_DONATION_EVENT		= "c85be1734ed6a0f4d6adf924d4d41406e2729878c652110a96e2fdec64e118d1"
	DONATED_NFT_CLAIMED		= "0d0e4b90a96d931964d5b0421a6b5b48dd73abf538cdb9ae015069d49e9a4679"
	RAFFLE_DEPOSIT_EVENT	= "cf6f6dcf9306290f700abd5d57c300a4ad7bf0d4086d5a5e88040fcecb0fef90"
	RAFFLE_WITHDRAWAL_EVENT = "49775450da95afb818c4651b894f124d05629f06572a18900bf29db74b04a0d8"
	RAFFLE_NFT_WINNER		= "0afb73a48ffcc4cf27ecb1889154b910544b89c39c5e5822c431062629134e56"
	RAFFLE_NFT_CLAIMED		= "2278d2acbf6ac7ebd6ad5d3171672894b0c220903ad9ad7bb45057d368c98040"
	TRANSFER_EVT			= "ddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef"
	STAKE_ACTION_EVENT		= "fc5cf5c5f6cbe80581e46dc67e0d882f36b5b4dd66ac607f7a95985d41aa83f4"
	UNSTAKE_ACTION_EVENT	= "33940a9fc3ee3c9150b29b34ef29ca397b5e1e340425a4f0da0346b5b90766c8"
	ETH_DEPOSIT_EVENT		= "dc0eacba8b1f88284dca5eec8be23173aefa7206298fe22de43e064b6ccd8418"
	CLAIM_REWARD_EVENT		= "dde81df5caa033e783e2a39d93e3a8718a7dc27ba95a4757f5433a01f794ec36"
	MARKETING_REWARD_SENT	= "dceb832c3abeeb39542a70bd553842dc873f15f6a5411c645092a93fef4ef9ba"

	/// Admin events
	CHARITY_PERCENTAGE_CHANGED		= "0918b499c15fcb0e087d411d53664cde23577e1aa4a9cbfbdf735ddd1871e7d5"
	PRIZE_PERCENTAGE_CHANGED	= "595fa5ba64cd6f66df19b53b59bb4a275bae1ec5b362da95e1fa4c5feb62be1e"
	RAFFLE_PERCENTAGE_CHANGED = "d2002a256ab2f8a5c1b865685754b20437c0a26e7249e40ae5df2993966f99f2"
	STAKE_PERCENTAGE_CHANGED = "9be9203c5d81ee2019998f7020cf412c92a35ba870318f660d4972660210fb4b"
	NUM_RAFFLE_ETH_WINNERS_PER_ROUND_CHANGED = "5e3567ae359e61c47ad2f11d2828dea7e062b2ff63dea385cdb7900a50504c7c"
	NUM_RAFFLE_NFT_WINNERS = "72e4278828b8a868e0ba5b4887f954797ea786f8bac991128636171f5eed471d"
	NUM_RAFFLE_NFT_HOLDERS = "0cc7e37c68566f67d0fe13bf38246d7447cf99a0c481c2ef9963969bb4f5ebc8"
	SYSTEM_MODE_CHANGED		= "f24e774cdaabee9b8782266728e442b7f1fa6ae9204755c0da1541e99f04aa4c"
	CHARITY_ADDRESS_CHANGED	= "1c7efd98583e33a9cc6adff48a97abdaaf43e5c9e918d5ec3e75e93e1dafcf6c"
	RWALK_ADDRESS_CHANGED	= "9dc3707e4b88be65295b0083b44ffa94059c80372add9b0d24d6a3b371a03b92"
	RAFFLE_ADDRESS_CHANGED	= "508d510537704d37ed64691a706392abc33e59ce8a327b4952e112820ceb88a6"
	STAKING_ADDRESS_CHANGED  = "3d112e567ad7f87ef5e5219a98118d33b03b247b007cfbadf4f133e7010f2c34"
	MARKETING_ADDRESS_CHANGED = "4d03942c29c20d1bccfe551e9d148c917c5a44fb558a4fc60270d8f76fb75f54"
	COSMIC_TOKEN_ADDRESS_CHANGED	= "3ab2dcf539aa3c30314265602fc86dee8e343e3c3905730f870835f36c400deb"
	COSMIC_SIGNATURE_ADDRESS_CHANGED	= "7142a592d5404b5fc4a294c66f70e32b2a7776bd807f722d59268def87c765d1"
	BUSINESS_LOGIC_ADDRESS_CHANGED	= "77ddb5e9e1495e15651bf87ccd8bbb7e637439fb260f0fda41b6ce4b3098aafd"
	TIME_INCREASE_CHANGED	= "ed46e73b174ced51fb529cdc6c1b4d8abf49387e6d849b71648afb63c81d12cd"
	TIMEOUT_CLAIMPRIZE_CHANGED	= "caf8e5c6bc1bb2c19935f84ddcdaefb232ad06f9f2abd2ad588bea4bbe631d26"
	PRICE_INCREASE_CHANGED	= "cbc1f49adfa29e2f2f0f5c9e057722496a4bc95a6a5446deaa423a02b30c64ac"
	NANOSECONDS_EXTRA_CHANGED = "678d086a09e1be49401b1e3a6e14db1878e8d8b88e71d0cfe24a32726d0e38b9"
	INITIAL_SECONDS_UNTIL_PRIZE_CHANGED = "6da281754ba85ee0c5983a8e8f05a92910c2a0c5b80e68c126216d65f162a305"
	INITIAL_BID_AMOUNT_FRACTION_CHANGED = "3b311f029da1b90c3de2e0b3168436c5ed7d8b0ae81b7d4894c12da03835c628"
	ACTIVATION_TIME_CHANGED = "584ff9a7b3a86db5397052f2e440da0ed60a95f646a3a884863cd92262e682b6"
	ETHCST_BID_RATIO_CHANGED = "22d4bc2aefb73ef04b923582d7ff2f2b74e699aaf7ff92ea24b348ed59161292"
	ROUND_START_AUCTION_LENGTH_CHANGED = "fa013ab3d3f17cb94a869ccc05fc842a00a83fd7318e55faeb0335c6d5d22aa3"
)
var (
	eclient 				*ethclient.Client
	rpcclient 				*rpc.Client

	// CosmicGame events:
	evt_prize_claim_event,_ = hex.DecodeString(PRIZE_CLAIM_EVENT)
	evt_bid_event,_			= hex.DecodeString(BID_EVENT)
	evt_donation_event,_	= hex.DecodeString(DONATION_EVENT)
	evt_nft_donation_event,_= hex.DecodeString(NFT_DONATION_EVENT)
	evt_raffle_nft_winner,_	= hex.DecodeString(RAFFLE_NFT_WINNER)
	evt_donated_nft_claimed,_= hex.DecodeString(DONATED_NFT_CLAIMED)
	evt_charity_percentage_changed,_= hex.DecodeString(CHARITY_PERCENTAGE_CHANGED)
	evt_prize_percentage_changed,_ = hex.DecodeString(PRIZE_PERCENTAGE_CHANGED)
	evt_raffle_percentage_changed,_ = hex.DecodeString(RAFFLE_PERCENTAGE_CHANGED)
	evt_staking_percentage_changed,_ = hex.DecodeString(STAKE_PERCENTAGE_CHANGED)
	evt_num_raffle_eth_winners_per_round_changed,_ = hex.DecodeString(NUM_RAFFLE_ETH_WINNERS_PER_ROUND_CHANGED)
	evt_num_raffle_nft_winners_per_round_changed,_ = hex.DecodeString(NUM_RAFFLE_NFT_WINNERS);
	evt_num_raffle_nft_holders_per_round_changed,_ = hex.DecodeString(NUM_RAFFLE_NFT_HOLDERS);
	evt_charity_address_changed,_	= hex.DecodeString(CHARITY_ADDRESS_CHANGED);
	evt_rwalk_address_changed,_	= hex.DecodeString(RWALK_ADDRESS_CHANGED);
	evt_raffle_address_changed,_	= hex.DecodeString(RAFFLE_ADDRESS_CHANGED);
	evt_staking_address_changed,_	= hex.DecodeString(STAKING_ADDRESS_CHANGED);
	evt_marketing_address_changed,_	= hex.DecodeString(MARKETING_ADDRESS_CHANGED);
	evt_costok_address_changed,_	= hex.DecodeString(COSMIC_TOKEN_ADDRESS_CHANGED);
	evt_cossig_address_changed,_	= hex.DecodeString(COSMIC_SIGNATURE_ADDRESS_CHANGED);
	evt_blogic_address_changed,_	= hex.DecodeString(BUSINESS_LOGIC_ADDRESS_CHANGED);
	evt_time_increase_changed,_	= hex.DecodeString(TIME_INCREASE_CHANGED);
	evt_timeout_claimprize_changed,_ = hex.DecodeString(TIMEOUT_CLAIMPRIZE_CHANGED);
	evt_price_increase_changed,_	= hex.DecodeString(PRICE_INCREASE_CHANGED);
	evt_nanoseconds_extra_changed,_	= hex.DecodeString(NANOSECONDS_EXTRA_CHANGED);
	evt_initial_seconds_until_prize_changed,_	= hex.DecodeString(INITIAL_SECONDS_UNTIL_PRIZE_CHANGED)
	evt_initial_bid_amount_fraction_changed,_	= hex.DecodeString(INITIAL_BID_AMOUNT_FRACTION_CHANGED)
	evt_activation_time_changed,_	= hex.DecodeString(ACTIVATION_TIME_CHANGED)
	evt_ethcst_bid_ratio_changed,_	= hex.DecodeString(ETHCST_BID_RATIO_CHANGED)
	evt_round_start_auction_length_changed,_ = hex.DecodeString(ROUND_START_AUCTION_LENGTH_CHANGED)
	evt_system_mode_changed,_ = hex.DecodeString(SYSTEM_MODE_CHANGED);

	// CharityWallet events
	evt_donation_received_event,_=hex.DecodeString(DONATION_RECEIVED_EVENT)
	evt_donation_sent_event,_= hex.DecodeString(DONATION_SENT_EVENT)
	evt_charity_updated,_	= hex.DecodeString(CHARITY_UPDATED)

	// CosmicSignature events
	evt_token_name_event,_	= hex.DecodeString(TOKEN_NAME_EVENT)
	evt_mint_event,_		= hex.DecodeString(MINT_EVENT)

	// RaffleWallet events
	evt_raffle_deposit,_	= hex.DecodeString(RAFFLE_DEPOSIT_EVENT)
	evt_raffle_withdrawal,_	= hex.DecodeString(RAFFLE_WITHDRAWAL_EVENT)

	// ERC20 events
	evt_transfer,_			= hex.DecodeString(TRANSFER_EVT)

	// StakingWallet events
	evt_stake_action,_		= hex.DecodeString(STAKE_ACTION_EVENT)
	evt_unstake_action,_	= hex.DecodeString(UNSTAKE_ACTION_EVENT)
	evt_claim_reward,_		= hex.DecodeString(CLAIM_REWARD_EVENT)
	evt_eth_deposit,_		= hex.DecodeString(ETH_DEPOSIT_EVENT)

	// MarketingWallet events
	evt_marketing_reward_sent,_		= hex.DecodeString(MARKETING_REWARD_SENT)

	inspected_events []InspectedEvent

	cosmic_game_abi			*abi.ABI
	cosmic_signature_abi	*abi.ABI
	cosmic_token_abi		*abi.ABI
	charity_wallet_abi		*abi.ABI
	raffle_wallet_abi		*abi.ABI
	staking_wallet_abi		*abi.ABI
	marketing_wallet_abi	*abi.ABI
	erc20_abi				*abi.ABI
	erc721_abi				*abi.ABI

	cosmic_game_addr		common.Address
	cosmic_signature_addr	common.Address
	cosmic_token_addr		common.Address
	cosmic_dao_addr			common.Address
	charity_wallet_addr		common.Address
	raffle_wallet_addr		common.Address
	staking_wallet_addr		common.Address
	marketing_wallet_addr	common.Address
	cosmic_sig_aid			int64

	cg_contracts			CosmicGameContractAddrs
	storagew				SQLStorageWrapper
	RPC_URL					 = os.Getenv("RPC_URL")
	Error					*log.Logger
	Info					*log.Logger
)

func get_event_ids(from_evt_id,to_evt_id int64) []int64 {
	output := make([]int64 ,0, 1024)
	for _,e := range inspected_events {
		var event_list []int64
		event_list = storagew.S.Get_evtlogs_by_signature_only_in_range(
				e.Signature,from_evt_id,to_evt_id,
		)
		output = append(output,event_list...)
	}
	sort.Slice(output, func(i, j int) bool { return output[i] < output[j] })
	num_elts:=Remove_duplicates_int64(output)
	return output[0:num_elts]
}
func process_events(exit_chan chan bool) {

	var max_batch_size int64 = 1024*200
	for {
		status := storagew.Get_cosmic_game_processing_status()
		select {
			case exit_flag := <-exit_chan:
				if exit_flag {
					Info.Println("Exiting by user request.")
					os.Exit(0)
				}
			default:
		}
	/*	Info.Printf(
			"scanning event range from %v to %v\n",
			status.LastEvtIdProcessed,status.LastEvtIdProcessed+max_batch_size,
		)*/
		id_upper_limit := status.LastEvtIdProcessed + max_batch_size
		last_evt_id,err := storagew.S.Get_last_evtlog_id()
		if err != nil {
			Error.Printf("Error: %v. Possibly 'evt_log' table is empty, aborting",err)
			os.Exit(1)
		}
		if  id_upper_limit > last_evt_id {
			id_upper_limit = last_evt_id
		}
		events := get_event_ids(status.LastEvtIdProcessed,id_upper_limit)
		for _,evt_id := range events {
			err := process_single_event(evt_id)
			if err != nil {
				Error.Printf("Pausing event processing loop for 5 sec due to error")
				time.Sleep(5 * time.Second)
				break
			}
			status.LastEvtIdProcessed=evt_id
			storagew.Update_cosmic_game_process_status(&status)
		}
		if len(events) == 0 {
			status.LastEvtIdProcessed = id_upper_limit
			storagew.Update_cosmic_game_process_status(&status)
			time.Sleep(1 * time.Second) // sleep only if there is no data
		}
	}
}
func get_abi(abi_str string) *abi.ABI {
	abi_parsed := strings.NewReader(abi_str)
	abi_obj,err := abi.JSON(abi_parsed)
	if err!= nil {
		Info.Printf("Can't parse ABI: %v\n",err)
		os.Exit(1)
	}
	return &abi_obj
}
func main() {

	log_dir:=fmt.Sprintf("%v/%v",os.Getenv("HOME"),DEFAULT_LOG_DIR)
	os.MkdirAll(log_dir, os.ModePerm)
	db_log_file:=fmt.Sprintf("%v/cosmicgame_%v",log_dir,DEFAULT_DB_LOG)

	fname:=fmt.Sprintf("%v/cosmicgame_info.log",log_dir)
	logfile, err := os.OpenFile(fname, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err!=nil {
		fmt.Printf("Can't start: %v\n",err)
		os.Exit(1)
	}
	Info = log.New(logfile,"INFO: ",log.Ltime|log.Lshortfile)

	fname=fmt.Sprintf("%v/cosmicgame_error.log",log_dir)
	if err!=nil {
		fmt.Printf("Can't start: %v\n",err)
		os.Exit(1)
	}
	logfile, err = os.OpenFile(fname, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	Error = log.New(logfile,"ERROR: ",log.Ltime|log.Lshortfile)

	rpcclient, err=rpc.DialContext(context.Background(), RPC_URL)
	if err != nil {
		log.Fatal(err)
	}
	Info.Printf("Connected to ETH node: %v\n",RPC_URL)
	eclient = ethclient.NewClient(rpcclient)

	storagew.S = Connect_to_storage(Info)
	storagew.S.Db_set_schema_name("public");
	storagew.S.Init_log(db_log_file)
	storagew.S.Log_msg("Log initialized\n")

	cosmic_game_abi = get_abi(CosmicGameABI)
	cosmic_signature_abi = get_abi(CosmicSignatureABI)
	cosmic_token_abi = get_abi(CosmicTokenABI)
	charity_wallet_abi = get_abi(CharityWalletABI);
	raffle_wallet_abi = get_abi(RaffleWalletABI);
	staking_wallet_abi = get_abi(StakingWalletABI);
	marketing_wallet_abi = get_abi(MarketingWalletABI);
	erc20_abi = get_abi(ERC20ABI)
	erc721_abi = get_abi(ERC721ABI)

	cg_contracts = storagew.Get_cosmic_game_contract_addrs()
	cosmic_sig_aid,err  = storagew.S.Nonfatal_lookup_address_id(cg_contracts.CosmicSignatureAddr)
	if err != nil {
		fmt.Printf("Lookup of CosmicSignatureAddr failed: %v",err)
		os.Exit(1)
	}
	cosmic_game_addr = common.HexToAddress(cg_contracts.CosmicGameAddr)
	cosmic_signature_addr = common.HexToAddress(cg_contracts.CosmicSignatureAddr)
	cosmic_token_addr = common.HexToAddress(cg_contracts.CosmicTokenAddr)
	cosmic_dao_addr = common.HexToAddress(cg_contracts.CosmicDaoAddr)
	charity_wallet_addr = common.HexToAddress(cg_contracts.CharityWalletAddr)
	raffle_wallet_addr = common.HexToAddress(cg_contracts.RaffleWalletAddr)
	staking_wallet_addr = common.HexToAddress(cg_contracts.StakingWalletAddr)
	marketing_wallet_addr = common.HexToAddress(cg_contracts.MarketingWalletAddr)

	c := make(chan os.Signal)
	exit_chan := make(chan bool)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-c
		Info.Printf("Got SIGINT signal, will exit after processing is over." +
					" To interrupt abruptly send SIGKILL (9) to the kernel.\n")
		exit_chan <- true
	}()


	inspected_events = build_list_of_inspected_events_layer1(cosmic_sig_aid)
	process_events(exit_chan)
}
