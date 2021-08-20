
CREATE INDEX pol_buysell_uaid_idx		ON	pol_buysell			(user_aid);
CREATE INDEX pol_buysell_ctrct_idx		ON	pol_buysell			(contract_aid);
CREATE INDEX pol_fund_addrem_ctrct_idx	ON	pol_fund_addrem		(contract_aid);
CREATE INDEX pol_mktmkr_idx				ON  pol_market			(mkt_mkr_aid);
CREATE INDEX pol_cond_split_idx			ON	pol_pos_split		(condition_id);
CREATE INDEX pol_cond_split_ctrct_idx	ON	pol_pos_split		(contract_aid);
CREATE INDEX pol_cond_merge_ctrct_idx	ON	pol_pos_merge		(contract_aid);
CREATE INDEX pol_cond_prep_quest_idx	ON	pol_cond_prep		(question_id);
CREATE INDEX pol_cond_res_quest_idx		ON	pol_cond_res		(question_id);
