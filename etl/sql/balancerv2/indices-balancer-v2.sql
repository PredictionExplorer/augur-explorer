
CREATE INDEX preg_poolid_idx			ON	pool_reg					(pool_id);
CREATE INDEX swapfee_ctrct_idx			ON	swap_fee					(contract_aid);
CREATE INDEX swapfee_block_num_idx		ON	swap_fee					(block_num);
CREATE INDEX swapfee_txindex_idx		ON	swap_fee					(tx_index);
CREATE INDEX swap_bnum_idx				ON	swap						(block_num);
