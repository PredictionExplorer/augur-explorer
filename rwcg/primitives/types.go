// Package primitives - Core types used across the RWCG project
package primitives

import (
	"math/big"

	"github.com/ethereum/go-ethereum/common"
)

// =====================================================================
// Blockchain Transaction Types
// =====================================================================

// AugurTx is a wrapper for Ethereum Transaction object in our own format
type AugurTx struct {
	TxId              int64 // once inserted tx_id is stored here
	BlockNum          int64
	Gas               int64
	GasUsed           int64
	TimeStamp         int64
	TxIndex           int32
	NumLogs           int32
	CtrctCreate       bool
	CumulativeGasUsed uint64
	EffectiveGasPrice *big.Int
	BlockHash         string
	GasPrice          string
	TxHash            string
	From              string
	To                string
	Value             string
	Input             []byte
}

// AddrStatsLog for address statistics logging
type AddrStatsLog struct {
	BlockNum int64
	TxIndex  int64
	Aid      int64 // Address ID
}

// TxShort is a shortened transaction record
type TxShort struct {
	BlockNum int64
	TxIndex  int64
	TxFee    string
}

// =====================================================================
// User Statistics Types
// =====================================================================

// RankStats holds user ranking statistics for leaderboards
type RankStats struct {
	Aid          int64
	TotalTrades  int64
	ProfitLoss   float64
	VolumeTraded float64
}

// ProfitMaker represents a user entry in the profit leaderboard
type ProfitMaker struct {
	Percentage float64
	ProfitLoss float64
	Addr       string
}

// TradeMaker represents a user entry in the trades leaderboard
type TradeMaker struct {
	Percentage  float64
	TotalTrades int64
	Addr        string
}

// VolumeMaker represents a user entry in the volume leaderboard
type VolumeMaker struct {
	Percentage float64
	Volume     float64
	Addr       string
}

// =====================================================================
// User Info Types
// =====================================================================

// UserInfo contains comprehensive information about a user account
type UserInfo struct {
	Aid                int64
	BlockNum           int64
	TimeStamp          int64   // user registration timestamp (from block table)
	ProfitLoss         float64 // profit/loss for the (account) lifetime
	TradeFreq          float64 // trade frequency as percentile of all users
	ReportProfits      float64 // amount of money user has made in profits in outcome reporting
	AffProfits         float64 // profits made in affiliate commissions
	MoneyAtStake       float64 // how much money User has invested
	ValidityBonds      float64 // amount of validity bonds for all markets user created
	TotalWithdrawn     float64 // amount of money User has deposited
	TotalDeposited     float64 // amount of money User has withdrawn
	TopTrades          float64
	TopProfit          float64
	UnclaimedProfit    float64
	WalletAid          int64 // Filled only if present
	EOAAid             int64 // Filled only if present
	BalancerNumSwaps   int64 // statistics: how many swaps at Balancer
	UniswapNumSwaps    int64 // statistics: how many swaps at Uniswap
	HedgingProfits     bool  // Flag to indicate negative 'MoneyAtStake' field
	NoActivity         bool  // True if doesn't have entry in 'ustats' table
	TotalTrades        int32 // how many trades were made by this User
	MarketsCreated     int32 // how many markets this User has created
	MarketsTraded      int32 // how many markets this User has traded
	WithdrawReqs       int32 // number of withdrawal requests
	DepositReqs        int32 // number of Deposit requests
	TotalReports       int32 // amount of reports User has made
	TotalDesignated    int32 // total reports submitted as designated reporter
	Addr               string
	AddrSh             string // short version of the above addr
	WalletAddr         string // Wallet contract address, filled only if present
	EOAAddr            string // EOA address (controlling account) for wallet contract
}

// =====================================================================
// Statistics Types
// =====================================================================

