-- Do not use, for Development only!
INSERT INTO universe(universe_addr) VALUES('0x919274CE3Ba5a0D0f11071Bb9fb402AE5c23381B');
INSERT INTO universe(universe_addr) VALUES('0x2C41A5516ED977577eeAEf423e80cd98BAbC204a');
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
		42,
		18855701,
		'0xAad8776770e1f38F81DeA5d1C1404B7fdBA2DA2B',-- augur (main contract)
		'0x388B7756a0738d8Cf3a5f6F36c367842f3F163cE',-- augur trading
		'0x446D9a3B94E97cE3CF45240e1d34858D3B8443F7',-- profit loss
		'0x68e888AD28D8E41Aa07dcdEb666432Cd18b26aBa',-- dai cash
		'0x4F0738B7396dD2A9C05E42a40cF34351317C517D',-- rep token
		'0x75057695Fb8C3bB4A12eE0C6421f613640Afd0a1',-- zerox trade
		'0x4eacd0aF335451709e1e7B570B8Ea68EdEC8bc97',-- zerox exchange
		'0x7f4139C357602081D5d0B62Cb87292dD87E9d2C6',-- wallet registry (v1)
		'0x942818dfD9E191D8A129b8bD21EE7a97fF9F441D',-- wallet registry (v2)
		'0x586875DDf4041d3ada196e575A53eebd45291ed8',-- fill order
		'0x493d6E70eD7a3DDf3650Ae0A3D93029b0d4348eE',-- eth exchange
		'0x39E7A4318d4BD69499c424B571B2eC545c6740Bd',-- share token
		'0x919274CE3Ba5a0D0f11071Bb9fb402AE5c23381B',-- universe
		'0xE125E74df3691ed5Ec5bcd681c8e09A251825d82'-- CreateOrder
);
INSERT INTO last_block  VALUES(18855700);
INSERT INTO ooconfig(osize_threshold) VALUES(0.0);
