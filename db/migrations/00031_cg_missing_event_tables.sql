-- +goose Up
-- Two events that no handler decoded until now.
--
-- EthBidRefundAmountInGasToSwallowMaxLimitChanged (ISystemEvents.sol:42,
-- emitted from SystemManagementV2.sol:48) is an owner-settable configuration
-- value reachable in V3 through CosmicSignatureGameAdminModuleV3, which
-- inherits SystemManagementV2. It joins the admin-event UNION as record_type
-- 45 (43 is retired; see migration 00030).
--
-- ArbitrumError (CosmicSignatureEvents.sol:30, emitted from
-- ArbitrumHelpers.sol:72) reports that one of the four Arbitrum precompile
-- reads feeding the random-number seed failed, so the seed was built from
-- fewer entropy sources than intended. It is informational only and is
-- deliberately absent from the API.

CREATE TABLE cg_adm_eth_bid_refund_gas_limit ( -- ISystemEvents.sol:EthBidRefundAmountInGasToSwallowMaxLimitChanged
	id				BIGSERIAL PRIMARY KEY,
	evtlog_id		BIGINT REFERENCES evt_log(id) ON DELETE CASCADE,
	block_num		BIGINT NOT NULL,
	tx_id			BIGINT NOT NULL,
	time_stamp		TIMESTAMPTZ NOT NULL,
	contract_aid	BIGINT NOT NULL,
	new_value		DECIMAL NOT NULL,
	UNIQUE (evtlog_id)
);

CREATE TABLE cg_arbitrum_error ( -- CosmicSignatureEvents.sol:ArbitrumError
	id				BIGSERIAL PRIMARY KEY,
	evtlog_id		BIGINT REFERENCES evt_log(id) ON DELETE CASCADE,
	block_num		BIGINT NOT NULL,
	tx_id			BIGINT NOT NULL,
	time_stamp		TIMESTAMPTZ NOT NULL,
	contract_aid	BIGINT NOT NULL,
	err_str			TEXT NOT NULL,
	UNIQUE (evtlog_id)
);

-- A precompile failure repeats on every claim once it starts, so the common
-- query is "which failures happened, most recent first".
CREATE INDEX cg_arbitrum_error_block_idx ON cg_arbitrum_error(block_num DESC);

-- +goose Down
DROP TABLE cg_arbitrum_error;
DROP TABLE cg_adm_eth_bid_refund_gas_limit;
