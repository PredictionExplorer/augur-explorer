-- Do not use, for Development only!
INSERT INTO universe(universe_addr) VALUES('0x50F0b9e78dc52A834EEd7378DC98a57a5a3C9900');
INSERT INTO main_stats(universe_id) VALUES(1);
INSERT INTO contract_addresses(
		augur,
		augur_trading,
		profit_loss,
		dai_cash,
		rep_token,
		zerox,
		wallet_reg,
		fill_order,
		eth_xchg,
		share_token,
		universe
) VALUES (
		'0xbc362B362dF54f77e6786d63652bc82245cF1af4',-- augur (main contract)
		'0x9a9bd63f57c264893E6a01e1Ca737e2265f0266B',-- augur trading
		'0xc7830Ac6DE781352F7CaE877025AC7368d0Eff94',-- profit loss
		'0xb6085Abd65E21d205AEaD0b1b9981B8B221fA14E',-- dai cash
		'0x08f7904Ab81CFA0615C9D73f69F0c17521B36Fb0',-- rep token
		'0x5Fe53EB2C230C5dbCd0Aa14C727B1797536Df5c9',-- zerox
		'0x17069733eD2aC451BE4f7A559E4C7d960b061C41',-- wallet registry
		'0x4f2E91D2A29D77c040AAA4527378C3FCb7aDc000',-- fill order
		'0xFacdc564F29C751d4d6315B8d2706087AF1BED3d',-- eth exchange
		'0xFFaFFda91CF0a333DA64a046B1506213A053B142',-- share token
		'0x50F0b9e78dc52A834EEd7378DC98a57a5a3C9900'-- universe
);
INSERT INTO last_block  VALUES(19712444);
