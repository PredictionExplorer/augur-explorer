-- Do not use, for Development only!
INSERT INTO universe(universe_addr) VALUES('0x82a37C54267b1e9D94C37895Fe26EC232aA55030');
INSERT INTO main_stats(universe_id) VALUES(1);
INSERT INTO contract_addresses(
		dai_cash,
		rep_token,
		zerox,
		wallet_reg,
		fill_order,
		eth_xchg,
		share_token
) VALUES (
		'0x3a043108953aACf3505503867F8Db7C1585577c7',--dai cash
		'0x82a37C54267b1e9D94C37895Fe26EC232aA55030',--rep token
		'0x6cfC125DF7Ba27B26138CdBCd5804137dC3BA1A6',--zerox
		'0x6e968FE21894A35Ba59ee8EC6f60Ea0DDC3a59E5',--wallet registry (v1)
		'0xa2207BB135287a4EB3ae4De32A0b99d112ae57B0',--fill order
		'0x36A829b02Ab0bAF4eB3Ffe544335472FC45C0eA1',--eth exchange
		'0xE60c9fe85aEE7B4848a97271dA8c86323CdFb897'-- share token
);
