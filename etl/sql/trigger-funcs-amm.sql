CREATE OR REPLACE FUNCTION on_liquidity_changed_insert() RETURNS trigger AS  $$
DECLARE
BEGIN
	UPDATE aa_sports_market SET liquidity = liquidity - NEW.collateral
		WHERE contract_aid=NEW.factory_aid AND market_id=NEW.market_id;
	RETURN NEW;
END;
$$ LANGUAGE plpgsql;
CREATE OR REPLACE FUNCTION on_liquidity_changed_delete() RETURNS trigger AS  $$
DECLARE
BEGIN

	UPDATE aa_sports_market SET liquidity = liquidity + NEW.collateral
		WHERE contract_aid=NEW.factory_aid AND market_id=NEW.market_id;
	RETURN OLD;
END;
$$ LANGUAGE plpgsql;
CREATE OR REPLACE FUNCTION amm_insert_abstract_market(
	p_evtlog_id BIGINT,p_block_num BIGINT,p_tx_id BIGINT,p_contract_aid BIGINT,
	p_market_id BIGINT,p_timestamp TIMESTAMPTZ,p_sharetoken_ids TEXT
) RETURNS void AS  $$
DECLARE
BEGIN
	
END;
