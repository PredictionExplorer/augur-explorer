-- +goose Up
CREATE TABLE cg_prize_claim ( -- ICosmicSignatureGame.sol:MainPrizeClaimed
	id						BIGSERIAL PRIMARY KEY,
	evtlog_id				BIGINT REFERENCES evt_log(id) ON DELETE CASCADE,
	block_num				BIGINT NOT NULL,
	tx_id					BIGINT NOT NULL,
	time_stamp				TIMESTAMPTZ NOT NULL,
	contract_aid			BIGINT NOT NULL,
	round_num				BIGINT NOT NULL,
	winner_aid				BIGINT NOT NULL,
	token_id				BIGINT NOT NULL,
	timeout					BIGINT NOT NULL,	-- timeoutTimeToWithdrawSecondaryPrizes
	amount					DECIMAL DEFAULT 0,	-- ethPrizeAmount
	cst_amount				DECIMAL DEFAULT 0,	-- cstPrizeAmount
	UNIQUE(evtlog_id)
);
CREATE TABLE cg_prize( -- Generic prize record , that unifies all prizes , populated automatically with triggers
	round_num				BIGINT NOT NULL,
	winner_index			BIGINT NOT NULL,
	ptype					SMALLINT DEFAULT -1, -- provided by each prize winning event Codes: 
														-- 0 - Main Prize ETH
														-- 1 - Main Prize CST (ERC20)
														-- 2 - Main Prize CS NFT
														-- 3 - Last CST Bidder CS NFT (ERC721)
														-- 4 - Last CST Bidder ERC20 (CST)
														-- 5 - Endurance Champion CS NFT
														-- 6 - Endurance Champion ERC20 (CST)
														-- 7 - Chrono Warrior ETH
														-- 8 - Chrono Warrior CST (ERC20)
														-- 9 - Chrono Warrior CS NFT
														-- 10 - Raffle ETH (for bidders)
														-- 11 - Raffle CST (for bidders)
														-- 12 - Raffle CS NFT (for bidders)
														-- 13 - Raffle CST (for RandomWalk stakers)
														-- 14 - Raffle CS NFT (for RandomWalk stakers)
														-- 15 - Staking Deposit ETH (for CS NFT stakers)
	PRIMARY KEY(round_num,winner_index,ptype)
);
CREATE TABLE cg_bid ( -- ICosmicSignatureGame.sol:BidPlaced
	id				BIGSERIAL PRIMARY KEY,
	evtlog_id		BIGINT REFERENCES evt_log(id) ON DELETE CASCADE,
	block_num		BIGINT NOT NULL,
	tx_id			BIGINT NOT NULL,
	time_stamp		TIMESTAMPTZ NOT NULL,
	contract_aid	BIGINT NOT NULL,
	bidder_aid		BIGINT NOT NULL,
	rwalk_nft_id	BIGINT NOT NULL,	--token_id of RandomWalk, if present
	round_num		BIGINT NOT NULL,
	bid_type		SMALLINT NOT NULL,  --  0 = ETH, 1 = RandomWalk, 2 = CST
	bid_position	BIGINT NOT NULL,	-- Position of this bid within the round (1st, 2nd, 3rd, etc.)
	prize_time		TIMESTAMPTZ NOT NULL,
	eth_price		DECIMAL NOT NULL,	-- PaidEthPrice (or -1 if CST bid)
	cst_price		DECIMAL NOT NULL,	-- PaidCstPrice (or -1 if ETH bid)
	cst_reward		DECIMAL DEFAULT 0,	-- CST reward amount for this bid (from cstRewardAmountForBidding at time of bid)
	bid_cst_reward_amount	DECIMAL DEFAULT -1,	-- IBiddingV2 BidPlaced (topic 0x1d1f406c…); -1 = legacy BidPlaced
	cst_dutch_auction_duration DECIMAL DEFAULT -1,	-- per-bid auction duration from IBiddingV2 BidPlaced; -1 = legacy
	msg				TEXT,
	UNIQUE(evtlog_id)
);
CREATE TABLE cg_eth_donated ( -- IEthDonations.sol:EthDonated
	id				BIGSERIAL PRIMARY KEY,
	evtlog_id		BIGINT REFERENCES evt_log(id) ON DELETE CASCADE,
	block_num		BIGINT NOT NULL,
	tx_id			BIGINT NOT NULL,
	time_stamp		TIMESTAMPTZ NOT NULL,
	contract_aid	BIGINT NOT NULL,
	donor_aid		BIGINT NOT NULL,
	round_num		BIGINT NOT NULL,
	amount			DECIMAL NOT NULL,
	UNIQUE(evtlog_id)
);
CREATE TABLE cg_eth_donated_wi ( -- IEthDonations.sol:EthDonatedWithInfo
	id				BIGSERIAL PRIMARY KEY,
	evtlog_id		BIGINT REFERENCES evt_log(id) ON DELETE CASCADE,
	block_num		BIGINT NOT NULL,
	tx_id			BIGINT NOT NULL,
	time_stamp		TIMESTAMPTZ NOT NULL,
	contract_aid	BIGINT NOT NULL,
	donor_aid		BIGINT NOT NULL,
	round_num		BIGINT NOT NULL,
	record_id		BIGINT NOT NULL,
	amount			DECIMAL NOT NULL,
	UNIQUE(evtlog_id),
	UNIQUE(record_id)
);
CREATE TABLE cg_donation_json ( -- JSON data related to donation (this table is complementary to cg_eth_donated_wi table)
	record_id		BIGINT PRIMARY KEY REFERENCES cg_eth_donated_wi(record_id) ON DELETE CASCADE,
	data			TEXT
);
CREATE TABLE cg_donation_received ( -- ICharityWallet.sol:DonationReceived
	id				BIGSERIAL PRIMARY KEY,
	evtlog_id		BIGINT REFERENCES evt_log(id) ON DELETE CASCADE,
	block_num		BIGINT NOT NULL,
	tx_id			BIGINT NOT NULL,
	time_stamp		TIMESTAMPTZ NOT NULL,
	contract_aid	BIGINT NOT NULL,
	donor_aid		BIGINT NOT NULL,
	amount			DECIMAL NOT NULL,
	round_num		BIGINT NOT NULL DEFAULT -1,
	UNIQUE(evtlog_id)
);
CREATE TABLE cg_donation_sent ( -- ICharityWallet.sol:DonationSent
	id				BIGSERIAL PRIMARY KEY,
	evtlog_id		BIGINT REFERENCES evt_log(id) ON DELETE CASCADE,
	block_num		BIGINT NOT NULL,
	tx_id			BIGINT NOT NULL,
	time_stamp		TIMESTAMPTZ NOT NULL,
	contract_aid	BIGINT NOT NULL,
	charity_aid		BIGINT NOT NULL,
	amount			DECIMAL NOT NULL,
	UNIQUE(evtlog_id)
);
CREATE TABLE cg_erc20_donation ( -- IPrizesWallet.sol:TokenDonated
	id				BIGSERIAL PRIMARY KEY,
	evtlog_id		BIGINT REFERENCES evt_log(id) ON DELETE CASCADE,
	block_num		BIGINT NOT NULL,
	tx_id			BIGINT NOT NULL,
	time_stamp		TIMESTAMPTZ NOT NULL,
	contract_aid	BIGINT NOT NULL,
	round_num		BIGINT NOT NULL,
	donor_aid		BIGINT NOT NULL,
	token_aid		BIGINT NOT NULL,	-- this is address id (table address)
	amount			DECIMAL NOT NULL,
	bid_id			BIGINT NOT NULL,		-- id of the related `cg_bid` record
	UNIQUE(evtlog_id)
);
CREATE TABLE cg_nft_donation ( -- IPrizesWallet.sol:NftDonated
	id				BIGSERIAL PRIMARY KEY,
	evtlog_id		BIGINT REFERENCES evt_log(id) ON DELETE CASCADE,
	block_num		BIGINT NOT NULL,
	tx_id			BIGINT NOT NULL,
	time_stamp		TIMESTAMPTZ NOT NULL,
	contract_aid	BIGINT NOT NULL,
	round_num		BIGINT NOT NULL,
	donor_aid		BIGINT NOT NULL,
	token_aid		BIGINT NOT NULL,	-- this is address id (table address)
	token_id		BIGINT NOT NULL,	-- this is tokenID
	idx				BIGINT NOT NULL,	-- Index field of NFTDonationEvent
	bid_id			BIGINT NOT NULL,		-- id of the related `cg_bid` record
	token_uri		TEXT NOT NULL,
	UNIQUE(evtlog_id)
);
CREATE TABLE cg_charity_receiver_changed ( -- ICharityWallet.sol:CharityAddressChanged
	id				BIGSERIAL PRIMARY KEY,
	evtlog_id		BIGINT REFERENCES evt_log(id) ON DELETE CASCADE,
	block_num		BIGINT NOT NULL,
	tx_id			BIGINT NOT NULL,
	time_stamp		TIMESTAMPTZ NOT NULL,
	contract_aid	BIGINT NOT NULL,
	charity_aid		BIGINT NOT NULL,
	UNIQUE(evtlog_id)
);
CREATE TABLE cg_token_name ( -- ICosmicSignatureNft.sol:NftNameChanged
	id				BIGSERIAL PRIMARY KEY,
	evtlog_id		BIGINT REFERENCES evt_log(id) ON DELETE CASCADE,
	block_num		BIGINT NOT NULL,
	tx_id			BIGINT NOT NULL,
	time_stamp		TIMESTAMPTZ NOT NULL,
	contract_aid	BIGINT NOT NULL,
	token_id		BIGINT NOT NULL,
	token_name		TEXT NOT NULL,
	UNIQUE(evtlog_id)
);
CREATE TABLE cg_mint_event ( -- ICosmicSignatureNft.sol:NftMinted
	id				BIGSERIAL PRIMARY KEY,
	evtlog_id		BIGINT REFERENCES evt_log(id) ON DELETE CASCADE,
	block_num		BIGINT NOT NULL,
	tx_id			BIGINT NOT NULL,
	time_stamp		TIMESTAMPTZ NOT NULL,
	contract_aid	BIGINT NOT NULL,
	owner_aid		BIGINT NOT NULL,
	token_id		BIGINT NOT NULL,
	cur_owner_aid	BIGINT NOT NULL,
	round_num		BIGINT NOT NULL,
	seed			TEXT NOT NULL,
	token_name		TEXT DEFAULT '', -- last name set via setTokenName()
	UNIQUE(evtlog_id)
);
CREATE TABLE cg_prize_deposit ( -- IPrizesWallet.sol:EthReceived
	id				BIGSERIAL PRIMARY KEY,
	evtlog_id		BIGINT REFERENCES evt_log(id) ON DELETE CASCADE,
	block_num		BIGINT NOT NULL,
	tx_id			BIGINT NOT NULL,
	time_stamp		TIMESTAMPTZ NOT NULL,
	contract_aid	BIGINT NOT NULL,
	winner_aid		BIGINT NOT NULL,
	round_num		BIGINT NOT NULL,
	winner_index	BIGINT NOT NULL,
	amount			DECIMAL NOT NULL,
	claimed			BOOLEAN DEFAULT 'F',	-- upon withdrawal is set to TRUE
	withdrawal_id	BIGINT DEFAULT 0, -- at withdrawal set to evtlog_id of bw_raffle_Withdrawal
	UNIQUE(evtlog_id)
);
CREATE TABLE cg_prize_withdrawal ( -- IPrizesWallet.sol:EthWithdrawn
	id				BIGSERIAL PRIMARY KEY,
	evtlog_id		BIGINT REFERENCES evt_log(id) ON DELETE CASCADE,
	block_num		BIGINT NOT NULL,
	tx_id			BIGINT NOT NULL,
	time_stamp		TIMESTAMPTZ NOT NULL,
	contract_aid	BIGINT NOT NULL,
	round_num		BIGINT NOT NULL,
	winner_aid		BIGINT NOT NULL,
	beneficiary_aid	BIGINT NOT NULL,	-- Who actually claimed (can differ from winner after timeout)
	amount			DECIMAL NOT NULL,
	UNIQUE(evtlog_id)
);
CREATE TABLE cg_raffle_nft_prize ( -- ICosmicSignatureGame.sol:RaffleWinnerPrizePaid
	id				BIGSERIAL PRIMARY KEY,
	evtlog_id		BIGINT REFERENCES evt_log(id) ON DELETE CASCADE,
	block_num		BIGINT NOT NULL,
	tx_id			BIGINT NOT NULL,
	time_stamp		TIMESTAMPTZ NOT NULL,
	contract_aid	BIGINT NOT NULL,
	winner_aid		BIGINT NOT NULL,
	round_num		BIGINT NOT NULL,
	token_id		BIGINT NOT NULL,
	winner_idx		BIGINT NOT NULL,
	cst_amount		DECIMAL NOT NULL,
	is_rwalk		BOOLEAN NOT NULL,
	is_staker		BOOLEAN NOT NULL,
	UNIQUE(evtlog_id)
);
CREATE TABLE cg_raffle_eth_prize ( -- ICosmicSignatureGame.sol:RaffleWinnerBidderEthPrizeAllocated
	id				BIGSERIAL PRIMARY KEY,
	evtlog_id		BIGINT REFERENCES evt_log(id) ON DELETE CASCADE,
	block_num		BIGINT NOT NULL,
	tx_id			BIGINT NOT NULL,
	time_stamp		TIMESTAMPTZ NOT NULL,
	contract_aid	BIGINT NOT NULL,
	winner_aid		BIGINT NOT NULL,
	round_num		BIGINT NOT NULL,
	winner_idx		BIGINT NOT NULL,
	amount			DECIMAL NOT NULL,
	UNIQUE(evtlog_id)
);
CREATE TABLE cg_endurance_prize ( -- ICosmicSignatureGame.sol:EnduranceChampionPrizePaid
	-- Note: The Solidity event does not emit a winner_index field.
	-- There is exactly one endurance champion per round, so winner_index is implicitly 0.
	id				BIGSERIAL PRIMARY KEY,
	evtlog_id		BIGINT REFERENCES evt_log(id) ON DELETE CASCADE,
	block_num		BIGINT NOT NULL,
	tx_id			BIGINT NOT NULL,
	time_stamp		TIMESTAMPTZ NOT NULL,
	contract_aid	BIGINT NOT NULL,
	winner_aid		BIGINT NOT NULL,
	round_num		BIGINT NOT NULL,
	erc721_token_id		BIGINT NOT NULL,
	erc20_amount	DECIMAL NOT NULL,
	UNIQUE(round_num),
	UNIQUE(evtlog_id)
);
CREATE TABLE cg_lastcst_prize ( -- ICosmicSignatureGame.sol:LastCstBidderPrizePaid
	-- Note: The Solidity event does not emit a winner_index field.
	-- There is exactly one last CST bidder per round, so winner_index is implicitly 0.
	id				BIGSERIAL PRIMARY KEY,
	evtlog_id		BIGINT REFERENCES evt_log(id) ON DELETE CASCADE,
	block_num		BIGINT NOT NULL,
	tx_id			BIGINT NOT NULL,
	time_stamp		TIMESTAMPTZ NOT NULL,
	contract_aid	BIGINT NOT NULL,
	winner_aid		BIGINT NOT NULL,
	round_num		BIGINT NOT NULL,
	erc721_token_id		BIGINT NOT NULL,
	erc20_amount	DECIMAL NOT NULL,
	UNIQUE(round_num),
	UNIQUE(evtlog_id)
);
CREATE TABLE cg_chrono_warrior_prize ( -- ICosmicSignatureGame.sol:ChronoWarriorPrizePaid
	id				BIGSERIAL PRIMARY KEY,
	evtlog_id		BIGINT REFERENCES evt_log(id) ON DELETE CASCADE,
	block_num		BIGINT NOT NULL,
	tx_id			BIGINT NOT NULL,
	time_stamp		TIMESTAMPTZ NOT NULL,
	contract_aid	BIGINT NOT NULL,
	winner_aid		BIGINT NOT NULL,
	round_num		BIGINT NOT NULL,
	winner_index	BIGINT NOT NULL,
	eth_amount		DECIMAL NOT NULL,
	cst_amount		DECIMAL NOT NULL,
	nft_id			BIGINT NOT NULL,
	UNIQUE(round_num),
	UNIQUE(evtlog_id)
);
CREATE TABLE cg_donated_tok_claimed ( -- IPrizesWallet.sol:DonatedTokenClaimed
	id				BIGSERIAL PRIMARY KEY,
	evtlog_id		BIGINT REFERENCES evt_log(id) ON DELETE CASCADE,
	block_num		BIGINT NOT NULL,
	tx_id			BIGINT NOT NULL,
	time_stamp		TIMESTAMPTZ NOT NULL,
	contract_aid	BIGINT NOT NULL,
	round_num		BIGINT NOT NULL,
	idx				BIGINT NOT NULL,
	token_aid		BIGINT NOT NULL,
	winner_aid		BIGINT NOT NULL,
	amount			DECIMAL NOT NULL,
	UNIQUE(evtlog_id)
);
CREATE TABLE cg_donated_nft_claimed ( -- IPrizesWallet.sol:DonatedNftClaimed
	id				BIGSERIAL PRIMARY KEY,
	evtlog_id		BIGINT REFERENCES evt_log(id) ON DELETE CASCADE,
	block_num		BIGINT NOT NULL,
	tx_id			BIGINT NOT NULL,
	time_stamp		TIMESTAMPTZ NOT NULL,
	contract_aid	BIGINT NOT NULL,
	round_num		BIGINT NOT NULL,
	idx				BIGINT NOT NULL,
	token_aid		BIGINT NOT NULL,
	winner_aid		BIGINT NOT NULL,
	token_id		DECIMAL NOT NULL,
	UNIQUE(evtlog_id)
);
CREATE TABLE cg_nft_unstaked_rwalk ( -- IStakingWalletNftBase.sol:NftUnstaked
	-- Note: The Solidity NftUnstaked event does not emit a round_num field.
	-- The round_num is derived via SQL trigger from the most recent prize claim.
	-- Default -1 indicates round has not been determined yet (will be populated by trigger).
	id				BIGSERIAL PRIMARY KEY,
	evtlog_id		BIGINT REFERENCES evt_log(id) ON DELETE CASCADE,
	block_num		BIGINT NOT NULL,
	tx_id			BIGINT NOT NULL,
	time_stamp		TIMESTAMPTZ NOT NULL,
	contract_aid	BIGINT NOT NULL,
	action_id		BIGINT NOT NULL,
	token_id		BIGINT NOT NULL,
	round_num		BIGINT DEFAULT -1,
	num_staked_nfts	BIGINT NOT NULL,
	staker_aid		BIGINT NOT NULL,
	UNIQUE(evtlog_id)
);
CREATE TABLE cg_nft_unstaked_cst ( -- IStakingWalletNftBase.sol:NftUnstaked
	-- Note: The Solidity NftUnstaked event does not emit a round_num field.
	-- The round_num is derived via SQL trigger from the most recent prize claim.
	-- Default -1 indicates round has not been determined yet (will be populated by trigger).
	id				BIGSERIAL PRIMARY KEY,
	evtlog_id		BIGINT REFERENCES evt_log(id) ON DELETE CASCADE,
	block_num		BIGINT NOT NULL,
	tx_id			BIGINT NOT NULL,
	time_stamp		TIMESTAMPTZ NOT NULL,
	contract_aid	BIGINT NOT NULL,
	action_id		BIGINT NOT NULL,
	action_counter	BIGINT NOT NULL,
	token_id		BIGINT NOT NULL,
	num_staked_nfts	BIGINT NOT NULL,
	round_num		BIGINT DEFAULT -1,
	staker_aid		BIGINT NOT NULL,
	reward			DECIMAL NOT NULL,
	reward_per_tok	DECIMAL NOT NULL, -- reward per token at the time of unstake
	UNIQUE(evtlog_id)
);
CREATE TABLE cg_nft_staked_cst ( -- IStakingWalletNftBase.sol:NftStaked
	-- Note: The Solidity NftStaked event does not emit a round_num field.
	-- The round_num is derived via SQL trigger from the most recent prize claim.
	-- Default -1 indicates round has not been determined yet (will be populated by trigger).
	id				BIGSERIAL PRIMARY KEY,
	evtlog_id		BIGINT REFERENCES evt_log(id) ON DELETE CASCADE,
	block_num		BIGINT NOT NULL,
	tx_id			BIGINT NOT NULL,
	time_stamp		TIMESTAMPTZ NOT NULL,
	contract_aid	BIGINT NOT NULL,
	action_id		BIGINT NOT NULL,
	token_id		BIGINT NOT NULL,
	round_num		BIGINT DEFAULT -1,
	num_staked_nfts	BIGINT NOT NULL,
	reward_per_staker	DECIMAL NOT NULL,
	staker_aid		BIGINT NOT NULL,
	claimed			BOOLEAN DEFAULT 'F',
	UNIQUE(evtlog_id)
);
CREATE TABLE cg_nft_staked_rwalk ( -- IStakingWalletNftBase.sol:NftStaked
	-- Note: The Solidity NftStaked event does not emit a round_num field.
	-- The round_num is derived via SQL trigger from the most recent prize claim.
	-- Default -1 indicates round has not been determined yet (will be populated by trigger).
	id				BIGSERIAL PRIMARY KEY,
	evtlog_id		BIGINT REFERENCES evt_log(id) ON DELETE CASCADE,
	block_num		BIGINT NOT NULL,
	tx_id			BIGINT NOT NULL,
	time_stamp		TIMESTAMPTZ NOT NULL,
	contract_aid	BIGINT NOT NULL,
	action_id		BIGINT NOT NULL,
	token_id		BIGINT NOT NULL,
	round_num		BIGINT DEFAULT -1,
	num_staked_nfts	BIGINT NOT NULL,
	staker_aid		BIGINT NOT NULL,
	UNIQUE(evtlog_id)
);
CREATE TABLE cg_staking_eth_deposit ( -- IStakingWalletCosmicSignatureNft.sol:EthDepositReceived
	id				BIGSERIAL PRIMARY KEY,
	evtlog_id		BIGINT REFERENCES evt_log(id) ON DELETE CASCADE,
	block_num		BIGINT NOT NULL,
	tx_id			BIGINT NOT NULL,
	time_stamp		TIMESTAMPTZ NOT NULL,
	contract_aid	BIGINT NOT NULL,
	round_num		BIGINT NOT NULL ,
	deposit_time	TIMESTAMPTZ NOT NULL,
	deposit_id		BIGINT NOT NULL,	-- action counter
	num_staked_nfts	BIGINT NOT NULL,		-- new tokens added between previous deposit and this deposit
	accumulated_nfts	BIGINT DEFAULT 0,	-- accumulated number of staked tokesn from previous deposits
	deposit_amount		DECIMAL NOT NULL,
	accumulated_amount	DECIMAL DEFAULT 0,
	amount_per_token	DECIMAL NOT NULL,	-- this value is for current deposit, not
	accumulated_per_token	DECIMAL DEFAULT 0,	-- this is the accumulated value from previous deposits to current deposit
	modulo			DECIMAL NOT NULL,
	accum_modulo	DECIMAL DEFAULT 0,
	UNIQUE(evtlog_id)
);
CREATE TABLE cg_first_bid ( -- ICosmicSignatureGame.sol:FirstBidPlacedInRound
	id				BIGSERIAL PRIMARY KEY,
	evtlog_id		BIGINT REFERENCES evt_log(id) ON DELETE CASCADE,
	block_num		BIGINT NOT NULL,
	tx_id			BIGINT NOT NULL,
	time_stamp		TIMESTAMPTZ NOT NULL,
	contract_aid	BIGINT NOT NULL,
	round_num		BIGINT NOT NULL,
	start_ts		BIGINT NOT NULL,
	UNIQUE(evtlog_id)
);
CREATE TABLE cg_mkt_reward ( -- IMarketingWallet.sol:RewardPaid
	id				BIGSERIAL PRIMARY KEY,
	evtlog_id		BIGINT REFERENCES evt_log(id) ON DELETE CASCADE,
	block_num		BIGINT NOT NULL,
	tx_id			BIGINT NOT NULL,
	time_stamp		TIMESTAMPTZ NOT NULL,
	contract_aid	BIGINT NOT NULL,
	amount			DECIMAL NOT NULL,
	marketer_aid	BIGINT NOT NULL,
	UNIQUE(evtlog_id)
);
CREATE TABLE cg_erc721_transfer ( -- IERC721.sol:Transfer
	id              BIGSERIAL PRIMARY KEY,
	evtlog_id       BIGINT REFERENCES evt_log(id) ON DELETE CASCADE,
	block_num       BIGINT NOT NULL,
	tx_id           BIGINT NOT NULL,
	time_stamp      TIMESTAMPTZ NOT NULL,
	contract_aid    BIGINT NOT NULL,
	token_id        BIGINT NOT NULL,
	from_aid        BIGINT NOT NULL,
	to_aid          BIGINT NOT NULL,
	otype           SMALLINT NOT NULL,-- 0-regular transfer,1-Mint,2-Burn
	UNIQUE(evtlog_id)
);
CREATE TABLE cg_erc20_transfer ( -- IERC20.sol:Transfer
	id              BIGSERIAL PRIMARY KEY,
	evtlog_id       BIGINT REFERENCES evt_log(id) ON DELETE CASCADE,
	block_num       BIGINT NOT NULL,
	tx_id           BIGINT NOT NULL,
	time_stamp      TIMESTAMPTZ NOT NULL,
	contract_aid    BIGINT NOT NULL,
	value			DECIMAL NOT NULL,
	from_aid        BIGINT NOT NULL,
	to_aid          BIGINT NOT NULL,
	otype           SMALLINT NOT NULL,-- 0-regular transfer,1-Mint,2-Burn
	UNIQUE(evtlog_id)
);
CREATE TABLE cg_costok_owner( -- CosmicToken owner info (about balance)
	owner_aid		BIGINT PRIMARY KEY,
	cur_balance		DECIMAL DEFAULT 0 --- user's balance in CosmicToken denomination
);
CREATE TABLE cg_fund_transf_err ( -- ICosmicSignatureErrors.sol:FundTransferFailed
	id              BIGSERIAL PRIMARY KEY,
	evtlog_id       BIGINT REFERENCES evt_log(id) ON DELETE CASCADE,
	block_num       BIGINT NOT NULL,
	tx_id           BIGINT NOT NULL,
	time_stamp      TIMESTAMPTZ NOT NULL,
	contract_aid    BIGINT NOT NULL,
	destination_aid	BIGINT NOT NULL,
	amount			DECIMAL NOT NULL,
	UNIQUE(evtlog_id)
);
CREATE TABLE cg_erc20_transf_err ( -- ICosmicSignatureErrors.sol:ERC20TransferFailed
	id              BIGSERIAL PRIMARY KEY,
	evtlog_id       BIGINT REFERENCES evt_log(id) ON DELETE CASCADE,
	block_num       BIGINT NOT NULL,
	tx_id           BIGINT NOT NULL,
	time_stamp      TIMESTAMPTZ NOT NULL,
	contract_aid    BIGINT NOT NULL,
	destination_aid	BIGINT NOT NULL,
	amount			DECIMAL NOT NULL,
	UNIQUE(evtlog_id)
);
CREATE TABLE cg_funds_to_charity ( -- ICosmicSignatureEvents.sol:FundsTransferredToCharity
	id              BIGSERIAL PRIMARY KEY,
	evtlog_id       BIGINT REFERENCES evt_log(id) ON DELETE CASCADE,
	block_num       BIGINT NOT NULL,
	tx_id           BIGINT NOT NULL,
	time_stamp      TIMESTAMPTZ NOT NULL,
	contract_aid    BIGINT NOT NULL,
	charity_aid		BIGINT NOT NULL,
	amount			DECIMAL NOT NULL,
	UNIQUE(evtlog_id)
);
CREATE TABLE cg_delay_duration ( -- ISystemManagement.sol:DelayDurationBeforeRoundActivationChanged
	id              BIGSERIAL PRIMARY KEY,
	evtlog_id       BIGINT REFERENCES evt_log(id) ON DELETE CASCADE,
	block_num       BIGINT NOT NULL,
	tx_id           BIGINT NOT NULL,
	time_stamp      TIMESTAMPTZ NOT NULL,
	contract_aid    BIGINT NOT NULL,
	new_value		DECIMAL NOT NULL,
	UNIQUE(evtlog_id)
);
CREATE TABLE cg_st_reward ( -- CST Staking rewards, per deposit, per token. This is the smallest reward unit (from which other accumulators are composed)
	-- This table is internal, it is populated via SQL triggers
	staker_aid		BIGINT NOT NULL,
	action_id		BIGINT NOT NULL,
	token_id		BIGINT NOT NULL,
	deposit_id		BIGINT NOT NULL,
	round_num		BIGINT NOT NULL,
	reward			DECIMAL NOT NULL,
	collected		BOOLEAN DEFAULT 'F',
	is_unstake		BOOLEAN DEFAULT 'F',	-- true if reward is generated on unstake() transaction
	UNIQUE(action_id,deposit_id)
);
CREATE TABLE cg_staker_cst ( -- counts statistics per user for staking CosmicSignature tokens
	staker_aid				BIGINT PRIMARY KEY,
	total_tokens_staked		BIGINT DEFAULT 0,
	num_stake_actions		BIGINT DEFAULT 0,
	num_unstake_actions		BIGINT DEFAULT 0,
	total_reward			DECIMAL DEFAULT 0,
	unclaimed_reward		DECIMAL DEFAULT 0,
	num_tokens_minted		BIGINT DEFAULT 0	-- this field is no longer used
);
CREATE TABLE cg_staker_deposit (-- accumulates rewards per staker (this is for CST staking wallet only)
	staker_aid				BIGINT NOT NULL,
	deposit_id				BIGINT NOT NULL, 
	tokens_staked			BIGINT DEFAULT 0,
	amount_deposited		DECIMAL DEFAULT 0,
	amount_to_claim			DECIMAL DEFAULT 0,
	PRIMARY KEY(staker_aid,deposit_id)
);
CREATE TABLE cg_staked_token_cst (	-- accumulates rewards per token (this table is NOT redundant with cg_st_reward (check count(*) on both)
	staker_aid				BIGINT NOT NULL,
	token_id				BIGINT NOT NULL,
	stake_action_id			BIGINT NOT NULL,
	PRIMARY KEY(token_id),
	UNIQUE(stake_action_id)
);
CREATE TABLE cg_staked_token_cst_rewards (-- Accumulates sum of rewards per action for all deposits
	staker_aid				BIGINT NOT NULL,
	token_id				BIGINT NOT NULL,
	stake_action_id			BIGINT NOT NULL,
	accumulated_reward		DECIMAL DEFAULT 0,	-- since staking can only be once per token, this table will keep the history forever
	claimed_reward			DECIMAL DEFAULT 0,	-- amount that has been claimed (can't be larger than accumulated_reward)
	PRIMARY KEY(token_id),
	UNIQUE(stake_action_id)
);
CREATE TABLE cg_staker_rwalk ( -- counts statistics per user for staking RandomWalk tokens
	staker_aid				BIGINT PRIMARY KEY,
	total_tokens_staked		BIGINT DEFAULT 0,
	num_stake_actions		BIGINT DEFAULT 0,
	num_unstake_actions		BIGINT DEFAULT 0,
	num_tokens_minted		BIGINT DEFAULT 0
);
CREATE TABLE cg_staked_token_rwalk (
	staker_aid				BIGINT NOT NULL,
	token_id				BIGINT NOT NULL,
	stake_action_id			BIGINT NOT NULL,
	PRIMARY KEY(token_id),
	UNIQUE(stake_action_id)
);
CREATE TABLE cg_contracts (
	cosmic_game_addr		TEXT NOT NULL,
	cosmic_signature_addr	TEXT NOT NULL,
	cosmic_token_addr		TEXT NOT NULL,
	cosmic_dao_addr			TEXT NOT NULL,
	charity_wallet_addr		TEXT NOT NULL,
	prizes_wallet_addr		TEXT NOT NULL,
	random_walk_addr		TEXT NOT NULL,
	staking_wallet_cst_addr		TEXT NOT NULL,
	staking_wallet_rwalk_addr	TEXT NOT NULL,
	marketing_wallet_addr	TEXT NOT NULL,
	implementation_addr		TEXT NOT NULL
);
CREATE TABLE cg_proc_status (
	last_evt_id             BIGINT DEFAULT 0,
	last_block_num          BIGINT DEFAULT 0
);

-- Game management tables (cg_adm_*): admin and config events.
-- Requires: evt_log. Load after base cosmicgame tables if split.

CREATE TABLE cg_adm_cst_min_limit ( -- StartingBidPriceCSTMinLimitChanged event
	id				BIGSERIAL PRIMARY KEY,
	evtlog_id		BIGINT REFERENCES evt_log(id) ON DELETE CASCADE,
	block_num		BIGINT NOT NULL,
	tx_id			BIGINT NOT NULL,
	time_stamp		TIMESTAMPTZ NOT NULL,
	contract_aid	BIGINT NOT NULL,
	min_limit		DECIMAL NOT NULL,
	UNIQUE(evtlog_id)
);
CREATE TABLE cg_adm_charity_pcent( -- ISystemEvents.sol:CharityEthDonationAmountPercentageChanged event
	id              BIGSERIAL PRIMARY KEY,
	evtlog_id       BIGINT REFERENCES evt_log(id) ON DELETE CASCADE,
	block_num       BIGINT NOT NULL,
	tx_id           BIGINT NOT NULL,
	time_stamp      TIMESTAMPTZ NOT NULL,
	contract_aid    BIGINT NOT NULL,
	percentage		DECIMAL NOT NULL,
	UNIQUE(evtlog_id)
);
CREATE TABLE cg_adm_main_prize_pcent( -- ISystemEvents.sol:MainEthPrizeAmountPercentageChanged
	id              BIGSERIAL PRIMARY KEY,
	evtlog_id       BIGINT REFERENCES evt_log(id) ON DELETE CASCADE,
	block_num       BIGINT NOT NULL,
	tx_id           BIGINT NOT NULL,
	time_stamp      TIMESTAMPTZ NOT NULL,
	contract_aid    BIGINT NOT NULL,
	percentage		DECIMAL NOT NULL,
	UNIQUE(evtlog_id)
);
CREATE TABLE cg_adm_stake_pcent( -- ISystemEvents.sol:StakingTotalEthRewardAmountPercentageChanged event
	id              BIGSERIAL PRIMARY KEY,
	evtlog_id       BIGINT REFERENCES evt_log(id) ON DELETE CASCADE,
	block_num       BIGINT NOT NULL,
	tx_id           BIGINT NOT NULL,
	time_stamp      TIMESTAMPTZ NOT NULL,
	contract_aid    BIGINT NOT NULL,
	percentage		DECIMAL NOT NULL,
	UNIQUE(evtlog_id)
);
CREATE TABLE cg_adm_raffle_pcent( -- ISystemEvents.sol:RaffleTotalEthPrizeAmountPercentageChanged event
	id              BIGSERIAL PRIMARY KEY,
	evtlog_id       BIGINT REFERENCES evt_log(id) ON DELETE CASCADE,
	block_num       BIGINT NOT NULL,
	tx_id           BIGINT NOT NULL,
	time_stamp      TIMESTAMPTZ NOT NULL,
	contract_aid    BIGINT NOT NULL,
	percentage		DECIMAL NOT NULL,
	UNIQUE(evtlog_id)
);
CREATE TABLE cg_adm_chrono_pcent( -- ISystemEvents.sol:ChronoWarriorEthPrizeAmountPercentageChanged event
	id              BIGSERIAL PRIMARY KEY,
	evtlog_id       BIGINT REFERENCES evt_log(id) ON DELETE CASCADE,
	block_num       BIGINT NOT NULL,
	tx_id           BIGINT NOT NULL,
	time_stamp      TIMESTAMPTZ NOT NULL,
	contract_aid    BIGINT NOT NULL,
	percentage		DECIMAL NOT NULL,
	UNIQUE(evtlog_id)
);
CREATE TABLE cg_adm_raf_eth_bidding( -- ISystemEvents.sol:NumRaffleEthPrizesForBiddersChanged event
	id              BIGSERIAL PRIMARY KEY,
	evtlog_id       BIGINT REFERENCES evt_log(id) ON DELETE CASCADE,
	block_num       BIGINT NOT NULL,
	tx_id           BIGINT NOT NULL,
	time_stamp      TIMESTAMPTZ NOT NULL,
	contract_aid    BIGINT NOT NULL,
	num_winners		DECIMAL NOT NULL,	-- newNumRaffleETHWinnersBidding
	UNIQUE(evtlog_id)
);
CREATE TABLE cg_adm_raf_nft_bidding( -- ISystemEvents.sol:NumRaffleCosmicSignatureNftsForBiddersChanged event
	id              BIGSERIAL PRIMARY KEY,
	evtlog_id       BIGINT REFERENCES evt_log(id) ON DELETE CASCADE,
	block_num       BIGINT NOT NULL,
	tx_id           BIGINT NOT NULL,
	time_stamp      TIMESTAMPTZ NOT NULL,
	contract_aid    BIGINT NOT NULL,
	num_winners		DECIMAL NOT NULL,	-- newNumRaffleNFTWinnersBidding 
	UNIQUE(evtlog_id)
);
CREATE TABLE cg_adm_raf_nft_staking_rwalk( -- ISystemEvents.sol:NumRaffleCosmicSignatureNftsForRandomWalkNftStakersChanged event
	id              BIGSERIAL PRIMARY KEY,
	evtlog_id       BIGINT REFERENCES evt_log(id) ON DELETE CASCADE,
	block_num       BIGINT NOT NULL,
	tx_id           BIGINT NOT NULL,
	time_stamp      TIMESTAMPTZ NOT NULL,
	contract_aid    BIGINT NOT NULL,
	num_winners		DECIMAL NOT NULL,	-- newNumRaffleNFTWinnersStakingRWalkChanged
	UNIQUE(evtlog_id)
);
CREATE TABLE cg_adm_charity_wallet( -- ISystemEvents.sol:CharityAddressChanged event (contract CosmicGame - renamed to CharityWalletChanged for clarity)
	id              BIGSERIAL PRIMARY KEY,
	evtlog_id       BIGINT REFERENCES evt_log(id) ON DELETE CASCADE,
	block_num       BIGINT NOT NULL,
	tx_id           BIGINT NOT NULL,
	time_stamp      TIMESTAMPTZ NOT NULL,
	contract_aid    BIGINT NOT NULL,
	new_charity_aid	BIGINT NOT NULL,
	UNIQUE(evtlog_id)
);
CREATE TABLE cg_adm_rwalk_addr( -- ISystemEvents.sol:RandomWalkAddressChanged event (contract CosmicGame)
	id              BIGSERIAL PRIMARY KEY,
	evtlog_id       BIGINT REFERENCES evt_log(id) ON DELETE CASCADE,
	block_num       BIGINT NOT NULL,
	tx_id           BIGINT NOT NULL,
	time_stamp      TIMESTAMPTZ NOT NULL,
	contract_aid    BIGINT NOT NULL,
	new_rwalk_aid	BIGINT NOT NULL,
	UNIQUE(evtlog_id)
);
CREATE TABLE cg_adm_prizes_wallet_addr( -- ISystemEvents.sol:PrizesWalletAddressChangedevent (contract CosmicGame)
	id              BIGSERIAL PRIMARY KEY,
	evtlog_id       BIGINT REFERENCES evt_log(id) ON DELETE CASCADE,
	block_num       BIGINT NOT NULL,
	tx_id           BIGINT NOT NULL,
	time_stamp      TIMESTAMPTZ NOT NULL,
	contract_aid    BIGINT NOT NULL,
	new_wallet_aid	BIGINT NOT NULL,
	UNIQUE(evtlog_id)
);
CREATE TABLE cg_adm_staking_cst_addr( -- ISystemEvents.sol:StakingWalletCosmicSignatureNftAddressChangedevent (contract CosmicGame)
	id              BIGSERIAL PRIMARY KEY,
	evtlog_id       BIGINT REFERENCES evt_log(id) ON DELETE CASCADE,
	block_num       BIGINT NOT NULL,
	tx_id           BIGINT NOT NULL,
	time_stamp      TIMESTAMPTZ NOT NULL,
	contract_aid    BIGINT NOT NULL,
	new_staking_aid	BIGINT NOT NULL,
	UNIQUE(evtlog_id)
);
CREATE TABLE cg_adm_staking_rwalk_addr( -- ISystemEVents.sol:StakingWalletRandomWalkNftAddressChanged(contract CosmicGame)
	id              BIGSERIAL PRIMARY KEY,
	evtlog_id       BIGINT REFERENCES evt_log(id) ON DELETE CASCADE,
	block_num       BIGINT NOT NULL,
	tx_id           BIGINT NOT NULL,
	time_stamp      TIMESTAMPTZ NOT NULL,
	contract_aid    BIGINT NOT NULL,
	new_staking_aid	BIGINT NOT NULL,
	UNIQUE(evtlog_id)
);
CREATE TABLE cg_adm_marketing_addr( -- ISystemEvents.sol:MarketingWalletAddressChanged event (contract CosmicGame)
	id              BIGSERIAL PRIMARY KEY,
	evtlog_id       BIGINT REFERENCES evt_log(id) ON DELETE CASCADE,
	block_num       BIGINT NOT NULL,
	tx_id           BIGINT NOT NULL,
	time_stamp      TIMESTAMPTZ NOT NULL,
	contract_aid    BIGINT NOT NULL,
	new_marketing_aid	BIGINT NOT NULL,
	UNIQUE(evtlog_id)
);
CREATE TABLE cg_adm_treasurer_addr( -- IMarketingWallet.sol:TreasurerAddressChanged
	id              BIGSERIAL PRIMARY KEY,
	evtlog_id       BIGINT REFERENCES evt_log(id) ON DELETE CASCADE,
	block_num       BIGINT NOT NULL,
	tx_id           BIGINT NOT NULL,
	time_stamp      TIMESTAMPTZ NOT NULL,
	contract_aid    BIGINT NOT NULL,
	new_treasurer_aid	BIGINT NOT NULL,
	UNIQUE(evtlog_id)
);
CREATE TABLE cg_adm_costok_addr( -- ISystemEvents.sol:CosmicTokenContractAddressChanged event (contract CosmicGame)
	id              BIGSERIAL PRIMARY KEY,
	evtlog_id       BIGINT REFERENCES evt_log(id) ON DELETE CASCADE,
	block_num       BIGINT NOT NULL,
	tx_id           BIGINT NOT NULL,
	time_stamp      TIMESTAMPTZ NOT NULL,
	contract_aid    BIGINT NOT NULL,
	new_costok_aid	BIGINT NOT NULL,
	UNIQUE(evtlog_id)
);
CREATE TABLE cg_adm_cossig_addr( -- ISystemEvents.sol:CosmicSignatureNftAddressChangedevent (contract CosmicGame)
	id              BIGSERIAL PRIMARY KEY,
	evtlog_id       BIGINT REFERENCES evt_log(id) ON DELETE CASCADE,
	block_num       BIGINT NOT NULL,
	tx_id           BIGINT NOT NULL,
	time_stamp      TIMESTAMPTZ NOT NULL,
	contract_aid    BIGINT NOT NULL,
	new_cossig_aid	BIGINT NOT NULL,
	UNIQUE(evtlog_id)
);
CREATE TABLE cg_adm_upgraded ( -- IERC1967.sol:Upgraded
	id              BIGSERIAL PRIMARY KEY,
	evtlog_id       BIGINT REFERENCES evt_log(id) ON DELETE CASCADE,
	block_num       BIGINT NOT NULL,
	tx_id           BIGINT NOT NULL,
	time_stamp      TIMESTAMPTZ NOT NULL,
	contract_aid    BIGINT NOT NULL,
	implementation_aid	BIGINT NOT NULL,
	UNIQUE(evtlog_id)
);
CREATE TABLE cg_adm_admin_changed ( -- IERC1967.sol:AdminChanged
	id              BIGSERIAL PRIMARY KEY,
	evtlog_id       BIGINT REFERENCES evt_log(id) ON DELETE CASCADE,
	block_num       BIGINT NOT NULL,
	tx_id           BIGINT NOT NULL,
	time_stamp      TIMESTAMPTZ NOT NULL,
	contract_aid    BIGINT NOT NULL,
	old_admin_aid	BIGINT NOT NULL,
	new_admin_aid	BIGINT NOT NULL,
	UNIQUE(evtlog_id)
);
CREATE TABLE cg_adm_time_inc( -- ISystemEvents.sol:TimeIncreaseChanged event (contract CosmicGame)
	id              BIGSERIAL PRIMARY KEY,
	evtlog_id       BIGINT REFERENCES evt_log(id) ON DELETE CASCADE,
	block_num       BIGINT NOT NULL,
	tx_id           BIGINT NOT NULL,
	time_stamp      TIMESTAMPTZ NOT NULL,
	contract_aid    BIGINT NOT NULL,
	new_time_inc	DECIMAL NOT NULL,
	UNIQUE(evtlog_id)
);
CREATE TABLE cg_adm_timeout_claimprize( -- ISystemEvents.sol:TimeoutDurationToClaimMainPrizeChanged event (contract CosmicGame)
	id              BIGSERIAL PRIMARY KEY,
	evtlog_id       BIGINT REFERENCES evt_log(id) ON DELETE CASCADE,
	block_num       BIGINT NOT NULL,
	tx_id           BIGINT NOT NULL,
	time_stamp      TIMESTAMPTZ NOT NULL,
	contract_aid    BIGINT NOT NULL,
	new_timeout		BIGINT NOT NULL,
	UNIQUE(evtlog_id)
);
CREATE TABLE cg_adm_timeout_withdraw( -- IPrizesWallet.sol:TimeoutDurationToWithdrawPrizesChanged event (contract CosmicGame)
	id              BIGSERIAL PRIMARY KEY,
	evtlog_id       BIGINT REFERENCES evt_log(id) ON DELETE CASCADE,
	block_num       BIGINT NOT NULL,
	tx_id           BIGINT NOT NULL,
	time_stamp      TIMESTAMPTZ NOT NULL,
	contract_aid    BIGINT NOT NULL,
	new_timeout		BIGINT NOT NULL,
	UNIQUE(evtlog_id)
);
CREATE TABLE cg_adm_price_inc( -- ISystemEvents.sol:PriceIncreaseChanged event (contract CosmicGame)
	id              BIGSERIAL PRIMARY KEY,
	evtlog_id       BIGINT REFERENCES evt_log(id) ON DELETE CASCADE,
	block_num       BIGINT NOT NULL,
	tx_id           BIGINT NOT NULL,
	time_stamp      TIMESTAMPTZ NOT NULL,
	contract_aid    BIGINT NOT NULL,
	new_price_increase	DECIMAL NOT NULL,
	UNIQUE(evtlog_id)
);
CREATE TABLE cg_adm_prize_microsec ( -- ISystemEvents.sol:MainPrizeTimeIncrementInMicroSecondsChanged event (contract CosmicGamez)
	id              BIGSERIAL PRIMARY KEY,
	evtlog_id       BIGINT REFERENCES evt_log(id) ON DELETE CASCADE,
	block_num       BIGINT NOT NULL,
	tx_id           BIGINT NOT NULL,
	time_stamp      TIMESTAMPTZ NOT NULL,
	contract_aid    BIGINT NOT NULL,
	new_microseconds	DECIMAL NOT NULL,
	UNIQUE(evtlog_id)
);
CREATE TABLE cg_adm_inisecprize ( -- ISystemEvents.sol:InitialSecondsUntilPrizeChanged event (contract CosmicGame)
	id              BIGSERIAL PRIMARY KEY,
	evtlog_id       BIGINT REFERENCES evt_log(id) ON DELETE CASCADE,
	block_num       BIGINT NOT NULL,
	tx_id           BIGINT NOT NULL,
	time_stamp      TIMESTAMPTZ NOT NULL,
	contract_aid    BIGINT NOT NULL,
	new_inisec		DECIMAL NOT NULL,
	UNIQUE(evtlog_id)
);
CREATE TABLE cg_adm_acttime ( -- ISystemEvents.sol:ActivationTimeChanged event (contract CosmicGame)
	id              BIGSERIAL PRIMARY KEY,
	evtlog_id       BIGINT REFERENCES evt_log(id) ON DELETE CASCADE,
	block_num       BIGINT NOT NULL,
	tx_id           BIGINT NOT NULL,
	time_stamp      TIMESTAMPTZ NOT NULL,
	contract_aid    BIGINT NOT NULL,
	new_atime		DECIMAL NOT NULL,
	UNIQUE(evtlog_id)
);
CREATE TABLE cg_adm_cst_auclen ( -- ISystemEvents.sol:CstDutchAuctionDurationDivisorChanged / ISystemEventsV2.sol:CstDutchAuctionDurationChanged
	-- Previously ISystemEvents.sol:RoundStartCSTAuctionLengthChanged event (contract CosmicGame)
	id              BIGSERIAL PRIMARY KEY,
	evtlog_id       BIGINT REFERENCES evt_log(id) ON DELETE CASCADE,
	block_num       BIGINT NOT NULL,
	tx_id           BIGINT NOT NULL,
	time_stamp      TIMESTAMPTZ NOT NULL,
	contract_aid    BIGINT NOT NULL,
	new_len			DECIMAL NOT NULL,
	UNIQUE(evtlog_id)
);
CREATE TABLE cg_adm_cst_auclen_chg_div ( -- ISystemEventsV2.sol:CstDutchAuctionDurationChangeDivisorChanged
	id              BIGSERIAL PRIMARY KEY,
	evtlog_id       BIGINT REFERENCES evt_log(id) ON DELETE CASCADE,
	block_num       BIGINT NOT NULL,
	tx_id           BIGINT NOT NULL,
	time_stamp      TIMESTAMPTZ NOT NULL,
	contract_aid    BIGINT NOT NULL,
	new_len			DECIMAL NOT NULL,
	UNIQUE(evtlog_id)
);
CREATE TABLE cg_adm_eth_auclen ( -- ISystemEvents.sol:EthDutchAuctionDurationDivisorChanged
	id              BIGSERIAL PRIMARY KEY,
	evtlog_id       BIGINT REFERENCES evt_log(id) ON DELETE CASCADE,
	block_num       BIGINT NOT NULL,
	tx_id           BIGINT NOT NULL,
	time_stamp      TIMESTAMPTZ NOT NULL,
	contract_aid    BIGINT NOT NULL,
	new_len			DECIMAL NOT NULL,
	UNIQUE(evtlog_id)
);
CREATE TABLE cg_adm_eth_auc_endprice ( -- ISystemEvents.sol:EthDutchAuctionEndingBidPriceDivisorChanged
	id              BIGSERIAL PRIMARY KEY,
	evtlog_id       BIGINT REFERENCES evt_log(id) ON DELETE CASCADE,
	block_num       BIGINT NOT NULL,
	tx_id           BIGINT NOT NULL,
	time_stamp      TIMESTAMPTZ NOT NULL,
	contract_aid    BIGINT NOT NULL,
	new_len			DECIMAL NOT NULL,
	UNIQUE(evtlog_id)
);
CREATE TABLE cg_adm_erc_rwd_mul ( -- ISystemEvents.sol:CstPrizeAmountChanged
	id              BIGSERIAL PRIMARY KEY,
	evtlog_id       BIGINT REFERENCES evt_log(id) ON DELETE CASCADE,
	block_num       BIGINT NOT NULL,
	tx_id           BIGINT NOT NULL,
	time_stamp      TIMESTAMPTZ NOT NULL,
	contract_aid    BIGINT NOT NULL,
	new_reward		DECIMAL NOT NULL,	-- static CST prize amount (no longer multiplier)
	UNIQUE(evtlog_id)
);
CREATE TABLE cg_adm_mkt_reward ( -- ISystemEvents.sol:MarketingWalletCstContributionAmountChanged
	id              BIGSERIAL PRIMARY KEY,
	evtlog_id       BIGINT REFERENCES evt_log(id) ON DELETE CASCADE,
	block_num       BIGINT NOT NULL,
	tx_id           BIGINT NOT NULL,
	time_stamp      TIMESTAMPTZ NOT NULL,
	contract_aid    BIGINT NOT NULL,
	new_reward		DECIMAL NOT NULL,
	UNIQUE(evtlog_id)
);
CREATE TABLE cg_adm_erc20_reward ( -- ISystemEvents.sol:CstRewardAmountForBiddingChanged / BidCstRewardAmountChanged / ISystemEventsV2.sol:BidCstRewardAmountMultiplierChanged
	id              BIGSERIAL PRIMARY KEY,
	evtlog_id       BIGINT REFERENCES evt_log(id) ON DELETE CASCADE,
	block_num       BIGINT NOT NULL,
	tx_id           BIGINT NOT NULL,
	time_stamp      TIMESTAMPTZ NOT NULL,
	contract_aid    BIGINT NOT NULL,
	new_reward		DECIMAL NOT NULL,
	UNIQUE(evtlog_id)
);
CREATE TABLE cg_adm_msg_len ( -- ISystemEvents.sol:BidMessageLengthMaxLimitChanged
	id              BIGSERIAL PRIMARY KEY,
	evtlog_id       BIGINT REFERENCES evt_log(id) ON DELETE CASCADE,
	block_num       BIGINT NOT NULL,
	tx_id           BIGINT NOT NULL,
	time_stamp      TIMESTAMPTZ NOT NULL,
	contract_aid    BIGINT NOT NULL,
	new_length		DECIMAL NOT NULL,
	UNIQUE(evtlog_id)
);
CREATE TABLE cg_adm_script_url ( -- CosmicSignatureNft.sol:TokenGenerationScriptURLEvent(admin event)
	id              BIGSERIAL PRIMARY KEY,
	evtlog_id       BIGINT REFERENCES evt_log(id) ON DELETE CASCADE,
	block_num       BIGINT NOT NULL,
	tx_id           BIGINT NOT NULL,
	time_stamp      TIMESTAMPTZ NOT NULL,
	contract_aid    BIGINT NOT NULL,
	new_url			TEXT NOT NULL,
	UNIQUE(evtlog_id)
);
CREATE TABLE cg_adm_base_uri_cs( -- CosmicSignatureNft.sol:BaseURI 
	id              BIGSERIAL PRIMARY KEY,
	evtlog_id       BIGINT REFERENCES evt_log(id) ON DELETE CASCADE,
	block_num       BIGINT NOT NULL,
	tx_id           BIGINT NOT NULL,
	time_stamp      TIMESTAMPTZ NOT NULL,
	contract_aid    BIGINT NOT NULL,
	new_uri			TEXT NOT NULL,
	UNIQUE(evtlog_id)
);
CREATE TABLE cg_adm_ownership ( -- OwnershipTransferred event (OpenZeppelin) (not to confuse with ERC721 token owner)
	id              BIGSERIAL PRIMARY KEY,
	evtlog_id       BIGINT REFERENCES evt_log(id) ON DELETE CASCADE,
	block_num       BIGINT NOT NULL,
	tx_id           BIGINT NOT NULL,
	time_stamp      TIMESTAMPTZ NOT NULL,
	contract_aid    BIGINT NOT NULL,
	prev_owner_aid	BIGINT NOT NULL,
	new_owner_aid	BIGINT NOT NULL,
	contract_code	INT NOT NULL,		-- 1: CosmicGame, 2: CosmicSignature, 3: CosmicToken, 3: CharityWallet, 4: EthPrizesWallet, 5: StakingWallet CST, 6: StakingWallet RWalk
	UNIQUE(evtlog_id)
);
CREATE TABLE cg_adm_initialized( -- Initialized event (OpenZeppelin)
	id              BIGSERIAL PRIMARY KEY,
	evtlog_id       BIGINT REFERENCES evt_log(id) ON DELETE CASCADE,
	block_num       BIGINT NOT NULL,
	tx_id           BIGINT NOT NULL,
	time_stamp      TIMESTAMPTZ NOT NULL,
	contract_aid    BIGINT NOT NULL,
	version			BIGINT NOT NULL,
	UNIQUE(evtlog_id)
);

-- Banned bids (admin/manual bans; API: get_banned_bids, ban_bid, unban_bid)
CREATE TABLE cg_banned_bids (
	id          BIGSERIAL PRIMARY KEY,
	bid_id      BIGINT NOT NULL,
	user_addr   VARCHAR(255) NOT NULL,
	created_at  BIGINT NOT NULL
);
CREATE INDEX idx_cg_banned_bids_bid_id ON cg_banned_bids(bid_id);

-- Statistical and aggregate tables for Cosmic Game.
-- Load after base cosmicgame tables (and game-management.sql if used).

CREATE TABLE cg_transfer_stats( -- table to keep tracking of the statistical counters for tokent transfers
    user_aid                BIGINT NOT NULL,
    erc20_num_transfers     BIGINT DEFAULT 0, -- CosmicToken
    erc721_num_transfers    BIGINT DEFAULT 0  -- CosmicSignature
);
CREATE TABLE cg_round_stats( -- collects statistics per round 
	round_num					BIGINT NOT NULL PRIMARY KEY,
	total_bids					BIGINT DEFAULT 0,
	total_nft_donated			BIGINT DEFAULT 0,
	num_erc20_donations			BIGINT DEFAULT 0,		-- number of donations made during the round (ERC20 tokens)
	total_raffle_eth_deposits	DECIMAL DEFAULT 0,
	chrono_warrior_prize_eth	DECIMAL DEFAULT 0,
	total_cst_paid_in_prizes	DECIMAL DEFAULT 0,
	total_nfts_minted			BIGINT DEFAULT 0,
	num_contracts_donated_erc20	BIGINT DEFAULT 0,		-- number of unique ERC20 token contracts that donated in this round
	total_raffle_nfts			BIGINT DEFAULT 0,		-- counts only raffle NFTs
	donations_round_total		DECIMAL DEFAULT 0,		-- total donations for current round (reset on claimPrize())
	donations_round_count		BIGINT DEFAULT 0,		-- total number of donations for the current round
	total_eth_in_bids			DECIMAL DEFAULT 0,		-- sum of ETH in all bids
	total_cst_in_bids			DECIMAL DEFAULT 0,		-- sum of CST in all bids
	-- Round timing fields (added 2025-11-06)
	param_window_start_time		TIMESTAMPTZ,			-- When parameter setting window starts (previous round ends)
	activation_time				TIMESTAMPTZ,			-- When admin sets round activation (param window ends)
	param_window_duration_seconds BIGINT,				-- Duration of parameter setting window
	round_start_time			TIMESTAMPTZ,			-- When FirstBidPlacedInRound fires (actual round start)
	round_end_time				TIMESTAMPTZ,			-- When prize is claimed (round ends)
	round_duration_seconds		BIGINT					-- Duration of active round (end - start)
);
CREATE TABLE cg_bidder ( -- collects statistics per bidder
	bidder_aid		BIGINT PRIMARY KEY,
	num_bids		BIGINT DEFAULT 0,
	max_bid			DECIMAL DEFAULT 0,
	tokens_minted	DECIMAL DEFAULT 0, -- total tokens minted
	total_eth_spent	DECIMAL DEFAULT 0, -- sum of eth_price across this bidder's ETH/RandomWalk bids (wei)
	total_cst_spent	DECIMAL DEFAULT 0  -- sum of cst_price across this bidder's CST bids (wei)
);
CREATE TABLE cg_winner ( -- collects statistics per winner of any prize type
	winner_aid				BIGINT PRIMARY KEY,
	max_win_amount			DECIMAL DEFAULT 0,	-- max ETH won in main prize
	prizes_count			BIGINT DEFAULT 0,	-- total prize count (all types)
	prizes_sum				DECIMAL DEFAULT 0,	-- sum of ETH won (main + raffle + chrono warrior)
	tokens_count			BIGINT DEFAULT 0,	-- DEPRECATED: use erc721_count instead
	erc20_count				BIGINT DEFAULT 0,	-- count of ERC20 (CST) prizes won
	erc721_count			BIGINT DEFAULT 0,	-- count of ERC721 (NFT) prizes won
	unclaimed_nfts			BIGINT DEFAULT 0	-- donated NFTs (erc721) pending claim by winner
);
CREATE TABLE cg_donor (--counts statistics for unique donors (who donate ETH to cosmic game)
	donor_aid				BIGINT PRIMARY KEY,
	count_donations			BIGINT DEFAULT 0,
	total_eth_donated		DECIMAL DEFAULT 0
);
CREATE TABLE cg_raffle_winner_stats (	-- prizes in ETH
	winner_aid		BIGINT PRIMARY KEY,
	amount_sum		DECIMAL DEFAULT 0,
	withdrawal_sum	DECIMAL DEFAULT 0,
	raffles_count	BIGINT DEFAULT 0
);
CREATE TABLE cg_raffle_nft_winner_stats ( -- prizes in NFT
	winner_aid		BIGINT PRIMARY KEY,
	num_won			BIGINT DEFAULT 0	-- num tokens won
);
CREATE TABLE cg_glob_stats ( -- global statistics
	num_vol_donations		BIGINT DEFAULT 0,		-- total number of voluntary donations
	vol_donations_total		DECIMAL DEFAULT 0,		-- sum of voluntary donations
	cg_donations_total		DECIMAL DEFAULT 0,		-- sum of all donatinos deposited by CosmicGame contract
	num_cg_donations		BIGINT DEFAULT 0,		-- number of donations deposited by CosmicGame contract
	direct_donations		DECIMAL DEFAULT 0,		-- direct donations to CosmicGame contract via donate() function
	num_direct_donations	BIGINT DEFAULT 0,		-- number of direct donationons
	sum_withdrawals			DECIMAL DEFAULT 0,		-- sum of withdrawals from CharityWallet to recipient
	num_withdrawals			BIGINT DEFAULT 0,
	num_bids				BIGINT DEFAULT 0, 		-- total bids made
	num_wins				BIGINT DEFAULT 0,		-- total prizes given
	num_rwalk_used			BIGINT DEFAULT 0,
	num_mints				BIGINT DEFAULT 0,
	cur_num_bids			BIGINT DEFAULT 0,		-- num bids since new round
	num_bids_cst			BIGINT DEFAULT 0,		-- amount of bids made with CST
	total_raffle_eth_deposits DECIMAL DEFAULT 0,
	total_raffle_eth_withdrawn DECIMAL DEFAULT 0,
	total_chrono_warrior_eth_deposits DECIMAL DEFAULT 0,
	total_cst_given_in_prizes DECIMAL DEFAULT 0,
	cst_reward_for_bidding DECIMAL DEFAULT 0, -- V1: fixed CST reward wei; post-V2 upgrade stores bidCstRewardAmountMultiplier
	total_nft_donated		BIGINT DEFAULT 0,
	total_erc20_donations	BIGINT DEFAULT 0,		-- the number of donations, not the number of tokens
	total_cst_consumed		DECIMAL DEFAULT 0,		-- or burned, sum of the tokens that was burned as bid price
	total_mkt_rewards		DECIMAL DEFAULT 0,
	num_mkt_rewards			BIGINT DEFAULT 0
);
CREATE TABLE cg_nft_stats ( -- stats for donated NFTs (donated with bidAndDonateNFT())
	contract_aid			BIGINT PRIMARY KEY,
	num_donated				BIGINT DEFAULT 0		-- how many NFTs were donated
);
CREATE TABLE cg_erc20_donation_stats ( -- stats for donated NFTs (donated with bidAndDonateNFT())
	token_aid				BIGINT NOT NULL,
	round_num				BIGINT NOT NULL,
	total_amount			DECIMAL DEFAULT 0,		-- the sum for all donations for the round on single ERC200 token
	claimed					BOOLEAN DEFAULT 'F',
	winner_aid				BIGINT DEFAULT 0,		-- stored winner_aid when the guy actually makes the claim
	PRIMARY KEY(round_num,token_aid)
);
CREATE TABLE cg_stake_stats_cst ( -- gloal staking statistics (StakinWalletCST)
	total_tokens_staked		BIGINT DEFAULT 0,
	total_reward_amount		DECIMAL DEFAULT 0,
	total_unclaimed_reward	DECIMAL DEFAULT 0,
	total_num_stakers		BIGINT DEFAULT 0,
	num_deposits			BIGINT DEFAULT 0,
	total_modulo			DECIMAL DEFAULT 0,
	num_charity_deposits	BIGINT DEFAULT 0,
	total_charity_amount	DECIMAL DEFAULT 0
);
CREATE TABLE cg_stake_stats_rwalk ( -- gloal staking statistics (StakinWalletRWalk)
	total_tokens_staked		BIGINT DEFAULT 0,
	total_num_stakers		BIGINT DEFAULT 0,
	total_nft_mints			BIGINT DEFAULT 0
);

INSERT INTO cg_glob_stats DEFAULT VALUES;
INSERT INTO cg_stake_stats_cst DEFAULT VALUES;
INSERT INTO cg_stake_stats_rwalk DEFAULT VALUES;

-- +goose StatementBegin
CREATE OR REPLACE FUNCTION on_bid_insert() RETURNS trigger AS  $$
DECLARE
	v_max_bid				DECIMAL;
	v_cnt					NUMERIC;
BEGIN

	SELECT MAX(eth_price) FROM cg_bid INTO v_max_bid WHERE bidder_aid = NEW.bidder_aid;
	IF v_max_bid IS NULL THEN
		v_max_bid := 0;
	END IF;
	UPDATE cg_bidder
		SET
			num_bids = (num_bids + 1),
			max_bid	 = v_max_bid,
			total_eth_spent = total_eth_spent + (CASE WHEN NEW.eth_price > 0 THEN NEW.eth_price ELSE 0 END),
			total_cst_spent = total_cst_spent + (CASE WHEN NEW.cst_price > 0 THEN NEW.cst_price ELSE 0 END)
		WHERE bidder_aid = NEW.bidder_aid;
	GET DIAGNOSTICS v_cnt = ROW_COUNT;
	IF v_cnt = 0 THEN
		INSERT INTO cg_bidder(bidder_aid,num_bids,max_bid,total_eth_spent,total_cst_spent)
			VALUES(
				NEW.bidder_aid,
				1,
				v_max_bid,
				(CASE WHEN NEW.eth_price > 0 THEN NEW.eth_price ELSE 0 END),
				(CASE WHEN NEW.cst_price > 0 THEN NEW.cst_price ELSE 0 END)
			);
	END IF;
	UPDATE cg_glob_stats SET num_bids = (num_bids + 1);
	GET DIAGNOSTICS v_cnt = ROW_COUNT;
	IF v_cnt = 0 THEN
		RAISE EXCEPTION 'cg_glob_stats table wasnt initialized (no record found)';
	END IF;
	IF NEW.rwalk_nft_id > -1 THEN
		UPDATE cg_glob_stats SET num_rwalk_used = (num_rwalk_used + 1);
	ELSE 
		IF NEW.bid_type = 2 THEN
			UPDATE cg_glob_stats SET
				num_bids_cst = (num_bids_cst + 1),
				total_cst_consumed = (total_cst_consumed + NEW.cst_price);
		END IF;
	END IF;
	UPDATE cg_glob_stats SET cur_num_bids = (cur_num_bids + 1);
	UPDATE cg_round_stats SET 
			total_bids = (total_bids + 1),
			total_cst_in_bids = (total_cst_in_bids + CASE WHEN NEW.bid_type = 2 THEN NEW.cst_price ELSE 0 END),
			total_eth_in_bids = (total_eth_in_bids + NEW.eth_price)
	   	WHERE round_num=NEW.round_num;
	GET DIAGNOSTICS v_cnt = ROW_COUNT;
	IF v_cnt = 0 THEN
		INSERT INTO cg_round_stats(round_num,total_bids) VALUES (NEW.round_num,1);
	END IF;
	RETURN NEW;
END;
$$ LANGUAGE plpgsql;
-- +goose StatementEnd
-- +goose StatementBegin
CREATE OR REPLACE FUNCTION on_bid_delete() RETURNS trigger AS  $$
DECLARE
	v_max_bid				DECIMAL;
	v_cnt					NUMERIC;
BEGIN

	SELECT MAX(eth_price) FROM cg_bid INTO v_max_bid WHERE bidder_aid = OLD.bidder_aid;
	IF v_max_bid IS NULL THEN
		v_max_bid := 0;
	END IF;
	UPDATE cg_bidder
		SET
			num_bids = (num_bids - 1),
			max_bid	 = v_max_bid,
			total_eth_spent = total_eth_spent - (CASE WHEN OLD.eth_price > 0 THEN OLD.eth_price ELSE 0 END),
			total_cst_spent = total_cst_spent - (CASE WHEN OLD.cst_price > 0 THEN OLD.cst_price ELSE 0 END)
		WHERE bidder_aid = OLD.bidder_aid;
	UPDATE cg_glob_stats SET num_bids = (num_bids - 1);
	GET DIAGNOSTICS v_cnt = ROW_COUNT;
	IF v_cnt = 0 THEN
		RAISE EXCEPTION 'cg_glob_stats table wasnt initialized (no record found)';
	END IF;
	IF OLD.rwalk_nft_id > -1 THEN
		UPDATE cg_glob_stats SET num_rwalk_used = (num_rwalk_used - 1);
	ELSE
		IF OLD.bid_type = 2 THEN
			UPDATE cg_glob_stats SET 
				num_bids_cst = (num_bids_cst - 1),
				total_cst_consumed = (total_cst_consumed - OLD.cst_price);
		END IF;
	END IF;
	UPDATE cg_glob_stats SET cur_num_bids = (cur_num_bids - 1) WHERE cur_num_bids>0;
	UPDATE cg_round_stats SET 
			total_bids = (total_bids - 1),
			total_cst_in_bids = (total_cst_in_bids - CASE WHEN OLD.bid_type = 2 THEN OLD.cst_price ELSE 0 END),
			total_eth_in_bids = (total_eth_in_bids - OLD.eth_price)
		WHERE round_num=OLD.round_num;
	RETURN OLD;
END;
$$ LANGUAGE plpgsql;
-- +goose StatementEnd
-- +goose StatementBegin
CREATE OR REPLACE FUNCTION on_prize_claim_insert() RETURNS trigger AS  $$
DECLARE
	v_max_prize				DECIMAL;
	v_prizes_sum			DECIMAL;
	v_prizes_count			BIGINT;
	v_donated_nfts			BIGINT;
	v_cnt					NUMERIC;
BEGIN

	-- Main Prize adds 3 prizes (ETH, CST, NFT)
	v_prizes_count := 3;
	
	-- Get max prize amount and sum from Main Prize only (for backward compatibility)
	SELECT MAX(amount), SUM(amount)
		FROM cg_prize_claim
		WHERE winner_aid=NEW.winner_aid
		INTO v_max_prize, v_prizes_sum;
	IF v_max_prize IS NULL THEN
		v_max_prize := 0;
		v_prizes_sum := 0;
	END IF;
	IF v_prizes_count IS NULL THEN
		v_prizes_count := 0;
	END IF;
	SELECT total_nft_donated FROM cg_round_stats WHERE round_num=NEW.round_num INTO v_donated_nfts;
	IF v_donated_nfts IS NULL THEN
		v_donated_nfts := 0;
	END IF;
	UPDATE cg_winner
		SET
			prizes_count = (prizes_count + v_prizes_count),
			max_win_amount	 = v_max_prize,
			prizes_sum = v_prizes_sum,
			unclaimed_nfts = (unclaimed_nfts + v_donated_nfts),
			erc20_count = (erc20_count + 1),   -- Main prize awards 1 CST prize
			erc721_count = (erc721_count + 1)  -- Main prize awards 1 NFT
		WHERE winner_aid = NEW.winner_aid;
	GET DIAGNOSTICS v_cnt = ROW_COUNT;
	IF v_cnt = 0 THEN
		INSERT INTO cg_winner(winner_aid,max_win_amount,prizes_count,prizes_sum,unclaimed_nfts,erc20_count,erc721_count)
			VALUES(NEW.winner_aid,v_max_prize,v_prizes_count,v_prizes_sum,v_donated_nfts,1,1);
	END IF;
	UPDATE cg_glob_stats SET num_wins = (num_wins + 1);
	GET DIAGNOSTICS v_cnt = ROW_COUNT;
	IF v_cnt = 0 THEN
		RAISE EXCEPTION 'cg_glob_stats table wasnt initialized (no record found)';
	END IF;
	UPDATE cg_glob_stats SET cur_num_bids = 0;
	GET DIAGNOSTICS v_cnt = ROW_COUNT;
	IF v_cnt = 0 THEN
		RAISE EXCEPTION 'on_prize_claim_insert() failed to update cur_num_bids in cg_glob_stats';
	END IF;
	
	UPDATE cg_erc20_donation_stats SET winner_aid=NEW.winner_aid WHERE round_num=NEW.round_num;
	-- Note: This UPDATE may affect 0 rows if there are no ERC20 donations for this round (OK)
	
	UPDATE cg_glob_stats SET total_cst_given_in_prizes = (total_cst_given_in_prizes + NEW.cst_amount);
	GET DIAGNOSTICS v_cnt = ROW_COUNT;
	IF v_cnt = 0 THEN
		RAISE EXCEPTION 'on_prize_claim_insert() failed to update total_cst_given_in_prizes in cg_glob_stats';
	END IF;
	
	-- Update round stats (create round record if it doesn't exist)
	UPDATE cg_round_stats SET total_cst_paid_in_prizes = (total_cst_paid_in_prizes + NEW.cst_amount), total_nfts_minted = (total_nfts_minted + 1) WHERE round_num = NEW.round_num;
	GET DIAGNOSTICS v_cnt = ROW_COUNT;
	IF v_cnt = 0 THEN
		INSERT INTO cg_round_stats(round_num, total_cst_paid_in_prizes, total_nfts_minted) VALUES (NEW.round_num, NEW.cst_amount, 1);
		RAISE NOTICE 'on_prize_claim_insert() created new round_stats record for round_num=%', NEW.round_num;
	END IF;
	
	-- Insert THREE records in cg_prize table for Main Prize
	-- 1) Main Prize ETH (ptype=0)
	INSERT INTO cg_prize(round_num,winner_index,ptype) VALUES(NEW.round_num,0,0);
	-- 2) Main Prize CST (ptype=1)
	INSERT INTO cg_prize(round_num,winner_index,ptype) VALUES(NEW.round_num,0,1);
	-- 3) Main Prize CS NFT (ptype=2)
	INSERT INTO cg_prize(round_num,winner_index,ptype) VALUES(NEW.round_num,0,2);
	
	RETURN NEW;
