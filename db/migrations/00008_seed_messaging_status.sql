-- +goose Up
-- rw_messaging_status (the notibot notification watermark) was created
-- without a row, and Update_messaging_status uses a plain UPDATE, which
-- silently affects zero rows on an empty table. On a freshly migrated
-- database the watermark therefore never persisted and every notibot restart
-- re-notified the entire event history to Twitter/Discord. Legacy production
-- databases were seeded manually; make the migrations self-sufficient
-- (same defect family as last_block, migration 00005). Found by the store
-- read suite.
INSERT INTO rw_messaging_status(last_tx_id, last_evtlog_id, last_block_num, last_timestamp)
SELECT 0, 0, 0, 0
WHERE NOT EXISTS (SELECT 1 FROM rw_messaging_status);

-- +goose Down
-- Removing the seed row would break the UPDATE-based writer again; the Down
-- migration intentionally leaves the row in place.
SELECT 1;
