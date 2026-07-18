// Event-row deletes for the CosmicGame ETL, one method per event table.
//
// The ETL processes events idempotently: every handler deletes its event's
// rows by evtlog_id before inserting, so re-processing (reorg recovery,
// batch retry) is state-neutral. Deletes are per evt_log row; the plpgsql
// delete triggers of migrations 00002/00003 reverse the aggregate side
// effects.

package cosmicgame

import (
	"context"

	"github.com/PredictionExplorer/augur-explorer/internal/store"
)

// deleteByEvtlogID removes every row of table carrying evtlogID — the
// delete-before-insert half of the ETL's re-processing contract. table is
// always a compile-time literal supplied by the named methods below, never
// request or chain input.
func (r *Repo) deleteByEvtlogID(ctx context.Context, table string, evtlogID int64) error {
	_, err := r.q(ctx).Exec(ctx, "DELETE FROM "+table+" WHERE evtlog_id=$1", evtlogID)
	return store.WrapError("delete from "+table, err)
}

// Game events.

// DeletePrizeClaim removes a MainPrizeClaimed row.
func (r *Repo) DeletePrizeClaim(ctx context.Context, evtlogID int64) error {
	return r.deleteByEvtlogID(ctx, "cg_prize_claim", evtlogID)
}

// DeleteBid removes a BidPlaced row.
func (r *Repo) DeleteBid(ctx context.Context, evtlogID int64) error {
	return r.deleteByEvtlogID(ctx, "cg_bid", evtlogID)
}

// DeleteRoundStarted removes a FirstBidPlacedInRound row.
func (r *Repo) DeleteRoundStarted(ctx context.Context, evtlogID int64) error {
	return r.deleteByEvtlogID(ctx, "cg_first_bid", evtlogID)
}

// Donations.

// DeleteEthDonation removes an EthDonated row.
func (r *Repo) DeleteEthDonation(ctx context.Context, evtlogID int64) error {
	return r.deleteByEvtlogID(ctx, "cg_eth_donated", evtlogID)
}

// DeleteEthDonationWithInfo removes an EthDonatedWithInfo row (the
// cg_donation_json companion row follows via ON DELETE CASCADE).
func (r *Repo) DeleteEthDonationWithInfo(ctx context.Context, evtlogID int64) error {
	return r.deleteByEvtlogID(ctx, "cg_eth_donated_wi", evtlogID)
}

// DeleteDonationReceived removes a CharityWallet DonationReceived row.
func (r *Repo) DeleteDonationReceived(ctx context.Context, evtlogID int64) error {
	return r.deleteByEvtlogID(ctx, "cg_donation_received", evtlogID)
}

// DeleteDonationSent removes a CharityWallet FundsTransferredToCharity row.
func (r *Repo) DeleteDonationSent(ctx context.Context, evtlogID int64) error {
	return r.deleteByEvtlogID(ctx, "cg_donation_sent", evtlogID)
}

// DeleteERC20Donation removes a PrizesWallet TokenDonated row.
func (r *Repo) DeleteERC20Donation(ctx context.Context, evtlogID int64) error {
	return r.deleteByEvtlogID(ctx, "cg_erc20_donation", evtlogID)
}

// DeleteNFTDonation removes a PrizesWallet NftDonated row.
func (r *Repo) DeleteNFTDonation(ctx context.Context, evtlogID int64) error {
	return r.deleteByEvtlogID(ctx, "cg_nft_donation", evtlogID)
}

// DeleteDonatedTokenClaim removes a DonatedTokenClaimed row.
func (r *Repo) DeleteDonatedTokenClaim(ctx context.Context, evtlogID int64) error {
	return r.deleteByEvtlogID(ctx, "cg_donated_tok_claimed", evtlogID)
}

// DeleteDonatedNFTClaim removes a DonatedNftClaimed row.
func (r *Repo) DeleteDonatedNFTClaim(ctx context.Context, evtlogID int64) error {
	return r.deleteByEvtlogID(ctx, "cg_donated_nft_claimed", evtlogID)
}

// DeleteFundsToCharity removes a FundsTransferredToCharity (game) row.
func (r *Repo) DeleteFundsToCharity(ctx context.Context, evtlogID int64) error {
	return r.deleteByEvtlogID(ctx, "cg_funds_to_charity", evtlogID)
}

// Tokens.

// DeleteTokenName removes an NftNameChanged row.
func (r *Repo) DeleteTokenName(ctx context.Context, evtlogID int64) error {
	return r.deleteByEvtlogID(ctx, "cg_token_name", evtlogID)
}

