-- Do not use, for Development only!
INSERT INTO universe(universe_addr) VALUES('0x50F0b9e78dc52A834EEd7378DC98a57a5a3C9900');
INSERT INTO main_stats(universe_id) VALUES(1);
INSERT INTO contract_addresses(
		chain_id,
		upload_block,
		augur,
		augur_trading,
		profit_loss,
		dai_cash,
		rep_token,
		zerox_trade,
		zerox_xchg,
		wallet_reg,
		wallet_reg2,
		fill_order,
		eth_xchg,
		share_token,
		universe,
		create_order
) VALUES (
		42
		19712445,
		'0xbc362B362dF54f77e6786d63652bc82245cF1af4',-- augur (main contract)
		'0x9a9bd63f57c264893E6a01e1Ca737e2265f0266B',-- augur trading
		'0xc7830Ac6DE781352F7CaE877025AC7368d0Eff94',-- profit loss
		'0xb6085Abd65E21d205AEaD0b1b9981B8B221fA14E',-- dai cash
		'0x7859eF3B6D0Be006675990192Eea09cba46d9C35',-- rep token
		'0x5Fe53EB2C230C5dbCd0Aa14C727B1797536Df5c9',-- zerox trade
		'0x4eacd0aF335451709e1e7B570B8Ea68EdEC8bc97',-- zerox exchange
		'0x17069733eD2aC451BE4f7A559E4C7d960b061C41',-- wallet registry (v1)
		'0x93E4533C9eF5A44D12D679AE0257322Ec6F9007C',-- wallet registry (v2)
		'0x4f2E91D2A29D77c040AAA4527378C3FCb7aDc000',-- fill order
		'0xFacdc564F29C751d4d6315B8d2706087AF1BED3d',-- eth exchange
		'0xFFaFFda91CF0a333DA64a046B1506213A053B142',-- share token
		'0x50F0b9e78dc52A834EEd7378DC98a57a5a3C9900',-- universe
		'0xB3E80567fa4655A2fa67dD2eC92a6dE597156848'-- CreateOrder
);
INSERT INTO last_block  VALUES(19712444);
