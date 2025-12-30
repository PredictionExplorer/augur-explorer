-- Archive tables for RandomWalk AND CosmicGame
-- These store historical data that may be pruned from RPC nodes
-- No IDs, no foreign keys - just raw data for recovery

-- Archived blocks
CREATE TABLE arch_block (
    block_num   BIGINT NOT NULL,
    num_tx      BIGINT DEFAULT 0,
    ts          TIMESTAMPTZ NOT NULL,
    cash_flow   NUMERIC(64,18) DEFAULT 0.0,
    block_hash  CHAR(66) NOT NULL PRIMARY KEY,
    parent_hash CHAR(66) NOT NULL
);
CREATE INDEX idx_arch_block_hash ON arch_block(block_hash);

-- Archived transactions
CREATE TABLE arch_tx (
    block_num    BIGINT NOT NULL,
    from_aid     BIGINT DEFAULT 0,
    to_aid       BIGINT DEFAULT 0,
    gas_used     BIGINT DEFAULT 0,
    tx_index     INT DEFAULT 0,
    num_logs     INT DEFAULT 0,
    ctrct_create BOOLEAN DEFAULT FALSE,
    value        NUMERIC(80,18) DEFAULT 0.0,
    gas_price    NUMERIC(80,18) DEFAULT 0.0,
    tx_hash      CHAR(66) NOT NULL PRIMARY KEY,
    input_sig    CHAR(10)
);
CREATE INDEX idx_arch_tx_block ON arch_tx(block_num);

-- Archived event logs
CREATE TABLE arch_evtlog (
    block_num     BIGINT NOT NULL,
    evt_id        BIGINT NOT NULL PRIMARY KEY,
    tx_hash       CHAR(66) NOT NULL,
    contract_addr CHAR(42) NOT NULL,
    topic0_sig    CHAR(8) NOT NULL,
    log_rlp       BYTEA NOT NULL
);
CREATE INDEX idx_arch_evtlog_tx ON arch_evtlog(tx_hash);
CREATE INDEX idx_arch_evtlog_block ON arch_evtlog(block_num);

