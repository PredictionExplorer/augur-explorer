CREATE OR REPLACE FUNCTION on_bid_insert() RETURNS trigger AS  $$
DECLARE
	v_max_bid				DECIMAL;
	v_cnt					NUMERIC;
BEGIN

	SELECT MAX(bid_price) FROM cg_bid INTO v_max_bid;
	IF v_max_bid IS NULL THEN
		v_max_bid := 0;
	END IF;
	UPDATE cg_bidder
		SET
			num_bids = (num_bids + 1),
			max_bid	 = v_max_bid
		WHERE bidder_aid = NEW.bidder_aid;
	GET DIAGNOSTICS v_cnt = ROW_COUNT;
	IF v_cnt = 0 THEN
		INSERT INTO cg_bidder(bidder_aid,num_bids,max_bid)
			VALUES(NEW.bidder_aid,1,v_max_bid);
	END IF;
	UPDATE cg_glob_stats SET num_bids = (num_bids + 1);
	GET DIAGNOSTICS v_cnt = ROW_COUNT;
	IF v_cnt = 0 THEN
		RAISE EXCEPTION 'cg_glob_stats table wasnt initialized (no record found)';
	END IF;
	IF NEW.rwalk_nft_id > -1 THEN
		UPDATE cg_glob_stats SET num_rwalk_used = (num_rwalk_used + 1);
	END IF;
	UPDATE cg_glob_stats SET cur_num_bids = (cur_num_bids + 1);
	UPDATE cg_round_stats SET total_bids = (total_bids + 1) WHERE round_num=NEW.round_num;
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

	SELECT MAX(bid_price) FROM cg_bid INTO v_max_bid;
	IF v_max_bid IS NULL THEN
		v_max_bid := 0;
	END IF;
	UPDATE cg_bidder
		SET
			num_bids = (num_bids - 1),
			max_bid	 = v_max_bid
		WHERE bidder_aid = OLD.bidder_aid;
	UPDATE cg_glob_stats SET num_bids = (num_bids - 1);
	GET DIAGNOSTICS v_cnt = ROW_COUNT;
	IF v_cnt = 0 THEN
		RAISE EXCEPTION 'cg_glob_stats table wasnt initialized (no record found)';
	END IF;
	IF OLD.rwalk_nft_id > -1 THEN
		UPDATE cg_glob_stats SET num_rwalk_used = (num_rwalk_used - 1);
	END IF;
	UPDATE cg_glob_stats SET cur_num_bids = (cur_num_bids - 1) WHERE cur_num_bids>0;
	UPDATE cg_round_stats SET total_bids = (total_bids - 1) WHERE round_num=OLD.round_num;
	RETURN OLD;
END;
$$ LANGUAGE plpgsql;
CREATE OR REPLACE FUNCTION on_prize_claim_insert() RETURNS trigger AS  $$
DECLARE
	v_max_prize				DECIMAL;
	v_prizes_sum			DECIMAL;
	v_prizes_count			BIGINT;
	v_donated_nfts			BIGINT;
	v_cnt					NUMERIC;
BEGIN

	SELECT
			MAX(amount),
 			COUNT(*) as prizes_count,
			SUM(amount) as prizes_sum
		FROM cg_prize_claim
		WHERE winner_aid=NEW.winner_aid
		INTO v_max_prize,v_prizes_count,v_prizes_sum;
	IF v_max_prize IS NULL THEN
		v_max_prize := 0;
		v_prizes_sum := 0;
		v_prizes_count := 0;
	END IF;
	SELECT total_nft_donated FROm cg_round_stats WHERE round_num=NEW.prize_num INTO v_donated_nfts;
	IF v_donated_nfts IS NULL THEN
		v_donated_nfts := 0;
	END IF;
	UPDATE cg_winner
		SET
			prizes_count = v_prizes_count,
			max_win_amount	 = v_max_prize,
			prizes_sum = v_prizes_sum,
			unclaimed_nfts = (unclaimed_nfts + v_donated_nfts)
		WHERE winner_aid = NEW.winner_aid;
	GET DIAGNOSTICS v_cnt = ROW_COUNT;
	IF v_cnt = 0 THEN
		INSERT INTO cg_winner(winner_aid,max_win_amount,prizes_count,prizes_sum,unclaimed_nfts)
			VALUES(NEW.winner_aid,v_max_prize,v_prizes_count,v_prizes_sum,v_donated_nfts);
	END IF;
	UPDATE cg_glob_stats SET num_wins = (num_wins + 1);
	GET DIAGNOSTICS v_cnt = ROW_COUNT;
	IF v_cnt = 0 THEN
		RAISE EXCEPTION 'cg_glob_stats table wasnt initialized (no record found)';
	END IF;
	UPDATE cg_glob_stats SET cur_num_bids = 0;
	RETURN NEW;
