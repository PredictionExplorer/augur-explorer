// Event-row inserts for the CosmicGame ETL, one method per event type.
// Every method resolves its address foreign keys through the Store's
// lookup-or-create cache, runs a single INSERT (the plpgsql triggers of
// migrations 00002/00003 maintain the aggregates) and returns errors instead
// of terminating the process — the polling loop leaves the batch
// unacknowledged so it re-processes after restart.

package cosmicgame

import (
	"context"
	"fmt"

	cgmodel "github.com/PredictionExplorer/augur-explorer/internal/model/cosmicgame"
	"github.com/PredictionExplorer/augur-explorer/internal/store"
)

// insertAdminValue inserts the five columns every admin/parameter event
// shares plus the one event-specific value column. table and column are
// compile-time literals supplied by the callers below.
func (r *Repo) insertAdminValue(ctx context.Context, table, column string, evtID, blockNum, txID, timeStamp, contractAid int64, value any) error {
	query := "INSERT INTO " + table +
		"(evtlog_id,block_num,tx_id,time_stamp,contract_aid," + column +
		") VALUES($1,$2,$3,TO_TIMESTAMP($4),$5,$6)"
	_, err := r.q(ctx).Exec(ctx, query, evtID, blockNum, txID, timeStamp, contractAid, value)
	return store.WrapError("insert into "+table, err)
}

// Game events.

// InsertPrizeClaim records a MainPrizeClaimed event.
func (r *Repo) InsertPrizeClaim(ctx context.Context, evt *cgmodel.CGPrizeClaimEvent) error {
	const op = "insert into cg_prize_claim"
	contractAid, err := r.addrID(ctx, evt.ContractAddr, 0, 0)
	if err != nil {
		return store.WrapError(op, err)
	}
	winnerAid, err := r.addrID(ctx, evt.WinnerAddr, 0, 0)
	if err != nil {
		return store.WrapError(op, err)
	}
	query := "INSERT INTO cg_prize_claim(" +
		"evtlog_id,block_num,time_stamp,tx_id,contract_aid," +
		"round_num,token_id,winner_aid,timeout,amount,cst_amount" +
		") VALUES($1,$2,TO_TIMESTAMP($3),$4,$5,$6,$7,$8,$9,$10,$11)"
	_, err = r.q(ctx).Exec(ctx, query,
		evt.EvtId,
		evt.BlockNum,
		evt.TimeStamp,
		evt.TxId,
		contractAid,
		evt.RoundNum,
		evt.TokenId,
		winnerAid,
		evt.Timeout,
		evt.Amount,
		evt.CstAmount,
	)
	return store.WrapError(op, err)
}

// InsertBid records a BidPlaced event (V1 and V2 share the shape; V2 carries
// real BidCstRewardAmount/CstDutchAuctionDuration values, V1 passes "-1").
// The bid position is derived from the bids already stored for the round;
// the CST bidding reward falls back to cg_glob_stats.cst_reward_for_bidding
// (populated by admin events or the ETL startup chain sync) for V1 events.
func (r *Repo) InsertBid(ctx context.Context, evt *cgmodel.CGBidEvent) error {
	const op = "insert into cg_bid"
	contractAid, err := r.addrID(ctx, evt.ContractAddr, 0, 0)
	if err != nil {
		return store.WrapError(op, err)
	}
	bidderAid, err := r.addrID(ctx, evt.LastBidderAddr, 0, 0)
	if err != nil {
		return store.WrapError(op, err)
	}

	var bidPosition int64
	err = r.q(ctx).QueryRow(ctx,
		"SELECT COALESCE(MAX(bid_position), 0) + 1 FROM cg_bid WHERE round_num = $1",
		evt.RoundNum).Scan(&bidPosition)
	if err != nil {
		// The legacy layer silently defaulted to position 1 here, mislabeling
		// every later bid of the round on a DB failure; real errors propagate now.
		return store.WrapError(op+": bid position", err)
	}

	cstReward, rewardErr := r.GlobStatsCstRewardForBidding(ctx)
	switch {
	case evt.BidCstRewardAmount != "-1":
		cstReward = evt.BidCstRewardAmount
	case rewardErr != nil:
		return fmt.Errorf("%s: cst_reward_for_bidding unset in cg_glob_stats (process admin events or restart ETL for chain sync): %w",
			op, rewardErr)
	case cstReward == "" || cstReward == "0":
		return fmt.Errorf("%s: cst_reward_for_bidding unset in cg_glob_stats (process admin events or restart ETL for chain sync): value=%q",
			op, cstReward)
	}

	// ETH and RandomWalk bids carry no CST price; CST bids no ETH price.
	ethPrice := evt.EthPrice
	cstPrice := evt.CstPrice
	if evt.BidType == 2 {
		ethPrice = "-1"
	} else {
		cstPrice = "-1"
	}

	query := "INSERT INTO cg_bid(" +
		"evtlog_id,block_num,time_stamp,tx_id,contract_aid," +
		"bidder_aid,rwalk_nft_id,eth_price,cst_price,cst_reward,bid_cst_reward_amount,cst_dutch_auction_duration,prize_time,msg,round_num,bid_type,bid_position" +
		") VALUES($1,$2,TO_TIMESTAMP($3),$4,$5,$6,$7,$8,$9,$10,$11,$12,TO_TIMESTAMP($13),$14,$15,$16,$17)"
	_, err = r.q(ctx).Exec(ctx, query,
		evt.EvtId,
		evt.BlockNum,
		evt.TimeStamp,
		evt.TxId,
		contractAid,
		bidderAid,
		evt.RandomWalkTokenId,
		ethPrice,
		cstPrice,
		cstReward,
		evt.BidCstRewardAmount,
		evt.CstDutchAuctionDuration,
		evt.PrizeTime,
		evt.Message,
		evt.RoundNum,
		evt.BidType,
		bidPosition,
	)
	return store.WrapError(op, err)
}

// InsertRoundStarted records a FirstBidPlacedInRound event.
func (r *Repo) InsertRoundStarted(ctx context.Context, evt *cgmodel.CGRoundStarted) error {
	const op = "insert into cg_first_bid"
	contractAid, err := r.addrID(ctx, evt.Contract, evt.BlockNum, evt.TxId)
	if err != nil {
		return store.WrapError(op, err)
	}
	query := "INSERT INTO cg_first_bid(" +
		"evtlog_id,block_num,tx_id,time_stamp,contract_aid," +
		"round_num,start_ts" +
		") VALUES($1,$2,$3,TO_TIMESTAMP($4),$5,$6,$7)"
	_, err = r.q(ctx).Exec(ctx, query,
		evt.EvtId,
		evt.BlockNum,
		evt.TxId,
		evt.TimeStamp,
		contractAid,
		evt.RoundNum,
		evt.StartTimestamp,
	)
	return store.WrapError(op, err)
}

// Donations.

// InsertEthDonation records an EthDonated event.
func (r *Repo) InsertEthDonation(ctx context.Context, evt *cgmodel.CGDonationEvent) error {
	const op = "insert into cg_eth_donated"
	contractAid, err := r.addrID(ctx, evt.ContractAddr, 0, 0)
	if err != nil {
		return store.WrapError(op, err)
	}
	donorAid, err := r.addrID(ctx, evt.DonorAddr, 0, 0)
	if err != nil {
		return store.WrapError(op, err)
	}
	query := "INSERT INTO cg_eth_donated(" +
		"evtlog_id,block_num,time_stamp,tx_id,contract_aid," +
		"donor_aid,round_num,amount" +
		") VALUES($1,$2,TO_TIMESTAMP($3),$4,$5,$6,$7,$8)"
	_, err = r.q(ctx).Exec(ctx, query,
		evt.EvtId,
		evt.BlockNum,
		evt.TimeStamp,
		evt.TxId,
		contractAid,
		donorAid,
		evt.RoundNum,
		evt.Amount,
	)
	return store.WrapError(op, err)
}

// InsertEthDonationWithInfo records an EthDonatedWithInfo event; the JSON
// payload read from the contract follows via InsertDonationJSON.
func (r *Repo) InsertEthDonationWithInfo(ctx context.Context, evt *cgmodel.CGDonationWithInfoEvent) error {
	const op = "insert into cg_eth_donated_wi"
	contractAid, err := r.addrID(ctx, evt.ContractAddr, 0, 0)
	if err != nil {
		return store.WrapError(op, err)
	}
	donorAid, err := r.addrID(ctx, evt.DonorAddr, 0, 0)
	if err != nil {
		return store.WrapError(op, err)
	}
	query := "INSERT INTO cg_eth_donated_wi(" +
		"evtlog_id,block_num,time_stamp,tx_id,contract_aid," +
		"donor_aid,round_num,record_id,amount" +
		") VALUES($1,$2,TO_TIMESTAMP($3),$4,$5,$6,$7,$8,$9)"
	_, err = r.q(ctx).Exec(ctx, query,
		evt.EvtId,
		evt.BlockNum,
		evt.TimeStamp,
		evt.TxId,
		contractAid,
		donorAid,
		evt.RoundNum,
		evt.RecordId,
		evt.Amount,
	)
	return store.WrapError(op, err)
}