END;
$$ LANGUAGE plpgsql;
-- +goose StatementEnd
-- +goose StatementBegin
CREATE OR REPLACE FUNCTION on_prize_claim_delete() RETURNS trigger AS  $$
DECLARE
	v_max_prize				DECIMAL;
	v_prizes_sum			DECIMAL;
	v_prizes_count			BIGINT;
	v_cnt					NUMERIC;
	v_donated_nfts			BIGINT;
BEGIN

	SELECT
			MAX(amount),
 			COUNT(*) as prizes_count,
			SUM(amount) as prizes_sum
		FROM cg_prize_claim
		WHERE winner_aid=OLD.winner_aid
		INTO v_max_prize,v_prizes_count,v_prizes_sum;
	IF v_max_prize IS NULL THEN
		v_max_prize := 0;
		v_prizes_sum := 0;
		v_prizes_count := 0;
	END IF;
	SELECT total_nft_donated FROM cg_round_stats WHERE round_num=OLD.round_num INTO v_donated_nfts;
	IF v_donated_nfts IS NULL THEN
		v_donated_nfts := 0;
	END IF;
	UPDATE cg_winner
		SET
			prizes_count = v_prizes_count,
			max_win_amount	 = v_max_prize,
			prizes_sum = v_prizes_sum,
			unclaimed_nfts = GREATEST(0, unclaimed_nfts - v_donated_nfts),
			erc20_count = GREATEST(0, erc20_count - 1),
			erc721_count = GREATEST(0, erc721_count - 1)
		WHERE winner_aid = OLD.winner_aid;

	UPDATE cg_glob_stats SET num_wins = (num_wins - 1);
	GET DIAGNOSTICS v_cnt = ROW_COUNT;
	IF v_cnt = 0 THEN
		RAISE EXCEPTION 'cg_glob_stats table wasnt initialized (no record found)';
	END IF;
	UPDATE cg_erc20_donation_stats SET winner_aid=0 WHERE round_num=OLD.round_num;
	UPDATE cg_glob_stats SET total_cst_given_in_prizes = (total_cst_given_in_prizes - OLD.cst_amount);
	
	-- Update round stats
	UPDATE cg_round_stats SET total_cst_paid_in_prizes = (total_cst_paid_in_prizes - OLD.cst_amount) WHERE round_num = OLD.round_num;
	UPDATE cg_round_stats SET total_nfts_minted = (total_nfts_minted - 1) WHERE round_num = OLD.round_num;
	
	-- Remove THREE corresponding records from cg_prize table
	-- 1) Main Prize ETH (ptype=0)
	DELETE FROM cg_prize WHERE round_num=OLD.round_num AND winner_index=0 AND ptype=0;
	-- 2) Main Prize CST (ptype=1)
	DELETE FROM cg_prize WHERE round_num=OLD.round_num AND winner_index=0 AND ptype=1;
	-- 3) Main Prize CS NFT (ptype=2)
	DELETE FROM cg_prize WHERE round_num=OLD.round_num AND winner_index=0 AND ptype=2;
	
	RETURN OLD;
