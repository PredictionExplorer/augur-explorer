-- RandomWalk domain fixture.
--
-- Insert order matters: rw_transfer creates the rw_token row (mint), then
-- rw_mint_evt fills seed/price, then names, offers, sales, cancellations and
-- withdrawals. Triggers maintain rw_token / rw_stats / rw_mkt_stats /
-- rw_user_stats / rw_user_rwtok.

-- Baseline stats rows (a long-running deployment always has these; the offer
-- triggers UPDATE them without an insert-if-missing fallback).
INSERT INTO rw_stats(rwalk_aid) VALUES (8);
INSERT INTO rw_mkt_stats(contract_aid) VALUES (12);

-- === Mints: tokens 10..13 ===================================================
INSERT INTO rw_transfer(evtlog_id, block_num, tx_id, time_stamp, contract_aid, token_id, from_aid, to_aid, otype) VALUES
  (5079, 130, 1030, TO_TIMESTAMP(1767228600), 8, 10, 1, 23, 1);
INSERT INTO rw_mint_evt(evtlog_id, block_num, tx_id, time_stamp, contract_aid, token_id, owner_aid, seed, seed_num, price) VALUES
  (5080, 130, 1030, TO_TIMESTAMP(1767228600), 8, 10, 23, 'aa00000000000000000000000000000000000000000000000000000000000010', 16, 50000000000000000);

INSERT INTO rw_transfer(evtlog_id, block_num, tx_id, time_stamp, contract_aid, token_id, from_aid, to_aid, otype) VALUES
  (5081, 131, 1031, TO_TIMESTAMP(1767228700), 8, 11, 1, 24, 1);
INSERT INTO rw_mint_evt(evtlog_id, block_num, tx_id, time_stamp, contract_aid, token_id, owner_aid, seed, seed_num, price) VALUES
  (5082, 131, 1031, TO_TIMESTAMP(1767228700), 8, 11, 24, 'aa00000000000000000000000000000000000000000000000000000000000011', 17, 55000000000000000);

INSERT INTO rw_transfer(evtlog_id, block_num, tx_id, time_stamp, contract_aid, token_id, from_aid, to_aid, otype) VALUES
  (5083, 132, 1032, TO_TIMESTAMP(1767228800), 8, 12, 1, 21, 1);
INSERT INTO rw_mint_evt(evtlog_id, block_num, tx_id, time_stamp, contract_aid, token_id, owner_aid, seed, seed_num, price) VALUES
  (5084, 132, 1032, TO_TIMESTAMP(1767228800), 8, 12, 21, 'aa00000000000000000000000000000000000000000000000000000000000012', 18, 60000000000000000);

INSERT INTO rw_transfer(evtlog_id, block_num, tx_id, time_stamp, contract_aid, token_id, from_aid, to_aid, otype) VALUES
  (5085, 133, 1033, TO_TIMESTAMP(1767228900), 8, 13, 1, 22, 1);
INSERT INTO rw_mint_evt(evtlog_id, block_num, tx_id, time_stamp, contract_aid, token_id, owner_aid, seed, seed_num, price) VALUES
  (5086, 133, 1033, TO_TIMESTAMP(1767228900), 8, 13, 22, 'aa00000000000000000000000000000000000000000000000000000000000013', 19, 65000000000000000);

-- === Naming =================================================================
INSERT INTO rw_token_name(evtlog_id, block_num, tx_id, time_stamp, contract_aid, token_id, new_name) VALUES
  (5087, 134, 1034, TO_TIMESTAMP(1767229000), 8, 10, 'Wanderer');

-- === Marketplace ============================================================
-- Offer 1: carol sells #10 for 1 ETH; dave buys it (profit over 0.05 mint price).
INSERT INTO rw_new_offer(evtlog_id, block_num, tx_id, time_stamp, contract_aid, rwalk_aid, offer_id, seller_aid, buyer_aid, otype, token_id, active, price) VALUES
  (5088, 135, 1035, TO_TIMESTAMP(1767229100), 12, 8, 1, 23, 1, 1, 10, TRUE, 1000000000000000000);

INSERT INTO rw_item_bought(evtlog_id, block_num, tx_id, time_stamp, contract_aid, offer_id, seller_aid, buyer_aid) VALUES
  (5089, 136, 1036, TO_TIMESTAMP(1767229200), 12, 1, 23, 24);

INSERT INTO rw_transfer(evtlog_id, block_num, tx_id, time_stamp, contract_aid, token_id, from_aid, to_aid, otype) VALUES
  (5090, 136, 1036, TO_TIMESTAMP(1767229200), 8, 10, 23, 24, 0);

-- Offer 2: dave sells #11 (active).
INSERT INTO rw_new_offer(evtlog_id, block_num, tx_id, time_stamp, contract_aid, rwalk_aid, offer_id, seller_aid, buyer_aid, otype, token_id, active, price) VALUES
  (5091, 137, 1037, TO_TIMESTAMP(1767229300), 12, 8, 2, 24, 1, 1, 11, TRUE, 2000000000000000000);

