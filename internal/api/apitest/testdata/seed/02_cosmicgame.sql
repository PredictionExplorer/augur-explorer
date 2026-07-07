-- CosmicGame domain fixture.
--
-- Rows are inserted in chain order so the load-bearing plpgsql triggers from
-- migration 00002 (cg_bidder / cg_winner / cg_glob_stats / cg_round_stats /
-- staking accumulators / cg_prize) compute every aggregate exactly as they
-- would in production. Do not reorder without checking trigger dependencies.
--
-- Story: rounds 0..2 complete (winners alice, dave, emma); round 3 is open.
-- Round 0 exercises every prize type, donations of all kinds and staking.

-- === Before round 0: admin config ==========================================
-- Activation time for round 0 (no prize claims yet -> applies to round 0).
INSERT INTO cg_adm_acttime(evtlog_id, block_num, tx_id, time_stamp, contract_aid, new_atime)
VALUES (5001, 100, 1001, TO_TIMESTAMP(1767225600), 2, 1767225650);

INSERT INTO cg_adm_charity_pcent(evtlog_id, block_num, tx_id, time_stamp, contract_aid, percentage)
VALUES (5002, 100, 1001, TO_TIMESTAMP(1767225600), 2, 10);

-- === Round 0: bidding =======================================================
INSERT INTO cg_first_bid(evtlog_id, block_num, tx_id, time_stamp, contract_aid, round_num, start_ts)
VALUES (5003, 101, 1002, TO_TIMESTAMP(1767225700), 2, 0, 1767225700);

INSERT INTO cg_bid(id, evtlog_id, block_num, tx_id, time_stamp, contract_aid, bidder_aid,
                   rwalk_nft_id, round_num, bid_type, bid_position, prize_time,
                   eth_price, cst_price, cst_reward, bid_cst_reward_amount, cst_dutch_auction_duration, msg) VALUES
  (2001, 5004, 101, 1002, TO_TIMESTAMP(1767225700), 2, 21, -1, 0, 0, 1, TO_TIMESTAMP(1767229300),
   100000000000000000, -1, 100000000000000000000, -1, -1, 'hello world'),
  (2002, 5006, 102, 1003, TO_TIMESTAMP(1767225800), 2, 22, 13, 0, 1, 2, TO_TIMESTAMP(1767229400),
   50000000000000000, -1, 100000000000000000000, -1, -1, ''),
  (2003, 5008, 103, 1004, TO_TIMESTAMP(1767225900), 2, 23, -1, 0, 2, 3, TO_TIMESTAMP(1767229500),
   -1, 200000000000000000000, 100000000000000000000, 98000000000000000000, 1800, 'cst bid'),
  (2004, 5010, 104, 1005, TO_TIMESTAMP(1767226000), 2, 21, -1, 0, 0, 4, TO_TIMESTAMP(1767229600),
   120000000000000000, -1, 100000000000000000000, -1, -1, '');

-- CST bidding rewards (ERC-20 mints; zero address -> bidder).
INSERT INTO cg_erc20_transfer(evtlog_id, block_num, tx_id, time_stamp, contract_aid, value, from_aid, to_aid, otype) VALUES
  (5005, 101, 1002, TO_TIMESTAMP(1767225700), 4, 100000000000000000000, 1, 21, 1),
  (5007, 102, 1003, TO_TIMESTAMP(1767225800), 4, 100000000000000000000, 1, 22, 1),
  (5009, 103, 1004, TO_TIMESTAMP(1767225900), 4, 100000000000000000000, 1, 23, 1),
  (5011, 104, 1005, TO_TIMESTAMP(1767226000), 4, 100000000000000000000, 1, 21, 1);

-- === Round 0: donations =====================================================
INSERT INTO cg_eth_donated(evtlog_id, block_num, tx_id, time_stamp, contract_aid, donor_aid, round_num, amount)
VALUES (5012, 105, 1006, TO_TIMESTAMP(1767226100), 2, 24, 0, 200000000000000000);

