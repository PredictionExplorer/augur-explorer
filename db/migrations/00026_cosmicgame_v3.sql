-- +goose Up
-- CosmicSignatureGame V3 adds a multi-NFT main prize, split bid rewards,
-- five evented configuration values, and two event-less per-round champion
-- duration values.

ALTER TABLE cg_prize_claim
	ADD COLUMN num_cs_nfts BIGINT NOT NULL DEFAULT 1 CHECK (num_cs_nfts > 0);

ALTER TABLE cg_round_stats
	ADD COLUMN endurance_champion_duration BIGINT NOT NULL DEFAULT 0 CHECK (endurance_champion_duration >= 0),
	ADD COLUMN chrono_warrior_duration BIGINT NOT NULL DEFAULT 0 CHECK (chrono_warrior_duration >= 0);

CREATE TABLE cg_bid_reward (
	id				BIGSERIAL PRIMARY KEY,
	evtlog_id		BIGINT REFERENCES evt_log(id) ON DELETE CASCADE,
	bid_id			BIGINT NOT NULL,
	round_num		BIGINT NOT NULL,
	recipient_aid	BIGINT NOT NULL,
	reward_type		SMALLINT NOT NULL CHECK (reward_type IN (0, 1)),
	amount			DECIMAL NOT NULL DEFAULT 0 CHECK (amount >= 0),
	UNIQUE (bid_id, reward_type)
);
COMMENT ON COLUMN cg_bid_reward.reward_type IS
	'0 = bidder placing this bid; 1 = outbid previous-last bidder';
CREATE INDEX idx_cg_bid_reward_recipient
	ON cg_bid_reward (recipient_aid, round_num, bid_id);

-- Pre-V3 bids credited the complete mint to the bidder placing the bid.
-- GREATEST protects historical V1 sentinel values while NOT EXISTS keeps
-- the migration safe if an operator pre-populated part of the table.
INSERT INTO cg_bid_reward(evtlog_id,bid_id,round_num,recipient_aid,reward_type,amount)
	SELECT b.evtlog_id, b.id, b.round_num, b.bidder_aid, 0,
		GREATEST(COALESCE(b.cst_reward, 0), 0)
	FROM cg_bid b
	WHERE NOT EXISTS (
		SELECT 1 FROM cg_bid_reward r
		WHERE r.bid_id = b.id AND r.reward_type = 0
	);

CREATE TABLE cg_adm_late_bid_dur_divisor (
	id				BIGSERIAL PRIMARY KEY,
	evtlog_id		BIGINT REFERENCES evt_log(id) ON DELETE CASCADE,
	block_num		BIGINT NOT NULL,
	tx_id			BIGINT NOT NULL,
	time_stamp		TIMESTAMPTZ NOT NULL,
	contract_aid	BIGINT NOT NULL,
	new_value		DECIMAL NOT NULL,
	UNIQUE (evtlog_id)
);

CREATE TABLE cg_adm_late_bid_premium_base_mul (
	id				BIGSERIAL PRIMARY KEY,
	evtlog_id		BIGINT REFERENCES evt_log(id) ON DELETE CASCADE,
	block_num		BIGINT NOT NULL,
	tx_id			BIGINT NOT NULL,
	time_stamp		TIMESTAMPTZ NOT NULL,
	contract_aid	BIGINT NOT NULL,
	new_value		DECIMAL NOT NULL,
	UNIQUE (evtlog_id)
);

CREATE TABLE cg_adm_late_bid_premium_exponent (
	id				BIGSERIAL PRIMARY KEY,
	evtlog_id		BIGINT REFERENCES evt_log(id) ON DELETE CASCADE,
	block_num		BIGINT NOT NULL,
	tx_id			BIGINT NOT NULL,
	time_stamp		TIMESTAMPTZ NOT NULL,
	contract_aid	BIGINT NOT NULL,
	new_value		DECIMAL NOT NULL,
	UNIQUE (evtlog_id)
);

CREATE TABLE cg_adm_last_bidder_reward_pct (
	id				BIGSERIAL PRIMARY KEY,
	evtlog_id		BIGINT REFERENCES evt_log(id) ON DELETE CASCADE,
	block_num		BIGINT NOT NULL,
	tx_id			BIGINT NOT NULL,
	time_stamp		TIMESTAMPTZ NOT NULL,
	contract_aid	BIGINT NOT NULL,
	new_value		DECIMAL NOT NULL CHECK (new_value BETWEEN 0 AND 100),
	UNIQUE (evtlog_id)
);

CREATE TABLE cg_adm_main_prize_num_nfts (
	id				BIGSERIAL PRIMARY KEY,
	evtlog_id		BIGINT REFERENCES evt_log(id) ON DELETE CASCADE,
	block_num		BIGINT NOT NULL,
	tx_id			BIGINT NOT NULL,
	time_stamp		TIMESTAMPTZ NOT NULL,
	contract_aid	BIGINT NOT NULL,
	new_value		DECIMAL NOT NULL CHECK (new_value > 0),
	UNIQUE (evtlog_id)
);

