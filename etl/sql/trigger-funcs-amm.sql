CREATE OR REPLACE FUNCTION on_liquidity_changed_insert() RETURNS trigger AS  $$
DECLARE
BEGIN
	UPDATE aa_market SET liquidity = liquidity - NEW.collateral
		WHERE contract_aid=NEW.factory_aid AND market_id=NEW.market_id;
	RETURN NEW;
END;
$$ LANGUAGE plpgsql;
CREATE OR REPLACE FUNCTION on_liquidity_changed_delete() RETURNS trigger AS  $$
DECLARE
BEGIN

	UPDATE aa_market SET liquidity = liquidity + NEW.collateral
		WHERE contract_aid=NEW.factory_aid AND market_id=NEW.market_id;
	RETURN OLD;
END;
$$ LANGUAGE plpgsql;
CREATE OR REPLACE FUNCTION amm_insert_abstract_market(
	p_evtlog_id BIGINT,p_block_num BIGINT,p_tx_id BIGINT,p_contract_aid BIGINT,p_timestamp TIMESTAMPTZ,
	p_created_time TIMESTAMPTZ,p_end_time TIMESTAMPTZ,p_market_id BIGINT,p_factory_aid BIGINT,
	p_collateral_aid BIGINT,p_protocol_aid BIGINT,p_settlement_aid BIGINT,p_feepot_aid BIGINT,
	p_settlement_fee DECIMAL,p_staker_fee DECIMAL,p_protocol_fee DECIMAL,
	p_sharefactor DECIMAL,p_sharetokens TEXT
) RETURNS BIGINT AS  $$
DECLARE
	v_aa_mkt_id BIGINT;
	v_token_addr TEXT;
	v_token_aid BIGINT;
BEGIN
	INSERT INTO aa_market(
		evtlog_id,block_num,tx_id,contract_aid,time_stamp,
		created_time,end_time,market_id,factory_aid,
		collateral_aid,protocol_aid,settlement_aid,feepot_aid,
		settlement_fee,staker_fee,protocol_fee,
		sharefactor
	) VALUES(
		p_evtlog_id,p_block_num,p_tx_id,p_contract_aid,p_timestamp,
		p_created_time,p_end_time,p_market_id,p_factory_aid,
		p_collateral_aid,p_protocol_aid,p_settlement_aid,p_feepot_aid,
		p_settlement_fee,p_staker_fee,p_protocol_fee,
		p_sharefactor
	) RETURNING id INTO v_aa_mkt_id;
	FOREACH v_token_addr IN ARRAY STRING_TO_ARRAY(p_sharetokens,',')
		LOOP
			INSERT INTO address(block_num,tx_id,addr) VALUES(p_block_num,p_tx_id,v_token_addr)
				ON CONFLICT DO NOTHING
				RETURNING address_id INTO v_token_aid ;
			IF v_token_aid IS NULL THEN
				SELECT address_id FROM address WHERE addr=v_token_addr INTO v_token_aid;
				IF v_token_aid IS NULL THEN
					RAISE EXCEPTION 'can''t find address % in address table and INSERT failed',v_token_addr;
				END IF;
			END IF;
		END LOOP;

	RETURN v_aa_mkt_id;
END;
$$ LANGUAGE plpgsql;
CREATE OR REPLACE FUNCTION on_before_sports_market_insert() RETURNS trigger AS  $$
DECLARE
	v_market_id BIGINT;
BEGIN

END;
$$ LANGUAGE plpgsql;
CREATE OR REPLACE FUNCTION on_after_sports_market_delete() RETURNS trigger AS  $$
DECLARE
BEGIN

	DELETE FROM aa_market WHERE market_id = OLD.market_id;
	RETURN OLD;
