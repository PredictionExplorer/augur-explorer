CREATE OR REPLACE FUNCTION on_fpmm_fund_add_insert() RETURNS trigger AS  $$
DECLARE
	v_cnt			NUMERIC;
BEGIN

	UPDATE pol_mkt_stats
		SET
			open_interest = (open_interest + NEW.shares_minted),
			num_liquidity_ops = (num_liquidity_ops + 1)
		WHERE contract_aid = NEW.contract_aid;
	GET DIAGNOSTICS v_cnt = ROW_COUNT;
	IF v_cnt = 0 THEN
		INSERT INTO pol_mkt_stats(contract_aid,open_interest,num_liquidity_ops)
			VALUES(NEW.contract_aid,NEW.shares_minted,1);
	END IF;
	RETURN NEW;
END;
$$ LANGUAGE plpgsql;
CREATE OR REPLACE FUNCTION on_fpmm_fund_add_delete() RETURNS trigger AS  $$
DECLARE
BEGIN

	UPDATE pol_mkt_stats
		SET
			open_interest = (open_interest - OLD.shares_minted),
			num_liquidity_ops = (num_liquidity_ops - 1)
		WHERE contract_aid = OLD.contract_aid;
	RETURN OLD;
END;
$$ LANGUAGE plpgsql;
CREATE OR REPLACE FUNCTION on_fpmm_fund_rem_insert() RETURNS trigger AS  $$
DECLARE
	v_cnt			NUMERIC;
BEGIN

	UPDATE pol_mkt_stats
		SET
			open_interest = (open_interest - NEW.shares_burnt),
			num_liquidity_ops = (num_liquidity_ops + 1)
		WHERE contract_aid = NEW.contract_aid;
	GET DIAGNOSTICS v_cnt = ROW_COUNT;
	IF v_cnt = 0 THEN
		INSERT INTO pol_mkt_stats(contract_aid,open_interest,num_liquidity_ops)
			VALUES(NEW.contract_aid,NEW.shares_burnt,1);
	END IF;
	RETURN NEW;
END;
$$ LANGUAGE plpgsql;
CREATE OR REPLACE FUNCTION on_fpmm_fund_rem_delete() RETURNS trigger AS  $$
DECLARE
BEGIN

	UPDATE pol_mkt_stats
		SET
			open_interest = (open_interest + OLD.shares_burnt),
			num_liquidity_ops = (num_liquidity_ops - 1)
		WHERE contract_aid = OLD.contract_aid;
	RETURN OLD;
END;
$$ LANGUAGE plpgsql;
CREATE OR REPLACE FUNCTION on_fpmm_buysell_insert() RETURNS trigger AS  $$
DECLARE
	v_cnt				NUMERIC;
BEGIN
	-- Noote: actually these are copies, but they meant to be for the future, if code diverges
	IF NEW.op_type = 0 THEN -- BUY
		UPDATE pol_mkt_stats
			SET
				total_volume = (total_volume + NEW.collateral_amount),
				total_fees = (total_fees + NEW.fee_amount),
				num_trades = (num_trades + 1)
			WHERE contract_aid = NEW.contract_aid;
		GET DIAGNOSTICS v_cnt = ROW_COUNT;
		IF v_cnt = 0 THEN
			INSERT INTO pol_mkt_stats (contract_aid,total_volume,num_trades,total_fees)
				VALUES(NEW.contract_aid,NEW.collateral_amount,1,NEW.fee_amount);
		END IF;
	END IF;
	IF NEW.op_type = 1 THEN -- SELL
		UPDATE pol_mkt_stats
			SET
				total_volume = (total_volume + NEW.collateral_amount),
				total_fees = (total_fees + NEW.fee_amount),
				num_trades = (num_trades + 1)
			WHERE contract_aid = NEW.contract_aid;
		GET DIAGNOSTICS v_cnt = ROW_COUNT;
		IF v_cnt = 0 THEN
			INSERT INTO pol_mkt_stats (contract_aid,total_volume,num_trades,total_fees)
				VALUES(NEW.contract_aid,NEW.collateral_amount,1,NEW.fee_amount);
		END IF;
	END IF;
	RETURN NEW;
END;
$$ LANGUAGE plpgsql;
CREATE OR REPLACE FUNCTION on_fpmm_buysell_delete() RETURNS trigger AS  $$
DECLARE
BEGIN

	IF OLD.op_type = 0 THEN -- BUY
		UPDATE pol_mkt_stats
			SET
				total_volume = (total_volume - OLD.investment_amount),
				total_fees = (total_fees - OLD.fee_amount),
				num_trades = (num_trades - 1)
			WHERE contract_aid = OLD.contract_aid;
	END IF;
	IF OLD.op_type = 1 THEN -- SELL
		UPDATE pol_mkt_stats
			SET
				total_volume = (total_volume - OLD.investment_amount),
				total_fees = (total_fees - OLD.fee_amount),
				num_trades = (num_trades - 1)
			WHERE contract_aid = OLD.contract_aid;
	END IF;
	RETURN OLD;
END;
$$ LANGUAGE plpgsql;
