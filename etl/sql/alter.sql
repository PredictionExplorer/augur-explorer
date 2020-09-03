ALTER TABLE outcome_vol ADD highest_bid DECIMAL(64,18) DEFAULT 0.0;
ALTER TABLE outcome_vol ADD lowest_ask DECIMAL(64,18) DEFAULT 0.0;
ALTER TABLE outcome_vol ADD cur_spread DECIMAL(64,18) DEFAULT 0.0;
ALTER TABLE outcome_vol ADD price_estimate DECIMAL(64,18) DEFAULT 0.0;
DELETE FROM oorders;
ALTER TABLE oorders ADD opcode SMALLINT NOT NULL;
ALTER TABLE mktord RENAME COLUMN order_id TO order_hash;
ALTER TABLE oorders RENAME COLUMN order_id TO order_hash;