-- Voluntary donation into the charity wallet (donor is not the CosmicGame contract).
INSERT INTO cg_donation_received(evtlog_id, block_num, tx_id, time_stamp, contract_aid, donor_aid, amount, round_num)
VALUES (5013, 105, 1006, TO_TIMESTAMP(1767226100), 6, 24, 10000000000000000, -1);

INSERT INTO cg_eth_donated_wi(evtlog_id, block_num, tx_id, time_stamp, contract_aid, donor_aid, round_num, record_id, amount)
VALUES (5014, 106, 1007, TO_TIMESTAMP(1767226200), 2, 25, 0, 0, 300000000000000000);

INSERT INTO cg_donation_json(record_id, data)
VALUES (0, '{"title":"apitest donation","url":"https://example.org/apitest"}');

INSERT INTO cg_erc20_donation(evtlog_id, block_num, tx_id, time_stamp, contract_aid, round_num, donor_aid, token_aid, amount, bid_id)
VALUES (5015, 107, 1008, TO_TIMESTAMP(1767226300), 7, 0, 21, 26, 500000000000000000000, 2001);

INSERT INTO cg_nft_donation(evtlog_id, block_num, tx_id, time_stamp, contract_aid, round_num, donor_aid, token_aid, token_id, idx, bid_id, token_uri)
VALUES (5016, 108, 1009, TO_TIMESTAMP(1767226400), 7, 0, 22, 27, 777, 0, 2002, 'https://nft.example/777');

-- === Round 0: marketing reward =============================================
INSERT INTO cg_mkt_reward(evtlog_id, block_num, tx_id, time_stamp, contract_aid, amount, marketer_aid)
VALUES (5017, 109, 1010, TO_TIMESTAMP(1767226500), 11, 50000000000000000000, 25);

-- === Round 0: prize claim (alice) ===========================================
INSERT INTO cg_prize_claim(evtlog_id, block_num, tx_id, time_stamp, contract_aid, round_num, winner_aid, token_id, timeout, amount, cst_amount)
VALUES (5018, 110, 1011, TO_TIMESTAMP(1767226600), 2, 0, 21, 1, 1768089600, 500000000000000000, 100000000000000000000);

INSERT INTO cg_mint_event(evtlog_id, block_num, tx_id, time_stamp, contract_aid, owner_aid, token_id, cur_owner_aid, round_num, seed) VALUES
  (5019, 110, 1011, TO_TIMESTAMP(1767226600), 3, 21, 1, 21, 0, 'seed0000000000000000000000000000000000000000000000000000000001');

INSERT INTO cg_erc721_transfer(evtlog_id, block_num, tx_id, time_stamp, contract_aid, token_id, from_aid, to_aid, otype) VALUES
  (5020, 110, 1011, TO_TIMESTAMP(1767226600), 3, 1, 1, 21, 1);

-- Raffle ETH winners (bob, carol) with their PrizesWallet deposits.
INSERT INTO cg_raffle_eth_prize(evtlog_id, block_num, tx_id, time_stamp, contract_aid, winner_aid, round_num, winner_idx, amount) VALUES
  (5021, 110, 1011, TO_TIMESTAMP(1767226600), 2, 22, 0, 0, 50000000000000000),
  (5023, 110, 1011, TO_TIMESTAMP(1767226600), 2, 23, 0, 1, 50000000000000000);

INSERT INTO cg_prize_deposit(evtlog_id, block_num, tx_id, time_stamp, contract_aid, winner_aid, round_num, winner_index, amount) VALUES
  (5022, 110, 1011, TO_TIMESTAMP(1767226600), 7, 22, 0, 0, 50000000000000000),
  (5024, 110, 1011, TO_TIMESTAMP(1767226600), 7, 23, 0, 1, 50000000000000000);

