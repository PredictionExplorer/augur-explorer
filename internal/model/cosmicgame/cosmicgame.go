// Package cosmicgame defines the CosmicGame data model: the event and row
// types shared by the ETL, storage and API layers.
package cosmicgame

// ContractAddrs is the deployed-contract address registry for one
// CosmicGame deployment (cg_contracts row), joined with the RandomWalk
// marketplace address for the dashboard response.
type ContractAddrs struct {
	CosmicGameAddr         string
	CosmicSignatureAddr    string
	CosmicTokenAddr        string
	CosmicDaoAddr          string
	CharityWalletAddr      string
	PrizesWalletAddr       string
	RandomWalkAddr         string
	StakingWalletCSTAddr   string
	StakingWalletRWalkAddr string
	MarketingWalletAddr    string
	// MarketplaceAddr is rw_contracts.marketplace_addr (RandomWalk NFT marketplace); included on dashboard ContractAddrs.
	MarketplaceAddr    string
	ImplementationAddr string
}

// ProcStatus is the cg-etl progress watermark (cg_proc_status row): the
// last processed event-log ID and block number.
type ProcStatus struct {
	LastEvtIdProcessed int64
	LastBlockNum       int64
}

// CGPrizeClaimEvent records a MainPrizeClaimed contract event: the main-prize
// winner of a round with the ETH and CST amounts awarded.
type CGPrizeClaimEvent struct {
	EvtId        int64
	BlockNum     int64
	TimeStamp    int64
	TxId         int64
	ContractAddr string
	RoundNum     int64
	TokenId      int64
	WinnerAddr   string
	Timeout      int64
	Amount       string
	CstAmount    string
}

// CGBidEvent records a BidPlaced contract event (one bid in a round). It
// covers both mechanics generations: the V2-only fields carry "-1" for V1
// events, and exactly one of EthPrice/CstPrice is "-1" depending on bid type.
type CGBidEvent struct {
	EvtId                   int64
	BlockNum                int64
	TimeStamp               int64
	TxId                    int64
	RandomWalkTokenId       int64
	PrizeTime               int64
	RoundNum                int64
	BidType                 int64
	ContractAddr            string
	LastBidderAddr          string
	EthPrice                string // PaidEthPrice (or -1 for CST bids)
	CstPrice                string // PaidCstPrice (or -1 for ETH bids)
	ERC20Value              string // reward of CST tokens earned for bidding
	Message                 string
	BidCstRewardAmount      string // IBiddingV2 BidPlaced; "-1" if V1 event
	CstDutchAuctionDuration string // IBiddingV2 BidPlaced; "-1" if V1 event
}

// CGDonationEvent records a plain EthDonated contract event: a direct ETH
// donation to the game without an accompanying info record.
type CGDonationEvent struct {
	EvtId        int64
	BlockNum     int64
	TimeStamp    int64
	TxId         int64
	ContractAddr string
	DonorAddr    string
	RoundNum     int64
	Amount       string
}

// CGDonationWithInfoEvent records an EthDonatedWithInfo contract event: an
// ETH donation carrying a JSON info record stored at the contract side.
type CGDonationWithInfoEvent struct {
	EvtId        int64
	BlockNum     int64
	TimeStamp    int64
	TxId         int64
	ContractAddr string
	DonorAddr    string
	RoundNum     int64
	RecordId     int64 // record ID at the contract side
	Amount       string
}

// CGDonationReceivedEvent records a charity-wallet DonationReceived contract
// event: ETH arriving at the CharityWallet.
type CGDonationReceivedEvent struct {
	EvtId        int64
	BlockNum     int64
	TimeStamp    int64
	TxId         int64
	RoundNum     int64
	ContractAddr string
	DonorAddr    string
	Amount       string
}

