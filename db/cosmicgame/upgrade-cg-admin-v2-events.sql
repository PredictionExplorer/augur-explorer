-- Apply on existing databases before running ETL with CosmicSignatureGameV2 admin event support.
CREATE TABLE IF NOT EXISTS cg_adm_cst_auclen_chg_div (
	id              BIGSERIAL PRIMARY KEY,
	evtlog_id       BIGINT REFERENCES evt_log(id) ON DELETE CASCADE,
	block_num       BIGINT NOT NULL,
	tx_id           BIGINT NOT NULL,
	time_stamp      TIMESTAMPTZ NOT NULL,
	contract_aid    BIGINT NOT NULL,
	new_len			DECIMAL NOT NULL,
	UNIQUE(evtlog_id)
);
