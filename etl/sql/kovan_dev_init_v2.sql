-- Do not use, for Development only!
INSERT INTO universe(universe_addr) VALUES('0x39558F07B0123bb9C73c046153B5bed0c8bbc8B5');
INSERT INTO main_stats(universe_id) VALUES(1);
INSERT INTO contract_addresses(
		upload_block,
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
		19520325,
		'0x08f7904Ab81CFA0615C9D73f69F0c17521B36Fb0',-- augur (main contract)
		'0xFe3cfcc677873488937D31370Bf28DB424a82310',-- augur trading
		'0xd6aCE86ba608b012453a70C5d96B904d3cAE6aA3',-- profit loss
		'0xaCA1207624B246952fc18f7BB5D8523aF81e7d05',-- dai cash
		'0x14dF26Dcae09e954b85aA7f60D3b0FfCEDa69A37',-- rep token
		'0x5360148F1e5FA2A96241FE2B0710fe73879aFfAF',-- zerox
		'0x1aD2Fb709e8B22430bfBa302f3Fd8F993877879D',-- wallet registry
		'0xc275B2FEab763d53e9B200bF35f5b88051e23DE9',-- fill order
		'0xB03Cd91915F66c2ABacE5135112F3A4d891429AE',-- eth exchange
		'0xA2821b3F0302927d0cEbeB0C85060672dF397F14',-- share token
		'0x39558F07B0123bb9C73c046153B5bed0c8bbc8B5'-- universe
);
INSERT INTO last_block  VALUES(19520324);
INSERT INTO ooconfig(osize_threshold) VALUES(0.0);
