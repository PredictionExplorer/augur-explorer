CREATE TRIGGER bjoin_insert AFTER INSERT on bjoin FOR EACH ROW EXECUTE PROCEDURE on_bjoin_insert();
CREATE TRIGGER bjoin_delete AFTER DELETE on bjoin FOR EACH ROW EXECUTE PROCEDURE on_bjoin_delete();
CREATE TRIGGER bexit_insert AFTER INSERT on bexit FOR EACH ROW EXECUTE PROCEDURE on_bexit_insert();
CREATE TRIGGER bexit_delete AFTER DELETE on bexit FOR EACH ROW EXECUTE PROCEDURE on_bexit_delete();
CREATE TRIGGER bbind_insert AFTER INSERT on bbind FOR EACH ROW EXECUTE PROCEDURE on_bbind_insert();
CREATE TRIGGER bbind_delete AFTER DELETE on bbind FOR EACH ROW EXECUTE PROCEDURE on_bbind_delete();
