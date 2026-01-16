-- pair aid
CREATE INDEX upair_pair_aid_idx		ON	upair		(pair_aid);
CREATE INDEX uswap1_pair_aid_idx	ON	uswap1		(pair_aid);


CREATE INDEX unwap2_id1_idx			ON	uswap2		(uswap1_id);
CREATE INDEX uswap2_token_aid_idx	ON	uswap2		(token_aid);
CREATE INDEX upair_evt_idx			ON	upair		(evtlog_id);
CREATE INDEX upair_token0_idx		ON	upair		(token0_aid);
CREATE INDEX upair_token1_idx		ON	upair		(token1_aid);
CREATE INDEX uswap2_txid_idx		ON	uswap2		(tx_id);
CREATE INDEX uswap1_tx_id_idx		ON	uswap1		(tx_id);
CREATE INDEX uswqp1_aid_idx			ON	uswap1		(recipient_aid);

CREATE INDEX uswap2_evtlog_idx		ON	uswap2		(evtlog_id);
