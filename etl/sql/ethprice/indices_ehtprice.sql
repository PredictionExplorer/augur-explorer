CREATE INDEX ep_block_num_idx			ON	ep_swap			(block_num);
CREATE INDEX ep_timestamp_idx			ON	ep_swap			(time_stamp);
CREATE INDEX ep_blk_tx_log_idx			ON	ep_swap			(block_num,tx_idx,log_idx);