END;
$$ LANGUAGE plpgsql;
CREATE OR REPLACE FUNCTION on_prize_claim_delete() RETURNS trigger AS  $$
DECLARE
	v_max_prize				DECIMAL;
	v_prizes_sum			DECIMAL;
	v_prizes_count			BIGINT;
	v_cnt					NUMERIC;
BEGIN

	SELECT
			MAX(amount),
 			COUNT(*) as prizes_count,
			SUM(amount) as prizes_sum
		FROM cg_prize_claim
		WHERE winner_aid=OLD.winner_aid
		INTO v_max_prize,v_prizes_count,v_prizes_sum;
	IF v_max_prize IS NULL THEN
		v_max_prize := 0;
		v_prizes_sum := 0;
		v_prizes_count := 0;
	END IF;
	UPDATE cg_winner
		SET
			prizes_count = v_prizes_count,
			max_win_amount	 = v_max_prize,
			prizes_sum = v_prizes_sum
		WHERE winner_aid = OLD.winner_aid;

	UPDATE cg_glob_stats SET num_wins = (num_wins - 1);
	GET DIAGNOSTICS v_cnt = ROW_COUNT;
	IF v_cnt = 0 THEN
		RAISE EXCEPTION 'cg_glob_stats table wasnt initialized (no record found)';
	END IF;
	RETURN OLD;
END;
$$ LANGUAGE plpgsql;
CREATE OR REPLACE FUNCTION on_donation_received_insert() RETURNS trigger AS  $$
DECLARE
	v_cnt						NUMERIC;
	v_cosmicgame_aid			BIGINT;
	v_cosmicgame_addr			TEXT;
BEGIN

	SELECT cosmic_game_addr FROM cg_contracts LIMIT 1 INTO v_cosmicgame_addr;
	IF v_cosmicgame_addr IS NULL THEN
		RAISE EXCEPTION 'CosmicGame contract address is not defined';
	END IF;
	SELECT address_id FROM address WHERE addr=v_cosmicgame_addr INTO v_cosmicgame_aid;
	IF v_cosmicgame_aid IS NULL THEN
		RAISE EXCEPTION 'CosmicGame address id not found in address table';
	END IF;
	IF NEW.donor_aid != v_cosmicgame_aid THEN
		UPDATE cg_glob_stats 
			SET 
				num_vol_donations = (num_vol_donations + 1),
				vol_donations_total = (vol_donations_total + NEW.amount);
	ELSE 
		UPDATE cg_glob_stats 
			SET 
				num_cg_donations = (num_cg_donations + 1),
				cg_donations_total = (cg_donations_total + NEW.amount);
		
	END IF;
	GET DIAGNOSTICS v_cnt = ROW_COUNT;
	IF v_cnt = 0 THEN
		RAISE EXCEPTION 'cg_glob_stats table wasnt initialized (no record found)';
	END IF;

	RETURN NEW;
END;
$$ LANGUAGE plpgsql;
CREATE OR REPLACE FUNCTION on_donation_received_delete() RETURNS trigger AS  $$
DECLARE
	v_cnt					NUMERIC;
	v_cosmicgame_aid			BIGINT;
	v_cosmicgame_addr			TEXT;
