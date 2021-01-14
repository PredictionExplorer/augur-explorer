CREATE TRIGGER ens_name_insert_before BEFORE INSERT on ens_name
	FOR EACH ROW EXECUTE PROCEDURE on_ens_name_insert_before();
CREATE TRIGGER ens_name_insert_after AFTER INSERT on ens_name
	FOR EACH ROW EXECUTE PROCEDURE on_ens_name_insert_after();
CREATE TRIGGER ens_name_delete AFTER DELETE on ens_name
	FOR EACH ROW EXECUTE PROCEDURE on_ens_name_delete();
CREATE TRIGGER ens_new_owner_insert AFTER INSERT on ens_new_owner
	FOR EACH ROW EXECUTE PROCEDURE on_ens_new_owner_insert();
CREATE TRIGGER ens_reg_transfer_insert AFTER INSERT on ens_reg_transf
	FOR EACH ROW EXECUTE PROCEDURE on_ens_reg_transfer_insert();
CREATE TRIGGER ens_reg_transfer_delete AFTER DELETE on ens_reg_transf
	FOR EACH ROW EXECUTE PROCEDURE on_ens_reg_transfer_delete();
CREATE TRIGGER ens_text_insert_after AFTER INSERT on ens_text_chg
	FOR EACH ROW EXECUTE PROCEDURE on_ens_text_chg_insert();
CREATE TRIGGER ens_text_delete AFTER DELETE on ens_text_chg
	FOR EACH ROW EXECUTE PROCEDURE on_ens_text_chg_delete();
