-- +goose Up
-- Reorg-path fix for the staking triggers, found by the cg-etl reorg
-- simulation test (TestReorgRollbackAndReplay).
--
-- on_nft_unstaked_cst_insert() deletes the staker's row from
-- cg_staked_token_cst, but on_nft_unstaked_cst_delete() never restored it
-- (the body carried a "We aren't restoring state here (To Do)" comment).
-- During a chain reorg blocks are deleted tip-down, so the unstake event is
-- deleted before the reward deposit event; when on_eth_deposit_delete() then
-- iterates cg_staked_token_cst to reverse the per-staker reward accounting it
-- finds no rows and reverses nothing. Re-processing the replacement fork then
-- applies the deposit a second time, permanently inflating
-- cg_staker_cst.total_reward / unclaimed_reward. The RandomWalk variant had
-- the same gap for cg_staked_token_rwalk.
--
-- Restoring the staked-token row on unstake-event deletion makes the
-- delete triggers a true inverse of the insert triggers, so
-- rollback-and-replay reproduces the original state exactly.

-- +goose StatementBegin
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

	-- Restore the staked-token row removed by on_nft_unstaked_cst_insert(),
	-- so earlier events (e.g. reward deposits) can be unwound correctly.
	INSERT INTO cg_staked_token_cst(staker_aid,token_id,stake_action_id)
		VALUES(OLD.staker_aid,OLD.token_id,OLD.action_id)
		ON CONFLICT (token_id) DO NOTHING;

	FOR v_rec IN (SELECT action_id,deposit_id FROM cg_st_reward WHERE action_id=OLD.action_id ORDER BY deposit_id DESC,action_id DESC)
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
-- +goose StatementEnd

-- +goose StatementBegin
CREATE OR REPLACE FUNCTION on_nft_unstaked_rwalk_delete() RETURNS trigger AS  $$
DECLARE
	v_active_stakers			INT;
BEGIN

	UPDATE cg_staker_rwalk
		SET total_tokens_staked = (total_tokens_staked + 1),
			num_unstake_actions = (num_unstake_actions - 1)
		WHERE staker_aid=OLD.staker_aid;
	UPDATE cg_stake_stats_rwalk SET total_tokens_staked = (total_tokens_staked + 1);
	SELECT COUNT(*) FROM cg_staker_rwalk WHERE total_tokens_staked > 0 INTO v_active_stakers;
	IF v_active_stakers IS NOT NULL THEN
		UPDATE cg_stake_stats_rwalk SET total_num_stakers=v_active_stakers;
	END IF;

	-- Restore the staked-token row removed by on_nft_unstaked_rwalk_insert().
	INSERT INTO cg_staked_token_rwalk(staker_aid,token_id,stake_action_id)
		VALUES(OLD.staker_aid,OLD.token_id,OLD.action_id)
		ON CONFLICT (token_id) DO NOTHING;

	RETURN OLD;
END;
$$ LANGUAGE plpgsql;
-- +goose StatementEnd

-- +goose Down
-- Restore the original (non-restoring) trigger bodies from migration 00002.

-- +goose StatementBegin
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

	FOR v_rec IN (SELECT action_id,deposit_id FROM cg_st_reward WHERE action_id=OLD.action_id ORDER BY deposit_id DESC,action_id DESC)
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
-- +goose StatementEnd

-- +goose StatementBegin
CREATE OR REPLACE FUNCTION on_nft_unstaked_rwalk_delete() RETURNS trigger AS  $$
DECLARE
	v_active_stakers			INT;
BEGIN

	UPDATE cg_staker_rwalk
		SET total_tokens_staked = (total_tokens_staked + 1),
			num_unstake_actions = (num_unstake_actions - 1)
		WHERE staker_aid=OLD.staker_aid;
	UPDATE cg_stake_stats_rwalk SET total_tokens_staked = (total_tokens_staked + 1);
	SELECT COUNT(*) FROM cg_staker_rwalk WHERE total_tokens_staked > 0 INTO v_active_stakers;
	IF v_active_stakers IS NOT NULL THEN
		UPDATE cg_stake_stats_rwalk SET total_num_stakers=v_active_stakers;
	END IF;

	RETURN OLD;
END;
$$ LANGUAGE plpgsql;
-- +goose StatementEnd
