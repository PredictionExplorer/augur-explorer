-- Noote: some triggers are empty because they were discontinued over time. 
-- (will be deactivated completely in the future)
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
			SET num_bids = (num_bids + 1)
			WHERE	(s.market_aid = NEW.market_aid) AND
					(s.eoa_aid = NEW.eoa_aid) AND
					(s.outcome_idx = NEW.outcome_idx);
		GET DIAGNOSTICS v_cnt = ROW_COUNT;
		IF v_cnt = 0 THEN
			INSERT	INTO oostats(market_aid,eoa_aid,outcome_idx,num_bids)
					VALUES(NEW.market_aid,NEW.eoa_aid,NEW.outcome_idx,1)
					ON CONFLICT DO NOTHING;

		END IF;
	END IF;
	IF NEW.otype = 1 THEN
		UPDATE oostats AS s
			SET num_asks = (num_asks + 1)
			WHERE	(s.market_aid = NEW.market_aid) AND
					(s.eoa_aid = NEW.eoa_aid) AND
					(s.outcome_idx = NEW.outcome_idx);
		GET DIAGNOSTICS v_cnt = ROW_COUNT;
		IF v_cnt = 0 THEN
			INSERT	INTO oostats(market_aid,eoa_aid,outcome_idx,num_asks)
					VALUES(NEW.market_aid,NEW.eoa_aid,NEW.outcome_idx,1)
					ON CONFLICT DO NOTHING;
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
			SET num_bids = (num_bids - 1)
			WHERE	(s.market_aid = OLD.market_aid) AND
					(s.eoa_aid = OLD.eoa_aid) AND
					(s.outcome_idx = OLD.outcome_idx);
	END IF;
	IF OLD.otype = 1 THEN
		UPDATE oostats AS s
			SET num_asks = (num_asks - 1)
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
			SET markets_created = (markets_created + 1),
				validity_bonds = (validity_bonds + NEW.validity_bond)
			WHERE s.eoa_aid = NEW.eoa_aid;

	GET DIAGNOSTICS v_cnt = ROW_COUNT;
	IF v_cnt = 0 THEN
		INSERT	INTO ustats(eoa_aid,wallet_aid,markets_created,validity_bonds)
				VALUES(NEW.eoa_aid,NEW.wallet_aid,1,NEW.validity_bond);
	END IF;

	UPDATE main_stats
		SET markets_count = (markets_count + 1), active_count = (active_count +1);
	IF NEW.market_type = 0 THEN
		UPDATE main_stats SET yesno_count = (yesno_count + 1);
	END IF;
	IF NEW.market_type = 1 THEN
		UPDATE main_stats SET categ_count = (categ_count + 1);
	END IF;
	IF NEW.market_type = 2 THEN
		UPDATE main_stats SET scalar_count = (scalar_count + 1);
	END IF;

	UPDATE category set total_markets = (total_markets + 1) WHERE cat_id=NEW.cat_id;

	RETURN NEW;
END;
$$ LANGUAGE plpgsql;
CREATE OR REPLACE FUNCTION on_market_delete() RETURNS trigger AS  $$
DECLARE
	v_cnt numeric;
BEGIN

	UPDATE ustats AS s
			SET markets_created = markets_created - 1,
				validity_bonds = validity_bonds - OLD.validity_bond
			WHERE s.eoa_aid = OLD.eoa_aid;

	GET DIAGNOSTICS v_cnt = ROW_COUNT;
	IF v_cnt = 0 THEN	-- this condition won't be true during normal operation
		INSERT	INTO ustats(eoa_aid,wallet_aid,markets_created)
				VALUES(OLD.eoa_aid,OLD.wallet_aid,0);
	END IF;

	UPDATE main_stats
		SET markets_count = (markets_count - 1), active_count = (active_count - 1);
	IF OLD.market_type = 0 THEN
		UPDATE main_stats SET yesno_count = (yesno_count - 1);
	END IF;
	IF OLD.market_type = 1 THEN
		UPDATE main_stats SET categ_count = (categ_count - 1);
	END IF;
	IF OLD.market_type = 2 THEN
		UPDATE main_stats SET scalar_count = (scalar_count - 1);
	END IF;

	UPDATE category set total_markets = (total_markets - 1) WHERE cat_id=OLD.cat_id;

	RETURN OLD;
