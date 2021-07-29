CREATE OR REPLACE FUNCTION on_fpmm_fund_addrem_insert() RETURNS trigger AS  $$
DECLARE
	v_cnt			NUMERIC;
BEGIN

	-- Noote: actually these are copies, but they meant to be for the future, if code diverges
	IF NEW.op_type = 0 THEN -- ADD LIQUIDITY
		UPDATE pol_mkt_stats
			SET
				open_interest = (open_interest + NEW.shares_minted),
				total_collateral = (total_collateral + NEW.sum_amounts),
				num_liquidity_ops = (num_liquidity_ops + 1)
			WHERE contract_aid = NEW.contract_aid;
		GET DIAGNOSTICS v_cnt = ROW_COUNT;
		IF v_cnt = 0 THEN
			INSERT INTO pol_mkt_stats(contract_aid,open_interest,num_liquidity_ops)
				VALUES(NEW.contract_aid,NEW.shares_minted,1);
		END IF;
	END IF;
	IF NEW.op_type = 1 THEN -- REMOVE LIQUIDITY
		UPDATE pol_mkt_stats
			SET
				open_interest = (open_interest + NEW.shares_minted),
				total_collateral = (total_collateral + NEW.sum_amounts),
				num_liquidity_ops = (num_liquidity_ops + 1)
			WHERE contract_aid = NEW.contract_aid;
		GET DIAGNOSTICS v_cnt = ROW_COUNT;
		IF v_cnt = 0 THEN
			INSERT INTO pol_mkt_stats(contract_aid,open_interest,num_liquidity_ops)
				VALUES(NEW.contract_aid,NEW.shares_minted,1);
		END IF;
	END IF;
	RETURN NEW;
END;
$$ LANGUAGE plpgsql;
CREATE OR REPLACE FUNCTION on_fpmm_fund_addrem_delete() RETURNS trigger AS  $$
DECLARE
BEGIN

	IF OLD.op_type = 0 THEN -- ADD LIQUIDITY
		UPDATE pol_mkt_stats
			SET
				open_interest = (open_interest - OLD.shares_minted),
				total_collateral = (total_collateral - OLD.sum_amounts),
				num_liquidity_ops = (num_liquidity_ops - 1)
			WHERE contract_aid = OLD.contract_aid;
	END IF;
	IF OLD.op_type = 0 THEN -- REMOVE LOQUIDITY
		UPDATE pol_mkt_stats
			SET
				open_interest = (open_interest - OLD.shares_minted),
				total_collateral = (total_collateral - OLD.sum_amounts),
				num_liquidity_ops = (num_liquidity_ops - 1)
			WHERE contract_aid = OLD.contract_aid;
	RETURN OLD;
END;
$$ LANGUAGE plpgsql;
CREATE OR REPLACE FUNCTION on_fpmm_buysell_insert() RETURNS trigger AS  $$
DECLARE
	v_cnt					NUMERIC;
	v_prev_profit_loss		DECIMAL;
	v_profit_loss			DECIMAL;
	v_accum_profit_loss		DECIMAL;
	v_accum_collateral		DECIMAL;
BEGIN
	SELECT accum_profit_loss,accum_collateral FROM pol_buysell WHERE user_aid = NEW.user_aid 
		ORDER BY evtlog_id DESC LIMIT 1 INTO v_prev_profit_loss,v_accum_collateral;
	IF v_prev_profit_loss IS NULL THEN
		v_prev_profit_loss := 0;
	END IF;
	IF v_accum_collateral IS NULL THEN
		v_accum_collateral := 0;
	END IF;
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

		v_profit_loss := v_prev_profit_loss - NEW.collateral_amount;
		v_accum_profit_loss := v_accum_profit_loss + v_profit_loss;

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

		v_profit_loss := v_prev_profit_loss + NEW.collateral_amount;

	END IF;
	UPDATE pol_buysell SET
			profit_loss = v_profit_loss,
			accum_profit_loss = v_provit_loss
		WHERE id=NEW.id;
	RETURN NEW;
END;
$$ LANGUAGE plpgsql;
CREATE OR REPLACE FUNCTION on_fpmm_buysell_delete() RETURNS trigger AS  $$
DECLARE
	v_profit_loss			DECIMAL;
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

	UPDATE pol_ustats SET
		total_buysell_ops = (total_buysell_ops - 1),
		WHERE user_id = OLD.user_id
	RETURN OLD;
END;
$$ LANGUAGE plpgsql;
