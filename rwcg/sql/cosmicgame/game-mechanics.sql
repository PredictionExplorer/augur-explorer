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
