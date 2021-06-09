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
