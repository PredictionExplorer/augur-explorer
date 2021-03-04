CREATE TRIGGER ens_name_reg1_insert_before BEFORE INSERT on ens_name_reg1
	FOR EACH ROW EXECUTE PROCEDURE on_ens_name_reg1_insert_before();
CREATE TRIGGER ens_name_reg2_insert_before BEFORE INSERT on ens_name_reg2
	FOR EACH ROW EXECUTE PROCEDURE on_ens_name_reg2_insert_before();
CREATE TRIGGER ens_new_owner_insert AFTER INSERT on ens_new_owner
	FOR EACH ROW EXECUTE PROCEDURE on_ens_new_owner_insert();
CREATE TRIGGER ens_addr1 AFTER INSERT on ens_addr1
	FOR EACH ROW EXECUTE PROCEDURE on_ens_addr_changed1_insert();
CREATE TRIGGER ens_addr2 AFTER INSERT on ens_addr2
	FOR EACH ROW EXECUTE PROCEDURE on_ens_addr_changed2_insert();
CREATE TRIGGER ens_reg_transfer_insert AFTER INSERT on ens_reg_transf
	FOR EACH ROW EXECUTE PROCEDURE on_ens_reg_transfer_insert();
CREATE TRIGGER ens_reg_transfer_delete AFTER DELETE on ens_reg_transf
	FOR EACH ROW EXECUTE PROCEDURE on_ens_reg_transfer_delete();
CREATE TRIGGER ens_text_insert AFTER INSERT on ens_text_chg
	FOR EACH ROW EXECUTE PROCEDURE on_ens_text_chg_insert();
CREATE TRIGGER ens_text_delete AFTER DELETE on ens_text_chg
	FOR EACH ROW EXECUTE PROCEDURE on_ens_text_chg_delete();
CREATE TRIGGER ens_text_key_insert AFTER INSERT on ens_text_key
	FOR EACH ROW EXECUTE PROCEDURE on_ens_text_key_insert();
CREATE TRIGGER ens_text_key_delete AFTER DELETE on ens_text_key
	FOR EACH ROW EXECUTE PROCEDURE on_ens_text_key_delete();
