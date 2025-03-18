CREATE OR REPLACE FUNCTION on_bid_insert() RETURNS trigger AS  $$
DECLARE
	v_max_bid				DECIMAL;
	v_cnt					NUMERIC;
BEGIN

	SELECT MAX(bid_price) FROM cg_bid INTO v_max_bid WHERE bidder_aid = NEW.bidder_aid;
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
	ELSE 
		IF NEW.bid_type = 2 THEN
			UPDATE cg_glob_stats SET
				num_bids_cst = (num_bids_cst + 1),
				total_cst_consumed = (total_cst_consumed + NEW.num_cst_tokens);
		END IF;
	END IF;
	UPDATE cg_glob_stats SET cur_num_bids = (cur_num_bids + 1);
	UPDATE cg_round_stats SET 
			total_bids = (total_bids + 1),
			total_cst_in_bids = (total_cst_in_bids + NEW.num_cst_tokens),
			total_eth_in_Bids = (total_eth_in_bids + NEW.bid_price)
	   	WHERE round_num=NEW.round_num;
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

	SELECT MAX(bid_price) FROM cg_bid INTO v_max_bid WHERE bidder_aid = OLD.bidder_aid;
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
	ELSE
		IF OLD.bid_type = 2 THEN
			UPDATE cg_glob_stats SET 
				num_bids_cst = (num_bids_cst - 1),
				total_cst_consumed = (total_cst_consumed - OLD.num_cst_tokens);
		END IF;
	END IF;
	UPDATE cg_glob_stats SET cur_num_bids = (cur_num_bids - 1) WHERE cur_num_bids>0;
	UPDATE cg_round_stats SET 
			total_bids = (total_bids - 1),
			total_cst_in_bids = (total_cst_in_bids + NEW.num_cst_tokens),
			total_eth_in_Bids = (total_eth_in_bids + NEW.bid_price)
		WHERE round_num=OLD.round_num;
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
	SELECT total_nft_donated FROM cg_round_stats WHERE round_num=NEW.round_num INTO v_donated_nfts;
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
	UPDATE cg_erc20_donation_stats SET winner_aid=NEW.winner_aid WHERE round_num=NEW.round_num;
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
	UPDATE cg_erc20_donation_stats SET winner_aid=0 WHERE round_num=OLD.round_num;
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

	IF NEW.round_num <> -1 THEN
		UPDATE cg_prize_claim SET donation_evt_id=NEW.evtlog_id WHERE round_num=NEW.round_num;
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
	IF OLD.round_num <> -1 THEN
		UPDATE cg_prize_claim SET donation_evt_id=0 WHERE round_num=OLD.round_num;
	END IF;

	RETURN OLD;
END;
$$ LANGUAGE plpgsql;
CREATE OR REPLACE FUNCTION on_erc20_donation_insert() RETURNS trigger AS  $$
DECLARE
	v_cnt					NUMERIC;
BEGIN

	UPDATE cg_erc20_donation_stats SET total_amount = (total_amount + NEW.amount) WHERE round_num=NEW.round_num AND token_aid=NEW.token_aid;
	GET DIAGNOSTICS v_cnt = ROW_COUNT;
	IF v_cnt = 0 THEN
		INSERT INTO cg_erc20_donation_stats(token_aid,round_num,total_amount) VALUES (NEW.token_aid,NEW.round_num,NEW.amount);
	END IF;
	UPDATE cg_round_stats SET num_erc20_donations = (num_erc20_donations + 1) WHERE round_num=NEW.round_num;
	GET DIAGNOSTICS v_cnt = ROW_COUNT;
	IF v_cnt = 0 THEN
		INSERT INTO cg_round_stats(round_num,num_erc20_donations) VALUES (NEW.round_num,1);
	END IF;
	UPDATE cg_glob_stats SET total_erc20_donations = (total_erc20_donations + 1);
	RETURN NEW;
END;
$$ LANGUAGE plpgsql;
CREATE OR REPLACE FUNCTION on_erc20_donation_delete() RETURNS trigger AS  $$
DECLARE
	v_cnt					NUMERIC;
