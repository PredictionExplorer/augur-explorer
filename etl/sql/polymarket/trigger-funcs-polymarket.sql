CREATE OR REPLACE FUNCTION on_fpmm_fund_addrem_insert() RETURNS trigger AS  $$
DECLARE
	v_cnt					NUMERIC;
	v_normalized_collateral DECIMAL;
	v_accumulated_oi		DECIMAL;
BEGIN

	IF NEW.op_type = 0 THEN
		v_normalized_collateral := -NEW.transfer_amount;--SharesAdded (this is the collateral put into pool)
	ELSE 
		v_normalized_collateral := NEW.transfer_amount;	-- collateral taken out of the pool
	END IF;
	UPDATE pol_fund_addrem SET norm_collateral = v_normalized_collateral WHERE id=NEW.id;
	-- market statistics
	UPDATE pol_mkt_stats
		SET
			num_liq_ops = (num_liq_ops + 1),
			open_interest = (open_interest + v_normalized_collateral),
			total_volume = (total_volume + NEW.transfer_amount),
			total_liquidity = (total_liquidity + v_normalized_collateral)
		WHERE contract_aid = NEW.contract_aid
		RETURNING open_interest INTO v_accumulated_oi;
	GET DIAGNOSTICS v_cnt = ROW_COUNT;
	IF v_cnt = 0 THEN
		INSERT INTO pol_mkt_stats(contract_aid,num_liq_ops,open_interest,total_volume)
			VALUES(NEW.contract_aid,1,v_normalized_collateral,NEW.transfer_amount);
		v_accumulated_oi := v_normalized_collateral;
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
			tot_volume = (tot_volume + NEW.transfer_amount),
			profit = (profit + v_normalized_collateral)
		WHERE (user_aid = NEW.funder_aid) AND (contract_aid=NEW.contract_aid);
	GET DIAGNOSTICS v_cnt = ROW_COUNT;
	IF v_cnt = 0 THEN
		INSERT INTO pol_ustats_mkt(user_aid,contract_aid,tot_liq_ops,tot_liq_given,tot_volume)
			VALUES(NEW.funder_aid,NEW.contract_aid,1,v_normalized_collateral,NEW.transfer_amount);
	END IF;
	INSERT INTO pol_oi_hist(evtlog_id,tx_id,user_aid,parent_id,contract_aid,op_type,amount,accum)
		VALUES(
			NEW.evtlog_id,NEW.tx_id,NEW.funder_aid,NEW.id,NEW.contract_aid,
			(NEW.op_type+1),v_normalized_collateral,v_accumulated_oi
		);

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
			total_volume = (total_volume - OLD.transfer_amount),
			total_liquidity = (total_liquidity - OLD.norm_collateral)
		WHERE contract_aid = OLD.contract_aid;
	UPDATE pol_ustats
		SET
			tot_liq_ops = (tot_liq_ops - 1),
			tot_liq_given = (tot_liq_given - OLD.norm_collateral),
			profit = (profit - OLD.norm_collateral)
		WHERE user_aid = OLD.funder_aid;
	UPDATE pol_ustats_mkt
		SET
			tot_liq_ops = (tot_liq_ops - 1),
			tot_liq_given = (tot_liq_given - OLD.norm_collateral),
			tot_volume = (tot_volume - OLD.transfer_amount),
			profit = (profit - OLD.norm_collateral)
		WHERE (user_aid = OLD.funder_aid) AND (contract_aid=OLD.contract_aid);
	RETURN OLD;
END;
$$ LANGUAGE plpgsql;
CREATE OR REPLACE FUNCTION on_fpmm_buysell_insert() RETURNS trigger AS  $$
DECLARE
	v_cnt					NUMERIC;
	v_accum_collateral		DECIMAL;
	v_normalized_amount		DECIMAL;
	v_accumulated_oi		DECIMAL;
