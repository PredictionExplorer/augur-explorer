
CREATE INDEX preg_poolid_idx			ON	pool_reg					(pool_id);
CREATE INDEX swapfee_ctrct_idx			ON	swap_fee					(contract_aid);
CREATE INDEX swapfee_block_num_idx		ON	swap_fee					(block_num);
CREATE INDEX swapfee_txindex_idx		ON	swap_fee					(tx_index);
CREATE INDEX swap_bnum_idx				ON	swap						(block_num);
CREATE INDEX block_block_hash_idx		ON	block						(block_hash);
CREATE INDEX swap_fees_idx				ON	swap						(token_in_aid,token_out_aid,time_stamp DESC);
CREATE INDEX tok_bal_pooltok_idx		ON	tok_bal						(pool_aid,tok_aid);
CREATE INDEX tok_bal_block_num_idx		ON	tok_bal						(block_num);
CREATE INDEX tok_bal_txindex_idx		ON	tok_bal						(tx_index);
CREATE INDEX bpt_from_idx				ON	bpt_transf					(pool_aid,from_aid);
CREATE INDEX bpt_to_idx					ON	bpt_transf					(pool_aid,to_aid);
CREATE INDEX bpt_bal_aid_idx			ON	bpt_bal						(pool_aid,aid);
CREATE INDEX swfh_pool_aid_idx			ON	swf_hist					(pool_aid);
CREATE INDEX swfh_pool_aid_idx2			ON	swf_hist					(pool_aid,time_stamp);