// MainStats contains summary statistics for the main dashboard
type MainStats struct {
	LastBlockNum   int64
	MarketsCount   int64
	YesNoCount     int64
	CategCount     int64
	ScalarCount    int64
	ActiveCount    int64
	FinalizedCount int64
	InvalidCount   int64
	MoneyAtStake   float64
	TradesCount    int64
}

// =====================================================================
// Contract Address Configuration
// =====================================================================

// ContractAddresses stores all the contract addresses used by the system
type ContractAddresses struct {
	ChainId                 int64
	Augur                   common.Address
	AugurTrading            common.Address
	PL                      common.Address
	ZeroxTrade              common.Address
	ZeroxXchg               common.Address
	Dai                     common.Address
	Reputation              common.Address
	WalletReg               common.Address
	WalletReg2              common.Address
	FillOrder               common.Address
	EthXchg                 common.Address
	ShareToken              common.Address
	GenesisUniverse         common.Address
	CreateOrder             common.Address
	LegacyReputationToken   common.Address
	BuyParticipationTokens  common.Address
	RedeemStake             common.Address
	WarpSync                common.Address
	HotLoading              common.Address
	Affiliates              common.Address
	AffiliateValidator      common.Address
	Time                    common.Address
	CancelOrder             common.Address
	Orders                  common.Address
	SimulateTrade           common.Address
	Trade                   common.Address
	OICash                  common.Address
	UniswapV2Factory        common.Address
	UniswapV2Router02       common.Address
	AuditFunds              common.Address
	WETH9                   common.Address
	USDC                    common.Address
	USDT                    common.Address
	RelayHubV2              common.Address
	AccountLoader           common.Address
}

// =====================================================================
// ETL Event Types
// =====================================================================

// EthereumEventLog represents an Ethereum event log entry
type EthereumEventLog struct {
	EvtId           int64
	BlockNum        int64
	TxId            int64
	ContractAid     int64
	TimeStamp       int64
	ContractAddress string
	TxHash          string
	Topic0_Sig      string
	RlpLog          []byte
}

// InspectedEvent for ETL event tracking
type InspectedEvent struct {
	ContractAid int64
	Signature   string
}

// ShortEvtLog is a shortened event log reference
type ShortEvtLog struct {
	EvtId int64
	TxId  int64
}

// EthereumEventTopic represents an event topic
type EthereumEventTopic struct {
	BlockNum    int64
	TxId        int64
	EventLogId  int64
	ContractAid int64
	Pos         int
	Value       string // Important note: this isn't 0x-prefixed
}

// =====================================================================
// Market Types
// =====================================================================

// MarketVeryShortInfo contains minimal market information
type MarketVeryShortInfo struct {
	MktAddr string
	MktDesc string
}

// =====================================================================
// Token Types
// =====================================================================

// ERC20Info contains ERC20 token information
type ERC20Info struct {
	Id           int64
	Aid          int64
	TotalSupplyF float64
	Decimals     int
	TotalSupply  string
	Address      string
	Name         string
	Symbol       string
}

// =====================================================================
// DeFi Types (for TxInfo fields)
// =====================================================================

// BalancerSwap represents a Balancer swap event
type BalancerSwap struct {
	PoolAid    int64
	PoolAddr   string
	TokenIn    string
	TokenOut   string
	AmountIn   float64
	AmountOut  float64
}

// UniswapSwap represents a Uniswap swap event
type UniswapSwap struct {
	PairAid   int64
	PairAddr  string
	Token0    string
	Token1    string
	Amount0   float64
	Amount1   float64
}

// PoolVeryShortInfo contains minimal pool information
type PoolVeryShortInfo struct {
	PoolAid  int64
	PoolAddr string
}

// PairVeryShortInfo contains minimal pair information
type PairVeryShortInfo struct {
	PairAid  int64
	PairAddr string
}

// TokenVeryShortInfo contains minimal token information
type TokenVeryShortInfo struct {
	TokenAid  int64
	TokenAddr string
	Symbol    string
}
