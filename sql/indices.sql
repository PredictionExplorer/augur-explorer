-- indices for cascading DELETEs
CREATE INDEX tx_block_num_idx       ON  transaction     USING   btree   (block_num);
CREATE INDEX mkt_tx_idx				ON	market			USING	btree	(tx_id);
CREATE INDEX sbal_tx_idx			ON	sbalances		USING	btree	(tx_id);
CREATE INDEX mord_tx_idx			ON	mktord			USING	btree	(tx_id);
CREATE INDEX rep_tx_idx				ON	report			USING	btree	(tx_id);
CREATE INDEX vol_tx_idx				ON	volume			USING	btree	(tx_id);
CREATE INDEX oichg_mtk_idx			ON	oi_chg			USING	btree	(market_aid);
CREATE INDEX mktfin_idx				ON	mkt_fin			USING	btree	(market_aid);
-- market_aid indices
CREATE INDEX sbal_mkt_idx			ON	sbalances		USING	btree	(market_aid);
CREATE INDEX mo_mkt_idx				ON	mktord			USING	btree	(market_aid);
CREATE INDEX oord_mkt_idx			ON	oorders			USING	btree	(market_aid);
CREATE INDEX rep_mkt_idx			ON	report			USING	btree	(market_aid);
CREATE INDEX vol_mkt_idx			ON	volume			USING	btree	(market_aid);
CREATE INDEX oich_mkt_idx			ON	oi_chg			USING	btree	(market_aid);
CREATE INDEX mfin_mkt_idx			ON	mkt_fin			USING	btree	(market_aid);
-- other indices
CREATE INDEX blk_ph_idx				ON block			USING	btree	(parent_hash);
CREATE INDEX mord_ts_idx			ON mktord			USING	btree	(time_stamp);
CREATE UNIQUE INDEX ovol_idx		ON outcome_vol		USING	btree	(market_aid,outcome_idx);
CREATE INDEX oo_depth_idx			ON oorders			USING	btree	(market_aid,outcome,otype);
CREATE UNIQUE INDEX oostats_uniq	ON oostats			USING	btree	(market_aid,eoa_aid,outcome_idx);
