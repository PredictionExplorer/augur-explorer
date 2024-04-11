package cosmicgame

import (
	"os"
	"fmt"

	p "github.com/PredictionExplorer/augur-explorer/primitives/cosmicgame"
)
func (sw *SQLStorageWrapper) Insert_prize_claim_event(evt *p.CGPrizeClaimEvent) {

	contract_aid := sw.S.Lookup_or_create_address(evt.ContractAddr,0, 0)
	winner_aid := sw.S.Lookup_or_create_address(evt.WinnerAddr,0, 0)
	var query string
	query =  "INSERT INTO "+sw.S.SchemaName()+".cg_prize_claim("+
					"evtlog_id,block_num,time_stamp,tx_id,contract_aid,"+
					"prize_num,token_id,winner_aid,amount,donation_evt_id"+
					") VALUES($1,$2,TO_TIMESTAMP($3),$4,$5,$6,$7,$8,$9,$10)"
	_,err := sw.S.Db().Exec(query,
		evt.EvtId,
		evt.BlockNum,
		evt.TimeStamp,
		evt.TxId,
		contract_aid,
		evt.PrizeNum,
		evt.TokenId,
		winner_aid,
		evt.Amount,
		evt.DonationEvtId,
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
	query =  "INSERT INTO "+sw.S.SchemaName()+".cg_bid("+
					"evtlog_id,block_num,time_stamp,tx_id,contract_aid,"+
					"bidder_aid,rwalk_nft_id,bid_price,erc20_amount,prize_time,msg,round_num,bid_type,num_cst_tokens"+
					") VALUES($1,$2,TO_TIMESTAMP($3),$4,$5,$6,$7,$8,$9,TO_TIMESTAMP($10),$11,$12,$13,$14)"
	_,err := sw.S.Db().Exec(query,
		evt.EvtId,
		evt.BlockNum,
		evt.TimeStamp,
		evt.TxId,
		contract_aid,
		bidder_aid,
		evt.RandomWalkTokenId,
		evt.BidPrice,
		evt.ERC20_Value,
		evt.PrizeTime,
		evt.Message,
		evt.RoundNum,
		evt.BidType,
		evt.NumCSTTokens,
	)
	if err != nil {
		sw.S.Log_msg(fmt.Sprintf("DB error: can't insert into cg_bid table: %v\n",err))
		os.Exit(1)
	}
}
func (sw *SQLStorageWrapper) Insert_donation(evt *p.CGDonationEvent ) {

	contract_aid := sw.S.Lookup_or_create_address(evt.ContractAddr,0, 0)
	donor_aid := sw.S.Lookup_or_create_address(evt.DonorAddr,0, 0)
	var query string
	query =  "INSERT INTO "+sw.S.SchemaName()+".cg_donation("+
					"evtlog_id,block_num,time_stamp,tx_id,contract_aid,"+
					"donor_aid,amount"+
					") VALUES($1,$2,TO_TIMESTAMP($3),$4,$5,$6,$7)"
	_,err := sw.S.Db().Exec(query,
		evt.EvtId,
		evt.BlockNum,
		evt.TimeStamp,
		evt.TxId,
		contract_aid,
		donor_aid,
		evt.Amount,
	)
	if err != nil {
		sw.S.Log_msg(fmt.Sprintf("DB error: can't insert into cg_donation table: %v\n",err))
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
	query =  "INSERT INTO "+sw.S.SchemaName()+".cg_charity_updated("+
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
		sw.S.Log_msg(fmt.Sprintf("DB error: can't insert into cg_charity_updated table: %v\n",err))
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
func (sw *SQLStorageWrapper) Insert_raffle_deposit(evt *p.CGRaffleDeposit) {

	contract_aid := sw.S.Lookup_or_create_address(evt.ContractAddr,0, 0)
	winner_aid := sw.S.Lookup_or_create_address(evt.WinnerAddr,0, 0)

	var query string
	query =  "INSERT INTO "+sw.S.SchemaName()+".cg_raffle_deposit ("+
					"evtlog_id,block_num,time_stamp,tx_id,contract_aid,"+
					"winner_aid,round_num,amount"+
					") VALUES($1,$2,TO_TIMESTAMP($3),$4,$5,$6,$7,$8)"
	_,err := sw.S.Db().Exec(query,
		evt.EvtId,
		evt.BlockNum,
		evt.TimeStamp,
		evt.TxId,
		contract_aid,
		winner_aid,
		evt.Round,
		evt.Amount,
	)
	if err != nil {
		sw.S.Log_msg(fmt.Sprintf("DB error: can't insert into cg_raffle_deposit table: %v\n",err))
		os.Exit(1)
	}
}
func (sw *SQLStorageWrapper) Insert_raffle_withdrawal(evt *p.CGRaffleWithdrawal) {

	contract_aid := sw.S.Lookup_or_create_address(evt.ContractAddr,0, 0)
	winner_aid := sw.S.Lookup_or_create_address(evt.WinnerAddr,0, 0)

	var query string
	query =  "INSERT INTO "+sw.S.SchemaName()+".cg_raffle_withdrawal("+
					"evtlog_id,block_num,time_stamp,tx_id,contract_aid,"+
					"winner_aid,amount"+
					") VALUES($1,$2,TO_TIMESTAMP($3),$4,$5,$6,$7)"
	_,err := sw.S.Db().Exec(query,
		evt.EvtId,
		evt.BlockNum,
		evt.TimeStamp,
		evt.TxId,
		contract_aid,
		winner_aid,
		evt.Amount,
	)
	if err != nil {
		sw.S.Log_msg(fmt.Sprintf("DB error: can't insert into cg_raffle_withdrawal table: %v\n",err))
		os.Exit(1)
	}
}
func (sw *SQLStorageWrapper) Insert_raffle_nft_winner(evt *p.CGRaffleNFTWinner) {

	contract_aid := sw.S.Lookup_or_create_address(evt.ContractAddr,0, 0)
	winner_aid := sw.S.Lookup_or_create_address(evt.WinnerAddr,0, 0)

	var query string
	query =  "INSERT INTO "+sw.S.SchemaName()+".cg_raffle_nft_winner ("+
					"evtlog_id,block_num,time_stamp,tx_id,contract_aid,"+
					"winner_aid,round_num,token_id,winner_idx"+
					") VALUES($1,$2,TO_TIMESTAMP($3),$4,$5,$6,$7,$8,$9)"
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
	)
	if err != nil {
		sw.S.Log_msg(fmt.Sprintf("DB error: can't insert into cg_raffle_nft_winner table: %v\n",err))
		os.Exit(1)
	}
}
func (sw *SQLStorageWrapper) Insert_raffle_nft_claimed(evt *p.CGRaffleNFTClaimed) {

	contract_aid := sw.S.Lookup_or_create_address(evt.ContractAddr,0, 0)
	winner_aid := sw.S.Lookup_or_create_address(evt.WinnerAddr,0, 0)

	var query string
	query =  "INSERT INTO "+sw.S.SchemaName()+".cg_raffle_nft_claimed ("+
					"evtlog_id,block_num,time_stamp,tx_id,contract_aid,"+
					"winner_aid,token_id"+
					") VALUES($1,$2,TO_TIMESTAMP($3),$4,$5,$6,$7)"
	_,err := sw.S.Db().Exec(query,
		evt.EvtId,
		evt.BlockNum,
		evt.TimeStamp,
		evt.TxId,
		contract_aid,
		winner_aid,
		evt.TokenId,
	)
	if err != nil {
		sw.S.Log_msg(fmt.Sprintf("DB error: can't insert into cg_raffle_nft_claimed table: %v\n",err))
		os.Exit(1)
	}
}
func (sw *SQLStorageWrapper) Insert_donated_nft_claimed(evt *p.CGDonatedNFTClaimed) {

	contract_aid := sw.S.Lookup_or_create_address(evt.ContractAddr,0, 0)
	token_aid := sw.S.Lookup_or_create_address(evt.TokenAddr,0, 0)
	winner_aid := sw.S.Lookup_or_create_address(evt.WinnerAddr,0, 0)

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
func (sw *SQLStorageWrapper) Insert_stake_action_event(evt *p.CGStakeAction) {

	contract_aid:=sw.S.Lookup_or_create_address(evt.ContractAddr,evt.BlockNum,evt.TxId)
	staker_aid:=sw.S.Lookup_or_create_address(evt.Staker,evt.BlockNum,evt.TxId)
	var query string
	query = "INSERT INTO cg_stake_action (" +
				"evtlog_id,block_num,tx_id,time_stamp,contract_aid, "+
				"action_id,token_id,num_staked_nfts,unstake_time,staker_aid" +
			") VALUES (" +
				"$1,$2,$3,TO_TIMESTAMP($4),$5,$6,$7,$8,TO_TIMESTAMP($9),$10"+
			")"
	_,err := sw.S.Db().Exec(query,
		evt.EvtId,
		evt.BlockNum,
		evt.TxId,
		evt.TimeStamp,
		contract_aid,
		evt.ActionId,
		evt.TokenId,
		evt.TotalNfts,
		evt.UnstakeTime,
		staker_aid,
	)
	if err != nil {
		sw.S.Log_msg(fmt.Sprintf("DB error: can't insert into cg_stake_action table: %v\n",err))
		os.Exit(1)
	}
}
func (sw *SQLStorageWrapper) Insert_unstake_action_event(evt *p.CGUnstakeAction) {

	contract_aid:=sw.S.Lookup_or_create_address(evt.ContractAddr,evt.BlockNum,evt.TxId)
	staker_aid:=sw.S.Lookup_or_create_address(evt.Staker,evt.BlockNum,evt.TxId)
	var query string
	query = "INSERT INTO cg_unstake_action (" +
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
		evt.TokenId,
		evt.TotalNfts,
		staker_aid,
	)
	if err != nil {
		sw.S.Log_msg(fmt.Sprintf("DB error: can't insert into cg_unstake_action table: %v\n",err))
		os.Exit(1)
	}
}
func (sw *SQLStorageWrapper) Insert_eth_deposit_event(evt *p.CGEthDeposit) {

	contract_aid:=sw.S.Lookup_or_create_address(evt.ContractAddr,evt.BlockNum,evt.TxId)
	var query string
	query = "INSERT INTO cg_eth_deposit(" +
				"evtlog_id,block_num,tx_id,time_stamp,contract_aid, "+
				"deposit_time,round_num,deposit_num,num_staked_nfts,amount,amount_per_staker,modulo,accum_modulo" +
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
		evt.DepositNum,
		evt.NumStakedNfts,
		evt.Amount,
		evt.AmountPerStaker,
		evt.Modulo,
		evt.AccumModulo,
	)
	if err != nil {
		sw.S.Log_msg(fmt.Sprintf("DB error: can't insert into cg_eth_deposit table: %v\n",err))
		os.Exit(1)
	}
}
func (sw *SQLStorageWrapper) Insert_claim_reward_event(evt *p.CGClaimReward) {

	contract_aid:=sw.S.Lookup_or_create_address(evt.ContractAddr,evt.BlockNum,evt.TxId)
	staker_aid:=sw.S.Lookup_or_create_address(evt.Staker,evt.BlockNum,evt.TxId)
	var query string
	query = "INSERT INTO cg_claim_reward(" +
				"evtlog_id,block_num,tx_id,time_stamp,contract_aid, "+
				"action_id,deposit_id,reward,staker_aid" +
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
		evt.DepositId,
		evt.Reward,
		staker_aid,
	)
	if err != nil {
		sw.S.Log_msg(fmt.Sprintf("DB error: can't insert into cg_claim_Reward table: %v\n",err))
		os.Exit(1)
	}
}
func (sw *SQLStorageWrapper) Insert_marketing_reward_sent_event(evt *p.CGMarketingRewardSent) {

	contract_aid:=sw.S.Lookup_or_create_address(evt.ContractAddr,evt.BlockNum,evt.TxId)
	marketer_aid:=sw.S.Lookup_or_create_address(evt.Marketer,evt.BlockNum,evt.TxId)
	var query string
	query = "INSERT INTO cg_mkt_reward(" +
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
	query = "INSERT INTO cg_transfer(" +
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
		sw.S.Log_msg(fmt.Sprintf("DB error: can't insert into cg_transfer table: %v\n",err))
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
	query = "INSERT INTO cg_erc20_transfer(" +
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
	query = "INSERT INTO cg_adm_charity_pcent (" +
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
	query = "INSERT INTO cg_adm_prize_pcent (" +
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
		sw.S.Log_msg(fmt.Sprintf("DB error: can't insert into cg_adm_prize_pcent table: %v\n",err))
		os.Exit(1)
	}
}
func (sw *SQLStorageWrapper) Insert_cosmic_game_raffle_percentage_changed_event(evt *p.CGRafflePercentageChanged) {

	contract_aid:=sw.S.Lookup_or_create_address(evt.Contract,evt.BlockNum,evt.TxId)
	var query string
	query = "INSERT INTO cg_adm_raffle_pcent (" +
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
func (sw *SQLStorageWrapper) Insert_cosmic_game_num_raffle_eth_winners_per_round_changed_event(evt *p.CGNumRaffleWinnersPerRoundChanged) {

	contract_aid:=sw.S.Lookup_or_create_address(evt.Contract,evt.BlockNum,evt.TxId)
	var query string
	query = "INSERT INTO cg_adm_raf_eth_winners(" +
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
		evt.NewNumRaffleWinnersPerRound,
	)
	if err != nil {
		sw.S.Log_msg(fmt.Sprintf("DB error: can't insert into cg_adm_raf_eth_winners table: %v\n",err))
		os.Exit(1)
	}
}
func (sw *SQLStorageWrapper) Insert_cosmic_game_num_raffle_nft_winners_per_round_changed_event(evt *p.CGNumRaffleNFTWinnersPerRoundChanged) {

	contract_aid:=sw.S.Lookup_or_create_address(evt.Contract,evt.BlockNum,evt.TxId)
	var query string
	query = "INSERT INTO cg_adm_raf_nft_winners(" +
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
		evt.NewNumRaffleNFTWinnersPerRound,
	)
	if err != nil {
		sw.S.Log_msg(fmt.Sprintf("DB error: can't insert into cg_adm_raf_eth_winners table: %v\n",err))
		os.Exit(1)
	}
}
func (sw *SQLStorageWrapper) Insert_cosmic_game_num_raffle_nft_holders_per_round_changed_event(evt *p.CGNumRaffleNFTHoldersPerRoundChanged) {

	contract_aid:=sw.S.Lookup_or_create_address(evt.Contract,evt.BlockNum,evt.TxId)
	var query string
	query = "INSERT INTO cg_adm_raf_nft_holders(" +
				"evtlog_id,block_num,tx_id,time_stamp,contract_aid, "+
				"num_holders" +
			") VALUES (" +
				"$1,$2,$3,TO_TIMESTAMP($4),$5,$6"+
			")"
	_,err := sw.S.Db().Exec(query,
		evt.EvtId,
		evt.BlockNum,
		evt.TxId,
		evt.TimeStamp,
		contract_aid,
		evt.NewNumRaffleNFTHoldersPerRound,
	)
	if err != nil {
		sw.S.Log_msg(fmt.Sprintf("DB error: can't insert into cg_adm_raf_eth_winners table: %v\n",err))
		os.Exit(1)
	}
}
func (sw *SQLStorageWrapper) Insert_cosmic_game_staking_percentage_changed_event(evt *p.CGStakingPercentageChanged) {

	contract_aid:=sw.S.Lookup_or_create_address(evt.Contract,evt.BlockNum,evt.TxId)
	var query string
	query = "INSERT INTO cg_adm_stake_pcent (" +
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
func (sw *SQLStorageWrapper) Insert_cosmic_game_system_mode_changed_event(evt *p.CGSystemModeChanged) {

	contract_aid:=sw.S.Lookup_or_create_address(evt.Contract,evt.BlockNum,evt.TxId)
	var query string
	query = "INSERT INTO cg_adm_sysmode (" +
				"evtlog_id,block_num,tx_id,time_stamp,contract_aid, "+
				"sysmode" +
			") VALUES (" +
				"$1,$2,$3,TO_TIMESTAMP($4),$5,$6"+
			")"
	_,err := sw.S.Db().Exec(query,
		evt.EvtId,
		evt.BlockNum,
		evt.TxId,
		evt.TimeStamp,
		contract_aid,
		evt.NewSystemMode,
	)
	if err != nil {
		sw.S.Log_msg(fmt.Sprintf("DB error: can't insert into cg_adm_sysmode table: %v\n",err))
		os.Exit(1)
	}
}
func (sw *SQLStorageWrapper) Insert_cosmic_game_charity_address_changed_event(evt *p.CGCharityAddressChanged) {

	contract_aid:=sw.S.Lookup_or_create_address(evt.Contract,evt.BlockNum,evt.TxId)
	new_charity_aid:=sw.S.Lookup_or_create_address(evt.NewCharity,evt.BlockNum,evt.TxId)
	var query string
	query = "INSERT INTO cg_adm_charity_addr(" +
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
		sw.S.Log_msg(fmt.Sprintf("DB error: can't insert into cg_adm_charity_addr table: %v\n",err))
		os.Exit(1)
	}
}
func (sw *SQLStorageWrapper) Insert_cosmic_game_random_walk_address_changed_event(evt *p.CGRandomWalkAddressChanged) {

	contract_aid:=sw.S.Lookup_or_create_address(evt.Contract,evt.BlockNum,evt.TxId)
	new_rwalk_aid:=sw.S.Lookup_or_create_address(evt.NewRandomWalk,evt.BlockNum,evt.TxId)
	var query string
	query = "INSERT INTO cg_adm_rwalk_addr(" +
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
func (sw *SQLStorageWrapper) Insert_cosmic_game_raffle_wallet_address_changed_event(evt *p.CGRaffleWalletAddressChanged) {

	contract_aid:=sw.S.Lookup_or_create_address(evt.Contract,evt.BlockNum,evt.TxId)
	new_raffle_aid:=sw.S.Lookup_or_create_address(evt.NewRaffleWallet,evt.BlockNum,evt.TxId)
	var query string
	query = "INSERT INTO cg_adm_raffle_addr(" +
				"evtlog_id,block_num,tx_id,time_stamp,contract_aid, "+
				"new_raffle_aid" +
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
		sw.S.Log_msg(fmt.Sprintf("DB error: can't insert into cg_adm_rwalk_addr table: %v\n",err))
		os.Exit(1)
	}
}
func (sw *SQLStorageWrapper) Insert_cosmic_game_staking_wallet_address_changed_event(evt *p.CGStakingWalletAddressChanged) {

	contract_aid:=sw.S.Lookup_or_create_address(evt.Contract,evt.BlockNum,evt.TxId)
	new_staking_aid:=sw.S.Lookup_or_create_address(evt.NewStakingWallet,evt.BlockNum,evt.TxId)
	var query string
	query = "INSERT INTO cg_adm_staking_addr(" +
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
		sw.S.Log_msg(fmt.Sprintf("DB error: can't insert into cg_adm_staking_addr table: %v\n",err))
		os.Exit(1)
	}
}
func (sw *SQLStorageWrapper) Insert_cosmic_game_marketing_wallet_address_changed_event(evt *p.CGMarketingWalletAddressChanged) {

	contract_aid:=sw.S.Lookup_or_create_address(evt.Contract,evt.BlockNum,evt.TxId)
	new_marketing_aid:=sw.S.Lookup_or_create_address(evt.NewMarketingWallet,evt.BlockNum,evt.TxId)
	var query string
	query = "INSERT INTO cg_adm_marketing_addr(" +
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
		new_marketing_aid,
	)
	if err != nil {
		sw.S.Log_msg(fmt.Sprintf("DB error: can't insert into cg_adm_marketing_addr table: %v\n",err))
		os.Exit(1)
	}
}
func (sw *SQLStorageWrapper) Insert_cosmic_token_address_changed_event(evt *p.CGCosmicTokenAddressChanged) {

	contract_aid:=sw.S.Lookup_or_create_address(evt.Contract,evt.BlockNum,evt.TxId)
	new_costok_aid:=sw.S.Lookup_or_create_address(evt.NewCosmicToken,evt.BlockNum,evt.TxId)
	var query string
	query = "INSERT INTO cg_adm_costok_addr(" +
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
	query = "INSERT INTO cg_adm_cossig_addr(" +
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
func (sw *SQLStorageWrapper) Insert_business_logic_address_changed_event(evt *p.CGBusinessLogicAddressChanged) {

	contract_aid:=sw.S.Lookup_or_create_address(evt.Contract,evt.BlockNum,evt.TxId)
	new_blogic_aid:=sw.S.Lookup_or_create_address(evt.NewContractAddress,evt.BlockNum,evt.TxId)
	var query string
	query = "INSERT INTO cg_adm_blogic_addr(" +
				"evtlog_id,block_num,tx_id,time_stamp,contract_aid, "+
				"new_blogic_aid" +
			") VALUES (" +
				"$1,$2,$3,TO_TIMESTAMP($4),$5,$6"+
			")"
	_,err := sw.S.Db().Exec(query,
		evt.EvtId,
		evt.BlockNum,
		evt.TxId,
		evt.TimeStamp,
		contract_aid,
		new_blogic_aid,
	)
	if err != nil {
		sw.S.Log_msg(fmt.Sprintf("DB error: can't insert into cg_adm_blogic_addr table: %v\n",err))
		os.Exit(1)
	}
}
func (sw *SQLStorageWrapper) Insert_time_increase_changed_event(evt *p.CGTimeIncreaseChanged) {

	contract_aid:=sw.S.Lookup_or_create_address(evt.Contract,evt.BlockNum,evt.TxId)
	var query string
	query = "INSERT INTO cg_adm_time_inc(" +
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
	query = "INSERT INTO cg_adm_timeout_claimprize(" +
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
