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

	PRIZE_CLAIM_EVENT		= "4e6dc8ff50108e18c5ebeabb472c87e32464277e1aafd81888c8ac1b4cdde672"
	BID_EVENT				= "bcb004d688d0951e50c218ded0d0d574bde915630e29b92987b1f2eab9556549"
	DONATION_EVENT			= "e32cacf203d00685e2b4d8b0a90e7cd8f3f8a208fdf116f4bb36abe08b7d548e"
	DONATION_WITH_INFO_EVENT= "a08049565b10d44a06dca9bf05685b39bc370352043c5a003e8d35d45ebdc53f"
	DONATION_RECEIVED_EVENT	= "264f630d9efa0d07053a31163641d9fcc0adafc9d9e76f1c37c2ce3a558d2c52"
	DONATION_SENT_EVENT		= "1222634ba80f397fa40371bab63974e6f97da8200777ec79e731c69bb6a2351d"
	CHARITY_UPDATED			= "1c7efd98583e33a9cc6adff48a97abdaaf43e5c9e918d5ec3e75e93e1dafcf6c"
	TOKEN_NAME_EVENT		= "a14cfb0fe69c0b55eaaa4d9400bdba2a72e1860cade89c2a8a055e6cfde8936d"
	MINT_EVENT				= "c2115f21464937bfdcd1560f96f0e20b70e88accbdcd1069084c80c8797ef106"
	NFT_DONATION_EVENT		= "b12e72bab0c2a8fe98175a3c02792645ebdf65d37cf3983517e6d1c5ab842e23"
	ERC20_DONATED			= "3f94f6171c96ab491d8c268f80da5e460b148878065711a0bb099bbc297615af"
	DONATED_TOKEN_CLAIMED	= "af1adae2e1e983ca738335dc2e37194114142793f394ac934a45ea632a8a5bd0"
	DONATED_NFT_CLAIMED		= "03c2b6e01c9ca39e4073132f279b61b987d41a82b08cea3dd4a2fabc37067ac3"
//	ETH_DEPOSIT_EVENT		= "85177f287940f2f05425a4029951af0e047a7f9c4eaa9a6e6917bcd869f86695"
	ETH_PRIZE_DEPOSIT_EVENT		= "999946acc98c7b7dacc26921697d55abbcb1637484b0a73040f0b06287e43326"
	ETH_PRIZE_WITHDRAWAL_EVENT = "4f43b861ba36494acfe938f3815fba7fac6981bdc611b6ccdc14c08f59292383"
	RAFFLE_ETH_WINNER		= "9c62e2cb8cbd10bf3b8a5760977d719fc3475fa67cb7ef9d2e1701f275e885c4"
	RAFFLE_NFT_WINNER		= "b74b8b17fe663bfa97ab0bd69aa757a370a7e071e9a9171a3fae7aa8a8a850a5"
	ENDURANCE_WINNER		= "838ec9dd2530548892bff113f5ffb0138d2efc63c7f59bb6571e8c923b631260"	//EnduranceChampionPrizePaid
	LASTCST_BIDDER_WINNER	= "3901b6430c99dc290ee88ff84c4de6091ad7eac335b58e92ef5cbb0793abf4f6"
	CHRONO_WARRIOR			= "4325f4809052852398280e22eac8be724fa5616bd6cad1bc20fdcbcd114a3d75"	//ChronoWarriorPrizeAllocated
	TRANSFER_EVT			= "ddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef"
	//STAKE_ACTION_EVENT		= "cd7bda73d32f8fada3eeee8d21563aa19eb2fe86d90b4449cf5252e6f3da7ef8"
	NFT_STAKED_EVENT		= "cbd8e5368101e7829397d044213d371dac0f6727ea4bdc42d2f5a1f2e2f00829"
