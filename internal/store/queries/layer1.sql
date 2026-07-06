-- Layer-1 (blockchain infrastructure) queries.
-- Converted from hand-written SQL in internal/store; the sqlc-generated code
-- lives in internal/store/sqlcgen. See sqlc.yaml.

-- name: GetAddressID :one
SELECT address_id FROM address WHERE addr = $1;

-- name: InsertAddress :one
INSERT INTO address (block_num, tx_id, addr)
VALUES ($1, $2, $3)
ON CONFLICT (addr) DO UPDATE SET addr = EXCLUDED.addr
RETURNING address_id;

-- name: GetLastBlockNum :one
SELECT block_num FROM last_block LIMIT 1;

-- name: SetLastBlockNum :exec
UPDATE last_block SET block_num = $1;

-- name: GetBlockByHash :one
SELECT block_num, num_tx, ts, block_hash, parent_hash
FROM block
WHERE block_hash = $1;

-- name: InsertBlock :exec
INSERT INTO block (block_num, num_tx, ts, block_hash, parent_hash)
VALUES ($1, $2, TO_TIMESTAMP($3), $4, $5);

-- name: GetTransactionID :one
SELECT id FROM transaction WHERE tx_hash = $1;

-- name: GetEventLogByTxAndIndex :one
SELECT id FROM evt_log WHERE tx_id = $1 AND log_index = $2;