END;
$$ LANGUAGE plpgsql;
-- update all the statistics on insertion into 'mktord' table
CREATE OR REPLACE FUNCTION on_mktord_insert() RETURNS trigger AS  $$
DECLARE
	v_cnt numeric;
BEGIN

	-- Make sure user stats record exists
	INSERT INTO mkts_traded(eoa_aid,market_aid) VALUES(NEW.eoa_aid,NEW.market_aid)
		ON CONFLICT DO NOTHING;
	INSERT INTO mkts_traded(eoa_aid,market_aid) VALUES(NEW.eoa_fill_aid,NEW.market_aid)
		ON CONFLICT DO NOTHING;

	-- Update statistics for the Creator of the Order (Seller)
	UPDATE trd_mkt_stats AS s
			SET total_trades = (total_trades + 1),
				volume_traded = (volume_traded + (NEW.price * NEW.amount_filled))
			WHERE	s.eoa_aid = NEW.eoa_aid AND
					s.market_aid = NEW.market_aid;

	GET DIAGNOSTICS v_cnt = ROW_COUNT;
	IF v_cnt = 0 THEN
		INSERT	INTO trd_mkt_stats(eoa_aid,wallet_aid,market_aid,total_trades,volume_traded)
				VALUES(NEW.eoa_aid,NEW.wallet_aid,NEW.market_aid,1,(NEW.price * NEW.amount_filled));
	END IF;
	UPDATE ustats
		SET total_trades = (total_trades + 1),
			volume_traded = (volume_traded + (NEW.price * NEW.amount_filled))
		WHERE eoa_aid=NEW.eoa_aid;

	-- Update statistics for the Filler of the Order (Buyer)
	UPDATE trd_mkt_stats AS s
			SET total_trades = (total_trades + 1),
				volume_traded = (volume_traded + (NEW.price * NEW.amount_filled))
			WHERE	s.eoa_aid = NEW.eoa_fill_aid AND
					s.market_aid = NEW.market_aid;

	GET DIAGNOSTICS v_cnt = ROW_COUNT;
	IF v_cnt = 0 THEN
		INSERT	INTO trd_mkt_stats(eoa_aid,wallet_aid,market_aid,total_trades,volume_traded)
			VALUES(NEW.eoa_fill_aid,NEW.wallet_fill_aid,NEW.market_aid,1,(NEW.price * NEW.amount_filled));
	END IF;
	UPDATE ustats
		SET total_trades = (total_trades + 1),
			volume_traded = (volume_traded + (NEW.price * NEW.amount_filled))
		WHERE eoa_aid=NEW.eoa_fill_aid;

	-- Noote: for Main statistics a trade between 2 users is counted as single trade (i.e its a +1)_
	-- 			but from the point of the User we have +1 for Creator and +1 for Filler (so, its 2 trades)
	UPDATE main_stats SET trades_count = (trades_count + 1);
	UPDATE market SET total_trades = (total_trades + 1) WHERE market_aid = NEW.market_aid;

	RETURN NEW;
END;
$$ LANGUAGE plpgsql;
CREATE OR REPLACE FUNCTION on_mktord_delete() RETURNS trigger AS  $$
DECLARE
	v_cnt numeric;
	v_total_trades bigint;
