CREATE OR REPLACE FUNCTION on_fpmm_fund_add_insert() RETURNS trigger AS  $$
DECLARE
	v_market_id		BIGINT;
	v_cnt			NUMERIC;
BEGIN

	SELECT mkt_mkr_aid FROM pol_market WHERE NEW.contract_aid = mkt_mkr_aid INTO v_market_id;
	IF v_market_id IS NULL THEN
		INSERT INTO update_needed(market_update) VALUES(TRUE);
		RAISE EXCEPTION 
			'Market id not determined using market maker aid = %. Market update requested',
			NEW.contract_aid;
	END IF;
	UPDATE pol_mkt_stats
		SET
			open_interest = (open_interest + NEW.shares_minted),
			num_liquidity_ops = (num_liquidity_ops + 1)
		WHERE market_id = v_market_id;
	GET DIAGNOSTICS v_cnt = ROW_COUNT;
	IF v_cnt = 0 THEN
		INSERT INTO pol_mkt_stats (market_id,open_interest,num_liquidity_ops)
			VALUES(v_market_id,NEW.shares_minted,1);
	END IF;
	RETURN NEW;
END;
$$ LANGUAGE plpgsql;
CREATE OR REPLACE FUNCTION on_fpmm_fund_add_delete() RETURNS trigger AS  $$
DECLARE
	v_market_id		BIGINT;
BEGIN

	SELECT mkt_mkr_aid FROM pol_market WHERE NEW.contract_aid = mkt_mkr_aid INTO v_market_id;
	IF v_market_id IS NULL THEN
		RETURN OLD;
	END IF;
	UPDATE pol_mkt_stats
		SET
			open_interest = (open_interest - OLD.shares_minted),
			num_liquidity_ops = (num_liquidity_ops - 1)
		WHERE market_id = v_market_id;
	RETURN OLD;
END;
$$ LANGUAGE plpgsql;
CREATE OR REPLACE FUNCTION on_fpmm_fund_rem_insert() RETURNS trigger AS  $$
DECLARE
	v_market_id		BIGINT;
	v_cnt			NUMERIC;
BEGIN

	SELECT mkt_mkr_aid FROM pol_market WHERE NEW.contract_aid = mkt_mkr_aid INTO v_market_id;
	IF v_market_id IS NULL THEN
		INSERT INTO update_needed(market_update) VALUES(TRUE);
		RAISE EXCEPTION 
			'Market id not determined using market maker aid = %. Market update requested',
			NEW.contract_aid;
	END IF;
	UPDATE pol_mkt_stats
		SET
			open_interest = (open_interest - NEW.shares_burnt),
			num_liquidity_ops = (num_liquidity_ops + 1)
		WHERE market_id = v_market_id;
	GET DIAGNOSTICS v_cnt = ROW_COUNT;
	IF v_cnt = 0 THEN
		INSERT INTO pol_mkt_stats (market_id,open_interest,num_liquidity_ops)
			VALUES(v_market_id,NEW.shares_burnt,1);
	END IF;
	RETURN NEW;
END;
$$ LANGUAGE plpgsql;
CREATE OR REPLACE FUNCTION on_fpmm_fund_rem_delete() RETURNS trigger AS  $$
DECLARE
	v_market_id		BIGINT;
BEGIN

	SELECT mkt_mkr_aid FROM pol_market WHERE NEW.contract_aid = mkt_mkr_aid INTO v_market_id;
	IF v_market_id IS NULL THEN
		RETURN OLD;
	END IF;
	UPDATE pol_mkt_stats
		SET
			open_interest = (open_interest + OLD.shares_burnt),
			num_liquidity_ops = (num_liquidity_ops - 1)
		WHERE market_id = v_market_id;
	RETURN OLD;
END;
$$ LANGUAGE plpgsql;
CREATE OR REPLACE FUNCTION on_fpmm_buy_insert() RETURNS trigger AS  $$
DECLARE
	v_market_id		BIGINT;
	v_cnt			NUMERIC;
