CREATE INDEX aname_label_idx		ON	active_name(label);
CREATE INDEX aname_expire_idx		ON	active_name(expires);
CREATE INDEX ens_name_label_idx		ON	ens_name(label);
CREATE INDEX newowner_label_idx		ON	ens_new_owner(label);
CREATE INDEX newowner_node_idx		ON	ens_new_owner(node);