BEGIN

	SELECT cosmic_game_addr FROM cg_contracts LIMIT 1 INTO v_cosmicgame_addr;
	IF v_cosmicgame_addr IS NULL THEN
		RAISE EXCEPTION 'CosmicGame contract address is not defined';
	END IF;
	SELECT address_id FROM address WHERE addr=v_cosmicgame_addr INTO v_cosmicgame_aid;
	IF v_cosmicgame_aid IS NULL THEN
		RAISE EXCEPTION 'CosmicGame address id not found in address table';
	END IF;
	IF OLD.donor_aid != v_cosmicgame_aid THEN
		UPDATE cg_glob_stats
			SET
				num_vol_donations = (num_vol_donations - 1),
				vol_donations_total = (vol_donations_total - OLD.amount);
	ELSE
		UPDATE cg_glob_stats
			SET
				num_cg_donations = (num_cg_donations - 1),
				cg_donations_total = (cg_donations_total - OLD.amount);
	END IF;
	GET DIAGNOSTICS v_cnt = ROW_COUNT;
	IF v_cnt = 0 THEN
		RAISE EXCEPTION 'cg_glob_stats table wasnt initialized (no record found)';
	END IF;

	RETURN OLD;
END;
$$ LANGUAGE plpgsql;
CREATE OR REPLACE FUNCTION on_nft_donation_insert() RETURNS trigger AS  $$
DECLARE
	v_cnt					NUMERIC;
BEGIN

	UPDATE cg_nft_stats
		SET
			num_donated = (num_donated + 1)
		WHERE contract_aid = NEW.token_aid;
	GET DIAGNOSTICS v_cnt = ROW_COUNT;
	IF v_cnt = 0 THEN
		INSERT INTO cg_nft_stats(contract_aid,num_donated)
			VALUES(NEW.token_aid,1);
	END IF;
	UPDATE cg_round_stats SET total_nft_donated = (total_nft_donated + 1) WHERE round_num=NEW.round_num;
	GET DIAGNOSTICS v_cnt = ROW_COUNT;
	IF v_cnt = 0 THEN
		INSERT INTO cg_round_stats(round_num,total_nft_donated) VALUES (NEW.round_num,1);
	END IF;
	UPDATE cg_glob_stats SET total_nft_donated = (total_nft_donated + 1);
	RETURN NEW;
END;
$$ LANGUAGE plpgsql;
CREATE OR REPLACE FUNCTION on_nft_donation_delete() RETURNS trigger AS  $$
DECLARE
	v_cnt					NUMERIC;
BEGIN

	UPDATE cg_nft_stats
		SET
			num_donated = (num_donated - 1)
		WHERE contract_aid = OLD.token_aid;
	UPDATE cg_round_stats SET total_nft_donated = (total_nft_donated - 1) WHERE round_num=OLD.round_num;
	UPDATE cg_glob_stats SET total_nft_donated = (total_nft_donated - 1);
	RETURN OLD;
END;
$$ LANGUAGE plpgsql;
CREATE OR REPLACE FUNCTION on_raffle_deposit_insert() RETURNS trigger AS  $$
DECLARE
	v_cnt						NUMERIC;
BEGIN

	UPDATE cg_raffle_winner_stats
		SET
			amount_sum	 = (amount_sum + NEW.amount),
			raffles_count = (raffles_count + 1)
		WHERE winner_aid = NEW.winner_aid;
	GET DIAGNOSTICS v_cnt = ROW_COUNT;
	IF v_cnt = 0 THEN
		INSERT INTO cg_raffle_winner_stats(winner_aid,amount_sum,raffles_count)
			VALUES(NEW.winner_aid,NEW.amount,1);
	END IF;
	UPDATE cg_round_stats
		SET
			total_raffle_eth_deposits = (total_raffle_eth_deposits + NEW.amount)
		WHERE round_num=NEW.round_num;
	GET DIAGNOSTICS v_cnt = ROW_COUNT;
	IF v_cnt = 0 THEN
		INSERT INTO cg_round_stats(round_num,total_raffle_eth_deposits)
			VALUES(NEW.round_num,NEW.amount);
	END IF;
	UPDATE cg_glob_stats SET total_raffle_eth_deposits = (total_raffle_eth_deposits + NEW.amount);
	RETURN NEW;