BEGIN

	UPDATE cg_erc20_donation_stats SET total_amount = (total_amount - OLD.amount) WHERE round_num=OLD.round_num AND token_aid=OLD.token_aid;
	UPDATE cg_round_stats SET num_erc20_donations = (num_erc20_donations - 1) WHERE round_num=OLD.round_num;
	UPDATE cg_glob_stats SET total_erc20_donations = (total_erc20_donations - 1);
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
CREATE OR REPLACE FUNCTION on_prize_deposit_insert() RETURNS trigger AS  $$
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
CREATE OR REPLACE FUNCTION on_prize_deposit_delete() RETURNS trigger AS  $$
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
	IF NEW.is_staker THEN
		IF NEW.is_rwalk THEN
			UPDATE cg_staker_rwalk SET num_tokens_minted = (num_tokens_minted + 1)
				WHERE staker_aid=NEW.winner_aid;
			GET DIAGNOSTICS v_cnt = ROW_COUNT;
			IF v_cnt = 0 THEN
				INSERT INTO cg_staker_rwalk(staker_aid,num_tokens_minted) VALUES(NEW.winner_aid,1);
			END IF;
			UPDATE cg_stake_stats_rwalk SET total_nft_mints = (total_nft_mints + 1);
		ELSE
			UPDATE cg_staker_cst SET num_tokens_minted = (num_tokens_minted + 1)
				WHERE staker_aid=NEW.winner_aid;
			GET DIAGNOSTICS v_cnt = ROW_COUNT;
			IF v_cnt = 0 THEN
				INSERT INTO cg_staker_cst(staker_aid,num_tokens_minted) VALUES(NEW.winner_aid,1);
			END IF;
			UPDATE cg_stake_stats_cst SET total_nft_mints = (total_nft_mints + 1);
		END IF;
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
	IF OLD.is_staker THEN
		IF OLD.is_rwalk THEN
			UPDATE cg_staker_rwalk SET num_tokens_minted = (num_tokens_minted + 1)
				WHERE staker_aid=OLD.winner_aid;
			UPDATE cg_stake_stats_rwalk SET total_nft_mints = (total_nft_mints - 1);
		ELSE
			UPDATE cg_staker_cst SET num_tokens_minted = (num_tokens_minted + 1)
				WHERE staker_aid=OLD.winner_aid;
			UPDATE cg_stake_stats_cst SET total_nft_mints = (total_nft_mints - 1);
		END IF;
	END IF;
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
CREATE OR REPLACE FUNCTION on_donated_tok_claimed_insert() RETURNS trigger AS  $$
DECLARE
	v_amount					DECIMAL;
BEGIN

	UPDATE cg_erc20_donation_stats
		SET
			total_amount = (total_amount - NEW.amount)
		WHERE round_num = NEW.round_num AND token_aid = NEW.token_aid;
	SELECT total_amount FROM cg_erc20_donation_stats 
		WHERE round_num=NEW.round_num AND token_aid=NEW.token_aid
		INTO v_amount;
	IF v_amount = 0 THEN
		UPDATE cg_erc20_donation_stats SET claimed = 't'
			WHERE round_num = NEW.round_num AND token_aid = NEW.token_aid;
	END IF;

	RETURN NEW;
END;
$$ LANGUAGE plpgsql;
CREATE OR REPLACE FUNCTION on_donated_tok_claimed_delete() RETURNS trigger AS  $$
DECLARE
	v_amount					DECIMAL;
BEGIN

	UPDATE cg_erc20_donation_stats
		SET
			total_amount = (total_amount + OLD.amount)
		WHERE round_num = OLD.round_num AND token_aid = OLD.token_aid;
	SELECT total_amount FROM cg_erc20_donation_stats 
		WHERE round_num=OLD.round_num AND token_aid=OLD.token_aid
		INTO v_amount;
	IF v_amount <> 0 THEN
		UPDATE cg_erc20_donation_stats SET claimed = 'f'
			WHERE round_num = OLD.round_num AND token_aid = OLD.token_aid;
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
CREATE OR REPLACE FUNCTION on_mint_update() RETURNS trigger AS  $$
DECLARE
	v_cnt						NUMERIC;
