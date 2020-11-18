CREATE INDEX mkt_tx_idx				ON	market			(tx_id);
CREATE INDEX mord_tx_idx			ON	mktord			(tx_id);
CREATE INDEX rep_tx_idx				ON	report			(tx_id);
CREATE INDEX vol_tx_idx				ON	volume			(tx_id);
CREATE INDEX mktfin_idx				ON	mkt_fin			(tx_id);
CREATE INDEX outc_vol_idx			ON	outcome_vol		(tx_id);
CREATE INDEX oichg_tx_idx			ON	oi_chg			(tx_id);
CREATE INDEX pl_idx					ON	profit_loss		(tx_id);
CREATE INDEX cf_idx					ON	claim_funds		(tx_id);
CREATE INDEX agtx_idx				ON	agtx_status		(tx_id);
CREATE INDEX execwtx_idx			ON	exec_wtx		(tx_id);
CREATE INDEX regctrct_idx			ON	register_contract	(tx_id);
CREATE INDEX tbc_idx				ON	tbc				(tx_id);
CREATE INDEX stbc_idx				ON	stbc			(tx_id);
CREATE INDEX toktr_idx				ON	tok_transf		(tx_id);
CREATE INDEX sbal_tx_idx			ON	sbalances		(tx_id);
CREATE INDEX pldbg_idx				ON	pl_debug		(block_num);
-- market_aid indices
CREATE INDEX mo_mkt_idx				ON	mktord			(market_aid);
CREATE INDEX oord_mkt_idx			ON	oorders			(market_aid);
CREATE INDEX rep_mkt_idx			ON	report			(market_aid);
CREATE INDEX vol_mkt_idx			ON	volume			(market_aid);
CREATE INDEX outc_vol_mkt_idx		ON	outcome_vol		(market_aid);
CREATE INDEX oich_mkt_idx			ON	oi_chg			(market_aid);
CREATE INDEX mfin_mkt_idx			ON	mkt_fin			(market_aid);
CREATE INDEX mpl_mkt_idx			ON	profit_loss		(market_aid);
CREATE INDEX cf_mkt_idx				ON	claim_funds		(market_aid);
CREATE INDEX trd_mkt_stat_idx		ON	trd_mkt_stats	(market_aid);
CREATE INDEX sbal_mkt_idx			ON	sbalances		(market_aid);
CREATE INDEX tbc_mkt_idx			ON	tbc				(market_aid);
CREATE INDEX tokt_mtk_idx			ON	tok_transf		(market_aid);
-- aid indices
CREATE INDEX uranks_aid_idx			ON	uranks			(aid);
CREATE UNIQUE INDEX ustats1_idx		ON	ustats			(aid);
CREATE INDEX pl_aid_idx				ON	profit_loss		(aid);
CREATE INDEX exec_wtx_aid_idx		ON	exec_wtx		(eoa_aid);
-- other indices
CREATE INDEX mord_ts_idx			ON mktord			(time_stamp);
CREATE UNIQUE INDEX ovol_idx		ON outcome_vol		(market_aid,outcome_idx);
CREATE INDEX oo_depth_idx			ON oorders			(market_aid,outcome_idx,otype);
CREATE INDEX oo_uniq				ON oorders			(order_hash);
CREATE UNIQUE INDEX oostats_uniq	ON oostats			(market_aid,aid,outcome_idx);
CREATE UNIQUE INDEX tmstats_uniq	ON trd_mkt_stats	(market_aid,aid);
CREATE INDEX pl_profit_srch_idx		ON profit_loss		(market_aid,aid,outcome_idx);
CREATE INDEX open_positions_idx		ON profit_loss		(aid,realized_profit) WHERE realized_profit = 0.0;
CREATE INDEX closed_positions_idx	ON profit_loss		(aid,realized_profit) WHERE realized_profit <> 0.0;
CREATE UNIQUE INDEX cl_uniq			ON claim_funds		(aid,market_aid,outcome_idx);
CREATE UNIQUE INDEX mkts_traded_unq	ON mkts_traded		(aid,market_aid);
CREATE UNIQUE INDEX sbal_uniq		ON sbalances		(market_aid,account_aid,outcome_idx);
CREATE INDEX exec_wtx_in_sig_idx	ON exec_wtx			(input_sig);
CREATE INDEX exec_wtx_referral_idx	ON exec_wtx			(referral_aid);
CREATE INDEX exec_wtx_to_idx		ON exec_wtx			(to_aid);
CREATE INDEX mktord_ts_idx			ON mktord			(time_stamp);
CREATE UNIQUE INDEX pldebug_uniq	ON pl_debug			(block_num,market_aid,aid,outcome_idx);
CREATE INDEX mkt_tsv_idx			ON mkt_words		USING gin(tokens);
CREATE UNIQUE INDEX mktwrds_uniq1	ON mkt_words		(market_aid) WHERE market_aid IS NOT NULL;
CREATE UNIQUE INDEX mktwrds_uniq2	ON mkt_words		(cat_id) WHERE cat_id IS NOT NULL;