END;
$$ LANGUAGE plpgsql;
CREATE OR REPLACE FUNCTION on_raffle_deposit_delete() RETURNS trigger AS  $$
DECLARE
BEGIN

	UPDATE cg_raffle_winner_stats
		SET
			amount_sum	 = (amount_sum - OLD.amount),
			raffles_count = (raffles_count - 1)
		WHERE winner_aid = OLD.winner_aid;
	UPDATE cg_round_stats
		SET
			total_raffle_eth_deposits = (total_raffle_eth_deposits - OLD.amount)
		WHERE round_num=OLD.round_num;
	UPDATE cg_glob_stats SET total_raffle_eth_deposits = (total_raffle_eth_deposits - OLD.amount);
	RETURN OLD;
END;
$$ LANGUAGE plpgsql;
CREATE OR REPLACE FUNCTION on_raffle_nft_winner_insert() RETURNS trigger AS  $$
DECLARE
	v_cnt						NUMERIC;
BEGIN

	UPDATE cg_raffle_nft_winner_stats
		SET
			num_won = (num_won + 1)
		WHERE winner_aid = NEW.winner_aid;
	GET DIAGNOSTICS v_cnt = ROW_COUNT;
	IF v_cnt = 0 THEN
		INSERT INTO cg_raffle_nft_winner_stats(winner_aid,num_won)
			VALUES(NEW.winner_aid,1);
	END IF;
	UPDATE cg_round_stats
		SET
			total_raffle_nfts = (total_raffle_nfts + 1)
		WHERE round_num=NEW.round_num;
	GET DIAGNOSTICS v_cnt = ROW_COUNT;
	IF v_cnt = 0 THEN
		INSERT INTO cg_round_stats(round_num,total_raffle_nfts) VALUES(NEW.round_num,1);
	END IF;

	RETURN NEW;
END;
$$ LANGUAGE plpgsql;
CREATE OR REPLACE FUNCTION on_raffle_nft_winner_delete() RETURNS trigger AS  $$
DECLARE
BEGIN

	UPDATE cg_raffle_nft_winner_stats
		SET
			num_won = (num_won - 1)
		WHERE winner_aid = OLD.winner_aid;
	UPDATE cg_round_stats
		SET
			total_raffle_nfts = (total_raffle_nfts - 1)
		WHERE round_num=OLD.round_num;
	RETURN OLD;
END;
$$ LANGUAGE plpgsql;
CREATE OR REPLACE FUNCTION on_erc721transfer_insert() RETURNS trigger AS  $$
DECLARE
	v_cnt						NUMERIC;
BEGIN

	UPDATE cg_mint_event
		SET
			cur_owner_aid = NEW.to_aid
		WHERE token_id=NEW.token_id;

	IF NEW.from_aid = NEW.to_aid THEN -- self transfer
		UPDATE cg_transfer_stats SET erc721_num_transfers = (erc721_num_transfers + 1)
		WHERE user_aid = NEW.from_aid;
		GET DIAGNOSTICS v_cnt = ROW_COUNT;
		IF v_cnt = 0 THEN
			INSERT INTO cg_transfer_stats(user_aid,erc721_num_transfers) VALUES(NEW.from_aid,1);
		END IF;
	ELSE
		UPDATE cg_transfer_stats SET erc721_num_transfers = (erc721_num_transfers + 1)
		WHERE user_aid = NEW.from_aid;
		GET DIAGNOSTICS v_cnt = ROW_COUNT;
		IF v_cnt = 0 THEN
			INSERT INTO cg_transfer_stats(user_aid,erc721_num_transfers) VALUES(NEW.from_aid,1);
		END IF;
		UPDATE cg_transfer_stats SET erc721_num_transfers = (erc721_num_transfers + 1)
		WHERE user_aid = NEW.to_aid;
		GET DIAGNOSTICS v_cnt = ROW_COUNT;
		IF v_cnt = 0 THEN
			INSERT INTO cg_transfer_stats(user_aid,erc721_num_transfers) VALUES(NEW.to_aid,1);
		END IF;
	END IF;

	RETURN NEW;
