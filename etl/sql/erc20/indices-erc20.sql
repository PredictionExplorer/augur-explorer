
CREATE INDEX erc20_ctrct_user_idx		ON	erc20_bal			(contract_aid,aid);
CREATE INDEX erc20_bal_parent_idx		ON	erc20_bal			(parent_id);
CREATE INDEX erc20_bal_aid_idx			ON	erc20_bal			(aid);
CREATE INDEX erc20_transf_from_idx		ON	erc20_transf		(from_aid);
CREATE INDEX erc20_transf_to_idx		ON	erc20_transf		(to_aid);
CREATE INDEX erc20_tx_idx				ON	erc20_transf		(tx_id);
CREATE INDEX erc20_bal_ctrct_idx		ON	erc20_bal			(contract_aid);
CREATE INDEX erc20_tr_ctrct_idx			ON	erc20_transf		(contract_aid);