BEGIN

	RETURN NEW;
END;
$$ LANGUAGE plpgsql;
CREATE OR REPLACE FUNCTION on_prize_withdrawal_insert() RETURNS trigger AS  $$
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
	UPDATE cg_prize_deposit SET claimed=TRUE,withdrawal_id=NEW.evtlog_id WHERE (evtlog_id<NEW.evtlog_id) AND (withdrawal_id=0) AND (winner_aid=NEW.winner_aid);
	UPDATE cg_glob_stats SET total_raffle_eth_withdrawn = (total_raffle_eth_withdrawn + NEW.amount);
	RETURN NEW;
END;
$$ LANGUAGE plpgsql;
CREATE OR REPLACE FUNCTION on_prize_withdrawal_delete() RETURNS trigger AS  $$
DECLARE
BEGIN

	UPDATE cg_raffle_winner_stats
		SET
			withdrawal_sum = (withdrawal_sum - OLD.amount),
			amount_sum = (amount_sum - OLD.amount)
		WHERE winner_aid = OLD.winner_aid;
	UPDATE cg_prize_deposit SET claimed=FALSE,withdrawal_id=0 WHERE withdrawal_id = OLD.evtlog_id;
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
CREATE OR REPLACE FUNCTION on_staker_update() RETURNS trigger AS  $$
DECLARE
BEGIN

--	IF NEW.total_tokens_staked = 0 THEN
--		IF OLD.total_tokens_staked = 1 THEN
--			UPDATE cg_stake_stats_cst SET total_num_stakers = (total_num_stakers - 1);
--		END IF;
--	ELSE
--		IF NEW.total_tokens_staked = 1 THEN
--			IF OLD.total_tokens_staked = 0 THEN
--				UPDATE cg_stake_stats_cst SET total_num_stakers = (total_num_stakers + 1);
--			END IF;
--		END IF;
--	END IF;
	RETURN NEW;
END;
$$ LANGUAGE plpgsql;
CREATE OR REPLACE FUNCTION on_eth_deposit_insert() RETURNS trigger AS  $$
DECLARE
	v_amount_per_token DECIMAL;
	v_mod DECIMAL;
	v_rec RECORD;
	v_cnt						NUMERIC;
	v_prev_num_tokens		INT;
	v_tokens_added			INT;
BEGIN

	IF NEW.num_staked_nfts > 0 THEN
		SELECT num_staked_nfts FROm cg_eth_deposit ORDER BY deposit_id DESC LIMIT 1 INTO v_prev_num_tokens;
		IF v_prev_num_tokens IS NULL THEN
			v_prev_num_tokens:=0;
			v_tokens_added:=NEW.num_staked_nfts;
		ELSE
			v_tokens_added:=NEW.num_staked_nfts-v_prev_num_tokens;
		END IF;
		v_mod := MOD(NEW.amount,NEW.num_staked_nfts);
		v_amount_per_token := (NEW.amount - v_mod) / v_tokens_added);
		UPDATE cg_staker_cst
			SET total_reward = (total_reward + (v_amount_per_token*total_tokens_staked)),
				unclaimed_reward = (unclaimed_reward + (v_amount_per_token*total_tokens_staked))
			WHERE total_tokens_staked > 0;
		UPDATE cg_stake_stats_cst
			SET
				total_reward_amount = (total_reward_amount + (NEW.amount - v_mod)),
				total_unclaimed_reward = (total_unclaimed_reward + (NEW.amount - v_mod)),
				num_deposits = (num_deposits + 1),
				total_modulo = (total_modulo + v_mod)
			;
		FOR v_rec IN (SELECT count(*) AS num_toks,staker_aid FROM cg_staked_token_cst GROUP BY staker_aid)
		LOOP
			INSERT INTO cg_staker_deposit(staker_aid,deposit_id,tokens_staked,amount_to_claim,amount_deposited)
				VALUES(v_rec.staker_aid,NEW.deposit_id,v_rec.num_toks,amount_per_token*v_rec.num_toks,amount_per_token*v_rec.num_toks);
		END LOOP;
		FOR v_rec IN (SELECT token_id,stake_action_id,staker_aid FROM cg_staked_token_cst)
		LOOP
			UPDATE cg_staked_token_cst_rewards 
				SET accumulated_reward = (accumulated_reward + v_amount_per_token)
				WHERE stake_action_id=v_rec.stake_action_id;
			GET DIAGNOSTICS v_cnt = ROW_COUNT;
			IF v_cnt = 0 THEN
				INSERT INTO cg_staked_token_cst_rewards(staker_aid,token_id,stake_action_id,accumulated_reward)
					VALUES(v_rec.staker_aid,v_rec.token_id,v_rec.stake_action_id,v_amount_per_token);
			END IF;
			INSERT INTO cg_st_reward(staker_aid,action_id,token_id,deposit_id,round_num,reward)
				VALUES(v_rec.staker_aid,v_rec.stake_action_id,v_rec.token_id,NEW.deposit_id,NEW.round_num,amount_per_token);
		END LOOP;
	END IF;

	RETURN NEW;
