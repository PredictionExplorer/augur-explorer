-- indices for cascading DELETEs
CREATE INDEX tx_block_num_idx       ON  transaction     (block_num);
CREATE INDEX txinp_idx				ON	tx_input		(tx_id);
CREATE INDEX evt_log_tx_idx			ON	evt_log			(tx_id);
CREATE INDEX tx_input_tx_idx		ON	tx_input		(tx_id);
-- other indices
CREATE INDEX blk_ph_idx				ON block			(parent_hash);
CREATE UNIQUE INDEX blk_hash_uniq	ON block			(block_hash);
CREATE INDEX blk_ts_idx				ON block			(ts);
CREATE INDEX pest_ts_idx			ON price_estimate	(time_stamp);
CREATE INDEX elog_ctrct_idx			ON evt_log			(contract_aid);
CREATE INDEX elog_topic0_sig		ON evt_log			(topic0_sig);
CREATE INDEX etop_val_key			ON evt_topic		(value);
CREATE INDEX etop_bnum_key			ON evt_topic		(block_num);
CREATE INDEX etop_ctrct_idx			ON evt_topic		(contract_aid);
CREATE INDEX tx_input_sig_idx		ON transaction		(input_sig);
