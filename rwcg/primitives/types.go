// Package primitives - Core types used across the RWCG project
package primitives

import (
	"github.com/ethereum/go-ethereum/common"
)

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