// DeleteMint removes an NftMinted row.
func (r *Repo) DeleteMint(ctx context.Context, evtlogID int64) error {
	return r.deleteByEvtlogID(ctx, "cg_mint_event", evtlogID)
}

// DeleteCosmicSignatureTransfer removes an ERC721 Transfer row.
func (r *Repo) DeleteCosmicSignatureTransfer(ctx context.Context, evtlogID int64) error {
	return r.deleteByEvtlogID(ctx, "cg_erc721_transfer", evtlogID)
}

// DeleteCosmicTokenTransfer removes an ERC20 Transfer row.
func (r *Repo) DeleteCosmicTokenTransfer(ctx context.Context, evtlogID int64) error {
	return r.deleteByEvtlogID(ctx, "cg_erc20_transfer", evtlogID)
}

// Prizes.

// DeletePrizeDeposit removes a PrizesWallet EthReceived row.
func (r *Repo) DeletePrizeDeposit(ctx context.Context, evtlogID int64) error {
	return r.deleteByEvtlogID(ctx, "cg_prize_deposit", evtlogID)
}

// DeletePrizeWithdrawal removes a PrizesWallet EthWithdrawn row.
func (r *Repo) DeletePrizeWithdrawal(ctx context.Context, evtlogID int64) error {
	return r.deleteByEvtlogID(ctx, "cg_prize_withdrawal", evtlogID)
}

// DeleteRaffleNFTWinner removes a RaffleWinnerPrizePaid row.
func (r *Repo) DeleteRaffleNFTWinner(ctx context.Context, evtlogID int64) error {
	return r.deleteByEvtlogID(ctx, "cg_raffle_nft_prize", evtlogID)
}

// DeleteRaffleETHWinner removes a RaffleWinnerBidderEthPrizeAllocated row.
func (r *Repo) DeleteRaffleETHWinner(ctx context.Context, evtlogID int64) error {
	return r.deleteByEvtlogID(ctx, "cg_raffle_eth_prize", evtlogID)
}

// DeleteEnduranceWinner removes an EnduranceChampionPrizePaid row.
func (r *Repo) DeleteEnduranceWinner(ctx context.Context, evtlogID int64) error {
	return r.deleteByEvtlogID(ctx, "cg_endurance_prize", evtlogID)
}

// DeleteLastCstBidderWinner removes a LastCstBidderPrizePaid row.
func (r *Repo) DeleteLastCstBidderWinner(ctx context.Context, evtlogID int64) error {
	return r.deleteByEvtlogID(ctx, "cg_lastcst_prize", evtlogID)
}

// DeleteChronoWarrior removes a ChronoWarriorPrizePaid row.
func (r *Repo) DeleteChronoWarrior(ctx context.Context, evtlogID int64) error {
	return r.deleteByEvtlogID(ctx, "cg_chrono_warrior_prize", evtlogID)
}

// DeleteFundTransferFailed removes a FundTransferFailed row.
func (r *Repo) DeleteFundTransferFailed(ctx context.Context, evtlogID int64) error {
	return r.deleteByEvtlogID(ctx, "cg_fund_transf_err", evtlogID)
}

// DeleteERC20TransferFailed removes an ERC20TransferFailed row.
func (r *Repo) DeleteERC20TransferFailed(ctx context.Context, evtlogID int64) error {
	return r.deleteByEvtlogID(ctx, "cg_erc20_transf_err", evtlogID)
}

// Staking.

// DeleteStakingEthDeposit removes an EthDepositReceived row.
func (r *Repo) DeleteStakingEthDeposit(ctx context.Context, evtlogID int64) error {
	return r.deleteByEvtlogID(ctx, "cg_staking_eth_deposit", evtlogID)
}

// DeleteNftStakedCST removes a CST-wallet NftStaked row.
func (r *Repo) DeleteNftStakedCST(ctx context.Context, evtlogID int64) error {
	return r.deleteByEvtlogID(ctx, "cg_nft_staked_cst", evtlogID)
}

// DeleteNftStakedRWalk removes a RandomWalk-wallet NftStaked row.
func (r *Repo) DeleteNftStakedRWalk(ctx context.Context, evtlogID int64) error {
	return r.deleteByEvtlogID(ctx, "cg_nft_staked_rwalk", evtlogID)
}

// DeleteNftUnstakedCST removes a CST-wallet NftUnstaked row.
func (r *Repo) DeleteNftUnstakedCST(ctx context.Context, evtlogID int64) error {
	return r.deleteByEvtlogID(ctx, "cg_nft_unstaked_cst", evtlogID)
}

