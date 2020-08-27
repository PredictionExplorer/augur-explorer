-- Do not use, for Development only!
INSERT INTO universe(universe_addr) VALUES('0x1b8dae4F281A437E797f6213C6564926a04D9959');
INSERT INTO main_stats(universe_id) VALUES(1);
INSERT INTO contract_addresses(
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
		'0xFCaf25bF38E7C86612a25ff18CB8e09aB07c9885',-- augur (main contract)
		'0x25Ff5dc79A7c4e34254ff0f4a19d69E491201DD3',-- augur trading
		'0x658655115E55fa3433B9686865f011874BD71083',--profit loss
		'0x3a043108953aACf3505503867F8Db7C1585577c7',--dai cash
		'0x82a37C54267b1e9D94C37895Fe26EC232aA55030',--rep token
		'0x6cfC125DF7Ba27B26138CdBCd5804137dC3BA1A6',--zerox trade
		'0x2F8F57DBd52Ee20f9258C5232cfbbAb0f8D478F1',--zerox exchange
		'0x6e968FE21894A35Ba59ee8EC6f60Ea0DDC3a59E5',--wallet registry (v1)
		'0x2eF25877B254d6391B843Df25Dd7A8b0A243BEe9',--wallet registry (v2)
		'0xa2207BB135287a4EB3ae4De32A0b99d112ae57B0',--fill order
		'0xaB120ec64Bf11438AF32C1586E28e2Bcea95E6df',--eth exchange
		'0xE60c9fe85aEE7B4848a97271dA8c86323CdFb897',-- share token
		'0x1b8dae4F281A437E797f6213C6564926a04D9959',-- Universe
		'0xD65f9c350eD35e8339a228fF27Efd2d9BBA17C0b' -- CreateOrder
);
INSERT INTO ooconfig(osize_threshold) VALUES(0.0);
