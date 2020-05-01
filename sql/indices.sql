-- indices for cascading
CREATE INDEX tx_block_num_idx       ON  transaction     USING   btree   (block_num);
CREATE INDEX mkt_tx_id				ON	market			USING	btree	(tx_id);
CREATE INDEX sbal_tx_id				ON	sbalances		USING	btree	(tx_id);
CREATE INDEX mord_tx_id				ON	mktord			USING	btree	(tx_id);
CREATE INDEX rep_tx_id				ON	report			USING	btree	(tx_id);
CREATE INDEX vol_tx_id				ON	volume			USING	btree	(tx_id);
CREATE INDEX oichg_mtk_id			ON	oi_chg			USING	btree	(market_aid);
CREATE INDEX mktfin_id				ON	mkt_fin			USING	btree	(market_aid);
-- market_aid indices
CREATE INDEX sbal_mkt_id			ON	sbalances		USING	btree	(market_aid);
CREATE INDEX mo_mkt_id				ON	mktord			USING	btree	(market_aid);
CREATE INDEX oord_mkt_id			ON	oorders			USING	btree	(market_aid);
CREATE INDEX rep_mkt_id				ON	report			USING	btree	(market_aid);
CREATE INDEX vol_mkt_id				ON	volume			USING	btree	(market_aid);
CREATE INDEX oich_mkt_id			ON	oi_chg			USING	btree	(market_aid);
CREATE INDEX mfin_mkt_id			ON	mkt_fin			USING	btree	(market_aid);
-- other indices
CREATE INDEX blk_ph_id				ON block			USING	btree	(parent_hash);
