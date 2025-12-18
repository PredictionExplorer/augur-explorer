package main

import (
	"context"
	"encoding/hex"
	"fmt"
	"log"
	"os"
	"os/signal"
	"strings"
	"syscall"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi"
	ethcommon "github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/rpc"

	. "github.com/PredictionExplorer/augur-explorer/rwcg/contracts/cosmicgame"
	. "github.com/PredictionExplorer/augur-explorer/rwcg/dbs"
	. "github.com/PredictionExplorer/augur-explorer/rwcg/dbs/cosmicgame"
	etlcommon "github.com/PredictionExplorer/augur-explorer/rwcg/etl/common"
	. "github.com/PredictionExplorer/augur-explorer/rwcg/primitives"
	. "github.com/PredictionExplorer/augur-explorer/rwcg/primitives/cosmicgame"
)
const (
	DEFAULT_DB_LOG			= "db.log"
	IMGGEN_PATH				= "v2/etl/cmd/cosmicgame/imggen_monitor/imggen_exec" // relative to $HOME

	PRIZE_CLAIM_EVENT		= "8c551ec2b6f186753e27f1cf46f84b57f4f83f721e8c1e6170ae512845ced591" // ICosmicSignatureGame.sol:MainPrizeClaimed
	BID_EVENT				= "bcb004d688d0951e50c218ded0d0d574bde915630e29b92987b1f2eab9556549" // ICosmicSignatureGame.sol:BidPlaced
	ETH_DONATED_EVENT		= "e32cacf203d00685e2b4d8b0a90e7cd8f3f8a208fdf116f4bb36abe08b7d548e" // IEthDonations.sol:EthDonated
	ETH_DONATED_WI_EVENT	= "a08049565b10d44a06dca9bf05685b39bc370352043c5a003e8d35d45ebdc53f" // IEthDonations.sol:EthDonatedWithInfo
	DONATION_RECEIVED_EVENT	= "264f630d9efa0d07053a31163641d9fcc0adafc9d9e76f1c37c2ce3a558d2c52" // ICharityWallet.sol:DonationReceived
	DONATION_SENT_EVENT		= "1222634ba80f397fa40371bab63974e6f97da8200777ec79e731c69bb6a2351d" // ICharityWallet.sol:DonationSent
	CHARITY_RECEIVER_CHANGED = "1c7efd98583e33a9cc6adff48a97abdaaf43e5c9e918d5ec3e75e93e1dafcf6c" // ICharityWallet.sol:CharityAddressChanged
	CHARITY_WALLET_CHANGED	= "1c7efd98583e33a9cc6adff48a97abdaaf43e5c9e918d5ec3e75e93e1dafcf6c" // ISystemEvents.sol:CharityAddressChanged
	TOKEN_NAME_EVENT		= "a14cfb0fe69c0b55eaaa4d9400bdba2a72e1860cade89c2a8a055e6cfde8936d" // ICosmicSignatureNft.sol:NftNameChanged
	MINT_EVENT				= "c2115f21464937bfdcd1560f96f0e20b70e88accbdcd1069084c80c8797ef106" // ICosmicSignatureNft.sol:NftMinted
	NFT_ETH_DONATED_EVENT		= "b12e72bab0c2a8fe98175a3c02792645ebdf65d37cf3983517e6d1c5ab842e23" // IPrizesWallet.sol:NftDonated
	ERC20_DONATED			= "3f94f6171c96ab491d8c268f80da5e460b148878065711a0bb099bbc297615af" // IPrizesWallet.sol:TokenDonated
	DONATED_TOKEN_CLAIMED	= "af1adae2e1e983ca738335dc2e37194114142793f394ac934a45ea632a8a5bd0" // IPrizesWallet.sol:DonatedTokenClaimed
	DONATED_NFT_CLAIMED		= "03c2b6e01c9ca39e4073132f279b61b987d41a82b08cea3dd4a2fabc37067ac3" // IPrizesWallet.sol:DonatedNftClaimed
	ETH_PRIZE_DEPOSIT_EVENT		= "8e369548f9ef214c7c1823c098a23763321fd761ac5cf78958e1db1b3648e7a2" // IPrizesWallet.sol:EthReceived
	ETH_PRIZE_WITHDRAWAL_EVENT = "172b54ba94575dba1c8dda35e4a6f6a0f761efe5c5416227b34c7c7632a673d4" // IPrizesWallet.sol:EthWithdrawn
	RAFFLE_ETH_PRIZE_EVENT	= "9c62e2cb8cbd10bf3b8a5760977d719fc3475fa67cb7ef9d2e1701f275e885c4" // ICosmicSignatureGame.sol:RaffleWinnerBidderEthPrizeAllocated
	RAFFLE_NFT_PRIZE_EVENT	= "27c21fe4cea1a3367aa491829dd4dd824296c00910626150464cba8ea5ebb3f4" // ICosmicSignatureGame.sol:RaffleWinnerPrizePaid
	ENDURANCE_PRIZE_EVENT	= "838ec9dd2530548892bff113f5ffb0138d2efc63c7f59bb6571e8c923b631260" // ICosmicSignatureGame.sol:EnduranceChampionPrizePaid
	LASTCST_BIDDER_PRIZE_EVENT = "3901b6430c99dc290ee88ff84c4de6091ad7eac335b58e92ef5cbb0793abf4f6" // ICosmicSignatureGame.sol:LastCstBidderPrizePaid
	CHRONO_WARRIOR_PRIZE_EVENT = "aa858ae20a26d00a9ea528972d537e68a51a0744226d9ea1fc9b22492dc282a5" // ICosmicSignatureGame.sol:ChronoWarriorPrizePaid
	TRANSFER_EVT			= "ddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef" // IERC721.sol:Transfer
	CST_NFT_STAKED_EVENT		= "e09cd972bcd125457d8f8a684b2b67ec513fbb7f770001bbebd7c22b41ad9da8" // IStakingWalletNftBase.sol:NftStaked
	RWALK_NFT_STAKED_EVENT		= "62773741191803b5cec48480156933243e422a1fb1ea9967dab3ee30df2da95f" // IStakingWalletNftBase.sol:NftStaked
	NFT_UNSTAKED_RWALK		= "08e7047cd3ef25f4a72589ed1d73eb0af1ca8a5957f9539e126dff45870979cc" // IStakingWalletNftBase.sol:NftUnstaked
	NFT_UNSTAKED_CST		= "ec478a78e4e3931ff728a54eeb6875304c891fa5fa253337b60d37fdc5e1feaf" // IStakingWalletNftBase.sol:NftUnstaked
	STAKING_ETH_DEPOSIT_EVENT= "26726e1a40953e6b9e06e1c1a1f53422299c7e00d2cc5d986bd4723796ca4588" // IStakingWalletCosmicSignatureNft.sol:EthDepositReceived
	CLAIM_REWARD_EVENT		= "dde81df5caa033e783e2a39d93e3a8718a7dc27ba95a4757f5433a01f794ec36" // IStakingWalletNftBase.sol:RewardClaimed
	FUND_TRANSFER_ERR		= "154fb6c686c977314af35b730a16b83facda5265b2abec7237c730f63e42215a" // ICosmicSignatureErrors.sol:FundTransferFailed
	ERC20_TRANSFER_ERR		= "f7fce645f12ae266a329c431e96ebea892316a1415809056621ffeea04efd4ab" // ICosmicSignatureErrors.sol:ERC20TransferFailed
	FIRST_BID_EVENT			= "028a52641badd593b7f30072734c0b97e449213f55b5c3663756427340accd3c" // ICosmicSignatureGame.sol:FirstBidPlacedInRound

	/// Admin events
	PROXY_UPGRADED			= "bc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b" // IERC1967.sol:Upgraded
	ADMIN_CHANGED			= "7e644d79422f17c01e4894b5f4f588d331ebfa28653d42ae832dc59e38c9798f" // IERC1967.sol:AdminChanged
	TREASURER_CHANGED		= "df73fc12cc071a4834f7ba0e7c6cfe7d23e98866e191ec9e86e6e61614d9e50e" // IMarketingWallet.sol:TreasurerAddressChanged
	INITIALIZED				= "c7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2" // Initializable.sol:Initialized
	CHARITY_PERCENTAGE_CHANGED		= "fe65b6d5007c66dadebc5095104ccd672c070f396dfdcfe1ee7f54201b9efa6d" // ISystemEvents.sol:CharityEthDonationAmountPercentageChanged
	PRIZE_PERCENTAGE_CHANGED	= "b5a05ec7911dd5450a7fa4ae54d50d9d83af6e256db8fc76c82edd7b659cf8bc" // ISystemEvents.sol:MainEthPrizeAmountPercentageChanged
	RAFFLE_PERCENTAGE_CHANGED = "bfcd8fb930a57c1598c9760db19c84ec766546c3c9a8565611df8302482bfb17" // ISystemEvents.sol:RaffleTotalEthPrizeAmountForBiddersPercentageChanged
	STAKE_PERCENTAGE_CHANGED = "9e44c04f534af356419a731b967f3d56bc748b1f3fdbda7a89f4e1a321ade934" // ISystemEvents.sol:CosmicSignatureNftStakingTotalEthRewardAmountPercentageChanged
	CHRONO_PERCENTAGE_CHANGED = "5581e31f5b8d4b3c45b8ab8bf67e3602ce588b423905eb6ad34bd6bc3c848699" // ISystemEvents.sol:ChronoWarriorEthPrizeAmountPercentageChanged
	NUM_RAFFLE_ETH_PRIZE_EVENTS_BIDDING_CHANGED = "4787028773c8f14bc3b4bc41f43a02329ae41105823287201a34782c530d35fd" // ISystemEvents.sol:NumRaffleEthPrizesForBiddersChanged
	NUM_RAFFLE_NFT_PRIZE_EVENTS_BIDDING_CHANGED = "85d8bf21006916836edd67a5afeed2e891cf876a6c9cd9babf7f42f4b007c24f" // ISystemEvents.sol:NumRaffleCosmicSignatureNftsForBiddersChanged
	NUM_RAFFLE_NFT_PRIZE_EVENTS_STAKING_RWALK_CHANGED = "3312247fcf207243294680d9103851c8709e19be3d46ee4b1bff5e12d1c5ca7b" // ISystemEvents.sol:NumRaffleCosmicSignatureNftsForRandomWalkNftStakersChanged
	SYSTEM_MODE_CHANGED		= "f24e774cdaabee9b8782266728e442b7f1fa6ae9204755c0da1541e99f04aa4c" // ISystemManagement.sol:SystemModeChanged
	RWALK_ADDRESS_CHANGED	= "dab38e33e6e11cbb8b085bba9d7426d5e5af01bcc46d2c5957645e4d09e8c49c" // ISystemEvents.sol:RandomWalkNftAddressChanged
	PRIZE_WALLET_ADDRESS_CHANGED	= "b4cecfe1346c94da27291cf5a02969d5fe0b5c36eca49b04fcd60841d28c5e13" // ISystemEvents.sol:PrizesWalletAddressChanged
	STAKING_WALLET_CST_ADDRESS_CHANGED  = "4da1815cd654922275d14d2335fd9a0dd0aa6a0d0ff87fb4cc872ebe9704596f" // ISystemEvents.sol:StakingWalletCosmicSignatureNftAddressChanged
	STAKING_WALLET_RWALK_ADDRESS_CHANGED  = "bf6e296f85d08cc1ab124aed644bf4b19e4a726a7aea53e3784ab1341738a040" // ISystemEvents.sol:StakingWalletRandomWalkNftAddressChanged
	MARKETING_ADDRESS_CHANGED = "4d03942c29c20d1bccfe551e9d148c917c5a44fb558a4fc60270d8f76fb75f54" // ISystemEvents.sol:MarketingWalletAddressChanged
	COSMIC_TOKEN_ADDRESS_CHANGED	= "9b3eda10f1724a2cf9f7dae4ac263c77908df4d00e92f1377b66fc8be37cd8c1" // ISystemEvents.sol:CosmicSignatureTokenAddressChanged
	COSMIC_SIGNATURE_ADDRESS_CHANGED	= "5bde6238168795ba4e77972a2bdaa5a465f7c9a5d05817f5e8d3fed2e5a4fa60" // ISystemEvents.sol:CosmicSignatureNftAddressChanged
	BUSINESS_LOGIC_ADDRESS_CHANGED	= "77ddb5e9e1495e15651bf87ccd8bbb7e637439fb260f0fda41b6ce4b3098aafd" // ISystemManagement.sol:BusinessLogicContractAddressChanged
	TIME_INCREASE_CHANGED	= "ed46e73b174ced51fb529cdc6c1b4d8abf49387e6d849b71648afb63c81d12cd" // ISystemEvents.sol:MainPrizeTimeIncrementIncreaseDivisorChanged
	TIMEOUT_CLAIMPRIZE_CHANGED	= "37a332914fac995349420c0419b4423a19dcb762017f691442a0782ce4bf417a" // ISystemEvents.sol:TimeoutDurationToClaimMainPrizeChanged
	TIMEOUT_TO_WITHDRAW_PRIZE = "8717bb199c6bc4a5dadb21547205f9ef8ec037dda246a5526d6a6471306ea52e" // IPrizesWallet.sol:TimeoutDurationToWithdrawPrizesChanged
	PRICE_INCREASE_CHANGED	= "deb71e1d117914859ddde484a9810084d9ee399077d5cd8dcbdf8770d33d8ae4" // ISystemEvents.sol:EthBidPriceIncreaseDivisorChanged
	MAIN_PRIZE_MICROSECOND_INCREASE= "07417920574ce0bdfe987af0575c8793cc73a29d7830760ad459d0e569b5b79b" // ISystemEvents.sol:MainPrizeTimeIncrementInMicroSecondsChanged
	INITIAL_SECONDS_UNTIL_PRIZE_CHANGED = "b5edd1f338b34c8f5dd3b1c5cc12f05653c495713c282bf588d34cf14fad0f89" // ISystemEvents.sol:InitialDurationUntilMainPrizeDivisorChanged
	ROUND_ACTIVATION_TIME_CHANGED = "9a2159c1f277ddd727551baedc6a6c4cba77cc5219c8563ee3b15fb67548d89b" // ISystemManagement.sol:RoundActivationTimeChanged
	ETH_DUTCH_AUCTION_DURATION_DIVISOR_CHANGED= "fdf6043c89a5f304289122dcc9f8bd78bb111b5d4f409e2fc2e6c141a1110b79" // ISystemEvents.sol:EthDutchAuctionDurationDivisorChanged
	CST_DUTCH_AUCTION_DURATION_DIVISOR_CHANGED= "c95d03f6c735a9e59c760fdb88e585aafe0a31b5c034fc7838155287ee32212f" // ISystemEvents.sol:CstDutchAuctionDurationDivisorChanged
	ETH_DUTCH_AUCTION_ENDING_PRICE_DIVISOR_CHANGED = "b6f6af60099e44041a78b3561ed029b98bf03fdb0efbbb2eb15e1f3d7d923037" // ISystemEvents.sol:EthDutchAuctionEndingBidPriceDivisorChanged
	MARKETING_REWARD_PAID	= "e2403640ba68fed3a2f88b7557551d1993f84b99bb10ff833f0cf8db0c5e0486" // IMarketingWallet.sol:RewardPaid
	MARKETING_REWARD_CHANGED = "2652e6657dd1ed89d7bdcb70f8827cc8449ac4536ebf604dbb2465cdad264020" // ISystemEvents.sol:MarketingWalletCstContributionAmountChanged
	CST_REWARD_FOR_BIDDING_CHANGED = "70ad04ce09c925ea466a5f603054f310bba5b7484bba77b382aade0bf93b55d0" // ISystemEvents.sol:CstRewardAmountForBiddingChanged
	STATIC_CST_REWARD		= "d95e7f967f9370c11deb15ffbb191b9f2e9795ab0738db5bc72bd2794978f32d" // ISystemEvents.sol:CstPrizeAmountChanged
	MAX_MESSAGE_LENGTH		= "157c413b0549fd4f45aab72b7828304fb2c45dad53de0f1128c5eabf3aaabaf8" // ISystemEvents.sol:BidMessageLengthMaxLimitChanged
	TOKEN_SCRIPT_URL		= "27e2bd70f498920ee0fd7d8204ae8845b75dc81330e3acafa32946be3503730c" // ICosmicSignatureNft.sol:NftGenerationScriptUrlChanged
	BASE_URI				= "bdfd815215fcee5bb949c941ab489c7ead076a7c8acd3527cd1b50f613ac67e6" // ICosmicSignatureNft.sol:BaseUriChanged
	OWNERSHIP_TRANSFERRED	= "8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0" // Ownable.sol:OwnershipTransferred
	STARTING_CST_MIN_LIM	= "4e8c80fe79d13b8663de9f6981925ae24c498cc07b0ebd4f4789fa78ca17caff" // ISystemEvents.sol:CstDutchAuctionBeginningBidPriceMinLimitChanged
	FUNDS_TO_CHARITY		= "1222634ba80f397fa40371bab63974e6f97da8200777ec79e731c69bb6a2351d" // ICosmicSignatureEvents.sol:FundsTransferredToCharity
	DELAY_DURATION_ROUND	= "b0868a729f47ae3829aaafe3ca2975d3db2148553c854112f598be6d91ef0d28" // ISystemManagement.sol:DelayDurationBeforeRoundActivationChanged

)
var (
	eclient 				*ethclient.Client
	rpcclient 				*rpc.Client

	// CosmicGame events:
	evt_prize_claim_event,_ = hex.DecodeString(PRIZE_CLAIM_EVENT)
	evt_bid_event,_			= hex.DecodeString(BID_EVENT)
	evt_eth_donated_event,_	= hex.DecodeString(ETH_DONATED_EVENT)
	evt_eth_donated_wi_event,_= hex.DecodeString(ETH_DONATED_WI_EVENT)
	evt_nft_donation_event,_= hex.DecodeString(NFT_ETH_DONATED_EVENT)
	evt_erc20_donated,_		= hex.DecodeString(ERC20_DONATED)
	evt_raffle_nft_prize_event,_	= hex.DecodeString(RAFFLE_NFT_PRIZE_EVENT)
	evt_raffle_eth_prize_event,_	= hex.DecodeString(RAFFLE_ETH_PRIZE_EVENT)
	evt_endurance_prize_event,_	= hex.DecodeString(ENDURANCE_PRIZE_EVENT)
	evt_lastcst_bidder_prize_event,_	= hex.DecodeString(LASTCST_BIDDER_PRIZE_EVENT)
	evt_chrono_warrior_prize_event,_	= hex.DecodeString(CHRONO_WARRIOR_PRIZE_EVENT)
	evt_donated_token_claimed,_	= hex.DecodeString(DONATED_TOKEN_CLAIMED)
	evt_donated_nft_claimed,_= hex.DecodeString(DONATED_NFT_CLAIMED)
	evt_charity_percentage_changed,_= hex.DecodeString(CHARITY_PERCENTAGE_CHANGED)
	evt_prize_percentage_changed,_ = hex.DecodeString(PRIZE_PERCENTAGE_CHANGED)
	evt_raffle_percentage_changed,_ = hex.DecodeString(RAFFLE_PERCENTAGE_CHANGED)
	evt_staking_percentage_changed,_ = hex.DecodeString(STAKE_PERCENTAGE_CHANGED)
	evt_chrono_percentage_changed,_ = hex.DecodeString(CHRONO_PERCENTAGE_CHANGED)
	evt_num_raffle_eth_winners_bidding_changed,_ = hex.DecodeString(NUM_RAFFLE_ETH_PRIZE_EVENTS_BIDDING_CHANGED)
	evt_num_raffle_nft_winners_bidding_changed,_ = hex.DecodeString(NUM_RAFFLE_NFT_PRIZE_EVENTS_BIDDING_CHANGED)
	evt_num_raffle_nft_winners_staking_rwalk_changed,_ = hex.DecodeString(NUM_RAFFLE_NFT_PRIZE_EVENTS_STAKING_RWALK_CHANGED)
	evt_charity_wallet_changed,_ = hex.DecodeString(CHARITY_WALLET_CHANGED)
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
	evt_activation_time_changed,_	= hex.DecodeString(ROUND_ACTIVATION_TIME_CHANGED)
	evt_cst_dutch_auction_duration_divisor_changed,_ = hex.DecodeString(CST_DUTCH_AUCTION_DURATION_DIVISOR_CHANGED) // previously	evt_round_start_auction_length_changed
	evt_eth_dutch_auction_duration_divisor_changed,_ = hex.DecodeString(ETH_DUTCH_AUCTION_DURATION_DIVISOR_CHANGED)
	evt_eth_dutch_auction_ending_bidprice_divisor,_ = hex.DecodeString(ETH_DUTCH_AUCTION_ENDING_PRICE_DIVISOR_CHANGED)
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
	evt_static_cst_reward,_				= hex.DecodeString(STATIC_CST_REWARD)
	evt_funds2charity,_				= hex.DecodeString(FUNDS_TO_CHARITY)
	evt_delay_duration_round,_		= hex.DecodeString(DELAY_DURATION_ROUND)
	evt_first_bid_event,_				= hex.DecodeString(FIRST_BID_EVENT)

	// CharityWallet events
	evt_donation_received_event,_=hex.DecodeString(DONATION_RECEIVED_EVENT)
	evt_donation_sent_event,_= hex.DecodeString(DONATION_SENT_EVENT)
	evt_charity_receiver_changed,_ = hex.DecodeString(CHARITY_RECEIVER_CHANGED)

	// CosmicSignature events
	evt_token_name_event,_	= hex.DecodeString(TOKEN_NAME_EVENT)
	evt_mint_event,_		= hex.DecodeString(MINT_EVENT)

	// PrizeWallet events
	evt_eth_prize_deposit,_		= hex.DecodeString(ETH_PRIZE_DEPOSIT_EVENT)
	evt_eth_prize_withdrawal,_	= hex.DecodeString(ETH_PRIZE_WITHDRAWAL_EVENT)

	// ERC20 events
	evt_transfer,_			= hex.DecodeString(TRANSFER_EVT)

	// StakingWallet events
	evt_cst_nft_staked,_		= hex.DecodeString(CST_NFT_STAKED_EVENT)
	evt_rwalk_nft_staked,_		= hex.DecodeString(RWALK_NFT_STAKED_EVENT)
	evt_nft_unstaked_rwalk,_= hex.DecodeString(NFT_UNSTAKED_RWALK)
	evt_nft_unstaked_cst,_	= hex.DecodeString(NFT_UNSTAKED_CST)
	evt_staking_eth_deposit,_		= hex.DecodeString(STAKING_ETH_DEPOSIT_EVENT)

	// MarketingWallet events
	evt_marketing_reward_paid,_		= hex.DecodeString(MARKETING_REWARD_PAID)
	evt_treasurer_changed,_		= hex.DecodeString(TREASURER_CHANGED)

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

	cosmic_game_addr		ethcommon.Address
	cosmic_signature_addr	ethcommon.Address
	cosmic_token_addr		ethcommon.Address
	cosmic_dao_addr			ethcommon.Address
	charity_wallet_addr		ethcommon.Address
	prizes_wallet_addr		ethcommon.Address
	staking_wallet_cst_addr		ethcommon.Address
	staking_wallet_rwalk_addr		ethcommon.Address
	marketing_wallet_addr	ethcommon.Address
	cosmic_sig_aid			int64
	cosmic_tok_aid			int64

	cg_contracts			CosmicGameContractAddrs
	storagew				SQLStorageWrapper
	RPC_URL					 = os.Getenv("RPC_URL")
	Error					*log.Logger
	Info					*log.Logger
)

