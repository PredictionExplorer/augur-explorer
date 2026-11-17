-- ChaosZero domain fixture: one settled round (0) with the full lifecycle
-- (initialize, fund, trades on three questions, split, merge, answer
-- observation, settlement, escrow, claim, vault payout) plus one open round
-- (1) funded partly by round 0's carry, and ERC-1155 position transfers and
-- approvals on top.
--
-- Deterministic conventions follow 01_layer1.sql: digit-only addresses,
-- block N at timestamp 1767225600 + (N-100)*100, tx 11NN in its own block,
-- evt_log ids 5301+. CST amounts are raw 1e18-scale DECIMALs chosen so the
-- API's /1e18 FLOAT8 scaling yields short exact doubles (12, 12.5, 0.75).
--
-- Realized cash flows (traders list order is profit DESC, addr ASC):
--   alice (21): buy 12, claim 17            -> profit +5
--   dave  (24): buy 4                       -> profit -4
--   bob   (22): buy 5                       -> profit -5
--   carol (23): split 20, sell 7, merge 5   -> profit -8
-- Round 0 pools: 100 funded + 17 bought - 7 sold + 20 split - 5 merged = 125,
-- settled as 35 escrowed ([20,0,15,0,...]) + 90 carried. Round 1 is funded
-- 10 donated + 90 carry = 100, plus dave's 4 CST buy -> 104 open liquidity.

-- ---------------------------------------------------------------------------
-- Contract addresses (fixture deployment); the migration-seeded production
-- row is repointed so address joins and /api/chaoszero/contracts agree.
-- ---------------------------------------------------------------------------
INSERT INTO address(address_id, block_num, tx_id, addr) VALUES
  (31, 100, 1001, '0x3100000000000000000000000000000000000031'), -- CZ market
  (32, 100, 1001, '0x3200000000000000000000000000000000000032'), -- CZ outcome source
  (33, 100, 1001, '0x3300000000000000000000000000000000000033'); -- CZ vault

UPDATE cz_contracts SET
  market_addr = '0x3100000000000000000000000000000000000031',
  source_addr = '0x3200000000000000000000000000000000000032',
  vault_addr  = '0x3300000000000000000000000000000000000033';

-- ---------------------------------------------------------------------------
-- Blocks 150..166 and transactions 1101..1117 (one tx per block)
-- ---------------------------------------------------------------------------
INSERT INTO block(block_num, ts, block_hash, parent_hash)
SELECT
  n,
  TO_TIMESTAMP(1767225600 + (n - 100) * 100),
  '0xb000000000000000000000000000000000000000000000000000000000000' || LPAD(n::TEXT, 3, '0'),
  '0xb000000000000000000000000000000000000000000000000000000000000' || LPAD((n - 1)::TEXT, 3, '0')
FROM generate_series(150, 166) AS n;

INSERT INTO transaction(id, block_num, from_aid, to_aid, tx_hash) VALUES
  (1101, 150, 21, 31, '0xf000000000000000000000000000000000000000000000000000000000001101'), -- round 0 initialized
  (1102, 151, 21, 31, '0xf000000000000000000000000000000000000000000000000000000000001102'), -- alice funds round 0
  (1103, 152, 21, 31, '0xf000000000000000000000000000000000000000000000000000000000001103'), -- alice buys q0 YES (+mint)
  (1104, 153, 22, 31, '0xf000000000000000000000000000000000000000000000000000000000001104'), -- bob buys q1 NO
  (1105, 154, 23, 31, '0xf000000000000000000000000000000000000000000000000000000000001105'), -- carol splits q2
  (1106, 155, 23, 31, '0xf000000000000000000000000000000000000000000000000000000000001106'), -- carol sells q2 YES
  (1107, 156, 23, 31, '0xf000000000000000000000000000000000000000000000000000000000001107'), -- carol merges
  (1108, 157, 32, 31, '0xf000000000000000000000000000000000000000000000000000000000001108'), -- answers observed
  (1109, 158, 21, 31, '0xf000000000000000000000000000000000000000000000000000000000001109'), -- round 0 settled (+escrow)
  (1110, 159, 21, 31, '0xf000000000000000000000000000000000000000000000000000000000001110'), -- alice claims (+vault payout)
  (1111, 160, 21, 31, '0xf000000000000000000000000000000000000000000000000000000000001111'), -- alice transfers to dave
  (1112, 161, 23, 31, '0xf000000000000000000000000000000000000000000000000000000000001112'), -- carol batch-transfers to dave
  (1113, 162, 24, 31, '0xf000000000000000000000000000000000000000000000000000000000001113'), -- dave approves bob
  (1114, 163, 21, 31, '0xf000000000000000000000000000000000000000000000000000000000001114'), -- round 1 initialized
  (1115, 164, 22, 31, '0xf000000000000000000000000000000000000000000000000000000000001115'), -- bob funds round 1
  (1116, 165, 24, 31, '0xf000000000000000000000000000000000000000000000000000000000001116'), -- dave buys q3 YES
  (1117, 166, 24, 31, '0xf000000000000000000000000000000000000000000000000000000000001117'); -- dave revokes bob

