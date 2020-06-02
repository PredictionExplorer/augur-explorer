// Global types, used anywhere in the package
package main
import (
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)
const (
	MAX_BLOCKS_CHAIN_SPLIT = 128
)
type OrderType uint8
const (
	OrderTypeBid		OrderType = 0
	OrderTypeAsk		OrderType = 1
)
type OrderAction uint8
const(
	OrderActionCreate	OrderAction = 0
	OrderActionCancel	OrderAction = 1
	OrderActionFill		OrderAction = 2
)
type TokenType uint8
const(
	ReputationToken		TokenType = 0
	DisputeCrowdsourcer TokenType = 1
	ParticipationToken	TokenType = 2
)
type MarketStatus uint8
const (
	MktStatusTraded		MarketStatus = 0
	MktStatusReporting	MarketStatus = 1
	MktStatusReported	MarketStatus = 2
	MktStatusDisputing	MarketStatus = 3
	MktStatusFinalized	MarketStatus = 4
	MktStatusFinInvalid	MarketStatus = 5
)
type BlockNumber int64	// -1 is used to mark 'block not set' for the database

type EventSequencer struct {	// determines the order for contained events
	unordered_list		[]*types.Log
}
type ExtraInfo struct {
	Categories			[]string	`json:"categories"`
	Description			string		`json:"description"`
	Tags				[]string	`json:"tags"`
	LongDescription		string		`json:"longDescription"`
}
type MarketCreatedEvt struct {
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
type MktOrderEvt struct {
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
type MarketOIChangedEvt struct {
	Universe common.Address
	Market   common.Address
	MarketOI *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}
type MktFinalizedEvt struct {
	Universe                common.Address
	Market                  common.Address
	Timestamp               *big.Int
	WinningPayoutNumerators []*big.Int
	Raw                     types.Log // Blockchain specific contextual infos
}
type InitialReportSubmittedEvt struct {
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
type MktVolumeChangedEvt struct {
	Universe       common.Address
	Market         common.Address
	Volume         *big.Int
	OutcomeVolumes []*big.Int
	Timestamp      *big.Int
	Raw            types.Log // Blockchain specific contextual infos
}
type DisputeCrowdsourcerContributionEvt struct {
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
type TokensTransferred struct {
	Universe  common.Address
	Token     common.Address
	From      common.Address
	To        common.Address
	Value     *big.Int
	TokenType uint8
	Market    common.Address
	Raw       types.Log // Blockchain specific contextual infos
}
type TokenBalanceChanged struct {
	Universe  common.Address
	Owner     common.Address
	Token     common.Address
	TokenType uint8
	Market    common.Address
	Balance   *big.Int
	Outcome   *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}
type ShareTokenBalanceChanged struct {
	Universe common.Address
	Account  common.Address
	Market   common.Address
	Outcome  *big.Int
	Balance  *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}
type CancelZeroXOrder struct {
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
type TransferBatch struct {
	Operator common.Address
	From     common.Address
	To       common.Address
	Ids      []*big.Int
	Values   []*big.Int
	Raw      types.Log // Blockchain specific contextual infos
}
type TransferSingle struct {
	Operator common.Address
	From     common.Address
	To       common.Address
	Id       *big.Int
	Value    *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}
type ProfitLossChanged struct {
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
type Transfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}
type FillEvt struct {
	MakerAddress           common.Address
	FeeRecipientAddress    common.Address
	MakerAssetData         []byte
	TakerAssetData         []byte
	MakerFeeAssetData      []byte
	TakerFeeAssetData      []byte
	OrderHash              [32]byte
	TakerAddress           common.Address
	SenderAddress          common.Address
	MakerAssetFilledAmount *big.Int
	TakerAssetFilledAmount *big.Int
	MakerFeePaid           *big.Int
	TakerFeePaid           *big.Int
	ProtocolFeePaid        *big.Int
	Raw                    types.Log // Blockchain specific contextual infos
}
type OwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}
type TradingProceedsClaimed struct {
	Universe        common.Address
	Sender          common.Address
	Market          common.Address
	Outcome         *big.Int
	NumShares       *big.Int
	NumPayoutTokens *big.Int
	Fees            *big.Int
	Timestamp       *big.Int
	Raw             types.Log // Blockchain specific contextual infos
}
type ApprovalForAll struct {
	Owner    common.Address
	Operator common.Address
	Approved bool
	Raw      types.Log // Blockchain specific contextual infos
}
type Approval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}
type InfoMarket struct {
	MktAid			int64
	Fee				float64
	OpenInterest	float64
	CurVolume		float64
	MoneyAtStake	float64
	TotalTrades		int64
	MktAddr			string
	MktAddrSh		string	// short address (with .. in the middle)
	Signer			string
	SignerSh		string
	MktCreator		string
	MktCreatorSh	string	// short address (with .. in the middle)
	Reporter		string
	ReporterSh		string
	EndDate			string
	Description		string
	LongDesc		string
	Categories		string
	Outcomes		string
	MktType			string
	Status			string
	CurOutcome		string	// calculated only if the query is made on specific outcome
}
type InfoCategories struct {
	CatId			int64
	Category		string
	Subcategories	[]string
}
type MarketTrade struct {
	OrderHash		string
	MktAddr			string
	MktAddrSh		string	// short address (with .. in the middle)
	CreatorAddr		string
	CreatorAddrSh	string	// short address (with .. in the middle)
	FillerAddr		string
	FillerAddrSh	string	// short address (with .. in the middle)
	Type			string
	Direction		string
	Date			string
	Price			float64
	Amount			float64
	Outcome			int
	OutcomeStr		string
}
type OutcomeVol struct {
	Outcome			int
	OutcomeStr		string
	Volume			float64
	LastPrice		float64
	MktType			int
	MktAddr			string
}
type ZxMeshOrderSpec struct {
	Market			common.Address
	Price			*big.Int
	Outcome			uint8
	Type			uint8
}
type DepthEntry struct {
	MktAid			int64
	ExpiresTs		int64
	Price			float64
	Volume			float64
	AccumVol		float64
	TotalBids		int32
	TotalAsks		int32
	TotalCancel		int32
	OutcomeIdx		int32
	WalletAddr		string
	WalletAddrSh	string	// short version of the addr
	EOAAddr			string
	EOAAddrSh		string	// short version of the addr
	DateCreated		string
	Expires			string
}
type MarketDepth struct {
	Bids			[]DepthEntry
	Asks			[]DepthEntry
}
type UserInfo struct {
	EOAAid			int64
	WalletAid		int64
	ProfitLoss		float64 // profit/loss for the (account) lifetime
	TradeFreq		float64	// trade frequency as percentil of all users (ex: top 15% of all users)
	ReportProfits	float64	// amount of money user has made in profits in outcome reporting
	AffProfits		float64	// profits made in affiliate commissions
	MoneyAtStake	float64	// how much money User has invested
	TotalWithdrawn	float64	// amount of money User has deposited
	TotalDeposited	float64	// amount of money User has withdrawn
	TopTrades		float64
	TopProfit		float64
	TotalTrades		uint32	// how many trades were made by this User
	MarketsCreated	uint32	// how many markets this User has created
	MarketsTraded	uint32	// how many markets this User has traded
	WithdrawReqs	uint32	// number of withdrawal requests
	DepositReqs		uint32	// number of Deposit requests
	TotalReports	uint32	// amount of reports User has made
	TotalDesignated	uint32	// total reports submitted as designated reporter
	EOAAddr			string	// User's Ethereum address (Externally Owned Account)
	EOAAddrSh		string	// short version of the above addr
	WalletAddr		string	// User's Wallet Contract Address
	WalletAddrSh	string	// short version of the above addr
}
type MainStats struct {
	MarketsCount	int64
	YesNoCount		int64
	CategCount		int64
	ScalarCount		int64
	ActiveCount		int64
	FinalizedCount	int64
	MoneyAtStake	float64
	TradesCount		int64
}
type MarketOrder struct {	// this is a short order info, to show in tables
	MktAid				int64
	TradeTs				int64
	Price				float64
	Volume				float64
	AccumVol			float64
	OutcomeIdx			int32
	OType				int32
	CreatedTs			int64
	OrderHash			string
	SellerWalletAddr	string
	SellerWalletAddrSh	string	// short version of the addr
	SellerEOAAddr		string
	SellerEOAAddrSh		string	// short version of the addr
	BuyerWalletAddr		string
	BuyerWalletAddrSh	string
	BuyerEOAAddr		string
	BuyerEOAAddrSh		string
	Direction			string
	Date				string
}
type PLEntry struct {	// profit loss entry
	MktAid				int64
	Timestamp			int64
	BlockNum			int64
	NetPosition			float64
	AccumPl				float64
	AccumFrozen			float64
	AvgPrice			float64
	FrozenFunds			float64
	RealizedProfit		float64
	RealizedCost		float64
	FinalProfit			float64
	MktType				int
	OutcomeIdx			int
	Date				string
	EOAAddr				string
	EOAAddrSh			string
	WalletAddr			string
	WalletAddrSh		string
	MktAddr				string
	MktAddrSh			string
	OutcomeStr			string
	MktDescr			string
	OrderHash			string
	CounterPAddr		string
	CounterPAddrSh		string
}
type RankStats struct {
	EoaAid				int64
	TotalTrades			int64
	ProfitLoss			float64
}
type OrderInfo struct {		// this is a full order information, to show in dedicated webpage
	MktAid				int64
	TradeTs				int64
	Price				float64
	Volume				float64
	OutcomeIdx			int32
	CreatedTs			int64
	OrderHashSh			string
	OrderHash			string
	OType				string
	CreatorrWalletAddr	string
	CreatorWalletAddrSh	string	// short version of the addr
	CreatorEOAAddr		string
	CreatorEOAAddrSh	string	// short version of the addr
	FillerWalletAddr	string
	FillerWalletAddrSh	string
	FillerEOAAddr		string
	FillerEOAAddrSh		string
	Date				string
	MarketAddr			string
	MarketAddrSh		string
}
