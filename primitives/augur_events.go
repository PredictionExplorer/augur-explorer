// Types related to Augur contracts

package primitives
import (
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)
type EUniverseCreated struct {//0xe36b09d83f9cfa88c37f071fc2cfb5ff30b764cbd98088e70d965573c9ce5bbd
	ParentUniverse    common.Address
	ChildUniverse     common.Address
	PayoutNumerators  []*big.Int
	CreationTimestamp *big.Int
	Raw               types.Log // Blockchain specific contextual infos
}
type EUniverseForked struct {//0xce5b6de2a0053ebc6c04e68bcbb9f0a1f2deeb7049c72881e198f95b5752db82
	Universe      common.Address
	ForkingMarket common.Address
	Raw           types.Log // Blockchain specific contextual infos
}
type ERegisterContract struct {//0xa037dd0e01f0488a530cb17065a6d2f284fae016004fc744ee2a41d5cacf85d5
	ContractAddress common.Address
	Key             [32]byte
	Raw             types.Log // Blockchain specific contextual infos
}
type EMarketMigrated struct {//0xc3cf07f8fa0fafc25a9dd0bad2cd6b961c55dad41b42c8ef8f931bc40e41e08c
	Market           common.Address
	OriginalUniverse common.Address
	NewUniverse      common.Address
	Raw              types.Log // Blockchain specific contextual infos
}
type EMarketTransferred struct {//0x55f2a7bfa32e835c3f3c3cff653a3d11c077ce1b00c5a41c6aaf09eedc1ac3b2
	Universe common.Address
	Market   common.Address
	From     common.Address
	To       common.Address
	Raw      types.Log // Blockchain specific contextual infos
}
type EValidityBondChanged struct {//0x69af68e366a0570364e3a086f3b5ac79f08ecc3f93eaccbfcf3864809b12b5d8
	Universe     common.Address
	ValidityBond *big.Int
	Raw          types.Log // Blockchain specific contextual infos
}
type WarpSyncDataUpdated struct {//0x7589653fe5a2ab3ccc12538316852339868efdd9d3bd0b84d055cf224cf96873
	Universe      common.Address
	WarpSyncHash  *big.Int
	MarketEndTime *big.Int
	Raw           types.Log // Blockchain specific contextual infos
}
type FinishDeployment struct {//0xf06c142f93fdd00fbcd1e8f3d82e6f22667d52df764b39570061a7dbeea09be0
	Raw types.Log // Blockchain specific contextual infos
}



/// Market Events
type EMarketCreated struct {//0xea17ae24b0d40ea7962a6d832db46d1f81eaec1562946d0830d1c21d4c000ec1
	Universe             common.Address
	EndTime              *big.Int
	ExtraInfo            string
	Market               common.Address
	MarketCreator        common.Address
	DesignatedReporter   common.Address
	FeePerCashInAttoCash *big.Int
	Prices               []*big.Int
	MarketType           uint8
	NumTicks             *big.Int
	Outcomes             [][32]byte
	NoShowBond           *big.Int
	Timestamp            *big.Int
	Raw                  types.Log // Blockchain specific contextual infos
}
type EMarketOIChanged struct {//0x213a05b9ad8567c2f8fa868e7375e5bf30e69add0dbb5913ca8a3e58c815c268
	Universe common.Address
	Market   common.Address
	MarketOI *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}
type EMarketFinalized struct {//0x6d39632c2dc10305bf5771cfff4af1851f07c03ea27b821cad382466bdf7a21f
	Universe                common.Address
	Market                  common.Address
	Timestamp               *big.Int
	WinningPayoutNumerators []*big.Int
	Raw                     types.Log // Blockchain specific contextual infos
}
type EMarketVolumeChanged_v1 struct {	// previous version of the event (to be deleted on Augur Release)
	Universe       common.Address
	Market         common.Address
	Volume         *big.Int
	OutcomeVolumes []*big.Int
	Timestamp      *big.Int
	Raw            types.Log // Blockchain specific contextual infos
}
type EMarketVolumeChanged_v2 struct {
	Universe       common.Address
	Market         common.Address
	Volume         *big.Int
	OutcomeVolumes []*big.Int
	TotalTrades    *big.Int
	Timestamp      *big.Int
	Raw            types.Log // Blockchain specific contextual infos
}
type ENoShowBondChanged struct {//0xd1fc3f2cb1387e602db0e6f8f22649df65df5246eeff281cf6d1ef62feda4ece
	Universe   common.Address
	NoShowBond *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}


