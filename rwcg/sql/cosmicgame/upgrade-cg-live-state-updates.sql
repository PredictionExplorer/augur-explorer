-- Upgrade: replace the synthetic evt_log "chain sync" mechanism with cg_live_state_updates.
--
-- Part 1 (CREATE TABLE) is safe to apply any time before running the new ETL build.
-- Part 2 and 3 are the production cleanup of the synthetic rows; they are meant to be
-- reviewed and RUN MANUALLY, step by step, after a backup (pg_dump).

-- ============================================================================
-- Part 1: new table for event-less state variables (V3 championDurations etc.)
-- ============================================================================
CREATE TABLE IF NOT EXISTS cg_live_state_updates (
	id            BIGSERIAL PRIMARY KEY,
	variable_name TEXT NOT NULL,
	contract_aid  BIGINT NOT NULL,
	round_num     BIGINT NOT NULL DEFAULT -1,
	block_num     BIGINT NOT NULL,
	time_stamp    TIMESTAMPTZ NOT NULL DEFAULT NOW(),
	new_value     DECIMAL NOT NULL
);
CREATE INDEX IF NOT EXISTS idx_cg_live_state_updates_var
	ON cg_live_state_updates(variable_name, round_num, id);

-- ============================================================================
-- Part 2: AUDIT (read-only). Run these before deleting anything.
-- ============================================================================
-- The synthetic anchor rows all hang off the sentinel transaction:
--   tx_hash = '0xcccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccc'
-- Find its id (5494 in the current production DB) and use it below.
--
-- SELECT id FROM transaction WHERE tx_hash =
--   '0xcccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccc';
--
-- Count the synthetic evt_log rows:
--   SELECT count(*) FROM evt_log WHERE tx_id = 5494;
--
-- Which anchors are actually referenced by admin correction rows (these correction
-- rows will be CASCADE-deleted together with their anchors in Part 3):
--
-- SELECT 'cg_adm_erc20_reward' AS tbl, count(*) FROM cg_adm_erc20_reward
--   WHERE evtlog_id IN (SELECT id FROM evt_log WHERE tx_id = 5494)
-- UNION ALL SELECT 'cg_adm_cst_auclen', count(*) FROM cg_adm_cst_auclen
--   WHERE evtlog_id IN (SELECT id FROM evt_log WHERE tx_id = 5494)
-- UNION ALL SELECT 'cg_adm_cst_auclen_chg_div', count(*) FROM cg_adm_cst_auclen_chg_div
--   WHERE evtlog_id IN (SELECT id FROM evt_log WHERE tx_id = 5494)
-- UNION ALL SELECT 'cg_adm_timeout_claimprize', count(*) FROM cg_adm_timeout_claimprize
--   WHERE evtlog_id IN (SELECT id FROM evt_log WHERE tx_id = 5494)
-- UNION ALL SELECT 'cg_adm_timeout_withdraw', count(*) FROM cg_adm_timeout_withdraw
--   WHERE evtlog_id IN (SELECT id FROM evt_log WHERE tx_id = 5494)
-- UNION ALL SELECT 'cg_adm_price_inc', count(*) FROM cg_adm_price_inc
--   WHERE evtlog_id IN (SELECT id FROM evt_log WHERE tx_id = 5494)
-- UNION ALL SELECT 'cg_adm_time_inc', count(*) FROM cg_adm_time_inc
--   WHERE evtlog_id IN (SELECT id FROM evt_log WHERE tx_id = 5494)
-- UNION ALL SELECT 'cg_adm_prize_microsec', count(*) FROM cg_adm_prize_microsec
--   WHERE evtlog_id IN (SELECT id FROM evt_log WHERE tx_id = 5494)
-- UNION ALL SELECT 'cg_adm_inisecprize', count(*) FROM cg_adm_inisecprize
--   WHERE evtlog_id IN (SELECT id FROM evt_log WHERE tx_id = 5494)
-- UNION ALL SELECT 'cg_adm_eth_auclen', count(*) FROM cg_adm_eth_auclen
--   WHERE evtlog_id IN (SELECT id FROM evt_log WHERE tx_id = 5494)
-- UNION ALL SELECT 'cg_adm_eth_auc_endprice', count(*) FROM cg_adm_eth_auc_endprice
--   WHERE evtlog_id IN (SELECT id FROM evt_log WHERE tx_id = 5494)
-- UNION ALL SELECT 'cg_adm_cst_min_limit', count(*) FROM cg_adm_cst_min_limit
--   WHERE evtlog_id IN (SELECT id FROM evt_log WHERE tx_id = 5494)
-- UNION ALL SELECT 'cg_adm_mkt_reward', count(*) FROM cg_adm_mkt_reward
--   WHERE evtlog_id IN (SELECT id FROM evt_log WHERE tx_id = 5494)
-- UNION ALL SELECT 'cg_delay_duration', count(*) FROM cg_delay_duration
--   WHERE evtlog_id IN (SELECT id FROM evt_log WHERE tx_id = 5494);
--
-- NOTE on the referenced rows: they hold the *initial* (constructor-set) parameter
-- values seeded on the first ETL run — values that never emitted an event. After the
-- CASCADE delete, a "latest value" query on a parameter that was never changed on-chain
-- returns no rows. If you want to preserve those values, export the audit result above
-- before deleting (the new ETL's startup drift check will also log the live chain value
-- of every parameter, so nothing is lost permanently either way).

-- ============================================================================
-- Part 3: DELETE (manual). Run inside a transaction; verify counts before COMMIT.
-- ============================================================================
-- BEGIN;
--   -- Cascades to all cg_adm_* correction rows referencing the anchors:
--   DELETE FROM evt_log WHERE tx_id = 5494;
--   -- The sentinel transaction itself:
--   DELETE FROM transaction WHERE id = 5494;
--   -- Sanity: both must return 0
--   SELECT count(*) FROM evt_log WHERE tx_id = 5494;
--   SELECT count(*) FROM transaction WHERE tx_hash =
--     '0xcccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccc';
-- COMMIT;
--
-- (Block header rows inserted alongside the anchors are real chain headers and may be
-- shared with genuine events — leave them in place.)
--
-- Afterwards: run transaction-collector-verify; it should report RESULT: OK with
-- "missing receipt file: 0" and "SQL tx missing on disk: 0".
