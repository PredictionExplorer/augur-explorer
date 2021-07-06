package primitives
import (

	"github.com/ethereum/go-ethereum/common"
)
type AMM_TxId_Rec struct { // shares minted entry (short)
	RecordId				int64
	TxId					int64
	SharesSwappedId			int64
	LiquidityId				int64
	BalancerId				int64
}
type AMM_TxBalSwaps struct { // Tx ids for balancer swaps
	RecordId				int64
	TxId					int64
	SharesSwappedId			int64
	LiquidityId				int64
}
type AMM_CatSport struct {
	SportId					int64
	Name					string
	Categories				[]string
}
type AMM_CatEntries = map[int64]AMM_CatSport
type AMM_TeamJSON struct {
	TeamId					int64 	`json:"team_id" `
	SportId					string  `json:"sport_id", string`
	Name					string	`json:"name", string`
	Mascot					string	`json:"mascot", string`
	Abbreviation			string	`json:"abbreviation", string`
	Record					string	`json:"record", string`
}
type AMM_Team struct {
	TeamId					int64	`json:"team_id" `
	SportId					int64	`json:"sport_id"`
	Name					string	`json:"name", string`
	Mascot					string	`json:"mascot", string`
	Abbreviation			string	`json:"abbreviation", string`
	Record					string	`json:"record", string`
}
type AMM_TeamEntries = map[int64]AMM_TeamJSON
type AMM_Constants struct {
	Categories				AMM_CatEntries
	Teams					AMM_TeamEntries
}
type AMM_SportMarket struct {
	BlockNum				int64
	CreatedTs				int64
	MarketId				int64
	ContractAid				int64
	FactoryAid				int64
	StartTimeTs				int64
	EndTimeTs				int64
	ResolvedTs				int64
	Score					int64
	MarketTypeCode			int64
	HomeTeamId				int64
	AwayTeamId				int64
	EventId					int64
	Liquidity				float64
	CreatedDate				string
	TxHash					string
	CreatorAddr				string
	FactoryAddr				string
	StartTime				string
	EndTime					string
	HomeTeam				string
	AwayTeam				string
	Title					string
	Description				string
	ResolvedDate			string
	MarketRules				[]string
}
type AMM_ERC20_Op struct {
	EvtId					int64
	BlockNum				int64
	TxId					int64
	TimeStamp				int64
	FromAid					int64
	ToAid					int64
	Amount					float64
	Balance					float64
	Contract				string
	FromAddr				string
	ToAddr					string
	AmountStr				string
	BalanceStr				string
}
type AA_ContractAddrs struct {// Arbitrum-Augur contract addresses
	AMM_Factory				common.Address
	SportsBall1				common.Address
	SportsBall2				common.Address
	MMA						common.Address
	TrustedFactory			common.Address
}
type ArbitrumAugurProcessStatus struct {
	LastEvtId				int64
}
type AA_PoolCreated struct {// Arbitrum Augur PoolCreated event
	EvtId					int64
	BlockNum				int64
	TxId					int64
	TimeStamp				int64
	MarketId				int64
	Contract				string
	PoolAddr				string
	FactoryAddr				string
	CreatorAddr				string
	TokenRecipientAddr		string
}
type AA_PriceMarket struct {// Arbitrum Augur PriceMarket event
	EvtId					int64
	BlockNum				int64
	TxId					int64
	TimeStamp				int64
	EndTime					int64
	MarketId				int64
	Contract				string
	CreatorAddr				string
	SpotPrice				string
}
type AA_SportsMarket struct {// Arbitrum Augur Sports Market Created event
	EvtId					int64
	BlockNum				int64
	TxId					int64
	TimeStamp				int64
	EstimatedStarTime		int64
	EndTime					int64
	MarketId				int64
	EventId					int64
	HomeTeamId				int64
	AwayTeamId				int64
	Score					int64
	MarketType				int
	Contract				string
	CreatorAddr				string
	ShareTokens				string
	CollateralAddr			string
	ProtocolAddr			string
	SettlementAddr			string
	SettlementFee			string
	FeePotAddr				string
	StakerFee				string
	ProtocolFee				string
	ShareFactor				string
}
type AA_TrustedMarket struct {// Arbitrum Augur TrustedMarket Created event
	EvtId					int64
	BlockNum				int64
	TxId					int64
	TimeStamp				int64
	EstimatedStarTime		int64
	EndTime					int64
	MarketId				int64
	EventId					int64
	HomeTeamId				int64
	AwayTeamId				int64
	Score					int64
	MarketType				int
	Contract				string
	CreatorAddr				string
	ShareTokens				string
	CollateralAddr			string
	ProtocolAddr			string
	SettlementAddr			string
	SettlementFee			string
	FeePotAddr				string
	StakerFee				string
	ProtocolFee				string
	ShareFactor				string
	Description				string
	Outcomes				string
}
type AA_TurboCreated struct {// Arbitrum Augur TurboCreated event
	EvtId					int64
	BlockNum				int64
	TxId					int64
	TimeStamp				int64
	TurboId					int64
	NumTicks				int64
	Contract				string
	CreatorFee				string
	OutcomeSymbols		    string // comma separated
	OutcomeNames			string // comma separated
	ArbiterAddr             string
	ArbiterConfiguration	[]byte
	Index					string
}
type AA_SharesMinted struct {// Arbitrum Augur SharesMinted event
	EvtId					int64
	BlockNum				int64
	TxId					int64
	TimeStamp				int64
	MarketId				int64
	Contract				string
	Amount					string
	ReceiverAddr			string
}
type AA_SharesBurned struct {// Arbitrum Augur SharesBurned event
	EvtId					int64
	BlockNum				int64
	TxId					int64
	TimeStamp				int64
	MarketId				int64
	Contract				string
	Amount					string
	ReceiverAddr			string
}
type AA_SharesSwapped struct {
	EvtId					int64
	BlockNum				int64
	TxId					int64
	TimeStamp				int64
	MarketId				int64
	Outcome					int64
	Contract				string
	MarketFactoryAddr		string
	UserAddr				string
	Collateral				string
	Shares					string
}
type AA_SharesSwappedInfo struct {
	EvtId					int64
	BlockNum				int64
	TxId					int64
	CreatedTs				int64
	MarketId				int64
	Outcome					int64
	UserAid					int64
	Collateral				float64
	Shares					float64
	Buy						bool
	Contract				string
	MarketFactoryAddr		string
	UserAddr				string
	InOutRatio				string
	CreatedDate				string
	TxHash					string
	OutcomeStr				string
}
type AA_WinningsClaimed struct {// Arbitrum Augur Winnnings Claimed
	EvtId					int64
	BlockNum				int64
	TxId					int64
	TimeStamp				int64
	MarketId				int64
	Contract				string
	Amount					string
	WinningOutcomeAddr		string
	SettlementFee			string
	Payout					string
	ReceiverAddr			string
}
type AA_MarketResolved struct {// Arbitrum Augur Market Reolved event
	EvtId					int64
	BlockNum				int64
	TxId					int64
	TimeStamp				int64
	MarketId				int64
	Contract				string
	WinnerAddr				string
}
type AA_Pool struct {// Arbitrum Augur Pool 
	BlockNum				int64
	TimeStamp				int64
	PoolAid					int64
	FactoryAid				int64
	MarketId				int64
	DateTime				string
	PoolAddr				string
	FactoryAddr				string
	CreatorAddr				string
	TxHash					string
}
type AA_FeePotTransfer struct { //ERC20 transfer event for FeePot
	EvtId					int64
	BlockNum				int64
	TxId					int64
	TimeStamp				int64
	Contract				string
	From					string
	To						string
	Value					string
}
type AA_Turbo struct {
	BlockNum				int64
	TimeStamp				int64
	CratorFee				float64
	ArbiterAid				int64
	NumTicks				int
	ArbiterAddr				string
}
type AA_LiquidityChanged struct {
	EvtId					int64
	BlockNum				int64
	TxId					int64
	TimeStamp				int64
	MarketId				int64
	Contract				string
	MarketFactoryAddr		string
	UserAddr				string
	RecipientAddr			string
	LpTokens				string
	Collateral				string
	SharesReturned			string
}
type AMM_LiquidityChangedInfo struct {// struct for the API
	CreatedTs				int64
	BlockNum				int64
	MarketId				int64
	UserAid					int64
	Collateral				float64
	Tokens					float64
	In						bool	/// In - true, Out - false
	CreatedDate				string
	TxHash					string
	UserAddr				string
	RecipientAddr			string
}
type AA_SettlementFeeClaimed struct {
	EvtId					int64
	BlockNum				int64
	TxId					int64
	TimeStamp				int64
	Contract				string
	SettlementAddr			string
	Amount					string
	ReceiverAddr			string
}
type AA_ProtocolFeeClaimed struct {
	EvtId					int64
	BlockNum				int64
	TxId					int64
	TimeStamp				int64
	Contract				string
	ProtocolAddr			string
	Amount					string
}
type AA_ProtocolChanged struct {
	EvtId					int64
	BlockNum				int64
	TxId					int64
	TimeStamp				int64
	Contract				string
	ProtocolAddr			string
}
type AA_ProtocolFeeChanged struct {
	EvtId					int64
	BlockNum				int64
	TxId					int64
	TimeStamp				int64
	Contract				string
	Fee						string
}
type AA_SettlementFeeChanged struct {
	EvtId					int64
	BlockNum				int64
	TxId					int64
	TimeStamp				int64
	Contract				string
	Fee						string
}
type AA_StakerFeeChanged struct {
	EvtId					int64
	BlockNum				int64
	TxId					int64
	TimeStamp				int64
	Contract				string
	Fee						string
}