// CGDonationSentEvent records a charity-wallet DonationSent contract event:
// accumulated ETH forwarded from the CharityWallet to the charity address.
type CGDonationSentEvent struct {
	EvtId        int64
	BlockNum     int64
	TimeStamp    int64
	TxId         int64
	ContractAddr string
	CharityAddr  string
	Amount       string
}

// CGERC20DonationEvent records an Erc20TokenDonated contract event: an
// arbitrary ERC-20 donation attached to a bid.
type CGERC20DonationEvent struct {
	EvtId        int64
	BlockNum     int64
	TimeStamp    int64
	TxId         int64
	ContractAddr string
	TokenAddr    string
	RoundNum     int64
	DonorAddr    string
	Amount       string
	BidId        int64 // id of related bid record
}

// CGNFTDonationEvent records an NftDonated contract event: an arbitrary
// ERC-721 token donated to the game, claimable by the round winner.
type CGNFTDonationEvent struct {
	EvtId        int64
	BlockNum     int64
	TimeStamp    int64
	TxId         int64
	ContractAddr string
	TokenAddr    string
	RoundNum     int64
	DonorAddr    string
	TokenId      int64
	Index        int64
	BidId        int64 // id of related bid record
	NFTTokenURI  string
}

// CGCharityUpdatedEvent records a CharityAddressChanged event emitted by the
// CharityWallet: the destination charity was repointed.
type CGCharityUpdatedEvent struct {
	EvtId          int64
	BlockNum       int64
	TimeStamp      int64
	TxId           int64
	ContractAddr   string
	NewCharityAddr string
}

// CGTokenNameEvent records a NftNameChanged contract event: a Cosmic
// Signature token was (re)named by its owner.
type CGTokenNameEvent struct {
	EvtId        int64
	BlockNum     int64
	TimeStamp    int64
	TxId         int64
	ContractAddr string
	TokenId      int64
	TokenName    string
}

// CGMintEvent records a Cosmic Signature NFT mint (ERC-721 Transfer from the
// zero address) with the on-chain generation seed and originating round.
type CGMintEvent struct {
	EvtId        int64
	BlockNum     int64
	TimeStamp    int64
	TxId         int64
	ContractAddr string
	TokenId      int64
	OwnerAddr    string
	Seed         string
	RoundNum     int64
}

// CGPrizesEthDeposit records a PrizesWallet EthReceived event: prize ETH
// deposited for a round winner to withdraw later.
type CGPrizesEthDeposit struct {
	EvtId        int64
	BlockNum     int64
	TimeStamp    int64
	TxId         int64
	ContractAddr string
	WinnerAddr   string
	Round        int64
	WinnerIndex  int64
	Amount       string
}

// CGPrizesEthWithdrawal records a PrizesWallet EthWithdrawn event: a winner
// (or, after the timeout, any beneficiary) withdrawing deposited prize ETH.
type CGPrizesEthWithdrawal struct {
	EvtId           int64
	BlockNum        int64
	TimeStamp       int64
	TxId            int64
	ContractAddr    string
	Round           int64
	WinnerAddr      string
	BeneficiaryAddr string // Who actually claimed (can be different from winner after timeout)
	Amount          string
}

// CGRaffleNFTWinner records a raffle NFT win: a Cosmic Signature token
// awarded to a bidder or staker raffle winner of a round.
type CGRaffleNFTWinner struct {
	EvtId        int64
	BlockNum     int64
	TimeStamp    int64
	TxId         int64
	ContractAddr string
	WinnerAddr   string
	Round        int64
	WinnerIndex  int64
	CstAmount    string
	TokenId      int64
	IsRandomWalk bool
	IsStaker     bool
}

// CGRaffleETHWinner records a RaffleWinnerEthPrizeAllocated event: an ETH
// raffle prize allocated to a bidder of the round.
type CGRaffleETHWinner struct {
	EvtId        int64
	BlockNum     int64
	TimeStamp    int64
	TxId         int64
	ContractAddr string
	WinnerAddr   string
	Round        int64
	WinnerIndex  int64
	Amount       string
}

