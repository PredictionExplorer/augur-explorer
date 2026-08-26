-- ============================================================================
-- validate-durations.sql -- Endurance Champion / Chrono-Warrior DURATION
-- validator
--
-- Companion to validate-db.sql (game mechanics) and validate-stats.sql
-- (statistical counters). The stored per-round durations in cg_round_stats
-- (endurance_champion_duration, chrono_warrior_duration) are captured from
-- the V3 contract getter championDurations(roundNum) at MainPrizeClaimed
-- time (internal/indexer/cosmicgame/champion_durations.go); the indexer
-- never derives them from bids. This script recomputes both durations from
-- raw cg_bid timestamps plus the claim timestamp and compares them to the
-- stored values, so an indexer bug, a failed recovery pass, or a wrong
-- contract getter shows up as a violation.
--
-- Definitions (mirroring the contract and the frontend's utils/endurance.ts):
--   * lead stint: a bid holds the lead from its block timestamp until the
--     next bid; the round's last bid holds it until the main-prize claim.
--   * endurance champion duration: the longest single lead stint.
--   * champion lineage: the first stint always seeds the lineage (the first
--     bidder is the initial champion); a later stint must be STRICTLY longer
--     to dethrone (ties keep the incumbent).
--   * chrono-warrior duration: the longest contiguous reign as endurance
--     champion. Champion i's reign runs from start_i + championTime_{i-1}
--     (the moment it out-endured the predecessor) until
--     start_{i+1} + championTime_i (the moment the successor out-endured
--     it); the final champion reigns until the claim.
--   * same-timestamp bids are ordered by evtlog_id (transaction order).
--
-- Pre-V3 deployments do not implement championDurations, so every claimed
-- round legitimately stores 0/0 there; that situation is reported as a WARN
-- (coverage), and the recompute comparisons only apply to rounds where a
-- capture happened. A round with zero stored durations FOLLOWED by a round
-- with captured durations is a recovery gap and is an ERROR.
--
-- Usage (reads the same env vars as cg-etl: DATABASE_URL, or PGSQL_HOST /
-- PGSQL_USERNAME / PGSQL_PASSWORD / PGSQL_DATABASE):
--   scripts/validate-durations.sh
-- or invoke directly:
--   psql "$DATABASE_URL" -f scripts/validate-durations.sql
--
-- Read-only: only TEMPORARY tables are created and the whole run is rolled
-- back. Exits non-zero (raised exception) if any ERROR-severity check fails.
-- ============================================================================

\set ON_ERROR_STOP on
\pset pager off
\echo ''
\echo '=== Cosmic Signature DB validator (champion durations) ==='
\echo ''

BEGIN;
SET LOCAL client_min_messages = warning;

CREATE TEMP TABLE vr (
	id         BIGSERIAL PRIMARY KEY,
	section    TEXT,
	check_name TEXT,
	severity   TEXT,    -- ERROR | WARN
	violations BIGINT,
	details    TEXT
) ON COMMIT DROP;

-- One row per completed round.
CREATE TEMP TABLE claims ON COMMIT DROP AS
SELECT round_num, EXTRACT(EPOCH FROM time_stamp)::BIGINT AS claim_ts
FROM cg_prize_claim;

-- Lead stints: each bid leads until the next bid; the last one until claim.
CREATE TEMP TABLE stints ON COMMIT DROP AS
SELECT b.round_num,
       EXTRACT(EPOCH FROM b.time_stamp)::BIGINT AS start_ts,
       COALESCE(LEAD(EXTRACT(EPOCH FROM b.time_stamp)::BIGINT)
                  OVER (PARTITION BY b.round_num ORDER BY b.time_stamp, b.evtlog_id),
                c.claim_ts) AS end_ts
FROM cg_bid b JOIN claims c USING (round_num);

-- Champion lineage: record-setting stints (strictly longer than every
-- earlier stint in the round; the first stint always qualifies).
CREATE TEMP TABLE champs ON COMMIT DROP AS
SELECT round_num, start_ts, dur
FROM (
	SELECT round_num, start_ts, end_ts - start_ts AS dur,
	       MAX(end_ts - start_ts) OVER (PARTITION BY round_num
	           ORDER BY start_ts, end_ts
	           ROWS BETWEEN UNBOUNDED PRECEDING AND 1 PRECEDING) AS prev_record
	FROM stints
) s
WHERE prev_record IS NULL OR dur > prev_record;

-- Reign segments: champion i reigns from dethroning the predecessor until
-- being dethroned by the successor (the final champion until the claim).
CREATE TEMP TABLE segs ON COMMIT DROP AS
SELECT c.round_num,
       c.start_ts + COALESCE(LAG(c.dur) OVER w, 0)            AS seg_start,
       COALESCE(LEAD(c.start_ts) OVER w + c.dur, cl.claim_ts) AS seg_end
FROM champs c JOIN claims cl USING (round_num)
WINDOW w AS (PARTITION BY c.round_num ORDER BY c.start_ts, c.dur);

CREATE TEMP TABLE recomputed ON COMMIT DROP AS
SELECT c.round_num,
       MAX(c.dur) AS endurance_calc,
       (SELECT MAX(GREATEST(0, s.seg_end - s.seg_start))
        FROM segs s WHERE s.round_num = c.round_num) AS chrono_calc
FROM champs c
GROUP BY c.round_num;

-- Rounds whose durations were captured from the contract (the indexer skips
-- the store when the getter returns 0/0, so captured means "not both zero").
CREATE TEMP TABLE captured ON COMMIT DROP AS
SELECT rs.round_num,
       rs.endurance_champion_duration, rs.chrono_warrior_duration
FROM cg_round_stats rs JOIN claims USING (round_num)
WHERE COALESCE(rs.endurance_champion_duration,0) <> 0
   OR COALESCE(rs.chrono_warrior_duration,0)     <> 0;

-- ===========================================================================
-- SECTION A: coverage
-- ===========================================================================

-- A1: claimed rounds without captured durations. WARN, not ERROR: pre-V3
-- contracts do not implement championDurations, so all-zero is legitimate
-- for old deployments (this is the current state of mainnet).
INSERT INTO vr(section,check_name,severity,violations,details)
SELECT 'A. coverage',
       'claimed rounds have championDurations captured (all-zero = pre-V3 contract or pending recovery)',
       'WARN',COUNT(*),
       LEFT(STRING_AGG('round '||c.round_num,', ' ORDER BY c.round_num),300)
FROM claims c
LEFT JOIN captured cap USING (round_num)
WHERE cap.round_num IS NULL;

-- A2: recovery gap: a round without captured durations while a LATER round
-- has them. RecoverChampionDurations backfills every missing round on ETL
-- startup, so on a V3 deployment this must never happen.
INSERT INTO vr(section,check_name,severity,violations,details)
SELECT 'A. coverage',
       'no uncaptured round precedes a captured round (startup recovery gap)',
       'ERROR',COUNT(*),
       LEFT(STRING_AGG('round '||c.round_num,', ' ORDER BY c.round_num),300)
FROM claims c
LEFT JOIN captured cap USING (round_num)
WHERE cap.round_num IS NULL
  AND EXISTS (SELECT 1 FROM captured c2 WHERE c2.round_num > c.round_num);

-- ===========================================================================
-- SECTION B: stored contract values vs recomputation from cg_bid
-- ===========================================================================

INSERT INTO vr(section,check_name,severity,violations,details)
SELECT 'B. recomputation',
       'endurance_champion_duration matches longest lead stint from cg_bid',
       'ERROR',COUNT(*),
       LEFT(STRING_AGG('round '||cap.round_num||': stored '||cap.endurance_champion_duration
            ||' recomputed '||COALESCE(r.endurance_calc::TEXT,'<no bids>'),'; '),500)
FROM captured cap
LEFT JOIN recomputed r USING (round_num)
WHERE r.endurance_calc IS NULL
   OR cap.endurance_champion_duration <> r.endurance_calc;

INSERT INTO vr(section,check_name,severity,violations,details)
SELECT 'B. recomputation',
       'chrono_warrior_duration matches longest champion reign from cg_bid',
       'ERROR',COUNT(*),
       LEFT(STRING_AGG('round '||cap.round_num||': stored '||cap.chrono_warrior_duration
            ||' recomputed '||COALESCE(r.chrono_calc::TEXT,'<no bids>'),'; '),500)
FROM captured cap
LEFT JOIN recomputed r USING (round_num)
WHERE r.chrono_calc IS NULL
   OR cap.chrono_warrior_duration <> r.chrono_calc;

-- ===========================================================================
-- SECTION C: input-data sanity (applies to every claimed round with bids;
-- protects the recomputation itself from corrupt indexer data)
-- ===========================================================================

-- C1: recomputed durations must fit inside the round window
-- [first bid, claim].
INSERT INTO vr(section,check_name,severity,violations,details)
SELECT 'C. sanity',
       'recomputed durations within [0, claim - first bid]',
       'ERROR',COUNT(*),
       LEFT(STRING_AGG('round '||v.round_num||': endurance '||v.endurance_calc
            ||' chrono '||v.chrono_calc||' window '||v.round_window,'; '),500)
FROM (
	SELECT r.round_num, r.endurance_calc, r.chrono_calc,
	       cl.claim_ts - fb.first_ts AS round_window
	FROM recomputed r
	JOIN claims cl USING (round_num)
	JOIN (SELECT round_num, MIN(start_ts) AS first_ts
	      FROM stints GROUP BY round_num) fb USING (round_num)
	WHERE r.endurance_calc < 0 OR r.chrono_calc < 0
	   OR r.endurance_calc > cl.claim_ts - fb.first_ts
	   OR r.chrono_calc    > cl.claim_ts - fb.first_ts
) v;

-- C2: no bid of a claimed round may be timestamped after the claim (would
-- produce a negative final stint and indicates broken event ordering).
INSERT INTO vr(section,check_name,severity,violations,details)
SELECT 'C. sanity',
       'no bid timestamp exceeds the round''s claim timestamp',
       'ERROR',COUNT(*),
       LEFT(STRING_AGG('round '||v.round_num||' evtlog '||v.evtlog_id,', '),300)
FROM (
	SELECT b.round_num, b.evtlog_id
	FROM cg_bid b JOIN claims c USING (round_num)
	WHERE EXTRACT(EPOCH FROM b.time_stamp)::BIGINT > c.claim_ts
) v;

-- ===========================================================================
-- SECTION D: cg_live_state_updates consistency
-- ===========================================================================

-- D1: the latest live-state snapshot per round must agree with
-- cg_round_stats (both are written from the same contract read in
-- captureChampionDurations; reorg replays append a fresh live row). WARN
-- because InsertLiveStateUpdateIfChanged may legitimately skip inserts.
INSERT INTO vr(section,check_name,severity,violations,details)
SELECT 'D. live state',
       'latest cg_live_state_updates duration matches cg_round_stats',
       'WARN',COUNT(*),
       LEFT(STRING_AGG('round '||v.round_num||' '||v.variable_name
            ||': live '||v.live_value||' stored '||v.stored_value,'; '),500)
FROM (
	WITH last_live AS (
		SELECT DISTINCT ON (variable_name, round_num)
		       variable_name, round_num, new_value
		FROM cg_live_state_updates
		WHERE variable_name IN ('endurance_champion_duration',
		                        'chrono_warrior_duration')
		ORDER BY variable_name, round_num, id DESC)
	SELECT l.round_num, l.variable_name,
	       l.new_value::TEXT AS live_value,
	       CASE l.variable_name
	            WHEN 'endurance_champion_duration'
	                 THEN cap.endurance_champion_duration
	            ELSE cap.chrono_warrior_duration
	       END::TEXT AS stored_value
	FROM last_live l
	JOIN captured cap USING (round_num)
	WHERE l.new_value <> CASE l.variable_name
	                          WHEN 'endurance_champion_duration'
	                               THEN cap.endurance_champion_duration
	                          ELSE cap.chrono_warrior_duration
	                     END
) v;

-- ===========================================================================
-- REPORT
-- ===========================================================================

\echo ''
\echo '--- Recomputed vs stored (per round) ---------------------------------------'
SELECT r.round_num,
       cap.endurance_champion_duration AS endurance_stored,
       r.endurance_calc,
       cap.chrono_warrior_duration     AS chrono_stored,
       r.chrono_calc,
       CASE WHEN cap.round_num IS NULL THEN 'not captured (pre-V3?)'
            WHEN cap.endurance_champion_duration = r.endurance_calc
             AND cap.chrono_warrior_duration     = r.chrono_calc THEN 'match'
            ELSE 'MISMATCH' END AS verdict
FROM recomputed r
LEFT JOIN captured cap USING (round_num)
ORDER BY r.round_num;

\echo ''
\echo '--- Check results ---------------------------------------------------------'
SELECT
	CASE WHEN violations = 0 THEN 'PASS'
	     WHEN severity = 'WARN' THEN 'WARN' ELSE 'FAIL' END AS status,
	section, check_name, violations
FROM vr ORDER BY id;

\echo ''
\echo '--- Failing checks (details) ----------------------------------------------'
SELECT CASE WHEN severity='WARN' THEN 'WARN' ELSE 'FAIL' END AS status,
       section, check_name, violations, details
FROM vr WHERE violations > 0 ORDER BY id;

\echo ''
\echo '--- Summary ----------------------------------------------------------------'
SELECT
	COUNT(*)                                                            AS checks_run,
	COUNT(*) FILTER (WHERE violations = 0)                              AS passed,
	COUNT(*) FILTER (WHERE violations > 0 AND severity = 'ERROR')       AS failed,
	COUNT(*) FILTER (WHERE violations > 0 AND severity = 'WARN')        AS warnings
FROM vr;

SELECT COUNT(*) > 0 AS has_errors
FROM vr WHERE violations > 0 AND severity = 'ERROR' \gset

ROLLBACK;

\if :has_errors
\warn 'DURATION VALIDATION FAILED: at least one ERROR-severity check found violations.'
-- Force a non-zero psql exit status for cron/CI usage.
DO $$ BEGIN RAISE EXCEPTION 'champion duration validation failed'; END $$;
\else
\echo 'DURATION VALIDATION OK: no ERROR-severity violations found.'
\endif