END;
$$ LANGUAGE plpgsql;
CREATE OR REPLACE FUNCTION on_erc721transfer_delete() RETURNS trigger AS  $$
DECLARE
BEGIN

	IF OLD.from_aid = OLD.to_aid THEN -- self transfer
		UPDATE cg_transfer_stats SET erc721_num_transfers = (erc721_num_transfers - 1)
		WHERE user_aid = OLD.from_aid;
	ELSE
		UPDATE cg_transfer_stats SET erc721_num_transfers = (erc721_num_transfers - 1)
		WHERE user_aid = OLD.from_aid;
		UPDATE cg_transfer_stats SET erc721_num_transfers = (erc721_num_transfers - 1)
		WHERE user_aid = OLD.to_aid;
	END IF;
	RETURN OLD;
END;
$$ LANGUAGE plpgsql;
CREATE OR REPLACE FUNCTION on_erc20_transfer_insert() RETURNS trigger AS  $$
DECLARE
	v_cnt						NUMERIC;
	v_from_addr					TEXT;
	v_to_addr					TEXT;
BEGIN


	SELECT addr FROM address WHERE address_id=NEW.from_aid INTO v_from_addr;
	SELECT addr FROM address WHERE address_id=NEW.to_aid INTO v_to_addr;
	--- update From balance
	IF v_to_addr = '0x0000000000000000000000000000000000000000' THEN -- burn
		UPDATE cg_costok_owner SET cur_balance = (cur_balance - NEW.value)
		WHERE owner_aid=NEW.from_aid;
		GET DIAGNOSTICS v_cnt = ROW_COUNT;
		IF v_cnt = 0 THEN
			INSERT INTO cg_costok_owner(owner_aid,cur_balance) VALUES(NEW.from_aid,-NEW.value);
		END IF;
	ELSE 
		IF v_from_addr != '0x0000000000000000000000000000000000000000' THEN --regular transfer
			UPDATE cg_costok_owner SET cur_balance = (cur_balance - NEW.value)
			WHERE owner_aid = NEW.from_aid;
			GET DIAGNOSTICS v_cnt = ROW_COUNT;
			IF v_cnt = 0 THEN
				INSERT INTO cg_costok_owner(owner_aid,cur_balance) VALUES(NEW.from_aid,-NEW.value);
			END IF;
			UPDATE cg_costok_owner SET cur_balance = (cur_balance + NEW.value)
			WHERE owner_aid = NEW.to_aid;
			GET DIAGNOSTICS v_cnt = ROW_COUNT;
			IF v_cnt = 0 THEN
				INSERT INTO cg_costok_owner(owner_aid,cur_balance) VALUES(NEW.to_aid,NEW.value);
			END IF;
		ELSE  -- mint
			UPDATE cg_costok_owner SET cur_balance = (cur_balance + NEW.value)
			WHERE owner_aid=NEW.to_aid;
			GET DIAGNOSTICS v_cnt = ROW_COUNT;
			IF v_cnt = 0 THEN
				INSERT INTO cg_costok_owner(owner_aid,cur_balance) VALUES(NEW.to_aid,NEW.value);
			END IF;
		END IF;
	END IF;
	IF NEW.from_aid = NEW.to_aid THEN -- self transfer
		UPDATE cg_transfer_stats SET erc20_num_transfers = (erc20_num_transfers + 1)
		WHERE user_aid = NEW.from_aid;
		GET DIAGNOSTICS v_cnt = ROW_COUNT;
		IF v_cnt = 0 THEN
			INSERT INTO cg_transfer_stats(user_aid,erc20_num_transfers) VALUES(NEW.from_aid,1);
		END IF;
	ELSE
		UPDATE cg_transfer_stats SET erc20_num_transfers = (erc20_num_transfers + 1)
		WHERE user_aid = NEW.from_aid;
		GET DIAGNOSTICS v_cnt = ROW_COUNT;
		IF v_cnt = 0 THEN
			INSERT INTO cg_transfer_stats(user_aid,erc20_num_transfers) VALUES(NEW.from_aid,1);
		END IF;
		UPDATE cg_transfer_stats SET erc20_num_transfers = (erc20_num_transfers + 1)
		WHERE user_aid = NEW.to_aid;
		GET DIAGNOSTICS v_cnt = ROW_COUNT;
		IF v_cnt = 0 THEN
			INSERT INTO cg_transfer_stats(user_aid,erc20_num_transfers) VALUES(NEW.to_aid,1);
		END IF;
	END IF;

	RETURN NEW;