// CGEnduranceWinner records an EnduranceChampionPrizePaid event: the bidder
// who held last-bidder status for the longest single stretch of a round.
type CGEnduranceWinner struct {
	EvtId         int64
	BlockNum      int64
	TimeStamp     int64
	TxId          int64
	ContractAddr  string
	WinnerAddr    string
	Round         int64
	WinnerIndex   int64
	Erc721TokenId int64
	Erc20Amount   string
}

// CGLastBidderWinner records a LastCstBidderPrizePaid event: the prize for
// the final CST bidder of a round.
type CGLastBidderWinner struct {
	EvtId         int64
	BlockNum      int64
	TimeStamp     int64
	TxId          int64
	ContractAddr  string
	WinnerAddr    string
	Round         int64
	WinnerIndex   int64
	Erc721TokenId int64
	Erc20Amount   string
}

// CGChronoWarrior records a ChronoWarriorPrizePaid event: the bidder with the
// longest cumulative last-bidder time across the round.
type CGChronoWarrior struct {
	EvtId        int64
	BlockNum     int64
	TimeStamp    int64
	TxId         int64
	ContractAddr string
	WinnerAddr   string
	Round        int64
	WinnerIndex  int64
	EthAmount    string
	CstAmount    string
	NftId        int64
}

// CGDonatedTokenClaimed records a DonatedTokenClaimed event: a round winner
// (or timeout claimer) collecting a donated ERC-20 amount.
type CGDonatedTokenClaimed struct { // ERC20 tokens
	EvtId           int64
	BlockNum        int64
	TimeStamp       int64
	TxId            int64
	ContractAddr    string
	RoundNum        int64
	Index           int64
	TokenAddr       string
	Amount          string
	BeneficiaryAddr string
}

// CGDonatedNFTClaimed records a DonatedNftClaimed event: a round winner (or
// timeout claimer) collecting a donated ERC-721 token.
type CGDonatedNFTClaimed struct { // ERC721 tokens
	EvtId           int64
	BlockNum        int64
	TimeStamp       int64
	TxId            int64
	ContractAddr    string
	RoundNum        int64
	Index           int64
	TokenAddr       string
	TokenId         string
	BeneficiaryAddr string
}

// CGNftStakedCst records an NftStaked event of the Cosmic Signature staking
// wallet: one CST NFT entering the staking pool.
type CGNftStakedCst struct {
	EvtId           int64
	BlockNum        int64
	TimeStamp       int64
	TxId            int64
	ContractAddr    string
	ActionId        int64
	NftId           int64
	NumStakedNfts   int64
	StakerAddress   string
	RewardPerStaker string
}

// CGNftStakedRWalk records an NftStaked event of the RandomWalk staking
// wallet: one RandomWalk NFT entering the staking pool.
type CGNftStakedRWalk struct {
	EvtId         int64
	BlockNum      int64
	TimeStamp     int64
	TxId          int64
	ContractAddr  string
	ActionId      int64
	NftId         int64
	NumStakedNfts int64
	StakerAddress string
}

// CGEthDeposit records an EthDepositReceived event of the CST staking wallet:
// prize ETH distributed across the NFTs staked at deposit time.
type CGEthDeposit struct {
	EvtId           int64
	BlockNum        int64
	TimeStamp       int64
	TxId            int64
	ContractAddr    string
	RoundNum        int64
	DepositTime     int64
	DepositId       int64
	NumStakedNfts   int64
	Amount          string
	AmountPerStaker string
	AccumModulo     string
	Modulo          string
}

// CGNftUnstakedRWalk records an NftUnstaked event of the RandomWalk staking
// wallet: one RandomWalk NFT leaving the staking pool.
type CGNftUnstakedRWalk struct {
	EvtId         int64
	BlockNum      int64
	TimeStamp     int64
	TxId          int64
	ContractAddr  string
	RoundNum      int64
	ActionId      int64
	NftId         int64
	NumStakedNfts int64
	StakerAddress string
}

