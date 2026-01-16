package primitives
import (
)
type BalancerStatus struct {
	LastEvtId				int64
}
type BalancerPoolInfo struct {
	EvtId					int64
	BlockNum				int64
	TxId					int64
	TimeStamp				int64
	PoolAid					int64
	CallerAid				int64
	NumSwaps				int64
	NumHolders				int64
	NumTokens				int64
	NumAugurTokens			int64
	SwapFee					float64
	PoolAddr				string
	CallerAddr				string
	CreatedDate				string
	ControllerAddr			string
	MktAddr					string
	ContractAddr			string
	Tokens					[]BalancerToken
}
type BalancerToken struct {
	TimeStampAdded			int64
	TokenAid				int64
	Denorm					float64
	Weight					float64
	Balance					float64
	TokenAddr				string
	DateAdded				string
	WrappingContract		ERC20ShTokContract
}
type BalancerJoin struct {
	EvtId					int64
	BlockNum				int64
	TxId					int64
	TimeStamp				int64
	PoolAid					int64
	CallerAid				int64
	TokenAid				int64
	PoolAddr				string
	CallerAddr				string
	TokenInAddr				string
	AmountIn				string
}
type BalancerExit struct {
	EvtId					int64
	BlockNum				int64
	TxId					int64
	TimeStamp				int64
	PoolAid					int64
	CallerAid				int64
	TokenAid				int64
	PoolAddr				string
	CallerAddr				string
	TokenOutAddr			string
	AmountOut				string
}
type BalancerSwap struct {
	Id						int64
	EvtId					int64
	BlockNum				int64
	TxId					int64
	TimeStamp				int64
	PoolAid					int64
	CallerAid				int64
	TokenInAid				int64
	TokenOutAid				int64
	AmountInF				float64
	AmountOutF				float64
	DecimalsIn				int
	DecimalsOut				int
	PoolAddr				string
	CallerAddr				string
	TokenInAddr				string
	TokenOutAddr			string
	SymbolIn				string
	SymbolOut				string
	AmountIn				string
	AmountOut				string
	Date					string
}
type UserBalancerSwap struct {
	Id						int64
	EvtId					int64
	BlockNum				int64
	TxId					int64
	TimeStamp				int64
	PoolAid					int64
	CallerAid				int64
	TokenInAid				int64
	TokenOutAid				int64
	MktAid					int64
	AmountInF				float64
	AmountOutF				float64
	DecimalsIn				int
	DecimalsOut				int
	OutcomeIdx				int
	PoolAddr				string
	CallerAddr				string
	TokenInAddr				string
	TokenOutAddr			string
	SymbolIn				string
	SymbolOut				string
	NameIn					string
	NameOut					string
	AmountIn				string
	AmountOut				string
	Date					string
	MktAddr					string
	MktDescription			string
	Outcome					string
}
type SetSwapFee struct {
	Id						int64
	EvtId					int64
	BlockNum				int64
	TxId					int64
	TimeStamp				int64
	PoolAid					int64
	PoolAddr				string
	FeeStr					string
}
type SetController struct {
	Id						int64
	EvtId					int64
	BlockNum				int64
	TxId					int64
	TimeStamp				int64
	PoolAid					int64
	PoolAddr				string
	ControllerAddr			string
}
type SetPublic struct {
	Id						int64
	EvtId					int64
	BlockNum				int64
	TxId					int64
	TimeStamp				int64
	PoolAid					int64
	PoolAddr				string
	Public					bool
}
type Finalize struct {
	Id						int64
	EvtId					int64
	BlockNum				int64
	TxId					int64
	TimeStamp				int64
	PoolAid					int64
	PoolAddr				string
}
type PoolBind struct {
	Id						int64
	EvtId					int64
	BlockNum				int64
	TxId					int64
	TimeStamp				int64
	PoolAid					int64
	Denorm					string
	PoolAddr				string
	TokenAddr				string
	Balance					string
}
type PoolUnBind struct {
	Id						int64
	EvtId					int64
	BlockNum				int64
	TxId					int64
	TimeStamp				int64
	PoolAid					int64
	PoolAddr				string
	TokenAddr				string
}
type PoolReBind struct {
	Id						int64
	EvtId					int64
	BlockNum				int64
	TxId					int64
	TimeStamp				int64
	PoolAid					int64
	Denorm					string
	PoolAddr				string
	TokenAddr				string
	Balance					string
}
type PoolGulp struct {
	Id						int64
	EvtId					int64
	BlockNum				int64
	TxId					int64
	TimeStamp				int64
	PoolAid					int64
	PoolAddr				string
	TokenAddr				string
	AbsorbedBalance			string
}
type PoolInfo struct {
	PoolAid					int64
	NumHolders				int64
	NumSwaps				int64
	CreatedBlockNum			int64
	WentPublicTs			int64
	CreatedTs				int64
	FinalizedTs				int64
	NumTokens				int
	IsPublic				bool
	WasFinalized			bool
	UsdLiquidity			float64
	SwapFee					float64
	PoolAddr				string
}
type MarketPool struct {
	OutcomeIdx				int
	OutcomeStr				string
	MktAddress				string
}
type BSwapPrice struct {
	Id						int64
	TimeStamp				int64
	NumRecords				int64
	Price					float64
	Date					string
}
type BalancerTokenHolder struct {
	HolderAddr				string
	Balance					float64
	Percentage				float64
}
type PoolVeryShortInfo struct {
	PoolAid				int64
	NumSwaps			int64
	NumHolders			int64
	PoolAddr			string
}
