// Types related to Augur contracts
package primitives
import (
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)
type ZxMeshOrderSpec struct {
	Market			common.Address
	Price			*big.Int
	Outcome			uint8
	Type			uint8
}
type EMarketCreated struct {
	Universe             common.Address
	EndTime              *big.Int
	ExtraInfo            string
	Market               common.Address
	MarketCreator        common.Address
	DesignatedReporter   common.Address
	FeePerCashInAttoCash *big.Int
	Prices               []*big.Int
	MarketType           uint8
	NumTicks             *big.Int
	Outcomes             [][32]byte
	NoShowBond           *big.Int
	Timestamp            *big.Int
	Raw                  types.Log // Blockchain specific contextual infos
}
type EOrderEvent struct {
	Universe     common.Address
	Market       common.Address
	EventType    uint8
	OrderType    uint8
	OrderId      [32]byte
	TradeGroupId [32]byte
	AddressData  []common.Address
	Uint256Data  []*big.Int
	Raw          types.Log // Blockchain specific contextual infos
}
type EMarketOIChanged struct {
	Universe common.Address
	Market   common.Address
	MarketOI *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}
type EMarketFinalized struct {
	Universe                common.Address
	Market                  common.Address
	Timestamp               *big.Int
	WinningPayoutNumerators []*big.Int
	Raw                     types.Log // Blockchain specific contextual infos
}
type EInitialReportSubmitted struct {
	Universe             common.Address
	Reporter             common.Address
	Market               common.Address
	InitialReporter      common.Address
	AmountStaked         *big.Int
	IsDesignatedReporter bool
	PayoutNumerators     []*big.Int
	Description          string
	NextWindowStartTime  *big.Int
	NextWindowEndTime    *big.Int
	Timestamp            *big.Int
	Raw                  types.Log // Blockchain specific contextual infos
}
type EMarketVolumeChanged_v1 struct {	// previous version of the event (to be deleted on Augur Release)
	Universe       common.Address
	Market         common.Address
	Volume         *big.Int
	OutcomeVolumes []*big.Int
	Timestamp      *big.Int
	Raw            types.Log // Blockchain specific contextual infos
}
type EMarketVolumeChanged_v2 struct {
	Universe       common.Address
	Market         common.Address
	Volume         *big.Int
	OutcomeVolumes []*big.Int
	TotalTrades    *big.Int
	Timestamp      *big.Int
	Raw            types.Log // Blockchain specific contextual infos
}
type EDisputeCrowdsourcerContribution struct {
	Universe            common.Address
	Reporter            common.Address
	Market              common.Address
	DisputeCrowdsourcer common.Address
	AmountStaked        *big.Int
	Description         string
	PayoutNumerators    []*big.Int
	CurrentStake        *big.Int
	StakeRemaining      *big.Int
	DisputeRound        *big.Int
	Timestamp           *big.Int
	Raw                 types.Log // Blockchain specific contextual infos
}
type ETokensTransferred struct {
	Universe  common.Address
	Token     common.Address
	From      common.Address
	To        common.Address
	Value     *big.Int
	TokenType uint8
	Market    common.Address
	Raw       types.Log // Blockchain specific contextual infos
}
type ETokenBalanceChanged struct {
	Universe  common.Address
	Owner     common.Address
	Token     common.Address
	TokenType uint8
	Market    common.Address
	Balance   *big.Int
	Outcome   *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}
type EShareTokenBalanceChanged struct {
	Universe common.Address
	Account  common.Address
	Market   common.Address
	Outcome  *big.Int
	Balance  *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}
type ECancelZeroXOrder struct {
	Universe  common.Address
	Market    common.Address
	Account   common.Address
	Outcome   *big.Int
	Price     *big.Int
	Amount    *big.Int
	OrderType uint8
	OrderHash [32]byte
	Raw       types.Log // Blockchain specific contextual infos
}
type ETransferBatch struct {
	Operator common.Address
	From     common.Address
	To       common.Address
	Ids      []*big.Int
	Values   []*big.Int
	Raw      types.Log // Blockchain specific contextual infos
}
type ETransferSingle struct {
	Operator common.Address
	From     common.Address
	To       common.Address
	Id       *big.Int
	Value    *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}
type EProfitLossChanged struct {
	Universe       common.Address
	Market         common.Address
	Account        common.Address
	Outcome        *big.Int
	NetPosition    *big.Int
	AvgPrice       *big.Int
	RealizedProfit *big.Int
	FrozenFunds    *big.Int
	RealizedCost   *big.Int
	Timestamp      *big.Int
	Raw            types.Log // Blockchain specific contextual infos
}
type ETransfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}
type EFill struct {
	MakerAddress           common.Address
	FeeRecipientAddress    common.Address
	MakerAssetData         []byte
	TakerAssetData         []byte
	MakerFeeAssetData      []byte
	TakerFeeAssetData      []byte
	OrderHash              [32]byte
	TakerAddress           common.Address
	SenderAddress          common.Address
	MakerAssetFilledAmount *big.Int
	TakerAssetFilledAmount *big.Int
	MakerFeePaid           *big.Int
	TakerFeePaid           *big.Int
	ProtocolFeePaid        *big.Int
	Raw                    types.Log // Blockchain specific contextual infos
}
type EOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}
type ETradingProceedsClaimed struct {
	Universe        common.Address
	Sender          common.Address
	Market          common.Address
	Outcome         *big.Int
	NumShares       *big.Int
	NumPayoutTokens *big.Int
	Fees            *big.Int
	Timestamp       *big.Int
	Raw             types.Log // Blockchain specific contextual infos
}
type EApprovalForAll struct {
	Owner    common.Address
	Operator common.Address
	Approved bool
	Raw      types.Log // Blockchain specific contextual infos
}
type EApproval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}
type EExecuteTransactionStatus struct {// Augur's transaction status
	Success        bool
	FundingSuccess bool
	Raw            types.Log // Blockchain specific contextual infos
}
