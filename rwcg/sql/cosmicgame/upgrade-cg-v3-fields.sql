-- Apply on existing databases before running the ETL with CosmicSignatureGameV3 support.
-- The ETL routes MainPrizeClaimed by event topic (V2 0x8c551ec2…, V3 0x9314e785…) and populates the
-- new split bid-reward rows for every bid it processes; no upgrade-block bifurcation is needed.

-- 1. Multi-NFT main prize (V3 mints num_cs_nfts Cosmic Signature NFTs to the winner; V2 = 1).
ALTER TABLE cg_prize_claim ADD COLUMN IF NOT EXISTS num_cs_nfts BIGINT NOT NULL DEFAULT 1;

-- 1b. V3 champion durations (championDurations[roundNum]); fetched by the ETL via eth_call
-- (at V3 claim processing and via startup backfill). 0 = not yet fetched / pre-V3 round.
ALTER TABLE cg_round_stats ADD COLUMN IF NOT EXISTS endurance_champion_duration BIGINT DEFAULT 0;
ALTER TABLE cg_round_stats ADD COLUMN IF NOT EXISTS chrono_warrior_duration BIGINT DEFAULT 0;

-- 2. Bid CST reward 90/10 split table (Comment-202607161).
CREATE TABLE IF NOT EXISTS cg_bid_reward (
	id				BIGSERIAL PRIMARY KEY,
	evtlog_id		BIGINT REFERENCES evt_log(id) ON DELETE CASCADE,
	bid_id			BIGINT NOT NULL,
	round_num		BIGINT NOT NULL,
	recipient_aid	BIGINT NOT NULL,
	reward_type		SMALLINT NOT NULL,	-- 0 = new (this) bidder; 1 = outbid (previous) bidder
	amount			DECIMAL NOT NULL DEFAULT 0,
	UNIQUE(bid_id,reward_type)
);

-- Backfill: pre-existing (V2/V1) bids credited the whole reward to the bidder placing the bid.
INSERT INTO cg_bid_reward(evtlog_id,bid_id,round_num,recipient_aid,reward_type,amount)
	SELECT b.evtlog_id, b.id, b.round_num, b.bidder_aid, 0, GREATEST(COALESCE(b.cst_reward,0),0)
	FROM cg_bid b
	WHERE NOT EXISTS (SELECT 1 FROM cg_bid_reward r WHERE r.bid_id=b.id AND r.reward_type=0);

-- 3. V3 configuration-changed events (ISystemEventsV3).
CREATE TABLE IF NOT EXISTS cg_adm_late_bid_dur_divisor(
	id BIGSERIAL PRIMARY KEY, evtlog_id BIGINT REFERENCES evt_log(id) ON DELETE CASCADE,
	block_num BIGINT NOT NULL, tx_id BIGINT NOT NULL, time_stamp TIMESTAMPTZ NOT NULL,
	contract_aid BIGINT NOT NULL, new_value DECIMAL NOT NULL, UNIQUE(evtlog_id));
CREATE TABLE IF NOT EXISTS cg_adm_late_bid_premium_base_mul(
	id BIGSERIAL PRIMARY KEY, evtlog_id BIGINT REFERENCES evt_log(id) ON DELETE CASCADE,
	block_num BIGINT NOT NULL, tx_id BIGINT NOT NULL, time_stamp TIMESTAMPTZ NOT NULL,
	contract_aid BIGINT NOT NULL, new_value DECIMAL NOT NULL, UNIQUE(evtlog_id));
CREATE TABLE IF NOT EXISTS cg_adm_late_bid_premium_exponent(
	id BIGSERIAL PRIMARY KEY, evtlog_id BIGINT REFERENCES evt_log(id) ON DELETE CASCADE,
	block_num BIGINT NOT NULL, tx_id BIGINT NOT NULL, time_stamp TIMESTAMPTZ NOT NULL,
	contract_aid BIGINT NOT NULL, new_value DECIMAL NOT NULL, UNIQUE(evtlog_id));
CREATE TABLE IF NOT EXISTS cg_adm_last_bidder_reward_pct( -- new_value is a percentage 0..100
	id BIGSERIAL PRIMARY KEY, evtlog_id BIGINT REFERENCES evt_log(id) ON DELETE CASCADE,
	block_num BIGINT NOT NULL, tx_id BIGINT NOT NULL, time_stamp TIMESTAMPTZ NOT NULL,
	contract_aid BIGINT NOT NULL, new_value DECIMAL NOT NULL, UNIQUE(evtlog_id));
CREATE TABLE IF NOT EXISTS cg_adm_main_prize_num_nfts(
	id BIGSERIAL PRIMARY KEY, evtlog_id BIGINT REFERENCES evt_log(id) ON DELETE CASCADE,
	block_num BIGINT NOT NULL, tx_id BIGINT NOT NULL, time_stamp TIMESTAMPTZ NOT NULL,
	contract_aid BIGINT NOT NULL, new_value DECIMAL NOT NULL, UNIQUE(evtlog_id));