//	UNSTAKE_ACTION_EVENT	= "678afbb7bbf1c4f3df509d59b51d6e75969703762eb36ed41414dc7c49569c96"
	NFT_UNSTAKED_RWALK		= "1792a7a9d5e2f53a8d06f2ae40b4446d91b503e84dd7f6307f40cdeb11541668"
	NFT_UNSTAKED_CST		= "ec478a78e4e3931ff728a54eeb6875304c891fa5fa253337b60d37fdc5e1feaf"
	STAKING_ETH_DEPOSIT_EVENT= "b71b1087ee7f659cf742c29d8095c562e6e832337190e10dbe81db89955c0913"
	CLAIM_REWARD_EVENT		= "dde81df5caa033e783e2a39d93e3a8718a7dc27ba95a4757f5433a01f794ec36"
	REWARD_PAID_EVENT		= "f9bac74bc321a00ef2afeb1f44684e20e22f567699a5840df47967ea88c5b449"
	FUND_TRANSFER_ERR		= "154fb6c686c977314af35b730a16b83facda5265b2abec7237c730f63e42215a"
	ERC20_TRANSFER_ERR		= "f7fce645f12ae266a329c431e96ebea892316a1415809056621ffeea04efd4ab"
	ROUND_STARTED			= "028a52641badd593b7f30072734c0b97e449213f55b5c3663756427340accd3c"	//FirstBidPlacedInRound

	/// Admin events
	PROXY_UPGRADED			= "bc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b"
	ADMIN_CHANGED			= "7e644d79422f17c01e4894b5f4f588d331ebfa28653d42ae832dc59e38c9798f"	// AdminChanged (ERC1967Utils.sol)
	INITIALIZED				= "c7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2"
	CHARITY_PERCENTAGE_CHANGED		= "fe65b6d5007c66dadebc5095104ccd672c070f396dfdcfe1ee7f54201b9efa6d"
	PRIZE_PERCENTAGE_CHANGED	= "b5a05ec7911dd5450a7fa4ae54d50d9d83af6e256db8fc76c82edd7b659cf8bc"
	RAFFLE_PERCENTAGE_CHANGED = "69f3b6a0324a65c444ac42746c7084ea54bd6c7553de10f7ec7bc8295df21698"
	STAKE_PERCENTAGE_CHANGED = "a7a96cf336831afee0ab3de72b71a5af8c439675b697b47e08283c52ec4679ef"
	CHRONO_PERCENTAGE_CHANGED = "5581e31f5b8d4b3c45b8ab8bf67e3602ce588b423905eb6ad34bd6bc3c848699"
	NUM_RAFFLE_ETH_WINNERS_BIDDING_CHANGED = "4787028773c8f14bc3b4bc41f43a02329ae41105823287201a34782c530d35fd"
	NUM_RAFFLE_NFT_WINNERS_BIDDING_CHANGED = "85d8bf21006916836edd67a5afeed2e891cf876a6c9cd9babf7f42f4b007c24f"
	NUM_RAFFLE_NFT_WINNERS_STAKING_RWALK_CHANGED = "3312247fcf207243294680d9103851c8709e19be3d46ee4b1bff5e12d1c5ca7b"
	SYSTEM_MODE_CHANGED		= "f24e774cdaabee9b8782266728e442b7f1fa6ae9204755c0da1541e99f04aa4c"
	CHARITY_ADDRESS_CHANGED	= "1c7efd98583e33a9cc6adff48a97abdaaf43e5c9e918d5ec3e75e93e1dafcf6c"
	RWALK_ADDRESS_CHANGED	= "dab38e33e6e11cbb8b085bba9d7426d5e5af01bcc46d2c5957645e4d09e8c49c"
	PRIZE_WALLET_ADDRESS_CHANGED	= "b4cecfe1346c94da27291cf5a02969d5fe0b5c36eca49b04fcd60841d28c5e13"
	STAKING_WALLET_CST_ADDRESS_CHANGED  = "4da1815cd654922275d14d2335fd9a0dd0aa6a0d0ff87fb4cc872ebe9704596f"
	STAKING_WALLET_RWALK_ADDRESS_CHANGED  = "bf6e296f85d08cc1ab124aed644bf4b19e4a726a7aea53e3784ab1341738a040"
	MARKETING_ADDRESS_CHANGED = "4d03942c29c20d1bccfe551e9d148c917c5a44fb558a4fc60270d8f76fb75f54"
	COSMIC_TOKEN_ADDRESS_CHANGED	= "2ea51ddc5dacc666588569f7c6d26d9f79fe6bc7fea7cc7d89bcd6e38cdfb421"
	COSMIC_SIGNATURE_ADDRESS_CHANGED	= "5bde6238168795ba4e77972a2bdaa5a465f7c9a5d05817f5e8d3fed2e5a4fa60"
	BUSINESS_LOGIC_ADDRESS_CHANGED	= "77ddb5e9e1495e15651bf87ccd8bbb7e637439fb260f0fda41b6ce4b3098aafd"
	TIME_INCREASE_CHANGED	= "ed46e73b174ced51fb529cdc6c1b4d8abf49387e6d849b71648afb63c81d12cd"
	TIMEOUT_CLAIMPRIZE_CHANGED	= "37a332914fac995349420c0419b4423a19dcb762017f691442a0782ce4bf417a"
	TIMEOUT_TO_WITHDRAW_PRIZE = "8717bb199c6bc4a5dadb21547205f9ef8ec037dda246a5526d6a6471306ea52e"
	PRICE_INCREASE_CHANGED	= "a7c2c10f71f5776e9274342cd7414dfc1d0ad560dd3111fcddcb92ec2fa8575d"	//NextEthBidPriceIncreaseDivisorChanged
	MAIN_PRIZE_MICROSECOND_INCREASE= "07417920574ce0bdfe987af0575c8793cc73a29d7830760ad459d0e569b5b79b"//	NANOSECONDS_EXTRA_CHANGED = "07417920574ce0bdfe987af0575c8793cc73a29d7830760ad459d0e569b5b79b" // the variable was renamed to mainPrizeTimeIncrementInMicroSeconds
	INITIAL_SECONDS_UNTIL_PRIZE_CHANGED = "b5edd1f338b34c8f5dd3b1c5cc12f05653c495713c282bf588d34cf14fad0f89"
