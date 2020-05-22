CREATE OR REPLACE FUNCTION on_oi_chg_insert() RETURNS trigger AS  $$ --updates open interest of the market
DECLARE
BEGIN

	UPDATE market SET open_interest = NEW.oi WHERE market.market_aid=NEW.market_aid;

	RETURN NEW;
END;
$$ LANGUAGE plpgsql;
CREATE OR REPLACE FUNCTION on_volume_insert() RETURNS trigger AS  $$ --updates volume of the market
DECLARE
BEGIN

	UPDATE market SET cur_volume = NEW.volume WHERE market.market_aid=NEW.market_aid;

	RETURN NEW;
END;
$$ LANGUAGE plpgsql;
CREATE OR REPLACE FUNCTION on_oorders_insert() RETURNS trigger AS  $$ --updates open order statistics
DECLARE
	v_cnt numeric;
BEGIN

	IF NEW.otype = 0 THEN
		UPDATE oostats AS s
			SET num_bids = num_bids + 1
			WHERE	(s.market_aid = NEW.market_aid) AND
					(s.eoa_aid = NEW.eoa_aid) AND
					(s.outcome_idx = NEW.outcome_idx);
		GET DIAGNOSTICS v_cnt = ROW_COUNT;
		IF v_cnt = 0 THEN
			INSERT	INTO oostats(market_aid,eoa_aid,outcome_idx,num_bids)
					VALUES(NEW.market_aid,NEW.eoa_aid,NEW.outcome_idx,1);
		END IF;
	END IF;
	IF NEW.otype = 1 THEN
		UPDATE oostats AS s
			SET num_asks = num_asks + 1
			WHERE	(s.market_aid = NEW.market_aid) AND
					(s.eoa_aid = NEW.eoa_aid) AND
					(s.outcome_idx = NEW.outcome_idx);
		GET DIAGNOSTICS v_cnt = ROW_COUNT;
		IF v_cnt = 0 THEN
			INSERT	INTO oostats(market_aid,eoa_aid,outcome_idx,num_asks)
					VALUES(NEW.market_aid,NEW.eoa_aid,NEW.outcome_idx,1);
		END IF;
	END IF;

	RETURN NEW;
END;
$$ LANGUAGE plpgsql;
CREATE OR REPLACE FUNCTION on_oorders_delete() RETURNS trigger AS  $$ -- reverts order statistics on delete
DECLARE
BEGIN

	IF OLD.otype = 0 THEN
		UPDATE oostats AS s
			SET num_bids = num_bids - 1
			WHERE	(s.market_aid = OLD.market_aid) AND
					(s.eoa_aid = OLD.eoa_aid) AND
					(s.outcome_idx = OLD.outcome_idx);
	END IF;
	IF OLD.otype = 1 THEN
		UPDATE oostats AS s
			SET num_asks = num_asks - 1
			WHERE	(s.market_aid = OLD.market_aid) AND
					(s.eoa_aid = OLD.eoa_aid) AND
					(s.outcome_idx = OLD.outcome_idx);
	END IF;

	RETURN OLD;
END;
$$ LANGUAGE plpgsql;
-- update all the statistics on insertion into 'market' table
CREATE OR REPLACE FUNCTION on_market_insert() RETURNS trigger AS  $$
DECLARE
	v_cnt numeric;
BEGIN

	UPDATE ustats AS s
			SET markets_created = markets_created + 1
			WHERE s.eoa_aid = NEW.eoa_aid;

	GET DIAGNOSTICS v_cnt = ROW_COUNT;
	IF v_cnt = 0 THEN
		INSERT	INTO ustats(eoa_aid,wallet_aid,markets_created)
				VALUES(NEW.eoa_aid,NEW.wallet_aid,1);
	END IF;

	UPDATE main_stats
		SET markets_count = markets_count + 1, active_count = active_count +1;
	IF NEW.market_type = 0 THEN
		UPDATE main_stats SET yesno_count = yesno_count + 1;
	END IF;
	IF NEW.market_type = 1 THEN
		UPDATE main_stats SET categ_count = categ_count + 1;
	END IF;
	IF NEW.market_type = 2 THEN
		UPDATE main_stats SET scalar_count = scalar_count + 1;
	END IF;

	RETURN NEW;
END;
$$ LANGUAGE plpgsql;
CREATE OR REPLACE FUNCTION on_market_delete() RETURNS trigger AS  $$
DECLARE
	v_cnt numeric;
