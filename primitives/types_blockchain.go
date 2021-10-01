package primitives
import (
)
type AgtxInBlock struct {
	TxId				int64
	BlockNum			int64
	ContextAid			int64
	MarketAid			int64
	PoolAid				int64
	PairAid				int64
	TimeStamp			int64
	TxType				int
	ContextAddr			string		// Address related to transaction type
	MarketAddr			string		// Market address related to tx
	TxHash				string
	OrderHash			string		// if available, hash of the Order
	PoolAddr			string
	PairAddr			string
}
type BlockInfo struct {
	BlockNumFrom		int64
	BlockNumTo			int64
	FromTimeStamp		int64
	ToTimeStamp			int64
	NumBlocks			int64
	NumAugurTx			int64		// Only Augur-related transaction counter
	NumEvents			int64
	NumAugurEvents		int64
	NumDefiEvents		int64
	NumOtherEvents		int64
	NumBalSwaps			int64		// Num swaps at Balancer
	NumUniSwaps			int64		// Num swaps at Uniswap
	NumMarketsTraded	int64
	NumMarketsCreated	int64
	NumUniqueAddresses	int64
	NumUniqueOrders		int64
	GasUsed				int64
	TxCostEth			float64
	FromDate			string
	ToDate				string
	ActiveAddresses		[]string	//list of addresses participated in this block
	Transactions		[]string
	Orders				[]string
	MarketsTraded		[]MarketVeryShortInfo // list of market addresses created at this block
	MarketsCreated		[]MarketVeryShortInfo
//	BlockTransactions	[]AgtxInBlock	DISCONTINUED

}
type AgtxEvent struct {
	EvtType				int
	DeFiPlatformCode	int
	ReferenceId			int64	// Could be Market Order ID, or Event Log id
	Aid					int64
	MktAid				int64
	DeFiSwapId			int64
	Addr				string
	MktAddr				string
	OrderHash			string
	MktDescr			string
}
type TxInfo struct {
	TotalEvents			int
	NumAugurEvents		int
	NumDeFiEvents		int
	NumOtherEvents		int
	NumBalancerSwaps	int
	NumUniswapSwaps		int
	TxId				int64
	GasUsed				int64
	BlockNum			int64
	FromAid				int64
	ToAid				int64
	Value				float64
	TxFeeEth			float64
	Hash				string
	From				string
	To					string
	BalancerSwaps		[]BalancerSwap
	UniswapSwaps		[]UniswapSwap
	BalancerPools		[]PoolVeryShortInfo
	UniswapPairs		[]PairVeryShortInfo
	TokensTraded		[]TokenVeryShortInfo
	MarketsTraded		[]MarketVeryShortInfo
	FullEventList		[]AgtxEvent
}
type EvtLogEntry struct { // Layer1 entry (event)
	BlockNum				int64
	TxId					int64
	EvtId					int64
//	TxHash					*string
}
type ChainReorg struct {
	BlockNum				int64
	Hash					string
}
type BasicChainInfo struct { // piece of common information for storing in tables
	BlockNum				int64
	TxId					int64
	EvtId					int64
	TimeStamp				int64
}
type AddressInfo struct {
	Aid						int64
	Addr					string
}
