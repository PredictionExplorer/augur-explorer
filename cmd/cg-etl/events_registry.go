// Event-signature registry for the CosmicGame ETL: topic-hash constants, decoded
// signature variables, the inspected-events list and the event dispatch/lookup logic.
package main

import (
	"bytes"
	"encoding/hex"
	"fmt"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/rlp"

	. "github.com/PredictionExplorer/augur-explorer/internal/primitives"
)

const (
	PRIZE_CLAIM_EVENT          = "8c551ec2b6f186753e27f1cf46f84b57f4f83f721e8c1e6170ae512845ced591" // ICosmicSignatureGame.sol:MainPrizeClaimed
	BID_EVENT                  = "bcb004d688d0951e50c218ded0d0d574bde915630e29b92987b1f2eab9556549" // IBidding.sol:BidPlaced (legacy)
	BID_EVENT_V2               = "1d1f406c89e99504a7222aaba50bf96b9f91f522ebc6ea825255145919e102ec" // IBiddingV2.sol:BidPlaced; both topics handled in parallel
	ETH_DONATED_EVENT          = "e32cacf203d00685e2b4d8b0a90e7cd8f3f8a208fdf116f4bb36abe08b7d548e" // IEthDonations.sol:EthDonated
	ETH_DONATED_WI_EVENT       = "a08049565b10d44a06dca9bf05685b39bc370352043c5a003e8d35d45ebdc53f" // IEthDonations.sol:EthDonatedWithInfo
	DONATION_RECEIVED_EVENT    = "264f630d9efa0d07053a31163641d9fcc0adafc9d9e76f1c37c2ce3a558d2c52" // ICharityWallet.sol:DonationReceived
	DONATION_SENT_EVENT        = "1222634ba80f397fa40371bab63974e6f97da8200777ec79e731c69bb6a2351d" // ICharityWallet.sol:DonationSent
	CHARITY_RECEIVER_CHANGED   = "1c7efd98583e33a9cc6adff48a97abdaaf43e5c9e918d5ec3e75e93e1dafcf6c" // ICharityWallet.sol:CharityAddressChanged
	CHARITY_WALLET_CHANGED     = "1c7efd98583e33a9cc6adff48a97abdaaf43e5c9e918d5ec3e75e93e1dafcf6c" // ISystemEvents.sol:CharityAddressChanged
	TOKEN_NAME_EVENT           = "a14cfb0fe69c0b55eaaa4d9400bdba2a72e1860cade89c2a8a055e6cfde8936d" // ICosmicSignatureNft.sol:NftNameChanged
	MINT_EVENT                 = "c2115f21464937bfdcd1560f96f0e20b70e88accbdcd1069084c80c8797ef106" // ICosmicSignatureNft.sol:NftMinted
	NFT_ETH_DONATED_EVENT      = "b12e72bab0c2a8fe98175a3c02792645ebdf65d37cf3983517e6d1c5ab842e23" // IPrizesWallet.sol:NftDonated
	ERC20_DONATED              = "3f94f6171c96ab491d8c268f80da5e460b148878065711a0bb099bbc297615af" // IPrizesWallet.sol:TokenDonated
	DONATED_TOKEN_CLAIMED      = "af1adae2e1e983ca738335dc2e37194114142793f394ac934a45ea632a8a5bd0" // IPrizesWallet.sol:DonatedTokenClaimed
	DONATED_NFT_CLAIMED        = "03c2b6e01c9ca39e4073132f279b61b987d41a82b08cea3dd4a2fabc37067ac3" // IPrizesWallet.sol:DonatedNftClaimed
	ETH_PRIZE_DEPOSIT_EVENT    = "8e369548f9ef214c7c1823c098a23763321fd761ac5cf78958e1db1b3648e7a2" // IPrizesWallet.sol:EthReceived
	ETH_PRIZE_WITHDRAWAL_EVENT = "172b54ba94575dba1c8dda35e4a6f6a0f761efe5c5416227b34c7c7632a673d4" // IPrizesWallet.sol:EthWithdrawn
	RAFFLE_ETH_PRIZE_EVENT     = "9c62e2cb8cbd10bf3b8a5760977d719fc3475fa67cb7ef9d2e1701f275e885c4" // ICosmicSignatureGame.sol:RaffleWinnerBidderEthPrizeAllocated
	RAFFLE_NFT_PRIZE_EVENT     = "27c21fe4cea1a3367aa491829dd4dd824296c00910626150464cba8ea5ebb3f4" // ICosmicSignatureGame.sol:RaffleWinnerPrizePaid
	ENDURANCE_PRIZE_EVENT      = "838ec9dd2530548892bff113f5ffb0138d2efc63c7f59bb6571e8c923b631260" // ICosmicSignatureGame.sol:EnduranceChampionPrizePaid
	LASTCST_BIDDER_PRIZE_EVENT = "3901b6430c99dc290ee88ff84c4de6091ad7eac335b58e92ef5cbb0793abf4f6" // ICosmicSignatureGame.sol:LastCstBidderPrizePaid
	CHRONO_WARRIOR_PRIZE_EVENT = "aa858ae20a26d00a9ea528972d537e68a51a0744226d9ea1fc9b22492dc282a5" // ICosmicSignatureGame.sol:ChronoWarriorPrizePaid
	TRANSFER_EVT               = "ddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef" // IERC721.sol:Transfer
	CST_NFT_STAKED_EVENT       = "e09cd972bcd125457d8f8a684b2b67ec513fbb7f770001bbebd7c22b41ad9da8" // IStakingWalletNftBase.sol:NftStaked
	RWALK_NFT_STAKED_EVENT     = "62773741191803b5cec48480156933243e422a1fb1ea9967dab3ee30df2da95f" // IStakingWalletNftBase.sol:NftStaked
	NFT_UNSTAKED_RWALK         = "08e7047cd3ef25f4a72589ed1d73eb0af1ca8a5957f9539e126dff45870979cc" // IStakingWalletNftBase.sol:NftUnstaked
	NFT_UNSTAKED_CST           = "ec478a78e4e3931ff728a54eeb6875304c891fa5fa253337b60d37fdc5e1feaf" // IStakingWalletNftBase.sol:NftUnstaked
	STAKING_ETH_DEPOSIT_EVENT  = "26726e1a40953e6b9e06e1c1a1f53422299c7e00d2cc5d986bd4723796ca4588" // IStakingWalletCosmicSignatureNft.sol:EthDepositReceived
	CLAIM_REWARD_EVENT         = "dde81df5caa033e783e2a39d93e3a8718a7dc27ba95a4757f5433a01f794ec36" // IStakingWalletNftBase.sol:RewardClaimed
	FUND_TRANSFER_ERR          = "154fb6c686c977314af35b730a16b83facda5265b2abec7237c730f63e42215a" // ICosmicSignatureErrors.sol:FundTransferFailed
	ERC20_TRANSFER_ERR         = "f7fce645f12ae266a329c431e96ebea892316a1415809056621ffeea04efd4ab" // ICosmicSignatureErrors.sol:ERC20TransferFailed
	FIRST_BID_EVENT            = "028a52641badd593b7f30072734c0b97e449213f55b5c3663756427340accd3c" // ICosmicSignatureGame.sol:FirstBidPlacedInRound

	/// Admin events
	PROXY_UPGRADED                                    = "bc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b" // IERC1967.sol:Upgraded
	ADMIN_CHANGED                                     = "7e644d79422f17c01e4894b5f4f588d331ebfa28653d42ae832dc59e38c9798f" // IERC1967.sol:AdminChanged
	TREASURER_CHANGED                                 = "df73fc12cc071a4834f7ba0e7c6cfe7d23e98866e191ec9e86e6e61614d9e50e" // IMarketingWallet.sol:TreasurerAddressChanged
	INITIALIZED                                       = "c7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2" // Initializable.sol:Initialized
	CHARITY_PERCENTAGE_CHANGED                        = "fe65b6d5007c66dadebc5095104ccd672c070f396dfdcfe1ee7f54201b9efa6d" // ISystemEvents.sol:CharityEthDonationAmountPercentageChanged
	PRIZE_PERCENTAGE_CHANGED                          = "b5a05ec7911dd5450a7fa4ae54d50d9d83af6e256db8fc76c82edd7b659cf8bc" // ISystemEvents.sol:MainEthPrizeAmountPercentageChanged
	RAFFLE_PERCENTAGE_CHANGED                         = "bfcd8fb930a57c1598c9760db19c84ec766546c3c9a8565611df8302482bfb17" // ISystemEvents.sol:RaffleTotalEthPrizeAmountForBiddersPercentageChanged
	STAKE_PERCENTAGE_CHANGED                          = "9e44c04f534af356419a731b967f3d56bc748b1f3fdbda7a89f4e1a321ade934" // ISystemEvents.sol:CosmicSignatureNftStakingTotalEthRewardAmountPercentageChanged
	CHRONO_PERCENTAGE_CHANGED                         = "5581e31f5b8d4b3c45b8ab8bf67e3602ce588b423905eb6ad34bd6bc3c848699" // ISystemEvents.sol:ChronoWarriorEthPrizeAmountPercentageChanged
	NUM_RAFFLE_ETH_PRIZE_EVENTS_BIDDING_CHANGED       = "4787028773c8f14bc3b4bc41f43a02329ae41105823287201a34782c530d35fd" // ISystemEvents.sol:NumRaffleEthPrizesForBiddersChanged
	NUM_RAFFLE_NFT_PRIZE_EVENTS_BIDDING_CHANGED       = "85d8bf21006916836edd67a5afeed2e891cf876a6c9cd9babf7f42f4b007c24f" // ISystemEvents.sol:NumRaffleCosmicSignatureNftsForBiddersChanged
	NUM_RAFFLE_NFT_PRIZE_EVENTS_STAKING_RWALK_CHANGED = "3312247fcf207243294680d9103851c8709e19be3d46ee4b1bff5e12d1c5ca7b" // ISystemEvents.sol:NumRaffleCosmicSignatureNftsForRandomWalkNftStakersChanged
	SYSTEM_MODE_CHANGED                               = "f24e774cdaabee9b8782266728e442b7f1fa6ae9204755c0da1541e99f04aa4c" // ISystemManagement.sol:SystemModeChanged
	RWALK_ADDRESS_CHANGED                             = "dab38e33e6e11cbb8b085bba9d7426d5e5af01bcc46d2c5957645e4d09e8c49c" // ISystemEvents.sol:RandomWalkNftAddressChanged
	PRIZE_WALLET_ADDRESS_CHANGED                      = "b4cecfe1346c94da27291cf5a02969d5fe0b5c36eca49b04fcd60841d28c5e13" // ISystemEvents.sol:PrizesWalletAddressChanged
	STAKING_WALLET_CST_ADDRESS_CHANGED                = "4da1815cd654922275d14d2335fd9a0dd0aa6a0d0ff87fb4cc872ebe9704596f" // ISystemEvents.sol:StakingWalletCosmicSignatureNftAddressChanged
	STAKING_WALLET_RWALK_ADDRESS_CHANGED              = "bf6e296f85d08cc1ab124aed644bf4b19e4a726a7aea53e3784ab1341738a040" // ISystemEvents.sol:StakingWalletRandomWalkNftAddressChanged
	MARKETING_ADDRESS_CHANGED                         = "4d03942c29c20d1bccfe551e9d148c917c5a44fb558a4fc60270d8f76fb75f54" // ISystemEvents.sol:MarketingWalletAddressChanged
	COSMIC_TOKEN_ADDRESS_CHANGED                      = "9b3eda10f1724a2cf9f7dae4ac263c77908df4d00e92f1377b66fc8be37cd8c1" // ISystemEvents.sol:CosmicSignatureTokenAddressChanged
	COSMIC_SIGNATURE_ADDRESS_CHANGED                  = "5bde6238168795ba4e77972a2bdaa5a465f7c9a5d05817f5e8d3fed2e5a4fa60" // ISystemEvents.sol:CosmicSignatureNftAddressChanged
	BUSINESS_LOGIC_ADDRESS_CHANGED                    = "77ddb5e9e1495e15651bf87ccd8bbb7e637439fb260f0fda41b6ce4b3098aafd" // ISystemManagement.sol:BusinessLogicContractAddressChanged
	TIME_INCREASE_CHANGED                             = "ed46e73b174ced51fb529cdc6c1b4d8abf49387e6d849b71648afb63c81d12cd" // ISystemEvents.sol:MainPrizeTimeIncrementIncreaseDivisorChanged
	TIMEOUT_CLAIMPRIZE_CHANGED                        = "37a332914fac995349420c0419b4423a19dcb762017f691442a0782ce4bf417a" // ISystemEvents.sol:TimeoutDurationToClaimMainPrizeChanged
	TIMEOUT_TO_WITHDRAW_PRIZE                         = "8717bb199c6bc4a5dadb21547205f9ef8ec037dda246a5526d6a6471306ea52e" // IPrizesWallet.sol:TimeoutDurationToWithdrawPrizesChanged
	PRICE_INCREASE_CHANGED                            = "deb71e1d117914859ddde484a9810084d9ee399077d5cd8dcbdf8770d33d8ae4" // ISystemEvents.sol:EthBidPriceIncreaseDivisorChanged
	MAIN_PRIZE_MICROSECOND_INCREASE                   = "07417920574ce0bdfe987af0575c8793cc73a29d7830760ad459d0e569b5b79b" // ISystemEvents.sol:MainPrizeTimeIncrementInMicroSecondsChanged
	INITIAL_SECONDS_UNTIL_PRIZE_CHANGED               = "b5edd1f338b34c8f5dd3b1c5cc12f05653c495713c282bf588d34cf14fad0f89" // ISystemEvents.sol:InitialDurationUntilMainPrizeDivisorChanged
	ROUND_ACTIVATION_TIME_CHANGED                     = "9a2159c1f277ddd727551baedc6a6c4cba77cc5219c8563ee3b15fb67548d89b" // ISystemManagement.sol:RoundActivationTimeChanged
	ETH_DUTCH_AUCTION_DURATION_DIVISOR_CHANGED        = "fdf6043c89a5f304289122dcc9f8bd78bb111b5d4f409e2fc2e6c141a1110b79" // ISystemEvents.sol:EthDutchAuctionDurationDivisorChanged
	CST_DUTCH_AUCTION_DURATION_DIVISOR_CHANGED        = "c95d03f6c735a9e59c760fdb88e585aafe0a31b5c034fc7838155287ee32212f" // ISystemEvents.sol:CstDutchAuctionDurationDivisorChanged
	ETH_DUTCH_AUCTION_ENDING_PRICE_DIVISOR_CHANGED    = "b6f6af60099e44041a78b3561ed029b98bf03fdb0efbbb2eb15e1f3d7d923037" // ISystemEvents.sol:EthDutchAuctionEndingBidPriceDivisorChanged
	MARKETING_REWARD_PAID                             = "e2403640ba68fed3a2f88b7557551d1993f84b99bb10ff833f0cf8db0c5e0486" // IMarketingWallet.sol:RewardPaid
	MARKETING_REWARD_CHANGED                          = "2652e6657dd1ed89d7bdcb70f8827cc8449ac4536ebf604dbb2465cdad264020" // ISystemEvents.sol:MarketingWalletCstContributionAmountChanged
	CST_REWARD_FOR_BIDDING_CHANGED                    = "70ad04ce09c925ea466a5f603054f310bba5b7484bba77b382aade0bf93b55d0" // ISystemEvents.sol:CstRewardAmountForBiddingChanged
	BID_CST_REWARD_AMOUNT_CHANGED                     = "96978b83addd498dff54ab50bf4ed5b62e543d07c7935099eafe180248efe4b4" // ISystemEvents.sol:BidCstRewardAmountChanged
	BID_CST_REWARD_AMOUNT_MULTIPLIER_CHANGED          = "40b9c59af8c486ccf8c7cc73df5a51e7cc29747ea7d39f99632ecaf9caa2ed1f" // ISystemEventsV2.sol:BidCstRewardAmountMultiplierChanged
	CST_DUTCH_AUCTION_DURATION_CHANGED                = "4abea08c196329c357e3175d011af39a8625be99ef0ba5a0f3547a95534fedb7" // ISystemEventsV2.sol:CstDutchAuctionDurationChanged
	CST_DUTCH_AUCTION_DURATION_CHANGE_DIVISOR_CHANGED = "acbc6b6929088e4b2d043625fa7248e00fb6658a425eb9d9dc8c37b18c7f3e6f" // ISystemEventsV2.sol:CstDutchAuctionDurationChangeDivisorChanged
	STATIC_CST_REWARD                                 = "d95e7f967f9370c11deb15ffbb191b9f2e9795ab0738db5bc72bd2794978f32d" // ISystemEvents.sol:CstPrizeAmountChanged
	MAX_MESSAGE_LENGTH                                = "157c413b0549fd4f45aab72b7828304fb2c45dad53de0f1128c5eabf3aaabaf8" // ISystemEvents.sol:BidMessageLengthMaxLimitChanged
	TOKEN_SCRIPT_URL                                  = "27e2bd70f498920ee0fd7d8204ae8845b75dc81330e3acafa32946be3503730c" // ICosmicSignatureNft.sol:NftGenerationScriptUrlChanged
	BASE_URI                                          = "bdfd815215fcee5bb949c941ab489c7ead076a7c8acd3527cd1b50f613ac67e6" // ICosmicSignatureNft.sol:BaseUriChanged
	OWNERSHIP_TRANSFERRED                             = "8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0" // Ownable.sol:OwnershipTransferred
	STARTING_CST_MIN_LIM                              = "4e8c80fe79d13b8663de9f6981925ae24c498cc07b0ebd4f4789fa78ca17caff" // ISystemEvents.sol:CstDutchAuctionBeginningBidPriceMinLimitChanged
	FUNDS_TO_CHARITY                                  = "1222634ba80f397fa40371bab63974e6f97da8200777ec79e731c69bb6a2351d" // ICosmicSignatureEvents.sol:FundsTransferredToCharity
	DELAY_DURATION_ROUND                              = "b0868a729f47ae3829aaafe3ca2975d3db2148553c854112f598be6d91ef0d28" // ISystemManagement.sol:DelayDurationBeforeRoundActivationChanged

)

