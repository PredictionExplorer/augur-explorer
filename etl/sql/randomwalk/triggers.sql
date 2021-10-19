CREATE TRIGGER rw_item_bougnt_insert AFTER INSERT ON rw_item_bought FOR EACH ROW EXECUTE PROCEDURE on_item_bought_insert();
CREATE TRIGGER rw_item_bought_delete AFTER DELETE ON rw_item_bought FOR EACH ROW EXECUTE PROCEDURE on_item_bought_delete();

