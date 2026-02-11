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
	tokens_minted	DECIMAL DEFAULT 0 -- total tokens minted
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
	cst_reward_for_bidding DECIMAL DEFAULT 0,
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