// DeleteNftUnstakedRWalk removes a RandomWalk-wallet NftUnstaked row.
func (r *Repo) DeleteNftUnstakedRWalk(ctx context.Context, evtlogID int64) error {
	return r.deleteByEvtlogID(ctx, "cg_nft_unstaked_rwalk", evtlogID)
}

// Marketing.

// DeleteMarketingRewardPaid removes a MarketingWallet RewardPaid row.
func (r *Repo) DeleteMarketingRewardPaid(ctx context.Context, evtlogID int64) error {
	return r.deleteByEvtlogID(ctx, "cg_mkt_reward", evtlogID)
}

// Admin: percentages and counts.

// DeleteCharityPercentageChange removes a CharityEthDonationAmountPercentageChanged row.
func (r *Repo) DeleteCharityPercentageChange(ctx context.Context, evtlogID int64) error {
	return r.deleteByEvtlogID(ctx, "cg_adm_charity_pcent", evtlogID)
}

// DeletePrizePercentageChange removes a MainEthPrizeAmountPercentageChanged row.
func (r *Repo) DeletePrizePercentageChange(ctx context.Context, evtlogID int64) error {
	return r.deleteByEvtlogID(ctx, "cg_adm_main_prize_pcent", evtlogID)
}

// DeleteRafflePercentageChange removes a RaffleTotalEthPrizeAmountForBiddersPercentageChanged row.
func (r *Repo) DeleteRafflePercentageChange(ctx context.Context, evtlogID int64) error {
	return r.deleteByEvtlogID(ctx, "cg_adm_raffle_pcent", evtlogID)
}

// DeleteStakingPercentageChange removes a CosmicSignatureNftStakingTotalEthRewardAmountPercentageChanged row.
func (r *Repo) DeleteStakingPercentageChange(ctx context.Context, evtlogID int64) error {
	return r.deleteByEvtlogID(ctx, "cg_adm_stake_pcent", evtlogID)
}

// DeleteChronoPercentageChange removes a ChronoWarriorEthPrizeAmountPercentageChanged row.
func (r *Repo) DeleteChronoPercentageChange(ctx context.Context, evtlogID int64) error {
	return r.deleteByEvtlogID(ctx, "cg_adm_chrono_pcent", evtlogID)
}

// DeleteNumRaffleETHWinnersBiddingChange removes a NumRaffleEthPrizesForBiddersChanged row.
func (r *Repo) DeleteNumRaffleETHWinnersBiddingChange(ctx context.Context, evtlogID int64) error {
	return r.deleteByEvtlogID(ctx, "cg_adm_raf_eth_bidding", evtlogID)
}

// DeleteNumRaffleNFTWinnersBiddingChange removes a NumRaffleCosmicSignatureNftsForBiddersChanged row.
func (r *Repo) DeleteNumRaffleNFTWinnersBiddingChange(ctx context.Context, evtlogID int64) error {
	return r.deleteByEvtlogID(ctx, "cg_adm_raf_nft_bidding", evtlogID)
}

// DeleteNumRaffleNFTWinnersStakingRWalkChange removes a
// NumRaffleCosmicSignatureNftsForRandomWalkNftStakersChanged row.
func (r *Repo) DeleteNumRaffleNFTWinnersStakingRWalkChange(ctx context.Context, evtlogID int64) error {
	return r.deleteByEvtlogID(ctx, "cg_adm_raf_nft_staking_rwalk", evtlogID)
}

// Admin: address changes.

// DeleteCharityReceiverChange removes a CharityWallet CharityAddressChanged row.
func (r *Repo) DeleteCharityReceiverChange(ctx context.Context, evtlogID int64) error {
	return r.deleteByEvtlogID(ctx, "cg_charity_receiver_changed", evtlogID)
}

// DeleteCharityWalletAddressChange removes a game CharityAddressChanged row.
func (r *Repo) DeleteCharityWalletAddressChange(ctx context.Context, evtlogID int64) error {
	return r.deleteByEvtlogID(ctx, "cg_adm_charity_wallet", evtlogID)
}

// DeleteRandomWalkAddressChange removes a RandomWalkNftAddressChanged row.
func (r *Repo) DeleteRandomWalkAddressChange(ctx context.Context, evtlogID int64) error {
	return r.deleteByEvtlogID(ctx, "cg_adm_rwalk_addr", evtlogID)
}

// DeletePrizesWalletAddressChange removes a PrizesWalletAddressChanged row.
func (r *Repo) DeletePrizesWalletAddressChange(ctx context.Context, evtlogID int64) error {
	return r.deleteByEvtlogID(ctx, "cg_adm_prizes_wallet_addr", evtlogID)
}