END;
$$ LANGUAGE plpgsql;
CREATE OR REPLACE FUNCTION on_erc20_transfer_delete() RETURNS trigger AS  $$
DECLARE
	v_from_addr					TEXT;
	v_to_addr					TEXT;
BEGIN

	SELECT addr FROM address WHERE address_id=OLD.to_aid INTO v_from_addr;
	SELECT addr FROM address WHERE address_id=OLD.to_aid INTO v_to_addr;
	--- update From balance
	IF v_to_addr = '0x0000000000000000000000000000000000000000' THEN -- burn
		UPDATE cg_costok_owner SET cur_balance = (cur_balance + OLD.value)
		WHERE owner_aid=OLD.from_aid;
	ELSE 
		IF v_from_addr != '0x0000000000000000000000000000000000000000' THEN --regular transfer
			UPDATE cg_costok_owner SET cur_balance = (cur_balance + OLD.value)
			WHERE owner_aid = OLD.from_aid;
			UPDATE cg_costok_owner SET cur_balance = (cur_balance - OLD.value)
			WHERE owner_aid = OLD.to_aid;
		ELSE  -- mint
			UPDATE cg_costok_owner SET cur_balance = (cur_balance - OLD.value)
			WHERE owner_aid=OLD.to_aid;
		END IF;
	END IF;
	IF OLD.from_aid = OLD.to_aid THEN -- self transfer
		UPDATE cg_transfer_stats SET erc20_num_transfers = (erc20_num_transfers - 1)
		WHERE user_aid = OLD.from_aid;
	ELSE
		UPDATE cg_transfer_stats SET erc20_num_transfers = (erc20_num_transfers - 1)
		WHERE user_aid = OLD.from_aid;
		UPDATE cg_transfer_stats SET erc20_num_transfers = (erc20_num_transfers - 1)
		WHERE user_aid = OLD.to_aid;
	END IF;
	RETURN OLD;
END;
$$ LANGUAGE plpgsql;
CREATE OR REPLACE FUNCTION on_donated_nft_claimed_insert() RETURNS trigger AS  $$
DECLARE
	v_cnt						NUMERIC;
BEGIN

	UPDATE cg_winner
		SET
			unclaimed_nfts = (unclaimed_nfts - 1)
		WHERE winner_aid = NEW.winner_aid;
	GET DIAGNOSTICS v_cnt = ROW_COUNT;
	IF v_cnt = 0 THEN
		INSERT INTO cg_winner(winner_aid,unclaimed_nfts) VALUES(NEW.winner_aid,1);
	END IF;
	RETURN NEW;
END;
$$ LANGUAGE plpgsql;
CREATE OR REPLACE FUNCTION on_donated_nft_claimed_delete() RETURNS trigger AS  $$
DECLARE
BEGIN

	UPDATE cg_winner
		SET
			unclaimed_nfts = (unclaimed_nfts + 1)
		WHERE winner_aid = OLD .winner_aid;
	RETURN OLD;
END;
$$ LANGUAGE plpgsql;
CREATE OR REPLACE FUNCTION on_mint_insert() RETURNS trigger AS  $$
DECLARE
	v_cnt						NUMERIC;
BEGIN

	UPDATE cg_winner
		SET
			tokens_count = (tokens_count + 1)
		WHERE winner_aid = NEW.owner_aid;
	GET DIAGNOSTICS v_cnt = ROW_COUNT;
	IF v_cnt = 0 THEN
		INSERT INTO cg_winner(winner_aid,tokens_count) VALUES(NEW.owner_aid,1);
	END IF;
	UPDATE cg_glob_stats SET num_mints = (num_mints + 1);
	RETURN NEW;