END;
$$ LANGUAGE plpgsql;
CREATE OR REPLACE FUNCTION on_eth_deposit_delete() RETURNS trigger AS  $$
DECLARE
	v_amount_per_token DECIMAL;
	v_mod DECIMAL;
BEGIN

	IF OLD.num_staked_nfts > 0 THEN
		v_mod := MOD(OLD.amount,OLD.num_staked_nfts);
		v_amount_per_token := (OLD.amount - v_mod) / OLD.num_staked_nfts;
		UPDATE cg_staker_cst
			SET total_reward = (total_reward -  (v_amount_per_token*total_tokens_staked)),
				unclaimed_reward = (unclaimed_reward -  (v_amount_per_token*total_tokens_staked))
			WHERE total_tokens_staked > 0;
		UPDATE cg_stake_stats_cst
			SET 
				total_reward_amount = (total_reward_amount - (OLD.amount - v_mod)),
				total_unclaimed_reward = (total_unclaimed_reward - (OLD.amount - v_mod)),
				num_deposits = (num_deposits - 1),
				total_modulo = (total_modulo - v_mod)
			;
		DELETE FROM cg_staker_deposit WHERE deposit_id=OLD.deposit_id;
	ELSE   
	END IF;

	RETURN OLD;
END;
$$ LANGUAGE plpgsql;
CREATE OR REPLACE FUNCTION on_claim_reward_insert() RETURNS trigger AS  $$
DECLARE
	v_cnt						NUMERIC;
BEGIN

	UPDATE cg_stake_stats_cst
		SET total_unclaimed_reward = (total_unclaimed_reward - NEW.reward);
	UPDATE cg_staker_cst
		SET unclaimed_reward = (unclaimed_reward - NEW.reward)
		WHERE staker_aid=NEW.staker_aid;
	RETURN NEW;
END;
$$ LANGUAGE plpgsql;
CREATE OR REPLACE FUNCTION on_claim_reward_delete() RETURNS trigger AS  $$
DECLARE
BEGIN

	UPDATE cg_stake_stats_cst
		SET total_unclaimed_reward = (total_unclaimed_reward + OLD.reward);
	UPDATE cg_staker_cst
		SET unclaimed_reward = (unclaimed_reward + OLD.reward)
		WHERE staker_aid=OLD.staker_aid;

	RETURN OLD;
END;
$$ LANGUAGE plpgsql;
CREATE OR REPLACE FUNCTION on_unstake_action_cst_insert() RETURNS trigger AS  $$
DECLARE
	v_cnt						NUMERIC;
BEGIN

	UPDATE cg_staker_cst
		SET	total_tokens_staked = (total_tokens_staked - 1),
			num_unstake_actions = (num_unstake_actions + 1)
		WHERE staker_aid=NEW.staker_aid;
	UPDATE cg_stake_stats_cst SET total_tokens_staked = (total_tokens_staked - 1);

	DELETE from cg_staked_token_cst WHERE token_id=NEW.token_id AND staker_aid=NEW.staker_aid;
	RETURN NEW;
