CREATE TABLE bw_prize_claim(
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
CREATE TABLE bw_bid (
	id				BIGSERIAL PRIMARY KEY,
	evtlog_id		BIGINT REFERENCES evt_log(id) ON DELETE CASCADE,
	block_num		BIGINT NOT NULL,
	tx_id			BIGINT NOT NULL,
	time_stamp		TIMESTAMPTZ NOT NULL,
	contract_aid	BIGINT NOT NULL,
	bidder_aid		BIGINT NOT NULL,
	rwalk_nft_id	BIGINT NOT NULL,	--token_id of RandomWalk, if present
	round_num		BIGINT NOT NULL,
	prize_time		TIMESTAMPTZ NOT NULL,
	bid_price		DECIMAL NOT NULL,
	erc20_amount	DECIMAL DEFAULT 0,	-- amount of CosmicSignatureToken minted in ERC20
	msg				TEXT,
	UNIQUE(evtlog_id)
);
CREATE TABLE bw_donation (
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
CREATE TABLE bw_donation_received (
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
CREATE TABLE bw_donation_sent (
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
CREATE TABLE bw_nft_donation (
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
	bid_id			BIGINT NOT NULL,		-- id of the related `bw_bid` record
	token_uri		TEXT NOT NULL,
	UNIQUE(evtlog_id)
);
CREATE TABLE bw_charity_updated (
	id				BIGSERIAL PRIMARY KEY,
	evtlog_id		BIGINT REFERENCES evt_log(id) ON DELETE CASCADE,
	block_num		BIGINT NOT NULL,
	tx_id			BIGINT NOT NULL,
	time_stamp		TIMESTAMPTZ NOT NULL,
	contract_aid	BIGINT NOT NULL,
	charity_aid		BIGINT NOT NULL,
	UNIQUE(evtlog_id)
);
CREATE TABLE bw_token_name (
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
CREATE TABLE bw_mint_event (
	id				BIGSERIAL PRIMARY KEY,
	evtlog_id		BIGINT REFERENCES evt_log(id) ON DELETE CASCADE,
	block_num		BIGINT NOT NULL,
	tx_id			BIGINT NOT NULL,
	time_stamp		TIMESTAMPTZ NOT NULL,
	contract_aid	BIGINT NOT NULL,
	owner_aid		BIGINT NOT NULL,
	token_id		BIGINT NOT NULL,
	cur_owner_aid	BIGINT NOT NULL,
	seed			TEXT NOT NULL,
	UNIQUE(evtlog_id)
);
CREATE TABLE bw_raffle_deposit (
	id				BIGSERIAL PRIMARY KEY,
	evtlog_id		BIGINT REFERENCES evt_log(id) ON DELETE CASCADE,
	block_num		BIGINT NOT NULL,
	tx_id			BIGINT NOT NULL,
	time_stamp		TIMESTAMPTZ NOT NULL,
	contract_aid	BIGINT NOT NULL,
	winner_aid		BIGINT NOT NULL,
	round_num		BIGINT NOT NULL,
	amount			DECIMAL NOT NULL,
	UNIQUE(evtlog_id)
);
CREATE TABLE bw_raffle_withdrawal (
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
CREATE TABLE bw_raffle_nft_winner (
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
CREATE TABLE bw_donated_nft_claimed (
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
CREATE TABLE bw_transfer( -- cosmic signature ERC721 transfer
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
CREATE TABLE bw_round_stats( -- collects statistics per round 
	round_num					BIGINT NOT NULL PRIMARY KEY,
	total_bids					BIGINT DEFAULT 0,
	total_nft_donated			BIGINT DEFAULT 0,
	total_raffle_eth_deposits	DECIMAL DEFAULT 0,
	total_raffle_nfts			BIGINT DEFAULT 0
);
CREATE TABLE bw_bidder ( -- collects statistics per bidder
	bidder_aid		BIGINT PRIMARY KEY,
	num_bids		BIGINT DEFAULT 0,
	max_bid			DECIMAL DEFAULT 0,
	tokens_minted	DECIMAL DEFAULT 0 -- total tokens minted
);
CREATE TABLE bw_winner ( -- collects statistics per winer of prize
	winner_aid				BIGINT PRIMARY KEY,
	max_win_amount			DECIMAL DEFAULT 0,
	prizes_count			BIGINT DEFAULT 0,
	prizes_sum				DECIMAL DEFAULT 0,
	tokens_count			BIGINT DEFAULT 0,	-- tokens won in prizes + raffles
	unclaimed_nfts			BIGINT DEFAULT 0	-- donated NFTs
);
CREATE TABLE bw_raffle_winner_stats (	-- prizes in ETH
	winner_aid		BIGINT PRIMARY KEY,
	amount_sum		DECIMAL DEFAULT 0,
	withdrawal_sum	DECIMAL DEFAULT 0,
	raffles_count	BIGINT DEFAULT 0
);
CREATE TABLE bw_raffle_nft_winner_stats ( -- prizes in NFT
	winner_aid		BIGINT PRIMARY KEY,
	num_won			BIGINT DEFAULT 0	-- num tokens won
);
CREATE TABLE bw_glob_stats ( -- global statistics
	num_vol_donations		BIGINT DEFAULT 0,		-- total number of voluntary donations
	vol_donations_total		DECIMAL DEFAULT 0,		-- sum of voluntary donations
	num_bids				BIGINT DEFAULT 0, 		-- total bids made
	num_wins				BIGINT DEFAULT 0,		-- total prizes given
	num_rwalk_used			BIGINT DEFAULT 0,
	num_mints				BIGINT DEFAULT 0,
	cur_num_bids			BIGINT DEFAULT 0		-- num bids since new round
);
CREATE TABLE bw_nft_stats ( -- stats for donated NFTs (donated with bidAndDonateNFT())
	contract_aid			BIGINT PRIMARY KEY,
	num_donated				BIGINT DEFAULT 0		-- how many NFTs were donated
);
CREATE TABLE bw_contracts (
	cosmic_game_addr		TEXT,
	cosmic_signature_addr	TEXT,
	cosmic_token_addr		TEXT,
	cosmic_dao_addr			TEXT,
	charity_wallet_addr		TEXT,
	raffle_wallet_addr		TEXT,
	random_walk_addr		TEXT
);
CREATE TABLE bw_proc_status (
	last_evt_id             BIGINT DEFAULT 0
);
INSERT INTO bw_glob_stats DEFAULT VALUES;
