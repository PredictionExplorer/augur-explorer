package primitives

import (
	"log"
	"encoding/hex"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
)
func (evt *MarketCreatedEvt) Dump(l *log.Logger) {	// dumps struct to stdout for debugging
	l.Printf("MarketCreated {\n")
	l.Printf("\tUniverse: %v\n",evt.Universe.String())
	l.Printf("\tEndTime: %v\n",evt.EndTime)
	l.Printf("\tExtraInfo: %v\n",evt.ExtraInfo)
	l.Printf("\tMarket: %v\n",evt.Market.String())
	l.Printf("\tMarketCreator: %v\n",evt.MarketCreator.String())
	l.Printf("\tDesignatedReporter: %v\n",evt.DesignatedReporter.String())
	l.Printf("\tFeePerCashInAttoCash: %v\n",evt.FeePerCashInAttoCash);
	prices := Bigint_ptr_slice_to_str(&evt.Prices,",")
	l.Printf("\tPrices: %v\n",prices)
	l.Printf("\tMarketType: %v\n",evt.MarketType)
	l.Printf("\tNumTicks: %v\n",evt.NumTicks)
	outcomes := Outcomes_to_str(&evt.Outcomes,",")
	l.Printf("\tOutcomes: %v\n",outcomes)
	l.Printf("\tNoShowBond: %v\n",evt.NoShowBond)
	l.Printf("\tTimestamp: %v\n",evt.Timestamp)
	l.Printf("}\n")
}
func (evt *MarketOIChangedEvt) Dump(l *log.Logger) {	// dumps struct to stdout for debugging

	l.Printf("MarketOIChanged {\n")
	l.Printf("\tUniverse: %v\n",evt.Universe.String())
	l.Printf("\tMarket: %v\n",evt.Market.String())
	l.Printf("\tMarket Open Interest: %v\n",evt.MarketOI.String())
	l.Printf("}\n")
}
func (evt *MktOrderEvt) Dump(l *log.Logger) { // dumps struct to stdout for debugging

	l.Printf("OrderEvent {\n")
	l.Printf("\tUniverse: %v\n",evt.Universe.String())
	l.Printf("\tMarket: %v\n",evt.Market.String())
	l.Printf("\tEventType: %v\n",evt.EventType)
	l.Printf("\tOrderType: %v\n",evt.OrderType)
	l.Printf("\tOrderId: %v\n",hex.EncodeToString(evt.OrderId[:]))
	l.Printf("\tTradeGroupId: %v\n",evt.TradeGroupId)
	l.Printf("\tAddressData: %v\n",addresses_to_str(&evt.AddressData,","))
	uintdata := Bigint_ptr_slice_to_str(&evt.Uint256Data,",")
	l.Printf("\tUint256data: %v\n",uintdata)
	l.Printf("}\n")
}
func (evt *MktFinalizedEvt) Dump(l *log.Logger) { // dumps struct to stdout for debugging

	l.Printf("MarketFinalizedEvent {\n")
	l.Printf("\tUniverse: %v\n",evt.Universe.String())
	l.Printf("\tMarket: %v\n",evt.Market.String())
	l.Printf("\tTimestamp: %v\n",evt.Timestamp)
	payouts := Bigint_ptr_slice_to_str(&evt.WinningPayoutNumerators,",")
	l.Printf("\tWinningPayouts: %v\n",payouts)
	l.Printf("}\n")
}
func (evt *InitialReportSubmittedEvt) Dump(l *log.Logger) {

	l.Printf("InitialReportSubmitted {\n")
	l.Printf("\tUniverse: %v\n",evt.Universe.String())
	l.Printf("\tReporter: %v\n",evt.Reporter.String())
	l.Printf("\tMarket: %v\n",evt.Market.String())
	l.Printf("\tInitialReporter: %v\n",evt.InitialReporter.String())
	l.Printf("\tAmountStaked: %v\n",evt.AmountStaked)
	l.Printf("\tIsDesignatedReporter: %v\n",evt.IsDesignatedReporter)
	payout_numerators := Bigint_ptr_slice_to_str(&evt.PayoutNumerators,",")
	l.Printf("\tPayoutNumerators: %v\n",payout_numerators)
	l.Printf("\tDescription: %v\n",evt.Description)
	l.Printf("\tNextWindowStartTime: %v\n",evt.NextWindowStartTime)
	l.Printf("\tNextWindowEndTime: %v\n",evt.NextWindowEndTime)
	l.Printf("\tTimestamp: %v\n",evt.Timestamp)
	l.Printf("}\n")
}
func (evt *DisputeCrowdsourcerContributionEvt) Dump(l *log.Logger) {

	l.Printf("DisputeCrowdsourcerContribution {\n")
	l.Printf("\tUniverse: %v\n",evt.Universe.String())
	l.Printf("\tReporter: %v\n",evt.Reporter.String())
	l.Printf("\tMarket: %v\n",evt.Market.String())
	l.Printf("\tDisputedCrowdsourcer: %v\n",evt.DisputeCrowdsourcer.String())
	l.Printf("\tAmountStaked: %v\n",evt.AmountStaked)
	l.Printf("\tDescription: %v\n",evt.Description)
	payout_numerators := Bigint_ptr_slice_to_str(&evt.PayoutNumerators,",")
	l.Printf("\tPayoutNumerators: %v\n",payout_numerators)
	l.Printf("\tCurrentStake: %v\n",evt.CurrentStake)
	l.Printf("\tStakeRemaining: %v\n",evt.StakeRemaining)
	l.Printf("\tDisputeRound %v\n",evt.DisputeRound)
	l.Printf("\tTimestamp: %v\n",evt.Timestamp)
	l.Printf("}\n")
}
func (evt *MktVolumeChangedEvt_v1) Dump(l *log.Logger) { // dumps struct to stdout for debugging

	l.Printf("MarketVolumeChanged_v1 {\n")
	l.Printf("\tUniverse: %v\n",evt.Universe.String())
	l.Printf("\tMarket: %v\n",evt.Market.String())
	l.Printf("\tVolume: %v\n",evt.Volume.String())
	outcome_volumes := Bigint_ptr_slice_to_str(&evt.OutcomeVolumes,",")
	l.Printf("\tOutcomeVolumes: %v\n",outcome_volumes)
	l.Printf("\tTimestamp: %v\n",evt.Timestamp)
	l.Printf("}\n")
}
func (evt *MktVolumeChangedEvt_v2) Dump(l *log.Logger) { // dumps struct to stdout for debugging

	l.Printf("MarketVolumeChanged_v2 {\n")
	l.Printf("\tUniverse: %v\n",evt.Universe.String())
	l.Printf("\tMarket: %v\n",evt.Market.String())
	l.Printf("\tVolume: %v\n",evt.Volume.String())
	outcome_volumes := Bigint_ptr_slice_to_str(&evt.OutcomeVolumes,",")
	l.Printf("\tOutcomeVolumes: %v\n",outcome_volumes)
	l.Printf("\tTotalTrades: %v\n",evt.TotalTrades.String())
	l.Printf("\tTimestamp: %v\n",evt.Timestamp)
	l.Printf("}\n")
}
func (evt *TokensTransferred) Dump(l *log.Logger) {

	l.Printf("TokensTransferred {\n")
	l.Printf("\tUniverse: %v\n",evt.Universe.String())
	l.Printf("\tToken: %v\n",evt.Token.String())
	l.Printf("\tFrom: %v\n",evt.From.String())
	l.Printf("\tTo: %v\n",evt.To.String())
	l.Printf("\tValue: %v\n",evt.Value.String())
	l.Printf("\tTT:TokenType: %v\n",evt.TokenType)
	l.Printf("\tMarket: %v\n",evt.Market.String())
	l.Printf("}\n")
}
func (evt *TokenBalanceChanged) Dump(l *log.Logger) {

	l.Printf("TokensBalanceChanged {\n")
	l.Printf("\tUniverse: %v\n",evt.Universe.String())
	l.Printf("\tOwner: %v\n",evt.Owner.String())
	l.Printf("\tToken: %v\n",evt.Token.String())
	l.Printf("\tTBC:TokenType: %v\n",evt.TokenType)
	l.Printf("\tMarket: %v\n",evt.Market.String())
	l.Printf("\tBalance: %v\n",evt.Balance.String())
	l.Printf("\tOutcome: %v\n",evt.Outcome.String())
	l.Printf("}\n")
}
func (evt *ShareTokenBalanceChanged) Dump(l *log.Logger) {
	l.Printf("ShareTokensBalanceChanged {\n")
	l.Printf("\tUniverse: %v\n",evt.Universe.String())
	l.Printf("\tAccount: %v\n",evt.Account.String())
	l.Printf("\tMarket: %v\n",evt.Market.String())
	l.Printf("\tOutcome: %v\n",evt.Outcome.String())
	l.Printf("\tBalance: %v\n",evt.Balance.String())
	l.Printf("}\n")
}
func (evt *CancelZeroXOrder) Dump(l *log.Logger) {
	l.Printf("CancelZeroXOrder {\n")
	l.Printf("\tUniverse: %v\n",evt.Universe.String())
	l.Printf("\tMarket: %v\n",evt.Market.String())
	l.Printf("\tAccount: %v\n",evt.Account.String())
	l.Printf("\tOutcome: %v\n",evt.Outcome.String())
	l.Printf("\tPrice: %v\n",evt.Price.String())
	l.Printf("\tAmount: %v\n",evt.Amount.String())
	l.Printf("\tOrderType: %v\n",evt.OrderType)
	l.Printf("\tOrderHash: %v\n",hex.EncodeToString(evt.OrderHash[:]))
	l.Printf("\t\n")
}
func (evt *TransferBatch) Dump(zc *ZeroX,l *log.Logger) {
	l.Printf("TransferBatch {\n")
	l.Printf("\tOperator: %v\n",evt.Operator.String())
	l.Printf("\tFrom: %v\n",evt.From.String())
	l.Printf("\tTo: %v\n",evt.To.String())
	ids := Bigint_ptr_slice_to_str(&evt.Ids,",")
	l.Printf("\tIds: %v\n",ids)
	l.Printf("\tDecoded token IDs:\n")
	var copts = new(bind.CallOpts)
	copts.Pending = true
	for i:=0 ; i<len(evt.Ids); i++ {
		if false {
			l.Printf("\t\tcan't decode token info hex string: \n")
		} else {
			tok_info,err := zc.UnpackTokenId(copts,evt.Ids[i])
			if err == nil {
				l.Printf("\t\tMarket: %v\n",tok_info.Market.String())
				l.Printf("\t\tPrice: %v\n",tok_info.Price)
				l.Printf("\t\tOutcome: %v\n",tok_info.Outcome)
				l.Printf("\t\tType: %v\n",tok_info.Type)
			} else {
				l.Printf("\t\ttoken decode error: %v\n",err)
			}
		}
	}
	values := Bigint_ptr_slice_to_str(&evt.Values,",")
	l.Printf("\tValues: %v\n",values)
	l.Printf("}\n")
}
func (evt *TransferSingle) Dump(l *log.Logger) {
	l.Printf("TransferSingle {\n")
	l.Printf("\tOperator: %v\n",evt.Operator.String())
	l.Printf("\tFrom: %v\n",evt.From.String())
	l.Printf("\tTo: %v\n",evt.To.String())
	l.Printf("\tId: %v\n",evt.Id.String())
	l.Printf("\tValue: %v\n",evt.Value.String())
	l.Printf("}\n")
}
func (evt *ProfitLossChanged) Dump(l *log.Logger) {
	l.Printf("ProfitLossChanged {\n")
	l.Printf("\tUniverse: %v\n",evt.Universe.String())
	l.Printf("\tMarket: %v\n",evt.Market.String())
	l.Printf("\tAccount: %v\n",evt.Account.String())
	l.Printf("\tOutcome: %v\n",evt.Outcome.String())
	l.Printf("\tNetPosition: %v\n",evt.NetPosition.String())
	l.Printf("\tAvgPrice: %v\n",evt.AvgPrice.String())
	l.Printf("\tRealizedProfit: %v\n",evt.RealizedProfit.String())
	l.Printf("\tFrozenFunds: %v\n",evt.FrozenFunds.String())
	l.Printf("\tRealizedCost: %v\n",evt.RealizedCost.String())
	l.Printf("\tTimestamp: %v\n",evt.Timestamp.String())
	l.Printf("}\n")
}
func (evt *Transfer) Dump(l *log.Logger) {
	l.Printf("Transfer {\n")
	l.Printf("\tFrom: %v\n",evt.From.String())
	l.Printf("\tTo: %v\n",evt.To.String())
	l.Printf("\tValue: %v\n",evt.Value.String())
	l.Printf("}\n")
}
func (evt *FillEvt) Dump(l *log.Logger) {
	l.Printf("FillEvt {\n")
	l.Printf("\tMakerAddress: %v\n",evt.MakerAddress.String())
	l.Printf("\tFeeRecipientAddress: %v\n",evt.FeeRecipientAddress)
	l.Printf("\tMarketAssetData: %v\n",hex.EncodeToString(evt.MakerAssetData[:]))
	l.Printf("\tTakerAssetData: %v\n",hex.EncodeToString(evt.TakerAssetData[:]))
	l.Printf("\tMakerFeeAssetData: %v\n",hex.EncodeToString(evt.MakerFeeAssetData[:]))
	l.Printf("\tTakerFeeAssetData: %v\n",hex.EncodeToString(evt.TakerFeeAssetData[:]))
	l.Printf("\tOrderHash: %v\n",hex.EncodeToString(evt.OrderHash[:]))
	l.Printf("\tTakerAddress: %v\n",evt.TakerAddress.String())
	l.Printf("\tSenderAddress: %v\n",evt.SenderAddress.String())
	l.Printf("\tMakerAssetFilledAmount: %v\n",evt.MakerAssetFilledAmount.String())
	l.Printf("\tTakerAssetFilledAmount: %v\n",evt.TakerAssetFilledAmount.String())
	l.Printf("\tMakerFeePaid: %v\n",evt.MakerFeePaid.String())
	l.Printf("\tTakerFeePaid: %v\n",evt.TakerFeePaid.String())
	l.Printf("\tProtocolFeePaid: %v\n",evt.ProtocolFeePaid.String())
	l.Printf("}\n")
}
func (evt *OwnershipTransferred) Dump(l *log.Logger) {
	l.Printf("OwnershipTransferred {\n")
	l.Printf("\tPreviousOwner: %v\n",evt.PreviousOwner.String())
	l.Printf("\tNewOwner: %v\n",evt.NewOwner.String())
	l.Printf("}\n")
}
func (evt *TradingProceedsClaimed) Dump(l *log.Logger) {
	l.Printf("TradingProceedsClaimed {\n")
	l.Printf("\tUniverse: %v\n",evt.Universe.String())
	l.Printf("\tSender: %v\n",evt.Sender.String())
	l.Printf("\tMarket: %v\n",evt.Market.String())
	l.Printf("\tOutcome: %v\n",evt.Outcome.String())
	l.Printf("\tNumShares: %v\n",evt.NumShares.String())
	l.Printf("\tNumPayoutTokens: %v\n",evt.NumPayoutTokens.String())
	l.Printf("\tFees: %v\n",evt.Fees.String())
	l.Printf("\tTimestamp: %v\n",evt.Timestamp.String())
	l.Printf("}\n")
}
func (evt *ApprovalForAll) Dump(l *log.Logger) {
	l.Printf("ApprovalForAll {\n")
	l.Printf("\tOwner: %v\n",evt.Owner.String())
	l.Printf("\tOperator: %v\n",evt.Operator.String())
	l.Printf("\tApproved: %v\n",evt.Approved)
	l.Printf("}\n")
}
func (evt* Approval) Dump(l *log.Logger) {
	l.Printf("Approval {\n")
	l.Printf("\tOwner: %v\n",evt.Owner.String())
	l.Printf("\tSpender: %v\n",evt.Spender.String())
	l.Printf("\tValue: %v\n",evt.Value.String())
	l.Printf("}\n")
}
func (evt *ExecuteTransactionStatus) Dump(l *log.Logger) {	// dumps struct to stdout for debugging

	l.Printf("ExecuteTransactionStatus {\n")
	l.Printf("\tSuccess: %v\n",evt.Success)
	l.Printf("\tFundingSuccess: %v\n",evt.FundingSuccess)
	l.Printf("}\n")
}
func (obj *GasSpent) Dump(l *log.Logger) {
	l.Printf("GasSpent {\n")
	l.Printf(
		"\tTrading Gas: %v\tTrading Tx cost: %v\tTrading Txs: %v\n",
		obj.Trading,obj.EthTrading,obj.Num_trading,
	)
	l.Printf(
		"\tReporting Gas: %v\tReporting Tx cost: %v\tReporting Txs: %v\n",
		obj.Reporting,obj.EthReporting,obj.Num_reporting,
	)
	l.Printf(
		"\tMarkets created Gas: %v\tMarkets created TX cost: %v\nMarkets Txs: %v\n",
		obj.Markets,obj.EthMarkets,obj.Num_markets,
	)
	l.Printf(
		"\tTotal Gas spent: %v\tNum Txs: %v\n",
		obj.Total,obj.EthTotal,obj.Num_total,
	)
	l.Printf("}")
}