INSERT INTO evt_log(id, block_num, tx_id, contract_aid, topic0_sig, log_index, log_rlp) VALUES
  (5301, 150, 1101, 31, 'cz0ri001', 0, '\x00'), -- RoundInitialized 0
  (5302, 151, 1102, 31, 'cz0rf001', 0, '\x00'), -- RoundFunded alice
  (5303, 152, 1103, 31, 'cz0tr001', 0, '\x00'), -- Trade alice buy q0 YES
  (5304, 152, 1103, 31, 'cz0ts001', 1, '\x00'), -- TransferSingle mint to alice
  (5305, 153, 1104, 31, 'cz0tr001', 0, '\x00'), -- Trade bob buy q1 NO
  (5306, 154, 1105, 31, 'cz0ps001', 0, '\x00'), -- PositionSplit carol q2
  (5307, 155, 1106, 31, 'cz0tr001', 0, '\x00'), -- Trade carol sell q2 YES
  (5308, 156, 1107, 31, 'cz0pm001', 0, '\x00'), -- PositionsMerged carol
  (5309, 157, 1108, 31, 'cz0ao001', 0, '\x00'), -- AnswersObserved
  (5310, 158, 1109, 31, 'cz0rs001', 0, '\x00'), -- RoundSettled 0
  (5311, 158, 1109, 33, 'cz0ce001', 1, '\x00'), -- CycleEscrowed (vault)
  (5312, 159, 1110, 31, 'cz0cl001', 0, '\x00'), -- Claimed alice
  (5313, 159, 1110, 33, 'cz0pd001', 1, '\x00'), -- Paid alice (vault)
  (5314, 160, 1111, 31, 'cz0ts001', 0, '\x00'), -- TransferSingle alice->dave
  (5315, 161, 1112, 31, 'cz0tb001', 0, '\x00'), -- TransferBatch carol->dave (2 rows)
  (5316, 162, 1113, 31, 'cz0af001', 0, '\x00'), -- ApprovalForAll dave->bob true
  (5317, 163, 1114, 31, 'cz0ri001', 0, '\x00'), -- RoundInitialized 1
  (5318, 164, 1115, 31, 'cz0rf001', 0, '\x00'), -- RoundFunded bob (uses carry)
  (5319, 165, 1116, 31, 'cz0tr001', 0, '\x00'), -- Trade dave buy q3 YES
  (5320, 166, 1117, 31, 'cz0af001', 0, '\x00'); -- ApprovalForAll dave->bob false

-- ---------------------------------------------------------------------------
-- Round 0: full lifecycle
-- ---------------------------------------------------------------------------
INSERT INTO cz_round_initialized(evtlog_id, block_num, tx_id, time_stamp, contract_aid, round_id) VALUES
  (5301, 150, 1101, TO_TIMESTAMP(1767230600), 31, 0);

-- alice funds 100 CST of her own, split 12.5 per question pool.
INSERT INTO cz_round_funded(evtlog_id, block_num, tx_id, time_stamp, contract_aid, round_id,
    caller_aid, donated, transfers_used, carry_used, allocations) VALUES
  (5302, 151, 1102, TO_TIMESTAMP(1767230700), 31, 0, 21,
   100000000000000000000, 0, 0,
   ARRAY[12500000000000000000, 12500000000000000000, 12500000000000000000, 12500000000000000000,
         12500000000000000000, 12500000000000000000, 12500000000000000000, 12500000000000000000]::NUMERIC[]);

-- Trades: alice buys q0 YES at 0.75, bob buys q1 NO at 0.625, carol sells
-- q2 YES at 0.7 (after splitting 20 CST into q2 pairs).
INSERT INTO cz_trade(evtlog_id, block_num, tx_id, time_stamp, contract_aid, round_id,
    user_aid, question, yes, buy, cash, units) VALUES
  (5303, 152, 1103, TO_TIMESTAMP(1767230800), 31, 0, 21, 0, TRUE,  TRUE,
   12000000000000000000, 16000000000000000000),
  (5305, 153, 1104, TO_TIMESTAMP(1767230900), 31, 0, 22, 1, FALSE, TRUE,
   5000000000000000000, 8000000000000000000),
  (5307, 155, 1106, TO_TIMESTAMP(1767231100), 31, 0, 23, 2, TRUE,  FALSE,
   7000000000000000000, 10000000000000000000);

INSERT INTO cz_position_split(evtlog_id, block_num, tx_id, time_stamp, contract_aid, round_id,
    user_aid, question, amount) VALUES
  (5306, 154, 1105, TO_TIMESTAMP(1767231000), 31, 0, 23, 2, 20000000000000000000);

INSERT INTO cz_positions_merged(evtlog_id, block_num, tx_id, time_stamp, contract_aid, round_id,
    user_aid, amounts, collateral_out) VALUES
  (5308, 156, 1107, TO_TIMESTAMP(1767231200), 31, 0, 23,
   ARRAY[0, 0, 5000000000000000000, 0, 0, 0, 0, 0]::NUMERIC[], 5000000000000000000);

