CREATE INDEX swap_pool_aid_idx          ON  swap                    (pool_aid);
CREATE INDEX collect_pool_idx			ON	collect					(pool_aid);
CREATE INDEX collect_pool_sorted_idx	ON	collect					(pool_aid,time_stamp DESC);
