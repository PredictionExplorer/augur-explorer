-- Do not use, for Development only!
INSERT INTO universe(universe_addr) VALUES('0x49244BD018Ca9fd1f06ecC07B9E9De773246e5AA');
INSERT INTO main_stats(universe_id) VALUES(1);
INSERT INTO contract_addresses(
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
		10543755,
		'0x23916a8F5C3846e3100e5f587FF14F3098722F5d',-- augur (main contract)
		'0x63A1eed178323C5eE0aD72fbD8a8cF1a7902881e',-- augur trading
		'0x2c167231cF521AAABc8AbE09F4e2bCB728f26C01',-- profit loss
		'0x6B175474E89094C44Da98b954EedeAC495271d0F',-- dai cash
		'0x221657776846890989a759BA2973e427DfF5C9bB',-- rep token
		'0x8346F3074994FD9A813c735D629B257D93780Eed',-- zerox trade
		'0x61935CbDd02287B511119DDb11Aeb42F1593b7Ef',-- zerox exchange
		'0x9Fa160f92A10b431F255BF1a70a1c1e5808E5128',-- wallet registry (v1)
		'0x1dD864Ed6F291b31C86aAF228DB387cd60a20e18',-- wallet registry (v2)
		'0xc42E71b9A6E38DD05cFB51Be6751a4d10d66ba35',-- fill order
		'0xA478c2975Ab1Ea89e8196811F51A7B7Ade33eB11',-- eth exchange
		'0x9e4799ff2023819b1272eee430eadf510eDF85f0',-- share token
		'0x49244BD018Ca9fd1f06ecC07B9E9De773246e5AA'-- universe,
		'0x8a97CBe557F1153b04d4eDbE4ECa0159B8138937'-- CreateOrder
);
INSERT INTO last_block  VALUES(10543754);
INSERT INTO ooconfig(osize_threshold) VALUES(0.0);
