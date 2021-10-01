CREATE TRIGGER liquidity_changed_insert AFTER INSERT ON aa_liquidity_changed FOR EACH ROW EXECUTE PROCEDURE on_liquidity_changed_insert();
CREATE TRIGGER liquidity_changed_delete AFTER DELETE ON aa_liquidity_changed FOR EACH ROW EXECUTE PROCEDURE on_liquidity_changed_delete();
--CREATE TRIGGER sports_market_insert BEFORE INSERT ON aa_sports_market FOR EACH ROW EXECUTE PROCEDURE on_before_sports_market_insert();
--CREATE TRIGGER sports_market_delete AFTTER DELETE PN aa_sports_market FOR EACH ROW EXECUTE PROCEDURE on_after_sports_market_delete();