// InsertDonationJSON records the JSON payload of an EthDonatedWithInfo
// donation, keyed by the contract-side record id (FK to cg_eth_donated_wi
// with ON DELETE CASCADE, so re-processing the parent event replaces it).
func (r *Repo) InsertDonationJSON(ctx context.Context, recordID int64, data string) error {
	query := "INSERT INTO cg_donation_json(record_id,data) VALUES($1,$2)"
	_, err := r.q(ctx).Exec(ctx, query, recordID, data)
	return store.WrapError("insert into cg_donation_json", err)
}

// InsertDonationReceived records a CharityWallet DonationReceived event.
func (r *Repo) InsertDonationReceived(ctx context.Context, evt *cgmodel.CGDonationReceivedEvent) error {
	const op = "insert into cg_donation_received"
	contractAid, err := r.addrID(ctx, evt.ContractAddr, 0, 0)
	if err != nil {
		return store.WrapError(op, err)
	}
	donorAid, err := r.addrID(ctx, evt.DonorAddr, 0, 0)
	if err != nil {
		return store.WrapError(op, err)
	}
	query := "INSERT INTO cg_donation_received(" +
		"evtlog_id,block_num,time_stamp,tx_id,contract_aid," +
		"donor_aid,amount,round_num" +
		") VALUES($1,$2,TO_TIMESTAMP($3),$4,$5,$6,$7,$8)"
	_, err = r.q(ctx).Exec(ctx, query,
		evt.EvtId,
		evt.BlockNum,
		evt.TimeStamp,
		evt.TxId,
		contractAid,
		donorAid,
		evt.Amount,
		evt.RoundNum,
	)
	return store.WrapError(op, err)
}

// InsertDonationSent records a CharityWallet FundsTransferredToCharity event.
func (r *Repo) InsertDonationSent(ctx context.Context, evt *cgmodel.CGDonationSentEvent) error {
	const op = "insert into cg_donation_sent"
	contractAid, err := r.addrID(ctx, evt.ContractAddr, 0, 0)
	if err != nil {
		return store.WrapError(op, err)
	}
	charityAid, err := r.addrID(ctx, evt.CharityAddr, 0, 0)
	if err != nil {
		return store.WrapError(op, err)
	}
	query := "INSERT INTO cg_donation_sent(" +
		"evtlog_id,block_num,time_stamp,tx_id,contract_aid," +
		"charity_aid,amount" +
		") VALUES($1,$2,TO_TIMESTAMP($3),$4,$5,$6,$7)"
	_, err = r.q(ctx).Exec(ctx, query,
		evt.EvtId,
		evt.BlockNum,
		evt.TimeStamp,
		evt.TxId,
		contractAid,
		charityAid,
		evt.Amount,
	)
	return store.WrapError(op, err)
}

// InsertERC20Donation records a PrizesWallet TokenDonated event.
func (r *Repo) InsertERC20Donation(ctx context.Context, evt *cgmodel.CGERC20DonationEvent) error {
	const op = "insert into cg_erc20_donation"
	contractAid, err := r.addrID(ctx, evt.ContractAddr, 0, 0)
	if err != nil {
		return store.WrapError(op, err)
	}
	donorAid, err := r.addrID(ctx, evt.DonorAddr, 0, 0)
	if err != nil {
		return store.WrapError(op, err)
	}
	tokenAid, err := r.addrID(ctx, evt.TokenAddr, 0, 0)
	if err != nil {
		return store.WrapError(op, err)
	}
	query := "INSERT INTO cg_erc20_donation(" +
		"evtlog_id,block_num,time_stamp,tx_id,contract_aid," +
		"donor_aid,token_aid,round_num,bid_id,amount" +
		") VALUES($1,$2,TO_TIMESTAMP($3),$4,$5,$6,$7,$8,$9,$10)"
	_, err = r.q(ctx).Exec(ctx, query,
		evt.EvtId,
		evt.BlockNum,
		evt.TimeStamp,
		evt.TxId,
		contractAid,
		donorAid,
		tokenAid,
		evt.RoundNum,
		evt.BidId,
		evt.Amount,
	)
	return store.WrapError(op, err)
}

// InsertNFTDonation records a PrizesWallet NftDonated event.
func (r *Repo) InsertNFTDonation(ctx context.Context, evt *cgmodel.CGNFTDonationEvent) error {
	const op = "insert into cg_nft_donation"
	contractAid, err := r.addrID(ctx, evt.ContractAddr, 0, 0)
	if err != nil {
		return store.WrapError(op, err)
	}
	donorAid, err := r.addrID(ctx, evt.DonorAddr, 0, 0)
	if err != nil {
		return store.WrapError(op, err)
	}
	tokenAid, err := r.addrID(ctx, evt.TokenAddr, 0, 0)
	if err != nil {
		return store.WrapError(op, err)
	}
	query := "INSERT INTO cg_nft_donation(" +
		"evtlog_id,block_num,time_stamp,tx_id,contract_aid," +
		"donor_aid,token_aid,token_id,round_num,idx,bid_id,token_uri" +
		") VALUES($1,$2,TO_TIMESTAMP($3),$4,$5,$6,$7,$8,$9,$10,$11,$12)"
	_, err = r.q(ctx).Exec(ctx, query,
		evt.EvtId,
		evt.BlockNum,
		evt.TimeStamp,
		evt.TxId,
		contractAid,
		donorAid,
		tokenAid,
		evt.TokenId,
		evt.RoundNum,
		evt.Index,
		evt.BidId,
		evt.NFTTokenURI,
	)
	return store.WrapError(op, err)
}

// InsertDonatedTokenClaim records a PrizesWallet DonatedTokenClaimed event.
func (r *Repo) InsertDonatedTokenClaim(ctx context.Context, evt *cgmodel.CGDonatedTokenClaimed) error {
	const op = "insert into cg_donated_tok_claimed"
	contractAid, err := r.addrID(ctx, evt.ContractAddr, 0, 0)
	if err != nil {
		return store.WrapError(op, err)
	}
	tokenAid, err := r.addrID(ctx, evt.TokenAddr, 0, 0)
	if err != nil {
		return store.WrapError(op, err)
	}
	winnerAid, err := r.addrID(ctx, evt.BeneficiaryAddr, 0, 0)
	if err != nil {
		return store.WrapError(op, err)
	}
	query := "INSERT INTO cg_donated_tok_claimed(" +
		"evtlog_id,block_num,time_stamp,tx_id,contract_aid," +
		"round_num,idx,token_aid,winner_aid,amount" +
		") VALUES($1,$2,TO_TIMESTAMP($3),$4,$5,$6,$7,$8,$9,$10)"
	_, err = r.q(ctx).Exec(ctx, query,
		evt.EvtId,
		evt.BlockNum,
		evt.TimeStamp,
		evt.TxId,
		contractAid,
		evt.RoundNum,
		evt.Index,
		tokenAid,
		winnerAid,
		evt.Amount,
	)
	return store.WrapError(op, err)
}

// InsertDonatedNFTClaim records a PrizesWallet DonatedNftClaimed event.
func (r *Repo) InsertDonatedNFTClaim(ctx context.Context, evt *cgmodel.CGDonatedNFTClaimed) error {
	const op = "insert into cg_donated_nft_claimed"
	contractAid, err := r.addrID(ctx, evt.ContractAddr, 0, 0)
	if err != nil {
		return store.WrapError(op, err)
	}
	tokenAid, err := r.addrID(ctx, evt.TokenAddr, 0, 0)
	if err != nil {
		return store.WrapError(op, err)
	}
	winnerAid, err := r.addrID(ctx, evt.BeneficiaryAddr, 0, 0)
	if err != nil {
		return store.WrapError(op, err)
	}
	query := "INSERT INTO cg_donated_nft_claimed(" +
		"evtlog_id,block_num,time_stamp,tx_id,contract_aid," +
		"round_num,idx,token_aid,winner_aid,token_id" +
		") VALUES($1,$2,TO_TIMESTAMP($3),$4,$5,$6,$7,$8,$9,$10)"
	_, err = r.q(ctx).Exec(ctx, query,
		evt.EvtId,
		evt.BlockNum,
		evt.TimeStamp,
		evt.TxId,
		contractAid,
		evt.RoundNum,
		evt.Index,
		tokenAid,
		winnerAid,
		evt.TokenId,
	)
	return store.WrapError(op, err)
}

// InsertFundsToCharity records a game FundsTransferredToCharity event.
func (r *Repo) InsertFundsToCharity(ctx context.Context, evt *cgmodel.CGFundsToCharity) error {
	const op = "insert into cg_funds_to_charity"
	contractAid, err := r.addrID(ctx, evt.Contract, evt.BlockNum, evt.TxId)
	if err != nil {
		return store.WrapError(op, err)
	}
	charityAid, err := r.addrID(ctx, evt.CharityAddr, evt.BlockNum, evt.TxId)
	if err != nil {
		return store.WrapError(op, err)
	}
	query := "INSERT INTO cg_funds_to_charity(" +
		"evtlog_id,block_num,tx_id,time_stamp,contract_aid," +
		"charity_aid,amount" +
		") VALUES($1,$2,$3,TO_TIMESTAMP($4),$5,$6,$7)"
	_, err = r.q(ctx).Exec(ctx, query,
		evt.EvtId,
		evt.BlockNum,
		evt.TxId,
		evt.TimeStamp,
		contractAid,
		charityAid,
		evt.Amount,
	)
	return store.WrapError(op, err)
}

// Tokens.

