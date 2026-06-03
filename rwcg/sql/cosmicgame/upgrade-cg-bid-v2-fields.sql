-- Apply on existing databases before running ETL with CosmicSignatureGameV2 BidPlaced support.
-- ETL routes by event topic (0xbcb004d6… legacy, 0x1d1f406c… V2); no upgrade-block bifurcation needed.
ALTER TABLE cg_bid ADD COLUMN IF NOT EXISTS bid_cst_reward_amount DECIMAL DEFAULT -1;
ALTER TABLE cg_bid ADD COLUMN IF NOT EXISTS cst_dutch_auction_duration DECIMAL DEFAULT -1;
