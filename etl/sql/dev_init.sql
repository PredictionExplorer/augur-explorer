-- Do not use, for Development only!
INSERT INTO universe(universe_addr) VALUES('0x1b8dae4F281A437E797f6213C6564926a04D9959');
INSERT INTO main_stats(universe_id) VALUES(1);
INSERT INTO contract_addresses(
		chain_id,
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
		102,
		'0x25Ff5dc79A7c4e34254ff0f4a19d69E491201DD3',-- augur (main contract)
		'0xaf517E20601Df8d8584035EB895C02713bC1f3A4',-- augur trading
		'0x658655115E55fa3433B9686865f011874BD71083',--profit loss
		'0x021076fB9adafcf83869435F9d72A5873869B4ad',--dai cash
		'0x82a37C54267b1e9D94C37895Fe26EC232aA55030',--rep token
		'0x6cfC125DF7Ba27B26138CdBCd5804137dC3BA1A6',--zerox trade
		'0x2e90e3C430C1470d62587c8983755D717f46F617',--zerox exchange
		'0x6e968FE21894A35Ba59ee8EC6f60Ea0DDC3a59E5',--wallet registry (v1)
		'0x2eF25877B254d6391B843Df25Dd7A8b0A243BEe9',--wallet registry (v2)
		'0xa2207BB135287a4EB3ae4De32A0b99d112ae57B0',--fill order
		'0x780219783B532c837f83AD633579Bdbcb39441B7',--eth exchange
		'0xE60c9fe85aEE7B4848a97271dA8c86323CdFb897',-- share token
		'0x1b8dae4F281A437E797f6213C6564926a04D9959',-- Universe
		'0xD65f9c350eD35e8339a228fF27Efd2d9BBA17C0b' -- CreateOrder
);
INSERT INTO ooconfig(osize_threshold) VALUES(0.0);