/// Trading Events
type EOrderEvent struct {
	Universe     common.Address
	Market       common.Address
	EventType    uint8
	OrderType    uint8
	OrderId      [32]byte
	TradeGroupId [32]byte
	AddressData  []common.Address
	Uint256Data  []*big.Int
	Raw          types.Log // Blockchain specific contextual infos
}
type ECancelZeroXOrder struct {
	Universe  common.Address
	Market    common.Address
	Account   common.Address
	Outcome   *big.Int
	Price     *big.Int
	Amount    *big.Int
	OrderType uint8
	OrderHash [32]byte
	Raw       types.Log // Blockchain specific contextual infos
}
type EProfitLossChanged struct {
	Universe       common.Address
	Market         common.Address
	Account        common.Address
	Outcome        *big.Int
	NetPosition    *big.Int
	AvgPrice       *big.Int
	RealizedProfit *big.Int
	FrozenFunds    *big.Int
	RealizedCost   *big.Int
	Timestamp      *big.Int
	Raw            types.Log // Blockchain specific contextual infos
}
type ETradingProceedsClaimed struct {//0x95366b7f64c6bb45149f9f7c522403fceebe5170ff76b8ffde2b0ab943ac11ce
	Universe        common.Address
	Shareholder		common.Address
	Market          common.Address
	Outcome         *big.Int
	NumShares       *big.Int
	NumPayoutTokens *big.Int
	Fees            *big.Int
	Timestamp       *big.Int
	Raw             types.Log // Blockchain specific contextual infos
}
type ECompleteSetsPurchased struct {//0xfe06587917de7df83a446bcbb889cee699d7fc35b7b53e263282c2acb5a16499
	Universe        common.Address
	Market          common.Address
	Account         common.Address
	NumCompleteSets *big.Int
	Timestamp       *big.Int
	Raw             types.Log // Blockchain specific contextual infos
}
type ECompleteSetsSold struct {//0xdd7dcfa6708112395eb94e9b1889295fb19af21ef290e918256838c979b2dfbd
	Universe        common.Address
	Market          common.Address
	Account         common.Address
	NumCompleteSets *big.Int
	Fees            *big.Int
	Timestamp       *big.Int
	Raw             types.Log // Blockchain specific contextual infos
}
type EExecuteTransactionStatus struct {// Augur's transaction status
	Success        bool
	FundingSuccess bool
	Raw            types.Log // Blockchain specific contextual infos
}
type ETransactionRelayed struct { // RelayHub event (v1)
	Relay    common.Address		//Topics[1]
	From     common.Address		//Topics[2]
	To       common.Address		//Topics[3]
	Selector [4]byte
	Status   uint8
	Charge   *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}



/// Token events
type ETokensBurned struct {//0x145a4839b3d82d1e28f6ed93f52622b351892e835530386bb1fe4effba99aeea
	Universe    common.Address
	Token       common.Address
	Target      common.Address
	Amount      *big.Int
	TokenType   uint8
	Market      common.Address
	TotalSupply *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}
type ETokensMinted struct {//0x07f766729171db8cc73d96b25cc56784077e26c7ff48b0187877ace391c181a6
	Universe    common.Address
	Token       common.Address
	Target      common.Address
	Amount      *big.Int
	TokenType   uint8
	Market      common.Address
	TotalSupply *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}
type ETokensTransferred struct {//0x3c67396e9c55d2fc8ad68875fc5beca1d96ad2a2f23b210ccc1d986551ab6fdf
	Universe  common.Address
	Token     common.Address
	From      common.Address
	To        common.Address
	Value     *big.Int
	TokenType uint8
	Market    common.Address
	Raw       types.Log // Blockchain specific contextual infos
}
type ETokenBalanceChanged struct {//0x63fd58f559b73fc4da5511c341ec8a7b31c5c48538ef83c6077712b6edf5f7cb
	Universe  common.Address
	Owner     common.Address
	Token     common.Address
	TokenType uint8
	Market    common.Address
	Balance   *big.Int
	Outcome   *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}
type EShareTokenBalanceChanged struct {//0x350ea32dc29530b9557420816d743c436f8397086f98c96292138edd69e01cb3
	Universe common.Address
	Account  common.Address
	Market   common.Address
	Outcome  *big.Int
	Balance  *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}