BEGIN


	-- Update statistics for the Creator of the Order (Seller)
	UPDATE trd_mkt_stats AS s
			SET total_trades = (total_trades - 1),
				volume_traded = (volume_traded - (OLD.price * OLD.amount_filled))
			WHERE	s.eoa_aid = OLD.eoa_aid AND
					s.market_aid = OLD.market_aid
			RETURNING total_trades INTO v_total_trades;

	GET DIAGNOSTICS v_cnt = ROW_COUNT;
	IF v_cnt = 0 THEN	-- this condition won't be true during normal operation
		INSERT	INTO trd_mkt_stats(eoa_aid,wallet_aid,market_aid)
				VALUES(OLD.eoa_aid,OLD.wallet_aid,OLD.market_aid);
	END IF;

	IF v_total_trades = 0 THEN
		DELETE FROM mkts_traded WHERE eoa_aid = OLD.eoa_aid AND market_aid = OLD.market_aid;
	END IF;

	UPDATE ustats
		SET total_trades = (total_trades - 1),
			volume_traded = (volume_traded - (OLD.price * OLD.amount_filled))
		WHERE eoa_aid=OLD.eoa_aid;


	-- Update statistics for the Filler of the Order (Buyer)
	UPDATE trd_mkt_stats AS s
			SET total_trades = (total_trades - 1),
				volume_traded = (volume_traded - (OLD.price * OLD.amount_filled))
			WHERE	s.eoa_aid = OLD.eoa_fill_aid AND
					s.market_aid = OLD.market_aid
			RETURNING total_trades INTO v_total_trades;

	GET DIAGNOSTICS v_cnt = ROW_COUNT;
	IF v_cnt = 0 THEN
		INSERT	INTO trd_mkt_stats(eoa_aid,wallet_aid,market_aid)
				VALUES(OLD.eoa_fill_aid,OLD.wallet_fill_aid,OLD.market_aid);
	END IF;

	IF v_total_trades = 0 THEN
		DELETE FROM mkts_traded WHERE eoa_aid = OLD.eoa_fill_aid AND market_aid = OLD.market_aid;
	END IF;

	UPDATE ustats
		SET total_trades = (total_trades - 1),
			volume_traded = (volume_traded - (OLD.price * OLD.amount_filled))
		WHERE eoa_aid=OLD.eoa_fill_aid;

	--- Update global statistics
	UPDATE main_stats SET trades_count = (trades_count - 1);
	UPDATE market SET total_trades = (total_trades - 1) WHERE market_aid = OLD.market_aid;

	RETURN OLD;
END;
$$ LANGUAGE plpgsql;
-- update all the statistics on insertion into 'trd_mkt_stats' table
CREATE OR REPLACE FUNCTION on_trd_mkt_stats_insert() RETURNS trigger AS  $$
DECLARE
	v_cnt numeric;
BEGIN

	RETURN NEW;
END;
$$ LANGUAGE plpgsql;
CREATE OR REPLACE FUNCTION on_trd_mkt_stats_delete() RETURNS trigger AS  $$
DECLARE
	v_cnt numeric;
BEGIN

	RETURN OLD;
END;
$$ LANGUAGE plpgsql;
CREATE OR REPLACE FUNCTION on_trd_mkt_stats_update() RETURNS trigger AS  $$
DECLARE
	v_cnt numeric;
BEGIN

	-- This function consoldates statistics from different per-User market records 
	-- into a single User statistics record

	-- Begin of update profit loss
	UPDATE ustats AS s
			SET profit_loss = (profit_loss + (NEW.profit_loss - OLD.profit_loss)),
				money_at_stake = (money_at_stake + (NEW.frozen_funds - OLD.frozen_funds))
			WHERE	s.eoa_aid = NEW.eoa_aid;

	GET DIAGNOSTICS v_cnt = ROW_COUNT;
	IF v_cnt = 0 THEN
		--RAISE EXCEPTION 'Corresponding row in ustats ( % - % ) table doesnt exist',NEW.eoa_aid,NEW.wallet_aid;
	END IF;
	-- End of update profit loss

	RETURN NEW;
END;
$$ LANGUAGE plpgsql;
CREATE OR REPLACE FUNCTION on_mktfin_insert() RETURNS trigger AS  $$
DECLARE
	v_validity_bond decimal;
	v_eoa_aid bigint;
BEGIN

	UPDATE market
		SET fin_timestamp = NEW.fin_timestamp,
			winning_payouts=NEW.winning_payouts,
			winning_outcome=NEW.winning_outcome
		WHERE market.market_aid=NEW.market_aid;
	UPDATE main_stats SET active_count = (active_count - 1);
	SELECT eoa_aid,validity_bond FROM market WHERE market_aid = NEW.market_aid INTO v_eoa_aid,v_validity_bond;
	UPDATE ustats SET validity_bonds = validity_bonds - v_validity_bond
		WHERE eoa_aid = v_eoa_aid;

	RETURN NEW;
END;
$$ LANGUAGE plpgsql;
CREATE OR REPLACE FUNCTION on_mktfin_delete() RETURNS trigger AS  $$
DECLARE
	v_validity_bond decimal;
	v_eoa_aid bigint;