var (
	// CosmicGame events:
	evt_prize_claim_event, _                                 = hex.DecodeString(PRIZE_CLAIM_EVENT)
	evt_bid_event, _                                         = hex.DecodeString(BID_EVENT)
	evt_bid_event_v2, _                                      = hex.DecodeString(BID_EVENT_V2)
	evt_eth_donated_event, _                                 = hex.DecodeString(ETH_DONATED_EVENT)
	evt_eth_donated_wi_event, _                              = hex.DecodeString(ETH_DONATED_WI_EVENT)
	evt_nft_donation_event, _                                = hex.DecodeString(NFT_ETH_DONATED_EVENT)
	evt_erc20_donated, _                                     = hex.DecodeString(ERC20_DONATED)
	evt_raffle_nft_prize_event, _                            = hex.DecodeString(RAFFLE_NFT_PRIZE_EVENT)
	evt_raffle_eth_prize_event, _                            = hex.DecodeString(RAFFLE_ETH_PRIZE_EVENT)
	evt_endurance_prize_event, _                             = hex.DecodeString(ENDURANCE_PRIZE_EVENT)
	evt_lastcst_bidder_prize_event, _                        = hex.DecodeString(LASTCST_BIDDER_PRIZE_EVENT)
	evt_chrono_warrior_prize_event, _                        = hex.DecodeString(CHRONO_WARRIOR_PRIZE_EVENT)
	evt_donated_token_claimed, _                             = hex.DecodeString(DONATED_TOKEN_CLAIMED)
	evt_donated_nft_claimed, _                               = hex.DecodeString(DONATED_NFT_CLAIMED)
	evt_charity_percentage_changed, _                        = hex.DecodeString(CHARITY_PERCENTAGE_CHANGED)
	evt_prize_percentage_changed, _                          = hex.DecodeString(PRIZE_PERCENTAGE_CHANGED)
	evt_raffle_percentage_changed, _                         = hex.DecodeString(RAFFLE_PERCENTAGE_CHANGED)
	evt_staking_percentage_changed, _                        = hex.DecodeString(STAKE_PERCENTAGE_CHANGED)
	evt_chrono_percentage_changed, _                         = hex.DecodeString(CHRONO_PERCENTAGE_CHANGED)
	evt_num_raffle_eth_winners_bidding_changed, _            = hex.DecodeString(NUM_RAFFLE_ETH_PRIZE_EVENTS_BIDDING_CHANGED)
	evt_num_raffle_nft_winners_bidding_changed, _            = hex.DecodeString(NUM_RAFFLE_NFT_PRIZE_EVENTS_BIDDING_CHANGED)
	evt_num_raffle_nft_winners_staking_rwalk_changed, _      = hex.DecodeString(NUM_RAFFLE_NFT_PRIZE_EVENTS_STAKING_RWALK_CHANGED)
	evt_charity_wallet_changed, _                            = hex.DecodeString(CHARITY_WALLET_CHANGED)
	evt_rwalk_address_changed, _                             = hex.DecodeString(RWALK_ADDRESS_CHANGED)
	evt_prizes_wallet_address_changed, _                     = hex.DecodeString(PRIZE_WALLET_ADDRESS_CHANGED)
	evt_staking_wallet_cst_address_changed, _                = hex.DecodeString(STAKING_WALLET_CST_ADDRESS_CHANGED)
	evt_staking_wallet_rwalk_address_changed, _              = hex.DecodeString(STAKING_WALLET_RWALK_ADDRESS_CHANGED)
	evt_marketing_address_changed, _                         = hex.DecodeString(MARKETING_ADDRESS_CHANGED)
	evt_costok_address_changed, _                            = hex.DecodeString(COSMIC_TOKEN_ADDRESS_CHANGED)
	evt_cossig_address_changed, _                            = hex.DecodeString(COSMIC_SIGNATURE_ADDRESS_CHANGED)
	evt_time_increase_changed, _                             = hex.DecodeString(TIME_INCREASE_CHANGED)
	evt_timeout_claimprize_changed, _                        = hex.DecodeString(TIMEOUT_CLAIMPRIZE_CHANGED)
	evt_timeout_to_withdraw_prize, _                         = hex.DecodeString(TIMEOUT_TO_WITHDRAW_PRIZE)
	evt_price_increase_changed, _                            = hex.DecodeString(PRICE_INCREASE_CHANGED)
	evt_prize_microsecond_increase_changed, _                = hex.DecodeString(MAIN_PRIZE_MICROSECOND_INCREASE)
	evt_initial_seconds_until_prize_changed, _               = hex.DecodeString(INITIAL_SECONDS_UNTIL_PRIZE_CHANGED)
	evt_activation_time_changed, _                           = hex.DecodeString(ROUND_ACTIVATION_TIME_CHANGED)
	evt_cst_dutch_auction_duration_divisor_changed, _        = hex.DecodeString(CST_DUTCH_AUCTION_DURATION_DIVISOR_CHANGED) // previously	evt_round_start_auction_length_changed
	evt_eth_dutch_auction_duration_divisor_changed, _        = hex.DecodeString(ETH_DUTCH_AUCTION_DURATION_DIVISOR_CHANGED)
	evt_eth_dutch_auction_ending_bidprice_divisor, _         = hex.DecodeString(ETH_DUTCH_AUCTION_ENDING_PRICE_DIVISOR_CHANGED)
	evt_proxy_upgraded, _                                    = hex.DecodeString(PROXY_UPGRADED)
	evt_admin_changed, _                                     = hex.DecodeString(ADMIN_CHANGED)
	evt_cst_reward_for_bidding_changed, _                    = hex.DecodeString(CST_REWARD_FOR_BIDDING_CHANGED)
	evt_bid_cst_reward_amount_changed, _                     = hex.DecodeString(BID_CST_REWARD_AMOUNT_CHANGED)
	evt_bid_cst_reward_amount_multiplier_changed, _          = hex.DecodeString(BID_CST_REWARD_AMOUNT_MULTIPLIER_CHANGED)
	evt_cst_dutch_auction_duration_changed, _                = hex.DecodeString(CST_DUTCH_AUCTION_DURATION_CHANGED)
	evt_cst_dutch_auction_duration_change_divisor_changed, _ = hex.DecodeString(CST_DUTCH_AUCTION_DURATION_CHANGE_DIVISOR_CHANGED)
	evt_max_msg_length_changed, _                            = hex.DecodeString(MAX_MESSAGE_LENGTH)
	evt_token_script_url, _                                  = hex.DecodeString(TOKEN_SCRIPT_URL)
	evt_base_uri, _                                          = hex.DecodeString(BASE_URI)
	evt_marketing_reward_changed, _                          = hex.DecodeString(MARKETING_REWARD_CHANGED)
	evt_ownership_transferred, _                             = hex.DecodeString(OWNERSHIP_TRANSFERRED)
	evt_initialized, _                                       = hex.DecodeString(INITIALIZED)
	evt_cst_min_limit, _                                     = hex.DecodeString(STARTING_CST_MIN_LIM)
	evt_fund_transf_err, _                                   = hex.DecodeString(FUND_TRANSFER_ERR)
	evt_erc20_transf_err, _                                  = hex.DecodeString(ERC20_TRANSFER_ERR)
	evt_static_cst_reward, _                                 = hex.DecodeString(STATIC_CST_REWARD)
	evt_funds2charity, _                                     = hex.DecodeString(FUNDS_TO_CHARITY)
	evt_delay_duration_round, _                              = hex.DecodeString(DELAY_DURATION_ROUND)
	evt_first_bid_event, _                                   = hex.DecodeString(FIRST_BID_EVENT)

	// CharityWallet events
	evt_donation_received_event, _  = hex.DecodeString(DONATION_RECEIVED_EVENT)
	evt_donation_sent_event, _      = hex.DecodeString(DONATION_SENT_EVENT)
	evt_charity_receiver_changed, _ = hex.DecodeString(CHARITY_RECEIVER_CHANGED)

	// CosmicSignature events
	evt_token_name_event, _ = hex.DecodeString(TOKEN_NAME_EVENT)
	evt_mint_event, _       = hex.DecodeString(MINT_EVENT)

	// PrizeWallet events
	evt_eth_prize_deposit, _    = hex.DecodeString(ETH_PRIZE_DEPOSIT_EVENT)
	evt_eth_prize_withdrawal, _ = hex.DecodeString(ETH_PRIZE_WITHDRAWAL_EVENT)

	// ERC20 events
	evt_transfer, _ = hex.DecodeString(TRANSFER_EVT)

	// StakingWallet events
	evt_cst_nft_staked, _      = hex.DecodeString(CST_NFT_STAKED_EVENT)
	evt_rwalk_nft_staked, _    = hex.DecodeString(RWALK_NFT_STAKED_EVENT)
	evt_nft_unstaked_rwalk, _  = hex.DecodeString(NFT_UNSTAKED_RWALK)
	evt_nft_unstaked_cst, _    = hex.DecodeString(NFT_UNSTAKED_CST)
	evt_staking_eth_deposit, _ = hex.DecodeString(STAKING_ETH_DEPOSIT_EVENT)

	// MarketingWallet events
	evt_marketing_reward_paid, _ = hex.DecodeString(MARKETING_REWARD_PAID)
	evt_treasurer_changed, _     = hex.DecodeString(TREASURER_CHANGED)

	inspected_events []InspectedEvent
)