END;
$$ LANGUAGE plpgsql;
CREATE OR REPLACE FUNCTION on_unstake_action_cst_delete() RETURNS trigger AS  $$
DECLARE
BEGIN

	UPDATE cg_staker_cst
		SET total_tokens_staked = (total_tokens_staked + 1),
			num_unstake_actions = (num_unstake_actions - 1)
		WHERE staker_aid=OLD.staker_aid;
	UPDATE cg_stake_stats_cst SET total_tokens_staked = (total_tokens_staked + 1);
	-- We aren't restoring state here (To Do)

	RETURN OLD;
END;
$$ LANGUAGE plpgsql;
CREATE OR REPLACE FUNCTION on_marketing_rewards_insert() RETURNS trigger AS  $$
DECLARE
	v_cnt						NUMERIC;
BEGIN

	UPDATE cg_glob_stats
		SET	total_mkt_rewards= (total_mkt_rewards + NEW.amount),
			num_mkt_rewards = (num_mkt_rewards + 1)
		;
	RETURN NEW;
END;
$$ LANGUAGE plpgsql;
CREATE OR REPLACE FUNCTION on_marketing_rewards_delete() RETURNS trigger AS  $$
DECLARE
BEGIN

	UPDATE cg_glob_stats
		SET	total_mkt_rewards= (total_mkt_rewards - OLD.amount),
			num_mkt_rewards = (num_mkt_rewards - 1)
		;
	RETURN OLD;
END;
$$ LANGUAGE plpgsql;
CREATE OR REPLACE FUNCTION on_direct_donation_insert() RETURNS trigger AS  $$
DECLARE
	v_cnt						NUMERIC;
BEGIN

	UPDATE cg_glob_stats 
		SET 
			num_direct_donations = (num_direct_donations + 1),
			direct_donations = (direct_donations + NEW.amount);
	UPDATE cg_round_stats
		SET
			donations_round_total = (donations_round_total + NEW.amount),
			donations_round_count = (donations_round_count + 1)
		WHERE NEW.round_num=round_num;
	GET DIAGNOSTICS v_cnt = ROW_COUNT;
	IF v_cnt = 0 THEN
		INSERT INTO cg_round_stats(round_num,donations_round_total,donations_round_count) VALUES (NEW.round_num,NEW.amount,1);
	END IF;
	UPDATE cg_donor
		SET
			total_eth_donated = (total_eth_donated + NEW.amount),
			count_donations = (count_donations + 1)
		WHERE donor_aid = NEW.donor_aid;
	GET DIAGNOSTICS v_cnt = ROW_COUNT;
	IF v_cnt = 0 THEN
		INSERT INTO cg_donor(donor_aid,count_donations,total_eth_donated)
			VALUES(NEW.donor_aid,1,NEW.amount);
	END IF;
	RETURN NEW;
END;
$$ LANGUAGE plpgsql;
CREATE OR REPLACE FUNCTION on_direct_donation_delete() RETURNS trigger AS  $$
DECLARE
BEGIN

	UPDATE cg_glob_stats 
		SET 
			num_direct_donations = (num_direct_donations - 1),
			direct_donations = (direct_donations - OLD.amount);
	UPDATE cg_round_stats
		SET
			donations_round_total = (donations_round_total - OLD.amount),
			donations_round_count = (donations_round_count - 1)
		WHERE OLD.round_num=round_num;
	UPDATE cg_donor
		SET
			total_eth_donated = (total_eth_donated - OLD.amount),
			count_donations = (count_donations - 1)
		WHERE donor_aid = OLD.donor_aid;
	RETURN OLD;
END;
$$ LANGUAGE plpgsql;
CREATE OR REPLACE FUNCTION on_stake_action_rwalk_insert() RETURNS trigger AS  $$
DECLARE
	v_cnt						NUMERIC;
BEGIN
	RETURN NEW;
END;
$$ LANGUAGE plpgsql;
CREATE OR REPLACE FUNCTION on_stake_action_rwalk_delete() RETURNS trigger AS  $$
DECLARE
BEGIN
	RETURN OLD;
