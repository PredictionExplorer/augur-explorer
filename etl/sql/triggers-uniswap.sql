CREATE TRIGGER uswap1_insert AFTER INSERT on uswap1 FOR EACH ROW EXECUTE PROCEDURE on_uswap1_insert();
CREATE TRIGGER uswap1_delete AFTER DELETE on uswap1 FOR EACH ROW EXECUTE PROCEDURE on_uswap1_delete();