BEGIN

	UPDATE ustats AS s
			SET markets_created = markets_created - 1
			WHERE s.eoa_aid = OLD.eoa_aid;

	GET DIAGNOSTICS v_cnt = ROW_COUNT;
	IF v_cnt = 0 THEN	-- this condition won't be true during normal operation
		INSERT	INTO ustats(eoa_aid,wallet_aid,markets_created)
				VALUES(OLD.eoa_aid,OLD.wallet_aid,0);
	END IF;

	UPDATE main_stats
		SET markets_count = markets_count - 1, active_count = active_count - 1;
	IF NEW.market_type = 0 THEN
		UPDATE main_stats SET yesno_count = yesno_count - 1;
	END IF;
	IF NEW.market_type = 1 THEN
		UPDATE main_stats SET categ_count = categ_count - 1;
	END IF;
	IF NEW.market_type = 2 THEN
		UPDATE main_stats SET scalar_count = scalar_count - 1;
	END IF;

	RETURN OLD;
END;
$$ LANGUAGE plpgsql;
-- update all the statistics on insertion into 'mktord' table
CREATE OR REPLACE FUNCTION on_mktord_insert() RETURNS trigger AS  $$
DECLARE
	v_cnt numeric;
BEGIN

	-- Make sure user stats record exists
	INSERT INTO ustats(eoa_aid,wallet_aid) VALUES(NEW.eoa_aid,NEW.wallet_aid)
		ON CONFLICT(eoa_aid) DO NOTHING;
	INSERT INTO ustats(eoa_aid,wallet_aid) VALUES(NEW.eoa_fill_aid,NEW.wallet_fill_aid)
		ON CONFLICT(eoa_aid) DO NOTHING;

	-- Update statistics for the Creator of the Order (Seller)
	UPDATE trd_mkt_stats AS s
			SET total_trades = total_trades + 1
			WHERE	s.eoa_aid = NEW.eoa_aid AND
					s.market_aid = NEW.market_aid;

	GET DIAGNOSTICS v_cnt = ROW_COUNT;
	IF v_cnt = 0 THEN
		INSERT	INTO trd_mkt_stats(eoa_aid,wallet_aid,market_aid)
				VALUES(NEW.eoa_aid,NEW.wallet_aid,NEW.market_aid);
	END IF;

	-- Update statistics for the Filler of the Order (Buyer)
	UPDATE trd_mkt_stats AS s
			SET total_trades = total_trades + 1
			WHERE	s.eoa_aid = NEW.eoa_fill_aid AND
					s.market_aid = NEW.market_aid;

	GET DIAGNOSTICS v_cnt = ROW_COUNT;
	IF v_cnt = 0 THEN
		INSERT	INTO trd_mkt_stats(eoa_aid,wallet_aid,market_aid)
				VALUES(NEW.eoa_fill_aid,NEW.wallet_fill_aid,NEW.market_aid);
	END IF;

	UPDATE main_stats SET trades_count = trades_count + 1;

	RETURN NEW;
END;
$$ LANGUAGE plpgsql;
CREATE OR REPLACE FUNCTION on_mktord_delete() RETURNS trigger AS  $$
DECLARE
	v_cnt numeric;
BEGIN

	-- Make sure user stats record exists
	INSERT INTO ustats(eoa_aid,wallet_aid) VALUES(NEW.eoa_aid,NEW.wallet_aid)
		ON CONFLICT(eoa_aid) DO NOTHING;
	INSERT INTO ustats(eoa_aid,wallet_aid) VALUES(NEW.eoa_fill_aid,NEW.wallet_fill_aid)
		ON CONFLICT(eoa_aid) DO NOTHING;

	-- Update statistics for the Creator of the Order (Seller)
	UPDATE trd_mkt_stats AS s
			SET total_trades = total_trades - 1
			WHERE	s.eoa_aid = OLD.eoa_aid AND
					s.market_aid = OLD.market_aid;

	GET DIAGNOSTICS v_cnt = ROW_COUNT;
	IF v_cnt = 0 THEN	-- this condition won't be true during normal operation
		INSERT	INTO trd_mkt_stats(eoa_aid,wallet_aid,market_aid)
				VALUES(OLD.eoa_aid,OLD.wallet_aid,OLD.market_aid);
	END IF;

	-- Update statistics for the Filler of the Order (Buyer)
	UPDATE trd_mkt_stats AS s
			SET total_trades = total_trades - 1
			WHERE	s.eoa_aid = OLD.eoa_fill_aid AND
					s.market_aid = OLD.market_aid;

	GET DIAGNOSTICS v_cnt = ROW_COUNT;
	IF v_cnt = 0 THEN
		INSERT	INTO trd_mkt_stats(eoa_aid,wallet_aid,market_aid)
				VALUES(OLD.eoa_fill_aid,OLD.wallet_fill_aid,OLD.market_aid);
	END IF;

	UPDATE main_stats SET trades_count = trades_count - 1;

	RETURN OLD;
END;
$$ LANGUAGE plpgsql;
-- update all the statistics on insertion into 'trd_mkt_stats' table
CREATE OR REPLACE FUNCTION on_trd_mkt_stats_insert() RETURNS trigger AS  $$
DECLARE
	v_cnt numeric;
