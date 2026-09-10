-- +goose Up
-- V3.1 (Solidity commit 3bbd92e0) replaced the game's charity-transfer
-- failure emission of FundTransferFailed(string,address,uint256) with a
-- dedicated EthTransferToCharityFailed(address indexed,uint256) event.
-- Mirror of cg_fund_transf_err, which keeps serving the already-deployed
-- CharityWallet / StakingWalletCosmicSignatureNft emitters.
CREATE TABLE cg_eth_to_charity_failed ( -- CosmicSignatureEvents.sol:EthTransferToCharityFailed
	id              BIGSERIAL PRIMARY KEY,
	evtlog_id       BIGINT REFERENCES evt_log(id) ON DELETE CASCADE,
	block_num       BIGINT NOT NULL,
	tx_id           BIGINT NOT NULL,
	time_stamp      TIMESTAMPTZ NOT NULL,
	contract_aid    BIGINT NOT NULL,
	charity_aid     BIGINT NOT NULL,
	amount          DECIMAL NOT NULL,
	UNIQUE(evtlog_id)
);
CREATE INDEX cg_eth_to_charity_failed_block_idx ON cg_eth_to_charity_failed(block_num DESC);

-- +goose Down
DROP TABLE cg_eth_to_charity_failed;