-- Raffle NFT winners: dave (bidder raffle) and carol (RandomWalk-staker raffle).
INSERT INTO cg_raffle_nft_prize(evtlog_id, block_num, tx_id, time_stamp, contract_aid, winner_aid, round_num, token_id, winner_idx, cst_amount, is_rwalk, is_staker) VALUES
  (5025, 110, 1011, TO_TIMESTAMP(1767226600), 2, 24, 0, 2, 0, 30000000000000000000, FALSE, FALSE),
  (5028, 110, 1011, TO_TIMESTAMP(1767226600), 2, 23, 0, 3, 0, 30000000000000000000, TRUE, TRUE);

INSERT INTO cg_mint_event(evtlog_id, block_num, tx_id, time_stamp, contract_aid, owner_aid, token_id, cur_owner_aid, round_num, seed) VALUES
  (5026, 110, 1011, TO_TIMESTAMP(1767226600), 3, 24, 2, 24, 0, 'seed0000000000000000000000000000000000000000000000000000000002'),
  (5029, 110, 1011, TO_TIMESTAMP(1767226600), 3, 23, 3, 23, 0, 'seed0000000000000000000000000000000000000000000000000000000003');

INSERT INTO cg_erc721_transfer(evtlog_id, block_num, tx_id, time_stamp, contract_aid, token_id, from_aid, to_aid, otype) VALUES
  (5027, 110, 1011, TO_TIMESTAMP(1767226600), 3, 2, 1, 24, 1),
  (5030, 110, 1011, TO_TIMESTAMP(1767226600), 3, 3, 1, 23, 1);

-- Last CST bidder prize (carol).
INSERT INTO cg_lastcst_prize(evtlog_id, block_num, tx_id, time_stamp, contract_aid, winner_aid, round_num, erc721_token_id, erc20_amount)
VALUES (5031, 110, 1011, TO_TIMESTAMP(1767226600), 2, 23, 0, 4, 40000000000000000000);

INSERT INTO cg_mint_event(evtlog_id, block_num, tx_id, time_stamp, contract_aid, owner_aid, token_id, cur_owner_aid, round_num, seed) VALUES
  (5032, 110, 1011, TO_TIMESTAMP(1767226600), 3, 23, 4, 23, 0, 'seed0000000000000000000000000000000000000000000000000000000004');

INSERT INTO cg_erc721_transfer(evtlog_id, block_num, tx_id, time_stamp, contract_aid, token_id, from_aid, to_aid, otype) VALUES
  (5033, 110, 1011, TO_TIMESTAMP(1767226600), 3, 4, 1, 23, 1);

-- Endurance champion prize (bob).
INSERT INTO cg_endurance_prize(evtlog_id, block_num, tx_id, time_stamp, contract_aid, winner_aid, round_num, erc721_token_id, erc20_amount)
VALUES (5034, 110, 1011, TO_TIMESTAMP(1767226600), 2, 22, 0, 5, 45000000000000000000);

INSERT INTO cg_mint_event(evtlog_id, block_num, tx_id, time_stamp, contract_aid, owner_aid, token_id, cur_owner_aid, round_num, seed) VALUES
  (5035, 110, 1011, TO_TIMESTAMP(1767226600), 3, 22, 5, 22, 0, 'seed0000000000000000000000000000000000000000000000000000000005');

INSERT INTO cg_erc721_transfer(evtlog_id, block_num, tx_id, time_stamp, contract_aid, token_id, from_aid, to_aid, otype) VALUES
  (5036, 110, 1011, TO_TIMESTAMP(1767226600), 3, 5, 1, 22, 1);

-- Chrono warrior prize (alice) + its PrizesWallet deposit. winner_index is the
-- per-round EthReceived registration counter, shared with the raffle deposits
-- (bob 0, carol 1, chrono 2) — the by-user deposit queries join on it.
INSERT INTO cg_chrono_warrior_prize(evtlog_id, block_num, tx_id, time_stamp, contract_aid, winner_aid, round_num, winner_index, eth_amount, cst_amount, nft_id)
VALUES (5037, 110, 1011, TO_TIMESTAMP(1767226600), 2, 21, 0, 2, 80000000000000000, 35000000000000000000, 6);