END;
$$ LANGUAGE plpgsql;
-- +goose StatementEnd
-- +goose StatementBegin
CREATE OR REPLACE FUNCTION on_donation_received_insert() RETURNS trigger AS  $$
DECLARE
	v_cnt						NUMERIC;
	v_cosmicgame_aid			BIGINT;
	v_cosmicgame_addr			TEXT;
BEGIN

	SELECT cosmic_game_addr FROM cg_contracts LIMIT 1 INTO v_cosmicgame_addr;
	IF v_cosmicgame_addr IS NULL THEN
		RAISE EXCEPTION 'CosmicGame contract address is not defined';
	END IF;
	SELECT address_id FROM address WHERE addr=v_cosmicgame_addr INTO v_cosmicgame_aid;
	IF v_cosmicgame_aid IS NULL THEN
		RAISE EXCEPTION 'CosmicGame address id not found in address table';
	END IF;
	IF NEW.donor_aid != v_cosmicgame_aid THEN
		UPDATE cg_glob_stats 
			SET 
				num_vol_donations = (num_vol_donations + 1),
				vol_donations_total = (vol_donations_total + NEW.amount);
	ELSE 
		UPDATE cg_glob_stats 
			SET 
				num_cg_donations = (num_cg_donations + 1),
				cg_donations_total = (cg_donations_total + NEW.amount);
		
	END IF;
	GET DIAGNOSTICS v_cnt = ROW_COUNT;
	IF v_cnt = 0 THEN
		RAISE EXCEPTION 'cg_glob_stats table wasnt initialized (no record found)';
	END IF;

	RETURN NEW;
