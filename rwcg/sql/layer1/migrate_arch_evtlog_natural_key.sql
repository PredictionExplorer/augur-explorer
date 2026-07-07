-- Migrate arch_evtlog from PRIMARY KEY (evt_id) to PRIMARY KEY (tx_hash, log_index).
-- Run on a copy / after backup. Requires live evt_log + transaction for backfill.
--
-- Rationale: evt_id is database-specific; (tx_hash, log_index) matches the chain and
-- stays valid if evt_log is reindexed or restored from another DB.
--
-- Steps:
--   1. Backfill log_index from evt_log where evt_id joins.
--   2. Delete archive rows that cannot be joined (legacy imports from another DB).
--   3. Dedupe any (tx_hash, log_index) collisions keeping smallest evt_id.
--   4. Replace primary key.

BEGIN;

ALTER TABLE arch_evtlog ADD COLUMN IF NOT EXISTS log_index INT;

UPDATE arch_evtlog ae
SET log_index = e.log_index
FROM evt_log e
WHERE e.id = ae.evt_id;

-- Rows from old cross-DB exports with no matching live log
DELETE FROM arch_evtlog WHERE log_index IS NULL;

ALTER TABLE arch_evtlog ALTER COLUMN log_index SET NOT NULL;

-- Colliding duplicates (should be rare)
DELETE FROM arch_evtlog a
WHERE EXISTS (
    SELECT 1 FROM arch_evtlog b
    WHERE b.tx_hash = a.tx_hash
      AND b.log_index = a.log_index
      AND b.evt_id < a.evt_id
);

ALTER TABLE arch_evtlog DROP CONSTRAINT IF EXISTS arch_evtlog_pkey;

ALTER TABLE arch_evtlog ALTER COLUMN evt_id DROP NOT NULL;

ALTER TABLE arch_evtlog ADD PRIMARY KEY (tx_hash, log_index);

CREATE INDEX IF NOT EXISTS idx_arch_evtlog_evt_id ON arch_evtlog(evt_id);

COMMIT;