END;
$$ LANGUAGE plpgsql;
CREATE OR REPLACE FUNCTION amm_insert_sports_market(
	-- basic Ethereum params
	p_evtlog_id BIGINT,p_block_num BIGINT,p_tx_id BIGINT,p_contract_aid BIGINT,p_time_stamp TIMESTAMPTZ,
	-- Abstract Market params
	p_market_id BIGINT,p_creator_aid BIGINT,p_created_time TIMESTAMPTZ,p_end_time TIMESTAMPTZ,
	p_settlement_fee DECIMAL,p_staker_fee DECIMAL,p_protocol_fee DECIMAL,
	p_settlement_aid BIGINT,p_feepot_aid BIGINT,p_protocol_aid BIGINT,p_collateral_aid BIGINT,
	p_sharefactor DECIMAL,p_sharetokens TEXT,
	-- Specific Market (child object) params
	p_event_id BIGINT,p_home_team_id BIGINT,p_away_team_id BIGINT,
	p_estimated_start TIMESTAMPTZ,p_market_type INT,p_value0 DECIMAL
) RETURNS void AS $$
DECLARE
	v_aa_mkt_id BIGINT;
BEGIN

	SELECT amm_insert_abstract_market(
		p_evtlog_id,p_block_num,p_tx_id,p_contract_aid,	p_time_stamp,
		p_created_time,p_end_time,p_market_id,p_contract_aid,
		p_collateral_aid,p_protocol_aid,p_settlement_aid,p_feepot_aid,
		p_settlement_fee,p_staker_fee,p_protocol_fee,
		p_sharefactor,p_sharetokens
	) INTO v_aa_mkt_id;

	INSERT INTO aa_sports_market (
		evtlog_id,block_num,tx_id,contract_aid,time_stamp,
		end_time,est_start_time,
		market_id,creator_aid,aa_mkt_id,event_id,home_team_id,away_team_id,market_type,value0
	) VALUES (
		p_evtlog_id,p_block_num,p_tx_id,p_contract_aid,p_time_stamp,
		p_end_time,p_estimated_start,
		p_market_id,p_creator_aid,v_aa_mkt_id,p_event_id,p_home_team_id,p_away_team_id,
		p_market_type,p_value0
	);

END;
$$ LANGUAGE plpgsql;
CREATE OR REPLACE FUNCTION amm_insert_trusted_market(
	-- basic Ethereum params
	p_evtlog_id BIGINT,p_block_num BIGINT,p_tx_id BIGINT,p_contract_aid BIGINT,p_time_stamp TIMESTAMPTZ,
	-- Abstract Market params
	p_market_id BIGINT,p_creator_aid BIGINT,p_created_time TIMESTAMPTZ,p_end_time TIMESTAMPTZ,
	p_settlement_fee DECIMAL,p_staker_fee DECIMAL,p_protocol_fee DECIMAL,
	p_settlement_aid BIGINT,p_feepot_aid BIGINT,p_protocol_aid BIGINT,p_collateral_aid BIGINT,
	p_sharefactor DECIMAL,p_sharetokens TEXT,
	-- Specific Market (child object) params
	p_descr TEXT,p_outcomes TEXT
) RETURNS void AS $$
DECLARE
	v_aa_mkt_id BIGINT;
BEGIN

	SELECT amm_insert_abstract_market(
		p_evtlog_id,p_block_num,p_tx_id,p_contract_aid,	p_time_stamp,
		p_created_time,p_end_time,p_market_id,p_contract_aid,
		p_collateral_aid,p_protocol_aid,p_settlement_aid,p_feepot_aid,
		p_settlement_fee,p_staker_fee,p_protocol_fee,
		p_sharefactor,p_sharetokens
	) INTO v_aa_mkt_id;

	INSERT INTO aa_trusted_market (
		evtlog_id,block_num,tx_id,contract_aid,time_stamp,
		end_time,market_id,creator_aid,aa_mkt_id,
		descr,outcomes
	) VALUES (
		p_evtlog_id,p_block_num,p_tx_id,p_contract_aid,p_time_stamp,
		p_end_time,	p_market_id,p_creator_aid,v_aa_mkt_id,
		p_descr,p_outcomes
	);

END;
$$ LANGUAGE plpgsql;