// InsertTokenName records an NftNameChanged event.
func (r *Repo) InsertTokenName(ctx context.Context, evt *cgmodel.CGTokenNameEvent) error {
	const op = "insert into cg_token_name"
	contractAid, err := r.addrID(ctx, evt.ContractAddr, 0, 0)
	if err != nil {
		return store.WrapError(op, err)
	}
	query := "INSERT INTO cg_token_name(" +
		"evtlog_id,block_num,time_stamp,tx_id,contract_aid," +
		"token_id,token_name" +
		") VALUES($1,$2,TO_TIMESTAMP($3),$4,$5,$6,$7)"
	_, err = r.q(ctx).Exec(ctx, query,
		evt.EvtId,
		evt.BlockNum,
		evt.TimeStamp,
		evt.TxId,
		contractAid,
		evt.TokenId,
		evt.TokenName,
	)
	return store.WrapError(op, err)
}

// InsertMint records an NftMinted event; the owner starts as the current
// owner (cur_owner_aid tracks later transfers).
func (r *Repo) InsertMint(ctx context.Context, evt *cgmodel.CGMintEvent) error {
	const op = "insert into cg_mint_event"
	contractAid, err := r.addrID(ctx, evt.ContractAddr, 0, 0)
	if err != nil {
		return store.WrapError(op, err)
	}
	ownerAid, err := r.addrID(ctx, evt.OwnerAddr, 0, 0)
	if err != nil {
		return store.WrapError(op, err)
	}
	query := "INSERT INTO cg_mint_event(" +
		"evtlog_id,block_num,time_stamp,tx_id,contract_aid," +
		"owner_aid,cur_owner_aid,token_id,round_num,seed" +
		") VALUES($1,$2,TO_TIMESTAMP($3),$4,$5,$6,$7,$8,$9,$10)"
	_, err = r.q(ctx).Exec(ctx, query,
		evt.EvtId,
		evt.BlockNum,
		evt.TimeStamp,
		evt.TxId,
		contractAid,
		ownerAid,
		ownerAid,
		evt.TokenId,
		evt.RoundNum,
		evt.Seed,
	)
	return store.WrapError(op, err)
}

// InsertCosmicSignatureTransfer records an ERC721 Transfer of a
// CosmicSignature NFT (otype: 1 = mint, 2 = burn, 0 = regular transfer).
func (r *Repo) InsertCosmicSignatureTransfer(ctx context.Context, evt *cgmodel.CGERC721Transfer) error {
	const op = "insert into cg_erc721_transfer"
	contractAid, err := r.addrID(ctx, evt.Contract, evt.BlockNum, evt.TxId)
	if err != nil {
		return store.WrapError(op, err)
	}
	fromAid, err := r.addrID(ctx, evt.From, evt.BlockNum, evt.TxId)
	if err != nil {
		return store.WrapError(op, err)
	}
	toAid, err := r.addrID(ctx, evt.To, evt.BlockNum, evt.TxId)
	if err != nil {
		return store.WrapError(op, err)
	}
	otype := transferOType(evt.From, evt.To)
	query := "INSERT INTO cg_erc721_transfer(" +
		"evtlog_id,block_num,tx_id,time_stamp,contract_aid," +
		"token_id,from_aid,to_aid,otype" +
		") VALUES($1,$2,$3,TO_TIMESTAMP($4),$5,$6,$7,$8,$9)"
	_, err = r.q(ctx).Exec(ctx, query,
		evt.EvtId,
		evt.BlockNum,
		evt.TxId,
		evt.TimeStamp,
		contractAid,
		evt.TokenId,
		fromAid,
		toAid,
		otype,
	)
	return store.WrapError(op, err)
}

// InsertCosmicTokenTransfer records an ERC20 Transfer of CosmicToken
// (otype: 1 = mint, 2 = burn, 0 = regular transfer).
func (r *Repo) InsertCosmicTokenTransfer(ctx context.Context, evt *cgmodel.CGERC20Transfer) error {
	const op = "insert into cg_erc20_transfer"
	contractAid, err := r.addrID(ctx, evt.Contract, evt.BlockNum, evt.TxId)
	if err != nil {
		return store.WrapError(op, err)
	}
	fromAid, err := r.addrID(ctx, evt.From, evt.BlockNum, evt.TxId)
	if err != nil {
		return store.WrapError(op, err)
	}
	toAid, err := r.addrID(ctx, evt.To, evt.BlockNum, evt.TxId)
	if err != nil {
		return store.WrapError(op, err)
	}
	otype := transferOType(evt.From, evt.To)
	query := "INSERT INTO cg_erc20_transfer(" +
		"evtlog_id,block_num,tx_id,time_stamp,contract_aid," +
		"value,from_aid,to_aid,otype" +
		") VALUES($1,$2,$3,TO_TIMESTAMP($4),$5,$6,$7,$8,$9)"
	_, err = r.q(ctx).Exec(ctx, query,
		evt.EvtId,
		evt.BlockNum,
		evt.TxId,
		evt.TimeStamp,
		contractAid,
		evt.Value,
		fromAid,
		toAid,
		otype,
	)
	return store.WrapError(op, err)
}

const zeroAddress = "0x0000000000000000000000000000000000000000"

// transferOType classifies a token transfer: 1 = mint (from zero address),
// 2 = burn (to zero address), 0 = regular transfer.
func transferOType(from, to string) int {
	switch {
	case from == zeroAddress:
		return 1
	case to == zeroAddress:
		return 2
	default:
		return 0
	}
}

// Prizes.

// InsertPrizeDeposit records a PrizesWallet EthReceived event.
func (r *Repo) InsertPrizeDeposit(ctx context.Context, evt *cgmodel.CGPrizesEthDeposit) error {
	const op = "insert into cg_prize_deposit"
	contractAid, err := r.addrID(ctx, evt.ContractAddr, 0, 0)
	if err != nil {
		return store.WrapError(op, err)
	}
	winnerAid, err := r.addrID(ctx, evt.WinnerAddr, 0, 0)
	if err != nil {
		return store.WrapError(op, err)
	}
	query := "INSERT INTO cg_prize_deposit(" +
		"evtlog_id,block_num,time_stamp,tx_id,contract_aid," +
		"winner_aid,round_num,winner_index,amount" +
		") VALUES($1,$2,TO_TIMESTAMP($3),$4,$5,$6,$7,$8,$9)"
	_, err = r.q(ctx).Exec(ctx, query,
		evt.EvtId,
		evt.BlockNum,
		evt.TimeStamp,
		evt.TxId,
		contractAid,
		winnerAid,
		evt.Round,
		evt.WinnerIndex,
		evt.Amount,
	)
	return store.WrapError(op, err)
}

// InsertPrizeWithdrawal records a PrizesWallet EthWithdrawn event.
func (r *Repo) InsertPrizeWithdrawal(ctx context.Context, evt *cgmodel.CGPrizesEthWithdrawal) error {
	const op = "insert into cg_prize_withdrawal"
	contractAid, err := r.addrID(ctx, evt.ContractAddr, 0, 0)
	if err != nil {
		return store.WrapError(op, err)
	}
	winnerAid, err := r.addrID(ctx, evt.WinnerAddr, 0, 0)
	if err != nil {
		return store.WrapError(op, err)
	}
	beneficiaryAid, err := r.addrID(ctx, evt.BeneficiaryAddr, 0, 0)
	if err != nil {
		return store.WrapError(op, err)
	}
	query := "INSERT INTO cg_prize_withdrawal(" +
		"evtlog_id,block_num,time_stamp,tx_id,contract_aid," +
		"round_num,winner_aid,beneficiary_aid,amount" +
		") VALUES($1,$2,TO_TIMESTAMP($3),$4,$5,$6,$7,$8,$9)"
	_, err = r.q(ctx).Exec(ctx, query,
		evt.EvtId,
		evt.BlockNum,
		evt.TimeStamp,
		evt.TxId,
		contractAid,
		evt.Round,
		winnerAid,
		beneficiaryAid,
		evt.Amount,
	)
	return store.WrapError(op, err)
}

// InsertRaffleNFTWinner records a RaffleWinnerPrizePaid event.
func (r *Repo) InsertRaffleNFTWinner(ctx context.Context, evt *cgmodel.CGRaffleNFTWinner) error {
	const op = "insert into cg_raffle_nft_prize"
	contractAid, err := r.addrID(ctx, evt.ContractAddr, 0, 0)
	if err != nil {
		return store.WrapError(op, err)
	}
	winnerAid, err := r.addrID(ctx, evt.WinnerAddr, 0, 0)
	if err != nil {
		return store.WrapError(op, err)
	}
	query := "INSERT INTO cg_raffle_nft_prize(" +
		"evtlog_id,block_num,time_stamp,tx_id,contract_aid," +
		"winner_aid,round_num,token_id,winner_idx,cst_amount,is_rwalk,is_staker" +
		") VALUES($1,$2,TO_TIMESTAMP($3),$4,$5,$6,$7,$8,$9,$10,$11,$12)"
	_, err = r.q(ctx).Exec(ctx, query,
		evt.EvtId,
		evt.BlockNum,
		evt.TimeStamp,
		evt.TxId,
		contractAid,
		winnerAid,
		evt.Round,
		evt.TokenId,
		evt.WinnerIndex,
		evt.CstAmount,
		evt.IsRandomWalk,
		evt.IsStaker,
	)
	return store.WrapError(op, err)
}

