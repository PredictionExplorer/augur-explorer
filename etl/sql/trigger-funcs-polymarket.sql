CREATE OR REPLACE FUNCTION on_fpmm_fund_addrem_insert() RETURNS trigger AS  $$
DECLARE
	v_cnt					NUMERIC;
	v_normalized_collateral DECIMAL;
BEGIN

	IF NEW.op_type = 0 THEN
		v_normalized_collateral := -NEW.shares;--SharesAdded (this is the collateral put into pool)
	ELSE 
		v_normalized_collateral := NEW.shares;	-- collateral taken out of the pool
	END IF;
	UPDATE pol_fund_addrem SET norm_collateral = v_normalized_collateral WHERE id=NEW.id;
	-- market statistics
	UPDATE pol_mkt_stats
		SET
			num_liq_ops = (num_liq_ops + 1),
			open_interest = (open_interest + v_normalized_collateral),
			total_volume = (total_volume + NEW.shares)
		WHERE contract_aid = NEW.contract_aid;
	GET DIAGNOSTICS v_cnt = ROW_COUNT;
	IF v_cnt = 0 THEN
		INSERT INTO pol_mkt_stats(contract_aid,num_liq_ops,open_interest,total_volume)
			VALUES(NEW.contract_aid,1,v_normalized_collateral,NEW.shares);
	END IF;

	-- global user statistics
	UPDATE pol_ustats
		SET
			tot_liq_ops = (tot_liq_ops + 1),
			tot_liq_given = (tot_liq_given + v_normalized_collateral)
		WHERE user_aid = NEW.funder_aid;
	GET DIAGNOSTICS v_cnt = ROW_COUNT;
	IF v_cnt = 0 THEN
		INSERT INTO pol_ustats(user_aid,reg_time_stamp,tot_liq_ops,tot_liq_given)
			VALUES(NEW.funder_aid,NEW.time_stamp,1,v_normalized_collateral);
	END IF;

	-- user statistics for this market in particular
	UPDATE pol_ustats_mkt
		SET
			tot_liq_ops = (tot_liq_ops + 1),
			tot_liq_given = (tot_liq_given + v_normalized_collateral),
			tot_volume = (tot_volume + NEW.shares)
		WHERE (user_aid = NEW.funder_aid) AND (contract_aid=NEW.contract_aid);
	GET DIAGNOSTICS v_cnt = ROW_COUNT;
	IF v_cnt = 0 THEN
		INSERT INTO pol_ustats_mkt(user_aid,tot_liq_ops,tot_liq_given,tot_volume)
			VALUES(NEW.funder_aid,1,v_normalized_collateral,NEW.shares);
	END IF;

	RETURN NEW;
END;
$$ LANGUAGE plpgsql;
CREATE OR REPLACE FUNCTION on_fpmm_fund_addrem_delete() RETURNS trigger AS  $$
DECLARE
BEGIN

	UPDATE pol_mkt_stats
		SET
			num_liq_ops = (num_liq_ops - 1),
			open_interest = (open_interest - OLD.norm_collateral),
			total_volume = (total_volume - OLD.shares)
		WHERE contract_aid = OLD.contract_aid;
	UPDATE pol_ustats
		SET
			tot_liq_ops = (tot_liq_ops - 1),
			tot_liq_given = (tot_liq_given - OLD.norm_collateral)
		WHERE user_aid = OLD.funder_aid;
	UPDATE pol_ustats_mkt
		SET
			tot_liq_ops = (tot_liq_ops - 1),
			tot_liq_given = (tot_liq_given - OLD.norm_collateral),
			tot_volume = (tot_volume - OLD.shares)
		WHERE (user_aid = OLD.funder_aid) AND (contract_aid=OLD.contract_aid);
	RETURN OLD;
END;
$$ LANGUAGE plpgsql;
CREATE OR REPLACE FUNCTION on_fpmm_buysell_insert() RETURNS trigger AS  $$
DECLARE
	v_cnt					NUMERIC;
	v_accum_collateral		DECIMAL;
	v_normalized_amount		DECIMAL;