type ETransferBatch struct {
	Operator common.Address
	From     common.Address
	To       common.Address
	Ids      []*big.Int
	Values   []*big.Int
	Raw      types.Log // Blockchain specific contextual infos
}
type ETransferSingle struct {
	Operator common.Address
	From     common.Address
	To       common.Address
	Id       *big.Int
	Value    *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}
type ETransfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}
type EApprovalForAll struct {
	Owner    common.Address
	Operator common.Address
	Approved bool
	Raw      types.Log // Blockchain specific contextual infos
}
type EApproval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}
type EOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}
type EMarketRepBondTransferred struct {//0x0519ee50d0e6120223e58d0b52824ca4985c524f045a3d6a529936e511d2ba8d.
	Universe common.Address
	Market   common.Address
	From     common.Address
	To       common.Address
	Raw      types.Log // Blockchain specific contextual infos
}


/// Reporting Events
type EDesignatedReportStakeChanged struct {//0x9c75a088fcb0527d67a80a7d0a5006bbabe02f4b23984234ae68b2b146f001bc
	Universe              common.Address
	DesignatedReportStake *big.Int
	Raw                   types.Log // Blockchain specific contextual infos
}
type EInitialReportSubmitted struct {//0xc3ebb227c22e7644e9bef8822009f746a72c86f239760124d67fdc2c302b3115
	Universe             common.Address
	Reporter             common.Address
	Market               common.Address
	InitialReporter      common.Address
	AmountStaked         *big.Int
	IsDesignatedReporter bool
	PayoutNumerators     []*big.Int
	Description          string
	NextWindowStartTime  *big.Int
	NextWindowEndTime    *big.Int
	Timestamp            *big.Int
	Raw                  types.Log // Blockchain specific contextual infos
}
type EInitialReporterTransferred struct {//0xee62c58e2603b92f96a002e012f4f3bd5748102cfa3b711f6d778c6237fcaa96
	Universe common.Address
	Market   common.Address
	From     common.Address
	To       common.Address
	Raw      types.Log // Blockchain specific contextual infos
}
type EInitialReporterRedeemed struct {//0x3ffffb51f92f91faf4ba8c906f5a0180d1033be93b1e227cd92c872dc234fdf0
	Universe         common.Address
	Reporter         common.Address
	Market           common.Address
	InitialReporter  common.Address
	AmountRedeemed   *big.Int
	RepReceived      *big.Int
	PayoutNumerators []*big.Int
	Timestamp        *big.Int
	Raw              types.Log // Blockchain specific contextual infos
}
type EDisputeCrowdsourcerCreated struct {//0xf9a0b30bcf861874bf36630742f0d56b22648898d7cdd0cd785d74acd17e0d44
	Universe            common.Address
	Market              common.Address
	DisputeCrowdsourcer common.Address
	PayoutNumerators    []*big.Int
	Size                *big.Int
	DisputeRound        *big.Int
	Raw                 types.Log // Blockchain specific contextual infos
}
type EDisputeCrowdsourcerContribution struct {//0xe7f47639cdf56ec6c5451df334b73c9ca5cccd20da2c0f4e390e9bb71a6f672a
	Universe            common.Address
	Reporter            common.Address
	Market              common.Address
	DisputeCrowdsourcer common.Address
	AmountStaked        *big.Int
	Description         string
	PayoutNumerators    []*big.Int
	CurrentStake        *big.Int
	StakeRemaining      *big.Int
	DisputeRound        *big.Int
	Timestamp           *big.Int
	Raw                 types.Log // Blockchain specific contextual infos
}
type EDisputeCrowdsourcerCompleted struct {//0x81afc41f9f2f0d22a52a2ddb3a0b6db83baf39c05544fd25f2751b72b1943bb5
	Universe               common.Address
	Market                 common.Address
	DisputeCrowdsourcer    common.Address
	PayoutNumerators       []*big.Int
	NextWindowStartTime    *big.Int
	NextWindowEndTime      *big.Int
	PacingOn               bool
	TotalRepStakedInPayout *big.Int
	TotalRepStakedInMarket *big.Int
	DisputeRound           *big.Int
	Timestamp              *big.Int
	Raw                    types.Log // Blockchain specific contextual infos
}
type EDisputeCrowdsourcerRedeemed struct {//0x6afb0328cf957750be87a6f34b1cd21457ddf1382af65f9592ff2d333945633f
	Universe            common.Address
	Reporter            common.Address
	Market              common.Address
	DisputeCrowdsourcer common.Address
	AmountRedeemed      *big.Int
	RepReceived         *big.Int
	PayoutNumerators    []*big.Int
	Timestamp           *big.Int
	Raw                 types.Log // Blockchain specific contextual infos
}
type EDisputeWindowCreated struct {//0x97f8b399e255f30d56b759b645c86652624ee258937579ff4a747abaeae857c4
	Universe      common.Address
	DisputeWindow common.Address
	StartTime     *big.Int
	EndTime       *big.Int
	Id            *big.Int
	Initial       bool
	Raw           types.Log // Blockchain specific contextual infos
}
type EMarketParticipantsDisavowed struct {//0x3b4f3db017516414df2695e5b0052661779d7163a6cd4368fd74313be73fa0b8
	Universe common.Address
	Market   common.Address
	Raw      types.Log // Blockchain specific contextual infos
}
type EReportingParticipantDisavowed struct {//0xb20adf682c8f82b94a135452f54ac4483c9ee8c9b2324e946120696ab1d034b4
	Universe             common.Address
	Market               common.Address
	ReportingParticipant common.Address
	Raw                  types.Log // Blockchain specific contextual infos
}

