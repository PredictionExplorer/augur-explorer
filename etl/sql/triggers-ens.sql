CREATE TRIGGER ens_name_insert_before BEFORE INSERT on ens_name FOR EACH ROW EXECUTE PROCEDURE on_ens_name_insert_before();
CREATE TRIGGER ens_name_insert_after AFTER INSERT on ens_name FOR EACH ROW EXECUTE PROCEDURE on_ens_name_insert_after();
CREATE TRIGGER ens_name_delete AFTER DELETE on ens_name FOR EACH ROW EXECUTE PROCEDURE on_ens_name_delete();