// InsertRaffleETHWinner records a RaffleWinnerBidderEthPrizeAllocated event.
func (r *Repo) InsertRaffleETHWinner(ctx context.Context, evt *cgmodel.CGRaffleETHWinner) error {
	const op = "insert into cg_raffle_eth_prize"
	contractAid, err := r.addrID(ctx, evt.ContractAddr, 0, 0)
	if err != nil {
		return store.WrapError(op, err)
	}
	winnerAid, err := r.addrID(ctx, evt.WinnerAddr, 0, 0)
	if err != nil {
		return store.WrapError(op, err)
	}
	query := "INSERT INTO cg_raffle_eth_prize(" +
		"evtlog_id,block_num,time_stamp,tx_id,contract_aid," +
		"winner_aid,round_num,winner_idx,amount" +
		") VALUES($1,$2,TO_TIMESTAMP($3),$4,$5,$6,$7,$8,$9)"
	_, err = r.q(ctx).Exec(ctx, query,
		evt.EvtId,
		evt.BlockNum,
		evt.TimeStamp,
		evt.TxId,
		contractAid,
		winnerAid,
		evt.Round,
		evt.WinnerIndex,
		evt.Amount,
	)
	return store.WrapError(op, err)
}

// InsertEnduranceWinner records an EnduranceChampionPrizePaid event. The
// Solidity event emits no winner index — there is exactly one endurance
// champion per round.
func (r *Repo) InsertEnduranceWinner(ctx context.Context, evt *cgmodel.CGEnduranceWinner) error {
	const op = "insert into cg_endurance_prize"
	contractAid, err := r.addrID(ctx, evt.ContractAddr, 0, 0)
	if err != nil {
		return store.WrapError(op, err)
	}
	winnerAid, err := r.addrID(ctx, evt.WinnerAddr, 0, 0)
	if err != nil {
		return store.WrapError(op, err)
	}
	query := "INSERT INTO cg_endurance_prize(" +
		"evtlog_id,block_num,time_stamp,tx_id,contract_aid," +
		"winner_aid,round_num,erc721_token_id,erc20_amount" +
		") VALUES($1,$2,TO_TIMESTAMP($3),$4,$5,$6,$7,$8,$9)"
	_, err = r.q(ctx).Exec(ctx, query,
		evt.EvtId,
		evt.BlockNum,
		evt.TimeStamp,
		evt.TxId,
		contractAid,
		winnerAid,
		evt.Round,
		evt.Erc721TokenId,
		evt.Erc20Amount,
	)
	return store.WrapError(op, err)
}

// InsertLastCstBidderWinner records a LastCstBidderPrizePaid event (one per
// round, no winner index in the Solidity event).
func (r *Repo) InsertLastCstBidderWinner(ctx context.Context, evt *cgmodel.CGLastBidderWinner) error {
	const op = "insert into cg_lastcst_prize"
	contractAid, err := r.addrID(ctx, evt.ContractAddr, 0, 0)
	if err != nil {
		return store.WrapError(op, err)
	}
	winnerAid, err := r.addrID(ctx, evt.WinnerAddr, 0, 0)
	if err != nil {
		return store.WrapError(op, err)
	}
	query := "INSERT INTO cg_lastcst_prize(" +
		"evtlog_id,block_num,time_stamp,tx_id,contract_aid," +
		"winner_aid,round_num,erc721_token_id,erc20_amount" +
		") VALUES($1,$2,TO_TIMESTAMP($3),$4,$5,$6,$7,$8,$9)"
	_, err = r.q(ctx).Exec(ctx, query,
		evt.EvtId,
		evt.BlockNum,
		evt.TimeStamp,
		evt.TxId,
		contractAid,
		winnerAid,
		evt.Round,
		evt.Erc721TokenId,
		evt.Erc20Amount,
	)
	return store.WrapError(op, err)
}

// InsertChronoWarrior records a ChronoWarriorPrizePaid event.
func (r *Repo) InsertChronoWarrior(ctx context.Context, evt *cgmodel.CGChronoWarrior) error {
	const op = "insert into cg_chrono_warrior_prize"
	contractAid, err := r.addrID(ctx, evt.ContractAddr, 0, 0)
	if err != nil {
		return store.WrapError(op, err)
	}
	winnerAid, err := r.addrID(ctx, evt.WinnerAddr, 0, 0)
	if err != nil {
		return store.WrapError(op, err)
	}
	query := "INSERT INTO cg_chrono_warrior_prize(" +
		"evtlog_id,block_num,time_stamp,tx_id,contract_aid," +
		"winner_aid,round_num,winner_index,eth_amount,cst_amount,nft_id" +
		") VALUES($1,$2,TO_TIMESTAMP($3),$4,$5,$6,$7,$8,$9,$10,$11)"
	_, err = r.q(ctx).Exec(ctx, query,
		evt.EvtId,
		evt.BlockNum,
		evt.TimeStamp,
		evt.TxId,
		contractAid,
		winnerAid,
		evt.Round,
		evt.WinnerIndex,
		evt.EthAmount,
		evt.CstAmount,
		evt.NftId,
	)
	return store.WrapError(op, err)
}

// InsertFundTransferFailed records a FundTransferFailed event.
func (r *Repo) InsertFundTransferFailed(ctx context.Context, evt *cgmodel.CGFundTransferFailed) error {
	const op = "insert into cg_fund_transf_err"
	contractAid, err := r.addrID(ctx, evt.Contract, evt.BlockNum, evt.TxId)
	if err != nil {
		return store.WrapError(op, err)
	}
	destinationAid, err := r.addrID(ctx, evt.Destination, evt.BlockNum, evt.TxId)
	if err != nil {
		return store.WrapError(op, err)
	}
	query := "INSERT INTO cg_fund_transf_err(" +
		"evtlog_id,block_num,tx_id,time_stamp,contract_aid," +
		"destination_aid,amount" +
		") VALUES($1,$2,$3,TO_TIMESTAMP($4),$5,$6,$7)"
	_, err = r.q(ctx).Exec(ctx, query,
		evt.EvtId,
		evt.BlockNum,
		evt.TxId,
		evt.TimeStamp,
		contractAid,
		destinationAid,
		evt.Amount,
	)
	return store.WrapError(op, err)
}

// InsertERC20TransferFailed records an ERC20TransferFailed event.
func (r *Repo) InsertERC20TransferFailed(ctx context.Context, evt *cgmodel.CGErc20TransferFailed) error {
	const op = "insert into cg_erc20_transf_err"
	contractAid, err := r.addrID(ctx, evt.Contract, evt.BlockNum, evt.TxId)
	if err != nil {
		return store.WrapError(op, err)
	}
	destinationAid, err := r.addrID(ctx, evt.Destination, evt.BlockNum, evt.TxId)
	if err != nil {
		return store.WrapError(op, err)
	}
	query := "INSERT INTO cg_erc20_transf_err(" +
		"evtlog_id,block_num,tx_id,time_stamp,contract_aid," +
		"destination_aid,amount" +
		") VALUES($1,$2,$3,TO_TIMESTAMP($4),$5,$6,$7)"
	_, err = r.q(ctx).Exec(ctx, query,
		evt.EvtId,
		evt.BlockNum,
		evt.TxId,
		evt.TimeStamp,
		contractAid,
		destinationAid,
		evt.Amount,
	)
	return store.WrapError(op, err)
}

// Staking.

// InsertNftStakedCST records a CST-wallet NftStaked event.
func (r *Repo) InsertNftStakedCST(ctx context.Context, evt *cgmodel.CGNftStakedCst) error {
	const op = "insert into cg_nft_staked_cst"
	contractAid, err := r.addrID(ctx, evt.ContractAddr, evt.BlockNum, evt.TxId)
	if err != nil {
		return store.WrapError(op, err)
	}
	stakerAid, err := r.addrID(ctx, evt.StakerAddress, evt.BlockNum, evt.TxId)
	if err != nil {
		return store.WrapError(op, err)
	}
	query := "INSERT INTO cg_nft_staked_cst(" +
		"evtlog_id,block_num,tx_id,time_stamp,contract_aid," +
		"action_id,token_id,num_staked_nfts,reward_per_staker,staker_aid" +
		") VALUES($1,$2,$3,TO_TIMESTAMP($4),$5,$6,$7,$8,$9,$10)"
	_, err = r.q(ctx).Exec(ctx, query,
		evt.EvtId,
		evt.BlockNum,
		evt.TxId,
		evt.TimeStamp,
		contractAid,
		evt.ActionId,
		evt.NftId,
		evt.NumStakedNfts,
		evt.RewardPerStaker,
		stakerAid,
	)
	return store.WrapError(op, err)
}

// InsertNftStakedRWalk records a RandomWalk-wallet NftStaked event.
func (r *Repo) InsertNftStakedRWalk(ctx context.Context, evt *cgmodel.CGNftStakedRWalk) error {
	const op = "insert into cg_nft_staked_rwalk"
	contractAid, err := r.addrID(ctx, evt.ContractAddr, evt.BlockNum, evt.TxId)
	if err != nil {
		return store.WrapError(op, err)
	}
	stakerAid, err := r.addrID(ctx, evt.StakerAddress, evt.BlockNum, evt.TxId)
	if err != nil {
		return store.WrapError(op, err)
	}
	query := "INSERT INTO cg_nft_staked_rwalk(" +
		"evtlog_id,block_num,tx_id,time_stamp,contract_aid," +
		"action_id,token_id,num_staked_nfts,staker_aid" +
		") VALUES($1,$2,$3,TO_TIMESTAMP($4),$5,$6,$7,$8,$9)"
	_, err = r.q(ctx).Exec(ctx, query,
		evt.EvtId,
		evt.BlockNum,
		evt.TxId,
		evt.TimeStamp,
		contractAid,
		evt.ActionId,
		evt.NftId,
		evt.NumStakedNfts,
		stakerAid,
	)
	return store.WrapError(op, err)
}