type EParticipationTokensRedeemed struct {//0x18052b5e29020458e154999fa71891a5db3404a5b0b9c5ec60c90adca7d38d63
	Universe                common.Address
	DisputeWindow           common.Address
	Account                 common.Address
	AttoParticipationTokens *big.Int
	FeePayoutShare          *big.Int
	Timestamp               *big.Int
	Raw                     types.Log // Blockchain specific contextual infos
}
type EReportingFeeChanged struct {//0xadddfaec4505d90a6a211907536944e6e1af7ff5cf6d1873de43e36020f36009
	Universe     common.Address
	ReportingFee *big.Int
	Raw          types.Log // Blockchain specific contextual infos
}


/////////// (AUGUR) ARBITRUM EVENTS
type AMMFactoryPoolCreated struct {
	Pool             common.Address
	MarketFactory    common.Address
	MarketId         *big.Int
	Creator          common.Address
	LpTokenRecipient common.Address
	Raw              types.Log // Blockchain specific contextual infos
}
type TurboCreated struct {
	Id                   *big.Int
	CreatorFee           *big.Int
	OutcomeSymbols       []string
	OutcomeNames         [][32]byte
	NumTicks             *big.Int
	Arbiter              common.Address
	ArbiterConfiguration []byte
	Index                *big.Int
	Raw                  types.Log // Blockchain specific contextual infos
}
type WinningsClaimed struct {
	Id       *big.Int
	Amount   *big.Int
	Receiver common.Address
	Raw      types.Log // Blockchain specific contextual infos
}
type PriceMarketCreated struct {
	Id        *big.Int
	Creator   common.Address
	EndTime   *big.Int
	SpotPrice *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}
type SportsLinkMarketCreated struct {
	Id                *big.Int
	Creator           common.Address
	EndTime           *big.Int
	MarketType        uint8
	EventId           *big.Int
	HomeTeamId        *big.Int
	AwayTeamId        *big.Int
	EstimatedStarTime *big.Int
	Score             *big.Int
	Raw               types.Log // Blockchain specific contextual infos
}
type TrustedMarketCreated struct {
	Id          *big.Int
	Creator     common.Address
	EndTime     *big.Int
	Description string
	Outcomes    []string
	Raw         types.Log // Blockchain specific contextual infos
}
type MarketResolved struct {
	Id     *big.Int
	Winner common.Address
	Raw    types.Log // Blockchain specific contextual infos
}
type SharesMinted struct {
	Id       *big.Int
	Amount   *big.Int
	Receiver common.Address
	Raw      types.Log // Blockchain specific contextual infos
}
type SharesBurned struct {
	Id       *big.Int
	Amount   *big.Int
	Receiver common.Address
	Raw      types.Log // Blockchain specific contextual infos
}
type SharesSwapped struct {
	MarketFactory common.Address
	MarketId      *big.Int
	User          common.Address
	Outcome       *big.Int
	Collateral    *big.Int
	Shares        *big.Int
	Raw           types.Log // Blockchain specific contextual infos
}
type LiquidityChanged struct {
	MarketFactory common.Address
	MarketId      *big.Int
	User          common.Address
	Recipient     common.Address
	Collateral    *big.Int
	LpTokens      *big.Int
	Raw           types.Log // Blockchain specific contextual infos
}