BEGIN

	UPDATE main_stats SET active_count = (active_count + 1);
	SELECT eoa_aid,validity_bond FROM market WHERE market_aid = OLD.market_aid INTO v_eoa_aid,v_validity_bond;
	UPDATE ustats SET validity_bonds = validity_bonds + v_validity_bond
		WHERE eoa_aid = v_eoa_aid;

	RETURN OLD;
END;
$$ LANGUAGE plpgsql;
CREATE OR REPLACE FUNCTION on_profit_loss_insert() RETURNS trigger AS  $$
DECLARE
BEGIN

	IF NEW.mktord_id > 0 THEN 
		IF NEW.closed_position = 0 THEN
			UPDATE trd_mkt_stats
				SET frozen_funds = (frozen_funds + NEW.immediate_ff),
					profit_loss = (profit_loss + NEW.immediate_profit)
				WHERE market_aid = NEW.market_aid AND eoa_aid = NEW.eoa_aid;
			UPDATE main_stats SET money_at_stake = (money_at_stake + NEW.immediate_ff);
			UPDATE market
				SET money_at_stake = (money_at_stake + NEW.immediate_ff)
				WHERE market_aid = NEW.market_aid;
		END IF;
		IF NEW.closed_position = 1 THEN
			RAISE EXCEPTION 'You cant insert a record with closed_position = 1, undefined behavior';
		END IF;
	END IF;
	RETURN NEW;
END;
$$ LANGUAGE plpgsql;
CREATE OR REPLACE FUNCTION on_profit_loss_delete() RETURNS trigger AS  $$
DECLARE
BEGIN

	IF OLD.mktord_id > 0 THEN
		UPDATE trd_mkt_stats AS s
			SET frozen_funds = (frozen_funds - OLD.immediate_ff),
				profit_loss = (profit_loss - OLD.immediate_profit)
			WHERE market_aid = OLD.market_aid AND eoa_aid = OLD.eoa_aid;
		UPDATE main_stats SET money_at_stake = (money_at_stake - OLD.immediate_ff);
		UPDATE market
			SET money_at_stake = (money_at_stake - OLD.immediate_ff)
			WHERE market_aid = OLD.market_aid;
	END IF;

	RETURN OLD;
END;
$$ LANGUAGE plpgsql;
CREATE OR REPLACE FUNCTION on_profit_loss_update() RETURNS trigger AS  $$
DECLARE
BEGIN
	IF OLD.immediate_ff != NEW.immediate_ff THEN
		RAISE EXCEPTION 'You cant change immediate frozen funds on update, instead delete the whole block';
	END IF;
	IF OLD.immediate_profit != NEW.immediate_profit THEN
		RAISE EXCEPTION 'You cant change immediate profit on update, instead delete the whole block';
	END IF;
	RETURN NEW;
END;
$$ LANGUAGE plpgsql;


-- reporting triggers
CREATE OR REPLACE FUNCTION on_report_insert() RETURNS trigger AS  $$ --updates volume of the market
DECLARE
	v_cnt numeric;
BEGIN

	-- Update statistics for the Reporter
	UPDATE trd_mkt_stats AS s
			SET total_reports = (total_reports + 1)
			WHERE	s.eoa_aid = NEW.eoa_aid AND
					s.market_aid = NEW.market_aid;
	GET DIAGNOSTICS v_cnt = ROW_COUNT;
	IF v_cnt = 0 THEN
		INSERT	INTO trd_mkt_stats(eoa_aid,wallet_aid,market_aid)
				VALUES(NEW.eoa_aid,NEW.wallet_aid,NEW.market_aid);
	END IF;
	UPDATE ustats
		SET total_reports = (total_reports + 1)
		WHERE	eoa_aid = NEW.eoa_aid;
	IF NEW.is_designated IS TRUE THEN
		UPDATE market
			SET designated_outcome = NEW.outcome_idx
			WHERE market_aid = NEW.market_aid;
		UPDATE ustats
			SET total_designated = (total_designated + 1)
			WHERE eoa_aid = NEW.eoa_aid;
	END IF;
	IF NEW.is_initial THEN
		UPDATE market
			SET initial_outcome = NEW.outcome_idx
			WHERE market_aid = NEW.market_aid;
	END IF;
	RETURN NEW;
END;
$$ LANGUAGE plpgsql;
CREATE OR REPLACE FUNCTION on_report_delete() RETURNS trigger AS  $$
DECLARE
	v_cnt numeric;
