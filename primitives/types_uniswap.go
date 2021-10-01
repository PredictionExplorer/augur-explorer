package primitives
import (
	"github.com/ethereum/go-ethereum/common"
)
type UniswapStatus struct {
	LastEvtId				int64
}
type MarketUPair struct { // Uniswap Pair where the Market can be traded
	MktAid					int64
	OutcomeIdx				int64
	PairAid					int64
	Token0Aid				int64
	Token1Aid				int64
	TotalSwaps				int64
	CreatedTs				int64
	NumAugurTokens			int64
	Token0Decimals			int
	Token1Decimals			int
	Outcome					string
	MktAddr					string
	CreatedDate				string
	PairAddr				string
	Token0Addr				string
	Token1Addr				string
	Token0Name				string
	Token1Name				string
	Token0Symbol			string
	Token1Symbol			string
}
type UniswapSwap struct {
	Id						int64
	PairAid					int64
	BlockNum				int64
	RequesterAid			int64
	CreatedTs				int64
	Amount0_In				float64
	Amount1_In				float64
	Amount0_Out				float64
	Amount1_Out				float64
	CreatedDate				string
	RequesterAddr			string
	PairAddr				string
	Symbol0					string
	Symbol1					string
}
type UserUniswapSwap struct {
	Id                      int64
	PairAid                 int64
	BlockNum                int64
	RequesterAid            int64
	CreatedTs               int64
	MktAid                  int64
	Amount0_In              float64
	Amount1_In              float64
	Amount0_Out             float64
	Amount1_Out             float64
	OutcomeIdx              int
	CreatedDate             string
	RequesterAddr           string
	PairAddr                string
	Symbol0                 string
	Symbol1                 string
	Name0                   string
	Name1                   string
	MktAddr                 string
	MktDescription			string
	Outcome                 string
}
type UPairTokens struct {
	Decimals0				int
	Decimals1				int
	Token0Addr				common.Address
	Token1Addr				common.Address
}
type UPairPrice struct {
	Id						int64
	TimeStamp				int64
	NumRecords				int64
	Price					float64
	Date					string
}
type PairVeryShortInfo struct {
	PairAid				int64
	TotalSwaps			int64
	PairAddr			string
}