func build_list_of_inspected_events_layer1(cosmic_sig_aid int64) []InspectedEvent {

	// this is the list of all the events we read (not necesarilly insert into the DB, but check on them)
	inspected_events = make([]InspectedEvent, 0, 32)
	inspected_events = append(inspected_events,
		// this list matches the order of main.go event variables in `var` declaration
		InspectedEvent{
			Signature:   hex.EncodeToString(evt_prize_claim_event[:4]),
			ContractAid: 0,
		},
		InspectedEvent{
			Signature:   hex.EncodeToString(evt_bid_event[:4]),
			ContractAid: 0,
		},
		InspectedEvent{
			Signature:   hex.EncodeToString(evt_bid_event_v2[:4]),
			ContractAid: 0,
		},
		InspectedEvent{
			Signature:   hex.EncodeToString(evt_eth_donated_event[:4]),
			ContractAid: 0,
		},
		InspectedEvent{
			Signature:   hex.EncodeToString(evt_eth_donated_wi_event[:4]),
			ContractAid: 0,
		},
		InspectedEvent{
			Signature:   hex.EncodeToString(evt_erc20_donated[:4]),
			ContractAid: 0,
		},
		InspectedEvent{
			Signature:   hex.EncodeToString(evt_nft_donation_event[:4]),
			ContractAid: 0,
		},
		InspectedEvent{
			Signature:   hex.EncodeToString(evt_raffle_nft_prize_event[:4]),
			ContractAid: 0,
		},
		InspectedEvent{
			Signature:   hex.EncodeToString(evt_raffle_eth_prize_event[:4]),
			ContractAid: 0,
		},
		InspectedEvent{
			Signature:   hex.EncodeToString(evt_endurance_prize_event[:4]),
			ContractAid: 0,
		},
		InspectedEvent{
			Signature:   hex.EncodeToString(evt_lastcst_bidder_prize_event[:4]),
			ContractAid: 0,
		},
		InspectedEvent{
			Signature:   hex.EncodeToString(evt_donated_token_claimed[:4]),
			ContractAid: 0,
		},
		InspectedEvent{
			Signature:   hex.EncodeToString(evt_donated_nft_claimed[:4]),
			ContractAid: 0,
		},
		InspectedEvent{
			Signature:   hex.EncodeToString(evt_charity_percentage_changed[:4]),
			ContractAid: 0,
		},
		InspectedEvent{
			Signature:   hex.EncodeToString(evt_prize_percentage_changed[:4]),
			ContractAid: 0,
		},
		InspectedEvent{
			Signature:   hex.EncodeToString(evt_raffle_percentage_changed[:4]),
			ContractAid: 0,
		},
		InspectedEvent{
			Signature:   hex.EncodeToString(evt_staking_percentage_changed[:4]),
			ContractAid: 0,
		},
		InspectedEvent{
			Signature:   hex.EncodeToString(evt_chrono_percentage_changed[:4]),
			ContractAid: 0,
		},
		InspectedEvent{
			Signature:   hex.EncodeToString(evt_num_raffle_eth_winners_bidding_changed[:4]),
			ContractAid: 0,
		},
		InspectedEvent{
			Signature:   hex.EncodeToString(evt_num_raffle_nft_winners_bidding_changed[:4]),
			ContractAid: 0,
		},
		InspectedEvent{
			Signature:   hex.EncodeToString(evt_num_raffle_nft_winners_staking_rwalk_changed[:4]),
			ContractAid: 0,
		},
		InspectedEvent{
			Signature:   hex.EncodeToString(evt_charity_wallet_changed[:4]),
			ContractAid: 0,
		},
		InspectedEvent{
			Signature:   hex.EncodeToString(evt_rwalk_address_changed[:4]),
			ContractAid: 0,
		},
		InspectedEvent{
			Signature:   hex.EncodeToString(evt_prizes_wallet_address_changed[:4]),
			ContractAid: 0,
		},
		InspectedEvent{
			Signature:   hex.EncodeToString(evt_staking_wallet_cst_address_changed[:4]),
			ContractAid: 0,
		},
		InspectedEvent{
			Signature:   hex.EncodeToString(evt_staking_wallet_rwalk_address_changed[:4]),
			ContractAid: 0,
		},
		InspectedEvent{
			Signature:   hex.EncodeToString(evt_marketing_address_changed[:4]),
			ContractAid: 0,
		},
		InspectedEvent{
			Signature:   hex.EncodeToString(evt_treasurer_changed[:4]),
			ContractAid: 0,
		},
		InspectedEvent{
			Signature:   hex.EncodeToString(evt_costok_address_changed[:4]),
			ContractAid: 0,
		},
		InspectedEvent{
			Signature:   hex.EncodeToString(evt_cossig_address_changed[:4]),
			ContractAid: 0,
		},
		InspectedEvent{
			Signature:   hex.EncodeToString(evt_time_increase_changed[:4]),
			ContractAid: 0,
		},
		InspectedEvent{
			Signature:   hex.EncodeToString(evt_timeout_claimprize_changed[:4]),
			ContractAid: 0,
		},
		InspectedEvent{
			Signature:   hex.EncodeToString(evt_timeout_to_withdraw_prize[:4]),
			ContractAid: 0,
		},
		InspectedEvent{
			Signature:   hex.EncodeToString(evt_price_increase_changed[:4]),
			ContractAid: 0,
		},
		InspectedEvent{
			Signature:   hex.EncodeToString(evt_prize_microsecond_increase_changed[:4]),
			ContractAid: 0,
		},
		InspectedEvent{
			Signature:   hex.EncodeToString(evt_initial_seconds_until_prize_changed[:4]),
			ContractAid: 0,
		},
		InspectedEvent{
			Signature:   hex.EncodeToString(evt_activation_time_changed[:4]),
			ContractAid: 0,
		},
		InspectedEvent{
			Signature:   hex.EncodeToString(evt_cst_dutch_auction_duration_divisor_changed[:4]),
			ContractAid: 0,
		},
		InspectedEvent{
			Signature:   hex.EncodeToString(evt_cst_dutch_auction_duration_changed[:4]),
			ContractAid: 0,
		},
		InspectedEvent{
			Signature:   hex.EncodeToString(evt_cst_dutch_auction_duration_change_divisor_changed[:4]),
			ContractAid: 0,
		},
		InspectedEvent{
			Signature:   hex.EncodeToString(evt_eth_dutch_auction_duration_divisor_changed[:4]),
			ContractAid: 0,
		},
		InspectedEvent{
			Signature:   hex.EncodeToString(evt_eth_dutch_auction_ending_bidprice_divisor[:4]),
			ContractAid: 0,
		},
		InspectedEvent{
			Signature:   hex.EncodeToString(evt_donation_received_event[:4]),
			ContractAid: 0,
		},
		InspectedEvent{
			Signature:   hex.EncodeToString(evt_donation_sent_event[:4]),
			ContractAid: 0,
		},
		InspectedEvent{
			Signature:   hex.EncodeToString(evt_charity_receiver_changed[:4]),
			ContractAid: 0,
		},
		InspectedEvent{
			Signature:   hex.EncodeToString(evt_token_name_event[:4]),
			ContractAid: 0,
		},
		InspectedEvent{
			Signature:   hex.EncodeToString(evt_mint_event[:4]),
			ContractAid: 0,
		},
		InspectedEvent{
			Signature:   hex.EncodeToString(evt_eth_prize_deposit[:4]),
			ContractAid: 0,
		},
		InspectedEvent{
			Signature:   hex.EncodeToString(evt_staking_eth_deposit[:4]),
			ContractAid: 0,
		},
		InspectedEvent{
			Signature:   hex.EncodeToString(evt_eth_prize_withdrawal[:4]),
			ContractAid: 0,
		},
		InspectedEvent{
			Signature:   hex.EncodeToString(evt_transfer[:4]),
			ContractAid: cosmic_sig_aid,
		},
		InspectedEvent{
			Signature:   hex.EncodeToString(evt_cst_nft_staked[:4]),
			ContractAid: 0,
		},
		InspectedEvent{
			Signature:   hex.EncodeToString(evt_rwalk_nft_staked[:4]),
			ContractAid: 0,
		},
		InspectedEvent{
			Signature:   hex.EncodeToString(evt_nft_unstaked_rwalk[:4]),
			ContractAid: 0,
		},
		InspectedEvent{
			Signature:   hex.EncodeToString(evt_nft_unstaked_cst[:4]),
			ContractAid: 0,
		},
		InspectedEvent{
			Signature:   hex.EncodeToString(evt_marketing_reward_paid[:4]),
			ContractAid: 0,
		},
		InspectedEvent{
			Signature:   hex.EncodeToString(evt_cst_reward_for_bidding_changed[:4]),
			ContractAid: 0,
		},
		InspectedEvent{
			Signature:   hex.EncodeToString(evt_bid_cst_reward_amount_changed[:4]),
			ContractAid: 0,
		},
		InspectedEvent{
			Signature:   hex.EncodeToString(evt_bid_cst_reward_amount_multiplier_changed[:4]),
			ContractAid: 0,
		},
		InspectedEvent{
			Signature:   hex.EncodeToString(evt_static_cst_reward[:4]),
			ContractAid: 0,
		},
		InspectedEvent{
			Signature:   hex.EncodeToString(evt_max_msg_length_changed[:4]),
			ContractAid: 0,
		},
		InspectedEvent{
			Signature:   hex.EncodeToString(evt_token_script_url[:4]),
			ContractAid: 0,
		},
		InspectedEvent{
			Signature:   hex.EncodeToString(evt_base_uri[:4]),
			ContractAid: 0,
		},
		InspectedEvent{
			Signature:   hex.EncodeToString(evt_proxy_upgraded[:4]),
			ContractAid: 0,
		},
		InspectedEvent{
			Signature:   hex.EncodeToString(evt_admin_changed[:4]),
			ContractAid: 0,
		},
		InspectedEvent{
			Signature:   hex.EncodeToString(evt_marketing_reward_changed[:4]),
			ContractAid: 0,
		},
		InspectedEvent{
			Signature:   hex.EncodeToString(evt_ownership_transferred[:4]),
			ContractAid: 0,
		},
		InspectedEvent{
			Signature:   hex.EncodeToString(evt_initialized[:4]),
			ContractAid: 0,
		},
		InspectedEvent{
			Signature:   hex.EncodeToString(evt_chrono_warrior_prize_event[:4]),
			ContractAid: 0,
		},
		InspectedEvent{
			Signature:   hex.EncodeToString(evt_cst_min_limit[:4]),
			ContractAid: 0,
		},
		InspectedEvent{
			Signature:   hex.EncodeToString(evt_fund_transf_err[:4]),
			ContractAid: 0,
		},
		InspectedEvent{
			Signature:   hex.EncodeToString(evt_erc20_transf_err[:4]),
			ContractAid: 0,
		},
		InspectedEvent{
			Signature:   hex.EncodeToString(evt_funds2charity[:4]),
			ContractAid: 0,
		},
		InspectedEvent{
			Signature:   hex.EncodeToString(evt_delay_duration_round[:4]),
			ContractAid: 0,
		},
		InspectedEvent{
			Signature:   hex.EncodeToString(evt_first_bid_event[:4]),
			ContractAid: 0,
		},
	)
	return inspected_events
}