END;
$$ LANGUAGE plpgsql;
CREATE OR REPLACE FUNCTION on_unstake_action_rwalk_insert() RETURNS trigger AS  $$
DECLARE
	v_cnt						NUMERIC;
BEGIN
	RETURN NEW;
END;
$$ LANGUAGE plpgsql;
CREATE OR REPLACE FUNCTION on_unstake_action_rwalk_delete() RETURNS trigger AS  $$
DECLARE
BEGIN

	RETURN OLD;
END;
$$ LANGUAGE plpgsql;
CREATE OR REPLACE FUNCTION on_nft_staked_cst_insert() RETURNS trigger AS  $$
DECLARE
	v_cnt						NUMERIC;
	v_active_stakers			INT;
	v_round_num					INT;
BEGIN

	INSERT INTO cg_staked_token_cst(staker_aid,token_id,stake_action_id)
		VALUES(NEW.staker_aid,NEW.token_id,NEW.action_id);
	INSERT INTO cg_staked_token_cst_rewards(staker_aid,token_id,stake_action_id)
		VALUES(NEW.staker_aid,NEW.token_id,NEW.action_id);
	UPDATE cg_staker_cst SET 
			total_tokens_staked = (total_tokens_staked + 1),
			num_stake_actions = (num_stake_actions + 1)
		WHERE staker_aid=NEW.staker_aid;
	GET DIAGNOSTICS v_cnt = ROW_COUNT;
	IF v_cnt = 0 THEN
		INSERT INTO cg_staker_cst(staker_aid,num_stake_actions,total_tokens_staked) VALUES(NEW.staker_aid,1,1);
	END IF;
	UPDATE cg_stake_stats_cst SET total_tokens_staked = (total_tokens_staked + 1);
	SELECT COUNT(*) FROM cg_staker_cst WHERE total_tokens_staked > 0 INTO v_active_stakers;
	IF v_active_stakers IS NOT NULL THEN
		UPDATE cg_stake_stats_cst SET total_num_stakers=v_active_stakers;
	END IF;
	SELECT round_num FROM cg_prize_claim ORDER BY round_num DESC LIMIT 1 INTO v_round_num;
	IF v_round_num IS NOT NULL THEN
		UPDATE cg_nft_staked_cst SET round_num=(v_round_num+1) WHERE id=NEW.id;
	END IF;
	RETURN NEW;
END;
$$ LANGUAGE plpgsql;
CREATE OR REPLACE FUNCTION on_nft_staked_cst_delete() RETURNS trigger AS  $$
DECLARE
	v_active_stakers			INT;
BEGIN

	DELETE FROM cg_staked_token_cst WHERE token_id = OLD.token_id;
	DELETE FROM cg_staked_token_cst_rewards WHERE token_id = OLD.token_id;
	UPDATE cg_staker_cst SET 
			total_tokens_staked = (total_tokens_staked - 1),
			num_stake_actions = (num_stake_actions - 1)
		WHERE staker_aid=OLD.staker_aid;
	UPDATE cg_stake_stats_cst SET total_tokens_staked = (total_tokens_staked - 1);
	SELECT COUNT(*) FROM cg_staker_cst WHERE total_tokens_staked > 0 INTO v_active_stakers;
	IF v_active_stakers IS NOT NULL THEN
		UPDATE cg_stake_stats_cst SET total_num_stakers=v_active_stakers;
	END IF;

	RETURN OLD;
END;
$$ LANGUAGE plpgsql;
CREATE OR REPLACE FUNCTION on_nft_staked_rwalk_insert() RETURNS trigger AS  $$
DECLARE
	v_cnt						NUMERIC;
	v_round_num					INT;