// CGNftUnstakedCst records an NftUnstaked event of the CST staking wallet:
// one CST NFT leaving the pool together with its accumulated ETH reward.
type CGNftUnstakedCst struct {
	EvtId          int64
	BlockNum       int64
	TimeStamp      int64
	TxId           int64
	ContractAddr   string
	RoundNum       int64
	ActionId       int64
	ActionCounter  int64
	NftId          int64
	NumStakedNfts  int64
	StakerAddress  string
	RewardAmount   string
	RewardPerToken string
}

// CGMarketingRewardPaid records a RewardPaid event of the MarketingWallet:
// CST paid out to a marketer by the treasurer.
type CGMarketingRewardPaid struct {
	EvtId        int64
	BlockNum     int64
	TimeStamp    int64
	TxId         int64
	ContractAddr string
	ActionId     int64
	DepositId    int64
	Amount       string
	Marketer     string
}

// CGERC721Transfer records an ERC-721 Transfer event of the Cosmic Signature
// NFT contract (mints and burns use the zero address).
type CGERC721Transfer struct {
	EvtId     int64
	BlockNum  int64
	TxId      int64
	TimeStamp int64
	Contract  string
	From      string
	To        string
	TokenId   int64
}

// CGERC20Transfer records an ERC-20 Transfer event of the CosmicToken
// contract (mints and burns use the zero address).
type CGERC20Transfer struct {
	EvtId     int64
	BlockNum  int64
	TxId      int64
	TimeStamp int64
	Contract  string
	From      string
	To        string
	Value     string
}

// CGCharityPercentageChanged records an admin change of the charity
// percentage game parameter.
type CGCharityPercentageChanged struct {
	EvtId                int64
	BlockNum             int64
	TxId                 int64
	TimeStamp            int64
	Contract             string
	NewCharityPercentage string
}

// CGPrizePercentageChanged records an admin change of the main-prize
// percentage game parameter.
type CGPrizePercentageChanged struct {
	EvtId              int64
	BlockNum           int64
	TxId               int64
	TimeStamp          int64
	Contract           string
	NewPrizePercentage string
}

// CGRafflePercentageChanged records an admin change of the raffle percentage
// game parameter.
type CGRafflePercentageChanged struct {
	EvtId               int64
	BlockNum            int64
	TxId                int64
	TimeStamp           int64
	Contract            string
	NewRafflePercentage string
}

// CGStakingPercentageChanged records an admin change of the staking
// percentage game parameter.
type CGStakingPercentageChanged struct {
	EvtId                int64
	BlockNum             int64
	TxId                 int64
	TimeStamp            int64
	Contract             string
	NewStakingPercentage string
}

// CGChronoPercentageChanged records an admin change of the chrono-warrior
// percentage game parameter.
type CGChronoPercentageChanged struct {
	EvtId               int64
	BlockNum            int64
	TxId                int64
	TimeStamp           int64
	Contract            string
	NewChronoPercentage string
}

// CGNumRaffleETHWinnersBiddingChanged records an admin change of the number
// of bidder-raffle ETH winners per round.
type CGNumRaffleETHWinnersBiddingChanged struct {
	EvtId                         int64
	BlockNum                      int64
	TxId                          int64
	TimeStamp                     int64
	Contract                      string
	NewNumRaffleETHWinnersBidding int64
}

// CGNumRaffleNFTWinnersBiddingChanged records an admin change of the number
// of bidder-raffle NFT winners per round.
type CGNumRaffleNFTWinnersBiddingChanged struct {
	EvtId                         int64
	BlockNum                      int64
	TxId                          int64
	TimeStamp                     int64
	Contract                      string
	NewNumRaffleNFTWinnersBidding int64
}

