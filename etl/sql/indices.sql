-- indices for cascading DELETEs
CREATE INDEX tx_block_num_idx       ON  transaction     (block_num);
CREATE INDEX mkt_tx_idx				ON	market			(tx_id);
CREATE INDEX sbal_tx_idx			ON	sbalances		(tx_id);
CREATE INDEX mord_tx_idx			ON	mktord			(tx_id);
CREATE INDEX rep_tx_idx				ON	report			(tx_id);
CREATE INDEX vol_tx_idx				ON	volume			(tx_id);
CREATE INDEX oichg_mtk_idx			ON	oi_chg			(market_aid);
CREATE INDEX mktfin_idx				ON	mkt_fin			(market_aid);
CREATE INDEX outc_vol_idx			ON	outcome_vol		(market_aid);
CREATE INDEX pl_idx					ON	profit_loss		(tx_id);
CREATE INDEX cf_idx					ON	claim_funds		(tx_id);
CREATE INDEX daib_idx				ON	dai_bal			(tx_id);
CREATE INDEX dait_idx				ON	dai_transf		(tx_id);
CREATE INDEX rep_idx				ON	rep_transf		(tx_id);
CREATE INDEX tbc_idx				ON	tbc				(tx_id);
CREATE INDEX toktr_idx				ON	tok_transf		(tx_id);
CREATE INDEX pldbg_idx				ON	pl_debug		(block_num);
-- market_aid indices
CREATE INDEX sbal_mkt_idx			ON	sbalances		(market_aid);
CREATE INDEX mo_mkt_idx				ON	mktord			(market_aid);
CREATE INDEX oord_mkt_idx			ON	oorders			(market_aid);
CREATE INDEX rep_mkt_idx			ON	report			(market_aid);
CREATE INDEX vol_mkt_idx			ON	volume			(market_aid);
CREATE INDEX oich_mkt_idx			ON	oi_chg			(market_aid);
CREATE INDEX mfin_mkt_idx			ON	mkt_fin			(market_aid);
CREATE INDEX mpl_mkt_idx			ON	profit_loss		(market_aid);
CREATE INDEX cf_mkt_idx				ON	claim_funds		(market_aid);
CREATE INDEX trd_mkt_stat_idx		ON	trd_mkt_stats	(market_aid);
CREATE INDEX tbc_mkt_idx			ON	tbc				(market_aid);
CREATE INDEX tokt_mtk_idx			ON	tok_transf		(market_aid);
-- eoa_aid indices	(pure EOA, not composite)
CREATE INDEX uranks_eoa_idx			ON	uranks			(eoa_aid);
CREATE UNIQUE INDEX ustats1_idx		ON ustats			(eoa_aid);
CREATE INDEX pl_eoa_idx				ON profit_loss		(eoa_aid);
-- wallet aid indices (pure Wallet contract, not composite indices)
CREATE UNIQUE INDEX ustats2_idx		ON ustats			(wallet_aid);
CREATE INDEX pl_wallet_idx			ON profit_loss		(wallet_aid);
-- other indices
CREATE INDEX blk_ph_idx				ON block			(parent_hash);
CREATE UNIQUE INDEX blk_hash_uniq	ON block			(block_hash);
CREATE INDEX mord_ts_idx			ON mktord			(time_stamp);
CREATE UNIQUE INDEX ovol_idx		ON outcome_vol		(market_aid,outcome_idx);
CREATE INDEX oo_depth_idx			ON oorders			(market_aid,outcome_idx,otype);
CREATE UNIQUE INDEX oostats_uniq	ON oostats			(market_aid,eoa_aid,outcome_idx);
CREATE UNIQUE INDEX tmstats_uniq	ON trd_mkt_stats	(market_aid,eoa_aid);
CREATE INDEX pl_profit_srch_idx		ON profit_loss		(market_aid,eoa_aid,outcome_idx);
CREATE INDEX open_positions_idx		ON profit_loss		(eoa_aid,realized_profit) WHERE realized_profit = 0.0;
CREATE INDEX closed_positions_idx	ON profit_loss		(eoa_aid,realized_profit) WHERE realized_profit <> 0.0;
CREATE UNIQUE INDEX cl_uniq			ON claim_funds		(eoa_aid,market_aid,outcome_idx);
CREATE UNIQUE INDEX pldebug_uniq	ON pl_debug			(block_num,market_aid,wallet_aid,outcome_idx);
CREATE UNIQUE INDEX mkts_traded_unq	ON mkts_traded		(eoa_aid,market_aid);
CREATE UNIQUE INDEX sbal_uniq		ON sbalances		(market_aid,account_aid,outcome_idx);
CREATE INDEX daib_processed_idx		ON	dai_bal			(processed);
