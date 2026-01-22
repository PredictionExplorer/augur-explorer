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
	winner_idx		BIGINT NOT NULL,
	UNIQUE(round_num),
	UNIQUE(evtlog_id)
);
CREATE TABLE cg_lastcst_prize ( -- ICosmicSignatureGame.sol:LastCstBidderPrizePaid
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
	winner_idx		BIGINT NOT NULL,
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
CREATE TABLE cg_adm_cst_auclen ( -- ISystemEvents.sol:CstDutchAuctionDurationDivisorChanged
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
CREATE TABLE cg_adm_erc20_reward ( -- ISystemEvents.sol:CstRewardAmountForBiddingChanged
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
INSERT INTO cg_glob_stats DEFAULT VALUES;
INSERT INTO cg_stake_stats_cst DEFAULT VALUES;
INSERT INTO cg_stake_stats_rwalk DEFAULT VALUES;