-- Event-less contract state belongs in this audit table, never in evt_log.
-- In particular, the runtime records championDurations(round) observations
-- here while reporting evented cg_adm_* mismatches as read-only drift.
-- Existing 0xcccc... synthetic chain-sync rows are deliberately preserved:
-- deleting production history requires a separate, operator-reviewed audit.
CREATE TABLE cg_live_state_updates (
	id				BIGSERIAL PRIMARY KEY,
	variable_name	TEXT NOT NULL,
	contract_aid	BIGINT NOT NULL,
	round_num		BIGINT NOT NULL DEFAULT -1,
	block_num		BIGINT NOT NULL,
	time_stamp		TIMESTAMPTZ NOT NULL DEFAULT NOW(),
	new_value		DECIMAL NOT NULL
);
CREATE INDEX idx_cg_live_state_updates_var
	ON cg_live_state_updates (variable_name, round_num, id);

-- +goose StatementBegin
CREATE OR REPLACE FUNCTION on_block_delete_cg_live_state_updates() RETURNS trigger AS $$
BEGIN
	-- Startup recovery can observe a head newer than the indexer's stored
	-- tip. Removing every observation at or above the divergent block also
	-- clears those not-yet-materialized future-block observations.
	DELETE FROM cg_live_state_updates WHERE block_num>=OLD.block_num;
	RETURN OLD;
END;
$$ LANGUAGE plpgsql;
-- +goose StatementEnd
CREATE TRIGGER block_delete_cg_live_state_updates
	BEFORE DELETE ON block
	FOR EACH ROW EXECUTE PROCEDURE on_block_delete_cg_live_state_updates();

-- +goose StatementBegin
CREATE OR REPLACE FUNCTION on_prize_claim_insert() RETURNS trigger AS $$
DECLARE
	v_max_prize				DECIMAL;
	v_prizes_sum			DECIMAL;
	v_prizes_count			BIGINT;
	v_donated_nfts			BIGINT;
	v_cnt					NUMERIC;
BEGIN
	-- cg_prize stores one row per prize category; the multi-NFT cardinality
	-- is carried by cg_prize_claim.num_cs_nfts.
	v_prizes_count := 3;

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
			max_win_amount = v_max_prize,
			prizes_sum = v_prizes_sum,
			unclaimed_nfts = (unclaimed_nfts + v_donated_nfts),
			erc20_count = (erc20_count + 1),
			erc721_count = (erc721_count + 1)
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

	UPDATE cg_glob_stats SET total_cst_given_in_prizes = (total_cst_given_in_prizes + NEW.cst_amount);
	GET DIAGNOSTICS v_cnt = ROW_COUNT;
	IF v_cnt = 0 THEN
		RAISE EXCEPTION 'on_prize_claim_insert() failed to update total_cst_given_in_prizes in cg_glob_stats';
	END IF;

	UPDATE cg_round_stats SET
		total_cst_paid_in_prizes = (total_cst_paid_in_prizes + NEW.cst_amount),
		total_nfts_minted = (total_nfts_minted + NEW.num_cs_nfts)
		WHERE round_num = NEW.round_num;
	GET DIAGNOSTICS v_cnt = ROW_COUNT;
	IF v_cnt = 0 THEN
		INSERT INTO cg_round_stats(round_num, total_cst_paid_in_prizes, total_nfts_minted)
			VALUES (NEW.round_num, NEW.cst_amount, NEW.num_cs_nfts);
		RAISE NOTICE 'on_prize_claim_insert() created new round_stats record for round_num=%', NEW.round_num;
	END IF;

	INSERT INTO cg_prize(round_num,winner_index,ptype) VALUES(NEW.round_num,0,0);
	INSERT INTO cg_prize(round_num,winner_index,ptype) VALUES(NEW.round_num,0,1);
	INSERT INTO cg_prize(round_num,winner_index,ptype) VALUES(NEW.round_num,0,2);

	RETURN NEW;
END;
$$ LANGUAGE plpgsql;
-- +goose StatementEnd

-- +goose StatementBegin
CREATE OR REPLACE FUNCTION on_prize_claim_delete() RETURNS trigger AS $$
DECLARE
	v_max_prize				DECIMAL;
	v_prizes_sum			DECIMAL;
	v_prizes_count			BIGINT;
	v_cnt					NUMERIC;
	v_donated_nfts			BIGINT;