BEGIN

	-- first calculate accumulated collateral (which is the same as profit/loss on this market)
	IF NEW.op_type = 0 THEN
		v_normalized_amount := -(NEW.collateral_amount-NEW.fee_amount);
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
			total_volume = (total_volume + (NEW.collateral_amount-NEW.fee_amount)),
			total_fees = (total_fees + NEW.fee_amount),
			open_interest = (open_interest + v_normalized_amount)
		WHERE contract_aid = NEW.contract_aid
		RETURNING open_interest INTO v_accumulated_oi;
	GET DIAGNOSTICS v_cnt = ROW_COUNT;
	IF v_cnt = 0 THEN
		INSERT INTO pol_mkt_stats(contract_aid,open_interest,num_trades,total_volume,total_fees)
			VALUES(NEW.contract_aid,v_normalized_amount,1,NEW.collateral_amount,NEW.fee_amount);
		v_accumulated_oi := (NEW.collateral_amount-NEW.fee_amount);
	END IF;
	-- user statistics
	UPDATE pol_ustats
		SET
			tot_trades = (tot_trades + 1),
			tot_volume = (tot_volume + (NEW.collateral_amount-NEW.fee_amount)),
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
			tot_volume = (tot_volume + (NEW.collateral_amount-NEW.fee_amount)),
			tot_fees = (tot_fees + NEW.fee_amount),
			profit = (profit + v_normalized_amount)
		WHERE user_aid=NEW.user_aid AND contract_aid = NEW.contract_aid;
	GET DIAGNOSTICS v_cnt = ROW_COUNT;
	IF v_cnt = 0 THEN
		INSERT INTO pol_ustats_mkt(user_aid,contract_aid,tot_trades,tot_volume,tot_fees,profit)
			VALUES(NEW.user_aid,NEW.contract_aid,1,NEW.collateral_amount,NEW.fee_amount,v_normalized_amount);
	END IF;

	INSERT INTO pol_oi_hist(evtlog_id,tx_id,user_aid,parent_id,contract_aid,op_type,amount,accum)
		VALUES(
			NEW.evtlog_id,NEW.tx_id,NEW.user_aid,NEW.id,NEW.contract_aid,
			(NEW.op_type+3),v_normalized_amount,v_accumulated_oi
		);
	RETURN NEW;
END;
$$ LANGUAGE plpgsql;
CREATE OR REPLACE FUNCTION on_fpmm_buysell_delete() RETURNS trigger AS  $$
DECLARE
BEGIN

	UPDATE pol_mkt_stats
		SET
			num_trades = (num_trades - 1),
			total_volume = (total_volume - (OLD.collateral_amount-OLD.fee_amount)),
			total_fees = (total_fees - OLD.fee_amount),
			open_interest = (open_interest - OLD.normalized_amount)
		WHERE contract_aid = OLD.contract_aid;
	UPDATE pol_ustats
		SET
			tot_trades = (tot_trades - 1),
			tot_volume = (tot_volume - (OLD.collateral_amount-OLD.fee_amount)),
			tot_fees = (tot_fees - OLD.fee_amount),
			profit = (profit - (OLD.normalized_amount-OLD.fee_amount))
		WHERE user_aid = OLD.user_aid;
	UPDATE pol_ustats_mkt
		SET
			tot_trades = (tot_trades - 1),
			tot_volume = (tot_volume - (OLD.collateral_amount-OLD.fee_amount)),
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
			INSERT INTO pol_tok_id_ops(
				evtlog_id,tx_id,parent_split_id,contract_aid,token_id_hex,condition_id,
				outcome_idx,token_from,token_to,token_amount
			) VALUES(
				NEW.evtlog_id,NEW.tx_id,NEW.id,NEW.stakeholder_aid,v_token_id_hex,NEW.condition_id,
				v_counter-1,v_froms[v_counter],v_tos[v_counter],v_amount
			);
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
	v_id_ins				BIGINT;
BEGIN
	
	IF LENGTH(NEW.tok_ids) > 0 THEN
		v_counter := 1;
		v_amounts_str := STRING_TO_ARRAY(NEW.tok_amounts,',');
		v_froms := STRING_TO_ARRAY(NEW.tok_froms,',');
		v_tos := STRING_TO_ARRAY(NEW.tok_tos,',');
		FOREACH v_token_id_hex IN ARRAY STRING_TO_ARRAY(NEW.tok_ids,',')
		LOOP
			v_amount := v_amounts_str[v_counter]::DECIMAL;
			INSERT INTO pol_tok_id_ops(
				evtlog_id,tx_id,parent_merge_id,contract_aid,token_id_hex,condition_id,
				outcome_idx,token_from,token_to,token_amount
			) VALUES (
				NEW.evtlog_id,NEW.tx_id,NEW.id,NEW.stakeholder_aid,v_token_id_hex,NEW.condition_id,
				v_counter-1,v_froms[v_counter],v_tos[v_counter],v_amount
			) RETURNING id INTO v_id_ins;
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
CREATE OR REPLACE FUNCTION on_pol_pay_redem_insert() RETURNS trigger AS  $$
DECLARE
	v_mkt_mkr_aid			BIGINT;
	v_counter				BIGINT;
	v_froms					TEXT[];
	v_tos					TEXT[]; -- the To fields of ERC1155 transfer events
	v_amounts_str			TEXT[];
	v_token_id_hex			TEXT;
	v_amount				DECIMAL;
	v_accumulated_oi		DECIMAL;
