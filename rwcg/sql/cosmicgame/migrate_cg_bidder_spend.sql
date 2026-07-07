-- Migration: per-bidder spend tracking for the ROI leaderboard.
--
-- Adds total_eth_spent / total_cst_spent to cg_bidder, updates the bid
-- insert/delete triggers to maintain them, and backfills existing rows.
--
-- Idempotent and safe to re-run. Run inside the cosmicgame schema
-- (e.g. `SET search_path TO <schema>;` first, or qualify the table names).
--
-- Apply with:  psql "<conn>" -v ON_ERROR_STOP=1 -f migrate_cg_bidder_spend.sql

BEGIN;

-- 1) New columns (no-op if already present)
ALTER TABLE cg_bidder ADD COLUMN IF NOT EXISTS total_eth_spent DECIMAL DEFAULT 0;
ALTER TABLE cg_bidder ADD COLUMN IF NOT EXISTS total_cst_spent DECIMAL DEFAULT 0;

-- 2) Updated trigger functions (CREATE OR REPLACE = idempotent)
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

-- 3) Backfill from existing bids. eth_price = -1 on CST bids and cst_price = -1
--    on ETH/RW bids, so the > 0 guard counts only real payments.
UPDATE cg_bidder b
SET total_eth_spent = COALESCE(s.eth_spent, 0),
    total_cst_spent = COALESCE(s.cst_spent, 0)
FROM (
	SELECT bidder_aid,
	       SUM(CASE WHEN eth_price > 0 THEN eth_price ELSE 0 END) AS eth_spent,
	       SUM(CASE WHEN cst_price > 0 THEN cst_price ELSE 0 END) AS cst_spent
	FROM cg_bid
	GROUP BY bidder_aid
) s
WHERE b.bidder_aid = s.bidder_aid;

COMMIT;