BEGIN

	INSERT INTO cg_staked_token_rwalk(staker_aid,token_id,stake_action_id)
		VALUES(NEW.staker_aid,NEW.token_id,NEW.action_id);
	UPDATE cg_staker_rwalk SET 
			total_tokens_staked = (total_tokens_staked + 1),
			num_stake_actions = (num_stake_actions + 1)
		WHERE staker_aid=NEW.staker_aid;
	GET DIAGNOSTICS v_cnt = ROW_COUNT;
	IF v_cnt = 0 THEN
		INSERT INTO cg_staker_rwalk(staker_aid,total_tokens_staked,num_stake_actions) VALUES(NEW.staker_aid,1,1);
	END IF;
	UPDATE cg_stake_stats_rwalk SET total_tokens_staked = (total_tokens_staked + 1);
	SELECT round_num FROM cg_prize_claim ORDER BY round_num DESC LIMIT 1 INTO v_round_num;
	IF v_round_num IS NOT NULL THEN
		UPDATE cg_nft_staked_rwalk SET round_num=(v_round_num+1) WHERE id=NEW.id;
	END IF;
	RETURN NEW;
END;
$$ LANGUAGE plpgsql;
CREATE OR REPLACE FUNCTION on_nft_staked_rwalk_delete() RETURNS trigger AS  $$
DECLARE
BEGIN

	DELETE FROM cg_staked_token_rwalk WHERE token_id = OLD.token_id;
	UPDATE cg_staker_rwalk SET 
			total_tokens_staked = (total_tokens_staked - 1),
			num_stake_actions = (num_stake_actions - 1)
		WHERE staker_aid=OLD.staker_aid;
	UPDATE cg_stake_stats_rwalk SET total_tokens_staked = (total_tokens_staked - 1);

	RETURN OLD;
END;
$$ LANGUAGE plpgsql;
CREATE OR REPLACE FUNCTION on_nft_unstaked_cst_insert() RETURNS trigger AS  $$
DECLARE
	v_cnt						NUMERIC;
	v_rec RECORD;
	v_active_stakers			INT;
	v_round_num					INT;
BEGIN

	UPDATE cg_staker_cst
		SET	total_tokens_staked = (total_tokens_staked - 1),
			num_unstake_actions = (num_unstake_actions + 1)
		WHERE staker_aid=NEW.staker_aid;
	UPDATE cg_stake_stats_cst SET total_tokens_staked = (total_tokens_staked - 1);

	FOR v_rec IN (SELECT action_id,deposit_id FROM cg_st_reward WHERE action_id=NEW.action_id ORDER BY deposit_id DESC,action_id DESC)
		LOOP
			UPDATE cg_st_reward 
				SET collected = 'T',
			    	is_unstake = 'T'
				WHERE deposit_id=v_rec.deposit_id AND action_id=v_rec.action_id;
		END LOOP;

	DELETE from cg_staked_token_cst WHERE token_id=NEW.token_id AND staker_aid=NEW.staker_aid;
	SELECT COUNT(*) FROM cg_staker_cst WHERE total_tokens_staked > 0 INTO v_active_stakers;
	IF v_active_stakers IS NOT NULL THEN
		UPDATE cg_stake_stats_cst SET total_num_stakers=v_active_stakers;
	END IF;
	SELECT round_num FROM cg_prize_claim ORDER BY round_num DESC LIMIT 1 INTO v_round_num;
	IF v_round_num IS NOT NULL THEN
		UPDATE cg_nft_unstaked_cst SET round_num=(v_round_num+1) WHERE id=NEW.id;
	END IF;
	RETURN NEW;
END;
$$ LANGUAGE plpgsql;
CREATE OR REPLACE FUNCTION on_nft_unstaked_cst_delete() RETURNS trigger AS  $$
DECLARE
	v_rec RECORD;
	v_active_stakers			INT;
BEGIN

	UPDATE cg_staker_cst
		SET total_tokens_staked = (total_tokens_staked + 1),
			num_unstake_actions = (num_unstake_actions - 1)
		WHERE staker_aid=OLD.staker_aid;
	UPDATE cg_stake_stats_cst SET total_tokens_staked = (total_tokens_staked + 1);

	FOR v_rec IN (SELECT action_id,deposit_id FROM cg_st_reward ORDER BY deposit_id DESC,action_id DESC)
		LOOP
			UPDATE cg_st_reward
				SET collected = 'F',
			   		is_unstake = 'F'
				WHERE deposit_id=v_rec.deposit_id AND action_id=v_rec.action_id;
		END LOOP;
	SELECT COUNT(*) FROM cg_staker_cst WHERE total_tokens_staked > 0 INTO v_active_stakers;
	IF v_active_stakers IS NOT NULL THEN
		UPDATE cg_stake_stats_cst SET total_num_stakers=v_active_stakers;
	END IF;

	RETURN OLD;