INSERT INTO cg_mint_event(evtlog_id, block_num, tx_id, time_stamp, contract_aid, owner_aid, token_id, cur_owner_aid, round_num, seed) VALUES
  (5038, 110, 1011, TO_TIMESTAMP(1767226600), 3, 21, 6, 21, 0, 'seed0000000000000000000000000000000000000000000000000000000006');

INSERT INTO cg_erc721_transfer(evtlog_id, block_num, tx_id, time_stamp, contract_aid, token_id, from_aid, to_aid, otype) VALUES
  (5039, 110, 1011, TO_TIMESTAMP(1767226600), 3, 6, 1, 21, 1);

INSERT INTO cg_prize_deposit(evtlog_id, block_num, tx_id, time_stamp, contract_aid, winner_aid, round_num, winner_index, amount) VALUES
  (5040, 110, 1011, TO_TIMESTAMP(1767226600), 7, 21, 0, 2, 80000000000000000);

-- CosmicGame -> charity wallet transfer for round 0.
INSERT INTO cg_donation_received(evtlog_id, block_num, tx_id, time_stamp, contract_aid, donor_aid, amount, round_num)
VALUES (5041, 110, 1011, TO_TIMESTAMP(1767226600), 6, 2, 90000000000000000, 0);

INSERT INTO cg_funds_to_charity(evtlog_id, block_num, tx_id, time_stamp, contract_aid, charity_aid, amount)
VALUES (5042, 110, 1011, TO_TIMESTAMP(1767226600), 2, 6, 90000000000000000);

-- Main-prize CST reward mint for alice.
INSERT INTO cg_erc20_transfer(evtlog_id, block_num, tx_id, time_stamp, contract_aid, value, from_aid, to_aid, otype) VALUES
  (5043, 110, 1011, TO_TIMESTAMP(1767226600), 4, 100000000000000000000, 1, 21, 1);

-- === Round 0 aftermath ======================================================
-- Bob withdraws his raffle ETH (marks his round-0 deposit claimed).
INSERT INTO cg_prize_withdrawal(evtlog_id, block_num, tx_id, time_stamp, contract_aid, round_num, winner_aid, beneficiary_aid, amount)
VALUES (5044, 111, 1012, TO_TIMESTAMP(1767226700), 7, 0, 22, 22, 50000000000000000);

-- Alice claims the donated NFT and the donated ERC-20 tokens.
INSERT INTO cg_donated_nft_claimed(evtlog_id, block_num, tx_id, time_stamp, contract_aid, round_num, idx, token_aid, winner_aid, token_id)
VALUES (5045, 112, 1013, TO_TIMESTAMP(1767226800), 7, 0, 0, 27, 21, 777);

INSERT INTO cg_donated_tok_claimed(evtlog_id, block_num, tx_id, time_stamp, contract_aid, round_num, idx, token_aid, winner_aid, amount)
VALUES (5046, 112, 1013, TO_TIMESTAMP(1767226800), 7, 0, 0, 26, 21, 500000000000000000000);

-- Alice names her main-prize token.
INSERT INTO cg_token_name(evtlog_id, block_num, tx_id, time_stamp, contract_aid, token_id, token_name)
VALUES (5047, 113, 1014, TO_TIMESTAMP(1767226900), 3, 1, 'Genesis');

-- Secondary transfers: NFT #2 dave->bob, CST alice->bob, CST burn by carol.
INSERT INTO cg_erc721_transfer(evtlog_id, block_num, tx_id, time_stamp, contract_aid, token_id, from_aid, to_aid, otype) VALUES
  (5048, 113, 1014, TO_TIMESTAMP(1767226900), 3, 2, 24, 22, 0);

INSERT INTO cg_erc20_transfer(evtlog_id, block_num, tx_id, time_stamp, contract_aid, value, from_aid, to_aid, otype) VALUES
  (5049, 113, 1014, TO_TIMESTAMP(1767226900), 4, 10000000000000000000, 21, 22, 0),
  (5050, 113, 1014, TO_TIMESTAMP(1767226900), 4, 5000000000000000000, 23, 1, 2);