BEGIN
	-- we need to extract contract_aid because Conditional Token
	--	is the contract that emits these events (not market maker contract)
	SELECT
			mkt_mkr_aid
		FROM pol_market
		WHERE condition_id=CONCAT('0x',NEW.condition_id)
		LIMIT 1
		INTO v_mkt_mkr_aid;
	IF v_mkt_mkr_aid IS NULL THEN
		SELECT stakeholder_aid
			FROM pol_pos_split
			WHERE condition_id=NEW.condition_id
			LIMIT 1
			INTO v_mkt_mkr_aid;
			IF v_mkt_mkr_aid IS NULL THEN
				RAISE EXCEPTION
					'Market not registered for condition  % (block %), update pol_markets table',
					NEW.condition_id,NEW.block_num;
				RETURN NEW;
			END IF;
	END IF;

	IF LENGTH(NEW.tok_ids) > 0 THEN
		v_counter := 1;
		v_amounts_str := STRING_TO_ARRAY(NEW.tok_amounts,',');
		FOREACH v_token_id_hex IN ARRAY STRING_TO_ARRAY(NEW.tok_ids,',')
		LOOP
			v_amount := v_amounts_str[v_counter]::DECIMAL;
			v_froms := STRING_TO_ARRAY(NEW.tok_froms,',');
			v_tos := STRING_TO_ARRAY(NEW.tok_tos,',');
			-- Noote: here we do not insert 'contract_aid' field because we don't have it
			INSERT INTO pol_tok_id_ops(
				tx_id,evtlog_id,parent_split_id,token_id_hex,condition_id,
				outcome_idx,token_from,token_to,token_amount
			) VALUES(
				NEW.tx_id,NEW.evtlog_id,NEW.id,v_token_id_hex,NEW.condition_id,
				v_counter-1,v_froms[v_counter],v_tos[v_counter],v_amount
			);
			v_counter := (v_counter + 1);
		END LOOP;
	END IF;

	UPDATE pol_mkt_stats
		SET
			open_interest = (open_interest + NEW.payout)	-- since OI is negative we use +
		WHERE contract_aid = v_mkt_mkr_aid
		RETURNING open_interest INTO v_accumulated_oi;
	IF v_accumulated_oi IS NULL THEN
		v_accumulated_oi := 0;
	END IF;
	UPDATE pol_ustats
		SET
			profit = (profit + NEW.payout)
		WHERE user_aid = NEW.redeemer_aid;
	UPDATE pol_ustats_mkt
		SET
			profit = (profit + NEW.payout)
		WHERE user_aid=NEW.redeemer_aid AND contract_aid = v_mkt_mkr_aid;

	INSERT INTO pol_oi_hist(evtlog_id,tx_id,user_aid,parent_id,contract_aid,op_type,amount,accum)
		VALUES(
			NEW.evtlog_id,NEW.tx_id,NEW.redeemer_aid,NEW.id,v_mkt_mkr_aid,
			5,NEW.payout,v_accumulated_oi
		);

	RETURN NEW;
END;
$$ LANGUAGE plpgsql;
CREATE OR REPLACE FUNCTION on_pol_pay_redem_delete() RETURNS trigger AS  $$
DECLARE
	v_mkt_mkr_aid			BIGINT;
