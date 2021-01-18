
-- indices for cascading DELETEs
CREATE INDEX dai_transf_idx			ON	dai_transf		(evtlog_id);
CREATE INDEX daib_idx				ON	dai_bal			(dai_transf_id);
CREATE INDEX rep_idx				ON	rep_transf		(evtlog_id);
-- market_aid indices

-- other indices
CREATE INDEX daib_processed_idx		ON dai_bal			(processed);

-- 
CREATE INDEX wstok_from_idx			ON wstok_transf		(from_aid);
CREATE INDEX wstok_to_idx			ON wstok_transf		(to_aid);
CREATE INDEX wstok_wrapper_aid_idx	ON wstok_transf		(wrapper_aid);
CREATE INDEX ethusd_price_ts_idx	ON ethusd_price		(time_stamp);