-- === Staking (after round 0) ================================================
-- CST staking: alice stakes token #1, bob stakes token #5.
INSERT INTO cg_nft_staked_cst(evtlog_id, block_num, tx_id, time_stamp, contract_aid, action_id, token_id, num_staked_nfts, reward_per_staker, staker_aid) VALUES
  (5051, 114, 1015, TO_TIMESTAMP(1767227000), 9, 1, 1, 1, 0, 21),
  (5052, 114, 1015, TO_TIMESTAMP(1767227000), 9, 2, 5, 2, 0, 22);

-- RandomWalk staking: carol stakes #10, dave stakes #11, bob stakes #13
-- (bob thereby stakes in both wallets, feeding unique/stakers/both).
INSERT INTO cg_nft_staked_rwalk(evtlog_id, block_num, tx_id, time_stamp, contract_aid, action_id, token_id, num_staked_nfts, staker_aid) VALUES
  (5053, 114, 1015, TO_TIMESTAMP(1767227000), 10, 101, 10, 1, 23),
  (5054, 114, 1015, TO_TIMESTAMP(1767227000), 10, 102, 11, 2, 24),
  (5098, 114, 1015, TO_TIMESTAMP(1767227000), 10, 103, 13, 3, 22);

-- Staking reward deposit: 2 ETH over 2 staked tokens -> 1 ETH per token.
INSERT INTO cg_staking_eth_deposit(evtlog_id, block_num, tx_id, time_stamp, contract_aid, round_num, deposit_time, deposit_id, num_staked_nfts, deposit_amount, amount_per_token, modulo)
VALUES (5055, 115, 1016, TO_TIMESTAMP(1767227100), 9, 0, TO_TIMESTAMP(1767227100), 501, 2, 2000000000000000000, 1000000000000000000, 0);

-- Alice unstakes token #1 collecting her 1 ETH reward; carol unstakes #10.
INSERT INTO cg_nft_unstaked_cst(evtlog_id, block_num, tx_id, time_stamp, contract_aid, action_id, action_counter, token_id, num_staked_nfts, staker_aid, reward, reward_per_tok)
VALUES (5056, 116, 1017, TO_TIMESTAMP(1767227200), 9, 1, 3, 1, 1, 21, 1000000000000000000, 1000000000000000000);

INSERT INTO cg_nft_unstaked_rwalk(evtlog_id, block_num, tx_id, time_stamp, contract_aid, action_id, token_id, num_staked_nfts, staker_aid)
VALUES (5057, 116, 1017, TO_TIMESTAMP(1767227200), 10, 101, 10, 1, 23);

-- === Round 1 ================================================================
INSERT INTO cg_adm_acttime(evtlog_id, block_num, tx_id, time_stamp, contract_aid, new_atime)
VALUES (5058, 117, 1018, TO_TIMESTAMP(1767227300), 2, 1767227400);

INSERT INTO cg_first_bid(evtlog_id, block_num, tx_id, time_stamp, contract_aid, round_num, start_ts)
VALUES (5059, 118, 1019, TO_TIMESTAMP(1767227400), 2, 1, 1767227400);

INSERT INTO cg_bid(id, evtlog_id, block_num, tx_id, time_stamp, contract_aid, bidder_aid,
                   rwalk_nft_id, round_num, bid_type, bid_position, prize_time,
                   eth_price, cst_price, cst_reward, bid_cst_reward_amount, cst_dutch_auction_duration, msg) VALUES
  (2005, 5060, 118, 1019, TO_TIMESTAMP(1767227400), 2, 22, -1, 1, 0, 1, TO_TIMESTAMP(1767231000),
   100000000000000000, -1, 100000000000000000000, -1, -1, ''),
  (2006, 5061, 119, 1020, TO_TIMESTAMP(1767227500), 2, 24, -1, 1, 0, 2, TO_TIMESTAMP(1767231100),
   110000000000000000, -1, 100000000000000000000, -1, -1, 'go dave');

