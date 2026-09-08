-- +goose Up
-- Match ECMAScript String.trim(), including BOM but excluding U+0085. Keep
-- this predicate shared by the message query and its partial index.
-- +goose StatementBegin
CREATE FUNCTION cg_chat_has_text(value TEXT) RETURNS BOOLEAN
LANGUAGE SQL IMMUTABLE PARALLEL SAFE AS $$
    SELECT COALESCE(btrim(value, U&'\0009\000A\000B\000C\000D\0020\00A0\1680\2000\2001\2002\2003\2004\2005\2006\2007\2008\2009\200A\2028\2029\202F\205F\3000\FEFF') <> '', FALSE)
$$;
-- +goose StatementEnd

-- A durable invalidation generation, not a count/hash over the entire feed.
-- Ordinary appended events preserve older cursors. Edits, reorganizations
-- and moderation changes invalidate cached history in the same transaction.
CREATE TABLE cg_chat_revision (
    singleton BOOLEAN PRIMARY KEY DEFAULT TRUE CHECK (singleton),
    revision BIGINT NOT NULL DEFAULT 1 CHECK (revision > 0)
);
INSERT INTO cg_chat_revision (singleton) VALUES (TRUE);

-- +goose StatementBegin
CREATE FUNCTION cg_invalidate_chat_history() RETURNS TRIGGER
LANGUAGE plpgsql AS $$
BEGIN
    UPDATE cg_chat_revision SET revision = revision + 1 WHERE singleton;
    RETURN NULL;
END;
$$;
-- +goose StatementEnd

-- Row triggers do not invalidate on the indexer's routine no-op DELETE
-- before inserting a new event. Reprocessing an existing row does invalidate.
CREATE TRIGGER cg_chat_bid_changed AFTER UPDATE OR DELETE ON cg_bid
FOR EACH ROW EXECUTE FUNCTION cg_invalidate_chat_history();
CREATE TRIGGER cg_chat_ban_changed AFTER INSERT OR UPDATE OR DELETE ON cg_banned_bids
FOR EACH ROW EXECUTE FUNCTION cg_invalidate_chat_history();

-- +goose Down
DROP TRIGGER cg_chat_ban_changed ON cg_banned_bids;
DROP TRIGGER cg_chat_bid_changed ON cg_bid;
DROP FUNCTION cg_invalidate_chat_history();
DROP TABLE cg_chat_revision;
DROP FUNCTION cg_chat_has_text(TEXT);