BEGIN

	-- Update statistics for the Reporter
	UPDATE trd_mkt_stats AS s
			SET total_reports = (total_reports - 1)
			WHERE	s.eoa_aid = OLD.eoa_aid AND
					s.market_aid = OLD.market_aid;
	GET DIAGNOSTICS v_cnt = ROW_COUNT;
	IF v_cnt = 0 THEN
		INSERT	INTO trd_mkt_stats(eoa_aid,wallet_aid,market_aid)
				VALUES(OLD.eoa_aid,OLD.wallet_aid,OLD.market_aid);
	END IF;
	UPDATE ustats
		SET total_reports = (total_reports - 1)
		WHERE	eoa_aid = OLD.eoa_aid;
	IF OLD.is_designated THEN
		UPDATE market
			SET designated_outcome = -1
			WHERE market_aid = OLD.market_aid;
		UPDATE ustats
			SET total_designated = (total_designated - 1)
			WHERE	eoa_aid = OLD.eoa_aid;
	END IF;
	IF OLD.is_initial THEN
		UPDATE market
			SET initial_outcome = -1
			WHERE market_aid = OLD.market_aid;
	END IF;
	RETURN OLD;
END;
$$ LANGUAGE plpgsql;
CREATE OR REPLACE FUNCTION on_tx_insert() RETURNS trigger AS  $$
DECLARE
BEGIN

	UPDATE block
		SET  num_tx = num_tx + 1
		WHERE block_num=NEW.block_num;

	RETURN NEW;
END;
$$ LANGUAGE plpgsql;
CREATE OR REPLACE FUNCTION on_tx_delete() RETURNS trigger AS  $$
DECLARE
BEGIN

	UPDATE block
		SET  num_tx = num_tx - 1
		WHERE block_num=OLD.block_num;
	RETURN OLD;

END;
$$ LANGUAGE plpgsql;
CREATE OR REPLACE FUNCTION on_dai_transf_insert() RETURNS trigger AS  $$
DECLARE
	v_eoa_aid bigint;
	v_cnt numeric;
	v_augur bool;	-- true if this transfer is made to Augur Wallet account
	v_internal bool;
BEGIN

	v_augur := false;
	SELECT eoa_aid FROM ustats WHERE wallet_aid = NEW.from_aid INTO v_eoa_aid;
	GET DIAGNOSTICS v_cnt = ROW_COUNT;
	IF v_cnt > 0 THEN
		v_augur := true;
	END IF;
	INSERT INTO dai_bal(block_num,tx_id,dai_transf_id,aid,amount,augur,internal)
			VALUES(NEW.block_num,NEW.tx_id,NEW.id,NEW.from_aid,-NEW.amount,v_augur,NEW.from_internal);


	v_augur := false;
	SELECT eoa_aid FROM ustats WHERE wallet_aid = NEW.to_aid INTO v_eoa_aid;
	GET DIAGNOSTICS v_cnt = ROW_COUNT;
	IF v_cnt > 0 THEN
		v_augur := true;
	END IF;
	INSERT INTO dai_bal(block_num,tx_id,dai_transf_id,aid,amount,augur,internal)
			VALUES(NEW.block_num,NEW.tx_id,NEW.id,NEW.to_aid,NEW.amount,v_augur,NEW.to_internal);

	RETURN NEW;
END;
$$ LANGUAGE plpgsql;
CREATE OR REPLACE FUNCTION on_dai_transf_delete() RETURNS trigger AS  $$
DECLARE
BEGIN

	DELETE FROM dai_bal WHERE dai_transf_id = OLD.id;
	RETURN OLD;

END;
$$ LANGUAGE plpgsql;
CREATE OR REPLACE FUNCTION on_dai_bal_insert() RETURNS trigger AS  $$
DECLARE
	v_eoa_aid bigint;
	v_cnt numeric;
	v_augur bool;
BEGIN

	v_augur := false;
	SELECT eoa_aid FROM ustats WHERE wallet_aid = NEW.aid INTO v_eoa_aid;
	GET DIAGNOSTICS v_cnt = ROW_COUNT;
	IF v_cnt > 0 THEN
		v_augur := true;
	END IF;

	IF v_augur THEN
		if NEW.internal IS FALSE THEN
			UPDATE block AS b
				SET cash_flow = (cash_flow + NEW.amount)
				WHERE	b.block_num = NEW.block_num;
		END IF;
	END IF;
	RETURN NEW;
