CREATE TRIGGER dai_transf_insert AFTER INSERT on dai_transf FOR EACH ROW EXECUTE PROCEDURE on_dai_transf_insert();
CREATE TRIGGER dai_transf_delete AFTER DELETE on dai_transf FOR EACH ROW EXECUTE PROCEDURE on_dai_transf_delete();
CREATE TRIGGER dai_bal_insert AFTER INSERT on dai_bal FOR EACH ROW EXECUTE PROCEDURE on_dai_bal_insert();
CREATE TRIGGER dai_bal_delete AFTER DELETE on dai_bal FOR EACH ROW EXECUTE PROCEDURE on_dai_bal_delete();
CREATE TRIGGER dai_bal_update AFTER UPDATE ON dai_bal FOR EACH ROW EXECUTE PROCEDURE on_dai_bal_update();
