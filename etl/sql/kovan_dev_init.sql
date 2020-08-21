-- Do not use, for Development only!
INSERT INTO universe(universe_addr) VALUES('0x48CB0042B1c0AE139Fd6b83481613a10Af4a32A8');
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
		19830392,
		'0xd50475cC7Ae9dd46a09D8415A57F0DF59E89b0f9',-- augur (main contract)
		'0x9307C37077fd428f18f2a4E94BaeeeA7Cc5681dd',-- augur trading
		'0x5B659733415f9919D5d4FA8e7723D7b0e6434E4F',-- profit loss
		'0x5bc0ceB4d813282C4468288C8AB42Af5e7a43B27',-- dai cash
		'0x7859eF3B6D0Be006675990192Eea09cba46d9C35',-- rep token
		'0x5673627689E6435D3Bea15A7B94cf6E5D2d450E3',-- zerox
		'0x29fD93b1Ebb775306b018BdDd107e2F5c889e70c',-- wallet registry
		'0x58f1a6a327dC8A4c73F72B37B75802EB24670A94',-- fill order
		'0xA867A0291D23AC82bB9eA83DD52ee8B71922Fb42',-- eth exchange
		'0xfB4B6fc159E2eB3Ab7e2E08f0056238488c6D5AB',-- share token
		'0x48CB0042B1c0AE139Fd6b83481613a10Af4a32A8'-- universe
);
INSERT INTO last_block  VALUES(19830391);
INSERT INTO ooconfig(osize_threshold) VALUES(0.0);
