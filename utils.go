package main

import (
//	"runtime"
	"fmt"
//	"os"
//	"io"
	"encoding/hex"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
)
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
	fmt.Printf("\tUniverse: %v\n",evt.Universe.String())
	fmt.Printf("\tReporter: %v\n",evt.Reporter.String())
	fmt.Printf("\tMarket: %v\n",evt.Market.String())
	fmt.Printf("\tInitialReporter: %v\n",evt.InitialReporter.String())
	fmt.Printf("\tAmountStaked: %v\n",evt.AmountStaked)
	fmt.Printf("\tIsDesignatedReporter: %v\n",evt.IsDesignatedReporter)
	payout_numerators := bigint_ptr_slice_to_str(&evt.PayoutNumerators,",")
	fmt.Printf("\tPayoutNumerators: %v\n",payout_numerators)
	fmt.Printf("\tDescription: %v\n",evt.Description)
	fmt.Printf("\tNextWindowStartTime: %v\n",evt.NextWindowStartTime)
	fmt.Printf("\tNextWindowEndTime: %v\n",evt.NextWindowEndTime)
	fmt.Printf("\tTimestamp: %v\n",evt.Timestamp)
	fmt.Printf("}\n")
}
func (evt *DisputeCrowdsourcerContributionEvt) Dump() {

	fmt.Printf("DisputeCrowdsourcerContribution {\n")
	fmt.Printf("\tUniverse: %v\n",evt.Universe.String())
	fmt.Printf("\tReporter: %v\n",evt.Reporter.String())
	fmt.Printf("\tMarket: %v\n",evt.Market.String())
	fmt.Printf("\tDisputedCrowdsourcer: %v\n",evt.DisputeCrowdsourcer.String())
	fmt.Printf("\tAmountStaked: %v\n",evt.AmountStaked)
	fmt.Printf("\tDescription: %v\n",evt.Description)
	payout_numerators := bigint_ptr_slice_to_str(&evt.PayoutNumerators,",")
	fmt.Printf("\tPayoutNumerators: %v\n",payout_numerators)
	fmt.Printf("\tCurrentStake: %v\n",evt.CurrentStake)
	fmt.Printf("\tStakeRemaining: %v\n",evt.StakeRemaining)
	fmt.Printf("\tDisputeRound %v\n",evt.DisputeRound)
	fmt.Printf("\tTimestamp: %v\n",evt.Timestamp)
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
	fmt.Printf("\tDecoded token IDs:\n")
	var copts = new(bind.CallOpts)
	copts.Pending = true
	for i:=0 ; i<len(evt.Ids); i++ {
		//tok_data,err := hex.DecodeString(evt.Ids[i])
		if false {
			fmt.Printf("\t\tcan't decode token info hex string: \n")
		} else {
			tok_info,err := ctrct_zerox.UnpackTokenId(copts,evt.Ids[i])
			//tok_info,err := ctrct_zerox.DecodeAssetData(copts,tok_data)
			if err == nil {
				fmt.Printf("\t\tMarket: %v\n",tok_info.Market.String())
				fmt.Printf("\t\tPrice: %v\n",tok_info.Price)
				fmt.Printf("\t\tOutcome: %v\n",tok_info.Outcome)
				fmt.Printf("\t\tType: %v\n",tok_info.Type)
			} else {
				fmt.Printf("\t\ttoken decode error: %v\n",err)
			}
		}
	}
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
	fmt.Printf("}\n")
}
func (evt *FillEvt) Dump() {
	fmt.Printf("FillEvt {\n")
	fmt.Printf("\tMakerAddress: %v\n",evt.MakerAddress.String())
	fmt.Printf("\tFeeRecipientAddress: %v\n",evt.FeeRecipientAddress)
	fmt.Printf("\tMarketAssetData: %v\n",hex.EncodeToString(evt.MakerAssetData[:]))
	fmt.Printf("\tTakerAssetData: %v\n",hex.EncodeToString(evt.TakerAssetData[:]))
	fmt.Printf("\tMakerFeeAssetData: %v\n",hex.EncodeToString(evt.MakerFeeAssetData[:]))
	fmt.Printf("\tTakerFeeAssetData: %v\n",hex.EncodeToString(evt.TakerFeeAssetData[:]))
	fmt.Printf("\tOrderHash: %v\n",hex.EncodeToString(evt.OrderHash[:]))
	fmt.Printf("\tTakerAddress: %v\n",evt.TakerAddress.String())
	fmt.Printf("\tSenderAddress: %v\n",evt.SenderAddress.String())
	fmt.Printf("\tMakerAssetFilledAmount: %v\n",evt.MakerAssetFilledAmount.String())
	fmt.Printf("\tTakerAssetFilledAmount: %v\n",evt.TakerAssetFilledAmount.String())
	fmt.Printf("\tMakerFeePaid: %v\n",evt.MakerFeePaid.String())
	fmt.Printf("\tTakerFeePaid: %v\n",evt.TakerFeePaid.String())
	fmt.Printf("\tProtocolFeePaid: %v\n",evt.ProtocolFeePaid.String())
	fmt.Printf("}\n")
}
func (evt *OwnershipTransferred) Dump() {
	fmt.Printf("OwnershipTransferred {\n")
	fmt.Printf("\tPreviousOwner: %v\n",evt.PreviousOwner.String())
	fmt.Printf("\tNewOwner: %v\n",evt.NewOwner.String())
	fmt.Printf("}\n")
}
func (evt *TradingProceedsClaimed) Dump() {
	fmt.Printf("TradingProceedsClaimed {\n")
	fmt.Printf("\tUniverse: %v\n",evt.Universe.String())
	fmt.Printf("\tSender: %v\n",evt.Sender.String())
	fmt.Printf("\tMarket: %v\n",evt.Market.String())
	fmt.Printf("\tOutcome: %v\n",evt.Outcome.String())
	fmt.Printf("\tNumShares: %v\n",evt.NumShares.String())
	fmt.Printf("\tNumPayoutTokens: %v\n",evt.NumPayoutTokens.String())
	fmt.Printf("\tFees: %v\n",evt.Fees.String())
	fmt.Printf("\tTimestamp: %v\n",evt.Timestamp.String())
	fmt.Printf("}\n")
}
func (evt *ApprovalForAll) Dump() {
	fmt.Printf("ApprovalForAll {\n")
	fmt.Printf("\tOwner: %v\n",evt.Owner.String())
	fmt.Printf("\tOperator: %v\n",evt.Operator.String())
	fmt.Printf("\tApproved: %v\n",evt.Approved)
	fmt.Printf("}\n")
}
func (evt* Approval) Dump() {
	fmt.Printf("Approval {\n")
	fmt.Printf("\tOwner: %v\n",evt.Owner.String())
	fmt.Printf("\tSpender: %v\n",evt.Spender.String())
	fmt.Printf("\tValue: %v\n",evt.Value.String())
	fmt.Printf("}\n")
}
