CREATE OR REPLACE FUNCTION update_oi_on_insert() RETURNS trigger AS  $$ --updates open interest of the market
DECLARE
BEGIN

	UPDATE market SET open_interest = NEW.oi WHERE market.market_aid=NEW.market_aid;

	RETURN NEW;
END;
$$ LANGUAGE plpgsql;
CREATE OR REPLACE FUNCTION update_vol_on_insert() RETURNS trigger AS  $$ --updates volume of the market
DECLARE
BEGIN

	UPDATE market SET cur_volume = NEW.volume WHERE market.market_aid=NEW.market_aid;

	RETURN NEW;
END;
$$ LANGUAGE plpgsql;
CREATE OR REPLACE FUNCTION update_mkt_fin_on_insert() RETURNS trigger AS  $$ --updates market finalization event
DECLARE
BEGIN

	UPDATE market SET fin_timestamp = NEW.fin_timestamp,winning_payouts=NEW.winning_payouts WHERE market.market_aid=NEW.market_aid;

	RETURN NEW;
END;
$$ LANGUAGE plpgsql;
CREATE OR REPLACE FUNCTION update_oostats_on_insert() RETURNS trigger AS  $$ --updates open order statistics
DECLARE
BEGIN

	IF NEW.otype = 0 THEN
		UPDATE oostats AS s
			SET num_bids = num_bids + 1
			WHERE	(s.market_aid = NEW.market_aid) AND
					(s.eoa_aid = NEW.eoa_aid) AND
					(s.outcome_idx = NEW.outcome_idx);
	END IF;
	IF NEW.otype = 1 THEN
		UPDATE oostats AS s
			SET num_asks = num_asks + 1
			WHERE	(s.market_aid = NEW.market_aid) AND
					(s.eoa_aid = NEW.eoa_aid) AND
					(s.outcome_idx = NEW.outcome_idx);
	END IF;

	RETURN NEW;
END;
$$ LANGUAGE plpgsql;
CREATE OR REPLACE FUNCTION update_oostats_on_delete() RETURNS trigger AS  $$ -- reverts order statistics on delete
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

