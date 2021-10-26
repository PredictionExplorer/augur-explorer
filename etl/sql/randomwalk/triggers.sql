CREATE TRIGGER rw_item_bougnt_insert AFTER INSERT ON rw_item_bought FOR EACH ROW EXECUTE PROCEDURE on_item_bought_insert();
CREATE TRIGGER rw_item_bought_delete AFTER DELETE ON rw_item_bought FOR EACH ROW EXECUTE PROCEDURE on_item_bought_delete();
CREATE TRIGGER rw_token_name_insert AFTER INSERT ON rw_token_name FOR EACH ROW EXECUTE PROCEDURE on_token_name_insert();
CREATE TRIGGER rw_token_name_delete AFTER DELETE ON rw_token_name FOR EACH ROW EXECUTE PROCEDURE on_token_name_delete();
CREATE TRIGGER rw_mint_event_insert AFTER INSERT ON rw_mint_evt FOR EACH ROW EXECUTE PROCEDURE on_mint_event_insert();
CREATE TRIGGER rw_mint_event_delete AFTER DELETE ON rw_mint_evt FOR EACH ROW EXECUTE PROCEDURE on_mint_event_delete();
CREATE TRIGGER rw_offer_canceled_insert AFTER INSERT ON rw_offer_canceled FOR EACH ROW EXECUTE PROCEDURE on_offer_canceled_insert();
CREATE TRIGGER rw_offer_canceled_delete AFTER DELETE ON rw_offer_canceled FOR EACH ROW EXECUTE PROCEDURE on_offer_canceled_delete();
