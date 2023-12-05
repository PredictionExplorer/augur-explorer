CREATE TABLE cg_prize_claim(
	id						BIGSERIAL PRIMARY KEY,
	evtlog_id				BIGINT REFERENCES evt_log(id) ON DELETE CASCADE,
	block_num				BIGINT NOT NULL,
	tx_id					BIGINT NOT NULL,
	time_stamp				TIMESTAMPTZ NOT NULL,
	contract_aid			BIGINT NOT NULL,
	prize_num				BIGINT NOT NULL,
	winner_aid				BIGINT NOT NULL,
	token_id				BIGINT NOT NULL,
	amount					DECIMAL DEFAULT 0,
	donation_evt_id			BIGINT NOT NULL,
	UNIQUE(evtlog_id)
);
CREATE TABLE cg_bid (
	id				BIGSERIAL PRIMARY KEY,
	evtlog_id		BIGINT REFERENCES evt_log(id) ON DELETE CASCADE,
	block_num		BIGINT NOT NULL,
	tx_id			BIGINT NOT NULL,
	time_stamp		TIMESTAMPTZ NOT NULL,
	contract_aid	BIGINT NOT NULL,
	bidder_aid		BIGINT NOT NULL,
	rwalk_nft_id	BIGINT NOT NULL,	--token_id of RandomWalk, if present
	round_num		BIGINT NOT NULL,
	bid_type		SMALLINT NOT NULL,  -- ETH, RandomWalk, CST
	num_cst_tokens	BIGINT NOT NULL,
	prize_time		TIMESTAMPTZ NOT NULL,
	bid_price		DECIMAL NOT NULL,
	erc20_amount	DECIMAL DEFAULT 0,	-- amount of CosmicSignatureToken minted in ERC20
	msg				TEXT,
	UNIQUE(evtlog_id)
);
CREATE TABLE cg_donation (
	id				BIGSERIAL PRIMARY KEY,
	evtlog_id		BIGINT REFERENCES evt_log(id) ON DELETE CASCADE,
	block_num		BIGINT NOT NULL,
	tx_id			BIGINT NOT NULL,
	time_stamp		TIMESTAMPTZ NOT NULL,
	contract_aid	BIGINT NOT NULL,
	donor_aid		BIGINT NOT NULL,
	amount			DECIMAL NOT NULL,
	UNIQUE(evtlog_id)
);
CREATE TABLE cg_donation_received (
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
CREATE TABLE cg_donation_sent (
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
CREATE TABLE cg_nft_donation (
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
CREATE TABLE cg_charity_updated (
	id				BIGSERIAL PRIMARY KEY,
	evtlog_id		BIGINT REFERENCES evt_log(id) ON DELETE CASCADE,
	block_num		BIGINT NOT NULL,
	tx_id			BIGINT NOT NULL,
	time_stamp		TIMESTAMPTZ NOT NULL,
	contract_aid	BIGINT NOT NULL,
	charity_aid		BIGINT NOT NULL,
	UNIQUE(evtlog_id)
);
CREATE TABLE cg_token_name (
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
CREATE TABLE cg_mint_event (
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
	staked			BOOLEAN DEFAULT 'F',
	UNIQUE(evtlog_id)
);
CREATE TABLE cg_raffle_deposit (
	id				BIGSERIAL PRIMARY KEY,
	evtlog_id		BIGINT REFERENCES evt_log(id) ON DELETE CASCADE,
	block_num		BIGINT NOT NULL,
	tx_id			BIGINT NOT NULL,
	time_stamp		TIMESTAMPTZ NOT NULL,
	contract_aid	BIGINT NOT NULL,
	winner_aid		BIGINT NOT NULL,
	round_num		BIGINT NOT NULL,
	amount			DECIMAL NOT NULL,
	claimed			BOOLEAN DEFAULT 'F',	-- upon withdrawal is set to TRUE
	withdrawal_id	BIGINT DEFAULT 0, -- at withdrawal set to evtlog_id of bw_raffle_Withdrawal
	UNIQUE(evtlog_id)
);
CREATE TABLE cg_raffle_withdrawal (
	id				BIGSERIAL PRIMARY KEY,
	evtlog_id		BIGINT REFERENCES evt_log(id) ON DELETE CASCADE,
	block_num		BIGINT NOT NULL,
	tx_id			BIGINT NOT NULL,
	time_stamp		TIMESTAMPTZ NOT NULL,
	contract_aid	BIGINT NOT NULL,
	winner_aid		BIGINT NOT NULL,
	amount			DECIMAL NOT NULL,
	UNIQUE(evtlog_id)
);
CREATE TABLE cg_raffle_nft_winner (
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
	UNIQUE(evtlog_id)
);
CREATE TABLE cg_donated_nft_claimed (
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
CREATE TABLE cg_stake_action (
	id				BIGSERIAL PRIMARY KEY,
	evtlog_id		BIGINT REFERENCES evt_log(id) ON DELETE CASCADE,
	block_num		BIGINT NOT NULL,
	tx_id			BIGINT NOT NULL,
	time_stamp		TIMESTAMPTZ NOT NULL,
	contract_aid	BIGINT NOT NULL,
	round_num		BIGINT DEFAULT -1,
	action_id		BIGINT NOT NULL,
	token_id		BIGINT NOT NULL,
	num_staked_nfts	BIGINT NOT NULL,
	unstake_time	TIMESTAMPTZ NOT NULL,
	staker_aid		BIGINT NOT NULL,
	claimed			BOOLEAN DEFAULT 'F',
	UNIQUE(evtlog_id)
);
CREATE TABLE cg_unstake_action (
	id				BIGSERIAL PRIMARY KEY,
	evtlog_id		BIGINT REFERENCES evt_log(id) ON DELETE CASCADE,
	block_num		BIGINT NOT NULL,
	tx_id			BIGINT NOT NULL,
	time_stamp		TIMESTAMPTZ NOT NULL,
	contract_aid	BIGINT NOT NULL,
	round_num		BIGINT DEFAULT -1,
	action_id		BIGINT NOT NULL,
	token_id		BIGINT NOT NULL,
	num_staked_nfts	BIGINT NOT NULL,
	staker_aid		BIGINT NOT NULL,
	UNIQUE(evtlog_id)
);
CREATE TABLE cg_eth_deposit (
	id				BIGSERIAL PRIMARY KEY,
	evtlog_id		BIGINT REFERENCES evt_log(id) ON DELETE CASCADE,
	block_num		BIGINT NOT NULL,
	tx_id			BIGINT NOT NULL,
	time_stamp		TIMESTAMPTZ NOT NULL,
	contract_aid	BIGINT NOT NULL,
	round_num		BIGINT DEFAULT -1 ,
	deposit_time	TIMESTAMPTZ NOT NULL,
	deposit_num		BIGINT NOT NULL,
	num_staked_nfts	BIGINT NOT NULL,
	amount			DECIMAL NOT NULL,
	modulo			DECIMAL NOT NULL,
	UNIQUE(evtlog_id)
);
CREATE TABLE cg_claim_reward (
	id				BIGSERIAL PRIMARY KEY,
	evtlog_id		BIGINT REFERENCES evt_log(id) ON DELETE CASCADE,
	block_num		BIGINT NOT NULL,
	tx_id			BIGINT NOT NULL,
	time_stamp		TIMESTAMPTZ NOT NULL,
	contract_aid	BIGINT NOT NULL,
	action_id		BIGINT NOT NULL,
	deposit_id		BIGINT NOT NULL,
	reward			DECIMAL NOT NULL,
	staker_aid		BIGINT NOT NULL,
	UNIQUE(evtlog_id)
);
CREATE TABLE cg_transfer( -- cosmic signature ERC721 transfer
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
CREATE TABLE cg_erc20_transfer( -- cosmic token ERC20 transfer
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
CREATE TABLE cg_adm_charity_pcent( -- CharityPercentageChanged event
	id              BIGSERIAL PRIMARY KEY,
	evtlog_id       BIGINT REFERENCES evt_log(id) ON DELETE CASCADE,
	block_num       BIGINT NOT NULL,
	tx_id           BIGINT NOT NULL,
	time_stamp      TIMESTAMPTZ NOT NULL,
	contract_aid    BIGINT NOT NULL,
	percentage		DECIMAL NOT NULL,
	UNIQUE(evtlog_id)
);
CREATE TABLE cg_adm_prize_pcent( -- PrizePercentageChanged event
	id              BIGSERIAL PRIMARY KEY,
	evtlog_id       BIGINT REFERENCES evt_log(id) ON DELETE CASCADE,
	block_num       BIGINT NOT NULL,
	tx_id           BIGINT NOT NULL,
	time_stamp      TIMESTAMPTZ NOT NULL,
	contract_aid    BIGINT NOT NULL,
	percentage		DECIMAL NOT NULL,
	UNIQUE(evtlog_id)
);
CREATE TABLE cg_adm_raffle_pcent( -- RafflePercentageChanged event
	id              BIGSERIAL PRIMARY KEY,
	evtlog_id       BIGINT REFERENCES evt_log(id) ON DELETE CASCADE,
	block_num       BIGINT NOT NULL,
	tx_id           BIGINT NOT NULL,
	time_stamp      TIMESTAMPTZ NOT NULL,
	contract_aid    BIGINT NOT NULL,
	percentage		DECIMAL NOT NULL,
	UNIQUE(evtlog_id)
);
CREATE TABLE cg_adm_raf_eth_winners( -- NumRaffleWinnersPerRoundChanged event event
	id              BIGSERIAL PRIMARY KEY,
	evtlog_id       BIGINT REFERENCES evt_log(id) ON DELETE CASCADE,
	block_num       BIGINT NOT NULL,
	tx_id           BIGINT NOT NULL,
	time_stamp      TIMESTAMPTZ NOT NULL,
	contract_aid    BIGINT NOT NULL,
	num_winners		DECIMAL NOT NULL,	-- newNumRaffleWinnersPerRound
	UNIQUE(evtlog_id)
);
CREATE TABLE cg_transfer_stats( -- table to keep tracking of the statistical counters for tokent transfers
    user_aid                BIGINT NOT NULL,
    erc20_num_transfers     BIGINT DEFAULT 0, -- CosmicToken
    erc721_num_transfers    BIGINT DEFAULT 0  -- CosmicSignature
);
CREATE TABLE cg_round_stats( -- collects statistics per round 
	round_num					BIGINT NOT NULL PRIMARY KEY,
	total_bids					BIGINT DEFAULT 0,
	total_nft_donated			BIGINT DEFAULT 0,
	total_raffle_eth_deposits	DECIMAL DEFAULT 0,
	total_raffle_nfts			BIGINT DEFAULT 0
);
CREATE TABLE cg_bidder ( -- collects statistics per bidder
	bidder_aid		BIGINT PRIMARY KEY,
	num_bids		BIGINT DEFAULT 0,
	max_bid			DECIMAL DEFAULT 0,
	tokens_minted	DECIMAL DEFAULT 0 -- total tokens minted
);
CREATE TABLE cg_winner ( -- collects statistics per winer of prize
	winner_aid				BIGINT PRIMARY KEY,
	max_win_amount			DECIMAL DEFAULT 0,
	prizes_count			BIGINT DEFAULT 0,
	prizes_sum				DECIMAL DEFAULT 0,
	tokens_count			BIGINT DEFAULT 0,	-- tokens won in prizes + raffles
	unclaimed_nfts			BIGINT DEFAULT 0	-- donated NFTs
);
CREATE TABLE cg_staker ( -- counts statistics per user for StakingWallet actions
	staker_aid				BIGINT PRIMARY KEY,
	total_tokens_staked		BIGINT DEFAULT 0,
	num_stake_actions		BIGINT DEFAULT 0,
	num_unstake_actions		BIGINT DEFAULT 0,
	total_reward			DECIMAL DEFAULT 0,
	unclaimed_reward		DECIMAL DEFAULT 0
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
	sum_withdrawals			DECIMAL DEFAULT 0,		-- sum of withdrawals from CharityWallet to recipient
	num_withdrawals			BIGINT DEFAULT 0,
	num_bids				BIGINT DEFAULT 0, 		-- total bids made
	num_wins				BIGINT DEFAULT 0,		-- total prizes given
	num_rwalk_used			BIGINT DEFAULT 0,
	num_mints				BIGINT DEFAULT 0,
	cur_num_bids			BIGINT DEFAULT 0,		-- num bids since new round
	total_raffle_eth_deposits DECIMAL DEFAULT 0,
	total_raffle_eth_withdrawn DECIMAL DEFAULT 0,
	total_nft_donated		BIGINT DEFAULT 0
);
CREATE TABLE cg_nft_stats ( -- stats for donated NFTs (donated with bidAndDonateNFT())
	contract_aid			BIGINT PRIMARY KEY,
	num_donated				BIGINT DEFAULT 0		-- how many NFTs were donated
);
CREATE TABLE cg_stake_stats ( -- gloal staking statistics (StakinWallet)
	total_tokens_staked		BIGINT DEFAULT 0,
	total_reward_amount		DECIMAL DEFAULT 0,
	total_unclaimed_reward	DECIMAL DEFAULT 0,
	total_num_stakers		BIGINT DEFAULT 0,
	num_deposits			BIGINT DEFAULT 0,
	total_modulo			DECIMAL DEFAULT 0
);
CREATE TABLE cg_contracts (
	cosmic_game_addr		TEXT NOT NULL,
	cosmic_signature_addr	TEXT NOT NULL,
	cosmic_token_addr		TEXT NOT NULL,
	cosmic_dao_addr			TEXT NOT NULL,
	charity_wallet_addr		TEXT NOT NULL,
	raffle_wallet_addr		TEXT NOT NULL,
	random_walk_addr		TEXT NOT NULL,
	staking_wallet_addr		TEXT NOT NULL,
	marketing_wallet_addr	TEXT NOT NULL
);
CREATE TABLE cg_proc_status (
	last_evt_id             BIGINT DEFAULT 0
);
INSERT INTO cg_glob_stats DEFAULT VALUES;
