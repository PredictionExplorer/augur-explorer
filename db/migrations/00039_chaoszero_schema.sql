-- +goose Up
-- ChaosZero schema (cz_ prefix): the eight-market IndependentOutcomeMarket
-- and its OutcomeSettlementVault on Arbitrum One, indexed by cmd/cz-etl.
-- The contracts are immutable singletons (no proxy, no owner); a future
-- redeployment is a new row-less address change in cz_contracts, so every
-- event table carries contract_aid.

-- Contract registry (one row), production deployment of 21 October 2026.
-- The vault is deployed by the market's constructor; its address is pinned
-- here so the ETL needs no RPC read at startup. The outcome source emits no
-- events but is registered for address joins.
CREATE TABLE cz_contracts (
	market_addr		TEXT DEFAULT '0x2BC4759A2b44e1197EAdc4E035Eb8c52299658d2',
	source_addr		TEXT DEFAULT '0x099c29624E05506582655417344A17FB71c9d138',
	vault_addr		TEXT DEFAULT '0x60F79aF7378Dc06F7747827446b32157c796dd3e'
);
INSERT INTO cz_contracts DEFAULT VALUES;

-- ETL watermark. last_block defaults to the block before the market
-- deployment transaction (0x9360e041..., block 517037317) so a fresh
-- database starts scanning at deployment instead of block zero.
CREATE TABLE cz_proc_status (
	last_evt_id		BIGINT DEFAULT 0,
	last_block		BIGINT DEFAULT 517037316
);
INSERT INTO cz_proc_status DEFAULT VALUES;

-- IndependentOutcomeMarket.RoundInitialized(uint256 indexed roundId)
CREATE TABLE cz_round_initialized (
	id				BIGSERIAL PRIMARY KEY,
	evtlog_id		BIGINT REFERENCES evt_log(id) ON DELETE CASCADE,
	block_num		BIGINT NOT NULL,
	tx_id			BIGINT NOT NULL,
	time_stamp		TIMESTAMPTZ NOT NULL,
	contract_aid	BIGINT NOT NULL,
	round_id		BIGINT NOT NULL,
	UNIQUE(evtlog_id)
);

-- IndependentOutcomeMarket.RoundFunded: caller's own CST (donated), swept
-- unsolicited transfers and applied carry; allocations[8] is the per-pool
-- split (their sum equals donated+transfers_used+carry_used).
CREATE TABLE cz_round_funded (
	id				BIGSERIAL PRIMARY KEY,
	evtlog_id		BIGINT REFERENCES evt_log(id) ON DELETE CASCADE,
	block_num		BIGINT NOT NULL,
	tx_id			BIGINT NOT NULL,
	time_stamp		TIMESTAMPTZ NOT NULL,
	contract_aid	BIGINT NOT NULL,
	round_id		BIGINT NOT NULL,
	caller_aid		BIGINT NOT NULL,
	donated			DECIMAL NOT NULL,
	transfers_used	DECIMAL NOT NULL,
	carry_used		DECIMAL NOT NULL,
	allocations		NUMERIC[] NOT NULL,		-- 8 elements, one per question
	UNIQUE(evtlog_id)
);

-- IndependentOutcomeMarket.Trade: every buy/sell on one of the eight
-- binary questions. cash is CST in (buy) or out (sell); units the position
-- tokens out (buy) or in (sell).
CREATE TABLE cz_trade (
	id				BIGSERIAL PRIMARY KEY,
	evtlog_id		BIGINT REFERENCES evt_log(id) ON DELETE CASCADE,
	block_num		BIGINT NOT NULL,
	tx_id			BIGINT NOT NULL,
	time_stamp		TIMESTAMPTZ NOT NULL,
	contract_aid	BIGINT NOT NULL,
	round_id		BIGINT NOT NULL,
	user_aid		BIGINT NOT NULL,
	question		SMALLINT NOT NULL,		-- 0..7
	yes				BOOLEAN NOT NULL,		-- YES or NO side
	buy				BOOLEAN NOT NULL,		-- buy or sell
	cash			DECIMAL NOT NULL,		-- CST amount
	units			DECIMAL NOT NULL,		-- position-token amount
	UNIQUE(evtlog_id)
);

