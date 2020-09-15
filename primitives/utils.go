package primitives

import (
	"log"
	"encoding/hex"
	"bytes"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/0xProject/0x-mesh/common/types"
)
func (evt *EMarketCreated) Dump(l *log.Logger) {	// dumps struct to stdout for debugging
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
func (evt *EMarketOIChanged) Dump(l *log.Logger) {	// dumps struct to stdout for debugging

	l.Printf("MarketOIChanged {\n")
	l.Printf("\tUniverse: %v\n",evt.Universe.String())
	l.Printf("\tMarket: %v\n",evt.Market.String())
	l.Printf("\tMarket Open Interest: %v\n",evt.MarketOI.String())
	l.Printf("}\n")
}
func (evt *EOrderEvent) Dump(l *log.Logger) { // dumps struct to stdout for debugging

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
func (evt *EMarketFinalized) Dump(l *log.Logger) { // dumps struct to stdout for debugging

	l.Printf("MarketFinalizedEvent {\n")
	l.Printf("\tUniverse: %v\n",evt.Universe.String())
	l.Printf("\tMarket: %v\n",evt.Market.String())
	l.Printf("\tTimestamp: %v\n",evt.Timestamp)
	payouts := Bigint_ptr_slice_to_str(&evt.WinningPayoutNumerators,",")
	l.Printf("\tWinningPayouts: %v\n",payouts)
	l.Printf("}\n")
}
func (evt *EInitialReportSubmitted) Dump(l *log.Logger) {

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
func (evt *EDisputeCrowdsourcerContribution) Dump(l *log.Logger) {

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
func (evt *EMarketVolumeChanged_v1) Dump(l *log.Logger) { // dumps struct to stdout for debugging

	l.Printf("MarketVolumeChanged_v1 {\n")
	l.Printf("\tUniverse: %v\n",evt.Universe.String())
	l.Printf("\tMarket: %v\n",evt.Market.String())
	l.Printf("\tVolume: %v\n",evt.Volume.String())
	outcome_volumes := Bigint_ptr_slice_to_str(&evt.OutcomeVolumes,",")
	l.Printf("\tOutcomeVolumes: %v\n",outcome_volumes)
	l.Printf("\tTimestamp: %v\n",evt.Timestamp)
	l.Printf("}\n")
}
func (evt *EMarketVolumeChanged_v2) Dump(l *log.Logger) { // dumps struct to stdout for debugging

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
func (evt *ETokensTransferred) Dump(l *log.Logger) {

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
func (evt *ETokenBalanceChanged) Dump(l *log.Logger) {

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
func (evt *EShareTokenBalanceChanged) Dump(l *log.Logger) {
	l.Printf("ShareTokensBalanceChanged {\n")
	l.Printf("\tUniverse: %v\n",evt.Universe.String())
	l.Printf("\tAccount: %v\n",evt.Account.String())
	l.Printf("\tMarket: %v\n",evt.Market.String())
	l.Printf("\tOutcome: %v\n",evt.Outcome.String())
	l.Printf("\tBalance: %v\n",evt.Balance.String())
	l.Printf("}\n")
}
func (evt *ECancelZeroXOrder) Dump(l *log.Logger) {
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
func (evt *ETransferBatch) Dump(zc *ZeroX,l *log.Logger) {
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
func (evt *ETransferSingle) Dump(l *log.Logger) {
	l.Printf("TransferSingle {\n")
	l.Printf("\tOperator: %v\n",evt.Operator.String())
	l.Printf("\tFrom: %v\n",evt.From.String())
	l.Printf("\tTo: %v\n",evt.To.String())
	l.Printf("\tId: %v\n",evt.Id.String())
	l.Printf("\tValue: %v\n",evt.Value.String())
	l.Printf("}\n")
}
func (evt *EProfitLossChanged) Dump(l *log.Logger) {
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
func (evt *ETransfer) Dump(l *log.Logger) {
	l.Printf("Transfer {\n")
	l.Printf("\tFrom: %v\n",evt.From.String())
	l.Printf("\tTo: %v\n",evt.To.String())
	l.Printf("\tValue: %v\n",evt.Value.String())
	l.Printf("}\n")
}
func (evt *EFill) Dump(l *log.Logger) {
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
func (evt *EOwnershipTransferred) Dump(l *log.Logger) {
	l.Printf("OwnershipTransferred {\n")
	l.Printf("\tPreviousOwner: %v\n",evt.PreviousOwner.String())
	l.Printf("\tNewOwner: %v\n",evt.NewOwner.String())
	l.Printf("}\n")
}
func (evt *ETradingProceedsClaimed) Dump(l *log.Logger) {
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
func (evt *EApprovalForAll) Dump(l *log.Logger) {
	l.Printf("ApprovalForAll {\n")
	l.Printf("\tOwner: %v\n",evt.Owner.String())
	l.Printf("\tOperator: %v\n",evt.Operator.String())
	l.Printf("\tApproved: %v\n",evt.Approved)
	l.Printf("}\n")
}
func (evt* EApproval) Dump(l *log.Logger) {
	l.Printf("Approval {\n")
	l.Printf("\tOwner: %v\n",evt.Owner.String())
	l.Printf("\tSpender: %v\n",evt.Spender.String())
	l.Printf("\tValue: %v\n",evt.Value.String())
	l.Printf("}\n")
}
func (evt *EExecuteTransactionStatus) Dump(l *log.Logger) {	// dumps struct to stdout for debugging

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
func (obj *ExecuteWalletTx) Dump(l *log.Logger) {
	l.Printf("ExecuteWalletTransaction {\n")
	l.Printf("\tto: %v\n",obj.To)
	l.Printf("\tdata: %v\n",obj.CallData)
	l.Printf("\tinput_sig: %v\n",obj.InputSig)
	l.Printf("\tvalue: %v\n",obj.Value)
	l.Printf("\tpayment: %v\n",obj.Payment)
	l.Printf("\treferralAddress:  %v\n",obj.ReferralAddress)
	l.Printf("\tfingerprint: %v\n",obj.Fingerprint)
	l.Printf("\tdesiredSignerBalance: %v\n",obj.DesiredSignerBalance)
	l.Printf("\tmaxExchangeRateInDai: %v\n",obj.MaxExchangeRateInDAI)
	l.Printf("\trevertOnFaliure: %v\n",obj.RevertOnFailure)
	l.Printf("}\n")
}
func (evt *ERegisterContract) Dump(l *log.Logger) {
	l.Printf("RegisterContract {\n")
	l.Printf("\tContractAddress: %v\n",evt.ContractAddress.String())
	length := bytes.Index(evt.Key[:],[]byte{0})
	l.Printf("\tKey: %v\n",string(evt.Key[:length]))
	l.Printf("}\n")
}
func (caddrs *ContractAddresses) Dump(l *log.Logger) {
	l.Printf("ContractAddresses {\n")
	l.Printf("\tChainID: %v\n",caddrs.ChainId)
	l.Printf("\tAugur: %v\n",caddrs.Augur.String())
	l.Printf("\tAugurTrading: %v\n",caddrs.AugurTrading.String())
	l.Printf("\tProfitLoss: %v\n",caddrs.PL.String())
	l.Printf("\tCash: %v\n",caddrs.Dai.String())
	l.Printf("\tZeroXTrade: %v\n",caddrs.ZeroxTrade.String())
	l.Printf("\tExchange: %v\n",caddrs.ZeroxXchg.String())
	l.Printf("\tREPv2: %v\n",caddrs.Reputation.String())
	l.Printf("\tAugurWalletRegistry: %v\n",caddrs.WalletReg.String())
	l.Printf("\tAugurWalletRegistryV2: %v\n",caddrs.WalletReg2.String())
	l.Printf("\tFillOrder: %v\n",caddrs.FillOrder.String())
	l.Printf("\tEthExchange: %v\n",caddrs.EthXchg.String())
	l.Printf("\tShareToken: %v\n",caddrs.ShareToken.String())
	l.Printf("\tGenesisUniverse: %v\n",caddrs.GenesisUniverse.String())
	l.Printf("\tCreateOrder: %v\n",caddrs.CreateOrder.String())
	l.Printf("\tLegacyReputationToken: %v\n",caddrs.LegacyReputationToken.String())
	l.Printf("\tBuyParticipationTokens: %v\n",caddrs.BuyParticipationTokens.String())
	l.Printf("\tRedeemStake: %v\n",caddrs.RedeemStake.String())
	l.Printf("\tWarpSync: %v\n",caddrs.WarpSync.String())
	l.Printf("\tHotLoading:%v\n",caddrs.HotLoading.String())
	l.Printf("\tAffiliates: %v\n",caddrs.Affiliates.String())
	l.Printf("\tAffiliateValidator: %v\n",caddrs.AffiliateValidator.String())
	l.Printf("\tTime: %v\n",caddrs.Time.String())
	l.Printf("\tCancelOrder: %v\n",caddrs.CancelOrder.String())
	l.Printf("\tOrders: %v\n",caddrs.Orders.String())
	l.Printf("\tSimulateTrade: %v\n",caddrs.SimulateTrade.String())
	l.Printf("\tTrade: %v\n",caddrs.Trade.String())
	l.Printf("\tOICash: %v\n",caddrs.OICash.String())
	l.Printf("\tUniswapV2Factory: %v\n",caddrs.UniswapV2Factory.String())
	l.Printf("\tUniswapV2Router02: %v\n",caddrs.UniswapV2Router02.String())
	l.Printf("\tAuditFunds: %v\n",caddrs.AuditFunds.String())
	l.Printf("\tWETH9: %v\n",caddrs.WETH9.String())
	l.Printf("\tUSDC: %v\n",caddrs.USDC.String())
	l.Printf("\tUSDT: %v\n",caddrs.USDT.String())
	l.Printf("\tRelayHubV2: %v\n",caddrs.RelayHubV2.String())
	l.Printf("\tAccountLoader: %v\n",caddrs.AccountLoader.String())
	l.Printf("}\n")
}
func (evt *EUniverseCreated) Dump(l *log.Logger) {
	l.Printf("UniverseCreated {\n")
	l.Printf("\tParentUniverse: %v\n",evt.ParentUniverse.String())
	l.Printf("\tChildUniverse: %v\n",evt.ChildUniverse.String())
	l.Printf("\tPayoutNumerators: %v\n",Bigint_ptr_slice_to_str(&evt.PayoutNumerators,","))
	l.Printf("\tCreationTimestamp: %v\n",evt.CreationTimestamp.Int64())
	l.Printf("}\n")
}
func Dump_0x_mesh_order(l *log.Logger,o *types.OrderInfo) {
	l.Printf("0x Mesh Order {\n")
	l.Printf("\tOrderHash: %v\n",o.OrderHash.String())
	l.Printf("\tSignature: %v\n",hex.EncodeToString(o.SignedOrder.Signature))
	l.Printf("\tFillableTakerAssetAmount: %v\n",o.FillableTakerAssetAmount.String())
	l.Printf("\tChainId: %v\n",o.SignedOrder.ChainID.String())
	l.Printf("\tExchangeAddress: %v\n",o.SignedOrder.ExchangeAddress.String())
	l.Printf("\tMakerAddress: %v\n",o.SignedOrder.MakerAddress.String())
	l.Printf("\tMakerAssetData: %v\n",hex.EncodeToString(o.SignedOrder.MakerAssetData))
	l.Printf("\tMakerFeeAssetData: %v\n",hex.EncodeToString(o.SignedOrder.MakerFeeAssetData))
	l.Printf("\tMakerAssetAmount: %v\n",o.SignedOrder.MakerAssetAmount.String())
	l.Printf("\tMakerFee: %v\n",o.SignedOrder.MakerFee.String())
	l.Printf("\tTakerAddress: %v\n",o.SignedOrder.TakerAddress.String())
	l.Printf("\tTakerAssetData: %v\n",hex.EncodeToString(o.SignedOrder.TakerAssetData))
	l.Printf("\tTakerFeeAssetData: %v\n",hex.EncodeToString(o.SignedOrder.TakerFeeAssetData))
	l.Printf("\tTakerAssetAmount: %v\n",o.SignedOrder.TakerAssetAmount.String())
	l.Printf("\tTakerFee: %v\n",o.SignedOrder.TakerFee.String())
	l.Printf("\tSenderAddress: %v\n",o.SignedOrder.SenderAddress.String())
	l.Printf("\tFeeRecipientAddress: %v\n",o.SignedOrder.FeeRecipientAddress.String())
	l.Printf("\tExpirationTimeSeconds: %v\n",o.SignedOrder.ExpirationTimeSeconds.String())
	l.Printf("\tSalt: %v\n",o.SignedOrder.Salt.String())
	l.Printf("}\n")
}
