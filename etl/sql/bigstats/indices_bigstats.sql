CREATE INDEX addr_acct_type				ON	bs_addr			(is_contract);
CREATE INDEX bs_log_block_idx			ON	bs_log			(block_num);
CREATE INDEX block_ts_idx				ON	bs_block		(ts);
