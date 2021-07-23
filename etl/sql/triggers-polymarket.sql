CREATE TRIGGER fpmm_funds_add_insert AFTER INSERT ON pol_fund_add FOR EACH ROW EXECUTE PROCEDURE on_fpmm_fund_add_insert();
CREATE TRIGGER fpmm_funds_add_delete AFTER DELETE ON pol_fund_add FOR EACH ROW EXECUTE PROCEDURE on_fpmm_fund_add_delete();
CREATE TRIGGER fpmm_funds_rem_insert AFTER INSERT ON pol_fund_rem FOR EACH ROW EXECUTE PROCEDURE on_fpmm_fund_rem_insert();
CREATE TRIGGER fpmm_funds_rem_delete AFTER DELETE ON pol_fund_rem FOR EACH ROW EXECUTE PROCEDURE on_fpmm_fund_rem_delete();
CREATE TRIGGER fpmm_funds_buy_insert AFTER INSERT ON pol_buy FOR EACH ROW EXECUTE PROCEDURE on_fpmm_buy_insert();
CREATE TRIGGER fpmm_funds_buy_delete AFTER DELETE ON pol_buy FOR EACH ROW EXECUTE PROCEDURE on_fpmm_buy_delete();
CREATE TRIGGER fpmm_funds_sell_insert AFTER INSERT ON pol_sell FOR EACH ROW EXECUTE PROCEDURE on_fpmm_sell_insert();
CREATE TRIGGER fpmm_funds_sell_delete AFTER DELETE ON pol_sell FOR EACH ROW EXECUTE PROCEDURE on_fpmm_sell_delete();
