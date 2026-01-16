package randomwalk

// =====================================================================
// User Statistics Types (for leaderboards)
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

