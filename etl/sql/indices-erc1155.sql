
CREATE INDEX erc1155_bal_aid_idx				ON	erc1155_bal			(aid);
CREATE INDEX erc1155_bal_ctrct_idx				ON	erc1155_bal			(contract_aid);
CREATE INDEX erc1155_transf_from_idx			ON	erc1155_transf		(from_aid);
CREATE INDEX erc1155_transf_to_idx				ON	erc1155_transf		(to_aid);
CREATE INDEX erc1155_tx_idx						ON	erc1155_transf		(tx_id);
CREATE INDEX erc1155_tr_tok_id_idx				ON	erc1155_transf		(token_id);
CREATE INDEX erc1155_bal_tok_id_idx				ON	erc1155_bal			(token_id);
CREATE INDEX erc1155_uri_ttid_idx				ON	erc1155_uri			(token_type_id);
CREATE INDEX erc1155_ctrct_user_idx				ON	erc1155_bal			(contract_aid,aid);
