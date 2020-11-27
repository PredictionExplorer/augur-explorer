
-- indices for cascading DELETEs

-- pool_aid indices
CREATE INDEX bpool_aid_idx			ON	bpool			(pool_aid);
CREATE INDEX bswap_pool_aid_idx		ON	bswap			(pool_aid);
CREATE INDEX btoken_pool_aid_idx	ON	btoken			(pool_aid);
CREATE INDEX bjoin_pool_aid_idx		ON	bjoin			(pool_aid);
CREATE INDEX bexit_pool_aid_idx		ON	bexit			(pool_aid);
CREATE INDEX bholder_pool_idx		ON	bholder			(pool_aid);
CREATE INDEX b_set_public_pool_idx	ON	b_set_public	(pool_aid);
CREATE INDEX b_finalized			ON	b_finalized		(pool_aid);
CREATE INDEX b_bind					ON	b_bind			(pool_aid);
CREATE INDEX b_unbind				ON	b_unbind		(pool_aid);
CREATE INDEX b_rebind				ON	b_rebind		(pool_aid);
CREATE INDEX b_gulp					ON	b_gulp			(pool_aid);

-- other indices
CREATE INDEX btoken					ON	btoken			(token_aid);
CREATE INDEX bswap_token_in_idx     ON	bswap			(token_in_aid);
CREATE INDEX bswap_token_out_idx    ON	bswap			(token_out_aid);
CREATE INDEX bswap_txid_idx			ON	bswap			(tx_id);