END;
$$ LANGUAGE plpgsql;
-- +goose StatementEnd
-- +goose StatementBegin
CREATE OR REPLACE FUNCTION on_donation_received_delete() RETURNS trigger AS  $$
DECLARE
	v_cnt					NUMERIC;
	v_cosmicgame_aid			BIGINT;
	v_cosmicgame_addr			TEXT;
BEGIN

	SELECT cosmic_game_addr FROM cg_contracts LIMIT 1 INTO v_cosmicgame_addr;
	IF v_cosmicgame_addr IS NULL THEN
		RAISE EXCEPTION 'CosmicGame contract address is not defined';
	END IF;
	SELECT address_id FROM address WHERE addr=v_cosmicgame_addr INTO v_cosmicgame_aid;
	IF v_cosmicgame_aid IS NULL THEN
		RAISE EXCEPTION 'CosmicGame address id not found in address table';
	END IF;
	IF OLD.donor_aid != v_cosmicgame_aid THEN
		UPDATE cg_glob_stats
			SET
				num_vol_donations = (num_vol_donations - 1),
				vol_donations_total = (vol_donations_total - OLD.amount);
	ELSE
		UPDATE cg_glob_stats
			SET
				num_cg_donations = (num_cg_donations - 1),
				cg_donations_total = (cg_donations_total - OLD.amount);
	END IF;
	GET DIAGNOSTICS v_cnt = ROW_COUNT;
	IF v_cnt = 0 THEN
		RAISE EXCEPTION 'cg_glob_stats table wasnt initialized (no record found)';
	END IF;
	RETURN OLD;
END;
$$ LANGUAGE plpgsql;
-- +goose StatementEnd
-- +goose StatementBegin
CREATE OR REPLACE FUNCTION on_erc20_donation_insert() RETURNS trigger AS  $$
DECLARE
	v_cnt					NUMERIC;
BEGIN

	UPDATE cg_erc20_donation_stats SET total_amount = (total_amount + NEW.amount) WHERE round_num=NEW.round_num AND token_aid=NEW.token_aid;
	GET DIAGNOSTICS v_cnt = ROW_COUNT;
	IF v_cnt = 0 THEN
		-- New token contract for this round - increment unique contract count
		INSERT INTO cg_erc20_donation_stats(token_aid,round_num,total_amount) VALUES (NEW.token_aid,NEW.round_num,NEW.amount);
		UPDATE cg_round_stats SET num_contracts_donated_erc20 = (num_contracts_donated_erc20 + 1) WHERE round_num=NEW.round_num;
	END IF;
	UPDATE cg_round_stats SET num_erc20_donations = (num_erc20_donations + 1) WHERE round_num=NEW.round_num;
	GET DIAGNOSTICS v_cnt = ROW_COUNT;
	IF v_cnt = 0 THEN
		INSERT INTO cg_round_stats(round_num,num_erc20_donations) VALUES (NEW.round_num,1);
	END IF;
	UPDATE cg_glob_stats SET total_erc20_donations = (total_erc20_donations + 1);
	RETURN NEW;
END;
$$ LANGUAGE plpgsql;
-- +goose StatementEnd
-- +goose StatementBegin
CREATE OR REPLACE FUNCTION on_erc20_donation_delete() RETURNS trigger AS  $$
DECLARE
	v_cnt					NUMERIC;
	v_remaining_amount		DECIMAL;
BEGIN

	UPDATE cg_erc20_donation_stats SET total_amount = (total_amount - OLD.amount) WHERE round_num=OLD.round_num AND token_aid=OLD.token_aid RETURNING total_amount INTO v_remaining_amount;
	
	-- If this was the last donation for this token contract in this round, decrement unique contract count
	IF v_remaining_amount = 0 THEN
		DELETE FROM cg_erc20_donation_stats WHERE round_num=OLD.round_num AND token_aid=OLD.token_aid;
		UPDATE cg_round_stats SET num_contracts_donated_erc20 = (num_contracts_donated_erc20 - 1) WHERE round_num=OLD.round_num;
	END IF;
	
	UPDATE cg_round_stats SET num_erc20_donations = (num_erc20_donations - 1) WHERE round_num=OLD.round_num;
	UPDATE cg_glob_stats SET total_erc20_donations = (total_erc20_donations - 1);
	RETURN OLD;
END;
$$ LANGUAGE plpgsql;
-- +goose StatementEnd
-- +goose StatementBegin
CREATE OR REPLACE FUNCTION on_nft_donation_insert() RETURNS trigger AS  $$
DECLARE
	v_cnt					NUMERIC;
BEGIN

	UPDATE cg_nft_stats
		SET
			num_donated = (num_donated + 1)
		WHERE contract_aid = NEW.token_aid;
	GET DIAGNOSTICS v_cnt = ROW_COUNT;
	IF v_cnt = 0 THEN
		INSERT INTO cg_nft_stats(contract_aid,num_donated)
			VALUES(NEW.token_aid,1);
	END IF;
	UPDATE cg_round_stats SET total_nft_donated = (total_nft_donated + 1) WHERE round_num=NEW.round_num;
	GET DIAGNOSTICS v_cnt = ROW_COUNT;
	IF v_cnt = 0 THEN
		INSERT INTO cg_round_stats(round_num,total_nft_donated) VALUES (NEW.round_num,1);
	END IF;
	UPDATE cg_glob_stats SET total_nft_donated = (total_nft_donated + 1);
	RETURN NEW;
END;
$$ LANGUAGE plpgsql;
-- +goose StatementEnd
-- +goose StatementBegin
CREATE OR REPLACE FUNCTION on_nft_donation_delete() RETURNS trigger AS  $$
DECLARE
	v_cnt					NUMERIC;
BEGIN

	UPDATE cg_nft_stats
		SET
			num_donated = (num_donated - 1)
		WHERE contract_aid = OLD.token_aid;
	UPDATE cg_round_stats SET total_nft_donated = (total_nft_donated - 1) WHERE round_num=OLD.round_num;
	UPDATE cg_glob_stats SET total_nft_donated = (total_nft_donated - 1);
	RETURN OLD;
END;
$$ LANGUAGE plpgsql;
-- +goose StatementEnd
-- +goose StatementBegin
CREATE OR REPLACE FUNCTION on_prize_deposit_insert() RETURNS trigger AS  $$
DECLARE
	v_cnt						NUMERIC;
BEGIN

	-- Note: Statistics are updated by specific prize event triggers (cg_raffle_eth_prize, cg_chrono_warrior_prize, etc.)
	-- Not updating stats here to avoid double-counting
	RETURN NEW;
END;
$$ LANGUAGE plpgsql;
-- +goose StatementEnd
-- +goose StatementBegin
CREATE OR REPLACE FUNCTION on_prize_deposit_delete() RETURNS trigger AS  $$
DECLARE
BEGIN

	-- Note: Statistics are updated by specific prize event triggers (cg_raffle_eth_prize, cg_chrono_warrior_prize, etc.)
	-- Not updating stats here to avoid double-counting
	RETURN OLD;
END;
$$ LANGUAGE plpgsql;
-- +goose StatementEnd
-- +goose StatementBegin
CREATE OR REPLACE FUNCTION on_raffle_nft_winner_insert() RETURNS trigger AS  $$
DECLARE
	v_cnt						NUMERIC;
