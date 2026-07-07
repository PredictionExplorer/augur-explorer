-- +goose Up
-- Reorg-path fixes for the RandomWalk marketplace trigger, found by the
-- rw-etl reorg simulation test (TestReorgRollbackAndReplay). The delete
-- trigger must be a true inverse of on_item_bought_insert(), otherwise a
-- chain reorg that rolls back a sale corrupts the aggregates permanently:
--
--   1. It referenced NEW.contract_aid / NEW.seller_aid / NEW.buyer_aid — in a
--      DELETE trigger NEW is null, so the rw_mkt_stats volume/trade reversal
--      and the rw_user_rwtok updates silently matched no rows and the
--      replayed fork double-counted market volume and trade counts.
--   2. It restored the seller's rw_user_rwtok.price_bought to the SALE price
--      instead of the original purchase price, so the replayed sale computed
--      profit = price - price = 0 instead of the true profit.
--   3. It never reversed the profit bookkeeping (rw_new_offer.profit,
--      rw_user_stats.total_profit).
--   4. It left buyer_aid/seller_aid pointing at the trade counterparty; the
--      pre-sale value for the open side of a marketplace offer is always the
--      zero address.
--
-- The original purchase price is recovered from the recorded profit
-- (price_bought = sale price - profit), which is exactly how the insert
-- trigger derived the profit.

-- +goose StatementBegin
CREATE OR REPLACE FUNCTION on_item_bought_delete() RETURNS trigger AS  $$
DECLARE
	v_rwalk_aid				BIGINT;
	v_token_id				BIGINT;
	v_buyer_aid				BIGINT;
	v_seller_aid			BIGINT;
	v_price					DECIMAL;
	v_offer_type			SMALLINT;
	v_profit				DECIMAL;
	v_zero_aid				BIGINT;
BEGIN

	SELECT rwalk_aid,token_id,buyer_aid,seller_aid,price,otype,profit
		FROM rw_new_offer
		WHERE offer_id=OLD.offer_id AND contract_aid=OLD.contract_aid
		INTO v_rwalk_aid,v_token_id,v_buyer_aid,v_seller_aid,v_price,v_offer_type,v_profit;
	IF v_rwalk_aid IS NULL THEN
		RETURN OLD;
	END IF;
	UPDATE rw_stats
		SET		total_vol = (total_vol - v_price),
				total_num_trades = (total_num_trades -1)
		WHERE rwalk_aid = v_rwalk_aid;
	UPDATE rw_mkt_stats
		SET		total_vol = (total_vol - v_price),
				total_num_trades = (total_num_trades -1)
		WHERE contract_aid = OLD.contract_aid;

	UPDATE rw_new_offer SET active=TRUE, profit=NULL
		WHERE offer_id=OLD.offer_id AND contract_aid=OLD.contract_aid;

	-- The open side of the offer was the zero address before the sale.
	SELECT address_id FROM address
		WHERE addr='0x0000000000000000000000000000000000000000' INTO v_zero_aid;
	IF v_zero_aid IS NOT NULL THEN
		IF v_offer_type = 1 THEN
			UPDATE rw_new_offer SET buyer_aid=v_zero_aid
				WHERE offer_id=OLD.offer_id AND contract_aid=OLD.contract_aid;
		ELSE
			UPDATE rw_new_offer SET seller_aid=v_zero_aid
				WHERE offer_id=OLD.offer_id AND contract_aid=OLD.contract_aid;
		END IF;
	END IF;

	UPDATE rw_token SET
			num_trades=(num_trades - 1),
			total_vol=(total_vol-v_price)
		WHERE token_id=v_token_id AND rwalk_aid=v_rwalk_aid;
	UPDATE rw_user_stats SET
			total_num_trades = (total_num_trades-1),
			total_vol = (total_vol-v_price)
		WHERE user_aid=v_buyer_aid AND rwalk_aid=v_rwalk_aid;
	if v_buyer_aid != v_seller_aid THEN 
		UPDATE rw_user_stats SET
				total_num_trades = (total_num_trades-1),
				total_vol = (total_vol-v_price)
			WHERE user_aid=v_seller_aid AND rwalk_aid=v_rwalk_aid;
	END IF;
	IF v_offer_type = 1 THEN
		UPDATE rw_mkt_stats SET
				total_sell_orders = (total_sell_orders + 1)
			WHERE contract_aid=OLD.contract_aid;
	ELSE
		UPDATE rw_mkt_stats SET
				total_buy_orders = (total_buy_orders + 1)
			WHERE contract_aid=OLD.contract_aid;
	END IF;

	IF v_profit IS NOT NULL THEN
		UPDATE rw_user_stats SET total_profit = (total_profit - v_profit)
			WHERE user_aid=v_seller_aid AND rwalk_aid=v_rwalk_aid;
	END IF;
	UPDATE rw_user_rwtok
		SET price_bought = (CASE WHEN v_profit IS NOT NULL THEN v_price - v_profit ELSE NULL END)
		WHERE rwalk_aid=v_rwalk_aid AND user_aid=v_seller_aid AND token_id=v_token_id;
	UPDATE rw_user_rwtok
		SET price_bought = NULL
		WHERE rwalk_aid=v_rwalk_aid AND user_aid=v_buyer_aid AND token_id=v_token_id;
	RETURN OLD;