// InsertNftUnstakedCST records a CST-wallet NftUnstaked event (carries the
// staker's reward accounting).
func (r *Repo) InsertNftUnstakedCST(ctx context.Context, evt *cgmodel.CGNftUnstakedCst) error {
	const op = "insert into cg_nft_unstaked_cst"
	contractAid, err := r.addrID(ctx, evt.ContractAddr, evt.BlockNum, evt.TxId)
	if err != nil {
		return store.WrapError(op, err)
	}
	stakerAid, err := r.addrID(ctx, evt.StakerAddress, evt.BlockNum, evt.TxId)
	if err != nil {
		return store.WrapError(op, err)
	}
	query := "INSERT INTO cg_nft_unstaked_cst(" +
		"evtlog_id,block_num,tx_id,time_stamp,contract_aid," +
		"action_id,token_id,num_staked_nfts,staker_aid,reward,reward_per_tok,action_counter" +
		") VALUES($1,$2,$3,TO_TIMESTAMP($4),$5,$6,$7,$8,$9,$10,$11,$12)"
	_, err = r.q(ctx).Exec(ctx, query,
		evt.EvtId,
		evt.BlockNum,
		evt.TxId,
		evt.TimeStamp,
		contractAid,
		evt.ActionId,
		evt.NftId,
		evt.NumStakedNfts,
		stakerAid,
		evt.RewardAmount,
		evt.RewardPerToken,
		evt.ActionCounter,
	)
	return store.WrapError(op, err)
}

// InsertNftUnstakedRWalk records a RandomWalk-wallet NftUnstaked event.
func (r *Repo) InsertNftUnstakedRWalk(ctx context.Context, evt *cgmodel.CGNftUnstakedRWalk) error {
	const op = "insert into cg_nft_unstaked_rwalk"
	contractAid, err := r.addrID(ctx, evt.ContractAddr, evt.BlockNum, evt.TxId)
	if err != nil {
		return store.WrapError(op, err)
	}
	stakerAid, err := r.addrID(ctx, evt.StakerAddress, evt.BlockNum, evt.TxId)
	if err != nil {
		return store.WrapError(op, err)
	}
	query := "INSERT INTO cg_nft_unstaked_rwalk(" +
		"evtlog_id,block_num,tx_id,time_stamp,contract_aid," +
		"action_id,token_id,num_staked_nfts,staker_aid" +
		") VALUES($1,$2,$3,TO_TIMESTAMP($4),$5,$6,$7,$8,$9)"
	_, err = r.q(ctx).Exec(ctx, query,
		evt.EvtId,
		evt.BlockNum,
		evt.TxId,
		evt.TimeStamp,
		contractAid,
		evt.ActionId,
		evt.NftId,
		evt.NumStakedNfts,
		stakerAid,
	)
	return store.WrapError(op, err)
}

// InsertStakingEthDeposit records a staking-wallet EthDepositReceived event.
func (r *Repo) InsertStakingEthDeposit(ctx context.Context, evt *cgmodel.CGEthDeposit) error {
	const op = "insert into cg_staking_eth_deposit"
	contractAid, err := r.addrID(ctx, evt.ContractAddr, evt.BlockNum, evt.TxId)
	if err != nil {
		return store.WrapError(op, err)
	}
	query := "INSERT INTO cg_staking_eth_deposit(" +
		"evtlog_id,block_num,tx_id,time_stamp,contract_aid," +
		"deposit_time,round_num,deposit_id,num_staked_nfts,deposit_amount,amount_per_token,modulo,accum_modulo" +
		") VALUES($1,$2,$3,TO_TIMESTAMP($4),$5,TO_TIMESTAMP($6),$7,$8,$9,$10,$11,$12,$13)"
	_, err = r.q(ctx).Exec(ctx, query,
		evt.EvtId,
		evt.BlockNum,
		evt.TxId,
		evt.TimeStamp,
		contractAid,
		evt.DepositTime,
		evt.RoundNum,
		evt.DepositId,
		evt.NumStakedNfts,
		evt.Amount,
		evt.AmountPerStaker,
		evt.Modulo,
		evt.AccumModulo,
	)
	return store.WrapError(op, err)
}

// Marketing.

// InsertMarketingRewardPaid records a MarketingWallet RewardPaid event.
func (r *Repo) InsertMarketingRewardPaid(ctx context.Context, evt *cgmodel.CGMarketingRewardPaid) error {
	const op = "insert into cg_mkt_reward"
	contractAid, err := r.addrID(ctx, evt.ContractAddr, evt.BlockNum, evt.TxId)
	if err != nil {
		return store.WrapError(op, err)
	}
	marketerAid, err := r.addrID(ctx, evt.Marketer, evt.BlockNum, evt.TxId)
	if err != nil {
		return store.WrapError(op, err)
	}
	query := "INSERT INTO cg_mkt_reward(" +
		"evtlog_id,block_num,tx_id,time_stamp,contract_aid," +
		"amount,marketer_aid" +
		") VALUES($1,$2,$3,TO_TIMESTAMP($4),$5,$6,$7)"
	_, err = r.q(ctx).Exec(ctx, query,
		evt.EvtId,
		evt.BlockNum,
		evt.TxId,
		evt.TimeStamp,
		contractAid,
		evt.Amount,
		marketerAid,
	)
	return store.WrapError(op, err)
}

// Admin: percentages and counts.

// InsertCharityPercentageChange records a CharityEthDonationAmountPercentageChanged event.
func (r *Repo) InsertCharityPercentageChange(ctx context.Context, evt *cgmodel.CGCharityPercentageChanged) error {
	contractAid, err := r.addrID(ctx, evt.Contract, evt.BlockNum, evt.TxId)
	if err != nil {
		return store.WrapError("insert into cg_adm_charity_pcent", err)
	}
	return r.insertAdminValue(ctx, "cg_adm_charity_pcent", "percentage",
		evt.EvtId, evt.BlockNum, evt.TxId, evt.TimeStamp, contractAid, evt.NewCharityPercentage)
}

// InsertPrizePercentageChange records a MainEthPrizeAmountPercentageChanged event.
func (r *Repo) InsertPrizePercentageChange(ctx context.Context, evt *cgmodel.CGPrizePercentageChanged) error {
	contractAid, err := r.addrID(ctx, evt.Contract, evt.BlockNum, evt.TxId)
	if err != nil {
		return store.WrapError("insert into cg_adm_main_prize_pcent", err)
	}
	return r.insertAdminValue(ctx, "cg_adm_main_prize_pcent", "percentage",
		evt.EvtId, evt.BlockNum, evt.TxId, evt.TimeStamp, contractAid, evt.NewPrizePercentage)
}

// InsertRafflePercentageChange records a RaffleTotalEthPrizeAmountForBiddersPercentageChanged event.
func (r *Repo) InsertRafflePercentageChange(ctx context.Context, evt *cgmodel.CGRafflePercentageChanged) error {
	contractAid, err := r.addrID(ctx, evt.Contract, evt.BlockNum, evt.TxId)
	if err != nil {
		return store.WrapError("insert into cg_adm_raffle_pcent", err)
	}
	return r.insertAdminValue(ctx, "cg_adm_raffle_pcent", "percentage",
		evt.EvtId, evt.BlockNum, evt.TxId, evt.TimeStamp, contractAid, evt.NewRafflePercentage)
}

// InsertStakingPercentageChange records a CosmicSignatureNftStakingTotalEthRewardAmountPercentageChanged event.
func (r *Repo) InsertStakingPercentageChange(ctx context.Context, evt *cgmodel.CGStakingPercentageChanged) error {
	contractAid, err := r.addrID(ctx, evt.Contract, evt.BlockNum, evt.TxId)
	if err != nil {
		return store.WrapError("insert into cg_adm_stake_pcent", err)
	}
	return r.insertAdminValue(ctx, "cg_adm_stake_pcent", "percentage",
		evt.EvtId, evt.BlockNum, evt.TxId, evt.TimeStamp, contractAid, evt.NewStakingPercentage)
}

// InsertChronoPercentageChange records a ChronoWarriorEthPrizeAmountPercentageChanged event.
func (r *Repo) InsertChronoPercentageChange(ctx context.Context, evt *cgmodel.CGChronoPercentageChanged) error {
	contractAid, err := r.addrID(ctx, evt.Contract, evt.BlockNum, evt.TxId)
	if err != nil {
		return store.WrapError("insert into cg_adm_chrono_pcent", err)
	}
	return r.insertAdminValue(ctx, "cg_adm_chrono_pcent", "percentage",
		evt.EvtId, evt.BlockNum, evt.TxId, evt.TimeStamp, contractAid, evt.NewChronoPercentage)
}

// InsertNumRaffleETHWinnersBiddingChange records a NumRaffleEthPrizesForBiddersChanged event.
func (r *Repo) InsertNumRaffleETHWinnersBiddingChange(ctx context.Context, evt *cgmodel.CGNumRaffleETHWinnersBiddingChanged) error {
	contractAid, err := r.addrID(ctx, evt.Contract, evt.BlockNum, evt.TxId)
	if err != nil {
		return store.WrapError("insert into cg_adm_raf_eth_bidding", err)
	}
	return r.insertAdminValue(ctx, "cg_adm_raf_eth_bidding", "num_winners",
		evt.EvtId, evt.BlockNum, evt.TxId, evt.TimeStamp, contractAid, evt.NewNumRaffleETHWinnersBidding)
}