BEGIN

	UPDATE cg_raffle_nft_winner_stats
		SET
			num_won = (num_won + 1)
		WHERE winner_aid = NEW.winner_aid;
	GET DIAGNOSTICS v_cnt = ROW_COUNT;
	IF v_cnt = 0 THEN
		INSERT INTO cg_raffle_nft_winner_stats(winner_aid,num_won)
			VALUES(NEW.winner_aid,1);
	END IF;
	UPDATE cg_round_stats
		SET
			total_raffle_nfts = (total_raffle_nfts + 1)
		WHERE round_num=NEW.round_num;
	GET DIAGNOSTICS v_cnt = ROW_COUNT;
	IF v_cnt = 0 THEN
		INSERT INTO cg_round_stats(round_num,total_raffle_nfts) VALUES(NEW.round_num,1);
	END IF;
	IF NEW.is_staker THEN
		IF NEW.is_rwalk THEN
			UPDATE cg_staker_rwalk SET num_tokens_minted = (num_tokens_minted + 1)
				WHERE staker_aid=NEW.winner_aid;
			GET DIAGNOSTICS v_cnt = ROW_COUNT;
			IF v_cnt = 0 THEN
				INSERT INTO cg_staker_rwalk(staker_aid,num_tokens_minted) VALUES(NEW.winner_aid,1);
			END IF;
			UPDATE cg_stake_stats_rwalk SET total_nft_mints = (total_nft_mints + 1);
		ELSE
			UPDATE cg_staker_cst SET num_tokens_minted = (num_tokens_minted + 1)
				WHERE staker_aid=NEW.winner_aid;
			GET DIAGNOSTICS v_cnt = ROW_COUNT;
			IF v_cnt = 0 THEN
				INSERT INTO cg_staker_cst(staker_aid,num_tokens_minted) VALUES(NEW.winner_aid,1);
			END IF;
		END IF;
	END IF;
	
	-- Update CST prize total
	UPDATE cg_glob_stats SET total_cst_given_in_prizes = (total_cst_given_in_prizes + NEW.cst_amount);
	
	-- Update round stats
	UPDATE cg_round_stats SET total_cst_paid_in_prizes = (total_cst_paid_in_prizes + NEW.cst_amount), total_nfts_minted = (total_nfts_minted + 1) WHERE round_num = NEW.round_num;
	GET DIAGNOSTICS v_cnt = ROW_COUNT;
	IF v_cnt = 0 THEN
		INSERT INTO cg_round_stats(round_num, total_cst_paid_in_prizes, total_nfts_minted) VALUES (NEW.round_num, NEW.cst_amount, 1);
	END IF;
	
	-- Insert TWO records in cg_prize table with conditional ptype based on is_rwalk field
	IF NEW.is_rwalk THEN
		-- For RandomWalk stakers: CST (ptype=13) and CS NFT (ptype=14)
		INSERT INTO cg_prize(round_num,winner_index,ptype) VALUES(NEW.round_num,NEW.winner_idx,13);
		INSERT INTO cg_prize(round_num,winner_index,ptype) VALUES(NEW.round_num,NEW.winner_idx,14);
	ELSE
		-- For bidders: CST (ptype=11) and CS NFT (ptype=12)
		INSERT INTO cg_prize(round_num,winner_index,ptype) VALUES(NEW.round_num,NEW.winner_idx,11);
		INSERT INTO cg_prize(round_num,winner_index,ptype) VALUES(NEW.round_num,NEW.winner_idx,12);
	END IF;
	
	-- Update winner: prizes_count (+2), erc20_count (+1), erc721_count (+1)
	UPDATE cg_winner 
		SET 
			prizes_count = (prizes_count + 2),
			erc20_count = (erc20_count + 1),
			erc721_count = (erc721_count + 1)
		WHERE winner_aid = NEW.winner_aid;
	GET DIAGNOSTICS v_cnt = ROW_COUNT;
	IF v_cnt = 0 THEN
		INSERT INTO cg_winner(winner_aid, prizes_count, erc20_count, erc721_count) VALUES (NEW.winner_aid, 2, 1, 1);
	END IF;

	RETURN NEW;
END;
$$ LANGUAGE plpgsql;
-- +goose StatementEnd
-- +goose StatementBegin
CREATE OR REPLACE FUNCTION on_raffle_nft_winner_delete() RETURNS trigger AS  $$
DECLARE
BEGIN

	UPDATE cg_raffle_nft_winner_stats
		SET
			num_won = (num_won - 1)
		WHERE winner_aid = OLD.winner_aid;
	UPDATE cg_round_stats
		SET
			total_raffle_nfts = (total_raffle_nfts - 1)
		WHERE round_num=OLD.round_num;
	IF OLD.is_staker THEN
		IF OLD.is_rwalk THEN
			UPDATE cg_staker_rwalk SET num_tokens_minted = (num_tokens_minted - 1)
				WHERE staker_aid=OLD.winner_aid;
			UPDATE cg_stake_stats_rwalk SET total_nft_mints = (total_nft_mints - 1);
		ELSE
			UPDATE cg_staker_cst SET num_tokens_minted = (num_tokens_minted - 1)
				WHERE staker_aid=OLD.winner_aid;
		END IF;
	END IF;
	
	-- Update CST prize total
	UPDATE cg_glob_stats SET total_cst_given_in_prizes = (total_cst_given_in_prizes - OLD.cst_amount);
	
	-- Update round stats
	UPDATE cg_round_stats SET total_cst_paid_in_prizes = (total_cst_paid_in_prizes - OLD.cst_amount) WHERE round_num = OLD.round_num;
	UPDATE cg_round_stats SET total_nfts_minted = (total_nfts_minted - 1) WHERE round_num = OLD.round_num;
	
	-- Remove TWO corresponding records from cg_prize table with conditional ptype based on is_rwalk field
	IF OLD.is_rwalk THEN
		-- For RandomWalk stakers: CST (ptype=13) and CS NFT (ptype=14)
		DELETE FROM cg_prize WHERE round_num=OLD.round_num AND winner_index=OLD.winner_idx AND ptype=13;
		DELETE FROM cg_prize WHERE round_num=OLD.round_num AND winner_index=OLD.winner_idx AND ptype=14;
	ELSE
		-- For bidders: CST (ptype=11) and CS NFT (ptype=12)
		DELETE FROM cg_prize WHERE round_num=OLD.round_num AND winner_index=OLD.winner_idx AND ptype=11;
		DELETE FROM cg_prize WHERE round_num=OLD.round_num AND winner_index=OLD.winner_idx AND ptype=12;
	END IF;

	-- Update winner counts
	UPDATE cg_winner 
		SET 
			prizes_count = GREATEST(0, prizes_count - 2),
			erc20_count = GREATEST(0, erc20_count - 1),
			erc721_count = GREATEST(0, erc721_count - 1)
		WHERE winner_aid = OLD.winner_aid;

	RETURN OLD;
END;
$$ LANGUAGE plpgsql;
-- +goose StatementEnd
-- +goose StatementBegin
CREATE OR REPLACE FUNCTION on_endurance_winner_insert() RETURNS trigger AS  $$
DECLARE
	v_cnt						NUMERIC;
BEGIN
	-- Note: The Solidity EnduranceChampionPrizePaid event does not emit a winner_index.
	-- There is exactly one endurance champion per round, so winner_index is implicitly 0.

	UPDATE cg_glob_stats SET total_cst_given_in_prizes = (total_cst_given_in_prizes + NEW.erc20_amount);
	
	-- Update round stats
	UPDATE cg_round_stats SET total_cst_paid_in_prizes = (total_cst_paid_in_prizes + NEW.erc20_amount), total_nfts_minted = (total_nfts_minted + 1) WHERE round_num = NEW.round_num;
	GET DIAGNOSTICS v_cnt = ROW_COUNT;
	IF v_cnt = 0 THEN
		INSERT INTO cg_round_stats(round_num, total_cst_paid_in_prizes, total_nfts_minted) VALUES (NEW.round_num, NEW.erc20_amount, 1);
	END IF;

	-- Insert TWO records in cg_prize table for Endurance Champion prizes
	-- Winner index is always 0 (one endurance champion per round)
	-- 1) ERC721 CS NFT prize (ptype=5)
	INSERT INTO cg_prize(round_num,winner_index,ptype) VALUES(NEW.round_num,0,5);
	-- 2) ERC20 CST token prize (ptype=6)
	INSERT INTO cg_prize(round_num,winner_index,ptype) VALUES(NEW.round_num,0,6);
	
	-- Update winner: prizes_count (+2), erc20_count (+1), erc721_count (+1)
	UPDATE cg_winner 
		SET 
			prizes_count = (prizes_count + 2),
			erc20_count = (erc20_count + 1),
			erc721_count = (erc721_count + 1)
		WHERE winner_aid = NEW.winner_aid;
	GET DIAGNOSTICS v_cnt = ROW_COUNT;
	IF v_cnt = 0 THEN
		INSERT INTO cg_winner(winner_aid, prizes_count, erc20_count, erc721_count) VALUES (NEW.winner_aid, 2, 1, 1);
	END IF;

	RETURN NEW;
END;
$$ LANGUAGE plpgsql;
-- +goose StatementEnd
-- +goose StatementBegin
CREATE OR REPLACE FUNCTION on_endurance_winner_delete() RETURNS trigger AS  $$
DECLARE
BEGIN
	-- Note: Winner index is always 0 (one endurance champion per round)

	UPDATE cg_glob_stats SET total_cst_given_in_prizes = (total_cst_given_in_prizes - OLD.erc20_amount);
	
	-- Update round stats
	UPDATE cg_round_stats SET total_cst_paid_in_prizes = (total_cst_paid_in_prizes - OLD.erc20_amount) WHERE round_num = OLD.round_num;
	UPDATE cg_round_stats SET total_nfts_minted = (total_nfts_minted - 1) WHERE round_num = OLD.round_num;

	-- Remove BOTH corresponding records from cg_prize table
	-- 1) ERC721 CS NFT prize (ptype=5)
	DELETE FROM cg_prize WHERE round_num=OLD.round_num AND winner_index=0 AND ptype=5;
	-- 2) ERC20 CST token prize (ptype=6)
	DELETE FROM cg_prize WHERE round_num=OLD.round_num AND winner_index=0 AND ptype=6;

	-- Update winner counts
	UPDATE cg_winner 
		SET 
			prizes_count = GREATEST(0, prizes_count - 2),
			erc20_count = GREATEST(0, erc20_count - 1),
			erc721_count = GREATEST(0, erc721_count - 1)
		WHERE winner_aid = OLD.winner_aid;

	RETURN OLD;
END;
$$ LANGUAGE plpgsql;
-- +goose StatementEnd
-- +goose StatementBegin
CREATE OR REPLACE FUNCTION on_lastcst_winner_insert() RETURNS trigger AS  $$
DECLARE
	v_cnt						NUMERIC;
BEGIN
	-- Note: The Solidity LastCstBidderPrizePaid event does not emit a winner_index.
	-- There is exactly one last CST bidder per round, so winner_index is implicitly 0.

	UPDATE cg_glob_stats SET total_cst_given_in_prizes = (total_cst_given_in_prizes + NEW.erc20_amount);
	
	-- Update round stats
	UPDATE cg_round_stats SET total_cst_paid_in_prizes = (total_cst_paid_in_prizes + NEW.erc20_amount), total_nfts_minted = (total_nfts_minted + 1) WHERE round_num = NEW.round_num;
	GET DIAGNOSTICS v_cnt = ROW_COUNT;
	IF v_cnt = 0 THEN
		INSERT INTO cg_round_stats(round_num, total_cst_paid_in_prizes, total_nfts_minted) VALUES (NEW.round_num, NEW.erc20_amount, 1);
	END IF;

	-- Insert TWO records in cg_prize table for Last CST Bidder prizes
	-- Winner index is always 0 (one last CST bidder per round)
	-- 1) ERC721 CS NFT prize (ptype=3)
	INSERT INTO cg_prize(round_num,winner_index,ptype) VALUES(NEW.round_num,0,3);
	-- 2) ERC20 CST token prize (ptype=4)
	INSERT INTO cg_prize(round_num,winner_index,ptype) VALUES(NEW.round_num,0,4);
	
	-- Update winner: prizes_count (+2), erc20_count (+1), erc721_count (+1)
	UPDATE cg_winner 
		SET 
			prizes_count = (prizes_count + 2),
			erc20_count = (erc20_count + 1),
			erc721_count = (erc721_count + 1)
		WHERE winner_aid = NEW.winner_aid;
	GET DIAGNOSTICS v_cnt = ROW_COUNT;
	IF v_cnt = 0 THEN
		INSERT INTO cg_winner(winner_aid, prizes_count, erc20_count, erc721_count) VALUES (NEW.winner_aid, 2, 1, 1);
	END IF;

	RETURN NEW;
END;
$$ LANGUAGE plpgsql;
-- +goose StatementEnd
-- +goose StatementBegin
CREATE OR REPLACE FUNCTION on_lastcst_winner_delete() RETURNS trigger AS  $$
DECLARE
BEGIN
	-- Note: Winner index is always 0 (one last CST bidder per round)

	UPDATE cg_glob_stats SET total_cst_given_in_prizes = (total_cst_given_in_prizes - OLD.erc20_amount);
	
	-- Update round stats
	UPDATE cg_round_stats SET total_cst_paid_in_prizes = (total_cst_paid_in_prizes - OLD.erc20_amount) WHERE round_num = OLD.round_num;
	UPDATE cg_round_stats SET total_nfts_minted = (total_nfts_minted - 1) WHERE round_num = OLD.round_num;

	-- Remove BOTH corresponding records from cg_prize table
	-- 1) ERC721 CS NFT prize (ptype=3)
	DELETE FROM cg_prize WHERE round_num=OLD.round_num AND winner_index=0 AND ptype=3;
	-- 2) ERC20 CST token prize (ptype=4)
	DELETE FROM cg_prize WHERE round_num=OLD.round_num AND winner_index=0 AND ptype=4;

	-- Update winner counts
	UPDATE cg_winner 
		SET 
			prizes_count = GREATEST(0, prizes_count - 2),
			erc20_count = GREATEST(0, erc20_count - 1),
			erc721_count = GREATEST(0, erc721_count - 1)
		WHERE winner_aid = OLD.winner_aid;

	RETURN OLD;
END;
$$ LANGUAGE plpgsql;
-- +goose StatementEnd
-- +goose StatementBegin
CREATE OR REPLACE FUNCTION on_erc721transfer_insert() RETURNS trigger AS  $$
DECLARE
	v_cnt						NUMERIC;
BEGIN

	UPDATE cg_mint_event
		SET
			cur_owner_aid = NEW.to_aid
		WHERE token_id=NEW.token_id;

	IF NEW.from_aid = NEW.to_aid THEN -- self transfer
		UPDATE cg_transfer_stats SET erc721_num_transfers = (erc721_num_transfers + 1)
		WHERE user_aid = NEW.from_aid;
		GET DIAGNOSTICS v_cnt = ROW_COUNT;
		IF v_cnt = 0 THEN
			INSERT INTO cg_transfer_stats(user_aid,erc721_num_transfers) VALUES(NEW.from_aid,1);
		END IF;
	ELSE
		UPDATE cg_transfer_stats SET erc721_num_transfers = (erc721_num_transfers + 1)
		WHERE user_aid = NEW.from_aid;
		GET DIAGNOSTICS v_cnt = ROW_COUNT;
		IF v_cnt = 0 THEN
			INSERT INTO cg_transfer_stats(user_aid,erc721_num_transfers) VALUES(NEW.from_aid,1);
		END IF;
		UPDATE cg_transfer_stats SET erc721_num_transfers = (erc721_num_transfers + 1)
		WHERE user_aid = NEW.to_aid;
		GET DIAGNOSTICS v_cnt = ROW_COUNT;
		IF v_cnt = 0 THEN
			INSERT INTO cg_transfer_stats(user_aid,erc721_num_transfers) VALUES(NEW.to_aid,1);
		END IF;
	END IF;

	RETURN NEW;
END;
$$ LANGUAGE plpgsql;
-- +goose StatementEnd
-- +goose StatementBegin
CREATE OR REPLACE FUNCTION on_erc721transfer_delete() RETURNS trigger AS  $$
DECLARE
BEGIN

	IF OLD.from_aid = OLD.to_aid THEN -- self transfer
		UPDATE cg_transfer_stats SET erc721_num_transfers = (erc721_num_transfers - 1)
		WHERE user_aid = OLD.from_aid;
	ELSE
		UPDATE cg_transfer_stats SET erc721_num_transfers = (erc721_num_transfers - 1)
		WHERE user_aid = OLD.from_aid;
		UPDATE cg_transfer_stats SET erc721_num_transfers = (erc721_num_transfers - 1)
		WHERE user_aid = OLD.to_aid;
	END IF;
	RETURN OLD;
END;
$$ LANGUAGE plpgsql;
-- +goose StatementEnd
-- +goose StatementBegin
CREATE OR REPLACE FUNCTION on_erc20_transfer_insert() RETURNS trigger AS  $$
DECLARE
	v_cnt						NUMERIC;
	v_from_addr					TEXT;
	v_to_addr					TEXT;
BEGIN


	SELECT addr FROM address WHERE address_id=NEW.from_aid INTO v_from_addr;
	SELECT addr FROM address WHERE address_id=NEW.to_aid INTO v_to_addr;
	--- update From balance
	IF v_to_addr = '0x0000000000000000000000000000000000000000' THEN -- burn
		UPDATE cg_costok_owner SET cur_balance = (cur_balance - NEW.value)
		WHERE owner_aid=NEW.from_aid;
		GET DIAGNOSTICS v_cnt = ROW_COUNT;
		IF v_cnt = 0 THEN
			INSERT INTO cg_costok_owner(owner_aid,cur_balance) VALUES(NEW.from_aid,-NEW.value);
		END IF;
	ELSE 
		IF v_from_addr != '0x0000000000000000000000000000000000000000' THEN --regular transfer
			UPDATE cg_costok_owner SET cur_balance = (cur_balance - NEW.value)
			WHERE owner_aid = NEW.from_aid;
			GET DIAGNOSTICS v_cnt = ROW_COUNT;
			IF v_cnt = 0 THEN
				INSERT INTO cg_costok_owner(owner_aid,cur_balance) VALUES(NEW.from_aid,-NEW.value);
			END IF;
			UPDATE cg_costok_owner SET cur_balance = (cur_balance + NEW.value)
			WHERE owner_aid = NEW.to_aid;
			GET DIAGNOSTICS v_cnt = ROW_COUNT;
			IF v_cnt = 0 THEN
				INSERT INTO cg_costok_owner(owner_aid,cur_balance) VALUES(NEW.to_aid,NEW.value);
			END IF;
		ELSE  -- mint
			UPDATE cg_costok_owner SET cur_balance = (cur_balance + NEW.value)
			WHERE owner_aid=NEW.to_aid;
			GET DIAGNOSTICS v_cnt = ROW_COUNT;
			IF v_cnt = 0 THEN
				INSERT INTO cg_costok_owner(owner_aid,cur_balance) VALUES(NEW.to_aid,NEW.value);
			END IF;
		END IF;
	END IF;
	IF NEW.from_aid = NEW.to_aid THEN -- self transfer
		UPDATE cg_transfer_stats SET erc20_num_transfers = (erc20_num_transfers + 1)
		WHERE user_aid = NEW.from_aid;
		GET DIAGNOSTICS v_cnt = ROW_COUNT;
		IF v_cnt = 0 THEN
			INSERT INTO cg_transfer_stats(user_aid,erc20_num_transfers) VALUES(NEW.from_aid,1);
		END IF;
	ELSE
		UPDATE cg_transfer_stats SET erc20_num_transfers = (erc20_num_transfers + 1)
		WHERE user_aid = NEW.from_aid;
		GET DIAGNOSTICS v_cnt = ROW_COUNT;
		IF v_cnt = 0 THEN
			INSERT INTO cg_transfer_stats(user_aid,erc20_num_transfers) VALUES(NEW.from_aid,1);
		END IF;
		UPDATE cg_transfer_stats SET erc20_num_transfers = (erc20_num_transfers + 1)
		WHERE user_aid = NEW.to_aid;
		GET DIAGNOSTICS v_cnt = ROW_COUNT;
		IF v_cnt = 0 THEN
			INSERT INTO cg_transfer_stats(user_aid,erc20_num_transfers) VALUES(NEW.to_aid,1);
		END IF;
	END IF;

	RETURN NEW;
END;
$$ LANGUAGE plpgsql;
-- +goose StatementEnd
-- +goose StatementBegin
CREATE OR REPLACE FUNCTION on_erc20_transfer_delete() RETURNS trigger AS  $$
DECLARE
	v_from_addr					TEXT;
	v_to_addr					TEXT;
BEGIN

	SELECT addr FROM address WHERE address_id=OLD.from_aid INTO v_from_addr;
	SELECT addr FROM address WHERE address_id=OLD.to_aid INTO v_to_addr;
	--- update From balance
	IF v_to_addr = '0x0000000000000000000000000000000000000000' THEN -- burn
		UPDATE cg_costok_owner SET cur_balance = (cur_balance + OLD.value)
		WHERE owner_aid=OLD.from_aid;
	ELSE 
		IF v_from_addr != '0x0000000000000000000000000000000000000000' THEN --regular transfer
			UPDATE cg_costok_owner SET cur_balance = (cur_balance + OLD.value)
			WHERE owner_aid = OLD.from_aid;
			UPDATE cg_costok_owner SET cur_balance = (cur_balance - OLD.value)
			WHERE owner_aid = OLD.to_aid;
		ELSE  -- mint
			UPDATE cg_costok_owner SET cur_balance = (cur_balance - OLD.value)
			WHERE owner_aid=OLD.to_aid;
		END IF;
	END IF;
	IF OLD.from_aid = OLD.to_aid THEN -- self transfer
		UPDATE cg_transfer_stats SET erc20_num_transfers = (erc20_num_transfers - 1)
		WHERE user_aid = OLD.from_aid;
	ELSE
		UPDATE cg_transfer_stats SET erc20_num_transfers = (erc20_num_transfers - 1)
		WHERE user_aid = OLD.from_aid;
		UPDATE cg_transfer_stats SET erc20_num_transfers = (erc20_num_transfers - 1)
		WHERE user_aid = OLD.to_aid;
	END IF;
	RETURN OLD;