END;
$$ LANGUAGE plpgsql;
CREATE OR REPLACE FUNCTION on_dai_bal_delete() RETURNS trigger AS  $$
DECLARE
	v_augur bool;	-- true if this transfer is made to Augur Wallet account
	v_eoa_aid bigint;
	v_cnt numeric;
BEGIN

	v_augur := false;
	SELECT eoa_aid FROM ustats WHERE wallet_aid = OLD.aid INTO v_eoa_aid;
	GET DIAGNOSTICS v_cnt = ROW_COUNT;
	IF v_cnt > 0 THEN
		v_augur := true;
	END IF;
	IF v_augur THEN
		IF OLD.internal IS FALSE THEN
			UPDATE block AS b
				SET cash_flow = (cash_flow - OLD.amount)
				WHERE	b.block_num = OLD.block_num;
		END IF;
	END IF;
	RETURN OLD;
END;
$$ LANGUAGE plpgsql;
CREATE OR REPLACE FUNCTION on_dai_bal_update() RETURNS trigger AS  $$
DECLARE
	v_augur bool;	-- true if this transfer is made to Augur Wallet account
	v_eoa_aid bigint;
	v_cnt numeric;
BEGIN
	-- Noute: this trigger only calculates block.cash_flow. For another process
	--			use another trigger function
	-- cash_flow calculation starts
	IF NEW.internal != OLD.internal THEN
		RAISE EXCEPTION 'Changing dai_bal.internal field not possible. Delete the whole block';
	END IF;
	IF OLD.augur != NEW.augur THEN -- this update is coming from ustats table
		IF NEW.augur THEN
			IF NEW.internal IS FALSE THEN
				UPDATE block AS b
					SET cash_flow = (cash_flow + NEW.amount)
					WHERE	b.block_num = NEW.block_num;
			END IF;
		END IF;
	END IF;
	-- cash flow calculation ends
	RETURN NEW;
END;
$$ LANGUAGE plpgsql;
CREATE OR REPLACE FUNCTION on_ustats_insert() RETURNS trigger AS  $$
DECLARE
BEGIN

	-- this exception is for debugging purposes, to_do: remove later
	IF NEW.wallet_aid = 0 THEN
		RAISE EXCEPTION 'INSERT into ustats: wallet_aid cant be 0';
	END IF;


	-- The transfers of DAI can happen before wallet is created, so we fix it
	UPDATE dai_bal SET augur = true WHERE aid = NEW.wallet_aid;

	RETURN NEW;
END;
$$ LANGUAGE plpgsql;
CREATE OR REPLACE FUNCTION on_claim_funds_insert() RETURNS trigger AS  $$
DECLARE
BEGIN

	UPDATE trd_mkt_stats
		SET	profit_loss = (profit_loss + NEW.final_profit),
			frozen_funds = (frozen_funds - NEW.unfrozen_funds)
		WHERE market_aid = NEW.market_aid AND eoa_aid = NEW.eoa_aid;
	RETURN NEW;
END;
$$ LANGUAGE plpgsql;
CREATE OR REPLACE FUNCTION on_claim_funds_delete() RETURNS trigger AS  $$
DECLARE
BEGIN

	UPDATE trd_mkt_stats
		SET frozen_funds = (frozen_funds + OLD.unfrozen_funds),
			profit_loss = (profit_loss - OLD.final_profit)
		WHERE market_aid = OLd.market_aid AND eoa_aid = OLD.eoa_aid;
	RETURN OLD;
END;
$$ LANGUAGE plpgsql;
CREATE OR REPLACE FUNCTION on_mkts_traded_insert() RETURNS trigger AS  $$
DECLARE
BEGIN

	UPDATE ustats SET markets_traded = (markets_traded + 1) WHERE eoa_aid = NEW.eoa_aid;
	RETURN NEW;
END;
$$ LANGUAGE plpgsql;
CREATE OR REPLACE FUNCTION on_mkts_traded_delete() RETURNS trigger AS  $$
DECLARE
BEGIN

	UPDATE ustats SET markets_traded = (markets_traded - 1) WHERE eoa_aid = OLD.eoa_aid;
	RETURN OLD;
END;
$$ LANGUAGE plpgsql;
