CREATE TRIGGER tx_insert AFTER INSERT on transaction FOR EACH ROW EXECUTE PROCEDURE on_tx_insert();
CREATE TRIGGER tx_delete AFTER DELETE on transaction FOR EACH ROW EXECUTE PROCEDURE on_tx_delete();
CREATE TRIGGER mesh_evt_insert AFTER INSERT on mesh_evt FOR EACH ROW EXECUTE PROCEDURE on_mesh_evt_insert();
CREATE TRIGGER mesh_evt_delete BEFORE DELETE on mesh_evt FOR EACH ROW EXECUTE PROCEDURE on_mesh_evt_delete();
