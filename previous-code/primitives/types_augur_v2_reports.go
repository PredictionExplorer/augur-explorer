package primitives
import (
)
type Report struct {
	MktAid				int64
	TimeStamp			int64
	RepStake			float64
	Round				int
	OutcomeIdx			int
	MktType				int
	RType				int		// Report Type : 0-Initial, 1-Contribution
	IsInitial			bool
	IsDesignated		bool
	Reporter			string
	MktAddr				string
	MktAddrSh			string
	MktDescription		string
	OutcomeStr			string
	Date				string
	ReportType			string
	WinStart			string
	WinEnd				string
}
type InitialReportInfo struct {
	InitialReporterAid		int64
	ActualReporterAid		int64
	TimeStamp				int64
	WinStartTs				int64
	WinEndTs				int64
	ScalarValue				float64
	AmountStaked			float64
	MktType					int
	OutcomeIdx				int
	IsDesignated			bool
	OutcomeStr				string
	InitialReporterAddr		string
	ActualReporterAddr		string	// in case someone has to do the job of InitialReporter
	TxHash					string
	TxHashSh				string
	DateTime				string
	WinStartDate			string
	WinEndDate				string
}
type CrowdsourcerInfo struct {
	CrowdsourcerAid			int64
	Size					float64
	Round					int
	OutcomeIdx				int
	OutcomeStr				string
	CrowdsourcerAddr		string
	TxHash					string
	TxHashSh				string
}
type DisputeContribution struct {
	TimeStamp				int64
	ReporterAid				int64
	AmountStaked			float64
	CurrentStake			float64
	StakeRemaining			float64
	ScalarValue				float64
	MktType					int
	OutcomeIdx				int
	DisputeRound			int
	CrowdsourcerAddr		string
	ReporterAddr			string
	ReporterAddrSh			string
	OutcomeStr				string
	TxHash					string
	TxHashSh				string
	DateTime				string
}
type DisputeInfo struct{	
	// A dispute can be created, but not yet completed
	CrowdsourcerAid				int64
	CreatedTs					int64	// Dispute created timetamp
	CompletedTs					int64	// Dispute completed timestamp
	ReporterAid					int64
	DisputeWindowAid			int64
	WindowStartTs				int64
	WindowEndTs					int64
	MinDisputeSize				float64
	TotalRepPayout				float64
	RepInMarket					float64
	ScalarValue					float64
	MktType						int
	OutcomeIdx					int
	DisputeRoundStart			int
	DisputeRoundEnd				int
	PacingOn					bool	// temporal halt on all reporting (if true)
	Completed					bool
	CrowdsourcerAddr			string
	ReporterAddr				string
	OutcomeStr					string
	CreatedTxHash				string
	CreatedTxHashSh				string
	CompletedTxHash				string
	CompletedTxHashSh			string
	CreatedDate					string
	CompletedDate				string
	WindowStartDate				string
	WindowEndDate				string
	Contributions				[]DisputeContribution
}
type ReportingStatus struct {
	TentativeWinningOutcome			string
	InitialReport					InitialReportInfo
	Disputes						[]DisputeInfo

}
type DisputeRound struct {
	TimeStamp				int64
	WindowStartTs			int64
	WindowEndTs				int64
	DisputeWinAid			int64
	CompletedTs				int64
	RepPayout				float64 // collected REP (sum of all contributions) for this round
	MarketRep				float64	// accumulated REP amount
	MinDisputeSize			float64
	ScalarValue				float64
	MktType					int
	OutcomeIdx				int
	RoundNum				int
	Completed				bool
	PacingOn				bool	// temporal halt on all reporting (if true)
	Color					bool	// true if highlite the row
	DateTime				string
	OutcomeStr				string
	WindowStartDate			string
	WindowEndDate			string
	CompletedDate			string
}
type OutcomeRounds struct {
	MarketRep				float64
	TimeStamp				int64
	WindowStartTs			int64
	WindowEndTs				int64
	CompletedTs				int64
	ScalarValue				float64
	MktType					int
	WindowNum				int		// consecutive window number (local to current report)
	WindowSpan				int		// how many rounds does the window span
	RoundNum				int
	Completed				bool
	DateTime				string
	WindowStartDate			string
	WindowEndDate			string
	CompletedDate			string
	ORounds					[]DisputeRound
}
type RoundsRow struct {
	Rounds					OutcomeRounds
}
type IniRepRedeemed struct {
	ReporterAid				int64
	InitialReporterAid		int64
	TimeStamp				int64
	Amount					float64
	RepReceived				float64
	ScalarValue				float64
	MktType					int
	OutcomeIdx				int
	OutcomeStr				string
	DateTime				string
	ReporterAddr			string
	InitialReporterAddr		string
	TxHash					string
	TxHashSh				string
}
type RedeemedParticipant struct {
	ReporterAid				int64
	TimeStamp				int64
	MktType					int
	OutcomeIdx				int
	RepInvested				float64
	RepReturned				float64
	Profit					float64
	ScalarValue				float64
	OutcomeStr				string
	DateTime				string
	ReporterAddr			string
	TxHash					string
	TxHashSh				string
}
type UserRepProfit struct {
	TimeStamp				int64
	MarketAid				int64
	OutcomeIdx				int
	RType					int		// report type : 0-Initial, 1-Contribution
	RepInvested				float64
	RepReturned				float64
	Profit					float64
	ROI						float64
	MarketAddr				string
	MarketDescr				string
	OutcomeStr				string
	DateTime				string
	ReporterAddr			string
	TxHash					string
	TxHashSh				string
}
type RepLosingParticipant struct {
	ReporterAid				int64
	TimeStamp				int64
	OutcomeIdx				int
	RepLost					float64
	OutcomeStr				string
	DateTime				string
	ReporterAddr			string
	TxHash					string
	TxHashSh				string
}
type NoShowBondPrice struct {
	TimeStamp				int64
	Price					float64
	DateTime				string
}
