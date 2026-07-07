-- Store-suite extension seed: rw_uranks percentile rankings.
--
-- Production fills this table from the rwctl top-rated cron
-- (Update_randomwalk_top_*_rank), not from ETL triggers, so the shared
-- dataset leaves it empty. Seed carol (sold #10 for 1 ETH over a 0.05 mint)
-- and dave (bought #10) so the top-profit/trades/volume readers and
-- Get_random_walk_stats' unique-user count have real rows to pin.
INSERT INTO rw_uranks(aid, total_trades, top_profit, top_trades, top_volume, profit, volume) VALUES
  (23, 1, 1.0, 1.0, 1.0, 950000000000000000, 1000000000000000000),
  (24, 1, 50.0, 1.0, 1.0, 0.0, 1000000000000000000);