// CGNumRaffleNFTWinnersStakingRWalkChanged records an admin change of the
// number of RandomWalk-staker raffle NFT winners per round.
type CGNumRaffleNFTWinnersStakingRWalkChanged struct {
	EvtId                              int64
	BlockNum                           int64
	TxId                               int64
	TimeStamp                          int64
	Contract                           string
	NewNumRaffleNFTWinnersStakingRWalk int64
}

// CGCharityAddressChanged records an admin change of the game's charity
// address (the CosmicSignatureGame-emitted variant).
type CGCharityAddressChanged struct {
	EvtId      int64
	BlockNum   int64
	TxId       int64
	TimeStamp  int64
	Contract   string
	NewCharity string
}

// CGRandomWalkAddressChanged records an admin repoint of the RandomWalk NFT
// contract address used by the game.
type CGRandomWalkAddressChanged struct {
	EvtId         int64
	BlockNum      int64
	TxId          int64
	TimeStamp     int64
	Contract      string
	NewRandomWalk string
}

// CGPrizeWalletAddressChanged records an admin repoint of the PrizesWallet
// contract address used by the game.
type CGPrizeWalletAddressChanged struct {
	EvtId          int64
	BlockNum       int64
	TxId           int64
	TimeStamp      int64
	Contract       string
	NewPrizeWallet string
}

// CGStakingWalletCSTAddressChanged records an admin repoint of the Cosmic
// Signature staking wallet address used by the game.
type CGStakingWalletCSTAddressChanged struct {
	EvtId               int64
	BlockNum            int64
	TxId                int64
	TimeStamp           int64
	Contract            string
	NewStakingWalletCST string
}

// CGStakingWalletRWalkAddressChanged records an admin repoint of the
// RandomWalk staking wallet address used by the game.
type CGStakingWalletRWalkAddressChanged struct {
	EvtId                 int64
	BlockNum              int64
	TxId                  int64
	TimeStamp             int64
	Contract              string
	NewStakingWalletRWalk string
}

// CGMarketingWalletAddressChanged records an admin repoint of the
// MarketingWallet contract address used by the game.
type CGMarketingWalletAddressChanged struct {
	EvtId              int64
	BlockNum           int64
	TxId               int64
	TimeStamp          int64
	Contract           string
	NewMarketingWallet string
}

// CGTreasurerAddressChanged records a MarketingWallet treasurer change: the
// address allowed to pay marketing rewards.
type CGTreasurerAddressChanged struct {
	EvtId        int64
	BlockNum     int64
	TxId         int64
	TimeStamp    int64
	Contract     string
	NewTreasurer string
}

// CGCosmicTokenAddressChanged records an admin repoint of the CosmicToken
// (ERC-20) contract address used by the game.
type CGCosmicTokenAddressChanged struct {
	EvtId          int64
	BlockNum       int64
	TxId           int64
	TimeStamp      int64
	Contract       string
	NewCosmicToken string
}

// CGCosmicSignatureAddressChanged records an admin repoint of the Cosmic
// Signature NFT contract address used by the game.
type CGCosmicSignatureAddressChanged struct {
	EvtId              int64
	BlockNum           int64
	TxId               int64
	TimeStamp          int64
	Contract           string
	NewCosmicSignature string
}

// CGUpgraded records the ERC-1967 proxy Upgraded event: the game
// implementation contract behind the proxy changed.
type CGUpgraded struct { // openzeppelin proxy Upgraded event
	EvtId          int64
	BlockNum       int64
	TxId           int64
	TimeStamp      int64
	Contract       string
	Implementation string
}

// CGAdminChanged records the ERC-1967 proxy AdminChanged event: the proxy
// admin moved from OldAdmin to NewAdmin.
type CGAdminChanged struct { // openzeppelin proxy AdminChanged event
	EvtId     int64
	BlockNum  int64
	TxId      int64
	TimeStamp int64
	Contract  string
	OldAdmin  string
	NewAdmin  string
}