END;
$$ LANGUAGE plpgsql;
CREATE OR REPLACE FUNCTION on_nft_unstaked_rwalk_insert() RETURNS trigger AS  $$
DECLARE
	v_cnt						NUMERIC;
	v_round_num					INT;
BEGIN

	UPDATE cg_staker_rwalk
		SET	total_tokens_staked = (total_tokens_staked - 1),
			num_unstake_actions = (num_unstake_actions + 1)
		WHERE staker_aid=NEW.staker_aid;
	UPDATE cg_stake_stats_rwalk SET total_tokens_staked = (total_tokens_staked - 1);
	DELETE FROM cg_staked_token_rwalk WHERE token_id=NEW.token_id AND staker_aid=NEW.staker_aid;
	SELECT round_num FROM cg_prize_claim ORDER BY round_num DESC LIMIT 1 INTO v_round_num;
	IF v_round_num IS NOT NULL THEN
		UPDATE cg_nft_unstaked_rwalk SET round_num=(v_round_num+1) WHERE id=NEW.id;
	END IF;
	RETURN NEW;
END;
$$ LANGUAGE plpgsql;
CREATE OR REPLACE FUNCTION on_nft_unstaked_rwalk_delete() RETURNS trigger AS  $$
DECLARE
BEGIN

	UPDATE cg_staker_rwalk
		SET total_tokens_staked = (total_tokens_staked + 1),
			num_unstake_actions = (num_unstake_actions - 1)
		WHERE staker_aid=OLD.staker_aid;
	UPDATE cg_stake_stats_rwalk SET total_tokens_staked = (total_tokens_staked + 1);
	-- We aren't restoring state here (To Do)

	RETURN OLD;
END;
$$ LANGUAGE plpgsql;
CREATE OR REPLACE FUNCTION on_st_reward_insert() RETURNS trigger AS  $$
DECLARE
BEGIN

	IF NEW.collected THEN
		UPDATE cg_staker_deposit 
			SET amount_to_claim = (amount_to_claim - NEW.reward)
			WHERE deposit_id=NEW.deposit_id AND staker_aid=NEW.staker_aid;
	END IF;
	RETURN NEW;
END;
$$ LANGUAGE plpgsql;
CREATE OR REPLACE FUNCTION on_st_reward_delete() RETURNS trigger AS  $$
DECLARE
BEGIN

	IF OLD.collected THEN
		UPDATE cg_staker_deposit 
			SET amount_to_claim = (amount_to_claim + OLD.reward)
			WHERE deposit_id=OLD.deposit_id AND staker_aid=OLD.staker_aid;
	END IF;
	RETURN OLD;
END;
$$ LANGUAGE plpgsql;
CREATE OR REPLACE FUNCTION on_st_reward_update() RETURNS trigger AS  $$
DECLARE
	v_cnt						NUMERIC;
BEGIN
	IF OLD.collected <> NEW.collected THEN
		-- executed when rewardPaid() method is called or unstake() is called
		IF NEW.reward = OLD.reward THEN
			IF NEW.collected THEN
				UPDATE cg_staker_deposit
					SET amount_to_claim=(amount_to_claim - NEW.reward)
					WHERE staker_aid=NEW.staker_aid AND deposit_id=NEW.deposit_id;
			END IF;
			IF NOT NEW.collected THEN
				UPDATE cg_staker_deposit
					SET amount_to_claim=(amount_to_claim + NEW.reward)
					WHERE staker_aid=NEW.staker_aid AND deposit_id=NEW.deposit_id;
			END IF;
		END IF;
	END IF;
	RETURN NEW;
END;
$$ LANGUAGE plpgsql;