-- Offer 3: bob sells #13 (active).
INSERT INTO rw_new_offer(evtlog_id, block_num, tx_id, time_stamp, contract_aid, rwalk_aid, offer_id, seller_aid, buyer_aid, otype, token_id, active, price) VALUES
  (5092, 138, 1038, TO_TIMESTAMP(1767229400), 12, 8, 3, 22, 1, 1, 13, TRUE, 2500000000000000000);

-- Offer 4: dave relists #10, then cancels.
INSERT INTO rw_new_offer(evtlog_id, block_num, tx_id, time_stamp, contract_aid, rwalk_aid, offer_id, seller_aid, buyer_aid, otype, token_id, active, price) VALUES
  (5093, 139, 1039, TO_TIMESTAMP(1767229500), 12, 8, 4, 24, 1, 1, 10, TRUE, 3000000000000000000);

INSERT INTO rw_offer_canceled(evtlog_id, block_num, tx_id, time_stamp, contract_aid, offer_id) VALUES
  (5094, 140, 1040, TO_TIMESTAMP(1767229600), 12, 4);

-- === Withdrawal =============================================================
INSERT INTO rw_withdrawal(evtlog_id, block_num, tx_id, time_stamp, contract_aid, aid, token_id, amount) VALUES
  (5095, 141, 1041, TO_TIMESTAMP(1767229700), 8, 23, 10, 30000000000000000);

-- === Cosmic Signature on the shared marketplace ============================
-- Offer 101: alice lists Cosmic Signature #1 for 1.5 ETH; bob buys it.
INSERT INTO rw_new_offer(evtlog_id, block_num, tx_id, time_stamp, contract_aid, rwalk_aid, offer_id, seller_aid, buyer_aid, otype, token_id, active, price) VALUES
  (5201, 130, 1044, TO_TIMESTAMP(1767228610), 12, 3, 101, 21, 1, 1, 1, TRUE, 1500000000000000000);

INSERT INTO rw_item_bought(evtlog_id, block_num, tx_id, time_stamp, contract_aid, offer_id, seller_aid, buyer_aid) VALUES
  (5202, 131, 1045, TO_TIMESTAMP(1767228710), 12, 101, 21, 22);

-- Offer 102: dave lists #2 for 2 ETH (active).
INSERT INTO rw_new_offer(evtlog_id, block_num, tx_id, time_stamp, contract_aid, rwalk_aid, offer_id, seller_aid, buyer_aid, otype, token_id, active, price) VALUES
  (5203, 132, 1046, TO_TIMESTAMP(1767228810), 12, 3, 102, 24, 1, 1, 2, TRUE, 2000000000000000000);

-- Offer 103: bob places a 0.75 ETH buy bid for #3 (active).
INSERT INTO rw_new_offer(evtlog_id, block_num, tx_id, time_stamp, contract_aid, rwalk_aid, offer_id, seller_aid, buyer_aid, otype, token_id, active, price) VALUES
  (5204, 133, 1047, TO_TIMESTAMP(1767228910), 12, 3, 103, 1, 22, 0, 3, TRUE, 750000000000000000);

-- Offer 104: carol lists #3 for 3 ETH, then cancels it.
INSERT INTO rw_new_offer(evtlog_id, block_num, tx_id, time_stamp, contract_aid, rwalk_aid, offer_id, seller_aid, buyer_aid, otype, token_id, active, price) VALUES
  (5205, 134, 1048, TO_TIMESTAMP(1767229010), 12, 3, 104, 23, 1, 1, 3, TRUE, 3000000000000000000);

INSERT INTO rw_offer_canceled(evtlog_id, block_num, tx_id, time_stamp, contract_aid, offer_id) VALUES
  (5206, 135, 1049, TO_TIMESTAMP(1767229110), 12, 104);

-- Offer 105: dave lists #4 at the same price as #2, exercising the exact
-- price/event-log tie break.
INSERT INTO rw_new_offer(evtlog_id, block_num, tx_id, time_stamp, contract_aid, rwalk_aid, offer_id, seller_aid, buyer_aid, otype, token_id, active, price) VALUES
  (5207, 136, 1050, TO_TIMESTAMP(1767229210), 12, 3, 105, 24, 1, 1, 4, TRUE, 2000000000000000000);

-- === Beauty ranking =========================================================
INSERT INTO rw_token_ranking(token_id, rating, updated_at) VALUES
  (10, 1210.5, TO_TIMESTAMP(1767232000)),
  (11, 1189.5, TO_TIMESTAMP(1767232000)),
  (12, 1195,   TO_TIMESTAMP(1767232100)),
  (13, 1205,   TO_TIMESTAMP(1767232100));

INSERT INTO rw_ranking_match(nft1, nft2, nft1_won, voter_aid, created_at) VALUES
  (10, 11, TRUE,  21,   TO_TIMESTAMP(1767232000)),
  (12, 13, FALSE, NULL, TO_TIMESTAMP(1767232100));
