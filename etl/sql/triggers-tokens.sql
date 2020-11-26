CREATE TRIGGER dai_transf_insert AFTER INSERT on dai_transf FOR EACH ROW EXECUTE PROCEDURE on_dai_transf_insert();
CREATE TRIGGER dai_transf_delete AFTER DELETE on dai_transf FOR EACH ROW EXECUTE PROCEDURE on_dai_transf_delete();
CREATE TRIGGER dai_bal_insert AFTER INSERT on dai_bal FOR EACH ROW EXECUTE PROCEDURE on_dai_bal_insert();
CREATE TRIGGER dai_bal_delete AFTER DELETE on dai_bal FOR EACH ROW EXECUTE PROCEDURE on_dai_bal_delete();
CREATE TRIGGER dai_bal_update AFTER UPDATE ON dai_bal FOR EACH ROW EXECUTE PROCEDURE on_dai_bal_update();
CREATE TRIGGER af_wrapper_insert AFTER INSERT on af_wrapper FOR EACH ROW EXECUTE PROCEDURE on_af_wrapper_insert();
CREATE TRIGGER af_wrapper_delete AFTER DELETE on af_wrapper FOR EACH ROW EXECUTE PROCEDURE on_af_wrapper_delete();
