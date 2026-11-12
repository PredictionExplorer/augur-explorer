// API row types for the ChaosZero read endpoints (internal/api/chaoszero).
// CST amounts are returned both scaled to whole CST (float64, for display)
// and raw where exactness matters; timestamps are unix seconds plus a
// formatted date string, following the RandomWalk API conventions.
package chaoszero

// EventRow carries the columns every per-event list row shares.
type EventRow struct {
	RecordId  int64
	EvtLogId  int64
	BlockNum  int64
	TxId      int64
	TxHash    string
	TimeStamp int64
	DateTime  string
}

// ApiRoundSummary is one row of the combined per-round list: lifecycle,
// funding, trading and settlement facts aggregated from several cz_ tables.
type ApiRoundSummary struct {
	RoundId       int64
	FirstSeenTs   int64
	FirstSeenDate string
	NumFundings   int64
	FundedTotal   float64
	NumTrades     int64
	Volume        float64
	NumTraders    int64
	NumSells      int64   // early exits: sells are only possible while the round is open
	SoldTotal     float64 // CST withdrawn early through those sells
	Settled       bool
	ValuesMask    int16
	Carried       float64
	EscrowedTotal float64
	NumClaims     int64
	ClaimedTotal  float64
}

// ApiRoundDetail is the combined single-round view.
type ApiRoundDetail struct {
	Summary    ApiRoundSummary
	Odds       []ApiQuestionOdds
	Fundings   []ApiRoundFunded
	Settlement *ApiRoundSettled
	Escrow     *ApiCycleEscrowed
	Answers    []ApiAnswersObserved
	Claims     []ApiClaim
	Payouts    []ApiPaid
}

// ApiQuestionOdds is one of the eight per-question implied-odds rows of a
// round detail. The price of the question's most recent trade (CST paid or
// received per position unit) is read as the market-implied probability at
// that moment: a winning unit always pays exactly 1 CST, so buying YES at
// 0.66 CST/unit implies a 66% YES. Questions without trades have
// NumTrades = 0 and zero-valued price fields.
type ApiQuestionOdds struct {
	Question   int16
	NumTrades  int64
	LastYes    bool    // side of the most recent trade
	LastBuy    bool    // direction of the most recent trade
	LastPrice  float64 // CST per position unit in that trade
	ImpliedYes float64 // P(YES) implied by that price, in [0,1]
	LastTs     int64
	LastDate   string
}

// ApiRoundFunded is one cz_round_funded row.
type ApiRoundFunded struct {
	EventRow
	RoundId       int64
	CallerAddr    string
	Donated       float64
	TransfersUsed float64
	CarryUsed     float64
	Allocations   []float64
}

// ApiTrade is one cz_trade row.
type ApiTrade struct {
	EventRow
	RoundId  int64
	UserAddr string
	Question int16
	Yes      bool
	Buy      bool
	Cash     float64
	Units    float64
}

// ApiPositionSplit is one cz_position_split row.
type ApiPositionSplit struct {
	EventRow
	RoundId  int64
	UserAddr string
	Question int16
	Amount   float64
}

// ApiPositionsMerged is one cz_positions_merged row.
type ApiPositionsMerged struct {
	EventRow
	RoundId       int64
	UserAddr      string
	Amounts       []float64
	CollateralOut float64
}

// ApiRoundSettled is one cz_round_settled row.
type ApiRoundSettled struct {
	EventRow
	RoundId    int64
	ValuesMask int16
	Answers    []bool
	Escrowed   []float64
	Carried    float64
}

// ApiClaim is one cz_claim row.
type ApiClaim struct {
	EventRow
	RoundId  int64
	UserAddr string
	Amounts  []float64
	Total    float64
}

// ApiAnswersObserved is one cz_answers_observed row.
type ApiAnswersObserved struct {
	EventRow
	RoundId    int64
	KnownMask  int16
	ValuesMask int16
}

// ApiCycleEscrowed is one cz_cycle_escrowed row (vault).
type ApiCycleEscrowed struct {
	EventRow
	RoundId int64
	Amounts []float64
}

// ApiPaid is one cz_paid row (vault).
type ApiPaid struct {
	EventRow
	RoundId       int64
	Question      int16
	RecipientAddr string
	Amount        float64
}

// ApiPositionTransfer is one cz_position_transfer row.
type ApiPositionTransfer struct {
	EventRow
	BatchIndex   int32
	OperatorAddr string
	FromAddr     string
	ToAddr       string
	PositionId   string
	RoundId      int64
	Question     int16
	Yes          bool
	Amount       float64
}

// ApiApprovalForAll is one cz_approval_for_all row.
type ApiApprovalForAll struct {
	EventRow
	OwnerAddr    string
	OperatorAddr string
	Approved     bool
}

// ApiTraderSummary aggregates one address's market cash flows across every
// money-moving cz_ table (trades, splits, merges, claims). Profit is pure
// realized cash accounting — CST received back from the market minus CST
// paid in; positions still held in unsettled rounds are not valued, and a
// position received by ERC-1155 transfer counts as zero cost.
type ApiTraderSummary struct {
	UserAddr    string
	NumRounds   int64 // distinct rounds with any activity
	FirstTs     int64
	FirstDate   string
	LastTs      int64
	LastDate    string
	NumBuys     int64
	BuyVolume   float64 // CST paid buying positions
	NumSells    int64
	SellVolume  float64 // CST withdrawn early by selling to the pool
	NumSplits   int64
	SplitVolume float64 // CST deposited minting YES+NO pairs
	NumMerges   int64
	MergeVolume float64 // CST redeemed merging pairs back
	NumClaims   int64
	ClaimVolume float64 // CST claimed after settlement
	TotalIn     float64 // buys + splits
	TotalOut    float64 // sells + merges + claims
	Profit      float64 // TotalOut - TotalIn (realized)
}

// ApiStatistics is the combined dashboard summary over every cz_ table.
type ApiStatistics struct {
	NumRounds       int64
	LatestRound     int64
	NumFundings     int64
	FundedTotal     float64
	NumTrades       int64
	Volume          float64
	NumTraders      int64
	NumSells        int64   // early exits (see ApiRoundSummary.NumSells)
	SoldTotal       float64 // CST withdrawn early through sells
	NumSplits       int64
	NumMerges       int64
	NumSettled      int64
	CarriedTotal    float64
	NumClaims       int64
	ClaimedTotal    float64
	NumPayouts      int64
	PaidTotal       float64
	NumTransfers    int64
	NumApprovals    int64
	NumAnswerEvents int64
}