//	INITIAL_BID_AMOUNT_FRACTION_CHANGED = "3b311f029da1b90c3de2e0b3168436c5ed7d8b0ae81b7d4894c12da03835c628"		DISCONTINUED
	ROUND_ACTIVATION_TIME_CHANGED = "9a2159c1f277ddd727551baedc6a6c4cba77cc5219c8563ee3b15fb67548d89b"
	ETH_DUTCH_AUCTION_DURATION_DIVISOR_CHANGED=	"fdf6043c89a5f304289122dcc9f8bd78bb111b5d4f409e2fc2e6c141a1110b79"
	CST_DUTCH_AUCTION_DURATION_DIVISOR_CHANGED= "c95d03f6c735a9e59c760fdb88e585aafe0a31b5c034fc7838155287ee32212f" // previosly ROUND_START_AUCTION_LENGTH_CHANGED = "23dabd88e0a182dcd593bec053f3867f1bd6afc77d470cbc1ef48ad189bfd676"
	MARKETING_REWARD_SENT	= "e2403640ba68fed3a2f88b7557551d1993f84b99bb10ff833f0cf8db0c5e0486"
	MARKETING_REWARD_CHANGED = "aa59cda80c9b70b72f2ab15997b2622a0f94b107c401abfdc0f01f6f36489221"
	CST_REWARD_FOR_BIDDING_CHANGED = "70ad04ce09c925ea466a5f603054f310bba5b7484bba77b382aade0bf93b55d0"	// previous: ERC20_TOKEN_REWARD		= "762f994f6c24fece9d12a1eba3630058b2a0d8cb551a6496ef6e128aedf86353"
	ERC20_REWARD_MULT		= "44d50377242b2c165fd7ae3c2a9f2ccad8ecf04268512599ba6b81dedec0a59b"	//CstRewardAmountMultiplierChanged
	MAX_MESSAGE_LENGTH		= "157c413b0549fd4f45aab72b7828304fb2c45dad53de0f1128c5eabf3aaabaf8"	//BidMessageLengthMaxLimitChanged
	TOKEN_SCRIPT_URL		= "27e2bd70f498920ee0fd7d8204ae8845b75dc81330e3acafa32946be3503730c"
	BASE_URI				= "bdfd815215fcee5bb949c941ab489c7ead076a7c8acd3527cd1b50f613ac67e6"
	OWNERSHIP_TRANSFERRED	= "8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0"
	STARTING_CST_MIN_LIM	= "4e8c80fe79d13b8663de9f6981925ae24c498cc07b0ebd4f4789fa78ca17caff"
	FUNDS_TO_CHARITY		= "1222634ba80f397fa40371bab63974e6f97da8200777ec79e731c69bb6a2351d"
	DELAY_DURATION_ROUND	= "b0868a729f47ae3829aaafe3ca2975d3db2148553c854112f598be6d91ef0d28"
	STAKING_STATE_RESET		= "8a77a30ae54ca88ceda5f17cd3679c1b868abd2c6d36b67bcbe5e9beac7ec648"

)
var (
	eclient 				*ethclient.Client
	rpcclient 				*rpc.Client

	// CosmicGame events:
	evt_prize_claim_event,_ = hex.DecodeString(PRIZE_CLAIM_EVENT)
	evt_bid_event,_			= hex.DecodeString(BID_EVENT)
	evt_donation_event,_	= hex.DecodeString(DONATION_EVENT)
	evt_donation_with_info_event,_= hex.DecodeString(DONATION_WITH_INFO_EVENT)
	evt_nft_donation_event,_= hex.DecodeString(NFT_DONATION_EVENT)
	evt_erc20_donated,_		= hex.DecodeString(ERC20_DONATED)
	evt_raffle_nft_winner,_	= hex.DecodeString(RAFFLE_NFT_WINNER)
	evt_raffle_eth_winner,_	= hex.DecodeString(RAFFLE_ETH_WINNER)
	evt_endurance_winner,_	= hex.DecodeString(ENDURANCE_WINNER)
	evt_lastcst_bidder_winner,_	= hex.DecodeString(LASTCST_BIDDER_WINNER)
	evt_chrono_warrior,_	= hex.DecodeString(CHRONO_WARRIOR)
	evt_donated_token_claimed,_	= hex.DecodeString(DONATED_TOKEN_CLAIMED)
	evt_donated_nft_claimed,_= hex.DecodeString(DONATED_NFT_CLAIMED)
	evt_charity_percentage_changed,_= hex.DecodeString(CHARITY_PERCENTAGE_CHANGED)
	evt_prize_percentage_changed,_ = hex.DecodeString(PRIZE_PERCENTAGE_CHANGED)
	evt_raffle_percentage_changed,_ = hex.DecodeString(RAFFLE_PERCENTAGE_CHANGED)
	evt_staking_percentage_changed,_ = hex.DecodeString(STAKE_PERCENTAGE_CHANGED)
	evt_chrono_percentage_changed,_ = hex.DecodeString(CHRONO_PERCENTAGE_CHANGED)
	evt_num_raffle_eth_winners_bidding_changed,_ = hex.DecodeString(NUM_RAFFLE_ETH_WINNERS_BIDDING_CHANGED)
	evt_num_raffle_nft_winners_bidding_changed,_ = hex.DecodeString(NUM_RAFFLE_NFT_WINNERS_BIDDING_CHANGED)
	evt_num_raffle_nft_winners_staking_rwalk_changed,_ = hex.DecodeString(NUM_RAFFLE_NFT_WINNERS_STAKING_RWALK_CHANGED)
	evt_charity_address_changed,_	= hex.DecodeString(CHARITY_ADDRESS_CHANGED)
	evt_rwalk_address_changed,_	= hex.DecodeString(RWALK_ADDRESS_CHANGED)
	evt_prizes_wallet_address_changed,_	= hex.DecodeString(PRIZE_WALLET_ADDRESS_CHANGED)
	evt_staking_wallet_cst_address_changed,_	= hex.DecodeString(STAKING_WALLET_CST_ADDRESS_CHANGED)
	evt_staking_wallet_rwalk_address_changed,_	= hex.DecodeString(STAKING_WALLET_RWALK_ADDRESS_CHANGED)
	evt_marketing_address_changed,_	= hex.DecodeString(MARKETING_ADDRESS_CHANGED)
	evt_costok_address_changed,_	= hex.DecodeString(COSMIC_TOKEN_ADDRESS_CHANGED)
	evt_cossig_address_changed,_	= hex.DecodeString(COSMIC_SIGNATURE_ADDRESS_CHANGED)
	evt_time_increase_changed,_	= hex.DecodeString(TIME_INCREASE_CHANGED)
	evt_timeout_claimprize_changed,_ = hex.DecodeString(TIMEOUT_CLAIMPRIZE_CHANGED)
	evt_timeout_to_withdraw_prize,_	= hex.DecodeString(TIMEOUT_TO_WITHDRAW_PRIZE)
	evt_price_increase_changed,_	= hex.DecodeString(PRICE_INCREASE_CHANGED)
	evt_prize_microsecond_increase_changed,_	= hex.DecodeString(MAIN_PRIZE_MICROSECOND_INCREASE)
	evt_initial_seconds_until_prize_changed,_	= hex.DecodeString(INITIAL_SECONDS_UNTIL_PRIZE_CHANGED)
//	evt_initial_bid_amount_fraction_changed,_	= hex.DecodeString(INITIAL_BID_AMOUNT_FRACTION_CHANGED)		DISCONTINUED
	evt_activation_time_changed,_	= hex.DecodeString(ROUND_ACTIVATION_TIME_CHANGED)
	evt_cst_dutch_auction_duration_divisor_changed,_ = hex.DecodeString(CST_DUTCH_AUCTION_DURATION_DIVISOR_CHANGED) // previously	evt_round_start_auction_length_changed
	evt_eth_dutch_auction_duration_divisor_changed,_ = hex.DecodeString(ETH_DUTCH_AUCTION_DURATION_DIVISOR_CHANGED)
	evt_system_mode_changed,_ = hex.DecodeString(SYSTEM_MODE_CHANGED)
	evt_proxy_upgraded,_	= hex.DecodeString(PROXY_UPGRADED)
	evt_admin_changed,_		= hex.DecodeString(ADMIN_CHANGED)
	evt_cst_reward_for_bidding_changed,_	= hex.DecodeString(CST_REWARD_FOR_BIDDING_CHANGED)
	evt_max_msg_length_changed,_	= hex.DecodeString(MAX_MESSAGE_LENGTH)
	evt_token_script_url,_			= hex.DecodeString(TOKEN_SCRIPT_URL)
	evt_base_uri,_					= hex.DecodeString(BASE_URI)
	evt_marketing_reward_changed,_	= hex.DecodeString(MARKETING_REWARD_CHANGED)
	evt_ownership_transferred,_		= hex.DecodeString(OWNERSHIP_TRANSFERRED)
	evt_initialized,_		= hex.DecodeString(INITIALIZED)
	evt_cst_min_limit,_				= hex.DecodeString(STARTING_CST_MIN_LIM)
	evt_fund_transf_err,_		= hex.DecodeString(FUND_TRANSFER_ERR)
	evt_erc20_transf_err,_			= hex.DecodeString(ERC20_TRANSFER_ERR)
	evt_erc20_reward_mult,_				= hex.DecodeString(ERC20_REWARD_MULT)
	evt_funds2charity,_				= hex.DecodeString(FUNDS_TO_CHARITY)
	evt_delay_duration_round,_		= hex.DecodeString(DELAY_DURATION_ROUND)
	evt_round_started,_				= hex.DecodeString(ROUND_STARTED)
	evt_staking_state_reset,_		= hex.DecodeString(STAKING_STATE_RESET)

	// CharityWallet events
	evt_donation_received_event,_=hex.DecodeString(DONATION_RECEIVED_EVENT)
	evt_donation_sent_event,_= hex.DecodeString(DONATION_SENT_EVENT)
	evt_charity_updated,_	= hex.DecodeString(CHARITY_UPDATED)

	// CosmicSignature events
	evt_token_name_event,_	= hex.DecodeString(TOKEN_NAME_EVENT)
	evt_mint_event,_		= hex.DecodeString(MINT_EVENT)

	// PrizeWallet events
	evt_eth_prize_deposit,_		= hex.DecodeString(ETH_PRIZE_DEPOSIT_EVENT)
	evt_eth_prize_withdrawal,_	= hex.DecodeString(ETH_PRIZE_WITHDRAWAL_EVENT)

	// ERC20 events
	evt_transfer,_			= hex.DecodeString(TRANSFER_EVT)

	// StakingWallet events
	evt_nft_staked,_		= hex.DecodeString(NFT_STAKED_EVENT)
	evt_nft_unstaked_rwalk,_= hex.DecodeString(NFT_UNSTAKED_RWALK)
	evt_nft_unstaked_cst,_	= hex.DecodeString(NFT_UNSTAKED_CST)
	evt_reward_paid,_		= hex.DecodeString(REWARD_PAID_EVENT)
	evt_staking_eth_deposit,_		= hex.DecodeString(STAKING_ETH_DEPOSIT_EVENT)

	// MarketingWallet events
	evt_marketing_reward_sent,_		= hex.DecodeString(MARKETING_REWARD_SENT)

	inspected_events []InspectedEvent

	cosmic_game_abi			*abi.ABI
	cosmic_signature_abi	*abi.ABI
	cosmic_token_abi		*abi.ABI
	charity_wallet_abi		*abi.ABI
	prizes_wallet_abi		*abi.ABI
	staking_wallet_cst_abi		*abi.ABI
	staking_wallet_rwalk_abi		*abi.ABI
	marketing_wallet_abi	*abi.ABI
	erc20_abi				*abi.ABI
	erc721_abi				*abi.ABI

	cosmic_game_addr		common.Address
	cosmic_signature_addr	common.Address
	cosmic_token_addr		common.Address
	cosmic_dao_addr			common.Address
	charity_wallet_addr		common.Address
	prizes_wallet_addr		common.Address
	staking_wallet_cst_addr		common.Address
	staking_wallet_rwalk_addr		common.Address
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

	cosmic_game_abi = get_abi(CosmicSignatureGameABI)
//	blogic_abi = get_abi(BusinessLogicABI);
	cosmic_signature_abi = get_abi(CosmicSignatureNftABI)
	cosmic_token_abi = get_abi(CosmicSignatureTokenABI)
	charity_wallet_abi = get_abi(CharityWalletABI);
	prizes_wallet_abi = get_abi(PrizesWalletABI);
	staking_wallet_cst_abi = get_abi(IStakingWalletCosmicSignatureNftABI);
	staking_wallet_rwalk_abi = get_abi(IStakingWalletRandomWalkNftABI);
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
	prizes_wallet_addr = common.HexToAddress(cg_contracts.PrizesWalletAddr)
	staking_wallet_cst_addr = common.HexToAddress(cg_contracts.StakingWalletCSTAddr)
	staking_wallet_rwalk_addr = common.HexToAddress(cg_contracts.StakingWalletRWalkAddr)
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
