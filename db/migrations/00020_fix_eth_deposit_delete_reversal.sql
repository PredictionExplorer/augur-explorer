-- +goose Up
-- Reorg-path fix for on_eth_deposit_delete(), found by the v2 user-staking
-- store suite (TestUserStakingMultiDepositPagination).
--
-- on_eth_deposit_insert() rewrites the row's num_staked_nfts to the
-- inter-deposit delta (which is zero or negative whenever unstakes outpace
-- stakes) and stores the event's pool-wide staked count in accumulated_nfts.
-- The delete trigger guarded its whole reversal with
-- OLD.num_staked_nfts > 0 and divided by that column, so:
--   1. deleting any deposit whose pool shrank since the previous deposit
--      reversed nothing at all (replay then double-counts every staker's
--      total_reward/unclaimed_reward), and
--   2. deleting a deposit whose pool grew divided by the delta instead of
--      the pool size, over-reversing the per-staker rewards.
-- Independently, its per-staker loop applied every staker's token count to
-- every staker (WHERE total_tokens_staked > 0 instead of the loop row's
-- staker), over-reversing whenever more than one wallet was staked.
--
-- The fix reverses from the deposit's own recorded fan-out: per-staker
-- amounts come from cg_staker_deposit (whose amount_to_claim has already
-- been restored to fully-pending by the unstake-event deletions that
-- precede the deposit deletion in tip-down reorg order), and the pool-wide
-- amounts divide by OLD.accumulated_nfts exactly like the insert did.

-- +goose StatementBegin
CREATE OR REPLACE FUNCTION on_eth_deposit_delete() RETURNS trigger AS  $$
DECLARE
	v_mod DECIMAL;
	v_rec RECORD;
BEGIN

	IF OLD.accumulated_nfts > 0 THEN
		v_mod := MOD(OLD.deposit_amount,OLD.accumulated_nfts);
		FOR v_rec IN (SELECT staker_aid,amount_deposited FROM cg_staker_deposit WHERE deposit_id=OLD.deposit_id)
		LOOP
			UPDATE cg_staker_cst
				SET total_reward = (total_reward - v_rec.amount_deposited),
					unclaimed_reward = (unclaimed_reward - v_rec.amount_deposited)
				WHERE staker_aid=v_rec.staker_aid;
		END LOOP;
		UPDATE cg_stake_stats_cst
			SET
				total_reward_amount = (total_reward_amount - (OLD.deposit_amount - v_mod)),
				total_unclaimed_reward = (total_unclaimed_reward - (OLD.deposit_amount - v_mod)),
				num_deposits = (num_deposits - 1),
				total_modulo = (total_modulo - v_mod)
			;
		DELETE FROM cg_staker_deposit WHERE deposit_id=OLD.deposit_id;
		DELETE FROM cg_st_reward WHERE deposit_id=OLD.deposit_id;
	END IF;

	-- Remove corresponding record from cg_prize table
	DELETE FROM cg_prize WHERE round_num=OLD.round_num AND winner_index=0 AND ptype=15;

	RETURN OLD;
END;
$$ LANGUAGE plpgsql;
-- +goose StatementEnd

-- +goose Down
-- Restore the original trigger body from migration 00002.

-- +goose StatementBegin
CREATE OR REPLACE FUNCTION on_eth_deposit_delete() RETURNS trigger AS  $$
DECLARE
	v_amount_per_token DECIMAL;
	v_mod DECIMAL;
	v_rec RECORD;
BEGIN

	IF OLD.num_staked_nfts > 0 THEN
		v_mod := MOD(OLD.deposit_amount,OLD.num_staked_nfts);
		v_amount_per_token := (OLD.deposit_amount - v_mod) / OLD.num_staked_nfts;
		FOR v_rec IN (SELECT count(token_id) AS num_toks,staker_aid FROM cg_staked_token_cst GROUP BY staker_aid)
		LOOP
			UPDATE cg_staker_cst
				SET total_reward = (total_reward -  (v_amount_per_token*v_rec.num_toks)),
					unclaimed_reward = (unclaimed_reward -  (v_amount_per_token*v_rec.num_toks))
				WHERE total_tokens_staked > 0;
		END LOOP;
		UPDATE cg_stake_stats_cst
			SET 
				total_reward_amount = (total_reward_amount - (OLD.deposit_amount - v_mod)),
				total_unclaimed_reward = (total_unclaimed_reward - (OLD.deposit_amount - v_mod)),
				num_deposits = (num_deposits - 1),
				total_modulo = (total_modulo - v_mod)
			;
		DELETE FROM cg_staker_deposit WHERE deposit_id=OLD.deposit_id;
		DELETE FROM cg_st_reward WHERE deposit_id=OLD.deposit_id;
	ELSE   
	END IF;

	-- Remove corresponding record from cg_prize table
	DELETE FROM cg_prize WHERE round_num=OLD.round_num AND winner_index=0 AND ptype=15;

	RETURN OLD;
END;
$$ LANGUAGE plpgsql;
-- +goose StatementEnd