-- IndependentOutcomeMarket.PositionSplit: complete-set split on one question.
CREATE TABLE cz_position_split (
	id				BIGSERIAL PRIMARY KEY,
	evtlog_id		BIGINT REFERENCES evt_log(id) ON DELETE CASCADE,
	block_num		BIGINT NOT NULL,
	tx_id			BIGINT NOT NULL,
	time_stamp		TIMESTAMPTZ NOT NULL,
	contract_aid	BIGINT NOT NULL,
	round_id		BIGINT NOT NULL,
	user_aid		BIGINT NOT NULL,
	question		SMALLINT NOT NULL,
	amount			DECIMAL NOT NULL,
	UNIQUE(evtlog_id)
);

-- IndependentOutcomeMarket.PositionsMerged: matched YES/NO pairs returned
-- to CST across all eight questions in one call.
CREATE TABLE cz_positions_merged (
	id				BIGSERIAL PRIMARY KEY,
	evtlog_id		BIGINT REFERENCES evt_log(id) ON DELETE CASCADE,
	block_num		BIGINT NOT NULL,
	tx_id			BIGINT NOT NULL,
	time_stamp		TIMESTAMPTZ NOT NULL,
	contract_aid	BIGINT NOT NULL,
	round_id		BIGINT NOT NULL,
	user_aid		BIGINT NOT NULL,
	amounts			NUMERIC[] NOT NULL,		-- 8 elements
	collateral_out	DECIMAL NOT NULL,
	UNIQUE(evtlog_id)
);

-- IndependentOutcomeMarket.RoundSettled: cycle resolution. values_mask bit q
-- is question q's outcome; answers is the same mask unpacked at store time.
CREATE TABLE cz_round_settled (
	id				BIGSERIAL PRIMARY KEY,
	evtlog_id		BIGINT REFERENCES evt_log(id) ON DELETE CASCADE,
	block_num		BIGINT NOT NULL,
	tx_id			BIGINT NOT NULL,
	time_stamp		TIMESTAMPTZ NOT NULL,
	contract_aid	BIGINT NOT NULL,
	round_id		BIGINT NOT NULL,
	values_mask		SMALLINT NOT NULL,
	answers			BOOLEAN[] NOT NULL,		-- values_mask unpacked, 8 elements
	escrowed		NUMERIC[] NOT NULL,		-- 8 elements
	carried			DECIMAL NOT NULL,
	UNIQUE(evtlog_id)
);

-- IndependentOutcomeMarket.Claimed: winning-position payout.
CREATE TABLE cz_claim (
	id				BIGSERIAL PRIMARY KEY,
	evtlog_id		BIGINT REFERENCES evt_log(id) ON DELETE CASCADE,
	block_num		BIGINT NOT NULL,
	tx_id			BIGINT NOT NULL,
	time_stamp		TIMESTAMPTZ NOT NULL,
	contract_aid	BIGINT NOT NULL,
	round_id		BIGINT NOT NULL,
	user_aid		BIGINT NOT NULL,
	amounts			NUMERIC[] NOT NULL,		-- 8 elements
	total			DECIMAL NOT NULL,
	UNIQUE(evtlog_id)
);

-- IndependentOutcomeMarket.AnswersObserved: partial outcomes known mid-round.
CREATE TABLE cz_answers_observed (
	id				BIGSERIAL PRIMARY KEY,
	evtlog_id		BIGINT REFERENCES evt_log(id) ON DELETE CASCADE,
	block_num		BIGINT NOT NULL,
	tx_id			BIGINT NOT NULL,
	time_stamp		TIMESTAMPTZ NOT NULL,
	contract_aid	BIGINT NOT NULL,
	round_id		BIGINT NOT NULL,
	known_mask		SMALLINT NOT NULL,
	values_mask		SMALLINT NOT NULL,
	UNIQUE(evtlog_id)
);

-- OutcomeSettlementVault.CycleEscrowed: winning-claim collateral reserved
-- at settlement.
CREATE TABLE cz_cycle_escrowed (
	id				BIGSERIAL PRIMARY KEY,
	evtlog_id		BIGINT REFERENCES evt_log(id) ON DELETE CASCADE,
	block_num		BIGINT NOT NULL,
	tx_id			BIGINT NOT NULL,
	time_stamp		TIMESTAMPTZ NOT NULL,
	contract_aid	BIGINT NOT NULL,
	round_id		BIGINT NOT NULL,
	amounts			NUMERIC[] NOT NULL,		-- 8 elements
	UNIQUE(evtlog_id)
);