END;
$$ LANGUAGE plpgsql;
-- +goose StatementEnd
-- +goose StatementBegin
CREATE OR REPLACE FUNCTION on_donated_tok_claimed_insert() RETURNS trigger AS  $$
DECLARE
	v_amount					DECIMAL;
BEGIN

	UPDATE cg_erc20_donation_stats
		SET
			total_amount = (total_amount - NEW.amount)
		WHERE round_num = NEW.round_num AND token_aid = NEW.token_aid;
	SELECT total_amount FROM cg_erc20_donation_stats 
		WHERE round_num=NEW.round_num AND token_aid=NEW.token_aid
		INTO v_amount;
	IF v_amount = 0 THEN
		UPDATE cg_erc20_donation_stats SET claimed = 't'
			WHERE round_num = NEW.round_num AND token_aid = NEW.token_aid;
	END IF;

	RETURN NEW;
END;
$$ LANGUAGE plpgsql;
-- +goose StatementEnd
-- +goose StatementBegin
CREATE OR REPLACE FUNCTION on_donated_tok_claimed_delete() RETURNS trigger AS  $$
DECLARE
	v_amount					DECIMAL;
BEGIN

	UPDATE cg_erc20_donation_stats
		SET
			total_amount = (total_amount + OLD.amount)
		WHERE round_num = OLD.round_num AND token_aid = OLD.token_aid;
	SELECT total_amount FROM cg_erc20_donation_stats 
		WHERE round_num=OLD.round_num AND token_aid=OLD.token_aid
		INTO v_amount;
	IF v_amount <> 0 THEN
		UPDATE cg_erc20_donation_stats SET claimed = 'f'
			WHERE round_num = OLD.round_num AND token_aid = OLD.token_aid;
	END IF;
	RETURN OLD;
END;
$$ LANGUAGE plpgsql;
-- +goose StatementEnd
-- +goose StatementBegin
CREATE OR REPLACE FUNCTION on_donated_nft_claimed_insert() RETURNS trigger AS  $$
DECLARE
	v_cnt						BIGINT;
	new_unclaimed_value			BIGINT;
	v_round_winner_aid			BIGINT;
	v_timeout_ts				BIGINT;
	v_claim_ts					BIGINT;
BEGIN

	-- Check if the claimer is the winner of this round
	SELECT winner_aid, timeout FROM cg_prize_claim WHERE round_num = NEW.round_num INTO v_round_winner_aid, v_timeout_ts;
	
	IF v_round_winner_aid IS NULL THEN
		-- Round hasn't been won yet - this claim event is being processed before the prize claim event
		-- This can happen if events from different contracts are processed out of order
		-- We'll skip unclaimed_nfts tracking for now - it will be corrected when prize claim is processed
		RAISE NOTICE 'on_donated_nft_claimed_insert() no winner found yet for round_num=%. Claimer winner_aid=%, idx=%, evtlog_id=%. Available rounds in cg_prize_claim: %. Skipping unclaimed_nfts tracking - will be reconciled when MainPrizeClaimed is processed.',
			NEW.round_num,
			NEW.winner_aid,
			NEW.idx,
			NEW.evtlog_id,
			(SELECT STRING_AGG(round_num::TEXT, ',') FROM cg_prize_claim);
		-- Return without updating unclaimed_nfts
		RETURN NEW;
	END IF;
	
	-- Get the claim timestamp
	SELECT EXTRACT(EPOCH FROM time_stamp)::BIGINT FROM cg_donated_nft_claimed WHERE id = NEW.id INTO v_claim_ts;
	
	-- ALWAYS decrement unclaimed_nfts for the ORIGINAL ROUND WINNER, regardless of who claims
	-- The unclaimed_nfts counter tracks how many donated NFTs the round winner has pending
	-- When ANYONE claims (winner or non-winner after timeout), that count should decrease for the winner
	UPDATE cg_winner
		SET
			unclaimed_nfts = GREATEST(0, unclaimed_nfts - 1)
		WHERE winner_aid = v_round_winner_aid
		RETURNING unclaimed_nfts INTO new_unclaimed_value;
	GET DIAGNOSTICS v_cnt = ROW_COUNT;
	IF v_cnt = 0 THEN
		-- Winner record should exist from on_prize_claim_insert(), but handle edge case
		RAISE NOTICE 'on_donated_nft_claimed_insert() winner record does not exist for round winner_aid=% (round_num=%), claimer=%. Creating record.',
			v_round_winner_aid, NEW.round_num, NEW.winner_aid;
		INSERT INTO cg_winner(winner_aid, unclaimed_nfts) VALUES(v_round_winner_aid, 0);
	END IF;
	
	-- Log if this is a timeout claim by non-winner (for debugging)
	IF v_round_winner_aid != NEW.winner_aid THEN
		IF v_claim_ts >= v_timeout_ts THEN
			RAISE NOTICE 'on_donated_nft_claimed_insert() timeout claim: non-winner (%) claimed NFT from round % after timeout. Original winner: %, timeout_ts: %, claim_ts: %',
				NEW.winner_aid, NEW.round_num, v_round_winner_aid, v_timeout_ts, v_claim_ts;
		ELSE
			-- Non-winner claiming before timeout - this should not happen per Solidity logic
			RAISE WARNING 'on_donated_nft_claimed_insert() UNEXPECTED: non-winner (%) claimed NFT from round % BEFORE timeout. Original winner: %, timeout_ts: %, claim_ts: %. This may indicate a bug.',
				NEW.winner_aid, NEW.round_num, v_round_winner_aid, v_timeout_ts, v_claim_ts;
		END IF;
	END IF;
	
	RETURN NEW;
END;
$$ LANGUAGE plpgsql;
-- +goose StatementEnd
-- +goose StatementBegin
CREATE OR REPLACE FUNCTION on_donated_nft_claimed_delete() RETURNS trigger AS  $$
DECLARE
	v_round_winner_aid			BIGINT;
BEGIN
	-- Get the original round winner
	SELECT winner_aid FROM cg_prize_claim WHERE round_num = OLD.round_num INTO v_round_winner_aid;
	
	IF v_round_winner_aid IS NOT NULL THEN
		-- Increment unclaimed_nfts for the original round winner
		UPDATE cg_winner
			SET
				unclaimed_nfts = (unclaimed_nfts + 1)
			WHERE winner_aid = v_round_winner_aid;
	END IF;
	
	RETURN OLD;
END;
$$ LANGUAGE plpgsql;
-- +goose StatementEnd
-- +goose StatementBegin
CREATE OR REPLACE FUNCTION on_mint_insert() RETURNS trigger AS  $$
DECLARE
	v_cnt						NUMERIC;
BEGIN
	-- REMOVED: cg_winner updates (now tracked by prize-specific triggers)
	-- Main prize NFT is tracked by on_prize_claim_insert()
	-- Raffle NFT is tracked by on_raffle_nft_winner_insert()
	-- Endurance NFT is tracked by on_endurance_winner_insert()
	-- Last CST NFT is tracked by on_lastcst_winner_insert()
	-- Chrono Warrior NFT is tracked by on_chrono_warrior_insert()
	
	UPDATE cg_glob_stats SET num_mints = (num_mints + 1);
	RETURN NEW;
END;
$$ LANGUAGE plpgsql;
-- +goose StatementEnd
-- +goose StatementBegin
CREATE OR REPLACE FUNCTION on_mint_delete() RETURNS trigger AS  $$
DECLARE
BEGIN
	-- REMOVED: cg_winner updates (now tracked by prize-specific triggers)
	
	UPDATE cg_glob_stats SET num_mints = (num_mints - 1);
	RETURN OLD;
END;
$$ LANGUAGE plpgsql;
-- +goose StatementEnd
-- +goose StatementBegin
CREATE OR REPLACE FUNCTION on_mint_update() RETURNS trigger AS  $$
DECLARE
	v_cnt						NUMERIC;
BEGIN

	RETURN NEW;
END;
$$ LANGUAGE plpgsql;
-- +goose StatementEnd
-- +goose StatementBegin
CREATE OR REPLACE FUNCTION on_prize_withdrawal_insert() RETURNS trigger AS  $$
DECLARE
	v_cnt						NUMERIC;
BEGIN

	UPDATE cg_raffle_winner_stats
		SET
			withdrawal_sum = (withdrawal_sum + NEW.amount),	-- data for historical purposes
			amount_sum = (amount_sum - NEW.amount)			-- current amount available to withdraw
		WHERE winner_aid = NEW.winner_aid;
	GET DIAGNOSTICS v_cnt = ROW_COUNT;
	IF v_cnt = 0 THEN
		INSERT INTO cg_raffle_winner_stats(winner_aid,withdrawal_sum) VALUES(NEW.winner_aid,NEW.amount);
	END IF;
	-- Mark deposits as claimed for this specific round (per-round withdrawal)
	UPDATE cg_prize_deposit 
		SET claimed=TRUE, withdrawal_id=NEW.evtlog_id 
		WHERE (round_num=NEW.round_num) 
		AND (withdrawal_id=0) 
		AND (winner_aid=NEW.winner_aid);
	UPDATE cg_glob_stats SET total_raffle_eth_withdrawn = (total_raffle_eth_withdrawn + NEW.amount);
	RETURN NEW;
END;
$$ LANGUAGE plpgsql;
-- +goose StatementEnd
-- +goose StatementBegin
CREATE OR REPLACE FUNCTION on_prize_withdrawal_delete() RETURNS trigger AS  $$
DECLARE
BEGIN

	UPDATE cg_raffle_winner_stats
		SET
			withdrawal_sum = (withdrawal_sum - OLD.amount),
			amount_sum = (amount_sum + OLD.amount)
		WHERE winner_aid = OLD.winner_aid;
	-- Unclaim deposits for this specific round
	UPDATE cg_prize_deposit 
		SET claimed=FALSE, withdrawal_id=0 
		WHERE (round_num=OLD.round_num) 
		AND (withdrawal_id = OLD.evtlog_id) 
		AND (winner_aid=OLD.winner_aid);
	UPDATE cg_glob_stats SET total_raffle_eth_withdrawn = (total_raffle_eth_withdrawn - OLD.amount);
	RETURN OLD;
END;
$$ LANGUAGE plpgsql;
-- +goose StatementEnd
-- +goose StatementBegin
CREATE OR REPLACE FUNCTION on_token_name_insert() RETURNS trigger AS  $$
DECLARE
BEGIN

	UPDATE cg_mint_event SET token_name = NEW.token_name WHERE token_id=NEW.token_id;
	RETURN NEW;
END;
$$ LANGUAGE plpgsql;
-- +goose StatementEnd
-- +goose StatementBegin
CREATE OR REPLACE FUNCTION on_token_name_delete() RETURNS trigger AS  $$
DECLARE
BEGIN

	UPDATE cg_mint_event SET token_name = '' WHERE token_id = OLD.token_id;
	RETURN OLD;
END;
$$ LANGUAGE plpgsql;
-- +goose StatementEnd
-- +goose StatementBegin
CREATE OR REPLACE FUNCTION on_donation_sent_insert() RETURNS trigger AS  $$
DECLARE
	v_cnt						NUMERIC;
BEGIN

	UPDATE cg_glob_stats 
		SET 
			num_withdrawals = (num_withdrawals + 1),
			sum_withdrawals = (sum_withdrawals + NEW.amount);

	RETURN NEW;
END;
$$ LANGUAGE plpgsql;
-- +goose StatementEnd
-- +goose StatementBegin
CREATE OR REPLACE FUNCTION on_donation_sent_delete() RETURNS trigger AS  $$
DECLARE
BEGIN

	UPDATE cg_glob_stats
		SET
			num_withdrawals = (num_withdrawals - 1),
			sum_withdrawals = (sum_withdrawals - OLD.amount);

	RETURN OLD;
END;
$$ LANGUAGE plpgsql;
-- +goose StatementEnd
-- +goose StatementBegin
CREATE OR REPLACE FUNCTION on_eth_deposit_insert() RETURNS trigger AS  $$
DECLARE
	v_amount_per_token DECIMAL;
	v_mod DECIMAL;
	v_rec RECORD;
	v_cnt						NUMERIC;
	v_prev_num_tokens		INT;
	v_prev_amount			DECIMAL;
	v_prev_amount_per_token	DECIMAL;
	v_tokens_added			INT;
BEGIN

	IF NEW.num_staked_nfts > 0 THEN
		SELECT accumulated_nfts,accumulated_amount,accumulated_per_token FROM cg_staking_eth_deposit ORDER BY deposit_id DESC OFFSET 1 LIMIT 1 INTO v_prev_num_tokens,v_prev_amount,v_prev_amount_per_token;
		IF v_prev_num_tokens IS NULL THEN
			v_prev_num_tokens:=0;
			v_tokens_added:=NEW.num_staked_nfts;
		ELSE
			v_tokens_added:=NEW.num_staked_nfts-v_prev_num_tokens;
		END IF;
		IF v_prev_amount IS NULL THEN
			v_prev_amount:=0;
		END IF;
		v_mod := MOD(NEW.deposit_amount,NEW.num_staked_nfts);
		v_amount_per_token := (NEW.deposit_amount - v_mod) / NEW.num_staked_nfts;
		UPDATE cg_staker_cst
			SET total_reward = (total_reward + (v_amount_per_token*total_tokens_staked)),
				unclaimed_reward = (unclaimed_reward + (v_amount_per_token*total_tokens_staked))
			WHERE total_tokens_staked > 0;
		UPDATE cg_stake_stats_cst
			SET
				total_reward_amount = (total_reward_amount + (NEW.deposit_amount - v_mod)),
				total_unclaimed_reward = (total_unclaimed_reward + (NEW.deposit_amount - v_mod)),
				num_deposits = (num_deposits + 1),
				total_modulo = (total_modulo + v_mod)
			;
		FOR v_rec IN (SELECT count(*) AS num_toks,staker_aid FROM cg_staked_token_cst GROUP BY staker_aid)
		LOOP
			INSERT INTO cg_staker_deposit(staker_aid,deposit_id,tokens_staked,amount_to_claim,amount_deposited)
				VALUES(v_rec.staker_aid,NEW.deposit_id,v_rec.num_toks,v_amount_per_token*v_rec.num_toks,v_amount_per_token*v_rec.num_toks);
		END LOOP;
		FOR v_rec IN (SELECT token_id,stake_action_id,staker_aid FROM cg_staked_token_cst)
		LOOP
			UPDATE cg_staked_token_cst_rewards 
				SET accumulated_reward = (accumulated_reward + v_amount_per_token)
				WHERE stake_action_id=v_rec.stake_action_id;
			GET DIAGNOSTICS v_cnt = ROW_COUNT;
			IF v_cnt = 0 THEN
				INSERT INTO cg_staked_token_cst_rewards(staker_aid,token_id,stake_action_id,accumulated_reward)
					VALUES(v_rec.staker_aid,v_rec.token_id,v_rec.stake_action_id,v_amount_per_token);
			END IF;
			INSERT INTO cg_st_reward(staker_aid,action_id,token_id,deposit_id,round_num,reward)
				VALUES(v_rec.staker_aid,v_rec.stake_action_id,v_rec.token_id,NEW.deposit_id,NEW.round_num,v_amount_per_token);
		END LOOP;
		UPDATE cg_staking_eth_deposit 
			SET 
				
				accumulated_amount = (v_prev_amount + NEW.deposit_amount),
				accumulated_per_token = (v_prev_amount_per_token + v_amount_per_token),
				accumulated_nfts = num_staked_nfts,
				num_staked_nfts = v_tokens_added
			WHERE id=NEW.id;
	END IF;

	-- Insert record in cg_prize table with ptype=15 for Staking Deposit ETH (for CS NFT stakers)
	-- Using winner_index=0 to indicate this is a general staking reward, not a specific winner's prize
	INSERT INTO cg_prize(round_num,winner_index,ptype) VALUES(NEW.round_num,0,15);

	RETURN NEW;
END;
$$ LANGUAGE plpgsql;
-- +goose StatementEnd
-- +goose StatementBegin
CREATE OR REPLACE FUNCTION on_eth_deposit_delete() RETURNS trigger AS  $$
DECLARE
	v_amount_per_token DECIMAL;
	v_mod DECIMAL;
	v_rec RECORD;
BEGIN

	IF OLD.num_staked_nfts > 0 THEN
		v_mod := MOD(OLD.deposit_amount,OLD.num_staked_nfts);
		v_amount_per_token := (OLD.deposit_amount - v_mod) / OLD.num_staked_nfts;
		FOR v_rec IN (SELECT count(token_id) AS num_toks,staker_aid FROM cg_staked_token_cst GROUP BY staker_aid)
		LOOP
			UPDATE cg_staker_cst
				SET total_reward = (total_reward -  (v_amount_per_token*v_rec.num_toks)),
					unclaimed_reward = (unclaimed_reward -  (v_amount_per_token*v_rec.num_toks))
				WHERE total_tokens_staked > 0;
		END LOOP;
		UPDATE cg_stake_stats_cst
			SET 
				total_reward_amount = (total_reward_amount - (OLD.deposit_amount - v_mod)),
				total_unclaimed_reward = (total_unclaimed_reward - (OLD.deposit_amount - v_mod)),
				num_deposits = (num_deposits - 1),
				total_modulo = (total_modulo - v_mod)
			;
		DELETE FROM cg_staker_deposit WHERE deposit_id=OLD.deposit_id;
		DELETE FROM cg_st_reward WHERE deposit_id=OLD.deposit_id;
	ELSE   
	END IF;

	-- Remove corresponding record from cg_prize table
	DELETE FROM cg_prize WHERE round_num=OLD.round_num AND winner_index=0 AND ptype=15;

	RETURN OLD;
END;
$$ LANGUAGE plpgsql;
-- +goose StatementEnd
-- +goose StatementBegin
CREATE OR REPLACE FUNCTION on_marketing_rewards_insert() RETURNS trigger AS  $$
DECLARE
	v_cnt						NUMERIC;
BEGIN

	UPDATE cg_glob_stats
		SET	total_mkt_rewards= (total_mkt_rewards + NEW.amount),
			num_mkt_rewards = (num_mkt_rewards + 1)
		;

	RETURN NEW;
END;
$$ LANGUAGE plpgsql;
-- +goose StatementEnd
-- +goose StatementBegin
CREATE OR REPLACE FUNCTION on_marketing_rewards_delete() RETURNS trigger AS  $$
DECLARE
BEGIN

	UPDATE cg_glob_stats
		SET	total_mkt_rewards= (total_mkt_rewards - OLD.amount),
			num_mkt_rewards = (num_mkt_rewards - 1)
		;

	RETURN OLD;
END;
$$ LANGUAGE plpgsql;
-- +goose StatementEnd
-- +goose StatementBegin
CREATE OR REPLACE FUNCTION on_direct_donation_insert() RETURNS trigger AS  $$
DECLARE
	v_cnt						NUMERIC;
BEGIN

	UPDATE cg_glob_stats 
		SET 
			num_direct_donations = (num_direct_donations + 1),
			direct_donations = (direct_donations + NEW.amount);
	UPDATE cg_round_stats
		SET
			donations_round_total = (donations_round_total + NEW.amount),
			donations_round_count = (donations_round_count + 1)
		WHERE NEW.round_num=round_num;
	GET DIAGNOSTICS v_cnt = ROW_COUNT;
	IF v_cnt = 0 THEN
		INSERT INTO cg_round_stats(round_num,donations_round_total,donations_round_count) VALUES (NEW.round_num,NEW.amount,1);
	END IF;
	UPDATE cg_donor
		SET
			total_eth_donated = (total_eth_donated + NEW.amount),
			count_donations = (count_donations + 1)
		WHERE donor_aid = NEW.donor_aid;
	GET DIAGNOSTICS v_cnt = ROW_COUNT;
	IF v_cnt = 0 THEN
		INSERT INTO cg_donor(donor_aid,count_donations,total_eth_donated)
			VALUES(NEW.donor_aid,1,NEW.amount);
	END IF;
	RETURN NEW;
END;
$$ LANGUAGE plpgsql;
-- +goose StatementEnd
-- +goose StatementBegin
CREATE OR REPLACE FUNCTION on_direct_donation_delete() RETURNS trigger AS  $$
DECLARE
BEGIN

	UPDATE cg_glob_stats 
		SET 
			num_direct_donations = (num_direct_donations - 1),
			direct_donations = (direct_donations - OLD.amount);
	UPDATE cg_round_stats
		SET
			donations_round_total = (donations_round_total - OLD.amount),
			donations_round_count = (donations_round_count - 1)
		WHERE OLD.round_num=round_num;
	UPDATE cg_donor
		SET
			total_eth_donated = (total_eth_donated - OLD.amount),
			count_donations = (count_donations - 1)
		WHERE donor_aid = OLD.donor_aid;
	RETURN OLD;
END;
$$ LANGUAGE plpgsql;
-- +goose StatementEnd
-- +goose StatementBegin
CREATE OR REPLACE FUNCTION on_stake_action_rwalk_insert() RETURNS trigger AS  $$
DECLARE
	v_cnt						NUMERIC;
