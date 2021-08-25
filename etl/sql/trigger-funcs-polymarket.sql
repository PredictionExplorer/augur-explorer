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
			total_volume = (total_volume + NEW.shares),
			total_liquidity = (total_liquidity + v_normalized_collateral)
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
			tot_liq_given = (tot_liq_given + v_normalized_collateral),
			profit = (profit + v_normalized_collateral)
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
			tot_volume = (tot_volume + NEW.shares),
			profit = (profit + v_normalized_collateral)
		WHERE (user_aid = NEW.funder_aid) AND (contract_aid=NEW.contract_aid);
	GET DIAGNOSTICS v_cnt = ROW_COUNT;
	IF v_cnt = 0 THEN
		INSERT INTO pol_ustats_mkt(user_aid,contract_aid,tot_liq_ops,tot_liq_given,tot_volume,total_liquidity)
			VALUES(NEW.funder_aid,NEW.contract_aid,1,v_normalized_collateral,NEW.shares,v_normalized_collateral);
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
			total_volume = (total_volume - OLD.shares),
			total_liquidity = (total_liquidity - OLD.norm_collateral)
		WHERE contract_aid = OLD.contract_aid;
	UPDATE pol_ustats
		SET
			tot_liq_ops = (tot_liq_ops - 1),
			tot_liq_given = (tot_liq_given - OLD.norm_collateral),
			profit = (profit - OLD.normalized_amount)
		WHERE user_aid = OLD.funder_aid;
	UPDATE pol_ustats_mkt
		SET
			tot_liq_ops = (tot_liq_ops - 1),
			tot_liq_given = (tot_liq_given - OLD.norm_collateral),
			tot_volume = (tot_volume - OLD.shares),
			profit = (profit - OLD.normalized_amount)
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
			total_fees = (total_fees + NEW.fee_amount),
			open_interest = (open_interest + v_normalized_amount)
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
			total_fees = (total_fees - OLD.fee_amount),
			open_interest = (open_interest - OLD.normalized_amount)
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
CREATE OR REPLACE FUNCTION on_pol_pos_split_insert() RETURNS trigger AS  $$
DECLARE
	v_counter				BIGINT;
	v_froms					TEXT[];
	v_tos					TEXT[]; -- the To fields of ERC1155 transfer events
	v_amounts_str			TEXT[];
	v_token_id_hex			TEXT;
	v_amount				DECIMAL;
BEGIN

	IF LENGTH(NEW.tok_ids) > 0 THEN
		v_counter := 1;
		v_amounts_str := STRING_TO_ARRAY(NEW.tok_amounts,',');
		FOREACH v_token_id_hex IN ARRAY STRING_TO_ARRAY(NEW.tok_ids,',')
		LOOP
			v_amount := v_amounts_str[v_counter]::DECIMAL;
			v_froms := STRING_TO_ARRAY(NEW.tok_froms,',');
			v_tos := STRING_TO_ARRAY(NEW.tok_tos,',');
			INSERT INTO pol_pos_tok_ids(parent_split_id,contract_aid,token_id_hex,outcome_idx,token_from,token_to,token_amount)
				VALUES(NEW.id,NEW.stakeholder_aid,v_token_id_hex,v_counter-1,v_froms[v_counter],v_tos[v_counter],v_amount);
			INSERT INTO pol_tok_ids(contract_aid,token_id_hex,outcome_idx)
				VALUES(NEW.stakeholder_aid,v_token_id_hex,v_counter-1)
				ON CONFLICT DO NOTHING;
			v_counter := (v_counter + 1);
		END LOOP;
	END IF;


	RETURN NEW;
END;
$$ LANGUAGE plpgsql;
CREATE OR REPLACE FUNCTION on_pol_pos_split_delete() RETURNS trigger AS  $$
DECLARE
BEGIN

	-- nothing to do, all deletes are made via 'CASCADE'
	RETURN OLD;
END;
$$ LANGUAGE plpgsql;
CREATE OR REPLACE FUNCTION on_pol_pos_merge_insert() RETURNS trigger AS  $$
DECLARE
	v_counter				BIGINT;
	v_froms					TEXT[];
	v_tos					TEXT[]; -- the To fields of ERC1155 transfer events
	v_amounts_str			TEXT[];
	v_token_id_hex			TEXT;
	v_amount				DECIMAL;
BEGIN

	IF LENGTH(NEW.tok_ids) > 0 THEN
		v_counter := 1;
		v_amounts_str := STRING_TO_ARRAY(NEW.tok_amounts,',');
		v_froms := STRING_TO_ARRAY(NEW.tok_froms,',');
		v_tos := STRING_TO_ARRAY(NEW.tok_tos,',');
		FOREACH v_token_id_hex IN ARRAY STRING_TO_ARRAY(NEW.tok_ids,',')
		LOOP
			v_amount := v_amounts_str[v_counter]::DECIMAL;
			INSERT INTO pol_pos_tok_ids(parent_merge_id,contract_aid,token_id_hex,outcome_idx,token_from,token_to,token_amount)
				VALUES(NEW.id,NEW.stakeholder_aid,v_token_id_hex,v_counter-1,v_froms[v_counter],v_tos[v_counter],v_amount);
			INSERT INTO pol_tok_ids(contract_aid,token_id_hex,outcome_idx)
				VALUES(NEW.stakeholder_aid,v_token_id_hex,v_counter-1)
				ON CONFLICT DO NOTHING;
			v_counter := (v_counter + 1);
		END LOOP;
	END IF;

	RETURN NEW;
END;
$$ LANGUAGE plpgsql;
CREATE OR REPLACE FUNCTION on_pol_pos_merge_delete() RETURNS trigger AS  $$
DECLARE
BEGIN

	-- nothing to do, all deletes are made via 'CASCADE'
	RETURN OLD;
END;
$$ LANGUAGE plpgsql;