// InsertNumRaffleNFTWinnersBiddingChange records a NumRaffleCosmicSignatureNftsForBiddersChanged event.
func (r *Repo) InsertNumRaffleNFTWinnersBiddingChange(ctx context.Context, evt *cgmodel.CGNumRaffleNFTWinnersBiddingChanged) error {
	contractAid, err := r.addrID(ctx, evt.Contract, evt.BlockNum, evt.TxId)
	if err != nil {
		return store.WrapError("insert into cg_adm_raf_nft_bidding", err)
	}
	return r.insertAdminValue(ctx, "cg_adm_raf_nft_bidding", "num_winners",
		evt.EvtId, evt.BlockNum, evt.TxId, evt.TimeStamp, contractAid, evt.NewNumRaffleNFTWinnersBidding)
}

// InsertNumRaffleNFTWinnersStakingRWalkChange records a
// NumRaffleCosmicSignatureNftsForRandomWalkNftStakersChanged event.
func (r *Repo) InsertNumRaffleNFTWinnersStakingRWalkChange(ctx context.Context, evt *cgmodel.CGNumRaffleNFTWinnersStakingRWalkChanged) error {
	contractAid, err := r.addrID(ctx, evt.Contract, evt.BlockNum, evt.TxId)
	if err != nil {
		return store.WrapError("insert into cg_adm_raf_nft_staking_rwalk", err)
	}
	return r.insertAdminValue(ctx, "cg_adm_raf_nft_staking_rwalk", "num_winners",
		evt.EvtId, evt.BlockNum, evt.TxId, evt.TimeStamp, contractAid, evt.NewNumRaffleNFTWinnersStakingRWalk)
}

// Admin: address changes.

// InsertCharityReceiverChange records a CharityWallet CharityAddressChanged
// event (who receives the charity funds).
func (r *Repo) InsertCharityReceiverChange(ctx context.Context, evt *cgmodel.CGCharityUpdatedEvent) error {
	const op = "insert into cg_charity_receiver_changed"
	contractAid, err := r.addrID(ctx, evt.ContractAddr, 0, 0)
	if err != nil {
		return store.WrapError(op, err)
	}
	newCharityAid, err := r.addrID(ctx, evt.NewCharityAddr, 0, 0)
	if err != nil {
		return store.WrapError(op, err)
	}
	query := "INSERT INTO cg_charity_receiver_changed(" +
		"evtlog_id,block_num,time_stamp,tx_id,contract_aid," +
		"charity_aid" +
		") VALUES($1,$2,TO_TIMESTAMP($3),$4,$5,$6)"
	_, err = r.q(ctx).Exec(ctx, query,
		evt.EvtId,
		evt.BlockNum,
		evt.TimeStamp,
		evt.TxId,
		contractAid,
		newCharityAid,
	)
	return store.WrapError(op, err)
}

// InsertCharityWalletAddressChange records a game CharityAddressChanged
// event (which CharityWallet contract the game uses).
func (r *Repo) InsertCharityWalletAddressChange(ctx context.Context, evt *cgmodel.CGCharityAddressChanged) error {
	const op = "insert into cg_adm_charity_wallet"
	contractAid, err := r.addrID(ctx, evt.Contract, evt.BlockNum, evt.TxId)
	if err != nil {
		return store.WrapError(op, err)
	}
	newCharityAid, err := r.addrID(ctx, evt.NewCharity, evt.BlockNum, evt.TxId)
	if err != nil {
		return store.WrapError(op, err)
	}
	return r.insertAdminValue(ctx, "cg_adm_charity_wallet", "new_charity_aid",
		evt.EvtId, evt.BlockNum, evt.TxId, evt.TimeStamp, contractAid, newCharityAid)
}

// InsertRandomWalkAddressChange records a RandomWalkNftAddressChanged event.
func (r *Repo) InsertRandomWalkAddressChange(ctx context.Context, evt *cgmodel.CGRandomWalkAddressChanged) error {
	const op = "insert into cg_adm_rwalk_addr"
	contractAid, err := r.addrID(ctx, evt.Contract, evt.BlockNum, evt.TxId)
	if err != nil {
		return store.WrapError(op, err)
	}
	newRWalkAid, err := r.addrID(ctx, evt.NewRandomWalk, evt.BlockNum, evt.TxId)
	if err != nil {
		return store.WrapError(op, err)
	}
	return r.insertAdminValue(ctx, "cg_adm_rwalk_addr", "new_rwalk_aid",
		evt.EvtId, evt.BlockNum, evt.TxId, evt.TimeStamp, contractAid, newRWalkAid)
}

// InsertPrizesWalletAddressChange records a PrizesWalletAddressChanged event.
func (r *Repo) InsertPrizesWalletAddressChange(ctx context.Context, evt *cgmodel.CGPrizeWalletAddressChanged) error {
	const op = "insert into cg_adm_prizes_wallet_addr"
	contractAid, err := r.addrID(ctx, evt.Contract, evt.BlockNum, evt.TxId)
	if err != nil {
		return store.WrapError(op, err)
	}
	newWalletAid, err := r.addrID(ctx, evt.NewPrizeWallet, evt.BlockNum, evt.TxId)
	if err != nil {
		return store.WrapError(op, err)
	}
	return r.insertAdminValue(ctx, "cg_adm_prizes_wallet_addr", "new_wallet_aid",
		evt.EvtId, evt.BlockNum, evt.TxId, evt.TimeStamp, contractAid, newWalletAid)
}

// InsertStakingWalletCSTAddressChange records a StakingWalletCosmicSignatureNftAddressChanged event.
func (r *Repo) InsertStakingWalletCSTAddressChange(ctx context.Context, evt *cgmodel.CGStakingWalletCSTAddressChanged) error {
	const op = "insert into cg_adm_staking_cst_addr"
	contractAid, err := r.addrID(ctx, evt.Contract, evt.BlockNum, evt.TxId)
	if err != nil {
		return store.WrapError(op, err)
	}
	newStakingAid, err := r.addrID(ctx, evt.NewStakingWalletCST, evt.BlockNum, evt.TxId)
	if err != nil {
		return store.WrapError(op, err)
	}
	return r.insertAdminValue(ctx, "cg_adm_staking_cst_addr", "new_staking_aid",
		evt.EvtId, evt.BlockNum, evt.TxId, evt.TimeStamp, contractAid, newStakingAid)
}

// InsertStakingWalletRWalkAddressChange records a StakingWalletRandomWalkNftAddressChanged event.
func (r *Repo) InsertStakingWalletRWalkAddressChange(ctx context.Context, evt *cgmodel.CGStakingWalletRWalkAddressChanged) error {
	const op = "insert into cg_adm_staking_rwalk_addr"
	contractAid, err := r.addrID(ctx, evt.Contract, evt.BlockNum, evt.TxId)
	if err != nil {
		return store.WrapError(op, err)
	}
	newStakingAid, err := r.addrID(ctx, evt.NewStakingWalletRWalk, evt.BlockNum, evt.TxId)
	if err != nil {
		return store.WrapError(op, err)
	}
	return r.insertAdminValue(ctx, "cg_adm_staking_rwalk_addr", "new_staking_aid",
		evt.EvtId, evt.BlockNum, evt.TxId, evt.TimeStamp, contractAid, newStakingAid)
}

// InsertMarketingWalletAddressChange records a MarketingWalletAddressChanged event.
func (r *Repo) InsertMarketingWalletAddressChange(ctx context.Context, evt *cgmodel.CGMarketingWalletAddressChanged) error {
	const op = "insert into cg_adm_marketing_addr"
	contractAid, err := r.addrID(ctx, evt.Contract, evt.BlockNum, evt.TxId)
	if err != nil {
		return store.WrapError(op, err)
	}
	newMarketingAid, err := r.addrID(ctx, evt.NewMarketingWallet, evt.BlockNum, evt.TxId)
	if err != nil {
		return store.WrapError(op, err)
	}
	return r.insertAdminValue(ctx, "cg_adm_marketing_addr", "new_marketing_aid",
		evt.EvtId, evt.BlockNum, evt.TxId, evt.TimeStamp, contractAid, newMarketingAid)
}

// InsertTreasurerAddressChange records a TreasurerAddressChanged event.
func (r *Repo) InsertTreasurerAddressChange(ctx context.Context, evt *cgmodel.CGTreasurerAddressChanged) error {
	const op = "insert into cg_adm_treasurer_addr"
	contractAid, err := r.addrID(ctx, evt.Contract, evt.BlockNum, evt.TxId)
	if err != nil {
		return store.WrapError(op, err)
	}
	newTreasurerAid, err := r.addrID(ctx, evt.NewTreasurer, evt.BlockNum, evt.TxId)
	if err != nil {
		return store.WrapError(op, err)
	}
	return r.insertAdminValue(ctx, "cg_adm_treasurer_addr", "new_treasurer_aid",
		evt.EvtId, evt.BlockNum, evt.TxId, evt.TimeStamp, contractAid, newTreasurerAid)
}

