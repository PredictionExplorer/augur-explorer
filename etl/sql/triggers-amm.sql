CREATE TRIGGER liquidity_changed_insert AFTER INSERT ON aa_liquidity_changed FOR EACH ROW EXECUTE PROCEDURE on_liquidity_changed_insert();
CREATE TRIGGER liquidity_changed_delete AFTER DELETE ON aa_liquidity_changed FOR EACH ROW EXECUTE PROCEDURE on_liquidity_changed_delete();