-- OutcomeSettlementVault.Paid: actual CST payout to a winner.
CREATE TABLE cz_paid (
	id				BIGSERIAL PRIMARY KEY,
	evtlog_id		BIGINT REFERENCES evt_log(id) ON DELETE CASCADE,
	block_num		BIGINT NOT NULL,
	tx_id			BIGINT NOT NULL,
	time_stamp		TIMESTAMPTZ NOT NULL,
	contract_aid	BIGINT NOT NULL,
	round_id		BIGINT NOT NULL,
	question		SMALLINT NOT NULL,
	recipient_aid	BIGINT NOT NULL,
	amount			DECIMAL NOT NULL,
	UNIQUE(evtlog_id)
);

-- ERC-1155 TransferSingle/TransferBatch of position tokens. One row per
-- (id, value) pair; a TransferBatch of n ids yields n rows sharing an
-- evtlog_id, distinguished by batch_index. The position id encodes
-- (roundId << 4) | (question << 1) | yes, decoded into the three columns.
-- Mints have from_aid = zero address, burns to_aid = zero address.
CREATE TABLE cz_position_transfer (
	id				BIGSERIAL PRIMARY KEY,
	evtlog_id		BIGINT REFERENCES evt_log(id) ON DELETE CASCADE,
	batch_index		INT NOT NULL DEFAULT 0,
	block_num		BIGINT NOT NULL,
	tx_id			BIGINT NOT NULL,
	time_stamp		TIMESTAMPTZ NOT NULL,
	contract_aid	BIGINT NOT NULL,
	operator_aid	BIGINT NOT NULL,
	from_aid		BIGINT NOT NULL,
	to_aid			BIGINT NOT NULL,
	position_id		DECIMAL NOT NULL,		-- raw ERC-1155 token id
	round_id		BIGINT NOT NULL,
	question		SMALLINT NOT NULL,
	yes				BOOLEAN NOT NULL,
	amount			DECIMAL NOT NULL,
	UNIQUE(evtlog_id, batch_index)
);

-- ERC-1155 ApprovalForAll on the market.
CREATE TABLE cz_approval_for_all (
	id				BIGSERIAL PRIMARY KEY,
	evtlog_id		BIGINT REFERENCES evt_log(id) ON DELETE CASCADE,
	block_num		BIGINT NOT NULL,
	tx_id			BIGINT NOT NULL,
	time_stamp		TIMESTAMPTZ NOT NULL,
	contract_aid	BIGINT NOT NULL,
	owner_aid		BIGINT NOT NULL,
	operator_aid	BIGINT NOT NULL,
	approved		BOOLEAN NOT NULL,
	UNIQUE(evtlog_id)
);

-- Read indexes for the obvious access paths (per-round activity feeds and
-- per-user history).
CREATE INDEX cz_trade_round_idx				ON cz_trade (round_id, id);
CREATE INDEX cz_trade_user_idx				ON cz_trade (user_aid, id);
CREATE INDEX cz_claim_round_idx				ON cz_claim (round_id);
CREATE INDEX cz_claim_user_idx				ON cz_claim (user_aid);
CREATE INDEX cz_paid_round_idx				ON cz_paid (round_id);
CREATE INDEX cz_paid_recipient_idx			ON cz_paid (recipient_aid);
CREATE INDEX cz_position_transfer_round_idx	ON cz_position_transfer (round_id);
CREATE INDEX cz_position_transfer_from_idx	ON cz_position_transfer (from_aid);
CREATE INDEX cz_position_transfer_to_idx	ON cz_position_transfer (to_aid);

-- +goose Down
DROP TABLE IF EXISTS
	cz_round_initialized,
	cz_round_funded,
	cz_trade,
	cz_position_split,
	cz_positions_merged,
	cz_round_settled,
	cz_claim,
	cz_answers_observed,
	cz_cycle_escrowed,
	cz_paid,
	cz_position_transfer,
	cz_approval_for_all,
	cz_proc_status,
	cz_contracts
CASCADE;