BEGIN
	SELECT mkt_mkr_aid FROM pol_market WHERE NEW.contract_aid = mkt_mkr_aid INTO v_market_id;
	IF v_market_id IS NULL THEN
		INSERT INTO update_needed(market_update) VALUES(TRUE);
		RAISE EXCEPTION 
			'Market id not determined using market maker aid = %. Market update requested',
			NEW.contract_aid;
	END IF;
	UPDATE pol_mkt_stats
		SET
			total_volume = (total_volume + NEW.investment_amount),
			total_fees = (total_fees + NEW.fee_amount),
			num_trades = (num_trades + 1)
		WHERE market_id = v_market_id;
	GET DIAGNOSTICS v_cnt = ROW_COUNT;
	IF v_cnt = 0 THEN
		INSERT INTO pol_mkt_stats (market_id,total_volume,num_trades,total_fees)
			VALUES(v_market_id,NEW.total_volume,1,NEW.fee_amount);
	END IF;
	RETURN NEW;
END;
$$ LANGUAGE plpgsql;
CREATE OR REPLACE FUNCTION on_fpmm_buy_delete() RETURNS trigger AS  $$
DECLARE
	v_market_id		BIGINT;
BEGIN

	SELECT mkt_mkr_aid FROM pol_market WHERE NEW.contract_aid = mkt_mkr_aid INTO v_market_id;
	IF v_market_id IS NULL THEN
		INSERT INTO update_needed(market_update) VALUES(TRUE);
		RAISE EXCEPTION 
			'Market id not determined using market maker aid = %. Market update requested',
			NEW.contract_aid;
	END IF;
	UPDATE pol_mkt_stats
		SET
			total_volume = (total_volume - NEW.investment_amount),
			total_fees = (total_fees - NEW.fee_amount),
			num_trades = (num_trades - 1)
		WHERE market_id = v_market_id;
	RETURN OLD;
END;
$$ LANGUAGE plpgsql;
CREATE OR REPLACE FUNCTION on_fpmm_sell_insert() RETURNS trigger AS  $$
DECLARE
	v_market_id		BIGINT;
	v_cnt			NUMERIC;
BEGIN
	SELECT mkt_mkr_aid FROM pol_market WHERE NEW.contract_aid = mkt_mkr_aid INTO v_market_id;
	IF v_market_id IS NULL THEN
		INSERT INTO update_needed(market_update) VALUES(TRUE);
		RAISE EXCEPTION 
			'Market id not determined using market maker aid = %. Market update requested',
			NEW.contract_aid;
	END IF;
	UPDATE pol_mkt_stats
		SET
			total_volume = (total_volume - NEW.return_amount),
			total_fees = (total_fees + NEW.fee_amount),
			num_trades = (num_trades + 1)
		WHERE market_id = v_market_id;
	GET DIAGNOSTICS v_cnt = ROW_COUNT;
	IF v_cnt = 0 THEN
		INSERT INTO pol_mkt_stats (market_id,total_volume,num_trades,total_fees)
			VALUES(v_market_id,NEW.total_volume,1,NEW.fee_amount);
	END IF;
	RETURN NEW;
END;
$$ LANGUAGE plpgsql;
CREATE OR REPLACE FUNCTION on_fpmm_sell_delete() RETURNS trigger AS  $$
DECLARE
	v_market_id		BIGINT;
BEGIN

	SELECT mkt_mkr_aid FROM pol_market WHERE NEW.contract_aid = mkt_mkr_aid INTO v_market_id;
	IF v_market_id IS NULL THEN
		INSERT INTO update_needed(market_update) VALUES(TRUE);
		RAISE EXCEPTION 
			'Market id not determined using market maker aid = %. Market update requested',
			NEW.contract_aid;
	END IF;
	UPDATE pol_mkt_stats
		SET
			total_volume = (total_volume + NEW.return_amount),
			total_fees = (total_fees - NEW.fee_amount),
			num_trades = (num_trades - 1)
		WHERE market_id = v_market_id;
	RETURN OLD;
END;
$$ LANGUAGE plpgsql;