// CGTimeIncreaseChanged records an admin change of the per-bid prize-time
// increase parameter.
type CGTimeIncreaseChanged struct {
	EvtId           int64
	BlockNum        int64
	TxId            int64
	TimeStamp       int64
	Contract        string
	NewTimeIncrease string
}

// CGTimeoutClaimPrizeChanged records an admin change of the timeout after
// which anyone may claim an unclaimed main prize.
type CGTimeoutClaimPrizeChanged struct {
	EvtId      int64
	BlockNum   int64
	TxId       int64
	TimeStamp  int64
	Contract   string
	NewTimeout int64
}

// CGTimeoutToWithdrawPrizeChanged records an admin change of the timeout
// after which anyone may withdraw an unclaimed PrizesWallet balance.
type CGTimeoutToWithdrawPrizeChanged struct {
	EvtId      int64
	BlockNum   int64
	TxId       int64
	TimeStamp  int64
	Contract   string
	NewTimeout int64
}

// CGPriceIncreaseChanged records an admin change of the per-bid ETH price
// increase parameter.
type CGPriceIncreaseChanged struct {
	EvtId            int64
	BlockNum         int64
	TxId             int64
	TimeStamp        int64
	Contract         string
	NewPriceIncrease string
}

// CGMainPrizeMicroSecondsIncreaseChanged records an admin change of the
// V2-mechanics per-bid main-prize time increase (microseconds).
type CGMainPrizeMicroSecondsIncreaseChanged struct {
	EvtId           int64
	BlockNum        int64
	TxId            int64
	TimeStamp       int64
	Contract        string
	NewMicroseconds string
}

// CGInitialSecondsUntilPrizeChanged records an admin change of the initial
// main-prize countdown at round start.
type CGInitialSecondsUntilPrizeChanged struct {
	EvtId                       int64
	BlockNum                    int64
	TxId                        int64
	TimeStamp                   int64
	Contract                    string
	NewInitialSecondsUntilPrize string
}

// CGActivationTimeChanged records an admin change of the game activation
// time (when bidding opens for the next round).
type CGActivationTimeChanged struct {
	EvtId             int64
	BlockNum          int64
	TxId              int64
	TimeStamp         int64
	Contract          string
	NewActivationTime string
}

// CGCstDutchAuctionDurationDivisorChanged records an admin change of the CST
// Dutch-auction duration divisor (V2 mechanics).
type CGCstDutchAuctionDurationDivisorChanged struct {
	EvtId     int64
	BlockNum  int64
	TxId      int64
	TimeStamp int64
	Contract  string
	NewValue  string
}

// CGCstDutchAuctionDurationChangeDivisorChanged records an admin change of
// the CST Dutch-auction duration-change divisor (V2 mechanics).
type CGCstDutchAuctionDurationChangeDivisorChanged struct {
	EvtId     int64
	BlockNum  int64
	TxId      int64
	TimeStamp int64
	Contract  string
	NewValue  string
}

// CGEthDutchAuctionDurationDivisorChanged records an admin change of the ETH
// Dutch-auction duration divisor (V2 mechanics).
type CGEthDutchAuctionDurationDivisorChanged struct {
	EvtId     int64
	BlockNum  int64
	TxId      int64
	TimeStamp int64
	Contract  string
	NewValue  string
}

// CGEthDutchAuctionEndingBidPriceDivisorChanged records an admin change of
// the ETH Dutch-auction ending-bid-price divisor (V2 mechanics).
type CGEthDutchAuctionEndingBidPriceDivisorChanged struct {
	EvtId     int64
	BlockNum  int64
	TxId      int64
	TimeStamp int64
	Contract  string
	NewValue  string
}

// CGStaticCstReward records an admin change of the flat CST reward paid per
// bid (TokenReward).
type CGStaticCstReward struct {
	EvtId     int64
	BlockNum  int64
	TxId      int64
	TimeStamp int64
	Contract  string
	NewReward string
}

