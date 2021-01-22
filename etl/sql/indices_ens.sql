CREATE INDEX aname_label_idx		ON	active_name(label);
CREATE INDEX aname_expire_idx		ON	active_name(expires);
CREATE INDEX ens_name_label_idx		ON	ens_name(label);
CREATE INDEX newowner_label_idx		ON	ens_new_owner(label);
CREATE INDEX newowner_node_idx		ON	ens_new_owner(node);
CREATE INDEX ens_words_idx			ON	ens_node(fqdn_words);
CREATE INDEX ens_node_label_idx		ON	ens_node(label);
CREATE INDEX ens_reg_transf_idx		ON	ens_reg_transf(node);
CREATE INDEX ens_cur_own_idx		ON	ens_node(cur_owner_aid);
CREATE INDEX len_node_words_idx		ON	ens_node(LENGTH(fqdn_words));
CREATE INDEX ens_newown_owner_idx	ON	ens_new_owner(owner_aid);
CREATE INDEX ens_node_evtlog_idx	ON 	ens_node(evtlog_id);