INSERT INTO cg_prize_claim(evtlog_id, block_num, tx_id, time_stamp, contract_aid, round_num, winner_aid, token_id, timeout, amount, cst_amount)
VALUES (5062, 120, 1021, TO_TIMESTAMP(1767227600), 2, 1, 24, 7, 1768090600, 600000000000000000, 110000000000000000000);

INSERT INTO cg_mint_event(evtlog_id, block_num, tx_id, time_stamp, contract_aid, owner_aid, token_id, cur_owner_aid, round_num, seed) VALUES
  (5063, 120, 1021, TO_TIMESTAMP(1767227600), 3, 24, 7, 24, 1, 'seed0000000000000000000000000000000000000000000000000000000007');

INSERT INTO cg_erc721_transfer(evtlog_id, block_num, tx_id, time_stamp, contract_aid, token_id, from_aid, to_aid, otype) VALUES
  (5064, 120, 1021, TO_TIMESTAMP(1767227600), 3, 7, 1, 24, 1);

INSERT INTO cg_raffle_eth_prize(evtlog_id, block_num, tx_id, time_stamp, contract_aid, winner_aid, round_num, winner_idx, amount) VALUES
  (5065, 120, 1021, TO_TIMESTAMP(1767227600), 2, 21, 1, 0, 60000000000000000);

INSERT INTO cg_prize_deposit(evtlog_id, block_num, tx_id, time_stamp, contract_aid, winner_aid, round_num, winner_index, amount) VALUES
  (5066, 120, 1021, TO_TIMESTAMP(1767227600), 7, 21, 1, 0, 60000000000000000);

-- === Round 2 ================================================================
INSERT INTO cg_adm_acttime(evtlog_id, block_num, tx_id, time_stamp, contract_aid, new_atime)
VALUES (5067, 121, 1022, TO_TIMESTAMP(1767227700), 2, 1767227800);

INSERT INTO cg_first_bid(evtlog_id, block_num, tx_id, time_stamp, contract_aid, round_num, start_ts)
VALUES (5068, 122, 1023, TO_TIMESTAMP(1767227800), 2, 2, 1767227800);

INSERT INTO cg_bid(id, evtlog_id, block_num, tx_id, time_stamp, contract_aid, bidder_aid,
                   rwalk_nft_id, round_num, bid_type, bid_position, prize_time,
                   eth_price, cst_price, cst_reward, bid_cst_reward_amount, cst_dutch_auction_duration, msg) VALUES
  (2007, 5069, 122, 1023, TO_TIMESTAMP(1767227800), 2, 23, -1, 2, 2, 1, TO_TIMESTAMP(1767231400),
   -1, 210000000000000000000, 100000000000000000000, 98000000000000000000, 1750, ''),
  (2008, 5070, 123, 1024, TO_TIMESTAMP(1767227900), 2, 25, -1, 2, 0, 2, TO_TIMESTAMP(1767231500),
   90000000000000000, -1, 100000000000000000000, -1, -1, 'emma was here'),
  (2009, 5071, 124, 1025, TO_TIMESTAMP(1767228000), 2, 21, -1, 2, 0, 3, TO_TIMESTAMP(1767231600),
   100000000000000000, -1, 100000000000000000000, -1, -1, '');