BEGIN
	RETURN NEW;
END;
$$ LANGUAGE plpgsql;
-- +goose StatementEnd
-- +goose StatementBegin
CREATE OR REPLACE FUNCTION on_unstake_action_rwalk_delete() RETURNS trigger AS  $$
DECLARE
BEGIN

	RETURN OLD;
END;
$$ LANGUAGE plpgsql;
-- +goose StatementEnd
-- +goose StatementBegin
CREATE OR REPLACE FUNCTION on_nft_staked_cst_insert() RETURNS trigger AS  $$
DECLARE
	v_cnt						NUMERIC;
	v_active_stakers			INT;
	v_round_num					INT;
BEGIN

	INSERT INTO cg_staked_token_cst(staker_aid,token_id,stake_action_id)
		VALUES(NEW.staker_aid,NEW.token_id,NEW.action_id);
	INSERT INTO cg_staked_token_cst_rewards(staker_aid,token_id,stake_action_id)
		VALUES(NEW.staker_aid,NEW.token_id,NEW.action_id);
	UPDATE cg_staker_cst SET 
			total_tokens_staked = (total_tokens_staked + 1),
			num_stake_actions = (num_stake_actions + 1)
		WHERE staker_aid=NEW.staker_aid;
	GET DIAGNOSTICS v_cnt = ROW_COUNT;
	IF v_cnt = 0 THEN
		INSERT INTO cg_staker_cst(staker_aid,num_stake_actions,total_tokens_staked) VALUES(NEW.staker_aid,1,1);
	END IF;
	UPDATE cg_stake_stats_cst SET total_tokens_staked = (total_tokens_staked + 1);
	SELECT COUNT(*) FROM cg_staker_cst WHERE total_tokens_staked > 0 INTO v_active_stakers;
	IF v_active_stakers IS NOT NULL THEN
		UPDATE cg_stake_stats_cst SET total_num_stakers=v_active_stakers;
	END IF;
	SELECT round_num FROM cg_prize_claim ORDER BY round_num DESC LIMIT 1 INTO v_round_num;
	IF v_round_num IS NOT NULL THEN
		UPDATE cg_nft_staked_cst SET round_num=(v_round_num+1) WHERE id=NEW.id;
	END IF;
	RETURN NEW;
END;
$$ LANGUAGE plpgsql;
-- +goose StatementEnd
-- +goose StatementBegin
CREATE OR REPLACE FUNCTION on_nft_staked_cst_delete() RETURNS trigger AS  $$
DECLARE
	v_active_stakers			INT;
BEGIN

	DELETE FROM cg_staked_token_cst WHERE token_id = OLD.token_id;
	DELETE FROM cg_staked_token_cst_rewards WHERE token_id = OLD.token_id;
	UPDATE cg_staker_cst SET 
			total_tokens_staked = (total_tokens_staked - 1),
			num_stake_actions = (num_stake_actions - 1)
		WHERE staker_aid=OLD.staker_aid;
	UPDATE cg_stake_stats_cst SET total_tokens_staked = (total_tokens_staked - 1);
	SELECT COUNT(*) FROM cg_staker_cst WHERE total_tokens_staked > 0 INTO v_active_stakers;
	IF v_active_stakers IS NOT NULL THEN
		UPDATE cg_stake_stats_cst SET total_num_stakers=v_active_stakers;
	END IF;

	RETURN OLD;
END;
$$ LANGUAGE plpgsql;
-- +goose StatementEnd
-- +goose StatementBegin
CREATE OR REPLACE FUNCTION on_nft_staked_rwalk_insert() RETURNS trigger AS  $$
DECLARE
	v_cnt						NUMERIC;
	v_active_stakers			INT;
	v_round_num					INT;
BEGIN

	INSERT INTO cg_staked_token_rwalk(staker_aid,token_id,stake_action_id)
		VALUES(NEW.staker_aid,NEW.token_id,NEW.action_id);
	UPDATE cg_staker_rwalk SET 
			total_tokens_staked = (total_tokens_staked + 1),
			num_stake_actions = (num_stake_actions + 1)
		WHERE staker_aid=NEW.staker_aid;
	GET DIAGNOSTICS v_cnt = ROW_COUNT;
	IF v_cnt = 0 THEN
		INSERT INTO cg_staker_rwalk(staker_aid,total_tokens_staked,num_stake_actions) VALUES(NEW.staker_aid,1,1);
	END IF;
	UPDATE cg_stake_stats_rwalk SET total_tokens_staked = (total_tokens_staked + 1);
	SELECT COUNT(*) FROM cg_staker_rwalk WHERE total_tokens_staked > 0 INTO v_active_stakers;
	IF v_active_stakers IS NOT NULL THEN
		UPDATE cg_stake_stats_rwalk SET total_num_stakers=v_active_stakers;
	END IF;
	SELECT round_num FROM cg_prize_claim ORDER BY round_num DESC LIMIT 1 INTO v_round_num;
	IF v_round_num IS NOT NULL THEN
		UPDATE cg_nft_staked_rwalk SET round_num=(v_round_num+1) WHERE id=NEW.id;
	END IF;
	RETURN NEW;
END;
$$ LANGUAGE plpgsql;
-- +goose StatementEnd
-- +goose StatementBegin
CREATE OR REPLACE FUNCTION on_nft_staked_rwalk_delete() RETURNS trigger AS  $$
DECLARE
	v_active_stakers			INT;
BEGIN

	DELETE FROM cg_staked_token_rwalk WHERE token_id = OLD.token_id;
	UPDATE cg_staker_rwalk SET 
			total_tokens_staked = (total_tokens_staked - 1),
			num_stake_actions = (num_stake_actions - 1)
		WHERE staker_aid=OLD.staker_aid;
	UPDATE cg_stake_stats_rwalk SET total_tokens_staked = (total_tokens_staked - 1);
	SELECT COUNT(*) FROM cg_staker_rwalk WHERE total_tokens_staked > 0 INTO v_active_stakers;
	IF v_active_stakers IS NOT NULL THEN
		UPDATE cg_stake_stats_rwalk SET total_num_stakers=v_active_stakers;
	END IF;

	RETURN OLD;
END;
$$ LANGUAGE plpgsql;
-- +goose StatementEnd
-- +goose StatementBegin
CREATE OR REPLACE FUNCTION on_nft_unstaked_cst_insert() RETURNS trigger AS  $$
DECLARE
	v_cnt						NUMERIC;
	v_rec RECORD;
	v_active_stakers			INT;
	v_round_num					INT;
BEGIN

	UPDATE cg_staker_cst
		SET	total_tokens_staked = (total_tokens_staked - 1),
			num_unstake_actions = (num_unstake_actions + 1)
		WHERE staker_aid=NEW.staker_aid;
	UPDATE cg_stake_stats_cst SET total_tokens_staked = (total_tokens_staked - 1);

	FOR v_rec IN (SELECT action_id,deposit_id FROM cg_st_reward WHERE action_id=NEW.action_id ORDER BY deposit_id DESC,action_id DESC)
		LOOP
			UPDATE cg_st_reward 
				SET collected = 'T',
			    	is_unstake = 'T'
				WHERE deposit_id=v_rec.deposit_id AND action_id=v_rec.action_id;
		END LOOP;

	DELETE from cg_staked_token_cst WHERE token_id=NEW.token_id AND staker_aid=NEW.staker_aid;
	SELECT COUNT(*) FROM cg_staker_cst WHERE total_tokens_staked > 0 INTO v_active_stakers;
	IF v_active_stakers IS NOT NULL THEN
		UPDATE cg_stake_stats_cst SET total_num_stakers=v_active_stakers;
	END IF;
	SELECT round_num FROM cg_prize_claim ORDER BY round_num DESC LIMIT 1 INTO v_round_num;
	IF v_round_num IS NOT NULL THEN
		UPDATE cg_nft_unstaked_cst SET round_num=(v_round_num+1) WHERE id=NEW.id;
	END IF;
	RETURN NEW;
END;
$$ LANGUAGE plpgsql;
-- +goose StatementEnd
-- +goose StatementBegin
CREATE OR REPLACE FUNCTION on_nft_unstaked_cst_delete() RETURNS trigger AS  $$
DECLARE
	v_rec RECORD;
	v_active_stakers			INT;
BEGIN

	UPDATE cg_staker_cst
		SET total_tokens_staked = (total_tokens_staked + 1),
			num_unstake_actions = (num_unstake_actions - 1)
		WHERE staker_aid=OLD.staker_aid;
	UPDATE cg_stake_stats_cst SET total_tokens_staked = (total_tokens_staked + 1);

	FOR v_rec IN (SELECT action_id,deposit_id FROM cg_st_reward WHERE action_id=OLD.action_id ORDER BY deposit_id DESC,action_id DESC)
		LOOP
			UPDATE cg_st_reward
				SET collected = 'F',
			   		is_unstake = 'F'
				WHERE deposit_id=v_rec.deposit_id AND action_id=v_rec.action_id;
		END LOOP;
	SELECT COUNT(*) FROM cg_staker_cst WHERE total_tokens_staked > 0 INTO v_active_stakers;
	IF v_active_stakers IS NOT NULL THEN
		UPDATE cg_stake_stats_cst SET total_num_stakers=v_active_stakers;
	END IF;

	RETURN OLD;
END;
$$ LANGUAGE plpgsql;
-- +goose StatementEnd
-- +goose StatementBegin
CREATE OR REPLACE FUNCTION on_nft_unstaked_rwalk_insert() RETURNS trigger AS  $$
DECLARE
	v_cnt						NUMERIC;
	v_active_stakers			INT;
	v_round_num					INT;
BEGIN

	UPDATE cg_staker_rwalk
		SET	total_tokens_staked = (total_tokens_staked - 1),
			num_unstake_actions = (num_unstake_actions + 1)
		WHERE staker_aid=NEW.staker_aid;
	UPDATE cg_stake_stats_rwalk SET total_tokens_staked = (total_tokens_staked - 1);
	DELETE FROM cg_staked_token_rwalk WHERE token_id=NEW.token_id AND staker_aid=NEW.staker_aid;
	SELECT COUNT(*) FROM cg_staker_rwalk WHERE total_tokens_staked > 0 INTO v_active_stakers;
	IF v_active_stakers IS NOT NULL THEN
		UPDATE cg_stake_stats_rwalk SET total_num_stakers=v_active_stakers;
	END IF;
	SELECT round_num FROM cg_prize_claim ORDER BY round_num DESC LIMIT 1 INTO v_round_num;
	IF v_round_num IS NOT NULL THEN
		UPDATE cg_nft_unstaked_rwalk SET round_num=(v_round_num+1) WHERE id=NEW.id;
	END IF;
	RETURN NEW;
END;
$$ LANGUAGE plpgsql;
-- +goose StatementEnd
-- +goose StatementBegin
CREATE OR REPLACE FUNCTION on_nft_unstaked_rwalk_delete() RETURNS trigger AS  $$
DECLARE
	v_active_stakers			INT;
BEGIN

	UPDATE cg_staker_rwalk
		SET total_tokens_staked = (total_tokens_staked + 1),
			num_unstake_actions = (num_unstake_actions - 1)
		WHERE staker_aid=OLD.staker_aid;
	UPDATE cg_stake_stats_rwalk SET total_tokens_staked = (total_tokens_staked + 1);
	SELECT COUNT(*) FROM cg_staker_rwalk WHERE total_tokens_staked > 0 INTO v_active_stakers;
	IF v_active_stakers IS NOT NULL THEN
		UPDATE cg_stake_stats_rwalk SET total_num_stakers=v_active_stakers;
	END IF;
	-- We aren't restoring state here (To Do)

	RETURN OLD;
END;
$$ LANGUAGE plpgsql;
-- +goose StatementEnd
-- +goose StatementBegin
CREATE OR REPLACE FUNCTION on_st_reward_insert() RETURNS trigger AS  $$
DECLARE
BEGIN

	IF NEW.collected THEN
		UPDATE cg_staker_deposit 
			SET amount_to_claim = (amount_to_claim - NEW.reward)
			WHERE deposit_id=NEW.deposit_id AND staker_aid=NEW.staker_aid;
	END IF;
	RETURN NEW;
END;
$$ LANGUAGE plpgsql;
-- +goose StatementEnd
-- +goose StatementBegin
CREATE OR REPLACE FUNCTION on_st_reward_delete() RETURNS trigger AS  $$
DECLARE
BEGIN

	IF OLD.collected THEN
		UPDATE cg_staker_deposit 
			SET amount_to_claim = (amount_to_claim + OLD.reward)
			WHERE deposit_id=OLD.deposit_id AND staker_aid=OLD.staker_aid;
	END IF;
	RETURN OLD;
END;
$$ LANGUAGE plpgsql;
-- +goose StatementEnd
-- +goose StatementBegin
CREATE OR REPLACE FUNCTION on_st_reward_update() RETURNS trigger AS  $$
DECLARE
	v_cnt						NUMERIC;
BEGIN
	IF OLD.collected <> NEW.collected THEN
		-- executed when rewardPaid() method is called or unstake() is called
		IF NEW.reward = OLD.reward THEN
			IF NEW.collected THEN
				UPDATE cg_staker_deposit
					SET amount_to_claim=(amount_to_claim - NEW.reward)
					WHERE staker_aid=NEW.staker_aid AND deposit_id=NEW.deposit_id;
				UPDATE cg_staker_cst
					SET unclaimed_reward = (unclaimed_reward - NEW.reward)
					WHERE staker_aid=NEW.staker_aid;
				UPDATE cg_stake_stats_cst
					SET total_unclaimed_reward = (total_unclaimed_reward - NEW.reward);
			END IF;
			IF NOT NEW.collected THEN
				UPDATE cg_staker_deposit
					SET amount_to_claim=(amount_to_claim + NEW.reward)
					WHERE staker_aid=NEW.staker_aid AND deposit_id=NEW.deposit_id;
				UPDATE cg_staker_cst
					SET unclaimed_reward = (unclaimed_reward + NEW.reward)
					WHERE staker_aid=NEW.staker_aid;
				UPDATE cg_stake_stats_cst
					SET total_unclaimed_reward = (total_unclaimed_reward + NEW.reward);
			END IF;
		END IF;
	END IF;
	RETURN NEW;
END;
$$ LANGUAGE plpgsql;
-- +goose StatementEnd
-- +goose StatementBegin
CREATE OR REPLACE FUNCTION on_chrono_warrior_insert() RETURNS trigger AS  $$
DECLARE
	v_cnt						NUMERIC;
BEGIN

	UPDATE cg_glob_stats SET total_chrono_warrior_eth_deposits = (total_chrono_warrior_eth_deposits + NEW.eth_amount);
	UPDATE cg_glob_stats SET total_cst_given_in_prizes = (total_cst_given_in_prizes + NEW.cst_amount);
	
	-- Update round stats
	UPDATE cg_round_stats SET chrono_warrior_prize_eth = (chrono_warrior_prize_eth + NEW.eth_amount), total_cst_paid_in_prizes = (total_cst_paid_in_prizes + NEW.cst_amount), total_nfts_minted = (total_nfts_minted + 1) WHERE round_num = NEW.round_num;
	GET DIAGNOSTICS v_cnt = ROW_COUNT;
	IF v_cnt = 0 THEN
		INSERT INTO cg_round_stats(round_num, chrono_warrior_prize_eth, total_cst_paid_in_prizes, total_nfts_minted) VALUES (NEW.round_num, NEW.eth_amount, NEW.cst_amount, 1);
	END IF;

	-- Insert THREE records in cg_prize table for Chrono Warrior prizes
	-- 1) Chrono Warrior ETH (ptype=7)
	INSERT INTO cg_prize(round_num,winner_index,ptype) VALUES(NEW.round_num,NEW.winner_index,7);
	-- 2) Chrono Warrior CST (ptype=8)
	INSERT INTO cg_prize(round_num,winner_index,ptype) VALUES(NEW.round_num,NEW.winner_index,8);
	-- 3) Chrono Warrior CS NFT (ptype=9)
	INSERT INTO cg_prize(round_num,winner_index,ptype) VALUES(NEW.round_num,NEW.winner_index,9);
	
	-- Update winner: prizes_count (+3), prizes_sum (+ETH), erc20_count (+1), erc721_count (+1)
	UPDATE cg_winner 
		SET 
			prizes_count = (prizes_count + 3),
			prizes_sum = (prizes_sum + NEW.eth_amount),
			erc20_count = (erc20_count + 1),
			erc721_count = (erc721_count + 1)
		WHERE winner_aid = NEW.winner_aid;
	GET DIAGNOSTICS v_cnt = ROW_COUNT;
	IF v_cnt = 0 THEN
		INSERT INTO cg_winner(winner_aid, prizes_count, prizes_sum, erc20_count, erc721_count) 
			VALUES (NEW.winner_aid, 3, NEW.eth_amount, 1, 1);
	END IF;

	-- Credit chrono warrior ETH to raffle_winner_stats (shared pool for all withdrawable prizes)
	UPDATE cg_raffle_winner_stats
		SET amount_sum = (amount_sum + NEW.eth_amount)
		WHERE winner_aid = NEW.winner_aid;
	GET DIAGNOSTICS v_cnt = ROW_COUNT;
	IF v_cnt = 0 THEN
		INSERT INTO cg_raffle_winner_stats(winner_aid, amount_sum, raffles_count)
			VALUES(NEW.winner_aid, NEW.eth_amount, 0);
	END IF;

	RETURN NEW;
END;
$$ LANGUAGE plpgsql;
-- +goose StatementEnd
-- +goose StatementBegin
CREATE OR REPLACE FUNCTION on_chrono_warrior_delete() RETURNS trigger AS  $$
DECLARE
BEGIN

	UPDATE cg_glob_stats SET total_chrono_warrior_eth_deposits = (total_chrono_warrior_eth_deposits - OLD.eth_amount);
	UPDATE cg_glob_stats SET total_cst_given_in_prizes = (total_cst_given_in_prizes - OLD.cst_amount);
	
	-- Update round stats
	UPDATE cg_round_stats SET chrono_warrior_prize_eth = (chrono_warrior_prize_eth - OLD.eth_amount) WHERE round_num = OLD.round_num;
	UPDATE cg_round_stats SET total_cst_paid_in_prizes = (total_cst_paid_in_prizes - OLD.cst_amount) WHERE round_num = OLD.round_num;
	UPDATE cg_round_stats SET total_nfts_minted = (total_nfts_minted - 1) WHERE round_num = OLD.round_num;

	-- Remove THREE corresponding records from cg_prize table
	-- 1) Chrono Warrior ETH (ptype=7)
	DELETE FROM cg_prize WHERE round_num=OLD.round_num AND winner_index=OLD.winner_index AND ptype=7;
	-- 2) Chrono Warrior CST (ptype=8)
	DELETE FROM cg_prize WHERE round_num=OLD.round_num AND winner_index=OLD.winner_index AND ptype=8;
	-- 3) Chrono Warrior CS NFT (ptype=9)
	DELETE FROM cg_prize WHERE round_num=OLD.round_num AND winner_index=OLD.winner_index AND ptype=9;

	-- Update winner counts
	UPDATE cg_winner 
		SET 
			prizes_count = GREATEST(0, prizes_count - 3),
			prizes_sum = GREATEST(0, prizes_sum - OLD.eth_amount),
			erc20_count = GREATEST(0, erc20_count - 1),
			erc721_count = GREATEST(0, erc721_count - 1)
		WHERE winner_aid = OLD.winner_aid;

	-- Debit chrono warrior ETH from raffle_winner_stats
	UPDATE cg_raffle_winner_stats
		SET amount_sum = (amount_sum - OLD.eth_amount)
		WHERE winner_aid = OLD.winner_aid;

	RETURN OLD;
END;
$$ LANGUAGE plpgsql;
-- +goose StatementEnd
-- +goose StatementBegin
CREATE OR REPLACE FUNCTION on_raffle_eth_winner_insert() RETURNS trigger AS  $$
DECLARE
	v_cnt						NUMERIC;
BEGIN

	UPDATE cg_raffle_winner_stats
		SET
			amount_sum = (amount_sum + NEW.amount),
			raffles_count = (raffles_count + 1)
		WHERE winner_aid = NEW.winner_aid;
	GET DIAGNOSTICS v_cnt = ROW_COUNT;
	IF v_cnt = 0 THEN
		INSERT INTO cg_raffle_winner_stats(winner_aid,amount_sum,raffles_count) VALUES(NEW.winner_aid,NEW.amount,1);
	END IF;

	UPDATE cg_round_stats
		SET
			total_raffle_eth_deposits = (total_raffle_eth_deposits + NEW.amount)
		WHERE round_num=NEW.round_num;
	GET DIAGNOSTICS v_cnt = ROW_COUNT;
	IF v_cnt = 0 THEN
		INSERT INTO cg_round_stats(round_num,total_raffle_eth_deposits) VALUES(NEW.round_num,NEW.amount);
	END IF;

	UPDATE cg_glob_stats SET total_raffle_eth_deposits = (total_raffle_eth_deposits + NEW.amount);

	-- Insert record in cg_prize table with ptype=10 for Raffle ETH (for bidders)
	INSERT INTO cg_prize(round_num,winner_index,ptype) VALUES(NEW.round_num,NEW.winner_idx,10);
	
	-- Update winner prize count and prizes_sum (Raffle ETH = 1 prize, ETH amount)
	UPDATE cg_winner 
		SET 
			prizes_count = (prizes_count + 1),
			prizes_sum = (prizes_sum + NEW.amount)
		WHERE winner_aid = NEW.winner_aid;
	GET DIAGNOSTICS v_cnt = ROW_COUNT;
	IF v_cnt = 0 THEN
		INSERT INTO cg_winner(winner_aid, prizes_count, prizes_sum) VALUES (NEW.winner_aid, 1, NEW.amount);
	END IF;

	RETURN NEW;