BEGIN
	SELECT
			mkt_mkr_aid
		FROM pol_market
		WHERE condition_id=CONCAT('0x',NEW.condition_id)
		LIMIT 1
		INTO v_mkt_mkr_aid;
	IF v_mkt_mkr_aid IS NULL THEN
		--RAISE EXCEPTION 'Market not registered for condition  %, update pol_markets table',NEW.condition_id;
		RETURN NEW;
	END IF;
	IF v_contract_aid IS NULL THEN
		RAISE EXCEPTION 'Condition preparation for % not found (DELETE)',OLD.condition_id;
		RETURN OLD;
	END IF;
	UPDATE pol_mkt_stats
		SET
			open_interest = (open_interest - OLD.payout)	-- since OI is negative we use +
		WHERE contract_aid = v_mkt_mkr_aid;

	UPDATE pol_ustats
		SET
			profit = (profit - OLD.payout)
		WHERE user_aid = OLD.redeemer_aid;
	UPDATE pol_ustats_mkt
		SET
			profit = (profit - OLD.payout)
		WHERE user_aid=OLD.redeemer_aid AND contract_aid = v_mkt_mkr_aid;
	RETURN OLD;
END;
$$ LANGUAGE plpgsql;
CREATE OR REPLACE FUNCTION pol_insert_search_tokens(p_market_id BIGINT,p_contract_aid BIGINT,p_desc TEXT,p_title TEXT) RETURNS void AS  $$
DECLARE
BEGIN


	INSERT INTO pol_mkt_words(market_id,contract_aid,tok_type,tokens)
		VALUES(p_market_id,p_contract_aid,1,setweight(to_tsvector(coalesce(p_title,'')),'A'))
		ON CONFLICT DO NOTHING;
	INSERT INTO pol_mkt_words(market_id,contract_aid,tok_type,tokens)
		VALUES(p_market_id,p_contract_aid,0,setweight(to_tsvector(coalesce(p_desc,'')),'D'))
		ON CONFLICT DO NOTHING;

END;
$$ LANGUAGE plpgsql;
CREATE OR REPLACE FUNCTION pol_delete_search_tokens(p_market_id bigint,p_contract_aid BIGINT) RETURNS VOID AS  $$
DECLARE
BEGIN

	DELETE from pol_mkt_words WHERE market_id=p_market_id;
END;
$$ LANGUAGE plpgsql;
CREATE OR REPLACE FUNCTION on_pol_market_insert() RETURNS trigger AS  $$
DECLARE
BEGIN

	PERFORM pol_insert_search_tokens(NEW.market_id,NEW.mkt_mkr_aid,NEW.question,NEW.description);
	RETURN NEW;
END;
$$ LANGUAGE plpgsql;
CREATE OR REPLACE FUNCTION on_pol_market_delete() RETURNS trigger AS  $$
DECLARE
BEGIN
	PERFORM delete_search_tokens(OLD.market_id);
	RETURN OLD;
END;
$$ LANGUAGE plpgsql;
CREATE OR REPLACE FUNCTION oi_history_transactions(p_condition_id TEXT,p_usdc_aid BIGINT,p_fpmm_aid BIGINT)
	RETURNS TABLE(bal_id BIGINT) AS $$
DECLARE
BEGIN
	RETURN QUERY 
		SELECT DISTINCT data.bal_id AS bal_id FROM (
			(
				SELECT DISTINCT e20b.id bal_id
				FROM pol_tok_id_ops tops
					CROSS JOIN erc20_bal e20b
				WHERE
					tops.tx_id=e20b.tx_id AND
					tops.condition_id = p_condition_id
			)
			UNION ALL
			(
				SELECT DISTINCT e20b.id bal_id
				FROM erc20_bal e20b
				WHERE
					contract_aid=p_usdc_aid AND aid=p_fpmm_aid
			)
		) AS data
		ORDER BY bal_id;
END;
$$ LANGUAGE plpgsql;
CREATE OR REPLACE FUNCTION oi_history_transaction_ids(p_condition_id TEXT,p_usdc_aid BIGINT,p_fpmm_aid BIGINT)
	RETURNS TABLE(tx_id BIGINT) AS $$
DECLARE
BEGIN
	RETURN QUERY 
		SELECT DISTINCT data.tx_id AS tx_id FROM (
			(
				SELECT DISTINCT e20b.tx_id AS tx_id
				FROM pol_tok_id_ops tops
					CROSS JOIN erc20_bal e20b
				WHERE
					tops.tx_id=e20b.tx_id AND
					tops.condition_id = p_condition_id
			)
			UNION ALL
			(
				SELECT DISTINCT e20b.tx_id AS tx_id
				FROM erc20_bal e20b
				WHERE
					contract_aid=p_usdc_aid AND aid=p_fpmm_aid
			)
		) AS data;
END;
$$ LANGUAGE plpgsql;
