
CREATE INDEX bal_aid_idx				ON	erc20_bal			(aid);
CREATE INDEX erc20_transf_from_idx		ON	erc20_transf		(from_aid);
CREATE INDEX erc20_transf_to_idx		ON	erc20_transf		(to_aid);
CREATE INDEX erc20_tx_idx				ON	erc20_transf		(tx_id);
