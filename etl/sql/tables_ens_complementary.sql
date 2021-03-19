CREATE TABLE alexa_top1m(	-- Alexa's top 1M domain names, about 700k records
	name				TEXT,
	hash				TEXT UNIQUE	-- label hash
);
CREATE TABLE en_prop_names(	-- English proper names (list of 61k words)
	word				TEXT,
	hash				TEXT UNIQUE	-- label hash
);
CREATE TABLE email_tokens( -- Words extracted from 300million emails list dataset
	token				TEXT,
	hash				TEXT UNIQUE	-- label hash
);
CREATE TABLE pwd_db ( -- 36 million record password database
	password			TEXT,
	hash				TEXT UNIQUE	-- label hash
);
