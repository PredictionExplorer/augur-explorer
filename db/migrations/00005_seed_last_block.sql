-- +goose Up
-- The last_block watermark table was created without a row. Every writer
-- (Set_last_block_num, and Insert_block's watermark maintenance) uses a plain
-- UPDATE, which silently affects zero rows on an empty table, so on a freshly
-- migrated database the watermark never advanced and chain-split handling
-- (HandleChainSplit reads Get_last_block_num to know how far to roll back)
-- was inert. Legacy production databases were seeded manually; make the
-- migrations self-sufficient. Found by the ETL fixture suite.
INSERT INTO last_block(block_num)
SELECT 0
WHERE NOT EXISTS (SELECT 1 FROM last_block);

-- +goose Down
-- Removing the seed row would break the UPDATE-based writers again; the Down
-- migration intentionally leaves the row in place.
SELECT 1;