END;
$$ LANGUAGE plpgsql;
-- +goose StatementEnd

-- +goose Down
-- Restore the original (broken) body from migration 00003.

-- +goose StatementBegin
CREATE OR REPLACE FUNCTION on_item_bought_delete() RETURNS trigger AS  $$
DECLARE
	v_rwalk_aid				BIGINT;
	v_token_id				BIGINT;
	v_buyer_aid				BIGINT;
	v_seller_aid			BIGINT;
	v_price					DECIMAL;
	v_offer_type			SMALLINT;
BEGIN

	SELECT rwalk_aid,token_id,buyer_aid,seller_aid,price,otype
		FROM rw_new_offer
		WHERE offer_id=OLD.offer_id AND contract_aid=OLD.contract_aid
		INTO v_rwalk_aid,v_token_id,v_buyer_aid,v_seller_aid,v_price,v_offer_type;
	IF v_rwalk_aid IS NULL THEN
		RETURN OLD;
	END IF;
	UPDATE rw_stats
		SET		total_vol = (total_vol - v_price),
				total_num_trades = (total_num_trades -1)
		WHERE rwalk_aid = v_rwalk_aid;
	UPDATE rw_mkt_stats
		SET		total_vol = (total_vol - v_price),
				total_num_trades = (total_num_trades -1)
		WHERE contract_aid = NEW.contract_aid;

	UPDATE rw_new_offer SET active=TRUE WHERE offer_id=OLD.offer_id AND contract_aid=OLD.contract_aid;
	UPDATE rw_token SET
			num_trades=(num_trades - 1),
			total_vol=(total_vol-v_price)
		WHERE token_id=v_token_id AND rwalk_aid=v_rwalk_aid;
	UPDATE rw_user_stats SET
			total_num_trades = (total_num_trades-1),
			total_vol = (total_vol-v_price)
		WHERE user_aid=v_buyer_aid AND rwalk_aid=v_rwalk_aid;
	if v_buyer_aid != v_seller_aid THEN 
		UPDATE rw_user_stats SET
				total_num_trades = (total_num_trades-1),
				total_vol = (total_vol-v_price)
			WHERE user_aid=v_seller_aid AND rwalk_aid=v_rwalk_aid;
	END IF;
	IF v_offer_type = 1 THEN
		UPDATE rw_mkt_stats SET
				total_sell_orders = (total_sell_orders + 1)
			WHERE contract_aid=OLD.contract_aid;
	ELSE
		UPDATE rw_mkt_stats SET
				total_buy_orders = (total_buy_orders + 1)
			WHERE contract_aid=OLD.contract_aid;
	END IF;
	UPDATE rw_user_rwtok
		SET price_bought = v_price
		WHERE rwalk_aid=v_rwalk_aid AND user_aid=NEW.seller_aid AND token_id=v_token_id;
	UPDATE rw_user_rwtok
		SET price_bought = NULL
		WHERE rwalk_aid=v_rwalk_aid AND user_aid=NEW.buyer_aid AND token_id=v_token_id;
	RETURN OLD;
END;
$$ LANGUAGE plpgsql;
-- +goose StatementEnd