// select_event_and_process dispatches the log to every matching event handler.
// Handlers that touch the base store propagate DB errors, which stop the
// dispatch and are returned to the polling loop.
func select_event_and_process(log *types.Log, evtlog *EthereumEventLog) error {

	if 0 == bytes.Compare(log.Topics[0].Bytes(), evt_prize_claim_event) {
		proc_prize_claim_event(log, evtlog)
	}
	if 0 == bytes.Compare(log.Topics[0].Bytes(), evt_bid_event) {
		if err := proc_bid_event_v1(log, evtlog); err != nil {
			return err
		}
	}
	if 0 == bytes.Compare(log.Topics[0].Bytes(), evt_bid_event_v2) {
		if err := proc_bid_event_v2(log, evtlog); err != nil {
			return err
		}
	}
	if 0 == bytes.Compare(log.Topics[0].Bytes(), evt_eth_donated_event) {
		proc_donation_event(log, evtlog)
	}
	if 0 == bytes.Compare(log.Topics[0].Bytes(), evt_eth_donated_wi_event) {
		proc_donation_with_info_event(log, evtlog)
	}
	if 0 == bytes.Compare(log.Topics[0].Bytes(), evt_donation_received_event) {
		proc_donation_received_event(log, evtlog)
	}
	if 0 == bytes.Compare(log.Topics[0].Bytes(), evt_donation_sent_event) {
		proc_donation_sent_event(log, evtlog)
	}
	if 0 == bytes.Compare(log.Topics[0].Bytes(), evt_nft_donation_event) {
		proc_nft_donation_event(log, evtlog)
	}
	if 0 == bytes.Compare(log.Topics[0].Bytes(), evt_erc20_donated) {
		proc_erc20_donated_event(log, evtlog)
	}
	if 0 == bytes.Compare(log.Topics[0].Bytes(), evt_charity_receiver_changed) {
		proc_charity_address_changed_unified(log, evtlog)
	}
	if 0 == bytes.Compare(log.Topics[0].Bytes(), evt_token_name_event) {
		proc_token_name_event(log, evtlog)
	}
	if 0 == bytes.Compare(log.Topics[0].Bytes(), evt_mint_event) {
		proc_mint_event(log, evtlog)
	}
	if 0 == bytes.Compare(log.Topics[0].Bytes(), evt_eth_prize_deposit) {
		proc_prizes_eth_deposit_event(log, evtlog)
	}
	if 0 == bytes.Compare(log.Topics[0].Bytes(), evt_eth_prize_withdrawal) {
		proc_eth_prize_withdrawal_event(log, evtlog)
	}
	if 0 == bytes.Compare(log.Topics[0].Bytes(), evt_raffle_eth_prize_event) {
		proc_raffle_eth_winner_event(log, evtlog)
	}
	if 0 == bytes.Compare(log.Topics[0].Bytes(), evt_raffle_nft_prize_event) {
		proc_raffle_nft_winner_event(log, evtlog)
	}
	if 0 == bytes.Compare(log.Topics[0].Bytes(), evt_endurance_prize_event) {
		proc_endurance_winner_event(log, evtlog)
	}
	if 0 == bytes.Compare(log.Topics[0].Bytes(), evt_lastcst_bidder_prize_event) {
		proc_lastcst_bidder_winner_event(log, evtlog)
	}
	if 0 == bytes.Compare(log.Topics[0].Bytes(), evt_chrono_warrior_prize_event) {
		proc_chrono_warrior_event(log, evtlog)
	}
	if 0 == bytes.Compare(log.Topics[0].Bytes(), evt_donated_token_claimed) {
		proc_donated_token_claimed_event(log, evtlog)
	}
	if 0 == bytes.Compare(log.Topics[0].Bytes(), evt_donated_nft_claimed) {
		proc_donated_nft_claimed_event(log, evtlog)
	}
	if 0 == bytes.Compare(log.Topics[0].Bytes(), evt_transfer) {
		proc_transfer_event_common(log, evtlog)
	}
	if 0 == bytes.Compare(log.Topics[0].Bytes(), evt_cst_nft_staked) {
		proc_cst_nft_staked_event(log, evtlog)
	}
	if 0 == bytes.Compare(log.Topics[0].Bytes(), evt_rwalk_nft_staked) {
		proc_rwalk_nft_staked_event(log, evtlog)
	}
	if 0 == bytes.Compare(log.Topics[0].Bytes(), evt_nft_unstaked_rwalk) {
		proc_nft_unstaked_rwalk_event(log, evtlog)
	}
	if 0 == bytes.Compare(log.Topics[0].Bytes(), evt_nft_unstaked_cst) {
		proc_nft_unstaked_cst_event(log, evtlog)
	}
	if 0 == bytes.Compare(log.Topics[0].Bytes(), evt_staking_eth_deposit) {
		proc_staking_eth_deposit_event(log, evtlog)
	}
	if 0 == bytes.Compare(log.Topics[0].Bytes(), evt_marketing_reward_paid) {
		proc_marketing_reward_paid_event(log, evtlog)
	}
	if 0 == bytes.Compare(log.Topics[0].Bytes(), evt_charity_percentage_changed) {
		proc_charity_percentage_changed_event(log, evtlog)
	}
	if 0 == bytes.Compare(log.Topics[0].Bytes(), evt_prize_percentage_changed) {
		proc_prize_percentage_changed_event(log, evtlog)
	}
	if 0 == bytes.Compare(log.Topics[0].Bytes(), evt_raffle_percentage_changed) {
		proc_raffle_percentage_changed_event(log, evtlog)
	}
	if 0 == bytes.Compare(log.Topics[0].Bytes(), evt_staking_percentage_changed) {
		proc_staking_percentage_changed_event(log, evtlog)
	}
	if 0 == bytes.Compare(log.Topics[0].Bytes(), evt_chrono_percentage_changed) {
		proc_chrono_percentage_changed_event(log, evtlog)
	}
	if 0 == bytes.Compare(log.Topics[0].Bytes(), evt_num_raffle_eth_winners_bidding_changed) {
		proc_num_raffle_eth_winners_bidding_changed_event(log, evtlog)
	}
	if 0 == bytes.Compare(log.Topics[0].Bytes(), evt_num_raffle_nft_winners_bidding_changed) {
		proc_num_raffle_nft_winners_bidding_changed_event(log, evtlog)
	}
	if 0 == bytes.Compare(log.Topics[0].Bytes(), evt_num_raffle_nft_winners_staking_rwalk_changed) {
		proc_num_raffle_nft_winners_staking_rwalk_changed_event(log, evtlog)
	}
	if 0 == bytes.Compare(log.Topics[0].Bytes(), evt_charity_wallet_changed) {
		proc_charity_address_changed_unified(log, evtlog)
	}
	if 0 == bytes.Compare(log.Topics[0].Bytes(), evt_rwalk_address_changed) {
		proc_random_walk_address_changed_event(log, evtlog)
	}
	if 0 == bytes.Compare(log.Topics[0].Bytes(), evt_prizes_wallet_address_changed) {
		proc_raffle_address_changed_event(log, evtlog)
	}
	if 0 == bytes.Compare(log.Topics[0].Bytes(), evt_staking_wallet_cst_address_changed) {
		proc_staking_wallet_cst_address_changed_event(log, evtlog)
	}
	if 0 == bytes.Compare(log.Topics[0].Bytes(), evt_staking_wallet_rwalk_address_changed) {
		proc_staking_wallet_rwalk_address_changed_event(log, evtlog)
	}
	if 0 == bytes.Compare(log.Topics[0].Bytes(), evt_marketing_address_changed) {
		proc_marketing_wallet_address_changed_event(log, evtlog)
	}
	if 0 == bytes.Compare(log.Topics[0].Bytes(), evt_treasurer_changed) {
		proc_treasurer_changed_event(log, evtlog)
	}
	if 0 == bytes.Compare(log.Topics[0].Bytes(), evt_costok_address_changed) {
		proc_cosmic_token_address_changed_event(log, evtlog)
	}
	if 0 == bytes.Compare(log.Topics[0].Bytes(), evt_cossig_address_changed) {
		proc_cosmic_signature_address_changed_event(log, evtlog)
	}
	if 0 == bytes.Compare(log.Topics[0].Bytes(), evt_proxy_upgraded) {
		proc_proxy_upgraded_event(log, evtlog)
	}
	if 0 == bytes.Compare(log.Topics[0].Bytes(), evt_admin_changed) {
		proc_admin_changed_event(log, evtlog)
	}
	if 0 == bytes.Compare(log.Topics[0].Bytes(), evt_time_increase_changed) {
		proc_time_increase_changed_event(log, evtlog)
	}
	if 0 == bytes.Compare(log.Topics[0].Bytes(), evt_timeout_claimprize_changed) {
		proc_timeout_claimprize_changed_event(log, evtlog)
	}
	if 0 == bytes.Compare(log.Topics[0].Bytes(), evt_timeout_to_withdraw_prize) {
		proc_timeout_duration_to_withdraw_prize_event(log, evtlog)
	}
	if 0 == bytes.Compare(log.Topics[0].Bytes(), evt_price_increase_changed) {
		proc_price_increase_changed_event(log, evtlog)
	}
	if 0 == bytes.Compare(log.Topics[0].Bytes(), evt_prize_microsecond_increase_changed) {
		proc_mainprize_microsecond_increase_changed(log, evtlog)
	}
	if 0 == bytes.Compare(log.Topics[0].Bytes(), evt_initial_seconds_until_prize_changed) {
		proc_initial_seconds_until_prize_changed_event(log, evtlog)
	}
	if 0 == bytes.Compare(log.Topics[0].Bytes(), evt_activation_time_changed) {
		proc_activation_time_changed_event(log, evtlog)
	}
	if 0 == bytes.Compare(log.Topics[0].Bytes(), evt_cst_dutch_auction_duration_divisor_changed) {
		proc_cst_dutch_auction_duration_divisor_changed_event(log, evtlog)
	}
	if 0 == bytes.Compare(log.Topics[0].Bytes(), evt_cst_dutch_auction_duration_changed) {
		proc_cst_dutch_auction_duration_changed_event(log, evtlog)
	}
	if 0 == bytes.Compare(log.Topics[0].Bytes(), evt_cst_dutch_auction_duration_change_divisor_changed) {
		proc_cst_dutch_auction_duration_change_divisor_changed_event(log, evtlog)
	}
	if 0 == bytes.Compare(log.Topics[0].Bytes(), evt_eth_dutch_auction_duration_divisor_changed) {
		proc_eth_dutch_auction_duration_divisor_changed_event(log, evtlog)
	}
	if 0 == bytes.Compare(log.Topics[0].Bytes(), evt_eth_dutch_auction_ending_bidprice_divisor) {
		proc_eth_dutch_auction_ending_bid_price_divisor_changed__event(log, evtlog)
	}
	if 0 == bytes.Compare(log.Topics[0].Bytes(), evt_cst_reward_for_bidding_changed) {
		proc_erc20_token_reward_changed_event(log, evtlog)
	}
	if 0 == bytes.Compare(log.Topics[0].Bytes(), evt_bid_cst_reward_amount_changed) {
		proc_bid_cst_reward_amount_changed_event(log, evtlog)
	}
	if 0 == bytes.Compare(log.Topics[0].Bytes(), evt_bid_cst_reward_amount_multiplier_changed) {
		proc_bid_cst_reward_amount_multiplier_changed_event(log, evtlog)
	}
	if 0 == bytes.Compare(log.Topics[0].Bytes(), evt_static_cst_reward) {
		proc_static_cst_reward_changed_event(log, evtlog)
	}
	if 0 == bytes.Compare(log.Topics[0].Bytes(), evt_max_msg_length_changed) {
		proc_max_msg_length_changed_event(log, evtlog)
	}
	if 0 == bytes.Compare(log.Topics[0].Bytes(), evt_token_script_url) {
		proc_token_generation_script_url_event(log, evtlog)
	}
	if 0 == bytes.Compare(log.Topics[0].Bytes(), evt_base_uri) {
		proc_base_uri_event(log, evtlog)
	}
	if 0 == bytes.Compare(log.Topics[0].Bytes(), evt_marketing_reward_changed) {
		proc_marketing_reward_changed(log, evtlog)
	}
	if 0 == bytes.Compare(log.Topics[0].Bytes(), evt_ownership_transferred) {
		proc_ownership_transferred_event(log, evtlog)
	}
	if 0 == bytes.Compare(log.Topics[0].Bytes(), evt_initialized) {
		proc_initialized_event(log, evtlog)
	}
	if 0 == bytes.Compare(log.Topics[0].Bytes(), evt_cst_min_limit) {
		proc_starting_bid_price_cst_min_limit_event(log, evtlog)
	}
	if 0 == bytes.Compare(log.Topics[0].Bytes(), evt_fund_transf_err) {
		proc_fund_transfer_failed_event(log, evtlog)
	}
	if 0 == bytes.Compare(log.Topics[0].Bytes(), evt_funds2charity) {
		proc_funds_transferred_to_charity_event(log, evtlog)
	}
	if 0 == bytes.Compare(log.Topics[0].Bytes(), evt_delay_duration_round) {
		proc_delay_duration_before_next_round_changed_event(log, evtlog)
	}
	if 0 == bytes.Compare(log.Topics[0].Bytes(), evt_first_bid_event) {
		proc_round_started_event(log, evtlog)
	}
	return nil
}
func process_single_event(evt_id int64) error {

	evtlog, err := storagew.S.Get_event_log(evt_id)
	if err != nil {
		return fmt.Errorf("process_single_event(%v): %w", evt_id, err)
	}
	var log types.Log
	err = rlp.DecodeBytes(evtlog.RlpLog, &log)
	if err != nil {
		panic(fmt.Sprintf("RLP Decode error: %v", err))
	}
	log.BlockNumber = uint64(evtlog.BlockNum)
	log.TxHash.SetBytes(common.HexToHash(evtlog.TxHash).Bytes())
	log.Address.SetBytes(common.HexToHash(evtlog.ContractAddress).Bytes())
	num_topics := len(log.Topics)
	if num_topics > 0 {
		return select_event_and_process(&log, &evtlog)
	}
	return nil
}
