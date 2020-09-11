
-- indices for cascading DELETEs
CREATE INDEX dai_transf_idx			ON	dai_transf		(evtlog_id);
CREATE INDEX daib_idx				ON	dai_bal			(dai_transf_id);
CREATE INDEX rep_idx				ON	rep_transf		(evtlog_id);
CREATE INDEX tbc_idx				ON	tbc				(evtlog_id);
CREATE INDEX stbc_idx				ON	stbc			(evtlog_id);
CREATE INDEX toktr_idx				ON	tok_transf		(evtlog_id);
CREATE INDEX sbal_tx_idx			ON	sbalances		(evtlog_id);
-- market_aid indices
CREATE INDEX sbal_mkt_idx			ON	sbalances		(market_aid);
CREATE INDEX tbc_mkt_idx			ON	tbc				(market_aid);
CREATE INDEX tokt_mtk_idx			ON	tok_transf		(market_aid);

-- eoa_aid indices	(pure EOA, not composite)

-- other indices
CREATE INDEX daib_processed_idx		ON dai_bal			(processed);
CREATE UNIQUE INDEX sbal_uniq		ON sbalances		(market_aid,account_aid,outcome_idx);