END;
$$ LANGUAGE plpgsql;
-- +goose StatementEnd
-- +goose StatementBegin
CREATE OR REPLACE FUNCTION on_raffle_eth_winner_delete() RETURNS trigger AS  $$
DECLARE
BEGIN

	UPDATE cg_raffle_winner_stats
		SET
			amount_sum = (amount_sum - OLD.amount),
			raffles_count = (raffles_count - 1)
		WHERE winner_aid = OLD.winner_aid;

	UPDATE cg_round_stats
		SET
			total_raffle_eth_deposits = (total_raffle_eth_deposits - OLD.amount)
		WHERE round_num=OLD.round_num;

	UPDATE cg_glob_stats SET total_raffle_eth_deposits = (total_raffle_eth_deposits - OLD.amount);

	-- Remove corresponding record from cg_prize table
	DELETE FROM cg_prize WHERE round_num=OLD.round_num AND winner_index=OLD.winner_idx AND ptype=10;

	-- Update winner prize count and prizes_sum
	UPDATE cg_winner 
		SET 
			prizes_count = GREATEST(0, prizes_count - 1),
			prizes_sum = GREATEST(0, prizes_sum - OLD.amount)
		WHERE winner_aid = OLD.winner_aid;

	RETURN OLD;
END;
$$ LANGUAGE plpgsql;
-- +goose StatementEnd
-- +goose StatementBegin
CREATE OR REPLACE FUNCTION on_erc20_reward_insert() RETURNS trigger AS $$
BEGIN
	UPDATE cg_glob_stats SET cst_reward_for_bidding = NEW.new_reward;
	RETURN NEW;
END;
$$ LANGUAGE plpgsql;
-- +goose StatementEnd
CREATE TRIGGER erc20_reward_insert AFTER INSERT ON cg_adm_erc20_reward FOR EACH ROW EXECUTE FUNCTION on_erc20_reward_insert();
-- ============================================================================
-- Round Timing Trigger Functions
-- Purpose: Automatically track round lifecycle timing including parameter windows
-- Date: 2025-11-06
-- ============================================================================

-- Trigger function for cg_first_bid (FirstBidPlacedInRound event)
-- Sets round_start_time when the first bid is placed
-- +goose StatementBegin
CREATE OR REPLACE FUNCTION on_first_bid_update_round_timing() RETURNS trigger AS $$
BEGIN
	-- When FirstBidPlacedInRound event fires, set the round_start_time
	-- Use start_ts (Unix timestamp) instead of time_stamp which may be incorrectly populated
	UPDATE cg_round_stats
	SET round_start_time = TO_TIMESTAMP(NEW.start_ts)
	WHERE round_num = NEW.round_num;
	
	-- If round_stats doesn't exist yet, create it
	IF NOT FOUND THEN
		INSERT INTO cg_round_stats(round_num, round_start_time)
		VALUES (NEW.round_num, TO_TIMESTAMP(NEW.start_ts));
	END IF;
	
	RETURN NEW;
END;
$$ LANGUAGE plpgsql;
-- +goose StatementEnd

-- Trigger function for cg_prize_claim
-- Sets round_end_time and param_window_start_time for next round
-- +goose StatementBegin
CREATE OR REPLACE FUNCTION on_prize_claim_update_round_timing() RETURNS trigger AS $$
DECLARE
	v_start_time TIMESTAMPTZ;
	v_next_round_num BIGINT;
BEGIN
	-- Get the round start time for duration calculation
	SELECT round_start_time INTO v_start_time
	FROM cg_round_stats
	WHERE round_num = NEW.round_num;
	
	-- Update current round's end time and calculate duration
	UPDATE cg_round_stats
	SET 
		round_end_time = NEW.time_stamp,
		round_duration_seconds = CASE 
			WHEN round_start_time IS NOT NULL 
			THEN EXTRACT(EPOCH FROM (NEW.time_stamp - round_start_time))::BIGINT
			ELSE NULL
		END
	WHERE round_num = NEW.round_num;
	
	-- If round_stats doesn't exist yet, create it
	IF NOT FOUND THEN
		INSERT INTO cg_round_stats(round_num, round_end_time)
		VALUES (NEW.round_num, NEW.time_stamp);
	END IF;
	
	-- Set the param_window_start_time for the NEXT round
	-- (parameter setting window starts when current round ends)
	v_next_round_num := NEW.round_num + 1;
	
	UPDATE cg_round_stats
	SET param_window_start_time = NEW.time_stamp
	WHERE round_num = v_next_round_num;
	
	-- If next round doesn't exist yet, create it
	IF NOT FOUND THEN
		INSERT INTO cg_round_stats(round_num, param_window_start_time)
		VALUES (v_next_round_num, NEW.time_stamp);
	END IF;
	
	RETURN NEW;
END;
$$ LANGUAGE plpgsql;
-- +goose StatementEnd

-- Trigger function for cg_adm_acttime
-- Sets activation_time and calculates param_window_duration
-- +goose StatementBegin
CREATE OR REPLACE FUNCTION on_activation_time_set_update_round_timing() RETURNS trigger AS $$
DECLARE
	v_round_num BIGINT;
	v_param_start TIMESTAMPTZ;
BEGIN
	-- Round this activation time applies to: 0 when no prize has been claimed yet, else last claimed round + 1.
	SELECT COALESCE(MAX(round_num), -1) + 1 INTO v_round_num
	FROM cg_prize_claim;
	
	-- Get param window start time
	SELECT param_window_start_time INTO v_param_start
	FROM cg_round_stats
	WHERE round_num = v_round_num;
	
	-- Update the round with activation_time and calculate param window duration
	UPDATE cg_round_stats
	SET 
		activation_time = TO_TIMESTAMP(NEW.new_atime),
		param_window_duration_seconds = CASE
			WHEN param_window_start_time IS NOT NULL
			THEN EXTRACT(EPOCH FROM (TO_TIMESTAMP(NEW.new_atime) - param_window_start_time))::BIGINT
			ELSE NULL
		END
	WHERE round_num = v_round_num;
	
	-- If round_stats doesn't exist yet, create it
	IF NOT FOUND THEN
		INSERT INTO cg_round_stats(round_num, activation_time)
		VALUES (v_round_num, TO_TIMESTAMP(NEW.new_atime));
	END IF;
	
	RETURN NEW;
END;
$$ LANGUAGE plpgsql;
-- +goose StatementEnd

CREATE TRIGGER cg_bid_insert AFTER INSERT ON cg_bid FOR EACH ROW EXECUTE PROCEDURE on_bid_insert();
CREATE TRIGGER cg_bid_delete AFTER DELETE ON cg_bid FOR EACH ROW EXECUTE PROCEDURE on_bid_delete();
CREATE TRIGGER cg_prize_claim_insert AFTER INSERT ON cg_prize_claim FOR EACH ROW EXECUTE PROCEDURE on_prize_claim_insert();
CREATE TRIGGER cg_prize_claim_delete AFTER DELETE ON cg_prize_claim FOR EACH ROW EXECUTE PROCEDURE on_prize_claim_delete();
CREATE TRIGGER cg_donation_received_insert AFTER INSERT ON cg_donation_received FOR EACH ROW EXECUTE PROCEDURE on_donation_received_insert();
CREATE TRIGGER cg_donation_received_delete AFTER DELETE ON cg_donation_received FOR EACH ROW EXECUTE PROCEDURE on_donation_received_delete();
CREATE TRIGGER cg_erc20_donation_insert AFTER INSERT ON cg_erc20_donation FOR EACH ROW EXECUTE PROCEDURE on_erc20_donation_insert();
CREATE TRIGGER cg_erc20_donation_delete AFTER DELETE ON cg_erc20_donation FOR EACH ROW EXECUTE PROCEDURE on_erc20_donation_delete();
CREATE TRIGGER cg_nft_donation_insert AFTER INSERT ON cg_nft_donation FOR EACH ROW EXECUTE PROCEDURE on_nft_donation_insert();
CREATE TRIGGER cg_nft_donation_delete AFTER DELETE ON cg_nft_donation FOR EACH ROW EXECUTE PROCEDURE on_nft_donation_delete();
CREATE TRIGGER cg_prize_deposit_insert AFTER INSERT ON cg_prize_deposit FOR EACH ROW EXECUTE PROCEDURE on_prize_deposit_insert();
CREATE TRIGGER cg_prize_deposit_delete AFTER DELETE ON cg_prize_deposit FOR EACH ROW EXECUTE PROCEDURE on_prize_deposit_delete();
CREATE TRIGGER cg_prize_withdrawal_insert AFTER INSERT ON cg_prize_withdrawal FOR EACH ROW EXECUTE PROCEDURE on_prize_withdrawal_insert();
CREATE TRIGGER cg_prize_withdrawal_delete AFTER DELETE ON cg_prize_withdrawal FOR EACH ROW EXECUTE PROCEDURE on_prize_withdrawal_delete();
CREATE TRIGGER cg_raffle_nft_winner_insert AFTER INSERT ON cg_raffle_nft_prize FOR EACH ROW EXECUTE PROCEDURE on_raffle_nft_winner_insert();
CREATE TRIGGER cg_raffle_nft_winner_delete AFTER DELETE ON cg_raffle_nft_prize FOR EACH ROW EXECUTE PROCEDURE on_raffle_nft_winner_delete();
CREATE TRIGGER cg_raffle_eth_winner_insert AFTER INSERT ON cg_raffle_eth_prize FOR EACH ROW EXECUTE PROCEDURE on_raffle_eth_winner_insert();
CREATE TRIGGER cg_raffle_eth_winner_delete AFTER DELETE ON cg_raffle_eth_prize FOR EACH ROW EXECUTE PROCEDURE on_raffle_eth_winner_delete();
CREATE TRIGGER cg_endurance_winner_insert AFTER INSERT ON cg_endurance_prize FOR EACH ROW EXECUTE PROCEDURE on_endurance_winner_insert();
CREATE TRIGGER cg_endurance_winner_delete AFTER DELETE ON cg_endurance_prize FOR EACH ROW EXECUTE PROCEDURE on_endurance_winner_delete();
CREATE TRIGGER cg_chrono_warrior_insert AFTER INSERT ON cg_chrono_warrior_prize FOR EACH ROW EXECUTE PROCEDURE on_chrono_warrior_insert();
CREATE TRIGGER cg_chrono_warrior_delete AFTER DELETE ON cg_chrono_warrior_prize FOR EACH ROW EXECUTE PROCEDURE on_chrono_warrior_delete();
CREATE TRIGGER cg_lastcst_winner_insert AFTER INSERT ON cg_lastcst_prize FOR EACH ROW EXECUTE PROCEDURE on_lastcst_winner_insert();
CREATE TRIGGER cg_lastcst_winner_delete AFTER DELETE ON cg_lastcst_prize FOR EACH ROW EXECUTE PROCEDURE on_lastcst_winner_delete();
CREATE TRIGGER cg_transfer_insert AFTER INSERT ON cg_erc721_transfer FOR EACH ROW EXECUTE PROCEDURE on_erc721transfer_insert();
CREATE TRIGGER cg_transfer_delete AFTER DELETE ON cg_erc721_transfer FOR EACH ROW EXECUTE PROCEDURE on_erc721transfer_delete();
CREATE TRIGGER cg_erc20_transfer_insert AFTER INSERT ON cg_erc20_transfer FOR EACH ROW EXECUTE PROCEDURE on_erc20_transfer_insert();
CREATE TRIGGER cg_erc20_transfer_delete AFTER DELETE ON cg_erc20_transfer FOR EACH ROW EXECUTE PROCEDURE on_erc20_transfer_delete();
CREATE TRIGGER cg_donated_nft_claimed_insert AFTER INSERT ON cg_donated_nft_claimed FOR EACH ROW EXECUTE PROCEDURE on_donated_nft_claimed_insert();
CREATE TRIGGER cg_donated_nft_claimed_delete AFTER DELETE ON cg_donated_nft_claimed FOR EACH ROW EXECUTE PROCEDURE on_donated_nft_claimed_delete();
CREATE TRIGGER cg_donated_tok_claimed_insert AFTER INSERT ON cg_donated_tok_claimed FOR EACH ROW EXECUTE PROCEDURE on_donated_tok_claimed_insert();
CREATE TRIGGER cg_donated_tok_claimed_delete AFTER DELETE ON cg_donated_tok_claimed FOR EACH ROW EXECUTE PROCEDURE on_donated_tok_claimed_delete();
CREATE TRIGGER cg_mint_insert AFTER INSERT ON cg_mint_event FOR EACH ROW EXECUTE PROCEDURE on_mint_insert();
CREATE TRIGGER cg_mint_delete AFTER DELETE ON cg_mint_event FOR EACH ROW EXECUTE PROCEDURE on_mint_delete();
CREATE TRIGGER cg_mint_update AFTER UPDATE ON cg_mint_event FOR EACH ROW EXECUTE PROCEDURE on_mint_update();
CREATE TRIGGER cg_token_name_insert AFTER INSERT ON cg_token_name FOR EACH ROW EXECUTE PROCEDURE on_token_name_insert();
CREATE TRIGGER cg_token_name_delete AFTER DELETE ON cg_token_name FOR EACH ROW EXECUTE PROCEDURE on_token_name_delete();
CREATE TRIGGER cg_donation_sent_insert AFTER INSERT ON cg_donation_sent FOR EACH ROW EXECUTE PROCEDURE on_donation_sent_insert();
CREATE TRIGGER cg_donation_sent_delete AFTER DELETE ON cg_donation_sent FOR EACH ROW EXECUTE PROCEDURE on_donation_sent_delete();
CREATE TRIGGER cg_eth_deposit_insert AFTER INSERT ON cg_staking_eth_deposit FOR EACH ROW EXECUTE PROCEDURE on_eth_deposit_insert();
CREATE TRIGGER cg_eth_deposit_delete AFTER DELETE ON cg_staking_eth_deposit FOR EACH ROW EXECUTE PROCEDURE on_eth_deposit_delete();
CREATE TRIGGER cg_marketing_rewards_insert AFTER INSERT ON cg_mkt_reward FOR EACH ROW EXECUTE PROCEDURE on_marketing_rewards_insert();
CREATE TRIGGER cg_marketing_rewards_delete AFTER DELETE ON cg_mkt_reward FOR EACH ROW EXECUTE PROCEDURE on_marketing_rewards_delete();
CREATE TRIGGER cg_donation_insert AFTER INSERT ON cg_eth_donated FOR EACH ROW EXECUTE PROCEDURE on_direct_donation_insert();
CREATE TRIGGER cg_donation_delete AFTER DELETE ON cg_eth_donated FOR EACH ROW EXECUTE PROCEDURE on_direct_donation_delete();
CREATE TRIGGER cg_donation_wi_insert AFTER INSERT ON cg_eth_donated_wi FOR EACH ROW EXECUTE PROCEDURE on_direct_donation_insert();
CREATE TRIGGER cg_donation_wi_delete AFTER DELETE ON cg_eth_donated_wi FOR EACH ROW EXECUTE PROCEDURE on_direct_donation_delete();
CREATE TRIGGER cg_nft_staked_rwalk_insert AFTER INSERT ON cg_nft_staked_rwalk FOR EACH ROW EXECUTE PROCEDURE on_nft_staked_rwalk_insert();
CREATE TRIGGER cg_nft_staked_rwalk_delete AFTER DELETE ON cg_nft_staked_rwalk FOR EACH ROW EXECUTE PROCEDURE on_nft_staked_rwalk_delete();
CREATE TRIGGER cg_nft_staked_cst_insert AFTER INSERT ON cg_nft_staked_cst FOR EACH ROW EXECUTE PROCEDURE on_nft_staked_cst_insert();
CREATE TRIGGER cg_nft_staked_cst_delete AFTER DELETE ON cg_nft_staked_cst FOR EACH ROW EXECUTE PROCEDURE on_nft_staked_cst_delete();
CREATE TRIGGER cg_nft_unstaked_cst_insert AFTER INSERT ON cg_nft_unstaked_cst FOR EACH ROW EXECUTE PROCEDURE on_nft_unstaked_cst_insert();
CREATE TRIGGER cg_nft_unstaked_cst_delete AFTER DELETE ON cg_nft_unstaked_cst FOR EACH ROW EXECUTE PROCEDURE on_nft_unstaked_cst_delete();
CREATE TRIGGER cg_nft_unstaked_rwalk_insert AFTER INSERT ON cg_nft_unstaked_rwalk FOR EACH ROW EXECUTE PROCEDURE on_nft_unstaked_rwalk_insert();
CREATE TRIGGER cg_nft_unstaked_rwalk_delete AFTER DELETE ON cg_nft_unstaked_rwalk FOR EACH ROW EXECUTE PROCEDURE on_nft_unstaked_rwalk_delete();
CREATE TRIGGER cg_st_reward_insert AFTER INSERT ON cg_st_reward FOR EACH ROW EXECUTE PROCEDURE on_st_reward_insert();
CREATE TRIGGER cg_st_reward_delete AFTER DELETE ON cg_st_reward FOR EACH ROW EXECUTE PROCEDURE on_st_reward_delete();
CREATE TRIGGER cg_st_reward_update AFTER UPDATE ON cg_st_reward FOR EACH ROW EXECUTE PROCEDURE on_st_reward_update();
-- Round timing triggers (added 2025-11-06)
CREATE TRIGGER trigger_first_bid_update_round_timing AFTER INSERT ON cg_first_bid FOR EACH ROW EXECUTE PROCEDURE on_first_bid_update_round_timing();
CREATE TRIGGER trigger_prize_claim_update_round_timing AFTER INSERT ON cg_prize_claim FOR EACH ROW EXECUTE PROCEDURE on_prize_claim_update_round_timing();
CREATE TRIGGER trigger_activation_time_update_round_timing AFTER INSERT ON cg_adm_acttime FOR EACH ROW EXECUTE PROCEDURE on_activation_time_set_update_round_timing();

CREATE INDEX bid_bidder_aid_idx			ON cg_bid			(bidder_aid);
CREATE INDEX prize_winner_aid_idx		ON cg_prize_claim	(winner_aid);
CREATE INDEX prize_num_idx				ON cg_prize_claim	(round_num);
CREATE INDEX mint_tokid_idx				ON cg_mint_event	(token_id);

-- +goose Down

DROP TABLE IF EXISTS
	cg_prize_claim,
	cg_prize,
	cg_bid,
	cg_eth_donated,
	cg_eth_donated_wi,
	cg_donation_json,
	cg_donation_received,
	cg_donation_sent,
	cg_erc20_donation,
	cg_nft_donation,
	cg_charity_receiver_changed,
	cg_token_name,
	cg_mint_event,
	cg_prize_deposit,
	cg_prize_withdrawal,
	cg_raffle_nft_prize,
	cg_raffle_eth_prize,
	cg_endurance_prize,
	cg_lastcst_prize,
	cg_donated_nft_claimed,
	cg_donated_tok_claimed,
	cg_nft_staked_cst,
	cg_nft_staked_rwalk,
	cg_nft_unstaked_cst,
	cg_nft_unstaked_rwalk,
	cg_staking_eth_deposit,
	cg_first_bid,
	cg_mkt_reward,
	cg_erc721_transfer,
	cg_erc20_transfer,
	cg_costok_owner,
	cg_adm_charity_pcent,
	cg_adm_treasurer_addr,
	cg_adm_main_prize_pcent,
	cg_adm_raffle_pcent,
	cg_adm_stake_pcent,
	cg_adm_raf_eth_bidding,
	cg_adm_raf_nft_bidding,
	cg_adm_raf_nft_staking_rwalk,
	cg_adm_charity_wallet,
	cg_adm_rwalk_addr,
	cg_adm_prizes_wallet_addr,
	cg_adm_prize_microsec,
	cg_adm_staking_cst_addr,
	cg_adm_staking_rwalk_addr,
	cg_adm_marketing_addr,
	cg_adm_costok_addr,
	cg_adm_cossig_addr,
	cg_adm_upgraded,
	cg_adm_time_inc,
	cg_adm_admin_changed,
	cg_adm_timeout_claimprize,
	cg_adm_timeout_withdraw,
	cg_adm_price_inc,
	cg_adm_inisecprize,
	cg_adm_acttime,
	cg_adm_cst_auclen,
	cg_adm_cst_auclen_chg_div,
	cg_adm_eth_auclen,
	cg_adm_eth_auc_endprice,
	cg_adm_erc_rwd_mul,
	cg_adm_mkt_reward,
	cg_adm_erc20_reward,
	cg_adm_msg_len,
	cg_adm_script_url,
	cg_adm_base_uri_cs,
	cg_adm_ownership,
	cg_adm_initialized,
	cg_fund_transf_err,
	cg_erc20_transf_err,
	cg_funds_to_charity,
	cg_adm_chrono_pcent,
	cg_adm_cst_min_limit,
	cg_delay_duration,
	cg_chrono_warrior_prize,
	cg_banned_bids,
	cg_transfer_stats,
	cg_round_stats,
	cg_bidder,
	cg_winner,
	cg_donor,
	cg_staker_cst,
	cg_staker_rwalk,
	cg_staker_deposit,
	cg_staked_token_cst,
	cg_staked_token_cst_rewards,
	cg_staked_token_rwalk,
	cg_st_reward,
	cg_raffle_winner_stats,
	cg_erc20_donation_stats,
	cg_raffle_nft_winner_stats,
	cg_glob_stats,
	cg_nft_stats,
	cg_stake_stats_cst,
	cg_stake_stats_rwalk,
	cg_contracts,
	cg_proc_status
CASCADE;