-- Emma bids-and-donates an NFT in round 2; nobody claims it (pins the
-- unclaimed-donated-NFT views). idx is the PrizesWallet's global donation
-- counter (bob's round-0 donation took 0), not a per-round index.
INSERT INTO cg_nft_donation(evtlog_id, block_num, tx_id, time_stamp, contract_aid, round_num, donor_aid, token_aid, token_id, idx, bid_id, token_uri)
VALUES (5102, 123, 1024, TO_TIMESTAMP(1767227900), 7, 2, 25, 27, 888, 1, 2008, 'https://nft.example/888');

INSERT INTO cg_prize_claim(evtlog_id, block_num, tx_id, time_stamp, contract_aid, round_num, winner_aid, token_id, timeout, amount, cst_amount)
VALUES (5072, 125, 1026, TO_TIMESTAMP(1767228100), 2, 2, 25, 8, 1768091100, 700000000000000000, 120000000000000000000);

INSERT INTO cg_mint_event(evtlog_id, block_num, tx_id, time_stamp, contract_aid, owner_aid, token_id, cur_owner_aid, round_num, seed) VALUES
  (5073, 125, 1026, TO_TIMESTAMP(1767228100), 3, 25, 8, 25, 2, 'seed0000000000000000000000000000000000000000000000000000000008');

INSERT INTO cg_erc721_transfer(evtlog_id, block_num, tx_id, time_stamp, contract_aid, token_id, from_aid, to_aid, otype) VALUES
  (5074, 125, 1026, TO_TIMESTAMP(1767228100), 3, 8, 1, 25, 1);

-- Round-2 raffle NFT for bob as a CosmicSignature staker (is_staker without
-- is_rwalk feeds the staking/cst/mints views).
INSERT INTO cg_raffle_nft_prize(evtlog_id, block_num, tx_id, time_stamp, contract_aid, winner_aid, round_num, token_id, winner_idx, cst_amount, is_rwalk, is_staker) VALUES
  (5099, 125, 1026, TO_TIMESTAMP(1767228100), 2, 22, 2, 9, 0, 32000000000000000000, FALSE, TRUE);

INSERT INTO cg_mint_event(evtlog_id, block_num, tx_id, time_stamp, contract_aid, owner_aid, token_id, cur_owner_aid, round_num, seed) VALUES
  (5100, 125, 1026, TO_TIMESTAMP(1767228100), 3, 22, 9, 22, 2, 'seed0000000000000000000000000000000000000000000000000000000009');

INSERT INTO cg_erc721_transfer(evtlog_id, block_num, tx_id, time_stamp, contract_aid, token_id, from_aid, to_aid, otype) VALUES
  (5101, 125, 1026, TO_TIMESTAMP(1767228100), 3, 9, 1, 22, 1);

-- === Round 3 (open) =========================================================
INSERT INTO cg_adm_acttime(evtlog_id, block_num, tx_id, time_stamp, contract_aid, new_atime)
VALUES (5075, 126, 1027, TO_TIMESTAMP(1767228200), 2, 1767228300);

INSERT INTO cg_first_bid(evtlog_id, block_num, tx_id, time_stamp, contract_aid, round_num, start_ts)
VALUES (5076, 127, 1028, TO_TIMESTAMP(1767228300), 2, 3, 1767228300);

INSERT INTO cg_bid(id, evtlog_id, block_num, tx_id, time_stamp, contract_aid, bidder_aid,
                   rwalk_nft_id, round_num, bid_type, bid_position, prize_time,
                   eth_price, cst_price, cst_reward, bid_cst_reward_amount, cst_dutch_auction_duration, msg) VALUES
  (2010, 5077, 127, 1028, TO_TIMESTAMP(1767228300), 2, 21, -1, 3, 0, 1, TO_TIMESTAMP(1767231900),
   50000000000000000, -1, 100000000000000000000, -1, -1, 'round three begins'),
  (2011, 5078, 128, 1029, TO_TIMESTAMP(1767228400), 2, 22, -1, 3, 0, 2, TO_TIMESTAMP(1767232000),
   55000000000000000, -1, 100000000000000000000, -1, -1, ''),
  (2012, 5097, 129, 1043, TO_TIMESTAMP(1767228500), 2, 21, -1, 3, 0, 3, TO_TIMESTAMP(1767232100),
   60000000000000000, -1, 100000000000000000000, -1, -1, 'alice fifth bid');

-- Charity wallet pays out to the external charity receiver.
INSERT INTO cg_donation_sent(evtlog_id, block_num, tx_id, time_stamp, contract_aid, charity_aid, amount)
VALUES (5096, 142, 1042, TO_TIMESTAMP(1767229800), 6, 28, 50000000000000000);

-- === Moderation =============================================================
-- Bob's RandomWalk bid message was banned by an admin.
INSERT INTO cg_banned_bids(bid_id, user_addr, created_at)
VALUES (2002, '0x2200000000000000000000000000000000000022', 1767230000);
