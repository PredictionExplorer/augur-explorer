
-- indices for cascading DELETEs
CREATE INDEX dai_transf_idx			ON	dai_transf		(evtlog_id);
CREATE INDEX daib_idx				ON	dai_bal			(dai_transf_id);
CREATE INDEX rep_idx				ON	rep_transf		(evtlog_id);
-- market_aid indices

-- other indices
CREATE INDEX daib_processed_idx		ON dai_bal			(processed);
