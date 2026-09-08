-- +goose NO TRANSACTION
-- +goose Up
CREATE INDEX CONCURRENTLY cg_chat_round_time_event_idx
    ON cg_bid (round_num, time_stamp DESC, evtlog_id DESC)
    WHERE cg_chat_has_text(msg);

-- +goose Down
DROP INDEX CONCURRENTLY IF EXISTS cg_chat_round_time_event_idx;