BEGIN
	SELECT
			MAX(amount),
 			COUNT(*) AS prizes_count,
			SUM(amount) AS prizes_sum
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
			max_win_amount = v_max_prize,
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

	UPDATE cg_round_stats SET
		total_cst_paid_in_prizes = (total_cst_paid_in_prizes - OLD.cst_amount),
		total_nfts_minted = (total_nfts_minted - OLD.num_cs_nfts)
		WHERE round_num = OLD.round_num;

	DELETE FROM cg_prize WHERE round_num=OLD.round_num AND winner_index=0 AND ptype=0;
	DELETE FROM cg_prize WHERE round_num=OLD.round_num AND winner_index=0 AND ptype=1;
	DELETE FROM cg_prize WHERE round_num=OLD.round_num AND winner_index=0 AND ptype=2;

	RETURN OLD;
END;
$$ LANGUAGE plpgsql;
-- +goose StatementEnd

-- +goose Down
DROP TRIGGER block_delete_cg_live_state_updates ON block;
DROP FUNCTION on_block_delete_cg_live_state_updates();

-- Restore the V1/V2 trigger bodies before removing num_cs_nfts.
-- +goose StatementBegin
CREATE OR REPLACE FUNCTION on_prize_claim_insert() RETURNS trigger AS $$
DECLARE
	v_max_prize				DECIMAL;
	v_prizes_sum			DECIMAL;
	v_prizes_count			BIGINT;
	v_donated_nfts			BIGINT;
	v_cnt					NUMERIC;
BEGIN
	v_prizes_count := 3;

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
			max_win_amount = v_max_prize,
			prizes_sum = v_prizes_sum,
			unclaimed_nfts = (unclaimed_nfts + v_donated_nfts),
			erc20_count = (erc20_count + 1),
			erc721_count = (erc721_count + 1)
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

	UPDATE cg_glob_stats SET total_cst_given_in_prizes = (total_cst_given_in_prizes + NEW.cst_amount);
	GET DIAGNOSTICS v_cnt = ROW_COUNT;
	IF v_cnt = 0 THEN
		RAISE EXCEPTION 'on_prize_claim_insert() failed to update total_cst_given_in_prizes in cg_glob_stats';
	END IF;

	UPDATE cg_round_stats SET
		total_cst_paid_in_prizes = (total_cst_paid_in_prizes + NEW.cst_amount),
		total_nfts_minted = (total_nfts_minted + 1)
		WHERE round_num = NEW.round_num;
	GET DIAGNOSTICS v_cnt = ROW_COUNT;
	IF v_cnt = 0 THEN
		INSERT INTO cg_round_stats(round_num, total_cst_paid_in_prizes, total_nfts_minted)
			VALUES (NEW.round_num, NEW.cst_amount, 1);
		RAISE NOTICE 'on_prize_claim_insert() created new round_stats record for round_num=%', NEW.round_num;
	END IF;

	INSERT INTO cg_prize(round_num,winner_index,ptype) VALUES(NEW.round_num,0,0);
	INSERT INTO cg_prize(round_num,winner_index,ptype) VALUES(NEW.round_num,0,1);
	INSERT INTO cg_prize(round_num,winner_index,ptype) VALUES(NEW.round_num,0,2);

	RETURN NEW;
END;
$$ LANGUAGE plpgsql;
-- +goose StatementEnd

-- +goose StatementBegin
CREATE OR REPLACE FUNCTION on_prize_claim_delete() RETURNS trigger AS $$
DECLARE
	v_max_prize				DECIMAL;
	v_prizes_sum			DECIMAL;
	v_prizes_count			BIGINT;
	v_cnt					NUMERIC;
	v_donated_nfts			BIGINT;
BEGIN
	SELECT
			MAX(amount),
 			COUNT(*) AS prizes_count,
			SUM(amount) AS prizes_sum
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
			max_win_amount = v_max_prize,
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

	UPDATE cg_round_stats SET
		total_cst_paid_in_prizes = (total_cst_paid_in_prizes - OLD.cst_amount),
		total_nfts_minted = (total_nfts_minted - 1)
		WHERE round_num = OLD.round_num;

	DELETE FROM cg_prize WHERE round_num=OLD.round_num AND winner_index=0 AND ptype=0;
	DELETE FROM cg_prize WHERE round_num=OLD.round_num AND winner_index=0 AND ptype=1;
	DELETE FROM cg_prize WHERE round_num=OLD.round_num AND winner_index=0 AND ptype=2;

	RETURN OLD;
END;
$$ LANGUAGE plpgsql;
-- +goose StatementEnd

DROP TABLE cg_live_state_updates;
DROP TABLE cg_adm_main_prize_num_nfts;
DROP TABLE cg_adm_last_bidder_reward_pct;
DROP TABLE cg_adm_late_bid_premium_exponent;
DROP TABLE cg_adm_late_bid_premium_base_mul;
DROP TABLE cg_adm_late_bid_dur_divisor;
DROP TABLE cg_bid_reward;

ALTER TABLE cg_round_stats
	DROP COLUMN chrono_warrior_duration,
	DROP COLUMN endurance_champion_duration;
ALTER TABLE cg_prize_claim DROP COLUMN num_cs_nfts;