// DeleteStakingWalletCSTAddressChange removes a StakingWalletCosmicSignatureNftAddressChanged row.
func (r *Repo) DeleteStakingWalletCSTAddressChange(ctx context.Context, evtlogID int64) error {
	return r.deleteByEvtlogID(ctx, "cg_adm_staking_cst_addr", evtlogID)
}

// DeleteStakingWalletRWalkAddressChange removes a StakingWalletRandomWalkNftAddressChanged row.
func (r *Repo) DeleteStakingWalletRWalkAddressChange(ctx context.Context, evtlogID int64) error {
	return r.deleteByEvtlogID(ctx, "cg_adm_staking_rwalk_addr", evtlogID)
}

// DeleteMarketingWalletAddressChange removes a MarketingWalletAddressChanged row.
func (r *Repo) DeleteMarketingWalletAddressChange(ctx context.Context, evtlogID int64) error {
	return r.deleteByEvtlogID(ctx, "cg_adm_marketing_addr", evtlogID)
}

// DeleteTreasurerAddressChange removes a TreasurerAddressChanged row.
func (r *Repo) DeleteTreasurerAddressChange(ctx context.Context, evtlogID int64) error {
	return r.deleteByEvtlogID(ctx, "cg_adm_treasurer_addr", evtlogID)
}

// DeleteCosmicTokenAddressChange removes a CosmicSignatureTokenAddressChanged row.
func (r *Repo) DeleteCosmicTokenAddressChange(ctx context.Context, evtlogID int64) error {
	return r.deleteByEvtlogID(ctx, "cg_adm_costok_addr", evtlogID)
}

// DeleteCosmicSignatureAddressChange removes a CosmicSignatureNftAddressChanged row.
func (r *Repo) DeleteCosmicSignatureAddressChange(ctx context.Context, evtlogID int64) error {
	return r.deleteByEvtlogID(ctx, "cg_adm_cossig_addr", evtlogID)
}

// Admin: proxy and lifecycle.

// DeleteUpgraded removes an ERC-1967 Upgraded row.
func (r *Repo) DeleteUpgraded(ctx context.Context, evtlogID int64) error {
	return r.deleteByEvtlogID(ctx, "cg_adm_upgraded", evtlogID)
}

// DeleteAdminChanged removes an ERC-1967 AdminChanged row.
func (r *Repo) DeleteAdminChanged(ctx context.Context, evtlogID int64) error {
	return r.deleteByEvtlogID(ctx, "cg_adm_admin_changed", evtlogID)
}

// DeleteOwnershipTransfer removes an OwnershipTransferred row.
func (r *Repo) DeleteOwnershipTransfer(ctx context.Context, evtlogID int64) error {
	return r.deleteByEvtlogID(ctx, "cg_adm_ownership", evtlogID)
}

// DeleteInitialized removes an Initialized row.
func (r *Repo) DeleteInitialized(ctx context.Context, evtlogID int64) error {
	return r.deleteByEvtlogID(ctx, "cg_adm_initialized", evtlogID)
}

// Admin: timing and pricing parameters.

// DeleteTimeIncreaseChange removes a legacy TimeIncreaseChanged row.
func (r *Repo) DeleteTimeIncreaseChange(ctx context.Context, evtlogID int64) error {
	return r.deleteByEvtlogID(ctx, "cg_adm_time_inc", evtlogID)
}

// DeleteTimeoutClaimPrizeChange removes a TimeoutDurationToClaimMainPrizeChanged row.
func (r *Repo) DeleteTimeoutClaimPrizeChange(ctx context.Context, evtlogID int64) error {
	return r.deleteByEvtlogID(ctx, "cg_adm_timeout_claimprize", evtlogID)
}

// DeleteTimeoutToWithdrawPrizesChange removes a TimeoutDurationToWithdrawPrizesChanged row.
func (r *Repo) DeleteTimeoutToWithdrawPrizesChange(ctx context.Context, evtlogID int64) error {
	return r.deleteByEvtlogID(ctx, "cg_adm_timeout_withdraw", evtlogID)
}

// DeletePriceIncreaseChange removes an EthBidPriceIncreaseDivisorChanged row.
func (r *Repo) DeletePriceIncreaseChange(ctx context.Context, evtlogID int64) error {
	return r.deleteByEvtlogID(ctx, "cg_adm_price_inc", evtlogID)
}

// DeleteMainPrizeMicrosecondsChange removes a MainPrizeTimeIncrementInMicroSecondsChanged row.
func (r *Repo) DeleteMainPrizeMicrosecondsChange(ctx context.Context, evtlogID int64) error {
	return r.deleteByEvtlogID(ctx, "cg_adm_prize_microsec", evtlogID)
}

