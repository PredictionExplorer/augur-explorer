-- Do not use, for Development only!
INSERT INTO universe(universe_addr) VALUES('0x919274CE3Ba5a0D0f11071Bb9fb402AE5c23381B');
INSERT INTO universe(universe_addr) VALUES('0x2C41A5516ED977577eeAEf423e80cd98BAbC204a');
INSERT INTO main_stats(universe_id) VALUES(1);
INSERT INTO contract_addresses(
		dai_cash,
		rep_token,
		zerox,
		wallet_reg,
		fill_order,
		eth_xchg,
		share_token,
		universe
) VALUES (
		'0x68e888AD28D8E41Aa07dcdEb666432Cd18b26aBa',-- dai cash
		'0x4F0738B7396dD2A9C05E42a40cF34351317C517D',-- rep token
		'0x75057695Fb8C3bB4A12eE0C6421f613640Afd0a1',-- zerox
		'0x7f4139C357602081D5d0B62Cb87292dD87E9d2C6',-- wallet registry
		'0x586875DDf4041d3ada196e575A53eebd45291ed8',-- fill order
		'0x493d6E70eD7a3DDf3650Ae0A3D93029b0d4348eE',-- eth exchange
		'0x39E7A4318d4BD69499c424B571B2eC545c6740Bd',-- share token
		'0x919274CE3Ba5a0D0f11071Bb9fb402AE5c23381B'-- universe
);
INSERT INTO last_block  VALUES(18855700);

