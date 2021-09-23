CREATE TRIGGER erc20_transf_insert AFTER INSERT ON erc20_transf FOR EACH ROW EXECUTE PROCEDURE on_erc20_transf_insert();
CREATE TRIGGER erc20_transf_delete AFTER DELETE ON erc20_transf FOR EACH ROW EXECUTE PROCEDURE on_erc20_transf_delete();
CREATE TRIGGER erc20_bal_insert AFTER INSERT ON erc20_bal FOR EACH ROW EXECUTE PROCEDURE on_erc20_bal_insert();
CREATE TRIGGER erc20_bal_delete AFTER DELETE ON erc20_bal FOR EACH ROW EXECUTE PROCEDURE on_erc20_bal_delete();
