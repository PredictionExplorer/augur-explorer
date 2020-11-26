CREATE TRIGGER mesh_evt_insert AFTER INSERT on mesh_evt FOR EACH ROW EXECUTE PROCEDURE on_mesh_evt_insert();
CREATE TRIGGER mesh_evt_delete BEFORE DELETE on mesh_evt FOR EACH ROW EXECUTE PROCEDURE on_mesh_evt_delete();