// InsertCosmicTokenAddressChange records a CosmicSignatureTokenAddressChanged event.
func (r *Repo) InsertCosmicTokenAddressChange(ctx context.Context, evt *cgmodel.CGCosmicTokenAddressChanged) error {
	const op = "insert into cg_adm_costok_addr"
	contractAid, err := r.addrID(ctx, evt.Contract, evt.BlockNum, evt.TxId)
	if err != nil {
		return store.WrapError(op, err)
	}
	newCosTokAid, err := r.addrID(ctx, evt.NewCosmicToken, evt.BlockNum, evt.TxId)
	if err != nil {
		return store.WrapError(op, err)
	}
	return r.insertAdminValue(ctx, "cg_adm_costok_addr", "new_costok_aid",
		evt.EvtId, evt.BlockNum, evt.TxId, evt.TimeStamp, contractAid, newCosTokAid)
}

// InsertCosmicSignatureAddressChange records a CosmicSignatureNftAddressChanged event.
func (r *Repo) InsertCosmicSignatureAddressChange(ctx context.Context, evt *cgmodel.CGCosmicSignatureAddressChanged) error {
	const op = "insert into cg_adm_cossig_addr"
	contractAid, err := r.addrID(ctx, evt.Contract, evt.BlockNum, evt.TxId)
	if err != nil {
		return store.WrapError(op, err)
	}
	newCosSigAid, err := r.addrID(ctx, evt.NewCosmicSignature, evt.BlockNum, evt.TxId)
	if err != nil {
		return store.WrapError(op, err)
	}
	return r.insertAdminValue(ctx, "cg_adm_cossig_addr", "new_cossig_aid",
		evt.EvtId, evt.BlockNum, evt.TxId, evt.TimeStamp, contractAid, newCosSigAid)
}

// Admin: proxy and lifecycle.

// InsertUpgraded records an ERC-1967 Upgraded event.
func (r *Repo) InsertUpgraded(ctx context.Context, evt *cgmodel.CGUpgraded) error {
	const op = "insert into cg_adm_upgraded"
	contractAid, err := r.addrID(ctx, evt.Contract, evt.BlockNum, evt.TxId)
	if err != nil {
		return store.WrapError(op, err)
	}
	implementationAid, err := r.addrID(ctx, evt.Implementation, evt.BlockNum, evt.TxId)
	if err != nil {
		return store.WrapError(op, err)
	}
	return r.insertAdminValue(ctx, "cg_adm_upgraded", "implementation_aid",
		evt.EvtId, evt.BlockNum, evt.TxId, evt.TimeStamp, contractAid, implementationAid)
}

// InsertAdminChanged records an ERC-1967 AdminChanged event.
func (r *Repo) InsertAdminChanged(ctx context.Context, evt *cgmodel.CGAdminChanged) error {
	const op = "insert into cg_adm_admin_changed"
	contractAid, err := r.addrID(ctx, evt.Contract, evt.BlockNum, evt.TxId)
	if err != nil {
		return store.WrapError(op, err)
	}
	oldAdminAid, err := r.addrID(ctx, evt.OldAdmin, evt.BlockNum, evt.TxId)
	if err != nil {
		return store.WrapError(op, err)
	}
	newAdminAid, err := r.addrID(ctx, evt.NewAdmin, evt.BlockNum, evt.TxId)
	if err != nil {
		return store.WrapError(op, err)
	}
	query := "INSERT INTO cg_adm_admin_changed(" +
		"evtlog_id,block_num,tx_id,time_stamp,contract_aid," +
		"old_admin_aid,new_admin_aid" +
		") VALUES($1,$2,$3,TO_TIMESTAMP($4),$5,$6,$7)"
	_, err = r.q(ctx).Exec(ctx, query,
		evt.EvtId,
		evt.BlockNum,
		evt.TxId,
		evt.TimeStamp,
		contractAid,
		oldAdminAid,
		newAdminAid,
	)
	return store.WrapError(op, err)
}

// InsertOwnershipTransfer records an OwnershipTransferred event
// (contract_code identifies which platform contract changed owner).
func (r *Repo) InsertOwnershipTransfer(ctx context.Context, evt *cgmodel.CGOwnershipTransferred) error {
	const op = "insert into cg_adm_ownership"
	contractAid, err := r.addrID(ctx, evt.Contract, evt.BlockNum, evt.TxId)
	if err != nil {
		return store.WrapError(op, err)
	}
	newOwnerAid, err := r.addrID(ctx, evt.NewOwner, evt.BlockNum, evt.TxId)
	if err != nil {
		return store.WrapError(op, err)
	}
	prevOwnerAid, err := r.addrID(ctx, evt.PrevOwner, evt.BlockNum, evt.TxId)
	if err != nil {
		return store.WrapError(op, err)
	}
	query := "INSERT INTO cg_adm_ownership(" +
		"evtlog_id,block_num,tx_id,time_stamp,contract_aid," +
		"prev_owner_aid,new_owner_aid,contract_code" +
		") VALUES($1,$2,$3,TO_TIMESTAMP($4),$5,$6,$7,$8)"
	_, err = r.q(ctx).Exec(ctx, query,
		evt.EvtId,
		evt.BlockNum,
		evt.TxId,
		evt.TimeStamp,
		contractAid,
		prevOwnerAid,
		newOwnerAid,
		evt.ContractCode,
	)
	return store.WrapError(op, err)
}

// InsertInitialized records an Initializable Initialized event.
func (r *Repo) InsertInitialized(ctx context.Context, evt *cgmodel.CGInitialized) error {
	contractAid, err := r.addrID(ctx, evt.Contract, evt.BlockNum, evt.TxId)
	if err != nil {
		return store.WrapError("insert into cg_adm_initialized", err)
	}
	return r.insertAdminValue(ctx, "cg_adm_initialized", "version",
		evt.EvtId, evt.BlockNum, evt.TxId, evt.TimeStamp, contractAid, evt.Version)
}

// Admin: timing and pricing parameters.

// InsertTimeIncreaseChange records a legacy TimeIncreaseChanged event.
func (r *Repo) InsertTimeIncreaseChange(ctx context.Context, evt *cgmodel.CGTimeIncreaseChanged) error {
	contractAid, err := r.addrID(ctx, evt.Contract, evt.BlockNum, evt.TxId)
	if err != nil {
		return store.WrapError("insert into cg_adm_time_inc", err)
	}
	return r.insertAdminValue(ctx, "cg_adm_time_inc", "new_time_inc",
		evt.EvtId, evt.BlockNum, evt.TxId, evt.TimeStamp, contractAid, evt.NewTimeIncrease)
}

// InsertTimeoutClaimPrizeChange records a TimeoutDurationToClaimMainPrizeChanged event.
func (r *Repo) InsertTimeoutClaimPrizeChange(ctx context.Context, evt *cgmodel.CGTimeoutClaimPrizeChanged) error {
	contractAid, err := r.addrID(ctx, evt.Contract, evt.BlockNum, evt.TxId)
	if err != nil {
		return store.WrapError("insert into cg_adm_timeout_claimprize", err)
	}
	return r.insertAdminValue(ctx, "cg_adm_timeout_claimprize", "new_timeout",
		evt.EvtId, evt.BlockNum, evt.TxId, evt.TimeStamp, contractAid, evt.NewTimeout)
}

// InsertTimeoutToWithdrawPrizesChange records a TimeoutDurationToWithdrawPrizesChanged event.
func (r *Repo) InsertTimeoutToWithdrawPrizesChange(ctx context.Context, evt *cgmodel.CGTimeoutToWithdrawPrizeChanged) error {
	contractAid, err := r.addrID(ctx, evt.Contract, evt.BlockNum, evt.TxId)
	if err != nil {
		return store.WrapError("insert into cg_adm_timeout_withdraw", err)
	}
	return r.insertAdminValue(ctx, "cg_adm_timeout_withdraw", "new_timeout",
		evt.EvtId, evt.BlockNum, evt.TxId, evt.TimeStamp, contractAid, evt.NewTimeout)
}

// InsertPriceIncreaseChange records an EthBidPriceIncreaseDivisorChanged event.
func (r *Repo) InsertPriceIncreaseChange(ctx context.Context, evt *cgmodel.CGPriceIncreaseChanged) error {
	contractAid, err := r.addrID(ctx, evt.Contract, evt.BlockNum, evt.TxId)
	if err != nil {
		return store.WrapError("insert into cg_adm_price_inc", err)
	}
	return r.insertAdminValue(ctx, "cg_adm_price_inc", "new_price_increase",
		evt.EvtId, evt.BlockNum, evt.TxId, evt.TimeStamp, contractAid, evt.NewPriceIncrease)
}

// InsertMainPrizeMicrosecondsChange records a MainPrizeTimeIncrementInMicroSecondsChanged event.
func (r *Repo) InsertMainPrizeMicrosecondsChange(ctx context.Context, evt *cgmodel.CGMainPrizeMicroSecondsIncreaseChanged) error {
	contractAid, err := r.addrID(ctx, evt.Contract, evt.BlockNum, evt.TxId)
	if err != nil {
		return store.WrapError("insert into cg_adm_prize_microsec", err)
	}
	return r.insertAdminValue(ctx, "cg_adm_prize_microsec", "new_microseconds",
		evt.EvtId, evt.BlockNum, evt.TxId, evt.TimeStamp, contractAid, evt.NewMicroseconds)
}

