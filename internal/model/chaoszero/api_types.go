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
	Fundings   []ApiRoundFunded
	Settlement *ApiRoundSettled
	Escrow     *ApiCycleEscrowed
	Answers    []ApiAnswersObserved
	Claims     []ApiClaim
	Payouts    []ApiPaid
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

// ApiStatistics is the combined dashboard summary over every cz_ table.
type ApiStatistics struct {
	NumRounds       int64
	LatestRound     int64
	NumFundings     int64
	FundedTotal     float64
	NumTrades       int64
	Volume          float64
	NumTraders      int64
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
