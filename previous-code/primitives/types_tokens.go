package primitives
import (
)
type DaiB struct {
	Id					int64
	Aid					int64
	DaiTransfId			int64
	BlockNum			int64
	Processed			bool
	Address				string
	Amount				string
	Balance				string
	BlockHash			string
}
type DaiOp struct {
	BlockNum			int64
	Date				string
	Deposit				string
	Withdrawal			string
	FromAddr			string
	ToAddr				string
}
type ERC20B struct {
	Id					int64
	Aid					int64
	ParentId			int64
	BlockNum			int64
	ContractAid			int64
	Processed			bool
	Address				string
	ContractAddr		string
	Amount				string
	Balance				string
	BlockHash			string
}
type TokenVeryShortInfo struct {
	TokenAid			int64
	TokenAddr			string
	Name				string
	Symbol				string
}
type TokProcessStatus struct {
	LastEvtId				int64
}
type ERC20ProcessStatus struct {
	LastEvtId				int64
}
type ERC20ShTokContract struct {
	TimeStamp               int64
	WrapperAid              int64
	LastEvtId               int64
	MarketAid               int64
	OutcomeIdx              int
	Decimals                int
	Address                 string
	Symbol                  string
	Name                    string
	MktAddr                 string
	MktDescr                string
	DateTime                string
	Outcome                 string
}
type ERC20Info struct {
	Id						int64
	Aid						int64
	TotalSupplyF			float64
	Decimals				int
	TotalSupply				string
	Address					string
	Name					string
	Symbol					string
}
type UserShtokTransfer struct {
	TimeStamp				int64
	Amount					float64
	Balance					float64
	FromAid					int64
	ToAid					int64
	From					string
	To						string
	Date					string
}
type UserShTokens struct {
	WrapperAid				int64
	NumTransfers			int64
	MarketAid				int64
	Balance					float64
	OutcomeIdx				int
	Symbol					string
	Name					string
	WrapperAddr				string
	MarketAddr				string
}
