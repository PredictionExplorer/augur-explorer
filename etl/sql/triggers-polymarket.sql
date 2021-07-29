CREATE TRIGGER fpmm_funds_addrem_insert AFTER INSERT ON pol_fund_addrem FOR EACH ROW EXECUTE PROCEDURE on_fpmm_fund_addrem_insert();
CREATE TRIGGER fpmm_funds_addrem_delete AFTER DELETE ON pol_fund_addrem FOR EACH ROW EXECUTE PROCEDURE on_fpmm_fund_addrem_delete();
CREATE TRIGGER fpmm_buysell_insert AFTER INSERT ON pol_buysell FOR EACH ROW EXECUTE PROCEDURE on_fpmm_buysell_insert();
CREATE TRIGGER fpmm_buysell_delete AFTER DELETE ON pol_buysell FOR EACH ROW EXECUTE PROCEDURE on_fpmm_buysell_delete();