BEGIN

	-- first calculate accumulated collateral (which is the same as profit/loss on this market)
	IF NEW.op_type = 0 THEN
		v_normalized_amount := -NEW.collateral_amount;
	ELSE 
		v_normalized_amount := NEW.collateral_amount;
	END IF;

	SELECT accum_collateral
		FROM pol_buysell
		WHERE (contract_aid = NEW.contract_aid) AND (id<NEW.id)
		ORDER BY id DESC LIMIT 1 INTO v_accum_collateral;
	IF v_accum_collateral IS NULL THEN
		v_accum_collateral := 0;
	END IF;

	v_accum_collateral := v_accum_collateral + v_normalized_amount;
	UPDATE pol_buysell
		SET
			accum_collateral = v_accum_collateral,
			normalized_amount = v_normalized_amount
		WHERE id=NEW.id;

	-- now update statistics, market statistics first
	UPDATE pol_mkt_stats
		SET
			num_trades = (num_trades + 1),
			total_volume = (total_volume + NEW.collateral_amount),
			total_fees = (total_fees + NEW.fee_amount)
		WHERE contract_aid = NEW.contract_aid;
	GET DIAGNOSTICS v_cnt = ROW_COUNT;
	IF v_cnt = 0 THEN
		INSERT INTO pol_mkt_stats(contract_aid,open_interest,num_trades,total_volume,total_fees)
			VALUES(NEW.contract_aid,v_normalized_amount,1,NEW.collateral_amount,NEW.fee_amount);
	END IF;
	-- user statistics
	UPDATE pol_ustats
		SET
			tot_trades = (tot_trades + 1),
			tot_volume = (tot_volume + NEW.collateral_amount),
			tot_fees = (tot_fees + NEW.fee_amount),
			profit = (profit + v_normalized_amount)
		WHERE user_aid = NEW.user_aid;
	GET DIAGNOSTICS v_cnt = ROW_COUNT;
	IF v_cnt = 0 THEN
		INSERT INTO pol_ustats(user_aid,reg_time_stamp,tot_trades,tot_volume,tot_fees,profit)
			VALUES(NEW.user_aid,NEW.time_stamp,1,NEW.collateral_amount,NEW.fee_amount,v_normalized_amount);
	END IF;
	UPDATE pol_ustats_mkt
		SET
			tot_trades = (tot_trades + 1),
			tot_volume = (tot_volume + NEW.collateral_amount),
			tot_fees = (tot_fees + NEW.fee_amount),
			profit = (profit + v_normalized_amount)
		WHERE user_aid=NEW.user_aid AND contract_aid = NEW.contract_aid;
	GET DIAGNOSTICS v_cnt = ROW_COUNT;
	IF v_cnt = 0 THEN
		INSERT INTO pol_ustats_mkt(user_aid,contract_aid,tot_trades,tot_volume,tot_fees,profit)
			VALUES(NEW.user_aid,NEW.contract_aid,1,NEW.collateral_amount,NEW.fee_amount,v_normalized_amount);
	END IF;
	RETURN NEW;
END;
$$ LANGUAGE plpgsql;
CREATE OR REPLACE FUNCTION on_fpmm_buysell_delete() RETURNS trigger AS  $$
DECLARE
BEGIN

	UPDATE pol_mkt_stats
		SET
			num_trades = (num_trades - 1),
			total_volume = (total_volume - OLD.collateral_amount),
			total_fees = (total_fees - OLD.fee_amount)
		WHERE contract_aid = OLD.contract_aid;
	UPDATE pol_ustats
		SET
			tot_trades = (tot_trades - 1),
			tot_volume = (tot_volume - OLD.collateral_amount),
			tot_fees = (tot_fees - OLD.fee_amount),
			profit = (profit - OLD.normalized_amount)
		WHERE user_aid = OLD.user_aid;
	UPDATE pol_ustats_mkt
		SET
			tot_trades = (tot_trades - 1),
			tot_volume = (tot_volume - OLD.collateral_amount),
			tot_fees = (tot_fees - OLD.fee_amount),
			profit = (profit - OLD.normalized_amount)
		WHERE user_aid=OLD.user_aid AND contract_aid = OLD.contract_aid;
	RETURN OLD;
END;
$$ LANGUAGE plpgsql;
CREATE OR REPLACE FUNCTION on_ustats_mkt_insert() RETURNS trigger AS  $$
DECLARE
	v_cnt					NUMERIC;
BEGIN
	UPDATE pol_ustats SET markets_count = (markets_count+1) WHERE user_aid=NEW.user_aid;
	GET DIAGNOSTICS v_cnt = ROW_COUNT;
	IF v_cnt = 0 THEN
		RAISE EXCEPTION 'pol_ustats row doesnt exist for user_aid %',NEW.user_aid;
	END IF;
	RETURN NEW;
END;
$$ LANGUAGE plpgsql;
CREATE OR REPLACE FUNCTION on_ustats_mkt_delete() RETURNS trigger AS  $$
DECLARE
BEGIN
	UPDATE pol_ustats SET markets_count = (markets_count - 1) WHERE user_aid=OLD.user_aid;
	RETURN OLD;
END;
$$ LANGUAGE plpgsql;