// DeleteInitialSecondsUntilPrizeChange removes an InitialDurationUntilMainPrizeDivisorChanged row.
func (r *Repo) DeleteInitialSecondsUntilPrizeChange(ctx context.Context, evtlogID int64) error {
	return r.deleteByEvtlogID(ctx, "cg_adm_inisecprize", evtlogID)
}

// DeleteActivationTimeChange removes a RoundActivationTimeChanged row.
func (r *Repo) DeleteActivationTimeChange(ctx context.Context, evtlogID int64) error {
	return r.deleteByEvtlogID(ctx, "cg_adm_acttime", evtlogID)
}

// DeleteCstAuctionLengthChange removes a CST dutch-auction duration/divisor row.
func (r *Repo) DeleteCstAuctionLengthChange(ctx context.Context, evtlogID int64) error {
	return r.deleteByEvtlogID(ctx, "cg_adm_cst_auclen", evtlogID)
}

// DeleteCstAuctionDurationChangeDivisorChange removes a CstDutchAuctionDurationChangeDivisorChanged row.
func (r *Repo) DeleteCstAuctionDurationChangeDivisorChange(ctx context.Context, evtlogID int64) error {
	return r.deleteByEvtlogID(ctx, "cg_adm_cst_auclen_chg_div", evtlogID)
}

// DeleteEthAuctionDurationDivisorChange removes an EthDutchAuctionDurationDivisorChanged row.
func (r *Repo) DeleteEthAuctionDurationDivisorChange(ctx context.Context, evtlogID int64) error {
	return r.deleteByEvtlogID(ctx, "cg_adm_eth_auclen", evtlogID)
}

// DeleteEthAuctionEndingBidPriceDivisorChange removes an EthDutchAuctionEndingBidPriceDivisorChanged row.
func (r *Repo) DeleteEthAuctionEndingBidPriceDivisorChange(ctx context.Context, evtlogID int64) error {
	return r.deleteByEvtlogID(ctx, "cg_adm_eth_auc_endprice", evtlogID)
}

// Admin: rewards and limits.

// DeleteStaticCstRewardChange removes a CstPrizeAmountChanged row.
func (r *Repo) DeleteStaticCstRewardChange(ctx context.Context, evtlogID int64) error {
	return r.deleteByEvtlogID(ctx, "cg_adm_erc_rwd_mul", evtlogID)
}

// DeleteMarketingRewardChange removes a MarketingWalletCstContributionAmountChanged row.
func (r *Repo) DeleteMarketingRewardChange(ctx context.Context, evtlogID int64) error {
	return r.deleteByEvtlogID(ctx, "cg_adm_mkt_reward", evtlogID)
}

// DeleteCstRewardForBiddingChange removes a CST-bid-reward change row
// (CstRewardAmountForBiddingChanged and its V2 variants share the table).
func (r *Repo) DeleteCstRewardForBiddingChange(ctx context.Context, evtlogID int64) error {
	return r.deleteByEvtlogID(ctx, "cg_adm_erc20_reward", evtlogID)
}

// DeleteMaxMessageLengthChange removes a BidMessageLengthMaxLimitChanged row.
func (r *Repo) DeleteMaxMessageLengthChange(ctx context.Context, evtlogID int64) error {
	return r.deleteByEvtlogID(ctx, "cg_adm_msg_len", evtlogID)
}

// DeleteCstMinLimit removes a CstDutchAuctionBeginningBidPriceMinLimitChanged row.
func (r *Repo) DeleteCstMinLimit(ctx context.Context, evtlogID int64) error {
	return r.deleteByEvtlogID(ctx, "cg_adm_cst_min_limit", evtlogID)
}

// DeleteNextRoundDelayDurationChange removes a DelayDurationBeforeRoundActivationChanged row.
func (r *Repo) DeleteNextRoundDelayDurationChange(ctx context.Context, evtlogID int64) error {
	return r.deleteByEvtlogID(ctx, "cg_delay_duration", evtlogID)
}

// Admin: NFT metadata.

// DeleteTokenGenerationScriptURL removes an NftGenerationScriptUriChanged row.
func (r *Repo) DeleteTokenGenerationScriptURL(ctx context.Context, evtlogID int64) error {
	return r.deleteByEvtlogID(ctx, "cg_adm_script_url", evtlogID)
}

// DeleteBaseURI removes an NftBaseUriChanged row.
func (r *Repo) DeleteBaseURI(ctx context.Context, evtlogID int64) error {
	return r.deleteByEvtlogID(ctx, "cg_adm_base_uri_cs", evtlogID)
}