-- q0's outcome (YES) became observable mid-round.
INSERT INTO cz_answers_observed(evtlog_id, block_num, tx_id, time_stamp, contract_aid, round_id,
    known_mask, values_mask) VALUES
  (5309, 157, 1108, TO_TIMESTAMP(1767231300), 31, 0, 1, 1);

-- Settlement: q0 YES, q2 YES (mask 5); 35 CST escrowed for winners, 90 carried.
INSERT INTO cz_round_settled(evtlog_id, block_num, tx_id, time_stamp, contract_aid, round_id,
    values_mask, answers, escrowed, carried) VALUES
  (5310, 158, 1109, TO_TIMESTAMP(1767231400), 31, 0, 5,
   ARRAY[TRUE, FALSE, TRUE, FALSE, FALSE, FALSE, FALSE, FALSE]::BOOLEAN[],
   ARRAY[20000000000000000000, 0, 15000000000000000000, 0, 0, 0, 0, 0]::NUMERIC[],
   90000000000000000000);

INSERT INTO cz_cycle_escrowed(evtlog_id, block_num, tx_id, time_stamp, contract_aid, round_id, amounts) VALUES
  (5311, 158, 1109, TO_TIMESTAMP(1767231400), 33, 0,
   ARRAY[20000000000000000000, 0, 15000000000000000000, 0, 0, 0, 0, 0]::NUMERIC[]);

-- alice claims her winning q0 YES units; the vault pays out.
INSERT INTO cz_claim(evtlog_id, block_num, tx_id, time_stamp, contract_aid, round_id,
    user_aid, amounts, total) VALUES
  (5312, 159, 1110, TO_TIMESTAMP(1767231500), 31, 0, 21,
   ARRAY[17000000000000000000, 0, 0, 0, 0, 0, 0, 0]::NUMERIC[], 17000000000000000000);

INSERT INTO cz_paid(evtlog_id, block_num, tx_id, time_stamp, contract_aid, round_id,
    question, recipient_aid, amount) VALUES
  (5313, 159, 1110, TO_TIMESTAMP(1767231500), 33, 0, 0, 21, 17000000000000000000);

-- ---------------------------------------------------------------------------
-- ERC-1155 position transfers and approvals
-- position_id = (roundId << 4) | (question << 1) | yes
-- ---------------------------------------------------------------------------
INSERT INTO cz_position_transfer(evtlog_id, batch_index, block_num, tx_id, time_stamp, contract_aid,
    operator_aid, from_aid, to_aid, position_id, round_id, question, yes, amount) VALUES
  -- Mint of alice's q0 YES units bought in tx 1103 (from the zero address).
  (5304, 0, 152, 1103, TO_TIMESTAMP(1767230800), 31, 31, 1, 21, 1, 0, 0, TRUE, 16000000000000000000),
  -- alice hands 3 of those units to dave.
  (5314, 0, 160, 1111, TO_TIMESTAMP(1767231600), 31, 21, 21, 24, 1, 0, 0, TRUE, 3000000000000000000),
  -- carol batch-transfers leftover q2 positions to dave (one evt, two rows).
  (5315, 0, 161, 1112, TO_TIMESTAMP(1767231700), 31, 23, 23, 24, 5, 0, 2, TRUE, 2000000000000000000),
  (5315, 1, 161, 1112, TO_TIMESTAMP(1767231700), 31, 23, 23, 24, 4, 0, 2, FALSE, 3000000000000000000);

INSERT INTO cz_approval_for_all(evtlog_id, block_num, tx_id, time_stamp, contract_aid,
    owner_aid, operator_aid, approved) VALUES
  (5316, 162, 1113, TO_TIMESTAMP(1767231800), 31, 24, 22, TRUE),
  (5320, 166, 1117, TO_TIMESTAMP(1767232200), 31, 24, 22, FALSE);

-- ---------------------------------------------------------------------------
-- Round 1: open (initialized, funded with round 0's carry, one trade)
-- ---------------------------------------------------------------------------
INSERT INTO cz_round_initialized(evtlog_id, block_num, tx_id, time_stamp, contract_aid, round_id) VALUES
  (5317, 163, 1114, TO_TIMESTAMP(1767231900), 31, 1);

INSERT INTO cz_round_funded(evtlog_id, block_num, tx_id, time_stamp, contract_aid, round_id,
    caller_aid, donated, transfers_used, carry_used, allocations) VALUES
  (5318, 164, 1115, TO_TIMESTAMP(1767232000), 31, 1, 22,
   10000000000000000000, 0, 90000000000000000000,
   ARRAY[12500000000000000000, 12500000000000000000, 12500000000000000000, 12500000000000000000,
         12500000000000000000, 12500000000000000000, 12500000000000000000, 12500000000000000000]::NUMERIC[]);

INSERT INTO cz_trade(evtlog_id, block_num, tx_id, time_stamp, contract_aid, round_id,
    user_aid, question, yes, buy, cash, units) VALUES
  (5319, 165, 1116, TO_TIMESTAMP(1767232100), 31, 1, 24, 3, TRUE, TRUE,
   4000000000000000000, 5000000000000000000);
