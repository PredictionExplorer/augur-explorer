package main

import (
	"runtime"
	"fmt"
	"os"
	"io"
	"encoding/hex"
)
func check(e error) {
	if e != nil {
		panic(e)
	}
}
func Fatalf(format string, args ...interface{}) {
	w := io.MultiWriter(os.Stdout, os.Stderr)
	if runtime.GOOS == "windows" {
		// The SameFile check below doesn't work on Windows.
		// stdout is unlikely to get redirected though, so just print there.
		w = os.Stdout
	} else {
		outf, _ := os.Stdout.Stat()
		errf, _ := os.Stderr.Stat()
		if outf != nil && errf != nil && os.SameFile(outf, errf) {
			w = os.Stderr
		}
	}
	fmt.Fprintf(w, "Fatal: "+format+"\n", args...)
	os.Exit(1)
}
func (evt *MarketCreatedEvt) Dump() {	// dumps struct to stdout for debugging
	fmt.Printf("MarketCreated {\n")
	fmt.Printf("\tUniverse: %v\n",evt.Universe.String())
	fmt.Printf("\tEndTime: %v\n",evt.EndTime)
	fmt.Printf("\tExtraInfo: %v\n",evt.ExtraInfo)
	fmt.Printf("\tMarket: %v\n",evt.Market.String())
	fmt.Printf("\tMarketCreator: %v\n",evt.MarketCreator.String())
	fmt.Printf("\tDesignatedReporter: %v\n",evt.DesignatedReporter.String())
	fmt.Printf("\tFeePerCashInAttoCash: %v\n",evt.FeePerCashInAttoCash);
	prices := bigint_ptr_slice_to_str(&evt.Prices,",")
	fmt.Printf("\tPrices: %v\n",prices)
	fmt.Printf("\tMarketType: %v\n",evt.MarketType)
	fmt.Printf("\tNumTicks: %v\n",evt.NumTicks)
	outcomes := outcomes_to_str(&evt.Outcomes,",")
	fmt.Printf("\tOutcomes: %v\n",outcomes)
	fmt.Printf("\tNoShowBond: %v\n",evt.NoShowBond)
	fmt.Printf("\tTimestamp: %v\n",evt.Timestamp)
	fmt.Printf("}\n")
}
func (evt *MarketOIChangedEvt) Dump() {	// dumps struct to stdout for debugging

	fmt.Printf("MarketOIChanged {\n")
	fmt.Printf("\tUniverse: %v\n",evt.Universe.String())
	fmt.Printf("\tMarket: %v\n",evt.Market.String())
	fmt.Printf("\tMarket Open Interest: %v\n",evt.MarketOI.String())
	fmt.Printf("}\n")
}
func (evt *MktOrderEvt) Dump() { // dumps struct to stdout for debugging

	fmt.Printf("OrderEvent {\n")
	fmt.Printf("\tUniverse: %v\n",evt.Universe.String())
	fmt.Printf("\tMarket: %v\n",evt.Market.String())
	fmt.Printf("\tEventType: %v\n",evt.EventType)
	fmt.Printf("\tOrderType: %v\n",evt.OrderType)
	fmt.Printf("\tOrderId: %v\n",hex.EncodeToString(evt.OrderId[:]))
	fmt.Printf("\tTradeGroupId: %v\n",evt.TradeGroupId)
	fmt.Printf("\tAddressData: %v\n",addresses_to_str(&evt.AddressData,","))
	uintdata := bigint_ptr_slice_to_str(&evt.Uint256Data,",")
	fmt.Printf("\tUint256data: %v\n",uintdata)
	fmt.Printf("}\n")
}
func (evt *MktFinalizedEvt) Dump() { // dumps struct to stdout for debugging

	fmt.Printf("MarketFinalizedEvent {\n")
	fmt.Printf("\tUniverse: %v\n",evt.Universe.String())
	fmt.Printf("\tMarket: %v\n",evt.Market.String())
	fmt.Printf("\tTimestamp: %v\n",evt.Timestamp)
	payouts := bigint_ptr_slice_to_str(&evt.WinningPayoutNumerators,",")
	fmt.Printf("\tWinningPayouts: %v\n",payouts)
	fmt.Printf("}\n")
}
func (evt *InitialReportSubmittedEvt) Dump() {

	fmt.Printf("InitialReportSubmitted {\n")
	fmt.Printf("\tUniverse: %v",evt.Universe.String())
	fmt.Printf("\tReporter: %v",evt.Reporter.String())
	fmt.Printf("\tMarket: %v",evt.Market.String())
	fmt.Printf("\tInitialReporter: %v",evt.InitialReporter.String())
	fmt.Printf("\tAmountStaked: %v",evt.AmountStaked)
	fmt.Printf("\tIsDesignatedReporter: %v",evt.IsDesignatedReporter,evt.IsDesignatedReporter)
	payout_numerators := bigint_ptr_slice_to_str(&evt.PayoutNumerators,",")
	fmt.Printf("\tPayoutNumerators: %v",payout_numerators)
	fmt.Printf("\tDescription: %v",evt.Description)
	fmt.Printf("\tNextWindowStartTime: %v",evt.NextWindowStartTime)
	fmt.Printf("\tNextWindowEndTime: %v",evt.NextWindowEndTime)
	fmt.Printf("\tTimestamp: %v",evt.Timestamp)
	fmt.Printf("}\n")
}
func (evt *DisputeCrowdsourcerContributionEvt) Dump() {

	fmt.Printf("DisputeCrowdsourcerContribution {\n")
	fmt.Printf("\tUniverse: %v",evt.Universe.String())
	fmt.Printf("\tReporter: %v",evt.Reporter.String())
	fmt.Printf("\tMarket: %v",evt.Market.String())
	fmt.Printf("\tDisputedCrowdsourcer: %v",evt.DisputeCrowdsourcer.String())
	fmt.Printf("\tAmountStaked: %v",evt.AmountStaked)
	fmt.Printf("\tDescription: %v",evt.Description)
	payout_numerators := bigint_ptr_slice_to_str(&evt.PayoutNumerators,",")
	fmt.Printf("\tPayoutNumerators: %v",payout_numerators)
	fmt.Printf("\tCurrentStake: %v",evt.CurrentStake)
	fmt.Printf("\tStakeRemaining: %v",evt.StakeRemaining)
	fmt.Printf("\tDisputeRound %v",evt.DisputeRound)
	fmt.Printf("\tTimestamp: %v",evt.Timestamp)
	fmt.Printf("}\n")
}
func (evt *MktVolumeChangedEvt) Dump() { // dumps struct to stdout for debugging

	fmt.Printf("MarketVolumeChanged {\n")
	fmt.Printf("\tUniverse: %v\n",evt.Universe.String())
	fmt.Printf("\tMarket: %v\n",evt.Market.String())
	fmt.Printf("\tVolume: %v\n",evt.Volume.String())
	outcome_volumes := bigint_ptr_slice_to_str(&evt.OutcomeVolumes,",")
	fmt.Printf("\tOutcomeVolumes: %v\n",outcome_volumes)
	fmt.Printf("\tTimestamp: %v\n",evt.Timestamp)
	fmt.Printf("}\n")
}
func (evt *TokensTransferred) Dump() {

	fmt.Printf("TokensTransferred {\n")
	fmt.Printf("\tUniverse: %v\n",evt.Universe.String())
	fmt.Printf("\tToken: %v\n",evt.Token.String())
	fmt.Printf("\tFrom: %v\n",evt.From.String())
	fmt.Printf("\tTo: %v\n",evt.To.String())
	fmt.Printf("\tValue: %v\n",evt.Value.String())
	fmt.Printf("\tTT:TokenType: %v\n",evt.TokenType)
	fmt.Printf("\tMarket: %v\n",evt.Market.String())
	fmt.Printf("}\n")
}
func (evt *TokenBalanceChanged) Dump() {

	fmt.Printf("TokensBalanceChanged {\n")
	fmt.Printf("\tUniverse: %v\n",evt.Universe.String())
	fmt.Printf("\tOwner: %v\n",evt.Owner.String())
	fmt.Printf("\tToken: %v\n",evt.Token.String())
	fmt.Printf("\tTBC:TokenType: %v\n",evt.TokenType)
	fmt.Printf("\tMarket: %v\n",evt.Market.String())
	fmt.Printf("\tBalance: %v\n",evt.Balance.String())
	fmt.Printf("\tOutcome: %v\n",evt.Outcome.String())
	fmt.Printf("}\n")
}
func (evt *ShareTokenBalanceChanged) Dump() {
	fmt.Printf("ShareTokensBalanceChanged {\n")
	fmt.Printf("\tUniverse: %v\n",evt.Universe.String())
	fmt.Printf("\tAccount: %v\n",evt.Account.String())
	fmt.Printf("\tMarket: %v\n",evt.Market.String())
	fmt.Printf("\tOutcome: %v\n",evt.Outcome.String())
	fmt.Printf("\tBalance: %v\n",evt.Balance.String())
	fmt.Printf("}\n")
}
func (evt *CancelZeroXOrder) Dump() {
	fmt.Printf("CancelZeroXOrder {\n")
	fmt.Printf("\tUniverse: %v\n",evt.Universe.String())
	fmt.Printf("\tMarket: %v\n",evt.Market.String())
	fmt.Printf("\tAccount: %v\n",evt.Account.String())
	fmt.Printf("\tOutcome: %v\n",evt.Outcome.String())
	fmt.Printf("\tPrice: %v\n",evt.Price.String())
	fmt.Printf("\tAmount: %v\n",evt.Amount.String())
	fmt.Printf("\tOrderType: %v\n",evt.OrderType)
	fmt.Printf("\tOrderHash: %v\n",hex.EncodeToString(evt.OrderHash[:]))
	fmt.Printf("\t\n")
}
func (evt *TransferBatch) Dump() {
	fmt.Printf("TransferBatch {\n")
	fmt.Printf("\tOperator: %v\n",evt.Operator.String())
	fmt.Printf("\tFrom: %v\n",evt.From.String())
	fmt.Printf("\tTo: %v\n",evt.To.String())
	ids := bigint_ptr_slice_to_str(&evt.Ids,",")
	fmt.Printf("\tIds: %v\n",ids)
	values := bigint_ptr_slice_to_str(&evt.Values,",")
	fmt.Printf("\tValues: %v\n",values)
	fmt.Printf("}\n")
}
func (evt *TransferSingle) Dump() {
	fmt.Printf("TransferSingle {\n")
	fmt.Printf("\tOperator: %v\n",evt.Operator.String())
	fmt.Printf("\tFrom: %v\n",evt.From.String())
	fmt.Printf("\tTo: %v\n",evt.To.String())
	fmt.Printf("\tId: %v\n",evt.Id.String())
	fmt.Printf("\tValue: %v\n",evt.Value.String())
	fmt.Printf("}\n")
}
func (evt *ProfitLossChanged) Dump() {
	fmt.Printf("ProfitLossChanged {\n")
	fmt.Printf("\tUniverse: %v\n",evt.Universe.String())
	fmt.Printf("\tMarket: %v\n",evt.Market.String())
	fmt.Printf("\tAccount: %v\n",evt.Account.String())
	fmt.Printf("\tOutcome: %v\n",evt.Outcome.String())
	fmt.Printf("\tNetPosition: %v\n",evt.NetPosition.String())
	fmt.Printf("\tAvgPrice: %v\n",evt.AvgPrice.String())
	fmt.Printf("\tRealizedProfit: %v\n",evt.RealizedProfit.String())
	fmt.Printf("\tFrozenFunds: %v\n",evt.FrozenFunds.String())
	fmt.Printf("\tRealizedCost: %v\n",evt.RealizedCost.String())
	fmt.Printf("\tTimestamp: %v\n",evt.Timestamp.String())
	fmt.Printf("}\n")
}
func (evt *Transfer) Dump() {
	fmt.Printf("Transfer {\n")
	fmt.Printf("\tFrom: %v\n",evt.From.String())
	fmt.Printf("\tTo: %v\n",evt.To.String())
	fmt.Printf("\tValue: %v\n",evt.Value.String())
	fmt.Printf("}")
}
