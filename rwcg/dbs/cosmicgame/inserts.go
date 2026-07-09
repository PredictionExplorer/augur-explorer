package cosmicgame

import (
	"os"
	"fmt"

	p "github.com/PredictionExplorer/augur-explorer/rwcg/primitives/cosmicgame"
)
func (sw *SQLStorageWrapper) Insert_prize_claim_event(evt *p.CGPrizeClaimEvent) {

	contract_aid := sw.S.Lookup_or_create_address(evt.ContractAddr,0, 0)
	winner_aid := sw.S.Lookup_or_create_address(evt.WinnerAddr,0, 0)
	var query string
	num_cs_nfts := evt.NumCSNfts
	if num_cs_nfts <= 0 { num_cs_nfts = 1 } // defensive: V2 events (and any unset) award exactly one NFT
	query =  "INSERT INTO "+sw.S.SchemaName()+".cg_prize_claim("+
				"evtlog_id,block_num,time_stamp,tx_id,contract_aid,"+
				"round_num,token_id,num_cs_nfts,winner_aid,timeout,amount,cst_amount"+
				") VALUES($1,$2,TO_TIMESTAMP($3),$4,$5,$6,$7,$8,$9,$10,$11,$12)"
	_,err := sw.S.Db().Exec(query,
		evt.EvtId,
		evt.BlockNum,
		evt.TimeStamp,
		evt.TxId,
		contract_aid,
		evt.RoundNum,
		evt.TokenId,
		num_cs_nfts,
		winner_aid,
		evt.Timeout,
		evt.Amount,
		evt.CstAmount,
	)
	if err != nil {
		sw.S.Log_msg(fmt.Sprintf("DB error: can't insert into cg_prize_claim table: %v\n",err))
		os.Exit(1)
	}
}
func (sw *SQLStorageWrapper) Insert_bid_event(evt *p.CGBidEvent) {

	contract_aid := sw.S.Lookup_or_create_address(evt.ContractAddr,0, 0)
	bidder_aid := sw.S.Lookup_or_create_address(evt.LastBidderAddr,0, 0)
	var query string
	// Calculate bid position for this round
	var bid_position int64 = 1
	row := sw.S.Db().QueryRow("SELECT COALESCE(MAX(bid_position), 0) + 1 FROM "+sw.S.SchemaName()+".cg_bid WHERE round_num = $1", evt.RoundNum)
	err := row.Scan(&bid_position)
	if err != nil || bid_position == 0 {
		bid_position = 1 // First bid in round
	}
	
	// Get current cst_reward_for_bidding from settings (populated by admin events or ETL chain sync at startup).
	var cst_reward string
	row = sw.S.Db().QueryRow("SELECT cst_reward_for_bidding FROM "+sw.S.SchemaName()+".cg_glob_stats LIMIT 1")
	err = row.Scan(&cst_reward)
	if evt.BidCstRewardAmount != "-1" {
		cst_reward = evt.BidCstRewardAmount
	} else if err != nil || cst_reward == "" || cst_reward == "0" {
		sw.S.Log_msg(fmt.Sprintf("DB error: cst_reward_for_bidding unset in cg_glob_stats (process admin events or restart ETL for chain sync): %v, value=%q\n", err, cst_reward))
		os.Exit(1)
	}
	
	// Determine eth_price and cst_price based on bid type
	eth_price := evt.EthPrice
	cst_price := evt.CstPrice
	if evt.BidType == 2 { // CST bid
		eth_price = "-1"
		// cst_price already set
	} else { // ETH or RandomWalk bid
		// eth_price already set
		cst_price = "-1"
	}
	
	query =  "INSERT INTO "+sw.S.SchemaName()+".cg_bid("+
			"evtlog_id,block_num,time_stamp,tx_id,contract_aid,"+
			"bidder_aid,rwalk_nft_id,eth_price,cst_price,cst_reward,bid_cst_reward_amount,cst_dutch_auction_duration,prize_time,msg,round_num,bid_type,bid_position"+
			") VALUES($1,$2,TO_TIMESTAMP($3),$4,$5,$6,$7,$8,$9,$10,$11,$12,TO_TIMESTAMP($13),$14,$15,$16,$17) RETURNING id"
	var bid_id int64
	err = sw.S.Db().QueryRow(query,
		evt.EvtId,
		evt.BlockNum,
		evt.TimeStamp,
		evt.TxId,
		contract_aid,
		bidder_aid,
		evt.RandomWalkTokenId,
		eth_price,
		cst_price,
		cst_reward,
		evt.BidCstRewardAmount,
		evt.CstDutchAuctionDuration,
		evt.PrizeTime,
		evt.Message,
		evt.RoundNum,
		evt.BidType,
		bid_position,
	).Scan(&bid_id)
	if err != nil {
		sw.S.Log_msg(fmt.Sprintf("DB error: can't insert into cg_bid table: %v\n",err))
		os.Exit(1)
	}
	sw.insert_bid_reward_rows(evt, bid_id, bidder_aid)
}
// insert_bid_reward_rows records the V3 bid CST reward 90/10 split (Comment-202607161) as two rows in
// cg_bid_reward: reward_type 0 for the bidder placing the bid (~10% in V3, the whole reward in V2/V1) and
// reward_type 1 for the outbid (previous last) bidder (90% in V3, absent otherwise). The split is derived
// by the ETL from the CST mint Transfer events, so it is correct for both eras without an upgrade-block flag.
func (sw *SQLStorageWrapper) insert_bid_reward_rows(evt *p.CGBidEvent, bid_id, bidder_aid int64) {
	this_reward := evt.ThisBidderReward
	if this_reward == "" || this_reward == "-1" {
		// Legacy v1 bids do not carry a split; fall back to the whole cst_reward as the bidder's share.
		this_reward = "0"
	}
	query := "INSERT INTO "+sw.S.SchemaName()+".cg_bid_reward(evtlog_id,bid_id,round_num,recipient_aid,reward_type,amount) "+
		"VALUES($1,$2,$3,$4,0,$5)"
	if _,err := sw.S.Db().Exec(query,evt.EvtId,bid_id,evt.RoundNum,bidder_aid,this_reward); err != nil {
		sw.S.Log_msg(fmt.Sprintf("DB error: can't insert this-bidder row into cg_bid_reward: %v\n",err))
		os.Exit(1)
	}
	if evt.PrevBidderAddr != "" && evt.PrevBidderReward != "" && evt.PrevBidderReward != "0" {
		prev_aid := sw.S.Lookup_or_create_address(evt.PrevBidderAddr,0,0)
		query = "INSERT INTO "+sw.S.SchemaName()+".cg_bid_reward(evtlog_id,bid_id,round_num,recipient_aid,reward_type,amount) "+
			"VALUES($1,$2,$3,$4,1,$5)"
		if _,err := sw.S.Db().Exec(query,evt.EvtId,bid_id,evt.RoundNum,prev_aid,evt.PrevBidderReward); err != nil {
			sw.S.Log_msg(fmt.Sprintf("DB error: can't insert prev-bidder row into cg_bid_reward: %v\n",err))
			os.Exit(1)
		}
	}
}
func (sw *SQLStorageWrapper) Insert_donation_event(evt *p.CGDonationEvent ) {

	contract_aid := sw.S.Lookup_or_create_address(evt.ContractAddr,0, 0)
	donor_aid := sw.S.Lookup_or_create_address(evt.DonorAddr,0, 0)
	var query string
	query =  "INSERT INTO "+sw.S.SchemaName()+".cg_eth_donated("+
					"evtlog_id,block_num,time_stamp,tx_id,contract_aid,"+
					"donor_aid,round_num,amount"+
					") VALUES($1,$2,TO_TIMESTAMP($3),$4,$5,$6,$7,$8)"
	_,err := sw.S.Db().Exec(query,
		evt.EvtId,
		evt.BlockNum,
		evt.TimeStamp,
		evt.TxId,
		contract_aid,
		donor_aid,
		evt.RoundNum,
		evt.Amount,
	)
	if err != nil {
		sw.S.Log_msg(fmt.Sprintf("DB error: can't insert into cg_eth_donated table: %v\n",err))
		os.Exit(1)
	}
}
func (sw *SQLStorageWrapper) Insert_donation_with_info_event(evt *p.CGDonationWithInfoEvent ) {

	contract_aid := sw.S.Lookup_or_create_address(evt.ContractAddr,0, 0)
	donor_aid := sw.S.Lookup_or_create_address(evt.DonorAddr,0, 0)
	var query string
	query =  "INSERT INTO "+sw.S.SchemaName()+".cg_eth_donated_wi("+
					"evtlog_id,block_num,time_stamp,tx_id,contract_aid,"+
					"donor_aid,round_num,record_id,amount"+
					") VALUES($1,$2,TO_TIMESTAMP($3),$4,$5,$6,$7,$8,$9)"
	_,err := sw.S.Db().Exec(query,
		evt.EvtId,
		evt.BlockNum,
		evt.TimeStamp,
		evt.TxId,
		contract_aid,
		donor_aid,
		evt.RoundNum,
		evt.RecordId,
		evt.Amount,
	)
	if err != nil {
		sw.S.Log_msg(fmt.Sprintf("DB error: can't insert into cg_eth_donated_wi table: %v\n",err))
		os.Exit(1)
	}
}
func (sw *SQLStorageWrapper) Insert_donation_wi_data_json(recordId int64,data string) {

	var query string
	query =  "INSERT INTO "+sw.S.SchemaName()+".cg_donation_json("+
					"record_id,data"+
					") VALUES($1,$2)"
	_,err := sw.S.Db().Exec(query,recordId,data)
	if err != nil {
		sw.S.Log_msg(fmt.Sprintf("DB error: can't insert into cg_donation_json : %v\n",err))
		os.Exit(1)
	}
}
func (sw *SQLStorageWrapper) Insert_donation_received(evt *p.CGDonationReceivedEvent ) {

	contract_aid := sw.S.Lookup_or_create_address(evt.ContractAddr,0, 0)
	donor_aid := sw.S.Lookup_or_create_address(evt.DonorAddr,0, 0)
	var query string
	query =  "INSERT INTO "+sw.S.SchemaName()+".cg_donation_received("+
					"evtlog_id,block_num,time_stamp,tx_id,contract_aid,"+
					"donor_aid,amount,round_num"+
					") VALUES($1,$2,TO_TIMESTAMP($3),$4,$5,$6,$7,$8)"
	_,err := sw.S.Db().Exec(query,
		evt.EvtId,
		evt.BlockNum,
		evt.TimeStamp,
		evt.TxId,
		contract_aid,
		donor_aid,
		evt.Amount,
		evt.RoundNum,
	)
	if err != nil {
		sw.S.Log_msg(fmt.Sprintf("DB error: can't insert into cg_donation_received table: %v\n",err))
		os.Exit(1)
	}
}
func (sw *SQLStorageWrapper) Insert_donation_sent(evt *p.CGDonationSentEvent ) {

	contract_aid := sw.S.Lookup_or_create_address(evt.ContractAddr,0, 0)
	charity_aid := sw.S.Lookup_or_create_address(evt.CharityAddr,0, 0)
	var query string
	query =  "INSERT INTO "+sw.S.SchemaName()+".cg_donation_sent("+
					"evtlog_id,block_num,time_stamp,tx_id,contract_aid,"+
					"charity_aid,amount"+
					") VALUES($1,$2,TO_TIMESTAMP($3),$4,$5,$6,$7)"
	_,err := sw.S.Db().Exec(query,
		evt.EvtId,
		evt.BlockNum,
		evt.TimeStamp,
		evt.TxId,
		contract_aid,
		charity_aid,
		evt.Amount,
	)
	if err != nil {
		sw.S.Log_msg(fmt.Sprintf("DB error: can't insert into cg_donation_sent table: %v\n",err))
		os.Exit(1)
	}
}
func (sw *SQLStorageWrapper) Insert_erc20_donated_event(evt *p.CGERC20DonationEvent) {

	contract_aid := sw.S.Lookup_or_create_address(evt.ContractAddr,0, 0)
	donor_aid := sw.S.Lookup_or_create_address(evt.DonorAddr,0, 0)
	token_aid := sw.S.Lookup_or_create_address(evt.TokenAddr,0, 0)
	var query string
	query =  "INSERT INTO "+sw.S.SchemaName()+".cg_erc20_donation("+
					"evtlog_id,block_num,time_stamp,tx_id,contract_aid,"+
					"donor_aid,token_aid,round_num,bid_id,amount"+
				") VALUES($1,$2,TO_TIMESTAMP($3),$4,$5,$6,$7,$8,$9,$10)"
	_,err := sw.S.Db().Exec(query,
		evt.EvtId,
		evt.BlockNum,
		evt.TimeStamp,
		evt.TxId,
		contract_aid,
		donor_aid,
		token_aid,
		evt.RoundNum,
		evt.BidId,
		evt.Amount,
	)
	if err != nil {
		sw.S.Log_msg(fmt.Sprintf("DB error: can't insert into cg_erc20_donation table: %v\n",err))
		os.Exit(1)
	}
}
func (sw *SQLStorageWrapper) Insert_nft_donation_event(evt *p.CGNFTDonationEvent) {

	contract_aid := sw.S.Lookup_or_create_address(evt.ContractAddr,0, 0)
	donor_aid := sw.S.Lookup_or_create_address(evt.DonorAddr,0, 0)
	token_aid := sw.S.Lookup_or_create_address(evt.TokenAddr,0, 0)
	var query string
	query =  "INSERT INTO "+sw.S.SchemaName()+".cg_nft_donation("+
					"evtlog_id,block_num,time_stamp,tx_id,contract_aid,"+
					"donor_aid,token_aid,token_id,round_num,idx,bid_id,token_uri"+
				") VALUES($1,$2,TO_TIMESTAMP($3),$4,$5,$6,$7,$8,$9,$10,$11,$12)"
	_,err := sw.S.Db().Exec(query,
		evt.EvtId,
		evt.BlockNum,
		evt.TimeStamp,
		evt.TxId,
		contract_aid,
		donor_aid,
		token_aid,
		evt.TokenId,
		evt.RoundNum,
		evt.Index,
		evt.BidId,
		evt.NFTTokenURI,
	)
	if err != nil {
		sw.S.Log_msg(fmt.Sprintf("DB error: can't insert into cg_nft_donation table: %v\n",err))
		os.Exit(1)
	}
}
func (sw *SQLStorageWrapper) Insert_charity_updated_event(evt *p.CGCharityUpdatedEvent) {

	contract_aid := sw.S.Lookup_or_create_address(evt.ContractAddr,0, 0)
	new_charity_aid := sw.S.Lookup_or_create_address(evt.NewCharityAddr,0, 0)
	var query string
	query =  "INSERT INTO "+sw.S.SchemaName()+".cg_charity_receiver_changed("+
					"evtlog_id,block_num,time_stamp,tx_id,contract_aid,"+
					"charity_aid"+
					") VALUES($1,$2,TO_TIMESTAMP($3),$4,$5,$6)"
	_,err := sw.S.Db().Exec(query,
		evt.EvtId,
		evt.BlockNum,
		evt.TimeStamp,
		evt.TxId,
		contract_aid,
		new_charity_aid,
	)
	if err != nil {
		sw.S.Log_msg(fmt.Sprintf("DB error: can't insert into cg_charity_receiver_changed table: %v\n",err))
		os.Exit(1)
	}
}
func (sw *SQLStorageWrapper) Insert_token_name_event(evt *p.CGTokenNameEvent) {

	contract_aid := sw.S.Lookup_or_create_address(evt.ContractAddr,0, 0)
	var query string
	query =  "INSERT INTO "+sw.S.SchemaName()+".cg_token_name("+
					"evtlog_id,block_num,time_stamp,tx_id,contract_aid,"+
					"token_id,token_name"+
					") VALUES($1,$2,TO_TIMESTAMP($3),$4,$5,$6,$7)"
	_,err := sw.S.Db().Exec(query,
		evt.EvtId,
		evt.BlockNum,
		evt.TimeStamp,
		evt.TxId,
		contract_aid,
		evt.TokenId,
		evt.TokenName,
	)
	if err != nil {
		sw.S.Log_msg(fmt.Sprintf("DB error: can't insert into cg_token_name table: %v\n",err))
		os.Exit(1)
	}
}
func (sw *SQLStorageWrapper) Insert_mint_event(evt *p.CGMintEvent) {

	contract_aid := sw.S.Lookup_or_create_address(evt.ContractAddr,0, 0)
	owner_aid := sw.S.Lookup_or_create_address(evt.OwnerAddr,0, 0)

	var query string
	query =  "INSERT INTO "+sw.S.SchemaName()+".cg_mint_event("+
					"evtlog_id,block_num,time_stamp,tx_id,contract_aid,"+
					"owner_aid,cur_owner_aid,token_id,round_num,seed"+
					") VALUES($1,$2,TO_TIMESTAMP($3),$4,$5,$6,$7,$8,$9,$10)"
	_,err := sw.S.Db().Exec(query,
		evt.EvtId,
		evt.BlockNum,
		evt.TimeStamp,
		evt.TxId,
		contract_aid,
		owner_aid,
		owner_aid,
		evt.TokenId,
		evt.RoundNum,
		evt.Seed,
	)
	if err != nil {
		sw.S.Log_msg(fmt.Sprintf("DB error: can't insert into cg_mint_event table: %v\n",err))
		os.Exit(1)
	}
}
func (sw *SQLStorageWrapper) Insert_prize_deposit(evt *p.CGPrizesEthDeposit) {

	contract_aid := sw.S.Lookup_or_create_address(evt.ContractAddr,0, 0)
	winner_aid := sw.S.Lookup_or_create_address(evt.WinnerAddr,0, 0)

	var query string
	query =  "INSERT INTO "+sw.S.SchemaName()+".cg_prize_deposit ("+
					"evtlog_id,block_num,time_stamp,tx_id,contract_aid,"+
					"winner_aid,round_num,winner_index,amount"+
					") VALUES($1,$2,TO_TIMESTAMP($3),$4,$5,$6,$7,$8,$9)"
	_,err := sw.S.Db().Exec(query,
		evt.EvtId,
		evt.BlockNum,
		evt.TimeStamp,
		evt.TxId,
		contract_aid,
		winner_aid,
		evt.Round,
		evt.WinnerIndex,
		evt.Amount,
	)
	if err != nil {
		sw.S.Log_msg(fmt.Sprintf("DB error: can't insert into cg_prize_deposit table: %v\n",err))
		os.Exit(1)
	}
}
func (sw *SQLStorageWrapper) Insert_prize_withdrawal(evt *p.CGPrizesEthWithdrawal) {

	contract_aid := sw.S.Lookup_or_create_address(evt.ContractAddr,0, 0)
	winner_aid := sw.S.Lookup_or_create_address(evt.WinnerAddr,0, 0)
	beneficiary_aid := sw.S.Lookup_or_create_address(evt.BeneficiaryAddr,0, 0)

	var query string
	query =  "INSERT INTO "+sw.S.SchemaName()+".cg_prize_withdrawal("+
					"evtlog_id,block_num,time_stamp,tx_id,contract_aid,"+
					"round_num,winner_aid,beneficiary_aid,amount"+
					") VALUES($1,$2,TO_TIMESTAMP($3),$4,$5,$6,$7,$8,$9)"
	_,err := sw.S.Db().Exec(query,
		evt.EvtId,
		evt.BlockNum,
		evt.TimeStamp,
		evt.TxId,
		contract_aid,
		evt.Round,
		winner_aid,
		beneficiary_aid,
		evt.Amount,
	)
	if err != nil {
		sw.S.Log_msg(fmt.Sprintf("DB error: can't insert into cg_prize_withdrawal table: %v\n",err))
		os.Exit(1)
	}
}
func (sw *SQLStorageWrapper) Insert_raffle_nft_winner(evt *p.CGRaffleNFTWinner) {

	contract_aid := sw.S.Lookup_or_create_address(evt.ContractAddr,0, 0)
	winner_aid := sw.S.Lookup_or_create_address(evt.WinnerAddr,0, 0)

	var query string
	query =  "INSERT INTO "+sw.S.SchemaName()+".cg_raffle_nft_prize ("+
					"evtlog_id,block_num,time_stamp,tx_id,contract_aid,"+
					"winner_aid,round_num,token_id,winner_idx,cst_amount,is_rwalk,is_staker"+
					") VALUES($1,$2,TO_TIMESTAMP($3),$4,$5,$6,$7,$8,$9,$10,$11,$12)"
	_,err := sw.S.Db().Exec(query,
		evt.EvtId,
		evt.BlockNum,
		evt.TimeStamp,
		evt.TxId,
		contract_aid,
		winner_aid,
		evt.Round,
		evt.TokenId,
		evt.WinnerIndex,
		evt.CstAmount,
		evt.IsRandomWalk,
		evt.IsStaker,
	)
	if err != nil {
		sw.S.Log_msg(fmt.Sprintf("DB error: can't insert into cg_raffle_nft_prize table: %v\n",err))
		os.Exit(1)
	}
}
func (sw *SQLStorageWrapper) Insert_raffle_eth_winner(evt *p.CGRaffleETHWinner) {

	contract_aid := sw.S.Lookup_or_create_address(evt.ContractAddr,0, 0)
	winner_aid := sw.S.Lookup_or_create_address(evt.WinnerAddr,0, 0)

	var query string
	query =  "INSERT INTO "+sw.S.SchemaName()+".cg_raffle_eth_prize ("+
					"evtlog_id,block_num,time_stamp,tx_id,contract_aid,"+
					"winner_aid,round_num,winner_idx,amount"+
					") VALUES($1,$2,TO_TIMESTAMP($3),$4,$5,$6,$7,$8,$9)"
	_,err := sw.S.Db().Exec(query,
		evt.EvtId,
		evt.BlockNum,
		evt.TimeStamp,
		evt.TxId,
		contract_aid,
		winner_aid,
		evt.Round,
		evt.WinnerIndex,
		evt.Amount,
	)
	if err != nil {
		sw.S.Log_msg(fmt.Sprintf("DB error: can't insert into cg_raffle_eth_prize table: %v\n",err))
		os.Exit(1)
	}
}
func (sw *SQLStorageWrapper) Insert_endurance_winner(evt *p.CGEnduranceWinner) {
	// Note: The Solidity EnduranceChampionPrizePaid event does not emit a winner_index.
	// There is exactly one endurance champion per round, so winner_index is implicitly 0.

	contract_aid := sw.S.Lookup_or_create_address(evt.ContractAddr,0, 0)
	winner_aid := sw.S.Lookup_or_create_address(evt.WinnerAddr,0, 0)

	var query string
	query =  "INSERT INTO "+sw.S.SchemaName()+".cg_endurance_prize ("+
					"evtlog_id,block_num,time_stamp,tx_id,contract_aid,"+
					"winner_aid,round_num,erc721_token_id,erc20_amount"+
					") VALUES($1,$2,TO_TIMESTAMP($3),$4,$5,$6,$7,$8,$9)"
	_,err := sw.S.Db().Exec(query,
		evt.EvtId,
		evt.BlockNum,
		evt.TimeStamp,
		evt.TxId,
		contract_aid,
		winner_aid,
		evt.Round,
		evt.Erc721TokenId,
		evt.Erc20Amount,
	)
	if err != nil {
		sw.S.Log_msg(fmt.Sprintf("DB error: can't insert into cg_endurance_prize table: %v\n",err))
		os.Exit(1)
	}
}
func (sw *SQLStorageWrapper) Insert_lastcst_bidder_winner(evt *p.CGLastBidderWinner) {
	// Note: The Solidity LastCstBidderPrizePaid event does not emit a winner_index.
	// There is exactly one last CST bidder per round, so winner_index is implicitly 0.

	contract_aid := sw.S.Lookup_or_create_address(evt.ContractAddr,0, 0)
	winner_aid := sw.S.Lookup_or_create_address(evt.WinnerAddr,0, 0)

	var query string
	query =  "INSERT INTO "+sw.S.SchemaName()+".cg_lastcst_prize ("+
					"evtlog_id,block_num,time_stamp,tx_id,contract_aid,"+
					"winner_aid,round_num,erc721_token_id,erc20_amount"+
					") VALUES($1,$2,TO_TIMESTAMP($3),$4,$5,$6,$7,$8,$9)"
	_,err := sw.S.Db().Exec(query,
		evt.EvtId,
		evt.BlockNum,
		evt.TimeStamp,
		evt.TxId,
		contract_aid,
		winner_aid,
		evt.Round,
		evt.Erc721TokenId,
		evt.Erc20Amount,
	)
	if err != nil {
		sw.S.Log_msg(fmt.Sprintf("DB error: can't insert into cg_lastcst_prize table: %v\n",err))
		os.Exit(1)
	}
}
func (sw *SQLStorageWrapper) Insert_chrono_warrior_event(evt *p.CGChronoWarrior) {

	contract_aid := sw.S.Lookup_or_create_address(evt.ContractAddr,0, 0)
	winner_aid := sw.S.Lookup_or_create_address(evt.WinnerAddr,0, 0)

	var query string
	query =  "INSERT INTO "+sw.S.SchemaName()+".cg_chrono_warrior_prize("+
					"evtlog_id,block_num,time_stamp,tx_id,contract_aid,"+
					"winner_aid,round_num,winner_index,eth_amount,cst_amount,nft_id"+
					") VALUES($1,$2,TO_TIMESTAMP($3),$4,$5,$6,$7,$8,$9,$10,$11)"
	_,err := sw.S.Db().Exec(query,
		evt.EvtId,
		evt.BlockNum,
		evt.TimeStamp,
		evt.TxId,
		contract_aid,
		winner_aid,
		evt.Round,
		evt.WinnerIndex,
		evt.EthAmount,
		evt.CstAmount,
		evt.NftId,
	)
	if err != nil {
		sw.S.Log_msg(fmt.Sprintf("DB error: can't insert into cg_chrono_warrior_prize table: %v\n",err))
		os.Exit(1)
	}
}
func (sw *SQLStorageWrapper) Insert_donated_token_claimed(evt *p.CGDonatedTokenClaimed) {

	contract_aid := sw.S.Lookup_or_create_address(evt.ContractAddr,0, 0)
	token_aid := sw.S.Lookup_or_create_address(evt.TokenAddr,0, 0)
	winner_aid := sw.S.Lookup_or_create_address(evt.BeneficiaryAddr,0, 0)

	var query string
	query =  "INSERT INTO "+sw.S.SchemaName()+".cg_donated_tok_claimed ("+
					"evtlog_id,block_num,time_stamp,tx_id,contract_aid,"+
					"round_num,idx,token_aid,winner_aid,amount"+
					") VALUES($1,$2,TO_TIMESTAMP($3),$4,$5,$6,$7,$8,$9,$10)"
	_,err := sw.S.Db().Exec(query,
		evt.EvtId,
		evt.BlockNum,
		evt.TimeStamp,
		evt.TxId,
		contract_aid,
		evt.RoundNum,
		evt.Index,
		token_aid,
		winner_aid,
		evt.Amount,
	)
	if err != nil {
		sw.S.Log_msg(fmt.Sprintf("DB error: can't insert into cg_donated_tok_claimed table: %v\n",err))
		os.Exit(1)
	}
}
func (sw *SQLStorageWrapper) Insert_donated_nft_claimed(evt *p.CGDonatedNFTClaimed) {

	contract_aid := sw.S.Lookup_or_create_address(evt.ContractAddr,0, 0)
	token_aid := sw.S.Lookup_or_create_address(evt.TokenAddr,0, 0)
	winner_aid := sw.S.Lookup_or_create_address(evt.BeneficiaryAddr,0, 0)

	var query string
	query =  "INSERT INTO "+sw.S.SchemaName()+".cg_donated_nft_claimed ("+
					"evtlog_id,block_num,time_stamp,tx_id,contract_aid,"+
					"round_num,idx,token_aid,winner_aid,token_id"+
					") VALUES($1,$2,TO_TIMESTAMP($3),$4,$5,$6,$7,$8,$9,$10)"
	_,err := sw.S.Db().Exec(query,
		evt.EvtId,
		evt.BlockNum,
		evt.TimeStamp,
		evt.TxId,
		contract_aid,
		evt.RoundNum,
		evt.Index,
		token_aid,
		winner_aid,
		evt.TokenId,
	)
	if err != nil {
		sw.S.Log_msg(fmt.Sprintf("DB error: can't insert into cg_donated_nft_claimed table: %v\n",err))
		os.Exit(1)
	}
}
func (sw *SQLStorageWrapper) Insert_nft_unstaked_rwalk_event(evt *p.CGNftUnstakedRWalk) {

	contract_aid:=sw.S.Lookup_or_create_address(evt.ContractAddr,evt.BlockNum,evt.TxId)
	staker_aid:=sw.S.Lookup_or_create_address(evt.StakerAddress,evt.BlockNum,evt.TxId)
	var query string
	query = "INSERT INTO "+sw.S.SchemaName()+".cg_nft_unstaked_rwalk (" +
				"evtlog_id,block_num,tx_id,time_stamp,contract_aid, "+
				"action_id,token_id,num_staked_nfts,staker_aid" +
			") VALUES (" +
				"$1,$2,$3,TO_TIMESTAMP($4),$5,$6,$7,$8,$9"+
			")"
	_,err := sw.S.Db().Exec(query,
		evt.EvtId,
		evt.BlockNum,
		evt.TxId,
		evt.TimeStamp,
		contract_aid,
		evt.ActionId,
		evt.NftId,
		evt.NumStakedNfts,
		staker_aid,
	)
	if err != nil {
		sw.S.Log_msg(fmt.Sprintf("DB error: can't insert into cg_nft_unstaked_rwalk table: %v\n",err))
		os.Exit(1)
	}
}
func (sw *SQLStorageWrapper) Insert_nft_unstaked_cst_event(evt *p.CGNftUnstakedCst) {

	contract_aid:=sw.S.Lookup_or_create_address(evt.ContractAddr,evt.BlockNum,evt.TxId)
	staker_aid:=sw.S.Lookup_or_create_address(evt.StakerAddress,evt.BlockNum,evt.TxId)
	var query string
	query = "INSERT INTO "+sw.S.SchemaName()+".cg_nft_unstaked_cst (" +
				"evtlog_id,block_num,tx_id,time_stamp,contract_aid, "+
				"action_id,token_id,num_staked_nfts,staker_aid,reward,reward_per_tok,action_counter" +
			") VALUES (" +
				"$1,$2,$3,TO_TIMESTAMP($4),$5,$6,$7,$8,$9,$10,$11,$12"+
			")"
	_,err := sw.S.Db().Exec(query,
		evt.EvtId,
		evt.BlockNum,
		evt.TxId,
		evt.TimeStamp,
		contract_aid,
		evt.ActionId,
		evt.NftId,
		evt.NumStakedNfts,
		staker_aid,
		evt.RewardAmount,
		evt.RewardPerToken,
		evt.ActionCounter,
	)
	if err != nil {
		sw.S.Log_msg(fmt.Sprintf("DB error: can't insert into cg_nft_unstaked_cst  table: %v\n",err))
		os.Exit(1)
	}
}
func (sw *SQLStorageWrapper) Insert_eth_deposit_event(evt *p.CGEthDeposit) {

	contract_aid:=sw.S.Lookup_or_create_address(evt.ContractAddr,evt.BlockNum,evt.TxId)
	var query string
	query = "INSERT INTO "+sw.S.SchemaName()+".cg_staking_eth_deposit(" +
				"evtlog_id,block_num,tx_id,time_stamp,contract_aid, "+
				"deposit_time,round_num,deposit_id,num_staked_nfts,deposit_amount,amount_per_token,modulo,accum_modulo" +
			") VALUES (" +
				"$1,$2,$3,TO_TIMESTAMP($4),$5,TO_TIMESTAMP($6),$7,$8,$9,$10,$11,$12,$13"+
			")"
	_,err := sw.S.Db().Exec(query,
		evt.EvtId,
		evt.BlockNum,
		evt.TxId,
		evt.TimeStamp,
		contract_aid,
		evt.DepositTime,
		evt.RoundNum,
		evt.DepositId,
		evt.NumStakedNfts,
		evt.Amount,
		evt.AmountPerStaker,
		evt.Modulo,
		evt.AccumModulo,
	)
	if err != nil {
		sw.S.Log_msg(fmt.Sprintf("DB error: can't insert into cg_staking_eth_deposit table: %v\n",err))
		os.Exit(1)
	}
}
func (sw *SQLStorageWrapper) Insert_marketing_reward_paid_event(evt *p.CGMarketingRewardPaid) {

	contract_aid:=sw.S.Lookup_or_create_address(evt.ContractAddr,evt.BlockNum,evt.TxId)
	marketer_aid:=sw.S.Lookup_or_create_address(evt.Marketer,evt.BlockNum,evt.TxId)
	var query string
	query = "INSERT INTO "+sw.S.SchemaName()+".cg_mkt_reward(" +
				"evtlog_id,block_num,tx_id,time_stamp,contract_aid, "+
				"amount,marketer_aid" +
			") VALUES (" +
				"$1,$2,$3,TO_TIMESTAMP($4),$5,$6,$7"+
			")"
	_,err := sw.S.Db().Exec(query,
		evt.EvtId,
		evt.BlockNum,
		evt.TxId,
		evt.TimeStamp,
		contract_aid,
		evt.Amount,
		marketer_aid,
	)
	if err != nil {
		sw.S.Log_msg(fmt.Sprintf("DB error: can't insert into cg_mkt_reward table: %v\n",err))
		os.Exit(1)
	}
}
func (sw *SQLStorageWrapper) Insert_cosmic_signature_transfer_event(evt *p.CGERC721Transfer) {

	contract_aid:=sw.S.Lookup_or_create_address(evt.Contract,evt.BlockNum,evt.TxId)
	from_aid:=sw.S.Lookup_or_create_address(evt.From,evt.BlockNum,evt.TxId)
	to_aid:=sw.S.Lookup_or_create_address(evt.To,evt.BlockNum,evt.TxId)
	otype := int(0)
	if evt.From == "0x0000000000000000000000000000000000000000" {
		otype = 1
	}
	if evt.To == "0x0000000000000000000000000000000000000000" {
		otype = 2
	}
	var query string
	query = "INSERT INTO "+sw.S.SchemaName()+".cg_erc721_transfer(" +
				"evtlog_id,block_num,tx_id,time_stamp,contract_aid, "+
				"token_id,from_aid,to_aid,otype" +
			") VALUES (" +
				"$1,$2,$3,TO_TIMESTAMP($4),$5,$6,$7,$8,$9"+
			")"
	_,err := sw.S.Db().Exec(query,
		evt.EvtId,
		evt.BlockNum,
		evt.TxId,
		evt.TimeStamp,
		contract_aid,
		evt.TokenId,
		from_aid,
		to_aid,
		otype,
	)
	if err != nil {
		sw.S.Log_msg(fmt.Sprintf("DB error: can't insert into cg_erc721_transfer table: %v\n",err))
		os.Exit(1)
	}
}
func (sw *SQLStorageWrapper) Insert_cosmic_token_transfer_event(evt *p.CGERC20Transfer) {

	contract_aid:=sw.S.Lookup_or_create_address(evt.Contract,evt.BlockNum,evt.TxId)
	from_aid:=sw.S.Lookup_or_create_address(evt.From,evt.BlockNum,evt.TxId)
	to_aid:=sw.S.Lookup_or_create_address(evt.To,evt.BlockNum,evt.TxId)
	otype := int(0)
	if evt.From == "0x0000000000000000000000000000000000000000" {
		otype = 1
	}
	if evt.To == "0x0000000000000000000000000000000000000000" {
		otype = 2
	}
	var query string
	query = "INSERT INTO "+sw.S.SchemaName()+".cg_erc20_transfer(" +
				"evtlog_id,block_num,tx_id,time_stamp,contract_aid, "+
				"value,from_aid,to_aid,otype" +
			") VALUES (" +
				"$1,$2,$3,TO_TIMESTAMP($4),$5,$6,$7,$8,$9"+
			")"
	_,err := sw.S.Db().Exec(query,
		evt.EvtId,
		evt.BlockNum,
		evt.TxId,
		evt.TimeStamp,
		contract_aid,
		evt.Value,
		from_aid,
		to_aid,
		otype,
	)
	if err != nil {
		sw.S.Log_msg(fmt.Sprintf("DB error: can't insert into cg_erc20_transfer table: %v\n",err))
		os.Exit(1)
	}
}
func (sw *SQLStorageWrapper) Insert_cosmic_game_charity_percentage_changed_event(evt *p.CGCharityPercentageChanged) {

	contract_aid:=sw.S.Lookup_or_create_address(evt.Contract,evt.BlockNum,evt.TxId)
	var query string
	query = "INSERT INTO "+sw.S.SchemaName()+".cg_adm_charity_pcent (" +
				"evtlog_id,block_num,tx_id,time_stamp,contract_aid, "+
				"percentage" +
			") VALUES (" +
				"$1,$2,$3,TO_TIMESTAMP($4),$5,$6"+
			")"
	_,err := sw.S.Db().Exec(query,
		evt.EvtId,
		evt.BlockNum,
		evt.TxId,
		evt.TimeStamp,
		contract_aid,
		evt.NewCharityPercentage,
	)
	if err != nil {
		sw.S.Log_msg(fmt.Sprintf("DB error: can't insert into cg_adm_charity_pcent table: %v\n",err))
		os.Exit(1)
	}
}
func (sw *SQLStorageWrapper) Insert_cosmic_game_prize_percentage_changed_event(evt *p.CGPrizePercentageChanged) {

	contract_aid:=sw.S.Lookup_or_create_address(evt.Contract,evt.BlockNum,evt.TxId)
	var query string
	query = "INSERT INTO "+sw.S.SchemaName()+".cg_adm_main_prize_pcent (" +
				"evtlog_id,block_num,tx_id,time_stamp,contract_aid, "+
				"percentage" +
			") VALUES (" +
				"$1,$2,$3,TO_TIMESTAMP($4),$5,$6"+
			")"
	_,err := sw.S.Db().Exec(query,
		evt.EvtId,
		evt.BlockNum,
		evt.TxId,
		evt.TimeStamp,
		contract_aid,
		evt.NewPrizePercentage,
	)
	if err != nil {
		sw.S.Log_msg(fmt.Sprintf("DB error: can't insert into cg_adm_main_prize_pcent table: %v\n",err))
		os.Exit(1)
	}
}
func (sw *SQLStorageWrapper) Insert_cosmic_game_raffle_percentage_changed_event(evt *p.CGRafflePercentageChanged) {

	contract_aid:=sw.S.Lookup_or_create_address(evt.Contract,evt.BlockNum,evt.TxId)
	var query string
	query = "INSERT INTO "+sw.S.SchemaName()+".cg_adm_raffle_pcent (" +
				"evtlog_id,block_num,tx_id,time_stamp,contract_aid, "+
				"percentage" +
			") VALUES (" +
				"$1,$2,$3,TO_TIMESTAMP($4),$5,$6"+
			")"
	_,err := sw.S.Db().Exec(query,
		evt.EvtId,
		evt.BlockNum,
		evt.TxId,
		evt.TimeStamp,
		contract_aid,
		evt.NewRafflePercentage,
	)
	if err != nil {
		sw.S.Log_msg(fmt.Sprintf("DB error: can't insert into cg_adm_raffle_pcent table: %v\n",err))
		os.Exit(1)
	}
}
func (sw *SQLStorageWrapper) Insert_cosmic_game_num_raffle_eth_winners_bidding_changed_event(evt *p.CGNumRaffleETHWinnersBiddingChanged) {

	contract_aid:=sw.S.Lookup_or_create_address(evt.Contract,evt.BlockNum,evt.TxId)
	var query string
	query = "INSERT INTO "+sw.S.SchemaName()+".cg_adm_raf_eth_bidding(" +
				"evtlog_id,block_num,tx_id,time_stamp,contract_aid, "+
				"num_winners" +
			") VALUES (" +
				"$1,$2,$3,TO_TIMESTAMP($4),$5,$6"+
			")"
	_,err := sw.S.Db().Exec(query,
		evt.EvtId,
		evt.BlockNum,
		evt.TxId,
		evt.TimeStamp,
		contract_aid,
		evt.NewNumRaffleETHWinnersBidding,
	)
	if err != nil {
		sw.S.Log_msg(fmt.Sprintf("DB error: can't insert into cg_adm_raf_eth_bidding table: %v\n",err))
		os.Exit(1)
	}
}
func (sw *SQLStorageWrapper) Insert_cosmic_game_num_raffle_nft_winners_bidding_changed_event(evt *p.CGNumRaffleNFTWinnersBiddingChanged) {

	contract_aid:=sw.S.Lookup_or_create_address(evt.Contract,evt.BlockNum,evt.TxId)
	var query string
	query = "INSERT INTO "+sw.S.SchemaName()+".cg_adm_raf_nft_bidding(" +
				"evtlog_id,block_num,tx_id,time_stamp,contract_aid, "+
				"num_winners" +
			") VALUES (" +
				"$1,$2,$3,TO_TIMESTAMP($4),$5,$6"+
			")"
	_,err := sw.S.Db().Exec(query,
		evt.EvtId,
		evt.BlockNum,
		evt.TxId,
		evt.TimeStamp,
		contract_aid,
		evt.NewNumRaffleNFTWinnersBidding,
	)
	if err != nil {
		sw.S.Log_msg(fmt.Sprintf("DB error: can't insert into cg_adm_raf_nft_winners_bidding table: %v\n",err))
		os.Exit(1)
	}
}
func (sw *SQLStorageWrapper) Insert_cosmic_game_num_raffle_nft_winners_staking_rwalk_changed_event(evt *p.CGNumRaffleNFTWinnersStakingRWalkChanged) {

	contract_aid:=sw.S.Lookup_or_create_address(evt.Contract,evt.BlockNum,evt.TxId)
	var query string
	query = "INSERT INTO "+sw.S.SchemaName()+".cg_adm_raf_nft_staking_rwalk(" +
				"evtlog_id,block_num,tx_id,time_stamp,contract_aid, "+
				"num_winners" +
			") VALUES (" +
				"$1,$2,$3,TO_TIMESTAMP($4),$5,$6"+
			")"
	_,err := sw.S.Db().Exec(query,
		evt.EvtId,
		evt.BlockNum,
		evt.TxId,
		evt.TimeStamp,
		contract_aid,
		evt.NewNumRaffleNFTWinnersStakingRWalk,
	)
	if err != nil {
		sw.S.Log_msg(fmt.Sprintf("DB error: can't insert into cg_adm_raf_nft_staking_rwalk table: %v\n",err))
		os.Exit(1)
	}
}
func (sw *SQLStorageWrapper) Insert_cosmic_game_staking_percentage_changed_event(evt *p.CGStakingPercentageChanged) {

	contract_aid:=sw.S.Lookup_or_create_address(evt.Contract,evt.BlockNum,evt.TxId)
	var query string
	query = "INSERT INTO "+sw.S.SchemaName()+".cg_adm_stake_pcent (" +
				"evtlog_id,block_num,tx_id,time_stamp,contract_aid, "+
				"percentage" +
			") VALUES (" +
				"$1,$2,$3,TO_TIMESTAMP($4),$5,$6"+
			")"
	_,err := sw.S.Db().Exec(query,
		evt.EvtId,
		evt.BlockNum,
		evt.TxId,
		evt.TimeStamp,
		contract_aid,
		evt.NewStakingPercentage,
	)
	if err != nil {
		sw.S.Log_msg(fmt.Sprintf("DB error: can't insert into cg_adm_stake_pcent table: %v\n",err))
		os.Exit(1)
	}
}
func (sw *SQLStorageWrapper) Insert_cosmic_game_chrono_percentage_changed_event(evt *p.CGChronoPercentageChanged) {

	contract_aid:=sw.S.Lookup_or_create_address(evt.Contract,evt.BlockNum,evt.TxId)
	var query string
	query = "INSERT INTO "+sw.S.SchemaName()+".cg_adm_chrono_pcent (" +
				"evtlog_id,block_num,tx_id,time_stamp,contract_aid, "+
				"percentage" +
			") VALUES (" +
				"$1,$2,$3,TO_TIMESTAMP($4),$5,$6"+
			")"
	_,err := sw.S.Db().Exec(query,
		evt.EvtId,
		evt.BlockNum,
		evt.TxId,
		evt.TimeStamp,
		contract_aid,
		evt.NewChronoPercentage,
	)
	if err != nil {
		sw.S.Log_msg(fmt.Sprintf("DB error: can't insert into cg_adm_chrono_pcent table: %v\n",err))
		os.Exit(1)
	}
}
func (sw *SQLStorageWrapper) Insert_cosmic_game_charity_address_changed_event(evt *p.CGCharityAddressChanged) {

	contract_aid:=sw.S.Lookup_or_create_address(evt.Contract,evt.BlockNum,evt.TxId)
	new_charity_aid:=sw.S.Lookup_or_create_address(evt.NewCharity,evt.BlockNum,evt.TxId)
	var query string
	query = "INSERT INTO "+sw.S.SchemaName()+".cg_adm_charity_wallet(" +
				"evtlog_id,block_num,tx_id,time_stamp,contract_aid, "+
				"new_charity_aid" +
			") VALUES (" +
				"$1,$2,$3,TO_TIMESTAMP($4),$5,$6"+
			")"
	_,err := sw.S.Db().Exec(query,
		evt.EvtId,
		evt.BlockNum,
		evt.TxId,
		evt.TimeStamp,
		contract_aid,
		new_charity_aid,
	)
	if err != nil {
		sw.S.Log_msg(fmt.Sprintf("DB error: can't insert into cg_adm_charity_wallet table: %v\n",err))
		os.Exit(1)
	}
}
func (sw *SQLStorageWrapper) Insert_cosmic_game_random_walk_address_changed_event(evt *p.CGRandomWalkAddressChanged) {

	contract_aid:=sw.S.Lookup_or_create_address(evt.Contract,evt.BlockNum,evt.TxId)
	new_rwalk_aid:=sw.S.Lookup_or_create_address(evt.NewRandomWalk,evt.BlockNum,evt.TxId)
	var query string
	query = "INSERT INTO "+sw.S.SchemaName()+".cg_adm_rwalk_addr(" +
				"evtlog_id,block_num,tx_id,time_stamp,contract_aid, "+
				"new_rwalk_aid" +
			") VALUES (" +
				"$1,$2,$3,TO_TIMESTAMP($4),$5,$6"+
			")"
	_,err := sw.S.Db().Exec(query,
		evt.EvtId,
		evt.BlockNum,
		evt.TxId,
		evt.TimeStamp,
		contract_aid,
		new_rwalk_aid,
	)
	if err != nil {
		sw.S.Log_msg(fmt.Sprintf("DB error: can't insert into cg_adm_rwalk_addr table: %v\n",err))
		os.Exit(1)
	}
}
func (sw *SQLStorageWrapper) Insert_cosmic_game_prize_wallet_address_changed_event(evt *p.CGPrizeWalletAddressChanged) {

	contract_aid:=sw.S.Lookup_or_create_address(evt.Contract,evt.BlockNum,evt.TxId)
	new_raffle_aid:=sw.S.Lookup_or_create_address(evt.NewPrizeWallet,evt.BlockNum,evt.TxId)
	var query string
	query = "INSERT INTO "+sw.S.SchemaName()+".cg_adm_prizes_wallet_addr(" +
				"evtlog_id,block_num,tx_id,time_stamp,contract_aid, "+
				"new_wallet_aid" +
			") VALUES (" +
				"$1,$2,$3,TO_TIMESTAMP($4),$5,$6"+
			")"
	_,err := sw.S.Db().Exec(query,
		evt.EvtId,
		evt.BlockNum,
		evt.TxId,
		evt.TimeStamp,
		contract_aid,
		new_raffle_aid,
	)
	if err != nil {
		sw.S.Log_msg(fmt.Sprintf("DB error: can't insert into cg_adm_prizes_wallet table: %v\n",err))
		os.Exit(1)
	}
}
func (sw *SQLStorageWrapper) Insert_cosmic_game_staking_wallet_cst_address_changed_event(evt *p.CGStakingWalletCSTAddressChanged) {

	contract_aid:=sw.S.Lookup_or_create_address(evt.Contract,evt.BlockNum,evt.TxId)
	new_staking_aid:=sw.S.Lookup_or_create_address(evt.NewStakingWalletCST,evt.BlockNum,evt.TxId)
	var query string
	query = "INSERT INTO "+sw.S.SchemaName()+".cg_adm_staking_cst_addr(" +
				"evtlog_id,block_num,tx_id,time_stamp,contract_aid, "+
				"new_staking_aid" +
			") VALUES (" +
				"$1,$2,$3,TO_TIMESTAMP($4),$5,$6"+
			")"
	_,err := sw.S.Db().Exec(query,
		evt.EvtId,
		evt.BlockNum,
		evt.TxId,
		evt.TimeStamp,
		contract_aid,
		new_staking_aid,
	)
	if err != nil {
		sw.S.Log_msg(fmt.Sprintf("DB error: can't insert into cg_adm_staking_cst_addr table: %v\n",err))
		os.Exit(1)
	}
}
func (sw *SQLStorageWrapper) Insert_cosmic_game_staking_wallet_rwalk_address_changed_event(evt *p.CGStakingWalletRWalkAddressChanged) {

	contract_aid:=sw.S.Lookup_or_create_address(evt.Contract,evt.BlockNum,evt.TxId)
	new_staking_aid:=sw.S.Lookup_or_create_address(evt.NewStakingWalletRWalk,evt.BlockNum,evt.TxId)
	var query string
	query = "INSERT INTO "+sw.S.SchemaName()+".cg_adm_staking_rwalk_addr(" +
				"evtlog_id,block_num,tx_id,time_stamp,contract_aid, "+
				"new_staking_aid" +
			") VALUES (" +
				"$1,$2,$3,TO_TIMESTAMP($4),$5,$6"+
			")"
	_,err := sw.S.Db().Exec(query,
		evt.EvtId,
		evt.BlockNum,
		evt.TxId,
		evt.TimeStamp,
		contract_aid,
		new_staking_aid,
	)
	if err != nil {
		sw.S.Log_msg(fmt.Sprintf("DB error: can't insert into cg_adm_staking_rwalk_addr table: %v\n",err))
		os.Exit(1)
	}
}
func (sw *SQLStorageWrapper) Insert_cosmic_game_marketing_wallet_address_changed_event(evt *p.CGMarketingWalletAddressChanged) {

	contract_aid:=sw.S.Lookup_or_create_address(evt.Contract,evt.BlockNum,evt.TxId)
	new_marketing_aid:=sw.S.Lookup_or_create_address(evt.NewMarketingWallet,evt.BlockNum,evt.TxId)
	var query string
	query = "INSERT INTO "+sw.S.SchemaName()+".cg_adm_marketing_addr(" +
				"evtlog_id,block_num,tx_id,time_stamp,contract_aid, "+
				"new_marketing_aid" +
			") VALUES (" +
				"$1,$2,$3,TO_TIMESTAMP($4),$5,$6"+
			")"
	_,err := sw.S.Db().Exec(query,
		evt.EvtId,
		evt.BlockNum,
		evt.TxId,
		evt.TimeStamp,
		contract_aid,
		new_marketing_aid,
	)
	if err != nil {
		sw.S.Log_msg(fmt.Sprintf("DB error: can't insert into cg_adm_marketing_addr table: %v\n",err))
		os.Exit(1)
	}
}
func (sw *SQLStorageWrapper) Insert_treasurer_address_changed_event(evt *p.CGTreasurerAddressChanged) {

	contract_aid:=sw.S.Lookup_or_create_address(evt.Contract,evt.BlockNum,evt.TxId)
	new_treasurer_aid:=sw.S.Lookup_or_create_address(evt.NewTreasurer,evt.BlockNum,evt.TxId)
	var query string
	query = "INSERT INTO "+sw.S.SchemaName()+".cg_adm_treasurer_addr(" +
				"evtlog_id,block_num,tx_id,time_stamp,contract_aid, "+
				"new_treasurer_aid" +
			") VALUES (" +
				"$1,$2,$3,TO_TIMESTAMP($4),$5,$6"+
			")"
	_,err := sw.S.Db().Exec(query,
		evt.EvtId,
		evt.BlockNum,
		evt.TxId,
		evt.TimeStamp,
		contract_aid,
		new_treasurer_aid,
	)
	if err != nil {
		sw.S.Log_msg(fmt.Sprintf("DB error: can't insert into cg_adm_treasurer_addr table: %v\n",err))
		os.Exit(1)
	}
}
func (sw *SQLStorageWrapper) Insert_cosmic_token_address_changed_event(evt *p.CGCosmicTokenAddressChanged) {

	contract_aid:=sw.S.Lookup_or_create_address(evt.Contract,evt.BlockNum,evt.TxId)
	new_costok_aid:=sw.S.Lookup_or_create_address(evt.NewCosmicToken,evt.BlockNum,evt.TxId)
	var query string
	query = "INSERT INTO "+sw.S.SchemaName()+".cg_adm_costok_addr(" +
				"evtlog_id,block_num,tx_id,time_stamp,contract_aid, "+
				"new_costok_aid" +
			") VALUES (" +
				"$1,$2,$3,TO_TIMESTAMP($4),$5,$6"+
			")"
	_,err := sw.S.Db().Exec(query,
		evt.EvtId,
		evt.BlockNum,
		evt.TxId,
		evt.TimeStamp,
		contract_aid,
		new_costok_aid,
	)
	if err != nil {
		sw.S.Log_msg(fmt.Sprintf("DB error: can't insert into cg_adm_costok_addr table: %v\n",err))
		os.Exit(1)
	}
}
func (sw *SQLStorageWrapper) Insert_cosmic_signature_address_changed_event(evt *p.CGCosmicSignatureAddressChanged) {

	contract_aid:=sw.S.Lookup_or_create_address(evt.Contract,evt.BlockNum,evt.TxId)
	new_cossig_aid:=sw.S.Lookup_or_create_address(evt.NewCosmicSignature,evt.BlockNum,evt.TxId)
	var query string
	query = "INSERT INTO "+sw.S.SchemaName()+".cg_adm_cossig_addr(" +
				"evtlog_id,block_num,tx_id,time_stamp,contract_aid, "+
				"new_cossig_aid" +
			") VALUES (" +
				"$1,$2,$3,TO_TIMESTAMP($4),$5,$6"+
			")"
	_,err := sw.S.Db().Exec(query,
		evt.EvtId,
		evt.BlockNum,
		evt.TxId,
		evt.TimeStamp,
		contract_aid,
		new_cossig_aid,
	)
	if err != nil {
		sw.S.Log_msg(fmt.Sprintf("DB error: can't insert into cg_adm_cossig_addr table: %v\n",err))
		os.Exit(1)
	}
}
func (sw *SQLStorageWrapper) Insert_upgraded_event(evt *p.CGUpgraded) {

	contract_aid:=sw.S.Lookup_or_create_address(evt.Contract,evt.BlockNum,evt.TxId)
	implementation_aid:=sw.S.Lookup_or_create_address(evt.Implementation,evt.BlockNum,evt.TxId)
	var query string
	query = "INSERT INTO "+sw.S.SchemaName()+".cg_adm_upgraded(" +
				"evtlog_id,block_num,tx_id,time_stamp,contract_aid, "+
				"implementation_aid" +
			") VALUES (" +
				"$1,$2,$3,TO_TIMESTAMP($4),$5,$6"+
			")"
	_,err := sw.S.Db().Exec(query,
		evt.EvtId,
		evt.BlockNum,
		evt.TxId,
		evt.TimeStamp,
		contract_aid,
		implementation_aid,
	)
	if err != nil {
		sw.S.Log_msg(fmt.Sprintf("DB error: can't insert into cg_adm_upgraded table: %v\n",err))
		os.Exit(1)
	}
}
func (sw *SQLStorageWrapper) Insert_admin_changed_event(evt *p.CGAdminChanged) {

	contract_aid:=sw.S.Lookup_or_create_address(evt.Contract,evt.BlockNum,evt.TxId)
	old_admin_aid:=sw.S.Lookup_or_create_address(evt.OldAdmin,evt.BlockNum,evt.TxId)
	new_admin_aid:=sw.S.Lookup_or_create_address(evt.NewAdmin,evt.BlockNum,evt.TxId)
	var query string
	query = "INSERT INTO "+sw.S.SchemaName()+".cg_adm_admin_changed(" +
				"evtlog_id,block_num,tx_id,time_stamp,contract_aid, "+
				"old_admin_aid,new_admin_aid" +
			") VALUES (" +
				"$1,$2,$3,TO_TIMESTAMP($4),$5,$6,$7"+
			")"
	_,err := sw.S.Db().Exec(query,
		evt.EvtId,
		evt.BlockNum,
		evt.TxId,
		evt.TimeStamp,
		contract_aid,
		old_admin_aid,
		new_admin_aid,
	)
	if err != nil {
		sw.S.Log_msg(fmt.Sprintf("DB error: can't insert into cg_adm_admin_changed table: %v\n",err))
		os.Exit(1)
	}
}
func (sw *SQLStorageWrapper) Insert_time_increase_changed_event(evt *p.CGTimeIncreaseChanged) {

	contract_aid:=sw.S.Lookup_or_create_address(evt.Contract,evt.BlockNum,evt.TxId)
	var query string
	query = "INSERT INTO "+sw.S.SchemaName()+".cg_adm_time_inc(" +
				"evtlog_id,block_num,tx_id,time_stamp,contract_aid, "+
				"new_time_inc" +
			") VALUES (" +
				"$1,$2,$3,TO_TIMESTAMP($4),$5,$6"+
			")"
	_,err := sw.S.Db().Exec(query,
		evt.EvtId,
		evt.BlockNum,
		evt.TxId,
		evt.TimeStamp,
		contract_aid,
		evt.NewTimeIncrease,
	)
	if err != nil {
		sw.S.Log_msg(fmt.Sprintf("DB error: can't insert into cg_adm_time_inc table: %v\n",err))
		os.Exit(1)
	}
}
func (sw *SQLStorageWrapper) Insert_timeout_claimprize_changed_event(evt *p.CGTimeoutClaimPrizeChanged) {

	contract_aid:=sw.S.Lookup_or_create_address(evt.Contract,evt.BlockNum,evt.TxId)
	var query string
	query = "INSERT INTO "+sw.S.SchemaName()+".cg_adm_timeout_claimprize(" +
				"evtlog_id,block_num,tx_id,time_stamp,contract_aid, "+
				"new_timeout" +
			") VALUES (" +
				"$1,$2,$3,TO_TIMESTAMP($4),$5,$6"+
			")"
	_,err := sw.S.Db().Exec(query,
		evt.EvtId,
		evt.BlockNum,
		evt.TxId,
		evt.TimeStamp,
		contract_aid,
		evt.NewTimeout,
	)
	if err != nil {
		sw.S.Log_msg(fmt.Sprintf("DB error: can't insert into cg_adm_timeout_claimprize table: %v\n",err))
		os.Exit(1)
	}
}
func (sw *SQLStorageWrapper) Insert_timeout_to_withdraw_prizes_changed_event(evt *p.CGTimeoutToWithdrawPrizeChanged) {

	contract_aid:=sw.S.Lookup_or_create_address(evt.Contract,evt.BlockNum,evt.TxId)
	var query string
	query = "INSERT INTO "+sw.S.SchemaName()+".cg_adm_timeout_withdraw(" +
				"evtlog_id,block_num,tx_id,time_stamp,contract_aid, "+
				"new_timeout" +
			") VALUES (" +
				"$1,$2,$3,TO_TIMESTAMP($4),$5,$6"+
			")"
	_,err := sw.S.Db().Exec(query,
		evt.EvtId,
		evt.BlockNum,
		evt.TxId,
		evt.TimeStamp,
		contract_aid,
		evt.NewTimeout,
	)
	if err != nil {
		sw.S.Log_msg(fmt.Sprintf("DB error: can't insert into cg_adm_timeout_withdraw table: %v\n",err))
		os.Exit(1)
	}
}
func (sw *SQLStorageWrapper) Insert_price_increase_changed_event(evt *p.CGPriceIncreaseChanged) {

	contract_aid:=sw.S.Lookup_or_create_address(evt.Contract,evt.BlockNum,evt.TxId)
	var query string
	query = "INSERT INTO "+sw.S.SchemaName()+".cg_adm_price_inc(" +
				"evtlog_id,block_num,tx_id,time_stamp,contract_aid, "+
				"new_price_increase" +
			") VALUES (" +
				"$1,$2,$3,TO_TIMESTAMP($4),$5,$6"+
			")"
	_,err := sw.S.Db().Exec(query,
		evt.EvtId,
		evt.BlockNum,
		evt.TxId,
		evt.TimeStamp,
		contract_aid,
		evt.NewPriceIncrease,
	)
	if err != nil {
		sw.S.Log_msg(fmt.Sprintf("DB error: can't insert into cg_adm_price_increase table: %v\n",err))
		os.Exit(1)
	}
}
func (sw *SQLStorageWrapper) Insert_mainprize_microseconds_increase_changed_event(evt *p.CGMainPrizeMicroSecondsIncreaseChanged) {

	contract_aid:=sw.S.Lookup_or_create_address(evt.Contract,evt.BlockNum,evt.TxId)
	var query string
	query = "INSERT INTO "+sw.S.SchemaName()+".cg_adm_prize_microsec (" +
				"evtlog_id,block_num,tx_id,time_stamp,contract_aid, "+
				"new_microseconds " +
			") VALUES (" +
				"$1,$2,$3,TO_TIMESTAMP($4),$5,$6"+
			")"
	_,err := sw.S.Db().Exec(query,
		evt.EvtId,
		evt.BlockNum,
		evt.TxId,
		evt.TimeStamp,
		contract_aid,
		evt.NewMicroseconds,
	)
	if err != nil {
		sw.S.Log_msg(fmt.Sprintf("DB error: can't insert into cg_adm_prize_microsec table: %v\n",err))
		os.Exit(1)
	}
}
func (sw *SQLStorageWrapper) Insert_initial_seconds_until_prize_changed_event(evt *p.CGInitialSecondsUntilPrizeChanged) {

	contract_aid:=sw.S.Lookup_or_create_address(evt.Contract,evt.BlockNum,evt.TxId)
	var query string
	query = "INSERT INTO "+sw.S.SchemaName()+".cg_adm_inisecprize (" +
				"evtlog_id,block_num,tx_id,time_stamp,contract_aid, "+
				"new_inisec" +
			") VALUES (" +
				"$1,$2,$3,TO_TIMESTAMP($4),$5,$6"+
			")"
	_,err := sw.S.Db().Exec(query,
		evt.EvtId,
		evt.BlockNum,
		evt.TxId,
		evt.TimeStamp,
		contract_aid,
		evt.NewInitialSecondsUntilPrize,
	)
	if err != nil {
		sw.S.Log_msg(fmt.Sprintf("DB error: can't insert into cg_adm_inisecprize table: %v\n",err))
		os.Exit(1)
	}
}
func (sw *SQLStorageWrapper) Insert_activation_time_changed_event(evt *p.CGActivationTimeChanged) {

	contract_aid:=sw.S.Lookup_or_create_address(evt.Contract,evt.BlockNum,evt.TxId)
	var query string
	query = "INSERT INTO "+sw.S.SchemaName()+".cg_adm_acttime (" +
				"evtlog_id,block_num,tx_id,time_stamp,contract_aid, "+
				"new_atime" +
			") VALUES (" +
				"$1,$2,$3,TO_TIMESTAMP($4),$5,$6"+
			")"
	_,err := sw.S.Db().Exec(query,
		evt.EvtId,
		evt.BlockNum,
		evt.TxId,
		evt.TimeStamp,
		contract_aid,
		evt.NewActivationTime,
	)
	if err != nil {
		sw.S.Log_msg(fmt.Sprintf("DB error: can't insert into cg_adm_acttime table: %v\n",err))
		os.Exit(1)
	}
}
func (sw *SQLStorageWrapper) Insert_round_start_cst_auction_length_changed_event(evt *p.CGCstDutchAuctionDurationDivisorChanged) {

	contract_aid:=sw.S.Lookup_or_create_address(evt.Contract,evt.BlockNum,evt.TxId)
	var query string
	query = "INSERT INTO "+sw.S.SchemaName()+".cg_adm_cst_auclen (" +
				"evtlog_id,block_num,tx_id,time_stamp,contract_aid, "+
				"new_len" +
			") VALUES (" +
				"$1,$2,$3,TO_TIMESTAMP($4),$5,$6"+
			")"
	_,err := sw.S.Db().Exec(query,
		evt.EvtId,
		evt.BlockNum,
		evt.TxId,
		evt.TimeStamp,
		contract_aid,
		evt.NewValue,
	)
	if err != nil {
		sw.S.Log_msg(fmt.Sprintf("DB error: can't insert into cg_adm_cst_auclen table: %v\n",err))
		os.Exit(1)
	}
}
func (sw *SQLStorageWrapper) Insert_cst_dutch_auction_duration_change_divisor_changed_event(evt *p.CGCstDutchAuctionDurationChangeDivisorChanged) {

	contract_aid:=sw.S.Lookup_or_create_address(evt.Contract,evt.BlockNum,evt.TxId)
	var query string
	query = "INSERT INTO "+sw.S.SchemaName()+".cg_adm_cst_auclen_chg_div (" +
				"evtlog_id,block_num,tx_id,time_stamp,contract_aid, "+
				"new_len" +
			") VALUES (" +
				"$1,$2,$3,TO_TIMESTAMP($4),$5,$6"+
			")"
	_,err := sw.S.Db().Exec(query,
		evt.EvtId,
		evt.BlockNum,
		evt.TxId,
		evt.TimeStamp,
		contract_aid,
		evt.NewValue,
	)
	if err != nil {
		sw.S.Log_msg(fmt.Sprintf("DB error: can't insert into cg_adm_cst_auclen_chg_div table: %v\n",err))
		os.Exit(1)
	}
}
func (sw *SQLStorageWrapper) Insert_eth_auction_duration_divisor_changed_event(evt *p.CGEthDutchAuctionDurationDivisorChanged) {

	contract_aid:=sw.S.Lookup_or_create_address(evt.Contract,evt.BlockNum,evt.TxId)
	var query string
	query = "INSERT INTO "+sw.S.SchemaName()+".cg_adm_eth_auclen (" +
				"evtlog_id,block_num,tx_id,time_stamp,contract_aid, "+
				"new_len" +
			") VALUES (" +
				"$1,$2,$3,TO_TIMESTAMP($4),$5,$6"+
			")"
	_,err := sw.S.Db().Exec(query,
		evt.EvtId,
		evt.BlockNum,
		evt.TxId,
		evt.TimeStamp,
		contract_aid,
		evt.NewValue,
	)
	if err != nil {
		sw.S.Log_msg(fmt.Sprintf("DB error: can't insert into cg_adm_eth_auclen table: %v\n",err))
		os.Exit(1)
	}
}
func (sw *SQLStorageWrapper) Insert_eth_dutch_auction_ending_bidprice_divisor_changed_event(evt *p.CGEthDutchAuctionEndingBidPriceDivisorChanged) {

	contract_aid:=sw.S.Lookup_or_create_address(evt.Contract,evt.BlockNum,evt.TxId)
	var query string
	query = "INSERT INTO "+sw.S.SchemaName()+".cg_adm_eth_auc_endprice(" +
				"evtlog_id,block_num,tx_id,time_stamp,contract_aid, "+
				"new_len" +
			") VALUES (" +
				"$1,$2,$3,TO_TIMESTAMP($4),$5,$6"+
			")"
	_,err := sw.S.Db().Exec(query,
		evt.EvtId,
		evt.BlockNum,
		evt.TxId,
		evt.TimeStamp,
		contract_aid,
		evt.NewValue,
	)
	if err != nil {
		sw.S.Log_msg(fmt.Sprintf("DB error: can't insert into cg_adm_eth_auc_endprice table: %v\n",err))
		os.Exit(1)
	}
}
func (sw *SQLStorageWrapper) Insert_static_cst_reward_changed_event(evt *p.CGStaticCstReward) {

	contract_aid:=sw.S.Lookup_or_create_address(evt.Contract,evt.BlockNum,evt.TxId)
	var query string
	query = "INSERT INTO "+sw.S.SchemaName()+".cg_adm_erc_rwd_mul(" +
				"evtlog_id,block_num,tx_id,time_stamp,contract_aid, "+
				"new_reward" +
			") VALUES (" +
				"$1,$2,$3,TO_TIMESTAMP($4),$5,$6"+
			")"
	_,err := sw.S.Db().Exec(query,
		evt.EvtId,
		evt.BlockNum,
		evt.TxId,
		evt.TimeStamp,
		contract_aid,
		evt.NewReward,
	)
	if err != nil {
		sw.S.Log_msg(fmt.Sprintf("DB error: can't insert into cg_adm_erc_rwd_mul table: %v\n",err))
		os.Exit(1)
	}
}
func (sw *SQLStorageWrapper) Insert_marketing_reward_changed_event(evt *p.CGMarketingRewardChanged) {

	contract_aid:=sw.S.Lookup_or_create_address(evt.Contract,evt.BlockNum,evt.TxId)
	var query string
	query = "INSERT INTO "+sw.S.SchemaName()+".cg_adm_mkt_reward(" +
				"evtlog_id,block_num,tx_id,time_stamp,contract_aid, "+
				"new_reward" +
			") VALUES (" +
				"$1,$2,$3,TO_TIMESTAMP($4),$5,$6"+
			")"
	_,err := sw.S.Db().Exec(query,
		evt.EvtId,
		evt.BlockNum,
		evt.TxId,
		evt.TimeStamp,
		contract_aid,
		evt.NewReward,
	)
	if err != nil {
		sw.S.Log_msg(fmt.Sprintf("DB error: can't insert into cg_adm_mkt_reward table: %v\n",err))
		os.Exit(1)
	}
}
func (sw *SQLStorageWrapper) Insert_erc20_token_reward_changed_event(evt *p.CGCstRewardForBiddingChanged) {

	contract_aid:=sw.S.Lookup_or_create_address(evt.Contract,evt.BlockNum,evt.TxId)
	var query string
	query = "INSERT INTO "+sw.S.SchemaName()+".cg_adm_erc20_reward(" +
				"evtlog_id,block_num,tx_id,time_stamp,contract_aid, "+
				"new_reward" +
			") VALUES (" +
				"$1,$2,$3,TO_TIMESTAMP($4),$5,$6"+
			")"
	_,err := sw.S.Db().Exec(query,
		evt.EvtId,
		evt.BlockNum,
		evt.TxId,
		evt.TimeStamp,
		contract_aid,
		evt.NewReward,
	)
	if err != nil {
		sw.S.Log_msg(fmt.Sprintf("DB error: can't insert into cg_adm_erc20_reward table: %v\n",err))
		os.Exit(1)
	}
}
func (sw *SQLStorageWrapper) Insert_max_message_length_changed_event(evt *p.CGMaxMessageLengthChanged) {

	contract_aid:=sw.S.Lookup_or_create_address(evt.Contract,evt.BlockNum,evt.TxId)
	var query string
	query = "INSERT INTO "+sw.S.SchemaName()+".cg_adm_msg_len(" +
				"evtlog_id,block_num,tx_id,time_stamp,contract_aid, "+
				"new_length" +
			") VALUES (" +
				"$1,$2,$3,TO_TIMESTAMP($4),$5,$6"+
			")"
	_,err := sw.S.Db().Exec(query,
		evt.EvtId,
		evt.BlockNum,
		evt.TxId,
		evt.TimeStamp,
		contract_aid,
		evt.NewMessageLength,
	)
	if err != nil {
		sw.S.Log_msg(fmt.Sprintf("DB error: can't insert into cg_adm_msg_len table: %v\n",err))
		os.Exit(1)
	}
}
func (sw *SQLStorageWrapper) Insert_token_generation_script_url_event(evt *p.CGTokenGenerationScriptURL) {

	contract_aid:=sw.S.Lookup_or_create_address(evt.Contract,evt.BlockNum,evt.TxId)
	var query string
	query = "INSERT INTO "+sw.S.SchemaName()+".cg_adm_script_url(" +
				"evtlog_id,block_num,tx_id,time_stamp,contract_aid, "+
				"new_url" +
			") VALUES (" +
				"$1,$2,$3,TO_TIMESTAMP($4),$5,$6"+
			")"
	_,err := sw.S.Db().Exec(query,
		evt.EvtId,
		evt.BlockNum,
		evt.TxId,
		evt.TimeStamp,
		contract_aid,
		evt.NewURL,
	)
	if err != nil {
		sw.S.Log_msg(fmt.Sprintf("DB error: can't insert into cg_adm_script_url table: %v\n",err))
		os.Exit(1)
	}
}
func (sw *SQLStorageWrapper) Insert_base_uri_event(evt *p.CGBaseURIEvent) {

	contract_aid:=sw.S.Lookup_or_create_address(evt.Contract,evt.BlockNum,evt.TxId)
	var query string
	query = "INSERT INTO "+sw.S.SchemaName()+".cg_adm_base_uri_cs(" +
				"evtlog_id,block_num,tx_id,time_stamp,contract_aid, "+
				"new_uri" +
			") VALUES (" +
				"$1,$2,$3,TO_TIMESTAMP($4),$5,$6"+
			")"
	_,err := sw.S.Db().Exec(query,
		evt.EvtId,
		evt.BlockNum,
		evt.TxId,
		evt.TimeStamp,
		contract_aid,
		evt.NewURI,
	)
	if err != nil {
		sw.S.Log_msg(fmt.Sprintf("DB error: can't insert into cg_adm_base_uri table: %v\n",err))
		os.Exit(1)
	}
}
func (sw *SQLStorageWrapper) Insert_nft_staked_cst_event(evt *p.CGNftStakedCst) {

	contract_aid:=sw.S.Lookup_or_create_address(evt.ContractAddr,evt.BlockNum,evt.TxId)
	staker_aid:=sw.S.Lookup_or_create_address(evt.StakerAddress,evt.BlockNum,evt.TxId)
	var query string
	query = "INSERT INTO "+sw.S.SchemaName()+".cg_nft_staked_cst(" +
				"evtlog_id,block_num,tx_id,time_stamp,contract_aid, "+
				"action_id,token_id,num_staked_nfts,reward_per_staker,staker_aid" +
			") VALUES (" +
				"$1,$2,$3,TO_TIMESTAMP($4),$5,$6,$7,$8,$9,$10"+
			")"
	_,err := sw.S.Db().Exec(query,
		evt.EvtId,
		evt.BlockNum,
		evt.TxId,
		evt.TimeStamp,
		contract_aid,
		evt.ActionId,
		evt.NftId,
		evt.NumStakedNfts,
		evt.RewardPerStaker,
		staker_aid,
	)
	if err != nil {
		sw.S.Log_msg(fmt.Sprintf("DB error: can't insert into cg_nft_staked_cst table: %v\n",err))
		os.Exit(1)
	}
}
func (sw *SQLStorageWrapper) Insert_nft_staked_rwalk_event(evt *p.CGNftStakedRWalk) {

	contract_aid:=sw.S.Lookup_or_create_address(evt.ContractAddr,evt.BlockNum,evt.TxId)
	staker_aid:=sw.S.Lookup_or_create_address(evt.StakerAddress,evt.BlockNum,evt.TxId)
	var table = "cg_nft_staked_cst"
	var query string
	query = "INSERT INTO "+sw.S.SchemaName()+".cg_nft_staked_rwalk("+
				"evtlog_id,block_num,tx_id,time_stamp,contract_aid, "+
				"action_id,token_id,num_staked_nfts,staker_aid" +
			") VALUES (" +
				"$1,$2,$3,TO_TIMESTAMP($4),$5,$6,$7,$8,$9"+
			")"
	_,err := sw.S.Db().Exec(query,
		evt.EvtId,
		evt.BlockNum,
		evt.TxId,
		evt.TimeStamp,
		contract_aid,
		evt.ActionId,
		evt.NftId,
		evt.NumStakedNfts,
		staker_aid,
	)
	if err != nil {
		sw.S.Log_msg(fmt.Sprintf("DB error: can't insert into "+table+" table: %v\n",err))
		os.Exit(1)
	}
}
func (sw *SQLStorageWrapper) Insert_ownership_transferred_event(evt *p.CGOwnershipTransferred) {

	contract_aid:=sw.S.Lookup_or_create_address(evt.Contract,evt.BlockNum,evt.TxId)
	new_owner_aid:=sw.S.Lookup_or_create_address(evt.NewOwner,evt.BlockNum,evt.TxId)
	prev_owner_aid:=sw.S.Lookup_or_create_address(evt.PrevOwner,evt.BlockNum,evt.TxId)
	var query string
	query = "INSERT INTO "+sw.S.SchemaName()+".cg_adm_ownership(" +
				"evtlog_id,block_num,tx_id,time_stamp,contract_aid, "+
				"prev_owner_aid,new_owner_aid,contract_code" +
			") VALUES (" +
				"$1,$2,$3,TO_TIMESTAMP($4),$5,$6,$7,$8"+
			")"
	_,err := sw.S.Db().Exec(query,
		evt.EvtId,
		evt.BlockNum,
		evt.TxId,
		evt.TimeStamp,
		contract_aid,
		prev_owner_aid,
		new_owner_aid,
		evt.ContractCode,
	)
	if err != nil {
		sw.S.Log_msg(fmt.Sprintf("DB error: can't insert into cg_adm_ownership table: %v\n",err))
		os.Exit(1)
	}
}
func (sw *SQLStorageWrapper) Insert_initialized_event(evt *p.CGInitialized) {

	contract_aid:=sw.S.Lookup_or_create_address(evt.Contract,evt.BlockNum,evt.TxId)
	var query string
	query = "INSERT INTO "+sw.S.SchemaName()+".cg_adm_initialized(" +
				"evtlog_id,block_num,tx_id,time_stamp,contract_aid, "+
				"version" +
			") VALUES (" +
				"$1,$2,$3,TO_TIMESTAMP($4),$5,$6"+
			")"
	_,err := sw.S.Db().Exec(query,
		evt.EvtId,
		evt.BlockNum,
		evt.TxId,
		evt.TimeStamp,
		contract_aid,
		evt.Version,
	)
	if err != nil {
		sw.S.Log_msg(fmt.Sprintf("DB error: can't insert into cg_adm_initialized table: %v\n",err))
		os.Exit(1)
	}
}
func (sw *SQLStorageWrapper) Insert_cst_min_limit_event(evt *p.CGCstMinLimit) {

	contract_aid:=sw.S.Lookup_or_create_address(evt.Contract,evt.BlockNum,evt.TxId)
	var query string
	query = "INSERT INTO "+sw.S.SchemaName()+".cg_adm_cst_min_limit(" +
				"evtlog_id,block_num,tx_id,time_stamp,contract_aid, "+
				"min_limit" +
			") VALUES (" +
				"$1,$2,$3,TO_TIMESTAMP($4),$5,$6"+
			")"
	_,err := sw.S.Db().Exec(query,
		evt.EvtId,
		evt.BlockNum,
		evt.TxId,
		evt.TimeStamp,
		contract_aid,
		evt.CstMinLimit,
	)
	if err != nil {
		sw.S.Log_msg(fmt.Sprintf("DB error: can't insert into cg_adm_cst_min_limit table: %v\n",err))
		os.Exit(1)
	}
}
func (sw *SQLStorageWrapper) Insert_fund_transfer_failed_event(evt *p.CGFundTransferFailed ) {

	contract_aid:=sw.S.Lookup_or_create_address(evt.Contract,evt.BlockNum,evt.TxId)
	destination_aid:=sw.S.Lookup_or_create_address(evt.Destination,evt.BlockNum,evt.TxId)
	var query string
	query = "INSERT INTO "+sw.S.SchemaName()+".cg_fund_transf_err(" +
				"evtlog_id,block_num,tx_id,time_stamp,contract_aid, "+
				"destination_aid,amount" +
			") VALUES (" +
				"$1,$2,$3,TO_TIMESTAMP($4),$5,$6,$7"+
			")"
	_,err := sw.S.Db().Exec(query,
		evt.EvtId,
		evt.BlockNum,
		evt.TxId,
		evt.TimeStamp,
		contract_aid,
		destination_aid,
		evt.Amount,
	)
	if err != nil {
		sw.S.Log_msg(fmt.Sprintf("DB error: can't insert into cg_fund_transf_err table: %v\n",err))
		os.Exit(1)
	}
}
func (sw *SQLStorageWrapper) Insert_erc20_transfer_failed_event(evt *p.CGErc20TransferFailed ) {

	contract_aid:=sw.S.Lookup_or_create_address(evt.Contract,evt.BlockNum,evt.TxId)
	destination_aid:=sw.S.Lookup_or_create_address(evt.Destination,evt.BlockNum,evt.TxId)
	var query string
	query = "INSERT INTO "+sw.S.SchemaName()+".cg_erc20_transf_err(" +
				"evtlog_id,block_num,tx_id,time_stamp,contract_aid, "+
				"destination_aid,amount" +
			") VALUES (" +
				"$1,$2,$3,TO_TIMESTAMP($4),$5,$6,$7"+
			")"
	_,err := sw.S.Db().Exec(query,
		evt.EvtId,
		evt.BlockNum,
		evt.TxId,
		evt.TimeStamp,
		contract_aid,
		destination_aid,
		evt.Amount,
	)
	if err != nil {
		sw.S.Log_msg(fmt.Sprintf("DB error: can't insert into cg_erc20_transf_err table: %v\n",err))
		os.Exit(1)
	}
}
func (sw *SQLStorageWrapper) Insert_funds_transferred_to_charity_event(evt *p.CGFundsToCharity ) {

	contract_aid:=sw.S.Lookup_or_create_address(evt.Contract,evt.BlockNum,evt.TxId)
	charity_aid:=sw.S.Lookup_or_create_address(evt.CharityAddr,evt.BlockNum,evt.TxId)
	var query string
	query = "INSERT INTO "+sw.S.SchemaName()+".cg_funds_to_charity(" +
				"evtlog_id,block_num,tx_id,time_stamp,contract_aid, "+
				"charity_aid,amount" +
			") VALUES (" +
				"$1,$2,$3,TO_TIMESTAMP($4),$5,$6,$7"+
			")"
	_,err := sw.S.Db().Exec(query,
		evt.EvtId,
		evt.BlockNum,
		evt.TxId,
		evt.TimeStamp,
		contract_aid,
		charity_aid,
		evt.Amount,
	)
	if err != nil {
		sw.S.Log_msg(fmt.Sprintf("DB error: can't insert into cg_funds_to_charity table: %v\n",err))
		os.Exit(1)
	}
}
func (sw *SQLStorageWrapper) Insert_delay_duration_before_next_round_changed_event(evt *p.CGNextRoundDelayDuration) {

	contract_aid:=sw.S.Lookup_or_create_address(evt.Contract,evt.BlockNum,evt.TxId)
	var query string
	query = "INSERT INTO "+sw.S.SchemaName()+".cg_delay_duration(" +
				"evtlog_id,block_num,tx_id,time_stamp,contract_aid, "+
				"new_value" +
			") VALUES (" +
				"$1,$2,$3,TO_TIMESTAMP($4),$5,$6"+
			")"
	_,err := sw.S.Db().Exec(query,
		evt.EvtId,
		evt.BlockNum,
		evt.TxId,
		evt.TimeStamp,
		contract_aid,
		evt.NewValue,
	)
	if err != nil {
		sw.S.Log_msg(fmt.Sprintf("DB error: can't insert into cg_delay_duration table: %v\n",err))
		os.Exit(1)
	}
}
func (sw *SQLStorageWrapper) Insert_round_started_event(evt *p.CGRoundStarted) {

	contract_aid:=sw.S.Lookup_or_create_address(evt.Contract,evt.BlockNum,evt.TxId)
	var query string
	query = "INSERT INTO "+sw.S.SchemaName()+".cg_first_bid(" +
				"evtlog_id,block_num,tx_id,time_stamp,contract_aid, "+
				"round_num,start_ts" +
			") VALUES (" +
				"$1,$2,$3,TO_TIMESTAMP($4),$5,$6,$7"+
			")"
	_,err := sw.S.Db().Exec(query,
		evt.EvtId,
		evt.BlockNum,
		evt.TxId,
		evt.TimeStamp,
		contract_aid,
		evt.RoundNum,
		evt.StartTimestamp,
	)
	if err != nil {
		sw.S.Log_msg(fmt.Sprintf("DB error: can't insert into cg_first_bid table: %v\n",err))
		os.Exit(1)
	}
}
// insert_v3_config_changed inserts a single-uint256 V3 ISystemEventsV3 config-changed event into `table`.
func (sw *SQLStorageWrapper) insert_v3_config_changed(table string, evtId, blockNum, txId, timeStamp int64, contract, newValue string) {
	contract_aid := sw.S.Lookup_or_create_address(contract, blockNum, txId)
	query := "INSERT INTO "+sw.S.SchemaName()+"."+table+"("+
				"evtlog_id,block_num,tx_id,time_stamp,contract_aid,new_value"+
			") VALUES($1,$2,$3,TO_TIMESTAMP($4),$5,$6)"
	_,err := sw.S.Db().Exec(query, evtId, blockNum, txId, timeStamp, contract_aid, newValue)
	if err != nil {
		sw.S.Log_msg(fmt.Sprintf("DB error: can't insert into %v table: %v\n",table,err))
		os.Exit(1)
	}
}
func (sw *SQLStorageWrapper) Insert_round_late_bid_duration_divisor_changed_event(evt *p.CGRoundLateBidDurationDivisorChanged) {
	sw.insert_v3_config_changed("cg_adm_late_bid_dur_divisor",evt.EvtId,evt.BlockNum,evt.TxId,evt.TimeStamp,evt.Contract,evt.NewValue)
}
func (sw *SQLStorageWrapper) Insert_round_late_bid_premium_base_multiplier_changed_event(evt *p.CGRoundLateBidPricePremiumAmountBaseMultiplierChanged) {
	sw.insert_v3_config_changed("cg_adm_late_bid_premium_base_mul",evt.EvtId,evt.BlockNum,evt.TxId,evt.TimeStamp,evt.Contract,evt.NewValue)
}
func (sw *SQLStorageWrapper) Insert_round_late_bid_premium_exponent_changed_event(evt *p.CGRoundLateBidPricePremiumAmountExponentChanged) {
	sw.insert_v3_config_changed("cg_adm_late_bid_premium_exponent",evt.EvtId,evt.BlockNum,evt.TxId,evt.TimeStamp,evt.Contract,evt.NewValue)
}
func (sw *SQLStorageWrapper) Insert_last_bidder_bid_cst_reward_amount_percentage_changed_event(evt *p.CGLastBidderBidCstRewardAmountPercentageChanged) {
	sw.insert_v3_config_changed("cg_adm_last_bidder_reward_pct",evt.EvtId,evt.BlockNum,evt.TxId,evt.TimeStamp,evt.Contract,evt.NewValue)
}
func (sw *SQLStorageWrapper) Insert_main_prize_num_cs_nfts_changed_event(evt *p.CGMainPrizeNumCosmicSignatureNftsChanged) {
	sw.insert_v3_config_changed("cg_adm_main_prize_num_nfts",evt.EvtId,evt.BlockNum,evt.TxId,evt.TimeStamp,evt.Contract,evt.NewValue)
}