// getContractAddresses returns all contract addresses for FilterLogs
func getContractAddresses() []ethcommon.Address {
	return []ethcommon.Address{
		cosmic_game_addr,
		cosmic_signature_addr,
		cosmic_token_addr,
		charity_wallet_addr,
		prizes_wallet_addr,
		staking_wallet_cst_addr,
		staking_wallet_rwalk_addr,
		marketing_wallet_addr,
	}
}

// process_events_filterlog uses FilterLogs to get events directly from blockchain
func process_events_filterlog(exit_chan chan bool) {
	// Create ETL context for common operations
	ctx := &etlcommon.ETLContext{
		Storage:   storagew.S,
		EthClient: eclient,
		RpcClient: rpcclient,
		Info:      Info,
		Error:     Error,
	}

	// Adaptive batch sizing: start large, reduce if we get events
	var batchSize uint64 = 100000      // Start with 100k blocks
	var minBatchSize uint64 = 1000     // Minimum when processing events
	var maxBatchSize uint64 = 1000000  // Maximum when scanning empty ranges
	contracts := getContractAddresses()

	for {
		select {
			case exit_flag := <-exit_chan:
				if exit_flag {
					Info.Println("Exiting by user request.")
					os.Exit(0)
				}
			default:
		}

		// Get last processed block from status
		status := storagew.Get_cosmic_game_processing_status()
		lastProcessedBlock := status.LastBlockNum
		if lastProcessedBlock == 0 {
			// If no blocks processed yet, start from the block where contracts were deployed
			lastProcessedBlock, _ = storagew.S.Get_last_block_num()
		}

		// Get current block from chain
		currentBlock, err := etlcommon.GetCurrentBlockNumber(eclient)
		if err != nil {
			Error.Printf("Failed to get current block number: %v", err)
			time.Sleep(5 * time.Second)
			continue
		}

		// Calculate block range to process
		fromBlock := uint64(lastProcessedBlock + 1)
		toBlock := fromBlock + batchSize - 1
		if toBlock > currentBlock {
			toBlock = currentBlock
		}

		if fromBlock > currentBlock {
			// Already caught up, wait for new blocks
			time.Sleep(2 * time.Second)
			batchSize = minBatchSize // Reset to small batch for real-time
			continue
		}

		Info.Printf("Fetching events from block %d to %d (batch size: %d)\n", fromBlock, toBlock, batchSize)

		// Fetch events using FilterLogs
		logs, err := etlcommon.FetchEvents(eclient, fromBlock, toBlock, contracts)
		if err != nil {
			Error.Printf("FetchEvents failed: %v", err)
			// Reduce batch size on error (might be too large)
			batchSize = batchSize / 2
			if batchSize < minBatchSize {
				batchSize = minBatchSize
			}
			time.Sleep(5 * time.Second)
			continue
		}

		Info.Printf("Received %d events\n", len(logs))

		// Process each event
		var processingFailed bool
		var lastSuccessfulBlock uint64
		for _, log := range logs {
			// Ensure block exists with correct hash (chain split detection)
			_, err := etlcommon.EnsureBlockExists(ctx, int64(log.BlockNumber), log.BlockHash.Hex())
			if err != nil {
				Error.Printf("EnsureBlockExists failed for block %d: %v", log.BlockNumber, err)
				processingFailed = true
				time.Sleep(5 * time.Second)
				break
			}

			// Ensure transaction exists
			txId, _, err := etlcommon.EnsureTransactionExists(ctx, log.TxHash, int64(log.BlockNumber))
			if err != nil {
				Error.Printf("EnsureTransactionExists failed for tx %s: %v", log.TxHash.Hex(), err)
				processingFailed = true
				time.Sleep(5 * time.Second)
				break
			}

			// Insert event log
			evtId, err := etlcommon.InsertEventLog(ctx, log, txId)
			if err != nil {
				Error.Printf("InsertEventLog failed: %v", err)
				processingFailed = true
				time.Sleep(5 * time.Second)
				break
			}

			// Process the event using existing logic
			err = process_single_event(evtId)
			if err != nil {
				Error.Printf("process_single_event failed for evt %d: %v", evtId, err)
				// Continue processing other events
			}
			
			// Track last successfully processed block
			lastSuccessfulBlock = log.BlockNumber
		}

		// Only update status if processing succeeded
		if !processingFailed {
			status.LastBlockNum = int64(toBlock)
			storagew.Update_cosmic_game_process_status(&status)
		} else if lastSuccessfulBlock > 0 {
			// Update to last successfully processed block
			status.LastBlockNum = int64(lastSuccessfulBlock)
			storagew.Update_cosmic_game_process_status(&status)
		}
		// If processingFailed and lastSuccessfulBlock==0, don't update - will retry same batch

		// Adaptive batch sizing
		if len(logs) == 0 {
			// No events - increase batch size for faster scanning
			batchSize = batchSize * 2
			if batchSize > maxBatchSize {
				batchSize = maxBatchSize
			}
		} else {
			// Found events - use smaller batch for granularity
			batchSize = minBatchSize
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
	
	// Insert all contract addresses into address table (for fresh database)
	storagew.S.Lookup_or_create_address(cg_contracts.CosmicGameAddr, 0, 0)
	storagew.S.Lookup_or_create_address(cg_contracts.CosmicSignatureAddr, 0, 0)
	storagew.S.Lookup_or_create_address(cg_contracts.CosmicTokenAddr, 0, 0)
	storagew.S.Lookup_or_create_address(cg_contracts.CosmicDaoAddr, 0, 0)
	storagew.S.Lookup_or_create_address(cg_contracts.CharityWalletAddr, 0, 0)
	storagew.S.Lookup_or_create_address(cg_contracts.PrizesWalletAddr, 0, 0)
	storagew.S.Lookup_or_create_address(cg_contracts.RandomWalkAddr, 0, 0)
	storagew.S.Lookup_or_create_address(cg_contracts.StakingWalletCSTAddr, 0, 0)
	storagew.S.Lookup_or_create_address(cg_contracts.StakingWalletRWalkAddr, 0, 0)
	storagew.S.Lookup_or_create_address(cg_contracts.MarketingWalletAddr, 0, 0)
	storagew.S.Lookup_or_create_address(cg_contracts.ImplementationAddr, 0, 0)
	Info.Printf("All contract addresses registered in address table\n")
	
	// Now lookup the address IDs (they are guaranteed to exist now)
	cosmic_sig_aid,err  = storagew.S.Nonfatal_lookup_address_id(cg_contracts.CosmicSignatureAddr)
	if err != nil {
		fmt.Printf("Lookup of CosmicSignatureAddr failed: %v",err)
		os.Exit(1)
	}
	cosmic_tok_aid,err  = storagew.S.Nonfatal_lookup_address_id(cg_contracts.CosmicTokenAddr)
	if err != nil {
		fmt.Printf("Lookup of CosmicTokenAddr failed: %v",err)
		os.Exit(1)
	}
	cosmic_game_addr = ethcommon.HexToAddress(cg_contracts.CosmicGameAddr)
	cosmic_signature_addr = ethcommon.HexToAddress(cg_contracts.CosmicSignatureAddr)
	cosmic_token_addr = ethcommon.HexToAddress(cg_contracts.CosmicTokenAddr)
	cosmic_dao_addr = ethcommon.HexToAddress(cg_contracts.CosmicDaoAddr)
	charity_wallet_addr = ethcommon.HexToAddress(cg_contracts.CharityWalletAddr)
	prizes_wallet_addr = ethcommon.HexToAddress(cg_contracts.PrizesWalletAddr)
	staking_wallet_cst_addr = ethcommon.HexToAddress(cg_contracts.StakingWalletCSTAddr)
	staking_wallet_rwalk_addr = ethcommon.HexToAddress(cg_contracts.StakingWalletRWalkAddr)
	marketing_wallet_addr = ethcommon.HexToAddress(cg_contracts.MarketingWalletAddr)

	c := make(chan os.Signal, 1)
	exit_chan := make(chan bool)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM, syscall.SIGHUP)
	go func() {
		<-c
		Info.Printf("Got signal, will exit after current batch is processed." +
					" To interrupt abruptly send SIGKILL (9) to the kernel.\n")
		exit_chan <- true
	}()


	// Use new FilterLogs-based event processing
	process_events_filterlog(exit_chan)
}