// CGMarketingRewardChanged records an admin change of the CST amount minted
// to the MarketingWallet per round.
type CGMarketingRewardChanged struct {
	EvtId     int64
	BlockNum  int64
	TxId      int64
	TimeStamp int64
	Contract  string
	NewReward string
}

// CGCstRewardForBiddingChanged records an admin change of the CST reward
// multiplier for CST bids (V2 mechanics).
type CGCstRewardForBiddingChanged struct {
	EvtId     int64
	BlockNum  int64
	TxId      int64
	TimeStamp int64
	Contract  string
	NewReward string
}

// CGMaxMessageLengthChanged records an admin change of the maximum bid
// message length.
type CGMaxMessageLengthChanged struct {
	EvtId            int64
	BlockNum         int64
	TxId             int64
	TimeStamp        int64
	Contract         string
	NewMessageLength string
}

// CGTokenGenerationScriptURL records an admin change of the NFT generation
// script URL published by the Cosmic Signature NFT contract.
type CGTokenGenerationScriptURL struct {
	EvtId     int64
	BlockNum  int64
	TxId      int64
	TimeStamp int64
	Contract  string
	NewURL    string
}

// CGBaseURIEvent records an admin change of the Cosmic Signature NFT
// metadata base URI.
type CGBaseURIEvent struct {
	EvtId     int64
	BlockNum  int64
	TxId      int64
	TimeStamp int64
	Contract  string
	NewURI    string
}

// CGOwnershipTransferred records an Ownable OwnershipTransferred event from
// one of the game contracts, discriminated by ContractCode.
type CGOwnershipTransferred struct {
	EvtId        int64
	BlockNum     int64
	TxId         int64
	TimeStamp    int64
	Contract     string
	PrevOwner    string
	NewOwner     string
	ContractCode int64
}

// CGInitialized records the OpenZeppelin Initialized event emitted when a
// proxy implementation initializes with a version number.
type CGInitialized struct {
	EvtId     int64
	BlockNum  int64
	TxId      int64
	TimeStamp int64
	Contract  string
	Version   int64
}

// CGCstMinLimit records an admin change of the CST Dutch-auction minimum
// starting price limit.
type CGCstMinLimit struct {
	EvtId       int64
	BlockNum    int64
	TxId        int64
	TimeStamp   int64
	Contract    string
	CstMinLimit string
}

// CGFundTransferFailed records a FundTransferFailed contract error event: an
// ETH send from the game to Destination reverted.
type CGFundTransferFailed struct {
	EvtId       int64
	BlockNum    int64
	TxId        int64
	TimeStamp   int64
	Contract    string
	Destination string
	Amount      string
}

// CGErc20TransferFailed records an ERC20TransferFailed contract error event:
// a CST transfer from the game to Destination reverted.
type CGErc20TransferFailed struct {
	EvtId       int64
	BlockNum    int64
	TxId        int64
	TimeStamp   int64
	Contract    string
	Destination string
	Amount      string
}

// CGFundsToCharity records a FundsTransferredToCharity event: ETH moved
// directly from the game to the charity address.
type CGFundsToCharity struct {
	EvtId       int64
	BlockNum    int64
	TxId        int64
	TimeStamp   int64
	Contract    string
	CharityAddr string
	Amount      string
}

// CGNextRoundDelayDuration records an admin change of the delay between a
// main-prize claim and the next round's activation (V2 mechanics).
type CGNextRoundDelayDuration struct {
	EvtId     int64
	BlockNum  int64
	TxId      int64
	TimeStamp int64
	Contract  string
	NewValue  int64
}

// CGRoundStarted records a RoundActivated event: a new round opened for
// bidding at StartTimestamp (V2 mechanics).
type CGRoundStarted struct {
	EvtId          int64
	BlockNum       int64
	TxId           int64
	TimeStamp      int64
	Contract       string
	RoundNum       int64
	StartTimestamp int64
}