BEGIN

	-- Update statistics for the Creator of the Order (Seller)
	UPDATE ustats AS s
			SET markets_traded = markets_traded + 1
			WHERE	s.eoa_aid = NEW.eoa_aid;

	GET DIAGNOSTICS v_cnt = ROW_COUNT;
	IF v_cnt = 0 THEN
		RAISE EXCEPTION 'Corresponding row in ustats table doesnt exist';
	END IF;

	RETURN NEW;
END;
$$ LANGUAGE plpgsql;
CREATE OR REPLACE FUNCTION on_trd_mkt_stats_delete() RETURNS trigger AS  $$
DECLARE
	v_cnt numeric;
BEGIN

	-- Update statistics for the Creator of the Order (Seller)
	SELECT COUNT(*) AS num_rows
		FROM trd_mkt_stats AS s
		WHERE	s.eoa_aid = OLD.eoa_aid AND
				s.market_aid = OLD.market_aid;

	GET DIAGNOSTICS v_cnt = ROW_COUNT;
	IF v_cnt = 0 THEN
		UPDATE ustats AS s
			SET markets_traded = markets_traded - 1
			WHERE	s.eoa_aid = OLD.eoa_aid;
	END IF;

	RETURN OLD;
END;
$$ LANGUAGE plpgsql;
-- reporting triggers
CREATE OR REPLACE FUNCTION on_report_insert() RETURNS trigger AS  $$ --updates volume of the market
DECLARE
	v_cnt numeric;
BEGIN

	-- Make sure user stats record exists
	INSERT INTO ustats(eoa_aid,wallet_aid) VALUES(NEW.eoa_aid,NEW.wallet_aid)
		ON CONFLICT(eoa_aid) DO NOTHING;

	-- Update statistics for the Reporter
	UPDATE trd_mkt_stats AS s
			SET total_reports = total_reports + 1
			WHERE	s.eoa_aid = NEW.eoa_aid AND
					s.market_aid = NEW.market_aid;
	GET DIAGNOSTICS v_cnt = ROW_COUNT;
	IF v_cnt = 0 THEN
		INSERT	INTO trd_mkt_stats(eoa_aid,wallet_aid,market_aid)
				VALUES(NEW.eoa_aid,NEW.wallet_aid,NEW.market_aid);
	END IF;

	IF NEW.is_designated IS TRUE THEN
		UPDATE trd_mkt_stats AS s
			SET total_designated = total_designated + 1
			WHERE	s.eoa_aid = NEW.eoa_aid AND
					s.market_aid = NEW.market_aid;
	END IF;

	RETURN NEW;
END;
$$ LANGUAGE plpgsql;
CREATE OR REPLACE FUNCTION on_report_delete() RETURNS trigger AS  $$ -- reverts order statistics on delete
DECLARE
	v_cnt numeric;
BEGIN

	-- Make sure user stats record exists
	INSERT INTO ustats(eoa_aid,wallet_aid) VALUES(OLD.eoa_aid,OLD.wallet_aid)
		ON CONFLICT(eoa_aid) DO NOTHING;

	-- Update statistics for the Reporter
	UPDATE trd_mkt_stats AS s
			SET total_reports = total_reports - 1
			WHERE	s.eoa_aid = OLD.eoa_aid AND
					s.market_aid = OLD.market_aid;
	GET DIAGNOSTICS v_cnt = ROW_COUNT;
	IF v_cnt = 0 THEN
		INSERT	INTO trd_mkt_stats(eoa_aid,wallet_aid,market_aid)
				VALUES(OLD.eoa_aid,OLD.wallet_aid,OLD.market_aid);
	END IF;

	IF NEW.is_designated IS TRUE THEN
		UPDATE trd_mkt_stats AS s
			SET total_designated = total_designated - 1
			WHERE	s.eoa_aid = OLD.eoa_aid AND
					s.market_aid = OLD.market_aid;
	END IF;
	RETURN OLD;
END;
$$ LANGUAGE plpgsql;
CREATE OR REPLACE FUNCTION on_mktfin_insert() RETURNS trigger AS  $$
DECLARE
BEGIN

	UPDATE market
		SET fin_timestamp = NEW.fin_timestamp,
			winning_payouts=NEW.winning_payouts
		WHERE market.market_aid=NEW.market_aid;
	UPDATE main_stats SET active_count = active_count - 1;

	RETURN NEW;
END;
$$ LANGUAGE plpgsql;
CREATE OR REPLACE FUNCTION on_mktfin_delete() RETURNS trigger AS  $$
DECLARE
BEGIN

	UPDATE main_stats SET active_count = active_count + 1;
	RETURN OLD;
END;
$$ LANGUAGE plpgsql;