END;
$$ LANGUAGE plpgsql;
CREATE OR REPLACE FUNCTION on_mint_delete() RETURNS trigger AS  $$
DECLARE
BEGIN

	UPDATE cg_winner
		SET
			tokens_count = (tokens_count - 1)
		WHERE winner_aid = OLD.owner_aid;
	UPDATE cg_glob_stats SET num_mints = (num_mints - 1);
	RETURN OLD;
END;
$$ LANGUAGE plpgsql;
CREATE OR REPLACE FUNCTION on_raffle_withdrawal_insert() RETURNS trigger AS  $$
DECLARE
	v_cnt						NUMERIC;
BEGIN

	UPDATE cg_raffle_winner_stats
		SET
			withdrawal_sum = (withdrawal_sum + NEW.amount),	-- data for historical purposes
			amount_sum = (amount_sum - NEW.amount)			-- current amount available to withdraw
		WHERE winner_aid = NEW.winner_aid;
	GET DIAGNOSTICS v_cnt = ROW_COUNT;
	IF v_cnt = 0 THEN
		INSERT INTO cg_raffle_winner_stats(winner_aid,withdrawal_sum) VALUES(NEW.winner_aid,NEW.amount);
	END IF;
	UPDATE cg_raffle_deposit SET claimed=TRUE,withdrawal_id=NEW.evtlog_id WHERE (evtlog_id<NEW.evtlog_id) AND (withdrawal_id=0) AND (winner_aid=NEW.winner_aid);
	UPDATE cg_glob_stats SET total_raffle_eth_withdrawn = (total_raffle_eth_withdrawn + NEW.amount);
	RETURN NEW;
END;
$$ LANGUAGE plpgsql;
CREATE OR REPLACE FUNCTION on_raffle_withdrawal_delete() RETURNS trigger AS  $$
DECLARE
BEGIN

	UPDATE cg_raffle_winner_stats
		SET
			withdrawal_sum = (withdrawal_sum - OLD.amount),
			amount_sum = (amount_sum - OLD.amount)
		WHERE winner_aid = OLD.winner_aid;
	UPDATE cg_raffle_deposit SET claimed=FALSE,withdrawal_id=0 WHERE withdrawal_id = OLD.evtlog_id;
	UPDATE cg_glob_stats SET total_raffle_eth_withdrawn = (total_raffle_eth_withdrawn - OLD.amount);
	RETURN OLD;
END;
$$ LANGUAGE plpgsql;
CREATE OR REPLACE FUNCTION on_token_name_insert() RETURNS trigger AS  $$
DECLARE
BEGIN

	UPDATE cg_mint_event SET token_name = NEW.token_name WHERE token_id=NEW.token_id;
	RETURN NEW;
END;
$$ LANGUAGE plpgsql;
CREATE OR REPLACE FUNCTION on_token_name_delete() RETURNS trigger AS  $$
DECLARE
BEGIN

	UPDATE cg_mint_event SET token_name = '' WHERE token_id = OLD.token_id;
	RETURN OLD;
END;
$$ LANGUAGE plpgsql;
CREATE OR REPLACE FUNCTION on_donation_sent_insert() RETURNS trigger AS  $$
DECLARE
	v_cnt						NUMERIC;
BEGIN

	UPDATE cg_glob_stats 
		SET 
			num_withdrawals = (num_withdrawals + 1),
			sum_withdrawals = (sum_withdrawals + NEW.amount);

	RETURN NEW;
END;
$$ LANGUAGE plpgsql;
CREATE OR REPLACE FUNCTION on_donation_sent_delete() RETURNS trigger AS  $$
DECLARE
BEGIN

	UPDATE cg_glob_stats
		SET
			num_withdrawals = (num_withdrawals - 1),
			sum_withdrawals = (sum_withdrawals - OLD.amount);

	RETURN OLD;
END;
$$ LANGUAGE plpgsql;
