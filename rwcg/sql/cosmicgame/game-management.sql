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