// InsertInitialSecondsUntilPrizeChange records an InitialDurationUntilMainPrizeDivisorChanged event.
func (r *Repo) InsertInitialSecondsUntilPrizeChange(ctx context.Context, evt *cgmodel.CGInitialSecondsUntilPrizeChanged) error {
	contractAid, err := r.addrID(ctx, evt.Contract, evt.BlockNum, evt.TxId)
	if err != nil {
		return store.WrapError("insert into cg_adm_inisecprize", err)
	}
	return r.insertAdminValue(ctx, "cg_adm_inisecprize", "new_inisec",
		evt.EvtId, evt.BlockNum, evt.TxId, evt.TimeStamp, contractAid, evt.NewInitialSecondsUntilPrize)
}

// InsertActivationTimeChange records a RoundActivationTimeChanged event.
func (r *Repo) InsertActivationTimeChange(ctx context.Context, evt *cgmodel.CGActivationTimeChanged) error {
	contractAid, err := r.addrID(ctx, evt.Contract, evt.BlockNum, evt.TxId)
	if err != nil {
		return store.WrapError("insert into cg_adm_acttime", err)
	}
	return r.insertAdminValue(ctx, "cg_adm_acttime", "new_atime",
		evt.EvtId, evt.BlockNum, evt.TxId, evt.TimeStamp, contractAid, evt.NewActivationTime)
}

// InsertCstAuctionLengthChange records a CST dutch-auction duration change
// (CstDutchAuctionDurationDivisorChanged and the V2 duration event share the
// table).
func (r *Repo) InsertCstAuctionLengthChange(ctx context.Context, evt *cgmodel.CGCstDutchAuctionDurationDivisorChanged) error {
	contractAid, err := r.addrID(ctx, evt.Contract, evt.BlockNum, evt.TxId)
	if err != nil {
		return store.WrapError("insert into cg_adm_cst_auclen", err)
	}
	return r.insertAdminValue(ctx, "cg_adm_cst_auclen", "new_len",
		evt.EvtId, evt.BlockNum, evt.TxId, evt.TimeStamp, contractAid, evt.NewValue)
}

// InsertCstAuctionDurationChangeDivisorChange records a CstDutchAuctionDurationChangeDivisorChanged event.
func (r *Repo) InsertCstAuctionDurationChangeDivisorChange(ctx context.Context, evt *cgmodel.CGCstDutchAuctionDurationChangeDivisorChanged) error {
	contractAid, err := r.addrID(ctx, evt.Contract, evt.BlockNum, evt.TxId)
	if err != nil {
		return store.WrapError("insert into cg_adm_cst_auclen_chg_div", err)
	}
	return r.insertAdminValue(ctx, "cg_adm_cst_auclen_chg_div", "new_len",
		evt.EvtId, evt.BlockNum, evt.TxId, evt.TimeStamp, contractAid, evt.NewValue)
}

// InsertEthAuctionDurationDivisorChange records an EthDutchAuctionDurationDivisorChanged event.
func (r *Repo) InsertEthAuctionDurationDivisorChange(ctx context.Context, evt *cgmodel.CGEthDutchAuctionDurationDivisorChanged) error {
	contractAid, err := r.addrID(ctx, evt.Contract, evt.BlockNum, evt.TxId)
	if err != nil {
		return store.WrapError("insert into cg_adm_eth_auclen", err)
	}
	return r.insertAdminValue(ctx, "cg_adm_eth_auclen", "new_len",
		evt.EvtId, evt.BlockNum, evt.TxId, evt.TimeStamp, contractAid, evt.NewValue)
}

// InsertEthAuctionEndingBidPriceDivisorChange records an EthDutchAuctionEndingBidPriceDivisorChanged event.
func (r *Repo) InsertEthAuctionEndingBidPriceDivisorChange(ctx context.Context, evt *cgmodel.CGEthDutchAuctionEndingBidPriceDivisorChanged) error {
	contractAid, err := r.addrID(ctx, evt.Contract, evt.BlockNum, evt.TxId)
	if err != nil {
		return store.WrapError("insert into cg_adm_eth_auc_endprice", err)
	}
	return r.insertAdminValue(ctx, "cg_adm_eth_auc_endprice", "new_len",
		evt.EvtId, evt.BlockNum, evt.TxId, evt.TimeStamp, contractAid, evt.NewValue)
}

// Admin: rewards and limits.

// InsertStaticCstRewardChange records a CstPrizeAmountChanged event.
func (r *Repo) InsertStaticCstRewardChange(ctx context.Context, evt *cgmodel.CGStaticCstReward) error {
	contractAid, err := r.addrID(ctx, evt.Contract, evt.BlockNum, evt.TxId)
	if err != nil {
		return store.WrapError("insert into cg_adm_erc_rwd_mul", err)
	}
	return r.insertAdminValue(ctx, "cg_adm_erc_rwd_mul", "new_reward",
		evt.EvtId, evt.BlockNum, evt.TxId, evt.TimeStamp, contractAid, evt.NewReward)
}

// InsertMarketingRewardChange records a MarketingWalletCstContributionAmountChanged event.
func (r *Repo) InsertMarketingRewardChange(ctx context.Context, evt *cgmodel.CGMarketingRewardChanged) error {
	contractAid, err := r.addrID(ctx, evt.Contract, evt.BlockNum, evt.TxId)
	if err != nil {
		return store.WrapError("insert into cg_adm_mkt_reward", err)
	}
	return r.insertAdminValue(ctx, "cg_adm_mkt_reward", "new_reward",
		evt.EvtId, evt.BlockNum, evt.TxId, evt.TimeStamp, contractAid, evt.NewReward)
}

// InsertCstRewardForBiddingChange records a CST-bid-reward change
// (CstRewardAmountForBiddingChanged and its V2 variants share the table; the
// insert trigger updates cg_glob_stats.cst_reward_for_bidding).
func (r *Repo) InsertCstRewardForBiddingChange(ctx context.Context, evt *cgmodel.CGCstRewardForBiddingChanged) error {
	contractAid, err := r.addrID(ctx, evt.Contract, evt.BlockNum, evt.TxId)
	if err != nil {
		return store.WrapError("insert into cg_adm_erc20_reward", err)
	}
	return r.insertAdminValue(ctx, "cg_adm_erc20_reward", "new_reward",
		evt.EvtId, evt.BlockNum, evt.TxId, evt.TimeStamp, contractAid, evt.NewReward)
}

// InsertMaxMessageLengthChange records a BidMessageLengthMaxLimitChanged event.
func (r *Repo) InsertMaxMessageLengthChange(ctx context.Context, evt *cgmodel.CGMaxMessageLengthChanged) error {
	contractAid, err := r.addrID(ctx, evt.Contract, evt.BlockNum, evt.TxId)
	if err != nil {
		return store.WrapError("insert into cg_adm_msg_len", err)
	}
	return r.insertAdminValue(ctx, "cg_adm_msg_len", "new_length",
		evt.EvtId, evt.BlockNum, evt.TxId, evt.TimeStamp, contractAid, evt.NewMessageLength)
}

// InsertCstMinLimit records a CstDutchAuctionBeginningBidPriceMinLimitChanged event.
func (r *Repo) InsertCstMinLimit(ctx context.Context, evt *cgmodel.CGCstMinLimit) error {
	contractAid, err := r.addrID(ctx, evt.Contract, evt.BlockNum, evt.TxId)
	if err != nil {
		return store.WrapError("insert into cg_adm_cst_min_limit", err)
	}
	return r.insertAdminValue(ctx, "cg_adm_cst_min_limit", "min_limit",
		evt.EvtId, evt.BlockNum, evt.TxId, evt.TimeStamp, contractAid, evt.CstMinLimit)
}

// InsertNextRoundDelayDurationChange records a DelayDurationBeforeRoundActivationChanged event.
func (r *Repo) InsertNextRoundDelayDurationChange(ctx context.Context, evt *cgmodel.CGNextRoundDelayDuration) error {
	contractAid, err := r.addrID(ctx, evt.Contract, evt.BlockNum, evt.TxId)
	if err != nil {
		return store.WrapError("insert into cg_delay_duration", err)
	}
	return r.insertAdminValue(ctx, "cg_delay_duration", "new_value",
		evt.EvtId, evt.BlockNum, evt.TxId, evt.TimeStamp, contractAid, evt.NewValue)
}

// Admin: NFT metadata.

// InsertTokenGenerationScriptURL records an NftGenerationScriptUriChanged event.
func (r *Repo) InsertTokenGenerationScriptURL(ctx context.Context, evt *cgmodel.CGTokenGenerationScriptURL) error {
	contractAid, err := r.addrID(ctx, evt.Contract, evt.BlockNum, evt.TxId)
	if err != nil {
		return store.WrapError("insert into cg_adm_script_url", err)
	}
	return r.insertAdminValue(ctx, "cg_adm_script_url", "new_url",
		evt.EvtId, evt.BlockNum, evt.TxId, evt.TimeStamp, contractAid, evt.NewURL)
}

// InsertBaseURI records an NftBaseUriChanged event.
func (r *Repo) InsertBaseURI(ctx context.Context, evt *cgmodel.CGBaseURIEvent) error {
	contractAid, err := r.addrID(ctx, evt.Contract, evt.BlockNum, evt.TxId)
	if err != nil {
		return store.WrapError("insert into cg_adm_base_uri_cs", err)
	}
	return r.insertAdminValue(ctx, "cg_adm_base_uri_cs", "new_uri",
		evt.EvtId, evt.BlockNum, evt.TxId, evt.TimeStamp, contractAid, evt.NewURI)
}
